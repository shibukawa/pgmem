package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bt_index_check_callback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 float32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var __phi386 int32
	_ = __phi386
	var v388 int32
	_ = v388
	var __phi388 int32
	_ = __phi388
	var v390 int32
	_ = v390
	var __phi390 int32
	_ = __phi390
	var v391 int32
	_ = v391
	var __phi391 int32
	_ = __phi391
	var v393 int32
	_ = v393
	var __phi393 int32
	_ = __phi393
	var v399 int32
	_ = v399
	var __phi399 int32
	_ = __phi399
	var v400 int32
	_ = v400
	var __phi400 int32
	_ = __phi400
	var v405 int32
	_ = v405
	var __phi405 int32
	_ = __phi405
	var v407 int32
	_ = v407
	var __phi407 int32
	_ = __phi407
	var v408 int32
	_ = v408
	var __phi408 int32
	_ = __phi408
	var v410 int32
	_ = v410
	var __phi410 int32
	_ = __phi410
	var v411 int32
	_ = v411
	var __phi411 int32
	_ = __phi411
	var v412 int32
	_ = v412
	var __phi412 int32
	_ = __phi412
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int64
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v782 int64
	_ = v782
	var v785 int64
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
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
	var v835 int32
	_ = v835
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v889 int32
	_ = v889
	var v897 int32
	_ = v897
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v924 int32
	_ = v924
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v955 int32
	_ = v955
	var v967 int32
	_ = v967
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1055 int32
	_ = v1055
	var v1063 int32
	_ = v1063
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1090 int32
	_ = v1090
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int64
	_ = v1175
	var v1178 int64
	_ = v1178
	var v1187 int32
	_ = v1187
	var v1193 int32
	_ = v1193
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1301 int32
	_ = v1301
	var v1306 int32
	_ = v1306
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1328 int32
	_ = v1328
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1448 int32
	_ = v1448
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1464 int32
	_ = v1464
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1513 int32
	_ = v1513
	var v1514 int64
	_ = v1514
	var v1519 int64
	_ = v1519
	var v1525 int32
	_ = v1525
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1536 int32
	_ = v1536
	var v1542 int32
	_ = v1542
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1556 int32
	_ = v1556
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1571 int32
	_ = v1571
	var v1582 int32
	_ = v1582
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1653 int32
	_ = v1653
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1690 int32
	_ = v1690
	var v1695 int32
	_ = v1695
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1708 int32
	_ = v1708
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1753 int32
	_ = v1753
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1776 int32
	_ = v1776
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1808 int32
	_ = v1808
	var v1813 int32
	_ = v1813
	var v1818 int32
	_ = v1818
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1837 int32
	_ = v1837
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1881 int32
	_ = v1881
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1903 int32
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1912 int32
	_ = v1912
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1926 int32
	_ = v1926
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1953 int32
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1964 int32
	_ = v1964
	var v1990 int32
	_ = v1990
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2002 int32
	_ = v2002
	var v2007 int32
	_ = v2007
	var v2014 int32
	_ = v2014
	var v2019 int32
	_ = v2019
	var v2024 int32
	_ = v2024
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2040 int64
	_ = v2040
	var v2043 int64
	_ = v2043
	var v2055 int32
	_ = v2055
	var v2059 int32
	_ = v2059
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2089 int32
	_ = v2089
	var v2090 int64
	_ = v2090
	var v2095 int64
	_ = v2095
	var v2101 int32
	_ = v2101
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2112 int32
	_ = v2112
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2148 int32
	_ = v2148
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2164 int64
	_ = v2164
	var v2169 int64
	_ = v2169
	var v2175 int32
	_ = v2175
	var v2181 int32
	_ = v2181
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2200 int32
	_ = v2200
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2223 int32
	_ = v2223
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2249 int32
	_ = v2249
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2277 int32
	_ = v2277
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2292 int64
	_ = v2292
	var v2296 int64
	_ = v2296
	var v2302 int32
	_ = v2302
	var v2312 int32
	_ = v2312
	var v2317 int32
	_ = v2317
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2336 int32
	_ = v2336
	var v2341 int32
	_ = v2341
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2348 int32
	_ = v2348
	var v2350 int32
	_ = v2350
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2372 int64
	_ = v2372
	var v2377 int64
	_ = v2377
	var v2383 int32
	_ = v2383
	var v2388 int32
	_ = v2388
	var v2390 int32
	_ = v2390
	var v2395 int32
	_ = v2395
	var v2398 int32
	_ = v2398
	var v2421 int32
	_ = v2421
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2429 int32
	_ = v2429
	var v2434 int32
	_ = v2434
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2444 int32
	_ = v2444
	var v2449 int32
	_ = v2449
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2459 int32
	_ = v2459
	var v2463 int32
	_ = v2463
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2471 int32
	_ = v2471
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2497 int32
	_ = v2497
	var v2499 int32
	_ = v2499
	var v2508 int32
	_ = v2508
	var v2532 int32
	_ = v2532
	var v2535 int32
	_ = v2535
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2551 int32
	_ = v2551
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2558 int32
	_ = v2558
	var v2560 int32
	_ = v2560
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2576 int32
	_ = v2576
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2612 int32
	_ = v2612
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2625 int32
	_ = v2625
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2631 int32
	_ = v2631
	var v2638 int32
	_ = v2638
	var v2641 int32
	_ = v2641
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2660 int32
	_ = v2660
	var v2663 int32
	_ = v2663
	var v2666 int32
	_ = v2666
	var v2668 int32
	_ = v2668
	var v2669 int32
	_ = v2669
	var v2671 int32
	_ = v2671
	var v2673 int32
	_ = v2673
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2693 int32
	_ = v2693
	var v2730 int32
	_ = v2730
	var v2739 int32
	_ = v2739
	var v2764 int32
	_ = v2764
	var v2771 int32
	_ = v2771
	var v2774 int32
	_ = v2774
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2784 int32
	_ = v2784
	var v2785 int32
	_ = v2785
	var v2786 int64
	_ = v2786
	var v2790 int64
	_ = v2790
	var v2796 int32
	_ = v2796
	var v2801 int32
	_ = v2801
	var v2805 int32
	_ = v2805
	var v2808 int32
	_ = v2808
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2818 int32
	_ = v2818
	var v2819 int32
	_ = v2819
	var v2820 int64
	_ = v2820
	var v2825 int64
	_ = v2825
	var v2831 int32
	_ = v2831
	var v2836 int32
	_ = v2836
	var v2838 int32
	_ = v2838
	var v2842 int32
	_ = v2842
	var v2845 int32
	_ = v2845
	var v2846 int32
	_ = v2846
	var v2847 int32
	_ = v2847
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2864 int32
	_ = v2864
	var v2869 int32
	_ = v2869
	var v2878 int32
	_ = v2878
	var v2902 int32
	_ = v2902
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2909 int32
	_ = v2909
	var v2911 int32
	_ = v2911
	var v2913 int32
	_ = v2913
	var v2921 int32
	_ = v2921
	var v2935 int32
	_ = v2935
	var v2947 int32
	_ = v2947
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2952 int32
	_ = v2952
	var v2955 int32
	_ = v2955
	var v2958 int32
	_ = v2958
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
	var v2964 int32
	_ = v2964
	var v2965 int32
	_ = v2965
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2970 int32
	_ = v2970
	var v2971 int32
	_ = v2971
	var v2974 int32
	_ = v2974
	var v2975 int32
	_ = v2975
	var v2977 int32
	_ = v2977
	var v2979 int32
	_ = v2979
	var v2986 int32
	_ = v2986
	var v2988 int32
	_ = v2988
	var v2989 int32
	_ = v2989
	var v2991 int32
	_ = v2991
	var v3002 int32
	_ = v3002
	var v3005 int32
	_ = v3005
	var v3006 int32
	_ = v3006
	var v3015 int32
	_ = v3015
	var v3020 int32
	_ = v3020
	var v3024 int32
	_ = v3024
	var v3027 int32
	_ = v3027
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3038 int32
	_ = v3038
	var v3043 int32
	_ = v3043
	var v3047 int32
	_ = v3047
	var v3050 int32
	_ = v3050
	var v3051 int32
	_ = v3051
	var v3059 int32
	_ = v3059
	var v3063 int32
	_ = v3063
	var v3068 int32
	_ = v3068
	var v3069 int32
	_ = v3069
	var v3072 int32
	_ = v3072
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3077 int32
	_ = v3077
	var v3079 int32
	_ = v3079
	var v3082 int32
	_ = v3082
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3093 int32
	_ = v3093
	var v3095 int32
	_ = v3095
	var v3103 int32
	_ = v3103
	var v3104 int32
	_ = v3104
	var v3105 int32
	_ = v3105
	var v3106 int32
	_ = v3106
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3109 int32
	_ = v3109
	var v3119 int32
	_ = v3119
	var v3124 int32
	_ = v3124
	var v3126 int32
	_ = v3126
	var v3127 int32
	_ = v3127
	var v3129 int32
	_ = v3129
	var v3134 int32
	_ = v3134
	var v3135 int32
	_ = v3135
	var v3136 float64
	_ = v3136
	var v3137 int32
	_ = v3137
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3142 int64
	_ = v3142
	var v3143 int32
	_ = v3143
	var v3144 int64
	_ = v3144
	var v3145 int32
	_ = v3145
	var v3147 int32
	_ = v3147
	var v3148 int64
	_ = v3148
	var v3151 int32
	_ = v3151
	var v3157 int32
	_ = v3157
	var v3165 int32
	_ = v3165
	var v3166 int32
	_ = v3166
	var v3191 int64
	_ = v3191
	var v3195 int32
	_ = v3195
	var v3196 int32
	_ = v3196
	var v3197 int32
	_ = v3197
	var v3198 int64
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3200 int64
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3202 int64
	_ = v3202
	var v3203 int32
	_ = v3203
	var v3204 int64
	_ = v3204
	var v3208 int64
	_ = v3208
	var v3210 int32
	_ = v3210
	var v3216 int32
	_ = v3216
	var v3242 int64
	_ = v3242
	var v3248 int32
	_ = v3248
	var v3249 int32
	_ = v3249
	var v3275 int64
	_ = v3275
	var v3279 int32
	_ = v3279
	var v3280 int64
	_ = v3280
	var v3281 int64
	_ = v3281
	var v3282 int32
	_ = v3282
	var v3285 int32
	_ = v3285
	var v3287 int64
	_ = v3287
	var v3288 int32
	_ = v3288
	var v3295 int32
	_ = v3295
	var v3300 int32
	_ = v3300
	var v3302 int64
	_ = v3302
	var v3305 int32
	_ = v3305
	var v3306 int32
	_ = v3306
	var v3307 int32
	_ = v3307
	var v3308 int32
	_ = v3308
	var v3309 int64
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3311 int64
	_ = v3311
	var v3312 int32
	_ = v3312
	var v3313 int64
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3315 int64
	_ = v3315
	var v3319 int64
	_ = v3319
	var v3321 int32
	_ = v3321
	var v3325 int32
	_ = v3325
	var v3327 int64
	_ = v3327
	var v3332 int32
	_ = v3332
	var v3333 int32
	_ = v3333
	var v3334 int64
	_ = v3334
	var v3338 int32
	_ = v3338
	var v3339 int64
	_ = v3339
	var v3340 int64
	_ = v3340
	var v3341 int32
	_ = v3341
	var v3344 int32
	_ = v3344
	var v3348 int64
	_ = v3348
	var v3358 int64
	_ = v3358
	var v3359 int64
	_ = v3359
	var v3388 int64
	_ = v3388
	var v3389 int64
	_ = v3389
	var v3404 int32
	_ = v3404
	var v3409 int32
	_ = v3409
	var v3442 int32
	_ = v3442
	var v3444 int32
	_ = v3444
	var v3477 int32
	_ = v3477
	var v3479 int32
	_ = v3479
	var v3480 int32
	_ = v3480
	var v3482 int32
	_ = v3482
	var v3489 int32
	_ = v3489
	var v3493 int32
	_ = v3493
	var v3498 int32
	_ = v3498
	var v3507 int32
	_ = v3507
	var v3510 int32
	_ = v3510
	var v3511 int32
	_ = v3511
	var v3512 int32
	_ = v3512
	var v3520 int32
	_ = v3520
	var v3521 int32
	_ = v3521
	var v3522 int64
	_ = v3522
	var v3530 int64
	_ = v3530
	var v3536 int32
	_ = v3536
	var v3541 int32
	_ = v3541
	var v3542 int32
	_ = v3542
	var v3547 int32
	_ = v3547
	var v3552 int32
	_ = v3552
	var v3553 int32
	_ = v3553
	var v3558 int32
	_ = v3558
	var v3559 int32
	_ = v3559
	var v3565 int32
	_ = v3565
	var v3566 int32
	_ = v3566
	var v3567 int32
	_ = v3567
	var v3568 int32
	_ = v3568
	var v3569 int32
	_ = v3569
	var v3578 int32
	_ = v3578
	var v3579 int32
	_ = v3579
	var v3583 int32
	_ = v3583
	var v3586 int32
	_ = v3586
	var v3587 int32
	_ = v3587
	var v3588 int32
	_ = v3588
	var v3596 int32
	_ = v3596
	var v3597 int32
	_ = v3597
	var v3598 int64
	_ = v3598
	var v3603 int64
	_ = v3603
	var v3609 int32
	_ = v3609
	var v3615 int32
	_ = v3615
	var v3620 int32
	_ = v3620
	var v3621 int32
	_ = v3621
	var v3622 int32
	_ = v3622
	var v3626 int32
	_ = v3626
	var v3629 int32
	_ = v3629
	var v3630 int32
	_ = v3630
	var v3631 int32
	_ = v3631
	var v3639 int32
	_ = v3639
	var v3640 int32
	_ = v3640
	var v3643 int32
	_ = v3643
	var v3644 int32
	_ = v3644
	var v3649 int32
	_ = v3649
	var v3655 int32
	_ = v3655
	var v3656 int32
	_ = v3656
	var v3657 int32
	_ = v3657
	var v3659 int32
	_ = v3659
	var v3660 int32
	_ = v3660
	var v3661 int64
	_ = v3661
	var v3666 int64
	_ = v3666
	var v3672 int32
	_ = v3672
	var v3678 int32
	_ = v3678
	var v3683 int32
	_ = v3683
	var v3687 int32
	_ = v3687
	var v3688 int32
	_ = v3688
	var v3691 int32
	_ = v3691
	var v3695 int32
	_ = v3695
	var v3728 int32
	_ = v3728
	var v3732 int32
	_ = v3732
	var v3737 int32
	_ = v3737
	var v3740 int32
	_ = v3740
	var v3741 int32
	_ = v3741
	var v3747 int32
	_ = v3747
	var v3751 int32
	_ = v3751
	var v3787 int32
	_ = v3787
	var v3790 int32
	_ = v3790
	var v3791 int32
	_ = v3791
	var v3799 int32
	_ = v3799
	var v3836 int32
	_ = v3836
	var v3840 int32
	_ = v3840
	var v3843 int32
	_ = v3843
	var v3844 int32
	_ = v3844
	var v3852 int32
	_ = v3852
	var v3857 int32
	_ = v3857
	var v3861 int32
	_ = v3861
	var v3864 int32
	_ = v3864
	var v3865 int32
	_ = v3865
	var v3873 int32
	_ = v3873
	var v3878 int32
	_ = v3878
	v4 = l3
	v33 = m.G0
	v35 = v33 - int32(80)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v37 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3861 = m.ExcPending
	if v3861 != 0 {
		goto L5
	} else {
		goto L839
	}
