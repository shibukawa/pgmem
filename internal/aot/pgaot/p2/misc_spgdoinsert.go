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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v444 int32
	_ = v444
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
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int64
	_ = v640
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v650 int64
	_ = v650
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v671 int32
	_ = v671
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v780 int32
	_ = v780
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v891 int32
	_ = v891
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v955 int32
	_ = v955
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v998 int32
	_ = v998
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
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
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1319 int32
	_ = v1319
	var v1325 int32
	_ = v1325
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
	var v1341 int32
	_ = v1341
	var v1344 int64
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int64
	_ = v1346
	var v1347 int64
	_ = v1347
	var v1351 int64
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1464 int32
	_ = v1464
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1487 int32
	_ = v1487
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1556 int32
	_ = v1556
	var v1562 int32
	_ = v1562
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1574 int32
	_ = v1574
	var v1581 int32
	_ = v1581
	var v1586 int32
	_ = v1586
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1599 int32
	_ = v1599
	var v1604 int32
	_ = v1604
	var v1611 int32
	_ = v1611
	var v1621 int32
	_ = v1621
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1703 int32
	_ = v1703
	var v1709 int32
	_ = v1709
	var v1718 int32
	_ = v1718
	var v1725 int32
	_ = v1725
	var v1730 int32
	_ = v1730
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1743 int32
	_ = v1743
	var v1748 int32
	_ = v1748
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1770 int32
	_ = v1770
	var v1776 int32
	_ = v1776
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1853 int32
	_ = v1853
	var v1860 int32
	_ = v1860
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1880 int64
	_ = v1880
	var v1887 int32
	_ = v1887
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1907 int32
	_ = v1907
	var v1917 int32
	_ = v1917
	var v1919 int32
	_ = v1919
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1970 int32
	_ = v1970
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1988 int32
	_ = v1988
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2008 int32
	_ = v2008
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2016 int32
	_ = v2016
	var v2028 int32
	_ = v2028
	var v2030 int32
	_ = v2030
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2081 int32
	_ = v2081
	var v2088 int32
	_ = v2088
	var v2091 int32
	_ = v2091
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2112 int32
	_ = v2112
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2126 int32
	_ = v2126
	var v2128 int32
	_ = v2128
	var v2146 int32
	_ = v2146
	var v2170 int32
	_ = v2170
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2178 int32
	_ = v2178
	var v2187 int32
	_ = v2187
	var v2237 int32
	_ = v2237
	var v2240 int32
	_ = v2240
	var v2303 int32
	_ = v2303
	var v2305 int32
	_ = v2305
	var v2315 int32
	_ = v2315
	var v2362 int32
	_ = v2362
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2373 int32
	_ = v2373
	var v2376 int32
	_ = v2376
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2396 int32
	_ = v2396
	var v2443 int32
	_ = v2443
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2527 int32
	_ = v2527
	var v2531 int32
	_ = v2531
	var v2580 int32
	_ = v2580
	var v2584 int32
	_ = v2584
	var v2602 int32
	_ = v2602
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2630 int32
	_ = v2630
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2638 int32
	_ = v2638
	var v2647 int32
	_ = v2647
	var v2698 int32
	_ = v2698
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2710 int32
	_ = v2710
	var v2711 int32
	_ = v2711
	var v2721 int32
	_ = v2721
	var v2766 int32
	_ = v2766
	var v2767 int32
	_ = v2767
	var v2768 int32
	_ = v2768
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2840 int32
	_ = v2840
	var v2841 int32
	_ = v2841
	var v2845 int32
	_ = v2845
	var v2846 int32
	_ = v2846
	var v2905 int32
	_ = v2905
	var v2914 int32
	_ = v2914
	var v2962 int32
	_ = v2962
	var v2963 int32
	_ = v2963
	var v2965 int32
	_ = v2965
	var v2968 int32
	_ = v2968
	var v2970 int32
	_ = v2970
	var v2972 int32
	_ = v2972
	var v2973 int32
	_ = v2973
	var v2975 int32
	_ = v2975
	var v2976 int32
	_ = v2976
	var v2984 int32
	_ = v2984
	var v3039 int32
	_ = v3039
	var v3045 int32
	_ = v3045
	var v3049 int32
	_ = v3049
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3058 int32
	_ = v3058
	var v3061 int32
	_ = v3061
	var v3063 int32
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3066 int32
	_ = v3066
	var v3068 int32
	_ = v3068
	var v3070 int32
	_ = v3070
	var v3075 int32
	_ = v3075
	var v3077 int32
	_ = v3077
	var v3078 int32
	_ = v3078
	var v3080 int32
	_ = v3080
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3086 int32
	_ = v3086
	var v3089 int32
	_ = v3089
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3095 int32
	_ = v3095
	var v3096 int32
	_ = v3096
	var v3101 int32
	_ = v3101
	var v3102 int32
	_ = v3102
	var v3110 int32
	_ = v3110
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3115 int32
	_ = v3115
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
	var v3124 int32
	_ = v3124
	var v3130 int32
	_ = v3130
	var v3132 int32
	_ = v3132
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3142 int32
	_ = v3142
	var v3145 int32
	_ = v3145
	var v3146 int32
	_ = v3146
	var v3147 int32
	_ = v3147
	var v3156 int32
	_ = v3156
	var v3158 int32
	_ = v3158
	var v3170 int32
	_ = v3170
	var v3203 int32
	_ = v3203
	var v3206 int32
	_ = v3206
	var v3207 int32
	_ = v3207
	var v3209 int32
	_ = v3209
	var v3211 int32
	_ = v3211
	var v3213 int32
	_ = v3213
	var v3215 int32
	_ = v3215
	var v3217 int32
	_ = v3217
	var v3218 int32
	_ = v3218
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3231 int32
	_ = v3231
	var v3243 int32
	_ = v3243
	var v3281 int32
	_ = v3281
	var v3282 int32
	_ = v3282
	var v3284 int32
	_ = v3284
	var v3285 int32
	_ = v3285
	var v3286 int32
	_ = v3286
	var v3288 int32
	_ = v3288
	var v3291 int32
	_ = v3291
	var v3292 int32
	_ = v3292
	var v3294 int32
	_ = v3294
	var v3295 int32
	_ = v3295
	var v3305 int32
	_ = v3305
	var v3311 int32
	_ = v3311
	var v3313 int32
	_ = v3313
	var v3319 int32
	_ = v3319
	var v3320 int32
	_ = v3320
	var v3321 int32
	_ = v3321
	var v3322 int32
	_ = v3322
	var v3323 int32
	_ = v3323
	var v3326 int32
	_ = v3326
	var v3327 int32
	_ = v3327
	var v3328 int32
	_ = v3328
	var v3337 int32
	_ = v3337
	var v3342 int32
	_ = v3342
	var v3343 int32
	_ = v3343
	var v3384 int32
	_ = v3384
	var v3387 int32
	_ = v3387
	var v3388 int32
	_ = v3388
	var v3390 int32
	_ = v3390
	var v3392 int32
	_ = v3392
	var v3394 int32
	_ = v3394
	var v3396 int32
	_ = v3396
	var v3398 int32
	_ = v3398
	var v3399 int32
	_ = v3399
	var v3401 int32
	_ = v3401
	var v3402 int32
	_ = v3402
	var v3415 int32
	_ = v3415
	var v3416 int32
	_ = v3416
	var v3457 int32
	_ = v3457
	var v3464 int32
	_ = v3464
	var v3468 int32
	_ = v3468
	var v3473 int32
	_ = v3473
	var v3477 int32
	_ = v3477
	var v3481 int32
	_ = v3481
	var v3486 int32
	_ = v3486
	var v3490 int32
	_ = v3490
	var v3494 int32
	_ = v3494
	var v3499 int32
	_ = v3499
	var v3517 int32
	_ = v3517
	var v3529 int32
	_ = v3529
	var v3556 int32
	_ = v3556
	var v3557 int32
	_ = v3557
	var v3558 int32
	_ = v3558
	var v3571 int32
	_ = v3571
	var v3577 int32
	_ = v3577
	var v3619 int32
	_ = v3619
	var v3622 int32
	_ = v3622
	var v3624 int32
	_ = v3624
	var v3627 int32
	_ = v3627
	var v3632 int32
	_ = v3632
	var v3634 int32
	_ = v3634
	var v3637 int32
	_ = v3637
	var v3642 int32
	_ = v3642
	var v3644 int32
	_ = v3644
	var v3647 int32
	_ = v3647
	var v3652 int32
	_ = v3652
	var v3654 int32
	_ = v3654
	var v3656 int32
	_ = v3656
	var v3657 int32
	_ = v3657
	var v3659 int32
	_ = v3659
	var v3667 int32
	_ = v3667
	var v3722 int32
	_ = v3722
	var v3726 int32
	_ = v3726
	var v3773 int32
	_ = v3773
	var v3775 int32
	_ = v3775
	var v3777 int32
	_ = v3777
	var v3780 int32
	_ = v3780
	var v3797 int32
	_ = v3797
	var v3799 int32
	_ = v3799
	var v3811 int32
	_ = v3811
	var v3837 int32
	_ = v3837
	var v3841 int32
	_ = v3841
	var v3842 int32
	_ = v3842
	var v3843 int32
	_ = v3843
	var v3845 int32
	_ = v3845
	var v3849 int32
	_ = v3849
	var v3850 int32
	_ = v3850
	var v3853 int32
	_ = v3853
	var v3855 int32
	_ = v3855
	var v3857 int32
	_ = v3857
	var v3867 int32
	_ = v3867
	var v3872 int32
	_ = v3872
	var v3878 int32
	_ = v3878
	var v3880 int32
	_ = v3880
	var v3886 int32
	_ = v3886
	var v3890 int32
	_ = v3890
	var v3891 int32
	_ = v3891
	var v3892 int32
	_ = v3892
	var v3895 int32
	_ = v3895
	var v3902 int32
	_ = v3902
	var v3907 int32
	_ = v3907
	var v3908 int32
	_ = v3908
	var v3909 int32
	_ = v3909
	var v3914 int32
	_ = v3914
	var v3918 int32
	_ = v3918
	var v3923 int32
	_ = v3923
	var v3925 int32
	_ = v3925
	var v3926 int32
	_ = v3926
	var v3936 int32
	_ = v3936
	var v3941 int32
	_ = v3941
	var v3984 int32
	_ = v3984
	var v3985 int32
	_ = v3985
	var v3986 int32
	_ = v3986
	var v3987 int32
	_ = v3987
	var v3988 int32
	_ = v3988
	var v3989 int32
	_ = v3989
	var v3993 int32
	_ = v3993
	var v3999 int32
	_ = v3999
	var v4001 int32
	_ = v4001
	var v4007 int32
	_ = v4007
	var v4008 int32
	_ = v4008
	var v4009 int32
	_ = v4009
	var v4011 int32
	_ = v4011
	var v4014 int32
	_ = v4014
	var v4015 int32
	_ = v4015
	var v4018 int32
	_ = v4018
	var v4021 int32
	_ = v4021
	var v4026 int32
	_ = v4026
	var v4029 int32
	_ = v4029
	var v4031 int32
	_ = v4031
	var v4040 int32
	_ = v4040
	var v4046 int32
	_ = v4046
	var v4048 int32
	_ = v4048
	var v4054 int32
	_ = v4054
	var v4055 int32
	_ = v4055
	var v4060 int32
	_ = v4060
	var v4064 int32
	_ = v4064
	var v4065 int32
	_ = v4065
	var v4067 int32
	_ = v4067
	var v4071 int32
	_ = v4071
	var v4073 int32
	_ = v4073
	var v4074 int32
	_ = v4074
	var v4076 int32
	_ = v4076
	var v4077 int32
	_ = v4077
	var v4078 int32
	_ = v4078
	var v4079 int32
	_ = v4079
	var v4080 int32
	_ = v4080
	var v4083 int32
	_ = v4083
	var v4085 int32
	_ = v4085
	var v4098 int32
	_ = v4098
	var v4141 int32
	_ = v4141
	var v4145 int32
	_ = v4145
	var v4146 int32
	_ = v4146
	var v4147 int32
	_ = v4147
	var v4149 int32
	_ = v4149
	var v4150 int32
	_ = v4150
	var v4151 int32
	_ = v4151
	var v4154 int32
	_ = v4154
	var v4156 int32
	_ = v4156
	var v4161 int32
	_ = v4161
	var v4167 int32
	_ = v4167
	var v4170 int32
	_ = v4170
	var v4174 int32
	_ = v4174
	var v4179 int32
	_ = v4179
	var v4185 int32
	_ = v4185
	var v4187 int32
	_ = v4187
	var v4193 int32
	_ = v4193
	var v4194 int32
	_ = v4194
	var v4198 int32
	_ = v4198
	var v4204 int32
	_ = v4204
	var v4206 int32
	_ = v4206
	var v4212 int32
	_ = v4212
	var v4213 int32
	_ = v4213
	var v4215 int32
	_ = v4215
	var v4216 int32
	_ = v4216
	var v4219 int32
	_ = v4219
	var v4222 int32
	_ = v4222
	var v4224 int32
	_ = v4224
	var v4229 int32
	_ = v4229
	var v4235 int32
	_ = v4235
	var v4238 int32
	_ = v4238
	var v4242 int32
	_ = v4242
	var v4247 int32
	_ = v4247
	var v4253 int32
	_ = v4253
	var v4255 int32
	_ = v4255
	var v4261 int32
	_ = v4261
	var v4265 int32
	_ = v4265
	var v4266 int32
	_ = v4266
	var v4267 int32
	_ = v4267
	var v4270 int32
	_ = v4270
	var v4272 int32
	_ = v4272
	var v4274 int32
	_ = v4274
	var v4277 int32
	_ = v4277
	var v4278 int32
	_ = v4278
	var v4282 int32
	_ = v4282
	var v4289 int32
	_ = v4289
	var v4290 int32
	_ = v4290
	var v4296 int32
	_ = v4296
	var v4301 int32
	_ = v4301
	var v4304 int32
	_ = v4304
	var v4305 int32
	_ = v4305
	var v4306 int32
	_ = v4306
	var v4307 int32
	_ = v4307
	var v4308 int32
	_ = v4308
	var v4310 int32
	_ = v4310
	var v4311 int32
	_ = v4311
	var v4312 int32
	_ = v4312
	var v4316 int32
	_ = v4316
	var v4319 int32
	_ = v4319
	var v4320 int32
	_ = v4320
	var v4321 int32
	_ = v4321
	var v4323 int32
	_ = v4323
	var v4329 int32
	_ = v4329
	var v4330 int32
	_ = v4330
	var v4334 int32
	_ = v4334
	var v4335 int32
	_ = v4335
	var v4339 int32
	_ = v4339
	var v4340 int32
	_ = v4340
	var v4342 int32
	_ = v4342
	var v4343 int32
	_ = v4343
	var v4345 int32
	_ = v4345
	var v4348 int32
	_ = v4348
	var v4352 int32
	_ = v4352
	var v4353 int32
	_ = v4353
	var v4355 int32
	_ = v4355
	var v4359 int32
	_ = v4359
	var v4360 int32
	_ = v4360
	var v4362 int32
	_ = v4362
	var v4366 int32
	_ = v4366
	var v4367 int32
	_ = v4367
	var v4369 int32
	_ = v4369
	var v4370 int32
	_ = v4370
	var v4377 int32
	_ = v4377
	var v4380 int64
	_ = v4380
	var v4381 int32
	_ = v4381
	var v4385 int32
	_ = v4385
	var v4391 int32
	_ = v4391
	var v4393 int32
	_ = v4393
	var v4399 int32
	_ = v4399
	var v4406 int32
	_ = v4406
	var v4412 int32
	_ = v4412
	var v4414 int32
	_ = v4414
	var v4420 int32
	_ = v4420
	var v4421 int64
	_ = v4421
	var v4430 int32
	_ = v4430
	var v4431 int32
	_ = v4431
	var v4434 int32
	_ = v4434
	var v4437 int32
	_ = v4437
	var v4444 int32
	_ = v4444
	var v4446 int32
	_ = v4446
	var v4451 int32
	_ = v4451
	var v4453 int32
	_ = v4453
	var v4455 int32
	_ = v4455
	var v4457 int32
	_ = v4457
	var v4460 int32
	_ = v4460
	var v4473 int32
	_ = v4473
	var v4475 int32
	_ = v4475
	var v4483 int32
	_ = v4483
	var v4514 int32
	_ = v4514
	var v4517 int32
	_ = v4517
	var v4530 int32
	_ = v4530
	var v4532 int32
	_ = v4532
	var v4533 int32
	_ = v4533
	var v4540 int32
	_ = v4540
	var v4571 int32
	_ = v4571
	var v4572 int32
	_ = v4572
	var v4574 int32
	_ = v4574
	var v4578 int32
	_ = v4578
	var v4580 int32
	_ = v4580
	var v4585 int32
	_ = v4585
	var v4639 int32
	_ = v4639
	var v4640 int32
	_ = v4640
	var v4642 int32
	_ = v4642
	var v4647 int32
	_ = v4647
	var v4648 int32
	_ = v4648
	var v4652 int32
	_ = v4652
	var v4654 int32
	_ = v4654
	var v4659 int32
	_ = v4659
	var v4663 int32
	_ = v4663
	var v4664 int32
	_ = v4664
	var v4667 int32
	_ = v4667
	var v4669 int32
	_ = v4669
	var v4671 int32
	_ = v4671
	var v4677 int32
	_ = v4677
	var v4678 int32
	_ = v4678
	var v4679 int64
	_ = v4679
	var v4690 int32
	_ = v4690
	var v4691 int32
	_ = v4691
	var v4696 int32
	_ = v4696
	var v4697 int32
	_ = v4697
	var v4698 int32
	_ = v4698
	var v4699 int32
	_ = v4699
	var v4702 int32
	_ = v4702
	var v4703 int32
	_ = v4703
	var v4713 int32
	_ = v4713
	var v4717 int32
	_ = v4717
	var v4722 int32
	_ = v4722
	var v4723 int32
	_ = v4723
	var v4724 int64
	_ = v4724
	var v4731 int64
	_ = v4731
	var v4738 int64
	_ = v4738
	var v4740 int64
	_ = v4740
	var v4741 int64
	_ = v4741
	var v4744 int64
	_ = v4744
	var v4746 int64
	_ = v4746
	var v4750 int64
	_ = v4750
	var v4752 int64
	_ = v4752
	var v4760 int64
	_ = v4760
	var v4765 int64
	_ = v4765
	var v4778 int64
	_ = v4778
	var v4780 int32
	_ = v4780
	var v4781 int32
	_ = v4781
	var v4786 int32
	_ = v4786
	var v4787 int32
	_ = v4787
	var v4792 int32
	_ = v4792
	var v4794 int32
	_ = v4794
	var v4800 int32
	_ = v4800
	var v4805 int32
	_ = v4805
	var v4811 int32
	_ = v4811
	var v4814 int32
	_ = v4814
	var v4815 int32
	_ = v4815
	var v4817 int32
	_ = v4817
	var v4819 int32
	_ = v4819
	var v4821 int32
	_ = v4821
	var v4822 int32
	_ = v4822
	var v4823 int32
	_ = v4823
	var v4835 int32
	_ = v4835
	var v4836 int32
	_ = v4836
	var v4882 int32
	_ = v4882
	var v4883 int32
	_ = v4883
	var v4885 int32
	_ = v4885
	var v4886 int32
	_ = v4886
	var v4889 int32
	_ = v4889
	var v4890 int32
	_ = v4890
	var v4893 int32
	_ = v4893
	var v4894 int32
	_ = v4894
	var v4897 int32
	_ = v4897
	var v4899 int32
	_ = v4899
	var v4907 int32
	_ = v4907
	var v4961 int32
	_ = v4961
	var v4962 int32
	_ = v4962
	var v5008 int32
	_ = v5008
	var v5011 int32
	_ = v5011
	var v5013 int32
	_ = v5013
	var v5021 int32
	_ = v5021
	var v5076 int32
	_ = v5076
	var v5123 int32
	_ = v5123
	var v5124 int32
	_ = v5124
	var v5127 int32
	_ = v5127
	var v5130 int32
	_ = v5130
	var v5131 int32
	_ = v5131
	var v5135 int32
	_ = v5135
	var v5137 int32
	_ = v5137
	var v5138 int32
	_ = v5138
	var v5141 int32
	_ = v5141
	var v5145 int32
	_ = v5145
	var v5146 int32
	_ = v5146
	var v5149 int32
	_ = v5149
	var v5152 int32
	_ = v5152
	var v5160 int32
	_ = v5160
	var v5168 int32
	_ = v5168
	var v5171 int32
	_ = v5171
	var v5172 int32
	_ = v5172
	var v5175 int32
	_ = v5175
	var v5185 int32
	_ = v5185
	var v5189 int32
	_ = v5189
	var v5194 int32
	_ = v5194
	var v5195 int32
	_ = v5195
	var v5198 int32
	_ = v5198
	var v5199 int32
	_ = v5199
	var v5203 int32
	_ = v5203
	var v5204 int32
	_ = v5204
	var v5208 int32
	_ = v5208
	var v5210 int32
	_ = v5210
	var v5217 int32
	_ = v5217
	var v5218 int32
	_ = v5218
	var v5219 int32
	_ = v5219
	var v5232 int32
	_ = v5232
	var v5233 int32
	_ = v5233
	var v5280 int32
	_ = v5280
	var v5287 int32
	_ = v5287
	var v5288 int32
	_ = v5288
	var v5292 int32
	_ = v5292
	var v5293 int32
	_ = v5293
	var v5352 int32
	_ = v5352
	var v5357 int32
	_ = v5357
	var v5358 int32
	_ = v5358
	var v5360 int32
	_ = v5360
	var v5363 int32
	_ = v5363
	var v5366 int32
	_ = v5366
	var v5367 int32
	_ = v5367
	var v5376 int32
	_ = v5376
	var v5377 int32
	_ = v5377
	var v5378 int32
	_ = v5378
	var v5380 int32
	_ = v5380
	var v5385 int32
	_ = v5385
	var v5386 int32
	_ = v5386
	var v5387 int32
	_ = v5387
	var v5388 int32
	_ = v5388
	var v5391 int32
	_ = v5391
	var v5392 int32
	_ = v5392
	var v5393 int32
	_ = v5393
	var v5396 int32
	_ = v5396
	var v5398 int32
	_ = v5398
	var v5403 int32
	_ = v5403
	var v5404 int32
	_ = v5404
	var v5406 int32
	_ = v5406
	var v5407 int32
	_ = v5407
	var v5410 int32
	_ = v5410
	var v5411 int32
	_ = v5411
	var v5412 int32
	_ = v5412
	var v5416 int32
	_ = v5416
	var v5419 int32
	_ = v5419
	var v5420 int32
	_ = v5420
	var v5421 int32
	_ = v5421
	var v5423 int32
	_ = v5423
	var v5428 int32
	_ = v5428
	var v5429 int32
	_ = v5429
	var v5431 int32
	_ = v5431
	var v5435 int32
	_ = v5435
	var v5438 int64
	_ = v5438
	var v5439 int32
	_ = v5439
	var v5443 int32
	_ = v5443
	var v5445 int32
	_ = v5445
	var v5451 int32
	_ = v5451
	var v5453 int32
	_ = v5453
	var v5456 int32
	_ = v5456
	var v5457 int32
	_ = v5457
	var v5460 int32
	_ = v5460
	var v5461 int32
	_ = v5461
	var v5465 int32
	_ = v5465
	var v5471 int32
	_ = v5471
	var v5473 int32
	_ = v5473
	var v5479 int32
	_ = v5479
	var v5480 int32
	_ = v5480
	var v5484 int32
	_ = v5484
	var v5490 int32
	_ = v5490
	var v5492 int32
	_ = v5492
	var v5498 int32
	_ = v5498
	var v5500 int32
	_ = v5500
	var v5502 int32
	_ = v5502
	var v5503 int32
	_ = v5503
	var v5508 int32
	_ = v5508
	var v5510 int32
	_ = v5510
	var v5513 int32
	_ = v5513
	var v5515 int32
	_ = v5515
	var v5517 int32
	_ = v5517
	var v5518 int32
	_ = v5518
	var v5521 int32
	_ = v5521
	var v5525 int32
	_ = v5525
	var v5526 int32
	_ = v5526
	var v5533 int32
	_ = v5533
	var v5539 int32
	_ = v5539
	var v5541 int32
	_ = v5541
	var v5552 int32
	_ = v5552
	var v5560 int32
	_ = v5560
	var v5566 int32
	_ = v5566
	var v5568 int32
	_ = v5568
	var v5575 int32
	_ = v5575
	var v5577 int32
	_ = v5577
	var v5585 int32
	_ = v5585
	var v5587 int32
	_ = v5587
	var v5588 int32
	_ = v5588
	var v5592 int32
	_ = v5592
	var v5593 int32
	_ = v5593
	var v5595 int32
	_ = v5595
	var v5599 int32
	_ = v5599
	var v5600 int32
	_ = v5600
	var v5601 int32
	_ = v5601
	var v5602 int32
	_ = v5602
	var v5604 int32
	_ = v5604
	var v5607 int32
	_ = v5607
	var v5608 int32
	_ = v5608
	var v5609 int32
	_ = v5609
	var v5613 int32
	_ = v5613
	var v5616 int32
	_ = v5616
	var v5617 int32
	_ = v5617
	var v5618 int32
	_ = v5618
	var v5620 int32
	_ = v5620
	var v5624 int32
	_ = v5624
	var v5628 int32
	_ = v5628
	var v5629 int32
	_ = v5629
	var v5631 int32
	_ = v5631
	var v5632 int32
	_ = v5632
	var v5636 int32
	_ = v5636
	var v5639 int32
	_ = v5639
	var v5644 int32
	_ = v5644
	var v5645 int32
	_ = v5645
	var v5647 int32
	_ = v5647
	var v5650 int64
	_ = v5650
	var v5651 int32
	_ = v5651
	var v5652 int64
	_ = v5652
	var v5653 int64
	_ = v5653
	var v5655 int32
	_ = v5655
	var v5659 int64
	_ = v5659
	var v5663 int32
	_ = v5663
	var v5665 int32
	_ = v5665
	var v5670 int32
	_ = v5670
	var v5673 int32
	_ = v5673
	var v5675 int32
	_ = v5675
	var v5680 int32
	_ = v5680
	var v5681 int32
	_ = v5681
	var v5682 int32
	_ = v5682
	var v5683 int32
	_ = v5683
	var v5686 int32
	_ = v5686
	var v5688 int32
	_ = v5688
	var v5691 int32
	_ = v5691
	var v5692 int32
	_ = v5692
	var v5696 int32
	_ = v5696
	var v5700 int32
	_ = v5700
	var v5701 int32
	_ = v5701
	var v5702 int32
	_ = v5702
	var v5703 int32
	_ = v5703
	var v5712 int32
	_ = v5712
	var v5763 int32
	_ = v5763
	var v5767 int32
	_ = v5767
	var v5768 int32
	_ = v5768
	var v5771 int32
	_ = v5771
	var v5772 int32
	_ = v5772
	var v5775 int32
	_ = v5775
	var v5776 int32
	_ = v5776
	var v5785 int32
	_ = v5785
	var v5831 int32
	_ = v5831
	var v5832 int32
	_ = v5832
	var v5833 int32
	_ = v5833
	var v5834 int32
	_ = v5834
	var v5835 int32
	_ = v5835
	var v5836 int32
	_ = v5836
	var v5838 int32
	_ = v5838
	var v5843 int32
	_ = v5843
	var v5844 int32
	_ = v5844
	var v5845 int32
	_ = v5845
	var v5863 int32
	_ = v5863
	var v5864 int32
	_ = v5864
	var v5914 int32
	_ = v5914
	var v5915 int32
	_ = v5915
	var v5919 int32
	_ = v5919
	var v5920 int32
	_ = v5920
	var v5924 int32
	_ = v5924
	var v5936 int32
	_ = v5936
	var v5979 int32
	_ = v5979
	var v5980 int32
	_ = v5980
	var v5981 int32
	_ = v5981
	var v5982 int32
	_ = v5982
	var v5983 int32
	_ = v5983
	var v5986 int32
	_ = v5986
	var v5991 int32
	_ = v5991
	var v5995 int32
	_ = v5995
	var v5996 int32
	_ = v5996
	var v5997 int32
	_ = v5997
	var v5998 int32
	_ = v5998
	var v5999 int32
	_ = v5999
	var v6003 int32
	_ = v6003
	var v6006 int32
	_ = v6006
	var v6008 int32
	_ = v6008
	var v6009 int32
	_ = v6009
	var v6011 int32
	_ = v6011
	var v6013 int32
	_ = v6013
	var v6014 int32
	_ = v6014
	var v6019 int32
	_ = v6019
	var v6023 int32
	_ = v6023
	var v6024 int32
	_ = v6024
	var v6026 int32
	_ = v6026
	var v6027 int32
	_ = v6027
	var v6029 int32
	_ = v6029
	var v6034 int32
	_ = v6034
	var v6035 int32
	_ = v6035
	var v6037 int32
	_ = v6037
	var v6038 int32
	_ = v6038
	var v6043 int32
	_ = v6043
	var v6045 int32
	_ = v6045
	var v6046 int32
	_ = v6046
	var v6052 int32
	_ = v6052
	var v6058 int32
	_ = v6058
	var v6060 int32
	_ = v6060
	var v6066 int32
	_ = v6066
	var v6067 int32
	_ = v6067
	var v6068 int32
	_ = v6068
	var v6072 int32
	_ = v6072
	var v6078 int32
	_ = v6078
	var v6080 int32
	_ = v6080
	var v6086 int32
	_ = v6086
	var v6087 int32
	_ = v6087
	var v6089 int32
	_ = v6089
	var v6090 int32
	_ = v6090
	var v6093 int32
	_ = v6093
	var v6094 int32
	_ = v6094
	var v6095 int32
	_ = v6095
	var v6096 int32
	_ = v6096
	var v6098 int32
	_ = v6098
	var v6099 int32
	_ = v6099
	var v6103 int32
	_ = v6103
	var v6116 int32
	_ = v6116
	var v6117 int32
	_ = v6117
	var v6164 int32
	_ = v6164
	var v6169 int32
	_ = v6169
	var v6227 int32
	_ = v6227
	var v6233 int32
	_ = v6233
	var v6238 int32
	_ = v6238
	var v6242 int32
	_ = v6242
	var v6244 int32
	_ = v6244
	var v6245 int32
	_ = v6245
	var v6248 int32
	_ = v6248
	var v6249 int32
	_ = v6249
	var v6253 int32
	_ = v6253
	var v6266 int32
	_ = v6266
	var v6267 int32
	_ = v6267
	var v6314 int32
	_ = v6314
	var v6319 int32
	_ = v6319
	var v6377 int32
	_ = v6377
	var v6383 int32
	_ = v6383
	var v6388 int32
	_ = v6388
	var v6393 int32
	_ = v6393
	var v6394 int32
	_ = v6394
	var v6395 int32
	_ = v6395
	var v6399 int32
	_ = v6399
	var v6402 int32
	_ = v6402
	var v6403 int32
	_ = v6403
	var v6404 int32
	_ = v6404
	var v6406 int32
	_ = v6406
	var v6411 int32
	_ = v6411
	var v6412 int32
	_ = v6412
	var v6414 int32
	_ = v6414
	var v6415 int32
	_ = v6415
	var v6417 int32
	_ = v6417
	var v6421 int32
	_ = v6421
	var v6425 int32
	_ = v6425
	var v6426 int32
	_ = v6426
	var v6428 int32
	_ = v6428
	var v6431 int64
	_ = v6431
	var v6432 int32
	_ = v6432
	var v6438 int32
	_ = v6438
	var v6440 int32
	_ = v6440
	var v6447 int32
	_ = v6447
	var v6453 int32
	_ = v6453
	var v6455 int32
	_ = v6455
	var v6461 int32
	_ = v6461
	var v6464 int64
	_ = v6464
	var v6466 int32
	_ = v6466
	var v6468 int32
	_ = v6468
	var v6472 int32
	_ = v6472
	var v6474 int32
	_ = v6474
	var v6483 int32
	_ = v6483
	var v6485 int32
	_ = v6485
	var v6488 int32
	_ = v6488
	var v6490 int32
	_ = v6490
	var v6496 int32
	_ = v6496
	var v6497 int32
	_ = v6497
	var v6503 int32
	_ = v6503
	var v6508 int32
	_ = v6508
	var v6512 int32
	_ = v6512
	var v6516 int32
	_ = v6516
	var v6521 int32
	_ = v6521
	var v6525 int32
	_ = v6525
	var v6529 int32
	_ = v6529
	var v6534 int32
	_ = v6534
	var v6538 int32
	_ = v6538
	var v6539 int32
	_ = v6539
	var v6545 int32
	_ = v6545
	var v6550 int32
	_ = v6550
	var v6554 int32
	_ = v6554
	var v6558 int32
	_ = v6558
	var v6563 int32
	_ = v6563
	var v6567 int32
	_ = v6567
	var v6571 int32
	_ = v6571
	var v6576 int32
	_ = v6576
	var v6580 int32
	_ = v6580
	var v6581 int32
	_ = v6581
	var v6589 int32
	_ = v6589
	var v6594 int32
	_ = v6594
	var v6598 int32
	_ = v6598
	var v6599 int32
	_ = v6599
	var v6605 int32
	_ = v6605
	var v6610 int32
	_ = v6610
	var v6614 int32
	_ = v6614
	var v6615 int32
	_ = v6615
	var v6621 int32
	_ = v6621
	var v6626 int32
	_ = v6626
	var v6630 int32
	_ = v6630
	var v6634 int32
	_ = v6634
	var v6639 int32
	_ = v6639
	var v6643 int32
	_ = v6643
	var v6644 int32
	_ = v6644
	var v6650 int32
	_ = v6650
	var v6655 int32
	_ = v6655
	var v6656 int32
	_ = v6656
	var v6657 int32
	_ = v6657
	var v6659 int32
	_ = v6659
	var v6662 int32
	_ = v6662
	var v6718 int32
	_ = v6718
	var v6723 int32
	_ = v6723
	var v6726 int32
	_ = v6726
	var v6727 int32
	_ = v6727
	var v6736 int32
	_ = v6736
	var v6740 int32
	_ = v6740
	var v6745 int32
	_ = v6745
	var v6802 int32
	_ = v6802
	var v6808 int32
	_ = v6808
	var v6813 int32
	_ = v6813
	var v6874 int32
	_ = v6874
	var v6882 int32
	_ = v6882
	var v6925 int32
	_ = v6925
	var v6927 int32
	_ = v6927
	var v6934 int32
	_ = v6934
	var v6942 int32
	_ = v6942
	var v6981 int32
	_ = v6981
	var v6986 int32
	_ = v6986
	var v6988 int32
	_ = v6988
	var v6990 int32
	_ = v6990
	var v6994 int32
	_ = v6994
	var v7001 int32
	_ = v7001
	v6 = int32(0)
	v54 = m.G0
	v56 = v54 - int32(800)
	m.G0 = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v60 != 0 {
		v100 = v6
		v101 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+432)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if int32(2) <= v103 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v63 = F_index_getprocinfo(m, l0, int32(1), int32(2))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+6)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v69+v71*int32(0)<<(uint(int32(2))%32)+int32(24)-int32(4))))
	goto L5
