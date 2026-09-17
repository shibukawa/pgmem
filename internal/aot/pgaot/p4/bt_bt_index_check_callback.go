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
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
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
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
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
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 float32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v243 int64
	_ = v243
	var v257 int64
	_ = v257
	var v258 int64
	_ = v258
	var v260 int64
	_ = v260
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
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
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v378 int64
	_ = v378
	var v379 int64
	_ = v379
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var __phi464 int32
	_ = __phi464
	var v465 int32
	_ = v465
	var __phi465 int32
	_ = __phi465
	var v468 int32
	_ = v468
	var __phi468 int32
	_ = __phi468
	var v469 int32
	_ = v469
	var __phi469 int32
	_ = __phi469
	var v471 int32
	_ = v471
	var __phi471 int32
	_ = __phi471
	var v478 int32
	_ = v478
	var __phi478 int32
	_ = __phi478
	var v481 int32
	_ = v481
	var __phi481 int32
	_ = __phi481
	var v483 int32
	_ = v483
	var __phi483 int32
	_ = __phi483
	var v485 int32
	_ = v485
	var __phi485 int32
	_ = __phi485
	var v486 int32
	_ = v486
	var __phi486 int32
	_ = __phi486
	var v489 int32
	_ = v489
	var __phi489 int32
	_ = __phi489
	var v490 int32
	_ = v490
	var __phi490 int32
	_ = __phi490
	var v491 int32
	_ = v491
	var __phi491 int32
	_ = __phi491
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int64
	_ = v504
	var v505 int64
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v862 int64
	_ = v862
	var v865 int64
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v893 int32
	_ = v893
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v969 int32
	_ = v969
	var v977 int32
	_ = v977
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1004 int32
	_ = v1004
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1035 int32
	_ = v1035
	var v1047 int32
	_ = v1047
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
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
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1135 int32
	_ = v1135
	var v1143 int32
	_ = v1143
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1170 int32
	_ = v1170
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1219 int32
	_ = v1219
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
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int64
	_ = v1255
	var v1258 int64
	_ = v1258
	var v1267 int32
	_ = v1267
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
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
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1383 int32
	_ = v1383
	var v1388 int32
	_ = v1388
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1412 int32
	_ = v1412
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1530 int32
	_ = v1530
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1543 int32
	_ = v1543
	var v1548 int32
	_ = v1548
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1597 int32
	_ = v1597
	var v1598 int64
	_ = v1598
	var v1603 int64
	_ = v1603
	var v1609 int32
	_ = v1609
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1620 int32
	_ = v1620
	var v1626 int32
	_ = v1626
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1650 int32
	_ = v1650
	var v1655 int32
	_ = v1655
	var v1667 int32
	_ = v1667
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1716 int32
	_ = v1716
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1739 int32
	_ = v1739
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1781 int32
	_ = v1781
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1794 int32
	_ = v1794
	var v1801 int32
	_ = v1801
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1899 int32
	_ = v1899
	var v1904 int32
	_ = v1904
	var v1908 int32
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1925 int32
	_ = v1925
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1967 int32
	_ = v1967
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1989 int32
	_ = v1989
	var v1993 int32
	_ = v1993
	var v1998 int32
	_ = v1998
	var v2001 int32
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2012 int32
	_ = v2012
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2050 int32
	_ = v2050
	var v2076 int32
	_ = v2076
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2088 int32
	_ = v2088
	var v2093 int32
	_ = v2093
	var v2100 int32
	_ = v2100
	var v2105 int32
	_ = v2105
	var v2110 int32
	_ = v2110
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2126 int64
	_ = v2126
	var v2129 int64
	_ = v2129
	var v2141 int32
	_ = v2141
	var v2145 int32
	_ = v2145
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2162 int32
	_ = v2162
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2175 int32
	_ = v2175
	var v2176 int64
	_ = v2176
	var v2181 int64
	_ = v2181
	var v2187 int32
	_ = v2187
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2198 int32
	_ = v2198
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2234 int32
	_ = v2234
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2250 int64
	_ = v2250
	var v2255 int64
	_ = v2255
	var v2261 int32
	_ = v2261
	var v2267 int32
	_ = v2267
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2286 int32
	_ = v2286
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2313 int32
	_ = v2313
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2335 int32
	_ = v2335
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2363 int32
	_ = v2363
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2378 int64
	_ = v2378
	var v2382 int64
	_ = v2382
	var v2388 int32
	_ = v2388
	var v2398 int32
	_ = v2398
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2422 int32
	_ = v2422
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2436 int32
	_ = v2436
	var v2443 int32
	_ = v2443
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2448 int32
	_ = v2448
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2458 int64
	_ = v2458
	var v2463 int64
	_ = v2463
	var v2469 int32
	_ = v2469
	var v2474 int32
	_ = v2474
	var v2478 int32
	_ = v2478
	var v2481 int32
	_ = v2481
	var v2484 int32
	_ = v2484
	var v2507 int32
	_ = v2507
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2515 int32
	_ = v2515
	var v2520 int32
	_ = v2520
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2530 int32
	_ = v2530
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
	var v2545 int32
	_ = v2545
	var v2549 int32
	_ = v2549
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2557 int32
	_ = v2557
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2564 int32
	_ = v2564
	var v2565 int32
	_ = v2565
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2583 int32
	_ = v2583
	var v2585 int32
	_ = v2585
	var v2594 int32
	_ = v2594
	var v2618 int32
	_ = v2618
	var v2621 int32
	_ = v2621
	var v2624 int32
	_ = v2624
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
	var v2633 int32
	_ = v2633
	var v2634 int32
	_ = v2634
	var v2637 int32
	_ = v2637
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2646 int32
	_ = v2646
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2662 int32
	_ = v2662
	var v2664 int32
	_ = v2664
	var v2667 int32
	_ = v2667
	var v2698 int32
	_ = v2698
	var v2707 int32
	_ = v2707
	var v2708 int32
	_ = v2708
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2715 int32
	_ = v2715
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2724 int32
	_ = v2724
	var v2727 int32
	_ = v2727
	var v2729 int32
	_ = v2729
	var v2730 int32
	_ = v2730
	var v2731 int32
	_ = v2731
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2746 int32
	_ = v2746
	var v2749 int32
	_ = v2749
	var v2752 int32
	_ = v2752
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2757 int32
	_ = v2757
	var v2759 int32
	_ = v2759
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2766 int32
	_ = v2766
	var v2779 int32
	_ = v2779
	var v2816 int32
	_ = v2816
	var v2825 int32
	_ = v2825
	var v2850 int32
	_ = v2850
	var v2857 int32
	_ = v2857
	var v2860 int32
	_ = v2860
	var v2861 int32
	_ = v2861
	var v2862 int32
	_ = v2862
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2872 int64
	_ = v2872
	var v2876 int64
	_ = v2876
	var v2882 int32
	_ = v2882
	var v2887 int32
	_ = v2887
	var v2891 int32
	_ = v2891
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2904 int32
	_ = v2904
	var v2905 int32
	_ = v2905
	var v2906 int64
	_ = v2906
	var v2911 int64
	_ = v2911
	var v2917 int32
	_ = v2917
	var v2922 int32
	_ = v2922
	var v2924 int32
	_ = v2924
	var v2928 int32
	_ = v2928
	var v2931 int32
	_ = v2931
	var v2932 int32
	_ = v2932
	var v2933 int32
	_ = v2933
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2950 int32
	_ = v2950
	var v2955 int32
	_ = v2955
	var v2964 int32
	_ = v2964
	var v2988 int32
	_ = v2988
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2995 int32
	_ = v2995
	var v2997 int32
	_ = v2997
	var v2999 int32
	_ = v2999
	var v3007 int32
	_ = v3007
	var v3021 int32
	_ = v3021
	var v3033 int32
	_ = v3033
	var v3035 int32
	_ = v3035
	var v3036 int32
	_ = v3036
	var v3038 int32
	_ = v3038
	var v3041 int32
	_ = v3041
	var v3044 int32
	_ = v3044
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3050 int32
	_ = v3050
	var v3051 int32
	_ = v3051
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3056 int32
	_ = v3056
	var v3057 int32
	_ = v3057
	var v3060 int32
	_ = v3060
	var v3061 int32
	_ = v3061
	var v3063 int32
	_ = v3063
	var v3065 int32
	_ = v3065
	var v3072 int32
	_ = v3072
	var v3074 int32
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3077 int32
	_ = v3077
	var v3088 int32
	_ = v3088
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3101 int32
	_ = v3101
	var v3106 int32
	_ = v3106
	var v3110 int32
	_ = v3110
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3115 int32
	_ = v3115
	var v3124 int32
	_ = v3124
	var v3129 int32
	_ = v3129
	var v3133 int32
	_ = v3133
	var v3136 int32
	_ = v3136
	var v3137 int32
	_ = v3137
	var v3145 int32
	_ = v3145
	var v3149 int32
	_ = v3149
	var v3154 int32
	_ = v3154
	var v3158 int32
	_ = v3158
	var v3161 int32
	_ = v3161
	var v3162 int32
	_ = v3162
	var v3170 int32
	_ = v3170
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
	var v3182 int32
	_ = v3182
	var v3183 int32
	_ = v3183
	var v3184 int32
	_ = v3184
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3190 int32
	_ = v3190
	var v3191 int32
	_ = v3191
	var v3192 int32
	_ = v3192
	var v3194 int32
	_ = v3194
	var v3202 int32
	_ = v3202
	var v3203 int32
	_ = v3203
	var v3204 int32
	_ = v3204
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3207 int32
	_ = v3207
	var v3208 int32
	_ = v3208
	var v3218 int32
	_ = v3218
	var v3223 int32
	_ = v3223
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3228 int32
	_ = v3228
	var v3233 int32
	_ = v3233
	var v3234 int32
	_ = v3234
	var v3235 float64
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3239 int32
	_ = v3239
	var v3240 int32
	_ = v3240
	var v3241 int64
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3243 int64
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3246 int32
	_ = v3246
	var v3247 int64
	_ = v3247
	var v3250 int32
	_ = v3250
	var v3256 int32
	_ = v3256
	var v3264 int32
	_ = v3264
	var v3265 int32
	_ = v3265
	var v3290 int64
	_ = v3290
	var v3294 int32
	_ = v3294
	var v3295 int32
	_ = v3295
	var v3296 int32
	_ = v3296
	var v3297 int64
	_ = v3297
	var v3298 int32
	_ = v3298
	var v3299 int64
	_ = v3299
	var v3300 int32
	_ = v3300
	var v3301 int64
	_ = v3301
	var v3302 int32
	_ = v3302
	var v3303 int64
	_ = v3303
	var v3307 int64
	_ = v3307
	var v3309 int32
	_ = v3309
	var v3315 int32
	_ = v3315
	var v3341 int64
	_ = v3341
	var v3346 int32
	_ = v3346
	var v3348 int32
	_ = v3348
	var v3374 int64
	_ = v3374
	var v3378 int32
	_ = v3378
	var v3379 int64
	_ = v3379
	var v3380 int64
	_ = v3380
	var v3381 int32
	_ = v3381
	var v3384 int32
	_ = v3384
	var v3390 int64
	_ = v3390
	var v3400 int32
	_ = v3400
	var v3404 int32
	_ = v3404
	var v3406 int32
	_ = v3406
	var v3411 int32
	_ = v3411
	var v3412 int32
	_ = v3412
	var v3415 int32
	_ = v3415
	var v3417 int64
	_ = v3417
	var v3418 int32
	_ = v3418
	var v3419 int32
	_ = v3419
	var v3421 int32
	_ = v3421
	var v3422 int32
	_ = v3422
	var v3425 int32
	_ = v3425
	var v3428 int32
	_ = v3428
	var v3431 int32
	_ = v3431
	var v3437 int64
	_ = v3437
	var v3439 int32
	_ = v3439
	var v3443 int32
	_ = v3443
	var v3444 int32
	_ = v3444
	var v3449 int64
	_ = v3449
	var v3451 int32
	_ = v3451
	var v3452 int32
	_ = v3452
	var v3453 int32
	_ = v3453
	var v3457 int64
	_ = v3457
	var v3458 int32
	_ = v3458
	var v3459 int32
	_ = v3459
	var v3461 int32
	_ = v3461
	var v3462 int32
	_ = v3462
	var v3465 int64
	_ = v3465
	var v3467 int32
	_ = v3467
	var v3469 int32
	_ = v3469
	var v3470 int32
	_ = v3470
	var v3475 int64
	_ = v3475
	var v3479 int32
	_ = v3479
	var v3483 int32
	_ = v3483
	var v3485 int32
	_ = v3485
	var v3487 int32
	_ = v3487
	var v3489 int64
	_ = v3489
	var v3490 int32
	_ = v3490
	var v3491 int32
	_ = v3491
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3495 int64
	_ = v3495
	var v3496 int64
	_ = v3496
	var v3498 int32
	_ = v3498
	var v3500 int32
	_ = v3500
	var v3502 int32
	_ = v3502
	var v3506 int64
	_ = v3506
	var v3509 int32
	_ = v3509
	var v3511 int32
	_ = v3511
	var v3515 int64
	_ = v3515
	var v3516 int32
	_ = v3516
	var v3517 int64
	_ = v3517
	var v3518 int32
	_ = v3518
	var v3519 int64
	_ = v3519
	var v3520 int32
	_ = v3520
	var v3521 int64
	_ = v3521
	var v3522 int32
	_ = v3522
	var v3523 int64
	_ = v3523
	var v3527 int64
	_ = v3527
	var v3528 int32
	_ = v3528
	var v3531 int32
	_ = v3531
	var v3538 int64
	_ = v3538
	var v3539 int64
	_ = v3539
	var v3568 int64
	_ = v3568
	var v3569 int64
	_ = v3569
	var v3584 int32
	_ = v3584
	var v3589 int32
	_ = v3589
	var v3622 int32
	_ = v3622
	var v3624 int32
	_ = v3624
	var v3657 int32
	_ = v3657
	var v3659 int32
	_ = v3659
	var v3660 int32
	_ = v3660
	var v3662 int32
	_ = v3662
	var v3674 int32
	_ = v3674
	var v3677 int32
	_ = v3677
	var v3678 int32
	_ = v3678
	var v3679 int32
	_ = v3679
	var v3687 int32
	_ = v3687
	var v3688 int32
	_ = v3688
	var v3689 int64
	_ = v3689
	var v3697 int64
	_ = v3697
	var v3703 int32
	_ = v3703
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3714 int32
	_ = v3714
	var v3719 int32
	_ = v3719
	var v3720 int32
	_ = v3720
	var v3725 int32
	_ = v3725
	var v3726 int32
	_ = v3726
	var v3732 int32
	_ = v3732
	var v3733 int32
	_ = v3733
	var v3734 int32
	_ = v3734
	var v3735 int32
	_ = v3735
	var v3736 int32
	_ = v3736
	var v3745 int32
	_ = v3745
	var v3746 int32
	_ = v3746
	var v3750 int32
	_ = v3750
	var v3753 int32
	_ = v3753
	var v3754 int32
	_ = v3754
	var v3755 int32
	_ = v3755
	var v3763 int32
	_ = v3763
	var v3764 int32
	_ = v3764
	var v3765 int64
	_ = v3765
	var v3770 int64
	_ = v3770
	var v3776 int32
	_ = v3776
	var v3782 int32
	_ = v3782
	var v3787 int32
	_ = v3787
	var v3788 int32
	_ = v3788
	var v3789 int32
	_ = v3789
	var v3793 int32
	_ = v3793
	var v3796 int32
	_ = v3796
	var v3797 int32
	_ = v3797
	var v3798 int32
	_ = v3798
	var v3806 int32
	_ = v3806
	var v3807 int32
	_ = v3807
	var v3810 int32
	_ = v3810
	var v3811 int32
	_ = v3811
	var v3816 int32
	_ = v3816
	var v3822 int32
	_ = v3822
	var v3823 int32
	_ = v3823
	var v3824 int32
	_ = v3824
	var v3826 int32
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3828 int64
	_ = v3828
	var v3833 int64
	_ = v3833
	var v3839 int32
	_ = v3839
	var v3845 int32
	_ = v3845
	var v3850 int32
	_ = v3850
	var v3857 int32
	_ = v3857
	var v3860 int32
	_ = v3860
	var v3861 int32
	_ = v3861
	var v3869 int32
	_ = v3869
	var v3874 int32
	_ = v3874
	var v3878 int32
	_ = v3878
	var v3881 int32
	_ = v3881
	var v3882 int32
	_ = v3882
	var v3890 int32
	_ = v3890
	var v3895 int32
	_ = v3895
	v4 = l3
	v33 = m.G0
	v35 = v33 + int32(-64)
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
	v3878 = m.ExcPending
	if v3878 != 0 {
		goto L5
	} else {
		goto L843
	}
