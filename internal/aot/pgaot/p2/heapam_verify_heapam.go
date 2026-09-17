package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_verify_heapam(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int64
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int64
	_ = v310
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int64
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int64
	_ = v378
	var v386 int64
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int64
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int64
	_ = v405
	var v412 int64
	_ = v412
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v741 int32
	_ = v741
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int64
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int64
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int64
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int64
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
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
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int64
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int64
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int64
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int64
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int64
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int64
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1193 int32
	_ = v1193
	var v1199 int32
	_ = v1199
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int64
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1250 int32
	_ = v1250
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int64
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int64
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int64
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int64
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int64
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1499 int64
	_ = v1499
	var v1502 int64
	_ = v1502
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int64
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1541 int64
	_ = v1541
	var v1544 int64
	_ = v1544
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int64
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1597 int32
	_ = v1597
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1612 int64
	_ = v1612
	var v1615 int64
	_ = v1615
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1627 int64
	_ = v1627
	var v1630 int64
	_ = v1630
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1642 int64
	_ = v1642
	var v1645 int64
	_ = v1645
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1694 int64
	_ = v1694
	var v1697 int64
	_ = v1697
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1709 int64
	_ = v1709
	var v1712 int64
	_ = v1712
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1724 int64
	_ = v1724
	var v1727 int64
	_ = v1727
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1750 int32
	_ = v1750
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1767 int32
	_ = v1767
	var v1770 int32
	_ = v1770
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1792 int32
	_ = v1792
	var v1796 int32
	_ = v1796
	var v1803 int32
	_ = v1803
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1842 int32
	_ = v1842
	var v1851 int32
	_ = v1851
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1869 int32
	_ = v1869
	var v1871 int64
	_ = v1871
	var v1874 int64
	_ = v1874
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1884 int32
	_ = v1884
	var v1886 int64
	_ = v1886
	var v1889 int64
	_ = v1889
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1899 int32
	_ = v1899
	var v1901 int64
	_ = v1901
	var v1904 int64
	_ = v1904
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1930 int32
	_ = v1930
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1944 int64
	_ = v1944
	var v1947 int64
	_ = v1947
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1959 int64
	_ = v1959
	var v1962 int64
	_ = v1962
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1974 int64
	_ = v1974
	var v1977 int64
	_ = v1977
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1991 int32
	_ = v1991
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2008 int64
	_ = v2008
	var v2011 int64
	_ = v2011
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int64
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2050 int32
	_ = v2050
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2058 int64
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2075 int32
	_ = v2075
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2106 int64
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2123 int32
	_ = v2123
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2134 int32
	_ = v2134
	var v2144 int32
	_ = v2144
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2186 int32
	_ = v2186
	var v2187 int64
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2204 int32
	_ = v2204
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2221 int32
	_ = v2221
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2233 int32
	_ = v2233
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2244 int32
	_ = v2244
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2264 int64
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2275 int32
	_ = v2275
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2281 int32
	_ = v2281
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2293 int32
	_ = v2293
	var v2296 int32
	_ = v2296
	var v2304 int32
	_ = v2304
	var v2306 int32
	_ = v2306
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2326 int64
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2337 int32
	_ = v2337
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2343 int32
	_ = v2343
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2354 int32
	_ = v2354
	var v2358 int32
	_ = v2358
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2380 int64
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2391 int32
	_ = v2391
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2397 int32
	_ = v2397
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2408 int32
	_ = v2408
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2428 int64
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2439 int32
	_ = v2439
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2445 int32
	_ = v2445
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2465 int32
	_ = v2465
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2482 int64
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2493 int32
	_ = v2493
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2499 int32
	_ = v2499
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2527 int64
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2538 int32
	_ = v2538
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2544 int32
	_ = v2544
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2557 int32
	_ = v2557
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2573 int32
	_ = v2573
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2581 int64
	_ = v2581
	var v2583 int64
	_ = v2583
	var v2585 int32
	_ = v2585
	var v2587 int32
	_ = v2587
	var v2589 int32
	_ = v2589
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2605 int32
	_ = v2605
	var v2607 int32
	_ = v2607
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2633 int32
	_ = v2633
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2658 int32
	_ = v2658
	var v2661 int32
	_ = v2661
	var v2664 int32
	_ = v2664
	var v2670 int32
	_ = v2670
	var v2699 int32
	_ = v2699
	var v2701 int32
	_ = v2701
	var v2706 int32
	_ = v2706
	var v2709 int32
	_ = v2709
	var v2713 int32
	_ = v2713
	var v2735 int32
	_ = v2735
	var v2739 int32
	_ = v2739
	var v2745 int32
	_ = v2745
	var v2748 int32
	_ = v2748
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2758 int32
	_ = v2758
	var v2766 int32
	_ = v2766
	var v2773 int32
	_ = v2773
	var v2774 int32
	_ = v2774
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2778 int64
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2782 int32
	_ = v2782
	var v2783 int32
	_ = v2783
	var v2789 int32
	_ = v2789
	var v2791 int32
	_ = v2791
	var v2792 int32
	_ = v2792
	var v2795 int32
	_ = v2795
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2803 int32
	_ = v2803
	var v2804 int32
	_ = v2804
	var v2814 int32
	_ = v2814
	var v2815 int32
	_ = v2815
	var v2821 int32
	_ = v2821
	var v2822 int32
	_ = v2822
	var v2823 int32
	_ = v2823
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2840 int32
	_ = v2840
	var v2841 int32
	_ = v2841
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2844 int32
	_ = v2844
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2856 int32
	_ = v2856
	var v2858 int32
	_ = v2858
	var v2865 int32
	_ = v2865
	var v2866 int32
	_ = v2866
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2876 int32
	_ = v2876
	var v2881 int32
	_ = v2881
	var v2888 int32
	_ = v2888
	var v2889 int32
	_ = v2889
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2893 int64
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2897 int32
	_ = v2897
	var v2898 int32
	_ = v2898
	var v2904 int32
	_ = v2904
	var v2906 int32
	_ = v2906
	var v2907 int32
	_ = v2907
	var v2910 int32
	_ = v2910
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2918 int32
	_ = v2918
	var v2919 int32
	_ = v2919
	var v2921 int32
	_ = v2921
	var v2930 int32
	_ = v2930
	var v2937 int32
	_ = v2937
	var v2938 int32
	_ = v2938
	var v2939 int32
	_ = v2939
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2942 int64
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2946 int32
	_ = v2946
	var v2947 int32
	_ = v2947
	var v2953 int32
	_ = v2953
	var v2955 int32
	_ = v2955
	var v2956 int32
	_ = v2956
	var v2959 int32
	_ = v2959
	var v2964 int32
	_ = v2964
	var v2965 int32
	_ = v2965
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2977 int32
	_ = v2977
	var v2978 int32
	_ = v2978
	var v2982 int32
	_ = v2982
	var v2983 int32
	_ = v2983
	var v2984 int32
	_ = v2984
	var v2986 int32
	_ = v2986
	var v2988 int32
	_ = v2988
	var v2992 int32
	_ = v2992
	var v2993 int32
	_ = v2993
	var v2996 int32
	_ = v2996
	var v3000 int32
	_ = v3000
	var v3004 int32
	_ = v3004
	var v3005 int32
	_ = v3005
	var v3006 int32
	_ = v3006
	var v3007 int32
	_ = v3007
	var v3016 int32
	_ = v3016
	var v3017 int32
	_ = v3017
	var v3018 int32
	_ = v3018
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3021 int64
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3032 int32
	_ = v3032
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3038 int32
	_ = v3038
	var v3043 int32
	_ = v3043
	var v3044 int32
	_ = v3044
	var v3046 int32
	_ = v3046
	var v3047 int32
	_ = v3047
	var v3049 int32
	_ = v3049
	var v3050 int32
	_ = v3050
	var v3056 int32
	_ = v3056
	var v3058 int32
	_ = v3058
	var v3060 int32
	_ = v3060
	var v3064 int32
	_ = v3064
	var v3068 int32
	_ = v3068
	var v3072 int32
	_ = v3072
	var v3076 int32
	_ = v3076
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3093 int32
	_ = v3093
	var v3102 int32
	_ = v3102
	var v3103 int32
	_ = v3103
	var v3104 int32
	_ = v3104
	var v3105 int64
	_ = v3105
	var v3106 int32
	_ = v3106
	var v3109 int32
	_ = v3109
	var v3110 int32
	_ = v3110
	var v3116 int32
	_ = v3116
	var v3118 int32
	_ = v3118
	var v3119 int32
	_ = v3119
	var v3122 int32
	_ = v3122
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3130 int32
	_ = v3130
	var v3131 int32
	_ = v3131
	var v3143 int32
	_ = v3143
	var v3145 int32
	_ = v3145
	var v3150 int32
	_ = v3150
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3178 int32
	_ = v3178
	var v3182 int32
	_ = v3182
	var v3186 int32
	_ = v3186
	var v3192 int32
	_ = v3192
	var v3193 int32
	_ = v3193
	var v3195 int32
	_ = v3195
	var v3203 int32
	_ = v3203
	var v3208 int32
	_ = v3208
	var v3209 int32
	_ = v3209
	var v3210 int32
	_ = v3210
	var v3211 int32
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3213 int64
	_ = v3213
	var v3214 int32
	_ = v3214
	var v3217 int32
	_ = v3217
	var v3218 int32
	_ = v3218
	var v3224 int32
	_ = v3224
	var v3226 int32
	_ = v3226
	var v3227 int32
	_ = v3227
	var v3230 int32
	_ = v3230
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3238 int32
	_ = v3238
	var v3239 int32
	_ = v3239
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3248 int32
	_ = v3248
	var v3251 int32
	_ = v3251
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3277 int32
	_ = v3277
	var v3278 int32
	_ = v3278
	var v3289 int32
	_ = v3289
	var v3303 int32
	_ = v3303
	var v3307 int32
	_ = v3307
	var v3308 int32
	_ = v3308
	var v3310 int32
	_ = v3310
	var v3314 int32
	_ = v3314
	var v3316 int32
	_ = v3316
	var v3318 int32
	_ = v3318
	var v3322 int32
	_ = v3322
	var v3323 int32
	_ = v3323
	var v3324 int32
	_ = v3324
	var v3325 int32
	_ = v3325
	var v3326 int32
	_ = v3326
	var v3328 int32
	_ = v3328
	var v3329 int32
	_ = v3329
	var v3331 int32
	_ = v3331
	var v3332 int32
	_ = v3332
	var v3340 int32
	_ = v3340
	var v3342 int32
	_ = v3342
	var v3359 int32
	_ = v3359
	var v3360 int32
	_ = v3360
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3365 int32
	_ = v3365
	var v3368 int32
	_ = v3368
	var v3373 int32
	_ = v3373
	var v3374 int32
	_ = v3374
	var v3376 int32
	_ = v3376
	var v3383 int32
	_ = v3383
	var v3384 int32
	_ = v3384
	var v3385 int32
	_ = v3385
	var v3386 int64
	_ = v3386
	var v3387 int32
	_ = v3387
	var v3388 int32
	_ = v3388
	var v3389 int32
	_ = v3389
	var v3392 int32
	_ = v3392
	var v3393 int32
	_ = v3393
	var v3399 int32
	_ = v3399
	var v3401 int32
	_ = v3401
	var v3402 int32
	_ = v3402
	var v3405 int32
	_ = v3405
	var v3410 int32
	_ = v3410
	var v3411 int32
	_ = v3411
	var v3413 int32
	_ = v3413
	var v3414 int32
	_ = v3414
	var v3423 int32
	_ = v3423
	var v3425 int32
	_ = v3425
	var v3426 int32
	_ = v3426
	var v3429 int32
	_ = v3429
	var v3430 int32
	_ = v3430
	var v3431 int32
	_ = v3431
	var v3434 int32
	_ = v3434
	var v3440 int32
	_ = v3440
	var v3441 int32
	_ = v3441
	var v3442 int32
	_ = v3442
	var v3447 int32
	_ = v3447
	var v3456 int32
	_ = v3456
	var v3460 int32
	_ = v3460
	var v3462 int32
	_ = v3462
	var v3469 int32
	_ = v3469
	var v3470 int32
	_ = v3470
	var v3473 int32
	_ = v3473
	var v3475 int32
	_ = v3475
	var v3483 int32
	_ = v3483
	var v3484 int32
	_ = v3484
	var v3485 int32
	_ = v3485
	var v3486 int32
	_ = v3486
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3496 int32
	_ = v3496
	var v3502 int32
	_ = v3502
	var v3503 int64
	_ = v3503
	var v3504 int32
	_ = v3504
	var v3505 int32
	_ = v3505
	var v3506 int32
	_ = v3506
	var v3509 int32
	_ = v3509
	var v3510 int32
	_ = v3510
	var v3516 int32
	_ = v3516
	var v3518 int32
	_ = v3518
	var v3519 int32
	_ = v3519
	var v3522 int32
	_ = v3522
	var v3527 int32
	_ = v3527
	var v3528 int32
	_ = v3528
	var v3530 int32
	_ = v3530
	var v3531 int32
	_ = v3531
	var v3535 int32
	_ = v3535
	var v3542 int32
	_ = v3542
	var v3543 int32
	_ = v3543
	var v3545 int32
	_ = v3545
	var v3547 int32
	_ = v3547
	var v3554 int32
	_ = v3554
	var v3555 int32
	_ = v3555
	var v3557 int32
	_ = v3557
	var v3558 int32
	_ = v3558
	var v3563 int32
	_ = v3563
	var v3564 int32
	_ = v3564
	var v3568 int32
	_ = v3568
	var v3586 int32
	_ = v3586
	var v3587 int64
	_ = v3587
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3590 int32
	_ = v3590
	var v3593 int32
	_ = v3593
	var v3594 int32
	_ = v3594
	var v3600 int32
	_ = v3600
	var v3602 int32
	_ = v3602
	var v3603 int32
	_ = v3603
	var v3606 int32
	_ = v3606
	var v3611 int32
	_ = v3611
	var v3612 int32
	_ = v3612
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3639 int32
	_ = v3639
	var v3640 int32
	_ = v3640
	var v3642 int32
	_ = v3642
	var v3664 int32
	_ = v3664
	var v3666 int32
	_ = v3666
	var v3692 int32
	_ = v3692
	var v3719 int32
	_ = v3719
	var v3720 int32
	_ = v3720
	var v3722 int32
	_ = v3722
	var v3723 int32
	_ = v3723
	var v3724 int32
	_ = v3724
	var v3726 int32
	_ = v3726
	var v3727 int32
	_ = v3727
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3734 int32
	_ = v3734
	var v3756 int32
	_ = v3756
	var v3766 int32
	_ = v3766
	var v3773 int32
	_ = v3773
	var v3778 int32
	_ = v3778
	v2 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(_a_F_verify_heapam_0)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[0]))) = v2
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v29 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3766 = m.ExcPending
	if v3766 != 0 {
		goto L19
	} else {
		goto L806
	}