L5:
	;
	if v83 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v86 = F_index_getprocinfo(m, l0, int32(1), int32(6))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+24)))
	if v94 != int32(65535) {
		v100 = v63
		v101 = v93
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v91 = F_FunctionCall1Coll(m, v86, v89, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v100 = v63
	v101 = v91
	goto L1
L11:
	;
	v97 = F_pg_detoast_datum(m, v93)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v100 = v63
	v101 = v97
	goto L1
L13:
	;
	v115 = int32(1)
	v116 = v103
	goto L16
L14:
	;
	goto L15
L15:
	;
	v255 = F_SpGistGetLeafTupleSize(m, v58, v56+int32(432), l4)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L3
	} else {
		goto L33
	}
L16:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v115))))
	if v163 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L15
L18:
	;
	v198 = v115 + int32(1)
	if v198 < v194 {
		v115 = v198
		v116 = v194
		goto L16
	} else {
		goto L26
	}
L19:
	;
	v167 = v115 << (uint(int32(2)) % 32)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l3+v167)))
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58+int32(24)+v115<<(uint(int32(4))%32)))))
	if v173 == int32(65535) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56+int32(432)+v115<<(uint(int32(2))%32)))) = int32(0)
	v194 = v116
	goto L18
L22:
	;
	v179 = F_pg_detoast_datum(m, v169)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L3
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56+int32(432)+v167))) = v169
	v194 = v116
	goto L18
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56+int32(432)+v167))) = v179
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v194 = v182
	goto L18
L26:
	;
	goto L17
L27:
	;
	m.G0 = v56 + int32(800)
	return v7001
L28:
	;
	v6981 = *(*int32)(unsafe.Add(mBase, uint32(v56)+416))
	if v6981 == int32(0) {
		goto L1011
	} else {
		goto L1012
	}
L29:
	;
	if v6882 == int32(0) {
		goto L1006
	} else {
		goto L1007
	}
L30:
	;
	v6874 = int32(1)
	v6882 = v408
	goto L29
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6802 = m.ExcPending
	if v6802 != 0 {
		goto L3
	} else {
		goto L1003
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6723 = m.ExcPending
	if v6723 != 0 {
		goto L3
	} else {
		goto L998
	}
L33:
	;
	v258 = v255 + int32(4)
	if base.Ui32(int32(8160)) < base.Ui32(v258) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v60 != 0 {
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+428)) = int32(-1)
	v266 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+424)) = uint16(v266)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+420)) = v266
	*(*int64)(unsafe.Add(mBase, uint32(v56)+412)) = int64(4294967295)
	v273 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v273 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v261 == int32(0) {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L3
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v276 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+411)) = uint8(v276)
	v279 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v279 == v276 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	if v60 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v6718 = int32(0)
	v6934 = v6718
	v6942 = v6718
	goto L28
L46:
	;
	v284 = int32(2)
	goto L48
L47:
	;
	v284 = int32(1)
	goto L48
L48:
	;
	if v60 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v287 = int32(8)
	goto L51
L50:
	;
	v287 = int32(0)
	goto L51
L51:
	;
	if v60 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v290 = int32(12)
	goto L54
L53:
	;
	v290 = int32(4)
	goto L54
L54:
	;
	if v60 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v293 = int32(4)
	goto L57
L56:
	;
	v293 = int32(0)
	goto L57
L57:
	;
	if v60 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v296 = int32(7)
	goto L60
L59:
	;
	v296 = int32(3)
	goto L60
L60:
	;
	v298 = v56 + int32(628)
	v307 = int32(-1)
	v317 = int32(0)
	v318 = int32(1)
	v319 = v258
	v320 = v284
	v321 = v307
	v323 = v6
	v325 = v6
	v326 = v307
	v333 = v6
	v342 = v6
	v344 = v258
	goto L61
L61:
	;
	if v320 == int32(-1) {
		goto L69
	} else {
		goto L70
	}
L62:
	;
	goto L45
L63:
	;
	v4517 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v4517 != 0 {
		v6874 = int32(0)
		v6882 = v4475
		goto L29
	} else {
		goto L649
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+788)) = v1802
	if v60 != 0 {
		v1869 = int32(0)
		goto L304
	} else {
		goto L305
	}
L65:
	;
	v1782 = int32(0)
	v1802 = v1782
	v1804 = v1782
	v1805 = v1782
	goto L64
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L3
	} else {
		goto L301
	}
L67:
	;
	F_ReleaseBuffer(m, v401)
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L3
	} else {
		goto L299
	}
L68:
	;
	if v408 < int32(0) {
		goto L91
	} else {
		goto L92
	}
L69:
	;
	v366 = int32(8160)
	if base.Ui32(v366) <= base.Ui32(v344) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	if v325 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L72:
	;
	v369 = v366
	goto L74
L73:
	;
	v369 = v344
	goto L74
L74:
	;
	v372 = F_SpGistGetBuffer(m, l0, v296, v369, v56+int32(411))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	if v372 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v408 = v372
	v409 = v392
	goto L68
L77:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v377+(v372^int32(-1))<<(uint(int32(6))%32))+16))
	v392 = v383
	goto L76
L78:
	;
	goto L79
L79:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v385+v372<<(uint(int32(6))%32)+int32(-64))+16))
	v392 = v391
	goto L76
L80:
	;
	v408 = v407
	v409 = v320
	goto L68
L81:
	;
	v395 = F_ReadBuffer(m, l0, v320)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L3
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v320 == v326 {
		v408 = v325
		v409 = v326
		goto L68
	} else {
		goto L86
	}
L84:
	;
	F_LockBuffer(m, v395, int32(2))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	v407 = v395
	goto L80
L86:
	;
	v401 = F_ReadBuffer(m, l0, v320)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L3
	} else {
		goto L87
	}
L87:
	;
	v403 = F_ConditionalLockBuffer(m, v401)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	if v403 == int32(0) {
		goto L67
	} else {
		goto L89
	}
L89:
	;
	v407 = v401
	goto L80
L90:
	;
	v428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427)+16)))
	v430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v428+v427))))
	v433 = int32(0)
	if base.B2i32(v430&int32(8) == v433)^v60 == v433 {
		goto L66
	} else {
		goto L94
	}
L91:
	;
	v413 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v413+(v408^int32(-1))<<(uint(int32(2))%32))))
	v427 = v419
	goto L90
L92:
	;
	goto L93
L93:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v427 = v421 + v408<<(uint(int32(13))%32) + int32(-8192)
	goto L90
L94:
	;
	if v430&int32(4) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v4473 = v427
	v4475 = v408
	v4483 = v318
	v4514 = v409
	goto L63
L96:
	;
	goto L97
L97:
	;
	v444 = F_spgFormLeafTuple(m, l1, l2, v56+int32(432), l4)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	v451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427)+14)))
	v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427)+12)))
	v453 = v451 - v452
	v454 = int32(0)
	if v454 < v453 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v460 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427)+16)))
	v462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427+v460)+4)))
	if v462 != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v457 = v453
	goto L102
L101:
	;
	v457 = v454
	goto L102
L102:
	;
	goto L99
L103:
	;
	v463 = int32(20)
	goto L105
L104:
	;
	v463 = int32(0)
	goto L105
L105:
	;
	if base.Ui32(int32(base.Ui32(v446)>>(uint(int32(2))%32))+int32(4)) <= base.Ui32(v457+v463) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v466 = int32(4556740)
	v468 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v469 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v468 + v469
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+411)))
	*(*int64)(unsafe.Add(mBase, uint32(v56)+626)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+625)) = uint8(v60)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+624)) = uint8(v472)
	if base.Ui32(v469) < base.Ui32(v409-v469) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	goto L108
L108:
	;
	v660 = v409 - int32(1)
	v662 = base.B2i32(base.Ui32(v660) < base.Ui32(int32(2)))
	if base.Ui32(v660) < base.Ui32(int32(2)) {
		goto L155
	} else {
		goto L156
	}
L109:
	;
	F_MarkBufferDirty(m, v408)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L3
	} else {
		goto L133
	}
L110:
	;
	v484 = v318 & int32(65535)
	goto L112
L111:
	;
	v484 = int32(0)
	goto L112
