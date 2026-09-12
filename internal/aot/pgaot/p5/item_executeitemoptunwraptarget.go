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
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
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
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
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
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v374 int32
	_ = v374
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int64
	_ = v516
	var v518 int64
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
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
	var v562 int32
	_ = v562
	var v577 int32
	_ = v577
	var v593 int32
	_ = v593
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
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
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v828 int64
	_ = v828
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v907 int32
	_ = v907
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
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v957 int32
	_ = v957
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v974 int64
	_ = v974
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1078 int32
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1096 int64
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1133 int32
	_ = v1133
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1190 int32
	_ = v1190
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1207 int32
	_ = v1207
	var v1212 int32
	_ = v1212
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1314 int32
	_ = v1314
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1341 int64
	_ = v1341
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
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
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1440 int32
	_ = v1440
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1476 int64
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
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
	var v1499 int32
	_ = v1499
	var v1508 int32
	_ = v1508
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1522 int64
	_ = v1522
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1541 int32
	_ = v1541
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1559 int32
	_ = v1559
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1603 int32
	_ = v1603
	var v1608 int32
	_ = v1608
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1649 int32
	_ = v1649
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1681 int32
	_ = v1681
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1700 int32
	_ = v1700
	var v1705 int32
	_ = v1705
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1727 int32
	_ = v1727
	var v1730 int32
	_ = v1730
	var v1733 int64
	_ = v1733
	var v1734 int64
	_ = v1734
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1782 int32
	_ = v1782
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1858 int32
	_ = v1858
	var v1860 int32
	_ = v1860
	var v1869 int32
	_ = v1869
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1879 int32
	_ = v1879
	var v1885 int32
	_ = v1885
	var v1887 int64
	_ = v1887
	var v1889 int32
	_ = v1889
	var v1894 int32
	_ = v1894
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1904 int64
	_ = v1904
	var v1906 int64
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1922 int32
	_ = v1922
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1941 int32
	_ = v1941
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1949 int32
	_ = v1949
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1966 int32
	_ = v1966
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1996 int32
	_ = v1996
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2023 int32
	_ = v2023
	var v2028 int32
	_ = v2028
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2039 int32
	_ = v2039
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2048 int32
	_ = v2048
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2062 int32
	_ = v2062
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2081 int32
	_ = v2081
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2094 int32
	_ = v2094
	var v2098 int32
	_ = v2098
	var v2101 int64
	_ = v2101
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2111 int32
	_ = v2111
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2136 int32
	_ = v2136
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2148 int32
	_ = v2148
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2156 int32
	_ = v2156
	var v2160 int32
	_ = v2160
	var v2162 int32
	_ = v2162
	var v2166 int64
	_ = v2166
	var v2170 int32
	_ = v2170
	var v2192 int32
	_ = v2192
	var v2195 int32
	_ = v2195
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2212 int32
	_ = v2212
	var v2223 int32
	_ = v2223
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2229 int32
	_ = v2229
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2244 int32
	_ = v2244
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2257 int32
	_ = v2257
	var v2261 int32
	_ = v2261
	var v2266 int32
	_ = v2266
	var v2275 int32
	_ = v2275
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2290 int32
	_ = v2290
	var v2295 int32
	_ = v2295
	var v2299 int32
	_ = v2299
	var v2303 int32
	_ = v2303
	var v2308 int32
	_ = v2308
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2321 int32
	_ = v2321
	var v2326 int32
	_ = v2326
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2342 int32
	_ = v2342
	var v2347 int32
	_ = v2347
	var v2352 int32
	_ = v2352
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2369 int32
	_ = v2369
	var v2372 int32
	_ = v2372
	var v2381 int32
	_ = v2381
	var v2388 int32
	_ = v2388
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2401 int32
	_ = v2401
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2411 int32
	_ = v2411
	var v2416 int32
	_ = v2416
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2424 int32
	_ = v2424
	var v2435 int32
	_ = v2435
	var v2442 int32
	_ = v2442
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2455 int32
	_ = v2455
	var v2460 int32
	_ = v2460
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2475 int32
	_ = v2475
	var v2480 int32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
	var v2490 int32
	_ = v2490
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2495 int32
	_ = v2495
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2502 int64
	_ = v2502
	var v2505 int32
	_ = v2505
	var v2512 int32
	_ = v2512
	var v2515 int64
	_ = v2515
	var v2518 int64
	_ = v2518
	var v2519 int64
	_ = v2519
	var v2522 int64
	_ = v2522
	var v2523 int64
	_ = v2523
	var v2525 int64
	_ = v2525
	var v2526 int64
	_ = v2526
	var v2529 int64
	_ = v2529
	var v2535 int64
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2540 int32
	_ = v2540
	var v2543 int32
	_ = v2543
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2565 int32
	_ = v2565
	var v2570 int32
	_ = v2570
	var v2571 int32
	_ = v2571
	var v2575 int32
	_ = v2575
	var v2577 int32
	_ = v2577
	var v2579 int32
	_ = v2579
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2593 int32
	_ = v2593
	var v2596 int64
	_ = v2596
	var v2599 int64
	_ = v2599
	var v2600 int64
	_ = v2600
	var v2603 int64
	_ = v2603
	var v2604 int64
	_ = v2604
	var v2606 int64
	_ = v2606
	var v2607 int64
	_ = v2607
	var v2610 int64
	_ = v2610
	var v2620 int32
	_ = v2620
	var v2631 int32
	_ = v2631
	var v2638 int32
	_ = v2638
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2643 int32
	_ = v2643
	var v2651 int32
	_ = v2651
	var v2656 int32
	_ = v2656
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2670 int32
	_ = v2670
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2691 int32
	_ = v2691
	var v2694 int64
	_ = v2694
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2699 int64
	_ = v2699
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2710 int32
	_ = v2710
	var v2717 int32
	_ = v2717
	var v2720 int32
	_ = v2720
	var v2721 int32
	_ = v2721
	var v2722 int32
	_ = v2722
	var v2723 int32
	_ = v2723
	var v2729 int32
	_ = v2729
	var v2734 int32
	_ = v2734
	var v2735 int64
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2740 int32
	_ = v2740
	var v2743 int32
	_ = v2743
	var v2752 int32
	_ = v2752
	var v2758 int32
	_ = v2758
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2771 int32
	_ = v2771
	var v2776 int32
	_ = v2776
	var v2782 int32
	_ = v2782
	var v2783 int32
	_ = v2783
	var v2789 int32
	_ = v2789
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2799 int32
	_ = v2799
	var v2811 int32
	_ = v2811
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2820 int32
	_ = v2820
	var v2825 int32
	_ = v2825
	var v2828 int32
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2830 int32
	_ = v2830
	var v2833 int32
	_ = v2833
	var v2835 int32
	_ = v2835
	var v2841 int32
	_ = v2841
	var v2845 int32
	_ = v2845
	var v2846 int32
	_ = v2846
	var v2848 int32
	_ = v2848
	var v2856 int32
	_ = v2856
	var v2860 int32
	_ = v2860
	var v2870 int32
	_ = v2870
	var v2882 int32
	_ = v2882
	var v2883 int32
	_ = v2883
	var v2889 int32
	_ = v2889
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2901 int64
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2909 int32
	_ = v2909
	var v2910 int32
	_ = v2910
	var v2911 int32
	_ = v2911
	var v2915 int32
	_ = v2915
	var v2917 int32
	_ = v2917
	var v2918 int32
	_ = v2918
	var v2919 int32
	_ = v2919
	var v2923 int32
	_ = v2923
	var v2928 int32
	_ = v2928
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2933 int32
	_ = v2933
	var v2937 int32
	_ = v2937
	var v2940 int64
	_ = v2940
	var v2943 int32
	_ = v2943
	var v2944 int32
	_ = v2944
	var v2945 int64
	_ = v2945
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2956 int32
	_ = v2956
	var v2960 int64
	_ = v2960
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
	var v2965 int32
	_ = v2965
	var v2969 int32
	_ = v2969
	var v2972 int32
	_ = v2972
	var v2980 int32
	_ = v2980
	var v2981 int32
	_ = v2981
	var v2988 int32
	_ = v2988
	var v2989 int32
	_ = v2989
	var v2990 int32
	_ = v2990
	var v2994 int32
	_ = v2994
	var v2996 int32
	_ = v2996
	var v2998 int32
	_ = v2998
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3007 int32
	_ = v3007
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3012 int32
	_ = v3012
	var v3013 int32
	_ = v3013
	var v3019 int32
	_ = v3019
	var v3024 int32
	_ = v3024
	var v3025 int32
	_ = v3025
	var v3032 int32
	_ = v3032
	var v3035 int32
	_ = v3035
	var v3036 int32
	_ = v3036
	var v3037 int32
	_ = v3037
	var v3045 int32
	_ = v3045
	var v3050 int32
	_ = v3050
	var v3055 int32
	_ = v3055
	var v3056 int32
	_ = v3056
	var v3062 int32
	_ = v3062
	var v3067 int32
	_ = v3067
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3075 int32
	_ = v3075
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3082 int32
	_ = v3082
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3086 int32
	_ = v3086
	var v3089 int64
	_ = v3089
	var v3095 float64
	_ = v3095
	var v3096 int32
	_ = v3096
	var v3097 int32
	_ = v3097
	var v3100 int32
	_ = v3100
	var v3106 int32
	_ = v3106
	var v3109 int32
	_ = v3109
	var v3110 int32
	_ = v3110
	var v3111 int32
	_ = v3111
	var v3112 int32
	_ = v3112
	var v3121 int32
	_ = v3121
	var v3126 int32
	_ = v3126
	var v3127 float64
	_ = v3127
	var v3134 int32
	_ = v3134
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3142 int32
	_ = v3142
	var v3144 int32
	_ = v3144
	var v3147 int64
	_ = v3147
	var v3153 float64
	_ = v3153
	var v3154 int32
	_ = v3154
	var v3155 int32
	_ = v3155
	var v3158 int32
	_ = v3158
	var v3164 int32
	_ = v3164
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3170 int32
	_ = v3170
	var v3179 int32
	_ = v3179
	var v3184 int32
	_ = v3184
	var v3185 float64
	_ = v3185
	var v3192 int32
	_ = v3192
	var v3198 int32
	_ = v3198
	var v3204 int32
	_ = v3204
	var v3207 int32
	_ = v3207
	var v3208 int32
	_ = v3208
	var v3209 int32
	_ = v3209
	var v3210 int32
	_ = v3210
	var v3216 int32
	_ = v3216
	var v3221 int32
	_ = v3221
	var v3226 int32
	_ = v3226
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3235 int32
	_ = v3235
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3246 int32
	_ = v3246
	var v3249 int32
	_ = v3249
	var v3250 int32
	_ = v3250
	var v3251 int32
	_ = v3251
	var v3252 int32
	_ = v3252
	var v3258 int32
	_ = v3258
	var v3263 int32
	_ = v3263
	var v3267 int32
	_ = v3267
	var v3270 int32
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3279 int32
	_ = v3279
	var v3284 int32
	_ = v3284
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3289 int32
	_ = v3289
	var v3290 int32
	_ = v3290
	var v3292 int32
	_ = v3292
	var v3293 int32
	_ = v3293
	var v3294 int32
	_ = v3294
	var v3297 int32
	_ = v3297
	var v3298 int32
	_ = v3298
	var v3306 int32
	_ = v3306
	var v3307 int32
	_ = v3307
	var v3308 int32
	_ = v3308
	var v3309 int32
	_ = v3309
	var v3316 int32
	_ = v3316
	var v3319 int32
	_ = v3319
	var v3320 int32
	_ = v3320
	var v3321 int32
	_ = v3321
	var v3322 int32
	_ = v3322
	var v3328 int32
	_ = v3328
	var v3333 int32
	_ = v3333
	var v3335 int32
	_ = v3335
	var v3337 int32
	_ = v3337
	var v3338 int32
	_ = v3338
	var v3342 int32
	_ = v3342
	var v3343 int32
	_ = v3343
	var v3345 int32
	_ = v3345
	var v3347 int32
	_ = v3347
	var v3348 int32
	_ = v3348
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3354 int32
	_ = v3354
	var v3355 int32
	_ = v3355
	var v3356 int32
	_ = v3356
	var v3357 int32
	_ = v3357
	var v3359 int32
	_ = v3359
	var v3361 int32
	_ = v3361
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3365 int32
	_ = v3365
	var v3366 int32
	_ = v3366
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3378 int32
	_ = v3378
	var v3379 int32
	_ = v3379
	var v3381 int32
	_ = v3381
	var v3386 int32
	_ = v3386
	var v3387 int32
	_ = v3387
	var v3389 int32
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3391 int64
	_ = v3391
	var v3393 int32
	_ = v3393
	var v3394 int32
	_ = v3394
	var v3397 int32
	_ = v3397
	var v3398 int32
	_ = v3398
	var v3399 int32
	_ = v3399
	var v3404 int32
	_ = v3404
	var v3405 int32
	_ = v3405
	var v3408 int32
	_ = v3408
	var v3410 int32
	_ = v3410
	var v3411 int32
	_ = v3411
	var v3412 int32
	_ = v3412
	var v3413 int32
	_ = v3413
	var v3416 int32
	_ = v3416
	var v3420 int32
	_ = v3420
	var v3421 int32
	_ = v3421
	var v3423 int32
	_ = v3423
	var v3425 int32
	_ = v3425
	var v3429 int32
	_ = v3429
	var v3430 int32
	_ = v3430
	var v3431 int32
	_ = v3431
	var v3433 int32
	_ = v3433
	var v3434 int32
	_ = v3434
	var v3435 int32
	_ = v3435
	var v3439 int32
	_ = v3439
	var v3440 int32
	_ = v3440
	var v3441 int32
	_ = v3441
	var v3447 int32
	_ = v3447
	var v3448 int32
	_ = v3448
	var v3452 int32
	_ = v3452
	var v3453 int32
	_ = v3453
	var v3454 int32
	_ = v3454
	var v3456 int32
	_ = v3456
	var v3463 int32
	_ = v3463
	var v3466 int32
	_ = v3466
	var v3470 int32
	_ = v3470
	var v3475 int32
	_ = v3475
	var v3479 int32
	_ = v3479
	var v3482 int32
	_ = v3482
	var v3483 int32
	_ = v3483
	var v3484 int32
	_ = v3484
	var v3485 int32
	_ = v3485
	var v3486 int32
	_ = v3486
	var v3492 int32
	_ = v3492
	var v3497 int32
	_ = v3497
	var v3500 int32
	_ = v3500
	var v3501 int32
	_ = v3501
	var v3503 int32
	_ = v3503
	var v3504 int32
	_ = v3504
	var v3505 int32
	_ = v3505
	var v3510 int32
	_ = v3510
	var v3511 int32
	_ = v3511
	var v3514 int32
	_ = v3514
	var v3516 int32
	_ = v3516
	var v3522 int32
	_ = v3522
	var v3523 int32
	_ = v3523
	var v3524 int32
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3526 int32
	_ = v3526
	var v3528 int32
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3530 int32
	_ = v3530
	var v3531 int32
	_ = v3531
	var v3532 int32
	_ = v3532
	var v3539 int32
	_ = v3539
	var v3542 int32
	_ = v3542
	var v3546 int32
	_ = v3546
	var v3551 int32
	_ = v3551
	var v3556 int32
	_ = v3556
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
	v31 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
	return v3556
