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
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
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
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int64
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
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
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
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
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
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
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1116 int32
	_ = v1116
	var v1122 int32
	_ = v1122
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int64
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1178 int64
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1232 int32
	_ = v1232
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
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
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1416 int32
	_ = v1416
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1519 int32
	_ = v1519
	var v1522 int32
	_ = v1522
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1544 int32
	_ = v1544
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1614 int32
	_ = v1614
	var v1618 int32
	_ = v1618
	var v1623 int32
	_ = v1623
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
	v319 = m.G0
	v321 = v319 - int32(176)
	m.G0 = v321
	v323 = F_new_object_addresses(m)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L5
	} else {
		goto L105
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
	return
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L5
	} else {
		goto L101
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L5
	} else {
		goto L97
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L5
	} else {
		goto L93
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L5
	} else {
		goto L89
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
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
	if v68 == int32(0) {
		v214 = v3
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_performMultipleDeletions(m, v66, v222, v214)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L83
	}
L31:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v71 <= int32(0) {
		v214 = v3
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v80 = v3
	v85 = v3
	goto L33
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v85<<(uint(int32(2))%32))))
	v93 = F_makeRangeVarFromNameList(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L35
	}
L34:
	;
	v214 = v204
	goto L30
L35:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+96)) = uint8(v65)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+104)) = int64(0)
	if v98 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v103 = int32(4)
	goto L39
L38:
	;
	v103 = int32(8)
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+100)) = v103
	v109 = F_RangeVarGetRelidExtended(m, v93, v41, int32(1), int32(573), v24+int32(96))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L41
	}
L40:
	;
	v207 = v85 + int32(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v207 < v208 {
		v80 = v204
		v85 = v207
		goto L33
	} else {
		goto L82
	}
L41:
	;
	if v109 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	if v114 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+113)))
	if v174 != int32(116) {
		goto L67
	} else {
		goto L68
	}
L45:
	;
	switch v65 - int32(83) {
	case 0:
		goto L61
	default:
		v204 = v80
		goto L40
	case 16:
		goto L57
	case 19:
		goto L56
	case 22:
		goto L58
	case 26:
		goto L59
	case 29:
		goto L55
	case 31:
		v149 = int32(_a_F_ExecDropStmt_4)
		goto L54
	case 35:
		goto L60
	}
L46:
	;
	v117 = F_LookupNamespaceNoError(m, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	if v117 != 0 {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	if v113&int32(1) == int32(0) {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	v125 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	if v125 == int32(0) {
		v204 = v80
		goto L40
	} else {
		goto L51
	}
L51:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v129
	F_errmsg(m, int32(_a_F_ExecDropStmt_5), v24+int32(48))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1480), int32(_a_F_ExecDropStmt_6))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	v204 = v80
	goto L40
L54:
	;
	if v113&int32(1) == int32(0) {
		goto L9
	} else {
		goto L62
	}
L55:
	;
	v149 = int32(_a_F_ExecDropStmt_7)
	goto L54
L56:
	;
	v149 = int32(_a_F_ExecDropStmt_8)
	goto L54
L57:
	;
	v149 = int32(_a_F_ExecDropStmt_9)
	goto L54
L58:
	;
	v149 = int32(_a_F_ExecDropStmt_10)
	goto L54
L59:
	;
	v149 = int32(_a_F_ExecDropStmt_11)
	goto L54
L60:
	;
	v149 = int32(_a_F_ExecDropStmt_12)
	goto L54
L61:
	;
	v149 = int32(_a_F_ExecDropStmt_13)
	goto L54
L62:
	;
	v156 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	if v156 == int32(0) {
		v204 = v80
		goto L40
	} else {
		goto L64
	}
L64:
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
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1497), int32(_a_F_ExecDropStmt_6))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	v204 = v80
	goto L40
L67:
	;
	v177 = v80 | int32(2)
	goto L69
L68:
	;
	v177 = v80
	goto L69
L69:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v178 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v179 = v177
	goto L72
L71:
	;
	v179 = v80
	goto L72
L72:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+112)))
	if v183 == int32(73) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v186 = v179 & int32(2)
	goto L75
L74:
	;
	v186 = int32(0)
	goto L75
L75:
	;
	if v186 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	if v183 == int32(73) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	v192 = F_find_all_inheritors(m, v189, v190, int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+124)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+120)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v24)+116)) = int32(1259)
	F_add_exact_object_address(m, v24+int32(116), v66)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L5
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	v204 = v179
	goto L40
L82:
	;
	goto L34
L83:
	;
	F_free_object_addresses(m, v66)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	m.G0 = v24 + int32(128)
	goto L7
L85:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(_a_F_ExecDropStmt_14), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1559), int32(_a_F_ExecDropStmt_3))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
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
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(_a_F_ExecDropStmt_15), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1563), int32(_a_F_ExecDropStmt_3))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
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
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v269
	F_errmsg(m, int32(_a_F_ExecDropStmt_16), v24-int32(-64))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1474), int32(_a_F_ExecDropStmt_6))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
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
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	F_errcode(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v289
	F_errmsg(m, v288, v24+int32(32))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1493), int32(_a_F_ExecDropStmt_6))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
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
	v306 = m.ExcPending
	if v306 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v307
	F_errmsg(m, int32(_a_F_ExecDropStmt_17), v24+int32(80))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_2), int32(1668), int32(_a_F_ExecDropStmt_3))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
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
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v325 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	return
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L5
	} else {
		goto L500
	}
