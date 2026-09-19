package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BootstrapModeMain(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v387 int32
	_ = v387
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v420 int32
	_ = v420
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v453 int32
	_ = v453
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v486 int32
	_ = v486
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v528 int64
	_ = v528
	var v529 int64
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int64
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int64
	_ = v582
	var v595 int64
	_ = v595
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v604 int64
	_ = v604
	var v616 int32
	_ = v616
	var v618 int64
	_ = v618
	var v620 int32
	_ = v620
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v784 int64
	_ = v784
	var v786 int64
	_ = v786
	var v788 int64
	_ = v788
	var v790 int64
	_ = v790
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v828 int64
	_ = v828
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1125 int32
	_ = v1125
	var v1131 int64
	_ = v1131
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1231 int32
	_ = v1231
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1254 int32
	_ = v1254
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1266 int32
	_ = v1266
	var v1271 int32
	_ = v1271
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
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
	var v1396 int32
	_ = v1396
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1466 int32
	_ = v1466
	var v1479 int32
	_ = v1479
	var v1496 int32
	_ = v1496
	var v1502 int32
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1538 int32
	_ = v1538
	var v1541 int32
	_ = v1541
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1571 int32
	_ = v1571
	var v1577 int32
	_ = v1577
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1614 int32
	_ = v1614
	var v1618 int32
	_ = v1618
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1635 int32
	_ = v1635
	var v1639 int32
	_ = v1639
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1660 int32
	_ = v1660
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1677 int32
	_ = v1677
	var v1681 int32
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1698 int32
	_ = v1698
	var v1702 int32
	_ = v1702
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1719 int32
	_ = v1719
	var v1723 int32
	_ = v1723
	var v1727 int32
	_ = v1727
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1740 int32
	_ = v1740
	var v1744 int32
	_ = v1744
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1761 int32
	_ = v1761
	var v1765 int32
	_ = v1765
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1783 int32
	_ = v1783
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1819 int32
	_ = v1819
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1855 int32
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1872 int32
	_ = v1872
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1901 int32
	_ = v1901
	var v1905 int32
	_ = v1905
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1918 int32
	_ = v1918
	var v1922 int32
	_ = v1922
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1939 int32
	_ = v1939
	var v1943 int32
	_ = v1943
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1956 int32
	_ = v1956
	var v1960 int32
	_ = v1960
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1977 int32
	_ = v1977
	var v1981 int32
	_ = v1981
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1998 int32
	_ = v1998
	var v2002 int32
	_ = v2002
	var v2006 int32
	_ = v2006
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2019 int32
	_ = v2019
	var v2023 int32
	_ = v2023
	var v2027 int32
	_ = v2027
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2040 int32
	_ = v2040
	var v2044 int32
	_ = v2044
	var v2048 int32
	_ = v2048
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2061 int32
	_ = v2061
	var v2065 int32
	_ = v2065
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2082 int32
	_ = v2082
	var v2086 int32
	_ = v2086
	var v2090 int32
	_ = v2090
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2103 int32
	_ = v2103
	var v2107 int32
	_ = v2107
	var v2111 int32
	_ = v2111
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2124 int32
	_ = v2124
	var v2128 int32
	_ = v2128
	var v2132 int32
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2145 int32
	_ = v2145
	var v2149 int32
	_ = v2149
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2162 int32
	_ = v2162
	var v2166 int32
	_ = v2166
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2183 int32
	_ = v2183
	var v2187 int32
	_ = v2187
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2202 int32
	_ = v2202
	var v2206 int32
	_ = v2206
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2219 int32
	_ = v2219
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2239 int32
	_ = v2239
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2258 int32
	_ = v2258
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2277 int32
	_ = v2277
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
	var v2287 int32
	_ = v2287
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2304 int32
	_ = v2304
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2350 int32
	_ = v2350
	var v2353 int32
	_ = v2353
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2366 int32
	_ = v2366
	var v2370 int32
	_ = v2370
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2398 int32
	_ = v2398
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2415 int32
	_ = v2415
	var v2428 int32
	_ = v2428
	var v2445 int32
	_ = v2445
	var v2449 int32
	_ = v2449
	var v2451 int32
	_ = v2451
	var v2458 int32
	_ = v2458
	var v2483 int32
	_ = v2483
	var v2486 int32
	_ = v2486
	var v2490 int32
	_ = v2490
	var v2492 int32
	_ = v2492
	var v2497 int32
	_ = v2497
	var v2504 int32
	_ = v2504
	var v2525 int32
	_ = v2525
	var v2529 int32
	_ = v2529
	var v2530 int32
	_ = v2530
	var v2535 int32
	_ = v2535
	var v2537 int32
	_ = v2537
	var v2542 int32
	_ = v2542
	var v2546 int32
	_ = v2546
	var v2576 int32
	_ = v2576
	var v2580 int32
	_ = v2580
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2594 int32
	_ = v2594
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2609 int32
	_ = v2609
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2633 int32
	_ = v2633
	var v2635 int32
	_ = v2635
	var v2637 int32
	_ = v2637
	var v2639 int32
	_ = v2639
	var v2641 int32
	_ = v2641
	var v2643 int32
	_ = v2643
	var v2645 int32
	_ = v2645
	var v2647 int32
	_ = v2647
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2652 int32
	_ = v2652
	var v2654 int32
	_ = v2654
	var v2660 int32
	_ = v2660
	var v2663 int32
	_ = v2663
	var v2687 int32
	_ = v2687
	var v2690 int32
	_ = v2690
	var v2691 int32
	_ = v2691
	var v2711 int32
	_ = v2711
	var v2713 int32
	_ = v2713
	var v2718 int32
	_ = v2718
	var v2746 int32
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2751 int32
	_ = v2751
	var v2754 int32
	_ = v2754
	var v2758 int32
	_ = v2758
	var v2761 int32
	_ = v2761
	var v2778 int32
	_ = v2778
	var v2781 int32
	_ = v2781
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2796 int32
	_ = v2796
	var v2803 int32
	_ = v2803
	var v2820 int32
	_ = v2820
	var v2825 int32
	_ = v2825
	var v2827 int32
	_ = v2827
	var v2831 int32
	_ = v2831
	var v2833 int32
	_ = v2833
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2840 int32
	_ = v2840
	var v2841 int32
	_ = v2841
	var v2846 int32
	_ = v2846
	var v2848 int32
	_ = v2848
	var v2849 int32
	_ = v2849
	var v2853 int32
	_ = v2853
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2860 int32
	_ = v2860
	var v2863 int32
	_ = v2863
	var v2884 int32
	_ = v2884
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2893 int32
	_ = v2893
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2931 int32
	_ = v2931
	var v2933 int32
	_ = v2933
	var v2934 int32
	_ = v2934
	var v2935 int32
	_ = v2935
	var v2944 int32
	_ = v2944
	var v2945 int32
	_ = v2945
	var v2946 int32
	_ = v2946
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2954 int32
	_ = v2954
	var v2958 int32
	_ = v2958
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2968 int32
	_ = v2968
	var v2969 int32
	_ = v2969
	var v2972 int32
	_ = v2972
	var v2973 int32
	_ = v2973
	var v2974 int32
	_ = v2974
	var v2979 int32
	_ = v2979
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3012 int32
	_ = v3012
	var v3018 int32
	_ = v3018
	var v3019 int32
	_ = v3019
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3037 int32
	_ = v3037
	var v3040 int32
	_ = v3040
	var v3043 int32
	_ = v3043
	var v3048 int32
	_ = v3048
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3079 int32
	_ = v3079
	var v3101 int32
	_ = v3101
	var v3102 int32
	_ = v3102
	var v3105 int32
	_ = v3105
	var v3108 int32
	_ = v3108
	var v3109 int32
	_ = v3109
	var v3110 int32
	_ = v3110
	var v3114 int32
	_ = v3114
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3118 int32
	_ = v3118
	var v3119 int32
	_ = v3119
	var v3120 int32
	_ = v3120
	var v3121 int32
	_ = v3121
	var v3126 int32
	_ = v3126
	var v3128 int32
	_ = v3128
	var v3130 int32
	_ = v3130
	var v3132 int32
	_ = v3132
	var v3133 int32
	_ = v3133
	var v3137 int32
	_ = v3137
	var v3141 int32
	_ = v3141
	var v3143 int32
	_ = v3143
	var v3146 int32
	_ = v3146
	var v3153 int32
	_ = v3153
	var v3155 int32
	_ = v3155
	var v3158 int32
	_ = v3158
	var v3161 int32
	_ = v3161
	var v3162 int32
	_ = v3162
	var v3164 int32
	_ = v3164
	var v3166 int32
	_ = v3166
	var v3167 int32
	_ = v3167
	var v3170 int32
	_ = v3170
	var v3171 int32
	_ = v3171
	var v3173 int32
	_ = v3173
	var v3175 int32
	_ = v3175
	var v3178 int32
	_ = v3178
	var v3182 int32
	_ = v3182
	var v3183 int32
	_ = v3183
	var v3187 int32
	_ = v3187
	var v3195 int32
	_ = v3195
	var v3196 int32
	_ = v3196
	var v3199 int32
	_ = v3199
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3203 int32
	_ = v3203
	var v3204 int32
	_ = v3204
	var v3207 int32
	_ = v3207
	var v3208 int32
	_ = v3208
	var v3210 int32
	_ = v3210
	var v3213 int32
	_ = v3213
	var v3214 int32
	_ = v3214
	var v3215 int32
	_ = v3215
	var v3218 int32
	_ = v3218
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3245 int32
	_ = v3245
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3249 int32
	_ = v3249
	var v3252 int32
	_ = v3252
	var v3254 int32
	_ = v3254
	var v3255 int32
	_ = v3255
	var v3259 int32
	_ = v3259
	var v3260 int32
	_ = v3260
	var v3266 int32
	_ = v3266
	var v3267 int32
	_ = v3267
	var v3269 int32
	_ = v3269
	var v3271 int32
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3275 int32
	_ = v3275
	var v3278 int32
	_ = v3278
	var v3279 int32
	_ = v3279
	var v3281 int32
	_ = v3281
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3288 int32
	_ = v3288
	var v3289 int32
	_ = v3289
	var v3290 int32
	_ = v3290
	var v3294 int32
	_ = v3294
	var v3295 int32
	_ = v3295
	var v3298 int32
	_ = v3298
	var v3299 int32
	_ = v3299
	var v3300 int32
	_ = v3300
	var v3308 int32
	_ = v3308
	var v3309 int32
	_ = v3309
	var v3311 int32
	_ = v3311
	var v3312 int32
	_ = v3312
	var v3313 int32
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3322 int32
	_ = v3322
	var v3325 int32
	_ = v3325
	var v3344 int32
	_ = v3344
	var v3345 int32
	_ = v3345
	var v3346 int32
	_ = v3346
	var v3348 int32
	_ = v3348
	var v3351 int32
	_ = v3351
	var v3355 int32
	_ = v3355
	var v3358 int32
	_ = v3358
	var v3359 int32
	_ = v3359
	var v3364 int32
	_ = v3364
	var v3368 int32
	_ = v3368
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3396 int32
	_ = v3396
	var v3399 int32
	_ = v3399
	var v3400 int32
	_ = v3400
	var v3402 int32
	_ = v3402
	var v3403 int32
	_ = v3403
	var v3407 int32
	_ = v3407
	var v3408 int32
	_ = v3408
	var v3413 int32
	_ = v3413
	var v3426 int32
	_ = v3426
	var v3443 int32
	_ = v3443
	var v3447 int32
	_ = v3447
	var v3449 int32
	_ = v3449
	var v3451 int32
	_ = v3451
	var v3452 int32
	_ = v3452
	var v3453 int32
	_ = v3453
	var v3454 int32
	_ = v3454
	var v3458 int32
	_ = v3458
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3463 int32
	_ = v3463
	var v3480 int32
	_ = v3480
	var v3485 int32
	_ = v3485
	var v3486 int32
	_ = v3486
	var v3487 int32
	_ = v3487
	var v3488 int32
	_ = v3488
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3517 int32
	_ = v3517
	var v3518 int32
	_ = v3518
	var v3519 int32
	_ = v3519
	var v3524 int32
	_ = v3524
	var v3528 int32
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3533 int32
	_ = v3533
	var v3534 int32
	_ = v3534
	var v3539 int32
	_ = v3539
	var v3543 int32
	_ = v3543
	var v3546 int32
	_ = v3546
	var v3547 int32
	_ = v3547
	var v3571 int32
	_ = v3571
	var v3574 int32
	_ = v3574
	var v3575 int32
	_ = v3575
	var v3577 int32
	_ = v3577
	var v3578 int32
	_ = v3578
	var v3582 int32
	_ = v3582
	var v3583 int32
	_ = v3583
	var v3588 int32
	_ = v3588
	var v3601 int32
	_ = v3601
	var v3618 int32
	_ = v3618
	var v3622 int32
	_ = v3622
	var v3624 int32
	_ = v3624
	var v3628 int32
	_ = v3628
	var v3634 int32
	_ = v3634
	var v3635 int32
	_ = v3635
	var v3664 int32
	_ = v3664
	var v3668 int32
	_ = v3668
	var v3674 int32
	_ = v3674
	var v3676 int32
	_ = v3676
	var v3678 int32
	_ = v3678
	var v3681 int32
	_ = v3681
	var v3682 int32
	_ = v3682
	var v3684 int32
	_ = v3684
	var v3685 int32
	_ = v3685
	var v3687 int32
	_ = v3687
	var v3688 int32
	_ = v3688
	var v3689 int32
	_ = v3689
	var v3693 int32
	_ = v3693
	var v3699 int32
	_ = v3699
	var v3706 int32
	_ = v3706
	var v3707 int32
	_ = v3707
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3712 int32
	_ = v3712
	var v3714 int32
	_ = v3714
	var v3717 int32
	_ = v3717
	var v3730 int32
	_ = v3730
	var v3732 int32
	_ = v3732
	var v3734 int32
	_ = v3734
	var v3737 int32
	_ = v3737
	var v3738 int32
	_ = v3738
	var v3740 int32
	_ = v3740
	var v3741 int32
	_ = v3741
	var v3743 int32
	_ = v3743
	var v3744 int32
	_ = v3744
	var v3745 int32
	_ = v3745
	var v3749 int32
	_ = v3749
	var v3755 int32
	_ = v3755
	var v3761 int32
	_ = v3761
	var v3763 int32
	_ = v3763
	var v3765 int32
	_ = v3765
	var v3768 int32
	_ = v3768
	var v3769 int32
	_ = v3769
	var v3771 int32
	_ = v3771
	var v3772 int32
	_ = v3772
	var v3774 int32
	_ = v3774
	var v3775 int32
	_ = v3775
	var v3776 int32
	_ = v3776
	var v3780 int32
	_ = v3780
	var v3787 int32
	_ = v3787
	var v3792 int32
	_ = v3792
	var v3796 int32
	_ = v3796
	var v3801 int32
	_ = v3801
	var v3806 int32
	_ = v3806
	var v3807 int32
	_ = v3807
	var v3809 int32
	_ = v3809
	var v3812 int32
	_ = v3812
	var v3814 int32
	_ = v3814
	var v3816 int32
	_ = v3816
	var v3818 int32
	_ = v3818
	var v3821 int32
	_ = v3821
	var v3824 int32
	_ = v3824
	var v3828 int32
	_ = v3828
	var v3830 int32
	_ = v3830
	var v3833 int32
	_ = v3833
	var v3836 int32
	_ = v3836
	var v3837 int32
	_ = v3837
	var v3845 int32
	_ = v3845
	var v3850 int32
	_ = v3850
	var v3854 int32
	_ = v3854
	var v3855 int32
	_ = v3855
	var v3857 int32
	_ = v3857
	var v3858 int32
	_ = v3858
	var v3861 int32
	_ = v3861
	var v3862 int32
	_ = v3862
	var v3874 int32
	_ = v3874
	var v3895 int32
	_ = v3895
	var v3898 int32
	_ = v3898
	var v3902 int32
	_ = v3902
	var v3904 int32
	_ = v3904
	var v3905 int32
	_ = v3905
	var v3907 int32
	_ = v3907
	var v3909 int32
	_ = v3909
	var v3910 int32
	_ = v3910
	var v3911 int32
	_ = v3911
	var v3915 int32
	_ = v3915
	var v3922 int32
	_ = v3922
	var v3925 int32
	_ = v3925
	var v3926 int32
	_ = v3926
	var v3927 int32
	_ = v3927
	var v3928 int32
	_ = v3928
	var v3929 int32
	_ = v3929
	var v3939 int32
	_ = v3939
	var v3944 int32
	_ = v3944
	var v3948 int32
	_ = v3948
	var v3950 int32
	_ = v3950
	var v3983 int32
	_ = v3983
	var v3986 int32
	_ = v3986
	var v3988 int32
	_ = v3988
	var v3990 int32
	_ = v3990
	var v3992 int32
	_ = v3992
	var v3993 int32
	_ = v3993
	var v3994 int32
	_ = v3994
	var v4000 int32
	_ = v4000
	var v4001 int32
	_ = v4001
	var v4002 int32
	_ = v4002
	var v4004 int32
	_ = v4004
	var v4009 int32
	_ = v4009
	var v4014 int32
	_ = v4014
	var v4015 int32
	_ = v4015
	var v4017 int32
	_ = v4017
	var v4020 int32
	_ = v4020
	var v4022 int32
	_ = v4022
	var v4025 int32
	_ = v4025
	var v4028 int32
	_ = v4028
	var v4030 int32
	_ = v4030
	var v4032 int32
	_ = v4032
	var v4034 int32
	_ = v4034
	var v4035 int32
	_ = v4035
	var v4036 int32
	_ = v4036
	var v4042 int32
	_ = v4042
	var v4043 int32
	_ = v4043
	var v4044 int32
	_ = v4044
	var v4046 int32
	_ = v4046
	var v4051 int32
	_ = v4051
	var v4056 int32
	_ = v4056
	var v4057 int32
	_ = v4057
	var v4059 int32
	_ = v4059
	var v4063 int32
	_ = v4063
	var v4067 int32
	_ = v4067
	var v4068 int32
	_ = v4068
	var v4073 int32
	_ = v4073
	var v4076 int32
	_ = v4076
	var v4079 int64
	_ = v4079
	var v4083 int32
	_ = v4083
	var v4087 int32
	_ = v4087
	var v4091 int32
	_ = v4091
	var v4096 int32
	_ = v4096
	var v4099 int32
	_ = v4099
	var v4102 int32
	_ = v4102
	var v4104 int32
	_ = v4104
	var v4106 int32
	_ = v4106
	var v4108 int32
	_ = v4108
	var v4109 int32
	_ = v4109
	var v4110 int32
	_ = v4110
	var v4116 int32
	_ = v4116
	var v4117 int32
	_ = v4117
	var v4118 int32
	_ = v4118
	var v4120 int32
	_ = v4120
	var v4125 int32
	_ = v4125
	var v4130 int32
	_ = v4130
	var v4131 int32
	_ = v4131
	var v4133 int32
	_ = v4133
	var v4137 int32
	_ = v4137
	var v4139 int32
	_ = v4139
	var v4140 int32
	_ = v4140
	var v4143 int32
	_ = v4143
	var v4146 int32
	_ = v4146
	var v4148 int32
	_ = v4148
	var v4151 int32
	_ = v4151
	var v4152 int32
	_ = v4152
	var v4156 int32
	_ = v4156
	var v4161 int32
	_ = v4161
	var v4164 int32
	_ = v4164
	var v4168 int32
	_ = v4168
	var v4172 int32
	_ = v4172
	var v4175 int32
	_ = v4175
	var v4176 int32
	_ = v4176
	var v4179 int32
	_ = v4179
	var v4182 int32
	_ = v4182
	var v4189 int32
	_ = v4189
	var v4190 int32
	_ = v4190
	var v4194 int32
	_ = v4194
	var v4195 int32
	_ = v4195
	var v4201 int32
	_ = v4201
	var v4206 int32
	_ = v4206
	var v4209 int32
	_ = v4209
	var v4213 int32
	_ = v4213
	var v4216 int32
	_ = v4216
	var v4219 int32
	_ = v4219
	var v4220 int32
	_ = v4220
	var v4227 int32
	_ = v4227
	var v4235 int32
	_ = v4235
	var v4236 int32
	_ = v4236
	var v4239 int32
	_ = v4239
	var v4240 int32
	_ = v4240
	var v4248 int32
	_ = v4248
	var v4253 int32
	_ = v4253
	var v4257 int32
	_ = v4257
	var v4260 int32
	_ = v4260
	var v4262 int32
	_ = v4262
	var v4264 int32
	_ = v4264
	var v4266 int32
	_ = v4266
	var v4267 int32
	_ = v4267
	var v4268 int32
	_ = v4268
	var v4274 int32
	_ = v4274
	var v4275 int32
	_ = v4275
	var v4276 int32
	_ = v4276
	var v4278 int32
	_ = v4278
	var v4283 int32
	_ = v4283
	var v4288 int32
	_ = v4288
	var v4289 int32
	_ = v4289
	var v4291 int32
	_ = v4291
	var v4296 int32
	_ = v4296
	var v4297 int32
	_ = v4297
	var v4301 int32
	_ = v4301
	var v4306 int32
	_ = v4306
	var v4311 int32
	_ = v4311
	var v4313 int32
	_ = v4313
	var v4316 int32
	_ = v4316
	var v4319 int32
	_ = v4319
	var v4321 int32
	_ = v4321
	var v4325 int32
	_ = v4325
	var v4326 int32
	_ = v4326
	var v4328 int32
	_ = v4328
	var v4332 int32
	_ = v4332
	var v4337 int32
	_ = v4337
	var v4339 int32
	_ = v4339
	var v4341 int32
	_ = v4341
	var v4342 int32
	_ = v4342
	var v4345 int32
	_ = v4345
	var v4346 int32
	_ = v4346
	var v4348 int32
	_ = v4348
	var v4350 int32
	_ = v4350
	var v4352 int32
	_ = v4352
	var v4354 int32
	_ = v4354
	var v4357 int32
	_ = v4357
	var v4358 int32
	_ = v4358
	var v4362 int32
	_ = v4362
	var v4367 int32
	_ = v4367
	var v4369 int32
	_ = v4369
	var v4370 int32
	_ = v4370
	var v4385 int32
	_ = v4385
	var v4388 int32
	_ = v4388
	var v4390 int32
	_ = v4390
	var v4392 int32
	_ = v4392
	var v4394 int32
	_ = v4394
	var v4395 int32
	_ = v4395
	var v4396 int32
	_ = v4396
	var v4402 int32
	_ = v4402
	var v4403 int32
	_ = v4403
	var v4404 int32
	_ = v4404
	var v4406 int32
	_ = v4406
	var v4407 int32
	_ = v4407
	var v4412 int32
	_ = v4412
	var v4413 int32
	_ = v4413
	var v4416 int32
	_ = v4416
	var v4422 int32
	_ = v4422
	var v4427 int32
	_ = v4427
	var v4429 int32
	_ = v4429
	var v4434 int32
	_ = v4434
	var v4439 int32
	_ = v4439
	var v4440 int32
	_ = v4440
	var v4442 int32
	_ = v4442
	var v4447 int32
	_ = v4447
	var v4452 int32
	_ = v4452
	var v4454 int32
	_ = v4454
	var v4455 int32
	_ = v4455
	var v4459 int32
	_ = v4459
	var v4460 int32
	_ = v4460
	var v4465 int32
	_ = v4465
	var v4469 int64
	_ = v4469
	var v4489 int32
	_ = v4489
	var v4490 int32
	_ = v4490
	var v4493 int32
	_ = v4493
	var v4494 int32
	_ = v4494
	var v4503 int32
	_ = v4503
	var v4506 int32
	_ = v4506
	var v4509 int32
	_ = v4509
	var v4511 int32
	_ = v4511
	var v4513 int32
	_ = v4513
	var v4515 int32
	_ = v4515
	var v4516 int32
	_ = v4516
	var v4517 int32
	_ = v4517
	var v4523 int32
	_ = v4523
	var v4524 int32
	_ = v4524
	var v4525 int32
	_ = v4525
	var v4527 int32
	_ = v4527
	var v4528 int32
	_ = v4528
	var v4533 int32
	_ = v4533
	var v4534 int32
	_ = v4534
	var v4537 int32
	_ = v4537
	var v4543 int32
	_ = v4543
	var v4548 int32
	_ = v4548
	var v4550 int32
	_ = v4550
	var v4555 int32
	_ = v4555
	var v4560 int32
	_ = v4560
	var v4561 int32
	_ = v4561
	var v4563 int32
	_ = v4563
	var v4568 int32
	_ = v4568
	var v4573 int32
	_ = v4573
	var v4575 int32
	_ = v4575
	var v4576 int32
	_ = v4576
	var v4580 int32
	_ = v4580
	var v4581 int32
	_ = v4581
	var v4586 int32
	_ = v4586
	var v4587 int64
	_ = v4587
	var v4602 int32
	_ = v4602
	var v4612 int32
	_ = v4612
	var v4613 int32
	_ = v4613
	var v4616 int32
	_ = v4616
	var v4617 int32
	_ = v4617
	var v4626 int32
	_ = v4626
	var v4629 int32
	_ = v4629
	var v4632 int32
	_ = v4632
	var v4634 int32
	_ = v4634
	var v4636 int32
	_ = v4636
	var v4638 int32
	_ = v4638
	var v4639 int32
	_ = v4639
	var v4640 int32
	_ = v4640
	var v4646 int32
	_ = v4646
	var v4647 int32
	_ = v4647
	var v4648 int32
	_ = v4648
	var v4651 int32
	_ = v4651
	var v4652 int32
	_ = v4652
	var v4653 int32
	_ = v4653
	var v4659 int32
	_ = v4659
	var v4664 int32
	_ = v4664
	var v4666 int32
	_ = v4666
	var v4671 int32
	_ = v4671
	var v4676 int32
	_ = v4676
	var v4677 int32
	_ = v4677
	var v4679 int32
	_ = v4679
	var v4682 int32
	_ = v4682
	var v4685 int32
	_ = v4685
	var v4688 int32
	_ = v4688
	var v4689 int32
	_ = v4689
	var v4691 int32
	_ = v4691
	var v4695 int32
	_ = v4695
	var v4696 int32
	_ = v4696
	var v4698 int32
	_ = v4698
	var v4699 int32
	_ = v4699
	var v4700 int32
	_ = v4700
	var v4701 int32
	_ = v4701
	var v4707 int32
	_ = v4707
	var v4711 int32
	_ = v4711
	var v4716 int32
	_ = v4716
	var v4717 int32
	_ = v4717
	var v4721 int32
	_ = v4721
	var v4722 int32
	_ = v4722
	var v4728 int32
	_ = v4728
	var v4734 int32
	_ = v4734
	var v4739 int32
	_ = v4739
	var v4742 int32
	_ = v4742
	var v4748 int32
	_ = v4748
	var v4751 int32
	_ = v4751
	var v4753 int32
	_ = v4753
	var v4755 int32
	_ = v4755
	var v4757 int32
	_ = v4757
	var v4758 int32
	_ = v4758
	var v4759 int32
	_ = v4759
	var v4765 int32
	_ = v4765
	var v4766 int32
	_ = v4766
	var v4767 int32
	_ = v4767
	var v4769 int32
	_ = v4769
	var v4774 int32
	_ = v4774
	var v4779 int32
	_ = v4779
	var v4780 int32
	_ = v4780
	var v4782 int32
	_ = v4782
	var v4786 int32
	_ = v4786
	var v4791 int32
	_ = v4791
	var v4813 int32
	_ = v4813
	var v4815 int32
	_ = v4815
	var v4816 int32
	_ = v4816
	var v4818 int32
	_ = v4818
	var v4819 int32
	_ = v4819
	var v4821 int32
	_ = v4821
	var v4822 int32
	_ = v4822
	var v4824 int32
	_ = v4824
	var v4825 int32
	_ = v4825
	var v4826 int32
	_ = v4826
	var v4829 int32
	_ = v4829
	var v4832 int32
	_ = v4832
	var v4835 int32
	_ = v4835
	var v4836 int32
	_ = v4836
	var v4838 int32
	_ = v4838
	var v4839 int32
	_ = v4839
	var v4869 int32
	_ = v4869
	var v4872 int32
	_ = v4872
	var v4874 int32
	_ = v4874
	var v4876 int32
	_ = v4876
	var v4878 int32
	_ = v4878
	var v4879 int32
	_ = v4879
	var v4880 int32
	_ = v4880
	var v4886 int32
	_ = v4886
	var v4887 int32
	_ = v4887
	var v4888 int32
	_ = v4888
	var v4891 int32
	_ = v4891
	var v4892 int32
	_ = v4892
	var v4893 int32
	_ = v4893
	var v4894 int32
	_ = v4894
	var v4895 int32
	_ = v4895
	var v4901 int32
	_ = v4901
	var v4902 int32
	_ = v4902
	var v4904 int32
	_ = v4904
	var v4905 int32
	_ = v4905
	var v4910 int32
	_ = v4910
	var v4916 int32
	_ = v4916
	var v4917 int32
	_ = v4917
	var v4918 int32
	_ = v4918
	var v4924 int32
	_ = v4924
	var v4925 int32
	_ = v4925
	var v4931 int32
	_ = v4931
	var v4932 int32
	_ = v4932
	var v4934 int32
	_ = v4934
	var v4936 int32
	_ = v4936
	var v4942 int32
	_ = v4942
	var v4945 int32
	_ = v4945
	var v4946 int32
	_ = v4946
	var v4947 int32
	_ = v4947
	var v4948 int32
	_ = v4948
	var v4950 int32
	_ = v4950
	var v4953 int32
	_ = v4953
	var v4956 int32
	_ = v4956
	var v4957 int32
	_ = v4957
	var v4961 int32
	_ = v4961
	var v4966 int32
	_ = v4966
	var v4969 int32
	_ = v4969
	var v4971 int32
	_ = v4971
	var v4974 int32
	_ = v4974
	var v4978 int32
	_ = v4978
	var v4980 int32
	_ = v4980
	var v4981 int32
	_ = v4981
	var v4983 int32
	_ = v4983
	var v4984 int32
	_ = v4984
	var v4987 int32
	_ = v4987
	var v4991 int32
	_ = v4991
	var v4996 int32
	_ = v4996
	var v4997 int32
	_ = v4997
	var v4998 int32
	_ = v4998
	var v5007 int32
	_ = v5007
	var v5012 int32
	_ = v5012
	var v5014 int32
	_ = v5014
	var v5016 int32
	_ = v5016
	var v5019 int32
	_ = v5019
	var v5031 int32
	_ = v5031
	var v5051 int32
	_ = v5051
	var v5058 int32
	_ = v5058
	var v5059 int32
	_ = v5059
	var v5060 int32
	_ = v5060
	var v5061 int32
	_ = v5061
	var v5062 int32
	_ = v5062
	var v5064 int32
	_ = v5064
	var v5070 int32
	_ = v5070
	var v5073 int32
	_ = v5073
	var v5074 int32
	_ = v5074
	var v5075 int32
	_ = v5075
	var v5080 int32
	_ = v5080
	var v5082 int32
	_ = v5082
	var v5085 int32
	_ = v5085
	var v5089 int32
	_ = v5089
	var v5090 int32
	_ = v5090
	var v5101 int32
	_ = v5101
	var v5106 int32
	_ = v5106
	var v5107 int32
	_ = v5107
	var v5113 int32
	_ = v5113
	var v5118 int32
	_ = v5118
	var v5120 int32
	_ = v5120
	var v5121 int32
	_ = v5121
	var v5123 int32
	_ = v5123
	var v5134 int32
	_ = v5134
	var v5135 int32
	_ = v5135
	var v5152 int32
	_ = v5152
	var v5155 int32
	_ = v5155
	var v5165 int32
	_ = v5165
	var v5185 int32
	_ = v5185
	var v5187 int32
	_ = v5187
	var v5194 int32
	_ = v5194
	var v5195 int32
	_ = v5195
	var v5196 int32
	_ = v5196
	var v5197 int32
	_ = v5197
	var v5198 int32
	_ = v5198
	var v5200 int32
	_ = v5200
	var v5206 int32
	_ = v5206
	var v5209 int32
	_ = v5209
	var v5210 int32
	_ = v5210
	var v5211 int32
	_ = v5211
	var v5216 int32
	_ = v5216
	var v5218 int32
	_ = v5218
	var v5221 int32
	_ = v5221
	var v5225 int32
	_ = v5225
	var v5226 int32
	_ = v5226
	var v5237 int32
	_ = v5237
	var v5266 int32
	_ = v5266
	var v5271 int32
	_ = v5271
	var v5273 int32
	_ = v5273
	var v5276 int32
	_ = v5276
	var v5279 int32
	_ = v5279
	var v5290 int32
	_ = v5290
	var v5310 int32
	_ = v5310
	var v5312 int32
	_ = v5312
	var v5319 int32
	_ = v5319
	var v5320 int32
	_ = v5320
	var v5321 int32
	_ = v5321
	var v5322 int32
	_ = v5322
	var v5323 int32
	_ = v5323
	var v5325 int32
	_ = v5325
	var v5331 int32
	_ = v5331
	var v5334 int32
	_ = v5334
	var v5335 int32
	_ = v5335
	var v5336 int32
	_ = v5336
	var v5341 int32
	_ = v5341
	var v5343 int32
	_ = v5343
	var v5346 int32
	_ = v5346
	var v5350 int32
	_ = v5350
	var v5351 int32
	_ = v5351
	var v5362 int32
	_ = v5362
	var v5393 int32
	_ = v5393
	var v5397 int32
	_ = v5397
	var v5402 int32
	_ = v5402
	var v5403 int32
	_ = v5403
	var v5405 int32
	_ = v5405
	var v5406 int32
	_ = v5406
	var v5408 int32
	_ = v5408
	var v5409 int32
	_ = v5409
	var v5411 int32
	_ = v5411
	var v5412 int32
	_ = v5412
	var v5414 int32
	_ = v5414
	var v5415 int32
	_ = v5415
	var v5417 int32
	_ = v5417
	var v5418 int32
	_ = v5418
	var v5420 int32
	_ = v5420
	var v5421 int32
	_ = v5421
	var v5424 int32
	_ = v5424
	var v5425 int32
	_ = v5425
	var v5427 int32
	_ = v5427
	var v5432 int32
	_ = v5432
	var v5436 int32
	_ = v5436
	var v5461 int32
	_ = v5461
	var v5464 int32
	_ = v5464
	var v5465 int32
	_ = v5465
	var v5467 int32
	_ = v5467
	var v5468 int32
	_ = v5468
	var v5470 int32
	_ = v5470
	var v5471 int32
	_ = v5471
	var v5473 int32
	_ = v5473
	var v5474 int32
	_ = v5474
	var v5476 int32
	_ = v5476
	var v5477 int32
	_ = v5477
	var v5479 int32
	_ = v5479
	var v5482 int32
	_ = v5482
	var v5483 int32
	_ = v5483
	var v5485 int32
	_ = v5485
	var v5488 int32
	_ = v5488
	var v5491 int32
	_ = v5491
	var v5492 int32
	_ = v5492
	var v5494 int32
	_ = v5494
	var v5521 int32
	_ = v5521
	var v5524 int32
	_ = v5524
	var v5547 int32
	_ = v5547
	var v5548 int32
	_ = v5548
	var v5551 int32
	_ = v5551
	var v5552 int32
	_ = v5552
	var v5555 int32
	_ = v5555
	var v5556 int32
	_ = v5556
	var v5559 int32
	_ = v5559
	var v5563 int32
	_ = v5563
	var v5566 int32
	_ = v5566
	var v5578 int32
	_ = v5578
	var v5597 int32
	_ = v5597
	var v5598 int32
	_ = v5598
	var v5601 int32
	_ = v5601
	var v5605 int32
	_ = v5605
	var v5616 int32
	_ = v5616
	var v5638 int32
	_ = v5638
	var v5692 int32
	_ = v5692
	var v5696 int64
	_ = v5696
	var v5698 int32
	_ = v5698
	var v5699 int32
	_ = v5699
	var v5701 int32
	_ = v5701
	var v5705 int32
	_ = v5705
	var v5707 int32
	_ = v5707
	var v5711 int32
	_ = v5711
	var v5712 int32
	_ = v5712
	var v5719 int32
	_ = v5719
	var v5724 int32
	_ = v5724
	var v5726 int32
	_ = v5726
	var v5727 int32
	_ = v5727
	var v5728 int32
	_ = v5728
	var v5735 int32
	_ = v5735
	var v5751 int32
	_ = v5751
	var v5753 int32
	_ = v5753
	var v5756 int32
	_ = v5756
	var v5757 int32
	_ = v5757
	var v5759 int32
	_ = v5759
	var v5760 int32
	_ = v5760
	var v5764 int32
	_ = v5764
	var v5765 int32
	_ = v5765
	var v5766 int32
	_ = v5766
	var v5767 int32
	_ = v5767
	var v5768 int32
	_ = v5768
	var v5769 int32
	_ = v5769
	var v5773 int32
	_ = v5773
	var v5778 int32
	_ = v5778
	var v5782 int32
	_ = v5782
	var v5784 int32
	_ = v5784
	var v5788 int32
	_ = v5788
	var v5790 int32
	_ = v5790
	var v5794 int32
	_ = v5794
	var v5795 int32
	_ = v5795
	var v5801 int32
	_ = v5801
	var v5806 int32
	_ = v5806
	var v5808 int32
	_ = v5808
	var v5810 int32
	_ = v5810
	var v5811 int32
	_ = v5811
	var v5812 int32
	_ = v5812
	var v5817 int32
	_ = v5817
	var v5823 int32
	_ = v5823
	var v5825 int32
	_ = v5825
	var v5826 int32
	_ = v5826
	var v5827 int32
	_ = v5827
	var v5828 int32
	_ = v5828
	var v5829 int32
	_ = v5829
	var v5841 int32
	_ = v5841
	var v5846 int32
	_ = v5846
	var v5855 int32
	_ = v5855
	var v5860 int32
	_ = v5860
	var v5861 int32
	_ = v5861
	var v5862 int32
	_ = v5862
	var v5863 int32
	_ = v5863
	var v5864 int32
	_ = v5864
	var v5865 int32
	_ = v5865
	var v5866 int32
	_ = v5866
	var v5867 int32
	_ = v5867
	var v5868 int32
	_ = v5868
	var v5869 int32
	_ = v5869
	var v5870 int32
	_ = v5870
	var v5871 int32
	_ = v5871
	var v5872 int32
	_ = v5872
	var v5873 int32
	_ = v5873
	var v5874 int32
	_ = v5874
	var v5875 int32
	_ = v5875
	var v5876 int32
	_ = v5876
	var v5877 int32
	_ = v5877
	var v5878 int32
	_ = v5878
	var v5879 int32
	_ = v5879
	var v5880 int32
	_ = v5880
	var v5881 int32
	_ = v5881
	var v5882 int32
	_ = v5882
	var v5883 int32
	_ = v5883
	var v5884 int32
	_ = v5884
	var v5885 int32
	_ = v5885
	var v5886 int32
	_ = v5886
	var v5887 int32
	_ = v5887
	var v5888 int32
	_ = v5888
	var v5889 int32
	_ = v5889
	var v5890 int32
	_ = v5890
	var v5891 int32
	_ = v5891
	var v5892 int32
	_ = v5892
	var v5893 int32
	_ = v5893
	var v5894 int32
	_ = v5894
	var v5895 int32
	_ = v5895
	var v5896 int32
	_ = v5896
	var v5897 int32
	_ = v5897
	var v5898 int32
	_ = v5898
	var v5899 int32
	_ = v5899
	var v5900 int32
	_ = v5900
	var v5901 int32
	_ = v5901
	var v5902 int32
	_ = v5902
	var v5903 int32
	_ = v5903
	var v5904 int32
	_ = v5904
	var v5905 int32
	_ = v5905
	var v5906 int32
	_ = v5906
	var v5907 int32
	_ = v5907
	var v5908 int32
	_ = v5908
	var v5909 int32
	_ = v5909
	var v5910 int32
	_ = v5910
	var v5911 int32
	_ = v5911
	var v5912 int32
	_ = v5912
	var v5913 int32
	_ = v5913
	var v5914 int32
	_ = v5914
	var v5915 int32
	_ = v5915
	var v5916 int32
	_ = v5916
	var v5917 int32
	_ = v5917
	var v5920 int32
	_ = v5920
	var v5946 int32
	_ = v5946
	var v5949 int32
	_ = v5949
	var v5950 int32
	_ = v5950
	var v5951 int32
	_ = v5951
	var v5954 int32
	_ = v5954
	var v5957 int32
	_ = v5957
	var v5958 int32
	_ = v5958
	var v5961 int32
	_ = v5961
	var v5967 int32
	_ = v5967
	var v5970 int32
	_ = v5970
	var v5974 int32
	_ = v5974
	var v5976 int32
	_ = v5976
	var v5978 int32
	_ = v5978
	var v5981 int32
	_ = v5981
	var v5982 int32
	_ = v5982
	var v5984 int32
	_ = v5984
	var v5985 int32
	_ = v5985
	var v5987 int32
	_ = v5987
	var v5988 int32
	_ = v5988
	var v5989 int32
	_ = v5989
	var v5993 int32
	_ = v5993
	var v6001 int32
	_ = v6001
	var v6014 int32
	_ = v6014
	var v6015 int32
	_ = v6015
	var v6018 int32
	_ = v6018
	var v6044 int32
	_ = v6044
	var v6063 int32
	_ = v6063
	var v6067 int32
	_ = v6067
	var v6069 int32
	_ = v6069
	var v6072 int32
	_ = v6072
	var v6078 int32
	_ = v6078
	var v6083 int32
	_ = v6083
	var v6087 int32
	_ = v6087
	var v6091 int32
	_ = v6091
	var v6096 int32
	_ = v6096
	var v6100 int32
	_ = v6100
	var v6104 int32
	_ = v6104
	var v6109 int32
	_ = v6109
	var v6120 int32
	_ = v6120
	var v6121 int32
	_ = v6121
	var v6136 int32
	_ = v6136
	var v6146 int32
	_ = v6146
	var v6166 int32
	_ = v6166
	var v6171 int32
	_ = v6171
	var v6173 int32
	_ = v6173
	var v6177 int32
	_ = v6177
	var v6178 int32
	_ = v6178
	var v6180 int32
	_ = v6180
	var v6187 int32
	_ = v6187
	var v6189 int32
	_ = v6189
	var v6193 int32
	_ = v6193
	var v6195 int32
	_ = v6195
	var v6197 int32
	_ = v6197
	var v6199 int32
	_ = v6199
	var v6201 int32
	_ = v6201
	var v6205 int32
	_ = v6205
	var v6207 int32
	_ = v6207
	var v6210 int32
	_ = v6210
	var v6213 int32
	_ = v6213
	var v6217 int32
	_ = v6217
	var v6220 int32
	_ = v6220
	var v6222 int32
	_ = v6222
	var v6228 int32
	_ = v6228
	var v6233 int32
	_ = v6233
	var v6239 int32
	_ = v6239
	var v6244 int32
	_ = v6244
	var v6248 int32
	_ = v6248
	var v6251 int32
	_ = v6251
	var v6257 int32
	_ = v6257
	var v6260 int32
	_ = v6260
	var v6263 int32
	_ = v6263
	var v6269 int32
	_ = v6269
	var v6273 int32
	_ = v6273
	var v6277 int32
	_ = v6277
	var v6282 int32
	_ = v6282
	v4 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(96)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_InitStandaloneProcess(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v35 = l0 - int32(1)
	F_InitializeGUCOptions(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v42 = int32(0)
	v44 = v4
	goto L11
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6273 = m.ExcPending
	if v6273 != 0 {
		goto L1
	} else {
		goto L1197
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[0])) = int32(2)
	F_proc_exit(m, int32(0))
	mBase = m.M
	v6269 = m.ExcPending
	if v6269 != 0 {
		goto L1
	} else {
		goto L1196
	}
L6:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v6263 = m.ExcPending
	if v6263 != 0 {
		goto L1
	} else {
		goto L1195
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v31
	F_write_stderr(m, int32(_a_F_BootstrapModeMain_0), v29+int32(16))
	mBase = m.M
	v6257 = m.ExcPending
	if v6257 != 0 {
		goto L1
	} else {
		goto L1193
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v31
	F_write_stderr(m, int32(_a_F_BootstrapModeMain_1), v29)
	mBase = m.M
	v6248 = m.ExcPending
	if v6248 != 0 {
		goto L1
	} else {
		goto L1191
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v130
	F_errmsg(m, int32(_a_F_BootstrapModeMain_2), v29+int32(32))
	mBase = m.M
	v6239 = m.ExcPending
	if v6239 != 0 {
		goto L1
	} else {
		goto L1189
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6217 = m.ExcPending
	if v6217 != 0 {
		goto L1
	} else {
		goto L1185
	}
L11:
	;
	v69 = F_getopt(m, v35, l1+int32(4), int32(_a_F_BootstrapModeMain_3))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L12:
	;
	if v69 != int32(-1) {
		goto L8
	} else {
		goto L91
	}
L13:
	;
	goto L12
L14:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_4), v309, int32(0), int32(1))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L90
	}
L15:
	;
	v185 = int32(_a_F_BootstrapModeMain_5)
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	goto L62
L16:
	;
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_6), int32(_a_F_BootstrapModeMain_7), int32(1), int32(4))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L58
	}
