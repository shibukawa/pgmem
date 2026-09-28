package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v131 int64
	_ = v131
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v135 int32
	_ = v135
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v165 int64
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int64
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
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
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int64
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v490 int32
	_ = v490
	var v492 int64
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v503 int64
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
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
	var v748 int32
	_ = v748
	var v749 int64
	_ = v749
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v847 int64
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v858 int64
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
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
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1018 int64
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1029 int64
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
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
	var v1100 int32
	_ = v1100
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
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
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
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
	var v1198 int32
	_ = v1198
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
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
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
	var v1272 int32
	_ = v1272
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
	var v1304 int32
	_ = v1304
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
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
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1486 int32
	_ = v1486
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1649 int32
	_ = v1649
	var v1650 int64
	_ = v1650
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1693 int32
	_ = v1693
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
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
	var v1753 int32
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1757 int32
	_ = v1757
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1781 int32
	_ = v1781
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1790 int64
	_ = v1790
	var v1794 int32
	_ = v1794
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
	var v1820 int64
	_ = v1820
	var v1824 int32
	_ = v1824
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1861 int32
	_ = v1861
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1875 int32
	_ = v1875
	var v1892 int32
	_ = v1892
	var v1895 int64
	_ = v1895
	var v1908 int32
	_ = v1908
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1919 int32
	_ = v1919
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1930 int32
	_ = v1930
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1947 int32
	_ = v1947
	var v1975 int32
	_ = v1975
	var v1979 int32
	_ = v1979
	var v1984 int32
	_ = v1984
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2011 int32
	_ = v2011
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2049 int32
	_ = v2049
	var v2057 int32
	_ = v2057
	var v2059 int32
	_ = v2059
	var v2061 int32
	_ = v2061
	var v2066 int32
	_ = v2066
	var v2071 int32
	_ = v2071
	var v2073 int32
	_ = v2073
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2087 int32
	_ = v2087
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2129 int32
	_ = v2129
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
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2149 int32
	_ = v2149
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2156 int32
	_ = v2156
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2162 int32
	_ = v2162
	var v2167 int32
	_ = v2167
	var v2169 int32
	_ = v2169
	var v2171 int32
	_ = v2171
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2184 int32
	_ = v2184
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2217 int32
	_ = v2217
	var v2222 int32
	_ = v2222
	var v2224 int32
	_ = v2224
	var v2226 int32
	_ = v2226
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2235 int32
	_ = v2235
	var v2241 int32
	_ = v2241
	var v2246 int32
	_ = v2246
	var v2248 int32
	_ = v2248
	var v2252 int32
	_ = v2252
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2282 int64
	_ = v2282
	var v2284 int64
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2292 int32
	_ = v2292
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2301 int64
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2304 int64
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2307 int64
	_ = v2307
	var v2312 int32
	_ = v2312
	var v2315 int64
	_ = v2315
	var v2317 int32
	_ = v2317
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2348 int32
	_ = v2348
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
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
	var v2426 int64
	_ = v2426
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
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2460 int32
	_ = v2460
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2465 int64
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2468 int32
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2474 int64
	_ = v2474
	var v2476 int32
	_ = v2476
	var v2480 int32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2494 int32
	_ = v2494
	var v2505 int32
	_ = v2505
	var v2509 int32
	_ = v2509
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2520 int32
	_ = v2520
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2546 int64
	_ = v2546
	var v2548 int32
	_ = v2548
	var v2551 int32
	_ = v2551
	var v2558 int32
	_ = v2558
	var v2566 int64
	_ = v2566
	var v2570 int32
	_ = v2570
	var v2572 int64
	_ = v2572
	var v2574 int64
	_ = v2574
	var v2577 int64
	_ = v2577
	var v2578 int64
	_ = v2578
	var v2583 int64
	_ = v2583
	var v2590 int64
	_ = v2590
	var v2604 int32
	_ = v2604
	var v2606 int32
	_ = v2606
	var v2612 int32
	_ = v2612
	var v2617 int32
	_ = v2617
	var v2621 int32
	_ = v2621
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2628 int32
	_ = v2628
	var v2630 int32
	_ = v2630
	var v2634 int32
	_ = v2634
	var v2640 int32
	_ = v2640
	var v2645 int32
	_ = v2645
	var v2647 int32
	_ = v2647
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2654 int32
	_ = v2654
	var v2656 int32
	_ = v2656
	var v2658 int64
	_ = v2658
	var v2683 int32
	_ = v2683
	var v2685 int32
	_ = v2685
	var v2687 int32
	_ = v2687
	var v2709 int32
	_ = v2709
	var v2710 int32
	_ = v2710
	var v2716 int32
	_ = v2716
	var v2721 int32
	_ = v2721
	var v2724 int32
	_ = v2724
	var v2728 int32
	_ = v2728
	var v2733 int32
	_ = v2733
	var v2735 int32
	_ = v2735
	var v2738 int32
	_ = v2738
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2742 int64
	_ = v2742
	var v2745 int64
	_ = v2745
	var v2749 int32
	_ = v2749
	var v2752 int32
	_ = v2752
	var v2755 int32
	_ = v2755
	var v2758 int32
	_ = v2758
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2767 int32
	_ = v2767
	var v2769 int32
	_ = v2769
	var v2771 int32
	_ = v2771
	var v2772 int32
	_ = v2772
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2779 int32
	_ = v2779
	var v2781 int32
	_ = v2781
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2793 int32
	_ = v2793
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2807 int32
	_ = v2807
	var v2812 int32
	_ = v2812
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
	var v2862 int32
	_ = v2862
	var v2865 int32
	_ = v2865
	var v2891 int32
	_ = v2891
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2901 int32
	_ = v2901
	var v2906 int32
	_ = v2906
	var v2910 int32
	_ = v2910
	var v2917 int32
	_ = v2917
	var v2922 int32
	_ = v2922
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2928 int32
	_ = v2928
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2932 int32
	_ = v2932
	var v2933 int32
	_ = v2933
	var v2934 int32
	_ = v2934
	var v2937 int64
	_ = v2937
	var v2939 int64
	_ = v2939
	var v2942 int32
	_ = v2942
	var v2944 int32
	_ = v2944
	var v2945 int32
	_ = v2945
	var v2947 int32
	_ = v2947
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2955 int32
	_ = v2955
	var v2959 int32
	_ = v2959
	var v2963 int32
	_ = v2963
	var v2968 int32
	_ = v2968
	var v2969 int64
	_ = v2969
	var v2970 int32
	_ = v2970
	var v2972 int64
	_ = v2972
	var v2973 int32
	_ = v2973
	var v2975 int64
	_ = v2975
	var v2976 int32
	_ = v2976
	var v2984 int64
	_ = v2984
	var v2987 int32
	_ = v2987
	var v2988 int32
	_ = v2988
	var v2991 int32
	_ = v2991
	var v2994 int32
	_ = v2994
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
	var v3003 int64
	_ = v3003
	var v3005 int32
	_ = v3005
	var v3007 int32
	_ = v3007
	var v3009 int32
	_ = v3009
	var v3010 int32
	_ = v3010
	var v3011 int64
	_ = v3011
	var v3013 int32
	_ = v3013
	var v3017 int32
	_ = v3017
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3022 int32
	_ = v3022
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3032 int32
	_ = v3032
	var v3037 int32
	_ = v3037
	var v3040 int32
	_ = v3040
	var v3045 int32
	_ = v3045
	var v3046 int32
	_ = v3046
	var v3048 int32
	_ = v3048
	var v3049 int64
	_ = v3049
	var v3051 int32
	_ = v3051
	var v3053 int32
	_ = v3053
	var v3055 int32
	_ = v3055
	var v3062 int32
	_ = v3062
	var v3064 int32
	_ = v3064
	var v3066 int64
	_ = v3066
	var v3070 int32
	_ = v3070
	var v3072 int32
	_ = v3072
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3085 int32
	_ = v3085
	var v3090 int32
	_ = v3090
	var v3093 int64
	_ = v3093
	var v3095 int32
	_ = v3095
	var v3097 int32
	_ = v3097
	var v3115 int32
	_ = v3115
	var v3116 int32
	_ = v3116
	var v3119 int32
	_ = v3119
	var v3123 int32
	_ = v3123
	var v3124 int64
	_ = v3124
	var v3125 int32
	_ = v3125
	var v3129 int64
	_ = v3129
	var v3130 int32
	_ = v3130
	var v3134 int64
	_ = v3134
	var v3135 int32
	_ = v3135
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3142 int32
	_ = v3142
	var v3143 int32
	_ = v3143
	var v3144 int32
	_ = v3144
	var v3151 int32
	_ = v3151
	var v3155 int32
	_ = v3155
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3171 int32
	_ = v3171
	var v3175 int32
	_ = v3175
	var v3176 int32
	_ = v3176
	var v3178 int32
	_ = v3178
	var v3179 int32
	_ = v3179
	var v3180 int32
	_ = v3180
	var v3182 int32
	_ = v3182
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3190 int32
	_ = v3190
	var v3191 int32
	_ = v3191
	var v3194 int32
	_ = v3194
	var v3201 int32
	_ = v3201
	var v3202 int32
	_ = v3202
	var v3203 int32
	_ = v3203
	var v3206 int32
	_ = v3206
	var v3209 int32
	_ = v3209
	var v3214 int32
	_ = v3214
	var v3215 int32
	_ = v3215
	var v3217 int32
	_ = v3217
	var v3219 int32
	_ = v3219
	var v3223 int32
	_ = v3223
	var v3224 int32
	_ = v3224
	var v3225 int32
	_ = v3225
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3232 int32
	_ = v3232
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3237 int32
	_ = v3237
	var v3239 int32
	_ = v3239
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3246 int32
	_ = v3246
	var v3248 int32
	_ = v3248
	var v3250 int32
	_ = v3250
	var v3251 int32
	_ = v3251
	var v3254 int32
	_ = v3254
	var v3261 int32
	_ = v3261
	var v3267 int32
	_ = v3267
	var v3271 int32
	_ = v3271
	var v3276 int32
	_ = v3276
	var v3280 int32
	_ = v3280
	var v3284 int32
	_ = v3284
	var v3289 int32
	_ = v3289
	var v3291 int32
	_ = v3291
	var v3294 int64
	_ = v3294
	var v3299 int32
	_ = v3299
	var v3300 int64
	_ = v3300
	var v3309 int32
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3314 int64
	_ = v3314
	var v3317 int64
	_ = v3317
	var v3323 int32
	_ = v3323
	var v3328 int32
	_ = v3328
	var v3331 int32
	_ = v3331
	var v3340 int32
	_ = v3340
	var v3341 int64
	_ = v3341
	var v3343 int64
	_ = v3343
	var v3346 int32
	_ = v3346
	var v3350 int64
	_ = v3350
	var v3353 int32
	_ = v3353
	var v3354 int32
	_ = v3354
	var v3360 int32
	_ = v3360
	var v3362 int32
	_ = v3362
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3366 int32
	_ = v3366
	var v3369 int32
	_ = v3369
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3374 int32
	_ = v3374
	var v3378 int32
	_ = v3378
	var v3380 int32
	_ = v3380
	var v3381 int32
	_ = v3381
	var v3387 int32
	_ = v3387
	var v3389 int32
	_ = v3389
	var v3391 int64
	_ = v3391
	var v3394 int64
	_ = v3394
	var v3398 int32
	_ = v3398
	var v3399 int32
	_ = v3399
	var v3401 int32
	_ = v3401
	var v3403 int32
	_ = v3403
	var v3405 int32
	_ = v3405
	var v3407 int32
	_ = v3407
	var v3408 int32
	_ = v3408
	var v3409 int64
	_ = v3409
	var v3411 int32
	_ = v3411
	var v3412 int32
	_ = v3412
	var v3415 int32
	_ = v3415
	var v3420 int32
	_ = v3420
	var v3423 int32
	_ = v3423
	var v3424 int32
	_ = v3424
	var v3429 int32
	_ = v3429
	var v3431 int32
	_ = v3431
	var v3433 int32
	_ = v3433
	var v3436 int32
	_ = v3436
	var v3438 int32
	_ = v3438
	var v3445 int32
	_ = v3445
	var v3447 int64
	_ = v3447
	var v3450 int64
	_ = v3450
	var v3452 int32
	_ = v3452
	var v3455 int32
	_ = v3455
	var v3457 int64
	_ = v3457
	var v3462 int32
	_ = v3462
	var v3463 int32
	_ = v3463
	var v3465 int64
	_ = v3465
	var v3468 int64
	_ = v3468
	var v3474 int32
	_ = v3474
	var v3479 int32
	_ = v3479
	var v3485 int64
	_ = v3485
	var v3487 int32
	_ = v3487
	var v3489 int32
	_ = v3489
	var v3507 int32
	_ = v3507
	var v3508 int32
	_ = v3508
	var v3510 int32
	_ = v3510
	var v3512 int32
	_ = v3512
	var v3513 int32
	_ = v3513
	var v3515 int32
	_ = v3515
	var v3518 int64
	_ = v3518
	var v3519 int32
	_ = v3519
	var v3523 int64
	_ = v3523
	var v3524 int32
	_ = v3524
	var v3528 int64
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3532 int32
	_ = v3532
	var v3533 int32
	_ = v3533
	var v3536 int32
	_ = v3536
	var v3537 int32
	_ = v3537
	var v3538 int32
	_ = v3538
	var v3545 int32
	_ = v3545
	var v3549 int32
	_ = v3549
	var v3561 int32
	_ = v3561
	var v3562 int32
	_ = v3562
	var v3563 int32
	_ = v3563
	var v3565 int32
	_ = v3565
	var v3569 int32
	_ = v3569
	var v3570 int32
	_ = v3570
	var v3572 int32
	_ = v3572
	var v3573 int32
	_ = v3573
	var v3574 int32
	_ = v3574
	var v3576 int32
	_ = v3576
	var v3582 int32
	_ = v3582
	var v3583 int32
	_ = v3583
	var v3584 int32
	_ = v3584
	var v3585 int32
	_ = v3585
	var v3588 int32
	_ = v3588
	var v3595 int32
	_ = v3595
	var v3596 int32
	_ = v3596
	var v3597 int32
	_ = v3597
	var v3600 int32
	_ = v3600
	var v3603 int32
	_ = v3603
	var v3608 int32
	_ = v3608
	var v3609 int32
	_ = v3609
	var v3611 int32
	_ = v3611
	var v3613 int32
	_ = v3613
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3619 int32
	_ = v3619
	var v3624 int32
	_ = v3624
	var v3625 int32
	_ = v3625
	var v3626 int32
	_ = v3626
	var v3629 int32
	_ = v3629
	var v3630 int32
	_ = v3630
	var v3631 int32
	_ = v3631
	var v3633 int32
	_ = v3633
	var v3637 int32
	_ = v3637
	var v3638 int32
	_ = v3638
	var v3640 int32
	_ = v3640
	var v3642 int32
	_ = v3642
	var v3644 int32
	_ = v3644
	var v3645 int32
	_ = v3645
	var v3648 int32
	_ = v3648
	var v3655 int32
	_ = v3655
	var v3664 int32
	_ = v3664
	var v3668 int32
	_ = v3668
	var v3673 int32
	_ = v3673
	var v3677 int32
	_ = v3677
	var v3681 int32
	_ = v3681
	var v3686 int32
	_ = v3686
	var v3690 int32
	_ = v3690
	var v3694 int32
	_ = v3694
	var v3699 int32
	_ = v3699
	var v3701 int32
	_ = v3701
	var v3704 int64
	_ = v3704
	var v3707 int32
	_ = v3707
	var v3708 int32
	_ = v3708
	var v3712 int32
	_ = v3712
	var v3714 int32
	_ = v3714
	var v3718 int64
	_ = v3718
	var v3721 int32
	_ = v3721
	var v3722 int32
	_ = v3722
	var v3728 int32
	_ = v3728
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3732 int32
	_ = v3732
	var v3734 int32
	_ = v3734
	var v3737 int32
	_ = v3737
	var v3740 int64
	_ = v3740
	var v3743 int64
	_ = v3743
	var v3749 int32
	_ = v3749
	var v3751 int32
	_ = v3751
	var v3753 int32
	_ = v3753
	var v3755 int32
	_ = v3755
	var v3757 int32
	_ = v3757
	var v3758 int32
	_ = v3758
	var v3760 int64
	_ = v3760
	var v3761 int64
	_ = v3761
	var v3763 int32
	_ = v3763
	var v3764 int32
	_ = v3764
	var v3767 int32
	_ = v3767
	var v3772 int32
	_ = v3772
	var v3775 int32
	_ = v3775
	var v3776 int32
	_ = v3776
	var v3780 int32
	_ = v3780
	var v3782 int32
	_ = v3782
	var v3784 int32
	_ = v3784
	var v3787 int32
	_ = v3787
	var v3789 int32
	_ = v3789
	var v3796 int32
	_ = v3796
	var v3798 int64
	_ = v3798
	var v3801 int64
	_ = v3801
	var v3803 int32
	_ = v3803
	var v3806 int32
	_ = v3806
	var v3807 int64
	_ = v3807
	var v3809 int32
	_ = v3809
	var v3811 int32
	_ = v3811
	var v3829 int32
	_ = v3829
	var v3830 int32
	_ = v3830
	var v3832 int32
	_ = v3832
	var v3834 int32
	_ = v3834
	var v3835 int32
	_ = v3835
	var v3837 int32
	_ = v3837
	var v3840 int64
	_ = v3840
	var v3841 int32
	_ = v3841
	var v3845 int64
	_ = v3845
	var v3846 int32
	_ = v3846
	var v3850 int64
	_ = v3850
	var v3851 int32
	_ = v3851
	var v3853 int64
	_ = v3853
	var v3854 int32
	_ = v3854
	var v3857 int32
	_ = v3857
	var v3858 int32
	_ = v3858
	var v3861 int32
	_ = v3861
	var v3862 int32
	_ = v3862
	var v3863 int32
	_ = v3863
	var v3870 int32
	_ = v3870
	var v3874 int32
	_ = v3874
	var v3886 int32
	_ = v3886
	var v3887 int32
	_ = v3887
	var v3888 int32
	_ = v3888
	var v3890 int32
	_ = v3890
	var v3894 int32
	_ = v3894
	var v3895 int32
	_ = v3895
	var v3897 int32
	_ = v3897
	var v3898 int32
	_ = v3898
	var v3899 int32
	_ = v3899
	var v3901 int32
	_ = v3901
	var v3907 int32
	_ = v3907
	var v3908 int32
	_ = v3908
	var v3909 int32
	_ = v3909
	var v3910 int32
	_ = v3910
	var v3913 int32
	_ = v3913
	var v3920 int32
	_ = v3920
	var v3921 int32
	_ = v3921
	var v3922 int32
	_ = v3922
	var v3925 int32
	_ = v3925
	var v3928 int32
	_ = v3928
	var v3933 int32
	_ = v3933
	var v3934 int32
	_ = v3934
	var v3936 int32
	_ = v3936
	var v3938 int32
	_ = v3938
	var v3942 int32
	_ = v3942
	var v3943 int32
	_ = v3943
	var v3944 int32
	_ = v3944
	var v3949 int32
	_ = v3949
	var v3950 int32
	_ = v3950
	var v3951 int32
	_ = v3951
	var v3954 int32
	_ = v3954
	var v3955 int32
	_ = v3955
	var v3956 int32
	_ = v3956
	var v3958 int32
	_ = v3958
	var v3962 int32
	_ = v3962
	var v3963 int32
	_ = v3963
	var v3965 int32
	_ = v3965
	var v3967 int32
	_ = v3967
	var v3969 int32
	_ = v3969
	var v3970 int32
	_ = v3970
	var v3973 int32
	_ = v3973
	var v3980 int32
	_ = v3980
	var v3989 int32
	_ = v3989
	var v3993 int32
	_ = v3993
	var v3998 int32
	_ = v3998
	var v4002 int32
	_ = v4002
	var v4006 int32
	_ = v4006
	var v4011 int32
	_ = v4011
	var v4015 int32
	_ = v4015
	var v4019 int32
	_ = v4019
	var v4024 int32
	_ = v4024
	var v4026 int32
	_ = v4026
	var v4029 int64
	_ = v4029
	var v4032 int32
	_ = v4032
	var v4033 int32
	_ = v4033
	var v4035 int32
	_ = v4035
	var v4037 int32
	_ = v4037
	var v4038 int64
	_ = v4038
	var v4039 int64
	_ = v4039
	var v4040 int32
	_ = v4040
	var v4041 int32
	_ = v4041
	var v4043 int32
	_ = v4043
	var v4046 int32
	_ = v4046
	var v4050 int32
	_ = v4050
	var v4051 int32
	_ = v4051
	var v4053 int32
	_ = v4053
	var v4054 int32
	_ = v4054
	var v4057 int32
	_ = v4057
	var v4058 int32
	_ = v4058
	var v4079 int32
	_ = v4079
	var v4080 int32
	_ = v4080
	var v4084 int32
	_ = v4084
	var v4087 int32
	_ = v4087
	var v4090 int32
	_ = v4090
	var v4093 int32
	_ = v4093
	var v4094 int32
	_ = v4094
	var v4097 int32
	_ = v4097
	var v4098 int32
	_ = v4098
	var v4101 int32
	_ = v4101
	var v4108 int32
	_ = v4108
	var v4109 int32
	_ = v4109
	var v4111 int32
	_ = v4111
	var v4114 int64
	_ = v4114
	var v4116 int32
	_ = v4116
	var v4117 int32
	_ = v4117
	var v4118 int64
	_ = v4118
	var v4123 int32
	_ = v4123
	var v4124 int32
	_ = v4124
	var v4125 int32
	_ = v4125
	var v4126 int64
	_ = v4126
	var v4128 int64
	_ = v4128
	var v4131 int32
	_ = v4131
	var v4133 int32
	_ = v4133
	var v4134 int32
	_ = v4134
	var v4136 int32
	_ = v4136
	var v4137 int32
	_ = v4137
	var v4141 int32
	_ = v4141
	var v4144 int32
	_ = v4144
	var v4163 int32
	_ = v4163
	var v4167 int32
	_ = v4167
	var v4172 int64
	_ = v4172
	var v4175 int64
	_ = v4175
	var v4178 int32
	_ = v4178
	var v4182 int64
	_ = v4182
	var v4185 int32
	_ = v4185
	var v4186 int32
	_ = v4186
	var v4192 int32
	_ = v4192
	var v4194 int32
	_ = v4194
	var v4195 int32
	_ = v4195
	var v4196 int32
	_ = v4196
	var v4198 int32
	_ = v4198
	var v4201 int32
	_ = v4201
	var v4207 int32
	_ = v4207
	var v4209 int32
	_ = v4209
	var v4211 int32
	_ = v4211
	var v4213 int32
	_ = v4213
	var v4214 int64
	_ = v4214
	var v4216 int32
	_ = v4216
	var v4218 int32
	_ = v4218
	var v4219 int32
	_ = v4219
	var v4220 int64
	_ = v4220
	var v4222 int32
	_ = v4222
	var v4223 int32
	_ = v4223
	var v4226 int32
	_ = v4226
	var v4231 int32
	_ = v4231
	var v4234 int32
	_ = v4234
	var v4235 int32
	_ = v4235
	var v4240 int32
	_ = v4240
	var v4242 int32
	_ = v4242
	var v4244 int32
	_ = v4244
	var v4247 int32
	_ = v4247
	var v4249 int32
	_ = v4249
	var v4256 int32
	_ = v4256
	var v4258 int64
	_ = v4258
	var v4261 int64
	_ = v4261
	var v4263 int32
	_ = v4263
	var v4266 int32
	_ = v4266
	var v4268 int32
	_ = v4268
	var v4285 int64
	_ = v4285
	var v4287 int64
	_ = v4287
	var v4290 int32
	_ = v4290
	var v4292 int32
	_ = v4292
	var v4293 int32
	_ = v4293
	var v4296 int32
	_ = v4296
	var v4303 int32
	_ = v4303
	var v4305 int32
	_ = v4305
	var v4308 int64
	_ = v4308
	var v4311 int32
	_ = v4311
	var v4312 int32
	_ = v4312
	var v4315 int32
	_ = v4315
	var v4318 int32
	_ = v4318
	var v4319 int32
	_ = v4319
	var v4320 int32
	_ = v4320
	var v4321 int32
	_ = v4321
	var v4322 int32
	_ = v4322
	var v4323 int32
	_ = v4323
	var v4324 int32
	_ = v4324
	var v4327 int64
	_ = v4327
	var v4329 int32
	_ = v4329
	var v4331 int32
	_ = v4331
	var v4333 int32
	_ = v4333
	var v4334 int32
	_ = v4334
	var v4335 int32
	_ = v4335
	var v4336 int64
	_ = v4336
	var v4338 int32
	_ = v4338
	var v4342 int32
	_ = v4342
	var v4344 int32
	_ = v4344
	var v4345 int64
	_ = v4345
	var v4348 int32
	_ = v4348
	var v4350 int32
	_ = v4350
	var v4353 int32
	_ = v4353
	var v4354 int32
	_ = v4354
	var v4355 int32
	_ = v4355
	var v4357 int32
	_ = v4357
	var v4360 int32
	_ = v4360
	var v4361 int32
	_ = v4361
	var v4367 int32
	_ = v4367
	var v4372 int32
	_ = v4372
	var v4375 int32
	_ = v4375
	var v4376 int32
	_ = v4376
	var v4381 int32
	_ = v4381
	var v4382 int32
	_ = v4382
	var v4384 int32
	_ = v4384
	var v4385 int64
	_ = v4385
	var v4387 int32
	_ = v4387
	var v4389 int32
	_ = v4389
	var v4391 int32
	_ = v4391
	var v4396 int32
	_ = v4396
	var v4400 int64
	_ = v4400
	var v4403 int32
	_ = v4403
	var v4404 int32
	_ = v4404
	var v4410 int32
	_ = v4410
	var v4412 int32
	_ = v4412
	var v4413 int32
	_ = v4413
	var v4414 int32
	_ = v4414
	var v4416 int32
	_ = v4416
	var v4419 int32
	_ = v4419
	var v4424 int32
	_ = v4424
	var v4426 int32
	_ = v4426
	var v4428 int32
	_ = v4428
	var v4430 int32
	_ = v4430
	var v4432 int32
	_ = v4432
	var v4437 int32
	_ = v4437
	var v4439 int32
	_ = v4439
	var v4440 int32
	_ = v4440
	var v4442 int32
	_ = v4442
	var v4448 int32
	_ = v4448
	var v4449 int32
	_ = v4449
	var v4455 int32
	_ = v4455
	var v4460 int32
	_ = v4460
	var v4464 int32
	_ = v4464
	var v4465 int32
	_ = v4465
	var v4466 int64
	_ = v4466
	var v4468 int32
	_ = v4468
	var v4470 int64
	_ = v4470
	var v4475 int32
	_ = v4475
	var v4476 int32
	_ = v4476
	var v4478 int64
	_ = v4478
	var v4481 int64
	_ = v4481
	var v4487 int32
	_ = v4487
	var v4492 int32
	_ = v4492
	var v4498 int64
	_ = v4498
	var v4500 int32
	_ = v4500
	var v4502 int32
	_ = v4502
	var v4522 int32
	_ = v4522
	var v4525 int32
	_ = v4525
	var v4526 int64
	_ = v4526
	var v4529 int64
	_ = v4529
	var v4531 int64
	_ = v4531
	var v4532 int64
	_ = v4532
	var v4535 int64
	_ = v4535
	var v4541 int32
	_ = v4541
	var v4546 int32
	_ = v4546
	var v4550 int32
	_ = v4550
	var v4553 int32
	_ = v4553
	var v4557 int32
	_ = v4557
	var v4564 int32
	_ = v4564
	var v4569 int32
	_ = v4569
	var v4573 int32
	_ = v4573
	var v4576 int32
	_ = v4576
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
	var v4644 int32
	_ = v4644
	var v4649 int32
	_ = v4649
	var v4653 int32
	_ = v4653
	var v4656 int32
	_ = v4656
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
	var v4692 int32
	_ = v4692
	var v4697 int32
	_ = v4697
	var v4701 int32
	_ = v4701
	var v4704 int32
	_ = v4704
	var v4708 int32
	_ = v4708
	var v4713 int32
	_ = v4713
	var v4717 int32
	_ = v4717
	var v4720 int32
	_ = v4720
	var v4721 int64
	_ = v4721
	var v4724 int64
	_ = v4724
	var v4726 int64
	_ = v4726
	var v4727 int64
	_ = v4727
	var v4730 int64
	_ = v4730
	var v4736 int32
	_ = v4736
	var v4741 int32
	_ = v4741
	var v4745 int32
	_ = v4745
	var v4748 int32
	_ = v4748
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
	var v4784 int32
	_ = v4784
	var v4789 int32
	_ = v4789
	var v4793 int32
	_ = v4793
	var v4796 int32
	_ = v4796
	var v4800 int32
	_ = v4800
	var v4805 int32
	_ = v4805
	var v4808 int32
	_ = v4808
	var v4813 int32
	_ = v4813
	var v4815 int32
	_ = v4815
	var v4817 int32
	_ = v4817
	var v4820 int32
	_ = v4820
	var v4847 int32
	_ = v4847
	var v4858 int32
	_ = v4858
	var v4860 int32
	_ = v4860
	var v4861 int32
	_ = v4861
	var v4863 int32
	_ = v4863
	var v4864 int32
	_ = v4864
	var v4875 int32
	_ = v4875
	var v4877 int32
	_ = v4877
	var v4878 int32
	_ = v4878
	var v4884 int32
	_ = v4884
	var v4887 int32
	_ = v4887
	var v4888 int32
	_ = v4888
	var v4891 int32
	_ = v4891
	var v4910 int32
	_ = v4910
	var v4914 int32
	_ = v4914
	var v4917 int32
	_ = v4917
	var v4919 int32
	_ = v4919
	var v4920 int32
	_ = v4920
	var v4943 int32
	_ = v4943
	var v4947 int32
	_ = v4947
	var v4966 int32
	_ = v4966
	var v4970 int32
	_ = v4970
	var v4973 int32
	_ = v4973
	var v4975 int32
	_ = v4975
	var v4976 int32
	_ = v4976
	var v4984 int32
	_ = v4984
	var v4985 int32
	_ = v4985
	var v4987 int32
	_ = v4987
	var v4989 int32
	_ = v4989
	var v4992 int32
	_ = v4992
	var v4996 int32
	_ = v4996
	var v4998 int32
	_ = v4998
	var v5005 int32
	_ = v5005
	var v5006 int32
	_ = v5006
	var v5008 int32
	_ = v5008
	var v5012 int32
	_ = v5012
	var v5016 int32
	_ = v5016
	var v5023 int32
	_ = v5023
	var v5025 int32
	_ = v5025
	var v5026 int32
	_ = v5026
	var v5027 int32
	_ = v5027
	var v5029 int32
	_ = v5029
	var v5031 int32
	_ = v5031
	var v5032 int32
	_ = v5032
	var v5033 int32
	_ = v5033
	var v5036 int32
	_ = v5036
	var v5038 int32
	_ = v5038
	var v5040 int32
	_ = v5040
	var v5047 int32
	_ = v5047
	var v5049 int32
	_ = v5049
	var v5068 int32
	_ = v5068
	var v5089 int32
	_ = v5089
	var v5091 int32
	_ = v5091
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(1568)
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
	v26 = int32(_a_F_apply_dispatch_0)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v24
	switch v24 - int32(65) {
	case 0:
		goto L35
	case 1:
		goto L46
	case 2:
		goto L45
	case 3:
		goto L42
	case 4:
		goto L36
	default:
		goto L13
	case 8:
		goto L44
	case 10:
		goto L31
	case 12:
		goto L3
	case 14:
		goto L38
	case 15:
		goto L32
	case 17:
		goto L40
	case 18:
		goto L37
	case 19:
		goto L41
	case 20:
		goto L43
	case 24:
		goto L39
	case 33:
		goto L33
	case 34:
		goto L34
	case 47:
		goto L29
	case 49:
		goto L30
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v27
	m.G0 = v22 + int32(1568)
	return
L4:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v5089 = m.ExcPending
	if v5089 != 0 {
		goto L1
	} else {
		goto L1336
	}
L5:
	;
	F_logicalrep_rel_close(m, v604, v5049)
	mBase = m.M
	v5068 = m.ExcPending
	if v5068 != 0 {
		goto L1
	} else {
		goto L1335
	}
L6:
	;
	v5023 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	F_AfterTriggerEndQuery(m, v5023)
	mBase = m.M
	v5025 = m.ExcPending
	if v5025 != 0 {
		goto L1
	} else {
		goto L1324
	}
L7:
	;
	F_ExecCloseIndices(m, v746)
	mBase = m.M
	v5012 = m.ExcPending
	if v5012 != 0 {
		goto L1
	} else {
		goto L1322
	}
L8:
	;
	v4984 = F_FindDeletedTupleInLocalRel(m, v747, v744, v632, v22+int32(1368), v22+int32(1372), v22+int32(1376))
	mBase = m.M
	v4985 = m.ExcPending
	if v4985 != 0 {
		goto L1
	} else {
		goto L1314
	}
L9:
	;
	v4875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1360)))
	v4877 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v4878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4877)+39)))
	F_ExecuteTruncateGuts(m, v4860, v4858, v4864, int32(0), v4875, (v4878^int32(-1))&int32(1))
	mBase = m.M
	v4884 = m.ExcPending
	if v4884 != 0 {
		goto L1
	} else {
		goto L1299
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	v4847 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v4847
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v4847
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v4847
	goto L3
L11:
	;
	F_stream_open_and_write_change(m, v2317, int32(65), v22+int32(320))
	mBase = m.M
	v4813 = m.ExcPending
	if v4813 != 0 {
		goto L1
	} else {
		goto L1295
	}
L12:
	;
	F_pa_switch_to_partial_serialize(m, v2326, int32(1))
	mBase = m.M
	v4808 = m.ExcPending
	if v4808 != 0 {
		goto L1
	} else {
		goto L1294
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4793 = m.ExcPending
	if v4793 != 0 {
		goto L1
	} else {
		goto L1290
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4777 = m.ExcPending
	if v4777 != 0 {
		goto L1
	} else {
		goto L1287
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4761 = m.ExcPending
	if v4761 != 0 {
		goto L1
	} else {
		goto L1283
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4745 = m.ExcPending
	if v4745 != 0 {
		goto L1
	} else {
		goto L1279
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4717 = m.ExcPending
	if v4717 != 0 {
		goto L1
	} else {
		goto L1275
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4701 = m.ExcPending
	if v4701 != 0 {
		goto L1
	} else {
		goto L1271
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4685 = m.ExcPending
	if v4685 != 0 {
		goto L1
	} else {
		goto L1268
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4669 = m.ExcPending
	if v4669 != 0 {
		goto L1
	} else {
		goto L1264
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4653 = m.ExcPending
	if v4653 != 0 {
		goto L1
	} else {
		goto L1260
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4637 = m.ExcPending
	if v4637 != 0 {
		goto L1
	} else {
		goto L1257
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4621 = m.ExcPending
	if v4621 != 0 {
		goto L1
	} else {
		goto L1253
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4605 = m.ExcPending
	if v4605 != 0 {
		goto L1
	} else {
		goto L1250
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4589 = m.ExcPending
	if v4589 != 0 {
		goto L1
	} else {
		goto L1246
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4573 = m.ExcPending
	if v4573 != 0 {
		goto L1
	} else {
		goto L1242
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4550 = m.ExcPending
	if v4550 != 0 {
		goto L1
	} else {
		goto L1238
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4522 = m.ExcPending
	if v4522 != 0 {
		goto L1
	} else {
		goto L1234
	}
L29:
	;
	v4285 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+1368)) = v4285
	v4287 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+1360)) = v4287
	v4290 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v4290 != 0 {
		goto L16
	} else {
		goto L1159
	}
L30:
	;
	v3829 = v22 + int32(336)
	v3830 = m.G0
	v3832 = v3830 - int32(16)
	m.G0 = v3832
	v3834 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v3835 = m.ExcPending
	if v3835 != 0 {
		goto L1
	} else {
		goto L1044
	}
L31:
	;
	v3507 = v22 + int32(336)
	v3508 = m.G0
	v3510 = v3508 - int32(16)
	m.G0 = v3510
	v3512 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v3513 = m.ExcPending
	if v3513 != 0 {
		goto L1
	} else {
		goto L960
	}
L32:
	;
	F_logicalrep_read_prepare_common(m, l0, int32(_a_F_apply_dispatch_1), v22+int32(336))
	mBase = m.M
	v3340 = m.ExcPending
	if v3340 != 0 {
		goto L1
	} else {
		goto L910
	}
L33:
	;
	v3115 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v3116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3115)+16)))
	if v3116 == int32(1) {
		goto L852
	} else {
		goto L853
	}
L34:
	;
	v2937 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+1368)) = v2937
	v2939 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+1360)) = v2939
	v2942 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v2942 != 0 {
		goto L20
	} else {
		goto L799
	}
L35:
	;
	v2282 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+328)) = v2282
	v2284 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+320)) = v2284
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v2287 != 0 {
		goto L21
	} else {
		goto L633
	}
L36:
	;
	v2119 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v2119 == int32(0) {
		goto L23
	} else {
		goto L590
	}
L37:
	;
	v1602 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v1602 != 0 {
		goto L26
	} else {
		goto L486
	}
L38:
	;
	v1564 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v1564 != 0 {
		goto L3
	} else {
		goto L475
	}
L39:
	;
	v1542 = F_handle_streamed_transaction(m, int32(89), l0)
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L1
	} else {
		goto L465
	}
