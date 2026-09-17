package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_transformFromClauseItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v193 int32
	_ = v193
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v338 int32
	_ = v338
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v383 int32
	_ = v383
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
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
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
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
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 float64
	_ = v680
	var v681 int32
	_ = v681
	var v687 int32
	_ = v687
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v830 int32
	_ = v830
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v881 int32
	_ = v881
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v990 int32
	_ = v990
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1061 int32
	_ = v1061
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1102 int32
	_ = v1102
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1161 int32
	_ = v1161
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1203 int32
	_ = v1203
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1290 int32
	_ = v1290
	var v1315 int32
	_ = v1315
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1346 int32
	_ = v1346
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1392 int32
	_ = v1392
	var v1397 int32
	_ = v1397
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1469 int32
	_ = v1469
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1518 int32
	_ = v1518
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1610 int32
	_ = v1610
	var v1619 int32
	_ = v1619
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1681 int32
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1738 int32
	_ = v1738
	var v1747 int32
	_ = v1747
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1761 int32
	_ = v1761
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1770 int32
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1784 int32
	_ = v1784
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1797 int32
	_ = v1797
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1819 int32
	_ = v1819
	var v1830 int32
	_ = v1830
	var v1839 int32
	_ = v1839
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1904 int32
	_ = v1904
	var v1909 int32
	_ = v1909
	var v1920 int32
	_ = v1920
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1939 int32
	_ = v1939
	var v1944 int32
	_ = v1944
	var v1955 int32
	_ = v1955
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1972 int32
	_ = v1972
	var v1976 int32
	_ = v1976
	var v1979 int32
	_ = v1979
	var v1985 int32
	_ = v1985
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2035 int32
	_ = v2035
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2051 int32
	_ = v2051
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2061 int32
	_ = v2061
	var v2075 int32
	_ = v2075
	var v2087 int32
	_ = v2087
	var v2091 int32
	_ = v2091
	var v2093 int32
	_ = v2093
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2114 int32
	_ = v2114
	var v2120 int32
	_ = v2120
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2127 int32
	_ = v2127
	var v2130 int32
	_ = v2130
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2140 int64
	_ = v2140
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2155 int32
	_ = v2155
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2164 int32
	_ = v2164
	var v2167 int32
	_ = v2167
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2176 int32
	_ = v2176
	var v2181 int32
	_ = v2181
	var v2185 int32
	_ = v2185
	var v2188 int32
	_ = v2188
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2200 int32
	_ = v2200
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2213 int32
	_ = v2213
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2219 int32
	_ = v2219
	var v2221 int32
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2247 int32
	_ = v2247
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2265 int32
	_ = v2265
	var v2273 int32
	_ = v2273
	var v2289 int32
	_ = v2289
	var v2295 int32
	_ = v2295
	var v2306 int32
	_ = v2306
	var v2312 int32
	_ = v2312
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
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
	var v2333 int32
	_ = v2333
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2374 int32
	_ = v2374
	var v2381 int32
	_ = v2381
	var v2385 int32
	_ = v2385
	var v2390 int32
	_ = v2390
	var v2395 int32
	_ = v2395
	var v2399 int32
	_ = v2399
	var v2404 int32
	_ = v2404
	var v2408 int32
	_ = v2408
	var v2414 int32
	_ = v2414
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2425 int32
	_ = v2425
	var v2457 int32
	_ = v2457
	var v2481 int32
	_ = v2481
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2508 int32
	_ = v2508
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2514 int32
	_ = v2514
	var v2519 int32
	_ = v2519
	var v2526 int32
	_ = v2526
	var v2550 int32
	_ = v2550
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2557 int32
	_ = v2557
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
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2589 int32
	_ = v2589
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2603 int32
	_ = v2603
	var v2605 int32
	_ = v2605
	var v2608 int32
	_ = v2608
	var v2609 int32
	_ = v2609
	var v2611 int32
	_ = v2611
	var v2643 int32
	_ = v2643
	var v2647 int32
	_ = v2647
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2684 int32
	_ = v2684
	var v2687 int32
	_ = v2687
	var v2691 int32
	_ = v2691
	var v2693 int32
	_ = v2693
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2702 int32
	_ = v2702
	var v2704 int32
	_ = v2704
	var v2713 int32
	_ = v2713
	var v2738 int32
	_ = v2738
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2743 int32
	_ = v2743
	var v2747 int32
	_ = v2747
	var v2750 int32
	_ = v2750
	var v2774 int32
	_ = v2774
	var v2776 int32
	_ = v2776
	var v2780 int32
	_ = v2780
	var v2782 int32
	_ = v2782
	var v2783 int32
	_ = v2783
	var v2784 int32
	_ = v2784
	var v2786 int32
	_ = v2786
	var v2818 int32
	_ = v2818
	var v2822 int32
	_ = v2822
	var v2827 int32
	_ = v2827
	var v2830 int32
	_ = v2830
	var v2854 int32
	_ = v2854
	var v2863 int32
	_ = v2863
	var v2864 int32
	_ = v2864
	var v2869 int32
	_ = v2869
	var v2897 int32
	_ = v2897
	var v2898 int32
	_ = v2898
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2903 int32
	_ = v2903
	var v2905 int32
	_ = v2905
	var v2907 int32
	_ = v2907
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2910 int32
	_ = v2910
	var v2913 int32
	_ = v2913
	var v2914 int32
	_ = v2914
	var v2921 int32
	_ = v2921
	var v2948 int32
	_ = v2948
	var v2954 int32
	_ = v2954
	var v2955 int32
	_ = v2955
	var v2960 int32
	_ = v2960
	var v2962 int32
	_ = v2962
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2969 int32
	_ = v2969
	var v2971 int32
	_ = v2971
	var v2979 int32
	_ = v2979
	var v3013 int32
	_ = v3013
	var v3014 int32
	_ = v3014
	var v3015 int32
	_ = v3015
	var v3030 int32
	_ = v3030
	var v3031 int32
	_ = v3031
	var v3033 int32
	_ = v3033
	var v3038 int32
	_ = v3038
	var v3042 int32
	_ = v3042
	var v3045 int32
	_ = v3045
	var v3049 int32
	_ = v3049
	var v3050 int32
	_ = v3050
	var v3052 int32
	_ = v3052
	var v3057 int32
	_ = v3057
	var v3061 int32
	_ = v3061
	var v3064 int32
	_ = v3064
	var v3071 int32
	_ = v3071
	var v3072 int32
	_ = v3072
	var v3074 int32
	_ = v3074
	var v3079 int32
	_ = v3079
	var v3083 int32
	_ = v3083
	var v3086 int32
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3096 int32
	_ = v3096
	var v3097 int32
	_ = v3097
	var v3099 int32
	_ = v3099
	var v3104 int32
	_ = v3104
	var v3108 int32
	_ = v3108
	var v3111 int32
	_ = v3111
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3119 int32
	_ = v3119
	var v3124 int32
	_ = v3124
	var v3128 int32
	_ = v3128
	var v3131 int32
	_ = v3131
	var v3137 int32
	_ = v3137
	var v3138 int32
	_ = v3138
	var v3140 int32
	_ = v3140
	var v3145 int32
	_ = v3145
	var v3152 int32
	_ = v3152
	var v3153 int32
	_ = v3153
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3160 int32
	_ = v3160
	var v3162 int32
	_ = v3162
	var v3164 int32
	_ = v3164
	var v3166 int32
	_ = v3166
	var v3170 int32
	_ = v3170
	var v3173 int32
	_ = v3173
	var v3181 int32
	_ = v3181
	var v3183 int32
	_ = v3183
	var v3191 int32
	_ = v3191
	var v3196 int32
	_ = v3196
	var v3197 int32
	_ = v3197
	var v3199 int32
	_ = v3199
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3202 int32
	_ = v3202
	var v3205 int32
	_ = v3205
	var v3211 int32
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3215 int32
	_ = v3215
	var v3216 int32
	_ = v3216
	var v3218 int32
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3231 int32
	_ = v3231
	var v3233 int32
	_ = v3233
	var v3235 int32
	_ = v3235
	var v3239 int32
	_ = v3239
	var v3241 int32
	_ = v3241
	var v3244 int32
	_ = v3244
	var v3245 int32
	_ = v3245
	var v3249 int32
	_ = v3249
	var v3250 int32
	_ = v3250
	var v3251 int32
	_ = v3251
	var v3252 int32
	_ = v3252
	var v3254 int32
	_ = v3254
	var v3255 int32
	_ = v3255
	var v3256 int32
	_ = v3256
	var v3257 int32
	_ = v3257
	var v3261 int32
	_ = v3261
	var v3263 int32
	_ = v3263
	var v3266 int32
	_ = v3266
	var v3270 int32
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3275 int32
	_ = v3275
	var v3282 int32
	_ = v3282
	var v3285 int32
	_ = v3285
	var v3292 int32
	_ = v3292
	var v3296 int32
	_ = v3296
	var v3297 int32
	_ = v3297
	var v3298 int32
	_ = v3298
	var v3300 int32
	_ = v3300
	var v3305 int32
	_ = v3305
	var v3307 int32
	_ = v3307
	var v3308 int32
	_ = v3308
	var v3311 int32
	_ = v3311
	var v3313 int32
	_ = v3313
	var v3315 int32
	_ = v3315
	var v3316 int32
	_ = v3316
	var v3319 int32
	_ = v3319
	var v3320 int32
	_ = v3320
	var v3323 int32
	_ = v3323
	var v3324 int32
	_ = v3324
	var v3326 int32
	_ = v3326
	var v3327 int32
	_ = v3327
	var v3330 int32
	_ = v3330
	var v3331 int32
	_ = v3331
	var v3334 int32
	_ = v3334
	var v3337 int32
	_ = v3337
	var v3338 int32
	_ = v3338
	var v3342 int32
	_ = v3342
	var v3343 int32
	_ = v3343
	var v3344 int32
	_ = v3344
	var v3345 int32
	_ = v3345
	var v3348 int32
	_ = v3348
	var v3361 int32
	_ = v3361
	var v3382 int32
	_ = v3382
	var v3384 int32
	_ = v3384
	var v3385 int32
	_ = v3385
	var v3387 int32
	_ = v3387
	var v3388 int32
	_ = v3388
	var v3389 int32
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3391 int32
	_ = v3391
	var v3392 int32
	_ = v3392
	var v3393 int32
	_ = v3393
	var v3394 int32
	_ = v3394
	var v3396 int32
	_ = v3396
	var v3397 int32
	_ = v3397
	var v3402 int32
	_ = v3402
	var v3406 int32
	_ = v3406
	var v3407 int32
	_ = v3407
	var v3415 int32
	_ = v3415
	var v3416 int32
	_ = v3416
	var v3418 int32
	_ = v3418
	var v3419 int32
	_ = v3419
	var v3420 int32
	_ = v3420
	var v3421 int32
	_ = v3421
	var v3423 int32
	_ = v3423
	var v3424 int32
	_ = v3424
	var v3425 int32
	_ = v3425
	var v3426 int32
	_ = v3426
	var v3428 int32
	_ = v3428
	var v3429 int32
	_ = v3429
	var v3430 int32
	_ = v3430
	var v3431 int32
	_ = v3431
	var v3432 int32
	_ = v3432
	var v3433 int32
	_ = v3433
	var v3435 int32
	_ = v3435
	var v3437 int32
	_ = v3437
	var v3439 int32
	_ = v3439
	var v3440 int32
	_ = v3440
	var v3443 int32
	_ = v3443
	var v3444 int32
	_ = v3444
	var v3446 int32
	_ = v3446
	var v3447 int32
	_ = v3447
	var v3448 int32
	_ = v3448
	var v3450 int32
	_ = v3450
	var v3451 int32
	_ = v3451
	var v3452 int32
	_ = v3452
	var v3453 int32
	_ = v3453
	var v3455 int32
	_ = v3455
	var v3456 int32
	_ = v3456
	var v3458 int32
	_ = v3458
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3461 int32
	_ = v3461
	var v3462 int32
	_ = v3462
	var v3464 int32
	_ = v3464
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3468 int32
	_ = v3468
	var v3471 int32
	_ = v3471
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3475 int32
	_ = v3475
	var v3476 int32
	_ = v3476
	var v3483 int32
	_ = v3483
	var v3513 int32
	_ = v3513
	var v3516 int32
	_ = v3516
	var v3519 int32
	_ = v3519
	var v3522 int32
	_ = v3522
	var v3523 int32
	_ = v3523
	var v3526 int32
	_ = v3526
	var v3527 int32
	_ = v3527
	var v3530 int32
	_ = v3530
	var v3537 int32
	_ = v3537
	var v3538 int32
	_ = v3538
	var v3541 int32
	_ = v3541
	var v3546 int32
	_ = v3546
	var v3549 int32
	_ = v3549
	var v3550 int32
	_ = v3550
	var v3556 int32
	_ = v3556
	var v3557 int32
	_ = v3557
	var v3559 int32
	_ = v3559
	var v3564 int32
	_ = v3564
	var v3568 int32
	_ = v3568
	var v3571 int32
	_ = v3571
	var v3575 int32
	_ = v3575
	var v3576 int32
	_ = v3576
	var v3578 int32
	_ = v3578
	var v3583 int32
	_ = v3583
	var v3587 int32
	_ = v3587
	var v3590 int32
	_ = v3590
	var v3591 int32
	_ = v3591
	var v3597 int32
	_ = v3597
	var v3598 int32
	_ = v3598
	var v3600 int32
	_ = v3600
	var v3605 int32
	_ = v3605
	var v3640 int32
	_ = v3640
	var v3641 int32
	_ = v3641
	var v3646 int32
	_ = v3646
	var v3647 int32
	_ = v3647
	var v3651 int32
	_ = v3651
	var v3656 int32
	_ = v3656
	var v3660 int32
	_ = v3660
	var v3663 int32
	_ = v3663
	var v3667 int32
	_ = v3667
	var v3671 int32
	_ = v3671
	var v3672 int32
	_ = v3672
	var v3673 int32
	_ = v3673
	var v3675 int32
	_ = v3675
	var v3680 int32
	_ = v3680
	var v3684 int32
	_ = v3684
	var v3688 int32
	_ = v3688
	var v3689 int32
	_ = v3689
	var v3690 int32
	_ = v3690
	var v3692 int32
	_ = v3692
	var v3697 int32
	_ = v3697
	var v3701 int32
	_ = v3701
	var v3704 int32
	_ = v3704
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3710 int32
	_ = v3710
	var v3712 int32
	_ = v3712
	var v3717 int32
	_ = v3717
	var v3721 int32
	_ = v3721
	var v3724 int32
	_ = v3724
	var v3728 int32
	_ = v3728
	var v3729 int32
	_ = v3729
	var v3730 int32
	_ = v3730
	var v3732 int32
	_ = v3732
	var v3737 int32
	_ = v3737
	var v3741 int32
	_ = v3741
	var v3744 int32
	_ = v3744
	var v3748 int32
	_ = v3748
	var v3749 int32
	_ = v3749
	var v3750 int32
	_ = v3750
	var v3752 int32
	_ = v3752
	var v3757 int32
	_ = v3757
	var v3761 int32
	_ = v3761
	var v3765 int32
	_ = v3765
	var v3770 int32
	_ = v3770
	var v3771 int32
	_ = v3771
	var v3776 int32
	_ = v3776
	var v3777 int32
	_ = v3777
	var v3779 int32
	_ = v3779
	var v3782 int32
	_ = v3782
	var v3785 int32
	_ = v3785
	var v3793 int32
	_ = v3793
	var v3820 int32
	_ = v3820
	var v3824 int32
	_ = v3824
	var v3826 int32
	_ = v3826
	var v3829 int32
	_ = v3829
	var v3830 int32
	_ = v3830
	var v3863 int32
	_ = v3863
	var v3864 int32
	_ = v3864
	var v3866 int32
	_ = v3866
	var v3867 int32
	_ = v3867
	var v3868 int32
	_ = v3868
	var v3870 int32
	_ = v3870
	var v3875 int32
	_ = v3875
	var v3876 int32
	_ = v3876
	var v3878 int32
	_ = v3878
	var v3879 int32
	_ = v3879
	var v3886 int32
	_ = v3886
	var v3889 int32
	_ = v3889
	var v3891 int32
	_ = v3891
	var v3893 int32
	_ = v3893
	var v3894 int32
	_ = v3894
	var v3895 int32
	_ = v3895
	var v3896 int32
	_ = v3896
	var v3897 int32
	_ = v3897
	var v3898 int32
	_ = v3898
	var v3899 int32
	_ = v3899
	var v3900 int32
	_ = v3900
	var v3901 int32
	_ = v3901
	var v3902 int32
	_ = v3902
	var v3903 int32
	_ = v3903
	var v3904 int32
	_ = v3904
	var v3909 int32
	_ = v3909
	var v3918 int32
	_ = v3918
	var v3923 int32
	_ = v3923
	var v3943 int32
	_ = v3943
	var v3947 int32
	_ = v3947
	var v3948 int32
	_ = v3948
	var v3949 int32
	_ = v3949
	var v3950 int32
	_ = v3950
	var v3955 int32
	_ = v3955
	var v3958 int32
	_ = v3958
	var v3961 int32
	_ = v3961
	var v3962 int32
	_ = v3962
	var v3968 int32
	_ = v3968
	var v3998 int32
	_ = v3998
	var v3999 int32
	_ = v3999
	var v4002 int32
	_ = v4002
	var v4005 int32
	_ = v4005
	var v4008 int32
	_ = v4008
	var v4009 int32
	_ = v4009
	var v4012 int32
	_ = v4012
	var v4013 int32
	_ = v4013
	var v4016 int32
	_ = v4016
	var v4023 int32
	_ = v4023
	var v4024 int32
	_ = v4024
	var v4027 int32
	_ = v4027
	var v4029 int32
	_ = v4029
	var v4030 int32
	_ = v4030
	var v4033 int32
	_ = v4033
	var v4034 int32
	_ = v4034
	var v4046 int32
	_ = v4046
	var v4067 int32
	_ = v4067
	var v4068 int32
	_ = v4068
	var v4081 int32
	_ = v4081
	var v4133 int32
	_ = v4133
	var v4134 int32
	_ = v4134
	var v4136 int32
	_ = v4136
	var v4146 int32
	_ = v4146
	var v4147 int32
	_ = v4147
	var v4148 int32
	_ = v4148
	var v4150 int32
	_ = v4150
	var v4154 int32
	_ = v4154
	var v4155 int32
	_ = v4155
	var v4156 int32
	_ = v4156
	var v4157 int32
	_ = v4157
	var v4169 int32
	_ = v4169
	var v4177 int32
	_ = v4177
	var v4183 int32
	_ = v4183
	var v4186 int32
	_ = v4186
	var v4187 int32
	_ = v4187
	var v4188 int32
	_ = v4188
	var v4191 int32
	_ = v4191
	var v4194 int32
	_ = v4194
	var v4195 int32
	_ = v4195
	var v4196 int32
	_ = v4196
	var v4199 int32
	_ = v4199
	var v4202 int32
	_ = v4202
	var v4205 int32
	_ = v4205
	var v4206 int32
	_ = v4206
	var v4212 int32
	_ = v4212
	var v4242 int32
	_ = v4242
	var v4243 int32
	_ = v4243
	var v4246 int32
	_ = v4246
	var v4249 int32
	_ = v4249
	var v4252 int32
	_ = v4252
	var v4253 int32
	_ = v4253
	var v4256 int32
	_ = v4256
	var v4257 int32
	_ = v4257
	var v4260 int32
	_ = v4260
	var v4267 int32
	_ = v4267
	var v4268 int32
	_ = v4268
	var v4271 int32
	_ = v4271
	var v4276 int32
	_ = v4276
	var v4279 int32
	_ = v4279
	var v4285 int32
	_ = v4285
	var v4290 int32
	_ = v4290
	var v4324 int32
	_ = v4324
	var v4328 int32
	_ = v4328
	var v4331 int32
	_ = v4331
	var v4332 int32
	_ = v4332
	var v4339 int32
	_ = v4339
	var v4341 int32
	_ = v4341
	var v4369 int32
	_ = v4369
	var v4370 int32
	_ = v4370
	var v4373 int32
	_ = v4373
	var v4376 int32
	_ = v4376
	var v4379 int32
	_ = v4379
	var v4380 int32
	_ = v4380
	var v4383 int32
	_ = v4383
	var v4384 int32
	_ = v4384
	var v4387 int32
	_ = v4387
	var v4394 int32
	_ = v4394
	var v4395 int32
	_ = v4395
	var v4399 int32
	_ = v4399
	var v4403 int32
	_ = v4403
	var v4405 int32
	_ = v4405
	var v4413 int32
	_ = v4413
	var v4442 int32
	_ = v4442
	var v4443 int32
	_ = v4443
	var v4446 int32
	_ = v4446
	var v4450 int32
	_ = v4450
	var v4453 int32
	_ = v4453
	var v4454 int32
	_ = v4454
	var v4461 int32
	_ = v4461
	var v4468 int32
	_ = v4468
	var v4491 int32
	_ = v4491
	var v4492 int32
	_ = v4492
	var v4495 int32
	_ = v4495
	var v4498 int32
	_ = v4498
	var v4501 int32
	_ = v4501
	var v4502 int32
	_ = v4502
	var v4505 int32
	_ = v4505
	var v4506 int32
	_ = v4506
	var v4509 int32
	_ = v4509
	var v4516 int32
	_ = v4516
	var v4517 int32
	_ = v4517
	var v4521 int32
	_ = v4521
	var v4525 int32
	_ = v4525
	var v4527 int32
	_ = v4527
	var v4540 int32
	_ = v4540
	var v4564 int32
	_ = v4564
	var v4565 int32
	_ = v4565
	var v4568 int32
	_ = v4568
	var v4569 int32
	_ = v4569
	var v4570 int32
	_ = v4570
	var v4571 int32
	_ = v4571
	var v4572 int32
	_ = v4572
	var v4573 int32
	_ = v4573
	var v4575 int32
	_ = v4575
	var v4576 int32
	_ = v4576
	var v4577 int32
	_ = v4577
	var v4579 int32
	_ = v4579
	var v4581 int32
	_ = v4581
	var v4584 int32
	_ = v4584
	var v4585 int32
	_ = v4585
	var v4586 int32
	_ = v4586
	var v4589 int32
	_ = v4589
	var v4590 int32
	_ = v4590
	var v4591 int32
	_ = v4591
	var v4592 int32
	_ = v4592
	var v4593 int32
	_ = v4593
	var v4594 int32
	_ = v4594
	var v4596 int32
	_ = v4596
	var v4597 int32
	_ = v4597
	var v4598 int32
	_ = v4598
	var v4600 int32
	_ = v4600
	var v4602 int32
	_ = v4602
	var v4605 int32
	_ = v4605
	var v4606 int32
	_ = v4606
	var v4607 int32
	_ = v4607
	var v4608 int32
	_ = v4608
	var v4609 int32
	_ = v4609
	var v4610 int32
	_ = v4610
	var v4612 int32
	_ = v4612
	var v4613 int32
	_ = v4613
	var v4615 int32
	_ = v4615
	var v4618 int32
	_ = v4618
	var v4625 int32
	_ = v4625
	var v4652 int32
	_ = v4652
	var v4656 int32
	_ = v4656
	var v4657 int32
	_ = v4657
	var v4660 int32
	_ = v4660
	var v4661 int32
	_ = v4661
	var v4694 int32
	_ = v4694
	var v4697 int32
	_ = v4697
	var v4699 int32
	_ = v4699
	var v4700 int32
	_ = v4700
	var v4702 int32
	_ = v4702
	var v4703 int32
	_ = v4703
	var v4704 int32
	_ = v4704
	var v4741 int32
	_ = v4741
	var v4744 int32
	_ = v4744
	var v4750 int32
	_ = v4750
	var v4755 int32
	_ = v4755
	var v4790 int32
	_ = v4790
	var v4793 int32
	_ = v4793
	var v4799 int32
	_ = v4799
	var v4804 int32
	_ = v4804
	var v4808 int32
	_ = v4808
	var v4811 int32
	_ = v4811
	var v4817 int32
	_ = v4817
	var v4822 int32
	_ = v4822
	var v4857 int32
	_ = v4857
	var v4860 int32
	_ = v4860
	var v4866 int32
	_ = v4866
	var v4871 int32
	_ = v4871
	var v4881 int32
	_ = v4881
	var v4889 int32
	_ = v4889
	var v4895 int32
	_ = v4895
	var v4898 int32
	_ = v4898
	var v4899 int32
	_ = v4899
	var v4906 int32
	_ = v4906
	var v4912 int32
	_ = v4912
	var v4913 int32
	_ = v4913
	var v4939 int32
	_ = v4939
	var v4944 int32
	_ = v4944
	var v4946 int32
	_ = v4946
	var v4950 int32
	_ = v4950
	var v4953 int32
	_ = v4953
	var v4955 int32
	_ = v4955
	var v4960 int32
	_ = v4960
	var v4964 int32
	_ = v4964
	var v4967 int32
	_ = v4967
	var v4971 int32
	_ = v4971
	var v4972 int32
	_ = v4972
	var v4976 int32
	_ = v4976
	var v4977 int32
	_ = v4977
	var v4979 int32
	_ = v4979
	var v4981 int32
	_ = v4981
	var v4986 int32
	_ = v4986
	var v4987 int32
	_ = v4987
	var v4988 int32
	_ = v4988
	var v4989 int32
	_ = v4989
	var v4991 int32
	_ = v4991
	var v4992 int32
	_ = v4992
	var v4993 int32
	_ = v4993
	var v4994 int32
	_ = v4994
	var v4995 int32
	_ = v4995
	var v4996 int32
	_ = v4996
	var v4999 int32
	_ = v4999
	var v5001 int32
	_ = v5001
	var v5002 int32
	_ = v5002
	var v5004 int32
	_ = v5004
	var v5005 int32
	_ = v5005
	var v5016 int32
	_ = v5016
	var v5034 int32
	_ = v5034
	var v5038 int32
	_ = v5038
	var v5039 int32
	_ = v5039
	var v5043 int32
	_ = v5043
	var v5045 int32
	_ = v5045
	var v5046 int32
	_ = v5046
	var v5048 int32
	_ = v5048
	var v5049 int32
	_ = v5049
	var v5050 int32
	_ = v5050
	var v5052 int32
	_ = v5052
	var v5053 int32
	_ = v5053
	var v5055 int32
	_ = v5055
	var v5059 int32
	_ = v5059
	var v5060 int32
	_ = v5060
	var v5066 int32
	_ = v5066
	var v5071 int32
	_ = v5071
	var v5072 int32
	_ = v5072
	var v5074 int32
	_ = v5074
	var v5075 int32
	_ = v5075
	var v5079 int32
	_ = v5079
	var v5086 int32
	_ = v5086
	var v5107 int32
	_ = v5107
	var v5112 int32
	_ = v5112
	var v5116 int32
	_ = v5116
	var v5118 int32
	_ = v5118
	var v5122 int32
	_ = v5122
	var v5125 int32
	_ = v5125
	var v5129 int32
	_ = v5129
	var v5132 int32
	_ = v5132
	var v5138 int32
	_ = v5138
	var v5139 int32
	_ = v5139
	var v5142 int32
	_ = v5142
	var v5144 int32
	_ = v5144
	var v5145 int32
	_ = v5145
	var v5148 int32
	_ = v5148
	var v5151 int32
	_ = v5151
	var v5154 int32
	_ = v5154
	var v5157 int32
	_ = v5157
	var v5159 int32
	_ = v5159
	var v5160 int32
	_ = v5160
	var v5163 int32
	_ = v5163
	var v5167 int32
	_ = v5167
	var v5171 int32
	_ = v5171
	var v5174 int32
	_ = v5174
	var v5177 int32
	_ = v5177
	var v5179 int32
	_ = v5179
	var v5180 int32
	_ = v5180
	var v5183 int32
	_ = v5183
	var v5186 int32
	_ = v5186
	var v5189 int32
	_ = v5189
	var v5192 int32
	_ = v5192
	var v5194 int32
	_ = v5194
	var v5195 int32
	_ = v5195
	var v5198 int32
	_ = v5198
	var v5202 int32
	_ = v5202
	var v5206 int32
	_ = v5206
	var v5209 int32
	_ = v5209
	var v5210 int32
	_ = v5210
	var v5219 int32
	_ = v5219
	var v5220 int32
	_ = v5220
	var v5223 int32
	_ = v5223
	var v5224 int32
	_ = v5224
	var v5233 int32
	_ = v5233
	var v5234 int32
	_ = v5234
	var v5235 int32
	_ = v5235
	var v5236 int32
	_ = v5236
	var v5237 int32
	_ = v5237
	var v5242 int32
	_ = v5242
	var v5243 int32
	_ = v5243
	var v5244 int32
	_ = v5244
	var v5248 int32
	_ = v5248
	var v5249 int32
	_ = v5249
	var v5250 int32
	_ = v5250
	var v5251 int32
	_ = v5251
	var v5256 int32
	_ = v5256
	var v5257 int32
	_ = v5257
	var v5258 int32
	_ = v5258
	var v5262 int32
	_ = v5262
	var v5263 int32
	_ = v5263
	var v5264 int32
	_ = v5264
	var v5265 int32
	_ = v5265
	var v5268 int32
	_ = v5268
	var v5271 int32
	_ = v5271
	var v5273 int32
	_ = v5273
	var v5274 int32
	_ = v5274
	var v5286 int32
	_ = v5286
	var v5287 int32
	_ = v5287
	var v5294 int32
	_ = v5294
	var v5300 int32
	_ = v5300
	var v5305 int32
	_ = v5305
	var v5308 int32
	_ = v5308
	var v5310 int32
	_ = v5310
	var v5312 int32
	_ = v5312
	var v5315 int32
	_ = v5315
	var v5316 int32
	_ = v5316
	var v5317 int32
	_ = v5317
	var v5319 int64
	_ = v5319
	var v5321 int64
	_ = v5321
	var v5323 int64
	_ = v5323
	var v5325 int64
	_ = v5325
	var v5328 int64
	_ = v5328
	var v5330 int64
	_ = v5330
	var v5332 int64
	_ = v5332
	var v5334 int64
	_ = v5334
	var v5336 int32
	_ = v5336
	var v5339 int32
	_ = v5339
	var v5340 int32
	_ = v5340
	var v5342 int32
	_ = v5342
	var v5343 int32
	_ = v5343
	var v5345 int32
	_ = v5345
	var v5346 int32
	_ = v5346
	var v5348 int32
	_ = v5348
	var v5351 int32
	_ = v5351
	var v5353 int32
	_ = v5353
	var v5360 int32
	_ = v5360
	var v5389 int32
	_ = v5389
	var v5391 int32
	_ = v5391
	var v5397 int32
	_ = v5397
	var v5398 int32
	_ = v5398
	var v5399 int32
	_ = v5399
	var v5403 int32
	_ = v5403
	var v5404 int32
	_ = v5404
	var v5405 int32
	_ = v5405
	var v5408 int32
	_ = v5408
	var v5411 int32
	_ = v5411
	var v5412 int32
	_ = v5412
	var v5413 int32
	_ = v5413
	var v5428 int32
	_ = v5428
	var v5431 int32
	_ = v5431
	var v5455 int32
	_ = v5455
	var v5456 int32
	_ = v5456
	var v5458 int32
	_ = v5458
	var v5460 int32
	_ = v5460
	var v5463 int32
	_ = v5463
	var v5466 int32
	_ = v5466
	var v5468 int32
	_ = v5468
	var v5471 int32
	_ = v5471
	var v5474 int32
	_ = v5474
	var v5476 int32
	_ = v5476
	var v5479 int32
	_ = v5479
	var v5482 int32
	_ = v5482
	var v5483 int32
	_ = v5483
	var v5484 int32
	_ = v5484
	var v5488 int32
	_ = v5488
	var v5496 int32
	_ = v5496
	var v5527 int32
	_ = v5527
	var v5533 int32
	_ = v5533
	var v5554 int32
	_ = v5554
	var v5557 int32
	_ = v5557
	var v5558 int32
	_ = v5558
	var v5559 int32
	_ = v5559
	var v5563 int32
	_ = v5563
	var v5597 int32
	_ = v5597
	var v5598 int32
	_ = v5598
	var v5599 int32
	_ = v5599
	var v5600 int32
	_ = v5600
	var v5602 int32
	_ = v5602
	var v5603 int32
	_ = v5603
	var v5604 int32
	_ = v5604
	var v5605 int32
	_ = v5605
	var v5606 int32
	_ = v5606
	var v5607 int32
	_ = v5607
	var v5609 int32
	_ = v5609
	var v5610 int32
	_ = v5610
	var v5611 int32
	_ = v5611
	var v5612 int32
	_ = v5612
	var v5615 int32
	_ = v5615
	var v5616 int32
	_ = v5616
	var v5622 int32
	_ = v5622
	var v5628 int32
	_ = v5628
	var v5650 int32
	_ = v5650
	var v5651 int32
	_ = v5651
	var v5654 int32
	_ = v5654
	var v5655 int32
	_ = v5655
	var v5661 int32
	_ = v5661
	var v5688 int32
	_ = v5688
	var v5689 int32
	_ = v5689
	var v5691 int32
	_ = v5691
	var v5693 int32
	_ = v5693
	var v5694 int32
	_ = v5694
	var v5695 int32
	_ = v5695
	var v5697 int32
	_ = v5697
	var v5699 int32
	_ = v5699
	var v5711 int32
	_ = v5711
	var v5712 int32
	_ = v5712
	var v5714 int32
	_ = v5714
	var v5715 int32
	_ = v5715
	var v5716 int32
	_ = v5716
	var v5719 int32
	_ = v5719
	var v5720 int32
	_ = v5720
	var v5722 int32
	_ = v5722
	var v5723 int32
	_ = v5723
	var v5726 int32
	_ = v5726
	var v5727 int32
	_ = v5727
	var v5734 int32
	_ = v5734
	var v5761 int32
	_ = v5761
	var v5765 int32
	_ = v5765
	var v5766 int32
	_ = v5766
	var v5769 int32
	_ = v5769
	var v5770 int32
	_ = v5770
	var v5772 int32
	_ = v5772
	var v5807 int32
	_ = v5807
	var v5818 int32
	_ = v5818
	var v5839 int32
	_ = v5839
	var v5840 int32
	_ = v5840
	var v5846 int32
	_ = v5846
	var v5847 int32
	_ = v5847
	var v5849 int32
	_ = v5849
	var v5850 int32
	_ = v5850
	var v5851 int32
	_ = v5851
	var v5852 int32
	_ = v5852
	var v5853 int32
	_ = v5853
	var v5854 int32
	_ = v5854
	var v5855 int32
	_ = v5855
	var v5857 int32
	_ = v5857
	var v5869 int32
	_ = v5869
	var v5870 int32
	_ = v5870
	var v5874 int32
	_ = v5874
	var v5875 int32
	_ = v5875
	var v5876 int32
	_ = v5876
	var v5877 int32
	_ = v5877
	var v5880 int32
	_ = v5880
	var v5881 int32
	_ = v5881
	var v5883 int32
	_ = v5883
	var v5884 int32
	_ = v5884
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
	var v5895 int32
	_ = v5895
	var v5901 int32
	_ = v5901
	var v5903 int32
	_ = v5903
	var v5928 int32
	_ = v5928
	var v5932 int32
	_ = v5932
	var v5934 int32
	_ = v5934
	var v5938 int32
	_ = v5938
	var v5944 int32
	_ = v5944
	var v5947 int32
	_ = v5947
	var v5949 int32
	_ = v5949
	var v5951 int32
	_ = v5951
	var v5952 int32
	_ = v5952
	var v5956 int32
	_ = v5956
	var v5957 int32
	_ = v5957
	var v5960 int32
	_ = v5960
	var v5961 int32
	_ = v5961
	var v5963 int32
	_ = v5963
	var v5964 int32
	_ = v5964
	var v5970 int32
	_ = v5970
	var v5971 int32
	_ = v5971
	var v5973 int32
	_ = v5973
	var v5974 int32
	_ = v5974
	var v5976 int32
	_ = v5976
	var v5977 int32
	_ = v5977
	var v5979 int32
	_ = v5979
	var v5982 int32
	_ = v5982
	var v5983 int32
	_ = v5983
	var v5987 int32
	_ = v5987
	var v5990 int32
	_ = v5990
	var v5991 int32
	_ = v5991
	var v5992 int32
	_ = v5992
	var v5993 int32
	_ = v5993
	var v5999 int32
	_ = v5999
	var v6000 int32
	_ = v6000
	var v6002 int32
	_ = v6002
	var v6007 int32
	_ = v6007
	var v6011 int32
	_ = v6011
	var v6014 int32
	_ = v6014
	var v6015 int32
	_ = v6015
	var v6016 int32
	_ = v6016
	var v6017 int32
	_ = v6017
	var v6025 int32
	_ = v6025
	var v6026 int32
	_ = v6026
	var v6028 int32
	_ = v6028
	var v6033 int32
	_ = v6033
	var v6037 int32
	_ = v6037
	var v6040 int32
	_ = v6040
	var v6041 int32
	_ = v6041
	var v6042 int32
	_ = v6042
	var v6043 int32
	_ = v6043
	var v6049 int32
	_ = v6049
	var v6050 int32
	_ = v6050
	var v6052 int32
	_ = v6052
	var v6057 int32
	_ = v6057
	var v6062 int32
	_ = v6062
	var v6065 int32
	_ = v6065
	var v6069 int32
	_ = v6069
	var v6070 int32
	_ = v6070
	var v6071 int32
	_ = v6071
	var v6073 int32
	_ = v6073
	var v6078 int32
	_ = v6078
	var v6082 int32
	_ = v6082
	var v6085 int32
	_ = v6085
	var v6087 int32
	_ = v6087
	var v6088 int32
	_ = v6088
	var v6089 int32
	_ = v6089
	var v6090 int32
	_ = v6090
	var v6091 int32
	_ = v6091
	var v6092 int32
	_ = v6092
	var v6093 int32
	_ = v6093
	var v6094 int32
	_ = v6094
	var v6095 int32
	_ = v6095
	var v6096 int32
	_ = v6096
	var v6097 int32
	_ = v6097
	var v6099 int32
	_ = v6099
	var v6108 int32
	_ = v6108
	var v6109 int32
	_ = v6109
	var v6111 int32
	_ = v6111
	var v6116 int32
	_ = v6116
	var v6149 int32
	_ = v6149
	var v6150 int32
	_ = v6150
	var v6151 int32
	_ = v6151
	var v6160 int32
	_ = v6160
	var v6164 int32
	_ = v6164
	var v6174 int32
	_ = v6174
	var v6177 int32
	_ = v6177
	var v6186 int32
	_ = v6186
	var v6190 int32
	_ = v6190
	var v6191 int32
	_ = v6191
	var v6193 int32
	_ = v6193
	var v6194 int32
	_ = v6194
	var v6197 int32
	_ = v6197
	var v6198 int32
	_ = v6198
	var v6200 int32
	_ = v6200
	var v6201 int32
	_ = v6201
	var v6202 int32
	_ = v6202
	var v6203 int32
	_ = v6203
	var v6206 int32
	_ = v6206
	var v6209 int32
	_ = v6209
	var v6215 int32
	_ = v6215
	var v6245 int32
	_ = v6245
	var v6246 int32
	_ = v6246
	var v6249 int32
	_ = v6249
	var v6252 int32
	_ = v6252
	var v6255 int32
	_ = v6255
	var v6256 int32
	_ = v6256
	var v6259 int32
	_ = v6259
	var v6260 int32
	_ = v6260
	var v6263 int32
	_ = v6263
	var v6270 int32
	_ = v6270
	var v6271 int32
	_ = v6271
	var v6276 int32
	_ = v6276
	var v6309 int32
	_ = v6309
	var v6310 int32
	_ = v6310
	var v6311 int32
	_ = v6311
	var v6318 int32
	_ = v6318
	var v6321 int32
	_ = v6321
	var v6325 int32
	_ = v6325
	var v6326 int32
	_ = v6326
	var v6328 int32
	_ = v6328
	var v6333 int32
	_ = v6333
	var v6337 int32
	_ = v6337
	var v6340 int32
	_ = v6340
	var v6341 int32
	_ = v6341
	var v6347 int32
	_ = v6347
	var v6348 int32
	_ = v6348
	var v6350 int32
	_ = v6350
	var v6355 int32
	_ = v6355
	var v6360 int32
	_ = v6360
	var v6361 int32
	_ = v6361
	var v6387 int32
	_ = v6387
	var v6388 int32
	_ = v6388
	var v6390 int32
	_ = v6390
	var v6391 int32
	_ = v6391
	var v6412 int32
	_ = v6412
	var v6415 int32
	_ = v6415
	var v6457 int32
	_ = v6457
	var v6459 int32
	_ = v6459
	var v6462 int32
	_ = v6462
	var v6466 int32
	_ = v6466
	var v6467 int32
	_ = v6467
	var v6468 int32
	_ = v6468
	var v6469 int32
	_ = v6469
	var v6470 int32
	_ = v6470
	var v6471 int32
	_ = v6471
	var v6503 int32
	_ = v6503
	var v6510 int32
	_ = v6510
	var v6511 int32
	_ = v6511
	var v6514 int32
	_ = v6514
	var v6515 int32
	_ = v6515
	var v6518 int32
	_ = v6518
	var v6521 int32
	_ = v6521
	v5 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(336)
	m.G0 = v34
	F_check_stack_depth(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v40 - int32(85) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L15
	case 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38:
		goto L14
	case 4:
		goto L6
	case 39:
		goto L16
	default:
		goto L19
	}
