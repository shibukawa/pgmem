package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_greek_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
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
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1345 int32
	_ = v1345
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1400 int32
	_ = v1400
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1567 int32
	_ = v1567
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1579 int32
	_ = v1579
	var v1595 int32
	_ = v1595
	var v1602 int32
	_ = v1602
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1732 int32
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1744 int32
	_ = v1744
	var v1760 int32
	_ = v1760
	var v1767 int32
	_ = v1767
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1779 int32
	_ = v1779
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1798 int32
	_ = v1798
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
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1851 int32
	_ = v1851
	var v1855 int32
	_ = v1855
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1896 int32
	_ = v1896
	var v1898 int32
	_ = v1898
	var v1901 int32
	_ = v1901
	var v1903 int32
	_ = v1903
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1932 int32
	_ = v1932
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1941 int32
	_ = v1941
	var v1945 int32
	_ = v1945
	var v1949 int32
	_ = v1949
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1966 int32
	_ = v1966
	var v1969 int32
	_ = v1969
	var v1973 int32
	_ = v1973
	var v1977 int32
	_ = v1977
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2009 int32
	_ = v2009
	var v2012 int32
	_ = v2012
	var v2015 int32
	_ = v2015
	var v2019 int32
	_ = v2019
	var v2022 int32
	_ = v2022
	var v2024 int32
	_ = v2024
	var v2027 int32
	_ = v2027
	var v2029 int32
	_ = v2029
	var v2032 int32
	_ = v2032
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2068 int32
	_ = v2068
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2126 int32
	_ = v2126
	var v2130 int32
	_ = v2130
	var v2132 int32
	_ = v2132
	var v2138 int32
	_ = v2138
	var v2154 int32
	_ = v2154
	var v2161 int32
	_ = v2161
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2176 int32
	_ = v2176
	var v2179 int32
	_ = v2179
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2197 int32
	_ = v2197
	var v2200 int32
	_ = v2200
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2218 int32
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2234 int32
	_ = v2234
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2248 int32
	_ = v2248
	var v2251 int32
	_ = v2251
	var v2254 int32
	_ = v2254
	var v2258 int32
	_ = v2258
	var v2261 int32
	_ = v2261
	var v2263 int32
	_ = v2263
	var v2266 int32
	_ = v2266
	var v2268 int32
	_ = v2268
	var v2271 int32
	_ = v2271
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2305 int32
	_ = v2305
	var v2307 int32
	_ = v2307
	var v2314 int32
	_ = v2314
	var v2316 int32
	_ = v2316
	var v2318 int32
	_ = v2318
	var v2320 int32
	_ = v2320
	var v2333 int32
	_ = v2333
	var v2335 int32
	_ = v2335
	var v2337 int32
	_ = v2337
	var v2355 int32
	_ = v2355
	var v2357 int32
	_ = v2357
	var v2365 int32
	_ = v2365
	var v2369 int32
	_ = v2369
	var v2371 int32
	_ = v2371
	var v2377 int32
	_ = v2377
	var v2393 int32
	_ = v2393
	var v2400 int32
	_ = v2400
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2448 int32
	_ = v2448
	var v2449 int32
	_ = v2449
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2457 int32
	_ = v2457
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2463 int32
	_ = v2463
	var v2466 int32
	_ = v2466
	var v2470 int32
	_ = v2470
	var v2474 int32
	_ = v2474
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2484 int32
	_ = v2484
	var v2486 int32
	_ = v2486
	var v2489 int32
	_ = v2489
	var v2491 int32
	_ = v2491
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2500 int32
	_ = v2500
	var v2503 int32
	_ = v2503
	var v2506 int32
	_ = v2506
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2526 int32
	_ = v2526
	var v2530 int32
	_ = v2530
	var v2531 int32
	_ = v2531
	var v2536 int32
	_ = v2536
	var v2539 int32
	_ = v2539
	var v2542 int32
	_ = v2542
	var v2546 int32
	_ = v2546
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2556 int32
	_ = v2556
	var v2559 int32
	_ = v2559
	var v2563 int32
	_ = v2563
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2569 int32
	_ = v2569
	var v2572 int32
	_ = v2572
	var v2576 int32
	_ = v2576
	var v2580 int32
	_ = v2580
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2590 int32
	_ = v2590
	var v2592 int32
	_ = v2592
	var v2595 int32
	_ = v2595
	var v2597 int32
	_ = v2597
	var v2601 int32
	_ = v2601
	var v2606 int32
	_ = v2606
	var v2609 int32
	_ = v2609
	var v2612 int32
	_ = v2612
	var v2616 int32
	_ = v2616
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2630 int32
	_ = v2630
	var v2632 int32
	_ = v2632
	var v2635 int32
	_ = v2635
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2644 int32
	_ = v2644
	var v2646 int32
	_ = v2646
	var v2649 int32
	_ = v2649
	var v2652 int32
	_ = v2652
	var v2655 int32
	_ = v2655
	var v2659 int32
	_ = v2659
	var v2662 int32
	_ = v2662
	var v2664 int32
	_ = v2664
	var v2667 int32
	_ = v2667
	var v2669 int32
	_ = v2669
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2676 int32
	_ = v2676
	var v2678 int32
	_ = v2678
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2698 int32
	_ = v2698
	var v2700 int32
	_ = v2700
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2715 int32
	_ = v2715
	var v2719 int32
	_ = v2719
	var v2722 int32
	_ = v2722
	var v2724 int32
	_ = v2724
	var v2727 int32
	_ = v2727
	var v2729 int32
	_ = v2729
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2754 int32
	_ = v2754
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2766 int32
	_ = v2766
	var v2767 int32
	_ = v2767
	var v2768 int32
	_ = v2768
	var v2770 int32
	_ = v2770
	var v2773 int32
	_ = v2773
	var v2776 int32
	_ = v2776
	var v2778 int32
	_ = v2778
	var v2781 int32
	_ = v2781
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2789 int32
	_ = v2789
	var v2791 int32
	_ = v2791
	var v2794 int32
	_ = v2794
	var v2796 int32
	_ = v2796
	var v2799 int32
	_ = v2799
	var v2803 int32
	_ = v2803
	var v2804 int32
	_ = v2804
	var v2808 int32
	_ = v2808
	var v2809 int32
	_ = v2809
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2815 int32
	_ = v2815
	var v2819 int32
	_ = v2819
	var v2821 int32
	_ = v2821
	var v2822 int32
	_ = v2822
	var v2824 int32
	_ = v2824
	var v2826 int32
	_ = v2826
	var v2832 int32
	_ = v2832
	var v2833 int32
	_ = v2833
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2846 int32
	_ = v2846
	var v2847 int32
	_ = v2847
	var v2851 int32
	_ = v2851
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2857 int32
	_ = v2857
	var v2863 int32
	_ = v2863
	var v2864 int32
	_ = v2864
	var v2867 int32
	_ = v2867
	var v2869 int32
	_ = v2869
	var v2872 int32
	_ = v2872
	var v2874 int32
	_ = v2874
	var v2877 int32
	_ = v2877
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2886 int32
	_ = v2886
	var v2887 int32
	_ = v2887
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2897 int32
	_ = v2897
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2911 int32
	_ = v2911
	var v2912 int32
	_ = v2912
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2920 int32
	_ = v2920
	var v2923 int32
	_ = v2923
	var v2924 int32
	_ = v2924
	var v2926 int32
	_ = v2926
	var v2932 int32
	_ = v2932
	var v2933 int32
	_ = v2933
	var v2936 int32
	_ = v2936
	var v2938 int32
	_ = v2938
	var v2941 int32
	_ = v2941
	var v2943 int32
	_ = v2943
	var v2948 int32
	_ = v2948
	var v2949 int32
	_ = v2949
	var v2951 int32
	_ = v2951
	var v2952 int32
	_ = v2952
	var v2954 int32
	_ = v2954
	var v2960 int32
	_ = v2960
	var v2961 int32
	_ = v2961
	var v2964 int32
	_ = v2964
	var v2965 int32
	_ = v2965
	var v2970 int32
	_ = v2970
	var v2971 int32
	_ = v2971
	var v2974 int32
	_ = v2974
	var v2976 int32
	_ = v2976
	var v2979 int32
	_ = v2979
	var v2982 int32
	_ = v2982
	var v2983 int32
	_ = v2983
	var v2985 int32
	_ = v2985
	var v2992 int32
	_ = v2992
	var v2993 int32
	_ = v2993
	var v2996 int32
	_ = v2996
	var v2998 int32
	_ = v2998
	var v3001 int32
	_ = v3001
	var v3003 int32
	_ = v3003
	var v3006 int32
	_ = v3006
	var v3007 int32
	_ = v3007
	var v3012 int32
	_ = v3012
	var v3015 int32
	_ = v3015
	var v3018 int32
	_ = v3018
	var v3022 int32
	_ = v3022
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3030 int32
	_ = v3030
	var v3031 int32
	_ = v3031
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3037 int32
	_ = v3037
	var v3041 int32
	_ = v3041
	var v3042 int32
	_ = v3042
	var v3045 int32
	_ = v3045
	var v3046 int32
	_ = v3046
	var v3049 int32
	_ = v3049
	var v3050 int32
	_ = v3050
	var v3057 int32
	_ = v3057
	var v3058 int32
	_ = v3058
	var v3061 int32
	_ = v3061
	var v3062 int32
	_ = v3062
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3071 int32
	_ = v3071
	var v3072 int32
	_ = v3072
	var v3078 int32
	_ = v3078
	var v3081 int32
	_ = v3081
	var v3082 int32
	_ = v3082
	var v3086 int32
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3092 int32
	_ = v3092
	var v3095 int32
	_ = v3095
	var v3098 int32
	_ = v3098
	var v3102 int32
	_ = v3102
	var v3105 int32
	_ = v3105
	var v3107 int32
	_ = v3107
	var v3110 int32
	_ = v3110
	var v3112 int32
	_ = v3112
	var v3119 int32
	_ = v3119
	var v3120 int32
	_ = v3120
	var v3124 int32
	_ = v3124
	var v3125 int32
	_ = v3125
	var v3130 int32
	_ = v3130
	var v3131 int32
	_ = v3131
	var v3134 int32
	_ = v3134
	var v3136 int32
	_ = v3136
	var v3139 int32
	_ = v3139
	var v3142 int32
	_ = v3142
	var v3143 int32
	_ = v3143
	var v3145 int32
	_ = v3145
	var v3148 int32
	_ = v3148
	var v3152 int32
	_ = v3152
	var v3156 int32
	_ = v3156
	var v3162 int32
	_ = v3162
	var v3163 int32
	_ = v3163
	var v3166 int32
	_ = v3166
	var v3168 int32
	_ = v3168
	var v3171 int32
	_ = v3171
	var v3173 int32
	_ = v3173
	var v3180 int32
	_ = v3180
	var v3181 int32
	_ = v3181
	var v3185 int32
	_ = v3185
	var v3186 int32
	_ = v3186
	var v3191 int32
	_ = v3191
	var v3192 int32
	_ = v3192
	var v3195 int32
	_ = v3195
	var v3197 int32
	_ = v3197
	var v3200 int32
	_ = v3200
	var v3203 int32
	_ = v3203
	var v3204 int32
	_ = v3204
	var v3206 int32
	_ = v3206
	var v3209 int32
	_ = v3209
	var v3213 int32
	_ = v3213
	var v3217 int32
	_ = v3217
	var v3223 int32
	_ = v3223
	var v3224 int32
	_ = v3224
	var v3227 int32
	_ = v3227
	var v3229 int32
	_ = v3229
	var v3232 int32
	_ = v3232
	var v3234 int32
	_ = v3234
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3252 int32
	_ = v3252
	var v3253 int32
	_ = v3253
	var v3256 int32
	_ = v3256
	var v3258 int32
	_ = v3258
	var v3261 int32
	_ = v3261
	var v3264 int32
	_ = v3264
	var v3265 int32
	_ = v3265
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3275 int32
	_ = v3275
	var v3279 int32
	_ = v3279
	var v3280 int32
	_ = v3280
	var v3284 int32
	_ = v3284
	var v3286 int32
	_ = v3286
	var v3288 int32
	_ = v3288
	var v3289 int32
	_ = v3289
	var v3296 int32
	_ = v3296
	var v3297 int32
	_ = v3297
	var v3300 int32
	_ = v3300
	var v3303 int32
	_ = v3303
	var v3306 int32
	_ = v3306
	var v3307 int32
	_ = v3307
	var v3311 int32
	_ = v3311
	var v3312 int32
	_ = v3312
	var v3314 int32
	_ = v3314
	var v3317 int32
	_ = v3317
	var v3321 int32
	_ = v3321
	var v3325 int32
	_ = v3325
	var v3331 int32
	_ = v3331
	var v3332 int32
	_ = v3332
	var v3335 int32
	_ = v3335
	var v3338 int32
	_ = v3338
	var v3341 int32
	_ = v3341
	var v3342 int32
	_ = v3342
	var v3345 int32
	_ = v3345
	var v3349 int32
	_ = v3349
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v13 = v9
	goto L2
L1:
	;
	return v3349
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v13
	v21 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_0), int32(46), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v234
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v237 = int32(0)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v236-int32(4))))
	if v244 == v237 {
		goto L102
	} else {
		goto L103
	}
L4:
	;
	goto L3
L5:
	;
	return int32(0)
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v25
	switch v21 - int32(1) {
	case 0:
		goto L32
	case 1:
		goto L31
	case 2:
		goto L30
	case 3:
		goto L29
	case 4:
		goto L28
	case 5:
		goto L27
	case 6:
		goto L26
	case 7:
		goto L25
	case 8:
		goto L24
	case 9:
		goto L23
	case 10:
		goto L22
	case 11:
		goto L21
	case 12:
		goto L20
	case 13:
		goto L19
	case 14:
		goto L18
	case 15:
		goto L17
	case 16:
		goto L16
	case 17:
		goto L15
	case 18:
		goto L14
	case 19:
		goto L13
	case 20:
		goto L12
	case 21:
		goto L11
	case 22:
		goto L10
	case 23:
		goto L9
	case 24:
		goto L8
	default:
		goto L7
	}