L17:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v160
	v166 = F_psprintf(m, int32(_a_F_BootstrapModeMain_8), v29+int32(80))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L54
	}
L18:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	v157 = F_pstrdup(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L53
	}
L19:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	F_ParseLongOption(m, v112, v29+int32(92), v29+int32(88))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L41
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	v83 = F_strcmp(m, int32(_a_F_BootstrapModeMain_9), v81)
	mBase = m.M
	if v83 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_10), v75, int32(1), int32(4))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	switch v69 - int32(45) {
	case 0:
		goto L20
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 24, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 56, 57, 58, 59, 60, 61, 63, 64, 65, 66, 67, 68:
		goto L8
	case 21:
		goto L21
	case 23:
		goto L18
	case 25:
		goto L16
	case 43:
		goto L14
	case 54:
		goto L19
	case 55:
		goto L17
	case 62:
		v42 = int32(1)
		goto L11
	case 69:
		goto L15
	default:
		goto L13
	}
L23:
	;
	goto L11
L24:
	;
	if v108 != int32(5) {
		goto L10
	} else {
		goto L40
	}
L25:
	;
	v108 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v88 = F_strcmp(m, int32(_a_F_BootstrapModeMain_11), v81)
	mBase = m.M
	if v88 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v108 = int32(1)
	goto L24
L29:
	;
	goto L30
L30:
	;
	v94 = F_strncmp(m, int32(_a_F_BootstrapModeMain_12), v81, int32(9))
	mBase = m.M
	if v94 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v108 = int32(2)
	goto L24
L32:
	;
	goto L33
L33:
	;
	v99 = F_strcmp(m, int32(_a_F_BootstrapModeMain_13), v81)
	mBase = m.M
	if v99 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v108 = int32(3)
	goto L24
L35:
	;
	goto L36
L36:
	;
	v106 = F_strcmp(m, int32(_a_F_BootstrapModeMain_14), v81)
	mBase = m.M
	if v106 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v107 = int32(5)
	goto L39
L38:
	;
	v107 = int32(4)
	goto L39
L39:
	;
	v108 = v107
	goto L24
L40:
	;
	goto L19
L41:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
	if v119 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	F_SetConfigOption(m, v144, v119, int32(1), int32(4))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L50
	}
L45:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	if v69 == int32(45) {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v130
	F_errmsg(m, int32(_a_F_BootstrapModeMain_15), v29+int32(48))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(259), int32(_a_F_BootstrapModeMain_17))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	F_pfree(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
	F_pfree(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L11
L53:
	;
	v44 = v157
	goto L11
L54:
	;
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_18), v166, int32(1), int32(4))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_19), v166, int32(1), int32(4))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_pfree(m, v166)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L11
L58:
	;
	goto L11
L59:
	;
	goto L11
L60:
	;
	v304 = F_strlen(m, v293)
	mBase = m.M
	goto L59
L62:
	;
	goto L63
L63:
	;
	v194 = int32(1023)
	if (v185^v187)&int32(3) != 0 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v297 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v294))) = uint8(v297)
	goto L60
L65:
	;
	v278 = v273
	v279 = v274
	v280 = v275
	goto L86
L66:
	;
	if v268 == int32(0) {
		v293 = v266
		v294 = v267
		goto L64
	} else {
		goto L85
	}
L67:
	;
	v266 = v187
	v267 = v185
	v268 = v194
	goto L66
L68:
	;
	goto L69
L69:
	;
	v198 = int32(0)
	if base.B2i32(v187&int32(3) == v198)|int32(0) == v198 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v234 == int32(0) {
		v293 = v231
		v294 = v232
		goto L64
	} else {
		goto L79
	}
L71:
	;
	v210 = v187
	v211 = v185
	v212 = v194
	goto L74
L72:
	;
	goto L73
L73:
	;
	v231 = v187
	v232 = v185
	v233 = v194
	v234 = int32(1)
	goto L70
L74:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	*(*uint8)(unsafe.Add(mBase, uint32(v211))) = uint8(v214)
	if v214 == int32(0) {
		v273 = v210
		v274 = v211
		v275 = v212
		goto L65
	} else {
		goto L76
	}
L75:
	;
	v231 = v225
	v232 = v219
	v233 = v221
	v234 = v223
	goto L70
L76:
	;
	v218 = int32(1)
	v219 = v211 + v218
	v221 = v212 - v218
	v222 = int32(0)
	v223 = base.B2i32(v221 != v222)
	v225 = v210 + v218
	if v225&int32(3) == v222 {
		v231 = v225
		v232 = v219
		v233 = v221
		v234 = v223
		goto L70
	} else {
		goto L77
	}
L77:
	;
	if v221 != 0 {
		v210 = v225
		v211 = v219
		v212 = v221
		goto L74
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if base.B2i32(v237 == int32(0))|base.B2i32(base.Ui32(v233) < base.Ui32(int32(4))) != 0 {
		v266 = v231
		v267 = v232
		v268 = v233
		goto L66
	} else {
		goto L80
	}
L80:
	;
	v244 = v231
	v245 = v232
	v246 = v233
	goto L81
L81:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	v252 = int32(-2139062144)
	if (int32(16843008)-v249|v249)&v252 != v252 {
		v273 = v244
		v274 = v245
		v275 = v246
		goto L65
	} else {
		goto L83
	}
L82:
	;
	v266 = v260
	v267 = v258
	v268 = v262
	goto L66
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v249
	v257 = int32(4)
	v258 = v245 + v257
	v260 = v244 + v257
	v262 = v246 - v257
	if base.Ui32(int32(3)) < base.Ui32(v262) {
		v244 = v260
		v245 = v258
		v246 = v262
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v273 = v266
	v274 = v267
	v275 = v268
	goto L65
L86:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	*(*uint8)(unsafe.Add(mBase, uint32(v279))) = uint8(v282)
	if v282 == int32(0) {
		v293 = v278
		v294 = v279
		goto L64
	} else {
		goto L88
	}
L87:
	;
	v293 = v289
	v294 = v287
	goto L64
L88:
	;
	v286 = int32(1)
	v287 = v279 + v286
	v289 = v278 + v286
	v291 = v280 - v286
	if v291 != 0 {
		v278 = v289
		v279 = v287
		v280 = v291
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	goto L11
L91:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[2]))
	if v35 != v317 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	v319 = F_SelectConfigFiles(m, v44, v31)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v319 == int32(0) {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	F_checkDataDir(m)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_ChangeToDataDir(m)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_CreateDataDirLockFile(m, int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v331 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[3])) = uint8(v331)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[0])) = int32(0)
	F_InitializeMaxBackends(m)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_InitPostmasterChildSlots(m)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v344 = int32(1)
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[4]))
	if v347&(v347-v344) != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	F_CreateSharedMemoryAndSemaphores(m)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L110
	}
L101:
	;
	v354 = v344 << (uint(int32(32)-base.I32_clz(v347)) % 32)
	goto L103
L102:
	;
	v354 = v347
	goto L103
L103:
	;
	if base.Ui32(v354) <= base.Ui32(int32(31)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v357 = int32(31)
	goto L106
L105:
	;
	v357 = v354
	goto L106
L106:
	;
	if base.Ui32(int32(_a_F_BootstrapModeMain_20)) <= base.Ui32(v354) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v362 = int32(1024)
	goto L109
L108:
	;
	v362 = int32(base.Ui32(v357) >> (uint(int32(4)) % 32))
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[5])) = v362
	goto L100
L110:
	;
	F_set_max_safe_fds(m)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	F_InitProcess(m)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_BaseInit(m)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v373 = int32(0)
	v375 = m.G0
	v377 = v375 - int32(32)
	m.G0 = v377
	switch int32(2) {
	case 0, 2:
		v387 = v373
		goto L116
	default:
		goto L117
	}
