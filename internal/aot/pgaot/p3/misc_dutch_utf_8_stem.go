package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dutch_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
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
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v925 int32
	_ = v925
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
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
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1381 int32
	_ = v1381
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1446 int32
	_ = v1446
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1479 int32
	_ = v1479
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1573 int32
	_ = v1573
	var v1576 int32
	_ = v1576
	var v1580 int32
	_ = v1580
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1606 int32
	_ = v1606
	var v1609 int32
	_ = v1609
	var v1613 int32
	_ = v1613
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1646 int32
	_ = v1646
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1655 int32
	_ = v1655
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1708 int32
	_ = v1708
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
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
	var v1770 int32
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
	var v1819 int32
	_ = v1819
	var v1826 int32
	_ = v1826
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1867 int32
	_ = v1867
	var v1869 int32
	_ = v1869
	var v1877 int32
	_ = v1877
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1889 int32
	_ = v1889
	var v1906 int32
	_ = v1906
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1936 int32
	_ = v1936
	var v1942 int32
	_ = v1942
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1952 int32
	_ = v1952
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1992 int32
	_ = v1992
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2014 int32
	_ = v2014
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2022 int32
	_ = v2022
	var v2026 int32
	_ = v2026
	var v2028 int32
	_ = v2028
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2083 int32
	_ = v2083
	var v2087 int32
	_ = v2087
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2103 int32
	_ = v2103
	var v2116 int32
	_ = v2116
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2136 int32
	_ = v2136
	var v2142 int32
	_ = v2142
	var v2154 int32
	_ = v2154
	var v2161 int32
	_ = v2161
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2180 int32
	_ = v2180
	var v2182 int32
	_ = v2182
	var v2189 int32
	_ = v2189
	var v2192 int32
	_ = v2192
	var v2196 int32
	_ = v2196
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2218 int32
	_ = v2218
	var v2222 int32
	_ = v2222
	var v2224 int32
	_ = v2224
	var v2226 int32
	_ = v2226
	var v2230 int32
	_ = v2230
	var v2232 int32
	_ = v2232
	var v2236 int32
	_ = v2236
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2268 int32
	_ = v2268
	var v2270 int32
	_ = v2270
	var v2274 int32
	_ = v2274
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2304 int32
	_ = v2304
	var v2306 int32
	_ = v2306
	var v2310 int32
	_ = v2310
	var v2313 int32
	_ = v2313
	var v2315 int32
	_ = v2315
	var v2319 int32
	_ = v2319
	var v2329 int32
	_ = v2329
	var v2331 int32
	_ = v2331
	var v2335 int32
	_ = v2335
	var v2348 int32
	_ = v2348
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2368 int32
	_ = v2368
	var v2374 int32
	_ = v2374
	var v2386 int32
	_ = v2386
	var v2393 int32
	_ = v2393
	var v2398 int32
	_ = v2398
	var v2402 int32
	_ = v2402
	var v2404 int32
	_ = v2404
	var v2406 int32
	_ = v2406
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2426 int32
	_ = v2426
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2437 int32
	_ = v2437
	var v2439 int32
	_ = v2439
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2449 int32
	_ = v2449
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2475 int32
	_ = v2475
	var v2489 int32
	_ = v2489
	var v2496 int32
	_ = v2496
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2519 int32
	_ = v2519
	var v2527 int32
	_ = v2527
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2553 int32
	_ = v2553
	var v2556 int32
	_ = v2556
	var v2559 int32
	_ = v2559
	var v2567 int32
	_ = v2567
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2586 int32
	_ = v2586
	var v2588 int32
	_ = v2588
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
	var v2600 int32
	_ = v2600
	var v2601 int32
	_ = v2601
	var v2605 int32
	_ = v2605
	var v2607 int32
	_ = v2607
	var v2614 int32
	_ = v2614
	var v2616 int32
	_ = v2616
	var v2621 int32
	_ = v2621
	var v2623 int32
	_ = v2623
	var v2630 int32
	_ = v2630
	var v2633 int32
	_ = v2633
	var v2637 int32
	_ = v2637
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2659 int32
	_ = v2659
	var v2663 int32
	_ = v2663
	var v2674 int32
	_ = v2674
	var v2676 int32
	_ = v2676
	var v2678 int32
	_ = v2678
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2699 int32
	_ = v2699
	var v2701 int32
	_ = v2701
	var v2706 int32
	_ = v2706
	var v2708 int32
	_ = v2708
	var v2715 int32
	_ = v2715
	var v2718 int32
	_ = v2718
	var v2722 int32
	_ = v2722
	var v2729 int32
	_ = v2729
	var v2730 int32
	_ = v2730
	var v2744 int32
	_ = v2744
	var v2748 int32
	_ = v2748
	var v2750 int32
	_ = v2750
	var v2752 int32
	_ = v2752
	var v2756 int32
	_ = v2756
	var v2758 int32
	_ = v2758
	var v2762 int32
	_ = v2762
	var v2764 int32
	_ = v2764
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2802 int32
	_ = v2802
	var v2804 int32
	_ = v2804
	var v2808 int32
	_ = v2808
	var v2811 int32
	_ = v2811
	var v2813 int32
	_ = v2813
	var v2817 int32
	_ = v2817
	var v2827 int32
	_ = v2827
	var v2829 int32
	_ = v2829
	var v2833 int32
	_ = v2833
	var v2846 int32
	_ = v2846
	var v2861 int32
	_ = v2861
	var v2862 int32
	_ = v2862
	var v2866 int32
	_ = v2866
	var v2872 int32
	_ = v2872
	var v2884 int32
	_ = v2884
	var v2891 int32
	_ = v2891
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2903 int32
	_ = v2903
	var v2905 int32
	_ = v2905
	var v2910 int32
	_ = v2910
	var v2912 int32
	_ = v2912
	var v2919 int32
	_ = v2919
	var v2922 int32
	_ = v2922
	var v2926 int32
	_ = v2926
	var v2933 int32
	_ = v2933
	var v2934 int32
	_ = v2934
	var v2948 int32
	_ = v2948
	var v2952 int32
	_ = v2952
	var v2954 int32
	_ = v2954
	var v2956 int32
	_ = v2956
	var v2960 int32
	_ = v2960
	var v2962 int32
	_ = v2962
	var v2966 int32
	_ = v2966
	var v2989 int32
	_ = v2989
	var v2990 int32
	_ = v2990
	var v2992 int32
	_ = v2992
	var v2994 int32
	_ = v2994
	var v2998 int32
	_ = v2998
	var v3000 int32
	_ = v3000
	var v3004 int32
	_ = v3004
	var v3018 int32
	_ = v3018
	var v3019 int32
	_ = v3019
	var v3034 int32
	_ = v3034
	var v3036 int32
	_ = v3036
	var v3040 int32
	_ = v3040
	var v3043 int32
	_ = v3043
	var v3045 int32
	_ = v3045
	var v3049 int32
	_ = v3049
	var v3059 int32
	_ = v3059
	var v3061 int32
	_ = v3061
	var v3065 int32
	_ = v3065
	var v3078 int32
	_ = v3078
	var v3093 int32
	_ = v3093
	var v3094 int32
	_ = v3094
	var v3098 int32
	_ = v3098
	var v3104 int32
	_ = v3104
	var v3116 int32
	_ = v3116
	var v3123 int32
	_ = v3123
	var v3127 int32
	_ = v3127
	var v3129 int32
	_ = v3129
	var v3132 int32
	_ = v3132
	var v3135 int32
	_ = v3135
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3141 int32
	_ = v3141
	var v3143 int32
	_ = v3143
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3153 int32
	_ = v3153
	var v3159 int32
	_ = v3159
	var v3160 int32
	_ = v3160
	var v3165 int32
	_ = v3165
	var v3166 int32
	_ = v3166
	var v3172 int32
	_ = v3172
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3181 int32
	_ = v3181
	var v3183 int32
	_ = v3183
	var v3188 int32
	_ = v3188
	var v3190 int32
	_ = v3190
	var v3197 int32
	_ = v3197
	var v3200 int32
	_ = v3200
	var v3204 int32
	_ = v3204
	var v3211 int32
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3226 int32
	_ = v3226
	var v3233 int32
	_ = v3233
	var v3245 int32
	_ = v3245
	var v3256 int32
	_ = v3256
	var v3257 int32
	_ = v3257
	var v3262 int32
	_ = v3262
	var v3265 int32
	_ = v3265
	var v3268 int32
	_ = v3268
	var v3276 int32
	_ = v3276
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3296 int32
	_ = v3296
	var v3297 int32
	_ = v3297
	var v3302 int32
	_ = v3302
	var v3305 int32
	_ = v3305
	var v3308 int32
	_ = v3308
	var v3316 int32
	_ = v3316
	var v3326 int32
	_ = v3326
	var v3327 int32
	_ = v3327
	var v3335 int32
	_ = v3335
	var v3337 int32
	_ = v3337
	var v3338 int32
	_ = v3338
	var v3339 int32
	_ = v3339
	var v3341 int32
	_ = v3341
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
	var v3352 int32
	_ = v3352
	var v3356 int32
	_ = v3356
	var v3357 int32
	_ = v3357
	var v3359 int32
	_ = v3359
	var v3361 int32
	_ = v3361
	var v3367 int32
	_ = v3367
	var v3368 int32
	_ = v3368
	var v3371 int32
	_ = v3371
	var v3377 int32
	_ = v3377
	var v3378 int32
	_ = v3378
	var v3383 int32
	_ = v3383
	var v3384 int32
	_ = v3384
	var v3389 int32
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3396 int32
	_ = v3396
	var v3398 int32
	_ = v3398
	var v3399 int32
	_ = v3399
	var v3400 int32
	_ = v3400
	var v3404 int32
	_ = v3404
	var v3406 int32
	_ = v3406
	var v3407 int32
	_ = v3407
	var v3409 int32
	_ = v3409
	var v3412 int32
	_ = v3412
	var v3413 int32
	_ = v3413
	var v3415 int32
	_ = v3415
	var v3417 int32
	_ = v3417
	var v3419 int32
	_ = v3419
	var v3421 int32
	_ = v3421
	var v3436 int32
	_ = v3436
	var v3437 int32
	_ = v3437
	var v3440 int32
	_ = v3440
	var v3446 int32
	_ = v3446
	var v3447 int32
	_ = v3447
	var v3452 int32
	_ = v3452
	var v3453 int32
	_ = v3453
	var v3458 int32
	_ = v3458
	var v3459 int32
	_ = v3459
	var v3464 int32
	_ = v3464
	var v3465 int32
	_ = v3465
	var v3470 int32
	_ = v3470
	var v3471 int32
	_ = v3471
	var v3476 int32
	_ = v3476
	var v3477 int32
	_ = v3477
	var v3482 int32
	_ = v3482
	var v3483 int32
	_ = v3483
	var v3488 int32
	_ = v3488
	var v3489 int32
	_ = v3489
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3500 int32
	_ = v3500
	var v3501 int32
	_ = v3501
	var v3504 int32
	_ = v3504
	var v3506 int32
	_ = v3506
	var v3510 int32
	_ = v3510
	var v3514 int32
	_ = v3514
	var v3521 int32
	_ = v3521
	var v3522 int32
	_ = v3522
	var v3527 int32
	_ = v3527
	var v3528 int32
	_ = v3528
	var v3533 int32
	_ = v3533
	var v3534 int32
	_ = v3534
	var v3539 int32
	_ = v3539
	var v3540 int32
	_ = v3540
	var v3545 int32
	_ = v3545
	var v3546 int32
	_ = v3546
	var v3551 int32
	_ = v3551
	var v3552 int32
	_ = v3552
	var v3557 int32
	_ = v3557
	var v3558 int32
	_ = v3558
	var v3563 int32
	_ = v3563
	var v3564 int32
	_ = v3564
	var v3569 int32
	_ = v3569
	var v3570 int32
	_ = v3570
	var v3575 int32
	_ = v3575
	var v3576 int32
	_ = v3576
	var v3584 int32
	_ = v3584
	var v3587 int32
	_ = v3587
	var v3588 int32
	_ = v3588
	var v3591 int32
	_ = v3591
	var v3592 int32
	_ = v3592
	var v3598 int32
	_ = v3598
	var v3602 int32
	_ = v3602
	v2 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v14
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L2
L1:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
	v108 = int32(1)
	if v105 <= v103 {
		v595 = v2
		v596 = v108
		goto L28
	} else {
		goto L29
	}
L2:
	;
	v25 = int32(0)
	v26 = F_out_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), v25)
	mBase = m.M
	if v26 == v25 {
		goto L2
	} else {
		goto L4
	}
L3:
	;
	v31 = int32(1)
	goto L5
L4:
	;
	goto L3
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = F_eq_s(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_1))
	mBase = m.M
	if v37 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34
	if int32(0) < v31 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	goto L6
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34
	v45 = F_in_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v45 != 0 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v31 = v31 - int32(1)
	goto L5
L11:
	;
	goto L10
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17
	goto L1
L13:
	;
	v55 = F_out_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v55 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v56
	goto L15
L15:
	;
	v65 = int32(0)
	v66 = F_out_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), v65)
	mBase = m.M
	if v66 == v65 {
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v71 = int32(1)
	goto L18
L17:
	;
	goto L16
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v77 = F_eq_s(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_2))
	mBase = m.M
	if v77 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v74
	if int32(0) < v71 {
		goto L12
	} else {
		goto L25
	}
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v74
	v85 = F_in_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v85 != 0 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v71 = v71 - int32(1)
	goto L18
L24:
	;
	goto L23
L25:
	;
	v95 = F_out_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v95 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v96
	goto L12
L27:
	;
	return v3602
L28:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v597
	v601 = v597 - int32(1)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v601 <= v602 {
		goto L163
	} else {
		goto L164
	}
L29:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v112 = int32(1)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v105-v112))))
	if base.B2i32(v114&int32(224) != int32(96))|base.B2i32(v112<<(uint(v114)%32)&int32(_a_F_dutch_UTF_8_stem_3) == int32(0)) != 0 {
		v595 = v2
		v596 = v108
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v129 = F_find_among_b(m, l0, int32(_a_F_dutch_UTF_8_stem_4), int32(8), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	if v129 == int32(0) {
		v595 = v2
		v596 = v108
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v135
	v137 = int32(1)
	switch v129 - v137 {
	case 0:
		goto L43
	case 1:
		goto L42
	case 2:
		goto L41
	case 3:
		goto L40
	case 4:
		goto L39
	case 5:
		goto L38
	case 6:
		goto L37
	case 7:
		goto L36
	default:
		v595 = v137
		v596 = v108
		goto L28
	}
L34:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v587
	v590 = F_slice_del(m, l0)
	mBase = m.M
	if v590 < int32(0) {
		v3602 = v590
		goto L27
	} else {
		goto L161
	}
L35:
	;
	v595 = int32(0)
	v596 = v585
	goto L28
L36:
	;
	v578 = F_slice_from_s(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_5))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L31
	} else {
		goto L159
	}
L37:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v388 = int32(3)
	v390 = int32(0)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v392-v393 < v388 {
		v403 = v390
		goto L110
	} else {
		goto L111
	}
L38:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v135 < v354 {
		v585 = v108
		goto L35
	} else {
		goto L98
	}
L39:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v135 < v346 {
		v585 = v108
		goto L35
	} else {
		goto L95
	}
L40:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v196 = v195 - v135
	v197 = int32(2)
	v199 = int32(0)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v201-v202 < v197 {
		v212 = v199
		goto L61
	} else {
		goto L62
	}
L41:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v135 < v187 {
		v585 = v108
		goto L35
	} else {
		goto L56
	}
L42:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v135 < v143 {
		v585 = v108
		goto L35
	} else {
		goto L45
	}
L43:
	;
	v140 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v140 {
		v595 = v137
		v596 = v108
		goto L28
	} else {
		goto L44
	}
L44:
	;
	v3602 = v140
	goto L27
L45:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v135 <= v145 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v135
	v159 = int32(0)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v166 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v166 != 0 {
		v181 = v159
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v135-int32(1)))))
	if v151 != int32(116) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v135 - int32(1)
	if v143 < v135 {
		v585 = v108
		goto L35
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	if v181 == int32(0) {
		v585 = v108
		goto L35
	} else {
		goto L54
	}
L51:
	;
	goto L50
L52:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v168 = v163 - v135
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v167 - v168
	v175 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v175 != 0 {
		v181 = v159
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176 - v168
	v181 = int32(1)
	goto L51
L54:
	;
	v184 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v184 {
		v595 = v137
		v596 = v108
		goto L28
	} else {
		goto L55
	}
L55:
	;
	v3602 = v184
	goto L27
L56:
	;
	v191 = F_slice_from_s(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_7))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L31
	} else {
		goto L57
	}