L7:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = v233
	goto L2
L8:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L83
L9:
	;
	v169 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_1))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L79
	}
L10:
	;
	v163 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_2))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L77
	}
L11:
	;
	v157 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_3))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L75
	}
L12:
	;
	v151 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_4))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L73
	}
L13:
	;
	v145 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_5))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L71
	}
L14:
	;
	v139 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_6))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L69
	}
L15:
	;
	v133 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_7))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L67
	}
L16:
	;
	v127 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_8))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L65
	}
L17:
	;
	v121 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_9))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L63
	}
L18:
	;
	v115 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_10))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L61
	}
L19:
	;
	v109 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_11))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L59
	}
L20:
	;
	v103 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_12))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L57
	}
L21:
	;
	v97 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_13))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L55
	}
L22:
	;
	v91 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_14))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L53
	}
L23:
	;
	v85 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_15))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L51
	}
L24:
	;
	v79 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_16))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L49
	}
L25:
	;
	v73 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_17))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L47
	}
L26:
	;
	v67 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_18))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L45
	}
L27:
	;
	v61 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_19))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L43
	}
L28:
	;
	v55 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_20))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L41
	}
L29:
	;
	v49 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_21))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L39
	}
L30:
	;
	v43 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_22))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L37
	}
L31:
	;
	v37 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_23))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L35
	}
L32:
	;
	v31 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_24))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	if int32(0) <= v31 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v3349 = v31
	goto L1
L35:
	;
	if int32(0) <= v37 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v3349 = v37
	goto L1
L37:
	;
	if int32(0) <= v43 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v3349 = v43
	goto L1
L39:
	;
	if int32(0) <= v49 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v3349 = v49
	goto L1
L41:
	;
	if int32(0) <= v55 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	v3349 = v55
	goto L1
L43:
	;
	if int32(0) <= v61 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	v3349 = v61
	goto L1
L45:
	;
	if int32(0) <= v67 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v3349 = v67
	goto L1
L47:
	;
	if int32(0) <= v73 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v3349 = v73
	goto L1
L49:
	;
	if int32(0) <= v79 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	v3349 = v79
	goto L1
L51:
	;
	if int32(0) <= v85 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	v3349 = v85
	goto L1
L53:
	;
	if int32(0) <= v91 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	v3349 = v91
	goto L1
L55:
	;
	if int32(0) <= v97 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	v3349 = v97
	goto L1
L57:
	;
	if int32(0) <= v103 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	v3349 = v103
	goto L1
L59:
	;
	if int32(0) <= v109 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v3349 = v109
	goto L1
L61:
	;
	if int32(0) <= v115 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	v3349 = v115
	goto L1
L63:
	;
	if int32(0) <= v121 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	v3349 = v121
	goto L1
L65:
	;
	if int32(0) <= v127 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v3349 = v127
	goto L1
L67:
	;
	if int32(0) <= v133 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	v3349 = v133
	goto L1
L69:
	;
	if int32(0) <= v139 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	v3349 = v139
	goto L1
L71:
	;
	if int32(0) <= v145 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	v3349 = v145
	goto L1
L73:
	;
	if int32(0) <= v151 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v3349 = v151
	goto L1
L75:
	;
	if int32(0) <= v157 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	v3349 = v157
	goto L1
L77:
	;
	if int32(0) <= v163 {
		goto L7
	} else {
		goto L78
	}
L78:
	;
	v3349 = v163
	goto L1
L79:
	;
	if int32(0) <= v169 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	v3349 = v169
	goto L1
L81:
	;
	if v227 < int32(0) {
		goto L4
	} else {
		goto L100
	}
L83:
	;
	goto L84
L84:
	;
	goto L85
L85:
	;
	v182 = v25
	v184 = int32(1)
	goto L88
L87:
	;
	v227 = v209
	goto L81
L88:
	;
	if v182 <= v175 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L87
L90:
	;
	v227 = int32(-1)
	goto L81
L91:
	;
	goto L92
L92:
	;
	v189 = v182 - int32(1)
	v191 = int32(*(*int8)(unsafe.Add(mBase, uint32(v174+v189))))
	if base.B2i32(int32(0) <= v191)|base.B2i32(v189 <= v175) != 0 {
		v209 = v189
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v213 = int32(1)
	if v213 < v184 {
		v182 = v209
		v184 = v184 - v213
		goto L88
	} else {
		goto L99
	}
L94:
	;
	v197 = v189
	goto L95
L95:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v197))))
	if base.Ui32(int32(191)) < base.Ui32(v202) {
		v209 = v197
		goto L93
	} else {
		goto L97
	}
L96:
	;
	v209 = v175
	goto L93
L97:
	;
	v206 = v197 - int32(1)
	if v175 < v206 {
		v197 = v206
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	goto L89
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v227
	goto L7
L101:
	;
	if v318 < int32(3) {
		v3349 = int32(0)
		goto L1
	} else {
		goto L117
	}
L102:
	;
	v318 = int32(0)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v249 = v244 & int32(3)
	if base.Ui32(v244) < base.Ui32(int32(4)) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v318 = v307
	goto L101
L106:
	;
	v291 = v285
	v292 = v286
	v296 = v237
	goto L114
L107:
	;
	v285 = v236
	v286 = int32(0)
	goto L106
L108:
	;
	goto L109
L109:
	;
	v256 = v236
	v257 = int32(0)
	v260 = v237
	goto L110
L110:
	;
	v262 = int32(*(*int8)(unsafe.Add(mBase, uint32(v256))))
	v263 = int32(-65)
	v266 = int32(*(*int8)(unsafe.Add(mBase, uint32(v256)+1)))
	v270 = int32(*(*int8)(unsafe.Add(mBase, uint32(v256)+2)))
	v274 = int32(*(*int8)(unsafe.Add(mBase, uint32(v256)+3)))
	v277 = v257 + base.B2i32(v263 < v262) + base.B2i32(v263 < v266) + base.B2i32(v263 < v270) + base.B2i32(v263 < v274)
	v278 = int32(4)
	v279 = v256 + v278
	v281 = v260 + v278
	if v281 != v244&int32(-4) {
		v256 = v279
		v257 = v277
		v260 = v281
		goto L110
	} else {
		goto L112
	}
L111:
	;
	if v249 == int32(0) {
		v307 = v277
		goto L105
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	v285 = v279
	v286 = v277
	goto L106
L114:
	;
	v297 = int32(*(*int8)(unsafe.Add(mBase, uint32(v291))))
	v300 = v292 + base.B2i32(int32(-65) < v297)
	v301 = int32(1)
	v304 = v296 + v301
	if v304 != v249 {
		v291 = v291 + v301
		v292 = v300
		v296 = v304
		goto L114
	} else {
		goto L116
	}
L115:
	;
	v307 = v300
	goto L105
L116:
	;
	goto L115
L117:
	;
	v321 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v321)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v323
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v329 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_25), int32(40), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	if v329 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v331
	switch v329 - int32(1) {
	case 0:
		goto L133
	case 1:
		goto L132
	case 2:
		goto L131
	case 3:
		goto L130
	case 4:
		goto L129
	case 5:
		goto L128
	case 6:
		goto L127
	case 7:
		goto L126
	case 8:
		goto L125
	case 9:
		goto L124
	case 10:
		goto L123
	default:
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v405 = v323 - v325
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v407 = v405 + v406
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v407
	v413 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_26), int32(14), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L5
	} else {
		goto L157
	}
L122:
	;
	v402 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v402)
	goto L121
L123:
	;
	v397 = F_slice_from_s(m, l0, int32(10), int32(_a_F_greek_UTF_8_stem_27))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L5
	} else {
		goto L154
	}
L124:
	;
	v391 = F_slice_from_s(m, l0, int32(12), int32(_a_F_greek_UTF_8_stem_28))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L5
	} else {
		goto L152
	}
L125:
	;
	v385 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_29))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L5
	} else {
		goto L150
	}
L126:
	;
	v379 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_30))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L148
	}
L127:
	;
	v373 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_31))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L5
	} else {
		goto L146
	}
L128:
	;
	v367 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_32))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L144
	}
L129:
	;
	v361 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_33))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L5
	} else {
		goto L142
	}
L130:
	;
	v355 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_34))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L5
	} else {
		goto L140
	}
L131:
	;
	v349 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_35))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L5
	} else {
		goto L138
	}
L132:
	;
	v343 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_36))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L5
	} else {
		goto L136
	}
L133:
	;
	v337 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_37))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L5
	} else {
		goto L134
	}
L134:
	;
	if int32(0) <= v337 {
		goto L122
	} else {
		goto L135
	}
L135:
	;
	v3349 = v337
	goto L1
L136:
	;
	if int32(0) <= v343 {
		goto L122
	} else {
		goto L137
	}
L137:
	;
	v3349 = v343
	goto L1
L138:
	;
	if int32(0) <= v349 {
		goto L122
	} else {
		goto L139
	}
L139:
	;
	v3349 = v349
	goto L1
L140:
	;
	if int32(0) <= v355 {
		goto L122
	} else {
		goto L141
	}
L141:
	;
	v3349 = v355
	goto L1
L142:
	;
	if int32(0) <= v361 {
		goto L122
	} else {
		goto L143
	}
L143:
	;
	v3349 = v361
	goto L1
L144:
	;
	if int32(0) <= v367 {
		goto L122
	} else {
		goto L145
	}
L145:
	;
	v3349 = v367
	goto L1
L146:
	;
	if int32(0) <= v373 {
		goto L122
	} else {
		goto L147
	}
L147:
	;
	v3349 = v373
	goto L1
L148:
	;
	if int32(0) <= v379 {
		goto L122
	} else {
		goto L149
	}
L149:
	;
	v3349 = v379
	goto L1
L150:
	;
	if int32(0) <= v385 {
		goto L122
	} else {
		goto L151
	}
L151:
	;
	v3349 = v385
	goto L1
L152:
	;
	if int32(0) <= v391 {
		goto L122
	} else {
		goto L153
	}
L153:
	;
	v3349 = v391
	goto L1
L154:
	;
	if v397 < int32(0) {
		v3349 = v397
		goto L1
	} else {
		goto L155
	}
L155:
	;
	goto L122
L156:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v454 = v453 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v454
	v456 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v454
	v462 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_38), int32(7), v456)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L5
	} else {
		goto L171
	}
L157:
	;
	if v413 == int32(0) {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v417
	v419 = F_slice_del(m, l0)
	mBase = m.M
	if v419 < int32(0) {
		v3349 = v419
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v422 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v422)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v424
	v430 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_39), int32(31), v422)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L5
	} else {
		goto L160
	}
L160:
	;
	if v430 == int32(0) {
		goto L156
	} else {
		goto L161
	}
L161:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v435 < v434 {
		goto L156
	} else {
		goto L162
	}
L162:
	;
	switch v430 - int32(1) {
	case 0:
		goto L164
	case 1:
		goto L163
	default:
		goto L156
	}
L163:
	;
	v447 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_40))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L5
	} else {
		goto L167
	}
L164:
	;
	v441 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_41))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L5
	} else {
		goto L165
	}
L165:
	;
	if int32(0) <= v441 {
		goto L156
	} else {
		goto L166
	}
L166:
	;
	v3349 = v441
	goto L1
L167:
	;
	if v447 < int32(0) {
		v3349 = v447
		goto L1
	} else {
		goto L168
	}
L168:
	;
	goto L156
L169:
	;
	if v500 < int32(0) {
		v3349 = v500
		goto L1
	} else {
		goto L181
	}
L170:
	;
	v500 = v497
	goto L169
L171:
	;
	if v462 == int32(0) {
		v497 = v456
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v466
	v468 = F_slice_del(m, l0)
	mBase = m.M
	if v468 < int32(0) {
		v497 = v468
		goto L170
	} else {
		goto L173
	}
L173:
	;
	v471 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v471)
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v473
	v480 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_42), int32(8), v471)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L5
	} else {
		goto L174
	}
L174:
	;
	if v480 == int32(0) {
		v500 = v471
		goto L169
	} else {
		goto L175
	}
L175:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v486 < v485 {
		v497 = int32(0)
		goto L170
	} else {
		goto L176
	}
L176:
	;
	v491 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_43))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L5
	} else {
		goto L177
	}
L177:
	;
	if int32(0) <= v491 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v495 = int32(1)
	goto L180
L179:
	;
	v495 = v491
	goto L180
L180:
	;
	v497 = v495
	goto L170
L181:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v504 = v503 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v504
	v509 = int32(6)
	v511 = int32(0)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v504-v514 < v509 {
		v524 = v511
		goto L187
	} else {
		goto L188
	}
L182:
	;
	if v598 < int32(0) {
		v3349 = v598
		goto L1
	} else {
		goto L207
	}
L183:
	;
	v598 = v593
	goto L182
L184:
	;
	v546 = int32(0)
	v550 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_44), int32(7), v546)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L5
	} else {
		goto L194
	}
L185:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v540 = v538 + (v504 - v503)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v540
	goto L184
L186:
	;
	if v524 == int32(0) {
		goto L185
	} else {
		goto L190
	}
L187:
	;
	goto L186
L188:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v520 = F_memcmp(m, v517+v504-v509, int32(_a_F_greek_UTF_8_stem_45), v509)
	mBase = m.M
	if v520 != 0 {
		v524 = v511
		goto L187
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v504 - v509
	v524 = int32(1)
	goto L187
L190:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v527
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v529 < v527 {
		goto L185
	} else {
		goto L191
	}
L191:
	;
	v533 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_46))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	if int32(0) <= v533 {
		goto L184
	} else {
		goto L193
	}