L115:
	;
	v405 = int32(2)
	v406 = int32(0)
	v408 = m.G0
	v410 = v408 - int32(32)
	m.G0 = v410
	switch v405 {
	case 0, 2:
		v420 = v406
		goto L122
	default:
		goto L123
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v377)+12)) = v387
	F_sigemptyset(m, v377+int32(16))
	mBase = m.M
	goto L119
L117:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[6])) = v373
	v387 = int32(_a_F_BootstrapModeMain_21)
	goto L116
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v377)+24)) = int32(268435456)
	v401 = F___sigaction(m, int32(1), v377+int32(12), int32(0))
	mBase = m.M
	m.G0 = v377 + int32(32)
	goto L115
L121:
	;
	v439 = int32(0)
	v441 = m.G0
	v443 = v441 - int32(32)
	m.G0 = v443
	switch int32(2) {
	case 0, 2:
		v453 = v439
		goto L128
	default:
		goto L129
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410)+12)) = v420
	F_sigemptyset(m, v410+int32(16))
	mBase = m.M
	goto L125
L123:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[7])) = v406
	v420 = int32(_a_F_BootstrapModeMain_21)
	goto L122
L125:
	;
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410)+24)) = int32(268435456)
	v434 = F___sigaction(m, v405, v410+int32(12), int32(0))
	mBase = m.M
	m.G0 = v410 + int32(32)
	goto L121
L127:
	;
	v472 = int32(0)
	v474 = m.G0
	v476 = v474 - int32(32)
	m.G0 = v476
	switch int32(2) {
	case 0, 2:
		v486 = v472
		goto L134
	default:
		goto L135
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v443)+12)) = v453
	F_sigemptyset(m, v443+int32(16))
	mBase = m.M
	goto L131
L129:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[8])) = v439
	v453 = int32(_a_F_BootstrapModeMain_21)
	goto L128
L131:
	;
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v443)+24)) = int32(268435456)
	v467 = F___sigaction(m, int32(15), v443+int32(12), int32(0))
	mBase = m.M
	m.G0 = v443 + int32(32)
	goto L127
L133:
	;
	v504 = m.G0
	v506 = v504 - int32(_a_F_BootstrapModeMain_22)
	m.G0 = v506
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[9]))
	v513 = F_LWLockAcquire(m, v509+int32(1152), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L139
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v476)+12)) = v486
	F_sigemptyset(m, v476+int32(16))
	mBase = m.M
	goto L137
L135:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[10])) = v472
	v486 = int32(_a_F_BootstrapModeMain_21)
	goto L134
L137:
	;
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v476)+24)) = int32(268435456)
	v500 = F___sigaction(m, int32(3), v476+int32(12), int32(0))
	mBase = m.M
	m.G0 = v476 + int32(32)
	goto L133
L139:
	;
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[11]))
	v517 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v516)+320)) = uint8(v517)
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[9]))
	F_LWLockRelease(m, v520+int32(1152))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_gettimeofday(m, v506-int32(-64))
	mBase = m.M
	v528 = int64(*(*int32)(unsafe.Add(mBase, uint32(v506)+72)))
	v529 = *(*int64)(unsafe.Add(mBase, uint32(v506)+64))
	v530 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	v532 = F_palloc(m, int32(_a_F_BootstrapModeMain_20))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v537 = (v532 + int32(_a_F_BootstrapModeMain_23)) & int32(-8192)
	v538 = int32(0)
	base.MemoryFill(m, v537, v538, int32(_a_F_BootstrapModeMain_24))
	v542 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[12]))
	v544 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[13])))
	v546 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[14]))
	v547 = F_time(m)
	mBase = m.M
	v548 = int32(_a_F_BootstrapModeMain_25)
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v549))) = int32(_a_F_BootstrapModeMain_26)
	*(*int64)(unsafe.Add(mBase, uint32(v549)+8)) = int64(3)
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v555)+4)) = v538
	F_MultiXactSetNextMXact(m, int32(1), v538)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_AdvanceOldestClogXid(m, int32(3))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_SetTransactionIdLimit(m, int32(3), int32(1))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v569 = int32(1)
	F_SetMultiXactIdLimit(m, v569, v569, v569)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v574 = int32(0)
	F_SetCommitTsLimit(m, v574, v574)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v537))) = int64(4295151896)
	v581 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[12]))
	v582 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v537)+48)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v537)+36)) = int32(_a_F_BootstrapModeMain_24)
	*(*int32)(unsafe.Add(mBase, uint32(v537)+32)) = v581
	v595 = base.I64_extend_i32_u(v530&int32(4095)) | (v528<<(uint(int64(12))%64) | v529<<(uint(int64(32))%64))
	*(*int64)(unsafe.Add(mBase, uint32(v537)+24)) = v595
	v597 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v537)+146)) = v597
	*(*int64)(unsafe.Add(mBase, uint32(v537)+138)) = v582
	*(*int64)(unsafe.Add(mBase, uint32(v537)+130)) = v547
	v602 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v537)+122)) = v602
	v604 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v537)+114)) = v604
	*(*int64)(unsafe.Add(mBase, uint32(v537)+106)) = int64(12884901888)
	*(*int64)(unsafe.Add(mBase, uint32(v537)+98)) = int64(4294977296)
	*(*int64)(unsafe.Add(mBase, uint32(v537)+90)) = int64(3)
	*(*int32)(unsafe.Add(mBase, uint32(v537)+86)) = v546
	*(*uint8)(unsafe.Add(mBase, uint32(v537)+82)) = uint8(v544)
	*(*int64)(unsafe.Add(mBase, uint32(v537)+74)) = v604
	v616 = int32(40)
	v618 = base.I64_extend_i32_u(v542 + v616)
	*(*int64)(unsafe.Add(mBase, uint32(v537)+66)) = v618
	v620 = int32(_a_F_BootstrapModeMain_27)
	*(*uint16)(unsafe.Add(mBase, uint32(v537)+64)) = uint16(v620)
	*(*uint16)(unsafe.Add(mBase, uint32(v537)+56)) = uint16(v597)
	*(*int64)(unsafe.Add(mBase, uint32(v537)+40)) = int64(114)
	*(*int64)(unsafe.Add(mBase, uint32(v537)+8)) = base.I64_extend_i32_s(v581)
	v628 = int32(-1)
	v632 = m.Env.Pgmem_crc32c(m, v628, v537|int32(64), int32(90))
	mBase = m.M
	v636 = m.Env.Pgmem_crc32c(m, v632, v537|v616, int32(20))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v537)+60)) = v636 ^ v628
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[16])) = v602
	v646 = F_XLogFileInit(m, int64(1), v602)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[17])) = v646
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = int32(0)
	v653 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v653))) = int32(167772229)
	v656 = int32(_a_F_BootstrapModeMain_24)
	v657 = F_write(m, v646, v537, v656)
	mBase = m.M
	if v657 != v656 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v661 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	if v661 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	v682 = int32(_a_F_BootstrapModeMain_28)
	v683 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	v684 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v683))) = v684
	v687 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v687))) = int32(167772228)
	v691 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[17]))
	v694 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[20])))
	if v694 != int32(1) {
		v708 = v684
		goto L165
	} else {
		goto L166
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = int32(51)
	goto L153
L152:
	;
	goto L153
L153:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_29), int32(0))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_30), int32(_a_F_BootstrapModeMain_31), int32(_a_F_BootstrapModeMain_32))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	v1118 = int32(0)
	F_InitPostgres(m, v1118, v1118, v1118, v1118, v1118, v1118)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L252
	}
L159:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L1
	} else {
		goto L248
	}
L160:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L1
	} else {
		goto L244
	}
L161:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L240
	}
L162:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L236
	}
L163:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L1
	} else {
		goto L232
	}
L164:
	;
	if v708 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L165:
	;
	goto L164
L166:
	;
	goto L167
L167:
	;
	v699 = F_fsync(m, v691)
	mBase = m.M
	if v699 != int32(-1) {
		v708 = v699
		goto L165
	} else {
		goto L169
	}
L168:
	;
	v708 = int32(-1)
	goto L165
L169:
	;
	v703 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	if v703 == int32(27) {
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v712 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v712))) = int32(0)
	v716 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[17]))
	v717 = F_close(m, v716)
	mBase = m.M
	if v717 != 0 {
		goto L163
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L228
	}
L174:
	;
	v719 = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[17])) = v719
	v722 = v506 + int32(80)
	v724 = int32(0)
	v728 = m.G0
	v730 = v728 - int32(16)
	m.G0 = v730
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = v724
	v736 = F_open(m, int32(_a_F_BootstrapModeMain_33), v724, v730)
	mBase = m.M
	if v736 != v719 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	if v769 == int32(0) {
		goto L162
	} else {
		goto L188
	}
L176:
	;
	goto L180
L177:
	;
	v769 = v724
	goto L178
L178:
	;
	m.G0 = v730 + int32(16)
	goto L175
L179:
	;
	v764 = F_close(m, v736)
	mBase = m.M
	v769 = v762
	goto L178
L180:
	;
	v742 = v722
	v743 = int32(32)
	goto L181
L181:
	;
	v748 = F_read(m, v736, v742, v743)
	mBase = m.M
	if v748 <= int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v762 = int32(1)
	goto L179
L183:
	;
	v752 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	if v752 == int32(27) {
		goto L181
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v757 = v743 - v748
	if v757 != 0 {
		v742 = v742 + v748
		v743 = v757
		goto L181
	} else {
		goto L187
	}
L186:
	;
	v762 = int32(0)
	goto L179
L187:
	;
	goto L182
L188:
	;
	v776 = int32(_a_F_BootstrapModeMain_34)
	v777 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[21]))
	v778 = int32(8)
	v780 = int32(0)
	base.MemoryFill(m, v777+v778, v780, int32(288))
	*(*int64)(unsafe.Add(mBase, uint32(v777))) = v595
	v784 = *(*int64)(unsafe.Add(mBase, uint32(v506)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v777)+257)) = v784
	v786 = *(*int64)(unsafe.Add(mBase, uint32(v506)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v777)+265)) = v786
	v788 = *(*int64)(unsafe.Add(mBase, uint32(v506)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v777)+273)) = v788
	v790 = *(*int64)(unsafe.Add(mBase, uint32(v506)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v777)+281)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v777)+128)) = int64(1000)
	v794 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v777)+16)) = v794
	v797 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v777)+180)) = v797
	v800 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v777)+184)) = v800
	v803 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[24]))
	*(*int32)(unsafe.Add(mBase, uint32(v777)+188)) = v803
	v806 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[25]))
	*(*int32)(unsafe.Add(mBase, uint32(v777)+192)) = v806
	v809 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v777)+196)) = v809
	v812 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v777)+172)) = v812
	v815 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[26])))
	*(*uint8)(unsafe.Add(mBase, uint32(v777)+176)) = uint8(v815)
	v818 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[27])))
	*(*int32)(unsafe.Add(mBase, uint32(v777)+252)) = v42
	*(*uint8)(unsafe.Add(mBase, uint32(v777)+200)) = uint8(v818)
	*(*int32)(unsafe.Add(mBase, uint32(v777)+120)) = v780
	*(*int64)(unsafe.Add(mBase, uint32(v777)+112)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v777)+104)) = v547
	*(*int32)(unsafe.Add(mBase, uint32(v777)+96)) = v794
	v828 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v777)+88)) = v828
	*(*int64)(unsafe.Add(mBase, uint32(v777)+80)) = int64(12884901888)
	*(*int64)(unsafe.Add(mBase, uint32(v777)+72)) = int64(4294977296)
	*(*int64)(unsafe.Add(mBase, uint32(v777)+64)) = int64(3)
	*(*int32)(unsafe.Add(mBase, uint32(v777)+60)) = v546
	*(*uint8)(unsafe.Add(mBase, uint32(v777)+56)) = uint8(v544)
	*(*int64)(unsafe.Add(mBase, uint32(v777)+48)) = v828
	*(*int64)(unsafe.Add(mBase, uint32(v777)+40)) = v618
	*(*int64)(unsafe.Add(mBase, uint32(v777)+32)) = v618
	*(*int64)(unsafe.Add(mBase, uint32(v777)+24)) = v547
	*(*int32)(unsafe.Add(mBase, uint32(v777)+224)) = int32(_a_F_BootstrapModeMain_24)
	*(*int64)(unsafe.Add(mBase, uint32(v777)+216)) = int64(562949953429504)
	*(*int64)(unsafe.Add(mBase, uint32(v777)+208)) = int64(4698053236609777664)
	*(*int32)(unsafe.Add(mBase, uint32(v777)+204)) = v778
	*(*int64)(unsafe.Add(mBase, uint32(v777)+8)) = int64(869757897079260936)
	v854 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[12]))
	v855 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v777)+292)) = v855
	*(*uint8)(unsafe.Add(mBase, uint32(v777)+256)) = uint8(v794)
	*(*uint8)(unsafe.Add(mBase, uint32(v777)+248)) = uint8(v780)
	*(*int64)(unsafe.Add(mBase, uint32(v777)+240)) = int64(8796093024204)
	*(*int64)(unsafe.Add(mBase, uint32(v777)+232)) = int64(137438953536)
	*(*int32)(unsafe.Add(mBase, uint32(v777)+228)) = v854
	v868 = m.Env.Pgmem_crc32c(m, v855, v777, int32(292))
	mBase = m.M
	v870 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v870)+292)) = v868 ^ v855
	base.MemoryFill(m, v506+int32(376), v780, int32(_a_F_BootstrapModeMain_35))
	base.MemoryCopy(m, v722, v870, int32(296))
	v883 = F_BasicOpenFile(m, int32(_a_F_BootstrapModeMain_36), int32(194))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	if v883 < int32(0) {
		goto L161
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = int32(0)
	v891 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = int32(167772172)
	v894 = int32(_a_F_BootstrapModeMain_24)
	v895 = F_write(m, v883, v722, v894)
	mBase = m.M
	if v895 != v894 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v899 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	if v899 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L192:
	;
	goto L193
L193:
	;
	v923 = int32(_a_F_BootstrapModeMain_28)
	v924 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	v925 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v924))) = v925
	v928 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v928))) = int32(167772170)
	v933 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[20])))
	if v933 != int32(1) {
		v947 = v925
		goto L202
	} else {
		goto L203
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = int32(51)
	goto L196
L195:
	;
	goto L196
L196:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+48)) = int32(_a_F_BootstrapModeMain_36)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_37), v506+int32(48))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_30), int32(_a_F_BootstrapModeMain_38), int32(_a_F_BootstrapModeMain_39))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L1
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
	if v947 != 0 {
		goto L160
	} else {
		goto L208
	}
L202:
	;
	goto L201
L203:
	;
	goto L204
L204:
	;
	v938 = F_fsync(m, v883)
	mBase = m.M
	if v938 != int32(-1) {
		v947 = v938
		goto L202
	} else {
		goto L206
	}
L205:
	;
	v947 = int32(-1)
	goto L202
L206:
	;
	v942 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	if v942 == int32(27) {
		goto L204
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	v949 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v949))) = int32(0)
	v952 = F_close(m, v883)
	mBase = m.M
	if v952 != 0 {
		goto L159
	} else {
		goto L209
	}
L209:
	;
	v954 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[28]))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v954)+28))
	v957 = F_LWLockAcquire(m, v955, int32(0))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v959 = int32(_a_F_BootstrapModeMain_40)
	v962 = F_SimpleLruZeroPage(m, v959, int64(0))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_SimpleLruWritePage(m, v959, v962)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_LWLockRelease(m, v955)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	v969 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[29]))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v969)+28))
	v972 = F_LWLockAcquire(m, v970, int32(0))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v974 = int32(_a_F_BootstrapModeMain_41)
	v977 = F_SimpleLruZeroPage(m, v974, int64(0))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_SimpleLruWritePage(m, v974, v977)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	F_LWLockRelease(m, v970)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v984 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30]))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v984)+28))
	v987 = F_LWLockAcquire(m, v985, int32(0))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	v989 = int32(_a_F_BootstrapModeMain_42)
	v992 = F_SimpleLruZeroPage(m, v989, int64(0))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	F_SimpleLruWritePage(m, v989, v992)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	F_LWLockRelease(m, v985)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v999 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[31]))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v999)+28))
	v1002 = F_LWLockAcquire(m, v1000, int32(0))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v1004 = int32(_a_F_BootstrapModeMain_43)
	v1007 = F_SimpleLruZeroPage(m, v1004, int64(0))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	F_SimpleLruWritePage(m, v1004, v1007)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	F_LWLockRelease(m, v1000)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	F_pfree(m, v532)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	F_ReadControlFile(m)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	m.G0 = v506 + int32(_a_F_BootstrapModeMain_22)
	goto L158
L228:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_44), int32(0))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_30), int32(_a_F_BootstrapModeMain_45), int32(_a_F_BootstrapModeMain_32))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_46), int32(0))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_30), int32(_a_F_BootstrapModeMain_47), int32(_a_F_BootstrapModeMain_32))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_48), int32(0))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_30), int32(_a_F_BootstrapModeMain_49), int32(_a_F_BootstrapModeMain_50))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506))) = int32(_a_F_BootstrapModeMain_36)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_51), v506)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_30), int32(_a_F_BootstrapModeMain_52), int32(_a_F_BootstrapModeMain_39))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+32)) = int32(_a_F_BootstrapModeMain_36)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_53), v506+int32(32))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_30), int32(_a_F_BootstrapModeMain_54), int32(_a_F_BootstrapModeMain_39))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+16)) = int32(_a_F_BootstrapModeMain_36)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_55), v506+int32(16))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_30), int32(_a_F_BootstrapModeMain_56), int32(_a_F_BootstrapModeMain_39))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	base.MemoryFill(m, int32(_a_F_BootstrapModeMain_57), int32(0), int32(160))
	v1131 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[32])) = v1131
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[33])) = v1131
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[34])) = v1131
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[35])) = v1131
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[36])) = v1131
	v1147 = F_boot_yylex_init(m, v29+int32(92))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	if v1147 != 0 {
		goto L4
	} else {
		goto L254
	}
L254:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	v1152 = m.G0
	v1154 = v1152 - int32(1136)
	m.G0 = v1154
	*(*int32)(unsafe.Add(mBase, uint32(v1154)+1132)) = int32(0)
	v1159 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[37]))
	v1162 = v1154 + int32(128)
	v1164 = v1154 + int32(928)
	v1169 = v1151
	v1171 = int32(-2)
	v1173 = v1162
	v1176 = v1154
	v1177 = v1164
	v1179 = v4
	v1180 = v1164
	v1182 = v1159
	v1183 = v1162
	v1184 = int32(200)
	v1188 = v4
	goto L263
L256:
	;
	if v6146+int32(928) != v6136 {
		goto L1171
	} else {
		goto L1172
	}
L257:
	;
	v6136 = v6121
	v6146 = v6120
	goto L256
L258:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6100 = m.ExcPending
	if v6100 != 0 {
		goto L1
	} else {
		goto L1168
	}
L259:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6087 = m.ExcPending
	if v6087 != 0 {
		goto L1
	} else {
		goto L1165
	}
L260:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6067 = m.ExcPending
	if v6067 != 0 {
		goto L1
	} else {
		goto L1162
	}
L261:
	;
	F_boot_yyerror(m, v1169, int32(_a_F_BootstrapModeMain_58))
	mBase = m.M
	v6063 = m.ExcPending
	if v6063 != 0 {
		goto L1
	} else {
		goto L1161
	}
L262:
	;
	v6044 = v6018
	goto L1158
L263:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1180))) = uint8(v1179)
	if base.Ui32(v1177+v1184-int32(1)) <= base.Ui32(v1180) {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	switch v3749 {
	case 0:
		goto L1155
	default:
		v6014 = v3737
		v6015 = v3738
		v6018 = v3741
		goto L262
	case 3:
		goto L1154
	}
L265:
	;
	if int32(_a_F_BootstrapModeMain_59) < v1184 {
		goto L261
	} else {
		goto L268
	}
L266:
	;
	v1243 = v1173
	v1244 = v1177
	v1245 = v1180
	v1246 = v1183
	v1247 = v1184
	goto L267
L267:
	;
	if v1179 == int32(46) {
		goto L285
	} else {
		goto L286
	}
L268:
	;
	v1199 = int32(_a_F_BootstrapModeMain_26)
	v1201 = v1184 << (uint(int32(1)) % 32)
	if v1199 <= v1201 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1204 = v1199
	goto L271
L270:
	;
	v1204 = v1201
	goto L271
L271:
	;
	v1209 = F_palloc(m, v1204*int32(5)+int32(3))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	if v1209 == int32(0) {
		goto L261
	} else {
		goto L273
	}
L273:
	;
	v1213 = v1180 - v1177
	v1215 = v1213 + int32(1)
	if v1215 != 0 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	base.MemoryCopy(m, v1209, v1177, v1215)
	goto L276
L275:
	;
	goto L276
L276:
	;
	v1220 = base.I32_div_s(v1204+int32(3), int32(4))
	v1221 = int32(2)
	v1223 = v1209 + v1220<<(uint(v1221)%32)
	v1225 = v1215 << (uint(v1221) % 32)
	if v1225 != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	base.MemoryCopy(m, v1223, v1183, v1225)
	goto L279
L278:
	;
	goto L279
L279:
	;
	if v1176+int32(928) != v1177 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	F_pfree(m, v1177)
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L1
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	if v1204-int32(1) <= v1213 {
		v6136 = v1209
		v6146 = v1176
		goto L256
	} else {
		goto L284
	}
L283:
	;
	goto L282
L284:
	;
	v1243 = v1223 + v1225 - int32(4)
	v1244 = v1209
	v1245 = v1209 + v1213
	v1246 = v1223
	v1247 = v1204
	goto L267
L285:
	;
	v6136 = v1244
	v6146 = v1176
	goto L256
L286:
	;
	goto L287
L287:
	;
	v1254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1179<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[38]))))
	if v1254 == int32(-53) {
		v3730 = v1169
		v3732 = v1171
		v3734 = v1243
		v3737 = v1176
		v3738 = v1244
		v3740 = v1179
		v3741 = v1245
		v3743 = v1182
		v3744 = v1246
		v3745 = v1247
		v3749 = v1188
		goto L291
	} else {
		goto L292
	}
L288:
	;
	goto L264
L289:
	;
	v1169 = v5974
	v1171 = v5976
	v1173 = v5978
	v1176 = v5981
	v1177 = v5982
	v1179 = v5984
	v1180 = v5985 + int32(1)
	v1182 = v5987
	v1183 = v5988
	v1184 = v5989
	v1188 = v5993
	goto L263
L290:
	;
	v3787 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3771)+uint32(_c_F_BootstrapModeMain[39]))))
	v3792 = *(*int32)(unsafe.Add(mBase, uint32(v3765+(int32(1)-v3787)<<(uint(int32(2))%32))))
	switch v3771 - int32(14) {
	case 0:
		goto L713
	case 1:
		goto L712
	case 2:
		goto L711
	case 3:
		goto L710
	case 4:
		goto L709
	case 5:
		goto L708
	case 6:
		goto L707
	case 7:
		goto L706
	case 8:
		goto L705
	case 9:
		goto L704
	case 10:
		goto L703
	case 11:
		goto L702
	case 12:
		goto L701
	case 13:
		goto L700
	case 14, 16, 25:
		goto L699
	case 15, 17, 19:
		goto L698
	case 18:
		goto L697
	default:
		v5920 = v3792
		goto L670
	case 22:
		goto L696
	case 23:
		goto L695
	case 24:
		goto L694
	case 26:
		goto L693
	case 30:
		goto L692
	case 31:
		goto L691
	case 32:
		goto L690
	case 33:
		goto L689
	case 34:
		goto L688
	case 35:
		goto L687
	case 36:
		goto L686
	case 37:
		goto L685
	case 38:
		goto L684
	case 39:
		goto L683
	case 40:
		goto L682
	case 41:
		goto L681
	case 42:
		goto L680
	case 43:
		goto L679
	case 44:
		goto L678
	case 45:
		goto L677
	case 46:
		goto L676
	case 47:
		goto L675
	case 48:
		goto L674
	case 49:
		goto L673
	case 50:
		goto L672
	case 51:
		goto L671
	}
L291:
	;
	v3755 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3740)+uint32(_c_F_BootstrapModeMain[40]))))
	if v3755 == int32(0) {
		goto L288
	} else {
		goto L669
	}
L292:
	;
	if v1171 == int32(-2) {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	v3709 = v1254 + v3708
	if base.Ui32(int32(169)) < base.Ui32(v3709) {
		v3730 = v3674
		v3732 = v3707
		v3734 = v3678
		v3737 = v3681
		v3738 = v3682
		v3740 = v3684
		v3741 = v3685
		v3743 = v3687
		v3744 = v3688
		v3745 = v3689
		v3749 = v3693
		goto L291
	} else {
		goto L664
	}
L294:
	;
	v1259 = m.G0
	v1261 = v1259 - int32(16)
	m.G0 = v1261
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+92)) = v1176 + int32(1132)
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+40))
	if v1266 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L295:
	;
	v3674 = v1169
	v3676 = v1171
	v3678 = v1243
	v3681 = v1176
	v3682 = v1244
	v3684 = v1179
	v3685 = v1245
	v3687 = v1182
	v3688 = v1246
	v3689 = v1247
	v3693 = v1188
	goto L296
L296:
	;
	if v3676 <= int32(0) {
		goto L659
	} else {
		goto L660
	}
L297:
	;
	v3674 = v1377
	v3676 = v2202
	v3678 = v1381
	v3681 = v1384
	v3682 = v1385
	v3684 = v1387
	v3685 = v1388
	v3687 = v1390
	v3688 = v1391
	v3689 = v1392
	v3693 = v1396
	goto L296
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+40)) = int32(1)
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+44))
	if v1271 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L299:
	;
	goto L300
L300:
	;
	v1339 = v1169
	v1343 = v1243
	v1346 = v1176
	v1347 = v1244
	v1349 = v1179
	v1350 = v1245
	v1352 = v1182
	v1353 = v1246
	v1354 = v1247
	v1355 = v1261
	v1358 = v1188
	goto L317
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+44)) = int32(1)
	goto L303
L302:
	;
	goto L303
L303:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+4))
	if v1276 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[41]))
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+4)) = v1280
	goto L306
L305:
	;
	goto L306
L306:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+8))
	if v1282 == int32(0) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[37]))
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+8)) = v1286
	goto L309
L308:
	;
	goto L309
