package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_apply_dispatch(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
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
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v166 int64
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int64
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
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
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v494 int32
	_ = v494
	var v496 int64
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v507 int64
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
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
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
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
	var v756 int64
	_ = v756
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v857 int64
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v868 int64
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
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
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1026 int64
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1037 int64
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1138 int32
	_ = v1138
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
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
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
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1271 int32
	_ = v1271
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
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
	var v1438 int32
	_ = v1438
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1618 int32
	_ = v1618
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1639 int32
	_ = v1639
	var v1640 int64
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1648 int32
	_ = v1648
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1683 int32
	_ = v1683
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1771 int32
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1777 int64
	_ = v1777
	var v1783 int32
	_ = v1783
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1806 int64
	_ = v1806
	var v1812 int32
	_ = v1812
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1865 int32
	_ = v1865
	var v1879 int32
	_ = v1879
	var v1883 int32
	_ = v1883
	var v1884 int64
	_ = v1884
	var v1895 int32
	_ = v1895
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1923 int32
	_ = v1923
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1938 int32
	_ = v1938
	var v1966 int32
	_ = v1966
	var v1970 int32
	_ = v1970
	var v1975 int32
	_ = v1975
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2016 int32
	_ = v2016
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2059 int32
	_ = v2059
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2071 int32
	_ = v2071
	var v2074 int32
	_ = v2074
	var v2077 int32
	_ = v2077
	var v2080 int32
	_ = v2080
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2092 int32
	_ = v2092
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2102 int32
	_ = v2102
	var v2113 int32
	_ = v2113
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2123 int32
	_ = v2123
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2143 int32
	_ = v2143
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2152 int32
	_ = v2152
	var v2154 int32
	_ = v2154
	var v2156 int32
	_ = v2156
	var v2161 int32
	_ = v2161
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2178 int32
	_ = v2178
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2200 int32
	_ = v2200
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2207 int32
	_ = v2207
	var v2209 int32
	_ = v2209
	var v2211 int32
	_ = v2211
	var v2216 int32
	_ = v2216
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2235 int32
	_ = v2235
	var v2240 int32
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2246 int32
	_ = v2246
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2276 int64
	_ = v2276
	var v2278 int64
	_ = v2278
	var v2281 int32
	_ = v2281
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2295 int64
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2298 int64
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2301 int64
	_ = v2301
	var v2306 int32
	_ = v2306
	var v2309 int64
	_ = v2309
	var v2311 int32
	_ = v2311
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2317 int32
	_ = v2317
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2332 int32
	_ = v2332
	var v2334 int32
	_ = v2334
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2356 int32
	_ = v2356
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2374 int32
	_ = v2374
	var v2376 int32
	_ = v2376
	var v2380 int64
	_ = v2380
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2390 int32
	_ = v2390
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2399 int32
	_ = v2399
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2407 int32
	_ = v2407
	var v2409 int64
	_ = v2409
	var v2427 int64
	_ = v2427
	var v2432 int64
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2437 int32
	_ = v2437
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2462 int32
	_ = v2462
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2467 int64
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2474 int32
	_ = v2474
	var v2476 int64
	_ = v2476
	var v2478 int32
	_ = v2478
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2504 int32
	_ = v2504
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2519 int32
	_ = v2519
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2534 int32
	_ = v2534
	var v2538 int32
	_ = v2538
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2544 int32
	_ = v2544
	var v2545 int32
	_ = v2545
	var v2546 int64
	_ = v2546
	var v2548 int32
	_ = v2548
	var v2551 int32
	_ = v2551
	var v2558 int32
	_ = v2558
	var v2567 int64
	_ = v2567
	var v2570 int32
	_ = v2570
	var v2572 int64
	_ = v2572
	var v2574 int64
	_ = v2574
	var v2578 int32
	_ = v2578
	var v2579 int64
	_ = v2579
	var v2601 int32
	_ = v2601
	var v2603 int32
	_ = v2603
	var v2609 int32
	_ = v2609
	var v2614 int32
	_ = v2614
	var v2618 int32
	_ = v2618
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2625 int32
	_ = v2625
	var v2627 int32
	_ = v2627
	var v2631 int32
	_ = v2631
	var v2637 int32
	_ = v2637
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2651 int32
	_ = v2651
	var v2653 int32
	_ = v2653
	var v2655 int64
	_ = v2655
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2713 int32
	_ = v2713
	var v2718 int32
	_ = v2718
	var v2721 int32
	_ = v2721
	var v2725 int32
	_ = v2725
	var v2730 int32
	_ = v2730
	var v2732 int32
	_ = v2732
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2739 int64
	_ = v2739
	var v2742 int64
	_ = v2742
	var v2746 int32
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2754 int32
	_ = v2754
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2765 int32
	_ = v2765
	var v2767 int32
	_ = v2767
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2774 int32
	_ = v2774
	var v2775 int32
	_ = v2775
	var v2777 int32
	_ = v2777
	var v2779 int32
	_ = v2779
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2795 int32
	_ = v2795
	var v2796 int32
	_ = v2796
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2806 int32
	_ = v2806
	var v2811 int32
	_ = v2811
	var v2814 int32
	_ = v2814
	var v2815 int32
	_ = v2815
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2837 int32
	_ = v2837
	var v2840 int32
	_ = v2840
	var v2844 int32
	_ = v2844
	var v2849 int32
	_ = v2849
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2860 int32
	_ = v2860
	var v2863 int32
	_ = v2863
	var v2889 int32
	_ = v2889
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2899 int32
	_ = v2899
	var v2904 int32
	_ = v2904
	var v2908 int32
	_ = v2908
	var v2915 int32
	_ = v2915
	var v2920 int32
	_ = v2920
	var v2922 int32
	_ = v2922
	var v2923 int32
	_ = v2923
	var v2924 int32
	_ = v2924
	var v2929 int32
	_ = v2929
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2932 int32
	_ = v2932
	var v2933 int32
	_ = v2933
	var v2938 int64
	_ = v2938
	var v2940 int64
	_ = v2940
	var v2943 int32
	_ = v2943
	var v2945 int32
	_ = v2945
	var v2946 int32
	_ = v2946
	var v2948 int32
	_ = v2948
	var v2951 int32
	_ = v2951
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2954 int32
	_ = v2954
	var v2956 int32
	_ = v2956
	var v2960 int32
	_ = v2960
	var v2964 int32
	_ = v2964
	var v2969 int32
	_ = v2969
	var v2970 int64
	_ = v2970
	var v2971 int32
	_ = v2971
	var v2973 int64
	_ = v2973
	var v2974 int32
	_ = v2974
	var v2976 int64
	_ = v2976
	var v2977 int32
	_ = v2977
	var v2985 int64
	_ = v2985
	var v2988 int32
	_ = v2988
	var v2989 int32
	_ = v2989
	var v2992 int32
	_ = v2992
	var v2995 int32
	_ = v2995
	var v2996 int32
	_ = v2996
	var v2997 int32
	_ = v2997
	var v2998 int32
	_ = v2998
	var v2999 int32
	_ = v2999
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3004 int64
	_ = v3004
	var v3006 int32
	_ = v3006
	var v3008 int32
	_ = v3008
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3012 int64
	_ = v3012
	var v3014 int32
	_ = v3014
	var v3018 int32
	_ = v3018
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3023 int32
	_ = v3023
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
	var v3046 int32
	_ = v3046
	var v3047 int32
	_ = v3047
	var v3049 int32
	_ = v3049
	var v3050 int64
	_ = v3050
	var v3052 int32
	_ = v3052
	var v3054 int32
	_ = v3054
	var v3056 int32
	_ = v3056
	var v3063 int32
	_ = v3063
	var v3065 int32
	_ = v3065
	var v3067 int64
	_ = v3067
	var v3071 int32
	_ = v3071
	var v3073 int32
	_ = v3073
	var v3079 int32
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3086 int32
	_ = v3086
	var v3091 int32
	_ = v3091
	var v3094 int64
	_ = v3094
	var v3096 int32
	_ = v3096
	var v3098 int32
	_ = v3098
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3120 int32
	_ = v3120
	var v3124 int32
	_ = v3124
	var v3125 int64
	_ = v3125
	var v3126 int32
	_ = v3126
	var v3130 int64
	_ = v3130
	var v3131 int32
	_ = v3131
	var v3135 int64
	_ = v3135
	var v3136 int32
	_ = v3136
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3143 int32
	_ = v3143
	var v3144 int32
	_ = v3144
	var v3145 int32
	_ = v3145
	var v3152 int32
	_ = v3152
	var v3165 int32
	_ = v3165
	var v3166 int32
	_ = v3166
	var v3167 int32
	_ = v3167
	var v3169 int32
	_ = v3169
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3176 int32
	_ = v3176
	var v3177 int32
	_ = v3177
	var v3178 int32
	_ = v3178
	var v3180 int32
	_ = v3180
	var v3186 int32
	_ = v3186
	var v3187 int32
	_ = v3187
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3192 int32
	_ = v3192
	var v3198 int32
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3200 int32
	_ = v3200
	var v3203 int32
	_ = v3203
	var v3206 int32
	_ = v3206
	var v3211 int32
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3214 int32
	_ = v3214
	var v3216 int32
	_ = v3216
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3232 int32
	_ = v3232
	var v3233 int32
	_ = v3233
	var v3234 int32
	_ = v3234
	var v3236 int32
	_ = v3236
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3243 int32
	_ = v3243
	var v3245 int32
	_ = v3245
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3251 int32
	_ = v3251
	var v3258 int32
	_ = v3258
	var v3264 int32
	_ = v3264
	var v3268 int32
	_ = v3268
	var v3273 int32
	_ = v3273
	var v3277 int32
	_ = v3277
	var v3281 int32
	_ = v3281
	var v3286 int32
	_ = v3286
	var v3288 int32
	_ = v3288
	var v3291 int64
	_ = v3291
	var v3296 int32
	_ = v3296
	var v3297 int64
	_ = v3297
	var v3305 int32
	_ = v3305
	var v3306 int32
	_ = v3306
	var v3310 int64
	_ = v3310
	var v3313 int64
	_ = v3313
	var v3319 int32
	_ = v3319
	var v3324 int32
	_ = v3324
	var v3327 int32
	_ = v3327
	var v3336 int32
	_ = v3336
	var v3337 int64
	_ = v3337
	var v3339 int64
	_ = v3339
	var v3342 int32
	_ = v3342
	var v3346 int64
	_ = v3346
	var v3349 int32
	_ = v3349
	var v3350 int32
	_ = v3350
	var v3356 int32
	_ = v3356
	var v3358 int32
	_ = v3358
	var v3359 int32
	_ = v3359
	var v3360 int32
	_ = v3360
	var v3362 int32
	_ = v3362
	var v3365 int32
	_ = v3365
	var v3368 int32
	_ = v3368
	var v3369 int32
	_ = v3369
	var v3370 int32
	_ = v3370
	var v3374 int32
	_ = v3374
	var v3376 int32
	_ = v3376
	var v3377 int32
	_ = v3377
	var v3383 int32
	_ = v3383
	var v3385 int32
	_ = v3385
	var v3387 int64
	_ = v3387
	var v3390 int64
	_ = v3390
	var v3394 int32
	_ = v3394
	var v3395 int32
	_ = v3395
	var v3397 int32
	_ = v3397
	var v3399 int32
	_ = v3399
	var v3401 int32
	_ = v3401
	var v3403 int32
	_ = v3403
	var v3404 int32
	_ = v3404
	var v3405 int64
	_ = v3405
	var v3407 int32
	_ = v3407
	var v3408 int32
	_ = v3408
	var v3411 int32
	_ = v3411
	var v3416 int32
	_ = v3416
	var v3419 int32
	_ = v3419
	var v3420 int32
	_ = v3420
	var v3425 int32
	_ = v3425
	var v3427 int32
	_ = v3427
	var v3429 int32
	_ = v3429
	var v3432 int32
	_ = v3432
	var v3434 int32
	_ = v3434
	var v3441 int32
	_ = v3441
	var v3443 int64
	_ = v3443
	var v3446 int64
	_ = v3446
	var v3448 int32
	_ = v3448
	var v3451 int32
	_ = v3451
	var v3453 int64
	_ = v3453
	var v3458 int32
	_ = v3458
	var v3459 int32
	_ = v3459
	var v3461 int64
	_ = v3461
	var v3464 int64
	_ = v3464
	var v3470 int32
	_ = v3470
	var v3475 int32
	_ = v3475
	var v3481 int64
	_ = v3481
	var v3483 int32
	_ = v3483
	var v3485 int32
	_ = v3485
	var v3503 int32
	_ = v3503
	var v3504 int32
	_ = v3504
	var v3506 int32
	_ = v3506
	var v3508 int32
	_ = v3508
	var v3509 int32
	_ = v3509
	var v3511 int32
	_ = v3511
	var v3514 int64
	_ = v3514
	var v3515 int32
	_ = v3515
	var v3519 int64
	_ = v3519
	var v3520 int32
	_ = v3520
	var v3524 int64
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3528 int32
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3532 int32
	_ = v3532
	var v3533 int32
	_ = v3533
	var v3534 int32
	_ = v3534
	var v3541 int32
	_ = v3541
	var v3554 int32
	_ = v3554
	var v3555 int32
	_ = v3555
	var v3556 int32
	_ = v3556
	var v3558 int32
	_ = v3558
	var v3562 int32
	_ = v3562
	var v3563 int32
	_ = v3563
	var v3565 int32
	_ = v3565
	var v3566 int32
	_ = v3566
	var v3567 int32
	_ = v3567
	var v3569 int32
	_ = v3569
	var v3575 int32
	_ = v3575
	var v3576 int32
	_ = v3576
	var v3577 int32
	_ = v3577
	var v3578 int32
	_ = v3578
	var v3581 int32
	_ = v3581
	var v3587 int32
	_ = v3587
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3592 int32
	_ = v3592
	var v3595 int32
	_ = v3595
	var v3600 int32
	_ = v3600
	var v3601 int32
	_ = v3601
	var v3603 int32
	_ = v3603
	var v3605 int32
	_ = v3605
	var v3609 int32
	_ = v3609
	var v3610 int32
	_ = v3610
	var v3611 int32
	_ = v3611
	var v3616 int32
	_ = v3616
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3621 int32
	_ = v3621
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3625 int32
	_ = v3625
	var v3629 int32
	_ = v3629
	var v3630 int32
	_ = v3630
	var v3632 int32
	_ = v3632
	var v3634 int32
	_ = v3634
	var v3636 int32
	_ = v3636
	var v3637 int32
	_ = v3637
	var v3640 int32
	_ = v3640
	var v3647 int32
	_ = v3647
	var v3656 int32
	_ = v3656
	var v3660 int32
	_ = v3660
	var v3665 int32
	_ = v3665
	var v3669 int32
	_ = v3669
	var v3673 int32
	_ = v3673
	var v3678 int32
	_ = v3678
	var v3682 int32
	_ = v3682
	var v3686 int32
	_ = v3686
	var v3691 int32
	_ = v3691
	var v3693 int32
	_ = v3693
	var v3696 int64
	_ = v3696
	var v3699 int32
	_ = v3699
	var v3700 int32
	_ = v3700
	var v3704 int32
	_ = v3704
	var v3706 int32
	_ = v3706
	var v3710 int64
	_ = v3710
	var v3713 int32
	_ = v3713
	var v3714 int32
	_ = v3714
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
	var v3729 int32
	_ = v3729
	var v3732 int64
	_ = v3732
	var v3735 int64
	_ = v3735
	var v3741 int32
	_ = v3741
	var v3743 int32
	_ = v3743
	var v3745 int32
	_ = v3745
	var v3747 int32
	_ = v3747
	var v3749 int32
	_ = v3749
	var v3750 int32
	_ = v3750
	var v3752 int64
	_ = v3752
	var v3753 int64
	_ = v3753
	var v3755 int32
	_ = v3755
	var v3756 int32
	_ = v3756
	var v3759 int32
	_ = v3759
	var v3764 int32
	_ = v3764
	var v3767 int32
	_ = v3767
	var v3768 int32
	_ = v3768
	var v3772 int32
	_ = v3772
	var v3774 int32
	_ = v3774
	var v3776 int32
	_ = v3776
	var v3779 int32
	_ = v3779
	var v3781 int32
	_ = v3781
	var v3788 int32
	_ = v3788
	var v3790 int64
	_ = v3790
	var v3793 int64
	_ = v3793
	var v3795 int32
	_ = v3795
	var v3798 int32
	_ = v3798
	var v3799 int64
	_ = v3799
	var v3801 int32
	_ = v3801
	var v3803 int32
	_ = v3803
	var v3821 int32
	_ = v3821
	var v3822 int32
	_ = v3822
	var v3824 int32
	_ = v3824
	var v3826 int32
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3829 int32
	_ = v3829
	var v3832 int64
	_ = v3832
	var v3833 int32
	_ = v3833
	var v3837 int64
	_ = v3837
	var v3838 int32
	_ = v3838
	var v3842 int64
	_ = v3842
	var v3843 int32
	_ = v3843
	var v3845 int64
	_ = v3845
	var v3846 int32
	_ = v3846
	var v3849 int32
	_ = v3849
	var v3850 int32
	_ = v3850
	var v3853 int32
	_ = v3853
	var v3854 int32
	_ = v3854
	var v3855 int32
	_ = v3855
	var v3862 int32
	_ = v3862
	var v3875 int32
	_ = v3875
	var v3876 int32
	_ = v3876
	var v3877 int32
	_ = v3877
	var v3879 int32
	_ = v3879
	var v3883 int32
	_ = v3883
	var v3884 int32
	_ = v3884
	var v3886 int32
	_ = v3886
	var v3887 int32
	_ = v3887
	var v3888 int32
	_ = v3888
	var v3890 int32
	_ = v3890
	var v3896 int32
	_ = v3896
	var v3897 int32
	_ = v3897
	var v3898 int32
	_ = v3898
	var v3899 int32
	_ = v3899
	var v3902 int32
	_ = v3902
	var v3908 int32
	_ = v3908
	var v3909 int32
	_ = v3909
	var v3910 int32
	_ = v3910
	var v3913 int32
	_ = v3913
	var v3916 int32
	_ = v3916
	var v3921 int32
	_ = v3921
	var v3922 int32
	_ = v3922
	var v3924 int32
	_ = v3924
	var v3926 int32
	_ = v3926
	var v3930 int32
	_ = v3930
	var v3931 int32
	_ = v3931
	var v3932 int32
	_ = v3932
	var v3937 int32
	_ = v3937
	var v3938 int32
	_ = v3938
	var v3939 int32
	_ = v3939
	var v3942 int32
	_ = v3942
	var v3943 int32
	_ = v3943
	var v3944 int32
	_ = v3944
	var v3946 int32
	_ = v3946
	var v3950 int32
	_ = v3950
	var v3951 int32
	_ = v3951
	var v3953 int32
	_ = v3953
	var v3955 int32
	_ = v3955
	var v3957 int32
	_ = v3957
	var v3958 int32
	_ = v3958
	var v3961 int32
	_ = v3961
	var v3968 int32
	_ = v3968
	var v3977 int32
	_ = v3977
	var v3981 int32
	_ = v3981
	var v3986 int32
	_ = v3986
	var v3990 int32
	_ = v3990
	var v3994 int32
	_ = v3994
	var v3999 int32
	_ = v3999
	var v4003 int32
	_ = v4003
	var v4007 int32
	_ = v4007
	var v4012 int32
	_ = v4012
	var v4014 int32
	_ = v4014
	var v4017 int64
	_ = v4017
	var v4020 int32
	_ = v4020
	var v4021 int32
	_ = v4021
	var v4025 int32
	_ = v4025
	var v4027 int32
	_ = v4027
	var v4028 int64
	_ = v4028
	var v4029 int64
	_ = v4029
	var v4030 int32
	_ = v4030
	var v4031 int32
	_ = v4031
	var v4033 int32
	_ = v4033
	var v4036 int32
	_ = v4036
	var v4040 int32
	_ = v4040
	var v4041 int32
	_ = v4041
	var v4043 int32
	_ = v4043
	var v4044 int32
	_ = v4044
	var v4047 int32
	_ = v4047
	var v4048 int32
	_ = v4048
	var v4069 int32
	_ = v4069
	var v4070 int32
	_ = v4070
	var v4074 int32
	_ = v4074
	var v4077 int32
	_ = v4077
	var v4078 int32
	_ = v4078
	var v4082 int32
	_ = v4082
	var v4083 int32
	_ = v4083
	var v4086 int32
	_ = v4086
	var v4087 int32
	_ = v4087
	var v4090 int32
	_ = v4090
	var v4097 int32
	_ = v4097
	var v4098 int32
	_ = v4098
	var v4100 int32
	_ = v4100
	var v4103 int32
	_ = v4103
	var v4105 int32
	_ = v4105
	var v4106 int32
	_ = v4106
	var v4107 int64
	_ = v4107
	var v4112 int32
	_ = v4112
	var v4113 int32
	_ = v4113
	var v4114 int32
	_ = v4114
	var v4115 int64
	_ = v4115
	var v4117 int64
	_ = v4117
	var v4120 int32
	_ = v4120
	var v4122 int32
	_ = v4122
	var v4123 int32
	_ = v4123
	var v4125 int32
	_ = v4125
	var v4126 int32
	_ = v4126
	var v4130 int32
	_ = v4130
	var v4133 int32
	_ = v4133
	var v4152 int32
	_ = v4152
	var v4156 int32
	_ = v4156
	var v4161 int64
	_ = v4161
	var v4164 int64
	_ = v4164
	var v4167 int32
	_ = v4167
	var v4171 int64
	_ = v4171
	var v4174 int32
	_ = v4174
	var v4175 int32
	_ = v4175
	var v4181 int32
	_ = v4181
	var v4183 int32
	_ = v4183
	var v4184 int32
	_ = v4184
	var v4185 int32
	_ = v4185
	var v4187 int32
	_ = v4187
	var v4190 int32
	_ = v4190
	var v4196 int32
	_ = v4196
	var v4198 int32
	_ = v4198
	var v4200 int32
	_ = v4200
	var v4202 int32
	_ = v4202
	var v4203 int64
	_ = v4203
	var v4205 int32
	_ = v4205
	var v4207 int32
	_ = v4207
	var v4208 int32
	_ = v4208
	var v4209 int64
	_ = v4209
	var v4211 int32
	_ = v4211
	var v4212 int32
	_ = v4212
	var v4215 int32
	_ = v4215
	var v4220 int32
	_ = v4220
	var v4223 int32
	_ = v4223
	var v4224 int32
	_ = v4224
	var v4229 int32
	_ = v4229
	var v4231 int32
	_ = v4231
	var v4233 int32
	_ = v4233
	var v4236 int32
	_ = v4236
	var v4238 int32
	_ = v4238
	var v4245 int32
	_ = v4245
	var v4247 int64
	_ = v4247
	var v4250 int64
	_ = v4250
	var v4252 int32
	_ = v4252
	var v4255 int32
	_ = v4255
	var v4257 int32
	_ = v4257
	var v4276 int64
	_ = v4276
	var v4278 int64
	_ = v4278
	var v4281 int32
	_ = v4281
	var v4283 int32
	_ = v4283
	var v4284 int32
	_ = v4284
	var v4287 int32
	_ = v4287
	var v4294 int32
	_ = v4294
	var v4296 int32
	_ = v4296
	var v4299 int64
	_ = v4299
	var v4302 int32
	_ = v4302
	var v4303 int32
	_ = v4303
	var v4306 int32
	_ = v4306
	var v4309 int32
	_ = v4309
	var v4310 int32
	_ = v4310
	var v4311 int32
	_ = v4311
	var v4312 int32
	_ = v4312
	var v4313 int32
	_ = v4313
	var v4314 int32
	_ = v4314
	var v4315 int32
	_ = v4315
	var v4318 int64
	_ = v4318
	var v4320 int32
	_ = v4320
	var v4322 int32
	_ = v4322
	var v4324 int32
	_ = v4324
	var v4325 int32
	_ = v4325
	var v4326 int32
	_ = v4326
	var v4327 int64
	_ = v4327
	var v4329 int32
	_ = v4329
	var v4333 int32
	_ = v4333
	var v4335 int32
	_ = v4335
	var v4336 int64
	_ = v4336
	var v4339 int32
	_ = v4339
	var v4341 int32
	_ = v4341
	var v4344 int32
	_ = v4344
	var v4345 int32
	_ = v4345
	var v4346 int32
	_ = v4346
	var v4348 int32
	_ = v4348
	var v4351 int32
	_ = v4351
	var v4352 int32
	_ = v4352
	var v4358 int32
	_ = v4358
	var v4363 int32
	_ = v4363
	var v4366 int32
	_ = v4366
	var v4367 int32
	_ = v4367
	var v4372 int32
	_ = v4372
	var v4373 int32
	_ = v4373
	var v4375 int32
	_ = v4375
	var v4376 int64
	_ = v4376
	var v4378 int32
	_ = v4378
	var v4380 int32
	_ = v4380
	var v4382 int32
	_ = v4382
	var v4387 int32
	_ = v4387
	var v4391 int64
	_ = v4391
	var v4394 int32
	_ = v4394
	var v4395 int32
	_ = v4395
	var v4401 int32
	_ = v4401
	var v4403 int32
	_ = v4403
	var v4404 int32
	_ = v4404
	var v4405 int32
	_ = v4405
	var v4407 int32
	_ = v4407
	var v4410 int32
	_ = v4410
	var v4415 int32
	_ = v4415
	var v4417 int32
	_ = v4417
	var v4419 int32
	_ = v4419
	var v4421 int32
	_ = v4421
	var v4423 int32
	_ = v4423
	var v4428 int32
	_ = v4428
	var v4430 int32
	_ = v4430
	var v4431 int32
	_ = v4431
	var v4433 int32
	_ = v4433
	var v4439 int32
	_ = v4439
	var v4440 int32
	_ = v4440
	var v4446 int32
	_ = v4446
	var v4451 int32
	_ = v4451
	var v4455 int32
	_ = v4455
	var v4456 int32
	_ = v4456
	var v4457 int64
	_ = v4457
	var v4459 int32
	_ = v4459
	var v4461 int64
	_ = v4461
	var v4466 int32
	_ = v4466
	var v4467 int32
	_ = v4467
	var v4469 int64
	_ = v4469
	var v4472 int64
	_ = v4472
	var v4478 int32
	_ = v4478
	var v4483 int32
	_ = v4483
	var v4489 int64
	_ = v4489
	var v4491 int32
	_ = v4491
	var v4493 int32
	_ = v4493
	var v4513 int32
	_ = v4513
	var v4516 int32
	_ = v4516
	var v4517 int64
	_ = v4517
	var v4520 int64
	_ = v4520
	var v4522 int64
	_ = v4522
	var v4523 int64
	_ = v4523
	var v4526 int64
	_ = v4526
	var v4532 int32
	_ = v4532
	var v4537 int32
	_ = v4537
	var v4541 int32
	_ = v4541
	var v4544 int32
	_ = v4544
	var v4548 int32
	_ = v4548
	var v4553 int32
	_ = v4553
	var v4557 int32
	_ = v4557
	var v4560 int32
	_ = v4560
	var v4564 int32
	_ = v4564
	var v4569 int32
	_ = v4569
	var v4573 int32
	_ = v4573
	var v4580 int32
	_ = v4580
	var v4585 int32
	_ = v4585
	var v4589 int32
	_ = v4589
	var v4592 int32
	_ = v4592
	var v4596 int32
	_ = v4596
	var v4601 int32
	_ = v4601
	var v4605 int32
	_ = v4605
	var v4612 int32
	_ = v4612
	var v4617 int32
	_ = v4617
	var v4621 int32
	_ = v4621
	var v4624 int32
	_ = v4624
	var v4628 int32
	_ = v4628
	var v4633 int32
	_ = v4633
	var v4637 int32
	_ = v4637
	var v4640 int32
	_ = v4640
	var v4644 int32
	_ = v4644
	var v4649 int32
	_ = v4649
	var v4653 int32
	_ = v4653
	var v4660 int32
	_ = v4660
	var v4665 int32
	_ = v4665
	var v4669 int32
	_ = v4669
	var v4672 int32
	_ = v4672
	var v4676 int32
	_ = v4676
	var v4681 int32
	_ = v4681
	var v4685 int32
	_ = v4685
	var v4688 int32
	_ = v4688
	var v4689 int64
	_ = v4689
	var v4692 int64
	_ = v4692
	var v4694 int64
	_ = v4694
	var v4695 int64
	_ = v4695
	var v4698 int64
	_ = v4698
	var v4704 int32
	_ = v4704
	var v4709 int32
	_ = v4709
	var v4713 int32
	_ = v4713
	var v4716 int32
	_ = v4716
	var v4720 int32
	_ = v4720
	var v4725 int32
	_ = v4725
	var v4729 int32
	_ = v4729
	var v4732 int32
	_ = v4732
	var v4736 int32
	_ = v4736
	var v4741 int32
	_ = v4741
	var v4745 int32
	_ = v4745
	var v4752 int32
	_ = v4752
	var v4757 int32
	_ = v4757
	var v4761 int32
	_ = v4761
	var v4764 int32
	_ = v4764
	var v4768 int32
	_ = v4768
	var v4773 int32
	_ = v4773
	var v4777 int32
	_ = v4777
	var v4783 int32
	_ = v4783
	var v4785 int32
	_ = v4785
	var v4787 int32
	_ = v4787
	var v4790 int32
	_ = v4790
	var v4817 int32
	_ = v4817
	var v4830 int32
	_ = v4830
	var v4831 int32
	_ = v4831
	var v4832 int32
	_ = v4832
	var v4833 int32
	_ = v4833
	var v4835 int32
	_ = v4835
	var v4845 int32
	_ = v4845
	var v4847 int32
	_ = v4847
	var v4848 int32
	_ = v4848
	var v4854 int32
	_ = v4854
	var v4857 int32
	_ = v4857
	var v4858 int32
	_ = v4858
	var v4861 int32
	_ = v4861
	var v4880 int32
	_ = v4880
	var v4884 int32
	_ = v4884
	var v4887 int32
	_ = v4887
	var v4889 int32
	_ = v4889
	var v4890 int32
	_ = v4890
	var v4913 int32
	_ = v4913
	var v4914 int32
	_ = v4914
	var v4917 int32
	_ = v4917
	var v4936 int32
	_ = v4936
	var v4940 int32
	_ = v4940
	var v4943 int32
	_ = v4943
	var v4945 int32
	_ = v4945
	var v4946 int32
	_ = v4946
	var v4951 int32
	_ = v4951
	var v4953 int32
	_ = v4953
	var v4963 int32
	_ = v4963
	var v4964 int32
	_ = v4964
	var v4966 int32
	_ = v4966
	var v4970 int32
	_ = v4970
	var v4974 int32
	_ = v4974
	var v4981 int32
	_ = v4981
	var v4983 int32
	_ = v4983
	var v4984 int32
	_ = v4984
	var v4985 int32
	_ = v4985
	var v4987 int32
	_ = v4987
	var v4988 int32
	_ = v4988
	var v4989 int32
	_ = v4989
	var v4992 int32
	_ = v4992
	var v4994 int32
	_ = v4994
	var v4996 int32
	_ = v4996
	var v5003 int32
	_ = v5003
	var v5005 int32
	_ = v5005
	var v5024 int32
	_ = v5024
	var v5045 int32
	_ = v5045
	var v5047 int32
	_ = v5047
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(1552)
	m.G0 = v22
	v24 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v26 = int32(4386280)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[534]))
	*(*int32)(unsafe.Add(mBase, _consts[534])) = v24
	switch v24 - int32(65) {
	case 0:
		goto L34
	case 1:
		goto L45
	case 2:
		goto L44
	case 3:
		goto L41
	case 4:
		goto L35
	default:
		goto L13
	case 8:
		goto L43
	case 10:
		goto L30
	case 12:
		goto L3
	case 14:
		goto L37
	case 15:
		goto L31
	case 17:
		goto L39
	case 18:
		goto L36
	case 19:
		goto L40
	case 20:
		goto L42
	case 24:
		goto L38
	case 33:
		goto L32
	case 34:
		goto L33
	case 47:
		goto L28
	case 49:
		goto L29
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[534])) = v27
	m.G0 = v22 + int32(1552)
	return