L112:
	;
	if v484 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444)+4)))
	v489 = v487 & int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v444)+4)) = uint16(v489)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	v495 = F_SpGistPageAddNewItem(m, v427, v444, int32(base.Ui32(v491)>>(uint(int32(2))%32)), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L3
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v507 = v318 & int32(65535)
	v512 = v507<<(uint(int32(2))%32) + v427 + int32(20)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v516 = v427 + v513&int32(32767)
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	switch v517 & int32(3) {
	case 0:
		goto L120
	default:
		goto L121
	case 2:
		goto L122
	}
L116:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+626)) = uint16(v495)
	if v325 == int32(0) {
		goto L109
	} else {
		goto L117
	}
L117:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+632)) = uint16(v321)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+630)) = uint16(v333)
	F_saveNodeLink(m, v56+int32(412), v409, v495)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L3
	} else {
		goto L118
	}
L118:
	;
	goto L109
L119:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+628)) = uint16(v318)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+626)) = uint16(v595)
	goto L109
L120:
	;
	v569 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v516)+4)))
	v572 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444)+4)))
	v575 = v569&int32(16383) | v572&int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v444)+4)) = uint16(v575)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	v581 = F_SpGistPageAddNewItem(m, v427, v444, int32(base.Ui32(v577)>>(uint(int32(2))%32)), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L3
	} else {
		goto L132
	}
L121:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L3
	} else {
		goto L129
	}
L122:
	;
	v520 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444)+4)))
	v522 = v520 & int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v444)+4)) = uint16(v522)
	F_PageIndexTupleDelete(m, v427, v507)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L3
	} else {
		goto L123
	}
L123:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	v530 = F_PageAddItemExtended(m, v427, v444, int32(base.Ui32(v526)>>(uint(int32(2))%32)), v507, int32(0))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L3
	} else {
		goto L124
	}
L124:
	;
	if v530 == v507 {
		v595 = v318
		goto L119
	} else {
		goto L125
	}
L125:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+320)) = int32(base.Ui32(v537) >> (uint(int32(2)) % 32))
	F_errmsg_internal(m, int32(426927), v56+int32(320))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(516634), int32(280), int32(404011))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+304)) = v555 & int32(3)
	F_errmsg_internal(m, int32(506912), v56+int32(304))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L3
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(516634), int32(287), int32(404011))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v586 = v427 + v583&int32(32767)
	v587 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v586)+4)))
	v592 = v587&int32(49152) | v581&int32(16383)
	*(*uint16)(unsafe.Add(mBase, uint32(v586)+4)) = uint16(v592)
	v595 = v581
	goto L119
L133:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603)+118)))
	if v604 != int32(112) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v653 = int32(4556740)
	v655 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v655 - int32(1)
	goto L30
L135:
	;
	v608 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v608 <= int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v611 != 0 {
		goto L134
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v613 != 0 {
		goto L134
	} else {
		goto L141
	}
L139:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v612 != 0 {
		goto L134
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L3
	} else {
		goto L142
	}
L142:
	;
	F_XLogRegisterData(m, v56+int32(624), int32(10))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L3
	} else {
		goto L143
	}
L143:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	F_XLogRegisterData(m, v444, int32(base.Ui32(v621)>>(uint(int32(2))%32)))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L3
	} else {
		goto L144
	}
L144:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+624)))
	if v629 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v630 = int32(14)
	goto L147
L146:
	;
	v630 = int32(8)
	goto L147
L147:
	;
	F_XLogRegisterBuffer(m, int32(0), v408, v630)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L3
	} else {
		goto L148
	}
L148:
	;
	v633 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+630)))
	if v633 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	F_XLogRegisterBuffer(m, int32(1), v325, int32(8))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L3
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v638 = int32(16)
	v640 = F_XLogInsert(m, v638, v638)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L3
	} else {
		goto L153
	}
L152:
	;
	goto L151
L153:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v427))) = base.I64_rotr(v640, int64(32))
	v645 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+630)))
	if v645 == int32(0) {
		goto L134
	} else {
		goto L154
	}
L154:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v323)+4)) = uint32(v640)
	v650 = int64(base.Ui64(v640) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v323))) = uint32(v650)
	goto L134
L155:
	;
	v1422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+411)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+796)) = v342
	v1424 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427)+12)))
	v1430 = int32(base.Ui32(v1424+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if base.Ui32(int32(25)) <= base.Ui32(v1424) {
		goto L247
	} else {
		goto L248
	}
L156:
	;
	if v318&int32(65535) == int32(0) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	if base.Ui32(int32(4079)) < base.Ui32(v780) {
		goto L155
	} else {
		goto L170
	}
L158:
	;
	v780 = int32(0)
	v825 = int32(1)
	goto L157
L159:
	;
	goto L160
L160:
	;
	v671 = int32(0)
	v679 = v318
	v681 = v671
	v683 = v671
	goto L161
L161:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v679&int32(65535)<<(uint(int32(2))%32)+(v427+int32(24))-int32(4))))
	v736 = v427 + v733&int32(32767)
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v736)))
	switch v737 & int32(3) {
	case 0:
		goto L164
	default:
		goto L165
	case 2:
		v765 = v681
		v766 = v683
		goto L163
	}
L162:
	;
	v780 = v765
	v825 = base.B2i32(v766 < int32(64))
	goto L157
L163:
	;
	v767 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v736)+4)))
	v769 = v767 & int32(16383)
	if v769 != 0 {
		v679 = v769
		v681 = v765
		v683 = v766
		goto L161
	} else {
		goto L169
	}
L164:
	;
	v765 = v681 + int32(base.Ui32(v737)>>(uint(int32(2))%32)) + int32(4)
	v766 = v683 + int32(1)
	goto L163
L165:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L3
	} else {
		goto L166
	}
L166:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v736)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+288)) = v744 & int32(3)
	F_errmsg_internal(m, int32(506912), v56+int32(288))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L3
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(516634), int32(369), int32(149202))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L3
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	goto L162
L170:
	;
	if v825 == int32(0) {
		goto L155
	} else {
		goto L171
	}
L171:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	if base.Ui32(int32(8160)) < base.Ui32(v780+int32(base.Ui32(v830)>>(uint(int32(2))%32))+int32(4)) {
		goto L155
	} else {
		goto L172
	}
L172:
	;
	v838 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+592)) = uint16(v838)
	v841 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v841) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v851 = int32(base.Ui32(v841+int32(262120))>>(uint(int32(1))%32)) & int32(131070)
	goto L175
L174:
	;
	v851 = v838
	goto L175
L175:
	;
	v852 = F_palloc(m, v851)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L3
	} else {
		goto L176
	}
L176:
	;
	v856 = F_palloc(m, v851+int32(2))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L3
	} else {
		goto L177
	}
L177:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	v862 = int32(base.Ui32(v858)>>(uint(int32(2))%32)) + int32(4)
	if v318&int32(65535) == int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v1035 = F_SpGistGetBuffer(m, l0, v296, v988, v56+int32(624)|int32(2))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L3
	} else {
		goto L193
	}
L179:
	;
	v987 = int32(0)
	v988 = v862
	v998 = v838
	goto L178
L180:
	;
	goto L181
L181:
	;
	v878 = v318
	v880 = int32(0)
	v881 = v862
	v891 = v838
	goto L182
L182:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v878&int32(65535)<<(uint(int32(2))%32)+(v427+int32(24))-int32(4))))
	v934 = v427 + v931&int32(32767)
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v934)))
	switch v935 & int32(3) {
	case 0:
		goto L185
	default:
		goto L186
	case 2:
		goto L187
	}
L183:
	;
	v987 = v974
	v988 = v971
	v998 = v972
	goto L178
L184:
	;
	v974 = v880 + int32(1)
	v975 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+4)))
	v977 = v975 & int32(16383)
	if v977 != 0 {
		v878 = v977
		v880 = v974
		v881 = v971
		v891 = v972
		goto L182
	} else {
		goto L191
	}
L185:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v852+v880<<(uint(int32(1))%32)))) = uint16(v878)
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v934)))
	v971 = v881 + int32(base.Ui32(v965)>>(uint(int32(2))%32)) + int32(4)
	v972 = v891
	goto L184
L186:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L3
	} else {
		goto L188
	}
L187:
	;
	v938 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v852+v880<<(uint(v938)%32)))) = uint16(v878)
	v971 = v881
	v972 = v938
	goto L184
L188:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v934)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+272)) = v947 & int32(3)
	F_errmsg_internal(m, int32(506912), v56+int32(272))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L3
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(516634), int32(446), int32(166949))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
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
	if v1035 < int32(0) {
		goto L198
	} else {
		goto L199
	}
L193:
	;
	if v1035 < int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1040+(v1035^int32(-1))<<(uint(int32(2))%32))))
	v1054 = v1046
	goto L192
L195:
	;
	goto L196
L196:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v1054 = v1048 + v1035<<(uint(int32(13))%32) + int32(-8192)
	goto L192
L197:
	;
	v1074 = F_palloc(m, v988)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L3
	} else {
		goto L201
	}
L198:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1058+(v1035^int32(-1))<<(uint(int32(6))%32))+16))
	v1073 = v1064
	goto L197
L199:
	;
	goto L200
L200:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1066+v1035<<(uint(int32(6))%32)+int32(-64))+16))
	v1073 = v1072
	goto L197
L201:
	;
	v1076 = int32(4556740)
	v1078 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1079 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1078 + v1079
	v1082 = int32(0)
	if (v998|base.B2i32(v987 <= v1082))&v1079 == v1082 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v1101 = v1082
	v1104 = v1074
	v1106 = int32(0)
	goto L205
L203:
	;
	v1196 = v1082
	v1197 = v1082
	v1200 = v1074
	goto L204
L204:
	;
	v1243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444)+4)))
	v1246 = v1243&int32(49152) | v1196
	*(*uint16)(unsafe.Add(mBase, uint32(v444)+4)) = uint16(v1246)
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	v1256 = F_SpGistPageAddNewItem(m, v1054, v444, int32(base.Ui32(v1251)>>(uint(int32(2))%32)), v56+int32(592))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L3
	} else {
		goto L213
	}
L205:
	;
	v1148 = v1101 << (uint(int32(1)) % 32)
	v1150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v852+v1148))))
	v1151 = int32(2)
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1150<<(uint(v1151)%32)+(v427+int32(24))-int32(4))))
	v1159 = v427 + v1156&int32(32767)
	v1160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1159)+4)))
	v1165 = v1160&int32(49152) | v1106&int32(16383)
	*(*uint16)(unsafe.Add(mBase, uint32(v1159)+4)) = uint16(v1165)
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1159)))
	v1173 = F_SpGistPageAddNewItem(m, v1054, v1159, int32(base.Ui32(v1168)>>(uint(v1151)%32)), v56+int32(592))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L3
	} else {
		goto L207
	}
L206:
	;
	v1196 = v1173 & int32(16383)
	v1197 = v987
	v1200 = v1184
	goto L204
L207:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v856+v1148))) = uint16(v1173)
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1159)))
	v1178 = int32(base.Ui32(v1176) >> (uint(int32(2)) % 32))
	if v1178 != 0 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1159)))
	v1184 = v1180 + int32(base.Ui32(v1181)>>(uint(int32(2))%32))
	v1186 = v1101 + int32(1)
	if v1186 != v987 {
		v1101 = v1186
		v1104 = v1184
		v1106 = v1173
		goto L205
	} else {
		goto L212
	}
L209:
	;
	v1179 = F__emscripten_memcpy_bulkmem(m, v1104, v1159, v1178)
	mBase = m.M
	v1180 = v1179
	goto L211
L210:
	;
	v1180 = v1104
	goto L211
L211:
	;
	goto L208
L212:
	;
	goto L206
L213:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v856+v1197<<(uint(int32(1))%32)))) = uint16(v1256)
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	v1261 = int32(base.Ui32(v1259) >> (uint(int32(2)) % 32))
	if v1261 != 0 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v1267 != 0 {
		goto L218
	} else {
		goto L219
	}
L215:
	;
	v1262 = F__emscripten_memcpy_bulkmem(m, v1200, v444, v1261)
	mBase = m.M
	v1263 = v1262
	goto L217
L216:
	;
	v1263 = v1200
	goto L217
L217:
	;
	goto L214
L218:
	;
	v1268 = int32(3)
	goto L220
L219:
	;
	v1268 = int32(1)
	goto L220
L220:
	;
	F_spgPageIndexMultiDelete(m, l1, v427, v852, v987, v1268, int32(3), v1073, v1256)
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L3
	} else {
		goto L221
	}
L221:
	;
	F_saveNodeLink(m, v56+int32(412), v1073, v1256)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L3
	} else {
		goto L222
	}
L222:
	;
	F_MarkBufferDirty(m, v408)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L3
	} else {
		goto L223
	}
L223:
	;
	F_MarkBufferDirty(m, v1035)
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L3
	} else {
		goto L224
	}
L224:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280)+118)))
	if v1281 != int32(112) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1359 = int32(4556740)
	v1361 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1361 - int32(1)
	F_SpGistSetLastUsedPage(m, l0, v1035)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L3
	} else {
		goto L245
	}
L226:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v1285 <= int32(0) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1288 != 0 {
		goto L225
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v1290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v1290 != 0 {
		goto L225
	} else {
		goto L232
	}
L230:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1289 != 0 {
		goto L225
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+640)) = uint8(v1290)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+636)) = v1291
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+628)) = uint8(v60)
	v1296 = v998 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+627)) = uint8(v1296)
	v1298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+424)))
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+630)) = uint16(v1298)
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v56)+428))
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+632)) = uint16(v1300)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+624)) = uint16(v987)
	F_XLogBeginInsert(m)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L3
	} else {
		goto L233
	}
L233:
	;
	F_XLogRegisterData(m, v56+int32(624), int32(20))
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L3
	} else {
		goto L234
	}
L234:
	;
	F_XLogRegisterData(m, v852, v987<<(uint(int32(1))%32))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L3
	} else {
		goto L235
	}
L235:
	;
	F_XLogRegisterData(m, v856, v1197<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L3
	} else {
		goto L236
	}
L236:
	;
	F_XLogRegisterData(m, v1074, v1263+int32(base.Ui32(v1264)>>(uint(int32(2))%32))-v1074)
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L3
	} else {
		goto L237
	}
L237:
	;
	F_XLogRegisterBuffer(m, int32(0), v408, int32(8))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L3
	} else {
		goto L238
	}
L238:
	;
	v1333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+626)))
	if v1333 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1334 = int32(14)
	goto L241
L240:
	;
	v1334 = int32(8)
	goto L241
L241:
	;
	F_XLogRegisterBuffer(m, int32(1), v1035, v1334)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L3
	} else {
		goto L242
	}
L242:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v56)+416))
	F_XLogRegisterBuffer(m, int32(2), v1338, int32(8))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L3
	} else {
		goto L243
	}
L243:
	;
	v1344 = F_XLogInsert(m, int32(16), int32(32))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L3
	} else {
		goto L244
	}
L244:
	;
	v1346 = int64(32)
	v1347 = base.I64_rotr(v1344, v1346)
	*(*int64)(unsafe.Add(mBase, uint32(v427))) = v1347
	*(*uint32)(unsafe.Add(mBase, uint32(v1054)+4)) = uint32(v1344)
	v1351 = int64(base.Ui64(v1344) >> (uint(v1346) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1054))) = uint32(v1351)
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v56)+420))
	*(*int64)(unsafe.Add(mBase, uint32(v1353))) = v1347
	goto L225
L245:
	;
	F_UnlockReleaseBuffer(m, v1035)
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L3
	} else {
		goto L246
	}
L246:
	;
	goto L30
L247:
	;
	v1434 = v1430
	goto L249
L248:
	;
	v1434 = int32(0)
	goto L249
L249:
	;
	v1436 = v1434 + int32(1)
	v1438 = v1436 << (uint(int32(2)) % 32)
	v1439 = F_palloc(m, v1438)
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L3
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+792)) = v1439
	v1443 = v1436 << (uint(int32(1)) % 32)
	v1444 = F_palloc(m, v1443)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L3
	} else {
		goto L251
	}
L251:
	;
	v1446 = F_palloc(m, v1443)
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L3
	} else {
		goto L252
	}
L252:
	;
	v1448 = F_palloc(m, v1438)
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L3
	} else {
		goto L253
	}
L253:
	;
	v1450 = F_palloc(m, v1438)
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L3
	} else {
		goto L254
	}
L254:
	;
	v1452 = F_palloc(m, v1436)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L3
	} else {
		goto L255
	}
L255:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+584)) = v1454
	v1456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+588)) = uint8(v1456)
	if base.Ui32(v660) <= base.Ui32(int32(1)) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	if v1434 == int32(0) {
		goto L65
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	if v318&int32(65535) == int32(0) {
		goto L65
	} else {
		goto L278
	}
L259:
	;
	v1464 = int32(0)
	v1474 = int32(1)
	v1476 = v1464
	v1487 = v1464
	goto L262
L260:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L3
	} else {
		goto L275
	}
L261:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L3
	} else {
		goto L272
	}
L262:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1474&int32(65535)<<(uint(int32(2))%32)+(v427+int32(24))-int32(4))))
	v1530 = v427 + v1527&int32(32767)
	v1531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1530))))
	if v1531&int32(3) != 0 {
		goto L260
	} else {
		goto L264
	}
L263:
	;
	v1802 = v1430
	v1804 = v1430
	v1805 = v1567
	goto L64
L264:
	;
	if v60 != 0 {
		v1548 = int32(0)
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1549 = int32(2)
	v1550 = v1476 << (uint(v1549) % 32)
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v56)+792))
	*(*int32)(unsafe.Add(mBase, uint32(v1550+v1551))) = v1548
	*(*int32)(unsafe.Add(mBase, uint32(v1448+v1550))) = v1530
	v1556 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1444+v1476<<(uint(v1556)%32)))) = uint16(v1474)
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1530)))
	v1567 = v1487 + int32(base.Ui32(v1562)>>(uint(v1549)%32)) + int32(4)
	v1569 = v1476 + v1556
	if v1569 != v1434 {
		v1474 = v1474 + v1556
		v1476 = v1569
		v1487 = v1567
		goto L262
	} else {
		goto L271
	}
L266:
	;
	v1536 = v1530 + int32(16)
	v1537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v1537 != int32(1) {
		v1548 = v1536
		goto L265
	} else {
		goto L267
	}
L267:
	;
	v1540 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)))
	switch v1540 - int32(1) {
	case 0:
		goto L270
	case 1:
		goto L269
	default:
		goto L261
	case 3:
		goto L268
	}
L268:
	;
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1536)))
	v1548 = v1545
	goto L265
L269:
	;
	v1544 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1536))))
	v1548 = v1544
	goto L265
L270:
	;
	v1543 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1536))))
	v1548 = v1543
	goto L265
L271:
	;
	goto L263
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+208)) = base.I32_extend16_s(v1540)
	F_errmsg_internal(m, int32(506326), v56+int32(208))
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L3
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(343236), int32(70), int32(73749))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L3
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v1530)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+224)) = v1591 & int32(3)
	F_errmsg_internal(m, int32(506912), v56+int32(224))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L3
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(516634), int32(767), int32(109607))
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L3
	} else {
		goto L277
	}
L277:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L278:
	;
	v1611 = int32(0)
	v1621 = v318
	v1631 = v1611
	v1633 = v1611
	v1634 = v1611
	goto L279
L279:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1621&int32(65535)<<(uint(int32(2))%32)+(v427+int32(24))-int32(4))))
	v1677 = v427 + v1674&int32(32767)
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1677)))
	switch v1678 & int32(3) {
	case 0:
		goto L284
	default:
		goto L283
	case 2:
		goto L282
	}
L280:
	;
	v1802 = v1755
	v1804 = v1758
	v1805 = v1756
	goto L64
L281:
	;
	v1758 = v1633 + int32(1)
	v1759 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1677)+4)))
	v1761 = v1759 & int32(16383)
	if v1761 != 0 {
		v1621 = v1761
		v1631 = v1755
		v1633 = v1758
		v1634 = v1756
		goto L279
	} else {
		goto L298
	}
L282:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1444+v1633<<(uint(int32(1))%32)))) = uint16(v1621)
	v1755 = v1631
	v1756 = v1634
	goto L281
L283:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L3
	} else {
		goto L295
	}
L284:
	;
	if v60 != 0 {
		v1695 = int32(0)
		goto L286
	} else {
		goto L287
	}
L285:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L3
	} else {
		goto L292
	}
L286:
	;
	v1696 = int32(2)
	v1697 = v1631 << (uint(v1696) % 32)
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v56)+792))
	*(*int32)(unsafe.Add(mBase, uint32(v1697+v1698))) = v1695
	*(*int32)(unsafe.Add(mBase, uint32(v1448+v1697))) = v1677
	v1703 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1444+v1633<<(uint(v1703)%32)))) = uint16(v1621)
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1677)))
	v1755 = v1631 + v1703
	v1756 = v1634 + int32(base.Ui32(v1709)>>(uint(v1696)%32)) - int32(16)
	goto L281
L287:
	;
	v1683 = v1677 + int32(16)
	v1684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v1684 != int32(1) {
		v1695 = v1683
		goto L286
	} else {
		goto L288
	}
L288:
	;
	v1687 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)))
	switch v1687 - int32(1) {
	case 0:
		goto L291
	case 1:
		goto L290
	default:
		goto L285
	case 3:
		goto L289
	}
L289:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1683)))
	v1695 = v1692
	goto L286
L290:
	;
	v1691 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1683))))
	v1695 = v1691
	goto L286
L291:
	;
	v1690 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1683))))
	v1695 = v1690
	goto L286
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+256)) = base.I32_extend16_s(v1687)
	F_errmsg_internal(m, int32(506326), v56+int32(256))
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L3
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(343236), int32(70), int32(73749))
	mBase = m.M
	v1730 = m.ExcPending
	if v1730 != 0 {
		goto L3
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v1677)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+240)) = v1735 & int32(3)
	F_errmsg_internal(m, int32(506912), v56+int32(240))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L3
	} else {
		goto L296
	}
L296:
	;
	F_errfinish(m, int32(516634), int32(803), int32(109607))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L3
	} else {
		goto L297
	}