L3:
	;
	m.G0 = v34 + int32(336)
	return v6521
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v6503
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v6503
	*(*int32)(unsafe.Add(mBase, uint32(v34)+296)) = v6503
	v6510 = F_list_make1_impl(m, int32(1), v34+int32(220))
	mBase = m.M
	v6511 = m.ExcPending
	if v6511 != 0 {
		goto L1
	} else {
		goto L1190
	}
L5:
	;
	F_pfree(m, v3343)
	mBase = m.M
	v6149 = m.ExcPending
	if v6149 != 0 {
		goto L1
	} else {
		goto L1135
	}
L6:
	;
	v5849 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5850 = F_transformFromClauseItem(m, l0, v5849, l2, l3)
	mBase = m.M
	v5851 = m.ExcPending
	if v5851 != 0 {
		goto L1
	} else {
		goto L1051
	}
L7:
	;
	v3771 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3776 = F_transformFromClauseItem(m, l0, v3771, v34+int32(292), v34+int32(284))
	mBase = m.M
	v3777 = m.ExcPending
	if v3777 != 0 {
		goto L1
	} else {
		goto L719
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3761 = m.ExcPending
	if v3761 != 0 {
		goto L1
	} else {
		goto L716
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3741 = m.ExcPending
	if v3741 != 0 {
		goto L1
	} else {
		goto L711
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3721 = m.ExcPending
	if v3721 != 0 {
		goto L1
	} else {
		goto L706
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3701 = m.ExcPending
	if v3701 != 0 {
		goto L1
	} else {
		goto L701
	}
L12:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_0), int32(0))
	mBase = m.M
	v3684 = m.ExcPending
	if v3684 != 0 {
		goto L1
	} else {
		goto L697
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3660 = m.ExcPending
	if v3660 != 0 {
		goto L1
	} else {
		goto L691
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3646 = m.ExcPending
	if v3646 != 0 {
		goto L1
	} else {
		goto L688
	}
L15:
	;
	v3307 = F_palloc0(m, int32(72))
	mBase = m.M
	v3308 = m.ExcPending
	if v3308 != 0 {
		goto L1
	} else {
		goto L608
	}
L16:
	;
	v3162 = m.G0
	v3164 = v3162 - int32(96)
	m.G0 = v3164
	v3166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3164)+44)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v3164)+48)) = int64(0)
	v3170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v3170 == int32(0) {
		goto L580
	} else {
		goto L581
	}
L17:
	;
	v1708 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v1708)
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1710 == int32(0) {
		v1939 = v5
		v1944 = v5
		v1955 = v5
		goto L298
	} else {
		goto L299
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(4)
	v1442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v1442)
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1446 != 0 {
		goto L262
	} else {
		goto L263
	}
L19:
	;
	if v40 == int32(64) {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	if v40 != int32(3) {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v47 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v34)+308)) = v1397
	v1430 = F_list_make1_impl(m, int32(1), v34+int32(12))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L259
	}
L23:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v1026 = F_palloc0(m, int32(136))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L1
	} else {
		goto L205
	}
L24:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if l0 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if v990 != 0 {
		v1397 = v990
		goto L22
	} else {
		goto L204
	}
L26:
	;
	if v257 != 0 {
		goto L52
	} else {
		goto L53
	}
L27:
	;
	v57 = l0
	v62 = v5
	goto L30
L28:
	;
	goto L29
L29:
	;
	v257 = int32(0)
	goto L26
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v57)+36))
	if v82 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v193 != 0 {
		v57 = v193
		v62 = v62 + int32(1)
		goto L30
	} else {
		goto L51
	}
L33:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v85 <= int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v88 = int32(0)
	if v88 < v85 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v91 = v85
	goto L37
L36:
	;
	v91 = v88
	goto L37
L37:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v99 = int32(0)
	goto L38
L38:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v92+v99<<(uint(int32(2))%32))))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if base.B2i32(v132 == int32(0))|base.B2i32(v132 != v135) != 0 {
		v153 = v132
		v154 = v135
		goto L41
	} else {
		goto L42
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34+int32(332)))) = v62
	v257 = v128
	goto L26
L40:
	;
	if v153-v154 != 0 {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	v138 = v129
	v139 = v48
	goto L43
L43:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+1)))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	if v143 == int32(0) {
		v153 = v143
		v154 = v142
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v153 = v143
	v154 = v142
	goto L41
L45:
	;
	v146 = int32(1)
	if v143 == v142 {
		v138 = v138 + v146
		v139 = v139 + v146
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v157 = v99 + int32(1)
	if v91 != v157 {
		v99 = v157
		goto L38
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	goto L39
L50:
	;
	goto L32
L51:
	;
	goto L31
L52:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v34)+332))
	v259 = m.G0
	v261 = v259 - int32(32)
	m.G0 = v261
	v264 = F_palloc0(m, int32(136))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v605 = F_get_visible_ENR_metadata(m, v604, v603)
	mBase = m.M
	goto L143
L55:
	;
	v990 = v508
	goto L25
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264))) = int32(101)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v268 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v269 = v268
	goto L59
L58:
	;
	v269 = v257
	goto L59
L59:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+12)) = int32(6)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+88)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v264)+84)) = v273
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v257)+16))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	v279 = base.B2i32(v277 != int32(67))
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+92)) = uint8(v279)
	if v279 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v257)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v257)+36)) = v283 + int32(1)
	goto L62
L61:
	;
	goto L62
L62:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v257)+16))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	if v288 != int32(67) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L139
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L134
	}
L65:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v257)+44))
	v298 = F_list_copy(m, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L69
	}
L66:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	if v291 == int32(1) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
	if v294 == int32(0) {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+96)) = v298
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v257)+48))
	v302 = F_list_copy(m, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+100)) = v302
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v257)+52))
	v306 = F_list_copy(m, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v264)+104)) = v306
	if v268 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v316 = int32(0)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	if v318 != 0 {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	v310 = F_copyObjectImpl(m, v268)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v313 = F_makeAlias(m, v270, int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L77
	}
L76:
	;
	v315 = v310
	goto L72
L77:
	;
	v315 = v313
	goto L72
L78:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	v320 = v319
	goto L80
L79:
	;
	v320 = v316
	goto L80
L80:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v257)+40))
	if v321 == int32(0) {
		v383 = v316
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if v383 < v320 {
		goto L63
	} else {
		goto L91
	}
L82:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	if v324 <= int32(0) {
		v383 = v316
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v338 = v316
	v350 = v318
	goto L84
L84:
	;
	if v320 <= v338 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v383 = v369
	goto L81
L86:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v321)+12))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359+v338<<(uint(int32(2))%32))))
	v364 = F_lappend(m, v350, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	v367 = v350
	goto L88
L88:
	;
	v369 = v338 + int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	if v369 < v370 {
		v338 = v369
		v350 = v367
		goto L84
	} else {
		goto L90
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315)+8)) = v364
	v367 = v364
	goto L88
L90:
	;
	goto L85
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+8)) = v315
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v257)+20))
	if v405 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v405)+12))
	v408 = F_makeString(m, v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	v436 = int32(0)
	goto L94
L94:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v257)+24))
	if v437 != 0 {
		goto L103
	} else {
		goto L104
	}
L95:
	;
	v410 = F_lappend(m, v406, v408)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v412)+8)) = v410
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v264)+96))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v257)+20))
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417)+8)))
	if v418 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v419 = int32(2249)
	goto L99
L98:
	;
	v419 = int32(2287)
	goto L99
L99:
	;
	v420 = F_lappend_oid(m, v414, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+96)) = v420
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v264)+100))
	v425 = F_lappend_int(m, v423, int32(-1))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+100)) = v425
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v264)+104))
	v430 = F_lappend_oid(m, v428, int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+104)) = v430
	v436 = int32(1)
	goto L94
L103:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+8))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v437)+8))
	v441 = F_makeString(m, v440)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	v493 = v436
	goto L105
L105:
	;
	v494 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+125)) = uint8(v494)
	v496 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+124)) = uint8(v496)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v499 = F_lappend(m, v498, v264)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L116
	}
L106:
	;
	v443 = F_lappend(m, v439, v441)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+8)) = v443
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v264)+96))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v257)+24))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+28))
	v450 = F_lappend_oid(m, v447, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+96)) = v450
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v264)+100))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v257)+24))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)+32))
	v456 = F_lappend_int(m, v453, v455)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+100)) = v456
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v264)+104))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v257)+24))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+36))
	v462 = F_lappend_oid(m, v459, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+104)) = v462
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+8))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v257)+24))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+20))
	v469 = F_makeString(m, v468)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v471 = F_lappend(m, v466, v469)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v473)+8)) = v471
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v264)+96))
	v477 = F_lappend_oid(m, v475, int32(2287))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+96)) = v477
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v264)+100))
	v482 = F_lappend_int(m, v480, int32(-1))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+100)) = v482
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v264)+104))
	v487 = F_lappend_oid(m, v485, int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+104)) = v487
	v493 = v436 | int32(2)
	goto L105
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v499
	if v499 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	v504 = v502
	goto L119
L118:
	;
	v504 = int32(0)
	goto L119
L119:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v264)+96))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v264)+100))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v264)+104))
	v508 = F_buildNSItemFromLists(m, v264, v504, v505, v506, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v264)+88))
	v511 = int32(0)
	if base.B2i32(v510 == v511)|base.B2i32(v493 == v511) != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	m.G0 = v261 + int32(32)
	goto L55
L122:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v508)+16))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v517)+8))
	if v518 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	v521 = v519
	goto L125
L124:
	;
	v521 = int32(0)
	goto L125
L125:
	;
	v527 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v516+v521<<(uint(int32(5))%32)-int32(2)))) = uint8(v527)
	if v493 == v527 {
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v508)+16))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)+8))
	if v533 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	v536 = v534
	goto L129
L128:
	;
	v536 = int32(0)
	goto L129
L129:
	;
	v542 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v531+v536<<(uint(int32(5))%32)-int32(34)))) = uint8(v542)
	if v493 == int32(2) {
		goto L121
	} else {
		goto L130
	}
L130:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v508)+16))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)+8))
	if v548 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	v551 = v549
	goto L133
L132:
	;
	v551 = int32(0)
	goto L133
L133:
	;
	v557 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v546+v551<<(uint(int32(5))%32)-int32(66)))) = uint8(v557)
	goto L121
