package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_spgdoinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v371 int32
	_ = v371
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int64
	_ = v664
	var v665 int32
	_ = v665
	var v667 int64
	_ = v667
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v692 int32
	_ = v692
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v709 int32
	_ = v709
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v801 int32
	_ = v801
	var v852 int32
	_ = v852
	var v858 int32
	_ = v858
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1338 int32
	_ = v1338
	var v1343 int32
	_ = v1343
	var v1347 int32
	_ = v1347
	var v1353 int32
	_ = v1353
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1377 int64
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1380 int64
	_ = v1380
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1462 int32
	_ = v1462
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1497 int32
	_ = v1497
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1521 int32
	_ = v1521
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1580 int64
	_ = v1580
	var v1581 int64
	_ = v1581
	var v1582 int64
	_ = v1582
	var v1583 int64
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1593 int32
	_ = v1593
	var v1597 int64
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1603 int32
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1613 int32
	_ = v1613
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1634 int32
	_ = v1634
	var v1639 int32
	_ = v1639
	var v1646 int32
	_ = v1646
	var v1658 int32
	_ = v1658
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1711 int32
	_ = v1711
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1724 int32
	_ = v1724
	var v1729 int64
	_ = v1729
	var v1730 int64
	_ = v1730
	var v1731 int64
	_ = v1731
	var v1732 int64
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1742 int32
	_ = v1742
	var v1746 int64
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1762 int32
	_ = v1762
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1780 int32
	_ = v1780
	var v1785 int32
	_ = v1785
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1808 int32
	_ = v1808
	var v1814 int32
	_ = v1814
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1839 int32
	_ = v1839
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1892 int64
	_ = v1892
	var v1893 int64
	_ = v1893
	var v1894 int64
	_ = v1894
	var v1895 int64
	_ = v1895
	var v1899 int32
	_ = v1899
	var v1905 int32
	_ = v1905
	var v1909 int64
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1915 int32
	_ = v1915
	var v1920 int64
	_ = v1920
	var v1923 int32
	_ = v1923
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1940 int64
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1955 int32
	_ = v1955
	var v1963 int32
	_ = v1963
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2012 int32
	_ = v2012
	var v2019 int32
	_ = v2019
	var v2038 int32
	_ = v2038
	var v2042 int64
	_ = v2042
	var v2044 int32
	_ = v2044
	var v2047 int32
	_ = v2047
	var v2050 int32
	_ = v2050
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2065 int32
	_ = v2065
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2089 int32
	_ = v2089
	var v2097 int32
	_ = v2097
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2146 int32
	_ = v2146
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2153 int32
	_ = v2153
	var v2156 int32
	_ = v2156
	var v2161 int32
	_ = v2161
	var v2174 int32
	_ = v2174
	var v2177 int32
	_ = v2177
	var v2180 int32
	_ = v2180
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2206 int32
	_ = v2206
	var v2210 int32
	_ = v2210
	var v2215 int32
	_ = v2215
	var v2231 int32
	_ = v2231
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2264 int32
	_ = v2264
	var v2274 int32
	_ = v2274
	var v2327 int32
	_ = v2327
	var v2330 int32
	_ = v2330
	var v2397 int32
	_ = v2397
	var v2399 int32
	_ = v2399
	var v2410 int32
	_ = v2410
	var v2460 int32
	_ = v2460
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2471 int32
	_ = v2471
	var v2477 int64
	_ = v2477
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2494 int32
	_ = v2494
	var v2544 int32
	_ = v2544
	var v2550 int32
	_ = v2550
	var v2551 int32
	_ = v2551
	var v2610 int32
	_ = v2610
	var v2611 int32
	_ = v2611
	var v2613 int32
	_ = v2613
	var v2614 int32
	_ = v2614
	var v2617 int32
	_ = v2617
	var v2618 int32
	_ = v2618
	var v2631 int32
	_ = v2631
	var v2635 int32
	_ = v2635
	var v2640 int32
	_ = v2640
	var v2656 int32
	_ = v2656
	var v2683 int32
	_ = v2683
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2691 int32
	_ = v2691
	var v2701 int32
	_ = v2701
	var v2754 int32
	_ = v2754
	var v2758 int64
	_ = v2758
	var v2760 int64
	_ = v2760
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2767 int32
	_ = v2767
	var v2768 int32
	_ = v2768
	var v2778 int32
	_ = v2778
	var v2827 int32
	_ = v2827
	var v2828 int64
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2830 int32
	_ = v2830
	var v2831 int32
	_ = v2831
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2910 int32
	_ = v2910
	var v2911 int32
	_ = v2911
	var v2974 int32
	_ = v2974
	var v2984 int32
	_ = v2984
	var v3035 int32
	_ = v3035
	var v3036 int32
	_ = v3036
	var v3038 int32
	_ = v3038
	var v3041 int32
	_ = v3041
	var v3043 int32
	_ = v3043
	var v3045 int32
	_ = v3045
	var v3046 int32
	_ = v3046
	var v3048 int32
	_ = v3048
	var v3049 int32
	_ = v3049
	var v3057 int32
	_ = v3057
	var v3116 int32
	_ = v3116
	var v3122 int32
	_ = v3122
	var v3126 int32
	_ = v3126
	var v3129 int32
	_ = v3129
	var v3130 int32
	_ = v3130
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3135 int32
	_ = v3135
	var v3138 int32
	_ = v3138
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3143 int32
	_ = v3143
	var v3145 int32
	_ = v3145
	var v3147 int32
	_ = v3147
	var v3152 int32
	_ = v3152
	var v3154 int32
	_ = v3154
	var v3155 int32
	_ = v3155
	var v3157 int32
	_ = v3157
	var v3160 int32
	_ = v3160
	var v3161 int32
	_ = v3161
	var v3162 int32
	_ = v3162
	var v3163 int32
	_ = v3163
	var v3166 int32
	_ = v3166
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3172 int32
	_ = v3172
	var v3173 int32
	_ = v3173
	var v3181 int32
	_ = v3181
	var v3189 int32
	_ = v3189
	var v3192 int32
	_ = v3192
	var v3193 int32
	_ = v3193
	var v3194 int32
	_ = v3194
	var v3196 int32
	_ = v3196
	var v3197 int32
	_ = v3197
	var v3198 int32
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3200 int32
	_ = v3200
	var v3204 int32
	_ = v3204
	var v3210 int32
	_ = v3210
	var v3212 int32
	_ = v3212
	var v3218 int32
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3227 int32
	_ = v3227
	var v3236 int32
	_ = v3236
	var v3237 int32
	_ = v3237
	var v3244 int32
	_ = v3244
	var v3287 int32
	_ = v3287
	var v3290 int32
	_ = v3290
	var v3291 int32
	_ = v3291
	var v3293 int32
	_ = v3293
	var v3295 int32
	_ = v3295
	var v3297 int32
	_ = v3297
	var v3299 int32
	_ = v3299
	var v3301 int32
	_ = v3301
	var v3302 int32
	_ = v3302
	var v3304 int32
	_ = v3304
	var v3305 int32
	_ = v3305
	var v3313 int32
	_ = v3313
	var v3321 int32
	_ = v3321
	var v3368 int32
	_ = v3368
	var v3369 int32
	_ = v3369
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3375 int32
	_ = v3375
	var v3378 int32
	_ = v3378
	var v3379 int32
	_ = v3379
	var v3381 int32
	_ = v3381
	var v3382 int32
	_ = v3382
	var v3392 int32
	_ = v3392
	var v3398 int32
	_ = v3398
	var v3400 int32
	_ = v3400
	var v3406 int32
	_ = v3406
	var v3407 int32
	_ = v3407
	var v3408 int32
	_ = v3408
	var v3409 int32
	_ = v3409
	var v3410 int32
	_ = v3410
	var v3413 int32
	_ = v3413
	var v3414 int32
	_ = v3414
	var v3415 int32
	_ = v3415
	var v3425 int32
	_ = v3425
	var v3428 int32
	_ = v3428
	var v3439 int32
	_ = v3439
	var v3475 int32
	_ = v3475
	var v3478 int32
	_ = v3478
	var v3479 int32
	_ = v3479
	var v3481 int32
	_ = v3481
	var v3483 int32
	_ = v3483
	var v3485 int32
	_ = v3485
	var v3487 int32
	_ = v3487
	var v3489 int32
	_ = v3489
	var v3490 int32
	_ = v3490
	var v3492 int32
	_ = v3492
	var v3493 int32
	_ = v3493
	var v3505 int32
	_ = v3505
	var v3516 int32
	_ = v3516
	var v3552 int32
	_ = v3552
	var v3559 int32
	_ = v3559
	var v3563 int32
	_ = v3563
	var v3568 int32
	_ = v3568
	var v3572 int32
	_ = v3572
	var v3576 int32
	_ = v3576
	var v3581 int32
	_ = v3581
	var v3585 int32
	_ = v3585
	var v3589 int32
	_ = v3589
	var v3594 int32
	_ = v3594
	var v3611 int32
	_ = v3611
	var v3626 int32
	_ = v3626
	var v3655 int32
	_ = v3655
	var v3656 int32
	_ = v3656
	var v3657 int32
	_ = v3657
	var v3671 int32
	_ = v3671
	var v3685 int32
	_ = v3685
	var v3722 int32
	_ = v3722
	var v3725 int32
	_ = v3725
	var v3727 int32
	_ = v3727
	var v3730 int32
	_ = v3730
	var v3735 int32
	_ = v3735
	var v3737 int32
	_ = v3737
	var v3740 int32
	_ = v3740
	var v3745 int32
	_ = v3745
	var v3747 int32
	_ = v3747
	var v3750 int32
	_ = v3750
	var v3755 int32
	_ = v3755
	var v3757 int32
	_ = v3757
	var v3759 int32
	_ = v3759
	var v3760 int32
	_ = v3760
	var v3762 int32
	_ = v3762
	var v3773 int32
	_ = v3773
	var v3830 int32
	_ = v3830
	var v3833 int32
	_ = v3833
	var v3884 int32
	_ = v3884
	var v3886 int32
	_ = v3886
	var v3888 int32
	_ = v3888
	var v3891 int32
	_ = v3891
	var v3901 int32
	_ = v3901
	var v3909 int32
	_ = v3909
	var v3924 int32
	_ = v3924
	var v3952 int32
	_ = v3952
	var v3956 int32
	_ = v3956
	var v3957 int32
	_ = v3957
	var v3958 int32
	_ = v3958
	var v3960 int32
	_ = v3960
	var v3964 int32
	_ = v3964
	var v3965 int32
	_ = v3965
	var v3968 int32
	_ = v3968
	var v3970 int32
	_ = v3970
	var v3972 int32
	_ = v3972
	var v3982 int32
	_ = v3982
	var v3987 int32
	_ = v3987
	var v3993 int32
	_ = v3993
	var v3995 int32
	_ = v3995
	var v4001 int32
	_ = v4001
	var v4005 int32
	_ = v4005
	var v4006 int32
	_ = v4006
	var v4007 int32
	_ = v4007
	var v4010 int32
	_ = v4010
	var v4017 int32
	_ = v4017
	var v4022 int32
	_ = v4022
	var v4023 int32
	_ = v4023
	var v4024 int32
	_ = v4024
	var v4029 int32
	_ = v4029
	var v4033 int32
	_ = v4033
	var v4038 int32
	_ = v4038
	var v4040 int32
	_ = v4040
	var v4041 int32
	_ = v4041
	var v4052 int32
	_ = v4052
	var v4065 int32
	_ = v4065
	var v4103 int32
	_ = v4103
	var v4104 int32
	_ = v4104
	var v4105 int32
	_ = v4105
	var v4106 int32
	_ = v4106
	var v4107 int32
	_ = v4107
	var v4108 int32
	_ = v4108
	var v4112 int32
	_ = v4112
	var v4118 int32
	_ = v4118
	var v4120 int32
	_ = v4120
	var v4121 int32
	_ = v4121
	var v4126 int32
	_ = v4126
	var v4127 int32
	_ = v4127
	var v4128 int32
	_ = v4128
	var v4130 int32
	_ = v4130
	var v4133 int32
	_ = v4133
	var v4134 int32
	_ = v4134
	var v4137 int32
	_ = v4137
	var v4140 int32
	_ = v4140
	var v4145 int32
	_ = v4145
	var v4148 int32
	_ = v4148
	var v4150 int32
	_ = v4150
	var v4159 int32
	_ = v4159
	var v4165 int32
	_ = v4165
	var v4167 int32
	_ = v4167
	var v4173 int32
	_ = v4173
	var v4174 int32
	_ = v4174
	var v4179 int32
	_ = v4179
	var v4183 int32
	_ = v4183
	var v4184 int32
	_ = v4184
	var v4186 int32
	_ = v4186
	var v4190 int32
	_ = v4190
	var v4192 int32
	_ = v4192
	var v4193 int32
	_ = v4193
	var v4195 int32
	_ = v4195
	var v4197 int32
	_ = v4197
	var v4198 int32
	_ = v4198
	var v4201 int32
	_ = v4201
	var v4203 int32
	_ = v4203
	var v4225 int32
	_ = v4225
	var v4263 int32
	_ = v4263
	var v4264 int32
	_ = v4264
	var v4270 int32
	_ = v4270
	var v4272 int32
	_ = v4272
	var v4273 int32
	_ = v4273
	var v4276 int32
	_ = v4276
	var v4282 int32
	_ = v4282
	var v4288 int32
	_ = v4288
	var v4291 int32
	_ = v4291
	var v4295 int32
	_ = v4295
	var v4300 int32
	_ = v4300
	var v4306 int32
	_ = v4306
	var v4308 int32
	_ = v4308
	var v4309 int32
	_ = v4309
	var v4314 int32
	_ = v4314
	var v4315 int32
	_ = v4315
	var v4319 int32
	_ = v4319
	var v4325 int32
	_ = v4325
	var v4327 int32
	_ = v4327
	var v4333 int32
	_ = v4333
	var v4334 int32
	_ = v4334
	var v4336 int32
	_ = v4336
	var v4337 int32
	_ = v4337
	var v4340 int32
	_ = v4340
	var v4348 int32
	_ = v4348
	var v4354 int32
	_ = v4354
	var v4357 int32
	_ = v4357
	var v4361 int32
	_ = v4361
	var v4366 int32
	_ = v4366
	var v4372 int32
	_ = v4372
	var v4374 int32
	_ = v4374
	var v4380 int32
	_ = v4380
	var v4384 int32
	_ = v4384
	var v4385 int32
	_ = v4385
	var v4386 int32
	_ = v4386
	var v4389 int32
	_ = v4389
	var v4391 int32
	_ = v4391
	var v4393 int32
	_ = v4393
	var v4396 int32
	_ = v4396
	var v4397 int32
	_ = v4397
	var v4401 int32
	_ = v4401
	var v4408 int32
	_ = v4408
	var v4409 int32
	_ = v4409
	var v4415 int32
	_ = v4415
	var v4420 int32
	_ = v4420
	var v4422 int32
	_ = v4422
	var v4423 int32
	_ = v4423
	var v4425 int32
	_ = v4425
	var v4426 int32
	_ = v4426
	var v4427 int32
	_ = v4427
	var v4429 int32
	_ = v4429
	var v4430 int32
	_ = v4430
	var v4431 int32
	_ = v4431
	var v4435 int32
	_ = v4435
	var v4438 int32
	_ = v4438
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
	var v4453 int32
	_ = v4453
	var v4454 int32
	_ = v4454
	var v4458 int32
	_ = v4458
	var v4459 int32
	_ = v4459
	var v4461 int32
	_ = v4461
	var v4462 int32
	_ = v4462
	var v4464 int32
	_ = v4464
	var v4467 int32
	_ = v4467
	var v4471 int32
	_ = v4471
	var v4472 int32
	_ = v4472
	var v4474 int32
	_ = v4474
	var v4478 int32
	_ = v4478
	var v4479 int32
	_ = v4479
	var v4481 int32
	_ = v4481
	var v4485 int32
	_ = v4485
	var v4486 int32
	_ = v4486
	var v4488 int32
	_ = v4488
	var v4489 int32
	_ = v4489
	var v4498 int32
	_ = v4498
	var v4501 int64
	_ = v4501
	var v4502 int32
	_ = v4502
	var v4506 int32
	_ = v4506
	var v4512 int32
	_ = v4512
	var v4514 int32
	_ = v4514
	var v4520 int32
	_ = v4520
	var v4531 int32
	_ = v4531
	var v4537 int32
	_ = v4537
	var v4539 int32
	_ = v4539
	var v4545 int32
	_ = v4545
	var v4547 int64
	_ = v4547
	var v4549 int64
	_ = v4549
	var v4555 int32
	_ = v4555
	var v4557 int32
	_ = v4557
	var v4562 int32
	_ = v4562
	var v4564 int32
	_ = v4564
	var v4566 int32
	_ = v4566
	var v4568 int32
	_ = v4568
	var v4571 int32
	_ = v4571
	var v4582 int32
	_ = v4582
	var v4583 int32
	_ = v4583
	var v4586 int32
	_ = v4586
	var v4590 int32
	_ = v4590
	var v4631 int32
	_ = v4631
	var v4642 int32
	_ = v4642
	var v4643 int32
	_ = v4643
	var v4646 int32
	_ = v4646
	var v4650 int32
	_ = v4650
	var v4689 int32
	_ = v4689
	var v4690 int32
	_ = v4690
	var v4694 int32
	_ = v4694
	var v4696 int32
	_ = v4696
	var v4701 int32
	_ = v4701
	var v4759 int32
	_ = v4759
	var v4760 int64
	_ = v4760
	var v4762 int64
	_ = v4762
	var v4767 int32
	_ = v4767
	var v4768 int32
	_ = v4768
	var v4772 int32
	_ = v4772
	var v4774 int32
	_ = v4774
	var v4779 int32
	_ = v4779
	var v4783 int32
	_ = v4783
	var v4784 int32
	_ = v4784
	var v4787 int64
	_ = v4787
	var v4790 int64
	_ = v4790
	var v4792 int32
	_ = v4792
	var v4798 int32
	_ = v4798
	var v4799 int32
	_ = v4799
	var v4800 int64
	_ = v4800
	var v4815 int32
	_ = v4815
	var v4816 int32
	_ = v4816
	var v4817 int64
	_ = v4817
	var v4818 int32
	_ = v4818
	var v4819 int32
	_ = v4819
	var v4820 int32
	_ = v4820
	var v4823 int32
	_ = v4823
	var v4824 int32
	_ = v4824
	var v4834 int32
	_ = v4834
	var v4838 int32
	_ = v4838
	var v4843 int32
	_ = v4843
	var v4844 int32
	_ = v4844
	var v4845 int64
	_ = v4845
	var v4852 int64
	_ = v4852
	var v4859 int64
	_ = v4859
	var v4861 int64
	_ = v4861
	var v4862 int64
	_ = v4862
	var v4865 int64
	_ = v4865
	var v4867 int64
	_ = v4867
	var v4871 int64
	_ = v4871
	var v4873 int64
	_ = v4873
	var v4881 int64
	_ = v4881
	var v4886 int64
	_ = v4886
	var v4899 int64
	_ = v4899
	var v4901 int32
	_ = v4901
	var v4902 int32
	_ = v4902
	var v4907 int32
	_ = v4907
	var v4908 int32
	_ = v4908
	var v4915 int32
	_ = v4915
	var v4917 int32
	_ = v4917
	var v4923 int32
	_ = v4923
	var v4928 int32
	_ = v4928
	var v4934 int32
	_ = v4934
	var v4937 int32
	_ = v4937
	var v4938 int32
	_ = v4938
	var v4940 int32
	_ = v4940
	var v4942 int32
	_ = v4942
	var v4944 int32
	_ = v4944
	var v4945 int32
	_ = v4945
	var v4946 int32
	_ = v4946
	var v4959 int32
	_ = v4959
	var v4960 int32
	_ = v4960
	var v5009 int32
	_ = v5009
	var v5010 int32
	_ = v5010
	var v5012 int32
	_ = v5012
	var v5013 int32
	_ = v5013
	var v5016 int32
	_ = v5016
	var v5017 int32
	_ = v5017
	var v5020 int32
	_ = v5020
	var v5021 int32
	_ = v5021
	var v5024 int32
	_ = v5024
	var v5026 int32
	_ = v5026
	var v5037 int32
	_ = v5037
	var v5095 int32
	_ = v5095
	var v5096 int32
	_ = v5096
	var v5145 int32
	_ = v5145
	var v5148 int32
	_ = v5148
	var v5150 int32
	_ = v5150
	var v5159 int32
	_ = v5159
	var v5218 int32
	_ = v5218
	var v5268 int32
	_ = v5268
	var v5269 int32
	_ = v5269
	var v5270 int32
	_ = v5270
	var v5275 int32
	_ = v5275
	var v5276 int32
	_ = v5276
	var v5279 int64
	_ = v5279
	var v5283 int32
	_ = v5283
	var v5284 int32
	_ = v5284
	var v5287 int32
	_ = v5287
	var v5290 int32
	_ = v5290
	var v5298 int32
	_ = v5298
	var v5306 int32
	_ = v5306
	var v5309 int32
	_ = v5309
	var v5310 int32
	_ = v5310
	var v5313 int32
	_ = v5313
	var v5323 int32
	_ = v5323
	var v5327 int32
	_ = v5327
	var v5332 int32
	_ = v5332
	var v5333 int32
	_ = v5333
	var v5336 int64
	_ = v5336
	var v5337 int32
	_ = v5337
	var v5341 int32
	_ = v5341
	var v5342 int32
	_ = v5342
	var v5346 int32
	_ = v5346
	var v5348 int32
	_ = v5348
	var v5352 int32
	_ = v5352
	var v5353 int32
	_ = v5353
	var v5354 int32
	_ = v5354
	var v5368 int32
	_ = v5368
	var v5369 int32
	_ = v5369
	var v5420 int32
	_ = v5420
	var v5424 int32
	_ = v5424
	var v5425 int32
	_ = v5425
	var v5429 int32
	_ = v5429
	var v5430 int32
	_ = v5430
	var v5497 int32
	_ = v5497
	var v5498 int32
	_ = v5498
	var v5501 int32
	_ = v5501
	var v5503 int32
	_ = v5503
	var v5506 int32
	_ = v5506
	var v5509 int64
	_ = v5509
	var v5511 int64
	_ = v5511
	var v5520 int32
	_ = v5520
	var v5521 int32
	_ = v5521
	var v5522 int32
	_ = v5522
	var v5524 int32
	_ = v5524
	var v5529 int32
	_ = v5529
	var v5530 int32
	_ = v5530
	var v5531 int32
	_ = v5531
	var v5532 int32
	_ = v5532
	var v5535 int32
	_ = v5535
	var v5536 int32
	_ = v5536
	var v5537 int32
	_ = v5537
	var v5540 int32
	_ = v5540
	var v5542 int32
	_ = v5542
	var v5547 int32
	_ = v5547
	var v5548 int32
	_ = v5548
	var v5550 int32
	_ = v5550
	var v5551 int32
	_ = v5551
	var v5554 int32
	_ = v5554
	var v5555 int32
	_ = v5555
	var v5556 int32
	_ = v5556
	var v5560 int32
	_ = v5560
	var v5563 int32
	_ = v5563
	var v5564 int32
	_ = v5564
	var v5565 int32
	_ = v5565
	var v5567 int32
	_ = v5567
	var v5572 int32
	_ = v5572
	var v5573 int32
	_ = v5573
	var v5575 int32
	_ = v5575
	var v5579 int32
	_ = v5579
	var v5582 int64
	_ = v5582
	var v5583 int32
	_ = v5583
	var v5587 int32
	_ = v5587
	var v5589 int32
	_ = v5589
	var v5598 int32
	_ = v5598
	var v5599 int32
	_ = v5599
	var v5602 int32
	_ = v5602
	var v5603 int32
	_ = v5603
	var v5607 int32
	_ = v5607
	var v5613 int32
	_ = v5613
	var v5615 int32
	_ = v5615
	var v5616 int32
	_ = v5616
	var v5621 int32
	_ = v5621
	var v5622 int32
	_ = v5622
	var v5626 int32
	_ = v5626
	var v5632 int32
	_ = v5632
	var v5634 int32
	_ = v5634
	var v5640 int32
	_ = v5640
	var v5642 int32
	_ = v5642
	var v5644 int32
	_ = v5644
	var v5645 int32
	_ = v5645
	var v5652 int32
	_ = v5652
	var v5653 int32
	_ = v5653
	var v5654 int32
	_ = v5654
	var v5656 int32
	_ = v5656
	var v5658 int32
	_ = v5658
	var v5659 int32
	_ = v5659
	var v5662 int32
	_ = v5662
	var v5666 int32
	_ = v5666
	var v5667 int32
	_ = v5667
	var v5674 int32
	_ = v5674
	var v5680 int32
	_ = v5680
	var v5682 int32
	_ = v5682
	var v5693 int32
	_ = v5693
	var v5701 int32
	_ = v5701
	var v5707 int32
	_ = v5707
	var v5709 int32
	_ = v5709
	var v5716 int32
	_ = v5716
	var v5718 int32
	_ = v5718
	var v5726 int32
	_ = v5726
	var v5728 int32
	_ = v5728
	var v5729 int32
	_ = v5729
	var v5733 int32
	_ = v5733
	var v5734 int32
	_ = v5734
	var v5736 int32
	_ = v5736
	var v5740 int32
	_ = v5740
	var v5741 int32
	_ = v5741
	var v5742 int32
	_ = v5742
	var v5743 int32
	_ = v5743
	var v5745 int32
	_ = v5745
	var v5748 int32
	_ = v5748
	var v5749 int32
	_ = v5749
	var v5750 int32
	_ = v5750
	var v5754 int32
	_ = v5754
	var v5757 int32
	_ = v5757
	var v5758 int32
	_ = v5758
	var v5759 int32
	_ = v5759
	var v5761 int32
	_ = v5761
	var v5765 int32
	_ = v5765
	var v5769 int32
	_ = v5769
	var v5770 int32
	_ = v5770
	var v5772 int32
	_ = v5772
	var v5773 int32
	_ = v5773
	var v5779 int32
	_ = v5779
	var v5784 int32
	_ = v5784
	var v5785 int32
	_ = v5785
	var v5787 int32
	_ = v5787
	var v5790 int64
	_ = v5790
	var v5791 int32
	_ = v5791
	var v5793 int64
	_ = v5793
	var v5798 int32
	_ = v5798
	var v5800 int32
	_ = v5800
	var v5810 int32
	_ = v5810
	var v5812 int32
	_ = v5812
	var v5813 int32
	_ = v5813
	var v5817 int32
	_ = v5817
	var v5818 int32
	_ = v5818
	var v5819 int32
	_ = v5819
	var v5821 int32
	_ = v5821
	var v5823 int32
	_ = v5823
	var v5825 int32
	_ = v5825
	var v5828 int32
	_ = v5828
	var v5829 int32
	_ = v5829
	var v5833 int32
	_ = v5833
	var v5836 int32
	_ = v5836
	var v5837 int32
	_ = v5837
	var v5838 int32
	_ = v5838
	var v5839 int32
	_ = v5839
	var v5849 int32
	_ = v5849
	var v5902 int32
	_ = v5902
	var v5906 int64
	_ = v5906
	var v5908 int64
	_ = v5908
	var v5911 int32
	_ = v5911
	var v5912 int32
	_ = v5912
	var v5915 int32
	_ = v5915
	var v5916 int32
	_ = v5916
	var v5926 int32
	_ = v5926
	var v5975 int32
	_ = v5975
	var v5976 int64
	_ = v5976
	var v5977 int32
	_ = v5977
	var v5978 int32
	_ = v5978
	var v5979 int32
	_ = v5979
	var v5980 int32
	_ = v5980
	var v5983 int32
	_ = v5983
	var v5988 int32
	_ = v5988
	var v5989 int32
	_ = v5989
	var v5990 int32
	_ = v5990
	var v6009 int32
	_ = v6009
	var v6010 int32
	_ = v6010
	var v6063 int32
	_ = v6063
	var v6064 int32
	_ = v6064
	var v6068 int32
	_ = v6068
	var v6069 int32
	_ = v6069
	var v6073 int32
	_ = v6073
	var v6090 int32
	_ = v6090
	var v6132 int32
	_ = v6132
	var v6133 int64
	_ = v6133
	var v6134 int32
	_ = v6134
	var v6135 int32
	_ = v6135
	var v6136 int32
	_ = v6136
	var v6139 int32
	_ = v6139
	var v6144 int32
	_ = v6144
	var v6148 int32
	_ = v6148
	var v6149 int32
	_ = v6149
	var v6150 int32
	_ = v6150
	var v6151 int32
	_ = v6151
	var v6152 int32
	_ = v6152
	var v6156 int32
	_ = v6156
	var v6159 int32
	_ = v6159
	var v6161 int32
	_ = v6161
	var v6162 int32
	_ = v6162
	var v6164 int32
	_ = v6164
	var v6166 int32
	_ = v6166
	var v6167 int32
	_ = v6167
	var v6172 int32
	_ = v6172
	var v6176 int32
	_ = v6176
	var v6177 int32
	_ = v6177
	var v6179 int32
	_ = v6179
	var v6180 int32
	_ = v6180
	var v6182 int32
	_ = v6182
	var v6187 int32
	_ = v6187
	var v6188 int32
	_ = v6188
	var v6190 int32
	_ = v6190
	var v6191 int32
	_ = v6191
	var v6196 int32
	_ = v6196
	var v6198 int32
	_ = v6198
	var v6199 int32
	_ = v6199
	var v6205 int32
	_ = v6205
	var v6211 int32
	_ = v6211
	var v6213 int32
	_ = v6213
	var v6214 int32
	_ = v6214
	var v6219 int32
	_ = v6219
	var v6220 int32
	_ = v6220
	var v6221 int32
	_ = v6221
	var v6225 int32
	_ = v6225
	var v6231 int32
	_ = v6231
	var v6233 int32
	_ = v6233
	var v6239 int32
	_ = v6239
	var v6240 int32
	_ = v6240
	var v6242 int32
	_ = v6242
	var v6243 int32
	_ = v6243
	var v6246 int32
	_ = v6246
	var v6247 int32
	_ = v6247
	var v6248 int32
	_ = v6248
	var v6249 int32
	_ = v6249
	var v6251 int32
	_ = v6251
	var v6252 int32
	_ = v6252
	var v6256 int32
	_ = v6256
	var v6270 int32
	_ = v6270
	var v6271 int32
	_ = v6271
	var v6321 int32
	_ = v6321
	var v6326 int32
	_ = v6326
	var v6388 int32
	_ = v6388
	var v6394 int32
	_ = v6394
	var v6399 int32
	_ = v6399
	var v6403 int32
	_ = v6403
	var v6405 int32
	_ = v6405
	var v6406 int32
	_ = v6406
	var v6409 int32
	_ = v6409
	var v6410 int32
	_ = v6410
	var v6414 int32
	_ = v6414
	var v6428 int32
	_ = v6428
	var v6429 int32
	_ = v6429
	var v6479 int32
	_ = v6479
	var v6484 int32
	_ = v6484
	var v6546 int32
	_ = v6546
	var v6552 int32
	_ = v6552
	var v6557 int32
	_ = v6557
	var v6562 int32
	_ = v6562
	var v6563 int32
	_ = v6563
	var v6564 int32
	_ = v6564
	var v6568 int32
	_ = v6568
	var v6571 int32
	_ = v6571
	var v6572 int32
	_ = v6572
	var v6573 int32
	_ = v6573
	var v6575 int32
	_ = v6575
	var v6580 int32
	_ = v6580
	var v6581 int32
	_ = v6581
	var v6583 int32
	_ = v6583
	var v6584 int32
	_ = v6584
	var v6586 int32
	_ = v6586
	var v6590 int32
	_ = v6590
	var v6594 int32
	_ = v6594
	var v6595 int32
	_ = v6595
	var v6597 int32
	_ = v6597
	var v6600 int64
	_ = v6600
	var v6601 int32
	_ = v6601
	var v6603 int64
	_ = v6603
	var v6608 int32
	_ = v6608
	var v6614 int32
	_ = v6614
	var v6616 int32
	_ = v6616
	var v6622 int32
	_ = v6622
	var v6624 int32
	_ = v6624
	var v6626 int32
	_ = v6626
	var v6632 int64
	_ = v6632
	var v6633 int32
	_ = v6633
	var v6637 int32
	_ = v6637
	var v6639 int32
	_ = v6639
	var v6643 int32
	_ = v6643
	var v6645 int32
	_ = v6645
	var v6653 int32
	_ = v6653
	var v6655 int32
	_ = v6655
	var v6657 int32
	_ = v6657
	var v6659 int32
	_ = v6659
	var v6665 int32
	_ = v6665
	var v6666 int32
	_ = v6666
	var v6672 int32
	_ = v6672
	var v6677 int32
	_ = v6677
	var v6681 int32
	_ = v6681
	var v6685 int32
	_ = v6685
	var v6690 int32
	_ = v6690
	var v6694 int32
	_ = v6694
	var v6698 int32
	_ = v6698
	var v6703 int32
	_ = v6703
	var v6707 int32
	_ = v6707
	var v6708 int32
	_ = v6708
	var v6714 int32
	_ = v6714
	var v6719 int32
	_ = v6719
	var v6723 int32
	_ = v6723
	var v6727 int32
	_ = v6727
	var v6732 int32
	_ = v6732
	var v6736 int32
	_ = v6736
	var v6740 int32
	_ = v6740
	var v6745 int32
	_ = v6745
	var v6749 int32
	_ = v6749
	var v6750 int32
	_ = v6750
	var v6758 int32
	_ = v6758
	var v6763 int32
	_ = v6763
	var v6767 int32
	_ = v6767
	var v6768 int32
	_ = v6768
	var v6774 int32
	_ = v6774
	var v6779 int32
	_ = v6779
	var v6783 int32
	_ = v6783
	var v6784 int32
	_ = v6784
	var v6790 int32
	_ = v6790
	var v6795 int32
	_ = v6795
	var v6799 int32
	_ = v6799
	var v6803 int32
	_ = v6803
	var v6808 int32
	_ = v6808
	var v6812 int32
	_ = v6812
	var v6813 int32
	_ = v6813
	var v6819 int32
	_ = v6819
	var v6824 int32
	_ = v6824
	var v6825 int32
	_ = v6825
	var v6826 int32
	_ = v6826
	var v6828 int32
	_ = v6828
	var v6831 int32
	_ = v6831
	var v6845 int32
	_ = v6845
	var v6891 int32
	_ = v6891
	var v6896 int32
	_ = v6896
	var v6899 int32
	_ = v6899
	var v6900 int32
	_ = v6900
	var v6909 int32
	_ = v6909
	var v6913 int32
	_ = v6913
	var v6918 int32
	_ = v6918
	var v6979 int32
	_ = v6979
	var v6985 int32
	_ = v6985
	var v6990 int32
	_ = v6990
	var v7056 int32
	_ = v7056
	var v7060 int32
	_ = v7060
	var v7110 int32
	_ = v7110
	var v7112 int32
	_ = v7112
	var v7120 int32
	_ = v7120
	var v7124 int32
	_ = v7124
	var v7130 int32
	_ = v7130
	var v7170 int32
	_ = v7170
	var v7177 int32
	_ = v7177
	var v7179 int32
	_ = v7179
	var v7181 int32
	_ = v7181
	var v7185 int32
	_ = v7185
	var v7193 int32
	_ = v7193
	var v7308 int32
	_ = v7308
	v6 = int32(0)
	v58 = m.G0
	v60 = v58 - int32(1104)
	m.G0 = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v64 != 0 {
		v105 = v6
		v107 = int64(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v60)+464)) = v107
	v109 = int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v109 < v110 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v67 = F_index_getprocinfo(m, l0, int32(1), int32(2))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+6)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v73+v75*int32(0)<<(uint(int32(2))%32)+int32(24)-int32(4))))
	goto L5