L57:
	;
	if int32(0) <= v191 {
		v595 = v137
		v596 = v108
		goto L28
	} else {
		goto L58
	}
L58:
	;
	v3602 = v191
	goto L27
L59:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v254 = v253 - v196
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v254
	v256 = int32(2)
	v258 = int32(0)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v254-v261 < v256 {
		v271 = v258
		goto L76
	} else {
		goto L77
	}
L60:
	;
	if v212 == int32(0) {
		goto L59
	} else {
		goto L64
	}
L61:
	;
	goto L60
L62:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v208 = F_memcmp(m, v205+v201-v197, int32(_a_F_dutch_UTF_8_stem_8), v197)
	mBase = m.M
	if v208 != 0 {
		v212 = v199
		goto L61
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v201 - v197
	v212 = int32(1)
	goto L61
L64:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v216 < v215 {
		goto L59
	} else {
		goto L65
	}
L65:
	;
	v218 = int32(0)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v225 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v225 != 0 {
		v240 = v218
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v240 == int32(0) {
		goto L59
	} else {
		goto L70
	}
L67:
	;
	goto L66
L68:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v227 = v222 - v221
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v226 - v227
	v234 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v234 != 0 {
		v240 = v218
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v235 - v227
	v240 = int32(1)
	goto L67
L70:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v243 - v196
	v246 = F_slice_del(m, l0)
	mBase = m.M
	if v246 < int32(0) {
		v3602 = v246
		goto L27
	} else {
		goto L71
	}
L71:
	;
	v249 = F_r_lengthen_V_2(m, l0)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L31
	} else {
		goto L72
	}
L72:
	;
	if int32(0) <= v249 {
		v595 = v137
		v596 = v108
		goto L28
	} else {
		goto L73
	}
L73:
	;
	v3602 = v249
	goto L27
L74:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v309 = v308 - v196
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v309
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v309 < v311 {
		v585 = v108
		goto L35
	} else {
		goto L87
	}
L75:
	;
	if v271 == int32(0) {
		goto L74
	} else {
		goto L79
	}
L76:
	;
	goto L75
L77:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v267 = F_memcmp(m, v264+v254-v256, int32(_a_F_dutch_UTF_8_stem_9), v256)
	mBase = m.M
	if v267 != 0 {
		v271 = v258
		goto L76
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v254 - v256
	v271 = int32(1)
	goto L76
L79:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v275 < v274 {
		goto L74
	} else {
		goto L80
	}
L80:
	;
	v277 = int32(0)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v284 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v284 != 0 {
		v299 = v277
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v299 == int32(0) {
		goto L74
	} else {
		goto L85
	}
L82:
	;
	goto L81
L83:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v286 = v281 - v280
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v285 - v286
	v293 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v293 != 0 {
		v299 = v277
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v294 - v286
	v299 = int32(1)
	goto L82
L85:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v302 - v196
	v305 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v305 {
		v595 = v137
		v596 = v108
		goto L28
	} else {
		goto L86
	}
L86:
	;
	v3602 = v305
	goto L27
L87:
	;
	v313 = int32(0)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v321 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v321 != 0 {
		v336 = v313
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v336 == int32(0) {
		v595 = v313
		v596 = v108
		goto L28
	} else {
		goto L92
	}
L89:
	;
	goto L88
L90:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v323 = v318 - v317
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v322 - v323
	v330 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v330 != 0 {
		v336 = v313
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331 - v323
	v336 = int32(1)
	goto L89
L92:
	;
	v339 = int32(1)
	v342 = F_slice_from_s(m, l0, v339, int32(_a_F_dutch_UTF_8_stem_10))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L31
	} else {
		goto L93
	}
L93:
	;
	if int32(0) <= v342 {
		v595 = v339
		v596 = v108
		goto L28
	} else {
		goto L94
	}
L94:
	;
	v3602 = v342
	goto L27
L95:
	;
	v350 = F_slice_from_s(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_11))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L31
	} else {
		goto L96
	}
L96:
	;
	if int32(0) <= v350 {
		v595 = v137
		v596 = v108
		goto L28
	} else {
		goto L97
	}
L97:
	;
	v3602 = v350
	goto L27
L98:
	;
	v356 = int32(0)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v360 = v358 - v359
	v365 = F_in_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), v356)
	mBase = m.M
	if v365 != 0 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	if v378 == int32(0) {
		v585 = v108
		goto L35
	} else {
		goto L105
	}
L100:
	;
	goto L99
L101:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v366 - v360
	v371 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_12))
	mBase = m.M
	if v371 == int32(0) {
		v378 = v356
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v374 - v360
	v378 = int32(1)
	goto L100
L104:
	;
	goto L103
L105:
	;
	v383 = F_slice_from_s(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_13))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L31
	} else {
		goto L106
	}
L106:
	;
	if int32(0) <= v383 {
		v595 = v137
		v596 = v108
		goto L28
	} else {
		goto L107
	}
L107:
	;
	v3602 = v383
	goto L27
L108:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v419 = v387 - v135
	v420 = v418 - v419
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v420
	v422 = int32(2)
	v424 = int32(0)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v420-v427 < v422 {
		v437 = v424
		goto L118
	} else {
		goto L119
	}
L109:
	;
	if v403 == int32(0) {
		goto L108
	} else {
		goto L113
	}
L110:
	;
	goto L109
L111:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v399 = F_memcmp(m, v396+v392-v388, int32(_a_F_dutch_UTF_8_stem_14), v388)
	mBase = m.M
	if v399 != 0 {
		v403 = v390
		goto L110
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v392 - v388
	v403 = int32(1)
	goto L110
L113:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v406 < v407 {
		goto L108
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v406
	v413 = F_slice_from_s(m, l0, int32(4), int32(_a_F_dutch_UTF_8_stem_15))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L31
	} else {
		goto L115
	}
L115:
	;
	if int32(0) <= v413 {
		v595 = int32(1)
		v596 = v108
		goto L28
	} else {
		goto L116
	}
L116:
	;
	v3602 = v413
	goto L27
L117:
	;
	if v437 != 0 {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	goto L117
L119:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v433 = F_memcmp(m, v430+v420-v422, int32(_a_F_dutch_UTF_8_stem_16), v422)
	mBase = m.M
	if v433 != 0 {
		v437 = v424
		goto L118
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v420 - v422
	v437 = int32(1)
	goto L118
L121:
	;
	v439 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v439 {
		v595 = int32(1)
		v596 = v108
		goto L28
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v443 = v442 - v419
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v443
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v443 <= v445 {
		v485 = v443
		v486 = v445
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v3602 = v439
	goto L27
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v485
	if v485 <= v486 {
		v528 = v485
		goto L135
	} else {
		goto L136
	}
L126:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447+v443-int32(1)))))
	if v451 != int32(100) {
		v485 = v443
		v486 = v445
		goto L125
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v443 - int32(1)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v443 <= v457 {
		v485 = v443
		v486 = v445
		goto L125
	} else {
		goto L128
	}
L128:
	;
	v459 = int32(0)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v466 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v466 != 0 {
		v481 = v459
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v481 != 0 {
		goto L34
	} else {
		goto L133
	}
L130:
	;
	goto L129
L131:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v468 = v463 - v462
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v467 - v468
	v475 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v475 != 0 {
		v481 = v459
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v476 - v468
	v481 = int32(1)
	goto L130
L133:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v485 = v482 - v419
	v486 = v484
	goto L125
L134:
	;
	v572 = F_slice_del(m, l0)
	mBase = m.M
	if v572 < int32(0) {
		v3602 = v572
		goto L27
	} else {
		goto L158
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v528
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v528 < v530 {
		v585 = v108
		goto L35
	} else {
		goto L145
	}
L136:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v491 = int32(1)
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489+v485-v491))))
	if base.Ui32(v491) < base.Ui32((v493-int32(105))&int32(255)) {
		v528 = v485
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v501 = v485 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v501
	v503 = int32(0)
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v507 = v505 - v501
	v512 = F_in_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), v503)
	mBase = m.M
	if v512 != 0 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	if v525 != 0 {
		goto L134
	} else {
		goto L144
	}
L139:
	;
	goto L138
L140:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v513 - v507
	v518 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_12))
	mBase = m.M
	if v518 == int32(0) {
		v525 = v503
		goto L139
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v521 - v507
	v525 = int32(1)
	goto L139
L143:
	;
	goto L142
L144:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v528 = v526 - v419
	goto L135
L145:
	;
	v532 = int32(0)
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v539 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v539 != 0 {
		v554 = v532
		goto L147
	} else {
		goto L148
	}
L146:
	;
	if v554 == int32(0) {
		v585 = v108
		goto L35
	} else {
		goto L150
	}
L147:
	;
	goto L146
L148:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v541 = v536 - v535
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v540 - v541
	v548 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v548 != 0 {
		v554 = v532
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v549 - v541
	v554 = int32(1)
	goto L147
L150:
	;
	v557 = F_slice_del(m, l0)
	mBase = m.M
	if v557 < int32(0) {
		v3602 = v557
		goto L27
	} else {
		goto L151
	}
L151:
	;
	v561 = F_r_lengthen_V_2(m, l0)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L31
	} else {
		goto L152
	}
L152:
	;
	if int32(0) < v561 {
		v595 = int32(1)
		v596 = v108
		goto L28
	} else {
		goto L153
	}
L153:
	;
	v565 = int32(1)
	if base.Ui32(v561) <= base.Ui32(v565) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v568 = v565
	goto L156
L155:
	;
	v568 = v561
	goto L156
L156:
	;
	if v561 == int32(0) {
		v585 = v568
		goto L35
	} else {
		goto L157
	}
L157:
	;
	return v568
L158:
	;
	v595 = int32(1)
	v596 = v108
	goto L28
L159:
	;
	if int32(0) <= v578 {
		v595 = v137
		v596 = v108
		goto L28
	} else {
		goto L160
	}
L160:
	;
	v3602 = v578
	goto L27
L161:
	;
	v595 = int32(1)
	v596 = v108
	goto L28
L162:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1259
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1259
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1259-int32(2) <= v1262 {
		goto L394
	} else {
		goto L395
	}
L163:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L164:
	;
	goto L165
L165:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604+v601))))
	if v606 != int32(101) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L167:
	;
	goto L168
L168:
	;
	v612 = F_find_among_b(m, l0, int32(_a_F_dutch_UTF_8_stem_17), int32(11), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L31
	} else {
		goto L169
	}
L169:
	;
	if v612 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L171:
	;
	goto L172
L172:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v616
	v618 = int32(1)
	switch v612 - v618 {
	case 0:
		goto L186
	case 1:
		goto L185
	case 2:
		goto L184
	case 3:
		goto L183
	case 4:
		goto L182
	case 5:
		goto L181
	case 6:
		goto L180
	case 7:
		goto L179
	case 8:
		goto L178
	case 9:
		goto L177
	case 10:
		goto L175
	default:
		v1252 = v618
		v1254 = v596
		goto L162
	}
L173:
	;
	v1248 = F_slice_del(m, l0)
	mBase = m.M
	if v1248 < int32(0) {
		v3602 = v1248
		goto L27
	} else {
		goto L392
	}
L174:
	;
	if v1210 < int32(0) {
		v3602 = v1214
		goto L27
	} else {
		goto L391
	}
L175:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v616 < v1215 {
		goto L379
	} else {
		goto L380
	}
L176:
	;
	v1212 = base.B2i32(v1210 < int32(0))
	if v1210 < int32(0) {
		goto L372
	} else {
		goto L373
	}
L177:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v616 < v1169 {
		goto L357
	} else {
		goto L358
	}
L178:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v616 < v1153 {
		goto L349
	} else {
		goto L350
	}
L179:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v616 < v1145 {
		goto L344
	} else {
		goto L345
	}
L180:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v616 < v1137 {
		goto L339
	} else {
		goto L340
	}
L181:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v616 < v1129 {
		goto L334
	} else {
		goto L335
	}
L182:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v616 < v1099 {
		goto L323
	} else {
		goto L324
	}
L183:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v616 < v1091 {
		goto L318
	} else {
		goto L319
	}
L184:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v616 < v1083 {
		goto L313
	} else {
		goto L314
	}
L185:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v616 < v1075 {
		goto L308
	} else {
		goto L309
	}
L186:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v622 = int32(2)
	v624 = int32(0)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v626-v627 < v622 {
		v637 = v624
		goto L188
	} else {
		goto L189
	}
L187:
	;
	if v637 != 0 {
		goto L191
	} else {
		goto L192
	}
L188:
	;
	goto L187
