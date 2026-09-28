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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int64
	_ = v158
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v164 int64
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
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
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v335 int32
	_ = v335
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int64
	_ = v470
	var v472 int64
	_ = v472
	var v474 int64
	_ = v474
	var v476 int64
	_ = v476
	var v478 int32
	_ = v478
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v513 int64
	_ = v513
	var v515 int64
	_ = v515
	var v517 int64
	_ = v517
	var v519 int64
	_ = v519
	var v523 int32
	_ = v523
	var v531 int32
	_ = v531
	var v545 int32
	_ = v545
	var v565 int32
	_ = v565
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int64
	_ = v721
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int64
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v762 int64
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int64
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int64
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int64
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v811 int64
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int64
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v916 int64
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int64
	_ = v958
	var v960 int64
	_ = v960
	var v962 int64
	_ = v962
	var v964 int64
	_ = v964
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v987 int64
	_ = v987
	var v989 int64
	_ = v989
	var v991 int64
	_ = v991
	var v993 int64
	_ = v993
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1083 int64
	_ = v1083
	var v1084 int64
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1093 int64
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1120 int32
	_ = v1120
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1159 int64
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1182 int64
	_ = v1182
	var v1183 int64
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1196 int32
	_ = v1196
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1210 int64
	_ = v1210
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1225 int32
	_ = v1225
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1253 int64
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1279 int64
	_ = v1279
	var v1283 int64
	_ = v1283
	var v1288 int64
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1317 int32
	_ = v1317
	var v1324 int32
	_ = v1324
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1342 int32
	_ = v1342
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1354 int64
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1361 int64
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1401 int32
	_ = v1401
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1415 int32
	_ = v1415
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1443 int32
	_ = v1443
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1464 int32
	_ = v1464
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1482 int64
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
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
	var v1499 int32
	_ = v1499
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
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
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1584 int32
	_ = v1584
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1605 int32
	_ = v1605
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1632 int32
	_ = v1632
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1650 int32
	_ = v1650
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
	var v1669 int64
	_ = v1669
	var v1670 int64
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1673 int32
	_ = v1673
	var v1676 int64
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1692 int32
	_ = v1692
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1714 int32
	_ = v1714
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1754 int32
	_ = v1754
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1769 int32
	_ = v1769
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1794 int32
	_ = v1794
	var v1799 int32
	_ = v1799
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1813 int32
	_ = v1813
	var v1816 int64
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1821 int64
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1833 int32
	_ = v1833
	var v1836 int32
	_ = v1836
	var v1839 int64
	_ = v1839
	var v1840 int64
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1853 int32
	_ = v1853
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1867 int64
	_ = v1867
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1882 int32
	_ = v1882
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1904 int32
	_ = v1904
	var v1909 int32
	_ = v1909
	var v1910 int64
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1918 int32
	_ = v1918
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1930 int32
	_ = v1930
	var v1935 int32
	_ = v1935
	var v1938 int64
	_ = v1938
	var v1943 int64
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1961 int32
	_ = v1961
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1982 int32
	_ = v1982
	var v1986 int32
	_ = v1986
	var v1991 int32
	_ = v1991
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2018 int32
	_ = v2018
	var v2025 int32
	_ = v2025
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2037 int32
	_ = v2037
	var v2042 int32
	_ = v2042
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2054 int32
	_ = v2054
	var v2071 int32
	_ = v2071
	var v2074 int32
	_ = v2074
	var v2077 int64
	_ = v2077
	var v2078 int64
	_ = v2078
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2126 int32
	_ = v2126
	var v2132 int32
	_ = v2132
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2138 int64
	_ = v2138
	var v2143 int32
	_ = v2143
	var v2147 int32
	_ = v2147
	var v2152 int32
	_ = v2152
	var v2157 int32
	_ = v2157
	var v2162 int32
	_ = v2162
	var v2165 int32
	_ = v2165
	var v2170 int32
	_ = v2170
	var v2175 int32
	_ = v2175
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2186 int32
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2194 int32
	_ = v2194
	var v2197 int32
	_ = v2197
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2211 int32
	_ = v2211
	var v2217 int32
	_ = v2217
	var v2219 int64
	_ = v2219
	var v2221 int32
	_ = v2221
	var v2226 int32
	_ = v2226
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2237 int32
	_ = v2237
	var v2238 int64
	_ = v2238
	var v2240 int64
	_ = v2240
	var v2242 int64
	_ = v2242
	var v2244 int64
	_ = v2244
	var v2246 int32
	_ = v2246
	var v2250 int32
	_ = v2250
	var v2252 int32
	_ = v2252
	var v2255 int32
	_ = v2255
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2267 int64
	_ = v2267
	var v2269 int64
	_ = v2269
	var v2271 int64
	_ = v2271
	var v2273 int64
	_ = v2273
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2296 int32
	_ = v2296
	var v2300 int32
	_ = v2300
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2318 int32
	_ = v2318
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2369 int32
	_ = v2369
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2380 int32
	_ = v2380
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2389 int32
	_ = v2389
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2403 int32
	_ = v2403
	var v2410 int32
	_ = v2410
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2422 int32
	_ = v2422
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2435 int32
	_ = v2435
	var v2439 int32
	_ = v2439
	var v2442 int64
	_ = v2442
	var v2445 int32
	_ = v2445
	var v2447 int32
	_ = v2447
	var v2448 int32
	_ = v2448
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2470 int64
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2477 int32
	_ = v2477
	var v2481 int32
	_ = v2481
	var v2484 int64
	_ = v2484
	var v2487 int32
	_ = v2487
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2493 int32
	_ = v2493
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2501 int32
	_ = v2501
	var v2506 int32
	_ = v2506
	var v2508 int32
	_ = v2508
	var v2510 int64
	_ = v2510
	var v2514 int32
	_ = v2514
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2545 int32
	_ = v2545
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2551 int32
	_ = v2551
	var v2555 int32
	_ = v2555
	var v2565 int64
	_ = v2565
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2569 int32
	_ = v2569
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2585 int32
	_ = v2585
	var v2588 int32
	_ = v2588
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2598 int32
	_ = v2598
	var v2602 int32
	_ = v2602
	var v2607 int32
	_ = v2607
	var v2616 int32
	_ = v2616
	var v2619 int32
	_ = v2619
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2631 int32
	_ = v2631
	var v2636 int32
	_ = v2636
	var v2640 int32
	_ = v2640
	var v2644 int32
	_ = v2644
	var v2649 int32
	_ = v2649
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2662 int32
	_ = v2662
	var v2667 int32
	_ = v2667
	var v2671 int32
	_ = v2671
	var v2674 int32
	_ = v2674
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2683 int32
	_ = v2683
	var v2688 int32
	_ = v2688
	var v2691 int32
	_ = v2691
	var v2697 int32
	_ = v2697
	var v2707 int64
	_ = v2707
	var v2711 int32
	_ = v2711
	var v2714 int32
	_ = v2714
	var v2723 int32
	_ = v2723
	var v2730 int32
	_ = v2730
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2743 int32
	_ = v2743
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2753 int32
	_ = v2753
	var v2758 int32
	_ = v2758
	var v2760 int64
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2763 int64
	_ = v2763
	var v2766 int32
	_ = v2766
	var v2777 int32
	_ = v2777
	var v2784 int32
	_ = v2784
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2797 int32
	_ = v2797
	var v2802 int32
	_ = v2802
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2817 int32
	_ = v2817
	var v2822 int32
	_ = v2822
	var v2825 int32
	_ = v2825
	var v2826 int32
	_ = v2826
	var v2827 int32
	_ = v2827
	var v2830 int32
	_ = v2830
	var v2832 int32
	_ = v2832
	var v2834 int64
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2838 int64
	_ = v2838
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2846 int32
	_ = v2846
	var v2853 int32
	_ = v2853
	var v2854 int64
	_ = v2854
	var v2855 int64
	_ = v2855
	var v2856 int64
	_ = v2856
	var v2859 int64
	_ = v2859
	var v2860 int64
	_ = v2860
	var v2862 int64
	_ = v2862
	var v2863 int64
	_ = v2863
	var v2866 int64
	_ = v2866
	var v2873 int64
	_ = v2873
	var v2875 int64
	_ = v2875
	var v2878 int32
	_ = v2878
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2900 int32
	_ = v2900
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2910 int32
	_ = v2910
	var v2912 int32
	_ = v2912
	var v2914 int64
	_ = v2914
	var v2915 int32
	_ = v2915
	var v2916 int64
	_ = v2916
	var v2919 int32
	_ = v2919
	var v2921 int32
	_ = v2921
	var v2922 int32
	_ = v2922
	var v2929 int32
	_ = v2929
	var v2930 int64
	_ = v2930
	var v2931 int64
	_ = v2931
	var v2932 int64
	_ = v2932
	var v2935 int64
	_ = v2935
	var v2936 int64
	_ = v2936
	var v2938 int64
	_ = v2938
	var v2939 int64
	_ = v2939
	var v2942 int64
	_ = v2942
	var v2952 int64
	_ = v2952
	var v2955 int32
	_ = v2955
	var v2966 int32
	_ = v2966
	var v2973 int32
	_ = v2973
	var v2976 int32
	_ = v2976
	var v2977 int32
	_ = v2977
	var v2978 int32
	_ = v2978
	var v2986 int32
	_ = v2986
	var v2991 int32
	_ = v2991
	var v2998 int32
	_ = v2998
	var v2999 int32
	_ = v2999
	var v3005 int32
	_ = v3005
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3015 int32
	_ = v3015
	var v3017 int32
	_ = v3017
	var v3019 int64
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3022 int64
	_ = v3022
	var v3026 int32
	_ = v3026
	var v3029 int64
	_ = v3029
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3039 int32
	_ = v3039
	var v3040 int32
	_ = v3040
	var v3041 int32
	_ = v3041
	var v3044 int32
	_ = v3044
	var v3051 int32
	_ = v3051
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3056 int32
	_ = v3056
	var v3057 int32
	_ = v3057
	var v3063 int32
	_ = v3063
	var v3068 int32
	_ = v3068
	var v3070 int64
	_ = v3070
	var v3072 int64
	_ = v3072
	var v3075 int32
	_ = v3075
	var v3084 int32
	_ = v3084
	var v3090 int32
	_ = v3090
	var v3093 int32
	_ = v3093
	var v3094 int32
	_ = v3094
	var v3095 int32
	_ = v3095
	var v3103 int32
	_ = v3103
	var v3108 int32
	_ = v3108
	var v3114 int32
	_ = v3114
	var v3115 int32
	_ = v3115
	var v3121 int32
	_ = v3121
	var v3126 int32
	_ = v3126
	var v3127 int32
	_ = v3127
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3144 int32
	_ = v3144
	var v3145 int32
	_ = v3145
	var v3146 int32
	_ = v3146
	var v3147 int32
	_ = v3147
	var v3153 int32
	_ = v3153
	var v3158 int32
	_ = v3158
	var v3161 int32
	_ = v3161
	var v3162 int32
	_ = v3162
	var v3163 int32
	_ = v3163
	var v3166 int32
	_ = v3166
	var v3168 int32
	_ = v3168
	var v3174 int32
	_ = v3174
	var v3178 int32
	_ = v3178
	var v3179 int32
	_ = v3179
	var v3181 int32
	_ = v3181
	var v3189 int32
	_ = v3189
	var v3193 int32
	_ = v3193
	var v3203 int32
	_ = v3203
	var v3215 int32
	_ = v3215
	var v3216 int32
	_ = v3216
	var v3222 int32
	_ = v3222
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3232 int32
	_ = v3232
	var v3233 int32
	_ = v3233
	var v3234 int32
	_ = v3234
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3247 int32
	_ = v3247
	var v3249 int32
	_ = v3249
	var v3250 int32
	_ = v3250
	var v3251 int32
	_ = v3251
	var v3255 int32
	_ = v3255
	var v3260 int32
	_ = v3260
	var v3262 int64
	_ = v3262
	var v3263 int32
	_ = v3263
	var v3265 int64
	_ = v3265
	var v3269 int32
	_ = v3269
	var v3272 int64
	_ = v3272
	var v3275 int32
	_ = v3275
	var v3276 int32
	_ = v3276
	var v3282 int32
	_ = v3282
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3287 int32
	_ = v3287
	var v3292 int64
	_ = v3292
	var v3294 int64
	_ = v3294
	var v3299 int64
	_ = v3299
	var v3301 int32
	_ = v3301
	var v3302 int32
	_ = v3302
	var v3307 int32
	_ = v3307
	var v3308 int32
	_ = v3308
	var v3317 int32
	_ = v3317
	var v3319 int32
	_ = v3319
	var v3321 int32
	_ = v3321
	var v3327 int32
	_ = v3327
	var v3328 int32
	_ = v3328
	var v3332 int32
	_ = v3332
	var v3335 int32
	_ = v3335
	var v3336 int32
	_ = v3336
	var v3337 int32
	_ = v3337
	var v3338 int32
	_ = v3338
	var v3344 int32
	_ = v3344
	var v3349 int32
	_ = v3349
	var v3350 int32
	_ = v3350
	var v3357 int32
	_ = v3357
	var v3360 int32
	_ = v3360
	var v3361 int32
	_ = v3361
	var v3362 int32
	_ = v3362
	var v3370 int32
	_ = v3370
	var v3375 int32
	_ = v3375
	var v3380 int32
	_ = v3380
	var v3381 int32
	_ = v3381
	var v3387 int32
	_ = v3387
	var v3392 int32
	_ = v3392
	var v3395 int32
	_ = v3395
	var v3396 int32
	_ = v3396
	var v3400 int32
	_ = v3400
	var v3401 int32
	_ = v3401
	var v3402 int32
	_ = v3402
	var v3407 int64
	_ = v3407
	var v3408 int64
	_ = v3408
	var v3409 int32
	_ = v3409
	var v3411 int32
	_ = v3411
	var v3414 int64
	_ = v3414
	var v3416 int32
	_ = v3416
	var v3421 float64
	_ = v3421
	var v3422 int32
	_ = v3422
	var v3423 int32
	_ = v3423
	var v3426 int32
	_ = v3426
	var v3432 int32
	_ = v3432
	var v3435 int32
	_ = v3435
	var v3436 int32
	_ = v3436
	var v3437 int32
	_ = v3437
	var v3438 int32
	_ = v3438
	var v3447 int32
	_ = v3447
	var v3452 int32
	_ = v3452
	var v3453 float64
	_ = v3453
	var v3460 int32
	_ = v3460
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3467 int32
	_ = v3467
	var v3468 int32
	_ = v3468
	var v3470 int32
	_ = v3470
	var v3473 int64
	_ = v3473
	var v3479 float64
	_ = v3479
	var v3480 int32
	_ = v3480
	var v3481 int32
	_ = v3481
	var v3484 int32
	_ = v3484
	var v3490 int32
	_ = v3490
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3496 int32
	_ = v3496
	var v3505 int32
	_ = v3505
	var v3510 int32
	_ = v3510
	var v3511 float64
	_ = v3511
	var v3518 int32
	_ = v3518
	var v3524 int32
	_ = v3524
	var v3530 int32
	_ = v3530
	var v3533 int32
	_ = v3533
	var v3534 int32
	_ = v3534
	var v3535 int32
	_ = v3535
	var v3536 int32
	_ = v3536
	var v3542 int32
	_ = v3542
	var v3547 int32
	_ = v3547
	var v3553 int64
	_ = v3553
	var v3554 int32
	_ = v3554
	var v3556 int32
	_ = v3556
	var v3557 int32
	_ = v3557
	var v3561 int32
	_ = v3561
	var v3566 int32
	_ = v3566
	var v3567 int32
	_ = v3567
	var v3571 int32
	_ = v3571
	var v3574 int32
	_ = v3574
	var v3575 int32
	_ = v3575
	var v3576 int32
	_ = v3576
	var v3577 int32
	_ = v3577
	var v3583 int32
	_ = v3583
	var v3588 int32
	_ = v3588
	var v3592 int32
	_ = v3592
	var v3595 int32
	_ = v3595
	var v3596 int32
	_ = v3596
	var v3597 int32
	_ = v3597
	var v3598 int32
	_ = v3598
	var v3604 int32
	_ = v3604
	var v3609 int32
	_ = v3609
	var v3611 int32
	_ = v3611
	var v3612 int32
	_ = v3612
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3619 int32
	_ = v3619
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3633 int32
	_ = v3633
	var v3635 int32
	_ = v3635
	var v3636 int32
	_ = v3636
	var v3643 int32
	_ = v3643
	var v3646 int32
	_ = v3646
	var v3647 int32
	_ = v3647
	var v3648 int32
	_ = v3648
	var v3649 int32
	_ = v3649
	var v3655 int32
	_ = v3655
	var v3660 int32
	_ = v3660
	var v3662 int64
	_ = v3662
	var v3665 int32
	_ = v3665
	var v3666 int32
	_ = v3666
	var v3671 int32
	_ = v3671
	var v3672 int32
	_ = v3672
	var v3675 int32
	_ = v3675
	var v3676 int32
	_ = v3676
	var v3677 int32
	_ = v3677
	var v3678 int32
	_ = v3678
	var v3680 int32
	_ = v3680
	var v3685 int32
	_ = v3685
	var v3686 int32
	_ = v3686
	var v3687 int32
	_ = v3687
	var v3688 int32
	_ = v3688
	var v3692 int32
	_ = v3692
	var v3694 int32
	_ = v3694
	var v3695 int32
	_ = v3695
	var v3698 int32
	_ = v3698
	var v3699 int32
	_ = v3699
	var v3701 int32
	_ = v3701
	var v3705 int32
	_ = v3705
	var v3706 int32
	_ = v3706
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3710 int64
	_ = v3710
	var v3712 int32
	_ = v3712
	var v3713 int32
	_ = v3713
	var v3716 int32
	_ = v3716
	var v3717 int32
	_ = v3717
	var v3718 int32
	_ = v3718
	var v3722 int32
	_ = v3722
	var v3723 int32
	_ = v3723
	var v3726 int32
	_ = v3726
	var v3727 int32
	_ = v3727
	var v3728 int32
	_ = v3728
	var v3729 int32
	_ = v3729
	var v3730 int32
	_ = v3730
	var v3733 int32
	_ = v3733
	var v3737 int32
	_ = v3737
	var v3738 int32
	_ = v3738
	var v3740 int32
	_ = v3740
	var v3742 int32
	_ = v3742
	var v3746 int32
	_ = v3746
	var v3747 int32
	_ = v3747
	var v3749 int32
	_ = v3749
	var v3750 int32
	_ = v3750
	var v3752 int32
	_ = v3752
	var v3753 int32
	_ = v3753
	var v3754 int32
	_ = v3754
	var v3760 int32
	_ = v3760
	var v3761 int32
	_ = v3761
	var v3765 int32
	_ = v3765
	var v3766 int32
	_ = v3766
	var v3767 int32
	_ = v3767
	var v3769 int32
	_ = v3769
	var v3776 int32
	_ = v3776
	var v3779 int32
	_ = v3779
	var v3783 int32
	_ = v3783
	var v3788 int32
	_ = v3788
	var v3792 int32
	_ = v3792
	var v3795 int32
	_ = v3795
	var v3796 int32
	_ = v3796
	var v3797 int32
	_ = v3797
	var v3798 int32
	_ = v3798
	var v3799 int32
	_ = v3799
	var v3805 int32
	_ = v3805
	var v3810 int32
	_ = v3810
	var v3812 int32
	_ = v3812
	var v3813 int32
	_ = v3813
	var v3814 int32
	_ = v3814
	var v3816 int32
	_ = v3816
	var v3817 int32
	_ = v3817
	var v3818 int32
	_ = v3818
	var v3820 int32
	_ = v3820
	var v3821 int32
	_ = v3821
	var v3825 int32
	_ = v3825
	var v3827 int32
	_ = v3827
	var v3833 int32
	_ = v3833
	var v3834 int32
	_ = v3834
	var v3835 int32
	_ = v3835
	var v3836 int32
	_ = v3836
	var v3837 int32
	_ = v3837
	var v3839 int32
	_ = v3839
	var v3840 int32
	_ = v3840
	var v3841 int32
	_ = v3841
	var v3842 int32
	_ = v3842
	var v3843 int32
	_ = v3843
	var v3850 int32
	_ = v3850
	var v3853 int32
	_ = v3853
	var v3857 int32
	_ = v3857
	var v3862 int32
	_ = v3862
	var v3864 int32
	_ = v3864
	var v3865 int64
	_ = v3865
	var v3867 int32
	_ = v3867
	var v3868 int32
	_ = v3868
	var v3870 int32
	_ = v3870
	var v3871 int32
	_ = v3871
	var v3872 int32
	_ = v3872
	var v3877 int32
	_ = v3877
	var v3878 int32
	_ = v3878
	var v3879 int32
	_ = v3879
	var v3884 int32
	_ = v3884
	var v3887 int32
	_ = v3887
	var v3888 int32
	_ = v3888
	var v3889 int32
	_ = v3889
	var v3890 int32
	_ = v3890
	var v3896 int32
	_ = v3896
	var v3901 int32
	_ = v3901
	var v3902 int32
	_ = v3902
	var v3905 int64
	_ = v3905
	var v3907 int64
	_ = v3907
	var v3909 int64
	_ = v3909
	var v3911 int64
	_ = v3911
	var v3913 int32
	_ = v3913
	var v3914 int32
	_ = v3914
	var v3919 int32
	_ = v3919
	var v3923 int32
	_ = v3923
	var v3925 int64
	_ = v3925
	var v3926 int32
	_ = v3926
	var v3933 int32
	_ = v3933
	var v3934 int32
	_ = v3934
	var v3935 int32
	_ = v3935
	var v3939 int32
	_ = v3939
	var v3940 int32
	_ = v3940
	var v3941 int32
	_ = v3941
	var v3942 int32
	_ = v3942
	var v3946 int32
	_ = v3946
	var v3947 int64
	_ = v3947
	var v3949 int64
	_ = v3949
	var v3951 int64
	_ = v3951
	var v3953 int64
	_ = v3953
	var v3955 int32
	_ = v3955
	var v3959 int32
	_ = v3959
	var v3961 int32
	_ = v3961
	var v3964 int32
	_ = v3964
	var v3969 int32
	_ = v3969
	var v3970 int32
	_ = v3970
	var v3976 int64
	_ = v3976
	var v3978 int64
	_ = v3978
	var v3980 int64
	_ = v3980
	var v3982 int64
	_ = v3982
	var v3997 int32
	_ = v3997
	var v4027 int32
	_ = v4027
	v6 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(1104)
	m.G0 = v25
	F_check_stack_depth(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[0]))
	if v32 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v35 {
	case 0, 1, 2, 3, 28:
		goto L45
	case 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 30, 41, 42:
		goto L44
	case 14:
		goto L33
	case 15:
		goto L43
	case 16:
		goto L42
	case 17:
		goto L41
	case 18:
		goto L40
	case 19:
		goto L39
	case 20:
		goto L38
	case 21:
		goto L37
	case 22:
		goto L36
	case 23:
		goto L35
	case 24:
		goto L12
	case 25:
		goto L13
	case 26:
		goto L14
	case 27:
		goto L15
	case 29:
		goto L16
	case 31:
		goto L17
	case 32:
		goto L18
	case 33:
		goto L19
	case 34:
		goto L20
	case 35:
		goto L21
	case 36:
		goto L22
	case 37, 45, 50, 51, 52, 53:
		goto L23
	case 38:
		goto L24
	default:
		goto L32
	case 40:
		goto L25
	case 43:
		goto L26
	case 44:
		goto L27
	case 46, 48:
		goto L28
	case 47:
		goto L29
	case 49:
		goto L30
	case 54, 55, 56, 57, 58, 59, 60, 61:
		goto L31
	}