L2:
	;
	v63 = v37
	goto L4
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+72)) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+64)) = v41
	v45 = F_smgropen(m, v35-int32(-64), v38)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v65 = F_smgrexists(m, v63, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L11
	}
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v45
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+72))
	if v49 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v63 = v61
	goto L4
L8:
	;
	v57 = v49
	goto L10
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+76))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v45)+72))
	v57 = v55
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+72)) = v57 + int32(1)
	goto L7
L11:
	;
	if v65 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F__bt_metaversion(m, l0, v35+int32(79), v35+int32(78))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
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
	v3840 = m.ExcPending
	if v3840 != 0 {
		goto L5
	} else {
		goto L835
	}
L15:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+79)))
	v74 = int32(1)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+78)))
	if base.B2i32(v73&v74 == int32(0))&base.B2i32(v78 == v74) != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v78 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v3687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v3688 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3687)+10)))
	if v3688 <= int32(0) {
		goto L819
	} else {
		goto L820
	}
L18:
	;
	v83 = F__bt_allequalimage(m, l0, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	v88 = v73
	goto L20
L20:
	;
	v90 = v88 & int32(1)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+3)))
	v94 = m.G0
	v96 = v94 - int32(1168)
	m.G0 = v96
	v100 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L23
	}
L21:
	;
	if v83 == int32(0) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+79)))
	v88 = v87
	goto L20
L23:
	;
	if v100 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+1104)) = v102 + int32(4)
	if v4 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v121 = F_palloc0(m, int32(72))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L35
	}
L27:
	;
	v108 = int32(_a_F_bt_index_check_callback_0)
	goto L29
L28:
	;
	v108 = int32(_a_F_bt_index_check_callback_1)
	goto L29
L29:
	;
	F_errmsg_internal(m, v108, v96+int32(1104))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	if v4 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v116 = int32(395)
	goto L33
L32:
	;
	v116 = int32(392)
	goto L33
L33:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), v116, int32(_a_F_bt_index_check_callback_3))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	goto L26
L35:
	;
	v123 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+28)) = v123
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v93)
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+11)) = uint8(v92)
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+10)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+9)) = uint8(v4)
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+8)) = uint8(v90)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = l0
	if v91 == v123 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)))
	if v240 != int32(1) {
		goto L59
	} else {
		goto L60
	}
L37:
	;
	v135 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+48))
	v139 = *(*float32)(unsafe.Add(mBase, uint32(v138)+100))
	v140 = int32(_a_F_bt_index_check_callback_4)
	v141 = int32(_a_F_bt_index_check_callback_5)
	v142 = *(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[0]))
	v144 = *(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[1]))
	v145 = v142 ^ v144
	*(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[1])) = base.I64_rotl(v145, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[0])) = v145<<(uint(int64(16))%64) ^ base.I64_rotl(v142, int64(24)) ^ v145
	v159 = base.I64_extend_i32_u(v135) * int64(452)
	v160 = base.I64_trunc_sat_f32_s(v139)
	if v160 < v159 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v162 = v159
	goto L41
L40:
	;
	v162 = v160
	goto L41
L41:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[2]))
	v171 = F_bloom_create(m, v162, v164, base.I64_rotl(v142*int64(5), int64(7))*int64(9))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v121)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+60)) = v171
	v176 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v178 = F_RegisterSnapshot(m, v176)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+28)) = v178
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[3]))
	if v182 < int32(2) {
		goto L36
	} else {
		goto L45
	}
L45:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+19)))
	if v186 != int32(1) {
		goto L36
	} else {
		goto L46
	}
L46:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+20)))
	v192 = int32(768)
	if v191&v192 == v192 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L55
	}
L48:
	;
	if base.Ui32(v209) < base.Ui32(v208) {
		goto L36
	} else {
		goto L54
	}
L49:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v208 = v196
	v209 = int32(2)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v199 = int32(3)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	if base.B2i32(base.Ui32(v198) < base.Ui32(v199))|base.B2i32(base.Ui32(v201) < base.Ui32(v199)) != 0 {
		v208 = v201
		v209 = v198
		goto L48
	} else {
		goto L52
	}
L52:
	;
	if int32(0) <= v198-v201 {
		goto L47
	} else {
		goto L53
	}
L53:
	;
	goto L36
L54:
	;
	goto L47
L55:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+1088)) = v220 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_6), v96+int32(1088))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(465), int32(_a_F_bt_index_check_callback_3))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+11)))
	if v257 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L60:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v244 = F_BuildIndexInfo(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+24)) = v244
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+116)))
	if v247 != int32(1) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v121)+28))
	if v250 != 0 {
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v251 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	v253 = F_RegisterSnapshot(m, v251)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+28)) = v253
	goto L59
L66:
	;
	m.G0 = v332 + int32(80)
	return
L67:
	;
	v3621 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v3622 = *(*int32)(unsafe.Add(mBase, uint32(v830)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3626 = m.ExcPending
	if v3626 != 0 {
		goto L5
	} else {
		goto L806
	}
L68:
	;
	v3542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+7)))
	if v3542&int32(32) == int32(0) {
		v3558 = v995
		goto L793
	} else {
		goto L794
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3507 = m.ExcPending
	if v3507 != 0 {
		goto L5
	} else {
		goto L787
	}
L70:
	;
	v3069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+10)))
	if v3069 == int32(1) {
		goto L726
	} else {
		goto L727
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3047 = m.ExcPending
	if v3047 != 0 {
		goto L5
	} else {
		goto L720
	}
L72:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+8)))
	if v260 == int32(0) {
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[4]))
	v269 = F_AllocSetContextCreateInternal(m, v264, int32(_a_F_bt_index_check_callback_7), int32(0), int32(_a_F_bt_index_check_callback_8), int32(_a_F_bt_index_check_callback_9))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L5
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+16)) = v269
	v273 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+20)) = v273
	v277 = F_palloc_btree_page(m, v121, int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L5
	} else {
		goto L79
	}
L78:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v277)+32))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v277)+36))
	v319 = v315
	v323 = v96
	v324 = v121
	v326 = int32(-1)
	v332 = v35
	v333 = l0
	v338 = v316
	v341 = int32(1)
	v343 = l1
	goto L88
L79:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v277)+40))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v277)+32))
	if v279 == v280 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v284 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	if v284 == int32(0) {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+1056)) = v291 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_10), v96+int32(1056))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v277)+40))
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v277)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+1048)) = v301
	*(*int64)(unsafe.Add(mBase, uint32(v96)+1040)) = v300
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_11), v96+int32(1040))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(514), int32(_a_F_bt_index_check_callback_3))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	goto L78
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3024 = m.ExcPending
	if v3024 != 0 {
		goto L5
	} else {
		goto L716
	}
L88:
	;
	if v319 == int32(0) {
		goto L70
	} else {
		goto L90
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3002 = m.ExcPending
	if v3002 != 0 {
		goto L5
	} else {
		goto L712
	}
L90:
	;
	v353 = int32(_a_F_bt_index_check_callback_12)
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[4]))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v324)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[4])) = v356
	v360 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	if v360 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+1024)) = v338
	if v338 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	v379 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v324)+56)) = uint8(v379)
	v382 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v324)+52)) = v382
	__phi386 = v319
	__phi388 = v379
	__phi390 = v323
	__phi391 = v324
	__phi393 = v382
	__phi399 = v332
	__phi400 = v333
	__phi405 = v338
	__phi407 = v382
	__phi408 = v341
	__phi410 = v343
	__phi411 = v354
	__phi412 = v326
	v386 = __phi386
	v388 = __phi388
	v390 = __phi390
	v391 = __phi391
	v393 = __phi393
	v399 = __phi399
	v400 = __phi400
	v405 = __phi405
	v407 = __phi407
	v408 = __phi408
	v410 = __phi410
	v411 = __phi411
	v412 = __phi412
	goto L103
L95:
	;
	v366 = int32(_a_F_bt_index_check_callback_13)
	goto L97
L96:
	;
	v366 = int32(_a_F_bt_index_check_callback_14)
	goto L97
L97:
	;
	if v341 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v367 = int32(_a_F_bt_index_check_callback_15)
	goto L100
L99:
	;
	v367 = v366
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+1028)) = v367
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_16), v323+int32(1024))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(647), int32(_a_F_bt_index_check_callback_17))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	goto L94
L103:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[5]))
	if v419 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v2989 = *(*int32)(unsafe.Add(mBase, uint32(v391)+48))
	if v2989 != 0 {
		goto L707
	} else {
		goto L708
	}
L105:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L5
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+36)) = v386
	v423 = F_palloc_btree_page(m, v391, v386)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L5
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+32)) = v423
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v423)))
	*(*int64)(unsafe.Add(mBase, uint32(v391)+40)) = base.I64_rotl(v426, int64(32))
	v430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v423)+16)))
	v431 = v423 + v430
	v432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v431)+12)))
	if v432&int32(20) != 0 {
		goto L116
	} else {
		goto L117
	}
L110:
	;
	if v386 == v388 {
		goto L87
	} else {
		goto L693
	}
L111:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v431)+8))
	if v779 == v405 {
		goto L204
	} else {
		goto L205
	}
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L5
	} else {
		goto L199
	}
L113:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L5
	} else {
		goto L195
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L5
	} else {
		goto L191
	}
L115:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L5
	} else {
		goto L186
	}
L116:
	;
	if v432&int32(4) != 0 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	if v407 != int32(-1) {
		v510 = v393
		v511 = v407
		goto L129
	} else {
		goto L130
	}
L119:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+9)))
	if v437&int32(1) != 0 {
		goto L115
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v440 == int32(0) {
		goto L114
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	v445 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	if v445 == int32(0) {
		v2921 = v393
		v2935 = v407
		goto L110
	} else {
		goto L125
	}
L125:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L5
	} else {
		goto L126
	}
L126:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+1008)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v390)+1012)) = v453 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_18), v390+int32(1008))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L5
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(694), int32(_a_F_bt_index_check_callback_17))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	v2921 = v393
	v2935 = v407
	goto L110
L129:
	;
	if v388 == int32(0) {
		goto L111
	} else {
		goto L144
	}
L130:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+9)))
	if v470 == int32(1) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v473 = F_bt_leftmost_ignoring_half_dead(m, v391, v386, v431)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L5
	} else {
		goto L134
	}
L132:
	;
	v483 = v432
	goto L133
L133:
	;
	if v483&int32(1) != 0 {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	if v473 == int32(0) {
		goto L113
	} else {
		goto L135
	}
L135:
	;
	v477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v431)+12)))
	if base.B2i32(v477&int32(130) == int32(0))&v408 != 0 {
		goto L112
	} else {
		goto L136
	}
L136:
	;
	v483 = v477
	goto L133
L137:
	;
	v510 = int32(-1)
	v511 = int32(0)
	goto L129
L138:
	;
	goto L139
L139:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v492 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v493 = int32(2)
	goto L142
L141:
	;
	v493 = int32(1)
	goto L142
L142:
	;
	v494 = F_PageGetItemIdCareful_2(m, v391, v488, v489, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L5
	} else {
		goto L143
	}
L143:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	v500 = v496 + v497&int32(_a_F_bt_index_check_callback_19)
	v501 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v500))))
	v504 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v500)+2)))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v431)+8))
	v510 = v506 - int32(1)
	v511 = v501<<(uint(int32(16))%32) | v504
	goto L129
L144:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	if v514 == v388 {
		goto L111
	} else {
		goto L145
	}
L145:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+9)))
	if v516 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v520 = int32(0)
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v391)+20))
	v523 = F_ReadBufferExtended(m, v519, v520, v388, v520, v522)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L5
	} else {
		goto L149
	}
L147:
	;
	v641 = v514
	goto L148
L148:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L5
	} else {
		goto L181
	}
L149:
	;
	F_LockBufferInternal(m, v523, int32(1))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L5
	} else {
		goto L150
	}
L150:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	F__bt_checkpage(m, v528, v523)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L5
	} else {
		goto L151
	}
L151:
	;
	if v523 < int32(0) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v549 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v548)+16)))
	v550 = v549 + v548
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+12)))
	if v551&int32(4) != 0 {
		goto L156
	} else {
		goto L157
	}
L153:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[6]))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v534+(v523^int32(-1))<<(uint(int32(2))%32))))
	v548 = v540
	goto L152
L154:
	;
	goto L155
L155:
	;
	v542 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[7]))
	v548 = v542 + v523<<(uint(int32(13))%32) + int32(-8192)
	goto L152
L156:
	;
	F_UnlockReleaseBuffer(m, v523)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L5
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v550)+4))
	if v557 == v388 {
		v601 = int32(-1)
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L111
L160:
	;
	F_UnlockReleaseBuffer(m, v523)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L5
	} else {
		goto L171
	}
L161:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v560 = int32(0)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v391)+20))
	v563 = F_ReadBufferExtended(m, v559, v560, v557, v560, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L5
	} else {
		goto L162
	}
L162:
	;
	F_LockBufferInternal(m, v563, int32(1))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L5
	} else {
		goto L163
	}
L163:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	F__bt_checkpage(m, v568, v563)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L5
	} else {
		goto L164
	}
L164:
	;
	if v563 < int32(0) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	F_UnlockReleaseBuffer(m, v563)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L5
	} else {
		goto L170
	}
L166:
	;
	v574 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[6]))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v574+(v563^int32(-1))<<(uint(int32(2))%32))))
	v581 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v580)+16)))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v580+v581)))
	v598 = v583
	goto L165
L167:
	;
	goto L168
L168:
	;
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[7]))
	v588 = v585 + v563<<(uint(int32(13))%32)
	v591 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v588-int32(_a_F_bt_index_check_callback_20)))))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v588+v591)+uint32(_c_F_bt_index_check_callback[8])))
	if v563 == int32(0) {
		v601 = v595
		goto L160
	} else {
		goto L169
	}
L169:
	;
	v598 = v595
	goto L165
L170:
	;
	v601 = v598
	goto L160
L171:
	;
	if v601 == v388 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v608 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L5
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+36)) = v557
	v641 = v601
	goto L148
L175:
	;
	if v608 == int32(0) {
		goto L111
	} else {
		goto L176
	}
L176:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L5
	} else {
		goto L177
	}
L177:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v615)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+928)) = v616 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_21), v390+int32(928))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L5
	} else {
		goto L178
	}
