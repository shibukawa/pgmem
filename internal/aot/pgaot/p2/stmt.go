package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecDropStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int64
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
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
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v866 int32
	_ = v866
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
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1049 int32
	_ = v1049
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1100 int32
	_ = v1100
	var v1107 int32
	_ = v1107
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1137 int32
	_ = v1137
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int64
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int64
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1236 int32
	_ = v1236
	var v1240 int32
	_ = v1240
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1257 int32
	_ = v1257
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1272 int32
	_ = v1272
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1435 int32
	_ = v1435
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
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
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1569 int32
	_ = v1569
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1639 int32
	_ = v1639
	var v1643 int32
	_ = v1643
	var v1648 int32
	_ = v1648
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch v13 - int32(18) {
	case 0, 5, 19, 23, 33:
		goto L2
	default:
		goto L1
	case 2:
		goto L3
	}
L1:
	;
	v321 = m.G0
	v323 = v321 - int32(176)
	m.G0 = v323
	v325 = F_new_object_addresses(m)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L5
	} else {
		goto L107
	}
L2:
	;
	v22 = m.G0
	v24 = v22 - int32(128)
	m.G0 = v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v26 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v16 != int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	F_PreventInTransactionBlock(m, l1, int32(_a_F_ExecDropStmt_0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	goto L2
L7:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_performMultipleDeletions(m, v66, v313, v308)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L105
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L5
	} else {
		goto L101
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L5
	} else {
		goto L97
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L5
	} else {
		goto L93
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L89
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L5
	} else {
		goto L85
	}
L13:
	;
	v41 = int32(8)
	goto L15
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v30 == int32(0) {
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch v43 - int32(18) {
	case 0:
		goto L22
	default:
		goto L21
	case 2:
		goto L20
	case 5:
		goto L23
	case 19:
		goto L25
	case 23:
		v65 = int32(114)
		goto L19
	case 33:
		goto L24
	}
L16:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v33 != int32(1) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v36 == int32(1) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v41 = int32(4)
	goto L15
L19:
	;
	v66 = F_new_object_addresses(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L29
	}
L20:
	;
	v65 = int32(105)
	goto L19
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L26
	}
L22:
	;
	v65 = int32(102)
	goto L19
L23:
	;
	v65 = int32(109)
	goto L19
L24:
	;
	v65 = int32(118)
	goto L19
L25:
	;
	v65 = int32(83)
	goto L19
L26:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v54
	F_errmsg_internal(m, int32(_a_F_ExecDropStmt_1), v24)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1601), int32(_a_F_ExecDropStmt_3))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v68 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v80 = v69
	v83 = v3
	goto L35
L31:
	;
	v69 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v69 < v70 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v308 = v3
	goto L7
L34:
	;
	goto L33
L35:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v80<<(uint(int32(2))%32))))
	v93 = F_makeRangeVarFromNameList(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L37
	}
L36:
	;
	v308 = v207
	goto L7
L37:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+96)) = uint8(v65)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+104)) = int64(0)
	if v98 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v103 = int32(4)
	goto L41
L40:
	;
	v103 = int32(8)
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+100)) = v103
	v109 = F_RangeVarGetRelidExtended(m, v93, v41, int32(1), int32(572), v24+int32(96))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L43
	}
L42:
	;
	v209 = v80 + int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v209 < v210 {
		v80 = v209
		v83 = v207
		goto L35
	} else {
		goto L84
	}
L43:
	;
	if v109 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	if v114 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+113)))
	if v174 != int32(116) {
		goto L69
	} else {
		goto L70
	}
L47:
	;
	switch v65 - int32(83) {
	case 0:
		goto L63
	default:
		v207 = v83
		goto L42
	case 16:
		goto L59
	case 19:
		goto L58
	case 22:
		goto L60
	case 26:
		goto L61
	case 29:
		goto L57
	case 31:
		v149 = int32(_a_F_ExecDropStmt_4)
		goto L56
	case 35:
		goto L62
	}
L48:
	;
	v117 = F_LookupNamespaceNoError(m, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	if v117 != 0 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	if v113&int32(1) == int32(0) {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	v125 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	if v125 == int32(0) {
		v207 = v83
		goto L42
	} else {
		goto L53
	}
L53:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v129
	F_errmsg(m, int32(_a_F_ExecDropStmt_5), v24+int32(48))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1480), int32(_a_F_ExecDropStmt_6))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v207 = v83
	goto L42
L56:
	;
	if v113&int32(1) == int32(0) {
		goto L9
	} else {
		goto L64
	}
L57:
	;
	v149 = int32(_a_F_ExecDropStmt_7)
	goto L56
L58:
	;
	v149 = int32(_a_F_ExecDropStmt_8)
	goto L56
L59:
	;
	v149 = int32(_a_F_ExecDropStmt_9)
	goto L56
L60:
	;
	v149 = int32(_a_F_ExecDropStmt_10)
	goto L56
L61:
	;
	v149 = int32(_a_F_ExecDropStmt_11)
	goto L56
L62:
	;
	v149 = int32(_a_F_ExecDropStmt_12)
	goto L56
L63:
	;
	v149 = int32(_a_F_ExecDropStmt_13)
	goto L56
L64:
	;
	v156 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	if v156 == int32(0) {
		v207 = v83
		goto L42
	} else {
		goto L66
	}
L66:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v161
	F_errmsg(m, v160, v24+int32(16))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1497), int32(_a_F_ExecDropStmt_6))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v207 = v83
	goto L42
L69:
	;
	v177 = v83 | int32(2)
	goto L71
L70:
	;
	v177 = v83
	goto L71
L71:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v178 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v179 = v177
	goto L74
L73:
	;
	v179 = v83
	goto L74
L74:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+112)))
	if v183 == int32(73) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v186 = v179 & int32(2)
	goto L77
L76:
	;
	v186 = int32(0)
	goto L77