L4:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v5045 = m.ExcPending
	if v5045 != 0 {
		goto L1
	} else {
		goto L1312
	}
L5:
	;
	F_logicalrep_rel_close(m, v608, v5005)
	mBase = m.M
	v5024 = m.ExcPending
	if v5024 != 0 {
		goto L1
	} else {
		goto L1311
	}
L6:
	;
	v4981 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	F_AfterTriggerEndQuery(m, v4981)
	mBase = m.M
	v4983 = m.ExcPending
	if v4983 != 0 {
		goto L1
	} else {
		goto L1301
	}
L7:
	;
	F_ExecCloseIndices(m, v751)
	mBase = m.M
	v4970 = m.ExcPending
	if v4970 != 0 {
		goto L1
	} else {
		goto L1299
	}
L8:
	;
	F_slot_store_data(m, v781, v750, v22+int32(276))
	mBase = m.M
	v4951 = m.ExcPending
	if v4951 != 0 {
		goto L1
	} else {
		goto L1296
	}
L9:
	;
	v4845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1344)))
	v4847 = *(*int32)(unsafe.Add(mBase, _consts[529]))
	v4848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4847)+31)))
	F_ExecuteTruncateGuts(m, v4833, v4830, v4835, int32(0), v4845, (v4848^int32(-1))&int32(1))
	mBase = m.M
	v4854 = m.ExcPending
	if v4854 != 0 {
		goto L1
	} else {
		goto L1282
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[535])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _consts[536])) = int64(0)
	v4817 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[537])) = v4817
	*(*int32)(unsafe.Add(mBase, _consts[534])) = v4817
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v4817
	goto L3
L11:
	;
	F_stream_open_and_write_change(m, v2311, int32(65), v22+int32(304))
	mBase = m.M
	v4783 = m.ExcPending
	if v4783 != 0 {
		goto L1
	} else {
		goto L1278
	}
L12:
	;
	F_pa_switch_to_partial_serialize(m, v2320, int32(1))
	mBase = m.M
	v4777 = m.ExcPending
	if v4777 != 0 {
		goto L1
	} else {
		goto L1277
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4761 = m.ExcPending
	if v4761 != 0 {
		goto L1
	} else {
		goto L1273
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4745 = m.ExcPending
	if v4745 != 0 {
		goto L1
	} else {
		goto L1270
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4729 = m.ExcPending
	if v4729 != 0 {
		goto L1
	} else {
		goto L1266
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4713 = m.ExcPending
	if v4713 != 0 {
		goto L1
	} else {
		goto L1262
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4685 = m.ExcPending
	if v4685 != 0 {
		goto L1
	} else {
		goto L1258
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4669 = m.ExcPending
	if v4669 != 0 {
		goto L1
	} else {
		goto L1254
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4653 = m.ExcPending
	if v4653 != 0 {
		goto L1
	} else {
		goto L1251
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4637 = m.ExcPending
	if v4637 != 0 {
		goto L1
	} else {
		goto L1247
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4621 = m.ExcPending
	if v4621 != 0 {
		goto L1
	} else {
		goto L1243
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4605 = m.ExcPending
	if v4605 != 0 {
		goto L1
	} else {
		goto L1240
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4589 = m.ExcPending
	if v4589 != 0 {
		goto L1
	} else {
		goto L1236
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4573 = m.ExcPending
	if v4573 != 0 {
		goto L1
	} else {
		goto L1233
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4557 = m.ExcPending
	if v4557 != 0 {
		goto L1
	} else {
		goto L1229
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4541 = m.ExcPending
	if v4541 != 0 {
		goto L1
	} else {
		goto L1225
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4513 = m.ExcPending
	if v4513 != 0 {
		goto L1
	} else {
		goto L1221
	}
L28:
	;
	v4276 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(1352)))) = v4276
	v4278 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+1344)) = v4278
	v4281 = int32(*(*uint8)(unsafe.Add(mBase, _consts[539])))
	if v4281 != 0 {
		goto L16
	} else {
		goto L1146
	}
L29:
	;
	v3821 = v22 + int32(320)
	v3822 = m.G0
	v3824 = v3822 - int32(16)
	m.G0 = v3824
	v3826 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v3827 = m.ExcPending
	if v3827 != 0 {
		goto L1
	} else {
		goto L1029
	}
L30:
	;
	v3503 = v22 + int32(320)
	v3504 = m.G0
	v3506 = v3504 - int32(16)
	m.G0 = v3506
	v3508 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v3509 = m.ExcPending
	if v3509 != 0 {
		goto L1
	} else {
		goto L944
	}
L31:
	;
	F_logicalrep_read_prepare_common(m, l0, int32(361420), v22+int32(320))
	mBase = m.M
	v3336 = m.ExcPending
	if v3336 != 0 {
		goto L1
	} else {
		goto L894
	}
L32:
	;
	v3116 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v3117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3116)+16)))
	if v3117 == int32(1) {
		goto L834
	} else {
		goto L835
	}
L33:
	;
	v2938 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(1352)))) = v2938
	v2940 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+1344)) = v2940
	v2943 = int32(*(*uint8)(unsafe.Add(mBase, _consts[539])))
	if v2943 != 0 {
		goto L20
	} else {
		goto L781
	}
L34:
	;
	v2276 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+312)) = v2276
	v2278 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+304)) = v2278
	v2281 = int32(*(*uint8)(unsafe.Add(mBase, _consts[539])))
	if v2281 != 0 {
		goto L21
	} else {
		goto L615
	}
L35:
	;
	v2113 = int32(*(*uint8)(unsafe.Add(mBase, _consts[539])))
	if v2113 == int32(0) {
		goto L23
	} else {
		goto L572
	}
L36:
	;
	v1592 = int32(*(*uint8)(unsafe.Add(mBase, _consts[539])))
	if v1592 != 0 {
		goto L26
	} else {
		goto L468
	}
L37:
	;
	v1554 = int32(*(*uint8)(unsafe.Add(mBase, _consts[539])))
	if v1554 != 0 {
		goto L3
	} else {
		goto L457
	}
L38:
	;
	v1532 = F_handle_streamed_transaction(m, int32(89), l0)
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L1
	} else {
		goto L447
	}
L39:
	;
	v1349 = F_handle_streamed_transaction(m, int32(82), l0)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L1
	} else {
		goto L402
	}
L40:
	;
	v1021 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+320)) = uint8(v1021)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+1344)) = uint8(v1021)
	v1026 = *(*int64)(unsafe.Add(mBase, _consts[540]))
	if v1026 != int64(0) {
		goto L3
	} else {
		goto L318
	}
L41:
	;
	v857 = *(*int64)(unsafe.Add(mBase, _consts[540]))
	if v857 != int64(0) {
		goto L3
	} else {
		goto L259
	}
L42:
	;
	v496 = *(*int64)(unsafe.Add(mBase, _consts[540]))
	if v496 != int64(0) {
		goto L3
	} else {
		goto L160
	}
L43:
	;
	v166 = *(*int64)(unsafe.Add(mBase, _consts[540]))
	if v166 != int64(0) {
		goto L3
	} else {
		goto L75
	}
L44:
	;
	v105 = v22 + int32(320)
	v106 = m.G0
	v108 = v106 - int32(16)
	m.G0 = v108
	v110 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L62
	}
L45:
	;
	v33 = v22 + int32(320)
	v34 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v34
	if v34 == int64(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v52 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	F_errmsg_internal(m, int32(400183), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(490650), int32(68), int32(274111))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v52
	v56 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v22)+336))
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v60
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	*(*int64)(unsafe.Add(mBase, _consts[536])) = v63
	*(*int64)(unsafe.Add(mBase, _consts[541])) = v63
	v68 = *(*int32)(unsafe.Add(mBase, _consts[529]))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
	if v69 == int64(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v99 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[542])) = uint8(v99)
	F_pgstat_report_activity(m, int32(3), int32(0))
	mBase = m.M
	goto L3
L56:
	;
	if v63 != v69 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	*(*int64)(unsafe.Add(mBase, _consts[540])) = v63
	v77 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v77 == int32(0) {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v82 = *(*int64)(unsafe.Add(mBase, _consts[540]))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+20)) = uint32(v82)
	v85 = int64(base.Ui64(v82) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+16)) = uint32(v85)
	F_errmsg(m, int32(509741), v22+int32(16))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(490005), int32(4924), int32(167834))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L55
L62:
	;
	v113 = v110 & int32(255)
	if v113 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v127 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L69
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v113
	F_errmsg_internal(m, int32(399990), v108)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(490650), int32(104), int32(99768))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v105))) = v127
	v130 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v105)+8)) = v130
	v133 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v105)+16)) = v133
	m.G0 = v108 + int32(16)
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	v141 = *(*int64)(unsafe.Add(mBase, _consts[541]))
	if v139 != v141 {
		goto L27
	} else {
		goto L72
	}
L72:
	;
	F_apply_handle_commit_internal(m, v22+int32(320))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	F_process_syncing_tables(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v151 = int32(0)
	F_pgstat_report_activity(m, int32(2), v151)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[535])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _consts[536])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[537])) = v151
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v151
	goto L3
L75:
	;
	v170 = F_handle_streamed_transaction(m, int32(73), l0)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v170 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	if v173 < int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
	goto L82
L79:
	;
	v177 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[291])) = v177
	goto L81
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	if base.B2i32(v181 == int32(2)) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v190 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	F_PushActiveSnapshot(m, v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v196
	v201 = m.G0
	v203 = v201 - int32(16)
	m.G0 = v203
	v206 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v208 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v211 = v208 << (uint(int32(24)) % 32)
	if v211 != int32(1308622848) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	F_logicalrep_read_tuple(m, l0, v22+int32(320))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L98
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v211 >> (uint(int32(24)) % 32)
	F_errmsg_internal(m, int32(462855), v203)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(490650), int32(439), int32(81232))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	m.G0 = v203 + int32(16)
	v235 = F_logicalrep_rel_open(m, v206, int32(3))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L100
	}
L99:
	;
	F_logicalrep_rel_close(m, v235, v475)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L159
	}
L100:
	;
	v237 = F_should_apply_changes_for_rel(m, v235)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	if v237 == int32(0) {
		v475 = int32(3)
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _consts[529]))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+31)))
	if v243 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v235)+40))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+48))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+80))
	F_SwitchToUntrustedUser(m, v248, v22+int32(1344))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, _consts[537])) = v235
	v255 = F_create_edata_for_relation(m, v235)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v235)+40))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+52))
	v261 = F_ExecInitExtraTupleSlot(m, v257, v259, int32(1596068))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v257)+152))
	if v263 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v266 = F_MakePerTupleExprContext(m, v257)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	v268 = v263
	goto L111
L111:
	;
	v269 = int32(4476144)
	v270 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v272
	F_slot_store_data(m, v261, v235, v22+int32(320))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L113
	}
L112:
	;
	v268 = v266
	goto L111
L113:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v235)+40))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+52))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v257)+152))
	if v281 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v284 = F_MakePerTupleExprContext(m, v257)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	v286 = v281
	goto L116
L116:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v235)+12))
	if v280 == v287 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v286 = v284
	goto L116
L118:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v270
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v235)+40))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+48))
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+119)))
	if v427 == int32(112) {
		goto L140
	} else {
		goto L141
	}
L119:
	;
	v290 = v280 << (uint(int32(2)) % 32)
	v291 = F_palloc(m, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v293 = F_palloc(m, v290)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	if v280 <= int32(0) {
		goto L118
	} else {
		goto L122
	}
L122:
	;
	v300 = int32(0)
	v306 = v2
	goto L123
L123:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v325 = v279 + int32(20) + v319<<(uint(int32(4))%32) + v300*int32(100)
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+91)))
	if v326 != 0 {
		v357 = v306
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if v357 <= int32(0) {
		goto L118
	} else {
		goto L134
	}
L125:
	;
	v360 = v300 + int32(1)
	if v360 != v280 {
		v300 = v360
		v306 = v357
		goto L123
	} else {
		goto L133
	}
L126:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+90)))
	if v327 != 0 {
		v357 = v306
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v235)+44))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v333 = int32(*(*int16)(unsafe.Add(mBase, uint32(v329+v300<<(uint(int32(1))%32)))))
	if int32(0) <= v333 {
		v357 = v306
		goto L125
	} else {
		goto L128
	}
L128:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v235)+40))
	v339 = F_build_column_default(m, v336, v300+int32(1))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	if v339 == int32(0) {
		v357 = v306
		goto L125
	} else {
		goto L130
	}
L130:
	;
	v344 = v306 << (uint(int32(2)) % 32)
	v346 = F_expression_planner(m, v339)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v349 = F_ExecInitExpr(m, v346, int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293+v344))) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v291+v344))) = v300
	v357 = v306 + int32(1)
	goto L125
L133:
	;
	goto L124
L134:
	;
	v365 = int32(0)
	goto L135
L135:
	;
	v385 = v365 << (uint(int32(2)) % 32)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v293+v385)))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v261)+20))
	v389 = v385 + v291
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v387)+20))
	v393 = m.T0[v392].(func(*base.Module, int32, int32, int32) int32)(m, v387, v286, v388+v390)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L137
	}
L136:
	;
	goto L118
L137:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	*(*int32)(unsafe.Add(mBase, uint32(v395+v396<<(uint(int32(2))%32)))) = v393
	v402 = v365 + int32(1)
	if v402 != v357 {
		v365 = v402
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	F_AfterTriggerEndQuery(m, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L149
	}
L140:
	;
	F_apply_handle_tuple_routing(m, v255, v261, int32(0), int32(3))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	F_ExecOpenIndices(m, v434, int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L144
	}
L143:
	;
	goto L139
L144:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	F_InitConflictIndexes(m, v434)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v434)+8))
	F_TargetPrivilegesCheck(m, v441, int64(1))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_ExecSimpleRelationInsert(m, v434, v438, v261)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_ExecCloseIndices(m, v434)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	goto L139
L149:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v255)+16))
	if v454 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	F_ExecCleanupTupleRouting(m, v455, v454)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v458 = int32(0)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v451)+104))
	F_ExecResetTupleTable(m, v459, v458)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L154
	}
L153:
	;
	goto L152
L154:
	;
	F_FreeExecutorState(m, v451)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_pfree(m, v255)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, _consts[537])) = int32(0)
	if v243 != 0 {
		v475 = v458
		goto L99
	} else {
		goto L157
	}
L157:
	;
	F_RestoreUserContext(m, v22+int32(1344))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v475 = v458
	goto L99
L159:
	;
	goto L4
L160:
	;
	v500 = F_handle_streamed_transaction(m, int32(85), l0)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	if v500 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	v503 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	if v503 < int32(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v510 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v510)+20))
	goto L167
L164:
	;
	v507 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[291])) = v507
	goto L166
L165:
	;
	goto L166
L166:
	;
	goto L163
L167:
	;
	if base.B2i32(v511 == int32(2)) == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v520 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	F_PushActiveSnapshot(m, v520)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v526
	v530 = v22 + int32(275)
	v535 = m.G0
	v537 = v535 - int32(32)
	m.G0 = v537
	v540 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L176
	}
L175:
	;
	v608 = F_logicalrep_rel_open(m, v540, int32(3))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L198
	}
L176:
	;
	v542 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L195
	}
L178:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L192
	}
L179:
	;
	if v542&int32(251) != int32(75) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v549 = v542 << (uint(int32(24)) % 32)
	if v549 != int32(1308622848) {
		goto L178
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v553 = int32(75)
	if v542&v553 == v553 {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	goto L182
L184:
	;
	v567 = v565 << (uint(int32(24)) % 32)
	if v567 != int32(1308622848) {
		goto L177
	} else {
		goto L190
	}
L185:
	;
	F_logicalrep_read_tuple(m, l0, v22+int32(288))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v563 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v530))) = uint8(v563)
	v565 = v542
	goto L184
L188:
	;
	v559 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v530))) = uint8(v559)
	v561 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v565 = v561
	goto L184
L190:
	;
	F_logicalrep_read_tuple(m, l0, v22+int32(276))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	m.G0 = v537 + int32(32)
	goto L175
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v549 >> (uint(int32(24)) % 32)
	F_errmsg_internal(m, int32(495027), v537)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(490650), int32(501), int32(352105))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537)+16)) = v567 >> (uint(int32(24)) % 32)
	F_errmsg_internal(m, int32(494964), v537+int32(16))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(490650), int32(517), int32(352105))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	v610 = F_should_apply_changes_for_rel(m, v608)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	if v610 == int32(0) {
		v5005 = int32(3)
		goto L5
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, _consts[537])) = v608
	F_check_relation_updatable(m, v608)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	v619 = *(*int32)(unsafe.Add(mBase, _consts[529]))
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+31)))
	if v620 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v608)+40))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v623)+48))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v624)+80))
	F_SwitchToUntrustedUser(m, v625, v22+int32(304))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v630 = F_create_edata_for_relation(m, v608)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L206
	}
L205:
	;
	goto L204