L178:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+920)) = v625
	*(*int32)(unsafe.Add(mBase, uint32(v390)+916)) = v557
	*(*int32)(unsafe.Add(mBase, uint32(v390)+912)) = v388
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_22), v390+int32(912))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L5
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1179), int32(_a_F_bt_index_check_callback_23))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L5
	} else {
		goto L180
	}
L180:
	;
	goto L111
L181:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L5
	} else {
		goto L182
	}
L182:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v651)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+80)) = v652 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_24), v390+int32(80))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L5
	} else {
		goto L183
	}
L183:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+72)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v390)+68)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v390)+64)) = v661
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_25), v390-int32(-64))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L5
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1199), int32(_a_F_bt_index_check_callback_23))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L5
	} else {
		goto L185
	}
L185:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L186:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L5
	} else {
		goto L187
	}
L187:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v682)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+976)) = v683 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_26), v390+int32(976))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+968)) = v692
	*(*int32)(unsafe.Add(mBase, uint32(v390)+964)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v390)+960)) = v386
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_25), v390+int32(960))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L5
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(683), int32(_a_F_bt_index_check_callback_17))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L5
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L191:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+992)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v390)+996)) = v714 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_27), v390+int32(992))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L5
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(689), int32(_a_F_bt_index_check_callback_17))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L5
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
	F_errcode(m, int32(33557032))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L5
	} else {
		goto L196
	}
L196:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v736)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+944)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v390)+948)) = v737 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_28), v390+int32(944))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L5
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(712), int32(_a_F_bt_index_check_callback_17))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L5
	} else {
		goto L200
	}
L200:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v759)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+48)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v390)+52)) = v760 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_29), v390+int32(48))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L5
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(718), int32(_a_F_bt_index_check_callback_17))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L5
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	v2902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2878)+12)))
	if v2902&int32(1) != 0 {
		v2921 = v510
		v2935 = v511
		goto L110
	} else {
		goto L689
	}
L204:
	;
	v782 = *(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v390)+1128)) = v782
	v785 = *(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[10]))
	*(*int64)(unsafe.Add(mBase, uint32(v390)+1120)) = v785
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v788 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v787)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v788) {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	goto L206
L206:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2842 = m.ExcPending
	if v2842 != 0 {
		goto L5
	} else {
		goto L684
	}
L207:
	;
	v796 = int32(base.Ui32(v788+int32(_a_F_bt_index_check_callback_30)) >> (uint(int32(2)) % 32))
	goto L209
L208:
	;
	v796 = int32(0)
	goto L209
L209:
	;
	v797 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v787)+16)))
	v798 = v787 + v797
	v801 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L5
	} else {
		goto L210
	}
L210:
	;
	if v801 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v803 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v798)+12)))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+872)) = v804
	*(*int32)(unsafe.Add(mBase, uint32(v390)+864)) = v796 & int32(_a_F_bt_index_check_callback_31)
	if v803&int32(1) != 0 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	goto L213
L213:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v798)+4))
	if v826 != 0 {
		goto L219
	} else {
		goto L220
	}
L214:
	;
	v813 = int32(_a_F_bt_index_check_callback_32)
	goto L216
L215:
	;
	v813 = int32(_a_F_bt_index_check_callback_33)
	goto L216
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+868)) = v813
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_34), v390+int32(864))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L5
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1254), int32(_a_F_bt_index_check_callback_35))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L5
	} else {
		goto L218
	}
L218:
	;
	goto L213
L219:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v830 = F_PageGetItemIdCareful_2(m, v391, v827, v828, int32(1))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L5
	} else {
		goto L222
	}
L220:
	;
	v943 = int32(1)
	goto L221
L221:
	;
	v945 = v796 & int32(_a_F_bt_index_check_callback_31)
	if base.Ui32(v945) < base.Ui32(v943) {
		v2878 = v798
		goto L203
	} else {
		goto L261
	}
L222:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+8)))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v835 = int32(1)
	v843 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v834)+16)))
	v844 = v834 + v843
	v845 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v844)+12)))
	if v845&int32(20) != 0 {
		v924 = v835
		goto L224
	} else {
		goto L225
	}
L223:
	;
	if v934 == int32(0) {
		goto L67
	} else {
		goto L257
	}
L224:
	;
	v934 = v924
	goto L223
L225:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v832)+192))
	v849 = int32(*(*int16)(unsafe.Add(mBase, uint32(v848)+10)))
	v850 = int32(*(*int16)(unsafe.Add(mBase, uint32(v848)+8)))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v834+int32(4))+20))
	v857 = v834 + v854&int32(_a_F_bt_index_check_callback_19)
	v858 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v857)+6)))
	v860 = v858 & int32(_a_F_bt_index_check_callback_8)
	if v860 == int32(0) {
		v878 = v850
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v844)+4))
	if v881 != 0 {
		goto L232
	} else {
		goto L233
	}
L227:
	;
	v863 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v857)+4)))
	if v863&int32(_a_F_bt_index_check_callback_8) != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v866 = int32(0)
	if base.B2i32(v833 == v866)|v863&int32(_a_F_bt_index_check_callback_36)|base.B2i32(v849 != v850) != 0 {
		v924 = v866
		goto L224
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v878 = v863 & int32(4095)
	goto L226
L231:
	;
	v878 = v850
	goto L226
L232:
	;
	v882 = int32(2)
	goto L234
L233:
	;
	v882 = int32(1)
	goto L234
L234:
	;
	if v845&int32(1) != 0 {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v908 = int32(0)
	if v860 == v908 {
		v924 = v908
		goto L224
	} else {
		goto L251
	}
L236:
	;
	if base.Ui32(v882) <= base.Ui32(v835) {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	goto L238
L238:
	;
	if v835 == v882 {
		goto L246
	} else {
		goto L247
	}
L239:
	;
	if v860 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L240:
	;
	goto L241
L241:
	;
	if v833 != 0 {
		goto L235
	} else {
		goto L245
	}
L242:
	;
	v934 = base.B2i32(v878 == v850)
	goto L223
L243:
	;
	goto L244
L244:
	;
	v889 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v857)+4)))
	v934 = int32(base.Ui32(v889)>>(uint(int32(13))%32)) & base.B2i32(v878 == v850)
	goto L223
L245:
	;
	v934 = base.B2i32(v878 == v849)
	goto L223
L246:
	;
	v897 = base.B2i32(v878 == int32(0))
	if v833|v897 != 0 {
		v924 = v897 | (v833 ^ int32(1))
		goto L224
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	if v833 != 0 {
		goto L235
	} else {
		goto L250
	}
L249:
	;
	v904 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v857)+4)))
	v934 = base.B2i32(v904 == int32(1))
	goto L223
L250:
	;
	v934 = base.B2i32(v878 == v849)
	goto L223
L251:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857)+5)))
	if v911&int32(32) != 0 {
		v924 = v908
		goto L224
	} else {
		goto L252
	}
L252:
	;
	v914 = F_BTreeTupleGetHeapTID(m, v857)
	mBase = m.M
	if v878 != v849 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v917 = v914
	goto L255
L254:
	;
	v917 = int32(0)
	goto L255
L255:
	;
	if v917 != 0 {
		v924 = v908
		goto L224
	} else {
		goto L256
	}
L256:
	;
	v924 = base.B2i32(v878 <= v849) & base.B2i32(int32(0) < v878)
	goto L224
L257:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v798)+4))
	if v939 != 0 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v940 = int32(2)
	goto L260
L259:
	;
	v940 = int32(1)
	goto L260
L260:
	;
	v943 = v940
	goto L221
L261:
	;
	v955 = v798
	v967 = v943
	goto L264
L262:
	;
	F_pfree(m, v2477)
	mBase = m.M
	v2838 = m.ExcPending
	if v2838 != 0 {
		goto L5
	} else {
		goto L683
	}
L263:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2805 = m.ExcPending
	if v2805 != 0 {
		goto L5
	} else {
		goto L678
	}
L264:
	;
	v980 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[5]))
	if v980 != 0 {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2771 = m.ExcPending
	if v2771 != 0 {
		goto L5
	} else {
		goto L673
	}
L266:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L5
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v986 = v967 & int32(_a_F_bt_index_check_callback_31)
	v987 = F_PageGetItemIdCareful_2(m, v391, v983, v984, v986)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L5
	} else {
		goto L278
	}
L269:
	;
	goto L268
L270:
	;
	goto L265
L271:
	;
	v2764 = v967 + int32(1)
	if base.Ui32(v2764&int32(_a_F_bt_index_check_callback_31)) <= base.Ui32(v945) {
		v955 = v2739
		v967 = v2764
		goto L264
	} else {
		goto L672
	}
L272:
	;
	v2532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2508)+12)))
	if v2532&int32(1) != 0 {
		v2739 = v2508
		goto L271
	} else {
		goto L611
	}
L273:
	;
	v2421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+12)))
	if v2421 != int32(1) {
		v2508 = v955
		goto L272
	} else {
		goto L582
	}
L274:
	;
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(v2318)))
	v2327 = F__bt_mkscankey(m, v2322, v1872+v2323&int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v2328 = m.ExcPending
	if v2328 != 0 {
		goto L5
	} else {
		goto L569
	}
L275:
	;
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+484)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v390)+480)) = v2187
	v2193 = F_psprintf(m, int32(_a_F_bt_index_check_callback_37), v390+int32(480))
	mBase = m.M
	v2194 = m.ExcPending
	if v2194 != 0 {
		goto L5
	} else {
		goto L548
	}
L276:
	;
	v2107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+7)))
	if v2107&int32(32) == int32(0) {
		v2123 = v995
		goto L535
	} else {
		goto L536
	}
L277:
	;
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+612)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v390)+608)) = v2065
	v2071 = F_psprintf(m, int32(_a_F_bt_index_check_callback_37), v390+int32(608))
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L5
	} else {
		goto L528
	}
L278:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v987)))
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v995 = v992 + v989&int32(_a_F_bt_index_check_callback_19)
	v996 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+6)))
	v998 = v996 & int32(_a_F_bt_index_check_callback_38)
	if int32(base.Ui32(v989)>>(uint(int32(17))%32)) == v998 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+8)))
	v1009 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v992)+16)))
	v1010 = v992 + v1009
	v1011 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1010)+12)))
	if v1011&int32(20) != 0 {
		v1090 = int32(1)
		goto L283
	} else {
		goto L284
	}
L280:
	;
	goto L281
L281:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2024 = m.ExcPending
	if v2024 != 0 {
		goto L5
	} else {
		goto L522
	}
L282:
	;
	if v1100 == int32(0) {
		goto L316
	} else {
		goto L317
	}
L283:
	;
	v1100 = v1090
	goto L282
L284:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+192))
	v1015 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1014)+10)))
	v1016 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1014)+8)))
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v992+v986<<(uint(int32(2))%32))+20))
	v1023 = v992 + v1020&int32(_a_F_bt_index_check_callback_19)
	v1024 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1023)+6)))
	v1026 = v1024 & int32(_a_F_bt_index_check_callback_8)
	if v1026 == int32(0) {
		v1044 = v1016
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1010)+4))
	if v1047 != 0 {
		goto L291
	} else {
		goto L292
	}
L286:
	;
	v1029 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1023)+4)))
	if v1029&int32(_a_F_bt_index_check_callback_8) != 0 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1032 = int32(0)
	if base.B2i32(v1001 == v1032)|v1029&int32(_a_F_bt_index_check_callback_36)|base.B2i32(v1015 != v1016) != 0 {
		v1090 = v1032
		goto L283
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	v1044 = v1029 & int32(4095)
	goto L285
L290:
	;
	v1044 = v1016
	goto L285
L291:
	;
	v1048 = int32(2)
	goto L293
L292:
	;
	v1048 = int32(1)
	goto L293
L293:
	;
	if v1011&int32(1) != 0 {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	v1074 = int32(0)
	if v1026 == v1074 {
		v1090 = v1074
		goto L283
	} else {
		goto L310
	}
L295:
	;
	if base.Ui32(v1048) <= base.Ui32(v986) {
		goto L298
	} else {
		goto L299
	}
L296:
	;
	goto L297
L297:
	;
	if v986 == v1048 {
		goto L305
	} else {
		goto L306
	}
L298:
	;
	if v1026 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L299:
	;
	goto L300
L300:
	;
	if v1001 != 0 {
		goto L294
	} else {
		goto L304
	}
L301:
	;
	v1100 = base.B2i32(v1044 == v1016)
	goto L282
L302:
	;
	goto L303
L303:
	;
	v1055 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1023)+4)))
	v1100 = int32(base.Ui32(v1055)>>(uint(int32(13))%32)) & base.B2i32(v1044 == v1016)
	goto L282
L304:
	;
	v1100 = base.B2i32(v1044 == v1015)
	goto L282
L305:
	;
	v1063 = base.B2i32(v1044 == int32(0))
	if v1001|v1063 != 0 {
		v1090 = v1063 | (v1001 ^ int32(1))
		goto L283
	} else {
		goto L308
	}
L306:
	;
	goto L307
L307:
	;
	if v1001 != 0 {
		goto L294
	} else {
		goto L309
	}
L308:
	;
	v1070 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1023)+4)))
	v1100 = base.B2i32(v1070 == int32(1))
	goto L282
L309:
	;
	v1100 = base.B2i32(v1044 == v1015)
	goto L282
L310:
	;
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1023)+5)))
	if v1077&int32(32) != 0 {
		v1090 = v1074
		goto L283
	} else {
		goto L311
	}
L311:
	;
	v1080 = F_BTreeTupleGetHeapTID(m, v1023)
	mBase = m.M
	if v1044 != v1015 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1083 = v1080
	goto L314
L313:
	;
	v1083 = int32(0)
	goto L314
L314:
	;
	if v1083 != 0 {
		v1090 = v1074
		goto L283
	} else {
		goto L315
	}
L315:
	;
	v1090 = base.B2i32(v1044 <= v1015) & base.B2i32(int32(0) < v1044)
	goto L283
L316:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+756)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v390)+752)) = v1103
	v1109 = F_psprintf(m, int32(_a_F_bt_index_check_callback_37), v390+int32(752))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L5
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955)+12)))
	if v1199&int32(1) == int32(0) {
		goto L342
	} else {
		goto L343
	}
L319:
	;
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+7)))
	if v1111&int32(32) == int32(0) {
		v1127 = v995
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1127)+2)))
	v1129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1127))))
	v1130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1127)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+740)) = v1130
	*(*int32)(unsafe.Add(mBase, uint32(v390)+736)) = v1128 | v1129<<(uint(int32(16))%32)
	v1139 = F_psprintf(m, int32(_a_F_bt_index_check_callback_37), v390+int32(736))
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L5
	} else {
		goto L323
	}
L321:
	;
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+5)))
	if v1116&int32(32) == int32(0) {
		v1127 = v995
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+2)))
	v1122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995))))
	v1127 = v1121 + (v995 + v1122<<(uint(int32(16))%32))
	goto L320
L323:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L5
	} else {
		goto L324
	}
L324:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L5
	} else {
		goto L325
	}