L40:
	;
	v1360 = F_handle_streamed_transaction(m, int32(82), l0)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L1
	} else {
		goto L421
	}
L41:
	;
	v1013 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+336)) = uint8(v1013)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+1360)) = uint8(v1013)
	v1018 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	if v1018 != int64(0) {
		goto L3
	} else {
		goto L320
	}
L42:
	;
	v847 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	if v847 != int64(0) {
		goto L3
	} else {
		goto L260
	}
L43:
	;
	v492 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	if v492 != int64(0) {
		goto L3
	} else {
		goto L160
	}
L44:
	;
	v165 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	if v165 != int64(0) {
		goto L3
	} else {
		goto L75
	}
L45:
	;
	v105 = m.G0
	v107 = v105 - int32(16)
	m.G0 = v107
	v109 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L62
	}
L46:
	;
	v33 = v22 + int32(336)
	v34 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v34
	if v34 == int64(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v52 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L54
	}
L51:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_2), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(68), int32(_a_F_apply_dispatch_4))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v52
	v56 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v22)+352))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v60
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v63
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[9])) = v63
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)+16))
	if base.B2i32(v69 == int64(0))|base.B2i32(v63 != v69) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v100 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])) = uint8(v100)
	F_pgstat_report_activity(m, int32(3), int32(0))
	mBase = m.M
	goto L3
L57:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8])) = v63
	v78 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v78 == int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v83 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+20)) = uint32(v83)
	v86 = int64(base.Ui64(v83) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+16)) = uint32(v86)
	F_errmsg(m, int32(_a_F_apply_dispatch_5), v22+int32(16))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(_a_F_apply_dispatch_7), int32(_a_F_apply_dispatch_8))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L56
L62:
	;
	v112 = v109 & int32(255)
	if v112 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v127 = v22 + int32(336)
	v128 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L69
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v112
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_9), v107)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(104), int32(_a_F_apply_dispatch_10))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = v128
	v131 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v127)+8)) = v131
	v134 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v127)+16)) = v134
	m.G0 = v107 + int32(16)
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	v142 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[9]))
	if v140 != v142 {
		goto L28
	} else {
		goto L72
	}
L72:
	;
	F_apply_handle_commit_internal(m, v127)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	F_ProcessSyncingRelations(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v150 = int32(0)
	F_pgstat_report_activity(m, int32(2), v150)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v150
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v150
	goto L3
L75:
	;
	v169 = F_handle_streamed_transaction(m, int32(73), l0)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v169 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v172 < int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	goto L82
L79:
	;
	v176 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v176
	goto L81
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	if base.B2i32(v180 == int32(2)) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v189 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	F_PushActiveSnapshot(m, v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v195
	v200 = m.G0
	v202 = v200 - int32(16)
	m.G0 = v202
	v205 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v207 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v210 = v207 << (uint(int32(24)) % 32)
	if v210 != int32(1308622848) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	F_logicalrep_read_tuple(m, l0, v22+int32(336))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L98
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v210 >> (uint(int32(24)) % 32)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_11), v202)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(439), int32(_a_F_apply_dispatch_12))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
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
	m.G0 = v202 + int32(16)
	v234 = F_logicalrep_rel_open(m, v205, int32(3))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L100
	}
L99:
	;
	F_logicalrep_rel_close(m, v234, v473)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L159
	}
L100:
	;
	v236 = F_should_apply_changes_for_rel(m, v234)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	if v236 == int32(0) {
		v473 = int32(3)
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+39)))
	if v242 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v234)+40))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+48))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+80))
	F_SwitchToUntrustedUser(m, v247, v22+int32(1360))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v234
	v254 = F_create_edata_for_relation(m, v234)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v234)+40))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+52))
	v260 = F_ExecInitExtraTupleSlot(m, v256, v258, int32(_a_F_apply_dispatch_13))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v256)+152))
	if v262 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v265 = F_MakePerTupleExprContext(m, v256)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	v267 = v262
	goto L111
L111:
	;
	v268 = int32(_a_F_apply_dispatch_14)
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15]))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v271
	F_slot_store_data(m, v260, v234, v22+int32(336))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L113
	}
L112:
	;
	v267 = v265
	goto L111
L113:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v234)+40))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+52))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v256)+152))
	if v280 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v283 = F_MakePerTupleExprContext(m, v256)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	v285 = v280
	goto L116
L116:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	if v279 == v286 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v285 = v283
	goto L116
L118:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v269
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v234)+40))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)+48))
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420)+119)))
	if v421 == int32(112) {
		goto L139
	} else {
		goto L140
	}
L119:
	;
	v289 = F_palloc_mul(m, int32(4), v279)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v292 = F_palloc_mul(m, int32(4), v279)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	if v279 <= int32(0) {
		goto L118
	} else {
		goto L122
	}
L122:
	;
	v297 = int32(0)
	v303 = v2
	goto L123
L123:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278+v297<<(uint(int32(3))%32))+34)))
	if v319&int32(12) != 0 {
		v350 = v303
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if v350 <= int32(0) {
		goto L118
	} else {
		goto L133
	}
L125:
	;
	v354 = v297 + int32(1)
	if v354 != v279 {
		v297 = v354
		v303 = v350
		goto L123
	} else {
		goto L132
	}
L126:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v234)+44))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v327 = int32(*(*int16)(unsafe.Add(mBase, uint32(v323+v297<<(uint(int32(1))%32)))))
	if int32(0) <= v327 {
		v350 = v303
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v234)+40))
	v333 = F_build_column_default(m, v330, v297+int32(1))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	if v333 == int32(0) {
		v350 = v303
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v338 = v303 << (uint(int32(2)) % 32)
	v340 = F_expression_planner(m, v333)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v343 = F_ExecInitExpr(m, v340, int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292+v338))) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v289+v338))) = v297
	v350 = v303 + int32(1)
	goto L125
L132:
	;
	goto L124
L133:
	;
	v359 = int32(0)
	goto L134
L134:
	;
	v379 = v359 << (uint(int32(2)) % 32)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v292+v379)))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
	v383 = v379 + v289
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v381)+24))
	v387 = m.T0[v386].(func(*base.Module, int32, int32, int32) int64)(m, v381, v285, v382+v384)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L136
	}
L135:
	;
	goto L118
L136:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	*(*int64)(unsafe.Add(mBase, uint32(v389+v390<<(uint(int32(3))%32)))) = v387
	v396 = v359 + int32(1)
	if v396 != v350 {
		v359 = v396
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	F_AfterTriggerEndQuery(m, v445)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L148
	}
L139:
	;
	F_apply_handle_tuple_routing(m, v254, v260, int32(0), int32(3))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v254)+8))
	F_ExecOpenIndices(m, v428, int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L143
	}
L142:
	;
	goto L138
L143:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	F_InitConflictIndexes(m, v428)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v428)+8))
	F_TargetPrivilegesCheck(m, v435, int64(1))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_ExecSimpleRelationInsert(m, v428, v432, v260)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_ExecCloseIndices(m, v428)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	goto L138
L148:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v254)+16))
	if v448 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	F_ExecCleanupTupleRouting(m, v449, v448)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	F_ExecCloseTrigTargetRelations(m, v445)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L153
	}
L152:
	;
	goto L151
L153:
	;
	v454 = int32(0)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v445)+104))
	F_ExecResetTupleTable(m, v455, v454)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_FreeExecutorState(m, v445)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_pfree(m, v254)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = int32(0)
	if v242 != 0 {
		v473 = v454
		goto L99
	} else {
		goto L157
	}
L157:
	;
	F_RestoreUserContext(m, v22+int32(1360))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v473 = v454
	goto L99
L159:
	;
	goto L4
L160:
	;
	v496 = F_handle_streamed_transaction(m, int32(85), l0)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	if v496 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v499 < int32(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v506)+20))
	goto L167
L164:
	;
	v503 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v503
	goto L166
L165:
	;
	goto L166
L166:
	;
	goto L163
L167:
	;
	if base.B2i32(v507 == int32(2)) == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v516 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	F_PushActiveSnapshot(m, v516)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v522
	v526 = v22 + int32(291)
	v531 = m.G0
	v533 = v531 - int32(32)
	m.G0 = v533
	v536 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L176
	}
L175:
	;
	v604 = F_logicalrep_rel_open(m, v536, int32(3))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L198
	}
L176:
	;
	v538 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L195
	}
L178:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L192
	}
L179:
	;
	if v538&int32(251) != int32(75) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v545 = v538 << (uint(int32(24)) % 32)
	if v545 != int32(1308622848) {
		goto L178
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v549 = int32(75)
	if v538&v549 == v549 {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	goto L182
L184:
	;
	v563 = v561 << (uint(int32(24)) % 32)
	if v563 != int32(1308622848) {
		goto L177
	} else {
		goto L190
	}
L185:
	;
	F_logicalrep_read_tuple(m, l0, v22+int32(304))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v559 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v526))) = uint8(v559)
	v561 = v538
	goto L184
L188:
	;
	v555 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v526))) = uint8(v555)
	v557 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v561 = v557
	goto L184
L190:
	;
	F_logicalrep_read_tuple(m, l0, v22+int32(292))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	m.G0 = v533 + int32(32)
	goto L175
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533))) = v545 >> (uint(int32(24)) % 32)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_15), v533)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(501), int32(_a_F_apply_dispatch_16))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v533)+16)) = v563 >> (uint(int32(24)) % 32)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_17), v533+int32(16))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(517), int32(_a_F_apply_dispatch_16))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
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
	v606 = F_should_apply_changes_for_rel(m, v604)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	if v606 == int32(0) {
		v5049 = int32(3)
		goto L5
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v604
	F_check_relation_updatable(m, v604)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+39)))
	if v616 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v604)+40))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)+48))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)+80))
	F_SwitchToUntrustedUser(m, v621, v22+int32(320))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v626 = F_create_edata_for_relation(m, v604)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L206
	}
L205:
	;
	goto L204
L206:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v604)+40))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v629)+52))
	v632 = F_ExecInitExtraTupleSlot(m, v628, v630, int32(_a_F_apply_dispatch_13))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v632)+12))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)))
	if int32(0) < v635 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v628)+32))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)+12))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v639)))
	v642 = int32(0)
	v645 = v634
	v646 = v635
	goto L211