L297:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L298:
	;
	goto L280
L299:
	;
	F_UnlockReleaseBuffer(m, v325)
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L3
	} else {
		goto L300
	}
L300:
	;
	v7001 = int32(0)
	goto L27
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+336)) = v409
	F_errmsg_internal(m, int32(354911), v56+int32(336))
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		goto L3
	} else {
		goto L302
	}
L302:
	;
	F_errfinish(m, int32(516634), int32(2105), int32(87728))
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L3
	} else {
		goto L303
	}
L303:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L304:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(v56)+792))
	v1871 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1870+v1802<<(uint(v1871)%32)))) = v1869
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	*(*int32)(unsafe.Add(mBase, uint32(v1448+v1875<<(uint(v1871)%32)))) = v444
	v1880 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v56)+768)) = v1880
	*(*int64)(unsafe.Add(mBase, uint32(v56)+776)) = v1880
	*(*int64)(unsafe.Add(mBase, uint32(v56)+760)) = v1880
	v1887 = v1875 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+788)) = v1887
	if v60 == int32(0) {
		goto L316
	} else {
		goto L317
	}
L305:
	;
	v1841 = v444 + int32(16)
	v1842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v1842 != int32(1) {
		v1869 = v1841
		goto L304
	} else {
		goto L306
	}
L306:
	;
	v1845 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)))
	switch v1845 - int32(1) {
	case 0:
		goto L307
	case 1:
		goto L310
	default:
		goto L308
	case 3:
		goto L309
	}
L307:
	;
	v1866 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1841))))
	v1869 = v1866
	goto L304
L308:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L3
	} else {
		goto L311
	}
L309:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1841)))
	v1869 = v1849
	goto L304
L310:
	;
	v1848 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1841))))
	v1869 = v1848
	goto L304
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+192)) = base.I32_extend16_s(v1845)
	F_errmsg_internal(m, int32(506326), v56+int32(192))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L3
	} else {
		goto L312
	}
L312:
	;
	F_errfinish(m, int32(343236), int32(70), int32(73749))
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L3
	} else {
		goto L313
	}
L313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L314:
	;
	v2627 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	v2630 = F_palloc(m, v2627<<(uint(int32(2))%32))
	mBase = m.M
	v2631 = m.ExcPending
	if v2631 != 0 {
		goto L3
	} else {
		goto L364
	}
L315:
	;
	v2170 = int32(0)
	if v2128 < int32(2) {
		v2527 = v2170
		v2531 = v2128
		goto L340
	} else {
		goto L341
	}
L316:
	;
	v1891 = int32(1)
	v1894 = F_index_getprocinfo(m, l0, v1891, int32(3))
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L3
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	v2008 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+768)) = v2008
	v2013 = F_palloc0(m, v1887<<(uint(int32(2))%32))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L3
	} else {
		goto L330
	}
L319:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1896)))
	v1902 = F_FunctionCall2Coll(m, v1894, v1897, v56+int32(788), v56+int32(760))
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L3
	} else {
		goto L320
	}
L320:
	;
	v1904 = int32(0)
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	if v1907 <= v1904 {
		v2580 = v1904
		v2584 = v1907
		v2602 = v1891
		v2626 = v1904
		goto L314
	} else {
		goto L321
	}
L321:
	;
	v1917 = v1904
	v1919 = v1904
	goto L322
L322:
	;
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1963)))
	if int32(2) <= v1964 {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v2126 = v2003
	v2128 = v2006
	v2146 = v1891
	goto L315
L324:
	;
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v1448+v1917<<(uint(int32(2))%32))))
	F_spgDeformLeafTuple(m, v1970, v1963, v56+int32(624), v56+int32(592), int32(0))
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		goto L3
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	v1979 = v1917 << (uint(int32(2)) % 32)
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v56)+780))
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v1979+v1980)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+624)) = v1982
	v1984 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+592)) = uint8(v1984)
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v1448+v1979)))
	v1995 = F_spgFormLeafTuple(m, l1, v1988+int32(6), v56+int32(624), v56+int32(592))
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L3
	} else {
		goto L328
	}
L327:
	;
	goto L326
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1979+v1450))) = v1995
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v1995)))
	v2003 = v1919 + int32(base.Ui32(v1998)>>(uint(int32(2))%32)) + int32(4)
	v2005 = v1917 + int32(1)
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	if v2005 < v2006 {
		v1917 = v2005
		v1919 = v2003
		goto L322
	} else {
		goto L329
	}
L329:
	;
	goto L323
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+776)) = v2013
	v2016 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(v1875) {
		v2580 = v2016
		v2584 = v1887
		v2602 = v2008
		v2626 = v2016
		goto L314
	} else {
		goto L331
	}
L331:
	;
	v2028 = v2016
	v2030 = v2016
	goto L332
L332:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v2074)))
	if int32(2) <= v2075 {
		goto L334
	} else {
		goto L335
	}
L333:
	;
	v2126 = v2112
	v2128 = v2115
	v2146 = v2008
	goto L315
L334:
	;
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(v1448+v2028<<(uint(int32(2))%32))))
	F_spgDeformLeafTuple(m, v2081, v2074, v56+int32(624), v56+int32(592), int32(1))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L3
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+624)) = int32(0)
	v2091 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+592)) = uint8(v2091)
	v2094 = v2028 << (uint(int32(2)) % 32)
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v1448+v2094)))
	v2104 = F_spgFormLeafTuple(m, l1, v2097+int32(6), v56+int32(624), v56+int32(592))
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L3
	} else {
		goto L338
	}
L337:
	;
	goto L336
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1450+v2094))) = v2104
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v2104)))
	v2112 = v2030 + int32(base.Ui32(v2107)>>(uint(int32(2))%32)) + int32(4)
	v2114 = v2028 + int32(1)
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	if v2114 < v2115 {
		v2028 = v2114
		v2030 = v2112
		goto L332
	} else {
		goto L339
	}
L339:
	;
	goto L333
L340:
	;
	v2580 = v2527
	v2584 = v2531
	v2602 = v2146
	v2626 = v2126
	goto L314
L341:
	;
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v56)+776))
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v2173)))
	v2175 = int32(1)
	v2178 = v2128 - base.B2i32(base.Ui32(int32(8160)) < base.Ui32(v2126))
	if base.Ui32(v2178) <= base.Ui32(v2175) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	if base.Ui32(int32(8161)) <= base.Ui32(v2126) {
		goto L350
	} else {
		goto L351
	}
L343:
	;
	v2187 = v2175
	goto L344
L344:
	;
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v2173+v2187<<(uint(int32(2))%32))))
	if v2174 == v2237 {
		goto L346
	} else {
		goto L347
	}
L345:
	;
	v2580 = v2170
	v2584 = v2128
	v2602 = v2146
	v2626 = v2126
	goto L314
L346:
	;
	v2240 = v2187 + int32(1)
	if v2178 != v2240 {
		v2187 = v2240
		goto L344
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	goto L345
L349:
	;
	goto L342
L350:
	;
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(v2173+v2128<<(uint(int32(2))%32)-int32(4))))
	v2305 = base.B2i32(v2303 == v2174)
	goto L352
L351:
	;
	v2305 = int32(1)
	goto L352
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+768)) = int32(8)
	v2315 = int32(0)
	goto L353
L353:
	;
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v56)+776))
	v2366 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	v2367 = base.I32_rem_s(v2315, v2366)
	*(*int32)(unsafe.Add(mBase, uint32(v2362+v2315<<(uint(int32(2))%32)))) = v2367
	v2370 = v2315 + int32(1)
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	if v2370 < v2371 {
		v2315 = v2370
		goto L353
	} else {
		goto L355
	}
L354:
	;
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(v56)+772))
	if v2373 == int32(0) {
		goto L356
	} else {
		goto L357
	}
L355:
	;
	goto L354
L356:
	;
	v2505 = int32(4)
	v2506 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	if v2305 != 0 {
		v2527 = v2505
		v2531 = v2506
		goto L340
	} else {
		goto L363
	}
L357:
	;
	v2376 = int32(2)
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(v2373+v2174<<(uint(v2376)%32))))
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	v2383 = F_palloc(m, v2380<<(uint(v2376)%32))
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
		goto L3
	} else {
		goto L358
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+772)) = v2383
	v2386 = int32(0)
	v2387 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v2387 <= v2386 {
		goto L356
	} else {
		goto L359
	}
L359:
	;
	v2396 = v2386
	goto L360
L360:
	;
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(v56)+772))
	*(*int32)(unsafe.Add(mBase, uint32(v2443+v2396<<(uint(int32(2))%32)))) = v2379
	v2449 = v2396 + int32(1)
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v2449 < v2450 {
		v2396 = v2449
		goto L360
	} else {
		goto L362
	}
L361:
	;
	goto L356
L362:
	;
	goto L361
L363:
	;
	v2509 = v2506 - int32(1)
	v2510 = int32(2)
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(v1450+v2509<<(uint(v2510)%32))))
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(v2513)))
	v2580 = v2505
	v2584 = v2509
	v2602 = int32(0)
	v2626 = v2126 - int32(base.Ui32(v2514)>>(uint(v2510)%32)) - int32(4)
	goto L314
L364:
	;
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	v2635 = F_palloc0(m, v2632<<(uint(int32(2))%32))
	mBase = m.M
	v2636 = m.ExcPending
	if v2636 != 0 {
		goto L3
	} else {
		goto L365
	}
L365:
	;
	v2637 = int32(0)
	v2638 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v2637 < v2638 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v2647 = v2637
	goto L369
L367:
	;
	v2721 = v2638
	goto L368
L368:
	;
	v2766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+760)))
	v2767 = *(*int32)(unsafe.Add(mBase, uint32(v56)+764))
	v2768 = F_spgFormInnerTuple(m, l1, v2766, v2767, v2721, v2630)
	mBase = m.M
	v2769 = m.ExcPending
	if v2769 != 0 {
		goto L3
	} else {
		goto L376
	}
L369:
	;
	v2698 = *(*int32)(unsafe.Add(mBase, uint32(v56)+772))
	if v2698 != 0 {
		goto L371
	} else {
		goto L372
	}
L370:
	;
	v2721 = v2711
	goto L368
L371:
	;
	v2702 = *(*int32)(unsafe.Add(mBase, uint32(v2698+v2647<<(uint(int32(2))%32))))
	v2703 = v2702
	goto L373
L372:
	;
	v2703 = int32(0)
	goto L373
L373:
	;
	v2706 = F_spgFormNodeTuple(m, l1, v2703, base.B2i32(v2698 == int32(0)))
	mBase = m.M
	v2707 = m.ExcPending
	if v2707 != 0 {
		goto L3
	} else {
		goto L374
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2630+v2647<<(uint(int32(2))%32)))) = v2706
	v2710 = v2647 + int32(1)
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v2710 < v2711 {
		v2647 = v2710
		goto L369
	} else {
		goto L375
	}
L375:
	;
	goto L370
L376:
	;
	v2770 = *(*int32)(unsafe.Add(mBase, uint32(v2768)))
	*(*int32)(unsafe.Add(mBase, uint32(v2768))) = v2770&int32(-5) | v2580
	if v2770&int32(65528) != 0 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v2789 = v2768 + int32(base.Ui32(v2770)>>(uint(int32(16))%32)) + int32(8)
	v2790 = int32(0)
	goto L380
L378:
	;
	goto L379
L379:
	;
	v2905 = int32(0)
	if v2905 < v2584 {
		goto L386
	} else {
		goto L387
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2630+v2790<<(uint(int32(2))%32)))) = v2789
	v2840 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2789)+6)))
	v2841 = int32(8191)
	v2845 = v2790 + int32(1)
	v2846 = *(*int32)(unsafe.Add(mBase, uint32(v2768)))
	if base.Ui32(v2845) < base.Ui32(int32(base.Ui32(v2846)>>(uint(int32(3))%32))&v2841) {
		v2789 = v2789 + v2840&v2841
		v2790 = v2845
		goto L380
	} else {
		goto L382
	}
L381:
	;
	goto L379
L382:
	;
	goto L381
L383:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+575)) = uint8(v60)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+570)) = uint8(v1422)
	v3837 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+566)) = uint16(v3837)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+564)) = uint8(v662)
	v3841 = F_palloc(m, v2626)
	mBase = m.M
	v3842 = m.ExcPending
	if v3842 != 0 {
		goto L3
	} else {
		goto L492
	}
L384:
	;
	if v3517 <= int32(0) {
		v3797 = v3114
		v3799 = v3517
		v3811 = v3529
		goto L383
	} else {
		goto L481
	}
L385:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3490 = m.ExcPending
	if v3490 != 0 {
		goto L3
	} else {
		goto L478
	}
L386:
	;
	v2914 = v2905
	goto L389
L387:
	;
	goto L388
L388:
	;
	v3039 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+574)) = uint8(v3039)
	if v325 == v3039 {
		v3080 = v3039
		goto L394
	} else {
		goto L395
	}
L389:
	;
	v2962 = v2914 << (uint(int32(2)) % 32)
	v2963 = *(*int32)(unsafe.Add(mBase, uint32(v56)+776))
	v2965 = *(*int32)(unsafe.Add(mBase, uint32(v2962+v2963)))
	if v2965 < int32(0) {
		goto L385
	} else {
		goto L391
	}
L390:
	;
	goto L388
L391:
	;
	v2968 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v2968 <= v2965 {
		goto L385
	} else {
		goto L392
	}
L392:
	;
	v2970 = int32(2)
	v2972 = v2635 + v2965<<(uint(v2970)%32)
	v2973 = *(*int32)(unsafe.Add(mBase, uint32(v2972)))
	v2975 = *(*int32)(unsafe.Add(mBase, uint32(v2962+v1450)))
	v2976 = *(*int32)(unsafe.Add(mBase, uint32(v2975)))
	*(*int32)(unsafe.Add(mBase, uint32(v2972))) = v2973 + int32(base.Ui32(v2976)>>(uint(v2970)%32)) + int32(4)
	v2984 = v2914 + int32(1)
	if v2984 != v2584 {
		v2914 = v2984
		goto L389
	} else {
		goto L393
	}
L393:
	;
	goto L390
L394:
	;
	if v662 == int32(0) {
		goto L409
	} else {
		goto L410
	}
L395:
	;
	v3045 = int32(1)
	if base.Ui32(v326-v3045) <= base.Ui32(v3045) {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	v3075 = base.I32_rem_u_s(v326+int32(1), int32(3))
	v3077 = F_SpGistGetBuffer(m, l0, v3075|v293, v3070, v56+int32(574))
	mBase = m.M
	v3078 = m.ExcPending
	if v3078 != 0 {
		goto L3
	} else {
		goto L408
	}
L397:
	;
	v3049 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2768)+4)))
	v3070 = v3049 + int32(4)
	goto L396
L398:
	;
	goto L399
L399:
	;
	v3052 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v323)+14)))
	v3053 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v323)+12)))
	v3054 = v3052 - v3053
	v3055 = int32(0)
	if v3055 < v3054 {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	v3061 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v323)+16)))
	v3063 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v323+v3061)+4)))
	if v3063 != 0 {
		goto L404
	} else {
		goto L405
	}
L401:
	;
	v3058 = v3054
	goto L403
L402:
	;
	v3058 = v3055
	goto L403
L403:
	;
	goto L400
L404:
	;
	v3064 = int32(20)
	goto L406
L405:
	;
	v3064 = int32(0)
	goto L406
L406:
	;
	v3066 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2768)+4)))
	v3068 = v3066 + int32(4)
	if base.Ui32(v3068) <= base.Ui32(v3058+v3064) {
		v3080 = v325
		goto L394
	} else {
		goto L407
	}
L407:
	;
	v3070 = v3068
	goto L396
L408:
	;
	v3080 = v3077
	goto L394
L409:
	;
	v3083 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427)+14)))
	v3084 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427)+12)))
	v3085 = v3083 - v3084
	v3086 = int32(0)
	if v3086 < v3085 {
		goto L413
	} else {
		goto L414
	}
L410:
	;
	v3091 = v3039
	goto L411
L411:
	;
	v3092 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+571)) = uint8(v3092)
	if v2626 <= v3091 {
		goto L416
	} else {
		goto L417
	}
L412:
	;
	v3091 = v3089 + v1805
	goto L411
L413:
	;
	v3089 = v3085
	goto L415
L414:
	;
	v3089 = v3086
	goto L415
L415:
	;
	goto L412
L416:
	;
	v3095 = int32(0)
	v3096 = v1802 + v2602
	if v3096 <= v3095 {
		v3797 = v3095
		v3799 = v3096
		v3811 = v2602
		goto L383
	} else {
		goto L419
	}
L417:
	;
	goto L418
L418:
	;
	v3102 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	if v3102 != int32(1) {
		goto L421
	} else {
		goto L422
	}
L419:
	;
	v3101 = F__emscripten_memset_bulkmem(m, v1452, base.I32_extend8_s(int32(0)), v3096)
	mBase = m.M
	goto L420
L420:
	;
	v3797 = v3095
	v3799 = v3096
	v3811 = v2602
	goto L383
L421:
	;
	v3110 = int32(8160)
	if base.Ui32(v3110) <= base.Ui32(v2626) {
		goto L424
	} else {
		goto L425
	}
L422:
	;
	if base.Ui32(v2626) <= base.Ui32(int32(8160)) {
		goto L421
	} else {
		goto L423
	}
L423:
	;
	v3797 = int32(0)
	v3799 = v1802
	v3811 = int32(0)
	goto L383
L424:
	;
	v3113 = v3110
	goto L426
L425:
	;
	v3113 = v2626
	goto L426
L426:
	;
	v3114 = F_SpGistGetBuffer(m, l0, v296, v3113, v56+int32(571))
	mBase = m.M
	v3115 = m.ExcPending
	if v3115 != 0 {
		goto L3
	} else {
		goto L427
	}
L427:
	;
	v3116 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	v3117 = F_palloc(m, v3116)
	mBase = m.M
	v3118 = m.ExcPending
	if v3118 != 0 {
		goto L3
	} else {
		goto L428
	}
L428:
	;
	v3119 = int32(0)
	v3120 = base.B2i32(v3119 <= v3114)
	if v3120 == v3119 {
		goto L430
	} else {
		goto L431
	}
L429:
	;
	v3139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3138)+14)))
	v3140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3138)+12)))
	v3141 = v3139 - v3140
	v3142 = int32(0)
	if v3142 < v3141 {
		goto L434
	} else {
		goto L435
	}
L430:
	;
	v3124 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v3130 = *(*int32)(unsafe.Add(mBase, uint32(v3124+(v3114^int32(-1))<<(uint(int32(2))%32))))
	v3138 = v3130
	goto L429
L431:
	;
	goto L432
L432:
	;
	v3132 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v3138 = v3132 + v3114<<(uint(int32(13))%32) + int32(-8192)
	goto L429
L433:
	;
	v3146 = int32(0)
	v3147 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v3146 < v3147 {
		goto L437
	} else {
		goto L438
	}
L434:
	;
	v3145 = v3141
	goto L436
L435:
	;
	v3145 = v3142
	goto L436
L436:
	;
	goto L433
L437:
	;
	v3156 = v3146
	v3158 = v3091
	v3170 = v3145
	goto L440
L438:
	;
	v3231 = v3091
	v3243 = v3145
	goto L439
L439:
	;
	if v3231 < int32(0) {
		goto L447
	} else {
		goto L448
	}
L440:
	;
	v3203 = v3156 + v3117
	v3206 = v2635 + v3156<<(uint(int32(2))%32)
	v3207 = *(*int32)(unsafe.Add(mBase, uint32(v3206)))
	if v3207 <= v3158 {
		goto L443
	} else {
		goto L444
	}
L441:
	;
	v3231 = v3217
	v3243 = v3218
	goto L439
L442:
	;
	v3220 = v3156 + int32(1)
	v3221 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v3220 < v3221 {
		v3156 = v3220
		v3158 = v3217
		v3170 = v3218
		goto L440
	} else {
		goto L446
	}
L443:
	;
	v3209 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3203))) = uint8(v3209)
	v3211 = *(*int32)(unsafe.Add(mBase, uint32(v3206)))
	v3217 = v3158 - v3211
	v3218 = v3170
	goto L442
L444:
	;
	goto L445
L445:
	;
	v3213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3203))) = uint8(v3213)
	v3215 = *(*int32)(unsafe.Add(mBase, uint32(v3206)))
	v3217 = v3158
	v3218 = v3170 - v3215
	goto L442
L446:
	;
	goto L441
L447:
	;
	if v2602 != 0 {
		goto L450
	} else {
		goto L451
	}
L448:
	;
	if v3243 < int32(0) {
		goto L447
	} else {
		goto L449
	}
L449:
	;
	v3517 = v1802 + v2602
	v3529 = v2602
	goto L384
L450:
	;
	v3281 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	v3282 = int32(2)
	v3284 = int32(4)
	v3285 = v3281<<(uint(v3282)%32) - v3284
	v3286 = *(*int32)(unsafe.Add(mBase, uint32(v56)+776))
	v3288 = *(*int32)(unsafe.Add(mBase, uint32(v3285+v3286)))
	v3291 = v2635 + v3288<<(uint(v3282)%32)
	v3292 = *(*int32)(unsafe.Add(mBase, uint32(v3291)))
	v3294 = *(*int32)(unsafe.Add(mBase, uint32(v3285+v1450)))
	v3295 = *(*int32)(unsafe.Add(mBase, uint32(v3294)))
	*(*int32)(unsafe.Add(mBase, uint32(v3291))) = v3292 - int32(base.Ui32(v3295)>>(uint(v3282)%32)) - v3284
	if v3120 == int32(0) {
		goto L454
	} else {
		goto L455
	}
L451:
	;
	goto L452
L452:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3477 = m.ExcPending
	if v3477 != 0 {
		goto L3
	} else {
		goto L475
	}