L193:
	;
	v593 = v533
	goto L183
L194:
	;
	if v550 == int32(0) {
		v593 = v546
		goto L183
	} else {
		goto L195
	}
L195:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v554
	v556 = F_slice_del(m, l0)
	mBase = m.M
	if v556 < int32(0) {
		v593 = v556
		goto L183
	} else {
		goto L196
	}
L196:
	;
	v559 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v559)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v561
	v568 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_47), int32(32), v559)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L5
	} else {
		goto L197
	}
L197:
	;
	if v568 == int32(0) {
		v598 = v559
		goto L182
	} else {
		goto L198
	}
L198:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v574 < v573 {
		v593 = int32(0)
		goto L183
	} else {
		goto L199
	}
L199:
	;
	switch v568 - int32(1) {
	case 0:
		goto L202
	case 1:
		goto L201
	default:
		goto L200
	}
L200:
	;
	v593 = int32(1)
	goto L183
L201:
	;
	v586 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_48))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L5
	} else {
		goto L205
	}
L202:
	;
	v580 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_49))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L5
	} else {
		goto L203
	}
L203:
	;
	if int32(0) <= v580 {
		goto L200
	} else {
		goto L204
	}
L204:
	;
	v593 = v580
	goto L183
L205:
	;
	if v586 < int32(0) {
		v593 = v586
		goto L183
	} else {
		goto L206
	}
L206:
	;
	goto L200
L207:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v602 = v601 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
	v604 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v602
	v610 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_50), int32(7), v604)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L5
	} else {
		goto L210
	}
L208:
	;
	if v668 < int32(0) {
		v3349 = v668
		goto L1
	} else {
		goto L223
	}
L209:
	;
	v668 = v665
	goto L208
L210:
	;
	if v610 == int32(0) {
		v665 = v604
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v614
	v616 = F_slice_del(m, l0)
	mBase = m.M
	if v616 < int32(0) {
		v665 = v616
		goto L209
	} else {
		goto L212
	}
L212:
	;
	v619 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v619)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v621
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v621
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v621-int32(3) <= v625 {
		v668 = v619
		goto L208
	} else {
		goto L213
	}
L213:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630+v621-int32(1)))))
	if v634&int32(224) != int32(160) {
		v668 = int32(0)
		goto L208
	} else {
		goto L214
	}
L214:
	;
	v639 = int32(0)
	if int32(1)<<(uint(v634)%32)&int32(-2145255424) == v639 {
		v665 = v639
		goto L209
	} else {
		goto L215
	}
L215:
	;
	v649 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_51), int32(19), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L5
	} else {
		goto L216
	}
L216:
	;
	if v649 == int32(0) {
		v665 = v639
		goto L209
	} else {
		goto L217
	}
L217:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v654 < v653 {
		v665 = v639
		goto L209
	} else {
		goto L218
	}
L218:
	;
	v659 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_52))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L5
	} else {
		goto L219
	}
L219:
	;
	if int32(0) <= v659 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v663 = int32(1)
	goto L222
L221:
	;
	v663 = v659
	goto L222
L222:
	;
	v665 = v663
	goto L209
L223:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v672 = v671 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v672
	v674 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v672
	v680 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_53), int32(11), v674)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L5
	} else {
		goto L226
	}
L224:
	;
	if v726 < int32(0) {
		v3349 = v726
		goto L1
	} else {
		goto L239
	}
L225:
	;
	v726 = v723
	goto L224
L226:
	;
	if v680 == int32(0) {
		v723 = v674
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v684
	v686 = F_slice_del(m, l0)
	mBase = m.M
	if v686 < int32(0) {
		v723 = v686
		goto L225
	} else {
		goto L228
	}
L228:
	;
	v689 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v689)
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v691
	v698 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_54), int32(40), v689)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L5
	} else {
		goto L229
	}
L229:
	;
	if v698 == int32(0) {
		v726 = v689
		goto L224
	} else {
		goto L230
	}
L230:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v704 < v703 {
		v723 = int32(0)
		goto L225
	} else {
		goto L231
	}
L231:
	;
	switch v698 - int32(1) {
	case 0:
		goto L234
	case 1:
		goto L233
	default:
		goto L232
	}
L232:
	;
	v723 = int32(1)
	goto L225
L233:
	;
	v716 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_55))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L5
	} else {
		goto L237
	}
L234:
	;
	v710 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_56))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L5
	} else {
		goto L235
	}
L235:
	;
	if int32(0) <= v710 {
		goto L232
	} else {
		goto L236
	}
L236:
	;
	v723 = v710
	goto L225
L237:
	;
	if v716 < int32(0) {
		v723 = v716
		goto L225
	} else {
		goto L238
	}
L238:
	;
	goto L232
L239:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v730 = v729 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v730
	v732 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v730
	v738 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_57), int32(6), v732)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L5
	} else {
		goto L241
	}
L240:
	;
	if v882 < int32(0) {
		v3349 = v882
		goto L1
	} else {
		goto L291
	}
L241:
	;
	if v738 == int32(0) {
		v882 = v732
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v742
	v744 = F_slice_del(m, l0)
	mBase = m.M
	if v744 < int32(0) {
		v882 = v744
		goto L240
	} else {
		goto L243
	}
L243:
	;
	v747 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v747)
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v749
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v749
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v749-int32(3) <= v753 {
		v790 = v753
		goto L245
	} else {
		goto L246
	}
L244:
	;
	v882 = v880
	goto L240
L245:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v793 = v791 + (v749 - v752)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v793
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v793
	v796 = int32(0)
	if v793-int32(9) <= v790 {
		v880 = v796
		goto L244
	} else {
		goto L257
	}
L246:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757+v749-int32(1)))))
	if v761 != int32(181) {
		v790 = v753
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v767 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_58), int32(7), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L5
	} else {
		goto L248
	}
L248:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v767 == int32(0) {
		v790 = v769
		goto L245
	} else {
		goto L249
	}
L249:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v769 < v772 {
		v790 = v769
		goto L245
	} else {
		goto L250
	}
L250:
	;
	v774 = int32(1)
	switch v767 - v774 {
	case 0:
		goto L252
	case 1:
		goto L251
	default:
		v882 = v774
		goto L240
	}
L251:
	;
	v785 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_59))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L5
	} else {
		goto L255
	}
L252:
	;
	v779 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_60))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L5
	} else {
		goto L253
	}
L253:
	;
	if v779 < int32(0) {
		v880 = v779
		goto L244
	} else {
		goto L254
	}
L254:
	;
	v882 = v774
	goto L240
L255:
	;
	if v785 < int32(0) {
		v880 = v785
		goto L244
	} else {
		goto L256
	}
L256:
	;
	v882 = v774
	goto L240
L257:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800+v793-int32(1)))))
	switch v804 - int32(186) {
	case 0, 3:
		goto L258
	default:
		v880 = v796
		goto L244
	}
L258:
	;
	v810 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_61), int32(10), int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L5
	} else {
		goto L259
	}
L259:
	;
	if v810 == int32(0) {
		v880 = v796
		goto L244
	} else {
		goto L260
	}
L260:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v814
	v816 = int32(1)
	switch v810 - v816 {
	case 0:
		goto L270
	case 1:
		goto L269
	case 2:
		goto L268
	case 3:
		goto L267
	case 4:
		goto L266
	case 5:
		goto L265
	case 6:
		goto L264
	case 7:
		goto L263
	case 8:
		goto L262
	case 9:
		goto L261
	default:
		v882 = v816
		goto L240
	}
L261:
	;
	v875 = F_slice_from_s(m, l0, int32(10), int32(_a_F_greek_UTF_8_stem_62))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L5
	} else {
		goto L289
	}
L262:
	;
	v869 = F_slice_from_s(m, l0, int32(12), int32(_a_F_greek_UTF_8_stem_63))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L5
	} else {
		goto L287
	}
L263:
	;
	v863 = F_slice_from_s(m, l0, int32(16), int32(_a_F_greek_UTF_8_stem_64))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L5
	} else {
		goto L285
	}
L264:
	;
	v857 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_65))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L5
	} else {
		goto L283
	}
L265:
	;
	v851 = F_slice_from_s(m, l0, int32(10), int32(_a_F_greek_UTF_8_stem_66))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L5
	} else {
		goto L281
	}
L266:
	;
	v845 = F_slice_from_s(m, l0, int32(12), int32(_a_F_greek_UTF_8_stem_67))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L5
	} else {
		goto L279
	}
L267:
	;
	v839 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_68))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L5
	} else {
		goto L277
	}
L268:
	;
	v833 = F_slice_from_s(m, l0, int32(10), int32(_a_F_greek_UTF_8_stem_69))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L5
	} else {
		goto L275
	}
L269:
	;
	v827 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_70))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L5
	} else {
		goto L273
	}
L270:
	;
	v821 = F_slice_from_s(m, l0, int32(12), int32(_a_F_greek_UTF_8_stem_71))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L5
	} else {
		goto L271
	}
L271:
	;
	if v821 < int32(0) {
		v880 = v821
		goto L244
	} else {
		goto L272
	}
L272:
	;
	v882 = v816
	goto L240
L273:
	;
	if v827 < int32(0) {
		v880 = v827
		goto L244
	} else {
		goto L274
	}
L274:
	;
	v882 = v816
	goto L240
L275:
	;
	if v833 < int32(0) {
		v880 = v833
		goto L244
	} else {
		goto L276
	}
L276:
	;
	v882 = v816
	goto L240
L277:
	;
	if v839 < int32(0) {
		v880 = v839
		goto L244
	} else {
		goto L278
	}
L278:
	;
	v882 = v816
	goto L240
L279:
	;
	if v845 < int32(0) {
		v880 = v845
		goto L244
	} else {
		goto L280
	}
L280:
	;
	v882 = v816
	goto L240
L281:
	;
	if v851 < int32(0) {
		v880 = v851
		goto L244
	} else {
		goto L282
	}
L282:
	;
	v882 = v816
	goto L240
L283:
	;
	if v857 < int32(0) {
		v880 = v857
		goto L244
	} else {
		goto L284
	}
L284:
	;
	v882 = v816
	goto L240
L285:
	;
	if v863 < int32(0) {
		v880 = v863
		goto L244
	} else {
		goto L286
	}
L286:
	;
	v882 = v816
	goto L240
L287:
	;
	if v869 < int32(0) {
		v880 = v869
		goto L244
	} else {
		goto L288
	}
L288:
	;
	v882 = v816
	goto L240
L289:
	;
	if int32(0) <= v875 {
		v882 = v816
		goto L240
	} else {
		goto L290
	}
L290:
	;
	v880 = v875
	goto L244
L291:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v889 = v888 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v889
	v891 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v889
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v889-int32(9) <= v894 {
		v951 = v891
		goto L293
	} else {
		goto L294
	}
L292:
	;
	if v955 < int32(0) {
		v3349 = v955
		goto L1
	} else {
		goto L308
	}
L293:
	;
	v955 = v951
	goto L292
L294:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898+v889-int32(1)))))
	switch v902 - int32(177) {
	case 0, 8:
		goto L295
	default:
		v951 = v891
		goto L293
	}
L295:
	;
	v908 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_72), int32(4), int32(0))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L5
	} else {
		goto L296
	}
L296:
	;
	if v908 == int32(0) {
		v951 = v891
		goto L293
	} else {
		goto L297
	}
L297:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v912
	v914 = F_slice_del(m, l0)
	mBase = m.M
	if v914 < int32(0) {
		v951 = v914
		goto L293
	} else {
		goto L298
	}
L298:
	;
	v917 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v917)
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v919
	v924 = v919 - int32(1)
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v924 <= v925 {
		v955 = v917
		goto L292
	} else {
		goto L299
	}
L299:
	;
	v927 = int32(0)
	v928 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v928+v924))))
	switch v930 - int32(131) {
	case 0, 4:
		goto L300
	default:
		v951 = v927
		goto L293
	}
L300:
	;
	v936 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_73), int32(2), int32(0))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L5
	} else {
		goto L301
	}
L301:
	;
	if v936 == int32(0) {
		v951 = v927
		goto L293
	} else {
		goto L302
	}
L302:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v941 < v940 {
		v951 = v927
		goto L293
	} else {
		goto L303
	}
L303:
	;
	v946 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_74))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L5
	} else {
		goto L304
	}
L304:
	;
	if int32(0) <= v946 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v950 = int32(1)
	goto L307
L306:
	;
	v950 = v946
	goto L307
L307:
	;
	v951 = v950
	goto L293
L308:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v959 = v958 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v959
	v961 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v959
	v967 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_75), int32(8), v961)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L5
	} else {
		goto L310
	}
L309:
	;
	if v1041 < int32(0) {
		v3349 = v1041
		goto L1
	} else {
		goto L333
	}
L310:
	;
	if v967 == int32(0) {
		v1041 = v961
		goto L309
	} else {
		goto L311
	}
L311:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v971
	v973 = F_slice_del(m, l0)
	mBase = m.M
	if v973 < int32(0) {
		v1041 = v973
		goto L309
	} else {
		goto L312
	}
L312:
	;
	v976 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v976)
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v978
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v978
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v985 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_76), int32(46), v976)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L5
	} else {
		goto L315
	}
L313:
	;
	v1041 = v1040
	goto L309
L314:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1009 = v1007 + (v978 - v981)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1009
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1009
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1009
	v1013 = int32(6)
	v1015 = int32(0)
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1009-v1018 < v1013 {
		v1028 = v1015
		goto L325
	} else {
		goto L326
	}
L315:
	;
	if v985 == int32(0) {
		goto L314
	} else {
		goto L316
	}
L316:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v990 < v989 {
		goto L314
	} else {
		goto L317
	}
L317:
	;
	v992 = int32(1)
	switch v985 - v992 {
	case 0:
		goto L319
	case 1:
		goto L318
	default:
		v1041 = v992
		goto L309
	}