L77:
	;
	if v186 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	if v183&int32(255) == int32(73) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	v194 = F_find_all_inheritors(m, v191, v192, int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L5
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+124)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+120)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v24)+116)) = int32(1259)
	F_add_exact_object_address(m, v24+int32(116), v66)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L5
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v207 = v179
	goto L42
L84:
	;
	goto L36
L85:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(_a_F_ExecDropStmt_14), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1559), int32(_a_F_ExecDropStmt_3))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(_a_F_ExecDropStmt_15), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1563), int32(_a_F_ExecDropStmt_3))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v251
	F_errmsg(m, int32(_a_F_ExecDropStmt_16), v24-int32(-64))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1474), int32(_a_F_ExecDropStmt_6))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	F_errcode(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v271
	F_errmsg(m, v270, v24+int32(32))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1493), int32(_a_F_ExecDropStmt_6))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v289
	F_errmsg(m, int32(_a_F_ExecDropStmt_17), v24+int32(80))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1668), int32(_a_F_ExecDropStmt_3))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_free_object_addresses(m, v66)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	m.G0 = v24 + int32(128)
	return
L107:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v327 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	return
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		goto L5
	} else {
		goto L505
	}
L110:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_performMultipleDeletions(m, v325, v1615, int32(0))
	mBase = m.M
	v1618 = m.ExcPending
	if v1618 != 0 {
		goto L5
	} else {
		goto L503
	}
L111:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v330 <= int32(0) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v343 = v3
	goto L113
L113:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v327)+12))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v345+v343<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+148)) = int32(0)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	F_get_object_address(m, v323+int32(136), v354, v349, v323+int32(148), int32(8), v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L5
	} else {
		goto L115
	}
L114:
	;
	goto L110
L115:
	;
	v362 = v323 + int32(160)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v323)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v363
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v323)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v323)+152)) = v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v323)+156))
	if v369 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L116:
	;
	v1600 = v343 + int32(1)
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v1600 < v1601 {
		v343 = v1600
		goto L113
	} else {
		goto L502
	}
L117:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_18), v1583, int32(_a_F_ExecDropStmt_19))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L5
	} else {
		goto L501
	}
L118:
	;
	v1560 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L5
	} else {
		goto L498
	}
L119:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v323)+136))
	if v1527 == int32(0) {
		goto L109
	} else {
		goto L493
	}
L120:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v920)+4))
	v1497 = F_makeRangeVarFromNameList(m, v1496)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L5
	} else {
		goto L488
	}
L121:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1414)))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+60)) = v1415
	*(*int32)(unsafe.Add(mBase, uint32(v323)+168)) = v1415
	v1421 = F_list_make1_impl(m, int32(1), v323+int32(60))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L5
	} else {
		goto L467
	}
L122:
	;
	v1394 = F_owningrel_does_not_exist_skipping(m, v349, v323+int32(136), v323+int32(172))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L5
	} else {
		goto L461
	}
L123:
	;
	v1369 = F_owningrel_does_not_exist_skipping(m, v349, v323+int32(136), v323+int32(172))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L5
	} else {
		goto L457
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v1353
	v1547 = v1352
	goto L118
L125:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v1352 = int32(_a_F_ExecDropStmt_20)
	v1353 = v1349
	goto L124
L126:
	;
	v1329 = F_owningrel_does_not_exist_skipping(m, v349, v323+int32(136), v323+int32(172))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L5
	} else {
		goto L453
	}
L127:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v1352 = int32(_a_F_ExecDropStmt_21)
	v1353 = v1322
	goto L124
L128:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v1352 = int32(_a_F_ExecDropStmt_22)
	v1353 = v1320
	goto L124
L129:
	;
	v1301 = F_list_copy_tail(m, v349, int32(1))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L5
	} else {
		goto L447
	}
L130:
	;
	v1281 = F_list_copy_tail(m, v349, int32(1))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L5
	} else {
		goto L440
	}
L131:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v1352 = int32(_a_F_ExecDropStmt_23)
	v1353 = v1278
	goto L124
L132:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L5
	} else {
		goto L436
	}
L133:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L5
	} else {
		goto L433
	}
L134:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v1352 = int32(_a_F_ExecDropStmt_24)
	v1353 = v1246
	goto L124
L135:
	;
	v372 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v372
	switch v367 {
	case 0:
		goto L134
	case 1:
		goto L141
	case 2, 3, 4, 10, 11, 13, 22, 27, 31, 32, 40, 50:
		goto L133
	case 5:
		goto L138
	case 6, 9, 18, 20, 23, 33, 37, 38, 41, 42, 51:
		goto L132
	case 7:
		goto L153
	case 8:
		goto L152
	case 12, 49:
		goto L154
	case 14:
		goto L125
	case 15:
		goto L145
	case 16:
		goto L127
	case 17:
		goto L128
	case 19:
		goto L144
	case 21:
		goto L139
	case 24:
		goto L129
	case 25:
		goto L140
	case 26:
		goto L130
	case 28:
		goto L123
	case 29:
		goto L143
	case 30:
		goto L131
	case 34:
		goto L142
	case 35:
		goto L126
	case 36:
		goto L151
	case 39:
		goto L150
	case 43:
		goto L121
	case 44:
		goto L122
	case 45:
		goto L146
	case 46:
		goto L148
	case 47:
		goto L149
	case 48:
		goto L147
	default:
		goto L109
	}
L136:
	;
	goto L137
L137:
	;
	if v367 == int32(19) {
		goto L367
	} else {
		goto L368
	}
L138:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v889)))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+56)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v323)+168)) = v890
	v896 = F_list_make1_impl(m, int32(1), v323+int32(56))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L5
	} else {
		goto L333
	}
L139:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v1352 = int32(_a_F_ExecDropStmt_25)
	v1353 = v887
	goto L124
L140:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v810 = F_makeRangeVarFromNameList(m, v809)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L5
	} else {
		goto L308
	}
L141:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v727 = F_makeRangeVarFromNameList(m, v726)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L5
	} else {
		goto L282
	}