L206:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v608)+40))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v633)+52))
	v636 = F_ExecInitExtraTupleSlot(m, v632, v634, int32(1596068))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v636)+12))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	if int32(0) < v639 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v632)+32))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+12))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v643)))
	v646 = int32(0)
	v647 = v639
	v649 = v638
	goto L211
L209:
	;
	goto L210
L210:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v632)+152))
	if v718 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L211:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649+v647<<(uint(int32(4))%32)+v646*int32(100))+111)))
	if v671 != 0 {
		v693 = v647
		v694 = v649
		goto L213
	} else {
		goto L214
	}
L212:
	;
	goto L210
L213:
	;
	v697 = v646 + int32(1)
	if v697 < v693 {
		v646 = v697
		v647 = v693
		v649 = v694
		goto L211
	} else {
		goto L218
	}
L214:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v608)+44))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v672)))
	v677 = int32(*(*int16)(unsafe.Add(mBase, uint32(v673+v646<<(uint(int32(1))%32)))))
	if v677 < int32(0) {
		v693 = v647
		v694 = v649
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v22)+280))
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680+v677))))
	if v682 == int32(117) {
		v693 = v647
		v694 = v649
		goto L213
	} else {
		goto L216
	}
L216:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v644)+36))
	v688 = F_bms_add_member(m, v685, v646+int32(8))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v644)+36)) = v688
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v636)+12))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v691)))
	v693 = v692
	v694 = v691
	goto L213
L218:
	;
	goto L212
L219:
	;
	v721 = F_MakePerTupleExprContext(m, v632)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	v723 = v718
	goto L221
L221:
	;
	v724 = int32(4476144)
	v725 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v723)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v727
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+275)))
	if v733 != 0 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	v723 = v721
	goto L221
L223:
	;
	v734 = v22 + int32(288)
	goto L225
L224:
	;
	v734 = v22 + int32(276)
	goto L225
L225:
	;
	F_slot_store_data(m, v636, v608, v734)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v725
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v608)+40))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v739)+48))
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740)+119)))
	if v741 == int32(112) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	F_apply_handle_tuple_routing(m, v630, v636, v22+int32(276), int32(2))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v608)+52))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v630)+8))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v751)+8))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	v756 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(1360)))) = v756
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(1352)))) = v756
	*(*int64)(unsafe.Add(mBase, uint32(v22)+1344)) = v756
	v766 = int32(0)
	F_EvalPlanQualInit(m, v22+int32(320), v753, v766, v766, int32(-1), v766)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L231
	}
L230:
	;
	goto L6
L231:
	;
	F_ExecOpenIndices(m, v751, int32(0))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	F_TargetPrivilegesCheck(m, v752, int64(2))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	v781 = F_table_slot_create(m, v752, v775+int32(104))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	if v749 != 0 {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v795 = F_GetTupleTransactionInfo(m, v781, v22+int32(1352), v22+int32(1356), v22+int32(1360))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L244
	}
L236:
	;
	v783 = F_RelationFindReplTupleByIndex(m, v752, v749, v636, v781)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L1
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v785 = F_RelationFindReplTupleSeq(m, v752, v636, v781)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L241
	}
L239:
	;
	if v783 != 0 {
		goto L235
	} else {
		goto L240
	}
L240:
	;
	goto L8
L241:
	;
	if v785 == int32(0) {
		goto L8
	} else {
		goto L242
	}
L242:
	;
	goto L235
L243:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v753)+152))
	if v828 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L244:
	;
	if v795 == int32(0) {
		goto L243
	} else {
		goto L245
	}
L245:
	;
	v799 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+1356)))
	v801 = int32(*(*uint16)(unsafe.Add(mBase, _consts[107])))
	if v799 == v801 {
		goto L243
	} else {
		goto L246
	}
L246:
	;
	v805 = F_table_slot_create(m, v752, v753+int32(104))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	F_slot_store_data(m, v805, v750, v22+int32(276))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+1344)) = v781
	v813 = v22 + int32(1344)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+300)) = v813
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v813
	v819 = int32(1)
	v823 = F_list_make1_impl(m, v819, v22+int32(56))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	F_ReportApplyConflict(m, v753, v751, int32(15), v819, v636, v805, v823)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	goto L243
L251:
	;
	v831 = F_MakePerTupleExprContext(m, v753)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L254
	}
L252:
	;
	v833 = v828
	goto L253
L253:
	;
	v834 = int32(4476144)
	v835 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v833)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v837
	F_slot_modify_data(m, v636, v781, v750, v22+int32(276))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L255
	}
L254:
	;
	v833 = v831
	goto L253
L255:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v835
	*(*int32)(unsafe.Add(mBase, uint32(v22)+348)) = v636
	F_InitConflictIndexes(m, v751)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v751)+8))
	F_TargetPrivilegesCheck(m, v848, int64(4))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	F_ExecSimpleRelationUpdate(m, v751, v753, v22+int32(320), v781, v636)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	goto L7
L259:
	;
	v861 = F_handle_streamed_transaction(m, int32(68), l0)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	if v861 != 0 {
		goto L3
	} else {
		goto L261
	}
L261:
	;
	v864 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	if v864 < int32(0) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	v871 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v871)+20))
	goto L266
L263:
	;
	v868 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[291])) = v868
	goto L265
L264:
	;
	goto L265
L265:
	;
	goto L262
L266:
	;
	if base.B2i32(v872 == int32(2)) == int32(0) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L1
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v881 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L1
	} else {
		goto L272
	}
L270:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	F_PushActiveSnapshot(m, v881)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v887 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v887
	v892 = m.G0
	v894 = v892 - int32(16)
	m.G0 = v894
	v897 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	v899 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	if v899&int32(251) != int32(75) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	F_logicalrep_read_tuple(m, l0, v22+int32(320))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L1
	} else {
		goto L282
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v894))) = base.I32_extend8_s(v899)
	F_errmsg_internal(m, int32(494992), v894)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	F_errfinish(m, int32(490650), int32(572), int32(347436))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L282:
	;
	m.G0 = v894 + int32(16)
	v925 = F_logicalrep_rel_open(m, v897, int32(3))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L1
	} else {
		goto L284
	}
L283:
	;
	F_logicalrep_rel_close(m, v925, v1017)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L1
	} else {
		goto L317
	}
L284:
	;
	v927 = F_should_apply_changes_for_rel(m, v925)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	if v927 == int32(0) {
		v1017 = int32(3)
		goto L283
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, _consts[537])) = v925
	F_check_relation_updatable(m, v925)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	v936 = *(*int32)(unsafe.Add(mBase, _consts[529]))
	v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v936)+31)))
	if v937 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v925)+40))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v940)+48))
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v941)+80))
	F_SwitchToUntrustedUser(m, v942, v22+int32(1344))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L1
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v947 = F_create_edata_for_relation(m, v925)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L292
	}
L291:
	;
	goto L290
L292:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v947)))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v925)+40))
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v950)+52))
	v953 = F_ExecInitExtraTupleSlot(m, v949, v951, int32(1596068))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v949)+152))
	if v955 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v958 = F_MakePerTupleExprContext(m, v949)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L297
	}
L295:
	;
	v960 = v955
	goto L296
L296:
	;
	v961 = int32(4476144)
	v962 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v960)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v964
	F_slot_store_data(m, v953, v925, v22+int32(320))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L298
	}
L297:
	;
	v960 = v958
	goto L296
L298:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v962
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v925)+40))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v972)+48))
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v973)+119)))
	if v974 == int32(112) {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v947)))
	F_AfterTriggerEndQuery(m, v991)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L1
	} else {
		goto L307
	}
L300:
	;
	F_apply_handle_tuple_routing(m, v947, v953, int32(0), int32(4))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v947)+8))
	F_ExecOpenIndices(m, v981, int32(0))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L1
	} else {
		goto L304
	}
L303:
	;
	goto L299
L304:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v925)+52))
	F_apply_handle_delete_internal(m, v947, v981, v953, v985)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	F_ExecCloseIndices(m, v981)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	goto L299
L307:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v947)+16))
	if v994 != 0 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v947)+12))
	F_ExecCleanupTupleRouting(m, v995, v994)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	v998 = int32(0)
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v991)+104))
	F_ExecResetTupleTable(m, v999, v998)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L312
	}
L311:
	;
	goto L310
L312:
	;
	F_FreeExecutorState(m, v991)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	F_pfree(m, v947)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, _consts[537])) = int32(0)
	if v937 != 0 {
		v1017 = v998
		goto L283
	} else {
		goto L315
	}
L315:
	;
	F_RestoreUserContext(m, v22+int32(1344))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v1017 = v998
	goto L283
L317:
	;
	goto L4
L318:
	;
	v1030 = F_handle_streamed_transaction(m, int32(84), l0)
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	if v1030 != 0 {
		goto L3
	} else {
		goto L320
	}
L320:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	if v1033 < int32(0) {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+20))
	goto L325
L322:
	;
	v1037 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[291])) = v1037
	goto L324
L323:
	;
	goto L324
L324:
	;
	goto L321
L325:
	;
	if base.B2i32(v1041 == int32(2)) == int32(0) {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L1
	} else {
		goto L329
	}
L327:
	;
	goto L328
L328:
	;
	v1050 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L1
	} else {
		goto L331
	}
L329:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	goto L328
L331:
	;
	F_PushActiveSnapshot(m, v1050)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1056
	v1063 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L334
	}
L333:
	;
	if v1127 == int32(0) {
		v4830 = v2
		v4831 = v2
		v4832 = v2
		v4833 = v2
		v4835 = v2
		goto L9
	} else {
		goto L342
	}
L334:
	;
	v1066 = F_pq_getmsgint(m, l0, int32(1))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	v1068 = int32(1)
	v1069 = v1066 & v1068
	*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(320)))) = uint8(v1069)
	v1074 = int32(base.Ui32(v1066)>>(uint(v1068)%32)) & v1068
	*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(1344)))) = uint8(v1074)
	v1076 = int32(0)
	if v1063 <= v1076 {
		v1127 = v1076
		goto L333
	} else {
		goto L336
	}
L336:
	;
	v1082 = int32(0)
	v1085 = v1076
	goto L337
L337:
	;
	v1101 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L339
	}
L338:
	;
	v1127 = v1103
	goto L333
L339:
	;
	v1103 = F_lappend_oid(m, v1082, v1101)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	v1106 = v1085 + int32(1)
	if v1106 != v1063 {
		v1082 = v1103
		v1085 = v1106
		goto L337
	} else {
		goto L341
	}
L341:
	;
	goto L338
L342:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+4))
	if v1130 <= int32(0) {
		v4830 = v2
		v4831 = v2
		v4832 = v2
		v4833 = v2
		v4835 = v2
		goto L9
	} else {
		goto L343
	}
L343:
	;
	v1138 = v2
	v1139 = v2
	v1140 = v2
	v1141 = v2
	v1142 = v2
	v1143 = v2
	goto L344
L344:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+12))
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1152+v1142<<(uint(int32(2))%32))))
	v1158 = F_logicalrep_rel_open(m, v1156, int32(8))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L1
	} else {
		goto L347
	}
L345:
	;
	v4830 = v1330
	v4831 = v1331
	v4832 = v1332
	v4833 = v1333
	v4835 = v1335
	goto L9
L346:
	;
	v1345 = v1142 + int32(1)
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+4))
	if v1345 < v1346 {
		v1138 = v1330
		v1139 = v1331
		v1140 = v1332
		v1141 = v1333
		v1142 = v1345
		v1143 = v1335
		goto L344
	} else {
		goto L401
	}
L347:
	;
	v1160 = F_should_apply_changes_for_rel(m, v1158)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	if v1160 == int32(0) {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	F_logicalrep_rel_close(m, v1158, int32(8))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L1
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1167 = F_lappend(m, v1139, v1158)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L353
	}
L352:
	;
	v1330 = v1138
	v1331 = v1139
	v1332 = v1140
	v1333 = v1141
	v1335 = v1143
	goto L346
L353:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+40))
	F_TargetPrivilegesCheck(m, v1169, int64(16))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+40))
	v1174 = F_lappend(m, v1141, v1173)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+36))
	v1177 = F_lappend_oid(m, v1138, v1176)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v1180 < int32(2) {
		v1199 = v1143
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+40))
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+48))
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1201)+119)))
	if v1202 != int32(112) {
		v1330 = v1177
		v1331 = v1167
		v1332 = v1140
		v1333 = v1174
		v1335 = v1199
		goto L346
	} else {
		goto L364
	}
L358:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+40))
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1183)+48))
	v1185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1184)+118)))
	if v1185 != int32(112) {
		v1199 = v1143
		goto L357
	} else {
		goto L359
	}
L359:
	;
	v1188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1184)+119)))
	if v1188 == int32(102) {
		v1199 = v1143
		goto L357
	} else {
		goto L360
	}
L360:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1183)+56))
	goto L361
L361:
	;
	if base.Ui32(v1191) < base.Ui32(int32(12000)) {
		v1199 = v1143
		goto L357
	} else {
		goto L362
	}
L362:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+36))
	v1195 = F_lappend_oid(m, v1143, v1194)
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	v1199 = v1195
	goto L357
L364:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+36))
	v1208 = F_find_all_inheritors(m, v1205, int32(8), int32(0))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	if v1208 == int32(0) {
		v1330 = v1177
		v1331 = v1167
		v1332 = v1140
		v1333 = v1174
		v1335 = v1199
		goto L346
	} else {
		goto L366
	}
L366:
	;
	v1212 = int32(0)
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+4))
	if v1213 <= v1212 {
		v1330 = v1177
		v1331 = v1167
		v1332 = v1140
		v1333 = v1174
		v1335 = v1199
		goto L346
	} else {
		goto L367
	}
L367:
	;
	v1216 = v1212
	v1221 = v1177
	v1223 = v1140
	v1224 = v1174
	v1226 = v1199
	goto L368
L368:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+12))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1235+v1216<<(uint(int32(2))%32))))
	v1240 = int32(0)
	if v1221 == v1240 {
		goto L372
	} else {
		goto L373
	}
L369:
	;
	v1330 = v1316
	v1331 = v1167
	v1332 = v1317
	v1333 = v1318
	v1335 = v1319
	goto L346
L370:
	;
	v1322 = v1216 + int32(1)
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+4))
	if v1322 < v1323 {
		v1216 = v1322
		v1221 = v1316
		v1223 = v1317
		v1224 = v1318
		v1226 = v1319
		goto L368
	} else {
		goto L400
	}
L371:
	;
	if v1278 != 0 {
		v1316 = v1221
		v1317 = v1223
		v1318 = v1224
		v1319 = v1226
		goto L370
	} else {
		goto L384
	}
L372:
	;
	v1278 = int32(0)
	goto L371
L373:
	;
	goto L374
L374:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+4))
	if v1246 <= int32(0) {
		v1271 = v1240
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v1278 = v1271
	goto L371
L376:
	;
	v1249 = int32(0)
	if v1249 < v1246 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1252 = v1246
	goto L379
L378:
	;
	v1252 = v1249
	goto L379
L379:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+12))
	v1255 = int32(0)
	goto L380
L380:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1253+v1255<<(uint(int32(2))%32))))
	v1264 = base.B2i32(v1263 == v1239)
	if v1263 == v1239 {
		v1271 = v1264
		goto L375
	} else {
		goto L382
	}
L381:
	;
	v1271 = v1264
	goto L375
L382:
	;
	v1266 = v1255 + int32(1)
	if v1266 != v1252 {
		v1255 = v1266
		goto L380
	} else {
		goto L383
	}
L383:
	;
	goto L381
L384:
	;
	v1280 = F_table_open(m, v1239, int32(0))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L1
	} else {
		goto L386
	}
L385:
	;
	F_TargetPrivilegesCheck(m, v1280, int64(16))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L1
	} else {
		goto L390
	}
L386:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+48))
	v1283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1282)+118)))
	if v1283 != int32(116) {
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v1286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280)+24)))
	if v1286 != 0 {
		goto L385
	} else {
		goto L388
	}
L388:
	;
	F_sequence_close(m, v1280, int32(8))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	v1316 = v1221
	v1317 = v1223
	v1318 = v1224
	v1319 = v1226
	goto L370
L390:
	;
	v1293 = F_lappend(m, v1224, v1280)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	v1295 = F_lappend(m, v1223, v1280)
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	v1297 = F_lappend_oid(m, v1221, v1239)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v1300 < int32(2) {
		v1316 = v1297
		v1317 = v1295
		v1318 = v1293
		v1319 = v1226
		goto L370
	} else {
		goto L394
	}
L394:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+48))
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303)+118)))
	if v1304 != int32(112) {
		v1316 = v1297
		v1317 = v1295
		v1318 = v1293
		v1319 = v1226
		goto L370
	} else {
		goto L395
	}
L395:
	;
	v1307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303)+119)))
	if v1307 == int32(102) {
		v1316 = v1297
		v1317 = v1295
		v1318 = v1293
		v1319 = v1226
		goto L370
	} else {
		goto L396
	}
L396:
	;
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+56))
	goto L397
L397:
	;
	if base.Ui32(v1310) < base.Ui32(int32(12000)) {
		v1316 = v1297
		v1317 = v1295
		v1318 = v1293
		v1319 = v1226
		goto L370
	} else {
		goto L398
	}
L398:
	;
	v1313 = F_lappend_oid(m, v1226, v1239)
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	v1316 = v1297
	v1317 = v1295
	v1318 = v1293
	v1319 = v1313
	goto L370
L400:
	;
	goto L369
L401:
	;
	goto L345
L402:
	;
	if v1349 != 0 {
		goto L3
	} else {
		goto L403
	}
L403:
	;
	v1352 = F_palloc(m, int32(32))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	v1355 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1352))) = v1355
	v1358 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	v1361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1358))))
	if v1361 != 0 {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v1362 = v1358
	goto L409
L408:
	;
	v1362 = int32(323805)
	goto L409
L409:
	;
	v1363 = F_pstrdup(m, v1362)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1352)+4)) = v1363
	v1366 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	v1368 = F_pstrdup(m, v1366)
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1352)+8)) = v1368
	v1371 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1352)+24)) = uint8(v1371)
	v1376 = F_pq_getmsgint(m, l0, int32(2))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	v1379 = v1376 << (uint(int32(2)) % 32)
	v1380 = F_palloc(m, v1379)
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	v1382 = F_palloc(m, v1379)
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	if int32(0) < v1376 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v1390 = int32(0)
	v1393 = v2
	goto L420
L418:
	;
	v1438 = v2
	goto L419
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1352)+28)) = v1438
	*(*int32)(unsafe.Add(mBase, uint32(v1352)+20)) = v1382
	*(*int32)(unsafe.Add(mBase, uint32(v1352)+16)) = v1380
	*(*int32)(unsafe.Add(mBase, uint32(v1352)+12)) = v1376
	F_logicalrep_relmap_update(m, v1352)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L1
	} else {
		goto L432
	}
L420:
	;
	v1405 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L1
	} else {
		goto L422
	}
L421:
	;
	v1438 = v1411
	goto L419
L422:
	;
	if v1405&int32(1) != 0 {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v1409 = F_bms_add_member(m, v1393, v1390)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L1
	} else {
		goto L426
	}
L424:
	;
	v1411 = v1393
	goto L425
L425:
	;
	v1413 = v1390 << (uint(int32(2)) % 32)
	v1415 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L1
	} else {
		goto L427
	}
L426:
	;
	v1411 = v1409
	goto L425
L427:
	;
	v1417 = F_pstrdup(m, v1415)
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1380+v1413))) = v1417
	v1422 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L1
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1413+v1382))) = v1422
	v1426 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	v1429 = v1390 + int32(1)
	if v1429 != v1376 {
		v1390 = v1429
		v1393 = v1411
		goto L420
	} else {
		goto L431
	}
L431:
	;
	goto L421
L432:
	;
	v1456 = m.G0
	v1458 = v1456 - int32(32)
	m.G0 = v1458
	v1461 = *(*int32)(unsafe.Add(mBase, _consts[543]))
	if v1461 == int32(0) {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	m.G0 = v1458 + int32(32)
	goto L3
L434:
	;
	F_hash_seq_init(m, v1458+int32(12), v1461)
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L1
	} else {
		goto L435
	}
L435:
	;
	v1470 = F_hash_seq_search(m, v1458+int32(12))
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L1
	} else {
		goto L436
	}
L436:
	;
	if v1470 == int32(0) {
		goto L433
	} else {
		goto L437
	}
L437:
	;
	v1475 = v1470
	goto L438
L438:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1475)+8))
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1352)))
	if v1493 == v1494 {
		goto L440
	} else {
		goto L441
	}
L439:
	;
	goto L433
L440:
	;
	v1497 = v1475 + int32(8)
	F_logicalrep_relmap_free_entry(m, v1497)
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L1
	} else {
		goto L443
	}
L441:
	;
	goto L442
L442:
	;
	v1507 = F_hash_seq_search(m, v1458+int32(12))
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L1
	} else {
		goto L445
	}
L443:
	;
	v1503 = F__emscripten_memset_bulkmem(m, v1497, base.I32_extend8_s(int32(0)), int32(72))
	mBase = m.M
	goto L444
L444:
	;
	goto L442
L445:
	;
	if v1507 != 0 {
		v1475 = v1507
		goto L438
	} else {
		goto L446
	}
L446:
	;
	goto L439
L447:
	;
	if v1532 != 0 {
		goto L3
	} else {
		goto L448
	}
L448:
	;
	v1535 = v22 + int32(320)
	v1537 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L1
	} else {
		goto L449
	}
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1535))) = v1537
	v1540 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L1
	} else {
		goto L450
	}
L450:
	;
	v1543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1540))))
	if v1543 != 0 {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v1544 = v1540
	goto L453
L452:
	;
	v1544 = int32(323805)
	goto L453
L453:
	;
	v1545 = F_pstrdup(m, v1544)
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L1
	} else {
		goto L454
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1535)+4)) = v1545
	v1548 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	v1550 = F_pstrdup(m, v1548)
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1535)+8)) = v1550
	goto L3
L457:
	;
	v1556 = int32(*(*uint8)(unsafe.Add(mBase, _consts[542])))
	if v1556 != int32(1) {
		goto L458
	} else {
		goto L459
	}
L458:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L1
	} else {
		goto L464
	}
L459:
	;
	v1560 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+20))
	goto L460
L460:
	;
	if base.B2i32(v1561 == int32(2)) == int32(0) {
		goto L3
	} else {
		goto L461
	}
L461:
	;
	v1567 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v1568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1567)+16)))
	if v1568 != int32(1) {
		goto L458
	} else {
		goto L462
	}
L462:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1567)))
	if v1571 == int32(1) {
		goto L3
	} else {
		goto L463
	}
L463:
	;
	goto L458
L464:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	F_errmsg_internal(m, int32(224832), int32(0))
	mBase = m.M
	v1585 = m.ExcPending
	if v1585 != 0 {
		goto L1
	} else {
		goto L466
	}
L466:
	;
	F_errfinish(m, int32(490005), int32(1436), int32(273753))
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
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
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1597 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[539])) = uint8(v1597)
	v1602 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	v1604 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L1
	} else {
		goto L470
	}
L470:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(304)))) = uint8(base.B2i32(v1604 == int32(1)))
	*(*int32)(unsafe.Add(mBase, _consts[544])) = v1602
	if v1602 == int32(0) {
		goto L25
	} else {
		goto L471
	}
L471:
	;
	*(*int64)(unsafe.Add(mBase, _consts[536])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v1602
	v1618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+304)))
	if v1618 == int32(1) {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v1621 = m.G0
	v1623 = v1621 + int32(-64)
	m.G0 = v1623
	*(*int32)(unsafe.Add(mBase, uint32(v1623)+60)) = v1602
	v1627 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1627)))
	if v1628 != int32(2) {
		goto L477
	} else {
		goto L478
	}
L473:
	;
	v1979 = v1602
	goto L474
L474:
	;
	v1998 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v1999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1998)+16)))
	if v1999 == int32(1) {
		goto L538
	} else {
		goto L539
	}
L475:
	;
	v1977 = *(*int32)(unsafe.Add(mBase, _consts[544]))
	v1979 = v1977
	goto L474
L476:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		goto L1
	} else {
		goto L533
	}
L477:
	;
	m.G0 = v1623 - int32(-64)
	goto L475
L478:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L1
	} else {
		goto L479
	}
L479:
	;
	v1634 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v1635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1634)+68)))
	if v1635 != int32(1) {
		goto L477
	} else {
		goto L480
	}
L480:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, _consts[529]))
	v1640 = *(*int64)(unsafe.Add(mBase, uint32(v1639)+8))
	if v1640 != int64(0) {
		goto L477
	} else {
		goto L481
	}