L108:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_performMultipleDeletions(m, v323, v1590, int32(0))
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L5
	} else {
		goto L498
	}
L109:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	if v328 <= int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v341 = v3
	goto L111
L111:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v325)+12))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v343+v341<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+148)) = int32(0)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	F_get_object_address(m, v321+int32(136), v352, v347, v321+int32(148), int32(8), v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L113
	}
L112:
	;
	goto L108
L113:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v321)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+160)) = v359
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v321)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v321)+152)) = v361
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v321)+156))
	if v365 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L114:
	;
	v1575 = v341 + int32(1)
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	if v1575 < v1576 {
		v341 = v1575
		goto L111
	} else {
		goto L497
	}
L115:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_18), v1558, int32(_a_F_ExecDropStmt_19))
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L5
	} else {
		goto L496
	}
L116:
	;
	v1535 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L5
	} else {
		goto L493
	}
L117:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v321)+136))
	if v1502 == int32(0) {
		goto L107
	} else {
		goto L488
	}
L118:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v916)+4))
	v1472 = F_makeRangeVarFromNameList(m, v1471)
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L5
	} else {
		goto L483
	}
L119:
	;
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1389)))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+60)) = v1390
	*(*int32)(unsafe.Add(mBase, uint32(v321)+168)) = v1390
	v1396 = F_list_make1_impl(m, int32(1), v321+int32(60))
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L5
	} else {
		goto L462
	}
L120:
	;
	v1369 = F_owningrel_does_not_exist_skipping(m, v347, v321+int32(136), v321+int32(172))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L5
	} else {
		goto L456
	}
L121:
	;
	v1344 = F_owningrel_does_not_exist_skipping(m, v347, v321+int32(136), v321+int32(172))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L5
	} else {
		goto L452
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v1328
	v1522 = v1327
	goto L116
L123:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v1327 = int32(_a_F_ExecDropStmt_20)
	v1328 = v1324
	goto L122
L124:
	;
	v1304 = F_owningrel_does_not_exist_skipping(m, v347, v321+int32(136), v321+int32(172))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L5
	} else {
		goto L448
	}
L125:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v1327 = int32(_a_F_ExecDropStmt_21)
	v1328 = v1297
	goto L122
L126:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v1327 = int32(_a_F_ExecDropStmt_22)
	v1328 = v1295
	goto L122
L127:
	;
	v1276 = F_list_copy_tail(m, v347, int32(1))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L5
	} else {
		goto L442
	}
L128:
	;
	v1256 = F_list_copy_tail(m, v347, int32(1))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L5
	} else {
		goto L435
	}
L129:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v1327 = int32(_a_F_ExecDropStmt_23)
	v1328 = v1253
	goto L122
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L5
	} else {
		goto L431
	}
L131:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L5
	} else {
		goto L428
	}
L132:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v1327 = int32(_a_F_ExecDropStmt_24)
	v1328 = v1221
	goto L122
L133:
	;
	v368 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v368
	switch v363 {
	case 0:
		goto L132
	case 1:
		goto L139
	case 2, 3, 4, 10, 11, 13, 22, 27, 31, 32, 40, 50:
		goto L131
	case 5:
		goto L136
	case 6, 9, 18, 20, 23, 33, 37, 38, 41, 42, 51:
		goto L130
	case 7:
		goto L151
	case 8:
		goto L150
	case 12, 49:
		goto L152
	case 14:
		goto L123
	case 15:
		goto L143
	case 16:
		goto L125
	case 17:
		goto L126
	case 19:
		goto L142
	case 21:
		goto L137
	case 24:
		goto L127
	case 25:
		goto L138
	case 26:
		goto L128
	case 28:
		goto L121
	case 29:
		goto L141
	case 30:
		goto L129
	case 34:
		goto L140
	case 35:
		goto L124
	case 36:
		goto L149
	case 39:
		goto L148
	case 43:
		goto L119
	case 44:
		goto L120
	case 45:
		goto L144
	case 46:
		goto L146
	case 47:
		goto L147
	case 48:
		goto L145
	default:
		goto L107
	}
L134:
	;
	goto L135
L135:
	;
	if v363 == int32(19) {
		goto L365
	} else {
		goto L366
	}
L136:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v885)))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+56)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v321)+168)) = v886
	v892 = F_list_make1_impl(m, int32(1), v321+int32(56))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L5
	} else {
		goto L331
	}
L137:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v1327 = int32(_a_F_ExecDropStmt_25)
	v1328 = v883
	goto L122
L138:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v806 = F_makeRangeVarFromNameList(m, v805)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L5
	} else {
		goto L306
	}
L139:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v723 = F_makeRangeVarFromNameList(m, v722)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L5
	} else {
		goto L280
	}
L140:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v640 = F_makeRangeVarFromNameList(m, v639)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L5
	} else {
		goto L254
	}
L141:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v557 = F_makeRangeVarFromNameList(m, v556)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L5
	} else {
		goto L228
	}
L142:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v474 = F_makeRangeVarFromNameList(m, v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L5
	} else {
		goto L202
	}
L143:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v1327 = int32(_a_F_ExecDropStmt_26)
	v1328 = v471
	goto L122
L144:
	;
	v459 = F_makeRangeVarFromNameList(m, v347)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L5
	} else {
		goto L196
	}
L145:
	;
	v447 = F_makeRangeVarFromNameList(m, v347)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L5
	} else {
		goto L190
	}
L146:
	;
	v435 = F_makeRangeVarFromNameList(m, v347)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L5
	} else {
		goto L184
	}