L189:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v633 = F_memcmp(m, v630+v626-v622, int32(_a_F_dutch_UTF_8_stem_18), v622)
	mBase = m.M
	if v633 != 0 {
		v637 = v624
		goto L188
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v626 - v622
	v637 = int32(1)
	goto L188
L191:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v638
	v640 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v640 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v644 = v621 - v616
	v645 = v643 - v644
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v645
	v647 = int32(2)
	v649 = int32(0)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v645-v652 < v647 {
		v662 = v649
		goto L197
	} else {
		goto L198
	}
L194:
	;
	v3602 = v640
	goto L27
L195:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v699 = v698 - v644
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v699
	v701 = int32(3)
	v703 = int32(0)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v699-v706 < v701 {
		v716 = v703
		goto L209
	} else {
		goto L210
	}
L196:
	;
	if v662 == int32(0) {
		goto L195
	} else {
		goto L200
	}
L197:
	;
	goto L196
L198:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v658 = F_memcmp(m, v655+v645-v647, int32(_a_F_dutch_UTF_8_stem_19), v647)
	mBase = m.M
	if v658 != 0 {
		v662 = v649
		goto L197
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v645 - v647
	v662 = int32(1)
	goto L197
L200:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v665
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v665 < v667 {
		goto L195
	} else {
		goto L201
	}
L201:
	;
	v669 = int32(0)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v676 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v676 != 0 {
		v691 = v669
		goto L203
	} else {
		goto L204
	}
L202:
	;
	if v691 == int32(0) {
		goto L195
	} else {
		goto L206
	}
L203:
	;
	goto L202
L204:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v678 = v673 - v672
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v677 - v678
	v685 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v685 != 0 {
		v691 = v669
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v686 - v678
	v691 = int32(1)
	goto L203
L206:
	;
	v694 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v694 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L207
	}
L207:
	;
	v3602 = v694
	goto L27
L208:
	;
	if v716 != 0 {
		goto L212
	} else {
		goto L213
	}
L209:
	;
	goto L208
L210:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v712 = F_memcmp(m, v709+v699-v701, int32(_a_F_dutch_UTF_8_stem_20), v701)
	mBase = m.M
	if v712 != 0 {
		v716 = v703
		goto L209
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v699 - v701
	v716 = int32(1)
	goto L209
L212:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v717
	v721 = F_slice_from_s(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_21))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L31
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v726 = v725 - v644
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v726
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v726 <= v728 {
		v964 = v726
		goto L217
	} else {
		goto L218
	}
L215:
	;
	if int32(0) <= v721 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L216
	}
L216:
	;
	v3602 = v721
	goto L27
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v964
	v970 = int32(3)
	v972 = int32(0)
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v964-v975 < v970 {
		v985 = v972
		goto L275
	} else {
		goto L276
	}
L218:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730+v726-int32(1)))))
	if v734 != int32(116) {
		v964 = v726
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v738 = v726 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v738
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v738
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v726 <= v741 {
		v964 = v726
		goto L217
	} else {
		goto L220
	}
L220:
	;
	v743 = int32(0)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L224
L221:
	;
	if v960 != 0 {
		goto L173
	} else {
		goto L273
	}
L222:
	;
	if v799 < int32(0) {
		v960 = v743
		goto L221
	} else {
		goto L241
	}
L224:
	;
	goto L225
L225:
	;
	goto L226
L226:
	;
	v754 = v746
	v756 = int32(1)
	goto L229
L228:
	;
	v799 = v781
	goto L222
L229:
	;
	if v754 <= v747 {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	goto L228
L231:
	;
	v799 = int32(-1)
	goto L222
L232:
	;
	goto L233
L233:
	;
	v761 = v754 - int32(1)
	v763 = int32(*(*int8)(unsafe.Add(mBase, uint32(v745+v761))))
	if base.B2i32(int32(0) <= v763)|base.B2i32(v761 <= v747) != 0 {
		v781 = v761
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v785 = int32(1)
	if v785 < v756 {
		v754 = v781
		v756 = v756 - v785
		goto L229
	} else {
		goto L240
	}
L235:
	;
	v769 = v761
	goto L236
L236:
	;
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v745+v769))))
	if base.Ui32(int32(191)) < base.Ui32(v774) {
		v781 = v769
		goto L234
	} else {
		goto L238
	}
L237:
	;
	v781 = v747
	goto L234
L238:
	;
	v778 = v769 - int32(1)
	if v747 < v778 {
		v769 = v778
		goto L236
	} else {
		goto L239
	}
L239:
	;
	goto L237
L240:
	;
	goto L230
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v799
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L244
L242:
	;
	if v932 != 0 {
		goto L265
	} else {
		goto L266
	}
L243:
	;
	v932 = v925
	goto L242
L244:
	;
	if v799 <= v817 {
		v925 = int32(-1)
		goto L243
	} else {
		goto L246
	}
L245:
	;
	v925 = int32(0)
	goto L243
L246:
	;
	v834 = int32(1)
	v835 = v799 - v834
	v837 = int32(*(*int8)(unsafe.Add(mBase, uint32(v818+v835))))
	v839 = v837 & int32(255)
	if base.B2i32(v835 == v817)|base.B2i32(int32(0) <= v837) != 0 {
		v897 = v839
		v901 = v834
		goto L247
	} else {
		goto L248
	}
L247:
	;
	if int32(252) < v897 {
		goto L255
	} else {
		goto L256
	}
L248:
	;
	v846 = v839 & int32(63)
	v848 = v799 - int32(2)
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818+v848))))
	v852 = v850 << (uint(int32(6)) % 32)
	if base.B2i32(v848 != v817)&base.B2i32(base.Ui32(v850) < base.Ui32(int32(192))) == int32(0) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v897 = v852&int32(1984) | v846
	v901 = int32(2)
	goto L247
L250:
	;
	goto L251
L251:
	;
	v865 = v852&int32(4032) | v846
	v867 = v799 - int32(3)
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818+v867))))
	if base.B2i32(v867 != v817)&base.B2i32(base.Ui32(v869) < base.Ui32(int32(224))) == int32(0) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v897 = v869<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_22) | v865
	v901 = int32(3)
	goto L247
L253:
	;
	goto L254
L254:
	;
	v887 = int32(4)
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v799+v818-v887))))
	v897 = v869<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_23) | v889&int32(7)<<(uint(int32(18))%32) | v865
	v901 = v887
	goto L247
L255:
	;
	v932 = v901
	goto L242
L256:
	;
	goto L257
L257:
	;
	v903 = v897 - int32(97)
	if v903 < int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v932 = v901
	goto L242
L259:
	;
	goto L260
L260:
	;
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v903)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[0]))))
	if int32(base.Ui32(v909)>>(uint(v903&int32(7))%32))&int32(1) == int32(0) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v932 = v901
	goto L242
L262:
	;
	goto L263
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v799 - v901
	goto L264
L264:
	;
	goto L245
L265:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v935 = v933 + (v799 - v803)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v935
	v937 = int32(2)
	v939 = int32(0)
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v935-v942 < v937 {
		v952 = v939
		goto L269
	} else {
		goto L270
	}
L266:
	;
	goto L267
L267:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v955 + (v746 - v744)
	v960 = int32(1)
	goto L221
L268:
	;
	if v952 == int32(0) {
		v960 = v743
		goto L221
	} else {
		goto L272
	}
L269:
	;
	goto L268
L270:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v948 = F_memcmp(m, v945+v935-v937, int32(_a_F_dutch_UTF_8_stem_24), v937)
	mBase = m.M
	if v948 != 0 {
		v952 = v939
		goto L269
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v935 - v937
	v952 = int32(1)
	goto L269
L272:
	;
	goto L267
L273:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v964 = v962 - v644
	goto L217
L274:
	;
	if v985 != 0 {
		goto L278
	} else {
		goto L279
	}
L275:
	;
	goto L274
L276:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v981 = F_memcmp(m, v978+v964-v970, int32(_a_F_dutch_UTF_8_stem_25), v970)
	mBase = m.M
	if v981 != 0 {
		v985 = v972
		goto L275
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v964 - v970
	v985 = int32(1)
	goto L275
L278:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v986
	v990 = F_slice_from_s(m, l0, int32(3), int32(_a_F_dutch_UTF_8_stem_26))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L31
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v995 = v994 - v644
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v995
	v997 = int32(2)
	v999 = int32(0)
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v995-v1002 < v997 {
		v1012 = v999
		goto L284
	} else {
		goto L285
	}
L281:
	;
	if int32(0) <= v990 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L282
	}
L282:
	;
	v3602 = v990
	goto L27
L283:
	;
	if v1012 != 0 {
		goto L287
	} else {
		goto L288
	}
L284:
	;
	goto L283
L285:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1008 = F_memcmp(m, v1005+v995-v997, int32(_a_F_dutch_UTF_8_stem_27), v997)
	mBase = m.M
	if v1008 != 0 {
		v1012 = v999
		goto L284
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v995 - v997
	v1012 = int32(1)
	goto L284
L287:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1013
	v1017 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_28))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L31
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1022 = v1021 - v644
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1022
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1022 <= v1024 {
		goto L292
	} else {
		goto L293
	}
L290:
	;
	if int32(0) <= v1017 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L291
	}
L291:
	;
	v3602 = v1017
	goto L27
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1022
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1022 < v1045 {
		goto L297
	} else {
		goto L298
	}
L293:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1026+v1022-int32(1)))))
	if v1030 != int32(39) {
		goto L292
	} else {
		goto L294
	}
L294:
	;
	v1034 = v1022 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1034
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1034
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1022 <= v1037 {
		goto L292
	} else {
		goto L295
	}
L295:
	;
	v1039 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1039 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L296
	}
L296:
	;
	v3602 = v1039
	goto L27
L297:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L298:
	;
	goto L299
L299:
	;
	v1047 = int32(0)
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1054 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v1054 != 0 {
		v1069 = v1047
		goto L301
	} else {
		goto L302
	}
L300:
	;
	if v1069 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L301:
	;
	goto L300
L302:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1056 = v1051 - v1050
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1055 - v1056
	v1063 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v1063 != 0 {
		v1069 = v1047
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1064 - v1056
	v1069 = int32(1)
	goto L301
L304:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L305:
	;
	goto L306
L306:
	;
	v1072 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1072 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L307
	}
L307:
	;
	v3602 = v1072
	goto L27
L308:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L309:
	;
	goto L310
L310:
	;
	v1079 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_29))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L31
	} else {
		goto L311
	}
L311:
	;
	if int32(0) <= v1079 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L312
	}
L312:
	;
	v3602 = v1079
	goto L27
L313:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L314:
	;
	goto L315
L315:
	;
	v1087 = F_slice_from_s(m, l0, int32(4), int32(_a_F_dutch_UTF_8_stem_30))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L31
	} else {
		goto L316
	}
L316:
	;
	if int32(0) <= v1087 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L317
	}
L317:
	;
	v3602 = v1087
	goto L27
L318:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L319:
	;
	goto L320
L320:
	;
	v1095 = F_slice_from_s(m, l0, int32(4), int32(_a_F_dutch_UTF_8_stem_31))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L31
	} else {
		goto L321
	}
L321:
	;
	if int32(0) <= v1095 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L322
	}
L322:
	;
	v3602 = v1095
	goto L27
L323:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L324:
	;
	goto L325
L325:
	;
	v1101 = int32(0)
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1108 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v1108 != 0 {
		v1123 = v1101
		goto L327
	} else {
		goto L328
	}
L326:
	;
	if v1123 == int32(0) {
		goto L330
	} else {
		goto L331
	}
L327:
	;
	goto L326
L328:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1110 = v1105 - v1104
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1109 - v1110
	v1117 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v1117 != 0 {
		v1123 = v1101
		goto L327
	} else {
		goto L329
	}
L329:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1118 - v1110
	v1123 = int32(1)
	goto L327
L330:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L331:
	;
	goto L332
L332:
	;
	v1126 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1126 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L333
	}
L333:
	;
	v3602 = v1126
	goto L27
L334:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L335:
	;
	goto L336
L336:
	;
	v1133 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_32))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L31
	} else {
		goto L337
	}
L337:
	;
	if int32(0) <= v1133 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L338
	}
L338:
	;
	v3602 = v1133
	goto L27
L339:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L340:
	;
	goto L341
L341:
	;
	v1141 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_33))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L31
	} else {
		goto L342
	}
L342:
	;
	if int32(0) <= v1141 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L343
	}
L343:
	;
	v3602 = v1141
	goto L27
L344:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L345:
	;
	goto L346
L346:
	;
	v1149 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_34))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L31
	} else {
		goto L347
	}
L347:
	;
	if int32(0) <= v1149 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L348
	}
L348:
	;
	v3602 = v1149
	goto L27
L349:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L350:
	;
	goto L351
L351:
	;
	v1155 = F_slice_del(m, l0)
	mBase = m.M
	if v1155 < int32(0) {
		v3602 = v1155
		goto L27
	} else {
		goto L352
	}
L352:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1161 = F_insert_s(m, l0, v1158, v1158, int32(1), int32(_a_F_dutch_UTF_8_stem_35))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L31
	} else {
		goto L353
	}
L353:
	;
	if v1161 < int32(0) {
		v3602 = v1161
		goto L27
	} else {
		goto L354
	}
L354:
	;
	v1165 = F_r_lengthen_V_2(m, l0)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L31
	} else {
		goto L355
	}
L355:
	;
	if v1165 <= int32(0) {
		v1210 = v1165
		goto L176
	} else {
		goto L356
	}
L356:
	;
	v1252 = v618
	v1254 = v596
	goto L162
L357:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L358:
	;
	goto L359
L359:
	;
	v1171 = int32(0)
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1178 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v1178 != 0 {
		v1193 = v1171
		goto L361
	} else {
		goto L362
	}
L360:
	;
	if v1193 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L361:
	;
	goto L360
L362:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1180 = v1175 - v1174
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1179 - v1180
	v1187 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v1187 != 0 {
		v1193 = v1171
		goto L361
	} else {
		goto L363
	}
L363:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1188 - v1180
	v1193 = int32(1)
	goto L361
L364:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L365:
	;
	goto L366
L366:
	;
	v1196 = F_slice_del(m, l0)
	mBase = m.M
	if v1196 < int32(0) {
		v3602 = v1196
		goto L27
	} else {
		goto L367
	}
L367:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1202 = F_insert_s(m, l0, v1199, v1199, int32(2), int32(_a_F_dutch_UTF_8_stem_36))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L31
	} else {
		goto L368
	}
L368:
	;
	if v1202 < int32(0) {
		v3602 = v1202
		goto L27
	} else {
		goto L369
	}
L369:
	;
	v1206 = F_r_lengthen_V_2(m, l0)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L31
	} else {
		goto L370
	}
L370:
	;
	if int32(0) < v1206 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L371
	}
L371:
	;
	v1210 = v1206
	goto L176
L372:
	;
	v1213 = v1210
	goto L374
L373:
	;
	v1213 = v596
	goto L374
L374:
	;
	if v1210 != 0 {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v1214 = v1213
	goto L377
L376:
	;
	v1214 = v596
	goto L377
L377:
	;
	if v1210 != 0 {
		goto L174
	} else {
		goto L378
	}
L378:
	;
	v1252 = v595
	v1254 = v1214
	goto L162
L379:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L380:
	;
	goto L381
L381:
	;
	v1217 = int32(0)
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1224 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v1224 != 0 {
		v1239 = v1217
		goto L383
	} else {
		goto L384
	}
L382:
	;
	if v1239 == int32(0) {
		goto L386
	} else {
		goto L387
	}
L383:
	;
	goto L382
L384:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1226 = v1221 - v1220
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1225 - v1226
	v1233 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v1233 != 0 {
		v1239 = v1217
		goto L383
	} else {
		goto L385
	}
L385:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1234 - v1226
	v1239 = int32(1)
	goto L383
L386:
	;
	v1252 = v595
	v1254 = v596
	goto L162
L387:
	;
	goto L388
L388:
	;
	v1244 = F_slice_from_s(m, l0, int32(3), int32(_a_F_dutch_UTF_8_stem_37))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L31
	} else {
		goto L389
	}
L389:
	;
	if int32(0) <= v1244 {
		v1252 = v618
		v1254 = v596
		goto L162
	} else {
		goto L390
	}
L390:
	;
	v3602 = v1244
	goto L27
L391:
	;
	v1252 = v618
	v1254 = v1214
	goto L162
L392:
	;
	v1252 = v618
	v1254 = v596
	goto L162
L393:
	;
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1500
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1500
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1500-int32(2) <= v1503 {
		goto L513
	} else {
		goto L514
	}
L394:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L395:
	;
	goto L396
L396:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1266+v1259-int32(1)))))
	if v1270&int32(224) != int32(96) {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L398:
	;
	goto L399
L399:
	;
	if int32(1)<<(uint(v1270)%32)&int32(_a_F_dutch_UTF_8_stem_38) == int32(0) {
		goto L400
	} else {
		goto L401
	}
L400:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L401:
	;
	goto L402
L402:
	;
	v1284 = F_find_among_b(m, l0, int32(_a_F_dutch_UTF_8_stem_39), int32(14), int32(0))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L31
	} else {
		goto L403
	}