L209:
	;
	goto L210
L210:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v628)+152))
	if v715 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L211:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645+v642<<(uint(int32(3))%32))+34)))
	if v664&int32(4) != 0 {
		v690 = v645
		v691 = v646
		goto L213
	} else {
		goto L214
	}
L212:
	;
	goto L210
L213:
	;
	v694 = v642 + int32(1)
	if v694 < v691 {
		v642 = v694
		v645 = v690
		v646 = v691
		goto L211
	} else {
		goto L219
	}
L214:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v604)+44))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	v672 = int32(*(*int16)(unsafe.Add(mBase, uint32(v668+v642<<(uint(int32(1))%32)))))
	if v672 < int32(0) {
		v690 = v645
		v691 = v646
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v22)+300))
	if v675 <= v672 {
		goto L27
	} else {
		goto L216
	}
L216:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v22)+296))
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677+v672))))
	if v679 == int32(117) {
		v690 = v645
		v691 = v646
		goto L213
	} else {
		goto L217
	}
L217:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v640)+36))
	v685 = F_bms_add_member(m, v682, v642+int32(8))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v640)+36)) = v685
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v632)+12))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v688)))
	v690 = v688
	v691 = v689
	goto L213
L219:
	;
	goto L212
L220:
	;
	v718 = F_MakePerTupleExprContext(m, v628)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L223
	}
L221:
	;
	v720 = v715
	goto L222
L222:
	;
	v721 = int32(_a_F_apply_dispatch_14)
	v722 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15]))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v720)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v724
	v729 = v22 + int32(292)
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+291)))
	if v730 != 0 {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v720 = v718
	goto L222
L224:
	;
	v731 = v22 + int32(304)
	goto L226
L225:
	;
	v731 = v729
	goto L226
L226:
	;
	F_slot_store_data(m, v632, v604, v731)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v722
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v604)+40))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v736)+48))
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737)+119)))
	if v738 == int32(112) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	F_apply_handle_tuple_routing(m, v626, v632, v729, int32(2))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L1
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v604)+52))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v626)+4))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v626)+8))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v746)+8))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	v749 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+1376)) = v749
	*(*int64)(unsafe.Add(mBase, uint32(v22)+1368)) = v749
	*(*int64)(unsafe.Add(mBase, uint32(v22)+1360)) = v749
	v757 = int32(0)
	F_EvalPlanQualInit(m, v22+int32(336), v748, v757, v757, int32(-1), v757)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L232
	}
L231:
	;
	goto L6
L232:
	;
	F_ExecOpenIndices(m, v746, int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	F_TargetPrivilegesCheck(m, v747, int64(2))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	v772 = F_table_slot_create(m, v747, v766+int32(104))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	if v744 != 0 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v786 = F_GetTupleTransactionInfo(m, v772, v22+int32(1368), v22+int32(1372), v22+int32(1376))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L245
	}
L237:
	;
	v774 = F_RelationFindReplTupleByIndex(m, v747, v744, v632, v772)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v776 = F_RelationFindReplTupleSeq(m, v747, v632, v772)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L1
	} else {
		goto L242
	}
L240:
	;
	if v774 != 0 {
		goto L236
	} else {
		goto L241
	}
L241:
	;
	goto L8
L242:
	;
	if v776 == int32(0) {
		goto L8
	} else {
		goto L243
	}
L243:
	;
	goto L236
L244:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v748)+152))
	if v818 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L245:
	;
	if v786 == int32(0) {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v790 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+1372)))
	v792 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_apply_dispatch[16])))
	if v790 == v792 {
		goto L244
	} else {
		goto L247
	}
L247:
	;
	v796 = F_table_slot_create(m, v747, v748+int32(104))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	F_slot_store_data(m, v796, v745, v22+int32(292))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+1360)) = v772
	v804 = v22 + int32(1360)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+316)) = v804
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v804
	v808 = int32(1)
	v812 = F_list_make1_impl(m, v808, v22+int32(56))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	F_ReportApplyConflict(m, v748, v746, int32(15), v808, v632, v796, v812)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	goto L244
L252:
	;
	v821 = F_MakePerTupleExprContext(m, v748)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L255
	}
L253:
	;
	v823 = v818
	goto L254
L254:
	;
	v824 = int32(_a_F_apply_dispatch_14)
	v825 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15]))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v823)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v827
	F_slot_modify_data(m, v632, v772, v745, v22+int32(292))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L256
	}
L255:
	;
	v823 = v821
	goto L254
L256:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v825
	*(*int32)(unsafe.Add(mBase, uint32(v22)+364)) = v632
	F_InitConflictIndexes(m, v746)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v746)+8))
	F_TargetPrivilegesCheck(m, v838, int64(4))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	F_ExecSimpleRelationUpdate(m, v746, v748, v22+int32(336), v772, v632)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	goto L7
L260:
	;
	v851 = F_handle_streamed_transaction(m, int32(68), l0)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	if v851 != 0 {
		goto L3
	} else {
		goto L262
	}
L262:
	;
	v854 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v854 < int32(0) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v861 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v861)+20))
	goto L267
L264:
	;
	v858 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v858
	goto L266
L265:
	;
	goto L266
L266:
	;
	goto L263
L267:
	;
	if base.B2i32(v862 == int32(2)) == int32(0) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v871 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L1
	} else {
		goto L273
	}
L271:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	goto L270
L273:
	;
	F_PushActiveSnapshot(m, v871)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	v877 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v877
	v882 = m.G0
	v884 = v882 - int32(16)
	m.G0 = v884
	v887 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	v889 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	if v889&int32(251) != int32(75) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	F_logicalrep_read_tuple(m, l0, v22+int32(336))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L283
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v884))) = base.I32_extend8_s(v889)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_18), v884)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(572), int32(_a_F_apply_dispatch_19))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	m.G0 = v884 + int32(16)
	v915 = F_logicalrep_rel_open(m, v887, int32(3))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L285
	}
L284:
	;
	F_logicalrep_rel_close(m, v915, v1007)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L319
	}
L285:
	;
	v917 = F_should_apply_changes_for_rel(m, v915)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	if v917 == int32(0) {
		v1007 = int32(3)
		goto L284
	} else {
		goto L287
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v915
	F_check_relation_updatable(m, v915)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	v926 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926)+39)))
	if v927 == int32(0) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v915)+40))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v930)+48))
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v931)+80))
	F_SwitchToUntrustedUser(m, v932, v22+int32(1360))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	v937 = F_create_edata_for_relation(m, v915)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L1
	} else {
		goto L293
	}
L292:
	;
	goto L291
L293:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v937)))
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v915)+40))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v940)+52))
	v943 = F_ExecInitExtraTupleSlot(m, v939, v941, int32(_a_F_apply_dispatch_13))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v939)+152))
	if v945 == int32(0) {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v948 = F_MakePerTupleExprContext(m, v939)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L1
	} else {
		goto L298
	}
L296:
	;
	v950 = v945
	goto L297
L297:
	;
	v951 = int32(_a_F_apply_dispatch_14)
	v952 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15]))
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v950)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v954
	F_slot_store_data(m, v943, v915, v22+int32(336))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L299
	}
L298:
	;
	v950 = v948
	goto L297
L299:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v952
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v915)+40))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v962)+48))
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963)+119)))
	if v964 == int32(112) {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v937)))
	F_AfterTriggerEndQuery(m, v981)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L1
	} else {
		goto L308
	}
L301:
	;
	F_apply_handle_tuple_routing(m, v937, v943, int32(0), int32(4))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v937)+8))
	F_ExecOpenIndices(m, v971, int32(0))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L305
	}
L304:
	;
	goto L300
L305:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v915)+52))
	F_apply_handle_delete_internal(m, v937, v971, v943, v975)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_ExecCloseIndices(m, v971)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	goto L300
L308:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v937)+16))
	if v984 != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v937)+12))
	F_ExecCleanupTupleRouting(m, v985, v984)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	F_ExecCloseTrigTargetRelations(m, v981)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L313
	}
L312:
	;
	goto L311
L313:
	;
	v990 = int32(0)
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v981)+104))
	F_ExecResetTupleTable(m, v991, v990)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	F_FreeExecutorState(m, v981)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	F_pfree(m, v937)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = int32(0)
	if v927 != 0 {
		v1007 = v990
		goto L284
	} else {
		goto L317
	}
L317:
	;
	F_RestoreUserContext(m, v22+int32(1360))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	v1007 = v990
	goto L284
L319:
	;
	goto L4
L320:
	;
	v1022 = F_handle_streamed_transaction(m, int32(84), l0)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	if v1022 != 0 {
		goto L3
	} else {
		goto L322
	}
L322:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v1025 < int32(0) {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+20))
	goto L327
L324:
	;
	v1029 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v1029
	goto L326
L325:
	;
	goto L326
L326:
	;
	goto L323
L327:
	;
	if base.B2i32(v1033 == int32(2)) == int32(0) {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	v1042 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L333
	}
L331:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	goto L330
L333:
	;
	F_PushActiveSnapshot(m, v1042)
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v1048
	v1055 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	v1058 = F_pq_getmsgint(m, l0, int32(1))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	v1060 = int32(1)
	v1061 = v1058 & v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(336)))) = uint8(v1061)
	v1066 = int32(base.Ui32(v1058)>>(uint(v1060)%32)) & v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(1360)))) = uint8(v1066)
	v1068 = int32(0)
	if v1068 < v1055 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1073 = v1068
	v1076 = int32(0)
	goto L340
L338:
	;
	v1100 = v1068
	goto L339
L339:
	;
	if v1100 == int32(0) {
		v4858 = v2
		v4860 = v2
		v4861 = v2
		v4863 = v2
		v4864 = v2
		goto L9
	} else {
		goto L345
	}
L340:
	;
	v1092 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L342
	}
L341:
	;
	v1100 = v1094
	goto L339
L342:
	;
	v1094 = F_lappend_oid(m, v1073, v1092)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	v1097 = v1076 + int32(1)
	if v1097 != v1055 {
		v1073 = v1094
		v1076 = v1097
		goto L340
	} else {
		goto L344
	}
L344:
	;
	goto L341
L345:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+4))
	if v1120 <= int32(0) {
		v4858 = v2
		v4860 = v2
		v4861 = v2
		v4863 = v2
		v4864 = v2
		goto L9
	} else {
		goto L346
	}
L346:
	;
	v1126 = v2
	v1128 = v2
	v1129 = v2
	v1130 = v2
	v1131 = v2
	v1132 = v2
	goto L347
L347:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+12))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1142+v1130<<(uint(int32(2))%32))))
	v1148 = F_logicalrep_rel_open(m, v1146, int32(8))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L1
	} else {
		goto L350
	}
L348:
	;
	v4858 = v1339
	v4860 = v1341
	v4861 = v1342
	v4863 = v1344
	v4864 = v1345
	goto L9
L349:
	;
	v1356 = v1130 + int32(1)
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+4))
	if v1356 < v1357 {
		v1126 = v1339
		v1128 = v1341
		v1129 = v1342
		v1130 = v1356
		v1131 = v1344
		v1132 = v1345
		goto L347
	} else {
		goto L420
	}
L350:
	;
	v1150 = F_should_apply_changes_for_rel(m, v1148)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	if v1150 == int32(0) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	F_logicalrep_rel_close(m, v1148, int32(8))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L1
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v1157 = F_lappend(m, v1131, v1148)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L1
	} else {
		goto L356
	}
L355:
	;
	v1339 = v1126
	v1341 = v1128
	v1342 = v1129
	v1344 = v1131
	v1345 = v1132
	goto L349
L356:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+40))
	F_TargetPrivilegesCheck(m, v1159, int64(16))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+40))
	v1164 = F_lappend(m, v1128, v1163)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L1
	} else {
		goto L358
	}
L358:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+36))
	v1167 = F_lappend_oid(m, v1126, v1166)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L359
	}
L359:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[17]))
	if v1170 <= int32(1) {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+40))
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+48))
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1201)+119)))
	if v1202 != int32(112) {
		v1339 = v1167
		v1341 = v1164
		v1342 = v1129
		v1344 = v1157
		v1345 = v1198
		goto L349
	} else {
		goto L375
	}
L361:
	;
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[18])))
	if v1174&int32(1) == int32(0) {
		v1198 = v1132
		goto L360
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+40))
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1179)+48))
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180)+118)))
	if v1181 != int32(112) {
		v1198 = v1132
		goto L360
	} else {
		goto L365
	}
L364:
	;
	goto L363
L365:
	;
	if v1170 <= int32(0) {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1179)+32))
	if v1186 != 0 {
		v1198 = v1132
		goto L360
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	v1188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180)+119)))
	if v1188 == int32(102) {
		v1198 = v1132
		goto L360
	} else {
		goto L371
	}
L369:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1179)+40))
	if v1187 != 0 {
		v1198 = v1132
		goto L360
	} else {
		goto L370
	}
L370:
	;
	goto L368
L371:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1179)+56))
	goto L372
L372:
	;
	if base.Ui32(v1191) < base.Ui32(int32(_a_F_apply_dispatch_20)) {
		v1198 = v1132
		goto L360
	} else {
		goto L373
	}
L373:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+36))
	v1195 = F_lappend_oid(m, v1132, v1194)
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L1
	} else {
		goto L374
	}
L374:
	;
	v1198 = v1195
	goto L360
L375:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+36))
	v1208 = F_find_all_inheritors(m, v1205, int32(8), int32(0))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L1
	} else {
		goto L376
	}
L376:
	;
	if v1208 == int32(0) {
		v1339 = v1167
		v1341 = v1164
		v1342 = v1129
		v1344 = v1157
		v1345 = v1198
		goto L349
	} else {
		goto L377
	}
L377:
	;
	v1212 = int32(0)
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+4))
	if v1213 <= v1212 {
		v1339 = v1167
		v1341 = v1164
		v1342 = v1129
		v1344 = v1157
		v1345 = v1198
		goto L349
	} else {
		goto L378
	}
L378:
	;
	v1216 = v1212
	v1219 = v1167
	v1221 = v1164
	v1222 = v1129
	v1225 = v1198
	goto L379
L379:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+12))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1235+v1216<<(uint(int32(2))%32))))
	v1240 = int32(0)
	if v1219 == v1240 {
		goto L383
	} else {
		goto L384
	}
L380:
	;
	v1339 = v1325
	v1341 = v1327
	v1342 = v1328
	v1344 = v1157
	v1345 = v1329
	goto L349
L381:
	;
	v1333 = v1216 + int32(1)
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+4))
	if v1333 < v1334 {
		v1216 = v1333
		v1219 = v1325
		v1221 = v1327
		v1222 = v1328
		v1225 = v1329
		goto L379
	} else {
		goto L419
	}
L382:
	;
	if v1278 != 0 {
		v1325 = v1219
		v1327 = v1221
		v1328 = v1222
		v1329 = v1225
		goto L381
	} else {
		goto L395
	}
L383:
	;
	v1278 = int32(0)
	goto L382
L384:
	;
	goto L385
L385:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1219)+4))
	if v1246 <= int32(0) {
		v1272 = v1240
		goto L386
	} else {
		goto L387
	}
L386:
	;
	v1278 = v1272
	goto L382
L387:
	;
	v1249 = int32(0)
	if v1249 < v1246 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v1252 = v1246
	goto L390
L389:
	;
	v1252 = v1249
	goto L390
L390:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1219)+12))
	v1255 = int32(0)
	goto L391
L391:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1253+v1255<<(uint(int32(2))%32))))
	v1264 = base.B2i32(v1263 == v1239)
	if v1263 == v1239 {
		v1272 = v1264
		goto L386
	} else {
		goto L393
	}
L392:
	;
	v1272 = v1264
	goto L386
L393:
	;
	v1266 = v1255 + int32(1)
	if v1266 != v1252 {
		v1255 = v1266
		goto L391
	} else {
		goto L394
	}
L394:
	;
	goto L392
L395:
	;
	v1280 = F_table_open(m, v1239, int32(0))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L1
	} else {
		goto L397
	}
L396:
	;
	F_TargetPrivilegesCheck(m, v1280, int64(16))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L1
	} else {
		goto L401
	}
L397:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+48))
	v1283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1282)+118)))
	if v1283 != int32(116) {
		goto L396
	} else {
		goto L398
	}
L398:
	;
	v1286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280)+24)))
	if v1286 != 0 {
		goto L396
	} else {
		goto L399
	}
L399:
	;
	F_relation_close(m, v1280, int32(8))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
	;
	v1325 = v1219
	v1327 = v1221
	v1328 = v1222
	v1329 = v1225
	goto L381
L401:
	;
	v1293 = F_lappend(m, v1221, v1280)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	v1295 = F_lappend(m, v1222, v1280)
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	v1297 = F_lappend_oid(m, v1219, v1239)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[17]))
	if v1300 <= int32(1) {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[18])))
	if v1304&int32(1) == int32(0) {
		v1325 = v1297
		v1327 = v1293
		v1328 = v1295
		v1329 = v1225
		goto L381
	} else {
		goto L408
	}
L406:
	;
	goto L407
L407:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+48))
	v1310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1309)+118)))
	if v1310 != int32(112) {
		v1325 = v1297
		v1327 = v1293
		v1328 = v1295
		v1329 = v1225
		goto L381
	} else {
		goto L409
	}
L408:
	;
	goto L407
L409:
	;
	if v1300 <= int32(0) {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+32))
	if v1315 != 0 {
		v1325 = v1297
		v1327 = v1293
		v1328 = v1295
		v1329 = v1225
		goto L381
	} else {
		goto L413
	}
L411:
	;
	goto L412
L412:
	;
	v1317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1309)+119)))
	if v1317 == int32(102) {
		v1325 = v1297
		v1327 = v1293
		v1328 = v1295
		v1329 = v1225
		goto L381
	} else {
		goto L415
	}
L413:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+40))
	if v1316 != 0 {
		v1325 = v1297
		v1327 = v1293
		v1328 = v1295
		v1329 = v1225
		goto L381
	} else {
		goto L414
	}
L414:
	;
	goto L412
L415:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+56))
	goto L416
L416:
	;
	if base.Ui32(v1320) < base.Ui32(int32(_a_F_apply_dispatch_20)) {
		v1325 = v1297
		v1327 = v1293
		v1328 = v1295
		v1329 = v1225
		goto L381
	} else {
		goto L417
	}
L417:
	;
	v1323 = F_lappend_oid(m, v1225, v1239)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L1
	} else {
		goto L418
	}
L418:
	;
	v1325 = v1297
	v1327 = v1293
	v1328 = v1295
	v1329 = v1323
	goto L381
L419:
	;
	goto L380
L420:
	;
	goto L348
L421:
	;
	if v1360 != 0 {
		goto L3
	} else {
		goto L422
	}
L422:
	;
	v1363 = F_palloc(m, int32(32))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L1
	} else {
		goto L423
	}
L423:
	;
	v1366 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1363))) = v1366
	v1369 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	v1372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1369))))
	if v1372 != 0 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v1373 = v1369
	goto L428
L427:
	;
	v1373 = int32(_a_F_apply_dispatch_21)
	goto L428
L428:
	;
	v1374 = F_pstrdup(m, v1373)
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L1
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1363)+4)) = v1374
	v1377 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	v1379 = F_pstrdup(m, v1377)
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1363)+8)) = v1379
	v1382 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L1
	} else {
		goto L432
	}
L432:
	;
	v1384 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1363)+25)) = uint8(v1384)
	*(*uint8)(unsafe.Add(mBase, uint32(v1363)+24)) = uint8(v1382)
	v1389 = F_pq_getmsgint(m, l0, int32(2))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L1
	} else {
		goto L433
	}
L433:
	;
	v1391 = F_palloc_mul(m, int32(4), v1389)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L434
	}
L434:
	;
	v1394 = F_palloc_mul(m, int32(4), v1389)
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L1
	} else {
		goto L435
	}
L435:
	;
	if int32(0) < v1389 {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v1402 = v2
	v1403 = int32(0)
	goto L439
L437:
	;
	v1447 = v2
	goto L438
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1363)+28)) = v1447
	*(*int32)(unsafe.Add(mBase, uint32(v1363)+20)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v1363)+16)) = v1391
	*(*int32)(unsafe.Add(mBase, uint32(v1363)+12)) = v1389
	F_logicalrep_relmap_update(m, v1363)
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L1
	} else {
		goto L451
	}
L439:
	;
	v1418 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L1
	} else {
		goto L441
	}
L440:
	;
	v1447 = v1424
	goto L438
L441:
	;
	if v1418&int32(1) != 0 {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v1422 = F_bms_add_member(m, v1402, v1403)
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L1
	} else {
		goto L445
	}
L443:
	;
	v1424 = v1402
	goto L444
L444:
	;
	v1426 = v1403 << (uint(int32(2)) % 32)
	v1428 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L1
	} else {
		goto L446
	}
L445:
	;
	v1424 = v1422
	goto L444
L446:
	;
	v1430 = F_pstrdup(m, v1428)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L447
	}
L447:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1391+v1426))) = v1430
	v1435 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1426+v1394))) = v1435
	v1439 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L1
	} else {
		goto L449
	}
L449:
	;
	v1442 = v1403 + int32(1)
	if v1442 != v1389 {
		v1402 = v1424
		v1403 = v1442
		goto L439
	} else {
		goto L450
	}
L450:
	;
	goto L440
L451:
	;
	v1469 = m.G0
	v1471 = v1469 - int32(32)
	m.G0 = v1471
	v1474 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[19]))
	if v1474 == int32(0) {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	m.G0 = v1471 + int32(32)
	goto L3
L453:
	;
	v1478 = v1471 + int32(12)
	F_hash_seq_init(m, v1478, v1474)
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L1
	} else {
		goto L454
	}
L454:
	;
	v1481 = F_hash_seq_search(m, v1478)
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	if v1481 == int32(0) {
		goto L452
	} else {
		goto L456
	}
L456:
	;
	v1486 = v1481
	goto L457
L457:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+8))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1363)))
	if v1504 == v1505 {
		goto L459
	} else {
		goto L460
	}
L458:
	;
	goto L452
L459:
	;
	v1508 = v1486 + int32(8)
	F_logicalrep_relmap_free_entry(m, v1508)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L1
	} else {
		goto L462
	}
L460:
	;
	goto L461
L461:
	;
	v1517 = F_hash_seq_search(m, v1471+int32(12))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L1
	} else {
		goto L463
	}
L462:
	;
	base.MemoryFill(m, v1508, int32(0), int32(72))
	goto L461
L463:
	;
	if v1517 != 0 {
		v1486 = v1517
		goto L457
	} else {
		goto L464
	}
L464:
	;
	goto L458
L465:
	;
	if v1542 != 0 {
		goto L3
	} else {
		goto L466
	}
L466:
	;
	v1545 = v22 + int32(336)
	v1547 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L1
	} else {
		goto L467
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1545))) = v1547
	v1550 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L1
	} else {
		goto L468
	}
L468:
	;
	v1553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1550))))
	if v1553 != 0 {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v1554 = v1550
	goto L471
L470:
	;
	v1554 = int32(_a_F_apply_dispatch_21)
	goto L471
L471:
	;
	v1555 = F_pstrdup(m, v1554)
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L1
	} else {
		goto L472
	}
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1545)+4)) = v1555
	v1558 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L1
	} else {
		goto L473
	}
L473:
	;
	v1560 = F_pstrdup(m, v1558)
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L1
	} else {
		goto L474
	}
L474:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1545)+8)) = v1560
	goto L3
L475:
	;
	v1566 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])))
	if v1566 != int32(1) {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L1
	} else {
		goto L482
	}
L477:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1570)+20))
	goto L478
L478:
	;
	if base.B2i32(v1571 == int32(2)) == int32(0) {
		goto L3
	} else {
		goto L479
	}
L479:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v1578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1577)+16)))
	if v1578 != int32(1) {
		goto L476
	} else {
		goto L480
	}