L481:
	;
	v1643 = F_AllTablesyncsReady(m)
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	if v1643 == int32(0) {
		goto L477
	} else {
		goto L483
	}
L483:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	if v1648 == int32(0) {
		goto L485
	} else {
		goto L486
	}
L484:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, _consts[545]))
	if v1879 == int32(0) {
		goto L523
	} else {
		goto L524
	}
L485:
	;
	v1704 = int32(4476144)
	v1705 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1708 = *(*int32)(unsafe.Add(mBase, _consts[533]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1708
	v1711 = F_palloc0(m, int32(20))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L1
	} else {
		goto L492
	}
L486:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+4))
	if v1651 <= int32(0) {
		goto L485
	} else {
		goto L487
	}
L487:
	;
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+12))
	v1657 = int32(0)
	goto L488
L488:
	;
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1654+v1657<<(uint(int32(2))%32))))
	v1679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1678)+13)))
	if v1679 == int32(0) {
		v1865 = v1678
		goto L484
	} else {
		goto L490
	}
L489:
	;
	goto L485
L490:
	;
	v1683 = v1657 + int32(1)
	if v1651 != v1683 {
		v1657 = v1683
		goto L488
	} else {
		goto L491
	}
L491:
	;
	goto L489
L492:
	;
	v1715 = F_add_size(m, int32(0), int32(96))
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L1
	} else {
		goto L493
	}
L493:
	;
	v1718 = F_add_size(m, v1715, int32(16777216))
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	v1721 = F_add_size(m, v1718, int32(16384))
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L1
	} else {
		goto L495
	}
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1623)+8)) = v1721
	v1726 = F_add_size(m, int32(0), int32(3))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1623)+12)) = v1726
	v1731 = F_shm_toc_estimate(m, v1621+int32(-56))
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	v1735 = F_shm_toc_estimate(m, v1621+int32(-56))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	v1738 = F_dsm_create(m, v1735, int32(0))
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	if v1738 == int32(0) {
		goto L500
	} else {
		goto L501
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1705
	F_pfree(m, v1711)
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L1
	} else {
		goto L503
	}
L501:
	;
	goto L502
L502:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v1738)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1747)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1747)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1747))) = int64(2021433447)
	*(*int32)(unsafe.Add(mBase, uint32(v1747)+12)) = v1731 & int32(-32)
	goto L504
L503:
	;
	goto L477
L504:
	;
	v1757 = F_shm_toc_allocate(m, v1747, int32(80))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	v1759 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1757)+20)) = v1759
	*(*int32)(unsafe.Add(mBase, uint32(v1757)+8)) = v1759
	*(*int32)(unsafe.Add(mBase, uint32(v1757))) = v1759
	*(*int32)(unsafe.Add(mBase, uint32(v1757)+32)) = v1759
	*(*int64)(unsafe.Add(mBase, uint32(v1757)+24)) = int64(0)
	F_shm_toc_insert(m, v1747, int64(1), v1757)
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	v1774 = F_shm_toc_allocate(m, v1747, int32(16777216))
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L1
	} else {
		goto L507
	}
L507:
	;
	v1777 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1774)+16)) = v1777
	*(*int32)(unsafe.Add(mBase, uint32(v1774)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1774))) = v1777
	v1783 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v1774)+36)) = uint16(v1783)
	*(*int64)(unsafe.Add(mBase, uint32(v1774)+24)) = v1777
	*(*int32)(unsafe.Add(mBase, uint32(v1774)+32)) = int32(16777176)
	goto L508
L508:
	;
	F_shm_toc_insert(m, v1747, int64(2), v1774)
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L1
	} else {
		goto L509
	}
L509:
	;
	v1795 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	F_shm_mq_set_sender(m, v1774, v1795)
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L1
	} else {
		goto L510
	}
L510:
	;
	v1798 = F_shm_mq_attach(m, v1774, v1738)
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1711))) = v1798
	v1803 = F_shm_toc_allocate(m, v1747, int32(16384))
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v1806 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1803)+16)) = v1806
	*(*int32)(unsafe.Add(mBase, uint32(v1803)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1803))) = v1806
	v1812 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v1803)+36)) = uint16(v1812)
	*(*int64)(unsafe.Add(mBase, uint32(v1803)+24)) = v1806
	*(*int32)(unsafe.Add(mBase, uint32(v1803)+32)) = int32(16344)
	goto L513
L513:
	;
	F_shm_toc_insert(m, v1747, int64(3), v1803)
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	F_shm_mq_set_receiver(m, v1803, v1824)
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	v1827 = F_shm_mq_attach(m, v1803, v1738)
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1711)+16)) = v1757
	*(*int32)(unsafe.Add(mBase, uint32(v1711)+8)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v1711)+4)) = v1827
	v1834 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1834)+24))
	v1837 = *(*int32)(unsafe.Add(mBase, _consts[529]))
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1837)))
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1837)+16))
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1834)+28))
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1738)+12))
	v1843 = F_logicalrep_worker_launch(m, int32(3), v1835, v1838, v1839, v1840, int32(0), v1842)
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	if v1843 == int32(0) {
		goto L518
	} else {
		goto L519
	}
L518:
	;
	F_pa_free_worker_info(m, v1711)
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L1
	} else {
		goto L521
	}
L519:
	;
	goto L520
L520:
	;
	v1852 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	v1853 = F_lappend(m, v1852, v1711)
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L1
	} else {
		goto L522
	}
L521:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1705
	goto L477
L522:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1705
	*(*int32)(unsafe.Add(mBase, _consts[511])) = v1853
	v1865 = v1711
	goto L484
L523:
	;
	v1883 = v1621 + int32(-16)
	v1884 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1883))) = v1884
	*(*int64)(unsafe.Add(mBase, uint32(v1623)+40)) = v1884
	*(*int64)(unsafe.Add(mBase, uint32(v1623)+32)) = v1884
	*(*int64)(unsafe.Add(mBase, uint32(v1623)+16)) = v1884
	*(*int64)(unsafe.Add(mBase, uint32(v1623)+24)) = int64(34359738372)
	v1895 = *(*int32)(unsafe.Add(mBase, _consts[533]))
	*(*int32)(unsafe.Add(mBase, uint32(v1883))) = v1895
	*(*int64)(unsafe.Add(mBase, uint32(v1623)+8)) = v1884
	v1905 = F_hash_create(m, int32(319797), int32(16), v1621+int32(-56), int32(1064))
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L1
	} else {
		goto L526
	}
L524:
	;
	v1908 = v1879
	goto L525
L525:
	;
	v1914 = F_hash_search(m, v1908, v1621+int32(-4), int32(1), v1621+int32(-56))
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L1
	} else {
		goto L527
	}
L526:
	;
	*(*int32)(unsafe.Add(mBase, _consts[545])) = v1905
	v1908 = v1905
	goto L525
L527:
	;
	v1916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1623)+8)))
	if v1916 == int32(1) {
		goto L476
	} else {
		goto L528
	}
L528:
	;
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+16))
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v1919)))
	*(*int32)(unsafe.Add(mBase, uint32(v1919))) = int32(1)
	if v1920 != 0 {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+16))
	F_s_lock(m, v1923, int32(489965), int32(504), int32(218431))
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L1
	} else {
		goto L532
	}
L530:
	;
	goto L531
L531:
	;
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+16))
	v1930 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1929)+8)) = v1930
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+16))
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1932)+4)) = v1933
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1935))) = v1930
	v1938 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v1865)+12)) = uint16(v1938)
	*(*int32)(unsafe.Add(mBase, uint32(v1914)+4)) = v1865
	goto L477
L532:
	;
	goto L531
L533:
	;
	F_errmsg_internal(m, int32(440232), int32(0))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L1
	} else {
		goto L534
	}
L534:
	;
	F_errfinish(m, int32(489965), int32(501), int32(218431))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L1
	} else {
		goto L535
	}
L535:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L536:
	;
	F_pgstat_report_activity(m, int32(3), int32(0))
	mBase = m.M
	goto L3
L537:
	;
	v2080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+304)))
	if v2080 == int32(1) {
		goto L566
	} else {
		goto L567
	}
L538:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v1998)))
	if v2002 == int32(3) {
		goto L537
	} else {
		goto L541
	}
L539:
	;
	goto L540
L540:
	;
	v2005 = F_pa_find_worker(m, v1979)
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L1
	} else {
		goto L544
	}
L541:
	;
	goto L540
L542:
	;
	v2052 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+1344)) = uint8(v2052)
	v2054 = v1594 - v1593
	*(*int32)(unsafe.Add(mBase, uint32(v22)+320)) = v2054 + int32(1)
	v2059 = *(*int32)(unsafe.Add(mBase, _consts[532]))
	F_BufFileWrite(m, v2059, v22+int32(320), int32(4))
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L1
	} else {
		goto L562
	}
L543:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2026 = F_pa_send_data(m, v2005, v2024, v2025)
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L1
	} else {
		goto L552
	}
L544:
	;
	if v2005 != 0 {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v2007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2005)+12)))
	if v2007 == int32(0) {
		goto L543
	} else {
		goto L548
	}
L546:
	;
	goto L547
L547:
	;
	v2016 = int32(*(*uint8)(unsafe.Add(mBase, _consts[539])))
	if v2016 == int32(0) {
		goto L24
	} else {
		goto L550
	}
L548:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, _consts[544]))
	v2012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+304)))
	F_stream_start_internal(m, v2011, v2012)
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L1
	} else {
		goto L549
	}
L549:
	;
	goto L542
L550:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, _consts[544]))
	v2021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+304)))
	F_stream_start_internal(m, v2020, v2021)
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L1
	} else {
		goto L551
	}
L551:
	;
	goto L536
L552:
	;
	v2028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+304)))
	if v2026 != 0 {
		goto L553
	} else {
		goto L554
	}
L553:
	;
	if v2028&int32(1) == int32(0) {
		goto L556
	} else {
		goto L557
	}
L554:
	;
	goto L555
L555:
	;
	F_pa_switch_to_partial_serialize(m, v2005, (v2028^int32(-1))&int32(1))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L1
	} else {
		goto L561
	}
L556:
	;
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+16))
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v2033)+4))
	F_pa_unlock_stream(m, v2034)
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L1
	} else {
		goto L559
	}
L557:
	;
	goto L558
L558:
	;
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+16))
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v2037)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2037)+20)) = v2038 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[546])) = v2005
	goto L560
L559:
	;
	goto L558
L560:
	;
	goto L536
L561:
	;
	goto L542
L562:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, _consts[532]))
	F_BufFileWrite(m, v2066, v22+int32(1344), int32(1))
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+320)) = v2054
	v2074 = *(*int32)(unsafe.Add(mBase, _consts[532]))
	F_BufFileWrite(m, v2074, v1593+v1595, v2054)
	mBase = m.M
	v2077 = m.ExcPending
	if v2077 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	*(*int32)(unsafe.Add(mBase, _consts[546])) = v2005
	goto L565
L565:
	;
	goto L536
L566:
	;
	v2084 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v2084)+32))
	v2087 = *(*int32)(unsafe.Add(mBase, _consts[547]))
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v2087)+4))
	F_LockApplyTransactionForSession(m, v2085, v2088, int32(1), int32(8))
	mBase = m.M
	v2092 = m.ExcPending
	if v2092 != 0 {
		goto L1
	} else {
		goto L569
	}
L567:
	;
	goto L568
L568:
	;
	*(*int32)(unsafe.Add(mBase, _consts[548])) = int32(0)
	goto L536
L569:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, _consts[547]))
	F_pa_set_xact_state(m, v2094, int32(1))
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		goto L1
	} else {
		goto L570
	}
L570:
	;
	v2099 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v2099)+32))
	F_logicalrep_worker_wakeup(m, v2100)
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	goto L568
L572:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, _consts[544]))
	v2119 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2119)+16)))
	if v2120 == int32(1) {
		goto L575
	} else {
		goto L576
	}
L573:
	;
	v2246 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[544])) = v2246
	*(*uint8)(unsafe.Add(mBase, _consts[539])) = uint8(v2246)
	v2254 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v2255 = *(*int32)(unsafe.Add(mBase, uint32(v2254)+24))
	goto L611
L574:
	;
	v2226 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L1
	} else {
		goto L604
	}
L575:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v2119)))
	if v2123 == int32(3) {
		goto L574
	} else {
		goto L578
	}
L576:
	;
	goto L577
L577:
	;
	v2126 = F_pa_find_worker(m, v2117)
	mBase = m.M
	v2127 = m.ExcPending
	if v2127 != 0 {
		goto L1
	} else {
		goto L581
	}
L578:
	;
	goto L577
L579:
	;
	v2169 = int32(69)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+1344)) = uint8(v2169)
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+320)) = v2171 - v2172 + int32(1)
	v2178 = *(*int32)(unsafe.Add(mBase, _consts[532]))
	F_BufFileWrite(m, v2178, v22+int32(320), int32(4))
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L1
	} else {
		goto L596
	}
L580:
	;
	F_pa_switch_to_partial_serialize(m, v2126, int32(1))
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L1
	} else {
		goto L595
	}
L581:
	;
	if v2126 != 0 {
		goto L582
	} else {
		goto L583
	}
L582:
	;
	v2128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2126)+12)))
	if v2128 != 0 {
		goto L579
	} else {
		goto L585
	}
L583:
	;
	goto L584
L584:
	;
	v2143 = int32(*(*uint8)(unsafe.Add(mBase, _consts[539])))
	if v2143 == int32(0) {
		goto L22
	} else {
		goto L590
	}
L585:
	;
	v2129 = *(*int32)(unsafe.Add(mBase, uint32(v2126)+16))
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v2129)+4))
	F_pa_lock_stream(m, v2130)
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2135 = F_pa_send_data(m, v2126, v2133, v2134)
	mBase = m.M
	v2136 = m.ExcPending
	if v2136 != 0 {
		goto L1
	} else {
		goto L587
	}
L587:
	;
	if v2135 == int32(0) {
		goto L580
	} else {
		goto L588
	}
L588:
	;
	*(*int32)(unsafe.Add(mBase, _consts[546])) = int32(0)
	goto L589
L589:
	;
	goto L573
L590:
	;
	v2147 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v2147)+32))
	v2150 = *(*int32)(unsafe.Add(mBase, _consts[544]))
	F_subxact_info_write(m, v2148, v2150)
	mBase = m.M
	v2152 = m.ExcPending
	if v2152 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	v2154 = *(*int32)(unsafe.Add(mBase, _consts[532]))
	F_BufFileClose(m, v2154)
	mBase = m.M
	v2156 = m.ExcPending
	if v2156 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	*(*int32)(unsafe.Add(mBase, _consts[532])) = int32(0)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2161 = m.ExcPending
	if v2161 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	v2163 = *(*int32)(unsafe.Add(mBase, _consts[531]))
	F_MemoryContextReset(m, v2163)
	mBase = m.M
	v2165 = m.ExcPending
	if v2165 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	goto L573
L595:
	;
	goto L579
L596:
	;
	v2185 = *(*int32)(unsafe.Add(mBase, _consts[532]))
	F_BufFileWrite(m, v2185, v22+int32(1344), int32(1))
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2193 = v2191 - v2192
	*(*int32)(unsafe.Add(mBase, uint32(v22)+320)) = v2193
	v2196 = *(*int32)(unsafe.Add(mBase, _consts[532]))
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_BufFileWrite(m, v2196, v2192+v2197, v2193)
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	v2202 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v2202)+32))
	v2205 = *(*int32)(unsafe.Add(mBase, _consts[544]))
	F_subxact_info_write(m, v2203, v2205)
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	v2209 = *(*int32)(unsafe.Add(mBase, _consts[532]))
	F_BufFileClose(m, v2209)
	mBase = m.M
	v2211 = m.ExcPending
	if v2211 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	*(*int32)(unsafe.Add(mBase, _consts[532])) = int32(0)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L1
	} else {
		goto L601
	}
L601:
	;
	v2218 = *(*int32)(unsafe.Add(mBase, _consts[531]))
	F_MemoryContextReset(m, v2218)
	mBase = m.M
	v2220 = m.ExcPending
	if v2220 != 0 {
		goto L1
	} else {
		goto L602
	}
L602:
	;
	*(*int32)(unsafe.Add(mBase, _consts[546])) = int32(0)
	goto L603
L603:
	;
	goto L573
L604:
	;
	if v2226 != 0 {
		goto L605
	} else {
		goto L606
	}
L605:
	;
	v2229 = *(*int32)(unsafe.Add(mBase, _consts[548]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v2229
	F_errmsg_internal(m, int32(312204), v22+int32(80))
	mBase = m.M
	v2235 = m.ExcPending
	if v2235 != 0 {
		goto L1
	} else {
		goto L608
	}
L606:
	;
	goto L607
L607:
	;
	F_pa_decr_and_wait_stream_block(m)
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L1
	} else {
		goto L610
	}
L608:
	;
	F_errfinish(m, int32(490005), int32(1693), int32(231870))
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L1
	} else {
		goto L609
	}
L609:
	;
	goto L607
L610:
	;
	goto L573
L611:
	;
	if v2255 != v2246 {
		goto L612
	} else {
		goto L613
	}
L612:
	;
	v2258 = int32(4)
	goto L614
L613:
	;
	v2258 = int32(2)
	goto L614
L614:
	;
	v2259 = int32(0)
	F_pgstat_report_activity(m, v2258, v2259)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[535])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _consts[536])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[537])) = v2259
	*(*int32)(unsafe.Add(mBase, _consts[534])) = v2259
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v2259
	goto L3
L615:
	;
	v2283 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2283)+68)))
	v2286 = v22 + int32(1344)
	v2288 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2286))) = v2288
	v2292 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v2293 = m.ExcPending
	if v2293 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2286)+4)) = v2292
	if v2284 != 0 {
		goto L619
	} else {
		goto L620
	}
L618:
	;
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(v22)+1348))
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v2306
	v2309 = *(*int64)(unsafe.Add(mBase, uint32(v22)+1352))
	*(*int64)(unsafe.Add(mBase, _consts[536])) = v2309
	v2311 = *(*int32)(unsafe.Add(mBase, uint32(v22)+1344))
	v2313 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2313)+16)))
	if v2314 == int32(1) {
		goto L627
	} else {
		goto L628
	}
L619:
	;
	v2295 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2296 = m.ExcPending
	if v2296 != 0 {
		goto L1
	} else {
		goto L622
	}
L620:
	;
	goto L621
L621:
	;
	v2301 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2286)+8)) = v2301
	*(*int64)(unsafe.Add(mBase, uint32(v2286)+16)) = v2301
	goto L618
L622:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2286)+8)) = v2295
	v2298 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L1
	} else {
		goto L623
	}
L623:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2286)+16)) = v2298
	goto L618
L624:
	;
	F_pa_unlock_stream(m, v2311)
	mBase = m.M
	v2922 = m.ExcPending
	if v2922 != 0 {
		goto L1
	} else {
		goto L777
	}
L625:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2908 = m.ExcPending
	if v2908 != 0 {
		goto L1
	} else {
		goto L774
	}
L626:
	;
	if v2306 != v2311 {
		goto L723
	} else {
		goto L724
	}
L627:
	;
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(v2313)))
	if v2317 == int32(3) {
		goto L626
	} else {
		goto L630
	}
L628:
	;
	goto L629
L629:
	;
	v2320 = F_pa_find_worker(m, v2311)
	mBase = m.M
	v2321 = m.ExcPending
	if v2321 != 0 {
		goto L1
	} else {
		goto L631
	}
L630:
	;
	goto L629
L631:
	;
	if v2320 != 0 {
		goto L632
	} else {
		goto L633
	}
L632:
	;
	v2322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2320)+12)))
	if v2322 != 0 {
		goto L11
	} else {
		goto L635
	}
L633:
	;
	goto L634
L634:
	;
	v2334 = int32(*(*uint8)(unsafe.Add(mBase, _consts[539])))
	if v2334 != 0 {
		goto L625
	} else {
		goto L640
	}
L635:
	;
	if v2306 != v2311 {
		goto L624
	} else {
		goto L636
	}
L636:
	;
	v2324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2326 = F_pa_send_data(m, v2320, v2324, v2325)
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		goto L1
	} else {
		goto L637
	}
L637:
	;
	if v2326 == int32(0) {
		goto L12
	} else {
		goto L638
	}
L638:
	;
	F_pa_xact_finish(m, v2320, int64(0))
	mBase = m.M
	v2332 = m.ExcPending
	if v2332 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	goto L10
L640:
	;
	if v2306 == v2311 {
		goto L642
	} else {
		goto L643
	}
L641:
	;
	v2706 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2707 = m.ExcPending
	if v2707 != 0 {
		goto L1
	} else {
		goto L719
	}
L642:
	;
	v2337 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(v2337)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v2338
	*(*int32)(unsafe.Add(mBase, uint32(v22)+148)) = v2311
	v2347 = F_pg_snprintf(m, v22+int32(320), int32(1024), int32(167885), v22+int32(144))
	mBase = m.M
	v2348 = m.ExcPending
	if v2348 != 0 {
		goto L1
	} else {
		goto L645
	}
L643:
	;
	goto L644
L644:
	;
	v2376 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	if v2376 < int32(0) {
		goto L650
	} else {
		goto L651
	}
L645:
	;
	v2350 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(v2350)+60))
	F_BufFileDeleteFileSet(m, v2351, v22+int32(320), int32(0))
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L1
	} else {
		goto L646
	}
L646:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+132)) = v2311
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v2338
	v2365 = F_pg_snprintf(m, v22+int32(320), int32(1024), int32(124348), v22+int32(128))
	mBase = m.M
	v2366 = m.ExcPending
	if v2366 != 0 {
		goto L1
	} else {
		goto L647
	}
L647:
	;
	v2368 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2369 = *(*int32)(unsafe.Add(mBase, uint32(v2368)+60))
	F_BufFileDeleteFileSet(m, v2369, v22+int32(320), int32(1))
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L1
	} else {
		goto L648
	}
L648:
	;
	goto L641
L649:
	;
	v2383 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(v2383)+20))
	goto L653
L650:
	;
	v2380 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[291])) = v2380
	goto L652
L651:
	;
	goto L652
L652:
	;
	goto L649
L653:
	;
	if base.B2i32(v2384 == int32(2)) == int32(0) {
		goto L654
	} else {
		goto L655
	}
L654:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2390 = m.ExcPending
	if v2390 != 0 {
		goto L1
	} else {
		goto L657
	}
L655:
	;
	goto L656
L656:
	;
	v2393 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L1
	} else {
		goto L659
	}
L657:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L1
	} else {
		goto L658
	}
L658:
	;
	goto L656
L659:
	;
	F_PushActiveSnapshot(m, v2393)
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L1
	} else {
		goto L660
	}
L660:
	;
	v2399 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2399
	v2402 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v2402)+32))
	F_subxact_info_read(m, v2403, v2311)
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L1
	} else {
		goto L661
	}
L661:
	;
	v2407 = *(*int32)(unsafe.Add(mBase, _consts[549]))
	v2409 = int64(*(*uint32)(unsafe.Add(mBase, _consts[550])))
	v2427 = v2409
	goto L664
L662:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2680 = m.ExcPending
	if v2680 != 0 {
		goto L1
	} else {
		goto L716
	}
L663:
	;
	if v2407 != 0 {
		goto L712
	} else {
		goto L713
	}
L664:
	;
	if v2427 <= int64(0) {
		goto L663
	} else {
		goto L666
	}
L665:
	;
	v2440 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(v2440)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = v2441
	*(*int32)(unsafe.Add(mBase, uint32(v22)+164)) = v2311
	v2450 = F_pg_snprintf(m, v22+int32(320), int32(1024), int32(167885), v22+int32(160))
	mBase = m.M
	v2451 = m.ExcPending
	if v2451 != 0 {
		goto L1
	} else {
		goto L668
	}
L666:
	;
	v2432 = v2427 - int64(1)
	v2433 = base.I32_wrap_i64(v2432)
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(v2407+v2433<<(uint(int32(4))%32))))
	if v2437 != v2306 {
		v2427 = v2432
		goto L664
	} else {
		goto L667
	}
L667:
	;
	goto L665
L668:
	;
	v2453 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v2453)+60))
	v2459 = F_BufFileOpenFileSet(m, v2454, v22+int32(320), int32(2), int32(0))
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L1
	} else {
		goto L669
	}
L669:
	;
	v2462 = *(*int32)(unsafe.Add(mBase, _consts[549]))
	v2465 = v2462 + v2433<<(uint(int32(4))%32)
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(v2465)+4))
	v2467 = *(*int64)(unsafe.Add(mBase, uint32(v2465)+8))
	v2468 = m.G0
	v2470 = v2468 - int32(1072)
	m.G0 = v2470
	v2472 = *(*int32)(unsafe.Add(mBase, uint32(v2459)))
	v2474 = v2472 - int32(1)
	if v2474 < v2466 {
		goto L674
	} else {
		goto L675
	}