L403:
	;
	if v1284 == int32(0) {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L405:
	;
	goto L406
L406:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1288
	v1290 = int32(1)
	switch v1284 - v1290 {
	case 0:
		goto L417
	case 1:
		goto L416
	case 2:
		goto L415
	case 3:
		goto L414
	case 4:
		goto L413
	case 5:
		goto L412
	case 6:
		goto L411
	case 7:
		goto L410
	case 8:
		goto L409
	case 9:
		goto L408
	default:
		v1497 = v1290
		v1499 = v1254
		goto L393
	}
L407:
	;
	v1491 = base.B2i32(v1488 < int32(0))
	if v1488 < int32(0) {
		goto L502
	} else {
		goto L503
	}
L408:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1288 < v1455 {
		goto L490
	} else {
		goto L491
	}
L409:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1288 < v1422 {
		goto L478
	} else {
		goto L479
	}
L410:
	;
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1288 < v1406 {
		goto L470
	} else {
		goto L471
	}
L411:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1288 < v1390 {
		goto L462
	} else {
		goto L463
	}
L412:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1288 < v1357 {
		goto L450
	} else {
		goto L451
	}
L413:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1322 = int32(3)
	v1324 = int32(0)
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1326-v1327 < v1322 {
		v1337 = v1324
		goto L436
	} else {
		goto L437
	}
L414:
	;
	v1317 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_40))
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		goto L31
	} else {
		goto L433
	}
L415:
	;
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1288 < v1310 {
		goto L429
	} else {
		goto L430
	}
L416:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1288 < v1301 {
		goto L423
	} else {
		goto L424
	}
L417:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1288 < v1293 {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L419:
	;
	goto L420
L420:
	;
	v1297 = F_slice_from_s(m, l0, int32(3), int32(_a_F_dutch_UTF_8_stem_41))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L31
	} else {
		goto L421
	}
L421:
	;
	if int32(0) <= v1297 {
		v1497 = v1290
		v1499 = v1254
		goto L393
	} else {
		goto L422
	}
L422:
	;
	v3602 = v1297
	goto L27
L423:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L424:
	;
	goto L425
L425:
	;
	v1303 = F_slice_del(m, l0)
	mBase = m.M
	if v1303 < int32(0) {
		v3602 = v1303
		goto L27
	} else {
		goto L426
	}
L426:
	;
	v1306 = F_r_lengthen_V_2(m, l0)
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L31
	} else {
		goto L427
	}
L427:
	;
	if int32(0) < v1306 {
		v1497 = v1290
		v1499 = v1254
		goto L393
	} else {
		goto L428
	}
L428:
	;
	v1488 = v1306
	goto L407
L429:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L430:
	;
	goto L431
L431:
	;
	v1312 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1312 {
		v1497 = v1290
		v1499 = v1254
		goto L393
	} else {
		goto L432
	}
L432:
	;
	v3602 = v1312
	goto L27
L433:
	;
	if int32(0) <= v1317 {
		v1497 = v1290
		v1499 = v1254
		goto L393
	} else {
		goto L434
	}
L434:
	;
	v3602 = v1317
	goto L27
L435:
	;
	if v1337 != 0 {
		goto L439
	} else {
		goto L440
	}
L436:
	;
	goto L435
L437:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1333 = F_memcmp(m, v1330+v1326-v1322, int32(_a_F_dutch_UTF_8_stem_42), v1322)
	mBase = m.M
	if v1333 != 0 {
		v1337 = v1324
		goto L436
	} else {
		goto L438
	}
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1326 - v1322
	v1337 = int32(1)
	goto L436
L439:
	;
	v1340 = F_slice_from_s(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_43))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L31
	} else {
		goto L442
	}
L440:
	;
	goto L441
L441:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1346 = v1344 + (v1288 - v1321)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1346
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1346 < v1348 {
		goto L444
	} else {
		goto L445
	}
L442:
	;
	if int32(0) <= v1340 {
		v1497 = v1290
		v1499 = v1254
		goto L393
	} else {
		goto L443
	}
L443:
	;
	v3602 = v1340
	goto L27
L444:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L445:
	;
	goto L446
L446:
	;
	v1350 = F_slice_del(m, l0)
	mBase = m.M
	if v1350 < int32(0) {
		v3602 = v1350
		goto L27
	} else {
		goto L447
	}
L447:
	;
	v1353 = F_r_lengthen_V_2(m, l0)
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L31
	} else {
		goto L448
	}
L448:
	;
	if v1353 <= int32(0) {
		v1488 = v1353
		goto L407
	} else {
		goto L449
	}
L449:
	;
	v1497 = v1290
	v1499 = v1254
	goto L393
L450:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L451:
	;
	goto L452
L452:
	;
	v1359 = int32(0)
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1366 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v1366 != 0 {
		v1381 = v1359
		goto L454
	} else {
		goto L455
	}
L453:
	;
	if v1381 == int32(0) {
		goto L457
	} else {
		goto L458
	}
L454:
	;
	goto L453
L455:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1368 = v1363 - v1362
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1367 - v1368
	v1375 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v1375 != 0 {
		v1381 = v1359
		goto L454
	} else {
		goto L456
	}
L456:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1376 - v1368
	v1381 = int32(1)
	goto L454
L457:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L458:
	;
	goto L459
L459:
	;
	v1386 = F_slice_from_s(m, l0, int32(3), int32(_a_F_dutch_UTF_8_stem_44))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L31
	} else {
		goto L460
	}
L460:
	;
	if int32(0) <= v1386 {
		v1497 = v1290
		v1499 = v1254
		goto L393
	} else {
		goto L461
	}
L461:
	;
	v3602 = v1386
	goto L27
L462:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L463:
	;
	goto L464
L464:
	;
	v1392 = F_slice_del(m, l0)
	mBase = m.M
	if v1392 < int32(0) {
		v3602 = v1392
		goto L27
	} else {
		goto L465
	}
L465:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1398 = F_insert_s(m, l0, v1395, v1395, int32(1), int32(_a_F_dutch_UTF_8_stem_45))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L31
	} else {
		goto L466
	}
L466:
	;
	if v1398 < int32(0) {
		v3602 = v1398
		goto L27
	} else {
		goto L467
	}
L467:
	;
	v1402 = F_r_lengthen_V_2(m, l0)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L31
	} else {
		goto L468
	}
L468:
	;
	if v1402 <= int32(0) {
		v1488 = v1402
		goto L407
	} else {
		goto L469
	}
L469:
	;
	v1497 = v1290
	v1499 = v1254
	goto L393
L470:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L471:
	;
	goto L472
L472:
	;
	v1408 = F_slice_del(m, l0)
	mBase = m.M
	if v1408 < int32(0) {
		v3602 = v1408
		goto L27
	} else {
		goto L473
	}
L473:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1414 = F_insert_s(m, l0, v1411, v1411, int32(1), int32(_a_F_dutch_UTF_8_stem_46))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L31
	} else {
		goto L474
	}
L474:
	;
	if v1414 < int32(0) {
		v3602 = v1414
		goto L27
	} else {
		goto L475
	}
L475:
	;
	v1418 = F_r_lengthen_V_2(m, l0)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L31
	} else {
		goto L476
	}
L476:
	;
	if v1418 <= int32(0) {
		v1488 = v1418
		goto L407
	} else {
		goto L477
	}
L477:
	;
	v1497 = v1290
	v1499 = v1254
	goto L393
L478:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L479:
	;
	goto L480
L480:
	;
	v1424 = int32(0)
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1431 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v1431 != 0 {
		v1446 = v1424
		goto L482
	} else {
		goto L483
	}
L481:
	;
	if v1446 == int32(0) {
		goto L485
	} else {
		goto L486
	}
L482:
	;
	goto L481
L483:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1433 = v1428 - v1427
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1432 - v1433
	v1440 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v1440 != 0 {
		v1446 = v1424
		goto L482
	} else {
		goto L484
	}
L484:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1441 - v1433
	v1446 = int32(1)
	goto L482
L485:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L486:
	;
	goto L487
L487:
	;
	v1451 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_47))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L31
	} else {
		goto L488
	}
L488:
	;
	if int32(0) <= v1451 {
		v1497 = v1290
		v1499 = v1254
		goto L393
	} else {
		goto L489
	}
L489:
	;
	v3602 = v1451
	goto L27
L490:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L491:
	;
	goto L492
L492:
	;
	v1457 = int32(0)
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1464 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v1464 != 0 {
		v1479 = v1457
		goto L494
	} else {
		goto L495
	}
L493:
	;
	if v1479 == int32(0) {
		goto L497
	} else {
		goto L498
	}
L494:
	;
	goto L493
L495:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1466 = v1461 - v1460
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1465 - v1466
	v1473 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v1473 != 0 {
		v1479 = v1457
		goto L494
	} else {
		goto L496
	}
L496:
	;
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1474 - v1466
	v1479 = int32(1)
	goto L494
L497:
	;
	v1497 = v1252
	v1499 = v1254
	goto L393
L498:
	;
	goto L499
L499:
	;
	v1484 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_48))
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L31
	} else {
		goto L500
	}
L500:
	;
	if int32(0) <= v1484 {
		v1497 = v1290
		v1499 = v1254
		goto L393
	} else {
		goto L501
	}
L501:
	;
	v3602 = v1484
	goto L27
L502:
	;
	v1492 = v1488
	goto L504
L503:
	;
	v1492 = v1254
	goto L504
L504:
	;
	if v1488 != 0 {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v1493 = v1492
	goto L507
L506:
	;
	v1493 = v1254
	goto L507
L507:
	;
	if v1488 == int32(0) {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v1497 = v1252
	v1499 = v1493
	goto L393
L509:
	;
	goto L510
L510:
	;
	if v1488 < int32(0) {
		v3602 = v1493
		goto L27
	} else {
		goto L511
	}
L511:
	;
	v1497 = v1290
	v1499 = v1493
	goto L393
L512:
	;
	v1934 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v1934)
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1936
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1936
	v1942 = int32(2)
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1946-v1936 < v1942 {
		v1956 = v1934
		goto L638
	} else {
		goto L639
	}
L513:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1708
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1708
	v1712 = v1708 - int32(1)
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1712 <= v1713 {
		goto L580
	} else {
		goto L581
	}
L514:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1509 = int32(1)
	v1511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1507+v1500-v1509))))
	if base.B2i32(v1511&int32(224) != int32(96))|base.B2i32(v1509<<(uint(v1511)%32)&int32(_a_F_dutch_UTF_8_stem_49) == int32(0)) != 0 {
		goto L513
	} else {
		goto L515
	}
L515:
	;
	v1526 = F_find_among_b(m, l0, int32(_a_F_dutch_UTF_8_stem_50), int32(16), int32(0))
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L31
	} else {
		goto L516
	}
L516:
	;
	if v1526 == int32(0) {
		goto L513
	} else {
		goto L517
	}
L517:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1530
	v1532 = int32(1)
	switch v1526 - v1532 {
	case 0:
		goto L526
	case 1:
		goto L525
	case 2:
		goto L524
	case 3:
		goto L523
	case 4:
		goto L522
	case 5:
		goto L521
	case 6:
		goto L520
	case 7:
		goto L519
	case 8:
		goto L518
	default:
		v1931 = v1532
		v1933 = v1499
		goto L512
	}
L518:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1530 < v1671 {
		goto L513
	} else {
		goto L571
	}
L519:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1530 < v1663 {
		goto L513
	} else {
		goto L568
	}
L520:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1530 < v1655 {
		goto L513
	} else {
		goto L565
	}
L521:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1530 < v1622 {
		goto L513
	} else {
		goto L555
	}
L522:
	;
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1530 < v1589 {
		goto L513
	} else {
		goto L545
	}
L523:
	;
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1530 < v1556 {
		goto L513
	} else {
		goto L535
	}
L524:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1530 < v1551 {
		goto L513
	} else {
		goto L533
	}
L525:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1530 < v1543 {
		goto L513
	} else {
		goto L530
	}
L526:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1530 < v1535 {
		goto L513
	} else {
		goto L527
	}
L527:
	;
	v1539 = F_slice_from_s(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_51))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L31
	} else {
		goto L528
	}
L528:
	;
	if int32(0) <= v1539 {
		v1931 = v1532
		v1933 = v1499
		goto L512
	} else {
		goto L529
	}
L529:
	;
	v3602 = v1539
	goto L27
L530:
	;
	v1547 = F_slice_from_s(m, l0, int32(3), int32(_a_F_dutch_UTF_8_stem_52))
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L31
	} else {
		goto L531
	}
L531:
	;
	if int32(0) <= v1547 {
		v1931 = v1532
		v1933 = v1499
		goto L512
	} else {
		goto L532
	}
L532:
	;
	v3602 = v1547
	goto L27
L533:
	;
	v1553 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1553 {
		v1931 = v1532
		v1933 = v1499
		goto L512
	} else {
		goto L534
	}
L534:
	;
	v3602 = v1553
	goto L27
L535:
	;
	v1558 = int32(0)
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1562 = v1560 - v1561
	v1567 = F_in_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), v1558)
	mBase = m.M
	if v1567 != 0 {
		goto L538
	} else {
		goto L539
	}
L536:
	;
	if v1580 == int32(0) {
		goto L513
	} else {
		goto L542
	}
L537:
	;
	goto L536
L538:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1568 - v1562
	v1573 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_12))
	mBase = m.M
	if v1573 == int32(0) {
		v1580 = v1558
		goto L537
	} else {
		goto L541
	}
L539:
	;
	goto L540
L540:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1576 - v1562
	v1580 = int32(1)
	goto L537
L541:
	;
	goto L540
L542:
	;
	v1585 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_53))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L31
	} else {
		goto L543
	}
L543:
	;
	if int32(0) <= v1585 {
		v1931 = v1532
		v1933 = v1499
		goto L512
	} else {
		goto L544
	}
L544:
	;
	v3602 = v1585
	goto L27
L545:
	;
	v1591 = int32(0)
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1595 = v1593 - v1594
	v1600 = F_in_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), v1591)
	mBase = m.M
	if v1600 != 0 {
		goto L548
	} else {
		goto L549
	}
L546:
	;
	if v1613 == int32(0) {
		goto L513
	} else {
		goto L552
	}
L547:
	;
	goto L546
L548:
	;
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1601 - v1595
	v1606 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_12))
	mBase = m.M
	if v1606 == int32(0) {
		v1613 = v1591
		goto L547
	} else {
		goto L551
	}
L549:
	;
	goto L550
L550:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1609 - v1595
	v1613 = int32(1)
	goto L547
L551:
	;
	goto L550
L552:
	;
	v1618 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_54))
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L31
	} else {
		goto L553
	}
L553:
	;
	if int32(0) <= v1618 {
		v1931 = v1532
		v1933 = v1499
		goto L512
	} else {
		goto L554
	}
L554:
	;
	v3602 = v1618
	goto L27
L555:
	;
	v1624 = int32(0)
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1628 = v1626 - v1627
	v1633 = F_in_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), v1624)
	mBase = m.M
	if v1633 != 0 {
		goto L558
	} else {
		goto L559
	}
L556:
	;
	if v1646 == int32(0) {
		goto L513
	} else {
		goto L562
	}
L557:
	;
	goto L556
