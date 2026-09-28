package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_verify_heapam(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int64
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int64
	_ = v307
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int64
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int64
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v375 int64
	_ = v375
	var v383 int64
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int64
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int64
	_ = v402
	var v409 int64
	_ = v409
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v736 int32
	_ = v736
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int64
	_ = v749
	var v750 int64
	_ = v750
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int64
	_ = v788
	var v789 int64
	_ = v789
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int64
	_ = v835
	var v836 int64
	_ = v836
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int64
	_ = v872
	var v873 int64
	_ = v873
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int64
	_ = v909
	var v910 int64
	_ = v910
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
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
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int64
	_ = v961
	var v962 int64
	_ = v962
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int64
	_ = v1002
	var v1003 int64
	_ = v1003
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int64
	_ = v1045
	var v1046 int64
	_ = v1046
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int64
	_ = v1111
	var v1112 int64
	_ = v1112
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int64
	_ = v1156
	var v1157 int64
	_ = v1157
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1188 int32
	_ = v1188
	var v1194 int32
	_ = v1194
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int64
	_ = v1209
	var v1210 int64
	_ = v1210
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1245 int32
	_ = v1245
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int64
	_ = v1255
	var v1256 int64
	_ = v1256
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int64
	_ = v1319
	var v1320 int64
	_ = v1320
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int64
	_ = v1358
	var v1359 int64
	_ = v1359
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int64
	_ = v1400
	var v1401 int64
	_ = v1401
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int64
	_ = v1440
	var v1441 int64
	_ = v1441
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1457 int32
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1494 int64
	_ = v1494
	var v1497 int64
	_ = v1497
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1507 int64
	_ = v1507
	var v1508 int64
	_ = v1508
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1524 int32
	_ = v1524
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1536 int64
	_ = v1536
	var v1539 int64
	_ = v1539
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int64
	_ = v1549
	var v1550 int64
	_ = v1550
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1592 int32
	_ = v1592
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1607 int64
	_ = v1607
	var v1610 int64
	_ = v1610
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1622 int64
	_ = v1622
	var v1625 int64
	_ = v1625
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1637 int64
	_ = v1637
	var v1640 int64
	_ = v1640
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1689 int64
	_ = v1689
	var v1692 int64
	_ = v1692
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1704 int64
	_ = v1704
	var v1707 int64
	_ = v1707
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1719 int64
	_ = v1719
	var v1722 int64
	_ = v1722
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1745 int32
	_ = v1745
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1757 int32
	_ = v1757
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1769 int32
	_ = v1769
	var v1773 int32
	_ = v1773
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1783 int32
	_ = v1783
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1798 int32
	_ = v1798
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1835 int32
	_ = v1835
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1864 int64
	_ = v1864
	var v1867 int64
	_ = v1867
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1877 int32
	_ = v1877
	var v1879 int64
	_ = v1879
	var v1882 int64
	_ = v1882
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1894 int64
	_ = v1894
	var v1897 int64
	_ = v1897
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1927 int32
	_ = v1927
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1937 int64
	_ = v1937
	var v1940 int64
	_ = v1940
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1950 int32
	_ = v1950
	var v1952 int64
	_ = v1952
	var v1955 int64
	_ = v1955
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1967 int64
	_ = v1967
	var v1970 int64
	_ = v1970
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2001 int64
	_ = v2001
	var v2004 int64
	_ = v2004
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2014 int64
	_ = v2014
	var v2015 int64
	_ = v2015
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2031 int32
	_ = v2031
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2050 int64
	_ = v2050
	var v2051 int64
	_ = v2051
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2067 int32
	_ = v2067
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2098 int64
	_ = v2098
	var v2099 int64
	_ = v2099
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2115 int32
	_ = v2115
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2126 int32
	_ = v2126
	var v2135 int32
	_ = v2135
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2178 int64
	_ = v2178
	var v2179 int64
	_ = v2179
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2195 int32
	_ = v2195
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2206 int32
	_ = v2206
	var v2212 int32
	_ = v2212
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2228 int32
	_ = v2228
	var v2230 int32
	_ = v2230
	var v2235 int32
	_ = v2235
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2255 int64
	_ = v2255
	var v2256 int64
	_ = v2256
	var v2265 int32
	_ = v2265
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2272 int32
	_ = v2272
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2295 int32
	_ = v2295
	var v2297 int32
	_ = v2297
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2317 int64
	_ = v2317
	var v2318 int64
	_ = v2318
	var v2327 int32
	_ = v2327
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2334 int32
	_ = v2334
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2345 int32
	_ = v2345
	var v2349 int32
	_ = v2349
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2371 int64
	_ = v2371
	var v2372 int64
	_ = v2372
	var v2381 int32
	_ = v2381
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2388 int32
	_ = v2388
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2399 int32
	_ = v2399
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2419 int64
	_ = v2419
	var v2420 int64
	_ = v2420
	var v2429 int32
	_ = v2429
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2436 int32
	_ = v2436
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2444 int32
	_ = v2444
	var v2445 int32
	_ = v2445
	var v2456 int32
	_ = v2456
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2473 int64
	_ = v2473
	var v2474 int64
	_ = v2474
	var v2483 int32
	_ = v2483
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2490 int32
	_ = v2490
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2518 int64
	_ = v2518
	var v2519 int64
	_ = v2519
	var v2528 int32
	_ = v2528
	var v2530 int32
	_ = v2530
	var v2531 int32
	_ = v2531
	var v2535 int32
	_ = v2535
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2546 int32
	_ = v2546
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2557 int32
	_ = v2557
	var v2558 int32
	_ = v2558
	var v2560 int32
	_ = v2560
	var v2561 int32
	_ = v2561
	var v2564 int32
	_ = v2564
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2571 int32
	_ = v2571
	var v2572 int64
	_ = v2572
	var v2574 int64
	_ = v2574
	var v2576 int32
	_ = v2576
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2594 int32
	_ = v2594
	var v2596 int32
	_ = v2596
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
	var v2621 int32
	_ = v2621
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2648 int32
	_ = v2648
	var v2651 int32
	_ = v2651
	var v2657 int32
	_ = v2657
	var v2685 int32
	_ = v2685
	var v2687 int32
	_ = v2687
	var v2692 int32
	_ = v2692
	var v2695 int32
	_ = v2695
	var v2700 int32
	_ = v2700
	var v2720 int32
	_ = v2720
	var v2724 int32
	_ = v2724
	var v2730 int32
	_ = v2730
	var v2733 int32
	_ = v2733
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2738 int32
	_ = v2738
	var v2739 int32
	_ = v2739
	var v2743 int32
	_ = v2743
	var v2751 int32
	_ = v2751
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2763 int64
	_ = v2763
	var v2764 int64
	_ = v2764
	var v2773 int32
	_ = v2773
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2780 int32
	_ = v2780
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2806 int32
	_ = v2806
	var v2807 int32
	_ = v2807
	var v2808 int32
	_ = v2808
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2822 int32
	_ = v2822
	var v2823 int32
	_ = v2823
	var v2824 int32
	_ = v2824
	var v2825 int32
	_ = v2825
	var v2826 int32
	_ = v2826
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2835 int32
	_ = v2835
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2841 int32
	_ = v2841
	var v2843 int32
	_ = v2843
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2857 int32
	_ = v2857
	var v2858 int32
	_ = v2858
	var v2859 int32
	_ = v2859
	var v2861 int32
	_ = v2861
	var v2866 int32
	_ = v2866
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2876 int32
	_ = v2876
	var v2877 int32
	_ = v2877
	var v2878 int64
	_ = v2878
	var v2879 int64
	_ = v2879
	var v2888 int32
	_ = v2888
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2895 int32
	_ = v2895
	var v2900 int32
	_ = v2900
	var v2901 int32
	_ = v2901
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2906 int32
	_ = v2906
	var v2915 int32
	_ = v2915
	var v2922 int32
	_ = v2922
	var v2923 int32
	_ = v2923
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2927 int64
	_ = v2927
	var v2928 int64
	_ = v2928
	var v2937 int32
	_ = v2937
	var v2939 int32
	_ = v2939
	var v2940 int32
	_ = v2940
	var v2944 int32
	_ = v2944
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2970 int32
	_ = v2970
	var v2972 int32
	_ = v2972
	var v2976 int32
	_ = v2976
	var v2977 int32
	_ = v2977
	var v2980 int32
	_ = v2980
	var v2984 int32
	_ = v2984
	var v2988 int32
	_ = v2988
	var v2989 int32
	_ = v2989
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3004 int32
	_ = v3004
	var v3005 int64
	_ = v3005
	var v3006 int64
	_ = v3006
	var v3015 int32
	_ = v3015
	var v3017 int32
	_ = v3017
	var v3018 int32
	_ = v3018
	var v3022 int32
	_ = v3022
	var v3027 int32
	_ = v3027
	var v3028 int32
	_ = v3028
	var v3030 int32
	_ = v3030
	var v3031 int32
	_ = v3031
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3040 int32
	_ = v3040
	var v3042 int32
	_ = v3042
	var v3044 int32
	_ = v3044
	var v3048 int32
	_ = v3048
	var v3052 int32
	_ = v3052
	var v3056 int32
	_ = v3056
	var v3060 int32
	_ = v3060
	var v3067 int32
	_ = v3067
	var v3068 int32
	_ = v3068
	var v3075 int32
	_ = v3075
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3085 int32
	_ = v3085
	var v3086 int32
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3088 int64
	_ = v3088
	var v3089 int64
	_ = v3089
	var v3098 int32
	_ = v3098
	var v3100 int32
	_ = v3100
	var v3101 int32
	_ = v3101
	var v3105 int32
	_ = v3105
	var v3110 int32
	_ = v3110
	var v3111 int32
	_ = v3111
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3125 int32
	_ = v3125
	var v3127 int32
	_ = v3127
	var v3132 int32
	_ = v3132
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3159 int32
	_ = v3159
	var v3163 int32
	_ = v3163
	var v3167 int32
	_ = v3167
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3176 int32
	_ = v3176
	var v3184 int32
	_ = v3184
	var v3189 int32
	_ = v3189
	var v3190 int32
	_ = v3190
	var v3191 int32
	_ = v3191
	var v3192 int32
	_ = v3192
	var v3193 int32
	_ = v3193
	var v3194 int64
	_ = v3194
	var v3195 int64
	_ = v3195
	var v3204 int32
	_ = v3204
	var v3206 int32
	_ = v3206
	var v3207 int32
	_ = v3207
	var v3211 int32
	_ = v3211
	var v3216 int32
	_ = v3216
	var v3217 int32
	_ = v3217
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3222 int32
	_ = v3222
	var v3224 int32
	_ = v3224
	var v3229 int32
	_ = v3229
	var v3232 int32
	_ = v3232
	var v3254 int32
	_ = v3254
	var v3256 int32
	_ = v3256
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3270 int32
	_ = v3270
	var v3282 int32
	_ = v3282
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3289 int32
	_ = v3289
	var v3293 int64
	_ = v3293
	var v3295 int32
	_ = v3295
	var v3297 int32
	_ = v3297
	var v3301 int32
	_ = v3301
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3305 int32
	_ = v3305
	var v3307 int32
	_ = v3307
	var v3308 int32
	_ = v3308
	var v3310 int32
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3319 int32
	_ = v3319
	var v3321 int32
	_ = v3321
	var v3337 int32
	_ = v3337
	var v3338 int32
	_ = v3338
	var v3341 int64
	_ = v3341
	var v3342 int32
	_ = v3342
	var v3343 int32
	_ = v3343
	var v3346 int32
	_ = v3346
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3353 int32
	_ = v3353
	var v3355 int32
	_ = v3355
	var v3362 int32
	_ = v3362
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3365 int64
	_ = v3365
	var v3366 int64
	_ = v3366
	var v3367 int32
	_ = v3367
	var v3368 int32
	_ = v3368
	var v3377 int32
	_ = v3377
	var v3379 int32
	_ = v3379
	var v3380 int32
	_ = v3380
	var v3384 int32
	_ = v3384
	var v3389 int32
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3392 int32
	_ = v3392
	var v3393 int32
	_ = v3393
	var v3400 int32
	_ = v3400
	var v3401 int32
	_ = v3401
	var v3404 int64
	_ = v3404
	var v3405 int32
	_ = v3405
	var v3406 int32
	_ = v3406
	var v3407 int32
	_ = v3407
	var v3408 int32
	_ = v3408
	var v3411 int32
	_ = v3411
	var v3417 int32
	_ = v3417
	var v3418 int32
	_ = v3418
	var v3419 int32
	_ = v3419
	var v3420 int32
	_ = v3420
	var v3425 int32
	_ = v3425
	var v3434 int32
	_ = v3434
	var v3438 int32
	_ = v3438
	var v3440 int32
	_ = v3440
	var v3447 int32
	_ = v3447
	var v3448 int32
	_ = v3448
	var v3451 int32
	_ = v3451
	var v3453 int32
	_ = v3453
	var v3461 int32
	_ = v3461
	var v3462 int32
	_ = v3462
	var v3463 int32
	_ = v3463
	var v3464 int32
	_ = v3464
	var v3471 int32
	_ = v3471
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3474 int32
	_ = v3474
	var v3480 int32
	_ = v3480
	var v3481 int64
	_ = v3481
	var v3482 int64
	_ = v3482
	var v3483 int32
	_ = v3483
	var v3484 int32
	_ = v3484
	var v3493 int32
	_ = v3493
	var v3495 int32
	_ = v3495
	var v3496 int32
	_ = v3496
	var v3500 int32
	_ = v3500
	var v3505 int32
	_ = v3505
	var v3506 int32
	_ = v3506
	var v3508 int32
	_ = v3508
	var v3509 int32
	_ = v3509
	var v3511 int32
	_ = v3511
	var v3519 int32
	_ = v3519
	var v3520 int32
	_ = v3520
	var v3522 int32
	_ = v3522
	var v3524 int32
	_ = v3524
	var v3531 int32
	_ = v3531
	var v3532 int32
	_ = v3532
	var v3534 int32
	_ = v3534
	var v3535 int32
	_ = v3535
	var v3540 int32
	_ = v3540
	var v3541 int32
	_ = v3541
	var v3545 int32
	_ = v3545
	var v3562 int32
	_ = v3562
	var v3563 int64
	_ = v3563
	var v3564 int64
	_ = v3564
	var v3565 int32
	_ = v3565
	var v3566 int32
	_ = v3566
	var v3575 int32
	_ = v3575
	var v3577 int32
	_ = v3577
	var v3578 int32
	_ = v3578
	var v3582 int32
	_ = v3582
	var v3587 int32
	_ = v3587
	var v3588 int32
	_ = v3588
	var v3590 int32
	_ = v3590
	var v3591 int32
	_ = v3591
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3617 int32
	_ = v3617
	var v3638 int32
	_ = v3638
	var v3640 int32
	_ = v3640
	var v3665 int32
	_ = v3665
	var v3691 int32
	_ = v3691
	var v3692 int32
	_ = v3692
	var v3694 int32
	_ = v3694
	var v3695 int32
	_ = v3695
	var v3696 int32
	_ = v3696
	var v3698 int32
	_ = v3698
	var v3699 int32
	_ = v3699
	var v3702 int32
	_ = v3702
	var v3703 int32
	_ = v3703
	var v3706 int32
	_ = v3706
	var v3727 int32
	_ = v3727
	var v3737 int32
	_ = v3737
	var v3744 int32
	_ = v3744
	var v3749 int32
	_ = v3749
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(_a_F_verify_heapam_0)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[0]))) = v2
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v28 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3737 = m.ExcPending
	if v3737 != 0 {
		goto L19
	} else {
		goto L760
	}