L147:
	;
	v423 = F_makeRangeVarFromNameList(m, v347)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L5
	} else {
		goto L178
	}
L148:
	;
	v411 = F_makeRangeVarFromNameList(m, v347)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L5
	} else {
		goto L172
	}
L149:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v409
	goto L122
L150:
	;
	v397 = F_makeRangeVarFromNameList(m, v347)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L5
	} else {
		goto L166
	}
L151:
	;
	v385 = F_makeRangeVarFromNameList(m, v347)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L5
	} else {
		goto L160
	}
L152:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v373 = F_makeRangeVarFromNameList(m, v372)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L5
	} else {
		goto L154
	}
L153:
	;
	v383 = F_TypeNameToString(m, v347)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L5
	} else {
		goto L158
	}
L154:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	if v375 == int32(0) {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v378 = F_LookupNamespaceNoError(m, v375)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L5
	} else {
		goto L156
	}
L156:
	;
	if v378 != 0 {
		goto L153
	} else {
		goto L157
	}
L157:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v380
	goto L122
L158:
	;
	v1327 = int32(_a_F_ExecDropStmt_27)
	v1328 = v383
	goto L122
L159:
	;
	v395 = F_NameListToString(m, v347)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L5
	} else {
		goto L164
	}
L160:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v385)+8))
	if v387 == int32(0) {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v390 = F_LookupNamespaceNoError(m, v387)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L162
	}
L162:
	;
	if v390 != 0 {
		goto L159
	} else {
		goto L163
	}
L163:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v385)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v392
	goto L122
L164:
	;
	v1327 = int32(_a_F_ExecDropStmt_28)
	v1328 = v395
	goto L122
L165:
	;
	v407 = F_NameListToString(m, v347)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L5
	} else {
		goto L170
	}
L166:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v397)+8))
	if v399 == int32(0) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v402 = F_LookupNamespaceNoError(m, v399)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L5
	} else {
		goto L168
	}
L168:
	;
	if v402 != 0 {
		goto L165
	} else {
		goto L169
	}
L169:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v397)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v404
	goto L122
L170:
	;
	v1327 = int32(_a_F_ExecDropStmt_29)
	v1328 = v407
	goto L122
L171:
	;
	v421 = F_NameListToString(m, v347)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L5
	} else {
		goto L176
	}
L172:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v411)+8))
	if v413 == int32(0) {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v416 = F_LookupNamespaceNoError(m, v413)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L5
	} else {
		goto L174
	}
L174:
	;
	if v416 != 0 {
		goto L171
	} else {
		goto L175
	}
L175:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v411)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v418
	goto L122
L176:
	;
	v1327 = int32(_a_F_ExecDropStmt_30)
	v1328 = v421
	goto L122
L177:
	;
	v433 = F_NameListToString(m, v347)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L5
	} else {
		goto L182
	}
L178:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v423)+8))
	if v425 == int32(0) {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v428 = F_LookupNamespaceNoError(m, v425)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L5
	} else {
		goto L180
	}
L180:
	;
	if v428 != 0 {
		goto L177
	} else {
		goto L181
	}
L181:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v423)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v430
	goto L122
L182:
	;
	v1327 = int32(_a_F_ExecDropStmt_31)
	v1328 = v433
	goto L122
L183:
	;
	v445 = F_NameListToString(m, v347)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L5
	} else {
		goto L188
	}
L184:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v435)+8))
	if v437 == int32(0) {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v440 = F_LookupNamespaceNoError(m, v437)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L5
	} else {
		goto L186
	}
L186:
	;
	if v440 != 0 {
		goto L183
	} else {
		goto L187
	}
L187:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v435)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v442
	goto L122
L188:
	;
	v1327 = int32(_a_F_ExecDropStmt_32)
	v1328 = v445
	goto L122
L189:
	;
	v457 = F_NameListToString(m, v347)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L5
	} else {
		goto L194
	}
L190:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v447)+8))
	if v449 == int32(0) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v452 = F_LookupNamespaceNoError(m, v449)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	if v452 != 0 {
		goto L189
	} else {
		goto L193
	}
L193:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v447)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v454
	goto L122
L194:
	;
	v1327 = int32(_a_F_ExecDropStmt_33)
	v1328 = v457
	goto L122
L195:
	;
	v469 = F_NameListToString(m, v347)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L5
	} else {
		goto L200
	}
L196:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v459)+8))
	if v461 == int32(0) {
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v464 = F_LookupNamespaceNoError(m, v461)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	if v464 != 0 {
		goto L195
	} else {
		goto L199
	}
L199:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v459)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v466
	goto L122
L200:
	;
	v1327 = int32(_a_F_ExecDropStmt_34)
	v1328 = v469
	goto L122
L201:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	if v483 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L202:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v474)+8))
	if v476 == int32(0) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v479 = F_LookupNamespaceNoError(m, v476)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L5
	} else {
		goto L204
	}
L204:
	;
	if v479 != 0 {
		goto L201
	} else {
		goto L205
	}
L205:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v474)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v481
	goto L122
L206:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	v539 = F_makeRangeVarFromNameList(m, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L5
	} else {
		goto L222
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_35)
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v532 = F_NameListToString(m, v531)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L5
	} else {
		goto L218
	}
L208:
	;
	v486 = int32(0)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v483)+4))
	if v487 <= v486 {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v491 = v486
	v494 = v487
	goto L210
L210:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v483)+12))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v502+v491<<(uint(int32(2))%32))))
	if v506 != 0 {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	goto L207
L212:
	;
	v508 = F_LookupTypeNameOid(m, v506, int32(1))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L5
	} else {
		goto L215
	}