L480:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1577)))
	if v1581 == int32(1) {
		goto L3
	} else {
		goto L481
	}
L481:
	;
	goto L476
L482:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L1
	} else {
		goto L483
	}
L483:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_22), int32(0))
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L1
	} else {
		goto L484
	}
L484:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1703), int32(_a_F_apply_dispatch_23))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
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
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1607 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])) = uint8(v1607)
	v1612 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L1
	} else {
		goto L487
	}
L487:
	;
	v1614 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(320)))) = uint8(base.B2i32(v1614 == int32(1)))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[20])) = v1612
	if v1612 == int32(0) {
		goto L25
	} else {
		goto L489
	}
L489:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v1612
	v1628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+320)))
	if v1628 == int32(1) {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v1631 = m.G0
	v1633 = v1631 + int32(-64)
	m.G0 = v1633
	*(*int32)(unsafe.Add(mBase, uint32(v1633)+60)) = v1612
	v1637 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1637)))
	if v1638 != int32(3) {
		goto L495
	} else {
		goto L496
	}
L491:
	;
	v1988 = v1612
	goto L492
L492:
	;
	v2007 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2007)+16)))
	if v2008 == int32(1) {
		goto L556
	} else {
		goto L557
	}
L493:
	;
	v1986 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[20]))
	v1988 = v1986
	goto L492
L494:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L1
	} else {
		goto L551
	}
L495:
	;
	m.G0 = v1633 - int32(-64)
	goto L493
L496:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v1645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1644)+68)))
	if v1645 != int32(1) {
		goto L495
	} else {
		goto L498
	}
L498:
	;
	v1649 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v1650 = *(*int64)(unsafe.Add(mBase, uint32(v1649)+16))
	if v1650 != int64(0) {
		goto L495
	} else {
		goto L499
	}
L499:
	;
	v1653 = F_AllTablesyncsReady(m)
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	if v1653 == int32(0) {
		goto L495
	} else {
		goto L501
	}
L501:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[21]))
	if v1658 == int32(0) {
		goto L503
	} else {
		goto L504
	}
L502:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[22]))
	if v1892 == int32(0) {
		goto L541
	} else {
		goto L542
	}
L503:
	;
	v1714 = int32(_a_F_apply_dispatch_14)
	v1715 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15]))
	v1718 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v1718
	v1721 = F_palloc0(m, int32(20))
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L1
	} else {
		goto L510
	}
L504:
	;
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v1658)+4))
	if v1661 <= int32(0) {
		goto L503
	} else {
		goto L505
	}
L505:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1658)+12))
	v1667 = int32(0)
	goto L506
L506:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v1664+v1667<<(uint(int32(2))%32))))
	v1689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1688)+13)))
	if v1689 == int32(0) {
		v1875 = v1688
		goto L502
	} else {
		goto L508
	}
L507:
	;
	goto L503
L508:
	;
	v1693 = v1667 + int32(1)
	if v1661 != v1693 {
		v1667 = v1693
		goto L506
	} else {
		goto L509
	}
L509:
	;
	goto L507
L510:
	;
	v1725 = F_add_size(m, int32(0), int32(96))
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	v1728 = F_add_size(m, v1725, int32(16777216))
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v1731 = F_add_size(m, v1728, int32(_a_F_apply_dispatch_24))
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L1
	} else {
		goto L513
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1633)+8)) = v1731
	v1736 = F_add_size(m, int32(0), int32(3))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1633)+12)) = v1736
	v1740 = v1631 + int32(-56)
	v1741 = F_shm_toc_estimate(m, v1740)
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	v1743 = F_shm_toc_estimate(m, v1740)
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	v1746 = F_dsm_create(m, v1743, int32(0))
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	if v1746 == int32(0) {
		goto L518
	} else {
		goto L519
	}
L518:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v1715
	F_pfree(m, v1721)
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L1
	} else {
		goto L521
	}
L519:
	;
	goto L520
L520:
	;
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1746)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1755))) = int64(2021433447)
	v1757 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1755)+8)), uint32(v1757))
	*(*int64)(unsafe.Add(mBase, uint32(v1755)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1755)+12)) = v1741 & int32(-32)
	goto L522
L521:
	;
	goto L495
L522:
	;
	v1766 = F_shm_toc_allocate(m, v1755, int32(80))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L1
	} else {
		goto L523
	}
L523:
	;
	v1768 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1766))), uint32(v1768))
	*(*int32)(unsafe.Add(mBase, uint32(v1766)+20)) = v1768
	*(*int32)(unsafe.Add(mBase, uint32(v1766)+8)) = v1768
	*(*int32)(unsafe.Add(mBase, uint32(v1766)+32)) = v1768
	*(*int64)(unsafe.Add(mBase, uint32(v1766)+24)) = int64(0)
	F_shm_toc_insert(m, v1755, int64(1), v1766)
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	v1784 = F_shm_toc_allocate(m, v1755, int32(16777216))
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	v1787 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1784))), uint32(v1787))
	v1790 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1784)+16)) = v1790
	*(*int64)(unsafe.Add(mBase, uint32(v1784)+4)) = v1790
	v1794 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v1784)+36)) = uint16(v1794)
	*(*int64)(unsafe.Add(mBase, uint32(v1784)+24)) = v1790
	*(*int32)(unsafe.Add(mBase, uint32(v1784)+32)) = int32(16777176)
	goto L526
L526:
	;
	F_shm_toc_insert(m, v1755, int64(2), v1784)
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L1
	} else {
		goto L527
	}
L527:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24]))
	F_shm_mq_set_sender(m, v1784, v1806)
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L1
	} else {
		goto L528
	}
L528:
	;
	v1809 = F_shm_mq_attach(m, v1784, v1746)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1721))) = v1809
	v1814 = F_shm_toc_allocate(m, v1755, int32(_a_F_apply_dispatch_24))
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L1
	} else {
		goto L530
	}
L530:
	;
	v1817 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1814))), uint32(v1817))
	v1820 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1814)+16)) = v1820
	*(*int64)(unsafe.Add(mBase, uint32(v1814)+4)) = v1820
	v1824 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v1814)+36)) = uint16(v1824)
	*(*int64)(unsafe.Add(mBase, uint32(v1814)+24)) = v1820
	*(*int32)(unsafe.Add(mBase, uint32(v1814)+32)) = int32(_a_F_apply_dispatch_25)
	goto L531
L531:
	;
	F_shm_toc_insert(m, v1755, int64(3), v1814)
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24]))
	F_shm_mq_set_receiver(m, v1814, v1836)
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L1
	} else {
		goto L533
	}
L533:
	;
	v1839 = F_shm_mq_attach(m, v1814, v1746)
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L1
	} else {
		goto L534
	}
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1721)+16)) = v1766
	*(*int32)(unsafe.Add(mBase, uint32(v1721)+8)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v1721)+4)) = v1839
	v1846 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1846)+24))
	v1849 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v1849)+4))
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v1849)+24))
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1846)+28))
	v1853 = int32(0)
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v1746)+12))
	v1856 = F_logicalrep_worker_launch(m, int32(4), v1847, v1850, v1851, v1852, v1853, v1854, v1853)
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L1
	} else {
		goto L535
	}
L535:
	;
	if v1856 == int32(0) {
		goto L536
	} else {
		goto L537
	}
L536:
	;
	F_pa_free_worker_info(m, v1721)
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L1
	} else {
		goto L539
	}
L537:
	;
	goto L538
L538:
	;
	v1865 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[21]))
	v1866 = F_lappend(m, v1865, v1721)
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L1
	} else {
		goto L540
	}
L539:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v1715
	goto L495
L540:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v1715
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[21])) = v1866
	v1875 = v1721
	goto L502
L541:
	;
	v1895 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1633)+40)) = v1895
	*(*int64)(unsafe.Add(mBase, uint32(v1633)+48)) = v1895
	*(*int64)(unsafe.Add(mBase, uint32(v1633)+32)) = v1895
	*(*int64)(unsafe.Add(mBase, uint32(v1633)+24)) = v1895
	*(*int64)(unsafe.Add(mBase, uint32(v1633)+8)) = v1895
	*(*int64)(unsafe.Add(mBase, uint32(v1633)+16)) = int64(34359738372)
	v1908 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v1633)+44)) = v1908
	v1916 = F_hash_create(m, int32(_a_F_apply_dispatch_26), int64(16), v1631+int32(-56), int32(1064))
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L1
	} else {
		goto L544
	}
L542:
	;
	v1919 = v1892
	goto L543
L543:
	;
	v1925 = F_hash_search(m, v1919, v1631+int32(-4), int32(1), v1631+int32(-56))
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L1
	} else {
		goto L545
	}
L544:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[22])) = v1916
	v1919 = v1916
	goto L543
L545:
	;
	v1927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1633)+8)))
	if v1927 == int32(1) {
		goto L494
	} else {
		goto L546
	}
L546:
	;
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+16))
	v1933 = base.AtomicRmwXchg32(m, v1930, int32(0), int32(1))
	if v1933 != 0 {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	F_s_lock(m, v1930, int32(_a_F_apply_dispatch_27))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L1
	} else {
		goto L550
	}
L548:
	;
	goto L549
L549:
	;
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+16))
	v1938 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1937)+8)) = v1938
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+16))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1633)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1940)+4)) = v1941
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+16))
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1943))), uint32(v1938))
	v1947 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v1875)+12)) = uint16(v1947)
	*(*int32)(unsafe.Add(mBase, uint32(v1925)+4)) = v1875
	goto L495
L550:
	;
	goto L549
L551:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_28), int32(0))
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L1
	} else {
		goto L552
	}
L552:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_29), int32(505), int32(_a_F_apply_dispatch_30))
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L554:
	;
	F_pgstat_report_activity(m, int32(3), int32(0))
	mBase = m.M
	goto L3
L555:
	;
	v2087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+320)))
	if v2087 == int32(1) {
		goto L584
	} else {
		goto L585
	}
L556:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v2007)))
	if v2011 == int32(4) {
		goto L555
	} else {
		goto L559
	}
L557:
	;
	goto L558
L558:
	;
	v2014 = F_pa_find_worker(m, v1988)
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L1
	} else {
		goto L562
	}
L559:
	;
	goto L558
L560:
	;
	v2059 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+1360)) = uint8(v2059)
	v2061 = v1604 - v1603
	*(*int32)(unsafe.Add(mBase, uint32(v22)+336)) = v2061 + int32(1)
	v2066 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25]))
	F_BufFileWrite(m, v2066, v22+int32(336), int32(4))
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L1
	} else {
		goto L580
	}
L561:
	;
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2035 = F_pa_send_data(m, v2014, v2033, v2034)
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L1
	} else {
		goto L570
	}
L562:
	;
	if v2014 != 0 {
		goto L563
	} else {
		goto L564
	}
L563:
	;
	v2016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2014)+12)))
	if v2016 == int32(0) {
		goto L561
	} else {
		goto L566
	}
L564:
	;
	goto L565
L565:
	;
	v2025 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v2025 == int32(0) {
		goto L24
	} else {
		goto L568
	}
L566:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[20]))
	v2021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+320)))
	F_stream_start_internal(m, v2020, v2021)
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L1
	} else {
		goto L567
	}
L567:
	;
	goto L560
L568:
	;
	v2029 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[20]))
	v2030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+320)))
	F_stream_start_internal(m, v2029, v2030)
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L1
	} else {
		goto L569
	}
L569:
	;
	goto L554
L570:
	;
	v2037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+320)))
	if v2035 != 0 {
		goto L571
	} else {
		goto L572
	}
L571:
	;
	if v2037&int32(1) == int32(0) {
		goto L574
	} else {
		goto L575
	}
L572:
	;
	goto L573
L573:
	;
	F_pa_switch_to_partial_serialize(m, v2014, (v2037^int32(-1))&int32(1))
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L1
	} else {
		goto L579
	}
L574:
	;
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v2014)+16))
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+4))
	F_pa_unlock_stream(m, v2043)
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L1
	} else {
		goto L577
	}
L575:
	;
	goto L576
L576:
	;
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(v2014)+16))
	v2049 = base.AtomicRmwAdd32(m, v2046, int32(20), int32(1))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[26])) = v2014
	goto L578
L577:
	;
	goto L576
L578:
	;
	goto L554
L579:
	;
	goto L560
L580:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25]))
	F_BufFileWrite(m, v2073, v22+int32(1360), int32(1))
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L1
	} else {
		goto L581
	}
L581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+336)) = v2061
	v2081 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25]))
	F_BufFileWrite(m, v2081, v1603+v1605, v2061)
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L1
	} else {
		goto L582
	}
L582:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[26])) = v2014
	goto L583
L583:
	;
	goto L554
L584:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v2091)+32))
	v2094 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[27]))
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(v2094)+4))
	F_LockApplyTransactionForSession(m, v2092, v2095, int32(1), int32(8))
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L1
	} else {
		goto L587
	}
L585:
	;
	goto L586
L586:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[28])) = int32(0)
	goto L554
L587:
	;
	v2101 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[27]))
	F_pa_set_xact_state(m, v2101, int32(1))
	mBase = m.M
	v2104 = m.ExcPending
	if v2104 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v2106)+32))
	F_logicalrep_worker_wakeup(m, v2107)
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	goto L586
L590:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[20]))
	v2125 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2125)+16)))
	if v2126 == int32(1) {
		goto L593
	} else {
		goto L594
	}
L591:
	;
	v2252 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[20])) = v2252
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])) = uint8(v2252)
	v2260 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(v2260)+24))
	goto L629
L592:
	;
	v2232 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2233 = m.ExcPending
	if v2233 != 0 {
		goto L1
	} else {
		goto L622
	}
L593:
	;
	v2129 = *(*int32)(unsafe.Add(mBase, uint32(v2125)))
	if v2129 == int32(4) {
		goto L592
	} else {
		goto L596
	}
L594:
	;
	goto L595
L595:
	;
	v2132 = F_pa_find_worker(m, v2123)
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L1
	} else {
		goto L599
	}
L596:
	;
	goto L595
L597:
	;
	v2175 = int32(69)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+1360)) = uint8(v2175)
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+336)) = v2177 - v2178 + int32(1)
	v2184 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25]))
	F_BufFileWrite(m, v2184, v22+int32(336), int32(4))
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
		goto L1
	} else {
		goto L614
	}
L598:
	;
	F_pa_switch_to_partial_serialize(m, v2132, int32(1))
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L1
	} else {
		goto L613
	}
L599:
	;
	if v2132 != 0 {
		goto L600
	} else {
		goto L601
	}
L600:
	;
	v2134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2132)+12)))
	if v2134 != 0 {
		goto L597
	} else {
		goto L603
	}
L601:
	;
	goto L602
L602:
	;
	v2149 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v2149 == int32(0) {
		goto L22
	} else {
		goto L608
	}
L603:
	;
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+16))
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+4))
	F_pa_lock_stream(m, v2136)
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L1
	} else {
		goto L604
	}
L604:
	;
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2141 = F_pa_send_data(m, v2132, v2139, v2140)
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L1
	} else {
		goto L605
	}
L605:
	;
	if v2141 == int32(0) {
		goto L598
	} else {
		goto L606
	}
L606:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[26])) = int32(0)
	goto L607
L607:
	;
	goto L591
L608:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v2153)+32))
	v2156 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[20]))
	F_subxact_info_write(m, v2154, v2156)
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L1
	} else {
		goto L609
	}
L609:
	;
	v2160 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25]))
	F_BufFileClose(m, v2160)
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25])) = int32(0)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L1
	} else {
		goto L611
	}
L611:
	;
	v2169 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[29]))
	F_MemoryContextReset(m, v2169)
	mBase = m.M
	v2171 = m.ExcPending
	if v2171 != 0 {
		goto L1
	} else {
		goto L612
	}
L612:
	;
	goto L591
L613:
	;
	goto L597
L614:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25]))
	F_BufFileWrite(m, v2191, v22+int32(1360), int32(1))
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2199 = v2197 - v2198
	*(*int32)(unsafe.Add(mBase, uint32(v22)+336)) = v2199
	v2202 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25]))
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_BufFileWrite(m, v2202, v2198+v2203, v2199)
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	v2208 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+32))
	v2211 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[20]))
	F_subxact_info_write(m, v2209, v2211)
	mBase = m.M
	v2213 = m.ExcPending
	if v2213 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	v2215 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25]))
	F_BufFileClose(m, v2215)
	mBase = m.M
	v2217 = m.ExcPending
	if v2217 != 0 {
		goto L1
	} else {
		goto L618
	}
L618:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25])) = int32(0)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L1
	} else {
		goto L619
	}
L619:
	;
	v2224 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[29]))
	F_MemoryContextReset(m, v2224)
	mBase = m.M
	v2226 = m.ExcPending
	if v2226 != 0 {
		goto L1
	} else {
		goto L620
	}
L620:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[26])) = int32(0)
	goto L621
L621:
	;
	goto L591
L622:
	;
	if v2232 != 0 {
		goto L623
	} else {
		goto L624
	}
L623:
	;
	v2235 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v2235
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_31), v22+int32(96))
	mBase = m.M
	v2241 = m.ExcPending
	if v2241 != 0 {
		goto L1
	} else {
		goto L626
	}
L624:
	;
	goto L625
L625:
	;
	F_pa_decr_and_wait_stream_block(m)
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L1
	} else {
		goto L628
	}
L626:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1961), int32(_a_F_apply_dispatch_32))
	mBase = m.M
	v2246 = m.ExcPending
	if v2246 != 0 {
		goto L1
	} else {
		goto L627
	}
L627:
	;
	goto L625
L628:
	;
	goto L591
L629:
	;
	if v2261 != v2252 {
		goto L630
	} else {
		goto L631
	}
L630:
	;
	v2264 = int32(4)
	goto L632
L631:
	;
	v2264 = int32(2)
	goto L632
L632:
	;
	v2265 = int32(0)
	F_pgstat_report_activity(m, v2264, v2265)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v2265
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v2265
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v2265
	goto L3
L633:
	;
	v2289 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2289)+68)))
	v2292 = v22 + int32(1360)
	v2294 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v2295 = m.ExcPending
	if v2295 != 0 {
		goto L1
	} else {
		goto L634
	}
L634:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2292))) = v2294
	v2298 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L1
	} else {
		goto L635
	}
L635:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2292)+4)) = v2298
	if v2290 != 0 {
		goto L637
	} else {
		goto L638
	}
L636:
	;
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(v22)+1364))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v2312
	v2315 = *(*int64)(unsafe.Add(mBase, uint32(v22)+1368))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v2315
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(v22)+1360))
	v2319 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2319)+16)))
	if v2320 == int32(1) {
		goto L645
	} else {
		goto L646
	}
L637:
	;
	v2301 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2302 = m.ExcPending
	if v2302 != 0 {
		goto L1
	} else {
		goto L640
	}
L638:
	;
	goto L639
L639:
	;
	v2307 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2292)+8)) = v2307
	*(*int64)(unsafe.Add(mBase, uint32(v2292)+16)) = v2307
	goto L636
L640:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2292)+8)) = v2301
	v2304 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2305 = m.ExcPending
	if v2305 != 0 {
		goto L1
	} else {
		goto L641
	}
L641:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2292)+16)) = v2304
	goto L636
L642:
	;
	F_pa_unlock_stream(m, v2317)
	mBase = m.M
	v2924 = m.ExcPending
	if v2924 != 0 {
		goto L1
	} else {
		goto L795
	}
L643:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2910 = m.ExcPending
	if v2910 != 0 {
		goto L1
	} else {
		goto L792
	}
L644:
	;
	if v2312 != v2317 {
		goto L741
	} else {
		goto L742
	}
L645:
	;
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(v2319)))
	if v2323 == int32(4) {
		goto L644
	} else {
		goto L648
	}
L646:
	;
	goto L647
L647:
	;
	v2326 = F_pa_find_worker(m, v2317)
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		goto L1
	} else {
		goto L649
	}
L648:
	;
	goto L647
L649:
	;
	if v2326 != 0 {
		goto L650
	} else {
		goto L651
	}
L650:
	;
	v2328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2326)+12)))
	if v2328 != 0 {
		goto L11
	} else {
		goto L653
	}
L651:
	;
	goto L652
L652:
	;
	v2340 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v2340 != 0 {
		goto L643
	} else {
		goto L658
	}
L653:
	;
	if v2312 != v2317 {
		goto L642
	} else {
		goto L654
	}
L654:
	;
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2331 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2332 = F_pa_send_data(m, v2326, v2330, v2331)
	mBase = m.M
	v2333 = m.ExcPending
	if v2333 != 0 {
		goto L1
	} else {
		goto L655
	}
L655:
	;
	if v2332 == int32(0) {
		goto L12
	} else {
		goto L656
	}
L656:
	;
	F_pa_xact_finish(m, v2326, int64(0))
	mBase = m.M
	v2338 = m.ExcPending
	if v2338 != 0 {
		goto L1
	} else {
		goto L657
	}
L657:
	;
	goto L10
L658:
	;
	if v2312 == v2317 {
		goto L660
	} else {
		goto L661
	}
L659:
	;
	v2709 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2710 = m.ExcPending
	if v2710 != 0 {
		goto L1
	} else {
		goto L737
	}
L660:
	;
	v2343 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v2343)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = v2344
	*(*int32)(unsafe.Add(mBase, uint32(v22)+164)) = v2317
	v2348 = v22 + int32(336)
	v2353 = F_pg_snprintf(m, v2348, int32(1024), int32(_a_F_apply_dispatch_33), v22+int32(160))
	mBase = m.M
	v2354 = m.ExcPending
	if v2354 != 0 {
		goto L1
	} else {
		goto L663
	}
L661:
	;
	goto L662
L662:
	;
	v2376 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v2376 < int32(0) {
		goto L668
	} else {
		goto L669
	}
L663:
	;
	v2356 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2357 = *(*int32)(unsafe.Add(mBase, uint32(v2356)+60))
	F_BufFileDeleteFileSet(m, v2357, v2348, int32(0))
	mBase = m.M
	v2360 = m.ExcPending
	if v2360 != 0 {
		goto L1
	} else {
		goto L664
	}
L664:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+148)) = v2317
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v2344
	v2367 = F_pg_snprintf(m, v2348, int32(1024), int32(_a_F_apply_dispatch_34), v22+int32(144))
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L1
	} else {
		goto L665
	}
L665:
	;
	v2370 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+60))
	F_BufFileDeleteFileSet(m, v2371, v2348, int32(1))
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L1
	} else {
		goto L666
	}
L666:
	;
	goto L659
L667:
	;
	v2383 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(v2383)+20))
	goto L671
L668:
	;
	v2380 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v2380
	goto L670
L669:
	;
	goto L670
L670:
	;
	goto L667
L671:
	;
	if base.B2i32(v2384 == int32(2)) == int32(0) {
		goto L672
	} else {
		goto L673
	}
L672:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2390 = m.ExcPending
	if v2390 != 0 {
		goto L1
	} else {
		goto L675
	}
L673:
	;
	goto L674
L674:
	;
	v2393 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L1
	} else {
		goto L677
	}
L675:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L1
	} else {
		goto L676
	}
L676:
	;
	goto L674
L677:
	;
	F_PushActiveSnapshot(m, v2393)
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L1
	} else {
		goto L678
	}
L678:
	;
	v2399 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v2399
	v2402 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v2402)+32))
	F_subxact_info_read(m, v2403, v2317)
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L1
	} else {
		goto L679
	}
L679:
	;
	v2407 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[30]))
	v2409 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_apply_dispatch[31])))
	v2426 = v2409
	goto L682
L680:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2683 = m.ExcPending
	if v2683 != 0 {
		goto L1
	} else {
		goto L734
	}
L681:
	;
	if v2407 != 0 {
		goto L730
	} else {
		goto L731
	}
L682:
	;
	if v2426 <= int64(0) {
		goto L681
	} else {
		goto L684
	}