L325:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+720)) = v1149 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_39), v390+int32(720))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L5
	} else {
		goto L326
	}
L326:
	;
	v1158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+7)))
	if v1158&int32(32) == int32(0) {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v1174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v955)+12)))
	v1175 = *(*int64)(unsafe.Add(mBase, uint32(v391)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+708)) = uint32(v1175)
	v1178 = int64(base.Ui64(v1175) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+704)) = uint32(v1178)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+700)) = v1139
	*(*int32)(unsafe.Add(mBase, uint32(v390)+692)) = v1173
	*(*int32)(unsafe.Add(mBase, uint32(v390)+688)) = v1109
	if v1174&int32(1) != 0 {
		goto L331
	} else {
		goto L332
	}
L328:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+192))
	v1171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1170)+8)))
	v1173 = v1171
	goto L327
L329:
	;
	v1163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+4)))
	if v1163&int32(_a_F_bt_index_check_callback_8) != 0 {
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v1173 = v1163 & int32(4095)
	goto L327
L331:
	;
	v1187 = int32(_a_F_bt_index_check_callback_40)
	goto L333
L332:
	;
	v1187 = int32(_a_F_bt_index_check_callback_41)
	goto L333
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+696)) = v1187
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_42), v390+int32(688))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L5
	} else {
		goto L334
	}
L334:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1354), int32(_a_F_bt_index_check_callback_35))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L5
	} else {
		goto L335
	}
L335:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L336:
	;
	if base.Ui32(v1551) < base.Ui32(v998) {
		goto L276
	} else {
		goto L403
	}
L337:
	;
	v1532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955)+12)))
	if v1532&int32(1) != 0 {
		v1551 = int32(2704)
		goto L336
	} else {
		goto L397
	}
L338:
	;
	v1459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+7)))
	if v1459&int32(32) == int32(0) {
		v1475 = v995
		goto L387
	} else {
		goto L388
	}
L339:
	;
	F_pfree(m, v1220)
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L5
	} else {
		goto L385
	}
L340:
	;
	F_UnlockReleaseBuffer(m, v1254)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L5
	} else {
		goto L384
	}
L341:
	;
	v1301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+7)))
	if v1301&int32(32) == int32(0) {
		goto L369
	} else {
		goto L370
	}
L342:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v955)+4))
	if v1206 != 0 {
		goto L345
	} else {
		goto L346
	}
L343:
	;
	goto L344
L344:
	;
	v1216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+11)))
	if v1216 != int32(1) {
		goto L341
	} else {
		goto L351
	}
L345:
	;
	v1207 = int32(2)
	goto L347
L346:
	;
	v1207 = int32(1)
	goto L347
L347:
	;
	if v1207 != v986 {
		goto L341
	} else {
		goto L348
	}
L348:
	;
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+9)))
	if v1209 != int32(1) {
		v2739 = v955
		goto L271
	} else {
		goto L349
	}
L349:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v955)+8))
	F_bt_child_highkey_check(m, v391, v986, int32(0), v1213)
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L5
	} else {
		goto L350
	}
L350:
	;
	v2739 = v955
	goto L271
L351:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v1220 = F__bt_mkscankey(m, v1219, v995)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L5
	} else {
		goto L352
	}
L352:
	;
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v1223 = int32(0)
	v1228 = F__bt_search(m, v1222, v1223, v1220, v390+int32(1164), int32(1), v1223)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L5
	} else {
		goto L353
	}
L353:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v390)+1164))
	if v1230 == int32(0) {
		goto L339
	} else {
		goto L354
	}
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+1136)) = v995
	v1234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+6)))
	v1235 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+1160)) = v1235
	*(*int32)(unsafe.Add(mBase, uint32(v390)+1144)) = v1220
	*(*uint8)(unsafe.Add(mBase, uint32(v390)+1152)) = uint8(v1235)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+1148)) = v1230
	*(*int32)(unsafe.Add(mBase, uint32(v390)+1140)) = (v1234&int32(_a_F_bt_index_check_callback_38) + int32(7)) & int32(_a_F_bt_index_check_callback_43)
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v1251 = F__bt_binsrch_insert(m, v1248, v390+int32(1136))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L5
	} else {
		goto L355
	}
L355:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v390)+1164))
	if v1254 < int32(0) {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	v1273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1272)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1273) {
		goto L360
	} else {
		goto L361
	}
L357:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[6]))
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1258+(v1254^int32(-1))<<(uint(int32(2))%32))))
	v1272 = v1264
	goto L356
L358:
	;
	goto L359
L359:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[7]))
	v1272 = v1266 + v1254<<(uint(int32(13))%32) + int32(-8192)
	goto L356
L360:
	;
	v1281 = int32(base.Ui32(v1273+int32(_a_F_bt_index_check_callback_30)) >> (uint(int32(2)) % 32))
	goto L362
L361:
	;
	v1281 = int32(0)
	goto L362
L362:
	;
	if base.Ui32(v1281&int32(_a_F_bt_index_check_callback_31)) < base.Ui32(v1251) {
		goto L340
	} else {
		goto L363
	}
L363:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v390)+1160))
	if int32(0) < v1285 {
		goto L340
	} else {
		goto L364
	}
L364:
	;
	v1288 = F__bt_compare(m, v1253, v1220, v1272, v1251)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L5
	} else {
		goto L365
	}
L365:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v390)+1164))
	F_UnlockReleaseBuffer(m, v1291)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L5
	} else {
		goto L366
	}
L366:
	;
	F_pfree(m, v1220)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L5
	} else {
		goto L367
	}
L367:
	;
	if v1288 != 0 {
		goto L338
	} else {
		goto L368
	}
L368:
	;
	goto L341
L369:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v1439 = F__bt_mkscankey(m, v1438, v995)
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L5
	} else {
		goto L382
	}
L370:
	;
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+5)))
	if v1306&int32(32) == int32(0) {
		goto L369
	} else {
		goto L371
	}
L371:
	;
	v1311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+2)))
	v1312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995))))
	v1316 = v1311 + (v995 + v1312<<(uint(int32(16))%32))
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1316)))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+1136)) = v1317
	v1319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1316)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v390)+1140)) = uint16(v1319)
	v1322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+4)))
	if v1322&int32(4094) == int32(0) {
		goto L369
	} else {
		goto L372
	}
L372:
	;
	v1328 = int32(1)
	goto L373
L373:
	;
	v1359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+2)))
	v1360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995))))
	v1361 = int32(16)
	v1367 = v1359 + (v995 + v1360<<(uint(v1361)%32)) + v1328*int32(6)
	v1369 = v390 + int32(1136)
	v1373 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1367)+2)))
	v1374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1367))))
	v1377 = v1373 | v1374<<(uint(v1361)%32)
	v1378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1369)+2)))
	v1379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1369))))
	v1382 = v1378 | v1379<<(uint(v1361)%32)
	if base.Ui32(v1377) < base.Ui32(v1382) {
		v1393 = int32(-1)
		goto L376
	} else {
		goto L377
	}
L374:
	;
	goto L369
L375:
	;
	if v1393 <= int32(0) {
		goto L277
	} else {
		goto L380
	}
L376:
	;
	goto L375
L377:
	;
	if base.Ui32(v1382) < base.Ui32(v1377) {
		v1393 = int32(1)
		goto L376
	} else {
		goto L378
	}
L378:
	;
	v1387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1367)+4)))
	v1388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1369)+4)))
	if base.Ui32(v1387) < base.Ui32(v1388) {
		v1393 = int32(-1)
		goto L376
	} else {
		goto L379
	}
L379:
	;
	v1393 = base.B2i32(base.Ui32(v1388) < base.Ui32(v1387))
	goto L376
L380:
	;
	v1396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1367)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v390)+1140)) = uint16(v1396)
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+1136)) = v1398
	v1401 = v1328 + int32(1)
	v1402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+4)))
	if base.Ui32(v1401) < base.Ui32(v1402&int32(4095)) {
		v1328 = v1401
		goto L373
	} else {
		goto L381
	}
L381:
	;
	goto L374
L382:
	;
	v1441 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1439)+4)) = uint8(v1441)
	v1443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1439))))
	if v1443 == v1441 {
		goto L337
	} else {
		goto L383
	}
L383:
	;
	v1551 = int32(2712)
	goto L336
L384:
	;
	goto L339
L385:
	;
	goto L338
L386:
	;
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+676)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v390)+672)) = v1476
	v1482 = F_psprintf(m, int32(_a_F_bt_index_check_callback_37), v390+int32(672))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L5
	} else {
		goto L390
	}
L387:
	;
	goto L386
L388:
	;
	v1464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+5)))
	if v1464&int32(32) == int32(0) {
		v1475 = v995
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v1469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+2)))
	v1470 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995))))
	v1475 = v1469 + (v995 + v1470<<(uint(int32(16))%32))
	goto L387
L390:
	;
	v1484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1475)+2)))
	v1485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1475))))
	v1486 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1475)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+660)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v390)+656)) = v1484 | v1485<<(uint(int32(16))%32)
	v1495 = F_psprintf(m, int32(_a_F_bt_index_check_callback_37), v390+int32(656))
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L5
	} else {
		goto L391
	}
L391:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L5
	} else {
		goto L392
	}
L392:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L5
	} else {
		goto L393
	}
L393:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1504)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+640)) = v1505 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_44), v390+int32(640))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L5
	} else {
		goto L394
	}
L394:
	;
	v1514 = *(*int64)(unsafe.Add(mBase, uint32(v391)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+636)) = uint32(v1514)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+628)) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v390)+624)) = v1482
	v1519 = int64(base.Ui64(v1514) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+632)) = uint32(v1519)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_45), v390+int32(624))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L5
	} else {
		goto L395
	}
L395:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1401), int32(_a_F_bt_index_check_callback_35))
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L5
	} else {
		goto L396
	}
L396:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L397:
	;
	v1536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+7)))
	if v1536&int32(32) == int32(0) {
		v1551 = int32(2712)
		goto L336
	} else {
		goto L398
	}
L398:
	;
	v1542 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+4)))
	if v1542&int32(_a_F_bt_index_check_callback_8) != 0 {
		v1551 = int32(2712)
		goto L336
	} else {
		goto L399
	}
L399:
	;
	if v1542&int32(_a_F_bt_index_check_callback_36) != 0 {
		goto L400
	} else {
		goto L401
	}
L400:
	;
	v1549 = int32(2712)
	goto L402
L401:
	;
	v1549 = int32(2704)
	goto L402
L402:
	;
	v1551 = v1549
	goto L336
L403:
	;
	v1553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+10)))
	if v1553 != int32(1) {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+8))
	v1687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+8)))
	if v1687 != int32(1) {
		goto L425
	} else {
		goto L426
	}
L405:
	;
	v1556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955)+12)))
	if v1556&int32(1) == int32(0) {
		goto L404
	} else {
		goto L406
	}
L406:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v987)))
	v1562 = int32(_a_F_bt_index_check_callback_46)
	if v1561&v1562 == v1562 {
		goto L404
	} else {
		goto L407
	}
L407:
	;
	v1566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+7)))
	if v1566&int32(32) == int32(0) {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v1644 = F_bt_normalize_tuple(m, v391, v995)
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L5
	} else {
		goto L422
	}
L409:
	;
	v1571 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+4)))
	if v1571&int32(_a_F_bt_index_check_callback_8) == int32(0) {
		goto L408
	} else {
		goto L410
	}
L410:
	;
	if v1571&int32(4095) == int32(0) {
		goto L404
	} else {
		goto L411
	}
L411:
	;
	v1582 = int32(0)
	goto L412
L412:
	;
	v1613 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+2)))
	v1614 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995))))
	v1623 = F__bt_form_posting(m, v995, v1613+(v995+v1614<<(uint(int32(16))%32))+v1582*int32(6), int32(1))
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L5
	} else {
		goto L414
	}
L413:
	;
	goto L404
L414:
	;
	v1625 = F_bt_normalize_tuple(m, v391, v1623)
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L5
	} else {
		goto L415
	}
L415:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v391)+60))
	v1628 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1625)+6)))
	F_bloom_add_element(m, v1627, v1625, v1628&int32(_a_F_bt_index_check_callback_38))
	mBase = m.M
	if v1625 != v1623 {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	F_pfree(m, v1625)
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		goto L5
	} else {
		goto L419
	}
L417:
	;
	goto L418
L418:
	;
	F_pfree(m, v1623)
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L5
	} else {
		goto L420
	}
L419:
	;
	goto L418
L420:
	;
	v1638 = v1582 + int32(1)
	v1639 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+4)))
	if base.Ui32(v1638) < base.Ui32(v1639&int32(4095)) {
		v1582 = v1638
		goto L412
	} else {
		goto L421
	}
L421:
	;
	goto L413
L422:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v391)+60))
	v1647 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1644)+6)))
	F_bloom_add_element(m, v1646, v1644, v1647&int32(_a_F_bt_index_check_callback_38))
	mBase = m.M
	if v995 == v1644 {
		goto L404
	} else {
		goto L423
	}
L423:
	;
	F_pfree(m, v1644)
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L5
	} else {
		goto L424
	}
L424:
	;
	goto L404
L425:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v955)+4))
	if v1715 == int32(0) {
		goto L429
	} else {
		goto L430
	}
L426:
	;
	v1690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+7)))
	if v1690&int32(32) == int32(0) {
		goto L425
	} else {
		goto L427
	}
L427:
	;
	v1695 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+4)))
	if v1695&int32(_a_F_bt_index_check_callback_8) == int32(0) {
		goto L425
	} else {
		goto L428
	}
L428:
	;
	v1700 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+2)))
	v1701 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995))))
	v1708 = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v1439)+8)) = v1700 + (v995 + v1701<<(uint(int32(16))%32)) + v1695&int32(4095)*v1708 - v1708
	goto L425
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1439)+8)) = v1686
	v1735 = v967 + int32(1)
	v1737 = v1735 & int32(_a_F_bt_index_check_callback_31)
	v1738 = base.B2i32(base.Ui32(v945) < base.Ui32(v1737))
	if v1738 == int32(0) {
		goto L438
	} else {
		goto L439
	}
L430:
	;
	v1718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955)+12)))
	if v1718&int32(1) != 0 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v1724 = F__bt_compare(m, v1721, v1439, v1722, int32(1))
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L5
	} else {
		goto L434
	}
L432:
	;
	goto L433
L433:
	;
	v1729 = F_invariant_l_offset(m, v391, v1439, int32(1))
	mBase = m.M
	v1730 = m.ExcPending
	if v1730 != 0 {
		goto L5
	} else {
		goto L436
	}
L434:
	;
	if v1724 <= int32(0) {
		goto L429
	} else {
		goto L435
	}
L435:
	;
	goto L68
L436:
	;
	if v1729 == int32(0) {
		goto L68
	} else {
		goto L437
	}
L437:
	;
	goto L429
L438:
	;
	v1741 = F_invariant_l_offset(m, v391, v1439, v1737)
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L5
	} else {
		goto L441
	}
L439:
	;
	goto L440