L2:
	;
	v3756 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v3756)
	m.G0 = v24 + int32(_a_F_verify_heapam_0)
	return int32(0)
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[1]))) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[2]))) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[3]))) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[4]))) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[5]))) = v24 + int32(_a_F_verify_heapam_1)
	if v185 != 0 {
		goto L163
	} else {
		goto L164
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[6]))) = base.I64_extend_i32_u(v396)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v395)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[7]))) = v559
	goto L3
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L19
	} else {
		goto L159
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L19
	} else {
		goto L155
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L19
	} else {
		goto L151
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L19
	} else {
		goto L146
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L19
	} else {
		goto L142
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L19
	} else {
		goto L138
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L19
	} else {
		goto L134
	}
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v32 == int32(1) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L19
	} else {
		goto L130
	}
L15:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v35 == int32(1) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v38 == int32(1) {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v46 = F_pg_detoast_datum_packed(m, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	base.MemoryFill(m, v24+int32(_a_F_verify_heapam_2), int32(0), int32(144))
	v192 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L19
	} else {
		goto L66
	}
L19:
	;
	return int32(0)
L20:
	;
	v50 = F_text_to_cstring(m, v46)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v55 = v50
	v56 = int32(_a_F_verify_heapam_3)
	goto L23
L22:
	;
	if v93 == int32(0) {
		v185 = v2
		v186 = int32(1)
		goto L18
	} else {
		goto L35
	}
L23:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v59 == v60 {
		v82 = v59
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v93 = int32(0)
	goto L22
L25:
	;
	v84 = int32(1)
	if v82 != 0 {
		v55 = v55 + v84
		v56 = v56 + v84
		goto L23
	} else {
		goto L34
	}
L26:
	;
	if base.Ui32((v59-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v70 = v59 | int32(32)
	goto L29
L28:
	;
	v70 = v59
	goto L29
L29:
	;
	if base.Ui32((v60-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v79 = v60 | int32(32)
	goto L32
L31:
	;
	v79 = v60
	goto L32
L32:
	;
	if v70 == v79 {
		v82 = v70
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v93 = v70 - v79
	goto L22
L34:
	;
	goto L24
L35:
	;
	v99 = v50
	v100 = int32(_a_F_verify_heapam_4)
	goto L37
L36:
	;
	if v137 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L37:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v103 == v104 {
		v126 = v103
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v137 = int32(0)
	goto L36
L39:
	;
	v128 = int32(1)
	if v126 != 0 {
		v99 = v99 + v128
		v100 = v100 + v128
		goto L37
	} else {
		goto L48
	}
L40:
	;
	if base.Ui32((v103-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v114 = v103 | int32(32)
	goto L43
L42:
	;
	v114 = v103
	goto L43
L43:
	;
	if base.Ui32((v104-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v123 = v104 | int32(32)
	goto L46
L45:
	;
	v123 = v104
	goto L46
L46:
	;
	if v114 == v123 {
		v126 = v114
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v137 = v114 - v123
	goto L36
L48:
	;
	goto L38
L49:
	;
	v185 = v2
	v186 = int32(0)
	goto L18
L50:
	;
	goto L51
L51:
	;
	v144 = v50
	v145 = int32(_a_F_verify_heapam_5)
	goto L53
L52:
	;
	if v182 != 0 {
		goto L8
	} else {
		goto L65
	}
L53:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v148 == v149 {
		v171 = v148
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v182 = int32(0)
	goto L52
L55:
	;
	v173 = int32(1)
	if v171 != 0 {
		v144 = v144 + v173
		v145 = v145 + v173
		goto L53
	} else {
		goto L64
	}
L56:
	;
	if base.Ui32((v148-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v159 = v148 | int32(32)
	goto L59
L58:
	;
	v159 = v148
	goto L59
L59:
	;
	if base.Ui32((v149-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v168 = v149 | int32(32)
	goto L62
L61:
	;
	v168 = v149
	goto L62
L62:
	;
	if v159 == v168 {
		v171 = v159
		goto L55
	} else {
		goto L63
	}
L63:
	;
	v182 = v159 - v168
	goto L52
L64:
	;
	goto L54
L65:
	;
	v185 = int32(1)
	v186 = int32(2)
	goto L18
L66:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v195 = int32(_a_F_verify_heapam_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))) = uint16(v195)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[9]))) = v194
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L19
	} else {
		goto L67
	}
L67:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10]))) = v201
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11]))) = v203
	v206 = F_relation_open(m, v41, int32(1))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L19
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[12]))) = v206
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206)+48))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+119)))
	switch v210 - int32(83) {
	case 0:
		goto L69
	default:
		goto L71
	case 26, 31, 33:
		goto L70
	}
L69:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+118)))
	if v239 != int32(117) {
		goto L78
	} else {
		goto L79
	}
L70:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v209)+84))
	if v236 != int32(2) {
		goto L7
	} else {
		goto L77
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L19
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L19
	} else {
		goto L73
	}
L73:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v206)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v220 + int32(4)
	F_errmsg(m, int32(_a_F_verify_heapam_7), v24)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L19
	} else {
		goto L74
	}
L74:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v206)+48))
	v228 = int32(*(*int8)(unsafe.Add(mBase, uint32(v227)+119)))
	F_errdetail_relkind_not_supported(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L19
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(339), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L19
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
	goto L69
L78:
	;
	v281 = int32(0)
	v283 = F_RelationGetNumberOfBlocksInFork(m, v206, v281)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L19
	} else {
		goto L93
	}
L79:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_verify_heapam[13])))
	if v244 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v254 == int32(0) {
		goto L78
	} else {
		goto L84
	}
L81:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[14]))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+316))
	v252 = base.B2i32(v250 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_verify_heapam[13])) = uint8(v252)
	v254 = v252
	goto L83
L82:
	;
	v254 = int32(0)
	goto L83
L83:
	;
	goto L80
L84:
	;
	v259 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L19
	} else {
		goto L85
	}
L85:
	;
	if v259 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_errcode(m, int32(100663618))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L19
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	F_relation_close(m, v206, int32(1))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L19
	} else {
		goto L92
	}
L89:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v206)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v264 + int32(4)
	F_errmsg(m, int32(_a_F_verify_heapam_10), v24+int32(16))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L19
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(362), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L19
	} else {
		goto L91
	}
L91:
	;
	goto L88
L92:
	;
	goto L2
L93:
	;
	if v283 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	F_relation_close(m, v206, int32(1))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L19
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v291 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L19
	} else {
		goto L98
	}
L97:
	;
	goto L2
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[15]))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[16]))) = v291
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v296 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v299)))
	if base.Ui64(base.I64_extend_i32_u(v283)) <= base.Ui64(v300) {
		goto L6
	} else {
		goto L102
	}
L100:
	;
	v304 = v281
	goto L101
L101:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
	if v306 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v304 = base.I32_wrap_i64(v300)
	goto L101
L103:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v309)))
	if base.Ui64(base.I64_extend_i32_u(v283)) <= base.Ui64(v310) {
		goto L5
	} else {
		goto L106
	}
L104:
	;
	v316 = v283
	goto L105
L105:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v206)+48))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+112))
	v320 = int32(0)
	if base.B2i32(v319 == v320)|base.B2i32(v43 == v320) == v320 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v316 = base.I32_wrap_i64(v310) + int32(1)
	goto L105
L107:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[17]))
	v354 = F_LWLockAcquire(m, v350+int32(384), int32(1))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L19
	} else {
		goto L113
	}
L108:
	;
	v328 = F_table_open(m, v319, int32(1))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L19
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[18]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[19]))) = int64(0)
	goto L107
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[19]))) = v328
	v336 = F_toast_open_indexes(m, v328, int32(1), v24+int32(_a_F_verify_heapam_11), v24+int32(_a_F_verify_heapam_12))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L19
	} else {
		goto L112
	}
L112:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[20])))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v338+v336<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[21]))) = v342
	goto L107
L113:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[22]))
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v357)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[23]))) = v358
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v357)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[24]))) = v360
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[17]))
	F_LWLockRelease(m, v363+int32(384))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L19
	} else {
		goto L114
	}
L114:
	;
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[23])))
	v369 = base.I32_wrap_i64(v368)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[25]))) = v369
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[24])))
	if base.Ui32(v371) <= base.Ui32(int32(2)) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[26]))) = v386
	v389 = v24 + int32(_a_F_verify_heapam_13)
	v391 = v24 + int32(_a_F_verify_heapam_14)
	F_ReadMultiXactIdRange(m, v389, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L19
	} else {
		goto L123
	}
L116:
	;
	v386 = base.I64_extend_i32_u(v371)
	goto L115
L117:
	;
	goto L118
L118:
	;
	v375 = v369 - v371
	if int32(0) < v375 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v378 = int64(3)
	if base.Ui64(v368-v378) < base.Ui64(base.I64_extend_i32_u(v375)) {
		v386 = v378
		goto L115
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v386 = v368 - base.I64_extend_i32_s(v375)
	goto L115
L122:
	;
	goto L121
L123:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[12])))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+48))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[27]))) = v396
	if base.Ui32(v396) < base.Ui32(int32(3)) {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v400 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[23])))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[25])))
	v402 = v401 - v396
	if int32(0) < v402 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[6]))) = v412
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v395)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[24]))) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[7]))) = v414
	goto L3
L126:
	;
	v405 = int64(3)
	if base.Ui64(v400-v405) < base.Ui64(base.I64_extend_i32_u(v402)) {
		v412 = v405
		goto L125
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v412 = v400 - base.I64_extend_i32_s(v402)
	goto L125
L129:
	;
	goto L128
L130:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L19
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_verify_heapam_15), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L19
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(273), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L19
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L19
	} else {
		goto L135
	}
L135:
	;
	F_errmsg(m, int32(_a_F_verify_heapam_16), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L19
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(279), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L19
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L19
	} else {
		goto L139
	}
L139:
	;
	F_errmsg(m, int32(_a_F_verify_heapam_17), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L19
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(285), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L19
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L19
	} else {
		goto L143
	}
L143:
	;
	F_errmsg(m, int32(_a_F_verify_heapam_18), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L19
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(291), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L19
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L19
	} else {
		goto L147
	}
L147:
	;
	F_errmsg(m, int32(_a_F_verify_heapam_19), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L19
	} else {
		goto L148
	}
L148:
	;
	F_errhint(m, int32(_a_F_verify_heapam_20), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L19
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(303), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L19
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L19
	} else {
		goto L152
	}
L152:
	;
	F_errmsg(m, int32(_a_F_verify_heapam_21), int32(0))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L19
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(349), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L19
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L19
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+1024)) = v283 - int32(1)
	F_errmsg(m, int32(_a_F_verify_heapam_22), v24+int32(1024))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L19
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(390), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L19
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L19
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+1008)) = v283 - int32(1)
	F_errmsg(m, int32(_a_F_verify_heapam_23), v24+int32(1008))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L19
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(403), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L19
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	v573 = int32(14)
	goto L165
L164:
	;
	v573 = int32(0)
	goto L165
L165:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[16])))
	if v185 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v578 = int32(120)
	goto L168
L167:
	;
	v578 = int32(_a_F_verify_heapam_24)
	goto L168
L168:
	;
	v582 = F_read_stream_begin_relation(m, v573, v574, v394, int32(0), v578, v24+int32(_a_F_verify_heapam_25), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L19
	} else {
		goto L169
	}
L169:
	;
	goto L170
L170:
	;
	v606 = F_read_stream_next_buffer(m, v582, int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L19
	} else {
		goto L173
	}
L171:
	;
	F_read_stream_end(m, v582)
	mBase = m.M
	v3719 = m.ExcPending
	if v3719 != 0 {
		goto L19
	} else {
		goto L792
	}
L172:
	;
	goto L171
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[15]))) = v606
	if v606 == int32(0) {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v612 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[28]))
	if v612 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L19
	} else {
		goto L178
	}
L176:
	;
	v616 = v606
	goto L177
L177:
	;
	base.MemoryFill(m, v24+int32(_a_F_verify_heapam_26), int32(0), int32(_a_F_verify_heapam_27))
	F_LockBuffer(m, v616, int32(1))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L19
	} else {
		goto L179
	}
L178:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[15])))
	v616 = v615
	goto L177
L179:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[15])))
	if v625 < int32(0) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))) = v644
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[15])))
	if v646 < int32(0) {
		goto L185
	} else {
		goto L186
	}
L181:
	;
	v629 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[30]))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v629+(v625^int32(-1))<<(uint(int32(6))%32))+16))
	v644 = v635
	goto L180
L182:
	;
	goto L183
L183:
	;
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[31]))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v637+v625<<(uint(int32(6))%32)+int32(-64))+16))
	v644 = v643
	goto L180
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[32]))) = v664
	v666 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v664)+12)))
	v667 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))) = uint16(v667)
	if base.Ui32(int32(25)) <= base.Ui32(v666) {
		goto L189
	} else {
		goto L190
	}
L185:
	;
	v650 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[34]))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v650+(v646^int32(-1))<<(uint(int32(2))%32))))
	v664 = v656
	goto L184
L186:
	;
	goto L187
L187:
	;
	v658 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[35]))
	v664 = v658 + v646<<(uint(int32(13))%32) + int32(-8192)
	goto L184
L188:
	;
	v3274 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[15])))
	F_UnlockReleaseBuffer(m, v3274)
	mBase = m.M
	v3276 = m.ExcPending
	if v3276 != 0 {
		goto L19
	} else {
		goto L714
	}
L189:
	;
	v677 = int32(base.Ui32(v666+int32(_a_F_verify_heapam_28)) >> (uint(int32(2)) % 32))
	goto L191
L190:
	;
	v677 = int32(0)
	goto L191
L191:
	;
	v679 = v677 & int32(_a_F_verify_heapam_6)
	if v679 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v682 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))) = uint16(v682)
	v684 = int32(_a_F_verify_heapam_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))) = uint16(v684)
	goto L188
L193:
	;
	goto L194
L194:
	;
	v688 = v667
	goto L195
L195:
	;
	v708 = v688 & int32(_a_F_verify_heapam_6)
	v711 = v708 + (v24 + int32(_a_F_verify_heapam_29))
	v712 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v711))) = uint8(v712)
	v716 = v24 + int32(_a_F_verify_heapam_30) + v708
	*(*uint8)(unsafe.Add(mBase, uint32(v716))) = uint8(v712)
	v721 = int32(1)
	v723 = v24 + int32(_a_F_verify_heapam_31) + v708<<(uint(v721)%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v723))) = uint16(v712)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[32])))
	v728 = v726 + int32(20)
	v730 = v708 << (uint(int32(2)) % 32)
	v731 = v728 + v730
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[36]))) = v731
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	switch int32(base.Ui32(v733)>>(uint(int32(15))%32))&int32(3) - v721 {
	case 0:
		goto L198
	case 1:
		goto L199
	default:
		goto L197
	}
L196:
	;
	v2706 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))) = uint16(v2706)
	v2709 = int32(_a_F_verify_heapam_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))) = uint16(v2709)
	v2713 = v2706
	goto L615
L197:
	;
	v2699 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2701 = v2699 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))) = uint16(v2701)
	if base.Ui32(v2701&int32(_a_F_verify_heapam_6)) <= base.Ui32(v679) {
		v688 = v2701
		goto L195
	} else {
		goto L614
	}
L198:
	;
	v946 = int32(base.Ui32(v733) >> (uint(int32(17)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[37]))) = uint16(v946)
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v950 = v948 & int32(_a_F_verify_heapam_32)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[38]))) = uint16(v950)
	if (v950+int32(7))&int32(_a_F_verify_heapam_33) != v950 {
		goto L240
	} else {
		goto L241
	}