L683:
	;
	v2440 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(v2440)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v2441
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v2317
	v2445 = v22 + int32(336)
	v2450 = F_pg_snprintf(m, v2445, int32(1024), int32(_a_F_apply_dispatch_33), v22+int32(176))
	mBase = m.M
	v2451 = m.ExcPending
	if v2451 != 0 {
		goto L1
	} else {
		goto L686
	}
L684:
	;
	v2432 = v2426 - int64(1)
	v2433 = base.I32_wrap_i64(v2432)
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(v2407+v2433<<(uint(int32(4))%32))))
	if v2437 != v2312 {
		v2426 = v2432
		goto L682
	} else {
		goto L685
	}
L685:
	;
	goto L683
L686:
	;
	v2453 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v2453)+60))
	v2457 = F_BufFileOpenFileSet(m, v2454, v2445, int32(2), int32(0))
	mBase = m.M
	v2458 = m.ExcPending
	if v2458 != 0 {
		goto L1
	} else {
		goto L687
	}
L687:
	;
	v2460 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[30]))
	v2463 = v2460 + v2433<<(uint(int32(4))%32)
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(v2463)+4))
	v2465 = *(*int64)(unsafe.Add(mBase, uint32(v2463)+8))
	v2466 = m.G0
	v2468 = v2466 - int32(1072)
	m.G0 = v2468
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(v2457)))
	v2472 = v2470 - int32(1)
	if v2472 < v2464 {
		goto L692
	} else {
		goto L693
	}
L688:
	;
	F_BufFileClose(m, v2457)
	mBase = m.M
	v2647 = m.ExcPending
	if v2647 != 0 {
		goto L1
	} else {
		goto L728
	}
L689:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2621 = m.ExcPending
	if v2621 != 0 {
		goto L1
	} else {
		goto L723
	}
L690:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2604 = m.ExcPending
	if v2604 != 0 {
		goto L1
	} else {
		goto L719
	}
L691:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2457))) = v2558
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(v2457)+24))
	if v2551 != v2570 {
		goto L709
	} else {
		goto L710
	}
L692:
	;
	v2474 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+32))
	v2551 = v2464
	v2558 = v2470
	v2566 = v2474
	goto L691
L693:
	;
	goto L694
L694:
	;
	v2476 = v2464
	v2480 = v2472
	v2483 = v2470
	goto L695
L695:
	;
	v2494 = int32(0)
	if base.B2i32(v2480 == v2494)|base.B2i32(base.B2i32(v2465 == int64(0))|base.B2i32(v2464 != v2480) == v2494) == v2494 {
		goto L698
	} else {
		goto L699
	}
L696:
	;
	v2551 = v2543
	v2558 = v2544
	v2566 = v2546
	goto L691
L697:
	;
	v2548 = v2480 - int32(1)
	if v2464 <= v2548 {
		v2476 = v2543
		v2480 = v2548
		v2483 = v2544
		goto L695
	} else {
		goto L707
	}
L698:
	;
	v2505 = *(*int32)(unsafe.Add(mBase, uint32(v2457)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2468)+16)) = v2505
	*(*int32)(unsafe.Add(mBase, uint32(v2468)+20)) = v2480
	v2509 = v2468 + int32(48)
	v2514 = F_pg_snprintf(m, v2509, int32(1024), int32(_a_F_apply_dispatch_35), v2468+int32(16))
	mBase = m.M
	v2515 = m.ExcPending
	if v2515 != 0 {
		goto L1
	} else {
		goto L701
	}
L699:
	;
	goto L700
L700:
	;
	v2533 = *(*int32)(unsafe.Add(mBase, uint32(v2457)+4))
	v2537 = *(*int32)(unsafe.Add(mBase, uint32(v2533+v2480<<(uint(int32(2))%32))))
	v2539 = F_FileTruncate(m, v2537, v2465, int32(167772167))
	mBase = m.M
	v2540 = m.ExcPending
	if v2540 != 0 {
		goto L1
	} else {
		goto L705
	}
L701:
	;
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(v2457)+4))
	v2520 = *(*int32)(unsafe.Add(mBase, uint32(v2516+v2480<<(uint(int32(2))%32))))
	F_FileClose(m, v2520)
	mBase = m.M
	v2522 = m.ExcPending
	if v2522 != 0 {
		goto L1
	} else {
		goto L702
	}
L702:
	;
	v2523 = *(*int32)(unsafe.Add(mBase, uint32(v2457)+12))
	v2524 = F_FileSetDelete(m, v2523, v2509)
	mBase = m.M
	v2525 = m.ExcPending
	if v2525 != 0 {
		goto L1
	} else {
		goto L703
	}
L703:
	;
	if v2524 == int32(0) {
		goto L690
	} else {
		goto L704
	}
L704:
	;
	v2543 = v2476 - base.B2i32(v2464 == v2480)
	v2544 = v2483 - int32(1)
	v2546 = int64(1073741824)
	goto L697
L705:
	;
	if v2539 < int32(0) {
		goto L689
	} else {
		goto L706
	}
L706:
	;
	v2543 = v2476
	v2544 = v2483
	v2546 = v2465
	goto L697
L707:
	;
	goto L696
L708:
	;
	m.G0 = v2468 + int32(1072)
	goto L688
L709:
	;
	if v2570 <= v2551 {
		goto L708
	} else {
		goto L718
	}
L710:
	;
	v2572 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+32))
	if v2572 <= v2566 {
		goto L711
	} else {
		goto L712
	}
L711:
	;
	v2574 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+48))
	if v2574+v2572 < v2566 {
		goto L709
	} else {
		goto L714
	}
L712:
	;
	goto L713
L713:
	;
	v2583 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2457)+40)) = v2583
	*(*int64)(unsafe.Add(mBase, uint32(v2457)+32)) = v2566
	*(*int64)(unsafe.Add(mBase, uint32(v2457)+48)) = v2583
	goto L708
L714:
	;
	v2577 = v2566 - v2572
	v2578 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+40))
	if v2566 <= v2578+v2572 {
		goto L715
	} else {
		goto L716
	}
L715:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2457)+40)) = v2577
	goto L717
L716:
	;
	goto L717
L717:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2457)+48)) = v2577
	goto L708
L718:
	;
	v2590 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2457)+40)) = v2590
	*(*int64)(unsafe.Add(mBase, uint32(v2457)+32)) = v2566
	*(*int32)(unsafe.Add(mBase, uint32(v2457)+24)) = v2551
	*(*int64)(unsafe.Add(mBase, uint32(v2457)+48)) = v2590
	goto L708
L719:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2606 = m.ExcPending
	if v2606 != 0 {
		goto L1
	} else {
		goto L720
	}
L720:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2468))) = v2468 + int32(48)
	F_errmsg(m, int32(_a_F_apply_dispatch_36), v2468)
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L1
	} else {
		goto L721
	}
L721:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_37), int32(952), int32(_a_F_apply_dispatch_38))
	mBase = m.M
	v2617 = m.ExcPending
	if v2617 != 0 {
		goto L1
	} else {
		goto L722
	}
L722:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L723:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2623 = m.ExcPending
	if v2623 != 0 {
		goto L1
	} else {
		goto L724
	}
L724:
	;
	v2624 = *(*int32)(unsafe.Add(mBase, uint32(v2457)+4))
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v2624+v2480<<(uint(int32(2))%32))))
	v2630 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[32]))
	v2634 = *(*int32)(unsafe.Add(mBase, uint32(v2630+v2628*int32(48))+32))
	goto L725
L725:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2468)+32)) = v2634
	F_errmsg(m, int32(_a_F_apply_dispatch_39), v2468+int32(32))
	mBase = m.M
	v2640 = m.ExcPending
	if v2640 != 0 {
		goto L1
	} else {
		goto L726
	}
L726:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_37), int32(970), int32(_a_F_apply_dispatch_38))
	mBase = m.M
	v2645 = m.ExcPending
	if v2645 != 0 {
		goto L1
	} else {
		goto L727
	}
L727:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L728:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[31])) = v2433
	v2651 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2652 = *(*int32)(unsafe.Add(mBase, uint32(v2651)+32))
	F_subxact_info_write(m, v2652, v2317)
	mBase = m.M
	v2654 = m.ExcPending
	if v2654 != 0 {
		goto L1
	} else {
		goto L729
	}
L729:
	;
	goto L680
L730:
	;
	F_pfree(m, v2407)
	mBase = m.M
	v2656 = m.ExcPending
	if v2656 != 0 {
		goto L1
	} else {
		goto L733
	}
L731:
	;
	goto L732
L732:
	;
	v2658 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[33])) = v2658
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[31])) = v2658
	goto L680
L733:
	;
	goto L732
L734:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2685 = m.ExcPending
	if v2685 != 0 {
		goto L1
	} else {
		goto L735
	}
L735:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2687 = m.ExcPending
	if v2687 != 0 {
		goto L1
	} else {
		goto L736
	}
L736:
	;
	goto L659
L737:
	;
	if v2709 == int32(0) {
		goto L10
	} else {
		goto L738
	}
L738:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_40), int32(0))
	mBase = m.M
	v2716 = m.ExcPending
	if v2716 != 0 {
		goto L1
	} else {
		goto L739
	}
L739:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(2136), int32(_a_F_apply_dispatch_41))
	mBase = m.M
	v2721 = m.ExcPending
	if v2721 != 0 {
		goto L1
	} else {
		goto L740
	}
L740:
	;
	goto L10
L741:
	;
	v2733 = m.G0
	v2735 = v2733 - int32(96)
	m.G0 = v2735
	v2738 = v22 + int32(1360)
	v2739 = *(*int32)(unsafe.Add(mBase, uint32(v2738)+4))
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v2738)))
	v2742 = *(*int64)(unsafe.Add(mBase, uint32(v2738)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[34])) = v2742
	v2745 = *(*int64)(unsafe.Add(mBase, uint32(v2738)+16))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[35])) = v2745
	if v2739 == v2740 {
		goto L746
	} else {
		goto L747
	}
L742:
	;
	v2724 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25]))
	if v2724 == int32(0) {
		goto L741
	} else {
		goto L743
	}
L743:
	;
	F_BufFileClose(m, v2724)
	mBase = m.M
	v2728 = m.ExcPending
	if v2728 != 0 {
		goto L1
	} else {
		goto L744
	}
L744:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25])) = int32(0)
	goto L741
L745:
	;
	m.G0 = v2735 + int32(96)
	if v2312 != v2317 {
		goto L784
	} else {
		goto L785
	}
L746:
	;
	v2749 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[27]))
	v2752 = base.AtomicRmwXchg32(m, v2749, int32(0), int32(1))
	if v2752 != 0 {
		goto L749
	} else {
		goto L750
	}
L747:
	;
	goto L748
L748:
	;
	v2788 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(v2788)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2735)+16)) = v2789
	*(*int32)(unsafe.Add(mBase, uint32(v2735)+20)) = v2739
	v2793 = v2735 + int32(32)
	v2798 = F_pg_snprintf(m, v2793, int32(64), int32(_a_F_apply_dispatch_42), v2735+int32(16))
	mBase = m.M
	v2799 = m.ExcPending
	if v2799 != 0 {
		goto L1
	} else {
		goto L761
	}
L749:
	;
	F_s_lock(m, v2749, int32(_a_F_apply_dispatch_27))
	mBase = m.M
	v2755 = m.ExcPending
	if v2755 != 0 {
		goto L1
	} else {
		goto L752
	}
L750:
	;
	goto L751
L751:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2749)+8)) = int32(2)
	v2758 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2749))), uint32(v2758))
	v2762 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2763 = *(*int32)(unsafe.Add(mBase, uint32(v2762)+32))
	F_UnlockApplyTransactionForSession(m, v2763, v2740, int32(1), int32(8))
	mBase = m.M
	v2767 = m.ExcPending
	if v2767 != 0 {
		goto L1
	} else {
		goto L753
	}
L752:
	;
	goto L751
L753:
	;
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v2769 = m.ExcPending
	if v2769 != 0 {
		goto L1
	} else {
		goto L754
	}
L754:
	;
	v2771 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v2772 = *(*int32)(unsafe.Add(mBase, uint32(v2771)+24))
	goto L755
L755:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2772) {
		goto L756
	} else {
		goto L757
	}
L756:
	;
	v2776 = F_EndTransactionBlock(m, int32(0))
	mBase = m.M
	v2777 = m.ExcPending
	if v2777 != 0 {
		goto L1
	} else {
		goto L759
	}
L757:
	;
	goto L758
L758:
	;
	v2781 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[36])) = v2781
	F_pgstat_report_activity(m, int32(2), v2781)
	mBase = m.M
	goto L745
L759:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2779 = m.ExcPending
	if v2779 != 0 {
		goto L1
	} else {
		goto L760
	}
L760:
	;
	goto L758
L761:
	;
	v2802 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2803 = m.ExcPending
	if v2803 != 0 {
		goto L1
	} else {
		goto L762
	}
L762:
	;
	if v2802 != 0 {
		goto L763
	} else {
		goto L764
	}
L763:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2735))) = v2793
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_43), v2735)
	mBase = m.M
	v2807 = m.ExcPending
	if v2807 != 0 {
		goto L1
	} else {
		goto L766
	}
L764:
	;
	goto L765
L765:
	;
	v2814 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[36]))
	if v2814 != 0 {
		goto L768
	} else {
		goto L769
	}
L766:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_29), int32(1487), int32(_a_F_apply_dispatch_44))
	mBase = m.M
	v2812 = m.ExcPending
	if v2812 != 0 {
		goto L1
	} else {
		goto L767
	}
L767:
	;
	goto L765
L768:
	;
	v2815 = *(*int32)(unsafe.Add(mBase, uint32(v2814)+4))
	v2816 = v2815
	goto L770
L769:
	;
	v2816 = int32(0)
	goto L770
L770:
	;
	v2817 = v2816
	goto L771
L771:
	;
	v2837 = v2817 - int32(1)
	if v2837 < int32(0) {
		goto L745
	} else {
		goto L773
	}
L772:
	;
	F_RollbackToSavepoint(m, v2735+int32(32))
	mBase = m.M
	v2849 = m.ExcPending
	if v2849 != 0 {
		goto L1
	} else {
		goto L775
	}
L773:
	;
	v2840 = *(*int32)(unsafe.Add(mBase, uint32(v2814)+12))
	v2844 = *(*int32)(unsafe.Add(mBase, uint32(v2840+v2837<<(uint(int32(2))%32))))
	if v2844 != v2739 {
		v2817 = v2837
		goto L771
	} else {
		goto L774
	}
L774:
	;
	goto L772
L775:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2851 = m.ExcPending
	if v2851 != 0 {
		goto L1
	} else {
		goto L776
	}
L776:
	;
	v2852 = int32(_a_F_apply_dispatch_45)
	v2854 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[36]))
	v2855 = int32(0)
	if base.B2i32(v2854 == v2855)|base.B2i32(v2837 <= v2855) != 0 {
		goto L778
	} else {
		goto L779
	}
L777:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[36])) = v2865
	goto L745
L778:
	;
	v2865 = int32(0)
	goto L780
L779:
	;
	v2862 = *(*int32)(unsafe.Add(mBase, uint32(v2854)+4))
	if v2837 < v2862 {
		goto L781
	} else {
		goto L782
	}
L780:
	;
	goto L777
L781:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2854)+4)) = v2837
	goto L783
L782:
	;
	goto L783
L783:
	;
	v2865 = v2854
	goto L780
L784:
	;
	F_pa_decr_and_wait_stream_block(m)
	mBase = m.M
	v2891 = m.ExcPending
	if v2891 != 0 {
		goto L1
	} else {
		goto L787
	}
L785:
	;
	goto L786
L786:
	;
	v2894 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2895 = m.ExcPending
	if v2895 != 0 {
		goto L1
	} else {
		goto L788
	}
L787:
	;
	goto L786
L788:
	;
	if v2894 == int32(0) {
		goto L10
	} else {
		goto L789
	}
L789:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_40), int32(0))
	mBase = m.M
	v2901 = m.ExcPending
	if v2901 != 0 {
		goto L1
	} else {
		goto L790
	}
L790:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(2239), int32(_a_F_apply_dispatch_41))
	mBase = m.M
	v2906 = m.ExcPending
	if v2906 != 0 {
		goto L1
	} else {
		goto L791
	}
L791:
	;
	goto L10
L792:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = int32(1)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_46), v22+int32(128))
	mBase = m.M
	v2917 = m.ExcPending
	if v2917 != 0 {
		goto L1
	} else {
		goto L793
	}
L793:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(2243), int32(_a_F_apply_dispatch_41))
	mBase = m.M
	v2922 = m.ExcPending
	if v2922 != 0 {
		goto L1
	} else {
		goto L794
	}
L794:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L795:
	;
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v2326)+16))
	v2928 = base.AtomicRmwAdd32(m, v2925, int32(20), int32(1))
	F_pa_lock_stream(m, v2317)
	mBase = m.M
	v2930 = m.ExcPending
	if v2930 != 0 {
		goto L1
	} else {
		goto L796
	}
L796:
	;
	v2931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2932 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2933 = F_pa_send_data(m, v2326, v2931, v2932)
	mBase = m.M
	v2934 = m.ExcPending
	if v2934 != 0 {
		goto L1
	} else {
		goto L797
	}
L797:
	;
	if v2933 == int32(0) {
		goto L12
	} else {
		goto L798
	}
L798:
	;
	goto L10
L799:
	;
	v2944 = v22 + int32(336)
	v2945 = m.G0
	v2947 = v2945 - int32(16)
	m.G0 = v2947
	v2950 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v2951 = m.ExcPending
	if v2951 != 0 {
		goto L1
	} else {
		goto L800
	}
L800:
	;
	v2952 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v2953 = m.ExcPending
	if v2953 != 0 {
		goto L1
	} else {
		goto L801
	}
L801:
	;
	v2955 = v2952 & int32(255)
	if v2955 != 0 {
		goto L802
	} else {
		goto L803
	}
L802:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2959 = m.ExcPending
	if v2959 != 0 {
		goto L1
	} else {
		goto L805
	}
L803:
	;
	goto L804
L804:
	;
	v2969 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2970 = m.ExcPending
	if v2970 != 0 {
		goto L1
	} else {
		goto L808
	}
L805:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2947))) = v2955
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_9), v2947)
	mBase = m.M
	v2963 = m.ExcPending
	if v2963 != 0 {
		goto L1
	} else {
		goto L806
	}
L806:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(1143), int32(_a_F_apply_dispatch_47))
	mBase = m.M
	v2968 = m.ExcPending
	if v2968 != 0 {
		goto L1
	} else {
		goto L807
	}
L807:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L808:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2944))) = v2969
	v2972 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2973 = m.ExcPending
	if v2973 != 0 {
		goto L1
	} else {
		goto L809
	}
L809:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2944)+8)) = v2972
	v2975 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2976 = m.ExcPending
	if v2976 != 0 {
		goto L1
	} else {
		goto L810
	}
L810:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2944)+16)) = v2975
	m.G0 = v2947 + int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v2950
	v2984 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v2984
	v2987 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2987)+16)))
	if v2988 == int32(1) {
		goto L813
	} else {
		goto L814
	}
L811:
	;
	v3093 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	F_ProcessSyncingRelations(m, v3093)
	mBase = m.M
	v3095 = m.ExcPending
	if v3095 != 0 {
		goto L1
	} else {
		goto L851
	}
L812:
	;
	v3053 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25]))
	if v3053 != 0 {
		goto L839
	} else {
		goto L840
	}
L813:
	;
	v2991 = *(*int32)(unsafe.Add(mBase, uint32(v2987)))
	if v2991 == int32(4) {
		goto L812
	} else {
		goto L816
	}
L814:
	;
	goto L815
L815:
	;
	v2994 = F_pa_find_worker(m, v2950)
	mBase = m.M
	v2995 = m.ExcPending
	if v2995 != 0 {
		goto L1
	} else {
		goto L819
	}
L816:
	;
	goto L815
L817:
	;
	F_stream_open_and_write_change(m, v2950, int32(99), v22+int32(1360))
	mBase = m.M
	v3045 = m.ExcPending
	if v3045 != 0 {
		goto L1
	} else {
		goto L836
	}
L818:
	;
	F_pa_switch_to_partial_serialize(m, v2994, int32(1))
	mBase = m.M
	v3040 = m.ExcPending
	if v3040 != 0 {
		goto L1
	} else {
		goto L835
	}
L819:
	;
	if v2994 != 0 {
		goto L820
	} else {
		goto L821
	}
L820:
	;
	v2996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2994)+12)))
	if v2996 != 0 {
		goto L817
	} else {
		goto L823
	}
L821:
	;
	goto L822
L822:
	;
	v3007 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v3007 != 0 {
		goto L19
	} else {
		goto L827
	}
L823:
	;
	v2997 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2998 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2999 = F_pa_send_data(m, v2994, v2997, v2998)
	mBase = m.M
	v3000 = m.ExcPending
	if v3000 != 0 {
		goto L1
	} else {
		goto L824
	}
L824:
	;
	if v2999 == int32(0) {
		goto L818
	} else {
		goto L825
	}
L825:
	;
	v3003 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	F_pa_xact_finish(m, v2994, v3003)
	mBase = m.M
	v3005 = m.ExcPending
	if v3005 != 0 {
		goto L1
	} else {
		goto L826
	}
L826:
	;
	goto L811
L827:
	;
	v3009 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v3010 = *(*int32)(unsafe.Add(mBase, uint32(v3009)+60))
	v3011 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	F_apply_spooled_messages(m, v3010, v2950, v3011)
	mBase = m.M
	v3013 = m.ExcPending
	if v3013 != 0 {
		goto L1
	} else {
		goto L828
	}
L828:
	;
	F_apply_handle_commit_internal(m, v22+int32(336))
	mBase = m.M
	v3017 = m.ExcPending
	if v3017 != 0 {
		goto L1
	} else {
		goto L829
	}
L829:
	;
	v3019 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v3020 = *(*int32)(unsafe.Add(mBase, uint32(v3019)+32))
	F_stream_cleanup_files(m, v3020, v2950)
	mBase = m.M
	v3022 = m.ExcPending
	if v3022 != 0 {
		goto L1
	} else {
		goto L830
	}
L830:
	;
	v3025 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3026 = m.ExcPending
	if v3026 != 0 {
		goto L1
	} else {
		goto L831
	}
L831:
	;
	if v3025 == int32(0) {
		goto L811
	} else {
		goto L832
	}
L832:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_48), int32(0))
	mBase = m.M
	v3032 = m.ExcPending
	if v3032 != 0 {
		goto L1
	} else {
		goto L833
	}
L833:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(2452), int32(_a_F_apply_dispatch_49))
	mBase = m.M
	v3037 = m.ExcPending
	if v3037 != 0 {
		goto L1
	} else {
		goto L834
	}
L834:
	;
	goto L811
L835:
	;
	goto L817
L836:
	;
	v3046 = *(*int32)(unsafe.Add(mBase, uint32(v2994)+16))
	F_pa_set_fileset_state(m, v3046)
	mBase = m.M
	v3048 = m.ExcPending
	if v3048 != 0 {
		goto L1
	} else {
		goto L837
	}
L837:
	;
	v3049 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	F_pa_xact_finish(m, v2994, v3049)
	mBase = m.M
	v3051 = m.ExcPending
	if v3051 != 0 {
		goto L1
	} else {
		goto L838
	}
L838:
	;
	goto L811
L839:
	;
	F_BufFileClose(m, v3053)
	mBase = m.M
	v3055 = m.ExcPending
	if v3055 != 0 {
		goto L1
	} else {
		goto L842
	}
L840:
	;
	goto L841
L841:
	;
	F_apply_handle_commit_internal(m, v22+int32(336))
	mBase = m.M
	v3062 = m.ExcPending
	if v3062 != 0 {
		goto L1
	} else {
		goto L843
	}
L842:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25])) = int32(0)
	goto L841