L6:
	;
	goto L5
L7:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(3947), int32(_a_F_executeItemOptUnwrapTarget_1))
	mBase = m.M
	v4027 = m.ExcPending
	if v4027 != 0 {
		goto L1
	} else {
		goto L1224
	}
L8:
	;
	m.G0 = v25 + int32(1104)
	return v3997
L9:
	;
	v3926 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v3926 {
		goto L1209
	} else {
		goto L1210
	}
L10:
	;
	v3867 = v25 + int32(1000)
	if v3867 != 0 {
		goto L1191
	} else {
		goto L1192
	}
L11:
	;
	v3842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3842 != 0 {
		v3997 = v283
		goto L8
	} else {
		goto L1182
	}
L12:
	;
	v3812 = v25 + int32(1032)
	v3813 = F_jspGetNext(m, l1, v3812)
	mBase = m.M
	v3814 = m.ExcPending
	if v3814 != 0 {
		goto L1
	} else {
		goto L1172
	}
L13:
	;
	v3729 = F_JsonbType(m, l2)
	mBase = m.M
	v3730 = m.ExcPending
	if v3730 != 0 {
		goto L1
	} else {
		goto L1136
	}
L14:
	;
	v3726 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3727 = F_executeNextItem(m, l0, l1, int32(0), v3726, l3)
	mBase = m.M
	v3728 = m.ExcPending
	if v3728 != 0 {
		goto L1
	} else {
		goto L1134
	}
L15:
	;
	v3710 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v3712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3713 = *(*int32)(unsafe.Add(mBase, uint32(v3712)))
	if v3713 == int32(18) {
		goto L1130
	} else {
		goto L1131
	}
L16:
	;
	if l4 != 0 {
		goto L1120
	} else {
		goto L1121
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+864)) = int32(1)
	v3675 = F_JsonbTypeName(m, l2)
	mBase = m.M
	v3676 = m.ExcPending
	if v3676 != 0 {
		goto L1
	} else {
		goto L1116
	}
L18:
	;
	v3619 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v3619 != int32(18) {
		goto L1101
	} else {
		goto L1102
	}
L19:
	;
	v3617 = F_executeNumericItemMethod(m, l0, l1, l2, l4, int32(1404), l3)
	mBase = m.M
	v3618 = m.ExcPending
	if v3618 != 0 {
		goto L1
	} else {
		goto L1099
	}
L20:
	;
	v3614 = F_executeNumericItemMethod(m, l0, l1, l2, l4, int32(1570), l3)
	mBase = m.M
	v3615 = m.ExcPending
	if v3615 != 0 {
		goto L1
	} else {
		goto L1098
	}
L21:
	;
	v3611 = F_executeNumericItemMethod(m, l0, l1, l2, l4, int32(1569), l3)
	mBase = m.M
	v3612 = m.ExcPending
	if v3612 != 0 {
		goto L1
	} else {
		goto L1097
	}
L22:
	;
	if l4 == int32(0) {
		goto L1038
	} else {
		goto L1039
	}
L23:
	;
	v2348 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l4 == int32(0) {
		goto L738
	} else {
		goto L739
	}
L24:
	;
	if l4 == int32(0) {
		goto L655
	} else {
		goto L656
	}
L25:
	;
	v1955 = v25 + int32(1032)
	v1956 = F_jspGetNext(m, l1, v1955)
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L1
	} else {
		goto L645
	}
L26:
	;
	if l4 == int32(0) {
		goto L600
	} else {
		goto L601
	}
L27:
	;
	if l4 == int32(0) {
		goto L553
	} else {
		goto L554
	}
L28:
	;
	if l4 == int32(0) {
		goto L441
	} else {
		goto L442
	}
L29:
	;
	if l4 == int32(0) {
		goto L396
	} else {
		goto L397
	}
L30:
	;
	if l4 == int32(0) {
		goto L369
	} else {
		goto L370
	}
L31:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l4 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L226
	}
L33:
	;
	v627 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1547), l3)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L225
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L222
	}
L35:
	;
	v279 = F_JsonbType(m, l2)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L136
	}
L36:
	;
	v227 = F_JsonbType(m, l2)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L111
	}
L37:
	;
	v186 = F_JsonbType(m, l2)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L90
	}
L38:
	;
	v184 = F_executeUnaryArithmExpr(m, l0, l1, l2, int32(1546), l3)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L89
	}
L39:
	;
	v181 = F_executeUnaryArithmExpr(m, l0, l1, l2, int32(0), l3)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L88
	}
L40:
	;
	v178 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1545), l3)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L87
	}
L41:
	;
	v175 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1544), l3)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L86
	}
L42:
	;
	v172 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1543), l3)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L85
	}
L43:
	;
	v169 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1542), l3)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L84
	}
L44:
	;
	v88 = F_executeBoolItem(m, l0, l1, l2, int32(1))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L63
	}
L45:
	;
	v38 = F_jspGetNext(m, l1, v25+int32(1032))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l3|v38 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if v40 != int32(28) {
		v3997 = int32(0)
		goto L8
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v51 = l0 + int32(16)
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	switch v40 {
	case 0:
		goto L51
	case 1:
		goto L53
	case 2:
		goto L54
	case 3:
		goto L55
	default:
		goto L52
	case 28:
		v3864 = v51
		v3865 = v52
		goto L10
	}
L50:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v3864 = l0 + int32(16)
	v3865 = v49
	goto L10
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1072)) = int32(0)
	v3923 = v51
	v3925 = v52
	goto L9
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L60
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1072)) = int32(1)
	v67 = v25 + int32(1080)
	if v67 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1072)) = int32(2)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1080)) = v62
	v3923 = v51
	v3925 = v52
	goto L9
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1072)) = int32(3)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+1080)) = uint8(base.B2i32(v56 != int32(0)))
	v3923 = v51
	v3925 = v52
	goto L9
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1084)) = v70
	v3923 = v51
	v3925 = v52
	goto L9
L57:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v68
	goto L59
L58:
	;
	goto L59
L59:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L56
L60:
	;
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_2), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(3248), int32(_a_F_executeItemOptUnwrapTarget_3))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	v92 = F_jspGetNext(m, l1, v25+int32(1072))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if l3 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v96 = int32(0)
	if v92 == v96 {
		v3997 = v96
		goto L8
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v100 = int32(0)
	if v88 != int32(2) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L67
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+872)) = uint8(base.B2i32(v88 == int32(1)))
	v108 = int32(3)
	goto L71
L70:
	;
	v108 = int32(0)
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+864)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v110 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v118 = F_executeItemOptUnwrapTarget(m, l0, v25+int32(1072), v25+int32(864), l3, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if l3 == int32(0) {
		v3997 = v100
		goto L8
	} else {
		goto L76
	}
L75:
	;
	v3997 = v118
	goto L8
L76:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v123 < v124 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v128 = v122 + v123<<(uint(int32(5))%32)
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v25)+888))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+40)) = v129
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v25)+880))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+32)) = v131
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v25)+872))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+24)) = v133
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v25)+864))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+16)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v137 + int32(1)
	v3997 = v100
	goto L8
L78:
	;
	goto L79
L79:
	;
	v141 = int32(16)
	v143 = v124 << (uint(int32(1)) % 32)
	if v143 <= v141 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v146 = v141
	goto L82
L81:
	;
	v146 = v143
	goto L82
L82:
	;
	v151 = F_palloc(m, v146<<(uint(int32(5))%32)|int32(16))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(1)
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v25)+864))
	*(*int64)(unsafe.Add(mBase, uint32(v151)+16)) = v158
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v25)+872))
	*(*int64)(unsafe.Add(mBase, uint32(v151)+24)) = v160
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v25)+880))
	*(*int64)(unsafe.Add(mBase, uint32(v151)+32)) = v162
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v25)+888))
	*(*int64)(unsafe.Add(mBase, uint32(v151)+40)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v122)+8)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v151
	v3997 = v100
	goto L8
L84:
	;
	v3997 = v169
	goto L8
L85:
	;
	v3997 = v172
	goto L8
L86:
	;
	v3997 = v175
	goto L8
L87:
	;
	v3997 = v178
	goto L8
L88:
	;
	v3997 = v181
	goto L8
L89:
	;
	v3997 = v184
	goto L8
L90:
	;
	if v186 == int32(16) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v191 = v25 + int32(1032)
	v193 = F_jspGetNext(m, l1, v191)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v199 = int32(1)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v200 == v199 {
		goto L99
	} else {
		goto L100
	}
L94:
	;
	if v193 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v195 = v191
	goto L97
L96:
	;
	v195 = int32(0)
	goto L97
L97:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v197 = F_executeItemUnwrapTargetArray(m, l0, v195, l2, l3, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v3997 = v197
	goto L8
L99:
	;
	v204 = F_executeNextItem(m, l0, l1, int32(0), l2, l3)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v206 != 0 {
		v3997 = v199
		goto L8
	} else {
		goto L103
	}
L102:
	;
	v3997 = v204
	goto L8
L103:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v207 != int32(1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v3997 = int32(2)
	goto L8
L105:
	;
	goto L106
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errcode(m, int32(151781506))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_4), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(890), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	if v227 == int32(17) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v232 = v25 + int32(1032)
	v233 = F_jspGetNext(m, l1, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	if l4 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L115:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v235 != int32(18) {
		goto L34
	} else {
		goto L116
	}
L116:
	;
	if v233 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v239 = v232
	goto L119
L118:
	;
	v239 = int32(0)
	goto L119
L119:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v241 = int32(1)
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v246 = F_executeAnyItem(m, l0, v239, v240, l3, v241, v241, v241, int32(0), v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v3997 = v246
	goto L8
L121:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v257 != 0 {
		goto L126
	} else {
		goto L127
	}
L122:
	;
	v250 = F_JsonbType(m, l2)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	if v250 != int32(16) {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v255 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v3997 = v255
	goto L8
L126:
	;
	v3997 = int32(1)
	goto L8
L127:
	;
	goto L128
L128:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v259 != int32(1) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v3997 = int32(2)
	goto L8
L130:
	;
	goto L131
L131:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errcode(m, int32(319553666))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_6), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(913), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	if v279 != int32(16) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v283 = int32(1)
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v284 != v283 {
		goto L11
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v288 != int32(18) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	goto L139
L141:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v307 = F_jspGetNext(m, l1, v25+int32(1032))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L145
	}
L142:
	;
	v302 = int32(-1)
	goto L141
L143:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	if v292&int32(1342177280) != int32(1073741824) {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v302 = v292 & int32(268435455)
	goto L141
L145:
	;
	v311 = base.B2i32(v302 < int32(0))
	if v302 < int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v312 = int32(1)
	goto L148
L147:
	;
	v312 = v302
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v312
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if int32(0) < v314 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v318 = v312 - int32(1)
	v335 = v6
	goto L152
L150:
	;
	v591 = int32(1)
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v303
	v3997 = v591
	goto L8
L152:
	;
	v344 = int32(2)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v351 = v335 << (uint(int32(3)) % 32)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v351+v352)))
	F_jspInitByBuffer(m, v25+int32(864), v349, v354)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L154
	}
L153:
	;
	v591 = v565
	goto L151
L154:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v357+v351)+4))
	if v359 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v25+int32(1072), v360, v359)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v369 = F_getArrayIndex(m, l0, v25+int32(864), l2, v25+int32(1000))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L161
	}
L158:
	;
	goto L157
L159:
	;
	v584 = v335 + int32(1)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v584 < v585 {
		v335 = v584
		goto L152
	} else {
		goto L221
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v303
	v3997 = v545
	goto L8
L161:
	;
	if v369 == int32(2) {
		v545 = v344
		goto L160
	} else {
		goto L162
	}
L162:
	;
	if v359 != int32(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v385 == int32(1) {
		goto L170
	} else {
		goto L171
	}
L164:
	;
	v377 = F_getArrayIndex(m, l0, v25+int32(1072), l2, v25+int32(1016))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1000))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1016)) = v382
	v384 = v382
	goto L163
L167:
	;
	if v377 == int32(2) {
		v545 = v344
		goto L160
	} else {
		goto L168
	}
L168:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1000))
	v384 = v381
	goto L163
L169:
	;
	v416 = int32(1)
	v417 = int32(0)
	if v417 < v384 {
		goto L184
	} else {
		goto L185
	}
L170:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1016))
	v415 = v388
	goto L169
L171:
	;
	goto L172
L172:
	;
	if v384 < int32(0) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v395 != int32(1) {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1016))
	if v391 < v384 {
		goto L173
	} else {
		goto L175
	}
L175:
	;
	if v391 < v312 {
		v415 = v391
		goto L169
	} else {
		goto L176
	}
L176:
	;
	goto L173
L177:
	;
	v3997 = int32(2)
	goto L8
L178:
	;
	goto L179
L179:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	F_errcode(m, int32(51118210))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_7), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(962), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	v420 = v384
	goto L186
L185:
	;
	v420 = v417
	goto L186
L186:
	;
	if v415 < v318 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v422 = v415
	goto L189
L188:
	;
	v422 = v318
	goto L189
L189:
	;
	if v422 < v420 {
		v565 = v416
		goto L159
	} else {
		goto L190
	}
L190:
	;
	v428 = v416
	v430 = v420
	goto L191
L191:
	;
	if v311 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L192:
	;
	if l3|v531 != 0 {
		v565 = v531
		goto L159
	} else {
		goto L220
	}
L193:
	;
	goto L192
L194:
	;
	if v430 != v422 {
		v428 = v523
		v430 = v430 + int32(1)
		goto L191
	} else {
		goto L219
	}
L195:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v449 = F_getIthJsonbValueFromContainer(m, v448, v430)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	v453 = l2
	goto L197
L197:
	;
	if base.B2i32(l3 != int32(0))|v307 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	if v449 == int32(0) {
		v523 = v428
		goto L194
	} else {
		goto L199
	}
L199:
	;
	v453 = v449
	goto L197
L200:
	;
	v3997 = int32(0)
	goto L8
L201:
	;
	goto L202
L202:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v457 <= int32(0) {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	v496 = int32(16)
	v498 = v465 << (uint(int32(1)) % 32)
	if v498 <= v496 {
		goto L215
	} else {
		goto L216
	}
L204:
	;
	if l3|v491 != 0 {
		v523 = v491
		goto L194
	} else {
		goto L214
	}
L205:
	;
	if l3 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	goto L207
L207:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v487 = F_executeItemOptUnwrapTarget(m, l0, v25+int32(1032), v453, l3, v486)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L212
	}