L670:
	;
	F_BufFileClose(m, v2459)
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L1
	} else {
		goto L710
	}
L671:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2618 = m.ExcPending
	if v2618 != 0 {
		goto L1
	} else {
		goto L705
	}
L672:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2601 = m.ExcPending
	if v2601 != 0 {
		goto L1
	} else {
		goto L701
	}
L673:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2459))) = v2558
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(v2459)+24))
	if v2551 != v2570 {
		goto L691
	} else {
		goto L692
	}
L674:
	;
	v2476 = *(*int64)(unsafe.Add(mBase, uint32(v2459)+32))
	v2551 = v2466
	v2558 = v2472
	v2567 = v2476
	goto L673
L675:
	;
	goto L676
L676:
	;
	v2478 = v2466
	v2484 = v2474
	v2485 = v2472
	goto L677
L677:
	;
	if v2484 == int32(0) {
		goto L680
	} else {
		goto L681
	}
L678:
	;
	v2551 = v2544
	v2558 = v2545
	v2567 = v2546
	goto L673
L679:
	;
	v2548 = v2484 - int32(1)
	if v2466 <= v2548 {
		v2478 = v2544
		v2484 = v2548
		v2485 = v2545
		goto L677
	} else {
		goto L689
	}
L680:
	;
	v2534 = *(*int32)(unsafe.Add(mBase, uint32(v2459)+4))
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(v2534+v2484<<(uint(int32(2))%32))))
	v2540 = F_FileTruncate(m, v2538, v2467, int32(167772167))
	mBase = m.M
	v2541 = m.ExcPending
	if v2541 != 0 {
		goto L1
	} else {
		goto L687
	}
L681:
	;
	if base.B2i32(v2467 == int64(0))|base.B2i32(v2466 != v2484) == int32(0) {
		goto L680
	} else {
		goto L682
	}
L682:
	;
	v2504 = *(*int32)(unsafe.Add(mBase, uint32(v2459)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2470)+16)) = v2504
	*(*int32)(unsafe.Add(mBase, uint32(v2470)+20)) = v2484
	v2513 = F_pg_snprintf(m, v2470+int32(48), int32(1024), int32(461698), v2470+int32(16))
	mBase = m.M
	v2514 = m.ExcPending
	if v2514 != 0 {
		goto L1
	} else {
		goto L683
	}
L683:
	;
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(v2459)+4))
	v2519 = *(*int32)(unsafe.Add(mBase, uint32(v2515+v2484<<(uint(int32(2))%32))))
	F_FileClose(m, v2519)
	mBase = m.M
	v2521 = m.ExcPending
	if v2521 != 0 {
		goto L1
	} else {
		goto L684
	}
L684:
	;
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(v2459)+12))
	v2525 = F_FileSetDelete(m, v2522, v2470+int32(48))
	mBase = m.M
	v2526 = m.ExcPending
	if v2526 != 0 {
		goto L1
	} else {
		goto L685
	}
L685:
	;
	if v2525 == int32(0) {
		goto L672
	} else {
		goto L686
	}
L686:
	;
	v2544 = v2478 - base.B2i32(v2466 == v2484)
	v2545 = v2485 - int32(1)
	v2546 = int64(1073741824)
	goto L679
L687:
	;
	if v2540 < int32(0) {
		goto L671
	} else {
		goto L688
	}
L688:
	;
	v2544 = v2478
	v2545 = v2485
	v2546 = v2467
	goto L679
L689:
	;
	goto L678
L690:
	;
	m.G0 = v2470 + int32(1072)
	goto L670
L691:
	;
	if v2570 <= v2551 {
		goto L690
	} else {
		goto L700
	}
L692:
	;
	v2572 = *(*int64)(unsafe.Add(mBase, uint32(v2459)+32))
	if v2572 <= v2567 {
		goto L693
	} else {
		goto L694
	}
L693:
	;
	v2574 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2459)+44)))
	if v2572+v2574 < v2567 {
		goto L691
	} else {
		goto L696
	}
L694:
	;
	goto L695
L695:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2459)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2459)+32)) = v2567
	goto L690
L696:
	;
	v2578 = base.I32_wrap_i64(v2567 - v2572)
	v2579 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2459)+40)))
	if v2567 <= v2572+v2579 {
		goto L697
	} else {
		goto L698
	}
L697:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2459)+40)) = v2578
	goto L699
L698:
	;
	goto L699
L699:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2459)+44)) = v2578
	goto L690
L700:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2459)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2459)+32)) = v2567
	*(*int32)(unsafe.Add(mBase, uint32(v2459)+24)) = v2551
	goto L690
L701:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2603 = m.ExcPending
	if v2603 != 0 {
		goto L1
	} else {
		goto L702
	}
L702:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2470))) = v2470 + int32(48)
	F_errmsg(m, int32(293941), v2470)
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		goto L1
	} else {
		goto L703
	}
L703:
	;
	F_errfinish(m, int32(493502), int32(952), int32(107555))
	mBase = m.M
	v2614 = m.ExcPending
	if v2614 != 0 {
		goto L1
	} else {
		goto L704
	}
L704:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L705:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2620 = m.ExcPending
	if v2620 != 0 {
		goto L1
	} else {
		goto L706
	}
L706:
	;
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v2459)+4))
	v2625 = *(*int32)(unsafe.Add(mBase, uint32(v2621+v2484<<(uint(int32(2))%32))))
	v2627 = *(*int32)(unsafe.Add(mBase, _consts[551]))
	v2631 = *(*int32)(unsafe.Add(mBase, uint32(v2627+v2625*int32(48))+32))
	goto L707
L707:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2470)+32)) = v2631
	F_errmsg(m, int32(296215), v2470+int32(32))
	mBase = m.M
	v2637 = m.ExcPending
	if v2637 != 0 {
		goto L1
	} else {
		goto L708
	}
L708:
	;
	F_errfinish(m, int32(493502), int32(970), int32(107555))
	mBase = m.M
	v2642 = m.ExcPending
	if v2642 != 0 {
		goto L1
	} else {
		goto L709
	}
L709:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L710:
	;
	*(*int32)(unsafe.Add(mBase, _consts[550])) = v2433
	v2648 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(v2648)+32))
	F_subxact_info_write(m, v2649, v2311)
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
		goto L1
	} else {
		goto L711
	}
L711:
	;
	goto L662
L712:
	;
	F_pfree(m, v2407)
	mBase = m.M
	v2653 = m.ExcPending
	if v2653 != 0 {
		goto L1
	} else {
		goto L715
	}
L713:
	;
	goto L714
L714:
	;
	v2655 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[552])) = v2655
	*(*int64)(unsafe.Add(mBase, _consts[550])) = v2655
	goto L662
L715:
	;
	goto L714
L716:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2682 = m.ExcPending
	if v2682 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2684 = m.ExcPending
	if v2684 != 0 {
		goto L1
	} else {
		goto L718
	}
L718:
	;
	goto L641
L719:
	;
	if v2706 == int32(0) {
		goto L10
	} else {
		goto L720
	}
L720:
	;
	F_errmsg_internal(m, int32(423975), int32(0))
	mBase = m.M
	v2713 = m.ExcPending
	if v2713 != 0 {
		goto L1
	} else {
		goto L721
	}
L721:
	;
	F_errfinish(m, int32(490005), int32(1868), int32(80633))
	mBase = m.M
	v2718 = m.ExcPending
	if v2718 != 0 {
		goto L1
	} else {
		goto L722
	}
L722:
	;
	goto L10
L723:
	;
	v2730 = m.G0
	v2732 = v2730 - int32(96)
	m.G0 = v2732
	v2735 = v22 + int32(1344)
	v2736 = *(*int32)(unsafe.Add(mBase, uint32(v2735)+4))
	v2737 = *(*int32)(unsafe.Add(mBase, uint32(v2735)))
	v2739 = *(*int64)(unsafe.Add(mBase, uint32(v2735)+8))
	*(*int64)(unsafe.Add(mBase, _consts[110])) = v2739
	v2742 = *(*int64)(unsafe.Add(mBase, uint32(v2735)+16))
	*(*int64)(unsafe.Add(mBase, _consts[111])) = v2742
	if v2736 == v2737 {
		goto L728
	} else {
		goto L729
	}
L724:
	;
	v2721 = *(*int32)(unsafe.Add(mBase, _consts[532]))
	if v2721 == int32(0) {
		goto L723
	} else {
		goto L725
	}
L725:
	;
	F_BufFileClose(m, v2721)
	mBase = m.M
	v2725 = m.ExcPending
	if v2725 != 0 {
		goto L1
	} else {
		goto L726
	}
L726:
	;
	*(*int32)(unsafe.Add(mBase, _consts[532])) = int32(0)
	goto L723
L727:
	;
	m.G0 = v2732 + int32(96)
	if v2306 != v2311 {
		goto L766
	} else {
		goto L767
	}
L728:
	;
	v2746 = *(*int32)(unsafe.Add(mBase, _consts[547]))
	v2747 = *(*int32)(unsafe.Add(mBase, uint32(v2746)))
	*(*int32)(unsafe.Add(mBase, uint32(v2746))) = int32(1)
	if v2747 != 0 {
		goto L731
	} else {
		goto L732
	}
L729:
	;
	goto L730
L730:
	;
	v2785 = *(*int32)(unsafe.Add(mBase, _consts[529]))
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(v2785)))
	*(*int32)(unsafe.Add(mBase, uint32(v2732)+16)) = v2786
	*(*int32)(unsafe.Add(mBase, uint32(v2732)+20)) = v2736
	v2795 = F_pg_snprintf(m, v2732+int32(32), int32(64), int32(38204), v2732+int32(16))
	mBase = m.M
	v2796 = m.ExcPending
	if v2796 != 0 {
		goto L1
	} else {
		goto L743
	}
L731:
	;
	F_s_lock(m, v2746, int32(489965), int32(1317), int32(347726))
	mBase = m.M
	v2754 = m.ExcPending
	if v2754 != 0 {
		goto L1
	} else {
		goto L734
	}
L732:
	;
	goto L733
L733:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2746))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2746)+8)) = int32(2)
	v2760 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2761 = *(*int32)(unsafe.Add(mBase, uint32(v2760)+32))
	F_UnlockApplyTransactionForSession(m, v2761, v2737, int32(1), int32(8))
	mBase = m.M
	v2765 = m.ExcPending
	if v2765 != 0 {
		goto L1
	} else {
		goto L735
	}
L734:
	;
	goto L733
L735:
	;
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v2767 = m.ExcPending
	if v2767 != 0 {
		goto L1
	} else {
		goto L736
	}
L736:
	;
	v2769 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v2770 = *(*int32)(unsafe.Add(mBase, uint32(v2769)+24))
	goto L737
L737:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2770) {
		goto L738
	} else {
		goto L739
	}
L738:
	;
	v2774 = F_EndTransactionBlock(m, int32(0))
	mBase = m.M
	v2775 = m.ExcPending
	if v2775 != 0 {
		goto L1
	} else {
		goto L741
	}
L739:
	;
	goto L740
L740:
	;
	v2779 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[553])) = v2779
	F_pgstat_report_activity(m, int32(2), v2779)
	mBase = m.M
	goto L727
L741:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2777 = m.ExcPending
	if v2777 != 0 {
		goto L1
	} else {
		goto L742
	}
L742:
	;
	goto L740
L743:
	;
	v2799 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2800 = m.ExcPending
	if v2800 != 0 {
		goto L1
	} else {
		goto L744
	}
L744:
	;
	if v2799 != 0 {
		goto L745
	} else {
		goto L746
	}
L745:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2732))) = v2732 + int32(32)
	F_errmsg_internal(m, int32(218540), v2732)
	mBase = m.M
	v2806 = m.ExcPending
	if v2806 != 0 {
		goto L1
	} else {
		goto L748
	}
L746:
	;
	goto L747
L747:
	;
	v2814 = *(*int32)(unsafe.Add(mBase, _consts[553]))
	if v2814 != 0 {
		goto L750
	} else {
		goto L751
	}
L748:
	;
	F_errfinish(m, int32(489965), int32(1475), int32(80659))
	mBase = m.M
	v2811 = m.ExcPending
	if v2811 != 0 {
		goto L1
	} else {
		goto L749
	}
L749:
	;
	goto L747
L750:
	;
	v2815 = *(*int32)(unsafe.Add(mBase, uint32(v2814)+4))
	v2816 = v2815
	goto L752
L751:
	;
	v2816 = int32(0)
	goto L752
L752:
	;
	v2817 = v2816
	goto L753
L753:
	;
	v2837 = v2817 - int32(1)
	if v2837 < int32(0) {
		goto L727
	} else {
		goto L755
	}
L754:
	;
	F_RollbackToSavepoint(m, v2732+int32(32))
	mBase = m.M
	v2849 = m.ExcPending
	if v2849 != 0 {
		goto L1
	} else {
		goto L757
	}
L755:
	;
	v2840 = *(*int32)(unsafe.Add(mBase, uint32(v2814)+12))
	v2844 = *(*int32)(unsafe.Add(mBase, uint32(v2840+v2837<<(uint(int32(2))%32))))
	if v2844 != v2736 {
		v2817 = v2837
		goto L753
	} else {
		goto L756
	}
L756:
	;
	goto L754
L757:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2851 = m.ExcPending
	if v2851 != 0 {
		goto L1
	} else {
		goto L758
	}
L758:
	;
	v2852 = int32(4386076)
	v2854 = *(*int32)(unsafe.Add(mBase, _consts[553]))
	v2855 = int32(0)
	if v2854 == v2855 {
		v2863 = v2855
		goto L760
	} else {
		goto L761
	}
L759:
	;
	*(*int32)(unsafe.Add(mBase, _consts[553])) = v2863
	goto L727
L760:
	;
	goto L759
L761:
	;
	if v2837 <= int32(0) {
		v2863 = v2855
		goto L760
	} else {
		goto L762
	}
L762:
	;
	v2860 = *(*int32)(unsafe.Add(mBase, uint32(v2854)+4))
	if v2837 < v2860 {
		goto L763
	} else {
		goto L764
	}
L763:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2854)+4)) = v2837
	goto L765
L764:
	;
	goto L765
L765:
	;
	v2863 = v2854
	goto L760
L766:
	;
	F_pa_decr_and_wait_stream_block(m)
	mBase = m.M
	v2889 = m.ExcPending
	if v2889 != 0 {
		goto L1
	} else {
		goto L769
	}
L767:
	;
	goto L768
L768:
	;
	v2892 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L1
	} else {
		goto L770
	}
L769:
	;
	goto L768
L770:
	;
	if v2892 == int32(0) {
		goto L10
	} else {
		goto L771
	}
L771:
	;
	F_errmsg_internal(m, int32(423975), int32(0))
	mBase = m.M
	v2899 = m.ExcPending
	if v2899 != 0 {
		goto L1
	} else {
		goto L772
	}
L772:
	;
	F_errfinish(m, int32(490005), int32(1971), int32(80633))
	mBase = m.M
	v2904 = m.ExcPending
	if v2904 != 0 {
		goto L1
	} else {
		goto L773
	}
L773:
	;
	goto L10
L774:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = int32(1)
	F_errmsg_internal(m, int32(477647), v22+int32(112))
	mBase = m.M
	v2915 = m.ExcPending
	if v2915 != 0 {
		goto L1
	} else {
		goto L775
	}
L775:
	;
	F_errfinish(m, int32(490005), int32(1975), int32(80633))
	mBase = m.M
	v2920 = m.ExcPending
	if v2920 != 0 {
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
	v2923 = *(*int32)(unsafe.Add(mBase, uint32(v2320)+16))
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v2923)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2923)+20)) = v2924 + int32(1)
	F_pa_lock_stream(m, v2311)
	mBase = m.M
	v2929 = m.ExcPending
	if v2929 != 0 {
		goto L1
	} else {
		goto L778
	}
L778:
	;
	v2930 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2931 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2932 = F_pa_send_data(m, v2320, v2930, v2931)
	mBase = m.M
	v2933 = m.ExcPending
	if v2933 != 0 {
		goto L1
	} else {
		goto L779
	}
L779:
	;
	if v2932 == int32(0) {
		goto L12
	} else {
		goto L780
	}
L780:
	;
	goto L10
L781:
	;
	v2945 = v22 + int32(320)
	v2946 = m.G0
	v2948 = v2946 - int32(16)
	m.G0 = v2948
	v2951 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v2952 = m.ExcPending
	if v2952 != 0 {
		goto L1
	} else {
		goto L782
	}
L782:
	;
	v2953 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v2954 = m.ExcPending
	if v2954 != 0 {
		goto L1
	} else {
		goto L783
	}
L783:
	;
	v2956 = v2953 & int32(255)
	if v2956 != 0 {
		goto L784
	} else {
		goto L785
	}
L784:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2960 = m.ExcPending
	if v2960 != 0 {
		goto L1
	} else {
		goto L787
	}
L785:
	;
	goto L786
L786:
	;
	v2970 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2971 = m.ExcPending
	if v2971 != 0 {
		goto L1
	} else {
		goto L790
	}
L787:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2948))) = v2956
	F_errmsg_internal(m, int32(399990), v2948)
	mBase = m.M
	v2964 = m.ExcPending
	if v2964 != 0 {
		goto L1
	} else {
		goto L788
	}
L788:
	;
	F_errfinish(m, int32(490650), int32(1140), int32(99718))
	mBase = m.M
	v2969 = m.ExcPending
	if v2969 != 0 {
		goto L1
	} else {
		goto L789
	}
L789:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L790:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2945))) = v2970
	v2973 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2974 = m.ExcPending
	if v2974 != 0 {
		goto L1
	} else {
		goto L791
	}
L791:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2945)+8)) = v2973
	v2976 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2977 = m.ExcPending
	if v2977 != 0 {
		goto L1
	} else {
		goto L792
	}
L792:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2945)+16)) = v2976
	m.G0 = v2948 + int32(16)
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v2951
	v2985 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	*(*int64)(unsafe.Add(mBase, _consts[536])) = v2985
	v2988 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v2989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2988)+16)))
	if v2989 == int32(1) {
		goto L795
	} else {
		goto L796
	}
L793:
	;
	v3094 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	F_process_syncing_tables(m, v3094)
	mBase = m.M
	v3096 = m.ExcPending
	if v3096 != 0 {
		goto L1
	} else {
		goto L833
	}
L794:
	;
	v3054 = *(*int32)(unsafe.Add(mBase, _consts[532]))
	if v3054 != 0 {
		goto L821
	} else {
		goto L822
	}
L795:
	;
	v2992 = *(*int32)(unsafe.Add(mBase, uint32(v2988)))
	if v2992 == int32(3) {
		goto L794
	} else {
		goto L798
	}
L796:
	;
	goto L797
L797:
	;
	v2995 = F_pa_find_worker(m, v2951)
	mBase = m.M
	v2996 = m.ExcPending
	if v2996 != 0 {
		goto L1
	} else {
		goto L801
	}
L798:
	;
	goto L797
L799:
	;
	F_stream_open_and_write_change(m, v2951, int32(99), v22+int32(1344))
	mBase = m.M
	v3046 = m.ExcPending
	if v3046 != 0 {
		goto L1
	} else {
		goto L818
	}
L800:
	;
	F_pa_switch_to_partial_serialize(m, v2995, int32(1))
	mBase = m.M
	v3041 = m.ExcPending
	if v3041 != 0 {
		goto L1
	} else {
		goto L817
	}
L801:
	;
	if v2995 != 0 {
		goto L802
	} else {
		goto L803
	}
L802:
	;
	v2997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2995)+12)))
	if v2997 != 0 {
		goto L799
	} else {
		goto L805
	}
L803:
	;
	goto L804
L804:
	;
	v3008 = int32(*(*uint8)(unsafe.Add(mBase, _consts[539])))
	if v3008 != 0 {
		goto L19
	} else {
		goto L809
	}
L805:
	;
	v2998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2999 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3000 = F_pa_send_data(m, v2995, v2998, v2999)
	mBase = m.M
	v3001 = m.ExcPending
	if v3001 != 0 {
		goto L1
	} else {
		goto L806
	}
L806:
	;
	if v3000 == int32(0) {
		goto L800
	} else {
		goto L807
	}
L807:
	;
	v3004 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	F_pa_xact_finish(m, v2995, v3004)
	mBase = m.M
	v3006 = m.ExcPending
	if v3006 != 0 {
		goto L1
	} else {
		goto L808
	}
L808:
	;
	goto L793
L809:
	;
	v3010 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v3011 = *(*int32)(unsafe.Add(mBase, uint32(v3010)+60))
	v3012 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	F_apply_spooled_messages(m, v3011, v2951, v3012)
	mBase = m.M
	v3014 = m.ExcPending
	if v3014 != 0 {
		goto L1
	} else {
		goto L810
	}
L810:
	;
	F_apply_handle_commit_internal(m, v22+int32(320))
	mBase = m.M
	v3018 = m.ExcPending
	if v3018 != 0 {
		goto L1
	} else {
		goto L811
	}
L811:
	;
	v3020 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v3021 = *(*int32)(unsafe.Add(mBase, uint32(v3020)+32))
	F_stream_cleanup_files(m, v3021, v2951)
	mBase = m.M
	v3023 = m.ExcPending
	if v3023 != 0 {
		goto L1
	} else {
		goto L812
	}
L812:
	;
	v3026 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3027 = m.ExcPending
	if v3027 != 0 {
		goto L1
	} else {
		goto L813
	}
L813:
	;
	if v3026 == int32(0) {
		goto L793
	} else {
		goto L814
	}
L814:
	;
	F_errmsg_internal(m, int32(424020), int32(0))
	mBase = m.M
	v3033 = m.ExcPending
	if v3033 != 0 {
		goto L1
	} else {
		goto L815
	}
L815:
	;
	F_errfinish(m, int32(490005), int32(2184), int32(99691))
	mBase = m.M
	v3038 = m.ExcPending
	if v3038 != 0 {
		goto L1
	} else {
		goto L816
	}
L816:
	;
	goto L793
L817:
	;
	goto L799
L818:
	;
	v3047 = *(*int32)(unsafe.Add(mBase, uint32(v2995)+16))
	F_pa_set_fileset_state(m, v3047)
	mBase = m.M
	v3049 = m.ExcPending
	if v3049 != 0 {
		goto L1
	} else {
		goto L819
	}
L819:
	;
	v3050 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	F_pa_xact_finish(m, v2995, v3050)
	mBase = m.M
	v3052 = m.ExcPending
	if v3052 != 0 {
		goto L1
	} else {
		goto L820
	}
L820:
	;
	goto L793
L821:
	;
	F_BufFileClose(m, v3054)
	mBase = m.M
	v3056 = m.ExcPending
	if v3056 != 0 {
		goto L1
	} else {
		goto L824
	}
L822:
	;
	goto L823
L823:
	;
	F_apply_handle_commit_internal(m, v22+int32(320))
	mBase = m.M
	v3063 = m.ExcPending
	if v3063 != 0 {
		goto L1
	} else {
		goto L825
	}
L824:
	;
	*(*int32)(unsafe.Add(mBase, _consts[532])) = int32(0)
	goto L823
L825:
	;
	v3065 = *(*int32)(unsafe.Add(mBase, _consts[547]))
	v3067 = *(*int64)(unsafe.Add(mBase, _consts[554]))
	*(*int64)(unsafe.Add(mBase, uint32(v3065)+24)) = v3067
	F_pa_set_xact_state(m, v3065, int32(2))
	mBase = m.M
	v3071 = m.ExcPending
	if v3071 != 0 {
		goto L1
	} else {
		goto L826
	}
L826:
	;
	F_pa_unlock_transaction(m, v2951)
	mBase = m.M
	v3073 = m.ExcPending
	if v3073 != 0 {
		goto L1
	} else {
		goto L827
	}
L827:
	;
	*(*int32)(unsafe.Add(mBase, _consts[553])) = int32(0)
	goto L828
L828:
	;
	v3079 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3080 = m.ExcPending
	if v3080 != 0 {
		goto L1
	} else {
		goto L829
	}
L829:
	;
	if v3079 == int32(0) {
		goto L793
	} else {
		goto L830
	}
L830:
	;
	F_errmsg_internal(m, int32(424020), int32(0))
	mBase = m.M
	v3086 = m.ExcPending
	if v3086 != 0 {
		goto L1
	} else {
		goto L831
	}
L831:
	;
	F_errfinish(m, int32(490005), int32(2238), int32(99691))
	mBase = m.M
	v3091 = m.ExcPending
	if v3091 != 0 {
		goto L1
	} else {
		goto L832
	}