L134:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v261)+16)) = v570
	F_errmsg(m, int32(_a_F_transformFromClauseItem_1), v261+int32(16))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_parser_errposition(m, l0, v577)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(2377), int32(_a_F_transformFromClauseItem_3))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errcode(m, int32(_a_F_transformFromClauseItem_4))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+8)) = v320
	*(*int32)(unsafe.Add(mBase, uint32(v261)+4)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v270
	F_errmsg(m, int32(_a_F_transformFromClauseItem_5), v261)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(2403), int32(_a_F_transformFromClauseItem_3))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	if base.B2i32(v605 != int32(0)) == int32(0) {
		goto L23
	} else {
		goto L144
	}
L144:
	;
	v610 = m.G0
	v612 = v610 - int32(32)
	m.G0 = v612
	v615 = F_palloc0(m, int32(136))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L146
	}
L145:
	;
	v990 = v915
	goto L25
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615))) = int32(101)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v623 = l1 + int32(12)
	if v619 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v624 = v619 + int32(4)
	goto L149
L148:
	;
	v624 = v623
	goto L149
L149:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v624)))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v628 = int32(0)
	if v626 == v628 {
		v660 = v628
		goto L152
	} else {
		goto L153
	}
L150:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L201
	}
L151:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v660)+12))
	if v663 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L152:
	;
	goto L151
L153:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	if v633 == int32(0) {
		v660 = v628
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v633)+4))
	if v636 <= int32(0) {
		v660 = v628
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v633)+12))
	v641 = int32(0)
	goto L156
L156:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v639+v641<<(uint(int32(2))%32))))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	v651 = F_strcmp(m, v650, v627)
	mBase = m.M
	if v651 == int32(0) {
		v660 = v649
		goto L152
	} else {
		goto L158
	}
L157:
	;
	v660 = int32(0)
	goto L152
L158:
	;
	v655 = v641 + int32(1)
	if v636 != v655 {
		v641 = v655
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615)+12)) = int32(7)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v660)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v615)+16)) = v668
	v670 = F_ENRMetadataGetTupDesc(m, v660)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L198
	}
L163:
	;
	v673 = F_makeAlias(m, v625, int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615)+8)) = v673
	F_buildRelationAliases(m, v670, v619, v673)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v660)))
	*(*int32)(unsafe.Add(mBase, uint32(v615)+108)) = v678
	v680 = *(*float64)(unsafe.Add(mBase, uint32(v660)+16))
	v681 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v615)+104)) = v681
	*(*int64)(unsafe.Add(mBase, uint32(v615)+96)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v615)+112)) = v680
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v670)))
	if v681 < v687 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v697 = v687
	v701 = int32(1)
	goto L169
L167:
	;
	goto L168
L168:
	;
	v801 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v615)+125)) = uint8(v801)
	v803 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v615)+124)) = uint8(v803)
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v806 = F_lappend(m, v805, v615)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L183
	}
L169:
	;
	v726 = v670 + v697<<(uint(int32(4))%32) + v701*int32(100)
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726)+11)))
	if v727 == int32(1) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L168
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615)+104)) = v764
	v767 = v701 + int32(1)
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v670)))
	if v767 <= v768 {
		v697 = v768
		v701 = v767
		goto L169
	} else {
		goto L182
	}
L172:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v615)+96))
	v732 = F_lappend_oid(m, v730, int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v745 = v726 - int32(80)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v745)+68))
	if v746 == int32(0) {
		goto L150
	} else {
		goto L178
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615)+96)) = v732
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v615)+100))
	v737 = F_lappend_int(m, v735, int32(0))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615)+100)) = v737
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v615)+104))
	v742 = F_lappend_oid(m, v740, int32(0))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v764 = v742
	goto L171
L178:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v615)+96))
	v750 = F_lappend_oid(m, v749, v746)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615)+96)) = v750
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v615)+100))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v745)+76))
	v755 = F_lappend_int(m, v753, v754)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615)+100)) = v755
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v615)+104))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v745)+96))
	v760 = F_lappend_oid(m, v758, v759)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v764 = v760
	goto L171
L182:
	;
	goto L170
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v806
	if v806 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v806)+4))
	v810 = v809
	goto L186
L185:
	;
	v810 = v5
	goto L186
L186:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v670)))
	v814 = F_palloc0(m, v811<<(uint(int32(5))%32))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	if int32(0) < v811 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v830 = int32(0)
	goto L191
L189:
	;
	goto L190
L190:
	;
	v915 = F_palloc(m, int32(28))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L197
	}
L191:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v670)))
	v856 = v670 + v850<<(uint(int32(4))%32) + v830*int32(100)
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856)+111)))
	if v857 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	goto L190
L193:
	;
	v862 = v814 + v830<<(uint(int32(5))%32)
	v864 = v830 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v862)+4)) = uint16(v864)
	*(*int32)(unsafe.Add(mBase, uint32(v862))) = v810
	v868 = v856 + int32(20)
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v868)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v862)+8)) = v869
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v868)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v862)+12)) = v871
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v868)+96))
	*(*uint16)(unsafe.Add(mBase, uint32(v862)+28)) = uint16(v864)
	*(*int32)(unsafe.Add(mBase, uint32(v862)+24)) = v810
	*(*int32)(unsafe.Add(mBase, uint32(v862)+16)) = v873
	goto L195
L194:
	;
	goto L195
L195:
	;
	v881 = v830 + int32(1)
	if v881 != v811 {
		v830 = v881
		goto L191
	} else {
		goto L196
	}
L196:
	;
	goto L192
L197:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v615)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v915)+20)) = int64(16777473)
	*(*int32)(unsafe.Add(mBase, uint32(v915)+16)) = v814
	*(*int32)(unsafe.Add(mBase, uint32(v915)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v915)+8)) = v810
	*(*int32)(unsafe.Add(mBase, uint32(v915)+4)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v915))) = v917
	m.G0 = v612 + int32(32)
	goto L145
L198:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v660)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v612)+16)) = v933
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_6), v612+int32(16))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(2506), int32(_a_F_transformFromClauseItem_7))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
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
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v623)))
	*(*int32)(unsafe.Add(mBase, uint32(v612))) = v949
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_8), v612)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(2546), int32(_a_F_transformFromClauseItem_7))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	goto L23
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1026))) = int32(101)
	if v1022 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v1034 = v1022 + int32(4)
	goto L208
L207:
	;
	v1034 = l1 + int32(12)
	goto L208
L208:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1034)))
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+4)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+12)) = int32(0)
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)))
	if v1039 != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v1229 = F_parserOpenTable(m, l0, l1, v1203)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L1
	} else {
		goto L235
	}
L210:
	;
	v1203 = int32(2)
	goto L209
L211:
	;
	goto L212
L212:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v1041 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1203 = int32(1)
	goto L209
L214:
	;
	goto L215
L215:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+4))
	if v1046 <= int32(0) {
		v1203 = int32(1)
		goto L209
	} else {
		goto L216
	}
L216:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+12))
	v1061 = int32(0)
	goto L217
L217:
	;
	v1081 = int32(2)
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1049+v1061<<(uint(v1081)%32))))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+4))
	if v1086 == int32(0) {
		v1203 = v1081
		goto L209
	} else {
		goto L219
	}
L218:
	;
	v1203 = v1194
	goto L209
L219:
	;
	if v1035 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v1194 = int32(1)
	v1196 = v1061 + v1194
	if v1046 != v1196 {
		v1061 = v1196
		goto L217
	} else {
		goto L234
	}
L221:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+4))
	if v1091 <= int32(0) {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+12))
	v1102 = int32(0)
	goto L223
L223:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1094+v1102<<(uint(int32(2))%32))))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+12))
	v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035))))
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131))))
	if base.B2i32(v1134 == int32(0))|base.B2i32(v1134 != v1137) != 0 {
		v1155 = v1134
		v1156 = v1137
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L220
L225:
	;
	if v1155-v1156 == int32(0) {
		v1203 = v1081
		goto L209
	} else {
		goto L232
	}
L226:
	;
	goto L225
L227:
	;
	v1140 = v1035
	v1141 = v1131
	goto L228
L228:
	;
	v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1141)+1)))
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1140)+1)))
	if v1145 == int32(0) {
		v1155 = v1145
		v1156 = v1144
		goto L226
	} else {
		goto L230
	}
L229:
	;
	v1155 = v1145
	v1156 = v1144
	goto L226
L230:
	;
	v1148 = int32(1)
	if v1145 == v1144 {
		v1140 = v1140 + v1148
		v1141 = v1141 + v1148
		goto L228
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v1161 = v1102 + int32(1)
	if v1161 != v1091 {
		v1102 = v1161
		goto L223
	} else {
		goto L233
	}
L233:
	;
	goto L224
L234:
	;
	goto L218
L235:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+56))
	*(*uint8)(unsafe.Add(mBase, uint32(v1026)+20)) = uint8(v1023)
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+16)) = v1231
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+48))
	v1235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1234)+119)))
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+24)) = v1203
	*(*uint8)(unsafe.Add(mBase, uint32(v1026)+21)) = uint8(v1235)
	v1239 = F_makeAlias(m, v1035, int32(0))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+8)) = v1239
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+52))
	F_buildRelationAliases(m, v1242, v1022, v1239)
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v1245 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1026)+125)) = uint8(v1245)
	v1247 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1026)+124)) = uint8(v1247)
	v1250 = F_palloc0(m, int32(40))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1250))) = int32(102)
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1026)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1250)+4)) = v1254
	v1256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1026)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1250)+8)) = uint8(v1256)
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1259 = F_lappend(m, v1258, v1250)
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1259
	if v1259 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1259)+4))
	v1264 = v1262
	goto L242
L241:
	;
	v1264 = int32(0)
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+28)) = v1264
	*(*int64)(unsafe.Add(mBase, uint32(v1250)+16)) = int64(2)
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1269 = F_lappend(m, v1268, v1026)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1269
	v1272 = int32(0)
	if v1269 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1269)+4))
	v1275 = v1274
	goto L246
L245:
	;
	v1275 = v1272
	goto L246
L246:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+52))
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1276)))
	v1280 = F_palloc0(m, v1277<<(uint(int32(5))%32))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	if int32(0) < v1277 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1290 = v1272
	goto L251
L249:
	;
	goto L250
L250:
	;
	v1380 = F_palloc(m, int32(28))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L1
	} else {
		goto L257
	}
L251:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1276)))
	v1321 = v1276 + v1315<<(uint(int32(4))%32) + v1290*int32(100)
	v1322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1321)+111)))
	if v1322 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	goto L250
L253:
	;
	v1327 = v1280 + v1290<<(uint(int32(5))%32)
	v1329 = v1290 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1327)+4)) = uint16(v1329)
	*(*int32)(unsafe.Add(mBase, uint32(v1327))) = v1275
	v1333 = v1321 + int32(20)
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1327)+8)) = v1334
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1327)+12)) = v1336
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+96))
	*(*uint16)(unsafe.Add(mBase, uint32(v1327)+28)) = uint16(v1329)
	*(*int32)(unsafe.Add(mBase, uint32(v1327)+24)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v1327)+16)) = v1338
	goto L255
L254:
	;
	goto L255
L255:
	;
	v1346 = v1290 + int32(1)
	if v1346 != v1277 {
		v1290 = v1346
		goto L251
	} else {
		goto L256
	}
L256:
	;
	goto L252
L257:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1026)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1380)+20)) = int64(16777473)
	*(*int32)(unsafe.Add(mBase, uint32(v1380)+16)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v1380)+12)) = v1250
	*(*int32)(unsafe.Add(mBase, uint32(v1380)+8)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v1380)+4)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v1380))) = v1382
	F_relation_close(m, v1229, int32(0))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	v1397 = v1380
	goto L22
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1430
	v1434 = F_palloc0(m, int32(8))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1434))) = int32(63)
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1434)+4)) = v1438
	v6521 = v1434
	goto L3
L261:
	;
	v1675 = F_parse_sub_analyze(m, v1444, l0, int32(0), v1674)
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L1
	} else {
		goto L292
	}
L262:
	;
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+4))
	v1449 = v1447
	goto L264
L263:
	;
	v1449 = int32(0)
	goto L264
L264:
	;
	v1452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)))
	if v1452 != 0 {
		v1674 = int32(1)
		goto L261
	} else {
		goto L265
	}
L265:
	;
	v1453 = int32(0)
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v1454 == v1453 {
		v1674 = v1453
		goto L261
	} else {
		goto L266
	}
L266:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+4))
	if v1457 <= int32(0) {
		v1619 = v5
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1674 = v1619
	goto L261
L268:
	;
	v1460 = int32(0)
	if v1460 < v1457 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1463 = v1457
	goto L271
L270:
	;
	v1463 = v1460
	goto L271
L271:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+12))
	v1469 = int32(0)
	goto L272
L272:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v1464+v1469<<(uint(int32(2))%32))))
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+4))
	v1502 = base.B2i32(v1500 == int32(0))
	if v1500 == int32(0) {
		v1619 = v1502
		goto L267
	} else {
		goto L274
	}
L273:
	;
	v1619 = v1502
	goto L267
L274:
	;
	if v1449 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1610 = v1469 + int32(1)
	if v1610 != v1463 {
		v1469 = v1610
		goto L272
	} else {
		goto L291
	}
L276:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1500)+4))
	if v1507 <= int32(0) {
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1500)+12))
	v1518 = int32(0)
	goto L278
L278:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1510+v1518<<(uint(int32(2))%32))))
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1546)+12))
	v1550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1449))))
	v1553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1547))))
	if base.B2i32(v1550 == int32(0))|base.B2i32(v1550 != v1553) != 0 {
		v1571 = v1550
		v1572 = v1553
		goto L281
	} else {
		goto L282
	}
L279:
	;
	v1674 = int32(1)
	goto L261
L280:
	;
	if v1571-v1572 != 0 {
		goto L287
	} else {
		goto L288
	}
L281:
	;
	goto L280
L282:
	;
	v1556 = v1449
	v1557 = v1547
	goto L283
L283:
	;
	v1560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1557)+1)))
	v1561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1556)+1)))
	if v1561 == int32(0) {
		v1571 = v1561
		v1572 = v1560
		goto L281
	} else {
		goto L285
	}
L284:
	;
	v1571 = v1561
	v1572 = v1560
	goto L281
L285:
	;
	v1564 = int32(1)
	if v1561 == v1560 {
		v1556 = v1556 + v1564
		v1557 = v1557 + v1564
		goto L283
	} else {
		goto L286
	}
L286:
	;
	goto L284
L287:
	;
	v1575 = v1518 + int32(1)
	if v1507 != v1575 {
		v1518 = v1575
		goto L278
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	goto L279
L290:
	;
	goto L275
L291:
	;
	goto L273
L292:
	;
	v1677 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v1677
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v1677)
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1675)))
	if v1681 != int32(67) {
		goto L8
	} else {
		goto L293
	}
L293:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1675)+4))
	if v1684 != int32(1) {
		goto L8
	} else {
		goto L294
	}
L294:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v1690 = F_addRangeTableEntryForSubquery(m, l0, v1675, v1687, v1688, int32(1))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1690
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v1690
	*(*int32)(unsafe.Add(mBase, uint32(v34)+304)) = v1690
	v1698 = F_list_make1_impl(m, int32(1), v34+int32(16))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1698
	v1702 = F_palloc0(m, int32(8))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1702))) = int32(63)
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1702)+4)) = v1706
	v6521 = v1702
	goto L3
L298:
	;
	v1964 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v1964)
	F_assign_list_collations(m, l0, v1939)
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L1
	} else {
		goto L347
	}
L299:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1710)+4))
	if v1713 <= int32(0) {
		v1939 = v5
		v1944 = v5
		v1955 = v5
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v1722 = v5
	v1726 = v5
	v1727 = v5
	v1738 = v5
	goto L301
L301:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v1710)+12))
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1747+v1726<<(uint(int32(2))%32))))
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+12))
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v1752)+4))
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1752)))
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1754)))
	if v1755 != int32(76) {
		goto L304
	} else {
		goto L305
	}
L302:
	;
	v1939 = v1904
	v1944 = v1909
	v1955 = v1920
	goto L298
L303:
	;
	v1930 = v1726 + int32(1)
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1710)+4))
	if v1930 < v1931 {
		v1722 = v1904
		v1726 = v1930
		v1727 = v1909
		v1738 = v1920
		goto L301
	} else {
		goto L346
	}
L304:
	;
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v1883 = F_transformExpr(m, l0, v1754, int32(5))
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L1
	} else {
		goto L336
	}
L305:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+4))
	if v1758 == int32(0) {
		goto L304
	} else {
		goto L306
	}
L306:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+4))
	if v1761 != int32(1) {
		goto L304
	} else {
		goto L307
	}
L307:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+12))
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v1764)))
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v1765)+4))
	v1767 = int32(_a_F_transformFromClauseItem_9)
	v1770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1766))))
	v1773 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformFromClauseItem[0])))
	if base.B2i32(v1770 == int32(0))|base.B2i32(v1770 != v1773) != 0 {
		v1791 = v1770
		v1792 = v1773
		goto L309
	} else {
		goto L310
	}
L308:
	;
	if v1791-v1792 != 0 {
		goto L304
	} else {
		goto L315
	}
L309:
	;
	goto L308
L310:
	;
	v1776 = v1766
	v1777 = v1767
	goto L311
L311:
	;
	v1780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1777)+1)))
	v1781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1776)+1)))
	if v1781 == int32(0) {
		v1791 = v1781
		v1792 = v1780
		goto L309
	} else {
		goto L313
	}
L312:
	;
	v1791 = v1781
	v1792 = v1780
	goto L309
L313:
	;
	v1784 = int32(1)
	if v1781 == v1780 {
		v1776 = v1776 + v1784
		v1777 = v1777 + v1784
		goto L311
	} else {
		goto L314
	}
L314:
	;
	goto L312
L315:
	;
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+8))
	if v1794 == int32(0) {
		goto L304
	} else {
		goto L316
	}
L316:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1794)+4))
	if v1797 < int32(2) {
		goto L304
	} else {
		goto L317
	}
L317:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+12))
	if v1800 != 0 {
		goto L304
	} else {
		goto L318
	}
L318:
	;
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+16))
	if v1801 != 0 {
		goto L304
	} else {
		goto L319
	}
L319:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+20))
	if v1802 != 0 {
		goto L304
	} else {
		goto L320
	}
L320:
	;
	v1803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1754)+25)))
	if v1803 != 0 {
		goto L304
	} else {
		goto L321
	}
L321:
	;
	v1804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1754)+26)))
	if v1804 != 0 {
		goto L304
	} else {
		goto L322
	}
L322:
	;
	v1805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1754)+27)))
	if v1805|v1753 != 0 {
		goto L304
	} else {
		goto L323
	}
L323:
	;
	v1814 = v1722
	v1815 = int32(0)
	v1819 = v1727
	v1830 = v1738
	goto L324
L324:
	;
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1794)+12))
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1839+v1815<<(uint(int32(2))%32))))
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v1846 = F_SystemFuncName(m, int32(_a_F_transformFromClauseItem_9))
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L1
	} else {
		goto L326
	}
L325:
	;
	v1904 = v1866
	v1909 = v1873
	v1920 = v1870
	goto L303
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v1843
	*(*int32)(unsafe.Add(mBase, uint32(v34)+332)) = v1843
	v1853 = F_list_make1_impl(m, int32(1), v34+int32(28))
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+32))
	v1857 = F_makeFuncCall(m, v1846, v1853, int32(0), v1856)
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	v1860 = F_transformExpr(m, l0, v1857, int32(5))
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if base.B2i32(v1862 != v1844)&base.B2i32(v1862 != v1860) != 0 {
		goto L9
	} else {
		goto L330
	}
L330:
	;
	v1866 = F_lappend(m, v1814, v1860)
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	v1868 = F_FigureColname(m, v1857)
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	v1870 = F_lappend(m, v1830, v1868)
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	v1873 = F_lappend(m, v1819, int32(0))
	mBase = m.M
	v1874 = m.ExcPending
	if v1874 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	v1876 = v1815 + int32(1)
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1794)+4))
	if v1876 < v1877 {
		v1814 = v1866
		v1815 = v1876
		v1819 = v1873
		v1830 = v1870
		goto L324
	} else {
		goto L335
	}
L335:
	;
	goto L325
L336:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if base.B2i32(v1881 != v1885)&base.B2i32(v1883 != v1885) != 0 {
		goto L10
	} else {
		goto L337
	}
L337:
	;
	v1889 = F_lappend(m, v1722, v1883)
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	v1891 = F_FigureColname(m, v1754)
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	v1893 = F_lappend(m, v1738, v1891)
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	if v1753 != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1895 != 0 {
		goto L11
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	v1896 = F_lappend(m, v1727, v1753)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L1
	} else {
		goto L345
	}
L344:
	;
	goto L343
L345:
	;
	v1904 = v1889
	v1909 = v1896
	v1920 = v1893
	goto L303
L346:
	;
	goto L302
L347:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1968 != 0 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	if v1939 != 0 {
		goto L352
	} else {
		goto L353
	}
L349:
	;
	v2009 = v1944
	goto L350
L350:
	;
	v2010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v2010 != 0 {
		goto L366
	} else {
		goto L367
	}
L351:
	;
	v1999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v1999 == int32(1) {
		goto L13
	} else {
		goto L363
	}
L352:
	;
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v1939)+4))
	if v1969 == int32(1) {
		goto L351
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v1972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L1
	} else {
		goto L356
	}
L355:
	;
	goto L354
L356:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	if v1972 == int32(1) {
		goto L12
	} else {
		goto L358
	}
L358:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_10), int32(0))
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L1
	} else {
		goto L359
	}
L359:
	;
	F_errhint(m, int32(_a_F_transformFromClauseItem_11), int32(0))
	mBase = m.M
	v1989 = m.ExcPending
	if v1989 != 0 {
		goto L1
	} else {
		goto L360
	}
L360:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1991 = F_exprLocation(m, v1990)
	mBase = m.M
	F_parser_errposition(m, l0, v1991)
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(650), int32(_a_F_transformFromClauseItem_13))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v1968
	*(*int32)(unsafe.Add(mBase, uint32(v34)+328)) = v1968
	v2007 = F_list_make1_impl(m, int32(1), v34+int32(24))
	mBase = m.M
	v2008 = m.ExcPending
	if v2008 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	v2009 = v2007
	goto L350
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v3013
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v3013
	*(*int32)(unsafe.Add(mBase, uint32(v34)+300)) = v3013
	v3152 = F_list_make1_impl(m, int32(1), v34+int32(20))
	mBase = m.M
	v3153 = m.ExcPending
	if v3153 != 0 {
		goto L1
	} else {
		goto L576
	}
L366:
	;
	v2015 = int32(1)
	goto L368
L367:
	;
	v2013 = F_contain_vars_of_level(m, v1939, int32(0))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L1
	} else {
		goto L369
	}
L368:
	;
	v2017 = m.G0
	v2019 = v2017 - int32(80)
	m.G0 = v2019
	v2022 = F_palloc0(m, int32(136))
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L1
	} else {
		goto L370
	}
L369:
	;
	v2015 = v2013
	goto L368
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2022))) = int32(101)
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1939 != 0 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v1939)+4))
	v2028 = v2027
	goto L373
L372:
	;
	v2028 = v5
	goto L373
L373:
	;
	v2029 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2022)+68)) = v2029
	*(*int32)(unsafe.Add(mBase, uint32(v2022)+36)) = v2029
	*(*int64)(unsafe.Add(mBase, uint32(v2022)+12)) = int64(3)
	v2035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	*(*int32)(unsafe.Add(mBase, uint32(v2022)+4)) = v2026
	*(*uint8)(unsafe.Add(mBase, uint32(v2022)+72)) = uint8(v2035)
	if v2026 != 0 {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v2041 = v2026 + int32(4)
	goto L376
L375:
	;
	v2040 = *(*int32)(unsafe.Add(mBase, uint32(v1955)+12))
	v2041 = v2040
	goto L376
L376:
	;
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v2041)))
	v2044 = F_makeAlias(m, v2042, int32(0))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L1
	} else {
		goto L377
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2022)+8)) = v2044
	v2051 = base.B2i32(v2026 == int32(0)) | base.B2i32(v2028 != int32(1))
	v2054 = F_palloc(m, v2028<<(uint(int32(2))%32))
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L1
	} else {
		goto L378
	}
L378:
	;
	v2061 = int32(0)
	v2075 = v5
	goto L386
L379:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3128 = m.ExcPending
	if v3128 != 0 {
		goto L1
	} else {
		goto L571
	}
L380:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3108 = m.ExcPending
	if v3108 != 0 {
		goto L1
	} else {
		goto L566
	}
L381:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3083 = m.ExcPending
	if v3083 != 0 {
		goto L1
	} else {
		goto L560
	}
L382:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3061 = m.ExcPending
	if v3061 != 0 {
		goto L1
	} else {
		goto L555
	}
L383:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3042 = m.ExcPending
	if v3042 != 0 {
		goto L1
	} else {
		goto L550
	}
L384:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_14), int32(0))
	mBase = m.M
	v3030 = m.ExcPending
	if v3030 != 0 {
		goto L1
	} else {
		goto L547
	}
L385:
	;
	F_buildRelationAliases(m, v2869, v2026, v2044)
	mBase = m.M
	v2897 = m.ExcPending
	if v2897 != 0 {
		goto L1
	} else {
		goto L531
	}
L386:
	;
	v2087 = int32(0)
	if v1939 == v2087 {
		v2097 = v2087
		goto L388
	} else {
		goto L389
	}
L387:
	;
	v2693 = v2122 + v2123
	if int32(1665) <= v2693 {
		goto L380
	} else {
		goto L514
	}
L388:
	;
	v2098 = int32(0)
	if v1955 == v2098 {
		v2109 = v2098
		goto L391
	} else {
		goto L392
	}
L389:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v1939)+4))
	if v2091 <= v2075 {
		v2097 = int32(0)
		goto L388
	} else {
		goto L390
	}
