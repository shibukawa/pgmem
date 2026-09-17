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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
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
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
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
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v164 int64
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int64
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
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
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
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
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v486 int32
	_ = v486
	var v488 int64
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int64
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
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
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int64
	_ = v742
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v840 int64
	_ = v840
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v851 int64
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1009 int64
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1020 int64
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1109 int32
	_ = v1109
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
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
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
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
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
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
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1409 int32
	_ = v1409
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1554 int32
	_ = v1554
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1608 int32
	_ = v1608
	var v1609 int64
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1626 int32
	_ = v1626
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1651 int32
	_ = v1651
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1736 int32
	_ = v1736
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1742 int64
	_ = v1742
	var v1748 int32
	_ = v1748
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1771 int64
	_ = v1771
	var v1777 int32
	_ = v1777
	var v1787 int32
	_ = v1787
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1813 int32
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1827 int32
	_ = v1827
	var v1843 int32
	_ = v1843
	var v1846 int64
	_ = v1846
	var v1859 int32
	_ = v1859
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1885 int32
	_ = v1885
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1927 int32
	_ = v1927
	var v1931 int32
	_ = v1931
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1962 int32
	_ = v1962
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2009 int32
	_ = v2009
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2018 int32
	_ = v2018
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2030 int32
	_ = v2030
	var v2033 int32
	_ = v2033
	var v2036 int32
	_ = v2036
	var v2039 int32
	_ = v2039
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2056 int32
	_ = v2056
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2061 int32
	_ = v2061
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
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
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2114 int32
	_ = v2114
	var v2119 int32
	_ = v2119
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2136 int32
	_ = v2136
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2169 int32
	_ = v2169
	var v2174 int32
	_ = v2174
	var v2176 int32
	_ = v2176
	var v2178 int32
	_ = v2178
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2187 int32
	_ = v2187
	var v2193 int32
	_ = v2193
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2204 int32
	_ = v2204
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2234 int64
	_ = v2234
	var v2236 int64
	_ = v2236
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2253 int64
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2256 int64
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2259 int64
	_ = v2259
	var v2264 int32
	_ = v2264
	var v2267 int64
	_ = v2267
	var v2269 int32
	_ = v2269
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2275 int32
	_ = v2275
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2290 int32
	_ = v2290
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2300 int32
	_ = v2300
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2328 int32
	_ = v2328
	var v2332 int64
	_ = v2332
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2348 int32
	_ = v2348
	var v2351 int32
	_ = v2351
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2357 int32
	_ = v2357
	var v2359 int32
	_ = v2359
	var v2361 int64
	_ = v2361
	var v2378 int64
	_ = v2378
	var v2383 int64
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2388 int32
	_ = v2388
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2396 int32
	_ = v2396
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2411 int32
	_ = v2411
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2416 int64
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2419 int32
	_ = v2419
	var v2421 int32
	_ = v2421
	var v2423 int32
	_ = v2423
	var v2425 int64
	_ = v2425
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2432 int32
	_ = v2432
	var v2444 int32
	_ = v2444
	var v2455 int32
	_ = v2455
	var v2459 int32
	_ = v2459
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2470 int32
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2483 int32
	_ = v2483
	var v2487 int32
	_ = v2487
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2496 int64
	_ = v2496
	var v2498 int32
	_ = v2498
	var v2501 int32
	_ = v2501
	var v2506 int32
	_ = v2506
	var v2516 int64
	_ = v2516
	var v2519 int32
	_ = v2519
	var v2521 int64
	_ = v2521
	var v2523 int64
	_ = v2523
	var v2527 int32
	_ = v2527
	var v2528 int64
	_ = v2528
	var v2550 int32
	_ = v2550
	var v2552 int32
	_ = v2552
	var v2558 int32
	_ = v2558
	var v2563 int32
	_ = v2563
	var v2567 int32
	_ = v2567
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2574 int32
	_ = v2574
	var v2576 int32
	_ = v2576
	var v2580 int32
	_ = v2580
	var v2586 int32
	_ = v2586
	var v2591 int32
	_ = v2591
	var v2593 int32
	_ = v2593
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2600 int32
	_ = v2600
	var v2602 int32
	_ = v2602
	var v2604 int64
	_ = v2604
	var v2628 int32
	_ = v2628
	var v2630 int32
	_ = v2630
	var v2632 int32
	_ = v2632
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2660 int32
	_ = v2660
	var v2665 int32
	_ = v2665
	var v2668 int32
	_ = v2668
	var v2672 int32
	_ = v2672
	var v2677 int32
	_ = v2677
	var v2679 int32
	_ = v2679
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2684 int32
	_ = v2684
	var v2686 int64
	_ = v2686
	var v2689 int64
	_ = v2689
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2701 int32
	_ = v2701
	var v2707 int32
	_ = v2707
	var v2708 int32
	_ = v2708
	var v2712 int32
	_ = v2712
	var v2714 int32
	_ = v2714
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2721 int32
	_ = v2721
	var v2722 int32
	_ = v2722
	var v2724 int32
	_ = v2724
	var v2726 int32
	_ = v2726
	var v2732 int32
	_ = v2732
	var v2733 int32
	_ = v2733
	var v2737 int32
	_ = v2737
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2746 int32
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2751 int32
	_ = v2751
	var v2756 int32
	_ = v2756
	var v2759 int32
	_ = v2759
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2781 int32
	_ = v2781
	var v2784 int32
	_ = v2784
	var v2788 int32
	_ = v2788
	var v2793 int32
	_ = v2793
	var v2795 int32
	_ = v2795
	var v2796 int32
	_ = v2796
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2806 int32
	_ = v2806
	var v2809 int32
	_ = v2809
	var v2834 int32
	_ = v2834
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2844 int32
	_ = v2844
	var v2849 int32
	_ = v2849
	var v2853 int32
	_ = v2853
	var v2860 int32
	_ = v2860
	var v2865 int32
	_ = v2865
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2876 int32
	_ = v2876
	var v2877 int32
	_ = v2877
	var v2878 int32
	_ = v2878
	var v2881 int64
	_ = v2881
	var v2883 int64
	_ = v2883
	var v2886 int32
	_ = v2886
	var v2888 int32
	_ = v2888
	var v2889 int32
	_ = v2889
	var v2891 int32
	_ = v2891
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2897 int32
	_ = v2897
	var v2899 int32
	_ = v2899
	var v2903 int32
	_ = v2903
	var v2907 int32
	_ = v2907
	var v2912 int32
	_ = v2912
	var v2913 int64
	_ = v2913
	var v2914 int32
	_ = v2914
	var v2916 int64
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2919 int64
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2928 int64
	_ = v2928
	var v2931 int32
	_ = v2931
	var v2932 int32
	_ = v2932
	var v2935 int32
	_ = v2935
	var v2938 int32
	_ = v2938
	var v2939 int32
	_ = v2939
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2944 int32
	_ = v2944
	var v2947 int64
	_ = v2947
	var v2949 int32
	_ = v2949
	var v2951 int32
	_ = v2951
	var v2953 int32
	_ = v2953
	var v2954 int32
	_ = v2954
	var v2955 int64
	_ = v2955
	var v2957 int32
	_ = v2957
	var v2961 int32
	_ = v2961
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2966 int32
	_ = v2966
	var v2969 int32
	_ = v2969
	var v2970 int32
	_ = v2970
	var v2976 int32
	_ = v2976
	var v2981 int32
	_ = v2981
	var v2984 int32
	_ = v2984
	var v2989 int32
	_ = v2989
	var v2990 int32
	_ = v2990
	var v2992 int32
	_ = v2992
	var v2993 int64
	_ = v2993
	var v2995 int32
	_ = v2995
	var v2997 int32
	_ = v2997
	var v2999 int32
	_ = v2999
	var v3006 int32
	_ = v3006
	var v3008 int32
	_ = v3008
	var v3010 int64
	_ = v3010
	var v3014 int32
	_ = v3014
	var v3016 int32
	_ = v3016
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3029 int32
	_ = v3029
	var v3034 int32
	_ = v3034
	var v3037 int64
	_ = v3037
	var v3039 int32
	_ = v3039
	var v3041 int32
	_ = v3041
	var v3059 int32
	_ = v3059
	var v3060 int32
	_ = v3060
	var v3063 int32
	_ = v3063
	var v3067 int32
	_ = v3067
	var v3068 int64
	_ = v3068
	var v3069 int32
	_ = v3069
	var v3073 int64
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3078 int64
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3082 int32
	_ = v3082
	var v3083 int32
	_ = v3083
	var v3086 int32
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3088 int32
	_ = v3088
	var v3095 int32
	_ = v3095
	var v3099 int32
	_ = v3099
	var v3111 int32
	_ = v3111
	var v3112 int32
	_ = v3112
	var v3113 int32
	_ = v3113
	var v3115 int32
	_ = v3115
	var v3119 int32
	_ = v3119
	var v3120 int32
	_ = v3120
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3124 int32
	_ = v3124
	var v3126 int32
	_ = v3126
	var v3132 int32
	_ = v3132
	var v3133 int32
	_ = v3133
	var v3134 int32
	_ = v3134
	var v3135 int32
	_ = v3135
	var v3138 int32
	_ = v3138
	var v3145 int32
	_ = v3145
	var v3146 int32
	_ = v3146
	var v3147 int32
	_ = v3147
	var v3150 int32
	_ = v3150
	var v3153 int32
	_ = v3153
	var v3158 int32
	_ = v3158
	var v3159 int32
	_ = v3159
	var v3161 int32
	_ = v3161
	var v3163 int32
	_ = v3163
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3174 int32
	_ = v3174
	var v3175 int32
	_ = v3175
	var v3176 int32
	_ = v3176
	var v3179 int32
	_ = v3179
	var v3180 int32
	_ = v3180
	var v3181 int32
	_ = v3181
	var v3183 int32
	_ = v3183
	var v3187 int32
	_ = v3187
	var v3188 int32
	_ = v3188
	var v3190 int32
	_ = v3190
	var v3192 int32
	_ = v3192
	var v3194 int32
	_ = v3194
	var v3195 int32
	_ = v3195
	var v3198 int32
	_ = v3198
	var v3205 int32
	_ = v3205
	var v3211 int32
	_ = v3211
	var v3215 int32
	_ = v3215
	var v3220 int32
	_ = v3220
	var v3224 int32
	_ = v3224
	var v3228 int32
	_ = v3228
	var v3233 int32
	_ = v3233
	var v3235 int32
	_ = v3235
	var v3238 int64
	_ = v3238
	var v3243 int32
	_ = v3243
	var v3244 int64
	_ = v3244
	var v3253 int32
	_ = v3253
	var v3254 int32
	_ = v3254
	var v3258 int64
	_ = v3258
	var v3261 int64
	_ = v3261
	var v3267 int32
	_ = v3267
	var v3272 int32
	_ = v3272
	var v3275 int32
	_ = v3275
	var v3284 int32
	_ = v3284
	var v3285 int64
	_ = v3285
	var v3287 int64
	_ = v3287
	var v3290 int32
	_ = v3290
	var v3294 int64
	_ = v3294
	var v3297 int32
	_ = v3297
	var v3298 int32
	_ = v3298
	var v3304 int32
	_ = v3304
	var v3306 int32
	_ = v3306
	var v3307 int32
	_ = v3307
	var v3308 int32
	_ = v3308
	var v3310 int32
	_ = v3310
	var v3313 int32
	_ = v3313
	var v3316 int32
	_ = v3316
	var v3317 int32
	_ = v3317
	var v3318 int32
	_ = v3318
	var v3322 int32
	_ = v3322
	var v3324 int32
	_ = v3324
	var v3325 int32
	_ = v3325
	var v3331 int32
	_ = v3331
	var v3333 int32
	_ = v3333
	var v3335 int64
	_ = v3335
	var v3338 int64
	_ = v3338
	var v3342 int32
	_ = v3342
	var v3343 int32
	_ = v3343
	var v3345 int32
	_ = v3345
	var v3347 int32
	_ = v3347
	var v3349 int32
	_ = v3349
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3353 int64
	_ = v3353
	var v3355 int32
	_ = v3355
	var v3356 int32
	_ = v3356
	var v3359 int32
	_ = v3359
	var v3364 int32
	_ = v3364
	var v3367 int32
	_ = v3367
	var v3368 int32
	_ = v3368
	var v3373 int32
	_ = v3373
	var v3375 int32
	_ = v3375
	var v3377 int32
	_ = v3377
	var v3380 int32
	_ = v3380
	var v3382 int32
	_ = v3382
	var v3389 int32
	_ = v3389
	var v3391 int64
	_ = v3391
	var v3394 int64
	_ = v3394
	var v3396 int32
	_ = v3396
	var v3399 int32
	_ = v3399
	var v3401 int64
	_ = v3401
	var v3406 int32
	_ = v3406
	var v3407 int32
	_ = v3407
	var v3409 int64
	_ = v3409
	var v3412 int64
	_ = v3412
	var v3418 int32
	_ = v3418
	var v3423 int32
	_ = v3423
	var v3429 int64
	_ = v3429
	var v3431 int32
	_ = v3431
	var v3433 int32
	_ = v3433
	var v3451 int32
	_ = v3451
	var v3452 int32
	_ = v3452
	var v3454 int32
	_ = v3454
	var v3456 int32
	_ = v3456
	var v3457 int32
	_ = v3457
	var v3459 int32
	_ = v3459
	var v3462 int64
	_ = v3462
	var v3463 int32
	_ = v3463
	var v3467 int64
	_ = v3467
	var v3468 int32
	_ = v3468
	var v3472 int64
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3476 int32
	_ = v3476
	var v3477 int32
	_ = v3477
	var v3480 int32
	_ = v3480
	var v3481 int32
	_ = v3481
	var v3482 int32
	_ = v3482
	var v3489 int32
	_ = v3489
	var v3493 int32
	_ = v3493
	var v3505 int32
	_ = v3505
	var v3506 int32
	_ = v3506
	var v3507 int32
	_ = v3507
	var v3509 int32
	_ = v3509
	var v3513 int32
	_ = v3513
	var v3514 int32
	_ = v3514
	var v3516 int32
	_ = v3516
	var v3517 int32
	_ = v3517
	var v3518 int32
	_ = v3518
	var v3520 int32
	_ = v3520
	var v3526 int32
	_ = v3526
	var v3527 int32
	_ = v3527
	var v3528 int32
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3532 int32
	_ = v3532
	var v3539 int32
	_ = v3539
	var v3540 int32
	_ = v3540
	var v3541 int32
	_ = v3541
	var v3544 int32
	_ = v3544
	var v3547 int32
	_ = v3547
	var v3552 int32
	_ = v3552
	var v3553 int32
	_ = v3553
	var v3555 int32
	_ = v3555
	var v3557 int32
	_ = v3557
	var v3561 int32
	_ = v3561
	var v3562 int32
	_ = v3562
	var v3563 int32
	_ = v3563
	var v3568 int32
	_ = v3568
	var v3569 int32
	_ = v3569
	var v3570 int32
	_ = v3570
	var v3573 int32
	_ = v3573
	var v3574 int32
	_ = v3574
	var v3575 int32
	_ = v3575
	var v3577 int32
	_ = v3577
	var v3581 int32
	_ = v3581
	var v3582 int32
	_ = v3582
	var v3584 int32
	_ = v3584
	var v3586 int32
	_ = v3586
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3592 int32
	_ = v3592
	var v3599 int32
	_ = v3599
	var v3608 int32
	_ = v3608
	var v3612 int32
	_ = v3612
	var v3617 int32
	_ = v3617
	var v3621 int32
	_ = v3621
	var v3625 int32
	_ = v3625
	var v3630 int32
	_ = v3630
	var v3634 int32
	_ = v3634
	var v3638 int32
	_ = v3638
	var v3643 int32
	_ = v3643
	var v3645 int32
	_ = v3645
	var v3648 int64
	_ = v3648
	var v3651 int32
	_ = v3651
	var v3652 int32
	_ = v3652
	var v3656 int32
	_ = v3656
	var v3658 int32
	_ = v3658
	var v3662 int64
	_ = v3662
	var v3665 int32
	_ = v3665
	var v3666 int32
	_ = v3666
	var v3672 int32
	_ = v3672
	var v3674 int32
	_ = v3674
	var v3675 int32
	_ = v3675
	var v3676 int32
	_ = v3676
	var v3678 int32
	_ = v3678
	var v3681 int32
	_ = v3681
	var v3684 int64
	_ = v3684
	var v3687 int64
	_ = v3687
	var v3693 int32
	_ = v3693
	var v3695 int32
	_ = v3695
	var v3697 int32
	_ = v3697
	var v3699 int32
	_ = v3699
	var v3701 int32
	_ = v3701
	var v3702 int32
	_ = v3702
	var v3704 int64
	_ = v3704
	var v3705 int64
	_ = v3705
	var v3707 int32
	_ = v3707
	var v3708 int32
	_ = v3708
	var v3711 int32
	_ = v3711
	var v3716 int32
	_ = v3716
	var v3719 int32
	_ = v3719
	var v3720 int32
	_ = v3720
	var v3724 int32
	_ = v3724
	var v3726 int32
	_ = v3726
	var v3728 int32
	_ = v3728
	var v3731 int32
	_ = v3731
	var v3733 int32
	_ = v3733
	var v3740 int32
	_ = v3740
	var v3742 int64
	_ = v3742
	var v3745 int64
	_ = v3745
	var v3747 int32
	_ = v3747
	var v3750 int32
	_ = v3750
	var v3751 int64
	_ = v3751
	var v3753 int32
	_ = v3753
	var v3755 int32
	_ = v3755
	var v3773 int32
	_ = v3773
	var v3774 int32
	_ = v3774
	var v3776 int32
	_ = v3776
	var v3778 int32
	_ = v3778
	var v3779 int32
	_ = v3779
	var v3781 int32
	_ = v3781
	var v3784 int64
	_ = v3784
	var v3785 int32
	_ = v3785
	var v3789 int64
	_ = v3789
	var v3790 int32
	_ = v3790
	var v3794 int64
	_ = v3794
	var v3795 int32
	_ = v3795
	var v3797 int64
	_ = v3797
	var v3798 int32
	_ = v3798
	var v3801 int32
	_ = v3801
	var v3802 int32
	_ = v3802
	var v3805 int32
	_ = v3805
	var v3806 int32
	_ = v3806
	var v3807 int32
	_ = v3807
	var v3814 int32
	_ = v3814
	var v3818 int32
	_ = v3818
	var v3830 int32
	_ = v3830
	var v3831 int32
	_ = v3831
	var v3832 int32
	_ = v3832
	var v3834 int32
	_ = v3834
	var v3838 int32
	_ = v3838
	var v3839 int32
	_ = v3839
	var v3841 int32
	_ = v3841
	var v3842 int32
	_ = v3842
	var v3843 int32
	_ = v3843
	var v3845 int32
	_ = v3845
	var v3851 int32
	_ = v3851
	var v3852 int32
	_ = v3852
	var v3853 int32
	_ = v3853
	var v3854 int32
	_ = v3854
	var v3857 int32
	_ = v3857
	var v3864 int32
	_ = v3864
	var v3865 int32
	_ = v3865
	var v3866 int32
	_ = v3866
	var v3869 int32
	_ = v3869
	var v3872 int32
	_ = v3872
	var v3877 int32
	_ = v3877
	var v3878 int32
	_ = v3878
	var v3880 int32
	_ = v3880
	var v3882 int32
	_ = v3882
	var v3886 int32
	_ = v3886
	var v3887 int32
	_ = v3887
	var v3888 int32
	_ = v3888
	var v3893 int32
	_ = v3893
	var v3894 int32
	_ = v3894
	var v3895 int32
	_ = v3895
	var v3898 int32
	_ = v3898
	var v3899 int32
	_ = v3899
	var v3900 int32
	_ = v3900
	var v3902 int32
	_ = v3902
	var v3906 int32
	_ = v3906
	var v3907 int32
	_ = v3907
	var v3909 int32
	_ = v3909
	var v3911 int32
	_ = v3911
	var v3913 int32
	_ = v3913
	var v3914 int32
	_ = v3914
	var v3917 int32
	_ = v3917
	var v3924 int32
	_ = v3924
	var v3933 int32
	_ = v3933
	var v3937 int32
	_ = v3937
	var v3942 int32
	_ = v3942
	var v3946 int32
	_ = v3946
	var v3950 int32
	_ = v3950
	var v3955 int32
	_ = v3955
	var v3959 int32
	_ = v3959
	var v3963 int32
	_ = v3963
	var v3968 int32
	_ = v3968
	var v3970 int32
	_ = v3970
	var v3973 int64
	_ = v3973
	var v3976 int32
	_ = v3976
	var v3977 int32
	_ = v3977
	var v3979 int32
	_ = v3979
	var v3981 int32
	_ = v3981
	var v3982 int64
	_ = v3982
	var v3983 int64
	_ = v3983
	var v3984 int32
	_ = v3984
	var v3985 int32
	_ = v3985
	var v3987 int32
	_ = v3987
	var v3990 int32
	_ = v3990
	var v3994 int32
	_ = v3994
	var v3995 int32
	_ = v3995
	var v3997 int32
	_ = v3997
	var v3998 int32
	_ = v3998
	var v4001 int32
	_ = v4001
	var v4002 int32
	_ = v4002
	var v4022 int32
	_ = v4022
	var v4023 int32
	_ = v4023
	var v4027 int32
	_ = v4027
	var v4030 int32
	_ = v4030
	var v4033 int32
	_ = v4033
	var v4036 int32
	_ = v4036
	var v4037 int32
	_ = v4037
	var v4040 int32
	_ = v4040
	var v4041 int32
	_ = v4041
	var v4044 int32
	_ = v4044
	var v4051 int32
	_ = v4051
	var v4052 int32
	_ = v4052
	var v4054 int32
	_ = v4054
	var v4057 int32
	_ = v4057
	var v4059 int32
	_ = v4059
	var v4060 int32
	_ = v4060
	var v4061 int64
	_ = v4061
	var v4066 int32
	_ = v4066
	var v4067 int32
	_ = v4067
	var v4068 int32
	_ = v4068
	var v4069 int64
	_ = v4069
	var v4071 int64
	_ = v4071
	var v4074 int32
	_ = v4074
	var v4076 int32
	_ = v4076
	var v4077 int32
	_ = v4077
	var v4079 int32
	_ = v4079
	var v4080 int32
	_ = v4080
	var v4084 int32
	_ = v4084
	var v4087 int32
	_ = v4087
	var v4105 int32
	_ = v4105
	var v4109 int32
	_ = v4109
	var v4114 int64
	_ = v4114
	var v4117 int64
	_ = v4117
	var v4120 int32
	_ = v4120
	var v4124 int64
	_ = v4124
	var v4127 int32
	_ = v4127
	var v4128 int32
	_ = v4128
	var v4134 int32
	_ = v4134
	var v4136 int32
	_ = v4136
	var v4137 int32
	_ = v4137
	var v4138 int32
	_ = v4138
	var v4140 int32
	_ = v4140
	var v4143 int32
	_ = v4143
	var v4149 int32
	_ = v4149
	var v4151 int32
	_ = v4151
	var v4153 int32
	_ = v4153
	var v4155 int32
	_ = v4155
	var v4156 int64
	_ = v4156
	var v4158 int32
	_ = v4158
	var v4160 int32
	_ = v4160
	var v4161 int32
	_ = v4161
	var v4162 int64
	_ = v4162
	var v4164 int32
	_ = v4164
	var v4165 int32
	_ = v4165
	var v4168 int32
	_ = v4168
	var v4173 int32
	_ = v4173
	var v4176 int32
	_ = v4176
	var v4177 int32
	_ = v4177
	var v4182 int32
	_ = v4182
	var v4184 int32
	_ = v4184
	var v4186 int32
	_ = v4186
	var v4189 int32
	_ = v4189
	var v4191 int32
	_ = v4191
	var v4198 int32
	_ = v4198
	var v4200 int64
	_ = v4200
	var v4203 int64
	_ = v4203
	var v4205 int32
	_ = v4205
	var v4208 int32
	_ = v4208
	var v4210 int32
	_ = v4210
	var v4227 int64
	_ = v4227
	var v4229 int64
	_ = v4229
	var v4232 int32
	_ = v4232
	var v4234 int32
	_ = v4234
	var v4235 int32
	_ = v4235
	var v4238 int32
	_ = v4238
	var v4245 int32
	_ = v4245
	var v4247 int32
	_ = v4247
	var v4250 int64
	_ = v4250
	var v4253 int32
	_ = v4253
	var v4254 int32
	_ = v4254
	var v4257 int32
	_ = v4257
	var v4260 int32
	_ = v4260
	var v4261 int32
	_ = v4261
	var v4262 int32
	_ = v4262
	var v4263 int32
	_ = v4263
	var v4264 int32
	_ = v4264
	var v4265 int32
	_ = v4265
	var v4266 int32
	_ = v4266
	var v4269 int64
	_ = v4269
	var v4271 int32
	_ = v4271
	var v4273 int32
	_ = v4273
	var v4275 int32
	_ = v4275
	var v4276 int32
	_ = v4276
	var v4277 int32
	_ = v4277
	var v4278 int64
	_ = v4278
	var v4280 int32
	_ = v4280
	var v4284 int32
	_ = v4284
	var v4286 int32
	_ = v4286
	var v4287 int64
	_ = v4287
	var v4290 int32
	_ = v4290
	var v4292 int32
	_ = v4292
	var v4295 int32
	_ = v4295
	var v4296 int32
	_ = v4296
	var v4297 int32
	_ = v4297
	var v4299 int32
	_ = v4299
	var v4302 int32
	_ = v4302
	var v4303 int32
	_ = v4303
	var v4309 int32
	_ = v4309
	var v4314 int32
	_ = v4314
	var v4317 int32
	_ = v4317
	var v4318 int32
	_ = v4318
	var v4323 int32
	_ = v4323
	var v4324 int32
	_ = v4324
	var v4326 int32
	_ = v4326
	var v4327 int64
	_ = v4327
	var v4329 int32
	_ = v4329
	var v4331 int32
	_ = v4331
	var v4333 int32
	_ = v4333
	var v4338 int32
	_ = v4338
	var v4342 int64
	_ = v4342
	var v4345 int32
	_ = v4345
	var v4346 int32
	_ = v4346
	var v4352 int32
	_ = v4352
	var v4354 int32
	_ = v4354
	var v4355 int32
	_ = v4355
	var v4356 int32
	_ = v4356
	var v4358 int32
	_ = v4358
	var v4361 int32
	_ = v4361
	var v4366 int32
	_ = v4366
	var v4368 int32
	_ = v4368
	var v4370 int32
	_ = v4370
	var v4372 int32
	_ = v4372
	var v4374 int32
	_ = v4374
	var v4379 int32
	_ = v4379
	var v4381 int32
	_ = v4381
	var v4382 int32
	_ = v4382
	var v4384 int32
	_ = v4384
	var v4390 int32
	_ = v4390
	var v4391 int32
	_ = v4391
	var v4397 int32
	_ = v4397
	var v4402 int32
	_ = v4402
	var v4406 int32
	_ = v4406
	var v4407 int32
	_ = v4407
	var v4408 int64
	_ = v4408
	var v4410 int32
	_ = v4410
	var v4412 int64
	_ = v4412
	var v4417 int32
	_ = v4417
	var v4418 int32
	_ = v4418
	var v4420 int64
	_ = v4420
	var v4423 int64
	_ = v4423
	var v4429 int32
	_ = v4429
	var v4434 int32
	_ = v4434
	var v4440 int64
	_ = v4440
	var v4442 int32
	_ = v4442
	var v4444 int32
	_ = v4444
	var v4464 int32
	_ = v4464
	var v4467 int32
	_ = v4467
	var v4468 int64
	_ = v4468
	var v4471 int64
	_ = v4471
	var v4473 int64
	_ = v4473
	var v4474 int64
	_ = v4474
	var v4477 int64
	_ = v4477
	var v4483 int32
	_ = v4483
	var v4488 int32
	_ = v4488
	var v4492 int32
	_ = v4492
	var v4495 int32
	_ = v4495
	var v4499 int32
	_ = v4499
	var v4504 int32
	_ = v4504
	var v4508 int32
	_ = v4508
	var v4511 int32
	_ = v4511
	var v4515 int32
	_ = v4515
	var v4520 int32
	_ = v4520
	var v4524 int32
	_ = v4524
	var v4531 int32
	_ = v4531
	var v4536 int32
	_ = v4536
	var v4540 int32
	_ = v4540
	var v4543 int32
	_ = v4543
	var v4547 int32
	_ = v4547
	var v4552 int32
	_ = v4552
	var v4556 int32
	_ = v4556
	var v4563 int32
	_ = v4563
	var v4568 int32
	_ = v4568
	var v4572 int32
	_ = v4572
	var v4575 int32
	_ = v4575
	var v4579 int32
	_ = v4579
	var v4584 int32
	_ = v4584
	var v4588 int32
	_ = v4588
	var v4591 int32
	_ = v4591
	var v4595 int32
	_ = v4595
	var v4600 int32
	_ = v4600
	var v4604 int32
	_ = v4604
	var v4611 int32
	_ = v4611
	var v4616 int32
	_ = v4616
	var v4620 int32
	_ = v4620
	var v4623 int32
	_ = v4623
	var v4627 int32
	_ = v4627
	var v4632 int32
	_ = v4632
	var v4636 int32
	_ = v4636
	var v4639 int32
	_ = v4639
	var v4640 int64
	_ = v4640
	var v4643 int64
	_ = v4643
	var v4645 int64
	_ = v4645
	var v4646 int64
	_ = v4646
	var v4649 int64
	_ = v4649
	var v4655 int32
	_ = v4655
	var v4660 int32
	_ = v4660
	var v4664 int32
	_ = v4664
	var v4667 int32
	_ = v4667
	var v4671 int32
	_ = v4671
	var v4676 int32
	_ = v4676
	var v4680 int32
	_ = v4680
	var v4683 int32
	_ = v4683
	var v4687 int32
	_ = v4687
	var v4692 int32
	_ = v4692
	var v4696 int32
	_ = v4696
	var v4703 int32
	_ = v4703
	var v4708 int32
	_ = v4708
	var v4712 int32
	_ = v4712
	var v4715 int32
	_ = v4715
	var v4719 int32
	_ = v4719
	var v4724 int32
	_ = v4724
	var v4728 int32
	_ = v4728
	var v4734 int32
	_ = v4734
	var v4736 int32
	_ = v4736
	var v4738 int32
	_ = v4738
	var v4741 int32
	_ = v4741
	var v4767 int32
	_ = v4767
	var v4778 int32
	_ = v4778
	var v4779 int32
	_ = v4779
	var v4781 int32
	_ = v4781
	var v4782 int32
	_ = v4782
	var v4783 int32
	_ = v4783
	var v4794 int32
	_ = v4794
	var v4796 int32
	_ = v4796
	var v4797 int32
	_ = v4797
	var v4803 int32
	_ = v4803
	var v4806 int32
	_ = v4806
	var v4807 int32
	_ = v4807
	var v4810 int32
	_ = v4810
	var v4828 int32
	_ = v4828
	var v4832 int32
	_ = v4832
	var v4835 int32
	_ = v4835
	var v4837 int32
	_ = v4837
	var v4838 int32
	_ = v4838
	var v4860 int32
	_ = v4860
	var v4864 int32
	_ = v4864
	var v4882 int32
	_ = v4882
	var v4886 int32
	_ = v4886
	var v4889 int32
	_ = v4889
	var v4891 int32
	_ = v4891
	var v4892 int32
	_ = v4892
	var v4897 int32
	_ = v4897
	var v4899 int32
	_ = v4899
	var v4907 int32
	_ = v4907
	var v4908 int32
	_ = v4908
	var v4910 int32
	_ = v4910
	var v4914 int32
	_ = v4914
	var v4918 int32
	_ = v4918
	var v4925 int32
	_ = v4925
	var v4927 int32
	_ = v4927
	var v4928 int32
	_ = v4928
	var v4929 int32
	_ = v4929
	var v4931 int32
	_ = v4931
	var v4932 int32
	_ = v4932
	var v4933 int32
	_ = v4933
	var v4936 int32
	_ = v4936
	var v4938 int32
	_ = v4938
	var v4940 int32
	_ = v4940
	var v4947 int32
	_ = v4947
	var v4949 int32
	_ = v4949
	var v4967 int32
	_ = v4967
	var v4987 int32
	_ = v4987
	var v4989 int32
	_ = v4989
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(1552)
	m.G0 = v21
	v23 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = int32(_a_F_apply_dispatch_0)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v23
	switch v23 - int32(65) {
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
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v26
	m.G0 = v21 + int32(1552)
	return
L4:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v4987 = m.ExcPending
	if v4987 != 0 {
		goto L1
	} else {
		goto L1306
	}
L5:
	;
	F_logicalrep_rel_close(m, v600, v4949)
	mBase = m.M
	v4967 = m.ExcPending
	if v4967 != 0 {
		goto L1
	} else {
		goto L1305
	}
L6:
	;
	v4925 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	F_AfterTriggerEndQuery(m, v4925)
	mBase = m.M
	v4927 = m.ExcPending
	if v4927 != 0 {
		goto L1
	} else {
		goto L1295
	}
L7:
	;
	F_ExecCloseIndices(m, v739)
	mBase = m.M
	v4914 = m.ExcPending
	if v4914 != 0 {
		goto L1
	} else {
		goto L1293
	}
L8:
	;
	F_slot_store_data(m, v765, v738, v21+int32(276))
	mBase = m.M
	v4897 = m.ExcPending
	if v4897 != 0 {
		goto L1
	} else {
		goto L1290
	}
L9:
	;
	v4794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1344)))
	v4796 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v4797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4796)+31)))
	F_ExecuteTruncateGuts(m, v4782, v4778, v4783, int32(0), v4794, (v4797^int32(-1))&int32(1))
	mBase = m.M
	v4803 = m.ExcPending
	if v4803 != 0 {
		goto L1
	} else {
		goto L1276
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	v4767 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v4767
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v4767
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v4767
	goto L3
L11:
	;
	F_stream_open_and_write_change(m, v2269, int32(65), v21+int32(304))
	mBase = m.M
	v4734 = m.ExcPending
	if v4734 != 0 {
		goto L1
	} else {
		goto L1272
	}
L12:
	;
	F_pa_switch_to_partial_serialize(m, v2278, int32(1))
	mBase = m.M
	v4728 = m.ExcPending
	if v4728 != 0 {
		goto L1
	} else {
		goto L1271
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4712 = m.ExcPending
	if v4712 != 0 {
		goto L1
	} else {
		goto L1267
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4696 = m.ExcPending
	if v4696 != 0 {
		goto L1
	} else {
		goto L1264
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4680 = m.ExcPending
	if v4680 != 0 {
		goto L1
	} else {
		goto L1260
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4664 = m.ExcPending
	if v4664 != 0 {
		goto L1
	} else {
		goto L1256
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4636 = m.ExcPending
	if v4636 != 0 {
		goto L1
	} else {
		goto L1252
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4620 = m.ExcPending
	if v4620 != 0 {
		goto L1
	} else {
		goto L1248
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4604 = m.ExcPending
	if v4604 != 0 {
		goto L1
	} else {
		goto L1245
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4588 = m.ExcPending
	if v4588 != 0 {
		goto L1
	} else {
		goto L1241
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4572 = m.ExcPending
	if v4572 != 0 {
		goto L1
	} else {
		goto L1237
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4556 = m.ExcPending
	if v4556 != 0 {
		goto L1
	} else {
		goto L1234
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4540 = m.ExcPending
	if v4540 != 0 {
		goto L1
	} else {
		goto L1230
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4524 = m.ExcPending
	if v4524 != 0 {
		goto L1
	} else {
		goto L1227
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4508 = m.ExcPending
	if v4508 != 0 {
		goto L1
	} else {
		goto L1223
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4492 = m.ExcPending
	if v4492 != 0 {
		goto L1
	} else {
		goto L1219
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4464 = m.ExcPending
	if v4464 != 0 {
		goto L1
	} else {
		goto L1215
	}
L28:
	;
	v4227 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+1352)) = v4227
	v4229 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+1344)) = v4229
	v4232 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v4232 != 0 {
		goto L16
	} else {
		goto L1140
	}
L29:
	;
	v3773 = v21 + int32(320)
	v3774 = m.G0
	v3776 = v3774 - int32(16)
	m.G0 = v3776
	v3778 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v3779 = m.ExcPending
	if v3779 != 0 {
		goto L1
	} else {
		goto L1025
	}
L30:
	;
	v3451 = v21 + int32(320)
	v3452 = m.G0
	v3454 = v3452 - int32(16)
	m.G0 = v3454
	v3456 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v3457 = m.ExcPending
	if v3457 != 0 {
		goto L1
	} else {
		goto L941
	}
L31:
	;
	F_logicalrep_read_prepare_common(m, l0, int32(_a_F_apply_dispatch_1), v21+int32(320))
	mBase = m.M
	v3284 = m.ExcPending
	if v3284 != 0 {
		goto L1
	} else {
		goto L891
	}
L32:
	;
	v3059 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v3060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3059)+16)))
	if v3060 == int32(1) {
		goto L833
	} else {
		goto L834
	}
L33:
	;
	v2881 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+1352)) = v2881
	v2883 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+1344)) = v2883
	v2886 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v2886 != 0 {
		goto L20
	} else {
		goto L780
	}
L34:
	;
	v2234 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+312)) = v2234
	v2236 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+304)) = v2236
	v2239 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v2239 != 0 {
		goto L21
	} else {
		goto L614
	}