L2:
	;
	v63 = v37
	goto L4
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v41
	v45 = F_smgropen(m, v33+int32(-16), v38)
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
	F__bt_metaversion(m, l0, v33+int32(-1), v33+int32(-2))
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
	v3857 = m.ExcPending
	if v3857 != 0 {
		goto L5
	} else {
		goto L839
	}
L15:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+62)))
	if v73 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+63)))
	if v76&int32(1) == int32(0) {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v73 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+63)))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+3)))
	v192 = m.G0
	v194 = v192 - int32(1168)
	m.G0 = v194
	v198 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L5
	} else {
		goto L36
	}
L21:
	;
	v84 = F__bt_allequalimage(m, l0, int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	if v84 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v86)+10)))
	if v87 <= int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v99 = int32(0)
	goto L25
L25:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v90+v99<<(uint(int32(2))%32))))
	if v127 != int32(1982) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L31
	}
L27:
	;
	v131 = v99 + int32(1)
	if v87 != v131 {
		v99 = v131
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L20
L31:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v140 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_0), v35)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	F_errhint(m, int32(_a_F_bt_index_check_callback_1), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(345), int32(_a_F_bt_index_check_callback_3))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	if v198 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+1104)) = v200 + int32(4)
	if v4 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v219 = F_palloc0(m, int32(72))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L5
	} else {
		goto L48
	}
L40:
	;
	v206 = int32(_a_F_bt_index_check_callback_4)
	goto L42
L41:
	;
	v206 = int32(_a_F_bt_index_check_callback_5)
	goto L42
L42:
	;
	F_errmsg_internal(m, v206, v194+int32(1104))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	if v4 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v214 = int32(393)
	goto L46
L45:
	;
	v214 = int32(390)
	goto L46
L46:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), v214, int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	goto L39
L48:
	;
	v221 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v219)+28)) = v221
	*(*uint8)(unsafe.Add(mBase, uint32(v219)+12)) = uint8(v191)
	*(*uint8)(unsafe.Add(mBase, uint32(v219)+11)) = uint8(v190)
	*(*uint8)(unsafe.Add(mBase, uint32(v219)+10)) = uint8(v189)
	*(*uint8)(unsafe.Add(mBase, uint32(v219)+9)) = uint8(v4)
	*(*uint8)(unsafe.Add(mBase, uint32(v219)+8)) = uint8(v188)
	*(*int32)(unsafe.Add(mBase, uint32(v219)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = l0
	if v189 == v221 {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	m.G0 = v411 - int32(-64)
	return
L50:
	;
	v3788 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v3789 = *(*int32)(unsafe.Add(mBase, uint32(v910)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3793 = m.ExcPending
	if v3793 != 0 {
		goto L5
	} else {
		goto L827
	}
L51:
	;
	v3709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+7)))
	if v3709&int32(32) == int32(0) {
		v3725 = v1075
		goto L814
	} else {
		goto L815
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3674 = m.ExcPending
	if v3674 != 0 {
		goto L5
	} else {
		goto L808
	}
L53:
	;
	v3176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+10)))
	if v3176 == int32(1) {
		goto L743
	} else {
		goto L744
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3158 = m.ExcPending
	if v3158 != 0 {
		goto L5
	} else {
		goto L739
	}
L55:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+12)))
	if v318 != int32(1) {
		goto L74
	} else {
		goto L75
	}
L56:
	;
	v233 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+48))
	v237 = *(*float32)(unsafe.Add(mBase, uint32(v236)+100))
	v238 = int32(_a_F_bt_index_check_callback_7)
	v239 = int32(_a_F_bt_index_check_callback_8)
	v240 = *(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[0]))
	v242 = *(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[1]))
	v243 = v240 ^ v242
	*(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[1])) = base.I64_rotl(v243, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[0])) = v243<<(uint(int64(16))%64) ^ base.I64_rotl(v240, int64(24)) ^ v243
	v257 = base.I64_extend_i32_u(v233) * int64(452)
	v258 = base.I64_trunc_sat_f32_s(v237)
	if v258 < v257 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v260 = v257
	goto L60
L59:
	;
	v260 = v258
	goto L60
L60:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[2]))
	v269 = F_bloom_create(m, v260, v262, base.I64_rotl(v240*int64(5), int64(7))*int64(9))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v219)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v219)+60)) = v269
	v274 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v276 = F_RegisterSnapshot(m, v274)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219)+28)) = v276
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[3]))
	if v280 < int32(2) {
		goto L55
	} else {
		goto L64
	}
L64:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+19)))
	if v284 != int32(1) {
		goto L55
	} else {
		goto L65
	}
L65:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+16))
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v288)+20)))
	v290 = int32(768)
	if v289&v290 != v290 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v296 = v294
	goto L68
L67:
	;
	v296 = int32(2)
	goto L68
L68:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v297))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v296)) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v309 == int32(0) {
		goto L54
	} else {
		goto L73
	}
L70:
	;
	v309 = base.B2i32(base.Ui32(v296) < base.Ui32(v297))
	goto L69
L71:
	;
	goto L72
L72:
	;
	v309 = int32(base.Ui32(v296-v297) >> (uint(int32(31)) % 32))
	goto L69
L73:
	;
	goto L55
L74:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+11)))
	if v335 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L75:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v322 = F_BuildIndexInfo(m, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219)+24)) = v322
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322)+116)))
	if v325 != int32(1) {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v219)+28))
	if v328 != 0 {
		goto L74
	} else {
		goto L78
	}
L78:
	;
	v329 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	v331 = F_RegisterSnapshot(m, v329)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219)+28)) = v331
	goto L74
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3133 = m.ExcPending
	if v3133 != 0 {
		goto L5
	} else {
		goto L734
	}
L82:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+8)))
	if v338 == int32(0) {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[4]))
	v347 = F_AllocSetContextCreateInternal(m, v342, int32(_a_F_bt_index_check_callback_9), int32(0), int32(_a_F_bt_index_check_callback_10), int32(_a_F_bt_index_check_callback_11))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L5
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219)+16)) = v347
	v351 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219)+20)) = v351
	v355 = F_palloc_btree_page(m, v219, int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L5
	} else {
		goto L89
	}
L88:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v355)+32))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v355)+36))
	v398 = v393
	v401 = v194
	v402 = v219
	v404 = int32(-1)
	v411 = v35
	v414 = l0
	v416 = v394
	v419 = int32(1)
	v423 = l1
	goto L98
L89:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v355)+40))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v355)+32))
	if v357 == v358 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v362 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	if v362 == int32(0) {
		goto L88
	} else {
		goto L92
	}
L92:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L93
	}
L93:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+1056)) = v369 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_12), v194+int32(1056))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v355)+40))
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v355)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v194)+1048)) = v379
	*(*int64)(unsafe.Add(mBase, uint32(v194)+1040)) = v378
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_13), v194+int32(1040))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(512), int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	goto L88
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3110 = m.ExcPending
	if v3110 != 0 {
		goto L5
	} else {
		goto L730
	}
L98:
	;
	if v398 == int32(0) {
		goto L53
	} else {
		goto L100
	}
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3088 = m.ExcPending
	if v3088 != 0 {
		goto L5
	} else {
		goto L726
	}
L100:
	;
	v431 = int32(_a_F_bt_index_check_callback_14)
	v432 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[4]))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v402)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[4])) = v434
	v438 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	if v438 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v401)+1024)) = v416
	if v416 != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	v457 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v402)+56)) = uint8(v457)
	v460 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v402)+52)) = v460
	__phi464 = v457
	__phi465 = v398
	__phi468 = v401
	__phi469 = v402
	__phi471 = v460
	__phi478 = v411
	__phi481 = v414
	__phi483 = v416
	__phi485 = v460
	__phi486 = v419
	__phi489 = v432
	__phi490 = v423
	__phi491 = v404
	v464 = __phi464
	v465 = __phi465
	v468 = __phi468
	v469 = __phi469
	v471 = __phi471
	v478 = __phi478
	v481 = __phi481
	v483 = __phi483
	v485 = __phi485
	v486 = __phi486
	v489 = __phi489
	v490 = __phi490
	v491 = __phi491
	goto L113
L105:
	;
	v444 = int32(_a_F_bt_index_check_callback_15)
	goto L107
L106:
	;
	v444 = int32(_a_F_bt_index_check_callback_16)
	goto L107
L107:
	;
	if v419 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v445 = int32(_a_F_bt_index_check_callback_17)
	goto L110
L109:
	;
	v445 = v444
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v401)+1028)) = v445
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_18), v401+int32(1024))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(645), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	goto L104
L113:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[5]))
	if v497 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v3075 = *(*int32)(unsafe.Add(mBase, uint32(v469)+48))
	if v3075 != 0 {
		goto L721
	} else {
		goto L722
	}
L115:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L5
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+36)) = v465
	v501 = F_palloc_btree_page(m, v469, v465)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L5
	} else {
		goto L119
	}
L118:
	;
	goto L117
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+32)) = v501
	v504 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v501)+4)))
	v505 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v501))))
	*(*int64)(unsafe.Add(mBase, uint32(v469)+40)) = v504 | v505<<(uint(int64(32))%64)
	v510 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v501)+16)))
	v511 = v501 + v510
	v512 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v511)+12)))
	if v512&int32(20) != 0 {
		goto L126
	} else {
		goto L127
	}
L120:
	;
	if v465 == v464 {
		goto L97
	} else {
		goto L707
	}
L121:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v511)+8))
	if v859 == v483 {
		goto L214
	} else {
		goto L215
	}
L122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L5
	} else {
		goto L209
	}
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L5
	} else {
		goto L205
	}
L124:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L5
	} else {
		goto L201
	}
L125:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L5
	} else {
		goto L196
	}
L126:
	;
	if v512&int32(4) != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	if v485 != int32(-1) {
		v590 = v471
		v591 = v485
		goto L139
	} else {
		goto L140
	}
L129:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+9)))
	if v517&int32(1) != 0 {
		goto L125
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	if v520 == int32(0) {
		goto L124
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	v525 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L5
	} else {
		goto L134
	}
L134:
	;
	if v525 == int32(0) {
		v3007 = v471
		v3021 = v485
		goto L120
	} else {
		goto L135
	}
L135:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+1008)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v468)+1012)) = v533 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_20), v468+int32(1008))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L5
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(692), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L5
	} else {
		goto L138
	}
L138:
	;
	v3007 = v471
	v3021 = v485
	goto L120
L139:
	;
	if v464 == int32(0) {
		goto L121
	} else {
		goto L154
	}
L140:
	;
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+9)))
	if v550 == int32(1) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v553 = F_bt_leftmost_ignoring_half_dead(m, v469, v465, v511)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L5
	} else {
		goto L144
	}
L142:
	;
	v563 = v512
	goto L143
L143:
	;
	if v563&int32(1) != 0 {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	if v553 == int32(0) {
		goto L123
	} else {
		goto L145
	}
L145:
	;
	v557 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v511)+12)))
	if base.B2i32(v557&int32(130) == int32(0))&v486 != 0 {
		goto L122
	} else {
		goto L146
	}
L146:
	;
	v563 = v557
	goto L143
L147:
	;
	v590 = int32(-1)
	v591 = int32(0)
	goto L139
L148:
	;
	goto L149
L149:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	if v572 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v573 = int32(2)
	goto L152
L151:
	;
	v573 = int32(1)
	goto L152
L152:
	;
	v574 = F_PageGetItemIdCareful_2(m, v469, v568, v569, v573)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L5
	} else {
		goto L153
	}
L153:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	v580 = v576 + v577&int32(_a_F_bt_index_check_callback_21)
	v581 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v580))))
	v584 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v580)+2)))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v511)+8))
	v590 = v586 - int32(1)
	v591 = v581<<(uint(int32(16))%32) | v584
	goto L139
L154:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v511)))
	if v594 == v464 {
		goto L121
	} else {
		goto L155
	}
L155:
	;
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+9)))
	if v596 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v600 = int32(0)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v469)+20))
	v603 = F_ReadBufferExtended(m, v599, v600, v464, v600, v602)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L5
	} else {
		goto L159
	}
L157:
	;
	v722 = v594
	goto L158
L158:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L5
	} else {
		goto L191
	}
L159:
	;
	F_LockBuffer(m, v603, int32(1))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L5
	} else {
		goto L160
	}
L160:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	F__bt_checkpage(m, v608, v603)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L5
	} else {
		goto L161
	}
L161:
	;
	if v603 < int32(0) {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v629 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v628)+16)))
	v630 = v629 + v628
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630)+12)))
	if v631&int32(4) != 0 {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	v614 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[6]))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v614+(v603^int32(-1))<<(uint(int32(2))%32))))
	v628 = v620
	goto L162
L164:
	;
	goto L165
L165:
	;
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[7]))
	v628 = v622 + v603<<(uint(int32(13))%32) + int32(-8192)
	goto L162
L166:
	;
	F_UnlockReleaseBuffer(m, v603)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L5
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	if v637 == v464 {
		v682 = int32(-1)
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L121
L170:
	;
	F_UnlockReleaseBuffer(m, v603)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L5
	} else {
		goto L181
	}
L171:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v640 = int32(0)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v469)+20))
	v643 = F_ReadBufferExtended(m, v639, v640, v637, v640, v642)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L5
	} else {
		goto L172
	}
L172:
	;
	F_LockBuffer(m, v643, int32(1))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L5
	} else {
		goto L173
	}
L173:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	F__bt_checkpage(m, v648, v643)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L5
	} else {
		goto L174
	}
L174:
	;
	if v643 < int32(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	F_UnlockReleaseBuffer(m, v643)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L5
	} else {
		goto L180
	}
L176:
	;
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[6]))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v654+(v643^int32(-1))<<(uint(int32(2))%32))))
	v661 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v660)+16)))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v660+v661)))
	v678 = v663
	goto L175
L177:
	;
	goto L178