L318:
	;
	v1003 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_77))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L5
	} else {
		goto L322
	}
L319:
	;
	v997 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_78))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L5
	} else {
		goto L320
	}
L320:
	;
	if v997 < int32(0) {
		v1040 = v997
		goto L313
	} else {
		goto L321
	}
L321:
	;
	v1041 = v992
	goto L309
L322:
	;
	if v1003 < int32(0) {
		v1040 = v1003
		goto L313
	} else {
		goto L323
	}
L323:
	;
	v1041 = v992
	goto L309
L324:
	;
	if v1028 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L325:
	;
	goto L324
L326:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1024 = F_memcmp(m, v1021+v1009-v1013, int32(_a_F_greek_UTF_8_stem_79), v1013)
	mBase = m.M
	if v1024 != 0 {
		v1028 = v1015
		goto L325
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1009 - v1013
	v1028 = int32(1)
	goto L325
L328:
	;
	v1040 = int32(0)
	goto L313
L329:
	;
	goto L330
L330:
	;
	v1035 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_80))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L5
	} else {
		goto L331
	}
L331:
	;
	if int32(0) <= v1035 {
		v1041 = int32(1)
		goto L309
	} else {
		goto L332
	}
L332:
	;
	v1040 = v1035
	goto L313
L333:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1047 = v1046 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1047
	v1049 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1047
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1047-int32(7) <= v1052 {
		v1140 = v1049
		goto L335
	} else {
		goto L336
	}
L334:
	;
	if v1146 < int32(0) {
		v3349 = v1146
		goto L1
	} else {
		goto L357
	}
L335:
	;
	v1146 = v1140
	goto L334
L336:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1058 = int32(1)
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056+v1047-v1058))))
	if base.B2i32(v1060&int32(224) != int32(160))|base.B2i32(v1058<<(uint(v1060)%32)&int32(-1610481664) == int32(0)) != 0 {
		v1140 = v1049
		goto L335
	} else {
		goto L337
	}
L337:
	;
	v1075 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_81), int32(3), int32(0))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L5
	} else {
		goto L338
	}
L338:
	;
	if v1075 == int32(0) {
		v1140 = v1049
		goto L335
	} else {
		goto L339
	}
L339:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1079
	v1081 = F_slice_del(m, l0)
	mBase = m.M
	if v1081 < int32(0) {
		v1140 = v1081
		goto L335
	} else {
		goto L340
	}
L340:
	;
	v1084 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v1084)
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1086
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1093 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_82), int32(4), v1084)
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L5
	} else {
		goto L341
	}
L341:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1093 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1110 = v1108 + (v1086 - v1089)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1110
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1110
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1110
	v1114 = int32(0)
	v1116 = v1110 - int32(1)
	if v1116 <= v1095 {
		v1140 = v1114
		goto L335
	} else {
		goto L349
	}
L343:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1095 < v1098 {
		goto L342
	} else {
		goto L344
	}
L344:
	;
	v1103 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_83))
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L5
	} else {
		goto L345
	}
L345:
	;
	if int32(0) <= v1103 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1107 = int32(1)
	goto L348
L347:
	;
	v1107 = v1103
	goto L348
L348:
	;
	v1146 = v1107
	goto L334
L349:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1118+v1116))))
	switch v1120 - int32(181) {
	case 0, 8:
		goto L350
	default:
		v1140 = v1114
		goto L335
	}
L350:
	;
	v1126 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_84), int32(2), int32(0))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L5
	} else {
		goto L351
	}
L351:
	;
	if v1126 == int32(0) {
		v1140 = v1114
		goto L335
	} else {
		goto L352
	}
L352:
	;
	v1133 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_85))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L5
	} else {
		goto L353
	}
L353:
	;
	if int32(0) <= v1133 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1137 = int32(1)
	goto L356
L355:
	;
	v1137 = v1133
	goto L356
L356:
	;
	v1140 = v1137
	goto L335
L357:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1150 = v1149 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1150
	v1152 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1150
	v1158 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_86), int32(4), v1152)
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L5
	} else {
		goto L360
	}
L358:
	;
	if v1196 < int32(0) {
		v3349 = v1196
		goto L1
	} else {
		goto L370
	}
L359:
	;
	v1196 = v1193
	goto L358
L360:
	;
	if v1158 == int32(0) {
		v1193 = v1152
		goto L359
	} else {
		goto L361
	}
L361:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1162
	v1164 = F_slice_del(m, l0)
	mBase = m.M
	if v1164 < int32(0) {
		v1193 = v1164
		goto L359
	} else {
		goto L362
	}
L362:
	;
	v1167 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v1167)
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1169
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1169
	v1176 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_87), int32(7), v1167)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L5
	} else {
		goto L363
	}
L363:
	;
	if v1176 == int32(0) {
		v1196 = v1167
		goto L358
	} else {
		goto L364
	}
L364:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1182 < v1181 {
		v1193 = int32(0)
		goto L359
	} else {
		goto L365
	}
L365:
	;
	v1187 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_88))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L5
	} else {
		goto L366
	}
L366:
	;
	if int32(0) <= v1187 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1191 = int32(1)
	goto L369
L368:
	;
	v1191 = v1187
	goto L369
L369:
	;
	v1193 = v1191
	goto L359
L370:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1200 = v1199 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1200
	v1202 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1200
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1200-int32(7) <= v1205 {
		v1252 = v1202
		goto L371
	} else {
		goto L372
	}
L371:
	;
	if v1252 < int32(0) {
		v3349 = v1252
		goto L1
	} else {
		goto L383
	}
L372:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209+v1200-int32(1)))))
	if base.B2i32(v1213 != int32(189))&base.B2i32(v1213 != int32(131)) != 0 {
		v1252 = v1202
		goto L371
	} else {
		goto L373
	}
L373:
	;
	v1222 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_89), int32(2), int32(0))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L5
	} else {
		goto L374
	}
L374:
	;
	if v1222 == int32(0) {
		v1252 = v1202
		goto L371
	} else {
		goto L375
	}
L375:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1226
	v1228 = F_slice_del(m, l0)
	mBase = m.M
	if v1228 < int32(0) {
		v1252 = v1228
		goto L371
	} else {
		goto L376
	}
L376:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1233 = int32(0)
	v1237 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_90), int32(10), v1233)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L5
	} else {
		goto L377
	}
L377:
	;
	if v1237 != 0 {
		v1252 = v1233
		goto L371
	} else {
		goto L378
	}
L378:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1241 = v1239 + (v1231 - v1232)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1241
	v1245 = F_insert_s(m, l0, v1241, v1241, int32(4), int32(_a_F_greek_UTF_8_stem_91))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L5
	} else {
		goto L379
	}
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1241
	if int32(0) <= v1245 {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v1251 = int32(1)
	goto L382
L381:
	;
	v1251 = v1245
	goto L382
L382:
	;
	v1252 = v1251
	goto L371
L383:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1258 = v1257 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1258
	v1260 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1258
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1258-int32(7) <= v1263 {
		v1324 = v1260
		goto L385
	} else {
		goto L386
	}
L384:
	;
	if v1328 < int32(0) {
		v3349 = v1328
		goto L1
	} else {
		goto L402
	}
L385:
	;
	v1328 = v1324
	goto L384
L386:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1267+v1258-int32(1)))))
	if base.B2i32(v1271 != int32(189))&base.B2i32(v1271 != int32(131)) != 0 {
		v1324 = v1260
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v1280 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_92), int32(2), int32(0))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L5
	} else {
		goto L388
	}
L388:
	;
	if v1280 == int32(0) {
		v1324 = v1260
		goto L385
	} else {
		goto L389
	}
L389:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1284
	v1286 = F_slice_del(m, l0)
	mBase = m.M
	if v1286 < int32(0) {
		v1324 = v1286
		goto L385
	} else {
		goto L390
	}
L390:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1289
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1289-int32(3) <= v1293 {
		v1328 = int32(0)
		goto L384
	} else {
		goto L391
	}
L391:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1297+v1289-int32(1)))))
	if v1301 != int32(187) {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	if v1301 != int32(128) {
		v1324 = int32(0)
		goto L385
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	v1308 = int32(0)
	v1312 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_93), int32(8), v1308)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L5
	} else {
		goto L396
	}
L395:
	;
	goto L394
L396:
	;
	if v1312 == int32(0) {
		v1324 = v1308
		goto L385
	} else {
		goto L397
	}
L397:
	;
	v1319 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_94))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L5
	} else {
		goto L398
	}
L398:
	;
	if int32(0) <= v1319 {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v1323 = int32(1)
	goto L401
L400:
	;
	v1323 = v1319
	goto L401
L401:
	;
	v1324 = v1323
	goto L385
L402:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1332 = v1331 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1332
	v1334 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1332
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1332-int32(9) <= v1337 {
		v1383 = v1334
		goto L403
	} else {
		goto L404
	}
L403:
	;
	if v1383 < int32(0) {
		v3349 = v1383
		goto L1
	} else {
		goto L415
	}
L404:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1341+v1332-int32(1)))))
	if base.B2i32(v1345 != int32(189))&base.B2i32(v1345 != int32(131)) != 0 {
		v1383 = v1334
		goto L403
	} else {
		goto L405
	}
L405:
	;
	v1354 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_95), int32(2), int32(0))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L5
	} else {
		goto L406
	}
L406:
	;
	if v1354 == int32(0) {
		v1383 = v1334
		goto L403
	} else {
		goto L407
	}
L407:
	;
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1358
	v1360 = F_slice_del(m, l0)
	mBase = m.M
	if v1360 < int32(0) {
		v1383 = v1360
		goto L403
	} else {
		goto L408
	}
L408:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1363
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1363
	v1366 = int32(0)
	v1370 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_96), int32(15), v1366)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L5
	} else {
		goto L409
	}
L409:
	;
	if v1370 == int32(0) {
		v1383 = v1366
		goto L403
	} else {
		goto L410
	}
L410:
	;
	v1377 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_97))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L5
	} else {
		goto L411
	}
L411:
	;
	if int32(0) <= v1377 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v1381 = int32(1)
	goto L414
L413:
	;
	v1381 = v1377
	goto L414
L414:
	;
	v1383 = v1381
	goto L403
L415:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1387 = v1386 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1387
	v1389 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1387
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1387-int32(5) <= v1392 {
		v1444 = v1389
		goto L417
	} else {
		goto L418
	}
L416:
	;
	if v1447 < int32(0) {
		v3349 = v1447
		goto L1
	} else {
		goto L430
	}
L417:
	;
	v1447 = v1444
	goto L416
L418:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1396+v1387-int32(1)))))
	if base.B2i32(v1400 != int32(189))&base.B2i32(v1400 != int32(131)) != 0 {
		v1444 = v1389
		goto L417
	} else {
		goto L419
	}
L419:
	;
	v1409 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_98), int32(2), int32(0))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L5
	} else {
		goto L420
	}
L420:
	;
	if v1409 == int32(0) {
		v1444 = v1389
		goto L417
	} else {
		goto L421
	}
L421:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1413
	v1415 = F_slice_del(m, l0)
	mBase = m.M
	if v1415 < int32(0) {
		v1444 = v1415
		goto L417
	} else {
		goto L422
	}
L422:
	;
	v1418 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v1418)
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1420
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1420
	v1427 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_99), int32(8), v1418)
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L5
	} else {
		goto L423
	}
L423:
	;
	if v1427 == int32(0) {
		v1447 = v1418
		goto L416
	} else {
		goto L424
	}
L424:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1433 < v1432 {
		v1444 = int32(0)
		goto L417
	} else {
		goto L425
	}
L425:
	;
	v1438 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_100))
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L5
	} else {
		goto L426
	}
L426:
	;
	if int32(0) <= v1438 {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v1442 = int32(1)
	goto L429
L428:
	;
	v1442 = v1438
	goto L429
L429:
	;
	v1444 = v1442
	goto L417
L430:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1451 = v1450 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1451
	v1453 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1451
	v1459 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_101), int32(3), v1453)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L5
	} else {
		goto L432
	}
L431:
	;
	if v1612 < int32(0) {
		v3349 = v1612
		goto L1
	} else {
		goto L463
	}
L432:
	;
	if v1459 == int32(0) {
		v1612 = v1453
		goto L431
	} else {
		goto L433
	}
L433:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1463
	v1465 = F_slice_del(m, l0)
	mBase = m.M
	if v1465 < int32(0) {
		v1612 = v1465
		goto L431
	} else {
		goto L434
	}
L434:
	;
	v1468 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v1468)
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1470
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L437
L435:
	;
	if v1602 != 0 {
		v1612 = v1468
		goto L431
	} else {
		goto L458
	}
L436:
	;
	v1602 = v1595
	goto L435
L437:
	;
	if v1470 <= v1487 {
		v1595 = int32(-1)
		goto L436
	} else {
		goto L439
	}
L438:
	;
	v1595 = int32(0)
	goto L436
L439:
	;
	v1504 = int32(1)
	v1505 = v1470 - v1504
	v1507 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1488+v1505))))
	v1509 = v1507 & int32(255)
	if base.B2i32(v1505 == v1487)|base.B2i32(int32(0) <= v1507) != 0 {
		v1567 = v1509
		v1571 = v1504
		goto L440
	} else {
		goto L441
	}
L440:
	;
	if int32(969) < v1567 {
		goto L448
	} else {
		goto L449
	}
L441:
	;
	v1516 = v1509 & int32(63)
	v1518 = v1470 - int32(2)
	v1520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1488+v1518))))
	v1522 = v1520 << (uint(int32(6)) % 32)
	if base.B2i32(v1518 != v1487)&base.B2i32(base.Ui32(v1520) < base.Ui32(int32(192))) == int32(0) {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v1567 = v1522&int32(1984) | v1516
	v1571 = int32(2)
	goto L440
L443:
	;
	goto L444
L444:
	;
	v1535 = v1522&int32(4032) | v1516
	v1537 = v1470 - int32(3)
	v1539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1488+v1537))))
	if base.B2i32(v1537 != v1487)&base.B2i32(base.Ui32(v1539) < base.Ui32(int32(224))) == int32(0) {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v1567 = v1539<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_102) | v1535
	v1571 = int32(3)
	goto L440