L199:
	;
	v741 = v733 & int32(_a_F_verify_heapam_32)
	if v741 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+288)) = int64(4294967296)
	v749 = F_psprintf(m, int32(_a_F_verify_heapam_34), v24+int32(288))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L19
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	if base.Ui32(v679) < base.Ui32(v741) {
		goto L209
	} else {
		goto L210
	}
L203:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v753 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v754 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v755 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v758 = F_Int64GetDatum(m, v754)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L19
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v755)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v753
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v758
	v765 = int32(base.Ui32(v755) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v765)
	v767 = F_cstring_to_text(m, v749)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L19
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v767
	F_pfree(m, v749)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L19
	} else {
		goto L206
	}
L206:
	;
	v776 = F_heap_form_tuple(m, v752, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L19
	} else {
		goto L207
	}
L207:
	;
	F_tuplestore_puttuple(m, v751, v776)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L19
	} else {
		goto L208
	}
L208:
	;
	v780 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v780)
	goto L197
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+308)) = v679
	*(*int32)(unsafe.Add(mBase, uint32(v24)+304)) = v741
	v788 = F_psprintf(m, int32(_a_F_verify_heapam_37), v24+int32(304))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L19
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v728+v741<<(uint(int32(2))%32))))
	switch int32(base.Ui32(v824)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L218
	case 1:
		goto L219
	case 2:
		goto L220
	default:
		goto L221
	}
L212:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v792 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v793 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v794 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v797 = F_Int64GetDatum(m, v793)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L19
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v794)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v792
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v797
	v804 = int32(base.Ui32(v794) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v804)
	v806 = F_cstring_to_text(m, v788)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L19
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v806
	F_pfree(m, v788)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L19
	} else {
		goto L215
	}
L215:
	;
	v815 = F_heap_form_tuple(m, v791, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L19
	} else {
		goto L216
	}
L216:
	;
	F_tuplestore_puttuple(m, v790, v815)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L19
	} else {
		goto L217
	}
L217:
	;
	v819 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v819)
	goto L197
L218:
	;
	v942 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v711))) = uint8(v942)
	*(*uint16)(unsafe.Add(mBase, uint32(v723))) = uint16(v741)
	goto L197
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+352)) = v741
	v909 = F_psprintf(m, int32(_a_F_verify_heapam_38), v24+int32(352))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L19
	} else {
		goto L234
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+336)) = v741
	v872 = F_psprintf(m, int32(_a_F_verify_heapam_39), v24+int32(336))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L19
	} else {
		goto L228
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+320)) = v741
	v835 = F_psprintf(m, int32(_a_F_verify_heapam_40), v24+int32(320))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L19
	} else {
		goto L222
	}
L222:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v839 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v840 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v841 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v844 = F_Int64GetDatum(m, v840)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L19
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v841)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v839
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v844
	v851 = int32(base.Ui32(v841) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v851)
	v853 = F_cstring_to_text(m, v835)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L19
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v853
	F_pfree(m, v835)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L19
	} else {
		goto L225
	}
L225:
	;
	v862 = F_heap_form_tuple(m, v838, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L19
	} else {
		goto L226
	}
L226:
	;
	F_tuplestore_puttuple(m, v837, v862)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L19
	} else {
		goto L227
	}
L227:
	;
	v866 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v866)
	goto L197
L228:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v876 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v877 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v878 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v881 = F_Int64GetDatum(m, v877)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L19
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v878)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v881
	v888 = int32(base.Ui32(v878) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v888)
	v890 = F_cstring_to_text(m, v872)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L19
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v890
	F_pfree(m, v872)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L19
	} else {
		goto L231
	}
L231:
	;
	v899 = F_heap_form_tuple(m, v875, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L19
	} else {
		goto L232
	}
L232:
	;
	F_tuplestore_puttuple(m, v874, v899)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L19
	} else {
		goto L233
	}
L233:
	;
	v903 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v903)
	goto L197
L234:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v913 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v914 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v915 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v918 = F_Int64GetDatum(m, v914)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L19
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v915)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v913
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v918
	v925 = int32(base.Ui32(v915) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v925)
	v927 = F_cstring_to_text(m, v909)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L19
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v927
	F_pfree(m, v909)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L19
	} else {
		goto L237
	}
L237:
	;
	v936 = F_heap_form_tuple(m, v912, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L19
	} else {
		goto L238
	}
L238:
	;
	F_tuplestore_puttuple(m, v911, v936)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L19
	} else {
		goto L239
	}
L239:
	;
	v940 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v940)
	goto L197
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+992)) = v950
	v961 = F_psprintf(m, int32(_a_F_verify_heapam_41), v24+int32(992))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L19
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	if base.Ui32(v733) <= base.Ui32(int32(_a_F_verify_heapam_42)) {
		goto L249
	} else {
		goto L250
	}
L243:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v965 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v966 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v967 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v970 = F_Int64GetDatum(m, v966)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L19
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v967)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v965
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v970
	v977 = int32(base.Ui32(v967) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v977)
	v979 = F_cstring_to_text(m, v961)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L19
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v979
	F_pfree(m, v961)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L19
	} else {
		goto L246
	}
L246:
	;
	v988 = F_heap_form_tuple(m, v964, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L19
	} else {
		goto L247
	}
L247:
	;
	F_tuplestore_puttuple(m, v963, v988)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L19
	} else {
		goto L248
	}
L248:
	;
	v992 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v992)
	goto L197
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+372)) = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+368)) = v946
	v1002 = F_psprintf(m, int32(_a_F_verify_heapam_43), v24+int32(368))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L19
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	if base.Ui32(int32(_a_F_verify_heapam_44)) <= base.Ui32(v950+v946) {
		goto L258
	} else {
		goto L259
	}
L252:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v1006 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v1007 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v1008 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v1011 = F_Int64GetDatum(m, v1007)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L19
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v1008)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v1011
	v1018 = int32(base.Ui32(v1008) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v1018)
	v1020 = F_cstring_to_text(m, v1002)
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L19
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v1020
	F_pfree(m, v1002)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L19
	} else {
		goto L255
	}
L255:
	;
	v1029 = F_heap_form_tuple(m, v1005, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L19
	} else {
		goto L256
	}
L256:
	;
	F_tuplestore_puttuple(m, v1004, v1029)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L19
	} else {
		goto L257
	}
L257:
	;
	v1033 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v1033)
	goto L197
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+392)) = int32(_a_F_verify_heapam_45)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+388)) = v946
	*(*int32)(unsafe.Add(mBase, uint32(v24)+384)) = v950
	v1045 = F_psprintf(m, int32(_a_F_verify_heapam_46), v24+int32(384))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L19
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v1078 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v711))) = uint8(v1078)
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v1083 = v726 + v1080&int32(_a_F_verify_heapam_32)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[46]))) = v1083
	v1085 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1083)+18)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[47]))) = v1085 & int32(2047)
	v1089 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1083)+20)))
	if v1089&int32(_a_F_verify_heapam_47) == int32(_a_F_verify_heapam_27) {
		goto L268
	} else {
		goto L269
	}
L261:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v1049 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v1050 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v1051 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v1054 = F_Int64GetDatum(m, v1050)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L19
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v1051)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v1049
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v1054
	v1061 = int32(base.Ui32(v1051) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v1061)
	v1063 = F_cstring_to_text(m, v1045)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L19
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v1063
	F_pfree(m, v1045)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L19
	} else {
		goto L264
	}
L264:
	;
	v1072 = F_heap_form_tuple(m, v1048, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L19
	} else {
		goto L265
	}
L265:
	;
	F_tuplestore_puttuple(m, v1047, v1072)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L19
	} else {
		goto L266
	}
L266:
	;
	v1076 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v1076)
	goto L197
L267:
	;
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101)+22)))
	v1104 = v1100 & int32(_a_F_verify_heapam_6)
	if base.Ui32(v1104) < base.Ui32(v1102) {
		goto L272
	} else {
		goto L273
	}
L268:
	;
	v1094 = F_HeapTupleGetUpdateXid(m, v1083)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L19
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+4))
	v1099 = v1098
	v1100 = v946
	v1101 = v1083
	goto L267
L271:
	;
	v1096 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[37]))))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[46])))
	v1099 = v1094
	v1100 = v1096
	v1101 = v1097
	goto L267
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+980)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v24)+976)) = v1102
	v1111 = F_psprintf(m, int32(_a_F_verify_heapam_48), v24+int32(976))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L19
	} else {
		goto L275
	}
L273:
	;
	v1149 = v1101
	goto L274
L274:
	;
	v1150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1149)+20)))
	v1151 = int32(_a_F_verify_heapam_49)
	if v1150&v1151 == v1151 {
		goto L281
	} else {
		goto L282
	}
L275:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v1115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v1116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v1117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v1120 = F_Int64GetDatum(m, v1116)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L19
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v1117)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v1115
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v1120
	v1127 = int32(base.Ui32(v1117) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v1127)
	v1129 = F_cstring_to_text(m, v1111)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L19
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v1129
	F_pfree(m, v1111)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L19
	} else {
		goto L278
	}
L278:
	;
	v1138 = F_heap_form_tuple(m, v1114, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L19
	} else {
		goto L279
	}
L279:
	;
	F_tuplestore_puttuple(m, v1113, v1138)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L19
	} else {
		goto L280
	}
L280:
	;
	v1142 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v1142)
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[46])))
	v1149 = v1144
	goto L274
L281:
	;
	v1156 = F_pstrdup(m, int32(_a_F_verify_heapam_50))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L19
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v1193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1083)+18)))
	if v1099|base.B2i32(v1193&int32(_a_F_verify_heapam_51) == int32(0)) != 0 {
		v1244 = v1193
		goto L290
	} else {
		goto L291
	}
L284:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v1160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v1161 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v1162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v1165 = F_Int64GetDatum(m, v1161)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L19
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v1162)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v1165
	v1172 = int32(base.Ui32(v1162) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v1172)
	v1174 = F_cstring_to_text(m, v1156)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L19
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v1174
	F_pfree(m, v1156)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L19
	} else {
		goto L287
	}
L287:
	;
	v1183 = F_heap_form_tuple(m, v1159, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L19
	} else {
		goto L288
	}
L288:
	;
	F_tuplestore_puttuple(m, v1158, v1183)
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L19
	} else {
		goto L289
	}
L289:
	;
	v1187 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v1187)
	goto L283
L290:
	;
	if int32(0) <= base.I32_extend16_s(v1244) {
		goto L299
	} else {
		goto L300
	}
L291:
	;
	v1199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1083)+20)))
	if v1199&int32(2048)|base.B2i32(v1199&int32(768) == int32(512)) != 0 {
		v1244 = v1193
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v1209 = F_psprintf(m, int32(_a_F_verify_heapam_52), int32(0))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L19
	} else {
		goto L293
	}
L293:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v1213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v1214 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v1215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v1218 = F_Int64GetDatum(m, v1214)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L19
	} else {
		goto L294
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v1215)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v1218
	v1225 = int32(base.Ui32(v1215) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v1225)
	v1227 = F_cstring_to_text(m, v1209)
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L19
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v1227
	F_pfree(m, v1209)
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L19
	} else {
		goto L296
	}
L296:
	;
	v1236 = F_heap_form_tuple(m, v1212, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L19
	} else {
		goto L297
	}
L297:
	;
	F_tuplestore_puttuple(m, v1211, v1236)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L19
	} else {
		goto L298
	}
L298:
	;
	v1240 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v1240)
	v1242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1083)+18)))
	v1244 = v1242
	goto L290
L299:
	;
	if v1089&int32(1) == int32(0) {
		goto L311
	} else {
		goto L312
	}
L300:
	;
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083)+21)))
	if v1250&int32(32) != 0 {
		goto L299
	} else {
		goto L301
	}
L301:
	;
	v1255 = F_psprintf(m, int32(_a_F_verify_heapam_53), int32(0))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L19
	} else {
		goto L302
	}
L302:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v1259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v1260 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v1261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v1264 = F_Int64GetDatum(m, v1260)
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L19
	} else {
		goto L303
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v1261)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v1259
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v1264
	v1271 = int32(base.Ui32(v1261) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v1271)
	v1273 = F_cstring_to_text(m, v1255)
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L19
	} else {
		goto L304
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v1273
	F_pfree(m, v1255)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L19
	} else {
		goto L305
	}
L305:
	;
	v1282 = F_heap_form_tuple(m, v1258, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L19
	} else {
		goto L306
	}
L306:
	;
	F_tuplestore_puttuple(m, v1257, v1282)
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L19
	} else {
		goto L307
	}
L307:
	;
	v1286 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v1286)
	goto L299
L308:
	;
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29])))
	v2657 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[46])))
	v2658 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2657)+12)))
	v2661 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2657)+14)))
	if v2656 != v2658<<(uint(int32(16))%32)|v2661 {
		goto L197
	} else {
		goto L611
	}
L309:
	;
	if base.Ui32(v1104) < base.Ui32(v1102) {
		goto L308
	} else {
		goto L346
	}
L310:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[47])))
	if v1391 == int32(1) {
		goto L331
	} else {
		goto L332
	}
L311:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[46])))
	v1297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1296)+22)))
	if v1297 != int32(24) {
		goto L310
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[47])))
	v1304 = base.I32_div_s(v1300+int32(7), int32(8))
	v1308 = (v1304 + int32(30)) & int32(-8)
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[46])))
	v1310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1309)+22)))
	if v1308 == v1310 {
		v1473 = v1309
		goto L309
	} else {
		goto L315
	}
L314:
	;
	v1473 = v1296
	goto L309
L315:
	;
	if v1300 == int32(1) {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+948)) = v1310
	*(*int32)(unsafe.Add(mBase, uint32(v24)+944)) = v1308
	v1319 = F_psprintf(m, int32(_a_F_verify_heapam_54), v24+int32(944))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L19
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+968)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v24)+964)) = v1310
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v1308
	v1358 = F_psprintf(m, int32(_a_F_verify_heapam_55), v24+int32(960))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L19
	} else {
		goto L325
	}
L319:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v1323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v1324 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v1325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v1328 = F_Int64GetDatum(m, v1324)
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L19
	} else {
		goto L320
	}
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v1325)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v1328
	v1335 = int32(base.Ui32(v1325) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v1335)
	v1337 = F_cstring_to_text(m, v1319)
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L19
	} else {
		goto L321
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v1337
	F_pfree(m, v1319)
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L19
	} else {
		goto L322
	}
L322:
	;
	v1346 = F_heap_form_tuple(m, v1322, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L19
	} else {
		goto L323
	}
L323:
	;
	F_tuplestore_puttuple(m, v1321, v1346)
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L19
	} else {
		goto L324
	}
L324:
	;
	v1350 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v1350)
	goto L308
L325:
	;
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v1362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v1363 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v1364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v1367 = F_Int64GetDatum(m, v1363)
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L19
	} else {
		goto L326
	}
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v1364)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v1362
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v1367
	v1374 = int32(base.Ui32(v1364) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v1374)
	v1376 = F_cstring_to_text(m, v1358)
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L19
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v1376
	F_pfree(m, v1358)
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L19
	} else {
		goto L328
	}