L453:
	;
	v3320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3319)+14)))
	v3321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3319)+12)))
	v3322 = v3320 - v3321
	v3323 = int32(0)
	if v3323 < v3322 {
		goto L458
	} else {
		goto L459
	}
L454:
	;
	v3305 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v3311 = *(*int32)(unsafe.Add(mBase, uint32(v3305+(v3114^int32(-1))<<(uint(int32(2))%32))))
	v3319 = v3311
	goto L453
L455:
	;
	goto L456
L456:
	;
	v3313 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v3319 = v3313 + v3114<<(uint(int32(13))%32) + int32(-8192)
	goto L453
L457:
	;
	v3327 = int32(0)
	v3328 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v3327 < v3328 {
		goto L461
	} else {
		goto L462
	}
L458:
	;
	v3326 = v3322
	goto L460
L459:
	;
	v3326 = v3323
	goto L460
L460:
	;
	goto L457
L461:
	;
	v3337 = v3327
	v3342 = v3091
	v3343 = v3326
	goto L464
L462:
	;
	v3415 = v3091
	v3416 = v3326
	goto L463
L463:
	;
	v3457 = int32(0)
	if v3457 <= v3415|v3416 {
		v3517 = v1802
		v3529 = v3457
		goto L384
	} else {
		goto L471
	}
L464:
	;
	v3384 = v3337 + v3117
	v3387 = v2635 + v3337<<(uint(int32(2))%32)
	v3388 = *(*int32)(unsafe.Add(mBase, uint32(v3387)))
	if v3388 <= v3342 {
		goto L467
	} else {
		goto L468
	}
L465:
	;
	v3415 = v3398
	v3416 = v3399
	goto L463
L466:
	;
	v3401 = v3337 + int32(1)
	v3402 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v3401 < v3402 {
		v3337 = v3401
		v3342 = v3398
		v3343 = v3399
		goto L464
	} else {
		goto L470
	}
L467:
	;
	v3390 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3384))) = uint8(v3390)
	v3392 = *(*int32)(unsafe.Add(mBase, uint32(v3387)))
	v3398 = v3342 - v3392
	v3399 = v3343
	goto L466
L468:
	;
	goto L469
L469:
	;
	v3394 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3384))) = uint8(v3394)
	v3396 = *(*int32)(unsafe.Add(mBase, uint32(v3387)))
	v3398 = v3342
	v3399 = v3343 - v3396
	goto L466
L470:
	;
	goto L465
L471:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3464 = m.ExcPending
	if v3464 != 0 {
		goto L3
	} else {
		goto L472
	}
L472:
	;
	F_errmsg_internal(m, int32(181265), int32(0))
	mBase = m.M
	v3468 = m.ExcPending
	if v3468 != 0 {
		goto L3
	} else {
		goto L473
	}
L473:
	;
	F_errfinish(m, int32(516634), int32(1112), int32(109607))
	mBase = m.M
	v3473 = m.ExcPending
	if v3473 != 0 {
		goto L3
	} else {
		goto L474
	}
L474:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L475:
	;
	F_errmsg_internal(m, int32(181265), int32(0))
	mBase = m.M
	v3481 = m.ExcPending
	if v3481 != 0 {
		goto L3
	} else {
		goto L476
	}
L476:
	;
	F_errfinish(m, int32(516634), int32(1117), int32(109607))
	mBase = m.M
	v3486 = m.ExcPending
	if v3486 != 0 {
		goto L3
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
	F_errmsg_internal(m, int32(264735), int32(0))
	mBase = m.M
	v3494 = m.ExcPending
	if v3494 != 0 {
		goto L3
	} else {
		goto L479
	}
L479:
	;
	F_errfinish(m, int32(516634), int32(957), int32(109607))
	mBase = m.M
	v3499 = m.ExcPending
	if v3499 != 0 {
		goto L3
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
	v3556 = v3517 & int32(3)
	v3557 = int32(0)
	v3558 = *(*int32)(unsafe.Add(mBase, uint32(v56)+776))
	if base.Ui32(int32(4)) <= base.Ui32(v3517) {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	v3571 = v3557
	v3577 = int32(0)
	goto L485
L483:
	;
	v3667 = v3557
	goto L484
L484:
	;
	if v3556 == int32(0) {
		v3797 = v3114
		v3799 = v3517
		v3811 = v3529
		goto L383
	} else {
		goto L488
	}
L485:
	;
	v3619 = int32(2)
	v3622 = *(*int32)(unsafe.Add(mBase, uint32(v3558+v3571<<(uint(v3619)%32))))
	v3624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3117+v3622))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3571+v1452))) = uint8(v3624)
	v3627 = v3571 | int32(1)
	v3632 = *(*int32)(unsafe.Add(mBase, uint32(v3558+v3627<<(uint(v3619)%32))))
	v3634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3117+v3632))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1452+v3627))) = uint8(v3634)
	v3637 = v3571 | v3619
	v3642 = *(*int32)(unsafe.Add(mBase, uint32(v3558+v3637<<(uint(v3619)%32))))
	v3644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3117+v3642))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1452+v3637))) = uint8(v3644)
	v3647 = v3571 | int32(3)
	v3652 = *(*int32)(unsafe.Add(mBase, uint32(v3558+v3647<<(uint(v3619)%32))))
	v3654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3117+v3652))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1452+v3647))) = uint8(v3654)
	v3656 = int32(4)
	v3657 = v3571 + v3656
	v3659 = v3577 + v3656
	if v3659 != v3517&int32(2147483644) {
		v3571 = v3657
		v3577 = v3659
		goto L485
	} else {
		goto L487
	}
L486:
	;
	v3667 = v3657
	goto L484
L487:
	;
	goto L486
L488:
	;
	v3722 = v3667
	v3726 = v3557
	goto L489
L489:
	;
	v3773 = *(*int32)(unsafe.Add(mBase, uint32(v3558+v3722<<(uint(int32(2))%32))))
	v3775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3117+v3773))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3722+v1452))) = uint8(v3775)
	v3777 = int32(1)
	v3780 = v3726 + v3777
	if v3780 != v3556 {
		v3722 = v3722 + v3777
		v3726 = v3780
		goto L489
	} else {
		goto L491
	}
L490:
	;
	v3797 = v3114
	v3799 = v3517
	v3811 = v3529
	goto L383
L491:
	;
	goto L490
L492:
	;
	v3843 = int32(4556740)
	v3845 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v3845 + int32(1)
	v3849 = int32(0)
	if base.Ui32(v660) < base.Ui32(int32(2)) {
		v3925 = v3849
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v3926 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+756)) = v3926
	if v3926 < v3799 {
		goto L516
	} else {
		goto L517
	}
L494:
	;
	v3850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v3850 == int32(1) {
		goto L496
	} else {
		goto L497
	}
L495:
	;
	if v1422&int32(1) != 0 {
		v3925 = v3849
		goto L493
	} else {
		goto L514
	}
L496:
	;
	v3853 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427)+16)))
	v3855 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427+v3853)+4)))
	v3857 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v3857) {
		goto L499
	} else {
		goto L500
	}
L497:
	;
	goto L498
L498:
	;
	if v1422&int32(1) != 0 {
		v3925 = v3849
		goto L493
	} else {
		goto L508
	}
L499:
	;
	v3867 = int32(base.Ui32(v3857+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	goto L501
L500:
	;
	v3867 = int32(0)
	goto L501
L501:
	;
	if v1804+v3855 != v3867 {
		goto L495
	} else {
		goto L502
	}
L502:
	;
	if v408 < int32(0) {
		goto L505
	} else {
		goto L506
	}
L503:
	;
	v3895 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+570)) = uint8(v3895)
	v3925 = v3849
	goto L493
L504:
	;
	F_PageInit(m, v3886, int32(8192), int32(8))
	mBase = m.M
	v3890 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3886)+16)))
	v3891 = v3886 + v3890
	v3892 = int32(65410)
	*(*uint16)(unsafe.Add(mBase, uint32(v3891)+6)) = uint16(v3892)
	*(*uint16)(unsafe.Add(mBase, uint32(v3891))) = uint16(v290)
	goto L503
L505:
	;
	v3872 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v3878 = *(*int32)(unsafe.Add(mBase, uint32(v3872+(v408^int32(-1))<<(uint(int32(2))%32))))
	v3886 = v3878
	goto L504
L506:
	;
	goto L507
L507:
	;
	v3880 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v3886 = v3880 + v408<<(uint(int32(13))%32) + int32(-8192)
	goto L504
L508:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+566)) = uint16(v1804)
	if v1804 <= int32(0) {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	v3902 = int32(1)
	F_spgPageIndexMultiDelete(m, l1, v427, v1444, v1804, v3902, int32(3), int32(0), v3902)
	mBase = m.M
	v3907 = m.ExcPending
	if v3907 != 0 {
		goto L3
	} else {
		goto L512
	}
L510:
	;
	goto L511
L511:
	;
	v3908 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1444))))
	v3909 = int32(1)
	F_spgPageIndexMultiDelete(m, l1, v427, v1444, v1804, v3909, int32(3), int32(0), v3909)
	mBase = m.M
	v3914 = m.ExcPending
	if v3914 != 0 {
		goto L3
	} else {
		goto L513
	}
L512:
	;
	v3925 = v3849
	goto L493
L513:
	;
	v3925 = v3908
	goto L493
L514:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+566)) = uint16(v1804)
	v3918 = int32(3)
	F_spgPageIndexMultiDelete(m, l1, v427, v1444, v1804, v3918, v3918, int32(-1), int32(0))
	mBase = m.M
	v3923 = m.ExcPending
	if v3923 != 0 {
		goto L3
	} else {
		goto L515
	}
L515:
	;
	v3925 = v3849
	goto L493
L516:
	;
	v3936 = v3837
	v3941 = v3841
	goto L519
L517:
	;
	v4098 = v3841
	goto L518
L518:
	;
	if v3797 != 0 {
		goto L542
	} else {
		goto L543
	}
L519:
	;
	v3984 = v3936 << (uint(int32(2)) % 32)
	v3985 = v1450 + v3984
	v3986 = *(*int32)(unsafe.Add(mBase, uint32(v3985)))
	v3987 = v3936 + v1452
	v3988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3987))))
	if v3988 != 0 {
		goto L521
	} else {
		goto L522
	}
L520:
	;
	v4098 = v4083
	goto L518
L521:
	;
	v3989 = v3797
	goto L523
L522:
	;
	v3989 = v408
	goto L523
L523:
	;
	if v3989 < int32(0) {
		goto L525
	} else {
		goto L526
	}
L524:
	;
	v4009 = *(*int32)(unsafe.Add(mBase, uint32(v56)+776))
	v4011 = *(*int32)(unsafe.Add(mBase, uint32(v4009+v3984)))
	v4014 = v2630 + v4011<<(uint(int32(2))%32)
	v4015 = *(*int32)(unsafe.Add(mBase, uint32(v4014)))
	if v4015 == int32(0) {
		goto L529
	} else {
		goto L530
	}
L525:
	;
	v3993 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v3999 = *(*int32)(unsafe.Add(mBase, uint32(v3993+(v3989^int32(-1))<<(uint(int32(6))%32))+16))
	v4008 = v3999
	goto L524
L526:
	;
	goto L527
L527:
	;
	v4001 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v4007 = *(*int32)(unsafe.Add(mBase, uint32(v4001+v3989<<(uint(int32(6))%32)+int32(-64))+16))
	v4008 = v4007
	goto L524
L528:
	;
	if v3989 < int32(0) {
		goto L533
	} else {
		goto L534
	}
L529:
	;
	v4029 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3986)+4)))
	v4031 = v4029 & int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v3986)+4)) = uint16(v4031)
	goto L528
L530:
	;
	v4018 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4015)+4)))
	if v4018 == int32(0) {
		goto L529
	} else {
		goto L531
	}
L531:
	;
	v4021 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3986)+4)))
	v4026 = v4021&int32(49152) | v4018&int32(16383)
	*(*uint16)(unsafe.Add(mBase, uint32(v3986)+4)) = uint16(v4026)
	goto L528
L532:
	;
	v4055 = *(*int32)(unsafe.Add(mBase, uint32(v3986)))
	v4060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3987))))
	v4064 = F_SpGistPageAddNewItem(m, v4054, v3986, int32(base.Ui32(v4055)>>(uint(int32(2))%32)), v56+int32(756)+v4060<<(uint(int32(1))%32))
	mBase = m.M
	v4065 = m.ExcPending
	if v4065 != 0 {
		goto L3
	} else {
		goto L536
	}
L533:
	;
	v4040 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v4046 = *(*int32)(unsafe.Add(mBase, uint32(v4040+(v3989^int32(-1))<<(uint(int32(2))%32))))
	v4054 = v4046
	goto L532
L534:
	;
	goto L535
L535:
	;
	v4048 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v4054 = v4048 + v3989<<(uint(int32(13))%32) + int32(-8192)
	goto L532
L536:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1446+v3936<<(uint(int32(1))%32)))) = uint16(v4064)
	v4067 = *(*int32)(unsafe.Add(mBase, uint32(v4014)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4067)+4)) = uint16(v4064)
	*(*uint16)(unsafe.Add(mBase, uint32(v4067)+2)) = uint16(v4008)
	v4071 = int32(base.Ui32(v4008) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v4067))) = uint16(v4071)
	v4073 = *(*int32)(unsafe.Add(mBase, uint32(v3985)))
	v4074 = *(*int32)(unsafe.Add(mBase, uint32(v4073)))
	v4076 = int32(base.Ui32(v4074) >> (uint(int32(2)) % 32))
	if v4076 != 0 {
		goto L538
	} else {
		goto L539
	}
L537:
	;
	v4079 = *(*int32)(unsafe.Add(mBase, uint32(v3985)))
	v4080 = *(*int32)(unsafe.Add(mBase, uint32(v4079)))
	v4083 = v4078 + int32(base.Ui32(v4080)>>(uint(int32(2))%32))
	v4085 = v3936 + int32(1)
	if v4085 != v3799 {
		v3936 = v4085
		v3941 = v4083
		goto L519
	} else {
		goto L541
	}
L538:
	;
	v4077 = F__emscripten_memcpy_bulkmem(m, v3941, v4073, v4076)
	mBase = m.M
	v4078 = v4077
	goto L540
L539:
	;
	v4078 = v3941
	goto L540
L540:
	;
	goto L537
L541:
	;
	goto L520
L542:
	;
	F_MarkBufferDirty(m, v3797)
	mBase = m.M
	v4141 = m.ExcPending
	if v4141 != 0 {
		goto L3
	} else {
		goto L545
	}
L543:
	;
	goto L544
L544:
	;
	if v3080 == int32(0) {
		goto L547
	} else {
		goto L548
	}
L545:
	;
	goto L544
L546:
	;
	F_MarkBufferDirty(m, v408)
	mBase = m.M
	v4310 = m.ExcPending
	if v4310 != 0 {
		goto L3
	} else {
		goto L581
	}
L547:
	;
	if v325 != 0 {
		goto L554
	} else {
		goto L555
	}
L548:
	;
	if v325 != v3080 {
		goto L547
	} else {
		goto L549
	}
L549:
	;
	v4145 = *(*int32)(unsafe.Add(mBase, uint32(v56)+412))
	v4146 = *(*int32)(unsafe.Add(mBase, uint32(v56)+420))
	v4147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2768)+4)))
	v4149 = F_SpGistPageAddNewItem(m, v4146, v2768, v4147, int32(0))
	mBase = m.M
	v4150 = m.ExcPending
	if v4150 != 0 {
		goto L3
	} else {
		goto L550
	}
L550:
	;
	v4151 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+576)) = uint8(v4151)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+572)) = uint16(v4149)
	v4154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+424)))
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+578)) = uint16(v4154)
	v4156 = *(*int32)(unsafe.Add(mBase, uint32(v56)+428))
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+580)) = uint16(v4156)
	F_saveNodeLink(m, v56+int32(412), v4145, v4149)
	mBase = m.M
	v4161 = m.ExcPending
	if v4161 != 0 {
		goto L3
	} else {
		goto L551
	}
L551:
	;
	if v3925 == int32(0) {
		v4304 = v4146
		v4305 = v325
		v4306 = v4149
		v4307 = v4145
		v4308 = v408
		goto L546
	} else {
		goto L552
	}
L552:
	;
	v4167 = *(*int32)(unsafe.Add(mBase, uint32(v3925<<(uint(int32(2))%32)+v427)+20))
	v4170 = v427 + v4167&int32(32767)
	*(*uint16)(unsafe.Add(mBase, uint32(v4170)+10)) = uint16(v4149)
	*(*uint16)(unsafe.Add(mBase, uint32(v4170)+8)) = uint16(v4145)
	v4174 = int32(base.Ui32(v4145) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v4170)+6)) = uint16(v4174)
	v4304 = v4146
	v4305 = v325
	v4306 = v4149
	v4307 = v4145
	v4308 = v408
	goto L546
L553:
	;
	v4304 = v4212
	v4305 = v3080
	v4306 = v4215
	v4307 = v4194
	v4308 = v408
	goto L546
L554:
	;
	if v3080 < int32(0) {
		goto L558
	} else {
		goto L559
	}
L555:
	;
	goto L556
L556:
	;
	if v408 < int32(0) {
		goto L571
	} else {
		goto L572
	}
L557:
	;
	if v3080 < int32(0) {
		goto L562
	} else {
		goto L563
	}
L558:
	;
	v4179 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v4185 = *(*int32)(unsafe.Add(mBase, uint32(v4179+(v3080^int32(-1))<<(uint(int32(6))%32))+16))
	v4194 = v4185
	goto L557
L559:
	;
	goto L560
L560:
	;
	v4187 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v4193 = *(*int32)(unsafe.Add(mBase, uint32(v4187+v3080<<(uint(int32(6))%32)+int32(-64))+16))
	v4194 = v4193
	goto L557
L561:
	;
	v4213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2768)+4)))
	v4215 = F_SpGistPageAddNewItem(m, v4212, v2768, v4213, int32(0))
	mBase = m.M
	v4216 = m.ExcPending
	if v4216 != 0 {
		goto L3
	} else {
		goto L565
	}
L562:
	;
	v4198 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v4204 = *(*int32)(unsafe.Add(mBase, uint32(v4198+(v3080^int32(-1))<<(uint(int32(2))%32))))
	v4212 = v4204
	goto L561
L563:
	;
	goto L564
L564:
	;
	v4206 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v4212 = v4206 + v3080<<(uint(int32(13))%32) + int32(-8192)
	goto L561
L565:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+572)) = uint16(v4215)
	F_MarkBufferDirty(m, v3080)
	mBase = m.M
	v4219 = m.ExcPending
	if v4219 != 0 {
		goto L3
	} else {
		goto L566
	}
L566:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+576)) = uint8(base.B2i32(v325 == v3080))
	v4222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+424)))
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+578)) = uint16(v4222)
	v4224 = *(*int32)(unsafe.Add(mBase, uint32(v56)+428))
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+580)) = uint16(v4224)
	F_saveNodeLink(m, v56+int32(412), v4194, v4215)
	mBase = m.M
	v4229 = m.ExcPending
	if v4229 != 0 {
		goto L3
	} else {
		goto L567
	}
L567:
	;
	if v3925 == int32(0) {
		goto L553
	} else {
		goto L568
	}
L568:
	;
	v4235 = *(*int32)(unsafe.Add(mBase, uint32(v3925<<(uint(int32(2))%32)+v427)+20))
	v4238 = v427 + v4235&int32(32767)
	*(*uint16)(unsafe.Add(mBase, uint32(v4238)+10)) = uint16(v4215)
	*(*uint16)(unsafe.Add(mBase, uint32(v4238)+8)) = uint16(v4194)
	v4242 = int32(base.Ui32(v4194) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v4238)+6)) = uint16(v4242)
	goto L553
L569:
	;
	v4270 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+576)) = uint8(v4270)
	v4272 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+574)) = uint8(v4272)
	v4274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2768)+4)))
	v4277 = F_PageAddItemExtended(m, v427, v2768, v4274, v4270, v4270)
	mBase = m.M
	v4278 = m.ExcPending
	if v4278 != 0 {
		goto L3
	} else {
		goto L574
	}
L570:
	;
	F_PageInit(m, v4261, int32(8192), int32(8))
	mBase = m.M
	v4265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4261)+16)))
	v4266 = v4261 + v4265
	v4267 = int32(65410)
	*(*uint16)(unsafe.Add(mBase, uint32(v4266)+6)) = uint16(v4267)
	*(*uint16)(unsafe.Add(mBase, uint32(v4266))) = uint16(v287)
	goto L569
L571:
	;
	v4247 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v4253 = *(*int32)(unsafe.Add(mBase, uint32(v4247+(v408^int32(-1))<<(uint(int32(2))%32))))
	v4261 = v4253
	goto L570
L572:
	;
	goto L573
L573:
	;
	v4255 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v4261 = v4255 + v408<<(uint(int32(13))%32) + int32(-8192)
	goto L570
L574:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+572)) = uint16(v4277)
	if v4277 == int32(1) {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	v4282 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+578)) = v4282
	v4304 = v427
	v4305 = v408
	v4306 = int32(1)
	v4307 = v409
	v4308 = v4282
	goto L546
L576:
	;
	goto L577
L577:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4289 = m.ExcPending
	if v4289 != 0 {
		goto L3
	} else {
		goto L578
	}
L578:
	;
	v4290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2768)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+176)) = v4290
	F_errmsg_internal(m, int32(426927), v56+int32(176))
	mBase = m.M
	v4296 = m.ExcPending
	if v4296 != 0 {
		goto L3
	} else {
		goto L579
	}
L579:
	;
	F_errfinish(m, int32(516634), int32(1347), int32(109607))
	mBase = m.M
	v4301 = m.ExcPending
	if v4301 != 0 {
		goto L3
	} else {
		goto L580
	}
L580:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L581:
	;
	v4311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4311)+118)))
	if v4312 != int32(112) {
		goto L582
	} else {
		goto L583
	}