L142:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v644 = F_makeRangeVarFromNameList(m, v643)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L5
	} else {
		goto L256
	}
L143:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v561 = F_makeRangeVarFromNameList(m, v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L5
	} else {
		goto L230
	}
L144:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v478 = F_makeRangeVarFromNameList(m, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L5
	} else {
		goto L204
	}
L145:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v1352 = int32(_a_F_ExecDropStmt_26)
	v1353 = v475
	goto L124
L146:
	;
	v463 = F_makeRangeVarFromNameList(m, v349)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L5
	} else {
		goto L198
	}
L147:
	;
	v451 = F_makeRangeVarFromNameList(m, v349)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L5
	} else {
		goto L192
	}
L148:
	;
	v439 = F_makeRangeVarFromNameList(m, v349)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L5
	} else {
		goto L186
	}
L149:
	;
	v427 = F_makeRangeVarFromNameList(m, v349)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L5
	} else {
		goto L180
	}
L150:
	;
	v415 = F_makeRangeVarFromNameList(m, v349)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L5
	} else {
		goto L174
	}
L151:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v413
	goto L124
L152:
	;
	v401 = F_makeRangeVarFromNameList(m, v349)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L5
	} else {
		goto L168
	}
L153:
	;
	v389 = F_makeRangeVarFromNameList(m, v349)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L5
	} else {
		goto L162
	}
L154:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v377 = F_makeRangeVarFromNameList(m, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L5
	} else {
		goto L156
	}
L155:
	;
	v387 = F_TypeNameToString(m, v349)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L5
	} else {
		goto L160
	}
L156:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v377)+8))
	if v379 == int32(0) {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v382 = F_LookupNamespaceNoError(m, v379)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L5
	} else {
		goto L158
	}
L158:
	;
	if v382 != 0 {
		goto L155
	} else {
		goto L159
	}
L159:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v377)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v384
	goto L124
L160:
	;
	v1352 = int32(_a_F_ExecDropStmt_27)
	v1353 = v387
	goto L124
L161:
	;
	v399 = F_NameListToString(m, v349)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L5
	} else {
		goto L166
	}
L162:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v389)+8))
	if v391 == int32(0) {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v394 = F_LookupNamespaceNoError(m, v391)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L5
	} else {
		goto L164
	}
L164:
	;
	if v394 != 0 {
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v389)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v396
	goto L124
L166:
	;
	v1352 = int32(_a_F_ExecDropStmt_28)
	v1353 = v399
	goto L124
L167:
	;
	v411 = F_NameListToString(m, v349)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L5
	} else {
		goto L172
	}
L168:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v401)+8))
	if v403 == int32(0) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v406 = F_LookupNamespaceNoError(m, v403)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L5
	} else {
		goto L170
	}
L170:
	;
	if v406 != 0 {
		goto L167
	} else {
		goto L171
	}
L171:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v401)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v408
	goto L124
L172:
	;
	v1352 = int32(_a_F_ExecDropStmt_29)
	v1353 = v411
	goto L124
L173:
	;
	v425 = F_NameListToString(m, v349)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L5
	} else {
		goto L178
	}
L174:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v415)+8))
	if v417 == int32(0) {
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v420 = F_LookupNamespaceNoError(m, v417)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L5
	} else {
		goto L176
	}
L176:
	;
	if v420 != 0 {
		goto L173
	} else {
		goto L177
	}
L177:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v415)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v422
	goto L124
L178:
	;
	v1352 = int32(_a_F_ExecDropStmt_30)
	v1353 = v425
	goto L124
L179:
	;
	v437 = F_NameListToString(m, v349)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L5
	} else {
		goto L184
	}
L180:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v427)+8))
	if v429 == int32(0) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v432 = F_LookupNamespaceNoError(m, v429)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L5
	} else {
		goto L182
	}
L182:
	;
	if v432 != 0 {
		goto L179
	} else {
		goto L183
	}
L183:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v427)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v434
	goto L124
L184:
	;
	v1352 = int32(_a_F_ExecDropStmt_31)
	v1353 = v437
	goto L124
L185:
	;
	v449 = F_NameListToString(m, v349)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L5
	} else {
		goto L190
	}
L186:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v439)+8))
	if v441 == int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v444 = F_LookupNamespaceNoError(m, v441)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	if v444 != 0 {
		goto L185
	} else {
		goto L189
	}
L189:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v439)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v446
	goto L124
L190:
	;
	v1352 = int32(_a_F_ExecDropStmt_32)
	v1353 = v449
	goto L124
L191:
	;
	v461 = F_NameListToString(m, v349)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L5
	} else {
		goto L196
	}
L192:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v451)+8))
	if v453 == int32(0) {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v456 = F_LookupNamespaceNoError(m, v453)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	if v456 != 0 {
		goto L191
	} else {
		goto L195
	}
L195:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v451)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v458
	goto L124
L196:
	;
	v1352 = int32(_a_F_ExecDropStmt_33)
	v1353 = v461
	goto L124
L197:
	;
	v473 = F_NameListToString(m, v349)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L5
	} else {
		goto L202
	}
L198:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v463)+8))
	if v465 == int32(0) {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v468 = F_LookupNamespaceNoError(m, v465)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L5
	} else {
		goto L200
	}
L200:
	;
	if v468 != 0 {
		goto L197
	} else {
		goto L201
	}
L201:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v463)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v470
	goto L124
L202:
	;
	v1352 = int32(_a_F_ExecDropStmt_34)
	v1353 = v473
	goto L124
L203:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	if v487 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L204:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v478)+8))
	if v480 == int32(0) {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v483 = F_LookupNamespaceNoError(m, v480)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L5
	} else {
		goto L206
	}
L206:
	;
	if v483 != 0 {
		goto L203
	} else {
		goto L207
	}
L207:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v478)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v485
	goto L124
L208:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	v543 = F_makeRangeVarFromNameList(m, v542)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L5
	} else {
		goto L224
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_35)
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v536 = F_NameListToString(m, v535)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L5
	} else {
		goto L220
	}