L328:
	;
	v1385 = F_heap_form_tuple(m, v1361, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L19
	} else {
		goto L329
	}
L329:
	;
	F_tuplestore_puttuple(m, v1360, v1385)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L19
	} else {
		goto L330
	}
L330:
	;
	v1389 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v1389)
	goto L308
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+916)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v24)+912)) = int32(24)
	v1400 = F_psprintf(m, int32(_a_F_verify_heapam_56), v24+int32(912))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L19
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+936)) = v1391
	*(*int32)(unsafe.Add(mBase, uint32(v24)+932)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v24)+928)) = int32(24)
	v1440 = F_psprintf(m, int32(_a_F_verify_heapam_57), v24+int32(928))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L19
	} else {
		goto L340
	}
L334:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v1404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v1405 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v1406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v1409 = F_Int64GetDatum(m, v1405)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L19
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v1406)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v1409
	v1416 = int32(base.Ui32(v1406) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v1416)
	v1418 = F_cstring_to_text(m, v1400)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L19
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v1418
	F_pfree(m, v1400)
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L19
	} else {
		goto L337
	}
L337:
	;
	v1427 = F_heap_form_tuple(m, v1403, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L19
	} else {
		goto L338
	}
L338:
	;
	F_tuplestore_puttuple(m, v1402, v1427)
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L19
	} else {
		goto L339
	}
L339:
	;
	v1431 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v1431)
	goto L308
L340:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v1444 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v1445 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v1446 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v1449 = F_Int64GetDatum(m, v1445)
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L19
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v1446)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v1444
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v1449
	v1456 = int32(base.Ui32(v1446) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v1456)
	v1458 = F_cstring_to_text(m, v1440)
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L19
	} else {
		goto L342
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v1458
	F_pfree(m, v1440)
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L19
	} else {
		goto L343
	}
L343:
	;
	v1467 = F_heap_form_tuple(m, v1443, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L19
	} else {
		goto L344
	}
L344:
	;
	F_tuplestore_puttuple(m, v1442, v1467)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L19
	} else {
		goto L345
	}
L345:
	;
	v1471 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v1471)
	goto L308
L346:
	;
	v1478 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[48]))) = uint8(v1478)
	v1480 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v716))) = uint8(v1480)
	v1483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1473)+20)))
	v1484 = int32(768)
	if v1483&v1484 != v1484 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1473)))
	v1489 = v1488
	goto L349
L348:
	;
	v1489 = int32(2)
	goto L349
L349:
	;
	v1494 = F_get_xid_status(m, v1489, v24+int32(_a_F_verify_heapam_2), v24+int32(_a_F_verify_heapam_58))
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L19
	} else {
		goto L357
	}
L350:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[12])))
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v2091)+52))
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v2092)))
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[47])))
	if v2093 < v2094 {
		goto L508
	} else {
		goto L509
	}
L351:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v2057 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2058 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v2059 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v2062 = F_Int64GetDatum(m, v2058)
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L19
	} else {
		goto L503
	}
L352:
	;
	F_ReadMultiXactIdRange(m, v389, v391)
	mBase = m.M
	v2050 = m.ExcPending
	if v2050 != 0 {
		goto L19
	} else {
		goto L501
	}
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+864)) = v1489
	v2008 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[23])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+872)) = uint32(v2008)
	v2011 = int64(base.Ui64(v2008) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+868)) = uint32(v2011)
	v2016 = F_psprintf(m, int32(_a_F_verify_heapam_59), v24+int32(864))
	mBase = m.M
	v2017 = m.ExcPending
	if v2017 != 0 {
		goto L19
	} else {
		goto L495
	}
L354:
	;
	v1582 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v716))) = uint8(v1582)
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[49])))
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(1040)+v730))) = v1587
	v1589 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1473)+20)))
	if v1589&int32(256) != 0 {
		goto L370
	} else {
		goto L371
	}
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+896)) = v1489
	v1541 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[6])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+904)) = uint32(v1541)
	v1544 = int64(base.Ui64(v1541) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+900)) = uint32(v1544)
	v1549 = F_psprintf(m, int32(_a_F_verify_heapam_60), v24+int32(896))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L19
	} else {
		goto L364
	}
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+880)) = v1489
	v1499 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[26])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+888)) = uint32(v1499)
	v1502 = int64(base.Ui64(v1499) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+884)) = uint32(v1502)
	v1507 = F_psprintf(m, int32(_a_F_verify_heapam_61), v24+int32(880))
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L19
	} else {
		goto L358
	}
L357:
	;
	switch v1494 - int32(1) {
	case 0:
		goto L353
	case 1:
		goto L356
	case 2:
		goto L355
	case 3:
		goto L354
	default:
		goto L308
	}
L358:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v1511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v1512 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v1513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v1516 = F_Int64GetDatum(m, v1512)
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L19
	} else {
		goto L359
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v1513)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v1511
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v1516
	v1523 = int32(base.Ui32(v1513) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v1523)
	v1525 = F_cstring_to_text(m, v1507)
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L19
	} else {
		goto L360
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v1525
	F_pfree(m, v1507)
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L19
	} else {
		goto L361
	}
L361:
	;
	v1534 = F_heap_form_tuple(m, v1510, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L19
	} else {
		goto L362
	}
L362:
	;
	F_tuplestore_puttuple(m, v1509, v1534)
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L19
	} else {
		goto L363
	}
L363:
	;
	v1538 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v1538)
	goto L308
L364:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v1553 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v1554 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v1555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v1558 = F_Int64GetDatum(m, v1554)
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L19
	} else {
		goto L365
	}
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v1555)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v1553
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v1558
	v1565 = int32(base.Ui32(v1555) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v1565)
	v1567 = F_cstring_to_text(m, v1549)
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L19
	} else {
		goto L366
	}
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v1567
	F_pfree(m, v1549)
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L19
	} else {
		goto L367
	}
L367:
	;
	v1576 = F_heap_form_tuple(m, v1552, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L19
	} else {
		goto L368
	}
L368:
	;
	F_tuplestore_puttuple(m, v1551, v1576)
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L19
	} else {
		goto L369
	}
L369:
	;
	v1580 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v1580)
	goto L308
L370:
	;
	v1762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1473)+21)))
	if v1762&int32(16) == int32(0) {
		goto L420
	} else {
		goto L421
	}
L371:
	;
	v1592 = base.I32_extend16_s(v1589)
	if v1592&int32(512) != 0 {
		goto L308
	} else {
		goto L372
	}
L372:
	;
	if v1592&int32(_a_F_verify_heapam_51) != 0 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1473)+8))
	v1602 = F_get_xid_status(m, v1597, v24+int32(_a_F_verify_heapam_2), v24+int32(_a_F_verify_heapam_62))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L19
	} else {
		goto L381
	}
L374:
	;
	goto L375
L375:
	;
	if v1592 < int32(0) {
		goto L396
	} else {
		goto L397
	}
L376:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[50])))
	switch v1656 {
	case 0:
		goto L350
	case 1:
		goto L391
	case 2:
		goto L390
	default:
		goto L370
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+848)) = v1597
	v1642 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[26])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+856)) = uint32(v1642)
	v1645 = int64(base.Ui64(v1642) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+852)) = uint32(v1645)
	v1652 = F_psprintf(m, int32(_a_F_verify_heapam_63), v24+int32(848))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L19
	} else {
		goto L388
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+832)) = v1597
	v1627 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[6])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+840)) = uint32(v1627)
	v1630 = int64(base.Ui64(v1627) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+836)) = uint32(v1630)
	v1637 = F_psprintf(m, int32(_a_F_verify_heapam_64), v24+int32(832))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L19
	} else {
		goto L386
	}
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+816)) = v1597
	v1612 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[23])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+824)) = uint32(v1612)
	v1615 = int64(base.Ui64(v1612) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+820)) = uint32(v1615)
	v1622 = F_psprintf(m, int32(_a_F_verify_heapam_65), v24+int32(816))
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L19
	} else {
		goto L384
	}
L380:
	;
	v1607 = F_pstrdup(m, int32(_a_F_verify_heapam_66))
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L19
	} else {
		goto L382
	}
L381:
	;
	switch v1602 {
	case 0:
		goto L380
	case 1:
		goto L379
	case 2:
		goto L377
	case 3:
		goto L378
	default:
		goto L376
	}
L382:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1607)
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L19
	} else {
		goto L383
	}
L383:
	;
	goto L308
L384:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1622)
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L19
	} else {
		goto L385
	}
L385:
	;
	goto L308
L386:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1637)
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L19
	} else {
		goto L387
	}
L387:
	;
	goto L308
L388:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1652)
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L19
	} else {
		goto L389
	}
L389:
	;
	goto L308
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+800)) = v1597
	v1673 = F_psprintf(m, int32(_a_F_verify_heapam_67), v24+int32(800))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L19
	} else {
		goto L394
	}
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+784)) = v1597
	v1663 = F_psprintf(m, int32(_a_F_verify_heapam_68), v24+int32(784))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L19
	} else {
		goto L392
	}
L392:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1663)
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L19
	} else {
		goto L393
	}
L393:
	;
	goto L308
L394:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1673)
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L19
	} else {
		goto L395
	}
L395:
	;
	goto L308
L396:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1473)+8))
	v1684 = F_get_xid_status(m, v1679, v24+int32(_a_F_verify_heapam_2), v24+int32(_a_F_verify_heapam_62))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L19
	} else {
		goto L404
	}
L397:
	;
	goto L398
L398:
	;
	if v1587 != 0 {
		goto L308
	} else {
		goto L419
	}
L399:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[50])))
	switch v1738 - int32(1) {
	case 0:
		goto L414
	case 1:
		goto L413
	case 2:
		goto L350
	default:
		goto L370
	}
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+768)) = v1679
	v1724 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[26])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+776)) = uint32(v1724)
	v1727 = int64(base.Ui64(v1724) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+772)) = uint32(v1727)
	v1734 = F_psprintf(m, int32(_a_F_verify_heapam_69), v24+int32(768))
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L19
	} else {
		goto L411
	}
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = v1679
	v1709 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[6])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+760)) = uint32(v1709)
	v1712 = int64(base.Ui64(v1709) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+756)) = uint32(v1712)
	v1719 = F_psprintf(m, int32(_a_F_verify_heapam_70), v24+int32(752))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L19
	} else {
		goto L409
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+736)) = v1679
	v1694 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[23])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+744)) = uint32(v1694)
	v1697 = int64(base.Ui64(v1694) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+740)) = uint32(v1697)
	v1704 = F_psprintf(m, int32(_a_F_verify_heapam_71), v24+int32(736))
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L19
	} else {
		goto L407
	}
L403:
	;
	v1689 = F_pstrdup(m, int32(_a_F_verify_heapam_72))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L19
	} else {
		goto L405
	}
L404:
	;
	switch v1684 {
	case 0:
		goto L403
	case 1:
		goto L402
	case 2:
		goto L400
	case 3:
		goto L401
	default:
		goto L399
	}
L405:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1689)
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L19
	} else {
		goto L406
	}
L406:
	;
	goto L308
L407:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1704)
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L19
	} else {
		goto L408
	}
L408:
	;
	goto L308
L409:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1719)
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L19
	} else {
		goto L410
	}
L410:
	;
	goto L308
L411:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1734)
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L19
	} else {
		goto L412
	}
L412:
	;
	goto L308
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+720)) = v1679
	v1757 = F_psprintf(m, int32(_a_F_verify_heapam_73), v24+int32(720))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L19
	} else {
		goto L417
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+704)) = v1679
	v1747 = F_psprintf(m, int32(_a_F_verify_heapam_74), v24+int32(704))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L19
	} else {
		goto L415
	}
L415:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1747)
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L19
	} else {
		goto L416
	}
L416:
	;
	goto L308
L417:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1757)
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L19
	} else {
		goto L418
	}
L418:
	;
	goto L308
L419:
	;
	goto L370
L420:
	;
	v1835 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1473)+20)))
	if v1835&int32(2048) != 0 {
		goto L445
	} else {
		goto L446
	}
L421:
	;
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v1473)+4))
	if v1767 == int32(0) {
		goto L352
	} else {
		goto L422
	}
L422:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[7])))
	goto L424
L423:
	;
	F_ReadMultiXactIdRange(m, v389, v391)
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L19
	} else {
		goto L430
	}
L424:
	;
	if int32(base.Ui32(v1767-v1770)>>(uint(int32(31))%32)) != 0 {
		goto L423
	} else {
		goto L425
	}
L425:
	;
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[51])))
	goto L426
L426:
	;
	if int32(base.Ui32(v1767-v1774)>>(uint(int32(31))%32)) != 0 {
		goto L423
	} else {
		goto L427
	}
L427:
	;
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[52])))
	goto L428
L428:
	;
	if base.B2i32(v1778-v1767 <= int32(0)) == int32(0) {
		goto L420
	} else {
		goto L429
	}
L429:
	;
	goto L423
L430:
	;
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[7])))
	goto L432
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+672)) = v1767
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[51])))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+676)) = v1823
	v1830 = F_psprintf(m, int32(_a_F_verify_heapam_75), v24+int32(672))
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L19
	} else {
		goto L443
	}
L432:
	;
	if int32(base.Ui32(v1767-v1786)>>(uint(int32(31))%32)) == int32(0) {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[51])))
	goto L436
L434:
	;
	goto L435
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+656)) = v1767
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+660)) = v1815
	v1820 = F_psprintf(m, int32(_a_F_verify_heapam_76), v24+int32(656))
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		goto L19
	} else {
		goto L442
	}
L436:
	;
	if int32(base.Ui32(v1767-v1792)>>(uint(int32(31))%32)) != 0 {
		goto L431
	} else {
		goto L437
	}
L437:
	;
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[52])))
	goto L438
L438:
	;
	if base.B2i32(v1796-v1767 <= int32(0)) == int32(0) {
		goto L420
	} else {
		goto L439
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+688)) = v1767
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[52])))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+692)) = v1803
	v1810 = F_psprintf(m, int32(_a_F_verify_heapam_77), v24+int32(688))
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L19
	} else {
		goto L440
	}
L440:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1810)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L19
	} else {
		goto L441
	}
L441:
	;
	goto L350
L442:
	;
	v2054 = v1820
	goto L351
L443:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1830)
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L19
	} else {
		goto L444
	}
L444:
	;
	goto L350
L445:
	;
	v1838 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[48]))) = uint8(v1838)
	goto L350
L446:
	;
	goto L447
L447:
	;
	v1842 = int32(0)
	if base.B2i32(v1835&int32(128) == v1842)&base.B2i32(v1835&int32(_a_F_verify_heapam_78) != int32(64)) == v1842 {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v1851 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[48]))) = uint8(v1851)
	goto L350
L449:
	;
	goto L450
L450:
	;
	if v1835&int32(_a_F_verify_heapam_27) != 0 {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v1855 = F_HeapTupleGetUpdateXid(m, v1473)
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L19
	} else {
		goto L459
	}
L452:
	;
	goto L453
L453:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1473)+4))
	v1939 = F_get_xid_status(m, v1934, v24+int32(_a_F_verify_heapam_2), v24+int32(_a_F_verify_heapam_79))
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L19
	} else {
		goto L481
	}