L35:
	;
	v2071 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v2071 == int32(0) {
		goto L23
	} else {
		goto L571
	}
L36:
	;
	v1561 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v1561 != 0 {
		goto L26
	} else {
		goto L467
	}
L37:
	;
	v1523 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v1523 != 0 {
		goto L3
	} else {
		goto L456
	}
L38:
	;
	v1501 = F_handle_streamed_transaction(m, int32(89), l0)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L1
	} else {
		goto L446
	}
L39:
	;
	v1325 = F_handle_streamed_transaction(m, int32(82), l0)
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L1
	} else {
		goto L402
	}
L40:
	;
	v1004 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+320)) = uint8(v1004)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1344)) = uint8(v1004)
	v1009 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	if v1009 != int64(0) {
		goto L3
	} else {
		goto L317
	}
L41:
	;
	v840 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	if v840 != int64(0) {
		goto L3
	} else {
		goto L258
	}
L42:
	;
	v488 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	if v488 != int64(0) {
		goto L3
	} else {
		goto L159
	}
L43:
	;
	v164 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	if v164 != int64(0) {
		goto L3
	} else {
		goto L74
	}
L44:
	;
	v104 = m.G0
	v106 = v104 - int32(16)
	m.G0 = v106
	v108 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L61
	}
L45:
	;
	v32 = v21 + int32(320)
	v33 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32))) = v33
	if v33 == int64(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v51 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_2), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(68), int32(_a_F_apply_dispatch_4))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v51
	v55 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v55
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v21)+336))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v59
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v62
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[9])) = v62
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	if base.B2i32(v68 == int64(0))|base.B2i32(v62 != v68) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v99 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])) = uint8(v99)
	F_pgstat_report_activity(m, int32(3), int32(0))
	mBase = m.M
	goto L3
L56:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8])) = v62
	v77 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v77 == int32(0) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v82 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+20)) = uint32(v82)
	v85 = int64(base.Ui64(v82) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+16)) = uint32(v85)
	F_errmsg(m, int32(_a_F_apply_dispatch_5), v21+int32(16))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(_a_F_apply_dispatch_7), int32(_a_F_apply_dispatch_8))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L55
L61:
	;
	v111 = v108 & int32(255)
	if v111 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v126 = v21 + int32(320)
	v127 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L68
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v111
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_9), v106)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(104), int32(_a_F_apply_dispatch_10))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v126))) = v127
	v130 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = v130
	v133 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v126)+16)) = v133
	m.G0 = v106 + int32(16)
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	v141 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[9]))
	if v139 != v141 {
		goto L27
	} else {
		goto L71
	}
L71:
	;
	F_apply_handle_commit_internal(m, v126)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	F_process_syncing_tables(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v149 = int32(0)
	F_pgstat_report_activity(m, int32(2), v149)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v149
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v149
	goto L3
L74:
	;
	v168 = F_handle_streamed_transaction(m, int32(73), l0)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if v168 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v171 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+20))
	goto L81
L78:
	;
	v175 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v175
	goto L80
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	if base.B2i32(v179 == int32(2)) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v188 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	F_PushActiveSnapshot(m, v188)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v194
	v199 = m.G0
	v201 = v199 - int32(16)
	m.G0 = v201
	v204 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v206 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v209 = v206 << (uint(int32(24)) % 32)
	if v209 != int32(1308622848) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	F_logicalrep_read_tuple(m, l0, v21+int32(320))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L97
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v209 >> (uint(int32(24)) % 32)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_11), v201)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(439), int32(_a_F_apply_dispatch_12))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
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
	m.G0 = v201 + int32(16)
	v233 = F_logicalrep_rel_open(m, v204, int32(3))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L99
	}
L98:
	;
	F_logicalrep_rel_close(m, v233, v468)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L158
	}
L99:
	;
	v235 = F_should_apply_changes_for_rel(m, v233)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	if v235 == int32(0) {
		v468 = int32(3)
		goto L98
	} else {
		goto L101
	}
L101:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+31)))
	if v241 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v233)+40))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+48))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+80))
	F_SwitchToUntrustedUser(m, v246, v21+int32(1344))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v233
	v253 = F_create_edata_for_relation(m, v233)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v233)+40))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+52))
	v259 = F_ExecInitExtraTupleSlot(m, v255, v257, int32(_a_F_apply_dispatch_13))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v255)+152))
	if v261 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v264 = F_MakePerTupleExprContext(m, v255)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	v266 = v261
	goto L110
L110:
	;
	v267 = int32(_a_F_apply_dispatch_14)
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15]))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v270
	F_slot_store_data(m, v259, v233, v21+int32(320))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	v266 = v264
	goto L110
L112:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v233)+40))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+52))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v255)+152))
	if v279 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v282 = F_MakePerTupleExprContext(m, v255)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L116
	}
L114:
	;
	v284 = v279
	goto L115
L115:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	if v278 == v285 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v284 = v282
	goto L115
L117:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v268
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v233)+40))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+48))
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+119)))
	if v420 == int32(112) {
		goto L139
	} else {
		goto L140
	}
L118:
	;
	v288 = v278 << (uint(int32(2)) % 32)
	v289 = F_palloc(m, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v291 = F_palloc(m, v288)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	if v278 <= int32(0) {
		goto L117
	} else {
		goto L121
	}
L121:
	;
	v296 = int32(0)
	v303 = v2
	goto L122
L122:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v320 = v277 + v314<<(uint(int32(4))%32) + v296*int32(100)
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+111)))
	if v321 != 0 {
		v351 = v303
		goto L124
	} else {
		goto L125
	}
L123:
	;
	if v351 <= int32(0) {
		goto L117
	} else {
		goto L133
	}
L124:
	;
	v355 = v296 + int32(1)
	if v355 != v278 {
		v296 = v355
		v303 = v351
		goto L122
	} else {
		goto L132
	}
L125:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+110)))
	if v322 != 0 {
		v351 = v303
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v233)+44))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	v328 = int32(*(*int16)(unsafe.Add(mBase, uint32(v324+v296<<(uint(int32(1))%32)))))
	if int32(0) <= v328 {
		v351 = v303
		goto L124
	} else {
		goto L127
	}
L127:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v233)+40))
	v334 = F_build_column_default(m, v331, v296+int32(1))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	if v334 == int32(0) {
		v351 = v303
		goto L124
	} else {
		goto L129
	}
L129:
	;
	v339 = v303 << (uint(int32(2)) % 32)
	v341 = F_expression_planner(m, v334)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v344 = F_ExecInitExpr(m, v341, int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291+v339))) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v289+v339))) = v296
	v351 = v303 + int32(1)
	goto L124
L132:
	;
	goto L123
L133:
	;
	v360 = int32(0)
	goto L134
L134:
	;
	v379 = v360 << (uint(int32(2)) % 32)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v291+v379)))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v259)+20))
	v383 = v379 + v289
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v381)+20))
	v387 = m.T0[v386].(func(*base.Module, int32, int32, int32) int32)(m, v381, v284, v382+v384)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L136
	}
L135:
	;
	goto L117
L136:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	*(*int32)(unsafe.Add(mBase, uint32(v389+v390<<(uint(int32(2))%32)))) = v387
	v396 = v360 + int32(1)
	if v396 != v351 {
		v360 = v396
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	F_AfterTriggerEndQuery(m, v444)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L148
	}
L139:
	;
	F_apply_handle_tuple_routing(m, v253, v259, int32(0), int32(3))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v253)+8))
	F_ExecOpenIndices(m, v427, int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L143
	}
L142:
	;
	goto L138
L143:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	F_InitConflictIndexes(m, v427)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v427)+8))
	F_TargetPrivilegesCheck(m, v434, int64(1))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_ExecSimpleRelationInsert(m, v427, v431, v259)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_ExecCloseIndices(m, v427)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	goto L138
L148:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v253)+16))
	if v447 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	F_ExecCleanupTupleRouting(m, v448, v447)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v451 = int32(0)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v444)+104))
	F_ExecResetTupleTable(m, v452, v451)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L153
	}
L152:
	;
	goto L151
L153:
	;
	F_FreeExecutorState(m, v444)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_pfree(m, v253)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = int32(0)
	if v241 != 0 {
		v468 = v451
		goto L98
	} else {
		goto L156
	}
L156:
	;
	F_RestoreUserContext(m, v21+int32(1344))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v468 = v451
	goto L98
L158:
	;
	goto L4
L159:
	;
	v492 = F_handle_streamed_transaction(m, int32(85), l0)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	if v492 != 0 {
		goto L3
	} else {
		goto L161
	}
L161:
	;
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v495 < int32(0) {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v502 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)+20))
	goto L166
L163:
	;
	v499 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v499
	goto L165
L164:
	;
	goto L165
L165:
	;
	goto L162
L166:
	;
	if base.B2i32(v503 == int32(2)) == int32(0) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v512 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L172
	}
L170:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	F_PushActiveSnapshot(m, v512)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v518
	v522 = v21 + int32(275)
	v527 = m.G0
	v529 = v527 - int32(32)
	m.G0 = v529
	v532 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L175
	}
L174:
	;
	v600 = F_logicalrep_rel_open(m, v532, int32(3))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L197
	}
L175:
	;
	v534 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L178
	}
L176:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L194
	}
L177:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L191
	}