L178:
	;
	v665 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[7]))
	v668 = v665 + v643<<(uint(int32(13))%32)
	v671 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v668-int32(_a_F_bt_index_check_callback_22)))))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v668+v671)+uint32(_c_F_bt_index_check_callback[8])))
	if v643 == int32(0) {
		v682 = v675
		goto L170
	} else {
		goto L179
	}
L179:
	;
	v678 = v675
	goto L175
L180:
	;
	v682 = v678
	goto L170
L181:
	;
	if v682 == v464 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v688 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L5
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+36)) = v637
	v722 = v682
	goto L158
L185:
	;
	if v688 == int32(0) {
		goto L121
	} else {
		goto L186
	}
L186:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L5
	} else {
		goto L187
	}
L187:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+928)) = v696 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_23), v468+int32(928))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+920)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v468)+916)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v468)+912)) = v464
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_24), v468+int32(912))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L5
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1177), int32(_a_F_bt_index_check_callback_25))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L5
	} else {
		goto L190
	}
L190:
	;
	goto L121
L191:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+80)) = v732 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_26), v468+int32(80))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L5
	} else {
		goto L193
	}
L193:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+72)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v468)+68)) = v464
	*(*int32)(unsafe.Add(mBase, uint32(v468)+64)) = v741
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_27), v468-int32(-64))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1197), int32(_a_F_bt_index_check_callback_25))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L5
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L5
	} else {
		goto L197
	}
L197:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v762)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+976)) = v763 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_28), v468+int32(976))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v511)))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+968)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v468)+964)) = v464
	*(*int32)(unsafe.Add(mBase, uint32(v468)+960)) = v465
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_27), v468+int32(960))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L5
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(681), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L5
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L5
	} else {
		goto L202
	}
L202:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v793)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+992)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v468)+996)) = v794 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_29), v468+int32(992))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L5
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(687), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L5
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L5
	} else {
		goto L206
	}
L206:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+944)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v468)+948)) = v817 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_30), v468+int32(944))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L5
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(710), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L5
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L5
	} else {
		goto L210
	}
L210:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v839)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+48)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v468)+52)) = v840 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_31), v468+int32(48))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L5
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(716), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L5
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	v2988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2964)+12)))
	if v2988&int32(1) != 0 {
		v3007 = v590
		v3021 = v591
		goto L120
	} else {
		goto L703
	}
L214:
	;
	v862 = *(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v468)+1128)) = v862
	v865 = *(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[10]))
	*(*int64)(unsafe.Add(mBase, uint32(v468)+1120)) = v865
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v868 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v867)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v868) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	goto L216
L216:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2928 = m.ExcPending
	if v2928 != 0 {
		goto L5
	} else {
		goto L698
	}
L217:
	;
	v876 = int32(base.Ui32(v868+int32(_a_F_bt_index_check_callback_32)) >> (uint(int32(2)) % 32))
	goto L219
L218:
	;
	v876 = int32(0)
	goto L219
L219:
	;
	v877 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v867)+16)))
	v878 = v867 + v877
	v881 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L5
	} else {
		goto L220
	}
L220:
	;
	if v881 != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v883 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v878)+12)))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+872)) = v884
	*(*int32)(unsafe.Add(mBase, uint32(v468)+864)) = v876 & int32(_a_F_bt_index_check_callback_33)
	if v883&int32(1) != 0 {
		goto L224
	} else {
		goto L225
	}
L222:
	;
	goto L223
L223:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v878)+4))
	if v906 != 0 {
		goto L229
	} else {
		goto L230
	}
L224:
	;
	v893 = int32(_a_F_bt_index_check_callback_34)
	goto L226
L225:
	;
	v893 = int32(_a_F_bt_index_check_callback_35)
	goto L226
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v468)+868)) = v893
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_36), v468+int32(864))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L5
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1252), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L5
	} else {
		goto L228
	}
L228:
	;
	goto L223
L229:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v910 = F_PageGetItemIdCareful_2(m, v469, v907, v908, int32(1))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L5
	} else {
		goto L232
	}
L230:
	;
	v1023 = int32(1)
	goto L231
L231:
	;
	v1025 = v876 & int32(_a_F_bt_index_check_callback_33)
	if base.Ui32(v1025) < base.Ui32(v1023) {
		v2964 = v878
		goto L213
	} else {
		goto L271
	}
L232:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+8)))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v915 = int32(1)
	v923 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v914)+16)))
	v924 = v914 + v923
	v925 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v924)+12)))
	if v925&int32(20) != 0 {
		v1004 = v915
		goto L234
	} else {
		goto L235
	}
L233:
	;
	if v1014 == int32(0) {
		goto L50
	} else {
		goto L267
	}
L234:
	;
	v1014 = v1004
	goto L233
L235:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v912)+192))
	v929 = int32(*(*int16)(unsafe.Add(mBase, uint32(v928)+10)))
	v930 = int32(*(*int16)(unsafe.Add(mBase, uint32(v928)+8)))
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v914+int32(4))+20))
	v937 = v914 + v934&int32(_a_F_bt_index_check_callback_21)
	v938 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v937)+6)))
	v940 = v938 & int32(_a_F_bt_index_check_callback_10)
	if v940 == int32(0) {
		v958 = v930
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v924)+4))
	if v961 != 0 {
		goto L242
	} else {
		goto L243
	}
L237:
	;
	v943 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v937)+4)))
	if v943&int32(_a_F_bt_index_check_callback_10) != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v946 = int32(0)
	if base.B2i32(v913 == v946)|v943&int32(_a_F_bt_index_check_callback_38)|base.B2i32(v929 != v930) != 0 {
		v1004 = v946
		goto L234
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v958 = v943 & int32(4095)
	goto L236
L241:
	;
	v958 = v930
	goto L236
L242:
	;
	v962 = int32(2)
	goto L244
L243:
	;
	v962 = int32(1)
	goto L244
L244:
	;
	if v925&int32(1) != 0 {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	v988 = int32(0)
	if v940 == v988 {
		v1004 = v988
		goto L234
	} else {
		goto L261
	}
L246:
	;
	if base.Ui32(v962) <= base.Ui32(v915) {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	goto L248
L248:
	;
	if v915 == v962 {
		goto L256
	} else {
		goto L257
	}
L249:
	;
	if v940 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L250:
	;
	goto L251
L251:
	;
	if v913 != 0 {
		goto L245
	} else {
		goto L255
	}
L252:
	;
	v1014 = base.B2i32(v958 == v930)
	goto L233
L253:
	;
	goto L254
L254:
	;
	v969 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v937)+4)))
	v1014 = int32(base.Ui32(v969)>>(uint(int32(13))%32)) & base.B2i32(v958 == v930)
	goto L233
L255:
	;
	v1014 = base.B2i32(v958 == v929)
	goto L233
L256:
	;
	v977 = base.B2i32(v958 == int32(0))
	if v913|v977 != 0 {
		v1004 = v977 | (v913 ^ int32(1))
		goto L234
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	if v913 != 0 {
		goto L245
	} else {
		goto L260
	}
L259:
	;
	v984 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v937)+4)))
	v1014 = base.B2i32(v984 == int32(1))
	goto L233
L260:
	;
	v1014 = base.B2i32(v958 == v929)
	goto L233
L261:
	;
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937)+5)))
	if v991&int32(32) != 0 {
		v1004 = v988
		goto L234
	} else {
		goto L262
	}
L262:
	;
	v994 = F_BTreeTupleGetHeapTID(m, v937)
	mBase = m.M
	if v958 != v929 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v997 = v994
	goto L265
L264:
	;
	v997 = int32(0)
	goto L265
L265:
	;
	if v997 != 0 {
		v1004 = v988
		goto L234
	} else {
		goto L266
	}
L266:
	;
	v1004 = base.B2i32(v958 <= v929) & base.B2i32(int32(0) < v958)
	goto L234
L267:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v878)+4))
	if v1019 != 0 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v1020 = int32(2)
	goto L270
L269:
	;
	v1020 = int32(1)
	goto L270
L270:
	;
	v1023 = v1020
	goto L231
L271:
	;
	v1035 = v878
	v1047 = v1023
	goto L274
L272:
	;
	F_pfree(m, v2563)
	mBase = m.M
	v2924 = m.ExcPending
	if v2924 != 0 {
		goto L5
	} else {
		goto L697
	}
L273:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2891 = m.ExcPending
	if v2891 != 0 {
		goto L5
	} else {
		goto L692
	}
L274:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[5]))
	if v1060 != 0 {
		goto L276
	} else {
		goto L277
	}
L275:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2857 = m.ExcPending
	if v2857 != 0 {
		goto L5
	} else {
		goto L687
	}
L276:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L5
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v1066 = v1047 & int32(_a_F_bt_index_check_callback_33)
	v1067 = F_PageGetItemIdCareful_2(m, v469, v1063, v1064, v1066)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L5
	} else {
		goto L288
	}
L279:
	;
	goto L278
L280:
	;
	goto L275
L281:
	;
	v2850 = v1047 + int32(1)
	if base.Ui32(v2850&int32(_a_F_bt_index_check_callback_33)) <= base.Ui32(v1025) {
		v1035 = v2825
		v1047 = v2850
		goto L274
	} else {
		goto L686
	}
L282:
	;
	v2618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2594)+12)))
	if v2618&int32(1) != 0 {
		v2825 = v2594
		goto L281
	} else {
		goto L625
	}
L283:
	;
	v2507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+12)))
	if v2507 != int32(1) {
		v2594 = v1035
		goto L282
	} else {
		goto L596
	}
L284:
	;
	v2408 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2409 = *(*int32)(unsafe.Add(mBase, uint32(v2404)))
	v2413 = F__bt_mkscankey(m, v2408, v1958+v2409&int32(_a_F_bt_index_check_callback_21))
	mBase = m.M
	v2414 = m.ExcPending
	if v2414 != 0 {
		goto L5
	} else {
		goto L583
	}
L285:
	;
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+484)) = v1066
	*(*int32)(unsafe.Add(mBase, uint32(v468)+480)) = v2273
	v2279 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v468+int32(480))
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L5
	} else {
		goto L562
	}
L286:
	;
	v2193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+7)))
	if v2193&int32(32) == int32(0) {
		v2209 = v1075
		goto L549
	} else {
		goto L550
	}
L287:
	;
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+612)) = v1066
	*(*int32)(unsafe.Add(mBase, uint32(v468)+608)) = v2151
	v2157 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v468+int32(608))
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L5
	} else {
		goto L542
	}
L288:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1067)))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v1075 = v1072 + v1069&int32(_a_F_bt_index_check_callback_21)
	v1076 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+6)))
	v1078 = v1076 & int32(_a_F_bt_index_check_callback_40)
	if int32(base.Ui32(v1069)>>(uint(int32(17))%32)) == v1078 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v1081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+8)))
	v1089 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1072)+16)))
	v1090 = v1072 + v1089
	v1091 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1090)+12)))
	if v1091&int32(20) != 0 {
		v1170 = int32(1)
		goto L293
	} else {
		goto L294
	}
L290:
	;
	goto L291
L291:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L5
	} else {
		goto L536
	}
L292:
	;
	if v1180 == int32(0) {
		goto L326
	} else {
		goto L327
	}
L293:
	;
	v1180 = v1170
	goto L292
L294:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1080)+192))
	v1095 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1094)+10)))
	v1096 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1094)+8)))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1072+v1066<<(uint(int32(2))%32))+20))
	v1103 = v1072 + v1100&int32(_a_F_bt_index_check_callback_21)
	v1104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1103)+6)))
	v1106 = v1104 & int32(_a_F_bt_index_check_callback_10)
	if v1106 == int32(0) {
		v1124 = v1096
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+4))
	if v1127 != 0 {
		goto L301
	} else {
		goto L302
	}
L296:
	;
	v1109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1103)+4)))
	if v1109&int32(_a_F_bt_index_check_callback_10) != 0 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1112 = int32(0)
	if base.B2i32(v1081 == v1112)|v1109&int32(_a_F_bt_index_check_callback_38)|base.B2i32(v1095 != v1096) != 0 {
		v1170 = v1112
		goto L293
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	v1124 = v1109 & int32(4095)
	goto L295
L300:
	;
	v1124 = v1096
	goto L295
L301:
	;
	v1128 = int32(2)
	goto L303
L302:
	;
	v1128 = int32(1)
	goto L303
L303:
	;
	if v1091&int32(1) != 0 {
		goto L305
	} else {
		goto L306
	}
L304:
	;
	v1154 = int32(0)
	if v1106 == v1154 {
		v1170 = v1154
		goto L293
	} else {
		goto L320
	}
L305:
	;
	if base.Ui32(v1128) <= base.Ui32(v1066) {
		goto L308
	} else {
		goto L309
	}
L306:
	;
	goto L307
L307:
	;
	if v1066 == v1128 {
		goto L315
	} else {
		goto L316
	}
L308:
	;
	if v1106 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L309:
	;
	goto L310
L310:
	;
	if v1081 != 0 {
		goto L304
	} else {
		goto L314
	}
L311:
	;
	v1180 = base.B2i32(v1124 == v1096)
	goto L292
L312:
	;
	goto L313
L313:
	;
	v1135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1103)+4)))
	v1180 = int32(base.Ui32(v1135)>>(uint(int32(13))%32)) & base.B2i32(v1124 == v1096)
	goto L292
L314:
	;
	v1180 = base.B2i32(v1124 == v1095)
	goto L292
L315:
	;
	v1143 = base.B2i32(v1124 == int32(0))
	if v1081|v1143 != 0 {
		v1170 = v1143 | (v1081 ^ int32(1))
		goto L293
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	if v1081 != 0 {
		goto L304
	} else {
		goto L319
	}
L318:
	;
	v1150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1103)+4)))
	v1180 = base.B2i32(v1150 == int32(1))
	goto L292
L319:
	;
	v1180 = base.B2i32(v1124 == v1095)
	goto L292
L320:
	;
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1103)+5)))
	if v1157&int32(32) != 0 {
		v1170 = v1154
		goto L293
	} else {
		goto L321
	}
L321:
	;
	v1160 = F_BTreeTupleGetHeapTID(m, v1103)
	mBase = m.M
	if v1124 != v1095 {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v1163 = v1160
	goto L324
L323:
	;
	v1163 = int32(0)
	goto L324
L324:
	;
	if v1163 != 0 {
		v1170 = v1154
		goto L293
	} else {
		goto L325
	}
L325:
	;
	v1170 = base.B2i32(v1124 <= v1095) & base.B2i32(int32(0) < v1124)
	goto L293
L326:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+756)) = v1066
	*(*int32)(unsafe.Add(mBase, uint32(v468)+752)) = v1183
	v1189 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v468+int32(752))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L5
	} else {
		goto L329
	}
L327:
	;
	goto L328
L328:
	;
	v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035)+12)))
	if v1279&int32(1) == int32(0) {
		goto L352
	} else {
		goto L353
	}
L329:
	;
	v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+7)))
	if v1191&int32(32) == int32(0) {
		v1207 = v1075
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1207)+2)))
	v1209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1207))))
	v1210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1207)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+740)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v468)+736)) = v1208 | v1209<<(uint(int32(16))%32)
	v1219 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v468+int32(736))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L5
	} else {
		goto L333
	}
L331:
	;
	v1196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+5)))
	if v1196&int32(32) == int32(0) {
		v1207 = v1075
		goto L330
	} else {
		goto L332
	}
L332:
	;
	v1201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+2)))
	v1202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075))))
	v1207 = v1201 + (v1075 + v1202<<(uint(int32(16))%32))
	goto L330