L454:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[53])))
	switch v1915 {
	case 0:
		goto L470
	case 1, 2:
		goto L471
	case 3:
		goto L469
	default:
		goto L350
	}
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+640)) = v1855
	v1901 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[26])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+648)) = uint32(v1901)
	v1904 = int64(base.Ui64(v1901) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+644)) = uint32(v1904)
	v1911 = F_psprintf(m, int32(_a_F_verify_heapam_80), v24+int32(640))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L19
	} else {
		goto L467
	}
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+624)) = v1855
	v1886 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[6])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+632)) = uint32(v1886)
	v1889 = int64(base.Ui64(v1886) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+628)) = uint32(v1889)
	v1896 = F_psprintf(m, int32(_a_F_verify_heapam_81), v24+int32(624))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L19
	} else {
		goto L465
	}
L457:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+608)) = v1855
	v1871 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[23])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+616)) = uint32(v1871)
	v1874 = int64(base.Ui64(v1871) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+612)) = uint32(v1874)
	v1881 = F_psprintf(m, int32(_a_F_verify_heapam_82), v24+int32(608))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L19
	} else {
		goto L463
	}
L458:
	;
	v1866 = F_pstrdup(m, int32(_a_F_verify_heapam_83))
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L19
	} else {
		goto L461
	}
L459:
	;
	v1861 = F_get_xid_status(m, v1855, v24+int32(_a_F_verify_heapam_2), v24+int32(_a_F_verify_heapam_79))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L19
	} else {
		goto L460
	}
L460:
	;
	switch v1861 {
	case 0:
		goto L458
	case 1:
		goto L457
	case 2:
		goto L455
	case 3:
		goto L456
	default:
		goto L454
	}
L461:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1866)
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L19
	} else {
		goto L462
	}
L462:
	;
	goto L350
L463:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1881)
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L19
	} else {
		goto L464
	}
L464:
	;
	goto L350
L465:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1896)
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L19
	} else {
		goto L466
	}
L466:
	;
	goto L350
L467:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1911)
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L19
	} else {
		goto L468
	}
L468:
	;
	goto L350
L469:
	;
	v1932 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[48]))) = uint8(v1932)
	goto L350
L470:
	;
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[9])))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1918))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1855)) == int32(0) {
		goto L473
	} else {
		goto L474
	}
L471:
	;
	v1916 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[48]))) = uint8(v1916)
	goto L350
L472:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[48]))) = uint8(v1930)
	goto L350
L473:
	;
	v1930 = base.B2i32(base.Ui32(v1855) < base.Ui32(v1918))
	goto L472
L474:
	;
	goto L475
L475:
	;
	v1930 = int32(base.Ui32(v1855-v1918) >> (uint(int32(31)) % 32))
	goto L472
L476:
	;
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[53])))
	switch v1988 {
	case 0:
		goto L489
	case 1, 2:
		goto L490
	case 3:
		goto L488
	default:
		goto L350
	}
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+592)) = v1934
	v1974 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[26])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+600)) = uint32(v1974)
	v1977 = int64(base.Ui64(v1974) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+596)) = uint32(v1977)
	v1984 = F_psprintf(m, int32(_a_F_verify_heapam_84), v24+int32(592))
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L19
	} else {
		goto L486
	}
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+576)) = v1934
	v1959 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[6])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+584)) = uint32(v1959)
	v1962 = int64(base.Ui64(v1959) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+580)) = uint32(v1962)
	v1969 = F_psprintf(m, int32(_a_F_verify_heapam_85), v24+int32(576))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L19
	} else {
		goto L484
	}
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+560)) = v1934
	v1944 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[23])))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+568)) = uint32(v1944)
	v1947 = int64(base.Ui64(v1944) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+564)) = uint32(v1947)
	v1954 = F_psprintf(m, int32(_a_F_verify_heapam_86), v24+int32(560))
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L19
	} else {
		goto L482
	}
L480:
	;
	v1941 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[48]))) = uint8(v1941)
	goto L350
L481:
	;
	switch v1939 {
	case 0:
		goto L480
	case 1:
		goto L479
	case 2:
		goto L477
	case 3:
		goto L478
	default:
		goto L476
	}
L482:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1954)
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L19
	} else {
		goto L483
	}
L483:
	;
	goto L308
L484:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1969)
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L19
	} else {
		goto L485
	}
L485:
	;
	goto L308
L486:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v1984)
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		goto L19
	} else {
		goto L487
	}
L487:
	;
	goto L308
L488:
	;
	v2005 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[48]))) = uint8(v2005)
	goto L350
L489:
	;
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[9])))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1991))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1934)) == int32(0) {
		goto L492
	} else {
		goto L493
	}
L490:
	;
	v1989 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[48]))) = uint8(v1989)
	goto L350
L491:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[48]))) = uint8(v2003)
	goto L350
L492:
	;
	v2003 = base.B2i32(base.Ui32(v1934) < base.Ui32(v1991))
	goto L491
L493:
	;
	goto L494
L494:
	;
	v2003 = int32(base.Ui32(v1934-v1991) >> (uint(int32(31)) % 32))
	goto L491
L495:
	;
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v2020 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2021 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v2022 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v2025 = F_Int64GetDatum(m, v2021)
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L19
	} else {
		goto L496
	}
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v2022)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v2020
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v2025
	v2032 = int32(base.Ui32(v2022) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v2032)
	v2034 = F_cstring_to_text(m, v2016)
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L19
	} else {
		goto L497
	}
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v2034
	F_pfree(m, v2016)
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L19
	} else {
		goto L498
	}
L498:
	;
	v2043 = F_heap_form_tuple(m, v2019, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L19
	} else {
		goto L499
	}
L499:
	;
	F_tuplestore_puttuple(m, v2018, v2043)
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L19
	} else {
		goto L500
	}
L500:
	;
	v2047 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v2047)
	goto L308
L501:
	;
	v2052 = F_pstrdup(m, int32(_a_F_verify_heapam_87))
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L19
	} else {
		goto L502
	}
L502:
	;
	v2054 = v2052
	goto L351
L503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v2059)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v2057
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v2062
	v2069 = int32(base.Ui32(v2059) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v2069)
	v2071 = F_cstring_to_text(m, v2054)
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L19
	} else {
		goto L504
	}
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v2071
	F_pfree(m, v2054)
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L19
	} else {
		goto L505
	}
L505:
	;
	v2080 = F_heap_form_tuple(m, v2056, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L19
	} else {
		goto L506
	}
L506:
	;
	F_tuplestore_puttuple(m, v2055, v2080)
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		goto L19
	} else {
		goto L507
	}
L507:
	;
	v2084 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v2084)
	goto L350
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+404)) = v2093
	*(*int32)(unsafe.Add(mBase, uint32(v24)+400)) = v2094
	v2101 = F_psprintf(m, int32(_a_F_verify_heapam_88), v24+int32(400))
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L19
	} else {
		goto L511
	}
L509:
	;
	goto L510
L510:
	;
	v2134 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))) = uint16(v2134)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[54]))) = v2134
	if v2094 <= v2134 {
		goto L517
	} else {
		goto L518
	}
L511:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v2105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2106 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v2107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v2110 = F_Int64GetDatum(m, v2106)
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L19
	} else {
		goto L512
	}
L512:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v2107)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v2105
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v2110
	v2117 = int32(base.Ui32(v2107) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v2117)
	v2119 = F_cstring_to_text(m, v2101)
	mBase = m.M
	v2120 = m.ExcPending
	if v2120 != 0 {
		goto L19
	} else {
		goto L513
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v2119
	F_pfree(m, v2101)
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L19
	} else {
		goto L514
	}
L514:
	;
	v2128 = F_heap_form_tuple(m, v2104, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2129 = m.ExcPending
	if v2129 != 0 {
		goto L19
	} else {
		goto L515
	}
L515:
	;
	F_tuplestore_puttuple(m, v2103, v2128)
	mBase = m.M
	v2131 = m.ExcPending
	if v2131 != 0 {
		goto L19
	} else {
		goto L516
	}
L516:
	;
	v2132 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v2132)
	goto L308
L517:
	;
	v2633 = int32(_a_F_verify_heapam_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))) = uint16(v2633)
	goto L308
L518:
	;
	v2144 = v2134
	goto L519
L519:
	;
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[12])))
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v2162)+52))
	v2168 = v2163 + v2144<<(uint(int32(4))%32) + int32(20)
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[54])))
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[46])))
	v2171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2170)+22)))
	v2172 = v2169 + v2171
	v2173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[37]))))
	if base.Ui32(v2173) < base.Ui32(v2172) {
		goto L521
	} else {
		goto L522
	}
L520:
	;
	goto L517
L521:
	;
	v2175 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2168)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+424)) = v2173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+420)) = v2172
	*(*int32)(unsafe.Add(mBase, uint32(v24)+416)) = v2175
	v2182 = F_psprintf(m, int32(_a_F_verify_heapam_89), v24+int32(416))
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L19
	} else {
		goto L524
	}
L522:
	;
	goto L523
L523:
	;
	v2215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2170)+20)))
	if v2215&int32(1) != 0 {
		goto L531
	} else {
		goto L532
	}
L524:
	;
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v2186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2187 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v2188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v2191 = F_Int64GetDatum(m, v2187)
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L19
	} else {
		goto L525
	}
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v2188)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v2186
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v2191
	v2198 = int32(base.Ui32(v2188) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v2198)
	v2200 = F_cstring_to_text(m, v2182)
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L19
	} else {
		goto L526
	}
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v2200
	F_pfree(m, v2182)
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L19
	} else {
		goto L527
	}
L527:
	;
	v2209 = F_heap_form_tuple(m, v2185, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2210 = m.ExcPending
	if v2210 != 0 {
		goto L19
	} else {
		goto L528
	}
L528:
	;
	F_tuplestore_puttuple(m, v2184, v2209)
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L19
	} else {
		goto L529
	}
L529:
	;
	v2213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v2213)
	goto L517
L530:
	;
	v2605 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	v2607 = v2605 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))) = uint16(v2607)
	v2609 = base.I32_extend16_s(v2607)
	v2610 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[47])))
	if v2609 < v2610 {
		v2144 = v2609
		goto L519
	} else {
		goto L610
	}
L531:
	;
	v2221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2170+v2144>>(uint(int32(3))%32))+23)))
	if int32(base.Ui32(v2221)>>(uint(v2144&int32(7))%32))&int32(1) == int32(0) {
		goto L530
	} else {
		goto L534
	}
L532:
	;
	goto L533
L533:
	;
	v2229 = v2171 + v2170
	v2230 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2168)+4)))
	if v2230 != int32(-1) {
		goto L535
	} else {
		goto L536
	}
L534:
	;
	goto L533
L535:
	;
	v2233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2168)+12)))
	v2237 = int32(0)
	v2239 = (v2169 + v2233 - int32(1)) & (v2237 - v2233)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[54]))) = v2239
	if v2230 <= v2237 {
		goto L538
	} else {
		goto L539
	}
L536:
	;
	goto L537
L537:
	;
	v2293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2229+v2169))))
	if v2293 == int32(0) {
		goto L548
	} else {
		goto L549
	}
L538:
	;
	v2244 = F_strlen(m, v2229+v2239)
	mBase = m.M
	v2247 = v2244 + int32(1)
	goto L540
L539:
	;
	v2247 = v2230
	goto L540
L540:
	;
	v2248 = v2247 + v2239
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[54]))) = v2248
	v2250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2170)+22)))
	v2251 = v2248 + v2250
	if base.Ui32(v2251) <= base.Ui32(v2173) {
		goto L530
	} else {
		goto L541
	}
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+552)) = v2173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+548)) = v2251
	*(*int32)(unsafe.Add(mBase, uint32(v24)+544)) = v2230
	v2259 = F_psprintf(m, int32(_a_F_verify_heapam_90), v24+int32(544))
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L19
	} else {
		goto L542
	}
L542:
	;
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v2262 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v2263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2264 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v2265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v2268 = F_Int64GetDatum(m, v2264)
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L19
	} else {
		goto L543
	}
L543:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v2265)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v2268
	v2275 = int32(base.Ui32(v2265) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v2275)
	v2277 = F_cstring_to_text(m, v2259)
	mBase = m.M
	v2278 = m.ExcPending
	if v2278 != 0 {
		goto L19
	} else {
		goto L544
	}
L544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v2277
	F_pfree(m, v2259)
	mBase = m.M
	v2281 = m.ExcPending
	if v2281 != 0 {
		goto L19
	} else {
		goto L545
	}
L545:
	;
	v2286 = F_heap_form_tuple(m, v2262, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L19
	} else {
		goto L546
	}
L546:
	;
	F_tuplestore_puttuple(m, v2261, v2286)
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
		goto L19
	} else {
		goto L547
	}
L547:
	;
	v2290 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v2290)
	goto L517
L548:
	;
	v2296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2168)+12)))
	v2304 = (v2169 + v2296 - int32(1)) & (int32(0) - v2296)
	goto L550
L549:
	;
	v2304 = v2169
	goto L550
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[54]))) = v2304
	v2306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2168)+6)))
	if v2306 == int32(1) {
		goto L1
	} else {
		goto L551
	}
L551:
	;
	v2309 = v2229 + v2304
	v2310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2309))))
	if v2310 == int32(1) {
		goto L553
	} else {
		goto L554
	}
L552:
	;
	v2363 = v2362 + v2304
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[54]))) = v2363
	v2365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2170)+22)))
	v2366 = v2363 + v2365
	if base.Ui32(v2173) < base.Ui32(v2366) {
		goto L564
	} else {
		goto L565
	}
L553:
	;
	v2313 = int32(18)
	v2314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2309)+1)))
	if v2314 == v2313 {
		v2362 = v2313
		goto L552
	} else {
		goto L556
	}
L554:
	;
	goto L555
L555:
	;
	v2354 = int32(1)
	if v2310&v2354 != 0 {
		v2362 = int32(base.Ui32(v2310) >> (uint(v2354) % 32))
		goto L552
	} else {
		goto L563
	}
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+528)) = v2314
	v2321 = F_psprintf(m, int32(_a_F_verify_heapam_91), v24+int32(528))
	mBase = m.M
	v2322 = m.ExcPending
	if v2322 != 0 {
		goto L19
	} else {
		goto L557
	}
L557:
	;
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v2324 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v2325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2326 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v2327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v2330 = F_Int64GetDatum(m, v2326)
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L19
	} else {
		goto L558
	}
L558:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v2327)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v2325
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v2330
	v2337 = int32(base.Ui32(v2327) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v2337)
	v2339 = F_cstring_to_text(m, v2321)
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		goto L19
	} else {
		goto L559
	}
L559:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v2339
	F_pfree(m, v2321)
	mBase = m.M
	v2343 = m.ExcPending
	if v2343 != 0 {
		goto L19
	} else {
		goto L560
	}
L560:
	;
	v2348 = F_heap_form_tuple(m, v2324, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2349 = m.ExcPending
	if v2349 != 0 {
		goto L19
	} else {
		goto L561
	}
L561:
	;
	F_tuplestore_puttuple(m, v2323, v2348)
	mBase = m.M
	v2351 = m.ExcPending
	if v2351 != 0 {
		goto L19
	} else {
		goto L562
	}
L562:
	;
	v2352 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v2352)
	goto L517