L208:
	;
	v491 = int32(0)
	goto L204
L209:
	;
	goto L210
L210:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	if v465 <= v464 {
		goto L203
	} else {
		goto L211
	}
L211:
	;
	v469 = v463 + v464<<(uint(int32(5))%32)
	v470 = *(*int64)(unsafe.Add(mBase, uint32(v453)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v469)+40)) = v470
	v472 = *(*int64)(unsafe.Add(mBase, uint32(v453)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v469)+32)) = v472
	v474 = *(*int64)(unsafe.Add(mBase, uint32(v453)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v469)+24)) = v474
	v476 = *(*int64)(unsafe.Add(mBase, uint32(v453)))
	*(*int64)(unsafe.Add(mBase, uint32(v469)+16)) = v476
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v463)))
	*(*int32)(unsafe.Add(mBase, uint32(v463))) = v478 + int32(1)
	v523 = int32(0)
	goto L194
L212:
	;
	if v487 == int32(2) {
		v545 = int32(2)
		goto L160
	} else {
		goto L213
	}
L213:
	;
	v491 = v487
	goto L204
L214:
	;
	v531 = int32(0)
	goto L193
L215:
	;
	v501 = v496
	goto L217
L216:
	;
	v501 = v498
	goto L217
L217:
	;
	v506 = F_palloc(m, v501<<(uint(int32(5))%32)|int32(16))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v506)+4)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v506))) = int32(1)
	v513 = *(*int64)(unsafe.Add(mBase, uint32(v453)))
	*(*int64)(unsafe.Add(mBase, uint32(v506)+16)) = v513
	v515 = *(*int64)(unsafe.Add(mBase, uint32(v453)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v506)+24)) = v515
	v517 = *(*int64)(unsafe.Add(mBase, uint32(v453)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v506)+32)) = v517
	v519 = *(*int64)(unsafe.Add(mBase, uint32(v453)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v506)+40)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v463)+8)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v506
	v523 = int32(0)
	goto L194
L219:
	;
	v531 = v523
	goto L193
L220:
	;
	v545 = int32(0)
	goto L160
L221:
	;
	goto L153
L222:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v614
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_8), v25+int32(32))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(899), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L225:
	;
	v3997 = v627
	goto L8
L226:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v633
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_9), v25)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1690), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v719 = F_cstring_to_text_with_len(m, v717, v718)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L251
	}
L230:
	;
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v693 != int32(1) {
		v3997 = int32(2)
		goto L8
	} else {
		goto L245
	}
L231:
	;
	if v643 == int32(1) {
		goto L229
	} else {
		goto L244
	}
L232:
	;
	switch v643 - int32(16) {
	case 0:
		goto L234
	default:
		goto L231
	case 2:
		goto L235
	}
L233:
	;
	v681 = int32(1)
	v684 = int32(0)
	v686 = F_executeAnyItem(m, l0, l1, v648, l3, v681, v681, v681, v684, v684)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L243
	}
L234:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L240
	}
L235:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v648)))
	if v649&int32(536870912) != 0 {
		goto L230
	} else {
		goto L236
	}
L236:
	;
	if v649&int32(1073741824) != 0 {
		goto L233
	} else {
		goto L237
	}
L237:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v648)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+832)) = v658
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_10), v25+int32(832))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	goto L7
L240:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+848)) = v669
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_11), v25+int32(848))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1707), int32(_a_F_executeItemOptUnwrapTarget_12))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	v3997 = v686
	goto L8
L244:
	;
	goto L230
L245:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v704 = F_jspOperationName(m, v703)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+816)) = v704
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_13), v25+int32(816))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2918), int32(_a_F_executeItemOptUnwrapTarget_14))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	v721 = base.I64_extend_i32_u(v719)
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v724 - int32(54) {
	case 0:
		goto L267
	case 1:
		goto L259
	case 2:
		goto L266
	case 3:
		v771 = int32(1548)
		v772 = int32(1549)
		goto L263
	case 4:
		goto L265
	case 5:
		goto L264
	case 6:
		goto L262
	case 7:
		goto L261
	default:
		goto L260
	}
L252:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L366
	}
L253:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L363
	}
L254:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L1
	} else {
		goto L359
	}
L255:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L356
	}
L256:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L1
	} else {
		goto L353
	}
L257:
	;
	v3997 = int32(2)
	goto L8
L258:
	;
	v926 = F_jspGetNext(m, l1, v25+int32(1072))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L337
	}
L259:
	;
	v916 = F_DirectFunctionCall1Coll(m, int32(1557), int32(100), v721)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L335
	}
L260:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L1
	} else {
		goto L332
	}
L261:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1040)) = v808
	v811 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1032)) = v811
	v814 = v25 + int32(1072)
	F_jspGetArg(m, l1, v814)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L302
	}
L262:
	;
	v802 = F_DirectFunctionCall1Coll(m, int32(1555), int32(100), v721)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L300
	}
L263:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v773 != 0 {
		goto L286
	} else {
		goto L287
	}
L264:
	;
	v771 = int32(1553)
	v772 = int32(1554)
	goto L263
L265:
	;
	v771 = int32(1551)
	v772 = int32(1552)
	goto L263
L266:
	;
	v762 = F_DirectFunctionCall1Coll(m, int32(1550), int32(100), v721)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L284
	}
L267:
	;
	v728 = v25 + int32(1072)
	F_jspGetArg(m, l1, v728)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1072))
	if v731 != int32(1) {
		goto L256
	} else {
		goto L269
	}
L269:
	;
	goto L272
L270:
	;
	F_jspGetRightArg(m, l1, v728)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v728)+12))
	goto L270
L274:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1072))
	if v740 != int32(1) {
		goto L255
	} else {
		goto L275
	}
L275:
	;
	goto L278
L276:
	;
	v749 = F_cstring_to_text(m, v737)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v728)+12))
	goto L276
L280:
	;
	v752 = F_cstring_to_text(m, v746)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	v755 = F_DirectFunctionCall3Coll(m, int32(596), int32(100), v721, base.I64_extend_i32_u(v749), base.I64_extend_i32_u(v752))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	v758 = F_text_to_cstring(m, base.I32_wrap_i64(v755))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	v923 = v758
	goto L258
L284:
	;
	v765 = F_text_to_cstring(m, base.I32_wrap_i64(v762))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	v923 = v765
	goto L258
L286:
	;
	v775 = v25 + int32(1072)
	F_jspGetArg(m, l1, v775)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L1
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	v795 = F_DirectFunctionCall1Coll(m, v772, int32(100), v721)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L298
	}
L289:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1072))
	if v778 != int32(1) {
		goto L254
	} else {
		goto L290
	}
L290:
	;
	goto L293
L291:
	;
	v786 = F_cstring_to_text(m, v785)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v775)+12))
	goto L291
L295:
	;
	v789 = F_DirectFunctionCall2Coll(m, v771, int32(100), v721, base.I64_extend_i32_u(v786))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	v792 = F_text_to_cstring(m, base.I32_wrap_i64(v789))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	v923 = v792
	goto L258
L298:
	;
	v798 = F_text_to_cstring(m, base.I32_wrap_i64(v795))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	v923 = v798
	goto L258
L300:
	;
	v805 = F_text_to_cstring(m, base.I32_wrap_i64(v802))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	v923 = v805
	goto L258
L302:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1072))
	if v817 != int32(1) {
		goto L253
	} else {
		goto L303
	}
L303:
	;
	goto L306
L304:
	;
	F_jspGetRightArg(m, l1, v814)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L308
	}
L306:
	;
	goto L307
L307:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v814)+12))
	goto L304
L308:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1072))
	if v826 != int32(2) {
		goto L252
	} else {
		goto L309
	}
L309:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v814)+12))
	v832 = F_numeric_int4_safe(m, v829, v25+int32(1032))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1036)))
	if v834 == int32(1) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v837 != int32(1) {
		goto L257
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	if v832 == int32(0) {
		goto L320
	} else {
		goto L321
	}
L314:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v848 = F_jspOperationName(m, v847)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+768)) = v848
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_15), v25+int32(768))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(3027), int32(_a_F_executeItemOptUnwrapTarget_14))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L320:
	;
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v863 != int32(1) {
		goto L257
	} else {
		goto L323
	}
L321:
	;
	goto L322
L322:
	;
	v889 = F_cstring_to_text(m, v823)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L329
	}
L323:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v874 = F_jspOperationName(m, v873)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+784)) = v874
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_16), v25+int32(784))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(3033), int32(_a_F_executeItemOptUnwrapTarget_14))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
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
	v893 = F_DirectFunctionCall3Coll(m, int32(1556), int32(100), v721, base.I64_extend_i32_u(v889), base.I64_extend_i32_s(v832))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	v896 = F_text_to_cstring(m, base.I32_wrap_i64(v893))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	v923 = v896
	goto L258
L332:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+752)) = v902
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_17), v25+int32(752))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(3043), int32(_a_F_executeItemOptUnwrapTarget_14))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L335:
	;
	v919 = F_text_to_cstring(m, base.I32_wrap_i64(v916))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	v923 = v919
	goto L258
L337:
	;
	if v926|l3 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v3997 = base.B2i32(v923 == int32(0))
	goto L8
L339:
	;
	goto L340
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+876)) = v923
	*(*int32)(unsafe.Add(mBase, uint32(v25)+864)) = int32(1)
	v936 = F_strlen(m, v923)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v25)+872)) = v936
	v938 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v938 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v946 = F_executeItemOptUnwrapTarget(m, l0, v25+int32(1072), v25+int32(864), l3, v945)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	v948 = int32(0)
	if l3 == v948 {
		v3997 = v948
		goto L8
	} else {
		goto L345
	}
L344:
	;
	v3997 = v946
	goto L8
L345:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v951)))
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v951)+4))
	if v952 < v953 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v957 = v951 + v952<<(uint(int32(5))%32)
	v958 = *(*int64)(unsafe.Add(mBase, uint32(v25)+888))
	*(*int64)(unsafe.Add(mBase, uint32(v957)+40)) = v958
	v960 = *(*int64)(unsafe.Add(mBase, uint32(v25)+880))
	*(*int64)(unsafe.Add(mBase, uint32(v957)+32)) = v960
	v962 = *(*int64)(unsafe.Add(mBase, uint32(v25)+872))
	*(*int64)(unsafe.Add(mBase, uint32(v957)+24)) = v962
	v964 = *(*int64)(unsafe.Add(mBase, uint32(v25)+864))
	*(*int64)(unsafe.Add(mBase, uint32(v957)+16)) = v964
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v951)))
	*(*int32)(unsafe.Add(mBase, uint32(v951))) = v966 + int32(1)
	v3997 = v948
	goto L8
L347:
	;
	goto L348
L348:
	;
	v970 = int32(16)
	v972 = v953 << (uint(int32(1)) % 32)
	if v972 <= v970 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v975 = v970
	goto L351
L350:
	;
	v975 = v972
	goto L351
L351:
	;
	v980 = F_palloc(m, v975<<(uint(int32(5))%32)|int32(16))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v980)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v980)+4)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v980))) = int32(1)
	v987 = *(*int64)(unsafe.Add(mBase, uint32(v25)+864))
	*(*int64)(unsafe.Add(mBase, uint32(v980)+16)) = v987
	v989 = *(*int64)(unsafe.Add(mBase, uint32(v25)+872))
	*(*int64)(unsafe.Add(mBase, uint32(v980)+24)) = v989
	v991 = *(*int64)(unsafe.Add(mBase, uint32(v25)+880))
	*(*int64)(unsafe.Add(mBase, uint32(v980)+32)) = v991
	v993 = *(*int64)(unsafe.Add(mBase, uint32(v25)+888))
	*(*int64)(unsafe.Add(mBase, uint32(v980)+40)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v951)+8)) = v980
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v980
	v3997 = v948
	goto L8
L353:
	;
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_18), int32(0))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2932), int32(_a_F_executeItemOptUnwrapTarget_14))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L356:
	;
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_19), int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2938), int32(_a_F_executeItemOptUnwrapTarget_14))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L358
	}
L358:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L359:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1029 = F_jspOperationName(m, v1028)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L1
	} else {
		goto L360
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+800)) = v1029
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_20), v25+int32(800))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2987), int32(_a_F_executeItemOptUnwrapTarget_14))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L363:
	;
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_21), int32(0))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(3013), int32(_a_F_executeItemOptUnwrapTarget_14))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_21), int32(0))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L1
	} else {
		goto L367
	}
L367:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(3019), int32(_a_F_executeItemOptUnwrapTarget_14))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L1
	} else {
		goto L368
	}
L368:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L369:
	;
	v1077 = F_JsonbType(m, l2)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L381
	}
L370:
	;
	v1070 = F_JsonbType(m, l2)
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	if v1070 != int32(16) {
		goto L369
	} else {
		goto L372
	}
L372:
	;
	v1075 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L373
	}
L373:
	;
	v3997 = v1075
	goto L8
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1084)) = v1132
	v1134 = F_strlen(m, v1132)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1072)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1080)) = v1134
	v1141 = F_executeNextItem(m, l0, l1, int32(0), v25+int32(1072), l3)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L1
	} else {
		goto L395
	}
L375:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1129 = F_pnstrdup(m, v1127, v1128)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L1
	} else {
		goto L394
	}
L376:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L377:
	;
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1102 != int32(1) {
		v3997 = int32(2)
		goto L8
	} else {
		goto L388
	}
L378:
	;
	v1092 = v25 + int32(864)
	v1093 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1097 = F_JsonEncodeDateTime(m, v1092, v1093, v1094, l2+int32(24))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L1
	} else {
		goto L386
	}
L379:
	;
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v1089 != 0 {
		goto L383
	} else {
		goto L384
	}
L380:
	;
	v1083 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+8)))
	v1084 = F_DirectFunctionCall1Coll(m, int32(664), int32(0), v1083)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L1
	} else {
		goto L382
	}
L381:
	;
	switch v1077 - int32(1) {
	case 0:
		goto L375
	case 1:
		goto L380
	case 2:
		goto L379
	case 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30:
		goto L376
	default:
		goto L377
	case 31:
		goto L378
	}
L382:
	;
	v1132 = base.I32_wrap_i64(v1084)
	goto L374
L383:
	;
	v1090 = int32(_a_F_executeItemOptUnwrapTarget_22)
	goto L385
L384:
	;
	v1090 = int32(_a_F_executeItemOptUnwrapTarget_23)
	goto L385
L385:
	;
	v1132 = v1090
	goto L374
L386:
	;
	v1099 = F_pstrdup(m, v1092)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	v1132 = v1099
	goto L374
L388:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L390
	}
L390:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1113 = F_jspOperationName(m, v1112)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+736)) = v1113
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_24), v25+int32(736))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1660), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	goto L376
L394:
	;
	v1132 = v1129
	goto L374
L395:
	;
	v3997 = v1141
	goto L8
L396:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v1152 - int32(1) {
	case 0:
		goto L404
	case 1:
		goto L405
	default:
		goto L403
	}
L397:
	;
	v1145 = F_JsonbType(m, l2)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
	;
	if v1145 != int32(16) {
		goto L396
	} else {
		goto L399
	}
L399:
	;
	v1150 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
	;
	v3997 = v1150
	goto L8
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+864)) = int32(2)
	v1288 = F_DirectFunctionCall1Coll(m, int32(1538), int32(0), v1283)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L1
	} else {
		goto L438
	}
L402:
	;
	v1279 = base.I64_extend_i32_s(v1164)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1032)) = v1279
	v1283 = v1279
	goto L401
L403:
	;
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1255 != int32(1) {
		v3997 = int32(2)
		goto L8
	} else {
		goto L432
	}
L404:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1204 = F_pnstrdup(m, v1202, v1203)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L1
	} else {
		goto L417
	}
L405:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1080)) = v1156
	v1159 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1072)) = v1159
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1164 = F_numeric_int4_safe(m, v1161, v25+int32(1072))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1076)))
	if v1166 != int32(1) {
		goto L402
	} else {
		goto L407
	}
L407:
	;
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1169 == int32(0) {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v3997 = int32(2)
	goto L8
L409:
	;
	goto L410
L410:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	v1182 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+8)))
	v1183 = F_DirectFunctionCall1Coll(m, int32(664), int32(0), v1182)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1186 = F_jspOperationName(m, v1185)
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+712)) = int32(_a_F_executeItemOptUnwrapTarget_25)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+708)) = v1186
	*(*uint32)(unsafe.Add(mBase, uint32(v25)+704)) = uint32(v1183)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_26), v25+int32(704))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1576), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L417:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1080)) = v1207
	v1210 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1072)) = v1210
	v1218 = F_DirectInputFunctionCallSafe(m, int32(1142), v1204, int32(-1), v25+int32(1072), v25+int32(1032))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L1
	} else {
		goto L419
	}
L418:
	;
	v1253 = *(*int64)(unsafe.Add(mBase, uint32(v25)+1032))
	v1283 = v1253
	goto L401
L419:
	;
	if v1218 != 0 {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1076)))
	if v1220&int32(1) == int32(0) {
		goto L418
	} else {
		goto L423
	}
L421:
	;
	goto L422
L422:
	;
	v1225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1225 == int32(0) {
		goto L424
	} else {
		goto L425
	}
L423:
	;
	goto L422
L424:
	;
	v3997 = int32(2)
	goto L8
L425:
	;
	goto L426
L426:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L1
	} else {
		goto L427
	}
L427:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1237 = F_jspOperationName(m, v1236)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L1
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+728)) = int32(_a_F_executeItemOptUnwrapTarget_25)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+724)) = v1237
	*(*int32)(unsafe.Add(mBase, uint32(v25)+720)) = v1204
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_26), v25+int32(720))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1598), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L432:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L1
	} else {
		goto L433
	}
L433:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L1
	} else {
		goto L434
	}
L434:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1266 = F_jspOperationName(m, v1265)
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L1
	} else {
		goto L435
	}
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v1266
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_27), v25+int32(688))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L1
	} else {
		goto L436
	}
L436:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1606), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L1
	} else {
		goto L437
	}
L437:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L438:
	;
	v1291 = F_pg_detoast_datum(m, base.I32_wrap_i64(v1288))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L1
	} else {
		goto L439
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+872)) = v1291
	v1297 = F_executeNextItem(m, l0, l1, int32(0), v25+int32(864), l3)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L1
	} else {
		goto L440
	}
L440:
	;
	v3997 = v1297
	goto L8
L441:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v1308 - int32(1) {
	case 0:
		goto L453
	case 1:
		goto L454
	default:
		goto L452
	}
L442:
	;
	v1301 = F_JsonbType(m, l2)
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L1
	} else {
		goto L443
	}
L443:
	;
	if v1301 != int32(16) {
		goto L441
	} else {
		goto L444
	}
L444:
	;
	v1306 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	v3997 = v1306
	goto L8
L446:
	;
	v3997 = int32(2)
	goto L8
L447:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L1
	} else {
		goto L550
	}
L448:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L1
	} else {
		goto L547
	}
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+872)) = v1614
	*(*int32)(unsafe.Add(mBase, uint32(v25)+864)) = int32(2)
	v1623 = F_executeNextItem(m, l0, l1, int32(0), v25+int32(864), l3)
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L1
	} else {
		goto L546
	}