L5:
	;
	if v87 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v90 = F_index_getprocinfo(m, l0, int32(1), int32(6))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v97 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+24)))
	if v98 != int32(_a_F_spgdoinsert_0) {
		v105 = v67
		v107 = v97
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	v95 = F_FunctionCall1Coll(m, v90, v93, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v105 = v67
	v107 = v95
	goto L1
L11:
	;
	v102 = F_pg_detoast_datum(m, base.I32_wrap_i64(v97))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v105 = v67
	v107 = base.I64_extend_i32_u(v102)
	goto L1
L13:
	;
	v120 = v109
	v122 = v110
	goto L16
L14:
	;
	goto L15
L15:
	;
	v267 = F_SpGistGetLeafTupleSize(m, v62, v60+int32(464), l4)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L3
	} else {
		goto L34
	}
L16:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v120))))
	if v171 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L15
L18:
	;
	v206 = v120 + int32(1)
	if v206 < v203 {
		v120 = v206
		v122 = v203
		goto L16
	} else {
		goto L26
	}
L19:
	;
	v175 = v120 << (uint(int32(3)) % 32)
	v177 = *(*int64)(unsafe.Add(mBase, uint32(l3+v175)))
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175+v62)+30)))
	if v179 == int32(_a_F_spgdoinsert_0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v60+int32(464)+v120<<(uint(int32(3))%32)))) = int64(0)
	v203 = v122
	goto L18
L22:
	;
	v186 = F_pg_detoast_datum(m, base.I32_wrap_i64(v177))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L3
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v60+int32(464)+v175))) = v177
	v203 = v122
	goto L18
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v60+int32(464)+v175))) = base.I64_extend_i32_u(v186)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v203 = v190
	goto L18
L26:
	;
	goto L17
L27:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_1), int32(123), int32(_a_F_spgdoinsert_2))
	mBase = m.M
	v7308 = m.ExcPending
	if v7308 != 0 {
		goto L3
	} else {
		goto L1033
	}
L28:
	;
	m.G0 = v60 + int32(1104)
	return v7193
L29:
	;
	v7170 = int32(0)
	if base.B2i32(v7130 == v7170)|base.B2i32(v7124 == v7130) == v7170 {
		goto L1026
	} else {
		goto L1027
	}
L30:
	;
	if v7060 == int32(0) {
		goto L1021
	} else {
		goto L1022
	}
L31:
	;
	v7056 = int32(1)
	v7060 = v431
	goto L30
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6979 = m.ExcPending
	if v6979 != 0 {
		goto L3
	} else {
		goto L1018
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6896 = m.ExcPending
	if v6896 != 0 {
		goto L3
	} else {
		goto L1013
	}
L34:
	;
	v270 = v267 + int32(4)
	if base.Ui32(int32(_a_F_spgdoinsert_3)) < base.Ui32(v270) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v64 != 0 {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+460)) = int32(-1)
	v278 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+456)) = uint16(v278)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+452)) = v278
	*(*int64)(unsafe.Add(mBase, uint32(v60)+444)) = int64(4294967295)
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v285 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v273 == int32(0) {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L3
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v288 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+443)) = uint8(v288)
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v291 == v288 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	if v64 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v6845 = v6
	goto L46
L46:
	;
	v6891 = int32(0)
	v7120 = v6891
	v7124 = v6891
	v7130 = v6845
	goto L29
L47:
	;
	v296 = int32(2)
	goto L49
L48:
	;
	v296 = int32(1)
	goto L49
L49:
	;
	if v64 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v299 = int32(8)
	goto L52
L51:
	;
	v299 = int32(0)
	goto L52
L52:
	;
	if v64 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v302 = int32(12)
	goto L55
L54:
	;
	v302 = int32(4)
	goto L55
L55:
	;
	if v64 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v305 = int32(4)
	goto L58
L57:
	;
	v305 = int32(0)
	goto L58
L58:
	;
	if v64 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v308 = int32(7)
	goto L61
L60:
	;
	v308 = int32(3)
	goto L61
L61:
	;
	v310 = v60 + int32(788)
	v327 = int32(-1)
	v336 = v270
	v337 = v296
	v338 = v6
	v339 = int32(1)
	v340 = v6
	v341 = v6
	v344 = v6
	v348 = v327
	v353 = v327
	v356 = v270
	v371 = v6
	goto L62
L62:
	;
	if v337 == int32(-1) {
		goto L70
	} else {
		goto L71
	}
L63:
	;
	v6845 = v4643
	goto L46
L64:
	;
	v4631 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v4631 != 0 {
		v7056 = int32(0)
		v7060 = v4583
		goto L30
	} else {
		goto L659
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+1092)) = v1839
	if v64 != 0 {
		v1909 = int64(0)
		goto L308
	} else {
		goto L309
	}
L66:
	;
	v1820 = int32(0)
	v1839 = v1820
	v1843 = v1820
	v1844 = v1820
	goto L65
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L3
	} else {
		goto L305
	}
L68:
	;
	F_ReleaseBuffer(m, v424)
	mBase = m.M
	v1801 = m.ExcPending
	if v1801 != 0 {
		goto L3
	} else {
		goto L303
	}
L69:
	;
	if v431 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L70:
	;
	v389 = int32(_a_F_spgdoinsert_3)
	if base.Ui32(v389) <= base.Ui32(v356) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	if v341 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L73:
	;
	v392 = v389
	goto L75
L74:
	;
	v392 = v356
	goto L75
L75:
	;
	v395 = F_SpGistGetBuffer(m, l0, v308, v392, v60+int32(443))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	if v395 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v431 = v395
	v432 = v415
	goto L69
L78:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[1]))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v400+(v395^int32(-1))*int32(56))+16))
	v415 = v406
	goto L77
L79:
	;
	goto L80
L80:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[2]))
	v409 = int32(56)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v408+v395*v409-v409)+16))
	v415 = v414
	goto L77
L81:
	;
	v431 = v430
	v432 = v337
	goto L69
L82:
	;
	v418 = F_ReadBuffer(m, l0, v337)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L3
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v337 == v348 {
		v431 = v341
		v432 = v348
		goto L69
	} else {
		goto L87
	}
L85:
	;
	F_LockBufferInternal(m, v418, int32(3))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L3
	} else {
		goto L86
	}
L86:
	;
	v430 = v418
	goto L81
L87:
	;
	v424 = F_ReadBuffer(m, l0, v337)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	v426 = F_ConditionalLockBuffer(m, v424)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L3
	} else {
		goto L89
	}
L89:
	;
	if v426 == int32(0) {
		goto L68
	} else {
		goto L90
	}
L90:
	;
	v430 = v424
	goto L81
L91:
	;
	v451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+16)))
	v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v451+v450))))
	v456 = int32(0)
	if base.B2i32(v453&int32(8) == v456)^v64 == v456 {
		goto L67
	} else {
		goto L95
	}
L92:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v436+(v431^int32(-1))<<(uint(int32(2))%32))))
	v450 = v442
	goto L91
L93:
	;
	goto L94
L94:
	;
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v450 = v444 + v431<<(uint(int32(13))%32) + int32(-8192)
	goto L91
L95:
	;
	if v453&int32(4) == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v4582 = v450
	v4583 = v431
	v4586 = v339
	v4590 = v432
	goto L64
L97:
	;
	goto L98
L98:
	;
	v467 = F_spgFormLeafTuple(m, l1, l2, v60+int32(464), l4)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v474 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+14)))
	v475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+12)))
	v476 = v474 - v475
	v477 = int32(0)
	if v477 < v476 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+16)))
	v485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450+v483)+4)))
	if v485 != 0 {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	v480 = v476
	goto L103
L102:
	;
	v480 = v477
	goto L103
L103:
	;
	goto L100
L104:
	;
	v486 = int32(20)
	goto L106
L105:
	;
	v486 = int32(0)
	goto L106
L106:
	;
	if base.Ui32(int32(base.Ui32(v469)>>(uint(int32(2))%32))+int32(4)) <= base.Ui32(v480+v486) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v489 = int32(_a_F_spgdoinsert_4)
	v491 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	v492 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v491 + v492
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+443)))
	*(*int64)(unsafe.Add(mBase, uint32(v60)+786)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+785)) = uint8(v64)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+784)) = uint8(v495)
	if base.Ui32(v492) < base.Ui32(v432-v492) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	goto L109
L109:
	;
	v681 = v432 - int32(1)
	v683 = base.B2i32(base.Ui32(v681) < base.Ui32(int32(2)))
	if base.Ui32(v681) < base.Ui32(int32(2)) {
		goto L156
	} else {
		goto L157
	}
L110:
	;
	F_MarkBufferDirty(m, v431)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L3
	} else {
		goto L134
	}