L563:
	;
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v2309)))
	v2362 = int32(base.Ui32(v2358) >> (uint(int32(2)) % 32))
	goto L552
L564:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+456)) = v2173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+452)) = v2366
	*(*int32)(unsafe.Add(mBase, uint32(v24)+448)) = int32(-1)
	v2375 = F_psprintf(m, int32(_a_F_verify_heapam_90), v24+int32(448))
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L19
	} else {
		goto L567
	}
L565:
	;
	goto L566
L566:
	;
	v2408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2309))))
	if v2408 != int32(1) {
		goto L530
	} else {
		goto L573
	}
L567:
	;
	v2377 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v2378 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v2379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2380 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v2381 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v2384 = F_Int64GetDatum(m, v2380)
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L19
	} else {
		goto L568
	}
L568:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v2381)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v2379
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v2384
	v2391 = int32(base.Ui32(v2381) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v2391)
	v2393 = F_cstring_to_text(m, v2375)
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L19
	} else {
		goto L569
	}
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v2393
	F_pfree(m, v2375)
	mBase = m.M
	v2397 = m.ExcPending
	if v2397 != 0 {
		goto L19
	} else {
		goto L570
	}
L570:
	;
	v2402 = F_heap_form_tuple(m, v2378, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2403 = m.ExcPending
	if v2403 != 0 {
		goto L19
	} else {
		goto L571
	}
L571:
	;
	F_tuplestore_puttuple(m, v2377, v2402)
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L19
	} else {
		goto L572
	}
L572:
	;
	v2406 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v2406)
	goto L517
L573:
	;
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(v2309)+10))
	v2412 = *(*int32)(unsafe.Add(mBase, uint32(v2309)+6))
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v2309)+2))
	if int32(1073741824) <= v2413 {
		goto L574
	} else {
		goto L575
	}
L574:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+520)) = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+516)) = v2413
	*(*int32)(unsafe.Add(mBase, uint32(v24)+512)) = v2411
	v2423 = F_psprintf(m, int32(_a_F_verify_heapam_92), v24+int32(512))
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L19
	} else {
		goto L577
	}
L575:
	;
	goto L576
L576:
	;
	v2465 = int32(0)
	if base.B2i32(base.Ui32(v2413-int32(4)) <= base.Ui32(v2412&int32(1073741823)))|base.B2i32(v2465 <= v2412) == v2465 {
		goto L583
	} else {
		goto L584
	}
L577:
	;
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v2426 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v2427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2428 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v2429 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v2432 = F_Int64GetDatum(m, v2428)
	mBase = m.M
	v2433 = m.ExcPending
	if v2433 != 0 {
		goto L19
	} else {
		goto L578
	}
L578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v2429)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v2427
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v2432
	v2439 = int32(base.Ui32(v2429) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v2439)
	v2441 = F_cstring_to_text(m, v2423)
	mBase = m.M
	v2442 = m.ExcPending
	if v2442 != 0 {
		goto L19
	} else {
		goto L579
	}
L579:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v2441
	F_pfree(m, v2423)
	mBase = m.M
	v2445 = m.ExcPending
	if v2445 != 0 {
		goto L19
	} else {
		goto L580
	}
L580:
	;
	v2450 = F_heap_form_tuple(m, v2426, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2451 = m.ExcPending
	if v2451 != 0 {
		goto L19
	} else {
		goto L581
	}
L581:
	;
	F_tuplestore_puttuple(m, v2425, v2450)
	mBase = m.M
	v2453 = m.ExcPending
	if v2453 != 0 {
		goto L19
	} else {
		goto L582
	}
L582:
	;
	v2454 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v2454)
	goto L576
L583:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+496)) = v2411
	*(*int32)(unsafe.Add(mBase, uint32(v24)+500)) = int32(base.Ui32(v2412) >> (uint(int32(30)) % 32))
	v2477 = F_psprintf(m, int32(_a_F_verify_heapam_93), v24+int32(496))
	mBase = m.M
	v2478 = m.ExcPending
	if v2478 != 0 {
		goto L19
	} else {
		goto L586
	}
L584:
	;
	goto L585
L585:
	;
	if v2215&int32(4) == int32(0) {
		goto L592
	} else {
		goto L593
	}
L586:
	;
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v2480 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v2481 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2482 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v2483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v2486 = F_Int64GetDatum(m, v2482)
	mBase = m.M
	v2487 = m.ExcPending
	if v2487 != 0 {
		goto L19
	} else {
		goto L587
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v2483)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v2481
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v2486
	v2493 = int32(base.Ui32(v2483) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v2493)
	v2495 = F_cstring_to_text(m, v2477)
	mBase = m.M
	v2496 = m.ExcPending
	if v2496 != 0 {
		goto L19
	} else {
		goto L588
	}
L588:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v2495
	F_pfree(m, v2477)
	mBase = m.M
	v2499 = m.ExcPending
	if v2499 != 0 {
		goto L19
	} else {
		goto L589
	}
L589:
	;
	v2504 = F_heap_form_tuple(m, v2480, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2505 = m.ExcPending
	if v2505 != 0 {
		goto L19
	} else {
		goto L590
	}
L590:
	;
	F_tuplestore_puttuple(m, v2479, v2504)
	mBase = m.M
	v2507 = m.ExcPending
	if v2507 != 0 {
		goto L19
	} else {
		goto L591
	}
L591:
	;
	v2508 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v2508)
	goto L585
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+464)) = v2411
	v2522 = F_psprintf(m, int32(_a_F_verify_heapam_94), v24+int32(464))
	mBase = m.M
	v2523 = m.ExcPending
	if v2523 != 0 {
		goto L19
	} else {
		goto L595
	}
L593:
	;
	goto L594
L594:
	;
	v2555 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[12])))
	v2556 = *(*int32)(unsafe.Add(mBase, uint32(v2555)+48))
	v2557 = *(*int32)(unsafe.Add(mBase, uint32(v2556)+112))
	if v2557 == int32(0) {
		goto L601
	} else {
		goto L602
	}
L595:
	;
	v2524 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v2526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2527 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v2528 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v2531 = F_Int64GetDatum(m, v2527)
	mBase = m.M
	v2532 = m.ExcPending
	if v2532 != 0 {
		goto L19
	} else {
		goto L596
	}
L596:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v2528)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v2526
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v2531
	v2538 = int32(base.Ui32(v2528) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v2538)
	v2540 = F_cstring_to_text(m, v2522)
	mBase = m.M
	v2541 = m.ExcPending
	if v2541 != 0 {
		goto L19
	} else {
		goto L597
	}
L597:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v2540
	F_pfree(m, v2522)
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L19
	} else {
		goto L598
	}
L598:
	;
	v2549 = F_heap_form_tuple(m, v2525, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2550 = m.ExcPending
	if v2550 != 0 {
		goto L19
	} else {
		goto L599
	}
L599:
	;
	F_tuplestore_puttuple(m, v2524, v2549)
	mBase = m.M
	v2552 = m.ExcPending
	if v2552 != 0 {
		goto L19
	} else {
		goto L600
	}
L600:
	;
	v2553 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v2553)
	goto L530
L601:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+480)) = v2411
	v2566 = F_psprintf(m, int32(_a_F_verify_heapam_95), v24+int32(480))
	mBase = m.M
	v2567 = m.ExcPending
	if v2567 != 0 {
		goto L19
	} else {
		goto L604
	}
L602:
	;
	goto L603
L603:
	;
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[19])))
	if v2570 == int32(0) {
		goto L530
	} else {
		goto L606
	}
L604:
	;
	F_report_corruption(m, v24+int32(_a_F_verify_heapam_2), v2566)
	mBase = m.M
	v2569 = m.ExcPending
	if v2569 != 0 {
		goto L19
	} else {
		goto L605
	}
L605:
	;
	goto L530
L606:
	;
	v2573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[48]))))
	if v2573&int32(1) != 0 {
		goto L530
	} else {
		goto L607
	}
L607:
	;
	v2577 = F_palloc0(m, int32(24))
	mBase = m.M
	v2578 = m.ExcPending
	if v2578 != 0 {
		goto L19
	} else {
		goto L608
	}
L608:
	;
	v2580 = v2309 + int32(2)
	v2581 = *(*int64)(unsafe.Add(mBase, uint32(v2580)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2577)+8)) = v2581
	v2583 = *(*int64)(unsafe.Add(mBase, uint32(v2580)))
	*(*int64)(unsafe.Add(mBase, uint32(v2577))) = v2583
	v2585 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29])))
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+16)) = v2585
	v2587 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2577)+20)) = uint16(v2587)
	v2589 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2577)+22)) = uint16(v2589)
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[55])))
	v2592 = F_lappend(m, v2591, v2577)
	mBase = m.M
	v2593 = m.ExcPending
	if v2593 != 0 {
		goto L19
	} else {
		goto L609
	}
L609:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[55]))) = v2592
	goto L530
L610:
	;
	goto L520
L611:
	;
	v2664 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2657)+16)))
	if base.Ui32(v679) <= base.Ui32((v2664-int32(1))&int32(_a_F_verify_heapam_6)) {
		goto L197
	} else {
		goto L612
	}
L612:
	;
	v2670 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	if v2670 == v2664 {
		goto L197
	} else {
		goto L613
	}
L613:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(_a_F_verify_heapam_31)+v2670<<(uint(int32(1))%32)))) = uint16(v2664)
	goto L197
L614:
	;
	goto L196
L615:
	;
	v2735 = v2713 & int32(_a_F_verify_heapam_6)
	v2739 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(_a_F_verify_heapam_31)+v2735<<(uint(int32(1))%32)))))
	if v2739 == int32(0) {
		goto L617
	} else {
		goto L618
	}
L616:
	;
	v3150 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))) = uint16(v3150)
	v3156 = v3150
	v3157 = v3150
	goto L699
L617:
	;
	v3143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v3145 = v3143 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))) = uint16(v3145)
	if base.Ui32(v3145&int32(_a_F_verify_heapam_6)) <= base.Ui32(v679) {
		v2713 = v3145
		goto L615
	} else {
		goto L698
	}
L618:
	;
	v2745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(_a_F_verify_heapam_29)+v2739))))
	if v2745 != int32(1) {
		goto L617
	} else {
		goto L619
	}
L619:
	;
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[32])))
	v2750 = v2748 + int32(20)
	v2751 = int32(2)
	v2752 = v2739 << (uint(v2751) % 32)
	v2753 = v2750 + v2752
	v2754 = *(*int32)(unsafe.Add(mBase, uint32(v2753)))
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(v2750+v2735<<(uint(v2751)%32))))
	if v2758&int32(_a_F_verify_heapam_96) == int32(_a_F_verify_heapam_97) {
		goto L621
	} else {
		goto L622
	}
L620:
	;
	v3102 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v3103 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v3104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v3105 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v3106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v3109 = F_Int64GetDatum(m, v3105)
	mBase = m.M
	v3110 = m.ExcPending
	if v3110 != 0 {
		goto L19
	} else {
		goto L693
	}
L621:
	;
	v2766 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2748+v2754&int32(_a_F_verify_heapam_32))+18)))
	if int32(0) <= v2766 {
		goto L624
	} else {
		goto L625
	}
L622:
	;
	goto L623
L623:
	;
	if v2754&int32(_a_F_verify_heapam_96) == int32(_a_F_verify_heapam_97) {
		goto L617
	} else {
		goto L637
	}
L624:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+176)) = v2739
	v2773 = F_psprintf(m, int32(_a_F_verify_heapam_98), v24+int32(176))
	mBase = m.M
	v2774 = m.ExcPending
	if v2774 != 0 {
		goto L19
	} else {
		goto L627
	}
L625:
	;
	goto L626
L626:
	;
	v2814 = v24 + int32(_a_F_verify_heapam_26) + v2739<<(uint(int32(1))%32)
	v2815 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2814))))
	if v2815 != 0 {
		goto L633
	} else {
		goto L634
	}
L627:
	;
	v2775 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v2777 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2778 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v2779 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v2782 = F_Int64GetDatum(m, v2778)
	mBase = m.M
	v2783 = m.ExcPending
	if v2783 != 0 {
		goto L19
	} else {
		goto L628
	}
L628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v2779)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v2777
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v2782
	v2789 = int32(base.Ui32(v2779) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v2789)
	v2791 = F_cstring_to_text(m, v2773)
	mBase = m.M
	v2792 = m.ExcPending
	if v2792 != 0 {
		goto L19
	} else {
		goto L629
	}
L629:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v2791
	F_pfree(m, v2773)
	mBase = m.M
	v2795 = m.ExcPending
	if v2795 != 0 {
		goto L19
	} else {
		goto L630
	}
L630:
	;
	v2800 = F_heap_form_tuple(m, v2776, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2801 = m.ExcPending
	if v2801 != 0 {
		goto L19
	} else {
		goto L631
	}
L631:
	;
	F_tuplestore_puttuple(m, v2775, v2800)
	mBase = m.M
	v2803 = m.ExcPending
	if v2803 != 0 {
		goto L19
	} else {
		goto L632
	}
L632:
	;
	v2804 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v2804)
	goto L626
L633:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+164)) = v2815
	*(*int32)(unsafe.Add(mBase, uint32(v24)+160)) = v2739
	v2821 = F_psprintf(m, int32(_a_F_verify_heapam_99), v24+int32(160))
	mBase = m.M
	v2822 = m.ExcPending
	if v2822 != 0 {
		goto L19
	} else {
		goto L636
	}
L634:
	;
	goto L635
L635:
	;
	v2823 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2814))) = uint16(v2823)
	goto L617
L636:
	;
	v3093 = v2821
	goto L620
L637:
	;
	v2831 = v2748 + v2758&int32(_a_F_verify_heapam_32)
	v2832 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2831)+20)))
	if v2832&int32(_a_F_verify_heapam_47) == int32(_a_F_verify_heapam_27) {
		goto L639
	} else {
		goto L640
	}
L638:
	;
	v2850 = v2842 + v2844&int32(_a_F_verify_heapam_32)
	v2851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2850)+20)))
	v2852 = int32(768)
	if v2851&v2852 != v2852 {
		goto L643
	} else {
		goto L644
	}
L639:
	;
	v2837 = F_HeapTupleGetUpdateXid(m, v2831)
	mBase = m.M
	v2838 = m.ExcPending
	if v2838 != 0 {
		goto L19
	} else {
		goto L642
	}
L640:
	;
	goto L641
L641:
	;
	v2841 = *(*int32)(unsafe.Add(mBase, uint32(v2831)+4))
	v2842 = v2748
	v2843 = v2841
	v2844 = v2754
	goto L638
L642:
	;
	v2839 = *(*int32)(unsafe.Add(mBase, uint32(v2753)))
	v2840 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[32])))
	v2842 = v2840
	v2843 = v2837
	v2844 = v2839
	goto L638
L643:
	;
	v2856 = *(*int32)(unsafe.Add(mBase, uint32(v2850)))
	v2858 = v2856
	goto L645
L644:
	;
	v2858 = int32(2)
	goto L645
L645:
	;
	if base.B2i32(v2843 == int32(0))|base.B2i32(v2858 != v2843) != 0 {
		goto L617
	} else {
		goto L646
	}