L390:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v1939)+12))
	v2097 = v2093 + v2075<<(uint(int32(2))%32)
	goto L388
L391:
	;
	if v2009 != 0 {
		goto L396
	} else {
		goto L397
	}
L392:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v1955)+4))
	if v2103 <= v2075 {
		v2109 = int32(0)
		goto L391
	} else {
		goto L393
	}
L393:
	;
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(v1955)+12))
	v2109 = v2105 + v2075<<(uint(int32(2))%32)
	goto L391
L394:
	;
	goto L387
L395:
	;
	v2130 = v2075 << (uint(int32(2)) % 32)
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v2120+v2130)))
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v2109)))
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2097)))
	v2136 = F_palloc0(m, int32(32))
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		goto L1
	} else {
		goto L404
	}
L396:
	;
	v2110 = int32(0)
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v2009)+4))
	if base.B2i32(v2109 == v2110)|(base.B2i32(v2097 == v2110)|base.B2i32(v2114 <= v2075)) == v2110 {
		goto L399
	} else {
		goto L400
	}
L397:
	;
	v2122 = v2098
	goto L398
L398:
	;
	v2123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v2123|base.B2i32(int32(1) < v2028) != 0 {
		goto L394
	} else {
		goto L403
	}
L399:
	;
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v2009)+12))
	if v2120 != 0 {
		goto L395
	} else {
		goto L402
	}
L400:
	;
	goto L401
L401:
	;
	v2122 = v2061
	goto L398
L402:
	;
	goto L401
L403:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v2054)))
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+76)) = v2127
	v2869 = v2127
	goto L385
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2136))) = int32(103)
	v2140 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2136)+12)) = v2140
	*(*int32)(unsafe.Add(mBase, uint32(v2136)+4)) = v2134
	*(*int64)(unsafe.Add(mBase, uint32(v2136)+20)) = v2140
	*(*int32)(unsafe.Add(mBase, uint32(v2136)+28)) = int32(0)
	v2151 = F_get_expr_result_type(m, v2134, v2019+int32(72), v2019+int32(76))
	mBase = m.M
	v2152 = m.ExcPending
	if v2152 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	if v2132 != 0 {
		goto L408
	} else {
		goto L409
	}
L406:
	;
	v2679 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+76))
	v2680 = *(*int32)(unsafe.Add(mBase, uint32(v2679)))
	*(*int32)(unsafe.Add(mBase, uint32(v2136)+8)) = v2680
	v2682 = *(*int32)(unsafe.Add(mBase, uint32(v2022)+68))
	v2683 = F_lappend(m, v2682, v2136)
	mBase = m.M
	v2684 = m.ExcPending
	if v2684 != 0 {
		goto L1
	} else {
		goto L513
	}
L407:
	;
	v2508 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+4))
	if int32(1601) <= v2508 {
		goto L382
	} else {
		goto L493
	}
L408:
	;
	if v2151 == int32(3) {
		goto L407
	} else {
		goto L411
	}
L409:
	;
	goto L410
L410:
	;
	if v2151 == int32(3) {
		goto L383
	} else {
		goto L427
	}
L411:
	;
	v2155 = int32(1)
	if base.Ui32(v2151-v2155) <= base.Ui32(v2155) {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v2159 = F_exprType(m, v2134)
	mBase = m.M
	v2160 = m.ExcPending
	if v2160 != 0 {
		goto L1
	} else {
		goto L415
	}
L413:
	;
	goto L414
L414:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L1
	} else {
		goto L422
	}
L415:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L1
	} else {
		goto L417
	}
L417:
	;
	if v2159 == int32(2249) {
		goto L384
	} else {
		goto L418
	}
L418:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_15), int32(0))
	mBase = m.M
	v2173 = m.ExcPending
	if v2173 != 0 {
		goto L1
	} else {
		goto L419
	}
L419:
	;
	v2174 = F_exprLocation(m, v2132)
	mBase = m.M
	F_parser_errposition(m, l0, v2174)
	mBase = m.M
	v2176 = m.ExcPending
	if v2176 != 0 {
		goto L1
	} else {
		goto L420
	}
L420:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1858), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L422:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L1
	} else {
		goto L423
	}
L423:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_17), int32(0))
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	v2193 = F_exprLocation(m, v2132)
	mBase = m.M
	F_parser_errposition(m, l0, v2193)
	mBase = m.M
	v2195 = m.ExcPending
	if v2195 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1865), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L427:
	;
	if base.Ui32(v2151-int32(1)) < base.Ui32(int32(2)) {
		goto L406
	} else {
		goto L428
	}
L428:
	;
	if v2151 != 0 {
		goto L381
	} else {
		goto L429
	}
L429:
	;
	v2208 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+76)) = v2208
	if v2134 == int32(0) {
		goto L433
	} else {
		goto L434
	}
L431:
	;
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+72))
	v2491 = F_exprTypmod(m, v2134)
	mBase = m.M
	v2492 = m.ExcPending
	if v2492 != 0 {
		goto L1
	} else {
		goto L489
	}
L432:
	;
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+4))
	v2481 = v2457
	goto L431
L433:
	;
	if v2051 != 0 {
		v2481 = v2133
		goto L431
	} else {
		goto L488
	}
L434:
	;
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v2134)))
	if v2213 != int32(15) {
		goto L433
	} else {
		goto L435
	}
L435:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v2134)+4))
	v2217 = int32(0)
	v2219 = m.G0
	v2221 = v2219 - int32(32)
	m.G0 = v2221
	v2224 = F_SearchSysCache1(m, int32(47), v2216)
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		goto L1
	} else {
		goto L439
	}
L436:
	;
	v2420 = int32(0)
	if v2051|base.B2i32(v2342 != v2420) == v2420 {
		goto L432
	} else {
		goto L484
	}
L437:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L1
	} else {
		goto L481
	}
L438:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L1
	} else {
		goto L478
	}
L439:
	;
	if v2224 != 0 {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v2228 = F_heap_attisnull(m, v2224, int32(22), int32(0))
	mBase = m.M
	v2229 = m.ExcPending
	if v2229 != 0 {
		goto L1
	} else {
		goto L444
	}
L441:
	;
	goto L442
L442:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L1
	} else {
		goto L475
	}
L443:
	;
	F_ReleaseCatCache(m, v2224)
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L1
	} else {
		goto L474
	}
L444:
	;
	if v2228 != 0 {
		v2342 = v2217
		goto L443
	} else {
		goto L445
	}
L445:
	;
	v2232 = F_heap_attisnull(m, v2224, int32(23), int32(0))
	mBase = m.M
	v2233 = m.ExcPending
	if v2233 != 0 {
		goto L1
	} else {
		goto L446
	}
L446:
	;
	if v2232 != 0 {
		v2342 = v2217
		goto L443
	} else {
		goto L447
	}
L447:
	;
	v2236 = F_SysCacheGetAttrNotNull(m, int32(47), v2224, int32(22))
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	v2240 = F_SysCacheGetAttrNotNull(m, int32(47), v2224, int32(23))
	mBase = m.M
	v2241 = m.ExcPending
	if v2241 != 0 {
		goto L1
	} else {
		goto L449
	}
L449:
	;
	v2242 = F_pg_detoast_datum(m, v2236)
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L1
	} else {
		goto L450
	}
L450:
	;
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v2242)+4))
	if v2244 != int32(1) {
		goto L438
	} else {
		goto L451
	}
L451:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(v2242)+16))
	if v2247 < int32(0) {
		goto L438
	} else {
		goto L452
	}
L452:
	;
	v2250 = *(*int32)(unsafe.Add(mBase, uint32(v2242)+8))
	if v2250 != 0 {
		goto L438
	} else {
		goto L453
	}
L453:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v2242)+12))
	if v2251 != int32(18) {
		goto L438
	} else {
		goto L454
	}
L454:
	;
	v2254 = F_pg_detoast_datum(m, v2240)
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(v2254)+4))
	if v2256 != int32(1) {
		goto L437
	} else {
		goto L456
	}
L456:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v2254)+16))
	if v2259 != v2247 {
		goto L437
	} else {
		goto L457
	}
L457:
	;
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(v2254)+8))
	if v2261 != 0 {
		goto L437
	} else {
		goto L458
	}
L458:
	;
	v2262 = *(*int32)(unsafe.Add(mBase, uint32(v2254)+12))
	if v2262 != int32(25) {
		goto L437
	} else {
		goto L459
	}
L459:
	;
	v2265 = int32(0)
	F_deconstruct_array_builtin(m, v2254, int32(25), v2221+int32(28), v2265, v2221+int32(24))
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L1
	} else {
		goto L460
	}
L460:
	;
	if v2247 == int32(0) {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v2342 = int32(0)
	goto L443
L462:
	;
	goto L463
L463:
	;
	v2289 = v2265
	v2295 = int32(0)
	v2306 = v2217
	goto L464
L464:
	;
	v2312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2242+int32(24)+v2295))))
	v2314 = v2312 - int32(105)
	v2315 = int32(0)
	if base.B2i32(v2314 == v2315)|base.B2i32(v2314 == int32(13)) == v2315 {
		goto L466
	} else {
		goto L467
	}
L465:
	;
	v2342 = v2337
	goto L443
L466:
	;
	v2322 = int32(0)
	if v2306 != 0 {
		v2342 = v2322
		goto L443
	} else {
		goto L469
	}
L467:
	;
	v2337 = v2289
	v2338 = v2306
	goto L468
L468:
	;
	v2340 = v2295 + int32(1)
	if v2340 != v2247 {
		v2289 = v2337
		v2295 = v2340
		v2306 = v2338
		goto L464
	} else {
		goto L473
	}
L469:
	;
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(v2221)+28))
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(v2323+v2295<<(uint(int32(2))%32))))
	v2328 = F_text_to_cstring(m, v2327)
	mBase = m.M
	v2329 = m.ExcPending
	if v2329 != 0 {
		goto L1
	} else {
		goto L470
	}
L470:
	;
	if v2328 == int32(0) {
		v2342 = v2322
		goto L443
	} else {
		goto L471
	}
L471:
	;
	v2333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2328))))
	if v2333 == int32(0) {
		v2342 = v2322
		goto L443
	} else {
		goto L472
	}
L472:
	;
	v2337 = v2328
	v2338 = int32(1)
	goto L468
L473:
	;
	goto L465
L474:
	;
	m.G0 = v2221 + int32(32)
	goto L436
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2221))) = v2216
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_18), v2221)
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L1
	} else {
		goto L476
	}
L476:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_19), int32(1624), int32(_a_F_transformFromClauseItem_20))
	mBase = m.M
	v2390 = m.ExcPending
	if v2390 != 0 {
		goto L1
	} else {
		goto L477
	}
L477:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L478:
	;
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_21), int32(0))
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
		goto L1
	} else {
		goto L479
	}
L479:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_19), int32(1650), int32(_a_F_transformFromClauseItem_20))
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2221)+16)) = v2247
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_22), v2221+int32(16))
	mBase = m.M
	v2414 = m.ExcPending
	if v2414 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_19), int32(1658), int32(_a_F_transformFromClauseItem_20))
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L1
	} else {
		goto L483
	}
L483:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L484:
	;
	if v2342 != 0 {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v2425 = v2342
	goto L487
L486:
	;
	v2425 = v2133
	goto L487
L487:
	;
	v2481 = v2425
	goto L431
L488:
	;
	goto L432
L489:
	;
	F_TupleDescInitEntry(m, v2208, int32(1), v2481, v2490, v2491, int32(0))
	mBase = m.M
	v2495 = m.ExcPending
	if v2495 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	v2496 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+76))
	v2498 = F_exprCollation(m, v2134)
	mBase = m.M
	v2499 = m.ExcPending
	if v2499 != 0 {
		goto L1
	} else {
		goto L491
	}
L491:
	;
	v2500 = *(*int32)(unsafe.Add(mBase, uint32(v2496)))
	*(*int32)(unsafe.Add(mBase, uint32(v2496+v2500<<(uint(int32(4))%32)+int32(100))+16)) = v2498
	goto L492
L492:
	;
	goto L406
L493:
	;
	v2511 = F_CreateTemplateTupleDesc(m, v2508)
	mBase = m.M
	v2512 = m.ExcPending
	if v2512 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+76)) = v2511
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+4))
	if int32(0) < v2514 {
		goto L495
	} else {
		goto L496
	}
L495:
	;
	v2519 = int32(1)
	v2526 = int32(0)
	goto L498
L496:
	;
	v2643 = v2511
	goto L497
L497:
	;
	F_CheckAttributeNamesTypes(m, v2643, int32(99), int32(2))
	mBase = m.M
	v2647 = m.ExcPending
	if v2647 != 0 {
		goto L1
	} else {
		goto L512
	}
L498:
	;
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+12))
	v2554 = *(*int32)(unsafe.Add(mBase, uint32(v2550+v2526<<(uint(int32(2))%32))))
	v2555 = *(*int32)(unsafe.Add(mBase, uint32(v2554)+4))
	v2556 = *(*int32)(unsafe.Add(mBase, uint32(v2554)+8))
	v2557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2556)+12)))
	if v2557 != 0 {
		goto L379
	} else {
		goto L500
	}
L499:
	;
	v2611 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+76))
	v2643 = v2611
	goto L497
L500:
	;
	F_typenameTypeIdAndMod(m, l0, v2556, v2019+int32(68), v2019-int32(-64))
	mBase = m.M
	v2563 = m.ExcPending
	if v2563 != 0 {
		goto L1
	} else {
		goto L501
	}
L501:
	;
	v2564 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+68))
	v2565 = F_GetColumnDefCollation(m, l0, v2554, v2564)
	mBase = m.M
	v2566 = m.ExcPending
	if v2566 != 0 {
		goto L1
	} else {
		goto L502
	}
L502:
	;
	v2567 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+76))
	v2568 = base.I32_extend16_s(v2519)
	v2569 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+68))
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+64))
	F_TupleDescInitEntry(m, v2567, v2568, v2555, v2569, v2570, int32(0))
	mBase = m.M
	v2573 = m.ExcPending
	if v2573 != 0 {
		goto L1
	} else {
		goto L503
	}
L503:
	;
	v2574 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+76))
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(v2574)))
	*(*int32)(unsafe.Add(mBase, uint32(v2574+v2575<<(uint(int32(4))%32)+v2568*int32(100))+16)) = v2565
	goto L504
L504:
	;
	v2583 = *(*int32)(unsafe.Add(mBase, uint32(v2136)+12))
	v2584 = F_pstrdup(m, v2555)
	mBase = m.M
	v2585 = m.ExcPending
	if v2585 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	v2586 = F_makeString(m, v2584)
	mBase = m.M
	v2587 = m.ExcPending
	if v2587 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	v2588 = F_lappend(m, v2583, v2586)
	mBase = m.M
	v2589 = m.ExcPending
	if v2589 != 0 {
		goto L1
	} else {
		goto L507
	}
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2136)+12)) = v2588
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(v2136)+16))
	v2592 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+68))
	v2593 = F_lappend_oid(m, v2591, v2592)
	mBase = m.M
	v2594 = m.ExcPending
	if v2594 != 0 {
		goto L1
	} else {
		goto L508
	}
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2136)+16)) = v2593
	v2596 = *(*int32)(unsafe.Add(mBase, uint32(v2136)+20))
	v2597 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+64))
	v2598 = F_lappend_int(m, v2596, v2597)
	mBase = m.M
	v2599 = m.ExcPending
	if v2599 != 0 {
		goto L1
	} else {
		goto L509
	}
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2136)+20)) = v2598
	v2601 = *(*int32)(unsafe.Add(mBase, uint32(v2136)+24))
	v2602 = F_lappend_oid(m, v2601, v2565)
	mBase = m.M
	v2603 = m.ExcPending
	if v2603 != 0 {
		goto L1
	} else {
		goto L510
	}
L510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2136)+24)) = v2602
	v2605 = int32(1)
	v2608 = v2526 + v2605
	v2609 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+4))
	if v2608 < v2609 {
		v2519 = v2519 + v2605
		v2526 = v2608
		goto L498
	} else {
		goto L511
	}
L511:
	;
	goto L499
L512:
	;
	goto L406
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2022)+68)) = v2683
	v2687 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v2130+v2054))) = v2687
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(v2687)))
	v2061 = v2691 + v2061
	v2075 = v2075 + int32(1)
	goto L386
L514:
	;
	v2696 = F_CreateTemplateTupleDesc(m, v2693)
	mBase = m.M
	v2697 = m.ExcPending
	if v2697 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+76)) = v2696
	if int32(0) < v2028 {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v2702 = int32(0)
	v2704 = v2702
	v2713 = v2702
	goto L519
L517:
	;
	v2827 = v2696
	v2830 = int32(1)
	goto L518
L518:
	;
	v2854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v2854 != int32(1) {
		v2869 = v2827
		goto L385
	} else {
		goto L529
	}
L519:
	;
	v2738 = v2054 + v2713<<(uint(int32(2))%32)
	v2739 = *(*int32)(unsafe.Add(mBase, uint32(v2738)))
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v2739)))
	if int32(0) < v2740 {
		goto L521
	} else {
		goto L522
	}
L520:
	;
	v2822 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+76))
	v2827 = v2822
	v2830 = v2786 + int32(1)
	goto L518
L521:
	;
	v2743 = v2704
	v2747 = int32(1)
	v2750 = v2739
	goto L524
L522:
	;
	v2786 = v2704
	goto L523
L523:
	;
	v2818 = v2713 + int32(1)
	if v2818 != v2028 {
		v2704 = v2786
		v2713 = v2818
		goto L519
	} else {
		goto L528
	}
L524:
	;
	v2774 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+76))
	v2776 = v2743 + int32(1)
	F_TupleDescCopyEntry(m, v2774, base.I32_extend16_s(v2776), v2750, base.I32_extend16_s(v2747))
	mBase = m.M
	v2780 = m.ExcPending
	if v2780 != 0 {
		goto L1
	} else {
		goto L526
	}
L525:
	;
	v2786 = v2776
	goto L523
L526:
	;
	v2782 = v2747 + int32(1)
	v2783 = *(*int32)(unsafe.Add(mBase, uint32(v2738)))
	v2784 = *(*int32)(unsafe.Add(mBase, uint32(v2783)))
	if v2782 <= v2784 {
		v2743 = v2776
		v2747 = v2782
		v2750 = v2783
		goto L524
	} else {
		goto L527
	}
L527:
	;
	goto L525
L528:
	;
	goto L520
L529:
	;
	F_TupleDescInitEntry(m, v2827, base.I32_extend16_s(v2830), int32(_a_F_transformFromClauseItem_23), int32(20), int32(-1), int32(0))
	mBase = m.M
	v2863 = m.ExcPending
	if v2863 != 0 {
		goto L1
	} else {
		goto L530
	}
L530:
	;
	v2864 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+76))
	v2869 = v2864
	goto L385
L531:
	;
	v2898 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2022)+125)) = uint8(v2898)
	*(*uint8)(unsafe.Add(mBase, uint32(v2022)+124)) = uint8(v2015)
	v2901 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2902 = F_lappend(m, v2901, v2022)
	mBase = m.M
	v2903 = m.ExcPending
	if v2903 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2902
	v2905 = int32(0)
	if v2902 != 0 {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	v2907 = *(*int32)(unsafe.Add(mBase, uint32(v2902)+4))
	v2908 = v2907
	goto L535
L534:
	;
	v2908 = v2905
	goto L535
L535:
	;
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+76))
	v2910 = *(*int32)(unsafe.Add(mBase, uint32(v2909)))
	v2913 = F_palloc0(m, v2910<<(uint(int32(5))%32))
	mBase = m.M
	v2914 = m.ExcPending
	if v2914 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	if int32(0) < v2910 {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	v2921 = v2905
	goto L540
L538:
	;
	goto L539
L539:
	;
	v3013 = F_palloc(m, int32(28))
	mBase = m.M
	v3014 = m.ExcPending
	if v3014 != 0 {
		goto L1
	} else {
		goto L546
	}
L540:
	;
	v2948 = *(*int32)(unsafe.Add(mBase, uint32(v2909)))
	v2954 = v2909 + v2948<<(uint(int32(4))%32) + v2921*int32(100)
	v2955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2954)+111)))
	if v2955 == int32(0) {
		goto L542
	} else {
		goto L543
	}
L541:
	;
	goto L539
L542:
	;
	v2960 = v2913 + v2921<<(uint(int32(5))%32)
	v2962 = v2921 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2960)+4)) = uint16(v2962)
	*(*int32)(unsafe.Add(mBase, uint32(v2960))) = v2908
	v2966 = v2954 + int32(20)
	v2967 = *(*int32)(unsafe.Add(mBase, uint32(v2966)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v2960)+8)) = v2967
	v2969 = *(*int32)(unsafe.Add(mBase, uint32(v2966)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v2960)+12)) = v2969
	v2971 = *(*int32)(unsafe.Add(mBase, uint32(v2966)+96))
	*(*uint16)(unsafe.Add(mBase, uint32(v2960)+28)) = uint16(v2962)
	*(*int32)(unsafe.Add(mBase, uint32(v2960)+24)) = v2908
	*(*int32)(unsafe.Add(mBase, uint32(v2960)+16)) = v2971
	goto L544
L543:
	;
	goto L544
L544:
	;
	v2979 = v2921 + int32(1)
	if v2979 != v2910 {
		v2921 = v2979
		goto L540
	} else {
		goto L545
	}
L545:
	;
	goto L541
L546:
	;
	v3015 = *(*int32)(unsafe.Add(mBase, uint32(v2022)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3013)+20)) = int64(16777473)
	*(*int32)(unsafe.Add(mBase, uint32(v3013)+16)) = v2913
	*(*int32)(unsafe.Add(mBase, uint32(v3013)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3013)+8)) = v2908
	*(*int32)(unsafe.Add(mBase, uint32(v3013)+4)) = v2022
	*(*int32)(unsafe.Add(mBase, uint32(v3013))) = v3015
	m.G0 = v2019 + int32(80)
	goto L365
L547:
	;
	v3031 = F_exprLocation(m, v2132)
	mBase = m.M
	F_parser_errposition(m, l0, v3031)
	mBase = m.M
	v3033 = m.ExcPending
	if v3033 != 0 {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1852), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v3038 = m.ExcPending
	if v3038 != 0 {
		goto L1
	} else {
		goto L549
	}
L549:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L550:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3045 = m.ExcPending
	if v3045 != 0 {
		goto L1
	} else {
		goto L551
	}
L551:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_24), int32(0))
	mBase = m.M
	v3049 = m.ExcPending
	if v3049 != 0 {
		goto L1
	} else {
		goto L552
	}
L552:
	;
	v3050 = F_exprLocation(m, v2134)
	mBase = m.M
	F_parser_errposition(m, l0, v3050)
	mBase = m.M
	v3052 = m.ExcPending
	if v3052 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1875), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v3057 = m.ExcPending
	if v3057 != 0 {
		goto L1
	} else {
		goto L554
	}
L554:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L555:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v3064 = m.ExcPending
	if v3064 != 0 {
		goto L1
	} else {
		goto L556
	}
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+32)) = int32(1600)
	F_errmsg(m, int32(_a_F_transformFromClauseItem_25), v2019+int32(32))
	mBase = m.M
	v3071 = m.ExcPending
	if v3071 != 0 {
		goto L1
	} else {
		goto L557
	}
L557:
	;
	v3072 = F_exprLocation(m, v2132)
	mBase = m.M
	F_parser_errposition(m, l0, v3072)
	mBase = m.M
	v3074 = m.ExcPending
	if v3074 != 0 {
		goto L1
	} else {
		goto L558
	}
L558:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1914), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v3079 = m.ExcPending
	if v3079 != 0 {
		goto L1
	} else {
		goto L559
	}
L559:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L560:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v3086 = m.ExcPending
	if v3086 != 0 {
		goto L1
	} else {
		goto L561
	}
L561:
	;
	v3087 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+72))
	v3088 = F_format_type_be(m, v3087)
	mBase = m.M
	v3089 = m.ExcPending
	if v3089 != 0 {
		goto L1
	} else {
		goto L562
	}
L562:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+20)) = v3088
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+16)) = v2133
	F_errmsg(m, int32(_a_F_transformFromClauseItem_26), v2019+int32(16))
	mBase = m.M
	v3096 = m.ExcPending
	if v3096 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	v3097 = F_exprLocation(m, v2134)
	mBase = m.M
	F_parser_errposition(m, l0, v3097)
	mBase = m.M
	v3099 = m.ExcPending
	if v3099 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1973), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v3104 = m.ExcPending
	if v3104 != 0 {
		goto L1
	} else {
		goto L565
	}
L565:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L566:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v3111 = m.ExcPending
	if v3111 != 0 {
		goto L1
	} else {
		goto L567
	}
L567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2019))) = int32(1664)
	F_errmsg(m, int32(_a_F_transformFromClauseItem_27), v2019)
	mBase = m.M
	v3116 = m.ExcPending
	if v3116 != 0 {
		goto L1
	} else {
		goto L568
	}
L568:
	;
	v3117 = F_exprLocation(m, v1939)
	mBase = m.M
	F_parser_errposition(m, l0, v3117)
	mBase = m.M
	v3119 = m.ExcPending
	if v3119 != 0 {
		goto L1
	} else {
		goto L569
	}
L569:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(2001), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v3124 = m.ExcPending
	if v3124 != 0 {
		goto L1
	} else {
		goto L570
	}
L570:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L571:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v3131 = m.ExcPending
	if v3131 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2019)+48)) = v2555
	F_errmsg(m, int32(_a_F_transformFromClauseItem_28), v2019+int32(48))
	mBase = m.M
	v3137 = m.ExcPending
	if v3137 != 0 {
		goto L1
	} else {
		goto L573
	}
L573:
	;
	v3138 = *(*int32)(unsafe.Add(mBase, uint32(v2554)+64))
	F_parser_errposition(m, l0, v3138)
	mBase = m.M
	v3140 = m.ExcPending
	if v3140 != 0 {
		goto L1
	} else {
		goto L574
	}