L8:
	;
	v3531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3531 != 0 {
		v3556 = v329
		goto L7
	} else {
		goto L1109
	}
L9:
	;
	v3500 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v3501 = m.ExcPending
	if v3501 != 0 {
		goto L1
	} else {
		goto L1098
	}
L10:
	;
	v3412 = F_JsonbType(m, l2)
	mBase = m.M
	v3413 = m.ExcPending
	if v3413 != 0 {
		goto L1
	} else {
		goto L1058
	}
L11:
	;
	v3408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3410 = F_executeNextItem(m, l0, l1, int32(0), v3408, l3, int32(1))
	mBase = m.M
	v3411 = m.ExcPending
	if v3411 != 0 {
		goto L1
	} else {
		goto L1056
	}
L12:
	;
	v3391 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v3393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3394 = *(*int32)(unsafe.Add(mBase, uint32(v3393)))
	if v3394 == int32(18) {
		goto L1052
	} else {
		goto L1053
	}
L13:
	;
	if l4 != 0 {
		goto L1042
	} else {
		goto L1043
	}
L14:
	;
	v3350 = F_palloc(m, int32(20))
	mBase = m.M
	v3351 = m.ExcPending
	if v3351 != 0 {
		goto L1
	} else {
		goto L1037
	}
L15:
	;
	v3294 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v3294 != int32(18) {
		goto L1021
	} else {
		goto L1022
	}
L16:
	;
	v3292 = F_executeNumericItemMethod(m, l0, l1, l2, l4, int32(1292), l3)
	mBase = m.M
	v3293 = m.ExcPending
	if v3293 != 0 {
		goto L1
	} else {
		goto L1019
	}
L17:
	;
	v3289 = F_executeNumericItemMethod(m, l0, l1, l2, l4, int32(1438), l3)
	mBase = m.M
	v3290 = m.ExcPending
	if v3290 != 0 {
		goto L1
	} else {
		goto L1018
	}
L18:
	;
	v3286 = F_executeNumericItemMethod(m, l0, l1, l2, l4, int32(1437), l3)
	mBase = m.M
	v3287 = m.ExcPending
	if v3287 != 0 {
		goto L1
	} else {
		goto L1017
	}
L19:
	;
	if l4 == int32(0) {
		goto L957
	} else {
		goto L958
	}
L20:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l4 == int32(0) {
		goto L651
	} else {
		goto L652
	}
L21:
	;
	if l4 == int32(0) {
		goto L565
	} else {
		goto L566
	}
L22:
	;
	v1611 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L1
	} else {
		goto L551
	}
L23:
	;
	if l4 == int32(0) {
		goto L508
	} else {
		goto L509
	}
L24:
	;
	if l4 == int32(0) {
		goto L461
	} else {
		goto L462
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
		goto L295
	} else {
		goto L296
	}
L27:
	;
	if l4 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L265
	}
L29:
	;
	v675 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1423), l3)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L264
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L259
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L256
	}
L32:
	;
	v325 = F_JsonbType(m, l2)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L160
	}
L33:
	;
	v271 = F_JsonbType(m, l2)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L135
	}
L34:
	;
	v227 = F_JsonbType(m, l2)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L114
	}
L35:
	;
	v225 = F_executeUnaryArithmExpr(m, l0, l1, l2, int32(1422), l3)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L113
	}
L36:
	;
	v222 = F_executeUnaryArithmExpr(m, l0, l1, l2, int32(0), l3)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L112
	}
L37:
	;
	v219 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1421), l3)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L111
	}
L38:
	;
	v216 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1420), l3)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L110
	}
L39:
	;
	v213 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1419), l3)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L109
	}
L40:
	;
	v210 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1418), l3)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L108
	}
L41:
	;
	v181 = F_executeBoolItem(m, l0, l1, l2, int32(1))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
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
	if l3 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if v37 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	if v37 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v40 != int32(28) {
		v3556 = int32(0)
		goto L7
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v49 = F_palloc(m, int32(20))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	v51 = v24 + int32(908)
	goto L50
L50:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v53 {
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
	v51 = v49
	goto L50
L52:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v126 {
		goto L77
	} else {
		goto L78
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(0)
	goto L52
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L74
	}
L55:
	;
	v74 = v24 + int32(936)
	if v74 != 0 {
		goto L64
	} else {
		goto L65
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(1)
	v68 = v51 + int32(4)
	if v68 != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(2)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v63
	goto L52
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(3)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)) = uint8(base.B2i32(v57 != int32(0)))
	goto L52
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v71
	goto L52
L60:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v69
	goto L62
L61:
	;
	goto L62
L62:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L59
L63:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v78 == int32(0) {
		goto L30
	} else {
		goto L67
	}
L64:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v75
	goto L66
L65:
	;
	goto L66
L66:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L63
L67:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v24)+936))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v87 = m.T0[v86].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v78, v77, v81, v24+int32(952), v24+int32(888))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	if v87 == int32(0) {
		goto L30
	} else {
		goto L69
	}
L69:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v24)+888))
	if v91 <= int32(0) {
		goto L52
	} else {
		goto L70
	}
L70:
	;
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v96
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v24)+960))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v24)+952))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v91
	if v101 == int32(18) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v106 = v100
	goto L73
L72:
	;
	v106 = int32(0)
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v106
	goto L52
L74:
	;
	F_errmsg_internal(m, int32(369188), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(500183), int32(2982), int32(291427))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
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
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v132 = F_executeItemOptUnwrapTarget(m, l0, v24+int32(752), v51, l3, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v52
	v3556 = v132
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
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v52
	v3556 = int32(0)
	goto L7
L84:
	;
	v148 = m.G0
	v150 = v148 - int32(16)
	m.G0 = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v152 != 0 {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v147 = v51
	goto L84
L86:
	;
	goto L87
L87:
	;
	v139 = F_palloc(m, int32(20))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+16)) = v141
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v143
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	*(*int64)(unsafe.Add(mBase, uint32(v139))) = v145
	v147 = v139
	goto L84
L89:
	;
	m.G0 = v150 + int32(16)
	goto L83
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+8)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v150)+12)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v147
	v159 = F_list_make2_impl(m, v150+int32(4), v150)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v164 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v159
	goto L89
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v147
	goto L89
L95:
	;
	goto L96
L96:
	;
	v168 = F_lappend(m, v164, v147)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v168
	goto L89
L98:
	;
	v185 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
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
	v189 = int32(0)
	if v185 == v189 {
		v3556 = v189
		goto L7
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v181 != int32(2) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L102
L104:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)) = uint8(base.B2i32(v181 == int32(1)))
	v200 = int32(3)
	goto L106
L105:
	;
	v200 = int32(0)
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+952)) = v200
	v207 = F_executeNextItem(m, l0, l1, v24+int32(752), v24+int32(952), l3, int32(1))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v3556 = v207
	goto L7
L108:
	;
	v3556 = v210
	goto L7
L109:
	;
	v3556 = v213
	goto L7
L110:
	;
	v3556 = v216
	goto L7
L111:
	;
	v3556 = v219
	goto L7
L112:
	;
	v3556 = v222
	goto L7
L113:
	;
	v3556 = v225
	goto L7
L114:
	;
	if v227 == int32(16) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v232 = v24 + int32(752)
	v236 = F_jspGetNext(m, l1, v232)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v242 = int32(1)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v243 == v242 {
		goto L123
	} else {
		goto L124
	}
L118:
	;
	if v236 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v238 = v232
	goto L121
L120:
	;
	v238 = int32(0)
	goto L121
L121:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v240 = F_executeItemUnwrapTargetArray(m, l0, v238, l2, l3, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v3556 = v240
	goto L7
L123:
	;
	v248 = F_executeNextItem(m, l0, l1, int32(0), l2, l3, int32(1))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v250 != 0 {
		v3556 = v242
		goto L7
	} else {
		goto L127
	}
L126:
	;
	v3556 = v248
	goto L7
L127:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v251 != int32(1) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v3556 = int32(2)
	goto L7
L129:
	;
	goto L130
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errcode(m, int32(151781506))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(25361), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(500183), int32(849), int32(108435))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
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
	if v271 == int32(17) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v277 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
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
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v279 != int32(18) {
		goto L31
	} else {
		goto L140
	}
L140:
	;
	if v277 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v285 = v24 + int32(752)
	goto L143
L142:
	;
	v285 = int32(0)
	goto L143
L143:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v287 = int32(1)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v292 = F_executeAnyItem(m, l0, v285, v286, l3, v287, v287, v287, int32(0), v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v3556 = v292
	goto L7
L145:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v303 != 0 {
		goto L150
	} else {
		goto L151
	}
L146:
	;
	v296 = F_JsonbType(m, l2)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	if v296 != int32(16) {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v301 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v3556 = v301
	goto L7
L150:
	;
	v3556 = int32(1)
	goto L7
L151:
	;
	goto L152
L152:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v305 != int32(1) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v3556 = int32(2)
	goto L7
L154:
	;
	goto L155
L155:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errcode(m, int32(319553666))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(110944), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(500183), int32(872), int32(108435))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
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
	if v325 != int32(16) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v329 = int32(1)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v330 != v329 {
		goto L8
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v334 != int32(18) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L163
L165:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v353 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L169
	}
L166:
	;
	v348 = int32(-1)
	goto L165
L167:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	if v338&int32(1342177280) != int32(1073741824) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v348 = v338 & int32(268435455)
	goto L165
L169:
	;
	v357 = base.B2i32(v348 < int32(0))
	if v348 < int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v358 = int32(1)
	goto L172
L171:
	;
	v358 = v348
	goto L172
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v358
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if int32(0) < v360 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v374 = v6
	goto L176
L174:
	;
	v618 = int32(1)
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v349
	v3556 = v618
	goto L7
L176:
	;
	v389 = int32(2)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v396 = v374 << (uint(int32(3)) % 32)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v396+v397)))
	F_jspInitByBuffer(m, v24+int32(952), v394, v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L178
	}
L177:
	;
	v618 = v593
	goto L175
L178:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v402+v396)+4))
	if v404 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v24+int32(908), v405, v404)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v414 = F_getArrayIndex(m, l0, v24+int32(952), l2, v24+int32(900))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L185
	}
L182:
	;
	goto L181
L183:
	;
	v611 = v374 + int32(1)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v611 < v612 {
		v374 = v611
		goto L176
	} else {
		goto L255
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v349
	v3556 = v577
	goto L7
L185:
	;
	if v414 == int32(2) {
		v577 = v389
		goto L184
	} else {
		goto L186
	}
L186:
	;
	if v404 != int32(0) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v430 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L188:
	;
	v422 = F_getArrayIndex(m, l0, v24+int32(908), l2, v24+int32(988))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v427
	v429 = v427
	goto L187
L191:
	;
	if v422 == int32(2) {
		v577 = v389
		goto L184
	} else {
		goto L192
	}
L192:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	v429 = v426
	goto L187
L193:
	;
	if v460 < v358 {
		goto L208
	} else {
		goto L209
	}
L194:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	v460 = v433
	goto L193
L195:
	;
	goto L196
L196:
	;
	if v429 < int32(0) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v440 != int32(1) {
		goto L201
	} else {
		goto L202
	}
L198:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	if v436 < v429 {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	if v436 < v358 {
		v460 = v436
		goto L193
	} else {
		goto L200
	}
L200:
	;
	goto L197
L201:
	;
	v3556 = int32(2)
	goto L7
L202:
	;
	goto L203
L203:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	F_errcode(m, int32(51118210))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	F_errmsg(m, int32(172414), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(500183), int32(921), int32(108435))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
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
	v462 = v460
	goto L210
L209:
	;
	v462 = v358 - int32(1)
	goto L210
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v462
	v464 = int32(0)
	if v464 < v429 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v467 = v429
	goto L213
L212:
	;
	v467 = v464
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = v467
	v469 = int32(1)
	if v462 < v467 {
		v593 = v469
		goto L183
	} else {
		goto L214
	}
L214:
	;
	v475 = v469
	v479 = v467
	v480 = v462
	goto L216
L215:
	;
	if l3 != 0 {
		v593 = v562
		goto L183
	} else {
		goto L253
	}
L216:
	;
	if v357 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	if v554 == int32(2) {
		v577 = v389
		goto L184
	} else {
		goto L252
	}
L218:
	;
	v558 = v479 + int32(1)
	if v558 <= v556 {
		v475 = v554
		v479 = v558
		v480 = v556
		goto L216
	} else {
		goto L251
	}
L219:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v495 = F_getIthJsonbValueFromContainer(m, v494, v479)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	v499 = l2
	goto L221
L221:
	;
	if base.B2i32(l3 != int32(0))|v353 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L222:
	;
	if v495 == int32(0) {
		v554 = v475
		v556 = v480
		goto L218
	} else {
		goto L223
	}
L223:
	;
	v499 = v495
	goto L221
L224:
	;
	v3556 = int32(0)
	goto L7
L225:
	;
	goto L226
L226:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v503 <= int32(0) {
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
	v546 = F_lappend(m, v535, v520)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L250
	}
L229:
	;
	if l3 != 0 {
		v554 = v544
		v556 = v480
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
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v540 = F_executeItemOptUnwrapTarget(m, l0, v24+int32(752), v499, l3, v539)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L246
	}
L233:
	;
	v544 = int32(0)
	goto L229
L234:
	;
	goto L235
L235:
	;
	if int32(0) <= v348 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v521 != 0 {
		goto L241
	} else {
		goto L242
	}
L237:
	;
	v520 = v499
	goto L236
L238:
	;
	goto L239
L239:
	;
	v512 = F_palloc(m, int32(20))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v499)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v512)+16)) = v514
	v516 = *(*int64)(unsafe.Add(mBase, uint32(v499)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v512)+8)) = v516
	v518 = *(*int64)(unsafe.Add(mBase, uint32(v499)))
	*(*int64)(unsafe.Add(mBase, uint32(v512))) = v518
	v520 = v512
	goto L236
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+888)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v24)+936)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v520
	v530 = F_list_make2_impl(m, v24+int32(44), v24+int32(40))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v535 != 0 {
		goto L228
	} else {
		goto L245
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v530
	v552 = v480
	goto L227
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v520
	v552 = v480
	goto L227
L246:
	;
	if v540 == int32(2) {
		v577 = v389
		goto L184
	} else {
		goto L247
	}
L247:
	;
	v544 = v540
	goto L229
L248:
	;
	if v544 != 0 {
		v554 = v544
		v556 = v480
		goto L218
	} else {
		goto L249
	}
L249:
	;
	v562 = int32(0)
	goto L215
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v546
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	v552 = v549
	goto L227
L251:
	;
	goto L217
L252:
	;
	v562 = v554
	goto L215
L253:
	;
	if v562 != 0 {
		v593 = v562
		goto L183
	} else {
		goto L254
	}
L254:
	;
	v577 = int32(0)
	goto L184
L255:
	;
	goto L177
L256:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v640
	F_errmsg_internal(m, int32(485026), v24+int32(32))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(500183), int32(858), int32(108435))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v24)+936))
	v661 = F_pnstrdup(m, v77, v660)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v661
	F_errmsg(m, int32(721501), v24+int32(16))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(500183), int32(3158), int32(396987))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	v3556 = v675
	goto L7
L265:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v681
	F_errmsg_internal(m, int32(485419), v24)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(500183), int32(1663), int32(108435))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	v700 = F_JsonbType(m, l2)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L280
	}
L269:
	;
	v693 = F_JsonbType(m, l2)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	if v693 != int32(16) {
		goto L268
	} else {
		goto L271
	}
