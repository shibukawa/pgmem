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
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int64
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int64
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v451 int32
	_ = v451
	var v453 int64
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int64
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int64
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v559 int64
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v648 int64
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v737 int64
	_ = v737
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v809 int64
	_ = v809
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v900 int32
	_ = v900
	var v904 int64
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v952 int64
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v967 int32
	_ = v967
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
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
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1000 int64
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1049 int64
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
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
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1097 int64
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1165 int64
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1308 int32
	_ = v1308
	var v1310 int64
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1325 int32
	_ = v1325
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1370 int32
	_ = v1370
	var v1375 int32
	_ = v1375
	var v1383 int32
	_ = v1383
	var v1385 int64
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1400 int32
	_ = v1400
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1422 int64
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1437 int32
	_ = v1437
	var v1442 int32
	_ = v1442
	var v1445 int64
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
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
	var v1506 int32
	_ = v1506
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1531 int64
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1546 int32
	_ = v1546
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1597 int32
	_ = v1597
	var v1599 int64
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1614 int32
	_ = v1614
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1672 int32
	_ = v1672
	var v1677 int32
	_ = v1677
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1687 int64
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1710 int32
	_ = v1710
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1738 int32
	_ = v1738
	var v1746 int32
	_ = v1746
	var v1755 int32
	_ = v1755
	var v1763 int32
	_ = v1763
	var v1772 int32
	_ = v1772
	var v1780 int32
	_ = v1780
	var v1789 int32
	_ = v1789
	var v1797 int32
	_ = v1797
	var v1805 int32
	_ = v1805
	var v1813 int32
	_ = v1813
	var v1822 int32
	_ = v1822
	var v1830 int32
	_ = v1830
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1853 int32
	_ = v1853
	v9 = m.G0
	v11 = v9 - int32(1440)
	m.G0 = v11
	F_initStringInfo(m, v11+int32(1424))
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
					v1250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v1251 = F_get_tablespace_name(m, v1250)
					mBase = m.M
					v1252 = m.ExcPending
					if v1252 != 0 {
						return int32(0)
					} else {
						if v1251 == int32(0) {
							if l1 != 0 {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v1258 = m.ExcPending
								if v1258 != 0 {
									return int32(0)
								} else {
									v1259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+800)) = v1259
									F_errmsg_internal(m, int32(_a_F_getObjectDescription_0), v11+int32(800))
									mBase = m.M
									v1265 = m.ExcPending
									if v1265 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3749), int32(_a_F_getObjectDescription_2))
										mBase = m.M
										v1270 = m.ExcPending
										if v1270 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v11)+816)) = v1251
							F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_3), v11+int32(816))
							mBase = m.M
							v1278 = m.ExcPending
							if v1278 != 0 {
								return int32(0)
							} else {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							}
						}
					}
				case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 35, 36, 37, 38, 39, 40, 41, 43, 44, 45:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v1646 = m.ExcPending
					if v1646 != 0 {
						return int32(0)
					} else {
						v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1647
						F_errmsg_internal(m, int32(_a_F_getObjectDescription_4), v11)
						mBase = m.M
						v1651 = m.ExcPending
						if v1651 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(_a_F_getObjectDescription_5), int32(_a_F_getObjectDescription_2))
							mBase = m.M
							v1656 = m.ExcPending
							if v1656 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 34:
					v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v137 = F_format_type_extended(m, v134, int32(-1), int32(8))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return int32(0)
					} else {
						if v137 == int32(0) {
							v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
							v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
							m.G0 = v11 + int32(1440)
							if v1847 != 0 {
								v1853 = v1848
							} else {
								v1853 = int32(0)
							}
							return v1853
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v137
							F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_6), v11+int32(48))
							mBase = m.M
							v148 = m.ExcPending
							if v148 != 0 {
								return int32(0)
							} else {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							}
						}
					}
				case 42:
					v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v122 = F_format_procedure_extended(m, v120, int32(1))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int32(0)
					} else {
						if v122 == int32(0) {
							v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
							v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
							m.G0 = v11 + int32(1440)
							if v1847 != 0 {
								v1853 = v1848
							} else {
								v1853 = int32(0)
							}
							return v1853
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v122
							F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_7), v11+int32(32))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return int32(0)
							} else {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							}
						}
					}
				case 46:
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v86 == int32(0) {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_getRelationDescription(m, v11+int32(1424), v91, l1)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
							v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
							m.G0 = v11 + int32(1440)
							if v1847 != 0 {
								v1853 = v1848
							} else {
								v1853 = int32(0)
							}
							return v1853
						}
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v96 = F_get_attname(m, v94, base.I32_extend16_s(v86), l1)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							if v96 == int32(0) {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							} else {
								v101 = v11 + int32(1360)
								F_initStringInfo(m, v101)
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									F_getRelationDescription(m, v101, v104, l1)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v96
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v108
										F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_8), v11+int32(16))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											v117 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
											F_pfree(m, v117)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
												v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
												m.G0 = v11 + int32(1440)
												if v1847 != 0 {
													v1853 = v1848
												} else {
													v1853 = int32(0)
												}
												return v1853
											}
										}
									}
								}
							}
						}
					}
				case 47:
					v1143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v1144 = F_GetUserNameFromId(m, v1143, l1)
					mBase = m.M
					v1145 = m.ExcPending
					if v1145 != 0 {
						return int32(0)
					} else {
						if v1144 == int32(0) {
							v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
							v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
							m.G0 = v11 + int32(1440)
							if v1847 != 0 {
								v1853 = v1848
							} else {
								v1853 = int32(0)
							}
							return v1853
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+720)) = v1144
							F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_9), v11+int32(720))
							mBase = m.M
							v1155 = m.ExcPending
							if v1155 != 0 {
								return int32(0)
							} else {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							}
						}
					}
				case 48:
					v1158 = F_table_open(m, int32(1261), int32(1))
					mBase = m.M
					v1159 = m.ExcPending
					if v1159 != 0 {
						return int32(0)
					} else {
						v1161 = v11 + int32(1360)
						v1165 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						F_ScanKeyInit(m, v1161, int32(1), int32(3), int32(184), v1165)
						mBase = m.M
						v1167 = m.ExcPending
						if v1167 != 0 {
							return int32(0)
						} else {
							v1169 = int32(1)
							v1172 = F_systable_beginscan(m, v1158, int32(_a_F_getObjectDescription_10), v1169, int32(0), v1169, v1161)
							mBase = m.M
							v1173 = m.ExcPending
							if v1173 != 0 {
								return int32(0)
							} else {
								v1174 = F_systable_getnext(m, v1172)
								mBase = m.M
								v1175 = m.ExcPending
								if v1175 != 0 {
									return int32(0)
								} else {
									if v1174 == int32(0) {
										if l1 != 0 {
											F_systable_endscan(m, v1172)
											mBase = m.M
											v1217 = m.ExcPending
											if v1217 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v1158, int32(1))
												mBase = m.M
												v1220 = m.ExcPending
												if v1220 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v1181 = m.ExcPending
											if v1181 != 0 {
												return int32(0)
											} else {
												v1182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+736)) = v1182
												F_errmsg_internal(m, int32(_a_F_getObjectDescription_11), v11+int32(736))
												mBase = m.M
												v1188 = m.ExcPending
												if v1188 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3706), int32(_a_F_getObjectDescription_2))
													mBase = m.M
													v1193 = m.ExcPending
													if v1193 != 0 {
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
										v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+16))
										v1195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1194)+22)))
										v1196 = v1194 + v1195
										v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1196)+8))
										v1199 = F_GetUserNameFromId(m, v1197, int32(0))
										mBase = m.M
										v1200 = m.ExcPending
										if v1200 != 0 {
											return int32(0)
										} else {
											v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1196)+4))
											v1203 = F_GetUserNameFromId(m, v1201, int32(0))
											mBase = m.M
											v1204 = m.ExcPending
											if v1204 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+756)) = v1203
												*(*int32)(unsafe.Add(mBase, uint32(v11)+752)) = v1199
												F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_12), v11+int32(752))
												mBase = m.M
												v1213 = m.ExcPending
												if v1213 != 0 {
													return int32(0)
												} else {
													F_systable_endscan(m, v1172)
													mBase = m.M
													v1217 = m.ExcPending
													if v1217 != 0 {
														return int32(0)
													} else {
														F_relation_close(m, v1158, int32(1))
														mBase = m.M
														v1220 = m.ExcPending
														if v1220 != 0 {
															return int32(0)
														} else {
															v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
															v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
															m.G0 = v11 + int32(1440)
															if v1847 != 0 {
																v1853 = v1848
															} else {
																v1853 = int32(0)
															}
															return v1853
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
					v1221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v1222 = F_get_database_name(m, v1221)
					mBase = m.M
					v1223 = m.ExcPending
					if v1223 != 0 {
						return int32(0)
					} else {
						if v1222 == int32(0) {
							if l1 != 0 {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v1229 = m.ExcPending
								if v1229 != 0 {
									return int32(0)
								} else {
									v1230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+768)) = v1230
									F_errmsg_internal(m, int32(_a_F_getObjectDescription_13), v11+int32(768))
									mBase = m.M
									v1236 = m.ExcPending
									if v1236 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3733), int32(_a_F_getObjectDescription_2))
										mBase = m.M
										v1241 = m.ExcPending
										if v1241 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v11)+784)) = v1222
							F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_14), v11+int32(784))
							mBase = m.M
							v1249 = m.ExcPending
							if v1249 != 0 {
								return int32(0)
							} else {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							}
						}
					}
				default:
					switch v19 - int32(1417) {
					case 0:
						v1295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1296 = F_GetForeignServerExtended(m, v1295, l1)
						mBase = m.M
						v1297 = m.ExcPending
						if v1297 != 0 {
							return int32(0)
						} else {
							if v1296 == int32(0) {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							} else {
								v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+848)) = v1300
								F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_15), v11+int32(848))
								mBase = m.M
								v1308 = m.ExcPending
								if v1308 != 0 {
									return int32(0)
								} else {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								}
							}
						}
					case 1:
						v1310 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						v1311 = F_SearchSysCache1(m, int32(83), v1310)
						mBase = m.M
						v1312 = m.ExcPending
						if v1312 != 0 {
							return int32(0)
						} else {
							if v1311 == int32(0) {
								if l1 != 0 {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1318 = m.ExcPending
									if v1318 != 0 {
										return int32(0)
									} else {
										v1319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+864)) = v1319
										F_errmsg_internal(m, int32(_a_F_getObjectDescription_16), v11+int32(864))
										mBase = m.M
										v1325 = m.ExcPending
										if v1325 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3791), int32(_a_F_getObjectDescription_2))
											mBase = m.M
											v1330 = m.ExcPending
											if v1330 != 0 {
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
								v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+16))
								v1332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1331)+22)))
								v1333 = v1331 + v1332
								v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+4))
								v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+8))
								v1336 = F_GetForeignServer(m, v1335)
								mBase = m.M
								v1337 = m.ExcPending
								if v1337 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v1311)
									mBase = m.M
									v1339 = m.ExcPending
									if v1339 != 0 {
										return int32(0)
									} else {
										if v1334 != 0 {
											v1341 = F_GetUserNameFromId(m, v1334, int32(0))
											mBase = m.M
											v1342 = m.ExcPending
											if v1342 != 0 {
												return int32(0)
											} else {
												v1344 = v1341
												v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1336)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+884)) = v1345
												*(*int32)(unsafe.Add(mBase, uint32(v11)+880)) = v1344
												F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_17), v11+int32(880))
												mBase = m.M
												v1354 = m.ExcPending
												if v1354 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
												}
											}
										} else {
											v1344 = int32(_a_F_getObjectDescription_18)
											v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1336)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+884)) = v1345
											*(*int32)(unsafe.Add(mBase, uint32(v11)+880)) = v1344
											F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_17), v11+int32(880))
											mBase = m.M
											v1354 = m.ExcPending
											if v1354 != 0 {
												return int32(0)
											} else {
												v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
												v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
												m.G0 = v11 + int32(1440)
												if v1847 != 0 {
													v1853 = v1848
												} else {
													v1853 = int32(0)
												}
												return v1853
											}
										}
									}
								}
							}
						}
					default:
						if v19 == int32(826) {
							v1680 = F_table_open(m, int32(826), int32(1))
							mBase = m.M
							v1681 = m.ExcPending
							if v1681 != 0 {
								return int32(0)
							} else {
								v1683 = v11 + int32(1360)
								v1687 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
								F_ScanKeyInit(m, v1683, int32(1), int32(3), int32(184), v1687)
								mBase = m.M
								v1689 = m.ExcPending
								if v1689 != 0 {
									return int32(0)
								} else {
									v1691 = int32(1)
									v1694 = F_systable_beginscan(m, v1680, int32(828), v1691, int32(0), v1691, v1683)
									mBase = m.M
									v1695 = m.ExcPending
									if v1695 != 0 {
										return int32(0)
									} else {
										v1696 = F_systable_getnext(m, v1694)
										mBase = m.M
										v1697 = m.ExcPending
										if v1697 != 0 {
											return int32(0)
										} else {
											if v1696 == int32(0) {
												if l1 != 0 {
													F_systable_endscan(m, v1694)
													mBase = m.M
													v1836 = m.ExcPending
													if v1836 != 0 {
														return int32(0)
													} else {
														F_relation_close(m, v1680, int32(1))
														mBase = m.M
														v1839 = m.ExcPending
														if v1839 != 0 {
															return int32(0)
														} else {
															v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
															v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
															m.G0 = v11 + int32(1440)
															if v1847 != 0 {
																v1853 = v1848
															} else {
																v1853 = int32(0)
															}
															return v1853
														}
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v1703 = m.ExcPending
													if v1703 != 0 {
														return int32(0)
													} else {
														v1704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v11)+896)) = v1704
														F_errmsg_internal(m, int32(_a_F_getObjectDescription_19), v11+int32(896))
														mBase = m.M
														v1710 = m.ExcPending
														if v1710 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3837), int32(_a_F_getObjectDescription_2))
															mBase = m.M
															v1715 = m.ExcPending
															if v1715 != 0 {
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
												v1716 = *(*int32)(unsafe.Add(mBase, uint32(v1696)+16))
												v1717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1716)+22)))
												v1718 = v1716 + v1717
												v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1718)+4))
												v1721 = F_GetUserNameFromId(m, v1719, int32(0))
												mBase = m.M
												v1722 = m.ExcPending
												if v1722 != 0 {
													return int32(0)
												} else {
													v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1718)+8))
													if v1723 != 0 {
														v1724 = F_get_namespace_name(m, v1723)
														mBase = m.M
														v1725 = m.ExcPending
														if v1725 != 0 {
															return int32(0)
														} else {
															v1726 = v1724
															v1727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1718)+12)))
															switch v1727 - int32(76) {
															case 0:
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1088)) = v1721
																F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_20), v11+int32(1088))
																mBase = m.M
																v1813 = m.ExcPending
																if v1813 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1694)
																	mBase = m.M
																	v1836 = m.ExcPending
																	if v1836 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v1680, int32(1))
																		mBase = m.M
																		v1839 = m.ExcPending
																		if v1839 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
																		}
																	}
																}
															default:
																if v1726 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+932)) = v1726
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+928)) = v1721
																	F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_21), v11+int32(928))
																	mBase = m.M
																	v1822 = m.ExcPending
																	if v1822 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1694)
																		mBase = m.M
																		v1836 = m.ExcPending
																		if v1836 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v1680, int32(1))
																			mBase = m.M
																			v1839 = m.ExcPending
																			if v1839 != 0 {
																				return int32(0)
																			} else {
																				v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																				v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																				m.G0 = v11 + int32(1440)
																				if v1847 != 0 {
																					v1853 = v1848
																				} else {
																					v1853 = int32(0)
																				}
																				return v1853
																			}
																		}
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+912)) = v1721
																	F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_22), v11+int32(912))
																	mBase = m.M
																	v1830 = m.ExcPending
																	if v1830 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1694)
																		mBase = m.M
																		v1836 = m.ExcPending
																		if v1836 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v1680, int32(1))
																			mBase = m.M
																			v1839 = m.ExcPending
																			if v1839 != 0 {
																				return int32(0)
																			} else {
																				v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																				v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																				m.G0 = v11 + int32(1440)
																				if v1847 != 0 {
																					v1853 = v1848
																				} else {
																					v1853 = int32(0)
																				}
																				return v1853
																			}
																		}
																	}
																}
															case 7:
																if v1726 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+996)) = v1726
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+992)) = v1721
																	F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_23), v11+int32(992))
																	mBase = m.M
																	v1755 = m.ExcPending
																	if v1755 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1694)
																		mBase = m.M
																		v1836 = m.ExcPending
																		if v1836 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v1680, int32(1))
																			mBase = m.M
																			v1839 = m.ExcPending
																			if v1839 != 0 {
																				return int32(0)
																			} else {
																				v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																				v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																				m.G0 = v11 + int32(1440)
																				if v1847 != 0 {
																					v1853 = v1848
																				} else {
																					v1853 = int32(0)
																				}
																				return v1853
																			}
																		}
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+976)) = v1721
																	F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_24), v11+int32(976))
																	mBase = m.M
																	v1763 = m.ExcPending
																	if v1763 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1694)
																		mBase = m.M
																		v1836 = m.ExcPending
																		if v1836 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v1680, int32(1))
																			mBase = m.M
																			v1839 = m.ExcPending
																			if v1839 != 0 {
																				return int32(0)
																			} else {
																				v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																				v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																				m.G0 = v11 + int32(1440)
																				if v1847 != 0 {
																					v1853 = v1848
																				} else {
																					v1853 = int32(0)
																				}
																				return v1853
																			}
																		}
																	}
																}
															case 8:
																if v1726 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+1060)) = v1726
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+1056)) = v1721
																	F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_25), v11+int32(1056))
																	mBase = m.M
																	v1789 = m.ExcPending
																	if v1789 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1694)
																		mBase = m.M
																		v1836 = m.ExcPending
																		if v1836 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v1680, int32(1))
																			mBase = m.M
																			v1839 = m.ExcPending
																			if v1839 != 0 {
																				return int32(0)
																			} else {
																				v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																				v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																				m.G0 = v11 + int32(1440)
																				if v1847 != 0 {
																					v1853 = v1848
																				} else {
																					v1853 = int32(0)
																				}
																				return v1853
																			}
																		}
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+1040)) = v1721
																	F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_26), v11+int32(1040))
																	mBase = m.M
																	v1797 = m.ExcPending
																	if v1797 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1694)
																		mBase = m.M
																		v1836 = m.ExcPending
																		if v1836 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v1680, int32(1))
																			mBase = m.M
																			v1839 = m.ExcPending
																			if v1839 != 0 {
																				return int32(0)
																			} else {
																				v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																				v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																				m.G0 = v11 + int32(1440)
																				if v1847 != 0 {
																					v1853 = v1848
																				} else {
																					v1853 = int32(0)
																				}
																				return v1853
																			}
																		}
																	}
																}
															case 26:
																if v1726 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+1028)) = v1726
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+1024)) = v1721
																	F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_27), v11+int32(1024))
																	mBase = m.M
																	v1772 = m.ExcPending
																	if v1772 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1694)
																		mBase = m.M
																		v1836 = m.ExcPending
																		if v1836 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v1680, int32(1))
																			mBase = m.M
																			v1839 = m.ExcPending
																			if v1839 != 0 {
																				return int32(0)
																			} else {
																				v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																				v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																				m.G0 = v11 + int32(1440)
																				if v1847 != 0 {
																					v1853 = v1848
																				} else {
																					v1853 = int32(0)
																				}
																				return v1853
																			}
																		}
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+1008)) = v1721
																	F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_28), v11+int32(1008))
																	mBase = m.M
																	v1780 = m.ExcPending
																	if v1780 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1694)
																		mBase = m.M
																		v1836 = m.ExcPending
																		if v1836 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v1680, int32(1))
																			mBase = m.M
																			v1839 = m.ExcPending
																			if v1839 != 0 {
																				return int32(0)
																			} else {
																				v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																				v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																				m.G0 = v11 + int32(1440)
																				if v1847 != 0 {
																					v1853 = v1848
																				} else {
																					v1853 = int32(0)
																				}
																				return v1853
																			}
																		}
																	}
																}
															case 34:
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1072)) = v1721
																F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_29), v11+int32(1072))
																mBase = m.M
																v1805 = m.ExcPending
																if v1805 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1694)
																	mBase = m.M
																	v1836 = m.ExcPending
																	if v1836 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v1680, int32(1))
																		mBase = m.M
																		v1839 = m.ExcPending
																		if v1839 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
																		}
																	}
																}
															case 38:
																if v1726 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+964)) = v1726
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+960)) = v1721
																	F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_30), v11+int32(960))
																	mBase = m.M
																	v1738 = m.ExcPending
																	if v1738 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1694)
																		mBase = m.M
																		v1836 = m.ExcPending
																		if v1836 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v1680, int32(1))
																			mBase = m.M
																			v1839 = m.ExcPending
																			if v1839 != 0 {
																				return int32(0)
																			} else {
																				v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																				v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																				m.G0 = v11 + int32(1440)
																				if v1847 != 0 {
																					v1853 = v1848
																				} else {
																					v1853 = int32(0)
																				}
																				return v1853
																			}
																		}
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+944)) = v1721
																	F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_31), v11+int32(944))
																	mBase = m.M
																	v1746 = m.ExcPending
																	if v1746 != 0 {
																		return int32(0)
																	} else {
																		F_systable_endscan(m, v1694)
																		mBase = m.M
																		v1836 = m.ExcPending
																		if v1836 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v1680, int32(1))
																			mBase = m.M
																			v1839 = m.ExcPending
																			if v1839 != 0 {
																				return int32(0)
																			} else {
																				v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																				v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																				m.G0 = v11 + int32(1440)
																				if v1847 != 0 {
																					v1853 = v1848
																				} else {
																					v1853 = int32(0)
																				}
																				return v1853
																			}
																		}
																	}
																}
															}
														}
													} else {
														v1726 = int32(0)
														v1727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1718)+12)))
														switch v1727 - int32(76) {
														case 0:
															*(*int32)(unsafe.Add(mBase, uint32(v11)+1088)) = v1721
															F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_20), v11+int32(1088))
															mBase = m.M
															v1813 = m.ExcPending
															if v1813 != 0 {
																return int32(0)
															} else {
																F_systable_endscan(m, v1694)
																mBase = m.M
																v1836 = m.ExcPending
																if v1836 != 0 {
																	return int32(0)
																} else {
																	F_relation_close(m, v1680, int32(1))
																	mBase = m.M
																	v1839 = m.ExcPending
																	if v1839 != 0 {
																		return int32(0)
																	} else {
																		v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																		v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																		m.G0 = v11 + int32(1440)
																		if v1847 != 0 {
																			v1853 = v1848
																		} else {
																			v1853 = int32(0)
																		}
																		return v1853
																	}
																}
															}
														default:
															if v1726 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+932)) = v1726
																*(*int32)(unsafe.Add(mBase, uint32(v11)+928)) = v1721
																F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_21), v11+int32(928))
																mBase = m.M
																v1822 = m.ExcPending
																if v1822 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1694)
																	mBase = m.M
																	v1836 = m.ExcPending
																	if v1836 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v1680, int32(1))
																		mBase = m.M
																		v1839 = m.ExcPending
																		if v1839 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
																		}
																	}
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+912)) = v1721
																F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_22), v11+int32(912))
																mBase = m.M
																v1830 = m.ExcPending
																if v1830 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1694)
																	mBase = m.M
																	v1836 = m.ExcPending
																	if v1836 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v1680, int32(1))
																		mBase = m.M
																		v1839 = m.ExcPending
																		if v1839 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
																		}
																	}
																}
															}
														case 7:
															if v1726 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+996)) = v1726
																*(*int32)(unsafe.Add(mBase, uint32(v11)+992)) = v1721
																F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_23), v11+int32(992))
																mBase = m.M
																v1755 = m.ExcPending
																if v1755 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1694)
																	mBase = m.M
																	v1836 = m.ExcPending
																	if v1836 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v1680, int32(1))
																		mBase = m.M
																		v1839 = m.ExcPending
																		if v1839 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
																		}
																	}
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+976)) = v1721
																F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_24), v11+int32(976))
																mBase = m.M
																v1763 = m.ExcPending
																if v1763 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1694)
																	mBase = m.M
																	v1836 = m.ExcPending
																	if v1836 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v1680, int32(1))
																		mBase = m.M
																		v1839 = m.ExcPending
																		if v1839 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
																		}
																	}
																}
															}
														case 8:
															if v1726 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1060)) = v1726
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1056)) = v1721
																F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_25), v11+int32(1056))
																mBase = m.M
																v1789 = m.ExcPending
																if v1789 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1694)
																	mBase = m.M
																	v1836 = m.ExcPending
																	if v1836 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v1680, int32(1))
																		mBase = m.M
																		v1839 = m.ExcPending
																		if v1839 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
																		}
																	}
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1040)) = v1721
																F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_26), v11+int32(1040))
																mBase = m.M
																v1797 = m.ExcPending
																if v1797 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1694)
																	mBase = m.M
																	v1836 = m.ExcPending
																	if v1836 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v1680, int32(1))
																		mBase = m.M
																		v1839 = m.ExcPending
																		if v1839 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
																		}
																	}
																}
															}
														case 26:
															if v1726 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1028)) = v1726
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1024)) = v1721
																F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_27), v11+int32(1024))
																mBase = m.M
																v1772 = m.ExcPending
																if v1772 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1694)
																	mBase = m.M
																	v1836 = m.ExcPending
																	if v1836 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v1680, int32(1))
																		mBase = m.M
																		v1839 = m.ExcPending
																		if v1839 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
																		}
																	}
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+1008)) = v1721
																F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_28), v11+int32(1008))
																mBase = m.M
																v1780 = m.ExcPending
																if v1780 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1694)
																	mBase = m.M
																	v1836 = m.ExcPending
																	if v1836 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v1680, int32(1))
																		mBase = m.M
																		v1839 = m.ExcPending
																		if v1839 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
																		}
																	}
																}
															}
														case 34:
															*(*int32)(unsafe.Add(mBase, uint32(v11)+1072)) = v1721
															F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_29), v11+int32(1072))
															mBase = m.M
															v1805 = m.ExcPending
															if v1805 != 0 {
																return int32(0)
															} else {
																F_systable_endscan(m, v1694)
																mBase = m.M
																v1836 = m.ExcPending
																if v1836 != 0 {
																	return int32(0)
																} else {
																	F_relation_close(m, v1680, int32(1))
																	mBase = m.M
																	v1839 = m.ExcPending
																	if v1839 != 0 {
																		return int32(0)
																	} else {
																		v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																		v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																		m.G0 = v11 + int32(1440)
																		if v1847 != 0 {
																			v1853 = v1848
																		} else {
																			v1853 = int32(0)
																		}
																		return v1853
																	}
																}
															}
														case 38:
															if v1726 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+964)) = v1726
																*(*int32)(unsafe.Add(mBase, uint32(v11)+960)) = v1721
																F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_30), v11+int32(960))
																mBase = m.M
																v1738 = m.ExcPending
																if v1738 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1694)
																	mBase = m.M
																	v1836 = m.ExcPending
																	if v1836 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v1680, int32(1))
																		mBase = m.M
																		v1839 = m.ExcPending
																		if v1839 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
																		}
																	}
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+944)) = v1721
																F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_31), v11+int32(944))
																mBase = m.M
																v1746 = m.ExcPending
																if v1746 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v1694)
																	mBase = m.M
																	v1836 = m.ExcPending
																	if v1836 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v1680, int32(1))
																		mBase = m.M
																		v1839 = m.ExcPending
																		if v1839 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
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
							v1646 = m.ExcPending
							if v1646 != 0 {
								return int32(0)
							} else {
								v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1647
								F_errmsg_internal(m, int32(_a_F_getObjectDescription_4), v11)
								mBase = m.M
								v1651 = m.ExcPending
								if v1651 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(_a_F_getObjectDescription_5), int32(_a_F_getObjectDescription_2))
									mBase = m.M
									v1656 = m.ExcPending
									if v1656 != 0 {
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
					v514 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
					v515 = F_SearchSysCache1(m, int32(2), v514)
					mBase = m.M
					v516 = m.ExcPending
					if v516 != 0 {
						return int32(0)
					} else {
						if v515 == int32(0) {
							if l1 != 0 {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v522 = m.ExcPending
								if v522 != 0 {
									return int32(0)
								} else {
									v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+336)) = v523
									F_errmsg_internal(m, int32(_a_F_getObjectDescription_32), v11+int32(336))
									mBase = m.M
									v529 = m.ExcPending
									if v529 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3269), int32(_a_F_getObjectDescription_2))
										mBase = m.M
										v534 = m.ExcPending
										if v534 != 0 {
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
							v535 = *(*int32)(unsafe.Add(mBase, uint32(v515)+16))
							v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535)+22)))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+352)) = v535 + v536 + int32(4)
							F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_33), v11+int32(352))
							mBase = m.M
							v547 = m.ExcPending
							if v547 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v515)
								mBase = m.M
								v549 = m.ExcPending
								if v549 != 0 {
									return int32(0)
								} else {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								}
							}
						}
					}
				case 1:
					v552 = F_table_open(m, int32(2602), int32(1))
					mBase = m.M
					v553 = m.ExcPending
					if v553 != 0 {
						return int32(0)
					} else {
						v555 = v11 + int32(1360)
						v559 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						F_ScanKeyInit(m, v555, int32(1), int32(3), int32(184), v559)
						mBase = m.M
						v561 = m.ExcPending
						if v561 != 0 {
							return int32(0)
						} else {
							v563 = int32(1)
							v566 = F_systable_beginscan(m, v552, int32(2756), v563, int32(0), v563, v555)
							mBase = m.M
							v567 = m.ExcPending
							if v567 != 0 {
								return int32(0)
							} else {
								v568 = F_systable_getnext(m, v566)
								mBase = m.M
								v569 = m.ExcPending
								if v569 != 0 {
									return int32(0)
								} else {
									if v568 == int32(0) {
										if l1 != 0 {
											F_systable_endscan(m, v566)
											mBase = m.M
											v635 = m.ExcPending
											if v635 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v552, int32(1))
												mBase = m.M
												v638 = m.ExcPending
												if v638 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v575 = m.ExcPending
											if v575 != 0 {
												return int32(0)
											} else {
												v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+368)) = v576
												F_errmsg_internal(m, int32(_a_F_getObjectDescription_34), v11+int32(368))
												mBase = m.M
												v582 = m.ExcPending
												if v582 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3305), int32(_a_F_getObjectDescription_2))
													mBase = m.M
													v587 = m.ExcPending
													if v587 != 0 {
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
										v588 = *(*int32)(unsafe.Add(mBase, uint32(v568)+16))
										v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588)+22)))
										v591 = v11 + int32(1344)
										F_initStringInfo(m, v591)
										mBase = m.M
										v593 = m.ExcPending
										if v593 != 0 {
											return int32(0)
										} else {
											v594 = v588 + v589
											v595 = *(*int32)(unsafe.Add(mBase, uint32(v594)+4))
											F_getOpFamilyDescription(m, v591, v595, int32(0))
											mBase = m.M
											v598 = m.ExcPending
											if v598 != 0 {
												return int32(0)
											} else {
												v599 = int32(*(*int16)(unsafe.Add(mBase, uint32(v594)+16)))
												v600 = *(*int32)(unsafe.Add(mBase, uint32(v594)+8))
												v603 = F_format_type_extended(m, v600, int32(-1), int32(2))
												mBase = m.M
												v604 = m.ExcPending
												if v604 != 0 {
													return int32(0)
												} else {
													v605 = *(*int32)(unsafe.Add(mBase, uint32(v594)+12))
													v608 = F_format_type_extended(m, v605, int32(-1), int32(2))
													mBase = m.M
													v609 = m.ExcPending
													if v609 != 0 {
														return int32(0)
													} else {
														v610 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
														v611 = *(*int32)(unsafe.Add(mBase, uint32(v594)+20))
														v612 = F_format_operator(m, v611)
														mBase = m.M
														v613 = m.ExcPending
														if v613 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+400)) = v612
															*(*int32)(unsafe.Add(mBase, uint32(v11)+396)) = v610
															*(*int32)(unsafe.Add(mBase, uint32(v11)+392)) = v608
															*(*int32)(unsafe.Add(mBase, uint32(v11)+388)) = v603
															*(*int32)(unsafe.Add(mBase, uint32(v11)+384)) = v599
															F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_35), v11+int32(384))
															mBase = m.M
															v625 = m.ExcPending
															if v625 != 0 {
																return int32(0)
															} else {
																v626 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
																F_pfree(m, v626)
																mBase = m.M
																v628 = m.ExcPending
																if v628 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v566)
																	mBase = m.M
																	v635 = m.ExcPending
																	if v635 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v552, int32(1))
																		mBase = m.M
																		v638 = m.ExcPending
																		if v638 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
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
					v641 = F_table_open(m, int32(2603), int32(1))
					mBase = m.M
					v642 = m.ExcPending
					if v642 != 0 {
						return int32(0)
					} else {
						v644 = v11 + int32(1360)
						v648 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						F_ScanKeyInit(m, v644, int32(1), int32(3), int32(184), v648)
						mBase = m.M
						v650 = m.ExcPending
						if v650 != 0 {
							return int32(0)
						} else {
							v652 = int32(1)
							v655 = F_systable_beginscan(m, v641, int32(2757), v652, int32(0), v652, v644)
							mBase = m.M
							v656 = m.ExcPending
							if v656 != 0 {
								return int32(0)
							} else {
								v657 = F_systable_getnext(m, v655)
								mBase = m.M
								v658 = m.ExcPending
								if v658 != 0 {
									return int32(0)
								} else {
									if v657 == int32(0) {
										if l1 != 0 {
											F_systable_endscan(m, v655)
											mBase = m.M
											v724 = m.ExcPending
											if v724 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v641, int32(1))
												mBase = m.M
												v727 = m.ExcPending
												if v727 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v664 = m.ExcPending
											if v664 != 0 {
												return int32(0)
											} else {
												v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+416)) = v665
												F_errmsg_internal(m, int32(_a_F_getObjectDescription_36), v11+int32(416))
												mBase = m.M
												v671 = m.ExcPending
												if v671 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3370), int32(_a_F_getObjectDescription_2))
													mBase = m.M
													v676 = m.ExcPending
													if v676 != 0 {
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
										v677 = *(*int32)(unsafe.Add(mBase, uint32(v657)+16))
										v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+22)))
										v680 = v11 + int32(1344)
										F_initStringInfo(m, v680)
										mBase = m.M
										v682 = m.ExcPending
										if v682 != 0 {
											return int32(0)
										} else {
											v683 = v677 + v678
											v684 = *(*int32)(unsafe.Add(mBase, uint32(v683)+4))
											F_getOpFamilyDescription(m, v680, v684, int32(0))
											mBase = m.M
											v687 = m.ExcPending
											if v687 != 0 {
												return int32(0)
											} else {
												v688 = int32(*(*int16)(unsafe.Add(mBase, uint32(v683)+16)))
												v689 = *(*int32)(unsafe.Add(mBase, uint32(v683)+8))
												v692 = F_format_type_extended(m, v689, int32(-1), int32(2))
												mBase = m.M
												v693 = m.ExcPending
												if v693 != 0 {
													return int32(0)
												} else {
													v694 = *(*int32)(unsafe.Add(mBase, uint32(v683)+12))
													v697 = F_format_type_extended(m, v694, int32(-1), int32(2))
													mBase = m.M
													v698 = m.ExcPending
													if v698 != 0 {
														return int32(0)
													} else {
														v699 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
														v700 = *(*int32)(unsafe.Add(mBase, uint32(v683)+20))
														v701 = F_format_procedure(m, v700)
														mBase = m.M
														v702 = m.ExcPending
														if v702 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+448)) = v701
															*(*int32)(unsafe.Add(mBase, uint32(v11)+444)) = v699
															*(*int32)(unsafe.Add(mBase, uint32(v11)+440)) = v697
															*(*int32)(unsafe.Add(mBase, uint32(v11)+436)) = v692
															*(*int32)(unsafe.Add(mBase, uint32(v11)+432)) = v688
															F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_37), v11+int32(432))
															mBase = m.M
															v714 = m.ExcPending
															if v714 != 0 {
																return int32(0)
															} else {
																v715 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
																F_pfree(m, v715)
																mBase = m.M
																v717 = m.ExcPending
																if v717 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v655)
																	mBase = m.M
																	v724 = m.ExcPending
																	if v724 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v641, int32(1))
																		mBase = m.M
																		v727 = m.ExcPending
																		if v727 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
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
					v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_GetAttrDefaultColumnAddress(m, v11+int32(1360), v372)
					mBase = m.M
					v374 = m.ExcPending
					if v374 != 0 {
						return int32(0)
					} else {
						v375 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1364))
						if v375 == int32(0) {
							if l1 != 0 {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v381 = m.ExcPending
								if v381 != 0 {
									return int32(0)
								} else {
									v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v382
									F_errmsg_internal(m, int32(_a_F_getObjectDescription_38), v11+int32(208))
									mBase = m.M
									v388 = m.ExcPending
									if v388 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3171), int32(_a_F_getObjectDescription_2))
										mBase = m.M
										v393 = m.ExcPending
										if v393 != 0 {
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
							v397 = F_getObjectDescription(m, v11+int32(1360), int32(0))
							mBase = m.M
							v398 = m.ExcPending
							if v398 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+224)) = v397
								F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_39), v11+int32(224))
								mBase = m.M
								v406 = m.ExcPending
								if v406 != 0 {
									return int32(0)
								} else {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								}
							}
						}
					}
				case 4:
					v151 = F_table_open(m, int32(2605), int32(1))
					mBase = m.M
					v152 = m.ExcPending
					if v152 != 0 {
						return int32(0)
					} else {
						v154 = v11 + int32(1360)
						v158 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						F_ScanKeyInit(m, v154, int32(1), int32(3), int32(184), v158)
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return int32(0)
						} else {
							v162 = int32(1)
							v165 = F_systable_beginscan(m, v151, int32(2660), v162, int32(0), v162, v154)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								v167 = F_systable_getnext(m, v165)
								mBase = m.M
								v168 = m.ExcPending
								if v168 != 0 {
									return int32(0)
								} else {
									if v167 == int32(0) {
										if l1 != 0 {
											F_systable_endscan(m, v165)
											mBase = m.M
											v208 = m.ExcPending
											if v208 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v151, int32(1))
												mBase = m.M
												v211 = m.ExcPending
												if v211 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return int32(0)
											} else {
												v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v175
												F_errmsg_internal(m, int32(_a_F_getObjectDescription_40), v11-int32(-64))
												mBase = m.M
												v181 = m.ExcPending
												if v181 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3043), int32(_a_F_getObjectDescription_2))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
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
										v187 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
										v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+22)))
										v189 = v187 + v188
										v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
										v191 = F_format_type_be(m, v190)
										mBase = m.M
										v192 = m.ExcPending
										if v192 != 0 {
											return int32(0)
										} else {
											v193 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
											v194 = F_format_type_be(m, v193)
											mBase = m.M
											v195 = m.ExcPending
											if v195 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v194
												*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v191
												F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_41), v11+int32(80))
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return int32(0)
												} else {
													F_systable_endscan(m, v165)
													mBase = m.M
													v208 = m.ExcPending
													if v208 != 0 {
														return int32(0)
													} else {
														F_relation_close(m, v151, int32(1))
														mBase = m.M
														v211 = m.ExcPending
														if v211 != 0 {
															return int32(0)
														} else {
															v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
															v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
															m.G0 = v11 + int32(1440)
															if v1847 != 0 {
																v1853 = v1848
															} else {
																v1853 = int32(0)
															}
															return v1853
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
					v260 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
					v261 = F_SearchSysCache1(m, int32(19), v260)
					mBase = m.M
					v262 = m.ExcPending
					if v262 != 0 {
						return int32(0)
					} else {
						if v261 == int32(0) {
							if l1 != 0 {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v268 = m.ExcPending
								if v268 != 0 {
									return int32(0)
								} else {
									v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v269
									F_errmsg_internal(m, int32(_a_F_getObjectDescription_42), v11+int32(128))
									mBase = m.M
									v275 = m.ExcPending
									if v275 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3103), int32(_a_F_getObjectDescription_2))
										mBase = m.M
										v280 = m.ExcPending
										if v280 != 0 {
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
							v281 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
							v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+22)))
							v283 = v281 + v282
							v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+80))
							if v284 != 0 {
								v286 = v11 + int32(1360)
								F_initStringInfo(m, v286)
								mBase = m.M
								v288 = m.ExcPending
								if v288 != 0 {
									return int32(0)
								} else {
									v289 = *(*int32)(unsafe.Add(mBase, uint32(v283)+80))
									F_getRelationDescription(m, v286, v289, int32(0))
									mBase = m.M
									v292 = m.ExcPending
									if v292 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v283 + int32(4)
										v296 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = v296
										F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_43), v11+int32(160))
										mBase = m.M
										v304 = m.ExcPending
										if v304 != 0 {
											return int32(0)
										} else {
											v305 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
											F_pfree(m, v305)
											mBase = m.M
											v307 = m.ExcPending
											if v307 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v261)
												mBase = m.M
												v309 = m.ExcPending
												if v309 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = v283 + int32(4)
								F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_44), v11+int32(144))
								mBase = m.M
								v319 = m.ExcPending
								if v319 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v261)
									mBase = m.M
									v321 = m.ExcPending
									if v321 != 0 {
										return int32(0)
									} else {
										v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
										v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
										m.G0 = v11 + int32(1440)
										if v1847 != 0 {
											v1853 = v1848
										} else {
											v1853 = int32(0)
										}
										return v1853
									}
								}
							}
						}
					}
				case 6:
					v323 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
					v324 = F_SearchSysCache1(m, int32(20), v323)
					mBase = m.M
					v325 = m.ExcPending
					if v325 != 0 {
						return int32(0)
					} else {
						if v324 == int32(0) {
							if l1 != 0 {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v331 = m.ExcPending
								if v331 != 0 {
									return int32(0)
								} else {
									v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = v332
									F_errmsg_internal(m, int32(_a_F_getObjectDescription_45), v11+int32(176))
									mBase = m.M
									v338 = m.ExcPending
									if v338 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3142), int32(_a_F_getObjectDescription_2))
										mBase = m.M
										v343 = m.ExcPending
										if v343 != 0 {
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
							v344 = *(*int32)(unsafe.Add(mBase, uint32(v324)+16))
							v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+22)))
							v346 = v344 + v345
							v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v349 = F_ConversionIsVisibleExt(m, v347, int32(0))
							mBase = m.M
							v350 = m.ExcPending
							if v350 != 0 {
								return int32(0)
							} else {
								if v349 != 0 {
									v355 = int32(0)
									v358 = F_quote_qualified_identifier(m, v355, v346+int32(4))
									mBase = m.M
									v359 = m.ExcPending
									if v359 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v358
										F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_46), v11+int32(192))
										mBase = m.M
										v367 = m.ExcPending
										if v367 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v324)
											mBase = m.M
											v369 = m.ExcPending
											if v369 != 0 {
												return int32(0)
											} else {
												v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
												v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
												m.G0 = v11 + int32(1440)
												if v1847 != 0 {
													v1853 = v1848
												} else {
													v1853 = int32(0)
												}
												return v1853
											}
										}
									}
								} else {
									v352 = *(*int32)(unsafe.Add(mBase, uint32(v346)+68))
									v353 = F_get_namespace_name(m, v352)
									mBase = m.M
									v354 = m.ExcPending
									if v354 != 0 {
										return int32(0)
									} else {
										v355 = v353
										v358 = F_quote_qualified_identifier(m, v355, v346+int32(4))
										mBase = m.M
										v359 = m.ExcPending
										if v359 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v358
											F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_46), v11+int32(192))
											mBase = m.M
											v367 = m.ExcPending
											if v367 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v324)
												mBase = m.M
												v369 = m.ExcPending
												if v369 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
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
					v1646 = m.ExcPending
					if v1646 != 0 {
						return int32(0)
					} else {
						v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1647
						F_errmsg_internal(m, int32(_a_F_getObjectDescription_4), v11)
						mBase = m.M
						v1651 = m.ExcPending
						if v1651 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(_a_F_getObjectDescription_5), int32(_a_F_getObjectDescription_2))
							mBase = m.M
							v1656 = m.ExcPending
							if v1656 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 11:
					v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v408 = F_get_language_name(m, v407, l1)
					mBase = m.M
					v409 = m.ExcPending
					if v409 != 0 {
						return int32(0)
					} else {
						if v408 == int32(0) {
							v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
							v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
							m.G0 = v11 + int32(1440)
							if v1847 != 0 {
								v1853 = v1848
							} else {
								v1853 = int32(0)
							}
							return v1853
						} else {
							v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v414 = F_get_language_name(m, v412, int32(0))
							mBase = m.M
							v415 = m.ExcPending
							if v415 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+240)) = v414
								F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_47), v11+int32(240))
								mBase = m.M
								v423 = m.ExcPending
								if v423 != 0 {
									return int32(0)
								} else {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								}
							}
						}
					}
				case 12:
					v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v425 = F_LargeObjectExists(m, v424)
					mBase = m.M
					v426 = m.ExcPending
					if v426 != 0 {
						return int32(0)
					} else {
						if v425 == int32(0) {
							v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
							v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
							m.G0 = v11 + int32(1440)
							if v1847 != 0 {
								v1853 = v1848
							} else {
								v1853 = int32(0)
							}
							return v1853
						} else {
							v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+256)) = v429
							F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_48), v11+int32(256))
							mBase = m.M
							v437 = m.ExcPending
							if v437 != 0 {
								return int32(0)
							} else {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							}
						}
					}
				case 14:
					v872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v873 = F_get_namespace_name(m, v872)
					mBase = m.M
					v874 = m.ExcPending
					if v874 != 0 {
						return int32(0)
					} else {
						if v873 == int32(0) {
							if l1 != 0 {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v880 = m.ExcPending
								if v880 != 0 {
									return int32(0)
								} else {
									v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+528)) = v881
									F_errmsg_internal(m, int32(_a_F_getObjectDescription_49), v11+int32(528))
									mBase = m.M
									v887 = m.ExcPending
									if v887 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3510), int32(_a_F_getObjectDescription_2))
										mBase = m.M
										v892 = m.ExcPending
										if v892 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v11)+544)) = v873
							F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_50), v11+int32(544))
							mBase = m.M
							v900 = m.ExcPending
							if v900 != 0 {
								return int32(0)
							} else {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							}
						}
					}
				case 15:
					v453 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
					v454 = F_SearchSysCache1(m, int32(14), v453)
					mBase = m.M
					v455 = m.ExcPending
					if v455 != 0 {
						return int32(0)
					} else {
						if v454 == int32(0) {
							if l1 != 0 {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v461 = m.ExcPending
								if v461 != 0 {
									return int32(0)
								} else {
									v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+288)) = v462
									F_errmsg_internal(m, int32(_a_F_getObjectDescription_51), v11+int32(288))
									mBase = m.M
									v468 = m.ExcPending
									if v468 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3226), int32(_a_F_getObjectDescription_2))
										mBase = m.M
										v473 = m.ExcPending
										if v473 != 0 {
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
							v475 = *(*int32)(unsafe.Add(mBase, uint32(v454)+16))
							v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+22)))
							v477 = v475 + v476
							v478 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v477)+4)))
							v479 = F_SearchSysCache1(m, int32(2), v478)
							mBase = m.M
							v480 = m.ExcPending
							if v480 != 0 {
								return int32(0)
							} else {
								if v479 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1665 = m.ExcPending
									if v1665 != 0 {
										return int32(0)
									} else {
										v1666 = *(*int32)(unsafe.Add(mBase, uint32(v477)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+304)) = v1666
										F_errmsg_internal(m, int32(_a_F_getObjectDescription_32), v11+int32(304))
										mBase = m.M
										v1672 = m.ExcPending
										if v1672 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3236), int32(_a_F_getObjectDescription_2))
											mBase = m.M
											v1677 = m.ExcPending
											if v1677 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v483 = *(*int32)(unsafe.Add(mBase, uint32(v479)+16))
									v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+22)))
									v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v487 = F_OpclassIsVisible(m, v486)
									mBase = m.M
									v488 = m.ExcPending
									if v488 != 0 {
										return int32(0)
									} else {
										if v487 != 0 {
											v493 = int32(0)
											v496 = F_quote_qualified_identifier(m, v493, v477+int32(8))
											mBase = m.M
											v497 = m.ExcPending
											if v497 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+324)) = v483 + v484 + int32(4)
												*(*int32)(unsafe.Add(mBase, uint32(v11)+320)) = v496
												F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_52), v11+int32(320))
												mBase = m.M
												v508 = m.ExcPending
												if v508 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v479)
													mBase = m.M
													v510 = m.ExcPending
													if v510 != 0 {
														return int32(0)
													} else {
														F_ReleaseCatCache(m, v454)
														mBase = m.M
														v512 = m.ExcPending
														if v512 != 0 {
															return int32(0)
														} else {
															v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
															v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
															m.G0 = v11 + int32(1440)
															if v1847 != 0 {
																v1853 = v1848
															} else {
																v1853 = int32(0)
															}
															return v1853
														}
													}
												}
											}
										} else {
											v490 = *(*int32)(unsafe.Add(mBase, uint32(v477)+72))
											v491 = F_get_namespace_name(m, v490)
											mBase = m.M
											v492 = m.ExcPending
											if v492 != 0 {
												return int32(0)
											} else {
												v493 = v491
												v496 = F_quote_qualified_identifier(m, v493, v477+int32(8))
												mBase = m.M
												v497 = m.ExcPending
												if v497 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+324)) = v483 + v484 + int32(4)
													*(*int32)(unsafe.Add(mBase, uint32(v11)+320)) = v496
													F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_52), v11+int32(320))
													mBase = m.M
													v508 = m.ExcPending
													if v508 != 0 {
														return int32(0)
													} else {
														F_ReleaseCatCache(m, v479)
														mBase = m.M
														v510 = m.ExcPending
														if v510 != 0 {
															return int32(0)
														} else {
															F_ReleaseCatCache(m, v454)
															mBase = m.M
															v512 = m.ExcPending
															if v512 != 0 {
																return int32(0)
															} else {
																v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																m.G0 = v11 + int32(1440)
																if v1847 != 0 {
																	v1853 = v1848
																} else {
																	v1853 = int32(0)
																}
																return v1853
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
					v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v440 = F_format_operator_extended(m, v438, int32(1))
					mBase = m.M
					v441 = m.ExcPending
					if v441 != 0 {
						return int32(0)
					} else {
						if v440 == int32(0) {
							v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
							v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
							m.G0 = v11 + int32(1440)
							if v1847 != 0 {
								v1853 = v1848
							} else {
								v1853 = int32(0)
							}
							return v1853
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+272)) = v440
							F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_53), v11+int32(272))
							mBase = m.M
							v451 = m.ExcPending
							if v451 != 0 {
								return int32(0)
							} else {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							}
						}
					}
				case 17:
					v730 = F_table_open(m, int32(2618), int32(1))
					mBase = m.M
					v731 = m.ExcPending
					if v731 != 0 {
						return int32(0)
					} else {
						v733 = v11 + int32(1360)
						v737 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						F_ScanKeyInit(m, v733, int32(1), int32(3), int32(184), v737)
						mBase = m.M
						v739 = m.ExcPending
						if v739 != 0 {
							return int32(0)
						} else {
							v741 = int32(1)
							v744 = F_systable_beginscan(m, v730, int32(2692), v741, int32(0), v741, v733)
							mBase = m.M
							v745 = m.ExcPending
							if v745 != 0 {
								return int32(0)
							} else {
								v746 = F_systable_getnext(m, v744)
								mBase = m.M
								v747 = m.ExcPending
								if v747 != 0 {
									return int32(0)
								} else {
									if v746 == int32(0) {
										if l1 != 0 {
											F_systable_endscan(m, v744)
											mBase = m.M
											v796 = m.ExcPending
											if v796 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v730, int32(1))
												mBase = m.M
												v799 = m.ExcPending
												if v799 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v753 = m.ExcPending
											if v753 != 0 {
												return int32(0)
											} else {
												v754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+464)) = v754
												F_errmsg_internal(m, int32(_a_F_getObjectDescription_54), v11+int32(464))
												mBase = m.M
												v760 = m.ExcPending
												if v760 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3434), int32(_a_F_getObjectDescription_2))
													mBase = m.M
													v765 = m.ExcPending
													if v765 != 0 {
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
										v766 = *(*int32)(unsafe.Add(mBase, uint32(v746)+16))
										v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766)+22)))
										v769 = v11 + int32(1344)
										F_initStringInfo(m, v769)
										mBase = m.M
										v771 = m.ExcPending
										if v771 != 0 {
											return int32(0)
										} else {
											v772 = v766 + v767
											v773 = *(*int32)(unsafe.Add(mBase, uint32(v772)+68))
											F_getRelationDescription(m, v769, v773, int32(0))
											mBase = m.M
											v776 = m.ExcPending
											if v776 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+480)) = v772 + int32(4)
												v780 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+484)) = v780
												F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_55), v11+int32(480))
												mBase = m.M
												v788 = m.ExcPending
												if v788 != 0 {
													return int32(0)
												} else {
													v789 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
													F_pfree(m, v789)
													mBase = m.M
													v791 = m.ExcPending
													if v791 != 0 {
														return int32(0)
													} else {
														F_systable_endscan(m, v744)
														mBase = m.M
														v796 = m.ExcPending
														if v796 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v730, int32(1))
															mBase = m.M
															v799 = m.ExcPending
															if v799 != 0 {
																return int32(0)
															} else {
																v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																m.G0 = v11 + int32(1440)
																if v1847 != 0 {
																	v1853 = v1848
																} else {
																	v1853 = int32(0)
																}
																return v1853
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
					v802 = F_table_open(m, int32(2620), int32(1))
					mBase = m.M
					v803 = m.ExcPending
					if v803 != 0 {
						return int32(0)
					} else {
						v805 = v11 + int32(1360)
						v809 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						F_ScanKeyInit(m, v805, int32(1), int32(3), int32(184), v809)
						mBase = m.M
						v811 = m.ExcPending
						if v811 != 0 {
							return int32(0)
						} else {
							v813 = int32(1)
							v816 = F_systable_beginscan(m, v802, int32(2702), v813, int32(0), v813, v805)
							mBase = m.M
							v817 = m.ExcPending
							if v817 != 0 {
								return int32(0)
							} else {
								v818 = F_systable_getnext(m, v816)
								mBase = m.M
								v819 = m.ExcPending
								if v819 != 0 {
									return int32(0)
								} else {
									if v818 == int32(0) {
										if l1 != 0 {
											F_systable_endscan(m, v816)
											mBase = m.M
											v868 = m.ExcPending
											if v868 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v802, int32(1))
												mBase = m.M
												v871 = m.ExcPending
												if v871 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v825 = m.ExcPending
											if v825 != 0 {
												return int32(0)
											} else {
												v826 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+496)) = v826
												F_errmsg_internal(m, int32(_a_F_getObjectDescription_56), v11+int32(496))
												mBase = m.M
												v832 = m.ExcPending
												if v832 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3480), int32(_a_F_getObjectDescription_2))
													mBase = m.M
													v837 = m.ExcPending
													if v837 != 0 {
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
										v838 = *(*int32)(unsafe.Add(mBase, uint32(v818)+16))
										v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838)+22)))
										v841 = v11 + int32(1344)
										F_initStringInfo(m, v841)
										mBase = m.M
										v843 = m.ExcPending
										if v843 != 0 {
											return int32(0)
										} else {
											v844 = v838 + v839
											v845 = *(*int32)(unsafe.Add(mBase, uint32(v844)+4))
											F_getRelationDescription(m, v841, v845, int32(0))
											mBase = m.M
											v848 = m.ExcPending
											if v848 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+512)) = v844 + int32(12)
												v852 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+516)) = v852
												F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_57), v11+int32(512))
												mBase = m.M
												v860 = m.ExcPending
												if v860 != 0 {
													return int32(0)
												} else {
													v861 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
													F_pfree(m, v861)
													mBase = m.M
													v863 = m.ExcPending
													if v863 != 0 {
														return int32(0)
													} else {
														F_systable_endscan(m, v816)
														mBase = m.M
														v868 = m.ExcPending
														if v868 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v802, int32(1))
															mBase = m.M
															v871 = m.ExcPending
															if v871 != 0 {
																return int32(0)
															} else {
																v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																m.G0 = v11 + int32(1440)
																if v1847 != 0 {
																	v1853 = v1848
																} else {
																	v1853 = int32(0)
																}
																return v1853
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
						v1646 = m.ExcPending
						if v1646 != 0 {
							return int32(0)
						} else {
							v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1647
							F_errmsg_internal(m, int32(_a_F_getObjectDescription_4), v11)
							mBase = m.M
							v1651 = m.ExcPending
							if v1651 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(_a_F_getObjectDescription_5), int32(_a_F_getObjectDescription_2))
								mBase = m.M
								v1656 = m.ExcPending
								if v1656 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v1281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1282 = F_GetForeignDataWrapperExtended(m, v1281, l1)
						mBase = m.M
						v1283 = m.ExcPending
						if v1283 != 0 {
							return int32(0)
						} else {
							if v1282 == int32(0) {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							} else {
								v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1282)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+832)) = v1286
								F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_58), v11+int32(832))
								mBase = m.M
								v1294 = m.ExcPending
								if v1294 != 0 {
									return int32(0)
								} else {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
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
						v1659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_getOpFamilyDescription(m, v11+int32(1424), v1659, l1)
						mBase = m.M
						v1661 = m.ExcPending
						if v1661 != 0 {
							return int32(0)
						} else {
							v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
							v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
							m.G0 = v11 + int32(1440)
							if v1847 != 0 {
								v1853 = v1848
							} else {
								v1853 = int32(0)
							}
							return v1853
						}
					} else {
						if v19 == int32(3079) {
							v1355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v1356 = F_get_extension_name(m, v1355)
							mBase = m.M
							v1357 = m.ExcPending
							if v1357 != 0 {
								return int32(0)
							} else {
								if v1356 == int32(0) {
									if l1 != 0 {
										v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
										v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
										m.G0 = v11 + int32(1440)
										if v1847 != 0 {
											v1853 = v1848
										} else {
											v1853 = int32(0)
										}
										return v1853
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v1363 = m.ExcPending
										if v1363 != 0 {
											return int32(0)
										} else {
											v1364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+1104)) = v1364
											F_errmsg_internal(m, int32(_a_F_getObjectDescription_59), v11+int32(1104))
											mBase = m.M
											v1370 = m.ExcPending
											if v1370 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3934), int32(_a_F_getObjectDescription_2))
												mBase = m.M
												v1375 = m.ExcPending
												if v1375 != 0 {
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
									*(*int32)(unsafe.Add(mBase, uint32(v11)+1120)) = v1356
									F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_60), v11+int32(1120))
									mBase = m.M
									v1383 = m.ExcPending
									if v1383 != 0 {
										return int32(0)
									} else {
										v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
										v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
										m.G0 = v11 + int32(1440)
										if v1847 != 0 {
											v1853 = v1848
										} else {
											v1853 = int32(0)
										}
										return v1853
									}
								}
							}
						} else {
							if v19 != int32(3256) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v1646 = m.ExcPending
								if v1646 != 0 {
									return int32(0)
								} else {
									v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1647
									F_errmsg_internal(m, int32(_a_F_getObjectDescription_4), v11)
									mBase = m.M
									v1651 = m.ExcPending
									if v1651 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(_a_F_getObjectDescription_5), int32(_a_F_getObjectDescription_2))
										mBase = m.M
										v1656 = m.ExcPending
										if v1656 != 0 {
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
									v45 = v11 + int32(1360)
									v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
									F_ScanKeyInit(m, v45, int32(1), int32(3), int32(184), v49)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										v53 = int32(1)
										v56 = F_systable_beginscan(m, v42, int32(3257), v53, int32(0), v53, v45)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v58 = F_systable_getnext(m, v56)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int32(0)
											} else {
												if v58 != 0 {
													v1460 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
													v1461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1460)+22)))
													v1463 = v11 + int32(1344)
													F_initStringInfo(m, v1463)
													mBase = m.M
													v1465 = m.ExcPending
													if v1465 != 0 {
														return int32(0)
													} else {
														v1466 = v1460 + v1461
														v1467 = *(*int32)(unsafe.Add(mBase, uint32(v1466)+68))
														F_getRelationDescription(m, v1463, v1467, int32(0))
														mBase = m.M
														v1470 = m.ExcPending
														if v1470 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+1216)) = v1466 + int32(4)
															v1474 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+1220)) = v1474
															F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_61), v11+int32(1216))
															mBase = m.M
															v1482 = m.ExcPending
															if v1482 != 0 {
																return int32(0)
															} else {
																v1483 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
																F_pfree(m, v1483)
																mBase = m.M
																v1485 = m.ExcPending
																if v1485 != 0 {
																	return int32(0)
																} else {
																	F_systable_endscan(m, v56)
																	mBase = m.M
																	v1490 = m.ExcPending
																	if v1490 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v42, int32(1))
																		mBase = m.M
																		v1493 = m.ExcPending
																		if v1493 != 0 {
																			return int32(0)
																		} else {
																			v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																			v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																			m.G0 = v11 + int32(1440)
																			if v1847 != 0 {
																				v1853 = v1848
																			} else {
																				v1853 = int32(0)
																			}
																			return v1853
																		}
																	}
																}
															}
														}
													}
												} else {
													if l1 != 0 {
														F_systable_endscan(m, v56)
														mBase = m.M
														v1490 = m.ExcPending
														if v1490 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v42, int32(1))
															mBase = m.M
															v1493 = m.ExcPending
															if v1493 != 0 {
																return int32(0)
															} else {
																v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
																v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
																m.G0 = v11 + int32(1440)
																if v1847 != 0 {
																	v1853 = v1848
																} else {
																	v1853 = int32(0)
																}
																return v1853
															}
														}
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
															return int32(0)
														} else {
															v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+1200)) = v64
															F_errmsg_internal(m, int32(_a_F_getObjectDescription_62), v11+int32(1200))
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(4008), int32(_a_F_getObjectDescription_2))
																mBase = m.M
																v75 = m.ExcPending
																if v75 != 0 {
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
						v213 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						v214 = F_SearchSysCache1(m, int32(16), v213)
						mBase = m.M
						v215 = m.ExcPending
						if v215 != 0 {
							return int32(0)
						} else {
							if v214 == int32(0) {
								if l1 != 0 {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return int32(0)
									} else {
										v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v222
										F_errmsg_internal(m, int32(_a_F_getObjectDescription_63), v11+int32(96))
										mBase = m.M
										v228 = m.ExcPending
										if v228 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3073), int32(_a_F_getObjectDescription_2))
											mBase = m.M
											v233 = m.ExcPending
											if v233 != 0 {
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
								v234 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
								v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+22)))
								v236 = v234 + v235
								v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v238 = F_CollationIsVisible(m, v237)
								mBase = m.M
								v239 = m.ExcPending
								if v239 != 0 {
									return int32(0)
								} else {
									if v238 != 0 {
										v244 = int32(0)
										v247 = F_quote_qualified_identifier(m, v244, v236+int32(4))
										mBase = m.M
										v248 = m.ExcPending
										if v248 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v247
											F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_64), v11+int32(112))
											mBase = m.M
											v256 = m.ExcPending
											if v256 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v214)
												mBase = m.M
												v258 = m.ExcPending
												if v258 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
												}
											}
										}
									} else {
										v241 = *(*int32)(unsafe.Add(mBase, uint32(v236)+68))
										v242 = F_get_namespace_name(m, v241)
										mBase = m.M
										v243 = m.ExcPending
										if v243 != 0 {
											return int32(0)
										} else {
											v244 = v242
											v247 = F_quote_qualified_identifier(m, v244, v236+int32(4))
											mBase = m.M
											v248 = m.ExcPending
											if v248 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v247
												F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_64), v11+int32(112))
												mBase = m.M
												v256 = m.ExcPending
												if v256 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v214)
													mBase = m.M
													v258 = m.ExcPending
													if v258 != 0 {
														return int32(0)
													} else {
														v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
														v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
														m.G0 = v11 + int32(1440)
														if v1847 != 0 {
															v1853 = v1848
														} else {
															v1853 = int32(0)
														}
														return v1853
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
						v1646 = m.ExcPending
						if v1646 != 0 {
							return int32(0)
						} else {
							v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1647
							F_errmsg_internal(m, int32(_a_F_getObjectDescription_4), v11)
							mBase = m.M
							v1651 = m.ExcPending
							if v1651 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(_a_F_getObjectDescription_5), int32(_a_F_getObjectDescription_2))
								mBase = m.M
								v1656 = m.ExcPending
								if v1656 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 10:
						v1385 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						v1386 = F_SearchSysCache1(m, int32(26), v1385)
						mBase = m.M
						v1387 = m.ExcPending
						if v1387 != 0 {
							return int32(0)
						} else {
							if v1386 == int32(0) {
								if l1 != 0 {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1393 = m.ExcPending
									if v1393 != 0 {
										return int32(0)
									} else {
										v1394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+1136)) = v1394
										F_errmsg_internal(m, int32(_a_F_getObjectDescription_65), v11+int32(1136))
										mBase = m.M
										v1400 = m.ExcPending
										if v1400 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3951), int32(_a_F_getObjectDescription_2))
											mBase = m.M
											v1405 = m.ExcPending
											if v1405 != 0 {
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
								v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+16))
								v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1406)+22)))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+1152)) = v1406 + v1407 + int32(4)
								F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_66), v11+int32(1152))
								mBase = m.M
								v1418 = m.ExcPending
								if v1418 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v1386)
									mBase = m.M
									v1420 = m.ExcPending
									if v1420 != 0 {
										return int32(0)
									} else {
										v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
										v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
										m.G0 = v11 + int32(1440)
										if v1847 != 0 {
											v1853 = v1848
										} else {
											v1853 = int32(0)
										}
										return v1853
									}
								}
							}
						}
					default:
						if v19 != int32(3381) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v1646 = m.ExcPending
							if v1646 != 0 {
								return int32(0)
							} else {
								v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1647
								F_errmsg_internal(m, int32(_a_F_getObjectDescription_4), v11)
								mBase = m.M
								v1651 = m.ExcPending
								if v1651 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(_a_F_getObjectDescription_5), int32(_a_F_getObjectDescription_2))
									mBase = m.M
									v1656 = m.ExcPending
									if v1656 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v904 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
							v905 = F_SearchSysCache1(m, int32(64), v904)
							mBase = m.M
							v906 = m.ExcPending
							if v906 != 0 {
								return int32(0)
							} else {
								if v905 == int32(0) {
									if l1 != 0 {
										v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
										v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
										m.G0 = v11 + int32(1440)
										if v1847 != 0 {
											v1853 = v1848
										} else {
											v1853 = int32(0)
										}
										return v1853
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v912 = m.ExcPending
										if v912 != 0 {
											return int32(0)
										} else {
											v913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+560)) = v913
											F_errmsg_internal(m, int32(_a_F_getObjectDescription_67), v11+int32(560))
											mBase = m.M
											v919 = m.ExcPending
											if v919 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3529), int32(_a_F_getObjectDescription_2))
												mBase = m.M
												v924 = m.ExcPending
												if v924 != 0 {
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
									v925 = *(*int32)(unsafe.Add(mBase, uint32(v905)+16))
									v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v925)+22)))
									v927 = v925 + v926
									v928 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v930 = F_StatisticsObjIsVisibleExt(m, v928, int32(0))
									mBase = m.M
									v931 = m.ExcPending
									if v931 != 0 {
										return int32(0)
									} else {
										if v930 != 0 {
											v936 = int32(0)
											v939 = F_quote_qualified_identifier(m, v936, v927+int32(8))
											mBase = m.M
											v940 = m.ExcPending
											if v940 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+576)) = v939
												F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_68), v11+int32(576))
												mBase = m.M
												v948 = m.ExcPending
												if v948 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v905)
													mBase = m.M
													v950 = m.ExcPending
													if v950 != 0 {
														return int32(0)
													} else {
														v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
														v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
														m.G0 = v11 + int32(1440)
														if v1847 != 0 {
															v1853 = v1848
														} else {
															v1853 = int32(0)
														}
														return v1853
													}
												}
											}
										} else {
											v933 = *(*int32)(unsafe.Add(mBase, uint32(v927)+72))
											v934 = F_get_namespace_name(m, v933)
											mBase = m.M
											v935 = m.ExcPending
											if v935 != 0 {
												return int32(0)
											} else {
												v936 = v934
												v939 = F_quote_qualified_identifier(m, v936, v927+int32(8))
												mBase = m.M
												v940 = m.ExcPending
												if v940 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+576)) = v939
													F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_68), v11+int32(576))
													mBase = m.M
													v948 = m.ExcPending
													if v948 != 0 {
														return int32(0)
													} else {
														F_ReleaseCatCache(m, v905)
														mBase = m.M
														v950 = m.ExcPending
														if v950 != 0 {
															return int32(0)
														} else {
															v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
															v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
															m.G0 = v11 + int32(1440)
															if v1847 != 0 {
																v1853 = v1848
															} else {
																v1853 = int32(0)
															}
															return v1853
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
				if v19 <= int32(_a_F_getObjectDescription_69) {
					switch v19 - int32(3576) {
					case 0:
						v1599 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						v1600 = F_SearchSysCache1(m, int32(70), v1599)
						mBase = m.M
						v1601 = m.ExcPending
						if v1601 != 0 {
							return int32(0)
						} else {
							if v1600 == int32(0) {
								if l1 != 0 {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1607 = m.ExcPending
									if v1607 != 0 {
										return int32(0)
									} else {
										v1608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+1312)) = v1608
										F_errmsg_internal(m, int32(_a_F_getObjectDescription_70), v11+int32(1312))
										mBase = m.M
										v1614 = m.ExcPending
										if v1614 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(_a_F_getObjectDescription_71), int32(_a_F_getObjectDescription_2))
											mBase = m.M
											v1619 = m.ExcPending
											if v1619 != 0 {
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
								v1620 = *(*int32)(unsafe.Add(mBase, uint32(v1600)+16))
								v1621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1620)+22)))
								v1622 = v1620 + v1621
								v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+4))
								v1624 = F_format_type_be(m, v1623)
								mBase = m.M
								v1625 = m.ExcPending
								if v1625 != 0 {
									return int32(0)
								} else {
									v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+8))
									v1628 = F_get_language_name(m, v1626, int32(0))
									mBase = m.M
									v1629 = m.ExcPending
									if v1629 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+1332)) = v1628
										*(*int32)(unsafe.Add(mBase, uint32(v11)+1328)) = v1624
										F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_72), v11+int32(1328))
										mBase = m.M
										v1638 = m.ExcPending
										if v1638 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v1600)
											mBase = m.M
											v1640 = m.ExcPending
											if v1640 != 0 {
												return int32(0)
											} else {
												v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
												v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
												m.G0 = v11 + int32(1440)
												if v1847 != 0 {
													v1853 = v1848
												} else {
													v1853 = int32(0)
												}
												return v1853
											}
										}
									}
								}
							}
						}
					case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v1646 = m.ExcPending
						if v1646 != 0 {
							return int32(0)
						} else {
							v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1647
							F_errmsg_internal(m, int32(_a_F_getObjectDescription_4), v11)
							mBase = m.M
							v1651 = m.ExcPending
							if v1651 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(_a_F_getObjectDescription_5), int32(_a_F_getObjectDescription_2))
								mBase = m.M
								v1656 = m.ExcPending
								if v1656 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 24:
						v1000 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						v1001 = F_SearchSysCache1(m, int32(76), v1000)
						mBase = m.M
						v1002 = m.ExcPending
						if v1002 != 0 {
							return int32(0)
						} else {
							if v1001 == int32(0) {
								if l1 != 0 {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1008 = m.ExcPending
									if v1008 != 0 {
										return int32(0)
									} else {
										v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+624)) = v1009
										F_errmsg_internal(m, int32(_a_F_getObjectDescription_73), v11+int32(624))
										mBase = m.M
										v1015 = m.ExcPending
										if v1015 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3591), int32(_a_F_getObjectDescription_2))
											mBase = m.M
											v1020 = m.ExcPending
											if v1020 != 0 {
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
								v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+16))
								v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1021)+22)))
								v1023 = v1021 + v1022
								v1024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v1025 = F_TSDictionaryIsVisible(m, v1024)
								mBase = m.M
								v1026 = m.ExcPending
								if v1026 != 0 {
									return int32(0)
								} else {
									if v1025 != 0 {
										v1031 = int32(0)
										v1034 = F_quote_qualified_identifier(m, v1031, v1023+int32(4))
										mBase = m.M
										v1035 = m.ExcPending
										if v1035 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+640)) = v1034
											F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_74), v11+int32(640))
											mBase = m.M
											v1043 = m.ExcPending
											if v1043 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v1001)
												mBase = m.M
												v1045 = m.ExcPending
												if v1045 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
												}
											}
										}
									} else {
										v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+68))
										v1029 = F_get_namespace_name(m, v1028)
										mBase = m.M
										v1030 = m.ExcPending
										if v1030 != 0 {
											return int32(0)
										} else {
											v1031 = v1029
											v1034 = F_quote_qualified_identifier(m, v1031, v1023+int32(4))
											mBase = m.M
											v1035 = m.ExcPending
											if v1035 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+640)) = v1034
												F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_74), v11+int32(640))
												mBase = m.M
												v1043 = m.ExcPending
												if v1043 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v1001)
													mBase = m.M
													v1045 = m.ExcPending
													if v1045 != 0 {
														return int32(0)
													} else {
														v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
														v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
														m.G0 = v11 + int32(1440)
														if v1847 != 0 {
															v1853 = v1848
														} else {
															v1853 = int32(0)
														}
														return v1853
													}
												}
											}
										}
									}
								}
							}
						}
					case 25:
						v952 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						v953 = F_SearchSysCache1(m, int32(78), v952)
						mBase = m.M
						v954 = m.ExcPending
						if v954 != 0 {
							return int32(0)
						} else {
							if v953 == int32(0) {
								if l1 != 0 {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v960 = m.ExcPending
									if v960 != 0 {
										return int32(0)
									} else {
										v961 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+592)) = v961
										F_errmsg_internal(m, int32(_a_F_getObjectDescription_75), v11+int32(592))
										mBase = m.M
										v967 = m.ExcPending
										if v967 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3561), int32(_a_F_getObjectDescription_2))
											mBase = m.M
											v972 = m.ExcPending
											if v972 != 0 {
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
								v973 = *(*int32)(unsafe.Add(mBase, uint32(v953)+16))
								v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v973)+22)))
								v975 = v973 + v974
								v976 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v978 = F_TSParserIsVisibleExt(m, v976, int32(0))
								mBase = m.M
								v979 = m.ExcPending
								if v979 != 0 {
									return int32(0)
								} else {
									if v978 != 0 {
										v984 = int32(0)
										v987 = F_quote_qualified_identifier(m, v984, v975+int32(4))
										mBase = m.M
										v988 = m.ExcPending
										if v988 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+608)) = v987
											F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_76), v11+int32(608))
											mBase = m.M
											v996 = m.ExcPending
											if v996 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v953)
												mBase = m.M
												v998 = m.ExcPending
												if v998 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
												}
											}
										}
									} else {
										v981 = *(*int32)(unsafe.Add(mBase, uint32(v975)+68))
										v982 = F_get_namespace_name(m, v981)
										mBase = m.M
										v983 = m.ExcPending
										if v983 != 0 {
											return int32(0)
										} else {
											v984 = v982
											v987 = F_quote_qualified_identifier(m, v984, v975+int32(4))
											mBase = m.M
											v988 = m.ExcPending
											if v988 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+608)) = v987
												F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_76), v11+int32(608))
												mBase = m.M
												v996 = m.ExcPending
												if v996 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v953)
													mBase = m.M
													v998 = m.ExcPending
													if v998 != 0 {
														return int32(0)
													} else {
														v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
														v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
														m.G0 = v11 + int32(1440)
														if v1847 != 0 {
															v1853 = v1848
														} else {
															v1853 = int32(0)
														}
														return v1853
													}
												}
											}
										}
									}
								}
							}
						}
					case 26:
						v1097 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						v1098 = F_SearchSysCache1(m, int32(74), v1097)
						mBase = m.M
						v1099 = m.ExcPending
						if v1099 != 0 {
							return int32(0)
						} else {
							if v1098 == int32(0) {
								if l1 != 0 {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1105 = m.ExcPending
									if v1105 != 0 {
										return int32(0)
									} else {
										v1106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+688)) = v1106
										F_errmsg_internal(m, int32(_a_F_getObjectDescription_77), v11+int32(688))
										mBase = m.M
										v1112 = m.ExcPending
										if v1112 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3653), int32(_a_F_getObjectDescription_2))
											mBase = m.M
											v1117 = m.ExcPending
											if v1117 != 0 {
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
								v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+16))
								v1119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1118)+22)))
								v1120 = v1118 + v1119
								v1121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v1122 = F_TSConfigIsVisible(m, v1121)
								mBase = m.M
								v1123 = m.ExcPending
								if v1123 != 0 {
									return int32(0)
								} else {
									if v1122 != 0 {
										v1128 = int32(0)
										v1131 = F_quote_qualified_identifier(m, v1128, v1120+int32(4))
										mBase = m.M
										v1132 = m.ExcPending
										if v1132 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+704)) = v1131
											F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_78), v11+int32(704))
											mBase = m.M
											v1140 = m.ExcPending
											if v1140 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v1098)
												mBase = m.M
												v1142 = m.ExcPending
												if v1142 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
												}
											}
										}
									} else {
										v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+68))
										v1126 = F_get_namespace_name(m, v1125)
										mBase = m.M
										v1127 = m.ExcPending
										if v1127 != 0 {
											return int32(0)
										} else {
											v1128 = v1126
											v1131 = F_quote_qualified_identifier(m, v1128, v1120+int32(4))
											mBase = m.M
											v1132 = m.ExcPending
											if v1132 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+704)) = v1131
												F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_78), v11+int32(704))
												mBase = m.M
												v1140 = m.ExcPending
												if v1140 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v1098)
													mBase = m.M
													v1142 = m.ExcPending
													if v1142 != 0 {
														return int32(0)
													} else {
														v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
														v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
														m.G0 = v11 + int32(1440)
														if v1847 != 0 {
															v1853 = v1848
														} else {
															v1853 = int32(0)
														}
														return v1853
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
							v1646 = m.ExcPending
							if v1646 != 0 {
								return int32(0)
							} else {
								v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1647
								F_errmsg_internal(m, int32(_a_F_getObjectDescription_4), v11)
								mBase = m.M
								v1651 = m.ExcPending
								if v1651 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(_a_F_getObjectDescription_5), int32(_a_F_getObjectDescription_2))
									mBase = m.M
									v1656 = m.ExcPending
									if v1656 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v1049 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
							v1050 = F_SearchSysCache1(m, int32(80), v1049)
							mBase = m.M
							v1051 = m.ExcPending
							if v1051 != 0 {
								return int32(0)
							} else {
								if v1050 == int32(0) {
									if l1 != 0 {
										v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
										v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
										m.G0 = v11 + int32(1440)
										if v1847 != 0 {
											v1853 = v1848
										} else {
											v1853 = int32(0)
										}
										return v1853
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v1057 = m.ExcPending
										if v1057 != 0 {
											return int32(0)
										} else {
											v1058 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+656)) = v1058
											F_errmsg_internal(m, int32(_a_F_getObjectDescription_79), v11+int32(656))
											mBase = m.M
											v1064 = m.ExcPending
											if v1064 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3622), int32(_a_F_getObjectDescription_2))
												mBase = m.M
												v1069 = m.ExcPending
												if v1069 != 0 {
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
									v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+16))
									v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1070)+22)))
									v1072 = v1070 + v1071
									v1073 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v1075 = F_TSTemplateIsVisibleExt(m, v1073, int32(0))
									mBase = m.M
									v1076 = m.ExcPending
									if v1076 != 0 {
										return int32(0)
									} else {
										if v1075 != 0 {
											v1081 = int32(0)
											v1084 = F_quote_qualified_identifier(m, v1081, v1072+int32(4))
											mBase = m.M
											v1085 = m.ExcPending
											if v1085 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+672)) = v1084
												F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_80), v11+int32(672))
												mBase = m.M
												v1093 = m.ExcPending
												if v1093 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v1050)
													mBase = m.M
													v1095 = m.ExcPending
													if v1095 != 0 {
														return int32(0)
													} else {
														v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
														v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
														m.G0 = v11 + int32(1440)
														if v1847 != 0 {
															v1853 = v1848
														} else {
															v1853 = int32(0)
														}
														return v1853
													}
												}
											}
										} else {
											v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+68))
											v1079 = F_get_namespace_name(m, v1078)
											mBase = m.M
											v1080 = m.ExcPending
											if v1080 != 0 {
												return int32(0)
											} else {
												v1081 = v1079
												v1084 = F_quote_qualified_identifier(m, v1081, v1072+int32(4))
												mBase = m.M
												v1085 = m.ExcPending
												if v1085 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+672)) = v1084
													F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_80), v11+int32(672))
													mBase = m.M
													v1093 = m.ExcPending
													if v1093 != 0 {
														return int32(0)
													} else {
														F_ReleaseCatCache(m, v1050)
														mBase = m.M
														v1095 = m.ExcPending
														if v1095 != 0 {
															return int32(0)
														} else {
															v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
															v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
															m.G0 = v11 + int32(1440)
															if v1847 != 0 {
																v1853 = v1848
															} else {
																v1853 = int32(0)
															}
															return v1853
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
					switch v19 - int32(_a_F_getObjectDescription_81) {
					case 0:
						v1585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1586 = F_get_subscription_name(m, v1585, l1)
						mBase = m.M
						v1587 = m.ExcPending
						if v1587 != 0 {
							return int32(0)
						} else {
							if v1586 == int32(0) {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+1296)) = v1586
								F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_82), v11+int32(1296))
								mBase = m.M
								v1597 = m.ExcPending
								if v1597 != 0 {
									return int32(0)
								} else {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								}
							}
						}
					case 1, 2, 3, 5:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v1646 = m.ExcPending
						if v1646 != 0 {
							return int32(0)
						} else {
							v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1647
							F_errmsg_internal(m, int32(_a_F_getObjectDescription_4), v11)
							mBase = m.M
							v1651 = m.ExcPending
							if v1651 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(_a_F_getObjectDescription_5), int32(_a_F_getObjectDescription_2))
								mBase = m.M
								v1656 = m.ExcPending
								if v1656 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 4:
						v1494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v1495 = F_get_publication_name(m, v1494, l1)
						mBase = m.M
						v1496 = m.ExcPending
						if v1496 != 0 {
							return int32(0)
						} else {
							if v1495 == int32(0) {
								v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
								v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
								m.G0 = v11 + int32(1440)
								if v1847 != 0 {
									v1853 = v1848
								} else {
									v1853 = int32(0)
								}
								return v1853
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+1232)) = v1495
								F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_83), v11+int32(1232))
								mBase = m.M
								v1506 = m.ExcPending
								if v1506 != 0 {
									return int32(0)
								} else {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								}
							}
						}
					case 6:
						v1531 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
						v1532 = F_SearchSysCache1(m, int32(52), v1531)
						mBase = m.M
						v1533 = m.ExcPending
						if v1533 != 0 {
							return int32(0)
						} else {
							if v1532 == int32(0) {
								if l1 != 0 {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1539 = m.ExcPending
									if v1539 != 0 {
										return int32(0)
									} else {
										v1540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+1264)) = v1540
										F_errmsg_internal(m, int32(_a_F_getObjectDescription_84), v11+int32(1264))
										mBase = m.M
										v1546 = m.ExcPending
										if v1546 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(4068), int32(_a_F_getObjectDescription_2))
											mBase = m.M
											v1551 = m.ExcPending
											if v1551 != 0 {
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
								v1552 = *(*int32)(unsafe.Add(mBase, uint32(v1532)+16))
								v1553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1552)+22)))
								v1554 = v1552 + v1553
								v1555 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+4))
								v1557 = F_get_publication_name(m, v1555, int32(0))
								mBase = m.M
								v1558 = m.ExcPending
								if v1558 != 0 {
									return int32(0)
								} else {
									v1560 = v11 + int32(1360)
									F_initStringInfo(m, v1560)
									mBase = m.M
									v1562 = m.ExcPending
									if v1562 != 0 {
										return int32(0)
									} else {
										v1563 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+8))
										F_getRelationDescription(m, v1560, v1563, int32(0))
										mBase = m.M
										v1566 = m.ExcPending
										if v1566 != 0 {
											return int32(0)
										} else {
											v1567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1554)+12)))
											v1568 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+1280)) = v1568
											*(*int32)(unsafe.Add(mBase, uint32(v11)+1284)) = v1557
											if v1567 != 0 {
												v1575 = int32(_a_F_getObjectDescription_85)
											} else {
												v1575 = int32(_a_F_getObjectDescription_86)
											}
											F_appendStringInfo(m, v11+int32(1424), v1575, v11+int32(1280))
											mBase = m.M
											v1579 = m.ExcPending
											if v1579 != 0 {
												return int32(0)
											} else {
												v1580 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
												F_pfree(m, v1580)
												mBase = m.M
												v1582 = m.ExcPending
												if v1582 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v1532)
													mBase = m.M
													v1584 = m.ExcPending
													if v1584 != 0 {
														return int32(0)
													} else {
														v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
														v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
														m.G0 = v11 + int32(1440)
														if v1847 != 0 {
															v1853 = v1848
														} else {
															v1853 = int32(0)
														}
														return v1853
													}
												}
											}
										}
									}
								}
							}
						}
					default:
						switch v19 - int32(_a_F_getObjectDescription_87) {
						case 0:
							v1511 = F_getPublicationSchemaInfo(m, l0, l1, v11+int32(1360), v11+int32(1344))
							mBase = m.M
							v1512 = m.ExcPending
							if v1512 != 0 {
								return int32(0)
							} else {
								if v1511 == int32(0) {
									v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
									v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
									m.G0 = v11 + int32(1440)
									if v1847 != 0 {
										v1853 = v1848
									} else {
										v1853 = int32(0)
									}
									return v1853
								} else {
									v1515 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1344))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+1248)) = v1515
									v1517 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1360))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+1252)) = v1517
									F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_88), v11+int32(1248))
									mBase = m.M
									v1525 = m.ExcPending
									if v1525 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v1517)
										mBase = m.M
										v1527 = m.ExcPending
										if v1527 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v1515)
											mBase = m.M
											v1529 = m.ExcPending
											if v1529 != 0 {
												return int32(0)
											} else {
												v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
												v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
												m.G0 = v11 + int32(1440)
												if v1847 != 0 {
													v1853 = v1848
												} else {
													v1853 = int32(0)
												}
												return v1853
											}
										}
									}
								}
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v1646 = m.ExcPending
							if v1646 != 0 {
								return int32(0)
							} else {
								v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v1647
								F_errmsg_internal(m, int32(_a_F_getObjectDescription_4), v11)
								mBase = m.M
								v1651 = m.ExcPending
								if v1651 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(_a_F_getObjectDescription_5), int32(_a_F_getObjectDescription_2))
									mBase = m.M
									v1656 = m.ExcPending
									if v1656 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 6:
							v1422 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
							v1423 = F_SearchSysCache1(m, int32(44), v1422)
							mBase = m.M
							v1424 = m.ExcPending
							if v1424 != 0 {
								return int32(0)
							} else {
								if v1423 == int32(0) {
									if l1 != 0 {
										v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
										v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
										m.G0 = v11 + int32(1440)
										if v1847 != 0 {
											v1853 = v1848
										} else {
											v1853 = int32(0)
										}
										return v1853
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v1430 = m.ExcPending
										if v1430 != 0 {
											return int32(0)
										} else {
											v1431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+1168)) = v1431
											F_errmsg_internal(m, int32(_a_F_getObjectDescription_89), v11+int32(1168))
											mBase = m.M
											v1437 = m.ExcPending
											if v1437 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_getObjectDescription_1), int32(3972), int32(_a_F_getObjectDescription_2))
												mBase = m.M
												v1442 = m.ExcPending
												if v1442 != 0 {
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
									v1445 = F_SysCacheGetAttrNotNull(m, int32(44), v1423, int32(2))
									mBase = m.M
									v1446 = m.ExcPending
									if v1446 != 0 {
										return int32(0)
									} else {
										v1448 = F_text_to_cstring(m, base.I32_wrap_i64(v1445))
										mBase = m.M
										v1449 = m.ExcPending
										if v1449 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+1184)) = v1448
											F_appendStringInfo(m, v11+int32(1424), int32(_a_F_getObjectDescription_90), v11+int32(1184))
											mBase = m.M
											v1457 = m.ExcPending
											if v1457 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v1423)
												mBase = m.M
												v1459 = m.ExcPending
												if v1459 != 0 {
													return int32(0)
												} else {
													v1847 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1428))
													v1848 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1424))
													m.G0 = v11 + int32(1440)
													if v1847 != 0 {
														v1853 = v1848
													} else {
														v1853 = int32(0)
													}
													return v1853
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