L2:
	;
	v3727 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v3727)
	m.G0 = v23 + int32(_a_F_verify_heapam_0)
	return int64(0)
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[1]))) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[2]))) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[3]))) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[4]))) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[5]))) = v23 + int32(_a_F_verify_heapam_1)
	if v184 != 0 {
		goto L163
	} else {
		goto L164
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[6]))) = base.I64_extend_i32_u(v393)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v392)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[7]))) = v556
	goto L3
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L19
	} else {
		goto L159
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L19
	} else {
		goto L155
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L19
	} else {
		goto L151
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L19
	} else {
		goto L146
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L19
	} else {
		goto L142
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L19
	} else {
		goto L138
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L19
	} else {
		goto L134
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v31 == int32(1) {
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
	v417 = m.ExcPending
	if v417 != 0 {
		goto L19
	} else {
		goto L130
	}
L15:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
	if v34 == int32(1) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
	if v37 == int32(1) {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v45 = F_pg_detoast_datum_packed(m, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	base.MemoryFill(m, v23+int32(_a_F_verify_heapam_2), int32(0), int32(144))
	v191 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L19
	} else {
		goto L66
	}
L19:
	;
	return int64(0)
L20:
	;
	v49 = F_text_to_cstring(m, v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v54 = v49
	v55 = int32(_a_F_verify_heapam_3)
	goto L23
L22:
	;
	if v92 == int32(0) {
		v184 = v2
		v185 = int32(1)
		goto L18
	} else {
		goto L35
	}
L23:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v58 == v59 {
		v81 = v58
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v92 = int32(0)
	goto L22
L25:
	;
	v83 = int32(1)
	if v81 != 0 {
		v54 = v54 + v83
		v55 = v55 + v83
		goto L23
	} else {
		goto L34
	}
L26:
	;
	if base.Ui32((v58-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v69 = v58 | int32(32)
	goto L29
L28:
	;
	v69 = v58
	goto L29
L29:
	;
	if base.Ui32((v59-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v78 = v59 | int32(32)
	goto L32
L31:
	;
	v78 = v59
	goto L32
L32:
	;
	if v69 == v78 {
		v81 = v69
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v92 = v69 - v78
	goto L22
L34:
	;
	goto L24
L35:
	;
	v98 = v49
	v99 = int32(_a_F_verify_heapam_4)
	goto L37
L36:
	;
	if v136 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L37:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v102 == v103 {
		v125 = v102
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v136 = int32(0)
	goto L36
L39:
	;
	v127 = int32(1)
	if v125 != 0 {
		v98 = v98 + v127
		v99 = v99 + v127
		goto L37
	} else {
		goto L48
	}
L40:
	;
	if base.Ui32((v102-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v113 = v102 | int32(32)
	goto L43
L42:
	;
	v113 = v102
	goto L43
L43:
	;
	if base.Ui32((v103-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v122 = v103 | int32(32)
	goto L46
L45:
	;
	v122 = v103
	goto L46
L46:
	;
	if v113 == v122 {
		v125 = v113
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v136 = v113 - v122
	goto L36
L48:
	;
	goto L38
L49:
	;
	v184 = v2
	v185 = int32(0)
	goto L18
L50:
	;
	goto L51
L51:
	;
	v143 = v49
	v144 = int32(_a_F_verify_heapam_5)
	goto L53
L52:
	;
	if v181 != 0 {
		goto L8
	} else {
		goto L65
	}
L53:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v147 == v148 {
		v170 = v147
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v181 = int32(0)
	goto L52
L55:
	;
	v172 = int32(1)
	if v170 != 0 {
		v143 = v143 + v172
		v144 = v144 + v172
		goto L53
	} else {
		goto L64
	}
L56:
	;
	if base.Ui32((v147-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v158 = v147 | int32(32)
	goto L59
L58:
	;
	v158 = v147
	goto L59
L59:
	;
	if base.Ui32((v148-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v167 = v148 | int32(32)
	goto L62
L61:
	;
	v167 = v148
	goto L62
L62:
	;
	if v158 == v167 {
		v170 = v158
		goto L55
	} else {
		goto L63
	}
L63:
	;
	v181 = v158 - v167
	goto L52
L64:
	;
	goto L54
L65:
	;
	v184 = int32(1)
	v185 = int32(2)
	goto L18
L66:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v194 = int32(_a_F_verify_heapam_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))) = uint16(v194)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[9]))) = v193
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L19
	} else {
		goto L67
	}
L67:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10]))) = v200
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11]))) = v202
	v205 = F_relation_open(m, v40, int32(1))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L19
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[12]))) = v205
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v205)+48))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+119)))
	switch v209 - int32(83) {
	case 0:
		goto L69
	default:
		goto L71
	case 26, 31, 33:
		goto L70
	}
L69:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+118)))
	if v238 != int32(117) {
		goto L78
	} else {
		goto L79
	}
L70:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v208)+84))
	if v235 != int32(2) {
		goto L7
	} else {
		goto L77
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L19
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L19
	} else {
		goto L73
	}
L73:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v205)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v219 + int32(4)
	F_errmsg(m, int32(_a_F_verify_heapam_7), v23)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L19
	} else {
		goto L74
	}
L74:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v205)+48))
	v227 = int32(*(*int8)(unsafe.Add(mBase, uint32(v226)+119)))
	F_errdetail_relkind_not_supported(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L19
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(341), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
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
	v280 = int32(0)
	v282 = F_RelationGetNumberOfBlocksInFork(m, v205, v280)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L19
	} else {
		goto L93
	}
L79:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_verify_heapam[13])))
	if v243 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v253 == int32(0) {
		goto L78
	} else {
		goto L84
	}
L81:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[14]))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+308))
	v251 = base.B2i32(v249 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_verify_heapam[13])) = uint8(v251)
	v253 = v251
	goto L83
L82:
	;
	v253 = int32(0)
	goto L83
L83:
	;
	goto L80
L84:
	;
	v258 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L19
	} else {
		goto L85
	}
L85:
	;
	if v258 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_errcode(m, int32(100663618))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L19
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	F_relation_close(m, v205, int32(1))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L19
	} else {
		goto L92
	}
L89:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v205)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v263 + int32(4)
	F_errmsg(m, int32(_a_F_verify_heapam_10), v23+int32(16))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L19
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(364), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
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
	if v282 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	F_relation_close(m, v205, int32(1))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L19
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v290 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L19
	} else {
		goto L98
	}
L97:
	;
	goto L2
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[15]))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[16]))) = v290
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	if v295 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v298 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	if base.Ui64(base.I64_extend_i32_u(v282)) <= base.Ui64(v298) {
		goto L6
	} else {
		goto L102
	}
L100:
	;
	v302 = v280
	goto L101
L101:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+112)))
	if v304 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v302 = base.I32_wrap_i64(v298)
	goto L101
L103:
	;
	v307 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui64(base.I64_extend_i32_u(v282)) <= base.Ui64(v307) {
		goto L5
	} else {
		goto L106
	}
L104:
	;
	v313 = v282
	goto L105
L105:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v205)+48))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+112))
	v317 = int32(0)
	if base.B2i32(v316 == v317)|base.B2i32(v42 == int64(0)) == v317 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v313 = base.I32_wrap_i64(v307) + int32(1)
	goto L105
L107:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[17]))
	v351 = F_LWLockAcquire(m, v347+int32(384), int32(1))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L19
	} else {
		goto L113
	}
L108:
	;
	v325 = F_table_open(m, v316, int32(1))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L19
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[18]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[19]))) = int64(0)
	goto L107
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[19]))) = v325
	v333 = F_toast_open_indexes(m, v325, int32(1), v23+int32(_a_F_verify_heapam_11), v23+int32(_a_F_verify_heapam_12))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L19
	} else {
		goto L112
	}
L112:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[20])))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v335+v333<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[21]))) = v339
	goto L107
L113:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[22]))
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v354)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[23]))) = v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v354)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[24]))) = v357
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[17]))
	F_LWLockRelease(m, v360+int32(384))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L19
	} else {
		goto L114
	}
L114:
	;
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[23])))
	v366 = base.I32_wrap_i64(v365)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[25]))) = v366
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[24])))
	if base.Ui32(v368) <= base.Ui32(int32(2)) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[26]))) = v383
	v386 = v23 + int32(_a_F_verify_heapam_13)
	v388 = v23 + int32(_a_F_verify_heapam_14)
	F_ReadMultiXactIdRange(m, v386, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L19
	} else {
		goto L123
	}
L116:
	;
	v383 = base.I64_extend_i32_u(v368)
	goto L115
L117:
	;
	goto L118
L118:
	;
	v372 = v366 - v368
	if int32(0) < v372 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v375 = int64(3)
	if base.Ui64(v365-v375) < base.Ui64(base.I64_extend_i32_u(v372)) {
		v383 = v375
		goto L115
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v383 = v365 - base.I64_extend_i32_s(v372)
	goto L115
L122:
	;
	goto L121
L123:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[12])))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+48))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[27]))) = v393
	if base.Ui32(v393) < base.Ui32(int32(3)) {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v397 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[23])))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[25])))
	v399 = v398 - v393
	if int32(0) < v399 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[6]))) = v409
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v392)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[24]))) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[7]))) = v411
	goto L3
L126:
	;
	v402 = int64(3)
	if base.Ui64(v397-v402) < base.Ui64(base.I64_extend_i32_u(v399)) {
		v409 = v402
		goto L125
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v409 = v397 - base.I64_extend_i32_s(v399)
	goto L125
L129:
	;
	goto L128
L130:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L19
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_verify_heapam_15), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L19
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(275), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
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
	v436 = m.ExcPending
	if v436 != 0 {
		goto L19
	} else {
		goto L135
	}
L135:
	;
	F_errmsg(m, int32(_a_F_verify_heapam_16), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L19
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(281), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
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
	v452 = m.ExcPending
	if v452 != 0 {
		goto L19
	} else {
		goto L139
	}
L139:
	;
	F_errmsg(m, int32(_a_F_verify_heapam_17), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L19
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(287), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
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
	v468 = m.ExcPending
	if v468 != 0 {
		goto L19
	} else {
		goto L143
	}
L143:
	;
	F_errmsg(m, int32(_a_F_verify_heapam_18), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L19
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(293), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
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
	v484 = m.ExcPending
	if v484 != 0 {
		goto L19
	} else {
		goto L147
	}
L147:
	;
	F_errmsg(m, int32(_a_F_verify_heapam_19), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L19
	} else {
		goto L148
	}
L148:
	;
	F_errhint(m, int32(_a_F_verify_heapam_20), int32(0))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L19
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(305), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
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
	v504 = m.ExcPending
	if v504 != 0 {
		goto L19
	} else {
		goto L152
	}
L152:
	;
	F_errmsg(m, int32(_a_F_verify_heapam_21), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L19
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(351), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
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
	v520 = m.ExcPending
	if v520 != 0 {
		goto L19
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+1024)) = v282 - int32(1)
	F_errmsg(m, int32(_a_F_verify_heapam_22), v23+int32(1024))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L19
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(392), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
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
	v540 = m.ExcPending
	if v540 != 0 {
		goto L19
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+1008)) = v282 - int32(1)
	F_errmsg(m, int32(_a_F_verify_heapam_23), v23+int32(1008))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L19
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_8), int32(405), int32(_a_F_verify_heapam_9))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
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
	v570 = int32(14)
	goto L165
L164:
	;
	v570 = int32(0)
	goto L165
L165:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[16])))
	if v184 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v575 = int32(3)
	goto L168
L167:
	;
	v575 = int32(_a_F_verify_heapam_24)
	goto L168
L168:
	;
	v579 = F_read_stream_begin_relation(m, v570, v571, v391, int32(0), v575, v23+int32(_a_F_verify_heapam_25), int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L19
	} else {
		goto L169
	}
L169:
	;
	goto L170
L170:
	;
	v602 = F_read_stream_next_buffer(m, v579, int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L19
	} else {
		goto L173
	}
L171:
	;
	F_read_stream_end(m, v579)
	mBase = m.M
	v3691 = m.ExcPending
	if v3691 != 0 {
		goto L19
	} else {
		goto L746
	}
L172:
	;
	goto L171
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[15]))) = v602
	if v602 == int32(0) {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v608 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[28]))
	if v608 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L19
	} else {
		goto L178
	}
L176:
	;
	v612 = v602
	goto L177
L177:
	;
	base.MemoryFill(m, v23+int32(_a_F_verify_heapam_26), int32(0), int32(_a_F_verify_heapam_27))
	F_LockBufferInternal(m, v612, int32(1))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L19
	} else {
		goto L179
	}
L178:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[15])))
	v612 = v611
	goto L177
L179:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[15])))
	if v621 < int32(0) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))) = v640
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[15])))
	if v642 < int32(0) {
		goto L185
	} else {
		goto L186
	}
L181:
	;
	v625 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[30]))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v625+(v621^int32(-1))*int32(56))+16))
	v640 = v631
	goto L180
L182:
	;
	goto L183
L183:
	;
	v633 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[31]))
	v634 = int32(56)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v633+v621*v634-v634)+16))
	v640 = v639
	goto L180
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[32]))) = v660
	v662 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v660)+12)))
	v663 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))) = uint16(v663)
	if base.Ui32(int32(25)) <= base.Ui32(v662) {
		goto L189
	} else {
		goto L190
	}
L185:
	;
	v646 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[34]))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v646+(v642^int32(-1))<<(uint(int32(2))%32))))
	v660 = v652
	goto L184
L186:
	;
	goto L187
L187:
	;
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_verify_heapam[35]))
	v660 = v654 + v642<<(uint(int32(13))%32) + int32(-8192)
	goto L184
L188:
	;
	v3254 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[15])))
	F_UnlockReleaseBuffer(m, v3254)
	mBase = m.M
	v3256 = m.ExcPending
	if v3256 != 0 {
		goto L19
	} else {
		goto L671
	}
L189:
	;
	v673 = int32(base.Ui32(v662+int32(_a_F_verify_heapam_28)) >> (uint(int32(2)) % 32))
	goto L191
L190:
	;
	v673 = int32(0)
	goto L191
L191:
	;
	v675 = v673 & int32(_a_F_verify_heapam_6)
	if v675 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v678 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))) = uint16(v678)
	v680 = int32(_a_F_verify_heapam_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))) = uint16(v680)
	goto L188
L193:
	;
	goto L194
L194:
	;
	v685 = v663
	goto L195