L333:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L5
	} else {
		goto L334
	}
L334:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L5
	} else {
		goto L335
	}
L335:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1228)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+720)) = v1229 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_41), v468+int32(720))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L5
	} else {
		goto L336
	}
L336:
	;
	v1238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+7)))
	if v1238&int32(32) == int32(0) {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	v1254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1035)+12)))
	v1255 = *(*int64)(unsafe.Add(mBase, uint32(v469)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+708)) = uint32(v1255)
	v1258 = int64(base.Ui64(v1255) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+704)) = uint32(v1258)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+700)) = v1219
	*(*int32)(unsafe.Add(mBase, uint32(v468)+692)) = v1253
	*(*int32)(unsafe.Add(mBase, uint32(v468)+688)) = v1189
	if v1254&int32(1) != 0 {
		goto L341
	} else {
		goto L342
	}
L338:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1249)+192))
	v1251 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1250)+8)))
	v1253 = v1251
	goto L337
L339:
	;
	v1243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+4)))
	if v1243&int32(_a_F_bt_index_check_callback_10) != 0 {
		goto L338
	} else {
		goto L340
	}
L340:
	;
	v1253 = v1243 & int32(4095)
	goto L337
L341:
	;
	v1267 = int32(_a_F_bt_index_check_callback_42)
	goto L343
L342:
	;
	v1267 = int32(_a_F_bt_index_check_callback_43)
	goto L343
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v468)+696)) = v1267
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_44), v468+int32(688))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L5
	} else {
		goto L344
	}
L344:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1352), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L5
	} else {
		goto L345
	}
L345:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L346:
	;
	if base.Ui32(v1635) < base.Ui32(v1078) {
		goto L286
	} else {
		goto L415
	}
L347:
	;
	v1616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035)+12)))
	if v1616&int32(1) != 0 {
		v1635 = int32(2704)
		goto L346
	} else {
		goto L409
	}
L348:
	;
	v1543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+7)))
	if v1543&int32(32) == int32(0) {
		v1559 = v1075
		goto L399
	} else {
		goto L400
	}
L349:
	;
	F__bt_freestack(m, v1307)
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L5
	} else {
		goto L396
	}
L350:
	;
	F__bt_relbuf(m, v1333)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L5
	} else {
		goto L395
	}
L351:
	;
	v1383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+7)))
	if v1383&int32(32) == int32(0) {
		goto L380
	} else {
		goto L381
	}
L352:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+4))
	if v1286 != 0 {
		goto L355
	} else {
		goto L356
	}
L353:
	;
	goto L354
L354:
	;
	v1296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+11)))
	if v1296 != int32(1) {
		goto L351
	} else {
		goto L361
	}
L355:
	;
	v1287 = int32(2)
	goto L357
L356:
	;
	v1287 = int32(1)
	goto L357
L357:
	;
	if v1287 != v1066 {
		goto L351
	} else {
		goto L358
	}
L358:
	;
	v1289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+9)))
	if v1289 != int32(1) {
		v2825 = v1035
		goto L281
	} else {
		goto L359
	}
L359:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+8))
	F_bt_child_highkey_check(m, v469, v1066, int32(0), v1293)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L5
	} else {
		goto L360
	}
L360:
	;
	v2825 = v1035
	goto L281
L361:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v1300 = F__bt_mkscankey(m, v1299, v1075)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L5
	} else {
		goto L362
	}
L362:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v1307 = F__bt_search(m, v1302, int32(0), v1300, v468+int32(1164), int32(1))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L5
	} else {
		goto L363
	}
L363:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v468)+1164))
	if v1309 == int32(0) {
		goto L349
	} else {
		goto L364
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v468)+1136)) = v1075
	v1313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+6)))
	v1314 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+1160)) = v1314
	*(*int32)(unsafe.Add(mBase, uint32(v468)+1144)) = v1300
	*(*uint8)(unsafe.Add(mBase, uint32(v468)+1152)) = uint8(v1314)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+1148)) = v1309
	*(*int32)(unsafe.Add(mBase, uint32(v468)+1140)) = (v1313&int32(_a_F_bt_index_check_callback_40) + int32(7)) & int32(_a_F_bt_index_check_callback_45)
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v1330 = F__bt_binsrch_insert(m, v1327, v468+int32(1136))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L5
	} else {
		goto L365
	}
L365:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v468)+1164))
	if v1333 < int32(0) {
		goto L367
	} else {
		goto L368
	}
L366:
	;
	v1352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1351)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1352) {
		goto L370
	} else {
		goto L371
	}
L367:
	;
	v1337 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[6]))
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1337+(v1333^int32(-1))<<(uint(int32(2))%32))))
	v1351 = v1343
	goto L366
L368:
	;
	goto L369
L369:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[7]))
	v1351 = v1345 + v1333<<(uint(int32(13))%32) + int32(-8192)
	goto L366
L370:
	;
	v1360 = int32(base.Ui32(v1352+int32(_a_F_bt_index_check_callback_32)) >> (uint(int32(2)) % 32))
	goto L372
L371:
	;
	v1360 = int32(0)
	goto L372
L372:
	;
	if base.Ui32(v1360&int32(_a_F_bt_index_check_callback_33)) < base.Ui32(v1330) {
		goto L350
	} else {
		goto L373
	}
L373:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v468)+1160))
	if int32(0) < v1364 {
		goto L350
	} else {
		goto L374
	}
L374:
	;
	v1367 = F__bt_compare(m, v1332, v1300, v1351, v1330)
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L5
	} else {
		goto L375
	}
L375:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v468)+1164))
	F__bt_relbuf(m, v1370)
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L5
	} else {
		goto L376
	}
L376:
	;
	F__bt_freestack(m, v1307)
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L5
	} else {
		goto L377
	}
L377:
	;
	F_pfree(m, v1300)
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L5
	} else {
		goto L378
	}
L378:
	;
	if v1367 != 0 {
		goto L348
	} else {
		goto L379
	}
L379:
	;
	goto L351
L380:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v1521 = F__bt_mkscankey(m, v1520, v1075)
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L5
	} else {
		goto L393
	}
L381:
	;
	v1388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+5)))
	if v1388&int32(32) == int32(0) {
		goto L380
	} else {
		goto L382
	}
L382:
	;
	v1393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+2)))
	v1394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075))))
	v1398 = v1393 + (v1075 + v1394<<(uint(int32(16))%32))
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1398)))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+1136)) = v1399
	v1401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1398)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v468)+1140)) = uint16(v1401)
	v1404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+4)))
	if v1404&int32(4094) == int32(0) {
		goto L380
	} else {
		goto L383
	}
L383:
	;
	v1412 = int32(1)
	goto L384
L384:
	;
	v1441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+2)))
	v1442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075))))
	v1443 = int32(16)
	v1449 = v1441 + (v1075 + v1442<<(uint(v1443)%32)) + v1412*int32(6)
	v1451 = v468 + int32(1136)
	v1455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1449)+2)))
	v1456 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1449))))
	v1459 = v1455 | v1456<<(uint(v1443)%32)
	v1460 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1451)+2)))
	v1461 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1451))))
	v1464 = v1460 | v1461<<(uint(v1443)%32)
	if base.Ui32(v1459) < base.Ui32(v1464) {
		v1475 = int32(-1)
		goto L387
	} else {
		goto L388
	}
L385:
	;
	goto L380
L386:
	;
	if v1475 <= int32(0) {
		goto L287
	} else {
		goto L391
	}
L387:
	;
	goto L386
L388:
	;
	if base.Ui32(v1464) < base.Ui32(v1459) {
		v1475 = int32(1)
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v1469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1449)+4)))
	v1470 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1451)+4)))
	if base.Ui32(v1469) < base.Ui32(v1470) {
		v1475 = int32(-1)
		goto L387
	} else {
		goto L390
	}
L390:
	;
	v1475 = base.B2i32(base.Ui32(v1470) < base.Ui32(v1469))
	goto L387
L391:
	;
	v1478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1449)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v468)+1140)) = uint16(v1478)
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1449)))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+1136)) = v1480
	v1483 = v1412 + int32(1)
	v1484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+4)))
	if base.Ui32(v1483) < base.Ui32(v1484&int32(4095)) {
		v1412 = v1483
		goto L384
	} else {
		goto L392
	}
L392:
	;
	goto L385
L393:
	;
	v1523 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1521)+4)) = uint8(v1523)
	v1525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1521))))
	if v1525 == v1523 {
		goto L347
	} else {
		goto L394
	}
L394:
	;
	v1635 = int32(2712)
	goto L346
L395:
	;
	goto L349
L396:
	;
	F_pfree(m, v1300)
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L5
	} else {
		goto L397
	}
L397:
	;
	goto L348
L398:
	;
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+676)) = v1066
	*(*int32)(unsafe.Add(mBase, uint32(v468)+672)) = v1560
	v1566 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v468+int32(672))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L5
	} else {
		goto L402
	}
L399:
	;
	goto L398
L400:
	;
	v1548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+5)))
	if v1548&int32(32) == int32(0) {
		v1559 = v1075
		goto L399
	} else {
		goto L401
	}
L401:
	;
	v1553 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+2)))
	v1554 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075))))
	v1559 = v1553 + (v1075 + v1554<<(uint(int32(16))%32))
	goto L399
L402:
	;
	v1568 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1559)+2)))
	v1569 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1559))))
	v1570 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1559)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+660)) = v1570
	*(*int32)(unsafe.Add(mBase, uint32(v468)+656)) = v1568 | v1569<<(uint(int32(16))%32)
	v1579 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v468+int32(656))
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L5
	} else {
		goto L403
	}
L403:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L5
	} else {
		goto L404
	}
L404:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L5
	} else {
		goto L405
	}
L405:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1588)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+640)) = v1589 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_46), v468+int32(640))
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L5
	} else {
		goto L406
	}
L406:
	;
	v1598 = *(*int64)(unsafe.Add(mBase, uint32(v469)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+636)) = uint32(v1598)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+628)) = v1579
	*(*int32)(unsafe.Add(mBase, uint32(v468)+624)) = v1566
	v1603 = int64(base.Ui64(v1598) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+632)) = uint32(v1603)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_47), v468+int32(624))
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L5
	} else {
		goto L407
	}
L407:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1399), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L5
	} else {
		goto L408
	}
L408:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L409:
	;
	v1620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+7)))
	if v1620&int32(32) == int32(0) {
		v1635 = int32(2712)
		goto L346
	} else {
		goto L410
	}
L410:
	;
	v1626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+4)))
	if v1626&int32(_a_F_bt_index_check_callback_10) != 0 {
		v1635 = int32(2712)
		goto L346
	} else {
		goto L411
	}
L411:
	;
	if v1626&int32(_a_F_bt_index_check_callback_38) != 0 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v1633 = int32(2712)
	goto L414
L413:
	;
	v1633 = int32(2704)
	goto L414
L414:
	;
	v1635 = v1633
	goto L346
L415:
	;
	v1637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+10)))
	if v1637 != int32(1) {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1521)+8))
	v1773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+8)))
	if v1773 != int32(1) {
		goto L439
	} else {
		goto L440
	}
L417:
	;
	v1640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035)+12)))
	if v1640&int32(1) == int32(0) {
		goto L416
	} else {
		goto L418
	}
L418:
	;
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1067)))
	v1646 = int32(_a_F_bt_index_check_callback_48)
	if v1645&v1646 == v1646 {
		goto L416
	} else {
		goto L419
	}
L419:
	;
	v1650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+7)))
	if v1650&int32(32) == int32(0) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v1729 = F_bt_normalize_tuple(m, v469, v1075)
	mBase = m.M
	v1730 = m.ExcPending
	if v1730 != 0 {
		goto L5
	} else {
		goto L435
	}
L421:
	;
	v1655 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+4)))
	if v1655&int32(_a_F_bt_index_check_callback_10) == int32(0) {
		goto L420
	} else {
		goto L422
	}
L422:
	;
	if v1655&int32(4095) == int32(0) {
		goto L416
	} else {
		goto L423
	}
L423:
	;
	v1667 = int32(0)
	goto L424
L424:
	;
	v1697 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+2)))
	v1698 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075))))
	v1707 = F__bt_form_posting(m, v1075, v1697+(v1075+v1698<<(uint(int32(16))%32))+v1667*int32(6), int32(1))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L5
	} else {
		goto L426
	}
L425:
	;
	goto L416
L426:
	;
	v1709 = F_bt_normalize_tuple(m, v469, v1707)
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L5
	} else {
		goto L427
	}
L427:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v469)+60))
	v1712 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1709)+6)))
	F_bloom_add_element(m, v1711, v1709, v1712&int32(_a_F_bt_index_check_callback_40))
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L5
	} else {
		goto L428
	}
L428:
	;
	if v1709 != v1707 {
		goto L429
	} else {
		goto L430
	}
L429:
	;
	F_pfree(m, v1709)
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L5
	} else {
		goto L432
	}
L430:
	;
	goto L431
L431:
	;
	F_pfree(m, v1707)
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L5
	} else {
		goto L433
	}
L432:
	;
	goto L431
L433:
	;
	v1723 = v1667 + int32(1)
	v1724 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+4)))
	if base.Ui32(v1723) < base.Ui32(v1724&int32(4095)) {
		v1667 = v1723
		goto L424
	} else {
		goto L434
	}
L434:
	;
	goto L425
L435:
	;
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v469)+60))
	v1732 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1729)+6)))
	F_bloom_add_element(m, v1731, v1729, v1732&int32(_a_F_bt_index_check_callback_40))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L5
	} else {
		goto L436
	}
L436:
	;
	if v1075 == v1729 {
		goto L416
	} else {
		goto L437
	}
L437:
	;
	F_pfree(m, v1729)
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L5
	} else {
		goto L438
	}
L438:
	;
	goto L416
L439:
	;
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+4))
	if v1801 == int32(0) {
		goto L443
	} else {
		goto L444
	}
L440:
	;
	v1776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+7)))
	if v1776&int32(32) == int32(0) {
		goto L439
	} else {
		goto L441
	}
L441:
	;
	v1781 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+4)))
	if v1781&int32(_a_F_bt_index_check_callback_10) == int32(0) {
		goto L439
	} else {
		goto L442
	}
L442:
	;
	v1786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+2)))
	v1787 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075))))
	v1794 = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v1521)+8)) = v1786 + (v1075 + v1787<<(uint(int32(16))%32)) + v1781&int32(4095)*v1794 - v1794
	goto L439
L443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1521)+8)) = v1772
	v1821 = v1047 + int32(1)
	v1823 = v1821 & int32(_a_F_bt_index_check_callback_33)
	v1824 = base.B2i32(base.Ui32(v1025) < base.Ui32(v1823))
	if v1824 == int32(0) {
		goto L452
	} else {
		goto L453
	}
L444:
	;
	v1804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035)+12)))
	if v1804&int32(1) != 0 {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v1810 = F__bt_compare(m, v1807, v1521, v1808, int32(1))
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L5
	} else {
		goto L448
	}
L446:
	;
	goto L447
L447:
	;
	v1815 = F_invariant_l_offset(m, v469, v1521, int32(1))
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L5
	} else {
		goto L450
	}
L448:
	;
	if v1810 <= int32(0) {
		goto L443
	} else {
		goto L449
	}
L449:
	;
	goto L51
L450:
	;
	if v1815 == int32(0) {
		goto L51
	} else {
		goto L451
	}
L451:
	;
	goto L443