L210:
	;
	v490 = int32(0)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	if v491 <= v490 {
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v495 = v490
	v500 = v491
	goto L212
L212:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v487)+12))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v506+v495<<(uint(int32(2))%32))))
	if v510 != 0 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	goto L209
L214:
	;
	v512 = F_LookupTypeNameOid(m, v510, int32(1))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L5
	} else {
		goto L217
	}
L215:
	;
	v517 = v500
	goto L216
L216:
	;
	v519 = v495 + int32(1)
	if v519 < v517 {
		v495 = v519
		v500 = v517
		goto L212
	} else {
		goto L219
	}
L217:
	;
	if v512 == int32(0) {
		goto L208
	} else {
		goto L218
	}
L218:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	v517 = v516
	goto L216
L219:
	;
	goto L213
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v536
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	v540 = F_TypeNameListToString(m, v539)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L5
	} else {
		goto L221
	}
L221:
	;
	v1526 = v540
	goto L119
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v557
	v1526 = int32(0)
	goto L119
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_27)
	v555 = F_TypeNameToString(m, v510)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L5
	} else {
		goto L228
	}
L224:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v543)+8))
	if v545 == int32(0) {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v548 = F_LookupNamespaceNoError(m, v545)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L5
	} else {
		goto L226
	}
L226:
	;
	if v548 != 0 {
		goto L223
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_5)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v543)+8))
	v557 = v552
	goto L222
L228:
	;
	v557 = v555
	goto L222
L229:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	if v570 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L230:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v561)+8))
	if v563 == int32(0) {
		goto L229
	} else {
		goto L231
	}
L231:
	;
	v566 = F_LookupNamespaceNoError(m, v563)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L5
	} else {
		goto L232
	}
L232:
	;
	if v566 != 0 {
		goto L229
	} else {
		goto L233
	}
L233:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v561)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v568
	goto L124
L234:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v593)+4))
	v626 = F_makeRangeVarFromNameList(m, v625)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L5
	} else {
		goto L250
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_36)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v619 = F_NameListToString(m, v618)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L5
	} else {
		goto L246
	}
L236:
	;
	v573 = int32(0)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	if v574 <= v573 {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v578 = v573
	v583 = v574
	goto L238
L238:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v570)+12))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v589+v578<<(uint(int32(2))%32))))
	if v593 != 0 {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	goto L235
L240:
	;
	v595 = F_LookupTypeNameOid(m, v593, int32(1))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L5
	} else {
		goto L243
	}
L241:
	;
	v600 = v583
	goto L242
L242:
	;
	v602 = v578 + int32(1)
	if v602 < v600 {
		v578 = v602
		v583 = v600
		goto L238
	} else {
		goto L245
	}
L243:
	;
	if v595 == int32(0) {
		goto L234
	} else {
		goto L244
	}
L244:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	v600 = v599
	goto L242
L245:
	;
	goto L239
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v619
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	v623 = F_TypeNameListToString(m, v622)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L5
	} else {
		goto L247
	}
L247:
	;
	v1526 = v623
	goto L119
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v640
	v1526 = int32(0)
	goto L119
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_27)
	v638 = F_TypeNameToString(m, v593)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L5
	} else {
		goto L254
	}
L250:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v626)+8))
	if v628 == int32(0) {
		goto L249
	} else {
		goto L251
	}
L251:
	;
	v631 = F_LookupNamespaceNoError(m, v628)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L5
	} else {
		goto L252
	}
L252:
	;
	if v631 != 0 {
		goto L249
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_5)
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v626)+8))
	v640 = v635
	goto L248
L254:
	;
	v640 = v638
	goto L248
L255:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	if v653 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L256:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v644)+8))
	if v646 == int32(0) {
		goto L255
	} else {
		goto L257
	}
L257:
	;
	v649 = F_LookupNamespaceNoError(m, v646)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L5
	} else {
		goto L258
	}
L258:
	;
	if v649 != 0 {
		goto L255
	} else {
		goto L259
	}
L259:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v644)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v651
	goto L124
L260:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v676)+4))
	v709 = F_makeRangeVarFromNameList(m, v708)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L5
	} else {
		goto L276
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_37)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v702 = F_NameListToString(m, v701)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L5
	} else {
		goto L272
	}
L262:
	;
	v656 = int32(0)
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v653)+4))
	if v657 <= v656 {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v661 = v656
	v666 = v657
	goto L264
L264:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v653)+12))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v672+v661<<(uint(int32(2))%32))))
	if v676 != 0 {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	goto L261
L266:
	;
	v678 = F_LookupTypeNameOid(m, v676, int32(1))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L5
	} else {
		goto L269
	}
L267:
	;
	v683 = v666
	goto L268
L268:
	;
	v685 = v661 + int32(1)
	if v685 < v683 {
		v661 = v685
		v666 = v683
		goto L264
	} else {
		goto L271
	}
L269:
	;
	if v678 == int32(0) {
		goto L260
	} else {
		goto L270
	}
L270:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v653)+4))
	v683 = v682
	goto L268
L271:
	;
	goto L265
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v702
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	v706 = F_TypeNameListToString(m, v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L5
	} else {
		goto L273
	}
L273:
	;
	v1526 = v706
	goto L119
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v723
	v1526 = int32(0)
	goto L119
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_27)
	v721 = F_TypeNameToString(m, v676)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L5
	} else {
		goto L280
	}
L276:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v709)+8))
	if v711 == int32(0) {
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v714 = F_LookupNamespaceNoError(m, v711)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L5
	} else {
		goto L278
	}
L278:
	;
	if v714 != 0 {
		goto L275
	} else {
		goto L279
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_5)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v709)+8))
	v723 = v718
	goto L274
L280:
	;
	v723 = v721
	goto L274
L281:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	if v736 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L282:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v727)+8))
	if v729 == int32(0) {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v732 = F_LookupNamespaceNoError(m, v729)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L5
	} else {
		goto L284
	}