L195:
	;
	v703 = v685 & int32(_a_F_verify_heapam_6)
	v706 = v703 + (v23 + int32(_a_F_verify_heapam_29))
	v707 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v706))) = uint8(v707)
	v711 = v23 + int32(_a_F_verify_heapam_30) + v703
	*(*uint8)(unsafe.Add(mBase, uint32(v711))) = uint8(v707)
	v716 = int32(1)
	v718 = v23 + int32(_a_F_verify_heapam_31) + v703<<(uint(v716)%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v718))) = uint16(v707)
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[32])))
	v723 = v721 + int32(20)
	v725 = v703 << (uint(int32(2)) % 32)
	v726 = v723 + v725
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[36]))) = v726
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	switch int32(base.Ui32(v728)>>(uint(int32(15))%32))&int32(3) - v716 {
	case 0:
		goto L198
	case 1:
		goto L199
	default:
		goto L197
	}
L196:
	;
	v2692 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))) = uint16(v2692)
	v2695 = int32(_a_F_verify_heapam_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))) = uint16(v2695)
	v2700 = v2692
	goto L578
L197:
	;
	v2685 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	v2687 = v2685 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))) = uint16(v2687)
	if base.Ui32(v2687&int32(_a_F_verify_heapam_6)) <= base.Ui32(v675) {
		v685 = v2687
		goto L195
	} else {
		goto L577
	}
L198:
	;
	v941 = int32(base.Ui32(v728) >> (uint(int32(17)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[37]))) = uint16(v941)
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	v945 = v943 & int32(_a_F_verify_heapam_32)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[38]))) = uint16(v945)
	if (v945+int32(7))&int32(_a_F_verify_heapam_33) != v945 {
		goto L235
	} else {
		goto L236
	}
L199:
	;
	v736 = v728 & int32(_a_F_verify_heapam_32)
	if v736 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+288)) = int64(4294967296)
	v744 = F_psprintf(m, int32(_a_F_verify_heapam_34), v23+int32(288))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L19
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	if base.Ui32(v675) < base.Ui32(v736) {
		goto L208
	} else {
		goto L209
	}
L203:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v748 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v749 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v750 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v750
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v749
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v748))
	v759 = int32(base.Ui32(v748) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v759)
	v761 = F_cstring_to_text(m, v744)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L19
	} else {
		goto L204
	}
L204:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v761)
	F_pfree(m, v744)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L19
	} else {
		goto L205
	}
L205:
	;
	v771 = F_heap_form_tuple(m, v747, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L19
	} else {
		goto L206
	}
L206:
	;
	F_tuplestore_puttuple(m, v746, v771)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L19
	} else {
		goto L207
	}
L207:
	;
	v775 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v775)
	goto L197
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+308)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v23)+304)) = v736
	v783 = F_psprintf(m, int32(_a_F_verify_heapam_37), v23+int32(304))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L19
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v723+v736<<(uint(int32(2))%32))))
	switch int32(base.Ui32(v819)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L216
	case 1:
		goto L217
	case 2:
		goto L218
	default:
		goto L219
	}
L211:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v787 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v788 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v789 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v788
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v787))
	v798 = int32(base.Ui32(v787) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v798)
	v800 = F_cstring_to_text(m, v783)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L19
	} else {
		goto L212
	}
L212:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v800)
	F_pfree(m, v783)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L19
	} else {
		goto L213
	}
L213:
	;
	v810 = F_heap_form_tuple(m, v786, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L19
	} else {
		goto L214
	}
L214:
	;
	F_tuplestore_puttuple(m, v785, v810)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L19
	} else {
		goto L215
	}
L215:
	;
	v814 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v814)
	goto L197
L216:
	;
	v937 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v706))) = uint8(v937)
	*(*uint16)(unsafe.Add(mBase, uint32(v718))) = uint16(v736)
	goto L197
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+352)) = v736
	v904 = F_psprintf(m, int32(_a_F_verify_heapam_38), v23+int32(352))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L19
	} else {
		goto L230
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+336)) = v736
	v867 = F_psprintf(m, int32(_a_F_verify_heapam_39), v23+int32(336))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L19
	} else {
		goto L225
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+320)) = v736
	v830 = F_psprintf(m, int32(_a_F_verify_heapam_40), v23+int32(320))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L19
	} else {
		goto L220
	}
L220:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v834 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v835 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v836 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v836
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v835
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v834))
	v845 = int32(base.Ui32(v834) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v845)
	v847 = F_cstring_to_text(m, v830)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L19
	} else {
		goto L221
	}
L221:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v847)
	F_pfree(m, v830)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L19
	} else {
		goto L222
	}
L222:
	;
	v857 = F_heap_form_tuple(m, v833, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L19
	} else {
		goto L223
	}
L223:
	;
	F_tuplestore_puttuple(m, v832, v857)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L19
	} else {
		goto L224
	}
L224:
	;
	v861 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v861)
	goto L197
L225:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v871 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v872 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v873 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v873
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v872
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v871))
	v882 = int32(base.Ui32(v871) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v882)
	v884 = F_cstring_to_text(m, v867)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L19
	} else {
		goto L226
	}
L226:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v884)
	F_pfree(m, v867)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L19
	} else {
		goto L227
	}
L227:
	;
	v894 = F_heap_form_tuple(m, v870, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L19
	} else {
		goto L228
	}
L228:
	;
	F_tuplestore_puttuple(m, v869, v894)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L19
	} else {
		goto L229
	}
L229:
	;
	v898 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v898)
	goto L197
L230:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v908 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v909 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v910 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v910
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v909
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v908))
	v919 = int32(base.Ui32(v908) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v919)
	v921 = F_cstring_to_text(m, v904)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L19
	} else {
		goto L231
	}
L231:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v921)
	F_pfree(m, v904)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L19
	} else {
		goto L232
	}
L232:
	;
	v931 = F_heap_form_tuple(m, v907, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L19
	} else {
		goto L233
	}
L233:
	;
	F_tuplestore_puttuple(m, v906, v931)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L19
	} else {
		goto L234
	}
L234:
	;
	v935 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v935)
	goto L197
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+992)) = v945
	v956 = F_psprintf(m, int32(_a_F_verify_heapam_41), v23+int32(992))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L19
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	if base.Ui32(v728) <= base.Ui32(int32(_a_F_verify_heapam_42)) {
		goto L243
	} else {
		goto L244
	}
L238:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v960 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v961 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v962 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v962
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v961
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v960))
	v971 = int32(base.Ui32(v960) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v971)
	v973 = F_cstring_to_text(m, v956)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L19
	} else {
		goto L239
	}
L239:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v973)
	F_pfree(m, v956)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L19
	} else {
		goto L240
	}
L240:
	;
	v983 = F_heap_form_tuple(m, v959, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L19
	} else {
		goto L241
	}
L241:
	;
	F_tuplestore_puttuple(m, v958, v983)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L19
	} else {
		goto L242
	}
L242:
	;
	v987 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v987)
	goto L197
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+372)) = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+368)) = v941
	v997 = F_psprintf(m, int32(_a_F_verify_heapam_43), v23+int32(368))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L19
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	if base.Ui32(int32(_a_F_verify_heapam_44)) <= base.Ui32(v945+v941) {
		goto L251
	} else {
		goto L252
	}
L246:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v1001 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v1002 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v1003 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v1003
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v1002
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v1001))
	v1012 = int32(base.Ui32(v1001) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v1012)
	v1014 = F_cstring_to_text(m, v997)
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L19
	} else {
		goto L247
	}
L247:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v1014)
	F_pfree(m, v997)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L19
	} else {
		goto L248
	}
L248:
	;
	v1024 = F_heap_form_tuple(m, v1000, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L19
	} else {
		goto L249
	}
L249:
	;
	F_tuplestore_puttuple(m, v999, v1024)
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L19
	} else {
		goto L250
	}
L250:
	;
	v1028 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v1028)
	goto L197
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+392)) = int32(_a_F_verify_heapam_45)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+388)) = v941
	*(*int32)(unsafe.Add(mBase, uint32(v23)+384)) = v945
	v1040 = F_psprintf(m, int32(_a_F_verify_heapam_46), v23+int32(384))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L19
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v1073 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v706))) = uint8(v1073)
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	v1078 = v721 + v1075&int32(_a_F_verify_heapam_32)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[46]))) = v1078
	v1080 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1078)+18)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[47]))) = v1080 & int32(2047)
	v1084 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1078)+20)))
	if v1084&int32(_a_F_verify_heapam_47) == int32(_a_F_verify_heapam_27) {
		goto L260
	} else {
		goto L261
	}
L254:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v1044 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v1045 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v1046 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v1046
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v1045
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v1044))
	v1055 = int32(base.Ui32(v1044) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v1055)
	v1057 = F_cstring_to_text(m, v1040)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L19
	} else {
		goto L255
	}
L255:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v1057)
	F_pfree(m, v1040)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L19
	} else {
		goto L256
	}
L256:
	;
	v1067 = F_heap_form_tuple(m, v1043, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L19
	} else {
		goto L257
	}
L257:
	;
	F_tuplestore_puttuple(m, v1042, v1067)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L19
	} else {
		goto L258
	}
L258:
	;
	v1071 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v1071)
	goto L197
L259:
	;
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1096)+22)))
	v1099 = v1095 & int32(_a_F_verify_heapam_6)
	if base.Ui32(v1099) < base.Ui32(v1097) {
		goto L264
	} else {
		goto L265
	}
L260:
	;
	v1089 = F_HeapTupleGetUpdateXid(m, v1078)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L19
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+4))
	v1094 = v1093
	v1095 = v941
	v1096 = v1078
	goto L259
L263:
	;
	v1091 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[37]))))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[46])))
	v1094 = v1089
	v1095 = v1091
	v1096 = v1092
	goto L259
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+980)) = v1099
	*(*int32)(unsafe.Add(mBase, uint32(v23)+976)) = v1097
	v1106 = F_psprintf(m, int32(_a_F_verify_heapam_48), v23+int32(976))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L19
	} else {
		goto L267
	}
L265:
	;
	v1144 = v1096
	goto L266
L266:
	;
	v1145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1144)+20)))
	v1146 = int32(_a_F_verify_heapam_49)
	if v1145&v1146 == v1146 {
		goto L272
	} else {
		goto L273
	}
L267:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v1110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v1111 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v1112 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v1112
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v1111
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v1110))
	v1121 = int32(base.Ui32(v1110) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v1121)
	v1123 = F_cstring_to_text(m, v1106)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L19
	} else {
		goto L268
	}
L268:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v1123)
	F_pfree(m, v1106)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L19
	} else {
		goto L269
	}
L269:
	;
	v1133 = F_heap_form_tuple(m, v1109, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L19
	} else {
		goto L270
	}
L270:
	;
	F_tuplestore_puttuple(m, v1108, v1133)
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L19
	} else {
		goto L271
	}
L271:
	;
	v1137 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v1137)
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[46])))
	v1144 = v1139
	goto L266
L272:
	;
	v1151 = F_pstrdup(m, int32(_a_F_verify_heapam_50))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L19
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v1188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1078)+18)))
	if v1094|base.B2i32(v1188&int32(_a_F_verify_heapam_51) == int32(0)) != 0 {
		v1238 = v1188
		goto L280
	} else {
		goto L281
	}
L275:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v1155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v1156 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v1157 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v1157
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v1156
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v1155))
	v1166 = int32(base.Ui32(v1155) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v1166)
	v1168 = F_cstring_to_text(m, v1151)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L19
	} else {
		goto L276
	}
L276:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v1168)
	F_pfree(m, v1151)
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L19
	} else {
		goto L277
	}
L277:
	;
	v1178 = F_heap_form_tuple(m, v1154, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L19
	} else {
		goto L278
	}
L278:
	;
	F_tuplestore_puttuple(m, v1153, v1178)
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L19
	} else {
		goto L279
	}
L279:
	;
	v1182 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v1182)
	goto L274
L280:
	;
	if int32(0) <= base.I32_extend16_s(v1238) {
		goto L288
	} else {
		goto L289
	}
L281:
	;
	v1194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1078)+20)))
	if v1194&int32(2048)|base.B2i32(v1194&int32(768) == int32(512)) != 0 {
		v1238 = v1188
		goto L280
	} else {
		goto L282
	}
L282:
	;
	v1204 = F_psprintf(m, int32(_a_F_verify_heapam_52), int32(0))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L19
	} else {
		goto L283
	}
L283:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v1208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v1209 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v1210 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v1210
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v1209
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v1208))
	v1219 = int32(base.Ui32(v1208) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v1219)
	v1221 = F_cstring_to_text(m, v1204)
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L19
	} else {
		goto L284
	}
L284:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v1221)
	F_pfree(m, v1204)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L19
	} else {
		goto L285
	}
L285:
	;
	v1231 = F_heap_form_tuple(m, v1207, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L19
	} else {
		goto L286
	}
L286:
	;
	F_tuplestore_puttuple(m, v1206, v1231)
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L19
	} else {
		goto L287
	}
L287:
	;
	v1235 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v1235)
	v1237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1078)+18)))
	v1238 = v1237
	goto L280
L288:
	;
	if v1084&int32(1) == int32(0) {
		goto L299
	} else {
		goto L300
	}
L289:
	;
	v1245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1078)+21)))
	if v1245&int32(32) != 0 {
		goto L288
	} else {
		goto L290
	}
L290:
	;
	v1250 = F_psprintf(m, int32(_a_F_verify_heapam_53), int32(0))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L19
	} else {
		goto L291
	}
L291:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v1254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v1255 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v1256 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v1256
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v1255
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v1254))
	v1265 = int32(base.Ui32(v1254) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v1265)
	v1267 = F_cstring_to_text(m, v1250)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L19
	} else {
		goto L292
	}
L292:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v1267)
	F_pfree(m, v1250)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L19
	} else {
		goto L293
	}
L293:
	;
	v1277 = F_heap_form_tuple(m, v1253, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L19
	} else {
		goto L294
	}
L294:
	;
	F_tuplestore_puttuple(m, v1252, v1277)
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L19
	} else {
		goto L295
	}
L295:
	;
	v1281 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v1281)
	goto L288
L296:
	;
	v2643 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29])))
	v2644 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[46])))
	v2645 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2644)+12)))
	v2648 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2644)+14)))
	if v2643 != v2645<<(uint(int32(16))%32)|v2648 {
		goto L197
	} else {
		goto L574
	}
L297:
	;
	if base.Ui32(v1099) < base.Ui32(v1097) {
		goto L296
	} else {
		goto L330
	}
L298:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[47])))
	if v1386 == int32(1) {
		goto L317
	} else {
		goto L318
	}
L299:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[46])))
	v1292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+22)))
	if v1292 != int32(24) {
		goto L298
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[47])))
	v1299 = base.I32_div_s(v1295+int32(7), int32(8))
	v1303 = (v1299 + int32(30)) & int32(-8)
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[46])))
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304)+22)))
	if v1303 == v1305 {
		v1469 = v1304
		goto L297
	} else {
		goto L303
	}
L302:
	;
	v1469 = v1291
	goto L297