L452:
	;
	v1827 = F_invariant_l_offset(m, v469, v1521, v1823)
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L5
	} else {
		goto L455
	}
L453:
	;
	goto L454
L454:
	;
	v1831 = int32(0)
	v1832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+12)))
	if v1832 != int32(1) {
		v1912 = v1831
		goto L457
	} else {
		goto L458
	}
L455:
	;
	if v1827 == int32(0) {
		goto L285
	} else {
		goto L456
	}
L456:
	;
	goto L454
L457:
	;
	if v1066 != v1025 {
		v2594 = v1035
		goto L282
	} else {
		goto L483
	}
L458:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v469)+24))
	v1836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1835)+116)))
	if v1836 != int32(1) {
		v1912 = v1831
		goto L457
	} else {
		goto L459
	}
L459:
	;
	v1839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035)+12)))
	if v1839&int32(1) == int32(0) {
		v1912 = v1831
		goto L457
	} else {
		goto L460
	}
L460:
	;
	v1844 = int32(0)
	v1846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1521)+2)))
	if v1846 != 0 {
		v1882 = v1844
		v1883 = v1844
		goto L461
	} else {
		goto L462
	}
L461:
	;
	if v1882|v1824 != 0 {
		v1912 = v1883
		goto L457
	} else {
		goto L473
	}
L462:
	;
	v1847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+7)))
	if v1847&int32(32) != 0 {
		goto L464
	} else {
		goto L465
	}
L463:
	;
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	F_bt_entry_unique_check(m, v469, v1075, v1862, v1066, v468+int32(1120))
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L5
	} else {
		goto L470
	}
L464:
	;
	v1850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+5)))
	if v1850&int32(32) != 0 {
		goto L463
	} else {
		goto L467
	}
L465:
	;
	goto L466
L466:
	;
	v1853 = int32(0)
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v468)+1132))
	if v1854 == v1853 {
		v1882 = v1844
		v1883 = v1853
		goto L461
	} else {
		goto L468
	}
L467:
	;
	goto L466
L468:
	;
	v1857 = int32(0)
	v1858 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1854)+4)))
	if v1858 == v1857 {
		v1882 = v1844
		v1883 = v1857
		goto L461
	} else {
		goto L469
	}
L469:
	;
	goto L463
L470:
	;
	v1867 = int32(1)
	v1868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+12)))
	if v1868 != v1867 {
		v1912 = v1867
		goto L457
	} else {
		goto L471
	}
L471:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v469)+24))
	v1872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1871)+116)))
	if v1872 != int32(1) {
		v1912 = v1867
		goto L457
	} else {
		goto L472
	}
L472:
	;
	v1875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035)+12)))
	v1876 = int32(1)
	v1882 = base.B2i32(v1875&v1876 == int32(0))
	v1883 = v1876
	goto L461
L473:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1521)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1521)+8)) = int32(0)
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v1890 = F__bt_compare(m, v1888, v1521, v1889, v1823)
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L5
	} else {
		goto L476
	}
L474:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1521)+8)) = v1885
	v1912 = v1910
	goto L457
L475:
	;
	if v1883 != 0 {
		v1910 = int32(1)
		goto L474
	} else {
		goto L481
	}
L476:
	;
	if v1890 == int32(0) {
		goto L477
	} else {
		goto L478
	}
L477:
	;
	v1894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1521)+2)))
	if v1894 != int32(1) {
		goto L475
	} else {
		goto L480
	}
L478:
	;
	goto L479
L479:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v468)+1128)) = int64(4294967295)
	v1899 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v468)+1124)) = uint16(v1899)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+1120)) = int32(-1)
	v1910 = v1883
	goto L474
L480:
	;
	goto L479
L481:
	;
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	F_bt_entry_unique_check(m, v469, v1075, v1904, v1066, v468+int32(1120))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L5
	} else {
		goto L482
	}
L482:
	;
	v1910 = int32(0)
	goto L474
L483:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v1916 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1915)+16)))
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v1915+v1916)+4))
	if v1918 == int32(0) {
		goto L486
	} else {
		goto L487
	}
L484:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v2080 = m.ExcPending
	if v2080 != 0 {
		goto L5
	} else {
		goto L530
	}
L485:
	;
	v2076 = int32(0)
	v2478 = v2076
	v2481 = v2050
	v2484 = v2076
	goto L283
L486:
	;
	v2050 = int32(0)
	goto L485
L487:
	;
	goto L488
L488:
	;
	v1925 = v1918
	goto L489
L489:
	;
	v1955 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[5]))
	if v1955 != 0 {
		goto L491
	} else {
		goto L492
	}
L490:
	;
	v2004 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1958)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v2004) {
		goto L508
	} else {
		goto L509
	}
L491:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L5
	} else {
		goto L494
	}
L492:
	;
	goto L493
L493:
	;
	v1958 = F_palloc_btree_page(m, v469, v1925)
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L5
	} else {
		goto L496
	}
L494:
	;
	goto L493
L495:
	;
	goto L490
L496:
	;
	v1960 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1958)+16)))
	v1961 = v1958 + v1960
	v1962 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1961)+12)))
	if v1962&int32(20) == int32(0) {
		goto L495
	} else {
		goto L497
	}
L497:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v1961)+4))
	if v1967 == int32(0) {
		goto L495
	} else {
		goto L498
	}
L498:
	;
	v1972 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L5
	} else {
		goto L499
	}
L499:
	;
	if v1972 != 0 {
		goto L500
	} else {
		goto L501
	}
L500:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L5
	} else {
		goto L503
	}
L501:
	;
	goto L502
L502:
	;
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v1961)+4))
	F_pfree(m, v1958)
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L5
	} else {
		goto L507
	}
L503:
	;
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v1977)+48))
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v1961)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+372)) = v1925
	*(*int32)(unsafe.Add(mBase, uint32(v468)+368)) = v1979
	*(*int32)(unsafe.Add(mBase, uint32(v468)+376)) = v1978 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_49), v468+int32(368))
	mBase = m.M
	v1989 = m.ExcPending
	if v1989 != 0 {
		goto L5
	} else {
		goto L504
	}
L504:
	;
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_50), int32(0))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L5
	} else {
		goto L505
	}
L505:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1926), int32(_a_F_bt_index_check_callback_51))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L5
	} else {
		goto L506
	}
L506:
	;
	goto L502
L507:
	;
	v1925 = v2001
	goto L489
L508:
	;
	v2012 = int32(base.Ui32(v2004+int32(_a_F_bt_index_check_callback_32)) >> (uint(int32(2)) % 32))
	goto L510
L509:
	;
	v2012 = int32(0)
	goto L510
L510:
	;
	if v1962&int32(1) != 0 {
		goto L512
	} else {
		goto L513
	}
L511:
	;
	v2039 = int32(0)
	v2042 = F_errstart(m, int32(13), v2039)
	mBase = m.M
	v2043 = m.ExcPending
	if v2043 != 0 {
		goto L5
	} else {
		goto L528
	}
L512:
	;
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v1961)+4))
	if v2017 != 0 {
		goto L515
	} else {
		goto L516
	}
L513:
	;
	goto L514
L514:
	;
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v1961)+4))
	if v2030 != 0 {
		goto L523
	} else {
		goto L524
	}
L515:
	;
	v2018 = int32(2)
	goto L517
L516:
	;
	v2018 = int32(1)
	goto L517
L517:
	;
	if base.Ui32(v2012&int32(_a_F_bt_index_check_callback_33)) < base.Ui32(v2018) {
		goto L511
	} else {
		goto L518
	}
L518:
	;
	v2022 = F_PageGetItemIdCareful_2(m, v469, v1925, v1958, v2018)
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L5
	} else {
		goto L519
	}
L519:
	;
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v1961)+4))
	if v2026 != 0 {
		goto L520
	} else {
		goto L521
	}
L520:
	;
	v2027 = int32(2)
	goto L522
L521:
	;
	v2027 = int32(1)
	goto L522
L522:
	;
	v2404 = v2022
	v2405 = v2027
	goto L284
L523:
	;
	v2031 = int32(3)
	goto L525
L524:
	;
	v2031 = int32(2)
	goto L525
L525:
	;
	if base.Ui32(v2012&int32(_a_F_bt_index_check_callback_33)) < base.Ui32(v2031) {
		goto L511
	} else {
		goto L526
	}
L526:
	;
	v2036 = F_PageGetItemIdCareful_2(m, v469, v1925, v1958, v2031)
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L5
	} else {
		goto L527
	}
L527:
	;
	v2404 = v2036
	v2405 = int32(0)
	goto L284
L528:
	;
	if v2042 != 0 {
		goto L484
	} else {
		goto L529
	}
L529:
	;
	v2050 = v2039
	goto L485
L530:
	;
	v2081 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1961)+12)))
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(v2082)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+324)) = v1925
	*(*int32)(unsafe.Add(mBase, uint32(v468)+328)) = v2083 + int32(4)
	v2088 = int32(0)
	if v2081&int32(1) != 0 {
		goto L531
	} else {
		goto L532
	}
L531:
	;
	v2093 = int32(_a_F_bt_index_check_callback_34)
	goto L533
L532:
	;
	v2093 = int32(_a_F_bt_index_check_callback_35)
	goto L533
L533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v468)+320)) = v2088 + v2093
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_52), v468+int32(320))
	mBase = m.M
	v2100 = m.ExcPending
	if v2100 != 0 {
		goto L5
	} else {
		goto L534
	}
L534:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(2057), int32(_a_F_bt_index_check_callback_51))
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L5
	} else {
		goto L535
	}
L535:
	;
	v2478 = v2088
	v2481 = v2039
	v2484 = int32(0)
	goto L283
L536:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L5
	} else {
		goto L537
	}
L537:
	;
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v2114)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+800)) = v2115 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_53), v468+int32(800))
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L5
	} else {
		goto L538
	}
L538:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v1067)))
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v2126 = *(*int64)(unsafe.Add(mBase, uint32(v469)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+788)) = uint32(v2126)
	v2129 = int64(base.Ui64(v2126) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+784)) = uint32(v2129)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+776)) = v1078
	*(*int32)(unsafe.Add(mBase, uint32(v468)+772)) = v1066
	*(*int32)(unsafe.Add(mBase, uint32(v468)+768)) = v2125
	*(*int32)(unsafe.Add(mBase, uint32(v468)+780)) = int32(base.Ui32(v2124) >> (uint(int32(17)) % 32))
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_54), v468+int32(768))
	mBase = m.M
	v2141 = m.ExcPending
	if v2141 != 0 {
		goto L5
	} else {
		goto L539
	}
L539:
	;
	F_errhint(m, int32(_a_F_bt_index_check_callback_55), int32(0))
	mBase = m.M
	v2145 = m.ExcPending
	if v2145 != 0 {
		goto L5
	} else {
		goto L540
	}
L540:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1327), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		goto L5
	} else {
		goto L541
	}
L541:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L542:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L5
	} else {
		goto L543
	}
L543:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2165 = m.ExcPending
	if v2165 != 0 {
		goto L5
	} else {
		goto L544
	}
L544:
	;
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+592)) = v2167 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_56), v468+int32(592))
	mBase = m.M
	v2175 = m.ExcPending
	if v2175 != 0 {
		goto L5
	} else {
		goto L545
	}
L545:
	;
	v2176 = *(*int64)(unsafe.Add(mBase, uint32(v469)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+588)) = uint32(v2176)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+580)) = v1412
	*(*int32)(unsafe.Add(mBase, uint32(v468)+576)) = v2157
	v2181 = int64(base.Ui64(v2176) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+584)) = uint32(v2181)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_57), v468+int32(576))
	mBase = m.M
	v2187 = m.ExcPending
	if v2187 != 0 {
		goto L5
	} else {
		goto L546
	}
L546:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1428), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
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
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+196)) = v1066
	*(*int32)(unsafe.Add(mBase, uint32(v468)+192)) = v2210
	v2216 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v468+int32(192))
	mBase = m.M
	v2217 = m.ExcPending
	if v2217 != 0 {
		goto L5
	} else {
		goto L552
	}
L549:
	;
	goto L548
L550:
	;
	v2198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+5)))
	if v2198&int32(32) == int32(0) {
		v2209 = v1075
		goto L549
	} else {
		goto L551
	}
L551:
	;
	v2203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+2)))
	v2204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075))))
	v2209 = v2203 + (v1075 + v2204<<(uint(int32(16))%32))
	goto L549
L552:
	;
	v2218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2209)+2)))
	v2219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2209))))
	v2220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2209)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+180)) = v2220
	*(*int32)(unsafe.Add(mBase, uint32(v468)+176)) = v2218 | v2219<<(uint(int32(16))%32)
	v2229 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v468+int32(176))
	mBase = m.M
	v2230 = m.ExcPending
	if v2230 != 0 {
		goto L5
	} else {
		goto L553
	}
L553:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2234 = m.ExcPending
	if v2234 != 0 {
		goto L5
	} else {
		goto L554
	}
L554:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L5
	} else {
		goto L555
	}
L555:
	;
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+160)) = v1078
	*(*int32)(unsafe.Add(mBase, uint32(v468)+164)) = v2239 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_58), v468+int32(160))
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L5
	} else {
		goto L556
	}
L556:
	;
	v2249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1035)+12)))
	v2250 = *(*int64)(unsafe.Add(mBase, uint32(v469)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+144)) = uint32(v2250)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+136)) = v2229
	*(*int32)(unsafe.Add(mBase, uint32(v468)+128)) = v2216
	v2255 = int64(base.Ui64(v2250) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+140)) = uint32(v2255)
	if v2249&int32(1) != 0 {
		goto L557
	} else {
		goto L558
	}
L557:
	;
	v2261 = int32(_a_F_bt_index_check_callback_42)
	goto L559
L558:
	;
	v2261 = int32(_a_F_bt_index_check_callback_43)
	goto L559
L559:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v468)+132)) = v2261
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_59), v468+int32(128))
	mBase = m.M
	v2267 = m.ExcPending
	if v2267 != 0 {
		goto L5
	} else {
		goto L560
	}
L560:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1483), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L5
	} else {
		goto L561
	}
L561:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L562:
	;
	v2281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+7)))
	if v2281&int32(32) == int32(0) {
		v2297 = v1075
		goto L564
	} else {
		goto L565
	}
L563:
	;
	v2298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2297)+2)))
	v2299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2297))))
	v2300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2297)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+468)) = v2300
	*(*int32)(unsafe.Add(mBase, uint32(v468)+464)) = v2298 | v2299<<(uint(int32(16))%32)
	v2309 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v468+int32(464))
	mBase = m.M
	v2310 = m.ExcPending
	if v2310 != 0 {
		goto L5
	} else {
		goto L567
	}
L564:
	;
	goto L563
L565:
	;
	v2286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+5)))
	if v2286&int32(32) == int32(0) {
		v2297 = v1075
		goto L564
	} else {
		goto L566
	}
L566:
	;
	v2291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+2)))
	v2292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075))))
	v2297 = v2291 + (v1075 + v2292<<(uint(int32(16))%32))
	goto L564
L567:
	;
	v2311 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v2313 = v1821 & int32(_a_F_bt_index_check_callback_33)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+452)) = v2313
	*(*int32)(unsafe.Add(mBase, uint32(v468)+448)) = v2311
	v2319 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v468+int32(448))
	mBase = m.M
	v2320 = m.ExcPending
	if v2320 != 0 {
		goto L5
	} else {
		goto L568
	}