L111:
	;
	v507 = v339 & int32(_a_F_spgdoinsert_0)
	goto L113
L112:
	;
	v507 = int32(0)
	goto L113
L113:
	;
	if v507 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v510 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v467)+4)))
	v512 = v510 & int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v467)+4)) = uint16(v512)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v518 = F_SpGistPageAddNewItem(m, v450, v467, int32(base.Ui32(v514)>>(uint(int32(2))%32)), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L3
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v530 = v339 & int32(_a_F_spgdoinsert_0)
	v535 = v450 + v530<<(uint(int32(2))%32) + int32(20)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)))
	v539 = v450 + v536&int32(_a_F_spgdoinsert_6)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	switch v540 & int32(3) {
	case 0:
		goto L121
	default:
		goto L122
	case 2:
		goto L123
	}
L117:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+786)) = uint16(v518)
	if v341 == int32(0) {
		goto L110
	} else {
		goto L118
	}
L118:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+792)) = uint16(v353)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+790)) = uint16(v344)
	F_saveNodeLink(m, v60+int32(444), v432, v518)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L3
	} else {
		goto L119
	}
L119:
	;
	goto L110
L120:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+788)) = uint16(v339)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+786)) = uint16(v618)
	goto L110
L121:
	;
	v592 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v539)+4)))
	v595 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v467)+4)))
	v598 = v592&int32(_a_F_spgdoinsert_7) | v595&int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v467)+4)) = uint16(v598)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v604 = F_SpGistPageAddNewItem(m, v450, v467, int32(base.Ui32(v600)>>(uint(int32(2))%32)), int32(0))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L3
	} else {
		goto L133
	}
L122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L3
	} else {
		goto L130
	}
L123:
	;
	v543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v467)+4)))
	v545 = v543 & int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v467)+4)) = uint16(v545)
	F_PageIndexTupleDelete(m, v450, v530)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L3
	} else {
		goto L124
	}
L124:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v553 = F_PageAddItemExtended(m, v450, v467, int32(base.Ui32(v549)>>(uint(int32(2))%32)), v530, int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	if v553 == v530 {
		v618 = v339
		goto L120
	} else {
		goto L126
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+320)) = int32(base.Ui32(v560) >> (uint(int32(2)) % 32))
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_8), v60+int32(320))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(279), int32(_a_F_spgdoinsert_10))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+304)) = v578 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_11), v60+int32(304))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(286), int32(_a_F_spgdoinsert_10))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L3
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v535)))
	v609 = v450 + v606&int32(_a_F_spgdoinsert_6)
	v610 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v609)+4)))
	v615 = v610&int32(_a_F_spgdoinsert_5) | v604&int32(_a_F_spgdoinsert_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v609)+4)) = uint16(v615)
	v618 = v604
	goto L120
L134:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627)+118)))
	if v628 != int32(112) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v674 = int32(_a_F_spgdoinsert_4)
	v676 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v676 - int32(1)
	goto L31
L136:
	;
	v632 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[6]))
	if v632 <= int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v635 != 0 {
		goto L135
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v637 != 0 {
		goto L135
	} else {
		goto L142
	}
L140:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v636 != 0 {
		goto L135
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L3
	} else {
		goto L143
	}
L143:
	;
	F_XLogRegisterData(m, v60+int32(784), int32(10))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L3
	} else {
		goto L144
	}
L144:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	F_XLogRegisterData(m, v467, int32(base.Ui32(v645)>>(uint(int32(2))%32)))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L3
	} else {
		goto L145
	}
L145:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+784)))
	if v653 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v654 = int32(14)
	goto L148
L147:
	;
	v654 = int32(8)
	goto L148
L148:
	;
	F_XLogRegisterBuffer(m, int32(0), v431, v654)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L3
	} else {
		goto L149
	}
L149:
	;
	v657 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+790)))
	if v657 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	F_XLogRegisterBuffer(m, int32(1), v341, int32(8))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L3
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v662 = int32(16)
	v664 = F_XLogInsert(m, v662, v662)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L3
	} else {
		goto L154
	}
L153:
	;
	goto L152
L154:
	;
	v667 = base.I64_rotl(v664, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v450))) = v667
	v669 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+790)))
	if v669 == int32(0) {
		goto L135
	} else {
		goto L155
	}
L155:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v340))) = v667
	goto L135
L156:
	;
	v1453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+443)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+1100)) = v371
	v1456 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+12)))
	v1462 = int32(base.Ui32(v1456+int32(_a_F_spgdoinsert_12))>>(uint(int32(2))%32)) & int32(_a_F_spgdoinsert_0)
	if base.Ui32(int32(25)) <= base.Ui32(v1456) {
		goto L245
	} else {
		goto L246
	}
L157:
	;
	if v339&int32(_a_F_spgdoinsert_0) == int32(0) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	if base.B2i32(v852 == int32(0))|base.B2i32(base.Ui32(int32(4079)) < base.Ui32(v801)) != 0 {
		goto L156
	} else {
		goto L171
	}
L159:
	;
	v801 = int32(0)
	v852 = int32(1)
	goto L158
L160:
	;
	goto L161
L161:
	;
	v692 = int32(0)
	v700 = v692
	v701 = v339
	v709 = v692
	goto L162
L162:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v450+int32(20)+v701&int32(_a_F_spgdoinsert_0)<<(uint(int32(2))%32))))
	v759 = v450 + v756&int32(_a_F_spgdoinsert_6)
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v759)))
	switch v760 & int32(3) {
	case 0:
		goto L165
	default:
		goto L166
	case 2:
		v788 = v700
		v789 = v709
		goto L164
	}
L163:
	;
	v801 = v788
	v852 = base.B2i32(v789 < int32(64))
	goto L158
L164:
	;
	v790 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v759)+4)))
	v792 = v790 & int32(_a_F_spgdoinsert_7)
	if v792 != 0 {
		v700 = v788
		v701 = v792
		v709 = v789
		goto L162
	} else {
		goto L170
	}
L165:
	;
	v788 = v700 + int32(base.Ui32(v760)>>(uint(int32(2))%32)) + int32(4)
	v789 = v709 + int32(1)
	goto L164
L166:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L3
	} else {
		goto L167
	}
L167:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v759)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+288)) = v767 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_11), v60+int32(288))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L3
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(368), int32(_a_F_spgdoinsert_13))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L3
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	goto L163
L171:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	if base.Ui32(int32(_a_F_spgdoinsert_3)) < base.Ui32(v801+int32(base.Ui32(v858)>>(uint(int32(2))%32))+int32(4)) {
		goto L156
	} else {
		goto L172
	}
L172:
	;
	v866 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+1056)) = uint16(v866)
	v869 = int32(2)
	v870 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v870) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v880 = int32(base.Ui32(v870+int32(_a_F_spgdoinsert_12))>>(uint(v869)%32)) & int32(_a_F_spgdoinsert_0)
	goto L175
L174:
	;
	v880 = v866
	goto L175
L175:
	;
	v881 = F_palloc_mul(m, v869, v880)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L3
	} else {
		goto L176
	}
L176:
	;
	v886 = F_palloc_mul(m, int32(2), v880+int32(1))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L3
	} else {
		goto L177
	}
L177:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v892 = int32(base.Ui32(v888)>>(uint(int32(2))%32)) + int32(4)
	if v339&int32(_a_F_spgdoinsert_0) == int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v1071 = F_SpGistGetBuffer(m, l0, v308, v1016, v60+int32(784)|int32(2))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L3
	} else {
		goto L193
	}
L179:
	;
	v1016 = v892
	v1018 = int32(0)
	v1021 = v866
	goto L178
L180:
	;
	goto L181
L181:
	;
	v907 = v892
	v909 = int32(0)
	v910 = v339
	v912 = v866
	goto L182
L182:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v450+int32(20)+v910&int32(_a_F_spgdoinsert_0)<<(uint(int32(2))%32))))
	v966 = v450 + v963&int32(_a_F_spgdoinsert_6)
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v966)))
	switch v967 & int32(3) {
	case 0:
		goto L185
	default:
		goto L186
	case 2:
		goto L187
	}
L183:
	;
	v1016 = v1003
	v1018 = v1006
	v1021 = v1004
	goto L178
L184:
	;
	v1006 = v909 + int32(1)
	v1007 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v966)+4)))
	v1009 = v1007 & int32(_a_F_spgdoinsert_7)
	if v1009 != 0 {
		v907 = v1003
		v909 = v1006
		v910 = v1009
		v912 = v1004
		goto L182
	} else {
		goto L191
	}
L185:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v881+v909<<(uint(int32(1))%32)))) = uint16(v910)
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v966)))
	v1003 = v907 + int32(base.Ui32(v997)>>(uint(int32(2))%32)) + int32(4)
	v1004 = v912
	goto L184
L186:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L3
	} else {
		goto L188
	}
L187:
	;
	v970 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v881+v909<<(uint(v970)%32)))) = uint16(v910)
	v1003 = v907
	v1004 = v970
	goto L184
L188:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v966)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+272)) = v979 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_11), v60+int32(272))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L3
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(445), int32(_a_F_spgdoinsert_14))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L3
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
	goto L183
L192:
	;
	if v1071 < int32(0) {
		goto L198
	} else {
		goto L199
	}
L193:
	;
	if v1071 < int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1076+(v1071^int32(-1))<<(uint(int32(2))%32))))
	v1090 = v1082
	goto L192
L195:
	;
	goto L196
L196:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v1090 = v1084 + v1071<<(uint(int32(13))%32) + int32(-8192)
	goto L192
L197:
	;
	v1110 = F_palloc(m, v1016)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L3
	} else {
		goto L201
	}
L198:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[1]))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1094+(v1071^int32(-1))*int32(56))+16))
	v1109 = v1100
	goto L197
L199:
	;
	goto L200
L200:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[2]))
	v1103 = int32(56)
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1102+v1071*v1103-v1103)+16))
	v1109 = v1108
	goto L197
L201:
	;
	v1112 = int32(_a_F_spgdoinsert_4)
	v1114 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v1114 + int32(1)
	v1118 = int32(0)
	if base.B2i32(v1018 <= v1118)|v1021 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v1229 = v1110
	v1232 = v1118
	v1280 = int32(0)
	goto L204
L203:
	;
	v1132 = v1110
	v1135 = v1118
	v1136 = int32(0)
	goto L205
L204:
	;
	v1281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v467)+4)))
	v1284 = v1280 | v1281&int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v467)+4)) = uint16(v1284)
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v1294 = F_SpGistPageAddNewItem(m, v1090, v467, int32(base.Ui32(v1289)>>(uint(int32(2))%32)), v60+int32(1056))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L3
	} else {
		goto L212
	}
L205:
	;
	v1184 = v1135 << (uint(int32(1)) % 32)
	v1186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v881+v1184))))
	v1187 = int32(2)
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v450+int32(20)+v1186<<(uint(v1187)%32))))
	v1193 = v450 + v1190&int32(_a_F_spgdoinsert_6)
	v1194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1193)+4)))
	v1199 = v1194&int32(_a_F_spgdoinsert_5) | v1136&int32(_a_F_spgdoinsert_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v1193)+4)) = uint16(v1199)
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1193)))
	v1207 = F_SpGistPageAddNewItem(m, v1090, v1193, int32(base.Ui32(v1202)>>(uint(v1187)%32)), v60+int32(1056))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L3
	} else {
		goto L207
	}
L206:
	;
	v1229 = v1217
	v1232 = v1018
	v1280 = v1207 & int32(_a_F_spgdoinsert_7)
	goto L204
L207:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1184+v886))) = uint16(v1207)
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1193)))
	v1212 = int32(base.Ui32(v1210) >> (uint(int32(2)) % 32))
	if v1212 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	base.MemoryCopy(m, v1132, v1193, v1212)
	goto L210
L209:
	;
	goto L210
L210:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1193)))
	v1217 = v1132 + int32(base.Ui32(v1214)>>(uint(int32(2))%32))
	v1219 = v1135 + int32(1)
	if v1219 != v1018 {
		v1132 = v1217
		v1135 = v1219
		v1136 = v1207
		goto L205
	} else {
		goto L211
	}
L211:
	;
	goto L206
L212:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v886+v1232<<(uint(int32(1))%32)))) = uint16(v1294)
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v1299 = int32(base.Ui32(v1297) >> (uint(int32(2)) % 32))
	if v1299 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	base.MemoryCopy(m, v1229, v467, v1299)
	goto L215
L214:
	;
	goto L215
L215:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v1304 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1305 = int32(3)
	goto L218
L217:
	;
	v1305 = int32(1)
	goto L218
L218:
	;
	F_spgPageIndexMultiDelete(m, l1, v450, v881, v1018, v1305, int32(3), v1109, v1294)
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L3
	} else {
		goto L219
	}
L219:
	;
	F_saveNodeLink(m, v60+int32(444), v1109, v1294)
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L3
	} else {
		goto L220
	}
L220:
	;
	F_MarkBufferDirty(m, v431)
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L3
	} else {
		goto L221
	}
L221:
	;
	F_MarkBufferDirty(m, v1071)
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L3
	} else {
		goto L222
	}
L222:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1317)+118)))
	if v1318 != int32(112) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1386 = int32(_a_F_spgdoinsert_4)
	v1388 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v1388 - int32(1)
	F_SpGistSetLastUsedPage(m, l0, v1071)
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L3
	} else {
		goto L243
	}
L224:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[6]))
	if v1322 <= int32(0) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1325 != 0 {
		goto L223
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v1327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v1327 != 0 {
		goto L223
	} else {
		goto L230
	}
L228:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1326 != 0 {
		goto L223
	} else {
		goto L229
	}
L229:
	;
	goto L227
L230:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v1329 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+800)) = uint8(v1329)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+796)) = v1328
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+792)) = uint16(v353)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+790)) = uint16(v344)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+788)) = uint8(v64)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+787)) = uint8(v1021)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+784)) = uint16(v1018)
	F_XLogBeginInsert(m)
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L3
	} else {
		goto L231
	}
L231:
	;
	F_XLogRegisterData(m, v60+int32(784), int32(20))
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L3
	} else {
		goto L232
	}
L232:
	;
	F_XLogRegisterData(m, v881, v1018<<(uint(int32(1))%32))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L3
	} else {
		goto L233
	}
L233:
	;
	F_XLogRegisterData(m, v886, v1232<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L3
	} else {
		goto L234
	}
L234:
	;
	F_XLogRegisterData(m, v1110, v1229+int32(base.Ui32(v1301)>>(uint(int32(2))%32))-v1110)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L3
	} else {
		goto L235
	}
L235:
	;
	F_XLogRegisterBuffer(m, int32(0), v431, int32(8))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L3
	} else {
		goto L236
	}
L236:
	;
	v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+786)))
	if v1367 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1368 = int32(14)
	goto L239
L238:
	;
	v1368 = int32(8)
	goto L239
L239:
	;
	F_XLogRegisterBuffer(m, int32(1), v1071, v1368)
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L3
	} else {
		goto L240
	}
L240:
	;
	F_XLogRegisterBuffer(m, int32(2), v341, int32(8))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L3
	} else {
		goto L241
	}
L241:
	;
	v1377 = F_XLogInsert(m, int32(16), int32(32))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L3
	} else {
		goto L242
	}
L242:
	;
	v1380 = base.I64_rotl(v1377, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v450))) = v1380
	*(*int64)(unsafe.Add(mBase, uint32(v1090))) = v1380
	*(*int64)(unsafe.Add(mBase, uint32(v340))) = v1380
	goto L223
L243:
	;
	F_UnlockReleaseBuffer(m, v1071)
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L3
	} else {
		goto L244
	}
L244:
	;
	goto L31
L245:
	;
	v1466 = v1462
	goto L247
L246:
	;
	v1466 = int32(0)
	goto L247
L247:
	;
	v1468 = v1466 + int32(1)
	v1469 = F_palloc_mul(m, int32(8), v1468)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L3
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+1096)) = v1469
	v1473 = F_palloc_mul(m, int32(2), v1468)
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L3
	} else {
		goto L249
	}
L249:
	;
	v1476 = F_palloc_mul(m, int32(2), v1468)
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L3
	} else {
		goto L250
	}
L250:
	;
	v1479 = F_palloc_mul(m, int32(4), v1468)
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L3
	} else {
		goto L251
	}
L251:
	;
	v1482 = F_palloc_mul(m, int32(4), v1468)
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L3
	} else {
		goto L252
	}
L252:
	;
	v1485 = F_palloc_mul(m, int32(1), v1468)
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L3
	} else {
		goto L253
	}
L253:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+744)) = v1487
	v1489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+748)) = uint8(v1489)
	if base.Ui32(v681) <= base.Ui32(int32(1)) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	if v1466 == int32(0) {
		goto L66
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	if v339&int32(_a_F_spgdoinsert_0) == int32(0) {
		goto L66
	} else {
		goto L279
	}
L257:
	;
	v1497 = int32(0)
	v1507 = v1497
	v1509 = int32(1)
	v1521 = v1497
	goto L259
L258:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L3
	} else {
		goto L276
	}
L259:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v450+int32(20)+v1509&int32(_a_F_spgdoinsert_0)<<(uint(int32(2))%32))))
	v1565 = v450 + v1562&int32(_a_F_spgdoinsert_6)
	v1566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1565))))
	if v1566&int32(3) != 0 {
		goto L258
	} else {
		goto L261
	}
L260:
	;
	v1839 = v1462
	v1843 = v1462
	v1844 = v1618
	goto L65
L261:
	;
	if v64 != 0 {
		v1597 = int64(0)
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1096))
	*(*int64)(unsafe.Add(mBase, uint32(v1598+v1507<<(uint(int32(3))%32)))) = v1597
	v1603 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1479+v1507<<(uint(v1603)%32)))) = v1565
	v1607 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1473+v1507<<(uint(v1607)%32)))) = uint16(v1509)
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1565)))
	v1618 = v1521 + int32(base.Ui32(v1613)>>(uint(v1603)%32)) + int32(4)
	v1620 = v1507 + v1607
	if v1620 != v1466 {
		v1507 = v1620
		v1509 = v1509 + v1607
		v1521 = v1618
		goto L259
	} else {
		goto L275
	}
L263:
	;
	v1571 = v1565 + int32(16)
	v1572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v1572 == int32(1) {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1575 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+36)))
	if base.I32_popcnt(v1575) != int32(1) {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	goto L266
L266:
	;
	v1597 = base.I64_extend_i32_u(v1571)
	goto L262
L267:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L3
	} else {
		goto L273
	}
L268:
	;
	switch base.I32_ctz(v1575) {
	case 0:
		goto L272
	case 1:
		goto L271
	case 2:
		goto L270
	case 3:
		goto L269
	default:
		goto L267
	}
L269:
	;
	v1583 = *(*int64)(unsafe.Add(mBase, uint32(v1571)))
	v1597 = v1583
	goto L262
L270:
	;
	v1582 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1571))))
	v1597 = v1582
	goto L262
L271:
	;
	v1581 = int64(*(*int16)(unsafe.Add(mBase, uint32(v1571))))
	v1597 = v1581
	goto L262
L272:
	;
	v1580 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1571))))
	v1597 = v1580
	goto L262
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+208)) = v1575
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_15), v60+int32(208))
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L3
	} else {
		goto L274
	}
L274:
	;
	goto L27
L275:
	;
	goto L260
L276:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1565)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+224)) = v1626 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_11), v60+int32(224))
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		goto L3
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(763), int32(_a_F_spgdoinsert_16))
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		goto L3
	} else {
		goto L278
	}
L278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L279:
	;
	v1646 = int32(0)
	v1658 = v339
	v1665 = v1646
	v1669 = v1646
	v1670 = v1646
	goto L280
L280:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v450+int32(20)+v1658&int32(_a_F_spgdoinsert_0)<<(uint(int32(2))%32))))
	v1714 = v450 + v1711&int32(_a_F_spgdoinsert_6)
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1714)))
	switch v1715 & int32(3) {
	case 0:
		goto L285
	default:
		goto L284
	case 2:
		goto L283
	}
L281:
	;
	v1839 = v1792
	v1843 = v1796
	v1844 = v1793
	goto L65
L282:
	;
	v1796 = v1669 + int32(1)
	v1797 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1714)+4)))
	v1799 = v1797 & int32(_a_F_spgdoinsert_7)
	if v1799 != 0 {
		v1658 = v1799
		v1665 = v1792
		v1669 = v1796
		v1670 = v1793
		goto L280
	} else {
		goto L302
	}
L283:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1473+v1669<<(uint(int32(1))%32)))) = uint16(v1658)
	v1792 = v1665
	v1793 = v1670
	goto L282
L284:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L3
	} else {
		goto L299
	}
L285:
	;
	if v64 != 0 {
		v1746 = int64(0)
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1096))
	*(*int64)(unsafe.Add(mBase, uint32(v1747+v1665<<(uint(int32(3))%32)))) = v1746
	v1752 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1479+v1665<<(uint(v1752)%32)))) = v1714
	v1756 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1473+v1669<<(uint(v1756)%32)))) = uint16(v1658)
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1714)))
	v1792 = v1665 + v1756
	v1793 = v1670 + int32(base.Ui32(v1762)>>(uint(v1752)%32)) - int32(16)
	goto L282
L287:
	;
	v1720 = v1714 + int32(16)
	v1721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v1721 == int32(1) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1724 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+36)))
	if base.I32_popcnt(v1724) != int32(1) {
		goto L291
	} else {
		goto L292
	}
L289:
	;
	goto L290
L290:
	;
	v1746 = base.I64_extend_i32_u(v1720)
	goto L286
L291:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L3
	} else {
		goto L297
	}
L292:
	;
	switch base.I32_ctz(v1724) {
	case 0:
		goto L296
	case 1:
		goto L295
	case 2:
		goto L294
	case 3:
		goto L293
	default:
		goto L291
	}
L293:
	;
	v1732 = *(*int64)(unsafe.Add(mBase, uint32(v1720)))
	v1746 = v1732
	goto L286
L294:
	;
	v1731 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1720))))
	v1746 = v1731
	goto L286
L295:
	;
	v1730 = int64(*(*int16)(unsafe.Add(mBase, uint32(v1720))))
	v1746 = v1730
	goto L286
L296:
	;
	v1729 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1720))))
	v1746 = v1729
	goto L286
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+256)) = v1724
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_15), v60+int32(256))
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L3
	} else {
		goto L298
	}