L284:
	;
	if v732 != 0 {
		goto L281
	} else {
		goto L285
	}
L285:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v727)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v734
	goto L124
L286:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v759)+4))
	v792 = F_makeRangeVarFromNameList(m, v791)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L5
	} else {
		goto L302
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_38)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v785 = F_NameListToString(m, v784)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L5
	} else {
		goto L298
	}
L288:
	;
	v739 = int32(0)
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v736)+4))
	if v740 <= v739 {
		goto L287
	} else {
		goto L289
	}
L289:
	;
	v744 = v739
	v749 = v740
	goto L290
L290:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v736)+12))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v755+v744<<(uint(int32(2))%32))))
	if v759 != 0 {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	goto L287
L292:
	;
	v761 = F_LookupTypeNameOid(m, v759, int32(1))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L5
	} else {
		goto L295
	}
L293:
	;
	v766 = v749
	goto L294
L294:
	;
	v768 = v744 + int32(1)
	if v768 < v766 {
		v744 = v768
		v749 = v766
		goto L290
	} else {
		goto L297
	}
L295:
	;
	if v761 == int32(0) {
		goto L286
	} else {
		goto L296
	}
L296:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v736)+4))
	v766 = v765
	goto L294
L297:
	;
	goto L291
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v785
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	v789 = F_TypeNameListToString(m, v788)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L5
	} else {
		goto L299
	}
L299:
	;
	v1526 = v789
	goto L119
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v806
	v1526 = int32(0)
	goto L119
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_27)
	v804 = F_TypeNameToString(m, v759)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L5
	} else {
		goto L306
	}
L302:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v792)+8))
	if v794 == int32(0) {
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v797 = F_LookupNamespaceNoError(m, v794)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L5
	} else {
		goto L304
	}
L304:
	;
	if v797 != 0 {
		goto L301
	} else {
		goto L305
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_5)
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v792)+8))
	v806 = v801
	goto L300
L306:
	;
	v806 = v804
	goto L300
L307:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	if v819 == int32(0) {
		goto L313
	} else {
		goto L314
	}
L308:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v810)+8))
	if v812 == int32(0) {
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v815 = F_LookupNamespaceNoError(m, v812)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L5
	} else {
		goto L310
	}
L310:
	;
	if v815 != 0 {
		goto L307
	} else {
		goto L311
	}
L311:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v810)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v817
	goto L124
L312:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v842)+4))
	v870 = F_makeRangeVarFromNameList(m, v869)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L5
	} else {
		goto L327
	}
L313:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v867 = F_NameListToString(m, v866)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L5
	} else {
		goto L324
	}
L314:
	;
	v822 = int32(0)
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v819)+4))
	if v823 <= v822 {
		goto L313
	} else {
		goto L315
	}
L315:
	;
	v827 = v822
	v832 = v823
	goto L316
L316:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v819)+12))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v838+v827<<(uint(int32(2))%32))))
	if v842 != 0 {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	goto L313
L318:
	;
	v844 = F_LookupTypeNameOid(m, v842, int32(1))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L5
	} else {
		goto L321
	}
L319:
	;
	v849 = v832
	goto L320
L320:
	;
	v851 = v827 + int32(1)
	if v851 < v849 {
		v827 = v851
		v832 = v849
		goto L316
	} else {
		goto L323
	}
L321:
	;
	if v844 == int32(0) {
		goto L312
	} else {
		goto L322
	}
L322:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v819)+4))
	v849 = v848
	goto L320
L323:
	;
	goto L317
L324:
	;
	v1352 = int32(_a_F_ExecDropStmt_39)
	v1353 = v867
	goto L124
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v884
	v1526 = int32(0)
	goto L119
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_27)
	v882 = F_TypeNameToString(m, v842)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L5
	} else {
		goto L331
	}
L327:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v870)+8))
	if v872 == int32(0) {
		goto L326
	} else {
		goto L328
	}
L328:
	;
	v875 = F_LookupNamespaceNoError(m, v872)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L5
	} else {
		goto L329
	}
L329:
	;
	if v875 != 0 {
		goto L326
	} else {
		goto L330
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_5)
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v870)+8))
	v884 = v879
	goto L325
L331:
	;
	v884 = v882
	goto L325
L332:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+52)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v323)+164)) = v944
	v950 = F_list_make1_impl(m, int32(1), v323+int32(52))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L5
	} else {
		goto L345
	}
L333:
	;
	if v896 == int32(0) {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v900 = int32(0)
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v896)+4))
	if v901 <= v900 {
		goto L332
	} else {
		goto L335
	}
L335:
	;
	v905 = v900
	v910 = v901
	goto L336
L336:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v896)+12))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v916+v905<<(uint(int32(2))%32))))
	if v920 != 0 {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	goto L332
L338:
	;
	v922 = F_LookupTypeNameOid(m, v920, int32(1))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L5
	} else {
		goto L341
	}
L339:
	;
	v927 = v910
	goto L340
L340:
	;
	v929 = v905 + int32(1)
	if v929 < v927 {
		v905 = v929
		v910 = v927
		goto L336
	} else {
		goto L343
	}
L341:
	;
	if v922 == int32(0) {
		goto L120
	} else {
		goto L342
	}
L342:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v896)+4))
	v927 = v926
	goto L340
L343:
	;
	goto L337
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_40)
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1017)))
	v1019 = F_TypeNameToString(m, v1018)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L5
	} else {
		goto L364
	}
L345:
	;
	if v950 == int32(0) {
		goto L344
	} else {
		goto L346
	}
L346:
	;
	v954 = int32(0)
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v950)+4))
	if v955 <= v954 {
		goto L344
	} else {
		goto L347
	}
L347:
	;
	v959 = v954
	v964 = v955
	goto L348
L348:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v950)+12))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v970+v959<<(uint(int32(2))%32))))
	if v974 == int32(0) {
		v999 = v964
		goto L350
	} else {
		goto L351
	}
L349:
	;
	goto L344