L178:
	;
	if v534&int32(251) != int32(75) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v541 = v534 << (uint(int32(24)) % 32)
	if v541 != int32(1308622848) {
		goto L177
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v545 = int32(75)
	if v534&v545 == v545 {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	goto L181
L183:
	;
	v559 = v557 << (uint(int32(24)) % 32)
	if v559 != int32(1308622848) {
		goto L176
	} else {
		goto L189
	}
L184:
	;
	F_logicalrep_read_tuple(m, l0, v21+int32(288))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v555 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v522))) = uint8(v555)
	v557 = v534
	goto L183
L187:
	;
	v551 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v522))) = uint8(v551)
	v553 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	v557 = v553
	goto L183
L189:
	;
	F_logicalrep_read_tuple(m, l0, v21+int32(276))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	m.G0 = v529 + int32(32)
	goto L174
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v529))) = v541 >> (uint(int32(24)) % 32)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_15), v529)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(501), int32(_a_F_apply_dispatch_16))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v529)+16)) = v559 >> (uint(int32(24)) % 32)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_17), v529+int32(16))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(517), int32(_a_F_apply_dispatch_16))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	v602 = F_should_apply_changes_for_rel(m, v600)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	if v602 == int32(0) {
		v4949 = int32(3)
		goto L5
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v600
	F_check_relation_updatable(m, v600)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v611 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+31)))
	if v612 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v600)+40))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v615)+48))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v616)+80))
	F_SwitchToUntrustedUser(m, v617, v21+int32(304))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v622 = F_create_edata_for_relation(m, v600)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v600)+40))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v625)+52))
	v628 = F_ExecInitExtraTupleSlot(m, v624, v626, int32(_a_F_apply_dispatch_13))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v628)+12))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	if int32(0) < v631 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v624)+32))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)+12))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	v638 = int32(0)
	v639 = v631
	v641 = v630
	goto L210
L208:
	;
	goto L209
L209:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v624)+152))
	if v708 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L210:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641+v639<<(uint(int32(4))%32)+v638*int32(100))+111)))
	if v662 != 0 {
		v684 = v639
		v685 = v641
		goto L212
	} else {
		goto L213
	}
L211:
	;
	goto L209
L212:
	;
	v688 = v638 + int32(1)
	if v688 < v684 {
		v638 = v688
		v639 = v684
		v641 = v685
		goto L210
	} else {
		goto L217
	}
L213:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v600)+44))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	v668 = int32(*(*int16)(unsafe.Add(mBase, uint32(v664+v638<<(uint(int32(1))%32)))))
	if v668 < int32(0) {
		v684 = v639
		v685 = v641
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v21)+280))
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671+v668))))
	if v673 == int32(117) {
		v684 = v639
		v685 = v641
		goto L212
	} else {
		goto L215
	}
L215:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v636)+36))
	v679 = F_bms_add_member(m, v676, v638+int32(8))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v636)+36)) = v679
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v628)+12))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v682)))
	v684 = v683
	v685 = v682
	goto L212
L217:
	;
	goto L211
L218:
	;
	v711 = F_MakePerTupleExprContext(m, v624)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L221
	}
L219:
	;
	v713 = v708
	goto L220
L220:
	;
	v714 = int32(_a_F_apply_dispatch_14)
	v715 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15]))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v713)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v717
	v722 = v21 + int32(276)
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+275)))
	if v723 != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v713 = v711
	goto L220
L222:
	;
	v724 = v21 + int32(288)
	goto L224
L223:
	;
	v724 = v722
	goto L224
L224:
	;
	F_slot_store_data(m, v628, v600, v724)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v715
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v600)+40))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v729)+48))
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730)+119)))
	if v731 == int32(112) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	F_apply_handle_tuple_routing(m, v622, v628, v722, int32(2))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v600)+52))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v622)+4))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v622)+8))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v739)+8))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	v742 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+1360)) = v742
	*(*int64)(unsafe.Add(mBase, uint32(v21)+1352)) = v742
	*(*int64)(unsafe.Add(mBase, uint32(v21)+1344)) = v742
	v750 = int32(0)
	F_EvalPlanQualInit(m, v21+int32(320), v741, v750, v750, int32(-1), v750)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L230
	}
L229:
	;
	goto L6
L230:
	;
	F_ExecOpenIndices(m, v739, int32(0))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	F_TargetPrivilegesCheck(m, v740, int64(2))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v765 = F_table_slot_create(m, v740, v759+int32(104))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	if v737 != 0 {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v779 = F_GetTupleTransactionInfo(m, v765, v21+int32(1352), v21+int32(1356), v21+int32(1360))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L243
	}
L235:
	;
	v767 = F_RelationFindReplTupleByIndex(m, v740, v737, v628, v765)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v769 = F_RelationFindReplTupleSeq(m, v740, v628, v765)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L240
	}
L238:
	;
	if v767 != 0 {
		goto L234
	} else {
		goto L239
	}
L239:
	;
	goto L8
L240:
	;
	if v769 == int32(0) {
		goto L8
	} else {
		goto L241
	}
L241:
	;
	goto L234
L242:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v741)+152))
	if v811 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L243:
	;
	if v779 == int32(0) {
		goto L242
	} else {
		goto L244
	}
L244:
	;
	v783 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+1356)))
	v785 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_apply_dispatch[16])))
	if v783 == v785 {
		goto L242
	} else {
		goto L245
	}
L245:
	;
	v789 = F_table_slot_create(m, v740, v741+int32(104))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	F_slot_store_data(m, v789, v738, v21+int32(276))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+1344)) = v765
	v797 = v21 + int32(1344)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+300)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v797
	v801 = int32(1)
	v805 = F_list_make1_impl(m, v801, v21+int32(56))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	F_ReportApplyConflict(m, v741, v739, int32(15), v801, v628, v789, v805)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	goto L242
L250:
	;
	v814 = F_MakePerTupleExprContext(m, v741)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L253
	}
L251:
	;
	v816 = v811
	goto L252
L252:
	;
	v817 = int32(_a_F_apply_dispatch_14)
	v818 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15]))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v816)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v820
	F_slot_modify_data(m, v628, v765, v738, v21+int32(276))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L254
	}
L253:
	;
	v816 = v814
	goto L252
L254:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v21)+348)) = v628
	F_InitConflictIndexes(m, v739)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v739)+8))
	F_TargetPrivilegesCheck(m, v831, int64(4))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	F_ExecSimpleRelationUpdate(m, v739, v741, v21+int32(320), v765, v628)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	goto L7
L258:
	;
	v844 = F_handle_streamed_transaction(m, int32(68), l0)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	if v844 != 0 {
		goto L3
	} else {
		goto L260
	}
L260:
	;
	v847 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v847 < int32(0) {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v854 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v854)+20))
	goto L265
L262:
	;
	v851 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v851
	goto L264
L263:
	;
	goto L264
L264:
	;
	goto L261
L265:
	;
	if base.B2i32(v855 == int32(2)) == int32(0) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L1
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	v864 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L271
	}
L269:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	goto L268
L271:
	;
	F_PushActiveSnapshot(m, v864)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	v870 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v870
	v875 = m.G0
	v877 = v875 - int32(16)
	m.G0 = v877
	v880 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v882 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	if v882&int32(251) != int32(75) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L1
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	F_logicalrep_read_tuple(m, l0, v21+int32(320))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L1
	} else {
		goto L281
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v877))) = base.I32_extend8_s(v882)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_18), v877)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(572), int32(_a_F_apply_dispatch_19))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	m.G0 = v877 + int32(16)
	v908 = F_logicalrep_rel_open(m, v880, int32(3))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L1
	} else {
		goto L283
	}
L282:
	;
	F_logicalrep_rel_close(m, v908, v998)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L1
	} else {
		goto L316
	}
L283:
	;
	v910 = F_should_apply_changes_for_rel(m, v908)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	if v910 == int32(0) {
		v998 = int32(3)
		goto L282
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v908
	F_check_relation_updatable(m, v908)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	v919 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919)+31)))
	if v920 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v908)+40))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v923)+48))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v924)+80))
	F_SwitchToUntrustedUser(m, v925, v21+int32(1344))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	v930 = F_create_edata_for_relation(m, v908)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L291
	}
L290:
	;
	goto L289
L291:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v930)))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v908)+40))
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v933)+52))
	v936 = F_ExecInitExtraTupleSlot(m, v932, v934, int32(_a_F_apply_dispatch_13))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v932)+152))
	if v938 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v941 = F_MakePerTupleExprContext(m, v932)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L296
	}
L294:
	;
	v943 = v938
	goto L295
L295:
	;
	v944 = int32(_a_F_apply_dispatch_14)
	v945 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15]))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v943)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v947
	F_slot_store_data(m, v936, v908, v21+int32(320))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L297
	}
L296:
	;
	v943 = v941
	goto L295
L297:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v945
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v908)+40))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v955)+48))
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956)+119)))
	if v957 == int32(112) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v930)))
	F_AfterTriggerEndQuery(m, v974)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L1
	} else {
		goto L306
	}
L299:
	;
	F_apply_handle_tuple_routing(m, v930, v936, int32(0), int32(4))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v930)+8))
	F_ExecOpenIndices(m, v964, int32(0))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L303
	}
L302:
	;
	goto L298
L303:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v908)+52))
	F_apply_handle_delete_internal(m, v930, v964, v936, v968)
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	F_ExecCloseIndices(m, v964)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	goto L298
L306:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v930)+16))
	if v977 != 0 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v930)+12))
	F_ExecCleanupTupleRouting(m, v978, v977)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	v981 = int32(0)
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v974)+104))
	F_ExecResetTupleTable(m, v982, v981)
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L1
	} else {
		goto L311
	}
L310:
	;
	goto L309
L311:
	;
	F_FreeExecutorState(m, v974)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	F_pfree(m, v930)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = int32(0)
	if v920 != 0 {
		v998 = v981
		goto L282
	} else {
		goto L314
	}
L314:
	;
	F_RestoreUserContext(m, v21+int32(1344))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	v998 = v981
	goto L282
L316:
	;
	goto L4
L317:
	;
	v1013 = F_handle_streamed_transaction(m, int32(84), l0)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	if v1013 != 0 {
		goto L3
	} else {
		goto L319
	}
L319:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v1016 < int32(0) {
		goto L321
	} else {
		goto L322
	}
L320:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+20))
	goto L324
L321:
	;
	v1020 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v1020
	goto L323
L322:
	;
	goto L323
L323:
	;
	goto L320
L324:
	;
	if base.B2i32(v1024 == int32(2)) == int32(0) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L1
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	v1033 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L330
	}
L328:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	goto L327
L330:
	;
	F_PushActiveSnapshot(m, v1033)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v1039
	v1046 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	v1049 = F_pq_getmsgint(m, l0, int32(1))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	v1051 = int32(1)
	v1052 = v1049 & v1051
	*(*uint8)(unsafe.Add(mBase, uint32(v21+int32(320)))) = uint8(v1052)
	v1057 = int32(base.Ui32(v1049)>>(uint(v1051)%32)) & v1051
	*(*uint8)(unsafe.Add(mBase, uint32(v21+int32(1344)))) = uint8(v1057)
	v1059 = int32(0)
	if v1059 < v1046 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1064 = v1059
	v1068 = int32(0)
	goto L337
L335:
	;
	v1090 = v1059
	goto L336
L336:
	;
	if v1090 == int32(0) {
		v4778 = v2
		v4779 = v2
		v4781 = v2
		v4782 = v2
		v4783 = v2
		goto L9
	} else {
		goto L342
	}
L337:
	;
	v1082 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L339
	}
L338:
	;
	v1090 = v1084
	goto L336
L339:
	;
	v1084 = F_lappend_oid(m, v1064, v1082)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	v1087 = v1068 + int32(1)
	if v1087 != v1046 {
		v1064 = v1084
		v1068 = v1087
		goto L337
	} else {
		goto L341
	}
L341:
	;
	goto L338
L342:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+4))
	if v1109 <= int32(0) {
		v4778 = v2
		v4779 = v2
		v4781 = v2
		v4782 = v2
		v4783 = v2
		goto L9
	} else {
		goto L343
	}
L343:
	;
	v1115 = v2
	v1116 = v2
	v1118 = v2
	v1119 = v2
	v1120 = v2
	v1123 = v2
	goto L344
L344:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+12))
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1130+v1123<<(uint(int32(2))%32))))
	v1136 = F_logicalrep_rel_open(m, v1134, int32(8))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L1
	} else {
		goto L347
	}
L345:
	;
	v4778 = v1305
	v4779 = v1306
	v4781 = v1308
	v4782 = v1309
	v4783 = v1310
	goto L9
L346:
	;
	v1321 = v1123 + int32(1)
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+4))
	if v1321 < v1322 {
		v1115 = v1305
		v1116 = v1306
		v1118 = v1308
		v1119 = v1309
		v1120 = v1310
		v1123 = v1321
		goto L344
	} else {
		goto L401
	}
L347:
	;
	v1138 = F_should_apply_changes_for_rel(m, v1136)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	if v1138 == int32(0) {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	F_logicalrep_rel_close(m, v1136, int32(8))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L1
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1145 = F_lappend(m, v1118, v1136)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L353
	}
L352:
	;
	v1305 = v1115
	v1306 = v1116
	v1308 = v1118
	v1309 = v1119
	v1310 = v1120
	goto L346
L353:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+40))
	F_TargetPrivilegesCheck(m, v1147, int64(16))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+40))
	v1152 = F_lappend(m, v1119, v1151)
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+36))
	v1155 = F_lappend_oid(m, v1115, v1154)
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[17]))
	if v1158 < int32(2) {
		v1176 = v1120
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+40))
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1178)+48))
	v1180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1179)+119)))
	if v1180 != int32(112) {
		v1305 = v1155
		v1306 = v1116
		v1308 = v1145
		v1309 = v1152
		v1310 = v1176
		goto L346
	} else {
		goto L364
	}
L358:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+40))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1161)+48))
	v1163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1162)+118)))
	if v1163 != int32(112) {
		v1176 = v1120
		goto L357
	} else {
		goto L359
	}
L359:
	;
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1162)+119)))
	if v1166 == int32(102) {
		v1176 = v1120
		goto L357
	} else {
		goto L360
	}
L360:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1161)+56))
	goto L361
L361:
	;
	if base.Ui32(v1169) < base.Ui32(int32(_a_F_apply_dispatch_20)) {
		v1176 = v1120
		goto L357
	} else {
		goto L362
	}
L362:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+36))
	v1173 = F_lappend_oid(m, v1120, v1172)
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	v1176 = v1173
	goto L357
L364:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+36))
	v1186 = F_find_all_inheritors(m, v1183, int32(8), int32(0))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	if v1186 == int32(0) {
		v1305 = v1155
		v1306 = v1116
		v1308 = v1145
		v1309 = v1152
		v1310 = v1176
		goto L346
	} else {
		goto L366
	}
L366:
	;
	v1190 = int32(0)
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+4))
	if v1191 <= v1190 {
		v1305 = v1155
		v1306 = v1116
		v1308 = v1145
		v1309 = v1152
		v1310 = v1176
		goto L346
	} else {
		goto L367
	}
L367:
	;
	v1194 = v1190
	v1197 = v1155
	v1198 = v1116
	v1201 = v1152
	v1202 = v1176
	goto L368
L368:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+12))
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1212+v1194<<(uint(int32(2))%32))))
	v1217 = int32(0)
	if v1197 == v1217 {
		goto L372
	} else {
		goto L373
	}
L369:
	;
	v1305 = v1292
	v1306 = v1293
	v1308 = v1145
	v1309 = v1295
	v1310 = v1296
	goto L346
L370:
	;
	v1299 = v1194 + int32(1)
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+4))
	if v1299 < v1300 {
		v1194 = v1299
		v1197 = v1292
		v1198 = v1293
		v1201 = v1295
		v1202 = v1296
		goto L368
	} else {
		goto L400
	}
L371:
	;
	if v1255 != 0 {
		v1292 = v1197
		v1293 = v1198
		v1295 = v1201
		v1296 = v1202
		goto L370
	} else {
		goto L384
	}
L372:
	;
	v1255 = int32(0)
	goto L371
L373:
	;
	goto L374
L374:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+4))
	if v1223 <= int32(0) {
		v1249 = v1217
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v1255 = v1249
	goto L371
L376:
	;
	v1226 = int32(0)
	if v1226 < v1223 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1229 = v1223
	goto L379
L378:
	;
	v1229 = v1226
	goto L379
L379:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+12))
	v1232 = int32(0)
	goto L380
L380:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1230+v1232<<(uint(int32(2))%32))))
	v1241 = base.B2i32(v1240 == v1216)
	if v1240 == v1216 {
		v1249 = v1241
		goto L375
	} else {
		goto L382
	}
L381:
	;
	v1249 = v1241
	goto L375
L382:
	;
	v1243 = v1232 + int32(1)
	if v1243 != v1229 {
		v1232 = v1243
		goto L380
	} else {
		goto L383
	}
L383:
	;
	goto L381
L384:
	;
	v1257 = F_table_open(m, v1216, int32(0))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L1
	} else {
		goto L386
	}
L385:
	;
	F_TargetPrivilegesCheck(m, v1257, int64(16))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L1
	} else {
		goto L390
	}
L386:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+48))
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1259)+118)))
	if v1260 != int32(116) {
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v1263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1257)+24)))
	if v1263 != 0 {
		goto L385
	} else {
		goto L388
	}
L388:
	;
	F_relation_close(m, v1257, int32(8))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	v1292 = v1197
	v1293 = v1198
	v1295 = v1201
	v1296 = v1202
	goto L370
L390:
	;
	v1270 = F_lappend(m, v1201, v1257)
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	v1272 = F_lappend(m, v1198, v1257)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	v1274 = F_lappend_oid(m, v1197, v1216)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[17]))
	if v1277 < int32(2) {
		v1292 = v1274
		v1293 = v1272
		v1295 = v1270
		v1296 = v1202
		goto L370
	} else {
		goto L394
	}
L394:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+48))
	v1281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280)+118)))
	if v1281 != int32(112) {
		v1292 = v1274
		v1293 = v1272
		v1295 = v1270
		v1296 = v1202
		goto L370
	} else {
		goto L395
	}
L395:
	;
	v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280)+119)))
	if v1284 == int32(102) {
		v1292 = v1274
		v1293 = v1272
		v1295 = v1270
		v1296 = v1202
		goto L370
	} else {
		goto L396
	}
L396:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+56))
	goto L397
L397:
	;
	if base.Ui32(v1287) < base.Ui32(int32(_a_F_apply_dispatch_20)) {
		v1292 = v1274
		v1293 = v1272
		v1295 = v1270
		v1296 = v1202
		goto L370
	} else {
		goto L398
	}
L398:
	;
	v1290 = F_lappend_oid(m, v1202, v1216)
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	v1292 = v1274
	v1293 = v1272
	v1295 = v1270
	v1296 = v1290
	goto L370
L400:
	;
	goto L369
L401:
	;
	goto L345
L402:
	;
	if v1325 != 0 {
		goto L3
	} else {
		goto L403
	}
L403:
	;
	v1328 = F_palloc(m, int32(32))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	v1331 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1328))) = v1331
	v1334 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334))))
	if v1337 != 0 {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v1338 = v1334
	goto L409
L408:
	;
	v1338 = int32(_a_F_apply_dispatch_21)
	goto L409
L409:
	;
	v1339 = F_pstrdup(m, v1338)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+4)) = v1339
	v1342 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	v1344 = F_pstrdup(m, v1342)
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+8)) = v1344
	v1347 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1328)+24)) = uint8(v1347)
	v1351 = F_pq_getmsgint(m, l0, int32(2))
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	v1354 = v1351 << (uint(int32(2)) % 32)
	v1355 = F_palloc(m, v1354)
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	v1357 = F_palloc(m, v1354)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	if int32(0) < v1351 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v1365 = v2
	v1367 = int32(0)
	goto L420
L418:
	;
	v1409 = v2
	goto L419
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+28)) = v1409
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+20)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+16)) = v1355
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+12)) = v1351
	F_logicalrep_relmap_update(m, v1328)
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L1
	} else {
		goto L432
	}
L420:
	;
	v1380 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L1
	} else {
		goto L422
	}
L421:
	;
	v1409 = v1386
	goto L419
L422:
	;
	if v1380&int32(1) != 0 {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v1384 = F_bms_add_member(m, v1365, v1367)
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L1
	} else {
		goto L426
	}
L424:
	;
	v1386 = v1365
	goto L425
L425:
	;
	v1388 = v1367 << (uint(int32(2)) % 32)
	v1390 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L1
	} else {
		goto L427
	}
L426:
	;
	v1386 = v1384
	goto L425
L427:
	;
	v1392 = F_pstrdup(m, v1390)
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1355+v1388))) = v1392
	v1397 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L1
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1357+v1388))) = v1397
	v1401 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	v1404 = v1367 + int32(1)
	if v1404 != v1351 {
		v1365 = v1386
		v1367 = v1404
		goto L420
	} else {
		goto L431
	}
L431:
	;
	goto L421
L432:
	;
	v1430 = m.G0
	v1432 = v1430 - int32(32)
	m.G0 = v1432
	v1435 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[18]))
	if v1435 == int32(0) {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	m.G0 = v1432 + int32(32)
	goto L3
L434:
	;
	v1439 = v1432 + int32(12)
	F_hash_seq_init(m, v1439, v1435)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L1
	} else {
		goto L435
	}
L435:
	;
	v1442 = F_hash_seq_search(m, v1439)
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L1
	} else {
		goto L436
	}
L436:
	;
	if v1442 == int32(0) {
		goto L433
	} else {
		goto L437
	}
L437:
	;
	v1447 = v1442
	goto L438
L438:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+8))
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1328)))
	if v1464 == v1465 {
		goto L440
	} else {
		goto L441
	}
L439:
	;
	goto L433
L440:
	;
	v1468 = v1447 + int32(8)
	F_logicalrep_relmap_free_entry(m, v1468)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L1
	} else {
		goto L443
	}
L441:
	;
	goto L442
L442:
	;
	v1477 = F_hash_seq_search(m, v1432+int32(12))
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L1
	} else {
		goto L444
	}
L443:
	;
	base.MemoryFill(m, v1468, int32(0), int32(72))
	goto L442
L444:
	;
	if v1477 != 0 {
		v1447 = v1477
		goto L438
	} else {
		goto L445
	}
L445:
	;
	goto L439
L446:
	;
	if v1501 != 0 {
		goto L3
	} else {
		goto L447
	}
L447:
	;
	v1504 = v21 + int32(320)
	v1506 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1504))) = v1506
	v1509 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L1
	} else {
		goto L449
	}
L449:
	;
	v1512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1509))))
	if v1512 != 0 {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v1513 = v1509
	goto L452
L451:
	;
	v1513 = int32(_a_F_apply_dispatch_21)
	goto L452
L452:
	;
	v1514 = F_pstrdup(m, v1513)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1504)+4)) = v1514
	v1517 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L1
	} else {
		goto L454
	}
L454:
	;
	v1519 = F_pstrdup(m, v1517)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1504)+8)) = v1519
	goto L3
L456:
	;
	v1525 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])))
	if v1525 != int32(1) {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L1
	} else {
		goto L463
	}
L458:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v1529)+20))
	goto L459
L459:
	;
	if base.B2i32(v1530 == int32(2)) == int32(0) {
		goto L3
	} else {
		goto L460
	}
L460:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v1537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1536)+16)))
	if v1537 != int32(1) {
		goto L457
	} else {
		goto L461
	}
L461:
	;
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v1536)))
	if v1540 == int32(1) {
		goto L3
	} else {
		goto L462
	}
L462:
	;
	goto L457
L463:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_22), int32(0))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1436), int32(_a_F_apply_dispatch_23))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L1
	} else {
		goto L466
	}