L309:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+20))
	if v1288 != 0 {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1316)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+28)) = v1317
	v1321 = v1315 + v1314<<(uint(int32(2))%32)
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1321)))
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1322)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+80)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+36)) = v1323
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1321)))
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1326)))
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+4)) = v1327
	v1329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1323))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1169)+24)) = uint8(v1329)
	goto L300
L311:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+12))
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1288+v1289<<(uint(int32(2))%32))))
	if v1293 != 0 {
		v1314 = v1289
		v1315 = v1288
		v1316 = v1293
		goto L310
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	F_boot_yyensure_buffer_stack(m, v1169)
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L1
	} else {
		goto L315
	}
L314:
	;
	goto L313
L315:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+4))
	v1299 = F_boot_yy_create_buffer(m, v1298, v1169)
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+20))
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+12))
	v1303 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1301+v1302<<(uint(v1303)%32)))) = v1299
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+20))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+12))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1307+v1308<<(uint(v1303)%32))))
	v1314 = v1308
	v1315 = v1307
	v1316 = v1312
	goto L310
L317:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+36))
	v1363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1339)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1362))) = uint8(v1363)
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+20))
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+12))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1365+v1366<<(uint(int32(2))%32))))
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1370)+28))
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+44))
	v1374 = v1362
	v1377 = v1339
	v1378 = v1362
	v1379 = v1371 + v1372
	v1381 = v1343
	v1384 = v1346
	v1385 = v1347
	v1387 = v1349
	v1388 = v1350
	v1390 = v1352
	v1391 = v1353
	v1392 = v1354
	v1393 = v1355
	v1396 = v1358
	goto L319
L319:
	;
	v1400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1378))))
	v1401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1400)+uint32(_c_F_BootstrapModeMain[42]))))
	v1403 = v1379 << (uint(int32(1)) % 32)
	v1406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1403)+uint32(_c_F_BootstrapModeMain[43]))))
	if v1406 != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+68)) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+64)) = v1379
	goto L323
L322:
	;
	goto L323
L323:
	;
	v1411 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1403)+uint32(_c_F_BootstrapModeMain[44]))))
	v1412 = v1411 + v1401
	v1417 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1412<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[45]))))
	if v1417 != v1379 {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1421 = v1401
	v1424 = v1379
	v1425 = v1401
	goto L327
L325:
	;
	v1479 = v1412
	goto L326
L326:
	;
	v1496 = int32(1)
	v1502 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1479<<(uint(v1496)%32))+uint32(_c_F_BootstrapModeMain[46]))))
	if v1502 != int32(127) {
		v1378 = v1378 + v1496
		v1379 = v1502
		goto L319
	} else {
		goto L333
	}
L327:
	;
	v1449 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1424<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[47]))))
	if int32(128) <= v1449 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	v1479 = v1461
	goto L326
L329:
	;
	v1452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1425)+uint32(_c_F_BootstrapModeMain[48]))))
	v1453 = v1452
	goto L331
L330:
	;
	v1453 = v1421
	goto L331
L331:
	;
	v1455 = v1453 & int32(255)
	v1456 = int32(1)
	v1460 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1449<<(uint(v1456)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	v1461 = v1455 + v1460
	v1466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1461<<(uint(v1456)%32))+uint32(_c_F_BootstrapModeMain[45]))))
	if v1466 != v1449&int32(_a_F_BootstrapModeMain_60) {
		v1421 = v1453
		v1424 = v1449
		v1425 = v1455
		goto L327
	} else {
		goto L332
	}
L332:
	;
	goto L328
L333:
	;
	v1506 = v1374
	goto L334
L334:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+64))
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+68))
	v1534 = v1506
	v1538 = v1531
	v1541 = v1532
	goto L336
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+80)) = v1534
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+32)) = v1541 - v1534
	v1562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1541))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1377)+24)) = uint8(v1562)
	v1564 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1541))) = uint8(v1564)
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+36)) = v1541
	v1571 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1538<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[43]))))
	v1577 = v1571
	goto L338
L338:
	;
	switch v1577 {
	case 0:
		goto L384
	case 1:
		goto L383
	case 2:
		goto L382
	case 3:
		goto L381
	case 4:
		goto L380
	case 5:
		goto L379
	case 6:
		goto L378
	case 7:
		goto L377
	case 8:
		goto L376
	case 9:
		goto L375
	case 10:
		goto L374
	case 11:
		goto L373
	case 12:
		goto L372
	case 13:
		goto L371
	case 14:
		goto L370
	case 15:
		goto L369
	case 16:
		goto L368
	case 17:
		goto L367
	case 18:
		goto L366
	case 19:
		goto L365
	case 20:
		goto L364
	case 21:
		goto L363
	case 22:
		goto L362
	case 23:
		goto L361
	case 24:
		goto L360
	case 25:
		goto L359
	case 26:
		goto L358
	case 27:
		goto L357
	case 28:
		goto L356
	case 29:
		goto L355
	case 30:
		goto L352
	case 31:
		goto L351
	case 32:
		goto L350
	case 33:
		v2202 = int32(0)
		goto L353
	default:
		goto L349
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+36)) = v3635
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+48)) = int32(0)
	v3664 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+44))
	v3668 = base.I32_div_s(v3664-int32(1), int32(2))
	v1577 = v3668 + int32(33)
	goto L338
L341:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_61))
	mBase = m.M
	v3634 = m.ExcPending
	if v3634 != 0 {
		goto L1
	} else {
		goto L658
	}
L342:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_62))
	mBase = m.M
	v3628 = m.ExcPending
	if v3628 != 0 {
		goto L1
	} else {
		goto L657
	}
L343:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L344:
	;
	v3480 = v3458 + v3459
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+36)) = v3480
	v3485 = *(*int32)(unsafe.Add(mBase, uint32(v3460+v3463<<(uint(int32(2))%32))))
	v3486 = *(*int32)(unsafe.Add(mBase, uint32(v3485)+28))
	v3487 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+44))
	v3488 = v3486 + v3487
	if base.Ui32(v3480) <= base.Ui32(v3454) {
		v1534 = v3454
		v1538 = v3488
		v1541 = v3480
		goto L336
	} else {
		goto L638
	}
L345:
	;
	v3102 = *(*int32)(unsafe.Add(mBase, uint32(v3101)))
	*(*int32)(unsafe.Add(mBase, uint32(v3102)+16)) = v3079
	v3105 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+28))
	if v3105 != 0 {
		v3227 = int32(0)
		goto L582
	} else {
		goto L583
	}
L346:
	;
	v3070 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v3071 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3079 = v3048
	v3101 = v3070 + v3071<<(uint(int32(2))%32)
	goto L345
L347:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_63))
	mBase = m.M
	v3043 = m.ExcPending
	if v3043 != 0 {
		goto L1
	} else {
		goto L581
	}
L348:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_64))
	mBase = m.M
	v3040 = m.ExcPending
	if v3040 != 0 {
		goto L1
	} else {
		goto L580
	}
L349:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_65))
	mBase = m.M
	v3037 = m.ExcPending
	if v3037 != 0 {
		goto L1
	} else {
		goto L579
	}
L350:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1377)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1541))) = uint8(v2266)
	v2268 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2269 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2272 = v2268 + v2269<<(uint(int32(2))%32)
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v2272)))
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(v2273)+44))
	if v2274 == int32(0) {
		goto L470
	} else {
		goto L471
	}
L351:
	;
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v2245 {
		goto L466
	} else {
		goto L467
	}
L352:
	;
	v2206 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v2206 {
		goto L460
	} else {
		goto L461
	}
L353:
	;
	m.G0 = v1393 + int32(16)
	goto L297
L354:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2197))) = v2195
	v2202 = int32(258)
	goto L353
L355:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v2175 {
		goto L456
	} else {
		goto L457
	}
L356:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v2154 {
		goto L452
	} else {
		goto L453
	}
L357:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v2132 {
		goto L449
	} else {
		goto L450
	}
L358:
	;
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v2111 {
		goto L446
	} else {
		goto L447
	}
L359:
	;
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v2090 {
		goto L443
	} else {
		goto L444
	}
L360:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v2069 {
		goto L440
	} else {
		goto L441
	}
L361:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v2048 {
		goto L437
	} else {
		goto L438
	}
L362:
	;
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v2027 {
		goto L434
	} else {
		goto L435
	}
L363:
	;
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v2006 {
		goto L431
	} else {
		goto L432
	}
L364:
	;
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v1985 {
		goto L428
	} else {
		goto L429
	}
L365:
	;
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v1964 {
		goto L425
	} else {
		goto L426
	}
L366:
	;
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v1943 {
		goto L422
	} else {
		goto L423
	}
L367:
	;
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v1922 {
		goto L419
	} else {
		goto L420
	}
L368:
	;
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if v1905 <= int32(0) {
		v1339 = v1377
		v1343 = v1381
		v1346 = v1384
		v1347 = v1385
		v1349 = v1387
		v1350 = v1388
		v1352 = v1390
		v1353 = v1391
		v1354 = v1392
		v1355 = v1393
		v1358 = v1396
		goto L317
	} else {
		goto L418
	}
L369:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if v1888 <= int32(0) {
		v1339 = v1377
		v1343 = v1381
		v1346 = v1384
		v1347 = v1385
		v1349 = v1387
		v1350 = v1388
		v1352 = v1390
		v1353 = v1391
		v1354 = v1392
		v1355 = v1393
		v1358 = v1396
		goto L317
	} else {
		goto L417
	}
L370:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1861 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v1861 {
		goto L414
	} else {
		goto L415
	}
L371:
	;
	v1841 = int32(262)
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if v1842 <= int32(0) {
		v2202 = v1841
		goto L353
	} else {
		goto L413
	}
L372:
	;
	v1823 = int32(261)
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if v1824 <= int32(0) {
		v2202 = v1823
		goto L353
	} else {
		goto L412
	}
L373:
	;
	v1805 = int32(260)
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if v1806 <= int32(0) {
		v2202 = v1805
		goto L353
	} else {
		goto L411
	}
L374:
	;
	v1787 = int32(259)
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if v1788 <= int32(0) {
		v2202 = v1787
		goto L353
	} else {
		goto L410
	}
L375:
	;
	v1769 = int32(263)
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if v1770 <= int32(0) {
		v2202 = v1769
		goto L353
	} else {
		goto L409
	}
L376:
	;
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v1748 {
		goto L406
	} else {
		goto L407
	}
L377:
	;
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v1727 {
		goto L403
	} else {
		goto L404
	}
L378:
	;
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v1706 {
		goto L400
	} else {
		goto L401
	}
L379:
	;
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v1685 {
		goto L397
	} else {
		goto L398
	}
L380:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v1664 {
		goto L394
	} else {
		goto L395
	}
L381:
	;
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v1643 {
		goto L391
	} else {
		goto L392
	}
L382:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v1622 {
		goto L388
	} else {
		goto L389
	}
L383:
	;
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+32))
	if int32(0) < v1601 {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	v1599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1377)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1541))) = uint8(v1599)
	v1506 = v1534
	goto L334
L385:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1604+v1605<<(uint(int32(2))%32))))
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1610+v1601-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1609)+28)) = base.B2i32(v1614 == int32(10))
	goto L387
L386:
	;
	goto L387
L387:
	;
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1618))) = int32(_a_F_BootstrapModeMain_66)
	v2202 = int32(264)
	goto L353
L388:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1625+v1626<<(uint(int32(2))%32))))
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1631+v1622-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1630)+28)) = base.B2i32(v1635 == int32(10))
	goto L390
L389:
	;
	goto L390
L390:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1639))) = int32(_a_F_BootstrapModeMain_67)
	v2202 = int32(265)
	goto L353
L391:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1646+v1647<<(uint(int32(2))%32))))
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1652+v1643-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1651)+28)) = base.B2i32(v1656 == int32(10))
	goto L393
L392:
	;
	goto L393
L393:
	;
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1660))) = int32(_a_F_BootstrapModeMain_68)
	v2202 = int32(266)
	goto L353
L394:
	;
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1667+v1668<<(uint(int32(2))%32))))
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1673+v1664-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1672)+28)) = base.B2i32(v1677 == int32(10))
	goto L396
L395:
	;
	goto L396
L396:
	;
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1681))) = int32(_a_F_BootstrapModeMain_69)
	v2202 = int32(276)
	goto L353
L397:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1688+v1689<<(uint(int32(2))%32))))
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1694+v1685-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1693)+28)) = base.B2i32(v1698 == int32(10))
	goto L399
L398:
	;
	goto L399
L399:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1702))) = int32(_a_F_BootstrapModeMain_70)
	v2202 = int32(277)
	goto L353
L400:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1709+v1710<<(uint(int32(2))%32))))
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1715+v1706-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1714)+28)) = base.B2i32(v1719 == int32(10))
	goto L402
L401:
	;
	goto L402
L402:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1723))) = int32(_a_F_BootstrapModeMain_71)
	v2202 = int32(278)
	goto L353
L403:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v1730+v1731<<(uint(int32(2))%32))))
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1736+v1727-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1735)+28)) = base.B2i32(v1740 == int32(10))
	goto L405
L404:
	;
	goto L405
L405:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1744))) = int32(_a_F_BootstrapModeMain_72)
	v2202 = int32(279)
	goto L353
L406:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1751+v1752<<(uint(int32(2))%32))))
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1757+v1748-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+28)) = base.B2i32(v1761 == int32(10))
	goto L408
L407:
	;
	goto L408
L408:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1765))) = int32(_a_F_BootstrapModeMain_73)
	v2202 = int32(267)
	goto L353
L409:
	;
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v1773+v1774<<(uint(int32(2))%32))))
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1779+v1770-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1778)+28)) = base.B2i32(v1783 == int32(10))
	v2202 = v1769
	goto L353
L410:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1791+v1792<<(uint(int32(2))%32))))
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1797+v1788-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1796)+28)) = base.B2i32(v1801 == int32(10))
	v2202 = v1787
	goto L353
L411:
	;
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v1809+v1810<<(uint(int32(2))%32))))
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1815+v1806-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1814)+28)) = base.B2i32(v1819 == int32(10))
	v2202 = v1805
	goto L353
L412:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1827+v1828<<(uint(int32(2))%32))))
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1833+v1824-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1832)+28)) = base.B2i32(v1837 == int32(10))
	v2202 = v1823
	goto L353
L413:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v1845+v1846<<(uint(int32(2))%32))))
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1851+v1842-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1850)+28)) = base.B2i32(v1855 == int32(10))
	v2202 = v1841
	goto L353
L414:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v1860+v1859<<(uint(int32(2))%32))))
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1868+v1861-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1867)+28)) = base.B2i32(v1872 == int32(10))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1878 = v1876
	v1879 = v1877
	goto L416
L415:
	;
	v1878 = v1860
	v1879 = v1859
	goto L416
L416:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1879<<(uint(int32(2))%32)+v1878)))
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1883)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1883)+32)) = v1884 + int32(1)
	v1339 = v1377
	v1343 = v1381
	v1346 = v1384
	v1347 = v1385
	v1349 = v1387
	v1350 = v1388
	v1352 = v1390
	v1353 = v1391
	v1354 = v1392
	v1355 = v1393
	v1358 = v1396
	goto L317
L417:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1891+v1892<<(uint(int32(2))%32))))
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1897+v1888-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1896)+28)) = base.B2i32(v1901 == int32(10))
	v1339 = v1377
	v1343 = v1381
	v1346 = v1384
	v1347 = v1385
	v1349 = v1387
	v1350 = v1388
	v1352 = v1390
	v1353 = v1391
	v1354 = v1392
	v1355 = v1393
	v1358 = v1396
	goto L317
L418:
	;
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v1908+v1909<<(uint(int32(2))%32))))
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1914+v1905-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1913)+28)) = base.B2i32(v1918 == int32(10))
	v1339 = v1377
	v1343 = v1381
	v1346 = v1384
	v1347 = v1385
	v1349 = v1387
	v1350 = v1388
	v1352 = v1390
	v1353 = v1391
	v1354 = v1392
	v1355 = v1393
	v1358 = v1396
	goto L317
L419:
	;
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1925+v1926<<(uint(int32(2))%32))))
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1931+v1922-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1930)+28)) = base.B2i32(v1935 == int32(10))
	goto L421
L420:
	;
	goto L421
L421:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1939))) = int32(_a_F_BootstrapModeMain_74)
	v2202 = int32(268)
	goto L353
L422:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1946+v1947<<(uint(int32(2))%32))))
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1952+v1943-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1951)+28)) = base.B2i32(v1956 == int32(10))
	goto L424
L423:
	;
	goto L424
L424:
	;
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1960))) = int32(_a_F_BootstrapModeMain_75)
	v2202 = int32(272)
	goto L353
L425:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v1967+v1968<<(uint(int32(2))%32))))
	v1973 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973+v1964-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1972)+28)) = base.B2i32(v1977 == int32(10))
	goto L427
L426:
	;
	goto L427
L427:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1981))) = int32(_a_F_BootstrapModeMain_76)
	v2202 = int32(273)
	goto L353
L428:
	;
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1988+v1989<<(uint(int32(2))%32))))
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v1998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1994+v1985-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1993)+28)) = base.B2i32(v1998 == int32(10))
	goto L430
L429:
	;
	goto L430
L430:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2002))) = int32(_a_F_BootstrapModeMain_77)
	v2202 = int32(274)
	goto L353
L431:
	;
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v2009+v2010<<(uint(int32(2))%32))))
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2015+v2006-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2014)+28)) = base.B2i32(v2019 == int32(10))
	goto L433
L432:
	;
	goto L433
L433:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2023))) = int32(_a_F_BootstrapModeMain_78)
	v2202 = int32(269)
	goto L353
L434:
	;
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v2030+v2031<<(uint(int32(2))%32))))
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2036+v2027-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2035)+28)) = base.B2i32(v2040 == int32(10))
	goto L436
L435:
	;
	goto L436
L436:
	;
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2044))) = int32(_a_F_BootstrapModeMain_79)
	v2202 = int32(270)
	goto L353
L437:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v2051+v2052<<(uint(int32(2))%32))))
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2057+v2048-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2056)+28)) = base.B2i32(v2061 == int32(10))
	goto L439
L438:
	;
	goto L439
L439:
	;
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2065))) = int32(_a_F_BootstrapModeMain_80)
	v2202 = int32(271)
	goto L353
L440:
	;
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v2072+v2073<<(uint(int32(2))%32))))
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2078+v2069-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2077)+28)) = base.B2i32(v2082 == int32(10))
	goto L442
L441:
	;
	goto L442
L442:
	;
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2086))) = int32(_a_F_BootstrapModeMain_81)
	v2202 = int32(275)
	goto L353
L443:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v2093+v2094<<(uint(int32(2))%32))))
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2099+v2090-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2098)+28)) = base.B2i32(v2103 == int32(10))
	goto L445
L444:
	;
	goto L445
L445:
	;
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2107))) = int32(_a_F_BootstrapModeMain_82)
	v2202 = int32(280)
	goto L353
L446:
	;
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v2114+v2115<<(uint(int32(2))%32))))
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2120+v2111-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2119)+28)) = base.B2i32(v2124 == int32(10))
	goto L448
L447:
	;
	goto L448
L448:
	;
	v2128 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2128))) = int32(_a_F_BootstrapModeMain_83)
	v2202 = int32(281)
	goto L353
L449:
	;
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v2135+v2136<<(uint(int32(2))%32))))
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2141+v2132-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2140)+28)) = base.B2i32(v2145 == int32(10))
	goto L451
L450:
	;
	goto L451
L451:
	;
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2149))) = int32(_a_F_BootstrapModeMain_84)
	v2202 = int32(282)
	goto L353
L452:
	;
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2158 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v2157+v2158<<(uint(int32(2))%32))))
	v2166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2153+v2154-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2162)+28)) = base.B2i32(v2166 == int32(10))
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2171 = v2170
	goto L454
L453:
	;
	v2171 = v2153
	goto L454
L454:
	;
	v2172 = F_pstrdup(m, v2171)
	mBase = m.M
	v2173 = m.ExcPending
	if v2173 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	v2195 = v2172
	goto L354
L456:
	;
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v2178+v2179<<(uint(int32(2))%32))))
	v2187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2174+v2175-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2183)+28)) = base.B2i32(v2187 == int32(10))
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2192 = v2191
	goto L458
L457:
	;
	v2192 = v2174
	goto L458
L458:
	;
	v2193 = F_DeescapeQuotedString(m, v2192)
	mBase = m.M
	v2194 = m.ExcPending
	if v2194 != 0 {
		goto L1
	} else {
		goto L459
	}
L459:
	;
	v2195 = v2193
	goto L354
L460:
	;
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(v2209+v2210<<(uint(int32(2))%32))))
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2215+v2206-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2214)+28)) = base.B2i32(v2219 == int32(10))
	goto L462
L461:
	;
	goto L462
L462:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2226 = m.ExcPending
	if v2226 != 0 {
		goto L1
	} else {
		goto L463
	}
L463:
	;
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(v2227+v2228<<(uint(int32(2))%32))))
	v2233 = *(*int32)(unsafe.Add(mBase, uint32(v2232)+32))
	v2234 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1393)+4)) = v2234
	*(*int32)(unsafe.Add(mBase, uint32(v1393))) = v2233
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_85), v1393)
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_86), int32(124), int32(_a_F_BootstrapModeMain_87))
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L466:
	;
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2253 = *(*int32)(unsafe.Add(mBase, uint32(v2248+v2249<<(uint(int32(2))%32))))
	v2254 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2254+v2245-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2253)+28)) = base.B2i32(v2258 == int32(10))
	goto L468
L467:
	;
	goto L468
L468:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_88))
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L470:
	;
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(v2273)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+28)) = v2277
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(v2272)))
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2279))) = v2280
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2284 = int32(2)
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(v2282+v2283<<(uint(v2284)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2287)+44)) = int32(1)
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2291 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(v2290+v2291<<(uint(v2284)%32))))
	v2296 = v2295
	v2297 = v2290
	v2298 = v2291
	goto L472
L471:
	;
	v2296 = v2273
	v2297 = v2268
	v2298 = v2269
	goto L472
L472:
	;
	v2299 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+36))
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+4))
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+28))
	v2302 = v2300 + v2301
	if base.Ui32(v2299) <= base.Ui32(v2302) {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2307 = v2265 ^ int32(-1) + v1541
	v2308 = v2304 + v2307
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+36)) = v2308
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(v2297+v2298<<(uint(int32(2))%32))))
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v2313)+28))
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+44))
	v2316 = v2314 + v2315
	if int32(0) < v2307 {
		goto L476
	} else {
		goto L477
	}
L474:
	;
	goto L475
L475:
	;
	if base.Ui32(v2302+int32(1)) < base.Ui32(v2299) {
		goto L348
	} else {
		goto L508
	}
L476:
	;
	v2323 = v2304
	v2324 = v2316
	goto L479
L477:
	;
	v2458 = v2316
	goto L478
L478:
	;
	v2483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2458<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[43]))))
	if v2483 != 0 {
		goto L497
	} else {
		goto L498
	}
L479:
	;
	v2346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2323))))
	if v2346 != 0 {
		goto L481
	} else {
		goto L482
	}
L480:
	;
	v2458 = v2449
	goto L478
L481:
	;
	v2347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2346)+uint32(_c_F_BootstrapModeMain[42]))))
	v2348 = v2347
	goto L483
L482:
	;
	v2348 = int32(1)
	goto L483
L483:
	;
	v2350 = v2324 << (uint(int32(1)) % 32)
	v2353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2350)+uint32(_c_F_BootstrapModeMain[43]))))
	if v2353 != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+68)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+64)) = v2324
	goto L486
L485:
	;
	goto L486
L486:
	;
	v2357 = v2348 & int32(255)
	v2360 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2350)+uint32(_c_F_BootstrapModeMain[44]))))
	v2361 = v2357 + v2360
	v2366 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2361<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[45]))))
	if v2366 != v2324 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v2370 = v2348
	v2373 = v2324
	v2374 = v2357
	goto L490
L488:
	;
	v2428 = v2361
	goto L489
L489:
	;
	v2445 = int32(1)
	v2449 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2428<<(uint(v2445)%32))+uint32(_c_F_BootstrapModeMain[46]))))
	v2451 = v2323 + v2445
	if v2451 != v2308 {
		v2323 = v2451
		v2324 = v2449
		goto L479
	} else {
		goto L496
	}
L490:
	;
	v2398 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2373<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[47]))))
	if int32(128) <= v2398 {
		goto L492
	} else {
		goto L493
	}
L491:
	;
	v2428 = v2410
	goto L489
L492:
	;
	v2401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2374)+uint32(_c_F_BootstrapModeMain[48]))))
	v2402 = v2401
	goto L494
L493:
	;
	v2402 = v2370
	goto L494
L494:
	;
	v2404 = v2402 & int32(255)
	v2405 = int32(1)
	v2409 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2398<<(uint(v2405)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	v2410 = v2404 + v2409
	v2415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2410<<(uint(v2405)%32))+uint32(_c_F_BootstrapModeMain[45]))))
	if v2415 != v2398&int32(_a_F_BootstrapModeMain_60) {
		v2370 = v2402
		v2373 = v2398
		v2374 = v2404
		goto L490
	} else {
		goto L495
	}
L495:
	;
	goto L491
L496:
	;
	goto L480
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+68)) = v2308
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+64)) = v2458
	goto L499
L498:
	;
	goto L499
L499:
	;
	v2486 = int32(1)
	v2490 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2458<<(uint(v2486)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	v2492 = v2490 + v2486
	v2497 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2492<<(uint(v2486)%32))+uint32(_c_F_BootstrapModeMain[45]))))
	if v2497 != v2458 {
		goto L500
	} else {
		goto L501
	}
L500:
	;
	v2504 = v2458
	goto L503
L501:
	;
	v2546 = v2492
	goto L502
L502:
	;
	if v2546 == int32(0) {
		v1506 = v2304
		goto L334
	} else {
		goto L506
	}
L503:
	;
	v2525 = int32(1)
	v2529 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2504<<(uint(v2525)%32))+uint32(_c_F_BootstrapModeMain[47]))))
	v2530 = base.I32_extend16_s(v2529)
	v2535 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2530<<(uint(v2525)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	v2537 = v2535 + v2525
	v2542 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2537<<(uint(v2525)%32))+uint32(_c_F_BootstrapModeMain[45]))))
	if v2529 != v2542 {
		v2504 = v2530
		goto L503
	} else {
		goto L505
	}
L504:
	;
	v2546 = v2537
	goto L502
L505:
	;
	goto L504
L506:
	;
	v2576 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2546<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[46]))))
	if v2576 == int32(127) {
		v1506 = v2304
		goto L334
	} else {
		goto L507
	}