L446:
	;
	goto L447
L447:
	;
	v1557 = int32(4)
	v1559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1470+v1488-v1557))))
	v1567 = v1539<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_103) | v1559&int32(7)<<(uint(int32(18))%32) | v1535
	v1571 = v1557
	goto L440
L448:
	;
	v1602 = v1571
	goto L435
L449:
	;
	goto L450
L450:
	;
	v1573 = v1567 - int32(945)
	if v1573 < int32(0) {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v1602 = v1571
	goto L435
L452:
	;
	goto L453
L453:
	;
	v1579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1573)>>(uint(int32(3))%32)))+uint32(_c_F_greek_UTF_8_stem[0]))))
	if int32(base.Ui32(v1579)>>(uint(v1573&int32(7))%32))&int32(1) == int32(0) {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v1602 = v1571
	goto L435
L455:
	;
	goto L456
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1470 - v1571
	goto L457
L457:
	;
	goto L438
L458:
	;
	v1606 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_104))
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L5
	} else {
		goto L459
	}
L459:
	;
	if int32(0) <= v1606 {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v1610 = int32(1)
	goto L462
L461:
	;
	v1610 = v1606
	goto L462
L462:
	;
	v1612 = v1610
	goto L431
L463:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1616 = v1615 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1616
	v1618 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1616
	v1624 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_105), int32(4), v1618)
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L5
	} else {
		goto L465
	}
L464:
	;
	if v1803 < int32(0) {
		v3349 = v1803
		goto L1
	} else {
		goto L504
	}
L465:
	;
	if v1624 == int32(0) {
		v1803 = v1618
		goto L464
	} else {
		goto L466
	}
L466:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1628
	v1630 = F_slice_del(m, l0)
	mBase = m.M
	if v1630 < int32(0) {
		v1803 = v1630
		goto L464
	} else {
		goto L467
	}
L467:
	;
	v1633 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v1633)
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1635
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L471
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1782
	v1784 = int32(0)
	v1788 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_106), int32(36), v1784)
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L5
	} else {
		goto L497
	}
L469:
	;
	if v1767 == int32(0) {
		goto L492
	} else {
		goto L493
	}
L470:
	;
	v1767 = v1760
	goto L469
L471:
	;
	if v1635 <= v1652 {
		v1760 = int32(-1)
		goto L470
	} else {
		goto L473
	}
L472:
	;
	v1760 = int32(0)
	goto L470
L473:
	;
	v1669 = int32(1)
	v1670 = v1635 - v1669
	v1672 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1653+v1670))))
	v1674 = v1672 & int32(255)
	if base.B2i32(v1670 == v1652)|base.B2i32(int32(0) <= v1672) != 0 {
		v1732 = v1674
		v1736 = v1669
		goto L474
	} else {
		goto L475
	}
L474:
	;
	if int32(969) < v1732 {
		goto L482
	} else {
		goto L483
	}
L475:
	;
	v1681 = v1674 & int32(63)
	v1683 = v1635 - int32(2)
	v1685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1653+v1683))))
	v1687 = v1685 << (uint(int32(6)) % 32)
	if base.B2i32(v1683 != v1652)&base.B2i32(base.Ui32(v1685) < base.Ui32(int32(192))) == int32(0) {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v1732 = v1687&int32(1984) | v1681
	v1736 = int32(2)
	goto L474
L477:
	;
	goto L478
L478:
	;
	v1700 = v1687&int32(4032) | v1681
	v1702 = v1635 - int32(3)
	v1704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1653+v1702))))
	if base.B2i32(v1702 != v1652)&base.B2i32(base.Ui32(v1704) < base.Ui32(int32(224))) == int32(0) {
		goto L479
	} else {
		goto L480
	}
L479:
	;
	v1732 = v1704<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_102) | v1700
	v1736 = int32(3)
	goto L474
L480:
	;
	goto L481
L481:
	;
	v1722 = int32(4)
	v1724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1635+v1653-v1722))))
	v1732 = v1704<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_103) | v1724&int32(7)<<(uint(int32(18))%32) | v1700
	v1736 = v1722
	goto L474
L482:
	;
	v1767 = v1736
	goto L469
L483:
	;
	goto L484
L484:
	;
	v1738 = v1732 - int32(945)
	if v1738 < int32(0) {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v1767 = v1736
	goto L469
L486:
	;
	goto L487
L487:
	;
	v1744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1738)>>(uint(int32(3))%32)))+uint32(_c_F_greek_UTF_8_stem[0]))))
	if int32(base.Ui32(v1744)>>(uint(v1738&int32(7))%32))&int32(1) == int32(0) {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v1767 = v1736
	goto L469
L489:
	;
	goto L490
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1635 - v1736
	goto L491
L491:
	;
	goto L472
L492:
	;
	v1772 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_107))
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L5
	} else {
		goto L495
	}
L493:
	;
	goto L494
L494:
	;
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1779 = v1777 + (v1635 - v1638)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1779
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1779
	v1782 = v1779
	goto L468
L495:
	;
	if v1772 < int32(0) {
		v1803 = v1772
		goto L464
	} else {
		goto L496
	}
L496:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1782 = v1776
	goto L468
L497:
	;
	if v1788 == int32(0) {
		v1803 = v1784
		goto L464
	} else {
		goto L498
	}
L498:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1793 < v1792 {
		v1803 = v1784
		goto L464
	} else {
		goto L499
	}
L499:
	;
	v1798 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_108))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L5
	} else {
		goto L500
	}
L500:
	;
	if int32(0) <= v1798 {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	v1802 = int32(1)
	goto L503
L502:
	;
	v1802 = v1798
	goto L503
L503:
	;
	v1803 = v1802
	goto L464
L504:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1808 = v1807 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1808
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1808
	v1813 = int32(10)
	v1815 = int32(0)
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1808-v1818 < v1813 {
		v1828 = v1815
		goto L509
	} else {
		goto L510
	}
L505:
	;
	if v1932 < int32(0) {
		v3349 = v1932
		goto L1
	} else {
		goto L535
	}
L506:
	;
	v1932 = v1928
	goto L505
L507:
	;
	v1842 = v1808 - v1807
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1844 = v1842 + v1843
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1844
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1844
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1844-int32(9) <= v1847 {
		goto L516
	} else {
		goto L517
	}
L508:
	;
	if v1828 == int32(0) {
		goto L507
	} else {
		goto L512
	}
L509:
	;
	goto L508
L510:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1824 = F_memcmp(m, v1821+v1808-v1813, int32(_a_F_greek_UTF_8_stem_109), v1813)
	mBase = m.M
	if v1824 != 0 {
		v1828 = v1815
		goto L509
	} else {
		goto L511
	}
L511:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1808 - v1813
	v1828 = int32(1)
	goto L509
L512:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1831
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1833 < v1831 {
		goto L507
	} else {
		goto L513
	}
L513:
	;
	v1837 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_110))
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L5
	} else {
		goto L514
	}
L514:
	;
	if v1837 < int32(0) {
		v1928 = v1837
		goto L506
	} else {
		goto L515
	}
L515:
	;
	goto L507
L516:
	;
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1874 = v1873 + v1842
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1874
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1874
	v1877 = int32(0)
	v1878 = int32(6)
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1874-v1883 < v1878 {
		v1893 = v1877
		goto L523
	} else {
		goto L524
	}
L517:
	;
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1851+v1844-int32(1)))))
	if v1855 != int32(181) {
		goto L516
	} else {
		goto L518
	}
L518:
	;
	v1861 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_111), int32(5), int32(0))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L5
	} else {
		goto L519
	}
L519:
	;
	if v1861 == int32(0) {
		goto L516
	} else {
		goto L520
	}
L520:
	;
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1865
	v1867 = F_slice_del(m, l0)
	mBase = m.M
	if v1867 < int32(0) {
		v1928 = v1867
		goto L506
	} else {
		goto L521
	}
L521:
	;
	v1870 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v1870)
	goto L516
L522:
	;
	if v1893 == int32(0) {
		v1932 = v1877
		goto L505
	} else {
		goto L526
	}
L523:
	;
	goto L522
L524:
	;
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1889 = F_memcmp(m, v1886+v1874-v1878, int32(_a_F_greek_UTF_8_stem_112), v1878)
	mBase = m.M
	if v1889 != 0 {
		v1893 = v1877
		goto L523
	} else {
		goto L525
	}
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1874 - v1878
	v1893 = int32(1)
	goto L523
L526:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1896
	v1898 = F_slice_del(m, l0)
	mBase = m.M
	if v1898 < int32(0) {
		v1928 = v1898
		goto L506
	} else {
		goto L527
	}
L527:
	;
	v1901 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v1901)
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1903
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1903
	v1910 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_113), int32(12), v1901)
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L5
	} else {
		goto L528
	}
L528:
	;
	if v1910 == int32(0) {
		v1932 = v1901
		goto L505
	} else {
		goto L529
	}
L529:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1916 < v1915 {
		v1928 = int32(0)
		goto L506
	} else {
		goto L530
	}
L530:
	;
	v1921 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_114))
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L5
	} else {
		goto L531
	}
L531:
	;
	if int32(0) <= v1921 {
		goto L532
	} else {
		goto L533
	}
L532:
	;
	v1925 = int32(1)
	goto L534
L533:
	;
	v1925 = v1921
	goto L534
L534:
	;
	v1928 = v1925
	goto L506
L535:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1936 = v1935 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1936
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1936
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1936-int32(9) <= v1941 {
		goto L538
	} else {
		goto L539
	}
L536:
	;
	if v2204 < int32(0) {
		v3349 = v2204
		goto L1
	} else {
		goto L593
	}
L537:
	;
	v2204 = v2200
	goto L536
L538:
	;
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2000 = v1998 + (v1936 - v1935)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2000
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2000
	v2003 = int32(0)
	v2004 = int32(6)
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2000-v2009 < v2004 {
		v2019 = v2003
		goto L552
	} else {
		goto L553
	}
L539:
	;
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1945+v1936-int32(1)))))
	if v1949 != int32(181) {
		goto L538
	} else {
		goto L540
	}
L540:
	;
	v1955 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_115), int32(11), int32(0))
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L5
	} else {
		goto L541
	}
L541:
	;
	if v1955 == int32(0) {
		goto L538
	} else {
		goto L542
	}
L542:
	;
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1959
	v1961 = F_slice_del(m, l0)
	mBase = m.M
	if v1961 < int32(0) {
		v2200 = v1961
		goto L537
	} else {
		goto L543
	}
L543:
	;
	v1964 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v1964)
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1966
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1966
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1966-int32(3) <= v1969 {
		goto L538
	} else {
		goto L544
	}
L544:
	;
	v1973 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973+v1966-int32(1)))))
	switch v1977 - int32(129) {
	case 0, 2:
		goto L545
	default:
		goto L538
	}
L545:
	;
	v1983 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_116), int32(2), int32(0))
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L5
	} else {
		goto L546
	}
L546:
	;
	if v1983 == int32(0) {
		goto L538
	} else {
		goto L547
	}
L547:
	;
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1988 < v1987 {
		goto L538
	} else {
		goto L548
	}
L548:
	;
	v1992 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_117))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L5
	} else {
		goto L549
	}
L549:
	;
	if v1992 < int32(0) {
		v2200 = v1992
		goto L537
	} else {
		goto L550
	}
L550:
	;
	goto L538
L551:
	;
	if v2019 == int32(0) {
		v2204 = v2003
		goto L536
	} else {
		goto L555
	}
L552:
	;
	goto L551
L553:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2015 = F_memcmp(m, v2012+v2000-v2004, int32(_a_F_greek_UTF_8_stem_118), v2004)
	mBase = m.M
	if v2015 != 0 {
		v2019 = v2003
		goto L552
	} else {
		goto L554
	}
L554:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2000 - v2004
	v2019 = int32(1)
	goto L552
L555:
	;
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2022
	v2024 = F_slice_del(m, l0)
	mBase = m.M
	if v2024 < int32(0) {
		v2200 = v2024
		goto L537
	} else {
		goto L556
	}
L556:
	;
	v2027 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v2027)
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2029
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2029
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L560
L557:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2176
	v2179 = int32(0)
	v2183 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_119), int32(95), v2179)
	mBase = m.M
	v2184 = m.ExcPending
	if v2184 != 0 {
		goto L5
	} else {
		goto L586
	}
L558:
	;
	if v2161 == int32(0) {
		goto L581
	} else {
		goto L582
	}
L559:
	;
	v2161 = v2154
	goto L558
L560:
	;
	if v2029 <= v2046 {
		v2154 = int32(-1)
		goto L559
	} else {
		goto L562
	}
L561:
	;
	v2154 = int32(0)
	goto L559
L562:
	;
	v2063 = int32(1)
	v2064 = v2029 - v2063
	v2066 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2047+v2064))))
	v2068 = v2066 & int32(255)
	if base.B2i32(v2064 == v2046)|base.B2i32(int32(0) <= v2066) != 0 {
		v2126 = v2068
		v2130 = v2063
		goto L563
	} else {
		goto L564
	}
L563:
	;
	if int32(969) < v2126 {
		goto L571
	} else {
		goto L572
	}
L564:
	;
	v2075 = v2068 & int32(63)
	v2077 = v2029 - int32(2)
	v2079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2047+v2077))))
	v2081 = v2079 << (uint(int32(6)) % 32)
	if base.B2i32(v2077 != v2046)&base.B2i32(base.Ui32(v2079) < base.Ui32(int32(192))) == int32(0) {
		goto L565
	} else {
		goto L566
	}