L568:
	;
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v2323 = F_PageGetItemIdCareful_2(m, v469, v2321, v2322, v2313)
	mBase = m.M
	v2324 = m.ExcPending
	if v2324 != 0 {
		goto L5
	} else {
		goto L569
	}
L569:
	;
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(v2323)))
	v2329 = v2325 + v2326&int32(_a_F_bt_index_check_callback_21)
	v2330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2329)+7)))
	if v2330&int32(32) == int32(0) {
		v2346 = v2329
		goto L571
	} else {
		goto L572
	}
L570:
	;
	v2347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2346)+2)))
	v2348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2346))))
	v2349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2346)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+436)) = v2349
	*(*int32)(unsafe.Add(mBase, uint32(v468)+432)) = v2347 | v2348<<(uint(int32(16))%32)
	v2358 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v468+int32(432))
	mBase = m.M
	v2359 = m.ExcPending
	if v2359 != 0 {
		goto L5
	} else {
		goto L574
	}
L571:
	;
	goto L570
L572:
	;
	v2335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2329)+5)))
	if v2335&int32(32) == int32(0) {
		v2346 = v2329
		goto L571
	} else {
		goto L573
	}
L573:
	;
	v2340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2329)+2)))
	v2341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2329))))
	v2346 = v2340 + (v2329 + v2341<<(uint(int32(16))%32))
	goto L571
L574:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2363 = m.ExcPending
	if v2363 != 0 {
		goto L5
	} else {
		goto L575
	}
L575:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2366 = m.ExcPending
	if v2366 != 0 {
		goto L5
	} else {
		goto L576
	}
L576:
	;
	v2367 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2368 = *(*int32)(unsafe.Add(mBase, uint32(v2367)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+416)) = v2368 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_60), v468+int32(416))
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L5
	} else {
		goto L577
	}
L577:
	;
	v2377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1035)+12)))
	v2378 = *(*int64)(unsafe.Add(mBase, uint32(v469)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+412)) = uint32(v2378)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+404)) = v2358
	v2382 = int64(base.Ui64(v2378) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+408)) = uint32(v2382)
	if v2377&int32(1) != 0 {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v2388 = int32(_a_F_bt_index_check_callback_42)
	goto L580
L579:
	;
	v2388 = int32(_a_F_bt_index_check_callback_43)
	goto L580
L580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v468)+400)) = v2388
	*(*int32)(unsafe.Add(mBase, uint32(v468)+396)) = v2319
	*(*int32)(unsafe.Add(mBase, uint32(v468)+392)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v468)+384)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v468)+388)) = v2388
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_61), v468+int32(384))
	mBase = m.M
	v2398 = m.ExcPending
	if v2398 != 0 {
		goto L5
	} else {
		goto L581
	}
L581:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1641), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2403 = m.ExcPending
	if v2403 != 0 {
		goto L5
	} else {
		goto L582
	}
L582:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L583:
	;
	v2415 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2413)+4)) = uint8(v2415)
	v2417 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v2419 = F__bt_compare(m, v2417, v2413, v2418, v1025)
	mBase = m.M
	v2420 = m.ExcPending
	if v2420 != 0 {
		goto L5
	} else {
		goto L584
	}
L584:
	;
	v2422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2413))))
	if int32(0)-(v2422^int32(1)) < v2419 {
		v2478 = v2413
		v2481 = int32(1)
		v2484 = v2405
		goto L283
	} else {
		goto L585
	}
L585:
	;
	v2427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+9)))
	if v2427 == int32(0) {
		goto L586
	} else {
		goto L587
	}
L586:
	;
	v2430 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v2431 = F_palloc_btree_page(m, v469, v2430)
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L5
	} else {
		goto L589
	}
L587:
	;
	goto L588
L588:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L5
	} else {
		goto L591
	}
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+32)) = v2431
	v2434 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2431)+16)))
	v2436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2431+v2434)+12)))
	if v2436&int32(20) != 0 {
		v3007 = v590
		v3021 = v591
		goto L120
	} else {
		goto L590
	}
L590:
	;
	goto L588
L591:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L5
	} else {
		goto L592
	}
L592:
	;
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2448 = *(*int32)(unsafe.Add(mBase, uint32(v2447)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+352)) = v2448 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_62), v468+int32(352))
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		goto L5
	} else {
		goto L593
	}
L593:
	;
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v2458 = *(*int64)(unsafe.Add(mBase, uint32(v469)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+348)) = uint32(v2458)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+340)) = v1066
	*(*int32)(unsafe.Add(mBase, uint32(v468)+336)) = v2457
	v2463 = int64(base.Ui64(v2458) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+344)) = uint32(v2463)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_63), v468+int32(336))
	mBase = m.M
	v2469 = m.ExcPending
	if v2469 != 0 {
		goto L5
	} else {
		goto L594
	}
L594:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1753), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2474 = m.ExcPending
	if v2474 != 0 {
		goto L5
	} else {
		goto L595
	}
L595:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L596:
	;
	v2510 = *(*int32)(unsafe.Add(mBase, uint32(v469)+24))
	v2511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2510)+116)))
	if v2481&v2511 != int32(1) {
		v2594 = v1035
		goto L282
	} else {
		goto L597
	}
L597:
	;
	v2515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035)+12)))
	if v2515&int32(1) == int32(0) {
		v2594 = v1035
		goto L282
	} else {
		goto L598
	}
L598:
	;
	v2520 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+4))
	if v2520 == int32(0) {
		v2594 = v1035
		goto L282
	} else {
		goto L599
	}
L599:
	;
	v2525 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v2526 = m.ExcPending
	if v2526 != 0 {
		goto L5
	} else {
		goto L600
	}
L600:
	;
	if v2525 != 0 {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_64), int32(0))
	mBase = m.M
	v2530 = m.ExcPending
	if v2530 != 0 {
		goto L5
	} else {
		goto L604
	}
L602:
	;
	goto L603
L603:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2478)+8)) = int32(0)
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v2540 = F__bt_compare(m, v2538, v2478, v2539, v1025)
	mBase = m.M
	v2541 = m.ExcPending
	if v2541 != 0 {
		goto L5
	} else {
		goto L606
	}
L604:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1765), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2535 = m.ExcPending
	if v2535 != 0 {
		goto L5
	} else {
		goto L605
	}
L605:
	;
	goto L603
L606:
	;
	if v2540 != 0 {
		v2594 = v1035
		goto L282
	} else {
		goto L607
	}
L607:
	;
	v2542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2478)+2)))
	if v2542 != 0 {
		v2594 = v1035
		goto L282
	} else {
		goto L608
	}
L608:
	;
	if v1912 == int32(0) {
		goto L609
	} else {
		goto L610
	}
L609:
	;
	v2545 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	F_bt_entry_unique_check(m, v469, v1075, v2545, v1025, v468+int32(1120))
	mBase = m.M
	v2549 = m.ExcPending
	if v2549 != 0 {
		goto L5
	} else {
		goto L612
	}
L610:
	;
	goto L611
L611:
	;
	v2552 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L5
	} else {
		goto L613
	}
L612:
	;
	goto L611
L613:
	;
	if v2552 != 0 {
		goto L614
	} else {
		goto L615
	}
L614:
	;
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_65), int32(0))
	mBase = m.M
	v2557 = m.ExcPending
	if v2557 != 0 {
		goto L5
	} else {
		goto L617
	}
L615:
	;
	goto L616
L616:
	;
	v2563 = F_palloc_btree_page(m, v469, v2520)
	mBase = m.M
	v2564 = m.ExcPending
	if v2564 != 0 {
		goto L5
	} else {
		goto L619
	}
L617:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1788), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2562 = m.ExcPending
	if v2562 != 0 {
		goto L5
	} else {
		goto L618
	}
L618:
	;
	goto L616
L619:
	;
	v2565 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2563)+16)))
	v2566 = v2563 + v2565
	v2567 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2566)+12)))
	if v2567&int32(20) != 0 {
		goto L272
	} else {
		goto L620
	}
L620:
	;
	if v2567&int32(1) == int32(0) {
		goto L280
	} else {
		goto L621
	}
L621:
	;
	v2574 = F_PageGetItemIdCareful_2(m, v469, v2520, v2563, v2484)
	mBase = m.M
	v2575 = m.ExcPending
	if v2575 != 0 {
		goto L5
	} else {
		goto L622
	}
L622:
	;
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(v2574)))
	F_bt_entry_unique_check(m, v469, v2563+v2576&int32(_a_F_bt_index_check_callback_21), v2520, v2484, v468+int32(1120))
	mBase = m.M
	v2583 = m.ExcPending
	if v2583 != 0 {
		goto L5
	} else {
		goto L623
	}
L623:
	;
	F_pfree(m, v2563)
	mBase = m.M
	v2585 = m.ExcPending
	if v2585 != 0 {
		goto L5
	} else {
		goto L624
	}
L624:
	;
	v2594 = v2566
	goto L282
L625:
	;
	v2621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+9)))
	if v2621 != int32(1) {
		v2825 = v2594
		goto L281
	} else {
		goto L626
	}
L626:
	;
	v2624 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v2625 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v2626 = F_PageGetItemIdCareful_2(m, v469, v2624, v2625, v1066)
	mBase = m.M
	v2627 = m.ExcPending
	if v2627 != 0 {
		goto L5
	} else {
		goto L627
	}
L627:
	;
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v2629 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2628)+16)))
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v2626)))
	v2633 = v2628 + v2630&int32(_a_F_bt_index_check_callback_21)
	v2634 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2633))))
	v2637 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2633)+2)))
	v2638 = v2634<<(uint(int32(16))%32) | v2637
	v2639 = F_palloc_btree_page(m, v469, v2638)
	mBase = m.M
	v2640 = m.ExcPending
	if v2640 != 0 {
		goto L5
	} else {
		goto L628
	}
L628:
	;
	v2641 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2639)+16)))
	v2642 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2639)+12)))
	v2644 = *(*int32)(unsafe.Add(mBase, uint32(v2628+v2629)+8))
	F_bt_child_highkey_check(m, v469, v1066, v2639, v2644)
	mBase = m.M
	v2646 = m.ExcPending
	if v2646 != 0 {
		goto L5
	} else {
		goto L629
	}
L629:
	;
	v2647 = v2639 + v2641
	v2648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+12)))
	if v2648&int32(4) != 0 {
		goto L273
	} else {
		goto L630
	}
L630:
	;
	v2653 = *(*int32)(unsafe.Add(mBase, uint32(v2647)+4))
	if v2653 != 0 {
		goto L631
	} else {
		goto L632
	}
L631:
	;
	v2654 = int32(2)
	goto L633
L632:
	;
	v2654 = int32(1)
	goto L633
L633:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v2642) {
		goto L634
	} else {
		goto L635
	}
L634:
	;
	v2662 = int32(base.Ui32(v2642+int32(_a_F_bt_index_check_callback_32)) >> (uint(int32(2)) % 32))
	goto L636
L635:
	;
	v2662 = int32(0)
	goto L636
L636:
	;
	v2664 = v2662 & int32(_a_F_bt_index_check_callback_33)
	if base.Ui32(v2654) <= base.Ui32(v2664) {
		goto L637
	} else {
		goto L638
	}
L637:
	;
	v2667 = v2654
	goto L640
L638:
	;
	goto L639
L639:
	;
	F_pfree(m, v2639)
	mBase = m.M
	v2816 = m.ExcPending
	if v2816 != 0 {
		goto L5
	} else {
		goto L685
	}
L640:
	;
	v2698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+12)))
	if v2698&int32(1) == int32(0) {
		goto L643
	} else {
		goto L644
	}
L641:
	;
	goto L639
L642:
	;
	v2779 = v2667 + int32(1)
	if base.Ui32(v2779&int32(_a_F_bt_index_check_callback_33)) <= base.Ui32(v2664) {
		v2667 = v2779
		goto L640
	} else {
		goto L684
	}
L643:
	;
	v2707 = *(*int32)(unsafe.Add(mBase, uint32(v2647)+4))
	if v2707 != 0 {
		goto L646
	} else {
		goto L647
	}
L644:
	;
	goto L645
L645:
	;
	v2711 = v2667 & int32(_a_F_bt_index_check_callback_33)
	v2712 = F_PageGetItemIdCareful_2(m, v469, v2638, v2639, v2711)
	mBase = m.M
	v2713 = m.ExcPending
	if v2713 != 0 {
		goto L5
	} else {
		goto L650
	}
L646:
	;
	v2708 = int32(2)
	goto L648
L647:
	;
	v2708 = int32(1)
	goto L648
L648:
	;
	if v2667&int32(_a_F_bt_index_check_callback_33) == v2708 {
		goto L642
	} else {
		goto L649
	}
L649:
	;
	goto L645
L650:
	;
	v2714 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2715 = F__bt_compare(m, v2714, v1521, v2639, v2711)
	mBase = m.M
	v2716 = m.ExcPending
	if v2716 != 0 {
		goto L5
	} else {
		goto L651
	}
L651:
	;
	v2717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1521))))
	if v2717 == int32(0) {
		goto L652
	} else {
		goto L653
	}
L652:
	;
	if int32(0) < v2715 {
		goto L52
	} else {
		goto L655
	}
L653:
	;
	goto L654
L654:
	;
	if v2715 == int32(0) {
		goto L657
	} else {
		goto L658
	}
L655:
	;
	goto L642
L656:
	;
	if v2759 <= v2764 {
		goto L52
	} else {
		goto L683
	}
L657:
	;
	v2724 = *(*int32)(unsafe.Add(mBase, uint32(v2712)))
	v2727 = v2639 + v2724&int32(_a_F_bt_index_check_callback_21)
	v2729 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2639)+16)))
	v2730 = v2639 + v2729
	v2731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2730)+12)))
	if v2731&int32(1) != 0 {
		goto L660
	} else {
		goto L661
	}
L658:
	;
	goto L659
L659:
	;
	if v2715 < int32(0) {
		goto L642
	} else {
		goto L682
	}
L660:
	;
	v2736 = *(*int32)(unsafe.Add(mBase, uint32(v2730)+4))
	if v2736 != 0 {
		goto L663
	} else {
		goto L664
	}
L661:
	;
	v2739 = int32(0)
	goto L662
L662:
	;
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2741 = *(*int32)(unsafe.Add(mBase, uint32(v2740)+192))
	v2742 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2741)+10)))
	v2743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2727)+7)))
	if v2743&int32(32) != 0 {
		goto L668
	} else {
		goto L669
	}
L663:
	;
	v2737 = int32(2)
	goto L665
L664:
	;
	v2737 = int32(1)
	goto L665
L665:
	;
	v2739 = base.B2i32(base.Ui32(v2737) <= base.Ui32(v2711))
	goto L662
L666:
	;
	v2762 = F_BTreeTupleGetHeapTIDCareful(m, v469, v2727, v2739)
	mBase = m.M
	v2763 = m.ExcPending
	if v2763 != 0 {
		goto L5
	} else {
		goto L679
	}
L667:
	;
	v2759 = v2757
	goto L666
L668:
	;
	v2746 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2727)+4)))
	if v2746&int32(_a_F_bt_index_check_callback_10) != 0 {
		goto L671
	} else {
		goto L672
	}
L669:
	;
	goto L670
L670:
	;
	v2755 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2741)+8)))
	if v2742 < v2755 {
		v2759 = v2742
		goto L666
	} else {
		goto L678
	}