L303:
	;
	if v1295 == int32(1) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+948)) = v1305
	*(*int32)(unsafe.Add(mBase, uint32(v23)+944)) = v1303
	v1314 = F_psprintf(m, int32(_a_F_verify_heapam_54), v23+int32(944))
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L19
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+968)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v23)+964)) = v1305
	*(*int32)(unsafe.Add(mBase, uint32(v23)+960)) = v1303
	v1353 = F_psprintf(m, int32(_a_F_verify_heapam_55), v23+int32(960))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L19
	} else {
		goto L312
	}
L307:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v1318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v1319 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v1320 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v1320
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v1319
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v1318))
	v1329 = int32(base.Ui32(v1318) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v1329)
	v1331 = F_cstring_to_text(m, v1314)
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L19
	} else {
		goto L308
	}
L308:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v1331)
	F_pfree(m, v1314)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L19
	} else {
		goto L309
	}
L309:
	;
	v1341 = F_heap_form_tuple(m, v1317, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L19
	} else {
		goto L310
	}
L310:
	;
	F_tuplestore_puttuple(m, v1316, v1341)
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L19
	} else {
		goto L311
	}
L311:
	;
	v1345 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v1345)
	goto L296
L312:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v1357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v1358 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v1359 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v1359
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v1358
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v1357))
	v1368 = int32(base.Ui32(v1357) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v1368)
	v1370 = F_cstring_to_text(m, v1353)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L19
	} else {
		goto L313
	}
L313:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v1370)
	F_pfree(m, v1353)
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L19
	} else {
		goto L314
	}
L314:
	;
	v1380 = F_heap_form_tuple(m, v1356, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L19
	} else {
		goto L315
	}
L315:
	;
	F_tuplestore_puttuple(m, v1355, v1380)
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L19
	} else {
		goto L316
	}
L316:
	;
	v1384 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v1384)
	goto L296
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+916)) = v1292
	*(*int32)(unsafe.Add(mBase, uint32(v23)+912)) = int32(24)
	v1395 = F_psprintf(m, int32(_a_F_verify_heapam_56), v23+int32(912))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L19
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+936)) = v1386
	*(*int32)(unsafe.Add(mBase, uint32(v23)+932)) = v1292
	*(*int32)(unsafe.Add(mBase, uint32(v23)+928)) = int32(24)
	v1435 = F_psprintf(m, int32(_a_F_verify_heapam_57), v23+int32(928))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L19
	} else {
		goto L325
	}
L320:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v1399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v1400 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v1401 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v1401
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v1400
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v1399))
	v1410 = int32(base.Ui32(v1399) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v1410)
	v1412 = F_cstring_to_text(m, v1395)
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L19
	} else {
		goto L321
	}
L321:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v1412)
	F_pfree(m, v1395)
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L19
	} else {
		goto L322
	}
L322:
	;
	v1422 = F_heap_form_tuple(m, v1398, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L19
	} else {
		goto L323
	}
L323:
	;
	F_tuplestore_puttuple(m, v1397, v1422)
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L19
	} else {
		goto L324
	}
L324:
	;
	v1426 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v1426)
	goto L296
L325:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v1439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v1440 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v1441 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v1441
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v1440
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v1439))
	v1450 = int32(base.Ui32(v1439) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v1450)
	v1452 = F_cstring_to_text(m, v1435)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L19
	} else {
		goto L326
	}
L326:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v1452)
	F_pfree(m, v1435)
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L19
	} else {
		goto L327
	}
L327:
	;
	v1462 = F_heap_form_tuple(m, v1438, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L19
	} else {
		goto L328
	}
L328:
	;
	F_tuplestore_puttuple(m, v1437, v1462)
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L19
	} else {
		goto L329
	}
L329:
	;
	v1466 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v1466)
	goto L296
L330:
	;
	v1473 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[48]))) = uint8(v1473)
	v1475 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v711))) = uint8(v1475)
	v1478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1469)+20)))
	v1479 = int32(768)
	if v1478&v1479 != v1479 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1469)))
	v1484 = v1483
	goto L333
L332:
	;
	v1484 = int32(2)
	goto L333
L333:
	;
	v1489 = F_get_xid_status(m, v1484, v23+int32(_a_F_verify_heapam_2), v23+int32(_a_F_verify_heapam_58))
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L19
	} else {
		goto L340
	}
L334:
	;
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[12])))
	v2084 = *(*int32)(unsafe.Add(mBase, uint32(v2083)+52))
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v2084)))
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[47])))
	if v2085 < v2086 {
		goto L479
	} else {
		goto L480
	}
L335:
	;
	F_ReadMultiXactIdRange(m, v386, v388)
	mBase = m.M
	v2043 = m.ExcPending
	if v2043 != 0 {
		goto L19
	} else {
		goto L473
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+864)) = v1484
	v2001 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[23])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+872)) = uint32(v2001)
	v2004 = int64(base.Ui64(v2001) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+868)) = uint32(v2004)
	v2009 = F_psprintf(m, int32(_a_F_verify_heapam_59), v23+int32(864))
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L19
	} else {
		goto L468
	}
L337:
	;
	v1577 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v711))) = uint8(v1577)
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[49])))
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(1040)+v725))) = v1582
	v1584 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1469)+20)))
	if v1584&int32(256) != 0 {
		goto L351
	} else {
		goto L352
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+896)) = v1484
	v1536 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[6])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+904)) = uint32(v1536)
	v1539 = int64(base.Ui64(v1536) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+900)) = uint32(v1539)
	v1544 = F_psprintf(m, int32(_a_F_verify_heapam_60), v23+int32(896))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L19
	} else {
		goto L346
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+880)) = v1484
	v1494 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[26])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+888)) = uint32(v1494)
	v1497 = int64(base.Ui64(v1494) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+884)) = uint32(v1497)
	v1502 = F_psprintf(m, int32(_a_F_verify_heapam_61), v23+int32(880))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L19
	} else {
		goto L341
	}
L340:
	;
	switch v1489 - int32(1) {
	case 0:
		goto L336
	case 1:
		goto L339
	case 2:
		goto L338
	case 3:
		goto L337
	default:
		goto L296
	}
L341:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v1506 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v1507 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v1508 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v1508
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v1507
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v1506))
	v1517 = int32(base.Ui32(v1506) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v1517)
	v1519 = F_cstring_to_text(m, v1502)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L19
	} else {
		goto L342
	}
L342:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v1519)
	F_pfree(m, v1502)
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		goto L19
	} else {
		goto L343
	}
L343:
	;
	v1529 = F_heap_form_tuple(m, v1505, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L19
	} else {
		goto L344
	}
L344:
	;
	F_tuplestore_puttuple(m, v1504, v1529)
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L19
	} else {
		goto L345
	}
L345:
	;
	v1533 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v1533)
	goto L296
L346:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v1548 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v1549 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v1550 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v1550
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v1549
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v1548))
	v1559 = int32(base.Ui32(v1548) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v1559)
	v1561 = F_cstring_to_text(m, v1544)
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L19
	} else {
		goto L347
	}
L347:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v1561)
	F_pfree(m, v1544)
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L19
	} else {
		goto L348
	}
L348:
	;
	v1571 = F_heap_form_tuple(m, v1547, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L19
	} else {
		goto L349
	}
L349:
	;
	F_tuplestore_puttuple(m, v1546, v1571)
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L19
	} else {
		goto L350
	}
L350:
	;
	v1575 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v1575)
	goto L296
L351:
	;
	v1757 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1469)+20)))
	if v1757&int32(_a_F_verify_heapam_27) == int32(0) {
		v1792 = v1757
		goto L405
	} else {
		goto L406
	}
L352:
	;
	v1587 = base.I32_extend16_s(v1584)
	if v1587&int32(512) != 0 {
		goto L296
	} else {
		goto L353
	}
L353:
	;
	if v1587&int32(_a_F_verify_heapam_51) != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+8))
	v1597 = F_get_xid_status(m, v1592, v23+int32(_a_F_verify_heapam_2), v23+int32(_a_F_verify_heapam_62))
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L19
	} else {
		goto L362
	}
L355:
	;
	goto L356
L356:
	;
	if v1587 < int32(0) {
		goto L377
	} else {
		goto L378
	}
L357:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[50])))
	switch v1651 {
	case 0:
		goto L334
	case 1:
		goto L372
	case 2:
		goto L371
	default:
		goto L351
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+848)) = v1592
	v1637 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[26])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+856)) = uint32(v1637)
	v1640 = int64(base.Ui64(v1637) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+852)) = uint32(v1640)
	v1647 = F_psprintf(m, int32(_a_F_verify_heapam_63), v23+int32(848))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L19
	} else {
		goto L369
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+832)) = v1592
	v1622 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[6])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+840)) = uint32(v1622)
	v1625 = int64(base.Ui64(v1622) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+836)) = uint32(v1625)
	v1632 = F_psprintf(m, int32(_a_F_verify_heapam_64), v23+int32(832))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L19
	} else {
		goto L367
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+816)) = v1592
	v1607 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[23])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+824)) = uint32(v1607)
	v1610 = int64(base.Ui64(v1607) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+820)) = uint32(v1610)
	v1617 = F_psprintf(m, int32(_a_F_verify_heapam_65), v23+int32(816))
	mBase = m.M
	v1618 = m.ExcPending
	if v1618 != 0 {
		goto L19
	} else {
		goto L365
	}
L361:
	;
	v1602 = F_pstrdup(m, int32(_a_F_verify_heapam_66))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L19
	} else {
		goto L363
	}
L362:
	;
	switch v1597 {
	case 0:
		goto L361
	case 1:
		goto L360
	case 2:
		goto L358
	case 3:
		goto L359
	default:
		goto L357
	}
L363:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1602)
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L19
	} else {
		goto L364
	}
L364:
	;
	goto L296
L365:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1617)
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L19
	} else {
		goto L366
	}
L366:
	;
	goto L296
L367:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1632)
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L19
	} else {
		goto L368
	}
L368:
	;
	goto L296
L369:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1647)
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L19
	} else {
		goto L370
	}
L370:
	;
	goto L296
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+800)) = v1592
	v1668 = F_psprintf(m, int32(_a_F_verify_heapam_67), v23+int32(800))
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L19
	} else {
		goto L375
	}
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+784)) = v1592
	v1658 = F_psprintf(m, int32(_a_F_verify_heapam_68), v23+int32(784))
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L19
	} else {
		goto L373
	}
L373:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1658)
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L19
	} else {
		goto L374
	}
L374:
	;
	goto L296
L375:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1668)
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L19
	} else {
		goto L376
	}
L376:
	;
	goto L296
L377:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+8))
	v1679 = F_get_xid_status(m, v1674, v23+int32(_a_F_verify_heapam_2), v23+int32(_a_F_verify_heapam_62))
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L19
	} else {
		goto L385
	}
L378:
	;
	goto L379
L379:
	;
	if v1582 != 0 {
		goto L296
	} else {
		goto L400
	}
L380:
	;
	v1733 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[50])))
	switch v1733 - int32(1) {
	case 0:
		goto L395
	case 1:
		goto L394
	case 2:
		goto L334
	default:
		goto L351
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+768)) = v1674
	v1719 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[26])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+776)) = uint32(v1719)
	v1722 = int64(base.Ui64(v1719) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+772)) = uint32(v1722)
	v1729 = F_psprintf(m, int32(_a_F_verify_heapam_69), v23+int32(768))
	mBase = m.M
	v1730 = m.ExcPending
	if v1730 != 0 {
		goto L19
	} else {
		goto L392
	}
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+752)) = v1674
	v1704 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[6])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+760)) = uint32(v1704)
	v1707 = int64(base.Ui64(v1704) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+756)) = uint32(v1707)
	v1714 = F_psprintf(m, int32(_a_F_verify_heapam_70), v23+int32(752))
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L19
	} else {
		goto L390
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+736)) = v1674
	v1689 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[23])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+744)) = uint32(v1689)
	v1692 = int64(base.Ui64(v1689) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+740)) = uint32(v1692)
	v1699 = F_psprintf(m, int32(_a_F_verify_heapam_71), v23+int32(736))
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		goto L19
	} else {
		goto L388
	}
L384:
	;
	v1684 = F_pstrdup(m, int32(_a_F_verify_heapam_72))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L19
	} else {
		goto L386
	}
L385:
	;
	switch v1679 {
	case 0:
		goto L384
	case 1:
		goto L383
	case 2:
		goto L381
	case 3:
		goto L382
	default:
		goto L380
	}
L386:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1684)
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L19
	} else {
		goto L387
	}
L387:
	;
	goto L296
L388:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1699)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L19
	} else {
		goto L389
	}
L389:
	;
	goto L296
L390:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1714)
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L19
	} else {
		goto L391
	}
L391:
	;
	goto L296
L392:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1729)
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L19
	} else {
		goto L393
	}
L393:
	;
	goto L296
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+720)) = v1674
	v1752 = F_psprintf(m, int32(_a_F_verify_heapam_73), v23+int32(720))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L19
	} else {
		goto L398
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+704)) = v1674
	v1742 = F_psprintf(m, int32(_a_F_verify_heapam_74), v23+int32(704))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L19
	} else {
		goto L396
	}
L396:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1742)
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L19
	} else {
		goto L397
	}
L397:
	;
	goto L296
L398:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1752)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L19
	} else {
		goto L399
	}
L399:
	;
	goto L296
L400:
	;
	goto L351
L401:
	;
	v1835 = int32(0)
	if base.B2i32(v1792&int32(128) == v1835)&base.B2i32(v1792&int32(_a_F_verify_heapam_75) != int32(64)) == v1835 {
		goto L423
	} else {
		goto L424
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+692)) = v1787
	*(*int32)(unsafe.Add(mBase, uint32(v23)+688)) = v1762
	v1829 = F_psprintf(m, int32(_a_F_verify_heapam_76), v23+int32(688))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L19
	} else {
		goto L421
	}
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+676)) = v1783
	*(*int32)(unsafe.Add(mBase, uint32(v23)+672)) = v1762
	v1818 = F_psprintf(m, int32(_a_F_verify_heapam_77), v23+int32(672))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L19
	} else {
		goto L419
	}
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+660)) = v1779
	*(*int32)(unsafe.Add(mBase, uint32(v23)+656)) = v1762
	v1807 = F_psprintf(m, int32(_a_F_verify_heapam_78), v23+int32(656))
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L19
	} else {
		goto L417
	}
L405:
	;
	if v1792&int32(2048) == int32(0) {
		goto L401
	} else {
		goto L416
	}
L406:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+4))
	if v1762 == int32(0) {
		goto L335
	} else {
		goto L407
	}
L407:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[7])))
	if v1762-v1765 < int32(0) {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	F_ReadMultiXactIdRange(m, v386, v388)
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L19
	} else {
		goto L412
	}
L409:
	;
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[51])))
	if v1762-v1769 < int32(0) {
		goto L408
	} else {
		goto L410
	}
L410:
	;
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[52])))
	if int32(0) < v1773-v1762 {
		v1792 = v1757
		goto L405
	} else {
		goto L411
	}