L507:
	;
	v2580 = v2308 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+36)) = v2580
	v1374 = v2304
	v1378 = v2580
	v1379 = v2576
	goto L319
L508:
	;
	v2585 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+80))
	v2586 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+40))
	if v2586 == int32(0) {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	if v2299-v2585 != int32(1) {
		v3454 = v2585
		v3458 = v2301
		v3459 = v2300
		v3460 = v2297
		v3463 = v2298
		goto L344
	} else {
		goto L512
	}
L510:
	;
	goto L511
L511:
	;
	v2594 = v2585 ^ int32(-1) + v2299
	if int32(0) < v2594 {
		goto L513
	} else {
		goto L514
	}
L512:
	;
	v3635 = v2585
	goto L340
L513:
	;
	v2597 = int32(7)
	v2598 = v2594 & v2597
	if base.Ui32(v2299-v2585-int32(2)) < base.Ui32(v2597) {
		goto L518
	} else {
		goto L519
	}
L514:
	;
	v2754 = v2296
	v2758 = v2297
	v2761 = v2298
	goto L515
L515:
	;
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(v2754)+44))
	if v2778 == int32(2) {
		goto L528
	} else {
		goto L529
	}
L516:
	;
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2747 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2746+v2747<<(uint(int32(2))%32))))
	v2754 = v2751
	v2758 = v2746
	v2761 = v2747
	goto L515
L517:
	;
	v2687 = v2660
	v2690 = v2663
	v2691 = int32(0)
	goto L525
L518:
	;
	v2660 = v2585
	v2663 = v2300
	goto L517
L519:
	;
	goto L520
L520:
	;
	v2609 = v2585
	v2612 = v2300
	v2613 = int32(0)
	goto L521
L521:
	;
	v2633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2609))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2612))) = uint8(v2633)
	v2635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2609)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2612)+1)) = uint8(v2635)
	v2637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2609)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2612)+2)) = uint8(v2637)
	v2639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2609)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2612)+3)) = uint8(v2639)
	v2641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2609)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2612)+4)) = uint8(v2641)
	v2643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2609)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2612)+5)) = uint8(v2643)
	v2645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2609)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2612)+6)) = uint8(v2645)
	v2647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2609)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2612)+7)) = uint8(v2647)
	v2649 = int32(8)
	v2650 = v2612 + v2649
	v2652 = v2609 + v2649
	v2654 = v2613 + v2649
	if v2654 != v2594&int32(2147483640) {
		v2609 = v2652
		v2612 = v2650
		v2613 = v2654
		goto L521
	} else {
		goto L523
	}
L522:
	;
	if v2598 == int32(0) {
		goto L516
	} else {
		goto L524
	}
L523:
	;
	goto L522
L524:
	;
	v2660 = v2652
	v2663 = v2650
	goto L517
L525:
	;
	v2711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2687))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2690))) = uint8(v2711)
	v2713 = int32(1)
	v2718 = v2691 + v2713
	if v2718 != v2598 {
		v2687 = v2687 + v2713
		v2690 = v2690 + v2713
		v2691 = v2718
		goto L525
	} else {
		goto L527
	}
L526:
	;
	goto L516
L527:
	;
	goto L526
L528:
	;
	v2781 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+28)) = v2781
	v3079 = v2781
	v3101 = v2758 + v2761<<(uint(int32(2))%32)
	goto L345
L529:
	;
	goto L530
L530:
	;
	v2787 = int32(0)
	v2788 = *(*int32)(unsafe.Add(mBase, uint32(v2754)+12))
	v2789 = v2585 - v2299
	v2790 = v2788 + v2789
	if v2790 <= v2787 {
		goto L531
	} else {
		goto L532
	}
L531:
	;
	v2793 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+36))
	v2794 = v2793
	v2796 = v2754
	v2803 = v2788
	goto L534
L532:
	;
	v2860 = v2754
	v2863 = v2790
	goto L533
L533:
	;
	v2884 = int32(_a_F_BootstrapModeMain_24)
	if base.Ui32(v2884) <= base.Ui32(v2863) {
		goto L550
	} else {
		goto L551
	}
L534:
	;
	v2820 = *(*int32)(unsafe.Add(mBase, uint32(v2796)+20))
	if v2820 == int32(0) {
		goto L536
	} else {
		goto L537
	}
L535:
	;
	v2860 = v2853
	v2863 = v2855
	goto L533
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2796)+4)) = int32(0)
	goto L341
L537:
	;
	goto L538
L538:
	;
	v2825 = *(*int32)(unsafe.Add(mBase, uint32(v2796)+4))
	v2827 = v2803 << (uint(int32(1)) % 32)
	if v2827 <= int32(0) {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	v2831 = base.I32_div_s(v2803, int32(8))
	v2833 = v2831 + v2803
	goto L541
L540:
	;
	v2833 = v2827
	goto L541
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2796)+12)) = v2833
	v2836 = v2833 + int32(2)
	if v2825 != 0 {
		goto L543
	} else {
		goto L544
	}
L542:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2796)+4)) = v2841
	if v2841 == int32(0) {
		goto L341
	} else {
		goto L548
	}
L543:
	;
	v2837 = F_repalloc(m, v2825, v2836)
	mBase = m.M
	v2838 = m.ExcPending
	if v2838 != 0 {
		goto L1
	} else {
		goto L546
	}
L544:
	;
	goto L545
L545:
	;
	v2839 = F_palloc(m, v2836)
	mBase = m.M
	v2840 = m.ExcPending
	if v2840 != 0 {
		goto L1
	} else {
		goto L547
	}
L546:
	;
	v2841 = v2837
	goto L542
L547:
	;
	v2841 = v2839
	goto L542
L548:
	;
	v2846 = v2841 + (v2794 - v2825)
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+36)) = v2846
	v2848 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2849 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2853 = *(*int32)(unsafe.Add(mBase, uint32(v2848+v2849<<(uint(int32(2))%32))))
	v2854 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+12))
	v2855 = v2854 + v2789
	if v2855 <= int32(0) {
		v2794 = v2846
		v2796 = v2853
		v2803 = v2854
		goto L534
	} else {
		goto L549
	}
L549:
	;
	goto L535
L550:
	;
	v2887 = v2884
	goto L552
L551:
	;
	v2887 = v2863
	goto L552
L552:
	;
	v2888 = *(*int32)(unsafe.Add(mBase, uint32(v2860)+24))
	if v2888 != 0 {
		goto L553
	} else {
		goto L554
	}
L553:
	;
	v2893 = v2787
	goto L557
L554:
	;
	goto L555
L555:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = int32(0)
	v2963 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2964 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2968 = *(*int32)(unsafe.Add(mBase, uint32(v2963+v2964<<(uint(int32(2))%32))))
	v2969 = *(*int32)(unsafe.Add(mBase, uint32(v2968)+4))
	v2972 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+4))
	v2973 = F_fread(m, v2969+v2594, int32(1), v2887, v2972)
	mBase = m.M
	v2974 = m.ExcPending
	if v2974 != 0 {
		goto L1
	} else {
		goto L568
	}
L556:
	;
	switch v2919 {
	case 0:
		goto L564
	default:
		v2958 = v2933
		goto L562
	case 11:
		goto L563
	}
L557:
	;
	v2915 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+4))
	v2916 = F_do_getc(m, v2915)
	mBase = m.M
	v2917 = m.ExcPending
	if v2917 != 0 {
		goto L1
	} else {
		goto L560
	}
L558:
	;
	v2933 = v2887
	goto L556
L559:
	;
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v2920+v2921<<(uint(int32(2))%32))))
	v2926 = *(*int32)(unsafe.Add(mBase, uint32(v2925)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v2926+v2594+v2893))) = uint8(v2916)
	v2931 = v2893 + int32(1)
	if v2931 != v2887 {
		v2893 = v2931
		goto L557
	} else {
		goto L561
	}
L560:
	;
	v2919 = v2916 + int32(1)
	switch v2919 {
	case 0, 11:
		v2933 = v2893
		goto L556
	default:
		goto L559
	}
L561:
	;
	goto L558
L562:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+28)) = v2958
	v3048 = v2958
	goto L346
L563:
	;
	v2945 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v2946 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v2950 = *(*int32)(unsafe.Add(mBase, uint32(v2945+v2946<<(uint(int32(2))%32))))
	v2951 = *(*int32)(unsafe.Add(mBase, uint32(v2950)+4))
	v2954 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v2951+v2594+v2933))) = uint8(v2954)
	v2958 = v2933 + int32(1)
	goto L562
L564:
	;
	v2934 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+4))
	v2935 = *(*int32)(unsafe.Add(mBase, uint32(v2934)))
	goto L565
L565:
	;
	if int32(base.Ui32(v2935)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v2958 = v2933
		goto L562
	} else {
		goto L566
	}
L566:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_63))
	mBase = m.M
	v2944 = m.ExcPending
	if v2944 != 0 {
		goto L1
	} else {
		goto L567
	}
L567:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L568:
	;
	v2979 = v2973
	goto L569
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+28)) = v2979
	if v2979 != 0 {
		v3048 = v2979
		goto L346
	} else {
		goto L571
	}
L571:
	;
	v3002 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+4))
	v3003 = *(*int32)(unsafe.Add(mBase, uint32(v3002)))
	goto L572
L572:
	;
	if int32(base.Ui32(v3003)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v3048 = int32(0)
	goto L346
L574:
	;
	goto L575
L575:
	;
	v3012 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	if v3012 != int32(27) {
		goto L347
	} else {
		goto L576
	}
L576:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = int32(0)
	v3018 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+4))
	v3019 = *(*int32)(unsafe.Add(mBase, uint32(v3018)))
	*(*int32)(unsafe.Add(mBase, uint32(v3018))) = v3019 & int32(-49)
	goto L577
L577:
	;
	v3023 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v3024 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3028 = *(*int32)(unsafe.Add(mBase, uint32(v3023+v3024<<(uint(int32(2))%32))))
	v3029 = *(*int32)(unsafe.Add(mBase, uint32(v3028)+4))
	v3032 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+4))
	v3033 = F_fread(m, v3029+v2594, int32(1), v2887, v3032)
	mBase = m.M
	v3034 = m.ExcPending
	if v3034 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	v2979 = v3033
	goto L569
L579:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L580:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L581:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L582:
	;
	v3228 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+28))
	v3229 = v3228 + v2594
	v3230 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v3231 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3235 = *(*int32)(unsafe.Add(mBase, uint32(v3230+v3231<<(uint(int32(2))%32))))
	v3236 = *(*int32)(unsafe.Add(mBase, uint32(v3235)+12))
	if v3236 < v3229 {
		goto L606
	} else {
		goto L607
	}
L583:
	;
	if v2594 == int32(0) {
		goto L584
	} else {
		goto L585
	}
L584:
	;
	v3108 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+4))
	v3109 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	if v3109 != 0 {
		goto L589
	} else {
		goto L590
	}
L585:
	;
	goto L586
L586:
	;
	v3213 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v3214 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3215 = int32(2)
	v3218 = *(*int32)(unsafe.Add(mBase, uint32(v3213+v3214<<(uint(v3215)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3218)+44)) = v3215
	v3227 = v3215
	goto L582
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3175)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3175))) = v3108
	v3182 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	if v3182 != 0 {
		goto L602
	} else {
		goto L603
	}
L588:
	;
	v3132 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	v3133 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3137 = *(*int32)(unsafe.Add(mBase, uint32(v3130+v3133<<(uint(int32(2))%32))))
	if v3137 == int32(0) {
		goto L596
	} else {
		goto L597
	}
L589:
	;
	v3110 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3114 = *(*int32)(unsafe.Add(mBase, uint32(v3109+v3110<<(uint(int32(2))%32))))
	if v3114 != 0 {
		v3130 = v3109
		goto L588
	} else {
		goto L592
	}
L590:
	;
	goto L591
L591:
	;
	F_boot_yyensure_buffer_stack(m, v1377)
	mBase = m.M
	v3116 = m.ExcPending
	if v3116 != 0 {
		goto L1
	} else {
		goto L593
	}
L592:
	;
	goto L591
L593:
	;
	v3117 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+4))
	v3118 = F_boot_yy_create_buffer(m, v3117, v1377)
	mBase = m.M
	v3119 = m.ExcPending
	if v3119 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	v3120 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v3121 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3120+v3121<<(uint(int32(2))%32)))) = v3118
	v3126 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	if v3126 != 0 {
		v3130 = v3126
		goto L588
	} else {
		goto L595
	}
L595:
	;
	v3128 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	v3175 = int32(0)
	v3178 = v3128
	goto L587
L596:
	;
	v3175 = int32(0)
	v3178 = v3132
	goto L587
L597:
	;
	goto L598
L598:
	;
	v3141 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+16)) = v3141
	v3143 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v3143))) = uint8(v3141)
	v3146 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v3146)+1)) = uint8(v3141)
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+44)) = v3141
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+28)) = int32(1)
	v3153 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+8)) = v3153
	v3155 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	if v3155 == v3141 {
		v3175 = v3137
		v3178 = v3132
		goto L587
	} else {
		goto L599
	}
L599:
	;
	v3158 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3161 = v3155 + v3158<<(uint(int32(2))%32)
	v3162 = *(*int32)(unsafe.Add(mBase, uint32(v3161)))
	if v3137 != v3162 {
		v3175 = v3137
		v3178 = v3132
		goto L587
	} else {
		goto L600
	}
L600:
	;
	v3164 = *(*int32)(unsafe.Add(mBase, uint32(v3162)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+28)) = v3164
	v3166 = *(*int32)(unsafe.Add(mBase, uint32(v3161)))
	v3167 = *(*int32)(unsafe.Add(mBase, uint32(v3166)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+80)) = v3167
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+36)) = v3167
	v3170 = *(*int32)(unsafe.Add(mBase, uint32(v3161)))
	v3171 = *(*int32)(unsafe.Add(mBase, uint32(v3170)))
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+4)) = v3171
	v3173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3167))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1377)+24)) = uint8(v3173)
	v3175 = v3137
	v3178 = v3132
	goto L587
L601:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3175)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = v3178
	v3195 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v3196 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3199 = v3195 + v3196<<(uint(int32(2))%32)
	v3200 = *(*int32)(unsafe.Add(mBase, uint32(v3199)))
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(v3200)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+28)) = v3201
	v3203 = *(*int32)(unsafe.Add(mBase, uint32(v3199)))
	v3204 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+36)) = v3204
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+80)) = v3204
	v3207 = *(*int32)(unsafe.Add(mBase, uint32(v3199)))
	v3208 = *(*int32)(unsafe.Add(mBase, uint32(v3207)))
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+4)) = v3208
	v3210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3204))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1377)+24)) = uint8(v3210)
	v3227 = int32(1)
	goto L582
L602:
	;
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3187 = *(*int32)(unsafe.Add(mBase, uint32(v3182+v3183<<(uint(int32(2))%32))))
	if v3175 == v3187 {
		goto L601
	} else {
		goto L605
	}
L603:
	;
	goto L604
L604:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3175)+32)) = int64(1)
	goto L601
L605:
	;
	goto L604
L606:
	;
	v3240 = v3229 + v3228>>(uint(int32(1))%32)
	v3241 = *(*int32)(unsafe.Add(mBase, uint32(v3235)+4))
	if v3241 != 0 {
		goto L610
	} else {
		goto L611
	}
L607:
	;
	v3271 = v3230
	v3272 = v3229
	v3273 = v3231
	goto L608
L608:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+28)) = v3272
	v3275 = int32(2)
	v3278 = *(*int32)(unsafe.Add(mBase, uint32(v3271+v3273<<(uint(v3275)%32))))
	v3279 = *(*int32)(unsafe.Add(mBase, uint32(v3278)+4))
	v3281 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3279+v3272))) = uint8(v3281)
	v3283 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v3284 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3288 = *(*int32)(unsafe.Add(mBase, uint32(v3283+v3284<<(uint(v3275)%32))))
	v3289 = *(*int32)(unsafe.Add(mBase, uint32(v3288)+4))
	v3290 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v3289+v3290)+1)) = uint8(v3281)
	v3294 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v3295 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3298 = v3294 + v3295<<(uint(v3275)%32)
	v3299 = *(*int32)(unsafe.Add(mBase, uint32(v3298)))
	v3300 = *(*int32)(unsafe.Add(mBase, uint32(v3299)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+80)) = v3300
	if v3227 == int32(1) {
		v3635 = v3300
		goto L340
	} else {
		goto L616
	}
L609:
	;
	v3247 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v3248 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3249 = int32(2)
	v3252 = *(*int32)(unsafe.Add(mBase, uint32(v3247+v3248<<(uint(v3249)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3252)+4)) = v3246
	v3254 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v3255 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3259 = *(*int32)(unsafe.Add(mBase, uint32(v3254+v3255<<(uint(v3249)%32))))
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(v3259)+4))
	if v3260 == int32(0) {
		goto L342
	} else {
		goto L615
	}
L610:
	;
	v3242 = F_repalloc(m, v3241, v3240)
	mBase = m.M
	v3243 = m.ExcPending
	if v3243 != 0 {
		goto L1
	} else {
		goto L613
	}
L611:
	;
	goto L612
L612:
	;
	v3244 = F_palloc(m, v3240)
	mBase = m.M
	v3245 = m.ExcPending
	if v3245 != 0 {
		goto L1
	} else {
		goto L614
	}
L613:
	;
	v3246 = v3242
	goto L609
L614:
	;
	v3246 = v3244
	goto L609
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3259)+12)) = v3240 - int32(2)
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+12))
	v3267 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+28))
	v3269 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+20))
	v3271 = v3269
	v3272 = v3267 + v2594
	v3273 = v3266
	goto L608
L616:
	;
	switch v3227 - int32(1) {
	case 0:
		goto L343
	case 1:
		goto L617
	default:
		goto L618
	}
L617:
	;
	v3451 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+28))
	v3452 = *(*int32)(unsafe.Add(mBase, uint32(v3298)))
	v3453 = *(*int32)(unsafe.Add(mBase, uint32(v3452)+4))
	v3454 = v3300
	v3458 = v3451
	v3459 = v3453
	v3460 = v3294
	v3463 = v3295
	goto L344
L618:
	;
	v3308 = v2265 ^ int32(-1) + v1541
	v3309 = v3300 + v3308
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+36)) = v3309
	v3311 = *(*int32)(unsafe.Add(mBase, uint32(v3298)))
	v3312 = *(*int32)(unsafe.Add(mBase, uint32(v3311)+28))
	v3313 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+44))
	v3314 = v3312 + v3313
	if v3308 <= int32(0) {
		v1374 = v3300
		v1378 = v3309
		v1379 = v3314
		goto L319
	} else {
		goto L619
	}
L619:
	;
	v3322 = v3314
	v3325 = v3300
	goto L620
L620:
	;
	v3344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3325))))
	if v3344 != 0 {
		goto L622
	} else {
		goto L623
	}
L621:
	;
	v1374 = v3300
	v1378 = v3309
	v1379 = v3447
	goto L319
L622:
	;
	v3345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3344)+uint32(_c_F_BootstrapModeMain[42]))))
	v3346 = v3345
	goto L624
L623:
	;
	v3346 = int32(1)
	goto L624
L624:
	;
	v3348 = v3322 << (uint(int32(1)) % 32)
	v3351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3348)+uint32(_c_F_BootstrapModeMain[43]))))
	if v3351 != 0 {
		goto L625
	} else {
		goto L626
	}
L625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+68)) = v3325
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+64)) = v3322
	goto L627
L626:
	;
	goto L627
L627:
	;
	v3355 = v3346 & int32(255)
	v3358 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3348)+uint32(_c_F_BootstrapModeMain[44]))))
	v3359 = v3355 + v3358
	v3364 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3359<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[45]))))
	if v3364 != v3322 {
		goto L628
	} else {
		goto L629
	}
L628:
	;
	v3368 = v3346
	v3371 = v3322
	v3372 = v3355
	goto L631
L629:
	;
	v3426 = v3359
	goto L630
L630:
	;
	v3443 = int32(1)
	v3447 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3426<<(uint(v3443)%32))+uint32(_c_F_BootstrapModeMain[46]))))
	v3449 = v3325 + v3443
	if v3309 != v3449 {
		v3322 = v3447
		v3325 = v3449
		goto L620
	} else {
		goto L637
	}
L631:
	;
	v3396 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3371<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[47]))))
	if int32(128) <= v3396 {
		goto L633
	} else {
		goto L634
	}
L632:
	;
	v3426 = v3408
	goto L630
L633:
	;
	v3399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3372)+uint32(_c_F_BootstrapModeMain[48]))))
	v3400 = v3399
	goto L635
L634:
	;
	v3400 = v3368
	goto L635
L635:
	;
	v3402 = v3400 & int32(255)
	v3403 = int32(1)
	v3407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3396<<(uint(v3403)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	v3408 = v3402 + v3407
	v3413 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3408<<(uint(v3403)%32))+uint32(_c_F_BootstrapModeMain[45]))))
	if v3413 != v3396&int32(_a_F_BootstrapModeMain_60) {
		v3368 = v3400
		v3371 = v3396
		v3372 = v3402
		goto L631
	} else {
		goto L636
	}
L636:
	;
	goto L632
L637:
	;
	goto L621
L638:
	;
	v3494 = v3454
	v3495 = v3488
	goto L639
L639:
	;
	v3517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3494))))
	if v3517 != 0 {
		goto L641
	} else {
		goto L642
	}
L640:
	;
	v1534 = v3454
	v1538 = v3622
	v1541 = v3480
	goto L336
L641:
	;
	v3518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3517)+uint32(_c_F_BootstrapModeMain[42]))))
	v3519 = v3518
	goto L643
L642:
	;
	v3519 = int32(1)
	goto L643
L643:
	;
	v3524 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3495<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[43]))))
	if v3524 != 0 {
		goto L644
	} else {
		goto L645
	}
L644:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+68)) = v3494
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+64)) = v3495
	goto L646
L645:
	;
	goto L646
L646:
	;
	v3528 = v3519 & int32(255)
	v3529 = int32(1)
	v3533 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3495<<(uint(v3529)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	v3534 = v3528 + v3533
	v3539 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3534<<(uint(v3529)%32))+uint32(_c_F_BootstrapModeMain[45]))))
	if v3539 != v3495 {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v3543 = v3519
	v3546 = v3495
	v3547 = v3528
	goto L650
L648:
	;
	v3601 = v3534
	goto L649
L649:
	;
	v3618 = int32(1)
	v3622 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3601<<(uint(v3618)%32))+uint32(_c_F_BootstrapModeMain[46]))))
	v3624 = v3494 + v3618
	if v3624 != v3480 {
		v3494 = v3624
		v3495 = v3622
		goto L639
	} else {
		goto L656
	}
L650:
	;
	v3571 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3546<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[47]))))
	if int32(128) <= v3571 {
		goto L652
	} else {
		goto L653
	}
L651:
	;
	v3601 = v3583
	goto L649
L652:
	;
	v3574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3547)+uint32(_c_F_BootstrapModeMain[48]))))
	v3575 = v3574
	goto L654
L653:
	;
	v3575 = v3543
	goto L654
L654:
	;
	v3577 = v3575 & int32(255)
	v3578 = int32(1)
	v3582 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3571<<(uint(v3578)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	v3583 = v3577 + v3582
	v3588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3583<<(uint(v3578)%32))+uint32(_c_F_BootstrapModeMain[45]))))
	if v3588 != v3571&int32(_a_F_BootstrapModeMain_60) {
		v3543 = v3575
		v3546 = v3571
		v3547 = v3577
		goto L650
	} else {
		goto L655
	}
L655:
	;
	goto L651
L656:
	;
	goto L640
L657:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L658:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L659:
	;
	v3699 = int32(0)
	v3707 = v3699
	v3708 = v3699
	goto L293
L660:
	;
	goto L661
L661:
	;
	if v3676 == int32(256) {
		v6014 = v3681
		v6015 = v3682
		v6018 = v3685
		goto L262
	} else {
		goto L662
	}
L662:
	;
	if base.Ui32(int32(282)) < base.Ui32(v3676) {
		v3707 = v3676
		v3708 = int32(2)
		goto L293
	} else {
		goto L663
	}
L663:
	;
	v3706 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3676)+uint32(_c_F_BootstrapModeMain[49]))))
	v3707 = v3676
	v3708 = v3706
	goto L293
L664:
	;
	v3712 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3709)+uint32(_c_F_BootstrapModeMain[50]))))
	if v3708 != v3712 {
		v3730 = v3674
		v3732 = v3707
		v3734 = v3678
		v3737 = v3681
		v3738 = v3682
		v3740 = v3684
		v3741 = v3685
		v3743 = v3687
		v3744 = v3688
		v3745 = v3689
		v3749 = v3693
		goto L291
	} else {
		goto L665
	}
L665:
	;
	v3714 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3709)+uint32(_c_F_BootstrapModeMain[51]))))
	if int32(0) < v3714 {
		goto L666
	} else {
		goto L667
	}
L666:
	;
	v3717 = *(*int32)(unsafe.Add(mBase, uint32(v3681)+1132))
	*(*int32)(unsafe.Add(mBase, uint32(v3678)+4)) = v3717
	v5974 = v3674
	v5976 = int32(-2)
	v5978 = v3678 + int32(4)
	v5981 = v3681
	v5982 = v3682
	v5984 = v3714
	v5985 = v3685
	v5987 = v3687
	v5988 = v3688
	v5989 = v3689
	v5993 = v3693 - base.B2i32(v3693 != int32(0))
	goto L289
L667:
	;
	goto L668
L668:
	;
	v3761 = v3674
	v3763 = v3707
	v3765 = v3678
	v3768 = v3681
	v3769 = v3682
	v3771 = int32(0) - v3714
	v3772 = v3685
	v3774 = v3687
	v3775 = v3688
	v3776 = v3689
	v3780 = v3693
	goto L290
L669:
	;
	v3761 = v3730
	v3763 = v3732
	v3765 = v3734
	v3768 = v3737
	v3769 = v3738
	v3771 = v3755
	v3772 = v3741
	v3774 = v3743
	v3775 = v3744
	v3776 = v3745
	v3780 = v3749
	goto L290
L670:
	;
	v5946 = v3765 - v3787<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v5946)+4)) = v5920
	v5949 = v5946 + int32(4)
	v5950 = v3772 - v3787
	v5951 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5950))))
	v5954 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3771)+uint32(_c_F_BootstrapModeMain[52]))))
	v5957 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5954)+uint32(_c_F_BootstrapModeMain[53]))))
	v5958 = v5951 + v5957
	if base.Ui32(int32(169)) < base.Ui32(v5958) {
		goto L1151
	} else {
		goto L1152
	}