L671:
	;
	v2749 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2741)+8)))
	if v2749 <= v2742 {
		v2757 = v2749
		goto L667
	} else {
		goto L674
	}
L672:
	;
	goto L673
L673:
	;
	v2752 = v2746 & int32(4095)
	if v2742 < v2752 {
		goto L675
	} else {
		goto L676
	}
L674:
	;
	v2759 = v2742
	goto L666
L675:
	;
	v2754 = v2742
	goto L677
L676:
	;
	v2754 = v2752
	goto L677
L677:
	;
	v2759 = v2754
	goto L666
L678:
	;
	v2757 = v2755
	goto L667
L679:
	;
	v2764 = *(*int32)(unsafe.Add(mBase, uint32(v1521)+12))
	if v2764 != v2759 {
		goto L656
	} else {
		goto L680
	}
L680:
	;
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(v1521)+8))
	if v2766|base.B2i32(v2762 == int32(0)) != 0 {
		goto L52
	} else {
		goto L681
	}
L681:
	;
	goto L642
L682:
	;
	goto L52
L683:
	;
	goto L642
L684:
	;
	goto L641
L685:
	;
	v2825 = v2594
	goto L281
L686:
	;
	v2964 = v2825
	goto L213
L687:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2860 = m.ExcPending
	if v2860 != 0 {
		goto L5
	} else {
		goto L688
	}
L688:
	;
	v2861 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2862 = *(*int32)(unsafe.Add(mBase, uint32(v2861)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+304)) = v2862 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_66), v468+int32(304))
	mBase = m.M
	v2870 = m.ExcPending
	if v2870 != 0 {
		goto L5
	} else {
		goto L689
	}
L689:
	;
	v2871 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v2872 = *(*int64)(unsafe.Add(mBase, uint32(v469)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+296)) = uint32(v2872)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+288)) = v2871
	v2876 = int64(base.Ui64(v2872) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+292)) = uint32(v2876)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_67), v468+int32(288))
	mBase = m.M
	v2882 = m.ExcPending
	if v2882 != 0 {
		goto L5
	} else {
		goto L690
	}
L690:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1806), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2887 = m.ExcPending
	if v2887 != 0 {
		goto L5
	} else {
		goto L691
	}
L691:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L692:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2894 = m.ExcPending
	if v2894 != 0 {
		goto L5
	} else {
		goto L693
	}
L693:
	;
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(v2895)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+272)) = v2896 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_68), v468+int32(272))
	mBase = m.M
	v2904 = m.ExcPending
	if v2904 != 0 {
		goto L5
	} else {
		goto L694
	}
L694:
	;
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v2906 = *(*int64)(unsafe.Add(mBase, uint32(v469)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+268)) = uint32(v2906)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+260)) = v2638
	*(*int32)(unsafe.Add(mBase, uint32(v468)+256)) = v2905
	v2911 = int64(base.Ui64(v2906) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+264)) = uint32(v2911)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_69), v468+int32(256))
	mBase = m.M
	v2917 = m.ExcPending
	if v2917 != 0 {
		goto L5
	} else {
		goto L695
	}
L695:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(2498), int32(_a_F_bt_index_check_callback_70))
	mBase = m.M
	v2922 = m.ExcPending
	if v2922 != 0 {
		goto L5
	} else {
		goto L696
	}
L696:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L697:
	;
	v2964 = v2566
	goto L213
L698:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2931 = m.ExcPending
	if v2931 != 0 {
		goto L5
	} else {
		goto L699
	}
L699:
	;
	v2932 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v2933 = *(*int32)(unsafe.Add(mBase, uint32(v2932)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+896)) = v2933 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_71), v468+int32(896))
	mBase = m.M
	v2941 = m.ExcPending
	if v2941 != 0 {
		goto L5
	} else {
		goto L700
	}
L700:
	;
	v2942 = *(*int32)(unsafe.Add(mBase, uint32(v511)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+888)) = v2942
	*(*int32)(unsafe.Add(mBase, uint32(v468)+884)) = v483
	*(*int32)(unsafe.Add(mBase, uint32(v468)+880)) = v465
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_72), v468+int32(880))
	mBase = m.M
	v2950 = m.ExcPending
	if v2950 != 0 {
		goto L5
	} else {
		goto L701
	}
L701:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(779), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v2955 = m.ExcPending
	if v2955 != 0 {
		goto L5
	} else {
		goto L702
	}
L702:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L703:
	;
	v2991 = *(*int32)(unsafe.Add(mBase, uint32(v2964)+4))
	if v2991 != 0 {
		v3007 = v590
		v3021 = v591
		goto L120
	} else {
		goto L704
	}
L704:
	;
	v2992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+9)))
	if v2992 != int32(1) {
		v3007 = v590
		v3021 = v591
		goto L120
	} else {
		goto L705
	}
L705:
	;
	v2995 = int32(0)
	v2997 = *(*int32)(unsafe.Add(mBase, uint32(v2964)+8))
	F_bt_child_highkey_check(m, v469, v2995, v2995, v2997)
	mBase = m.M
	v2999 = m.ExcPending
	if v2999 != 0 {
		goto L5
	} else {
		goto L706
	}
L706:
	;
	v3007 = v590
	v3021 = v591
	goto L120
L707:
	;
	v3033 = *(*int32)(unsafe.Add(mBase, uint32(v511)))
	if v465 == v3033 {
		goto L97
	} else {
		goto L708
	}
L708:
	;
	v3035 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	v3036 = *(*int32)(unsafe.Add(mBase, uint32(v469)+48))
	if v3036 != 0 {
		goto L709
	} else {
		goto L710
	}
L709:
	;
	F_pfree(m, v3036)
	mBase = m.M
	v3038 = m.ExcPending
	if v3038 != 0 {
		goto L5
	} else {
		goto L712
	}
L710:
	;
	goto L711
L711:
	;
	v3041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+9)))
	if v3041 != int32(1) {
		goto L713
	} else {
		goto L714
	}
L712:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+48)) = int32(0)
	goto L711
L713:
	;
	v3072 = *(*int32)(unsafe.Add(mBase, uint32(v469)+16))
	F_MemoryContextReset(m, v3072)
	mBase = m.M
	v3074 = m.ExcPending
	if v3074 != 0 {
		goto L5
	} else {
		goto L719
	}
L714:
	;
	v3044 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	if v3044 == int32(0) {
		goto L713
	} else {
		goto L715
	}
L715:
	;
	v3047 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v3048 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v3050 = F_PageGetItemIdCareful_2(m, v469, v3047, v3048, int32(1))
	mBase = m.M
	v3051 = m.ExcPending
	if v3051 != 0 {
		goto L5
	} else {
		goto L716
	}
L716:
	;
	v3052 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	v3053 = *(*int32)(unsafe.Add(mBase, uint32(v3050)))
	v3056 = v3052 + v3053&int32(_a_F_bt_index_check_callback_21)
	v3057 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3056)+6)))
	v3060 = F_MemoryContextAlloc(m, v489, v3057&int32(_a_F_bt_index_check_callback_40))
	mBase = m.M
	v3061 = m.ExcPending
	if v3061 != 0 {
		goto L5
	} else {
		goto L717
	}
L717:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+48)) = v3060
	v3063 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3056)+6)))
	v3065 = v3063 & int32(_a_F_bt_index_check_callback_40)
	if v3065 == int32(0) {
		goto L713
	} else {
		goto L718
	}
L718:
	;
	base.MemoryCopy(m, v3060, v3056, v3065)
	goto L713
L719:
	;
	if v3035 != 0 {
		__phi464 = v465
		__phi465 = v3035
		__phi471 = v3007
		__phi485 = v3021
		v464 = __phi464
		v465 = __phi465
		v471 = __phi471
		v485 = __phi485
		goto L113
	} else {
		goto L720
	}
L720:
	;
	goto L114
L721:
	;
	F_pfree(m, v3075)
	mBase = m.M
	v3077 = m.ExcPending
	if v3077 != 0 {
		goto L5
	} else {
		goto L724
	}
L722:
	;
	goto L723
L723:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[4])) = v489
	if v3021 != int32(-1) {
		v398 = v3021
		v401 = v468
		v402 = v469
		v404 = v3007
		v411 = v478
		v414 = v481
		v416 = v3007
		v419 = int32(0)
		v423 = v490
		goto L98
	} else {
		goto L725
	}
L724:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+48)) = int32(0)
	goto L723
L725:
	;
	goto L99
L726:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3091 = m.ExcPending
	if v3091 != 0 {
		goto L5
	} else {
		goto L727
	}
L727:
	;
	v3092 = *(*int32)(unsafe.Add(mBase, uint32(v481)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+116)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v468)+112)) = v3092 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_73), v468+int32(112))
	mBase = m.M
	v3101 = m.ExcPending
	if v3101 != 0 {
		goto L5
	} else {
		goto L728
	}
L728:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(535), int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v3106 = m.ExcPending
	if v3106 != 0 {
		goto L5
	} else {
		goto L729
	}
L729:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L730:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3113 = m.ExcPending
	if v3113 != 0 {
		goto L5
	} else {
		goto L731
	}
L731:
	;
	v3114 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v3115 = *(*int32)(unsafe.Add(mBase, uint32(v3114)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+96)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v468)+100)) = v3115 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_74), v468+int32(96))
	mBase = m.M
	v3124 = m.ExcPending
	if v3124 != 0 {
		goto L5
	} else {
		goto L732
	}
L732:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(791), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v3129 = m.ExcPending
	if v3129 != 0 {
		goto L5
	} else {
		goto L733
	}
L733:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L734:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3136 = m.ExcPending
	if v3136 != 0 {
		goto L5
	} else {
		goto L735
	}
L735:
	;
	v3137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+1072)) = v3137 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_75), v194+int32(1072))
	mBase = m.M
	v3145 = m.ExcPending
	if v3145 != 0 {
		goto L5
	} else {
		goto L736
	}
L736:
	;
	F_errhint(m, int32(_a_F_bt_index_check_callback_76), int32(0))
	mBase = m.M
	v3149 = m.ExcPending
	if v3149 != 0 {
		goto L5
	} else {
		goto L737
	}
L737:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(485), int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v3154 = m.ExcPending
	if v3154 != 0 {
		goto L5
	} else {
		goto L738
	}
L738:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L739:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v3161 = m.ExcPending
	if v3161 != 0 {
		goto L5
	} else {
		goto L740
	}
L740:
	;
	v3162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+1088)) = v3162 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_77), v194+int32(1088))
	mBase = m.M
	v3170 = m.ExcPending
	if v3170 != 0 {
		goto L5
	} else {
		goto L741
	}
L741:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(463), int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v3175 = m.ExcPending
	if v3175 != 0 {
		goto L5
	} else {
		goto L742
	}
L742:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L743:
	;
	v3179 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v3180 = F_BuildIndexInfo(m, v3179)
	mBase = m.M
	v3181 = m.ExcPending
	if v3181 != 0 {
		goto L5
	} else {
		goto L746
	}
L744:
	;
	goto L745
L745:
	;
	v3657 = *(*int32)(unsafe.Add(mBase, uint32(v402)+28))
	if v3657 != 0 {
		goto L803
	} else {
		goto L804
	}
L746:
	;
	v3182 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(v402)+28))
	v3184 = int32(0)
	v3188 = *(*int32)(unsafe.Add(mBase, uint32(v3182)+188))
	v3189 = *(*int32)(unsafe.Add(mBase, uint32(v3188)+8))
	v3190 = m.T0[v3189].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v3182, v3183, v3184, v3184, v3184, int32(449))
	mBase = m.M
	v3191 = m.ExcPending
	if v3191 != 0 {
		goto L5
	} else {
		goto L747
	}
L747:
	;
	v3192 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3180)+116)) = uint8(v3192)
	v3194 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3180)+121)) = uint8(v3194)
	*(*int32)(unsafe.Add(mBase, uint32(v3180)+100)) = v3192
	*(*int64)(unsafe.Add(mBase, uint32(v3180)+92)) = int64(0)
	v3202 = F_errstart(m, int32(14), v3192)
	mBase = m.M
	v3203 = m.ExcPending
	if v3203 != 0 {
		goto L5
	} else {
		goto L748
	}
L748:
	;
	if v3202 != 0 {
		goto L749
	} else {
		goto L750
	}
L749:
	;
	v3204 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v3205 = *(*int32)(unsafe.Add(mBase, uint32(v3204)+48))
	v3206 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	v3207 = *(*int32)(unsafe.Add(mBase, uint32(v3206)+48))
	v3208 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v401)+36)) = v3207 + v3208
	*(*int32)(unsafe.Add(mBase, uint32(v401)+32)) = v3205 + v3208
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_78), v401+int32(32))
	mBase = m.M
	v3218 = m.ExcPending
	if v3218 != 0 {
		goto L5
	} else {
		goto L752
	}
L750:
	;
	goto L751
L751:
	;
	v3225 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	v3226 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v3228 = int32(0)
	v3233 = *(*int32)(unsafe.Add(mBase, uint32(v3225)+188))
	v3234 = *(*int32)(unsafe.Add(mBase, uint32(v3233)+140))
	v3235 = m.T0[v3234].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v3225, v3226, v3180, int32(1), v3228, v3228, v3228, int32(-1), int32(_a_F_bt_index_check_callback_79), v402, v3190)
	mBase = m.M
	v3236 = m.ExcPending
	if v3236 != 0 {
		goto L5
	} else {
		goto L754
	}
L752:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(586), int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v3223 = m.ExcPending
	if v3223 != 0 {
		goto L5
	} else {
		goto L753
	}
L753:
	;
	goto L751
L754:
	;
	v3239 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3240 = m.ExcPending
	if v3240 != 0 {
		goto L5
	} else {
		goto L755
	}
L755:
	;
	if v3239 != 0 {
		goto L756
	} else {
		goto L757
	}
L756:
	;
	v3241 = *(*int64)(unsafe.Add(mBase, uint32(v402)+64))
	v3242 = *(*int32)(unsafe.Add(mBase, uint32(v423)+48))
	v3243 = int64(0)
	v3244 = *(*int32)(unsafe.Add(mBase, uint32(v402)+60))
	v3246 = v3244 + int32(24)
	v3247 = *(*int64)(unsafe.Add(mBase, uint32(v3244)+16))
	v3250 = base.I32_wrap_i64(int64(base.Ui64(v3247) >> (uint(int64(3)) % 64)))
	if v3250 <= int32(3) {
		goto L760
	} else {
		goto L761
	}
L757:
	;
	goto L758
L758:
	;
	v3622 = *(*int32)(unsafe.Add(mBase, uint32(v402)+60))
	F_pfree(m, v3622)
	mBase = m.M
	v3624 = m.ExcPending
	if v3624 != 0 {
		goto L5
	} else {
		goto L802
	}
L759:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v401)+16)) = base.F64_mul(base.F64_div(base.F64_convert_i64_u(v3568), base.F64_convert_i64_u(v3569)), float64(100))
	*(*int32)(unsafe.Add(mBase, uint32(v401)+8)) = v3242 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v401))) = v3241
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_80), v401)
	mBase = m.M
	v3584 = m.ExcPending
	if v3584 != 0 {
		goto L5
	} else {
		goto L800
	}
L760:
	;
	if v3250 == int32(0) {
		v3568 = v3243
		v3569 = v3247
		goto L759
	} else {
		goto L763
	}
L761:
	;
	goto L762