L298:
	;
	goto L27
L299:
	;
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1714)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+240)) = v1772 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_11), v60+int32(240))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L3
	} else {
		goto L300
	}
L300:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(799), int32(_a_F_spgdoinsert_16))
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L3
	} else {
		goto L301
	}
L301:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L302:
	;
	goto L281
L303:
	;
	F_UnlockReleaseBuffer(m, v341)
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L3
	} else {
		goto L304
	}
L304:
	;
	v7193 = int32(0)
	goto L28
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+336)) = v432
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_17), v60+int32(336))
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L3
	} else {
		goto L306
	}
L306:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(2100), int32(_a_F_spgdoinsert_18))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L3
	} else {
		goto L307
	}
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L308:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1096))
	*(*int64)(unsafe.Add(mBase, uint32(v1910+v1839<<(uint(int32(3))%32)))) = v1909
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1092))
	*(*int32)(unsafe.Add(mBase, uint32(v1479+v1915<<(uint(int32(2))%32)))) = v467
	v1920 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+1056)) = v1920
	v1923 = v1915 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+1092)) = v1923
	*(*int64)(unsafe.Add(mBase, uint32(v60)+1064)) = v1920
	*(*int64)(unsafe.Add(mBase, uint32(v60)+1072)) = v1920
	*(*int64)(unsafe.Add(mBase, uint32(v60)+1080)) = v1920
	if v64 == int32(0) {
		goto L323
	} else {
		goto L324
	}
L309:
	;
	v1883 = v467 + int32(16)
	v1884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v1884 == int32(1) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1887 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+36)))
	if base.I32_popcnt(v1887) != int32(1) {
		goto L313
	} else {
		goto L314
	}
L311:
	;
	goto L312
L312:
	;
	v1909 = base.I64_extend_i32_u(v1883)
	goto L308
L313:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L3
	} else {
		goto L319
	}
L314:
	;
	switch base.I32_ctz(v1887) {
	case 0:
		goto L318
	case 1:
		goto L317
	case 2:
		goto L316
	case 3:
		goto L315
	default:
		goto L313
	}
L315:
	;
	v1895 = *(*int64)(unsafe.Add(mBase, uint32(v1883)))
	v1909 = v1895
	goto L308
L316:
	;
	v1894 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1883))))
	v1909 = v1894
	goto L308
L317:
	;
	v1893 = int64(*(*int16)(unsafe.Add(mBase, uint32(v1883))))
	v1909 = v1893
	goto L308
L318:
	;
	v1892 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1883))))
	v1909 = v1892
	goto L308
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+192)) = v1887
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_15), v60+int32(192))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L3
	} else {
		goto L320
	}
L320:
	;
	goto L27
L321:
	;
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	v2684 = F_palloc_mul(m, int32(4), v2683)
	mBase = m.M
	v2685 = m.ExcPending
	if v2685 != 0 {
		goto L3
	} else {
		goto L375
	}
L322:
	;
	if v2206 < int32(2) {
		v2631 = v2206
		v2635 = v2210
		v2640 = v2215
		v2656 = v2231
		goto L321
	} else {
		goto L353
	}
L323:
	;
	v1933 = int32(1)
	v1936 = F_index_getprocinfo(m, l0, v1933, int32(3))
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L3
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v2070 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+1072)) = v2070
	v2074 = F_palloc0_mul(m, int32(4), v1923)
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L3
	} else {
		goto L340
	}
L326:
	;
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1938)))
	v1940 = F_FunctionCall2Coll(m, v1936, v1939, base.I64_extend_i32_u(v60+int32(1092)), base.I64_extend_i32_u(v60+int32(1056)))
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L3
	} else {
		goto L327
	}
L327:
	;
	v1942 = int32(0)
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1092))
	if v1945 <= v1942 {
		v2631 = v1945
		v2635 = v1942
		v2640 = v1942
		v2656 = v1933
		goto L321
	} else {
		goto L328
	}
L328:
	;
	v1955 = v1942
	v1963 = v1942
	goto L329
L329:
	;
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(v2005)))
	if int32(2) <= v2006 {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	v2206 = v2068
	v2210 = v1942
	v2215 = v2065
	v2231 = v1933
	goto L322
L331:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v1479+v1955<<(uint(int32(2))%32))))
	v2019 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2012)+4)))
	goto L335
L332:
	;
	goto L333
L333:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1084))
	v2042 = *(*int64)(unsafe.Add(mBase, uint32(v2038+v1955<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v60)+784)) = v2042
	v2044 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+752)) = uint8(v2044)
	v2047 = v1955 << (uint(int32(2)) % 32)
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(v2047+v1479)))
	v2057 = F_spgFormLeafTuple(m, l1, v2050+int32(6), v60+int32(784), v60+int32(752))
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L3
	} else {
		goto L338
	}
L334:
	;
	goto L333
L335:
	;
	F_index_deform_tuple_internal(m, v2005, v60+int32(784), v60+int32(752), v2012+int32(16), v2012+int32(12), int32(base.Ui32(v2019&int32(_a_F_spgdoinsert_19))>>(uint(int32(15))%32)))
	mBase = m.M
	goto L334
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1482+v2047))) = v2057
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(v2057)))
	v2065 = v1963 + int32(base.Ui32(v2060)>>(uint(int32(2))%32)) + int32(4)
	v2067 = v1955 + int32(1)
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1092))
	if v2067 < v2068 {
		v1955 = v2067
		v1963 = v2065
		goto L329
	} else {
		goto L339
	}
L339:
	;
	goto L330
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+1080)) = v2074
	v2077 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(v1915) {
		v2631 = v1923
		v2635 = v2077
		v2640 = v2077
		v2656 = v2070
		goto L321
	} else {
		goto L341
	}
L341:
	;
	v2089 = v2077
	v2097 = v2077
	goto L342
L342:
	;
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v2139)))
	if int32(2) <= v2140 {
		goto L344
	} else {
		goto L345
	}
L343:
	;
	v2206 = v2198
	v2210 = v2077
	v2215 = v2195
	v2231 = v2070
	goto L322
L344:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v1479+v2089<<(uint(int32(2))%32))))
	v2148 = v60 + int32(784)
	v2150 = v60 + int32(752)
	v2153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2146)+4)))
	goto L349
L345:
	;
	goto L346
L346:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v60)+784)) = int64(0)
	v2174 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+752)) = uint8(v2174)
	v2177 = v2089 << (uint(int32(2)) % 32)
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v2177+v1479)))
	v2187 = F_spgFormLeafTuple(m, l1, v2180+int32(6), v60+int32(784), v60+int32(752))
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L3
	} else {
		goto L351
	}
L347:
	;
	goto L346
L348:
	;
	F_index_deform_tuple_internal(m, v2139, v2148, v2150, v2146+int32(16), v2146+int32(12), int32(base.Ui32(v2153&int32(_a_F_spgdoinsert_19))>>(uint(int32(15))%32)))
	mBase = m.M
	goto L347
L349:
	;
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v2139)))
	if v2156 != int32(1) {
		goto L348
	} else {
		goto L350
	}
L350:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2148))) = int64(0)
	v2161 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2150))) = uint8(v2161)
	goto L347
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1482+v2177))) = v2187
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v2187)))
	v2195 = v2097 + int32(base.Ui32(v2190)>>(uint(int32(2))%32)) + int32(4)
	v2197 = v2089 + int32(1)
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1092))
	if v2197 < v2198 {
		v2089 = v2197
		v2097 = v2195
		goto L342
	} else {
		goto L352
	}
L352:
	;
	goto L343
L353:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1080))
	v2260 = *(*int32)(unsafe.Add(mBase, uint32(v2259)))
	v2261 = int32(1)
	v2264 = v2206 - base.B2i32(base.Ui32(int32(_a_F_spgdoinsert_3)) < base.Ui32(v2215))
	if base.Ui32(v2261) < base.Ui32(v2264) {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v2274 = v2261
	goto L357
L355:
	;
	goto L356
L356:
	;
	if base.Ui32(int32(_a_F_spgdoinsert_20)) <= base.Ui32(v2215) {
		goto L361
	} else {
		goto L362
	}
L357:
	;
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(v2259+v2274<<(uint(int32(2))%32))))
	if v2327 != v2260 {
		v2631 = v2206
		v2635 = v2210
		v2640 = v2215
		v2656 = v2231
		goto L321
	} else {
		goto L359
	}
L358:
	;
	goto L356
L359:
	;
	v2330 = v2274 + int32(1)
	if v2330 != v2264 {
		v2274 = v2330
		goto L357
	} else {
		goto L360
	}
L360:
	;
	goto L358
L361:
	;
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v2259+v2206<<(uint(int32(2))%32)-int32(4))))
	v2399 = base.B2i32(v2397 == v2260)
	goto L363
L362:
	;
	v2399 = int32(1)
	goto L363
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+1072)) = int32(8)
	v2410 = int32(0)
	goto L364
L364:
	;
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1080))
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	v2465 = base.I32_rem_s(v2410, v2464)
	*(*int32)(unsafe.Add(mBase, uint32(v2460+v2410<<(uint(int32(2))%32)))) = v2465
	v2468 = v2410 + int32(1)
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1092))
	if v2468 < v2469 {
		v2410 = v2468
		goto L364
	} else {
		goto L366
	}
L365:
	;
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1076))
	if v2471 == int32(0) {
		goto L367
	} else {
		goto L368
	}
L366:
	;
	goto L365
L367:
	;
	v2610 = int32(4)
	v2611 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1092))
	if v2399 != 0 {
		v2631 = v2611
		v2635 = v2610
		v2640 = v2215
		v2656 = v2231
		goto L321
	} else {
		goto L374
	}
L368:
	;
	v2477 = *(*int64)(unsafe.Add(mBase, uint32(v2471+v2260<<(uint(int32(3))%32))))
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	v2480 = F_palloc_mul(m, int32(8), v2479)
	mBase = m.M
	v2481 = m.ExcPending
	if v2481 != 0 {
		goto L3
	} else {
		goto L369
	}
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+1076)) = v2480
	v2483 = int32(0)
	v2484 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	if v2484 <= v2483 {
		goto L367
	} else {
		goto L370
	}
L370:
	;
	v2494 = v2483
	goto L371
L371:
	;
	v2544 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1076))
	*(*int64)(unsafe.Add(mBase, uint32(v2544+v2494<<(uint(int32(3))%32)))) = v2477
	v2550 = v2494 + int32(1)
	v2551 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	if v2550 < v2551 {
		v2494 = v2550
		goto L371
	} else {
		goto L373
	}
L372:
	;
	goto L367
L373:
	;
	goto L372
L374:
	;
	v2613 = v2611 - int32(1)
	v2614 = int32(2)
	v2617 = *(*int32)(unsafe.Add(mBase, uint32(v1482+v2613<<(uint(v2614)%32))))
	v2618 = *(*int32)(unsafe.Add(mBase, uint32(v2617)))
	v2631 = v2613
	v2635 = v2610
	v2640 = v2215 - int32(base.Ui32(v2618)>>(uint(v2614)%32)) - int32(4)
	v2656 = int32(0)
	goto L321
L375:
	;
	v2687 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	v2688 = F_palloc0_mul(m, int32(4), v2687)
	mBase = m.M
	v2689 = m.ExcPending
	if v2689 != 0 {
		goto L3
	} else {
		goto L376
	}
L376:
	;
	v2690 = int32(0)
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	if v2690 < v2691 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v2701 = v2690
	goto L380
L378:
	;
	v2778 = v2691
	goto L379
L379:
	;
	v2827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1056)))
	v2828 = *(*int64)(unsafe.Add(mBase, uint32(v60)+1064))
	v2829 = F_spgFormInnerTuple(m, l1, v2827, v2828, v2778, v2684)
	mBase = m.M
	v2830 = m.ExcPending
	if v2830 != 0 {
		goto L3
	} else {
		goto L387
	}
L380:
	;
	v2754 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1076))
	if v2754 != 0 {
		goto L382
	} else {
		goto L383
	}
L381:
	;
	v2778 = v2768
	goto L379
L382:
	;
	v2758 = *(*int64)(unsafe.Add(mBase, uint32(v2754+v2701<<(uint(int32(3))%32))))
	v2760 = v2758
	goto L384
L383:
	;
	v2760 = int64(0)
	goto L384
L384:
	;
	v2763 = F_spgFormNodeTuple(m, l1, v2760, base.B2i32(v2754 == int32(0)))
	mBase = m.M
	v2764 = m.ExcPending
	if v2764 != 0 {
		goto L3
	} else {
		goto L385
	}
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2684+v2701<<(uint(int32(2))%32)))) = v2763
	v2767 = v2701 + int32(1)
	v2768 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	if v2767 < v2768 {
		v2701 = v2767
		goto L380
	} else {
		goto L386
	}
L386:
	;
	goto L381
L387:
	;
	v2831 = *(*int32)(unsafe.Add(mBase, uint32(v2829)))
	*(*int32)(unsafe.Add(mBase, uint32(v2829))) = v2831&int32(-5) | v2635
	if v2831&int32(_a_F_spgdoinsert_21) != 0 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v2851 = v2829 + int32(base.Ui32(v2831)>>(uint(int32(16))%32)) + int32(8)
	v2852 = int32(0)
	goto L391
L389:
	;
	goto L390
L390:
	;
	v2974 = int32(0)
	if v2974 < v2631 {
		goto L397
	} else {
		goto L398
	}
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2684+v2852<<(uint(int32(2))%32)))) = v2851
	v2905 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2851)+6)))
	v2906 = int32(_a_F_spgdoinsert_22)
	v2910 = v2852 + int32(1)
	v2911 = *(*int32)(unsafe.Add(mBase, uint32(v2829)))
	if base.Ui32(v2910) < base.Ui32(int32(base.Ui32(v2911)>>(uint(int32(3))%32))&v2906) {
		v2851 = v2851 + v2905&v2906
		v2852 = v2910
		goto L391
	} else {
		goto L393
	}
L392:
	;
	goto L390
L393:
	;
	goto L392
L394:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+735)) = uint8(v64)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+730)) = uint8(v1453)
	v3952 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+726)) = uint16(v3952)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+724)) = uint8(v683)
	v3956 = F_palloc(m, v2640)
	mBase = m.M
	v3957 = m.ExcPending
	if v3957 != 0 {
		goto L3
	} else {
		goto L502
	}
L395:
	;
	if v3611 <= int32(0) {
		v3901 = v3193
		v3909 = v3611
		v3924 = v3626
		goto L394
	} else {
		goto L491
	}
L396:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3585 = m.ExcPending
	if v3585 != 0 {
		goto L3
	} else {
		goto L488
	}
L397:
	;
	v2984 = v2974
	goto L400
L398:
	;
	goto L399
L399:
	;
	v3116 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+734)) = uint8(v3116)
	if v341 == v3116 {
		v3157 = v3116
		goto L405
	} else {
		goto L406
	}
L400:
	;
	v3035 = v2984 << (uint(int32(2)) % 32)
	v3036 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1080))
	v3038 = *(*int32)(unsafe.Add(mBase, uint32(v3035+v3036)))
	if v3038 < int32(0) {
		goto L396
	} else {
		goto L402
	}
L401:
	;
	goto L399
L402:
	;
	v3041 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	if v3041 <= v3038 {
		goto L396
	} else {
		goto L403
	}
L403:
	;
	v3043 = int32(2)
	v3045 = v2688 + v3038<<(uint(v3043)%32)
	v3046 = *(*int32)(unsafe.Add(mBase, uint32(v3045)))
	v3048 = *(*int32)(unsafe.Add(mBase, uint32(v3035+v1482)))
	v3049 = *(*int32)(unsafe.Add(mBase, uint32(v3048)))
	*(*int32)(unsafe.Add(mBase, uint32(v3045))) = v3046 + int32(base.Ui32(v3049)>>(uint(v3043)%32)) + int32(4)
	v3057 = v2984 + int32(1)
	if v3057 != v2631 {
		v2984 = v3057
		goto L400
	} else {
		goto L404
	}
L404:
	;
	goto L401
L405:
	;
	if v683 == int32(0) {
		goto L420
	} else {
		goto L421
	}
L406:
	;
	v3122 = int32(1)
	if base.Ui32(v348-v3122) <= base.Ui32(v3122) {
		goto L408
	} else {
		goto L409
	}
L407:
	;
	v3152 = base.I32_rem_u_s(v348+int32(1), int32(3))
	v3154 = F_SpGistGetBuffer(m, l0, v3152|v305, v3147, v60+int32(734))
	mBase = m.M
	v3155 = m.ExcPending
	if v3155 != 0 {
		goto L3
	} else {
		goto L419
	}
L408:
	;
	v3126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2829)+4)))
	v3147 = v3126 + int32(4)
	goto L407
L409:
	;
	goto L410
L410:
	;
	v3129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v340)+14)))
	v3130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v340)+12)))
	v3131 = v3129 - v3130
	v3132 = int32(0)
	if v3132 < v3131 {
		goto L412
	} else {
		goto L413
	}
L411:
	;
	v3138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v340)+16)))
	v3140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v340+v3138)+4)))
	if v3140 != 0 {
		goto L415
	} else {
		goto L416
	}
L412:
	;
	v3135 = v3131
	goto L414
L413:
	;
	v3135 = v3132
	goto L414
L414:
	;
	goto L411
L415:
	;
	v3141 = int32(20)
	goto L417
L416:
	;
	v3141 = int32(0)
	goto L417
L417:
	;
	v3143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2829)+4)))
	v3145 = v3143 + int32(4)
	if base.Ui32(v3145) <= base.Ui32(v3135+v3141) {
		v3157 = v341
		goto L405
	} else {
		goto L418
	}
L418:
	;
	v3147 = v3145
	goto L407
L419:
	;
	v3157 = v3154
	goto L405
L420:
	;
	v3160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+14)))
	v3161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+12)))
	v3162 = v3160 - v3161
	v3163 = int32(0)
	if v3163 < v3162 {
		goto L424
	} else {
		goto L425
	}
L421:
	;
	v3168 = v3116
	goto L422
L422:
	;
	v3169 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+731)) = uint8(v3169)
	if v2640 <= v3168 {
		goto L427
	} else {
		goto L428
	}
L423:
	;
	v3168 = v3166 + v1844
	goto L422
L424:
	;
	v3166 = v3162
	goto L426
L425:
	;
	v3166 = v3163
	goto L426
L426:
	;
	goto L423
L427:
	;
	v3172 = int32(0)
	v3173 = v1839 + v2656
	if base.B2i32(v3173 <= v3172)|base.B2i32(v3173 == v3172) != 0 {
		v3901 = v3172
		v3909 = v3173
		v3924 = v2656
		goto L394
	} else {
		goto L430
	}
L428:
	;
	goto L429
L429:
	;
	v3181 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1092))
	if v3181 != int32(1) {
		goto L431
	} else {
		goto L432
	}
L430:
	;
	base.MemoryFill(m, v1485, int32(0), v3173)
	v3901 = v3172
	v3909 = v3173
	v3924 = v2656
	goto L394
L431:
	;
	v3189 = int32(_a_F_spgdoinsert_3)
	if base.Ui32(v3189) <= base.Ui32(v2640) {
		goto L434
	} else {
		goto L435
	}
L432:
	;
	if base.Ui32(v2640) <= base.Ui32(int32(_a_F_spgdoinsert_3)) {
		goto L431
	} else {
		goto L433
	}
L433:
	;
	v3901 = int32(0)
	v3909 = v1839
	v3924 = int32(0)
	goto L394
L434:
	;
	v3192 = v3189
	goto L436
L435:
	;
	v3192 = v2640
	goto L436
L436:
	;
	v3193 = F_SpGistGetBuffer(m, l0, v308, v3192, v60+int32(731))
	mBase = m.M
	v3194 = m.ExcPending
	if v3194 != 0 {
		goto L3
	} else {
		goto L437
	}
L437:
	;
	v3196 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	v3197 = F_palloc_mul(m, int32(1), v3196)
	mBase = m.M
	v3198 = m.ExcPending
	if v3198 != 0 {
		goto L3
	} else {
		goto L438
	}
L438:
	;
	v3199 = int32(0)
	v3200 = base.B2i32(v3199 <= v3193)
	if v3200 == v3199 {
		goto L440
	} else {
		goto L441
	}
L439:
	;
	v3219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3218)+14)))
	v3220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3218)+12)))
	v3221 = v3219 - v3220
	v3222 = int32(0)
	if v3222 < v3221 {
		goto L444
	} else {
		goto L445
	}
L440:
	;
	v3204 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v3210 = *(*int32)(unsafe.Add(mBase, uint32(v3204+(v3193^int32(-1))<<(uint(int32(2))%32))))
	v3218 = v3210
	goto L439
L441:
	;
	goto L442
L442:
	;
	v3212 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v3218 = v3212 + v3193<<(uint(int32(13))%32) + int32(-8192)
	goto L439
L443:
	;
	v3226 = int32(0)
	v3227 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	if v3226 < v3227 {
		goto L447
	} else {
		goto L448
	}
L444:
	;
	v3225 = v3221
	goto L446
L445:
	;
	v3225 = v3222
	goto L446
L446:
	;
	goto L443
L447:
	;
	v3236 = v3168
	v3237 = v3226
	v3244 = v3225
	goto L450
L448:
	;
	v3313 = v3168
	v3321 = v3225
	goto L449
L449:
	;
	if int32(0) <= v3313|v3321 {
		goto L457
	} else {
		goto L458
	}
L450:
	;
	v3287 = v3237 + v3197
	v3290 = v2688 + v3237<<(uint(int32(2))%32)
	v3291 = *(*int32)(unsafe.Add(mBase, uint32(v3290)))
	if v3291 <= v3236 {
		goto L453
	} else {
		goto L454
	}
L451:
	;
	v3313 = v3301
	v3321 = v3302
	goto L449
L452:
	;
	v3304 = v3237 + int32(1)
	v3305 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	if v3304 < v3305 {
		v3236 = v3301
		v3237 = v3304
		v3244 = v3302
		goto L450
	} else {
		goto L456
	}
L453:
	;
	v3293 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3287))) = uint8(v3293)
	v3295 = *(*int32)(unsafe.Add(mBase, uint32(v3290)))
	v3301 = v3236 - v3295
	v3302 = v3244
	goto L452
L454:
	;
	goto L455