L671:
	;
	v5915 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5916 = F_pstrdup(m, v5915)
	mBase = m.M
	v5917 = m.ExcPending
	if v5917 != 0 {
		goto L1
	} else {
		goto L1150
	}
L672:
	;
	v5912 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5913 = F_pstrdup(m, v5912)
	mBase = m.M
	v5914 = m.ExcPending
	if v5914 != 0 {
		goto L1
	} else {
		goto L1149
	}
L673:
	;
	v5909 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5910 = F_pstrdup(m, v5909)
	mBase = m.M
	v5911 = m.ExcPending
	if v5911 != 0 {
		goto L1
	} else {
		goto L1148
	}
L674:
	;
	v5906 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5907 = F_pstrdup(m, v5906)
	mBase = m.M
	v5908 = m.ExcPending
	if v5908 != 0 {
		goto L1
	} else {
		goto L1147
	}
L675:
	;
	v5903 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5904 = F_pstrdup(m, v5903)
	mBase = m.M
	v5905 = m.ExcPending
	if v5905 != 0 {
		goto L1
	} else {
		goto L1146
	}
L676:
	;
	v5900 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5901 = F_pstrdup(m, v5900)
	mBase = m.M
	v5902 = m.ExcPending
	if v5902 != 0 {
		goto L1
	} else {
		goto L1145
	}
L677:
	;
	v5897 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5898 = F_pstrdup(m, v5897)
	mBase = m.M
	v5899 = m.ExcPending
	if v5899 != 0 {
		goto L1
	} else {
		goto L1144
	}
L678:
	;
	v5894 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5895 = F_pstrdup(m, v5894)
	mBase = m.M
	v5896 = m.ExcPending
	if v5896 != 0 {
		goto L1
	} else {
		goto L1143
	}
L679:
	;
	v5891 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5892 = F_pstrdup(m, v5891)
	mBase = m.M
	v5893 = m.ExcPending
	if v5893 != 0 {
		goto L1
	} else {
		goto L1142
	}
L680:
	;
	v5888 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5889 = F_pstrdup(m, v5888)
	mBase = m.M
	v5890 = m.ExcPending
	if v5890 != 0 {
		goto L1
	} else {
		goto L1141
	}
L681:
	;
	v5885 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5886 = F_pstrdup(m, v5885)
	mBase = m.M
	v5887 = m.ExcPending
	if v5887 != 0 {
		goto L1
	} else {
		goto L1140
	}
L682:
	;
	v5882 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5883 = F_pstrdup(m, v5882)
	mBase = m.M
	v5884 = m.ExcPending
	if v5884 != 0 {
		goto L1
	} else {
		goto L1139
	}
L683:
	;
	v5879 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5880 = F_pstrdup(m, v5879)
	mBase = m.M
	v5881 = m.ExcPending
	if v5881 != 0 {
		goto L1
	} else {
		goto L1138
	}
L684:
	;
	v5876 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5877 = F_pstrdup(m, v5876)
	mBase = m.M
	v5878 = m.ExcPending
	if v5878 != 0 {
		goto L1
	} else {
		goto L1137
	}
L685:
	;
	v5873 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5874 = F_pstrdup(m, v5873)
	mBase = m.M
	v5875 = m.ExcPending
	if v5875 != 0 {
		goto L1
	} else {
		goto L1136
	}
L686:
	;
	v5870 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5871 = F_pstrdup(m, v5870)
	mBase = m.M
	v5872 = m.ExcPending
	if v5872 != 0 {
		goto L1
	} else {
		goto L1135
	}
L687:
	;
	v5867 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5868 = F_pstrdup(m, v5867)
	mBase = m.M
	v5869 = m.ExcPending
	if v5869 != 0 {
		goto L1
	} else {
		goto L1134
	}
L688:
	;
	v5864 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5865 = F_pstrdup(m, v5864)
	mBase = m.M
	v5866 = m.ExcPending
	if v5866 != 0 {
		goto L1
	} else {
		goto L1133
	}
L689:
	;
	v5861 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5862 = F_pstrdup(m, v5861)
	mBase = m.M
	v5863 = m.ExcPending
	if v5863 != 0 {
		goto L1
	} else {
		goto L1132
	}
L690:
	;
	v5860 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5920 = v5860
	goto L670
L691:
	;
	v5782 = int32(_a_F_BootstrapModeMain_89)
	v5784 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54])) = v5784 + int32(1)
	v5788 = m.G0
	v5790 = v5788 - int32(32)
	m.G0 = v5790
	v5794 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5795 = m.ExcPending
	if v5795 != 0 {
		goto L1
	} else {
		goto L1120
	}
L692:
	;
	v5698 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5699 = int32(_a_F_BootstrapModeMain_89)
	v5701 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54])) = v5701 + int32(1)
	v5705 = m.G0
	v5707 = v5705 - int32(48)
	m.G0 = v5707
	v5711 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5712 = m.ExcPending
	if v5712 != 0 {
		goto L1
	} else {
		goto L1105
	}
L693:
	;
	v5692 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5696 = F_strtox_2(m, v5692, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	goto L1104
L694:
	;
	v5920 = int32(2)
	goto L670
L695:
	;
	v5920 = int32(3)
	goto L670
L696:
	;
	v4932 = int32(_a_F_BootstrapModeMain_90)
	v4934 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	v4936 = v4934 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55])) = v4936
	if int32(41) <= v4936 {
		goto L258
	} else {
		goto L984
	}
L697:
	;
	v4931 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5920 = v4931
	goto L670
L698:
	;
	v5920 = int32(0)
	goto L670
L699:
	;
	v5920 = int32(1)
	goto L670
L700:
	;
	v4904 = F_palloc0(m, int32(36))
	mBase = m.M
	v4905 = m.ExcPending
	if v4905 != 0 {
		goto L1
	} else {
		goto L981
	}
L701:
	;
	v4895 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+96)) = v4895
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+108)) = v4895
	v4901 = F_list_make1_impl(m, int32(1), v3768+int32(96))
	mBase = m.M
	v4902 = m.ExcPending
	if v4902 != 0 {
		goto L1
	} else {
		goto L980
	}
L702:
	;
	v4891 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(8))))
	v4892 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v4893 = F_lappend(m, v4891, v4892)
	mBase = m.M
	v4894 = m.ExcPending
	if v4894 != 0 {
		goto L1
	} else {
		goto L979
	}
L703:
	;
	v4769 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4769 == int32(0) {
		goto L956
	} else {
		goto L957
	}
L704:
	;
	v4651 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4652 = m.ExcPending
	if v4652 != 0 {
		goto L1
	} else {
		goto L923
	}
L705:
	;
	v4527 = F_palloc0(m, int32(72))
	mBase = m.M
	v4528 = m.ExcPending
	if v4528 != 0 {
		goto L1
	} else {
		goto L901
	}
L706:
	;
	v4406 = F_palloc0(m, int32(72))
	mBase = m.M
	v4407 = m.ExcPending
	if v4407 != 0 {
		goto L1
	} else {
		goto L879
	}
L707:
	;
	v4311 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54]))
	v4313 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	if v4311 != v4313 {
		goto L260
	} else {
		goto L849
	}
L708:
	;
	v4278 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4278 == int32(0) {
		goto L839
	} else {
		goto L840
	}
L709:
	;
	v4120 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4120 == int32(0) {
		goto L796
	} else {
		goto L797
	}
L710:
	;
	v4099 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4099
	v4102 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4102)
	mBase = m.M
	v4104 = m.ExcPending
	if v4104 != 0 {
		goto L1
	} else {
		goto L788
	}
L711:
	;
	v4046 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4046 == int32(0) {
		goto L774
	} else {
		goto L775
	}
L712:
	;
	v4004 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4004 == int32(0) {
		goto L761
	} else {
		goto L762
	}
L713:
	;
	v3796 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v3796 == int32(0) {
		goto L714
	} else {
		goto L715
	}
L714:
	;
	v3801 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v3806 = F_AllocSetContextCreateInternal(m, v3801, int32(_a_F_BootstrapModeMain_91), int32(0), int32(_a_F_BootstrapModeMain_24), int32(_a_F_BootstrapModeMain_92))
	mBase = m.M
	v3807 = m.ExcPending
	if v3807 != 0 {
		goto L1
	} else {
		goto L717
	}
L715:
	;
	v3809 = v3796
	goto L716
L716:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v3809
	v3812 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v3814 = m.G0
	v3816 = v3814 - int32(48)
	m.G0 = v3816
	v3818 = F_strlen(m, v3812)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v3818) {
		goto L718
	} else {
		goto L719
	}
L717:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v3806
	v3809 = v3806
	goto L716
L718:
	;
	v3821 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3812)+63)) = uint8(v3821)
	goto L720
L719:
	;
	goto L720
L720:
	;
	v3824 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[59]))
	if v3824 == int32(0) {
		goto L721
	} else {
		goto L722
	}
L721:
	;
	F_populate_typ_list(m)
	mBase = m.M
	v3828 = m.ExcPending
	if v3828 != 0 {
		goto L1
	} else {
		goto L724
	}
L722:
	;
	goto L723
L723:
	;
	v3830 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	if v3830 != 0 {
		goto L725
	} else {
		goto L726
	}
L724:
	;
	goto L723
L725:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v3833 = m.ExcPending
	if v3833 != 0 {
		goto L1
	} else {
		goto L728
	}
L726:
	;
	goto L727
L727:
	;
	v3836 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v3837 = m.ExcPending
	if v3837 != 0 {
		goto L1
	} else {
		goto L729
	}
L728:
	;
	goto L727
L729:
	;
	if v3836 != 0 {
		goto L730
	} else {
		goto L731
	}
L730:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3816)+36)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v3816)+32)) = v3812
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_93), v3816+int32(32))
	mBase = m.M
	v3845 = m.ExcPending
	if v3845 != 0 {
		goto L1
	} else {
		goto L733
	}
L731:
	;
	goto L732
L732:
	;
	v3854 = F_makeRangeVar(m, int32(0), v3812, int32(-1))
	mBase = m.M
	v3855 = m.ExcPending
	if v3855 != 0 {
		goto L1
	} else {
		goto L735
	}
L733:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(458), int32(_a_F_BootstrapModeMain_94))
	mBase = m.M
	v3850 = m.ExcPending
	if v3850 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	goto L732
L735:
	;
	v3857 = F_table_openrv(m, v3854, int32(0))
	mBase = m.M
	v3858 = m.ExcPending
	if v3858 != 0 {
		goto L1
	} else {
		goto L736
	}
L736:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60])) = v3857
	v3861 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+48))
	v3862 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3861)+120)))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55])) = v3862
	if int32(0) < v3862 {
		goto L737
	} else {
		goto L738
	}
L737:
	;
	v3874 = int32(0)
	goto L740
L738:
	;
	goto L739
L739:
	;
	m.G0 = v3816 + int32(48)
	v3983 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v3983
	v3986 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v3986)
	mBase = m.M
	v3988 = m.ExcPending
	if v3988 != 0 {
		goto L1
	} else {
		goto L753
	}
L740:
	;
	v3895 = v3874 << (uint(int32(2)) % 32)
	v3898 = *(*int32)(unsafe.Add(mBase, uint32(v3895)+uint32(_c_F_BootstrapModeMain[61])))
	if v3898 == int32(0) {
		goto L742
	} else {
		goto L743
	}
L741:
	;
	goto L739
L742:
	;
	v3902 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[62]))
	v3904 = F_MemoryContextAllocZero(m, v3902, int32(100))
	mBase = m.M
	v3905 = m.ExcPending
	if v3905 != 0 {
		goto L1
	} else {
		goto L745
	}
L743:
	;
	v3907 = v3898
	goto L744
L744:
	;
	v3909 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	v3910 = *(*int32)(unsafe.Add(mBase, uint32(v3909)+52))
	v3911 = *(*int32)(unsafe.Add(mBase, uint32(v3910)))
	v3915 = int32(100)
	base.MemoryCopy(m, v3907, v3910+v3911<<(uint(int32(4))%32)+v3874*v3915+int32(20), v3915)
	v3922 = *(*int32)(unsafe.Add(mBase, uint32(v3895)+uint32(_c_F_BootstrapModeMain[61])))
	v3925 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v3926 = m.ExcPending
	if v3926 != 0 {
		goto L1
	} else {
		goto L746
	}
L745:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3895)+uint32(_c_F_BootstrapModeMain[61]))) = v3904
	v3907 = v3904
	goto L744
L746:
	;
	if v3925 != 0 {
		goto L747
	} else {
		goto L748
	}
L747:
	;
	v3927 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3922)+72)))
	v3928 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3922)+74)))
	v3929 = *(*int32)(unsafe.Add(mBase, uint32(v3922)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v3816+int32(16)))) = v3929
	*(*int32)(unsafe.Add(mBase, uint32(v3816)+12)) = v3928
	*(*int32)(unsafe.Add(mBase, uint32(v3816)+8)) = v3927
	*(*int32)(unsafe.Add(mBase, uint32(v3816)+4)) = v3922 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v3816))) = v3874
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_95), v3816)
	mBase = m.M
	v3939 = m.ExcPending
	if v3939 != 0 {
		goto L1
	} else {
		goto L750
	}
L748:
	;
	goto L749
L749:
	;
	v3948 = v3874 + int32(1)
	v3950 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	if v3948 < v3950 {
		v3874 = v3948
		goto L740
	} else {
		goto L752
	}
L750:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(475), int32(_a_F_BootstrapModeMain_94))
	mBase = m.M
	v3944 = m.ExcPending
	if v3944 != 0 {
		goto L1
	} else {
		goto L751
	}
L751:
	;
	goto L749
L752:
	;
	goto L741
L753:
	;
	v3990 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v3990 != 0 {
		goto L754
	} else {
		goto L755
	}
L754:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3992 = m.ExcPending
	if v3992 != 0 {
		goto L1
	} else {
		goto L757
	}
L755:
	;
	goto L756
L756:
	;
	v3993 = int32(0)
	v3994 = F_isatty(m, v3993)
	mBase = m.M
	if v3994 == v3993 {
		v5920 = v3792
		goto L670
	} else {
		goto L758
	}
L757:
	;
	goto L756
L758:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_96), int32(0))
	mBase = m.M
	v4000 = m.ExcPending
	if v4000 != 0 {
		goto L1
	} else {
		goto L759
	}
L759:
	;
	v4001 = F_fflush(m, v3774)
	mBase = m.M
	v4002 = m.ExcPending
	if v4002 != 0 {
		goto L1
	} else {
		goto L760
	}
L760:
	;
	v5920 = v3792
	goto L670
L761:
	;
	v4009 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4014 = F_AllocSetContextCreateInternal(m, v4009, int32(_a_F_BootstrapModeMain_91), int32(0), int32(_a_F_BootstrapModeMain_24), int32(_a_F_BootstrapModeMain_92))
	mBase = m.M
	v4015 = m.ExcPending
	if v4015 != 0 {
		goto L1
	} else {
		goto L764
	}
L762:
	;
	v4017 = v4004
	goto L763
L763:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4017
	v4020 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	F_closerel(m, v4020)
	mBase = m.M
	v4022 = m.ExcPending
	if v4022 != 0 {
		goto L1
	} else {
		goto L765
	}
L764:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4014
	v4017 = v4014
	goto L763
L765:
	;
	v4025 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4025
	v4028 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4028)
	mBase = m.M
	v4030 = m.ExcPending
	if v4030 != 0 {
		goto L1
	} else {
		goto L766
	}
L766:
	;
	v4032 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4032 != 0 {
		goto L767
	} else {
		goto L768
	}
L767:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4034 = m.ExcPending
	if v4034 != 0 {
		goto L1
	} else {
		goto L770
	}
L768:
	;
	goto L769
L769:
	;
	v4035 = int32(0)
	v4036 = F_isatty(m, v4035)
	mBase = m.M
	if v4036 == v4035 {
		v5920 = v3792
		goto L670
	} else {
		goto L771
	}
L770:
	;
	goto L769
L771:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_96), int32(0))
	mBase = m.M
	v4042 = m.ExcPending
	if v4042 != 0 {
		goto L1
	} else {
		goto L772
	}
L772:
	;
	v4043 = F_fflush(m, v3774)
	mBase = m.M
	v4044 = m.ExcPending
	if v4044 != 0 {
		goto L1
	} else {
		goto L773
	}
L773:
	;
	v5920 = v3792
	goto L670
L774:
	;
	v4051 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4056 = F_AllocSetContextCreateInternal(m, v4051, int32(_a_F_BootstrapModeMain_91), int32(0), int32(_a_F_BootstrapModeMain_24), int32(_a_F_BootstrapModeMain_92))
	mBase = m.M
	v4057 = m.ExcPending
	if v4057 != 0 {
		goto L1
	} else {
		goto L777
	}
L775:
	;
	v4059 = v4046
	goto L776
L776:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4059
	v4063 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55])) = v4063
	v4067 = F_errstart(m, int32(11), v4063)
	mBase = m.M
	v4068 = m.ExcPending
	if v4068 != 0 {
		goto L1
	} else {
		goto L778
	}
L777:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4056
	v4059 = v4056
	goto L776
L778:
	;
	if v4067 == int32(0) {
		v5920 = v3792
		goto L670
	} else {
		goto L779
	}
L779:
	;
	v4073 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(12))))
	v4076 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(8))))
	v4079 = *(*int64)(unsafe.Add(mBase, uint32(v3765-int32(20))))
	*(*int64)(unsafe.Add(mBase, uint32(v3768)+8)) = v4079
	if v4076 != 0 {
		goto L780
	} else {
		goto L781
	}
L780:
	;
	v4083 = int32(_a_F_BootstrapModeMain_97)
	goto L782
L781:
	;
	v4083 = int32(_a_F_BootstrapModeMain_98)
	goto L782
L782:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+4)) = v4083
	if v4073 != 0 {
		goto L783
	} else {
		goto L784
	}
L783:
	;
	v4087 = int32(_a_F_BootstrapModeMain_99)
	goto L785
L784:
	;
	v4087 = int32(_a_F_BootstrapModeMain_98)
	goto L785
L785:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3768))) = v4087
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_100), v3768)
	mBase = m.M
	v4091 = m.ExcPending
	if v4091 != 0 {
		goto L1
	} else {
		goto L786
	}
L786:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_101), int32(166), int32(_a_F_BootstrapModeMain_102))
	mBase = m.M
	v4096 = m.ExcPending
	if v4096 != 0 {
		goto L1
	} else {
		goto L787
	}
L787:
	;
	v5920 = v3792
	goto L670
L788:
	;
	v4106 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4106 != 0 {
		goto L789
	} else {
		goto L790
	}
L789:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4108 = m.ExcPending
	if v4108 != 0 {
		goto L1
	} else {
		goto L792
	}
L790:
	;
	goto L791
L791:
	;
	v4109 = int32(0)
	v4110 = F_isatty(m, v4109)
	mBase = m.M
	if v4110 == v4109 {
		v5920 = v3792
		goto L670
	} else {
		goto L793
	}
L792:
	;
	goto L791
L793:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_96), int32(0))
	mBase = m.M
	v4116 = m.ExcPending
	if v4116 != 0 {
		goto L1
	} else {
		goto L794
	}
L794:
	;
	v4117 = F_fflush(m, v3774)
	mBase = m.M
	v4118 = m.ExcPending
	if v4118 != 0 {
		goto L1
	} else {
		goto L795
	}
L795:
	;
	v5920 = v3792
	goto L670
L796:
	;
	v4125 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4130 = F_AllocSetContextCreateInternal(m, v4125, int32(_a_F_BootstrapModeMain_91), int32(0), int32(_a_F_BootstrapModeMain_24), int32(_a_F_BootstrapModeMain_92))
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		goto L1
	} else {
		goto L799
	}
L797:
	;
	v4133 = v4120
	goto L798
L798:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4133
	v4137 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	v4139 = F_CreateTupleDesc(m, v4137, int32(_a_F_BootstrapModeMain_57))
	mBase = m.M
	v4140 = m.ExcPending
	if v4140 != 0 {
		goto L1
	} else {
		goto L800
	}
L799:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4130
	v4133 = v4130
	goto L798
L800:
	;
	v4143 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(24))))
	v4146 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(28))))
	if v4146 != 0 {
		goto L802
	} else {
		goto L803
	}
L801:
	;
	v4257 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4257
	v4260 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4260)
	mBase = m.M
	v4262 = m.ExcPending
	if v4262 != 0 {
		goto L1
	} else {
		goto L831
	}
L802:
	;
	v4148 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	if v4148 != 0 {
		goto L805
	} else {
		goto L806
	}
L803:
	;
	goto L804
L804:
	;
	v4209 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(36))))
	if v4143 != 0 {
		goto L823
	} else {
		goto L824
	}
L805:
	;
	v4151 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4152 = m.ExcPending
	if v4152 != 0 {
		goto L1
	} else {
		goto L808
	}
L806:
	;
	goto L807
L807:
	;
	v4168 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(36))))
	if v4143 != 0 {
		goto L815
	} else {
		goto L816
	}
L808:
	;
	if v4151 != 0 {
		goto L809
	} else {
		goto L810
	}
L809:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_103), int32(0))
	mBase = m.M
	v4156 = m.ExcPending
	if v4156 != 0 {
		goto L1
	} else {
		goto L812
	}
L810:
	;
	goto L811
L811:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v4164 = m.ExcPending
	if v4164 != 0 {
		goto L1
	} else {
		goto L814
	}
L812:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_101), int32(202), int32(_a_F_BootstrapModeMain_102))
	mBase = m.M
	v4161 = m.ExcPending
	if v4161 != 0 {
		goto L1
	} else {
		goto L813
	}
L813:
	;
	goto L811
L814:
	;
	goto L807
L815:
	;
	v4172 = int32(1664)
	goto L817
L816:
	;
	v4172 = int32(0)
	goto L817
L817:
	;
	v4175 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(32))))
	v4176 = int32(0)
	v4179 = int32(112)
	v4182 = int32(1)
	v4189 = F_heap_create(m, v4168, int32(11), v4172, v4175, v4176, int32(2), v4139, int32(114), v4179, base.B2i32(v4143 != v4176), v4182, v4182, v3768+v4179, v3768+int32(124), v4182)
	mBase = m.M
	v4190 = m.ExcPending
	if v4190 != 0 {
		goto L1
	} else {
		goto L818
	}
L818:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60])) = v4189
	v4194 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4195 = m.ExcPending
	if v4195 != 0 {
		goto L1
	} else {
		goto L819
	}
L819:
	;
	if v4194 == int32(0) {
		goto L801
	} else {
		goto L820
	}
L820:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_104), int32(0))
	mBase = m.M
	v4201 = m.ExcPending
	if v4201 != 0 {
		goto L1
	} else {
		goto L821
	}
L821:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_101), int32(221), int32(_a_F_BootstrapModeMain_102))
	mBase = m.M
	v4206 = m.ExcPending
	if v4206 != 0 {
		goto L1
	} else {
		goto L822
	}
L822:
	;
	goto L801
L823:
	;
	v4213 = int32(1664)
	goto L825
L824:
	;
	v4213 = int32(0)
	goto L825
L825:
	;
	v4216 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(32))))
	v4219 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(20))))
	v4220 = int32(0)
	v4227 = base.B2i32(v4143 != v4220)
	v4235 = F_heap_create_with_catalog(m, v4209, int32(11), v4213, v4216, v4219, v4220, int32(10), int32(2), v4139, v4220, int32(114), int32(112), v4227, v4227, v4220, v4220, v4220, int32(1), v4220, v4220, v4220)
	mBase = m.M
	v4236 = m.ExcPending
	if v4236 != 0 {
		goto L1
	} else {
		goto L826
	}
L826:
	;
	v4239 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4240 = m.ExcPending
	if v4240 != 0 {
		goto L1
	} else {
		goto L827
	}
L827:
	;
	if v4239 == int32(0) {
		goto L801
	} else {
		goto L828
	}
L828:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+16)) = v4235
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_105), v3768+int32(16))
	mBase = m.M
	v4248 = m.ExcPending
	if v4248 != 0 {
		goto L1
	} else {
		goto L829
	}
L829:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_101), int32(248), int32(_a_F_BootstrapModeMain_102))
	mBase = m.M
	v4253 = m.ExcPending
	if v4253 != 0 {
		goto L1
	} else {
		goto L830
	}
L830:
	;
	goto L801
L831:
	;
	v4264 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4264 != 0 {
		goto L832
	} else {
		goto L833
	}
L832:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4266 = m.ExcPending
	if v4266 != 0 {
		goto L1
	} else {
		goto L835
	}
L833:
	;
	goto L834
L834:
	;
	v4267 = int32(0)
	v4268 = F_isatty(m, v4267)
	mBase = m.M
	if v4268 == v4267 {
		v5920 = v3792
		goto L670
	} else {
		goto L836
	}
L835:
	;
	goto L834
L836:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_96), int32(0))
	mBase = m.M
	v4274 = m.ExcPending
	if v4274 != 0 {
		goto L1
	} else {
		goto L837
	}
L837:
	;
	v4275 = F_fflush(m, v3774)
	mBase = m.M
	v4276 = m.ExcPending
	if v4276 != 0 {
		goto L1
	} else {
		goto L838
	}
L838:
	;
	v5920 = v3792
	goto L670
L839:
	;
	v4283 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4288 = F_AllocSetContextCreateInternal(m, v4283, int32(_a_F_BootstrapModeMain_91), int32(0), int32(_a_F_BootstrapModeMain_24), int32(_a_F_BootstrapModeMain_92))
	mBase = m.M
	v4289 = m.ExcPending
	if v4289 != 0 {
		goto L1
	} else {
		goto L842
	}
L840:
	;
	v4291 = v4278
	goto L841
L841:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4291
	v4296 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4297 = m.ExcPending
	if v4297 != 0 {
		goto L1
	} else {
		goto L843
	}
L842:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4288
	v4291 = v4288
	goto L841
L843:
	;
	if v4296 != 0 {
		goto L844
	} else {
		goto L845
	}
L844:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_106), int32(0))
	mBase = m.M
	v4301 = m.ExcPending
	if v4301 != 0 {
		goto L1
	} else {
		goto L847
	}
L845:
	;
	goto L846
L846:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54])) = int32(0)
	v5920 = v3792
	goto L670
L847:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_101), int32(258), int32(_a_F_BootstrapModeMain_102))
	mBase = m.M
	v4306 = m.ExcPending
	if v4306 != 0 {
		goto L1
	} else {
		goto L848
	}
L848:
	;
	goto L846
L849:
	;
	v4316 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	if v4316 == int32(0) {
		goto L259
	} else {
		goto L850
	}