L411:
	;
	goto L408
L412:
	;
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[7])))
	if v1762-v1779 < int32(0) {
		goto L404
	} else {
		goto L413
	}
L413:
	;
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[51])))
	if v1762-v1783 < int32(0) {
		goto L403
	} else {
		goto L414
	}
L414:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[52])))
	if v1787-v1762 <= int32(0) {
		goto L402
	} else {
		goto L415
	}
L415:
	;
	v1791 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1469)+20)))
	v1792 = v1791
	goto L405
L416:
	;
	v1798 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[48]))) = uint8(v1798)
	goto L334
L417:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1807)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L19
	} else {
		goto L418
	}
L418:
	;
	goto L334
L419:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1818)
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		goto L19
	} else {
		goto L420
	}
L420:
	;
	goto L334
L421:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1829)
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L19
	} else {
		goto L422
	}
L422:
	;
	goto L334
L423:
	;
	v1844 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[48]))) = uint8(v1844)
	goto L334
L424:
	;
	goto L425
L425:
	;
	if v1792&int32(_a_F_verify_heapam_27) != 0 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v1848 = F_HeapTupleGetUpdateXid(m, v1469)
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L19
	} else {
		goto L434
	}
L427:
	;
	goto L428
L428:
	;
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+4))
	v1932 = F_get_xid_status(m, v1927, v23+int32(_a_F_verify_heapam_2), v23+int32(_a_F_verify_heapam_79))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L19
	} else {
		goto L455
	}
L429:
	;
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[53])))
	switch v1908 {
	case 0:
		goto L445
	case 1, 2:
		goto L446
	case 3:
		goto L444
	default:
		goto L334
	}
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+640)) = v1848
	v1894 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[26])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+648)) = uint32(v1894)
	v1897 = int64(base.Ui64(v1894) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+644)) = uint32(v1897)
	v1904 = F_psprintf(m, int32(_a_F_verify_heapam_80), v23+int32(640))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L19
	} else {
		goto L442
	}
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+624)) = v1848
	v1879 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[6])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+632)) = uint32(v1879)
	v1882 = int64(base.Ui64(v1879) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+628)) = uint32(v1882)
	v1889 = F_psprintf(m, int32(_a_F_verify_heapam_81), v23+int32(624))
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L19
	} else {
		goto L440
	}
L432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+608)) = v1848
	v1864 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[23])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+616)) = uint32(v1864)
	v1867 = int64(base.Ui64(v1864) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+612)) = uint32(v1867)
	v1874 = F_psprintf(m, int32(_a_F_verify_heapam_82), v23+int32(608))
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
		goto L19
	} else {
		goto L438
	}
L433:
	;
	v1859 = F_pstrdup(m, int32(_a_F_verify_heapam_83))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L19
	} else {
		goto L436
	}
L434:
	;
	v1854 = F_get_xid_status(m, v1848, v23+int32(_a_F_verify_heapam_2), v23+int32(_a_F_verify_heapam_79))
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L19
	} else {
		goto L435
	}
L435:
	;
	switch v1854 {
	case 0:
		goto L433
	case 1:
		goto L432
	case 2:
		goto L430
	case 3:
		goto L431
	default:
		goto L429
	}
L436:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1859)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L19
	} else {
		goto L437
	}
L437:
	;
	goto L334
L438:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1874)
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L19
	} else {
		goto L439
	}
L439:
	;
	goto L334
L440:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1889)
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L19
	} else {
		goto L441
	}
L441:
	;
	goto L334
L442:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1904)
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L19
	} else {
		goto L443
	}
L443:
	;
	goto L334
L444:
	;
	v1925 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[48]))) = uint8(v1925)
	goto L334
L445:
	;
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[9])))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1911))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1848)) == int32(0) {
		goto L447
	} else {
		goto L448
	}
L446:
	;
	v1909 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[48]))) = uint8(v1909)
	goto L334
L447:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[48]))) = uint8(base.B2i32(base.Ui32(v1848) < base.Ui32(v1911)))
	goto L334
L448:
	;
	goto L449
L449:
	;
	v1923 = int32(base.Ui32(v1848-v1911) >> (uint(int32(31)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[48]))) = uint8(v1923)
	goto L334
L450:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[53])))
	switch v1981 {
	case 0:
		goto L463
	case 1, 2:
		goto L464
	case 3:
		goto L462
	default:
		goto L334
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+592)) = v1927
	v1967 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[26])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+600)) = uint32(v1967)
	v1970 = int64(base.Ui64(v1967) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+596)) = uint32(v1970)
	v1977 = F_psprintf(m, int32(_a_F_verify_heapam_84), v23+int32(592))
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L19
	} else {
		goto L460
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+576)) = v1927
	v1952 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[6])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+584)) = uint32(v1952)
	v1955 = int64(base.Ui64(v1952) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+580)) = uint32(v1955)
	v1962 = F_psprintf(m, int32(_a_F_verify_heapam_85), v23+int32(576))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L19
	} else {
		goto L458
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+560)) = v1927
	v1937 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[23])))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+568)) = uint32(v1937)
	v1940 = int64(base.Ui64(v1937) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+564)) = uint32(v1940)
	v1947 = F_psprintf(m, int32(_a_F_verify_heapam_86), v23+int32(560))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L19
	} else {
		goto L456
	}
L454:
	;
	v1934 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[48]))) = uint8(v1934)
	goto L334
L455:
	;
	switch v1932 {
	case 0:
		goto L454
	case 1:
		goto L453
	case 2:
		goto L451
	case 3:
		goto L452
	default:
		goto L450
	}
L456:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1947)
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L19
	} else {
		goto L457
	}
L457:
	;
	goto L296
L458:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1962)
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L19
	} else {
		goto L459
	}
L459:
	;
	goto L296
L460:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v1977)
	mBase = m.M
	v1980 = m.ExcPending
	if v1980 != 0 {
		goto L19
	} else {
		goto L461
	}
L461:
	;
	goto L296
L462:
	;
	v1998 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[48]))) = uint8(v1998)
	goto L334
L463:
	;
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[9])))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1984))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1927)) == int32(0) {
		goto L465
	} else {
		goto L466
	}
L464:
	;
	v1982 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[48]))) = uint8(v1982)
	goto L334
L465:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[48]))) = uint8(base.B2i32(base.Ui32(v1927) < base.Ui32(v1984)))
	goto L334
L466:
	;
	goto L467
L467:
	;
	v1996 = int32(base.Ui32(v1927-v1984) >> (uint(int32(31)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[48]))) = uint8(v1996)
	goto L334
L468:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v2013 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2014 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v2015 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v2015
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v2014
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v2013))
	v2024 = int32(base.Ui32(v2013) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v2024)
	v2026 = F_cstring_to_text(m, v2009)
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L19
	} else {
		goto L469
	}
L469:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v2026)
	F_pfree(m, v2009)
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L19
	} else {
		goto L470
	}
L470:
	;
	v2036 = F_heap_form_tuple(m, v2012, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L19
	} else {
		goto L471
	}
L471:
	;
	F_tuplestore_puttuple(m, v2011, v2036)
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L19
	} else {
		goto L472
	}
L472:
	;
	v2040 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v2040)
	goto L296
L473:
	;
	v2045 = F_pstrdup(m, int32(_a_F_verify_heapam_87))
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L19
	} else {
		goto L474
	}
L474:
	;
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v2049 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2050 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v2051 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v2051
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v2050
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v2049))
	v2060 = int32(base.Ui32(v2049) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v2060)
	v2062 = F_cstring_to_text(m, v2045)
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L19
	} else {
		goto L475
	}
L475:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v2062)
	F_pfree(m, v2045)
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L19
	} else {
		goto L476
	}
L476:
	;
	v2072 = F_heap_form_tuple(m, v2048, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2073 = m.ExcPending
	if v2073 != 0 {
		goto L19
	} else {
		goto L477
	}
L477:
	;
	F_tuplestore_puttuple(m, v2047, v2072)
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L19
	} else {
		goto L478
	}
L478:
	;
	v2076 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v2076)
	goto L334
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+404)) = v2085
	*(*int32)(unsafe.Add(mBase, uint32(v23)+400)) = v2086
	v2093 = F_psprintf(m, int32(_a_F_verify_heapam_88), v23+int32(400))
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L19
	} else {
		goto L482
	}
L480:
	;
	goto L481
L481:
	;
	v2126 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))) = uint16(v2126)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[54]))) = v2126
	if v2086 <= v2126 {
		goto L487
	} else {
		goto L488
	}
L482:
	;
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v2097 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2098 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v2099 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v2099
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v2098
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v2097))
	v2108 = int32(base.Ui32(v2097) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v2108)
	v2110 = F_cstring_to_text(m, v2093)
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L19
	} else {
		goto L483
	}
L483:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v2110)
	F_pfree(m, v2093)
	mBase = m.M
	v2115 = m.ExcPending
	if v2115 != 0 {
		goto L19
	} else {
		goto L484
	}
L484:
	;
	v2120 = F_heap_form_tuple(m, v2096, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2121 = m.ExcPending
	if v2121 != 0 {
		goto L19
	} else {
		goto L485
	}
L485:
	;
	F_tuplestore_puttuple(m, v2095, v2120)
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L19
	} else {
		goto L486
	}
L486:
	;
	v2124 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v2124)
	goto L296
L487:
	;
	v2621 = int32(_a_F_verify_heapam_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))) = uint16(v2621)
	goto L296
L488:
	;
	v2135 = v2126
	goto L489
L489:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[12])))
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v2153)+52))
	v2159 = v2154 + v2135<<(uint(int32(3))%32) + int32(28)
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[54])))
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[46])))
	v2162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2161)+22)))
	v2163 = v2160 + v2162
	v2164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[37]))))
	if base.Ui32(v2164) < base.Ui32(v2163) {
		goto L491
	} else {
		goto L492
	}
L490:
	;
	goto L487
L491:
	;
	v2166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2159)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+424)) = v2164
	*(*int32)(unsafe.Add(mBase, uint32(v23)+420)) = v2163
	*(*int32)(unsafe.Add(mBase, uint32(v23)+416)) = v2166
	v2173 = F_psprintf(m, int32(_a_F_verify_heapam_89), v23+int32(416))
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L19
	} else {
		goto L494
	}
L492:
	;
	goto L493
L493:
	;
	v2206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2161)+20)))
	if v2206&int32(1) != 0 {
		goto L500
	} else {
		goto L501
	}
L494:
	;
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v2177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2178 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v2179 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v2179
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v2178
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v2177))
	v2188 = int32(base.Ui32(v2177) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v2188)
	v2190 = F_cstring_to_text(m, v2173)
	mBase = m.M
	v2191 = m.ExcPending
	if v2191 != 0 {
		goto L19
	} else {
		goto L495
	}
L495:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v2190)
	F_pfree(m, v2173)
	mBase = m.M
	v2195 = m.ExcPending
	if v2195 != 0 {
		goto L19
	} else {
		goto L496
	}
L496:
	;
	v2200 = F_heap_form_tuple(m, v2176, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L19
	} else {
		goto L497
	}
L497:
	;
	F_tuplestore_puttuple(m, v2175, v2200)
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L19
	} else {
		goto L498
	}
L498:
	;
	v2204 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v2204)
	goto L487
L499:
	;
	v2594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2596 = v2594 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))) = uint16(v2596)
	v2598 = base.I32_extend16_s(v2596)
	v2599 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[47])))
	if v2598 < v2599 {
		v2135 = v2598
		goto L489
	} else {
		goto L573
	}
L500:
	;
	v2212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2161+v2135>>(uint(int32(3))%32))+23)))
	if int32(base.Ui32(v2212)>>(uint(v2135&int32(7))%32))&int32(1) == int32(0) {
		goto L499
	} else {
		goto L503
	}
L501:
	;
	goto L502
L502:
	;
	v2220 = v2161 + v2162
	v2221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2159)+2)))
	if v2221 != int32(-1) {
		goto L504
	} else {
		goto L505
	}
L503:
	;
	goto L502
L504:
	;
	v2224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2159)+5)))
	v2228 = int32(0)
	v2230 = (v2160 + v2224 - int32(1)) & (v2228 - v2224)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[54]))) = v2230
	if v2221 <= v2228 {
		goto L507
	} else {
		goto L508
	}
L505:
	;
	goto L506
L506:
	;
	v2284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2160+v2220))))
	if v2284 == int32(0) {
		goto L516
	} else {
		goto L517
	}
L507:
	;
	v2235 = F_strlen(m, v2230+v2220)
	mBase = m.M
	v2238 = v2235 + int32(1)
	goto L509
L508:
	;
	v2238 = v2221
	goto L509
L509:
	;
	v2239 = v2238 + v2230
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[54]))) = v2239
	v2241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2161)+22)))
	v2242 = v2239 + v2241
	if base.Ui32(v2242) <= base.Ui32(v2164) {
		goto L499
	} else {
		goto L510
	}
L510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+552)) = v2164
	*(*int32)(unsafe.Add(mBase, uint32(v23)+548)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v23)+544)) = v2221
	v2250 = F_psprintf(m, int32(_a_F_verify_heapam_90), v23+int32(544))
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L19
	} else {
		goto L511
	}
L511:
	;
	v2252 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v2253 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v2254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2255 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v2256 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v2256
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v2255
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v2254))
	v2265 = int32(base.Ui32(v2254) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v2265)
	v2267 = F_cstring_to_text(m, v2250)
	mBase = m.M
	v2268 = m.ExcPending
	if v2268 != 0 {
		goto L19
	} else {
		goto L512
	}
L512:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v2267)
	F_pfree(m, v2250)
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L19
	} else {
		goto L513
	}
L513:
	;
	v2277 = F_heap_form_tuple(m, v2253, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2278 = m.ExcPending
	if v2278 != 0 {
		goto L19
	} else {
		goto L514
	}
L514:
	;
	F_tuplestore_puttuple(m, v2252, v2277)
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L19
	} else {
		goto L515
	}
L515:
	;
	v2281 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v2281)
	goto L487
L516:
	;
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2159)+5)))
	v2295 = (v2160 + v2287 - int32(1)) & (int32(0) - v2287)
	goto L518
L517:
	;
	v2295 = v2160
	goto L518
L518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[54]))) = v2295
	v2297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2159)+4)))
	if v2297 == int32(1) {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	v2300 = v2295 + v2220
	v2301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2300))))
	if v2301 == int32(1) {
		goto L521
	} else {
		goto L522
	}
L520:
	;
	v2354 = v2353 + v2295
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[54]))) = v2354
	v2356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2161)+22)))
	v2357 = v2354 + v2356
	if base.Ui32(v2164) < base.Ui32(v2357) {
		goto L531
	} else {
		goto L532
	}
L521:
	;
	v2304 = int32(18)
	v2305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2300)+1)))
	if v2305 == v2304 {
		v2353 = v2304
		goto L520
	} else {
		goto L524
	}
L522:
	;
	goto L523
L523:
	;
	v2345 = int32(1)
	if v2301&v2345 != 0 {
		v2353 = int32(base.Ui32(v2301) >> (uint(v2345) % 32))
		goto L520
	} else {
		goto L530
	}