L574:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1931), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v3145 = m.ExcPending
	if v3145 != 0 {
		goto L1
	} else {
		goto L575
	}
L575:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L576:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v3152
	v3156 = F_palloc0(m, int32(8))
	mBase = m.M
	v3157 = m.ExcPending
	if v3157 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3156))) = int32(63)
	v3160 = *(*int32)(unsafe.Add(mBase, uint32(v3013)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3156)+4)) = v3160
	v6521 = v3156
	goto L3
L578:
	;
	v6503 = v3274
	goto L4
L579:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3282 = m.ExcPending
	if v3282 != 0 {
		goto L1
	} else {
		goto L602
	}
L580:
	;
	v3181 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3164)+60)) = v3181
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(v3166)+8))
	if v3183 == v3181 {
		goto L584
	} else {
		goto L585
	}
L581:
	;
	v3173 = *(*int32)(unsafe.Add(mBase, uint32(v3170)+4))
	if base.Ui32(v3173-int32(1)) < base.Ui32(int32(2)) {
		goto L580
	} else {
		goto L582
	}
L582:
	;
	if v3173 != int32(6) {
		goto L579
	} else {
		goto L583
	}
L583:
	;
	goto L580
L584:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3164)+60)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3164)+16)) = int32(0)
	v3191 = v3164 - int32(-64)
	v3196 = F_pg_snprintf(m, v3191, int32(32), int32(_a_F_transformFromClauseItem_29), v3164+int32(16))
	mBase = m.M
	v3197 = m.ExcPending
	if v3197 != 0 {
		goto L1
	} else {
		goto L587
	}
L585:
	;
	v3205 = v3183
	goto L586
L586:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3164)+12)) = v3205
	*(*int32)(unsafe.Add(mBase, uint32(v3164)+40)) = v3205
	v3211 = F_list_make1_impl(m, int32(1), v3164+int32(12))
	mBase = m.M
	v3212 = m.ExcPending
	if v3212 != 0 {
		goto L1
	} else {
		goto L590
	}
L587:
	;
	v3199 = F_pstrdup(m, v3191)
	mBase = m.M
	v3200 = m.ExcPending
	if v3200 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	v3201 = F_lappend(m, int32(0), v3199)
	mBase = m.M
	v3202 = m.ExcPending
	if v3202 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3166)+8)) = v3199
	v3205 = v3199
	goto L586
L590:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3164)+56)) = v3211
	v3215 = v3164 + int32(44)
	v3216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_CheckDuplicateColumnOrPathNames(m, v3215, v3216)
	mBase = m.M
	v3218 = m.ExcPending
	if v3218 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	v3219 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v3219)
	v3222 = F_palloc0(m, int32(72))
	mBase = m.M
	v3223 = m.ExcPending
	if v3223 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3222))) = int64(4294967300)
	v3227 = F_palloc0(m, int32(48))
	mBase = m.M
	v3228 = m.ExcPending
	if v3228 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3227))) = int64(12884902010)
	v3231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3227)+12)) = v3231
	v3233 = *(*int32)(unsafe.Add(mBase, uint32(v3166)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3227)+16)) = v3233
	v3235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3227)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3227)+20)) = v3235
	v3239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v3227)+32)) = v3239
	v3241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3227)+44)) = v3241
	v3244 = F_transformExpr(m, l0, v3227, int32(5))
	mBase = m.M
	v3245 = m.ExcPending
	if v3245 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3222)+16)) = v3244
	*(*int32)(unsafe.Add(mBase, uint32(v3164)+52)) = v3222
	*(*int32)(unsafe.Add(mBase, uint32(v3164)+48)) = l1
	v3249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3250 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3251 = F_transformJsonTableColumns(m, v3215, v3249, v3250, v3166)
	mBase = m.M
	v3252 = m.ExcPending
	if v3252 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3222)+60)) = v3251
	v3254 = *(*int32)(unsafe.Add(mBase, uint32(v3222)+16))
	v3255 = *(*int32)(unsafe.Add(mBase, uint32(v3254)+32))
	v3256 = F_copyObjectImpl(m, v3255)
	mBase = m.M
	v3257 = m.ExcPending
	if v3257 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3222)+64)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v3222)+52)) = v3256
	v3261 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3222)+68)) = v3261
	v3263 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v3263)
	v3266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	if v3266 == v3263 {
		goto L597
	} else {
		goto L598
	}
L597:
	;
	v3270 = F_contain_vars_of_level(m, v3222, int32(0))
	mBase = m.M
	v3271 = m.ExcPending
	if v3271 != 0 {
		goto L1
	} else {
		goto L600
	}
L598:
	;
	v3272 = int32(1)
	goto L599
L599:
	;
	v3273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v3274 = F_addRangeTableEntryForTableFunc(m, l0, v3222, v3273, v3272)
	mBase = m.M
	v3275 = m.ExcPending
	if v3275 != 0 {
		goto L1
	} else {
		goto L601
	}
L600:
	;
	v3272 = v3270
	goto L599
L601:
	;
	m.G0 = v3164 + int32(96)
	goto L578
L602:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3285 = m.ExcPending
	if v3285 != 0 {
		goto L1
	} else {
		goto L603
	}
L603:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3164)+32)) = int32(_a_F_transformFromClauseItem_30)
	F_errmsg(m, int32(_a_F_transformFromClauseItem_31), v3164+int32(32))
	mBase = m.M
	v3292 = m.ExcPending
	if v3292 != 0 {
		goto L1
	} else {
		goto L604
	}
L604:
	;
	F_errdetail(m, int32(_a_F_transformFromClauseItem_32), int32(0))
	mBase = m.M
	v3296 = m.ExcPending
	if v3296 != 0 {
		goto L1
	} else {
		goto L605
	}
L605:
	;
	v3297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v3298 = *(*int32)(unsafe.Add(mBase, uint32(v3297)+16))
	F_parser_errposition(m, l0, v3298)
	mBase = m.M
	v3300 = m.ExcPending
	if v3300 != 0 {
		goto L1
	} else {
		goto L606
	}
L606:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_33), int32(94), int32(_a_F_transformFromClauseItem_34))
	mBase = m.M
	v3305 = m.ExcPending
	if v3305 != 0 {
		goto L1
	} else {
		goto L607
	}
L607:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L608:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3307))) = int64(4)
	v3311 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v3311)
	v3313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3315 = F_transformExpr(m, l0, v3313, int32(5))
	mBase = m.M
	v3316 = m.ExcPending
	if v3316 != 0 {
		goto L1
	} else {
		goto L609
	}
L609:
	;
	v3319 = F_coerce_to_specific_type(m, l0, v3315, int32(25), int32(_a_F_transformFromClauseItem_35))
	mBase = m.M
	v3320 = m.ExcPending
	if v3320 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+20)) = v3319
	F_assign_expr_collations(m, l0, v3319)
	mBase = m.M
	v3323 = m.ExcPending
	if v3323 != 0 {
		goto L1
	} else {
		goto L611
	}
L611:
	;
	v3324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v3326 = F_transformExpr(m, l0, v3324, int32(5))
	mBase = m.M
	v3327 = m.ExcPending
	if v3327 != 0 {
		goto L1
	} else {
		goto L612
	}
L612:
	;
	v3330 = F_coerce_to_specific_type(m, l0, v3326, int32(142), int32(_a_F_transformFromClauseItem_35))
	mBase = m.M
	v3331 = m.ExcPending
	if v3331 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+16)) = v3330
	F_assign_expr_collations(m, l0, v3330)
	mBase = m.M
	v3334 = m.ExcPending
	if v3334 != 0 {
		goto L1
	} else {
		goto L614
	}
L614:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+64)) = int32(-1)
	v3337 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v3337 != 0 {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	v3338 = *(*int32)(unsafe.Add(mBase, uint32(v3337)+4))
	v3342 = v3338 << (uint(int32(2)) % 32)
	goto L617
L616:
	;
	v3342 = int32(0)
	goto L617
L617:
	;
	v3343 = F_palloc(m, v3342)
	mBase = m.M
	v3344 = m.ExcPending
	if v3344 != 0 {
		goto L1
	} else {
		goto L618
	}
L618:
	;
	v3345 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v3345 == int32(0) {
		goto L5
	} else {
		goto L619
	}
L619:
	;
	v3348 = *(*int32)(unsafe.Add(mBase, uint32(v3345)+4))
	if v3348 <= int32(0) {
		goto L5
	} else {
		goto L620
	}
L620:
	;
	v3361 = v5
	goto L621
L621:
	;
	v3382 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+24))
	v3384 = v3361 << (uint(int32(2)) % 32)
	v3385 = *(*int32)(unsafe.Add(mBase, uint32(v3345)+12))
	v3387 = *(*int32)(unsafe.Add(mBase, uint32(v3384+v3385)))
	v3388 = *(*int32)(unsafe.Add(mBase, uint32(v3387)+4))
	v3389 = F_pstrdup(m, v3388)
	mBase = m.M
	v3390 = m.ExcPending
	if v3390 != 0 {
		goto L1
	} else {
		goto L623
	}
L622:
	;
	goto L5
L623:
	;
	v3391 = F_makeString(m, v3389)
	mBase = m.M
	v3392 = m.ExcPending
	if v3392 != 0 {
		goto L1
	} else {
		goto L624
	}
L624:
	;
	v3393 = F_lappend(m, v3382, v3391)
	mBase = m.M
	v3394 = m.ExcPending
	if v3394 != 0 {
		goto L1
	} else {
		goto L625
	}
L625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+24)) = v3393
	v3396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3387)+12)))
	if v3396 != 0 {
		goto L630
	} else {
		goto L631
	}
L626:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3384+v3343))) = v3475
	v3640 = v3361 + int32(1)
	v3641 = *(*int32)(unsafe.Add(mBase, uint32(v3345)+4))
	if v3640 < v3641 {
		v3361 = v3640
		goto L621
	} else {
		goto L687
	}
L627:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3587 = m.ExcPending
	if v3587 != 0 {
		goto L1
	} else {
		goto L682
	}
L628:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3568 = m.ExcPending
	if v3568 != 0 {
		goto L1
	} else {
		goto L677
	}
L629:
	;
	v3419 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+28))
	v3420 = F_lappend_oid(m, v3419, v3418)
	mBase = m.M
	v3421 = m.ExcPending
	if v3421 != 0 {
		goto L1
	} else {
		goto L636
	}
L630:
	;
	v3397 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+64))
	if v3397 != int32(-1) {
		goto L628
	} else {
		goto L633
	}
L631:
	;
	goto L632
L632:
	;
	v3406 = *(*int32)(unsafe.Add(mBase, uint32(v3387)+8))
	v3407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3406)+12)))
	if v3407 == int32(1) {
		goto L627
	} else {
		goto L634
	}
L633:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+328)) = int32(-1)
	v3402 = int32(23)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+332)) = v3402
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+64)) = v3361
	v3418 = v3402
	goto L629
L634:
	;
	F_typenameTypeIdAndMod(m, l0, v3406, v34+int32(332), v34+int32(328))
	mBase = m.M
	v3415 = m.ExcPending
	if v3415 != 0 {
		goto L1
	} else {
		goto L635
	}
L635:
	;
	v3416 = *(*int32)(unsafe.Add(mBase, uint32(v34)+332))
	v3418 = v3416
	goto L629
L636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+28)) = v3420
	v3423 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+32))
	v3424 = *(*int32)(unsafe.Add(mBase, uint32(v34)+328))
	v3425 = F_lappend_int(m, v3423, v3424)
	mBase = m.M
	v3426 = m.ExcPending
	if v3426 != 0 {
		goto L1
	} else {
		goto L637
	}
L637:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+32)) = v3425
	v3428 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+36))
	v3429 = *(*int32)(unsafe.Add(mBase, uint32(v34)+332))
	v3430 = F_get_typcollation(m, v3429)
	mBase = m.M
	v3431 = m.ExcPending
	if v3431 != 0 {
		goto L1
	} else {
		goto L638
	}
L638:
	;
	v3432 = F_lappend_oid(m, v3428, v3430)
	mBase = m.M
	v3433 = m.ExcPending
	if v3433 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+36)) = v3432
	v3435 = int32(0)
	v3437 = *(*int32)(unsafe.Add(mBase, uint32(v3387)+16))
	if v3437 != 0 {
		goto L640
	} else {
		goto L641
	}
L640:
	;
	v3439 = F_transformExpr(m, l0, v3437, int32(5))
	mBase = m.M
	v3440 = m.ExcPending
	if v3440 != 0 {
		goto L1
	} else {
		goto L643
	}
L641:
	;
	v3447 = v3435
	goto L642
L642:
	;
	v3448 = *(*int32)(unsafe.Add(mBase, uint32(v3387)+20))
	if v3448 != 0 {
		goto L646
	} else {
		goto L647
	}
L643:
	;
	v3443 = F_coerce_to_specific_type(m, l0, v3439, int32(25), int32(_a_F_transformFromClauseItem_35))
	mBase = m.M
	v3444 = m.ExcPending
	if v3444 != 0 {
		goto L1
	} else {
		goto L644
	}
L644:
	;
	F_assign_expr_collations(m, l0, v3443)
	mBase = m.M
	v3446 = m.ExcPending
	if v3446 != 0 {
		goto L1
	} else {
		goto L645
	}
L645:
	;
	v3447 = v3443
	goto L642
L646:
	;
	v3450 = F_transformExpr(m, l0, v3448, int32(5))
	mBase = m.M
	v3451 = m.ExcPending
	if v3451 != 0 {
		goto L1
	} else {
		goto L649
	}
L647:
	;
	v3459 = v3435
	goto L648
L648:
	;
	v3460 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+40))
	v3461 = F_lappend(m, v3460, v3447)
	mBase = m.M
	v3462 = m.ExcPending
	if v3462 != 0 {
		goto L1
	} else {
		goto L652
	}
L649:
	;
	v3452 = *(*int32)(unsafe.Add(mBase, uint32(v34)+332))
	v3453 = *(*int32)(unsafe.Add(mBase, uint32(v34)+328))
	v3455 = F_coerce_to_specific_type_typmod(m, l0, v3450, v3452, v3453, int32(_a_F_transformFromClauseItem_35))
	mBase = m.M
	v3456 = m.ExcPending
	if v3456 != 0 {
		goto L1
	} else {
		goto L650
	}
L650:
	;
	F_assign_expr_collations(m, l0, v3455)
	mBase = m.M
	v3458 = m.ExcPending
	if v3458 != 0 {
		goto L1
	} else {
		goto L651
	}
L651:
	;
	v3459 = v3455
	goto L648
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+40)) = v3461
	v3464 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+44))
	v3465 = F_lappend(m, v3464, v3459)
	mBase = m.M
	v3466 = m.ExcPending
	if v3466 != 0 {
		goto L1
	} else {
		goto L653
	}
L653:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+44)) = v3465
	v3468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3387)+13)))
	if v3468 == int32(1) {
		goto L654
	} else {
		goto L655
	}
L654:
	;
	v3471 = *(*int32)(unsafe.Add(mBase, uint32(v3307)+56))
	v3472 = F_bms_add_member(m, v3471, v3361)
	mBase = m.M
	v3473 = m.ExcPending
	if v3473 != 0 {
		goto L1
	} else {
		goto L657
	}
L655:
	;
	goto L656
L656:
	;
	v3475 = *(*int32)(unsafe.Add(mBase, uint32(v3387)+4))
	v3476 = int32(0)
	if v3361 == v3476 {
		goto L626
	} else {
		goto L658
	}
L657:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+56)) = v3472
	goto L656
L658:
	;
	v3483 = v3476
	goto L659
L659:
	;
	v3513 = *(*int32)(unsafe.Add(mBase, uint32(v3343+v3483<<(uint(int32(2))%32))))
	v3516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3513))))
	v3519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3475))))
	if base.B2i32(v3516 == int32(0))|base.B2i32(v3516 != v3519) != 0 {
		v3537 = v3516
		v3538 = v3519
		goto L662
	} else {
		goto L663
	}
L660:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3546 = m.ExcPending
	if v3546 != 0 {
		goto L1
	} else {
		goto L672
	}
L661:
	;
	if v3537-v3538 != 0 {
		goto L668
	} else {
		goto L669
	}
L662:
	;
	goto L661
L663:
	;
	v3522 = v3513
	v3523 = v3475
	goto L664
L664:
	;
	v3526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3523)+1)))
	v3527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3522)+1)))
	if v3527 == int32(0) {
		v3537 = v3527
		v3538 = v3526
		goto L662
	} else {
		goto L666
	}
L665:
	;
	v3537 = v3527
	v3538 = v3526
	goto L662
L666:
	;
	v3530 = int32(1)
	if v3527 == v3526 {
		v3522 = v3522 + v3530
		v3523 = v3523 + v3530
		goto L664
	} else {
		goto L667
	}
L667:
	;
	goto L665
L668:
	;
	v3541 = v3483 + int32(1)
	if v3361 != v3541 {
		v3483 = v3541
		goto L659
	} else {
		goto L671
	}
L669:
	;
	goto L670
L670:
	;
	goto L660
L671:
	;
	goto L626
L672:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3549 = m.ExcPending
	if v3549 != 0 {
		goto L1
	} else {
		goto L673
	}
L673:
	;
	v3550 = *(*int32)(unsafe.Add(mBase, uint32(v3387)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+240)) = v3550
	F_errmsg(m, int32(_a_F_transformFromClauseItem_36), v34+int32(240))
	mBase = m.M
	v3556 = m.ExcPending
	if v3556 != 0 {
		goto L1
	} else {
		goto L674
	}
L674:
	;
	v3557 = *(*int32)(unsafe.Add(mBase, uint32(v3387)+24))
	F_parser_errposition(m, l0, v3557)
	mBase = m.M
	v3559 = m.ExcPending
	if v3559 != 0 {
		goto L1
	} else {
		goto L675
	}
L675:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(823), int32(_a_F_transformFromClauseItem_37))
	mBase = m.M
	v3564 = m.ExcPending
	if v3564 != 0 {
		goto L1
	} else {
		goto L676
	}
L676:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L677:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3571 = m.ExcPending
	if v3571 != 0 {
		goto L1
	} else {
		goto L678
	}
L678:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_38), int32(0))
	mBase = m.M
	v3575 = m.ExcPending
	if v3575 != 0 {
		goto L1
	} else {
		goto L679
	}
L679:
	;
	v3576 = *(*int32)(unsafe.Add(mBase, uint32(v3387)+24))
	F_parser_errposition(m, l0, v3576)
	mBase = m.M
	v3578 = m.ExcPending
	if v3578 != 0 {
		goto L1
	} else {
		goto L680
	}
L680:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(761), int32(_a_F_transformFromClauseItem_37))
	mBase = m.M
	v3583 = m.ExcPending
	if v3583 != 0 {
		goto L1
	} else {
		goto L681
	}
L681:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L682:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v3590 = m.ExcPending
	if v3590 != 0 {
		goto L1
	} else {
		goto L683
	}
L683:
	;
	v3591 = *(*int32)(unsafe.Add(mBase, uint32(v3387)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+256)) = v3591
	F_errmsg(m, int32(_a_F_transformFromClauseItem_28), v34+int32(256))
	mBase = m.M
	v3597 = m.ExcPending
	if v3597 != 0 {
		goto L1
	} else {
		goto L684
	}
L684:
	;
	v3598 = *(*int32)(unsafe.Add(mBase, uint32(v3387)+24))
	F_parser_errposition(m, l0, v3598)
	mBase = m.M
	v3600 = m.ExcPending
	if v3600 != 0 {
		goto L1
	} else {
		goto L685
	}
L685:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(774), int32(_a_F_transformFromClauseItem_37))
	mBase = m.M
	v3605 = m.ExcPending
	if v3605 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L687:
	;
	goto L622
L688:
	;
	v3647 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v3647
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_39), v34)
	mBase = m.M
	v3651 = m.ExcPending
	if v3651 != 0 {
		goto L1
	} else {
		goto L689
	}
L689:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1625), int32(_a_F_transformFromClauseItem_40))
	mBase = m.M
	v3656 = m.ExcPending
	if v3656 != 0 {
		goto L1
	} else {
		goto L690
	}
L690:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L691:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3663 = m.ExcPending
	if v3663 != 0 {
		goto L1
	} else {
		goto L692
	}
L692:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_41), int32(0))
	mBase = m.M
	v3667 = m.ExcPending
	if v3667 != 0 {
		goto L1
	} else {
		goto L693
	}
L693:
	;
	F_errhint(m, int32(_a_F_transformFromClauseItem_42), int32(0))
	mBase = m.M
	v3671 = m.ExcPending
	if v3671 != 0 {
		goto L1
	} else {
		goto L694
	}
L694:
	;
	v3672 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3673 = F_exprLocation(m, v3672)
	mBase = m.M
	F_parser_errposition(m, l0, v3673)
	mBase = m.M
	v3675 = m.ExcPending
	if v3675 != 0 {
		goto L1
	} else {
		goto L695
	}
L695:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(658), int32(_a_F_transformFromClauseItem_13))
	mBase = m.M
	v3680 = m.ExcPending
	if v3680 != 0 {
		goto L1
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
	F_errhint(m, int32(_a_F_transformFromClauseItem_43), int32(0))
	mBase = m.M
	v3688 = m.ExcPending
	if v3688 != 0 {
		goto L1
	} else {
		goto L698
	}
L698:
	;
	v3689 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3690 = F_exprLocation(m, v3689)
	mBase = m.M
	F_parser_errposition(m, l0, v3690)
	mBase = m.M
	v3692 = m.ExcPending
	if v3692 != 0 {
		goto L1
	} else {
		goto L699
	}
L699:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(643), int32(_a_F_transformFromClauseItem_13))
	mBase = m.M
	v3697 = m.ExcPending
	if v3697 != 0 {
		goto L1
	} else {
		goto L700
	}
L700:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L701:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3704 = m.ExcPending
	if v3704 != 0 {
		goto L1
	} else {
		goto L702
	}
L702:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_44), int32(0))
	mBase = m.M
	v3708 = m.ExcPending
	if v3708 != 0 {
		goto L1
	} else {
		goto L703
	}
L703:
	;
	v3709 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3710 = F_exprLocation(m, v3709)
	mBase = m.M
	F_parser_errposition(m, l0, v3710)
	mBase = m.M
	v3712 = m.ExcPending
	if v3712 != 0 {
		goto L1
	} else {
		goto L704
	}
L704:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(609), int32(_a_F_transformFromClauseItem_13))
	mBase = m.M
	v3717 = m.ExcPending
	if v3717 != 0 {
		goto L1
	} else {
		goto L705
	}
L705:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L706:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3724 = m.ExcPending
	if v3724 != 0 {
		goto L1
	} else {
		goto L707
	}
L707:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_45), int32(0))
	mBase = m.M
	v3728 = m.ExcPending
	if v3728 != 0 {
		goto L1
	} else {
		goto L708
	}
L708:
	;
	v3729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v3730 = F_exprLocation(m, v3729)
	mBase = m.M
	F_parser_errposition(m, l0, v3730)
	mBase = m.M
	v3732 = m.ExcPending
	if v3732 != 0 {
		goto L1
	} else {
		goto L709
	}
L709:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(597), int32(_a_F_transformFromClauseItem_13))
	mBase = m.M
	v3737 = m.ExcPending
	if v3737 != 0 {
		goto L1
	} else {
		goto L710
	}
L710:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L711:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3744 = m.ExcPending
	if v3744 != 0 {
		goto L1
	} else {
		goto L712
	}
L712:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_45), int32(0))
	mBase = m.M
	v3748 = m.ExcPending
	if v3748 != 0 {
		goto L1
	} else {
		goto L713
	}
L713:
	;
	v3749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v3750 = F_exprLocation(m, v3749)
	mBase = m.M
	F_parser_errposition(m, l0, v3750)
	mBase = m.M
	v3752 = m.ExcPending
	if v3752 != 0 {
		goto L1
	} else {
		goto L714
	}
L714:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(569), int32(_a_F_transformFromClauseItem_13))
	mBase = m.M
	v3757 = m.ExcPending
	if v3757 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_46), int32(0))
	mBase = m.M
	v3765 = m.ExcPending
	if v3765 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(446), int32(_a_F_transformFromClauseItem_47))
	mBase = m.M
	v3770 = m.ExcPending
	if v3770 != 0 {
		goto L1
	} else {
		goto L718
	}
L718:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L719:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v3776
	v3779 = *(*int32)(unsafe.Add(mBase, uint32(v34)+284))
	if v3779 == int32(0) {
		goto L720
	} else {
		goto L721
	}
L720:
	;
	v3863 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v3863 != 0 {
		goto L726
	} else {
		goto L727
	}
L721:
	;
	v3782 = *(*int32)(unsafe.Add(mBase, uint32(v3779)+4))
	if v3782 <= int32(0) {
		goto L720
	} else {
		goto L722
	}
L722:
	;
	v3785 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v3793 = int32(0)
	goto L723
L723:
	;
	v3820 = *(*int32)(unsafe.Add(mBase, uint32(v3779)+12))
	v3824 = *(*int32)(unsafe.Add(mBase, uint32(v3820+v3793<<(uint(int32(2))%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3824)+23)) = uint8(base.B2i32(base.Ui32(v3785) < base.Ui32(int32(2))))
	v3826 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3824)+22)) = uint8(v3826)
	v3829 = v3793 + v3826
	v3830 = *(*int32)(unsafe.Add(mBase, uint32(v3779)+4))
	if v3829 < v3830 {
		v3793 = v3829
		goto L723
	} else {
		goto L725
	}
L724:
	;
	goto L720
L725:
	;
	goto L724
L726:
	;
	v3864 = *(*int32)(unsafe.Add(mBase, uint32(v3863)+4))
	v3866 = v3864
	goto L728
L727:
	;
	v3866 = int32(0)
	goto L728
L728:
	;
	v3867 = F_list_concat(m, v3863, v3779)
	mBase = m.M
	v3868 = m.ExcPending
	if v3868 != 0 {
		goto L1
	} else {
		goto L729
	}