L440:
	;
	v1745 = int32(0)
	v1746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+12)))
	if v1746 != int32(1) {
		v1826 = v1745
		goto L443
	} else {
		goto L444
	}
L441:
	;
	if v1741 == int32(0) {
		goto L275
	} else {
		goto L442
	}
L442:
	;
	goto L440
L443:
	;
	if v986 != v945 {
		v2508 = v955
		goto L272
	} else {
		goto L469
	}
L444:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v391)+24))
	v1750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1749)+116)))
	if v1750 != int32(1) {
		v1826 = v1745
		goto L443
	} else {
		goto L445
	}
L445:
	;
	v1753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955)+12)))
	if v1753&int32(1) == int32(0) {
		v1826 = v1745
		goto L443
	} else {
		goto L446
	}
L446:
	;
	v1758 = int32(0)
	v1760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1439)+2)))
	if v1760 != 0 {
		v1796 = v1758
		v1797 = v1758
		goto L447
	} else {
		goto L448
	}
L447:
	;
	if v1796|v1738 != 0 {
		v1826 = v1797
		goto L443
	} else {
		goto L459
	}
L448:
	;
	v1761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+7)))
	if v1761&int32(32) != 0 {
		goto L450
	} else {
		goto L451
	}
L449:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	F_bt_entry_unique_check(m, v391, v995, v1776, v986, v390+int32(1120))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L5
	} else {
		goto L456
	}
L450:
	;
	v1764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+5)))
	if v1764&int32(32) != 0 {
		goto L449
	} else {
		goto L453
	}
L451:
	;
	goto L452
L452:
	;
	v1767 = int32(0)
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v390)+1132))
	if v1768 == v1767 {
		v1796 = v1758
		v1797 = v1767
		goto L447
	} else {
		goto L454
	}
L453:
	;
	goto L452
L454:
	;
	v1771 = int32(0)
	v1772 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1768)+4)))
	if v1772 == v1771 {
		v1796 = v1758
		v1797 = v1771
		goto L447
	} else {
		goto L455
	}
L455:
	;
	goto L449
L456:
	;
	v1781 = int32(1)
	v1782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+12)))
	if v1782 != v1781 {
		v1826 = v1781
		goto L443
	} else {
		goto L457
	}
L457:
	;
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v391)+24))
	v1786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1785)+116)))
	if v1786 != int32(1) {
		v1826 = v1781
		goto L443
	} else {
		goto L458
	}
L458:
	;
	v1789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955)+12)))
	v1790 = int32(1)
	v1796 = base.B2i32(v1789&v1790 == int32(0))
	v1797 = v1790
	goto L447
L459:
	;
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1439)+8)) = int32(0)
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v1804 = F__bt_compare(m, v1802, v1439, v1803, v1737)
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L5
	} else {
		goto L462
	}
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1439)+8)) = v1799
	v1826 = v1824
	goto L443
L461:
	;
	if v1797 != 0 {
		v1824 = int32(1)
		goto L460
	} else {
		goto L467
	}
L462:
	;
	if v1804 == int32(0) {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	v1808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1439)+2)))
	if v1808 != int32(1) {
		goto L461
	} else {
		goto L466
	}
L464:
	;
	goto L465
L465:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v390)+1128)) = int64(4294967295)
	v1813 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v390)+1124)) = uint16(v1813)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+1120)) = int32(-1)
	v1824 = v1797
	goto L460
L466:
	;
	goto L465
L467:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	F_bt_entry_unique_check(m, v391, v995, v1818, v986, v390+int32(1120))
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L5
	} else {
		goto L468
	}
L468:
	;
	v1824 = int32(0)
	goto L460
L469:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v1830 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1829)+16)))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1829+v1830)+4))
	if v1832 == int32(0) {
		goto L472
	} else {
		goto L473
	}
L470:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v1994 = m.ExcPending
	if v1994 != 0 {
		goto L5
	} else {
		goto L516
	}
L471:
	;
	v1990 = int32(0)
	v2390 = v1990
	v2395 = v1964
	v2398 = v1990
	goto L273
L472:
	;
	v1964 = int32(0)
	goto L471
L473:
	;
	goto L474
L474:
	;
	v1837 = v1832
	goto L475
L475:
	;
	v1869 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[5]))
	if v1869 != 0 {
		goto L477
	} else {
		goto L478
	}
L476:
	;
	v1918 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1872)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1918) {
		goto L494
	} else {
		goto L495
	}
L477:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L5
	} else {
		goto L480
	}
L478:
	;
	goto L479
L479:
	;
	v1872 = F_palloc_btree_page(m, v391, v1837)
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L5
	} else {
		goto L482
	}
L480:
	;
	goto L479
L481:
	;
	goto L476
L482:
	;
	v1874 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1872)+16)))
	v1875 = v1872 + v1874
	v1876 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1875)+12)))
	if v1876&int32(20) == int32(0) {
		goto L481
	} else {
		goto L483
	}
L483:
	;
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+4))
	if v1881 == int32(0) {
		goto L481
	} else {
		goto L484
	}
L484:
	;
	v1886 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		goto L5
	} else {
		goto L485
	}
L485:
	;
	if v1886 != 0 {
		goto L486
	} else {
		goto L487
	}
L486:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L5
	} else {
		goto L489
	}
L487:
	;
	goto L488
L488:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+4))
	F_pfree(m, v1872)
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L5
	} else {
		goto L493
	}
L489:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+48))
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+372)) = v1837
	*(*int32)(unsafe.Add(mBase, uint32(v390)+368)) = v1893
	*(*int32)(unsafe.Add(mBase, uint32(v390)+376)) = v1892 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_47), v390+int32(368))
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L5
	} else {
		goto L490
	}
L490:
	;
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_48), int32(0))
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L5
	} else {
		goto L491
	}
L491:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1926), int32(_a_F_bt_index_check_callback_49))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L5
	} else {
		goto L492
	}
L492:
	;
	goto L488
L493:
	;
	v1837 = v1915
	goto L475
L494:
	;
	v1926 = int32(base.Ui32(v1918+int32(_a_F_bt_index_check_callback_30)) >> (uint(int32(2)) % 32))
	goto L496
L495:
	;
	v1926 = int32(0)
	goto L496
L496:
	;
	if v1876&int32(1) != 0 {
		goto L498
	} else {
		goto L499
	}
L497:
	;
	v1953 = int32(0)
	v1956 = F_errstart(m, int32(13), v1953)
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L5
	} else {
		goto L514
	}
L498:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+4))
	if v1931 != 0 {
		goto L501
	} else {
		goto L502
	}
L499:
	;
	goto L500
L500:
	;
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+4))
	if v1944 != 0 {
		goto L509
	} else {
		goto L510
	}
L501:
	;
	v1932 = int32(2)
	goto L503
L502:
	;
	v1932 = int32(1)
	goto L503
L503:
	;
	if base.Ui32(v1926&int32(_a_F_bt_index_check_callback_31)) < base.Ui32(v1932) {
		goto L497
	} else {
		goto L504
	}
L504:
	;
	v1936 = F_PageGetItemIdCareful_2(m, v391, v1837, v1872, v1932)
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L5
	} else {
		goto L505
	}
L505:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+4))
	if v1940 != 0 {
		goto L506
	} else {
		goto L507
	}
L506:
	;
	v1941 = int32(2)
	goto L508
L507:
	;
	v1941 = int32(1)
	goto L508
L508:
	;
	v2318 = v1936
	v2319 = v1941
	goto L274
L509:
	;
	v1945 = int32(3)
	goto L511
L510:
	;
	v1945 = int32(2)
	goto L511
L511:
	;
	if base.Ui32(v1926&int32(_a_F_bt_index_check_callback_31)) < base.Ui32(v1945) {
		goto L497
	} else {
		goto L512
	}
L512:
	;
	v1950 = F_PageGetItemIdCareful_2(m, v391, v1837, v1872, v1945)
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L5
	} else {
		goto L513
	}
L513:
	;
	v2318 = v1950
	v2319 = int32(0)
	goto L274
L514:
	;
	if v1956 != 0 {
		goto L470
	} else {
		goto L515
	}
L515:
	;
	v1964 = v1953
	goto L471
L516:
	;
	v1995 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1875)+12)))
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+324)) = v1837
	*(*int32)(unsafe.Add(mBase, uint32(v390)+328)) = v1997 + int32(4)
	v2002 = int32(0)
	if v1995&int32(1) != 0 {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v2007 = int32(_a_F_bt_index_check_callback_32)
	goto L519
L518:
	;
	v2007 = int32(_a_F_bt_index_check_callback_33)
	goto L519
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+320)) = v2002 + v2007
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_50), v390+int32(320))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L5
	} else {
		goto L520
	}
L520:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(2057), int32(_a_F_bt_index_check_callback_49))
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L5
	} else {
		goto L521
	}
L521:
	;
	v2390 = v2002
	v2395 = v1953
	v2398 = int32(0)
	goto L273
L522:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L5
	} else {
		goto L523
	}
L523:
	;
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+800)) = v2029 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_51), v390+int32(800))
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L5
	} else {
		goto L524
	}
L524:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v987)))
	v2039 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v2040 = *(*int64)(unsafe.Add(mBase, uint32(v391)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+788)) = uint32(v2040)
	v2043 = int64(base.Ui64(v2040) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+784)) = uint32(v2043)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+776)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v390)+772)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v390)+768)) = v2039
	*(*int32)(unsafe.Add(mBase, uint32(v390)+780)) = int32(base.Ui32(v2038) >> (uint(int32(17)) % 32))
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_52), v390+int32(768))
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L5
	} else {
		goto L525
	}
L525:
	;
	F_errhint(m, int32(_a_F_bt_index_check_callback_53), int32(0))
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L5
	} else {
		goto L526
	}
L526:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1329), int32(_a_F_bt_index_check_callback_35))
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L5
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L5
	} else {
		goto L529
	}
L529:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L5
	} else {
		goto L530
	}
L530:
	;
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(v2080)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+592)) = v2081 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_54), v390+int32(592))
	mBase = m.M
	v2089 = m.ExcPending
	if v2089 != 0 {
		goto L5
	} else {
		goto L531
	}
L531:
	;
	v2090 = *(*int64)(unsafe.Add(mBase, uint32(v391)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+588)) = uint32(v2090)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+580)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v390)+576)) = v2071
	v2095 = int64(base.Ui64(v2090) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+584)) = uint32(v2095)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_55), v390+int32(576))
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L5
	} else {
		goto L532
	}
L532:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1430), int32(_a_F_bt_index_check_callback_35))
	mBase = m.M
	v2106 = m.ExcPending
	if v2106 != 0 {
		goto L5
	} else {
		goto L533
	}
L533:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L534:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+196)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v390)+192)) = v2124
	v2130 = F_psprintf(m, int32(_a_F_bt_index_check_callback_37), v390+int32(192))
	mBase = m.M
	v2131 = m.ExcPending
	if v2131 != 0 {
		goto L5
	} else {
		goto L538
	}
L535:
	;
	goto L534
L536:
	;
	v2112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+5)))
	if v2112&int32(32) == int32(0) {
		v2123 = v995
		goto L535
	} else {
		goto L537
	}
L537:
	;
	v2117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+2)))
	v2118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995))))
	v2123 = v2117 + (v995 + v2118<<(uint(int32(16))%32))
	goto L535
L538:
	;
	v2132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2123)+2)))
	v2133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2123))))
	v2134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2123)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+180)) = v2134
	*(*int32)(unsafe.Add(mBase, uint32(v390)+176)) = v2132 | v2133<<(uint(int32(16))%32)
	v2143 = F_psprintf(m, int32(_a_F_bt_index_check_callback_37), v390+int32(176))
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L5
	} else {
		goto L539
	}
L539:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2148 = m.ExcPending
	if v2148 != 0 {
		goto L5
	} else {
		goto L540
	}
L540:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L5
	} else {
		goto L541
	}
L541:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v2152)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+160)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v390)+164)) = v2153 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_56), v390+int32(160))
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L5
	} else {
		goto L542
	}
L542:
	;
	v2163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v955)+12)))
	v2164 = *(*int64)(unsafe.Add(mBase, uint32(v391)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+144)) = uint32(v2164)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+136)) = v2143
	*(*int32)(unsafe.Add(mBase, uint32(v390)+128)) = v2130
	v2169 = int64(base.Ui64(v2164) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+140)) = uint32(v2169)
	if v2163&int32(1) != 0 {
		goto L543
	} else {
		goto L544
	}
L543:
	;
	v2175 = int32(_a_F_bt_index_check_callback_40)
	goto L545
L544:
	;
	v2175 = int32(_a_F_bt_index_check_callback_41)
	goto L545
L545:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+132)) = v2175
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_57), v390+int32(128))
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L5
	} else {
		goto L546
	}
L546:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1485), int32(_a_F_bt_index_check_callback_35))
	mBase = m.M
	v2186 = m.ExcPending
	if v2186 != 0 {
		goto L5
	} else {
		goto L547
	}
L547:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L548:
	;
	v2195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+7)))
	if v2195&int32(32) == int32(0) {
		v2211 = v995
		goto L550
	} else {
		goto L551
	}
L549:
	;
	v2212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2211)+2)))
	v2213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2211))))
	v2214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2211)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+468)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v390)+464)) = v2212 | v2213<<(uint(int32(16))%32)
	v2223 = F_psprintf(m, int32(_a_F_bt_index_check_callback_37), v390+int32(464))
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L5
	} else {
		goto L553
	}
L550:
	;
	goto L549
L551:
	;
	v2200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+5)))
	if v2200&int32(32) == int32(0) {
		v2211 = v995
		goto L550
	} else {
		goto L552
	}
L552:
	;
	v2205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+2)))
	v2206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995))))
	v2211 = v2205 + (v995 + v2206<<(uint(int32(16))%32))
	goto L550
L553:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v2227 = v1735 & int32(_a_F_bt_index_check_callback_31)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+452)) = v2227
	*(*int32)(unsafe.Add(mBase, uint32(v390)+448)) = v2225
	v2233 = F_psprintf(m, int32(_a_F_bt_index_check_callback_37), v390+int32(448))
	mBase = m.M
	v2234 = m.ExcPending
	if v2234 != 0 {
		goto L5
	} else {
		goto L554
	}
L554:
	;
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v2237 = F_PageGetItemIdCareful_2(m, v391, v2235, v2236, v2227)
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L5
	} else {
		goto L555
	}
L555:
	;
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2237)))
	v2243 = v2239 + v2240&int32(_a_F_bt_index_check_callback_19)
	v2244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2243)+7)))
	if v2244&int32(32) == int32(0) {
		v2260 = v2243
		goto L557
	} else {
		goto L558
	}
L556:
	;
	v2261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2260)+2)))
	v2262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2260))))
	v2263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2260)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+436)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v390)+432)) = v2261 | v2262<<(uint(int32(16))%32)
	v2272 = F_psprintf(m, int32(_a_F_bt_index_check_callback_37), v390+int32(432))
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L5
	} else {
		goto L560
	}