L524:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+528)) = v2305
	v2312 = F_psprintf(m, int32(_a_F_verify_heapam_91), v23+int32(528))
	mBase = m.M
	v2313 = m.ExcPending
	if v2313 != 0 {
		goto L19
	} else {
		goto L525
	}
L525:
	;
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v2316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2317 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v2318 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v2318
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v2317
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v2316))
	v2327 = int32(base.Ui32(v2316) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v2327)
	v2329 = F_cstring_to_text(m, v2312)
	mBase = m.M
	v2330 = m.ExcPending
	if v2330 != 0 {
		goto L19
	} else {
		goto L526
	}
L526:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v2329)
	F_pfree(m, v2312)
	mBase = m.M
	v2334 = m.ExcPending
	if v2334 != 0 {
		goto L19
	} else {
		goto L527
	}
L527:
	;
	v2339 = F_heap_form_tuple(m, v2315, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		goto L19
	} else {
		goto L528
	}
L528:
	;
	F_tuplestore_puttuple(m, v2314, v2339)
	mBase = m.M
	v2342 = m.ExcPending
	if v2342 != 0 {
		goto L19
	} else {
		goto L529
	}
L529:
	;
	v2343 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v2343)
	goto L487
L530:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, uint32(v2300)))
	v2353 = int32(base.Ui32(v2349) >> (uint(int32(2)) % 32))
	goto L520
L531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v2164
	*(*int32)(unsafe.Add(mBase, uint32(v23)+452)) = v2357
	*(*int32)(unsafe.Add(mBase, uint32(v23)+448)) = int32(-1)
	v2366 = F_psprintf(m, int32(_a_F_verify_heapam_90), v23+int32(448))
	mBase = m.M
	v2367 = m.ExcPending
	if v2367 != 0 {
		goto L19
	} else {
		goto L534
	}
L532:
	;
	goto L533
L533:
	;
	v2399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2300))))
	if v2399 != int32(1) {
		goto L499
	} else {
		goto L539
	}
L534:
	;
	v2368 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v2369 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v2370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2371 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v2372 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v2372
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v2371
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v2370))
	v2381 = int32(base.Ui32(v2370) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v2381)
	v2383 = F_cstring_to_text(m, v2366)
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
		goto L19
	} else {
		goto L535
	}
L535:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v2383)
	F_pfree(m, v2366)
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L19
	} else {
		goto L536
	}
L536:
	;
	v2393 = F_heap_form_tuple(m, v2369, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L19
	} else {
		goto L537
	}
L537:
	;
	F_tuplestore_puttuple(m, v2368, v2393)
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L19
	} else {
		goto L538
	}
L538:
	;
	v2397 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v2397)
	goto L487
L539:
	;
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v2300)+10))
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v2300)+6))
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(v2300)+2))
	if int32(1073741824) <= v2404 {
		goto L540
	} else {
		goto L541
	}
L540:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+520)) = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+516)) = v2404
	*(*int32)(unsafe.Add(mBase, uint32(v23)+512)) = v2402
	v2414 = F_psprintf(m, int32(_a_F_verify_heapam_92), v23+int32(512))
	mBase = m.M
	v2415 = m.ExcPending
	if v2415 != 0 {
		goto L19
	} else {
		goto L543
	}
L541:
	;
	goto L542
L542:
	;
	v2456 = int32(0)
	if base.B2i32(base.Ui32(v2404-int32(4)) <= base.Ui32(v2403&int32(1073741823)))|base.B2i32(v2456 <= v2403) == v2456 {
		goto L548
	} else {
		goto L549
	}
L543:
	;
	v2416 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v2417 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v2418 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2419 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v2420 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v2420
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v2419
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v2418))
	v2429 = int32(base.Ui32(v2418) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v2429)
	v2431 = F_cstring_to_text(m, v2414)
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L19
	} else {
		goto L544
	}
L544:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v2431)
	F_pfree(m, v2414)
	mBase = m.M
	v2436 = m.ExcPending
	if v2436 != 0 {
		goto L19
	} else {
		goto L545
	}
L545:
	;
	v2441 = F_heap_form_tuple(m, v2417, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2442 = m.ExcPending
	if v2442 != 0 {
		goto L19
	} else {
		goto L546
	}
L546:
	;
	F_tuplestore_puttuple(m, v2416, v2441)
	mBase = m.M
	v2444 = m.ExcPending
	if v2444 != 0 {
		goto L19
	} else {
		goto L547
	}
L547:
	;
	v2445 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v2445)
	goto L542
L548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+496)) = v2402
	*(*int32)(unsafe.Add(mBase, uint32(v23)+500)) = int32(base.Ui32(v2403) >> (uint(int32(30)) % 32))
	v2468 = F_psprintf(m, int32(_a_F_verify_heapam_93), v23+int32(496))
	mBase = m.M
	v2469 = m.ExcPending
	if v2469 != 0 {
		goto L19
	} else {
		goto L551
	}
L549:
	;
	goto L550
L550:
	;
	if v2206&int32(4) == int32(0) {
		goto L556
	} else {
		goto L557
	}
L551:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v2472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2473 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v2474 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v2474
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v2473
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v2472))
	v2483 = int32(base.Ui32(v2472) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v2483)
	v2485 = F_cstring_to_text(m, v2468)
	mBase = m.M
	v2486 = m.ExcPending
	if v2486 != 0 {
		goto L19
	} else {
		goto L552
	}
L552:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v2485)
	F_pfree(m, v2468)
	mBase = m.M
	v2490 = m.ExcPending
	if v2490 != 0 {
		goto L19
	} else {
		goto L553
	}
L553:
	;
	v2495 = F_heap_form_tuple(m, v2471, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2496 = m.ExcPending
	if v2496 != 0 {
		goto L19
	} else {
		goto L554
	}
L554:
	;
	F_tuplestore_puttuple(m, v2470, v2495)
	mBase = m.M
	v2498 = m.ExcPending
	if v2498 != 0 {
		goto L19
	} else {
		goto L555
	}
L555:
	;
	v2499 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v2499)
	goto L550
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v2402
	v2513 = F_psprintf(m, int32(_a_F_verify_heapam_94), v23+int32(464))
	mBase = m.M
	v2514 = m.ExcPending
	if v2514 != 0 {
		goto L19
	} else {
		goto L559
	}
L557:
	;
	goto L558
L558:
	;
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[12])))
	v2547 = *(*int32)(unsafe.Add(mBase, uint32(v2546)+48))
	v2548 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+112))
	if v2548 == int32(0) {
		goto L564
	} else {
		goto L565
	}
L559:
	;
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v2517 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2518 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v2519 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v2519
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v2518
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v2517))
	v2528 = int32(base.Ui32(v2517) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v2528)
	v2530 = F_cstring_to_text(m, v2513)
	mBase = m.M
	v2531 = m.ExcPending
	if v2531 != 0 {
		goto L19
	} else {
		goto L560
	}
L560:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v2530)
	F_pfree(m, v2513)
	mBase = m.M
	v2535 = m.ExcPending
	if v2535 != 0 {
		goto L19
	} else {
		goto L561
	}
L561:
	;
	v2540 = F_heap_form_tuple(m, v2516, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2541 = m.ExcPending
	if v2541 != 0 {
		goto L19
	} else {
		goto L562
	}
L562:
	;
	F_tuplestore_puttuple(m, v2515, v2540)
	mBase = m.M
	v2543 = m.ExcPending
	if v2543 != 0 {
		goto L19
	} else {
		goto L563
	}
L563:
	;
	v2544 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v2544)
	goto L499
L564:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v2402
	v2557 = F_psprintf(m, int32(_a_F_verify_heapam_95), v23+int32(480))
	mBase = m.M
	v2558 = m.ExcPending
	if v2558 != 0 {
		goto L19
	} else {
		goto L567
	}
L565:
	;
	goto L566
L566:
	;
	v2561 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[19])))
	if v2561 == int32(0) {
		goto L499
	} else {
		goto L569
	}
L567:
	;
	F_report_corruption(m, v23+int32(_a_F_verify_heapam_2), v2557)
	mBase = m.M
	v2560 = m.ExcPending
	if v2560 != 0 {
		goto L19
	} else {
		goto L568
	}
L568:
	;
	goto L499
L569:
	;
	v2564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[48]))))
	if v2564&int32(1) != 0 {
		goto L499
	} else {
		goto L570
	}
L570:
	;
	v2568 = F_palloc0(m, int32(24))
	mBase = m.M
	v2569 = m.ExcPending
	if v2569 != 0 {
		goto L19
	} else {
		goto L571
	}
L571:
	;
	v2571 = v2300 + int32(2)
	v2572 = *(*int64)(unsafe.Add(mBase, uint32(v2571)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2568)+8)) = v2572
	v2574 = *(*int64)(unsafe.Add(mBase, uint32(v2571)))
	*(*int64)(unsafe.Add(mBase, uint32(v2568))) = v2574
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29])))
	*(*int32)(unsafe.Add(mBase, uint32(v2568)+16)) = v2576
	v2578 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2568)+20)) = uint16(v2578)
	v2580 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2568)+22)) = uint16(v2580)
	v2582 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[55])))
	v2583 = F_lappend(m, v2582, v2568)
	mBase = m.M
	v2584 = m.ExcPending
	if v2584 != 0 {
		goto L19
	} else {
		goto L572
	}
L572:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[55]))) = v2583
	goto L499
L573:
	;
	goto L490
L574:
	;
	v2651 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2644)+16)))
	if base.Ui32(v675) <= base.Ui32((v2651-int32(1))&int32(_a_F_verify_heapam_6)) {
		goto L197
	} else {
		goto L575
	}
L575:
	;
	v2657 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	if v2657 == v2651 {
		goto L197
	} else {
		goto L576
	}
L576:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(_a_F_verify_heapam_31)+v2657<<(uint(int32(1))%32)))) = uint16(v2651)
	goto L197
L577:
	;
	goto L196
L578:
	;
	v2720 = v2700 & int32(_a_F_verify_heapam_6)
	v2724 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(_a_F_verify_heapam_31)+v2720<<(uint(int32(1))%32)))))
	if v2724 == int32(0) {
		goto L580
	} else {
		goto L581
	}
L579:
	;
	v3132 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))) = uint16(v3132)
	v3138 = v3132
	v3139 = v3132
	goto L657
L580:
	;
	v3125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	v3127 = v3125 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))) = uint16(v3127)
	if base.Ui32(v3127&int32(_a_F_verify_heapam_6)) <= base.Ui32(v675) {
		v2700 = v3127
		goto L578
	} else {
		goto L656
	}
L581:
	;
	v2730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(_a_F_verify_heapam_29)+v2724))))
	if v2730 != int32(1) {
		goto L580
	} else {
		goto L582
	}
L582:
	;
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[32])))
	v2735 = v2733 + int32(20)
	v2736 = int32(2)
	v2737 = v2724 << (uint(v2736) % 32)
	v2738 = v2735 + v2737
	v2739 = *(*int32)(unsafe.Add(mBase, uint32(v2738)))
	v2743 = *(*int32)(unsafe.Add(mBase, uint32(v2735+v2720<<(uint(v2736)%32))))
	if v2743&int32(_a_F_verify_heapam_96) == int32(_a_F_verify_heapam_97) {
		goto L584
	} else {
		goto L585
	}
L583:
	;
	v3085 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v3086 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v3087 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v3088 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v3089 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v3089
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v3088
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v3087))
	v3098 = int32(base.Ui32(v3087) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v3098)
	v3100 = F_cstring_to_text(m, v3077)
	mBase = m.M
	v3101 = m.ExcPending
	if v3101 != 0 {
		goto L19
	} else {
		goto L652
	}
L584:
	;
	v2751 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2733+v2739&int32(_a_F_verify_heapam_32))+18)))
	if int32(0) <= v2751 {
		goto L587
	} else {
		goto L588
	}
L585:
	;
	goto L586
L586:
	;
	if v2739&int32(_a_F_verify_heapam_96) == int32(_a_F_verify_heapam_97) {
		goto L580
	} else {
		goto L599
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v2724
	v2758 = F_psprintf(m, int32(_a_F_verify_heapam_98), v23+int32(176))
	mBase = m.M
	v2759 = m.ExcPending
	if v2759 != 0 {
		goto L19
	} else {
		goto L590
	}
L588:
	;
	goto L589
L589:
	;
	v2799 = v23 + int32(_a_F_verify_heapam_26) + v2724<<(uint(int32(1))%32)
	v2800 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2799))))
	if v2800 != 0 {
		goto L595
	} else {
		goto L596
	}
L590:
	;
	v2760 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v2761 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v2762 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2763 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v2764 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v2764
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v2763
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v2762))
	v2773 = int32(base.Ui32(v2762) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v2773)
	v2775 = F_cstring_to_text(m, v2758)
	mBase = m.M
	v2776 = m.ExcPending
	if v2776 != 0 {
		goto L19
	} else {
		goto L591
	}
L591:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v2775)
	F_pfree(m, v2758)
	mBase = m.M
	v2780 = m.ExcPending
	if v2780 != 0 {
		goto L19
	} else {
		goto L592
	}
L592:
	;
	v2785 = F_heap_form_tuple(m, v2761, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2786 = m.ExcPending
	if v2786 != 0 {
		goto L19
	} else {
		goto L593
	}
L593:
	;
	F_tuplestore_puttuple(m, v2760, v2785)
	mBase = m.M
	v2788 = m.ExcPending
	if v2788 != 0 {
		goto L19
	} else {
		goto L594
	}
L594:
	;
	v2789 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v2789)
	goto L589
L595:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+164)) = v2800
	*(*int32)(unsafe.Add(mBase, uint32(v23)+160)) = v2724
	v2806 = F_psprintf(m, int32(_a_F_verify_heapam_99), v23+int32(160))
	mBase = m.M
	v2807 = m.ExcPending
	if v2807 != 0 {
		goto L19
	} else {
		goto L598
	}
L596:
	;
	goto L597
L597:
	;
	v2808 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2799))) = uint16(v2808)
	goto L580
L598:
	;
	v3077 = v2806
	goto L583
L599:
	;
	v2816 = v2733 + v2743&int32(_a_F_verify_heapam_32)
	v2817 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2816)+20)))
	if v2817&int32(_a_F_verify_heapam_47) == int32(_a_F_verify_heapam_27) {
		goto L601
	} else {
		goto L602
	}
L600:
	;
	v2835 = v2829 + v2828&int32(_a_F_verify_heapam_32)
	v2836 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2835)+20)))
	v2837 = int32(768)
	if v2836&v2837 != v2837 {
		goto L605
	} else {
		goto L606
	}
L601:
	;
	v2822 = F_HeapTupleGetUpdateXid(m, v2816)
	mBase = m.M
	v2823 = m.ExcPending
	if v2823 != 0 {
		goto L19
	} else {
		goto L604
	}
L602:
	;
	goto L603