L450:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v1472 != int32(46) {
		v1614 = v1470
		goto L449
	} else {
		goto L502
	}
L451:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L1
	} else {
		goto L497
	}
L452:
	;
	v1425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1425 != int32(1) {
		goto L446
	} else {
		goto L491
	}
L453:
	;
	v1358 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1080)) = v1358
	v1361 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1072)) = v1361
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1366 = F_pnstrdup(m, v1364, v1365)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L1
	} else {
		goto L471
	}
L454:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1311)+4)))
	goto L456
L455:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v1348 != int32(46) {
		v1614 = v1311
		goto L449
	} else {
		goto L468
	}
L456:
	;
	if base.B2i32(v1312 == int32(_a_F_executeItemOptUnwrapTarget_28)) == int32(0) {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	v1317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1311)+4)))
	goto L460
L458:
	;
	goto L459
L459:
	;
	v1324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1324 != int32(1) {
		goto L446
	} else {
		goto L462
	}
L460:
	;
	if base.B2i32(v1317&int32(_a_F_executeItemOptUnwrapTarget_29) == int32(_a_F_executeItemOptUnwrapTarget_30)) == int32(0) {
		goto L455
	} else {
		goto L461
	}
L461:
	;
	goto L459
L462:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L1
	} else {
		goto L463
	}
L463:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1335 = F_jspOperationName(m, v1334)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+592)) = v1335
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_31), v25+int32(592))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L1
	} else {
		goto L466
	}
L466:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1440), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L1
	} else {
		goto L467
	}
L467:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L468:
	;
	v1354 = F_DirectFunctionCall1Coll(m, int32(664), int32(0), base.I64_extend_i32_u(v1311))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	v1470 = v1311
	v1471 = base.I32_wrap_i64(v1354)
	goto L450
L470:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1032))
	v1408 = F_pg_detoast_datum(m, v1407)
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L1
	} else {
		goto L483
	}
L471:
	;
	v1373 = F_DirectInputFunctionCallSafe(m, int32(434), v1366, int32(-1), v25+int32(1072), v25+int32(1032))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L1
	} else {
		goto L472
	}
L472:
	;
	if v1373 != 0 {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v1375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1076)))
	if v1375&int32(1) == int32(0) {
		goto L470
	} else {
		goto L476
	}
L474:
	;
	goto L475
L475:
	;
	v1380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1380 != int32(1) {
		goto L446
	} else {
		goto L477
	}
L476:
	;
	goto L475
L477:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L1
	} else {
		goto L478
	}
L478:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L1
	} else {
		goto L479
	}
L479:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1391 = F_jspOperationName(m, v1390)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = int32(_a_F_executeItemOptUnwrapTarget_32)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v1391
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v1366
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_26), v25+int32(656))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L1
	} else {
		goto L481
	}
L481:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1465), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L483:
	;
	v1410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1408)+4)))
	goto L484
L484:
	;
	if base.B2i32(v1410 == int32(_a_F_executeItemOptUnwrapTarget_28)) == int32(0) {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v1415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1408)+4)))
	goto L488
L486:
	;
	goto L487
L487:
	;
	v1422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1422 == int32(1) {
		goto L451
	} else {
		goto L490
	}
L488:
	;
	if base.B2i32(v1415&int32(_a_F_executeItemOptUnwrapTarget_29) == int32(_a_F_executeItemOptUnwrapTarget_30)) == int32(0) {
		v1470 = v1408
		v1471 = v1366
		goto L450
	} else {
		goto L489
	}
L489:
	;
	goto L487
L490:
	;
	goto L446
L491:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L1
	} else {
		goto L493
	}
L493:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1436 = F_jspOperationName(m, v1435)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+576)) = v1436
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_27), v25+int32(576))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L1
	} else {
		goto L495
	}
L495:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1481), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L497:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1457 = F_jspOperationName(m, v1456)
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v1457
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_31), v25+int32(672))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1472), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
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
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1475 == int32(0) {
		v1614 = v1470
		goto L449
	} else {
		goto L503
	}
L503:
	;
	v1479 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1080)) = v1479
	v1482 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1072)) = v1482
	v1485 = v25 + int32(1032)
	F_jspGetArg(m, l1, v1485)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L1
	} else {
		goto L504
	}
L504:
	;
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1032))
	if v1488 != int32(2) {
		goto L448
	} else {
		goto L505
	}
L505:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1485)+12))
	v1494 = F_numeric_int4_safe(m, v1491, v25+int32(1072))
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	v1496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1076)))
	if v1496 == int32(1) {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v1499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1499 != int32(1) {
		goto L446
	} else {
		goto L510
	}
L508:
	;
	goto L509
L509:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1523 == int32(0) {
		v1565 = v6
		goto L516
	} else {
		goto L517
	}
L510:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1510 = F_jspOperationName(m, v1509)
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L1
	} else {
		goto L513
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+608)) = v1510
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_33), v25+int32(608))
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1508), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L516:
	;
	v1568 = v25 + int32(1072)
	v1569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1569 != 0 {
		goto L528
	} else {
		goto L529
	}
L517:
	;
	v1527 = v25 + int32(1032)
	F_jspGetRightArg(m, l1, v1527)
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L1
	} else {
		goto L518
	}
L518:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1032))
	if v1530 != int32(2) {
		goto L447
	} else {
		goto L519
	}
L519:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+12))
	v1536 = F_numeric_int4_safe(m, v1533, v25+int32(1072))
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	v1538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1076)))
	if v1538 != int32(1) {
		v1565 = v1536
		goto L516
	} else {
		goto L521
	}
L521:
	;
	v1541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1541 != int32(1) {
		goto L446
	} else {
		goto L522
	}
L522:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L1
	} else {
		goto L523
	}
L523:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1552 = F_jspOperationName(m, v1551)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+640)) = v1552
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_34), v25+int32(640))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L1
	} else {
		goto L526
	}
L526:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1522), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L1
	} else {
		goto L527
	}
L527:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L528:
	;
	v1570 = int32(0)
	goto L530
L529:
	;
	v1570 = v1568
	goto L530
L530:
	;
	v1571 = F_make_numeric_typmod_safe(m, v1494, v1565, v1570)
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	v1573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1076)))
	if v1573 != 0 {
		goto L446
	} else {
		goto L532
	}
L532:
	;
	v1577 = F_DirectInputFunctionCallSafe(m, int32(434), v1471, v1571, v1568, v25+int32(1000))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L1
	} else {
		goto L534
	}
L533:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1000))
	v1612 = F_pg_detoast_datum(m, v1611)
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L1
	} else {
		goto L545
	}
L534:
	;
	if v1577 != 0 {
		goto L535
	} else {
		goto L536
	}
L535:
	;
	v1579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1076)))
	if v1579&int32(1) == int32(0) {
		goto L533
	} else {
		goto L538
	}
L536:
	;
	goto L537
L537:
	;
	v1584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1584 != int32(1) {
		goto L446
	} else {
		goto L539
	}
L538:
	;
	goto L537
L539:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L1
	} else {
		goto L541
	}
L541:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1595 = F_jspOperationName(m, v1594)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+632)) = int32(_a_F_executeItemOptUnwrapTarget_32)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+628)) = v1595
	*(*int32)(unsafe.Add(mBase, uint32(v25)+624)) = v1471
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_26), v25+int32(624))
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L1
	} else {
		goto L543
	}
L543:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1542), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L1
	} else {
		goto L544
	}
L544:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L545:
	;
	v1614 = v1612
	goto L449
L546:
	;
	v3997 = v1623
	goto L8
L547:
	;
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_35), int32(0))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1500), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L1
	} else {
		goto L549
	}
L549:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L550:
	;
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_36), int32(0))
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L1
	} else {
		goto L551
	}
L551:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1514), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L1
	} else {
		goto L552
	}
L552:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L553:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v1664 - int32(1) {
	case 0:
		goto L563
	case 1:
		goto L564
	case 2:
		goto L561
	default:
		goto L562
	}
L554:
	;
	v1657 = F_JsonbType(m, l2)
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L1
	} else {
		goto L555
	}
L555:
	;
	if v1657 != int32(16) {
		goto L553
	} else {
		goto L556
	}
L556:
	;
	v1662 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L1
	} else {
		goto L557
	}
L557:
	;
	v3997 = v1662
	goto L8
L558:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L1
	} else {
		goto L595
	}
L559:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+872)) = uint8(v1769)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+864)) = int32(3)
	v1777 = F_executeNextItem(m, l0, l1, int32(0), v25+int32(864), l3)
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L1
	} else {
		goto L594
	}
L560:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+1000)) = uint8(v1767)
	v1769 = v1767
	goto L559
L561:
	;
	v1764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v1767 = v1764
	goto L560
L562:
	;
	v1760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1760 == int32(1) {
		goto L558
	} else {
		goto L593
	}
L563:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1725 = F_pnstrdup(m, v1723, v1724)
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L1
	} else {
		goto L580
	}
L564:
	;
	v1669 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+8)))
	v1670 = F_DirectFunctionCall1Coll(m, int32(664), int32(0), v1669)
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L1
	} else {
		goto L565
	}
L565:
	;
	v1673 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1080)) = v1673
	v1676 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1072)) = v1676
	v1679 = base.I32_wrap_i64(v1670)
	v1685 = F_DirectInputFunctionCallSafe(m, int32(1142), v1679, int32(-1), v25+int32(1072), v25+int32(1032))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L1
	} else {
		goto L567
	}
L566:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1032))
	v1767 = base.B2i32(v1720 != int32(0))
	goto L560
L567:
	;
	if v1685 != 0 {
		goto L568
	} else {
		goto L569
	}
L568:
	;
	v1687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1076)))
	if v1687&int32(1) == int32(0) {
		goto L566
	} else {
		goto L571
	}
L569:
	;
	goto L570
L570:
	;
	v1692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1692 == int32(0) {
		goto L572
	} else {
		goto L573
	}
L571:
	;
	goto L570
L572:
	;
	v3997 = int32(2)
	goto L8
L573:
	;
	goto L574
L574:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L1
	} else {
		goto L575
	}
L575:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L1
	} else {
		goto L576
	}
L576:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1704 = F_jspOperationName(m, v1703)
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+552)) = int32(_a_F_executeItemOptUnwrapTarget_37)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+548)) = v1704
	*(*int32)(unsafe.Add(mBase, uint32(v25)+544)) = v1679
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_26), v25+int32(544))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1384), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L580:
	;
	v1729 = F_strlen(m, v1725)
	mBase = m.M
	v1730 = F_parse_bool_with_len(m, v1725, v1729, v25+int32(1000))
	mBase = m.M
	goto L581
L581:
	;
	if v1730 != 0 {
		goto L582
	} else {
		goto L583
	}
L582:
	;
	v1731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1000)))
	v1769 = v1731
	goto L559
L583:
	;
	goto L584
L584:
	;
	v1732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1732 != int32(1) {
		goto L585
	} else {
		goto L586
	}
L585:
	;
	v3997 = int32(2)
	goto L8
L586:
	;
	goto L587
L587:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1744 = F_jspOperationName(m, v1743)
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L1
	} else {
		goto L590
	}
L590:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+568)) = int32(_a_F_executeItemOptUnwrapTarget_37)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+564)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v25)+560)) = v1725
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_26), v25+int32(560))
	mBase = m.M
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1404), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L593:
	;
	v3997 = int32(2)
	goto L8
L594:
	;
	v3997 = v1777
	goto L8
L595:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1787 = F_jspOperationName(m, v1786)
	mBase = m.M
	v1788 = m.ExcPending
	if v1788 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+528)) = v1787
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_38), v25+int32(528))
	mBase = m.M
	v1794 = m.ExcPending
	if v1794 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1413), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L600:
	;
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v1809 - int32(1) {
	case 0:
		goto L608
	case 1:
		goto L609
	default:
		goto L607
	}
L601:
	;
	v1802 = F_JsonbType(m, l2)
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L1
	} else {
		goto L602
	}
L602:
	;
	if v1802 != int32(16) {
		goto L600
	} else {
		goto L603
	}
L603:
	;
	v1807 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L1
	} else {
		goto L604
	}
L604:
	;
	v3997 = v1807
	goto L8
L605:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+864)) = int32(2)
	v1943 = F_DirectFunctionCall1Coll(m, int32(1539), int32(0), v1938)
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L1
	} else {
		goto L642
	}
L606:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1032)) = v1821
	v1938 = v1821
	goto L605
L607:
	;
	v1912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1912 != int32(1) {
		v3997 = int32(2)
		goto L8
	} else {
		goto L636
	}
L608:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1861 = F_pnstrdup(m, v1859, v1860)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L1
	} else {
		goto L621
	}
L609:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1080)) = v1813
	v1816 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1072)) = v1816
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1821 = F_numeric_int8_safe(m, v1818, v25+int32(1072))
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	v1823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1076)))
	if v1823 != int32(1) {
		goto L606
	} else {
		goto L611
	}
L611:
	;
	v1826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1826 == int32(0) {
		goto L612
	} else {
		goto L613
	}
L612:
	;
	v3997 = int32(2)
	goto L8
L613:
	;
	goto L614
L614:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1836 = m.ExcPending
	if v1836 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	v1839 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+8)))
	v1840 = F_DirectFunctionCall1Coll(m, int32(664), int32(0), v1839)
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1843 = F_jspOperationName(m, v1842)
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L1
	} else {
		goto L618
	}
L618:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+504)) = int32(_a_F_executeItemOptUnwrapTarget_39)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+500)) = v1843
	*(*uint32)(unsafe.Add(mBase, uint32(v25)+496)) = uint32(v1840)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_26), v25+int32(496))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L1
	} else {
		goto L619
	}
L619:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1311), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L1
	} else {
		goto L620
	}
L620:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L621:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1080)) = v1864
	v1867 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1072)) = v1867
	v1875 = F_DirectInputFunctionCallSafe(m, int32(587), v1861, int32(-1), v25+int32(1072), v25+int32(1032))
	mBase = m.M
	v1876 = m.ExcPending
	if v1876 != 0 {
		goto L1
	} else {
		goto L623
	}
L622:
	;
	v1910 = *(*int64)(unsafe.Add(mBase, uint32(v25)+1032))
	v1938 = v1910
	goto L605
L623:
	;
	if v1875 != 0 {
		goto L624
	} else {
		goto L625
	}
L624:
	;
	v1877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1076)))
	if v1877&int32(1) == int32(0) {
		goto L622
	} else {
		goto L627
	}
L625:
	;
	goto L626
L626:
	;
	v1882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1882 == int32(0) {
		goto L628
	} else {
		goto L629
	}
L627:
	;
	goto L626
L628:
	;
	v3997 = int32(2)
	goto L8
L629:
	;
	goto L630
L630:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L1
	} else {
		goto L631
	}
L631:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L1
	} else {
		goto L632
	}
L632:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1894 = F_jspOperationName(m, v1893)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L1
	} else {
		goto L633
	}
L633:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+520)) = int32(_a_F_executeItemOptUnwrapTarget_39)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+516)) = v1894
	*(*int32)(unsafe.Add(mBase, uint32(v25)+512)) = v1861
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_26), v25+int32(512))
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L1
	} else {
		goto L634
	}
L634:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1333), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L1
	} else {
		goto L635
	}
L635:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L636:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L1
	} else {
		goto L637
	}
L637:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L1
	} else {
		goto L638
	}
L638:
	;
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1923 = F_jspOperationName(m, v1922)
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+480)) = v1923
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_27), v25+int32(480))
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L1
	} else {
		goto L640
	}
L640:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1341), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L1
	} else {
		goto L641
	}
L641:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L642:
	;
	v1946 = F_pg_detoast_datum(m, base.I32_wrap_i64(v1943))
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L1
	} else {
		goto L643
	}
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+872)) = v1946
	v1952 = F_executeNextItem(m, l0, l1, int32(0), v25+int32(864), l3)
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L1
	} else {
		goto L644
	}
L644:
	;
	v3997 = v1952
	goto L8
L645:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(0) <= v1958 {
		goto L646
	} else {
		goto L647
	}
L646:
	;
	v1961 = int32(0)
	if base.B2i32(l3 == v1961)&(v1956^int32(1)) != 0 {
		v3997 = v1961
		goto L8
	} else {
		goto L649
	}
L647:
	;
	goto L648
L648:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L1
	} else {
		goto L652
	}
L649:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+864)) = int32(2)
	v1972 = F_int64_to_numeric(m, base.I64_extend_i32_s(v1958-int32(1)))
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L1
	} else {
		goto L650
	}
L650:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+872)) = v1972
	v1977 = F_executeNextItem(m, l0, l1, v1955, v25+int32(864), l3)
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L1
	} else {
		goto L651
	}
L651:
	;
	v3997 = v1977
	goto L8
L652:
	;
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_40), int32(0))
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L1
	} else {
		goto L653
	}
L653:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1270), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L1
	} else {
		goto L654
	}
L654:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L655:
	;
	v2001 = m.G0
	v2003 = v2001 - int32(320)
	m.G0 = v2003
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v2005 == int32(18) {
		goto L664
	} else {
		goto L665
	}
L656:
	;
	v1994 = F_JsonbType(m, l2)
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L1
	} else {
		goto L657
	}
L657:
	;
	if v1994 != int32(16) {
		goto L655
	} else {
		goto L658
	}
L658:
	;
	v1999 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
	;
	v3997 = v1999
	goto L8
L660:
	;
	v3997 = v2318
	goto L8
L661:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2342 = m.ExcPending
	if v2342 != 0 {
		goto L1
	} else {
		goto L734
	}
L662:
	;
	m.G0 = v2003 + int32(320)
	goto L660
L663:
	;
	if v2009&int32(268435455) == int32(0) {
		goto L677
	} else {
		goto L678
	}
L664:
	;
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v2008)))
	if v2009&int32(536870912) != 0 {
		goto L663
	} else {
		goto L667
	}
L665:
	;
	goto L666
L666:
	;
	v2018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2018 != int32(1) {
		goto L669
	} else {
		goto L670
	}
L667:
	;
	if v2009&int32(1073741824) == int32(0) {
		goto L661
	} else {
		goto L668
	}
L668:
	;
	goto L666
L669:
	;
	v2318 = int32(2)
	goto L662
L670:
	;
	goto L671
L671:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L1
	} else {
		goto L672
	}
L672:
	;
	F_errcode(m, int32(319553666))
	mBase = m.M
	v2028 = m.ExcPending
	if v2028 != 0 {
		goto L1
	} else {
		goto L673
	}
L673:
	;
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2030 = F_jspOperationName(m, v2029)
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L1
	} else {
		goto L674
	}
L674:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+16)) = v2030
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_41), v2003+int32(16))
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L1
	} else {
		goto L675
	}
L675:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(3106), int32(_a_F_executeItemOptUnwrapTarget_42))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L1
	} else {
		goto L676
	}
L676:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L677:
	;
	v2318 = int32(1)
	goto L662
L678:
	;
	goto L679