L582:
	;
	v4444 = int32(4556740)
	v4446 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v4446 - int32(1)
	if v3797 != 0 {
		goto L635
	} else {
		goto L636
	}
L583:
	;
	v4316 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v4316 <= int32(0) {
		goto L584
	} else {
		goto L585
	}
L584:
	;
	v4319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v4319 != 0 {
		goto L582
	} else {
		goto L587
	}
L585:
	;
	goto L586
L586:
	;
	v4321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v4321 != 0 {
		goto L582
	} else {
		goto L589
	}
L587:
	;
	v4320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4320 != 0 {
		goto L582
	} else {
		goto L588
	}
L588:
	;
	goto L586
L589:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v4323 = m.ExcPending
	if v4323 != 0 {
		goto L3
	} else {
		goto L590
	}
L590:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+568)) = uint16(v3799)
	F_XLogRegisterData(m, v56+int32(564), int32(28))
	mBase = m.M
	v4329 = m.ExcPending
	if v4329 != 0 {
		goto L3
	} else {
		goto L591
	}
L591:
	;
	v4330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+566)))
	F_XLogRegisterData(m, v1444, v4330<<(uint(int32(1))%32))
	mBase = m.M
	v4334 = m.ExcPending
	if v4334 != 0 {
		goto L3
	} else {
		goto L592
	}
L592:
	;
	v4335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+568)))
	F_XLogRegisterData(m, v1446, v4335<<(uint(int32(1))%32))
	mBase = m.M
	v4339 = m.ExcPending
	if v4339 != 0 {
		goto L3
	} else {
		goto L593
	}
L593:
	;
	v4340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+568)))
	F_XLogRegisterData(m, v1452, v4340)
	mBase = m.M
	v4342 = m.ExcPending
	if v4342 != 0 {
		goto L3
	} else {
		goto L594
	}
L594:
	;
	v4343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2768)+4)))
	F_XLogRegisterData(m, v2768, v4343)
	mBase = m.M
	v4345 = m.ExcPending
	if v4345 != 0 {
		goto L3
	} else {
		goto L595
	}
L595:
	;
	F_XLogRegisterData(m, v3841, v4098-v3841)
	mBase = m.M
	v4348 = m.ExcPending
	if v4348 != 0 {
		goto L3
	} else {
		goto L596
	}
L596:
	;
	if v4308 != 0 {
		goto L597
	} else {
		goto L598
	}
L597:
	;
	v4352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+570)))
	if v4352 != 0 {
		goto L600
	} else {
		goto L601
	}
L598:
	;
	goto L599
L599:
	;
	if v3797 != 0 {
		goto L604
	} else {
		goto L605
	}
L600:
	;
	v4353 = int32(14)
	goto L602
L601:
	;
	v4353 = int32(8)
	goto L602
L602:
	;
	F_XLogRegisterBuffer(m, int32(0), v4308, v4353)
	mBase = m.M
	v4355 = m.ExcPending
	if v4355 != 0 {
		goto L3
	} else {
		goto L603
	}
L603:
	;
	goto L599
L604:
	;
	v4359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+571)))
	if v4359 != 0 {
		goto L607
	} else {
		goto L608
	}
L605:
	;
	goto L606
L606:
	;
	v4366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+574)))
	if v4366 != 0 {
		goto L611
	} else {
		goto L612
	}
L607:
	;
	v4360 = int32(14)
	goto L609
L608:
	;
	v4360 = int32(8)
	goto L609
L609:
	;
	F_XLogRegisterBuffer(m, int32(1), v3797, v4360)
	mBase = m.M
	v4362 = m.ExcPending
	if v4362 != 0 {
		goto L3
	} else {
		goto L610
	}
L610:
	;
	goto L606
L611:
	;
	v4367 = int32(14)
	goto L613
L612:
	;
	v4367 = int32(8)
	goto L613
L613:
	;
	F_XLogRegisterBuffer(m, int32(2), v4305, v4367)
	mBase = m.M
	v4369 = m.ExcPending
	if v4369 != 0 {
		goto L3
	} else {
		goto L614
	}
L614:
	;
	v4370 = *(*int32)(unsafe.Add(mBase, uint32(v56)+416))
	if v4370 == int32(0) {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	v4380 = F_XLogInsert(m, int32(16), int32(80))
	mBase = m.M
	v4381 = m.ExcPending
	if v4381 != 0 {
		goto L3
	} else {
		goto L619
	}
L616:
	;
	if v4370 == v4305 {
		goto L615
	} else {
		goto L617
	}
L617:
	;
	F_XLogRegisterBuffer(m, int32(3), v4370, int32(8))
	mBase = m.M
	v4377 = m.ExcPending
	if v4377 != 0 {
		goto L3
	} else {
		goto L618
	}
L618:
	;
	goto L615
L619:
	;
	if v3797 != 0 {
		goto L620
	} else {
		goto L621
	}
L620:
	;
	if v3797 < int32(0) {
		goto L624
	} else {
		goto L625
	}
L621:
	;
	goto L622
L622:
	;
	if v4308 != 0 {
		goto L627
	} else {
		goto L628
	}
L623:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4399))) = base.I64_rotr(v4380, int64(32))
	goto L622
L624:
	;
	v4385 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v4391 = *(*int32)(unsafe.Add(mBase, uint32(v4385+(v3797^int32(-1))<<(uint(int32(2))%32))))
	v4399 = v4391
	goto L623
L625:
	;
	goto L626
L626:
	;
	v4393 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v4399 = v4393 + v3797<<(uint(int32(13))%32) + int32(-8192)
	goto L623
L627:
	;
	if v4308 < int32(0) {
		goto L631
	} else {
		goto L632
	}
L628:
	;
	v4430 = base.I32_wrap_i64(int64(base.Ui64(v4380) >> (uint(int64(32)) % 64)))
	goto L629
L629:
	;
	v4431 = base.I32_wrap_i64(v4380)
	*(*int32)(unsafe.Add(mBase, uint32(v4304)+4)) = v4431
	*(*int32)(unsafe.Add(mBase, uint32(v4304))) = v4430
	v4434 = *(*int32)(unsafe.Add(mBase, uint32(v56)+416))
	if v4434 == int32(0) {
		goto L582
	} else {
		goto L634
	}
L630:
	;
	v4421 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v4420))) = base.I64_rotr(v4380, v4421)
	v4430 = base.I32_wrap_i64(int64(base.Ui64(v4380) >> (uint(v4421) % 64)))
	goto L629
L631:
	;
	v4406 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v4412 = *(*int32)(unsafe.Add(mBase, uint32(v4406+(v4308^int32(-1))<<(uint(int32(2))%32))))
	v4420 = v4412
	goto L630
L632:
	;
	goto L633
L633:
	;
	v4414 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v4420 = v4414 + v4308<<(uint(int32(13))%32) + int32(-8192)
	goto L630
L634:
	;
	v4437 = *(*int32)(unsafe.Add(mBase, uint32(v56)+420))
	*(*int32)(unsafe.Add(mBase, uint32(v4437)+4)) = v4431
	*(*int32)(unsafe.Add(mBase, uint32(v4437))) = v4430
	goto L582
L635:
	;
	F_SpGistSetLastUsedPage(m, l0, v3797)
	mBase = m.M
	v4451 = m.ExcPending
	if v4451 != 0 {
		goto L3
	} else {
		goto L638
	}
L636:
	;
	goto L637
L637:
	;
	if v4308 != 0 {
		goto L640
	} else {
		goto L641
	}
L638:
	;
	F_UnlockReleaseBuffer(m, v3797)
	mBase = m.M
	v4453 = m.ExcPending
	if v4453 != 0 {
		goto L3
	} else {
		goto L639
	}
L639:
	;
	goto L637
L640:
	;
	F_SpGistSetLastUsedPage(m, l0, v4308)
	mBase = m.M
	v4455 = m.ExcPending
	if v4455 != 0 {
		goto L3
	} else {
		goto L643
	}
L641:
	;
	goto L642
L642:
	;
	if v3811 != 0 {
		goto L645
	} else {
		goto L646
	}
L643:
	;
	F_UnlockReleaseBuffer(m, v4308)
	mBase = m.M
	v4457 = m.ExcPending
	if v4457 != 0 {
		goto L3
	} else {
		goto L644
	}
L644:
	;
	goto L642
L645:
	;
	v6874 = int32(1)
	v6882 = v4305
	goto L29
L646:
	;
	goto L647
L647:
	;
	F_pfree(m, v444)
	mBase = m.M
	v4460 = m.ExcPending
	if v4460 != 0 {
		goto L3
	} else {
		goto L648
	}
L648:
	;
	v4473 = v4304
	v4475 = v4305
	v4483 = v4306
	v4514 = v4307
	goto L63
L649:
	;
	v4530 = v4473
	v4532 = v4475
	v4533 = v4514
	v4540 = v4483
	goto L661
L650:
	;
	v6659 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+411)) = uint8(v6659)
	v6662 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v6662 == v6659 {
		v317 = v6656
		v318 = v5135
		v319 = v6657
		v320 = v5137
		v321 = v4786
		v323 = v4530
		v325 = v4532
		v326 = v4533
		v333 = v4540
		v342 = v5138 + v342
		v344 = v5149
		goto L61
	} else {
		goto L997
	}
L651:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6643 = m.ExcPending
	if v6643 != 0 {
		goto L3
	} else {
		goto L994
	}
L652:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6630 = m.ExcPending
	if v6630 != 0 {
		goto L3
	} else {
		goto L991
	}
L653:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6614 = m.ExcPending
	if v6614 != 0 {
		goto L3
	} else {
		goto L988
	}
L654:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6598 = m.ExcPending
	if v6598 != 0 {
		goto L3
	} else {
		goto L985
	}
L655:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6580 = m.ExcPending
	if v6580 != 0 {
		goto L3
	} else {
		goto L982
	}
L656:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6567 = m.ExcPending
	if v6567 != 0 {
		goto L3
	} else {
		goto L979
	}
L657:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6554 = m.ExcPending
	if v6554 != 0 {
		goto L3
	} else {
		goto L976
	}
L658:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6538 = m.ExcPending
	if v6538 != 0 {
		goto L3
	} else {
		goto L973
	}
L659:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6525 = m.ExcPending
	if v6525 != 0 {
		goto L3
	} else {
		goto L970
	}
L660:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6512 = m.ExcPending
	if v6512 != 0 {
		goto L3
	} else {
		goto L967
	}
L661:
	;
	v4571 = int32(1)
	v4572 = v4533 - v4571
	v4574 = v4530 + int32(4)
	v4578 = base.I32_rem_u_s(v4533+v4571, int32(3))
	v4580 = v4540 & int32(65535)
	v4585 = v4580<<(uint(int32(2))%32) + v4530 + int32(20)
	goto L663
L662:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6496 = m.ExcPending
	if v6496 != 0 {
		goto L3
	} else {
		goto L964
	}
L663:
	;
	v4639 = *(*int32)(unsafe.Add(mBase, uint32(v4585)))
	v4640 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+388)) = v342
	v4642 = *(*int32)(unsafe.Add(mBase, uint32(v56)+432))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+384)) = v4642
	*(*int32)(unsafe.Add(mBase, uint32(v56)+380)) = v4640
	v4647 = v4530 + v4639&int32(32767)
	v4648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4647))))
	v4652 = int32(base.Ui32(v4648)>>(uint(int32(2))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+392)) = uint8(v4652)
	v4654 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+393)) = uint8(base.B2i32(base.Ui32(int32(65535)) < base.Ui32(v4654)))
	v4659 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	if base.Ui32(v4659) < base.Ui32(int32(65536)) {
		v4669 = int32(0)
		goto L665
	} else {
		goto L666
	}
L664:
	;
	goto L662
L665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+396)) = v4669
	v4671 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+400)) = int32(base.Ui32(v4671)>>(uint(int32(3))%32)) & int32(8191)
	v4677 = F_spgExtractNodeLabels(m, l1, v4647)
	mBase = m.M
	v4678 = m.ExcPending
	if v4678 != 0 {
		goto L3
	} else {
		goto L668
	}
L666:
	;
	v4663 = v4647 + int32(8)
	v4664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+50)))
	if v4664 != int32(1) {
		v4669 = v4663
		goto L665
	} else {
		goto L667
	}
L667:
	;
	v4667 = *(*int32)(unsafe.Add(mBase, uint32(v4663)))
	v4669 = v4667
	goto L665
L668:
	;
	v4679 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v56)+352)) = v4679
	*(*int64)(unsafe.Add(mBase, uint32(v56+int32(360)))) = v4679
	*(*int64)(unsafe.Add(mBase, uint32(v56+int32(368)))) = v4679
	*(*int64)(unsafe.Add(mBase, uint32(v56)+344)) = v4679
	*(*int32)(unsafe.Add(mBase, uint32(v56)+404)) = v4677
	if v60 == int32(0) {
		goto L670
	} else {
		goto L671
	}
L669:
	;
	v4703 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	if v4703&int32(4) == int32(0) {
		v4781 = v4702
		goto L674
	} else {
		goto L675
	}
L670:
	;
	v4690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v4691 = *(*int32)(unsafe.Add(mBase, uint32(v4690)))
	v4696 = F_FunctionCall2Coll(m, v100, v4691, v56+int32(380), v56+int32(344))
	mBase = m.M
	v4697 = m.ExcPending
	if v4697 != 0 {
		goto L3
	} else {
		goto L673
	}
L671:
	;
	goto L672
L672:
	;
	v4699 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+344)) = v4699
	v4702 = v4699
	goto L669
L673:
	;
	v4698 = *(*int32)(unsafe.Add(mBase, uint32(v56)+344))
	v4702 = v4698
	goto L669
L674:
	;
	if v4781 != int32(3) {
		goto L689
	} else {
		goto L690
	}
L675:
	;
	switch v4702 - int32(1) {
	case 0:
		goto L676
	case 1:
		goto L677
	default:
		v4781 = v4702
		goto L674
	}
L676:
	;
	v4723 = int32(4645584)
	v4724 = int64(0)
	v4731 = base.I64_extend_i32_s(int32(base.Ui32(v4703)>>(uint(int32(3))%32))&int32(8191) - int32(1))
	if base.Ui64(v4731) <= base.Ui64(v4724) {
		goto L682
	} else {
		goto L683
	}
L677:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4713 = m.ExcPending
	if v4713 != 0 {
		goto L3
	} else {
		goto L678
	}
L678:
	;
	F_errmsg_internal(m, int32(402807), int32(0))
	mBase = m.M
	v4717 = m.ExcPending
	if v4717 != 0 {
		goto L3
	} else {
		goto L679
	}
L679:
	;
	F_errfinish(m, int32(516634), int32(2212), int32(87728))
	mBase = m.M
	v4722 = m.ExcPending
	if v4722 != 0 {
		goto L3
	} else {
		goto L680
	}
L680:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L681:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v56)+348)) = uint32(v4778)
	v4780 = *(*int32)(unsafe.Add(mBase, uint32(v56)+344))
	v4781 = v4780
	goto L674
L682:
	;
	v4778 = v4724
	goto L681
L683:
	;
	goto L684
L684:
	;
	v4738 = v4731 - v4724
	v4740 = *(*int64)(unsafe.Add(mBase, _consts[92]))
	v4741 = *(*int64)(unsafe.Add(mBase, _consts[91]))
	v4744 = v4741
	v4746 = v4740
	goto L685
L685:
	;
	v4750 = v4744 ^ v4746
	v4752 = base.I64_rotl(v4750, int64(37))
	v4760 = v4750 ^ (v4750<<(uint(int64(16))%64) ^ base.I64_rotl(v4744, int64(24)))
	v4765 = int64(base.Ui64(base.I64_rotl(v4744*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v4738)) % 64))
	if base.Ui64(v4738) < base.Ui64(v4765) {
		v4744 = v4760
		v4746 = v4752
		goto L685
	} else {
		goto L687
	}
L686:
	;
	*(*int64)(unsafe.Add(mBase, _consts[92])) = v4752
	*(*int64)(unsafe.Add(mBase, _consts[91])) = v4760
	v4778 = v4724 + v4765
	goto L681
L687:
	;
	goto L686
L688:
	;
	goto L664
L689:
	;
	switch v4781 - int32(1) {
	case 0:
		goto L693
	case 1:
		goto L692
	default:
		goto L688
	}
L690:
	;
	goto L691
L691:
	;
	v5691 = *(*int32)(unsafe.Add(mBase, uint32(v56)+356))
	v5692 = int32(-8192)
	if base.Ui32(v5691+v5692) <= base.Ui32(v5692) {
		goto L654
	} else {
		goto L851
	}
L692:
	;
	v5195 = *(*int32)(unsafe.Add(mBase, uint32(v56)+404))
	if v5195 == int32(0) {
		goto L660
	} else {
		goto L740
	}
L693:
	;
	v4786 = *(*int32)(unsafe.Add(mBase, uint32(v56)+348))
	v4787 = *(*int32)(unsafe.Add(mBase, uint32(v56)+416))
	if v4787 == int32(0) {
		goto L694
	} else {
		goto L695
	}
L694:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+428)) = v4786
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+424)) = uint16(v4540)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+420)) = v4530
	*(*int32)(unsafe.Add(mBase, uint32(v56)+416)) = v4532
	*(*int32)(unsafe.Add(mBase, uint32(v56)+412)) = v4533
	v4800 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	v4805 = v4647 + int32(base.Ui32(v4800)>>(uint(int32(16))%32)) + int32(8)
	if v4786 == int32(0) {
		goto L700
	} else {
		goto L701
	}
L695:
	;
	if v4787 == v4532 {
		goto L694
	} else {
		goto L696
	}
L696:
	;
	F_SpGistSetLastUsedPage(m, l0, v4787)
	mBase = m.M
	v4792 = m.ExcPending
	if v4792 != 0 {
		goto L3
	} else {
		goto L697
	}
L697:
	;
	F_UnlockReleaseBuffer(m, v4787)
	mBase = m.M
	v4794 = m.ExcPending
	if v4794 != 0 {
		goto L3
	} else {
		goto L698
	}
L698:
	;
	goto L694
L699:
	;
	v5123 = int32(0)
	v5124 = int32(-1)
	if v5076 == v5123 {
		v5135 = v5123
		v5137 = v5124
		goto L720
	} else {
		goto L721
	}
L700:
	;
	if v4786 != 0 {
		goto L31
	} else {
		goto L719
	}
L701:
	;
	v4811 = int32(base.Ui32(v4800)>>(uint(int32(3))%32)) & int32(8191)
	if v4811 == int32(0) {
		goto L700
	} else {
		goto L702
	}
L702:
	;
	v4814 = int32(1)
	v4815 = v4786 - v4814
	v4817 = v4811 - v4814
	if base.Ui32(v4815) < base.Ui32(v4817) {
		goto L703
	} else {
		goto L704
	}
L703:
	;
	v4819 = v4815
	goto L705
L704:
	;
	v4819 = v4817
	goto L705
L705:
	;
	v4821 = v4819 + int32(1)
	v4822 = int32(3)
	v4823 = v4821 & v4822
	if base.Ui32(v4822) <= base.Ui32(v4819) {
		goto L706
	} else {
		goto L707
	}
L706:
	;
	v4835 = v4805
	v4836 = int32(0)
	goto L709
L707:
	;
	v4907 = v4805
	goto L708
L708:
	;
	if v4823 != 0 {
		goto L712
	} else {
		goto L713
	}
L709:
	;
	v4882 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4835)+6)))
	v4883 = int32(8191)
	v4885 = v4835 + v4882&v4883
	v4886 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4885)+6)))
	v4889 = v4885 + v4886&v4883
	v4890 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4889)+6)))
	v4893 = v4889 + v4890&v4883
	v4894 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4893)+6)))
	v4897 = v4893 + v4894&v4883
	v4899 = v4836 + int32(4)
	if v4899 != v4821&int32(-4) {
		v4835 = v4897
		v4836 = v4899
		goto L709
	} else {
		goto L711
	}
L710:
	;
	v4907 = v4897
	goto L708
L711:
	;
	goto L710
L712:
	;
	v4961 = v4907
	v4962 = int32(0)
	goto L715
L713:
	;
	v5021 = v4907
	goto L714
L714:
	;
	if v4786 == v4821 {
		v5076 = v5021
		goto L699
	} else {
		goto L718
	}
L715:
	;
	v5008 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4961)+6)))
	v5011 = v4961 + v5008&int32(8191)
	v5013 = v4962 + int32(1)
	if v5013 != v4823 {
		v4961 = v5011
		v4962 = v5013
		goto L715
	} else {
		goto L717
	}
L716:
	;
	v5021 = v5011
	goto L714
L717:
	;
	goto L716
L718:
	;
	goto L31
L719:
	;
	v5076 = v4805
	goto L699
L720:
	;
	v5138 = *(*int32)(unsafe.Add(mBase, uint32(v56)+352))
	if v60 == int32(0) {
		goto L723
	} else {
		goto L724
	}
L721:
	;
	v5127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5076)+4)))
	if v5127 == int32(0) {
		v5135 = v5123
		v5137 = v5124
		goto L720
	} else {
		goto L722
	}
L722:
	;
	v5130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5076)+2)))
	v5131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5076))))
	v5135 = v5127
	v5137 = v5130 | v5131<<(uint(int32(16))%32)
	goto L720
L723:
	;
	v5141 = *(*int32)(unsafe.Add(mBase, uint32(v56)+356))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+432)) = v5141
	v5145 = F_SpGistGetLeafTupleSize(m, v58, v56+int32(432), l4)
	mBase = m.M
	v5146 = m.ExcPending
	if v5146 != 0 {
		goto L3
	} else {
		goto L726
	}
L724:
	;
	v5149 = v344
	goto L725
L725:
	;
	if base.Ui32(v5149) < base.Ui32(int32(8161)) {
		goto L727
	} else {
		goto L728
	}