L646:
	;
	v2865 = v24 + int32(_a_F_verify_heapam_26) + v2739<<(uint(int32(1))%32)
	v2866 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2865))))
	if v2866 != 0 {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+276)) = v2866
	*(*int32)(unsafe.Add(mBase, uint32(v24)+272)) = v2739
	v2872 = F_psprintf(m, int32(_a_F_verify_heapam_100), v24+int32(272))
	mBase = m.M
	v2873 = m.ExcPending
	if v2873 != 0 {
		goto L19
	} else {
		goto L650
	}
L648:
	;
	goto L649
L649:
	;
	v2874 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2865))) = uint16(v2874)
	v2876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2831)+19)))
	if v2876&int32(64) == int32(0) {
		goto L652
	} else {
		goto L653
	}
L650:
	;
	v3093 = v2872
	goto L620
L651:
	;
	v2977 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2831)+20)))
	v2978 = int32(768)
	if v2977&v2978 != v2978 {
		goto L670
	} else {
		goto L671
	}
L652:
	;
	v2881 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2850)+18)))
	if int32(0) <= v2881 {
		goto L651
	} else {
		goto L655
	}
L653:
	;
	goto L654
L654:
	;
	v2930 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2850)+18)))
	if v2930 < int32(0) {
		goto L651
	} else {
		goto L663
	}
L655:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+256)) = v2739
	v2888 = F_psprintf(m, int32(_a_F_verify_heapam_101), v24+int32(256))
	mBase = m.M
	v2889 = m.ExcPending
	if v2889 != 0 {
		goto L19
	} else {
		goto L656
	}
L656:
	;
	v2890 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v2892 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2893 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v2894 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v2897 = F_Int64GetDatum(m, v2893)
	mBase = m.M
	v2898 = m.ExcPending
	if v2898 != 0 {
		goto L19
	} else {
		goto L657
	}
L657:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v2894)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v2892
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v2897
	v2904 = int32(base.Ui32(v2894) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v2904)
	v2906 = F_cstring_to_text(m, v2888)
	mBase = m.M
	v2907 = m.ExcPending
	if v2907 != 0 {
		goto L19
	} else {
		goto L658
	}
L658:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v2906
	F_pfree(m, v2888)
	mBase = m.M
	v2910 = m.ExcPending
	if v2910 != 0 {
		goto L19
	} else {
		goto L659
	}
L659:
	;
	v2915 = F_heap_form_tuple(m, v2891, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2916 = m.ExcPending
	if v2916 != 0 {
		goto L19
	} else {
		goto L660
	}
L660:
	;
	F_tuplestore_puttuple(m, v2890, v2915)
	mBase = m.M
	v2918 = m.ExcPending
	if v2918 != 0 {
		goto L19
	} else {
		goto L661
	}
L661:
	;
	v2919 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v2919)
	v2921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2831)+19)))
	if v2921&int32(64) == int32(0) {
		goto L651
	} else {
		goto L662
	}
L662:
	;
	goto L654
L663:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+240)) = v2739
	v2937 = F_psprintf(m, int32(_a_F_verify_heapam_102), v24+int32(240))
	mBase = m.M
	v2938 = m.ExcPending
	if v2938 != 0 {
		goto L19
	} else {
		goto L664
	}
L664:
	;
	v2939 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v2940 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v2941 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2942 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v2943 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v2946 = F_Int64GetDatum(m, v2942)
	mBase = m.M
	v2947 = m.ExcPending
	if v2947 != 0 {
		goto L19
	} else {
		goto L665
	}
L665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v2943)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v2941
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v2946
	v2953 = int32(base.Ui32(v2943) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v2953)
	v2955 = F_cstring_to_text(m, v2937)
	mBase = m.M
	v2956 = m.ExcPending
	if v2956 != 0 {
		goto L19
	} else {
		goto L666
	}
L666:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v2955
	F_pfree(m, v2937)
	mBase = m.M
	v2959 = m.ExcPending
	if v2959 != 0 {
		goto L19
	} else {
		goto L667
	}
L667:
	;
	v2964 = F_heap_form_tuple(m, v2940, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2965 = m.ExcPending
	if v2965 != 0 {
		goto L19
	} else {
		goto L668
	}
L668:
	;
	F_tuplestore_puttuple(m, v2939, v2964)
	mBase = m.M
	v2967 = m.ExcPending
	if v2967 != 0 {
		goto L19
	} else {
		goto L669
	}
L669:
	;
	v2968 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v2968)
	goto L651
L670:
	;
	v2982 = *(*int32)(unsafe.Add(mBase, uint32(v2831)))
	v2983 = v2982
	goto L672
L671:
	;
	v2983 = int32(2)
	goto L672
L672:
	;
	v2984 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v2986 = v24 + int32(_a_F_verify_heapam_30)
	v2988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2984+v2986))))
	if v2988 != int32(1) {
		v3050 = v2984
		goto L673
	} else {
		goto L674
	}
L673:
	;
	v3056 = v3050 & int32(_a_F_verify_heapam_6)
	v3058 = v24 + int32(_a_F_verify_heapam_30)
	v3060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3056+v3058))))
	if v3060 != int32(1) {
		goto L617
	} else {
		goto L686
	}
L674:
	;
	v2992 = v24 + int32(1040)
	v2993 = int32(2)
	v2996 = *(*int32)(unsafe.Add(mBase, uint32(v2992+v2984<<(uint(v2993)%32))))
	if v2996 != v2993 {
		v3050 = v2984
		goto L673
	} else {
		goto L675
	}
L675:
	;
	v3000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2739+v2986))))
	if v3000 != int32(1) {
		v3050 = v2984
		goto L673
	} else {
		goto L676
	}
L676:
	;
	v3004 = *(*int32)(unsafe.Add(mBase, uint32(v2992+v2752)))
	if v3004 != 0 {
		v3050 = v2984
		goto L673
	} else {
		goto L677
	}
L677:
	;
	v3005 = F_TransactionIdIsInProgress(m, v2983)
	mBase = m.M
	v3006 = m.ExcPending
	if v3006 != 0 {
		goto L19
	} else {
		goto L678
	}
L678:
	;
	v3007 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	if v3005 == int32(0) {
		v3050 = v3007
		goto L673
	} else {
		goto L679
	}
L679:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+232)) = v2843
	*(*int32)(unsafe.Add(mBase, uint32(v24)+228)) = v3007
	*(*int32)(unsafe.Add(mBase, uint32(v24)+224)) = v2983
	v3016 = F_psprintf(m, int32(_a_F_verify_heapam_103), v24+int32(224))
	mBase = m.M
	v3017 = m.ExcPending
	if v3017 != 0 {
		goto L19
	} else {
		goto L680
	}
L680:
	;
	v3018 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v3019 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v3020 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v3021 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v3022 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v3025 = F_Int64GetDatum(m, v3021)
	mBase = m.M
	v3026 = m.ExcPending
	if v3026 != 0 {
		goto L19
	} else {
		goto L681
	}
L681:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v3022)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v3020
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v3025
	v3032 = int32(base.Ui32(v3022) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v3032)
	v3034 = F_cstring_to_text(m, v3016)
	mBase = m.M
	v3035 = m.ExcPending
	if v3035 != 0 {
		goto L19
	} else {
		goto L682
	}
L682:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v3034
	F_pfree(m, v3016)
	mBase = m.M
	v3038 = m.ExcPending
	if v3038 != 0 {
		goto L19
	} else {
		goto L683
	}
L683:
	;
	v3043 = F_heap_form_tuple(m, v3019, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v3044 = m.ExcPending
	if v3044 != 0 {
		goto L19
	} else {
		goto L684
	}
L684:
	;
	F_tuplestore_puttuple(m, v3018, v3043)
	mBase = m.M
	v3046 = m.ExcPending
	if v3046 != 0 {
		goto L19
	} else {
		goto L685
	}
L685:
	;
	v3047 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v3047)
	v3049 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v3050 = v3049
	goto L673
L686:
	;
	v3064 = v24 + int32(1040)
	v3068 = *(*int32)(unsafe.Add(mBase, uint32(v3064+v3056<<(uint(int32(2))%32))))
	if v3068 != int32(3) {
		goto L617
	} else {
		goto L687
	}
L687:
	;
	v3072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2739+v3058))))
	if v3072 != int32(1) {
		goto L617
	} else {
		goto L688
	}
L688:
	;
	v3076 = *(*int32)(unsafe.Add(mBase, uint32(v3064+v2752)))
	switch v3076 {
	case 0:
		goto L689
	default:
		goto L617
	case 2:
		goto L690
	}
L689:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+216)) = v2843
	*(*int32)(unsafe.Add(mBase, uint32(v24)+212)) = v3056
	*(*int32)(unsafe.Add(mBase, uint32(v24)+208)) = v2983
	v3091 = F_psprintf(m, int32(_a_F_verify_heapam_104), v24+int32(208))
	mBase = m.M
	v3092 = m.ExcPending
	if v3092 != 0 {
		goto L19
	} else {
		goto L692
	}
L690:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+200)) = v2843
	*(*int32)(unsafe.Add(mBase, uint32(v24)+196)) = v3056
	*(*int32)(unsafe.Add(mBase, uint32(v24)+192)) = v2983
	v3083 = F_psprintf(m, int32(_a_F_verify_heapam_105), v24+int32(192))
	mBase = m.M
	v3084 = m.ExcPending
	if v3084 != 0 {
		goto L19
	} else {
		goto L691
	}
L691:
	;
	v3093 = v3083
	goto L620
L692:
	;
	v3093 = v3091
	goto L620
L693:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v3106)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v3104
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v3109
	v3116 = int32(base.Ui32(v3106) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v3116)
	v3118 = F_cstring_to_text(m, v3093)
	mBase = m.M
	v3119 = m.ExcPending
	if v3119 != 0 {
		goto L19
	} else {
		goto L694
	}
L694:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v3118
	F_pfree(m, v3093)
	mBase = m.M
	v3122 = m.ExcPending
	if v3122 != 0 {
		goto L19
	} else {
		goto L695
	}
L695:
	;
	v3127 = F_heap_form_tuple(m, v3103, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v3128 = m.ExcPending
	if v3128 != 0 {
		goto L19
	} else {
		goto L696
	}
L696:
	;
	F_tuplestore_puttuple(m, v3102, v3127)
	mBase = m.M
	v3130 = m.ExcPending
	if v3130 != 0 {
		goto L19
	} else {
		goto L697
	}
L697:
	;
	v3131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v3131)
	goto L617
L698:
	;
	goto L616
L699:
	;
	v3178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(_a_F_verify_heapam_30)+v3157))))
	if v3178 != int32(1) {
		v3242 = v3156
		goto L701
	} else {
		goto L702
	}
L700:
	;
	goto L188
L701:
	;
	v3248 = v3242 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))) = uint16(v3248)
	v3251 = v3248 & int32(_a_F_verify_heapam_6)
	if base.Ui32(v3251) <= base.Ui32(v679) {
		v3156 = v3248
		v3157 = v3251
		goto L699
	} else {
		goto L713
	}
L702:
	;
	v3182 = v3157 << (uint(int32(2)) % 32)
	v3186 = *(*int32)(unsafe.Add(mBase, uint32(v3182+(v24+int32(1040)))))
	switch v3186 {
	case 0, 2:
		goto L703
	default:
		v3242 = v3156
		goto L701
	}
L703:
	;
	v3192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(_a_F_verify_heapam_26)+v3157<<(uint(int32(1))%32)))))
	if v3192 != 0 {
		v3242 = v3156
		goto L701
	} else {
		goto L704
	}
L704:
	;
	v3193 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[32])))
	v3195 = *(*int32)(unsafe.Add(mBase, uint32(v3193+v3182)+20))
	if v3195&int32(_a_F_verify_heapam_96) == int32(_a_F_verify_heapam_97) {
		v3242 = v3156
		goto L701
	} else {
		goto L705
	}
L705:
	;
	v3203 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3193+v3195&int32(_a_F_verify_heapam_32))+18)))
	if int32(0) <= v3203 {
		v3242 = v3156
		goto L701
	} else {
		goto L706
	}
L706:
	;
	v3208 = F_psprintf(m, int32(_a_F_verify_heapam_106), int32(0))
	mBase = m.M
	v3209 = m.ExcPending
	if v3209 != 0 {
		goto L19
	} else {
		goto L707
	}
L707:
	;
	v3210 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v3211 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	v3212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v3213 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[29]))))
	v3214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	v3217 = F_Int64GetDatum(m, v3213)
	mBase = m.M
	v3218 = m.ExcPending
	if v3218 != 0 {
		goto L19
	} else {
		goto L708
	}
L708:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[40]))) = base.I32_extend16_s(v3214)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[41]))) = v3212
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[42]))) = v3217
	v3224 = int32(base.Ui32(v3214) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[43]))) = uint8(v3224)
	v3226 = F_cstring_to_text(m, v3208)
	mBase = m.M
	v3227 = m.ExcPending
	if v3227 != 0 {
		goto L19
	} else {
		goto L709
	}
L709:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[44]))) = v3226
	F_pfree(m, v3208)
	mBase = m.M
	v3230 = m.ExcPending
	if v3230 != 0 {
		goto L19
	} else {
		goto L710
	}
L710:
	;
	v3235 = F_heap_form_tuple(m, v3211, v24+int32(_a_F_verify_heapam_35), v24+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v3236 = m.ExcPending
	if v3236 != 0 {
		goto L19
	} else {
		goto L711
	}
L711:
	;
	F_tuplestore_puttuple(m, v3210, v3235)
	mBase = m.M
	v3238 = m.ExcPending
	if v3238 != 0 {
		goto L19
	} else {
		goto L712
	}
L712:
	;
	v3239 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v3239)
	v3241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[33]))))
	v3242 = v3241
	goto L701
L713:
	;
	goto L700
L714:
	;
	v3277 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[55])))
	if v3277 != 0 {
		goto L715
	} else {
		goto L716
	}
L715:
	;
	v3278 = *(*int32)(unsafe.Add(mBase, uint32(v3277)+4))
	if int32(0) < v3278 {
		goto L718
	} else {
		goto L719
	}
L716:
	;
	goto L717
L717:
	;
	if v42 == int32(0) {
		goto L170
	} else {
		goto L790
	}
L718:
	;
	v3289 = int32(0)
	goto L721
L719:
	;
	v3664 = v3277
	goto L720
L720:
	;
	F_list_free_deep(m, v3664)
	mBase = m.M
	v3666 = m.ExcPending
	if v3666 != 0 {
		goto L19
	} else {
		goto L789
	}
L721:
	;
	v3303 = *(*int32)(unsafe.Add(mBase, uint32(v3277)+12))
	v3307 = *(*int32)(unsafe.Add(mBase, uint32(v3303+v3289<<(uint(int32(2))%32))))
	v3308 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+4))
	v3310 = v24 + int32(_a_F_verify_heapam_35)
	v3314 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+8))
	F_ScanKeyInit(m, v3310, int32(1), int32(3), int32(184), v3314)
	mBase = m.M
	v3316 = m.ExcPending
	if v3316 != 0 {
		goto L19
	} else {
		goto L723
	}
L722:
	;
	v3642 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[55])))
	v3664 = v3642
	goto L720