L466:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L467:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1566 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])) = uint8(v1566)
	v1571 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L1
	} else {
		goto L468
	}
L468:
	;
	v1573 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v21+int32(304)))) = uint8(base.B2i32(v1573 == int32(1)))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[19])) = v1571
	if v1571 == int32(0) {
		goto L25
	} else {
		goto L470
	}
L470:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v1571
	v1587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+304)))
	if v1587 == int32(1) {
		goto L471
	} else {
		goto L472
	}
L471:
	;
	v1590 = m.G0
	v1592 = v1590 + int32(-64)
	m.G0 = v1592
	*(*int32)(unsafe.Add(mBase, uint32(v1592)+60)) = v1571
	v1596 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1596)))
	if v1597 != int32(2) {
		goto L476
	} else {
		goto L477
	}
L472:
	;
	v1940 = v1571
	goto L473
L473:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v1959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1958)+16)))
	if v1959 == int32(1) {
		goto L537
	} else {
		goto L538
	}
L474:
	;
	v1938 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[19]))
	v1940 = v1938
	goto L473
L475:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L1
	} else {
		goto L532
	}
L476:
	;
	m.G0 = v1592 - int32(-64)
	goto L474
L477:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L1
	} else {
		goto L478
	}
L478:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v1604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1603)+68)))
	if v1604 != int32(1) {
		goto L476
	} else {
		goto L479
	}
L479:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v1609 = *(*int64)(unsafe.Add(mBase, uint32(v1608)+8))
	if v1609 != int64(0) {
		goto L476
	} else {
		goto L480
	}
L480:
	;
	v1612 = F_AllTablesyncsReady(m)
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L1
	} else {
		goto L481
	}
L481:
	;
	if v1612 == int32(0) {
		goto L476
	} else {
		goto L482
	}
L482:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[20]))
	if v1617 == int32(0) {
		goto L484
	} else {
		goto L485
	}
L483:
	;
	v1843 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[21]))
	if v1843 == int32(0) {
		goto L522
	} else {
		goto L523
	}
L484:
	;
	v1671 = int32(_a_F_apply_dispatch_14)
	v1672 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15]))
	v1675 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v1675
	v1678 = F_palloc0(m, int32(20))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L1
	} else {
		goto L491
	}
L485:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v1617)+4))
	if v1620 <= int32(0) {
		goto L484
	} else {
		goto L486
	}
L486:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1617)+12))
	v1626 = int32(0)
	goto L487
L487:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1623+v1626<<(uint(int32(2))%32))))
	v1647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1646)+13)))
	if v1647 == int32(0) {
		v1827 = v1646
		goto L483
	} else {
		goto L489
	}
L488:
	;
	goto L484
L489:
	;
	v1651 = v1626 + int32(1)
	if v1620 != v1651 {
		v1626 = v1651
		goto L487
	} else {
		goto L490
	}
L490:
	;
	goto L488
L491:
	;
	v1682 = F_add_size(m, int32(0), int32(96))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	v1685 = F_add_size(m, v1682, int32(16777216))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L1
	} else {
		goto L493
	}
L493:
	;
	v1688 = F_add_size(m, v1685, int32(_a_F_apply_dispatch_24))
	mBase = m.M
	v1689 = m.ExcPending
	if v1689 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1592)+8)) = v1688
	v1693 = F_add_size(m, int32(0), int32(3))
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L1
	} else {
		goto L495
	}
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1592)+12)) = v1693
	v1697 = v1590 + int32(-56)
	v1698 = F_shm_toc_estimate(m, v1697)
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	v1700 = F_shm_toc_estimate(m, v1697)
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	v1703 = F_dsm_create(m, v1700, int32(0))
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	if v1703 == int32(0) {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v1672
	F_pfree(m, v1678)
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L1
	} else {
		goto L502
	}
L500:
	;
	goto L501
L501:
	;
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1703)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1712)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1712)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1712))) = int64(2021433447)
	*(*int32)(unsafe.Add(mBase, uint32(v1712)+12)) = v1698 & int32(-32)
	goto L503
L502:
	;
	goto L476
L503:
	;
	v1722 = F_shm_toc_allocate(m, v1712, int32(80))
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		goto L1
	} else {
		goto L504
	}
L504:
	;
	v1724 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+20)) = v1724
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+8)) = v1724
	*(*int32)(unsafe.Add(mBase, uint32(v1722))) = v1724
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+32)) = v1724
	*(*int64)(unsafe.Add(mBase, uint32(v1722)+24)) = int64(0)
	F_shm_toc_insert(m, v1712, int64(1), v1722)
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	v1739 = F_shm_toc_allocate(m, v1712, int32(16777216))
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	v1742 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1739)+16)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v1739)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1739))) = v1742
	v1748 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v1739)+36)) = uint16(v1748)
	*(*int64)(unsafe.Add(mBase, uint32(v1739)+24)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v1739)+32)) = int32(16777176)
	goto L507
L507:
	;
	F_shm_toc_insert(m, v1712, int64(2), v1739)
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L1
	} else {
		goto L508
	}
L508:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[23]))
	F_shm_mq_set_sender(m, v1739, v1760)
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L1
	} else {
		goto L509
	}
L509:
	;
	v1763 = F_shm_mq_attach(m, v1739, v1703)
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L1
	} else {
		goto L510
	}
L510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1678))) = v1763
	v1768 = F_shm_toc_allocate(m, v1712, int32(_a_F_apply_dispatch_24))
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	v1771 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1768)+16)) = v1771
	*(*int32)(unsafe.Add(mBase, uint32(v1768)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1768))) = v1771
	v1777 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v1768)+36)) = uint16(v1777)
	*(*int64)(unsafe.Add(mBase, uint32(v1768)+24)) = v1771
	*(*int32)(unsafe.Add(mBase, uint32(v1768)+32)) = int32(_a_F_apply_dispatch_25)
	goto L512
L512:
	;
	F_shm_toc_insert(m, v1712, int64(3), v1768)
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		goto L1
	} else {
		goto L513
	}
L513:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[23]))
	F_shm_mq_set_receiver(m, v1768, v1789)
	mBase = m.M
	v1791 = m.ExcPending
	if v1791 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	v1792 = F_shm_mq_attach(m, v1768, v1703)
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1678)+16)) = v1722
	*(*int32)(unsafe.Add(mBase, uint32(v1678)+8)) = v1703
	*(*int32)(unsafe.Add(mBase, uint32(v1678)+4)) = v1792
	v1799 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1799)+24))
	v1802 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1802)))
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v1802)+16))
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v1799)+28))
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1703)+12))
	v1808 = F_logicalrep_worker_launch(m, int32(3), v1800, v1803, v1804, v1805, int32(0), v1807)
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	if v1808 == int32(0) {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	F_pa_free_worker_info(m, v1678)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L1
	} else {
		goto L520
	}
L518:
	;
	goto L519
L519:
	;
	v1817 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[20]))
	v1818 = F_lappend(m, v1817, v1678)
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L1
	} else {
		goto L521
	}
L520:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v1672
	goto L476
L521:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v1672
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[20])) = v1818
	v1827 = v1678
	goto L483
L522:
	;
	v1846 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1592)+48)) = v1846
	*(*int64)(unsafe.Add(mBase, uint32(v1592)+40)) = v1846
	*(*int64)(unsafe.Add(mBase, uint32(v1592)+32)) = v1846
	*(*int64)(unsafe.Add(mBase, uint32(v1592)+16)) = v1846
	*(*int64)(unsafe.Add(mBase, uint32(v1592)+8)) = v1846
	*(*int64)(unsafe.Add(mBase, uint32(v1592)+24)) = int64(34359738372)
	v1859 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1592)+48)) = v1859
	v1867 = F_hash_create(m, int32(_a_F_apply_dispatch_26), int32(16), v1590+int32(-56), int32(1064))
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L1
	} else {
		goto L525
	}
L523:
	;
	v1870 = v1843
	goto L524
L524:
	;
	v1876 = F_hash_search(m, v1870, v1590+int32(-4), int32(1), v1590+int32(-56))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L1
	} else {
		goto L526
	}
L525:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[21])) = v1867
	v1870 = v1867
	goto L524
L526:
	;
	v1878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1592)+8)))
	if v1878 == int32(1) {
		goto L475
	} else {
		goto L527
	}
L527:
	;
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1827)+16))
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v1881)))
	*(*int32)(unsafe.Add(mBase, uint32(v1881))) = int32(1)
	if v1882 != 0 {
		goto L528
	} else {
		goto L529
	}
L528:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1827)+16))
	F_s_lock(m, v1885, int32(_a_F_apply_dispatch_27), int32(504), int32(_a_F_apply_dispatch_28))
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L1
	} else {
		goto L531
	}
L529:
	;
	goto L530
L530:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1827)+16))
	v1892 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+8)) = v1892
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1827)+16))
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1592)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+4)) = v1895
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1827)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1897))) = v1892
	v1900 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v1827)+12)) = uint16(v1900)
	*(*int32)(unsafe.Add(mBase, uint32(v1876)+4)) = v1827
	goto L476
L531:
	;
	goto L530
L532:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_29), int32(0))
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L1
	} else {
		goto L533
	}
L533:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_27), int32(501), int32(_a_F_apply_dispatch_28))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L1
	} else {
		goto L534
	}
L534:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L535:
	;
	F_pgstat_report_activity(m, int32(3), int32(0))
	mBase = m.M
	goto L3
L536:
	;
	v2039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+304)))
	if v2039 == int32(1) {
		goto L565
	} else {
		goto L566
	}
L537:
	;
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(v1958)))
	if v1962 == int32(3) {
		goto L536
	} else {
		goto L540
	}
L538:
	;
	goto L539
L539:
	;
	v1965 = F_pa_find_worker(m, v1940)
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		goto L1
	} else {
		goto L543
	}
L540:
	;
	goto L539
L541:
	;
	v2011 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1344)) = uint8(v2011)
	v2013 = v1563 - v1562
	*(*int32)(unsafe.Add(mBase, uint32(v21)+320)) = v2013 + int32(1)
	v2018 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24]))
	F_BufFileWrite(m, v2018, v21+int32(320), int32(4))
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L1
	} else {
		goto L561
	}
L542:
	;
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1986 = F_pa_send_data(m, v1965, v1984, v1985)
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		goto L1
	} else {
		goto L551
	}
L543:
	;
	if v1965 != 0 {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v1967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1965)+12)))
	if v1967 == int32(0) {
		goto L542
	} else {
		goto L547
	}
L545:
	;
	goto L546
L546:
	;
	v1976 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v1976 == int32(0) {
		goto L24
	} else {
		goto L549
	}
L547:
	;
	v1971 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[19]))
	v1972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+304)))
	F_stream_start_internal(m, v1971, v1972)
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	goto L541
L549:
	;
	v1980 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[19]))
	v1981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+304)))
	F_stream_start_internal(m, v1980, v1981)
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	goto L535
L551:
	;
	v1988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+304)))
	if v1986 != 0 {
		goto L552
	} else {
		goto L553
	}
L552:
	;
	if v1988&int32(1) == int32(0) {
		goto L555
	} else {
		goto L556
	}
L553:
	;
	goto L554
L554:
	;
	F_pa_switch_to_partial_serialize(m, v1965, (v1988^int32(-1))&int32(1))
	mBase = m.M
	v2009 = m.ExcPending
	if v2009 != 0 {
		goto L1
	} else {
		goto L560
	}
L555:
	;
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1965)+16))
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1993)+4))
	F_pa_unlock_stream(m, v1994)
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L1
	} else {
		goto L558
	}
L556:
	;
	goto L557
L557:
	;
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1965)+16))
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v1997)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1997)+20)) = v1998 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25])) = v1965
	goto L559
L558:
	;
	goto L557
L559:
	;
	goto L535
L560:
	;
	goto L541
L561:
	;
	v2025 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24]))
	F_BufFileWrite(m, v2025, v21+int32(1344), int32(1))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L1
	} else {
		goto L562
	}
L562:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+320)) = v2013
	v2033 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24]))
	F_BufFileWrite(m, v2033, v1562+v1564, v2013)
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25])) = v1965
	goto L564
L564:
	;
	goto L535
L565:
	;
	v2043 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v2043)+32))
	v2046 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[26]))
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v2046)+4))
	F_LockApplyTransactionForSession(m, v2044, v2047, int32(1), int32(8))
	mBase = m.M
	v2051 = m.ExcPending
	if v2051 != 0 {
		goto L1
	} else {
		goto L568
	}
L566:
	;
	goto L567
L567:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[27])) = int32(0)
	goto L535
L568:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[26]))
	F_pa_set_xact_state(m, v2053, int32(1))
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L1
	} else {
		goto L569
	}
L569:
	;
	v2058 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v2058)+32))
	F_logicalrep_worker_wakeup(m, v2059)
	mBase = m.M
	v2061 = m.ExcPending
	if v2061 != 0 {
		goto L1
	} else {
		goto L570
	}
L570:
	;
	goto L567
L571:
	;
	v2075 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[19]))
	v2077 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2077)+16)))
	if v2078 == int32(1) {
		goto L574
	} else {
		goto L575
	}
L572:
	;
	v2204 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[19])) = v2204
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])) = uint8(v2204)
	v2212 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v2212)+24))
	goto L610
L573:
	;
	v2184 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L1
	} else {
		goto L603
	}
L574:
	;
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(v2077)))
	if v2081 == int32(3) {
		goto L573
	} else {
		goto L577
	}
L575:
	;
	goto L576
L576:
	;
	v2084 = F_pa_find_worker(m, v2075)
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L1
	} else {
		goto L580
	}
L577:
	;
	goto L576
L578:
	;
	v2127 = int32(69)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1344)) = uint8(v2127)
	v2129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+320)) = v2129 - v2130 + int32(1)
	v2136 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24]))
	F_BufFileWrite(m, v2136, v21+int32(320), int32(4))
	mBase = m.M
	v2141 = m.ExcPending
	if v2141 != 0 {
		goto L1
	} else {
		goto L595
	}
L579:
	;
	F_pa_switch_to_partial_serialize(m, v2084, int32(1))
	mBase = m.M
	v2126 = m.ExcPending
	if v2126 != 0 {
		goto L1
	} else {
		goto L594
	}
L580:
	;
	if v2084 != 0 {
		goto L581
	} else {
		goto L582
	}
L581:
	;
	v2086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2084)+12)))
	if v2086 != 0 {
		goto L578
	} else {
		goto L584
	}
L582:
	;
	goto L583
L583:
	;
	v2101 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v2101 == int32(0) {
		goto L22
	} else {
		goto L589
	}
L584:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v2084)+16))
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v2087)+4))
	F_pa_lock_stream(m, v2088)
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L1
	} else {
		goto L585
	}
L585:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2093 = F_pa_send_data(m, v2084, v2091, v2092)
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	if v2093 == int32(0) {
		goto L579
	} else {
		goto L587
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25])) = int32(0)
	goto L588
L588:
	;
	goto L572
L589:
	;
	v2105 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+32))
	v2108 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[19]))
	F_subxact_info_write(m, v2106, v2108)
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L1
	} else {
		goto L590
	}
L590:
	;
	v2112 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24]))
	F_BufFileClose(m, v2112)
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24])) = int32(0)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	v2121 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[28]))
	F_MemoryContextReset(m, v2121)
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	goto L572
L594:
	;
	goto L578
L595:
	;
	v2143 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24]))
	F_BufFileWrite(m, v2143, v21+int32(1344), int32(1))
	mBase = m.M
	v2148 = m.ExcPending
	if v2148 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2151 = v2149 - v2150
	*(*int32)(unsafe.Add(mBase, uint32(v21)+320)) = v2151
	v2154 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24]))
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_BufFileWrite(m, v2154, v2150+v2155, v2151)
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	v2160 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v2160)+32))
	v2163 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[19]))
	F_subxact_info_write(m, v2161, v2163)
	mBase = m.M
	v2165 = m.ExcPending
	if v2165 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	v2167 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24]))
	F_BufFileClose(m, v2167)
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24])) = int32(0)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	v2176 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[28]))
	F_MemoryContextReset(m, v2176)
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L1
	} else {
		goto L601
	}
L601:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[25])) = int32(0)
	goto L602
L602:
	;
	goto L572
L603:
	;
	if v2184 != 0 {
		goto L604
	} else {
		goto L605
	}
L604:
	;
	v2187 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v2187
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_30), v21+int32(80))
	mBase = m.M
	v2193 = m.ExcPending
	if v2193 != 0 {
		goto L1
	} else {
		goto L607
	}
L605:
	;
	goto L606
L606:
	;
	F_pa_decr_and_wait_stream_block(m)
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L1
	} else {
		goto L609
	}
L607:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1693), int32(_a_F_apply_dispatch_31))
	mBase = m.M
	v2198 = m.ExcPending
	if v2198 != 0 {
		goto L1
	} else {
		goto L608
	}
L608:
	;
	goto L606
L609:
	;
	goto L572
L610:
	;
	if v2213 != v2204 {
		goto L611
	} else {
		goto L612
	}
L611:
	;
	v2216 = int32(4)
	goto L613
L612:
	;
	v2216 = int32(2)
	goto L613
L613:
	;
	v2217 = int32(0)
	F_pgstat_report_activity(m, v2216, v2217)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v2217
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v2217
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v2217
	goto L3
L614:
	;
	v2241 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2241)+68)))
	v2244 = v21 + int32(1344)
	v2246 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v2247 = m.ExcPending
	if v2247 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2244))) = v2246
	v2250 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2244)+4)) = v2250
	if v2242 != 0 {
		goto L618
	} else {
		goto L619
	}
L617:
	;
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(v21)+1348))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v2264
	v2267 = *(*int64)(unsafe.Add(mBase, uint32(v21)+1352))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v2267
	v2269 = *(*int32)(unsafe.Add(mBase, uint32(v21)+1344))
	v2271 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2271)+16)))
	if v2272 == int32(1) {
		goto L626
	} else {
		goto L627
	}
L618:
	;
	v2253 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2254 = m.ExcPending
	if v2254 != 0 {
		goto L1
	} else {
		goto L621
	}
L619:
	;
	goto L620
L620:
	;
	v2259 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2244)+8)) = v2259
	*(*int64)(unsafe.Add(mBase, uint32(v2244)+16)) = v2259
	goto L617
L621:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2244)+8)) = v2253
	v2256 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L1
	} else {
		goto L622
	}
L622:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2244)+16)) = v2256
	goto L617
L623:
	;
	F_pa_unlock_stream(m, v2269)
	mBase = m.M
	v2867 = m.ExcPending
	if v2867 != 0 {
		goto L1
	} else {
		goto L776
	}
L624:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2853 = m.ExcPending
	if v2853 != 0 {
		goto L1
	} else {
		goto L773
	}
L625:
	;
	if v2264 != v2269 {
		goto L722
	} else {
		goto L723
	}
L626:
	;
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(v2271)))
	if v2275 == int32(3) {
		goto L625
	} else {
		goto L629
	}
L627:
	;
	goto L628
L628:
	;
	v2278 = F_pa_find_worker(m, v2269)
	mBase = m.M
	v2279 = m.ExcPending
	if v2279 != 0 {
		goto L1
	} else {
		goto L630
	}
L629:
	;
	goto L628
L630:
	;
	if v2278 != 0 {
		goto L631
	} else {
		goto L632
	}
L631:
	;
	v2280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2278)+12)))
	if v2280 != 0 {
		goto L11
	} else {
		goto L634
	}
L632:
	;
	goto L633
L633:
	;
	v2292 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v2292 != 0 {
		goto L624
	} else {
		goto L639
	}
L634:
	;
	if v2264 != v2269 {
		goto L623
	} else {
		goto L635
	}
L635:
	;
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2284 = F_pa_send_data(m, v2278, v2282, v2283)
	mBase = m.M
	v2285 = m.ExcPending
	if v2285 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	if v2284 == int32(0) {
		goto L12
	} else {
		goto L637
	}
L637:
	;
	F_pa_xact_finish(m, v2278, int64(0))
	mBase = m.M
	v2290 = m.ExcPending
	if v2290 != 0 {
		goto L1
	} else {
		goto L638
	}
L638:
	;
	goto L10
L639:
	;
	if v2264 == v2269 {
		goto L641
	} else {
		goto L642
	}
L640:
	;
	v2653 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2654 = m.ExcPending
	if v2654 != 0 {
		goto L1
	} else {
		goto L718
	}
L641:
	;
	v2295 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2295)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+144)) = v2296
	*(*int32)(unsafe.Add(mBase, uint32(v21)+148)) = v2269
	v2300 = v21 + int32(320)
	v2305 = F_pg_snprintf(m, v2300, int32(1024), int32(_a_F_apply_dispatch_32), v21+int32(144))
	mBase = m.M
	v2306 = m.ExcPending
	if v2306 != 0 {
		goto L1
	} else {
		goto L644
	}
L642:
	;
	goto L643
L643:
	;
	v2328 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v2328 < int32(0) {
		goto L649
	} else {
		goto L650
	}
L644:
	;
	v2308 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2309 = *(*int32)(unsafe.Add(mBase, uint32(v2308)+60))
	F_BufFileDeleteFileSet(m, v2309, v2300, int32(0))
	mBase = m.M
	v2312 = m.ExcPending
	if v2312 != 0 {
		goto L1
	} else {
		goto L645
	}
L645:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+132)) = v2269
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v2296
	v2319 = F_pg_snprintf(m, v2300, int32(1024), int32(_a_F_apply_dispatch_33), v21+int32(128))
	mBase = m.M
	v2320 = m.ExcPending
	if v2320 != 0 {
		goto L1
	} else {
		goto L646
	}
L646:
	;
	v2322 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(v2322)+60))
	F_BufFileDeleteFileSet(m, v2323, v2300, int32(1))
	mBase = m.M
	v2326 = m.ExcPending
	if v2326 != 0 {
		goto L1
	} else {
		goto L647
	}
L647:
	;
	goto L640
L648:
	;
	v2335 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(v2335)+20))
	goto L652
L649:
	;
	v2332 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v2332
	goto L651
L650:
	;
	goto L651
L651:
	;
	goto L648
L652:
	;
	if base.B2i32(v2336 == int32(2)) == int32(0) {
		goto L653
	} else {
		goto L654
	}
L653:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2342 = m.ExcPending
	if v2342 != 0 {
		goto L1
	} else {
		goto L656
	}
L654:
	;
	goto L655
L655:
	;
	v2345 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L1
	} else {
		goto L658
	}
L656:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		goto L1
	} else {
		goto L657
	}
L657:
	;
	goto L655
L658:
	;
	F_PushActiveSnapshot(m, v2345)
	mBase = m.M
	v2348 = m.ExcPending
	if v2348 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
	;
	v2351 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v2351
	v2354 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2355 = *(*int32)(unsafe.Add(mBase, uint32(v2354)+32))
	F_subxact_info_read(m, v2355, v2269)
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		goto L1
	} else {
		goto L660
	}
L660:
	;
	v2359 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[29]))
	v2361 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_apply_dispatch[30])))
	v2378 = v2361
	goto L663
L661:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2628 = m.ExcPending
	if v2628 != 0 {
		goto L1
	} else {
		goto L715
	}
L662:
	;
	if v2359 != 0 {
		goto L711
	} else {
		goto L712
	}
L663:
	;
	if v2378 <= int64(0) {
		goto L662
	} else {
		goto L665
	}