L729:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v3867
	v3870 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3875 = F_transformFromClauseItem(m, l0, v3870, v34+int32(288), v34+int32(280))
	mBase = m.M
	v3876 = m.ExcPending
	if v3876 != 0 {
		goto L1
	} else {
		goto L730
	}
L730:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v3875
	v3878 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3879 = int32(0)
	if base.B2i32(v3878 == v3879)|base.B2i32(v3866 <= v3879) != 0 {
		goto L732
	} else {
		goto L733
	}
L731:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v3889
	v3891 = *(*int32)(unsafe.Add(mBase, uint32(v34)+280))
	F_checkNameSpaceConflicts(m, v3779, v3891)
	mBase = m.M
	v3893 = m.ExcPending
	if v3893 != 0 {
		goto L1
	} else {
		goto L738
	}
L732:
	;
	v3889 = int32(0)
	goto L734
L733:
	;
	v3886 = *(*int32)(unsafe.Add(mBase, uint32(v3878)+4))
	if v3866 < v3886 {
		goto L735
	} else {
		goto L736
	}
L734:
	;
	goto L731
L735:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3878)+4)) = v3866
	goto L737
L736:
	;
	goto L737
L737:
	;
	v3889 = v3878
	goto L734
L738:
	;
	v3894 = F_list_concat(m, v3779, v3891)
	mBase = m.M
	v3895 = m.ExcPending
	if v3895 != 0 {
		goto L1
	} else {
		goto L739
	}
L739:
	;
	v3896 = *(*int32)(unsafe.Add(mBase, uint32(v34)+288))
	v3897 = *(*int32)(unsafe.Add(mBase, uint32(v3896)+16))
	v3898 = *(*int32)(unsafe.Add(mBase, uint32(v34)+292))
	v3899 = *(*int32)(unsafe.Add(mBase, uint32(v3898)+16))
	v3900 = *(*int32)(unsafe.Add(mBase, uint32(v3896)))
	v3901 = *(*int32)(unsafe.Add(mBase, uint32(v3900)+8))
	v3902 = *(*int32)(unsafe.Add(mBase, uint32(v3898)))
	v3903 = *(*int32)(unsafe.Add(mBase, uint32(v3902)+8))
	v3904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v3904 == int32(1) {
		goto L740
	} else {
		goto L741
	}
L740:
	;
	if v3903 == int32(0) {
		v4081 = v5
		goto L743
	} else {
		goto L744
	}
L741:
	;
	goto L742
L742:
	;
	v4133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v4133 != 0 {
		goto L771
	} else {
		goto L772
	}
L743:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v4081
	goto L742
L744:
	;
	v3909 = *(*int32)(unsafe.Add(mBase, uint32(v3903)+4))
	if v3909 <= int32(0) {
		v4081 = v5
		goto L743
	} else {
		goto L745
	}
L745:
	;
	v3918 = v5
	v3923 = v5
	goto L746
L746:
	;
	v3943 = *(*int32)(unsafe.Add(mBase, uint32(v3903)+12))
	v3947 = *(*int32)(unsafe.Add(mBase, uint32(v3943+v3918<<(uint(int32(2))%32))))
	v3948 = *(*int32)(unsafe.Add(mBase, uint32(v3947)+4))
	v3949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3948))))
	v3950 = int32(0)
	if base.B2i32(v3949 == v3950)|base.B2i32(v3901 == v3950) != 0 {
		v4046 = v3923
		goto L748
	} else {
		goto L749
	}
L747:
	;
	v4081 = v4046
	goto L743
L748:
	;
	v4067 = v3918 + int32(1)
	v4068 = *(*int32)(unsafe.Add(mBase, uint32(v3903)+4))
	if v4067 < v4068 {
		v3918 = v4067
		v3923 = v4046
		goto L746
	} else {
		goto L770
	}
L749:
	;
	v3955 = *(*int32)(unsafe.Add(mBase, uint32(v3901)+4))
	if v3955 <= int32(0) {
		v4046 = v3923
		goto L748
	} else {
		goto L750
	}
L750:
	;
	v3958 = int32(0)
	if v3958 < v3955 {
		goto L751
	} else {
		goto L752
	}
L751:
	;
	v3961 = v3955
	goto L753
L752:
	;
	v3961 = v3958
	goto L753
L753:
	;
	v3962 = *(*int32)(unsafe.Add(mBase, uint32(v3901)+12))
	v3968 = int32(0)
	goto L754
L754:
	;
	v3998 = *(*int32)(unsafe.Add(mBase, uint32(v3962+v3968<<(uint(int32(2))%32))))
	v3999 = *(*int32)(unsafe.Add(mBase, uint32(v3998)+4))
	v4002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3948))))
	v4005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3999))))
	if base.B2i32(v4002 == int32(0))|base.B2i32(v4002 != v4005) != 0 {
		v4023 = v4002
		v4024 = v4005
		goto L757
	} else {
		goto L758
	}
L755:
	;
	v4029 = F_makeString(m, v3948)
	mBase = m.M
	v4030 = m.ExcPending
	if v4030 != 0 {
		goto L1
	} else {
		goto L767
	}
L756:
	;
	if v4023-v4024 != 0 {
		goto L763
	} else {
		goto L764
	}
L757:
	;
	goto L756
L758:
	;
	v4008 = v3948
	v4009 = v3999
	goto L759
L759:
	;
	v4012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4009)+1)))
	v4013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4008)+1)))
	if v4013 == int32(0) {
		v4023 = v4013
		v4024 = v4012
		goto L757
	} else {
		goto L761
	}
L760:
	;
	v4023 = v4013
	v4024 = v4012
	goto L757
L761:
	;
	v4016 = int32(1)
	if v4013 == v4012 {
		v4008 = v4008 + v4016
		v4009 = v4009 + v4016
		goto L759
	} else {
		goto L762
	}
L762:
	;
	goto L760
L763:
	;
	v4027 = v3968 + int32(1)
	if v3961 != v4027 {
		v3968 = v4027
		goto L754
	} else {
		goto L766
	}
L764:
	;
	goto L765
L765:
	;
	goto L755
L766:
	;
	v4046 = v3923
	goto L748
L767:
	;
	if v4029 == int32(0) {
		v4046 = v3923
		goto L748
	} else {
		goto L768
	}
L768:
	;
	v4033 = F_lappend(m, v3923, v4029)
	mBase = m.M
	v4034 = m.ExcPending
	if v4034 != 0 {
		goto L1
	} else {
		goto L769
	}
L769:
	;
	v4046 = v4033
	goto L748
L770:
	;
	goto L747
L771:
	;
	v4134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4133)+8)) = v4134
	goto L773
L772:
	;
	goto L773
L773:
	;
	v4136 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+268)) = v4136
	*(*int32)(unsafe.Add(mBase, uint32(v34)+272)) = v4136
	*(*int32)(unsafe.Add(mBase, uint32(v34)+276)) = v4136
	*(*int32)(unsafe.Add(mBase, uint32(v34)+264)) = v4136
	if v3903 != 0 {
		goto L774
	} else {
		goto L775
	}
L774:
	;
	v4146 = *(*int32)(unsafe.Add(mBase, uint32(v3903)+4))
	v4147 = v4146
	goto L776
L775:
	;
	v4147 = v4136
	goto L776
L776:
	;
	if v3901 != 0 {
		goto L777
	} else {
		goto L778
	}
L777:
	;
	v4148 = *(*int32)(unsafe.Add(mBase, uint32(v3901)+4))
	v4150 = v4148
	goto L779
L778:
	;
	v4150 = int32(0)
	goto L779
L779:
	;
	v4154 = F_palloc0(m, (v4150+v4147)<<(uint(int32(5))%32))
	mBase = m.M
	v4155 = m.ExcPending
	if v4155 != 0 {
		goto L1
	} else {
		goto L780
	}
L780:
	;
	v4156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v4156 != 0 {
		goto L787
	} else {
		goto L788
	}
L781:
	;
	v5038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v5038 != 0 {
		goto L927
	} else {
		goto L928
	}
L782:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+272)) = v4899
	*(*int32)(unsafe.Add(mBase, uint32(v34)+276)) = v4895
	*(*int32)(unsafe.Add(mBase, uint32(v34)+268)) = v4881
	v4906 = int32(0)
	v4912 = v4906
	v4913 = v4906
	goto L903
L783:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4857 = m.ExcPending
	if v4857 != 0 {
		goto L1
	} else {
		goto L898
	}
L784:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4808 = m.ExcPending
	if v4808 != 0 {
		goto L1
	} else {
		goto L894
	}
L785:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4790 = m.ExcPending
	if v4790 != 0 {
		goto L1
	} else {
		goto L890
	}
L786:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4741 = m.ExcPending
	if v4741 != 0 {
		goto L1
	} else {
		goto L886
	}
L787:
	;
	v4157 = *(*int32)(unsafe.Add(mBase, uint32(v4156)+4))
	if v4157 <= int32(0) {
		v4881 = v5
		v4889 = v5
		v4895 = v5
		v4898 = v5
		v4899 = v5
		goto L782
	} else {
		goto L790
	}
L788:
	;
	goto L789
L789:
	;
	v4615 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v4615 != 0 {
		goto L872
	} else {
		goto L873
	}
L790:
	;
	v4169 = v5
	v4177 = v5
	v4183 = v5
	v4186 = v5
	v4187 = v5
	v4188 = v5
	goto L791
L791:
	;
	v4191 = *(*int32)(unsafe.Add(mBase, uint32(v4156)+12))
	v4194 = v4191 + v4188<<(uint(int32(2))%32)
	v4195 = *(*int32)(unsafe.Add(mBase, uint32(v4194)))
	v4196 = *(*int32)(unsafe.Add(mBase, uint32(v4195)+4))
	if v4183 == int32(0) {
		goto L793
	} else {
		goto L794
	}
L792:
	;
	v4881 = v4564
	v4889 = v4585
	v4895 = v4609
	v4898 = v4606
	v4899 = v4442
	goto L782
L793:
	;
	if v3903 == int32(0) {
		goto L785
	} else {
		goto L816
	}
L794:
	;
	v4199 = *(*int32)(unsafe.Add(mBase, uint32(v4183)+4))
	if v4199 <= int32(0) {
		goto L793
	} else {
		goto L795
	}
L795:
	;
	v4202 = int32(0)
	if v4202 < v4199 {
		goto L796
	} else {
		goto L797
	}
L796:
	;
	v4205 = v4199
	goto L798
L797:
	;
	v4205 = v4202
	goto L798
L798:
	;
	v4206 = *(*int32)(unsafe.Add(mBase, uint32(v4183)+12))
	v4212 = int32(0)
	goto L799
L799:
	;
	v4242 = *(*int32)(unsafe.Add(mBase, uint32(v4206+v4212<<(uint(int32(2))%32))))
	v4243 = *(*int32)(unsafe.Add(mBase, uint32(v4242)+4))
	v4246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4243))))
	v4249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4196))))
	if base.B2i32(v4246 == int32(0))|base.B2i32(v4246 != v4249) != 0 {
		v4267 = v4246
		v4268 = v4249
		goto L802
	} else {
		goto L803
	}
L800:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4276 = m.ExcPending
	if v4276 != 0 {
		goto L1
	} else {
		goto L812
	}
L801:
	;
	if v4267-v4268 != 0 {
		goto L808
	} else {
		goto L809
	}
L802:
	;
	goto L801
L803:
	;
	v4252 = v4243
	v4253 = v4196
	goto L804
L804:
	;
	v4256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4253)+1)))
	v4257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4252)+1)))
	if v4257 == int32(0) {
		v4267 = v4257
		v4268 = v4256
		goto L802
	} else {
		goto L806
	}
L805:
	;
	v4267 = v4257
	v4268 = v4256
	goto L802
L806:
	;
	v4260 = int32(1)
	if v4257 == v4256 {
		v4252 = v4252 + v4260
		v4253 = v4253 + v4260
		goto L804
	} else {
		goto L807
	}
L807:
	;
	goto L805
L808:
	;
	v4271 = v4212 + int32(1)
	if v4205 != v4271 {
		v4212 = v4271
		goto L799
	} else {
		goto L811
	}
L809:
	;
	goto L810
L810:
	;
	goto L800
L811:
	;
	goto L793
L812:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v4279 = m.ExcPending
	if v4279 != 0 {
		goto L1
	} else {
		goto L813
	}
L813:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+144)) = v4196
	F_errmsg(m, int32(_a_F_transformFromClauseItem_48), v34+int32(144))
	mBase = m.M
	v4285 = m.ExcPending
	if v4285 != 0 {
		goto L1
	} else {
		goto L814
	}
L814:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1330), int32(_a_F_transformFromClauseItem_40))
	mBase = m.M
	v4290 = m.ExcPending
	if v4290 != 0 {
		goto L1
	} else {
		goto L815
	}
L815:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L816:
	;
	v4324 = *(*int32)(unsafe.Add(mBase, uint32(v3903)+4))
	if v4324 <= int32(0) {
		goto L818
	} else {
		goto L819
	}
L817:
	;
	if v4413 < int32(0) {
		goto L785
	} else {
		goto L838
	}
L818:
	;
	v4413 = int32(-1)
	goto L817
L819:
	;
	goto L820
L820:
	;
	v4328 = int32(0)
	if v4328 < v4324 {
		goto L821
	} else {
		goto L822
	}
L821:
	;
	v4331 = v4324
	goto L823
L822:
	;
	v4331 = v4328
	goto L823
L823:
	;
	v4332 = *(*int32)(unsafe.Add(mBase, uint32(v3903)+12))
	v4339 = int32(0)
	v4341 = int32(-1)
	goto L824
L824:
	;
	v4369 = *(*int32)(unsafe.Add(mBase, uint32(v4332+v4339<<(uint(int32(2))%32))))
	v4370 = *(*int32)(unsafe.Add(mBase, uint32(v4369)+4))
	v4373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4370))))
	v4376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4196))))
	if base.B2i32(v4373 == int32(0))|base.B2i32(v4373 != v4376) != 0 {
		v4394 = v4373
		v4395 = v4376
		goto L827
	} else {
		goto L828
	}
L825:
	;
	v4413 = v4403
	goto L817
L826:
	;
	if v4394-v4395 == int32(0) {
		goto L833
	} else {
		goto L834
	}
L827:
	;
	goto L826
L828:
	;
	v4379 = v4370
	v4380 = v4196
	goto L829
L829:
	;
	v4383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4380)+1)))
	v4384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4379)+1)))
	if v4384 == int32(0) {
		v4394 = v4384
		v4395 = v4383
		goto L827
	} else {
		goto L831
	}
L830:
	;
	v4394 = v4384
	v4395 = v4383
	goto L827
L831:
	;
	v4387 = int32(1)
	if v4384 == v4383 {
		v4379 = v4379 + v4387
		v4380 = v4380 + v4387
		goto L829
	} else {
		goto L832
	}
L832:
	;
	goto L830
L833:
	;
	v4399 = int32(0)
	if base.B2i32(v4341 < v4399) == v4399 {
		goto L786
	} else {
		goto L836
	}
L834:
	;
	v4403 = v4341
	goto L835
L835:
	;
	v4405 = v4339 + int32(1)
	if v4405 != v4331 {
		v4339 = v4405
		v4341 = v4403
		goto L824
	} else {
		goto L837
	}
L836:
	;
	v4403 = v4339
	goto L835
L837:
	;
	goto L825
L838:
	;
	v4442 = F_lappend_int(m, v4187, v4413+int32(1))
	mBase = m.M
	v4443 = m.ExcPending
	if v4443 != 0 {
		goto L1
	} else {
		goto L839
	}
L839:
	;
	if v3901 == int32(0) {
		goto L783
	} else {
		goto L840
	}
L840:
	;
	v4446 = *(*int32)(unsafe.Add(mBase, uint32(v3901)+4))
	if v4446 <= int32(0) {
		goto L842
	} else {
		goto L843
	}
L841:
	;
	if v4540 < int32(0) {
		goto L783
	} else {
		goto L862
	}
L842:
	;
	v4540 = int32(-1)
	goto L841
L843:
	;
	goto L844
L844:
	;
	v4450 = int32(0)
	if v4450 < v4446 {
		goto L845
	} else {
		goto L846
	}
L845:
	;
	v4453 = v4446
	goto L847
L846:
	;
	v4453 = v4450
	goto L847
L847:
	;
	v4454 = *(*int32)(unsafe.Add(mBase, uint32(v3901)+12))
	v4461 = int32(0)
	v4468 = int32(-1)
	goto L848
L848:
	;
	v4491 = *(*int32)(unsafe.Add(mBase, uint32(v4454+v4461<<(uint(int32(2))%32))))
	v4492 = *(*int32)(unsafe.Add(mBase, uint32(v4491)+4))
	v4495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4492))))
	v4498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4196))))
	if base.B2i32(v4495 == int32(0))|base.B2i32(v4495 != v4498) != 0 {
		v4516 = v4495
		v4517 = v4498
		goto L851
	} else {
		goto L852
	}
L849:
	;
	v4540 = v4525
	goto L841
L850:
	;
	if v4516-v4517 == int32(0) {
		goto L857
	} else {
		goto L858
	}
L851:
	;
	goto L850
L852:
	;
	v4501 = v4492
	v4502 = v4196
	goto L853
L853:
	;
	v4505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4502)+1)))
	v4506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4501)+1)))
	if v4506 == int32(0) {
		v4516 = v4506
		v4517 = v4505
		goto L851
	} else {
		goto L855
	}
L854:
	;
	v4516 = v4506
	v4517 = v4505
	goto L851
L855:
	;
	v4509 = int32(1)
	if v4506 == v4505 {
		v4501 = v4501 + v4509
		v4502 = v4502 + v4509
		goto L853
	} else {
		goto L856
	}
L856:
	;
	goto L854
L857:
	;
	v4521 = int32(0)
	if base.B2i32(v4468 < v4521) == v4521 {
		goto L784
	} else {
		goto L860
	}
L858:
	;
	v4525 = v4468
	goto L859
L859:
	;
	v4527 = v4461 + int32(1)
	if v4527 != v4453 {
		v4461 = v4527
		v4468 = v4525
		goto L848
	} else {
		goto L861
	}
L860:
	;
	v4525 = v4461
	goto L859
L861:
	;
	goto L849
L862:
	;
	v4564 = F_lappend_int(m, v4169, v4540+int32(1))
	mBase = m.M
	v4565 = m.ExcPending
	if v4565 != 0 {
		goto L1
	} else {
		goto L863
	}
L863:
	;
	v4568 = v3899 + v4413<<(uint(int32(5))%32)
	v4569 = *(*int32)(unsafe.Add(mBase, uint32(v4568)))
	v4570 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4568)+4)))
	v4571 = *(*int32)(unsafe.Add(mBase, uint32(v4568)+8))
	v4572 = *(*int32)(unsafe.Add(mBase, uint32(v4568)+12))
	v4573 = *(*int32)(unsafe.Add(mBase, uint32(v4568)+16))
	v4575 = F_makeVar(m, v4569, v4570, v4571, v4572, v4573, int32(0))
	mBase = m.M
	v4576 = m.ExcPending
	if v4576 != 0 {
		goto L1
	} else {
		goto L864
	}
L864:
	;
	v4577 = *(*int32)(unsafe.Add(mBase, uint32(v4568)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4575)+32)) = v4577
	v4579 = *(*int32)(unsafe.Add(mBase, uint32(v4568)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v4575)+36)) = v4579
	v4581 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4568)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4575)+40)) = uint16(v4581)
	F_markNullableIfNeeded(m, l0, v4575)
	mBase = m.M
	v4584 = m.ExcPending
	if v4584 != 0 {
		goto L1
	} else {
		goto L865
	}
L865:
	;
	v4585 = F_lappend(m, v4177, v4575)
	mBase = m.M
	v4586 = m.ExcPending
	if v4586 != 0 {
		goto L1
	} else {
		goto L866
	}
L866:
	;
	v4589 = v3897 + v4540<<(uint(int32(5))%32)
	v4590 = *(*int32)(unsafe.Add(mBase, uint32(v4589)))
	v4591 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4589)+4)))
	v4592 = *(*int32)(unsafe.Add(mBase, uint32(v4589)+8))
	v4593 = *(*int32)(unsafe.Add(mBase, uint32(v4589)+12))
	v4594 = *(*int32)(unsafe.Add(mBase, uint32(v4589)+16))
	v4596 = F_makeVar(m, v4590, v4591, v4592, v4593, v4594, int32(0))
	mBase = m.M
	v4597 = m.ExcPending
	if v4597 != 0 {
		goto L1
	} else {
		goto L867
	}
L867:
	;
	v4598 = *(*int32)(unsafe.Add(mBase, uint32(v4589)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4596)+32)) = v4598
	v4600 = *(*int32)(unsafe.Add(mBase, uint32(v4589)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v4596)+36)) = v4600
	v4602 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4589)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4596)+40)) = uint16(v4602)
	F_markNullableIfNeeded(m, l0, v4596)
	mBase = m.M
	v4605 = m.ExcPending
	if v4605 != 0 {
		goto L1
	} else {
		goto L868
	}
L868:
	;
	v4606 = F_lappend(m, v4186, v4596)
	mBase = m.M
	v4607 = m.ExcPending
	if v4607 != 0 {
		goto L1
	} else {
		goto L869
	}
L869:
	;
	v4608 = *(*int32)(unsafe.Add(mBase, uint32(v4194)))
	v4609 = F_lappend(m, v4183, v4608)
	mBase = m.M
	v4610 = m.ExcPending
	if v4610 != 0 {
		goto L1
	} else {
		goto L870
	}
L870:
	;
	v4612 = v4188 + int32(1)
	v4613 = *(*int32)(unsafe.Add(mBase, uint32(v4156)+4))
	if v4612 < v4613 {
		v4169 = v4564
		v4177 = v4585
		v4183 = v4609
		v4186 = v4606
		v4187 = v4442
		v4188 = v4612
		goto L791
	} else {
		goto L871
	}
L871:
	;
	goto L792
L872:
	;
	if v3894 == int32(0) {
		goto L875
	} else {
		goto L876
	}
L873:
	;
	goto L874
L874:
	;
	v5016 = v5
	v5034 = v5
	goto L781
L875:
	;
	v4694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v3894
	v4697 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v4697 != 0 {
		goto L881
	} else {
		goto L882
	}
L876:
	;
	v4618 = *(*int32)(unsafe.Add(mBase, uint32(v3894)+4))
	if v4618 <= int32(0) {
		goto L875
	} else {
		goto L877
	}
L877:
	;
	v4625 = v4136
	goto L878
L878:
	;
	v4652 = *(*int32)(unsafe.Add(mBase, uint32(v3894)+12))
	v4656 = *(*int32)(unsafe.Add(mBase, uint32(v4652+v4625<<(uint(int32(2))%32))))
	v4657 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v4656)+22)) = uint16(v4657)
	v4660 = v4625 + int32(1)
	v4661 = *(*int32)(unsafe.Add(mBase, uint32(v3894)+4))
	if v4660 < v4661 {
		v4625 = v4660
		goto L878
	} else {
		goto L880
	}
L879:
	;
	goto L875
L880:
	;
	goto L879
L881:
	;
	v4699 = F_transformExpr(m, l0, v4697, int32(2))
	mBase = m.M
	v4700 = m.ExcPending
	if v4700 != 0 {
		goto L1
	} else {
		goto L884
	}
L882:
	;
	v4704 = int32(0)
	goto L883
L883:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v4694
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v4704
	goto L874
L884:
	;
	v4702 = F_coerce_to_boolean(m, l0, v4699, int32(_a_F_transformFromClauseItem_49))
	mBase = m.M
	v4703 = m.ExcPending
	if v4703 != 0 {
		goto L1
	} else {
		goto L885
	}
L885:
	;
	v4704 = v4702
	goto L883
L886:
	;
	F_errcode(m, int32(33583236))
	mBase = m.M
	v4744 = m.ExcPending
	if v4744 != 0 {
		goto L1
	} else {
		goto L887
	}
L887:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+128)) = v4196
	F_errmsg(m, int32(_a_F_transformFromClauseItem_50), v34+int32(128))
	mBase = m.M
	v4750 = m.ExcPending
	if v4750 != 0 {
		goto L1
	} else {
		goto L888
	}
L888:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1345), int32(_a_F_transformFromClauseItem_40))
	mBase = m.M
	v4755 = m.ExcPending
	if v4755 != 0 {
		goto L1
	} else {
		goto L889
	}
L889:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L890:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v4793 = m.ExcPending
	if v4793 != 0 {
		goto L1
	} else {
		goto L891
	}
L891:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+80)) = v4196
	F_errmsg(m, int32(_a_F_transformFromClauseItem_51), v34+int32(80))
	mBase = m.M
	v4799 = m.ExcPending
	if v4799 != 0 {
		goto L1
	} else {
		goto L892
	}
L892:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1354), int32(_a_F_transformFromClauseItem_40))
	mBase = m.M
	v4804 = m.ExcPending
	if v4804 != 0 {
		goto L1
	} else {
		goto L893
	}
L893:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L894:
	;
	F_errcode(m, int32(33583236))
	mBase = m.M
	v4811 = m.ExcPending
	if v4811 != 0 {
		goto L1
	} else {
		goto L895
	}
L895:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+112)) = v4196
	F_errmsg(m, int32(_a_F_transformFromClauseItem_52), v34+int32(112))
	mBase = m.M
	v4817 = m.ExcPending
	if v4817 != 0 {
		goto L1
	} else {
		goto L896
	}
L896:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1369), int32(_a_F_transformFromClauseItem_40))
	mBase = m.M
	v4822 = m.ExcPending
	if v4822 != 0 {
		goto L1
	} else {
		goto L897
	}
L897:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L898:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v4860 = m.ExcPending
	if v4860 != 0 {
		goto L1
	} else {
		goto L899
	}
L899:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+96)) = v4196
	F_errmsg(m, int32(_a_F_transformFromClauseItem_53), v34+int32(96))
	mBase = m.M
	v4866 = m.ExcPending
	if v4866 != 0 {
		goto L1
	} else {
		goto L900
	}