L723:
	;
	v3318 = v3308 & int32(1073741823)
	v3322 = base.I32_div_u_s(v3318-int32(1), int32(1996))
	v3323 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[19])))
	v3324 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[21])))
	v3325 = F_get_toast_snapshot(m)
	mBase = m.M
	v3326 = m.ExcPending
	if v3326 != 0 {
		goto L19
	} else {
		goto L726
	}
L724:
	;
	v3639 = v3289 + int32(1)
	v3640 = *(*int32)(unsafe.Add(mBase, uint32(v3277)+4))
	if v3639 < v3640 {
		v3289 = v3639
		goto L721
	} else {
		goto L788
	}
L725:
	;
	v3586 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3307)+20)))
	v3587 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3307)+16)))
	v3588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3307)+22)))
	v3589 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v3590 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[49]))) = int32(0)
	v3593 = F_Int64GetDatum(m, v3587)
	mBase = m.M
	v3594 = m.ExcPending
	if v3594 != 0 {
		goto L19
	} else {
		goto L783
	}
L726:
	;
	v3328 = F_systable_beginscan_ordered(m, v3323, v3324, v3325, int32(1), v3310)
	mBase = m.M
	v3329 = m.ExcPending
	if v3329 != 0 {
		goto L19
	} else {
		goto L727
	}
L727:
	;
	v3331 = F_systable_getnext_ordered(m, v3328, int32(1))
	mBase = m.M
	v3332 = m.ExcPending
	if v3332 != 0 {
		goto L19
	} else {
		goto L728
	}
L728:
	;
	if v3331 != 0 {
		goto L729
	} else {
		goto L730
	}
L729:
	;
	v3340 = v3331
	v3342 = int32(0)
	goto L732
L730:
	;
	goto L731
L731:
	;
	F_systable_endscan_ordered(m, v3328)
	mBase = m.M
	v3557 = m.ExcPending
	if v3557 != 0 {
		goto L19
	} else {
		goto L781
	}
L732:
	;
	v3359 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[19])))
	v3360 = *(*int32)(unsafe.Add(mBase, uint32(v3359)+52))
	v3363 = F_fastgetattr_5(m, v3340, int32(2), v3360, v24+int32(_a_F_verify_heapam_62))
	mBase = m.M
	v3364 = m.ExcPending
	if v3364 != 0 {
		goto L19
	} else {
		goto L734
	}
L733:
	;
	F_systable_endscan_ordered(m, v3328)
	mBase = m.M
	v3545 = m.ExcPending
	if v3545 != 0 {
		goto L19
	} else {
		goto L778
	}
L734:
	;
	v3365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[50]))))
	if v3365 == int32(1) {
		goto L737
	} else {
		goto L738
	}
L735:
	;
	v3542 = F_systable_getnext_ordered(m, v3328, int32(1))
	mBase = m.M
	v3543 = m.ExcPending
	if v3543 != 0 {
		goto L19
	} else {
		goto L776
	}
L736:
	;
	v3502 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3307)+20)))
	v3503 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3307)+16)))
	v3504 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3307)+22)))
	v3505 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v3506 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[49]))) = int32(0)
	v3509 = F_Int64GetDatum(m, v3503)
	mBase = m.M
	v3510 = m.ExcPending
	if v3510 != 0 {
		goto L19
	} else {
		goto L771
	}
L737:
	;
	v3368 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v3368
	v3373 = F_psprintf(m, int32(_a_F_verify_heapam_107), v24-int32(-64))
	mBase = m.M
	v3374 = m.ExcPending
	if v3374 != 0 {
		goto L19
	} else {
		goto L740
	}
L738:
	;
	goto L739
L739:
	;
	if v3363 != v3342 {
		goto L741
	} else {
		goto L742
	}
L740:
	;
	v3495 = v3373
	v3496 = v3342
	goto L736
L741:
	;
	v3376 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+152)) = v3342
	*(*int32)(unsafe.Add(mBase, uint32(v24)+148)) = v3363
	*(*int32)(unsafe.Add(mBase, uint32(v24)+144)) = v3376
	v3383 = F_psprintf(m, int32(_a_F_verify_heapam_108), v24+int32(144))
	mBase = m.M
	v3384 = m.ExcPending
	if v3384 != 0 {
		goto L19
	} else {
		goto L744
	}
L742:
	;
	goto L743
L743:
	;
	v3423 = v3363 + int32(1)
	v3425 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[19])))
	v3426 = *(*int32)(unsafe.Add(mBase, uint32(v3425)+52))
	v3429 = F_fastgetattr_5(m, v3340, int32(3), v3426, v24+int32(_a_F_verify_heapam_62))
	mBase = m.M
	v3430 = m.ExcPending
	if v3430 != 0 {
		goto L19
	} else {
		goto L750
	}
L744:
	;
	v3385 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3307)+20)))
	v3386 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3307)+16)))
	v3387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3307)+22)))
	v3388 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[11])))
	v3389 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[10])))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[49]))) = int32(0)
	v3392 = F_Int64GetDatum(m, v3386)
	mBase = m.M
	v3393 = m.ExcPending
	if v3393 != 0 {
		goto L19
	} else {
		goto L745
	}
L745:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[56]))) = base.I32_extend16_s(v3387)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[57]))) = v3385
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = v3392
	v3399 = int32(base.Ui32(v3387) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[58]))) = uint8(v3399)
	v3401 = F_cstring_to_text(m, v3383)
	mBase = m.M
	v3402 = m.ExcPending
	if v3402 != 0 {
		goto L19
	} else {
		goto L746
	}
L746:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[59]))) = v3401
	F_pfree(m, v3383)
	mBase = m.M
	v3405 = m.ExcPending
	if v3405 != 0 {
		goto L19
	} else {
		goto L747
	}
L747:
	;
	v3410 = F_heap_form_tuple(m, v3389, v24+int32(_a_F_verify_heapam_36), v24+int32(_a_F_verify_heapam_58))
	mBase = m.M
	v3411 = m.ExcPending
	if v3411 != 0 {
		goto L19
	} else {
		goto L748
	}
L748:
	;
	F_tuplestore_puttuple(m, v3388, v3410)
	mBase = m.M
	v3413 = m.ExcPending
	if v3413 != 0 {
		goto L19
	} else {
		goto L749
	}
L749:
	;
	v3414 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v3414)
	goto L743
L750:
	;
	v3431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[50]))))
	if v3431 == int32(1) {
		goto L751
	} else {
		goto L752
	}
L751:
	;
	v3434 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = v3363
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v3434
	v3440 = F_psprintf(m, int32(_a_F_verify_heapam_109), v24+int32(80))
	mBase = m.M
	v3441 = m.ExcPending
	if v3441 != 0 {
		goto L19
	} else {
		goto L754
	}
L752:
	;
	goto L753
L753:
	;
	v3442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3429))))
	if v3442&int32(3) == int32(0) {
		goto L757
	} else {
		goto L758
	}
L754:
	;
	v3495 = v3440
	v3496 = v3423
	goto L736
L755:
	;
	v3485 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+8))
	v3486 = *(*int32)(unsafe.Add(mBase, uint32(v3429)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+136)) = v3486
	*(*int32)(unsafe.Add(mBase, uint32(v24)+132)) = v3363
	*(*int32)(unsafe.Add(mBase, uint32(v24)+128)) = v3485
	v3493 = F_psprintf(m, int32(_a_F_verify_heapam_110), v24+int32(128))
	mBase = m.M
	v3494 = m.ExcPending
	if v3494 != 0 {
		goto L19
	} else {
		goto L770
	}
L756:
	;
	if v3322 < v3363 {
		goto L761
	} else {
		goto L762
	}
L757:
	;
	v3447 = *(*int32)(unsafe.Add(mBase, uint32(v3429)))
	v3460 = int32(base.Ui32(v3447)>>(uint(int32(2))%32)) - int32(4)
	goto L756
L758:
	;
	goto L759
L759:
	;
	if v3442&int32(1) == int32(0) {
		goto L755
	} else {
		goto L760
	}
L760:
	;
	v3456 = int32(1)
	v3460 = int32(base.Ui32(v3442)>>(uint(v3456)%32)) - v3456
	goto L756
L761:
	;
	v3462 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+104)) = v3322
	*(*int32)(unsafe.Add(mBase, uint32(v24)+100)) = v3363
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v3462
	v3469 = F_psprintf(m, int32(_a_F_verify_heapam_111), v24+int32(96))
	mBase = m.M
	v3470 = m.ExcPending
	if v3470 != 0 {
		goto L19
	} else {
		goto L764
	}
L762:
	;
	goto L763
L763:
	;
	if v3363 < v3322 {
		goto L765
	} else {
		goto L766
	}
L764:
	;
	v3495 = v3469
	v3496 = v3423
	goto L736
L765:
	;
	v3473 = int32(1996)
	goto L767
L766:
	;
	v3473 = v3322*int32(-1996) + v3318
	goto L767
L767:
	;
	if v3460 == v3473 {
		v3535 = v3423
		goto L735
	} else {
		goto L768
	}
L768:
	;
	v3475 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+124)) = v3473
	*(*int32)(unsafe.Add(mBase, uint32(v24)+120)) = v3460
	*(*int32)(unsafe.Add(mBase, uint32(v24)+116)) = v3363
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = v3475
	v3483 = F_psprintf(m, int32(_a_F_verify_heapam_112), v24+int32(112))
	mBase = m.M
	v3484 = m.ExcPending
	if v3484 != 0 {
		goto L19
	} else {
		goto L769
	}
L769:
	;
	v3495 = v3483
	v3496 = v3423
	goto L736
L770:
	;
	v3495 = v3493
	v3496 = v3423
	goto L736
L771:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[56]))) = base.I32_extend16_s(v3504)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[57]))) = v3502
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = v3509
	v3516 = int32(base.Ui32(v3504) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[58]))) = uint8(v3516)
	v3518 = F_cstring_to_text(m, v3495)
	mBase = m.M
	v3519 = m.ExcPending
	if v3519 != 0 {
		goto L19
	} else {
		goto L772
	}
L772:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[59]))) = v3518
	F_pfree(m, v3495)
	mBase = m.M
	v3522 = m.ExcPending
	if v3522 != 0 {
		goto L19
	} else {
		goto L773
	}
L773:
	;
	v3527 = F_heap_form_tuple(m, v3506, v24+int32(_a_F_verify_heapam_36), v24+int32(_a_F_verify_heapam_58))
	mBase = m.M
	v3528 = m.ExcPending
	if v3528 != 0 {
		goto L19
	} else {
		goto L774
	}
L774:
	;
	F_tuplestore_puttuple(m, v3505, v3527)
	mBase = m.M
	v3530 = m.ExcPending
	if v3530 != 0 {
		goto L19
	} else {
		goto L775
	}
L775:
	;
	v3531 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v3531)
	v3535 = v3496
	goto L735
L776:
	;
	if v3542 != 0 {
		v3340 = v3542
		v3342 = v3535
		goto L732
	} else {
		goto L777
	}
L777:
	;
	goto L733
L778:
	;
	if v3322 < v3535 {
		goto L724
	} else {
		goto L779
	}
L779:
	;
	v3547 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = v3535
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v3322
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v3547
	v3554 = F_psprintf(m, int32(_a_F_verify_heapam_113), v24+int32(48))
	mBase = m.M
	v3555 = m.ExcPending
	if v3555 != 0 {
		goto L19
	} else {
		goto L780
	}
L780:
	;
	v3568 = v3554
	goto L725
L781:
	;
	v3558 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v3558
	v3563 = F_psprintf(m, int32(_a_F_verify_heapam_114), v24+int32(32))
	mBase = m.M
	v3564 = m.ExcPending
	if v3564 != 0 {
		goto L19
	} else {
		goto L782
	}
L782:
	;
	v3568 = v3563
	goto L725
L783:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[56]))) = base.I32_extend16_s(v3588)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[57]))) = v3586
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[39]))) = v3593
	v3600 = int32(base.Ui32(v3588) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[58]))) = uint8(v3600)
	v3602 = F_cstring_to_text(m, v3568)
	mBase = m.M
	v3603 = m.ExcPending
	if v3603 != 0 {
		goto L19
	} else {
		goto L784
	}
L784:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[59]))) = v3602
	F_pfree(m, v3568)
	mBase = m.M
	v3606 = m.ExcPending
	if v3606 != 0 {
		goto L19
	} else {
		goto L785
	}
L785:
	;
	v3611 = F_heap_form_tuple(m, v3590, v24+int32(_a_F_verify_heapam_36), v24+int32(_a_F_verify_heapam_58))
	mBase = m.M
	v3612 = m.ExcPending
	if v3612 != 0 {
		goto L19
	} else {
		goto L786
	}
L786:
	;
	F_tuplestore_puttuple(m, v3589, v3611)
	mBase = m.M
	v3614 = m.ExcPending
	if v3614 != 0 {
		goto L19
	} else {
		goto L787
	}
L787:
	;
	v3615 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))) = uint8(v3615)
	goto L724
L788:
	;
	goto L722
L789:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[55]))) = int32(0)
	goto L717
L790:
	;
	v3692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[45]))))
	if v3692&int32(1) == int32(0) {
		goto L170
	} else {
		goto L791
	}
L791:
	;
	goto L172
L792:
	;
	v3720 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[0])))
	if v3720 != 0 {
		goto L793
	} else {
		goto L794
	}
L793:
	;
	F_ReleaseBuffer(m, v3720)
	mBase = m.M
	v3722 = m.ExcPending
	if v3722 != 0 {
		goto L19
	} else {
		goto L796
	}
L794:
	;
	goto L795
L795:
	;
	v3723 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[20])))
	if v3723 != 0 {
		goto L797
	} else {
		goto L798
	}
L796:
	;
	goto L795
L797:
	;
	v3724 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[18])))
	F_toast_close_indexes(m, v3723, v3724)
	mBase = m.M
	v3726 = m.ExcPending
	if v3726 != 0 {
		goto L19
	} else {
		goto L800
	}
L798:
	;
	goto L799
L799:
	;
	v3727 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[19])))
	if v3727 != 0 {
		goto L801
	} else {
		goto L802
	}
L800:
	;
	goto L799
L801:
	;
	F_relation_close(m, v3727, int32(1))
	mBase = m.M
	v3730 = m.ExcPending
	if v3730 != 0 {
		goto L19
	} else {
		goto L804
	}
L802:
	;
	goto L803
L803:
	;
	v3731 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_verify_heapam[12])))
	F_relation_close(m, v3731, int32(1))
	mBase = m.M
	v3734 = m.ExcPending
	if v3734 != 0 {
		goto L19
	} else {
		goto L805
	}
L804:
	;
	goto L803
L805:
	;
	goto L2
L806:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+432)) = int32(-1)
	F_errmsg_internal(m, int32(_a_F_verify_heapam_115), v24+int32(432))
	mBase = m.M
	v3773 = m.ExcPending
	if v3773 != 0 {
		goto L19
	} else {
		goto L807
	}
L807:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_116), int32(70), int32(_a_F_verify_heapam_117))
	mBase = m.M
	v3778 = m.ExcPending
	if v3778 != 0 {
		goto L19
	} else {
		goto L808
	}
L808:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