L271:
	;
	v698 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	v3556 = v698
	goto L7
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v755
	v757 = F_strlen(m, v755)
	mBase = m.M
	v758 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+952)) = v758
	*(*int32)(unsafe.Add(mBase, uint32(v24)+956)) = v757
	v765 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(952), l3, v758)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L294
	}
L274:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v753 = F_pnstrdup(m, v751, v752)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L293
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v726 != int32(1) {
		v3556 = int32(2)
		goto L7
	} else {
		goto L287
	}
L277:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v719 = F_JsonEncodeDateTime(m, v24+int32(752), v715, v716, l2+int32(16))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L285
	}
L278:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v711 != 0 {
		goto L282
	} else {
		goto L283
	}
L279:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v707 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v706)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L281
	}
L280:
	;
	switch v700 - int32(1) {
	case 0:
		goto L274
	case 1:
		goto L279
	case 2:
		goto L278
	case 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30:
		goto L275
	default:
		goto L276
	case 31:
		goto L277
	}
L281:
	;
	v755 = v707
	goto L273
L282:
	;
	v712 = int32(344470)
	goto L284
L283:
	;
	v712 = int32(361486)
	goto L284
L284:
	;
	v755 = v712
	goto L273
L285:
	;
	v723 = F_pstrdup(m, v24+int32(752))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	v755 = v723
	goto L273
L287:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v737 = F_jspOperationName(m, v736)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+736)) = v737
	F_errmsg(m, int32(346248), v24+int32(736))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(500183), int32(1648), int32(108435))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	goto L275
L293:
	;
	v755 = v753
	goto L273
L294:
	;
	v3556 = v765
	goto L7
L295:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v776 - int32(1) {
	case 0:
		goto L304
	case 1:
		goto L305
	default:
		goto L303
	}
L296:
	;
	v769 = F_JsonbType(m, l2)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	if v769 != int32(16) {
		goto L295
	} else {
		goto L298
	}
L298:
	;
	v774 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	v3556 = v774
	goto L7
L300:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L1
	} else {
		goto L332
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = int32(2)
	v878 = F_DirectFunctionCall1Coll(m, int32(1414), int32(0), v873)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L329
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+908)) = v782
	v873 = v782
	goto L301
L303:
	;
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v847 != int32(1) {
		v3556 = int32(2)
		goto L7
	} else {
		goto L323
	}
L304:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v822 = F_pnstrdup(m, v820, v821)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L317
	}
L305:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v782 = F_numeric_int4_opt_error(m, v779, v24+int32(952))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+952)))
	if v784 != int32(1) {
		goto L302
	} else {
		goto L307
	}
L307:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v787 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v3556 = int32(2)
	goto L7
L309:
	;
	goto L310
L310:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v801 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v800)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v804 = F_jspOperationName(m, v803)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+712)) = int32(225160)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+708)) = v804
	*(*int32)(unsafe.Add(mBase, uint32(v24)+704)) = v801
	F_errmsg(m, int32(189987), v24+int32(704))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	F_errfinish(m, int32(500183), int32(1563), int32(108435))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L317:
	;
	v825 = *(*int32)(unsafe.Add(mBase, _consts[1074]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v825
	v828 = *(*int64)(unsafe.Add(mBase, _consts[1075]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v828
	v836 = F_DirectInputFunctionCallSafe(m, int32(1424), v822, int32(-1), v24+int32(952), v24+int32(908))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L319
	}
L318:
	;
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v842 == int32(1) {
		goto L300
	} else {
		goto L322
	}
L319:
	;
	if v836 == int32(0) {
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v840 != 0 {
		goto L318
	} else {
		goto L321
	}
L321:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v24)+908))
	v873 = v841
	goto L301
L322:
	;
	v3556 = int32(2)
	goto L7
L323:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v858 = F_jspOperationName(m, v857)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+688)) = v858
	F_errmsg(m, int32(346373), v24+int32(688))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	F_errfinish(m, int32(500183), int32(1593), int32(108435))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L329:
	;
	v880 = F_pg_detoast_datum(m, v878)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+756)) = v880
	v887 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(752), l3, int32(1))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	v3556 = v887
	goto L7
L332:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v897 = F_jspOperationName(m, v896)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+728)) = int32(225160)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+724)) = v897
	*(*int32)(unsafe.Add(mBase, uint32(v24)+720)) = v822
	F_errmsg(m, int32(189987), v24+int32(720))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(500183), int32(1585), int32(108435))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	v3556 = v920
	goto L7
L342:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L1
	} else {
		goto L456
	}
L343:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L1
	} else {
		goto L453
	}
L344:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L1
	} else {
		goto L450
	}
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+956)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v24)+952)) = int32(2)
	v1268 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(952), l3, int32(1))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L1
	} else {
		goto L449
	}
L346:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v1086 != int32(46) {
		v1258 = v1084
		goto L345
	} else {
		goto L403
	}
L347:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L398
	}
L348:
	;
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1038 != int32(1) {
		goto L390
	} else {
		goto L391
	}
L349:
	;
	v971 = *(*int32)(unsafe.Add(mBase, _consts[1074]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = v971
	v974 = *(*int64)(unsafe.Add(mBase, _consts[1075]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v974
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v979 = F_pnstrdup(m, v977, v978)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L370
	}
L350:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v926 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v925)+4)))
	goto L352
L351:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v963 != int32(46) {
		v1258 = v925
		goto L345
	} else {
		goto L366
	}
L352:
	;
	if base.B2i32(v926 == int32(49152)) == int32(0) {
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
		goto L358
	} else {
		goto L359
	}
L356:
	;
	if base.B2i32(v931&int32(57343) == int32(53248)) == int32(0) {
		goto L351
	} else {
		goto L357
	}
L357:
	;
	goto L355
L358:
	;
	v3556 = int32(2)
	goto L7
L359:
	;
	goto L360
L360:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v950 = F_jspOperationName(m, v949)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+592)) = v950
	F_errmsg(m, int32(683612), v24+int32(592))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	F_errfinish(m, int32(500183), int32(1414), int32(108435))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L366:
	;
	v968 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v925)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L367
	}
L367:
	;
	v1084 = v925
	v1085 = v968
	goto L346
L368:
	;
	v3556 = int32(2)
	goto L7
L369:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v24)+908))
	v1019 = F_pg_detoast_datum(m, v1018)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L1
	} else {
		goto L382
	}
L370:
	;
	v986 = F_DirectInputFunctionCallSafe(m, int32(408), v979, int32(-1), v24+int32(752), v24+int32(908))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	if v986 != 0 {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)))
	if v988 != int32(1) {
		goto L369
	} else {
		goto L375
	}
L373:
	;
	goto L374
L374:
	;
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v991 != int32(1) {
		goto L368
	} else {
		goto L376
	}
L375:
	;
	goto L374
L376:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L377
	}
L377:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L378
	}
L378:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1002 = F_jspOperationName(m, v1001)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+664)) = int32(491022)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+660)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v24)+656)) = v979
	F_errmsg(m, int32(189987), v24+int32(656))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	F_errfinish(m, int32(500183), int32(1439), int32(108435))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
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
	v1021 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1019)+4)))
	goto L383
L383:
	;
	if base.B2i32(v1021 == int32(49152)) == int32(0) {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1026 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1019)+4)))
	goto L387
L385:
	;
	goto L386
L386:
	;
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1033 == int32(1) {
		goto L347
	} else {
		goto L389
	}
L387:
	;
	if base.B2i32(v1026&int32(57343) == int32(53248)) == int32(0) {
		v1084 = v1019
		v1085 = v979
		goto L346
	} else {
		goto L388
	}
L388:
	;
	goto L386
L389:
	;
	goto L368
L390:
	;
	v3556 = int32(2)
	goto L7
L391:
	;
	goto L392
L392:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1050 = F_jspOperationName(m, v1049)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+576)) = v1050
	F_errmsg(m, int32(346373), v24+int32(576))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	F_errfinish(m, int32(500183), int32(1455), int32(108435))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
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
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1071 = F_jspOperationName(m, v1070)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+672)) = v1071
	F_errmsg(m, int32(683612), v24+int32(672))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	F_errfinish(m, int32(500183), int32(1446), int32(108435))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L403:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1089 == int32(0) {
		v1258 = v1084
		goto L345
	} else {
		goto L404
	}
L404:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, _consts[1074]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+896)) = v1093
	v1096 = *(*int64)(unsafe.Add(mBase, _consts[1075]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+888)) = v1096
	F_jspGetArg(m, l1, v24+int32(752))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v24)+752))
	if v1102 != int32(2) {
		goto L344
	} else {
		goto L406
	}
L406:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(752))+12))
	v1110 = F_numeric_int4_opt_error(m, v1107, v24+int32(984))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L407
	}
L407:
	;
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+984)))
	if v1112 == int32(1) {
		goto L410
	} else {
		goto L411
	}
L408:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	v1254 = F_pg_detoast_datum(m, v1253)
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L1
	} else {
		goto L447
	}
L409:
	;
	v3556 = int32(2)
	goto L7
L410:
	;
	v1115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1115 != int32(1) {
		goto L409
	} else {
		goto L413
	}
L411:
	;
	goto L412
L412:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1139 == int32(0) {
		v1183 = v6
		goto L419
	} else {
		goto L420
	}
L413:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1126 = F_jspOperationName(m, v1125)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+608)) = v1126
	F_errmsg(m, int32(224955), v24+int32(608))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L1
	} else {
		goto L417
	}
L417:
	;
	F_errfinish(m, int32(500183), int32(1487), int32(108435))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L1
	} else {
		goto L418
	}
L418:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L419:
	;
	v1185 = v24 + int32(908)
	if int32(0) <= v1110 {
		goto L432
	} else {
		goto L433
	}
L420:
	;
	F_jspGetRightArg(m, l1, v24+int32(752))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v24)+752))
	if v1146 != int32(2) {
		goto L343
	} else {
		goto L422
	}
L422:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(752))+12))
	v1154 = F_numeric_int4_opt_error(m, v1151, v24+int32(984))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L1
	} else {
		goto L423
	}
L423:
	;
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+984)))
	if v1156 != int32(1) {
		v1183 = v1154
		goto L419
	} else {
		goto L424
	}
L424:
	;
	v1159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1159 != int32(1) {
		goto L409
	} else {
		goto L425
	}
L425:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L427
	}
L427:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1170 = F_jspOperationName(m, v1169)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+640)) = v1170
	F_errmsg(m, int32(225028), v24+int32(640))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L1
	} else {
		goto L429
	}
L429:
	;
	F_errfinish(m, int32(500183), int32(1501), int32(108435))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = v24 + int32(908)
	v1207 = v24 + int32(936)
	if int32(0) <= v1183 {
		goto L436
	} else {
		goto L437
	}
L432:
	;
	v1195 = v1110
	v1196 = int32(0)
	goto L434
L433:
	;
	v1190 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1185))) = uint8(v1190)
	v1195 = int32(0) - v1110
	v1196 = int32(1)
	goto L434
L434:
	;
	v1198 = F_pg_ultoa_n(m, v1195, v1185+v1196)
	mBase = m.M
	v1201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1185+(v1198+v1196)))) = uint8(v1201)
	goto L431
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+904)) = v24 + int32(936)
	v1235 = F_construct_array_builtin(m, v24+int32(900), int32(2), int32(2275))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L1
	} else {
		goto L439
	}
L436:
	;
	v1217 = v1183
	v1218 = int32(0)
	goto L438
L437:
	;
	v1212 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1207))) = uint8(v1212)
	v1217 = int32(0) - v1183
	v1218 = int32(1)
	goto L438
L438:
	;
	v1220 = F_pg_ultoa_n(m, v1217, v1207+v1218)
	mBase = m.M
	v1223 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1207+(v1220+v1218)))) = uint8(v1223)
	goto L435
L439:
	;
	v1237 = F_DirectFunctionCall1Coll(m, int32(1425), int32(0), v1235)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L1
	} else {
		goto L440
	}
L440:
	;
	v1243 = F_DirectInputFunctionCallSafe(m, int32(408), v1085, v1237, v24+int32(888), v24+int32(988))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L1
	} else {
		goto L441
	}
L441:
	;
	if v1243 != 0 {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v1245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+892)))
	if v1245 != int32(1) {
		goto L408
	} else {
		goto L445
	}
L443:
	;
	goto L444
L444:
	;
	v1248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1248 == int32(1) {
		goto L342
	} else {
		goto L446
	}
L445:
	;
	goto L444
L446:
	;
	goto L409
L447:
	;
	F_pfree(m, v1235)
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	v1258 = v1254
	goto L345
L449:
	;
	v3556 = v1268
	goto L7
L450:
	;
	F_errmsg_internal(m, int32(271947), int32(0))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	F_errfinish(m, int32(500183), int32(1479), int32(108435))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L453:
	;
	F_errmsg_internal(m, int32(397674), int32(0))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L1
	} else {
		goto L454
	}
L454:
	;
	F_errfinish(m, int32(500183), int32(1493), int32(108435))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L1
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
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1304 = F_jspOperationName(m, v1303)
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+632)) = int32(491022)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+628)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v24)+624)) = v1085
	F_errmsg(m, int32(189987), v24+int32(624))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L1
	} else {
		goto L459
	}
L459:
	;
	F_errfinish(m, int32(500183), int32(1528), int32(108435))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L1
	} else {
		goto L460
	}
L460:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L461:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v1329 - int32(1) {
	case 0:
		goto L471
	case 1:
		goto L472
	case 2:
		goto L468
	default:
		goto L470
	}
L462:
	;
	v1322 = F_JsonbType(m, l2)
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L1
	} else {
		goto L463
	}
L463:
	;
	if v1322 != int32(16) {
		goto L461
	} else {
		goto L464
	}
L464:
	;
	v1327 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	v3556 = v1327
	goto L7
L466:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)) = uint8(v1451)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = int32(3)
	v1459 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(752), l3, int32(1))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L1
	} else {
		goto L507
	}
L467:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+936)) = uint8(v1449)
	v1451 = v1449
	goto L466
L468:
	;
	v1447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v1449 = v1447
	goto L467
L469:
	;
	v3556 = int32(2)
	goto L7
L470:
	;
	v1421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1421 != int32(1) {
		goto L499
	} else {
		goto L500
	}
L471:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1386 = F_pnstrdup(m, v1384, v1385)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L1
	} else {
		goto L486
	}
L472:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1335 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v1334)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L1
	} else {
		goto L473
	}
L473:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, _consts[1074]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v1338
	v1341 = *(*int64)(unsafe.Add(mBase, _consts[1075]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v1341
	v1349 = F_DirectInputFunctionCallSafe(m, int32(1424), v1335, int32(-1), v24+int32(952), v24+int32(908))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L1
	} else {
		goto L475
	}
L474:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v24)+908))
	v1449 = base.B2i32(v1381 != int32(0))
	goto L467
L475:
	;
	if v1349 != 0 {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v1351 != int32(1) {
		goto L474
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	v1354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1354 != int32(1) {
		goto L469
	} else {
		goto L480
	}
L479:
	;
	goto L478
L480:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L1
	} else {
		goto L481
	}
L481:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1365 = F_jspOperationName(m, v1364)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L1
	} else {
		goto L483
	}
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+552)) = int32(284269)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+548)) = v1365
	*(*int32)(unsafe.Add(mBase, uint32(v24)+544)) = v1335
	F_errmsg(m, int32(189987), v24+int32(544))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L1
	} else {
		goto L484
	}