L603:
	;
	v2826 = *(*int32)(unsafe.Add(mBase, uint32(v2816)+4))
	v2827 = v2826
	v2828 = v2739
	v2829 = v2733
	goto L600
L604:
	;
	v2824 = *(*int32)(unsafe.Add(mBase, uint32(v2738)))
	v2825 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[32])))
	v2827 = v2822
	v2828 = v2824
	v2829 = v2825
	goto L600
L605:
	;
	v2841 = *(*int32)(unsafe.Add(mBase, uint32(v2835)))
	v2843 = v2841
	goto L607
L606:
	;
	v2843 = int32(2)
	goto L607
L607:
	;
	if base.B2i32(v2827 == int32(0))|base.B2i32(v2843 != v2827) != 0 {
		goto L580
	} else {
		goto L608
	}
L608:
	;
	v2850 = v23 + int32(_a_F_verify_heapam_26) + v2724<<(uint(int32(1))%32)
	v2851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2850))))
	if v2851 != 0 {
		goto L609
	} else {
		goto L610
	}
L609:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+276)) = v2851
	*(*int32)(unsafe.Add(mBase, uint32(v23)+272)) = v2724
	v2857 = F_psprintf(m, int32(_a_F_verify_heapam_100), v23+int32(272))
	mBase = m.M
	v2858 = m.ExcPending
	if v2858 != 0 {
		goto L19
	} else {
		goto L612
	}
L610:
	;
	goto L611
L611:
	;
	v2859 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2850))) = uint16(v2859)
	v2861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+19)))
	if v2861&int32(64) == int32(0) {
		goto L614
	} else {
		goto L615
	}
L612:
	;
	v3077 = v2857
	goto L583
L613:
	;
	v2961 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2816)+20)))
	v2962 = int32(768)
	if v2961&v2962 != v2962 {
		goto L630
	} else {
		goto L631
	}
L614:
	;
	v2866 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2835)+18)))
	if int32(0) <= v2866 {
		goto L613
	} else {
		goto L617
	}
L615:
	;
	goto L616
L616:
	;
	v2915 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2835)+18)))
	if v2915 < int32(0) {
		goto L613
	} else {
		goto L624
	}
L617:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+256)) = v2724
	v2873 = F_psprintf(m, int32(_a_F_verify_heapam_101), v23+int32(256))
	mBase = m.M
	v2874 = m.ExcPending
	if v2874 != 0 {
		goto L19
	} else {
		goto L618
	}
L618:
	;
	v2875 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v2876 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v2877 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2878 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v2879 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v2879
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v2878
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v2877))
	v2888 = int32(base.Ui32(v2877) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v2888)
	v2890 = F_cstring_to_text(m, v2873)
	mBase = m.M
	v2891 = m.ExcPending
	if v2891 != 0 {
		goto L19
	} else {
		goto L619
	}
L619:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v2890)
	F_pfree(m, v2873)
	mBase = m.M
	v2895 = m.ExcPending
	if v2895 != 0 {
		goto L19
	} else {
		goto L620
	}
L620:
	;
	v2900 = F_heap_form_tuple(m, v2876, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2901 = m.ExcPending
	if v2901 != 0 {
		goto L19
	} else {
		goto L621
	}
L621:
	;
	F_tuplestore_puttuple(m, v2875, v2900)
	mBase = m.M
	v2903 = m.ExcPending
	if v2903 != 0 {
		goto L19
	} else {
		goto L622
	}
L622:
	;
	v2904 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v2904)
	v2906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+19)))
	if v2906&int32(64) == int32(0) {
		goto L613
	} else {
		goto L623
	}
L623:
	;
	goto L616
L624:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+240)) = v2724
	v2922 = F_psprintf(m, int32(_a_F_verify_heapam_102), v23+int32(240))
	mBase = m.M
	v2923 = m.ExcPending
	if v2923 != 0 {
		goto L19
	} else {
		goto L625
	}
L625:
	;
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v2926 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v2927 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v2928 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v2928
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v2927
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v2926))
	v2937 = int32(base.Ui32(v2926) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v2937)
	v2939 = F_cstring_to_text(m, v2922)
	mBase = m.M
	v2940 = m.ExcPending
	if v2940 != 0 {
		goto L19
	} else {
		goto L626
	}
L626:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v2939)
	F_pfree(m, v2922)
	mBase = m.M
	v2944 = m.ExcPending
	if v2944 != 0 {
		goto L19
	} else {
		goto L627
	}
L627:
	;
	v2949 = F_heap_form_tuple(m, v2925, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v2950 = m.ExcPending
	if v2950 != 0 {
		goto L19
	} else {
		goto L628
	}
L628:
	;
	F_tuplestore_puttuple(m, v2924, v2949)
	mBase = m.M
	v2952 = m.ExcPending
	if v2952 != 0 {
		goto L19
	} else {
		goto L629
	}
L629:
	;
	v2953 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v2953)
	goto L613
L630:
	;
	v2966 = *(*int32)(unsafe.Add(mBase, uint32(v2816)))
	v2967 = v2966
	goto L632
L631:
	;
	v2967 = int32(2)
	goto L632
L632:
	;
	v2968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	v2970 = v23 + int32(_a_F_verify_heapam_30)
	v2972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2968+v2970))))
	if v2972 != int32(1) {
		v3034 = v2968
		goto L633
	} else {
		goto L634
	}
L633:
	;
	v3040 = v3034 & int32(_a_F_verify_heapam_6)
	v3042 = v23 + int32(_a_F_verify_heapam_30)
	v3044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3040+v3042))))
	if v3044 != int32(1) {
		goto L580
	} else {
		goto L645
	}
L634:
	;
	v2976 = v23 + int32(1040)
	v2977 = int32(2)
	v2980 = *(*int32)(unsafe.Add(mBase, uint32(v2976+v2968<<(uint(v2977)%32))))
	if v2980 != v2977 {
		v3034 = v2968
		goto L633
	} else {
		goto L635
	}
L635:
	;
	v2984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2724+v2970))))
	if v2984 != int32(1) {
		v3034 = v2968
		goto L633
	} else {
		goto L636
	}
L636:
	;
	v2988 = *(*int32)(unsafe.Add(mBase, uint32(v2976+v2737)))
	if v2988 != 0 {
		v3034 = v2968
		goto L633
	} else {
		goto L637
	}
L637:
	;
	v2989 = F_TransactionIdIsInProgress(m, v2967)
	mBase = m.M
	v2990 = m.ExcPending
	if v2990 != 0 {
		goto L19
	} else {
		goto L638
	}
L638:
	;
	v2991 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	if v2989 == int32(0) {
		v3034 = v2991
		goto L633
	} else {
		goto L639
	}
L639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+232)) = v2827
	*(*int32)(unsafe.Add(mBase, uint32(v23)+228)) = v2991
	*(*int32)(unsafe.Add(mBase, uint32(v23)+224)) = v2967
	v3000 = F_psprintf(m, int32(_a_F_verify_heapam_103), v23+int32(224))
	mBase = m.M
	v3001 = m.ExcPending
	if v3001 != 0 {
		goto L19
	} else {
		goto L640
	}
L640:
	;
	v3002 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v3003 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v3004 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v3005 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v3006 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v3006
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v3005
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v3004))
	v3015 = int32(base.Ui32(v3004) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v3015)
	v3017 = F_cstring_to_text(m, v3000)
	mBase = m.M
	v3018 = m.ExcPending
	if v3018 != 0 {
		goto L19
	} else {
		goto L641
	}
L641:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v3017)
	F_pfree(m, v3000)
	mBase = m.M
	v3022 = m.ExcPending
	if v3022 != 0 {
		goto L19
	} else {
		goto L642
	}
L642:
	;
	v3027 = F_heap_form_tuple(m, v3003, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v3028 = m.ExcPending
	if v3028 != 0 {
		goto L19
	} else {
		goto L643
	}
L643:
	;
	F_tuplestore_puttuple(m, v3002, v3027)
	mBase = m.M
	v3030 = m.ExcPending
	if v3030 != 0 {
		goto L19
	} else {
		goto L644
	}
L644:
	;
	v3031 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v3031)
	v3033 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	v3034 = v3033
	goto L633
L645:
	;
	v3048 = v23 + int32(1040)
	v3052 = *(*int32)(unsafe.Add(mBase, uint32(v3048+v3040<<(uint(int32(2))%32))))
	if v3052 != int32(3) {
		goto L580
	} else {
		goto L646
	}
L646:
	;
	v3056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2724+v3042))))
	if v3056 != int32(1) {
		goto L580
	} else {
		goto L647
	}
L647:
	;
	v3060 = *(*int32)(unsafe.Add(mBase, uint32(v3048+v2737)))
	switch v3060 {
	case 0:
		goto L648
	default:
		goto L580
	case 2:
		goto L649
	}
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+216)) = v2827
	*(*int32)(unsafe.Add(mBase, uint32(v23)+212)) = v3040
	*(*int32)(unsafe.Add(mBase, uint32(v23)+208)) = v2967
	v3075 = F_psprintf(m, int32(_a_F_verify_heapam_104), v23+int32(208))
	mBase = m.M
	v3076 = m.ExcPending
	if v3076 != 0 {
		goto L19
	} else {
		goto L651
	}
L649:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v2827
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v3040
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v2967
	v3067 = F_psprintf(m, int32(_a_F_verify_heapam_105), v23+int32(192))
	mBase = m.M
	v3068 = m.ExcPending
	if v3068 != 0 {
		goto L19
	} else {
		goto L650
	}
L650:
	;
	v3077 = v3067
	goto L583
L651:
	;
	v3077 = v3075
	goto L583
L652:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v3100)
	F_pfree(m, v3077)
	mBase = m.M
	v3105 = m.ExcPending
	if v3105 != 0 {
		goto L19
	} else {
		goto L653
	}
L653:
	;
	v3110 = F_heap_form_tuple(m, v3086, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v3111 = m.ExcPending
	if v3111 != 0 {
		goto L19
	} else {
		goto L654
	}
L654:
	;
	F_tuplestore_puttuple(m, v3085, v3110)
	mBase = m.M
	v3113 = m.ExcPending
	if v3113 != 0 {
		goto L19
	} else {
		goto L655
	}
L655:
	;
	v3114 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v3114)
	goto L580
L656:
	;
	goto L579
L657:
	;
	v3159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(_a_F_verify_heapam_30)+v3138))))
	if v3159 != int32(1) {
		v3224 = v3139
		goto L659
	} else {
		goto L660
	}
L658:
	;
	goto L188
L659:
	;
	v3229 = v3224 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))) = uint16(v3229)
	v3232 = v3229 & int32(_a_F_verify_heapam_6)
	if base.Ui32(v3232) <= base.Ui32(v675) {
		v3138 = v3232
		v3139 = v3229
		goto L657
	} else {
		goto L670
	}
L660:
	;
	v3163 = v3138 << (uint(int32(2)) % 32)
	v3167 = *(*int32)(unsafe.Add(mBase, uint32(v3163+(v23+int32(1040)))))
	switch v3167 {
	case 0, 2:
		goto L661
	default:
		v3224 = v3139
		goto L659
	}
L661:
	;
	v3173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(_a_F_verify_heapam_26)+v3138<<(uint(int32(1))%32)))))
	if v3173 != 0 {
		v3224 = v3139
		goto L659
	} else {
		goto L662
	}
L662:
	;
	v3174 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[32])))
	v3176 = *(*int32)(unsafe.Add(mBase, uint32(v3174+v3163)+20))
	if v3176&int32(_a_F_verify_heapam_96) == int32(_a_F_verify_heapam_97) {
		v3224 = v3139
		goto L659
	} else {
		goto L663
	}
L663:
	;
	v3184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3174+v3176&int32(_a_F_verify_heapam_32))+18)))
	if int32(0) <= v3184 {
		v3224 = v3139
		goto L659
	} else {
		goto L664
	}
L664:
	;
	v3189 = F_psprintf(m, int32(_a_F_verify_heapam_106), int32(0))
	mBase = m.M
	v3190 = m.ExcPending
	if v3190 != 0 {
		goto L19
	} else {
		goto L665
	}
L665:
	;
	v3191 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v3192 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	v3193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[8]))))
	v3194 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[29]))))
	v3195 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[40]))) = v3195
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[41]))) = v3194
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[42]))) = base.I64_extend16_s(base.I64_extend_i32_u(v3193))
	v3204 = int32(base.Ui32(v3193) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[43]))) = uint8(v3204)
	v3206 = F_cstring_to_text(m, v3189)
	mBase = m.M
	v3207 = m.ExcPending
	if v3207 != 0 {
		goto L19
	} else {
		goto L666
	}
L666:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[44]))) = base.I64_extend_i32_u(v3206)
	F_pfree(m, v3189)
	mBase = m.M
	v3211 = m.ExcPending
	if v3211 != 0 {
		goto L19
	} else {
		goto L667
	}
L667:
	;
	v3216 = F_heap_form_tuple(m, v3192, v23+int32(_a_F_verify_heapam_35), v23+int32(_a_F_verify_heapam_36))
	mBase = m.M
	v3217 = m.ExcPending
	if v3217 != 0 {
		goto L19
	} else {
		goto L668
	}
L668:
	;
	F_tuplestore_puttuple(m, v3191, v3216)
	mBase = m.M
	v3219 = m.ExcPending
	if v3219 != 0 {
		goto L19
	} else {
		goto L669
	}
L669:
	;
	v3220 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v3220)
	v3222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[33]))))
	v3224 = v3222
	goto L659
L670:
	;
	goto L658
L671:
	;
	v3257 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[55])))
	if v3257 != 0 {
		goto L672
	} else {
		goto L673
	}
L672:
	;
	v3258 = *(*int32)(unsafe.Add(mBase, uint32(v3257)+4))
	if int32(0) < v3258 {
		goto L675
	} else {
		goto L676
	}
L673:
	;
	goto L674
L674:
	;
	if v41 == int64(0) {
		goto L170
	} else {
		goto L744
	}
L675:
	;
	v3270 = int32(0)
	goto L678
L676:
	;
	v3638 = v3257
	goto L677
L677:
	;
	F_list_free_deep(m, v3638)
	mBase = m.M
	v3640 = m.ExcPending
	if v3640 != 0 {
		goto L19
	} else {
		goto L743
	}
L678:
	;
	v3282 = *(*int32)(unsafe.Add(mBase, uint32(v3257)+12))
	v3286 = *(*int32)(unsafe.Add(mBase, uint32(v3282+v3270<<(uint(int32(2))%32))))
	v3287 = *(*int32)(unsafe.Add(mBase, uint32(v3286)+4))
	v3289 = v23 + int32(_a_F_verify_heapam_35)
	v3293 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3286)+8)))
	F_ScanKeyInit(m, v3289, int32(1), int32(3), int32(184), v3293)
	mBase = m.M
	v3295 = m.ExcPending
	if v3295 != 0 {
		goto L19
	} else {
		goto L680
	}
L679:
	;
	v3617 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[55])))
	v3638 = v3617
	goto L677