L664:
	;
	v2391 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2392 = *(*int32)(unsafe.Add(mBase, uint32(v2391)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+160)) = v2392
	*(*int32)(unsafe.Add(mBase, uint32(v21)+164)) = v2269
	v2396 = v21 + int32(320)
	v2401 = F_pg_snprintf(m, v2396, int32(1024), int32(_a_F_apply_dispatch_32), v21+int32(160))
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L1
	} else {
		goto L667
	}
L665:
	;
	v2383 = v2378 - int64(1)
	v2384 = base.I32_wrap_i64(v2383)
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v2359+v2384<<(uint(int32(4))%32))))
	if v2388 != v2264 {
		v2378 = v2383
		goto L663
	} else {
		goto L666
	}
L666:
	;
	goto L664
L667:
	;
	v2404 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2405 = *(*int32)(unsafe.Add(mBase, uint32(v2404)+60))
	v2408 = F_BufFileOpenFileSet(m, v2405, v2396, int32(2), int32(0))
	mBase = m.M
	v2409 = m.ExcPending
	if v2409 != 0 {
		goto L1
	} else {
		goto L668
	}
L668:
	;
	v2411 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[29]))
	v2414 = v2411 + v2384<<(uint(int32(4))%32)
	v2415 = *(*int32)(unsafe.Add(mBase, uint32(v2414)+4))
	v2416 = *(*int64)(unsafe.Add(mBase, uint32(v2414)+8))
	v2417 = m.G0
	v2419 = v2417 - int32(1072)
	m.G0 = v2419
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(v2408)))
	v2423 = v2421 - int32(1)
	if v2423 < v2415 {
		goto L673
	} else {
		goto L674
	}
L669:
	;
	F_BufFileClose(m, v2408)
	mBase = m.M
	v2593 = m.ExcPending
	if v2593 != 0 {
		goto L1
	} else {
		goto L709
	}
L670:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2567 = m.ExcPending
	if v2567 != 0 {
		goto L1
	} else {
		goto L704
	}
L671:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2550 = m.ExcPending
	if v2550 != 0 {
		goto L1
	} else {
		goto L700
	}
L672:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2408))) = v2506
	v2519 = *(*int32)(unsafe.Add(mBase, uint32(v2408)+24))
	if v2501 != v2519 {
		goto L690
	} else {
		goto L691
	}
L673:
	;
	v2425 = *(*int64)(unsafe.Add(mBase, uint32(v2408)+32))
	v2501 = v2415
	v2506 = v2421
	v2516 = v2425
	goto L672
L674:
	;
	goto L675
L675:
	;
	v2427 = v2415
	v2430 = v2423
	v2432 = v2421
	goto L676
L676:
	;
	v2444 = int32(0)
	if base.B2i32(v2430 == v2444)|base.B2i32(base.B2i32(v2416 == int64(0))|base.B2i32(v2415 != v2430) == v2444) == v2444 {
		goto L679
	} else {
		goto L680
	}
L677:
	;
	v2501 = v2493
	v2506 = v2494
	v2516 = v2496
	goto L672
L678:
	;
	v2498 = v2430 - int32(1)
	if v2415 <= v2498 {
		v2427 = v2493
		v2430 = v2498
		v2432 = v2494
		goto L676
	} else {
		goto L688
	}
L679:
	;
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(v2408)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2419)+16)) = v2455
	*(*int32)(unsafe.Add(mBase, uint32(v2419)+20)) = v2430
	v2459 = v2419 + int32(48)
	v2464 = F_pg_snprintf(m, v2459, int32(1024), int32(_a_F_apply_dispatch_34), v2419+int32(16))
	mBase = m.M
	v2465 = m.ExcPending
	if v2465 != 0 {
		goto L1
	} else {
		goto L682
	}
L680:
	;
	goto L681
L681:
	;
	v2483 = *(*int32)(unsafe.Add(mBase, uint32(v2408)+4))
	v2487 = *(*int32)(unsafe.Add(mBase, uint32(v2483+v2430<<(uint(int32(2))%32))))
	v2489 = F_FileTruncate(m, v2487, v2416, int32(167772167))
	mBase = m.M
	v2490 = m.ExcPending
	if v2490 != 0 {
		goto L1
	} else {
		goto L686
	}
L682:
	;
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(v2408)+4))
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(v2466+v2430<<(uint(int32(2))%32))))
	F_FileClose(m, v2470)
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		goto L1
	} else {
		goto L683
	}
L683:
	;
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v2408)+12))
	v2474 = F_FileSetDelete(m, v2473, v2459)
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L1
	} else {
		goto L684
	}
L684:
	;
	if v2474 == int32(0) {
		goto L671
	} else {
		goto L685
	}
L685:
	;
	v2493 = v2427 - base.B2i32(v2415 == v2430)
	v2494 = v2432 - int32(1)
	v2496 = int64(1073741824)
	goto L678
L686:
	;
	if v2489 < int32(0) {
		goto L670
	} else {
		goto L687
	}
L687:
	;
	v2493 = v2427
	v2494 = v2432
	v2496 = v2416
	goto L678
L688:
	;
	goto L677
L689:
	;
	m.G0 = v2419 + int32(1072)
	goto L669
L690:
	;
	if v2519 <= v2501 {
		goto L689
	} else {
		goto L699
	}
L691:
	;
	v2521 = *(*int64)(unsafe.Add(mBase, uint32(v2408)+32))
	if v2521 <= v2516 {
		goto L692
	} else {
		goto L693
	}
L692:
	;
	v2523 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2408)+44)))
	if v2521+v2523 < v2516 {
		goto L690
	} else {
		goto L695
	}
L693:
	;
	goto L694
L694:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2408)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2408)+32)) = v2516
	goto L689
L695:
	;
	v2527 = base.I32_wrap_i64(v2516 - v2521)
	v2528 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2408)+40)))
	if v2516 <= v2521+v2528 {
		goto L696
	} else {
		goto L697
	}
L696:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2408)+40)) = v2527
	goto L698
L697:
	;
	goto L698
L698:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2408)+44)) = v2527
	goto L689
L699:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2408)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2408)+32)) = v2516
	*(*int32)(unsafe.Add(mBase, uint32(v2408)+24)) = v2501
	goto L689
L700:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2552 = m.ExcPending
	if v2552 != 0 {
		goto L1
	} else {
		goto L701
	}
L701:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2419))) = v2419 + int32(48)
	F_errmsg(m, int32(_a_F_apply_dispatch_35), v2419)
	mBase = m.M
	v2558 = m.ExcPending
	if v2558 != 0 {
		goto L1
	} else {
		goto L702
	}
L702:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_36), int32(952), int32(_a_F_apply_dispatch_37))
	mBase = m.M
	v2563 = m.ExcPending
	if v2563 != 0 {
		goto L1
	} else {
		goto L703
	}
L703:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L704:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2569 = m.ExcPending
	if v2569 != 0 {
		goto L1
	} else {
		goto L705
	}
L705:
	;
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(v2408)+4))
	v2574 = *(*int32)(unsafe.Add(mBase, uint32(v2570+v2430<<(uint(int32(2))%32))))
	v2576 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[31]))
	v2580 = *(*int32)(unsafe.Add(mBase, uint32(v2576+v2574*int32(48))+32))
	goto L706
L706:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2419)+32)) = v2580
	F_errmsg(m, int32(_a_F_apply_dispatch_38), v2419+int32(32))
	mBase = m.M
	v2586 = m.ExcPending
	if v2586 != 0 {
		goto L1
	} else {
		goto L707
	}
L707:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_36), int32(970), int32(_a_F_apply_dispatch_37))
	mBase = m.M
	v2591 = m.ExcPending
	if v2591 != 0 {
		goto L1
	} else {
		goto L708
	}
L708:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L709:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[30])) = v2384
	v2597 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2598 = *(*int32)(unsafe.Add(mBase, uint32(v2597)+32))
	F_subxact_info_write(m, v2598, v2269)
	mBase = m.M
	v2600 = m.ExcPending
	if v2600 != 0 {
		goto L1
	} else {
		goto L710
	}
L710:
	;
	goto L661
L711:
	;
	F_pfree(m, v2359)
	mBase = m.M
	v2602 = m.ExcPending
	if v2602 != 0 {
		goto L1
	} else {
		goto L714
	}
L712:
	;
	goto L713
L713:
	;
	v2604 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[32])) = v2604
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[30])) = v2604
	goto L661
L714:
	;
	goto L713
L715:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2630 = m.ExcPending
	if v2630 != 0 {
		goto L1
	} else {
		goto L716
	}
L716:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2632 = m.ExcPending
	if v2632 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	goto L640
L718:
	;
	if v2653 == int32(0) {
		goto L10
	} else {
		goto L719
	}
L719:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_39), int32(0))
	mBase = m.M
	v2660 = m.ExcPending
	if v2660 != 0 {
		goto L1
	} else {
		goto L720
	}
L720:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1868), int32(_a_F_apply_dispatch_40))
	mBase = m.M
	v2665 = m.ExcPending
	if v2665 != 0 {
		goto L1
	} else {
		goto L721
	}
L721:
	;
	goto L10
L722:
	;
	v2677 = m.G0
	v2679 = v2677 - int32(96)
	m.G0 = v2679
	v2682 = v21 + int32(1344)
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v2682)+4))
	v2684 = *(*int32)(unsafe.Add(mBase, uint32(v2682)))
	v2686 = *(*int64)(unsafe.Add(mBase, uint32(v2682)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[33])) = v2686
	v2689 = *(*int64)(unsafe.Add(mBase, uint32(v2682)+16))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[34])) = v2689
	if v2683 == v2684 {
		goto L727
	} else {
		goto L728
	}
L723:
	;
	v2668 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24]))
	if v2668 == int32(0) {
		goto L722
	} else {
		goto L724
	}
L724:
	;
	F_BufFileClose(m, v2668)
	mBase = m.M
	v2672 = m.ExcPending
	if v2672 != 0 {
		goto L1
	} else {
		goto L725
	}
L725:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24])) = int32(0)
	goto L722
L726:
	;
	m.G0 = v2679 + int32(96)
	if v2264 != v2269 {
		goto L765
	} else {
		goto L766
	}
L727:
	;
	v2693 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[26]))
	v2694 = *(*int32)(unsafe.Add(mBase, uint32(v2693)))
	*(*int32)(unsafe.Add(mBase, uint32(v2693))) = int32(1)
	if v2694 != 0 {
		goto L730
	} else {
		goto L731
	}
L728:
	;
	goto L729
L729:
	;
	v2732 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(v2732)))
	*(*int32)(unsafe.Add(mBase, uint32(v2679)+16)) = v2733
	*(*int32)(unsafe.Add(mBase, uint32(v2679)+20)) = v2683
	v2737 = v2679 + int32(32)
	v2742 = F_pg_snprintf(m, v2737, int32(64), int32(_a_F_apply_dispatch_41), v2679+int32(16))
	mBase = m.M
	v2743 = m.ExcPending
	if v2743 != 0 {
		goto L1
	} else {
		goto L742
	}
L730:
	;
	F_s_lock(m, v2693, int32(_a_F_apply_dispatch_27), int32(1317), int32(_a_F_apply_dispatch_42))
	mBase = m.M
	v2701 = m.ExcPending
	if v2701 != 0 {
		goto L1
	} else {
		goto L733
	}
L731:
	;
	goto L732
L732:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2693))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2693)+8)) = int32(2)
	v2707 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2708 = *(*int32)(unsafe.Add(mBase, uint32(v2707)+32))
	F_UnlockApplyTransactionForSession(m, v2708, v2684, int32(1), int32(8))
	mBase = m.M
	v2712 = m.ExcPending
	if v2712 != 0 {
		goto L1
	} else {
		goto L734
	}
L733:
	;
	goto L732
L734:
	;
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v2714 = m.ExcPending
	if v2714 != 0 {
		goto L1
	} else {
		goto L735
	}
L735:
	;
	v2716 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v2716)+24))
	goto L736
L736:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2717) {
		goto L737
	} else {
		goto L738
	}
L737:
	;
	v2721 = F_EndTransactionBlock(m, int32(0))
	mBase = m.M
	v2722 = m.ExcPending
	if v2722 != 0 {
		goto L1
	} else {
		goto L740
	}
L738:
	;
	goto L739
L739:
	;
	v2726 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[35])) = v2726
	F_pgstat_report_activity(m, int32(2), v2726)
	mBase = m.M
	goto L726
L740:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2724 = m.ExcPending
	if v2724 != 0 {
		goto L1
	} else {
		goto L741
	}
L741:
	;
	goto L739
L742:
	;
	v2746 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2747 = m.ExcPending
	if v2747 != 0 {
		goto L1
	} else {
		goto L743
	}
L743:
	;
	if v2746 != 0 {
		goto L744
	} else {
		goto L745
	}
L744:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2679))) = v2737
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_43), v2679)
	mBase = m.M
	v2751 = m.ExcPending
	if v2751 != 0 {
		goto L1
	} else {
		goto L747
	}
L745:
	;
	goto L746
L746:
	;
	v2759 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[35]))
	if v2759 != 0 {
		goto L749
	} else {
		goto L750
	}
L747:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_27), int32(1475), int32(_a_F_apply_dispatch_44))
	mBase = m.M
	v2756 = m.ExcPending
	if v2756 != 0 {
		goto L1
	} else {
		goto L748
	}
L748:
	;
	goto L746
L749:
	;
	v2760 = *(*int32)(unsafe.Add(mBase, uint32(v2759)+4))
	v2761 = v2760
	goto L751
L750:
	;
	v2761 = int32(0)
	goto L751
L751:
	;
	v2762 = v2761
	goto L752
L752:
	;
	v2781 = v2762 - int32(1)
	if v2781 < int32(0) {
		goto L726
	} else {
		goto L754
	}
L753:
	;
	F_RollbackToSavepoint(m, v2679+int32(32))
	mBase = m.M
	v2793 = m.ExcPending
	if v2793 != 0 {
		goto L1
	} else {
		goto L756
	}
L754:
	;
	v2784 = *(*int32)(unsafe.Add(mBase, uint32(v2759)+12))
	v2788 = *(*int32)(unsafe.Add(mBase, uint32(v2784+v2781<<(uint(int32(2))%32))))
	if v2788 != v2683 {
		v2762 = v2781
		goto L752
	} else {
		goto L755
	}
L755:
	;
	goto L753
L756:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2795 = m.ExcPending
	if v2795 != 0 {
		goto L1
	} else {
		goto L757
	}
L757:
	;
	v2796 = int32(_a_F_apply_dispatch_45)
	v2798 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[35]))
	v2799 = int32(0)
	if base.B2i32(v2798 == v2799)|base.B2i32(v2781 <= v2799) != 0 {
		goto L759
	} else {
		goto L760
	}
L758:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[35])) = v2809
	goto L726
L759:
	;
	v2809 = int32(0)
	goto L761
L760:
	;
	v2806 = *(*int32)(unsafe.Add(mBase, uint32(v2798)+4))
	if v2781 < v2806 {
		goto L762
	} else {
		goto L763
	}
L761:
	;
	goto L758
L762:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2798)+4)) = v2781
	goto L764
L763:
	;
	goto L764
L764:
	;
	v2809 = v2798
	goto L761
L765:
	;
	F_pa_decr_and_wait_stream_block(m)
	mBase = m.M
	v2834 = m.ExcPending
	if v2834 != 0 {
		goto L1
	} else {
		goto L768
	}
L766:
	;
	goto L767
L767:
	;
	v2837 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2838 = m.ExcPending
	if v2838 != 0 {
		goto L1
	} else {
		goto L769
	}
L768:
	;
	goto L767
L769:
	;
	if v2837 == int32(0) {
		goto L10
	} else {
		goto L770
	}
L770:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_39), int32(0))
	mBase = m.M
	v2844 = m.ExcPending
	if v2844 != 0 {
		goto L1
	} else {
		goto L771
	}
L771:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1971), int32(_a_F_apply_dispatch_40))
	mBase = m.M
	v2849 = m.ExcPending
	if v2849 != 0 {
		goto L1
	} else {
		goto L772
	}
L772:
	;
	goto L10
L773:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = int32(1)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_46), v21+int32(112))
	mBase = m.M
	v2860 = m.ExcPending
	if v2860 != 0 {
		goto L1
	} else {
		goto L774
	}
L774:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1975), int32(_a_F_apply_dispatch_40))
	mBase = m.M
	v2865 = m.ExcPending
	if v2865 != 0 {
		goto L1
	} else {
		goto L775
	}
L775:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L776:
	;
	v2868 = *(*int32)(unsafe.Add(mBase, uint32(v2278)+16))
	v2869 = *(*int32)(unsafe.Add(mBase, uint32(v2868)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2868)+20)) = v2869 + int32(1)
	F_pa_lock_stream(m, v2269)
	mBase = m.M
	v2874 = m.ExcPending
	if v2874 != 0 {
		goto L1
	} else {
		goto L777
	}
L777:
	;
	v2875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2876 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2877 = F_pa_send_data(m, v2278, v2875, v2876)
	mBase = m.M
	v2878 = m.ExcPending
	if v2878 != 0 {
		goto L1
	} else {
		goto L778
	}
L778:
	;
	if v2877 == int32(0) {
		goto L12
	} else {
		goto L779
	}
L779:
	;
	goto L10
L780:
	;
	v2888 = v21 + int32(320)
	v2889 = m.G0
	v2891 = v2889 - int32(16)
	m.G0 = v2891
	v2894 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v2895 = m.ExcPending
	if v2895 != 0 {
		goto L1
	} else {
		goto L781
	}
L781:
	;
	v2896 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v2897 = m.ExcPending
	if v2897 != 0 {
		goto L1
	} else {
		goto L782
	}
L782:
	;
	v2899 = v2896 & int32(255)
	if v2899 != 0 {
		goto L783
	} else {
		goto L784
	}
L783:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2903 = m.ExcPending
	if v2903 != 0 {
		goto L1
	} else {
		goto L786
	}
L784:
	;
	goto L785
L785:
	;
	v2913 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2914 = m.ExcPending
	if v2914 != 0 {
		goto L1
	} else {
		goto L789
	}
L786:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2891))) = v2899
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_9), v2891)
	mBase = m.M
	v2907 = m.ExcPending
	if v2907 != 0 {
		goto L1
	} else {
		goto L787
	}
L787:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(1140), int32(_a_F_apply_dispatch_47))
	mBase = m.M
	v2912 = m.ExcPending
	if v2912 != 0 {
		goto L1
	} else {
		goto L788
	}
L788:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L789:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2888))) = v2913
	v2916 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2917 = m.ExcPending
	if v2917 != 0 {
		goto L1
	} else {
		goto L790
	}
L790:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2888)+8)) = v2916
	v2919 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v2920 = m.ExcPending
	if v2920 != 0 {
		goto L1
	} else {
		goto L791
	}
L791:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2888)+16)) = v2919
	m.G0 = v2891 + int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v2894
	v2928 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v2928
	v2931 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2931)+16)))
	if v2932 == int32(1) {
		goto L794
	} else {
		goto L795
	}
L792:
	;
	v3037 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	F_process_syncing_tables(m, v3037)
	mBase = m.M
	v3039 = m.ExcPending
	if v3039 != 0 {
		goto L1
	} else {
		goto L832
	}
L793:
	;
	v2997 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24]))
	if v2997 != 0 {
		goto L820
	} else {
		goto L821
	}
L794:
	;
	v2935 = *(*int32)(unsafe.Add(mBase, uint32(v2931)))
	if v2935 == int32(3) {
		goto L793
	} else {
		goto L797
	}
L795:
	;
	goto L796
L796:
	;
	v2938 = F_pa_find_worker(m, v2894)
	mBase = m.M
	v2939 = m.ExcPending
	if v2939 != 0 {
		goto L1
	} else {
		goto L800
	}
L797:
	;
	goto L796
L798:
	;
	F_stream_open_and_write_change(m, v2894, int32(99), v21+int32(1344))
	mBase = m.M
	v2989 = m.ExcPending
	if v2989 != 0 {
		goto L1
	} else {
		goto L817
	}
L799:
	;
	F_pa_switch_to_partial_serialize(m, v2938, int32(1))
	mBase = m.M
	v2984 = m.ExcPending
	if v2984 != 0 {
		goto L1
	} else {
		goto L816
	}
L800:
	;
	if v2938 != 0 {
		goto L801
	} else {
		goto L802
	}
L801:
	;
	v2940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2938)+12)))
	if v2940 != 0 {
		goto L798
	} else {
		goto L804
	}
L802:
	;
	goto L803
L803:
	;
	v2951 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v2951 != 0 {
		goto L19
	} else {
		goto L808
	}
L804:
	;
	v2941 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2942 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2943 = F_pa_send_data(m, v2938, v2941, v2942)
	mBase = m.M
	v2944 = m.ExcPending
	if v2944 != 0 {
		goto L1
	} else {
		goto L805
	}
L805:
	;
	if v2943 == int32(0) {
		goto L799
	} else {
		goto L806
	}
L806:
	;
	v2947 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	F_pa_xact_finish(m, v2938, v2947)
	mBase = m.M
	v2949 = m.ExcPending
	if v2949 != 0 {
		goto L1
	} else {
		goto L807
	}
L807:
	;
	goto L792
L808:
	;
	v2953 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2954 = *(*int32)(unsafe.Add(mBase, uint32(v2953)+60))
	v2955 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	F_apply_spooled_messages(m, v2954, v2894, v2955)
	mBase = m.M
	v2957 = m.ExcPending
	if v2957 != 0 {
		goto L1
	} else {
		goto L809
	}
L809:
	;
	F_apply_handle_commit_internal(m, v21+int32(320))
	mBase = m.M
	v2961 = m.ExcPending
	if v2961 != 0 {
		goto L1
	} else {
		goto L810
	}
L810:
	;
	v2963 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v2964 = *(*int32)(unsafe.Add(mBase, uint32(v2963)+32))
	F_stream_cleanup_files(m, v2964, v2894)
	mBase = m.M
	v2966 = m.ExcPending
	if v2966 != 0 {
		goto L1
	} else {
		goto L811
	}
L811:
	;
	v2969 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2970 = m.ExcPending
	if v2970 != 0 {
		goto L1
	} else {
		goto L812
	}
L812:
	;
	if v2969 == int32(0) {
		goto L792
	} else {
		goto L813
	}
L813:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_48), int32(0))
	mBase = m.M
	v2976 = m.ExcPending
	if v2976 != 0 {
		goto L1
	} else {
		goto L814
	}
L814:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(2184), int32(_a_F_apply_dispatch_49))
	mBase = m.M
	v2981 = m.ExcPending
	if v2981 != 0 {
		goto L1
	} else {
		goto L815
	}
L815:
	;
	goto L792
L816:
	;
	goto L798
L817:
	;
	v2990 = *(*int32)(unsafe.Add(mBase, uint32(v2938)+16))
	F_pa_set_fileset_state(m, v2990)
	mBase = m.M
	v2992 = m.ExcPending
	if v2992 != 0 {
		goto L1
	} else {
		goto L818
	}
L818:
	;
	v2993 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	F_pa_xact_finish(m, v2938, v2993)
	mBase = m.M
	v2995 = m.ExcPending
	if v2995 != 0 {
		goto L1
	} else {
		goto L819
	}
L819:
	;
	goto L792
L820:
	;
	F_BufFileClose(m, v2997)
	mBase = m.M
	v2999 = m.ExcPending
	if v2999 != 0 {
		goto L1
	} else {
		goto L823
	}
L821:
	;
	goto L822
L822:
	;
	F_apply_handle_commit_internal(m, v21+int32(320))
	mBase = m.M
	v3006 = m.ExcPending
	if v3006 != 0 {
		goto L1
	} else {
		goto L824
	}
L823:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24])) = int32(0)
	goto L822