L558:
	;
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1634 - v1628
	v1639 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_12))
	mBase = m.M
	if v1639 == int32(0) {
		v1646 = v1624
		goto L557
	} else {
		goto L561
	}
L559:
	;
	goto L560
L560:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1642 - v1628
	v1646 = int32(1)
	goto L557
L561:
	;
	goto L560
L562:
	;
	v1651 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_55))
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L31
	} else {
		goto L563
	}
L563:
	;
	if int32(0) <= v1651 {
		v1931 = v1532
		v1933 = v1499
		goto L512
	} else {
		goto L564
	}
L564:
	;
	v3602 = v1651
	goto L27
L565:
	;
	v1659 = F_slice_from_s(m, l0, int32(4), int32(_a_F_dutch_UTF_8_stem_56))
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L31
	} else {
		goto L566
	}
L566:
	;
	if int32(0) <= v1659 {
		v1931 = v1532
		v1933 = v1499
		goto L512
	} else {
		goto L567
	}
L567:
	;
	v3602 = v1659
	goto L27
L568:
	;
	v1667 = F_slice_from_s(m, l0, int32(4), int32(_a_F_dutch_UTF_8_stem_57))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L31
	} else {
		goto L569
	}
L569:
	;
	if int32(0) <= v1667 {
		v1931 = v1532
		v1933 = v1499
		goto L512
	} else {
		goto L570
	}
L570:
	;
	v3602 = v1667
	goto L27
L571:
	;
	v1673 = int32(0)
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1680 = F_eq_s_b(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	if v1680 != 0 {
		v1695 = v1673
		goto L573
	} else {
		goto L574
	}
L572:
	;
	if v1695 == int32(0) {
		goto L513
	} else {
		goto L576
	}
L573:
	;
	goto L572
L574:
	;
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1682 = v1677 - v1676
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1681 - v1682
	v1689 = F_out_grouping_b_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v1689 != 0 {
		v1695 = v1673
		goto L573
	} else {
		goto L575
	}
L575:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1690 - v1682
	v1695 = int32(1)
	goto L573
L576:
	;
	v1698 = F_slice_del(m, l0)
	mBase = m.M
	if v1698 < int32(0) {
		v3602 = v1698
		goto L27
	} else {
		goto L577
	}
L577:
	;
	v1701 = F_r_lengthen_V_2(m, l0)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L31
	} else {
		goto L578
	}
L578:
	;
	if int32(0) <= v1701 {
		v1931 = v1532
		v1933 = v1499
		goto L512
	} else {
		goto L579
	}
L579:
	;
	v3602 = v1701
	goto L27
L580:
	;
	v1931 = v1497
	v1933 = v1499
	goto L512
L581:
	;
	goto L582
L582:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1715+v1712))))
	if v1717&int32(224) != int32(96) {
		goto L583
	} else {
		goto L584
	}
L583:
	;
	v1931 = v1497
	v1933 = v1499
	goto L512
L584:
	;
	goto L585
L585:
	;
	if int32(1)<<(uint(v1717)%32)&int32(_a_F_dutch_UTF_8_stem_58) == int32(0) {
		goto L586
	} else {
		goto L587
	}
L586:
	;
	v1931 = v1497
	v1933 = v1499
	goto L512
L587:
	;
	goto L588
L588:
	;
	v1731 = F_find_among_b(m, l0, int32(_a_F_dutch_UTF_8_stem_59), int32(3), int32(0))
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L31
	} else {
		goto L589
	}
L589:
	;
	if v1731 == int32(0) {
		goto L590
	} else {
		goto L591
	}
L590:
	;
	v1931 = v1497
	v1933 = v1499
	goto L512
L591:
	;
	goto L592
L592:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1735
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1735 < v1737 {
		goto L593
	} else {
		goto L594
	}
L593:
	;
	v1931 = v1497
	v1933 = v1499
	goto L512
L594:
	;
	goto L595
L595:
	;
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1740 = int32(3)
	v1742 = int32(0)
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1744-v1745 < v1740 {
		v1755 = v1742
		goto L598
	} else {
		goto L599
	}
L596:
	;
	v1761 = v1735 - v1739
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1763 = v1761 + v1762
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1763
	v1765 = int32(2)
	v1767 = int32(0)
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1763-v1770 < v1765 {
		v1780 = v1767
		goto L604
	} else {
		goto L605
	}
L597:
	;
	if v1755 == int32(0) {
		goto L596
	} else {
		goto L601
	}
L598:
	;
	goto L597
L599:
	;
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1751 = F_memcmp(m, v1748+v1744-v1740, int32(_a_F_dutch_UTF_8_stem_60), v1740)
	mBase = m.M
	if v1751 != 0 {
		v1755 = v1742
		goto L598
	} else {
		goto L600
	}
L600:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1744 - v1740
	v1755 = int32(1)
	goto L598
L601:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1759 < v1758 {
		goto L596
	} else {
		goto L602
	}
L602:
	;
	v1931 = v1497
	v1933 = v1499
	goto L512
L603:
	;
	if v1780 != 0 {
		goto L607
	} else {
		goto L608
	}
L604:
	;
	goto L603
L605:
	;
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1776 = F_memcmp(m, v1773+v1763-v1765, int32(_a_F_dutch_UTF_8_stem_6), v1765)
	mBase = m.M
	if v1776 != 0 {
		v1780 = v1767
		goto L604
	} else {
		goto L606
	}
L606:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1763 - v1765
	v1780 = int32(1)
	goto L604
L607:
	;
	v1931 = v1497
	v1933 = v1499
	goto L512
L608:
	;
	goto L609
L609:
	;
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1782 = v1781 + v1761
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1782
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L612
L610:
	;
	if v1913 != 0 {
		goto L628
	} else {
		goto L629
	}
L611:
	;
	v1913 = v1906
	goto L610
L612:
	;
	if v1782 <= v1797 {
		v1906 = int32(-1)
		goto L611
	} else {
		goto L614
	}
L613:
	;
	v1906 = int32(0)
	goto L611
L614:
	;
	v1814 = int32(1)
	v1815 = v1782 - v1814
	v1817 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1798+v1815))))
	v1819 = v1817 & int32(255)
	if base.B2i32(v1815 == v1797)|base.B2i32(int32(0) <= v1817) != 0 {
		v1877 = v1819
		v1881 = v1814
		goto L615
	} else {
		goto L616
	}
L615:
	;
	if int32(252) < v1877 {
		goto L623
	} else {
		goto L624
	}
L616:
	;
	v1826 = v1819 & int32(63)
	v1828 = v1782 - int32(2)
	v1830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1798+v1828))))
	v1832 = v1830 << (uint(int32(6)) % 32)
	if base.B2i32(v1828 != v1797)&base.B2i32(base.Ui32(v1830) < base.Ui32(int32(192))) == int32(0) {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	v1877 = v1832&int32(1984) | v1826
	v1881 = int32(2)
	goto L615
L618:
	;
	goto L619
L619:
	;
	v1845 = v1832&int32(4032) | v1826
	v1847 = v1782 - int32(3)
	v1849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1798+v1847))))
	if base.B2i32(v1847 != v1797)&base.B2i32(base.Ui32(v1849) < base.Ui32(int32(224))) == int32(0) {
		goto L620
	} else {
		goto L621
	}
L620:
	;
	v1877 = v1849<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_22) | v1845
	v1881 = int32(3)
	goto L615
L621:
	;
	goto L622
L622:
	;
	v1867 = int32(4)
	v1869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1782+v1798-v1867))))
	v1877 = v1849<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_23) | v1869&int32(7)<<(uint(int32(18))%32) | v1845
	v1881 = v1867
	goto L615
L623:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1782 - v1881
	goto L627
L624:
	;
	v1883 = v1877 - int32(97)
	if v1883 < int32(0) {
		goto L623
	} else {
		goto L625
	}
L625:
	;
	v1889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1883)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[0]))))
	if int32(base.Ui32(v1889)>>(uint(v1883&int32(7))%32))&int32(1) == int32(0) {
		goto L623
	} else {
		goto L626
	}
L626:
	;
	v1913 = v1881
	goto L610
L627:
	;
	goto L613
L628:
	;
	v1931 = v1497
	v1933 = v1499
	goto L512
L629:
	;
	goto L630
L630:
	;
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1914 + v1761
	v1917 = F_slice_del(m, l0)
	mBase = m.M
	if v1917 < int32(0) {
		v3602 = v1917
		goto L27
	} else {
		goto L631
	}
L631:
	;
	v1921 = F_r_lengthen_V_2(m, l0)
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L31
	} else {
		goto L632
	}
L632:
	;
	if int32(0) < v1921 {
		v1931 = int32(1)
		v1933 = v1499
		goto L512
	} else {
		goto L633
	}
L633:
	;
	if v1921 == int32(0) {
		v1931 = v1497
		v1933 = v1499
		goto L512
	} else {
		goto L634
	}
L634:
	;
	if v1921 < int32(0) {
		v3602 = v1921
		goto L27
	} else {
		goto L635
	}
L635:
	;
	v1931 = int32(1)
	v1933 = v1921
	goto L512
L636:
	;
	if v2489 != 0 {
		goto L773
	} else {
		goto L774
	}
L637:
	;
	if v1956 == int32(0) {
		v2489 = v1934
		goto L636
	} else {
		goto L641
	}
L638:
	;
	goto L637
L639:
	;
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1952 = F_memcmp(m, v1950+v1936, int32(_a_F_dutch_UTF_8_stem_61), v1942)
	mBase = m.M
	if v1952 != 0 {
		v1956 = v1934
		goto L638
	} else {
		goto L640
	}
L640:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1942 + v1936
	v1956 = int32(1)
	goto L638
L641:
	;
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1959
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L645
L642:
	;
	v2489 = v2475
	goto L636
L643:
	;
	if v2014 < int32(0) {
		v2475 = v1934
		goto L642
	} else {
		goto L663
	}
L645:
	;
	goto L646
L646:
	;
	goto L647
L647:
	;
	v1969 = v1959
	v1971 = int32(3)
	goto L650
L649:
	;
	v2014 = v1999
	goto L643
L650:
	;
	if v1962 <= v1969 {
		goto L652
	} else {
		goto L653
	}
L651:
	;
	goto L649
L652:
	;
	v2014 = int32(-1)
	goto L643
L653:
	;
	goto L654
L654:
	;
	v1976 = v1969 + int32(1)
	v1978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1961+v1969))))
	if base.Ui32(v1978) < base.Ui32(int32(192)) {
		v1999 = v1976
		goto L655
	} else {
		goto L656
	}
L655:
	;
	v2000 = int32(1)
	if v2000 < v1971 {
		v1969 = v1999
		v1971 = v1971 - v2000
		goto L650
	} else {
		goto L662
	}
L656:
	;
	if v1962 <= v1976 {
		v1999 = v1976
		goto L655
	} else {
		goto L657
	}
L657:
	;
	v1985 = v1976
	goto L658
L658:
	;
	v1988 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1961+v1985))))
	if int32(-65) < v1988 {
		v1999 = v1985
		goto L655
	} else {
		goto L660
	}
L659:
	;
	v1999 = v1962
	goto L655
L660:
	;
	v1992 = v1985 + int32(1)
	if v1992 != v1962 {
		v1985 = v1992
		goto L658
	} else {
		goto L661
	}
L661:
	;
	goto L659
L662:
	;
	goto L651
L663:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1959
	v2018 = int32(2)
	v2020 = int32(0)
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2022-v1959 < v2018 {
		v2032 = v2020
		goto L666
	} else {
		goto L667
	}
L664:
	;
	goto L723
L665:
	;
	if v2032 != 0 {
		goto L664
	} else {
		goto L669
	}
L666:
	;
	goto L665
L667:
	;
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2028 = F_memcmp(m, v2026+v1959, int32(_a_F_dutch_UTF_8_stem_62), v2018)
	mBase = m.M
	if v2028 != 0 {
		v2032 = v2020
		goto L666
	} else {
		goto L668
	}
L668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2018 + v1959
	v2032 = int32(1)
	goto L666
L669:
	;
	v2034 = v1959
	goto L670
L670:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2034
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L674
L671:
	;
	goto L664
L672:
	;
	if v2161 == int32(0) {
		goto L664
	} else {
		goto L696
	}
L673:
	;
	v2161 = v2154
	goto L672
L674:
	;
	if v2056 <= v2034 {
		goto L676
	} else {
		goto L677
	}
L675:
	;
	v2154 = int32(0)
	goto L673
L676:
	;
	v2161 = int32(-1)
	goto L672
L677:
	;
	goto L678
L678:
	;
	v2072 = int32(1)
	v2074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2034+v2057))))
	if base.Ui32(v2074) < base.Ui32(int32(192)) {
		v2131 = v2074
		v2132 = v2072
		goto L679
	} else {
		goto L680
	}
L679:
	;
	if int32(252) < v2131 {
		v2154 = v2132
		goto L673
	} else {
		goto L692
	}
L680:
	;
	v2078 = v2034 + int32(1)
	if v2078 == v2056 {
		v2131 = v2074
		v2132 = v2072
		goto L679
	} else {
		goto L681
	}
L681:
	;
	v2081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2078+v2057))))
	v2083 = v2081 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v2074) {
		goto L683
	} else {
		goto L684
	}
L682:
	;
	v2097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2087+v2057))))
	v2099 = v2097 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v2074) {
		goto L688
	} else {
		goto L689
	}
L683:
	;
	v2087 = v2034 + int32(2)
	if v2087 != v2056 {
		goto L682
	} else {
		goto L686
	}
L684:
	;
	goto L685
L685:
	;
	v2131 = v2074<<(uint(int32(6))%32)&int32(1984) | v2083
	v2132 = int32(2)
	goto L679
L686:
	;
	goto L685
L687:
	;
	v2116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2057+v2103))))
	v2131 = v2116&int32(63) | (v2074<<(uint(int32(18))%32)&int32(_a_F_dutch_UTF_8_stem_63) | v2083<<(uint(int32(12))%32) | v2099<<(uint(int32(6))%32))
	v2132 = int32(4)
	goto L679
L688:
	;
	v2103 = v2034 + int32(3)
	if v2103 != v2056 {
		goto L687
	} else {
		goto L691
	}
L689:
	;
	goto L690
L690:
	;
	v2131 = v2074<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_22) | v2083<<(uint(int32(6))%32) | v2099
	v2132 = int32(3)
	goto L679
L691:
	;
	goto L690
L692:
	;
	v2136 = v2131 - int32(97)
	if v2136 < int32(0) {
		v2154 = v2132
		goto L673
	} else {
		goto L693
	}
L693:
	;
	v2142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2136)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[0]))))
	if int32(base.Ui32(v2142)>>(uint(v2136&int32(7))%32))&int32(1) == int32(0) {
		v2154 = v2132
		goto L673
	} else {
		goto L694
	}
L694:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2132 + v2034
	goto L695
L695:
	;
	goto L675
L696:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2034
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L699
L697:
	;
	if v2218 < int32(0) {
		v2475 = v1934
		goto L642
	} else {
		goto L717
	}
L699:
	;
	goto L700
L700:
	;
	goto L701
L701:
	;
	v2173 = v2034
	v2175 = int32(1)
	goto L704
L703:
	;
	v2218 = v2203
	goto L697
L704:
	;
	if v2166 <= v2173 {
		goto L706
	} else {
		goto L707
	}
L705:
	;
	goto L703