L213:
	;
	v513 = v494
	goto L214
L214:
	;
	v515 = v491 + int32(1)
	if v515 < v513 {
		v491 = v515
		v494 = v513
		goto L210
	} else {
		goto L217
	}
L215:
	;
	if v508 == int32(0) {
		goto L206
	} else {
		goto L216
	}
L216:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v483)+4))
	v513 = v512
	goto L214
L217:
	;
	goto L211
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v532
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	v536 = F_TypeNameListToString(m, v535)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L5
	} else {
		goto L219
	}
L219:
	;
	v1501 = v536
	goto L117
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v553
	v1501 = int32(0)
	goto L117
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_27)
	v551 = F_TypeNameToString(m, v506)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L5
	} else {
		goto L226
	}
L222:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v539)+8))
	if v541 == int32(0) {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v544 = F_LookupNamespaceNoError(m, v541)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L5
	} else {
		goto L224
	}
L224:
	;
	if v544 != 0 {
		goto L221
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_5)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v539)+8))
	v553 = v548
	goto L220
L226:
	;
	v553 = v551
	goto L220
L227:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	if v566 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L228:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v557)+8))
	if v559 == int32(0) {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v562 = F_LookupNamespaceNoError(m, v559)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L5
	} else {
		goto L230
	}
L230:
	;
	if v562 != 0 {
		goto L227
	} else {
		goto L231
	}
L231:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v557)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v564
	goto L122
L232:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v589)+4))
	v622 = F_makeRangeVarFromNameList(m, v621)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L5
	} else {
		goto L248
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_36)
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v615 = F_NameListToString(m, v614)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L5
	} else {
		goto L244
	}
L234:
	;
	v569 = int32(0)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	if v570 <= v569 {
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v574 = v569
	v577 = v570
	goto L236
L236:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v585+v574<<(uint(int32(2))%32))))
	if v589 != 0 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	goto L233
L238:
	;
	v591 = F_LookupTypeNameOid(m, v589, int32(1))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L5
	} else {
		goto L241
	}
L239:
	;
	v596 = v577
	goto L240
L240:
	;
	v598 = v574 + int32(1)
	if v598 < v596 {
		v574 = v598
		v577 = v596
		goto L236
	} else {
		goto L243
	}
L241:
	;
	if v591 == int32(0) {
		goto L232
	} else {
		goto L242
	}
L242:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	v596 = v595
	goto L240
L243:
	;
	goto L237
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v615
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	v619 = F_TypeNameListToString(m, v618)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L5
	} else {
		goto L245
	}
L245:
	;
	v1501 = v619
	goto L117
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v636
	v1501 = int32(0)
	goto L117
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_27)
	v634 = F_TypeNameToString(m, v589)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L5
	} else {
		goto L252
	}
L248:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v622)+8))
	if v624 == int32(0) {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v627 = F_LookupNamespaceNoError(m, v624)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L5
	} else {
		goto L250
	}
L250:
	;
	if v627 != 0 {
		goto L247
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_5)
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v622)+8))
	v636 = v631
	goto L246
L252:
	;
	v636 = v634
	goto L246
L253:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	if v649 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L254:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v640)+8))
	if v642 == int32(0) {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	v645 = F_LookupNamespaceNoError(m, v642)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L5
	} else {
		goto L256
	}
L256:
	;
	if v645 != 0 {
		goto L253
	} else {
		goto L257
	}
L257:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v640)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v647
	goto L122
L258:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	v705 = F_makeRangeVarFromNameList(m, v704)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L5
	} else {
		goto L274
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_37)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v698 = F_NameListToString(m, v697)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L5
	} else {
		goto L270
	}
L260:
	;
	v652 = int32(0)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v649)+4))
	if v653 <= v652 {
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v657 = v652
	v660 = v653
	goto L262
L262:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v649)+12))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v668+v657<<(uint(int32(2))%32))))
	if v672 != 0 {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	goto L259
L264:
	;
	v674 = F_LookupTypeNameOid(m, v672, int32(1))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L5
	} else {
		goto L267
	}
L265:
	;
	v679 = v660
	goto L266
L266:
	;
	v681 = v657 + int32(1)
	if v681 < v679 {
		v657 = v681
		v660 = v679
		goto L262
	} else {
		goto L269
	}
L267:
	;
	if v674 == int32(0) {
		goto L258
	} else {
		goto L268
	}
L268:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v649)+4))
	v679 = v678
	goto L266
L269:
	;
	goto L263
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v698
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	v702 = F_TypeNameListToString(m, v701)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L5
	} else {
		goto L271
	}
L271:
	;
	v1501 = v702
	goto L117
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v719
	v1501 = int32(0)
	goto L117
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_27)
	v717 = F_TypeNameToString(m, v672)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L5
	} else {
		goto L278
	}
L274:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v705)+8))
	if v707 == int32(0) {
		goto L273
	} else {
		goto L275
	}
L275:
	;
	v710 = F_LookupNamespaceNoError(m, v707)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L5
	} else {
		goto L276
	}
L276:
	;
	if v710 != 0 {
		goto L273
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_5)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v705)+8))
	v719 = v714
	goto L272
L278:
	;
	v719 = v717
	goto L272
L279:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	if v732 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L280:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v723)+8))
	if v725 == int32(0) {
		goto L279
	} else {
		goto L281
	}
L281:
	;
	v728 = F_LookupNamespaceNoError(m, v725)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L5
	} else {
		goto L282
	}