L850:
	;
	v4319 = m.G0
	v4321 = v4319 - int32(16)
	m.G0 = v4321
	v4325 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4326 = m.ExcPending
	if v4326 != 0 {
		goto L1
	} else {
		goto L851
	}
L851:
	;
	if v4325 != 0 {
		goto L852
	} else {
		goto L853
	}
L852:
	;
	v4328 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	*(*int32)(unsafe.Add(mBase, uint32(v4321))) = v4328
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_107), v4321)
	mBase = m.M
	v4332 = m.ExcPending
	if v4332 != 0 {
		goto L1
	} else {
		goto L855
	}
L853:
	;
	goto L854
L854:
	;
	v4339 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	v4341 = F_CreateTupleDesc(m, v4339, int32(_a_F_BootstrapModeMain_57))
	mBase = m.M
	v4342 = m.ExcPending
	if v4342 != 0 {
		goto L1
	} else {
		goto L857
	}
L855:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(635), int32(_a_F_BootstrapModeMain_108))
	mBase = m.M
	v4337 = m.ExcPending
	if v4337 != 0 {
		goto L1
	} else {
		goto L856
	}
L856:
	;
	goto L854
L857:
	;
	v4345 = F_heap_form_tuple(m, v4341, int32(_a_F_BootstrapModeMain_109), int32(_a_F_BootstrapModeMain_110))
	mBase = m.M
	v4346 = m.ExcPending
	if v4346 != 0 {
		goto L1
	} else {
		goto L858
	}
L858:
	;
	F_pfree(m, v4341)
	mBase = m.M
	v4348 = m.ExcPending
	if v4348 != 0 {
		goto L1
	} else {
		goto L859
	}
L859:
	;
	v4350 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	F_simple_heap_insert(m, v4350, v4345)
	mBase = m.M
	v4352 = m.ExcPending
	if v4352 != 0 {
		goto L1
	} else {
		goto L860
	}
L860:
	;
	F_pfree(m, v4345)
	mBase = m.M
	v4354 = m.ExcPending
	if v4354 != 0 {
		goto L1
	} else {
		goto L861
	}
L861:
	;
	v4357 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4358 = m.ExcPending
	if v4358 != 0 {
		goto L1
	} else {
		goto L862
	}
L862:
	;
	if v4357 != 0 {
		goto L863
	} else {
		goto L864
	}
L863:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_111), int32(0))
	mBase = m.M
	v4362 = m.ExcPending
	if v4362 != 0 {
		goto L1
	} else {
		goto L866
	}
L864:
	;
	goto L865
L865:
	;
	v4369 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	v4370 = int32(0)
	if base.B2i32(v4369 <= v4370)|base.B2i32(v4369 == v4370) == v4370 {
		goto L868
	} else {
		goto L869
	}
L866:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(643), int32(_a_F_BootstrapModeMain_108))
	mBase = m.M
	v4367 = m.ExcPending
	if v4367 != 0 {
		goto L1
	} else {
		goto L867
	}
L867:
	;
	goto L865
L868:
	;
	base.MemoryFill(m, int32(_a_F_BootstrapModeMain_110), int32(0), v4369)
	goto L870
L869:
	;
	goto L870
L870:
	;
	m.G0 = v4321 + int32(16)
	v4385 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4385
	v4388 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4388)
	mBase = m.M
	v4390 = m.ExcPending
	if v4390 != 0 {
		goto L1
	} else {
		goto L871
	}
L871:
	;
	v4392 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4392 != 0 {
		goto L872
	} else {
		goto L873
	}
L872:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4394 = m.ExcPending
	if v4394 != 0 {
		goto L1
	} else {
		goto L875
	}
L873:
	;
	goto L874
L874:
	;
	v4395 = int32(0)
	v4396 = F_isatty(m, v4395)
	mBase = m.M
	if v4396 == v4395 {
		v5920 = v3792
		goto L670
	} else {
		goto L876
	}
L875:
	;
	goto L874
L876:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_96), int32(0))
	mBase = m.M
	v4402 = m.ExcPending
	if v4402 != 0 {
		goto L1
	} else {
		goto L877
	}
L877:
	;
	v4403 = F_fflush(m, v3774)
	mBase = m.M
	v4404 = m.ExcPending
	if v4404 != 0 {
		goto L1
	} else {
		goto L878
	}
L878:
	;
	v5920 = v3792
	goto L670
L879:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4406))) = int32(204)
	v4412 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4413 = m.ExcPending
	if v4413 != 0 {
		goto L1
	} else {
		goto L880
	}
L880:
	;
	if v4412 != 0 {
		goto L881
	} else {
		goto L882
	}
L881:
	;
	v4416 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+48)) = v4416
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_112), v3768+int32(48))
	mBase = m.M
	v4422 = m.ExcPending
	if v4422 != 0 {
		goto L1
	} else {
		goto L884
	}
L882:
	;
	goto L883
L883:
	;
	v4429 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4429 == int32(0) {
		goto L886
	} else {
		goto L887
	}
L884:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_101), int32(279), int32(_a_F_BootstrapModeMain_102))
	mBase = m.M
	v4427 = m.ExcPending
	if v4427 != 0 {
		goto L1
	} else {
		goto L885
	}
L885:
	;
	goto L883
L886:
	;
	v4434 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4439 = F_AllocSetContextCreateInternal(m, v4434, int32(_a_F_BootstrapModeMain_91), int32(0), int32(_a_F_BootstrapModeMain_24), int32(_a_F_BootstrapModeMain_92))
	mBase = m.M
	v4440 = m.ExcPending
	if v4440 != 0 {
		goto L1
	} else {
		goto L889
	}
L887:
	;
	v4442 = v4429
	goto L888
L888:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4442
	v4447 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v4406)+4)) = v4447
	v4452 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(20))))
	v4454 = F_makeRangeVar(m, int32(0), v4452, int32(-1))
	mBase = m.M
	v4455 = m.ExcPending
	if v4455 != 0 {
		goto L1
	} else {
		goto L890
	}
L889:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4439
	v4442 = v4439
	goto L888
L890:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4406)+8)) = v4454
	v4459 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(12))))
	v4460 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4406)+16)) = v4460
	*(*int32)(unsafe.Add(mBase, uint32(v4406)+12)) = v4459
	v4465 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(4))))
	*(*uint16)(unsafe.Add(mBase, uint32(v4406)+62)) = uint16(v4460)
	*(*int32)(unsafe.Add(mBase, uint32(v4406)+20)) = v4465
	v4469 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4406)+24)) = v4469
	*(*int64)(unsafe.Add(mBase, uint32(v4406)+32)) = v4469
	*(*int64)(unsafe.Add(mBase, uint32(v4406)+40)) = v4469
	*(*int64)(unsafe.Add(mBase, uint32(v4406)+48)) = v4469
	*(*int64)(unsafe.Add(mBase, uint32(v4406)+53)) = v4469
	*(*int32)(unsafe.Add(mBase, uint32(v4406)+65)) = v4460
	*(*uint16)(unsafe.Add(mBase, uint32(v4406)+69)) = uint16(v4460)
	v4489 = F_RangeVarGetRelidExtended(m, v4454, v4460, v4460, v4460, v4460)
	mBase = m.M
	v4490 = m.ExcPending
	if v4490 != 0 {
		goto L1
	} else {
		goto L891
	}
L891:
	;
	v4493 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(28))))
	v4494 = int32(0)
	F_DefineIndex(m, v3768+int32(112), v4489, v4406, v4493, v4494, v4494, int32(-1), v4494, v4494, v4494, int32(1), v4494)
	mBase = m.M
	v4503 = m.ExcPending
	if v4503 != 0 {
		goto L1
	} else {
		goto L892
	}
L892:
	;
	v4506 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4506
	v4509 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4509)
	mBase = m.M
	v4511 = m.ExcPending
	if v4511 != 0 {
		goto L1
	} else {
		goto L893
	}
L893:
	;
	v4513 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4513 != 0 {
		goto L894
	} else {
		goto L895
	}
L894:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4515 = m.ExcPending
	if v4515 != 0 {
		goto L1
	} else {
		goto L897
	}
L895:
	;
	goto L896
L896:
	;
	v4516 = int32(0)
	v4517 = F_isatty(m, v4516)
	mBase = m.M
	if v4517 == v4516 {
		v5920 = v3792
		goto L670
	} else {
		goto L898
	}
L897:
	;
	goto L896
L898:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_96), int32(0))
	mBase = m.M
	v4523 = m.ExcPending
	if v4523 != 0 {
		goto L1
	} else {
		goto L899
	}
L899:
	;
	v4524 = F_fflush(m, v3774)
	mBase = m.M
	v4525 = m.ExcPending
	if v4525 != 0 {
		goto L1
	} else {
		goto L900
	}
L900:
	;
	v5920 = v3792
	goto L670
L901:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4527))) = int32(204)
	v4533 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4534 = m.ExcPending
	if v4534 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	if v4533 != 0 {
		goto L903
	} else {
		goto L904
	}
L903:
	;
	v4537 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+64)) = v4537
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_113), v3768-int32(-64))
	mBase = m.M
	v4543 = m.ExcPending
	if v4543 != 0 {
		goto L1
	} else {
		goto L906
	}
L904:
	;
	goto L905
L905:
	;
	v4550 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4550 == int32(0) {
		goto L908
	} else {
		goto L909
	}
L906:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_101), int32(332), int32(_a_F_BootstrapModeMain_102))
	mBase = m.M
	v4548 = m.ExcPending
	if v4548 != 0 {
		goto L1
	} else {
		goto L907
	}
L907:
	;
	goto L905
L908:
	;
	v4555 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4560 = F_AllocSetContextCreateInternal(m, v4555, int32(_a_F_BootstrapModeMain_91), int32(0), int32(_a_F_BootstrapModeMain_24), int32(_a_F_BootstrapModeMain_92))
	mBase = m.M
	v4561 = m.ExcPending
	if v4561 != 0 {
		goto L1
	} else {
		goto L911
	}
L909:
	;
	v4563 = v4550
	goto L910
L910:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4563
	v4568 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v4527)+4)) = v4568
	v4573 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(20))))
	v4575 = F_makeRangeVar(m, int32(0), v4573, int32(-1))
	mBase = m.M
	v4576 = m.ExcPending
	if v4576 != 0 {
		goto L1
	} else {
		goto L912
	}
L911:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4560
	v4563 = v4560
	goto L910
L912:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4527)+8)) = v4575
	v4580 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(12))))
	v4581 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4527)+16)) = v4581
	*(*int32)(unsafe.Add(mBase, uint32(v4527)+12)) = v4580
	v4586 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(4))))
	v4587 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4527)+24)) = v4587
	*(*int32)(unsafe.Add(mBase, uint32(v4527)+20)) = v4586
	*(*int64)(unsafe.Add(mBase, uint32(v4527)+32)) = v4587
	*(*int64)(unsafe.Add(mBase, uint32(v4527)+40)) = v4587
	*(*int64)(unsafe.Add(mBase, uint32(v4527)+48)) = v4587
	*(*int32)(unsafe.Add(mBase, uint32(v4527)+56)) = v4581
	*(*int32)(unsafe.Add(mBase, uint32(v4527)+65)) = v4581
	*(*uint16)(unsafe.Add(mBase, uint32(v4527)+62)) = uint16(v4581)
	v4602 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4527)+60)) = uint8(v4602)
	*(*uint16)(unsafe.Add(mBase, uint32(v4527)+69)) = uint16(v4581)
	v4612 = F_RangeVarGetRelidExtended(m, v4575, v4581, v4581, v4581, v4581)
	mBase = m.M
	v4613 = m.ExcPending
	if v4613 != 0 {
		goto L1
	} else {
		goto L913
	}
L913:
	;
	v4616 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(28))))
	v4617 = int32(0)
	F_DefineIndex(m, v3768+int32(112), v4612, v4527, v4616, v4617, v4617, int32(-1), v4617, v4617, v4617, int32(1), v4617)
	mBase = m.M
	v4626 = m.ExcPending
	if v4626 != 0 {
		goto L1
	} else {
		goto L914
	}
L914:
	;
	v4629 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4629
	v4632 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4632)
	mBase = m.M
	v4634 = m.ExcPending
	if v4634 != 0 {
		goto L1
	} else {
		goto L915
	}
L915:
	;
	v4636 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4636 != 0 {
		goto L916
	} else {
		goto L917
	}
L916:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4638 = m.ExcPending
	if v4638 != 0 {
		goto L1
	} else {
		goto L919
	}
L917:
	;
	goto L918
L918:
	;
	v4639 = int32(0)
	v4640 = F_isatty(m, v4639)
	mBase = m.M
	if v4640 == v4639 {
		v5920 = v3792
		goto L670
	} else {
		goto L920
	}
L919:
	;
	goto L918
L920:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_96), int32(0))
	mBase = m.M
	v4646 = m.ExcPending
	if v4646 != 0 {
		goto L1
	} else {
		goto L921
	}
L921:
	;
	v4647 = F_fflush(m, v3774)
	mBase = m.M
	v4648 = m.ExcPending
	if v4648 != 0 {
		goto L1
	} else {
		goto L922
	}
L922:
	;
	v5920 = v3792
	goto L670
L923:
	;
	if v4651 != 0 {
		goto L924
	} else {
		goto L925
	}
L924:
	;
	v4653 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+80)) = v4653
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_114), v3768+int32(80))
	mBase = m.M
	v4659 = m.ExcPending
	if v4659 != 0 {
		goto L1
	} else {
		goto L927
	}
L925:
	;
	goto L926
L926:
	;
	v4666 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4666 == int32(0) {
		goto L929
	} else {
		goto L930
	}
L927:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_101), int32(382), int32(_a_F_BootstrapModeMain_102))
	mBase = m.M
	v4664 = m.ExcPending
	if v4664 != 0 {
		goto L1
	} else {
		goto L928
	}
L928:
	;
	goto L926
L929:
	;
	v4671 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4676 = F_AllocSetContextCreateInternal(m, v4671, int32(_a_F_BootstrapModeMain_91), int32(0), int32(_a_F_BootstrapModeMain_24), int32(_a_F_BootstrapModeMain_92))
	mBase = m.M
	v4677 = m.ExcPending
	if v4677 != 0 {
		goto L1
	} else {
		goto L932
	}
L930:
	;
	v4679 = v4666
	goto L931
L931:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4679
	v4682 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v4685 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(12))))
	v4688 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(8))))
	v4689 = m.G0
	v4691 = v4689 - int32(32)
	m.G0 = v4691
	v4695 = F_makeRangeVar(m, int32(0), v4682, int32(-1))
	mBase = m.M
	v4696 = m.ExcPending
	if v4696 != 0 {
		goto L1
	} else {
		goto L935
	}
L932:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4676
	v4679 = v4676
	goto L931
L933:
	;
	v4717 = int32(0)
	v4721 = F_create_toast_table(m, v4698, v4685, v4688, v4717, int32(8), v4717, v4717)
	mBase = m.M
	v4722 = m.ExcPending
	if v4722 != 0 {
		goto L1
	} else {
		goto L940
	}
L934:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4707 = m.ExcPending
	if v4707 != 0 {
		goto L1
	} else {
		goto L937
	}
L935:
	;
	v4698 = F_table_openrv(m, v4695, int32(8))
	mBase = m.M
	v4699 = m.ExcPending
	if v4699 != 0 {
		goto L1
	} else {
		goto L936
	}
L936:
	;
	v4700 = *(*int32)(unsafe.Add(mBase, uint32(v4698)+48))
	v4701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4700)+119)))
	switch v4701 - int32(109) {
	case 0, 5:
		goto L933
	default:
		goto L934
	}
L937:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4691))) = v4682
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_115), v4691)
	mBase = m.M
	v4711 = m.ExcPending
	if v4711 != 0 {
		goto L1
	} else {
		goto L938
	}
L938:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_116), int32(107), int32(_a_F_BootstrapModeMain_117))
	mBase = m.M
	v4716 = m.ExcPending
	if v4716 != 0 {
		goto L1
	} else {
		goto L939
	}
L939:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L940:
	;
	if v4721 == int32(0) {
		goto L941
	} else {
		goto L942
	}
L941:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4728 = m.ExcPending
	if v4728 != 0 {
		goto L1
	} else {
		goto L944
	}
L942:
	;
	goto L943
L943:
	;
	F_relation_close(m, v4698, int32(0))
	mBase = m.M
	v4742 = m.ExcPending
	if v4742 != 0 {
		goto L1
	} else {
		goto L947
	}
L944:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4691)+16)) = v4682
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_118), v4691+int32(16))
	mBase = m.M
	v4734 = m.ExcPending
	if v4734 != 0 {
		goto L1
	} else {
		goto L945
	}
L945:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_116), int32(113), int32(_a_F_BootstrapModeMain_117))
	mBase = m.M
	v4739 = m.ExcPending
	if v4739 != 0 {
		goto L1
	} else {
		goto L946
	}
L946:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L947:
	;
	m.G0 = v4691 + int32(32)
	v4748 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4748
	v4751 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4751)
	mBase = m.M
	v4753 = m.ExcPending
	if v4753 != 0 {
		goto L1
	} else {
		goto L948
	}
L948:
	;
	v4755 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4755 != 0 {
		goto L949
	} else {
		goto L950
	}
L949:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4757 = m.ExcPending
	if v4757 != 0 {
		goto L1
	} else {
		goto L952
	}
L950:
	;
	goto L951
L951:
	;
	v4758 = int32(0)
	v4759 = F_isatty(m, v4758)
	mBase = m.M
	if v4759 == v4758 {
		v5920 = v3792
		goto L670
	} else {
		goto L953
	}
L952:
	;
	goto L951
L953:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_96), int32(0))
	mBase = m.M
	v4765 = m.ExcPending
	if v4765 != 0 {
		goto L1
	} else {
		goto L954
	}
L954:
	;
	v4766 = F_fflush(m, v3774)
	mBase = m.M
	v4767 = m.ExcPending
	if v4767 != 0 {
		goto L1
	} else {
		goto L955
	}
L955:
	;
	v5920 = v3792
	goto L670
L956:
	;
	v4774 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4779 = F_AllocSetContextCreateInternal(m, v4774, int32(_a_F_BootstrapModeMain_91), int32(0), int32(_a_F_BootstrapModeMain_24), int32(_a_F_BootstrapModeMain_92))
	mBase = m.M
	v4780 = m.ExcPending
	if v4780 != 0 {
		goto L1
	} else {
		goto L959
	}
L957:
	;
	v4782 = v4769
	goto L958
L958:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4782
	v4786 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[64]))
	if v4786 != 0 {
		goto L960
	} else {
		goto L961
	}
L959:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4779
	v4782 = v4779
	goto L958
L960:
	;
	v4791 = v4786
	goto L963
L961:
	;
	goto L962
L962:
	;
	v4869 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4869
	v4872 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4872)
	mBase = m.M
	v4874 = m.ExcPending
	if v4874 != 0 {
		goto L1
	} else {
		goto L971
	}
L963:
	;
	v4813 = *(*int32)(unsafe.Add(mBase, uint32(v4791)))
	v4815 = F_table_open(m, v4813, int32(0))
	mBase = m.M
	v4816 = m.ExcPending
	if v4816 != 0 {
		goto L1
	} else {
		goto L965
	}
L964:
	;
	goto L962
L965:
	;
	v4818 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[64]))
	v4819 = *(*int32)(unsafe.Add(mBase, uint32(v4818)+4))
	v4821 = F_index_open(m, v4819, int32(0))
	mBase = m.M
	v4822 = m.ExcPending
	if v4822 != 0 {
		goto L1
	} else {
		goto L966
	}
L966:
	;
	v4824 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[64]))
	v4825 = *(*int32)(unsafe.Add(mBase, uint32(v4824)+8))
	v4826 = int32(0)
	F_index_build(m, v4815, v4821, v4825, v4826, v4826)
	mBase = m.M
	v4829 = m.ExcPending
	if v4829 != 0 {
		goto L1
	} else {
		goto L967
	}
L967:
	;
	F_relation_close(m, v4821, int32(0))
	mBase = m.M
	v4832 = m.ExcPending
	if v4832 != 0 {
		goto L1
	} else {
		goto L968
	}
L968:
	;
	F_relation_close(m, v4815, int32(0))
	mBase = m.M
	v4835 = m.ExcPending
	if v4835 != 0 {
		goto L1
	} else {
		goto L969
	}
L969:
	;
	v4836 = int32(_a_F_BootstrapModeMain_119)
	v4838 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[64]))
	v4839 = *(*int32)(unsafe.Add(mBase, uint32(v4838)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[64])) = v4839
	if v4839 != 0 {
		v4791 = v4839
		goto L963
	} else {
		goto L970
	}
L970:
	;
	goto L964
L971:
	;
	v4876 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4876 != 0 {
		goto L972
	} else {
		goto L973
	}
L972:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4878 = m.ExcPending
	if v4878 != 0 {
		goto L1
	} else {
		goto L975
	}
L973:
	;
	goto L974
L974:
	;
	v4879 = int32(0)
	v4880 = F_isatty(m, v4879)
	mBase = m.M
	if v4880 == v4879 {
		v5920 = v3792
		goto L670
	} else {
		goto L976
	}
L975:
	;
	goto L974
L976:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_96), int32(0))
	mBase = m.M
	v4886 = m.ExcPending
	if v4886 != 0 {
		goto L1
	} else {
		goto L977
	}
L977:
	;
	v4887 = F_fflush(m, v3774)
	mBase = m.M
	v4888 = m.ExcPending
	if v4888 != 0 {
		goto L1
	} else {
		goto L978
	}
L978:
	;
	v5920 = v3792
	goto L670
L979:
	;
	v5920 = v4893
	goto L670
L980:
	;
	v5920 = v4901
	goto L670
L981:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4904))) = int32(92)
	v4910 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v4904)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4904)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4904)+4)) = v4910
	v4916 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v4917 = F_makeString(m, v4916)
	mBase = m.M
	v4918 = m.ExcPending
	if v4918 != 0 {
		goto L1
	} else {
		goto L982
	}
L982:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+100)) = v4917
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+104)) = v4917
	v4924 = F_list_make1_impl(m, int32(1), v3768+int32(100))
	mBase = m.M
	v4925 = m.ExcPending
	if v4925 != 0 {
		goto L1
	} else {
		goto L983
	}
L983:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4904)+28)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4904)+20)) = v4924
	v5920 = v4904
	goto L670
L984:
	;
	v4942 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(12))))
	v4945 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(4))))
	v4946 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v4947 = int32(0)
	v4948 = m.G0
	v4950 = v4948 - int32(48)
	m.G0 = v4950
	v4953 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	if v4953 != 0 {
		goto L985
	} else {
		goto L986
	}
L985:
	;
	v4956 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4957 = m.ExcPending
	if v4957 != 0 {
		goto L1
	} else {
		goto L988
	}
L986:
	;
	goto L987
L987:
	;
	v4971 = v4934 << (uint(int32(2)) % 32)
	v4974 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	if v4974 == int32(0) {
		goto L995
	} else {
		goto L996
	}
L988:
	;
	if v4956 != 0 {
		goto L989
	} else {
		goto L990
	}
L989:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_120), int32(0))
	mBase = m.M
	v4961 = m.ExcPending
	if v4961 != 0 {
		goto L1
	} else {
		goto L992
	}
L990:
	;
	goto L991
L991:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v4969 = m.ExcPending
	if v4969 != 0 {
		goto L1
	} else {
		goto L994
	}
L992:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(528), int32(_a_F_BootstrapModeMain_121))
	mBase = m.M
	v4966 = m.ExcPending
	if v4966 != 0 {
		goto L1
	} else {
		goto L993
	}
L993:
	;
	goto L991
L994:
	;
	goto L987
L995:
	;
	v4978 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[62]))
	v4980 = F_MemoryContextAllocZero(m, v4978, int32(100))
	mBase = m.M
	v4981 = m.ExcPending
	if v4981 != 0 {
		goto L1
	} else {
		goto L998
	}
L996:
	;
	v4983 = v4974
	goto L997
L997:
	;
	v4984 = int32(0)
	base.MemoryFill(m, v4983, v4984, int32(100))
	v4987 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v4991 = F_strncpy(m, v4987+int32(4), v4942, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v4991)+63)) = uint8(v4984)
	goto L999
L998:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61]))) = v4980
	v4983 = v4980
	goto L997
L999:
	;
	v4996 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4997 = m.ExcPending
	if v4997 != 0 {
		goto L1
	} else {
		goto L1000
	}
L1000:
	;
	if v4996 != 0 {
		goto L1001
	} else {
		goto L1002
	}
L1001:
	;
	v4998 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	*(*int32)(unsafe.Add(mBase, uint32(v4950)+36)) = v4945
	*(*int32)(unsafe.Add(mBase, uint32(v4950)+32)) = v4998 + int32(4)
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_122), v4950+int32(32))
	mBase = m.M
	v5007 = m.ExcPending
	if v5007 != 0 {
		goto L1
	} else {
		goto L1004
	}
L1002:
	;
	goto L1003
L1003:
	;
	v5014 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5016 = v4934 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5014)+74)) = uint16(v5016)
	v5019 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[59]))
	if v5019 == int32(0) {
		goto L1010
	} else {
		goto L1011
	}
L1004:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(537), int32(_a_F_BootstrapModeMain_121))
	mBase = m.M
	v5012 = m.ExcPending
	if v5012 != 0 {
		goto L1
	} else {
		goto L1005
	}
L1005:
	;
	goto L1003
L1006:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v5521)+80)) = uint16(v5524)
	v5547 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5548 = *(*int32)(unsafe.Add(mBase, uint32(v5547)+96))
	if v5548 != 0 {
		goto L1088
	} else {
		goto L1089
	}
L1007:
	;
	v5521 = v5494
	v5524 = int32(1)
	goto L1006
L1008:
	;
	v5461 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[65])) = v5436
	v5464 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5465 = *(*int32)(unsafe.Add(mBase, uint32(v5436)))
	*(*int32)(unsafe.Add(mBase, uint32(v5464)+68)) = v5465
	v5467 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5436)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v5467)+72)) = uint16(v5468)
	v5470 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5436)+82)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5470)+82)) = uint8(v5471)
	v5473 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5436)+132)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5473)+83)) = uint8(v5474)
	v5476 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5436)+133)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5476)+84)) = uint8(v5477)
	v5479 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	*(*uint8)(unsafe.Add(mBase, uint32(v5479)+85)) = uint8(v5461)
	v5482 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5483 = *(*int32)(unsafe.Add(mBase, uint32(v5436)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v5482)+96)) = v5483
	v5485 = *(*int32)(unsafe.Add(mBase, uint32(v5436)+96))
	if v5485 == v5461 {
		goto L1085
	} else {
		goto L1086
	}