L706:
	;
	v2218 = int32(-1)
	goto L697
L707:
	;
	goto L708
L708:
	;
	v2180 = v2173 + int32(1)
	v2182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2165+v2173))))
	if base.Ui32(v2182) < base.Ui32(int32(192)) {
		v2203 = v2180
		goto L709
	} else {
		goto L710
	}
L709:
	;
	v2204 = int32(1)
	if v2204 < v2175 {
		v2173 = v2203
		v2175 = v2175 - v2204
		goto L704
	} else {
		goto L716
	}
L710:
	;
	if v2166 <= v2180 {
		v2203 = v2180
		goto L709
	} else {
		goto L711
	}
L711:
	;
	v2189 = v2180
	goto L712
L712:
	;
	v2192 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2165+v2189))))
	if int32(-65) < v2192 {
		v2203 = v2189
		goto L709
	} else {
		goto L714
	}
L713:
	;
	v2203 = v2166
	goto L709
L714:
	;
	v2196 = v2189 + int32(1)
	if v2196 != v2166 {
		v2189 = v2196
		goto L712
	} else {
		goto L715
	}
L715:
	;
	goto L713
L716:
	;
	goto L705
L717:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2218
	v2222 = int32(2)
	v2224 = int32(0)
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2226-v2218 < v2222 {
		v2236 = v2224
		goto L719
	} else {
		goto L720
	}
L718:
	;
	if v2236 == int32(0) {
		v2034 = v2218
		goto L670
	} else {
		goto L722
	}
L719:
	;
	goto L718
L720:
	;
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2232 = F_memcmp(m, v2230+v2218, int32(_a_F_dutch_UTF_8_stem_62), v2222)
	mBase = m.M
	if v2232 != 0 {
		v2236 = v2224
		goto L719
	} else {
		goto L721
	}
L721:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2222 + v2218
	v2236 = int32(1)
	goto L719
L722:
	;
	goto L671
L723:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2260 = int32(2)
	v2262 = int32(0)
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2264-v2259 < v2260 {
		v2274 = v2262
		goto L726
	} else {
		goto L727
	}
L724:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2259
	v2398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2398 <= v2259 {
		v2489 = int32(0)
		goto L636
	} else {
		goto L755
	}
L725:
	;
	if v2274 != 0 {
		goto L723
	} else {
		goto L729
	}
L726:
	;
	goto L725
L727:
	;
	v2268 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2270 = F_memcmp(m, v2268+v2259, int32(_a_F_dutch_UTF_8_stem_64), v2260)
	mBase = m.M
	if v2270 != 0 {
		v2274 = v2262
		goto L726
	} else {
		goto L728
	}
L728:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2260 + v2259
	v2274 = int32(1)
	goto L726
L729:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2259
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L732
L730:
	;
	if v2393 == int32(0) {
		goto L723
	} else {
		goto L754
	}
L731:
	;
	v2393 = v2386
	goto L730
L732:
	;
	if v2288 <= v2259 {
		goto L734
	} else {
		goto L735
	}
L733:
	;
	v2386 = int32(0)
	goto L731
L734:
	;
	v2393 = int32(-1)
	goto L730
L735:
	;
	goto L736
L736:
	;
	v2304 = int32(1)
	v2306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2259+v2289))))
	if base.Ui32(v2306) < base.Ui32(int32(192)) {
		v2363 = v2306
		v2364 = v2304
		goto L737
	} else {
		goto L738
	}
L737:
	;
	if int32(252) < v2363 {
		v2386 = v2364
		goto L731
	} else {
		goto L750
	}
L738:
	;
	v2310 = v2259 + int32(1)
	if v2310 == v2288 {
		v2363 = v2306
		v2364 = v2304
		goto L737
	} else {
		goto L739
	}
L739:
	;
	v2313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2310+v2289))))
	v2315 = v2313 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v2306) {
		goto L741
	} else {
		goto L742
	}
L740:
	;
	v2329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2319+v2289))))
	v2331 = v2329 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v2306) {
		goto L746
	} else {
		goto L747
	}
L741:
	;
	v2319 = v2259 + int32(2)
	if v2319 != v2288 {
		goto L740
	} else {
		goto L744
	}
L742:
	;
	goto L743
L743:
	;
	v2363 = v2306<<(uint(int32(6))%32)&int32(1984) | v2315
	v2364 = int32(2)
	goto L737
L744:
	;
	goto L743
L745:
	;
	v2348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2289+v2335))))
	v2363 = v2348&int32(63) | (v2306<<(uint(int32(18))%32)&int32(_a_F_dutch_UTF_8_stem_63) | v2315<<(uint(int32(12))%32) | v2331<<(uint(int32(6))%32))
	v2364 = int32(4)
	goto L737
L746:
	;
	v2335 = v2259 + int32(3)
	if v2335 != v2288 {
		goto L745
	} else {
		goto L749
	}
L747:
	;
	goto L748
L748:
	;
	v2363 = v2306<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_22) | v2315<<(uint(int32(6))%32) | v2331
	v2364 = int32(3)
	goto L737
L749:
	;
	goto L748
L750:
	;
	v2368 = v2363 - int32(97)
	if v2368 < int32(0) {
		v2386 = v2364
		goto L731
	} else {
		goto L751
	}
L751:
	;
	v2374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2368)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[0]))))
	if int32(base.Ui32(v2374)>>(uint(v2368&int32(7))%32))&int32(1) == int32(0) {
		v2386 = v2364
		goto L731
	} else {
		goto L752
	}
L752:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2364 + v2259
	goto L753
L753:
	;
	goto L733
L754:
	;
	goto L724
L755:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1959
	v2402 = v1959 + int32(2)
	if v2398 <= v2402 {
		goto L756
	} else {
		goto L757
	}
L756:
	;
	v2426 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v2426)
	v2428 = F_slice_del(m, l0)
	mBase = m.M
	if v2428 < int32(0) {
		v2475 = v2428
		goto L642
	} else {
		goto L761
	}
L757:
	;
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2404+v2402))))
	if base.B2i32(v2406&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v2406)%32)&int32(_a_F_dutch_UTF_8_stem_65) == int32(0)) != 0 {
		goto L756
	} else {
		goto L758
	}
L758:
	;
	v2421 = F_find_among(m, l0, int32(_a_F_dutch_UTF_8_stem_66), int32(6), int32(0))
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		goto L31
	} else {
		goto L759
	}
L759:
	;
	if v2421 == int32(1) {
		v2475 = v1934
		goto L642
	} else {
		goto L760
	}
L760:
	;
	goto L756
L761:
	;
	v2431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2431
	v2434 = v2431 + int32(1)
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2435 <= v2434 {
		goto L762
	} else {
		goto L763
	}
L762:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2431
	v2475 = int32(1)
	goto L642
L763:
	;
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2437+v2434))))
	switch v2439 - int32(171) {
	case 0, 4:
		goto L764
	default:
		goto L762
	}
L764:
	;
	v2445 = F_find_among(m, l0, int32(_a_F_dutch_UTF_8_stem_67), int32(2), int32(0))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L31
	} else {
		goto L765
	}
L765:
	;
	if v2445 == int32(0) {
		goto L762
	} else {
		goto L766
	}
L766:
	;
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2449
	switch v2445 - int32(1) {
	case 0:
		goto L768
	case 1:
		goto L767
	default:
		goto L762
	}
L767:
	;
	v2461 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_68))
	mBase = m.M
	v2462 = m.ExcPending
	if v2462 != 0 {
		goto L31
	} else {
		goto L771
	}
L768:
	;
	v2455 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_69))
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		goto L31
	} else {
		goto L769
	}
L769:
	;
	if int32(0) <= v2455 {
		goto L762
	} else {
		goto L770
	}
L770:
	;
	v2475 = v2455
	goto L642
L771:
	;
	if v2461 < int32(0) {
		v2475 = v2461
		goto L642
	} else {
		goto L772
	}
L772:
	;
	goto L762
L773:
	;
	if v2489 < int32(0) {
		v3602 = v2489
		goto L27
	} else {
		goto L776
	}
L774:
	;
	goto L775
L775:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1936
	v2586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2586
	v2588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v2588 != 0 {
		goto L803
	} else {
		goto L804
	}
L776:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1936
	v2496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v2496
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v2496
	goto L778
L777:
	;
	goto L775
L778:
	;
	v2507 = int32(0)
	v2508 = F_out_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), v2507)
	mBase = m.M
	if v2508 == v2507 {
		goto L778
	} else {
		goto L780
	}
L779:
	;
	v2513 = int32(1)
	goto L781
L780:
	;
	goto L779
L781:
	;
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2519 = F_eq_s(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_1))
	mBase = m.M
	if v2519 == int32(0) {
		goto L784
	} else {
		goto L785
	}
L782:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2516
	if int32(0) < v2513 {
		goto L788
	} else {
		goto L789
	}
L783:
	;
	goto L782
L784:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2516
	v2527 = F_in_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v2527 != 0 {
		goto L783
	} else {
		goto L787
	}
L785:
	;
	goto L786
L786:
	;
	v2513 = v2513 - int32(1)
	goto L781
L787:
	;
	goto L786
L788:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1936
	goto L777
L789:
	;
	v2537 = F_out_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v2537 != 0 {
		goto L788
	} else {
		goto L790
	}
L790:
	;
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v2538
	goto L791
L791:
	;
	v2547 = int32(0)
	v2548 = F_out_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), v2547)
	mBase = m.M
	if v2548 == v2547 {
		goto L791
	} else {
		goto L793
	}
L792:
	;
	v2553 = int32(1)
	goto L794
L793:
	;
	goto L792
L794:
	;
	v2556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2559 = F_eq_s(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_2))
	mBase = m.M
	if v2559 == int32(0) {
		goto L797
	} else {
		goto L798
	}
L795:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2556
	if int32(0) < v2553 {
		goto L788
	} else {
		goto L801
	}
L796:
	;
	goto L795
L797:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2556
	v2567 = F_in_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v2567 != 0 {
		goto L796
	} else {
		goto L800
	}
L798:
	;
	goto L799
L799:
	;
	v2553 = v2553 - int32(1)
	goto L794
L800:
	;
	goto L799
L801:
	;
	v2577 = F_out_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v2577 != 0 {
		goto L788
	} else {
		goto L802
	}
L802:
	;
	v2578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v2578
	goto L788
L803:
	;
	v2589 = F_r_Step_1c_2(m, l0)
	mBase = m.M
	v2590 = m.ExcPending
	if v2590 != 0 {
		goto L31
	} else {
		goto L806
	}
L804:
	;
	v2598 = v1931
	v2599 = v1936
	v2600 = v1933
	goto L805
L805:
	;
	v2601 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v2601)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2599
	v2605 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L819
L806:
	;
	v2592 = base.B2i32(v2589 < int32(0))
	if v2589 < int32(0) {
		goto L807
	} else {
		goto L808
	}
L807:
	;
	v2593 = v2589
	goto L809
L808:
	;
	v2593 = v1933
	goto L809
L809:
	;
	if v2589 != 0 {
		goto L810
	} else {
		goto L811
	}
L810:
	;
	v2594 = v2593
	goto L812
L811:
	;
	v2594 = v1933
	goto L812
L812:
	;
	if v2589 < int32(0) {
		goto L813
	} else {
		goto L814
	}
L813:
	;
	return v2594
L814:
	;
	goto L815
L815:
	;
	v2597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2598 = int32(1)
	v2599 = v2597
	v2600 = v2594
	goto L805
L816:
	;
	if v3233 != 0 {
		goto L993
	} else {
		goto L994
	}
L817:
	;
	if v2659 < int32(0) {
		v3233 = v2601
		goto L816
	} else {
		goto L837
	}
L819:
	;
	goto L820
L820:
	;
	goto L821
L821:
	;
	v2614 = v2599
	v2616 = int32(1)
	goto L824
L823:
	;
	v2659 = v2644
	goto L817
L824:
	;
	if v2607 <= v2614 {
		goto L826
	} else {
		goto L827
	}
L825:
	;
	goto L823
L826:
	;
	v2659 = int32(-1)
	goto L817
L827:
	;
	goto L828
L828:
	;
	v2621 = v2614 + int32(1)
	v2623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2605+v2614))))
	if base.Ui32(v2623) < base.Ui32(int32(192)) {
		v2644 = v2621
		goto L829
	} else {
		goto L830
	}
L829:
	;
	v2645 = int32(1)
	if v2645 < v2616 {
		v2614 = v2644
		v2616 = v2616 - v2645
		goto L824
	} else {
		goto L836
	}
L830:
	;
	if v2607 <= v2621 {
		v2644 = v2621
		goto L829
	} else {
		goto L831
	}
L831:
	;
	v2630 = v2621
	goto L832
L832:
	;
	v2633 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2605+v2630))))
	if int32(-65) < v2633 {
		v2644 = v2630
		goto L829
	} else {
		goto L834
	}
L833:
	;
	v2644 = v2607
	goto L829
L834:
	;
	v2637 = v2630 + int32(1)
	if v2637 != v2607 {
		v2630 = v2637
		goto L832
	} else {
		goto L835
	}
L835:
	;
	goto L833
L836:
	;
	goto L825
L837:
	;
	v2663 = v2659
	goto L838
L838:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2663
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2663
	v2674 = int32(2)
	v2676 = int32(0)
	v2678 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2678-v2663 < v2674 {
		v2688 = v2676
		goto L841
	} else {
		goto L842
	}
L839:
	;
	v3233 = v2601
	goto L816
L840:
	;
	if v2688 != 0 {
		goto L844
	} else {
		goto L845
	}
L841:
	;
	goto L840
L842:
	;
	v2682 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2684 = F_memcmp(m, v2682+v2663, int32(_a_F_dutch_UTF_8_stem_70), v2674)
	mBase = m.M
	if v2684 != 0 {
		v2688 = v2676
		goto L841
	} else {
		goto L843
	}
L843:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2674 + v2663
	v2688 = int32(1)
	goto L841
L844:
	;
	v2689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2689
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L849
L845:
	;
	goto L846
L846:
	;
	v3172 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L974
L847:
	;
	if v2744 < int32(0) {
		v3233 = v2601
		goto L816
	} else {
		goto L867
	}
L849:
	;
	goto L850
L850:
	;
	goto L851
L851:
	;
	v2699 = v2689
	v2701 = int32(3)
	goto L854
L853:
	;
	v2744 = v2729
	goto L847
L854:
	;
	if v2692 <= v2699 {
		goto L856
	} else {
		goto L857
	}
L855:
	;
	goto L853
L856:
	;
	v2744 = int32(-1)
	goto L847
L857:
	;
	goto L858
L858:
	;
	v2706 = v2699 + int32(1)
	v2708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2691+v2699))))
	if base.Ui32(v2708) < base.Ui32(int32(192)) {
		v2729 = v2706
		goto L859
	} else {
		goto L860
	}
L859:
	;
	v2730 = int32(1)
	if v2730 < v2701 {
		v2699 = v2729
		v2701 = v2701 - v2730
		goto L854
	} else {
		goto L866
	}
L860:
	;
	if v2692 <= v2706 {
		v2729 = v2706
		goto L859
	} else {
		goto L861
	}
L861:
	;
	v2715 = v2706
	goto L862
L862:
	;
	v2718 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2691+v2715))))
	if int32(-65) < v2718 {
		v2729 = v2715
		goto L859
	} else {
		goto L864
	}
L863:
	;
	v2729 = v2692
	goto L859