L282:
	;
	if v728 != 0 {
		goto L279
	} else {
		goto L283
	}
L283:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v723)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v730
	goto L122
L284:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v755)+4))
	v788 = F_makeRangeVarFromNameList(m, v787)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L5
	} else {
		goto L300
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_38)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v781 = F_NameListToString(m, v780)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L5
	} else {
		goto L296
	}
L286:
	;
	v735 = int32(0)
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v732)+4))
	if v736 <= v735 {
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v740 = v735
	v743 = v736
	goto L288
L288:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v732)+12))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v751+v740<<(uint(int32(2))%32))))
	if v755 != 0 {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	goto L285
L290:
	;
	v757 = F_LookupTypeNameOid(m, v755, int32(1))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L5
	} else {
		goto L293
	}
L291:
	;
	v762 = v743
	goto L292
L292:
	;
	v764 = v740 + int32(1)
	if v764 < v762 {
		v740 = v764
		v743 = v762
		goto L288
	} else {
		goto L295
	}
L293:
	;
	if v757 == int32(0) {
		goto L284
	} else {
		goto L294
	}
L294:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v732)+4))
	v762 = v761
	goto L292
L295:
	;
	goto L289
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v781
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	v785 = F_TypeNameListToString(m, v784)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L5
	} else {
		goto L297
	}
L297:
	;
	v1501 = v785
	goto L117
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v802
	v1501 = int32(0)
	goto L117
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_27)
	v800 = F_TypeNameToString(m, v755)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L5
	} else {
		goto L304
	}
L300:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v788)+8))
	if v790 == int32(0) {
		goto L299
	} else {
		goto L301
	}
L301:
	;
	v793 = F_LookupNamespaceNoError(m, v790)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L5
	} else {
		goto L302
	}
L302:
	;
	if v793 != 0 {
		goto L299
	} else {
		goto L303
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_5)
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v788)+8))
	v802 = v797
	goto L298
L304:
	;
	v802 = v800
	goto L298
L305:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	if v815 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L306:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v806)+8))
	if v808 == int32(0) {
		goto L305
	} else {
		goto L307
	}
L307:
	;
	v811 = F_LookupNamespaceNoError(m, v808)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L5
	} else {
		goto L308
	}
L308:
	;
	if v811 != 0 {
		goto L305
	} else {
		goto L309
	}
L309:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v806)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v813
	goto L122
L310:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v838)+4))
	v866 = F_makeRangeVarFromNameList(m, v865)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L5
	} else {
		goto L325
	}
L311:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v863 = F_NameListToString(m, v862)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L5
	} else {
		goto L322
	}
L312:
	;
	v818 = int32(0)
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v815)+4))
	if v819 <= v818 {
		goto L311
	} else {
		goto L313
	}
L313:
	;
	v823 = v818
	v826 = v819
	goto L314
L314:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v815)+12))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v834+v823<<(uint(int32(2))%32))))
	if v838 != 0 {
		goto L316
	} else {
		goto L317
	}
L315:
	;
	goto L311
L316:
	;
	v840 = F_LookupTypeNameOid(m, v838, int32(1))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L5
	} else {
		goto L319
	}
L317:
	;
	v845 = v826
	goto L318
L318:
	;
	v847 = v823 + int32(1)
	if v847 < v845 {
		v823 = v847
		v826 = v845
		goto L314
	} else {
		goto L321
	}
L319:
	;
	if v840 == int32(0) {
		goto L310
	} else {
		goto L320
	}
L320:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v815)+4))
	v845 = v844
	goto L318
L321:
	;
	goto L315
L322:
	;
	v1327 = int32(_a_F_ExecDropStmt_39)
	v1328 = v863
	goto L122
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v880
	v1501 = int32(0)
	goto L117
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_27)
	v878 = F_TypeNameToString(m, v838)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L5
	} else {
		goto L329
	}
L325:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v866)+8))
	if v868 == int32(0) {
		goto L324
	} else {
		goto L326
	}
L326:
	;
	v871 = F_LookupNamespaceNoError(m, v868)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L5
	} else {
		goto L327
	}
L327:
	;
	if v871 != 0 {
		goto L324
	} else {
		goto L328
	}
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_5)
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v866)+8))
	v880 = v875
	goto L323
L329:
	;
	v880 = v878
	goto L323
L330:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v939)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+52)) = v940
	*(*int32)(unsafe.Add(mBase, uint32(v321)+164)) = v940
	v946 = F_list_make1_impl(m, int32(1), v321+int32(52))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L5
	} else {
		goto L343
	}
L331:
	;
	if v892 == int32(0) {
		goto L330
	} else {
		goto L332
	}
L332:
	;
	v896 = int32(0)
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v892)+4))
	if v897 <= v896 {
		goto L330
	} else {
		goto L333
	}
L333:
	;
	v901 = v896
	v904 = v897
	goto L334
L334:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v892)+12))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v912+v901<<(uint(int32(2))%32))))
	if v916 != 0 {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	goto L330
L336:
	;
	v918 = F_LookupTypeNameOid(m, v916, int32(1))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L5
	} else {
		goto L339
	}
L337:
	;
	v923 = v904
	goto L338
L338:
	;
	v925 = v901 + int32(1)
	if v925 < v923 {
		v901 = v925
		v904 = v923
		goto L334
	} else {
		goto L341
	}
L339:
	;
	if v918 == int32(0) {
		goto L118
	} else {
		goto L340
	}
L340:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v892)+4))
	v923 = v922
	goto L338