L557:
	;
	goto L556
L558:
	;
	v2249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2243)+5)))
	if v2249&int32(32) == int32(0) {
		v2260 = v2243
		goto L557
	} else {
		goto L559
	}
L559:
	;
	v2254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2243)+2)))
	v2255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2243))))
	v2260 = v2254 + (v2243 + v2255<<(uint(int32(16))%32))
	goto L557
L560:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2277 = m.ExcPending
	if v2277 != 0 {
		goto L5
	} else {
		goto L561
	}
L561:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L5
	} else {
		goto L562
	}
L562:
	;
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(v2281)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+416)) = v2282 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_58), v390+int32(416))
	mBase = m.M
	v2290 = m.ExcPending
	if v2290 != 0 {
		goto L5
	} else {
		goto L563
	}
L563:
	;
	v2291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v955)+12)))
	v2292 = *(*int64)(unsafe.Add(mBase, uint32(v391)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+412)) = uint32(v2292)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+404)) = v2272
	v2296 = int64(base.Ui64(v2292) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+408)) = uint32(v2296)
	if v2291&int32(1) != 0 {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v2302 = int32(_a_F_bt_index_check_callback_40)
	goto L566
L565:
	;
	v2302 = int32(_a_F_bt_index_check_callback_41)
	goto L566
L566:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+400)) = v2302
	*(*int32)(unsafe.Add(mBase, uint32(v390)+396)) = v2233
	*(*int32)(unsafe.Add(mBase, uint32(v390)+392)) = v2223
	*(*int32)(unsafe.Add(mBase, uint32(v390)+384)) = v2193
	*(*int32)(unsafe.Add(mBase, uint32(v390)+388)) = v2302
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_59), v390+int32(384))
	mBase = m.M
	v2312 = m.ExcPending
	if v2312 != 0 {
		goto L5
	} else {
		goto L567
	}
L567:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1641), int32(_a_F_bt_index_check_callback_35))
	mBase = m.M
	v2317 = m.ExcPending
	if v2317 != 0 {
		goto L5
	} else {
		goto L568
	}
L568:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L569:
	;
	v2329 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2327)+4)) = uint8(v2329)
	v2331 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v2333 = F__bt_compare(m, v2331, v2327, v2332, v945)
	mBase = m.M
	v2334 = m.ExcPending
	if v2334 != 0 {
		goto L5
	} else {
		goto L570
	}
L570:
	;
	v2336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2327))))
	if int32(0)-(v2336^int32(1)) < v2333 {
		v2390 = v2327
		v2395 = int32(1)
		v2398 = v2319
		goto L273
	} else {
		goto L571
	}
L571:
	;
	v2341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+9)))
	if v2341 == int32(0) {
		goto L572
	} else {
		goto L573
	}
L572:
	;
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v2345 = F_palloc_btree_page(m, v391, v2344)
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L5
	} else {
		goto L575
	}
L573:
	;
	goto L574
L574:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		goto L5
	} else {
		goto L577
	}
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+32)) = v2345
	v2348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2345)+16)))
	v2350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2345+v2348)+12)))
	if v2350&int32(20) != 0 {
		v2921 = v510
		v2935 = v511
		goto L110
	} else {
		goto L576
	}
L576:
	;
	goto L574
L577:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2360 = m.ExcPending
	if v2360 != 0 {
		goto L5
	} else {
		goto L578
	}
L578:
	;
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v2361)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+352)) = v2362 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_60), v390+int32(352))
	mBase = m.M
	v2370 = m.ExcPending
	if v2370 != 0 {
		goto L5
	} else {
		goto L579
	}
L579:
	;
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v2372 = *(*int64)(unsafe.Add(mBase, uint32(v391)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+348)) = uint32(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+340)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v390)+336)) = v2371
	v2377 = int64(base.Ui64(v2372) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+344)) = uint32(v2377)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_61), v390+int32(336))
	mBase = m.M
	v2383 = m.ExcPending
	if v2383 != 0 {
		goto L5
	} else {
		goto L580
	}
L580:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1753), int32(_a_F_bt_index_check_callback_35))
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L5
	} else {
		goto L581
	}
L581:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L582:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(v391)+24))
	v2425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2424)+116)))
	if v2395&v2425 != int32(1) {
		v2508 = v955
		goto L272
	} else {
		goto L583
	}
L583:
	;
	v2429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955)+12)))
	if v2429&int32(1) == int32(0) {
		v2508 = v955
		goto L272
	} else {
		goto L584
	}
L584:
	;
	v2434 = *(*int32)(unsafe.Add(mBase, uint32(v955)+4))
	if v2434 == int32(0) {
		v2508 = v955
		goto L272
	} else {
		goto L585
	}
L585:
	;
	v2439 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
		goto L5
	} else {
		goto L586
	}
L586:
	;
	if v2439 != 0 {
		goto L587
	} else {
		goto L588
	}
L587:
	;
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_62), int32(0))
	mBase = m.M
	v2444 = m.ExcPending
	if v2444 != 0 {
		goto L5
	} else {
		goto L590
	}
L588:
	;
	goto L589
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2390)+8)) = int32(0)
	v2452 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v2454 = F__bt_compare(m, v2452, v2390, v2453, v945)
	mBase = m.M
	v2455 = m.ExcPending
	if v2455 != 0 {
		goto L5
	} else {
		goto L592
	}
L590:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1765), int32(_a_F_bt_index_check_callback_35))
	mBase = m.M
	v2449 = m.ExcPending
	if v2449 != 0 {
		goto L5
	} else {
		goto L591
	}
L591:
	;
	goto L589
L592:
	;
	if v2454 != 0 {
		v2508 = v955
		goto L272
	} else {
		goto L593
	}
L593:
	;
	v2456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2390)+2)))
	if v2456 != 0 {
		v2508 = v955
		goto L272
	} else {
		goto L594
	}
L594:
	;
	if v1826 == int32(0) {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	F_bt_entry_unique_check(m, v391, v995, v2459, v945, v390+int32(1120))
	mBase = m.M
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L5
	} else {
		goto L598
	}
L596:
	;
	goto L597
L597:
	;
	v2466 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L5
	} else {
		goto L599
	}
L598:
	;
	goto L597
L599:
	;
	if v2466 != 0 {
		goto L600
	} else {
		goto L601
	}
L600:
	;
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_63), int32(0))
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L5
	} else {
		goto L603
	}
L601:
	;
	goto L602
L602:
	;
	v2477 = F_palloc_btree_page(m, v391, v2434)
	mBase = m.M
	v2478 = m.ExcPending
	if v2478 != 0 {
		goto L5
	} else {
		goto L605
	}
L603:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1788), int32(_a_F_bt_index_check_callback_35))
	mBase = m.M
	v2476 = m.ExcPending
	if v2476 != 0 {
		goto L5
	} else {
		goto L604
	}
L604:
	;
	goto L602
L605:
	;
	v2479 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2477)+16)))
	v2480 = v2477 + v2479
	v2481 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2480)+12)))
	if v2481&int32(20) != 0 {
		goto L262
	} else {
		goto L606
	}
L606:
	;
	if v2481&int32(1) == int32(0) {
		goto L270
	} else {
		goto L607
	}
L607:
	;
	v2488 = F_PageGetItemIdCareful_2(m, v391, v2434, v2477, v2398)
	mBase = m.M
	v2489 = m.ExcPending
	if v2489 != 0 {
		goto L5
	} else {
		goto L608
	}
L608:
	;
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(v2488)))
	F_bt_entry_unique_check(m, v391, v2477+v2490&int32(_a_F_bt_index_check_callback_19), v2434, v2398, v390+int32(1120))
	mBase = m.M
	v2497 = m.ExcPending
	if v2497 != 0 {
		goto L5
	} else {
		goto L609
	}
L609:
	;
	F_pfree(m, v2477)
	mBase = m.M
	v2499 = m.ExcPending
	if v2499 != 0 {
		goto L5
	} else {
		goto L610
	}
L610:
	;
	v2508 = v2480
	goto L272
L611:
	;
	v2535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+9)))
	if v2535 != int32(1) {
		v2739 = v2508
		goto L271
	} else {
		goto L612
	}
L612:
	;
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v2540 = F_PageGetItemIdCareful_2(m, v391, v2538, v2539, v986)
	mBase = m.M
	v2541 = m.ExcPending
	if v2541 != 0 {
		goto L5
	} else {
		goto L613
	}
L613:
	;
	v2542 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v2543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2542)+16)))
	v2544 = *(*int32)(unsafe.Add(mBase, uint32(v2540)))
	v2547 = v2542 + v2544&int32(_a_F_bt_index_check_callback_19)
	v2548 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2547))))
	v2551 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2547)+2)))
	v2552 = v2548<<(uint(int32(16))%32) | v2551
	v2553 = F_palloc_btree_page(m, v391, v2552)
	mBase = m.M
	v2554 = m.ExcPending
	if v2554 != 0 {
		goto L5
	} else {
		goto L614
	}
L614:
	;
	v2555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2553)+16)))
	v2556 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2553)+12)))
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(v2542+v2543)+8))
	F_bt_child_highkey_check(m, v391, v986, v2553, v2558)
	mBase = m.M
	v2560 = m.ExcPending
	if v2560 != 0 {
		goto L5
	} else {
		goto L615
	}
L615:
	;
	v2561 = v2553 + v2555
	v2562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2561)+12)))
	if v2562&int32(4) != 0 {
		goto L263
	} else {
		goto L616
	}
L616:
	;
	v2567 = *(*int32)(unsafe.Add(mBase, uint32(v2561)+4))
	if v2567 != 0 {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	v2568 = int32(2)
	goto L619
L618:
	;
	v2568 = int32(1)
	goto L619
L619:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v2556) {
		goto L620
	} else {
		goto L621
	}
L620:
	;
	v2576 = int32(base.Ui32(v2556+int32(_a_F_bt_index_check_callback_30)) >> (uint(int32(2)) % 32))
	goto L622
L621:
	;
	v2576 = int32(0)
	goto L622
L622:
	;
	v2578 = v2576 & int32(_a_F_bt_index_check_callback_31)
	if base.Ui32(v2568) <= base.Ui32(v2578) {
		goto L623
	} else {
		goto L624
	}
L623:
	;
	v2580 = v2568
	goto L626
L624:
	;
	goto L625
L625:
	;
	F_pfree(m, v2553)
	mBase = m.M
	v2730 = m.ExcPending
	if v2730 != 0 {
		goto L5
	} else {
		goto L671
	}
L626:
	;
	v2612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2561)+12)))
	if v2612&int32(1) == int32(0) {
		goto L629
	} else {
		goto L630
	}
L627:
	;
	goto L625
L628:
	;
	v2693 = v2580 + int32(1)
	if base.Ui32(v2693&int32(_a_F_bt_index_check_callback_31)) <= base.Ui32(v2578) {
		v2580 = v2693
		goto L626
	} else {
		goto L670
	}
L629:
	;
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v2561)+4))
	if v2621 != 0 {
		goto L632
	} else {
		goto L633
	}
L630:
	;
	goto L631
L631:
	;
	v2625 = v2580 & int32(_a_F_bt_index_check_callback_31)
	v2626 = F_PageGetItemIdCareful_2(m, v391, v2552, v2553, v2625)
	mBase = m.M
	v2627 = m.ExcPending
	if v2627 != 0 {
		goto L5
	} else {
		goto L636
	}
L632:
	;
	v2622 = int32(2)
	goto L634
L633:
	;
	v2622 = int32(1)
	goto L634
L634:
	;
	if v2580&int32(_a_F_bt_index_check_callback_31) == v2622 {
		goto L628
	} else {
		goto L635
	}
L635:
	;
	goto L631
L636:
	;
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v2629 = F__bt_compare(m, v2628, v1439, v2553, v2625)
	mBase = m.M
	v2630 = m.ExcPending
	if v2630 != 0 {
		goto L5
	} else {
		goto L637
	}
L637:
	;
	v2631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1439))))
	if v2631 == int32(0) {
		goto L638
	} else {
		goto L639
	}
L638:
	;
	if int32(0) < v2629 {
		goto L69
	} else {
		goto L641
	}
L639:
	;
	goto L640
L640:
	;
	if v2629 == int32(0) {
		goto L643
	} else {
		goto L644
	}
L641:
	;
	goto L628
L642:
	;
	if v2673 <= v2678 {
		goto L69
	} else {
		goto L669
	}
L643:
	;
	v2638 = *(*int32)(unsafe.Add(mBase, uint32(v2626)))
	v2641 = v2553 + v2638&int32(_a_F_bt_index_check_callback_19)
	v2643 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2553)+16)))
	v2644 = v2553 + v2643
	v2645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2644)+12)))
	if v2645&int32(1) != 0 {
		goto L646
	} else {
		goto L647
	}
L644:
	;
	goto L645
L645:
	;
	if v2629 < int32(0) {
		goto L628
	} else {
		goto L668
	}
L646:
	;
	v2650 = *(*int32)(unsafe.Add(mBase, uint32(v2644)+4))
	if v2650 != 0 {
		goto L649
	} else {
		goto L650
	}
L647:
	;
	v2653 = int32(0)
	goto L648
L648:
	;
	v2654 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v2655 = *(*int32)(unsafe.Add(mBase, uint32(v2654)+192))
	v2656 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2655)+10)))
	v2657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2641)+7)))
	if v2657&int32(32) != 0 {
		goto L654
	} else {
		goto L655
	}
L649:
	;
	v2651 = int32(2)
	goto L651
L650:
	;
	v2651 = int32(1)
	goto L651
L651:
	;
	v2653 = base.B2i32(base.Ui32(v2651) <= base.Ui32(v2625))
	goto L648
L652:
	;
	v2676 = F_BTreeTupleGetHeapTIDCareful(m, v391, v2641, v2653)
	mBase = m.M
	v2677 = m.ExcPending
	if v2677 != 0 {
		goto L5
	} else {
		goto L665
	}
L653:
	;
	v2673 = v2671
	goto L652
L654:
	;
	v2660 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2641)+4)))
	if v2660&int32(_a_F_bt_index_check_callback_8) != 0 {
		goto L657
	} else {
		goto L658
	}
L655:
	;
	goto L656
L656:
	;
	v2669 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2655)+8)))
	if v2656 < v2669 {
		v2673 = v2656
		goto L652
	} else {
		goto L664
	}
L657:
	;
	v2663 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2655)+8)))
	if v2663 <= v2656 {
		v2671 = v2663
		goto L653
	} else {
		goto L660
	}
L658:
	;
	goto L659
L659:
	;
	v2666 = v2660 & int32(4095)
	if v2656 < v2666 {
		goto L661
	} else {
		goto L662
	}
L660:
	;
	v2673 = v2656
	goto L652
L661:
	;
	v2668 = v2656
	goto L663
L662:
	;
	v2668 = v2666
	goto L663
L663:
	;
	v2673 = v2668
	goto L652
L664:
	;
	v2671 = v2669
	goto L653