L565:
	;
	v2126 = v2081&int32(1984) | v2075
	v2130 = int32(2)
	goto L563
L566:
	;
	goto L567
L567:
	;
	v2094 = v2081&int32(4032) | v2075
	v2096 = v2029 - int32(3)
	v2098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2047+v2096))))
	if base.B2i32(v2096 != v2046)&base.B2i32(base.Ui32(v2098) < base.Ui32(int32(224))) == int32(0) {
		goto L568
	} else {
		goto L569
	}
L568:
	;
	v2126 = v2098<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_102) | v2094
	v2130 = int32(3)
	goto L563
L569:
	;
	goto L570
L570:
	;
	v2116 = int32(4)
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2029+v2047-v2116))))
	v2126 = v2098<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_103) | v2118&int32(7)<<(uint(int32(18))%32) | v2094
	v2130 = v2116
	goto L563
L571:
	;
	v2161 = v2130
	goto L558
L572:
	;
	goto L573
L573:
	;
	v2132 = v2126 - int32(945)
	if v2132 < int32(0) {
		goto L574
	} else {
		goto L575
	}
L574:
	;
	v2161 = v2130
	goto L558
L575:
	;
	goto L576
L576:
	;
	v2138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2132)>>(uint(int32(3))%32)))+uint32(_c_F_greek_UTF_8_stem[1]))))
	if int32(base.Ui32(v2138)>>(uint(v2132&int32(7))%32))&int32(1) == int32(0) {
		goto L577
	} else {
		goto L578
	}
L577:
	;
	v2161 = v2130
	goto L558
L578:
	;
	goto L579
L579:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2029 - v2130
	goto L580
L580:
	;
	goto L561
L581:
	;
	v2166 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_120))
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L5
	} else {
		goto L584
	}
L582:
	;
	goto L583
L583:
	;
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2173 = v2171 + (v2029 - v2032)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2173
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2173
	v2176 = v2173
	goto L557
L584:
	;
	if v2166 < int32(0) {
		v2200 = v2166
		goto L537
	} else {
		goto L585
	}
L585:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2176 = v2170
	goto L557
L586:
	;
	if v2183 == int32(0) {
		v2200 = v2179
		goto L537
	} else {
		goto L587
	}
L587:
	;
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2188 < v2187 {
		v2200 = v2179
		goto L537
	} else {
		goto L588
	}
L588:
	;
	v2193 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_121))
	mBase = m.M
	v2194 = m.ExcPending
	if v2194 != 0 {
		goto L5
	} else {
		goto L589
	}
L589:
	;
	if int32(0) <= v2193 {
		goto L590
	} else {
		goto L591
	}
L590:
	;
	v2197 = int32(1)
	goto L592
L591:
	;
	v2197 = v2193
	goto L592
L592:
	;
	v2200 = v2197
	goto L537
L593:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2208 = v2207 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2208
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2208
	v2213 = int32(10)
	v2215 = int32(0)
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2208-v2218 < v2213 {
		v2228 = v2215
		goto L597
	} else {
		goto L598
	}
L594:
	;
	if v2457 < int32(0) {
		v3349 = v2457
		goto L1
	} else {
		goto L652
	}
L595:
	;
	v2457 = v2453
	goto L594
L596:
	;
	if v2228 != 0 {
		goto L600
	} else {
		goto L601
	}
L597:
	;
	goto L596
L598:
	;
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2224 = F_memcmp(m, v2221+v2208-v2213, int32(_a_F_greek_UTF_8_stem_122), v2213)
	mBase = m.M
	if v2224 != 0 {
		v2228 = v2215
		goto L597
	} else {
		goto L599
	}
L599:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2208 - v2213
	v2228 = int32(1)
	goto L597
L600:
	;
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2229
	v2231 = F_slice_del(m, l0)
	mBase = m.M
	if v2231 < int32(0) {
		v2453 = v2231
		goto L595
	} else {
		goto L603
	}
L601:
	;
	goto L602
L602:
	;
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2239 = v2237 + (v2208 - v2207)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2239
	v2242 = int32(0)
	v2243 = int32(6)
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2239-v2248 < v2243 {
		v2258 = v2242
		goto L605
	} else {
		goto L606
	}
L603:
	;
	v2234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v2234)
	goto L602
L604:
	;
	if v2258 == int32(0) {
		v2457 = v2242
		goto L594
	} else {
		goto L608
	}
L605:
	;
	goto L604
L606:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2254 = F_memcmp(m, v2251+v2239-v2243, int32(_a_F_greek_UTF_8_stem_123), v2243)
	mBase = m.M
	if v2254 != 0 {
		v2258 = v2242
		goto L605
	} else {
		goto L607
	}
L607:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2239 - v2243
	v2258 = int32(1)
	goto L605
L608:
	;
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2261
	v2263 = F_slice_del(m, l0)
	mBase = m.M
	if v2263 < int32(0) {
		v2453 = v2263
		goto L595
	} else {
		goto L609
	}
L609:
	;
	v2266 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v2266)
	v2268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2268
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2268
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L613
L610:
	;
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2432
	v2434 = int32(0)
	v2438 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_124), int32(25), v2434)
	mBase = m.M
	v2439 = m.ExcPending
	if v2439 != 0 {
		goto L5
	} else {
		goto L645
	}
L611:
	;
	if v2400 == int32(0) {
		goto L634
	} else {
		goto L635
	}
L612:
	;
	v2400 = v2393
	goto L611
L613:
	;
	if v2268 <= v2285 {
		v2393 = int32(-1)
		goto L612
	} else {
		goto L615
	}
L614:
	;
	v2393 = int32(0)
	goto L612
L615:
	;
	v2302 = int32(1)
	v2303 = v2268 - v2302
	v2305 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2286+v2303))))
	v2307 = v2305 & int32(255)
	if base.B2i32(v2303 == v2285)|base.B2i32(int32(0) <= v2305) != 0 {
		v2365 = v2307
		v2369 = v2302
		goto L616
	} else {
		goto L617
	}
L616:
	;
	if int32(969) < v2365 {
		goto L624
	} else {
		goto L625
	}
L617:
	;
	v2314 = v2307 & int32(63)
	v2316 = v2268 - int32(2)
	v2318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2286+v2316))))
	v2320 = v2318 << (uint(int32(6)) % 32)
	if base.B2i32(v2316 != v2285)&base.B2i32(base.Ui32(v2318) < base.Ui32(int32(192))) == int32(0) {
		goto L618
	} else {
		goto L619
	}
L618:
	;
	v2365 = v2320&int32(1984) | v2314
	v2369 = int32(2)
	goto L616
L619:
	;
	goto L620
L620:
	;
	v2333 = v2320&int32(4032) | v2314
	v2335 = v2268 - int32(3)
	v2337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2286+v2335))))
	if base.B2i32(v2335 != v2285)&base.B2i32(base.Ui32(v2337) < base.Ui32(int32(224))) == int32(0) {
		goto L621
	} else {
		goto L622
	}
L621:
	;
	v2365 = v2337<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_102) | v2333
	v2369 = int32(3)
	goto L616
L622:
	;
	goto L623
L623:
	;
	v2355 = int32(4)
	v2357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2268+v2286-v2355))))
	v2365 = v2337<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_103) | v2357&int32(7)<<(uint(int32(18))%32) | v2333
	v2369 = v2355
	goto L616
L624:
	;
	v2400 = v2369
	goto L611
L625:
	;
	goto L626
L626:
	;
	v2371 = v2365 - int32(945)
	if v2371 < int32(0) {
		goto L627
	} else {
		goto L628
	}
L627:
	;
	v2400 = v2369
	goto L611
L628:
	;
	goto L629
L629:
	;
	v2377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2371)>>(uint(int32(3))%32)))+uint32(_c_F_greek_UTF_8_stem[1]))))
	if int32(base.Ui32(v2377)>>(uint(v2371&int32(7))%32))&int32(1) == int32(0) {
		goto L630
	} else {
		goto L631
	}
L630:
	;
	v2400 = v2369
	goto L611
L631:
	;
	goto L632
L632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2268 - v2369
	goto L633
L633:
	;
	goto L614
L634:
	;
	v2405 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_125))
	mBase = m.M
	v2406 = m.ExcPending
	if v2406 != 0 {
		goto L5
	} else {
		goto L637
	}
L635:
	;
	goto L636
L636:
	;
	v2409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2410 = v2271 - v2268
	v2411 = v2409 - v2410
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2411
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2411
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2411
	v2418 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_126), int32(31), int32(0))
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L5
	} else {
		goto L639
	}
L637:
	;
	if int32(0) <= v2405 {
		goto L610
	} else {
		goto L638
	}
L638:
	;
	v2453 = v2405
	goto L595
L639:
	;
	if v2418 != 0 {
		goto L640
	} else {
		goto L641
	}
L640:
	;
	v2422 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_127))
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L5
	} else {
		goto L643
	}
L641:
	;
	goto L642
L642:
	;
	v2426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2427 = v2426 - v2410
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2427
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2427
	goto L610
L643:
	;
	if int32(0) <= v2422 {
		goto L610
	} else {
		goto L644
	}
L644:
	;
	v2453 = v2422
	goto L595
L645:
	;
	if v2438 == int32(0) {
		v2453 = v2434
		goto L595
	} else {
		goto L646
	}
L646:
	;
	v2442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2443 < v2442 {
		v2453 = v2434
		goto L595
	} else {
		goto L647
	}
L647:
	;
	v2448 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_128))
	mBase = m.M
	v2449 = m.ExcPending
	if v2449 != 0 {
		goto L5
	} else {
		goto L648
	}
L648:
	;
	if int32(0) <= v2448 {
		goto L649
	} else {
		goto L650
	}
L649:
	;
	v2452 = int32(1)
	goto L651
L650:
	;
	v2452 = v2448
	goto L651
L651:
	;
	v2453 = v2452
	goto L595
L652:
	;
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2461 = v2460 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2461
	v2463 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2461
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2461-int32(9) <= v2466 {
		v2559 = v2463
		goto L654
	} else {
		goto L655
	}
L653:
	;
	if v2563 < int32(0) {
		v3349 = v2563
		goto L1
	} else {
		goto L680
	}
L654:
	;
	v2563 = v2559
	goto L653
L655:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2470+v2461-int32(1)))))
	if v2474 != int32(131) {
		v2559 = v2463
		goto L654
	} else {
		goto L656
	}
L656:
	;
	v2480 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_129), int32(2), int32(0))
	mBase = m.M
	v2481 = m.ExcPending
	if v2481 != 0 {
		goto L5
	} else {
		goto L657
	}
L657:
	;
	if v2480 == int32(0) {
		v2559 = v2463
		goto L654
	} else {
		goto L658
	}
L658:
	;
	v2484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2484
	v2486 = F_slice_del(m, l0)
	mBase = m.M
	if v2486 < int32(0) {
		v2559 = v2486
		goto L654
	} else {
		goto L659
	}
L659:
	;
	v2489 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v2489)
	v2491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2491
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2491
	v2494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2495 = int32(6)
	v2500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2491-v2500 < v2495 {
		v2510 = v2489
		goto L662
	} else {
		goto L663
	}
L660:
	;
	v2524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2526 = v2524 + (v2491 - v2494)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2526
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2526
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2526
	v2530 = int32(0)
	v2531 = int32(6)
	v2536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2526-v2536 < v2531 {
		v2546 = v2530
		goto L672
	} else {
		goto L673
	}
L661:
	;
	if v2510 == int32(0) {
		goto L660
	} else {
		goto L665
	}
L662:
	;
	goto L661
L663:
	;
	v2503 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2506 = F_memcmp(m, v2503+v2491-v2495, int32(_a_F_greek_UTF_8_stem_130), v2495)
	mBase = m.M
	if v2506 != 0 {
		v2510 = v2489
		goto L662
	} else {
		goto L664
	}
L664:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2491 - v2495
	v2510 = int32(1)
	goto L662
L665:
	;
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2514 < v2513 {
		goto L660
	} else {
		goto L666
	}
L666:
	;
	v2519 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_131))
	mBase = m.M
	v2520 = m.ExcPending
	if v2520 != 0 {
		goto L5
	} else {
		goto L667
	}
L667:
	;
	if int32(0) <= v2519 {
		goto L668
	} else {
		goto L669
	}
L668:
	;
	v2523 = int32(1)
	goto L670
L669:
	;
	v2523 = v2519
	goto L670
L670:
	;
	v2563 = v2523
	goto L653
L671:
	;
	if v2546 == int32(0) {
		v2563 = v2530
		goto L653
	} else {
		goto L675
	}
L672:
	;
	goto L671
L673:
	;
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2542 = F_memcmp(m, v2539+v2526-v2531, int32(_a_F_greek_UTF_8_stem_132), v2531)
	mBase = m.M
	if v2542 != 0 {
		v2546 = v2530
		goto L672
	} else {
		goto L674
	}
L674:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2526 - v2531
	v2546 = int32(1)
	goto L672
L675:
	;
	v2552 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_133))
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L5
	} else {
		goto L676
	}
L676:
	;
	if int32(0) <= v2552 {
		goto L677
	} else {
		goto L678
	}
L677:
	;
	v2556 = int32(1)
	goto L679
L678:
	;
	v2556 = v2552
	goto L679
L679:
	;
	v2559 = v2556
	goto L654
L680:
	;
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2567 = v2566 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2567
	v2569 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2567
	v2572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2567-int32(11) <= v2572 {
		v2632 = v2569
		goto L682
	} else {
		goto L683
	}
L681:
	;
	if v2635 < int32(0) {
		v3349 = v2635
		goto L1
	} else {
		goto L698
	}
L682:
	;
	v2635 = v2632
	goto L681
L683:
	;
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2576+v2567-int32(1)))))
	if v2580 != int32(181) {
		v2632 = v2569
		goto L682
	} else {
		goto L684
	}