L824:
	;
	v3008 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[26]))
	v3010 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[36]))
	*(*int64)(unsafe.Add(mBase, uint32(v3008)+24)) = v3010
	F_pa_set_xact_state(m, v3008, int32(2))
	mBase = m.M
	v3014 = m.ExcPending
	if v3014 != 0 {
		goto L1
	} else {
		goto L825
	}
L825:
	;
	F_pa_unlock_transaction(m, v2894)
	mBase = m.M
	v3016 = m.ExcPending
	if v3016 != 0 {
		goto L1
	} else {
		goto L826
	}
L826:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[35])) = int32(0)
	goto L827
L827:
	;
	v3022 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3023 = m.ExcPending
	if v3023 != 0 {
		goto L1
	} else {
		goto L828
	}
L828:
	;
	if v3022 == int32(0) {
		goto L792
	} else {
		goto L829
	}
L829:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_48), int32(0))
	mBase = m.M
	v3029 = m.ExcPending
	if v3029 != 0 {
		goto L1
	} else {
		goto L830
	}
L830:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(2238), int32(_a_F_apply_dispatch_49))
	mBase = m.M
	v3034 = m.ExcPending
	if v3034 != 0 {
		goto L1
	} else {
		goto L831
	}
L831:
	;
	goto L792
L832:
	;
	v3041 = int32(0)
	F_pgstat_report_activity(m, int32(2), v3041)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v3041
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v3041
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v3041
	goto L3
L833:
	;
	v3063 = *(*int32)(unsafe.Add(mBase, uint32(v3059)))
	if v3063 == int32(1) {
		goto L18
	} else {
		goto L836
	}
L834:
	;
	goto L835
L835:
	;
	v3067 = v21 + int32(320)
	v3068 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3069 = m.ExcPending
	if v3069 != 0 {
		goto L1
	} else {
		goto L837
	}
L836:
	;
	goto L835
L837:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3067))) = v3068
	if v3068 != int64(0) {
		goto L840
	} else {
		goto L841
	}
L838:
	;
	v3235 = *(*int32)(unsafe.Add(mBase, uint32(v21)+344))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v3235
	v3238 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v3238
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[9])) = v3238
	v3243 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v3244 = *(*int64)(unsafe.Add(mBase, uint32(v3243)+8))
	if base.B2i32(v3244 == int64(0))|base.B2i32(v3238 != v3244) != 0 {
		goto L885
	} else {
		goto L886
	}
L839:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3224 = m.ExcPending
	if v3224 != 0 {
		goto L1
	} else {
		goto L882
	}
L840:
	;
	v3073 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3074 = m.ExcPending
	if v3074 != 0 {
		goto L1
	} else {
		goto L843
	}
L841:
	;
	goto L842
L842:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3211 = m.ExcPending
	if v3211 != 0 {
		goto L1
	} else {
		goto L879
	}
L843:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3067)+8)) = v3073
	if v3073 == int64(0) {
		goto L839
	} else {
		goto L844
	}
L844:
	;
	v3078 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3079 = m.ExcPending
	if v3079 != 0 {
		goto L1
	} else {
		goto L845
	}
L845:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3067)+16)) = v3078
	v3082 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v3083 = m.ExcPending
	if v3083 != 0 {
		goto L1
	} else {
		goto L846
	}
L846:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3067)+24)) = v3082
	v3086 = v21 + int32(348)
	v3087 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v3088 = m.ExcPending
	if v3088 != 0 {
		goto L1
	} else {
		goto L847
	}
L847:
	;
	goto L851
L848:
	;
	goto L838
L849:
	;
	v3205 = F_strlen(m, v3194)
	mBase = m.M
	goto L848
L851:
	;
	goto L852
L852:
	;
	v3095 = int32(199)
	if (v3086^v3087)&int32(3) != 0 {
		goto L856
	} else {
		goto L857
	}
L853:
	;
	v3198 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3195))) = uint8(v3198)
	goto L849
L854:
	;
	v3179 = v3174
	v3180 = v3175
	v3181 = v3176
	goto L875
L855:
	;
	if v3169 == int32(0) {
		v3194 = v3167
		v3195 = v3168
		goto L853
	} else {
		goto L874
	}
L856:
	;
	v3167 = v3087
	v3168 = v3086
	v3169 = v3095
	goto L855
L857:
	;
	goto L858
L858:
	;
	v3099 = int32(0)
	if base.B2i32(v3087&int32(3) == v3099)|int32(0) == v3099 {
		goto L860
	} else {
		goto L861
	}
L859:
	;
	if v3135 == int32(0) {
		v3194 = v3132
		v3195 = v3133
		goto L853
	} else {
		goto L868
	}
L860:
	;
	v3111 = v3087
	v3112 = v3086
	v3113 = v3095
	goto L863
L861:
	;
	goto L862
L862:
	;
	v3132 = v3087
	v3133 = v3086
	v3134 = v3095
	v3135 = int32(1)
	goto L859
L863:
	;
	v3115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3111))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3112))) = uint8(v3115)
	if v3115 == int32(0) {
		v3174 = v3111
		v3175 = v3112
		v3176 = v3113
		goto L854
	} else {
		goto L865
	}
L864:
	;
	v3132 = v3126
	v3133 = v3120
	v3134 = v3122
	v3135 = v3124
	goto L859
L865:
	;
	v3119 = int32(1)
	v3120 = v3112 + v3119
	v3122 = v3113 - v3119
	v3123 = int32(0)
	v3124 = base.B2i32(v3122 != v3123)
	v3126 = v3111 + v3119
	if v3126&int32(3) == v3123 {
		v3132 = v3126
		v3133 = v3120
		v3134 = v3122
		v3135 = v3124
		goto L859
	} else {
		goto L866
	}
L866:
	;
	if v3122 != 0 {
		v3111 = v3126
		v3112 = v3120
		v3113 = v3122
		goto L863
	} else {
		goto L867
	}
L867:
	;
	goto L864
L868:
	;
	v3138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3132))))
	if base.B2i32(v3138 == int32(0))|base.B2i32(base.Ui32(v3134) < base.Ui32(int32(4))) != 0 {
		v3167 = v3132
		v3168 = v3133
		v3169 = v3134
		goto L855
	} else {
		goto L869
	}
L869:
	;
	v3145 = v3132
	v3146 = v3133
	v3147 = v3134
	goto L870
L870:
	;
	v3150 = *(*int32)(unsafe.Add(mBase, uint32(v3145)))
	v3153 = int32(-2139062144)
	if (int32(16843008)-v3150|v3150)&v3153 != v3153 {
		v3174 = v3145
		v3175 = v3146
		v3176 = v3147
		goto L854
	} else {
		goto L872
	}
L871:
	;
	v3167 = v3161
	v3168 = v3159
	v3169 = v3163
	goto L855
L872:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3146))) = v3150
	v3158 = int32(4)
	v3159 = v3146 + v3158
	v3161 = v3145 + v3158
	v3163 = v3147 - v3158
	if base.Ui32(int32(3)) < base.Ui32(v3163) {
		v3145 = v3161
		v3146 = v3159
		v3147 = v3163
		goto L870
	} else {
		goto L873
	}
L873:
	;
	goto L871
L874:
	;
	v3174 = v3167
	v3175 = v3168
	v3176 = v3169
	goto L854
L875:
	;
	v3183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3179))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3180))) = uint8(v3183)
	if v3183 == int32(0) {
		v3194 = v3179
		v3195 = v3180
		goto L853
	} else {
		goto L877
	}
L876:
	;
	v3194 = v3190
	v3195 = v3188
	goto L853
L877:
	;
	v3187 = int32(1)
	v3188 = v3180 + v3187
	v3190 = v3179 + v3187
	v3192 = v3181 - v3187
	if v3192 != 0 {
		v3179 = v3190
		v3180 = v3188
		v3181 = v3192
		goto L875
	} else {
		goto L878
	}
L878:
	;
	goto L876
L879:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_50), int32(0))
	mBase = m.M
	v3215 = m.ExcPending
	if v3215 != 0 {
		goto L1
	} else {
		goto L880
	}
L880:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(139), int32(_a_F_apply_dispatch_51))
	mBase = m.M
	v3220 = m.ExcPending
	if v3220 != 0 {
		goto L1
	} else {
		goto L881
	}
L881:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L882:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_52), int32(0))
	mBase = m.M
	v3228 = m.ExcPending
	if v3228 != 0 {
		goto L1
	} else {
		goto L883
	}
L883:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(142), int32(_a_F_apply_dispatch_51))
	mBase = m.M
	v3233 = m.ExcPending
	if v3233 != 0 {
		goto L1
	} else {
		goto L884
	}
L884:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L885:
	;
	v3275 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])) = uint8(v3275)
	F_pgstat_report_activity(m, int32(3), int32(0))
	mBase = m.M
	goto L3
L886:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8])) = v3238
	v3253 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3254 = m.ExcPending
	if v3254 != 0 {
		goto L1
	} else {
		goto L887
	}
L887:
	;
	if v3253 == int32(0) {
		goto L885
	} else {
		goto L888
	}
L888:
	;
	v3258 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+196)) = uint32(v3258)
	v3261 = int64(base.Ui64(v3258) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+192)) = uint32(v3261)
	F_errmsg(m, int32(_a_F_apply_dispatch_5), v21+int32(192))
	mBase = m.M
	v3267 = m.ExcPending
	if v3267 != 0 {
		goto L1
	} else {
		goto L889
	}
L889:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(_a_F_apply_dispatch_7), int32(_a_F_apply_dispatch_8))
	mBase = m.M
	v3272 = m.ExcPending
	if v3272 != 0 {
		goto L1
	} else {
		goto L890
	}
L890:
	;
	goto L885
L891:
	;
	v3285 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	v3287 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[9]))
	if v3285 != v3287 {
		goto L17
	} else {
		goto L892
	}
L892:
	;
	v3290 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v3290 < int32(0) {
		goto L894
	} else {
		goto L895
	}
L893:
	;
	v3297 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v3298 = *(*int32)(unsafe.Add(mBase, uint32(v3297)+20))
	goto L897
L894:
	;
	v3294 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v3294
	goto L896
L895:
	;
	goto L896
L896:
	;
	goto L893
L897:
	;
	if base.B2i32(v3298 == int32(2)) == int32(0) {
		goto L898
	} else {
		goto L899
	}
L898:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v3304 = m.ExcPending
	if v3304 != 0 {
		goto L1
	} else {
		goto L901
	}
L899:
	;
	goto L900
L900:
	;
	v3307 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v3308 = m.ExcPending
	if v3308 != 0 {
		goto L1
	} else {
		goto L903
	}
L901:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v3306 = m.ExcPending
	if v3306 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	goto L900
L903:
	;
	F_PushActiveSnapshot(m, v3307)
	mBase = m.M
	v3310 = m.ExcPending
	if v3310 != 0 {
		goto L1
	} else {
		goto L904
	}
L904:
	;
	v3313 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v3313
	v3316 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v3317 = *(*int32)(unsafe.Add(mBase, uint32(v3316)))
	v3318 = *(*int32)(unsafe.Add(mBase, uint32(v21)+344))
	F_TwoPhaseTransactionGid(m, v3317, v3318, v21+int32(1344))
	mBase = m.M
	v3322 = m.ExcPending
	if v3322 != 0 {
		goto L1
	} else {
		goto L905
	}
L905:
	;
	v3324 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v3325 = *(*int32)(unsafe.Add(mBase, uint32(v3324)+24))
	goto L906
L906:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(v3325)) == int32(0) {
		goto L907
	} else {
		goto L908
	}
L907:
	;
	F_BeginTransactionBlock(m)
	mBase = m.M
	v3331 = m.ExcPending
	if v3331 != 0 {
		goto L1
	} else {
		goto L910
	}
L908:
	;
	goto L909
L909:
	;
	v3335 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[33])) = v3335
	v3338 = *(*int64)(unsafe.Add(mBase, uint32(v21)+336))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[34])) = v3338
	v3342 = F_PrepareTransactionBlock(m, v21+int32(1344))
	mBase = m.M
	v3343 = m.ExcPending
	if v3343 != 0 {
		goto L1
	} else {
		goto L912
	}
L910:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3333 = m.ExcPending
	if v3333 != 0 {
		goto L1
	} else {
		goto L911
	}
L911:
	;
	goto L909
L912:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v3345 = m.ExcPending
	if v3345 != 0 {
		goto L1
	} else {
		goto L913
	}
L913:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v3347 = m.ExcPending
	if v3347 != 0 {
		goto L1
	} else {
		goto L914
	}
L914:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3349 = m.ExcPending
	if v3349 != 0 {
		goto L1
	} else {
		goto L915
	}
L915:
	;
	v3351 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v3352 = m.ExcPending
	if v3352 != 0 {
		goto L1
	} else {
		goto L916
	}
L916:
	;
	v3353 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	v3355 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v3356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3355)+16)))
	if v3356 == int32(1) {
		goto L918
	} else {
		goto L919
	}
L917:
	;
	v3396 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])) = uint8(v3396)
	F_process_syncing_tables(m, v3394)
	mBase = m.M
	v3399 = m.ExcPending
	if v3399 != 0 {
		goto L1
	} else {
		goto L927
	}
L918:
	;
	v3359 = *(*int32)(unsafe.Add(mBase, uint32(v3355)))
	if v3359 == int32(3) {
		v3394 = v3353
		goto L917
	} else {
		goto L921
	}
L919:
	;
	goto L920
L920:
	;
	v3364 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v3364
	v3367 = F_palloc(m, int32(24))
	mBase = m.M
	v3368 = m.ExcPending
	if v3368 != 0 {
		goto L1
	} else {
		goto L922
	}
L921:
	;
	goto L920
L922:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3367)+16)) = v3353
	*(*int64)(unsafe.Add(mBase, uint32(v3367)+8)) = int64(0)
	v3373 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[37]))
	if v3373 != 0 {
		goto L924
	} else {
		goto L925
	}
L923:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3367))) = v3380
	v3382 = int32(_a_F_apply_dispatch_53)
	*(*int32)(unsafe.Add(mBase, uint32(v3367)+4)) = v3382
	*(*int32)(unsafe.Add(mBase, uint32(v3380)+4)) = v3367
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[38])) = v3367
	v3389 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v3389
	v3391 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	v3394 = v3391
	goto L917
L924:
	;
	v3375 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[38]))
	v3380 = v3375
	goto L923
L925:
	;
	goto L926
L926:
	;
	v3377 = int32(_a_F_apply_dispatch_53)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[37])) = v3377
	v3380 = v3377
	goto L923
L927:
	;
	v3401 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	if v3401 != int64(0) {
		goto L928
	} else {
		goto L929
	}
L928:
	;
	v3406 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3407 = m.ExcPending
	if v3407 != 0 {
		goto L1
	} else {
		goto L931
	}
L929:
	;
	goto L930
L930:
	;
	v3429 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	F_clear_subscription_skip_lsn(m, v3429)
	mBase = m.M
	v3431 = m.ExcPending
	if v3431 != 0 {
		goto L1
	} else {
		goto L937
	}
L931:
	;
	if v3406 != 0 {
		goto L932
	} else {
		goto L933
	}
L932:
	;
	v3409 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+212)) = uint32(v3409)
	v3412 = int64(base.Ui64(v3409) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+208)) = uint32(v3412)
	F_errmsg(m, int32(_a_F_apply_dispatch_54), v21+int32(208))
	mBase = m.M
	v3418 = m.ExcPending
	if v3418 != 0 {
		goto L1
	} else {
		goto L935
	}
L933:
	;
	goto L934
L934:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8])) = int64(0)
	goto L930
L935:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(_a_F_apply_dispatch_55), int32(_a_F_apply_dispatch_56))
	mBase = m.M
	v3423 = m.ExcPending
	if v3423 != 0 {
		goto L1
	} else {
		goto L936
	}
L936:
	;
	goto L934
L937:
	;
	v3433 = int32(0)
	F_pgstat_report_activity(m, int32(2), v3433)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v3433
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v3433
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v3433
	goto L3
L938:
	;
	v3645 = *(*int32)(unsafe.Add(mBase, uint32(v21)+344))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v3645
	v3648 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v3648
	v3651 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v3652 = *(*int32)(unsafe.Add(mBase, uint32(v3651)))
	F_TwoPhaseTransactionGid(m, v3652, v3645, v21+int32(1344))
	mBase = m.M
	v3656 = m.ExcPending
	if v3656 != 0 {
		goto L1
	} else {
		goto L992
	}
L939:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3634 = m.ExcPending
	if v3634 != 0 {
		goto L1
	} else {
		goto L989
	}
L940:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3621 = m.ExcPending
	if v3621 != 0 {
		goto L1
	} else {
		goto L986
	}
L941:
	;
	v3459 = v3456 & int32(255)
	if v3459 == int32(0) {
		goto L942
	} else {
		goto L943
	}
L942:
	;
	v3462 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3463 = m.ExcPending
	if v3463 != 0 {
		goto L1
	} else {
		goto L945
	}
L943:
	;
	goto L944
L944:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3608 = m.ExcPending
	if v3608 != 0 {
		goto L1
	} else {
		goto L983
	}
L945:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3451))) = v3462
	if v3462 == int64(0) {
		goto L940
	} else {
		goto L946
	}
L946:
	;
	v3467 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3468 = m.ExcPending
	if v3468 != 0 {
		goto L1
	} else {
		goto L947
	}
L947:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3451)+8)) = v3467
	if v3467 == int64(0) {
		goto L939
	} else {
		goto L948
	}
L948:
	;
	v3472 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3473 = m.ExcPending
	if v3473 != 0 {
		goto L1
	} else {
		goto L949
	}
L949:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3451)+16)) = v3472
	v3476 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v3477 = m.ExcPending
	if v3477 != 0 {
		goto L1
	} else {
		goto L950
	}
L950:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3451)+24)) = v3476
	v3480 = v21 + int32(348)
	v3481 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v3482 = m.ExcPending
	if v3482 != 0 {
		goto L1
	} else {
		goto L951
	}
L951:
	;
	goto L955
L952:
	;
	m.G0 = v3454 + int32(16)
	goto L938
L953:
	;
	v3599 = F_strlen(m, v3588)
	mBase = m.M
	goto L952
L955:
	;
	goto L956
L956:
	;
	v3489 = int32(199)
	if (v3480^v3481)&int32(3) != 0 {
		goto L960
	} else {
		goto L961
	}
L957:
	;
	v3592 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3589))) = uint8(v3592)
	goto L953
L958:
	;
	v3573 = v3568
	v3574 = v3569
	v3575 = v3570
	goto L979
L959:
	;
	if v3563 == int32(0) {
		v3588 = v3561
		v3589 = v3562
		goto L957
	} else {
		goto L978
	}
L960:
	;
	v3561 = v3481
	v3562 = v3480
	v3563 = v3489
	goto L959
L961:
	;
	goto L962
L962:
	;
	v3493 = int32(0)
	if base.B2i32(v3481&int32(3) == v3493)|int32(0) == v3493 {
		goto L964
	} else {
		goto L965
	}
L963:
	;
	if v3529 == int32(0) {
		v3588 = v3526
		v3589 = v3527
		goto L957
	} else {
		goto L972
	}
L964:
	;
	v3505 = v3481
	v3506 = v3480
	v3507 = v3489
	goto L967
L965:
	;
	goto L966
L966:
	;
	v3526 = v3481
	v3527 = v3480
	v3528 = v3489
	v3529 = int32(1)
	goto L963
L967:
	;
	v3509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3505))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3506))) = uint8(v3509)
	if v3509 == int32(0) {
		v3568 = v3505
		v3569 = v3506
		v3570 = v3507
		goto L958
	} else {
		goto L969
	}
L968:
	;
	v3526 = v3520
	v3527 = v3514
	v3528 = v3516
	v3529 = v3518
	goto L963
L969:
	;
	v3513 = int32(1)
	v3514 = v3506 + v3513
	v3516 = v3507 - v3513
	v3517 = int32(0)
	v3518 = base.B2i32(v3516 != v3517)
	v3520 = v3505 + v3513
	if v3520&int32(3) == v3517 {
		v3526 = v3520
		v3527 = v3514
		v3528 = v3516
		v3529 = v3518
		goto L963
	} else {
		goto L970
	}
L970:
	;
	if v3516 != 0 {
		v3505 = v3520
		v3506 = v3514
		v3507 = v3516
		goto L967
	} else {
		goto L971
	}
L971:
	;
	goto L968
L972:
	;
	v3532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3526))))
	if base.B2i32(v3532 == int32(0))|base.B2i32(base.Ui32(v3528) < base.Ui32(int32(4))) != 0 {
		v3561 = v3526
		v3562 = v3527
		v3563 = v3528
		goto L959
	} else {
		goto L973
	}
L973:
	;
	v3539 = v3526
	v3540 = v3527
	v3541 = v3528
	goto L974
L974:
	;
	v3544 = *(*int32)(unsafe.Add(mBase, uint32(v3539)))
	v3547 = int32(-2139062144)
	if (int32(16843008)-v3544|v3544)&v3547 != v3547 {
		v3568 = v3539
		v3569 = v3540
		v3570 = v3541
		goto L958
	} else {
		goto L976
	}
L975:
	;
	v3561 = v3555
	v3562 = v3553
	v3563 = v3557
	goto L959
L976:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3540))) = v3544
	v3552 = int32(4)
	v3553 = v3540 + v3552
	v3555 = v3539 + v3552
	v3557 = v3541 - v3552
	if base.Ui32(int32(3)) < base.Ui32(v3557) {
		v3539 = v3555
		v3540 = v3553
		v3541 = v3557
		goto L974
	} else {
		goto L977
	}
L977:
	;
	goto L975
L978:
	;
	v3568 = v3561
	v3569 = v3562
	v3570 = v3563
	goto L958
L979:
	;
	v3577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3573))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3574))) = uint8(v3577)
	if v3577 == int32(0) {
		v3588 = v3573
		v3589 = v3574
		goto L957
	} else {
		goto L981
	}
L980:
	;
	v3588 = v3584
	v3589 = v3582
	goto L957
L981:
	;
	v3581 = int32(1)
	v3582 = v3574 + v3581
	v3584 = v3573 + v3581
	v3586 = v3575 - v3581
	if v3586 != 0 {
		v3573 = v3584
		v3574 = v3582
		v3575 = v3586
		goto L979
	} else {
		goto L982
	}
L982:
	;
	goto L980
L983:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3454))) = v3459
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_57), v3454)
	mBase = m.M
	v3612 = m.ExcPending
	if v3612 != 0 {
		goto L1
	} else {
		goto L984
	}
L984:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(273), int32(_a_F_apply_dispatch_58))
	mBase = m.M
	v3617 = m.ExcPending
	if v3617 != 0 {
		goto L1
	} else {
		goto L985
	}
L985:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L986:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_59), int32(0))
	mBase = m.M
	v3625 = m.ExcPending
	if v3625 != 0 {
		goto L1
	} else {
		goto L987
	}
L987:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(278), int32(_a_F_apply_dispatch_58))
	mBase = m.M
	v3630 = m.ExcPending
	if v3630 != 0 {
		goto L1
	} else {
		goto L988
	}
L988:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L989:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_60), int32(0))
	mBase = m.M
	v3638 = m.ExcPending
	if v3638 != 0 {
		goto L1
	} else {
		goto L990
	}
L990:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(281), int32(_a_F_apply_dispatch_58))
	mBase = m.M
	v3643 = m.ExcPending
	if v3643 != 0 {
		goto L1
	} else {
		goto L991
	}