L455:
	;
	v3297 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3287))) = uint8(v3297)
	v3299 = *(*int32)(unsafe.Add(mBase, uint32(v3290)))
	v3301 = v3236
	v3302 = v3244 - v3299
	goto L452
L456:
	;
	goto L451
L457:
	;
	v3611 = v1839 + v2656
	v3626 = v2656
	goto L395
L458:
	;
	goto L459
L459:
	;
	if v2656 != 0 {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v3368 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1092))
	v3369 = int32(2)
	v3371 = int32(4)
	v3372 = v3368<<(uint(v3369)%32) - v3371
	v3373 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1080))
	v3375 = *(*int32)(unsafe.Add(mBase, uint32(v3372+v3373)))
	v3378 = v2688 + v3375<<(uint(v3369)%32)
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v3378)))
	v3381 = *(*int32)(unsafe.Add(mBase, uint32(v3372+v1482)))
	v3382 = *(*int32)(unsafe.Add(mBase, uint32(v3381)))
	*(*int32)(unsafe.Add(mBase, uint32(v3378))) = v3379 - int32(base.Ui32(v3382)>>(uint(v3369)%32)) - v3371
	if v3200 == int32(0) {
		goto L464
	} else {
		goto L465
	}
L461:
	;
	goto L462
L462:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3572 = m.ExcPending
	if v3572 != 0 {
		goto L3
	} else {
		goto L485
	}
L463:
	;
	v3407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3406)+14)))
	v3408 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3406)+12)))
	v3409 = v3407 - v3408
	v3410 = int32(0)
	if v3410 < v3409 {
		goto L468
	} else {
		goto L469
	}
L464:
	;
	v3392 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v3398 = *(*int32)(unsafe.Add(mBase, uint32(v3392+(v3193^int32(-1))<<(uint(int32(2))%32))))
	v3406 = v3398
	goto L463
L465:
	;
	goto L466
L466:
	;
	v3400 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v3406 = v3400 + v3193<<(uint(int32(13))%32) + int32(-8192)
	goto L463
L467:
	;
	v3414 = int32(0)
	v3415 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	if v3414 < v3415 {
		goto L471
	} else {
		goto L472
	}
L468:
	;
	v3413 = v3409
	goto L470
L469:
	;
	v3413 = v3410
	goto L470
L470:
	;
	goto L467
L471:
	;
	v3425 = v3414
	v3428 = v3168
	v3439 = v3413
	goto L474
L472:
	;
	v3505 = v3168
	v3516 = v3413
	goto L473
L473:
	;
	v3552 = int32(0)
	if v3552 <= v3505|v3516 {
		v3611 = v1839
		v3626 = v3552
		goto L395
	} else {
		goto L481
	}
L474:
	;
	v3475 = v3425 + v3197
	v3478 = v2688 + v3425<<(uint(int32(2))%32)
	v3479 = *(*int32)(unsafe.Add(mBase, uint32(v3478)))
	if v3479 <= v3428 {
		goto L477
	} else {
		goto L478
	}
L475:
	;
	v3505 = v3489
	v3516 = v3490
	goto L473
L476:
	;
	v3492 = v3425 + int32(1)
	v3493 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1072))
	if v3492 < v3493 {
		v3425 = v3492
		v3428 = v3489
		v3439 = v3490
		goto L474
	} else {
		goto L480
	}
L477:
	;
	v3481 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3475))) = uint8(v3481)
	v3483 = *(*int32)(unsafe.Add(mBase, uint32(v3478)))
	v3489 = v3428 - v3483
	v3490 = v3439
	goto L476
L478:
	;
	goto L479
L479:
	;
	v3485 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3475))) = uint8(v3485)
	v3487 = *(*int32)(unsafe.Add(mBase, uint32(v3478)))
	v3489 = v3428
	v3490 = v3439 - v3487
	goto L476
L480:
	;
	goto L475
L481:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3559 = m.ExcPending
	if v3559 != 0 {
		goto L3
	} else {
		goto L482
	}
L482:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_23), int32(0))
	mBase = m.M
	v3563 = m.ExcPending
	if v3563 != 0 {
		goto L3
	} else {
		goto L483
	}
L483:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1108), int32(_a_F_spgdoinsert_16))
	mBase = m.M
	v3568 = m.ExcPending
	if v3568 != 0 {
		goto L3
	} else {
		goto L484
	}
L484:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L485:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_23), int32(0))
	mBase = m.M
	v3576 = m.ExcPending
	if v3576 != 0 {
		goto L3
	} else {
		goto L486
	}
L486:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1113), int32(_a_F_spgdoinsert_16))
	mBase = m.M
	v3581 = m.ExcPending
	if v3581 != 0 {
		goto L3
	} else {
		goto L487
	}
L487:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L488:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_24), int32(0))
	mBase = m.M
	v3589 = m.ExcPending
	if v3589 != 0 {
		goto L3
	} else {
		goto L489
	}
L489:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(953), int32(_a_F_spgdoinsert_16))
	mBase = m.M
	v3594 = m.ExcPending
	if v3594 != 0 {
		goto L3
	} else {
		goto L490
	}
L490:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L491:
	;
	v3655 = v3611 & int32(3)
	v3656 = int32(0)
	v3657 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1080))
	if base.Ui32(int32(4)) <= base.Ui32(v3611) {
		goto L492
	} else {
		goto L493
	}
L492:
	;
	v3671 = v3656
	v3685 = int32(0)
	goto L495
L493:
	;
	v3773 = v3656
	goto L494
L494:
	;
	v3830 = v3773
	v3833 = v3656
	goto L499
L495:
	;
	v3722 = int32(2)
	v3725 = *(*int32)(unsafe.Add(mBase, uint32(v3657+v3671<<(uint(v3722)%32))))
	v3727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3197+v3725))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3671+v1485))) = uint8(v3727)
	v3730 = v3671 | int32(1)
	v3735 = *(*int32)(unsafe.Add(mBase, uint32(v3657+v3730<<(uint(v3722)%32))))
	v3737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3197+v3735))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1485+v3730))) = uint8(v3737)
	v3740 = v3671 | v3722
	v3745 = *(*int32)(unsafe.Add(mBase, uint32(v3657+v3740<<(uint(v3722)%32))))
	v3747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3197+v3745))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1485+v3740))) = uint8(v3747)
	v3750 = v3671 | int32(3)
	v3755 = *(*int32)(unsafe.Add(mBase, uint32(v3657+v3750<<(uint(v3722)%32))))
	v3757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3197+v3755))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1485+v3750))) = uint8(v3757)
	v3759 = int32(4)
	v3760 = v3671 + v3759
	v3762 = v3685 + v3759
	if v3762 != v3611&int32(2147483644) {
		v3671 = v3760
		v3685 = v3762
		goto L495
	} else {
		goto L497
	}
L496:
	;
	if v3655 == int32(0) {
		v3901 = v3193
		v3909 = v3611
		v3924 = v3626
		goto L394
	} else {
		goto L498
	}
L497:
	;
	goto L496
L498:
	;
	v3773 = v3760
	goto L494
L499:
	;
	v3884 = *(*int32)(unsafe.Add(mBase, uint32(v3657+v3830<<(uint(int32(2))%32))))
	v3886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3197+v3884))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3830+v1485))) = uint8(v3886)
	v3888 = int32(1)
	v3891 = v3833 + v3888
	if v3891 != v3655 {
		v3830 = v3830 + v3888
		v3833 = v3891
		goto L499
	} else {
		goto L501
	}
L500:
	;
	v3901 = v3193
	v3909 = v3611
	v3924 = v3626
	goto L394
L501:
	;
	goto L500
L502:
	;
	v3958 = int32(_a_F_spgdoinsert_4)
	v3960 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v3960 + int32(1)
	v3964 = int32(0)
	if base.Ui32(v681) < base.Ui32(int32(2)) {
		v4040 = v3964
		goto L503
	} else {
		goto L504
	}
L503:
	;
	v4041 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+1052)) = v4041
	if v4041 < v3909 {
		goto L526
	} else {
		goto L527
	}
L504:
	;
	v3965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v3965 == int32(1) {
		goto L506
	} else {
		goto L507
	}
L505:
	;
	if v1453&int32(1) != 0 {
		v4040 = v3964
		goto L503
	} else {
		goto L524
	}
L506:
	;
	v3968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+16)))
	v3970 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450+v3968)+4)))
	v3972 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v3972) {
		goto L509
	} else {
		goto L510
	}
L507:
	;
	goto L508
L508:
	;
	if v1453&int32(1) != 0 {
		v4040 = v3964
		goto L503
	} else {
		goto L518
	}
L509:
	;
	v3982 = int32(base.Ui32(v3972+int32(_a_F_spgdoinsert_12))>>(uint(int32(2))%32)) & int32(_a_F_spgdoinsert_0)
	goto L511
L510:
	;
	v3982 = int32(0)
	goto L511
L511:
	;
	if v1843+v3970 != v3982 {
		goto L505
	} else {
		goto L512
	}
L512:
	;
	if v431 < int32(0) {
		goto L515
	} else {
		goto L516
	}
L513:
	;
	v4010 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+730)) = uint8(v4010)
	v4040 = v3964
	goto L503
L514:
	;
	F_PageInit(m, v4001, int32(_a_F_spgdoinsert_25), int32(8))
	mBase = m.M
	v4005 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4001)+16)))
	v4006 = v4001 + v4005
	v4007 = int32(_a_F_spgdoinsert_26)
	*(*uint16)(unsafe.Add(mBase, uint32(v4006)+6)) = uint16(v4007)
	*(*uint16)(unsafe.Add(mBase, uint32(v4006))) = uint16(v302)
	goto L513
L515:
	;
	v3987 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v3993 = *(*int32)(unsafe.Add(mBase, uint32(v3987+(v431^int32(-1))<<(uint(int32(2))%32))))
	v4001 = v3993
	goto L514
L516:
	;
	goto L517
L517:
	;
	v3995 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v4001 = v3995 + v431<<(uint(int32(13))%32) + int32(-8192)
	goto L514
L518:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+726)) = uint16(v1843)
	if v1843 <= int32(0) {
		goto L519
	} else {
		goto L520
	}
L519:
	;
	v4017 = int32(1)
	F_spgPageIndexMultiDelete(m, l1, v450, v1473, v1843, v4017, int32(3), int32(0), v4017)
	mBase = m.M
	v4022 = m.ExcPending
	if v4022 != 0 {
		goto L3
	} else {
		goto L522
	}
L520:
	;
	goto L521
L521:
	;
	v4023 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1473))))
	v4024 = int32(1)
	F_spgPageIndexMultiDelete(m, l1, v450, v1473, v1843, v4024, int32(3), int32(0), v4024)
	mBase = m.M
	v4029 = m.ExcPending
	if v4029 != 0 {
		goto L3
	} else {
		goto L523
	}
L522:
	;
	v4040 = v3964
	goto L503
L523:
	;
	v4040 = v4023
	goto L503
L524:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+726)) = uint16(v1843)
	v4033 = int32(3)
	F_spgPageIndexMultiDelete(m, l1, v450, v1473, v1843, v4033, v4033, int32(-1), int32(0))
	mBase = m.M
	v4038 = m.ExcPending
	if v4038 != 0 {
		goto L3
	} else {
		goto L525
	}
L525:
	;
	v4040 = v3964
	goto L503
L526:
	;
	v4052 = v3952
	v4065 = v3956
	goto L529
L527:
	;
	v4225 = v3956
	goto L528
L528:
	;
	if v3901 != 0 {
		goto L551
	} else {
		goto L552
	}
L529:
	;
	v4103 = v4052 << (uint(int32(2)) % 32)
	v4104 = v1482 + v4103
	v4105 = *(*int32)(unsafe.Add(mBase, uint32(v4104)))
	v4106 = v4052 + v1485
	v4107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4106))))
	if v4107 != 0 {
		goto L531
	} else {
		goto L532
	}
L530:
	;
	v4225 = v4201
	goto L528
L531:
	;
	v4108 = v3901
	goto L533
L532:
	;
	v4108 = v431
	goto L533
L533:
	;
	if v4108 < int32(0) {
		goto L535
	} else {
		goto L536
	}
L534:
	;
	v4128 = *(*int32)(unsafe.Add(mBase, uint32(v60)+1080))
	v4130 = *(*int32)(unsafe.Add(mBase, uint32(v4128+v4103)))
	v4133 = v2684 + v4130<<(uint(int32(2))%32)
	v4134 = *(*int32)(unsafe.Add(mBase, uint32(v4133)))
	if v4134 == int32(0) {
		goto L539
	} else {
		goto L540
	}
L535:
	;
	v4112 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[1]))
	v4118 = *(*int32)(unsafe.Add(mBase, uint32(v4112+(v4108^int32(-1))*int32(56))+16))
	v4127 = v4118
	goto L534
L536:
	;
	goto L537
L537:
	;
	v4120 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[2]))
	v4121 = int32(56)
	v4126 = *(*int32)(unsafe.Add(mBase, uint32(v4120+v4108*v4121-v4121)+16))
	v4127 = v4126
	goto L534
L538:
	;
	if v4108 < int32(0) {
		goto L543
	} else {
		goto L544
	}
L539:
	;
	v4148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4105)+4)))
	v4150 = v4148 & int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v4105)+4)) = uint16(v4150)
	goto L538
L540:
	;
	v4137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4134)+4)))
	if v4137 == int32(0) {
		goto L539
	} else {
		goto L541
	}
L541:
	;
	v4140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4105)+4)))
	v4145 = v4140&int32(_a_F_spgdoinsert_5) | v4137&int32(_a_F_spgdoinsert_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v4105)+4)) = uint16(v4145)
	goto L538
L542:
	;
	v4174 = *(*int32)(unsafe.Add(mBase, uint32(v4105)))
	v4179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4106))))
	v4183 = F_SpGistPageAddNewItem(m, v4173, v4105, int32(base.Ui32(v4174)>>(uint(int32(2))%32)), v60+int32(1052)+v4179<<(uint(int32(1))%32))
	mBase = m.M
	v4184 = m.ExcPending
	if v4184 != 0 {
		goto L3
	} else {
		goto L546
	}
L543:
	;
	v4159 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v4165 = *(*int32)(unsafe.Add(mBase, uint32(v4159+(v4108^int32(-1))<<(uint(int32(2))%32))))
	v4173 = v4165
	goto L542
L544:
	;
	goto L545
L545:
	;
	v4167 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v4173 = v4167 + v4108<<(uint(int32(13))%32) + int32(-8192)
	goto L542
L546:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1476+v4052<<(uint(int32(1))%32)))) = uint16(v4183)
	v4186 = *(*int32)(unsafe.Add(mBase, uint32(v4133)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4186)+4)) = uint16(v4183)
	*(*uint16)(unsafe.Add(mBase, uint32(v4186)+2)) = uint16(v4127)
	v4190 = int32(base.Ui32(v4127) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v4186))) = uint16(v4190)
	v4192 = *(*int32)(unsafe.Add(mBase, uint32(v4104)))
	v4193 = *(*int32)(unsafe.Add(mBase, uint32(v4192)))
	v4195 = int32(base.Ui32(v4193) >> (uint(int32(2)) % 32))
	if v4195 != 0 {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	base.MemoryCopy(m, v4065, v4192, v4195)
	goto L549
L548:
	;
	goto L549
L549:
	;
	v4197 = *(*int32)(unsafe.Add(mBase, uint32(v4104)))
	v4198 = *(*int32)(unsafe.Add(mBase, uint32(v4197)))
	v4201 = v4065 + int32(base.Ui32(v4198)>>(uint(int32(2))%32))
	v4203 = v4052 + int32(1)
	if v4203 != v3909 {
		v4052 = v4203
		v4065 = v4201
		goto L529
	} else {
		goto L550
	}
L550:
	;
	goto L530
L551:
	;
	F_MarkBufferDirty(m, v3901)
	mBase = m.M
	v4263 = m.ExcPending
	if v4263 != 0 {
		goto L3
	} else {
		goto L554
	}
L552:
	;
	goto L553
L553:
	;
	v4264 = int32(0)
	if base.B2i32(v3157 == v4264)|base.B2i32(v3157 != v341) == v4264 {
		goto L557
	} else {
		goto L558
	}
L554:
	;
	goto L553
L555:
	;
	F_MarkBufferDirty(m, v431)
	mBase = m.M
	v4429 = m.ExcPending
	if v4429 != 0 {
		goto L3
	} else {
		goto L590
	}
L556:
	;
	v4422 = v340
	v4423 = v341
	v4425 = v4272
	v4426 = v348
	v4427 = v431
	goto L555
L557:
	;
	v4270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2829)+4)))
	v4272 = F_SpGistPageAddNewItem(m, v340, v2829, v4270, int32(0))
	mBase = m.M
	v4273 = m.ExcPending
	if v4273 != 0 {
		goto L3
	} else {
		goto L560
	}
L558:
	;
	goto L559
L559:
	;
	if v341 != 0 {
		goto L563
	} else {
		goto L564
	}
L560:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+740)) = uint16(v353)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+738)) = uint16(v344)
	v4276 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+736)) = uint8(v4276)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+732)) = uint16(v4272)
	F_saveNodeLink(m, v60+int32(444), v348, v4272)
	mBase = m.M
	v4282 = m.ExcPending
	if v4282 != 0 {
		goto L3
	} else {
		goto L561
	}
L561:
	;
	if v4040 == int32(0) {
		goto L556
	} else {
		goto L562
	}
L562:
	;
	v4288 = *(*int32)(unsafe.Add(mBase, uint32(v450+v4040<<(uint(int32(2))%32))+20))
	v4291 = v450 + v4288&int32(_a_F_spgdoinsert_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v4291)+10)) = uint16(v4272)
	*(*uint16)(unsafe.Add(mBase, uint32(v4291)+8)) = uint16(v348)
	v4295 = int32(base.Ui32(v348) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v4291)+6)) = uint16(v4295)
	goto L556
L563:
	;
	if v3157 < int32(0) {
		goto L567
	} else {
		goto L568
	}
L564:
	;
	goto L565
L565:
	;
	if v431 < int32(0) {
		goto L580
	} else {
		goto L581
	}
L566:
	;
	if v3157 < int32(0) {
		goto L571
	} else {
		goto L572
	}
L567:
	;
	v4300 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[1]))
	v4306 = *(*int32)(unsafe.Add(mBase, uint32(v4300+(v3157^int32(-1))*int32(56))+16))
	v4315 = v4306
	goto L566
L568:
	;
	goto L569
L569:
	;
	v4308 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[2]))
	v4309 = int32(56)
	v4314 = *(*int32)(unsafe.Add(mBase, uint32(v4308+v3157*v4309-v4309)+16))
	v4315 = v4314
	goto L566
L570:
	;
	v4334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2829)+4)))
	v4336 = F_SpGistPageAddNewItem(m, v4333, v2829, v4334, int32(0))
	mBase = m.M
	v4337 = m.ExcPending
	if v4337 != 0 {
		goto L3
	} else {
		goto L574
	}
L571:
	;
	v4319 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v4325 = *(*int32)(unsafe.Add(mBase, uint32(v4319+(v3157^int32(-1))<<(uint(int32(2))%32))))
	v4333 = v4325
	goto L570
L572:
	;
	goto L573
L573:
	;
	v4327 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v4333 = v4327 + v3157<<(uint(int32(13))%32) + int32(-8192)
	goto L570
L574:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+732)) = uint16(v4336)
	F_MarkBufferDirty(m, v3157)
	mBase = m.M
	v4340 = m.ExcPending
	if v4340 != 0 {
		goto L3
	} else {
		goto L575
	}
L575:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+740)) = uint16(v353)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+738)) = uint16(v344)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+736)) = uint8(base.B2i32(v3157 == v341))
	F_saveNodeLink(m, v60+int32(444), v4315, v4336)
	mBase = m.M
	v4348 = m.ExcPending
	if v4348 != 0 {
		goto L3
	} else {
		goto L576
	}
L576:
	;
	if v4040 == int32(0) {
		v4422 = v4333
		v4423 = v3157
		v4425 = v4336
		v4426 = v4315
		v4427 = v431
		goto L555
	} else {
		goto L577
	}
L577:
	;
	v4354 = *(*int32)(unsafe.Add(mBase, uint32(v450+v4040<<(uint(int32(2))%32))+20))
	v4357 = v450 + v4354&int32(_a_F_spgdoinsert_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v4357)+10)) = uint16(v4336)
	*(*uint16)(unsafe.Add(mBase, uint32(v4357)+8)) = uint16(v4315)
	v4361 = int32(base.Ui32(v4315) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v4357)+6)) = uint16(v4361)
	v4422 = v4333
	v4423 = v3157
	v4425 = v4336
	v4426 = v4315
	v4427 = v431
	goto L555
L578:
	;
	v4389 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+736)) = uint8(v4389)
	v4391 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+734)) = uint8(v4391)
	v4393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2829)+4)))
	v4396 = F_PageAddItemExtended(m, v450, v2829, v4393, v4389, v4389)
	mBase = m.M
	v4397 = m.ExcPending
	if v4397 != 0 {
		goto L3
	} else {
		goto L583
	}
L579:
	;
	F_PageInit(m, v4380, int32(_a_F_spgdoinsert_25), int32(8))
	mBase = m.M
	v4384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4380)+16)))
	v4385 = v4380 + v4384
	v4386 = int32(_a_F_spgdoinsert_26)
	*(*uint16)(unsafe.Add(mBase, uint32(v4385)+6)) = uint16(v4386)
	*(*uint16)(unsafe.Add(mBase, uint32(v4385))) = uint16(v299)
	goto L578
L580:
	;
	v4366 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v4372 = *(*int32)(unsafe.Add(mBase, uint32(v4366+(v431^int32(-1))<<(uint(int32(2))%32))))
	v4380 = v4372
	goto L579
L581:
	;
	goto L582
L582:
	;
	v4374 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v4380 = v4374 + v431<<(uint(int32(13))%32) + int32(-8192)
	goto L579
L583:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+732)) = uint16(v4396)
	if v4396 == int32(1) {
		goto L584
	} else {
		goto L585
	}
L584:
	;
	v4401 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+738)) = v4401
	v4422 = v450
	v4423 = v431
	v4425 = int32(1)
	v4426 = v432
	v4427 = v4401
	goto L555
L585:
	;
	goto L586