L484:
	;
	F_errfinish(m, int32(500183), int32(1357), int32(108435))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L1
	} else {
		goto L485
	}
L485:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L486:
	;
	v1390 = F_strlen(m, v1386)
	mBase = m.M
	v1391 = F_parse_bool_with_len(m, v1386, v1390, v24+int32(936))
	mBase = m.M
	goto L487
L487:
	;
	if v1391 != 0 {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v1392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+936)))
	v1451 = v1392
	goto L466
L489:
	;
	goto L490
L490:
	;
	v1393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1393 != int32(1) {
		goto L491
	} else {
		goto L492
	}
L491:
	;
	v3556 = int32(2)
	goto L7
L492:
	;
	goto L493
L493:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L1
	} else {
		goto L495
	}
L495:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1405 = F_jspOperationName(m, v1404)
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+568)) = int32(284269)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+564)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v24)+560)) = v1386
	F_errmsg(m, int32(189987), v24+int32(560))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	F_errfinish(m, int32(500183), int32(1377), int32(108435))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L499:
	;
	v3556 = int32(2)
	goto L7
L500:
	;
	goto L501
L501:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L1
	} else {
		goto L502
	}
L502:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L503
	}
L503:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1433 = F_jspOperationName(m, v1432)
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L1
	} else {
		goto L504
	}
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+528)) = v1433
	F_errmsg(m, int32(346449), v24+int32(528))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	F_errfinish(m, int32(500183), int32(1386), int32(108435))
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L507:
	;
	v3556 = v1459
	goto L7
L508:
	;
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v1470 - int32(1) {
	case 0:
		goto L517
	case 1:
		goto L518
	default:
		goto L516
	}
L509:
	;
	v1463 = F_JsonbType(m, l2)
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L1
	} else {
		goto L510
	}
L510:
	;
	if v1463 != int32(16) {
		goto L508
	} else {
		goto L511
	}
L511:
	;
	v1468 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v3556 = v1468
	goto L7
L513:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L1
	} else {
		goto L546
	}
L514:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = int32(2)
	v1574 = F_DirectFunctionCall1Coll(m, int32(1415), int32(0), v1568)
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L1
	} else {
		goto L543
	}
L515:
	;
	v1565 = F_Int64GetDatum(m, v1476)
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L1
	} else {
		goto L542
	}
L516:
	;
	v1541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1541 != int32(1) {
		v3556 = int32(2)
		goto L7
	} else {
		goto L536
	}
L517:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1516 = F_pnstrdup(m, v1514, v1515)
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L1
	} else {
		goto L530
	}
L518:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1476 = F_numeric_int8_opt_error(m, v1473, v24+int32(952))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	v1478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+952)))
	if v1478 != int32(1) {
		goto L515
	} else {
		goto L520
	}
L520:
	;
	v1481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1481 == int32(0) {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	v3556 = int32(2)
	goto L7
L522:
	;
	goto L523
L523:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1495 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v1494)
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L1
	} else {
		goto L526
	}
L526:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1498 = F_jspOperationName(m, v1497)
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L1
	} else {
		goto L527
	}
L527:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+504)) = int32(89347)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+500)) = v1498
	*(*int32)(unsafe.Add(mBase, uint32(v24)+496)) = v1495
	F_errmsg(m, int32(189987), v24+int32(496))
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L1
	} else {
		goto L528
	}
L528:
	;
	F_errfinish(m, int32(500183), int32(1283), int32(108435))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L530:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, _consts[1074]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v1519
	v1522 = *(*int64)(unsafe.Add(mBase, _consts[1075]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v1522
	v1530 = F_DirectInputFunctionCallSafe(m, int32(546), v1516, int32(-1), v24+int32(952), v24+int32(908))
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L1
	} else {
		goto L532
	}
L531:
	;
	v1536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1536 == int32(1) {
		goto L513
	} else {
		goto L535
	}
L532:
	;
	if v1530 == int32(0) {
		goto L531
	} else {
		goto L533
	}
L533:
	;
	v1534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v1534 != 0 {
		goto L531
	} else {
		goto L534
	}
L534:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v24)+908))
	v1568 = v1535
	goto L514
L535:
	;
	v3556 = int32(2)
	goto L7
L536:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L1
	} else {
		goto L537
	}
L537:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L1
	} else {
		goto L538
	}
L538:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1552 = F_jspOperationName(m, v1551)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L1
	} else {
		goto L539
	}
L539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+480)) = v1552
	F_errmsg(m, int32(346373), v24+int32(480))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	F_errfinish(m, int32(500183), int32(1313), int32(108435))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L1
	} else {
		goto L541
	}
L541:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L542:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+908)) = v1565
	v1568 = v1565
	goto L514
L543:
	;
	v1576 = F_pg_detoast_datum(m, v1574)
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L1
	} else {
		goto L544
	}
L544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+756)) = v1576
	v1583 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(752), l3, int32(1))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L1
	} else {
		goto L545
	}
L545:
	;
	v3556 = v1583
	goto L7
L546:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L1
	} else {
		goto L547
	}
L547:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1593 = F_jspOperationName(m, v1592)
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+520)) = int32(89347)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+516)) = v1593
	*(*int32)(unsafe.Add(mBase, uint32(v24)+512)) = v1516
	F_errmsg(m, int32(189987), v24+int32(512))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L1
	} else {
		goto L549
	}
L549:
	;
	F_errfinish(m, int32(500183), int32(1305), int32(108435))
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L551:
	;
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(0) <= v1613 {
		goto L552
	} else {
		goto L553
	}
L552:
	;
	v1616 = int32(0)
	if base.B2i32(l3 == v1616)&(v1611^int32(1)) != 0 {
		v3556 = v1616
		goto L7
	} else {
		goto L555
	}
L553:
	;
	goto L554
L554:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L1
	} else {
		goto L562
	}
L555:
	;
	if v1611 == int32(0) {
		goto L556
	} else {
		goto L557
	}
L556:
	;
	v1629 = F_palloc(m, int32(20))
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L1
	} else {
		goto L559
	}
L557:
	;
	v1631 = v24 + int32(952)
	goto L558
L558:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1631))) = int32(2)
	v1635 = F_int64_to_numeric(m, base.I64_extend_i32_s(v1613-int32(1)))
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L1
	} else {
		goto L560
	}
L559:
	;
	v1631 = v1629
	goto L558
L560:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1631)+4)) = v1635
	v1640 = F_executeNextItem(m, l0, l1, v24+int32(752), v1631, l3, v1611)
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L1
	} else {
		goto L561
	}
L561:
	;
	v3556 = v1640
	goto L7
L562:
	;
	F_errmsg_internal(m, int32(83981), int32(0))
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	F_errfinish(m, int32(500183), int32(1241), int32(108435))
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L565:
	;
	v1664 = m.G0
	v1666 = v1664 - int32(224)
	m.G0 = v1666
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v1668 == int32(18) {
		goto L574
	} else {
		goto L575
	}
L566:
	;
	v1657 = F_JsonbType(m, l2)
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L1
	} else {
		goto L567
	}
L567:
	;
	if v1657 != int32(16) {
		goto L565
	} else {
		goto L568
	}
L568:
	;
	v1662 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L1
	} else {
		goto L569
	}
L569:
	;
	v3556 = v1662
	goto L7
L570:
	;
	v3556 = v1966
	goto L7
L571:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L1
	} else {
		goto L646
	}
L572:
	;
	m.G0 = v1666 + int32(224)
	goto L570
L573:
	;
	if v1672&int32(268435455) == int32(0) {
		goto L587
	} else {
		goto L588
	}
L574:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1671)))
	if v1672&int32(536870912) != 0 {
		goto L573
	} else {
		goto L577
	}
L575:
	;
	goto L576
L576:
	;
	v1681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1681 != int32(1) {
		goto L579
	} else {
		goto L580
	}
L577:
	;
	if v1672&int32(1073741824) == int32(0) {
		goto L571
	} else {
		goto L578
	}
L578:
	;
	goto L576
L579:
	;
	v1966 = int32(2)
	goto L572
L580:
	;
	goto L581
L581:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L1
	} else {
		goto L582
	}
L582:
	;
	F_errcode(m, int32(319553666))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L1
	} else {
		goto L583
	}
L583:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1693 = F_jspOperationName(m, v1692)
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L1
	} else {
		goto L584
	}
L584:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+16)) = v1693
	F_errmsg(m, int32(111011), v1666+int32(16))
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		goto L1
	} else {
		goto L585
	}
L585:
	;
	F_errfinish(m, int32(500183), int32(2840), int32(423257))
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L587:
	;
	v1966 = int32(1)
	goto L572
L588:
	;
	goto L589
L589:
	;
	v1713 = F_jspGetNext(m, l1, v1666+int32(188))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L1
	} else {
		goto L590
	}
L590:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+116)) = int32(22517)
	*(*int64)(unsafe.Add(mBase, uint32(v1666)+108)) = int64(12884901889)
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+96)) = int32(347427)
	*(*int64)(unsafe.Add(mBase, uint32(v1666)+88)) = int64(21474836481)
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+76)) = int32(437750)
	*(*int64)(unsafe.Add(mBase, uint32(v1666)+68)) = int64(8589934593)
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v1727 == int32(18) {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1733 = base.I64_extend_i32_s(v1671 - v1730)
	goto L593
L592:
	;
	v1733 = int64(0)
	goto L593
L593:
	;
	v1734 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+128)) = int32(2)
	v1740 = F_int64_to_numeric(m, v1734*int64(10000000000)+v1733)
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+132)) = v1740
	v1743 = F_JsonbIteratorInit(m, v1671)
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+64)) = v1743
	v1746 = int32(1)
	v1752 = F_JsonbIteratorNext(m, v1666-int32(-64), v1666+int32(168), v1746)
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	if v1752 == int32(0) {
		v1966 = v1746
		goto L572
	} else {
		goto L597
	}
L597:
	;
	v1761 = v1746
	v1763 = v1752
	goto L598
L598:
	;
	if v1763 != int32(1) {
		v1949 = v1761
		goto L601
	} else {
		goto L602
	}
L599:
	;
	v1966 = v1961
	goto L572
L600:
	;
	goto L599
L601:
	;
	v1958 = F_JsonbIteratorNext(m, v1666-int32(-64), v1666+int32(168), int32(1))
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L1
	} else {
		goto L644
	}
L602:
	;
	v1782 = int32(0)
	if base.B2i32(l3 != int32(0))|v1713 == v1782 {
		v1961 = v1782
		goto L600
	} else {
		goto L603
	}
L603:
	;
	v1790 = F_JsonbIteratorNext(m, v1666-int32(-64), v1666+int32(148), int32(1))
	mBase = m.M
	v1791 = m.ExcPending
	if v1791 != 0 {
		goto L1
	} else {
		goto L604
	}
L604:
	;
	v1792 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+40)) = v1792
	v1798 = F_pushJsonbValue(m, v1666+int32(40), int32(6), v1792)
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L1
	} else {
		goto L605
	}
L605:
	;
	v1805 = F_pushJsonbValue(m, v1666+int32(40), int32(1), v1666+int32(108))
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L1
	} else {
		goto L606
	}
L606:
	;
	v1812 = F_pushJsonbValue(m, v1666+int32(40), int32(2), v1666+int32(168))
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L1
	} else {
		goto L607
	}
L607:
	;
	v1819 = F_pushJsonbValue(m, v1666+int32(40), int32(1), v1666+int32(88))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L1
	} else {
		goto L608
	}
L608:
	;
	v1826 = F_pushJsonbValue(m, v1666+int32(40), int32(2), v1666+int32(148))
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L1
	} else {
		goto L609
	}
L609:
	;
	v1833 = F_pushJsonbValue(m, v1666+int32(40), int32(1), v1666+int32(68))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	v1840 = F_pushJsonbValue(m, v1666+int32(40), int32(2), v1666+int32(128))
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L1
	} else {
		goto L611
	}
L611:
	;
	v1846 = F_pushJsonbValue(m, v1666+int32(40), int32(7), int32(0))
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L1
	} else {
		goto L612
	}
L612:
	;
	v1848 = F_JsonbValueToJsonb(m, v1846)
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+44)) = int32(18)
	v1853 = v1848 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+52)) = v1853
	v1855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1848))))
	if v1855 == int32(1) {
		goto L615
	} else {
		goto L616
	}
L614:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+48)) = v1885
	v1887 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1853
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1889
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1889 + int32(1)
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1894 <= int32(0) {
		goto L626
	} else {
		goto L627
	}
L615:
	;
	v1858 = int32(4)
	v1860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1848)+1)))
	if v1860&int32(254) == int32(2) {
		goto L618
	} else {
		goto L619
	}
L616:
	;
	goto L617
L617:
	;
	v1873 = int32(1)
	if v1855&v1873 != 0 {
		v1885 = int32(base.Ui32(v1855)>>(uint(v1873)%32)) - v1873
		goto L614
	} else {
		goto L624
	}
L618:
	;
	v1869 = v1858
	goto L620
L619:
	;
	v1869 = base.B2i32(v1860 == int32(18)) << (uint(v1858) % 32)
	goto L620
L620:
	;
	if v1860 == int32(1) {
		goto L621
	} else {
		goto L622
	}
L621:
	;
	v1872 = v1858
	goto L623
L622:
	;
	v1872 = v1869
	goto L623
L623:
	;
	v1885 = v1872
	goto L614
L624:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1848)))
	v1885 = int32(base.Ui32(v1879)>>(uint(int32(2))%32)) - int32(4)
	goto L614
L625:
	;
	if l3 != 0 {
		v1949 = v1944
		goto L601
	} else {
		goto L642
	}
L626:
	;
	if l3 == int32(0) {
		goto L629
	} else {
		goto L630
	}
L627:
	;
	goto L628
L628:
	;
	v1937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v1938 = F_executeItemOptUnwrapTarget(m, l0, v1666+int32(188), v1666+int32(44), l3, v1937)
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L1
	} else {
		goto L640
	}
L629:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v1887
	v1944 = int32(0)
	goto L625
L630:
	;
	v1900 = F_palloc(m, int32(20))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L1
	} else {
		goto L631
	}
L631:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v1666)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1900)+16)) = v1902
	v1904 = *(*int64)(unsafe.Add(mBase, uint32(v1666)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v1900)+8)) = v1904
	v1906 = *(*int64)(unsafe.Add(mBase, uint32(v1666)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v1900))) = v1906
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1908 != 0 {
		goto L632
	} else {
		goto L633
	}
L632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+216)) = v1900
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+220)) = v1908
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+36)) = v1908
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+32)) = v1900
	v1917 = F_list_make2_impl(m, v1666+int32(36), v1666+int32(32))
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L1
	} else {
		goto L635
	}
L633:
	;
	goto L634
L634:
	;
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v1922 == int32(0) {
		goto L636
	} else {
		goto L637
	}
L635:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v1917
	goto L629
L636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1900
	goto L629
L637:
	;
	goto L638
L638:
	;
	v1926 = F_lappend(m, v1922, v1900)
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v1926
	goto L629
L640:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v1887
	v1941 = int32(2)
	if v1938 == v1941 {
		v1961 = v1941
		goto L600
	} else {
		goto L641
	}
L641:
	;
	v1944 = v1938
	goto L625
L642:
	;
	v1946 = int32(0)
	if v1944 == v1946 {
		v1961 = v1946
		goto L600
	} else {
		goto L643
	}