L900:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1378), int32(_a_F_transformFromClauseItem_40))
	mBase = m.M
	v4871 = m.ExcPending
	if v4871 != 0 {
		goto L1
	} else {
		goto L901
	}
L901:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L902:
	;
	v5001 = F_transformExpr(m, l0, v4999, int32(3))
	mBase = m.M
	v5002 = m.ExcPending
	if v5002 != 0 {
		goto L1
	} else {
		goto L925
	}
L903:
	;
	v4939 = int32(0)
	if v4889 == v4939 {
		v4950 = v4939
		goto L905
	} else {
		goto L906
	}
L904:
	;
	v4995 = *(*int32)(unsafe.Add(mBase, uint32(v4913)+12))
	v4996 = *(*int32)(unsafe.Add(mBase, uint32(v4995)))
	v4999 = v4996
	goto L902
L905:
	;
	if v4898 == int32(0) {
		v4967 = v4939
		goto L910
	} else {
		goto L911
	}
L906:
	;
	v4944 = *(*int32)(unsafe.Add(mBase, uint32(v4889)+4))
	if v4944 <= v4912 {
		v4950 = int32(0)
		goto L905
	} else {
		goto L907
	}
L907:
	;
	v4946 = *(*int32)(unsafe.Add(mBase, uint32(v4889)+12))
	v4950 = v4946 + v4912<<(uint(int32(2))%32)
	goto L905
L908:
	;
	goto L904
L909:
	;
	v4976 = *(*int32)(unsafe.Add(mBase, uint32(v4960+v4912<<(uint(int32(2))%32))))
	v4977 = *(*int32)(unsafe.Add(mBase, uint32(v4950)))
	F_markVarForSelectPriv(m, l0, v4977)
	mBase = m.M
	v4979 = m.ExcPending
	if v4979 != 0 {
		goto L1
	} else {
		goto L919
	}
L910:
	;
	v4971 = F_makeBoolExpr(m, int32(0), v4967, int32(-1))
	mBase = m.M
	v4972 = m.ExcPending
	if v4972 != 0 {
		goto L1
	} else {
		goto L918
	}
L911:
	;
	v4953 = int32(0)
	v4955 = *(*int32)(unsafe.Add(mBase, uint32(v4898)+4))
	if base.B2i32(v4950 == v4953)|base.B2i32(v4955 <= v4912) == v4953 {
		goto L912
	} else {
		goto L913
	}
L912:
	;
	v4960 = *(*int32)(unsafe.Add(mBase, uint32(v4898)+12))
	if v4960 != 0 {
		goto L909
	} else {
		goto L915
	}
L913:
	;
	goto L914
L914:
	;
	if v4913 == int32(0) {
		v4967 = v4939
		goto L910
	} else {
		goto L916
	}
L915:
	;
	goto L914
L916:
	;
	v4964 = *(*int32)(unsafe.Add(mBase, uint32(v4913)+4))
	if v4964 == int32(1) {
		goto L908
	} else {
		goto L917
	}
L917:
	;
	v4967 = v4913
	goto L910
L918:
	;
	v4999 = v4971
	goto L902
L919:
	;
	F_markVarForSelectPriv(m, l0, v4976)
	mBase = m.M
	v4981 = m.ExcPending
	if v4981 != 0 {
		goto L1
	} else {
		goto L920
	}
L920:
	;
	v4986 = F_copyObjectImpl(m, v4977)
	mBase = m.M
	v4987 = m.ExcPending
	if v4987 != 0 {
		goto L1
	} else {
		goto L921
	}
L921:
	;
	v4988 = F_copyObjectImpl(m, v4976)
	mBase = m.M
	v4989 = m.ExcPending
	if v4989 != 0 {
		goto L1
	} else {
		goto L922
	}
L922:
	;
	v4991 = F_makeSimpleA_Expr(m, int32(0), int32(_a_F_transformFromClauseItem_54), v4986, v4988, int32(-1))
	mBase = m.M
	v4992 = m.ExcPending
	if v4992 != 0 {
		goto L1
	} else {
		goto L923
	}
L923:
	;
	v4993 = F_lappend(m, v4913, v4991)
	mBase = m.M
	v4994 = m.ExcPending
	if v4994 != 0 {
		goto L1
	} else {
		goto L924
	}
L924:
	;
	v4912 = v4912 + int32(1)
	v4913 = v4993
	goto L903
L925:
	;
	v5004 = F_coerce_to_boolean(m, l0, v5001, int32(_a_F_transformFromClauseItem_55))
	mBase = m.M
	v5005 = m.ExcPending
	if v5005 != 0 {
		goto L1
	} else {
		goto L926
	}
L926:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v5004
	v5016 = v4881
	v5034 = v4899
	goto L781
L927:
	;
	v5039 = *(*int32)(unsafe.Add(mBase, uint32(v5038)+4))
	v5043 = v5039 + int32(1)
	goto L929
L928:
	;
	v5043 = int32(1)
	goto L929
L929:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v5043
	v5045 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v5045 {
	case 0:
		goto L930
	case 1:
		goto L931
	case 2:
		goto L934
	case 3:
		goto L933
	default:
		goto L932
	}
L930:
	;
	v5075 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v5075 == int32(0) {
		goto L943
	} else {
		goto L944
	}
L931:
	;
	v5072 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_markRelsAsNulledBy(m, l0, v5072, v5043)
	mBase = m.M
	v5074 = m.ExcPending
	if v5074 != 0 {
		goto L1
	} else {
		goto L941
	}
L932:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5059 = m.ExcPending
	if v5059 != 0 {
		goto L1
	} else {
		goto L938
	}
L933:
	;
	v5053 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_markRelsAsNulledBy(m, l0, v5053, v5043)
	mBase = m.M
	v5055 = m.ExcPending
	if v5055 != 0 {
		goto L1
	} else {
		goto L937
	}
L934:
	;
	v5046 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_markRelsAsNulledBy(m, l0, v5046, v5043)
	mBase = m.M
	v5048 = m.ExcPending
	if v5048 != 0 {
		goto L1
	} else {
		goto L935
	}
L935:
	;
	v5049 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v5050 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	F_markRelsAsNulledBy(m, l0, v5049, v5050)
	mBase = m.M
	v5052 = m.ExcPending
	if v5052 != 0 {
		goto L1
	} else {
		goto L936
	}
L936:
	;
	goto L930
L937:
	;
	goto L930
L938:
	;
	v5060 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v5060
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_56), v34+int32(32))
	mBase = m.M
	v5066 = m.ExcPending
	if v5066 != 0 {
		goto L1
	} else {
		goto L939
	}
L939:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1447), int32(_a_F_transformFromClauseItem_40))
	mBase = m.M
	v5071 = m.ExcPending
	if v5071 != 0 {
		goto L1
	} else {
		goto L940
	}
L940:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L941:
	;
	goto L930
L942:
	;
	v5389 = v34 + int32(276)
	v5391 = v34 + int32(264)
	v5397 = F_extractRemainingColumns(m, l0, v3899, v3903, v34+int32(272), v5389, v5391, v4154+v5360<<(uint(int32(5))%32))
	mBase = m.M
	v5398 = m.ExcPending
	if v5398 != 0 {
		goto L1
	} else {
		goto L1004
	}
L943:
	;
	v5360 = int32(0)
	goto L942
L944:
	;
	goto L945
L945:
	;
	v5079 = int32(0)
	v5086 = v5079
	v5107 = v5079
	goto L946
L946:
	;
	v5112 = int32(0)
	if v5034 == v5112 {
		v5122 = v5112
		goto L948
	} else {
		goto L949
	}
L947:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+264)) = v5353
	v5360 = v5351
	goto L942
L948:
	;
	if v5016 == int32(0) {
		goto L952
	} else {
		goto L953
	}
L949:
	;
	v5116 = *(*int32)(unsafe.Add(mBase, uint32(v5034)+4))
	if v5116 <= v5086 {
		v5122 = int32(0)
		goto L948
	} else {
		goto L950
	}
L950:
	;
	v5118 = *(*int32)(unsafe.Add(mBase, uint32(v5034)+12))
	v5122 = v5118 + v5086<<(uint(int32(2))%32)
	goto L948
L951:
	;
	goto L947
L952:
	;
	v5125 = int32(0)
	v5351 = v5125
	v5353 = v5125
	goto L951
L953:
	;
	goto L954
L954:
	;
	v5129 = *(*int32)(unsafe.Add(mBase, uint32(v5016)+4))
	if base.B2i32(v5122 == int32(0))|base.B2i32(v5129 <= v5086) != 0 {
		v5351 = v5086
		v5353 = v5107
		goto L951
	} else {
		goto L955
	}
L955:
	;
	v5132 = *(*int32)(unsafe.Add(mBase, uint32(v5016)+12))
	if v5132 == int32(0) {
		v5351 = v5086
		v5353 = v5107
		goto L951
	} else {
		goto L956
	}
L956:
	;
	v5138 = *(*int32)(unsafe.Add(mBase, uint32(v5132+v5086<<(uint(int32(2))%32))))
	v5139 = *(*int32)(unsafe.Add(mBase, uint32(v5122)))
	v5142 = v3899 + v5139<<(uint(int32(5))%32)
	v5144 = v5142 - int32(32)
	v5145 = *(*int32)(unsafe.Add(mBase, uint32(v5144)))
	v5148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5142-int32(28)))))
	v5151 = *(*int32)(unsafe.Add(mBase, uint32(v5142-int32(24))))
	v5154 = *(*int32)(unsafe.Add(mBase, uint32(v5142-int32(20))))
	v5157 = *(*int32)(unsafe.Add(mBase, uint32(v5142-int32(16))))
	v5159 = F_makeVar(m, v5145, v5148, v5151, v5154, v5157, int32(0))
	mBase = m.M
	v5160 = m.ExcPending
	if v5160 != 0 {
		goto L1
	} else {
		goto L957
	}
L957:
	;
	v5163 = *(*int32)(unsafe.Add(mBase, uint32(v5142-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v5159)+32)) = v5163
	v5167 = *(*int32)(unsafe.Add(mBase, uint32(v5142-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v5159)+36)) = v5167
	v5171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5142-int32(4)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5159)+40)) = uint16(v5171)
	F_markNullableIfNeeded(m, l0, v5159)
	mBase = m.M
	v5174 = m.ExcPending
	if v5174 != 0 {
		goto L1
	} else {
		goto L958
	}
L958:
	;
	v5177 = v3897 + v5138<<(uint(int32(5))%32)
	v5179 = v5177 - int32(32)
	v5180 = *(*int32)(unsafe.Add(mBase, uint32(v5179)))
	v5183 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5177-int32(28)))))
	v5186 = *(*int32)(unsafe.Add(mBase, uint32(v5177-int32(24))))
	v5189 = *(*int32)(unsafe.Add(mBase, uint32(v5177-int32(20))))
	v5192 = *(*int32)(unsafe.Add(mBase, uint32(v5177-int32(16))))
	v5194 = F_makeVar(m, v5180, v5183, v5186, v5189, v5192, int32(0))
	mBase = m.M
	v5195 = m.ExcPending
	if v5195 != 0 {
		goto L1
	} else {
		goto L959
	}
L959:
	;
	v5198 = *(*int32)(unsafe.Add(mBase, uint32(v5177-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v5194)+32)) = v5198
	v5202 = *(*int32)(unsafe.Add(mBase, uint32(v5177-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v5194)+36)) = v5202
	v5206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5177-int32(4)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5194)+40)) = uint16(v5206)
	F_markNullableIfNeeded(m, l0, v5194)
	mBase = m.M
	v5209 = m.ExcPending
	if v5209 != 0 {
		goto L1
	} else {
		goto L960
	}
L960:
	;
	v5210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+76)) = v5159
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v5194
	*(*int32)(unsafe.Add(mBase, uint32(v34)+332)) = v5159
	*(*int32)(unsafe.Add(mBase, uint32(v34)+328)) = v5194
	v5219 = F_list_make2_impl(m, v34+int32(76), v34+int32(72))
	mBase = m.M
	v5220 = m.ExcPending
	if v5220 != 0 {
		goto L1
	} else {
		goto L961
	}
L961:
	;
	v5223 = F_select_common_type(m, l0, v5219, int32(_a_F_transformFromClauseItem_55), int32(0))
	mBase = m.M
	v5224 = m.ExcPending
	if v5224 != 0 {
		goto L1
	} else {
		goto L962
	}
L962:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+320)) = v5194
	*(*int32)(unsafe.Add(mBase, uint32(v34)+324)) = v5159
	*(*int32)(unsafe.Add(mBase, uint32(v34)+68)) = v5159
	*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v5194
	v5233 = F_list_make2_impl(m, v34+int32(68), v34-int32(-64))
	mBase = m.M
	v5234 = m.ExcPending
	if v5234 != 0 {
		goto L1
	} else {
		goto L963
	}
L963:
	;
	v5235 = F_select_common_typmod(m, v5233, v5223)
	mBase = m.M
	v5236 = m.ExcPending
	if v5236 != 0 {
		goto L1
	} else {
		goto L964
	}
L964:
	;
	v5237 = *(*int32)(unsafe.Add(mBase, uint32(v5159)+12))
	if v5237 != v5223 {
		goto L966
	} else {
		goto L967
	}
L965:
	;
	v5251 = *(*int32)(unsafe.Add(mBase, uint32(v5194)+12))
	if v5223 != v5251 {
		goto L973
	} else {
		goto L974
	}
L966:
	;
	v5242 = F_coerce_type(m, l0, v5159, v5237, v5223, v5235, int32(0), int32(2), int32(-1))
	mBase = m.M
	v5243 = m.ExcPending
	if v5243 != 0 {
		goto L1
	} else {
		goto L969
	}
L967:
	;
	goto L968
L968:
	;
	v5244 = *(*int32)(unsafe.Add(mBase, uint32(v5159)+16))
	if v5244 == v5235 {
		v5250 = v5159
		goto L965
	} else {
		goto L970
	}
L969:
	;
	v5250 = v5242
	goto L965
L970:
	;
	v5248 = F_makeRelabelType(m, v5159, v5223, v5235, int32(0), int32(2))
	mBase = m.M
	v5249 = m.ExcPending
	if v5249 != 0 {
		goto L1
	} else {
		goto L971
	}
L971:
	;
	v5250 = v5248
	goto L965
L972:
	;
	switch v5210 {
	case 0:
		goto L983
	case 1:
		v5308 = v5250
		goto L979
	case 2:
		goto L982
	case 3:
		goto L980
	default:
		goto L981
	}
L973:
	;
	v5256 = F_coerce_type(m, l0, v5194, v5251, v5223, v5235, int32(0), int32(2), int32(-1))
	mBase = m.M
	v5257 = m.ExcPending
	if v5257 != 0 {
		goto L1
	} else {
		goto L976
	}
L974:
	;
	goto L975
L975:
	;
	v5258 = *(*int32)(unsafe.Add(mBase, uint32(v5194)+16))
	if v5258 == v5235 {
		v5264 = v5194
		goto L972
	} else {
		goto L977
	}
L976:
	;
	v5264 = v5256
	goto L972
L977:
	;
	v5262 = F_makeRelabelType(m, v5194, v5223, v5235, int32(0), int32(2))
	mBase = m.M
	v5263 = m.ExcPending
	if v5263 != 0 {
		goto L1
	} else {
		goto L978
	}
L978:
	;
	v5264 = v5262
	goto L972
L979:
	;
	F_assign_expr_collations(m, l0, v5308)
	mBase = m.M
	v5310 = m.ExcPending
	if v5310 != 0 {
		goto L1
	} else {
		goto L993
	}
L980:
	;
	v5308 = v5264
	goto L979
L981:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5294 = m.ExcPending
	if v5294 != 0 {
		goto L1
	} else {
		goto L990
	}
L982:
	;
	v5273 = F_palloc0(m, int32(20))
	mBase = m.M
	v5274 = m.ExcPending
	if v5274 != 0 {
		goto L1
	} else {
		goto L988
	}
L983:
	;
	v5265 = *(*int32)(unsafe.Add(mBase, uint32(v5250)))
	if v5265 == int32(6) {
		v5308 = v5250
		goto L979
	} else {
		goto L984
	}
L984:
	;
	v5268 = *(*int32)(unsafe.Add(mBase, uint32(v5264)))
	if v5268 == int32(6) {
		goto L985
	} else {
		goto L986
	}
L985:
	;
	v5271 = v5264
	goto L987
L986:
	;
	v5271 = v5250
	goto L987
L987:
	;
	v5308 = v5271
	goto L979
L988:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5273)+4)) = v5223
	*(*int32)(unsafe.Add(mBase, uint32(v5273))) = int32(38)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+312)) = v5264
	*(*int32)(unsafe.Add(mBase, uint32(v34)+316)) = v5250
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v5250
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v5264
	v5286 = F_list_make2_impl(m, v34+int32(60), v34+int32(56))
	mBase = m.M
	v5287 = m.ExcPending
	if v5287 != 0 {
		goto L1
	} else {
		goto L989
	}
L989:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5273)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v5273)+12)) = v5286
	v5308 = v5273
	goto L979
L990:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v5210
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_56), v34+int32(48))
	mBase = m.M
	v5300 = m.ExcPending
	if v5300 != 0 {
		goto L1
	} else {
		goto L991
	}
L991:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1754), int32(_a_F_transformFromClauseItem_57))
	mBase = m.M
	v5305 = m.ExcPending
	if v5305 != 0 {
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
	v5312 = v5086 + int32(1)
	v5315 = v4154 + v5086<<(uint(int32(5))%32)
	v5316 = F_lappend(m, v5107, v5308)
	mBase = m.M
	v5317 = m.ExcPending
	if v5317 != 0 {
		goto L1
	} else {
		goto L994
	}
L994:
	;
	if v5308 == v5159 {
		goto L995
	} else {
		goto L996
	}
L995:
	;
	v5319 = *(*int64)(unsafe.Add(mBase, uint32(v5144)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v5315)+24)) = v5319
	v5321 = *(*int64)(unsafe.Add(mBase, uint32(v5144)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v5315)+16)) = v5321
	v5323 = *(*int64)(unsafe.Add(mBase, uint32(v5144)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v5315)+8)) = v5323
	v5325 = *(*int64)(unsafe.Add(mBase, uint32(v5144)))
	*(*int64)(unsafe.Add(mBase, uint32(v5315))) = v5325
	v5086 = v5312
	v5107 = v5316
	goto L946
L996:
	;
	goto L997
L997:
	;
	if v5308 == v5194 {
		goto L998
	} else {
		goto L999
	}
L998:
	;
	v5328 = *(*int64)(unsafe.Add(mBase, uint32(v5179)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v5315)+24)) = v5328
	v5330 = *(*int64)(unsafe.Add(mBase, uint32(v5179)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v5315)+16)) = v5330
	v5332 = *(*int64)(unsafe.Add(mBase, uint32(v5179)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v5315)+8)) = v5332
	v5334 = *(*int64)(unsafe.Add(mBase, uint32(v5179)))
	*(*int64)(unsafe.Add(mBase, uint32(v5315))) = v5334
	v5086 = v5312
	v5107 = v5316
	goto L946
L999:
	;
	goto L1000
L1000:
	;
	v5336 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v5315)+4)) = uint16(v5312)
	*(*int32)(unsafe.Add(mBase, uint32(v5315))) = v5336
	v5339 = F_exprType(m, v5308)
	mBase = m.M
	v5340 = m.ExcPending
	if v5340 != 0 {
		goto L1
	} else {
		goto L1001
	}
L1001:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5315)+8)) = v5339
	v5342 = F_exprTypmod(m, v5308)
	mBase = m.M
	v5343 = m.ExcPending
	if v5343 != 0 {
		goto L1
	} else {
		goto L1002
	}
L1002:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5315)+12)) = v5342
	v5345 = F_exprCollation(m, v5308)
	mBase = m.M
	v5346 = m.ExcPending
	if v5346 != 0 {
		goto L1
	} else {
		goto L1003
	}
L1003:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5315)+16)) = v5345
	v5348 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v5315)+28)) = uint16(v5312)
	*(*int32)(unsafe.Add(mBase, uint32(v5315)+24)) = v5348
	v5086 = v5312
	v5107 = v5316
	goto L946
L1004:
	;
	v5399 = v5397 + v5360
	v5403 = F_extractRemainingColumns(m, l0, v3897, v3901, v34+int32(268), v5389, v5391, v4154+v5399<<(uint(int32(5))%32))
	mBase = m.M
	v5404 = m.ExcPending
	if v5404 != 0 {
		goto L1
	} else {
		goto L1005
	}
L1005:
	;
	v5405 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v5405 == int32(0) {
		goto L1006
	} else {
		goto L1007
	}
L1006:
	;
	v5597 = *(*int32)(unsafe.Add(mBase, uint32(v34)+276))
	v5598 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5599 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v5599 != 0 {
		goto L1019
	} else {
		goto L1020
	}
L1007:
	;
	v5408 = v5399 + v5403
	if v5408 <= int32(0) {
		goto L1006
	} else {
		goto L1008
	}
L1008:
	;
	v5411 = int32(3)
	v5412 = v5408 & v5411
	v5413 = int32(0)
	if base.Ui32(v5411) <= base.Ui32(v5397+v5403+v5360-int32(1)) {
		goto L1009
	} else {
		goto L1010
	}
L1009:
	;
	v5428 = v5413
	v5431 = int32(0)
	goto L1012
L1010:
	;
	v5496 = v5413
	goto L1011
L1011:
	;
	v5527 = v5496
	v5533 = v5413
	goto L1016
L1012:
	;
	v5455 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v5456 = int32(5)
	v5458 = v4154 + v5428<<(uint(v5456)%32)
	v5460 = v5428 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5458)+28)) = uint16(v5460)
	*(*int32)(unsafe.Add(mBase, uint32(v5458)+24)) = v5455
	v5463 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v5466 = v4154 + v5460<<(uint(v5456)%32)
	v5468 = v5428 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v5466)+28)) = uint16(v5468)
	*(*int32)(unsafe.Add(mBase, uint32(v5466)+24)) = v5463
	v5471 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v5474 = v4154 + v5468<<(uint(v5456)%32)
	v5476 = v5428 | int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v5474)+28)) = uint16(v5476)
	*(*int32)(unsafe.Add(mBase, uint32(v5474)+24)) = v5471
	v5479 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v5482 = v4154 + v5476<<(uint(v5456)%32)
	v5483 = int32(4)
	v5484 = v5428 + v5483
	*(*uint16)(unsafe.Add(mBase, uint32(v5482)+28)) = uint16(v5484)
	*(*int32)(unsafe.Add(mBase, uint32(v5482)+24)) = v5479
	v5488 = v5431 + v5483
	if v5488 != v5408&int32(2147483644) {
		v5428 = v5484
		v5431 = v5488
		goto L1012
	} else {
		goto L1014
	}
L1013:
	;
	if v5412 == int32(0) {
		goto L1006
	} else {
		goto L1015
	}
L1014:
	;
	goto L1013
L1015:
	;
	v5496 = v5484
	goto L1011
L1016:
	;
	v5554 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v5557 = v4154 + v5527<<(uint(int32(5))%32)
	v5558 = int32(1)
	v5559 = v5527 + v5558
	*(*uint16)(unsafe.Add(mBase, uint32(v5557)+28)) = uint16(v5559)
	*(*int32)(unsafe.Add(mBase, uint32(v5557)+24)) = v5554
	v5563 = v5533 + v5558
	if v5563 != v5412 {
		v5527 = v5559
		v5533 = v5563
		goto L1016
	} else {
		goto L1018
	}
L1017:
	;
	goto L1006
L1018:
	;
	goto L1017
L1019:
	;
	v5600 = *(*int32)(unsafe.Add(mBase, uint32(v5599)+4))
	v5602 = v5600
	goto L1021
L1020:
	;
	v5602 = int32(0)
	goto L1021
L1021:
	;
	v5603 = *(*int32)(unsafe.Add(mBase, uint32(v34)+264))
	v5604 = *(*int32)(unsafe.Add(mBase, uint32(v34)+272))
	v5605 = *(*int32)(unsafe.Add(mBase, uint32(v34)+268))
	v5606 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v5607 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v5609 = F_addRangeTableEntryForJoin(m, l0, v5597, v4154, v5598, v5602, v5603, v5604, v5605, v5606, v5607, int32(1))
	mBase = m.M
	v5610 = m.ExcPending
	if v5610 != 0 {
		goto L1
	} else {
		goto L1022
	}
L1022:
	;
	v5611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v5611 != 0 {
		goto L1023
	} else {
		goto L1024
	}
L1023:
	;
	v5612 = *(*int32)(unsafe.Add(mBase, uint32(v5611)+4))
	v5615 = v5612 + int32(1)
	goto L1025
L1024:
	;
	v5615 = int32(1)
	goto L1025
L1025:
	;
	v5616 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v5615 < v5616 {
		goto L1026
	} else {
		goto L1027
	}
L1026:
	;
	v5622 = v5611
	v5628 = v5615
	goto L1029
L1027:
	;
	v5661 = v5611
	goto L1028
L1028:
	;
	v5688 = F_lappend(m, v5661, l1)
	mBase = m.M
	v5689 = m.ExcPending
	if v5689 != 0 {
		goto L1
	} else {
		goto L1033
	}
L1029:
	;
	v5650 = F_lappend(m, v5622, int32(0))
	mBase = m.M
	v5651 = m.ExcPending
	if v5651 != 0 {
		goto L1
	} else {
		goto L1031
	}
L1030:
	;
	v5661 = v5650
	goto L1028
L1031:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v5650
	v5654 = v5628 + int32(1)
	v5655 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v5654 < v5655 {
		v5622 = v5650
		v5628 = v5654
		goto L1029
	} else {
		goto L1032
	}
L1032:
	;
	goto L1030
L1033:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v5688
	v5691 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v5691 != 0 {
		goto L1034
	} else {
		goto L1035
	}