L864:
	;
	v2722 = v2715 + int32(1)
	if v2722 != v2692 {
		v2715 = v2722
		goto L862
	} else {
		goto L865
	}
L865:
	;
	goto L863
L866:
	;
	goto L855
L867:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2689
	v2748 = int32(2)
	v2750 = int32(0)
	v2752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2752-v2689 < v2748 {
		v2762 = v2750
		goto L870
	} else {
		goto L871
	}
L868:
	;
	goto L927
L869:
	;
	if v2762 != 0 {
		goto L868
	} else {
		goto L873
	}
L870:
	;
	goto L869
L871:
	;
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2758 = F_memcmp(m, v2756+v2689, int32(_a_F_dutch_UTF_8_stem_71), v2748)
	mBase = m.M
	if v2758 != 0 {
		v2762 = v2750
		goto L870
	} else {
		goto L872
	}
L872:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2748 + v2689
	v2762 = int32(1)
	goto L870
L873:
	;
	v2764 = v2689
	goto L874
L874:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2764
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2787 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L878
L875:
	;
	goto L868
L876:
	;
	if v2891 == int32(0) {
		goto L868
	} else {
		goto L900
	}
L877:
	;
	v2891 = v2884
	goto L876
L878:
	;
	if v2786 <= v2764 {
		goto L880
	} else {
		goto L881
	}
L879:
	;
	v2884 = int32(0)
	goto L877
L880:
	;
	v2891 = int32(-1)
	goto L876
L881:
	;
	goto L882
L882:
	;
	v2802 = int32(1)
	v2804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2764+v2787))))
	if base.Ui32(v2804) < base.Ui32(int32(192)) {
		v2861 = v2804
		v2862 = v2802
		goto L883
	} else {
		goto L884
	}
L883:
	;
	if int32(252) < v2861 {
		v2884 = v2862
		goto L877
	} else {
		goto L896
	}
L884:
	;
	v2808 = v2764 + int32(1)
	if v2808 == v2786 {
		v2861 = v2804
		v2862 = v2802
		goto L883
	} else {
		goto L885
	}
L885:
	;
	v2811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2808+v2787))))
	v2813 = v2811 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v2804) {
		goto L887
	} else {
		goto L888
	}
L886:
	;
	v2827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2817+v2787))))
	v2829 = v2827 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v2804) {
		goto L892
	} else {
		goto L893
	}
L887:
	;
	v2817 = v2764 + int32(2)
	if v2817 != v2786 {
		goto L886
	} else {
		goto L890
	}
L888:
	;
	goto L889
L889:
	;
	v2861 = v2804<<(uint(int32(6))%32)&int32(1984) | v2813
	v2862 = int32(2)
	goto L883
L890:
	;
	goto L889
L891:
	;
	v2846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2787+v2833))))
	v2861 = v2846&int32(63) | (v2804<<(uint(int32(18))%32)&int32(_a_F_dutch_UTF_8_stem_63) | v2813<<(uint(int32(12))%32) | v2829<<(uint(int32(6))%32))
	v2862 = int32(4)
	goto L883
L892:
	;
	v2833 = v2764 + int32(3)
	if v2833 != v2786 {
		goto L891
	} else {
		goto L895
	}
L893:
	;
	goto L894
L894:
	;
	v2861 = v2804<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_22) | v2813<<(uint(int32(6))%32) | v2829
	v2862 = int32(3)
	goto L883
L895:
	;
	goto L894
L896:
	;
	v2866 = v2861 - int32(97)
	if v2866 < int32(0) {
		v2884 = v2862
		goto L877
	} else {
		goto L897
	}
L897:
	;
	v2872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2866)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[0]))))
	if int32(base.Ui32(v2872)>>(uint(v2866&int32(7))%32))&int32(1) == int32(0) {
		v2884 = v2862
		goto L877
	} else {
		goto L898
	}
L898:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2862 + v2764
	goto L899
L899:
	;
	goto L879
L900:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2764
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L903
L901:
	;
	if v2948 < int32(0) {
		v3233 = v2601
		goto L816
	} else {
		goto L921
	}
L903:
	;
	goto L904
L904:
	;
	goto L905
L905:
	;
	v2903 = v2764
	v2905 = int32(1)
	goto L908
L907:
	;
	v2948 = v2933
	goto L901
L908:
	;
	if v2896 <= v2903 {
		goto L910
	} else {
		goto L911
	}
L909:
	;
	goto L907
L910:
	;
	v2948 = int32(-1)
	goto L901
L911:
	;
	goto L912
L912:
	;
	v2910 = v2903 + int32(1)
	v2912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2895+v2903))))
	if base.Ui32(v2912) < base.Ui32(int32(192)) {
		v2933 = v2910
		goto L913
	} else {
		goto L914
	}
L913:
	;
	v2934 = int32(1)
	if v2934 < v2905 {
		v2903 = v2933
		v2905 = v2905 - v2934
		goto L908
	} else {
		goto L920
	}
L914:
	;
	if v2896 <= v2910 {
		v2933 = v2910
		goto L913
	} else {
		goto L915
	}
L915:
	;
	v2919 = v2910
	goto L916
L916:
	;
	v2922 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2895+v2919))))
	if int32(-65) < v2922 {
		v2933 = v2919
		goto L913
	} else {
		goto L918
	}
L917:
	;
	v2933 = v2896
	goto L913
L918:
	;
	v2926 = v2919 + int32(1)
	if v2926 != v2896 {
		v2919 = v2926
		goto L916
	} else {
		goto L919
	}
L919:
	;
	goto L917
L920:
	;
	goto L909
L921:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2948
	v2952 = int32(2)
	v2954 = int32(0)
	v2956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2956-v2948 < v2952 {
		v2966 = v2954
		goto L923
	} else {
		goto L924
	}
L922:
	;
	if v2966 == int32(0) {
		v2764 = v2948
		goto L874
	} else {
		goto L926
	}
L923:
	;
	goto L922
L924:
	;
	v2960 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2962 = F_memcmp(m, v2960+v2948, int32(_a_F_dutch_UTF_8_stem_71), v2952)
	mBase = m.M
	if v2962 != 0 {
		v2966 = v2954
		goto L923
	} else {
		goto L925
	}
L925:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2952 + v2948
	v2966 = int32(1)
	goto L923
L926:
	;
	goto L875
L927:
	;
	v2989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2990 = int32(2)
	v2992 = int32(0)
	v2994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2994-v2989 < v2990 {
		v3004 = v2992
		goto L930
	} else {
		goto L931
	}
L928:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2989
	v3127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3127 <= v2989 {
		v3233 = v2601
		goto L816
	} else {
		goto L959
	}
L929:
	;
	if v3004 != 0 {
		goto L927
	} else {
		goto L933
	}
L930:
	;
	goto L929
L931:
	;
	v2998 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3000 = F_memcmp(m, v2998+v2989, int32(_a_F_dutch_UTF_8_stem_72), v2990)
	mBase = m.M
	if v3000 != 0 {
		v3004 = v2992
		goto L930
	} else {
		goto L932
	}
L932:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2990 + v2989
	v3004 = int32(1)
	goto L930
L933:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2989
	v3018 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3019 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L936
L934:
	;
	if v3123 == int32(0) {
		goto L927
	} else {
		goto L958
	}
L935:
	;
	v3123 = v3116
	goto L934
L936:
	;
	if v3018 <= v2989 {
		goto L938
	} else {
		goto L939
	}
L937:
	;
	v3116 = int32(0)
	goto L935
L938:
	;
	v3123 = int32(-1)
	goto L934
L939:
	;
	goto L940
L940:
	;
	v3034 = int32(1)
	v3036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2989+v3019))))
	if base.Ui32(v3036) < base.Ui32(int32(192)) {
		v3093 = v3036
		v3094 = v3034
		goto L941
	} else {
		goto L942
	}
L941:
	;
	if int32(252) < v3093 {
		v3116 = v3094
		goto L935
	} else {
		goto L954
	}
L942:
	;
	v3040 = v2989 + int32(1)
	if v3040 == v3018 {
		v3093 = v3036
		v3094 = v3034
		goto L941
	} else {
		goto L943
	}
L943:
	;
	v3043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3040+v3019))))
	v3045 = v3043 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v3036) {
		goto L945
	} else {
		goto L946
	}
L944:
	;
	v3059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3049+v3019))))
	v3061 = v3059 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v3036) {
		goto L950
	} else {
		goto L951
	}
L945:
	;
	v3049 = v2989 + int32(2)
	if v3049 != v3018 {
		goto L944
	} else {
		goto L948
	}
L946:
	;
	goto L947
L947:
	;
	v3093 = v3036<<(uint(int32(6))%32)&int32(1984) | v3045
	v3094 = int32(2)
	goto L941
L948:
	;
	goto L947
L949:
	;
	v3078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3019+v3065))))
	v3093 = v3078&int32(63) | (v3036<<(uint(int32(18))%32)&int32(_a_F_dutch_UTF_8_stem_63) | v3045<<(uint(int32(12))%32) | v3061<<(uint(int32(6))%32))
	v3094 = int32(4)
	goto L941
L950:
	;
	v3065 = v2989 + int32(3)
	if v3065 != v3018 {
		goto L949
	} else {
		goto L953
	}
L951:
	;
	goto L952
L952:
	;
	v3093 = v3036<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_22) | v3045<<(uint(int32(6))%32) | v3061
	v3094 = int32(3)
	goto L941
L953:
	;
	goto L952
L954:
	;
	v3098 = v3093 - int32(97)
	if v3098 < int32(0) {
		v3116 = v3094
		goto L935
	} else {
		goto L955
	}
L955:
	;
	v3104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v3098)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[0]))))
	if int32(base.Ui32(v3104)>>(uint(v3098&int32(7))%32))&int32(1) == int32(0) {
		v3116 = v3094
		goto L935
	} else {
		goto L956
	}
L956:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3094 + v2989
	goto L957
L957:
	;
	goto L937
L958:
	;
	goto L928
L959:
	;
	v3129 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v3129)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2689
	v3132 = F_slice_del(m, l0)
	mBase = m.M
	if v3132 < int32(0) {
		v3233 = v3132
		goto L816
	} else {
		goto L960
	}
L960:
	;
	v3135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3135
	v3138 = v3135 + int32(1)
	v3139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3139 <= v3138 {
		goto L961
	} else {
		goto L962
	}
L961:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3135
	v3233 = int32(1)
	goto L816
L962:
	;
	v3141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3141+v3138))))
	switch v3143 - int32(171) {
	case 0, 4:
		goto L963
	default:
		goto L961
	}
L963:
	;
	v3149 = F_find_among(m, l0, int32(_a_F_dutch_UTF_8_stem_73), int32(2), int32(0))
	mBase = m.M
	v3150 = m.ExcPending
	if v3150 != 0 {
		goto L31
	} else {
		goto L964
	}
L964:
	;
	if v3149 == int32(0) {
		goto L961
	} else {
		goto L965
	}
L965:
	;
	v3153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3153
	switch v3149 - int32(1) {
	case 0:
		goto L967
	case 1:
		goto L966
	default:
		goto L961
	}
L966:
	;
	v3165 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_74))
	mBase = m.M
	v3166 = m.ExcPending
	if v3166 != 0 {
		goto L31
	} else {
		goto L970
	}
L967:
	;
	v3159 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_75))
	mBase = m.M
	v3160 = m.ExcPending
	if v3160 != 0 {
		goto L31
	} else {
		goto L968
	}
L968:
	;
	if int32(0) <= v3159 {
		goto L961
	} else {
		goto L969
	}
L969:
	;
	v3233 = v3159
	goto L816
L970:
	;
	if v3165 < int32(0) {
		v3233 = v3165
		goto L816
	} else {
		goto L971
	}
L971:
	;
	goto L961
L972:
	;
	if int32(0) <= v3226 {
		v2663 = v3226
		goto L838
	} else {
		goto L992
	}
L974:
	;
	goto L975
L975:
	;
	goto L976
L976:
	;
	v3181 = v3173
	v3183 = int32(1)
	goto L979
L978:
	;
	v3226 = v3211
	goto L972
L979:
	;
	if v3174 <= v3181 {
		goto L981
	} else {
		goto L982
	}
L980:
	;
	goto L978
L981:
	;
	v3226 = int32(-1)
	goto L972
L982:
	;
	goto L983
L983:
	;
	v3188 = v3181 + int32(1)
	v3190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3172+v3181))))
	if base.Ui32(v3190) < base.Ui32(int32(192)) {
		v3211 = v3188
		goto L984
	} else {
		goto L985
	}
L984:
	;
	v3212 = int32(1)
	if v3212 < v3183 {
		v3181 = v3211
		v3183 = v3183 - v3212
		goto L979
	} else {
		goto L991
	}
L985:
	;
	if v3174 <= v3188 {
		v3211 = v3188
		goto L984
	} else {
		goto L986
	}
L986:
	;
	v3197 = v3188
	goto L987
L987:
	;
	v3200 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3172+v3197))))
	if int32(-65) < v3200 {
		v3211 = v3197
		goto L984
	} else {
		goto L989
	}
L988:
	;
	v3211 = v3174
	goto L984
L989:
	;
	v3204 = v3197 + int32(1)
	if v3204 != v3174 {
		v3197 = v3204
		goto L987
	} else {
		goto L990
	}
L990:
	;
	goto L988
L991:
	;
	goto L980
L992:
	;
	goto L839
L993:
	;
	if v3233 < int32(0) {
		v3602 = v3233
		goto L27
	} else {
		goto L996
	}
L994:
	;
	goto L995
L995:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2599
	v3335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3335
	v3337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v3337 != 0 {
		goto L1023
	} else {
		goto L1024
	}
L996:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2599
	v3245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v3245
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v3245
	goto L998
L997:
	;
	goto L995
L998:
	;
	v3256 = int32(0)
	v3257 = F_out_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), v3256)
	mBase = m.M
	if v3257 == v3256 {
		goto L998
	} else {
		goto L1000
	}
L999:
	;
	v3262 = int32(1)
	goto L1001
L1000:
	;
	goto L999
L1001:
	;
	v3265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3268 = F_eq_s(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_1))
	mBase = m.M
	if v3268 == int32(0) {
		goto L1004
	} else {
		goto L1005
	}
L1002:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3265
	if int32(0) < v3262 {
		goto L1008
	} else {
		goto L1009
	}
L1003:
	;
	goto L1002
L1004:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3265
	v3276 = F_in_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v3276 != 0 {
		goto L1003
	} else {
		goto L1007
	}
L1005:
	;
	goto L1006
L1006:
	;
	v3262 = v3262 - int32(1)
	goto L1001
L1007:
	;
	goto L1006
L1008:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2599
	goto L997
L1009:
	;
	v3286 = F_out_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v3286 != 0 {
		goto L1008
	} else {
		goto L1010
	}
L1010:
	;
	v3287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v3287
	goto L1011
L1011:
	;
	v3296 = int32(0)
	v3297 = F_out_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), v3296)
	mBase = m.M
	if v3297 == v3296 {
		goto L1011
	} else {
		goto L1013
	}
L1012:
	;
	v3302 = int32(1)
	goto L1014
L1013:
	;
	goto L1012
L1014:
	;
	v3305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3308 = F_eq_s(m, l0, int32(2), int32(_a_F_dutch_UTF_8_stem_2))
	mBase = m.M
	if v3308 == int32(0) {
		goto L1017
	} else {
		goto L1018
	}