L643:
	;
	v1949 = v1944
	goto L601
L644:
	;
	if v1958 != 0 {
		v1761 = v1949
		v1763 = v1958
		goto L598
	} else {
		goto L645
	}
L645:
	;
	v1966 = v1949
	goto L572
L646:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v1671)))
	*(*int32)(unsafe.Add(mBase, uint32(v1666))) = v1992
	F_errmsg_internal(m, int32(29593), v1666)
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L1
	} else {
		goto L647
	}
L647:
	;
	F_errfinish(m, int32(500183), int32(3629), int32(371872))
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L1
	} else {
		goto L648
	}
L648:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L649:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v2089 = F_cstring_to_text_with_len(m, v2087, v2088)
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L1
	} else {
		goto L674
	}
L650:
	;
	v2062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2062 != int32(1) {
		goto L666
	} else {
		goto L667
	}
L651:
	;
	v2052 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2052
	*(*int32)(unsafe.Add(mBase, uint32(v24)+984)) = int32(0)
	if v2002 == int32(1) {
		goto L649
	} else {
		goto L665
	}
L652:
	;
	switch v2002 - int32(16) {
	case 0:
		goto L654
	default:
		goto L651
	case 2:
		goto L655
	}
L653:
	;
	v2045 = int32(1)
	v2048 = int32(0)
	v2050 = F_executeAnyItem(m, l0, l1, v2007, l3, v2045, v2045, v2045, v2048, v2048)
	mBase = m.M
	v2051 = m.ExcPending
	if v2051 != 0 {
		goto L1
	} else {
		goto L664
	}
L654:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L1
	} else {
		goto L661
	}
L655:
	;
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v2007)))
	if v2008&int32(536870912) != 0 {
		goto L650
	} else {
		goto L656
	}
L656:
	;
	if v2008&int32(1073741824) != 0 {
		goto L653
	} else {
		goto L657
	}
L657:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		goto L1
	} else {
		goto L658
	}
L658:
	;
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v2007)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+448)) = v2017
	F_errmsg_internal(m, int32(29593), v24+int32(448))
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
	;
	F_errfinish(m, int32(500183), int32(3629), int32(371872))
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
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+464)) = v2033
	F_errmsg_internal(m, int32(485644), v24+int32(464))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L1
	} else {
		goto L662
	}
L662:
	;
	F_errfinish(m, int32(500183), int32(1680), int32(26236))
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L1
	} else {
		goto L663
	}
L663:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L664:
	;
	v3556 = v2050
	goto L7
L665:
	;
	goto L650
L666:
	;
	v3556 = int32(2)
	goto L7
L667:
	;
	goto L668
L668:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L1
	} else {
		goto L669
	}
L669:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L1
	} else {
		goto L670
	}
L670:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2074 = F_jspOperationName(m, v2073)
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L1
	} else {
		goto L671
	}
L671:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+432)) = v2074
	F_errmsg(m, int32(330720), v24+int32(432))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L1
	} else {
		goto L672
	}
L672:
	;
	F_errfinish(m, int32(500183), int32(2357), int32(423279))
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		goto L1
	} else {
		goto L673
	}
L673:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L674:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v2091 - int32(37) {
	case 0:
		goto L681
	default:
		goto L680
	case 8:
		v2160 = v2052
		goto L679
	}
L675:
	;
	v2369 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v2369 - int32(37) {
	case 0:
		v2969 = v2352
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
L676:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2330 = m.ExcPending
	if v2330 != 0 {
		goto L1
	} else {
		goto L737
	}
L677:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2312 = m.ExcPending
	if v2312 != 0 {
		goto L1
	} else {
		goto L733
	}
L678:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L1
	} else {
		goto L730
	}
L679:
	;
	v2162 = *(*int32)(unsafe.Add(mBase, _consts[1074]))
	v2166 = *(*int64)(unsafe.Add(mBase, _consts[1075]))
	v2170 = int32(0)
	goto L700
L680:
	;
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2136 == int32(0) {
		v2160 = v2052
		goto L679
	} else {
		goto L694
	}
L681:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2094 == int32(0) {
		v2160 = v2052
		goto L679
	} else {
		goto L682
	}
L682:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, _consts[1074]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = v2098
	v2101 = *(*int64)(unsafe.Add(mBase, _consts[1075]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v2101
	F_jspGetArg(m, l1, v24+int32(952))
	mBase = m.M
	v2106 = m.ExcPending
	if v2106 != 0 {
		goto L1
	} else {
		goto L683
	}
L683:
	;
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v24)+952))
	if v2107 != int32(1) {
		goto L678
	} else {
		goto L684
	}
L684:
	;
	v2111 = v24 + int32(952)
	v2113 = v24 + int32(936)
	if v2113 != 0 {
		goto L686
	} else {
		goto L687
	}
L685:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v24)+936))
	v2118 = F_cstring_to_text_with_len(m, v2116, v2117)
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L1
	} else {
		goto L689
	}
L686:
	;
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2113))) = v2114
	goto L688
L687:
	;
	goto L688
L688:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+12))
	goto L685
L689:
	;
	v2129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2129 != 0 {
		goto L690
	} else {
		goto L691
	}
L690:
	;
	v2130 = int32(0)
	goto L692
L691:
	;
	v2130 = v24 + int32(752)
	goto L692
L692:
	;
	v2131 = F_parse_datetime(m, v2089, v2118, v24+int32(900), v24+int32(988), v24+int32(984), v2130)
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L1
	} else {
		goto L693
	}
L693:
	;
	v2133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)))
	v2352 = v2131
	v2354 = v2052
	v2356 = v2133 ^ int32(1)
	goto L675
L694:
	;
	F_jspGetArg(m, l1, v24+int32(952))
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L1
	} else {
		goto L695
	}
L695:
	;
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(v24)+952))
	if v2143 != int32(2) {
		goto L677
	} else {
		goto L696
	}
L696:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(952))+12))
	v2151 = F_numeric_int4_opt_error(m, v2148, v24+int32(752))
	mBase = m.M
	v2152 = m.ExcPending
	if v2152 != 0 {
		goto L1
	} else {
		goto L697
	}
L697:
	;
	v2153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+752)))
	if v2153 != int32(1) {
		v2160 = v2151
		goto L679
	} else {
		goto L698
	}
L698:
	;
	v2156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2156 == int32(1) {
		goto L676
	} else {
		goto L699
	}
L699:
	;
	v3556 = int32(2)
	goto L7
L700:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(760)))) = v2162
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v2166
	v2192 = v2170 << (uint(int32(2)) % 32)
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v2192)+uint32(_consts[1076])))
	if v2195 == int32(0) {
		goto L702
	} else {
		goto L703
	}
L701:
	;
	v2232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	v2233 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v2233 == int32(37) {
		goto L709
	} else {
		goto L710
	}
L702:
	;
	v2198 = int32(4515488)
	v2199 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v2202 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2202
	v2206 = *(*int32)(unsafe.Add(mBase, uint32(v2192)+uint32(_consts[1077])))
	v2207 = F_cstring_to_text(m, v2206)
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L1
	} else {
		goto L705
	}
L703:
	;
	v2212 = v2195
	goto L704
L704:
	;
	v2223 = F_parse_datetime(m, v2089, v2212, v24+int32(900), v24+int32(988), v24+int32(984), v24+int32(752))
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L1
	} else {
		goto L706
	}
L705:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2192)+uint32(_consts[1076]))) = v2207
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2199
	v2212 = v2207
	goto L704
L706:
	;
	v2225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)))
	if v2225 == int32(0) {
		v2352 = v2223
		v2354 = v2160
		v2356 = int32(1)
		goto L675
	} else {
		goto L707
	}
L707:
	;
	v2229 = v2170 + int32(1)
	if v2229 != int32(13) {
		v2170 = v2229
		goto L700
	} else {
		goto L708
	}
L708:
	;
	goto L701
L709:
	;
	if v2232&int32(1) == int32(0) {
		goto L712
	} else {
		goto L713
	}
L710:
	;
	goto L711
L711:
	;
	if v2232&int32(1) == int32(0) {
		goto L721
	} else {
		goto L722
	}
L712:
	;
	v3556 = int32(2)
	goto L7
L713:
	;
	goto L714
L714:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		goto L1
	} else {
		goto L715
	}
L715:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2247 = m.ExcPending
	if v2247 != 0 {
		goto L1
	} else {
		goto L716
	}
L716:
	;
	v2248 = F_text_to_cstring(m, v2089)
	mBase = m.M
	v2249 = m.ExcPending
	if v2249 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+164)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v24)+160)) = int32(374681)
	F_errmsg(m, int32(727697), v24+int32(160))
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L1
	} else {
		goto L718
	}
L718:
	;
	F_errhint(m, int32(582914), int32(0))
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L1
	} else {
		goto L719
	}
L719:
	;
	F_errfinish(m, int32(500183), int32(2485), int32(423279))
	mBase = m.M
	v2266 = m.ExcPending
	if v2266 != 0 {
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
	v3556 = int32(2)
	goto L7
L722:
	;
	goto L723
L723:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2275 = m.ExcPending
	if v2275 != 0 {
		goto L1
	} else {
		goto L724
	}
L724:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2278 = m.ExcPending
	if v2278 != 0 {
		goto L1
	} else {
		goto L725
	}
L725:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2280 = F_jspOperationName(m, v2279)
	mBase = m.M
	v2281 = m.ExcPending
	if v2281 != 0 {
		goto L1
	} else {
		goto L726
	}
L726:
	;
	v2282 = F_text_to_cstring(m, v2089)
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L1
	} else {
		goto L727
	}
L727:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+180)) = v2282
	*(*int32)(unsafe.Add(mBase, uint32(v24)+176)) = v2280
	F_errmsg(m, int32(727697), v24+int32(176))
	mBase = m.M
	v2290 = m.ExcPending
	if v2290 != 0 {
		goto L1
	} else {
		goto L728
	}
L728:
	;
	F_errfinish(m, int32(500183), int32(2490), int32(423279))
	mBase = m.M
	v2295 = m.ExcPending
	if v2295 != 0 {
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
	F_errmsg_internal(m, int32(94625), int32(0))
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L1
	} else {
		goto L731
	}
L731:
	;
	F_errfinish(m, int32(500183), int32(2383), int32(423279))
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
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
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2314 = F_jspOperationName(m, v2313)
	mBase = m.M
	v2315 = m.ExcPending
	if v2315 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+416)) = v2314
	F_errmsg_internal(m, int32(94037), v24+int32(416))
	mBase = m.M
	v2321 = m.ExcPending
	if v2321 != 0 {
		goto L1
	} else {
		goto L735
	}
L735:
	;
	F_errfinish(m, int32(500183), int32(2442), int32(423279))
	mBase = m.M
	v2326 = m.ExcPending
	if v2326 != 0 {
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
	v2333 = m.ExcPending
	if v2333 != 0 {
		goto L1
	} else {
		goto L738
	}
L738:
	;
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2335 = F_jspOperationName(m, v2334)
	mBase = m.M
	v2336 = m.ExcPending
	if v2336 != 0 {
		goto L1
	} else {
		goto L739
	}
L739:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+400)) = v2335
	F_errmsg(m, int32(224950), v24+int32(400))
	mBase = m.M
	v2342 = m.ExcPending
	if v2342 != 0 {
		goto L1
	} else {
		goto L740
	}
L740:
	;
	F_errfinish(m, int32(500183), int32(2450), int32(423279))
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
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
	v3055 = m.ExcPending
	if v3055 != 0 {
		goto L1
	} else {
		goto L954
	}
L743:
	;
	v3025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3025 != int32(1) {
		goto L946
	} else {
		goto L947
	}
L744:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3007 = m.ExcPending
	if v3007 != 0 {
		goto L1
	} else {
		goto L941
	}
L745:
	;
	F_pfree(m, v2089)
	mBase = m.M
	v2972 = m.ExcPending
	if v2972 != 0 {
		goto L1
	} else {
		goto L928
	}
L746:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1184)
	v2969 = v2965
	goto L745
L747:
	;
	v2960 = *(*int64)(unsafe.Add(mBase, uint32(v24)+888))
	v2961 = F_Int64GetDatum(m, v2960)
	mBase = m.M
	v2962 = m.ExcPending
	if v2962 != 0 {
		goto L1
	} else {
		goto L927
	}
L748:
	;
	v3556 = int32(2)
	goto L7
L749:
	;
	if v2354 == int32(-1) {
		v2965 = v2933
		goto L746
	} else {
		goto L922
	}
L750:
	;
	v2930 = F_DirectFunctionCall1Coll(m, v2928, int32(0), v2352)
	mBase = m.M
	v2931 = m.ExcPending
	if v2931 != 0 {
		goto L1
	} else {
		goto L921
	}
L751:
	;
	v2915 = *(*int32)(unsafe.Add(mBase, _consts[329]))
	v2917 = m.G0
	v2918 = int32(16)
	v2919 = v2917 - v2918
	m.G0 = v2919
	v2923 = F_DetermineTimeZoneOffsetInternal(m, v24+int32(752), v2915, v2919+int32(8))
	mBase = m.M
	m.G0 = v2919 + v2918
	goto L920
L752:
	;
	v2895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2895, int32(237302), int32(7584))
	mBase = m.M
	v2899 = m.ExcPending
	if v2899 != 0 {
		goto L1
	} else {
		goto L917
	}
L753:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2882 = m.ExcPending
	if v2882 != 0 {
		goto L1
	} else {
		goto L914
	}
L754:
	;
	v2743 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2743 <= int32(1183) {
		goto L893
	} else {
		goto L894
	}
L755:
	;
	v2620 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2620 <= int32(1183) {
		goto L853
	} else {
		goto L854
	}
L756:
	;
	v2543 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2543 <= int32(1183) {
		goto L824
	} else {
		goto L825
	}
L757:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2424 <= int32(1183) {
		goto L787
	} else {
		goto L788
	}
L758:
	;
	v2372 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2372 <= int32(1183) {
		goto L764
	} else {
		goto L765
	}
L759:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1082)
	v2969 = v2420
	goto L745
L760:
	;
	v2418 = F_DirectFunctionCall1Coll(m, v2416, int32(0), v2352)
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L1
	} else {
		goto L779
	}
L761:
	;
	if v2372 != int32(1114) {
		goto L742
	} else {
		goto L778
	}
L762:
	;
	v2407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2407, int32(7584), int32(357362))
	mBase = m.M
	v2411 = m.ExcPending
	if v2411 != 0 {
		goto L1
	} else {
		goto L777
	}
L763:
	;
	v2381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2381 != int32(1) {
		goto L769
	} else {
		goto L770
	}
L764:
	;
	switch v2372 - int32(1082) {
	case 0:
		v2420 = v2352
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
	if v2372 == int32(1184) {
		goto L762
	} else {
		goto L767
	}
L767:
	;
	if v2372 != int32(1266) {
		goto L742
	} else {
		goto L768
	}
L768:
	;
	goto L763
L769:
	;
	v3556 = int32(2)
	goto L7
L770:
	;
	goto L771
L771:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L1
	} else {
		goto L772
	}
L772:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2391 = m.ExcPending
	if v2391 != 0 {
		goto L1
	} else {
		goto L773
	}
L773:
	;
	v2392 = F_text_to_cstring(m, v2089)
	mBase = m.M
	v2393 = m.ExcPending
	if v2393 != 0 {
		goto L1
	} else {
		goto L774
	}