L586:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4408 = m.ExcPending
	if v4408 != 0 {
		goto L3
	} else {
		goto L587
	}
L587:
	;
	v4409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2829)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+176)) = v4409
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_8), v60+int32(176))
	mBase = m.M
	v4415 = m.ExcPending
	if v4415 != 0 {
		goto L3
	} else {
		goto L588
	}
L588:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1343), int32(_a_F_spgdoinsert_16))
	mBase = m.M
	v4420 = m.ExcPending
	if v4420 != 0 {
		goto L3
	} else {
		goto L589
	}
L589:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L590:
	;
	v4430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4430)+118)))
	if v4431 != int32(112) {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	v4555 = int32(_a_F_spgdoinsert_4)
	v4557 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v4557 - int32(1)
	if v3901 != 0 {
		goto L645
	} else {
		goto L646
	}
L592:
	;
	v4435 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[6]))
	if v4435 <= int32(0) {
		goto L593
	} else {
		goto L594
	}
L593:
	;
	v4438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v4438 != 0 {
		goto L591
	} else {
		goto L596
	}
L594:
	;
	goto L595
L595:
	;
	v4440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v4440 != 0 {
		goto L591
	} else {
		goto L598
	}
L596:
	;
	v4439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4439 != 0 {
		goto L591
	} else {
		goto L597
	}
L597:
	;
	goto L595
L598:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v4442 = m.ExcPending
	if v4442 != 0 {
		goto L3
	} else {
		goto L599
	}
L599:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+728)) = uint16(v3909)
	F_XLogRegisterData(m, v60+int32(724), int32(28))
	mBase = m.M
	v4448 = m.ExcPending
	if v4448 != 0 {
		goto L3
	} else {
		goto L600
	}
L600:
	;
	v4449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+726)))
	F_XLogRegisterData(m, v1473, v4449<<(uint(int32(1))%32))
	mBase = m.M
	v4453 = m.ExcPending
	if v4453 != 0 {
		goto L3
	} else {
		goto L601
	}
L601:
	;
	v4454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+728)))
	F_XLogRegisterData(m, v1476, v4454<<(uint(int32(1))%32))
	mBase = m.M
	v4458 = m.ExcPending
	if v4458 != 0 {
		goto L3
	} else {
		goto L602
	}
L602:
	;
	v4459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+728)))
	F_XLogRegisterData(m, v1485, v4459)
	mBase = m.M
	v4461 = m.ExcPending
	if v4461 != 0 {
		goto L3
	} else {
		goto L603
	}
L603:
	;
	v4462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2829)+4)))
	F_XLogRegisterData(m, v2829, v4462)
	mBase = m.M
	v4464 = m.ExcPending
	if v4464 != 0 {
		goto L3
	} else {
		goto L604
	}
L604:
	;
	F_XLogRegisterData(m, v3956, v4225-v3956)
	mBase = m.M
	v4467 = m.ExcPending
	if v4467 != 0 {
		goto L3
	} else {
		goto L605
	}
L605:
	;
	if v4427 != 0 {
		goto L606
	} else {
		goto L607
	}
L606:
	;
	v4471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+730)))
	if v4471 != 0 {
		goto L609
	} else {
		goto L610
	}
L607:
	;
	goto L608
L608:
	;
	if v3901 != 0 {
		goto L613
	} else {
		goto L614
	}
L609:
	;
	v4472 = int32(14)
	goto L611
L610:
	;
	v4472 = int32(8)
	goto L611
L611:
	;
	F_XLogRegisterBuffer(m, int32(0), v4427, v4472)
	mBase = m.M
	v4474 = m.ExcPending
	if v4474 != 0 {
		goto L3
	} else {
		goto L612
	}
L612:
	;
	goto L608
L613:
	;
	v4478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+731)))
	if v4478 != 0 {
		goto L616
	} else {
		goto L617
	}
L614:
	;
	goto L615
L615:
	;
	v4485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+734)))
	if v4485 != 0 {
		goto L620
	} else {
		goto L621
	}
L616:
	;
	v4479 = int32(14)
	goto L618
L617:
	;
	v4479 = int32(8)
	goto L618
L618:
	;
	F_XLogRegisterBuffer(m, int32(1), v3901, v4479)
	mBase = m.M
	v4481 = m.ExcPending
	if v4481 != 0 {
		goto L3
	} else {
		goto L619
	}
L619:
	;
	goto L615
L620:
	;
	v4486 = int32(14)
	goto L622
L621:
	;
	v4486 = int32(8)
	goto L622
L622:
	;
	F_XLogRegisterBuffer(m, int32(2), v4423, v4486)
	mBase = m.M
	v4488 = m.ExcPending
	if v4488 != 0 {
		goto L3
	} else {
		goto L623
	}
L623:
	;
	v4489 = int32(0)
	if base.B2i32(v341 == v4489)|base.B2i32(v4423 == v341) == v4489 {
		goto L624
	} else {
		goto L625
	}
L624:
	;
	F_XLogRegisterBuffer(m, int32(3), v341, int32(8))
	mBase = m.M
	v4498 = m.ExcPending
	if v4498 != 0 {
		goto L3
	} else {
		goto L627
	}
L625:
	;
	goto L626
L626:
	;
	v4501 = F_XLogInsert(m, int32(16), int32(80))
	mBase = m.M
	v4502 = m.ExcPending
	if v4502 != 0 {
		goto L3
	} else {
		goto L628
	}
L627:
	;
	goto L626
L628:
	;
	if v3901 != 0 {
		goto L629
	} else {
		goto L630
	}
L629:
	;
	if v3901 < int32(0) {
		goto L633
	} else {
		goto L634
	}
L630:
	;
	goto L631
L631:
	;
	if v4427 == int32(0) {
		goto L637
	} else {
		goto L638
	}
L632:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4520))) = base.I64_rotl(v4501, int64(32))
	goto L631
L633:
	;
	v4506 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v4512 = *(*int32)(unsafe.Add(mBase, uint32(v4506+(v3901^int32(-1))<<(uint(int32(2))%32))))
	v4520 = v4512
	goto L632
L634:
	;
	goto L635
L635:
	;
	v4514 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v4520 = v4514 + v3901<<(uint(int32(13))%32) + int32(-8192)
	goto L632
L636:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4422))) = v4549
	if v341 == int32(0) {
		goto L591
	} else {
		goto L644
	}
L637:
	;
	v4549 = base.I64_rotl(v4501, int64(32))
	goto L636
L638:
	;
	goto L639
L639:
	;
	if v4427 < int32(0) {
		goto L641
	} else {
		goto L642
	}
L640:
	;
	v4547 = base.I64_rotl(v4501, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v4545))) = v4547
	v4549 = v4547
	goto L636
L641:
	;
	v4531 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v4537 = *(*int32)(unsafe.Add(mBase, uint32(v4531+(v4427^int32(-1))<<(uint(int32(2))%32))))
	v4545 = v4537
	goto L640
L642:
	;
	goto L643
L643:
	;
	v4539 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v4545 = v4539 + v4427<<(uint(int32(13))%32) + int32(-8192)
	goto L640
L644:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v340))) = v4549
	goto L591
L645:
	;
	F_SpGistSetLastUsedPage(m, l0, v3901)
	mBase = m.M
	v4562 = m.ExcPending
	if v4562 != 0 {
		goto L3
	} else {
		goto L648
	}
L646:
	;
	goto L647
L647:
	;
	if v4427 != 0 {
		goto L650
	} else {
		goto L651
	}
L648:
	;
	F_UnlockReleaseBuffer(m, v3901)
	mBase = m.M
	v4564 = m.ExcPending
	if v4564 != 0 {
		goto L3
	} else {
		goto L649
	}
L649:
	;
	goto L647
L650:
	;
	F_SpGistSetLastUsedPage(m, l0, v4427)
	mBase = m.M
	v4566 = m.ExcPending
	if v4566 != 0 {
		goto L3
	} else {
		goto L653
	}
L651:
	;
	goto L652
L652:
	;
	if v3924 != 0 {
		goto L655
	} else {
		goto L656
	}
L653:
	;
	F_UnlockReleaseBuffer(m, v4427)
	mBase = m.M
	v4568 = m.ExcPending
	if v4568 != 0 {
		goto L3
	} else {
		goto L654
	}
L654:
	;
	goto L652
L655:
	;
	v7056 = int32(1)
	v7060 = v4423
	goto L30
L656:
	;
	goto L657
L657:
	;
	F_pfree(m, v467)
	mBase = m.M
	v4571 = m.ExcPending
	if v4571 != 0 {
		goto L3
	} else {
		goto L658
	}
L658:
	;
	v4582 = v4422
	v4583 = v4423
	v4586 = v4425
	v4590 = v4426
	goto L64
L659:
	;
	v4642 = v4582
	v4643 = v4583
	v4646 = v4586
	v4650 = v4590
	goto L671
L660:
	;
	v6828 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+443)) = uint8(v6828)
	v6831 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v6831 == v6828 {
		v336 = v6825
		v337 = v5275
		v338 = v6826
		v339 = v5268
		v340 = v4642
		v341 = v4643
		v344 = v4646
		v348 = v4650
		v353 = v4907
		v356 = v5287
		v371 = v5276 + v371
		goto L62
	} else {
		goto L1012
	}
L661:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6812 = m.ExcPending
	if v6812 != 0 {
		goto L3
	} else {
		goto L1009
	}
L662:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6799 = m.ExcPending
	if v6799 != 0 {
		goto L3
	} else {
		goto L1006
	}
L663:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6783 = m.ExcPending
	if v6783 != 0 {
		goto L3
	} else {
		goto L1003
	}
L664:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6767 = m.ExcPending
	if v6767 != 0 {
		goto L3
	} else {
		goto L1000
	}
L665:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6749 = m.ExcPending
	if v6749 != 0 {
		goto L3
	} else {
		goto L997
	}
L666:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6736 = m.ExcPending
	if v6736 != 0 {
		goto L3
	} else {
		goto L994
	}
L667:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6723 = m.ExcPending
	if v6723 != 0 {
		goto L3
	} else {
		goto L991
	}
L668:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6707 = m.ExcPending
	if v6707 != 0 {
		goto L3
	} else {
		goto L988
	}
L669:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6694 = m.ExcPending
	if v6694 != 0 {
		goto L3
	} else {
		goto L985
	}
L670:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6681 = m.ExcPending
	if v6681 != 0 {
		goto L3
	} else {
		goto L982
	}
L671:
	;
	v4689 = int32(1)
	v4690 = v4650 - v4689
	v4694 = base.I32_rem_u_s(v4650+v4689, int32(3))
	v4696 = v4646 & int32(_a_F_spgdoinsert_0)
	v4701 = v4642 + v4696<<(uint(int32(2))%32) + int32(20)
	goto L673
L672:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6665 = m.ExcPending
	if v6665 != 0 {
		goto L3
	} else {
		goto L979
	}
L673:
	;
	v4759 = *(*int32)(unsafe.Add(mBase, uint32(v4701)))
	v4760 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+416)) = v371
	v4762 = *(*int64)(unsafe.Add(mBase, uint32(v60)+464))
	*(*int64)(unsafe.Add(mBase, uint32(v60)+408)) = v4762
	*(*int64)(unsafe.Add(mBase, uint32(v60)+400)) = v4760
	v4767 = v4642 + v4759&int32(_a_F_spgdoinsert_6)
	v4768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4767))))
	v4772 = int32(base.Ui32(v4768)>>(uint(int32(2))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+420)) = uint8(v4772)
	v4774 = *(*int32)(unsafe.Add(mBase, uint32(v4767)))
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+421)) = uint8(base.B2i32(base.Ui32(int32(_a_F_spgdoinsert_0)) < base.Ui32(v4774)))
	v4779 = *(*int32)(unsafe.Add(mBase, uint32(v4767)))
	if base.Ui32(v4779) < base.Ui32(int32(_a_F_spgdoinsert_27)) {
		v4790 = int64(0)
		goto L675
	} else {
		goto L676
	}
L674:
	;
	goto L672
L675:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v60)+424)) = v4790
	v4792 = *(*int32)(unsafe.Add(mBase, uint32(v4767)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+432)) = int32(base.Ui32(v4792)>>(uint(int32(3))%32)) & int32(_a_F_spgdoinsert_22)
	v4798 = F_spgExtractNodeLabels(m, l1, v4767)
	mBase = m.M
	v4799 = m.ExcPending
	if v4799 != 0 {
		goto L3
	} else {
		goto L680
	}
L676:
	;
	v4783 = v4767 + int32(8)
	v4784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+50)))
	if v4784 == int32(1) {
		goto L677
	} else {
		goto L678
	}
L677:
	;
	v4787 = *(*int64)(unsafe.Add(mBase, uint32(v4783)))
	v4790 = v4787
	goto L675
L678:
	;
	goto L679
L679:
	;
	v4790 = base.I64_extend_i32_u(v4783)
	goto L675
L680:
	;
	v4800 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+352)) = v4800
	*(*int32)(unsafe.Add(mBase, uint32(v60)+436)) = v4798
	*(*int64)(unsafe.Add(mBase, uint32(v60)+360)) = v4800
	*(*int64)(unsafe.Add(mBase, uint32(v60)+368)) = v4800
	*(*int64)(unsafe.Add(mBase, uint32(v60)+376)) = v4800
	*(*int64)(unsafe.Add(mBase, uint32(v60)+384)) = v4800
	*(*int64)(unsafe.Add(mBase, uint32(v60)+392)) = v4800
	if v64 == int32(0) {
		goto L682
	} else {
		goto L683
	}
L681:
	;
	v4824 = *(*int32)(unsafe.Add(mBase, uint32(v4767)))
	if v4824&int32(4) == int32(0) {
		v4902 = v4823
		goto L686
	} else {
		goto L687
	}
L682:
	;
	v4815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v4816 = *(*int32)(unsafe.Add(mBase, uint32(v4815)))
	v4817 = F_FunctionCall2Coll(m, v105, v4816, base.I64_extend_i32_u(v60+int32(400)), base.I64_extend_i32_u(v60+int32(352)))
	mBase = m.M
	v4818 = m.ExcPending
	if v4818 != 0 {
		goto L3
	} else {
		goto L685
	}
L683:
	;
	goto L684
L684:
	;
	v4820 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+352)) = v4820
	v4823 = v4820
	goto L681
L685:
	;
	v4819 = *(*int32)(unsafe.Add(mBase, uint32(v60)+352))
	v4823 = v4819
	goto L681
L686:
	;
	if v4902 != int32(3) {
		goto L701
	} else {
		goto L702
	}
L687:
	;
	switch v4823 - int32(1) {
	case 0:
		goto L688
	case 1:
		goto L689
	default:
		v4902 = v4823
		goto L686
	}
L688:
	;
	v4844 = int32(_a_F_spgdoinsert_28)
	v4845 = int64(0)
	v4852 = base.I64_extend_i32_s(int32(base.Ui32(v4824)>>(uint(int32(3))%32))&int32(_a_F_spgdoinsert_22) - int32(1))
	if base.Ui64(v4852) <= base.Ui64(v4845) {
		goto L694
	} else {
		goto L695
	}
L689:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4834 = m.ExcPending
	if v4834 != 0 {
		goto L3
	} else {
		goto L690
	}
L690:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_29), int32(0))
	mBase = m.M
	v4838 = m.ExcPending
	if v4838 != 0 {
		goto L3
	} else {
		goto L691
	}
L691:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(2207), int32(_a_F_spgdoinsert_18))
	mBase = m.M
	v4843 = m.ExcPending
	if v4843 != 0 {
		goto L3
	} else {
		goto L692
	}
L692:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L693:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v60)+360)) = uint32(v4899)
	v4901 = *(*int32)(unsafe.Add(mBase, uint32(v60)+352))
	v4902 = v4901
	goto L686
L694:
	;
	v4899 = v4845
	goto L693
L695:
	;
	goto L696
L696:
	;
	v4859 = v4852 - v4845
	v4861 = *(*int64)(unsafe.Add(mBase, _c_F_spgdoinsert[7]))
	v4862 = *(*int64)(unsafe.Add(mBase, _c_F_spgdoinsert[8]))
	v4865 = v4862
	v4867 = v4861
	goto L697
L697:
	;
	v4871 = v4865 ^ v4867
	v4873 = base.I64_rotl(v4871, int64(37))
	v4881 = v4871 ^ (v4871<<(uint(int64(16))%64) ^ base.I64_rotl(v4865, int64(24)))
	v4886 = int64(base.Ui64(base.I64_rotl(v4865*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v4859)) % 64))
	if base.Ui64(v4859) < base.Ui64(v4886) {
		v4865 = v4881
		v4867 = v4873
		goto L697
	} else {
		goto L699
	}
L698:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_spgdoinsert[7])) = v4873
	*(*int64)(unsafe.Add(mBase, _c_F_spgdoinsert[8])) = v4881
	v4899 = v4845 + v4886
	goto L693
L699:
	;
	goto L698
L700:
	;
	goto L674
L701:
	;
	switch v4902 - int32(1) {
	case 0:
		goto L705
	case 1:
		goto L704
	default:
		goto L700
	}
L702:
	;
	goto L703
L703:
	;
	v5828 = *(*int32)(unsafe.Add(mBase, uint32(v60)+376))
	v5829 = int32(-8192)
	if base.Ui32(v5828+v5829) <= base.Ui32(v5829) {
		goto L664
	} else {
		goto L868
	}
L704:
	;
	v5333 = *(*int32)(unsafe.Add(mBase, uint32(v60)+436))
	if v5333 == int32(0) {
		goto L670
	} else {
		goto L751
	}
L705:
	;
	v4907 = *(*int32)(unsafe.Add(mBase, uint32(v60)+360))
	v4908 = int32(0)
	if base.B2i32(v341 == v4908)|base.B2i32(v4643 == v341) == v4908 {
		goto L706
	} else {
		goto L707
	}
L706:
	;
	F_SpGistSetLastUsedPage(m, l0, v341)
	mBase = m.M
	v4915 = m.ExcPending
	if v4915 != 0 {
		goto L3
	} else {
		goto L709
	}
L707:
	;
	goto L708
L708:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+460)) = v4907
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+456)) = uint16(v4646)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+452)) = v4642
	*(*int32)(unsafe.Add(mBase, uint32(v60)+448)) = v4643
	*(*int32)(unsafe.Add(mBase, uint32(v60)+444)) = v4650
	v4923 = *(*int32)(unsafe.Add(mBase, uint32(v4767)))
	v4928 = v4767 + int32(base.Ui32(v4923)>>(uint(int32(16))%32)) + int32(8)
	if v4907 == int32(0) {
		goto L712
	} else {
		goto L713
	}
L709:
	;
	F_UnlockReleaseBuffer(m, v341)
	mBase = m.M
	v4917 = m.ExcPending
	if v4917 != 0 {
		goto L3
	} else {
		goto L710
	}
L710:
	;
	goto L708
L711:
	;
	v5268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5218)+4)))
	if v5268 != 0 {
		goto L731
	} else {
		goto L732
	}
L712:
	;
	if v4907 != 0 {
		goto L32
	} else {
		goto L730
	}
L713:
	;
	v4934 = int32(base.Ui32(v4923)>>(uint(int32(3))%32)) & int32(_a_F_spgdoinsert_22)
	if v4934 == int32(0) {
		goto L712
	} else {
		goto L714
	}
L714:
	;
	v4937 = int32(1)
	v4938 = v4907 - v4937
	v4940 = v4934 - v4937
	if base.Ui32(v4938) < base.Ui32(v4940) {
		goto L715
	} else {
		goto L716
	}
L715:
	;
	v4942 = v4938
	goto L717
L716:
	;
	v4942 = v4940
	goto L717
L717:
	;
	v4944 = v4942 + int32(1)
	v4945 = int32(3)
	v4946 = v4944 & v4945
	if base.Ui32(v4945) <= base.Ui32(v4942) {
		goto L719
	} else {
		goto L720
	}
L718:
	;
	if v4944 == v4907 {
		v5218 = v5159
		goto L711
	} else {
		goto L729
	}
L719:
	;
	v4959 = v4928
	v4960 = int32(0)
	goto L722
L720:
	;
	v5037 = v4928
	goto L721
L721:
	;
	v5095 = v5037
	v5096 = int32(0)
	goto L726
L722:
	;
	v5009 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4959)+6)))
	v5010 = int32(_a_F_spgdoinsert_22)
	v5012 = v4959 + v5009&v5010
	v5013 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5012)+6)))
	v5016 = v5012 + v5013&v5010
	v5017 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5016)+6)))
	v5020 = v5016 + v5017&v5010
	v5021 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5020)+6)))
	v5024 = v5020 + v5021&v5010
	v5026 = v4960 + int32(4)
	if v5026 != v4944&int32(-4) {
		v4959 = v5024
		v4960 = v5026
		goto L722
	} else {
		goto L724
	}
L723:
	;
	if v4946 == int32(0) {
		v5159 = v5024
		goto L718
	} else {
		goto L725
	}
L724:
	;
	goto L723
L725:
	;
	v5037 = v5024
	goto L721
L726:
	;
	v5145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5095)+6)))
	v5148 = v5095 + v5145&int32(_a_F_spgdoinsert_22)
	v5150 = v5096 + int32(1)
	if v5150 != v4946 {
		v5095 = v5148
		v5096 = v5150
		goto L726
	} else {
		goto L728
	}
L727:
	;
	v5159 = v5148
	goto L718
L728:
	;
	goto L727
L729:
	;
	goto L32
L730:
	;
	v5218 = v4928
	goto L711
L731:
	;
	v5269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5218)+2)))
	v5270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5218))))
	v5275 = v5269 | v5270<<(uint(int32(16))%32)
	goto L733
L732:
	;
	v5275 = int32(-1)
	goto L733
L733:
	;
	v5276 = *(*int32)(unsafe.Add(mBase, uint32(v60)+364))
	if v64 == int32(0) {
		goto L734
	} else {
		goto L735
	}
L734:
	;
	v5279 = *(*int64)(unsafe.Add(mBase, uint32(v60)+368))
	*(*int64)(unsafe.Add(mBase, uint32(v60)+464)) = v5279
	v5283 = F_SpGistGetLeafTupleSize(m, v62, v60+int32(464), l4)
	mBase = m.M
	v5284 = m.ExcPending
	if v5284 != 0 {
		goto L3
	} else {
		goto L737
	}