L1009:
	;
	v5403 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5405 = v5031 * int32(92)
	v5406 = *(*int32)(unsafe.Add(mBase, uint32(v5405)+uint32(_c_F_BootstrapModeMain[66])))
	*(*int32)(unsafe.Add(mBase, uint32(v5403)+68)) = v5406
	v5408 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5405)+uint32(_c_F_BootstrapModeMain[67]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5408)+72)) = uint16(v5409)
	v5411 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5405)+uint32(_c_F_BootstrapModeMain[68]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5411)+82)) = uint8(v5412)
	v5414 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5405)+uint32(_c_F_BootstrapModeMain[69]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5414)+83)) = uint8(v5415)
	v5417 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5405)+uint32(_c_F_BootstrapModeMain[70]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5417)+84)) = uint8(v5418)
	v5420 = int32(0)
	v5421 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	*(*uint8)(unsafe.Add(mBase, uint32(v5421)+85)) = uint8(v5420)
	v5424 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5425 = *(*int32)(unsafe.Add(mBase, uint32(v5405)+uint32(_c_F_BootstrapModeMain[71])))
	*(*int32)(unsafe.Add(mBase, uint32(v5424)+96)) = v5425
	v5427 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	if int32(1)<<(uint(v5031)%32)&int32(_a_F_BootstrapModeMain_123) != 0 {
		v5521 = v5427
		v5524 = v5420
		goto L1006
	} else {
		goto L1083
	}
L1010:
	;
	v5031 = v4947
	goto L1013
L1011:
	;
	v5134 = v5019
	v5135 = v4947
	goto L1012
L1012:
	;
	v5152 = *(*int32)(unsafe.Add(mBase, uint32(v5134)+4))
	if int32(0) < v5152 {
		goto L1038
	} else {
		goto L1039
	}
L1013:
	;
	v5051 = v5031*int32(92) + int32(_a_F_BootstrapModeMain_124)
	goto L1017
L1014:
	;
	v5134 = v5123
	v5135 = v5121
	goto L1012
L1015:
	;
	if v5089-v5090 == int32(0) {
		goto L1009
	} else {
		goto L1028
	}
L1017:
	;
	goto L1018
L1018:
	;
	v5058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4945))))
	if v5058 != 0 {
		goto L1019
	} else {
		goto L1020
	}
L1019:
	;
	v5059 = v4945
	v5060 = v5051
	v5061 = int32(64)
	v5062 = v5058
	goto L1023
L1020:
	;
	v5085 = v5051
	v5089 = int32(0)
	goto L1021
L1021:
	;
	v5090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5085))))
	goto L1015
L1022:
	;
	v5085 = v5080
	v5089 = v5082
	goto L1021
L1023:
	;
	v5064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5060))))
	if base.B2i32(v5062 != v5064)|base.B2i32(v5064 == int32(0)) != 0 {
		v5080 = v5060
		v5082 = v5062
		goto L1022
	} else {
		goto L1025
	}
L1024:
	;
	v5080 = v5074
	v5082 = int32(0)
	goto L1022
L1025:
	;
	v5070 = v5061 - int32(1)
	if v5070 == int32(0) {
		v5080 = v5060
		v5082 = v5062
		goto L1022
	} else {
		goto L1026
	}
L1026:
	;
	v5073 = int32(1)
	v5074 = v5060 + v5073
	v5075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5059)+1)))
	if v5075 != 0 {
		v5059 = v5059 + v5073
		v5060 = v5074
		v5061 = v5070
		v5062 = v5075
		goto L1023
	} else {
		goto L1027
	}
L1027:
	;
	goto L1024
L1028:
	;
	v5101 = v5031 + int32(1)
	if v5101 != int32(25) {
		v5031 = v5101
		goto L1013
	} else {
		goto L1029
	}
L1029:
	;
	v5106 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5107 = m.ExcPending
	if v5107 != 0 {
		goto L1
	} else {
		goto L1030
	}
L1030:
	;
	if v5106 != 0 {
		goto L1031
	} else {
		goto L1032
	}
L1031:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4950)+16)) = v4945
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_125), v4950+int32(16))
	mBase = m.M
	v5113 = m.ExcPending
	if v5113 != 0 {
		goto L1
	} else {
		goto L1034
	}
L1032:
	;
	goto L1033
L1033:
	;
	F_populate_typ_list(m)
	mBase = m.M
	v5120 = m.ExcPending
	if v5120 != 0 {
		goto L1
	} else {
		goto L1036
	}
L1034:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(817), int32(_a_F_BootstrapModeMain_126))
	mBase = m.M
	v5118 = m.ExcPending
	if v5118 != 0 {
		goto L1
	} else {
		goto L1035
	}
L1035:
	;
	goto L1033
L1036:
	;
	v5121 = int32(0)
	v5123 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[59]))
	if v5123 == v5121 {
		v5031 = v5121
		goto L1013
	} else {
		goto L1037
	}
L1037:
	;
	goto L1014
L1038:
	;
	v5155 = *(*int32)(unsafe.Add(mBase, uint32(v5134)+12))
	v5165 = v5135
	goto L1041
L1039:
	;
	goto L1040
L1040:
	;
	F_list_free_deep(m, v5134)
	mBase = m.M
	v5266 = m.ExcPending
	if v5266 != 0 {
		goto L1
	} else {
		goto L1058
	}
L1041:
	;
	v5185 = *(*int32)(unsafe.Add(mBase, uint32(v5155+v5165<<(uint(int32(2))%32))))
	v5187 = v5185 + int32(8)
	goto L1045
L1042:
	;
	goto L1040
L1043:
	;
	if v5225-v5226 == int32(0) {
		v5436 = v5185
		goto L1008
	} else {
		goto L1056
	}
L1045:
	;
	goto L1046
L1046:
	;
	v5194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5187))))
	if v5194 != 0 {
		goto L1047
	} else {
		goto L1048
	}
L1047:
	;
	v5195 = v5187
	v5196 = v4945
	v5197 = int32(64)
	v5198 = v5194
	goto L1051
L1048:
	;
	v5221 = v4945
	v5225 = int32(0)
	goto L1049
L1049:
	;
	v5226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5221))))
	goto L1043
L1050:
	;
	v5221 = v5216
	v5225 = v5218
	goto L1049
L1051:
	;
	v5200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5196))))
	if base.B2i32(v5198 != v5200)|base.B2i32(v5200 == int32(0)) != 0 {
		v5216 = v5196
		v5218 = v5198
		goto L1050
	} else {
		goto L1053
	}
L1052:
	;
	v5216 = v5210
	v5218 = int32(0)
	goto L1050
L1053:
	;
	v5206 = v5197 - int32(1)
	if v5206 == int32(0) {
		v5216 = v5196
		v5218 = v5198
		goto L1050
	} else {
		goto L1054
	}
L1054:
	;
	v5209 = int32(1)
	v5210 = v5196 + v5209
	v5211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5195)+1)))
	if v5211 != 0 {
		v5195 = v5195 + v5209
		v5196 = v5210
		v5197 = v5206
		v5198 = v5211
		goto L1051
	} else {
		goto L1055
	}
L1055:
	;
	goto L1052
L1056:
	;
	v5237 = v5165 + int32(1)
	if v5237 != v5152 {
		v5165 = v5237
		goto L1041
	} else {
		goto L1057
	}
L1057:
	;
	goto L1042
L1058:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[59])) = int32(0)
	F_populate_typ_list(m)
	mBase = m.M
	v5271 = m.ExcPending
	if v5271 != 0 {
		goto L1
	} else {
		goto L1059
	}
L1059:
	;
	v5273 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[59]))
	if v5273 == int32(0) {
		goto L1060
	} else {
		goto L1061
	}
L1060:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5393 = m.ExcPending
	if v5393 != 0 {
		goto L1
	} else {
		goto L1080
	}
L1061:
	;
	v5276 = *(*int32)(unsafe.Add(mBase, uint32(v5273)+4))
	if v5276 <= int32(0) {
		goto L1060
	} else {
		goto L1062
	}
L1062:
	;
	v5279 = *(*int32)(unsafe.Add(mBase, uint32(v5273)+12))
	v5290 = int32(0)
	goto L1063
L1063:
	;
	v5310 = *(*int32)(unsafe.Add(mBase, uint32(v5279+v5290<<(uint(int32(2))%32))))
	v5312 = v5310 + int32(8)
	goto L1067
L1064:
	;
	goto L1060
L1065:
	;
	if v5350-v5351 == int32(0) {
		v5436 = v5310
		goto L1008
	} else {
		goto L1078
	}
L1067:
	;
	goto L1068
L1068:
	;
	v5319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5312))))
	if v5319 != 0 {
		goto L1069
	} else {
		goto L1070
	}
L1069:
	;
	v5320 = v5312
	v5321 = v4945
	v5322 = int32(64)
	v5323 = v5319
	goto L1073
L1070:
	;
	v5346 = v4945
	v5350 = int32(0)
	goto L1071
L1071:
	;
	v5351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5346))))
	goto L1065
L1072:
	;
	v5346 = v5341
	v5350 = v5343
	goto L1071
L1073:
	;
	v5325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5321))))
	if base.B2i32(v5323 != v5325)|base.B2i32(v5325 == int32(0)) != 0 {
		v5341 = v5321
		v5343 = v5323
		goto L1072
	} else {
		goto L1075
	}
L1074:
	;
	v5341 = v5335
	v5343 = int32(0)
	goto L1072
L1075:
	;
	v5331 = v5322 - int32(1)
	if v5331 == int32(0) {
		v5341 = v5321
		v5343 = v5323
		goto L1072
	} else {
		goto L1076
	}
L1076:
	;
	v5334 = int32(1)
	v5335 = v5321 + v5334
	v5336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5320)+1)))
	if v5336 != 0 {
		v5320 = v5320 + v5334
		v5321 = v5335
		v5322 = v5331
		v5323 = v5336
		goto L1073
	} else {
		goto L1077
	}
L1077:
	;
	goto L1074
L1078:
	;
	v5362 = v5290 + int32(1)
	if v5276 != v5362 {
		v5290 = v5362
		goto L1063
	} else {
		goto L1079
	}
L1079:
	;
	goto L1064
L1080:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4950))) = v4945
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_127), v4950)
	mBase = m.M
	v5397 = m.ExcPending
	if v5397 != 0 {
		goto L1
	} else {
		goto L1081
	}
L1081:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(821), int32(_a_F_BootstrapModeMain_126))
	mBase = m.M
	v5402 = m.ExcPending
	if v5402 != 0 {
		goto L1
	} else {
		goto L1082
	}
L1082:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1083:
	;
	v5432 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5427)+72)))
	if int32(0) <= v5432 {
		v5521 = v5427
		v5524 = v5420
		goto L1006
	} else {
		goto L1084
	}
L1084:
	;
	v5494 = v5427
	goto L1007
L1085:
	;
	v5492 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5521 = v5492
	v5524 = v5461
	goto L1006
L1086:
	;
	v5488 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5436)+80)))
	if int32(0) <= v5488 {
		goto L1085
	} else {
		goto L1087
	}
L1087:
	;
	v5491 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5494 = v5491
	goto L1007
L1088:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5547)+96)) = int32(950)
	v5551 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	v5552 = v5551
	goto L1090
L1089:
	;
	v5552 = v5547
	goto L1090
L1090:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5552)+76)) = int32(-1)
	v5555 = int32(1)
	v5556 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	*(*uint8)(unsafe.Add(mBase, uint32(v5556)+92)) = uint8(v5555)
	v5559 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+uint32(_c_F_BootstrapModeMain[61])))
	switch v4946 - int32(2) {
	case 0:
		goto L1094
	case 1:
		v5638 = v5555
		goto L1092
	default:
		goto L1093
	}
L1091:
	;
	m.G0 = v4950 + int32(48)
	v5920 = v3792
	goto L670
L1092:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v5559)+86)) = uint8(v5638)
	goto L1091
L1093:
	;
	v5563 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5559)+72)))
	if v5563 <= int32(0) {
		goto L1091
	} else {
		goto L1095
	}
L1094:
	;
	v5638 = int32(0)
	goto L1092
L1095:
	;
	v5566 = int32(0)
	if v4934 <= v5566 {
		v5616 = v5566
		goto L1096
	} else {
		goto L1097
	}
L1096:
	;
	if v5616 != v4934 {
		goto L1091
	} else {
		goto L1103
	}
L1097:
	;
	v5578 = v5566
	goto L1098
L1098:
	;
	v5597 = *(*int32)(unsafe.Add(mBase, uint32(v5578<<(uint(int32(2))%32))+uint32(_c_F_BootstrapModeMain[61])))
	v5598 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5597)+72)))
	if v5598 <= int32(0) {
		v5616 = v5578
		goto L1096
	} else {
		goto L1100
	}
L1099:
	;
	v5638 = v5555
	goto L1092
L1100:
	;
	v5601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5597)+86)))
	if v5601 != int32(1) {
		v5616 = v5578
		goto L1096
	} else {
		goto L1101
	}
L1101:
	;
	v5605 = v5578 + int32(1)
	if v5605 != v4934 {
		v5578 = v5605
		goto L1098
	} else {
		goto L1102
	}
L1102:
	;
	goto L1099
L1103:
	;
	v5638 = v5555
	goto L1092
L1104:
	;
	v5920 = base.I32_wrap_i64(v5696)
	goto L670
L1105:
	;
	if v5711 != 0 {
		goto L1106
	} else {
		goto L1107
	}
L1106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+20)) = v5698
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+16)) = v5701
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_128), v5707+int32(16))
	mBase = m.M
	v5719 = m.ExcPending
	if v5719 != 0 {
		goto L1
	} else {
		goto L1109
	}
L1107:
	;
	goto L1108
L1108:
	;
	v5726 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	v5727 = *(*int32)(unsafe.Add(mBase, uint32(v5726)+52))
	v5728 = *(*int32)(unsafe.Add(mBase, uint32(v5727)))
	v5735 = *(*int32)(unsafe.Add(mBase, uint32(v5727+v5728<<(uint(int32(4))%32)+v5701*int32(100))+88))
	F_boot_get_type_io_data(m, v5735, v5707+int32(46), v5707+int32(45), v5707+int32(44), v5707+int32(43), v5707+int32(36), v5707+int32(32), v5707+int32(28))
	mBase = m.M
	v5751 = m.ExcPending
	if v5751 != 0 {
		goto L1
	} else {
		goto L1111
	}
L1109:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(670), int32(_a_F_BootstrapModeMain_129))
	mBase = m.M
	v5724 = m.ExcPending
	if v5724 != 0 {
		goto L1
	} else {
		goto L1110
	}
L1110:
	;
	goto L1108
L1111:
	;
	v5753 = v5701 << (uint(int32(2)) % 32)
	v5756 = *(*int32)(unsafe.Add(mBase, uint32(v5707)+32))
	v5757 = *(*int32)(unsafe.Add(mBase, uint32(v5707)+36))
	v5759 = F_OidInputFunctionCall(m, v5756, v5698, v5757, int32(-1))
	mBase = m.M
	v5760 = m.ExcPending
	if v5760 != 0 {
		goto L1
	} else {
		goto L1112
	}
L1112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5753)+uint32(_c_F_BootstrapModeMain[72]))) = v5759
	v5764 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5765 = m.ExcPending
	if v5765 != 0 {
		goto L1
	} else {
		goto L1113
	}
L1113:
	;
	if v5764 != 0 {
		goto L1114
	} else {
		goto L1115
	}
L1114:
	;
	v5766 = *(*int32)(unsafe.Add(mBase, uint32(v5707)+28))
	v5767 = *(*int32)(unsafe.Add(mBase, uint32(v5753)+uint32(_c_F_BootstrapModeMain[72])))
	v5768 = F_OidOutputFunctionCall(m, v5766, v5767)
	mBase = m.M
	v5769 = m.ExcPending
	if v5769 != 0 {
		goto L1
	} else {
		goto L1117
	}
L1115:
	;
	goto L1116
L1116:
	;
	m.G0 = v5707 + int32(48)
	v5920 = v3792
	goto L670
L1117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5707))) = v5768
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_130), v5707)
	mBase = m.M
	v5773 = m.ExcPending
	if v5773 != 0 {
		goto L1
	} else {
		goto L1118
	}
L1118:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(687), int32(_a_F_BootstrapModeMain_129))
	mBase = m.M
	v5778 = m.ExcPending
	if v5778 != 0 {
		goto L1
	} else {
		goto L1119
	}
L1119:
	;
	goto L1116
L1120:
	;
	if v5794 != 0 {
		goto L1121
	} else {
		goto L1122
	}
L1121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5790)+16)) = v5784
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_131), v5790+int32(16))
	mBase = m.M
	v5801 = m.ExcPending
	if v5801 != 0 {
		goto L1
	} else {
		goto L1124
	}
L1122:
	;
	goto L1123
L1123:
	;
	v5808 = v5784 * int32(100)
	v5810 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	v5811 = *(*int32)(unsafe.Add(mBase, uint32(v5810)+52))
	v5812 = *(*int32)(unsafe.Add(mBase, uint32(v5811)))
	v5817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5808+(v5811+v5812<<(uint(int32(4))%32)))+106)))
	if v5817 == int32(1) {
		goto L1126
	} else {
		goto L1127
	}
L1124:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(697), int32(_a_F_BootstrapModeMain_132))
	mBase = m.M
	v5806 = m.ExcPending
	if v5806 != 0 {
		goto L1
	} else {
		goto L1125
	}
L1125:
	;
	goto L1123
L1126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5823 = m.ExcPending
	if v5823 != 0 {
		goto L1
	} else {
		goto L1129
	}
L1127:
	;
	goto L1128
L1128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5784<<(uint(int32(2))%32))+uint32(_c_F_BootstrapModeMain[72]))) = int32(0)
	v5855 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v5784)+uint32(_c_F_BootstrapModeMain[36]))) = uint8(v5855)
	m.G0 = v5790 + int32(32)
	v5920 = v3792
	goto L670
L1129:
	;
	v5825 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	v5826 = *(*int32)(unsafe.Add(mBase, uint32(v5825)+52))
	v5827 = *(*int32)(unsafe.Add(mBase, uint32(v5826)))
	v5828 = *(*int32)(unsafe.Add(mBase, uint32(v5825)+48))
	v5829 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v5790)+4)) = v5828 + v5829
	*(*int32)(unsafe.Add(mBase, uint32(v5790))) = v5826 + v5827<<(uint(v5829)%32) + v5808 + int32(24)
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_133), v5790)
	mBase = m.M
	v5841 = m.ExcPending
	if v5841 != 0 {
		goto L1
	} else {
		goto L1130
	}
L1130:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(703), int32(_a_F_BootstrapModeMain_132))
	mBase = m.M
	v5846 = m.ExcPending
	if v5846 != 0 {
		goto L1
	} else {
		goto L1131
	}
L1131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1132:
	;
	v5920 = v5862
	goto L670
L1133:
	;
	v5920 = v5865
	goto L670
L1134:
	;
	v5920 = v5868
	goto L670
L1135:
	;
	v5920 = v5871
	goto L670
L1136:
	;
	v5920 = v5874
	goto L670
L1137:
	;
	v5920 = v5877
	goto L670
L1138:
	;
	v5920 = v5880
	goto L670
L1139:
	;
	v5920 = v5883
	goto L670
L1140:
	;
	v5920 = v5886
	goto L670
L1141:
	;
	v5920 = v5889
	goto L670
L1142:
	;
	v5920 = v5892
	goto L670
L1143:
	;
	v5920 = v5895
	goto L670
L1144:
	;
	v5920 = v5898
	goto L670
L1145:
	;
	v5920 = v5901
	goto L670
L1146:
	;
	v5920 = v5904
	goto L670
L1147:
	;
	v5920 = v5907
	goto L670
L1148:
	;
	v5920 = v5910
	goto L670
L1149:
	;
	v5920 = v5913
	goto L670
L1150:
	;
	v5920 = v5916
	goto L670
L1151:
	;
	v5970 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5954)+uint32(_c_F_BootstrapModeMain[73]))))
	v5974 = v3761
	v5976 = v3763
	v5978 = v5949
	v5981 = v3768
	v5982 = v3769
	v5984 = v5970
	v5985 = v5950
	v5987 = v3774
	v5988 = v3775
	v5989 = v3776
	v5993 = v3780
	goto L289
L1152:
	;
	v5961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5958)+uint32(_c_F_BootstrapModeMain[50]))))
	if v5961 != v5951&int32(255) {
		goto L1151
	} else {
		goto L1153
	}
L1153:
	;
	v5967 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5958)+uint32(_c_F_BootstrapModeMain[51]))))
	v5974 = v3761
	v5976 = v3763
	v5978 = v5949
	v5981 = v3768
	v5982 = v3769
	v5984 = v5967
	v5985 = v5950
	v5987 = v3774
	v5988 = v3775
	v5989 = v3776
	v5993 = v3780
	goto L289
L1154:
	;
	if v3732 == int32(0) {
		v6120 = v3737
		v6121 = v3738
		goto L257
	} else {
		goto L1157
	}
L1155:
	;
	F_boot_yyerror(m, v3730, int32(_a_F_BootstrapModeMain_134))
	mBase = m.M
	v6001 = m.ExcPending
	if v6001 != 0 {
		goto L1
	} else {
		goto L1156
	}
L1156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1157:
	;
	v6014 = v3737
	v6015 = v3738
	v6018 = v3741
	goto L262
L1158:
	;
	if v6015 == v6044 {
		v6120 = v6014
		v6121 = v6015
		goto L257
	} else {
		goto L1160
	}
L1160:
	;
	v6044 = v6044 - int32(1)
	goto L1158
L1161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1162:
	;
	v6069 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+32)) = v6069
	v6072 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54]))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+36)) = v6072
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_135), v3768+int32(32))
	mBase = m.M
	v6078 = m.ExcPending
	if v6078 != 0 {
		goto L1
	} else {
		goto L1163
	}
L1163:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_101), int32(265), int32(_a_F_BootstrapModeMain_102))
	mBase = m.M
	v6083 = m.ExcPending
	if v6083 != 0 {
		goto L1
	} else {
		goto L1164
	}
L1164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1165:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_136), int32(0))
	mBase = m.M
	v6091 = m.ExcPending
	if v6091 != 0 {
		goto L1
	} else {
		goto L1166
	}
L1166:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_101), int32(267), int32(_a_F_BootstrapModeMain_102))
	mBase = m.M
	v6096 = m.ExcPending
	if v6096 != 0 {
		goto L1
	} else {
		goto L1167
	}
L1167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1168:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_137), int32(0))
	mBase = m.M
	v6104 = m.ExcPending
	if v6104 != 0 {
		goto L1
	} else {
		goto L1169
	}
L1169:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_101), int32(446), int32(_a_F_BootstrapModeMain_102))
	mBase = m.M
	v6109 = m.ExcPending
	if v6109 != 0 {
		goto L1
	} else {
		goto L1170
	}
L1170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1171:
	;
	F_pfree(m, v6136)
	mBase = m.M
	v6166 = m.ExcPending
	if v6166 != 0 {
		goto L1
	} else {
		goto L1174
	}
L1172:
	;
	goto L1173
L1173:
	;
	m.G0 = v6146 + int32(1136)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v6171 = m.ExcPending
	if v6171 != 0 {
		goto L1
	} else {
		goto L1175
	}
L1174:
	;
	goto L1173
L1175:
	;
	v6173 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[9]))
	v6177 = F_LWLockAcquire(m, v6173+int32(3200), int32(0))
	mBase = m.M
	v6178 = m.ExcPending
	if v6178 != 0 {
		goto L1
	} else {
		goto L1176
	}
L1176:
	;
	v6180 = int32(0)
	F_write_relmap_file(m, int32(_a_F_BootstrapModeMain_138), v6180, v6180, v6180, v6180, int32(1664), int32(_a_F_BootstrapModeMain_139))
	mBase = m.M
	v6187 = m.ExcPending
	if v6187 != 0 {
		goto L1
	} else {
		goto L1177
	}
L1177:
	;
	v6189 = int32(0)
	v6193 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[74]))
	v6195 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	v6197 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	F_write_relmap_file(m, int32(_a_F_BootstrapModeMain_140), v6189, v6189, v6189, v6193, v6195, v6197)
	mBase = m.M
	v6199 = m.ExcPending
	if v6199 != 0 {
		goto L1
	} else {
		goto L1178
	}
L1178:
	;
	v6201 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[9]))
	F_LWLockRelease(m, v6201+int32(3200))
	mBase = m.M
	v6205 = m.ExcPending
	if v6205 != 0 {
		goto L1
	} else {
		goto L1179
	}
L1179:
	;
	v6207 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	if v6207 != 0 {
		goto L1180
	} else {
		goto L1181
	}
L1180:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v6210 = m.ExcPending
	if v6210 != 0 {
		goto L1
	} else {
		goto L1183
	}
L1181:
	;
	goto L1182
L1182:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v6213 = m.ExcPending
	if v6213 != 0 {
		goto L1
	} else {
		goto L1184
	}
L1183:
	;
	goto L1182
L1184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1185:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v6220 = m.ExcPending
	if v6220 != 0 {
		goto L1
	} else {
		goto L1186
	}
L1186:
	;
	v6222 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v6222
	F_errmsg(m, int32(_a_F_BootstrapModeMain_141), v29-int32(-64))
	mBase = m.M
	v6228 = m.ExcPending
	if v6228 != 0 {
		goto L1
	} else {
		goto L1187
	}
L1187:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(239), int32(_a_F_BootstrapModeMain_17))
	mBase = m.M
	v6233 = m.ExcPending
	if v6233 != 0 {
		goto L1
	} else {
		goto L1188
	}
L1188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1189:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(254), int32(_a_F_BootstrapModeMain_17))
	mBase = m.M
	v6244 = m.ExcPending
	if v6244 != 0 {
		goto L1
	} else {
		goto L1190
	}
L1190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1191:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v6251 = m.ExcPending
	if v6251 != 0 {
		goto L1
	} else {
		goto L1192
	}
L1192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1193:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v6260 = m.ExcPending
	if v6260 != 0 {
		goto L1
	} else {
		goto L1194
	}
L1194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1197:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_142), int32(0))
	mBase = m.M
	v6277 = m.ExcPending
	if v6277 != 0 {
		goto L1
	} else {
		goto L1198
	}
L1198:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_16), int32(383), int32(_a_F_BootstrapModeMain_17))
	mBase = m.M
	v6282 = m.ExcPending
	if v6282 != 0 {
		goto L1
	} else {
		goto L1199
	}
L1199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