L774:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+228)) = v2392
	*(*int32)(unsafe.Add(mBase, uint32(v24)+224)) = int32(357362)
	F_errmsg(m, int32(727697), v24+int32(224))
	mBase = m.M
	v2401 = m.ExcPending
	if v2401 != 0 {
		goto L1
	} else {
		goto L775
	}
L775:
	;
	F_errfinish(m, int32(500183), int32(2517), int32(423279))
	mBase = m.M
	v2406 = m.ExcPending
	if v2406 != 0 {
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
	v2416 = int32(1426)
	goto L760
L778:
	;
	v2416 = int32(1427)
	goto L760
L779:
	;
	v2420 = v2418
	goto L759
L780:
	;
	if v2354 != int32(-1) {
		goto L806
	} else {
		goto L807
	}
L781:
	;
	v2492 = F_DirectFunctionCall1Coll(m, v2490, int32(0), v2352)
	mBase = m.M
	v2493 = m.ExcPending
	if v2493 != 0 {
		goto L1
	} else {
		goto L805
	}
L782:
	;
	v2485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2485, v2484, int32(376092))
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L1
	} else {
		goto L804
	}
L783:
	;
	v2483 = int32(1430)
	v2484 = int32(7584)
	goto L782
L784:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2468 = m.ExcPending
	if v2468 != 0 {
		goto L1
	} else {
		goto L801
	}
L785:
	;
	if v2424 == int32(1114) {
		v2490 = int32(1429)
		goto L781
	} else {
		goto L800
	}
L786:
	;
	v2435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2435 != int32(1) {
		goto L792
	} else {
		goto L793
	}
L787:
	;
	switch v2424 - int32(1082) {
	case 0:
		goto L786
	case 1:
		v2495 = v2352
		goto L780
	default:
		goto L785
	}
L788:
	;
	goto L789
L789:
	;
	if v2424 == int32(1184) {
		goto L783
	} else {
		goto L790
	}
L790:
	;
	if v2424 != int32(1266) {
		goto L784
	} else {
		goto L791
	}
L791:
	;
	v2483 = int32(1428)
	v2484 = int32(7838)
	goto L782
L792:
	;
	v3556 = int32(2)
	goto L7
L793:
	;
	goto L794
L794:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2442 = m.ExcPending
	if v2442 != 0 {
		goto L1
	} else {
		goto L795
	}
L795:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2445 = m.ExcPending
	if v2445 != 0 {
		goto L1
	} else {
		goto L796
	}
L796:
	;
	v2446 = F_text_to_cstring(m, v2089)
	mBase = m.M
	v2447 = m.ExcPending
	if v2447 != 0 {
		goto L1
	} else {
		goto L797
	}
L797:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+260)) = v2446
	*(*int32)(unsafe.Add(mBase, uint32(v24)+256)) = int32(376092)
	F_errmsg(m, int32(727697), v24+int32(256))
	mBase = m.M
	v2455 = m.ExcPending
	if v2455 != 0 {
		goto L1
	} else {
		goto L798
	}
L798:
	;
	F_errfinish(m, int32(500183), int32(2545), int32(423279))
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
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
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+240)) = v2469
	F_errmsg_internal(m, int32(441783), v24+int32(240))
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L1
	} else {
		goto L802
	}
L802:
	;
	F_errfinish(m, int32(500183), int32(2566), int32(423279))
	mBase = m.M
	v2480 = m.ExcPending
	if v2480 != 0 {
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
	v2490 = v2483
	goto L781
L805:
	;
	v2495 = v2492
	goto L780
L806:
	;
	v2500 = F_anytime_typmod_check(m, int32(0), v2354)
	mBase = m.M
	v2501 = m.ExcPending
	if v2501 != 0 {
		goto L1
	} else {
		goto L809
	}
L807:
	;
	v2540 = v2495
	goto L808
L808:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1083)
	v2969 = v2540
	goto L745
L809:
	;
	v2502 = *(*int64)(unsafe.Add(mBase, uint32(v2495)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v2502
	v2505 = v24 + int32(752)
	if base.Ui32(v2500) <= base.Ui32(int32(6)) {
		goto L811
	} else {
		goto L812
	}
L810:
	;
	v2535 = *(*int64)(unsafe.Add(mBase, uint32(v24)+752))
	v2536 = F_Int64GetDatum(m, v2535)
	mBase = m.M
	v2537 = m.ExcPending
	if v2537 != 0 {
		goto L1
	} else {
		goto L818
	}
L811:
	;
	v2512 = v2500 << (uint(int32(3)) % 32)
	v2515 = *(*int64)(unsafe.Add(mBase, uint32(v2512)+uint32(_consts[1037])))
	v2518 = *(*int64)(unsafe.Add(mBase, uint32(v2512)+uint32(_consts[1038])))
	v2519 = *(*int64)(unsafe.Add(mBase, uint32(v2505)))
	if int64(0) <= v2519 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v2505))) = v2529
	goto L813
L815:
	;
	v2522 = v2518 + v2519
	v2523 = base.I64_rem_s(v2522, v2515)
	v2529 = v2522 - v2523
	goto L814
L816:
	;
	goto L817
L817:
	;
	v2525 = v2518 - v2519
	v2526 = base.I64_rem_s(v2525, v2515)
	v2529 = v2526 - v2525
	goto L814
L818:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2500
	v2540 = v2536
	goto L808
L819:
	;
	if v2354 != int32(-1) {
		goto L835
	} else {
		goto L836
	}
L820:
	;
	v2579 = F_DirectFunctionCall1Coll(m, v2577, int32(0), v2352)
	mBase = m.M
	v2580 = m.ExcPending
	if v2580 != 0 {
		goto L1
	} else {
		goto L834
	}
L821:
	;
	v2571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2571, int32(376092), int32(7838))
	mBase = m.M
	v2575 = m.ExcPending
	if v2575 != 0 {
		goto L1
	} else {
		goto L833
	}
L822:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2558 = m.ExcPending
	if v2558 != 0 {
		goto L1
	} else {
		goto L830
	}
L823:
	;
	if v2543 == int32(1114) {
		goto L743
	} else {
		goto L829
	}
L824:
	;
	switch v2543 - int32(1082) {
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
	if v2543 == int32(1184) {
		v2577 = int32(1431)
		goto L820
	} else {
		goto L827
	}
L827:
	;
	if v2543 != int32(1266) {
		goto L822
	} else {
		goto L828
	}
L828:
	;
	v2581 = v2352
	goto L819
L829:
	;
	goto L822
L830:
	;
	v2559 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+272)) = v2559
	F_errmsg_internal(m, int32(441783), v24+int32(272))
	mBase = m.M
	v2565 = m.ExcPending
	if v2565 != 0 {
		goto L1
	} else {
		goto L831
	}
L831:
	;
	F_errfinish(m, int32(500183), int32(2613), int32(423279))
	mBase = m.M
	v2570 = m.ExcPending
	if v2570 != 0 {
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
	v2577 = int32(1432)
	goto L820
L834:
	;
	v2581 = v2579
	goto L819
L835:
	;
	v2585 = F_anytime_typmod_check(m, int32(1), v2354)
	mBase = m.M
	v2586 = m.ExcPending
	if v2586 != 0 {
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
	v2969 = v2581
	goto L745
L838:
	;
	if base.Ui32(v2585) <= base.Ui32(int32(6)) {
		goto L840
	} else {
		goto L841
	}
L839:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2585
	goto L837
L840:
	;
	v2593 = v2585 << (uint(int32(3)) % 32)
	v2596 = *(*int64)(unsafe.Add(mBase, uint32(v2593)+uint32(_consts[1037])))
	v2599 = *(*int64)(unsafe.Add(mBase, uint32(v2593)+uint32(_consts[1038])))
	v2600 = *(*int64)(unsafe.Add(mBase, uint32(v2581)))
	if int64(0) <= v2600 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v2581))) = v2610
	goto L842
L844:
	;
	v2603 = v2599 + v2600
	v2604 = base.I64_rem_s(v2603, v2596)
	v2610 = v2603 - v2604
	goto L843
L845:
	;
	goto L846
L846:
	;
	v2606 = v2599 - v2600
	v2607 = base.I64_rem_s(v2606, v2596)
	v2610 = v2607 - v2606
	goto L843
L847:
	;
	if v2354 != int32(-1) {
		goto L872
	} else {
		goto L873
	}
L848:
	;
	v2684 = F_DirectFunctionCall1Coll(m, v2682, int32(0), v2352)
	mBase = m.M
	v2685 = m.ExcPending
	if v2685 != 0 {
		goto L1
	} else {
		goto L871
	}
L849:
	;
	v2676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2676, int32(7584), int32(237302))
	mBase = m.M
	v2680 = m.ExcPending
	if v2680 != 0 {
		goto L1
	} else {
		goto L870
	}
L850:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2663 = m.ExcPending
	if v2663 != 0 {
		goto L1
	} else {
		goto L867
	}
L851:
	;
	if v2620 == int32(1114) {
		v2686 = v2352
		goto L847
	} else {
		goto L866
	}
L852:
	;
	v2631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2631 != int32(1) {
		goto L858
	} else {
		goto L859
	}
L853:
	;
	switch v2620 - int32(1082) {
	case 0:
		v2682 = int32(1433)
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
	if v2620 == int32(1184) {
		goto L849
	} else {
		goto L856
	}
L856:
	;
	if v2620 != int32(1266) {
		goto L850
	} else {
		goto L857
	}
L857:
	;
	goto L852
L858:
	;
	v3556 = int32(2)
	goto L7
L859:
	;
	goto L860
L860:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2638 = m.ExcPending
	if v2638 != 0 {
		goto L1
	} else {
		goto L861
	}
L861:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2641 = m.ExcPending
	if v2641 != 0 {
		goto L1
	} else {
		goto L862
	}
L862:
	;
	v2642 = F_text_to_cstring(m, v2089)
	mBase = m.M
	v2643 = m.ExcPending
	if v2643 != 0 {
		goto L1
	} else {
		goto L863
	}
L863:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+340)) = v2642
	*(*int32)(unsafe.Add(mBase, uint32(v24)+336)) = int32(237302)
	F_errmsg(m, int32(727697), v24+int32(336))
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
		goto L1
	} else {
		goto L864
	}
L864:
	;
	F_errfinish(m, int32(500183), int32(2649), int32(423279))
	mBase = m.M
	v2656 = m.ExcPending
	if v2656 != 0 {
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
	v2664 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+304)) = v2664
	F_errmsg_internal(m, int32(441783), v24+int32(304))
	mBase = m.M
	v2670 = m.ExcPending
	if v2670 != 0 {
		goto L1
	} else {
		goto L868
	}
L868:
	;
	F_errfinish(m, int32(500183), int32(2660), int32(423279))
	mBase = m.M
	v2675 = m.ExcPending
	if v2675 != 0 {
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
	v2682 = int32(1434)
	goto L848
L871:
	;
	v2686 = v2684
	goto L847
L872:
	;
	v2691 = *(*int32)(unsafe.Add(mBase, _consts[1074]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = v2691
	v2694 = *(*int64)(unsafe.Add(mBase, _consts[1075]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v2694
	v2697 = F_anytimestamp_typmod_check(m, int32(0), v2354)
	mBase = m.M
	v2698 = m.ExcPending
	if v2698 != 0 {
		goto L1
	} else {
		goto L875
	}
L873:
	;
	v2740 = v2686
	goto L874
L874:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1114)
	v2969 = v2740
	goto L745
L875:
	;
	v2699 = *(*int64)(unsafe.Add(mBase, uint32(v2686)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+936)) = v2699
	F_AdjustTimestampForTypmod(m, v24+int32(936), v2697, v24+int32(752))
	mBase = m.M
	v2706 = m.ExcPending
	if v2706 != 0 {
		goto L1
	} else {
		goto L876
	}
L876:
	;
	v2707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)))
	if v2707 == int32(1) {
		goto L877
	} else {
		goto L878
	}
L877:
	;
	v2710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2710 == int32(0) {
		goto L880
	} else {
		goto L881
	}
L878:
	;
	goto L879
L879:
	;
	v2735 = *(*int64)(unsafe.Add(mBase, uint32(v24)+936))
	v2736 = F_Int64GetDatum(m, v2735)
	mBase = m.M
	v2737 = m.ExcPending
	if v2737 != 0 {
		goto L1
	} else {
		goto L888
	}
L880:
	;
	v3556 = int32(2)
	goto L7
L881:
	;
	goto L882
L882:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2717 = m.ExcPending
	if v2717 != 0 {
		goto L1
	} else {
		goto L883
	}
L883:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2720 = m.ExcPending
	if v2720 != 0 {
		goto L1
	} else {
		goto L884
	}
L884:
	;
	v2721 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2722 = F_jspOperationName(m, v2721)
	mBase = m.M
	v2723 = m.ExcPending
	if v2723 != 0 {
		goto L1
	} else {
		goto L885
	}
L885:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+320)) = v2722
	F_errmsg(m, int32(435593), v24+int32(320))
	mBase = m.M
	v2729 = m.ExcPending
	if v2729 != 0 {
		goto L1
	} else {
		goto L886
	}
L886:
	;
	F_errfinish(m, int32(500183), int32(2679), int32(423279))
	mBase = m.M
	v2734 = m.ExcPending
	if v2734 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2697
	v2740 = v2736
	goto L874
L889:
	;
	v2795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2795, int32(357362), int32(7584))
	mBase = m.M
	v2799 = m.ExcPending
	if v2799 != 0 {
		goto L1
	} else {
		goto L908
	}
L890:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2782 = m.ExcPending
	if v2782 != 0 {
		goto L1
	} else {
		goto L905
	}
L891:
	;
	if v2743 == int32(1114) {
		goto L752
	} else {
		goto L904
	}
L892:
	;
	v2752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2752 != int32(1) {
		goto L748
	} else {
		goto L898
	}
L893:
	;
	switch v2743 - int32(1082) {
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
	if v2743 == int32(1184) {
		v2933 = v2352
		goto L749
	} else {
		goto L896
	}
L896:
	;
	if v2743 != int32(1266) {
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
	v2758 = m.ExcPending
	if v2758 != 0 {
		goto L1
	} else {
		goto L899
	}
L899:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2761 = m.ExcPending
	if v2761 != 0 {
		goto L1
	} else {
		goto L900
	}
L900:
	;
	v2762 = F_text_to_cstring(m, v2089)
	mBase = m.M
	v2763 = m.ExcPending
	if v2763 != 0 {
		goto L1
	} else {
		goto L901
	}
L901:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+388)) = v2762
	*(*int32)(unsafe.Add(mBase, uint32(v24)+384)) = int32(7938)
	F_errmsg(m, int32(727697), v24+int32(384))
	mBase = m.M
	v2771 = m.ExcPending
	if v2771 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	F_errfinish(m, int32(500183), int32(2720), int32(423279))
	mBase = m.M
	v2776 = m.ExcPending
	if v2776 != 0 {
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
	v2783 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+352)) = v2783
	F_errmsg_internal(m, int32(441783), v24+int32(352))
	mBase = m.M
	v2789 = m.ExcPending
	if v2789 != 0 {
		goto L1
	} else {
		goto L906
	}