L350:
	;
	v1001 = v959 + int32(1)
	if v1001 < v999 {
		v959 = v1001
		v964 = v999
		goto L348
	} else {
		goto L363
	}
L351:
	;
	v978 = F_LookupTypeNameOid(m, v974, int32(1))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L5
	} else {
		goto L352
	}
L352:
	;
	if v978 != 0 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v950)+4))
	v999 = v980
	goto L350
L354:
	;
	goto L355
L355:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v974)+4))
	v982 = F_makeRangeVarFromNameList(m, v981)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L5
	} else {
		goto L358
	}
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v996
	v1526 = int32(0)
	goto L119
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_27)
	v994 = F_TypeNameToString(m, v974)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L5
	} else {
		goto L362
	}
L358:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v982)+8))
	if v984 == int32(0) {
		goto L357
	} else {
		goto L359
	}
L359:
	;
	v987 = F_LookupNamespaceNoError(m, v984)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L5
	} else {
		goto L360
	}
L360:
	;
	if v987 != 0 {
		goto L357
	} else {
		goto L361
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_5)
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v982)+8))
	v996 = v991
	goto L356
L362:
	;
	v996 = v994
	goto L356
L363:
	;
	goto L349
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v1019
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+4))
	v1024 = F_TypeNameToString(m, v1023)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L5
	} else {
		goto L365
	}
L365:
	;
	v1526 = v1024
	goto L119
L366:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L5
	} else {
		goto L427
	}
L367:
	;
	v1028 = F_get_func_prokind(m, v369)
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L5
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	v1033 = m.G0
	v1035 = v1033 - int32(32)
	m.G0 = v1035
	v1038 = v323 + int32(152)
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1038)))
	v1041 = *(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[0]))
	if v1041 != 0 {
		goto L376
	} else {
		goto L377
	}
L370:
	;
	if v1028 == int32(97) {
		goto L366
	} else {
		goto L371
	}
L371:
	;
	goto L369
L372:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[1]))
	if v1129 != 0 {
		goto L411
	} else {
		goto L412
	}
L373:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L5
	} else {
		goto L407
	}
L374:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L5
	} else {
		goto L404
	}
L375:
	;
	v1114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1107)+24)))
	if v1114 == int32(0) {
		goto L397
	} else {
		goto L398
	}
L376:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+4))
	if v1042 == v1039 {
		v1107 = v1041
		goto L375
	} else {
		goto L379
	}
L377:
	;
	goto L378
L378:
	;
	v1049 = int32(0)
	goto L381
L379:
	;
	goto L378
L380:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[0])) = v1100
	v1107 = v1100
	goto L375
L381:
	;
	v1058 = v1049 * int32(40)
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+uint32(_c_F_ExecDropStmt[2])))
	if v1039 != v1061 {
		goto L383
	} else {
		goto L384
	}
L382:
	;
	v1100 = v1058 + int32(_a_F_ExecDropStmt_41)
	goto L380
L383:
	;
	if v1049 == int32(36) {
		goto L374
	} else {
		goto L386
	}
L384:
	;
	goto L385
L385:
	;
	goto L382
L386:
	;
	v1068 = (v1049 | int32(1)) * int32(40)
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1068)+uint32(_c_F_ExecDropStmt[2])))
	if v1039 == v1071 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1100 = v1068 + int32(_a_F_ExecDropStmt_41)
	goto L380
L388:
	;
	goto L389
L389:
	;
	v1078 = (v1049 | int32(2)) * int32(40)
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+uint32(_c_F_ExecDropStmt[2])))
	if v1039 == v1081 {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	v1100 = v1078 + int32(_a_F_ExecDropStmt_41)
	goto L380
L391:
	;
	goto L392
L392:
	;
	v1088 = (v1049 | int32(3)) * int32(40)
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+uint32(_c_F_ExecDropStmt[2])))
	if v1039 == v1091 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1100 = v1088 + int32(_a_F_ExecDropStmt_41)
	goto L380
L394:
	;
	v1049 = v1049 + int32(4)
	goto L381
L396:
	;
	m.G0 = v1035 + int32(32)
	goto L372
L397:
	;
	v1129 = int32(0)
	goto L396
L398:
	;
	goto L399
L399:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+12))
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+4))
	v1120 = F_SearchSysCache1(m, v1118, v1119)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L5
	} else {
		goto L400
	}
L400:
	;
	if v1120 == int32(0) {
		goto L373
	} else {
		goto L401
	}
L401:
	;
	v1124 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1107)+24)))
	v1125 = F_SysCacheGetAttrNotNull(m, v1118, v1120, v1124)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L5
	} else {
		goto L402
	}
L402:
	;
	F_ReleaseCatCache(m, v1120)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L5
	} else {
		goto L403
	}
L403:
	;
	v1129 = v1125
	goto L396
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1035)+16)) = v1039
	F_errmsg_internal(m, int32(_a_F_ExecDropStmt_42), v1035+int32(16))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L5
	} else {
		goto L405
	}
L405:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_43), int32(2777), int32(_a_F_ExecDropStmt_44))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L5
	} else {
		goto L406
	}
L406:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L407:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1035)+4)) = v1153
	*(*int32)(unsafe.Add(mBase, uint32(v1035))) = v1118
	F_errmsg_internal(m, int32(_a_F_ExecDropStmt_45), v1035)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L5
	} else {
		goto L408
	}
L408:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_43), int32(2593), int32(_a_F_ExecDropStmt_46))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L5
	} else {
		goto L409
	}
L409:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L410:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v323)+148))
	if v1213 != 0 {
		goto L422
	} else {
		goto L423
	}
L411:
	;
	v1167 = F_object_ownercheck(m, int32(2615), v1129, v1165)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L5
	} else {
		goto L414
	}
L412:
	;
	goto L413
L413:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+112)) = v1201
	v1203 = *(*int64)(unsafe.Add(mBase, uint32(v323)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v323)+104)) = v1203
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v323)+148))
	F_check_object_ownership(m, v1165, v1200, v323+int32(104), v349, v1207)
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L5
	} else {
		goto L421
	}