L665:
	;
	v2678 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+12))
	if v2678 != v2673 {
		goto L642
	} else {
		goto L666
	}
L666:
	;
	v2680 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+8))
	if v2680|base.B2i32(v2676 == int32(0)) != 0 {
		goto L69
	} else {
		goto L667
	}
L667:
	;
	goto L628
L668:
	;
	goto L69
L669:
	;
	goto L628
L670:
	;
	goto L627
L671:
	;
	v2739 = v2508
	goto L271
L672:
	;
	v2878 = v2739
	goto L203
L673:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2774 = m.ExcPending
	if v2774 != 0 {
		goto L5
	} else {
		goto L674
	}
L674:
	;
	v2775 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(v2775)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+304)) = v2776 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_64), v390+int32(304))
	mBase = m.M
	v2784 = m.ExcPending
	if v2784 != 0 {
		goto L5
	} else {
		goto L675
	}
L675:
	;
	v2785 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v2786 = *(*int64)(unsafe.Add(mBase, uint32(v391)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+296)) = uint32(v2786)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+288)) = v2785
	v2790 = int64(base.Ui64(v2786) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+292)) = uint32(v2790)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_65), v390+int32(288))
	mBase = m.M
	v2796 = m.ExcPending
	if v2796 != 0 {
		goto L5
	} else {
		goto L676
	}
L676:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1806), int32(_a_F_bt_index_check_callback_35))
	mBase = m.M
	v2801 = m.ExcPending
	if v2801 != 0 {
		goto L5
	} else {
		goto L677
	}
L677:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L678:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2808 = m.ExcPending
	if v2808 != 0 {
		goto L5
	} else {
		goto L679
	}
L679:
	;
	v2809 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v2810 = *(*int32)(unsafe.Add(mBase, uint32(v2809)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+272)) = v2810 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_66), v390+int32(272))
	mBase = m.M
	v2818 = m.ExcPending
	if v2818 != 0 {
		goto L5
	} else {
		goto L680
	}
L680:
	;
	v2819 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v2820 = *(*int64)(unsafe.Add(mBase, uint32(v391)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+268)) = uint32(v2820)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+260)) = v2552
	*(*int32)(unsafe.Add(mBase, uint32(v390)+256)) = v2819
	v2825 = int64(base.Ui64(v2820) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+264)) = uint32(v2825)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_67), v390+int32(256))
	mBase = m.M
	v2831 = m.ExcPending
	if v2831 != 0 {
		goto L5
	} else {
		goto L681
	}
L681:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(2498), int32(_a_F_bt_index_check_callback_68))
	mBase = m.M
	v2836 = m.ExcPending
	if v2836 != 0 {
		goto L5
	} else {
		goto L682
	}
L682:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L683:
	;
	v2878 = v2480
	goto L203
L684:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2845 = m.ExcPending
	if v2845 != 0 {
		goto L5
	} else {
		goto L685
	}
L685:
	;
	v2846 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v2847 = *(*int32)(unsafe.Add(mBase, uint32(v2846)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+896)) = v2847 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_69), v390+int32(896))
	mBase = m.M
	v2855 = m.ExcPending
	if v2855 != 0 {
		goto L5
	} else {
		goto L686
	}
L686:
	;
	v2856 = *(*int32)(unsafe.Add(mBase, uint32(v431)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+888)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v390)+884)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v390)+880)) = v386
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_70), v390+int32(880))
	mBase = m.M
	v2864 = m.ExcPending
	if v2864 != 0 {
		goto L5
	} else {
		goto L687
	}
L687:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(781), int32(_a_F_bt_index_check_callback_17))
	mBase = m.M
	v2869 = m.ExcPending
	if v2869 != 0 {
		goto L5
	} else {
		goto L688
	}
L688:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L689:
	;
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(v2878)+4))
	if v2905 != 0 {
		v2921 = v510
		v2935 = v511
		goto L110
	} else {
		goto L690
	}
L690:
	;
	v2906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+9)))
	if v2906 != int32(1) {
		v2921 = v510
		v2935 = v511
		goto L110
	} else {
		goto L691
	}
L691:
	;
	v2909 = int32(0)
	v2911 = *(*int32)(unsafe.Add(mBase, uint32(v2878)+8))
	F_bt_child_highkey_check(m, v391, v2909, v2909, v2911)
	mBase = m.M
	v2913 = m.ExcPending
	if v2913 != 0 {
		goto L5
	} else {
		goto L692
	}
L692:
	;
	v2921 = v510
	v2935 = v511
	goto L110
L693:
	;
	v2947 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	if v386 == v2947 {
		goto L87
	} else {
		goto L694
	}
L694:
	;
	v2949 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	v2950 = *(*int32)(unsafe.Add(mBase, uint32(v391)+48))
	if v2950 != 0 {
		goto L695
	} else {
		goto L696
	}
L695:
	;
	F_pfree(m, v2950)
	mBase = m.M
	v2952 = m.ExcPending
	if v2952 != 0 {
		goto L5
	} else {
		goto L698
	}
L696:
	;
	goto L697
L697:
	;
	v2955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+9)))
	if v2955 != int32(1) {
		goto L699
	} else {
		goto L700
	}
L698:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+48)) = int32(0)
	goto L697
L699:
	;
	v2986 = *(*int32)(unsafe.Add(mBase, uint32(v391)+16))
	F_MemoryContextReset(m, v2986)
	mBase = m.M
	v2988 = m.ExcPending
	if v2988 != 0 {
		goto L5
	} else {
		goto L705
	}
L700:
	;
	v2958 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v2958 == int32(0) {
		goto L699
	} else {
		goto L701
	}
L701:
	;
	v2961 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v2962 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v2964 = F_PageGetItemIdCareful_2(m, v391, v2961, v2962, int32(1))
	mBase = m.M
	v2965 = m.ExcPending
	if v2965 != 0 {
		goto L5
	} else {
		goto L702
	}
L702:
	;
	v2966 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v2967 = *(*int32)(unsafe.Add(mBase, uint32(v2964)))
	v2970 = v2966 + v2967&int32(_a_F_bt_index_check_callback_19)
	v2971 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2970)+6)))
	v2974 = F_MemoryContextAlloc(m, v411, v2971&int32(_a_F_bt_index_check_callback_38))
	mBase = m.M
	v2975 = m.ExcPending
	if v2975 != 0 {
		goto L5
	} else {
		goto L703
	}
L703:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+48)) = v2974
	v2977 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2970)+6)))
	v2979 = v2977 & int32(_a_F_bt_index_check_callback_38)
	if v2979 == int32(0) {
		goto L699
	} else {
		goto L704
	}
L704:
	;
	base.MemoryCopy(m, v2974, v2970, v2979)
	goto L699
L705:
	;
	if v2949 != 0 {
		__phi386 = v2949
		__phi388 = v386
		__phi393 = v2921
		__phi407 = v2935
		v386 = __phi386
		v388 = __phi388
		v393 = __phi393
		v407 = __phi407
		goto L103
	} else {
		goto L706
	}
L706:
	;
	goto L104
L707:
	;
	F_pfree(m, v2989)
	mBase = m.M
	v2991 = m.ExcPending
	if v2991 != 0 {
		goto L5
	} else {
		goto L710
	}
L708:
	;
	goto L709
L709:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[4])) = v411
	if v2935 != int32(-1) {
		v319 = v2935
		v323 = v390
		v324 = v391
		v326 = v2921
		v332 = v399
		v333 = v400
		v338 = v2921
		v341 = int32(0)
		v343 = v410
		goto L88
	} else {
		goto L711
	}
L710:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+48)) = int32(0)
	goto L709
L711:
	;
	goto L89
L712:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3005 = m.ExcPending
	if v3005 != 0 {
		goto L5
	} else {
		goto L713
	}
L713:
	;
	v3006 = *(*int32)(unsafe.Add(mBase, uint32(v400)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+116)) = v412
	*(*int32)(unsafe.Add(mBase, uint32(v390)+112)) = v3006 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_71), v390+int32(112))
	mBase = m.M
	v3015 = m.ExcPending
	if v3015 != 0 {
		goto L5
	} else {
		goto L714
	}
L714:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(537), int32(_a_F_bt_index_check_callback_3))
	mBase = m.M
	v3020 = m.ExcPending
	if v3020 != 0 {
		goto L5
	} else {
		goto L715
	}
L715:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L716:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3027 = m.ExcPending
	if v3027 != 0 {
		goto L5
	} else {
		goto L717
	}
L717:
	;
	v3028 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v3029 = *(*int32)(unsafe.Add(mBase, uint32(v3028)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+96)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v390)+100)) = v3029 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_72), v390+int32(96))
	mBase = m.M
	v3038 = m.ExcPending
	if v3038 != 0 {
		goto L5
	} else {
		goto L718
	}
L718:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(793), int32(_a_F_bt_index_check_callback_17))
	mBase = m.M
	v3043 = m.ExcPending
	if v3043 != 0 {
		goto L5
	} else {
		goto L719
	}
L719:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L720:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3050 = m.ExcPending
	if v3050 != 0 {
		goto L5
	} else {
		goto L721
	}
L721:
	;
	v3051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+1072)) = v3051 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_73), v96+int32(1072))
	mBase = m.M
	v3059 = m.ExcPending
	if v3059 != 0 {
		goto L5
	} else {
		goto L722
	}
L722:
	;
	F_errhint(m, int32(_a_F_bt_index_check_callback_74), int32(0))
	mBase = m.M
	v3063 = m.ExcPending
	if v3063 != 0 {
		goto L5
	} else {
		goto L723
	}
L723:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(487), int32(_a_F_bt_index_check_callback_3))
	mBase = m.M
	v3068 = m.ExcPending
	if v3068 != 0 {
		goto L5
	} else {
		goto L724
	}
L724:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L725:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3489 = m.ExcPending
	if v3489 != 0 {
		goto L5
	} else {
		goto L784
	}
L726:
	;
	v3072 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v3073 = F_BuildIndexInfo(m, v3072)
	mBase = m.M
	v3074 = m.ExcPending
	if v3074 != 0 {
		goto L5
	} else {
		goto L729
	}
L727:
	;
	goto L728
L728:
	;
	v3477 = *(*int32)(unsafe.Add(mBase, uint32(v324)+28))
	if v3477 != 0 {
		goto L779
	} else {
		goto L780
	}
L729:
	;
	v3077 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[11]))
	v3079 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_bt_index_check_callback[12])))
	if v3079&int32(1) != 0 {
		goto L730
	} else {
		goto L731
	}
L730:
	;
	v3082 = int32(0)
	goto L732
L731:
	;
	v3082 = v3077
	goto L732
L732:
	;
	if v3082 != 0 {
		goto L725
	} else {
		goto L733
	}
L733:
	;
	v3083 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	v3084 = *(*int32)(unsafe.Add(mBase, uint32(v324)+28))
	v3085 = int32(0)
	v3089 = *(*int32)(unsafe.Add(mBase, uint32(v3083)+188))
	v3090 = *(*int32)(unsafe.Add(mBase, uint32(v3089)+8))
	v3091 = m.T0[v3090].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v3083, v3084, v3085, v3085, v3085, int32(449))
	mBase = m.M
	v3092 = m.ExcPending
	if v3092 != 0 {
		goto L5
	} else {
		goto L734
	}
L734:
	;
	v3093 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3073)+116)) = uint8(v3093)
	v3095 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3073)+121)) = uint8(v3095)
	*(*int32)(unsafe.Add(mBase, uint32(v3073)+100)) = v3093
	*(*int64)(unsafe.Add(mBase, uint32(v3073)+92)) = int64(0)
	v3103 = F_errstart(m, int32(14), v3093)
	mBase = m.M
	v3104 = m.ExcPending
	if v3104 != 0 {
		goto L5
	} else {
		goto L735
	}
L735:
	;
	if v3103 != 0 {
		goto L736
	} else {
		goto L737
	}
L736:
	;
	v3105 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v3106 = *(*int32)(unsafe.Add(mBase, uint32(v3105)+48))
	v3107 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	v3108 = *(*int32)(unsafe.Add(mBase, uint32(v3107)+48))
	v3109 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v323)+36)) = v3108 + v3109
	*(*int32)(unsafe.Add(mBase, uint32(v323)+32)) = v3106 + v3109
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_75), v323+int32(32))
	mBase = m.M
	v3119 = m.ExcPending
	if v3119 != 0 {
		goto L5
	} else {
		goto L739
	}
L737:
	;
	goto L738
L738:
	;
	v3126 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	v3127 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v3129 = int32(0)
	v3134 = *(*int32)(unsafe.Add(mBase, uint32(v3126)+188))
	v3135 = *(*int32)(unsafe.Add(mBase, uint32(v3134)+140))
	v3136 = m.T0[v3135].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v3126, v3127, v3073, int32(1), v3129, v3129, v3129, int32(-1), int32(_a_F_bt_index_check_callback_76), v324, v3091)
	mBase = m.M
	v3137 = m.ExcPending
	if v3137 != 0 {
		goto L5
	} else {
		goto L741
	}
L739:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(588), int32(_a_F_bt_index_check_callback_3))
	mBase = m.M
	v3124 = m.ExcPending
	if v3124 != 0 {
		goto L5
	} else {
		goto L740
	}
L740:
	;
	goto L738
L741:
	;
	v3140 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3141 = m.ExcPending
	if v3141 != 0 {
		goto L5
	} else {
		goto L742
	}
L742:
	;
	if v3140 != 0 {
		goto L743
	} else {
		goto L744
	}
L743:
	;
	v3142 = *(*int64)(unsafe.Add(mBase, uint32(v324)+64))
	v3143 = *(*int32)(unsafe.Add(mBase, uint32(v343)+48))
	v3144 = int64(0)
	v3145 = *(*int32)(unsafe.Add(mBase, uint32(v324)+60))
	v3147 = v3145 + int32(24)
	v3148 = *(*int64)(unsafe.Add(mBase, uint32(v3145)+16))
	v3151 = base.I32_wrap_i64(int64(base.Ui64(v3148) >> (uint(int64(3)) % 64)))
	if v3151 <= int32(7) {
		goto L747
	} else {
		goto L748
	}
L744:
	;
	goto L745
L745:
	;
	v3442 = *(*int32)(unsafe.Add(mBase, uint32(v324)+60))
	F_pfree(m, v3442)
	mBase = m.M
	v3444 = m.ExcPending
	if v3444 != 0 {
		goto L5
	} else {
		goto L778
	}
L746:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v323)+16)) = base.F64_mul(base.F64_div(base.F64_convert_i64_u(v3388), base.F64_convert_i64_u(v3389)), float64(100))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+8)) = v3143 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v323))) = v3142
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_77), v323)
	mBase = m.M
	v3404 = m.ExcPending
	if v3404 != 0 {
		goto L5
	} else {
		goto L776
	}
L747:
	;
	if v3151 == int32(0) {
		v3388 = v3144
		v3389 = v3148
		goto L746
	} else {
		goto L750
	}
L748:
	;
	goto L749