L1015:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3305
	if int32(0) < v3302 {
		goto L1008
	} else {
		goto L1021
	}
L1016:
	;
	goto L1015
L1017:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3305
	v3316 = F_in_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v3316 != 0 {
		goto L1016
	} else {
		goto L1020
	}
L1018:
	;
	goto L1019
L1019:
	;
	v3302 = v3302 - int32(1)
	goto L1014
L1020:
	;
	goto L1019
L1021:
	;
	v3326 = F_out_grouping_U(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(97), int32(252), int32(0))
	mBase = m.M
	if v3326 != 0 {
		goto L1008
	} else {
		goto L1022
	}
L1022:
	;
	v3327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v3327
	goto L1008
L1023:
	;
	v3338 = F_r_Step_1c_2(m, l0)
	mBase = m.M
	v3339 = m.ExcPending
	if v3339 != 0 {
		goto L31
	} else {
		goto L1026
	}
L1024:
	;
	v3347 = v3335
	v3348 = v2598
	v3350 = v2600
	goto L1025
L1025:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3347
	v3352 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3347
	v3356 = v3347 - int32(1)
	v3357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3356 <= v3357 {
		v3396 = v3352
		goto L1036
	} else {
		goto L1037
	}
L1026:
	;
	v3341 = base.B2i32(v3338 < int32(0))
	if v3338 < int32(0) {
		goto L1027
	} else {
		goto L1028
	}
L1027:
	;
	v3342 = v3338
	goto L1029
L1028:
	;
	v3342 = v2600
	goto L1029
L1029:
	;
	if v3338 != 0 {
		goto L1030
	} else {
		goto L1031
	}
L1030:
	;
	v3343 = v3342
	goto L1032
L1031:
	;
	v3343 = v2600
	goto L1032
L1032:
	;
	if v3338 < int32(0) {
		goto L1033
	} else {
		goto L1034
	}
L1033:
	;
	return v3343
L1034:
	;
	goto L1035
L1035:
	;
	v3345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3347 = v3345
	v3348 = int32(1)
	v3350 = v3343
	goto L1025
L1036:
	;
	v3398 = base.B2i32(v3396 < int32(0))
	if v3396 < int32(0) {
		goto L1051
	} else {
		goto L1052
	}
L1037:
	;
	v3359 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3359+v3356))))
	if v3361 != int32(116) {
		v3396 = v3352
		goto L1036
	} else {
		goto L1038
	}
L1038:
	;
	v3367 = F_find_among_b(m, l0, int32(_a_F_dutch_UTF_8_stem_76), int32(3), int32(0))
	mBase = m.M
	v3368 = m.ExcPending
	if v3368 != 0 {
		goto L31
	} else {
		goto L1039
	}
L1039:
	;
	if v3367 == int32(0) {
		v3396 = v3352
		goto L1036
	} else {
		goto L1040
	}
L1040:
	;
	v3371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3371
	switch v3367 - int32(1) {
	case 0:
		goto L1044
	case 1:
		goto L1043
	case 2:
		goto L1042
	default:
		goto L1041
	}
L1041:
	;
	v3396 = int32(1)
	goto L1036
L1042:
	;
	v3389 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_77))
	mBase = m.M
	v3390 = m.ExcPending
	if v3390 != 0 {
		goto L31
	} else {
		goto L1049
	}
L1043:
	;
	v3383 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_78))
	mBase = m.M
	v3384 = m.ExcPending
	if v3384 != 0 {
		goto L31
	} else {
		goto L1047
	}
L1044:
	;
	v3377 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_79))
	mBase = m.M
	v3378 = m.ExcPending
	if v3378 != 0 {
		goto L31
	} else {
		goto L1045
	}
L1045:
	;
	if int32(0) <= v3377 {
		goto L1041
	} else {
		goto L1046
	}
L1046:
	;
	v3396 = v3377
	goto L1036
L1047:
	;
	if int32(0) <= v3383 {
		goto L1041
	} else {
		goto L1048
	}
L1048:
	;
	v3396 = v3383
	goto L1036
L1049:
	;
	if v3389 < int32(0) {
		v3396 = v3389
		goto L1036
	} else {
		goto L1050
	}
L1050:
	;
	goto L1041
L1051:
	;
	v3399 = v3396
	goto L1053
L1052:
	;
	v3399 = v3350
	goto L1053
L1053:
	;
	if v3396 != 0 {
		goto L1054
	} else {
		goto L1055
	}
L1054:
	;
	v3400 = v3399
	goto L1056
L1055:
	;
	v3400 = v3350
	goto L1056
L1056:
	;
	if v3396 != 0 {
		goto L1061
	} else {
		goto L1062
	}
L1057:
	;
	if v3409 == int32(0) {
		goto L1065
	} else {
		goto L1066
	}
L1058:
	;
	if v3396 < int32(0) {
		v3602 = v3400
		goto L27
	} else {
		goto L1064
	}
L1059:
	;
	v3407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3407
	v3409 = v3406
	goto L1057
L1060:
	;
	v3406 = int32(1)
	goto L1059
L1061:
	;
	v3404 = int32(base.Ui32(v3396) >> (uint(int32(31)) % 32))
	goto L1063
L1062:
	;
	v3404 = int32(10)
	goto L1063
L1063:
	;
	switch v3404 {
	case 0:
		goto L1060
	default:
		goto L1058
	case 10:
		v3406 = v3348
		goto L1059
	}
L1064:
	;
	v3409 = v3348
	goto L1057
L1065:
	;
	v3598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3598
	v3602 = int32(1)
	goto L27
L1066:
	;
	v3412 = int32(0)
	v3413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3413
	v3415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3413 <= v3415 {
		v3584 = v3412
		goto L1067
	} else {
		goto L1068
	}
L1067:
	;
	v3587 = int32(0)
	v3588 = base.B2i32(v3584 < v3587)
	if v3588 == v3587 {
		goto L1065
	} else {
		goto L1137
	}
L1068:
	;
	v3417 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3419 = int32(1)
	v3421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3417+v3413-v3419))))
	if base.B2i32(v3421&int32(224) != int32(96))|base.B2i32(v3419<<(uint(v3421)%32)&int32(98532828) == int32(0)) != 0 {
		v3584 = v3412
		goto L1067
	} else {
		goto L1069
	}
L1069:
	;
	v3436 = F_find_among_b(m, l0, int32(_a_F_dutch_UTF_8_stem_80), int32(22), int32(0))
	mBase = m.M
	v3437 = m.ExcPending
	if v3437 != 0 {
		goto L31
	} else {
		goto L1070
	}
L1070:
	;
	if v3436 == int32(0) {
		v3584 = v3412
		goto L1067
	} else {
		goto L1071
	}
L1071:
	;
	v3440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3440
	switch v3436 - int32(1) {
	case 0:
		goto L1092
	case 1:
		goto L1091
	case 2:
		goto L1090
	case 3:
		goto L1089
	case 4:
		goto L1088
	case 5:
		goto L1087
	case 6:
		goto L1086
	case 7:
		goto L1085
	case 8:
		goto L1084
	case 9:
		goto L1083
	case 10:
		goto L1082
	case 11:
		goto L1081
	case 12:
		goto L1080
	case 13:
		goto L1079
	case 14:
		goto L1078
	case 15:
		goto L1077
	case 16:
		goto L1076
	case 17:
		goto L1075
	case 18:
		goto L1074
	case 19:
		goto L1073
	default:
		goto L1072
	}
L1072:
	;
	v3584 = int32(1)
	goto L1067
L1073:
	;
	v3575 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_81))
	mBase = m.M
	v3576 = m.ExcPending
	if v3576 != 0 {
		goto L31
	} else {
		goto L1135
	}
L1074:
	;
	v3569 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_82))
	mBase = m.M
	v3570 = m.ExcPending
	if v3570 != 0 {
		goto L31
	} else {
		goto L1133
	}
L1075:
	;
	v3563 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_83))
	mBase = m.M
	v3564 = m.ExcPending
	if v3564 != 0 {
		goto L31
	} else {
		goto L1131
	}
L1076:
	;
	v3557 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_84))
	mBase = m.M
	v3558 = m.ExcPending
	if v3558 != 0 {
		goto L31
	} else {
		goto L1129
	}
L1077:
	;
	v3551 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_85))
	mBase = m.M
	v3552 = m.ExcPending
	if v3552 != 0 {
		goto L31
	} else {
		goto L1127
	}
L1078:
	;
	v3545 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_86))
	mBase = m.M
	v3546 = m.ExcPending
	if v3546 != 0 {
		goto L31
	} else {
		goto L1125
	}
L1079:
	;
	v3539 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_87))
	mBase = m.M
	v3540 = m.ExcPending
	if v3540 != 0 {
		goto L31
	} else {
		goto L1123
	}
L1080:
	;
	v3533 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_88))
	mBase = m.M
	v3534 = m.ExcPending
	if v3534 != 0 {
		goto L31
	} else {
		goto L1121
	}
L1081:
	;
	v3527 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_89))
	mBase = m.M
	v3528 = m.ExcPending
	if v3528 != 0 {
		goto L31
	} else {
		goto L1119
	}
L1082:
	;
	v3504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3440 <= v3504 {
		goto L1113
	} else {
		goto L1114
	}
L1083:
	;
	v3500 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_90))
	mBase = m.M
	v3501 = m.ExcPending
	if v3501 != 0 {
		goto L31
	} else {
		goto L1111
	}
L1084:
	;
	v3494 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_91))
	mBase = m.M
	v3495 = m.ExcPending
	if v3495 != 0 {
		goto L31
	} else {
		goto L1109
	}
L1085:
	;
	v3488 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_92))
	mBase = m.M
	v3489 = m.ExcPending
	if v3489 != 0 {
		goto L31
	} else {
		goto L1107
	}
L1086:
	;
	v3482 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_93))
	mBase = m.M
	v3483 = m.ExcPending
	if v3483 != 0 {
		goto L31
	} else {
		goto L1105
	}
L1087:
	;
	v3476 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_94))
	mBase = m.M
	v3477 = m.ExcPending
	if v3477 != 0 {
		goto L31
	} else {
		goto L1103
	}
L1088:
	;
	v3470 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_95))
	mBase = m.M
	v3471 = m.ExcPending
	if v3471 != 0 {
		goto L31
	} else {
		goto L1101
	}
L1089:
	;
	v3464 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_96))
	mBase = m.M
	v3465 = m.ExcPending
	if v3465 != 0 {
		goto L31
	} else {
		goto L1099
	}
L1090:
	;
	v3458 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_97))
	mBase = m.M
	v3459 = m.ExcPending
	if v3459 != 0 {
		goto L31
	} else {
		goto L1097
	}
L1091:
	;
	v3452 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_98))
	mBase = m.M
	v3453 = m.ExcPending
	if v3453 != 0 {
		goto L31
	} else {
		goto L1095
	}
L1092:
	;
	v3446 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_99))
	mBase = m.M
	v3447 = m.ExcPending
	if v3447 != 0 {
		goto L31
	} else {
		goto L1093
	}
L1093:
	;
	if int32(0) <= v3446 {
		goto L1072
	} else {
		goto L1094
	}
L1094:
	;
	v3584 = v3446
	goto L1067
L1095:
	;
	if int32(0) <= v3452 {
		goto L1072
	} else {
		goto L1096
	}
L1096:
	;
	v3584 = v3452
	goto L1067
L1097:
	;
	if int32(0) <= v3458 {
		goto L1072
	} else {
		goto L1098
	}
L1098:
	;
	v3584 = v3458
	goto L1067
L1099:
	;
	if int32(0) <= v3464 {
		goto L1072
	} else {
		goto L1100
	}
L1100:
	;
	v3584 = v3464
	goto L1067
L1101:
	;
	if int32(0) <= v3470 {
		goto L1072
	} else {
		goto L1102
	}
L1102:
	;
	v3584 = v3470
	goto L1067
L1103:
	;
	if int32(0) <= v3476 {
		goto L1072
	} else {
		goto L1104
	}
L1104:
	;
	v3584 = v3476
	goto L1067
L1105:
	;
	if int32(0) <= v3482 {
		goto L1072
	} else {
		goto L1106
	}
L1106:
	;
	v3584 = v3482
	goto L1067
L1107:
	;
	if int32(0) <= v3488 {
		goto L1072
	} else {
		goto L1108
	}
L1108:
	;
	v3584 = v3488
	goto L1067
L1109:
	;
	if int32(0) <= v3494 {
		goto L1072
	} else {
		goto L1110
	}
L1110:
	;
	v3584 = v3494
	goto L1067
L1111:
	;
	if int32(0) <= v3500 {
		goto L1072
	} else {
		goto L1112
	}
L1112:
	;
	v3584 = v3500
	goto L1067
L1113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3440
	v3521 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_100))
	mBase = m.M
	v3522 = m.ExcPending
	if v3522 != 0 {
		goto L31
	} else {
		goto L1117
	}
L1114:
	;
	v3506 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3506+v3440-int32(1)))))
	if v3510 != int32(105) {
		goto L1113
	} else {
		goto L1115
	}
L1115:
	;
	v3514 = v3440 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3514
	if v3514 <= v3504 {
		v3584 = v3412
		goto L1067
	} else {
		goto L1116
	}
L1116:
	;
	goto L1113
L1117:
	;
	if int32(0) <= v3521 {
		goto L1072
	} else {
		goto L1118
	}
L1118:
	;
	v3584 = v3521
	goto L1067
L1119:
	;
	if int32(0) <= v3527 {
		goto L1072
	} else {
		goto L1120
	}
L1120:
	;
	v3584 = v3527
	goto L1067
L1121:
	;
	if int32(0) <= v3533 {
		goto L1072
	} else {
		goto L1122
	}
L1122:
	;
	v3584 = v3533
	goto L1067
L1123:
	;
	if int32(0) <= v3539 {
		goto L1072
	} else {
		goto L1124
	}
L1124:
	;
	v3584 = v3539
	goto L1067
L1125:
	;
	if int32(0) <= v3545 {
		goto L1072
	} else {
		goto L1126
	}
L1126:
	;
	v3584 = v3545
	goto L1067
L1127:
	;
	if int32(0) <= v3551 {
		goto L1072
	} else {
		goto L1128
	}
L1128:
	;
	v3584 = v3551
	goto L1067
L1129:
	;
	if int32(0) <= v3557 {
		goto L1072
	} else {
		goto L1130
	}
L1130:
	;
	v3584 = v3557
	goto L1067
L1131:
	;
	if int32(0) <= v3563 {
		goto L1072
	} else {
		goto L1132
	}
L1132:
	;
	v3584 = v3563
	goto L1067
L1133:
	;
	if int32(0) <= v3569 {
		goto L1072
	} else {
		goto L1134
	}
L1134:
	;
	v3584 = v3569
	goto L1067
L1135:
	;
	if v3575 < int32(0) {
		v3584 = v3575
		goto L1067
	} else {
		goto L1136
	}
L1136:
	;
	goto L1072
L1137:
	;
	if v3584 < v3587 {
		goto L1138
	} else {
		goto L1139
	}
L1138:
	;
	v3591 = v3584
	goto L1140
L1139:
	;
	v3591 = v3400
	goto L1140
L1140:
	;
	if v3584 != 0 {
		goto L1141
	} else {
		goto L1142
	}
L1141:
	;
	v3592 = v3591
	goto L1143
L1142:
	;
	v3592 = v3400
	goto L1143
L1143:
	;
	return v3592
}