L906:
	;
	F_errfinish(m, int32(500183), int32(2741), int32(423279))
	mBase = m.M
	v2794 = m.ExcPending
	if v2794 != 0 {
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
	v2811 = v2352 + int32(2483589)
	v2812 = int32(146097)
	v2813 = base.I32_div_u_s(v2811, v2812)
	v2814 = int32(3)
	v2820 = int32(2)
	v2825 = base.I32_div_u_s((v2813*int32(1073595727)+v2811)<<(uint(v2820)%32)|v2814, v2812)
	v2828 = v2352 + int32(2451545) + v2813*v2814 + v2825 + int32(32104)
	v2829 = int32(1461)
	v2830 = base.I32_div_u_s(v2828, v2829)
	v2833 = v2830*int32(-1461) + v2828
	v2835 = v2833 << (uint(v2820) % 32)
	if base.Ui32(v2829) <= base.Ui32(v2835) {
		goto L911
	} else {
		goto L912
	}
L909:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = int64(0)
	v2911 = int32(1435)
	goto L751
L910:
	;
	v2848 = base.I32_div_u_s(v2835, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(772)))) = v2848 + v2830<<(uint(int32(2))%32) - int32(4800)
	v2856 = v2846 + int32(123)
	v2860 = int32(base.Ui32(v2856*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(764)))) = v2856 - int32(base.Ui32(v2860*int32(7834))>>(uint(int32(8))%32))
	v2870 = base.I32_rem_u_s(v2860+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(768)))) = v2870 + int32(1)
	goto L909
L911:
	;
	v2841 = base.I32_rem_u_s(v2833+int32(305), int32(365))
	v2846 = v2841
	goto L910
L912:
	;
	goto L913
L913:
	;
	v2845 = base.I32_rem_u_s(v2833+int32(306), int32(366))
	v2846 = v2845
	goto L910
L914:
	;
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+192)) = v2883
	F_errmsg_internal(m, int32(485419), v24+int32(192))
	mBase = m.M
	v2889 = m.ExcPending
	if v2889 != 0 {
		goto L1
	} else {
		goto L915
	}
L915:
	;
	F_errfinish(m, int32(500183), int32(2771), int32(423279))
	mBase = m.M
	v2894 = m.ExcPending
	if v2894 != 0 {
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
	v2900 = int32(1436)
	v2901 = *(*int64)(unsafe.Add(mBase, uint32(v2352)))
	v2902 = int32(0)
	v2909 = F_timestamp2tm(m, v2901, v2902, v24+int32(752), v24+int32(948), v2902, v2902)
	mBase = m.M
	v2910 = m.ExcPending
	if v2910 != 0 {
		goto L1
	} else {
		goto L918
	}
L918:
	;
	if v2909 != 0 {
		v2928 = v2900
		goto L750
	} else {
		goto L919
	}
L919:
	;
	v2911 = v2900
	goto L751
L920:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+984)) = v2923
	v2928 = v2911
	goto L750
L921:
	;
	v2933 = v2930
	goto L749
L922:
	;
	v2937 = *(*int32)(unsafe.Add(mBase, _consts[1074]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+944)) = v2937
	v2940 = *(*int64)(unsafe.Add(mBase, _consts[1075]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+936)) = v2940
	v2943 = F_anytimestamp_typmod_check(m, int32(1), v2354)
	mBase = m.M
	v2944 = m.ExcPending
	if v2944 != 0 {
		goto L1
	} else {
		goto L923
	}
L923:
	;
	v2945 = *(*int64)(unsafe.Add(mBase, uint32(v2933)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+888)) = v2945
	F_AdjustTimestampForTypmod(m, v24+int32(888), v2943, v24+int32(936))
	mBase = m.M
	v2952 = m.ExcPending
	if v2952 != 0 {
		goto L1
	} else {
		goto L924
	}
L924:
	;
	v2953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+940)))
	if v2953 != int32(1) {
		goto L747
	} else {
		goto L925
	}
L925:
	;
	v2956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2956 != 0 {
		goto L744
	} else {
		goto L926
	}
L926:
	;
	goto L748
L927:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2943
	v2965 = v2961
	goto L746
L928:
	;
	if v2356&int32(1) == int32(0) {
		goto L929
	} else {
		goto L930
	}
L929:
	;
	v3556 = int32(2)
	goto L7
L930:
	;
	goto L931
L931:
	;
	v2980 = F_jspGetNext(m, l1, v24+int32(952))
	mBase = m.M
	v2981 = m.ExcPending
	if v2981 != 0 {
		goto L1
	} else {
		goto L932
	}
L932:
	;
	if l3 != 0 {
		goto L933
	} else {
		goto L934
	}
L933:
	;
	if v2980 == int32(0) {
		goto L936
	} else {
		goto L937
	}
L934:
	;
	if v2980 != 0 {
		goto L933
	} else {
		goto L935
	}
L935:
	;
	v3556 = int32(0)
	goto L7
L936:
	;
	v2988 = F_palloc(m, int32(20))
	mBase = m.M
	v2989 = m.ExcPending
	if v2989 != 0 {
		goto L1
	} else {
		goto L939
	}
L937:
	;
	v2990 = v24 + int32(908)
	goto L938
L938:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2990)+4)) = v2969
	*(*int32)(unsafe.Add(mBase, uint32(v2990))) = int32(32)
	v2994 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v2990)+8)) = v2994
	v2996 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	*(*int32)(unsafe.Add(mBase, uint32(v2990)+12)) = v2996
	v2998 = *(*int32)(unsafe.Add(mBase, uint32(v24)+984))
	*(*int32)(unsafe.Add(mBase, uint32(v2990)+16)) = v2998
	v3002 = F_executeNextItem(m, l0, l1, v24+int32(952), v2990, l3, v2980)
	mBase = m.M
	v3003 = m.ExcPending
	if v3003 != 0 {
		goto L1
	} else {
		goto L940
	}
L939:
	;
	v2990 = v2988
	goto L938
L940:
	;
	v3556 = v3002
	goto L7
L941:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v3010 = m.ExcPending
	if v3010 != 0 {
		goto L1
	} else {
		goto L942
	}
L942:
	;
	v3011 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3012 = F_jspOperationName(m, v3011)
	mBase = m.M
	v3013 = m.ExcPending
	if v3013 != 0 {
		goto L1
	} else {
		goto L943
	}
L943:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+368)) = v3012
	F_errmsg(m, int32(435593), v24+int32(368))
	mBase = m.M
	v3019 = m.ExcPending
	if v3019 != 0 {
		goto L1
	} else {
		goto L944
	}
L944:
	;
	F_errfinish(m, int32(500183), int32(2760), int32(423279))
	mBase = m.M
	v3024 = m.ExcPending
	if v3024 != 0 {
		goto L1
	} else {
		goto L945
	}
L945:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L946:
	;
	v3556 = int32(2)
	goto L7
L947:
	;
	goto L948
L948:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3032 = m.ExcPending
	if v3032 != 0 {
		goto L1
	} else {
		goto L949
	}
L949:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v3035 = m.ExcPending
	if v3035 != 0 {
		goto L1
	} else {
		goto L950
	}
L950:
	;
	v3036 = F_text_to_cstring(m, v2089)
	mBase = m.M
	v3037 = m.ExcPending
	if v3037 != 0 {
		goto L1
	} else {
		goto L951
	}
L951:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+292)) = v3036
	*(*int32)(unsafe.Add(mBase, uint32(v24)+288)) = int32(7971)
	F_errmsg(m, int32(727697), v24+int32(288))
	mBase = m.M
	v3045 = m.ExcPending
	if v3045 != 0 {
		goto L1
	} else {
		goto L952
	}
L952:
	;
	F_errfinish(m, int32(500183), int32(2598), int32(423279))
	mBase = m.M
	v3050 = m.ExcPending
	if v3050 != 0 {
		goto L1
	} else {
		goto L953
	}
L953:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L954:
	;
	v3056 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+208)) = v3056
	F_errmsg_internal(m, int32(441783), v24+int32(208))
	mBase = m.M
	v3062 = m.ExcPending
	if v3062 != 0 {
		goto L1
	} else {
		goto L955
	}
L955:
	;
	F_errfinish(m, int32(500183), int32(2530), int32(423279))
	mBase = m.M
	v3067 = m.ExcPending
	if v3067 != 0 {
		goto L1
	} else {
		goto L956
	}
L956:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L957:
	;
	v3077 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v3077 - int32(1) {
	case 0:
		goto L967
	case 1:
		goto L968
	default:
		goto L966
	}
L958:
	;
	v3070 = F_JsonbType(m, l2)
	mBase = m.M
	v3071 = m.ExcPending
	if v3071 != 0 {
		goto L1
	} else {
		goto L959
	}
L959:
	;
	if v3070 != int32(16) {
		goto L957
	} else {
		goto L960
	}
L960:
	;
	v3075 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v3076 = m.ExcPending
	if v3076 != 0 {
		goto L1
	} else {
		goto L961
	}
L961:
	;
	v3556 = v3075
	goto L7
L962:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3267 = m.ExcPending
	if v3267 != 0 {
		goto L1
	} else {
		goto L1012
	}
L963:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3246 = m.ExcPending
	if v3246 != 0 {
		goto L1
	} else {
		goto L1007
	}
L964:
	;
	v3241 = F_executeNextItem(m, l0, l1, int32(0), v3235, l3, int32(1))
	mBase = m.M
	v3242 = m.ExcPending
	if v3242 != 0 {
		goto L1
	} else {
		goto L1006
	}
L965:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = int32(2)
	v3226 = F_Float8GetDatum(m, v3153)
	mBase = m.M
	v3227 = m.ExcPending
	if v3227 != 0 {
		goto L1
	} else {
		goto L1003
	}
L966:
	;
	v3198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3198 != int32(1) {
		v3556 = int32(2)
		goto L7
	} else {
		goto L997
	}
L967:
	;
	v3139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3140 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v3141 = F_pnstrdup(m, v3139, v3140)
	mBase = m.M
	v3142 = m.ExcPending
	if v3142 != 0 {
		goto L1
	} else {
		goto L983
	}
L968:
	;
	v3082 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v3083 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v3082)
	mBase = m.M
	v3084 = m.ExcPending
	if v3084 != 0 {
		goto L1
	} else {
		goto L969
	}
L969:
	;
	v3086 = *(*int32)(unsafe.Add(mBase, _consts[1074]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v3086
	v3089 = *(*int64)(unsafe.Add(mBase, _consts[1075]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v3089
	v3095 = F_float8in_internal(m, v3083, int32(0), int32(271930), v3083, v24+int32(952))
	mBase = m.M
	v3096 = m.ExcPending
	if v3096 != 0 {
		goto L1
	} else {
		goto L970
	}
L970:
	;
	v3097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v3097 == int32(1) {
		goto L972
	} else {
		goto L973
	}
L971:
	;
	v3556 = int32(2)
	goto L7
L972:
	;
	v3100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3100 != int32(1) {
		goto L971
	} else {
		goto L975
	}
L973:
	;
	goto L974
L974:
	;
	v3127 = base.F64_abs(v3095)
	if base.F64_ne(v3127, math.Float64frombits(uint64(0x7ff0000000000000)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v3127)) < base.Ui64(int64(9218868437227405313))) != 0 {
		v3235 = l2
		goto L964
	} else {
		goto L981
	}
L975:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3106 = m.ExcPending
	if v3106 != 0 {
		goto L1
	} else {
		goto L976
	}
L976:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3109 = m.ExcPending
	if v3109 != 0 {
		goto L1
	} else {
		goto L977
	}
L977:
	;
	v3110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3111 = F_jspOperationName(m, v3110)
	mBase = m.M
	v3112 = m.ExcPending
	if v3112 != 0 {
		goto L1
	} else {
		goto L978
	}
L978:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+104)) = int32(271930)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+100)) = v3111
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v3083
	F_errmsg(m, int32(189987), v24+int32(96))
	mBase = m.M
	v3121 = m.ExcPending
	if v3121 != 0 {
		goto L1
	} else {
		goto L979
	}
L979:
	;
	F_errfinish(m, int32(500183), int32(1166), int32(108435))
	mBase = m.M
	v3126 = m.ExcPending
	if v3126 != 0 {
		goto L1
	} else {
		goto L980
	}
L980:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L981:
	;
	v3134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3134 == int32(1) {
		goto L963
	} else {
		goto L982
	}
L982:
	;
	goto L971
L983:
	;
	v3144 = *(*int32)(unsafe.Add(mBase, _consts[1074]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v3144
	v3147 = *(*int64)(unsafe.Add(mBase, _consts[1075]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v3147
	v3153 = F_float8in_internal(m, v3141, int32(0), int32(271930), v3141, v24+int32(952))
	mBase = m.M
	v3154 = m.ExcPending
	if v3154 != 0 {
		goto L1
	} else {
		goto L984
	}
L984:
	;
	v3155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v3155 == int32(1) {
		goto L986
	} else {
		goto L987
	}
L985:
	;
	v3556 = int32(2)
	goto L7
L986:
	;
	v3158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3158 != int32(1) {
		goto L985
	} else {
		goto L989
	}
L987:
	;
	goto L988
L988:
	;
	v3185 = base.F64_abs(v3153)
	if base.F64_ne(v3185, math.Float64frombits(uint64(0x7ff0000000000000)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v3185)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L965
	} else {
		goto L995
	}
L989:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3164 = m.ExcPending
	if v3164 != 0 {
		goto L1
	} else {
		goto L990
	}
L990:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3167 = m.ExcPending
	if v3167 != 0 {
		goto L1
	} else {
		goto L991
	}
L991:
	;
	v3168 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3169 = F_jspOperationName(m, v3168)
	mBase = m.M
	v3170 = m.ExcPending
	if v3170 != 0 {
		goto L1
	} else {
		goto L992
	}
L992:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+136)) = int32(271930)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+132)) = v3169
	*(*int32)(unsafe.Add(mBase, uint32(v24)+128)) = v3141
	F_errmsg(m, int32(189987), v24+int32(128))
	mBase = m.M
	v3179 = m.ExcPending
	if v3179 != 0 {
		goto L1
	} else {
		goto L993
	}
L993:
	;
	F_errfinish(m, int32(500183), int32(1192), int32(108435))
	mBase = m.M
	v3184 = m.ExcPending
	if v3184 != 0 {
		goto L1
	} else {
		goto L994
	}
L994:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L995:
	;
	v3192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3192 == int32(1) {
		goto L962
	} else {
		goto L996
	}
L996:
	;
	goto L985
L997:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3204 = m.ExcPending
	if v3204 != 0 {
		goto L1
	} else {
		goto L998
	}
L998:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3207 = m.ExcPending
	if v3207 != 0 {
		goto L1
	} else {
		goto L999
	}
L999:
	;
	v3208 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3209 = F_jspOperationName(m, v3208)
	mBase = m.M
	v3210 = m.ExcPending
	if v3210 != 0 {
		goto L1
	} else {
		goto L1000
	}
L1000:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v3209
	F_errmsg(m, int32(346373), v24+int32(80))
	mBase = m.M
	v3216 = m.ExcPending
	if v3216 != 0 {
		goto L1
	} else {
		goto L1001
	}
L1001:
	;
	F_errfinish(m, int32(500183), int32(1210), int32(108435))
	mBase = m.M
	v3221 = m.ExcPending
	if v3221 != 0 {
		goto L1
	} else {
		goto L1002
	}
L1002:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1003:
	;
	v3228 = F_DirectFunctionCall1Coll(m, int32(1417), int32(0), v3226)
	mBase = m.M
	v3229 = m.ExcPending
	if v3229 != 0 {
		goto L1
	} else {
		goto L1004
	}