L749:
	;
	v3287 = int64(0)
	v3288 = int32(0)
	if v3151 == v3288 {
		goto L762
	} else {
		goto L763
	}
L750:
	;
	v3157 = v3151 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v3151) {
		goto L751
	} else {
		goto L752
	}
L751:
	;
	v3165 = v3147
	v3166 = int32(0)
	v3191 = v3144
	goto L754
L752:
	;
	v3216 = v3147
	v3242 = v3144
	goto L753
L753:
	;
	v3248 = int32(0)
	v3249 = v3216
	v3275 = v3242
	goto L758
L754:
	;
	v3195 = int32(4)
	v3196 = v3165 + v3195
	v3197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3165)+3)))
	v3198 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3197)+uint32(_c_F_bt_index_check_callback[13]))))
	v3199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3165)+2)))
	v3200 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3199)+uint32(_c_F_bt_index_check_callback[13]))))
	v3201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3165)+1)))
	v3202 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3201)+uint32(_c_F_bt_index_check_callback[13]))))
	v3203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3165))))
	v3204 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3203)+uint32(_c_F_bt_index_check_callback[13]))))
	v3208 = v3198 + (v3200 + (v3202 + (v3191 + v3204)))
	v3210 = v3166 + v3195
	if v3210 != v3151&int32(-4) {
		v3165 = v3196
		v3166 = v3210
		v3191 = v3208
		goto L754
	} else {
		goto L756
	}
L755:
	;
	if v3157 == int32(0) {
		v3388 = v3208
		v3389 = v3148
		goto L746
	} else {
		goto L757
	}
L756:
	;
	goto L755
L757:
	;
	v3216 = v3196
	v3242 = v3208
	goto L753
L758:
	;
	v3279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3249))))
	v3280 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3279)+uint32(_c_F_bt_index_check_callback[13]))))
	v3281 = v3275 + v3280
	v3282 = int32(1)
	v3285 = v3248 + v3282
	if v3285 != v3157 {
		v3248 = v3285
		v3249 = v3249 + v3282
		v3275 = v3281
		goto L758
	} else {
		goto L760
	}
L759:
	;
	v3388 = v3281
	v3389 = v3148
	goto L746
L760:
	;
	goto L759
L761:
	;
	v3359 = *(*int64)(unsafe.Add(mBase, uint32(v3145)+16))
	v3388 = v3358
	v3389 = v3359
	goto L746
L762:
	;
	v3358 = int64(0)
	goto L761
L763:
	;
	goto L764
L764:
	;
	v3295 = v3151 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v3151) {
		goto L766
	} else {
		goto L767
	}
L765:
	;
	v3358 = v3348
	goto L761
L766:
	;
	v3300 = v3147
	v3302 = v3287
	v3305 = v3288
	goto L769
L767:
	;
	v3325 = v3147
	v3327 = v3287
	goto L768
L768:
	;
	v3332 = v3325
	v3333 = int32(0)
	v3334 = v3327
	goto L773
L769:
	;
	v3306 = int32(4)
	v3307 = v3300 + v3306
	v3308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3300)+3)))
	v3309 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3308)+uint32(_c_F_bt_index_check_callback[13]))))
	v3310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3300)+2)))
	v3311 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3310)+uint32(_c_F_bt_index_check_callback[13]))))
	v3312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3300)+1)))
	v3313 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3312)+uint32(_c_F_bt_index_check_callback[13]))))
	v3314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3300))))
	v3315 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3314)+uint32(_c_F_bt_index_check_callback[13]))))
	v3319 = v3309 + (v3311 + (v3313 + (v3302 + v3315)))
	v3321 = v3305 + v3306
	if v3321 != v3151&int32(-4) {
		v3300 = v3307
		v3302 = v3319
		v3305 = v3321
		goto L769
	} else {
		goto L771
	}
L770:
	;
	if v3295 == int32(0) {
		v3348 = v3319
		goto L765
	} else {
		goto L772
	}
L771:
	;
	goto L770
L772:
	;
	v3325 = v3307
	v3327 = v3319
	goto L768
L773:
	;
	v3338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3332))))
	v3339 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3338)+uint32(_c_F_bt_index_check_callback[13]))))
	v3340 = v3334 + v3339
	v3341 = int32(1)
	v3344 = v3333 + v3341
	if v3344 != v3295 {
		v3332 = v3332 + v3341
		v3333 = v3344
		v3334 = v3340
		goto L773
	} else {
		goto L775
	}
L774:
	;
	v3348 = v3340
	goto L765
L775:
	;
	goto L774
L776:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(596), int32(_a_F_bt_index_check_callback_3))
	mBase = m.M
	v3409 = m.ExcPending
	if v3409 != 0 {
		goto L5
	} else {
		goto L777
	}
L777:
	;
	goto L745
L778:
	;
	goto L728
L779:
	;
	F_UnregisterSnapshot(m, v3477)
	mBase = m.M
	v3479 = m.ExcPending
	if v3479 != 0 {
		goto L5
	} else {
		goto L782
	}
L780:
	;
	goto L781
L781:
	;
	v3480 = *(*int32)(unsafe.Add(mBase, uint32(v324)+16))
	F_MemoryContextDelete(m, v3480)
	mBase = m.M
	v3482 = m.ExcPending
	if v3482 != 0 {
		goto L5
	} else {
		goto L783
	}
L782:
	;
	goto L781
L783:
	;
	m.G0 = v323 + int32(1168)
	goto L66
L784:
	;
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_78), int32(0))
	mBase = m.M
	v3493 = m.ExcPending
	if v3493 != 0 {
		goto L5
	} else {
		goto L785
	}
L785:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_79), int32(931), int32(_a_F_bt_index_check_callback_80))
	mBase = m.M
	v3498 = m.ExcPending
	if v3498 != 0 {
		goto L5
	} else {
		goto L786
	}
L786:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L787:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3510 = m.ExcPending
	if v3510 != 0 {
		goto L5
	} else {
		goto L788
	}
L788:
	;
	v3511 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v3512 = *(*int32)(unsafe.Add(mBase, uint32(v3511)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+240)) = v3512 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_81), v390+int32(240))
	mBase = m.M
	v3520 = m.ExcPending
	if v3520 != 0 {
		goto L5
	} else {
		goto L789
	}
L789:
	;
	v3521 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v3522 = *(*int64)(unsafe.Add(mBase, uint32(v391)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+224)) = uint32(v3522)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+216)) = v2580 & int32(_a_F_bt_index_check_callback_31)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+212)) = v2552
	*(*int32)(unsafe.Add(mBase, uint32(v390)+208)) = v3521
	v3530 = int64(base.Ui64(v3522) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+220)) = uint32(v3530)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_82), v390+int32(208))
	mBase = m.M
	v3536 = m.ExcPending
	if v3536 != 0 {
		goto L5
	} else {
		goto L790
	}
L790:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(2539), int32(_a_F_bt_index_check_callback_68))
	mBase = m.M
	v3541 = m.ExcPending
	if v3541 != 0 {
		goto L5
	} else {
		goto L791
	}
L791:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L792:
	;
	v3559 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+564)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v390)+560)) = v3559
	v3565 = F_psprintf(m, int32(_a_F_bt_index_check_callback_37), v390+int32(560))
	mBase = m.M
	v3566 = m.ExcPending
	if v3566 != 0 {
		goto L5
	} else {
		goto L796
	}
L793:
	;
	goto L792
L794:
	;
	v3547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+5)))
	if v3547&int32(32) == int32(0) {
		v3558 = v995
		goto L793
	} else {
		goto L795
	}
L795:
	;
	v3552 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+2)))
	v3553 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995))))
	v3558 = v3552 + (v995 + v3553<<(uint(int32(16))%32))
	goto L793
L796:
	;
	v3567 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3558)+2)))
	v3568 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3558))))
	v3569 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3558)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+548)) = v3569
	*(*int32)(unsafe.Add(mBase, uint32(v390)+544)) = v3567 | v3568<<(uint(int32(16))%32)
	v3578 = F_psprintf(m, int32(_a_F_bt_index_check_callback_37), v390+int32(544))
	mBase = m.M
	v3579 = m.ExcPending
	if v3579 != 0 {
		goto L5
	} else {
		goto L797
	}
L797:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3583 = m.ExcPending
	if v3583 != 0 {
		goto L5
	} else {
		goto L798
	}
L798:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3586 = m.ExcPending
	if v3586 != 0 {
		goto L5
	} else {
		goto L799
	}
L799:
	;
	v3587 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v3588 = *(*int32)(unsafe.Add(mBase, uint32(v3587)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+528)) = v3588 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_83), v390+int32(528))
	mBase = m.M
	v3596 = m.ExcPending
	if v3596 != 0 {
		goto L5
	} else {
		goto L800
	}
L800:
	;
	v3597 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v955)+12)))
	v3598 = *(*int64)(unsafe.Add(mBase, uint32(v391)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+512)) = uint32(v3598)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+504)) = v3578
	*(*int32)(unsafe.Add(mBase, uint32(v390)+496)) = v3565
	v3603 = int64(base.Ui64(v3598) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+508)) = uint32(v3603)
	if v3597&int32(1) != 0 {
		goto L801
	} else {
		goto L802
	}
L801:
	;
	v3609 = int32(_a_F_bt_index_check_callback_40)
	goto L803
L802:
	;
	v3609 = int32(_a_F_bt_index_check_callback_41)
	goto L803
L803:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+500)) = v3609
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_57), v390+int32(496))
	mBase = m.M
	v3615 = m.ExcPending
	if v3615 != 0 {
		goto L5
	} else {
		goto L804
	}
L804:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1592), int32(_a_F_bt_index_check_callback_35))
	mBase = m.M
	v3620 = m.ExcPending
	if v3620 != 0 {
		goto L5
	} else {
		goto L805
	}
L805:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L806:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3629 = m.ExcPending
	if v3629 != 0 {
		goto L5
	} else {
		goto L807
	}
L807:
	;
	v3630 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v3631 = *(*int32)(unsafe.Add(mBase, uint32(v3630)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+848)) = v3631 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_84), v390+int32(848))
	mBase = m.M
	v3639 = m.ExcPending
	if v3639 != 0 {
		goto L5
	} else {
		goto L808
	}
L808:
	;
	v3640 = *(*int32)(unsafe.Add(mBase, uint32(v391)+36))
	v3643 = v3621 + v3622&int32(_a_F_bt_index_check_callback_19)
	v3644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3643)+7)))
	if v3644&int32(32) == int32(0) {
		goto L810
	} else {
		goto L811
	}
L809:
	;
	v3660 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v798)+12)))
	v3661 = *(*int64)(unsafe.Add(mBase, uint32(v391)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+832)) = uint32(v3661)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+820)) = v3659
	*(*int32)(unsafe.Add(mBase, uint32(v390)+816)) = v3640
	v3666 = int64(base.Ui64(v3661) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v390)+828)) = uint32(v3666)
	if v3660&int32(1) != 0 {
		goto L813
	} else {
		goto L814
	}
L810:
	;
	v3655 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v3656 = *(*int32)(unsafe.Add(mBase, uint32(v3655)+192))
	v3657 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3656)+8)))
	v3659 = v3657
	goto L809
L811:
	;
	v3649 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3643)+4)))
	if v3649&int32(_a_F_bt_index_check_callback_8) != 0 {
		goto L810
	} else {
		goto L812
	}
L812:
	;
	v3659 = v3649 & int32(4095)
	goto L809
L813:
	;
	v3672 = int32(_a_F_bt_index_check_callback_40)
	goto L815
L814:
	;
	v3672 = int32(_a_F_bt_index_check_callback_41)
	goto L815
L815:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+824)) = v3672
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_85), v390+int32(816))
	mBase = m.M
	v3678 = m.ExcPending
	if v3678 != 0 {
		goto L5
	} else {
		goto L816
	}
L816:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1280), int32(_a_F_bt_index_check_callback_35))
	mBase = m.M
	v3683 = m.ExcPending
	if v3683 != 0 {
		goto L5
	} else {
		goto L817
	}
L817:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L818:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(348), int32(_a_F_bt_index_check_callback_86))
	mBase = m.M
	v3836 = m.ExcPending
	if v3836 != 0 {
		goto L5
	} else {
		goto L834
	}
L819:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3787 = m.ExcPending
	if v3787 != 0 {
		goto L5
	} else {
		goto L831
	}
L820:
	;
	v3691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v3695 = int32(0)
	goto L821
L821:
	;
	v3728 = *(*int32)(unsafe.Add(mBase, uint32(v3691+v3695<<(uint(int32(2))%32))))
	if v3728 != int32(1982) {
		goto L823
	} else {
		goto L824
	}
L822:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3737 = m.ExcPending
	if v3737 != 0 {
		goto L5
	} else {
		goto L827
	}
L823:
	;
	v3732 = v3695 + int32(1)
	if v3688 != v3732 {
		v3695 = v3732
		goto L821
	} else {
		goto L826
	}
L824:
	;
	goto L825
L825:
	;
	goto L822
L826:
	;
	goto L819
L827:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3740 = m.ExcPending
	if v3740 != 0 {
		goto L5
	} else {
		goto L828
	}
L828:
	;
	v3741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v3741 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_87), v35)
	mBase = m.M
	v3747 = m.ExcPending
	if v3747 != 0 {
		goto L5
	} else {
		goto L829
	}
L829:
	;
	F_errhint(m, int32(_a_F_bt_index_check_callback_88), int32(0))
	mBase = m.M
	v3751 = m.ExcPending
	if v3751 != 0 {
		goto L5
	} else {
		goto L830
	}
L830:
	;
	goto L818
L831:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3790 = m.ExcPending
	if v3790 != 0 {
		goto L5
	} else {
		goto L832
	}
L832:
	;
	v3791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v3791 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_87), v35+int32(16))
	mBase = m.M
	v3799 = m.ExcPending
	if v3799 != 0 {
		goto L5
	} else {
		goto L833
	}
L833:
	;
	goto L818
L834:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L835:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3843 = m.ExcPending
	if v3843 != 0 {
		goto L5
	} else {
		goto L836
	}
L836:
	;
	v3844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+48)) = v3844 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_89), v35+int32(48))
	mBase = m.M
	v3852 = m.ExcPending
	if v3852 != 0 {
		goto L5
	} else {
		goto L837
	}
L837:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(322), int32(_a_F_bt_index_check_callback_86))
	mBase = m.M
	v3857 = m.ExcPending
	if v3857 != 0 {
		goto L5
	} else {
		goto L838
	}
L838:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L839:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3864 = m.ExcPending
	if v3864 != 0 {
		goto L5
	} else {
		goto L840
	}
L840:
	;
	v3865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+32)) = v3865 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_90), v35+int32(32))
	mBase = m.M
	v3873 = m.ExcPending
	if v3873 != 0 {
		goto L5
	} else {
		goto L841
	}
L841:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(330), int32(_a_F_bt_index_check_callback_86))
	mBase = m.M
	v3878 = m.ExcPending
	if v3878 != 0 {
		goto L5
	} else {
		goto L842
	}
L842:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