L735:
	;
	v5287 = v356
	goto L736
L736:
	;
	if base.Ui32(v5287) < base.Ui32(int32(_a_F_spgdoinsert_20)) {
		goto L738
	} else {
		goto L739
	}
L737:
	;
	v5287 = v5283 + int32(4)
	goto L736
L738:
	;
	v6825 = v336
	v6826 = v338
	goto L660
L739:
	;
	goto L740
L740:
	;
	if v64 != 0 {
		goto L741
	} else {
		goto L742
	}
L741:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5306 = m.ExcPending
	if v5306 != 0 {
		goto L3
	} else {
		goto L746
	}
L742:
	;
	v5290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v5290&int32(1) == int32(0) {
		goto L741
	} else {
		goto L743
	}
L743:
	;
	if v5287 < v336 {
		v6825 = v5287
		v6826 = int32(0)
		goto L660
	} else {
		goto L744
	}
L744:
	;
	v5298 = v338 + int32(1)
	if v5298 < int32(10) {
		v6825 = v336
		v6826 = v5298
		goto L660
	} else {
		goto L745
	}
L745:
	;
	goto L741
L746:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v5309 = m.ExcPending
	if v5309 != 0 {
		goto L3
	} else {
		goto L747
	}
L747:
	;
	v5310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+36)) = int32(_a_F_spgdoinsert_30)
	v5313 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+32)) = v5287 - v5313
	*(*int32)(unsafe.Add(mBase, uint32(v60)+40)) = v5310 + v5313
	F_errmsg(m, int32(_a_F_spgdoinsert_31), v60+int32(32))
	mBase = m.M
	v5323 = m.ExcPending
	if v5323 != 0 {
		goto L3
	} else {
		goto L748
	}
L748:
	;
	F_errhint(m, int32(_a_F_spgdoinsert_32), int32(0))
	mBase = m.M
	v5327 = m.ExcPending
	if v5327 != 0 {
		goto L3
	} else {
		goto L749
	}
L749:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(2277), int32(_a_F_spgdoinsert_18))
	mBase = m.M
	v5332 = m.ExcPending
	if v5332 != 0 {
		goto L3
	} else {
		goto L750
	}
L750:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L751:
	;
	v5336 = *(*int64)(unsafe.Add(mBase, uint32(v60)+360))
	v5337 = *(*int32)(unsafe.Add(mBase, uint32(v4767)))
	v5341 = int32(base.Ui32(v5337)>>(uint(int32(3))%32)) & int32(_a_F_spgdoinsert_22)
	v5342 = *(*int32)(unsafe.Add(mBase, uint32(v60)+368))
	if int32(0) <= v5342 {
		goto L752
	} else {
		goto L753
	}
L752:
	;
	if base.Ui32(v5341) < base.Ui32(v5342) {
		goto L669
	} else {
		goto L755
	}
L753:
	;
	v5346 = v5341
	goto L754
L754:
	;
	v5348 = v4767 + int32(8)
	v5352 = F_palloc_mul(m, int32(4), v5341+int32(1))
	mBase = m.M
	v5353 = m.ExcPending
	if v5353 != 0 {
		goto L3
	} else {
		goto L756
	}
L755:
	;
	v5346 = v5342
	goto L754
L756:
	;
	v5354 = *(*int32)(unsafe.Add(mBase, uint32(v4767)))
	if v5354&int32(_a_F_spgdoinsert_21) != 0 {
		goto L757
	} else {
		goto L758
	}
L757:
	;
	v5368 = v5348 + int32(base.Ui32(v5354)>>(uint(int32(16))%32))
	v5369 = int32(0)
	goto L760
L758:
	;
	goto L759
L759:
	;
	v5497 = F_spgFormNodeTuple(m, l1, v5336, int32(0))
	mBase = m.M
	v5498 = m.ExcPending
	if v5498 != 0 {
		goto L3
	} else {
		goto L767
	}
L760:
	;
	v5420 = v5352 + v5369<<(uint(int32(2))%32)
	if v5369 < v5346 {
		goto L763
	} else {
		goto L764
	}
L761:
	;
	goto L759
L762:
	;
	v5424 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5368)+6)))
	v5425 = int32(_a_F_spgdoinsert_22)
	v5429 = v5369 + int32(1)
	v5430 = *(*int32)(unsafe.Add(mBase, uint32(v4767)))
	if base.Ui32(v5429) < base.Ui32(int32(base.Ui32(v5430)>>(uint(int32(3))%32))&v5425) {
		v5368 = v5368 + v5424&v5425
		v5369 = v5429
		goto L760
	} else {
		goto L766
	}
L763:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5420))) = v5368
	goto L762
L764:
	;
	goto L765
L765:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5420)+4)) = v5368
	goto L762
L766:
	;
	goto L761
L767:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5352+v5346<<(uint(int32(2))%32)))) = v5497
	v5501 = *(*int32)(unsafe.Add(mBase, uint32(v4767)))
	v5503 = int32(base.Ui32(v5501) >> (uint(int32(16)) % 32))
	if v5503 == int32(0) {
		v5511 = int64(0)
		goto L768
	} else {
		goto L769
	}
L768:
	;
	v5520 = F_spgFormInnerTuple(m, l1, base.B2i32(v5503 != int32(0)), v5511, int32(base.Ui32(v5501)>>(uint(int32(3))%32))&int32(_a_F_spgdoinsert_22)+int32(1), v5352)
	mBase = m.M
	v5521 = m.ExcPending
	if v5521 != 0 {
		goto L3
	} else {
		goto L773
	}
L769:
	;
	v5506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+50)))
	if v5506 == int32(1) {
		goto L770
	} else {
		goto L771
	}
L770:
	;
	v5509 = *(*int64)(unsafe.Add(mBase, uint32(v5348)))
	v5511 = v5509
	goto L768
L771:
	;
	goto L772
L772:
	;
	v5511 = base.I64_extend_i32_u(v5348)
	goto L768
L773:
	;
	v5522 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+796)) = v5522
	v5524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+784)) = uint16(v4646)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+800)) = uint8(v5524)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+786)) = int64(4278190080)
	v5529 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4642)+14)))
	v5530 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4642)+12)))
	v5531 = v5529 - v5530
	v5532 = int32(0)
	if v5532 < v5531 {
		goto L776
	} else {
		goto L777
	}
L774:
	;
	v5823 = int32(0)
	v5825 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v5825 == v5823 {
		v4642 = v5817
		v4643 = v5818
		v4646 = v5819
		v4650 = v5821
		goto L671
	} else {
		goto L867
	}
L775:
	;
	v5536 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5520)+4)))
	v5537 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4767)+4)))
	if base.Ui32(v5536-v5537) <= base.Ui32(v5535) {
		goto L779
	} else {
		goto L780
	}
L776:
	;
	v5535 = v5531
	goto L778
L777:
	;
	v5535 = v5532
	goto L778
L778:
	;
	goto L775
L779:
	;
	v5540 = int32(_a_F_spgdoinsert_4)
	v5542 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v5542 + int32(1)
	F_PageIndexTupleDelete(m, v4642, v4696)
	mBase = m.M
	v5547 = m.ExcPending
	if v5547 != 0 {
		goto L3
	} else {
		goto L782
	}
L780:
	;
	goto L781
L781:
	;
	if base.Ui32(v4690) <= base.Ui32(int32(1)) {
		goto L667
	} else {
		goto L799
	}
L782:
	;
	v5548 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5520)+4)))
	v5550 = F_PageAddItemExtended(m, v4642, v5520, v5548, v4696, int32(0))
	mBase = m.M
	v5551 = m.ExcPending
	if v5551 != 0 {
		goto L3
	} else {
		goto L783
	}
L783:
	;
	if v5550 != v4696 {
		goto L668
	} else {
		goto L784
	}
L784:
	;
	F_MarkBufferDirty(m, v4643)
	mBase = m.M
	v5554 = m.ExcPending
	if v5554 != 0 {
		goto L3
	} else {
		goto L785
	}
L785:
	;
	v5555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5555)+118)))
	if v5556 != int32(112) {
		goto L786
	} else {
		goto L787
	}
L786:
	;
	v5587 = int32(_a_F_spgdoinsert_4)
	v5589 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v5589 - int32(1)
	v5817 = v4642
	v5818 = v4643
	v5819 = v4646
	v5821 = v4650
	goto L774
L787:
	;
	v5560 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[6]))
	if v5560 <= int32(0) {
		goto L788
	} else {
		goto L789
	}
L788:
	;
	v5563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v5563 != 0 {
		goto L786
	} else {
		goto L791
	}
L789:
	;
	goto L790
L790:
	;
	v5565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v5565 != 0 {
		goto L786
	} else {
		goto L793
	}
L791:
	;
	v5564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v5564 != 0 {
		goto L786
	} else {
		goto L792
	}
L792:
	;
	goto L790
L793:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v5567 = m.ExcPending
	if v5567 != 0 {
		goto L3
	} else {
		goto L794
	}
L794:
	;
	F_XLogRegisterData(m, v60+int32(784), int32(20))
	mBase = m.M
	v5572 = m.ExcPending
	if v5572 != 0 {
		goto L3
	} else {
		goto L795
	}
L795:
	;
	v5573 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5520)+4)))
	F_XLogRegisterData(m, v5520, v5573)
	mBase = m.M
	v5575 = m.ExcPending
	if v5575 != 0 {
		goto L3
	} else {
		goto L796
	}
L796:
	;
	F_XLogRegisterBuffer(m, int32(0), v4643, int32(8))
	mBase = m.M
	v5579 = m.ExcPending
	if v5579 != 0 {
		goto L3
	} else {
		goto L797
	}
L797:
	;
	v5582 = F_XLogInsert(m, int32(16), int32(48))
	mBase = m.M
	v5583 = m.ExcPending
	if v5583 != 0 {
		goto L3
	} else {
		goto L798
	}
L798:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4642))) = base.I64_rotl(v5582, int64(32))
	goto L786
L799:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+792)) = uint16(v353)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+790)) = uint16(v344)
	v5598 = base.I32_rem_u_s(v4650, int32(3))
	v5599 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5520)+4)))
	v5602 = F_SpGistGetBuffer(m, l0, v5598, v5599+int32(4), v310)
	mBase = m.M
	v5603 = m.ExcPending
	if v5603 != 0 {
		goto L3
	} else {
		goto L800
	}
L800:
	;
	if v5602 < int32(0) {
		goto L802
	} else {
		goto L803
	}
L801:
	;
	if v5602 < int32(0) {
		goto L806
	} else {
		goto L807
	}
L802:
	;
	v5607 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[1]))
	v5613 = *(*int32)(unsafe.Add(mBase, uint32(v5607+(v5602^int32(-1))*int32(56))+16))
	v5622 = v5613
	goto L801
L803:
	;
	goto L804
L804:
	;
	v5615 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[2]))
	v5616 = int32(56)
	v5621 = *(*int32)(unsafe.Add(mBase, uint32(v5615+v5602*v5616-v5616)+16))
	v5622 = v5621
	goto L801
L805:
	;
	if v5622 == v4650 {
		goto L666
	} else {
		goto L809
	}
L806:
	;
	v5626 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v5632 = *(*int32)(unsafe.Add(mBase, uint32(v5626+(v5602^int32(-1))<<(uint(int32(2))%32))))
	v5640 = v5632
	goto L805
L807:
	;
	goto L808
L808:
	;
	v5634 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v5640 = v5634 + v5602<<(uint(int32(13))%32) + int32(-8192)
	goto L805
L809:
	;
	v5642 = int32(_a_F_spgdoinsert_4)
	v5644 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	v5645 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v5644 + v5645
	if v5602 == v341 {
		goto L810
	} else {
		goto L811
	}
L810:
	;
	v5652 = v5645
	goto L812
L811:
	;
	v5652 = int32(2)
	goto L812
L812:
	;
	v5653 = base.B2i32(v4643 == v341)
	if v4643 == v341 {
		goto L813
	} else {
		goto L814
	}
L813:
	;
	v5654 = int32(0)
	goto L815
L814:
	;
	v5654 = v5652
	goto L815
L815:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+789)) = uint8(v5654)
	v5656 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5520)+4)))
	v5658 = F_SpGistPageAddNewItem(m, v5640, v5520, v5656, int32(0))
	mBase = m.M
	v5659 = m.ExcPending
	if v5659 != 0 {
		goto L3
	} else {
		goto L816
	}
L816:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+786)) = uint16(v5658)
	F_MarkBufferDirty(m, v5602)
	mBase = m.M
	v5662 = m.ExcPending
	if v5662 != 0 {
		goto L3
	} else {
		goto L817
	}
L817:
	;
	F_saveNodeLink(m, v60+int32(444), v5622, v5658)
	mBase = m.M
	v5666 = m.ExcPending
	if v5666 != 0 {
		goto L3
	} else {
		goto L818
	}
L818:
	;
	v5667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v5667 == int32(1) {
		goto L820
	} else {
		goto L821
	}
L819:
	;
	F_PageIndexTupleDelete(m, v4642, v4696)
	mBase = m.M
	v5728 = m.ExcPending
	if v5728 != 0 {
		goto L3
	} else {
		goto L831
	}
L820:
	;
	v5674 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v5674))) = int32(67)
	v5680 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5674)+4)))
	v5682 = v5680 & int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v5674)+4)) = uint16(v5682)
	goto L825
L821:
	;
	goto L822
L822:
	;
	v5701 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v5701))) = int32(65)
	v5707 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5701)+4)))
	v5709 = v5707 & int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v5701)+4)) = uint16(v5709)
	goto L828
L823:
	;
	v5726 = v5674
	goto L819
L825:
	;
	goto L826
L826:
	;
	v5693 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5674)+10)) = uint16(v5693)
	*(*int32)(unsafe.Add(mBase, uint32(v5674)+6)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v5674)+12)) = v5693
	goto L823
L827:
	;
	v5726 = v5701
	goto L819
L828:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v5701)+10)) = uint16(v5658)
	*(*uint16)(unsafe.Add(mBase, uint32(v5701)+8)) = uint16(v5622)
	v5716 = int32(base.Ui32(v5622) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v5701)+6)) = uint16(v5716)
	v5718 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v5701)+12)) = v5718
	goto L827
L831:
	;
	v5729 = *(*int32)(unsafe.Add(mBase, uint32(v5726)))
	v5733 = F_PageAddItemExtended(m, v4642, v5726, int32(base.Ui32(v5729)>>(uint(int32(2))%32)), v4696, int32(0))
	mBase = m.M
	v5734 = m.ExcPending
	if v5734 != 0 {
		goto L3
	} else {
		goto L832
	}
L832:
	;
	if v5733 != v4696 {
		goto L665
	} else {
		goto L833
	}
L833:
	;
	v5736 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4642)+16)))
	v5740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v5740 != 0 {
		goto L834
	} else {
		goto L835
	}
L834:
	;
	v5741 = int32(4)
	goto L836
L835:
	;
	v5741 = int32(2)
	goto L836
L836:
	;
	v5742 = v4642 + v5736 + v5741
	v5743 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5742))))
	v5745 = v5743 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5742))) = uint16(v5745)
	F_MarkBufferDirty(m, v4643)
	mBase = m.M
	v5748 = m.ExcPending
	if v5748 != 0 {
		goto L3
	} else {
		goto L837
	}
L837:
	;
	v5749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5749)+118)))
	if v5750 != int32(112) {
		goto L838
	} else {
		goto L839
	}
L838:
	;
	v5798 = int32(_a_F_spgdoinsert_4)
	v5800 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v5800 - int32(1)
	if base.B2i32(v5602 == v4643) == int32(0) {
		goto L859
	} else {
		goto L860
	}
L839:
	;
	v5754 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[6]))
	if v5754 <= int32(0) {
		goto L840
	} else {
		goto L841
	}
L840:
	;
	v5757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v5757 != 0 {
		goto L838
	} else {
		goto L843
	}
L841:
	;
	goto L842
L842:
	;
	v5759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v5759 != 0 {
		goto L838
	} else {
		goto L845
	}
L843:
	;
	v5758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v5758 != 0 {
		goto L838
	} else {
		goto L844
	}
L844:
	;
	goto L842
L845:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v5761 = m.ExcPending
	if v5761 != 0 {
		goto L3
	} else {
		goto L846
	}
L846:
	;
	F_XLogRegisterBuffer(m, int32(0), v4643, int32(8))
	mBase = m.M
	v5765 = m.ExcPending
	if v5765 != 0 {
		goto L3
	} else {
		goto L847
	}
L847:
	;
	v5769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+788)))
	if v5769 != 0 {
		goto L848
	} else {
		goto L849
	}
L848:
	;
	v5770 = int32(14)
	goto L850
L849:
	;
	v5770 = int32(8)
	goto L850
L850:
	;
	F_XLogRegisterBuffer(m, int32(1), v5602, v5770)
	mBase = m.M
	v5772 = m.ExcPending
	if v5772 != 0 {
		goto L3
	} else {
		goto L851
	}
L851:
	;
	v5773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+789)))
	if v5773 == int32(2) {
		goto L852
	} else {
		goto L853
	}
L852:
	;
	F_XLogRegisterBuffer(m, int32(2), v341, int32(8))
	mBase = m.M
	v5779 = m.ExcPending
	if v5779 != 0 {
		goto L3
	} else {
		goto L855
	}
L853:
	;
	goto L854
L854:
	;
	F_XLogRegisterData(m, v60+int32(784), int32(20))
	mBase = m.M
	v5784 = m.ExcPending
	if v5784 != 0 {
		goto L3
	} else {
		goto L856
	}
L855:
	;
	goto L854
L856:
	;
	v5785 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5520)+4)))
	F_XLogRegisterData(m, v5520, v5785)
	mBase = m.M
	v5787 = m.ExcPending
	if v5787 != 0 {
		goto L3
	} else {
		goto L857
	}
L857:
	;
	v5790 = F_XLogInsert(m, int32(16), int32(48))
	mBase = m.M
	v5791 = m.ExcPending
	if v5791 != 0 {
		goto L3
	} else {
		goto L858
	}
L858:
	;
	v5793 = base.I64_rotl(v5790, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v5640))) = v5793
	*(*int64)(unsafe.Add(mBase, uint32(v340))) = v5793
	*(*int64)(unsafe.Add(mBase, uint32(v4642))) = v5793
	goto L838
L859:
	;
	if v5653 == int32(0) {
		goto L862
	} else {
		goto L863
	}
L860:
	;
	v5813 = v4643
	goto L861
L861:
	;
	v5817 = v5640
	v5818 = v5813
	v5819 = v5658
	v5821 = v5622
	goto L774
L862:
	;
	F_SpGistSetLastUsedPage(m, l0, v4643)
	mBase = m.M
	v5810 = m.ExcPending
	if v5810 != 0 {
		goto L3
	} else {
		goto L865
	}
L863:
	;
	goto L864
L864:
	;
	v5813 = v5602
	goto L861
L865:
	;
	F_UnlockReleaseBuffer(m, v4643)
	mBase = m.M
	v5812 = m.ExcPending
	if v5812 != 0 {
		goto L3
	} else {
		goto L866
	}
L866:
	;
	goto L864
L867:
	;
	v7056 = v5823
	v7060 = v5818
	goto L30
L868:
	;
	v5833 = *(*int32)(unsafe.Add(mBase, uint32(v60)+384))
	if base.Ui32(v5828) <= base.Ui32(v5833) {
		goto L663
	} else {
		goto L869
	}
L869:
	;
	v5836 = F_palloc_mul(m, int32(4), v5828)
	mBase = m.M
	v5837 = m.ExcPending
	if v5837 != 0 {
		goto L3
	} else {
		goto L870
	}
L870:
	;
	v5838 = int32(0)
	v5839 = *(*int32)(unsafe.Add(mBase, uint32(v60)+376))
	if v5838 < v5839 {
		goto L871
	} else {
		goto L872
	}
L871:
	;
	v5849 = v5838
	goto L874
L872:
	;
	v5926 = v5839
	goto L873
L873:
	;
	v5975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+360)))
	v5976 = *(*int64)(unsafe.Add(mBase, uint32(v60)+368))
	v5977 = F_spgFormInnerTuple(m, l1, v5975, v5976, v5926, v5836)
	mBase = m.M
	v5978 = m.ExcPending
	if v5978 != 0 {
		goto L3
	} else {
		goto L881
	}
L874:
	;
	v5902 = *(*int32)(unsafe.Add(mBase, uint32(v60)+380))
	if v5902 != 0 {
		goto L876
	} else {
		goto L877
	}
L875:
	;
	v5926 = v5916
	goto L873
L876:
	;
	v5906 = *(*int64)(unsafe.Add(mBase, uint32(v5902+v5849<<(uint(int32(3))%32))))
	v5908 = v5906
	goto L878
L877:
	;
	v5908 = int64(0)
	goto L878
L878:
	;
	v5911 = F_spgFormNodeTuple(m, l1, v5908, base.B2i32(v5902 == int32(0)))
	mBase = m.M
	v5912 = m.ExcPending
	if v5912 != 0 {
		goto L3
	} else {
		goto L879
	}
L879:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5836+v5849<<(uint(int32(2))%32)))) = v5911
	v5915 = v5849 + int32(1)
	v5916 = *(*int32)(unsafe.Add(mBase, uint32(v60)+376))
	if v5915 < v5916 {
		v5849 = v5915
		goto L874
	} else {
		goto L880
	}
L880:
	;
	goto L875
L881:
	;
	v5979 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5977)+4)))
	v5980 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4767)+4)))
	if base.Ui32(v5980) < base.Ui32(v5979) {
		goto L662
	} else {
		goto L882
	}
L882:
	;
	v5983 = *(*int32)(unsafe.Add(mBase, uint32(v4767)))
	v5988 = F_palloc_mul(m, int32(4), int32(base.Ui32(v5983)>>(uint(int32(3))%32))&int32(_a_F_spgdoinsert_22))
	mBase = m.M
	v5989 = m.ExcPending
	if v5989 != 0 {
		goto L3
	} else {
		goto L883
	}
L883:
	;
	v5990 = *(*int32)(unsafe.Add(mBase, uint32(v4767)))
	if v5990&int32(_a_F_spgdoinsert_21) == int32(0) {
		goto L885
	} else {
		goto L886
	}