L832:
	;
	goto L793
L833:
	;
	v3098 = int32(0)
	F_pgstat_report_activity(m, int32(2), v3098)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[535])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _consts[536])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[537])) = v3098
	*(*int32)(unsafe.Add(mBase, _consts[534])) = v3098
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v3098
	goto L3
L834:
	;
	v3120 = *(*int32)(unsafe.Add(mBase, uint32(v3116)))
	if v3120 == int32(1) {
		goto L18
	} else {
		goto L837
	}
L835:
	;
	goto L836
L836:
	;
	v3124 = v22 + int32(320)
	v3125 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3126 = m.ExcPending
	if v3126 != 0 {
		goto L1
	} else {
		goto L838
	}
L837:
	;
	goto L836
L838:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3124))) = v3125
	if v3125 != int64(0) {
		goto L841
	} else {
		goto L842
	}
L839:
	;
	v3288 = *(*int32)(unsafe.Add(mBase, uint32(v22)+344))
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v3288
	v3291 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	*(*int64)(unsafe.Add(mBase, _consts[536])) = v3291
	*(*int64)(unsafe.Add(mBase, _consts[541])) = v3291
	v3296 = *(*int32)(unsafe.Add(mBase, _consts[529]))
	v3297 = *(*int64)(unsafe.Add(mBase, uint32(v3296)+8))
	if v3297 == int64(0) {
		goto L887
	} else {
		goto L888
	}
L840:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3277 = m.ExcPending
	if v3277 != 0 {
		goto L1
	} else {
		goto L884
	}
L841:
	;
	v3130 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3131 = m.ExcPending
	if v3131 != 0 {
		goto L1
	} else {
		goto L844
	}
L842:
	;
	goto L843
L843:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3264 = m.ExcPending
	if v3264 != 0 {
		goto L1
	} else {
		goto L881
	}
L844:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3124)+8)) = v3130
	if v3130 == int64(0) {
		goto L840
	} else {
		goto L845
	}
L845:
	;
	v3135 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3136 = m.ExcPending
	if v3136 != 0 {
		goto L1
	} else {
		goto L846
	}
L846:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3124)+16)) = v3135
	v3139 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v3140 = m.ExcPending
	if v3140 != 0 {
		goto L1
	} else {
		goto L847
	}
L847:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3124)+24)) = v3139
	v3143 = v22 + int32(348)
	v3144 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v3145 = m.ExcPending
	if v3145 != 0 {
		goto L1
	} else {
		goto L848
	}
L848:
	;
	goto L852
L849:
	;
	goto L839
L850:
	;
	v3258 = F_strlen(m, v3247)
	mBase = m.M
	goto L849
L852:
	;
	goto L853
L853:
	;
	v3152 = int32(199)
	if (v3143^v3144)&int32(3) != 0 {
		goto L857
	} else {
		goto L858
	}
L854:
	;
	v3251 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3248))) = uint8(v3251)
	goto L850
L855:
	;
	v3232 = v3227
	v3233 = v3228
	v3234 = v3229
	goto L877
L856:
	;
	if v3222 == int32(0) {
		v3247 = v3220
		v3248 = v3221
		goto L854
	} else {
		goto L876
	}
L857:
	;
	v3220 = v3144
	v3221 = v3143
	v3222 = v3152
	goto L856
L858:
	;
	goto L859
L859:
	;
	if v3144&int32(3) == int32(0) {
		goto L861
	} else {
		goto L862
	}
L860:
	;
	if v3189 == int32(0) {
		v3247 = v3186
		v3248 = v3187
		goto L854
	} else {
		goto L869
	}
L861:
	;
	v3186 = v3144
	v3187 = v3143
	v3188 = v3152
	v3189 = int32(1)
	goto L860
L862:
	;
	goto L863
L863:
	;
	v3165 = v3144
	v3166 = v3143
	v3167 = v3152
	goto L864
L864:
	;
	v3169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3165))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3166))) = uint8(v3169)
	if v3169 == int32(0) {
		v3227 = v3165
		v3228 = v3166
		v3229 = v3167
		goto L855
	} else {
		goto L866
	}
L865:
	;
	v3186 = v3180
	v3187 = v3174
	v3188 = v3176
	v3189 = v3178
	goto L860
L866:
	;
	v3173 = int32(1)
	v3174 = v3166 + v3173
	v3176 = v3167 - v3173
	v3177 = int32(0)
	v3178 = base.B2i32(v3176 != v3177)
	v3180 = v3165 + v3173
	if v3180&int32(3) == v3177 {
		v3186 = v3180
		v3187 = v3174
		v3188 = v3176
		v3189 = v3178
		goto L860
	} else {
		goto L867
	}
L867:
	;
	if v3176 != 0 {
		v3165 = v3180
		v3166 = v3174
		v3167 = v3176
		goto L864
	} else {
		goto L868
	}
L868:
	;
	goto L865
L869:
	;
	v3192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3186))))
	if v3192 == int32(0) {
		v3220 = v3186
		v3221 = v3187
		v3222 = v3188
		goto L856
	} else {
		goto L870
	}
L870:
	;
	if base.Ui32(v3188) < base.Ui32(int32(4)) {
		v3220 = v3186
		v3221 = v3187
		v3222 = v3188
		goto L856
	} else {
		goto L871
	}
L871:
	;
	v3198 = v3186
	v3199 = v3187
	v3200 = v3188
	goto L872
L872:
	;
	v3203 = *(*int32)(unsafe.Add(mBase, uint32(v3198)))
	v3206 = int32(-2139062144)
	if (int32(16843008)-v3203|v3203)&v3206 != v3206 {
		v3227 = v3198
		v3228 = v3199
		v3229 = v3200
		goto L855
	} else {
		goto L874
	}
L873:
	;
	v3220 = v3214
	v3221 = v3212
	v3222 = v3216
	goto L856
L874:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3199))) = v3203
	v3211 = int32(4)
	v3212 = v3199 + v3211
	v3214 = v3198 + v3211
	v3216 = v3200 - v3211
	if base.Ui32(int32(3)) < base.Ui32(v3216) {
		v3198 = v3214
		v3199 = v3212
		v3200 = v3216
		goto L872
	} else {
		goto L875
	}
L875:
	;
	goto L873
L876:
	;
	v3227 = v3220
	v3228 = v3221
	v3229 = v3222
	goto L855
L877:
	;
	v3236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3232))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3233))) = uint8(v3236)
	if v3236 == int32(0) {
		v3247 = v3232
		v3248 = v3233
		goto L854
	} else {
		goto L879
	}
L878:
	;
	v3247 = v3243
	v3248 = v3241
	goto L854
L879:
	;
	v3240 = int32(1)
	v3241 = v3233 + v3240
	v3243 = v3232 + v3240
	v3245 = v3234 - v3240
	if v3245 != 0 {
		v3232 = v3243
		v3233 = v3241
		v3234 = v3245
		goto L877
	} else {
		goto L880
	}
L880:
	;
	goto L878
L881:
	;
	F_errmsg_internal(m, int32(400352), int32(0))
	mBase = m.M
	v3268 = m.ExcPending
	if v3268 != 0 {
		goto L1
	} else {
		goto L882
	}
L882:
	;
	F_errfinish(m, int32(490650), int32(139), int32(361318))
	mBase = m.M
	v3273 = m.ExcPending
	if v3273 != 0 {
		goto L1
	} else {
		goto L883
	}
L883:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L884:
	;
	F_errmsg_internal(m, int32(400397), int32(0))
	mBase = m.M
	v3281 = m.ExcPending
	if v3281 != 0 {
		goto L1
	} else {
		goto L885
	}
L885:
	;
	F_errfinish(m, int32(490650), int32(142), int32(361318))
	mBase = m.M
	v3286 = m.ExcPending
	if v3286 != 0 {
		goto L1
	} else {
		goto L886
	}
L886:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L887:
	;
	v3327 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[542])) = uint8(v3327)
	F_pgstat_report_activity(m, int32(3), int32(0))
	mBase = m.M
	goto L3
L888:
	;
	if v3291 != v3297 {
		goto L887
	} else {
		goto L889
	}
L889:
	;
	*(*int64)(unsafe.Add(mBase, _consts[540])) = v3291
	v3305 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3306 = m.ExcPending
	if v3306 != 0 {
		goto L1
	} else {
		goto L890
	}
L890:
	;
	if v3305 == int32(0) {
		goto L887
	} else {
		goto L891
	}
L891:
	;
	v3310 = *(*int64)(unsafe.Add(mBase, _consts[540]))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+196)) = uint32(v3310)
	v3313 = int64(base.Ui64(v3310) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+192)) = uint32(v3313)
	F_errmsg(m, int32(509741), v22+int32(192))
	mBase = m.M
	v3319 = m.ExcPending
	if v3319 != 0 {
		goto L1
	} else {
		goto L892
	}
L892:
	;
	F_errfinish(m, int32(490005), int32(4924), int32(167834))
	mBase = m.M
	v3324 = m.ExcPending
	if v3324 != 0 {
		goto L1
	} else {
		goto L893
	}
L893:
	;
	goto L887
L894:
	;
	v3337 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	v3339 = *(*int64)(unsafe.Add(mBase, _consts[541]))
	if v3337 != v3339 {
		goto L17
	} else {
		goto L895
	}
L895:
	;
	v3342 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	if v3342 < int32(0) {
		goto L897
	} else {
		goto L898
	}
L896:
	;
	v3349 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v3350 = *(*int32)(unsafe.Add(mBase, uint32(v3349)+20))
	goto L900
L897:
	;
	v3346 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[291])) = v3346
	goto L899
L898:
	;
	goto L899
L899:
	;
	goto L896
L900:
	;
	if base.B2i32(v3350 == int32(2)) == int32(0) {
		goto L901
	} else {
		goto L902
	}
L901:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v3356 = m.ExcPending
	if v3356 != 0 {
		goto L1
	} else {
		goto L904
	}
L902:
	;
	goto L903
L903:
	;
	v3359 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v3360 = m.ExcPending
	if v3360 != 0 {
		goto L1
	} else {
		goto L906
	}
L904:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v3358 = m.ExcPending
	if v3358 != 0 {
		goto L1
	} else {
		goto L905
	}
L905:
	;
	goto L903
L906:
	;
	F_PushActiveSnapshot(m, v3359)
	mBase = m.M
	v3362 = m.ExcPending
	if v3362 != 0 {
		goto L1
	} else {
		goto L907
	}
L907:
	;
	v3365 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3365
	v3368 = *(*int32)(unsafe.Add(mBase, _consts[529]))
	v3369 = *(*int32)(unsafe.Add(mBase, uint32(v3368)))
	v3370 = *(*int32)(unsafe.Add(mBase, uint32(v22)+344))
	F_TwoPhaseTransactionGid(m, v3369, v3370, v22+int32(1344))
	mBase = m.M
	v3374 = m.ExcPending
	if v3374 != 0 {
		goto L1
	} else {
		goto L908
	}
L908:
	;
	v3376 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v3377 = *(*int32)(unsafe.Add(mBase, uint32(v3376)+24))
	goto L909
L909:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(v3377)) == int32(0) {
		goto L910
	} else {
		goto L911
	}
L910:
	;
	F_BeginTransactionBlock(m)
	mBase = m.M
	v3383 = m.ExcPending
	if v3383 != 0 {
		goto L1
	} else {
		goto L913
	}
L911:
	;
	goto L912
L912:
	;
	v3387 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	*(*int64)(unsafe.Add(mBase, _consts[110])) = v3387
	v3390 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	*(*int64)(unsafe.Add(mBase, _consts[111])) = v3390
	v3394 = F_PrepareTransactionBlock(m, v22+int32(1344))
	mBase = m.M
	v3395 = m.ExcPending
	if v3395 != 0 {
		goto L1
	} else {
		goto L915
	}
L913:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3385 = m.ExcPending
	if v3385 != 0 {
		goto L1
	} else {
		goto L914
	}
L914:
	;
	goto L912
L915:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v3397 = m.ExcPending
	if v3397 != 0 {
		goto L1
	} else {
		goto L916
	}
L916:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v3399 = m.ExcPending
	if v3399 != 0 {
		goto L1
	} else {
		goto L917
	}
L917:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3401 = m.ExcPending
	if v3401 != 0 {
		goto L1
	} else {
		goto L918
	}
L918:
	;
	v3403 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v3404 = m.ExcPending
	if v3404 != 0 {
		goto L1
	} else {
		goto L919
	}
L919:
	;
	v3405 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	v3407 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v3408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3407)+16)))
	if v3408 == int32(1) {
		goto L921
	} else {
		goto L922
	}
L920:
	;
	v3448 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[542])) = uint8(v3448)
	F_process_syncing_tables(m, v3446)
	mBase = m.M
	v3451 = m.ExcPending
	if v3451 != 0 {
		goto L1
	} else {
		goto L930
	}
L921:
	;
	v3411 = *(*int32)(unsafe.Add(mBase, uint32(v3407)))
	if v3411 == int32(3) {
		v3446 = v3405
		goto L920
	} else {
		goto L924
	}
L922:
	;
	goto L923
L923:
	;
	v3416 = *(*int32)(unsafe.Add(mBase, _consts[533]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3416
	v3419 = F_palloc(m, int32(24))
	mBase = m.M
	v3420 = m.ExcPending
	if v3420 != 0 {
		goto L1
	} else {
		goto L925
	}
L924:
	;
	goto L923
L925:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3419)+16)) = v3405
	*(*int64)(unsafe.Add(mBase, uint32(v3419)+8)) = int64(0)
	v3425 = *(*int32)(unsafe.Add(mBase, _consts[555]))
	if v3425 != 0 {
		goto L927
	} else {
		goto L928
	}
L926:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3419))) = v3432
	v3434 = int32(4094184)
	*(*int32)(unsafe.Add(mBase, uint32(v3419)+4)) = v3434
	*(*int32)(unsafe.Add(mBase, uint32(v3432)+4)) = v3419
	*(*int32)(unsafe.Add(mBase, _consts[556])) = v3419
	v3441 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3441
	v3443 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	v3446 = v3443
	goto L920
L927:
	;
	v3427 = *(*int32)(unsafe.Add(mBase, _consts[556]))
	v3432 = v3427
	goto L926
L928:
	;
	goto L929
L929:
	;
	v3429 = int32(4094184)
	*(*int32)(unsafe.Add(mBase, _consts[555])) = v3429
	v3432 = v3429
	goto L926
L930:
	;
	v3453 = *(*int64)(unsafe.Add(mBase, _consts[540]))
	if v3453 != int64(0) {
		goto L931
	} else {
		goto L932
	}
L931:
	;
	v3458 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3459 = m.ExcPending
	if v3459 != 0 {
		goto L1
	} else {
		goto L934
	}
L932:
	;
	goto L933
L933:
	;
	v3481 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	F_clear_subscription_skip_lsn(m, v3481)
	mBase = m.M
	v3483 = m.ExcPending
	if v3483 != 0 {
		goto L1
	} else {
		goto L940
	}
L934:
	;
	if v3458 != 0 {
		goto L935
	} else {
		goto L936
	}
L935:
	;
	v3461 = *(*int64)(unsafe.Add(mBase, _consts[540]))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+212)) = uint32(v3461)
	v3464 = int64(base.Ui64(v3461) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+208)) = uint32(v3464)
	F_errmsg(m, int32(509802), v22+int32(208))
	mBase = m.M
	v3470 = m.ExcPending
	if v3470 != 0 {
		goto L1
	} else {
		goto L938
	}
L936:
	;
	goto L937
L937:
	;
	*(*int64)(unsafe.Add(mBase, _consts[540])) = int64(0)
	goto L933
L938:
	;
	F_errfinish(m, int32(490005), int32(4938), int32(167863))
	mBase = m.M
	v3475 = m.ExcPending
	if v3475 != 0 {
		goto L1
	} else {
		goto L939
	}
L939:
	;
	goto L937
L940:
	;
	v3485 = int32(0)
	F_pgstat_report_activity(m, int32(2), v3485)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[535])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _consts[536])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[537])) = v3485
	*(*int32)(unsafe.Add(mBase, _consts[534])) = v3485
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v3485
	goto L3
L941:
	;
	v3693 = *(*int32)(unsafe.Add(mBase, uint32(v22)+344))
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v3693
	v3696 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	*(*int64)(unsafe.Add(mBase, _consts[536])) = v3696
	v3699 = *(*int32)(unsafe.Add(mBase, _consts[529]))
	v3700 = *(*int32)(unsafe.Add(mBase, uint32(v3699)))
	F_TwoPhaseTransactionGid(m, v3700, v3693, v22+int32(1344))
	mBase = m.M
	v3704 = m.ExcPending
	if v3704 != 0 {
		goto L1
	} else {
		goto L996
	}
L942:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3682 = m.ExcPending
	if v3682 != 0 {
		goto L1
	} else {
		goto L993
	}
L943:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3669 = m.ExcPending
	if v3669 != 0 {
		goto L1
	} else {
		goto L990
	}
L944:
	;
	v3511 = v3508 & int32(255)
	if v3511 == int32(0) {
		goto L945
	} else {
		goto L946
	}
L945:
	;
	v3514 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3515 = m.ExcPending
	if v3515 != 0 {
		goto L1
	} else {
		goto L948
	}
L946:
	;
	goto L947
L947:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3656 = m.ExcPending
	if v3656 != 0 {
		goto L1
	} else {
		goto L987
	}
L948:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3503))) = v3514
	if v3514 == int64(0) {
		goto L943
	} else {
		goto L949
	}
L949:
	;
	v3519 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3520 = m.ExcPending
	if v3520 != 0 {
		goto L1
	} else {
		goto L950
	}
L950:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3503)+8)) = v3519
	if v3519 == int64(0) {
		goto L942
	} else {
		goto L951
	}
L951:
	;
	v3524 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3525 = m.ExcPending
	if v3525 != 0 {
		goto L1
	} else {
		goto L952
	}
L952:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3503)+16)) = v3524
	v3528 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v3529 = m.ExcPending
	if v3529 != 0 {
		goto L1
	} else {
		goto L953
	}
L953:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3503)+24)) = v3528
	v3532 = v22 + int32(348)
	v3533 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v3534 = m.ExcPending
	if v3534 != 0 {
		goto L1
	} else {
		goto L954
	}
L954:
	;
	goto L958
L955:
	;
	m.G0 = v3506 + int32(16)
	goto L941
L956:
	;
	v3647 = F_strlen(m, v3636)
	mBase = m.M
	goto L955
L958:
	;
	goto L959
L959:
	;
	v3541 = int32(199)
	if (v3532^v3533)&int32(3) != 0 {
		goto L963
	} else {
		goto L964
	}
L960:
	;
	v3640 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3637))) = uint8(v3640)
	goto L956
L961:
	;
	v3621 = v3616
	v3622 = v3617
	v3623 = v3618
	goto L983
L962:
	;
	if v3611 == int32(0) {
		v3636 = v3609
		v3637 = v3610
		goto L960
	} else {
		goto L982
	}
L963:
	;
	v3609 = v3533
	v3610 = v3532
	v3611 = v3541
	goto L962
L964:
	;
	goto L965
L965:
	;
	if v3533&int32(3) == int32(0) {
		goto L967
	} else {
		goto L968
	}
L966:
	;
	if v3578 == int32(0) {
		v3636 = v3575
		v3637 = v3576
		goto L960
	} else {
		goto L975
	}
L967:
	;
	v3575 = v3533
	v3576 = v3532
	v3577 = v3541
	v3578 = int32(1)
	goto L966
L968:
	;
	goto L969
L969:
	;
	v3554 = v3533
	v3555 = v3532
	v3556 = v3541
	goto L970
L970:
	;
	v3558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3554))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3555))) = uint8(v3558)
	if v3558 == int32(0) {
		v3616 = v3554
		v3617 = v3555
		v3618 = v3556
		goto L961
	} else {
		goto L972
	}
L971:
	;
	v3575 = v3569
	v3576 = v3563
	v3577 = v3565
	v3578 = v3567
	goto L966
L972:
	;
	v3562 = int32(1)
	v3563 = v3555 + v3562
	v3565 = v3556 - v3562
	v3566 = int32(0)
	v3567 = base.B2i32(v3565 != v3566)
	v3569 = v3554 + v3562
	if v3569&int32(3) == v3566 {
		v3575 = v3569
		v3576 = v3563
		v3577 = v3565
		v3578 = v3567
		goto L966
	} else {
		goto L973
	}
L973:
	;
	if v3565 != 0 {
		v3554 = v3569
		v3555 = v3563
		v3556 = v3565
		goto L970
	} else {
		goto L974
	}
L974:
	;
	goto L971
L975:
	;
	v3581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3575))))
	if v3581 == int32(0) {
		v3609 = v3575
		v3610 = v3576
		v3611 = v3577
		goto L962
	} else {
		goto L976
	}
L976:
	;
	if base.Ui32(v3577) < base.Ui32(int32(4)) {
		v3609 = v3575
		v3610 = v3576
		v3611 = v3577
		goto L962
	} else {
		goto L977
	}
L977:
	;
	v3587 = v3575
	v3588 = v3576
	v3589 = v3577
	goto L978
L978:
	;
	v3592 = *(*int32)(unsafe.Add(mBase, uint32(v3587)))
	v3595 = int32(-2139062144)
	if (int32(16843008)-v3592|v3592)&v3595 != v3595 {
		v3616 = v3587
		v3617 = v3588
		v3618 = v3589
		goto L961
	} else {
		goto L980
	}
L979:
	;
	v3609 = v3603
	v3610 = v3601
	v3611 = v3605
	goto L962
L980:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3588))) = v3592
	v3600 = int32(4)
	v3601 = v3588 + v3600
	v3603 = v3587 + v3600
	v3605 = v3589 - v3600
	if base.Ui32(int32(3)) < base.Ui32(v3605) {
		v3587 = v3603
		v3588 = v3601
		v3589 = v3605
		goto L978
	} else {
		goto L981
	}
L981:
	;
	goto L979
L982:
	;
	v3616 = v3609
	v3617 = v3610
	v3618 = v3611
	goto L961
L983:
	;
	v3625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3621))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3622))) = uint8(v3625)
	if v3625 == int32(0) {
		v3636 = v3621
		v3637 = v3622
		goto L960
	} else {
		goto L985
	}
L984:
	;
	v3636 = v3632
	v3637 = v3630
	goto L960
L985:
	;
	v3629 = int32(1)
	v3630 = v3622 + v3629
	v3632 = v3621 + v3629
	v3634 = v3623 - v3629
	if v3634 != 0 {
		v3621 = v3632
		v3622 = v3630
		v3623 = v3634
		goto L983
	} else {
		goto L986
	}
L986:
	;
	goto L984
L987:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3506))) = v3511
	F_errmsg_internal(m, int32(400438), v3506)
	mBase = m.M
	v3660 = m.ExcPending
	if v3660 != 0 {
		goto L1
	} else {
		goto L988
	}
L988:
	;
	F_errfinish(m, int32(490650), int32(273), int32(446449))
	mBase = m.M
	v3665 = m.ExcPending
	if v3665 != 0 {
		goto L1
	} else {
		goto L989
	}
L989:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L990:
	;
	F_errmsg_internal(m, int32(400487), int32(0))
	mBase = m.M
	v3673 = m.ExcPending
	if v3673 != 0 {
		goto L1
	} else {
		goto L991
	}
L991:
	;
	F_errfinish(m, int32(490650), int32(278), int32(446449))
	mBase = m.M
	v3678 = m.ExcPending
	if v3678 != 0 {
		goto L1
	} else {
		goto L992
	}
L992:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L993:
	;
	F_errmsg_internal(m, int32(400536), int32(0))
	mBase = m.M
	v3686 = m.ExcPending
	if v3686 != 0 {
		goto L1
	} else {
		goto L994
	}
L994:
	;
	F_errfinish(m, int32(490650), int32(281), int32(446449))
	mBase = m.M
	v3691 = m.ExcPending
	if v3691 != 0 {
		goto L1
	} else {
		goto L995
	}
L995:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L996:
	;
	v3706 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	if v3706 < int32(0) {
		goto L998
	} else {
		goto L999
	}
L997:
	;
	v3713 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v3714 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+20))
	goto L1001
L998:
	;
	v3710 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[291])) = v3710
	goto L1000
L999:
	;
	goto L1000
L1000:
	;
	goto L997
L1001:
	;
	if base.B2i32(v3714 == int32(2)) == int32(0) {
		goto L1002
	} else {
		goto L1003
	}
L1002:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v3720 = m.ExcPending
	if v3720 != 0 {
		goto L1
	} else {
		goto L1005
	}