L1004:
	;
	v3230 = F_pg_detoast_datum(m, v3228)
	mBase = m.M
	v3231 = m.ExcPending
	if v3231 != 0 {
		goto L1
	} else {
		goto L1005
	}
L1005:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+756)) = v3230
	v3235 = v24 + int32(752)
	goto L964
L1006:
	;
	v3556 = v3241
	goto L7
L1007:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3249 = m.ExcPending
	if v3249 != 0 {
		goto L1
	} else {
		goto L1008
	}
L1008:
	;
	v3250 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3251 = F_jspOperationName(m, v3250)
	mBase = m.M
	v3252 = m.ExcPending
	if v3252 != 0 {
		goto L1
	} else {
		goto L1009
	}
L1009:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = v3251
	F_errmsg(m, int32(683612), v24+int32(112))
	mBase = m.M
	v3258 = m.ExcPending
	if v3258 != 0 {
		goto L1
	} else {
		goto L1010
	}
L1010:
	;
	F_errfinish(m, int32(500183), int32(1171), int32(108435))
	mBase = m.M
	v3263 = m.ExcPending
	if v3263 != 0 {
		goto L1
	} else {
		goto L1011
	}
L1011:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1012:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3270 = m.ExcPending
	if v3270 != 0 {
		goto L1
	} else {
		goto L1013
	}
L1013:
	;
	v3271 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3272 = F_jspOperationName(m, v3271)
	mBase = m.M
	v3273 = m.ExcPending
	if v3273 != 0 {
		goto L1
	} else {
		goto L1014
	}
L1014:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+144)) = v3272
	F_errmsg(m, int32(683612), v24+int32(144))
	mBase = m.M
	v3279 = m.ExcPending
	if v3279 != 0 {
		goto L1
	} else {
		goto L1015
	}
L1015:
	;
	F_errfinish(m, int32(500183), int32(1197), int32(108435))
	mBase = m.M
	v3284 = m.ExcPending
	if v3284 != 0 {
		goto L1
	} else {
		goto L1016
	}
L1016:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1017:
	;
	v3556 = v3286
	goto L7
L1018:
	;
	v3556 = v3289
	goto L7
L1019:
	;
	v3556 = v3292
	goto L7
L1020:
	;
	v3337 = F_palloc(m, int32(20))
	mBase = m.M
	v3338 = m.ExcPending
	if v3338 != 0 {
		goto L1
	} else {
		goto L1034
	}
L1021:
	;
	v3306 = int32(1)
	v3307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v3307 != 0 {
		v3335 = v3306
		goto L1020
	} else {
		goto L1024
	}
L1022:
	;
	v3297 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3298 = *(*int32)(unsafe.Add(mBase, uint32(v3297)))
	if v3298&int32(1342177280) != int32(1073741824) {
		goto L1021
	} else {
		goto L1023
	}
L1023:
	;
	v3335 = v3298 & int32(268435455)
	goto L1020
L1024:
	;
	v3308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3308 != 0 {
		v3556 = v3306
		goto L7
	} else {
		goto L1025
	}
L1025:
	;
	v3309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3309 != int32(1) {
		goto L1026
	} else {
		goto L1027
	}
L1026:
	;
	v3556 = int32(2)
	goto L7
L1027:
	;
	goto L1028
L1028:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3316 = m.ExcPending
	if v3316 != 0 {
		goto L1
	} else {
		goto L1029
	}
L1029:
	;
	F_errcode(m, int32(151781506))
	mBase = m.M
	v3319 = m.ExcPending
	if v3319 != 0 {
		goto L1
	} else {
		goto L1030
	}
L1030:
	;
	v3320 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3321 = F_jspOperationName(m, v3320)
	mBase = m.M
	v3322 = m.ExcPending
	if v3322 != 0 {
		goto L1
	} else {
		goto L1031
	}
L1031:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v3321
	F_errmsg(m, int32(25426), v24-int32(-64))
	mBase = m.M
	v3328 = m.ExcPending
	if v3328 != 0 {
		goto L1
	} else {
		goto L1032
	}
L1032:
	;
	F_errfinish(m, int32(500183), int32(1113), int32(108435))
	mBase = m.M
	v3333 = m.ExcPending
	if v3333 != 0 {
		goto L1
	} else {
		goto L1033
	}
L1033:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1034:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3337))) = int32(2)
	v3342 = F_int64_to_numeric(m, base.I64_extend_i32_u(v3335))
	mBase = m.M
	v3343 = m.ExcPending
	if v3343 != 0 {
		goto L1
	} else {
		goto L1035
	}
L1035:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3337)+4)) = v3342
	v3345 = int32(0)
	v3347 = F_executeNextItem(m, l0, l1, v3345, v3337, l3, v3345)
	mBase = m.M
	v3348 = m.ExcPending
	if v3348 != 0 {
		goto L1
	} else {
		goto L1036
	}
L1036:
	;
	v3556 = v3347
	goto L7
L1037:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3350))) = int32(1)
	v3354 = F_JsonbTypeName(m, l2)
	mBase = m.M
	v3355 = m.ExcPending
	if v3355 != 0 {
		goto L1
	} else {
		goto L1038
	}
L1038:
	;
	v3356 = F_pstrdup(m, v3354)
	mBase = m.M
	v3357 = m.ExcPending
	if v3357 != 0 {
		goto L1
	} else {
		goto L1039
	}
L1039:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3350)+8)) = v3356
	v3359 = F_strlen(m, v3356)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3350)+4)) = v3359
	v3361 = int32(0)
	v3363 = F_executeNextItem(m, l0, l1, v3361, v3350, l3, v3361)
	mBase = m.M
	v3364 = m.ExcPending
	if v3364 != 0 {
		goto L1
	} else {
		goto L1040
	}
L1040:
	;
	v3556 = v3363
	goto L7
L1041:
	;
	v3389 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v3390 = m.ExcPending
	if v3390 != 0 {
		goto L1
	} else {
		goto L1051
	}
L1042:
	;
	v3365 = F_JsonbType(m, l2)
	mBase = m.M
	v3366 = m.ExcPending
	if v3366 != 0 {
		goto L1
	} else {
		goto L1045
	}
L1043:
	;
	goto L1044
L1044:
	;
	F_jspGetArg(m, l1, v24+int32(752))
	mBase = m.M
	v3372 = m.ExcPending
	if v3372 != 0 {
		goto L1
	} else {
		goto L1047
	}
L1045:
	;
	if v3365 == int32(16) {
		goto L1041
	} else {
		goto L1046
	}
L1046:
	;
	goto L1044
L1047:
	;
	v3373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
	v3378 = F_executeBoolItem(m, l0, v24+int32(752), l2, int32(0))
	mBase = m.M
	v3379 = m.ExcPending
	if v3379 != 0 {
		goto L1
	} else {
		goto L1048
	}
L1048:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v3373
	v3381 = int32(1)
	if v3378 != v3381 {
		v3556 = v3381
		goto L7
	} else {
		goto L1049
	}
L1049:
	;
	v3386 = F_executeNextItem(m, l0, l1, int32(0), l2, l3, int32(1))
	mBase = m.M
	v3387 = m.ExcPending
	if v3387 != 0 {
		goto L1
	} else {
		goto L1050
	}
L1050:
	;
	v3556 = v3386
	goto L7
L1051:
	;
	v3556 = v3389
	goto L7
L1052:
	;
	v3397 = *(*int32)(unsafe.Add(mBase, uint32(v3393)+8))
	v3398 = v3397
	goto L1054
L1053:
	;
	v3398 = int32(0)
	goto L1054
L1054:
	;
	v3399 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3399
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3398
	v3404 = F_executeNextItem(m, l0, l1, v3399, v3393, l3, int32(1))
	mBase = m.M
	v3405 = m.ExcPending
	if v3405 != 0 {
		goto L1
	} else {
		goto L1055
	}
L1055:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v3391
	v3556 = v3404
	goto L7
L1056:
	;
	v3556 = v3410
	goto L7
L1057:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3479 = m.ExcPending
	if v3479 != 0 {
		goto L1
	} else {
		goto L1093
	}
L1058:
	;
	if v3412 == int32(17) {
		goto L1059
	} else {
		goto L1060
	}
L1059:
	;
	v3416 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = v3416
	v3420 = v24 + int32(756)
	if v3420 != 0 {
		goto L1063
	} else {
		goto L1064
	}
L1060:
	;
	goto L1061
L1061:
	;
	if l4 == int32(0) {
		goto L1078
	} else {
		goto L1079
	}
L1062:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = v3423
	v3425 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3429 = F_findJsonbValueFromContainer(m, v3425, int32(536870912), v24+int32(752))
	mBase = m.M
	v3430 = m.ExcPending
	if v3430 != 0 {
		goto L1
	} else {
		goto L1066
	}
L1063:
	;
	v3421 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v3420))) = v3421
	goto L1065
L1064:
	;
	goto L1065
L1065:
	;
	v3423 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L1062
L1066:
	;
	if v3429 != 0 {
		goto L1067
	} else {
		goto L1068
	}
L1067:
	;
	v3431 = int32(0)
	v3433 = F_executeNextItem(m, l0, l1, v3431, v3429, l3, v3431)
	mBase = m.M
	v3434 = m.ExcPending
	if v3434 != 0 {
		goto L1
	} else {
		goto L1070
	}
L1068:
	;
	goto L1069
L1069:
	;
	v3440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3440 != 0 {
		v3556 = v3416
		goto L7
	} else {
		goto L1076
	}
L1070:
	;
	if l3 != 0 {
		goto L1071
	} else {
		goto L1072
	}
L1071:
	;
	v3435 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v3435 <= int32(0) {
		v3556 = v3433
		goto L7
	} else {
		goto L1074
	}
L1072:
	;
	goto L1073
L1073:
	;
	F_pfree(m, v3429)
	mBase = m.M
	v3439 = m.ExcPending
	if v3439 != 0 {
		goto L1
	} else {
		goto L1075
	}
L1074:
	;
	goto L1073
L1075:
	;
	v3556 = v3433
	goto L7
L1076:
	;
	v3441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3441 == int32(1) {
		goto L1057
	} else {
		goto L1077
	}
L1077:
	;
	v3556 = int32(2)
	goto L7
L1078:
	;
	v3454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3454 != 0 {
		goto L1083
	} else {
		goto L1084
	}
L1079:
	;
	v3447 = F_JsonbType(m, l2)
	mBase = m.M
	v3448 = m.ExcPending
	if v3448 != 0 {
		goto L1
	} else {
		goto L1080
	}
L1080:
	;
	if v3447 != int32(16) {
		goto L1078
	} else {
		goto L1081
	}
L1081:
	;
	v3452 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v3453 = m.ExcPending
	if v3453 != 0 {
		goto L1
	} else {
		goto L1082
	}
L1082:
	;
	v3556 = v3452
	goto L7
L1083:
	;
	v3556 = int32(1)
	goto L7
L1084:
	;
	goto L1085
L1085:
	;
	v3456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3456 != int32(1) {
		goto L1086
	} else {
		goto L1087
	}
L1086:
	;
	v3556 = int32(2)
	goto L7
L1087:
	;
	goto L1088
L1088:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3463 = m.ExcPending
	if v3463 != 0 {
		goto L1
	} else {
		goto L1089
	}
L1089:
	;
	F_errcode(m, int32(285999234))
	mBase = m.M
	v3466 = m.ExcPending
	if v3466 != 0 {
		goto L1
	} else {
		goto L1090
	}
L1090:
	;
	F_errmsg(m, int32(110886), int32(0))
	mBase = m.M
	v3470 = m.ExcPending
	if v3470 != 0 {
		goto L1
	} else {
		goto L1091
	}
L1091:
	;
	F_errfinish(m, int32(500183), int32(1054), int32(108435))
	mBase = m.M
	v3475 = m.ExcPending
	if v3475 != 0 {
		goto L1
	} else {
		goto L1092
	}
L1092:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1093:
	;
	F_errcode(m, int32(285999234))
	mBase = m.M
	v3482 = m.ExcPending
	if v3482 != 0 {
		goto L1
	} else {
		goto L1094
	}
L1094:
	;
	v3483 = *(*int32)(unsafe.Add(mBase, uint32(v24)+760))
	v3484 = *(*int32)(unsafe.Add(mBase, uint32(v24)+756))
	v3485 = F_pnstrdup(m, v3483, v3484)
	mBase = m.M
	v3486 = m.ExcPending
	if v3486 != 0 {
		goto L1
	} else {
		goto L1095
	}
L1095:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v3485
	F_errmsg(m, int32(692889), v24+int32(48))
	mBase = m.M
	v3492 = m.ExcPending
	if v3492 != 0 {
		goto L1
	} else {
		goto L1096
	}
L1096:
	;
	F_errfinish(m, int32(500183), int32(1044), int32(108435))
	mBase = m.M
	v3497 = m.ExcPending
	if v3497 != 0 {
		goto L1
	} else {
		goto L1097
	}
L1097:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1098:
	;
	v3503 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v3503 != 0 {
		v3514 = int32(1)
		goto L1099
	} else {
		goto L1100
	}
L1099:
	;
	v3516 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v3516 != int32(18) {
		v3556 = v3514
		goto L7
	} else {
		goto L1104
	}
L1100:
	;
	v3504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	v3505 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v3505)
	v3510 = F_executeNextItem(m, l0, l1, v24+int32(752), l2, l3, v3505)
	mBase = m.M
	v3511 = m.ExcPending
	if v3511 != 0 {
		goto L1
	} else {
		goto L1101
	}
L1101:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v3504)
	if l3 != 0 {
		v3514 = v3510
		goto L1099
	} else {
		goto L1102
	}
L1102:
	;
	if v3510 != 0 {
		v3514 = v3510
		goto L1099
	} else {
		goto L1103
	}
L1103:
	;
	v3556 = int32(0)
	goto L7
L1104:
	;
	if v3500 != 0 {
		goto L1105
	} else {
		goto L1106
	}
L1105:
	;
	v3522 = v24 + int32(752)
	goto L1107
L1106:
	;
	v3522 = int32(0)
	goto L1107
L1107:
	;
	v3523 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3524 = int32(1)
	v3525 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v3529 = F_executeAnyItem(m, l0, v3522, v3523, l3, v3524, v3525, v3526, v3524, v3528)
	mBase = m.M
	v3530 = m.ExcPending
	if v3530 != 0 {
		goto L1
	} else {
		goto L1108
	}
L1108:
	;
	v3556 = v3529
	goto L7
L1109:
	;
	v3532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3532 != int32(1) {
		goto L1110
	} else {
		goto L1111
	}
L1110:
	;
	v3556 = int32(2)
	goto L7
L1111:
	;
	goto L1112
L1112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3539 = m.ExcPending
	if v3539 != 0 {
		goto L1
	} else {
		goto L1113
	}
L1113:
	;
	F_errcode(m, int32(151781506))
	mBase = m.M
	v3542 = m.ExcPending
	if v3542 != 0 {
		goto L1
	} else {
		goto L1114
	}
L1114:
	;
	F_errmsg(m, int32(25305), int32(0))
	mBase = m.M
	v3546 = m.ExcPending
	if v3546 != 0 {
		goto L1
	} else {
		goto L1115
	}
L1115:
	;
	F_errfinish(m, int32(500183), int32(978), int32(108435))
	mBase = m.M
	v3551 = m.ExcPending
	if v3551 != 0 {
		goto L1
	} else {
		goto L1116
	}
L1116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