L341:
	;
	goto L335
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_40)
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1013)))
	v1015 = F_TypeNameToString(m, v1014)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L5
	} else {
		goto L362
	}
L343:
	;
	if v946 == int32(0) {
		goto L342
	} else {
		goto L344
	}
L344:
	;
	v950 = int32(0)
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v946)+4))
	if v951 <= v950 {
		goto L342
	} else {
		goto L345
	}
L345:
	;
	v955 = v950
	v958 = v951
	goto L346
L346:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v946)+12))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v966+v955<<(uint(int32(2))%32))))
	if v970 == int32(0) {
		v995 = v958
		goto L348
	} else {
		goto L349
	}
L347:
	;
	goto L342
L348:
	;
	v997 = v955 + int32(1)
	if v997 < v995 {
		v955 = v997
		v958 = v995
		goto L346
	} else {
		goto L361
	}
L349:
	;
	v974 = F_LookupTypeNameOid(m, v970, int32(1))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L5
	} else {
		goto L350
	}
L350:
	;
	if v974 != 0 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v946)+4))
	v995 = v976
	goto L348
L352:
	;
	goto L353
L353:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v970)+4))
	v978 = F_makeRangeVarFromNameList(m, v977)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L5
	} else {
		goto L356
	}
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v992
	v1501 = int32(0)
	goto L117
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_27)
	v990 = F_TypeNameToString(m, v970)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L5
	} else {
		goto L360
	}
L356:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v978)+8))
	if v980 == int32(0) {
		goto L355
	} else {
		goto L357
	}
L357:
	;
	v983 = F_LookupNamespaceNoError(m, v980)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L5
	} else {
		goto L358
	}
L358:
	;
	if v983 != 0 {
		goto L355
	} else {
		goto L359
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_5)
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v978)+8))
	v992 = v987
	goto L354
L360:
	;
	v992 = v990
	goto L354
L361:
	;
	goto L347
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v1015
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v1018)+4))
	v1020 = F_TypeNameToString(m, v1019)
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L5
	} else {
		goto L363
	}
L363:
	;
	v1501 = v1020
	goto L117
L364:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L5
	} else {
		goto L422
	}
L365:
	;
	v1024 = F_get_func_prokind(m, v365)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L5
	} else {
		goto L368
	}
L366:
	;
	goto L367
L367:
	;
	v1029 = m.G0
	v1031 = v1029 - int32(32)
	m.G0 = v1031
	v1034 = v321 + int32(152)
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1034)))
	v1037 = *(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[0]))
	if v1037 != 0 {
		goto L374
	} else {
		goto L375
	}
L368:
	;
	if v1024 == int32(97) {
		goto L364
	} else {
		goto L369
	}
L369:
	;
	goto L367
L370:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[1]))
	if v1107 != 0 {
		goto L406
	} else {
		goto L407
	}
L371:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L5
	} else {
		goto L402
	}
L372:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L5
	} else {
		goto L399
	}
L373:
	;
	v1092 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1081)+24)))
	if v1092 == int32(0) {
		goto L392
	} else {
		goto L393
	}
L374:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+4))
	if v1038 == v1035 {
		v1081 = v1037
		goto L373
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	v1048 = int32(0)
	goto L382
L377:
	;
	goto L376
L378:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[0])) = v1078
	v1081 = v1078
	goto L373
L379:
	;
	v1078 = v1054 + int32(_a_F_ExecDropStmt_41)
	goto L378
L380:
	;
	v1078 = v1054 + int32(_a_F_ExecDropStmt_42)
	goto L378
L381:
	;
	v1078 = v1054 + int32(_a_F_ExecDropStmt_43)
	goto L378
L382:
	;
	v1054 = v1048 * int32(40)
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+uint32(_c_F_ExecDropStmt[2])))
	if v1035 != v1055 {
		goto L384
	} else {
		goto L385
	}
L383:
	;
	v1078 = v1054 + int32(_a_F_ExecDropStmt_44)
	goto L378
L384:
	;
	if v1048 == int32(36) {
		goto L372
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	goto L383
L387:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+uint32(_c_F_ExecDropStmt[3])))
	if v1061 == v1035 {
		goto L379
	} else {
		goto L388
	}
L388:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+uint32(_c_F_ExecDropStmt[4])))
	if v1063 == v1035 {
		goto L380
	} else {
		goto L389
	}
L389:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+uint32(_c_F_ExecDropStmt[5])))
	if v1065 == v1035 {
		goto L381
	} else {
		goto L390
	}
L390:
	;
	v1048 = v1048 + int32(4)
	goto L382
L391:
	;
	m.G0 = v1031 + int32(32)
	goto L370
L392:
	;
	v1107 = int32(0)
	goto L391
L393:
	;
	goto L394
L394:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+12))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+4))
	v1098 = F_SearchSysCache1(m, v1096, v1097)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L5
	} else {
		goto L395
	}
L395:
	;
	if v1098 == int32(0) {
		goto L371
	} else {
		goto L396
	}
L396:
	;
	v1102 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1081)+24)))
	v1103 = F_SysCacheGetAttrNotNull(m, v1096, v1098, v1102)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L5
	} else {
		goto L397
	}
L397:
	;
	F_ReleaseCatCache(m, v1098)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L5
	} else {
		goto L398
	}
L398:
	;
	v1107 = v1103
	goto L391
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1031)+16)) = v1035
	F_errmsg_internal(m, int32(_a_F_ExecDropStmt_45), v1031+int32(16))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L5
	} else {
		goto L400
	}