L1003:
	;
	goto L1004
L1004:
	;
	v3723 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v3724 = m.ExcPending
	if v3724 != 0 {
		goto L1
	} else {
		goto L1007
	}
L1005:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v3722 = m.ExcPending
	if v3722 != 0 {
		goto L1
	} else {
		goto L1006
	}
L1006:
	;
	goto L1004
L1007:
	;
	F_PushActiveSnapshot(m, v3723)
	mBase = m.M
	v3726 = m.ExcPending
	if v3726 != 0 {
		goto L1
	} else {
		goto L1008
	}
L1008:
	;
	v3729 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3729
	v3732 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	*(*int64)(unsafe.Add(mBase, _consts[110])) = v3732
	v3735 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	*(*int64)(unsafe.Add(mBase, _consts[111])) = v3735
	F_FinishPreparedTransaction(m, v22+int32(1344), int32(1))
	mBase = m.M
	v3741 = m.ExcPending
	if v3741 != 0 {
		goto L1
	} else {
		goto L1009
	}
L1009:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v3743 = m.ExcPending
	if v3743 != 0 {
		goto L1
	} else {
		goto L1010
	}
L1010:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v3745 = m.ExcPending
	if v3745 != 0 {
		goto L1
	} else {
		goto L1011
	}
L1011:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3747 = m.ExcPending
	if v3747 != 0 {
		goto L1
	} else {
		goto L1012
	}
L1012:
	;
	v3749 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v3750 = m.ExcPending
	if v3750 != 0 {
		goto L1
	} else {
		goto L1013
	}
L1013:
	;
	v3752 = *(*int64)(unsafe.Add(mBase, _consts[554]))
	v3753 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	v3755 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v3756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3755)+16)))
	if v3756 == int32(1) {
		goto L1015
	} else {
		goto L1016
	}
L1014:
	;
	v3795 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[542])) = uint8(v3795)
	F_process_syncing_tables(m, v3793)
	mBase = m.M
	v3798 = m.ExcPending
	if v3798 != 0 {
		goto L1
	} else {
		goto L1024
	}
L1015:
	;
	v3759 = *(*int32)(unsafe.Add(mBase, uint32(v3755)))
	if v3759 == int32(3) {
		v3793 = v3753
		goto L1014
	} else {
		goto L1018
	}
L1016:
	;
	goto L1017
L1017:
	;
	v3764 = *(*int32)(unsafe.Add(mBase, _consts[533]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3764
	v3767 = F_palloc(m, int32(24))
	mBase = m.M
	v3768 = m.ExcPending
	if v3768 != 0 {
		goto L1
	} else {
		goto L1019
	}
L1018:
	;
	goto L1017
L1019:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3767)+16)) = v3753
	*(*int64)(unsafe.Add(mBase, uint32(v3767)+8)) = v3752
	v3772 = *(*int32)(unsafe.Add(mBase, _consts[555]))
	if v3772 != 0 {
		goto L1021
	} else {
		goto L1022
	}
L1020:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3767))) = v3779
	v3781 = int32(4094184)
	*(*int32)(unsafe.Add(mBase, uint32(v3767)+4)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v3779)+4)) = v3767
	*(*int32)(unsafe.Add(mBase, _consts[556])) = v3767
	v3788 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3788
	v3790 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	v3793 = v3790
	goto L1014
L1021:
	;
	v3774 = *(*int32)(unsafe.Add(mBase, _consts[556]))
	v3779 = v3774
	goto L1020
L1022:
	;
	goto L1023
L1023:
	;
	v3776 = int32(4094184)
	*(*int32)(unsafe.Add(mBase, _consts[555])) = v3776
	v3779 = v3776
	goto L1020
L1024:
	;
	v3799 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	F_clear_subscription_skip_lsn(m, v3799)
	mBase = m.M
	v3801 = m.ExcPending
	if v3801 != 0 {
		goto L1
	} else {
		goto L1025
	}
L1025:
	;
	v3803 = int32(0)
	F_pgstat_report_activity(m, int32(2), v3803)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[535])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _consts[536])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[537])) = v3803
	*(*int32)(unsafe.Add(mBase, _consts[534])) = v3803
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v3803
	goto L3
L1026:
	;
	v4014 = *(*int32)(unsafe.Add(mBase, uint32(v22)+352))
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v4014
	v4017 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	*(*int64)(unsafe.Add(mBase, _consts[536])) = v4017
	v4020 = *(*int32)(unsafe.Add(mBase, _consts[529]))
	v4021 = *(*int32)(unsafe.Add(mBase, uint32(v4020)))
	F_TwoPhaseTransactionGid(m, v4021, v4014, v22+int32(1344))
	mBase = m.M
	v4025 = m.ExcPending
	if v4025 != 0 {
		goto L1
	} else {
		goto L1082
	}
L1027:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4003 = m.ExcPending
	if v4003 != 0 {
		goto L1
	} else {
		goto L1079
	}
L1028:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3990 = m.ExcPending
	if v3990 != 0 {
		goto L1
	} else {
		goto L1076
	}
L1029:
	;
	v3829 = v3826 & int32(255)
	if v3829 == int32(0) {
		goto L1030
	} else {
		goto L1031
	}
L1030:
	;
	v3832 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3833 = m.ExcPending
	if v3833 != 0 {
		goto L1
	} else {
		goto L1033
	}
L1031:
	;
	goto L1032
L1032:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3977 = m.ExcPending
	if v3977 != 0 {
		goto L1
	} else {
		goto L1073
	}
L1033:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3821))) = v3832
	if v3832 == int64(0) {
		goto L1028
	} else {
		goto L1034
	}
L1034:
	;
	v3837 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3838 = m.ExcPending
	if v3838 != 0 {
		goto L1
	} else {
		goto L1035
	}
L1035:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3821)+8)) = v3837
	if v3837 == int64(0) {
		goto L1027
	} else {
		goto L1036
	}
L1036:
	;
	v3842 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3843 = m.ExcPending
	if v3843 != 0 {
		goto L1
	} else {
		goto L1037
	}
L1037:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3821)+16)) = v3842
	v3845 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3846 = m.ExcPending
	if v3846 != 0 {
		goto L1
	} else {
		goto L1038
	}
L1038:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3821)+24)) = v3845
	v3849 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v3850 = m.ExcPending
	if v3850 != 0 {
		goto L1
	} else {
		goto L1039
	}
L1039:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3821)+32)) = v3849
	v3853 = v22 + int32(356)
	v3854 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v3855 = m.ExcPending
	if v3855 != 0 {
		goto L1
	} else {
		goto L1040
	}
L1040:
	;
	goto L1044
L1041:
	;
	m.G0 = v3824 + int32(16)
	goto L1026
L1042:
	;
	v3968 = F_strlen(m, v3957)
	mBase = m.M
	goto L1041
L1044:
	;
	goto L1045
L1045:
	;
	v3862 = int32(199)
	if (v3853^v3854)&int32(3) != 0 {
		goto L1049
	} else {
		goto L1050
	}
L1046:
	;
	v3961 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3958))) = uint8(v3961)
	goto L1042
L1047:
	;
	v3942 = v3937
	v3943 = v3938
	v3944 = v3939
	goto L1069
L1048:
	;
	if v3932 == int32(0) {
		v3957 = v3930
		v3958 = v3931
		goto L1046
	} else {
		goto L1068
	}
L1049:
	;
	v3930 = v3854
	v3931 = v3853
	v3932 = v3862
	goto L1048
L1050:
	;
	goto L1051
L1051:
	;
	if v3854&int32(3) == int32(0) {
		goto L1053
	} else {
		goto L1054
	}
L1052:
	;
	if v3899 == int32(0) {
		v3957 = v3896
		v3958 = v3897
		goto L1046
	} else {
		goto L1061
	}
L1053:
	;
	v3896 = v3854
	v3897 = v3853
	v3898 = v3862
	v3899 = int32(1)
	goto L1052
L1054:
	;
	goto L1055
L1055:
	;
	v3875 = v3854
	v3876 = v3853
	v3877 = v3862
	goto L1056
L1056:
	;
	v3879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3875))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3876))) = uint8(v3879)
	if v3879 == int32(0) {
		v3937 = v3875
		v3938 = v3876
		v3939 = v3877
		goto L1047
	} else {
		goto L1058
	}
L1057:
	;
	v3896 = v3890
	v3897 = v3884
	v3898 = v3886
	v3899 = v3888
	goto L1052
L1058:
	;
	v3883 = int32(1)
	v3884 = v3876 + v3883
	v3886 = v3877 - v3883
	v3887 = int32(0)
	v3888 = base.B2i32(v3886 != v3887)
	v3890 = v3875 + v3883
	if v3890&int32(3) == v3887 {
		v3896 = v3890
		v3897 = v3884
		v3898 = v3886
		v3899 = v3888
		goto L1052
	} else {
		goto L1059
	}
L1059:
	;
	if v3886 != 0 {
		v3875 = v3890
		v3876 = v3884
		v3877 = v3886
		goto L1056
	} else {
		goto L1060
	}
L1060:
	;
	goto L1057
L1061:
	;
	v3902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3896))))
	if v3902 == int32(0) {
		v3930 = v3896
		v3931 = v3897
		v3932 = v3898
		goto L1048
	} else {
		goto L1062
	}
L1062:
	;
	if base.Ui32(v3898) < base.Ui32(int32(4)) {
		v3930 = v3896
		v3931 = v3897
		v3932 = v3898
		goto L1048
	} else {
		goto L1063
	}
L1063:
	;
	v3908 = v3896
	v3909 = v3897
	v3910 = v3898
	goto L1064
L1064:
	;
	v3913 = *(*int32)(unsafe.Add(mBase, uint32(v3908)))
	v3916 = int32(-2139062144)
	if (int32(16843008)-v3913|v3913)&v3916 != v3916 {
		v3937 = v3908
		v3938 = v3909
		v3939 = v3910
		goto L1047
	} else {
		goto L1066
	}
L1065:
	;
	v3930 = v3924
	v3931 = v3922
	v3932 = v3926
	goto L1048
L1066:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3909))) = v3913
	v3921 = int32(4)
	v3922 = v3909 + v3921
	v3924 = v3908 + v3921
	v3926 = v3910 - v3921
	if base.Ui32(int32(3)) < base.Ui32(v3926) {
		v3908 = v3924
		v3909 = v3922
		v3910 = v3926
		goto L1064
	} else {
		goto L1067
	}
L1067:
	;
	goto L1065
L1068:
	;
	v3937 = v3930
	v3938 = v3931
	v3939 = v3932
	goto L1047
L1069:
	;
	v3946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3942))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3943))) = uint8(v3946)
	if v3946 == int32(0) {
		v3957 = v3942
		v3958 = v3943
		goto L1046
	} else {
		goto L1071
	}
L1070:
	;
	v3957 = v3953
	v3958 = v3951
	goto L1046
L1071:
	;
	v3950 = int32(1)
	v3951 = v3943 + v3950
	v3953 = v3942 + v3950
	v3955 = v3944 - v3950
	if v3955 != 0 {
		v3942 = v3953
		v3943 = v3951
		v3944 = v3955
		goto L1069
	} else {
		goto L1072
	}
L1072:
	;
	goto L1070
L1073:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3824))) = v3829
	F_errmsg_internal(m, int32(400582), v3824)
	mBase = m.M
	v3981 = m.ExcPending
	if v3981 != 0 {
		goto L1
	} else {
		goto L1074
	}
L1074:
	;
	F_errfinish(m, int32(490650), int32(332), int32(446481))
	mBase = m.M
	v3986 = m.ExcPending
	if v3986 != 0 {
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
	F_errmsg_internal(m, int32(400690), int32(0))
	mBase = m.M
	v3994 = m.ExcPending
	if v3994 != 0 {
		goto L1
	} else {
		goto L1077
	}
L1077:
	;
	F_errfinish(m, int32(490650), int32(337), int32(446481))
	mBase = m.M
	v3999 = m.ExcPending
	if v3999 != 0 {
		goto L1
	} else {
		goto L1078
	}
L1078:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1079:
	;
	F_errmsg_internal(m, int32(400633), int32(0))
	mBase = m.M
	v4007 = m.ExcPending
	if v4007 != 0 {
		goto L1
	} else {
		goto L1080
	}
L1080:
	;
	F_errfinish(m, int32(490650), int32(340), int32(446481))
	mBase = m.M
	v4012 = m.ExcPending
	if v4012 != 0 {
		goto L1
	} else {
		goto L1081
	}
L1081:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1082:
	;
	v4027 = v22 + int32(1344)
	v4028 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	v4029 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	v4030 = int32(0)
	v4031 = m.G0
	v4033 = v4031 - int32(16)
	m.G0 = v4033
	v4036 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v4040 = F_LWLockAcquire(m, v4036+int32(2304), int32(1))
	mBase = m.M
	v4041 = m.ExcPending
	if v4041 != 0 {
		goto L1
	} else {
		goto L1083
	}
L1083:
	;
	v4043 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v4044 = *(*int32)(unsafe.Add(mBase, uint32(v4043)+4))
	if v4044 <= int32(0) {
		v4133 = v4030
		goto L1084
	} else {
		goto L1085
	}
L1084:
	;
	v4152 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v4152+int32(2304))
	mBase = m.M
	v4156 = m.ExcPending
	if v4156 != 0 {
		goto L1
	} else {
		goto L1113
	}
L1085:
	;
	v4047 = v4043
	v4048 = v4030
	goto L1087
L1086:
	;
	F_pfree(m, v4114)
	mBase = m.M
	v4130 = m.ExcPending
	if v4130 != 0 {
		goto L1
	} else {
		goto L1112
	}
L1087:
	;
	v4069 = *(*int32)(unsafe.Add(mBase, uint32(v4047+v4048<<(uint(int32(2))%32))+8))
	v4070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4069)+44)))
	if v4070 != int32(1) {
		v4123 = v4047
		goto L1089
	} else {
		goto L1090
	}
L1088:
	;
	v4133 = int32(0)
	goto L1084
L1089:
	;
	v4125 = v4048 + int32(1)
	v4126 = *(*int32)(unsafe.Add(mBase, uint32(v4123)+4))
	if v4125 < v4126 {
		v4047 = v4123
		v4048 = v4125
		goto L1087
	} else {
		goto L1111
	}
L1090:
	;
	v4074 = v4069 + int32(47)
	v4077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4027))))
	v4078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4074))))
	if v4078 == int32(0) {
		v4097 = v4077
		v4098 = v4078
		goto L1092
	} else {
		goto L1093
	}
L1091:
	;
	if v4098-v4097 != 0 {
		v4123 = v4047
		goto L1089
	} else {
		goto L1099
	}
L1092:
	;
	goto L1091
L1093:
	;
	if v4077 != v4078 {
		v4097 = v4077
		v4098 = v4078
		goto L1092
	} else {
		goto L1094
	}
L1094:
	;
	v4082 = v4074
	v4083 = v4027
	goto L1095
L1095:
	;
	v4086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4083)+1)))
	v4087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4082)+1)))
	if v4087 == int32(0) {
		v4097 = v4086
		v4098 = v4087
		goto L1092
	} else {
		goto L1097
	}
L1096:
	;
	v4097 = v4086
	v4098 = v4087
	goto L1092
L1097:
	;
	v4090 = int32(1)
	if v4086 == v4087 {
		v4082 = v4082 + v4090
		v4083 = v4083 + v4090
		goto L1095
	} else {
		goto L1098
	}
L1098:
	;
	goto L1096
L1099:
	;
	v4100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4069)+45)))
	if v4100 == int32(1) {
		goto L1101
	} else {
		goto L1102
	}
L1100:
	;
	v4115 = *(*int64)(unsafe.Add(mBase, uint32(v4114)+56))
	if v4028 == v4115 {
		goto L1106
	} else {
		goto L1107
	}
L1101:
	;
	v4103 = *(*int32)(unsafe.Add(mBase, uint32(v4069)+32))
	v4105 = F_ReadTwoPhaseFile(m, v4103, int32(0))
	mBase = m.M
	v4106 = m.ExcPending
	if v4106 != 0 {
		goto L1
	} else {
		goto L1104
	}
L1102:
	;
	goto L1103
L1103:
	;
	v4107 = *(*int64)(unsafe.Add(mBase, uint32(v4069)+16))
	F_XlogReadTwoPhaseData(m, v4107, v4033+int32(12), int32(0))
	mBase = m.M
	v4112 = m.ExcPending
	if v4112 != 0 {
		goto L1
	} else {
		goto L1105
	}
L1104:
	;
	v4114 = v4105
	goto L1100
L1105:
	;
	v4113 = *(*int32)(unsafe.Add(mBase, uint32(v4033)+12))
	v4114 = v4113
	goto L1100
L1106:
	;
	v4117 = *(*int64)(unsafe.Add(mBase, uint32(v4114)+64))
	if v4117 == v4029 {
		goto L1086
	} else {
		goto L1109
	}
L1107:
	;
	goto L1108
L1108:
	;
	F_pfree(m, v4114)
	mBase = m.M
	v4120 = m.ExcPending
	if v4120 != 0 {
		goto L1
	} else {
		goto L1110
	}
L1109:
	;
	goto L1108
L1110:
	;
	v4122 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v4123 = v4122
	goto L1089
L1111:
	;
	goto L1088
L1112:
	;
	v4133 = int32(1)
	goto L1084
L1113:
	;
	m.G0 = v4033 + int32(16)
	if v4133 != 0 {
		goto L1114
	} else {
		goto L1115
	}
L1114:
	;
	v4161 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	*(*int64)(unsafe.Add(mBase, _consts[110])) = v4161
	v4164 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	*(*int64)(unsafe.Add(mBase, _consts[111])) = v4164
	v4167 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	if v4167 < int32(0) {
		goto L1118
	} else {
		goto L1119
	}
L1115:
	;
	goto L1116
L1116:
	;
	v4207 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v4208 = m.ExcPending
	if v4208 != 0 {
		goto L1
	} else {
		goto L1134
	}
L1117:
	;
	v4174 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v4175 = *(*int32)(unsafe.Add(mBase, uint32(v4174)+20))
	goto L1121
L1118:
	;
	v4171 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[291])) = v4171
	goto L1120
L1119:
	;
	goto L1120
L1120:
	;
	goto L1117
L1121:
	;
	if base.B2i32(v4175 == int32(2)) == int32(0) {
		goto L1122
	} else {
		goto L1123
	}
L1122:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v4181 = m.ExcPending
	if v4181 != 0 {
		goto L1
	} else {
		goto L1125
	}
L1123:
	;
	goto L1124
L1124:
	;
	v4184 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v4185 = m.ExcPending
	if v4185 != 0 {
		goto L1
	} else {
		goto L1127
	}
L1125:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v4183 = m.ExcPending
	if v4183 != 0 {
		goto L1
	} else {
		goto L1126
	}
L1126:
	;
	goto L1124
L1127:
	;
	F_PushActiveSnapshot(m, v4184)
	mBase = m.M
	v4187 = m.ExcPending
	if v4187 != 0 {
		goto L1
	} else {
		goto L1128
	}
L1128:
	;
	v4190 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v4190
	F_FinishPreparedTransaction(m, v22+int32(1344), int32(0))
	mBase = m.M
	v4196 = m.ExcPending
	if v4196 != 0 {
		goto L1
	} else {
		goto L1129
	}
L1129:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v4198 = m.ExcPending
	if v4198 != 0 {
		goto L1
	} else {
		goto L1130
	}
L1130:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v4200 = m.ExcPending
	if v4200 != 0 {
		goto L1
	} else {
		goto L1131
	}
L1131:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v4202 = m.ExcPending
	if v4202 != 0 {
		goto L1
	} else {
		goto L1132
	}
L1132:
	;
	v4203 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	F_clear_subscription_skip_lsn(m, v4203)
	mBase = m.M
	v4205 = m.ExcPending
	if v4205 != 0 {
		goto L1
	} else {
		goto L1133
	}
L1133:
	;
	goto L1116
L1134:
	;
	v4209 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	v4211 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v4212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4211)+16)))
	if v4212 == int32(1) {
		goto L1136
	} else {
		goto L1137
	}
L1135:
	;
	v4252 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[542])) = uint8(v4252)
	F_process_syncing_tables(m, v4250)
	mBase = m.M
	v4255 = m.ExcPending
	if v4255 != 0 {
		goto L1
	} else {
		goto L1145
	}
L1136:
	;
	v4215 = *(*int32)(unsafe.Add(mBase, uint32(v4211)))
	if v4215 == int32(3) {
		v4250 = v4209
		goto L1135
	} else {
		goto L1139
	}
L1137:
	;
	goto L1138
L1138:
	;
	v4220 = *(*int32)(unsafe.Add(mBase, _consts[533]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v4220
	v4223 = F_palloc(m, int32(24))
	mBase = m.M
	v4224 = m.ExcPending
	if v4224 != 0 {
		goto L1
	} else {
		goto L1140
	}
L1139:
	;
	goto L1138
L1140:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4223)+16)) = v4209
	*(*int64)(unsafe.Add(mBase, uint32(v4223)+8)) = int64(0)
	v4229 = *(*int32)(unsafe.Add(mBase, _consts[555]))
	if v4229 != 0 {
		goto L1142
	} else {
		goto L1143
	}
L1141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4223))) = v4236
	v4238 = int32(4094184)
	*(*int32)(unsafe.Add(mBase, uint32(v4223)+4)) = v4238
	*(*int32)(unsafe.Add(mBase, uint32(v4236)+4)) = v4223
	*(*int32)(unsafe.Add(mBase, _consts[556])) = v4223
	v4245 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v4245
	v4247 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	v4250 = v4247
	goto L1135
L1142:
	;
	v4231 = *(*int32)(unsafe.Add(mBase, _consts[556]))
	v4236 = v4231
	goto L1141
L1143:
	;
	goto L1144
L1144:
	;
	v4233 = int32(4094184)
	*(*int32)(unsafe.Add(mBase, _consts[555])) = v4233
	v4236 = v4233
	goto L1141
L1145:
	;
	v4257 = int32(0)
	F_pgstat_report_activity(m, int32(2), v4257)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[535])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _consts[536])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[537])) = v4257
	*(*int32)(unsafe.Add(mBase, _consts[534])) = v4257
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v4257
	goto L3
L1146:
	;
	v4283 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v4284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4283)+16)))
	if v4284 == int32(1) {
		goto L1147
	} else {
		goto L1148
	}
L1147:
	;
	v4287 = *(*int32)(unsafe.Add(mBase, uint32(v4283)))
	if v4287 == int32(1) {
		goto L15
	} else {
		goto L1150
	}
L1148:
	;
	goto L1149
L1149:
	;
	F_logicalrep_read_prepare_common(m, l0, int32(361413), v22+int32(320))
	mBase = m.M
	v4294 = m.ExcPending
	if v4294 != 0 {
		goto L1
	} else {
		goto L1151
	}
L1150:
	;
	goto L1149
L1151:
	;
	v4296 = *(*int32)(unsafe.Add(mBase, uint32(v22)+344))
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v4296
	v4299 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	*(*int64)(unsafe.Add(mBase, _consts[536])) = v4299
	v4302 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v4303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4302)+16)))
	if v4303 == int32(1) {
		goto L1154
	} else {
		goto L1155
	}
L1152:
	;
	v4455 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v4456 = m.ExcPending
	if v4456 != 0 {
		goto L1
	} else {
		goto L1209
	}
L1153:
	;
	v4380 = *(*int32)(unsafe.Add(mBase, _consts[532]))
	if v4380 != 0 {
		goto L1182
	} else {
		goto L1183
	}
L1154:
	;
	v4306 = *(*int32)(unsafe.Add(mBase, uint32(v4302)))
	if v4306 == int32(3) {
		goto L1153
	} else {
		goto L1157
	}
L1155:
	;
	goto L1156
L1156:
	;
	v4309 = F_pa_find_worker(m, v4296)
	mBase = m.M
	v4310 = m.ExcPending
	if v4310 != 0 {
		goto L1
	} else {
		goto L1160
	}
L1157:
	;
	goto L1156
L1158:
	;
	v4367 = *(*int32)(unsafe.Add(mBase, uint32(v22)+344))
	F_stream_open_and_write_change(m, v4367, int32(112), v22+int32(1344))
	mBase = m.M
	v4372 = m.ExcPending
	if v4372 != 0 {
		goto L1
	} else {
		goto L1179
	}
L1159:
	;
	F_pa_switch_to_partial_serialize(m, v4309, int32(1))
	mBase = m.M
	v4366 = m.ExcPending
	if v4366 != 0 {
		goto L1
	} else {
		goto L1178
	}
