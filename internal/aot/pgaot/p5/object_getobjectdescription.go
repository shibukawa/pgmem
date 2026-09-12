package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_getObjectDescription(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
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
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
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
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
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
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v989 int32
	_ = v989
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1086 int32
	_ = v1086
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
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1134 int32
	_ = v1134
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1212 int32
	_ = v1212
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
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1349 int32
	_ = v1349
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
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1424 int32
	_ = v1424
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1461 int32
	_ = v1461
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1530 int32
	_ = v1530
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1570 int32
	_ = v1570
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1637 int32
	_ = v1637
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1695 int32
	_ = v1695
	var v1700 int32
	_ = v1700
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1714 int32
	_ = v1714
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1735 int32
	_ = v1735
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1763 int32
	_ = v1763
	var v1771 int32
	_ = v1771
	var v1780 int32
	_ = v1780
	var v1788 int32
	_ = v1788
	var v1797 int32
	_ = v1797
	var v1805 int32
	_ = v1805
	var v1814 int32
	_ = v1814
	var v1822 int32
	_ = v1822
	var v1830 int32
	_ = v1830
	var v1838 int32
	_ = v1838
	var v1847 int32
	_ = v1847
	var v1855 int32
	_ = v1855
	var v1861 int32
	_ = v1861
	var v1864 int32
	_ = v1864
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1878 int32
	_ = v1878
	v9 = m.G0
	v11 = v9 - int32(1424)
	m.G0 = v11
	F_initStringInfo(m, v11+int32(1408))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v19 <= int32(2752) {
			if v19 <= int32(2327) {
				switch v19 - int32(1213) {
				case 0:
					v1274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v1275 = F_get_tablespace_name(m, v1274)
					mBase = m.M
					v1276 = m.ExcPending
					if v1276 != 0 {
						return int32(0)
					} else {
						if v1275 == int32(0) {
							if l1 != 0 {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v1282 = m.ExcPending
								if v1282 != 0 {
									return int32(0)
								} else {
									v1283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+800)) = v1283
									F_errmsg_internal(m, int32(50552), v11+int32(800))
									mBase = m.M
									v1289 = m.ExcPending
									if v1289 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(472250), int32(3699), int32(235362))
										mBase = m.M
										v1294 = m.ExcPending
										if v1294 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+816)) = v1275
							F_appendStringInfo(m, v11+int32(1408), int32(186293), v11+int32(816))
							mBase = m.M
							v1302 = m.ExcPending
							if v1302 != 0 {
								return int32(0)
							} else {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							}
						}
					}
				case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 35, 36, 37, 38, 39, 40, 41, 43, 44, 45:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v1669 = m.ExcPending
					if v1669 != 0 {
						return int32(0)
					} else {
						v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1670
						F_errmsg_internal(m, int32(55485), v11)
						mBase = m.M
						v1674 = m.ExcPending
						if v1674 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(472250), int32(4072), int32(235362))
							mBase = m.M
							v1679 = m.ExcPending
							if v1679 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 34:
					v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v141 = F_format_type_extended(m, v138, int32(-1), int32(8))
					mBase = m.M
					v142 = m.ExcPending
					if v142 != 0 {
						return int32(0)
					} else {
						if v141 == int32(0) {
							v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
							v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
							m.G0 = v11 + int32(1424)
							if v1872 != 0 {
								v1878 = v1873
							} else {
								v1878 = int32(0)
							}
							return v1878
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v141
							F_appendStringInfo(m, v11+int32(1408), int32(184124), v11+int32(48))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return int32(0)
							} else {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							}
						}
					}
				case 42:
					v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v126 = F_format_procedure_extended(m, v124, int32(1))
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return int32(0)
					} else {
						if v126 == int32(0) {
							v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
							v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
							m.G0 = v11 + int32(1424)
							if v1872 != 0 {
								v1878 = v1873
							} else {
								v1878 = int32(0)
							}
							return v1878
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v126
							F_appendStringInfo(m, v11+int32(1408), int32(174107), v11+int32(32))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return int32(0)
							} else {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							}
						}
					}
				case 46:
					v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v88 == int32(0) {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_getRelationDescription(m, v11+int32(1408), v93, l1)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
							v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
							m.G0 = v11 + int32(1424)
							if v1872 != 0 {
								v1878 = v1873
							} else {
								v1878 = int32(0)
							}
							return v1878
						}
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v98 = F_get_attname(m, v96, base.I32_extend16_s(v88), l1)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							if v98 == int32(0) {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							} else {
								F_initStringInfo(m, v11+int32(1360))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									F_getRelationDescription(m, v11+int32(1360), v108, l1)
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v98
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v112
										F_appendStringInfo(m, v11+int32(1408), int32(177015), v11+int32(16))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
											F_pfree(m, v121)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
												v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
												m.G0 = v11 + int32(1424)
												if v1872 != 0 {
													v1878 = v1873
												} else {
													v1878 = int32(0)
												}
												return v1878
											}
										}
									}
								}
							}
						}
					}
				case 47:
					v1165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v1166 = F_GetUserNameFromId(m, v1165, l1)
					mBase = m.M
					v1167 = m.ExcPending
					if v1167 != 0 {
						return int32(0)
					} else {
						if v1166 == int32(0) {
							v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
							v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
							m.G0 = v11 + int32(1424)
							if v1872 != 0 {
								v1878 = v1873
							} else {
								v1878 = int32(0)
							}
							return v1878
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+720)) = v1166
							F_appendStringInfo(m, v11+int32(1408), int32(185265), v11+int32(720))
							mBase = m.M
							v1177 = m.ExcPending
							if v1177 != 0 {
								return int32(0)
							} else {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							}
						}
					}
				case 48:
					v1180 = F_table_open(m, int32(1261), int32(1))
					mBase = m.M
					v1181 = m.ExcPending
					if v1181 != 0 {
						return int32(0)
					} else {
						v1187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_ScanKeyInit(m, v11+int32(1360), int32(1), int32(3), int32(184), v1187)
						mBase = m.M
						v1189 = m.ExcPending
						if v1189 != 0 {
							return int32(0)
						} else {
							v1191 = int32(1)
							v1196 = F_systable_beginscan(m, v1180, int32(6303), v1191, int32(0), v1191, v11+int32(1360))
							mBase = m.M
							v1197 = m.ExcPending
							if v1197 != 0 {
								return int32(0)
							} else {
								v1198 = F_systable_getnext(m, v1196)
								mBase = m.M
								v1199 = m.ExcPending
								if v1199 != 0 {
									return int32(0)
								} else {
									if v1198 == int32(0) {
										if l1 != 0 {
											F_systable_endscan(m, v1196)
											mBase = m.M
											v1241 = m.ExcPending
											if v1241 != 0 {
												return int32(0)
											} else {
												F_sequence_close(m, v1180, int32(1))
												mBase = m.M
												v1244 = m.ExcPending
												if v1244 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v1205 = m.ExcPending
											if v1205 != 0 {
												return int32(0)
											} else {
												v1206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+736)) = v1206
												F_errmsg_internal(m, int32(41903), v11+int32(736))
												mBase = m.M
												v1212 = m.ExcPending
												if v1212 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(472250), int32(3656), int32(235362))
													mBase = m.M
													v1217 = m.ExcPending
													if v1217 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+16))
										v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1218)+22)))
										v1220 = v1218 + v1219
										v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1220)+8))
										v1223 = F_GetUserNameFromId(m, v1221, int32(0))
										mBase = m.M
										v1224 = m.ExcPending
										if v1224 != 0 {
											return int32(0)
										} else {
											v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1220)+4))
											v1227 = F_GetUserNameFromId(m, v1225, int32(0))
											mBase = m.M
											v1228 = m.ExcPending
											if v1228 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+756)) = v1227
												*(*int32)(unsafe.Add(mBase, uint32(v11)+752)) = v1223
												F_appendStringInfo(m, v11+int32(1408), int32(185240), v11+int32(752))
												mBase = m.M
												v1237 = m.ExcPending
												if v1237 != 0 {
													return int32(0)
												} else {
													F_systable_endscan(m, v1196)
													mBase = m.M
													v1241 = m.ExcPending
													if v1241 != 0 {
														return int32(0)
													} else {
														F_sequence_close(m, v1180, int32(1))
														mBase = m.M
														v1244 = m.ExcPending
														if v1244 != 0 {
															return int32(0)
														} else {
															v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
															v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
															m.G0 = v11 + int32(1424)
															if v1872 != 0 {
																v1878 = v1873
															} else {
																v1878 = int32(0)
															}
															return v1878
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				case 49:
					v1245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v1246 = F_get_database_name(m, v1245)
					mBase = m.M
					v1247 = m.ExcPending
					if v1247 != 0 {
						return int32(0)
					} else {
						if v1246 == int32(0) {
							if l1 != 0 {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v1253 = m.ExcPending
								if v1253 != 0 {
									return int32(0)
								} else {
									v1254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+768)) = v1254
									F_errmsg_internal(m, int32(47107), v11+int32(768))
									mBase = m.M
									v1260 = m.ExcPending
									if v1260 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(472250), int32(3683), int32(235362))
										mBase = m.M
										v1265 = m.ExcPending
										if v1265 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+784)) = v1246
							F_appendStringInfo(m, v11+int32(1408), int32(178187), v11+int32(784))
							mBase = m.M
							v1273 = m.ExcPending
							if v1273 != 0 {
								return int32(0)
							} else {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							}
						}
					}
				default:
					switch v19 - int32(1417) {
					case 0:
						v1319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1320 = F_GetForeignServerExtended(m, v1319, l1)
						mBase = m.M
						v1321 = m.ExcPending
						if v1321 != 0 {
							return int32(0)
						} else {
							if v1320 == int32(0) {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							} else {
								v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1320)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+848)) = v1324
								F_appendStringInfo(m, v11+int32(1408), int32(172251), v11+int32(848))
								mBase = m.M
								v1332 = m.ExcPending
								if v1332 != 0 {
									return int32(0)
								} else {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								}
							}
						}
					case 1:
						v1334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1335 = F_SearchSysCache1(m, int32(83), v1334)
						mBase = m.M
						v1336 = m.ExcPending
						if v1336 != 0 {
							return int32(0)
						} else {
							if v1335 == int32(0) {
								if l1 != 0 {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1342 = m.ExcPending
									if v1342 != 0 {
										return int32(0)
									} else {
										v1343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+864)) = v1343
										F_errmsg_internal(m, int32(45940), v11+int32(864))
										mBase = m.M
										v1349 = m.ExcPending
										if v1349 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(472250), int32(3741), int32(235362))
											mBase = m.M
											v1354 = m.ExcPending
											if v1354 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+16))
								v1356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1355)+22)))
								v1357 = v1355 + v1356
								v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1357)+4))
								v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1357)+8))
								v1360 = F_GetForeignServer(m, v1359)
								mBase = m.M
								v1361 = m.ExcPending
								if v1361 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v1335)
									mBase = m.M
									v1363 = m.ExcPending
									if v1363 != 0 {
										return int32(0)
									} else {
										if v1358 != 0 {
											v1365 = F_GetUserNameFromId(m, v1358, int32(0))
											mBase = m.M
											v1366 = m.ExcPending
											if v1366 != 0 {
												return int32(0)
											} else {
												v1368 = v1365
												v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+884)) = v1369
												*(*int32)(unsafe.Add(mBase, uint32(v11)+880)) = v1368
												F_appendStringInfo(m, v11+int32(1408), int32(172153), v11+int32(880))
												mBase = m.M
												v1378 = m.ExcPending
												if v1378 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										} else {
											v1368 = int32(470204)
											v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+884)) = v1369
											*(*int32)(unsafe.Add(mBase, uint32(v11)+880)) = v1368
											F_appendStringInfo(m, v11+int32(1408), int32(172153), v11+int32(880))
											mBase = m.M
											v1378 = m.ExcPending
											if v1378 != 0 {
												return int32(0)
											} else {
												v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
												v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
												m.G0 = v11 + int32(1424)
												if v1872 != 0 {
													v1878 = v1873
												} else {
													v1878 = int32(0)
												}
												return v1878
											}
										}
									}
								}
							}
						}
					default:
						if v19 == int32(826) {
							v1703 = F_table_open(m, int32(826), int32(1))
							mBase = m.M
							v1704 = m.ExcPending
							if v1704 != 0 {
								return int32(0)
							} else {
								v1710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								F_ScanKeyInit(m, v11+int32(1360), int32(1), int32(3), int32(184), v1710)
								mBase = m.M
								v1712 = m.ExcPending
								if v1712 != 0 {
									return int32(0)
								} else {
									v1714 = int32(1)
									v1719 = F_systable_beginscan(m, v1703, int32(828), v1714, int32(0), v1714, v11+int32(1360))
									mBase = m.M
									v1720 = m.ExcPending
									if v1720 != 0 {
										return int32(0)
									} else {
										v1721 = F_systable_getnext(m, v1719)
										mBase = m.M
										v1722 = m.ExcPending
										if v1722 != 0 {
											return int32(0)
										} else {
											if v1721 == int32(0) {
												if l1 != 0 {
													F_systable_endscan(m, v1719)
													mBase = m.M
													v1861 = m.ExcPending
													if v1861 != 0 {
														return int32(0)
													} else {
														F_sequence_close(m, v1703, int32(1))
														mBase = m.M
														v1864 = m.ExcPending
														if v1864 != 0 {
															return int32(0)
														} else {
															v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
															v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
															m.G0 = v11 + int32(1424)
															if v1872 != 0 {
																v1878 = v1873
															} else {
																v1878 = int32(0)
															}
															return v1878
														}
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v1728 = m.ExcPending
													if v1728 != 0 {
														return int32(0)
													} else {
														v1729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v11)+896)) = v1729
														F_errmsg_internal(m, int32(53299), v11+int32(896))
														mBase = m.M
														v1735 = m.ExcPending
														if v1735 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(472250), int32(3787), int32(235362))
															mBase = m.M
															v1740 = m.ExcPending
															if v1740 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v1741 = *(*int32)(unsafe.Add(mBase, uint32(v1721)+16))
												v1742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1741)+22)))
												v1743 = v1741 + v1742
												v1744 = *(*int32)(unsafe.Add(mBase, uint32(v1743)+4))
												v1746 = F_GetUserNameFromId(m, v1744, int32(0))
												mBase = m.M
												v1747 = m.ExcPending
												if v1747 != 0 {
													return int32(0)
												} else {
													v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1743)+8))
													if v1748 != 0 {
														v1749 = F_get_namespace_name(m, v1748)
														mBase = m.M
														v1750 = m.ExcPending
														if v1750 != 0 {
															return int32(0)
														} else {
															v1751 = v1749
															v1752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1743)+12)))
															switch v1752 - int32(76) {
															case 0:
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1088)) = v1746
																F_appendStringInfo(m, v11+int32(1408), int32(184860), v11+int32(1088))
																mBase = m.M
																v1838 = m.ExcPending
																if v1838 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1719)
																	mBase = m.M
																	v1861 = m.ExcPending
																	if v1861 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v1703, int32(1))
																		mBase = m.M
																		v1864 = m.ExcPending
																		if v1864 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															default:
																if v1751 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+932)) = v1751
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+928)) = v1746
																	F_appendStringInfo(m, v11+int32(1408), int32(188017), v11+int32(928))
																	mBase = m.M
																	v1847 = m.ExcPending
																	if v1847 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1719)
																		mBase = m.M
																		v1861 = m.ExcPending
																		if v1861 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v1703, int32(1))
																			mBase = m.M
																			v1864 = m.ExcPending
																			if v1864 != 0 {
																				return int32(0)
																			} else {
																				v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																				v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																				m.G0 = v11 + int32(1424)
																				if v1872 != 0 {
																					v1878 = v1873
																				} else {
																					v1878 = int32(0)
																				}
																				return v1878
																			}
																		}
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+912)) = v1746
																	F_appendStringInfo(m, v11+int32(1408), int32(185088), v11+int32(912))
																	mBase = m.M
																	v1855 = m.ExcPending
																	if v1855 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1719)
																		mBase = m.M
																		v1861 = m.ExcPending
																		if v1861 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v1703, int32(1))
																			mBase = m.M
																			v1864 = m.ExcPending
																			if v1864 != 0 {
																				return int32(0)
																			} else {
																				v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																				v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																				m.G0 = v11 + int32(1424)
																				if v1872 != 0 {
																					v1878 = v1873
																				} else {
																					v1878 = int32(0)
																				}
																				return v1878
																			}
																		}
																	}
																}
															case 7:
																if v1751 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+996)) = v1751
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+992)) = v1746
																	F_appendStringInfo(m, v11+int32(1408), int32(188070), v11+int32(992))
																	mBase = m.M
																	v1780 = m.ExcPending
																	if v1780 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1719)
																		mBase = m.M
																		v1861 = m.ExcPending
																		if v1861 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v1703, int32(1))
																			mBase = m.M
																			v1864 = m.ExcPending
																			if v1864 != 0 {
																				return int32(0)
																			} else {
																				v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																				v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																				m.G0 = v11 + int32(1424)
																				if v1872 != 0 {
																					v1878 = v1873
																				} else {
																					v1878 = int32(0)
																				}
																				return v1878
																			}
																		}
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+976)) = v1746
																	F_appendStringInfo(m, v11+int32(1408), int32(185128), v11+int32(976))
																	mBase = m.M
																	v1788 = m.ExcPending
																	if v1788 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1719)
																		mBase = m.M
																		v1861 = m.ExcPending
																		if v1861 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v1703, int32(1))
																			mBase = m.M
																			v1864 = m.ExcPending
																			if v1864 != 0 {
																				return int32(0)
																			} else {
																				v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																				v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																				m.G0 = v11 + int32(1424)
																				if v1872 != 0 {
																					v1878 = v1873
																				} else {
																					v1878 = int32(0)
																				}
																				return v1878
																			}
																		}
																	}
																}
															case 8:
																if v1751 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+1060)) = v1751
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+1056)) = v1746
																	F_appendStringInfo(m, v11+int32(1408), int32(187951), v11+int32(1056))
																	mBase = m.M
																	v1814 = m.ExcPending
																	if v1814 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1719)
																		mBase = m.M
																		v1861 = m.ExcPending
																		if v1861 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v1703, int32(1))
																			mBase = m.M
																			v1864 = m.ExcPending
																			if v1864 != 0 {
																				return int32(0)
																			} else {
																				v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																				v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																				m.G0 = v11 + int32(1424)
																				if v1872 != 0 {
																					v1878 = v1873
																				} else {
																					v1878 = int32(0)
																				}
																				return v1878
																			}
																		}
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+1040)) = v1746
																	F_appendStringInfo(m, v11+int32(1408), int32(185035), v11+int32(1040))
																	mBase = m.M
																	v1822 = m.ExcPending
																	if v1822 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1719)
																		mBase = m.M
																		v1861 = m.ExcPending
																		if v1861 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v1703, int32(1))
																			mBase = m.M
																			v1864 = m.ExcPending
																			if v1864 != 0 {
																				return int32(0)
																			} else {
																				v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																				v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																				m.G0 = v11 + int32(1424)
																				if v1872 != 0 {
																					v1878 = v1873
																				} else {
																					v1878 = int32(0)
																				}
																				return v1878
																			}
																		}
																	}
																}
															case 26:
																if v1751 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+1028)) = v1751
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+1024)) = v1746
																	F_appendStringInfo(m, v11+int32(1408), int32(187811), v11+int32(1024))
																	mBase = m.M
																	v1797 = m.ExcPending
																	if v1797 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1719)
																		mBase = m.M
																		v1861 = m.ExcPending
																		if v1861 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v1703, int32(1))
																			mBase = m.M
																			v1864 = m.ExcPending
																			if v1864 != 0 {
																				return int32(0)
																			} else {
																				v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																				v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																				m.G0 = v11 + int32(1424)
																				if v1872 != 0 {
																					v1878 = v1873
																				} else {
																					v1878 = int32(0)
																				}
																				return v1878
																			}
																		}
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+1008)) = v1746
																	F_appendStringInfo(m, v11+int32(1408), int32(184921), v11+int32(1008))
																	mBase = m.M
																	v1805 = m.ExcPending
																	if v1805 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1719)
																		mBase = m.M
																		v1861 = m.ExcPending
																		if v1861 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v1703, int32(1))
																			mBase = m.M
																			v1864 = m.ExcPending
																			if v1864 != 0 {
																				return int32(0)
																			} else {
																				v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																				v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																				m.G0 = v11 + int32(1424)
																				if v1872 != 0 {
																					v1878 = v1873
																				} else {
																					v1878 = int32(0)
																				}
																				return v1878
																			}
																		}
																	}
																}
															case 34:
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1072)) = v1746
																F_appendStringInfo(m, v11+int32(1408), int32(185185), v11+int32(1072))
																mBase = m.M
																v1830 = m.ExcPending
																if v1830 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1719)
																	mBase = m.M
																	v1861 = m.ExcPending
																	if v1861 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v1703, int32(1))
																		mBase = m.M
																		v1864 = m.ExcPending
																		if v1864 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															case 38:
																if v1751 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+964)) = v1751
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+960)) = v1746
																	F_appendStringInfo(m, v11+int32(1408), int32(187881), v11+int32(960))
																	mBase = m.M
																	v1763 = m.ExcPending
																	if v1763 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1719)
																		mBase = m.M
																		v1861 = m.ExcPending
																		if v1861 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v1703, int32(1))
																			mBase = m.M
																			v1864 = m.ExcPending
																			if v1864 != 0 {
																				return int32(0)
																			} else {
																				v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																				v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																				m.G0 = v11 + int32(1424)
																				if v1872 != 0 {
																					v1878 = v1873
																				} else {
																					v1878 = int32(0)
																				}
																				return v1878
																			}
																		}
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+944)) = v1746
																	F_appendStringInfo(m, v11+int32(1408), int32(184978), v11+int32(944))
																	mBase = m.M
																	v1771 = m.ExcPending
																	if v1771 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1719)
																		mBase = m.M
																		v1861 = m.ExcPending
																		if v1861 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v1703, int32(1))
																			mBase = m.M
																			v1864 = m.ExcPending
																			if v1864 != 0 {
																				return int32(0)
																			} else {
																				v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																				v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																				m.G0 = v11 + int32(1424)
																				if v1872 != 0 {
																					v1878 = v1873
																				} else {
																					v1878 = int32(0)
																				}
																				return v1878
																			}
																		}
																	}
																}
															}
														}
													} else {
														v1751 = int32(0)
														v1752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1743)+12)))
														switch v1752 - int32(76) {
														case 0:
															*(*int32)(unsafe.Add(mBase, uint32(v11)+1088)) = v1746
															F_appendStringInfo(m, v11+int32(1408), int32(184860), v11+int32(1088))
															mBase = m.M
															v1838 = m.ExcPending
															if v1838 != 0 {
																return int32(0)
															} else {
																F_systable_endscan(m, v1719)
																mBase = m.M
																v1861 = m.ExcPending
																if v1861 != 0 {
																	return int32(0)
																} else {
																	F_sequence_close(m, v1703, int32(1))
																	mBase = m.M
																	v1864 = m.ExcPending
																	if v1864 != 0 {
																		return int32(0)
																	} else {
																		v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																		v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																		m.G0 = v11 + int32(1424)
																		if v1872 != 0 {
																			v1878 = v1873
																		} else {
																			v1878 = int32(0)
																		}
																		return v1878
																	}
																}
															}
														default:
															if v1751 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+932)) = v1751
																*(*int32)(unsafe.Add(mBase, uint32(v11)+928)) = v1746
																F_appendStringInfo(m, v11+int32(1408), int32(188017), v11+int32(928))
																mBase = m.M
																v1847 = m.ExcPending
																if v1847 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1719)
																	mBase = m.M
																	v1861 = m.ExcPending
																	if v1861 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v1703, int32(1))
																		mBase = m.M
																		v1864 = m.ExcPending
																		if v1864 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+912)) = v1746
																F_appendStringInfo(m, v11+int32(1408), int32(185088), v11+int32(912))
																mBase = m.M
																v1855 = m.ExcPending
																if v1855 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1719)
																	mBase = m.M
																	v1861 = m.ExcPending
																	if v1861 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v1703, int32(1))
																		mBase = m.M
																		v1864 = m.ExcPending
																		if v1864 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															}
														case 7:
															if v1751 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+996)) = v1751
																*(*int32)(unsafe.Add(mBase, uint32(v11)+992)) = v1746
																F_appendStringInfo(m, v11+int32(1408), int32(188070), v11+int32(992))
																mBase = m.M
																v1780 = m.ExcPending
																if v1780 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1719)
																	mBase = m.M
																	v1861 = m.ExcPending
																	if v1861 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v1703, int32(1))
																		mBase = m.M
																		v1864 = m.ExcPending
																		if v1864 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+976)) = v1746
																F_appendStringInfo(m, v11+int32(1408), int32(185128), v11+int32(976))
																mBase = m.M
																v1788 = m.ExcPending
																if v1788 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1719)
																	mBase = m.M
																	v1861 = m.ExcPending
																	if v1861 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v1703, int32(1))
																		mBase = m.M
																		v1864 = m.ExcPending
																		if v1864 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															}
														case 8:
															if v1751 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1060)) = v1751
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1056)) = v1746
																F_appendStringInfo(m, v11+int32(1408), int32(187951), v11+int32(1056))
																mBase = m.M
																v1814 = m.ExcPending
																if v1814 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1719)
																	mBase = m.M
																	v1861 = m.ExcPending
																	if v1861 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v1703, int32(1))
																		mBase = m.M
																		v1864 = m.ExcPending
																		if v1864 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1040)) = v1746
																F_appendStringInfo(m, v11+int32(1408), int32(185035), v11+int32(1040))
																mBase = m.M
																v1822 = m.ExcPending
																if v1822 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1719)
																	mBase = m.M
																	v1861 = m.ExcPending
																	if v1861 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v1703, int32(1))
																		mBase = m.M
																		v1864 = m.ExcPending
																		if v1864 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															}
														case 26:
															if v1751 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1028)) = v1751
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1024)) = v1746
																F_appendStringInfo(m, v11+int32(1408), int32(187811), v11+int32(1024))
																mBase = m.M
																v1797 = m.ExcPending
																if v1797 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1719)
																	mBase = m.M
																	v1861 = m.ExcPending
																	if v1861 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v1703, int32(1))
																		mBase = m.M
																		v1864 = m.ExcPending
																		if v1864 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1008)) = v1746
																F_appendStringInfo(m, v11+int32(1408), int32(184921), v11+int32(1008))
																mBase = m.M
																v1805 = m.ExcPending
																if v1805 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1719)
																	mBase = m.M
																	v1861 = m.ExcPending
																	if v1861 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v1703, int32(1))
																		mBase = m.M
																		v1864 = m.ExcPending
																		if v1864 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															}
														case 34:
															*(*int32)(unsafe.Add(mBase, uint32(v11)+1072)) = v1746
															F_appendStringInfo(m, v11+int32(1408), int32(185185), v11+int32(1072))
															mBase = m.M
															v1830 = m.ExcPending
															if v1830 != 0 {
																return int32(0)
															} else {
																F_systable_endscan(m, v1719)
																mBase = m.M
																v1861 = m.ExcPending
																if v1861 != 0 {
																	return int32(0)
																} else {
																	F_sequence_close(m, v1703, int32(1))
																	mBase = m.M
																	v1864 = m.ExcPending
																	if v1864 != 0 {
																		return int32(0)
																	} else {
																		v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																		v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																		m.G0 = v11 + int32(1424)
																		if v1872 != 0 {
																			v1878 = v1873
																		} else {
																			v1878 = int32(0)
																		}
																		return v1878
																	}
																}
															}
														case 38:
															if v1751 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+964)) = v1751
																*(*int32)(unsafe.Add(mBase, uint32(v11)+960)) = v1746
																F_appendStringInfo(m, v11+int32(1408), int32(187881), v11+int32(960))
																mBase = m.M
																v1763 = m.ExcPending
																if v1763 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1719)
																	mBase = m.M
																	v1861 = m.ExcPending
																	if v1861 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v1703, int32(1))
																		mBase = m.M
																		v1864 = m.ExcPending
																		if v1864 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+944)) = v1746
																F_appendStringInfo(m, v11+int32(1408), int32(184978), v11+int32(944))
																mBase = m.M
																v1771 = m.ExcPending
																if v1771 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1719)
																	mBase = m.M
																	v1861 = m.ExcPending
																	if v1861 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v1703, int32(1))
																		mBase = m.M
																		v1864 = m.ExcPending
																		if v1864 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v1669 = m.ExcPending
							if v1669 != 0 {
								return int32(0)
							} else {
								v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1670
								F_errmsg_internal(m, int32(55485), v11)
								mBase = m.M
								v1674 = m.ExcPending
								if v1674 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(472250), int32(4072), int32(235362))
									mBase = m.M
									v1679 = m.ExcPending
									if v1679 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			} else {
				switch v19 - int32(2601) {
				case 0:
					v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v523 = F_SearchSysCache1(m, int32(2), v522)
					mBase = m.M
					v524 = m.ExcPending
					if v524 != 0 {
						return int32(0)
					} else {
						if v523 == int32(0) {
							if l1 != 0 {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v530 = m.ExcPending
								if v530 != 0 {
									return int32(0)
								} else {
									v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+336)) = v531
									F_errmsg_internal(m, int32(50892), v11+int32(336))
									mBase = m.M
									v537 = m.ExcPending
									if v537 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(472250), int32(3219), int32(235362))
										mBase = m.M
										v542 = m.ExcPending
										if v542 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v543 = *(*int32)(unsafe.Add(mBase, uint32(v523)+16))
							v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543)+22)))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+352)) = v543 + v544 + int32(4)
							F_appendStringInfo(m, v11+int32(1408), int32(186647), v11+int32(352))
							mBase = m.M
							v555 = m.ExcPending
							if v555 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v523)
								mBase = m.M
								v557 = m.ExcPending
								if v557 != 0 {
									return int32(0)
								} else {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								}
							}
						}
					}
				case 1:
					v560 = F_table_open(m, int32(2602), int32(1))
					mBase = m.M
					v561 = m.ExcPending
					if v561 != 0 {
						return int32(0)
					} else {
						v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_ScanKeyInit(m, v11+int32(1360), int32(1), int32(3), int32(184), v567)
						mBase = m.M
						v569 = m.ExcPending
						if v569 != 0 {
							return int32(0)
						} else {
							v571 = int32(1)
							v576 = F_systable_beginscan(m, v560, int32(2756), v571, int32(0), v571, v11+int32(1360))
							mBase = m.M
							v577 = m.ExcPending
							if v577 != 0 {
								return int32(0)
							} else {
								v578 = F_systable_getnext(m, v576)
								mBase = m.M
								v579 = m.ExcPending
								if v579 != 0 {
									return int32(0)
								} else {
									if v578 == int32(0) {
										if l1 != 0 {
											F_systable_endscan(m, v576)
											mBase = m.M
											v647 = m.ExcPending
											if v647 != 0 {
												return int32(0)
											} else {
												F_sequence_close(m, v560, int32(1))
												mBase = m.M
												v650 = m.ExcPending
												if v650 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v585 = m.ExcPending
											if v585 != 0 {
												return int32(0)
											} else {
												v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+368)) = v586
												F_errmsg_internal(m, int32(37158), v11+int32(368))
												mBase = m.M
												v592 = m.ExcPending
												if v592 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(472250), int32(3255), int32(235362))
													mBase = m.M
													v597 = m.ExcPending
													if v597 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v598 = *(*int32)(unsafe.Add(mBase, uint32(v578)+16))
										v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+22)))
										F_initStringInfo(m, v11+int32(1344))
										mBase = m.M
										v603 = m.ExcPending
										if v603 != 0 {
											return int32(0)
										} else {
											v606 = v599 + v598
											v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)+4))
											F_getOpFamilyDescription(m, v11+int32(1344), v607, int32(0))
											mBase = m.M
											v610 = m.ExcPending
											if v610 != 0 {
												return int32(0)
											} else {
												v611 = int32(*(*int16)(unsafe.Add(mBase, uint32(v606)+16)))
												v612 = *(*int32)(unsafe.Add(mBase, uint32(v606)+8))
												v615 = F_format_type_extended(m, v612, int32(-1), int32(2))
												mBase = m.M
												v616 = m.ExcPending
												if v616 != 0 {
													return int32(0)
												} else {
													v617 = *(*int32)(unsafe.Add(mBase, uint32(v606)+12))
													v620 = F_format_type_extended(m, v617, int32(-1), int32(2))
													mBase = m.M
													v621 = m.ExcPending
													if v621 != 0 {
														return int32(0)
													} else {
														v622 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
														v623 = *(*int32)(unsafe.Add(mBase, uint32(v606)+20))
														v624 = F_format_operator(m, v623)
														mBase = m.M
														v625 = m.ExcPending
														if v625 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+400)) = v624
															*(*int32)(unsafe.Add(mBase, uint32(v11)+396)) = v622
															*(*int32)(unsafe.Add(mBase, uint32(v11)+392)) = v620
															*(*int32)(unsafe.Add(mBase, uint32(v11)+388)) = v615
															*(*int32)(unsafe.Add(mBase, uint32(v11)+384)) = v611
															F_appendStringInfo(m, v11+int32(1408), int32(190888), v11+int32(384))
															mBase = m.M
															v637 = m.ExcPending
															if v637 != 0 {
																return int32(0)
															} else {
																v638 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
																F_pfree(m, v638)
																mBase = m.M
																v640 = m.ExcPending
																if v640 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v576)
																	mBase = m.M
																	v647 = m.ExcPending
																	if v647 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v560, int32(1))
																		mBase = m.M
																		v650 = m.ExcPending
																		if v650 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				case 2:
					v653 = F_table_open(m, int32(2603), int32(1))
					mBase = m.M
					v654 = m.ExcPending
					if v654 != 0 {
						return int32(0)
					} else {
						v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_ScanKeyInit(m, v11+int32(1360), int32(1), int32(3), int32(184), v660)
						mBase = m.M
						v662 = m.ExcPending
						if v662 != 0 {
							return int32(0)
						} else {
							v664 = int32(1)
							v669 = F_systable_beginscan(m, v653, int32(2757), v664, int32(0), v664, v11+int32(1360))
							mBase = m.M
							v670 = m.ExcPending
							if v670 != 0 {
								return int32(0)
							} else {
								v671 = F_systable_getnext(m, v669)
								mBase = m.M
								v672 = m.ExcPending
								if v672 != 0 {
									return int32(0)
								} else {
									if v671 == int32(0) {
										if l1 != 0 {
											F_systable_endscan(m, v669)
											mBase = m.M
											v740 = m.ExcPending
											if v740 != 0 {
												return int32(0)
											} else {
												F_sequence_close(m, v653, int32(1))
												mBase = m.M
												v743 = m.ExcPending
												if v743 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v678 = m.ExcPending
											if v678 != 0 {
												return int32(0)
											} else {
												v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+416)) = v679
												F_errmsg_internal(m, int32(37197), v11+int32(416))
												mBase = m.M
												v685 = m.ExcPending
												if v685 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(472250), int32(3320), int32(235362))
													mBase = m.M
													v690 = m.ExcPending
													if v690 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v691 = *(*int32)(unsafe.Add(mBase, uint32(v671)+16))
										v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691)+22)))
										F_initStringInfo(m, v11+int32(1344))
										mBase = m.M
										v696 = m.ExcPending
										if v696 != 0 {
											return int32(0)
										} else {
											v699 = v692 + v691
											v700 = *(*int32)(unsafe.Add(mBase, uint32(v699)+4))
											F_getOpFamilyDescription(m, v11+int32(1344), v700, int32(0))
											mBase = m.M
											v703 = m.ExcPending
											if v703 != 0 {
												return int32(0)
											} else {
												v704 = int32(*(*int16)(unsafe.Add(mBase, uint32(v699)+16)))
												v705 = *(*int32)(unsafe.Add(mBase, uint32(v699)+8))
												v708 = F_format_type_extended(m, v705, int32(-1), int32(2))
												mBase = m.M
												v709 = m.ExcPending
												if v709 != 0 {
													return int32(0)
												} else {
													v710 = *(*int32)(unsafe.Add(mBase, uint32(v699)+12))
													v713 = F_format_type_extended(m, v710, int32(-1), int32(2))
													mBase = m.M
													v714 = m.ExcPending
													if v714 != 0 {
														return int32(0)
													} else {
														v715 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
														v716 = *(*int32)(unsafe.Add(mBase, uint32(v699)+20))
														v717 = F_format_procedure(m, v716)
														mBase = m.M
														v718 = m.ExcPending
														if v718 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+448)) = v717
															*(*int32)(unsafe.Add(mBase, uint32(v11)+444)) = v715
															*(*int32)(unsafe.Add(mBase, uint32(v11)+440)) = v713
															*(*int32)(unsafe.Add(mBase, uint32(v11)+436)) = v708
															*(*int32)(unsafe.Add(mBase, uint32(v11)+432)) = v704
															F_appendStringInfo(m, v11+int32(1408), int32(190919), v11+int32(432))
															mBase = m.M
															v730 = m.ExcPending
															if v730 != 0 {
																return int32(0)
															} else {
																v731 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
																F_pfree(m, v731)
																mBase = m.M
																v733 = m.ExcPending
																if v733 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v669)
																	mBase = m.M
																	v740 = m.ExcPending
																	if v740 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v653, int32(1))
																		mBase = m.M
																		v743 = m.ExcPending
																		if v743 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				case 3:
					v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_GetAttrDefaultColumnAddress(m, v11+int32(1360), v380)
					mBase = m.M
					v382 = m.ExcPending
					if v382 != 0 {
						return int32(0)
					} else {
						v383 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1364))
						if v383 == int32(0) {
							if l1 != 0 {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v389 = m.ExcPending
								if v389 != 0 {
									return int32(0)
								} else {
									v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v390
									F_errmsg_internal(m, int32(46743), v11+int32(208))
									mBase = m.M
									v396 = m.ExcPending
									if v396 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(472250), int32(3121), int32(235362))
										mBase = m.M
										v401 = m.ExcPending
										if v401 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v405 = F_getObjectDescription(m, v11+int32(1360), int32(0))
							mBase = m.M
							v406 = m.ExcPending
							if v406 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+224)) = v405
								F_appendStringInfo(m, v11+int32(1408), int32(171890), v11+int32(224))
								mBase = m.M
								v414 = m.ExcPending
								if v414 != 0 {
									return int32(0)
								} else {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								}
							}
						}
					}
				case 4:
					v155 = F_table_open(m, int32(2605), int32(1))
					mBase = m.M
					v156 = m.ExcPending
					if v156 != 0 {
						return int32(0)
					} else {
						v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_ScanKeyInit(m, v11+int32(1360), int32(1), int32(3), int32(184), v162)
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
							return int32(0)
						} else {
							v166 = int32(1)
							v171 = F_systable_beginscan(m, v155, int32(2660), v166, int32(0), v166, v11+int32(1360))
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return int32(0)
							} else {
								v173 = F_systable_getnext(m, v171)
								mBase = m.M
								v174 = m.ExcPending
								if v174 != 0 {
									return int32(0)
								} else {
									if v173 == int32(0) {
										if l1 != 0 {
											F_systable_endscan(m, v171)
											mBase = m.M
											v214 = m.ExcPending
											if v214 != 0 {
												return int32(0)
											} else {
												F_sequence_close(m, v155, int32(1))
												mBase = m.M
												v217 = m.ExcPending
												if v217 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v180 = m.ExcPending
											if v180 != 0 {
												return int32(0)
											} else {
												v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v181
												F_errmsg_internal(m, int32(38152), v11-int32(-64))
												mBase = m.M
												v187 = m.ExcPending
												if v187 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(472250), int32(2993), int32(235362))
													mBase = m.M
													v192 = m.ExcPending
													if v192 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v193 = *(*int32)(unsafe.Add(mBase, uint32(v173)+16))
										v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+22)))
										v195 = v193 + v194
										v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
										v197 = F_format_type_be(m, v196)
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return int32(0)
										} else {
											v199 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
											v200 = F_format_type_be(m, v199)
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v200
												*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v197
												F_appendStringInfo(m, v11+int32(1408), int32(173127), v11+int32(80))
												mBase = m.M
												v210 = m.ExcPending
												if v210 != 0 {
													return int32(0)
												} else {
													F_systable_endscan(m, v171)
													mBase = m.M
													v214 = m.ExcPending
													if v214 != 0 {
														return int32(0)
													} else {
														F_sequence_close(m, v155, int32(1))
														mBase = m.M
														v217 = m.ExcPending
														if v217 != 0 {
															return int32(0)
														} else {
															v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
															v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
															m.G0 = v11 + int32(1424)
															if v1872 != 0 {
																v1878 = v1873
															} else {
																v1878 = int32(0)
															}
															return v1878
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				case 5:
					v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v267 = F_SearchSysCache1(m, int32(19), v266)
					mBase = m.M
					v268 = m.ExcPending
					if v268 != 0 {
						return int32(0)
					} else {
						if v267 == int32(0) {
							if l1 != 0 {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v274 = m.ExcPending
								if v274 != 0 {
									return int32(0)
								} else {
									v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v275
									F_errmsg_internal(m, int32(38745), v11+int32(128))
									mBase = m.M
									v281 = m.ExcPending
									if v281 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(472250), int32(3053), int32(235362))
										mBase = m.M
										v286 = m.ExcPending
										if v286 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v287 = *(*int32)(unsafe.Add(mBase, uint32(v267)+16))
							v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+22)))
							v289 = v287 + v288
							v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+80))
							if v290 != 0 {
								F_initStringInfo(m, v11+int32(1360))
								mBase = m.M
								v294 = m.ExcPending
								if v294 != 0 {
									return int32(0)
								} else {
									v297 = *(*int32)(unsafe.Add(mBase, uint32(v289)+80))
									F_getRelationDescription(m, v11+int32(1360), v297, int32(0))
									mBase = m.M
									v300 = m.ExcPending
									if v300 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v289 + int32(4)
										v304 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = v304
										F_appendStringInfo(m, v11+int32(1408), int32(174793), v11+int32(160))
										mBase = m.M
										v312 = m.ExcPending
										if v312 != 0 {
											return int32(0)
										} else {
											v313 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
											F_pfree(m, v313)
											mBase = m.M
											v315 = m.ExcPending
											if v315 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v267)
												mBase = m.M
												v317 = m.ExcPending
												if v317 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = v289 + int32(4)
								F_appendStringInfo(m, v11+int32(1408), int32(169185), v11+int32(144))
								mBase = m.M
								v327 = m.ExcPending
								if v327 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v267)
									mBase = m.M
									v329 = m.ExcPending
									if v329 != 0 {
										return int32(0)
									} else {
										v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
										v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
										m.G0 = v11 + int32(1424)
										if v1872 != 0 {
											v1878 = v1873
										} else {
											v1878 = int32(0)
										}
										return v1878
									}
								}
							}
						}
					}
				case 6:
					v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v332 = F_SearchSysCache1(m, int32(20), v331)
					mBase = m.M
					v333 = m.ExcPending
					if v333 != 0 {
						return int32(0)
					} else {
						if v332 == int32(0) {
							if l1 != 0 {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v339 = m.ExcPending
								if v339 != 0 {
									return int32(0)
								} else {
									v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = v340
									F_errmsg_internal(m, int32(44498), v11+int32(176))
									mBase = m.M
									v346 = m.ExcPending
									if v346 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(472250), int32(3092), int32(235362))
										mBase = m.M
										v351 = m.ExcPending
										if v351 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v352 = *(*int32)(unsafe.Add(mBase, uint32(v332)+16))
							v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+22)))
							v354 = v352 + v353
							v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v357 = F_ConversionIsVisibleExt(m, v355, int32(0))
							mBase = m.M
							v358 = m.ExcPending
							if v358 != 0 {
								return int32(0)
							} else {
								if v357 != 0 {
									v363 = int32(0)
									v366 = F_quote_qualified_identifier(m, v363, v354+int32(4))
									mBase = m.M
									v367 = m.ExcPending
									if v367 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v366
										F_appendStringInfo(m, v11+int32(1408), int32(174568), v11+int32(192))
										mBase = m.M
										v375 = m.ExcPending
										if v375 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v332)
											mBase = m.M
											v377 = m.ExcPending
											if v377 != 0 {
												return int32(0)
											} else {
												v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
												v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
												m.G0 = v11 + int32(1424)
												if v1872 != 0 {
													v1878 = v1873
												} else {
													v1878 = int32(0)
												}
												return v1878
											}
										}
									}
								} else {
									v360 = *(*int32)(unsafe.Add(mBase, uint32(v354)+68))
									v361 = F_get_namespace_name(m, v360)
									mBase = m.M
									v362 = m.ExcPending
									if v362 != 0 {
										return int32(0)
									} else {
										v363 = v361
										v366 = F_quote_qualified_identifier(m, v363, v354+int32(4))
										mBase = m.M
										v367 = m.ExcPending
										if v367 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v366
											F_appendStringInfo(m, v11+int32(1408), int32(174568), v11+int32(192))
											mBase = m.M
											v375 = m.ExcPending
											if v375 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v332)
												mBase = m.M
												v377 = m.ExcPending
												if v377 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										}
									}
								}
							}
						}
					}
				case 7, 8, 9, 10, 13, 18:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v1669 = m.ExcPending
					if v1669 != 0 {
						return int32(0)
					} else {
						v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1670
						F_errmsg_internal(m, int32(55485), v11)
						mBase = m.M
						v1674 = m.ExcPending
						if v1674 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(472250), int32(4072), int32(235362))
							mBase = m.M
							v1679 = m.ExcPending
							if v1679 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 11:
					v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v416 = F_get_language_name(m, v415, l1)
					mBase = m.M
					v417 = m.ExcPending
					if v417 != 0 {
						return int32(0)
					} else {
						if v416 == int32(0) {
							v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
							v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
							m.G0 = v11 + int32(1424)
							if v1872 != 0 {
								v1878 = v1873
							} else {
								v1878 = int32(0)
							}
							return v1878
						} else {
							v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v422 = F_get_language_name(m, v420, int32(0))
							mBase = m.M
							v423 = m.ExcPending
							if v423 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+240)) = v422
								F_appendStringInfo(m, v11+int32(1408), int32(185854), v11+int32(240))
								mBase = m.M
								v431 = m.ExcPending
								if v431 != 0 {
									return int32(0)
								} else {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								}
							}
						}
					}
				case 12:
					v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v433 = F_LargeObjectExists(m, v432)
					mBase = m.M
					v434 = m.ExcPending
					if v434 != 0 {
						return int32(0)
					} else {
						if v433 == int32(0) {
							v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
							v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
							m.G0 = v11 + int32(1424)
							if v1872 != 0 {
								v1878 = v1873
							} else {
								v1878 = int32(0)
							}
							return v1878
						} else {
							v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+256)) = v437
							F_appendStringInfo(m, v11+int32(1408), int32(39802), v11+int32(256))
							mBase = m.M
							v445 = m.ExcPending
							if v445 != 0 {
								return int32(0)
							} else {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							}
						}
					}
				case 14:
					v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v895 = F_get_namespace_name(m, v894)
					mBase = m.M
					v896 = m.ExcPending
					if v896 != 0 {
						return int32(0)
					} else {
						if v895 == int32(0) {
							if l1 != 0 {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v902 = m.ExcPending
								if v902 != 0 {
									return int32(0)
								} else {
									v903 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+528)) = v903
									F_errmsg_internal(m, int32(50496), v11+int32(528))
									mBase = m.M
									v909 = m.ExcPending
									if v909 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(472250), int32(3460), int32(235362))
										mBase = m.M
										v914 = m.ExcPending
										if v914 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+544)) = v895
							F_appendStringInfo(m, v11+int32(1408), int32(188157), v11+int32(544))
							mBase = m.M
							v922 = m.ExcPending
							if v922 != 0 {
								return int32(0)
							} else {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							}
						}
					}
				case 15:
					v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v462 = F_SearchSysCache1(m, int32(14), v461)
					mBase = m.M
					v463 = m.ExcPending
					if v463 != 0 {
						return int32(0)
					} else {
						if v462 == int32(0) {
							if l1 != 0 {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v469 = m.ExcPending
								if v469 != 0 {
									return int32(0)
								} else {
									v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+288)) = v470
									F_errmsg_internal(m, int32(39996), v11+int32(288))
									mBase = m.M
									v476 = m.ExcPending
									if v476 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(472250), int32(3176), int32(235362))
										mBase = m.M
										v481 = m.ExcPending
										if v481 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v483 = *(*int32)(unsafe.Add(mBase, uint32(v462)+16))
							v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+22)))
							v485 = v483 + v484
							v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)+4))
							v487 = F_SearchSysCache1(m, int32(2), v486)
							mBase = m.M
							v488 = m.ExcPending
							if v488 != 0 {
								return int32(0)
							} else {
								if v487 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1688 = m.ExcPending
									if v1688 != 0 {
										return int32(0)
									} else {
										v1689 = *(*int32)(unsafe.Add(mBase, uint32(v485)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+304)) = v1689
										F_errmsg_internal(m, int32(50892), v11+int32(304))
										mBase = m.M
										v1695 = m.ExcPending
										if v1695 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(472250), int32(3186), int32(235362))
											mBase = m.M
											v1700 = m.ExcPending
											if v1700 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v491 = *(*int32)(unsafe.Add(mBase, uint32(v487)+16))
									v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+22)))
									v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v495 = F_OpclassIsVisible(m, v494)
									mBase = m.M
									v496 = m.ExcPending
									if v496 != 0 {
										return int32(0)
									} else {
										if v495 != 0 {
											v501 = int32(0)
											v504 = F_quote_qualified_identifier(m, v501, v485+int32(8))
											mBase = m.M
											v505 = m.ExcPending
											if v505 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+324)) = v491 + v492 + int32(4)
												*(*int32)(unsafe.Add(mBase, uint32(v11)+320)) = v504
												F_appendStringInfo(m, v11+int32(1408), int32(186566), v11+int32(320))
												mBase = m.M
												v516 = m.ExcPending
												if v516 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v487)
													mBase = m.M
													v518 = m.ExcPending
													if v518 != 0 {
														return int32(0)
													} else {
														F_ReleaseCatCache(m, v462)
														mBase = m.M
														v520 = m.ExcPending
														if v520 != 0 {
															return int32(0)
														} else {
															v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
															v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
															m.G0 = v11 + int32(1424)
															if v1872 != 0 {
																v1878 = v1873
															} else {
																v1878 = int32(0)
															}
															return v1878
														}
													}
												}
											}
										} else {
											v498 = *(*int32)(unsafe.Add(mBase, uint32(v485)+72))
											v499 = F_get_namespace_name(m, v498)
											mBase = m.M
											v500 = m.ExcPending
											if v500 != 0 {
												return int32(0)
											} else {
												v501 = v499
												v504 = F_quote_qualified_identifier(m, v501, v485+int32(8))
												mBase = m.M
												v505 = m.ExcPending
												if v505 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+324)) = v491 + v492 + int32(4)
													*(*int32)(unsafe.Add(mBase, uint32(v11)+320)) = v504
													F_appendStringInfo(m, v11+int32(1408), int32(186566), v11+int32(320))
													mBase = m.M
													v516 = m.ExcPending
													if v516 != 0 {
														return int32(0)
													} else {
														F_ReleaseCatCache(m, v487)
														mBase = m.M
														v518 = m.ExcPending
														if v518 != 0 {
															return int32(0)
														} else {
															F_ReleaseCatCache(m, v462)
															mBase = m.M
															v520 = m.ExcPending
															if v520 != 0 {
																return int32(0)
															} else {
																v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																m.G0 = v11 + int32(1424)
																if v1872 != 0 {
																	v1878 = v1873
																} else {
																	v1878 = int32(0)
																}
																return v1878
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				case 16:
					v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v448 = F_format_operator_extended(m, v446, int32(1))
					mBase = m.M
					v449 = m.ExcPending
					if v449 != 0 {
						return int32(0)
					} else {
						if v448 == int32(0) {
							v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
							v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
							m.G0 = v11 + int32(1424)
							if v1872 != 0 {
								v1878 = v1873
							} else {
								v1878 = int32(0)
							}
							return v1878
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+272)) = v448
							F_appendStringInfo(m, v11+int32(1408), int32(171724), v11+int32(272))
							mBase = m.M
							v459 = m.ExcPending
							if v459 != 0 {
								return int32(0)
							} else {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							}
						}
					}
				case 17:
					v746 = F_table_open(m, int32(2618), int32(1))
					mBase = m.M
					v747 = m.ExcPending
					if v747 != 0 {
						return int32(0)
					} else {
						v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_ScanKeyInit(m, v11+int32(1360), int32(1), int32(3), int32(184), v753)
						mBase = m.M
						v755 = m.ExcPending
						if v755 != 0 {
							return int32(0)
						} else {
							v757 = int32(1)
							v762 = F_systable_beginscan(m, v746, int32(2692), v757, int32(0), v757, v11+int32(1360))
							mBase = m.M
							v763 = m.ExcPending
							if v763 != 0 {
								return int32(0)
							} else {
								v764 = F_systable_getnext(m, v762)
								mBase = m.M
								v765 = m.ExcPending
								if v765 != 0 {
									return int32(0)
								} else {
									if v764 == int32(0) {
										if l1 != 0 {
											F_systable_endscan(m, v762)
											mBase = m.M
											v815 = m.ExcPending
											if v815 != 0 {
												return int32(0)
											} else {
												F_sequence_close(m, v746, int32(1))
												mBase = m.M
												v818 = m.ExcPending
												if v818 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v771 = m.ExcPending
											if v771 != 0 {
												return int32(0)
											} else {
												v772 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+464)) = v772
												F_errmsg_internal(m, int32(49423), v11+int32(464))
												mBase = m.M
												v778 = m.ExcPending
												if v778 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(472250), int32(3384), int32(235362))
													mBase = m.M
													v783 = m.ExcPending
													if v783 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v784 = *(*int32)(unsafe.Add(mBase, uint32(v764)+16))
										v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+22)))
										F_initStringInfo(m, v11+int32(1344))
										mBase = m.M
										v789 = m.ExcPending
										if v789 != 0 {
											return int32(0)
										} else {
											v792 = v784 + v785
											v793 = *(*int32)(unsafe.Add(mBase, uint32(v792)+68))
											F_getRelationDescription(m, v11+int32(1344), v793, int32(0))
											mBase = m.M
											v796 = m.ExcPending
											if v796 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+480)) = v792 + int32(4)
												v800 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+484)) = v800
												F_appendStringInfo(m, v11+int32(1408), int32(174867), v11+int32(480))
												mBase = m.M
												v808 = m.ExcPending
												if v808 != 0 {
													return int32(0)
												} else {
													v809 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
													F_pfree(m, v809)
													mBase = m.M
													v811 = m.ExcPending
													if v811 != 0 {
														return int32(0)
													} else {
														F_systable_endscan(m, v762)
														mBase = m.M
														v815 = m.ExcPending
														if v815 != 0 {
															return int32(0)
														} else {
															F_sequence_close(m, v746, int32(1))
															mBase = m.M
															v818 = m.ExcPending
															if v818 != 0 {
																return int32(0)
															} else {
																v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																m.G0 = v11 + int32(1424)
																if v1872 != 0 {
																	v1878 = v1873
																} else {
																	v1878 = int32(0)
																}
																return v1878
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				case 19:
					v821 = F_table_open(m, int32(2620), int32(1))
					mBase = m.M
					v822 = m.ExcPending
					if v822 != 0 {
						return int32(0)
					} else {
						v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_ScanKeyInit(m, v11+int32(1360), int32(1), int32(3), int32(184), v828)
						mBase = m.M
						v830 = m.ExcPending
						if v830 != 0 {
							return int32(0)
						} else {
							v832 = int32(1)
							v837 = F_systable_beginscan(m, v821, int32(2702), v832, int32(0), v832, v11+int32(1360))
							mBase = m.M
							v838 = m.ExcPending
							if v838 != 0 {
								return int32(0)
							} else {
								v839 = F_systable_getnext(m, v837)
								mBase = m.M
								v840 = m.ExcPending
								if v840 != 0 {
									return int32(0)
								} else {
									if v839 == int32(0) {
										if l1 != 0 {
											F_systable_endscan(m, v837)
											mBase = m.M
											v890 = m.ExcPending
											if v890 != 0 {
												return int32(0)
											} else {
												F_sequence_close(m, v821, int32(1))
												mBase = m.M
												v893 = m.ExcPending
												if v893 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v846 = m.ExcPending
											if v846 != 0 {
												return int32(0)
											} else {
												v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+496)) = v847
												F_errmsg_internal(m, int32(41598), v11+int32(496))
												mBase = m.M
												v853 = m.ExcPending
												if v853 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(472250), int32(3430), int32(235362))
													mBase = m.M
													v858 = m.ExcPending
													if v858 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v859 = *(*int32)(unsafe.Add(mBase, uint32(v839)+16))
										v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859)+22)))
										F_initStringInfo(m, v11+int32(1344))
										mBase = m.M
										v864 = m.ExcPending
										if v864 != 0 {
											return int32(0)
										} else {
											v867 = v859 + v860
											v868 = *(*int32)(unsafe.Add(mBase, uint32(v867)+4))
											F_getRelationDescription(m, v11+int32(1344), v868, int32(0))
											mBase = m.M
											v871 = m.ExcPending
											if v871 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+512)) = v867 + int32(12)
												v875 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+516)) = v875
												F_appendStringInfo(m, v11+int32(1408), int32(174813), v11+int32(512))
												mBase = m.M
												v883 = m.ExcPending
												if v883 != 0 {
													return int32(0)
												} else {
													v884 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
													F_pfree(m, v884)
													mBase = m.M
													v886 = m.ExcPending
													if v886 != 0 {
														return int32(0)
													} else {
														F_systable_endscan(m, v837)
														mBase = m.M
														v890 = m.ExcPending
														if v890 != 0 {
															return int32(0)
														} else {
															F_sequence_close(m, v821, int32(1))
															mBase = m.M
															v893 = m.ExcPending
															if v893 != 0 {
																return int32(0)
															} else {
																v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																m.G0 = v11 + int32(1424)
																if v1872 != 0 {
																	v1878 = v1873
																} else {
																	v1878 = int32(0)
																}
																return v1878
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				default:
					if v19 != int32(2328) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v1669 = m.ExcPending
						if v1669 != 0 {
							return int32(0)
						} else {
							v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1670
							F_errmsg_internal(m, int32(55485), v11)
							mBase = m.M
							v1674 = m.ExcPending
							if v1674 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(472250), int32(4072), int32(235362))
								mBase = m.M
								v1679 = m.ExcPending
								if v1679 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v1305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1306 = F_GetForeignDataWrapperExtended(m, v1305, l1)
						mBase = m.M
						v1307 = m.ExcPending
						if v1307 != 0 {
							return int32(0)
						} else {
							if v1306 == int32(0) {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							} else {
								v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1306)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+832)) = v1310
								F_appendStringInfo(m, v11+int32(1408), int32(172443), v11+int32(832))
								mBase = m.M
								v1318 = m.ExcPending
								if v1318 != 0 {
									return int32(0)
								} else {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								}
							}
						}
					}
				}
			}
		} else {
			if v19 <= int32(3575) {
				if v19 <= int32(3380) {
					if v19 == int32(2753) {
						v1682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_getOpFamilyDescription(m, v11+int32(1408), v1682, l1)
						mBase = m.M
						v1684 = m.ExcPending
						if v1684 != 0 {
							return int32(0)
						} else {
							v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
							v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
							m.G0 = v11 + int32(1424)
							if v1872 != 0 {
								v1878 = v1873
							} else {
								v1878 = int32(0)
							}
							return v1878
						}
					} else {
						if v19 == int32(3079) {
							v1379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v1380 = F_get_extension_name(m, v1379)
							mBase = m.M
							v1381 = m.ExcPending
							if v1381 != 0 {
								return int32(0)
							} else {
								if v1380 == int32(0) {
									if l1 != 0 {
										v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
										v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
										m.G0 = v11 + int32(1424)
										if v1872 != 0 {
											v1878 = v1873
										} else {
											v1878 = int32(0)
										}
										return v1878
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v1387 = m.ExcPending
										if v1387 != 0 {
											return int32(0)
										} else {
											v1388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+1104)) = v1388
											F_errmsg_internal(m, int32(44628), v11+int32(1104))
											mBase = m.M
											v1394 = m.ExcPending
											if v1394 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(472250), int32(3884), int32(235362))
												mBase = m.M
												v1399 = m.ExcPending
												if v1399 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+1120)) = v1380
									F_appendStringInfo(m, v11+int32(1408), int32(174703), v11+int32(1120))
									mBase = m.M
									v1407 = m.ExcPending
									if v1407 != 0 {
										return int32(0)
									} else {
										v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
										v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
										m.G0 = v11 + int32(1424)
										if v1872 != 0 {
											v1878 = v1873
										} else {
											v1878 = int32(0)
										}
										return v1878
									}
								}
							}
						} else {
							if v19 != int32(3256) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v1669 = m.ExcPending
								if v1669 != 0 {
									return int32(0)
								} else {
									v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1670
									F_errmsg_internal(m, int32(55485), v11)
									mBase = m.M
									v1674 = m.ExcPending
									if v1674 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(472250), int32(4072), int32(235362))
										mBase = m.M
										v1679 = m.ExcPending
										if v1679 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v42 = F_table_open(m, int32(3256), int32(1))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									F_ScanKeyInit(m, v11+int32(1360), int32(1), int32(3), int32(184), v49)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										v53 = int32(1)
										v58 = F_systable_beginscan(m, v42, int32(3257), v53, int32(0), v53, v11+int32(1360))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v60 = F_systable_getnext(m, v58)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return int32(0)
											} else {
												if v60 != 0 {
													v1483 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
													v1484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1483)+22)))
													F_initStringInfo(m, v11+int32(1344))
													mBase = m.M
													v1488 = m.ExcPending
													if v1488 != 0 {
														return int32(0)
													} else {
														v1491 = v1483 + v1484
														v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1491)+68))
														F_getRelationDescription(m, v11+int32(1344), v1492, int32(0))
														mBase = m.M
														v1495 = m.ExcPending
														if v1495 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+1216)) = v1491 + int32(4)
															v1499 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+1220)) = v1499
															F_appendStringInfo(m, v11+int32(1408), int32(174777), v11+int32(1216))
															mBase = m.M
															v1507 = m.ExcPending
															if v1507 != 0 {
																return int32(0)
															} else {
																v1508 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
																F_pfree(m, v1508)
																mBase = m.M
																v1510 = m.ExcPending
																if v1510 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v58)
																	mBase = m.M
																	v1514 = m.ExcPending
																	if v1514 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v42, int32(1))
																		mBase = m.M
																		v1517 = m.ExcPending
																		if v1517 != 0 {
																			return int32(0)
																		} else {
																			v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																			v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																			m.G0 = v11 + int32(1424)
																			if v1872 != 0 {
																				v1878 = v1873
																			} else {
																				v1878 = int32(0)
																			}
																			return v1878
																		}
																	}
																}
															}
														}
													}
												} else {
													if l1 != 0 {
														F_systable_endscan(m, v58)
														mBase = m.M
														v1514 = m.ExcPending
														if v1514 != 0 {
															return int32(0)
														} else {
															F_sequence_close(m, v42, int32(1))
															mBase = m.M
															v1517 = m.ExcPending
															if v1517 != 0 {
																return int32(0)
															} else {
																v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
																v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
																m.G0 = v11 + int32(1424)
																if v1872 != 0 {
																	v1878 = v1873
																} else {
																	v1878 = int32(0)
																}
																return v1878
															}
														}
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return int32(0)
														} else {
															v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+1200)) = v66
															F_errmsg_internal(m, int32(37677), v11+int32(1200))
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(472250), int32(3958), int32(235362))
																mBase = m.M
																v77 = m.ExcPending
																if v77 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					switch v19 - int32(3456) {
					case 0:
						v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v220 = F_SearchSysCache1(m, int32(16), v219)
						mBase = m.M
						v221 = m.ExcPending
						if v221 != 0 {
							return int32(0)
						} else {
							if v220 == int32(0) {
								if l1 != 0 {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v227 = m.ExcPending
									if v227 != 0 {
										return int32(0)
									} else {
										v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v228
										F_errmsg_internal(m, int32(43653), v11+int32(96))
										mBase = m.M
										v234 = m.ExcPending
										if v234 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(472250), int32(3023), int32(235362))
											mBase = m.M
											v239 = m.ExcPending
											if v239 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v240 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
								v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+22)))
								v242 = v240 + v241
								v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v244 = F_CollationIsVisible(m, v243)
								mBase = m.M
								v245 = m.ExcPending
								if v245 != 0 {
									return int32(0)
								} else {
									if v244 != 0 {
										v250 = int32(0)
										v253 = F_quote_qualified_identifier(m, v250, v242+int32(4))
										mBase = m.M
										v254 = m.ExcPending
										if v254 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v253
											F_appendStringInfo(m, v11+int32(1408), int32(174268), v11+int32(112))
											mBase = m.M
											v262 = m.ExcPending
											if v262 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v220)
												mBase = m.M
												v264 = m.ExcPending
												if v264 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										}
									} else {
										v247 = *(*int32)(unsafe.Add(mBase, uint32(v242)+68))
										v248 = F_get_namespace_name(m, v247)
										mBase = m.M
										v249 = m.ExcPending
										if v249 != 0 {
											return int32(0)
										} else {
											v250 = v248
											v253 = F_quote_qualified_identifier(m, v250, v242+int32(4))
											mBase = m.M
											v254 = m.ExcPending
											if v254 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v253
												F_appendStringInfo(m, v11+int32(1408), int32(174268), v11+int32(112))
												mBase = m.M
												v262 = m.ExcPending
												if v262 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v220)
													mBase = m.M
													v264 = m.ExcPending
													if v264 != 0 {
														return int32(0)
													} else {
														v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
														v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
														m.G0 = v11 + int32(1424)
														if v1872 != 0 {
															v1878 = v1873
														} else {
															v1878 = int32(0)
														}
														return v1878
													}
												}
											}
										}
									}
								}
							}
						}
					case 1, 2, 3, 4, 5, 6, 7, 8, 9:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v1669 = m.ExcPending
						if v1669 != 0 {
							return int32(0)
						} else {
							v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1670
							F_errmsg_internal(m, int32(55485), v11)
							mBase = m.M
							v1674 = m.ExcPending
							if v1674 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(472250), int32(4072), int32(235362))
								mBase = m.M
								v1679 = m.ExcPending
								if v1679 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 10:
						v1409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1410 = F_SearchSysCache1(m, int32(26), v1409)
						mBase = m.M
						v1411 = m.ExcPending
						if v1411 != 0 {
							return int32(0)
						} else {
							if v1410 == int32(0) {
								if l1 != 0 {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1417 = m.ExcPending
									if v1417 != 0 {
										return int32(0)
									} else {
										v1418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+1136)) = v1418
										F_errmsg_internal(m, int32(41527), v11+int32(1136))
										mBase = m.M
										v1424 = m.ExcPending
										if v1424 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(472250), int32(3901), int32(235362))
											mBase = m.M
											v1429 = m.ExcPending
											if v1429 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1410)+16))
								v1431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1430)+22)))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+1152)) = v1430 + v1431 + int32(4)
								F_appendStringInfo(m, v11+int32(1408), int32(172660), v11+int32(1152))
								mBase = m.M
								v1442 = m.ExcPending
								if v1442 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v1410)
									mBase = m.M
									v1444 = m.ExcPending
									if v1444 != 0 {
										return int32(0)
									} else {
										v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
										v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
										m.G0 = v11 + int32(1424)
										if v1872 != 0 {
											v1878 = v1873
										} else {
											v1878 = int32(0)
										}
										return v1878
									}
								}
							}
						}
					default:
						if v19 != int32(3381) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v1669 = m.ExcPending
							if v1669 != 0 {
								return int32(0)
							} else {
								v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1670
								F_errmsg_internal(m, int32(55485), v11)
								mBase = m.M
								v1674 = m.ExcPending
								if v1674 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(472250), int32(4072), int32(235362))
									mBase = m.M
									v1679 = m.ExcPending
									if v1679 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v927 = F_SearchSysCache1(m, int32(64), v926)
							mBase = m.M
							v928 = m.ExcPending
							if v928 != 0 {
								return int32(0)
							} else {
								if v927 == int32(0) {
									if l1 != 0 {
										v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
										v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
										m.G0 = v11 + int32(1424)
										if v1872 != 0 {
											v1878 = v1873
										} else {
											v1878 = int32(0)
										}
										return v1878
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v934 = m.ExcPending
										if v934 != 0 {
											return int32(0)
										} else {
											v935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+560)) = v935
											F_errmsg_internal(m, int32(39512), v11+int32(560))
											mBase = m.M
											v941 = m.ExcPending
											if v941 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(472250), int32(3479), int32(235362))
												mBase = m.M
												v946 = m.ExcPending
												if v946 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v947 = *(*int32)(unsafe.Add(mBase, uint32(v927)+16))
									v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947)+22)))
									v949 = v947 + v948
									v950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v952 = F_StatisticsObjIsVisibleExt(m, v950, int32(0))
									mBase = m.M
									v953 = m.ExcPending
									if v953 != 0 {
										return int32(0)
									} else {
										if v952 != 0 {
											v958 = int32(0)
											v961 = F_quote_qualified_identifier(m, v958, v949+int32(8))
											mBase = m.M
											v962 = m.ExcPending
											if v962 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+576)) = v961
												F_appendStringInfo(m, v11+int32(1408), int32(169701), v11+int32(576))
												mBase = m.M
												v970 = m.ExcPending
												if v970 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v927)
													mBase = m.M
													v972 = m.ExcPending
													if v972 != 0 {
														return int32(0)
													} else {
														v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
														v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
														m.G0 = v11 + int32(1424)
														if v1872 != 0 {
															v1878 = v1873
														} else {
															v1878 = int32(0)
														}
														return v1878
													}
												}
											}
										} else {
											v955 = *(*int32)(unsafe.Add(mBase, uint32(v949)+72))
											v956 = F_get_namespace_name(m, v955)
											mBase = m.M
											v957 = m.ExcPending
											if v957 != 0 {
												return int32(0)
											} else {
												v958 = v956
												v961 = F_quote_qualified_identifier(m, v958, v949+int32(8))
												mBase = m.M
												v962 = m.ExcPending
												if v962 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+576)) = v961
													F_appendStringInfo(m, v11+int32(1408), int32(169701), v11+int32(576))
													mBase = m.M
													v970 = m.ExcPending
													if v970 != 0 {
														return int32(0)
													} else {
														F_ReleaseCatCache(m, v927)
														mBase = m.M
														v972 = m.ExcPending
														if v972 != 0 {
															return int32(0)
														} else {
															v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
															v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
															m.G0 = v11 + int32(1424)
															if v1872 != 0 {
																v1878 = v1873
															} else {
																v1878 = int32(0)
															}
															return v1878
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				if v19 <= int32(6099) {
					switch v19 - int32(3576) {
					case 0:
						v1622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1623 = F_SearchSysCache1(m, int32(70), v1622)
						mBase = m.M
						v1624 = m.ExcPending
						if v1624 != 0 {
							return int32(0)
						} else {
							if v1623 == int32(0) {
								if l1 != 0 {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1630 = m.ExcPending
									if v1630 != 0 {
										return int32(0)
									} else {
										v1631 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+1312)) = v1631
										F_errmsg_internal(m, int32(44876), v11+int32(1312))
										mBase = m.M
										v1637 = m.ExcPending
										if v1637 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(472250), int32(4057), int32(235362))
											mBase = m.M
											v1642 = m.ExcPending
											if v1642 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+16))
								v1644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1643)+22)))
								v1645 = v1643 + v1644
								v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1645)+4))
								v1647 = F_format_type_be(m, v1646)
								mBase = m.M
								v1648 = m.ExcPending
								if v1648 != 0 {
									return int32(0)
								} else {
									v1649 = *(*int32)(unsafe.Add(mBase, uint32(v1645)+8))
									v1651 = F_get_language_name(m, v1649, int32(0))
									mBase = m.M
									v1652 = m.ExcPending
									if v1652 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+1332)) = v1651
										*(*int32)(unsafe.Add(mBase, uint32(v11)+1328)) = v1647
										F_appendStringInfo(m, v11+int32(1408), int32(185774), v11+int32(1328))
										mBase = m.M
										v1661 = m.ExcPending
										if v1661 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v1623)
											mBase = m.M
											v1663 = m.ExcPending
											if v1663 != 0 {
												return int32(0)
											} else {
												v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
												v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
												m.G0 = v11 + int32(1424)
												if v1872 != 0 {
													v1878 = v1873
												} else {
													v1878 = int32(0)
												}
												return v1878
											}
										}
									}
								}
							}
						}
					case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v1669 = m.ExcPending
						if v1669 != 0 {
							return int32(0)
						} else {
							v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1670
							F_errmsg_internal(m, int32(55485), v11)
							mBase = m.M
							v1674 = m.ExcPending
							if v1674 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(472250), int32(4072), int32(235362))
								mBase = m.M
								v1679 = m.ExcPending
								if v1679 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 24:
						v1022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1023 = F_SearchSysCache1(m, int32(76), v1022)
						mBase = m.M
						v1024 = m.ExcPending
						if v1024 != 0 {
							return int32(0)
						} else {
							if v1023 == int32(0) {
								if l1 != 0 {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1030 = m.ExcPending
									if v1030 != 0 {
										return int32(0)
									} else {
										v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+624)) = v1031
										F_errmsg_internal(m, int32(37238), v11+int32(624))
										mBase = m.M
										v1037 = m.ExcPending
										if v1037 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(472250), int32(3541), int32(235362))
											mBase = m.M
											v1042 = m.ExcPending
											if v1042 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+16))
								v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043)+22)))
								v1045 = v1043 + v1044
								v1046 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v1047 = F_TSDictionaryIsVisible(m, v1046)
								mBase = m.M
								v1048 = m.ExcPending
								if v1048 != 0 {
									return int32(0)
								} else {
									if v1047 != 0 {
										v1053 = int32(0)
										v1056 = F_quote_qualified_identifier(m, v1053, v1045+int32(4))
										mBase = m.M
										v1057 = m.ExcPending
										if v1057 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+640)) = v1056
											F_appendStringInfo(m, v11+int32(1408), int32(168136), v11+int32(640))
											mBase = m.M
											v1065 = m.ExcPending
											if v1065 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v1023)
												mBase = m.M
												v1067 = m.ExcPending
												if v1067 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										}
									} else {
										v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+68))
										v1051 = F_get_namespace_name(m, v1050)
										mBase = m.M
										v1052 = m.ExcPending
										if v1052 != 0 {
											return int32(0)
										} else {
											v1053 = v1051
											v1056 = F_quote_qualified_identifier(m, v1053, v1045+int32(4))
											mBase = m.M
											v1057 = m.ExcPending
											if v1057 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+640)) = v1056
												F_appendStringInfo(m, v11+int32(1408), int32(168136), v11+int32(640))
												mBase = m.M
												v1065 = m.ExcPending
												if v1065 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v1023)
													mBase = m.M
													v1067 = m.ExcPending
													if v1067 != 0 {
														return int32(0)
													} else {
														v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
														v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
														m.G0 = v11 + int32(1424)
														if v1872 != 0 {
															v1878 = v1873
														} else {
															v1878 = int32(0)
														}
														return v1878
													}
												}
											}
										}
									}
								}
							}
						}
					case 25:
						v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v975 = F_SearchSysCache1(m, int32(78), v974)
						mBase = m.M
						v976 = m.ExcPending
						if v976 != 0 {
							return int32(0)
						} else {
							if v975 == int32(0) {
								if l1 != 0 {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v982 = m.ExcPending
									if v982 != 0 {
										return int32(0)
									} else {
										v983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+592)) = v983
										F_errmsg_internal(m, int32(41400), v11+int32(592))
										mBase = m.M
										v989 = m.ExcPending
										if v989 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(472250), int32(3511), int32(235362))
											mBase = m.M
											v994 = m.ExcPending
											if v994 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v995 = *(*int32)(unsafe.Add(mBase, uint32(v975)+16))
								v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+22)))
								v997 = v995 + v996
								v998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v1000 = F_TSParserIsVisibleExt(m, v998, int32(0))
								mBase = m.M
								v1001 = m.ExcPending
								if v1001 != 0 {
									return int32(0)
								} else {
									if v1000 != 0 {
										v1006 = int32(0)
										v1009 = F_quote_qualified_identifier(m, v1006, v997+int32(4))
										mBase = m.M
										v1010 = m.ExcPending
										if v1010 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+608)) = v1009
											F_appendStringInfo(m, v11+int32(1408), int32(172358), v11+int32(608))
											mBase = m.M
											v1018 = m.ExcPending
											if v1018 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v975)
												mBase = m.M
												v1020 = m.ExcPending
												if v1020 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										}
									} else {
										v1003 = *(*int32)(unsafe.Add(mBase, uint32(v997)+68))
										v1004 = F_get_namespace_name(m, v1003)
										mBase = m.M
										v1005 = m.ExcPending
										if v1005 != 0 {
											return int32(0)
										} else {
											v1006 = v1004
											v1009 = F_quote_qualified_identifier(m, v1006, v997+int32(4))
											mBase = m.M
											v1010 = m.ExcPending
											if v1010 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+608)) = v1009
												F_appendStringInfo(m, v11+int32(1408), int32(172358), v11+int32(608))
												mBase = m.M
												v1018 = m.ExcPending
												if v1018 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v975)
													mBase = m.M
													v1020 = m.ExcPending
													if v1020 != 0 {
														return int32(0)
													} else {
														v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
														v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
														m.G0 = v11 + int32(1424)
														if v1872 != 0 {
															v1878 = v1873
														} else {
															v1878 = int32(0)
														}
														return v1878
													}
												}
											}
										}
									}
								}
							}
						}
					case 26:
						v1119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1120 = F_SearchSysCache1(m, int32(74), v1119)
						mBase = m.M
						v1121 = m.ExcPending
						if v1121 != 0 {
							return int32(0)
						} else {
							if v1120 == int32(0) {
								if l1 != 0 {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1127 = m.ExcPending
									if v1127 != 0 {
										return int32(0)
									} else {
										v1128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+688)) = v1128
										F_errmsg_internal(m, int32(43484), v11+int32(688))
										mBase = m.M
										v1134 = m.ExcPending
										if v1134 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(472250), int32(3603), int32(235362))
											mBase = m.M
											v1139 = m.ExcPending
											if v1139 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+16))
								v1141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1140)+22)))
								v1142 = v1140 + v1141
								v1143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v1144 = F_TSConfigIsVisible(m, v1143)
								mBase = m.M
								v1145 = m.ExcPending
								if v1145 != 0 {
									return int32(0)
								} else {
									if v1144 != 0 {
										v1150 = int32(0)
										v1153 = F_quote_qualified_identifier(m, v1150, v1142+int32(4))
										mBase = m.M
										v1154 = m.ExcPending
										if v1154 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+704)) = v1153
											F_appendStringInfo(m, v11+int32(1408), int32(174187), v11+int32(704))
											mBase = m.M
											v1162 = m.ExcPending
											if v1162 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v1120)
												mBase = m.M
												v1164 = m.ExcPending
												if v1164 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										}
									} else {
										v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+68))
										v1148 = F_get_namespace_name(m, v1147)
										mBase = m.M
										v1149 = m.ExcPending
										if v1149 != 0 {
											return int32(0)
										} else {
											v1150 = v1148
											v1153 = F_quote_qualified_identifier(m, v1150, v1142+int32(4))
											mBase = m.M
											v1154 = m.ExcPending
											if v1154 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+704)) = v1153
												F_appendStringInfo(m, v11+int32(1408), int32(174187), v11+int32(704))
												mBase = m.M
												v1162 = m.ExcPending
												if v1162 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v1120)
													mBase = m.M
													v1164 = m.ExcPending
													if v1164 != 0 {
														return int32(0)
													} else {
														v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
														v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
														m.G0 = v11 + int32(1424)
														if v1872 != 0 {
															v1878 = v1873
														} else {
															v1878 = int32(0)
														}
														return v1878
													}
												}
											}
										}
									}
								}
							}
						}
					default:
						if v19 != int32(3764) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v1669 = m.ExcPending
							if v1669 != 0 {
								return int32(0)
							} else {
								v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1670
								F_errmsg_internal(m, int32(55485), v11)
								mBase = m.M
								v1674 = m.ExcPending
								if v1674 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(472250), int32(4072), int32(235362))
									mBase = m.M
									v1679 = m.ExcPending
									if v1679 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v1071 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v1072 = F_SearchSysCache1(m, int32(80), v1071)
							mBase = m.M
							v1073 = m.ExcPending
							if v1073 != 0 {
								return int32(0)
							} else {
								if v1072 == int32(0) {
									if l1 != 0 {
										v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
										v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
										m.G0 = v11 + int32(1424)
										if v1872 != 0 {
											v1878 = v1873
										} else {
											v1878 = int32(0)
										}
										return v1878
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v1079 = m.ExcPending
										if v1079 != 0 {
											return int32(0)
										} else {
											v1080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+656)) = v1080
											F_errmsg_internal(m, int32(46886), v11+int32(656))
											mBase = m.M
											v1086 = m.ExcPending
											if v1086 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(472250), int32(3572), int32(235362))
												mBase = m.M
												v1091 = m.ExcPending
												if v1091 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+16))
									v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1092)+22)))
									v1094 = v1092 + v1093
									v1095 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v1097 = F_TSTemplateIsVisibleExt(m, v1095, int32(0))
									mBase = m.M
									v1098 = m.ExcPending
									if v1098 != 0 {
										return int32(0)
									} else {
										if v1097 != 0 {
											v1103 = int32(0)
											v1106 = F_quote_qualified_identifier(m, v1103, v1094+int32(4))
											mBase = m.M
											v1107 = m.ExcPending
											if v1107 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+672)) = v1106
												F_appendStringInfo(m, v11+int32(1408), int32(177876), v11+int32(672))
												mBase = m.M
												v1115 = m.ExcPending
												if v1115 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v1072)
													mBase = m.M
													v1117 = m.ExcPending
													if v1117 != 0 {
														return int32(0)
													} else {
														v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
														v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
														m.G0 = v11 + int32(1424)
														if v1872 != 0 {
															v1878 = v1873
														} else {
															v1878 = int32(0)
														}
														return v1878
													}
												}
											}
										} else {
											v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+68))
											v1101 = F_get_namespace_name(m, v1100)
											mBase = m.M
											v1102 = m.ExcPending
											if v1102 != 0 {
												return int32(0)
											} else {
												v1103 = v1101
												v1106 = F_quote_qualified_identifier(m, v1103, v1094+int32(4))
												mBase = m.M
												v1107 = m.ExcPending
												if v1107 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+672)) = v1106
													F_appendStringInfo(m, v11+int32(1408), int32(177876), v11+int32(672))
													mBase = m.M
													v1115 = m.ExcPending
													if v1115 != 0 {
														return int32(0)
													} else {
														F_ReleaseCatCache(m, v1072)
														mBase = m.M
														v1117 = m.ExcPending
														if v1117 != 0 {
															return int32(0)
														} else {
															v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
															v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
															m.G0 = v11 + int32(1424)
															if v1872 != 0 {
																v1878 = v1873
															} else {
																v1878 = int32(0)
															}
															return v1878
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					switch v19 - int32(6100) {
					case 0:
						v1608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1609 = F_get_subscription_name(m, v1608, l1)
						mBase = m.M
						v1610 = m.ExcPending
						if v1610 != 0 {
							return int32(0)
						} else {
							if v1609 == int32(0) {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+1296)) = v1609
								F_appendStringInfo(m, v11+int32(1408), int32(173873), v11+int32(1296))
								mBase = m.M
								v1620 = m.ExcPending
								if v1620 != 0 {
									return int32(0)
								} else {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								}
							}
						}
					case 1, 2, 3, 5:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v1669 = m.ExcPending
						if v1669 != 0 {
							return int32(0)
						} else {
							v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1670
							F_errmsg_internal(m, int32(55485), v11)
							mBase = m.M
							v1674 = m.ExcPending
							if v1674 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(472250), int32(4072), int32(235362))
								mBase = m.M
								v1679 = m.ExcPending
								if v1679 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 4:
						v1518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1519 = F_get_publication_name(m, v1518, l1)
						mBase = m.M
						v1520 = m.ExcPending
						if v1520 != 0 {
							return int32(0)
						} else {
							if v1519 == int32(0) {
								v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
								v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
								m.G0 = v11 + int32(1424)
								if v1872 != 0 {
									v1878 = v1873
								} else {
									v1878 = int32(0)
								}
								return v1878
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+1232)) = v1519
								F_appendStringInfo(m, v11+int32(1408), int32(174482), v11+int32(1232))
								mBase = m.M
								v1530 = m.ExcPending
								if v1530 != 0 {
									return int32(0)
								} else {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								}
							}
						}
					case 6:
						v1555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1556 = F_SearchSysCache1(m, int32(52), v1555)
						mBase = m.M
						v1557 = m.ExcPending
						if v1557 != 0 {
							return int32(0)
						} else {
							if v1556 == int32(0) {
								if l1 != 0 {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1563 = m.ExcPending
									if v1563 != 0 {
										return int32(0)
									} else {
										v1564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+1264)) = v1564
										F_errmsg_internal(m, int32(49521), v11+int32(1264))
										mBase = m.M
										v1570 = m.ExcPending
										if v1570 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(472250), int32(4018), int32(235362))
											mBase = m.M
											v1575 = m.ExcPending
											if v1575 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1556)+16))
								v1577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1576)+22)))
								v1578 = v1576 + v1577
								v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1578)+4))
								v1581 = F_get_publication_name(m, v1579, int32(0))
								mBase = m.M
								v1582 = m.ExcPending
								if v1582 != 0 {
									return int32(0)
								} else {
									F_initStringInfo(m, v11+int32(1360))
									mBase = m.M
									v1586 = m.ExcPending
									if v1586 != 0 {
										return int32(0)
									} else {
										v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1578)+8))
										F_getRelationDescription(m, v11+int32(1360), v1589, int32(0))
										mBase = m.M
										v1592 = m.ExcPending
										if v1592 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+1284)) = v1581
											v1594 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+1280)) = v1594
											F_appendStringInfo(m, v11+int32(1408), int32(174386), v11+int32(1280))
											mBase = m.M
											v1602 = m.ExcPending
											if v1602 != 0 {
												return int32(0)
											} else {
												v1603 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
												F_pfree(m, v1603)
												mBase = m.M
												v1605 = m.ExcPending
												if v1605 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v1556)
													mBase = m.M
													v1607 = m.ExcPending
													if v1607 != 0 {
														return int32(0)
													} else {
														v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
														v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
														m.G0 = v11 + int32(1424)
														if v1872 != 0 {
															v1878 = v1873
														} else {
															v1878 = int32(0)
														}
														return v1878
													}
												}
											}
										}
									}
								}
							}
						}
					default:
						switch v19 - int32(6237) {
						case 0:
							v1535 = F_getPublicationSchemaInfo(m, l0, l1, v11+int32(1360), v11+int32(1344))
							mBase = m.M
							v1536 = m.ExcPending
							if v1536 != 0 {
								return int32(0)
							} else {
								if v1535 == int32(0) {
									v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
									v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
									m.G0 = v11 + int32(1424)
									if v1872 != 0 {
										v1878 = v1873
									} else {
										v1878 = int32(0)
									}
									return v1878
								} else {
									v1539 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+1248)) = v1539
									v1541 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+1252)) = v1541
									F_appendStringInfo(m, v11+int32(1408), int32(174422), v11+int32(1248))
									mBase = m.M
									v1549 = m.ExcPending
									if v1549 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v1541)
										mBase = m.M
										v1551 = m.ExcPending
										if v1551 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v1539)
											mBase = m.M
											v1553 = m.ExcPending
											if v1553 != 0 {
												return int32(0)
											} else {
												v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
												v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
												m.G0 = v11 + int32(1424)
												if v1872 != 0 {
													v1878 = v1873
												} else {
													v1878 = int32(0)
												}
												return v1878
											}
										}
									}
								}
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v1669 = m.ExcPending
							if v1669 != 0 {
								return int32(0)
							} else {
								v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1670
								F_errmsg_internal(m, int32(55485), v11)
								mBase = m.M
								v1674 = m.ExcPending
								if v1674 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(472250), int32(4072), int32(235362))
									mBase = m.M
									v1679 = m.ExcPending
									if v1679 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 6:
							v1446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v1447 = F_SearchSysCache1(m, int32(44), v1446)
							mBase = m.M
							v1448 = m.ExcPending
							if v1448 != 0 {
								return int32(0)
							} else {
								if v1447 == int32(0) {
									if l1 != 0 {
										v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
										v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
										m.G0 = v11 + int32(1424)
										if v1872 != 0 {
											v1878 = v1873
										} else {
											v1878 = int32(0)
										}
										return v1878
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v1454 = m.ExcPending
										if v1454 != 0 {
											return int32(0)
										} else {
											v1455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+1168)) = v1455
											F_errmsg_internal(m, int32(53339), v11+int32(1168))
											mBase = m.M
											v1461 = m.ExcPending
											if v1461 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(472250), int32(3922), int32(235362))
												mBase = m.M
												v1466 = m.ExcPending
												if v1466 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v1469 = F_SysCacheGetAttrNotNull(m, int32(44), v1447, int32(2))
									mBase = m.M
									v1470 = m.ExcPending
									if v1470 != 0 {
										return int32(0)
									} else {
										v1471 = F_text_to_cstring(m, v1469)
										mBase = m.M
										v1472 = m.ExcPending
										if v1472 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+1184)) = v1471
											F_appendStringInfo(m, v11+int32(1408), int32(172345), v11+int32(1184))
											mBase = m.M
											v1480 = m.ExcPending
											if v1480 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v1447)
												mBase = m.M
												v1482 = m.ExcPending
												if v1482 != 0 {
													return int32(0)
												} else {
													v1872 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1412))
													v1873 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1408))
													m.G0 = v11 + int32(1424)
													if v1872 != 0 {
														v1878 = v1873
													} else {
														v1878 = int32(0)
													}
													return v1878
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