L680:
	;
	v3297 = v3287 & int32(1073741823)
	v3301 = base.I32_div_u_s(v3297-int32(1), int32(1996))
	v3302 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[19])))
	v3303 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[21])))
	v3304 = F_get_toast_snapshot(m)
	mBase = m.M
	v3305 = m.ExcPending
	if v3305 != 0 {
		goto L19
	} else {
		goto L683
	}
L681:
	;
	v3614 = v3270 + int32(1)
	v3615 = *(*int32)(unsafe.Add(mBase, uint32(v3257)+4))
	if v3614 < v3615 {
		v3270 = v3614
		goto L678
	} else {
		goto L742
	}
L682:
	;
	v3562 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3286)+22)))
	v3563 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3286)+16)))
	v3564 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3286)+20)))
	v3565 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v3566 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[49]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[56]))) = v3564
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = v3563
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[57]))) = base.I64_extend16_s(base.I64_extend_i32_u(v3562))
	v3575 = int32(base.Ui32(v3562) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[58]))) = uint8(v3575)
	v3577 = F_cstring_to_text(m, v3545)
	mBase = m.M
	v3578 = m.ExcPending
	if v3578 != 0 {
		goto L19
	} else {
		goto L738
	}
L683:
	;
	v3307 = F_systable_beginscan_ordered(m, v3302, v3303, v3304, int32(1), v3289)
	mBase = m.M
	v3308 = m.ExcPending
	if v3308 != 0 {
		goto L19
	} else {
		goto L684
	}
L684:
	;
	v3310 = F_systable_getnext_ordered(m, v3307, int32(1))
	mBase = m.M
	v3311 = m.ExcPending
	if v3311 != 0 {
		goto L19
	} else {
		goto L685
	}
L685:
	;
	if v3310 != 0 {
		goto L686
	} else {
		goto L687
	}
L686:
	;
	v3319 = int32(0)
	v3321 = v3310
	goto L689
L687:
	;
	goto L688
L688:
	;
	F_systable_endscan_ordered(m, v3307)
	mBase = m.M
	v3534 = m.ExcPending
	if v3534 != 0 {
		goto L19
	} else {
		goto L736
	}
L689:
	;
	v3337 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[19])))
	v3338 = *(*int32)(unsafe.Add(mBase, uint32(v3337)+52))
	v3341 = F_fastgetattr_5(m, v3321, int32(2), v3338, v23+int32(_a_F_verify_heapam_62))
	mBase = m.M
	v3342 = m.ExcPending
	if v3342 != 0 {
		goto L19
	} else {
		goto L691
	}
L690:
	;
	F_systable_endscan_ordered(m, v3307)
	mBase = m.M
	v3522 = m.ExcPending
	if v3522 != 0 {
		goto L19
	} else {
		goto L733
	}
L691:
	;
	v3343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[50]))))
	if v3343 == int32(1) {
		goto L694
	} else {
		goto L695
	}
L692:
	;
	v3519 = F_systable_getnext_ordered(m, v3307, int32(1))
	mBase = m.M
	v3520 = m.ExcPending
	if v3520 != 0 {
		goto L19
	} else {
		goto L731
	}
L693:
	;
	v3480 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3286)+22)))
	v3481 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3286)+16)))
	v3482 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3286)+20)))
	v3483 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v3484 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[49]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[56]))) = v3482
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = v3481
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[57]))) = base.I64_extend16_s(base.I64_extend_i32_u(v3480))
	v3493 = int32(base.Ui32(v3480) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[58]))) = uint8(v3493)
	v3495 = F_cstring_to_text(m, v3474)
	mBase = m.M
	v3496 = m.ExcPending
	if v3496 != 0 {
		goto L19
	} else {
		goto L727
	}
L694:
	;
	v3346 = *(*int32)(unsafe.Add(mBase, uint32(v3286)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v3346
	v3351 = F_psprintf(m, int32(_a_F_verify_heapam_107), v23-int32(-64))
	mBase = m.M
	v3352 = m.ExcPending
	if v3352 != 0 {
		goto L19
	} else {
		goto L697
	}
L695:
	;
	goto L696
L696:
	;
	v3353 = base.I32_wrap_i64(v3341)
	if v3353 != v3319 {
		goto L698
	} else {
		goto L699
	}
L697:
	;
	v3473 = v3319
	v3474 = v3351
	goto L693
L698:
	;
	v3355 = *(*int32)(unsafe.Add(mBase, uint32(v3286)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+152)) = v3319
	*(*int32)(unsafe.Add(mBase, uint32(v23)+148)) = v3353
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = v3355
	v3362 = F_psprintf(m, int32(_a_F_verify_heapam_108), v23+int32(144))
	mBase = m.M
	v3363 = m.ExcPending
	if v3363 != 0 {
		goto L19
	} else {
		goto L701
	}
L699:
	;
	goto L700
L700:
	;
	v3400 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[19])))
	v3401 = *(*int32)(unsafe.Add(mBase, uint32(v3400)+52))
	v3404 = F_fastgetattr_5(m, v3321, int32(3), v3401, v23+int32(_a_F_verify_heapam_62))
	mBase = m.M
	v3405 = m.ExcPending
	if v3405 != 0 {
		goto L19
	} else {
		goto L706
	}
L701:
	;
	v3364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3286)+22)))
	v3365 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3286)+16)))
	v3366 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3286)+20)))
	v3367 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[11])))
	v3368 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[10])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[49]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[56]))) = v3366
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[39]))) = v3365
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[57]))) = base.I64_extend16_s(base.I64_extend_i32_u(v3364))
	v3377 = int32(base.Ui32(v3364) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[58]))) = uint8(v3377)
	v3379 = F_cstring_to_text(m, v3362)
	mBase = m.M
	v3380 = m.ExcPending
	if v3380 != 0 {
		goto L19
	} else {
		goto L702
	}
L702:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[59]))) = base.I64_extend_i32_u(v3379)
	F_pfree(m, v3362)
	mBase = m.M
	v3384 = m.ExcPending
	if v3384 != 0 {
		goto L19
	} else {
		goto L703
	}
L703:
	;
	v3389 = F_heap_form_tuple(m, v3368, v23+int32(_a_F_verify_heapam_36), v23+int32(_a_F_verify_heapam_58))
	mBase = m.M
	v3390 = m.ExcPending
	if v3390 != 0 {
		goto L19
	} else {
		goto L704
	}
L704:
	;
	F_tuplestore_puttuple(m, v3367, v3389)
	mBase = m.M
	v3392 = m.ExcPending
	if v3392 != 0 {
		goto L19
	} else {
		goto L705
	}
L705:
	;
	v3393 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v3393)
	goto L700
L706:
	;
	v3406 = int32(1)
	v3407 = v3353 + v3406
	v3408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[50]))))
	if v3408 == v3406 {
		goto L707
	} else {
		goto L708
	}
L707:
	;
	v3411 = *(*int32)(unsafe.Add(mBase, uint32(v3286)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v3353
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v3411
	v3417 = F_psprintf(m, int32(_a_F_verify_heapam_109), v23+int32(80))
	mBase = m.M
	v3418 = m.ExcPending
	if v3418 != 0 {
		goto L19
	} else {
		goto L710
	}
L708:
	;
	goto L709
L709:
	;
	v3419 = base.I32_wrap_i64(v3404)
	v3420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3419))))
	if v3420&int32(3) == int32(0) {
		goto L713
	} else {
		goto L714
	}
L710:
	;
	v3473 = v3407
	v3474 = v3417
	goto L693
L711:
	;
	v3463 = *(*int32)(unsafe.Add(mBase, uint32(v3286)+8))
	v3464 = *(*int32)(unsafe.Add(mBase, uint32(v3419)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+136)) = v3464
	*(*int32)(unsafe.Add(mBase, uint32(v23)+132)) = v3353
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v3463
	v3471 = F_psprintf(m, int32(_a_F_verify_heapam_110), v23+int32(128))
	mBase = m.M
	v3472 = m.ExcPending
	if v3472 != 0 {
		goto L19
	} else {
		goto L726
	}
L712:
	;
	if v3301 < v3353 {
		goto L717
	} else {
		goto L718
	}
L713:
	;
	v3425 = *(*int32)(unsafe.Add(mBase, uint32(v3419)))
	v3438 = int32(base.Ui32(v3425)>>(uint(int32(2))%32)) - int32(4)
	goto L712
L714:
	;
	goto L715
L715:
	;
	if v3420&int32(1) == int32(0) {
		goto L711
	} else {
		goto L716
	}
L716:
	;
	v3434 = int32(1)
	v3438 = int32(base.Ui32(v3420)>>(uint(v3434)%32)) - v3434
	goto L712
L717:
	;
	v3440 = *(*int32)(unsafe.Add(mBase, uint32(v3286)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+104)) = v3301
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v3353
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v3440
	v3447 = F_psprintf(m, int32(_a_F_verify_heapam_111), v23+int32(96))
	mBase = m.M
	v3448 = m.ExcPending
	if v3448 != 0 {
		goto L19
	} else {
		goto L720
	}
L718:
	;
	goto L719
L719:
	;
	if v3353 < v3301 {
		goto L721
	} else {
		goto L722
	}
L720:
	;
	v3473 = v3407
	v3474 = v3447
	goto L693
L721:
	;
	v3451 = int32(1996)
	goto L723
L722:
	;
	v3451 = v3301*int32(-1996) + v3297
	goto L723
L723:
	;
	if v3438 == v3451 {
		v3511 = v3407
		goto L692
	} else {
		goto L724
	}
L724:
	;
	v3453 = *(*int32)(unsafe.Add(mBase, uint32(v3286)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+124)) = v3451
	*(*int32)(unsafe.Add(mBase, uint32(v23)+120)) = v3438
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = v3353
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v3453
	v3461 = F_psprintf(m, int32(_a_F_verify_heapam_112), v23+int32(112))
	mBase = m.M
	v3462 = m.ExcPending
	if v3462 != 0 {
		goto L19
	} else {
		goto L725
	}
L725:
	;
	v3473 = v3407
	v3474 = v3461
	goto L693
L726:
	;
	v3473 = v3407
	v3474 = v3471
	goto L693
L727:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[59]))) = base.I64_extend_i32_u(v3495)
	F_pfree(m, v3474)
	mBase = m.M
	v3500 = m.ExcPending
	if v3500 != 0 {
		goto L19
	} else {
		goto L728
	}
L728:
	;
	v3505 = F_heap_form_tuple(m, v3484, v23+int32(_a_F_verify_heapam_36), v23+int32(_a_F_verify_heapam_58))
	mBase = m.M
	v3506 = m.ExcPending
	if v3506 != 0 {
		goto L19
	} else {
		goto L729
	}
L729:
	;
	F_tuplestore_puttuple(m, v3483, v3505)
	mBase = m.M
	v3508 = m.ExcPending
	if v3508 != 0 {
		goto L19
	} else {
		goto L730
	}
L730:
	;
	v3509 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v3509)
	v3511 = v3473
	goto L692
L731:
	;
	if v3519 != 0 {
		v3319 = v3511
		v3321 = v3519
		goto L689
	} else {
		goto L732
	}
L732:
	;
	goto L690
L733:
	;
	if v3301 < v3511 {
		goto L681
	} else {
		goto L734
	}
L734:
	;
	v3524 = *(*int32)(unsafe.Add(mBase, uint32(v3286)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = v3511
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v3301
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v3524
	v3531 = F_psprintf(m, int32(_a_F_verify_heapam_113), v23+int32(48))
	mBase = m.M
	v3532 = m.ExcPending
	if v3532 != 0 {
		goto L19
	} else {
		goto L735
	}
L735:
	;
	v3545 = v3531
	goto L682
L736:
	;
	v3535 = *(*int32)(unsafe.Add(mBase, uint32(v3286)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v3535
	v3540 = F_psprintf(m, int32(_a_F_verify_heapam_114), v23+int32(32))
	mBase = m.M
	v3541 = m.ExcPending
	if v3541 != 0 {
		goto L19
	} else {
		goto L737
	}
L737:
	;
	v3545 = v3540
	goto L682
L738:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[59]))) = base.I64_extend_i32_u(v3577)
	F_pfree(m, v3545)
	mBase = m.M
	v3582 = m.ExcPending
	if v3582 != 0 {
		goto L19
	} else {
		goto L739
	}
L739:
	;
	v3587 = F_heap_form_tuple(m, v3566, v23+int32(_a_F_verify_heapam_36), v23+int32(_a_F_verify_heapam_58))
	mBase = m.M
	v3588 = m.ExcPending
	if v3588 != 0 {
		goto L19
	} else {
		goto L740
	}
L740:
	;
	F_tuplestore_puttuple(m, v3565, v3587)
	mBase = m.M
	v3590 = m.ExcPending
	if v3590 != 0 {
		goto L19
	} else {
		goto L741
	}
L741:
	;
	v3591 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))) = uint8(v3591)
	goto L681
L742:
	;
	goto L679
L743:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[55]))) = int32(0)
	goto L674
L744:
	;
	v3665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[45]))))
	if v3665&int32(1) == int32(0) {
		goto L170
	} else {
		goto L745
	}
L745:
	;
	goto L172
L746:
	;
	v3692 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[0])))
	if v3692 != 0 {
		goto L747
	} else {
		goto L748
	}
L747:
	;
	F_ReleaseBuffer(m, v3692)
	mBase = m.M
	v3694 = m.ExcPending
	if v3694 != 0 {
		goto L19
	} else {
		goto L750
	}
L748:
	;
	goto L749
L749:
	;
	v3695 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[20])))
	if v3695 != 0 {
		goto L751
	} else {
		goto L752
	}
L750:
	;
	goto L749
L751:
	;
	v3696 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[18])))
	F_toast_close_indexes(m, v3695, v3696)
	mBase = m.M
	v3698 = m.ExcPending
	if v3698 != 0 {
		goto L19
	} else {
		goto L754
	}
L752:
	;
	goto L753
L753:
	;
	v3699 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[19])))
	if v3699 != 0 {
		goto L755
	} else {
		goto L756
	}
L754:
	;
	goto L753
L755:
	;
	F_relation_close(m, v3699, int32(1))
	mBase = m.M
	v3702 = m.ExcPending
	if v3702 != 0 {
		goto L19
	} else {
		goto L758
	}
L756:
	;
	goto L757
L757:
	;
	v3703 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_verify_heapam[12])))
	F_relation_close(m, v3703, int32(1))
	mBase = m.M
	v3706 = m.ExcPending
	if v3706 != 0 {
		goto L19
	} else {
		goto L759
	}
L758:
	;
	goto L757
L759:
	;
	goto L2
L760:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+432)) = int32(-1)
	F_errmsg_internal(m, int32(_a_F_verify_heapam_115), v23+int32(432))
	mBase = m.M
	v3744 = m.ExcPending
	if v3744 != 0 {
		goto L19
	} else {
		goto L761
	}
L761:
	;
	F_errfinish(m, int32(_a_F_verify_heapam_116), int32(123), int32(_a_F_verify_heapam_117))
	mBase = m.M
	v3749 = m.ExcPending
	if v3749 != 0 {
		goto L19
	} else {
		goto L762
	}
L762:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