L1160:
	;
	if v4309 != 0 {
		goto L1161
	} else {
		goto L1162
	}
L1161:
	;
	v4311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4309)+12)))
	if v4311 != 0 {
		goto L1158
	} else {
		goto L1164
	}
L1162:
	;
	goto L1163
L1163:
	;
	v4322 = int32(*(*uint8)(unsafe.Add(mBase, _consts[539])))
	if v4322 != 0 {
		goto L14
	} else {
		goto L1168
	}
L1164:
	;
	v4312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4313 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4314 = F_pa_send_data(m, v4309, v4312, v4313)
	mBase = m.M
	v4315 = m.ExcPending
	if v4315 != 0 {
		goto L1
	} else {
		goto L1165
	}
L1165:
	;
	if v4314 == int32(0) {
		goto L1159
	} else {
		goto L1166
	}
L1166:
	;
	v4318 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	F_pa_xact_finish(m, v4309, v4318)
	mBase = m.M
	v4320 = m.ExcPending
	if v4320 != 0 {
		goto L1
	} else {
		goto L1167
	}
L1167:
	;
	goto L1152
L1168:
	;
	v4324 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v4325 = *(*int32)(unsafe.Add(mBase, uint32(v4324)+60))
	v4326 = *(*int32)(unsafe.Add(mBase, uint32(v22)+344))
	v4327 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	F_apply_spooled_messages(m, v4325, v4326, v4327)
	mBase = m.M
	v4329 = m.ExcPending
	if v4329 != 0 {
		goto L1
	} else {
		goto L1169
	}
L1169:
	;
	F_apply_handle_prepare_internal(m, v22+int32(320))
	mBase = m.M
	v4333 = m.ExcPending
	if v4333 != 0 {
		goto L1
	} else {
		goto L1170
	}
L1170:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v4335 = m.ExcPending
	if v4335 != 0 {
		goto L1
	} else {
		goto L1171
	}
L1171:
	;
	v4336 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	F_store_flush_position(m, v4336, int64(0))
	mBase = m.M
	v4339 = m.ExcPending
	if v4339 != 0 {
		goto L1
	} else {
		goto L1172
	}
L1172:
	;
	v4341 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[542])) = uint8(v4341)
	v4344 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v4345 = *(*int32)(unsafe.Add(mBase, uint32(v4344)+32))
	v4346 = *(*int32)(unsafe.Add(mBase, uint32(v22)+344))
	F_stream_cleanup_files(m, v4345, v4346)
	mBase = m.M
	v4348 = m.ExcPending
	if v4348 != 0 {
		goto L1
	} else {
		goto L1173
	}
L1173:
	;
	v4351 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v4352 = m.ExcPending
	if v4352 != 0 {
		goto L1
	} else {
		goto L1174
	}
L1174:
	;
	if v4351 == int32(0) {
		goto L1152
	} else {
		goto L1175
	}
L1175:
	;
	F_errmsg_internal(m, int32(424239), int32(0))
	mBase = m.M
	v4358 = m.ExcPending
	if v4358 != 0 {
		goto L1
	} else {
		goto L1176
	}
L1176:
	;
	F_errfinish(m, int32(490005), int32(1332), int32(361348))
	mBase = m.M
	v4363 = m.ExcPending
	if v4363 != 0 {
		goto L1
	} else {
		goto L1177
	}
L1177:
	;
	goto L1152
L1178:
	;
	goto L1158
L1179:
	;
	v4373 = *(*int32)(unsafe.Add(mBase, uint32(v4309)+16))
	F_pa_set_fileset_state(m, v4373)
	mBase = m.M
	v4375 = m.ExcPending
	if v4375 != 0 {
		goto L1
	} else {
		goto L1180
	}
L1180:
	;
	v4376 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	F_pa_xact_finish(m, v4309, v4376)
	mBase = m.M
	v4378 = m.ExcPending
	if v4378 != 0 {
		goto L1
	} else {
		goto L1181
	}
L1181:
	;
	goto L1152
L1182:
	;
	F_BufFileClose(m, v4380)
	mBase = m.M
	v4382 = m.ExcPending
	if v4382 != 0 {
		goto L1
	} else {
		goto L1185
	}
L1183:
	;
	goto L1184
L1184:
	;
	v4387 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	if v4387 < int32(0) {
		goto L1187
	} else {
		goto L1188
	}
L1185:
	;
	*(*int32)(unsafe.Add(mBase, _consts[532])) = int32(0)
	goto L1184
L1186:
	;
	v4394 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v4395 = *(*int32)(unsafe.Add(mBase, uint32(v4394)+20))
	goto L1190
L1187:
	;
	v4391 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[291])) = v4391
	goto L1189
L1188:
	;
	goto L1189
L1189:
	;
	goto L1186
L1190:
	;
	if base.B2i32(v4395 == int32(2)) == int32(0) {
		goto L1191
	} else {
		goto L1192
	}
L1191:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v4401 = m.ExcPending
	if v4401 != 0 {
		goto L1
	} else {
		goto L1194
	}
L1192:
	;
	goto L1193
L1193:
	;
	v4404 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v4405 = m.ExcPending
	if v4405 != 0 {
		goto L1
	} else {
		goto L1196
	}
L1194:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v4403 = m.ExcPending
	if v4403 != 0 {
		goto L1
	} else {
		goto L1195
	}
L1195:
	;
	goto L1193
L1196:
	;
	F_PushActiveSnapshot(m, v4404)
	mBase = m.M
	v4407 = m.ExcPending
	if v4407 != 0 {
		goto L1
	} else {
		goto L1197
	}
L1197:
	;
	v4410 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v4410
	F_apply_handle_prepare_internal(m, v22+int32(320))
	mBase = m.M
	v4415 = m.ExcPending
	if v4415 != 0 {
		goto L1
	} else {
		goto L1198
	}
L1198:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v4417 = m.ExcPending
	if v4417 != 0 {
		goto L1
	} else {
		goto L1199
	}
L1199:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v4419 = m.ExcPending
	if v4419 != 0 {
		goto L1
	} else {
		goto L1200
	}
L1200:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v4421 = m.ExcPending
	if v4421 != 0 {
		goto L1
	} else {
		goto L1201
	}
L1201:
	;
	v4423 = *(*int32)(unsafe.Add(mBase, _consts[547]))
	*(*int64)(unsafe.Add(mBase, uint32(v4423)+24)) = int64(0)
	F_pa_set_xact_state(m, v4423, int32(2))
	mBase = m.M
	v4428 = m.ExcPending
	if v4428 != 0 {
		goto L1
	} else {
		goto L1202
	}
L1202:
	;
	v4430 = *(*int32)(unsafe.Add(mBase, _consts[547]))
	v4431 = *(*int32)(unsafe.Add(mBase, uint32(v4430)+4))
	F_pa_unlock_transaction(m, v4431)
	mBase = m.M
	v4433 = m.ExcPending
	if v4433 != 0 {
		goto L1
	} else {
		goto L1203
	}
L1203:
	;
	*(*int32)(unsafe.Add(mBase, _consts[553])) = int32(0)
	goto L1204
L1204:
	;
	v4439 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v4440 = m.ExcPending
	if v4440 != 0 {
		goto L1
	} else {
		goto L1205
	}
L1205:
	;
	if v4439 == int32(0) {
		goto L1152
	} else {
		goto L1206
	}
L1206:
	;
	F_errmsg_internal(m, int32(424239), int32(0))
	mBase = m.M
	v4446 = m.ExcPending
	if v4446 != 0 {
		goto L1
	} else {
		goto L1207
	}
L1207:
	;
	F_errfinish(m, int32(490005), int32(1394), int32(361348))
	mBase = m.M
	v4451 = m.ExcPending
	if v4451 != 0 {
		goto L1
	} else {
		goto L1208
	}
L1208:
	;
	goto L1152
L1209:
	;
	v4457 = *(*int64)(unsafe.Add(mBase, uint32(v22)+328))
	F_process_syncing_tables(m, v4457)
	mBase = m.M
	v4459 = m.ExcPending
	if v4459 != 0 {
		goto L1
	} else {
		goto L1210
	}
L1210:
	;
	v4461 = *(*int64)(unsafe.Add(mBase, _consts[540]))
	if v4461 != int64(0) {
		goto L1211
	} else {
		goto L1212
	}
L1211:
	;
	v4466 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4467 = m.ExcPending
	if v4467 != 0 {
		goto L1
	} else {
		goto L1214
	}
L1212:
	;
	goto L1213
L1213:
	;
	v4489 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	F_clear_subscription_skip_lsn(m, v4489)
	mBase = m.M
	v4491 = m.ExcPending
	if v4491 != 0 {
		goto L1
	} else {
		goto L1220
	}
L1214:
	;
	if v4466 != 0 {
		goto L1215
	} else {
		goto L1216
	}
L1215:
	;
	v4469 = *(*int64)(unsafe.Add(mBase, _consts[540]))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+244)) = uint32(v4469)
	v4472 = int64(base.Ui64(v4469) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+240)) = uint32(v4472)
	F_errmsg(m, int32(509802), v22+int32(240))
	mBase = m.M
	v4478 = m.ExcPending
	if v4478 != 0 {
		goto L1
	} else {
		goto L1218
	}
L1216:
	;
	goto L1217
L1217:
	;
	*(*int64)(unsafe.Add(mBase, _consts[540])) = int64(0)
	goto L1213
L1218:
	;
	F_errfinish(m, int32(490005), int32(4938), int32(167863))
	mBase = m.M
	v4483 = m.ExcPending
	if v4483 != 0 {
		goto L1
	} else {
		goto L1219
	}
L1219:
	;
	goto L1217
L1220:
	;
	v4493 = int32(0)
	F_pgstat_report_activity(m, int32(2), v4493)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[535])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _consts[536])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[537])) = v4493
	*(*int32)(unsafe.Add(mBase, _consts[534])) = v4493
	*(*int32)(unsafe.Add(mBase, _consts[538])) = v4493
	goto L3
L1221:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4516 = m.ExcPending
	if v4516 != 0 {
		goto L1
	} else {
		goto L1222
	}
L1222:
	;
	v4517 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+36)) = uint32(v4517)
	v4520 = *(*int64)(unsafe.Add(mBase, _consts[541]))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+44)) = uint32(v4520)
	v4522 = int64(32)
	v4523 = int64(base.Ui64(v4517) >> (uint(v4522) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+32)) = uint32(v4523)
	v4526 = int64(base.Ui64(v4520) >> (uint(v4522) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+40)) = uint32(v4526)
	F_errmsg_internal(m, int32(660390), v22+int32(32))
	mBase = m.M
	v4532 = m.ExcPending
	if v4532 != 0 {
		goto L1
	} else {
		goto L1223
	}
L1223:
	;
	F_errfinish(m, int32(490005), int32(1021), int32(99748))
	mBase = m.M
	v4537 = m.ExcPending
	if v4537 != 0 {
		goto L1
	} else {
		goto L1224
	}
L1224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1225:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4544 = m.ExcPending
	if v4544 != 0 {
		goto L1
	} else {
		goto L1226
	}
L1226:
	;
	F_errmsg_internal(m, int32(400746), int32(0))
	mBase = m.M
	v4548 = m.ExcPending
	if v4548 != 0 {
		goto L1
	} else {
		goto L1227
	}
L1227:
	;
	F_errfinish(m, int32(490005), int32(1496), int32(81876))
	mBase = m.M
	v4553 = m.ExcPending
	if v4553 != 0 {
		goto L1
	} else {
		goto L1228
	}
L1228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1229:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4560 = m.ExcPending
	if v4560 != 0 {
		goto L1
	} else {
		goto L1230
	}
L1230:
	;
	F_errmsg_internal(m, int32(254351), int32(0))
	mBase = m.M
	v4564 = m.ExcPending
	if v4564 != 0 {
		goto L1
	} else {
		goto L1231
	}
L1231:
	;
	F_errfinish(m, int32(490005), int32(1510), int32(81876))
	mBase = m.M
	v4569 = m.ExcPending
	if v4569 != 0 {
		goto L1
	} else {
		goto L1232
	}
L1232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
	F_errmsg_internal(m, int32(477647), v22-int32(-64))
	mBase = m.M
	v4580 = m.ExcPending
	if v4580 != 0 {
		goto L1
	} else {
		goto L1234
	}
L1234:
	;
	F_errfinish(m, int32(490005), int32(1606), int32(81876))
	mBase = m.M
	v4585 = m.ExcPending
	if v4585 != 0 {
		goto L1
	} else {
		goto L1235
	}
L1235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1236:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4592 = m.ExcPending
	if v4592 != 0 {
		goto L1
	} else {
		goto L1237
	}
L1237:
	;
	F_errmsg_internal(m, int32(511850), int32(0))
	mBase = m.M
	v4596 = m.ExcPending
	if v4596 != 0 {
		goto L1
	} else {
		goto L1238
	}
L1238:
	;
	F_errfinish(m, int32(490005), int32(1651), int32(231870))
	mBase = m.M
	v4601 = m.ExcPending
	if v4601 != 0 {
		goto L1
	} else {
		goto L1239
	}
L1239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = int32(0)
	F_errmsg_internal(m, int32(477647), v22+int32(96))
	mBase = m.M
	v4612 = m.ExcPending
	if v4612 != 0 {
		goto L1
	} else {
		goto L1241
	}
L1241:
	;
	F_errfinish(m, int32(490005), int32(1722), int32(231870))
	mBase = m.M
	v4617 = m.ExcPending
	if v4617 != 0 {
		goto L1
	} else {
		goto L1242
	}
L1242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1243:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4624 = m.ExcPending
	if v4624 != 0 {
		goto L1
	} else {
		goto L1244
	}
L1244:
	;
	F_errmsg_internal(m, int32(520685), int32(0))
	mBase = m.M
	v4628 = m.ExcPending
	if v4628 != 0 {
		goto L1
	} else {
		goto L1245
	}
L1245:
	;
	F_errfinish(m, int32(490005), int32(1844), int32(80633))
	mBase = m.M
	v4633 = m.ExcPending
	if v4633 != 0 {
		goto L1
	} else {
		goto L1246
	}
L1246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1247:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4640 = m.ExcPending
	if v4640 != 0 {
		goto L1
	} else {
		goto L1248
	}
L1248:
	;
	F_errmsg_internal(m, int32(520726), int32(0))
	mBase = m.M
	v4644 = m.ExcPending
	if v4644 != 0 {
		goto L1
	} else {
		goto L1249
	}
L1249:
	;
	F_errfinish(m, int32(490005), int32(2161), int32(99691))
	mBase = m.M
	v4649 = m.ExcPending
	if v4649 != 0 {
		goto L1
	} else {
		goto L1250
	}
L1250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = int32(1)
	F_errmsg_internal(m, int32(477647), v22+int32(176))
	mBase = m.M
	v4660 = m.ExcPending
	if v4660 != 0 {
		goto L1
	} else {
		goto L1252
	}
L1252:
	;
	F_errfinish(m, int32(490005), int32(2242), int32(99691))
	mBase = m.M
	v4665 = m.ExcPending
	if v4665 != 0 {
		goto L1
	} else {
		goto L1253
	}
L1253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1254:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4672 = m.ExcPending
	if v4672 != 0 {
		goto L1
	} else {
		goto L1255
	}
L1255:
	;
	F_errmsg_internal(m, int32(400831), int32(0))
	mBase = m.M
	v4676 = m.ExcPending
	if v4676 != 0 {
		goto L1
	} else {
		goto L1256
	}
L1256:
	;
	F_errfinish(m, int32(490005), int32(1044), int32(361291))
	mBase = m.M
	v4681 = m.ExcPending
	if v4681 != 0 {
		goto L1
	} else {
		goto L1257
	}
L1257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1258:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4688 = m.ExcPending
	if v4688 != 0 {
		goto L1
	} else {
		goto L1259
	}
L1259:
	;
	v4689 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+228)) = uint32(v4689)
	v4692 = *(*int64)(unsafe.Add(mBase, _consts[541]))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+236)) = uint32(v4692)
	v4694 = int64(32)
	v4695 = int64(base.Ui64(v4689) >> (uint(v4694) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+224)) = uint32(v4695)
	v4698 = int64(base.Ui64(v4692) >> (uint(v4694) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+232)) = uint32(v4698)
	F_errmsg_internal(m, int32(660452), v22+int32(224))
	mBase = m.M
	v4704 = m.ExcPending
	if v4704 != 0 {
		goto L1
	} else {
		goto L1260
	}
L1260:
	;
	F_errfinish(m, int32(490005), int32(1113), int32(361376))
	mBase = m.M
	v4709 = m.ExcPending
	if v4709 != 0 {
		goto L1
	} else {
		goto L1261
	}
L1261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1262:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4716 = m.ExcPending
	if v4716 != 0 {
		goto L1
	} else {
		goto L1263
	}
L1263:
	;
	F_errmsg_internal(m, int32(520768), int32(0))
	mBase = m.M
	v4720 = m.ExcPending
	if v4720 != 0 {
		goto L1
	} else {
		goto L1264
	}
L1264:
	;
	F_errfinish(m, int32(490005), int32(1292), int32(361348))
	mBase = m.M
	v4725 = m.ExcPending
	if v4725 != 0 {
		goto L1
	} else {
		goto L1265
	}
L1265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1266:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4732 = m.ExcPending
	if v4732 != 0 {
		goto L1
	} else {
		goto L1267
	}
L1267:
	;
	F_errmsg_internal(m, int32(400881), int32(0))
	mBase = m.M
	v4736 = m.ExcPending
	if v4736 != 0 {
		goto L1
	} else {
		goto L1268
	}
L1268:
	;
	F_errfinish(m, int32(490005), int32(1298), int32(361348))
	mBase = m.M
	v4741 = m.ExcPending
	if v4741 != 0 {
		goto L1
	} else {
		goto L1269
	}
L1269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+256)) = int32(1)
	F_errmsg_internal(m, int32(477647), v22+int32(256))
	mBase = m.M
	v4752 = m.ExcPending
	if v4752 != 0 {
		goto L1
	} else {
		goto L1271
	}
L1271:
	;
	F_errfinish(m, int32(490005), int32(1398), int32(361348))
	mBase = m.M
	v4757 = m.ExcPending
	if v4757 != 0 {
		goto L1
	} else {
		goto L1272
	}
L1272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1273:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4764 = m.ExcPending
	if v4764 != 0 {
		goto L1
	} else {
		goto L1274
	}
L1274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v24
	F_errmsg(m, int32(708775), v22)
	mBase = m.M
	v4768 = m.ExcPending
	if v4768 != 0 {
		goto L1
	} else {
		goto L1275
	}
L1275:
	;
	F_errfinish(m, int32(490005), int32(3467), int32(321145))
	mBase = m.M
	v4773 = m.ExcPending
	if v4773 != 0 {
		goto L1
	} else {
		goto L1276
	}
L1276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1277:
	;
	goto L11
L1278:
	;
	if v2306 != v2311 {
		goto L10
	} else {
		goto L1279
	}
L1279:
	;
	v4785 = *(*int32)(unsafe.Add(mBase, uint32(v2320)+16))
	F_pa_set_fileset_state(m, v4785)
	mBase = m.M
	v4787 = m.ExcPending
	if v4787 != 0 {
		goto L1
	} else {
		goto L1280
	}
L1280:
	;
	F_pa_xact_finish(m, v2320, int64(0))
	mBase = m.M
	v4790 = m.ExcPending
	if v4790 != 0 {
		goto L1
	} else {
		goto L1281
	}
L1281:
	;
	goto L10
L1282:
	;
	if v4831 == int32(0) {
		goto L1283
	} else {
		goto L1284
	}
L1283:
	;
	if v4832 == int32(0) {
		goto L4
	} else {
		goto L1290
	}
L1284:
	;
	v4857 = int32(0)
	v4858 = *(*int32)(unsafe.Add(mBase, uint32(v4831)+4))
	if v4858 <= v4857 {
		goto L1283
	} else {
		goto L1285
	}
L1285:
	;
	v4861 = v4857
	goto L1286
L1286:
	;
	v4880 = *(*int32)(unsafe.Add(mBase, uint32(v4831)+12))
	v4884 = *(*int32)(unsafe.Add(mBase, uint32(v4880+v4861<<(uint(int32(2))%32))))
	F_logicalrep_rel_close(m, v4884, int32(0))
	mBase = m.M
	v4887 = m.ExcPending
	if v4887 != 0 {
		goto L1
	} else {
		goto L1288
	}
L1287:
	;
	goto L1283
L1288:
	;
	v4889 = v4861 + int32(1)
	v4890 = *(*int32)(unsafe.Add(mBase, uint32(v4831)+4))
	if v4889 < v4890 {
		v4861 = v4889
		goto L1286
	} else {
		goto L1289
	}
L1289:
	;
	goto L1287
L1290:
	;
	v4913 = int32(0)
	v4914 = *(*int32)(unsafe.Add(mBase, uint32(v4832)+4))
	if v4914 <= v4913 {
		goto L4
	} else {
		goto L1291
	}
L1291:
	;
	v4917 = v4913
	goto L1292
L1292:
	;
	v4936 = *(*int32)(unsafe.Add(mBase, uint32(v4832)+12))
	v4940 = *(*int32)(unsafe.Add(mBase, uint32(v4936+v4917<<(uint(int32(2))%32))))
	F_sequence_close(m, v4940, int32(0))
	mBase = m.M
	v4943 = m.ExcPending
	if v4943 != 0 {
		goto L1
	} else {
		goto L1294
	}
L1293:
	;
	goto L4
L1294:
	;
	v4945 = v4917 + int32(1)
	v4946 = *(*int32)(unsafe.Add(mBase, uint32(v4832)+4))
	if v4945 < v4946 {
		v4917 = v4945
		goto L1292
	} else {
		goto L1295
	}
L1295:
	;
	goto L1293
L1296:
	;
	v4953 = v22 + int32(1344)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v4953
	*(*int32)(unsafe.Add(mBase, uint32(v22)+300)) = v4953
	v4963 = F_list_make1_impl(m, int32(1), v22+int32(60))
	mBase = m.M
	v4964 = m.ExcPending
	if v4964 != 0 {
		goto L1
	} else {
		goto L1297
	}
L1297:
	;
	F_ReportApplyConflict(m, v753, v751, int32(15), int32(3), v636, v781, v4963)
	mBase = m.M
	v4966 = m.ExcPending
	if v4966 != 0 {
		goto L1
	} else {
		goto L1298
	}
L1298:
	;
	goto L7
L1299:
	;
	F_EvalPlanQualEnd(m, v22+int32(320))
	mBase = m.M
	v4974 = m.ExcPending
	if v4974 != 0 {
		goto L1
	} else {
		goto L1300
	}
L1300:
	;
	goto L6
L1301:
	;
	v4984 = *(*int32)(unsafe.Add(mBase, uint32(v630)+16))
	if v4984 != 0 {
		goto L1302
	} else {
		goto L1303
	}
L1302:
	;
	v4985 = *(*int32)(unsafe.Add(mBase, uint32(v630)+12))
	F_ExecCleanupTupleRouting(m, v4985, v4984)
	mBase = m.M
	v4987 = m.ExcPending
	if v4987 != 0 {
		goto L1
	} else {
		goto L1305
	}
L1303:
	;
	goto L1304
L1304:
	;
	v4988 = int32(0)
	v4989 = *(*int32)(unsafe.Add(mBase, uint32(v4981)+104))
	F_ExecResetTupleTable(m, v4989, v4988)
	mBase = m.M
	v4992 = m.ExcPending
	if v4992 != 0 {
		goto L1
	} else {
		goto L1306
	}
L1305:
	;
	goto L1304
L1306:
	;
	F_FreeExecutorState(m, v4981)
	mBase = m.M
	v4994 = m.ExcPending
	if v4994 != 0 {
		goto L1
	} else {
		goto L1307
	}
L1307:
	;
	F_pfree(m, v630)
	mBase = m.M
	v4996 = m.ExcPending
	if v4996 != 0 {
		goto L1
	} else {
		goto L1308
	}
L1308:
	;
	*(*int32)(unsafe.Add(mBase, _consts[537])) = int32(0)
	if v620 != 0 {
		v5005 = v4988
		goto L5
	} else {
		goto L1309
	}
L1309:
	;
	F_RestoreUserContext(m, v22+int32(304))
	mBase = m.M
	v5003 = m.ExcPending
	if v5003 != 0 {
		goto L1
	} else {
		goto L1310
	}
L1310:
	;
	v5005 = v4988
	goto L5
L1311:
	;
	goto L4
L1312:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v5047 = m.ExcPending
	if v5047 != 0 {
		goto L1
	} else {
		goto L1313
	}
L1313:
	;
	goto L3
}