L679:
	;
	v2050 = F_jspGetNext(m, l1, v2003+int32(292))
	mBase = m.M
	v2051 = m.ExcPending
	if v2051 != 0 {
		goto L1
	} else {
		goto L680
	}
L680:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+172)) = int32(_a_F_executeItemOptUnwrapTarget_43)
	v2054 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+160)) = v2054
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+168)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+140)) = int32(_a_F_executeItemOptUnwrapTarget_44)
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+128)) = v2054
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+136)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+108)) = int32(_a_F_executeItemOptUnwrapTarget_45)
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+96)) = v2054
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+104)) = int32(2)
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v2071 == int32(18) {
		goto L681
	} else {
		goto L682
	}
L681:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v2077 = base.I64_extend_i32_s(v2008 - v2074)
	goto L683
L682:
	;
	v2077 = int64(0)
	goto L683
L683:
	;
	v2078 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+192)) = int32(2)
	v2084 = F_int64_to_numeric(m, v2078*int64(10000000000)+v2077)
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L1
	} else {
		goto L684
	}
L684:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+200)) = v2084
	v2087 = F_JsonbIteratorInit(m, v2008)
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L1
	} else {
		goto L685
	}
L685:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+92)) = v2087
	v2095 = F_JsonbIteratorNext(m, v2003+int32(92), v2003+int32(256), int32(1))
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	if v2095 == int32(0) {
		v2318 = v2054
		goto L662
	} else {
		goto L687
	}
L687:
	;
	v2104 = v2095
	v2106 = v2054
	goto L688
L688:
	;
	if v2104 != int32(1) {
		v2300 = v2106
		goto L691
	} else {
		goto L692
	}
L689:
	;
	v2318 = v2310
	goto L662
L690:
	;
	goto L689
L691:
	;
	v2308 = F_JsonbIteratorNext(m, v2003+int32(92), v2003+int32(256), int32(1))
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L1
	} else {
		goto L732
	}
L692:
	;
	v2126 = int32(0)
	if base.B2i32(l3 != int32(0))|v2050 == v2126 {
		v2310 = v2126
		goto L690
	} else {
		goto L693
	}
L693:
	;
	v2132 = v2003 + int32(224)
	v2134 = F_JsonbIteratorNext(m, v2003+int32(92), v2132, int32(1))
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L1
	} else {
		goto L694
	}
L694:
	;
	v2136 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+48)) = v2136
	v2138 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2003)+40)) = v2138
	*(*int64)(unsafe.Add(mBase, uint32(v2003)+32)) = v2138
	v2143 = v2003 + int32(32)
	F_pushJsonbValue(m, v2143, int32(6), v2136)
	mBase = m.M
	v2147 = m.ExcPending
	if v2147 != 0 {
		goto L1
	} else {
		goto L695
	}
L695:
	;
	F_pushJsonbValue(m, v2143, int32(1), v2003+int32(160))
	mBase = m.M
	v2152 = m.ExcPending
	if v2152 != 0 {
		goto L1
	} else {
		goto L696
	}
L696:
	;
	F_pushJsonbValue(m, v2143, int32(2), v2003+int32(256))
	mBase = m.M
	v2157 = m.ExcPending
	if v2157 != 0 {
		goto L1
	} else {
		goto L697
	}
L697:
	;
	F_pushJsonbValue(m, v2143, int32(1), v2003+int32(128))
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L1
	} else {
		goto L698
	}
L698:
	;
	F_pushJsonbValue(m, v2143, int32(2), v2132)
	mBase = m.M
	v2165 = m.ExcPending
	if v2165 != 0 {
		goto L1
	} else {
		goto L699
	}
L699:
	;
	F_pushJsonbValue(m, v2143, int32(1), v2003+int32(96))
	mBase = m.M
	v2170 = m.ExcPending
	if v2170 != 0 {
		goto L1
	} else {
		goto L700
	}
L700:
	;
	F_pushJsonbValue(m, v2143, int32(2), v2003+int32(192))
	mBase = m.M
	v2175 = m.ExcPending
	if v2175 != 0 {
		goto L1
	} else {
		goto L701
	}
L701:
	;
	F_pushJsonbValue(m, v2143, int32(7), int32(0))
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L1
	} else {
		goto L702
	}
L702:
	;
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v2003)+32))
	v2181 = F_JsonbValueToJsonb(m, v2180)
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L1
	} else {
		goto L703
	}
L703:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+56)) = int32(18)
	v2186 = v2181 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+68)) = v2186
	v2188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2181))))
	if v2188 == int32(1) {
		goto L705
	} else {
		goto L706
	}
L704:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2003)+64)) = v2217
	v2219 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2186
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2221
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v2221 + int32(1)
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v2226 <= int32(0) {
		goto L716
	} else {
		goto L717
	}
L705:
	;
	v2194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2181)+1)))
	if v2194 == int32(18) {
		goto L708
	} else {
		goto L709
	}
L706:
	;
	goto L707
L707:
	;
	v2205 = int32(1)
	if v2188&v2205 != 0 {
		v2217 = int32(base.Ui32(v2188)>>(uint(v2205)%32)) - v2205
		goto L704
	} else {
		goto L714
	}
L708:
	;
	v2197 = int32(16)
	goto L710
L709:
	;
	v2197 = int32(0)
	goto L710
L710:
	;
	if base.Ui32((v2194-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L711
	} else {
		goto L712
	}
L711:
	;
	v2204 = int32(4)
	goto L713
L712:
	;
	v2204 = v2197
	goto L713
L713:
	;
	v2217 = v2204
	goto L704
L714:
	;
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(v2181)))
	v2217 = int32(base.Ui32(v2211)>>(uint(int32(2))%32)) - int32(4)
	goto L704
L715:
	;
	if l3 != 0 {
		v2300 = v2294
		goto L691
	} else {
		goto L730
	}
L716:
	;
	if l3 == int32(0) {
		goto L719
	} else {
		goto L720
	}
L717:
	;
	goto L718
L718:
	;
	v2286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v2287 = F_executeItemOptUnwrapTarget(m, l0, v2003+int32(292), v2003+int32(56), l3, v2286)
	mBase = m.M
	v2288 = m.ExcPending
	if v2288 != 0 {
		goto L1
	} else {
		goto L728
	}
L719:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v2219
	v2294 = int32(0)
	goto L715
L720:
	;
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(v2231)))
	v2233 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+4))
	if v2232 < v2233 {
		goto L721
	} else {
		goto L722
	}
L721:
	;
	v2237 = v2231 + v2232<<(uint(int32(5))%32)
	v2238 = *(*int64)(unsafe.Add(mBase, uint32(v2003)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v2237)+40)) = v2238
	v2240 = *(*int64)(unsafe.Add(mBase, uint32(v2003)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v2237)+32)) = v2240
	v2242 = *(*int64)(unsafe.Add(mBase, uint32(v2003)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v2237)+24)) = v2242
	v2244 = *(*int64)(unsafe.Add(mBase, uint32(v2003)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v2237)+16)) = v2244
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v2231)))
	*(*int32)(unsafe.Add(mBase, uint32(v2231))) = v2246 + int32(1)
	goto L719
L722:
	;
	goto L723
L723:
	;
	v2250 = int32(16)
	v2252 = v2233 << (uint(int32(1)) % 32)
	if v2252 <= v2250 {
		goto L724
	} else {
		goto L725
	}
L724:
	;
	v2255 = v2250
	goto L726
L725:
	;
	v2255 = v2252
	goto L726
L726:
	;
	v2260 = F_palloc(m, v2255<<(uint(int32(5))%32)|int32(16))
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L1
	} else {
		goto L727
	}
L727:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2260)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2260)+4)) = v2255
	*(*int32)(unsafe.Add(mBase, uint32(v2260))) = int32(1)
	v2267 = *(*int64)(unsafe.Add(mBase, uint32(v2003)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v2260)+16)) = v2267
	v2269 = *(*int64)(unsafe.Add(mBase, uint32(v2003)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v2260)+24)) = v2269
	v2271 = *(*int64)(unsafe.Add(mBase, uint32(v2003)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v2260)+32)) = v2271
	v2273 = *(*int64)(unsafe.Add(mBase, uint32(v2003)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v2260)+40)) = v2273
	*(*int32)(unsafe.Add(mBase, uint32(v2231)+8)) = v2260
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v2260
	goto L719
L728:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v2219
	v2290 = int32(2)
	if v2287 == v2290 {
		v2310 = v2290
		goto L690
	} else {
		goto L729
	}
L729:
	;
	v2294 = v2287
	goto L715
L730:
	;
	v2296 = int32(0)
	if v2294 == v2296 {
		v2310 = v2296
		goto L690
	} else {
		goto L731
	}
L731:
	;
	v2300 = v2294
	goto L691
L732:
	;
	if v2308 != 0 {
		v2104 = v2308
		v2106 = v2300
		goto L688
	} else {
		goto L733
	}
L733:
	;
	v2318 = v2300
	goto L662
L734:
	;
	v2343 = *(*int32)(unsafe.Add(mBase, uint32(v2008)))
	*(*int32)(unsafe.Add(mBase, uint32(v2003))) = v2343
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_10), v2003)
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L1
	} else {
		goto L735
	}
L735:
	;
	goto L7
L736:
	;
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v2429 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2430 = F_cstring_to_text_with_len(m, v2428, v2429)
	mBase = m.M
	v2431 = m.ExcPending
	if v2431 != 0 {
		goto L1
	} else {
		goto L760
	}
L737:
	;
	v2403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2403 != int32(1) {
		goto L752
	} else {
		goto L753
	}
L738:
	;
	v2393 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1064)) = v2393
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1060)) = int32(0)
	if v2348 == int32(1) {
		goto L736
	} else {
		goto L751
	}
L739:
	;
	switch v2348 - int32(16) {
	case 0:
		goto L741
	default:
		goto L738
	case 2:
		goto L742
	}
L740:
	;
	v2386 = int32(1)
	v2389 = int32(0)
	v2391 = F_executeAnyItem(m, l0, l1, v2353, l3, v2386, v2386, v2386, v2389, v2389)
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L1
	} else {
		goto L750
	}
L741:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2373 = m.ExcPending
	if v2373 != 0 {
		goto L1
	} else {
		goto L747
	}
L742:
	;
	v2353 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v2354 = *(*int32)(unsafe.Add(mBase, uint32(v2353)))
	if v2354&int32(536870912) != 0 {
		goto L737
	} else {
		goto L743
	}
L743:
	;
	if v2354&int32(1073741824) != 0 {
		goto L740
	} else {
		goto L744
	}
L744:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		goto L1
	} else {
		goto L745
	}
L745:
	;
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(v2353)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+448)) = v2363
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_10), v25+int32(448))
	mBase = m.M
	v2369 = m.ExcPending
	if v2369 != 0 {
		goto L1
	} else {
		goto L746
	}
L746:
	;
	goto L7
L747:
	;
	v2374 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+464)) = v2374
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_11), v25+int32(464))
	mBase = m.M
	v2380 = m.ExcPending
	if v2380 != 0 {
		goto L1
	} else {
		goto L748
	}
L748:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1707), int32(_a_F_executeItemOptUnwrapTarget_12))
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L1
	} else {
		goto L749
	}
L749:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L750:
	;
	v3997 = v2391
	goto L8
L751:
	;
	goto L737
L752:
	;
	v3997 = int32(2)
	goto L8
L753:
	;
	goto L754
L754:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2410 = m.ExcPending
	if v2410 != 0 {
		goto L1
	} else {
		goto L755
	}
L755:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L1
	} else {
		goto L756
	}
L756:
	;
	v2414 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2415 = F_jspOperationName(m, v2414)
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L1
	} else {
		goto L757
	}
L757:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+432)) = v2415
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_13), v25+int32(432))
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		goto L1
	} else {
		goto L758
	}
L758:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2452), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v2427 = m.ExcPending
	if v2427 != 0 {
		goto L1
	} else {
		goto L759
	}
L759:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L760:
	;
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v2432 - int32(37) {
	case 0:
		goto L767
	default:
		goto L766
	case 8:
		v2506 = v2393
		goto L765
	}
L761:
	;
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v2711 - int32(37) {
	case 0:
		v3299 = v2707
		goto L834
	default:
		goto L842
	case 8:
		goto L847
	case 13:
		goto L846
	case 14:
		goto L845
	case 15:
		goto L844
	case 16:
		goto L843
	}
L762:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2671 = m.ExcPending
	if v2671 != 0 {
		goto L1
	} else {
		goto L826
	}
L763:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2653 = m.ExcPending
	if v2653 != 0 {
		goto L1
	} else {
		goto L822
	}
L764:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2640 = m.ExcPending
	if v2640 != 0 {
		goto L1
	} else {
		goto L819
	}
L765:
	;
	v2508 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	v2510 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	v2514 = int32(0)
	goto L787
L766:
	;
	v2477 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2477 == int32(0) {
		v2506 = v2393
		goto L765
	} else {
		goto L780
	}
L767:
	;
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2435 == int32(0) {
		v2506 = v2393
		goto L765
	} else {
		goto L768
	}
L768:
	;
	v2439 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+872)) = v2439
	v2442 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+864)) = v2442
	v2445 = v25 + int32(1032)
	F_jspGetArg(m, l1, v2445)
	mBase = m.M
	v2447 = m.ExcPending
	if v2447 != 0 {
		goto L1
	} else {
		goto L769
	}
L769:
	;
	v2448 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1032))
	if v2448 != int32(1) {
		goto L764
	} else {
		goto L770
	}
L770:
	;
	v2452 = v25 + int32(1000)
	if v2452 != 0 {
		goto L772
	} else {
		goto L773
	}
L771:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1000))
	v2457 = F_cstring_to_text_with_len(m, v2455, v2456)
	mBase = m.M
	v2458 = m.ExcPending
	if v2458 != 0 {
		goto L1
	} else {
		goto L775
	}
L772:
	;
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v2445)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2452))) = v2453
	goto L774
L773:
	;
	goto L774
L774:
	;
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(v2445)+12))
	goto L771
L775:
	;
	v2468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2468 != 0 {
		goto L776
	} else {
		goto L777
	}
L776:
	;
	v2469 = int32(0)
	goto L778
L777:
	;
	v2469 = v25 + int32(864)
	goto L778
L778:
	;
	v2470 = F_parse_datetime(m, v2430, v2457, v25+int32(1068), v25+int32(1064), v25+int32(1060), v2469)
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L1
	} else {
		goto L779
	}
L779:
	;
	v2472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+868)))
	v2691 = v2472 << (uint(int32(1)) % 32) & int32(2)
	v2697 = v2393
	v2707 = v2470
	goto L761
L780:
	;
	v2481 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+872)) = v2481
	v2484 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+864)) = v2484
	v2487 = v25 + int32(1032)
	F_jspGetArg(m, l1, v2487)
	mBase = m.M
	v2489 = m.ExcPending
	if v2489 != 0 {
		goto L1
	} else {
		goto L781
	}
L781:
	;
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1032))
	if v2490 != int32(2) {
		goto L763
	} else {
		goto L782
	}
L782:
	;
	v2493 = *(*int32)(unsafe.Add(mBase, uint32(v2487)+12))
	v2496 = F_numeric_int4_safe(m, v2493, v25+int32(864))
	mBase = m.M
	v2497 = m.ExcPending
	if v2497 != 0 {
		goto L1
	} else {
		goto L783
	}
L783:
	;
	v2498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+868)))
	if v2498 != int32(1) {
		v2506 = v2496
		goto L765
	} else {
		goto L784
	}
L784:
	;
	v2501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2501 == int32(1) {
		goto L762
	} else {
		goto L785
	}
L785:
	;
	v3997 = int32(2)
	goto L8
L786:
	;
	v2573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	v2574 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v2574 == int32(37) {
		goto L798
	} else {
		goto L799
	}
L787:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+872)) = v2508
	*(*int64)(unsafe.Add(mBase, uint32(v25)+864)) = v2510
	v2537 = v2514 << (uint(int32(2)) % 32)
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(v2537)+uint32(_c_F_executeItemOptUnwrapTarget[3])))
	if v2538 == int32(0) {
		goto L789
	} else {
		goto L790
	}
L788:
	;
	v2691 = int32(0)
	v2697 = v2506
	v2707 = v2565
	goto L761
L789:
	;
	v2541 = int32(_a_F_executeItemOptUnwrapTarget_47)
	v2542 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[4]))
	v2545 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[4])) = v2545
	v2549 = *(*int32)(unsafe.Add(mBase, uint32(v2537)+uint32(_c_F_executeItemOptUnwrapTarget[6])))
	v2550 = F_cstring_to_text(m, v2549)
	mBase = m.M
	v2551 = m.ExcPending
	if v2551 != 0 {
		goto L1
	} else {
		goto L792
	}
L790:
	;
	v2555 = v2538
	goto L791
L791:
	;
	v2565 = F_parse_datetime(m, v2430, v2555, v25+int32(1068), v25+int32(1064), v25+int32(1060), v25+int32(864))
	mBase = m.M
	v2566 = m.ExcPending
	if v2566 != 0 {
		goto L1
	} else {
		goto L793
	}
L792:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2537)+uint32(_c_F_executeItemOptUnwrapTarget[3]))) = v2550
	*(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[4])) = v2542
	v2555 = v2550
	goto L791
L793:
	;
	v2567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+868)))
	if v2567 != 0 {
		goto L794
	} else {
		goto L795
	}
L794:
	;
	v2569 = v2514 + int32(1)
	if v2569 == int32(13) {
		goto L786
	} else {
		goto L797
	}
L795:
	;
	goto L796
L796:
	;
	goto L788
L797:
	;
	v2514 = v2569
	goto L787
L798:
	;
	if v2573&int32(1) == int32(0) {
		goto L801
	} else {
		goto L802
	}
L799:
	;
	goto L800
L800:
	;
	if v2573&int32(1) == int32(0) {
		goto L810
	} else {
		goto L811
	}
L801:
	;
	v3997 = int32(2)
	goto L8
L802:
	;
	goto L803
L803:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2585 = m.ExcPending
	if v2585 != 0 {
		goto L1
	} else {
		goto L804
	}
L804:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2588 = m.ExcPending
	if v2588 != 0 {
		goto L1
	} else {
		goto L805
	}
L805:
	;
	v2589 = F_text_to_cstring(m, v2430)
	mBase = m.M
	v2590 = m.ExcPending
	if v2590 != 0 {
		goto L1
	} else {
		goto L806
	}
L806:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+164)) = v2589
	*(*int32)(unsafe.Add(mBase, uint32(v25)+160)) = int32(_a_F_executeItemOptUnwrapTarget_48)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_49), v25+int32(160))
	mBase = m.M
	v2598 = m.ExcPending
	if v2598 != 0 {
		goto L1
	} else {
		goto L807
	}
L807:
	;
	F_errhint(m, int32(_a_F_executeItemOptUnwrapTarget_50), int32(0))
	mBase = m.M
	v2602 = m.ExcPending
	if v2602 != 0 {
		goto L1
	} else {
		goto L808
	}
L808:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2580), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v2607 = m.ExcPending
	if v2607 != 0 {
		goto L1
	} else {
		goto L809
	}