L843:
	;
	v3064 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[27]))
	v3066 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[37]))
	*(*int64)(unsafe.Add(mBase, uint32(v3064)+24)) = v3066
	F_pa_set_xact_state(m, v3064, int32(2))
	mBase = m.M
	v3070 = m.ExcPending
	if v3070 != 0 {
		goto L1
	} else {
		goto L844
	}
L844:
	;
	F_pa_unlock_transaction(m, v2950)
	mBase = m.M
	v3072 = m.ExcPending
	if v3072 != 0 {
		goto L1
	} else {
		goto L845
	}
L845:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[36])) = int32(0)
	goto L846
L846:
	;
	v3078 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3079 = m.ExcPending
	if v3079 != 0 {
		goto L1
	} else {
		goto L847
	}
L847:
	;
	if v3078 == int32(0) {
		goto L811
	} else {
		goto L848
	}
L848:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_48), int32(0))
	mBase = m.M
	v3085 = m.ExcPending
	if v3085 != 0 {
		goto L1
	} else {
		goto L849
	}
L849:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(2506), int32(_a_F_apply_dispatch_49))
	mBase = m.M
	v3090 = m.ExcPending
	if v3090 != 0 {
		goto L1
	} else {
		goto L850
	}
L850:
	;
	goto L811
L851:
	;
	v3097 = int32(0)
	F_pgstat_report_activity(m, int32(2), v3097)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v3097
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v3097
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v3097
	goto L3
L852:
	;
	v3119 = *(*int32)(unsafe.Add(mBase, uint32(v3115)))
	if v3119 == int32(1) {
		goto L18
	} else {
		goto L855
	}
L853:
	;
	goto L854
L854:
	;
	v3123 = v22 + int32(336)
	v3124 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3125 = m.ExcPending
	if v3125 != 0 {
		goto L1
	} else {
		goto L856
	}
L855:
	;
	goto L854
L856:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3123))) = v3124
	if v3124 != int64(0) {
		goto L859
	} else {
		goto L860
	}
L857:
	;
	v3291 = *(*int32)(unsafe.Add(mBase, uint32(v22)+360))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v3291
	v3294 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v3294
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[9])) = v3294
	v3299 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v3300 = *(*int64)(unsafe.Add(mBase, uint32(v3299)+16))
	if base.B2i32(v3300 == int64(0))|base.B2i32(v3294 != v3300) != 0 {
		goto L904
	} else {
		goto L905
	}
L858:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3280 = m.ExcPending
	if v3280 != 0 {
		goto L1
	} else {
		goto L901
	}
L859:
	;
	v3129 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3130 = m.ExcPending
	if v3130 != 0 {
		goto L1
	} else {
		goto L862
	}
L860:
	;
	goto L861
L861:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3267 = m.ExcPending
	if v3267 != 0 {
		goto L1
	} else {
		goto L898
	}
L862:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3123)+8)) = v3129
	if v3129 == int64(0) {
		goto L858
	} else {
		goto L863
	}
L863:
	;
	v3134 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3135 = m.ExcPending
	if v3135 != 0 {
		goto L1
	} else {
		goto L864
	}
L864:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3123)+16)) = v3134
	v3138 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v3139 = m.ExcPending
	if v3139 != 0 {
		goto L1
	} else {
		goto L865
	}
L865:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3123)+24)) = v3138
	v3142 = v22 + int32(364)
	v3143 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v3144 = m.ExcPending
	if v3144 != 0 {
		goto L1
	} else {
		goto L866
	}
L866:
	;
	goto L870
L867:
	;
	goto L857
L868:
	;
	v3261 = F_strlen(m, v3250)
	mBase = m.M
	goto L867
L870:
	;
	goto L871
L871:
	;
	v3151 = int32(199)
	if (v3142^v3143)&int32(3) != 0 {
		goto L875
	} else {
		goto L876
	}
L872:
	;
	v3254 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3251))) = uint8(v3254)
	goto L868
L873:
	;
	v3235 = v3230
	v3236 = v3231
	v3237 = v3232
	goto L894
L874:
	;
	if v3225 == int32(0) {
		v3250 = v3223
		v3251 = v3224
		goto L872
	} else {
		goto L893
	}
L875:
	;
	v3223 = v3143
	v3224 = v3142
	v3225 = v3151
	goto L874
L876:
	;
	goto L877
L877:
	;
	v3155 = int32(0)
	if base.B2i32(v3143&int32(3) == v3155)|int32(0) == v3155 {
		goto L879
	} else {
		goto L880
	}
L878:
	;
	if v3191 == int32(0) {
		v3250 = v3188
		v3251 = v3189
		goto L872
	} else {
		goto L887
	}
L879:
	;
	v3167 = v3143
	v3168 = v3142
	v3169 = v3151
	goto L882
L880:
	;
	goto L881
L881:
	;
	v3188 = v3143
	v3189 = v3142
	v3190 = v3151
	v3191 = int32(1)
	goto L878
L882:
	;
	v3171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3167))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3168))) = uint8(v3171)
	if v3171 == int32(0) {
		v3230 = v3167
		v3231 = v3168
		v3232 = v3169
		goto L873
	} else {
		goto L884
	}
L883:
	;
	v3188 = v3182
	v3189 = v3176
	v3190 = v3178
	v3191 = v3180
	goto L878
L884:
	;
	v3175 = int32(1)
	v3176 = v3168 + v3175
	v3178 = v3169 - v3175
	v3179 = int32(0)
	v3180 = base.B2i32(v3178 != v3179)
	v3182 = v3167 + v3175
	if v3182&int32(3) == v3179 {
		v3188 = v3182
		v3189 = v3176
		v3190 = v3178
		v3191 = v3180
		goto L878
	} else {
		goto L885
	}
L885:
	;
	if v3178 != 0 {
		v3167 = v3182
		v3168 = v3176
		v3169 = v3178
		goto L882
	} else {
		goto L886
	}
L886:
	;
	goto L883
L887:
	;
	v3194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3188))))
	if base.B2i32(v3194 == int32(0))|base.B2i32(base.Ui32(v3190) < base.Ui32(int32(4))) != 0 {
		v3223 = v3188
		v3224 = v3189
		v3225 = v3190
		goto L874
	} else {
		goto L888
	}
L888:
	;
	v3201 = v3188
	v3202 = v3189
	v3203 = v3190
	goto L889
L889:
	;
	v3206 = *(*int32)(unsafe.Add(mBase, uint32(v3201)))
	v3209 = int32(-2139062144)
	if (int32(16843008)-v3206|v3206)&v3209 != v3209 {
		v3230 = v3201
		v3231 = v3202
		v3232 = v3203
		goto L873
	} else {
		goto L891
	}
L890:
	;
	v3223 = v3217
	v3224 = v3215
	v3225 = v3219
	goto L874
L891:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3202))) = v3206
	v3214 = int32(4)
	v3215 = v3202 + v3214
	v3217 = v3201 + v3214
	v3219 = v3203 - v3214
	if base.Ui32(int32(3)) < base.Ui32(v3219) {
		v3201 = v3217
		v3202 = v3215
		v3203 = v3219
		goto L889
	} else {
		goto L892
	}
L892:
	;
	goto L890
L893:
	;
	v3230 = v3223
	v3231 = v3224
	v3232 = v3225
	goto L873
L894:
	;
	v3239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3235))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3236))) = uint8(v3239)
	if v3239 == int32(0) {
		v3250 = v3235
		v3251 = v3236
		goto L872
	} else {
		goto L896
	}
L895:
	;
	v3250 = v3246
	v3251 = v3244
	goto L872
L896:
	;
	v3243 = int32(1)
	v3244 = v3236 + v3243
	v3246 = v3235 + v3243
	v3248 = v3237 - v3243
	if v3248 != 0 {
		v3235 = v3246
		v3236 = v3244
		v3237 = v3248
		goto L894
	} else {
		goto L897
	}
L897:
	;
	goto L895
L898:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_50), int32(0))
	mBase = m.M
	v3271 = m.ExcPending
	if v3271 != 0 {
		goto L1
	} else {
		goto L899
	}
L899:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(139), int32(_a_F_apply_dispatch_51))
	mBase = m.M
	v3276 = m.ExcPending
	if v3276 != 0 {
		goto L1
	} else {
		goto L900
	}
L900:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L901:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_52), int32(0))
	mBase = m.M
	v3284 = m.ExcPending
	if v3284 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(142), int32(_a_F_apply_dispatch_51))
	mBase = m.M
	v3289 = m.ExcPending
	if v3289 != 0 {
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
	v3331 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])) = uint8(v3331)
	F_pgstat_report_activity(m, int32(3), int32(0))
	mBase = m.M
	goto L3
L905:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8])) = v3294
	v3309 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3310 = m.ExcPending
	if v3310 != 0 {
		goto L1
	} else {
		goto L906
	}
L906:
	;
	if v3309 == int32(0) {
		goto L904
	} else {
		goto L907
	}
L907:
	;
	v3314 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+212)) = uint32(v3314)
	v3317 = int64(base.Ui64(v3314) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+208)) = uint32(v3317)
	F_errmsg(m, int32(_a_F_apply_dispatch_5), v22+int32(208))
	mBase = m.M
	v3323 = m.ExcPending
	if v3323 != 0 {
		goto L1
	} else {
		goto L908
	}
L908:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(_a_F_apply_dispatch_7), int32(_a_F_apply_dispatch_8))
	mBase = m.M
	v3328 = m.ExcPending
	if v3328 != 0 {
		goto L1
	} else {
		goto L909
	}
L909:
	;
	goto L904
L910:
	;
	v3341 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	v3343 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[9]))
	if v3341 != v3343 {
		goto L17
	} else {
		goto L911
	}
L911:
	;
	v3346 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v3346 < int32(0) {
		goto L913
	} else {
		goto L914
	}
L912:
	;
	v3353 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v3354 = *(*int32)(unsafe.Add(mBase, uint32(v3353)+20))
	goto L916
L913:
	;
	v3350 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v3350
	goto L915
L914:
	;
	goto L915
L915:
	;
	goto L912
L916:
	;
	if base.B2i32(v3354 == int32(2)) == int32(0) {
		goto L917
	} else {
		goto L918
	}
L917:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v3360 = m.ExcPending
	if v3360 != 0 {
		goto L1
	} else {
		goto L920
	}
L918:
	;
	goto L919
L919:
	;
	v3363 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v3364 = m.ExcPending
	if v3364 != 0 {
		goto L1
	} else {
		goto L922
	}
L920:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v3362 = m.ExcPending
	if v3362 != 0 {
		goto L1
	} else {
		goto L921
	}
L921:
	;
	goto L919
L922:
	;
	F_PushActiveSnapshot(m, v3363)
	mBase = m.M
	v3366 = m.ExcPending
	if v3366 != 0 {
		goto L1
	} else {
		goto L923
	}
L923:
	;
	v3369 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v3369
	v3372 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v3373 = *(*int32)(unsafe.Add(mBase, uint32(v3372)+4))
	v3374 = *(*int32)(unsafe.Add(mBase, uint32(v22)+360))
	F_TwoPhaseTransactionGid(m, v3373, v3374, v22+int32(1360))
	mBase = m.M
	v3378 = m.ExcPending
	if v3378 != 0 {
		goto L1
	} else {
		goto L924
	}
L924:
	;
	v3380 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v3381 = *(*int32)(unsafe.Add(mBase, uint32(v3380)+24))
	goto L925
L925:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(v3381)) == int32(0) {
		goto L926
	} else {
		goto L927
	}
L926:
	;
	F_BeginTransactionBlock(m)
	mBase = m.M
	v3387 = m.ExcPending
	if v3387 != 0 {
		goto L1
	} else {
		goto L929
	}
L927:
	;
	goto L928
L928:
	;
	v3391 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[34])) = v3391
	v3394 = *(*int64)(unsafe.Add(mBase, uint32(v22)+352))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[35])) = v3394
	v3398 = F_PrepareTransactionBlock(m, v22+int32(1360))
	mBase = m.M
	v3399 = m.ExcPending
	if v3399 != 0 {
		goto L1
	} else {
		goto L931
	}
L929:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3389 = m.ExcPending
	if v3389 != 0 {
		goto L1
	} else {
		goto L930
	}
L930:
	;
	goto L928
L931:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v3401 = m.ExcPending
	if v3401 != 0 {
		goto L1
	} else {
		goto L932
	}
L932:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v3403 = m.ExcPending
	if v3403 != 0 {
		goto L1
	} else {
		goto L933
	}
L933:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3405 = m.ExcPending
	if v3405 != 0 {
		goto L1
	} else {
		goto L934
	}
L934:
	;
	v3407 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v3408 = m.ExcPending
	if v3408 != 0 {
		goto L1
	} else {
		goto L935
	}
L935:
	;
	v3409 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	v3411 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v3412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3411)+16)))
	if v3412 == int32(1) {
		goto L937
	} else {
		goto L938
	}
L936:
	;
	v3452 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])) = uint8(v3452)
	F_ProcessSyncingRelations(m, v3450)
	mBase = m.M
	v3455 = m.ExcPending
	if v3455 != 0 {
		goto L1
	} else {
		goto L946
	}
L937:
	;
	v3415 = *(*int32)(unsafe.Add(mBase, uint32(v3411)))
	if v3415 == int32(4) {
		v3450 = v3409
		goto L936
	} else {
		goto L940
	}
L938:
	;
	goto L939
L939:
	;
	v3420 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v3420
	v3423 = F_palloc(m, int32(24))
	mBase = m.M
	v3424 = m.ExcPending
	if v3424 != 0 {
		goto L1
	} else {
		goto L941
	}
L940:
	;
	goto L939
L941:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3423)+16)) = v3409
	*(*int64)(unsafe.Add(mBase, uint32(v3423)+8)) = int64(0)
	v3429 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[38]))
	if v3429 != 0 {
		goto L943
	} else {
		goto L944
	}
L942:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3423))) = v3436
	v3438 = int32(_a_F_apply_dispatch_53)
	*(*int32)(unsafe.Add(mBase, uint32(v3423)+4)) = v3438
	*(*int32)(unsafe.Add(mBase, uint32(v3436)+4)) = v3423
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[39])) = v3423
	v3445 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v3445
	v3447 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	v3450 = v3447
	goto L936
L943:
	;
	v3431 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[39]))
	v3436 = v3431
	goto L942
L944:
	;
	goto L945
L945:
	;
	v3433 = int32(_a_F_apply_dispatch_53)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[38])) = v3433
	v3436 = v3433
	goto L942
L946:
	;
	v3457 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	if v3457 != int64(0) {
		goto L947
	} else {
		goto L948
	}
L947:
	;
	v3462 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3463 = m.ExcPending
	if v3463 != 0 {
		goto L1
	} else {
		goto L950
	}
L948:
	;
	goto L949
L949:
	;
	v3485 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	F_clear_subscription_skip_lsn(m, v3485)
	mBase = m.M
	v3487 = m.ExcPending
	if v3487 != 0 {
		goto L1
	} else {
		goto L956
	}
L950:
	;
	if v3462 != 0 {
		goto L951
	} else {
		goto L952
	}
L951:
	;
	v3465 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+228)) = uint32(v3465)
	v3468 = int64(base.Ui64(v3465) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+224)) = uint32(v3468)
	F_errmsg(m, int32(_a_F_apply_dispatch_54), v22+int32(224))
	mBase = m.M
	v3474 = m.ExcPending
	if v3474 != 0 {
		goto L1
	} else {
		goto L954
	}
L952:
	;
	goto L953
L953:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8])) = int64(0)
	goto L949
L954:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(_a_F_apply_dispatch_55), int32(_a_F_apply_dispatch_56))
	mBase = m.M
	v3479 = m.ExcPending
	if v3479 != 0 {
		goto L1
	} else {
		goto L955
	}
L955:
	;
	goto L953
L956:
	;
	v3489 = int32(0)
	F_pgstat_report_activity(m, int32(2), v3489)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v3489
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v3489
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v3489
	goto L3
L957:
	;
	v3701 = *(*int32)(unsafe.Add(mBase, uint32(v22)+360))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v3701
	v3704 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v3704
	v3707 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v3708 = *(*int32)(unsafe.Add(mBase, uint32(v3707)+4))
	F_TwoPhaseTransactionGid(m, v3708, v3701, v22+int32(1360))
	mBase = m.M
	v3712 = m.ExcPending
	if v3712 != 0 {
		goto L1
	} else {
		goto L1011
	}
L958:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3690 = m.ExcPending
	if v3690 != 0 {
		goto L1
	} else {
		goto L1008
	}
L959:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3677 = m.ExcPending
	if v3677 != 0 {
		goto L1
	} else {
		goto L1005
	}
L960:
	;
	v3515 = v3512 & int32(255)
	if v3515 == int32(0) {
		goto L961
	} else {
		goto L962
	}
L961:
	;
	v3518 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3519 = m.ExcPending
	if v3519 != 0 {
		goto L1
	} else {
		goto L964
	}
L962:
	;
	goto L963
L963:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3664 = m.ExcPending
	if v3664 != 0 {
		goto L1
	} else {
		goto L1002
	}
L964:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3507))) = v3518
	if v3518 == int64(0) {
		goto L959
	} else {
		goto L965
	}
L965:
	;
	v3523 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3524 = m.ExcPending
	if v3524 != 0 {
		goto L1
	} else {
		goto L966
	}
L966:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3507)+8)) = v3523
	if v3523 == int64(0) {
		goto L958
	} else {
		goto L967
	}
L967:
	;
	v3528 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3529 = m.ExcPending
	if v3529 != 0 {
		goto L1
	} else {
		goto L968
	}
L968:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3507)+16)) = v3528
	v3532 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v3533 = m.ExcPending
	if v3533 != 0 {
		goto L1
	} else {
		goto L969
	}
L969:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3507)+24)) = v3532
	v3536 = v22 + int32(364)
	v3537 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v3538 = m.ExcPending
	if v3538 != 0 {
		goto L1
	} else {
		goto L970
	}
L970:
	;
	goto L974
L971:
	;
	m.G0 = v3510 + int32(16)
	goto L957
L972:
	;
	v3655 = F_strlen(m, v3644)
	mBase = m.M
	goto L971
L974:
	;
	goto L975
L975:
	;
	v3545 = int32(199)
	if (v3536^v3537)&int32(3) != 0 {
		goto L979
	} else {
		goto L980
	}
L976:
	;
	v3648 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3645))) = uint8(v3648)
	goto L972
L977:
	;
	v3629 = v3624
	v3630 = v3625
	v3631 = v3626
	goto L998
L978:
	;
	if v3619 == int32(0) {
		v3644 = v3617
		v3645 = v3618
		goto L976
	} else {
		goto L997
	}
L979:
	;
	v3617 = v3537
	v3618 = v3536
	v3619 = v3545
	goto L978
L980:
	;
	goto L981
L981:
	;
	v3549 = int32(0)
	if base.B2i32(v3537&int32(3) == v3549)|int32(0) == v3549 {
		goto L983
	} else {
		goto L984
	}
L982:
	;
	if v3585 == int32(0) {
		v3644 = v3582
		v3645 = v3583
		goto L976
	} else {
		goto L991
	}
L983:
	;
	v3561 = v3537
	v3562 = v3536
	v3563 = v3545
	goto L986
L984:
	;
	goto L985
L985:
	;
	v3582 = v3537
	v3583 = v3536
	v3584 = v3545
	v3585 = int32(1)
	goto L982
L986:
	;
	v3565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3561))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3562))) = uint8(v3565)
	if v3565 == int32(0) {
		v3624 = v3561
		v3625 = v3562
		v3626 = v3563
		goto L977
	} else {
		goto L988
	}
L987:
	;
	v3582 = v3576
	v3583 = v3570
	v3584 = v3572
	v3585 = v3574
	goto L982
L988:
	;
	v3569 = int32(1)
	v3570 = v3562 + v3569
	v3572 = v3563 - v3569
	v3573 = int32(0)
	v3574 = base.B2i32(v3572 != v3573)
	v3576 = v3561 + v3569
	if v3576&int32(3) == v3573 {
		v3582 = v3576
		v3583 = v3570
		v3584 = v3572
		v3585 = v3574
		goto L982
	} else {
		goto L989
	}
L989:
	;
	if v3572 != 0 {
		v3561 = v3576
		v3562 = v3570
		v3563 = v3572
		goto L986
	} else {
		goto L990
	}
L990:
	;
	goto L987
L991:
	;
	v3588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3582))))
	if base.B2i32(v3588 == int32(0))|base.B2i32(base.Ui32(v3584) < base.Ui32(int32(4))) != 0 {
		v3617 = v3582
		v3618 = v3583
		v3619 = v3584
		goto L978
	} else {
		goto L992
	}
L992:
	;
	v3595 = v3582
	v3596 = v3583
	v3597 = v3584
	goto L993
L993:
	;
	v3600 = *(*int32)(unsafe.Add(mBase, uint32(v3595)))
	v3603 = int32(-2139062144)
	if (int32(16843008)-v3600|v3600)&v3603 != v3603 {
		v3624 = v3595
		v3625 = v3596
		v3626 = v3597
		goto L977
	} else {
		goto L995
	}
L994:
	;
	v3617 = v3611
	v3618 = v3609
	v3619 = v3613
	goto L978
L995:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3596))) = v3600
	v3608 = int32(4)
	v3609 = v3596 + v3608
	v3611 = v3595 + v3608
	v3613 = v3597 - v3608
	if base.Ui32(int32(3)) < base.Ui32(v3613) {
		v3595 = v3611
		v3596 = v3609
		v3597 = v3613
		goto L993
	} else {
		goto L996
	}
L996:
	;
	goto L994
L997:
	;
	v3624 = v3617
	v3625 = v3618
	v3626 = v3619
	goto L977
L998:
	;
	v3633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3629))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3630))) = uint8(v3633)
	if v3633 == int32(0) {
		v3644 = v3629
		v3645 = v3630
		goto L976
	} else {
		goto L1000
	}
L999:
	;
	v3644 = v3640
	v3645 = v3638
	goto L976
L1000:
	;
	v3637 = int32(1)
	v3638 = v3630 + v3637
	v3640 = v3629 + v3637
	v3642 = v3631 - v3637
	if v3642 != 0 {
		v3629 = v3640
		v3630 = v3638
		v3631 = v3642
		goto L998
	} else {
		goto L1001
	}
L1001:
	;
	goto L999
L1002:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3510))) = v3515
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_57), v3510)
	mBase = m.M
	v3668 = m.ExcPending
	if v3668 != 0 {
		goto L1
	} else {
		goto L1003
	}
L1003:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(273), int32(_a_F_apply_dispatch_58))
	mBase = m.M
	v3673 = m.ExcPending
	if v3673 != 0 {
		goto L1
	} else {
		goto L1004
	}
L1004:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1005:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_59), int32(0))
	mBase = m.M
	v3681 = m.ExcPending
	if v3681 != 0 {
		goto L1
	} else {
		goto L1006
	}
L1006:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(278), int32(_a_F_apply_dispatch_58))
	mBase = m.M
	v3686 = m.ExcPending
	if v3686 != 0 {
		goto L1
	} else {
		goto L1007
	}
L1007:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1008:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_60), int32(0))
	mBase = m.M
	v3694 = m.ExcPending
	if v3694 != 0 {
		goto L1
	} else {
		goto L1009
	}
L1009:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(281), int32(_a_F_apply_dispatch_58))
	mBase = m.M
	v3699 = m.ExcPending
	if v3699 != 0 {
		goto L1
	} else {
		goto L1010
	}
L1010:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1011:
	;
	v3714 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v3714 < int32(0) {
		goto L1013
	} else {
		goto L1014
	}
L1012:
	;
	v3721 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v3722 = *(*int32)(unsafe.Add(mBase, uint32(v3721)+20))
	goto L1016
L1013:
	;
	v3718 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v3718
	goto L1015
L1014:
	;
	goto L1015
L1015:
	;
	goto L1012
L1016:
	;
	if base.B2i32(v3722 == int32(2)) == int32(0) {
		goto L1017
	} else {
		goto L1018
	}