L762:
	;
	v3390 = int64(0)
	if base.B2i32(v3246 != (v3244+int32(27))&int32(-4))|base.B2i32(v3250 < int32(4)) != 0 {
		v3469 = v3246
		v3470 = v3250
		v3475 = v3390
		goto L775
	} else {
		goto L776
	}
L763:
	;
	v3256 = v3250 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v3250) {
		goto L764
	} else {
		goto L765
	}
L764:
	;
	v3264 = v3246
	v3265 = int32(0)
	v3290 = v3243
	goto L767
L765:
	;
	v3315 = v3246
	v3341 = v3243
	goto L766
L766:
	;
	v3346 = int32(0)
	v3348 = v3315
	v3374 = v3341
	goto L771
L767:
	;
	v3294 = int32(4)
	v3295 = v3264 + v3294
	v3296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3264)+3)))
	v3297 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3296)+uint32(_c_F_bt_index_check_callback[11]))))
	v3298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3264)+2)))
	v3299 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3298)+uint32(_c_F_bt_index_check_callback[11]))))
	v3300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3264)+1)))
	v3301 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3300)+uint32(_c_F_bt_index_check_callback[11]))))
	v3302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3264))))
	v3303 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3302)+uint32(_c_F_bt_index_check_callback[11]))))
	v3307 = v3297 + (v3299 + (v3301 + (v3290 + v3303)))
	v3309 = v3265 + v3294
	if v3309 != v3250&int32(-4) {
		v3264 = v3295
		v3265 = v3309
		v3290 = v3307
		goto L767
	} else {
		goto L769
	}
L768:
	;
	if v3256 == int32(0) {
		v3568 = v3307
		v3569 = v3247
		goto L759
	} else {
		goto L770
	}
L769:
	;
	goto L768
L770:
	;
	v3315 = v3295
	v3341 = v3307
	goto L766
L771:
	;
	v3378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3348))))
	v3379 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3378)+uint32(_c_F_bt_index_check_callback[11]))))
	v3380 = v3374 + v3379
	v3381 = int32(1)
	v3384 = v3346 + v3381
	if v3384 != v3256 {
		v3346 = v3384
		v3348 = v3348 + v3381
		v3374 = v3380
		goto L771
	} else {
		goto L773
	}
L772:
	;
	v3568 = v3380
	v3569 = v3247
	goto L759
L773:
	;
	goto L772
L774:
	;
	v3539 = *(*int64)(unsafe.Add(mBase, uint32(v3244)+16))
	v3568 = v3538
	v3569 = v3539
	goto L759
L775:
	;
	if v3470 == int32(0) {
		v3538 = v3475
		goto L787
	} else {
		goto L788
	}
L776:
	;
	v3400 = v3250 - int32(4)
	v3404 = int32(base.Ui32(v3400)>>(uint(int32(2))%32)) + int32(1)
	v3406 = v3404 & int32(3)
	if base.Ui32(int32(12)) <= base.Ui32(v3400) {
		goto L777
	} else {
		goto L778
	}
L777:
	;
	v3411 = v3246
	v3412 = v3250
	v3415 = int32(0)
	v3417 = v3390
	goto L780
L778:
	;
	v3443 = v3246
	v3444 = v3250
	v3449 = v3390
	goto L779
L779:
	;
	v3451 = v3443
	v3452 = v3444
	v3453 = int32(0)
	v3457 = v3449
	goto L784
L780:
	;
	v3418 = int32(16)
	v3419 = v3412 - v3418
	v3421 = v3411 + v3418
	v3422 = *(*int32)(unsafe.Add(mBase, uint32(v3411)+12))
	v3425 = *(*int32)(unsafe.Add(mBase, uint32(v3411)+8))
	v3428 = *(*int32)(unsafe.Add(mBase, uint32(v3411)+4))
	v3431 = *(*int32)(unsafe.Add(mBase, uint32(v3411)))
	v3437 = base.I64_extend_i32_u(base.I32_popcnt(v3422)) + (base.I64_extend_i32_u(base.I32_popcnt(v3425)) + (base.I64_extend_i32_u(base.I32_popcnt(v3428)) + (v3417 + base.I64_extend_i32_u(base.I32_popcnt(v3431)))))
	v3439 = v3415 + int32(4)
	if v3439 != v3404&int32(2147483644) {
		v3411 = v3421
		v3412 = v3419
		v3415 = v3439
		v3417 = v3437
		goto L780
	} else {
		goto L782
	}
L781:
	;
	if v3406 == int32(0) {
		v3469 = v3421
		v3470 = v3419
		v3475 = v3437
		goto L775
	} else {
		goto L783
	}
L782:
	;
	goto L781
L783:
	;
	v3443 = v3421
	v3444 = v3419
	v3449 = v3437
	goto L779
L784:
	;
	v3458 = int32(4)
	v3459 = v3452 - v3458
	v3461 = v3451 + v3458
	v3462 = *(*int32)(unsafe.Add(mBase, uint32(v3451)))
	v3465 = v3457 + base.I64_extend_i32_u(base.I32_popcnt(v3462))
	v3467 = v3453 + int32(1)
	if v3467 != v3406 {
		v3451 = v3461
		v3452 = v3459
		v3453 = v3467
		v3457 = v3465
		goto L784
	} else {
		goto L786
	}
L785:
	;
	v3469 = v3461
	v3470 = v3459
	v3475 = v3465
	goto L775
L786:
	;
	goto L785
L787:
	;
	goto L774
L788:
	;
	v3479 = v3470 & int32(3)
	if v3479 == int32(0) {
		goto L790
	} else {
		goto L791
	}
L789:
	;
	if base.Ui32(v3470) < base.Ui32(int32(4)) {
		v3538 = v3506
		goto L787
	} else {
		goto L796
	}
L790:
	;
	v3500 = v3469
	v3502 = v3470
	v3506 = v3475
	goto L789
L791:
	;
	goto L792
L792:
	;
	v3483 = v3469
	v3485 = v3470
	v3487 = int32(0)
	v3489 = v3475
	goto L793
L793:
	;
	v3490 = int32(1)
	v3491 = v3483 + v3490
	v3493 = v3485 - v3490
	v3494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3483))))
	v3495 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3494)+uint32(_c_F_bt_index_check_callback[11]))))
	v3496 = v3489 + v3495
	v3498 = v3487 + v3490
	if v3498 != v3479 {
		v3483 = v3491
		v3485 = v3493
		v3487 = v3498
		v3489 = v3496
		goto L793
	} else {
		goto L795
	}
L794:
	;
	v3500 = v3491
	v3502 = v3493
	v3506 = v3496
	goto L789
L795:
	;
	goto L794
L796:
	;
	v3509 = v3500
	v3511 = v3502
	v3515 = v3506
	goto L797
L797:
	;
	v3516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3509)+3)))
	v3517 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3516)+uint32(_c_F_bt_index_check_callback[11]))))
	v3518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3509)+2)))
	v3519 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3518)+uint32(_c_F_bt_index_check_callback[11]))))
	v3520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3509)+1)))
	v3521 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3520)+uint32(_c_F_bt_index_check_callback[11]))))
	v3522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3509))))
	v3523 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3522)+uint32(_c_F_bt_index_check_callback[11]))))
	v3527 = v3517 + (v3519 + (v3521 + (v3515 + v3523)))
	v3528 = int32(4)
	v3531 = v3511 - v3528
	if v3531 != 0 {
		v3509 = v3509 + v3528
		v3511 = v3531
		v3515 = v3527
		goto L797
	} else {
		goto L799
	}
L798:
	;
	v3538 = v3527
	goto L787
L799:
	;
	goto L798
L800:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(594), int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v3589 = m.ExcPending
	if v3589 != 0 {
		goto L5
	} else {
		goto L801
	}
L801:
	;
	goto L758
L802:
	;
	goto L745
L803:
	;
	F_UnregisterSnapshot(m, v3657)
	mBase = m.M
	v3659 = m.ExcPending
	if v3659 != 0 {
		goto L5
	} else {
		goto L806
	}
L804:
	;
	goto L805
L805:
	;
	v3660 = *(*int32)(unsafe.Add(mBase, uint32(v402)+16))
	F_MemoryContextDelete(m, v3660)
	mBase = m.M
	v3662 = m.ExcPending
	if v3662 != 0 {
		goto L5
	} else {
		goto L807
	}
L806:
	;
	goto L805
L807:
	;
	m.G0 = v401 + int32(1168)
	goto L49
L808:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3677 = m.ExcPending
	if v3677 != 0 {
		goto L5
	} else {
		goto L809
	}
L809:
	;
	v3678 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v3679 = *(*int32)(unsafe.Add(mBase, uint32(v3678)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+240)) = v3679 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_81), v468+int32(240))
	mBase = m.M
	v3687 = m.ExcPending
	if v3687 != 0 {
		goto L5
	} else {
		goto L810
	}
L810:
	;
	v3688 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v3689 = *(*int64)(unsafe.Add(mBase, uint32(v469)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+224)) = uint32(v3689)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+216)) = v2667 & int32(_a_F_bt_index_check_callback_33)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+212)) = v2638
	*(*int32)(unsafe.Add(mBase, uint32(v468)+208)) = v3688
	v3697 = int64(base.Ui64(v3689) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+220)) = uint32(v3697)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_82), v468+int32(208))
	mBase = m.M
	v3703 = m.ExcPending
	if v3703 != 0 {
		goto L5
	} else {
		goto L811
	}
L811:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(2539), int32(_a_F_bt_index_check_callback_70))
	mBase = m.M
	v3708 = m.ExcPending
	if v3708 != 0 {
		goto L5
	} else {
		goto L812
	}
L812:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L813:
	;
	v3726 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+564)) = v1066
	*(*int32)(unsafe.Add(mBase, uint32(v468)+560)) = v3726
	v3732 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v468+int32(560))
	mBase = m.M
	v3733 = m.ExcPending
	if v3733 != 0 {
		goto L5
	} else {
		goto L817
	}
L814:
	;
	goto L813
L815:
	;
	v3714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+5)))
	if v3714&int32(32) == int32(0) {
		v3725 = v1075
		goto L814
	} else {
		goto L816
	}
L816:
	;
	v3719 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075)+2)))
	v3720 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075))))
	v3725 = v3719 + (v1075 + v3720<<(uint(int32(16))%32))
	goto L814
L817:
	;
	v3734 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3725)+2)))
	v3735 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3725))))
	v3736 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3725)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+548)) = v3736
	*(*int32)(unsafe.Add(mBase, uint32(v468)+544)) = v3734 | v3735<<(uint(int32(16))%32)
	v3745 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v468+int32(544))
	mBase = m.M
	v3746 = m.ExcPending
	if v3746 != 0 {
		goto L5
	} else {
		goto L818
	}
L818:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3750 = m.ExcPending
	if v3750 != 0 {
		goto L5
	} else {
		goto L819
	}
L819:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3753 = m.ExcPending
	if v3753 != 0 {
		goto L5
	} else {
		goto L820
	}
L820:
	;
	v3754 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v3755 = *(*int32)(unsafe.Add(mBase, uint32(v3754)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+528)) = v3755 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_83), v468+int32(528))
	mBase = m.M
	v3763 = m.ExcPending
	if v3763 != 0 {
		goto L5
	} else {
		goto L821
	}
L821:
	;
	v3764 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1035)+12)))
	v3765 = *(*int64)(unsafe.Add(mBase, uint32(v469)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+512)) = uint32(v3765)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+504)) = v3745
	*(*int32)(unsafe.Add(mBase, uint32(v468)+496)) = v3732
	v3770 = int64(base.Ui64(v3765) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+508)) = uint32(v3770)
	if v3764&int32(1) != 0 {
		goto L822
	} else {
		goto L823
	}
L822:
	;
	v3776 = int32(_a_F_bt_index_check_callback_42)
	goto L824
L823:
	;
	v3776 = int32(_a_F_bt_index_check_callback_43)
	goto L824
L824:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v468)+500)) = v3776
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_59), v468+int32(496))
	mBase = m.M
	v3782 = m.ExcPending
	if v3782 != 0 {
		goto L5
	} else {
		goto L825
	}
L825:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1590), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v3787 = m.ExcPending
	if v3787 != 0 {
		goto L5
	} else {
		goto L826
	}
L826:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L827:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3796 = m.ExcPending
	if v3796 != 0 {
		goto L5
	} else {
		goto L828
	}
L828:
	;
	v3797 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v3798 = *(*int32)(unsafe.Add(mBase, uint32(v3797)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+848)) = v3798 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_84), v468+int32(848))
	mBase = m.M
	v3806 = m.ExcPending
	if v3806 != 0 {
		goto L5
	} else {
		goto L829
	}
L829:
	;
	v3807 = *(*int32)(unsafe.Add(mBase, uint32(v469)+36))
	v3810 = v3788 + v3789&int32(_a_F_bt_index_check_callback_21)
	v3811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3810)+7)))
	if v3811&int32(32) == int32(0) {
		goto L831
	} else {
		goto L832
	}
L830:
	;
	v3827 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v878)+12)))
	v3828 = *(*int64)(unsafe.Add(mBase, uint32(v469)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+832)) = uint32(v3828)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+820)) = v3826
	*(*int32)(unsafe.Add(mBase, uint32(v468)+816)) = v3807
	v3833 = int64(base.Ui64(v3828) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v468)+828)) = uint32(v3833)
	if v3827&int32(1) != 0 {
		goto L834
	} else {
		goto L835
	}
L831:
	;
	v3822 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v3823 = *(*int32)(unsafe.Add(mBase, uint32(v3822)+192))
	v3824 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3823)+8)))
	v3826 = v3824
	goto L830
L832:
	;
	v3816 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+4)))
	if v3816&int32(_a_F_bt_index_check_callback_10) != 0 {
		goto L831
	} else {
		goto L833
	}
L833:
	;
	v3826 = v3816 & int32(4095)
	goto L830
L834:
	;
	v3839 = int32(_a_F_bt_index_check_callback_42)
	goto L836
L835:
	;
	v3839 = int32(_a_F_bt_index_check_callback_43)
	goto L836
L836:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v468)+824)) = v3839
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_85), v468+int32(816))
	mBase = m.M
	v3845 = m.ExcPending
	if v3845 != 0 {
		goto L5
	} else {
		goto L837
	}
L837:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1278), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v3850 = m.ExcPending
	if v3850 != 0 {
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
	v3860 = m.ExcPending
	if v3860 != 0 {
		goto L5
	} else {
		goto L840
	}
L840:
	;
	v3861 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+32)) = v3861 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_86), v33+int32(-32))
	mBase = m.M
	v3869 = m.ExcPending
	if v3869 != 0 {
		goto L5
	} else {
		goto L841
	}
L841:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(322), int32(_a_F_bt_index_check_callback_3))
	mBase = m.M
	v3874 = m.ExcPending
	if v3874 != 0 {
		goto L5
	} else {
		goto L842
	}
L842:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L843:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3881 = m.ExcPending
	if v3881 != 0 {
		goto L5
	} else {
		goto L844
	}
L844:
	;
	v3882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v3882 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_87), v33+int32(-48))
	mBase = m.M
	v3890 = m.ExcPending
	if v3890 != 0 {
		goto L5
	} else {
		goto L845
	}
L845:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(330), int32(_a_F_bt_index_check_callback_3))
	mBase = m.M
	v3895 = m.ExcPending
	if v3895 != 0 {
		goto L5
	} else {
		goto L846
	}
L846:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