L809:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L810:
	;
	v3997 = int32(2)
	goto L8
L811:
	;
	goto L812
L812:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2616 = m.ExcPending
	if v2616 != 0 {
		goto L1
	} else {
		goto L813
	}
L813:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2619 = m.ExcPending
	if v2619 != 0 {
		goto L1
	} else {
		goto L814
	}
L814:
	;
	v2620 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2621 = F_jspOperationName(m, v2620)
	mBase = m.M
	v2622 = m.ExcPending
	if v2622 != 0 {
		goto L1
	} else {
		goto L815
	}
L815:
	;
	v2623 = F_text_to_cstring(m, v2430)
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L1
	} else {
		goto L816
	}
L816:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+180)) = v2623
	*(*int32)(unsafe.Add(mBase, uint32(v25)+176)) = v2621
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_49), v25+int32(176))
	mBase = m.M
	v2631 = m.ExcPending
	if v2631 != 0 {
		goto L1
	} else {
		goto L817
	}
L817:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2585), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v2636 = m.ExcPending
	if v2636 != 0 {
		goto L1
	} else {
		goto L818
	}
L818:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L819:
	;
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_51), int32(0))
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L1
	} else {
		goto L820
	}
L820:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2478), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v2649 = m.ExcPending
	if v2649 != 0 {
		goto L1
	} else {
		goto L821
	}
L821:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L822:
	;
	v2654 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2655 = F_jspOperationName(m, v2654)
	mBase = m.M
	v2656 = m.ExcPending
	if v2656 != 0 {
		goto L1
	} else {
		goto L823
	}
L823:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+416)) = v2655
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_52), v25+int32(416))
	mBase = m.M
	v2662 = m.ExcPending
	if v2662 != 0 {
		goto L1
	} else {
		goto L824
	}
L824:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2537), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v2667 = m.ExcPending
	if v2667 != 0 {
		goto L1
	} else {
		goto L825
	}
L825:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L826:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2674 = m.ExcPending
	if v2674 != 0 {
		goto L1
	} else {
		goto L827
	}
L827:
	;
	v2675 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2676 = F_jspOperationName(m, v2675)
	mBase = m.M
	v2677 = m.ExcPending
	if v2677 != 0 {
		goto L1
	} else {
		goto L828
	}
L828:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+400)) = v2676
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_53), v25+int32(400))
	mBase = m.M
	v2683 = m.ExcPending
	if v2683 != 0 {
		goto L1
	} else {
		goto L829
	}
L829:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2545), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v2688 = m.ExcPending
	if v2688 != 0 {
		goto L1
	} else {
		goto L830
	}
L830:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L831:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3380 = m.ExcPending
	if v3380 != 0 {
		goto L1
	} else {
		goto L1035
	}
L832:
	;
	v3350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3350 != int32(1) {
		goto L1027
	} else {
		goto L1028
	}
L833:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3332 = m.ExcPending
	if v3332 != 0 {
		goto L1
	} else {
		goto L1022
	}
L834:
	;
	F_pfree(m, v2430)
	mBase = m.M
	v3301 = m.ExcPending
	if v3301 != 0 {
		goto L1
	} else {
		goto L1014
	}
L835:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1068)) = int32(1184)
	v3299 = v3294
	goto L834
L836:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1064)) = v3275
	v3292 = *(*int64)(unsafe.Add(mBase, uint32(v25)+1016))
	v3294 = v3292
	goto L835
L837:
	;
	v3997 = int32(2)
	goto L8
L838:
	;
	if v2697 == int32(-1) {
		v3294 = v3265
		goto L835
	} else {
		goto L1009
	}
L839:
	;
	v3262 = F_DirectFunctionCall1Coll(m, v3260, int32(0), v2707)
	mBase = m.M
	v3263 = m.ExcPending
	if v3263 != 0 {
		goto L1
	} else {
		goto L1008
	}
L840:
	;
	v3247 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[7]))
	v3249 = m.G0
	v3250 = int32(16)
	v3251 = v3249 - v3250
	m.G0 = v3251
	v3255 = F_DetermineTimeZoneOffsetInternal(m, v25+int32(864), v3247, v3251+int32(8))
	mBase = m.M
	m.G0 = v3251 + v3250
	goto L1007
L841:
	;
	v3228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v3228, int32(_a_F_executeItemOptUnwrapTarget_54), int32(_a_F_executeItemOptUnwrapTarget_55))
	mBase = m.M
	v3232 = m.ExcPending
	if v3232 != 0 {
		goto L1
	} else {
		goto L1004
	}
L842:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3215 = m.ExcPending
	if v3215 != 0 {
		goto L1
	} else {
		goto L1001
	}
L843:
	;
	v3075 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1068))
	if v3075 <= int32(1183) {
		goto L980
	} else {
		goto L981
	}
L844:
	;
	v2955 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1068))
	if v2955 <= int32(1183) {
		goto L941
	} else {
		goto L942
	}
L845:
	;
	v2878 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1068))
	if v2878 <= int32(1183) {
		goto L912
	} else {
		goto L913
	}
L846:
	;
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1068))
	if v2766 <= int32(1183) {
		goto L876
	} else {
		goto L877
	}
L847:
	;
	v2714 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1068))
	if v2714 <= int32(1183) {
		goto L853
	} else {
		goto L854
	}
L848:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1068)) = int32(1082)
	v3299 = v2763
	goto L834
L849:
	;
	v2760 = F_DirectFunctionCall1Coll(m, v2758, int32(0), v2707)
	mBase = m.M
	v2761 = m.ExcPending
	if v2761 != 0 {
		goto L1
	} else {
		goto L868
	}
L850:
	;
	if v2714 != int32(1114) {
		goto L831
	} else {
		goto L867
	}
L851:
	;
	v2749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2749, int32(_a_F_executeItemOptUnwrapTarget_55), int32(_a_F_executeItemOptUnwrapTarget_56))
	mBase = m.M
	v2753 = m.ExcPending
	if v2753 != 0 {
		goto L1
	} else {
		goto L866
	}
L852:
	;
	v2723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2723 != int32(1) {
		goto L858
	} else {
		goto L859
	}
L853:
	;
	switch v2714 - int32(1082) {
	case 0:
		v2763 = v2707
		goto L848
	case 1:
		goto L852
	default:
		goto L850
	}
L854:
	;
	goto L855
L855:
	;
	if v2714 == int32(1184) {
		goto L851
	} else {
		goto L856
	}
L856:
	;
	if v2714 != int32(1266) {
		goto L831
	} else {
		goto L857
	}
L857:
	;
	goto L852
L858:
	;
	v3997 = int32(2)
	goto L8
L859:
	;
	goto L860
L860:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2730 = m.ExcPending
	if v2730 != 0 {
		goto L1
	} else {
		goto L861
	}
L861:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2733 = m.ExcPending
	if v2733 != 0 {
		goto L1
	} else {
		goto L862
	}
L862:
	;
	v2734 = F_text_to_cstring(m, v2430)
	mBase = m.M
	v2735 = m.ExcPending
	if v2735 != 0 {
		goto L1
	} else {
		goto L863
	}
L863:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+228)) = v2734
	*(*int32)(unsafe.Add(mBase, uint32(v25)+224)) = int32(_a_F_executeItemOptUnwrapTarget_56)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_49), v25+int32(224))
	mBase = m.M
	v2743 = m.ExcPending
	if v2743 != 0 {
		goto L1
	} else {
		goto L864
	}
L864:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2612), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v2748 = m.ExcPending
	if v2748 != 0 {
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
	v2758 = int32(1558)
	goto L849
L867:
	;
	v2758 = int32(1559)
	goto L849
L868:
	;
	v2763 = v2760
	goto L848
L869:
	;
	if v2697 != int32(-1) {
		goto L895
	} else {
		goto L896
	}
L870:
	;
	v2834 = F_DirectFunctionCall1Coll(m, v2832, int32(0), v2707)
	mBase = m.M
	v2835 = m.ExcPending
	if v2835 != 0 {
		goto L1
	} else {
		goto L894
	}
L871:
	;
	v2827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2827, v2826, int32(_a_F_executeItemOptUnwrapTarget_57))
	mBase = m.M
	v2830 = m.ExcPending
	if v2830 != 0 {
		goto L1
	} else {
		goto L893
	}
L872:
	;
	v2825 = int32(1562)
	v2826 = int32(_a_F_executeItemOptUnwrapTarget_55)
	goto L871
L873:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2810 = m.ExcPending
	if v2810 != 0 {
		goto L1
	} else {
		goto L890
	}
L874:
	;
	if v2766 == int32(1114) {
		v2832 = int32(1561)
		goto L870
	} else {
		goto L889
	}
L875:
	;
	v2777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2777 != int32(1) {
		goto L881
	} else {
		goto L882
	}
L876:
	;
	switch v2766 - int32(1082) {
	case 0:
		goto L875
	case 1:
		v2838 = v2707
		goto L869
	default:
		goto L874
	}
L877:
	;
	goto L878
L878:
	;
	if v2766 == int32(1184) {
		goto L872
	} else {
		goto L879
	}
L879:
	;
	if v2766 != int32(1266) {
		goto L873
	} else {
		goto L880
	}
L880:
	;
	v2825 = int32(1560)
	v2826 = int32(_a_F_executeItemOptUnwrapTarget_58)
	goto L871
L881:
	;
	v3997 = int32(2)
	goto L8
L882:
	;
	goto L883
L883:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2784 = m.ExcPending
	if v2784 != 0 {
		goto L1
	} else {
		goto L884
	}
L884:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2787 = m.ExcPending
	if v2787 != 0 {
		goto L1
	} else {
		goto L885
	}
L885:
	;
	v2788 = F_text_to_cstring(m, v2430)
	mBase = m.M
	v2789 = m.ExcPending
	if v2789 != 0 {
		goto L1
	} else {
		goto L886
	}
L886:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+260)) = v2788
	*(*int32)(unsafe.Add(mBase, uint32(v25)+256)) = int32(_a_F_executeItemOptUnwrapTarget_57)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_49), v25+int32(256))
	mBase = m.M
	v2797 = m.ExcPending
	if v2797 != 0 {
		goto L1
	} else {
		goto L887
	}
L887:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2640), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v2802 = m.ExcPending
	if v2802 != 0 {
		goto L1
	} else {
		goto L888
	}
L888:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L889:
	;
	goto L873
L890:
	;
	v2811 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1068))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+240)) = v2811
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_59), v25+int32(240))
	mBase = m.M
	v2817 = m.ExcPending
	if v2817 != 0 {
		goto L1
	} else {
		goto L891
	}
L891:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2661), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v2822 = m.ExcPending
	if v2822 != 0 {
		goto L1
	} else {
		goto L892
	}
L892:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L893:
	;
	v2832 = v2825
	goto L870
L894:
	;
	v2838 = v2834
	goto L869
L895:
	;
	v2842 = F_anytime_typmod_check(m, int32(0), v2697)
	mBase = m.M
	v2843 = m.ExcPending
	if v2843 != 0 {
		goto L1
	} else {
		goto L898
	}
L896:
	;
	v2875 = v2838
	goto L897
L897:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1068)) = int32(1083)
	v3299 = v2875
	goto L834
L898:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+864)) = v2838
	v2846 = v25 + int32(864)
	if base.Ui32(v2842) <= base.Ui32(int32(6)) {
		goto L900
	} else {
		goto L901
	}
L899:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1064)) = v2842
	v2873 = *(*int64)(unsafe.Add(mBase, uint32(v25)+864))
	v2875 = v2873
	goto L897
L900:
	;
	v2853 = v2842 << (uint(int32(3)) % 32)
	v2854 = *(*int64)(unsafe.Add(mBase, uint32(v2853)+uint32(_c_F_executeItemOptUnwrapTarget[8])))
	v2855 = *(*int64)(unsafe.Add(mBase, uint32(v2853)+uint32(_c_F_executeItemOptUnwrapTarget[9])))
	v2856 = *(*int64)(unsafe.Add(mBase, uint32(v2846)))
	if int64(0) <= v2856 {
		goto L904
	} else {
		goto L905
	}
L901:
	;
	goto L902
L902:
	;
	goto L899
L903:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2846))) = v2866
	goto L902
L904:
	;
	v2859 = v2855 + v2856
	v2860 = base.I64_rem_s(v2859, v2854)
	v2866 = v2859 - v2860
	goto L903
L905:
	;
	goto L906
L906:
	;
	v2862 = v2855 - v2856
	v2863 = base.I64_rem_s(v2862, v2854)
	v2866 = v2863 - v2862
	goto L903
L907:
	;
	if v2697 != int32(-1) {
		goto L923
	} else {
		goto L924
	}
L908:
	;
	v2914 = F_DirectFunctionCall1Coll(m, v2912, int32(0), v2707)
	mBase = m.M
	v2915 = m.ExcPending
	if v2915 != 0 {
		goto L1
	} else {
		goto L922
	}
L909:
	;
	v2906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2906, int32(_a_F_executeItemOptUnwrapTarget_57), int32(_a_F_executeItemOptUnwrapTarget_58))
	mBase = m.M
	v2910 = m.ExcPending
	if v2910 != 0 {
		goto L1
	} else {
		goto L921
	}
L910:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L1
	} else {
		goto L918
	}
L911:
	;
	if v2878 == int32(1114) {
		goto L832
	} else {
		goto L917
	}
L912:
	;
	switch v2878 - int32(1082) {
	case 0:
		goto L832
	case 1:
		goto L909
	default:
		goto L911
	}
L913:
	;
	goto L914
L914:
	;
	if v2878 == int32(1184) {
		v2912 = int32(1563)
		goto L908
	} else {
		goto L915
	}
L915:
	;
	if v2878 != int32(1266) {
		goto L910
	} else {
		goto L916
	}
L916:
	;
	v2916 = v2707
	goto L907
L917:
	;
	goto L910
L918:
	;
	v2894 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1068))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+272)) = v2894
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_59), v25+int32(272))
	mBase = m.M
	v2900 = m.ExcPending
	if v2900 != 0 {
		goto L1
	} else {
		goto L919
	}
L919:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2708), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v2905 = m.ExcPending
	if v2905 != 0 {
		goto L1
	} else {
		goto L920
	}
L920:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L921:
	;
	v2912 = int32(1564)
	goto L908
L922:
	;
	v2916 = v2914
	goto L907
L923:
	;
	v2919 = base.I32_wrap_i64(v2916)
	v2921 = F_anytime_typmod_check(m, int32(1), v2697)
	mBase = m.M
	v2922 = m.ExcPending
	if v2922 != 0 {
		goto L1
	} else {
		goto L926
	}
L924:
	;
	v2952 = v2916
	goto L925
L925:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1068)) = int32(1266)
	v3299 = v2952
	goto L834
L926:
	;
	if base.Ui32(v2921) <= base.Ui32(int32(6)) {
		goto L928
	} else {
		goto L929
	}
L927:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1064)) = v2921
	v2952 = v2916 & int64(4294967295)
	goto L925
L928:
	;
	v2929 = v2921 << (uint(int32(3)) % 32)
	v2930 = *(*int64)(unsafe.Add(mBase, uint32(v2929)+uint32(_c_F_executeItemOptUnwrapTarget[8])))
	v2931 = *(*int64)(unsafe.Add(mBase, uint32(v2929)+uint32(_c_F_executeItemOptUnwrapTarget[9])))
	v2932 = *(*int64)(unsafe.Add(mBase, uint32(v2919)))
	if int64(0) <= v2932 {
		goto L932
	} else {
		goto L933
	}
L929:
	;
	goto L930
L930:
	;
	goto L927
L931:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2919))) = v2942
	goto L930
L932:
	;
	v2935 = v2931 + v2932
	v2936 = base.I64_rem_s(v2935, v2930)
	v2942 = v2935 - v2936
	goto L931
L933:
	;
	goto L934
L934:
	;
	v2938 = v2931 - v2932
	v2939 = base.I64_rem_s(v2938, v2930)
	v2942 = v2939 - v2938
	goto L931
L935:
	;
	if v2697 != int32(-1) {
		goto L960
	} else {
		goto L961
	}
L936:
	;
	v3019 = F_DirectFunctionCall1Coll(m, v3017, int32(0), v2707)
	mBase = m.M
	v3020 = m.ExcPending
	if v3020 != 0 {
		goto L1
	} else {
		goto L959
	}
L937:
	;
	v3011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v3011, int32(_a_F_executeItemOptUnwrapTarget_55), int32(_a_F_executeItemOptUnwrapTarget_54))
	mBase = m.M
	v3015 = m.ExcPending
	if v3015 != 0 {
		goto L1
	} else {
		goto L958
	}
L938:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2998 = m.ExcPending
	if v2998 != 0 {
		goto L1
	} else {
		goto L955
	}
L939:
	;
	if v2955 == int32(1114) {
		v3022 = v2707
		goto L935
	} else {
		goto L954
	}
L940:
	;
	v2966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2966 != int32(1) {
		goto L946
	} else {
		goto L947
	}
L941:
	;
	switch v2955 - int32(1082) {
	case 0:
		v3017 = int32(1565)
		goto L936
	case 1:
		goto L940
	default:
		goto L939
	}
L942:
	;
	goto L943
L943:
	;
	if v2955 == int32(1184) {
		goto L937
	} else {
		goto L944
	}
L944:
	;
	if v2955 != int32(1266) {
		goto L938
	} else {
		goto L945
	}
L945:
	;
	goto L940
L946:
	;
	v3997 = int32(2)
	goto L8
L947:
	;
	goto L948
L948:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2973 = m.ExcPending
	if v2973 != 0 {
		goto L1
	} else {
		goto L949
	}
L949:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2976 = m.ExcPending
	if v2976 != 0 {
		goto L1
	} else {
		goto L950
	}
L950:
	;
	v2977 = F_text_to_cstring(m, v2430)
	mBase = m.M
	v2978 = m.ExcPending
	if v2978 != 0 {
		goto L1
	} else {
		goto L951
	}
L951:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+340)) = v2977
	*(*int32)(unsafe.Add(mBase, uint32(v25)+336)) = int32(_a_F_executeItemOptUnwrapTarget_54)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_49), v25+int32(336))
	mBase = m.M
	v2986 = m.ExcPending
	if v2986 != 0 {
		goto L1
	} else {
		goto L952
	}
L952:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2744), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v2991 = m.ExcPending
	if v2991 != 0 {
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
	goto L938
L955:
	;
	v2999 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1068))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+304)) = v2999
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_59), v25+int32(304))
	mBase = m.M
	v3005 = m.ExcPending
	if v3005 != 0 {
		goto L1
	} else {
		goto L956
	}
L956:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2755), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v3010 = m.ExcPending
	if v3010 != 0 {
		goto L1
	} else {
		goto L957
	}
L957:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L958:
	;
	v3017 = int32(1566)
	goto L936
L959:
	;
	v3022 = v3019
	goto L935
L960:
	;
	v3026 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+872)) = v3026
	v3029 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+864)) = v3029
	v3032 = F_anytimestamp_typmod_check(m, int32(0), v2697)
	mBase = m.M
	v3033 = m.ExcPending
	if v3033 != 0 {
		goto L1
	} else {
		goto L963
	}