L1034:
	;
	v5693 = F_palloc(m, int32(28))
	mBase = m.M
	v5694 = m.ExcPending
	if v5694 != 0 {
		goto L1
	} else {
		goto L1037
	}
L1035:
	;
	v5719 = v3894
	goto L1036
L1036:
	;
	v5720 = int32(0)
	v5722 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v5722 != 0 {
		v5818 = v5720
		v5839 = int32(1)
		goto L1041
	} else {
		goto L1042
	}
L1037:
	;
	v5695 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v5693))) = v5695
	v5697 = *(*int32)(unsafe.Add(mBase, uint32(v5609)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5693)+4)) = v5697
	v5699 = *(*int32)(unsafe.Add(mBase, uint32(v5609)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v5693)+20)) = int64(16777473)
	*(*int32)(unsafe.Add(mBase, uint32(v5693)+16)) = v4154
	*(*int32)(unsafe.Add(mBase, uint32(v5693)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5693)+8)) = v5699
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v5693
	*(*int32)(unsafe.Add(mBase, uint32(v34)+260)) = v5693
	v5711 = F_list_make1_impl(m, int32(1), v34+int32(44))
	mBase = m.M
	v5712 = m.ExcPending
	if v5712 != 0 {
		goto L1
	} else {
		goto L1038
	}
L1038:
	;
	F_checkNameSpaceConflicts(m, v5711, v3894)
	mBase = m.M
	v5714 = m.ExcPending
	if v5714 != 0 {
		goto L1
	} else {
		goto L1039
	}
L1039:
	;
	v5715 = F_lappend(m, v3894, v5693)
	mBase = m.M
	v5716 = m.ExcPending
	if v5716 != 0 {
		goto L1
	} else {
		goto L1040
	}
L1040:
	;
	v5719 = v5715
	goto L1036
L1041:
	;
	v5840 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v5609)+23)) = uint8(v5840)
	*(*uint16)(unsafe.Add(mBase, uint32(v5609)+21)) = uint16(v5840)
	*(*uint8)(unsafe.Add(mBase, uint32(v5609)+20)) = uint8(v5839)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v5609
	v5846 = F_lappend(m, v5818, v5609)
	mBase = m.M
	v5847 = m.ExcPending
	if v5847 != 0 {
		goto L1
	} else {
		goto L1050
	}
L1042:
	;
	v5723 = int32(0)
	if v5719 == v5723 {
		v5818 = v5720
		v5839 = v5723
		goto L1041
	} else {
		goto L1043
	}
L1043:
	;
	v5726 = int32(0)
	v5727 = *(*int32)(unsafe.Add(mBase, uint32(v5719)+4))
	if v5726 < v5727 {
		goto L1044
	} else {
		goto L1045
	}
L1044:
	;
	v5734 = v5726
	goto L1047
L1045:
	;
	v5807 = int32(0)
	goto L1046
L1046:
	;
	v5818 = v5719
	v5839 = v5807
	goto L1041
L1047:
	;
	v5761 = *(*int32)(unsafe.Add(mBase, uint32(v5719)+12))
	v5765 = *(*int32)(unsafe.Add(mBase, uint32(v5761+v5734<<(uint(int32(2))%32))))
	v5766 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5765)+21)) = uint8(v5766)
	v5769 = v5734 + int32(1)
	v5770 = *(*int32)(unsafe.Add(mBase, uint32(v5719)+4))
	if v5769 < v5770 {
		v5734 = v5769
		goto L1047
	} else {
		goto L1049
	}
L1048:
	;
	v5772 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v5807 = base.B2i32(v5772 != int32(0))
	goto L1046
L1049:
	;
	goto L1048
L1050:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5846
	v6521 = l1
	goto L3
L1051:
	;
	v5852 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v5853 = *(*int32)(unsafe.Add(mBase, uint32(v5852)+4))
	v5854 = *(*int32)(unsafe.Add(mBase, uint32(v5853)+12))
	if v5854 != 0 {
		goto L1053
	} else {
		goto L1054
	}
L1052:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6082 = m.ExcPending
	if v6082 != 0 {
		goto L1
	} else {
		goto L1120
	}
L1053:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6062 = m.ExcPending
	if v6062 != 0 {
		goto L1
	} else {
		goto L1115
	}
L1054:
	;
	v5855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5853)+21)))
	v5857 = v5855 - int32(109)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v5857))|base.B2i32(int32(1)<<(uint(v5857)%32)&int32(41) == int32(0)) != 0 {
		goto L1053
	} else {
		goto L1055
	}
L1055:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+332)) = int32(2281)
	v5869 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v5870 = int32(1)
	v5874 = F_LookupFuncName(m, v5869, v5870, v34+int32(332), v5870)
	mBase = m.M
	v5875 = m.ExcPending
	if v5875 != 0 {
		goto L1
	} else {
		goto L1056
	}
L1056:
	;
	if v5874 != 0 {
		goto L1057
	} else {
		goto L1058
	}
L1057:
	;
	v5876 = F_get_func_rettype(m, v5874)
	mBase = m.M
	v5877 = m.ExcPending
	if v5877 != 0 {
		goto L1
	} else {
		goto L1060
	}
L1058:
	;
	goto L1059
L1059:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6037 = m.ExcPending
	if v6037 != 0 {
		goto L1
	} else {
		goto L1109
	}
L1060:
	;
	if v5876 == int32(3310) {
		goto L1061
	} else {
		goto L1062
	}
L1061:
	;
	v5880 = F_GetTsmRoutine(m, v5874)
	mBase = m.M
	v5881 = m.ExcPending
	if v5881 != 0 {
		goto L1
	} else {
		goto L1064
	}
L1062:
	;
	goto L1063
L1063:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6011 = m.ExcPending
	if v6011 != 0 {
		goto L1
	} else {
		goto L1103
	}
L1064:
	;
	v5883 = F_palloc0(m, int32(16))
	mBase = m.M
	v5884 = m.ExcPending
	if v5884 != 0 {
		goto L1
	} else {
		goto L1065
	}
L1065:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5883)+4)) = v5874
	*(*int32)(unsafe.Add(mBase, uint32(v5883))) = int32(104)
	v5888 = int32(0)
	v5889 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v5889 != 0 {
		goto L1066
	} else {
		goto L1067
	}
L1066:
	;
	v5890 = *(*int32)(unsafe.Add(mBase, uint32(v5889)+4))
	v5891 = v5890
	goto L1068
L1067:
	;
	v5891 = v5
	goto L1068
L1068:
	;
	v5892 = *(*int32)(unsafe.Add(mBase, uint32(v5880)+4))
	if v5892 != 0 {
		goto L1069
	} else {
		goto L1070
	}
L1069:
	;
	v5893 = *(*int32)(unsafe.Add(mBase, uint32(v5892)+4))
	v5895 = v5893
	goto L1071
L1070:
	;
	v5895 = int32(0)
	goto L1071
L1071:
	;
	if v5895 != v5891 {
		goto L1052
	} else {
		goto L1072
	}
L1072:
	;
	v5901 = v5888
	v5903 = v5
	goto L1073
L1073:
	;
	v5928 = int32(0)
	if v5889 == v5928 {
		v5938 = v5928
		goto L1075
	} else {
		goto L1076
	}
L1074:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5987 = m.ExcPending
	if v5987 != 0 {
		goto L1
	} else {
		goto L1097
	}
L1075:
	;
	if v5892 == int32(0) {
		goto L1081
	} else {
		goto L1082
	}
L1076:
	;
	v5932 = *(*int32)(unsafe.Add(mBase, uint32(v5889)+4))
	if v5932 <= v5901 {
		v5938 = int32(0)
		goto L1075
	} else {
		goto L1077
	}
L1077:
	;
	v5934 = *(*int32)(unsafe.Add(mBase, uint32(v5889)+12))
	v5938 = v5934 + v5901<<(uint(int32(2))%32)
	goto L1075
L1078:
	;
	goto L1074
L1079:
	;
	v5970 = *(*int32)(unsafe.Add(mBase, uint32(v5947+v5901<<(uint(int32(2))%32))))
	v5971 = *(*int32)(unsafe.Add(mBase, uint32(v5938)))
	v5973 = F_transformExpr(m, l0, v5971, int32(5))
	mBase = m.M
	v5974 = m.ExcPending
	if v5974 != 0 {
		goto L1
	} else {
		goto L1093
	}
L1080:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5883)+8)) = v5949
	v5951 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v5951 != 0 {
		goto L1086
	} else {
		goto L1087
	}
L1081:
	;
	v5949 = int32(0)
	goto L1080
L1082:
	;
	goto L1083
L1083:
	;
	v5944 = *(*int32)(unsafe.Add(mBase, uint32(v5892)+4))
	if base.B2i32(v5938 == int32(0))|base.B2i32(v5944 <= v5901) != 0 {
		v5949 = v5903
		goto L1080
	} else {
		goto L1084
	}
L1084:
	;
	v5947 = *(*int32)(unsafe.Add(mBase, uint32(v5892)+12))
	if v5947 != 0 {
		goto L1079
	} else {
		goto L1085
	}
L1085:
	;
	v5949 = v5903
	goto L1080
L1086:
	;
	v5952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5880)+8)))
	if v5952 == int32(0) {
		goto L1078
	} else {
		goto L1089
	}
L1087:
	;
	v5964 = v5
	goto L1088
L1088:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5883)+12)) = v5964
	*(*int32)(unsafe.Add(mBase, uint32(v5853)+32)) = v5883
	v6521 = v5850
	goto L3
L1089:
	;
	v5956 = F_transformExpr(m, l0, v5951, int32(5))
	mBase = m.M
	v5957 = m.ExcPending
	if v5957 != 0 {
		goto L1
	} else {
		goto L1090
	}
L1090:
	;
	v5960 = F_coerce_to_specific_type(m, l0, v5956, int32(701), int32(_a_F_transformFromClauseItem_58))
	mBase = m.M
	v5961 = m.ExcPending
	if v5961 != 0 {
		goto L1
	} else {
		goto L1091
	}
L1091:
	;
	F_assign_expr_collations(m, l0, v5960)
	mBase = m.M
	v5963 = m.ExcPending
	if v5963 != 0 {
		goto L1
	} else {
		goto L1092
	}
L1092:
	;
	v5964 = v5960
	goto L1088
L1093:
	;
	v5976 = F_coerce_to_specific_type(m, l0, v5973, v5970, int32(_a_F_transformFromClauseItem_59))
	mBase = m.M
	v5977 = m.ExcPending
	if v5977 != 0 {
		goto L1
	} else {
		goto L1094
	}
L1094:
	;
	F_assign_expr_collations(m, l0, v5976)
	mBase = m.M
	v5979 = m.ExcPending
	if v5979 != 0 {
		goto L1
	} else {
		goto L1095
	}
L1095:
	;
	v5982 = F_lappend(m, v5903, v5976)
	mBase = m.M
	v5983 = m.ExcPending
	if v5983 != 0 {
		goto L1
	} else {
		goto L1096
	}
L1096:
	;
	v5901 = v5901 + int32(1)
	v5903 = v5982
	goto L1073
L1097:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v5990 = m.ExcPending
	if v5990 != 0 {
		goto L1
	} else {
		goto L1098
	}
L1098:
	;
	v5991 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v5992 = F_NameListToString(m, v5991)
	mBase = m.M
	v5993 = m.ExcPending
	if v5993 != 0 {
		goto L1
	} else {
		goto L1099
	}
L1099:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+176)) = v5992
	F_errmsg(m, int32(_a_F_transformFromClauseItem_60), v34+int32(176))
	mBase = m.M
	v5999 = m.ExcPending
	if v5999 != 0 {
		goto L1
	} else {
		goto L1100
	}
L1100:
	;
	v6000 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_parser_errposition(m, l0, v6000)
	mBase = m.M
	v6002 = m.ExcPending
	if v6002 != 0 {
		goto L1
	} else {
		goto L1101
	}
L1101:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(991), int32(_a_F_transformFromClauseItem_61))
	mBase = m.M
	v6007 = m.ExcPending
	if v6007 != 0 {
		goto L1
	} else {
		goto L1102
	}
L1102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1103:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v6014 = m.ExcPending
	if v6014 != 0 {
		goto L1
	} else {
		goto L1104
	}
L1104:
	;
	v6015 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6016 = F_NameListToString(m, v6015)
	mBase = m.M
	v6017 = m.ExcPending
	if v6017 != 0 {
		goto L1
	} else {
		goto L1105
	}
L1105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = int32(_a_F_transformFromClauseItem_62)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v6016
	F_errmsg(m, int32(_a_F_transformFromClauseItem_63), v34+int32(208))
	mBase = m.M
	v6025 = m.ExcPending
	if v6025 != 0 {
		goto L1
	} else {
		goto L1106
	}
L1106:
	;
	v6026 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_parser_errposition(m, l0, v6026)
	mBase = m.M
	v6028 = m.ExcPending
	if v6028 != 0 {
		goto L1
	} else {
		goto L1107
	}
L1107:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(943), int32(_a_F_transformFromClauseItem_61))
	mBase = m.M
	v6033 = m.ExcPending
	if v6033 != 0 {
		goto L1
	} else {
		goto L1108
	}
L1108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1109:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v6040 = m.ExcPending
	if v6040 != 0 {
		goto L1
	} else {
		goto L1110
	}
L1110:
	;
	v6041 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6042 = F_NameListToString(m, v6041)
	mBase = m.M
	v6043 = m.ExcPending
	if v6043 != 0 {
		goto L1
	} else {
		goto L1111
	}
L1111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+160)) = v6042
	F_errmsg(m, int32(_a_F_transformFromClauseItem_64), v34+int32(160))
	mBase = m.M
	v6049 = m.ExcPending
	if v6049 != 0 {
		goto L1
	} else {
		goto L1112
	}
L1112:
	;
	v6050 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_parser_errposition(m, l0, v6050)
	mBase = m.M
	v6052 = m.ExcPending
	if v6052 != 0 {
		goto L1
	} else {
		goto L1113
	}
L1113:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(935), int32(_a_F_transformFromClauseItem_61))
	mBase = m.M
	v6057 = m.ExcPending
	if v6057 != 0 {
		goto L1
	} else {
		goto L1114
	}
L1114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1115:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v6065 = m.ExcPending
	if v6065 != 0 {
		goto L1
	} else {
		goto L1116
	}
L1116:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_65), int32(0))
	mBase = m.M
	v6069 = m.ExcPending
	if v6069 != 0 {
		goto L1
	} else {
		goto L1117
	}
L1117:
	;
	v6070 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6071 = F_exprLocation(m, v6070)
	mBase = m.M
	F_parser_errposition(m, l0, v6071)
	mBase = m.M
	v6073 = m.ExcPending
	if v6073 != 0 {
		goto L1
	} else {
		goto L1118
	}
L1118:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1143), int32(_a_F_transformFromClauseItem_40))
	mBase = m.M
	v6078 = m.ExcPending
	if v6078 != 0 {
		goto L1
	} else {
		goto L1119
	}
L1119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1120:
	;
	F_errcode(m, int32(403177602))
	mBase = m.M
	v6085 = m.ExcPending
	if v6085 != 0 {
		goto L1
	} else {
		goto L1121
	}
L1121:
	;
	v6087 = *(*int32)(unsafe.Add(mBase, uint32(v5880)+4))
	if v6087 != 0 {
		goto L1122
	} else {
		goto L1123
	}
L1122:
	;
	v6088 = *(*int32)(unsafe.Add(mBase, uint32(v6087)+4))
	v6089 = v6088
	goto L1124
L1123:
	;
	v6089 = int32(0)
	goto L1124
L1124:
	;
	v6090 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6091 = F_NameListToString(m, v6090)
	mBase = m.M
	v6092 = m.ExcPending
	if v6092 != 0 {
		goto L1
	} else {
		goto L1125
	}
L1125:
	;
	v6093 = *(*int32)(unsafe.Add(mBase, uint32(v5880)+4))
	if v6093 != 0 {
		goto L1126
	} else {
		goto L1127
	}
L1126:
	;
	v6094 = *(*int32)(unsafe.Add(mBase, uint32(v6093)+4))
	v6095 = v6094
	goto L1128
L1127:
	;
	v6095 = v5888
	goto L1128
L1128:
	;
	v6096 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v6096 != 0 {
		goto L1129
	} else {
		goto L1130
	}
L1129:
	;
	v6097 = *(*int32)(unsafe.Add(mBase, uint32(v6096)+4))
	v6099 = v6097
	goto L1131
L1130:
	;
	v6099 = int32(0)
	goto L1131
L1131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v6099
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v6095
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v6091
	F_errmsg_plural(m, int32(_a_F_transformFromClauseItem_66), int32(_a_F_transformFromClauseItem_67), v6089, v34+int32(192))
	mBase = m.M
	v6108 = m.ExcPending
	if v6108 != 0 {
		goto L1
	} else {
		goto L1132
	}
L1132:
	;
	v6109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_parser_errposition(m, l0, v6109)
	mBase = m.M
	v6111 = m.ExcPending
	if v6111 != 0 {
		goto L1
	} else {
		goto L1133
	}
L1133:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(961), int32(_a_F_transformFromClauseItem_61))
	mBase = m.M
	v6116 = m.ExcPending
	if v6116 != 0 {
		goto L1
	} else {
		goto L1134
	}
L1134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1135:
	;
	v6150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v6150 != 0 {
		goto L1136
	} else {
		goto L1137
	}
L1136:
	;
	v6151 = *(*int32)(unsafe.Add(mBase, uint32(v6150)+4))
	if v6151 <= int32(0) {
		goto L1140
	} else {
		goto L1141
	}
L1137:
	;
	goto L1138
L1138:
	;
	v6457 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+68)) = v6457
	v6459 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v6459)
	v6462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v6462 == v6459 {
		goto L1185
	} else {
		goto L1186
	}
L1139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+12)) = v6415
	*(*int32)(unsafe.Add(mBase, uint32(v3307)+8)) = v6412
	goto L1138
L1140:
	;
	v6412 = v5
	v6415 = v5
	goto L1139
L1141:
	;
	goto L1142
L1142:
	;
	v6160 = v5
	v6164 = int32(0)
	v6174 = v5
	v6177 = v5
	goto L1143
L1143:
	;
	v6186 = *(*int32)(unsafe.Add(mBase, uint32(v6150)+12))
	v6190 = *(*int32)(unsafe.Add(mBase, uint32(v6186+v6164<<(uint(int32(2))%32))))
	v6191 = *(*int32)(unsafe.Add(mBase, uint32(v6190)+12))
	v6193 = F_transformExpr(m, l0, v6191, int32(5))
	mBase = m.M
	v6194 = m.ExcPending
	if v6194 != 0 {
		goto L1
	} else {
		goto L1145
	}
L1144:
	;
	v6412 = v6201
	v6415 = v6387
	goto L1139
L1145:
	;
	v6197 = F_coerce_to_specific_type(m, l0, v6193, int32(25), int32(_a_F_transformFromClauseItem_35))
	mBase = m.M
	v6198 = m.ExcPending
	if v6198 != 0 {
		goto L1
	} else {
		goto L1146
	}
L1146:
	;
	F_assign_expr_collations(m, l0, v6197)
	mBase = m.M
	v6200 = m.ExcPending
	if v6200 != 0 {
		goto L1
	} else {
		goto L1147
	}
L1147:
	;
	v6201 = F_lappend(m, v6174, v6197)
	mBase = m.M
	v6202 = m.ExcPending
	if v6202 != 0 {
		goto L1
	} else {
		goto L1148
	}
L1148:
	;
	v6203 = *(*int32)(unsafe.Add(mBase, uint32(v6190)+4))
	if v6203 != 0 {
		goto L1151
	} else {
		goto L1152
	}
L1149:
	;
	v6387 = F_lappend(m, v6177, v6360)
	mBase = m.M
	v6388 = m.ExcPending
	if v6388 != 0 {
		goto L1
	} else {
		goto L1183
	}
L1150:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6337 = m.ExcPending
	if v6337 != 0 {
		goto L1
	} else {
		goto L1178
	}
L1151:
	;
	if v6177 == int32(0) {
		goto L1154
	} else {
		goto L1155
	}
L1152:
	;
	goto L1153
L1153:
	;
	v6311 = int32(0)
	if v6160 == v6311 {
		v6360 = v6311
		v6361 = int32(1)
		goto L1149
	} else {
		goto L1172
	}
L1154:
	;
	v6309 = F_makeString(m, v6203)
	mBase = m.M
	v6310 = m.ExcPending
	if v6310 != 0 {
		goto L1
	} else {
		goto L1171
	}
L1155:
	;
	v6206 = *(*int32)(unsafe.Add(mBase, uint32(v6177)+4))
	if v6206 <= int32(0) {
		goto L1154
	} else {
		goto L1156
	}
L1156:
	;
	v6209 = *(*int32)(unsafe.Add(mBase, uint32(v6177)+12))
	v6215 = int32(0)
	goto L1157
L1157:
	;
	v6245 = *(*int32)(unsafe.Add(mBase, uint32(v6209+v6215<<(uint(int32(2))%32))))
	if v6245 != 0 {
		goto L1159
	} else {
		goto L1160
	}
L1158:
	;
	goto L1154
L1159:
	;
	v6246 = *(*int32)(unsafe.Add(mBase, uint32(v6245)+4))
	v6249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6246))))
	v6252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6203))))
	if base.B2i32(v6249 == int32(0))|base.B2i32(v6249 != v6252) != 0 {
		v6270 = v6249
		v6271 = v6252
		goto L1163
	} else {
		goto L1164
	}
L1160:
	;
	goto L1161
L1161:
	;
	v6276 = v6215 + int32(1)
	if v6206 != v6276 {
		v6215 = v6276
		goto L1157
	} else {
		goto L1170
	}
L1162:
	;
	if v6270-v6271 == int32(0) {
		goto L1150
	} else {
		goto L1169
	}
L1163:
	;
	goto L1162
L1164:
	;
	v6255 = v6246
	v6256 = v6203
	goto L1165
L1165:
	;
	v6259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6256)+1)))
	v6260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6255)+1)))
	if v6260 == int32(0) {
		v6270 = v6260
		v6271 = v6259
		goto L1163
	} else {
		goto L1167
	}
L1166:
	;
	v6270 = v6260
	v6271 = v6259
	goto L1163
L1167:
	;
	v6263 = int32(1)
	if v6260 == v6259 {
		v6255 = v6255 + v6263
		v6256 = v6256 + v6263
		goto L1165
	} else {
		goto L1168
	}
L1168:
	;
	goto L1166
L1169:
	;
	goto L1161
L1170:
	;
	goto L1158
L1171:
	;
	v6360 = v6309
	v6361 = v6160
	goto L1149
L1172:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6318 = m.ExcPending
	if v6318 != 0 {
		goto L1
	} else {
		goto L1173
	}
L1173:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v6321 = m.ExcPending
	if v6321 != 0 {
		goto L1
	} else {
		goto L1174
	}
L1174:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_68), int32(0))
	mBase = m.M
	v6325 = m.ExcPending
	if v6325 != 0 {
		goto L1
	} else {
		goto L1175
	}
L1175:
	;
	v6326 = *(*int32)(unsafe.Add(mBase, uint32(v6190)+16))
	F_parser_errposition(m, l0, v6326)
	mBase = m.M
	v6328 = m.ExcPending
	if v6328 != 0 {
		goto L1
	} else {
		goto L1176
	}
L1176:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(874), int32(_a_F_transformFromClauseItem_37))
	mBase = m.M
	v6333 = m.ExcPending
	if v6333 != 0 {
		goto L1
	} else {
		goto L1177
	}
L1177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1178:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v6340 = m.ExcPending
	if v6340 != 0 {
		goto L1
	} else {
		goto L1179
	}
L1179:
	;
	v6341 = *(*int32)(unsafe.Add(mBase, uint32(v6190)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+224)) = v6341
	F_errmsg(m, int32(_a_F_transformFromClauseItem_69), v34+int32(224))
	mBase = m.M
	v6347 = m.ExcPending
	if v6347 != 0 {
		goto L1
	} else {
		goto L1180
	}
L1180:
	;
	v6348 = *(*int32)(unsafe.Add(mBase, uint32(v6190)+16))
	F_parser_errposition(m, l0, v6348)
	mBase = m.M
	v6350 = m.ExcPending
	if v6350 != 0 {
		goto L1
	} else {
		goto L1181
	}
L1181:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(865), int32(_a_F_transformFromClauseItem_37))
	mBase = m.M
	v6355 = m.ExcPending
	if v6355 != 0 {
		goto L1
	} else {
		goto L1182
	}
L1182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1183:
	;
	v6390 = v6164 + int32(1)
	v6391 = *(*int32)(unsafe.Add(mBase, uint32(v6150)+4))
	if v6390 < v6391 {
		v6160 = v6361
		v6164 = v6390
		v6174 = v6201
		v6177 = v6387
		goto L1143
	} else {
		goto L1184
	}
L1184:
	;
	goto L1144
L1185:
	;
	v6466 = F_contain_vars_of_level(m, v3307, int32(0))
	mBase = m.M
	v6467 = m.ExcPending
	if v6467 != 0 {
		goto L1
	} else {
		goto L1188
	}
L1186:
	;
	v6468 = int32(1)
	goto L1187
L1187:
	;
	v6469 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v6470 = F_addRangeTableEntryForTableFunc(m, l0, v3307, v6469, v6468)
	mBase = m.M
	v6471 = m.ExcPending
	if v6471 != 0 {
		goto L1
	} else {
		goto L1189
	}
L1188:
	;
	v6468 = v6466
	goto L1187
L1189:
	;
	v6503 = v6470
	goto L4
L1190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v6510
	v6514 = F_palloc0(m, int32(8))
	mBase = m.M
	v6515 = m.ExcPending
	if v6515 != 0 {
		goto L1
	} else {
		goto L1191
	}
L1191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6514))) = int32(63)
	v6518 = *(*int32)(unsafe.Add(mBase, uint32(v6503)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6514)+4)) = v6518
	v6521 = v6514
	goto L3
}