L400:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_46), int32(2777), int32(_a_F_ExecDropStmt_47))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L5
	} else {
		goto L401
	}
L401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L402:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1031)+4)) = v1132
	*(*int32)(unsafe.Add(mBase, uint32(v1031))) = v1096
	F_errmsg_internal(m, int32(_a_F_ExecDropStmt_48), v1031)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L5
	} else {
		goto L403
	}
L403:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_46), int32(2593), int32(_a_F_ExecDropStmt_49))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L5
	} else {
		goto L404
	}
L404:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L405:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v321)+148))
	if v1188 != 0 {
		goto L417
	} else {
		goto L418
	}
L406:
	;
	v1146 = F_object_ownercheck(m, int32(2615), v1107, v1144)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L5
	} else {
		goto L409
	}
L407:
	;
	goto L408
L408:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1178 = *(*int64)(unsafe.Add(mBase, uint32(v321)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v321)+104)) = v1178
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v321)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+112)) = v1180
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v321)+148))
	F_check_object_ownership(m, v1144, v1177, v321+int32(104), v347, v1184)
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L5
	} else {
		goto L416
	}
L409:
	;
	if v1146 == int32(0) {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[1]))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1153 = *(*int64)(unsafe.Add(mBase, uint32(v321)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v321)+120)) = v1153
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v321)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+128)) = v1155
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v321)+148))
	F_check_object_ownership(m, v1151, v1152, v321+int32(120), v347, v1159)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L5
	} else {
		goto L413
	}
L411:
	;
	goto L412
L412:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[6]))
	goto L414
L413:
	;
	goto L412
L414:
	;
	if base.B2i32(v1164 != int32(0))&base.B2i32(v1107 == v1164) == int32(0) {
		goto L405
	} else {
		goto L415
	}
L415:
	;
	v1171 = int32(_a_F_ExecDropStmt_50)
	v1173 = *(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecDropStmt[7])) = v1173 | int32(1)
	goto L405
L416:
	;
	goto L405
L417:
	;
	F_relation_close(m, v1188, int32(0))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L5
	} else {
		goto L420
	}
L418:
	;
	goto L419
L419:
	;
	F_add_exact_object_address(m, v321+int32(152), v323)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L5
	} else {
		goto L421
	}
L420:
	;
	goto L419
L421:
	;
	goto L114
L422:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L5
	} else {
		goto L423
	}
L423:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v1204 = F_NameListToString(m, v1203)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L5
	} else {
		goto L424
	}
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+96)) = v1204
	F_errmsg(m, int32(_a_F_ExecDropStmt_51), v321+int32(96))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L5
	} else {
		goto L425
	}
L425:
	;
	F_errhint(m, int32(_a_F_ExecDropStmt_52), int32(0))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L5
	} else {
		goto L426
	}
L426:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_18), int32(98), int32(_a_F_ExecDropStmt_53))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L5
	} else {
		goto L427
	}
L427:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+80)) = v363
	F_errmsg_internal(m, int32(_a_F_ExecDropStmt_54), v321+int32(80))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L5
	} else {
		goto L429
	}
L429:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_18), int32(512), int32(_a_F_ExecDropStmt_19))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L5
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
	*(*int32)(unsafe.Add(mBase, uint32(v321)+64)) = v363
	F_errmsg_internal(m, int32(_a_F_ExecDropStmt_54), v321-int32(-64))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L5
	} else {
		goto L432
	}
L432:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_18), int32(496), int32(_a_F_ExecDropStmt_19))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L5
	} else {
		goto L433
	}
L433:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_55)
	v1269 = F_NameListToString(m, v1256)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L5
	} else {
		goto L440
	}
L435:
	;
	v1258 = F_makeRangeVarFromNameList(m, v1256)
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L5
	} else {
		goto L436
	}
L436:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+8))
	if v1260 == int32(0) {
		goto L434
	} else {
		goto L437
	}
L437:
	;
	v1263 = F_LookupNamespaceNoError(m, v1260)
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L5
	} else {
		goto L438
	}
L438:
	;
	if v1263 != 0 {
		goto L434
	} else {
		goto L439
	}
L439:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v1265
	goto L122
L440:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v1269
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1272)))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+4))
	v1501 = v1274
	goto L117
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_56)
	v1289 = F_NameListToString(m, v1276)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L5
	} else {
		goto L447
	}
L442:
	;
	v1278 = F_makeRangeVarFromNameList(m, v1276)
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L5
	} else {
		goto L443
	}
L443:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+8))
	if v1280 == int32(0) {
		goto L441
	} else {
		goto L444
	}
L444:
	;
	v1283 = F_LookupNamespaceNoError(m, v1280)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L5
	} else {
		goto L445
	}
L445:
	;
	if v1283 != 0 {
		goto L441
	} else {
		goto L446
	}
L446:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+8))
	v1327 = int32(_a_F_ExecDropStmt_5)
	v1328 = v1285
	goto L122
L447:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v1289
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1292)))
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1293)+4))
	v1501 = v1294
	goto L117
L448:
	;
	if v1304 != 0 {
		v1501 = int32(0)
		goto L117
	} else {
		goto L449
	}
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_57)
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1308+v1309<<(uint(int32(2))%32)-int32(4))))
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v1316
	v1320 = F_list_copy_head(m, v347, v1309-int32(1))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L5
	} else {
		goto L450
	}
L450:
	;
	v1322 = F_NameListToString(m, v1320)
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L5
	} else {
		goto L451
	}
L451:
	;
	v1501 = v1322
	goto L117