L726:
	;
	v5149 = v5145 + int32(4)
	goto L725
L727:
	;
	v6656 = v317
	v6657 = v319
	goto L650
L728:
	;
	goto L729
L729:
	;
	if v60 != 0 {
		goto L730
	} else {
		goto L731
	}
L730:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5168 = m.ExcPending
	if v5168 != 0 {
		goto L3
	} else {
		goto L735
	}
L731:
	;
	v5152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v5152&int32(1) == int32(0) {
		goto L730
	} else {
		goto L732
	}
L732:
	;
	if v5149 < v319 {
		v6656 = int32(0)
		v6657 = v5149
		goto L650
	} else {
		goto L733
	}
L733:
	;
	v5160 = v317 + int32(1)
	if v5160 < int32(10) {
		v6656 = v5160
		v6657 = v319
		goto L650
	} else {
		goto L734
	}
L734:
	;
	goto L730
L735:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v5171 = m.ExcPending
	if v5171 != 0 {
		goto L3
	} else {
		goto L736
	}
L736:
	;
	v5172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+36)) = int32(8156)
	v5175 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+32)) = v5149 - v5175
	*(*int32)(unsafe.Add(mBase, uint32(v56)+40)) = v5172 + v5175
	F_errmsg(m, int32(724654), v56+int32(32))
	mBase = m.M
	v5185 = m.ExcPending
	if v5185 != 0 {
		goto L3
	} else {
		goto L737
	}
L737:
	;
	F_errhint(m, int32(675672), int32(0))
	mBase = m.M
	v5189 = m.ExcPending
	if v5189 != 0 {
		goto L3
	} else {
		goto L738
	}
L738:
	;
	F_errfinish(m, int32(516634), int32(2282), int32(87728))
	mBase = m.M
	v5194 = m.ExcPending
	if v5194 != 0 {
		goto L3
	} else {
		goto L739
	}
L739:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L740:
	;
	v5198 = *(*int32)(unsafe.Add(mBase, uint32(v56)+348))
	v5199 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	v5203 = int32(base.Ui32(v5199)>>(uint(int32(3))%32)) & int32(8191)
	v5204 = *(*int32)(unsafe.Add(mBase, uint32(v56)+352))
	if v5204 < int32(0) {
		goto L742
	} else {
		goto L743
	}
L741:
	;
	v5210 = v4647 + int32(8)
	v5217 = F_palloc(m, int32(base.Ui32(v5199)>>(uint(int32(1))%32))&int32(32764)+int32(4))
	mBase = m.M
	v5218 = m.ExcPending
	if v5218 != 0 {
		goto L3
	} else {
		goto L746
	}
L742:
	;
	v5208 = v5203
	goto L741
L743:
	;
	goto L744
L744:
	;
	if base.Ui32(v5203) < base.Ui32(v5204) {
		goto L659
	} else {
		goto L745
	}
L745:
	;
	v5208 = v5204
	goto L741
L746:
	;
	v5219 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	if v5219&int32(65528) != 0 {
		goto L747
	} else {
		goto L748
	}
L747:
	;
	v5232 = v5210 + int32(base.Ui32(v5219)>>(uint(int32(16))%32))
	v5233 = int32(0)
	goto L750
L748:
	;
	goto L749
L749:
	;
	v5352 = int32(0)
	v5357 = F_spgFormNodeTuple(m, l1, v5198, v5352)
	mBase = m.M
	v5358 = m.ExcPending
	if v5358 != 0 {
		goto L3
	} else {
		goto L753
	}
L750:
	;
	v5280 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v5217+base.B2i32(v5208 <= v5233)<<(uint(v5280)%32)+v5233<<(uint(v5280)%32)))) = v5232
	v5287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5232)+6)))
	v5288 = int32(8191)
	v5292 = v5233 + int32(1)
	v5293 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	if base.Ui32(v5292) < base.Ui32(int32(base.Ui32(v5293)>>(uint(int32(3))%32))&v5288) {
		v5232 = v5232 + v5287&v5288
		v5233 = v5292
		goto L750
	} else {
		goto L752
	}
L751:
	;
	goto L749
L752:
	;
	goto L751
L753:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5217+v5208<<(uint(int32(2))%32)))) = v5357
	v5360 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	if base.Ui32(v5360) < base.Ui32(int32(65536)) {
		v5367 = v5352
		goto L754
	} else {
		goto L755
	}
L754:
	;
	v5376 = F_spgFormInnerTuple(m, l1, base.B2i32(base.Ui32(int32(65535)) < base.Ui32(v5360)), v5367, int32(base.Ui32(v5360)>>(uint(int32(3))%32))&int32(8191)+int32(1), v5217)
	mBase = m.M
	v5377 = m.ExcPending
	if v5377 != 0 {
		goto L3
	} else {
		goto L759
	}
L755:
	;
	v5363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+50)))
	if v5363 == int32(1) {
		goto L756
	} else {
		goto L757
	}
L756:
	;
	v5366 = *(*int32)(unsafe.Add(mBase, uint32(v5210)))
	v5367 = v5366
	goto L754
L757:
	;
	goto L758
L758:
	;
	v5367 = v5210
	goto L754
L759:
	;
	v5378 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+636)) = v5378
	v5380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+624)) = uint16(v4540)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+640)) = uint8(v5380)
	*(*int64)(unsafe.Add(mBase, uint32(v56)+626)) = int64(4278190080)
	v5385 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4530)+14)))
	v5386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4530)+12)))
	v5387 = v5385 - v5386
	v5388 = int32(0)
	if v5388 < v5387 {
		goto L762
	} else {
		goto L763
	}
L760:
	;
	v5686 = int32(0)
	v5688 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v5688 == v5686 {
		v4530 = v5680
		v4532 = v5681
		v4533 = v5682
		v4540 = v5683
		goto L661
	} else {
		goto L850
	}
L761:
	;
	v5392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5376)+4)))
	v5393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4647)+4)))
	if base.Ui32(v5392-v5393) <= base.Ui32(v5391) {
		goto L765
	} else {
		goto L766
	}
L762:
	;
	v5391 = v5387
	goto L764
L763:
	;
	v5391 = v5388
	goto L764
L764:
	;
	goto L761
L765:
	;
	v5396 = int32(4556740)
	v5398 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v5398 + int32(1)
	F_PageIndexTupleDelete(m, v4530, v4580)
	mBase = m.M
	v5403 = m.ExcPending
	if v5403 != 0 {
		goto L3
	} else {
		goto L768
	}
L766:
	;
	goto L767
L767:
	;
	if base.Ui32(v4572) <= base.Ui32(int32(1)) {
		goto L657
	} else {
		goto L785
	}
L768:
	;
	v5404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5376)+4)))
	v5406 = F_PageAddItemExtended(m, v4530, v5376, v5404, v4580, int32(0))
	mBase = m.M
	v5407 = m.ExcPending
	if v5407 != 0 {
		goto L3
	} else {
		goto L769
	}
L769:
	;
	if v5406 != v4580 {
		goto L658
	} else {
		goto L770
	}
L770:
	;
	F_MarkBufferDirty(m, v4532)
	mBase = m.M
	v5410 = m.ExcPending
	if v5410 != 0 {
		goto L3
	} else {
		goto L771
	}
L771:
	;
	v5411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5411)+118)))
	if v5412 != int32(112) {
		goto L772
	} else {
		goto L773
	}
L772:
	;
	v5443 = int32(4556740)
	v5445 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v5445 - int32(1)
	v5680 = v4530
	v5681 = v4532
	v5682 = v4533
	v5683 = v4540
	goto L760
L773:
	;
	v5416 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v5416 <= int32(0) {
		goto L774
	} else {
		goto L775
	}
L774:
	;
	v5419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v5419 != 0 {
		goto L772
	} else {
		goto L777
	}
L775:
	;
	goto L776
L776:
	;
	v5421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v5421 != 0 {
		goto L772
	} else {
		goto L779
	}
L777:
	;
	v5420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v5420 != 0 {
		goto L772
	} else {
		goto L778
	}
L778:
	;
	goto L776
L779:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v5423 = m.ExcPending
	if v5423 != 0 {
		goto L3
	} else {
		goto L780
	}
L780:
	;
	F_XLogRegisterData(m, v56+int32(624), int32(20))
	mBase = m.M
	v5428 = m.ExcPending
	if v5428 != 0 {
		goto L3
	} else {
		goto L781
	}
L781:
	;
	v5429 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5376)+4)))
	F_XLogRegisterData(m, v5376, v5429)
	mBase = m.M
	v5431 = m.ExcPending
	if v5431 != 0 {
		goto L3
	} else {
		goto L782
	}
L782:
	;
	F_XLogRegisterBuffer(m, int32(0), v4532, int32(8))
	mBase = m.M
	v5435 = m.ExcPending
	if v5435 != 0 {
		goto L3
	} else {
		goto L783
	}
L783:
	;
	v5438 = F_XLogInsert(m, int32(16), int32(48))
	mBase = m.M
	v5439 = m.ExcPending
	if v5439 != 0 {
		goto L3
	} else {
		goto L784
	}
L784:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4530))) = base.I64_rotr(v5438, int64(32))
	goto L772
L785:
	;
	v5451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+424)))
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+630)) = uint16(v5451)
	v5453 = *(*int32)(unsafe.Add(mBase, uint32(v56)+428))
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+632)) = uint16(v5453)
	v5456 = base.I32_rem_u_s(v4533, int32(3))
	v5457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5376)+4)))
	v5460 = F_SpGistGetBuffer(m, l0, v5456, v5457+int32(4), v298)
	mBase = m.M
	v5461 = m.ExcPending
	if v5461 != 0 {
		goto L3
	} else {
		goto L786
	}
L786:
	;
	if v5460 < int32(0) {
		goto L788
	} else {
		goto L789
	}
L787:
	;
	if v5460 < int32(0) {
		goto L792
	} else {
		goto L793
	}
L788:
	;
	v5465 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v5471 = *(*int32)(unsafe.Add(mBase, uint32(v5465+(v5460^int32(-1))<<(uint(int32(6))%32))+16))
	v5480 = v5471
	goto L787
L789:
	;
	goto L790
L790:
	;
	v5473 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v5479 = *(*int32)(unsafe.Add(mBase, uint32(v5473+v5460<<(uint(int32(6))%32)+int32(-64))+16))
	v5480 = v5479
	goto L787
L791:
	;
	if v5480 == v4533 {
		goto L656
	} else {
		goto L795
	}
L792:
	;
	v5484 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v5490 = *(*int32)(unsafe.Add(mBase, uint32(v5484+(v5460^int32(-1))<<(uint(int32(2))%32))))
	v5498 = v5490
	goto L791
L793:
	;
	goto L794
L794:
	;
	v5492 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v5498 = v5492 + v5460<<(uint(int32(13))%32) + int32(-8192)
	goto L791
L795:
	;
	v5500 = int32(4556740)
	v5502 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v5503 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v5502 + v5503
	v5508 = *(*int32)(unsafe.Add(mBase, uint32(v56)+416))
	if v5508 == v5460 {
		goto L796
	} else {
		goto L797
	}
L796:
	;
	v5510 = v5503
	goto L798
L797:
	;
	v5510 = int32(2)
	goto L798
L798:
	;
	if v5508 != v4532 {
		goto L799
	} else {
		goto L800
	}
L799:
	;
	v5513 = v5510
	goto L801
L800:
	;
	v5513 = int32(0)
	goto L801
L801:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+629)) = uint8(v5513)
	v5515 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5376)+4)))
	v5517 = F_SpGistPageAddNewItem(m, v5498, v5376, v5515, int32(0))
	mBase = m.M
	v5518 = m.ExcPending
	if v5518 != 0 {
		goto L3
	} else {
		goto L802
	}
L802:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+626)) = uint16(v5517)
	F_MarkBufferDirty(m, v5460)
	mBase = m.M
	v5521 = m.ExcPending
	if v5521 != 0 {
		goto L3
	} else {
		goto L803
	}
L803:
	;
	F_saveNodeLink(m, v56+int32(412), v5480, v5517)
	mBase = m.M
	v5525 = m.ExcPending
	if v5525 != 0 {
		goto L3
	} else {
		goto L804
	}
L804:
	;
	v5526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v5526 == int32(1) {
		goto L806
	} else {
		goto L807
	}
L805:
	;
	F_PageIndexTupleDelete(m, v4530, v4580)
	mBase = m.M
	v5587 = m.ExcPending
	if v5587 != 0 {
		goto L3
	} else {
		goto L817
	}
L806:
	;
	v5533 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v5533))) = int32(67)
	v5539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5533)+4)))
	v5541 = v5539 & int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v5533)+4)) = uint16(v5541)
	goto L811
L807:
	;
	goto L808
L808:
	;
	v5560 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v5560))) = int32(65)
	v5566 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5560)+4)))
	v5568 = v5566 & int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v5560)+4)) = uint16(v5568)
	goto L814
L809:
	;
	v5585 = v5533
	goto L805
L811:
	;
	goto L812
L812:
	;
	v5552 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5533)+10)) = uint16(v5552)
	*(*int32)(unsafe.Add(mBase, uint32(v5533)+6)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v5533)+12)) = v5552
	goto L809
L813:
	;
	v5585 = v5560
	goto L805
L814:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v5560)+10)) = uint16(v5517)
	*(*uint16)(unsafe.Add(mBase, uint32(v5560)+8)) = uint16(v5480)
	v5575 = int32(base.Ui32(v5480) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v5560)+6)) = uint16(v5575)
	v5577 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v5560)+12)) = v5577
	goto L813
L817:
	;
	v5588 = *(*int32)(unsafe.Add(mBase, uint32(v5585)))
	v5592 = F_PageAddItemExtended(m, v4530, v5585, int32(base.Ui32(v5588)>>(uint(int32(2))%32)), v4580, int32(0))
	mBase = m.M
	v5593 = m.ExcPending
	if v5593 != 0 {
		goto L3
	} else {
		goto L818
	}
L818:
	;
	if v5592 != v4580 {
		goto L655
	} else {
		goto L819
	}
L819:
	;
	v5595 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4530)+16)))
	v5599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v5599 != 0 {
		goto L820
	} else {
		goto L821
	}
L820:
	;
	v5600 = int32(4)
	goto L822
L821:
	;
	v5600 = int32(2)
	goto L822
L822:
	;
	v5601 = v4530 + v5595 + v5600
	v5602 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5601))))
	v5604 = v5602 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5601))) = uint16(v5604)
	F_MarkBufferDirty(m, v4532)
	mBase = m.M
	v5607 = m.ExcPending
	if v5607 != 0 {
		goto L3
	} else {
		goto L823
	}
L823:
	;
	v5608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5608)+118)))
	if v5609 != int32(112) {
		goto L824
	} else {
		goto L825
	}
L824:
	;
	v5663 = int32(4556740)
	v5665 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v5665 - int32(1)
	if v5460 == v4532 {
		goto L845
	} else {
		goto L846
	}
L825:
	;
	v5613 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v5613 <= int32(0) {
		goto L826
	} else {
		goto L827
	}
L826:
	;
	v5616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v5616 != 0 {
		goto L824
	} else {
		goto L829
	}
L827:
	;
	goto L828
L828:
	;
	v5618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v5618 != 0 {
		goto L824
	} else {
		goto L831
	}
L829:
	;
	v5617 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v5617 != 0 {
		goto L824
	} else {
		goto L830
	}
L830:
	;
	goto L828
L831:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v5620 = m.ExcPending
	if v5620 != 0 {
		goto L3
	} else {
		goto L832
	}
L832:
	;
	F_XLogRegisterBuffer(m, int32(0), v4532, int32(8))
	mBase = m.M
	v5624 = m.ExcPending
	if v5624 != 0 {
		goto L3
	} else {
		goto L833
	}
L833:
	;
	v5628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+628)))
	if v5628 != 0 {
		goto L834
	} else {
		goto L835
	}
L834:
	;
	v5629 = int32(14)
	goto L836
L835:
	;
	v5629 = int32(8)
	goto L836
L836:
	;
	F_XLogRegisterBuffer(m, int32(1), v5460, v5629)
	mBase = m.M
	v5631 = m.ExcPending
	if v5631 != 0 {
		goto L3
	} else {
		goto L837
	}
L837:
	;
	v5632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+629)))
	if v5632 == int32(2) {
		goto L838
	} else {
		goto L839
	}
L838:
	;
	v5636 = *(*int32)(unsafe.Add(mBase, uint32(v56)+416))
	F_XLogRegisterBuffer(m, int32(2), v5636, int32(8))
	mBase = m.M
	v5639 = m.ExcPending
	if v5639 != 0 {
		goto L3
	} else {
		goto L841
	}
L839:
	;
	goto L840
L840:
	;
	F_XLogRegisterData(m, v56+int32(624), int32(20))
	mBase = m.M
	v5644 = m.ExcPending
	if v5644 != 0 {
		goto L3
	} else {
		goto L842
	}
L841:
	;
	goto L840
L842:
	;
	v5645 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5376)+4)))
	F_XLogRegisterData(m, v5376, v5645)
	mBase = m.M
	v5647 = m.ExcPending
	if v5647 != 0 {
		goto L3
	} else {
		goto L843
	}
L843:
	;
	v5650 = F_XLogInsert(m, int32(16), int32(48))
	mBase = m.M
	v5651 = m.ExcPending
	if v5651 != 0 {
		goto L3
	} else {
		goto L844
	}
L844:
	;
	v5652 = int64(32)
	v5653 = base.I64_rotr(v5650, v5652)
	*(*int64)(unsafe.Add(mBase, uint32(v5498))) = v5653
	v5655 = *(*int32)(unsafe.Add(mBase, uint32(v56)+420))
	*(*int64)(unsafe.Add(mBase, uint32(v5655))) = v5653
	*(*uint32)(unsafe.Add(mBase, uint32(v4574))) = uint32(v5650)
	v5659 = int64(base.Ui64(v5650) >> (uint(v5652) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v4530))) = uint32(v5659)
	goto L824
L845:
	;
	v5680 = v5498
	v5681 = v5460
	v5682 = v5480
	v5683 = v5517
	goto L760
L846:
	;
	v5670 = *(*int32)(unsafe.Add(mBase, uint32(v56)+416))
	if v4532 == v5670 {
		goto L845
	} else {
		goto L847
	}
L847:
	;
	F_SpGistSetLastUsedPage(m, l0, v4532)
	mBase = m.M
	v5673 = m.ExcPending
	if v5673 != 0 {
		goto L3
	} else {
		goto L848
	}
L848:
	;
	F_UnlockReleaseBuffer(m, v4532)
	mBase = m.M
	v5675 = m.ExcPending
	if v5675 != 0 {
		goto L3
	} else {
		goto L849
	}
L849:
	;
	goto L845
L850:
	;
	v6874 = v5686
	v6882 = v5681
	goto L29
L851:
	;
	v5696 = *(*int32)(unsafe.Add(mBase, uint32(v56)+364))
	if base.Ui32(v5691) <= base.Ui32(v5696) {
		goto L653
	} else {
		goto L852
	}
L852:
	;
	v5700 = F_palloc(m, v5691<<(uint(int32(2))%32))
	mBase = m.M
	v5701 = m.ExcPending
	if v5701 != 0 {
		goto L3
	} else {
		goto L853
	}
L853:
	;
	v5702 = int32(0)
	v5703 = *(*int32)(unsafe.Add(mBase, uint32(v56)+356))
	if v5702 < v5703 {
		goto L854
	} else {
		goto L855
	}
L854:
	;
	v5712 = v5702
	goto L857
L855:
	;
	v5785 = v5703
	goto L856
L856:
	;
	v5831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+348)))
	v5832 = *(*int32)(unsafe.Add(mBase, uint32(v56)+352))
	v5833 = F_spgFormInnerTuple(m, l1, v5831, v5832, v5785, v5700)
	mBase = m.M
	v5834 = m.ExcPending
	if v5834 != 0 {
		goto L3
	} else {
		goto L864
	}
L857:
	;
	v5763 = *(*int32)(unsafe.Add(mBase, uint32(v56)+360))
	if v5763 != 0 {
		goto L859
	} else {
		goto L860
	}
L858:
	;
	v5785 = v5776
	goto L856
L859:
	;
	v5767 = *(*int32)(unsafe.Add(mBase, uint32(v5763+v5712<<(uint(int32(2))%32))))
	v5768 = v5767
	goto L861
L860:
	;
	v5768 = int32(0)
	goto L861
L861:
	;
	v5771 = F_spgFormNodeTuple(m, l1, v5768, base.B2i32(v5763 == int32(0)))
	mBase = m.M
	v5772 = m.ExcPending
	if v5772 != 0 {
		goto L3
	} else {
		goto L862
	}
L862:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5700+v5712<<(uint(int32(2))%32)))) = v5771
	v5775 = v5712 + int32(1)
	v5776 = *(*int32)(unsafe.Add(mBase, uint32(v56)+356))
	if v5775 < v5776 {
		v5712 = v5775
		goto L857
	} else {
		goto L863
	}
L863:
	;
	goto L858
L864:
	;
	v5835 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5833)+4)))
	v5836 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4647)+4)))
	if base.Ui32(v5836) < base.Ui32(v5835) {
		goto L652
	} else {
		goto L865
	}
L865:
	;
	v5838 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	v5843 = F_palloc(m, int32(base.Ui32(v5838)>>(uint(int32(1))%32))&int32(32764))
	mBase = m.M
	v5844 = m.ExcPending
	if v5844 != 0 {
		goto L3
	} else {
		goto L866
	}
L866:
	;
	v5845 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	if v5845&int32(65528) == int32(0) {
		goto L868
	} else {
		goto L869
	}
L867:
	;
	v5979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+368)))
	v5980 = *(*int32)(unsafe.Add(mBase, uint32(v56)+372))
	v5981 = F_spgFormInnerTuple(m, l1, v5979, v5980, v5936, v5843)
	mBase = m.M
	v5982 = m.ExcPending
	if v5982 != 0 {
		goto L3
	} else {
		goto L874
	}