L991:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L992:
	;
	v3658 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v3658 < int32(0) {
		goto L994
	} else {
		goto L995
	}
L993:
	;
	v3665 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v3666 = *(*int32)(unsafe.Add(mBase, uint32(v3665)+20))
	goto L997
L994:
	;
	v3662 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v3662
	goto L996
L995:
	;
	goto L996
L996:
	;
	goto L993
L997:
	;
	if base.B2i32(v3666 == int32(2)) == int32(0) {
		goto L998
	} else {
		goto L999
	}
L998:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v3672 = m.ExcPending
	if v3672 != 0 {
		goto L1
	} else {
		goto L1001
	}
L999:
	;
	goto L1000
L1000:
	;
	v3675 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v3676 = m.ExcPending
	if v3676 != 0 {
		goto L1
	} else {
		goto L1003
	}
L1001:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v3674 = m.ExcPending
	if v3674 != 0 {
		goto L1
	} else {
		goto L1002
	}
L1002:
	;
	goto L1000
L1003:
	;
	F_PushActiveSnapshot(m, v3675)
	mBase = m.M
	v3678 = m.ExcPending
	if v3678 != 0 {
		goto L1
	} else {
		goto L1004
	}
L1004:
	;
	v3681 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v3681
	v3684 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[33])) = v3684
	v3687 = *(*int64)(unsafe.Add(mBase, uint32(v21)+336))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[34])) = v3687
	F_FinishPreparedTransaction(m, v21+int32(1344), int32(1))
	mBase = m.M
	v3693 = m.ExcPending
	if v3693 != 0 {
		goto L1
	} else {
		goto L1005
	}
L1005:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v3695 = m.ExcPending
	if v3695 != 0 {
		goto L1
	} else {
		goto L1006
	}
L1006:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v3697 = m.ExcPending
	if v3697 != 0 {
		goto L1
	} else {
		goto L1007
	}
L1007:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3699 = m.ExcPending
	if v3699 != 0 {
		goto L1
	} else {
		goto L1008
	}
L1008:
	;
	v3701 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v3702 = m.ExcPending
	if v3702 != 0 {
		goto L1
	} else {
		goto L1009
	}
L1009:
	;
	v3704 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[36]))
	v3705 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	v3707 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v3708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3707)+16)))
	if v3708 == int32(1) {
		goto L1011
	} else {
		goto L1012
	}
L1010:
	;
	v3747 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])) = uint8(v3747)
	F_process_syncing_tables(m, v3745)
	mBase = m.M
	v3750 = m.ExcPending
	if v3750 != 0 {
		goto L1
	} else {
		goto L1020
	}
L1011:
	;
	v3711 = *(*int32)(unsafe.Add(mBase, uint32(v3707)))
	if v3711 == int32(3) {
		v3745 = v3705
		goto L1010
	} else {
		goto L1014
	}
L1012:
	;
	goto L1013
L1013:
	;
	v3716 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v3716
	v3719 = F_palloc(m, int32(24))
	mBase = m.M
	v3720 = m.ExcPending
	if v3720 != 0 {
		goto L1
	} else {
		goto L1015
	}
L1014:
	;
	goto L1013
L1015:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3719)+16)) = v3705
	*(*int64)(unsafe.Add(mBase, uint32(v3719)+8)) = v3704
	v3724 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[37]))
	if v3724 != 0 {
		goto L1017
	} else {
		goto L1018
	}
L1016:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3719))) = v3731
	v3733 = int32(_a_F_apply_dispatch_53)
	*(*int32)(unsafe.Add(mBase, uint32(v3719)+4)) = v3733
	*(*int32)(unsafe.Add(mBase, uint32(v3731)+4)) = v3719
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[38])) = v3719
	v3740 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v3740
	v3742 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	v3745 = v3742
	goto L1010
L1017:
	;
	v3726 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[38]))
	v3731 = v3726
	goto L1016
L1018:
	;
	goto L1019
L1019:
	;
	v3728 = int32(_a_F_apply_dispatch_53)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[37])) = v3728
	v3731 = v3728
	goto L1016
L1020:
	;
	v3751 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	F_clear_subscription_skip_lsn(m, v3751)
	mBase = m.M
	v3753 = m.ExcPending
	if v3753 != 0 {
		goto L1
	} else {
		goto L1021
	}
L1021:
	;
	v3755 = int32(0)
	F_pgstat_report_activity(m, int32(2), v3755)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v3755
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v3755
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v3755
	goto L3
L1022:
	;
	v3970 = *(*int32)(unsafe.Add(mBase, uint32(v21)+352))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v3970
	v3973 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v3973
	v3976 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[1]))
	v3977 = *(*int32)(unsafe.Add(mBase, uint32(v3976)))
	v3979 = v21 + int32(1344)
	F_TwoPhaseTransactionGid(m, v3977, v3970, v3979)
	mBase = m.M
	v3981 = m.ExcPending
	if v3981 != 0 {
		goto L1
	} else {
		goto L1077
	}
L1023:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3959 = m.ExcPending
	if v3959 != 0 {
		goto L1
	} else {
		goto L1074
	}
L1024:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3946 = m.ExcPending
	if v3946 != 0 {
		goto L1
	} else {
		goto L1071
	}
L1025:
	;
	v3781 = v3778 & int32(255)
	if v3781 == int32(0) {
		goto L1026
	} else {
		goto L1027
	}
L1026:
	;
	v3784 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3785 = m.ExcPending
	if v3785 != 0 {
		goto L1
	} else {
		goto L1029
	}
L1027:
	;
	goto L1028
L1028:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3933 = m.ExcPending
	if v3933 != 0 {
		goto L1
	} else {
		goto L1068
	}
L1029:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3773))) = v3784
	if v3784 == int64(0) {
		goto L1024
	} else {
		goto L1030
	}
L1030:
	;
	v3789 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3790 = m.ExcPending
	if v3790 != 0 {
		goto L1
	} else {
		goto L1031
	}
L1031:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3773)+8)) = v3789
	if v3789 == int64(0) {
		goto L1023
	} else {
		goto L1032
	}
L1032:
	;
	v3794 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3795 = m.ExcPending
	if v3795 != 0 {
		goto L1
	} else {
		goto L1033
	}
L1033:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3773)+16)) = v3794
	v3797 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v3798 = m.ExcPending
	if v3798 != 0 {
		goto L1
	} else {
		goto L1034
	}
L1034:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3773)+24)) = v3797
	v3801 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v3802 = m.ExcPending
	if v3802 != 0 {
		goto L1
	} else {
		goto L1035
	}
L1035:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3773)+32)) = v3801
	v3805 = v21 + int32(356)
	v3806 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v3807 = m.ExcPending
	if v3807 != 0 {
		goto L1
	} else {
		goto L1036
	}
L1036:
	;
	goto L1040
L1037:
	;
	m.G0 = v3776 + int32(16)
	goto L1022
L1038:
	;
	v3924 = F_strlen(m, v3913)
	mBase = m.M
	goto L1037
L1040:
	;
	goto L1041
L1041:
	;
	v3814 = int32(199)
	if (v3805^v3806)&int32(3) != 0 {
		goto L1045
	} else {
		goto L1046
	}
L1042:
	;
	v3917 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3914))) = uint8(v3917)
	goto L1038
L1043:
	;
	v3898 = v3893
	v3899 = v3894
	v3900 = v3895
	goto L1064
L1044:
	;
	if v3888 == int32(0) {
		v3913 = v3886
		v3914 = v3887
		goto L1042
	} else {
		goto L1063
	}
L1045:
	;
	v3886 = v3806
	v3887 = v3805
	v3888 = v3814
	goto L1044
L1046:
	;
	goto L1047
L1047:
	;
	v3818 = int32(0)
	if base.B2i32(v3806&int32(3) == v3818)|int32(0) == v3818 {
		goto L1049
	} else {
		goto L1050
	}
L1048:
	;
	if v3854 == int32(0) {
		v3913 = v3851
		v3914 = v3852
		goto L1042
	} else {
		goto L1057
	}
L1049:
	;
	v3830 = v3806
	v3831 = v3805
	v3832 = v3814
	goto L1052
L1050:
	;
	goto L1051
L1051:
	;
	v3851 = v3806
	v3852 = v3805
	v3853 = v3814
	v3854 = int32(1)
	goto L1048
L1052:
	;
	v3834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3830))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3831))) = uint8(v3834)
	if v3834 == int32(0) {
		v3893 = v3830
		v3894 = v3831
		v3895 = v3832
		goto L1043
	} else {
		goto L1054
	}
L1053:
	;
	v3851 = v3845
	v3852 = v3839
	v3853 = v3841
	v3854 = v3843
	goto L1048
L1054:
	;
	v3838 = int32(1)
	v3839 = v3831 + v3838
	v3841 = v3832 - v3838
	v3842 = int32(0)
	v3843 = base.B2i32(v3841 != v3842)
	v3845 = v3830 + v3838
	if v3845&int32(3) == v3842 {
		v3851 = v3845
		v3852 = v3839
		v3853 = v3841
		v3854 = v3843
		goto L1048
	} else {
		goto L1055
	}
L1055:
	;
	if v3841 != 0 {
		v3830 = v3845
		v3831 = v3839
		v3832 = v3841
		goto L1052
	} else {
		goto L1056
	}
L1056:
	;
	goto L1053
L1057:
	;
	v3857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3851))))
	if base.B2i32(v3857 == int32(0))|base.B2i32(base.Ui32(v3853) < base.Ui32(int32(4))) != 0 {
		v3886 = v3851
		v3887 = v3852
		v3888 = v3853
		goto L1044
	} else {
		goto L1058
	}
L1058:
	;
	v3864 = v3851
	v3865 = v3852
	v3866 = v3853
	goto L1059
L1059:
	;
	v3869 = *(*int32)(unsafe.Add(mBase, uint32(v3864)))
	v3872 = int32(-2139062144)
	if (int32(16843008)-v3869|v3869)&v3872 != v3872 {
		v3893 = v3864
		v3894 = v3865
		v3895 = v3866
		goto L1043
	} else {
		goto L1061
	}
L1060:
	;
	v3886 = v3880
	v3887 = v3878
	v3888 = v3882
	goto L1044
L1061:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3865))) = v3869
	v3877 = int32(4)
	v3878 = v3865 + v3877
	v3880 = v3864 + v3877
	v3882 = v3866 - v3877
	if base.Ui32(int32(3)) < base.Ui32(v3882) {
		v3864 = v3880
		v3865 = v3878
		v3866 = v3882
		goto L1059
	} else {
		goto L1062
	}
L1062:
	;
	goto L1060
L1063:
	;
	v3893 = v3886
	v3894 = v3887
	v3895 = v3888
	goto L1043
L1064:
	;
	v3902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3898))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3899))) = uint8(v3902)
	if v3902 == int32(0) {
		v3913 = v3898
		v3914 = v3899
		goto L1042
	} else {
		goto L1066
	}
L1065:
	;
	v3913 = v3909
	v3914 = v3907
	goto L1042
L1066:
	;
	v3906 = int32(1)
	v3907 = v3899 + v3906
	v3909 = v3898 + v3906
	v3911 = v3900 - v3906
	if v3911 != 0 {
		v3898 = v3909
		v3899 = v3907
		v3900 = v3911
		goto L1064
	} else {
		goto L1067
	}
L1067:
	;
	goto L1065
L1068:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3776))) = v3781
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_61), v3776)
	mBase = m.M
	v3937 = m.ExcPending
	if v3937 != 0 {
		goto L1
	} else {
		goto L1069
	}
L1069:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(332), int32(_a_F_apply_dispatch_62))
	mBase = m.M
	v3942 = m.ExcPending
	if v3942 != 0 {
		goto L1
	} else {
		goto L1070
	}
L1070:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1071:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_63), int32(0))
	mBase = m.M
	v3950 = m.ExcPending
	if v3950 != 0 {
		goto L1
	} else {
		goto L1072
	}
L1072:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(337), int32(_a_F_apply_dispatch_62))
	mBase = m.M
	v3955 = m.ExcPending
	if v3955 != 0 {
		goto L1
	} else {
		goto L1073
	}
L1073:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1074:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_64), int32(0))
	mBase = m.M
	v3963 = m.ExcPending
	if v3963 != 0 {
		goto L1
	} else {
		goto L1075
	}
L1075:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_3), int32(340), int32(_a_F_apply_dispatch_62))
	mBase = m.M
	v3968 = m.ExcPending
	if v3968 != 0 {
		goto L1
	} else {
		goto L1076
	}
L1076:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1077:
	;
	v3982 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	v3983 = *(*int64)(unsafe.Add(mBase, uint32(v21)+336))
	v3984 = int32(0)
	v3985 = m.G0
	v3987 = v3985 - int32(16)
	m.G0 = v3987
	v3990 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[39]))
	v3994 = F_LWLockAcquire(m, v3990+int32(2304), int32(1))
	mBase = m.M
	v3995 = m.ExcPending
	if v3995 != 0 {
		goto L1
	} else {
		goto L1078
	}
L1078:
	;
	v3997 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[40]))
	v3998 = *(*int32)(unsafe.Add(mBase, uint32(v3997)+4))
	if v3998 <= int32(0) {
		v4087 = v3984
		goto L1079
	} else {
		goto L1080
	}
L1079:
	;
	v4105 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[39]))
	F_LWLockRelease(m, v4105+int32(2304))
	mBase = m.M
	v4109 = m.ExcPending
	if v4109 != 0 {
		goto L1
	} else {
		goto L1107
	}
L1080:
	;
	v4001 = v3997
	v4002 = v3984
	goto L1082
L1081:
	;
	F_pfree(m, v4068)
	mBase = m.M
	v4084 = m.ExcPending
	if v4084 != 0 {
		goto L1
	} else {
		goto L1106
	}
L1082:
	;
	v4022 = *(*int32)(unsafe.Add(mBase, uint32(v4001+v4002<<(uint(int32(2))%32))+8))
	v4023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4022)+44)))
	if v4023 != int32(1) {
		v4077 = v4001
		goto L1084
	} else {
		goto L1085
	}
L1083:
	;
	v4087 = int32(0)
	goto L1079
L1084:
	;
	v4079 = v4002 + int32(1)
	v4080 = *(*int32)(unsafe.Add(mBase, uint32(v4077)+4))
	if v4079 < v4080 {
		v4001 = v4077
		v4002 = v4079
		goto L1082
	} else {
		goto L1105
	}
L1085:
	;
	v4027 = v4022 + int32(47)
	v4030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4027))))
	v4033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3979))))
	if base.B2i32(v4030 == int32(0))|base.B2i32(v4030 != v4033) != 0 {
		v4051 = v4030
		v4052 = v4033
		goto L1087
	} else {
		goto L1088
	}
L1086:
	;
	if v4051-v4052 != 0 {
		v4077 = v4001
		goto L1084
	} else {
		goto L1093
	}
L1087:
	;
	goto L1086
L1088:
	;
	v4036 = v4027
	v4037 = v3979
	goto L1089
L1089:
	;
	v4040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4037)+1)))
	v4041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4036)+1)))
	if v4041 == int32(0) {
		v4051 = v4041
		v4052 = v4040
		goto L1087
	} else {
		goto L1091
	}
L1090:
	;
	v4051 = v4041
	v4052 = v4040
	goto L1087
L1091:
	;
	v4044 = int32(1)
	if v4041 == v4040 {
		v4036 = v4036 + v4044
		v4037 = v4037 + v4044
		goto L1089
	} else {
		goto L1092
	}
L1092:
	;
	goto L1090
L1093:
	;
	v4054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4022)+45)))
	if v4054 == int32(1) {
		goto L1095
	} else {
		goto L1096
	}
L1094:
	;
	v4069 = *(*int64)(unsafe.Add(mBase, uint32(v4068)+56))
	if v3982 == v4069 {
		goto L1100
	} else {
		goto L1101
	}
L1095:
	;
	v4057 = *(*int32)(unsafe.Add(mBase, uint32(v4022)+32))
	v4059 = F_ReadTwoPhaseFile(m, v4057, int32(0))
	mBase = m.M
	v4060 = m.ExcPending
	if v4060 != 0 {
		goto L1
	} else {
		goto L1098
	}
L1096:
	;
	goto L1097
L1097:
	;
	v4061 = *(*int64)(unsafe.Add(mBase, uint32(v4022)+16))
	F_XlogReadTwoPhaseData(m, v4061, v3987+int32(12), int32(0))
	mBase = m.M
	v4066 = m.ExcPending
	if v4066 != 0 {
		goto L1
	} else {
		goto L1099
	}
L1098:
	;
	v4068 = v4059
	goto L1094
L1099:
	;
	v4067 = *(*int32)(unsafe.Add(mBase, uint32(v3987)+12))
	v4068 = v4067
	goto L1094
L1100:
	;
	v4071 = *(*int64)(unsafe.Add(mBase, uint32(v4068)+64))
	if v4071 == v3983 {
		goto L1081
	} else {
		goto L1103
	}
L1101:
	;
	goto L1102
L1102:
	;
	F_pfree(m, v4068)
	mBase = m.M
	v4074 = m.ExcPending
	if v4074 != 0 {
		goto L1
	} else {
		goto L1104
	}
L1103:
	;
	goto L1102
L1104:
	;
	v4076 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[40]))
	v4077 = v4076
	goto L1084
L1105:
	;
	goto L1083
L1106:
	;
	v4087 = int32(1)
	goto L1079
L1107:
	;
	m.G0 = v3987 + int32(16)
	if v4087 != 0 {
		goto L1108
	} else {
		goto L1109
	}
L1108:
	;
	v4114 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[33])) = v4114
	v4117 = *(*int64)(unsafe.Add(mBase, uint32(v21)+344))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[34])) = v4117
	v4120 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v4120 < int32(0) {
		goto L1112
	} else {
		goto L1113
	}
L1109:
	;
	goto L1110
L1110:
	;
	v4160 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v4161 = m.ExcPending
	if v4161 != 0 {
		goto L1
	} else {
		goto L1128
	}
L1111:
	;
	v4127 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v4128 = *(*int32)(unsafe.Add(mBase, uint32(v4127)+20))
	goto L1115
L1112:
	;
	v4124 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v4124
	goto L1114
L1113:
	;
	goto L1114
L1114:
	;
	goto L1111
L1115:
	;
	if base.B2i32(v4128 == int32(2)) == int32(0) {
		goto L1116
	} else {
		goto L1117
	}
L1116:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v4134 = m.ExcPending
	if v4134 != 0 {
		goto L1
	} else {
		goto L1119
	}
L1117:
	;
	goto L1118
L1118:
	;
	v4137 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v4138 = m.ExcPending
	if v4138 != 0 {
		goto L1
	} else {
		goto L1121
	}
L1119:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v4136 = m.ExcPending
	if v4136 != 0 {
		goto L1
	} else {
		goto L1120
	}
L1120:
	;
	goto L1118
L1121:
	;
	F_PushActiveSnapshot(m, v4137)
	mBase = m.M
	v4140 = m.ExcPending
	if v4140 != 0 {
		goto L1
	} else {
		goto L1122
	}
L1122:
	;
	v4143 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v4143
	F_FinishPreparedTransaction(m, v21+int32(1344), int32(0))
	mBase = m.M
	v4149 = m.ExcPending
	if v4149 != 0 {
		goto L1
	} else {
		goto L1123
	}
L1123:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v4151 = m.ExcPending
	if v4151 != 0 {
		goto L1
	} else {
		goto L1124
	}
L1124:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v4153 = m.ExcPending
	if v4153 != 0 {
		goto L1
	} else {
		goto L1125
	}
L1125:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v4155 = m.ExcPending
	if v4155 != 0 {
		goto L1
	} else {
		goto L1126
	}
L1126:
	;
	v4156 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	F_clear_subscription_skip_lsn(m, v4156)
	mBase = m.M
	v4158 = m.ExcPending
	if v4158 != 0 {
		goto L1
	} else {
		goto L1127
	}
L1127:
	;
	goto L1110
L1128:
	;
	v4162 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	v4164 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v4165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4164)+16)))
	if v4165 == int32(1) {
		goto L1130
	} else {
		goto L1131
	}
L1129:
	;
	v4205 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])) = uint8(v4205)
	F_process_syncing_tables(m, v4203)
	mBase = m.M
	v4208 = m.ExcPending
	if v4208 != 0 {
		goto L1
	} else {
		goto L1139
	}
L1130:
	;
	v4168 = *(*int32)(unsafe.Add(mBase, uint32(v4164)))
	if v4168 == int32(3) {
		v4203 = v4162
		goto L1129
	} else {
		goto L1133
	}
L1131:
	;
	goto L1132
L1132:
	;
	v4173 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v4173
	v4176 = F_palloc(m, int32(24))
	mBase = m.M
	v4177 = m.ExcPending
	if v4177 != 0 {
		goto L1
	} else {
		goto L1134
	}
L1133:
	;
	goto L1132
L1134:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4176)+16)) = v4162
	*(*int64)(unsafe.Add(mBase, uint32(v4176)+8)) = int64(0)
	v4182 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[37]))
	if v4182 != 0 {
		goto L1136
	} else {
		goto L1137
	}
L1135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4176))) = v4189
	v4191 = int32(_a_F_apply_dispatch_53)
	*(*int32)(unsafe.Add(mBase, uint32(v4176)+4)) = v4191
	*(*int32)(unsafe.Add(mBase, uint32(v4189)+4)) = v4176
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[38])) = v4176
	v4198 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v4198
	v4200 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	v4203 = v4200
	goto L1129
L1136:
	;
	v4184 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[38]))
	v4189 = v4184
	goto L1135
L1137:
	;
	goto L1138
L1138:
	;
	v4186 = int32(_a_F_apply_dispatch_53)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[37])) = v4186
	v4189 = v4186
	goto L1135
L1139:
	;
	v4210 = int32(0)
	F_pgstat_report_activity(m, int32(2), v4210)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v4210
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v4210
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v4210
	goto L3
L1140:
	;
	v4234 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v4235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4234)+16)))
	if v4235 == int32(1) {
		goto L1141
	} else {
		goto L1142
	}
L1141:
	;
	v4238 = *(*int32)(unsafe.Add(mBase, uint32(v4234)))
	if v4238 == int32(1) {
		goto L15
	} else {
		goto L1144
	}
L1142:
	;
	goto L1143
L1143:
	;
	F_logicalrep_read_prepare_common(m, l0, int32(_a_F_apply_dispatch_65), v21+int32(320))
	mBase = m.M
	v4245 = m.ExcPending
	if v4245 != 0 {
		goto L1
	} else {
		goto L1145
	}
L1144:
	;
	goto L1143
L1145:
	;
	v4247 = *(*int32)(unsafe.Add(mBase, uint32(v21)+344))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v4247
	v4250 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = v4250
	v4253 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v4254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4253)+16)))
	if v4254 == int32(1) {
		goto L1148
	} else {
		goto L1149
	}
L1146:
	;
	v4406 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v4407 = m.ExcPending
	if v4407 != 0 {
		goto L1
	} else {
		goto L1203
	}
L1147:
	;
	v4331 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24]))
	if v4331 != 0 {
		goto L1176
	} else {
		goto L1177
	}
L1148:
	;
	v4257 = *(*int32)(unsafe.Add(mBase, uint32(v4253)))
	if v4257 == int32(3) {
		goto L1147
	} else {
		goto L1151
	}
L1149:
	;
	goto L1150
L1150:
	;
	v4260 = F_pa_find_worker(m, v4247)
	mBase = m.M
	v4261 = m.ExcPending
	if v4261 != 0 {
		goto L1
	} else {
		goto L1154
	}