L1017:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v3728 = m.ExcPending
	if v3728 != 0 {
		goto L1
	} else {
		goto L1020
	}
L1018:
	;
	goto L1019
L1019:
	;
	v3731 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v3732 = m.ExcPending
	if v3732 != 0 {
		goto L1
	} else {
		goto L1022
	}
L1020:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v3730 = m.ExcPending
	if v3730 != 0 {
		goto L1
	} else {
		goto L1021
	}
L1021:
	;
	goto L1019
L1022:
	;
	F_PushActiveSnapshot(m, v3731)
	mBase = m.M
	v3734 = m.ExcPending
	if v3734 != 0 {
		goto L1
	} else {
		goto L1023
	}
L1023:
	;
	v3737 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v3737
	v3740 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[34])) = v3740
	v3743 = *(*int64)(unsafe.Add(mBase, uint32(v22)+352))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[35])) = v3743
	F_FinishPreparedTransaction(m, v22+int32(1360), int32(1))
	mBase = m.M
	v3749 = m.ExcPending
	if v3749 != 0 {
		goto L1
	} else {
		goto L1024
	}
L1024:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v3751 = m.ExcPending
	if v3751 != 0 {
		goto L1
	} else {
		goto L1025
	}
L1025:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v3753 = m.ExcPending
	if v3753 != 0 {
		goto L1
	} else {
		goto L1026
	}
L1026:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3755 = m.ExcPending
	if v3755 != 0 {
		goto L1
	} else {
		goto L1027
	}
L1027:
	;
	v3757 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v3758 = m.ExcPending
	if v3758 != 0 {
		goto L1
	} else {
		goto L1028
	}
L1028:
	;
	v3760 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[37]))
	v3761 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	v3763 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v3764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3763)+16)))
	if v3764 == int32(1) {
		goto L1030
	} else {
		goto L1031
	}
L1029:
	;
	v3803 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])) = uint8(v3803)
	F_ProcessSyncingRelations(m, v3801)
	mBase = m.M
	v3806 = m.ExcPending
	if v3806 != 0 {
		goto L1
	} else {
		goto L1039
	}
L1030:
	;
	v3767 = *(*int32)(unsafe.Add(mBase, uint32(v3763)))
	if v3767 == int32(4) {
		v3801 = v3761
		goto L1029
	} else {
		goto L1033
	}
L1031:
	;
	goto L1032
L1032:
	;
	v3772 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v3772
	v3775 = F_palloc(m, int32(24))
	mBase = m.M
	v3776 = m.ExcPending
	if v3776 != 0 {
		goto L1
	} else {
		goto L1034
	}
L1033:
	;
	goto L1032
L1034:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3775)+16)) = v3761
	*(*int64)(unsafe.Add(mBase, uint32(v3775)+8)) = v3760
	v3780 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[38]))
	if v3780 != 0 {
		goto L1036
	} else {
		goto L1037
	}
L1035:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3775))) = v3787
	v3789 = int32(_a_F_apply_dispatch_53)
	*(*int32)(unsafe.Add(mBase, uint32(v3775)+4)) = v3789
	*(*int32)(unsafe.Add(mBase, uint32(v3787)+4)) = v3775
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[39])) = v3775
	v3796 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v3796
	v3798 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	v3801 = v3798
	goto L1029
L1036:
	;
	v3782 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[39]))
	v3787 = v3782
	goto L1035
L1037:
	;
	goto L1038
L1038:
	;
	v3784 = int32(_a_F_apply_dispatch_53)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[38])) = v3784
	v3787 = v3784
	goto L1035
L1039:
	;
	v3807 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	F_clear_subscription_skip_lsn(m, v3807)
	mBase = m.M
	v3809 = m.ExcPending
	if v3809 != 0 {
		goto L1
	} else {
		goto L1040
	}
L1040:
	;
	v3811 = int32(0)
	F_pgstat_report_activity(m, int32(2), v3811)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v3811
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v3811
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v3811
	goto L3
L1041:
	;
	v4026 = *(*int32)(unsafe.Add(mBase, uint32(v22)+368))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v4026
	v4029 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v4029
	v4032 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v4033 = *(*int32)(unsafe.Add(mBase, uint32(v4032)+4))
	v4035 = v22 + int32(1360)
	F_TwoPhaseTransactionGid(m, v4033, v4026, v4035)
	mBase = m.M
	v4037 = m.ExcPending
	if v4037 != 0 {
		goto L1
	} else {
		goto L1096
	}
L1042:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4015 = m.ExcPending
	if v4015 != 0 {
		goto L1
	} else {
		goto L1093
	}
L1043:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4002 = m.ExcPending
	if v4002 != 0 {
		goto L1
	} else {
		goto L1090
	}
L1044:
	;
	v3837 = v3834 & int32(255)
	if v3837 == int32(0) {
		goto L1045
	} else {
		goto L1046
	}
L1045:
	;
	v3840 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3841 = m.ExcPending
	if v3841 != 0 {
		goto L1
	} else {
		goto L1048
	}
L1046:
	;
	goto L1047
L1047:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3989 = m.ExcPending
	if v3989 != 0 {
		goto L1
	} else {
		goto L1087
	}
L1048:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3829))) = v3840
	if v3840 == int64(0) {
		goto L1043
	} else {
		goto L1049
	}
L1049:
	;
	v3845 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3846 = m.ExcPending
	if v3846 != 0 {
		goto L1
	} else {
		goto L1050
	}
L1050:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3829)+8)) = v3845
	if v3845 == int64(0) {
		goto L1042
	} else {
		goto L1051
	}
L1051:
	;
	v3850 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3851 = m.ExcPending
	if v3851 != 0 {
		goto L1
	} else {
		goto L1052
	}
L1052:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3829)+16)) = v3850
	v3853 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3854 = m.ExcPending
	if v3854 != 0 {
		goto L1
	} else {
		goto L1053
	}
L1053:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3829)+24)) = v3853
	v3857 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v3858 = m.ExcPending
	if v3858 != 0 {
		goto L1
	} else {
		goto L1054
	}
L1054:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3829)+32)) = v3857
	v3861 = v22 + int32(372)
	v3862 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v3863 = m.ExcPending
	if v3863 != 0 {
		goto L1
	} else {
		goto L1055
	}
L1055:
	;
	goto L1059
L1056:
	;
	m.G0 = v3832 + int32(16)
	goto L1041
L1057:
	;
	v3980 = F_strlen(m, v3969)
	mBase = m.M
	goto L1056
L1059:
	;
	goto L1060
L1060:
	;
	v3870 = int32(199)
	if (v3861^v3862)&int32(3) != 0 {
		goto L1064
	} else {
		goto L1065
	}
L1061:
	;
	v3973 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3970))) = uint8(v3973)
	goto L1057
L1062:
	;
	v3954 = v3949
	v3955 = v3950
	v3956 = v3951
	goto L1083
L1063:
	;
	if v3944 == int32(0) {
		v3969 = v3942
		v3970 = v3943
		goto L1061
	} else {
		goto L1082
	}
L1064:
	;
	v3942 = v3862
	v3943 = v3861
	v3944 = v3870
	goto L1063
L1065:
	;
	goto L1066
L1066:
	;
	v3874 = int32(0)
	if base.B2i32(v3862&int32(3) == v3874)|int32(0) == v3874 {
		goto L1068
	} else {
		goto L1069
	}
L1067:
	;
	if v3910 == int32(0) {
		v3969 = v3907
		v3970 = v3908
		goto L1061
	} else {
		goto L1076
	}
L1068:
	;
	v3886 = v3862
	v3887 = v3861
	v3888 = v3870
	goto L1071
L1069:
	;
	goto L1070
L1070:
	;
	v3907 = v3862
	v3908 = v3861
	v3909 = v3870
	v3910 = int32(1)
	goto L1067
L1071:
	;
	v3890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3886))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3887))) = uint8(v3890)
	if v3890 == int32(0) {
		v3949 = v3886
		v3950 = v3887
		v3951 = v3888
		goto L1062
	} else {
		goto L1073
	}
L1072:
	;
	v3907 = v3901
	v3908 = v3895
	v3909 = v3897
	v3910 = v3899
	goto L1067
L1073:
	;
	v3894 = int32(1)
	v3895 = v3887 + v3894
	v3897 = v3888 - v3894
	v3898 = int32(0)
	v3899 = base.B2i32(v3897 != v3898)
	v3901 = v3886 + v3894
	if v3901&int32(3) == v3898 {
		v3907 = v3901
		v3908 = v3895
		v3909 = v3897
		v3910 = v3899
		goto L1067
	} else {
		goto L1074
	}
L1074:
	;
	if v3897 != 0 {
		v3886 = v3901
		v3887 = v3895
		v3888 = v3897
		goto L1071
	} else {
		goto L1075
	}
L1075:
	;
	goto L1072
L1076:
	;
	v3913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3907))))
	if base.B2i32(v3913 == int32(0))|base.B2i32(base.Ui32(v3909) < base.Ui32(int32(4))) != 0 {
		v3942 = v3907
		v3943 = v3908
		v3944 = v3909
		goto L1063
	} else {
		goto L1077
	}
L1077:
	;
	v3920 = v3907
	v3921 = v3908
	v3922 = v3909
	goto L1078
L1078:
	;
	v3925 = *(*int32)(unsafe.Add(mBase, uint32(v3920)))
	v3928 = int32(-2139062144)
	if (int32(16843008)-v3925|v3925)&v3928 != v3928 {
		v3949 = v3920
		v3950 = v3921
		v3951 = v3922
		goto L1062
	} else {
		goto L1080
	}
L1079:
	;
	v3942 = v3936
	v3943 = v3934
	v3944 = v3938
	goto L1063
L1080:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3921))) = v3925
	v3933 = int32(4)
	v3934 = v3921 + v3933
	v3936 = v3920 + v3933
	v3938 = v3922 - v3933
	if base.Ui32(int32(3)) < base.Ui32(v3938) {
		v3920 = v3936
		v3921 = v3934
		v3922 = v3938
		goto L1078
	} else {
		goto L1081
	}
L1081:
	;
	goto L1079
L1082:
	;
	v3949 = v3942
	v3950 = v3943
	v3951 = v3944
	goto L1062
L1083:
	;
	v3958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3954))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3955))) = uint8(v3958)
	if v3958 == int32(0) {
		v3969 = v3954
		v3970 = v3955
		goto L1061
	} else {
		goto L1085
	}
L1084:
	;
	v3969 = v3965
	v3970 = v3963
	goto L1061
L1085:
	;
	v3962 = int32(1)
	v3963 = v3955 + v3962
	v3965 = v3954 + v3962
	v3967 = v3956 - v3962
	if v3967 != 0 {
		v3954 = v3965
		v3955 = v3963
		v3956 = v3967
		goto L1083
	} else {
		goto L1086
	}
L1086:
	;
	goto L1084
L1087:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3832))) = v3837
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_61), v3832)
	mBase = m.M
	v3993 = m.ExcPending
	if v3993 != 0 {
		goto L1
	} else {
		goto L1088
	}
L1088:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(332), int32(_a_F_apply_dispatch_62))
	mBase = m.M
	v3998 = m.ExcPending
	if v3998 != 0 {
		goto L1
	} else {
		goto L1089
	}
L1089:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1090:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_63), int32(0))
	mBase = m.M
	v4006 = m.ExcPending
	if v4006 != 0 {
		goto L1
	} else {
		goto L1091
	}
L1091:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(337), int32(_a_F_apply_dispatch_62))
	mBase = m.M
	v4011 = m.ExcPending
	if v4011 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_64), int32(0))
	mBase = m.M
	v4019 = m.ExcPending
	if v4019 != 0 {
		goto L1
	} else {
		goto L1094
	}
L1094:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(340), int32(_a_F_apply_dispatch_62))
	mBase = m.M
	v4024 = m.ExcPending
	if v4024 != 0 {
		goto L1
	} else {
		goto L1095
	}
L1095:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1096:
	;
	v4038 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	v4039 = *(*int64)(unsafe.Add(mBase, uint32(v22)+352))
	v4040 = int32(0)
	v4041 = m.G0
	v4043 = v4041 - int32(16)
	m.G0 = v4043
	v4046 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[40]))
	v4050 = F_LWLockAcquire(m, v4046+int32(2304), int32(1))
	mBase = m.M
	v4051 = m.ExcPending
	if v4051 != 0 {
		goto L1
	} else {
		goto L1097
	}
L1097:
	;
	v4053 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[41]))
	v4054 = *(*int32)(unsafe.Add(mBase, uint32(v4053)+4))
	if v4054 <= int32(0) {
		v4144 = v4040
		goto L1098
	} else {
		goto L1099
	}
L1098:
	;
	v4163 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[40]))
	F_LWLockRelease(m, v4163+int32(2304))
	mBase = m.M
	v4167 = m.ExcPending
	if v4167 != 0 {
		goto L1
	} else {
		goto L1126
	}
L1099:
	;
	v4057 = v4053
	v4058 = v4040
	goto L1101
L1100:
	;
	F_pfree(m, v4125)
	mBase = m.M
	v4141 = m.ExcPending
	if v4141 != 0 {
		goto L1
	} else {
		goto L1125
	}
L1101:
	;
	v4079 = *(*int32)(unsafe.Add(mBase, uint32(v4057+v4058<<(uint(int32(2))%32))+8))
	v4080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4079)+48)))
	if v4080 != int32(1) {
		v4134 = v4057
		goto L1103
	} else {
		goto L1104
	}
L1102:
	;
	v4144 = int32(0)
	goto L1098
L1103:
	;
	v4136 = v4058 + int32(1)
	v4137 = *(*int32)(unsafe.Add(mBase, uint32(v4134)+4))
	if v4136 < v4137 {
		v4057 = v4134
		v4058 = v4136
		goto L1101
	} else {
		goto L1124
	}
L1104:
	;
	v4084 = v4079 + int32(51)
	v4087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4084))))
	v4090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4035))))
	if base.B2i32(v4087 == int32(0))|base.B2i32(v4087 != v4090) != 0 {
		v4108 = v4087
		v4109 = v4090
		goto L1106
	} else {
		goto L1107
	}
L1105:
	;
	if v4108-v4109 != 0 {
		v4134 = v4057
		goto L1103
	} else {
		goto L1112
	}
L1106:
	;
	goto L1105
L1107:
	;
	v4093 = v4084
	v4094 = v4035
	goto L1108
L1108:
	;
	v4097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4094)+1)))
	v4098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4093)+1)))
	if v4098 == int32(0) {
		v4108 = v4098
		v4109 = v4097
		goto L1106
	} else {
		goto L1110
	}
L1109:
	;
	v4108 = v4098
	v4109 = v4097
	goto L1106
L1110:
	;
	v4101 = int32(1)
	if v4098 == v4097 {
		v4093 = v4093 + v4101
		v4094 = v4094 + v4101
		goto L1108
	} else {
		goto L1111
	}
L1111:
	;
	goto L1109
L1112:
	;
	v4111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4079)+49)))
	if v4111 == int32(1) {
		goto L1114
	} else {
		goto L1115
	}
L1113:
	;
	v4126 = *(*int64)(unsafe.Add(mBase, uint32(v4125)+56))
	if v4038 == v4126 {
		goto L1119
	} else {
		goto L1120
	}
L1114:
	;
	v4114 = *(*int64)(unsafe.Add(mBase, uint32(v4079)+32))
	v4116 = F_ReadTwoPhaseFile(m, v4114, int32(0))
	mBase = m.M
	v4117 = m.ExcPending
	if v4117 != 0 {
		goto L1
	} else {
		goto L1117
	}
L1115:
	;
	goto L1116
L1116:
	;
	v4118 = *(*int64)(unsafe.Add(mBase, uint32(v4079)+16))
	F_XlogReadTwoPhaseData(m, v4118, v4043+int32(12), int32(0))
	mBase = m.M
	v4123 = m.ExcPending
	if v4123 != 0 {
		goto L1
	} else {
		goto L1118
	}
L1117:
	;
	v4125 = v4116
	goto L1113
L1118:
	;
	v4124 = *(*int32)(unsafe.Add(mBase, uint32(v4043)+12))
	v4125 = v4124
	goto L1113
L1119:
	;
	v4128 = *(*int64)(unsafe.Add(mBase, uint32(v4125)+64))
	if v4128 == v4039 {
		goto L1100
	} else {
		goto L1122
	}
L1120:
	;
	goto L1121
L1121:
	;
	F_pfree(m, v4125)
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		goto L1
	} else {
		goto L1123
	}
L1122:
	;
	goto L1121
L1123:
	;
	v4133 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[41]))
	v4134 = v4133
	goto L1103
L1124:
	;
	goto L1102
L1125:
	;
	v4144 = int32(1)
	goto L1098
L1126:
	;
	m.G0 = v4043 + int32(16)
	if v4144 != 0 {
		goto L1127
	} else {
		goto L1128
	}
L1127:
	;
	v4172 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[34])) = v4172
	v4175 = *(*int64)(unsafe.Add(mBase, uint32(v22)+360))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[35])) = v4175
	v4178 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v4178 < int32(0) {
		goto L1131
	} else {
		goto L1132
	}
L1128:
	;
	goto L1129
L1129:
	;
	v4218 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v4219 = m.ExcPending
	if v4219 != 0 {
		goto L1
	} else {
		goto L1147
	}
L1130:
	;
	v4185 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v4186 = *(*int32)(unsafe.Add(mBase, uint32(v4185)+20))
	goto L1134
L1131:
	;
	v4182 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v4182
	goto L1133
L1132:
	;
	goto L1133
L1133:
	;
	goto L1130
L1134:
	;
	if base.B2i32(v4186 == int32(2)) == int32(0) {
		goto L1135
	} else {
		goto L1136
	}
L1135:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v4192 = m.ExcPending
	if v4192 != 0 {
		goto L1
	} else {
		goto L1138
	}
L1136:
	;
	goto L1137
L1137:
	;
	v4195 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v4196 = m.ExcPending
	if v4196 != 0 {
		goto L1
	} else {
		goto L1140
	}
L1138:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v4194 = m.ExcPending
	if v4194 != 0 {
		goto L1
	} else {
		goto L1139
	}
L1139:
	;
	goto L1137
L1140:
	;
	F_PushActiveSnapshot(m, v4195)
	mBase = m.M
	v4198 = m.ExcPending
	if v4198 != 0 {
		goto L1
	} else {
		goto L1141
	}
L1141:
	;
	v4201 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v4201
	F_FinishPreparedTransaction(m, v22+int32(1360), int32(0))
	mBase = m.M
	v4207 = m.ExcPending
	if v4207 != 0 {
		goto L1
	} else {
		goto L1142
	}
L1142:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v4209 = m.ExcPending
	if v4209 != 0 {
		goto L1
	} else {
		goto L1143
	}
L1143:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v4211 = m.ExcPending
	if v4211 != 0 {
		goto L1
	} else {
		goto L1144
	}
L1144:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v4213 = m.ExcPending
	if v4213 != 0 {
		goto L1
	} else {
		goto L1145
	}
L1145:
	;
	v4214 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	F_clear_subscription_skip_lsn(m, v4214)
	mBase = m.M
	v4216 = m.ExcPending
	if v4216 != 0 {
		goto L1
	} else {
		goto L1146
	}
L1146:
	;
	goto L1129
L1147:
	;
	v4220 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	v4222 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v4223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4222)+16)))
	if v4223 == int32(1) {
		goto L1149
	} else {
		goto L1150
	}
L1148:
	;
	v4263 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])) = uint8(v4263)
	F_ProcessSyncingRelations(m, v4261)
	mBase = m.M
	v4266 = m.ExcPending
	if v4266 != 0 {
		goto L1
	} else {
		goto L1158
	}
L1149:
	;
	v4226 = *(*int32)(unsafe.Add(mBase, uint32(v4222)))
	if v4226 == int32(4) {
		v4261 = v4220
		goto L1148
	} else {
		goto L1152
	}
L1150:
	;
	goto L1151
L1151:
	;
	v4231 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v4231
	v4234 = F_palloc(m, int32(24))
	mBase = m.M
	v4235 = m.ExcPending
	if v4235 != 0 {
		goto L1
	} else {
		goto L1153
	}
L1152:
	;
	goto L1151
L1153:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4234)+16)) = v4220
	*(*int64)(unsafe.Add(mBase, uint32(v4234)+8)) = int64(0)
	v4240 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[38]))
	if v4240 != 0 {
		goto L1155
	} else {
		goto L1156
	}
L1154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4234))) = v4247
	v4249 = int32(_a_F_apply_dispatch_53)
	*(*int32)(unsafe.Add(mBase, uint32(v4234)+4)) = v4249
	*(*int32)(unsafe.Add(mBase, uint32(v4247)+4)) = v4234
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[39])) = v4234
	v4256 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v4256
	v4258 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	v4261 = v4258
	goto L1148
L1155:
	;
	v4242 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[39]))
	v4247 = v4242
	goto L1154
L1156:
	;
	goto L1157
L1157:
	;
	v4244 = int32(_a_F_apply_dispatch_53)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[38])) = v4244
	v4247 = v4244
	goto L1154
L1158:
	;
	v4268 = int32(0)
	F_pgstat_report_activity(m, int32(2), v4268)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v4268
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v4268
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v4268
	goto L3
L1159:
	;
	v4292 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v4293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4292)+16)))
	if v4293 == int32(1) {
		goto L1160
	} else {
		goto L1161
	}
L1160:
	;
	v4296 = *(*int32)(unsafe.Add(mBase, uint32(v4292)))
	if v4296 == int32(1) {
		goto L15
	} else {
		goto L1163
	}
L1161:
	;
	goto L1162
L1162:
	;
	F_logicalrep_read_prepare_common(m, l0, int32(_a_F_apply_dispatch_65), v22+int32(336))
	mBase = m.M
	v4303 = m.ExcPending
	if v4303 != 0 {
		goto L1
	} else {
		goto L1164
	}
L1163:
	;
	goto L1162
L1164:
	;
	v4305 = *(*int32)(unsafe.Add(mBase, uint32(v22)+360))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v4305
	v4308 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v4308
	v4311 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v4312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4311)+16)))
	if v4312 == int32(1) {
		goto L1167
	} else {
		goto L1168
	}
L1165:
	;
	v4464 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v4465 = m.ExcPending
	if v4465 != 0 {
		goto L1
	} else {
		goto L1222
	}
L1166:
	;
	v4389 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25]))
	if v4389 != 0 {
		goto L1195
	} else {
		goto L1196
	}
L1167:
	;
	v4315 = *(*int32)(unsafe.Add(mBase, uint32(v4311)))
	if v4315 == int32(4) {
		goto L1166
	} else {
		goto L1170
	}
L1168:
	;
	goto L1169
L1169:
	;
	v4318 = F_pa_find_worker(m, v4305)
	mBase = m.M
	v4319 = m.ExcPending
	if v4319 != 0 {
		goto L1
	} else {
		goto L1173
	}
L1170:
	;
	goto L1169
L1171:
	;
	v4376 = *(*int32)(unsafe.Add(mBase, uint32(v22)+360))
	F_stream_open_and_write_change(m, v4376, int32(112), v22+int32(1360))
	mBase = m.M
	v4381 = m.ExcPending
	if v4381 != 0 {
		goto L1
	} else {
		goto L1192
	}
L1172:
	;
	F_pa_switch_to_partial_serialize(m, v4318, int32(1))
	mBase = m.M
	v4375 = m.ExcPending
	if v4375 != 0 {
		goto L1
	} else {
		goto L1191
	}
L1173:
	;
	if v4318 != 0 {
		goto L1174
	} else {
		goto L1175
	}
L1174:
	;
	v4320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4318)+12)))
	if v4320 != 0 {
		goto L1171
	} else {
		goto L1177
	}
L1175:
	;
	goto L1176
L1176:
	;
	v4331 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v4331 != 0 {
		goto L14
	} else {
		goto L1181
	}