L868:
	;
	v5936 = int32(0)
	goto L867
L869:
	;
	goto L870
L870:
	;
	v5863 = v4647 + int32(base.Ui32(v5845)>>(uint(int32(16))%32)) + int32(8)
	v5864 = int32(0)
	goto L871
L871:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5843+v5864<<(uint(int32(2))%32)))) = v5863
	v5914 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5863)+6)))
	v5915 = int32(8191)
	v5919 = v5864 + int32(1)
	v5920 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	v5924 = int32(base.Ui32(v5920)>>(uint(int32(3))%32)) & v5915
	if base.Ui32(v5919) < base.Ui32(v5924) {
		v5863 = v5863 + v5914&v5915
		v5864 = v5919
		goto L871
	} else {
		goto L873
	}
L872:
	;
	v5936 = v5924
	goto L867
L873:
	;
	goto L872
L874:
	;
	v5983 = *(*int32)(unsafe.Add(mBase, uint32(v5981)))
	v5986 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	*(*int32)(unsafe.Add(mBase, uint32(v5981))) = v5983&int32(-5) | v5986&int32(4)
	v5991 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+628)) = uint8(v5991)
	if base.Ui32(v4572) <= base.Ui32(int32(1)) {
		goto L877
	} else {
		goto L878
	}
L875:
	;
	v6027 = int32(4556740)
	v6029 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v6029 + int32(1)
	F_PageIndexTupleDelete(m, v4530, v4580)
	mBase = m.M
	v6034 = m.ExcPending
	if v6034 != 0 {
		goto L3
	} else {
		goto L889
	}
L876:
	;
	v6023 = F_SpGistGetBuffer(m, l0, v4578, v6019+int32(4), v298)
	mBase = m.M
	v6024 = m.ExcPending
	if v6024 != 0 {
		goto L3
	} else {
		goto L888
	}
L877:
	;
	v5995 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5981)+4)))
	v6019 = v5995
	goto L876
L878:
	;
	goto L879
L879:
	;
	v5996 = int32(0)
	v5997 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4530)+14)))
	v5998 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4530)+12)))
	v5999 = v5997 - v5998
	if v5996 < v5999 {
		goto L881
	} else {
		goto L882
	}
L880:
	;
	v6006 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4530)+16)))
	v6008 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4574+v6006))))
	if v6008 != 0 {
		goto L884
	} else {
		goto L885
	}
L881:
	;
	v6003 = v5999
	goto L883
L882:
	;
	v6003 = v5996
	goto L883
L883:
	;
	goto L880
L884:
	;
	v6009 = int32(20)
	goto L886
L885:
	;
	v6009 = int32(0)
	goto L886
L886:
	;
	v6011 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4647)+4)))
	v6013 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5981)+4)))
	v6014 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5833)+4)))
	if base.Ui32(v6013+v6014+int32(4)) <= base.Ui32(v6003+v6009+v6011) {
		v6026 = v5996
		goto L875
	} else {
		goto L887
	}
L887:
	;
	v6019 = v6013
	goto L876
L888:
	;
	v6026 = v6023
	goto L875
L889:
	;
	v6035 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5833)+4)))
	v6037 = F_PageAddItemExtended(m, v4530, v5833, v6035, v4580, int32(0))
	mBase = m.M
	v6038 = m.ExcPending
	if v6038 != 0 {
		goto L3
	} else {
		goto L890
	}
L890:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+624)) = uint16(v6037)
	if v6037 != v4580 {
		goto L651
	} else {
		goto L891
	}
L891:
	;
	if v6026 == int32(0) {
		goto L893
	} else {
		goto L894
	}
L892:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+629)) = uint8(v6094)
	v6098 = *(*int32)(unsafe.Add(mBase, uint32(v56)+364))
	v6099 = *(*int32)(unsafe.Add(mBase, uint32(v5833)))
	v6103 = int32(base.Ui32(v6099)>>(uint(int32(3))%32)) & int32(8191)
	if v6103 != 0 {
		goto L908
	} else {
		goto L909
	}
L893:
	;
	v6043 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5981)+4)))
	v6045 = F_SpGistPageAddNewItem(m, v4530, v5981, v6043, int32(0))
	mBase = m.M
	v6046 = m.ExcPending
	if v6046 != 0 {
		goto L3
	} else {
		goto L896
	}
L894:
	;
	goto L895
L895:
	;
	if v6026 < int32(0) {
		goto L898
	} else {
		goto L899
	}
L896:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+626)) = uint16(v6045)
	v6094 = int32(1)
	v6095 = v4533
	v6096 = v6045
	goto L892
L897:
	;
	v6068 = int32(0)
	if v6026 < v6068 {
		goto L902
	} else {
		goto L903
	}
L898:
	;
	v6052 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v6058 = *(*int32)(unsafe.Add(mBase, uint32(v6052+(v6026^int32(-1))<<(uint(int32(6))%32))+16))
	v6067 = v6058
	goto L897
L899:
	;
	goto L900
L900:
	;
	v6060 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v6066 = *(*int32)(unsafe.Add(mBase, uint32(v6060+v6026<<(uint(int32(6))%32)+int32(-64))+16))
	v6067 = v6066
	goto L897
L901:
	;
	v6087 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5981)+4)))
	v6089 = F_SpGistPageAddNewItem(m, v6086, v5981, v6087, int32(0))
	mBase = m.M
	v6090 = m.ExcPending
	if v6090 != 0 {
		goto L3
	} else {
		goto L905
	}
L902:
	;
	v6072 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v6078 = *(*int32)(unsafe.Add(mBase, uint32(v6072+(v6026^int32(-1))<<(uint(int32(2))%32))))
	v6086 = v6078
	goto L901
L903:
	;
	goto L904
L904:
	;
	v6080 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v6086 = v6080 + v6026<<(uint(int32(13))%32) + int32(-8192)
	goto L901
L905:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+626)) = uint16(v6089)
	F_MarkBufferDirty(m, v6026)
	mBase = m.M
	v6093 = m.ExcPending
	if v6093 != 0 {
		goto L3
	} else {
		goto L906
	}
L906:
	;
	v6094 = v6068
	v6095 = v6067
	v6096 = v6089
	goto L892
L907:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v6116)+4)) = uint16(v6096)
	*(*uint16)(unsafe.Add(mBase, uint32(v6116)+2)) = uint16(v6095)
	v6242 = int32(base.Ui32(v6095) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v6116))) = uint16(v6242)
	v6244 = *(*int32)(unsafe.Add(mBase, uint32(v56)+364))
	v6245 = *(*int32)(unsafe.Add(mBase, uint32(v4585)))
	v6248 = v4530 + v6245&int32(32767)
	v6249 = *(*int32)(unsafe.Add(mBase, uint32(v6248)))
	v6253 = int32(base.Ui32(v6249)>>(uint(int32(3))%32)) & int32(8191)
	if v6253 != 0 {
		goto L919
	} else {
		goto L920
	}
L908:
	;
	v6116 = v5833 + int32(base.Ui32(v6099)>>(uint(int32(16))%32)) + int32(8)
	v6117 = int32(0)
	goto L911
L909:
	;
	goto L910
L910:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6227 = m.ExcPending
	if v6227 != 0 {
		goto L3
	} else {
		goto L915
	}
L911:
	;
	if v6117 == v6098 {
		goto L907
	} else {
		goto L913
	}
L912:
	;
	goto L910
L913:
	;
	v6164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6116)+6)))
	v6169 = v6117 + int32(1)
	if v6169 != v6103 {
		v6116 = v6116 + v6164&int32(8191)
		v6117 = v6169
		goto L911
	} else {
		goto L914
	}
L914:
	;
	goto L912
L915:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+112)) = v6098
	F_errmsg_internal(m, int32(402967), v56+int32(112))
	mBase = m.M
	v6233 = m.ExcPending
	if v6233 != 0 {
		goto L3
	} else {
		goto L916
	}
L916:
	;
	F_errfinish(m, int32(516634), int32(68), int32(331387))
	mBase = m.M
	v6238 = m.ExcPending
	if v6238 != 0 {
		goto L3
	} else {
		goto L917
	}
L917:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L918:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v6266)+4)) = uint16(v6096)
	*(*uint16)(unsafe.Add(mBase, uint32(v6266)+2)) = uint16(v6095)
	*(*uint16)(unsafe.Add(mBase, uint32(v6266))) = uint16(v6242)
	F_MarkBufferDirty(m, v4532)
	mBase = m.M
	v6393 = m.ExcPending
	if v6393 != 0 {
		goto L3
	} else {
		goto L929
	}
L919:
	;
	v6266 = v6248 + int32(base.Ui32(v6249)>>(uint(int32(16))%32)) + int32(8)
	v6267 = int32(0)
	goto L922
L920:
	;
	goto L921
L921:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6377 = m.ExcPending
	if v6377 != 0 {
		goto L3
	} else {
		goto L926
	}
L922:
	;
	if v6267 == v6244 {
		goto L918
	} else {
		goto L924
	}
L923:
	;
	goto L921
L924:
	;
	v6314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6266)+6)))
	v6319 = v6267 + int32(1)
	if v6319 != v6253 {
		v6266 = v6266 + v6314&int32(8191)
		v6267 = v6319
		goto L922
	} else {
		goto L925
	}
L925:
	;
	goto L923
L926:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+128)) = v6244
	F_errmsg_internal(m, int32(402967), v56+int32(128))
	mBase = m.M
	v6383 = m.ExcPending
	if v6383 != 0 {
		goto L3
	} else {
		goto L927
	}
L927:
	;
	F_errfinish(m, int32(516634), int32(68), int32(331387))
	mBase = m.M
	v6388 = m.ExcPending
	if v6388 != 0 {
		goto L3
	} else {
		goto L928
	}
L928:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L929:
	;
	v6394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v6395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6394)+118)))
	if v6395 != int32(112) {
		goto L932
	} else {
		goto L933
	}
L930:
	;
	v6488 = int32(0)
	v6490 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v6490 == v6488 {
		goto L663
	} else {
		goto L963
	}
L931:
	;
	F_SpGistSetLastUsedPage(m, l0, v6026)
	mBase = m.M
	v6483 = m.ExcPending
	if v6483 != 0 {
		goto L3
	} else {
		goto L961
	}
L932:
	;
	v6472 = int32(4556740)
	v6474 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v6474 - int32(1)
	if v6026 == int32(0) {
		goto L930
	} else {
		goto L960
	}
L933:
	;
	v6399 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v6399 <= int32(0) {
		goto L934
	} else {
		goto L935
	}
L934:
	;
	v6402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v6402 != 0 {
		goto L932
	} else {
		goto L937
	}
L935:
	;
	goto L936
L936:
	;
	v6404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v6404 != 0 {
		goto L932
	} else {
		goto L939
	}
L937:
	;
	v6403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v6403 != 0 {
		goto L932
	} else {
		goto L938
	}
L938:
	;
	goto L936
L939:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v6406 = m.ExcPending
	if v6406 != 0 {
		goto L3
	} else {
		goto L940
	}
L940:
	;
	F_XLogRegisterData(m, v56+int32(624), int32(6))
	mBase = m.M
	v6411 = m.ExcPending
	if v6411 != 0 {
		goto L3
	} else {
		goto L941
	}
L941:
	;
	v6412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6248)+4)))
	F_XLogRegisterData(m, v6248, v6412)
	mBase = m.M
	v6414 = m.ExcPending
	if v6414 != 0 {
		goto L3
	} else {
		goto L942
	}
L942:
	;
	v6415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5981)+4)))
	F_XLogRegisterData(m, v5981, v6415)
	mBase = m.M
	v6417 = m.ExcPending
	if v6417 != 0 {
		goto L3
	} else {
		goto L943
	}
L943:
	;
	F_XLogRegisterBuffer(m, int32(0), v4532, int32(8))
	mBase = m.M
	v6421 = m.ExcPending
	if v6421 != 0 {
		goto L3
	} else {
		goto L944
	}
L944:
	;
	if v6026 != 0 {
		goto L945
	} else {
		goto L946
	}
L945:
	;
	v6425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+628)))
	if v6425 != 0 {
		goto L948
	} else {
		goto L949
	}
L946:
	;
	goto L947
L947:
	;
	v6431 = F_XLogInsert(m, int32(16), int32(64))
	mBase = m.M
	v6432 = m.ExcPending
	if v6432 != 0 {
		goto L3
	} else {
		goto L952
	}
L948:
	;
	v6426 = int32(14)
	goto L950
L949:
	;
	v6426 = int32(8)
	goto L950
L950:
	;
	F_XLogRegisterBuffer(m, int32(1), v6026, v6426)
	mBase = m.M
	v6428 = m.ExcPending
	if v6428 != 0 {
		goto L3
	} else {
		goto L951
	}
L951:
	;
	goto L947
L952:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4530))) = base.I64_rotr(v6431, int64(32))
	if v6026 == int32(0) {
		goto L953
	} else {
		goto L954
	}
L953:
	;
	v6438 = int32(4556740)
	v6440 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v6440 - int32(1)
	goto L930
L954:
	;
	goto L955
L955:
	;
	if v6026 < int32(0) {
		goto L957
	} else {
		goto L958
	}
L956:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v6461)+4)) = uint32(v6431)
	v6464 = int64(base.Ui64(v6431) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6461))) = uint32(v6464)
	v6466 = int32(4556740)
	v6468 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v6468 - int32(1)
	goto L931
L957:
	;
	v6447 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v6453 = *(*int32)(unsafe.Add(mBase, uint32(v6447+(v6026^int32(-1))<<(uint(int32(2))%32))))
	v6461 = v6453
	goto L956
L958:
	;
	goto L959
L959:
	;
	v6455 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v6461 = v6455 + v6026<<(uint(int32(13))%32) + int32(-8192)
	goto L956
L960:
	;
	goto L931
L961:
	;
	F_UnlockReleaseBuffer(m, v6026)
	mBase = m.M
	v6485 = m.ExcPending
	if v6485 != 0 {
		goto L3
	} else {
		goto L962
	}
L962:
	;
	goto L930
L963:
	;
	v6874 = v6488
	v6882 = v4532
	goto L29
L964:
	;
	v6497 = *(*int32)(unsafe.Add(mBase, uint32(v56)+344))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+16)) = v6497
	F_errmsg_internal(m, int32(504459), v56+int32(16))
	mBase = m.M
	v6503 = m.ExcPending
	if v6503 != 0 {
		goto L3
	} else {
		goto L965
	}
L965:
	;
	F_errfinish(m, int32(516634), int32(2318), int32(87728))
	mBase = m.M
	v6508 = m.ExcPending
	if v6508 != 0 {
		goto L3
	} else {
		goto L966
	}
L966:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L967:
	;
	F_errmsg_internal(m, int32(162848), int32(0))
	mBase = m.M
	v6516 = m.ExcPending
	if v6516 != 0 {
		goto L3
	} else {
		goto L968
	}
L968:
	;
	F_errfinish(m, int32(516634), int32(2295), int32(87728))
	mBase = m.M
	v6521 = m.ExcPending
	if v6521 != 0 {
		goto L3
	} else {
		goto L969
	}
L969:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L970:
	;
	F_errmsg_internal(m, int32(402854), int32(0))
	mBase = m.M
	v6529 = m.ExcPending
	if v6529 != 0 {
		goto L3
	} else {
		goto L971
	}
L971:
	;
	F_errfinish(m, int32(516634), int32(90), int32(434108))
	mBase = m.M
	v6534 = m.ExcPending
	if v6534 != 0 {
		goto L3
	} else {
		goto L972
	}
L972:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L973:
	;
	v6539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5376)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+80)) = v6539
	F_errmsg_internal(m, int32(426927), v56+int32(80))
	mBase = m.M
	v6545 = m.ExcPending
	if v6545 != 0 {
		goto L3
	} else {
		goto L974
	}
L974:
	;
	F_errfinish(m, int32(516634), int32(1553), int32(271419))
	mBase = m.M
	v6550 = m.ExcPending
	if v6550 != 0 {
		goto L3
	} else {
		goto L975
	}
L975:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L976:
	;
	F_errmsg_internal(m, int32(382833), int32(0))
	mBase = m.M
	v6558 = m.ExcPending
	if v6558 != 0 {
		goto L3
	} else {
		goto L977
	}
L977:
	;
	F_errfinish(m, int32(516634), int32(1588), int32(271419))
	mBase = m.M
	v6563 = m.ExcPending
	if v6563 != 0 {
		goto L3
	} else {
		goto L978
	}
L978:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L979:
	;
	F_errmsg_internal(m, int32(237278), int32(0))
	mBase = m.M
	v6571 = m.ExcPending
	if v6571 != 0 {
		goto L3
	} else {
		goto L980
	}
L980:
	;
	F_errfinish(m, int32(516634), int32(1616), int32(271419))
	mBase = m.M
	v6576 = m.ExcPending
	if v6576 != 0 {
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
	v6581 = *(*int32)(unsafe.Add(mBase, uint32(v5585)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+64)) = int32(base.Ui32(v6581) >> (uint(int32(2)) % 32))
	F_errmsg_internal(m, int32(426927), v56-int32(-64))
	mBase = m.M
	v6589 = m.ExcPending
	if v6589 != 0 {
		goto L3
	} else {
		goto L983
	}
L983:
	;
	F_errfinish(m, int32(516634), int32(1661), int32(271419))
	mBase = m.M
	v6594 = m.ExcPending
	if v6594 != 0 {
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
	v6599 = *(*int32)(unsafe.Add(mBase, uint32(v56)+356))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+96)) = v6599
	F_errmsg_internal(m, int32(504934), v56+int32(96))
	mBase = m.M
	v6605 = m.ExcPending
	if v6605 != 0 {
		goto L3
	} else {
		goto L986
	}
L986:
	;
	F_errfinish(m, int32(516634), int32(1736), int32(271381))
	mBase = m.M
	v6610 = m.ExcPending
	if v6610 != 0 {
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
	v6615 = *(*int32)(unsafe.Add(mBase, uint32(v56)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+160)) = v6615
	F_errmsg_internal(m, int32(505537), v56+int32(160))
	mBase = m.M
	v6621 = m.ExcPending
	if v6621 != 0 {
		goto L3
	} else {
		goto L989
	}
L989:
	;
	F_errfinish(m, int32(516634), int32(1741), int32(271381))
	mBase = m.M
	v6626 = m.ExcPending
	if v6626 != 0 {
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
	F_errmsg_internal(m, int32(28047), int32(0))
	mBase = m.M
	v6634 = m.ExcPending
	if v6634 != 0 {
		goto L3
	} else {
		goto L992
	}
L992:
	;
	F_errfinish(m, int32(516634), int32(1769), int32(271381))
	mBase = m.M
	v6639 = m.ExcPending
	if v6639 != 0 {
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
	v6644 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5833)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+144)) = v6644
	F_errmsg_internal(m, int32(426927), v56+int32(144))
	mBase = m.M
	v6650 = m.ExcPending
	if v6650 != 0 {
		goto L3
	} else {
		goto L995
	}
L995:
	;
	F_errfinish(m, int32(516634), int32(1825), int32(271381))
	mBase = m.M
	v6655 = m.ExcPending
	if v6655 != 0 {
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
	goto L62
L998:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v6726 = m.ExcPending
	if v6726 != 0 {
		goto L3
	} else {
		goto L999
	}
L999:
	;
	v6727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = int32(8156)
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v6727 + int32(4)
	F_errmsg(m, int32(724654), v56)
	mBase = m.M
	v6736 = m.ExcPending
	if v6736 != 0 {
		goto L3
	} else {
		goto L1000
	}
L1000:
	;
	F_errhint(m, int32(675672), int32(0))
	mBase = m.M
	v6740 = m.ExcPending
	if v6740 != 0 {
		goto L3
	} else {
		goto L1001
	}
L1001:
	;
	F_errfinish(m, int32(516634), int32(2005), int32(87728))
	mBase = m.M
	v6745 = m.ExcPending
	if v6745 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v56)+48)) = v4786
	F_errmsg_internal(m, int32(402967), v56+int32(48))
	mBase = m.M
	v6808 = m.ExcPending
	if v6808 != 0 {
		goto L3
	} else {
		goto L1004
	}
L1004:
	;
	F_errfinish(m, int32(516634), int32(1490), int32(271400))
	mBase = m.M
	v6813 = m.ExcPending
	if v6813 != 0 {
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
	v6934 = v6874
	v6942 = int32(0)
	goto L28
L1007:
	;
	goto L1008
L1008:
	;
	F_SpGistSetLastUsedPage(m, l0, v6882)
	mBase = m.M
	v6925 = m.ExcPending
	if v6925 != 0 {
		goto L3
	} else {
		goto L1009
	}
L1009:
	;
	F_UnlockReleaseBuffer(m, v6882)
	mBase = m.M
	v6927 = m.ExcPending
	if v6927 != 0 {
		goto L3
	} else {
		goto L1010
	}
L1010:
	;
	v6934 = v6874
	v6942 = v6882
	goto L28
L1011:
	;
	v6990 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v6990 == int32(0) {
		v7001 = v6934
		goto L27
	} else {
		goto L1016
	}
L1012:
	;
	if v6981 == v6942 {
		goto L1011
	} else {
		goto L1013
	}
L1013:
	;
	F_SpGistSetLastUsedPage(m, l0, v6981)
	mBase = m.M
	v6986 = m.ExcPending
	if v6986 != 0 {
		goto L3
	} else {
		goto L1014
	}
L1014:
	;
	F_UnlockReleaseBuffer(m, v6981)
	mBase = m.M
	v6988 = m.ExcPending
	if v6988 != 0 {
		goto L3
	} else {
		goto L1015
	}
L1015:
	;
	goto L1011
L1016:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v6994 = m.ExcPending
	if v6994 != 0 {
		goto L3
	} else {
		goto L1017
	}
L1017:
	;
	v7001 = v6934
	goto L27
}