L961:
	;
	v3072 = v3022
	goto L962
L962:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1068)) = int32(1114)
	v3299 = v3072
	goto L834
L963:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1000)) = v3022
	v3039 = F_AdjustTimestampForTypmod(m, v25+int32(1000), v3032, v25+int32(864))
	mBase = m.M
	v3040 = m.ExcPending
	if v3040 != 0 {
		goto L1
	} else {
		goto L964
	}
L964:
	;
	v3041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+868)))
	if v3041 == int32(1) {
		goto L965
	} else {
		goto L966
	}
L965:
	;
	v3044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3044 == int32(0) {
		goto L968
	} else {
		goto L969
	}
L966:
	;
	goto L967
L967:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1064)) = v3032
	v3070 = *(*int64)(unsafe.Add(mBase, uint32(v25)+1000))
	v3072 = v3070
	goto L962
L968:
	;
	v3997 = int32(2)
	goto L8
L969:
	;
	goto L970
L970:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3051 = m.ExcPending
	if v3051 != 0 {
		goto L1
	} else {
		goto L971
	}
L971:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v3054 = m.ExcPending
	if v3054 != 0 {
		goto L1
	} else {
		goto L972
	}
L972:
	;
	v3055 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3056 = F_jspOperationName(m, v3055)
	mBase = m.M
	v3057 = m.ExcPending
	if v3057 != 0 {
		goto L1
	} else {
		goto L973
	}
L973:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+320)) = v3056
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_60), v25+int32(320))
	mBase = m.M
	v3063 = m.ExcPending
	if v3063 != 0 {
		goto L1
	} else {
		goto L974
	}
L974:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2774), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v3068 = m.ExcPending
	if v3068 != 0 {
		goto L1
	} else {
		goto L975
	}
L975:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L976:
	;
	v3127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v3127, int32(_a_F_executeItemOptUnwrapTarget_56), int32(_a_F_executeItemOptUnwrapTarget_55))
	mBase = m.M
	v3131 = m.ExcPending
	if v3131 != 0 {
		goto L1
	} else {
		goto L995
	}
L977:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3114 = m.ExcPending
	if v3114 != 0 {
		goto L1
	} else {
		goto L992
	}
L978:
	;
	if v3075 == int32(1114) {
		goto L841
	} else {
		goto L991
	}
L979:
	;
	v3084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3084 != int32(1) {
		goto L837
	} else {
		goto L985
	}
L980:
	;
	switch v3075 - int32(1082) {
	case 0:
		goto L976
	case 1:
		goto L979
	default:
		goto L978
	}
L981:
	;
	goto L982
L982:
	;
	if v3075 == int32(1184) {
		v3265 = v2707
		goto L838
	} else {
		goto L983
	}
L983:
	;
	if v3075 != int32(1266) {
		goto L977
	} else {
		goto L984
	}
L984:
	;
	goto L979
L985:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3090 = m.ExcPending
	if v3090 != 0 {
		goto L1
	} else {
		goto L986
	}
L986:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v3093 = m.ExcPending
	if v3093 != 0 {
		goto L1
	} else {
		goto L987
	}
L987:
	;
	v3094 = F_text_to_cstring(m, v2430)
	mBase = m.M
	v3095 = m.ExcPending
	if v3095 != 0 {
		goto L1
	} else {
		goto L988
	}
L988:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+388)) = v3094
	*(*int32)(unsafe.Add(mBase, uint32(v25)+384)) = int32(_a_F_executeItemOptUnwrapTarget_61)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_49), v25+int32(384))
	mBase = m.M
	v3103 = m.ExcPending
	if v3103 != 0 {
		goto L1
	} else {
		goto L989
	}
L989:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2815), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v3108 = m.ExcPending
	if v3108 != 0 {
		goto L1
	} else {
		goto L990
	}
L990:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L991:
	;
	goto L977
L992:
	;
	v3115 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1068))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+352)) = v3115
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_59), v25+int32(352))
	mBase = m.M
	v3121 = m.ExcPending
	if v3121 != 0 {
		goto L1
	} else {
		goto L993
	}
L993:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2836), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v3126 = m.ExcPending
	if v3126 != 0 {
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
	v3132 = base.I32_wrap_i64(v2707)
	v3144 = v3132 + int32(_a_F_executeItemOptUnwrapTarget_62)
	v3145 = int32(_a_F_executeItemOptUnwrapTarget_63)
	v3146 = base.I32_div_u_s(v3144, v3145)
	v3147 = int32(3)
	v3153 = int32(2)
	v3158 = base.I32_div_u_s((v3146*int32(1073595727)+v3144)<<(uint(v3153)%32)|v3147, v3145)
	v3161 = v3132 + int32(_a_F_executeItemOptUnwrapTarget_64) + v3146*v3147 + v3158 + int32(_a_F_executeItemOptUnwrapTarget_65)
	v3162 = int32(1461)
	v3163 = base.I32_div_u_s(v3161, v3162)
	v3166 = v3163*int32(-1461) + v3161
	v3168 = v3166 << (uint(v3153) % 32)
	if base.Ui32(v3162) <= base.Ui32(v3168) {
		goto L998
	} else {
		goto L999
	}
L996:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+872)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+864)) = int64(0)
	v3243 = int32(1567)
	goto L840
L997:
	;
	v3181 = base.I32_div_u_s(v3168, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v25+int32(884)))) = v3181 + v3163<<(uint(int32(2))%32) - int32(_a_F_executeItemOptUnwrapTarget_66)
	v3189 = v3179 + int32(123)
	v3193 = int32(base.Ui32(v3189*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v25+int32(876)))) = v3189 - int32(base.Ui32(v3193*int32(_a_F_executeItemOptUnwrapTarget_67))>>(uint(int32(8))%32))
	v3203 = base.I32_rem_u_s(v3193+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v25+int32(880)))) = v3203 + int32(1)
	goto L996
L998:
	;
	v3174 = base.I32_rem_u_s(v3166+int32(305), int32(365))
	v3179 = v3174
	goto L997
L999:
	;
	goto L1000
L1000:
	;
	v3178 = base.I32_rem_u_s(v3166+int32(306), int32(366))
	v3179 = v3178
	goto L997
L1001:
	;
	v3216 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+192)) = v3216
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_9), v25+int32(192))
	mBase = m.M
	v3222 = m.ExcPending
	if v3222 != 0 {
		goto L1
	} else {
		goto L1002
	}
L1002:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2866), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v3227 = m.ExcPending
	if v3227 != 0 {
		goto L1
	} else {
		goto L1003
	}
L1003:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1004:
	;
	v3233 = int32(1568)
	v3234 = int32(0)
	v3241 = F_timestamp2tm(m, v2707, v3234, v25+int32(864), v25+int32(1028), v3234, v3234)
	mBase = m.M
	v3242 = m.ExcPending
	if v3242 != 0 {
		goto L1
	} else {
		goto L1005
	}
L1005:
	;
	if v3241 != 0 {
		v3260 = v3233
		goto L839
	} else {
		goto L1006
	}
L1006:
	;
	v3243 = v3233
	goto L840
L1007:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1060)) = v3255
	v3260 = v3243
	goto L839
L1008:
	;
	v3265 = v3262
	goto L838
L1009:
	;
	v3269 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1008)) = v3269
	v3272 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1000)) = v3272
	v3275 = F_anytimestamp_typmod_check(m, int32(1), v2697)
	mBase = m.M
	v3276 = m.ExcPending
	if v3276 != 0 {
		goto L1
	} else {
		goto L1010
	}
L1010:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1016)) = v3265
	v3282 = F_AdjustTimestampForTypmod(m, v25+int32(1016), v3275, v25+int32(1000))
	mBase = m.M
	v3283 = m.ExcPending
	if v3283 != 0 {
		goto L1
	} else {
		goto L1011
	}
L1011:
	;
	v3284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1004)))
	if v3284 != int32(1) {
		goto L836
	} else {
		goto L1012
	}
L1012:
	;
	v3287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3287 != 0 {
		goto L833
	} else {
		goto L1013
	}
L1013:
	;
	goto L837
L1014:
	;
	v3302 = int32(2)
	if v2691 == v3302 {
		v3997 = v3302
		goto L8
	} else {
		goto L1015
	}
L1015:
	;
	v3307 = F_jspGetNext(m, l1, v25+int32(1032))
	mBase = m.M
	v3308 = m.ExcPending
	if v3308 != 0 {
		goto L1
	} else {
		goto L1016
	}
L1016:
	;
	if l3 == int32(0) {
		goto L1017
	} else {
		goto L1018
	}
L1017:
	;
	if v3307 == int32(0) {
		v3997 = v2691
		goto L8
	} else {
		goto L1020
	}
L1018:
	;
	goto L1019
L1019:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1080)) = v3299
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1072)) = int32(32)
	v3317 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1068))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1088)) = v3317
	v3319 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1064))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1092)) = v3319
	v3321 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1060))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1096)) = v3321
	v3327 = F_executeNextItem(m, l0, l1, v25+int32(1032), v25+int32(1072), l3)
	mBase = m.M
	v3328 = m.ExcPending
	if v3328 != 0 {
		goto L1
	} else {
		goto L1021
	}
L1020:
	;
	goto L1019
L1021:
	;
	v3997 = v3327
	goto L8
L1022:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v3335 = m.ExcPending
	if v3335 != 0 {
		goto L1
	} else {
		goto L1023
	}
L1023:
	;
	v3336 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3337 = F_jspOperationName(m, v3336)
	mBase = m.M
	v3338 = m.ExcPending
	if v3338 != 0 {
		goto L1
	} else {
		goto L1024
	}
L1024:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+368)) = v3337
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_60), v25+int32(368))
	mBase = m.M
	v3344 = m.ExcPending
	if v3344 != 0 {
		goto L1
	} else {
		goto L1025
	}
L1025:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2855), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v3349 = m.ExcPending
	if v3349 != 0 {
		goto L1
	} else {
		goto L1026
	}
L1026:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1027:
	;
	v3997 = int32(2)
	goto L8
L1028:
	;
	goto L1029
L1029:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3357 = m.ExcPending
	if v3357 != 0 {
		goto L1
	} else {
		goto L1030
	}
L1030:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v3360 = m.ExcPending
	if v3360 != 0 {
		goto L1
	} else {
		goto L1031
	}
L1031:
	;
	v3361 = F_text_to_cstring(m, v2430)
	mBase = m.M
	v3362 = m.ExcPending
	if v3362 != 0 {
		goto L1
	} else {
		goto L1032
	}
L1032:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+292)) = v3361
	*(*int32)(unsafe.Add(mBase, uint32(v25)+288)) = int32(_a_F_executeItemOptUnwrapTarget_68)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_49), v25+int32(288))
	mBase = m.M
	v3370 = m.ExcPending
	if v3370 != 0 {
		goto L1
	} else {
		goto L1033
	}
L1033:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2693), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v3375 = m.ExcPending
	if v3375 != 0 {
		goto L1
	} else {
		goto L1034
	}
L1034:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1035:
	;
	v3381 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1068))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+208)) = v3381
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_59), v25+int32(208))
	mBase = m.M
	v3387 = m.ExcPending
	if v3387 != 0 {
		goto L1
	} else {
		goto L1036
	}
L1036:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(2625), int32(_a_F_executeItemOptUnwrapTarget_46))
	mBase = m.M
	v3392 = m.ExcPending
	if v3392 != 0 {
		goto L1
	} else {
		goto L1037
	}
L1037:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1038:
	;
	v3402 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v3402 - int32(1) {
	case 0:
		goto L1048
	case 1:
		goto L1049
	default:
		goto L1047
	}
L1039:
	;
	v3395 = F_JsonbType(m, l2)
	mBase = m.M
	v3396 = m.ExcPending
	if v3396 != 0 {
		goto L1
	} else {
		goto L1040
	}
L1040:
	;
	if v3395 != int32(16) {
		goto L1038
	} else {
		goto L1041
	}
L1041:
	;
	v3400 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v3401 = m.ExcPending
	if v3401 != 0 {
		goto L1
	} else {
		goto L1042
	}
L1042:
	;
	v3997 = v3400
	goto L8
L1043:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3592 = m.ExcPending
	if v3592 != 0 {
		goto L1
	} else {
		goto L1092
	}
L1044:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3571 = m.ExcPending
	if v3571 != 0 {
		goto L1
	} else {
		goto L1087
	}
L1045:
	;
	v3566 = F_executeNextItem(m, l0, l1, int32(0), v3561, l3)
	mBase = m.M
	v3567 = m.ExcPending
	if v3567 != 0 {
		goto L1
	} else {
		goto L1086
	}
L1046:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+864)) = int32(2)
	v3553 = F_DirectFunctionCall1Coll(m, int32(1541), int32(0), base.I64_reinterpret_f64(v3479))
	mBase = m.M
	v3554 = m.ExcPending
	if v3554 != 0 {
		goto L1
	} else {
		goto L1084
	}
L1047:
	;
	v3524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3524 != int32(1) {
		v3997 = int32(2)
		goto L8
	} else {
		goto L1078
	}
L1048:
	;
	v3465 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v3466 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3467 = F_pnstrdup(m, v3465, v3466)
	mBase = m.M
	v3468 = m.ExcPending
	if v3468 != 0 {
		goto L1
	} else {
		goto L1064
	}
L1049:
	;
	v3407 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+8)))
	v3408 = F_DirectFunctionCall1Coll(m, int32(664), int32(0), v3407)
	mBase = m.M
	v3409 = m.ExcPending
	if v3409 != 0 {
		goto L1
	} else {
		goto L1050
	}
L1050:
	;
	v3411 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1080)) = v3411
	v3414 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1072)) = v3414
	v3416 = base.I32_wrap_i64(v3408)
	v3421 = F_float8in_internal(m, v3416, int32(0), int32(_a_F_executeItemOptUnwrapTarget_69), v3416, v25+int32(1072))
	mBase = m.M
	v3422 = m.ExcPending
	if v3422 != 0 {
		goto L1
	} else {
		goto L1051
	}
L1051:
	;
	v3423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1076)))
	if v3423 == int32(1) {
		goto L1053
	} else {
		goto L1054
	}
L1052:
	;
	v3997 = int32(2)
	goto L8
L1053:
	;
	v3426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3426 != int32(1) {
		goto L1052
	} else {
		goto L1056
	}
L1054:
	;
	goto L1055
L1055:
	;
	v3453 = base.F64_abs(v3421)
	if base.F64_ne(v3453, math.Float64frombits(uint64(0x7ff0000000000000)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v3453)) < base.Ui64(int64(9218868437227405313))) != 0 {
		v3561 = l2
		goto L1045
	} else {
		goto L1062
	}
L1056:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3432 = m.ExcPending
	if v3432 != 0 {
		goto L1
	} else {
		goto L1057
	}
L1057:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3435 = m.ExcPending
	if v3435 != 0 {
		goto L1
	} else {
		goto L1058
	}
L1058:
	;
	v3436 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3437 = F_jspOperationName(m, v3436)
	mBase = m.M
	v3438 = m.ExcPending
	if v3438 != 0 {
		goto L1
	} else {
		goto L1059
	}
L1059:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+104)) = int32(_a_F_executeItemOptUnwrapTarget_69)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+100)) = v3437
	*(*int32)(unsafe.Add(mBase, uint32(v25)+96)) = v3416
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_26), v25+int32(96))
	mBase = m.M
	v3447 = m.ExcPending
	if v3447 != 0 {
		goto L1
	} else {
		goto L1060
	}
L1060:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1196), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v3452 = m.ExcPending
	if v3452 != 0 {
		goto L1
	} else {
		goto L1061
	}
L1061:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1062:
	;
	v3460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3460 == int32(1) {
		goto L1044
	} else {
		goto L1063
	}
L1063:
	;
	goto L1052
L1064:
	;
	v3470 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1080)) = v3470
	v3473 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1072)) = v3473
	v3479 = F_float8in_internal(m, v3467, int32(0), int32(_a_F_executeItemOptUnwrapTarget_69), v3467, v25+int32(1072))
	mBase = m.M
	v3480 = m.ExcPending
	if v3480 != 0 {
		goto L1
	} else {
		goto L1065
	}
L1065:
	;
	v3481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1076)))
	if v3481 == int32(1) {
		goto L1067
	} else {
		goto L1068
	}
L1066:
	;
	v3997 = int32(2)
	goto L8
L1067:
	;
	v3484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3484 != int32(1) {
		goto L1066
	} else {
		goto L1070
	}
L1068:
	;
	goto L1069
L1069:
	;
	v3511 = base.F64_abs(v3479)
	if base.F64_ne(v3511, math.Float64frombits(uint64(0x7ff0000000000000)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v3511)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L1046
	} else {
		goto L1076
	}
L1070:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3490 = m.ExcPending
	if v3490 != 0 {
		goto L1
	} else {
		goto L1071
	}
L1071:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3493 = m.ExcPending
	if v3493 != 0 {
		goto L1
	} else {
		goto L1072
	}
L1072:
	;
	v3494 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3495 = F_jspOperationName(m, v3494)
	mBase = m.M
	v3496 = m.ExcPending
	if v3496 != 0 {
		goto L1
	} else {
		goto L1073
	}
L1073:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+136)) = int32(_a_F_executeItemOptUnwrapTarget_69)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+132)) = v3495
	*(*int32)(unsafe.Add(mBase, uint32(v25)+128)) = v3467
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_26), v25+int32(128))
	mBase = m.M
	v3505 = m.ExcPending
	if v3505 != 0 {
		goto L1
	} else {
		goto L1074
	}
L1074:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1222), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v3510 = m.ExcPending
	if v3510 != 0 {
		goto L1
	} else {
		goto L1075
	}
L1075:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1076:
	;
	v3518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3518 == int32(1) {
		goto L1043
	} else {
		goto L1077
	}
L1077:
	;
	goto L1066
L1078:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3530 = m.ExcPending
	if v3530 != 0 {
		goto L1
	} else {
		goto L1079
	}
L1079:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3533 = m.ExcPending
	if v3533 != 0 {
		goto L1
	} else {
		goto L1080
	}
L1080:
	;
	v3534 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3535 = F_jspOperationName(m, v3534)
	mBase = m.M
	v3536 = m.ExcPending
	if v3536 != 0 {
		goto L1
	} else {
		goto L1081
	}
L1081:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+80)) = v3535
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_27), v25+int32(80))
	mBase = m.M
	v3542 = m.ExcPending
	if v3542 != 0 {
		goto L1
	} else {
		goto L1082
	}
L1082:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1240), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v3547 = m.ExcPending
	if v3547 != 0 {
		goto L1
	} else {
		goto L1083
	}
L1083:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1084:
	;
	v3556 = F_pg_detoast_datum(m, base.I32_wrap_i64(v3553))
	mBase = m.M
	v3557 = m.ExcPending
	if v3557 != 0 {
		goto L1
	} else {
		goto L1085
	}
L1085:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+872)) = v3556
	v3561 = v25 + int32(864)
	goto L1045
L1086:
	;
	v3997 = v3566
	goto L8
L1087:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3574 = m.ExcPending
	if v3574 != 0 {
		goto L1
	} else {
		goto L1088
	}