L414:
	;
	if v1167 == int32(0) {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[1]))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+128)) = v1174
	v1176 = *(*int64)(unsafe.Add(mBase, uint32(v323)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v323)+120)) = v1176
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v323)+148))
	F_check_object_ownership(m, v1172, v1173, v323+int32(120), v349, v1180)
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L5
	} else {
		goto L418
	}
L416:
	;
	goto L417
L417:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[3]))
	goto L419
L418:
	;
	goto L417
L419:
	;
	if base.B2i32(v1187 != int32(0))&base.B2i32(v1129 == v1187) == int32(0) {
		goto L410
	} else {
		goto L420
	}
L420:
	;
	v1194 = int32(_a_F_ExecDropStmt_47)
	v1196 = *(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[4])) = v1196 | int32(1)
	goto L410
L421:
	;
	goto L410
L422:
	;
	F_sequence_close(m, v1213, int32(0))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L5
	} else {
		goto L425
	}
L423:
	;
	goto L424
L424:
	;
	F_add_exact_object_address(m, v323+int32(152), v325)
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L5
	} else {
		goto L426
	}
L425:
	;
	goto L424
L426:
	;
	goto L116
L427:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L5
	} else {
		goto L428
	}
L428:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v1229 = F_NameListToString(m, v1228)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L5
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+96)) = v1229
	F_errmsg(m, int32(_a_F_ExecDropStmt_48), v323+int32(96))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L5
	} else {
		goto L430
	}
L430:
	;
	F_errhint(m, int32(_a_F_ExecDropStmt_49), int32(0))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L5
	} else {
		goto L431
	}
L431:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_18), int32(98), int32(_a_F_ExecDropStmt_50))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L5
	} else {
		goto L432
	}
L432:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+80)) = v367
	F_errmsg_internal(m, int32(_a_F_ExecDropStmt_51), v323+int32(80))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L5
	} else {
		goto L434
	}
L434:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_18), int32(512), int32(_a_F_ExecDropStmt_19))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L5
	} else {
		goto L435
	}
L435:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+64)) = v367
	F_errmsg_internal(m, int32(_a_F_ExecDropStmt_51), v323-int32(-64))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L5
	} else {
		goto L437
	}
L437:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_18), int32(496), int32(_a_F_ExecDropStmt_19))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L5
	} else {
		goto L438
	}
L438:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_52)
	v1294 = F_NameListToString(m, v1281)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L5
	} else {
		goto L445
	}
L440:
	;
	v1283 = F_makeRangeVarFromNameList(m, v1281)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L5
	} else {
		goto L441
	}
L441:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1283)+8))
	if v1285 == int32(0) {
		goto L439
	} else {
		goto L442
	}
L442:
	;
	v1288 = F_LookupNamespaceNoError(m, v1285)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L5
	} else {
		goto L443
	}
L443:
	;
	if v1288 != 0 {
		goto L439
	} else {
		goto L444
	}
L444:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1283)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v1290
	goto L124
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v1294
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1297)))
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v1298)+4))
	v1526 = v1299
	goto L119
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_53)
	v1314 = F_NameListToString(m, v1301)
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L5
	} else {
		goto L452
	}
L447:
	;
	v1303 = F_makeRangeVarFromNameList(m, v1301)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L5
	} else {
		goto L448
	}
L448:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+8))
	if v1305 == int32(0) {
		goto L446
	} else {
		goto L449
	}
L449:
	;
	v1308 = F_LookupNamespaceNoError(m, v1305)
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L5
	} else {
		goto L450
	}
L450:
	;
	if v1308 != 0 {
		goto L446
	} else {
		goto L451
	}
L451:
	;
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+8))
	v1352 = int32(_a_F_ExecDropStmt_5)
	v1353 = v1310
	goto L124
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v1314
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1317)))
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1318)+4))
	v1526 = v1319
	goto L119
L453:
	;
	if v1329 != 0 {
		v1526 = int32(0)
		goto L119
	} else {
		goto L454
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_54)
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1333+v1334<<(uint(int32(2))%32)-int32(4))))
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1340)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v1341
	v1345 = F_list_copy_head(m, v349, v1334-int32(1))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L5
	} else {
		goto L455
	}
L455:
	;
	v1347 = F_NameListToString(m, v1345)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L5
	} else {
		goto L456
	}
L456:
	;
	v1526 = v1347
	goto L119
L457:
	;
	if v1369 != 0 {
		v1526 = int32(0)
		goto L119
	} else {
		goto L458
	}
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_55)
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1373+v1374<<(uint(int32(2))%32)-int32(4))))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v1381
	v1385 = F_list_copy_head(m, v349, v1374-int32(1))
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L5
	} else {
		goto L459
	}
L459:
	;
	v1387 = F_NameListToString(m, v1385)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L5
	} else {
		goto L460
	}
L460:
	;
	v1526 = v1387
	goto L119
L461:
	;
	if v1394 != 0 {
		v1526 = int32(0)
		goto L119
	} else {
		goto L462
	}
L462:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_56)
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1398+v1399<<(uint(int32(2))%32)-int32(4))))
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v1406
	v1410 = F_list_copy_head(m, v349, v1399-int32(1))
	mBase = m.M
	v1411 = m.ExcPending
	if v1411 != 0 {
		goto L5
	} else {
		goto L463
	}
L463:
	;
	v1412 = F_NameListToString(m, v1410)
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L5
	} else {
		goto L464
	}
L464:
	;
	v1526 = v1412
	goto L119
L465:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+4))
	v1479 = F_makeRangeVarFromNameList(m, v1478)
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L5
	} else {
		goto L481
	}
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_57)
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1470)))
	v1472 = F_TypeNameToString(m, v1471)
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L5
	} else {
		goto L478
	}
L467:
	;
	if v1421 == int32(0) {
		goto L466
	} else {
		goto L468
	}
L468:
	;
	v1425 = int32(0)
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1421)+4))
	if v1426 <= v1425 {
		goto L466
	} else {
		goto L469
	}