L1151:
	;
	goto L1150
L1152:
	;
	v4318 = *(*int32)(unsafe.Add(mBase, uint32(v21)+344))
	F_stream_open_and_write_change(m, v4318, int32(112), v21+int32(1344))
	mBase = m.M
	v4323 = m.ExcPending
	if v4323 != 0 {
		goto L1
	} else {
		goto L1173
	}
L1153:
	;
	F_pa_switch_to_partial_serialize(m, v4260, int32(1))
	mBase = m.M
	v4317 = m.ExcPending
	if v4317 != 0 {
		goto L1
	} else {
		goto L1172
	}
L1154:
	;
	if v4260 != 0 {
		goto L1155
	} else {
		goto L1156
	}
L1155:
	;
	v4262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4260)+12)))
	if v4262 != 0 {
		goto L1152
	} else {
		goto L1158
	}
L1156:
	;
	goto L1157
L1157:
	;
	v4273 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[6])))
	if v4273 != 0 {
		goto L14
	} else {
		goto L1162
	}
L1158:
	;
	v4263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4264 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4265 = F_pa_send_data(m, v4260, v4263, v4264)
	mBase = m.M
	v4266 = m.ExcPending
	if v4266 != 0 {
		goto L1
	} else {
		goto L1159
	}
L1159:
	;
	if v4265 == int32(0) {
		goto L1153
	} else {
		goto L1160
	}
L1160:
	;
	v4269 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	F_pa_xact_finish(m, v4260, v4269)
	mBase = m.M
	v4271 = m.ExcPending
	if v4271 != 0 {
		goto L1
	} else {
		goto L1161
	}
L1161:
	;
	goto L1146
L1162:
	;
	v4275 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v4276 = *(*int32)(unsafe.Add(mBase, uint32(v4275)+60))
	v4277 = *(*int32)(unsafe.Add(mBase, uint32(v21)+344))
	v4278 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	F_apply_spooled_messages(m, v4276, v4277, v4278)
	mBase = m.M
	v4280 = m.ExcPending
	if v4280 != 0 {
		goto L1
	} else {
		goto L1163
	}
L1163:
	;
	F_apply_handle_prepare_internal(m, v21+int32(320))
	mBase = m.M
	v4284 = m.ExcPending
	if v4284 != 0 {
		goto L1
	} else {
		goto L1164
	}
L1164:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v4286 = m.ExcPending
	if v4286 != 0 {
		goto L1
	} else {
		goto L1165
	}
L1165:
	;
	v4287 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	F_store_flush_position(m, v4287, int64(0))
	mBase = m.M
	v4290 = m.ExcPending
	if v4290 != 0 {
		goto L1
	} else {
		goto L1166
	}
L1166:
	;
	v4292 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_dispatch[10])) = uint8(v4292)
	v4295 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[7]))
	v4296 = *(*int32)(unsafe.Add(mBase, uint32(v4295)+32))
	v4297 = *(*int32)(unsafe.Add(mBase, uint32(v21)+344))
	F_stream_cleanup_files(m, v4296, v4297)
	mBase = m.M
	v4299 = m.ExcPending
	if v4299 != 0 {
		goto L1
	} else {
		goto L1167
	}
L1167:
	;
	v4302 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v4303 = m.ExcPending
	if v4303 != 0 {
		goto L1
	} else {
		goto L1168
	}
L1168:
	;
	if v4302 == int32(0) {
		goto L1146
	} else {
		goto L1169
	}
L1169:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_66), int32(0))
	mBase = m.M
	v4309 = m.ExcPending
	if v4309 != 0 {
		goto L1
	} else {
		goto L1170
	}
L1170:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1332), int32(_a_F_apply_dispatch_67))
	mBase = m.M
	v4314 = m.ExcPending
	if v4314 != 0 {
		goto L1
	} else {
		goto L1171
	}
L1171:
	;
	goto L1146
L1172:
	;
	goto L1152
L1173:
	;
	v4324 = *(*int32)(unsafe.Add(mBase, uint32(v4260)+16))
	F_pa_set_fileset_state(m, v4324)
	mBase = m.M
	v4326 = m.ExcPending
	if v4326 != 0 {
		goto L1
	} else {
		goto L1174
	}
L1174:
	;
	v4327 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	F_pa_xact_finish(m, v4260, v4327)
	mBase = m.M
	v4329 = m.ExcPending
	if v4329 != 0 {
		goto L1
	} else {
		goto L1175
	}
L1175:
	;
	goto L1146
L1176:
	;
	F_BufFileClose(m, v4331)
	mBase = m.M
	v4333 = m.ExcPending
	if v4333 != 0 {
		goto L1
	} else {
		goto L1179
	}
L1177:
	;
	goto L1178
L1178:
	;
	v4338 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[11]))
	if v4338 < int32(0) {
		goto L1181
	} else {
		goto L1182
	}
L1179:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[24])) = int32(0)
	goto L1178
L1180:
	;
	v4345 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[12]))
	v4346 = *(*int32)(unsafe.Add(mBase, uint32(v4345)+20))
	goto L1184
L1181:
	;
	v4342 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[13])) = v4342
	goto L1183
L1182:
	;
	goto L1183
L1183:
	;
	goto L1180
L1184:
	;
	if base.B2i32(v4346 == int32(2)) == int32(0) {
		goto L1185
	} else {
		goto L1186
	}
L1185:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v4352 = m.ExcPending
	if v4352 != 0 {
		goto L1
	} else {
		goto L1188
	}
L1186:
	;
	goto L1187
L1187:
	;
	v4355 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v4356 = m.ExcPending
	if v4356 != 0 {
		goto L1
	} else {
		goto L1190
	}
L1188:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v4354 = m.ExcPending
	if v4354 != 0 {
		goto L1
	} else {
		goto L1189
	}
L1189:
	;
	goto L1187
L1190:
	;
	F_PushActiveSnapshot(m, v4355)
	mBase = m.M
	v4358 = m.ExcPending
	if v4358 != 0 {
		goto L1
	} else {
		goto L1191
	}
L1191:
	;
	v4361 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[15])) = v4361
	F_apply_handle_prepare_internal(m, v21+int32(320))
	mBase = m.M
	v4366 = m.ExcPending
	if v4366 != 0 {
		goto L1
	} else {
		goto L1192
	}
L1192:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v4368 = m.ExcPending
	if v4368 != 0 {
		goto L1
	} else {
		goto L1193
	}
L1193:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v4370 = m.ExcPending
	if v4370 != 0 {
		goto L1
	} else {
		goto L1194
	}
L1194:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v4372 = m.ExcPending
	if v4372 != 0 {
		goto L1
	} else {
		goto L1195
	}
L1195:
	;
	v4374 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[26]))
	*(*int64)(unsafe.Add(mBase, uint32(v4374)+24)) = int64(0)
	F_pa_set_xact_state(m, v4374, int32(2))
	mBase = m.M
	v4379 = m.ExcPending
	if v4379 != 0 {
		goto L1
	} else {
		goto L1196
	}
L1196:
	;
	v4381 = *(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[26]))
	v4382 = *(*int32)(unsafe.Add(mBase, uint32(v4381)+4))
	F_pa_unlock_transaction(m, v4382)
	mBase = m.M
	v4384 = m.ExcPending
	if v4384 != 0 {
		goto L1
	} else {
		goto L1197
	}
L1197:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[35])) = int32(0)
	goto L1198
L1198:
	;
	v4390 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v4391 = m.ExcPending
	if v4391 != 0 {
		goto L1
	} else {
		goto L1199
	}
L1199:
	;
	if v4390 == int32(0) {
		goto L1146
	} else {
		goto L1200
	}
L1200:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_66), int32(0))
	mBase = m.M
	v4397 = m.ExcPending
	if v4397 != 0 {
		goto L1
	} else {
		goto L1201
	}
L1201:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1394), int32(_a_F_apply_dispatch_67))
	mBase = m.M
	v4402 = m.ExcPending
	if v4402 != 0 {
		goto L1
	} else {
		goto L1202
	}
L1202:
	;
	goto L1146
L1203:
	;
	v4408 = *(*int64)(unsafe.Add(mBase, uint32(v21)+328))
	F_process_syncing_tables(m, v4408)
	mBase = m.M
	v4410 = m.ExcPending
	if v4410 != 0 {
		goto L1
	} else {
		goto L1204
	}
L1204:
	;
	v4412 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	if v4412 != int64(0) {
		goto L1205
	} else {
		goto L1206
	}
L1205:
	;
	v4417 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4418 = m.ExcPending
	if v4418 != 0 {
		goto L1
	} else {
		goto L1208
	}
L1206:
	;
	goto L1207
L1207:
	;
	v4440 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	F_clear_subscription_skip_lsn(m, v4440)
	mBase = m.M
	v4442 = m.ExcPending
	if v4442 != 0 {
		goto L1
	} else {
		goto L1214
	}
L1208:
	;
	if v4417 != 0 {
		goto L1209
	} else {
		goto L1210
	}
L1209:
	;
	v4420 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8]))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+244)) = uint32(v4420)
	v4423 = int64(base.Ui64(v4420) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+240)) = uint32(v4423)
	F_errmsg(m, int32(_a_F_apply_dispatch_54), v21+int32(240))
	mBase = m.M
	v4429 = m.ExcPending
	if v4429 != 0 {
		goto L1
	} else {
		goto L1212
	}
L1210:
	;
	goto L1211
L1211:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[8])) = int64(0)
	goto L1207
L1212:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(_a_F_apply_dispatch_55), int32(_a_F_apply_dispatch_56))
	mBase = m.M
	v4434 = m.ExcPending
	if v4434 != 0 {
		goto L1
	} else {
		goto L1213
	}
L1213:
	;
	goto L1211
L1214:
	;
	v4444 = int32(0)
	F_pgstat_report_activity(m, int32(2), v4444)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[2])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[3])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = v4444
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[0])) = v4444
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[5])) = v4444
	goto L3
L1215:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4467 = m.ExcPending
	if v4467 != 0 {
		goto L1
	} else {
		goto L1216
	}
L1216:
	;
	v4468 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+36)) = uint32(v4468)
	v4471 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[9]))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+44)) = uint32(v4471)
	v4473 = int64(32)
	v4474 = int64(base.Ui64(v4468) >> (uint(v4473) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+32)) = uint32(v4474)
	v4477 = int64(base.Ui64(v4471) >> (uint(v4473) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+40)) = uint32(v4477)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_68), v21+int32(32))
	mBase = m.M
	v4483 = m.ExcPending
	if v4483 != 0 {
		goto L1
	} else {
		goto L1217
	}
L1217:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1021), int32(_a_F_apply_dispatch_69))
	mBase = m.M
	v4488 = m.ExcPending
	if v4488 != 0 {
		goto L1
	} else {
		goto L1218
	}
L1218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1219:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4495 = m.ExcPending
	if v4495 != 0 {
		goto L1
	} else {
		goto L1220
	}
L1220:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_70), int32(0))
	mBase = m.M
	v4499 = m.ExcPending
	if v4499 != 0 {
		goto L1
	} else {
		goto L1221
	}
L1221:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1496), int32(_a_F_apply_dispatch_71))
	mBase = m.M
	v4504 = m.ExcPending
	if v4504 != 0 {
		goto L1
	} else {
		goto L1222
	}
L1222:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1223:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4511 = m.ExcPending
	if v4511 != 0 {
		goto L1
	} else {
		goto L1224
	}
L1224:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_72), int32(0))
	mBase = m.M
	v4515 = m.ExcPending
	if v4515 != 0 {
		goto L1
	} else {
		goto L1225
	}
L1225:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1510), int32(_a_F_apply_dispatch_71))
	mBase = m.M
	v4520 = m.ExcPending
	if v4520 != 0 {
		goto L1
	} else {
		goto L1226
	}
L1226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = int32(0)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_46), v21-int32(-64))
	mBase = m.M
	v4531 = m.ExcPending
	if v4531 != 0 {
		goto L1
	} else {
		goto L1228
	}
L1228:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1606), int32(_a_F_apply_dispatch_71))
	mBase = m.M
	v4536 = m.ExcPending
	if v4536 != 0 {
		goto L1
	} else {
		goto L1229
	}
L1229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1230:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4543 = m.ExcPending
	if v4543 != 0 {
		goto L1
	} else {
		goto L1231
	}
L1231:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_73), int32(0))
	mBase = m.M
	v4547 = m.ExcPending
	if v4547 != 0 {
		goto L1
	} else {
		goto L1232
	}
L1232:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1651), int32(_a_F_apply_dispatch_31))
	mBase = m.M
	v4552 = m.ExcPending
	if v4552 != 0 {
		goto L1
	} else {
		goto L1233
	}
L1233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = int32(0)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_46), v21+int32(96))
	mBase = m.M
	v4563 = m.ExcPending
	if v4563 != 0 {
		goto L1
	} else {
		goto L1235
	}
L1235:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1722), int32(_a_F_apply_dispatch_31))
	mBase = m.M
	v4568 = m.ExcPending
	if v4568 != 0 {
		goto L1
	} else {
		goto L1236
	}
L1236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1237:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4575 = m.ExcPending
	if v4575 != 0 {
		goto L1
	} else {
		goto L1238
	}
L1238:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_74), int32(0))
	mBase = m.M
	v4579 = m.ExcPending
	if v4579 != 0 {
		goto L1
	} else {
		goto L1239
	}
L1239:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1844), int32(_a_F_apply_dispatch_40))
	mBase = m.M
	v4584 = m.ExcPending
	if v4584 != 0 {
		goto L1
	} else {
		goto L1240
	}
L1240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1241:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4591 = m.ExcPending
	if v4591 != 0 {
		goto L1
	} else {
		goto L1242
	}
L1242:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_75), int32(0))
	mBase = m.M
	v4595 = m.ExcPending
	if v4595 != 0 {
		goto L1
	} else {
		goto L1243
	}
L1243:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(2161), int32(_a_F_apply_dispatch_49))
	mBase = m.M
	v4600 = m.ExcPending
	if v4600 != 0 {
		goto L1
	} else {
		goto L1244
	}
L1244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = int32(1)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_46), v21+int32(176))
	mBase = m.M
	v4611 = m.ExcPending
	if v4611 != 0 {
		goto L1
	} else {
		goto L1246
	}
L1246:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(2242), int32(_a_F_apply_dispatch_49))
	mBase = m.M
	v4616 = m.ExcPending
	if v4616 != 0 {
		goto L1
	} else {
		goto L1247
	}
L1247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1248:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4623 = m.ExcPending
	if v4623 != 0 {
		goto L1
	} else {
		goto L1249
	}
L1249:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_76), int32(0))
	mBase = m.M
	v4627 = m.ExcPending
	if v4627 != 0 {
		goto L1
	} else {
		goto L1250
	}
L1250:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1044), int32(_a_F_apply_dispatch_77))
	mBase = m.M
	v4632 = m.ExcPending
	if v4632 != 0 {
		goto L1
	} else {
		goto L1251
	}
L1251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1252:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4639 = m.ExcPending
	if v4639 != 0 {
		goto L1
	} else {
		goto L1253
	}
L1253:
	;
	v4640 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+228)) = uint32(v4640)
	v4643 = *(*int64)(unsafe.Add(mBase, _c_F_apply_dispatch[9]))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+236)) = uint32(v4643)
	v4645 = int64(32)
	v4646 = int64(base.Ui64(v4640) >> (uint(v4645) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+224)) = uint32(v4646)
	v4649 = int64(base.Ui64(v4643) >> (uint(v4645) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+232)) = uint32(v4649)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_78), v21+int32(224))
	mBase = m.M
	v4655 = m.ExcPending
	if v4655 != 0 {
		goto L1
	} else {
		goto L1254
	}
L1254:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1113), int32(_a_F_apply_dispatch_79))
	mBase = m.M
	v4660 = m.ExcPending
	if v4660 != 0 {
		goto L1
	} else {
		goto L1255
	}
L1255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1256:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4667 = m.ExcPending
	if v4667 != 0 {
		goto L1
	} else {
		goto L1257
	}
L1257:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_80), int32(0))
	mBase = m.M
	v4671 = m.ExcPending
	if v4671 != 0 {
		goto L1
	} else {
		goto L1258
	}
L1258:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1292), int32(_a_F_apply_dispatch_67))
	mBase = m.M
	v4676 = m.ExcPending
	if v4676 != 0 {
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
	v4683 = m.ExcPending
	if v4683 != 0 {
		goto L1
	} else {
		goto L1261
	}
L1261:
	;
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_81), int32(0))
	mBase = m.M
	v4687 = m.ExcPending
	if v4687 != 0 {
		goto L1
	} else {
		goto L1262
	}
L1262:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1298), int32(_a_F_apply_dispatch_67))
	mBase = m.M
	v4692 = m.ExcPending
	if v4692 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+256)) = int32(1)
	F_errmsg_internal(m, int32(_a_F_apply_dispatch_46), v21+int32(256))
	mBase = m.M
	v4703 = m.ExcPending
	if v4703 != 0 {
		goto L1
	} else {
		goto L1265
	}
L1265:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(1398), int32(_a_F_apply_dispatch_67))
	mBase = m.M
	v4708 = m.ExcPending
	if v4708 != 0 {
		goto L1
	} else {
		goto L1266
	}
L1266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1267:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4715 = m.ExcPending
	if v4715 != 0 {
		goto L1
	} else {
		goto L1268
	}
L1268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v23
	F_errmsg(m, int32(_a_F_apply_dispatch_82), v21)
	mBase = m.M
	v4719 = m.ExcPending
	if v4719 != 0 {
		goto L1
	} else {
		goto L1269
	}
L1269:
	;
	F_errfinish(m, int32(_a_F_apply_dispatch_6), int32(3467), int32(_a_F_apply_dispatch_83))
	mBase = m.M
	v4724 = m.ExcPending
	if v4724 != 0 {
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
	goto L11
L1272:
	;
	if v2264 != v2269 {
		goto L10
	} else {
		goto L1273
	}
L1273:
	;
	v4736 = *(*int32)(unsafe.Add(mBase, uint32(v2278)+16))
	F_pa_set_fileset_state(m, v4736)
	mBase = m.M
	v4738 = m.ExcPending
	if v4738 != 0 {
		goto L1
	} else {
		goto L1274
	}
L1274:
	;
	F_pa_xact_finish(m, v2278, int64(0))
	mBase = m.M
	v4741 = m.ExcPending
	if v4741 != 0 {
		goto L1
	} else {
		goto L1275
	}
L1275:
	;
	goto L10
L1276:
	;
	if v4781 == int32(0) {
		goto L1277
	} else {
		goto L1278
	}
L1277:
	;
	if v4779 == int32(0) {
		goto L4
	} else {
		goto L1284
	}
L1278:
	;
	v4806 = int32(0)
	v4807 = *(*int32)(unsafe.Add(mBase, uint32(v4781)+4))
	if v4807 <= v4806 {
		goto L1277
	} else {
		goto L1279
	}
L1279:
	;
	v4810 = v4806
	goto L1280
L1280:
	;
	v4828 = *(*int32)(unsafe.Add(mBase, uint32(v4781)+12))
	v4832 = *(*int32)(unsafe.Add(mBase, uint32(v4828+v4810<<(uint(int32(2))%32))))
	F_logicalrep_rel_close(m, v4832, int32(0))
	mBase = m.M
	v4835 = m.ExcPending
	if v4835 != 0 {
		goto L1
	} else {
		goto L1282
	}
L1281:
	;
	goto L1277
L1282:
	;
	v4837 = v4810 + int32(1)
	v4838 = *(*int32)(unsafe.Add(mBase, uint32(v4781)+4))
	if v4837 < v4838 {
		v4810 = v4837
		goto L1280
	} else {
		goto L1283
	}
L1283:
	;
	goto L1281
L1284:
	;
	v4860 = *(*int32)(unsafe.Add(mBase, uint32(v4779)+4))
	if v4860 <= int32(0) {
		goto L4
	} else {
		goto L1285
	}
L1285:
	;
	v4864 = int32(0)
	goto L1286
L1286:
	;
	v4882 = *(*int32)(unsafe.Add(mBase, uint32(v4779)+12))
	v4886 = *(*int32)(unsafe.Add(mBase, uint32(v4882+v4864<<(uint(int32(2))%32))))
	F_relation_close(m, v4886, int32(0))
	mBase = m.M
	v4889 = m.ExcPending
	if v4889 != 0 {
		goto L1
	} else {
		goto L1288
	}
L1287:
	;
	goto L4
L1288:
	;
	v4891 = v4864 + int32(1)
	v4892 = *(*int32)(unsafe.Add(mBase, uint32(v4779)+4))
	if v4891 < v4892 {
		v4864 = v4891
		goto L1286
	} else {
		goto L1289
	}
L1289:
	;
	goto L1287
L1290:
	;
	v4899 = v21 + int32(1344)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v4899
	*(*int32)(unsafe.Add(mBase, uint32(v21)+300)) = v4899
	v4907 = F_list_make1_impl(m, int32(1), v21+int32(60))
	mBase = m.M
	v4908 = m.ExcPending
	if v4908 != 0 {
		goto L1
	} else {
		goto L1291
	}
L1291:
	;
	F_ReportApplyConflict(m, v741, v739, int32(15), int32(3), v628, v765, v4907)
	mBase = m.M
	v4910 = m.ExcPending
	if v4910 != 0 {
		goto L1
	} else {
		goto L1292
	}
L1292:
	;
	goto L7
L1293:
	;
	F_EvalPlanQualEnd(m, v21+int32(320))
	mBase = m.M
	v4918 = m.ExcPending
	if v4918 != 0 {
		goto L1
	} else {
		goto L1294
	}
L1294:
	;
	goto L6
L1295:
	;
	v4928 = *(*int32)(unsafe.Add(mBase, uint32(v622)+16))
	if v4928 != 0 {
		goto L1296
	} else {
		goto L1297
	}
L1296:
	;
	v4929 = *(*int32)(unsafe.Add(mBase, uint32(v622)+12))
	F_ExecCleanupTupleRouting(m, v4929, v4928)
	mBase = m.M
	v4931 = m.ExcPending
	if v4931 != 0 {
		goto L1
	} else {
		goto L1299
	}
L1297:
	;
	goto L1298
L1298:
	;
	v4932 = int32(0)
	v4933 = *(*int32)(unsafe.Add(mBase, uint32(v4925)+104))
	F_ExecResetTupleTable(m, v4933, v4932)
	mBase = m.M
	v4936 = m.ExcPending
	if v4936 != 0 {
		goto L1
	} else {
		goto L1300
	}
L1299:
	;
	goto L1298
L1300:
	;
	F_FreeExecutorState(m, v4925)
	mBase = m.M
	v4938 = m.ExcPending
	if v4938 != 0 {
		goto L1
	} else {
		goto L1301
	}
L1301:
	;
	F_pfree(m, v622)
	mBase = m.M
	v4940 = m.ExcPending
	if v4940 != 0 {
		goto L1
	} else {
		goto L1302
	}
L1302:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_dispatch[4])) = int32(0)
	if v612 != 0 {
		v4949 = v4932
		goto L5
	} else {
		goto L1303
	}
L1303:
	;
	F_RestoreUserContext(m, v21+int32(304))
	mBase = m.M
	v4947 = m.ExcPending
	if v4947 != 0 {
		goto L1
	} else {
		goto L1304
	}
L1304:
	;
	v4949 = v4932
	goto L5
L1305:
	;
	goto L4
L1306:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v4989 = m.ExcPending
	if v4989 != 0 {
		goto L1
	} else {
		goto L1307
	}
L1307:
	;
	goto L3
}