L684:
	;
	v2586 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_134), int32(2), int32(0))
	mBase = m.M
	v2587 = m.ExcPending
	if v2587 != 0 {
		goto L5
	} else {
		goto L685
	}
L685:
	;
	if v2586 == int32(0) {
		v2632 = v2569
		goto L682
	} else {
		goto L686
	}
L686:
	;
	v2590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2590
	v2592 = F_slice_del(m, l0)
	mBase = m.M
	if v2592 < int32(0) {
		v2632 = v2592
		goto L682
	} else {
		goto L687
	}
L687:
	;
	v2595 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v2595)
	v2597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2597
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2597
	v2601 = int32(4)
	v2606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2597-v2606 < v2601 {
		v2616 = v2595
		goto L689
	} else {
		goto L690
	}
L688:
	;
	if v2616 == int32(0) {
		v2635 = v2595
		goto L681
	} else {
		goto L692
	}
L689:
	;
	goto L688
L690:
	;
	v2609 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2612 = F_memcmp(m, v2609+v2597-v2601, int32(_a_F_greek_UTF_8_stem_135), v2601)
	mBase = m.M
	if v2612 != 0 {
		v2616 = v2595
		goto L689
	} else {
		goto L691
	}
L691:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2597 - v2601
	v2616 = int32(1)
	goto L689
L692:
	;
	v2620 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2621 < v2620 {
		v2632 = int32(0)
		goto L682
	} else {
		goto L693
	}
L693:
	;
	v2626 = F_slice_from_s(m, l0, int32(10), int32(_a_F_greek_UTF_8_stem_136))
	mBase = m.M
	v2627 = m.ExcPending
	if v2627 != 0 {
		goto L5
	} else {
		goto L694
	}
L694:
	;
	if int32(0) <= v2626 {
		goto L695
	} else {
		goto L696
	}
L695:
	;
	v2630 = int32(1)
	goto L697
L696:
	;
	v2630 = v2626
	goto L697
L697:
	;
	v2632 = v2630
	goto L682
L698:
	;
	v2638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2639 = v2638 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2639
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2639
	v2644 = int32(10)
	v2646 = int32(0)
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2639-v2649 < v2644 {
		v2659 = v2646
		goto L703
	} else {
		goto L704
	}
L699:
	;
	if v2754 < int32(0) {
		v3349 = v2754
		goto L1
	} else {
		goto L728
	}
L700:
	;
	v2754 = v2752
	goto L699
L701:
	;
	v2698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2700 = v2698 + (v2639 - v2638)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2700
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2700
	v2703 = int32(0)
	v2704 = int32(8)
	v2709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2700-v2709 < v2704 {
		v2719 = v2703
		goto L716
	} else {
		goto L717
	}
L702:
	;
	if v2659 == int32(0) {
		goto L701
	} else {
		goto L706
	}
L703:
	;
	goto L702
L704:
	;
	v2652 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2655 = F_memcmp(m, v2652+v2639-v2644, int32(_a_F_greek_UTF_8_stem_137), v2644)
	mBase = m.M
	if v2655 != 0 {
		v2659 = v2646
		goto L703
	} else {
		goto L705
	}
L705:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2639 - v2644
	v2659 = int32(1)
	goto L703
L706:
	;
	v2662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2662
	v2664 = F_slice_del(m, l0)
	mBase = m.M
	if v2664 < int32(0) {
		v2752 = v2664
		goto L700
	} else {
		goto L707
	}
L707:
	;
	v2667 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v2667)
	v2669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2669
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2669
	v2673 = v2669 - int32(1)
	v2674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2673 <= v2674 {
		goto L701
	} else {
		goto L708
	}
L708:
	;
	v2676 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2676+v2673))))
	switch v2678 - int32(128) {
	case 0, 6:
		goto L709
	default:
		goto L701
	}
L709:
	;
	v2684 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_138), int32(6), int32(0))
	mBase = m.M
	v2685 = m.ExcPending
	if v2685 != 0 {
		goto L5
	} else {
		goto L710
	}
L710:
	;
	if v2684 == int32(0) {
		goto L701
	} else {
		goto L711
	}
L711:
	;
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2689 < v2688 {
		goto L701
	} else {
		goto L712
	}
L712:
	;
	v2693 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_139))
	mBase = m.M
	v2694 = m.ExcPending
	if v2694 != 0 {
		goto L5
	} else {
		goto L713
	}
L713:
	;
	if v2693 < int32(0) {
		v2752 = v2693
		goto L700
	} else {
		goto L714
	}
L714:
	;
	goto L701
L715:
	;
	if v2719 == int32(0) {
		v2754 = v2703
		goto L699
	} else {
		goto L719
	}
L716:
	;
	goto L715
L717:
	;
	v2712 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2715 = F_memcmp(m, v2712+v2700-v2704, int32(_a_F_greek_UTF_8_stem_140), v2704)
	mBase = m.M
	if v2715 != 0 {
		v2719 = v2703
		goto L716
	} else {
		goto L718
	}
L718:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2700 - v2704
	v2719 = int32(1)
	goto L716
L719:
	;
	v2722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2722
	v2724 = F_slice_del(m, l0)
	mBase = m.M
	if v2724 < int32(0) {
		v2752 = v2724
		goto L700
	} else {
		goto L720
	}
L720:
	;
	v2727 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v2727)
	v2729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2729
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2729
	v2736 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_141), int32(9), v2727)
	mBase = m.M
	v2737 = m.ExcPending
	if v2737 != 0 {
		goto L5
	} else {
		goto L721
	}
L721:
	;
	if v2736 == int32(0) {
		v2754 = v2727
		goto L699
	} else {
		goto L722
	}
L722:
	;
	v2741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2742 < v2741 {
		v2752 = int32(0)
		goto L700
	} else {
		goto L723
	}
L723:
	;
	v2747 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_142))
	mBase = m.M
	v2748 = m.ExcPending
	if v2748 != 0 {
		goto L5
	} else {
		goto L724
	}
L724:
	;
	if int32(0) <= v2747 {
		goto L725
	} else {
		goto L726
	}
L725:
	;
	v2751 = int32(1)
	goto L727
L726:
	;
	v2751 = v2747
	goto L727
L727:
	;
	v2752 = v2751
	goto L700
L728:
	;
	v2757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2758 = v2757 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2758
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2758
	v2766 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_143), int32(3), int32(0))
	mBase = m.M
	v2767 = m.ExcPending
	if v2767 != 0 {
		goto L5
	} else {
		goto L731
	}
L729:
	;
	if v2851 < int32(0) {
		v3349 = v2851
		goto L1
	} else {
		goto L756
	}
L730:
	;
	v2851 = v2847
	goto L729
L731:
	;
	if v2766 != 0 {
		goto L732
	} else {
		goto L733
	}
L732:
	;
	v2768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2768
	v2770 = F_slice_del(m, l0)
	mBase = m.M
	if v2770 < int32(0) {
		v2847 = v2770
		goto L730
	} else {
		goto L735
	}
L733:
	;
	goto L734
L734:
	;
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2778 = v2776 + (v2758 - v2757)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2778
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2778
	v2781 = int32(0)
	v2785 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_144), int32(3), v2781)
	mBase = m.M
	v2786 = m.ExcPending
	if v2786 != 0 {
		goto L5
	} else {
		goto L736
	}
L735:
	;
	v2773 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v2773)
	goto L734
L736:
	;
	if v2785 == int32(0) {
		v2847 = v2781
		goto L730
	} else {
		goto L737
	}
L737:
	;
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2789
	v2791 = F_slice_del(m, l0)
	mBase = m.M
	if v2791 < int32(0) {
		v2847 = v2791
		goto L730
	} else {
		goto L738
	}
L738:
	;
	v2794 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v2794)
	v2796 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2796
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2796
	v2799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2803 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_145), int32(6), v2794)
	mBase = m.M
	v2804 = m.ExcPending
	if v2804 != 0 {
		goto L5
	} else {
		goto L739
	}
L739:
	;
	if v2803 != 0 {
		goto L740
	} else {
		goto L741
	}
L740:
	;
	v2808 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_146))
	mBase = m.M
	v2809 = m.ExcPending
	if v2809 != 0 {
		goto L5
	} else {
		goto L743
	}
L741:
	;
	goto L742
L742:
	;
	v2813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2815 = v2813 + (v2796 - v2799)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2815
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2815
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2815
	v2819 = int32(0)
	v2821 = v2815 - int32(1)
	v2822 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2821 <= v2822 {
		v2847 = v2819
		goto L730
	} else {
		goto L747
	}
L743:
	;
	if int32(0) <= v2808 {
		goto L744
	} else {
		goto L745
	}
L744:
	;
	v2812 = int32(1)
	goto L746
L745:
	;
	v2812 = v2808
	goto L746
L746:
	;
	v2851 = v2812
	goto L729
L747:
	;
	v2824 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2824+v2821))))
	if v2826 != int32(184) {
		v2847 = v2819
		goto L730
	} else {
		goto L748
	}
L748:
	;
	v2832 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_147), int32(5), int32(0))
	mBase = m.M
	v2833 = m.ExcPending
	if v2833 != 0 {
		goto L5
	} else {
		goto L749
	}
L749:
	;
	if v2832 == int32(0) {
		v2847 = v2819
		goto L730
	} else {
		goto L750
	}
L750:
	;
	v2836 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2837 < v2836 {
		v2847 = v2819
		goto L730
	} else {
		goto L751
	}
L751:
	;
	v2842 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_148))
	mBase = m.M
	v2843 = m.ExcPending
	if v2843 != 0 {
		goto L5
	} else {
		goto L752
	}
L752:
	;
	if int32(0) <= v2842 {
		goto L753
	} else {
		goto L754
	}
L753:
	;
	v2846 = int32(1)
	goto L755
L754:
	;
	v2846 = v2842
	goto L755
L755:
	;
	v2847 = v2846
	goto L730
L756:
	;
	v2854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2855 = v2854 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2855
	v2857 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2855
	v2863 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_149), int32(3), v2857)
	mBase = m.M
	v2864 = m.ExcPending
	if v2864 != 0 {
		goto L5
	} else {
		goto L759
	}
L757:
	;
	if v2920 < int32(0) {
		v3349 = v2920
		goto L1
	} else {
		goto L777
	}
L758:
	;
	v2920 = v2916
	goto L757
L759:
	;
	if v2863 == int32(0) {
		v2916 = v2857
		goto L758
	} else {
		goto L760
	}
L760:
	;
	v2867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2867
	v2869 = F_slice_del(m, l0)
	mBase = m.M
	if v2869 < int32(0) {
		v2916 = v2869
		goto L758
	} else {
		goto L761
	}
L761:
	;
	v2872 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v2872)
	v2874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2874
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2874
	v2877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2881 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_150), int32(12), v2872)
	mBase = m.M
	v2882 = m.ExcPending
	if v2882 != 0 {
		goto L5
	} else {
		goto L762
	}
L762:
	;
	if v2881 != 0 {
		goto L763
	} else {
		goto L764
	}
L763:
	;
	v2886 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_151))
	mBase = m.M
	v2887 = m.ExcPending
	if v2887 != 0 {
		goto L5
	} else {
		goto L766
	}
L764:
	;
	goto L765
L765:
	;
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2893 = v2891 + (v2874 - v2877)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2893
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2893
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2893
	v2897 = int32(0)
	v2901 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_152), int32(25), v2897)
	mBase = m.M
	v2902 = m.ExcPending
	if v2902 != 0 {
		goto L5
	} else {
		goto L770
	}
L766:
	;
	if int32(0) <= v2886 {
		goto L767
	} else {
		goto L768
	}
L767:
	;
	v2890 = int32(1)
	goto L769
L768:
	;
	v2890 = v2886
	goto L769
L769:
	;
	v2920 = v2890
	goto L757
L770:
	;
	if v2901 == int32(0) {
		v2916 = v2897
		goto L758
	} else {
		goto L771
	}
L771:
	;
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2906 < v2905 {
		v2916 = v2897
		goto L758
	} else {
		goto L772
	}
L772:
	;
	v2911 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_153))
	mBase = m.M
	v2912 = m.ExcPending
	if v2912 != 0 {
		goto L5
	} else {
		goto L773
	}
L773:
	;
	if int32(0) <= v2911 {
		goto L774
	} else {
		goto L775
	}
L774:
	;
	v2915 = int32(1)
	goto L776
L775:
	;
	v2915 = v2911
	goto L776
L776:
	;
	v2916 = v2915
	goto L758
L777:
	;
	v2923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2924 = v2923 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2924
	v2926 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2924
	v2932 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_154), int32(3), v2926)
	mBase = m.M
	v2933 = m.ExcPending
	if v2933 != 0 {
		goto L5
	} else {
		goto L780
	}
L778:
	;
	if v2979 < int32(0) {
		v3349 = v2979
		goto L1
	} else {
		goto L792
	}
L779:
	;
	v2979 = v2976
	goto L778
L780:
	;
	if v2932 == int32(0) {
		v2976 = v2926
		goto L779
	} else {
		goto L781
	}
L781:
	;
	v2936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2936
	v2938 = F_slice_del(m, l0)
	mBase = m.M
	if v2938 < int32(0) {
		v2976 = v2938
		goto L779
	} else {
		goto L782
	}
L782:
	;
	v2941 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v2941)
	v2943 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2943
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2943
	v2948 = v2943 - int32(1)
	v2949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2948 <= v2949 {
		v2979 = v2941
		goto L778
	} else {
		goto L783
	}
L783:
	;
	v2951 = int32(0)
	v2952 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2952+v2948))))
	if v2954 != int32(189) {
		v2976 = v2951
		goto L779
	} else {
		goto L784
	}