L884:
	;
	v6132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+388)))
	v6133 = *(*int64)(unsafe.Add(mBase, uint32(v60)+392))
	v6134 = F_spgFormInnerTuple(m, l1, v6132, v6133, v6090, v5988)
	mBase = m.M
	v6135 = m.ExcPending
	if v6135 != 0 {
		goto L3
	} else {
		goto L891
	}
L885:
	;
	v6090 = int32(0)
	goto L884
L886:
	;
	goto L887
L887:
	;
	v6009 = v4767 + int32(base.Ui32(v5990)>>(uint(int32(16))%32)) + int32(8)
	v6010 = int32(0)
	goto L888
L888:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5988+v6010<<(uint(int32(2))%32)))) = v6009
	v6063 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6009)+6)))
	v6064 = int32(_a_F_spgdoinsert_22)
	v6068 = v6010 + int32(1)
	v6069 = *(*int32)(unsafe.Add(mBase, uint32(v4767)))
	v6073 = int32(base.Ui32(v6069)>>(uint(int32(3))%32)) & v6064
	if base.Ui32(v6068) < base.Ui32(v6073) {
		v6009 = v6009 + v6063&v6064
		v6010 = v6068
		goto L888
	} else {
		goto L890
	}
L889:
	;
	v6090 = v6073
	goto L884
L890:
	;
	goto L889
L891:
	;
	v6136 = *(*int32)(unsafe.Add(mBase, uint32(v6134)))
	v6139 = *(*int32)(unsafe.Add(mBase, uint32(v4767)))
	*(*int32)(unsafe.Add(mBase, uint32(v6134))) = v6136&int32(-5) | v6139&int32(4)
	v6144 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+788)) = uint8(v6144)
	if base.Ui32(v4690) <= base.Ui32(int32(1)) {
		goto L894
	} else {
		goto L895
	}
L892:
	;
	v6180 = int32(_a_F_spgdoinsert_4)
	v6182 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v6182 + int32(1)
	F_PageIndexTupleDelete(m, v4642, v4696)
	mBase = m.M
	v6187 = m.ExcPending
	if v6187 != 0 {
		goto L3
	} else {
		goto L906
	}
L893:
	;
	v6176 = F_SpGistGetBuffer(m, l0, v4694, v6172+int32(4), v310)
	mBase = m.M
	v6177 = m.ExcPending
	if v6177 != 0 {
		goto L3
	} else {
		goto L905
	}
L894:
	;
	v6148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6134)+4)))
	v6172 = v6148
	goto L893
L895:
	;
	goto L896
L896:
	;
	v6149 = int32(0)
	v6150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4642)+14)))
	v6151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4642)+12)))
	v6152 = v6150 - v6151
	if v6149 < v6152 {
		goto L898
	} else {
		goto L899
	}
L897:
	;
	v6159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4642)+16)))
	v6161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4642+v6159)+4)))
	if v6161 != 0 {
		goto L901
	} else {
		goto L902
	}
L898:
	;
	v6156 = v6152
	goto L900
L899:
	;
	v6156 = v6149
	goto L900
L900:
	;
	goto L897
L901:
	;
	v6162 = int32(20)
	goto L903
L902:
	;
	v6162 = int32(0)
	goto L903
L903:
	;
	v6164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4767)+4)))
	v6166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6134)+4)))
	v6167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5977)+4)))
	if base.Ui32(v6166+v6167+int32(4)) <= base.Ui32(v6156+v6162+v6164) {
		v6179 = v6149
		goto L892
	} else {
		goto L904
	}
L904:
	;
	v6172 = v6166
	goto L893
L905:
	;
	v6179 = v6176
	goto L892
L906:
	;
	v6188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5977)+4)))
	v6190 = F_PageAddItemExtended(m, v4642, v5977, v6188, v4696, int32(0))
	mBase = m.M
	v6191 = m.ExcPending
	if v6191 != 0 {
		goto L3
	} else {
		goto L907
	}
L907:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+784)) = uint16(v6190)
	if v6190 != v4696 {
		goto L661
	} else {
		goto L908
	}
L908:
	;
	if v6179 == int32(0) {
		goto L910
	} else {
		goto L911
	}
L909:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+789)) = uint8(v6247)
	v6251 = *(*int32)(unsafe.Add(mBase, uint32(v60)+384))
	v6252 = *(*int32)(unsafe.Add(mBase, uint32(v5977)))
	v6256 = int32(base.Ui32(v6252)>>(uint(int32(3))%32)) & int32(_a_F_spgdoinsert_22)
	if v6256 != 0 {
		goto L925
	} else {
		goto L926
	}
L910:
	;
	v6196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6134)+4)))
	v6198 = F_SpGistPageAddNewItem(m, v4642, v6134, v6196, int32(0))
	mBase = m.M
	v6199 = m.ExcPending
	if v6199 != 0 {
		goto L3
	} else {
		goto L913
	}
L911:
	;
	goto L912
L912:
	;
	if v6179 < int32(0) {
		goto L915
	} else {
		goto L916
	}
L913:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+786)) = uint16(v6198)
	v6247 = int32(1)
	v6248 = v4650
	v6249 = v6198
	goto L909
L914:
	;
	v6221 = int32(0)
	if v6179 < v6221 {
		goto L919
	} else {
		goto L920
	}
L915:
	;
	v6205 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[1]))
	v6211 = *(*int32)(unsafe.Add(mBase, uint32(v6205+(v6179^int32(-1))*int32(56))+16))
	v6220 = v6211
	goto L914
L916:
	;
	goto L917
L917:
	;
	v6213 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[2]))
	v6214 = int32(56)
	v6219 = *(*int32)(unsafe.Add(mBase, uint32(v6213+v6179*v6214-v6214)+16))
	v6220 = v6219
	goto L914
L918:
	;
	v6240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6134)+4)))
	v6242 = F_SpGistPageAddNewItem(m, v6239, v6134, v6240, int32(0))
	mBase = m.M
	v6243 = m.ExcPending
	if v6243 != 0 {
		goto L3
	} else {
		goto L922
	}
L919:
	;
	v6225 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v6231 = *(*int32)(unsafe.Add(mBase, uint32(v6225+(v6179^int32(-1))<<(uint(int32(2))%32))))
	v6239 = v6231
	goto L918
L920:
	;
	goto L921
L921:
	;
	v6233 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v6239 = v6233 + v6179<<(uint(int32(13))%32) + int32(-8192)
	goto L918
L922:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+786)) = uint16(v6242)
	F_MarkBufferDirty(m, v6179)
	mBase = m.M
	v6246 = m.ExcPending
	if v6246 != 0 {
		goto L3
	} else {
		goto L923
	}
L923:
	;
	v6247 = v6221
	v6248 = v6220
	v6249 = v6242
	goto L909
L924:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v6270)+4)) = uint16(v6249)
	*(*uint16)(unsafe.Add(mBase, uint32(v6270)+2)) = uint16(v6248)
	v6403 = int32(base.Ui32(v6248) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v6270))) = uint16(v6403)
	v6405 = *(*int32)(unsafe.Add(mBase, uint32(v60)+384))
	v6406 = *(*int32)(unsafe.Add(mBase, uint32(v4701)))
	v6409 = v4642 + v6406&int32(_a_F_spgdoinsert_6)
	v6410 = *(*int32)(unsafe.Add(mBase, uint32(v6409)))
	v6414 = int32(base.Ui32(v6410)>>(uint(int32(3))%32)) & int32(_a_F_spgdoinsert_22)
	if v6414 != 0 {
		goto L936
	} else {
		goto L937
	}
L925:
	;
	v6270 = v5977 + int32(base.Ui32(v6252)>>(uint(int32(16))%32)) + int32(8)
	v6271 = int32(0)
	goto L928
L926:
	;
	goto L927
L927:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6388 = m.ExcPending
	if v6388 != 0 {
		goto L3
	} else {
		goto L932
	}
L928:
	;
	if v6271 == v6251 {
		goto L924
	} else {
		goto L930
	}
L929:
	;
	goto L927
L930:
	;
	v6321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6270)+6)))
	v6326 = v6271 + int32(1)
	if v6326 != v6256 {
		v6270 = v6270 + v6321&int32(_a_F_spgdoinsert_22)
		v6271 = v6326
		goto L928
	} else {
		goto L931
	}
L931:
	;
	goto L929
L932:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+112)) = v6251
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_33), v60+int32(112))
	mBase = m.M
	v6394 = m.ExcPending
	if v6394 != 0 {
		goto L3
	} else {
		goto L933
	}
L933:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(68), int32(_a_F_spgdoinsert_34))
	mBase = m.M
	v6399 = m.ExcPending
	if v6399 != 0 {
		goto L3
	} else {
		goto L934
	}
L934:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L935:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v6428)+4)) = uint16(v6249)
	*(*uint16)(unsafe.Add(mBase, uint32(v6428)+2)) = uint16(v6248)
	*(*uint16)(unsafe.Add(mBase, uint32(v6428))) = uint16(v6403)
	F_MarkBufferDirty(m, v4643)
	mBase = m.M
	v6562 = m.ExcPending
	if v6562 != 0 {
		goto L3
	} else {
		goto L946
	}
L936:
	;
	v6428 = v6409 + int32(base.Ui32(v6410)>>(uint(int32(16))%32)) + int32(8)
	v6429 = int32(0)
	goto L939
L937:
	;
	goto L938
L938:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6546 = m.ExcPending
	if v6546 != 0 {
		goto L3
	} else {
		goto L943
	}
L939:
	;
	if v6429 == v6405 {
		goto L935
	} else {
		goto L941
	}
L940:
	;
	goto L938
L941:
	;
	v6479 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6428)+6)))
	v6484 = v6429 + int32(1)
	if v6484 != v6414 {
		v6428 = v6428 + v6479&int32(_a_F_spgdoinsert_22)
		v6429 = v6484
		goto L939
	} else {
		goto L942
	}
L942:
	;
	goto L940
L943:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+128)) = v6405
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_33), v60+int32(128))
	mBase = m.M
	v6552 = m.ExcPending
	if v6552 != 0 {
		goto L3
	} else {
		goto L944
	}
L944:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(68), int32(_a_F_spgdoinsert_34))
	mBase = m.M
	v6557 = m.ExcPending
	if v6557 != 0 {
		goto L3
	} else {
		goto L945
	}
L945:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L946:
	;
	v6563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v6564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6563)+118)))
	if v6564 != int32(112) {
		goto L949
	} else {
		goto L950
	}
L947:
	;
	v6657 = int32(0)
	v6659 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v6659 == v6657 {
		goto L673
	} else {
		goto L978
	}
L948:
	;
	F_SpGistSetLastUsedPage(m, l0, v6179)
	mBase = m.M
	v6653 = m.ExcPending
	if v6653 != 0 {
		goto L3
	} else {
		goto L976
	}
L949:
	;
	v6643 = int32(_a_F_spgdoinsert_4)
	v6645 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v6645 - int32(1)
	if v6179 == int32(0) {
		goto L947
	} else {
		goto L975
	}
L950:
	;
	v6568 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[6]))
	if v6568 <= int32(0) {
		goto L951
	} else {
		goto L952
	}
L951:
	;
	v6571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v6571 != 0 {
		goto L949
	} else {
		goto L954
	}
L952:
	;
	goto L953
L953:
	;
	v6573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v6573 != 0 {
		goto L949
	} else {
		goto L956
	}
L954:
	;
	v6572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v6572 != 0 {
		goto L949
	} else {
		goto L955
	}
L955:
	;
	goto L953
L956:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v6575 = m.ExcPending
	if v6575 != 0 {
		goto L3
	} else {
		goto L957
	}
L957:
	;
	F_XLogRegisterData(m, v60+int32(784), int32(6))
	mBase = m.M
	v6580 = m.ExcPending
	if v6580 != 0 {
		goto L3
	} else {
		goto L958
	}
L958:
	;
	v6581 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6409)+4)))
	F_XLogRegisterData(m, v6409, v6581)
	mBase = m.M
	v6583 = m.ExcPending
	if v6583 != 0 {
		goto L3
	} else {
		goto L959
	}
L959:
	;
	v6584 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6134)+4)))
	F_XLogRegisterData(m, v6134, v6584)
	mBase = m.M
	v6586 = m.ExcPending
	if v6586 != 0 {
		goto L3
	} else {
		goto L960
	}
L960:
	;
	F_XLogRegisterBuffer(m, int32(0), v4643, int32(8))
	mBase = m.M
	v6590 = m.ExcPending
	if v6590 != 0 {
		goto L3
	} else {
		goto L961
	}
L961:
	;
	if v6179 != 0 {
		goto L962
	} else {
		goto L963
	}
L962:
	;
	v6594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+788)))
	if v6594 != 0 {
		goto L965
	} else {
		goto L966
	}
L963:
	;
	goto L964
L964:
	;
	v6632 = F_XLogInsert(m, int32(16), int32(64))
	mBase = m.M
	v6633 = m.ExcPending
	if v6633 != 0 {
		goto L3
	} else {
		goto L974
	}
L965:
	;
	v6595 = int32(14)
	goto L967
L966:
	;
	v6595 = int32(8)
	goto L967
L967:
	;
	F_XLogRegisterBuffer(m, int32(1), v6179, v6595)
	mBase = m.M
	v6597 = m.ExcPending
	if v6597 != 0 {
		goto L3
	} else {
		goto L968
	}
L968:
	;
	v6600 = F_XLogInsert(m, int32(16), int32(64))
	mBase = m.M
	v6601 = m.ExcPending
	if v6601 != 0 {
		goto L3
	} else {
		goto L969
	}
L969:
	;
	v6603 = base.I64_rotl(v6600, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v4642))) = v6603
	if v6179 < int32(0) {
		goto L971
	} else {
		goto L972
	}
L970:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6622))) = v6603
	v6624 = int32(_a_F_spgdoinsert_4)
	v6626 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v6626 - int32(1)
	goto L948
L971:
	;
	v6608 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v6614 = *(*int32)(unsafe.Add(mBase, uint32(v6608+(v6179^int32(-1))<<(uint(int32(2))%32))))
	v6622 = v6614
	goto L970
L972:
	;
	goto L973
L973:
	;
	v6616 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v6622 = v6616 + v6179<<(uint(int32(13))%32) + int32(-8192)
	goto L970
L974:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4642))) = base.I64_rotl(v6632, int64(32))
	v6637 = int32(_a_F_spgdoinsert_4)
	v6639 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v6639 - int32(1)
	goto L947
L975:
	;
	goto L948
L976:
	;
	F_UnlockReleaseBuffer(m, v6179)
	mBase = m.M
	v6655 = m.ExcPending
	if v6655 != 0 {
		goto L3
	} else {
		goto L977
	}
L977:
	;
	goto L947
L978:
	;
	v7056 = v6657
	v7060 = v4643
	goto L30
L979:
	;
	v6666 = *(*int32)(unsafe.Add(mBase, uint32(v60)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v6666
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_35), v60+int32(16))
	mBase = m.M
	v6672 = m.ExcPending
	if v6672 != 0 {
		goto L3
	} else {
		goto L980
	}
L980:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(2313), int32(_a_F_spgdoinsert_18))
	mBase = m.M
	v6677 = m.ExcPending
	if v6677 != 0 {
		goto L3
	} else {
		goto L981
	}
L981:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L982:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_36), int32(0))
	mBase = m.M
	v6685 = m.ExcPending
	if v6685 != 0 {
		goto L3
	} else {
		goto L983
	}
L983:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(2290), int32(_a_F_spgdoinsert_18))
	mBase = m.M
	v6690 = m.ExcPending
	if v6690 != 0 {
		goto L3
	} else {
		goto L984
	}
L984:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L985:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_37), int32(0))
	mBase = m.M
	v6698 = m.ExcPending
	if v6698 != 0 {
		goto L3
	} else {
		goto L986
	}
L986:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(90), int32(_a_F_spgdoinsert_38))
	mBase = m.M
	v6703 = m.ExcPending
	if v6703 != 0 {
		goto L3
	} else {
		goto L987
	}
L987:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L988:
	;
	v6708 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5520)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+80)) = v6708
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_8), v60+int32(80))
	mBase = m.M
	v6714 = m.ExcPending
	if v6714 != 0 {
		goto L3
	} else {
		goto L989
	}
L989:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1549), int32(_a_F_spgdoinsert_39))
	mBase = m.M
	v6719 = m.ExcPending
	if v6719 != 0 {
		goto L3
	} else {
		goto L990
	}
L990:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L991:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_40), int32(0))
	mBase = m.M
	v6727 = m.ExcPending
	if v6727 != 0 {
		goto L3
	} else {
		goto L992
	}
L992:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1584), int32(_a_F_spgdoinsert_39))
	mBase = m.M
	v6732 = m.ExcPending
	if v6732 != 0 {
		goto L3
	} else {
		goto L993
	}
L993:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L994:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_41), int32(0))
	mBase = m.M
	v6740 = m.ExcPending
	if v6740 != 0 {
		goto L3
	} else {
		goto L995
	}
L995:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1612), int32(_a_F_spgdoinsert_39))
	mBase = m.M
	v6745 = m.ExcPending
	if v6745 != 0 {
		goto L3
	} else {
		goto L996
	}
L996:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L997:
	;
	v6750 = *(*int32)(unsafe.Add(mBase, uint32(v5726)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+64)) = int32(base.Ui32(v6750) >> (uint(int32(2)) % 32))
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_8), v60-int32(-64))
	mBase = m.M
	v6758 = m.ExcPending
	if v6758 != 0 {
		goto L3
	} else {
		goto L998
	}
L998:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1657), int32(_a_F_spgdoinsert_39))
	mBase = m.M
	v6763 = m.ExcPending
	if v6763 != 0 {
		goto L3
	} else {
		goto L999
	}
L999:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1000:
	;
	v6768 = *(*int32)(unsafe.Add(mBase, uint32(v60)+376))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+96)) = v6768
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_42), v60+int32(96))
	mBase = m.M
	v6774 = m.ExcPending
	if v6774 != 0 {
		goto L3
	} else {
		goto L1001
	}
L1001:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1732), int32(_a_F_spgdoinsert_43))
	mBase = m.M
	v6779 = m.ExcPending
	if v6779 != 0 {
		goto L3
	} else {
		goto L1002
	}
L1002:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1003:
	;
	v6784 = *(*int32)(unsafe.Add(mBase, uint32(v60)+384))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+160)) = v6784
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_44), v60+int32(160))
	mBase = m.M
	v6790 = m.ExcPending
	if v6790 != 0 {
		goto L3
	} else {
		goto L1004
	}
L1004:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1737), int32(_a_F_spgdoinsert_43))
	mBase = m.M
	v6795 = m.ExcPending
	if v6795 != 0 {
		goto L3
	} else {
		goto L1005
	}
L1005:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1006:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_45), int32(0))
	mBase = m.M
	v6803 = m.ExcPending
	if v6803 != 0 {
		goto L3
	} else {
		goto L1007
	}
L1007:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1764), int32(_a_F_spgdoinsert_43))
	mBase = m.M
	v6808 = m.ExcPending
	if v6808 != 0 {
		goto L3
	} else {
		goto L1008
	}
L1008:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1009:
	;
	v6813 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5977)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+144)) = v6813
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_8), v60+int32(144))
	mBase = m.M
	v6819 = m.ExcPending
	if v6819 != 0 {
		goto L3
	} else {
		goto L1010
	}
L1010:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1820), int32(_a_F_spgdoinsert_43))
	mBase = m.M
	v6824 = m.ExcPending
	if v6824 != 0 {
		goto L3
	} else {
		goto L1011
	}
L1011:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1012:
	;
	goto L63
L1013:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v6899 = m.ExcPending
	if v6899 != 0 {
		goto L3
	} else {
		goto L1014
	}
L1014:
	;
	v6900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = int32(_a_F_spgdoinsert_30)
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v6900 + int32(4)
	F_errmsg(m, int32(_a_F_spgdoinsert_31), v60)
	mBase = m.M
	v6909 = m.ExcPending
	if v6909 != 0 {
		goto L3
	} else {
		goto L1015
	}
L1015:
	;
	F_errhint(m, int32(_a_F_spgdoinsert_32), int32(0))
	mBase = m.M
	v6913 = m.ExcPending
	if v6913 != 0 {
		goto L3
	} else {
		goto L1016
	}
L1016:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(2000), int32(_a_F_spgdoinsert_18))
	mBase = m.M
	v6918 = m.ExcPending
	if v6918 != 0 {
		goto L3
	} else {
		goto L1017
	}
L1017:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1018:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+48)) = v4907
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_33), v60+int32(48))
	mBase = m.M
	v6985 = m.ExcPending
	if v6985 != 0 {
		goto L3
	} else {
		goto L1019
	}
L1019:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1486), int32(_a_F_spgdoinsert_46))
	mBase = m.M
	v6990 = m.ExcPending
	if v6990 != 0 {
		goto L3
	} else {
		goto L1020
	}
L1020:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1021:
	;
	v7120 = v7056
	v7124 = int32(0)
	v7130 = v341
	goto L29
L1022:
	;
	goto L1023
L1023:
	;
	F_SpGistSetLastUsedPage(m, l0, v7060)
	mBase = m.M
	v7110 = m.ExcPending
	if v7110 != 0 {
		goto L3
	} else {
		goto L1024
	}
L1024:
	;
	F_UnlockReleaseBuffer(m, v7060)
	mBase = m.M
	v7112 = m.ExcPending
	if v7112 != 0 {
		goto L3
	} else {
		goto L1025
	}
L1025:
	;
	v7120 = v7056
	v7124 = v7060
	v7130 = v341
	goto L29
L1026:
	;
	F_SpGistSetLastUsedPage(m, l0, v7130)
	mBase = m.M
	v7177 = m.ExcPending
	if v7177 != 0 {
		goto L3
	} else {
		goto L1029
	}
L1027:
	;
	goto L1028
L1028:
	;
	v7181 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v7181 == int32(0) {
		v7193 = v7120
		goto L28
	} else {
		goto L1031
	}
L1029:
	;
	F_UnlockReleaseBuffer(m, v7130)
	mBase = m.M
	v7179 = m.ExcPending
	if v7179 != 0 {
		goto L3
	} else {
		goto L1030
	}
L1030:
	;
	goto L1028
L1031:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v7185 = m.ExcPending
	if v7185 != 0 {
		goto L3
	} else {
		goto L1032
	}
L1032:
	;
	v7193 = v7120
	goto L28
L1033:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