L452:
	;
	if v1344 != 0 {
		v1501 = int32(0)
		goto L117
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_58)
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1348+v1349<<(uint(int32(2))%32)-int32(4))))
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v1356
	v1360 = F_list_copy_head(m, v347, v1349-int32(1))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L5
	} else {
		goto L454
	}
L454:
	;
	v1362 = F_NameListToString(m, v1360)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L5
	} else {
		goto L455
	}
L455:
	;
	v1501 = v1362
	goto L117
L456:
	;
	if v1369 != 0 {
		v1501 = int32(0)
		goto L117
	} else {
		goto L457
	}
L457:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_59)
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1373+v1374<<(uint(int32(2))%32)-int32(4))))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v1381
	v1385 = F_list_copy_head(m, v347, v1374-int32(1))
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L5
	} else {
		goto L458
	}
L458:
	;
	v1387 = F_NameListToString(m, v1385)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L5
	} else {
		goto L459
	}
L459:
	;
	v1501 = v1387
	goto L117
L460:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+4))
	v1454 = F_makeRangeVarFromNameList(m, v1453)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L5
	} else {
		goto L476
	}
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_60)
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1445)))
	v1447 = F_TypeNameToString(m, v1446)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L5
	} else {
		goto L473
	}
L462:
	;
	if v1396 == int32(0) {
		goto L461
	} else {
		goto L463
	}
L463:
	;
	v1400 = int32(0)
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+4))
	if v1401 <= v1400 {
		goto L461
	} else {
		goto L464
	}
L464:
	;
	v1405 = v1400
	v1408 = v1401
	goto L465
L465:
	;
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+12))
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1416+v1405<<(uint(int32(2))%32))))
	if v1420 != 0 {
		goto L467
	} else {
		goto L468
	}
L466:
	;
	goto L461
L467:
	;
	v1422 = F_LookupTypeNameOid(m, v1420, int32(1))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L5
	} else {
		goto L470
	}
L468:
	;
	v1427 = v1408
	goto L469
L469:
	;
	v1429 = v1405 + int32(1)
	if v1429 < v1427 {
		v1405 = v1429
		v1408 = v1427
		goto L465
	} else {
		goto L472
	}
L470:
	;
	if v1422 == int32(0) {
		goto L460
	} else {
		goto L471
	}
L471:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+4))
	v1427 = v1426
	goto L469
L472:
	;
	goto L466
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v1447
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+4))
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+4))
	v1501 = v1452
	goto L117
L474:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v1468
	v1501 = int32(0)
	goto L117
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_27)
	v1466 = F_TypeNameToString(m, v1420)
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L5
	} else {
		goto L480
	}
L476:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+8))
	if v1456 == int32(0) {
		goto L475
	} else {
		goto L477
	}
L477:
	;
	v1459 = F_LookupNamespaceNoError(m, v1456)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L5
	} else {
		goto L478
	}
L478:
	;
	if v1459 != 0 {
		goto L475
	} else {
		goto L479
	}
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_5)
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+8))
	v1468 = v1463
	goto L474
L480:
	;
	v1468 = v1466
	goto L474
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+172)) = v1486
	v1501 = int32(0)
	goto L117
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_27)
	v1484 = F_TypeNameToString(m, v916)
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L5
	} else {
		goto L487
	}
L483:
	;
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+8))
	if v1474 == int32(0) {
		goto L482
	} else {
		goto L484
	}
L484:
	;
	v1477 = F_LookupNamespaceNoError(m, v1474)
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L5
	} else {
		goto L485
	}
L485:
	;
	if v1477 != 0 {
		goto L482
	} else {
		goto L486
	}
L486:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+136)) = int32(_a_F_ExecDropStmt_5)
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+8))
	v1486 = v1481
	goto L481
L487:
	;
	v1486 = v1484
	goto L481
L488:
	;
	if v1501 == int32(0) {
		v1522 = v1502
		goto L116
	} else {
		goto L489
	}
L489:
	;
	v1509 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L5
	} else {
		goto L490
	}
L490:
	;
	if v1509 == int32(0) {
		goto L114
	} else {
		goto L491
	}
L491:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v321)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+32)) = v1513
	*(*int32)(unsafe.Add(mBase, uint32(v321)+36)) = v1501
	F_errmsg(m, v1502, v321+int32(32))
	mBase = m.M
	v1519 = m.ExcPending
	if v1519 != 0 {
		goto L5
	} else {
		goto L492
	}
L492:
	;
	v1558 = int32(523)
	goto L115
L493:
	;
	if v1535 == int32(0) {
		goto L114
	} else {
		goto L494
	}
L494:
	;
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v321)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+16)) = v1539
	F_errmsg(m, v1522, v321+int32(16))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L5
	} else {
		goto L495
	}
L495:
	;
	v1558 = int32(521)
	goto L115
L496:
	;
	goto L114
L497:
	;
	goto L112
L498:
	;
	F_free_object_addresses(m, v323)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L5
	} else {
		goto L499
	}
L499:
	;
	m.G0 = v321 + int32(176)
	goto L106
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v363
	F_errmsg_internal(m, int32(_a_F_ExecDropStmt_61), v321)
	mBase = m.M
	v1618 = m.ExcPending
	if v1618 != 0 {
		goto L5
	} else {
		goto L501
	}
L501:
	;
	F_errfinish(m, int32(_a_F_ExecDropStmt_18), int32(518), int32(_a_F_ExecDropStmt_19))
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L5
	} else {
		goto L502
	}
L502:
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