L784:
	;
	v2960 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_155), int32(6), int32(0))
	mBase = m.M
	v2961 = m.ExcPending
	if v2961 != 0 {
		goto L5
	} else {
		goto L785
	}
L785:
	;
	if v2960 == int32(0) {
		v2976 = v2951
		goto L779
	} else {
		goto L786
	}
L786:
	;
	v2964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2965 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2965 < v2964 {
		v2976 = v2951
		goto L779
	} else {
		goto L787
	}
L787:
	;
	v2970 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_156))
	mBase = m.M
	v2971 = m.ExcPending
	if v2971 != 0 {
		goto L5
	} else {
		goto L788
	}
L788:
	;
	if int32(0) <= v2970 {
		goto L789
	} else {
		goto L790
	}
L789:
	;
	v2974 = int32(1)
	goto L791
L790:
	;
	v2974 = v2970
	goto L791
L791:
	;
	v2976 = v2974
	goto L779
L792:
	;
	v2982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2983 = v2982 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2983
	v2985 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2983
	v2992 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_157), int32(3), v2985)
	mBase = m.M
	v2993 = m.ExcPending
	if v2993 != 0 {
		goto L5
	} else {
		goto L795
	}
L793:
	;
	if v3078 < int32(0) {
		v3349 = v3078
		goto L1
	} else {
		goto L820
	}
L794:
	;
	v3078 = v3072
	goto L793
L795:
	;
	if v2992 == int32(0) {
		v3072 = v2985
		goto L794
	} else {
		goto L796
	}
L796:
	;
	v2996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2996
	v2998 = F_slice_del(m, l0)
	mBase = m.M
	if v2998 < int32(0) {
		v3072 = v2998
		goto L794
	} else {
		goto L797
	}
L797:
	;
	v3001 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v3001)
	v3003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3003
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3003
	v3006 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3007 = int32(8)
	v3012 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3003-v3012 < v3007 {
		v3022 = v3001
		goto L799
	} else {
		goto L800
	}
L798:
	;
	if v3022 != 0 {
		goto L802
	} else {
		goto L803
	}
L799:
	;
	goto L798
L800:
	;
	v3015 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3018 = F_memcmp(m, v3015+v3003-v3007, int32(_a_F_greek_UTF_8_stem_158), v3007)
	mBase = m.M
	if v3018 != 0 {
		v3022 = v3001
		goto L799
	} else {
		goto L801
	}
L801:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3003 - v3007
	v3022 = int32(1)
	goto L799
L802:
	;
	v3026 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_159))
	mBase = m.M
	v3027 = m.ExcPending
	if v3027 != 0 {
		goto L5
	} else {
		goto L805
	}
L803:
	;
	goto L804
L804:
	;
	v3031 = v3003 - v3006
	v3032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3033 = v3031 + v3032
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3033
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3033
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3033
	v3037 = int32(1)
	v3041 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_160), int32(12), int32(0))
	mBase = m.M
	v3042 = m.ExcPending
	if v3042 != 0 {
		goto L5
	} else {
		goto L812
	}
L805:
	;
	if int32(0) <= v3026 {
		goto L806
	} else {
		goto L807
	}
L806:
	;
	v3030 = int32(1)
	goto L808
L807:
	;
	v3030 = v3026
	goto L808
L808:
	;
	v3078 = v3030
	goto L793
L809:
	;
	v3072 = v3071
	goto L794
L810:
	;
	v3049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3050 = v3049 + v3031
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3050
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3050
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3050
	v3057 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_161), int32(44), int32(0))
	mBase = m.M
	v3058 = m.ExcPending
	if v3058 != 0 {
		goto L5
	} else {
		goto L815
	}
L811:
	;
	v3045 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_162))
	mBase = m.M
	v3046 = m.ExcPending
	if v3046 != 0 {
		goto L5
	} else {
		goto L813
	}
L812:
	;
	switch v3041 {
	case 0:
		goto L810
	case 1:
		goto L811
	default:
		v3072 = v3037
		goto L794
	}
L813:
	;
	if v3045 < int32(0) {
		v3071 = v3045
		goto L809
	} else {
		goto L814
	}
L814:
	;
	v3072 = v3037
	goto L794
L815:
	;
	if v3057 == int32(0) {
		v3071 = v2985
		goto L809
	} else {
		goto L816
	}
L816:
	;
	v3061 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3062 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3062 < v3061 {
		v3071 = v2985
		goto L809
	} else {
		goto L817
	}
L817:
	;
	v3066 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_163))
	mBase = m.M
	v3067 = m.ExcPending
	if v3067 != 0 {
		goto L5
	} else {
		goto L818
	}
L818:
	;
	if int32(0) <= v3066 {
		v3072 = v3037
		goto L794
	} else {
		goto L819
	}
L819:
	;
	v3071 = v3066
	goto L809
L820:
	;
	v3081 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3082 = v3081 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3082
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3082
	v3086 = int32(0)
	v3087 = int32(8)
	v3092 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3082-v3092 < v3087 {
		v3102 = v3086
		goto L823
	} else {
		goto L824
	}
L821:
	;
	if v3139 < int32(0) {
		v3349 = v3139
		goto L1
	} else {
		goto L836
	}
L822:
	;
	if v3102 == int32(0) {
		v3139 = v3086
		goto L821
	} else {
		goto L826
	}
L823:
	;
	goto L822
L824:
	;
	v3095 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3098 = F_memcmp(m, v3095+v3082-v3087, int32(_a_F_greek_UTF_8_stem_164), v3087)
	mBase = m.M
	if v3098 != 0 {
		v3102 = v3086
		goto L823
	} else {
		goto L825
	}
L825:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3082 - v3087
	v3102 = int32(1)
	goto L823
L826:
	;
	v3105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3105
	v3107 = F_slice_del(m, l0)
	mBase = m.M
	if v3107 < int32(0) {
		v3136 = v3107
		goto L827
	} else {
		goto L828
	}
L827:
	;
	v3139 = v3136
	goto L821
L828:
	;
	v3110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v3110)
	v3112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3112
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3112
	v3119 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_165), int32(10), v3110)
	mBase = m.M
	v3120 = m.ExcPending
	if v3120 != 0 {
		goto L5
	} else {
		goto L829
	}
L829:
	;
	if v3119 == int32(0) {
		v3139 = v3110
		goto L821
	} else {
		goto L830
	}
L830:
	;
	v3124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3125 < v3124 {
		v3136 = int32(0)
		goto L827
	} else {
		goto L831
	}
L831:
	;
	v3130 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_166))
	mBase = m.M
	v3131 = m.ExcPending
	if v3131 != 0 {
		goto L5
	} else {
		goto L832
	}
L832:
	;
	if int32(0) <= v3130 {
		goto L833
	} else {
		goto L834
	}
L833:
	;
	v3134 = int32(1)
	goto L835
L834:
	;
	v3134 = v3130
	goto L835
L835:
	;
	v3136 = v3134
	goto L827
L836:
	;
	v3142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3143 = v3142 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3143
	v3145 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3143
	v3148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3143-int32(7) <= v3148 {
		v3197 = v3145
		goto L838
	} else {
		goto L839
	}
L837:
	;
	if v3200 < int32(0) {
		v3349 = v3200
		goto L1
	} else {
		goto L851
	}
L838:
	;
	v3200 = v3197
	goto L837
L839:
	;
	v3152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3152+v3143-int32(1)))))
	if v3156 != int32(181) {
		v3197 = v3145
		goto L838
	} else {
		goto L840
	}
L840:
	;
	v3162 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_167), int32(3), int32(0))
	mBase = m.M
	v3163 = m.ExcPending
	if v3163 != 0 {
		goto L5
	} else {
		goto L841
	}
L841:
	;
	if v3162 == int32(0) {
		v3197 = v3145
		goto L838
	} else {
		goto L842
	}
L842:
	;
	v3166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3166
	v3168 = F_slice_del(m, l0)
	mBase = m.M
	if v3168 < int32(0) {
		v3197 = v3168
		goto L838
	} else {
		goto L843
	}
L843:
	;
	v3171 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v3171)
	v3173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3173
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3173
	v3180 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_168), int32(6), v3171)
	mBase = m.M
	v3181 = m.ExcPending
	if v3181 != 0 {
		goto L5
	} else {
		goto L844
	}
L844:
	;
	if v3180 == int32(0) {
		v3200 = v3171
		goto L837
	} else {
		goto L845
	}
L845:
	;
	v3185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3186 < v3185 {
		v3197 = int32(0)
		goto L838
	} else {
		goto L846
	}
L846:
	;
	v3191 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_169))
	mBase = m.M
	v3192 = m.ExcPending
	if v3192 != 0 {
		goto L5
	} else {
		goto L847
	}
L847:
	;
	if int32(0) <= v3191 {
		goto L848
	} else {
		goto L849
	}
L848:
	;
	v3195 = int32(1)
	goto L850
L849:
	;
	v3195 = v3191
	goto L850
L850:
	;
	v3197 = v3195
	goto L838
L851:
	;
	v3203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3204 = v3203 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3204
	v3206 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3204
	v3209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3204-int32(7) <= v3209 {
		v3258 = v3206
		goto L853
	} else {
		goto L854
	}
L852:
	;
	if v3261 < int32(0) {
		v3349 = v3261
		goto L1
	} else {
		goto L866
	}
L853:
	;
	v3261 = v3258
	goto L852
L854:
	;
	v3213 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3213+v3204-int32(1)))))
	if v3217 != int32(181) {
		v3258 = v3206
		goto L853
	} else {
		goto L855
	}
L855:
	;
	v3223 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_170), int32(3), int32(0))
	mBase = m.M
	v3224 = m.ExcPending
	if v3224 != 0 {
		goto L5
	} else {
		goto L856
	}
L856:
	;
	if v3223 == int32(0) {
		v3258 = v3206
		goto L853
	} else {
		goto L857
	}
L857:
	;
	v3227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3227
	v3229 = F_slice_del(m, l0)
	mBase = m.M
	if v3229 < int32(0) {
		v3258 = v3229
		goto L853
	} else {
		goto L858
	}
L858:
	;
	v3232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v3232)
	v3234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3234
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3234
	v3241 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_171), int32(7), v3232)
	mBase = m.M
	v3242 = m.ExcPending
	if v3242 != 0 {
		goto L5
	} else {
		goto L859
	}
L859:
	;
	if v3241 == int32(0) {
		v3261 = v3232
		goto L852
	} else {
		goto L860
	}
L860:
	;
	v3246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3247 < v3246 {
		v3258 = int32(0)
		goto L853
	} else {
		goto L861
	}
L861:
	;
	v3252 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_172))
	mBase = m.M
	v3253 = m.ExcPending
	if v3253 != 0 {
		goto L5
	} else {
		goto L862
	}
L862:
	;
	if int32(0) <= v3252 {
		goto L863
	} else {
		goto L864
	}
L863:
	;
	v3256 = int32(1)
	goto L865
L864:
	;
	v3256 = v3252
	goto L865
L865:
	;
	v3258 = v3256
	goto L853
L866:
	;
	v3264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3265 = v3264 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3265
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3265
	v3273 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_173), int32(3), int32(0))
	mBase = m.M
	v3274 = m.ExcPending
	if v3274 != 0 {
		goto L5
	} else {
		goto L868
	}
L867:
	;
	if v3307 < int32(0) {
		v3349 = v3307
		goto L1
	} else {
		goto L880
	}
L868:
	;
	if v3273 != 0 {
		goto L869
	} else {
		goto L870
	}
L869:
	;
	v3275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3275
	v3279 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_174))
	mBase = m.M
	v3280 = m.ExcPending
	if v3280 != 0 {
		goto L5
	} else {
		goto L872
	}
L870:
	;
	goto L871
L871:
	;
	v3284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3286 = v3284 + (v3265 - v3264)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3286
	v3288 = int32(0)
	v3289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v3289 == v3288 {
		v3307 = v3288
		goto L867
	} else {
		goto L874
	}
L872:
	;
	if v3279 < int32(0) {
		v3307 = v3279
		goto L867
	} else {
		goto L873
	}
L873:
	;
	goto L871
L874:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3286
	v3296 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_175), int32(84), int32(0))
	mBase = m.M
	v3297 = m.ExcPending
	if v3297 != 0 {
		goto L5
	} else {
		goto L875
	}
L875:
	;
	if v3296 == int32(0) {
		v3307 = v3288
		goto L867
	} else {
		goto L876
	}
L876:
	;
	v3300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3300
	v3303 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v3303 {
		goto L877
	} else {
		goto L878
	}
L877:
	;
	v3306 = int32(1)
	goto L879
L878:
	;
	v3306 = v3303
	goto L879
L879:
	;
	v3307 = v3306
	goto L867
L880:
	;
	v3311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3312 = v3311 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3312
	v3314 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3312
	v3317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3312-int32(7) <= v3317 {
		v3342 = v3314
		goto L881
	} else {
		goto L882
	}
L881:
	;
	if v3342 < int32(0) {
		v3349 = v3342
		goto L1
	} else {
		goto L889
	}
L882:
	;
	v3321 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3321+v3312-int32(1)))))
	switch v3325 - int32(129) {
	case 0, 3:
		goto L883
	default:
		v3342 = v3314
		goto L881
	}
L883:
	;
	v3331 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_176), int32(8), int32(0))
	mBase = m.M
	v3332 = m.ExcPending
	if v3332 != 0 {
		goto L5
	} else {
		goto L884
	}
L884:
	;
	if v3331 == int32(0) {
		v3342 = v3314
		goto L881
	} else {
		goto L885
	}
L885:
	;
	v3335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3335
	v3338 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v3338 {
		goto L886
	} else {
		goto L887
	}
L886:
	;
	v3341 = int32(1)
	goto L888
L887:
	;
	v3341 = v3338
	goto L888
L888:
	;
	v3342 = v3341
	goto L881
L889:
	;
	v3345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3345
	v3349 = int32(1)
	goto L1
}