L469:
	;
	v1430 = v1425
	v1435 = v1426
	goto L470
L470:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1421)+12))
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1441+v1430<<(uint(int32(2))%32))))
	if v1445 != 0 {
		goto L472
	} else {
		goto L473
	}
L471:
	;
	goto L466
L472:
	;
	v1447 = F_LookupTypeNameOid(m, v1445, int32(1))
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L5
	} else {
		goto L475
	}
L473:
	;
	v1452 = v1435
	goto L474
L474:
	;
	v1454 = v1430 + int32(1)
	if v1454 < v1452 {
		v1430 = v1454
		v1435 = v1452
		goto L470
	} else {
		goto L477
	}
L475:
	;
	if v1447 == int32(0) {
		goto L465
	} else {
		goto L476
	}
L476:
	;
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1421)+4))
	v1452 = v1451
	goto L474
L477:
	;
	goto L471
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v1472
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1475)+4))
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1476)+4))
	v1526 = v1477
	goto L119
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v1493
	v1526 = int32(0)
	goto L119
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_27)
	v1491 = F_TypeNameToString(m, v1445)
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L5
	} else {
		goto L485
	}
L481:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1479)+8))
	if v1481 == int32(0) {
		goto L480
	} else {
		goto L482
	}
L482:
	;
	v1484 = F_LookupNamespaceNoError(m, v1481)
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L5
	} else {
		goto L483
	}
L483:
	;
	if v1484 != 0 {
		goto L480
	} else {
		goto L484
	}
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_5)
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1479)+8))
	v1493 = v1488
	goto L479
L485:
	;
	v1493 = v1491
	goto L479
L486:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+172)) = v1511
	v1526 = int32(0)
	goto L119
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_27)
	v1509 = F_TypeNameToString(m, v920)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L5
	} else {
		goto L492
	}
L488:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v1497)+8))
	if v1499 == int32(0) {
		goto L487
	} else {
		goto L489
	}
L489:
	;
	v1502 = F_LookupNamespaceNoError(m, v1499)
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L5
	} else {
		goto L490
	}
L490:
	;
	if v1502 != 0 {
		goto L487
	} else {
		goto L491
	}
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+136)) = int32(_a_F_ExecDropStmt_5)
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1497)+8))
	v1511 = v1506
	goto L486
L492:
	;
	v1511 = v1509
	goto L486
L493:
	;
	if v1526 == int32(0) {
		v1547 = v1527
		goto L118
	} else {
		goto L494
	}
L494:
	;
	v1534 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L5
	} else {
		goto L495
	}
L495:
	;
	if v1534 == int32(0) {
		goto L116
	} else {
		goto L496
	}
L496:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v323)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+32)) = v1538
	*(*int32)(unsafe.Add(mBase, uint32(v323)+36)) = v1526
	F_errmsg(m, v1527, v323+int32(32))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L5
	} else {
		goto L497
	}
L497:
	;
	v1583 = int32(523)
	goto L117
L498:
	;
	if v1560 == int32(0) {
		goto L116
	} else {
		goto L499
	}
L499:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v323)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+16)) = v1564
	F_errmsg(m, v1547, v323+int32(16))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L5
	} else {
		goto L500
	}
L500:
	;
	v1583 = int32(521)
	goto L117
L501:
	;
	goto L116
L502:
	;
	goto L114
L503:
	;
	F_free_object_addresses(m, v325)
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L5
	} else {
		goto L504
	}
L504:
	;
	m.G0 = v323 + int32(176)
	goto L108
L505:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v367
	F_errmsg_internal(m, int32(_a_F_ExecDropStmt_58), v323)
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L5
	} else {
		goto L506
	}
L506:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_18), int32(518), int32(_a_F_ExecDropStmt_19))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L5
	} else {
		goto L507
	}
L507:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_mark_stmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v10 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_plpgsql_statement_tree_walker_impl_x2especialized_x2e2(m, l0, l1)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L17
	} else {
		goto L43
	}
L6:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v90 = F_bms_copy(m, l1)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L17
	} else {
		goto L33
	}
L9:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v12 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v42 = v10
	goto L11
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v43 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L12:
	;
	v17 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v36 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L15:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v17<<(uint(int32(2))%32))))
	F_mark_stmt(m, v24, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	return
L18:
	;
	v29 = v17 + int32(1)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v29 < v30 {
		v17 = v29
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v42 = v36
	goto L11
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v46 <= int32(0) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v52 = int32(0)
	goto L23
L23:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54+v52<<(uint(int32(2))%32))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	if v59 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L1
L25:
	;
	v87 = v52 + int32(1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v87 < v88 {
		v52 = v87
		goto L23
	} else {
		goto L32
	}
L26:
	;
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v63 <= v62 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v66 = v62
	goto L28
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v66<<(uint(int32(2))%32))))
	F_mark_stmt(m, v74, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L17
	} else {
		goto L30
	}
L29:
	;
	goto L25
L30:
	;
	v79 = v66 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v79 < v80 {
		v66 = v79
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	goto L24
L33:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v92 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v97 = int32(0)
	v98 = v90
	goto L37
L35:
	;
	v113 = v90
	goto L36
L36:
	;
	F_plpgsql_statement_tree_walker_impl_x2especialized_x2e2(m, l0, v113)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L17
	} else {
		goto L41
	}
L37:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v97<<(uint(int32(2))%32))))
	v105 = F_bms_add_member(m, v98, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L17
	} else {
		goto L39
	}
L38:
	;
	v113 = v105
	goto L36
L39:
	;
	v108 = v97 + int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v108 < v109 {
		v97 = v108
		v98 = v105
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	F_bms_free(m, v113)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L17
	} else {
		goto L42
	}
L42:
	;
	return
L43:
	;
	goto L1
}
func F_stmt_requires_parse_analysis(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	switch v5 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v9 = int32(1)
	default:
		v9 = int32(0)
	}
	return v9
}