L1177:
	;
	v4321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4322 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4323 = F_pa_send_data(m, v4318, v4321, v4322)
	mBase = m.M
	v4324 = m.ExcPending
	if v4324 != 0 {
		goto L1
	} else {
		goto L1178
	}
L1178:
	;
	if v4323 == int32(0) {
		goto L1172
	} else {
		goto L1179
	}
L1179:
	;
	v4327 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	F_pa_xact_finish(m, v4318, v4327)
	mBase = m.M
	v4329 = m.ExcPending
	if v4329 != 0 {
		goto L1
	} else {
		goto L1180
	}
L1180:
	;
	goto L1165
L1181:
	;
	v4333 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v4334 = *(*int32)(unsafe.Add(mBase, uint32(v4333)+60))
	v4335 = *(*int32)(unsafe.Add(mBase, uint32(v22)+360))
	v4336 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	F_apply_spooled_messages(m, v4334, v4335, v4336)
	mBase = m.M
	v4338 = m.ExcPending
	if v4338 != 0 {
		goto L1
	} else {
		goto L1182
	}
L1182:
	;
	F_apply_handle_prepare_internal(m, v22+int32(336))
	mBase = m.M
	v4342 = m.ExcPending
	if v4342 != 0 {
		goto L1
	} else {
		goto L1183
	}
L1183:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v4344 = m.ExcPending
	if v4344 != 0 {
		goto L1
	} else {
		goto L1184
	}
L1184:
	;
	v4345 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	F_store_flush_position(m, v4345, int64(0))
	mBase = m.M
	v4348 = m.ExcPending
	if v4348 != 0 {
		goto L1
	} else {
		goto L1185
	}
L1185:
	;
	v4350 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])) = uint8(v4350)
	v4353 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v4354 = *(*int32)(unsafe.Add(mBase, uint32(v4353)+32))
	v4355 = *(*int32)(unsafe.Add(mBase, uint32(v22)+360))
	F_stream_cleanup_files(m, v4354, v4355)
	mBase = m.M
	v4357 = m.ExcPending
	if v4357 != 0 {
		goto L1
	} else {
		goto L1186
	}
L1186:
	;
	v4360 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v4361 = m.ExcPending
	if v4361 != 0 {
		goto L1
	} else {
		goto L1187
	}
L1187:
	;
	if v4360 == int32(0) {
		goto L1165
	} else {
		goto L1188
	}
L1188:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_66), int32(0))
	mBase = m.M
	v4367 = m.ExcPending
	if v4367 != 0 {
		goto L1
	} else {
		goto L1189
	}
L1189:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1596), int32(_a_F_apply_dispatch_67))
	mBase = m.M
	v4372 = m.ExcPending
	if v4372 != 0 {
		goto L1
	} else {
		goto L1190
	}
L1190:
	;
	goto L1165
L1191:
	;
	goto L1171
L1192:
	;
	v4382 = *(*int32)(unsafe.Add(mBase, uint32(v4318)+16))
	F_pa_set_fileset_state(m, v4382)
	mBase = m.M
	v4384 = m.ExcPending
	if v4384 != 0 {
		goto L1
	} else {
		goto L1193
	}
L1193:
	;
	v4385 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	F_pa_xact_finish(m, v4318, v4385)
	mBase = m.M
	v4387 = m.ExcPending
	if v4387 != 0 {
		goto L1
	} else {
		goto L1194
	}
L1194:
	;
	goto L1165
L1195:
	;
	F_BufFileClose(m, v4389)
	mBase = m.M
	v4391 = m.ExcPending
	if v4391 != 0 {
		goto L1
	} else {
		goto L1198
	}
L1196:
	;
	goto L1197
L1197:
	;
	v4396 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v4396 < int32(0) {
		goto L1200
	} else {
		goto L1201
	}
L1198:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25])) = int32(0)
	goto L1197
L1199:
	;
	v4403 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v4404 = *(*int32)(unsafe.Add(mBase, uint32(v4403)+20))
	goto L1203
L1200:
	;
	v4400 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v4400
	goto L1202
L1201:
	;
	goto L1202
L1202:
	;
	goto L1199
L1203:
	;
	if base.B2i32(v4404 == int32(2)) == int32(0) {
		goto L1204
	} else {
		goto L1205
	}
L1204:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v4410 = m.ExcPending
	if v4410 != 0 {
		goto L1
	} else {
		goto L1207
	}
L1205:
	;
	goto L1206
L1206:
	;
	v4413 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v4414 = m.ExcPending
	if v4414 != 0 {
		goto L1
	} else {
		goto L1209
	}
L1207:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v4412 = m.ExcPending
	if v4412 != 0 {
		goto L1
	} else {
		goto L1208
	}
L1208:
	;
	goto L1206
L1209:
	;
	F_PushActiveSnapshot(m, v4413)
	mBase = m.M
	v4416 = m.ExcPending
	if v4416 != 0 {
		goto L1
	} else {
		goto L1210
	}
L1210:
	;
	v4419 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v4419
	F_apply_handle_prepare_internal(m, v22+int32(336))
	mBase = m.M
	v4424 = m.ExcPending
	if v4424 != 0 {
		goto L1
	} else {
		goto L1211
	}
L1211:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v4426 = m.ExcPending
	if v4426 != 0 {
		goto L1
	} else {
		goto L1212
	}
L1212:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v4428 = m.ExcPending
	if v4428 != 0 {
		goto L1
	} else {
		goto L1213
	}
L1213:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v4430 = m.ExcPending
	if v4430 != 0 {
		goto L1
	} else {
		goto L1214
	}
L1214:
	;
	v4432 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[27]))
	*(*int64)(unsafe.Add(mBase, uint32(v4432)+24)) = int64(0)
	F_pa_set_xact_state(m, v4432, int32(2))
	mBase = m.M
	v4437 = m.ExcPending
	if v4437 != 0 {
		goto L1
	} else {
		goto L1215
	}
L1215:
	;
	v4439 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[27]))
	v4440 = *(*int32)(unsafe.Add(mBase, uint32(v4439)+4))
	F_pa_unlock_transaction(m, v4440)
	mBase = m.M
	v4442 = m.ExcPending
	if v4442 != 0 {
		goto L1
	} else {
		goto L1216
	}
L1216:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[36])) = int32(0)
	goto L1217
L1217:
	;
	v4448 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v4449 = m.ExcPending
	if v4449 != 0 {
		goto L1
	} else {
		goto L1218
	}
L1218:
	;
	if v4448 == int32(0) {
		goto L1165
	} else {
		goto L1219
	}
L1219:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_66), int32(0))
	mBase = m.M
	v4455 = m.ExcPending
	if v4455 != 0 {
		goto L1
	} else {
		goto L1220
	}
L1220:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1658), int32(_a_F_apply_dispatch_67))
	mBase = m.M
	v4460 = m.ExcPending
	if v4460 != 0 {
		goto L1
	} else {
		goto L1221
	}
L1221:
	;
	goto L1165
L1222:
	;
	v4466 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	F_ProcessSyncingRelations(m, v4466)
	mBase = m.M
	v4468 = m.ExcPending
	if v4468 != 0 {
		goto L1
	} else {
		goto L1223
	}
L1223:
	;
	v4470 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	if v4470 != int64(0) {
		goto L1224
	} else {
		goto L1225
	}
L1224:
	;
	v4475 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4476 = m.ExcPending
	if v4476 != 0 {
		goto L1
	} else {
		goto L1227
	}
L1225:
	;
	goto L1226
L1226:
	;
	v4498 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	F_clear_subscription_skip_lsn(m, v4498)
	mBase = m.M
	v4500 = m.ExcPending
	if v4500 != 0 {
		goto L1
	} else {
		goto L1233
	}
L1227:
	;
	if v4475 != 0 {
		goto L1228
	} else {
		goto L1229
	}
L1228:
	;
	v4478 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+260)) = uint32(v4478)
	v4481 = int64(base.Ui64(v4478) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+256)) = uint32(v4481)
	F_errmsg(m, int32(_a_F_apply_dispatch_54), v22+int32(256))
	mBase = m.M
	v4487 = m.ExcPending
	if v4487 != 0 {
		goto L1
	} else {
		goto L1231
	}
L1229:
	;
	goto L1230
L1230:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8])) = int64(0)
	goto L1226
L1231:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(_a_F_apply_dispatch_55), int32(_a_F_apply_dispatch_56))
	mBase = m.M
	v4492 = m.ExcPending
	if v4492 != 0 {
		goto L1
	} else {
		goto L1232
	}
L1232:
	;
	goto L1230
L1233:
	;
	v4502 = int32(0)
	F_pgstat_report_activity(m, int32(2), v4502)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v4502
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v4502
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v4502
	goto L3
L1234:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4525 = m.ExcPending
	if v4525 != 0 {
		goto L1
	} else {
		goto L1235
	}
L1235:
	;
	v4526 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+36)) = uint32(v4526)
	v4529 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[9]))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+44)) = uint32(v4529)
	v4531 = int64(32)
	v4532 = int64(base.Ui64(v4526) >> (uint(v4531) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+32)) = uint32(v4532)
	v4535 = int64(base.Ui64(v4529) >> (uint(v4531) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+40)) = uint32(v4535)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_68), v22+int32(32))
	mBase = m.M
	v4541 = m.ExcPending
	if v4541 != 0 {
		goto L1
	} else {
		goto L1236
	}
L1236:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1273), int32(_a_F_apply_dispatch_69))
	mBase = m.M
	v4546 = m.ExcPending
	if v4546 != 0 {
		goto L1
	} else {
		goto L1237
	}
L1237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1238:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4553 = m.ExcPending
	if v4553 != 0 {
		goto L1
	} else {
		goto L1239
	}
L1239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v672 + int32(1)
	v4557 = *(*int32)(unsafe.Add(mBase, uint32(v22)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v4557
	F_errmsg_plural(m, int32(_a_F_apply_dispatch_70), int32(_a_F_apply_dispatch_71), v4557, v22-int32(-64))
	mBase = m.M
	v4564 = m.ExcPending
	if v4564 != 0 {
		goto L1
	} else {
		goto L1240
	}
L1240:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(2898), int32(_a_F_apply_dispatch_72))
	mBase = m.M
	v4569 = m.ExcPending
	if v4569 != 0 {
		goto L1
	} else {
		goto L1241
	}
L1241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1242:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4576 = m.ExcPending
	if v4576 != 0 {
		goto L1
	} else {
		goto L1243
	}
L1243:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_73), int32(0))
	mBase = m.M
	v4580 = m.ExcPending
	if v4580 != 0 {
		goto L1
	} else {
		goto L1244
	}
L1244:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1763), int32(_a_F_apply_dispatch_74))
	mBase = m.M
	v4585 = m.ExcPending
	if v4585 != 0 {
		goto L1
	} else {
		goto L1245
	}
L1245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1246:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4592 = m.ExcPending
	if v4592 != 0 {
		goto L1
	} else {
		goto L1247
	}
L1247:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_75), int32(0))
	mBase = m.M
	v4596 = m.ExcPending
	if v4596 != 0 {
		goto L1
	} else {
		goto L1248
	}
L1248:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1777), int32(_a_F_apply_dispatch_74))
	mBase = m.M
	v4601 = m.ExcPending
	if v4601 != 0 {
		goto L1
	} else {
		goto L1249
	}
L1249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = int32(0)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_46), v22+int32(80))
	mBase = m.M
	v4612 = m.ExcPending
	if v4612 != 0 {
		goto L1
	} else {
		goto L1251
	}
L1251:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1874), int32(_a_F_apply_dispatch_74))
	mBase = m.M
	v4617 = m.ExcPending
	if v4617 != 0 {
		goto L1
	} else {
		goto L1252
	}
L1252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1253:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4624 = m.ExcPending
	if v4624 != 0 {
		goto L1
	} else {
		goto L1254
	}
L1254:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_76), int32(0))
	mBase = m.M
	v4628 = m.ExcPending
	if v4628 != 0 {
		goto L1
	} else {
		goto L1255
	}
L1255:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1919), int32(_a_F_apply_dispatch_32))
	mBase = m.M
	v4633 = m.ExcPending
	if v4633 != 0 {
		goto L1
	} else {
		goto L1256
	}
L1256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = int32(0)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_46), v22+int32(112))
	mBase = m.M
	v4644 = m.ExcPending
	if v4644 != 0 {
		goto L1
	} else {
		goto L1258
	}
L1258:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1990), int32(_a_F_apply_dispatch_32))
	mBase = m.M
	v4649 = m.ExcPending
	if v4649 != 0 {
		goto L1
	} else {
		goto L1259
	}
L1259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1260:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4656 = m.ExcPending
	if v4656 != 0 {
		goto L1
	} else {
		goto L1261
	}
L1261:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_77), int32(0))
	mBase = m.M
	v4660 = m.ExcPending
	if v4660 != 0 {
		goto L1
	} else {
		goto L1262
	}
L1262:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(2112), int32(_a_F_apply_dispatch_41))
	mBase = m.M
	v4665 = m.ExcPending
	if v4665 != 0 {
		goto L1
	} else {
		goto L1263
	}
L1263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1264:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4672 = m.ExcPending
	if v4672 != 0 {
		goto L1
	} else {
		goto L1265
	}
L1265:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_78), int32(0))
	mBase = m.M
	v4676 = m.ExcPending
	if v4676 != 0 {
		goto L1
	} else {
		goto L1266
	}
L1266:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(2429), int32(_a_F_apply_dispatch_49))
	mBase = m.M
	v4681 = m.ExcPending
	if v4681 != 0 {
		goto L1
	} else {
		goto L1267
	}
L1267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = int32(1)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_46), v22+int32(192))
	mBase = m.M
	v4692 = m.ExcPending
	if v4692 != 0 {
		goto L1
	} else {
		goto L1269
	}
L1269:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(2510), int32(_a_F_apply_dispatch_49))
	mBase = m.M
	v4697 = m.ExcPending
	if v4697 != 0 {
		goto L1
	} else {
		goto L1270
	}
L1270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1271:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4704 = m.ExcPending
	if v4704 != 0 {
		goto L1
	} else {
		goto L1272
	}
L1272:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_79), int32(0))
	mBase = m.M
	v4708 = m.ExcPending
	if v4708 != 0 {
		goto L1
	} else {
		goto L1273
	}
L1273:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1299), int32(_a_F_apply_dispatch_80))
	mBase = m.M
	v4713 = m.ExcPending
	if v4713 != 0 {
		goto L1
	} else {
		goto L1274
	}
L1274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1275:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4720 = m.ExcPending
	if v4720 != 0 {
		goto L1
	} else {
		goto L1276
	}
L1276:
	;
	v4721 = *(*int64)(unsafe.Add(mBase, uint32(v22)+336))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+244)) = uint32(v4721)
	v4724 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[9]))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+252)) = uint32(v4724)
	v4726 = int64(32)
	v4727 = int64(base.Ui64(v4721) >> (uint(v4726) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+240)) = uint32(v4727)
	v4730 = int64(base.Ui64(v4724) >> (uint(v4726) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+248)) = uint32(v4730)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_81), v22+int32(240))
	mBase = m.M
	v4736 = m.ExcPending
	if v4736 != 0 {
		goto L1
	} else {
		goto L1277
	}
L1277:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1368), int32(_a_F_apply_dispatch_82))
	mBase = m.M
	v4741 = m.ExcPending
	if v4741 != 0 {
		goto L1
	} else {
		goto L1278
	}
L1278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1279:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4748 = m.ExcPending
	if v4748 != 0 {
		goto L1
	} else {
		goto L1280
	}
L1280:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_83), int32(0))
	mBase = m.M
	v4752 = m.ExcPending
	if v4752 != 0 {
		goto L1
	} else {
		goto L1281
	}
L1281:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1556), int32(_a_F_apply_dispatch_67))
	mBase = m.M
	v4757 = m.ExcPending
	if v4757 != 0 {
		goto L1
	} else {
		goto L1282
	}
L1282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1283:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4764 = m.ExcPending
	if v4764 != 0 {
		goto L1
	} else {
		goto L1284
	}
L1284:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_84), int32(0))
	mBase = m.M
	v4768 = m.ExcPending
	if v4768 != 0 {
		goto L1
	} else {
		goto L1285
	}
L1285:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1562), int32(_a_F_apply_dispatch_67))
	mBase = m.M
	v4773 = m.ExcPending
	if v4773 != 0 {
		goto L1
	} else {
		goto L1286
	}
L1286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+272)) = int32(1)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_46), v22+int32(272))
	mBase = m.M
	v4784 = m.ExcPending
	if v4784 != 0 {
		goto L1
	} else {
		goto L1288
	}
L1288:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1662), int32(_a_F_apply_dispatch_67))
	mBase = m.M
	v4789 = m.ExcPending
	if v4789 != 0 {
		goto L1
	} else {
		goto L1289
	}
L1289:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1290:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4796 = m.ExcPending
	if v4796 != 0 {
		goto L1
	} else {
		goto L1291
	}
L1291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v24
	F_errmsg(m, int32(_a_F_apply_dispatch_85), v22)
	mBase = m.M
	v4800 = m.ExcPending
	if v4800 != 0 {
		goto L1
	} else {
		goto L1292
	}
L1292:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(3907), int32(_a_F_apply_dispatch_86))
	mBase = m.M
	v4805 = m.ExcPending
	if v4805 != 0 {
		goto L1
	} else {
		goto L1293
	}
L1293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1294:
	;
	goto L11
L1295:
	;
	if v2312 != v2317 {
		goto L10
	} else {
		goto L1296
	}
L1296:
	;
	v4815 = *(*int32)(unsafe.Add(mBase, uint32(v2326)+16))
	F_pa_set_fileset_state(m, v4815)
	mBase = m.M
	v4817 = m.ExcPending
	if v4817 != 0 {
		goto L1
	} else {
		goto L1297
	}
L1297:
	;
	F_pa_xact_finish(m, v2326, int64(0))
	mBase = m.M
	v4820 = m.ExcPending
	if v4820 != 0 {
		goto L1
	} else {
		goto L1298
	}
L1298:
	;
	goto L10
L1299:
	;
	if v4863 == int32(0) {
		goto L1300
	} else {
		goto L1301
	}
L1300:
	;
	if v4861 == int32(0) {
		goto L4
	} else {
		goto L1307
	}
L1301:
	;
	v4887 = int32(0)
	v4888 = *(*int32)(unsafe.Add(mBase, uint32(v4863)+4))
	if v4888 <= v4887 {
		goto L1300
	} else {
		goto L1302
	}
L1302:
	;
	v4891 = v4887
	goto L1303
L1303:
	;
	v4910 = *(*int32)(unsafe.Add(mBase, uint32(v4863)+12))
	v4914 = *(*int32)(unsafe.Add(mBase, uint32(v4910+v4891<<(uint(int32(2))%32))))
	F_logicalrep_rel_close(m, v4914, int32(0))
	mBase = m.M
	v4917 = m.ExcPending
	if v4917 != 0 {
		goto L1
	} else {
		goto L1305
	}
L1304:
	;
	goto L1300
L1305:
	;
	v4919 = v4891 + int32(1)
	v4920 = *(*int32)(unsafe.Add(mBase, uint32(v4863)+4))
	if v4919 < v4920 {
		v4891 = v4919
		goto L1303
	} else {
		goto L1306
	}
L1306:
	;
	goto L1304
L1307:
	;
	v4943 = *(*int32)(unsafe.Add(mBase, uint32(v4861)+4))
	if v4943 <= int32(0) {
		goto L4
	} else {
		goto L1308
	}
L1308:
	;
	v4947 = int32(0)
	goto L1309
L1309:
	;
	v4966 = *(*int32)(unsafe.Add(mBase, uint32(v4861)+12))
	v4970 = *(*int32)(unsafe.Add(mBase, uint32(v4966+v4947<<(uint(int32(2))%32))))
	F_relation_close(m, v4970, int32(0))
	mBase = m.M
	v4973 = m.ExcPending
	if v4973 != 0 {
		goto L1
	} else {
		goto L1311
	}
L1310:
	;
	goto L4
L1311:
	;
	v4975 = v4947 + int32(1)
	v4976 = *(*int32)(unsafe.Add(mBase, uint32(v4861)+4))
	if v4975 < v4976 {
		v4947 = v4975
		goto L1309
	} else {
		goto L1312
	}
L1312:
	;
	goto L1310
L1313:
	;
	F_slot_store_data(m, v772, v745, v22+int32(292))
	mBase = m.M
	v4996 = m.ExcPending
	if v4996 != 0 {
		goto L1
	} else {
		goto L1319
	}
L1314:
	;
	if v4984 != 0 {
		goto L1315
	} else {
		goto L1316
	}
L1315:
	;
	v4987 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+1372)))
	v4989 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_apply_dispatch[16])))
	if v4987 != v4989 {
		v4992 = int32(3)
		goto L1313
	} else {
		goto L1318
	}
L1316:
	;
	goto L1317
L1317:
	;
	v4992 = int32(4)
	goto L1313
L1318:
	;
	goto L1317
L1319:
	;
	v4998 = v22 + int32(1360)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v4998
	*(*int32)(unsafe.Add(mBase, uint32(v22)+316)) = v4998
	v5005 = F_list_make1_impl(m, int32(1), v22+int32(60))
	mBase = m.M
	v5006 = m.ExcPending
	if v5006 != 0 {
		goto L1
	} else {
		goto L1320
	}
L1320:
	;
	F_ReportApplyConflict(m, v748, v746, int32(15), v4992, v632, v772, v5005)
	mBase = m.M
	v5008 = m.ExcPending
	if v5008 != 0 {
		goto L1
	} else {
		goto L1321
	}
L1321:
	;
	goto L7
L1322:
	;
	F_EvalPlanQualEnd(m, v22+int32(336))
	mBase = m.M
	v5016 = m.ExcPending
	if v5016 != 0 {
		goto L1
	} else {
		goto L1323
	}
L1323:
	;
	goto L6
L1324:
	;
	v5026 = *(*int32)(unsafe.Add(mBase, uint32(v626)+16))
	if v5026 != 0 {
		goto L1325
	} else {
		goto L1326
	}
L1325:
	;
	v5027 = *(*int32)(unsafe.Add(mBase, uint32(v626)+12))
	F_ExecCleanupTupleRouting(m, v5027, v5026)
	mBase = m.M
	v5029 = m.ExcPending
	if v5029 != 0 {
		goto L1
	} else {
		goto L1328
	}
L1326:
	;
	goto L1327
L1327:
	;
	F_ExecCloseTrigTargetRelations(m, v5023)
	mBase = m.M
	v5031 = m.ExcPending
	if v5031 != 0 {
		goto L1
	} else {
		goto L1329
	}
L1328:
	;
	goto L1327
L1329:
	;
	v5032 = int32(0)
	v5033 = *(*int32)(unsafe.Add(mBase, uint32(v5023)+104))
	F_ExecResetTupleTable(m, v5033, v5032)
	mBase = m.M
	v5036 = m.ExcPending
	if v5036 != 0 {
		goto L1
	} else {
		goto L1330
	}
L1330:
	;
	F_FreeExecutorState(m, v5023)
	mBase = m.M
	v5038 = m.ExcPending
	if v5038 != 0 {
		goto L1
	} else {
		goto L1331
	}
L1331:
	;
	F_pfree(m, v626)
	mBase = m.M
	v5040 = m.ExcPending
	if v5040 != 0 {
		goto L1
	} else {
		goto L1332
	}
L1332:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = int32(0)
	if v616 != 0 {
		v5049 = v5032
		goto L5
	} else {
		goto L1333
	}
L1333:
	;
	F_RestoreUserContext(m, v22+int32(320))
	mBase = m.M
	v5047 = m.ExcPending
	if v5047 != 0 {
		goto L1
	} else {
		goto L1334
	}
L1334:
	;
	v5049 = v5032
	goto L5
L1335:
	;
	goto L4
L1336:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v5091 = m.ExcPending
	if v5091 != 0 {
		goto L1
	} else {
		goto L1337
	}
L1337:
	;
	goto L3
}