L1088:
	;
	v3575 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3576 = F_jspOperationName(m, v3575)
	mBase = m.M
	v3577 = m.ExcPending
	if v3577 != 0 {
		goto L1
	} else {
		goto L1089
	}
L1089:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+112)) = v3576
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_31), v25+int32(112))
	mBase = m.M
	v3583 = m.ExcPending
	if v3583 != 0 {
		goto L1
	} else {
		goto L1090
	}
L1090:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1201), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v3588 = m.ExcPending
	if v3588 != 0 {
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
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3595 = m.ExcPending
	if v3595 != 0 {
		goto L1
	} else {
		goto L1093
	}
L1093:
	;
	v3596 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3597 = F_jspOperationName(m, v3596)
	mBase = m.M
	v3598 = m.ExcPending
	if v3598 != 0 {
		goto L1
	} else {
		goto L1094
	}
L1094:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+144)) = v3597
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_31), v25+int32(144))
	mBase = m.M
	v3604 = m.ExcPending
	if v3604 != 0 {
		goto L1
	} else {
		goto L1095
	}
L1095:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1227), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v3609 = m.ExcPending
	if v3609 != 0 {
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
	v3997 = v3611
	goto L8
L1098:
	;
	v3997 = v3614
	goto L8
L1099:
	;
	v3997 = v3617
	goto L8
L1100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+864)) = int32(2)
	v3665 = F_int64_to_numeric(m, v3662)
	mBase = m.M
	v3666 = m.ExcPending
	if v3666 != 0 {
		goto L1
	} else {
		goto L1114
	}
L1101:
	;
	v3633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v3633 != 0 {
		v3662 = int64(1)
		goto L1100
	} else {
		goto L1104
	}
L1102:
	;
	v3622 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v3623 = *(*int32)(unsafe.Add(mBase, uint32(v3622)))
	if v3623&int32(1342177280) != int32(1073741824) {
		goto L1101
	} else {
		goto L1103
	}
L1103:
	;
	v3662 = base.I64_extend_i32_u(v3623 & int32(268435455))
	goto L1100
L1104:
	;
	v3635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3635 != 0 {
		v3997 = int32(1)
		goto L8
	} else {
		goto L1105
	}
L1105:
	;
	v3636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3636 != int32(1) {
		goto L1106
	} else {
		goto L1107
	}
L1106:
	;
	v3997 = int32(2)
	goto L8
L1107:
	;
	goto L1108
L1108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3643 = m.ExcPending
	if v3643 != 0 {
		goto L1
	} else {
		goto L1109
	}
L1109:
	;
	F_errcode(m, int32(151781506))
	mBase = m.M
	v3646 = m.ExcPending
	if v3646 != 0 {
		goto L1
	} else {
		goto L1110
	}
L1110:
	;
	v3647 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3648 = F_jspOperationName(m, v3647)
	mBase = m.M
	v3649 = m.ExcPending
	if v3649 != 0 {
		goto L1
	} else {
		goto L1111
	}
L1111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = v3648
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_70), v25-int32(-64))
	mBase = m.M
	v3655 = m.ExcPending
	if v3655 != 0 {
		goto L1
	} else {
		goto L1112
	}
L1112:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1145), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v3660 = m.ExcPending
	if v3660 != 0 {
		goto L1
	} else {
		goto L1113
	}
L1113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+872)) = v3665
	v3671 = F_executeNextItem(m, l0, l1, int32(0), v25+int32(864), l3)
	mBase = m.M
	v3672 = m.ExcPending
	if v3672 != 0 {
		goto L1
	} else {
		goto L1115
	}
L1115:
	;
	v3997 = v3671
	goto L8
L1116:
	;
	v3677 = F_pstrdup(m, v3675)
	mBase = m.M
	v3678 = m.ExcPending
	if v3678 != 0 {
		goto L1
	} else {
		goto L1117
	}
L1117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+876)) = v3677
	v3680 = F_strlen(m, v3677)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v25)+872)) = v3680
	v3685 = F_executeNextItem(m, l0, l1, int32(0), v25+int32(864), l3)
	mBase = m.M
	v3686 = m.ExcPending
	if v3686 != 0 {
		goto L1
	} else {
		goto L1118
	}
L1118:
	;
	v3997 = v3685
	goto L8
L1119:
	;
	v3708 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v3709 = m.ExcPending
	if v3709 != 0 {
		goto L1
	} else {
		goto L1129
	}
L1120:
	;
	v3687 = F_JsonbType(m, l2)
	mBase = m.M
	v3688 = m.ExcPending
	if v3688 != 0 {
		goto L1
	} else {
		goto L1123
	}
L1121:
	;
	goto L1122
L1122:
	;
	v3692 = v25 + int32(1032)
	F_jspGetArg(m, l1, v3692)
	mBase = m.M
	v3694 = m.ExcPending
	if v3694 != 0 {
		goto L1
	} else {
		goto L1125
	}
L1123:
	;
	if v3687 == int32(16) {
		goto L1119
	} else {
		goto L1124
	}
L1124:
	;
	goto L1122
L1125:
	;
	v3695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
	v3698 = F_executeBoolItem(m, l0, v3692, l2, int32(0))
	mBase = m.M
	v3699 = m.ExcPending
	if v3699 != 0 {
		goto L1
	} else {
		goto L1126
	}
L1126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v3695
	v3701 = int32(1)
	if v3698 != v3701 {
		v3997 = v3701
		goto L8
	} else {
		goto L1127
	}
L1127:
	;
	v3705 = F_executeNextItem(m, l0, l1, int32(0), l2, l3)
	mBase = m.M
	v3706 = m.ExcPending
	if v3706 != 0 {
		goto L1
	} else {
		goto L1128
	}
L1128:
	;
	v3997 = v3705
	goto L8
L1129:
	;
	v3997 = v3708
	goto L8
L1130:
	;
	v3716 = *(*int32)(unsafe.Add(mBase, uint32(v3712)+12))
	v3717 = v3716
	goto L1132
L1131:
	;
	v3717 = int32(0)
	goto L1132
L1132:
	;
	v3718 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3718
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3717
	v3722 = F_executeNextItem(m, l0, l1, v3718, v3712, l3)
	mBase = m.M
	v3723 = m.ExcPending
	if v3723 != 0 {
		goto L1
	} else {
		goto L1133
	}
L1133:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v3710
	v3997 = v3722
	goto L8
L1134:
	;
	v3997 = v3727
	goto L8
L1135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3792 = m.ExcPending
	if v3792 != 0 {
		goto L1
	} else {
		goto L1167
	}
L1136:
	;
	if v3729 == int32(17) {
		goto L1137
	} else {
		goto L1138
	}
L1137:
	;
	v3733 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+864)) = v3733
	v3737 = v25 + int32(872)
	if v3737 != 0 {
		goto L1141
	} else {
		goto L1142
	}
L1138:
	;
	goto L1139
L1139:
	;
	if l4 == int32(0) {
		goto L1152
	} else {
		goto L1153
	}
L1140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+876)) = v3740
	v3742 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v3746 = F_findJsonbValueFromContainer(m, v3742, int32(536870912), v25+int32(864))
	mBase = m.M
	v3747 = m.ExcPending
	if v3747 != 0 {
		goto L1
	} else {
		goto L1144
	}
L1141:
	;
	v3738 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v3737))) = v3738
	goto L1143
L1142:
	;
	goto L1143
L1143:
	;
	v3740 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L1140
L1144:
	;
	if v3746 != 0 {
		goto L1145
	} else {
		goto L1146
	}
L1145:
	;
	v3749 = F_executeNextItem(m, l0, l1, int32(0), v3746, l3)
	mBase = m.M
	v3750 = m.ExcPending
	if v3750 != 0 {
		goto L1
	} else {
		goto L1148
	}
L1146:
	;
	goto L1147
L1147:
	;
	v3753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3753 != 0 {
		v3997 = v3733
		goto L8
	} else {
		goto L1150
	}
L1148:
	;
	F_pfree(m, v3746)
	mBase = m.M
	v3752 = m.ExcPending
	if v3752 != 0 {
		goto L1
	} else {
		goto L1149
	}
L1149:
	;
	v3997 = v3749
	goto L8
L1150:
	;
	v3754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3754 == int32(1) {
		goto L1135
	} else {
		goto L1151
	}
L1151:
	;
	v3997 = int32(2)
	goto L8
L1152:
	;
	v3767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3767 != 0 {
		goto L1157
	} else {
		goto L1158
	}
L1153:
	;
	v3760 = F_JsonbType(m, l2)
	mBase = m.M
	v3761 = m.ExcPending
	if v3761 != 0 {
		goto L1
	} else {
		goto L1154
	}
L1154:
	;
	if v3760 != int32(16) {
		goto L1152
	} else {
		goto L1155
	}
L1155:
	;
	v3765 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v3766 = m.ExcPending
	if v3766 != 0 {
		goto L1
	} else {
		goto L1156
	}
L1156:
	;
	v3997 = v3765
	goto L8
L1157:
	;
	v3997 = int32(1)
	goto L8
L1158:
	;
	goto L1159
L1159:
	;
	v3769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3769 != int32(1) {
		goto L1160
	} else {
		goto L1161
	}
L1160:
	;
	v3997 = int32(2)
	goto L8
L1161:
	;
	goto L1162
L1162:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3776 = m.ExcPending
	if v3776 != 0 {
		goto L1
	} else {
		goto L1163
	}
L1163:
	;
	F_errcode(m, int32(285999234))
	mBase = m.M
	v3779 = m.ExcPending
	if v3779 != 0 {
		goto L1
	} else {
		goto L1164
	}
L1164:
	;
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_71), int32(0))
	mBase = m.M
	v3783 = m.ExcPending
	if v3783 != 0 {
		goto L1
	} else {
		goto L1165
	}
L1165:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1087), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v3788 = m.ExcPending
	if v3788 != 0 {
		goto L1
	} else {
		goto L1166
	}
L1166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1167:
	;
	F_errcode(m, int32(285999234))
	mBase = m.M
	v3795 = m.ExcPending
	if v3795 != 0 {
		goto L1
	} else {
		goto L1168
	}
L1168:
	;
	v3796 = *(*int32)(unsafe.Add(mBase, uint32(v25)+876))
	v3797 = *(*int32)(unsafe.Add(mBase, uint32(v25)+872))
	v3798 = F_pnstrdup(m, v3796, v3797)
	mBase = m.M
	v3799 = m.ExcPending
	if v3799 != 0 {
		goto L1
	} else {
		goto L1169
	}
L1169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v3798
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_72), v25+int32(48))
	mBase = m.M
	v3805 = m.ExcPending
	if v3805 != 0 {
		goto L1
	} else {
		goto L1170
	}
L1170:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1077), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v3810 = m.ExcPending
	if v3810 != 0 {
		goto L1
	} else {
		goto L1171
	}
L1171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1172:
	;
	v3816 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v3816 != 0 {
		v3825 = int32(1)
		goto L1173
	} else {
		goto L1174
	}
L1173:
	;
	v3827 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v3827 != int32(18) {
		v3997 = v3825
		goto L8
	} else {
		goto L1177
	}
L1174:
	;
	v3817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	v3818 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v3818)
	v3820 = F_executeNextItem(m, l0, l1, v3812, l2, l3)
	mBase = m.M
	v3821 = m.ExcPending
	if v3821 != 0 {
		goto L1
	} else {
		goto L1175
	}
L1175:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v3817)
	if l3|v3820 != 0 {
		v3825 = v3820
		goto L1173
	} else {
		goto L1176
	}
L1176:
	;
	v3997 = int32(0)
	goto L8
L1177:
	;
	if v3813 != 0 {
		goto L1178
	} else {
		goto L1179
	}
L1178:
	;
	v3833 = v25 + int32(1032)
	goto L1180
L1179:
	;
	v3833 = int32(0)
	goto L1180
L1180:
	;
	v3834 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v3835 = int32(1)
	v3836 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3837 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v3840 = F_executeAnyItem(m, l0, v3833, v3834, l3, v3835, v3836, v3837, v3835, v3839)
	mBase = m.M
	v3841 = m.ExcPending
	if v3841 != 0 {
		goto L1
	} else {
		goto L1181
	}
L1181:
	;
	v3997 = v3840
	goto L8
L1182:
	;
	v3843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3843 != int32(1) {
		goto L1183
	} else {
		goto L1184
	}
L1183:
	;
	v3997 = int32(2)
	goto L8
L1184:
	;
	goto L1185
L1185:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3850 = m.ExcPending
	if v3850 != 0 {
		goto L1
	} else {
		goto L1186
	}
L1186:
	;
	F_errcode(m, int32(151781506))
	mBase = m.M
	v3853 = m.ExcPending
	if v3853 != 0 {
		goto L1
	} else {
		goto L1187
	}
L1187:
	;
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_73), int32(0))
	mBase = m.M
	v3857 = m.ExcPending
	if v3857 != 0 {
		goto L1
	} else {
		goto L1188
	}
L1188:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(1014), int32(_a_F_executeItemOptUnwrapTarget_5))
	mBase = m.M
	v3862 = m.ExcPending
	if v3862 != 0 {
		goto L1
	} else {
		goto L1189
	}
L1189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1190:
	;
	v3871 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3871 != 0 {
		goto L1195
	} else {
		goto L1196
	}
L1191:
	;
	v3868 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v3867))) = v3868
	goto L1193
L1192:
	;
	goto L1193
L1193:
	;
	v3870 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L1190
L1194:
	;
	v3902 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1016))
	if v3902 <= int32(0) {
		v3923 = v3864
		v3925 = v3865
		goto L9
	} else {
		goto L1205
	}
L1195:
	;
	v3872 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1000))
	v3877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3878 = m.T0[v3877].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v3871, v3870, v3872, v25+int32(864), v25+int32(1016))
	mBase = m.M
	v3879 = m.ExcPending
	if v3879 != 0 {
		goto L1
	} else {
		goto L1198
	}
L1196:
	;
	goto L1197
L1197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3884 = m.ExcPending
	if v3884 != 0 {
		goto L1
	} else {
		goto L1200
	}
L1198:
	;
	if v3878 != 0 {
		goto L1194
	} else {
		goto L1199
	}
L1199:
	;
	goto L1197
L1200:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v3887 = m.ExcPending
	if v3887 != 0 {
		goto L1
	} else {
		goto L1201
	}
L1201:
	;
	v3888 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1000))
	v3889 = F_pnstrdup(m, v3870, v3888)
	mBase = m.M
	v3890 = m.ExcPending
	if v3890 != 0 {
		goto L1
	} else {
		goto L1202
	}
L1202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v3889
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_74), v25+int32(16))
	mBase = m.M
	v3896 = m.ExcPending
	if v3896 != 0 {
		goto L1
	} else {
		goto L1203
	}
L1203:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(3428), int32(_a_F_executeItemOptUnwrapTarget_75))
	mBase = m.M
	v3901 = m.ExcPending
	if v3901 != 0 {
		goto L1
	} else {
		goto L1204
	}
L1204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1205:
	;
	v3905 = *(*int64)(unsafe.Add(mBase, uint32(v3878)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1096)) = v3905
	v3907 = *(*int64)(unsafe.Add(mBase, uint32(v3878)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1088)) = v3907
	v3909 = *(*int64)(unsafe.Add(mBase, uint32(v3878)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1080)) = v3909
	v3911 = *(*int64)(unsafe.Add(mBase, uint32(v3878)))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1072)) = v3911
	v3913 = *(*int32)(unsafe.Add(mBase, uint32(v25)+876))
	v3914 = *(*int32)(unsafe.Add(mBase, uint32(v25)+864))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3902
	if v3914 == int32(18) {
		goto L1206
	} else {
		goto L1207
	}
L1206:
	;
	v3919 = v3913
	goto L1208
L1207:
	;
	v3919 = int32(0)
	goto L1208
L1208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3919
	v3923 = v3864
	v3925 = v3865
	goto L9
L1209:
	;
	v3933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v3934 = F_executeItemOptUnwrapTarget(m, l0, v25+int32(1032), v25+int32(1072), l3, v3933)
	mBase = m.M
	v3935 = m.ExcPending
	if v3935 != 0 {
		goto L1
	} else {
		goto L1212
	}
L1210:
	;
	goto L1211
L1211:
	;
	if l3 != 0 {
		goto L1213
	} else {
		goto L1214
	}
L1212:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3923))) = v3925
	v3997 = v3934
	goto L8
L1213:
	;
	v3939 = v25 + int32(1072)
	v3940 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v3941 = *(*int32)(unsafe.Add(mBase, uint32(v3940)))
	v3942 = *(*int32)(unsafe.Add(mBase, uint32(v3940)+4))
	if v3941 < v3942 {
		goto L1217
	} else {
		goto L1218
	}
L1214:
	;
	goto L1215
L1215:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3923))) = v3925
	v3997 = int32(0)
	goto L8
L1216:
	;
	goto L1215
L1217:
	;
	v3946 = v3940 + v3941<<(uint(int32(5))%32)
	v3947 = *(*int64)(unsafe.Add(mBase, uint32(v3939)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3946)+40)) = v3947
	v3949 = *(*int64)(unsafe.Add(mBase, uint32(v3939)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3946)+32)) = v3949
	v3951 = *(*int64)(unsafe.Add(mBase, uint32(v3939)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3946)+24)) = v3951
	v3953 = *(*int64)(unsafe.Add(mBase, uint32(v3939)))
	*(*int64)(unsafe.Add(mBase, uint32(v3946)+16)) = v3953
	v3955 = *(*int32)(unsafe.Add(mBase, uint32(v3940)))
	*(*int32)(unsafe.Add(mBase, uint32(v3940))) = v3955 + int32(1)
	goto L1216
L1218:
	;
	goto L1219
L1219:
	;
	v3959 = int32(16)
	v3961 = v3942 << (uint(int32(1)) % 32)
	if v3961 <= v3959 {
		goto L1220
	} else {
		goto L1221
	}
L1220:
	;
	v3964 = v3959
	goto L1222
L1221:
	;
	v3964 = v3961
	goto L1222
L1222:
	;
	v3969 = F_palloc(m, v3964<<(uint(int32(5))%32)|int32(16))
	mBase = m.M
	v3970 = m.ExcPending
	if v3970 != 0 {
		goto L1
	} else {
		goto L1223
	}
L1223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3969)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3969)+4)) = v3964
	*(*int32)(unsafe.Add(mBase, uint32(v3969))) = int32(1)
	v3976 = *(*int64)(unsafe.Add(mBase, uint32(v3939)))
	*(*int64)(unsafe.Add(mBase, uint32(v3969)+16)) = v3976
	v3978 = *(*int64)(unsafe.Add(mBase, uint32(v3939)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3969)+24)) = v3978
	v3980 = *(*int64)(unsafe.Add(mBase, uint32(v3939)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3969)+32)) = v3980
	v3982 = *(*int64)(unsafe.Add(mBase, uint32(v3939)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3969)+40)) = v3982
	*(*int32)(unsafe.Add(mBase, uint32(v3940)+8)) = v3969
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v3969
	goto L1216
L1224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
