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
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int64
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v644 int64
	_ = v644
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v665 int32
	_ = v665
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v772 int32
	_ = v772
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v897 int32
	_ = v897
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v1002 int32
	_ = v1002
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1045 int32
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1285 int32
	_ = v1285
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1300 int32
	_ = v1300
	var v1306 int32
	_ = v1306
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1324 int64
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int64
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1411 int32
	_ = v1411
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1445 int32
	_ = v1445
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1469 int32
	_ = v1469
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1535 int32
	_ = v1535
	var v1541 int32
	_ = v1541
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1553 int32
	_ = v1553
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1573 int32
	_ = v1573
	var v1578 int32
	_ = v1578
	var v1585 int32
	_ = v1585
	var v1595 int32
	_ = v1595
	var v1599 int32
	_ = v1599
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1675 int32
	_ = v1675
	var v1681 int32
	_ = v1681
	var v1690 int32
	_ = v1690
	var v1697 int32
	_ = v1697
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1710 int32
	_ = v1710
	var v1715 int32
	_ = v1715
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1737 int32
	_ = v1737
	var v1743 int32
	_ = v1743
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1763 int32
	_ = v1763
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1820 int32
	_ = v1820
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1837 int32
	_ = v1837
	var v1842 int64
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1869 int32
	_ = v1869
	var v1879 int32
	_ = v1879
	var v1882 int32
	_ = v1882
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1932 int32
	_ = v1932
	var v1939 int32
	_ = v1939
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1960 int32
	_ = v1960
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1970 int32
	_ = v1970
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1990 int32
	_ = v1990
	var v1993 int32
	_ = v1993
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2043 int32
	_ = v2043
	var v2050 int32
	_ = v2050
	var v2053 int32
	_ = v2053
	var v2056 int32
	_ = v2056
	var v2059 int32
	_ = v2059
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2069 int32
	_ = v2069
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2092 int32
	_ = v2092
	var v2111 int32
	_ = v2111
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2139 int32
	_ = v2139
	var v2148 int32
	_ = v2148
	var v2198 int32
	_ = v2198
	var v2201 int32
	_ = v2201
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2276 int32
	_ = v2276
	var v2323 int32
	_ = v2323
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2334 int32
	_ = v2334
	var v2337 int32
	_ = v2337
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2357 int32
	_ = v2357
	var v2404 int32
	_ = v2404
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2494 int32
	_ = v2494
	var v2513 int32
	_ = v2513
	var v2534 int32
	_ = v2534
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2545 int32
	_ = v2545
	var v2554 int32
	_ = v2554
	var v2602 int32
	_ = v2602
	var v2604 int32
	_ = v2604
	var v2606 int32
	_ = v2606
	var v2608 int32
	_ = v2608
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2615 int32
	_ = v2615
	var v2616 int32
	_ = v2616
	var v2625 int32
	_ = v2625
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2675 int32
	_ = v2675
	var v2694 int32
	_ = v2694
	var v2695 int32
	_ = v2695
	var v2745 int32
	_ = v2745
	var v2746 int32
	_ = v2746
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2810 int32
	_ = v2810
	var v2819 int32
	_ = v2819
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2870 int32
	_ = v2870
	var v2873 int32
	_ = v2873
	var v2875 int32
	_ = v2875
	var v2877 int32
	_ = v2877
	var v2878 int32
	_ = v2878
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2889 int32
	_ = v2889
	var v2944 int32
	_ = v2944
	var v2950 int32
	_ = v2950
	var v2954 int32
	_ = v2954
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2959 int32
	_ = v2959
	var v2960 int32
	_ = v2960
	var v2963 int32
	_ = v2963
	var v2966 int32
	_ = v2966
	var v2968 int32
	_ = v2968
	var v2969 int32
	_ = v2969
	var v2971 int32
	_ = v2971
	var v2973 int32
	_ = v2973
	var v2975 int32
	_ = v2975
	var v2980 int32
	_ = v2980
	var v2982 int32
	_ = v2982
	var v2983 int32
	_ = v2983
	var v2985 int32
	_ = v2985
	var v2988 int32
	_ = v2988
	var v2989 int32
	_ = v2989
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v2994 int32
	_ = v2994
	var v2996 int32
	_ = v2996
	var v2997 int32
	_ = v2997
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3009 int32
	_ = v3009
	var v3017 int32
	_ = v3017
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3031 int32
	_ = v3031
	var v3037 int32
	_ = v3037
	var v3039 int32
	_ = v3039
	var v3045 int32
	_ = v3045
	var v3046 int32
	_ = v3046
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3049 int32
	_ = v3049
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3063 int32
	_ = v3063
	var v3065 int32
	_ = v3065
	var v3072 int32
	_ = v3072
	var v3110 int32
	_ = v3110
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3116 int32
	_ = v3116
	var v3118 int32
	_ = v3118
	var v3120 int32
	_ = v3120
	var v3122 int32
	_ = v3122
	var v3124 int32
	_ = v3124
	var v3125 int32
	_ = v3125
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3138 int32
	_ = v3138
	var v3145 int32
	_ = v3145
	var v3187 int32
	_ = v3187
	var v3188 int32
	_ = v3188
	var v3190 int32
	_ = v3190
	var v3191 int32
	_ = v3191
	var v3192 int32
	_ = v3192
	var v3194 int32
	_ = v3194
	var v3197 int32
	_ = v3197
	var v3198 int32
	_ = v3198
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3211 int32
	_ = v3211
	var v3217 int32
	_ = v3217
	var v3219 int32
	_ = v3219
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
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
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3258 int32
	_ = v3258
	var v3290 int32
	_ = v3290
	var v3293 int32
	_ = v3293
	var v3294 int32
	_ = v3294
	var v3296 int32
	_ = v3296
	var v3298 int32
	_ = v3298
	var v3300 int32
	_ = v3300
	var v3302 int32
	_ = v3302
	var v3304 int32
	_ = v3304
	var v3305 int32
	_ = v3305
	var v3307 int32
	_ = v3307
	var v3308 int32
	_ = v3308
	var v3317 int32
	_ = v3317
	var v3331 int32
	_ = v3331
	var v3363 int32
	_ = v3363
	var v3370 int32
	_ = v3370
	var v3374 int32
	_ = v3374
	var v3379 int32
	_ = v3379
	var v3383 int32
	_ = v3383
	var v3387 int32
	_ = v3387
	var v3392 int32
	_ = v3392
	var v3396 int32
	_ = v3396
	var v3400 int32
	_ = v3400
	var v3405 int32
	_ = v3405
	var v3426 int32
	_ = v3426
	var v3438 int32
	_ = v3438
	var v3462 int32
	_ = v3462
	var v3463 int32
	_ = v3463
	var v3464 int32
	_ = v3464
	var v3477 int32
	_ = v3477
	var v3492 int32
	_ = v3492
	var v3525 int32
	_ = v3525
	var v3528 int32
	_ = v3528
	var v3530 int32
	_ = v3530
	var v3533 int32
	_ = v3533
	var v3538 int32
	_ = v3538
	var v3540 int32
	_ = v3540
	var v3543 int32
	_ = v3543
	var v3548 int32
	_ = v3548
	var v3550 int32
	_ = v3550
	var v3553 int32
	_ = v3553
	var v3558 int32
	_ = v3558
	var v3560 int32
	_ = v3560
	var v3562 int32
	_ = v3562
	var v3563 int32
	_ = v3563
	var v3565 int32
	_ = v3565
	var v3575 int32
	_ = v3575
	var v3628 int32
	_ = v3628
	var v3635 int32
	_ = v3635
	var v3679 int32
	_ = v3679
	var v3681 int32
	_ = v3681
	var v3683 int32
	_ = v3683
	var v3686 int32
	_ = v3686
	var v3708 int32
	_ = v3708
	var v3710 int32
	_ = v3710
	var v3720 int32
	_ = v3720
	var v3743 int32
	_ = v3743
	var v3747 int32
	_ = v3747
	var v3748 int32
	_ = v3748
	var v3749 int32
	_ = v3749
	var v3751 int32
	_ = v3751
	var v3755 int32
	_ = v3755
	var v3756 int32
	_ = v3756
	var v3759 int32
	_ = v3759
	var v3761 int32
	_ = v3761
	var v3763 int32
	_ = v3763
	var v3773 int32
	_ = v3773
	var v3778 int32
	_ = v3778
	var v3784 int32
	_ = v3784
	var v3786 int32
	_ = v3786
	var v3792 int32
	_ = v3792
	var v3796 int32
	_ = v3796
	var v3797 int32
	_ = v3797
	var v3798 int32
	_ = v3798
	var v3801 int32
	_ = v3801
	var v3808 int32
	_ = v3808
	var v3813 int32
	_ = v3813
	var v3814 int32
	_ = v3814
	var v3815 int32
	_ = v3815
	var v3820 int32
	_ = v3820
	var v3824 int32
	_ = v3824
	var v3829 int32
	_ = v3829
	var v3830 int32
	_ = v3830
	var v3832 int32
	_ = v3832
	var v3842 int32
	_ = v3842
	var v3847 int32
	_ = v3847
	var v3890 int32
	_ = v3890
	var v3891 int32
	_ = v3891
	var v3892 int32
	_ = v3892
	var v3893 int32
	_ = v3893
	var v3894 int32
	_ = v3894
	var v3895 int32
	_ = v3895
	var v3899 int32
	_ = v3899
	var v3905 int32
	_ = v3905
	var v3907 int32
	_ = v3907
	var v3913 int32
	_ = v3913
	var v3914 int32
	_ = v3914
	var v3915 int32
	_ = v3915
	var v3917 int32
	_ = v3917
	var v3920 int32
	_ = v3920
	var v3921 int32
	_ = v3921
	var v3924 int32
	_ = v3924
	var v3927 int32
	_ = v3927
	var v3932 int32
	_ = v3932
	var v3935 int32
	_ = v3935
	var v3937 int32
	_ = v3937
	var v3946 int32
	_ = v3946
	var v3952 int32
	_ = v3952
	var v3954 int32
	_ = v3954
	var v3960 int32
	_ = v3960
	var v3961 int32
	_ = v3961
	var v3966 int32
	_ = v3966
	var v3970 int32
	_ = v3970
	var v3971 int32
	_ = v3971
	var v3973 int32
	_ = v3973
	var v3977 int32
	_ = v3977
	var v3979 int32
	_ = v3979
	var v3980 int32
	_ = v3980
	var v3982 int32
	_ = v3982
	var v3984 int32
	_ = v3984
	var v3985 int32
	_ = v3985
	var v3988 int32
	_ = v3988
	var v3990 int32
	_ = v3990
	var v4003 int32
	_ = v4003
	var v4046 int32
	_ = v4046
	var v4047 int32
	_ = v4047
	var v4053 int32
	_ = v4053
	var v4055 int32
	_ = v4055
	var v4056 int32
	_ = v4056
	var v4059 int32
	_ = v4059
	var v4065 int32
	_ = v4065
	var v4071 int32
	_ = v4071
	var v4074 int32
	_ = v4074
	var v4078 int32
	_ = v4078
	var v4083 int32
	_ = v4083
	var v4089 int32
	_ = v4089
	var v4091 int32
	_ = v4091
	var v4097 int32
	_ = v4097
	var v4098 int32
	_ = v4098
	var v4102 int32
	_ = v4102
	var v4108 int32
	_ = v4108
	var v4110 int32
	_ = v4110
	var v4116 int32
	_ = v4116
	var v4117 int32
	_ = v4117
	var v4119 int32
	_ = v4119
	var v4120 int32
	_ = v4120
	var v4123 int32
	_ = v4123
	var v4131 int32
	_ = v4131
	var v4137 int32
	_ = v4137
	var v4140 int32
	_ = v4140
	var v4144 int32
	_ = v4144
	var v4149 int32
	_ = v4149
	var v4155 int32
	_ = v4155
	var v4157 int32
	_ = v4157
	var v4163 int32
	_ = v4163
	var v4167 int32
	_ = v4167
	var v4168 int32
	_ = v4168
	var v4169 int32
	_ = v4169
	var v4172 int32
	_ = v4172
	var v4174 int32
	_ = v4174
	var v4176 int32
	_ = v4176
	var v4179 int32
	_ = v4179
	var v4180 int32
	_ = v4180
	var v4184 int32
	_ = v4184
	var v4191 int32
	_ = v4191
	var v4192 int32
	_ = v4192
	var v4198 int32
	_ = v4198
	var v4203 int32
	_ = v4203
	var v4206 int32
	_ = v4206
	var v4207 int32
	_ = v4207
	var v4208 int32
	_ = v4208
	var v4209 int32
	_ = v4209
	var v4210 int32
	_ = v4210
	var v4212 int32
	_ = v4212
	var v4213 int32
	_ = v4213
	var v4214 int32
	_ = v4214
	var v4218 int32
	_ = v4218
	var v4221 int32
	_ = v4221
	var v4222 int32
	_ = v4222
	var v4223 int32
	_ = v4223
	var v4225 int32
	_ = v4225
	var v4231 int32
	_ = v4231
	var v4232 int32
	_ = v4232
	var v4236 int32
	_ = v4236
	var v4237 int32
	_ = v4237
	var v4241 int32
	_ = v4241
	var v4242 int32
	_ = v4242
	var v4244 int32
	_ = v4244
	var v4245 int32
	_ = v4245
	var v4247 int32
	_ = v4247
	var v4250 int32
	_ = v4250
	var v4254 int32
	_ = v4254
	var v4255 int32
	_ = v4255
	var v4257 int32
	_ = v4257
	var v4261 int32
	_ = v4261
	var v4262 int32
	_ = v4262
	var v4264 int32
	_ = v4264
	var v4268 int32
	_ = v4268
	var v4269 int32
	_ = v4269
	var v4271 int32
	_ = v4271
	var v4272 int32
	_ = v4272
	var v4281 int32
	_ = v4281
	var v4284 int64
	_ = v4284
	var v4285 int32
	_ = v4285
	var v4289 int32
	_ = v4289
	var v4295 int32
	_ = v4295
	var v4297 int32
	_ = v4297
	var v4303 int32
	_ = v4303
	var v4310 int32
	_ = v4310
	var v4316 int32
	_ = v4316
	var v4318 int32
	_ = v4318
	var v4324 int32
	_ = v4324
	var v4330 int32
	_ = v4330
	var v4331 int32
	_ = v4331
	var v4341 int32
	_ = v4341
	var v4343 int32
	_ = v4343
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
	var v4367 int32
	_ = v4367
	var v4368 int32
	_ = v4368
	var v4373 int32
	_ = v4373
	var v4376 int32
	_ = v4376
	var v4413 int32
	_ = v4413
	var v4423 int32
	_ = v4423
	var v4424 int32
	_ = v4424
	var v4429 int32
	_ = v4429
	var v4432 int32
	_ = v4432
	var v4467 int32
	_ = v4467
	var v4468 int32
	_ = v4468
	var v4472 int32
	_ = v4472
	var v4474 int32
	_ = v4474
	var v4479 int32
	_ = v4479
	var v4533 int32
	_ = v4533
	var v4534 int32
	_ = v4534
	var v4536 int32
	_ = v4536
	var v4541 int32
	_ = v4541
	var v4542 int32
	_ = v4542
	var v4546 int32
	_ = v4546
	var v4548 int32
	_ = v4548
	var v4553 int32
	_ = v4553
	var v4557 int32
	_ = v4557
	var v4558 int32
	_ = v4558
	var v4561 int32
	_ = v4561
	var v4563 int32
	_ = v4563
	var v4565 int32
	_ = v4565
	var v4571 int32
	_ = v4571
	var v4572 int32
	_ = v4572
	var v4573 int64
	_ = v4573
	var v4584 int32
	_ = v4584
	var v4585 int32
	_ = v4585
	var v4590 int32
	_ = v4590
	var v4591 int32
	_ = v4591
	var v4592 int32
	_ = v4592
	var v4593 int32
	_ = v4593
	var v4596 int32
	_ = v4596
	var v4597 int32
	_ = v4597
	var v4607 int32
	_ = v4607
	var v4611 int32
	_ = v4611
	var v4616 int32
	_ = v4616
	var v4617 int32
	_ = v4617
	var v4618 int64
	_ = v4618
	var v4625 int64
	_ = v4625
	var v4632 int64
	_ = v4632
	var v4634 int64
	_ = v4634
	var v4635 int64
	_ = v4635
	var v4638 int64
	_ = v4638
	var v4640 int64
	_ = v4640
	var v4644 int64
	_ = v4644
	var v4646 int64
	_ = v4646
	var v4654 int64
	_ = v4654
	var v4659 int64
	_ = v4659
	var v4672 int64
	_ = v4672
	var v4674 int32
	_ = v4674
	var v4675 int32
	_ = v4675
	var v4680 int32
	_ = v4680
	var v4681 int32
	_ = v4681
	var v4688 int32
	_ = v4688
	var v4690 int32
	_ = v4690
	var v4696 int32
	_ = v4696
	var v4701 int32
	_ = v4701
	var v4707 int32
	_ = v4707
	var v4710 int32
	_ = v4710
	var v4711 int32
	_ = v4711
	var v4713 int32
	_ = v4713
	var v4715 int32
	_ = v4715
	var v4717 int32
	_ = v4717
	var v4718 int32
	_ = v4718
	var v4719 int32
	_ = v4719
	var v4731 int32
	_ = v4731
	var v4732 int32
	_ = v4732
	var v4778 int32
	_ = v4778
	var v4779 int32
	_ = v4779
	var v4781 int32
	_ = v4781
	var v4782 int32
	_ = v4782
	var v4785 int32
	_ = v4785
	var v4786 int32
	_ = v4786
	var v4789 int32
	_ = v4789
	var v4790 int32
	_ = v4790
	var v4793 int32
	_ = v4793
	var v4795 int32
	_ = v4795
	var v4805 int32
	_ = v4805
	var v4859 int32
	_ = v4859
	var v4860 int32
	_ = v4860
	var v4906 int32
	_ = v4906
	var v4909 int32
	_ = v4909
	var v4911 int32
	_ = v4911
	var v4919 int32
	_ = v4919
	var v4974 int32
	_ = v4974
	var v5021 int32
	_ = v5021
	var v5022 int32
	_ = v5022
	var v5023 int32
	_ = v5023
	var v5028 int32
	_ = v5028
	var v5029 int32
	_ = v5029
	var v5032 int32
	_ = v5032
	var v5036 int32
	_ = v5036
	var v5037 int32
	_ = v5037
	var v5040 int32
	_ = v5040
	var v5043 int32
	_ = v5043
	var v5051 int32
	_ = v5051
	var v5059 int32
	_ = v5059
	var v5062 int32
	_ = v5062
	var v5063 int32
	_ = v5063
	var v5066 int32
	_ = v5066
	var v5076 int32
	_ = v5076
	var v5080 int32
	_ = v5080
	var v5085 int32
	_ = v5085
	var v5086 int32
	_ = v5086
	var v5089 int32
	_ = v5089
	var v5090 int32
	_ = v5090
	var v5094 int32
	_ = v5094
	var v5095 int32
	_ = v5095
	var v5099 int32
	_ = v5099
	var v5101 int32
	_ = v5101
	var v5108 int32
	_ = v5108
	var v5109 int32
	_ = v5109
	var v5110 int32
	_ = v5110
	var v5123 int32
	_ = v5123
	var v5124 int32
	_ = v5124
	var v5172 int32
	_ = v5172
	var v5176 int32
	_ = v5176
	var v5177 int32
	_ = v5177
	var v5181 int32
	_ = v5181
	var v5182 int32
	_ = v5182
	var v5241 int32
	_ = v5241
	var v5246 int32
	_ = v5246
	var v5247 int32
	_ = v5247
	var v5249 int32
	_ = v5249
	var v5251 int32
	_ = v5251
	var v5254 int32
	_ = v5254
	var v5257 int32
	_ = v5257
	var v5258 int32
	_ = v5258
	var v5267 int32
	_ = v5267
	var v5268 int32
	_ = v5268
	var v5269 int32
	_ = v5269
	var v5271 int32
	_ = v5271
	var v5276 int32
	_ = v5276
	var v5277 int32
	_ = v5277
	var v5278 int32
	_ = v5278
	var v5279 int32
	_ = v5279
	var v5282 int32
	_ = v5282
	var v5283 int32
	_ = v5283
	var v5284 int32
	_ = v5284
	var v5287 int32
	_ = v5287
	var v5289 int32
	_ = v5289
	var v5294 int32
	_ = v5294
	var v5295 int32
	_ = v5295
	var v5297 int32
	_ = v5297
	var v5298 int32
	_ = v5298
	var v5301 int32
	_ = v5301
	var v5302 int32
	_ = v5302
	var v5303 int32
	_ = v5303
	var v5307 int32
	_ = v5307
	var v5310 int32
	_ = v5310
	var v5311 int32
	_ = v5311
	var v5312 int32
	_ = v5312
	var v5314 int32
	_ = v5314
	var v5319 int32
	_ = v5319
	var v5320 int32
	_ = v5320
	var v5322 int32
	_ = v5322
	var v5326 int32
	_ = v5326
	var v5329 int64
	_ = v5329
	var v5330 int32
	_ = v5330
	var v5334 int32
	_ = v5334
	var v5336 int32
	_ = v5336
	var v5345 int32
	_ = v5345
	var v5346 int32
	_ = v5346
	var v5349 int32
	_ = v5349
	var v5350 int32
	_ = v5350
	var v5354 int32
	_ = v5354
	var v5360 int32
	_ = v5360
	var v5362 int32
	_ = v5362
	var v5368 int32
	_ = v5368
	var v5369 int32
	_ = v5369
	var v5373 int32
	_ = v5373
	var v5379 int32
	_ = v5379
	var v5381 int32
	_ = v5381
	var v5387 int32
	_ = v5387
	var v5389 int32
	_ = v5389
	var v5391 int32
	_ = v5391
	var v5392 int32
	_ = v5392
	var v5399 int32
	_ = v5399
	var v5400 int32
	_ = v5400
	var v5401 int32
	_ = v5401
	var v5403 int32
	_ = v5403
	var v5405 int32
	_ = v5405
	var v5406 int32
	_ = v5406
	var v5409 int32
	_ = v5409
	var v5413 int32
	_ = v5413
	var v5414 int32
	_ = v5414
	var v5421 int32
	_ = v5421
	var v5427 int32
	_ = v5427
	var v5429 int32
	_ = v5429
	var v5440 int32
	_ = v5440
	var v5448 int32
	_ = v5448
	var v5454 int32
	_ = v5454
	var v5456 int32
	_ = v5456
	var v5463 int32
	_ = v5463
	var v5465 int32
	_ = v5465
	var v5473 int32
	_ = v5473
	var v5475 int32
	_ = v5475
	var v5476 int32
	_ = v5476
	var v5480 int32
	_ = v5480
	var v5481 int32
	_ = v5481
	var v5483 int32
	_ = v5483
	var v5487 int32
	_ = v5487
	var v5488 int32
	_ = v5488
	var v5489 int32
	_ = v5489
	var v5490 int32
	_ = v5490
	var v5492 int32
	_ = v5492
	var v5495 int32
	_ = v5495
	var v5496 int32
	_ = v5496
	var v5497 int32
	_ = v5497
	var v5501 int32
	_ = v5501
	var v5504 int32
	_ = v5504
	var v5505 int32
	_ = v5505
	var v5506 int32
	_ = v5506
	var v5508 int32
	_ = v5508
	var v5512 int32
	_ = v5512
	var v5516 int32
	_ = v5516
	var v5517 int32
	_ = v5517
	var v5519 int32
	_ = v5519
	var v5520 int32
	_ = v5520
	var v5526 int32
	_ = v5526
	var v5531 int32
	_ = v5531
	var v5532 int32
	_ = v5532
	var v5534 int32
	_ = v5534
	var v5537 int64
	_ = v5537
	var v5538 int32
	_ = v5538
	var v5539 int64
	_ = v5539
	var v5542 int32
	_ = v5542
	var v5546 int32
	_ = v5546
	var v5553 int32
	_ = v5553
	var v5555 int32
	_ = v5555
	var v5565 int32
	_ = v5565
	var v5567 int32
	_ = v5567
	var v5568 int32
	_ = v5568
	var v5572 int32
	_ = v5572
	var v5573 int32
	_ = v5573
	var v5576 int32
	_ = v5576
	var v5577 int32
	_ = v5577
	var v5579 int32
	_ = v5579
	var v5581 int32
	_ = v5581
	var v5584 int32
	_ = v5584
	var v5585 int32
	_ = v5585
	var v5589 int32
	_ = v5589
	var v5593 int32
	_ = v5593
	var v5594 int32
	_ = v5594
	var v5595 int32
	_ = v5595
	var v5596 int32
	_ = v5596
	var v5605 int32
	_ = v5605
	var v5653 int32
	_ = v5653
	var v5655 int32
	_ = v5655
	var v5657 int32
	_ = v5657
	var v5659 int32
	_ = v5659
	var v5662 int32
	_ = v5662
	var v5663 int32
	_ = v5663
	var v5666 int32
	_ = v5666
	var v5667 int32
	_ = v5667
	var v5676 int32
	_ = v5676
	var v5722 int32
	_ = v5722
	var v5723 int32
	_ = v5723
	var v5724 int32
	_ = v5724
	var v5725 int32
	_ = v5725
	var v5726 int32
	_ = v5726
	var v5727 int32
	_ = v5727
	var v5729 int32
	_ = v5729
	var v5734 int32
	_ = v5734
	var v5735 int32
	_ = v5735
	var v5736 int32
	_ = v5736
	var v5754 int32
	_ = v5754
	var v5755 int32
	_ = v5755
	var v5805 int32
	_ = v5805
	var v5806 int32
	_ = v5806
	var v5810 int32
	_ = v5810
	var v5811 int32
	_ = v5811
	var v5815 int32
	_ = v5815
	var v5830 int32
	_ = v5830
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
	var v5877 int32
	_ = v5877
	var v5882 int32
	_ = v5882
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
	var v5894 int32
	_ = v5894
	var v5897 int32
	_ = v5897
	var v5899 int32
	_ = v5899
	var v5900 int32
	_ = v5900
	var v5902 int32
	_ = v5902
	var v5904 int32
	_ = v5904
	var v5905 int32
	_ = v5905
	var v5910 int32
	_ = v5910
	var v5914 int32
	_ = v5914
	var v5915 int32
	_ = v5915
	var v5917 int32
	_ = v5917
	var v5918 int32
	_ = v5918
	var v5920 int32
	_ = v5920
	var v5925 int32
	_ = v5925
	var v5926 int32
	_ = v5926
	var v5928 int32
	_ = v5928
	var v5929 int32
	_ = v5929
	var v5934 int32
	_ = v5934
	var v5936 int32
	_ = v5936
	var v5937 int32
	_ = v5937
	var v5943 int32
	_ = v5943
	var v5949 int32
	_ = v5949
	var v5951 int32
	_ = v5951
	var v5957 int32
	_ = v5957
	var v5958 int32
	_ = v5958
	var v5959 int32
	_ = v5959
	var v5963 int32
	_ = v5963
	var v5969 int32
	_ = v5969
	var v5971 int32
	_ = v5971
	var v5977 int32
	_ = v5977
	var v5978 int32
	_ = v5978
	var v5980 int32
	_ = v5980
	var v5981 int32
	_ = v5981
	var v5984 int32
	_ = v5984
	var v5985 int32
	_ = v5985
	var v5986 int32
	_ = v5986
	var v5987 int32
	_ = v5987
	var v5989 int32
	_ = v5989
	var v5990 int32
	_ = v5990
	var v5994 int32
	_ = v5994
	var v6007 int32
	_ = v6007
	var v6008 int32
	_ = v6008
	var v6055 int32
	_ = v6055
	var v6060 int32
	_ = v6060
	var v6118 int32
	_ = v6118
	var v6124 int32
	_ = v6124
	var v6129 int32
	_ = v6129
	var v6133 int32
	_ = v6133
	var v6135 int32
	_ = v6135
	var v6136 int32
	_ = v6136
	var v6139 int32
	_ = v6139
	var v6140 int32
	_ = v6140
	var v6144 int32
	_ = v6144
	var v6157 int32
	_ = v6157
	var v6158 int32
	_ = v6158
	var v6205 int32
	_ = v6205
	var v6210 int32
	_ = v6210
	var v6268 int32
	_ = v6268
	var v6274 int32
	_ = v6274
	var v6279 int32
	_ = v6279
	var v6284 int32
	_ = v6284
	var v6285 int32
	_ = v6285
	var v6286 int32
	_ = v6286
	var v6290 int32
	_ = v6290
	var v6293 int32
	_ = v6293
	var v6294 int32
	_ = v6294
	var v6295 int32
	_ = v6295
	var v6297 int32
	_ = v6297
	var v6302 int32
	_ = v6302
	var v6303 int32
	_ = v6303
	var v6305 int32
	_ = v6305
	var v6306 int32
	_ = v6306
	var v6308 int32
	_ = v6308
	var v6312 int32
	_ = v6312
	var v6316 int32
	_ = v6316
	var v6317 int32
	_ = v6317
	var v6319 int32
	_ = v6319
	var v6322 int64
	_ = v6322
	var v6323 int32
	_ = v6323
	var v6324 int64
	_ = v6324
	var v6334 int32
	_ = v6334
	var v6340 int32
	_ = v6340
	var v6342 int32
	_ = v6342
	var v6348 int32
	_ = v6348
	var v6351 int32
	_ = v6351
	var v6353 int32
	_ = v6353
	var v6359 int64
	_ = v6359
	var v6360 int32
	_ = v6360
	var v6364 int32
	_ = v6364
	var v6366 int32
	_ = v6366
	var v6370 int32
	_ = v6370
	var v6372 int32
	_ = v6372
	var v6383 int32
	_ = v6383
	var v6385 int32
	_ = v6385
	var v6390 int32
	_ = v6390
	var v6392 int32
	_ = v6392
	var v6398 int32
	_ = v6398
	var v6399 int32
	_ = v6399
	var v6405 int32
	_ = v6405
	var v6410 int32
	_ = v6410
	var v6414 int32
	_ = v6414
	var v6418 int32
	_ = v6418
	var v6423 int32
	_ = v6423
	var v6427 int32
	_ = v6427
	var v6431 int32
	_ = v6431
	var v6436 int32
	_ = v6436
	var v6440 int32
	_ = v6440
	var v6441 int32
	_ = v6441
	var v6447 int32
	_ = v6447
	var v6452 int32
	_ = v6452
	var v6456 int32
	_ = v6456
	var v6460 int32
	_ = v6460
	var v6465 int32
	_ = v6465
	var v6469 int32
	_ = v6469
	var v6473 int32
	_ = v6473
	var v6478 int32
	_ = v6478
	var v6482 int32
	_ = v6482
	var v6483 int32
	_ = v6483
	var v6491 int32
	_ = v6491
	var v6496 int32
	_ = v6496
	var v6500 int32
	_ = v6500
	var v6501 int32
	_ = v6501
	var v6507 int32
	_ = v6507
	var v6512 int32
	_ = v6512
	var v6516 int32
	_ = v6516
	var v6517 int32
	_ = v6517
	var v6523 int32
	_ = v6523
	var v6528 int32
	_ = v6528
	var v6532 int32
	_ = v6532
	var v6536 int32
	_ = v6536
	var v6541 int32
	_ = v6541
	var v6545 int32
	_ = v6545
	var v6546 int32
	_ = v6546
	var v6552 int32
	_ = v6552
	var v6557 int32
	_ = v6557
	var v6558 int32
	_ = v6558
	var v6559 int32
	_ = v6559
	var v6561 int32
	_ = v6561
	var v6564 int32
	_ = v6564
	var v6576 int32
	_ = v6576
	var v6620 int32
	_ = v6620
	var v6625 int32
	_ = v6625
	var v6628 int32
	_ = v6628
	var v6629 int32
	_ = v6629
	var v6638 int32
	_ = v6638
	var v6642 int32
	_ = v6642
	var v6647 int32
	_ = v6647
	var v6704 int32
	_ = v6704
	var v6710 int32
	_ = v6710
	var v6715 int32
	_ = v6715
	var v6776 int32
	_ = v6776
	var v6779 int32
	_ = v6779
	var v6827 int32
	_ = v6827
	var v6829 int32
	_ = v6829
	var v6836 int32
	_ = v6836
	var v6839 int32
	_ = v6839
	var v6846 int32
	_ = v6846
	var v6883 int32
	_ = v6883
	var v6890 int32
	_ = v6890
	var v6892 int32
	_ = v6892
	var v6894 int32
	_ = v6894
	var v6898 int32
	_ = v6898
	var v6905 int32
	_ = v6905
	var v7013 int32
	_ = v7013
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
	v103 = int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v103 < v104 {
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
	if v94 != int32(_a_F_spgdoinsert_0) {
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
	v113 = v103
	v114 = v104
	goto L16
L14:
	;
	goto L15
L15:
	;
	v253 = F_SpGistGetLeafTupleSize(m, v58, v56+int32(432), l4)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L3
	} else {
		goto L34
	}
L16:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v113))))
	if v161 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L15
L18:
	;
	v196 = v113 + int32(1)
	if v196 < v192 {
		v113 = v196
		v114 = v192
		goto L16
	} else {
		goto L26
	}
L19:
	;
	v165 = v113 << (uint(int32(2)) % 32)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l3+v165)))
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58+v113<<(uint(int32(4))%32))+24)))
	if v171 == int32(_a_F_spgdoinsert_0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56+int32(432)+v113<<(uint(int32(2))%32)))) = int32(0)
	v192 = v114
	goto L18
L22:
	;
	v177 = F_pg_detoast_datum(m, v167)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L3
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56+int32(432)+v165))) = v167
	v192 = v114
	goto L18
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56+int32(432)+v165))) = v177
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v192 = v180
	goto L18
L26:
	;
	goto L17
L27:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_1), int32(70), int32(_a_F_spgdoinsert_2))
	mBase = m.M
	v7013 = m.ExcPending
	if v7013 != 0 {
		goto L3
	} else {
		goto L1013
	}
L28:
	;
	m.G0 = v56 + int32(800)
	return v6905
L29:
	;
	v6883 = int32(0)
	if base.B2i32(v6846 == v6883)|base.B2i32(v6839 == v6846) == v6883 {
		goto L1006
	} else {
		goto L1007
	}
L30:
	;
	if v6779 == int32(0) {
		goto L1001
	} else {
		goto L1002
	}
L31:
	;
	v6776 = int32(1)
	v6779 = v401
	goto L30
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6704 = m.ExcPending
	if v6704 != 0 {
		goto L3
	} else {
		goto L998
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6625 = m.ExcPending
	if v6625 != 0 {
		goto L3
	} else {
		goto L993
	}
L34:
	;
	v256 = v253 + int32(4)
	if base.Ui32(int32(_a_F_spgdoinsert_3)) < base.Ui32(v256) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v60 != 0 {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+428)) = int32(-1)
	v264 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+424)) = uint16(v264)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+420)) = v264
	*(*int64)(unsafe.Add(mBase, uint32(v56)+412)) = int64(4294967295)
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v271 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v259 == int32(0) {
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
	v273 = m.ExcPending
	if v273 != 0 {
		goto L3
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v274 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+411)) = uint8(v274)
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v277 == v274 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	if v60 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v6576 = v6
	goto L46
L46:
	;
	v6620 = int32(0)
	v6836 = v6620
	v6839 = v6620
	v6846 = v6576
	goto L29
L47:
	;
	v282 = int32(2)
	goto L49
L48:
	;
	v282 = int32(1)
	goto L49
L49:
	;
	if v60 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v285 = int32(8)
	goto L52
L51:
	;
	v285 = int32(0)
	goto L52
L52:
	;
	if v60 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v288 = int32(12)
	goto L55
L54:
	;
	v288 = int32(4)
	goto L55
L55:
	;
	if v60 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v291 = int32(4)
	goto L58
L57:
	;
	v291 = int32(0)
	goto L58
L58:
	;
	if v60 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v294 = int32(7)
	goto L61
L60:
	;
	v294 = int32(3)
	goto L61
L61:
	;
	v296 = v56 + int32(628)
	v301 = int32(-1)
	v310 = v282
	v311 = int32(1)
	v312 = v256
	v313 = v6
	v314 = v6
	v315 = v6
	v319 = v6
	v322 = v301
	v330 = v301
	v334 = v256
	v345 = v6
	goto L62
L62:
	;
	if v310 == int32(-1) {
		goto L70
	} else {
		goto L71
	}
L63:
	;
	v6576 = v4423
	goto L46
L64:
	;
	v4413 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v4413 != 0 {
		v6776 = int32(0)
		v6779 = v4367
		goto L30
	} else {
		goto L640
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+788)) = v1772
	if v60 != 0 {
		v1831 = int32(0)
		goto L300
	} else {
		goto L301
	}
L66:
	;
	v1749 = int32(0)
	v1763 = v1749
	v1772 = v1749
	v1773 = v1749
	goto L65
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L3
	} else {
		goto L297
	}
L68:
	;
	F_ReleaseBuffer(m, v394)
	mBase = m.M
	v1730 = m.ExcPending
	if v1730 != 0 {
		goto L3
	} else {
		goto L295
	}
L69:
	;
	if v401 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L70:
	;
	v359 = int32(_a_F_spgdoinsert_3)
	if base.Ui32(v359) <= base.Ui32(v334) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	if v313 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L73:
	;
	v362 = v359
	goto L75
L74:
	;
	v362 = v334
	goto L75
L75:
	;
	v365 = F_SpGistGetBuffer(m, l0, v294, v362, v56+int32(411))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	if v365 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v401 = v365
	v402 = v385
	goto L69
L78:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[1]))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v370+(v365^int32(-1))<<(uint(int32(6))%32))+16))
	v385 = v376
	goto L77
L79:
	;
	goto L80
L80:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[2]))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v378+v365<<(uint(int32(6))%32)+int32(-64))+16))
	v385 = v384
	goto L77
L81:
	;
	v401 = v400
	v402 = v310
	goto L69
L82:
	;
	v388 = F_ReadBuffer(m, l0, v310)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L3
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v310 == v322 {
		v401 = v313
		v402 = v322
		goto L69
	} else {
		goto L87
	}
L85:
	;
	F_LockBuffer(m, v388, int32(2))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L3
	} else {
		goto L86
	}
L86:
	;
	v400 = v388
	goto L81
L87:
	;
	v394 = F_ReadBuffer(m, l0, v310)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	v396 = F_ConditionalLockBuffer(m, v394)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L3
	} else {
		goto L89
	}
L89:
	;
	if v396 == int32(0) {
		goto L68
	} else {
		goto L90
	}
L90:
	;
	v400 = v394
	goto L81
L91:
	;
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+16)))
	v423 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v421+v420))))
	v426 = int32(0)
	if base.B2i32(v423&int32(8) == v426)^v60 == v426 {
		goto L67
	} else {
		goto L95
	}
L92:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v406+(v401^int32(-1))<<(uint(int32(2))%32))))
	v420 = v412
	goto L91
L93:
	;
	goto L94
L94:
	;
	v414 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v420 = v414 + v401<<(uint(int32(13))%32) + int32(-8192)
	goto L91
L95:
	;
	if v423&int32(4) == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v4367 = v401
	v4368 = v420
	v4373 = v311
	v4376 = v402
	goto L64
L97:
	;
	goto L98
L98:
	;
	v437 = F_spgFormLeafTuple(m, l1, l2, v56+int32(432), l4)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v444 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+14)))
	v445 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+12)))
	v446 = v444 - v445
	v447 = int32(0)
	if v447 < v446 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+16)))
	v455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420+v453)+4)))
	if v455 != 0 {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	v450 = v446
	goto L103
L102:
	;
	v450 = v447
	goto L103
L103:
	;
	goto L100
L104:
	;
	v456 = int32(20)
	goto L106
L105:
	;
	v456 = int32(0)
	goto L106
L106:
	;
	if base.Ui32(int32(base.Ui32(v439)>>(uint(int32(2))%32))+int32(4)) <= base.Ui32(v450+v456) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v459 = int32(_a_F_spgdoinsert_4)
	v461 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	v462 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v461 + v462
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+411)))
	*(*int64)(unsafe.Add(mBase, uint32(v56)+626)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+625)) = uint8(v60)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+624)) = uint8(v465)
	if base.Ui32(v462) < base.Ui32(v402-v462) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	goto L109
L109:
	;
	v654 = v402 - int32(1)
	v656 = base.B2i32(base.Ui32(v654) < base.Ui32(int32(2)))
	if base.Ui32(v654) < base.Ui32(int32(2)) {
		goto L156
	} else {
		goto L157
	}
L110:
	;
	F_MarkBufferDirty(m, v401)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L3
	} else {
		goto L134
	}
L111:
	;
	v477 = v311 & int32(_a_F_spgdoinsert_0)
	goto L113
L112:
	;
	v477 = int32(0)
	goto L113
L113:
	;
	if v477 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v480 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v437)+4)))
	v482 = v480 & int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v437)+4)) = uint16(v482)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v488 = F_SpGistPageAddNewItem(m, v420, v437, int32(base.Ui32(v484)>>(uint(int32(2))%32)), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L3
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v500 = v311 & int32(_a_F_spgdoinsert_0)
	v505 = v420 + v500<<(uint(int32(2))%32) + int32(20)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v505)))
	v509 = v420 + v506&int32(_a_F_spgdoinsert_6)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	switch v510 & int32(3) {
	case 0:
		goto L121
	default:
		goto L122
	case 2:
		goto L123
	}
L117:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+626)) = uint16(v488)
	if v313 == int32(0) {
		goto L110
	} else {
		goto L118
	}
L118:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+632)) = uint16(v330)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+630)) = uint16(v319)
	F_saveNodeLink(m, v56+int32(412), v402, v488)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L3
	} else {
		goto L119
	}
L119:
	;
	goto L110
L120:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+628)) = uint16(v311)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+626)) = uint16(v588)
	goto L110
L121:
	;
	v562 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v509)+4)))
	v565 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v437)+4)))
	v568 = v562&int32(_a_F_spgdoinsert_7) | v565&int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v437)+4)) = uint16(v568)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v574 = F_SpGistPageAddNewItem(m, v420, v437, int32(base.Ui32(v570)>>(uint(int32(2))%32)), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L3
	} else {
		goto L133
	}
L122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L3
	} else {
		goto L130
	}
L123:
	;
	v513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v437)+4)))
	v515 = v513 & int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v437)+4)) = uint16(v515)
	F_PageIndexTupleDelete(m, v420, v500)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L3
	} else {
		goto L124
	}
L124:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v523 = F_PageAddItemExtended(m, v420, v437, int32(base.Ui32(v519)>>(uint(int32(2))%32)), v500, int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	if v523 == v500 {
		v588 = v311
		goto L120
	} else {
		goto L126
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+320)) = int32(base.Ui32(v530) >> (uint(int32(2)) % 32))
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_8), v56+int32(320))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(280), int32(_a_F_spgdoinsert_10))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
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
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+304)) = v548 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_11), v56+int32(304))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(287), int32(_a_F_spgdoinsert_10))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
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
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v505)))
	v579 = v420 + v576&int32(_a_F_spgdoinsert_6)
	v580 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v579)+4)))
	v585 = v580&int32(_a_F_spgdoinsert_5) | v574&int32(_a_F_spgdoinsert_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v579)+4)) = uint16(v585)
	v588 = v574
	goto L120
L134:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597)+118)))
	if v598 != int32(112) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v647 = int32(_a_F_spgdoinsert_4)
	v649 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v649 - int32(1)
	goto L31
L136:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[6]))
	if v602 <= int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v605 != 0 {
		goto L135
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v607 != 0 {
		goto L135
	} else {
		goto L142
	}
L140:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v606 != 0 {
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
	v609 = m.ExcPending
	if v609 != 0 {
		goto L3
	} else {
		goto L143
	}
L143:
	;
	F_XLogRegisterData(m, v56+int32(624), int32(10))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L3
	} else {
		goto L144
	}
L144:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	F_XLogRegisterData(m, v437, int32(base.Ui32(v615)>>(uint(int32(2))%32)))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L3
	} else {
		goto L145
	}
L145:
	;
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+624)))
	if v623 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v624 = int32(14)
	goto L148
L147:
	;
	v624 = int32(8)
	goto L148
L148:
	;
	F_XLogRegisterBuffer(m, int32(0), v401, v624)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L3
	} else {
		goto L149
	}
L149:
	;
	v627 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+630)))
	if v627 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	F_XLogRegisterBuffer(m, int32(1), v313, int32(8))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L3
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v632 = int32(16)
	v634 = F_XLogInsert(m, v632, v632)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L3
	} else {
		goto L154
	}
L153:
	;
	goto L152
L154:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v420))) = base.I64_rotr(v634, int64(32))
	v639 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+630)))
	if v639 == int32(0) {
		goto L135
	} else {
		goto L155
	}
L155:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v314)+4)) = uint32(v634)
	v644 = int64(base.Ui64(v634) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v314))) = uint32(v644)
	goto L135
L156:
	;
	v1403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+411)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+796)) = v345
	v1405 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+12)))
	v1411 = int32(base.Ui32(v1405+int32(_a_F_spgdoinsert_12))>>(uint(int32(2))%32)) & int32(_a_F_spgdoinsert_0)
	if base.Ui32(int32(25)) <= base.Ui32(v1405) {
		goto L245
	} else {
		goto L246
	}
L157:
	;
	if v311&int32(_a_F_spgdoinsert_0) == int32(0) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	if base.B2i32(v817 == int32(0))|base.B2i32(base.Ui32(int32(4079)) < base.Ui32(v772)) != 0 {
		goto L156
	} else {
		goto L171
	}
L159:
	;
	v772 = int32(0)
	v817 = int32(1)
	goto L158
L160:
	;
	goto L161
L161:
	;
	v665 = int32(0)
	v673 = v311
	v675 = v665
	v680 = v665
	goto L162
L162:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v420+int32(20)+v673&int32(_a_F_spgdoinsert_0)<<(uint(int32(2))%32))))
	v728 = v420 + v725&int32(_a_F_spgdoinsert_6)
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v728)))
	switch v729 & int32(3) {
	case 0:
		goto L165
	default:
		goto L166
	case 2:
		v757 = v675
		v758 = v680
		goto L164
	}
L163:
	;
	v772 = v757
	v817 = base.B2i32(v758 < int32(64))
	goto L158
L164:
	;
	v759 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v728)+4)))
	v761 = v759 & int32(_a_F_spgdoinsert_7)
	if v761 != 0 {
		v673 = v761
		v675 = v757
		v680 = v758
		goto L162
	} else {
		goto L170
	}
L165:
	;
	v757 = v675 + int32(base.Ui32(v729)>>(uint(int32(2))%32)) + int32(4)
	v758 = v680 + int32(1)
	goto L164
L166:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L3
	} else {
		goto L167
	}
L167:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v728)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+288)) = v736 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_11), v56+int32(288))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L3
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(369), int32(_a_F_spgdoinsert_13))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
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
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	if base.Ui32(int32(_a_F_spgdoinsert_3)) < base.Ui32(v772+int32(base.Ui32(v823)>>(uint(int32(2))%32))+int32(4)) {
		goto L156
	} else {
		goto L172
	}
L172:
	;
	v831 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+592)) = uint16(v831)
	v834 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v834) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v844 = int32(base.Ui32(v834+int32(_a_F_spgdoinsert_12))>>(uint(int32(1))%32)) & int32(_a_F_spgdoinsert_14)
	goto L175
L174:
	;
	v844 = v831
	goto L175
L175:
	;
	v845 = F_palloc(m, v844)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L3
	} else {
		goto L176
	}
L176:
	;
	v849 = F_palloc(m, v844+int32(2))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L3
	} else {
		goto L177
	}
L177:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v855 = int32(base.Ui32(v851)>>(uint(int32(2))%32)) + int32(4)
	if v311&int32(_a_F_spgdoinsert_0) == int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v1026 = F_SpGistGetBuffer(m, l0, v294, v977, v56+int32(624)|int32(2))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L3
	} else {
		goto L193
	}
L179:
	;
	v977 = v855
	v980 = int32(0)
	v1002 = v831
	goto L178
L180:
	;
	goto L181
L181:
	;
	v871 = v311
	v872 = v855
	v875 = int32(0)
	v897 = v831
	goto L182
L182:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v420+int32(20)+v871&int32(_a_F_spgdoinsert_0)<<(uint(int32(2))%32))))
	v925 = v420 + v922&int32(_a_F_spgdoinsert_6)
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v925)))
	switch v926 & int32(3) {
	case 0:
		goto L185
	default:
		goto L186
	case 2:
		goto L187
	}
L183:
	;
	v977 = v962
	v980 = v965
	v1002 = v963
	goto L178
L184:
	;
	v965 = v875 + int32(1)
	v966 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v925)+4)))
	v968 = v966 & int32(_a_F_spgdoinsert_7)
	if v968 != 0 {
		v871 = v968
		v872 = v962
		v875 = v965
		v897 = v963
		goto L182
	} else {
		goto L191
	}
L185:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v845+v875<<(uint(int32(1))%32)))) = uint16(v871)
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v925)))
	v962 = v872 + int32(base.Ui32(v956)>>(uint(int32(2))%32)) + int32(4)
	v963 = v897
	goto L184
L186:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L3
	} else {
		goto L188
	}
L187:
	;
	v929 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v845+v875<<(uint(v929)%32)))) = uint16(v871)
	v962 = v872
	v963 = v929
	goto L184
L188:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v925)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+272)) = v938 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_11), v56+int32(272))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L3
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(446), int32(_a_F_spgdoinsert_15))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
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
	if v1026 < int32(0) {
		goto L198
	} else {
		goto L199
	}
L193:
	;
	if v1026 < int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1031+(v1026^int32(-1))<<(uint(int32(2))%32))))
	v1045 = v1037
	goto L192
L195:
	;
	goto L196
L196:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v1045 = v1039 + v1026<<(uint(int32(13))%32) + int32(-8192)
	goto L192
L197:
	;
	v1065 = F_palloc(m, v977)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L3
	} else {
		goto L201
	}
L198:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[1]))
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1049+(v1026^int32(-1))<<(uint(int32(6))%32))+16))
	v1064 = v1055
	goto L197
L199:
	;
	goto L200
L200:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[2]))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1057+v1026<<(uint(int32(6))%32)+int32(-64))+16))
	v1064 = v1063
	goto L197
L201:
	;
	v1067 = int32(_a_F_spgdoinsert_4)
	v1069 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v1069 + int32(1)
	v1073 = int32(0)
	if base.B2i32(v980 <= v1073)|v1002 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v1181 = v1073
	v1182 = v1065
	v1227 = int32(0)
	goto L204
L203:
	;
	v1088 = v1073
	v1089 = v1065
	v1091 = int32(0)
	goto L205
L204:
	;
	v1228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v437)+4)))
	v1231 = v1227 | v1228&int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v437)+4)) = uint16(v1231)
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v1241 = F_SpGistPageAddNewItem(m, v1045, v437, int32(base.Ui32(v1236)>>(uint(int32(2))%32)), v56+int32(592))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L3
	} else {
		goto L212
	}
L205:
	;
	v1135 = v1088 << (uint(int32(1)) % 32)
	v1137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v845+v1135))))
	v1138 = int32(2)
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v420+int32(20)+v1137<<(uint(v1138)%32))))
	v1144 = v420 + v1141&int32(_a_F_spgdoinsert_6)
	v1145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1144)+4)))
	v1150 = v1145&int32(_a_F_spgdoinsert_5) | v1091&int32(_a_F_spgdoinsert_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v1144)+4)) = uint16(v1150)
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1144)))
	v1158 = F_SpGistPageAddNewItem(m, v1045, v1144, int32(base.Ui32(v1153)>>(uint(v1138)%32)), v56+int32(592))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L3
	} else {
		goto L207
	}
L206:
	;
	v1181 = v980
	v1182 = v1168
	v1227 = v1158 & int32(_a_F_spgdoinsert_7)
	goto L204
L207:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v849+v1135))) = uint16(v1158)
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1144)))
	v1163 = int32(base.Ui32(v1161) >> (uint(int32(2)) % 32))
	if v1163 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	base.MemoryCopy(m, v1089, v1144, v1163)
	goto L210
L209:
	;
	goto L210
L210:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1144)))
	v1168 = v1089 + int32(base.Ui32(v1165)>>(uint(int32(2))%32))
	v1170 = v1088 + int32(1)
	if v1170 != v980 {
		v1088 = v1170
		v1089 = v1168
		v1091 = v1158
		goto L205
	} else {
		goto L211
	}
L211:
	;
	goto L206
L212:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v849+v1181<<(uint(int32(1))%32)))) = uint16(v1241)
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v1246 = int32(base.Ui32(v1244) >> (uint(int32(2)) % 32))
	if v1246 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	base.MemoryCopy(m, v1182, v437, v1246)
	goto L215
L214:
	;
	goto L215
L215:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v1251 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1252 = int32(3)
	goto L218
L217:
	;
	v1252 = int32(1)
	goto L218
L218:
	;
	F_spgPageIndexMultiDelete(m, l1, v420, v845, v980, v1252, int32(3), v1064, v1241)
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L3
	} else {
		goto L219
	}
L219:
	;
	F_saveNodeLink(m, v56+int32(412), v1064, v1241)
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L3
	} else {
		goto L220
	}
L220:
	;
	F_MarkBufferDirty(m, v401)
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L3
	} else {
		goto L221
	}
L221:
	;
	F_MarkBufferDirty(m, v1026)
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L3
	} else {
		goto L222
	}
L222:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1264)+118)))
	if v1265 != int32(112) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1340 = int32(_a_F_spgdoinsert_4)
	v1342 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v1342 - int32(1)
	F_SpGistSetLastUsedPage(m, l0, v1026)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L3
	} else {
		goto L243
	}
L224:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[6]))
	if v1269 <= int32(0) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1272 != 0 {
		goto L223
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v1274 != 0 {
		goto L223
	} else {
		goto L230
	}
L228:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1273 != 0 {
		goto L223
	} else {
		goto L229
	}
L229:
	;
	goto L227
L230:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v1276 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+640)) = uint8(v1276)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+636)) = v1275
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+632)) = uint16(v330)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+630)) = uint16(v319)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+628)) = uint8(v60)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+627)) = uint8(v1002)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+624)) = uint16(v980)
	F_XLogBeginInsert(m)
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L3
	} else {
		goto L231
	}
L231:
	;
	F_XLogRegisterData(m, v56+int32(624), int32(20))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L3
	} else {
		goto L232
	}
L232:
	;
	F_XLogRegisterData(m, v845, v980<<(uint(int32(1))%32))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L3
	} else {
		goto L233
	}
L233:
	;
	F_XLogRegisterData(m, v849, v1181<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L3
	} else {
		goto L234
	}
L234:
	;
	F_XLogRegisterData(m, v1065, v1182+int32(base.Ui32(v1248)>>(uint(int32(2))%32))-v1065)
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L3
	} else {
		goto L235
	}
L235:
	;
	F_XLogRegisterBuffer(m, int32(0), v401, int32(8))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L3
	} else {
		goto L236
	}
L236:
	;
	v1314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+626)))
	if v1314 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1315 = int32(14)
	goto L239
L238:
	;
	v1315 = int32(8)
	goto L239
L239:
	;
	F_XLogRegisterBuffer(m, int32(1), v1026, v1315)
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L3
	} else {
		goto L240
	}
L240:
	;
	F_XLogRegisterBuffer(m, int32(2), v313, int32(8))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L3
	} else {
		goto L241
	}
L241:
	;
	v1324 = F_XLogInsert(m, int32(16), int32(32))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L3
	} else {
		goto L242
	}
L242:
	;
	v1326 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v420))) = base.I64_rotr(v1324, v1326)
	v1329 = base.I32_wrap_i64(v1324)
	*(*int32)(unsafe.Add(mBase, uint32(v1045)+4)) = v1329
	v1333 = base.I32_wrap_i64(int64(base.Ui64(v1324) >> (uint(v1326) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v1045))) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v314)+4)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v314))) = v1333
	goto L223
L243:
	;
	F_UnlockReleaseBuffer(m, v1026)
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L3
	} else {
		goto L244
	}
L244:
	;
	goto L31
L245:
	;
	v1415 = v1411
	goto L247
L246:
	;
	v1415 = int32(0)
	goto L247
L247:
	;
	v1417 = v1415 + int32(1)
	v1419 = v1417 << (uint(int32(2)) % 32)
	v1420 = F_palloc(m, v1419)
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L3
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+792)) = v1420
	v1424 = v1417 << (uint(int32(1)) % 32)
	v1425 = F_palloc(m, v1424)
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L3
	} else {
		goto L249
	}
L249:
	;
	v1427 = F_palloc(m, v1424)
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L3
	} else {
		goto L250
	}
L250:
	;
	v1429 = F_palloc(m, v1419)
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L3
	} else {
		goto L251
	}
L251:
	;
	v1431 = F_palloc(m, v1419)
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L3
	} else {
		goto L252
	}
L252:
	;
	v1433 = F_palloc(m, v1417)
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L3
	} else {
		goto L253
	}
L253:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+584)) = v1435
	v1437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+588)) = uint8(v1437)
	if base.Ui32(v654) <= base.Ui32(int32(1)) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	if v1415 == int32(0) {
		goto L66
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	if v311&int32(_a_F_spgdoinsert_0) == int32(0) {
		goto L66
	} else {
		goto L275
	}
L257:
	;
	v1445 = int32(0)
	v1455 = int32(1)
	v1459 = v1445
	v1469 = v1445
	goto L260
L258:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L3
	} else {
		goto L272
	}
L259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L3
	} else {
		goto L270
	}
L260:
	;
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v420+int32(20)+v1455&int32(_a_F_spgdoinsert_0)<<(uint(int32(2))%32))))
	v1509 = v420 + v1506&int32(_a_F_spgdoinsert_6)
	v1510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1509))))
	if v1510&int32(3) != 0 {
		goto L258
	} else {
		goto L262
	}
L261:
	;
	v1763 = v1411
	v1772 = v1411
	v1773 = v1546
	goto L65
L262:
	;
	if v60 != 0 {
		v1527 = int32(0)
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1528 = int32(2)
	v1529 = v1459 << (uint(v1528) % 32)
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v56)+792))
	*(*int32)(unsafe.Add(mBase, uint32(v1529+v1530))) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v1429+v1529))) = v1509
	v1535 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1425+v1459<<(uint(v1535)%32)))) = uint16(v1455)
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1509)))
	v1546 = v1469 + int32(base.Ui32(v1541)>>(uint(v1528)%32)) + int32(4)
	v1548 = v1459 + v1535
	if v1548 != v1415 {
		v1455 = v1455 + v1535
		v1459 = v1548
		v1469 = v1546
		goto L260
	} else {
		goto L269
	}
L264:
	;
	v1515 = v1509 + int32(16)
	v1516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v1516 != int32(1) {
		v1527 = v1515
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v1519 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)))
	switch v1519 - int32(1) {
	case 0:
		goto L268
	case 1:
		goto L267
	default:
		goto L259
	case 3:
		goto L266
	}
L266:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v1515)))
	v1527 = v1524
	goto L263
L267:
	;
	v1523 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1515))))
	v1527 = v1523
	goto L263
L268:
	;
	v1522 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1515))))
	v1527 = v1522
	goto L263
L269:
	;
	goto L261
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+208)) = base.I32_extend16_s(v1519)
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_16), v56+int32(208))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L3
	} else {
		goto L271
	}
L271:
	;
	goto L27
L272:
	;
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1509)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+224)) = v1565 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_11), v56+int32(224))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L3
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(767), int32(_a_F_spgdoinsert_17))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
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
	v1585 = int32(0)
	v1595 = v311
	v1599 = v1585
	v1608 = v1585
	v1609 = v1585
	goto L276
L276:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v420+int32(20)+v1595&int32(_a_F_spgdoinsert_0)<<(uint(int32(2))%32))))
	v1649 = v420 + v1646&int32(_a_F_spgdoinsert_6)
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1649)))
	switch v1650 & int32(3) {
	case 0:
		goto L281
	default:
		goto L280
	case 2:
		goto L279
	}
L277:
	;
	v1763 = v1725
	v1772 = v1722
	v1773 = v1723
	goto L65
L278:
	;
	v1725 = v1599 + int32(1)
	v1726 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1649)+4)))
	v1728 = v1726 & int32(_a_F_spgdoinsert_7)
	if v1728 != 0 {
		v1595 = v1728
		v1599 = v1725
		v1608 = v1722
		v1609 = v1723
		goto L276
	} else {
		goto L294
	}
L279:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1425+v1599<<(uint(int32(1))%32)))) = uint16(v1595)
	v1722 = v1608
	v1723 = v1609
	goto L278
L280:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L3
	} else {
		goto L291
	}
L281:
	;
	if v60 != 0 {
		v1667 = int32(0)
		goto L283
	} else {
		goto L284
	}
L282:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L3
	} else {
		goto L289
	}
L283:
	;
	v1668 = int32(2)
	v1669 = v1608 << (uint(v1668) % 32)
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v56)+792))
	*(*int32)(unsafe.Add(mBase, uint32(v1669+v1670))) = v1667
	*(*int32)(unsafe.Add(mBase, uint32(v1429+v1669))) = v1649
	v1675 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1425+v1599<<(uint(v1675)%32)))) = uint16(v1595)
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1649)))
	v1722 = v1608 + v1675
	v1723 = v1609 + int32(base.Ui32(v1681)>>(uint(v1668)%32)) - int32(16)
	goto L278
L284:
	;
	v1655 = v1649 + int32(16)
	v1656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v1656 != int32(1) {
		v1667 = v1655
		goto L283
	} else {
		goto L285
	}
L285:
	;
	v1659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)))
	switch v1659 - int32(1) {
	case 0:
		goto L288
	case 1:
		goto L287
	default:
		goto L282
	case 3:
		goto L286
	}
L286:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1655)))
	v1667 = v1664
	goto L283
L287:
	;
	v1663 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1655))))
	v1667 = v1663
	goto L283
L288:
	;
	v1662 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1655))))
	v1667 = v1662
	goto L283
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+256)) = base.I32_extend16_s(v1659)
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_16), v56+int32(256))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L3
	} else {
		goto L290
	}
L290:
	;
	goto L27
L291:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1649)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+240)) = v1702 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_11), v56+int32(240))
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L3
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(803), int32(_a_F_spgdoinsert_17))
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L3
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L294:
	;
	goto L277
L295:
	;
	F_UnlockReleaseBuffer(m, v313)
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L3
	} else {
		goto L296
	}
L296:
	;
	v6905 = int32(0)
	goto L28
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+336)) = v402
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_18), v56+int32(336))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L3
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(2105), int32(_a_F_spgdoinsert_19))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L3
	} else {
		goto L299
	}
L299:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L300:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v56)+792))
	v1833 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1832+v1772<<(uint(v1833)%32)))) = v1831
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	*(*int32)(unsafe.Add(mBase, uint32(v1429+v1837<<(uint(v1833)%32)))) = v437
	v1842 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v56)+760)) = v1842
	v1845 = v1837 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+788)) = v1845
	*(*int64)(unsafe.Add(mBase, uint32(v56)+768)) = v1842
	*(*int64)(unsafe.Add(mBase, uint32(v56)+776)) = v1842
	if v60 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L301:
	;
	v1808 = v437 + int32(16)
	v1809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v1809 != int32(1) {
		v1831 = v1808
		goto L300
	} else {
		goto L302
	}
L302:
	;
	v1812 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)))
	switch v1812 - int32(1) {
	case 0:
		goto L303
	case 1:
		goto L306
	default:
		goto L304
	case 3:
		goto L305
	}
L303:
	;
	v1828 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1808))))
	v1831 = v1828
	goto L300
L304:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L3
	} else {
		goto L307
	}
L305:
	;
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1808)))
	v1831 = v1816
	goto L300
L306:
	;
	v1815 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1808))))
	v1831 = v1815
	goto L300
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+192)) = base.I32_extend16_s(v1812)
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_16), v56+int32(192))
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L3
	} else {
		goto L308
	}
L308:
	;
	goto L27
L309:
	;
	v2534 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	v2537 = F_palloc(m, v2534<<(uint(int32(2))%32))
	mBase = m.M
	v2538 = m.ExcPending
	if v2538 != 0 {
		goto L3
	} else {
		goto L357
	}
L310:
	;
	if v2092 < int32(2) {
		v2490 = v2088
		v2491 = v2089
		v2494 = v2092
		v2513 = v2111
		goto L309
	} else {
		goto L335
	}
L311:
	;
	v1853 = int32(1)
	v1856 = F_index_getprocinfo(m, l0, v1853, int32(3))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L3
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	v1970 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+768)) = v1970
	v1975 = F_palloc0(m, v1845<<(uint(int32(2))%32))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L3
	} else {
		goto L325
	}
L314:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1858)))
	v1864 = F_FunctionCall2Coll(m, v1856, v1859, v56+int32(788), v56+int32(760))
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L3
	} else {
		goto L315
	}
L315:
	;
	v1866 = int32(0)
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	if v1869 <= v1866 {
		v2490 = v1866
		v2491 = v1866
		v2494 = v1869
		v2513 = v1853
		goto L309
	} else {
		goto L316
	}
L316:
	;
	v1879 = v1866
	v1882 = v1866
	goto L317
L317:
	;
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1925)))
	if int32(2) <= v1926 {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	v2088 = v1866
	v2089 = v1965
	v2092 = v1968
	v2111 = v1853
	goto L310
L319:
	;
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1429+v1879<<(uint(int32(2))%32))))
	F_spgDeformLeafTuple(m, v1932, v1925, v56+int32(624), v56+int32(592), int32(0))
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L3
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	v1941 = v1879 << (uint(int32(2)) % 32)
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v56)+780))
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1941+v1942)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+624)) = v1944
	v1946 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+592)) = uint8(v1946)
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v1941+v1429)))
	v1957 = F_spgFormLeafTuple(m, l1, v1950+int32(6), v56+int32(624), v56+int32(592))
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L3
	} else {
		goto L323
	}
L322:
	;
	goto L321
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1941+v1431))) = v1957
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(v1957)))
	v1965 = v1882 + int32(base.Ui32(v1960)>>(uint(int32(2))%32)) + int32(4)
	v1967 = v1879 + int32(1)
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	if v1967 < v1968 {
		v1879 = v1967
		v1882 = v1965
		goto L317
	} else {
		goto L324
	}
L324:
	;
	goto L318
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+776)) = v1975
	v1978 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(v1837) {
		v2490 = v1978
		v2491 = v1978
		v2494 = v1845
		v2513 = v1970
		goto L309
	} else {
		goto L326
	}
L326:
	;
	v1990 = v1978
	v1993 = v1978
	goto L327
L327:
	;
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v2036)))
	if int32(2) <= v2037 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	v2088 = v1978
	v2089 = v2074
	v2092 = v2077
	v2111 = v1970
	goto L310
L329:
	;
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v1429+v1990<<(uint(int32(2))%32))))
	F_spgDeformLeafTuple(m, v2043, v2036, v56+int32(624), v56+int32(592), int32(1))
	mBase = m.M
	v2050 = m.ExcPending
	if v2050 != 0 {
		goto L3
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+624)) = int32(0)
	v2053 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+592)) = uint8(v2053)
	v2056 = v1990 << (uint(int32(2)) % 32)
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v2056+v1429)))
	v2066 = F_spgFormLeafTuple(m, l1, v2059+int32(6), v56+int32(624), v56+int32(592))
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L3
	} else {
		goto L333
	}
L332:
	;
	goto L331
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1431+v2056))) = v2066
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2066)))
	v2074 = v1993 + int32(base.Ui32(v2069)>>(uint(int32(2))%32)) + int32(4)
	v2076 = v1990 + int32(1)
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	if v2076 < v2077 {
		v1990 = v2076
		v1993 = v2074
		goto L327
	} else {
		goto L334
	}
L334:
	;
	goto L328
L335:
	;
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v56)+776))
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v2134)))
	v2136 = int32(1)
	v2139 = v2092 - base.B2i32(base.Ui32(int32(_a_F_spgdoinsert_3)) < base.Ui32(v2089))
	if base.Ui32(v2136) < base.Ui32(v2139) {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v2148 = v2136
	goto L339
L337:
	;
	goto L338
L338:
	;
	if base.Ui32(int32(_a_F_spgdoinsert_20)) <= base.Ui32(v2089) {
		goto L343
	} else {
		goto L344
	}
L339:
	;
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(v2134+v2148<<(uint(int32(2))%32))))
	if v2198 != v2135 {
		v2490 = v2088
		v2491 = v2089
		v2494 = v2092
		v2513 = v2111
		goto L309
	} else {
		goto L341
	}
L340:
	;
	goto L338
L341:
	;
	v2201 = v2148 + int32(1)
	if v2201 != v2139 {
		v2148 = v2201
		goto L339
	} else {
		goto L342
	}
L342:
	;
	goto L340
L343:
	;
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(v2134+v2092<<(uint(int32(2))%32)-int32(4))))
	v2266 = base.B2i32(v2264 == v2135)
	goto L345
L344:
	;
	v2266 = int32(1)
	goto L345
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+768)) = int32(8)
	v2276 = int32(0)
	goto L346
L346:
	;
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(v56)+776))
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	v2328 = base.I32_rem_s(v2276, v2327)
	*(*int32)(unsafe.Add(mBase, uint32(v2323+v2276<<(uint(int32(2))%32)))) = v2328
	v2331 = v2276 + int32(1)
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	if v2331 < v2332 {
		v2276 = v2331
		goto L346
	} else {
		goto L348
	}
L347:
	;
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(v56)+772))
	if v2334 == int32(0) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	goto L347
L349:
	;
	v2466 = int32(4)
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	if v2266 != 0 {
		v2490 = v2466
		v2491 = v2089
		v2494 = v2467
		v2513 = v2111
		goto L309
	} else {
		goto L356
	}
L350:
	;
	v2337 = int32(2)
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v2334+v2135<<(uint(v2337)%32))))
	v2341 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	v2344 = F_palloc(m, v2341<<(uint(v2337)%32))
	mBase = m.M
	v2345 = m.ExcPending
	if v2345 != 0 {
		goto L3
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+772)) = v2344
	v2347 = int32(0)
	v2348 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v2348 <= v2347 {
		goto L349
	} else {
		goto L352
	}
L352:
	;
	v2357 = v2347
	goto L353
L353:
	;
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(v56)+772))
	*(*int32)(unsafe.Add(mBase, uint32(v2404+v2357<<(uint(int32(2))%32)))) = v2340
	v2410 = v2357 + int32(1)
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v2410 < v2411 {
		v2357 = v2410
		goto L353
	} else {
		goto L355
	}
L354:
	;
	goto L349
L355:
	;
	goto L354
L356:
	;
	v2469 = v2467 - int32(1)
	v2470 = int32(2)
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v1431+v2469<<(uint(v2470)%32))))
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v2473)))
	v2490 = v2466
	v2491 = v2089 - int32(base.Ui32(v2474)>>(uint(v2470)%32)) - int32(4)
	v2494 = v2469
	v2513 = int32(0)
	goto L309
L357:
	;
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	v2542 = F_palloc0(m, v2539<<(uint(int32(2))%32))
	mBase = m.M
	v2543 = m.ExcPending
	if v2543 != 0 {
		goto L3
	} else {
		goto L358
	}
L358:
	;
	v2544 = int32(0)
	v2545 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v2544 < v2545 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v2554 = v2544
	goto L362
L360:
	;
	v2625 = v2545
	goto L361
L361:
	;
	v2671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+760)))
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v56)+764))
	v2673 = F_spgFormInnerTuple(m, l1, v2671, v2672, v2625, v2537)
	mBase = m.M
	v2674 = m.ExcPending
	if v2674 != 0 {
		goto L3
	} else {
		goto L369
	}
L362:
	;
	v2602 = v2554 << (uint(int32(2)) % 32)
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(v56)+772))
	if v2604 != 0 {
		goto L364
	} else {
		goto L365
	}
L363:
	;
	v2625 = v2616
	goto L361
L364:
	;
	v2606 = *(*int32)(unsafe.Add(mBase, uint32(v2602+v2604)))
	v2608 = v2606
	goto L366
L365:
	;
	v2608 = int32(0)
	goto L366
L366:
	;
	v2611 = F_spgFormNodeTuple(m, l1, v2608, base.B2i32(v2604 == int32(0)))
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L3
	} else {
		goto L367
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2537+v2602))) = v2611
	v2615 = v2554 + int32(1)
	v2616 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v2615 < v2616 {
		v2554 = v2615
		goto L362
	} else {
		goto L368
	}
L368:
	;
	goto L363
L369:
	;
	v2675 = *(*int32)(unsafe.Add(mBase, uint32(v2673)))
	*(*int32)(unsafe.Add(mBase, uint32(v2673))) = v2490 | v2675&int32(-5)
	if v2675&int32(_a_F_spgdoinsert_21) != 0 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v2694 = v2673 + int32(base.Ui32(v2675)>>(uint(int32(16))%32)) + int32(8)
	v2695 = int32(0)
	goto L373
L371:
	;
	goto L372
L372:
	;
	v2810 = int32(0)
	if v2810 < v2494 {
		goto L379
	} else {
		goto L380
	}
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2537+v2695<<(uint(int32(2))%32)))) = v2694
	v2745 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2694)+6)))
	v2746 = int32(_a_F_spgdoinsert_22)
	v2750 = v2695 + int32(1)
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2673)))
	if base.Ui32(v2750) < base.Ui32(int32(base.Ui32(v2751)>>(uint(int32(3))%32))&v2746) {
		v2694 = v2694 + v2745&v2746
		v2695 = v2750
		goto L373
	} else {
		goto L375
	}
L374:
	;
	goto L372
L375:
	;
	goto L374
L376:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+575)) = uint8(v60)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+570)) = uint8(v1403)
	v3743 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+566)) = uint16(v3743)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+564)) = uint8(v656)
	v3747 = F_palloc(m, v2491)
	mBase = m.M
	v3748 = m.ExcPending
	if v3748 != 0 {
		goto L3
	} else {
		goto L484
	}
L377:
	;
	if v3426 <= int32(0) {
		v3708 = v3426
		v3710 = v3021
		v3720 = v3438
		goto L376
	} else {
		goto L473
	}
L378:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3396 = m.ExcPending
	if v3396 != 0 {
		goto L3
	} else {
		goto L470
	}
L379:
	;
	v2819 = v2810
	goto L382
L380:
	;
	goto L381
L381:
	;
	v2944 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+574)) = uint8(v2944)
	if v313 == v2944 {
		v2985 = v2944
		goto L387
	} else {
		goto L388
	}
L382:
	;
	v2867 = v2819 << (uint(int32(2)) % 32)
	v2868 = *(*int32)(unsafe.Add(mBase, uint32(v56)+776))
	v2870 = *(*int32)(unsafe.Add(mBase, uint32(v2867+v2868)))
	if v2870 < int32(0) {
		goto L378
	} else {
		goto L384
	}
L383:
	;
	goto L381
L384:
	;
	v2873 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v2873 <= v2870 {
		goto L378
	} else {
		goto L385
	}
L385:
	;
	v2875 = int32(2)
	v2877 = v2542 + v2870<<(uint(v2875)%32)
	v2878 = *(*int32)(unsafe.Add(mBase, uint32(v2877)))
	v2880 = *(*int32)(unsafe.Add(mBase, uint32(v2867+v1431)))
	v2881 = *(*int32)(unsafe.Add(mBase, uint32(v2880)))
	*(*int32)(unsafe.Add(mBase, uint32(v2877))) = v2878 + int32(base.Ui32(v2881)>>(uint(v2875)%32)) + int32(4)
	v2889 = v2819 + int32(1)
	if v2889 != v2494 {
		v2819 = v2889
		goto L382
	} else {
		goto L386
	}
L386:
	;
	goto L383
L387:
	;
	if v656 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L388:
	;
	v2950 = int32(1)
	if base.Ui32(v322-v2950) <= base.Ui32(v2950) {
		goto L390
	} else {
		goto L391
	}
L389:
	;
	v2980 = base.I32_rem_u_s(v322+int32(1), int32(3))
	v2982 = F_SpGistGetBuffer(m, l0, v2980|v291, v2975, v56+int32(574))
	mBase = m.M
	v2983 = m.ExcPending
	if v2983 != 0 {
		goto L3
	} else {
		goto L401
	}
L390:
	;
	v2954 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2673)+4)))
	v2975 = v2954 + int32(4)
	goto L389
L391:
	;
	goto L392
L392:
	;
	v2957 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v314)+14)))
	v2958 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v314)+12)))
	v2959 = v2957 - v2958
	v2960 = int32(0)
	if v2960 < v2959 {
		goto L394
	} else {
		goto L395
	}
L393:
	;
	v2966 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v314)+16)))
	v2968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v314+v2966)+4)))
	if v2968 != 0 {
		goto L397
	} else {
		goto L398
	}
L394:
	;
	v2963 = v2959
	goto L396
L395:
	;
	v2963 = v2960
	goto L396
L396:
	;
	goto L393
L397:
	;
	v2969 = int32(20)
	goto L399
L398:
	;
	v2969 = int32(0)
	goto L399
L399:
	;
	v2971 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2673)+4)))
	v2973 = v2971 + int32(4)
	if base.Ui32(v2973) <= base.Ui32(v2963+v2969) {
		v2985 = v313
		goto L387
	} else {
		goto L400
	}
L400:
	;
	v2975 = v2973
	goto L389
L401:
	;
	v2985 = v2982
	goto L387
L402:
	;
	v2988 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+14)))
	v2989 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+12)))
	v2990 = v2988 - v2989
	v2991 = int32(0)
	if v2991 < v2990 {
		goto L406
	} else {
		goto L407
	}
L403:
	;
	v2996 = v2944
	goto L404
L404:
	;
	v2997 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+571)) = uint8(v2997)
	if v2491 <= v2996 {
		goto L409
	} else {
		goto L410
	}
L405:
	;
	v2996 = v2994 + v1773
	goto L404
L406:
	;
	v2994 = v2990
	goto L408
L407:
	;
	v2994 = v2991
	goto L408
L408:
	;
	goto L405
L409:
	;
	v3000 = int32(0)
	v3001 = v1772 + v2513
	if base.B2i32(v3001 <= v3000)|base.B2i32(v3001 == v3000) != 0 {
		v3708 = v3001
		v3710 = v3000
		v3720 = v2513
		goto L376
	} else {
		goto L412
	}
L410:
	;
	goto L411
L411:
	;
	v3009 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	if v3009 != int32(1) {
		goto L413
	} else {
		goto L414
	}
L412:
	;
	base.MemoryFill(m, v1433, int32(0), v3001)
	v3708 = v3001
	v3710 = v3000
	v3720 = v2513
	goto L376
L413:
	;
	v3017 = int32(_a_F_spgdoinsert_3)
	if base.Ui32(v3017) <= base.Ui32(v2491) {
		goto L416
	} else {
		goto L417
	}
L414:
	;
	if base.Ui32(v2491) <= base.Ui32(int32(_a_F_spgdoinsert_3)) {
		goto L413
	} else {
		goto L415
	}
L415:
	;
	v3708 = v1772
	v3710 = int32(0)
	v3720 = int32(0)
	goto L376
L416:
	;
	v3020 = v3017
	goto L418
L417:
	;
	v3020 = v2491
	goto L418
L418:
	;
	v3021 = F_SpGistGetBuffer(m, l0, v294, v3020, v56+int32(571))
	mBase = m.M
	v3022 = m.ExcPending
	if v3022 != 0 {
		goto L3
	} else {
		goto L419
	}
L419:
	;
	v3023 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	v3024 = F_palloc(m, v3023)
	mBase = m.M
	v3025 = m.ExcPending
	if v3025 != 0 {
		goto L3
	} else {
		goto L420
	}
L420:
	;
	v3026 = int32(0)
	v3027 = base.B2i32(v3026 <= v3021)
	if v3027 == v3026 {
		goto L422
	} else {
		goto L423
	}
L421:
	;
	v3046 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3045)+14)))
	v3047 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3045)+12)))
	v3048 = v3046 - v3047
	v3049 = int32(0)
	if v3049 < v3048 {
		goto L426
	} else {
		goto L427
	}
L422:
	;
	v3031 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v3037 = *(*int32)(unsafe.Add(mBase, uint32(v3031+(v3021^int32(-1))<<(uint(int32(2))%32))))
	v3045 = v3037
	goto L421
L423:
	;
	goto L424
L424:
	;
	v3039 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v3045 = v3039 + v3021<<(uint(int32(13))%32) + int32(-8192)
	goto L421
L425:
	;
	v3053 = int32(0)
	v3054 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v3053 < v3054 {
		goto L429
	} else {
		goto L430
	}
L426:
	;
	v3052 = v3048
	goto L428
L427:
	;
	v3052 = v3049
	goto L428
L428:
	;
	goto L425
L429:
	;
	v3063 = v3053
	v3065 = v2996
	v3072 = v3052
	goto L432
L430:
	;
	v3138 = v2996
	v3145 = v3052
	goto L431
L431:
	;
	if int32(0) <= v3138|v3145 {
		goto L439
	} else {
		goto L440
	}
L432:
	;
	v3110 = v3063 + v3024
	v3113 = v2542 + v3063<<(uint(int32(2))%32)
	v3114 = *(*int32)(unsafe.Add(mBase, uint32(v3113)))
	if v3114 <= v3065 {
		goto L435
	} else {
		goto L436
	}
L433:
	;
	v3138 = v3124
	v3145 = v3125
	goto L431
L434:
	;
	v3127 = v3063 + int32(1)
	v3128 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v3127 < v3128 {
		v3063 = v3127
		v3065 = v3124
		v3072 = v3125
		goto L432
	} else {
		goto L438
	}
L435:
	;
	v3116 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3110))) = uint8(v3116)
	v3118 = *(*int32)(unsafe.Add(mBase, uint32(v3113)))
	v3124 = v3065 - v3118
	v3125 = v3072
	goto L434
L436:
	;
	goto L437
L437:
	;
	v3120 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3110))) = uint8(v3120)
	v3122 = *(*int32)(unsafe.Add(mBase, uint32(v3113)))
	v3124 = v3065
	v3125 = v3072 - v3122
	goto L434
L438:
	;
	goto L433
L439:
	;
	v3426 = v1772 + v2513
	v3438 = v2513
	goto L377
L440:
	;
	goto L441
L441:
	;
	if v2513 != 0 {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v3187 = *(*int32)(unsafe.Add(mBase, uint32(v56)+788))
	v3188 = int32(2)
	v3190 = int32(4)
	v3191 = v3187<<(uint(v3188)%32) - v3190
	v3192 = *(*int32)(unsafe.Add(mBase, uint32(v56)+776))
	v3194 = *(*int32)(unsafe.Add(mBase, uint32(v3191+v3192)))
	v3197 = v2542 + v3194<<(uint(v3188)%32)
	v3198 = *(*int32)(unsafe.Add(mBase, uint32(v3197)))
	v3200 = *(*int32)(unsafe.Add(mBase, uint32(v3191+v1431)))
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(v3200)))
	*(*int32)(unsafe.Add(mBase, uint32(v3197))) = v3198 - int32(base.Ui32(v3201)>>(uint(v3188)%32)) - v3190
	if v3027 == int32(0) {
		goto L446
	} else {
		goto L447
	}
L443:
	;
	goto L444
L444:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3383 = m.ExcPending
	if v3383 != 0 {
		goto L3
	} else {
		goto L467
	}
L445:
	;
	v3226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3225)+14)))
	v3227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3225)+12)))
	v3228 = v3226 - v3227
	v3229 = int32(0)
	if v3229 < v3228 {
		goto L450
	} else {
		goto L451
	}
L446:
	;
	v3211 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v3217 = *(*int32)(unsafe.Add(mBase, uint32(v3211+(v3021^int32(-1))<<(uint(int32(2))%32))))
	v3225 = v3217
	goto L445
L447:
	;
	goto L448
L448:
	;
	v3219 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v3225 = v3219 + v3021<<(uint(int32(13))%32) + int32(-8192)
	goto L445
L449:
	;
	v3233 = int32(0)
	v3234 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v3233 < v3234 {
		goto L453
	} else {
		goto L454
	}
L450:
	;
	v3232 = v3228
	goto L452
L451:
	;
	v3232 = v3229
	goto L452
L452:
	;
	goto L449
L453:
	;
	v3243 = v3233
	v3244 = v2996
	v3258 = v3232
	goto L456
L454:
	;
	v3317 = v2996
	v3331 = v3232
	goto L455
L455:
	;
	v3363 = int32(0)
	if v3363 <= v3317|v3331 {
		v3426 = v1772
		v3438 = v3363
		goto L377
	} else {
		goto L463
	}
L456:
	;
	v3290 = v3243 + v3024
	v3293 = v2542 + v3243<<(uint(int32(2))%32)
	v3294 = *(*int32)(unsafe.Add(mBase, uint32(v3293)))
	if v3294 <= v3244 {
		goto L459
	} else {
		goto L460
	}
L457:
	;
	v3317 = v3304
	v3331 = v3305
	goto L455
L458:
	;
	v3307 = v3243 + int32(1)
	v3308 = *(*int32)(unsafe.Add(mBase, uint32(v56)+768))
	if v3307 < v3308 {
		v3243 = v3307
		v3244 = v3304
		v3258 = v3305
		goto L456
	} else {
		goto L462
	}
L459:
	;
	v3296 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3290))) = uint8(v3296)
	v3298 = *(*int32)(unsafe.Add(mBase, uint32(v3293)))
	v3304 = v3244 - v3298
	v3305 = v3258
	goto L458
L460:
	;
	goto L461
L461:
	;
	v3300 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3290))) = uint8(v3300)
	v3302 = *(*int32)(unsafe.Add(mBase, uint32(v3293)))
	v3304 = v3244
	v3305 = v3258 - v3302
	goto L458
L462:
	;
	goto L457
L463:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3370 = m.ExcPending
	if v3370 != 0 {
		goto L3
	} else {
		goto L464
	}
L464:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_23), int32(0))
	mBase = m.M
	v3374 = m.ExcPending
	if v3374 != 0 {
		goto L3
	} else {
		goto L465
	}
L465:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1112), int32(_a_F_spgdoinsert_17))
	mBase = m.M
	v3379 = m.ExcPending
	if v3379 != 0 {
		goto L3
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
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_23), int32(0))
	mBase = m.M
	v3387 = m.ExcPending
	if v3387 != 0 {
		goto L3
	} else {
		goto L468
	}
L468:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1117), int32(_a_F_spgdoinsert_17))
	mBase = m.M
	v3392 = m.ExcPending
	if v3392 != 0 {
		goto L3
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
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_24), int32(0))
	mBase = m.M
	v3400 = m.ExcPending
	if v3400 != 0 {
		goto L3
	} else {
		goto L471
	}
L471:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(957), int32(_a_F_spgdoinsert_17))
	mBase = m.M
	v3405 = m.ExcPending
	if v3405 != 0 {
		goto L3
	} else {
		goto L472
	}
L472:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L473:
	;
	v3462 = v3426 & int32(3)
	v3463 = int32(0)
	v3464 = *(*int32)(unsafe.Add(mBase, uint32(v56)+776))
	if base.Ui32(int32(4)) <= base.Ui32(v3426) {
		goto L474
	} else {
		goto L475
	}
L474:
	;
	v3477 = v3463
	v3492 = int32(0)
	goto L477
L475:
	;
	v3575 = v3463
	goto L476
L476:
	;
	v3628 = v3575
	v3635 = v3463
	goto L481
L477:
	;
	v3525 = int32(2)
	v3528 = *(*int32)(unsafe.Add(mBase, uint32(v3464+v3477<<(uint(v3525)%32))))
	v3530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3024+v3528))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3477+v1433))) = uint8(v3530)
	v3533 = v3477 | int32(1)
	v3538 = *(*int32)(unsafe.Add(mBase, uint32(v3464+v3533<<(uint(v3525)%32))))
	v3540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3024+v3538))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1433+v3533))) = uint8(v3540)
	v3543 = v3477 | v3525
	v3548 = *(*int32)(unsafe.Add(mBase, uint32(v3464+v3543<<(uint(v3525)%32))))
	v3550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3024+v3548))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1433+v3543))) = uint8(v3550)
	v3553 = v3477 | int32(3)
	v3558 = *(*int32)(unsafe.Add(mBase, uint32(v3464+v3553<<(uint(v3525)%32))))
	v3560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3024+v3558))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1433+v3553))) = uint8(v3560)
	v3562 = int32(4)
	v3563 = v3477 + v3562
	v3565 = v3492 + v3562
	if v3565 != v3426&int32(2147483644) {
		v3477 = v3563
		v3492 = v3565
		goto L477
	} else {
		goto L479
	}
L478:
	;
	if v3462 == int32(0) {
		v3708 = v3426
		v3710 = v3021
		v3720 = v3438
		goto L376
	} else {
		goto L480
	}
L479:
	;
	goto L478
L480:
	;
	v3575 = v3563
	goto L476
L481:
	;
	v3679 = *(*int32)(unsafe.Add(mBase, uint32(v3464+v3628<<(uint(int32(2))%32))))
	v3681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3024+v3679))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3628+v1433))) = uint8(v3681)
	v3683 = int32(1)
	v3686 = v3635 + v3683
	if v3686 != v3462 {
		v3628 = v3628 + v3683
		v3635 = v3686
		goto L481
	} else {
		goto L483
	}
L482:
	;
	v3708 = v3426
	v3710 = v3021
	v3720 = v3438
	goto L376
L483:
	;
	goto L482
L484:
	;
	v3749 = int32(_a_F_spgdoinsert_4)
	v3751 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v3751 + int32(1)
	v3755 = int32(0)
	if base.Ui32(v654) < base.Ui32(int32(2)) {
		v3830 = v3755
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v3832 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+756)) = v3832
	if v3832 < v3708 {
		goto L508
	} else {
		goto L509
	}
L486:
	;
	v3756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v3756 == int32(1) {
		goto L488
	} else {
		goto L489
	}
L487:
	;
	if v1403&int32(1) != 0 {
		v3830 = v3755
		goto L485
	} else {
		goto L506
	}
L488:
	;
	v3759 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+16)))
	v3761 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420+v3759)+4)))
	v3763 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v3763) {
		goto L491
	} else {
		goto L492
	}
L489:
	;
	goto L490
L490:
	;
	if v1403&int32(1) != 0 {
		v3830 = v3755
		goto L485
	} else {
		goto L500
	}
L491:
	;
	v3773 = int32(base.Ui32(v3763+int32(_a_F_spgdoinsert_12))>>(uint(int32(2))%32)) & int32(_a_F_spgdoinsert_0)
	goto L493
L492:
	;
	v3773 = int32(0)
	goto L493
L493:
	;
	if v1763+v3761 != v3773 {
		goto L487
	} else {
		goto L494
	}
L494:
	;
	if v401 < int32(0) {
		goto L497
	} else {
		goto L498
	}
L495:
	;
	v3801 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+570)) = uint8(v3801)
	v3830 = v3755
	goto L485
L496:
	;
	F_PageInit(m, v3792, int32(_a_F_spgdoinsert_25), int32(8))
	mBase = m.M
	v3796 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3792)+16)))
	v3797 = v3792 + v3796
	v3798 = int32(_a_F_spgdoinsert_26)
	*(*uint16)(unsafe.Add(mBase, uint32(v3797)+6)) = uint16(v3798)
	*(*uint16)(unsafe.Add(mBase, uint32(v3797))) = uint16(v288)
	goto L495
L497:
	;
	v3778 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v3784 = *(*int32)(unsafe.Add(mBase, uint32(v3778+(v401^int32(-1))<<(uint(int32(2))%32))))
	v3792 = v3784
	goto L496
L498:
	;
	goto L499
L499:
	;
	v3786 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v3792 = v3786 + v401<<(uint(int32(13))%32) + int32(-8192)
	goto L496
L500:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+566)) = uint16(v1763)
	if v1763 <= int32(0) {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	v3808 = int32(1)
	F_spgPageIndexMultiDelete(m, l1, v420, v1425, v1763, v3808, int32(3), int32(0), v3808)
	mBase = m.M
	v3813 = m.ExcPending
	if v3813 != 0 {
		goto L3
	} else {
		goto L504
	}
L502:
	;
	goto L503
L503:
	;
	v3814 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1425))))
	v3815 = int32(1)
	F_spgPageIndexMultiDelete(m, l1, v420, v1425, v1763, v3815, int32(3), int32(0), v3815)
	mBase = m.M
	v3820 = m.ExcPending
	if v3820 != 0 {
		goto L3
	} else {
		goto L505
	}
L504:
	;
	v3830 = v3755
	goto L485
L505:
	;
	v3830 = v3814
	goto L485
L506:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+566)) = uint16(v1763)
	v3824 = int32(3)
	F_spgPageIndexMultiDelete(m, l1, v420, v1425, v1763, v3824, v3824, int32(-1), int32(0))
	mBase = m.M
	v3829 = m.ExcPending
	if v3829 != 0 {
		goto L3
	} else {
		goto L507
	}
L507:
	;
	v3830 = v3755
	goto L485
L508:
	;
	v3842 = v3743
	v3847 = v3747
	goto L511
L509:
	;
	v4003 = v3747
	goto L510
L510:
	;
	if v3710 != 0 {
		goto L533
	} else {
		goto L534
	}
L511:
	;
	v3890 = v3842 << (uint(int32(2)) % 32)
	v3891 = v1431 + v3890
	v3892 = *(*int32)(unsafe.Add(mBase, uint32(v3891)))
	v3893 = v3842 + v1433
	v3894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3893))))
	if v3894 != 0 {
		goto L513
	} else {
		goto L514
	}
L512:
	;
	v4003 = v3988
	goto L510
L513:
	;
	v3895 = v3710
	goto L515
L514:
	;
	v3895 = v401
	goto L515
L515:
	;
	if v3895 < int32(0) {
		goto L517
	} else {
		goto L518
	}
L516:
	;
	v3915 = *(*int32)(unsafe.Add(mBase, uint32(v56)+776))
	v3917 = *(*int32)(unsafe.Add(mBase, uint32(v3915+v3890)))
	v3920 = v2537 + v3917<<(uint(int32(2))%32)
	v3921 = *(*int32)(unsafe.Add(mBase, uint32(v3920)))
	if v3921 == int32(0) {
		goto L521
	} else {
		goto L522
	}
L517:
	;
	v3899 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[1]))
	v3905 = *(*int32)(unsafe.Add(mBase, uint32(v3899+(v3895^int32(-1))<<(uint(int32(6))%32))+16))
	v3914 = v3905
	goto L516
L518:
	;
	goto L519
L519:
	;
	v3907 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[2]))
	v3913 = *(*int32)(unsafe.Add(mBase, uint32(v3907+v3895<<(uint(int32(6))%32)+int32(-64))+16))
	v3914 = v3913
	goto L516
L520:
	;
	if v3895 < int32(0) {
		goto L525
	} else {
		goto L526
	}
L521:
	;
	v3935 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3892)+4)))
	v3937 = v3935 & int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v3892)+4)) = uint16(v3937)
	goto L520
L522:
	;
	v3924 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3921)+4)))
	if v3924 == int32(0) {
		goto L521
	} else {
		goto L523
	}
L523:
	;
	v3927 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3892)+4)))
	v3932 = v3927&int32(_a_F_spgdoinsert_5) | v3924&int32(_a_F_spgdoinsert_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v3892)+4)) = uint16(v3932)
	goto L520
L524:
	;
	v3961 = *(*int32)(unsafe.Add(mBase, uint32(v3892)))
	v3966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3893))))
	v3970 = F_SpGistPageAddNewItem(m, v3960, v3892, int32(base.Ui32(v3961)>>(uint(int32(2))%32)), v56+int32(756)+v3966<<(uint(int32(1))%32))
	mBase = m.M
	v3971 = m.ExcPending
	if v3971 != 0 {
		goto L3
	} else {
		goto L528
	}
L525:
	;
	v3946 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v3952 = *(*int32)(unsafe.Add(mBase, uint32(v3946+(v3895^int32(-1))<<(uint(int32(2))%32))))
	v3960 = v3952
	goto L524
L526:
	;
	goto L527
L527:
	;
	v3954 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v3960 = v3954 + v3895<<(uint(int32(13))%32) + int32(-8192)
	goto L524
L528:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1427+v3842<<(uint(int32(1))%32)))) = uint16(v3970)
	v3973 = *(*int32)(unsafe.Add(mBase, uint32(v3920)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3973)+4)) = uint16(v3970)
	*(*uint16)(unsafe.Add(mBase, uint32(v3973)+2)) = uint16(v3914)
	v3977 = int32(base.Ui32(v3914) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v3973))) = uint16(v3977)
	v3979 = *(*int32)(unsafe.Add(mBase, uint32(v3891)))
	v3980 = *(*int32)(unsafe.Add(mBase, uint32(v3979)))
	v3982 = int32(base.Ui32(v3980) >> (uint(int32(2)) % 32))
	if v3982 != 0 {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	base.MemoryCopy(m, v3847, v3979, v3982)
	goto L531
L530:
	;
	goto L531
L531:
	;
	v3984 = *(*int32)(unsafe.Add(mBase, uint32(v3891)))
	v3985 = *(*int32)(unsafe.Add(mBase, uint32(v3984)))
	v3988 = v3847 + int32(base.Ui32(v3985)>>(uint(int32(2))%32))
	v3990 = v3842 + int32(1)
	if v3990 != v3708 {
		v3842 = v3990
		v3847 = v3988
		goto L511
	} else {
		goto L532
	}
L532:
	;
	goto L512
L533:
	;
	F_MarkBufferDirty(m, v3710)
	mBase = m.M
	v4046 = m.ExcPending
	if v4046 != 0 {
		goto L3
	} else {
		goto L536
	}
L534:
	;
	goto L535
L535:
	;
	v4047 = int32(0)
	if base.B2i32(v2985 == v4047)|base.B2i32(v2985 != v313) == v4047 {
		goto L539
	} else {
		goto L540
	}
L536:
	;
	goto L535
L537:
	;
	F_MarkBufferDirty(m, v401)
	mBase = m.M
	v4212 = m.ExcPending
	if v4212 != 0 {
		goto L3
	} else {
		goto L572
	}
L538:
	;
	v4206 = v313
	v4207 = v314
	v4208 = v4055
	v4209 = v322
	v4210 = v401
	goto L537
L539:
	;
	v4053 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2673)+4)))
	v4055 = F_SpGistPageAddNewItem(m, v314, v2673, v4053, int32(0))
	mBase = m.M
	v4056 = m.ExcPending
	if v4056 != 0 {
		goto L3
	} else {
		goto L542
	}
L540:
	;
	goto L541
L541:
	;
	if v313 != 0 {
		goto L545
	} else {
		goto L546
	}
L542:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+580)) = uint16(v330)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+578)) = uint16(v319)
	v4059 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+576)) = uint8(v4059)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+572)) = uint16(v4055)
	F_saveNodeLink(m, v56+int32(412), v322, v4055)
	mBase = m.M
	v4065 = m.ExcPending
	if v4065 != 0 {
		goto L3
	} else {
		goto L543
	}
L543:
	;
	if v3830 == int32(0) {
		goto L538
	} else {
		goto L544
	}
L544:
	;
	v4071 = *(*int32)(unsafe.Add(mBase, uint32(v420+v3830<<(uint(int32(2))%32))+20))
	v4074 = v420 + v4071&int32(_a_F_spgdoinsert_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v4074)+10)) = uint16(v4055)
	*(*uint16)(unsafe.Add(mBase, uint32(v4074)+8)) = uint16(v322)
	v4078 = int32(base.Ui32(v322) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v4074)+6)) = uint16(v4078)
	goto L538
L545:
	;
	if v2985 < int32(0) {
		goto L549
	} else {
		goto L550
	}
L546:
	;
	goto L547
L547:
	;
	if v401 < int32(0) {
		goto L562
	} else {
		goto L563
	}
L548:
	;
	if v2985 < int32(0) {
		goto L553
	} else {
		goto L554
	}
L549:
	;
	v4083 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[1]))
	v4089 = *(*int32)(unsafe.Add(mBase, uint32(v4083+(v2985^int32(-1))<<(uint(int32(6))%32))+16))
	v4098 = v4089
	goto L548
L550:
	;
	goto L551
L551:
	;
	v4091 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[2]))
	v4097 = *(*int32)(unsafe.Add(mBase, uint32(v4091+v2985<<(uint(int32(6))%32)+int32(-64))+16))
	v4098 = v4097
	goto L548
L552:
	;
	v4117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2673)+4)))
	v4119 = F_SpGistPageAddNewItem(m, v4116, v2673, v4117, int32(0))
	mBase = m.M
	v4120 = m.ExcPending
	if v4120 != 0 {
		goto L3
	} else {
		goto L556
	}
L553:
	;
	v4102 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v4108 = *(*int32)(unsafe.Add(mBase, uint32(v4102+(v2985^int32(-1))<<(uint(int32(2))%32))))
	v4116 = v4108
	goto L552
L554:
	;
	goto L555
L555:
	;
	v4110 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v4116 = v4110 + v2985<<(uint(int32(13))%32) + int32(-8192)
	goto L552
L556:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+572)) = uint16(v4119)
	F_MarkBufferDirty(m, v2985)
	mBase = m.M
	v4123 = m.ExcPending
	if v4123 != 0 {
		goto L3
	} else {
		goto L557
	}
L557:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+580)) = uint16(v330)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+578)) = uint16(v319)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+576)) = uint8(base.B2i32(v2985 == v313))
	F_saveNodeLink(m, v56+int32(412), v4098, v4119)
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		goto L3
	} else {
		goto L558
	}
L558:
	;
	if v3830 == int32(0) {
		v4206 = v2985
		v4207 = v4116
		v4208 = v4119
		v4209 = v4098
		v4210 = v401
		goto L537
	} else {
		goto L559
	}
L559:
	;
	v4137 = *(*int32)(unsafe.Add(mBase, uint32(v420+v3830<<(uint(int32(2))%32))+20))
	v4140 = v420 + v4137&int32(_a_F_spgdoinsert_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v4140)+10)) = uint16(v4119)
	*(*uint16)(unsafe.Add(mBase, uint32(v4140)+8)) = uint16(v4098)
	v4144 = int32(base.Ui32(v4098) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v4140)+6)) = uint16(v4144)
	v4206 = v2985
	v4207 = v4116
	v4208 = v4119
	v4209 = v4098
	v4210 = v401
	goto L537
L560:
	;
	v4172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+576)) = uint8(v4172)
	v4174 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+574)) = uint8(v4174)
	v4176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2673)+4)))
	v4179 = F_PageAddItemExtended(m, v420, v2673, v4176, v4172, v4172)
	mBase = m.M
	v4180 = m.ExcPending
	if v4180 != 0 {
		goto L3
	} else {
		goto L565
	}
L561:
	;
	F_PageInit(m, v4163, int32(_a_F_spgdoinsert_25), int32(8))
	mBase = m.M
	v4167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4163)+16)))
	v4168 = v4163 + v4167
	v4169 = int32(_a_F_spgdoinsert_26)
	*(*uint16)(unsafe.Add(mBase, uint32(v4168)+6)) = uint16(v4169)
	*(*uint16)(unsafe.Add(mBase, uint32(v4168))) = uint16(v285)
	goto L560
L562:
	;
	v4149 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v4155 = *(*int32)(unsafe.Add(mBase, uint32(v4149+(v401^int32(-1))<<(uint(int32(2))%32))))
	v4163 = v4155
	goto L561
L563:
	;
	goto L564
L564:
	;
	v4157 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v4163 = v4157 + v401<<(uint(int32(13))%32) + int32(-8192)
	goto L561
L565:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+572)) = uint16(v4179)
	if v4179 == int32(1) {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	v4184 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+578)) = v4184
	v4206 = v401
	v4207 = v420
	v4208 = int32(1)
	v4209 = v402
	v4210 = v4184
	goto L537
L567:
	;
	goto L568
L568:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4191 = m.ExcPending
	if v4191 != 0 {
		goto L3
	} else {
		goto L569
	}
L569:
	;
	v4192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2673)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+176)) = v4192
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_8), v56+int32(176))
	mBase = m.M
	v4198 = m.ExcPending
	if v4198 != 0 {
		goto L3
	} else {
		goto L570
	}
L570:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1347), int32(_a_F_spgdoinsert_17))
	mBase = m.M
	v4203 = m.ExcPending
	if v4203 != 0 {
		goto L3
	} else {
		goto L571
	}
L571:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L572:
	;
	v4213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4213)+118)))
	if v4214 != int32(112) {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v4341 = int32(_a_F_spgdoinsert_4)
	v4343 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v4343 - int32(1)
	if v3710 != 0 {
		goto L626
	} else {
		goto L627
	}
L574:
	;
	v4218 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[6]))
	if v4218 <= int32(0) {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	v4221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v4221 != 0 {
		goto L573
	} else {
		goto L578
	}
L576:
	;
	goto L577
L577:
	;
	v4223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v4223 != 0 {
		goto L573
	} else {
		goto L580
	}
L578:
	;
	v4222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4222 != 0 {
		goto L573
	} else {
		goto L579
	}
L579:
	;
	goto L577
L580:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v4225 = m.ExcPending
	if v4225 != 0 {
		goto L3
	} else {
		goto L581
	}
L581:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+568)) = uint16(v3708)
	F_XLogRegisterData(m, v56+int32(564), int32(28))
	mBase = m.M
	v4231 = m.ExcPending
	if v4231 != 0 {
		goto L3
	} else {
		goto L582
	}
L582:
	;
	v4232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+566)))
	F_XLogRegisterData(m, v1425, v4232<<(uint(int32(1))%32))
	mBase = m.M
	v4236 = m.ExcPending
	if v4236 != 0 {
		goto L3
	} else {
		goto L583
	}
L583:
	;
	v4237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+568)))
	F_XLogRegisterData(m, v1427, v4237<<(uint(int32(1))%32))
	mBase = m.M
	v4241 = m.ExcPending
	if v4241 != 0 {
		goto L3
	} else {
		goto L584
	}
L584:
	;
	v4242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+568)))
	F_XLogRegisterData(m, v1433, v4242)
	mBase = m.M
	v4244 = m.ExcPending
	if v4244 != 0 {
		goto L3
	} else {
		goto L585
	}
L585:
	;
	v4245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2673)+4)))
	F_XLogRegisterData(m, v2673, v4245)
	mBase = m.M
	v4247 = m.ExcPending
	if v4247 != 0 {
		goto L3
	} else {
		goto L586
	}
L586:
	;
	F_XLogRegisterData(m, v3747, v4003-v3747)
	mBase = m.M
	v4250 = m.ExcPending
	if v4250 != 0 {
		goto L3
	} else {
		goto L587
	}
L587:
	;
	if v4210 != 0 {
		goto L588
	} else {
		goto L589
	}
L588:
	;
	v4254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+570)))
	if v4254 != 0 {
		goto L591
	} else {
		goto L592
	}
L589:
	;
	goto L590
L590:
	;
	if v3710 != 0 {
		goto L595
	} else {
		goto L596
	}
L591:
	;
	v4255 = int32(14)
	goto L593
L592:
	;
	v4255 = int32(8)
	goto L593
L593:
	;
	F_XLogRegisterBuffer(m, int32(0), v4210, v4255)
	mBase = m.M
	v4257 = m.ExcPending
	if v4257 != 0 {
		goto L3
	} else {
		goto L594
	}
L594:
	;
	goto L590
L595:
	;
	v4261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+571)))
	if v4261 != 0 {
		goto L598
	} else {
		goto L599
	}
L596:
	;
	goto L597
L597:
	;
	v4268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+574)))
	if v4268 != 0 {
		goto L602
	} else {
		goto L603
	}
L598:
	;
	v4262 = int32(14)
	goto L600
L599:
	;
	v4262 = int32(8)
	goto L600
L600:
	;
	F_XLogRegisterBuffer(m, int32(1), v3710, v4262)
	mBase = m.M
	v4264 = m.ExcPending
	if v4264 != 0 {
		goto L3
	} else {
		goto L601
	}
L601:
	;
	goto L597
L602:
	;
	v4269 = int32(14)
	goto L604
L603:
	;
	v4269 = int32(8)
	goto L604
L604:
	;
	F_XLogRegisterBuffer(m, int32(2), v4206, v4269)
	mBase = m.M
	v4271 = m.ExcPending
	if v4271 != 0 {
		goto L3
	} else {
		goto L605
	}
L605:
	;
	v4272 = int32(0)
	if base.B2i32(v313 == v4272)|base.B2i32(v4206 == v313) == v4272 {
		goto L606
	} else {
		goto L607
	}
L606:
	;
	F_XLogRegisterBuffer(m, int32(3), v313, int32(8))
	mBase = m.M
	v4281 = m.ExcPending
	if v4281 != 0 {
		goto L3
	} else {
		goto L609
	}
L607:
	;
	goto L608
L608:
	;
	v4284 = F_XLogInsert(m, int32(16), int32(80))
	mBase = m.M
	v4285 = m.ExcPending
	if v4285 != 0 {
		goto L3
	} else {
		goto L610
	}
L609:
	;
	goto L608
L610:
	;
	if v3710 != 0 {
		goto L611
	} else {
		goto L612
	}
L611:
	;
	if v3710 < int32(0) {
		goto L615
	} else {
		goto L616
	}
L612:
	;
	goto L613
L613:
	;
	if v4210 != 0 {
		goto L618
	} else {
		goto L619
	}
L614:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4303))) = base.I64_rotr(v4284, int64(32))
	goto L613
L615:
	;
	v4289 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v4295 = *(*int32)(unsafe.Add(mBase, uint32(v4289+(v3710^int32(-1))<<(uint(int32(2))%32))))
	v4303 = v4295
	goto L614
L616:
	;
	goto L617
L617:
	;
	v4297 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v4303 = v4297 + v3710<<(uint(int32(13))%32) + int32(-8192)
	goto L614
L618:
	;
	if v4210 < int32(0) {
		goto L622
	} else {
		goto L623
	}
L619:
	;
	goto L620
L620:
	;
	v4330 = base.I32_wrap_i64(int64(base.Ui64(v4284) >> (uint(int64(32)) % 64)))
	v4331 = base.I32_wrap_i64(v4284)
	*(*int32)(unsafe.Add(mBase, uint32(v4207)+4)) = v4331
	*(*int32)(unsafe.Add(mBase, uint32(v4207))) = v4330
	if v313 == int32(0) {
		goto L573
	} else {
		goto L625
	}
L621:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4324))) = base.I64_rotr(v4284, int64(32))
	goto L620
L622:
	;
	v4310 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v4316 = *(*int32)(unsafe.Add(mBase, uint32(v4310+(v4210^int32(-1))<<(uint(int32(2))%32))))
	v4324 = v4316
	goto L621
L623:
	;
	goto L624
L624:
	;
	v4318 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v4324 = v4318 + v4210<<(uint(int32(13))%32) + int32(-8192)
	goto L621
L625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+4)) = v4331
	*(*int32)(unsafe.Add(mBase, uint32(v314))) = v4330
	goto L573
L626:
	;
	F_SpGistSetLastUsedPage(m, l0, v3710)
	mBase = m.M
	v4348 = m.ExcPending
	if v4348 != 0 {
		goto L3
	} else {
		goto L629
	}
L627:
	;
	goto L628
L628:
	;
	if v4210 != 0 {
		goto L631
	} else {
		goto L632
	}
L629:
	;
	F_UnlockReleaseBuffer(m, v3710)
	mBase = m.M
	v4350 = m.ExcPending
	if v4350 != 0 {
		goto L3
	} else {
		goto L630
	}
L630:
	;
	goto L628
L631:
	;
	F_SpGistSetLastUsedPage(m, l0, v4210)
	mBase = m.M
	v4352 = m.ExcPending
	if v4352 != 0 {
		goto L3
	} else {
		goto L634
	}
L632:
	;
	goto L633
L633:
	;
	if v3720 != 0 {
		goto L636
	} else {
		goto L637
	}
L634:
	;
	F_UnlockReleaseBuffer(m, v4210)
	mBase = m.M
	v4354 = m.ExcPending
	if v4354 != 0 {
		goto L3
	} else {
		goto L635
	}
L635:
	;
	goto L633
L636:
	;
	v6776 = int32(1)
	v6779 = v4206
	goto L30
L637:
	;
	goto L638
L638:
	;
	F_pfree(m, v437)
	mBase = m.M
	v4357 = m.ExcPending
	if v4357 != 0 {
		goto L3
	} else {
		goto L639
	}
L639:
	;
	v4367 = v4206
	v4368 = v4207
	v4373 = v4208
	v4376 = v4209
	goto L64
L640:
	;
	v4423 = v4367
	v4424 = v4368
	v4429 = v4373
	v4432 = v4376
	goto L652
L641:
	;
	v6561 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+411)) = uint8(v6561)
	v6564 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v6564 == v6561 {
		v310 = v5028
		v311 = v5021
		v312 = v6558
		v313 = v4423
		v314 = v4424
		v315 = v6559
		v319 = v4429
		v322 = v4432
		v330 = v4680
		v334 = v5040
		v345 = v5029 + v345
		goto L62
	} else {
		goto L992
	}
L642:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6545 = m.ExcPending
	if v6545 != 0 {
		goto L3
	} else {
		goto L989
	}
L643:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6532 = m.ExcPending
	if v6532 != 0 {
		goto L3
	} else {
		goto L986
	}
L644:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6516 = m.ExcPending
	if v6516 != 0 {
		goto L3
	} else {
		goto L983
	}
L645:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6500 = m.ExcPending
	if v6500 != 0 {
		goto L3
	} else {
		goto L980
	}
L646:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6482 = m.ExcPending
	if v6482 != 0 {
		goto L3
	} else {
		goto L977
	}
L647:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6469 = m.ExcPending
	if v6469 != 0 {
		goto L3
	} else {
		goto L974
	}
L648:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6456 = m.ExcPending
	if v6456 != 0 {
		goto L3
	} else {
		goto L971
	}
L649:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6440 = m.ExcPending
	if v6440 != 0 {
		goto L3
	} else {
		goto L968
	}
L650:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6427 = m.ExcPending
	if v6427 != 0 {
		goto L3
	} else {
		goto L965
	}
L651:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6414 = m.ExcPending
	if v6414 != 0 {
		goto L3
	} else {
		goto L962
	}
L652:
	;
	v4467 = int32(1)
	v4468 = v4432 - v4467
	v4472 = base.I32_rem_u_s(v4432+v4467, int32(3))
	v4474 = v4429 & int32(_a_F_spgdoinsert_0)
	v4479 = v4424 + v4474<<(uint(int32(2))%32) + int32(20)
	goto L654
L653:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6398 = m.ExcPending
	if v6398 != 0 {
		goto L3
	} else {
		goto L959
	}
L654:
	;
	v4533 = *(*int32)(unsafe.Add(mBase, uint32(v4479)))
	v4534 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+388)) = v345
	v4536 = *(*int32)(unsafe.Add(mBase, uint32(v56)+432))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+384)) = v4536
	*(*int32)(unsafe.Add(mBase, uint32(v56)+380)) = v4534
	v4541 = v4424 + v4533&int32(_a_F_spgdoinsert_6)
	v4542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4541))))
	v4546 = int32(base.Ui32(v4542)>>(uint(int32(2))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+392)) = uint8(v4546)
	v4548 = *(*int32)(unsafe.Add(mBase, uint32(v4541)))
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+393)) = uint8(base.B2i32(base.Ui32(int32(_a_F_spgdoinsert_0)) < base.Ui32(v4548)))
	v4553 = *(*int32)(unsafe.Add(mBase, uint32(v4541)))
	if base.Ui32(v4553) < base.Ui32(int32(_a_F_spgdoinsert_27)) {
		v4563 = int32(0)
		goto L656
	} else {
		goto L657
	}
L655:
	;
	goto L653
L656:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+396)) = v4563
	v4565 = *(*int32)(unsafe.Add(mBase, uint32(v4541)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+400)) = int32(base.Ui32(v4565)>>(uint(int32(3))%32)) & int32(_a_F_spgdoinsert_22)
	v4571 = F_spgExtractNodeLabels(m, l1, v4541)
	mBase = m.M
	v4572 = m.ExcPending
	if v4572 != 0 {
		goto L3
	} else {
		goto L659
	}
L657:
	;
	v4557 = v4541 + int32(8)
	v4558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+50)))
	if v4558 != int32(1) {
		v4563 = v4557
		goto L656
	} else {
		goto L658
	}
L658:
	;
	v4561 = *(*int32)(unsafe.Add(mBase, uint32(v4557)))
	v4563 = v4561
	goto L656
L659:
	;
	v4573 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v56)+344)) = v4573
	*(*int32)(unsafe.Add(mBase, uint32(v56)+404)) = v4571
	*(*int64)(unsafe.Add(mBase, uint32(v56)+352)) = v4573
	*(*int64)(unsafe.Add(mBase, uint32(v56)+360)) = v4573
	*(*int64)(unsafe.Add(mBase, uint32(v56)+368)) = v4573
	if v60 == int32(0) {
		goto L661
	} else {
		goto L662
	}
L660:
	;
	v4597 = *(*int32)(unsafe.Add(mBase, uint32(v4541)))
	if v4597&int32(4) == int32(0) {
		v4675 = v4596
		goto L665
	} else {
		goto L666
	}
L661:
	;
	v4584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v4585 = *(*int32)(unsafe.Add(mBase, uint32(v4584)))
	v4590 = F_FunctionCall2Coll(m, v100, v4585, v56+int32(380), v56+int32(344))
	mBase = m.M
	v4591 = m.ExcPending
	if v4591 != 0 {
		goto L3
	} else {
		goto L664
	}
L662:
	;
	goto L663
L663:
	;
	v4593 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+344)) = v4593
	v4596 = v4593
	goto L660
L664:
	;
	v4592 = *(*int32)(unsafe.Add(mBase, uint32(v56)+344))
	v4596 = v4592
	goto L660
L665:
	;
	if v4675 != int32(3) {
		goto L680
	} else {
		goto L681
	}
L666:
	;
	switch v4596 - int32(1) {
	case 0:
		goto L667
	case 1:
		goto L668
	default:
		v4675 = v4596
		goto L665
	}
L667:
	;
	v4617 = int32(_a_F_spgdoinsert_28)
	v4618 = int64(0)
	v4625 = base.I64_extend_i32_s(int32(base.Ui32(v4597)>>(uint(int32(3))%32))&int32(_a_F_spgdoinsert_22) - int32(1))
	if base.Ui64(v4625) <= base.Ui64(v4618) {
		goto L673
	} else {
		goto L674
	}
L668:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4607 = m.ExcPending
	if v4607 != 0 {
		goto L3
	} else {
		goto L669
	}
L669:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_29), int32(0))
	mBase = m.M
	v4611 = m.ExcPending
	if v4611 != 0 {
		goto L3
	} else {
		goto L670
	}
L670:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(2212), int32(_a_F_spgdoinsert_19))
	mBase = m.M
	v4616 = m.ExcPending
	if v4616 != 0 {
		goto L3
	} else {
		goto L671
	}
L671:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L672:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v56)+348)) = uint32(v4672)
	v4674 = *(*int32)(unsafe.Add(mBase, uint32(v56)+344))
	v4675 = v4674
	goto L665
L673:
	;
	v4672 = v4618
	goto L672
L674:
	;
	goto L675
L675:
	;
	v4632 = v4625 - v4618
	v4634 = *(*int64)(unsafe.Add(mBase, _c_F_spgdoinsert[7]))
	v4635 = *(*int64)(unsafe.Add(mBase, _c_F_spgdoinsert[8]))
	v4638 = v4635
	v4640 = v4634
	goto L676
L676:
	;
	v4644 = v4638 ^ v4640
	v4646 = base.I64_rotl(v4644, int64(37))
	v4654 = v4644 ^ (v4644<<(uint(int64(16))%64) ^ base.I64_rotl(v4638, int64(24)))
	v4659 = int64(base.Ui64(base.I64_rotl(v4638*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v4632)) % 64))
	if base.Ui64(v4632) < base.Ui64(v4659) {
		v4638 = v4654
		v4640 = v4646
		goto L676
	} else {
		goto L678
	}
L677:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_spgdoinsert[7])) = v4646
	*(*int64)(unsafe.Add(mBase, _c_F_spgdoinsert[8])) = v4654
	v4672 = v4618 + v4659
	goto L672
L678:
	;
	goto L677
L679:
	;
	goto L655
L680:
	;
	switch v4675 - int32(1) {
	case 0:
		goto L684
	case 1:
		goto L683
	default:
		goto L679
	}
L681:
	;
	goto L682
L682:
	;
	v5584 = *(*int32)(unsafe.Add(mBase, uint32(v56)+356))
	v5585 = int32(-8192)
	if base.Ui32(v5584+v5585) <= base.Ui32(v5585) {
		goto L645
	} else {
		goto L848
	}
L683:
	;
	v5086 = *(*int32)(unsafe.Add(mBase, uint32(v56)+404))
	if v5086 == int32(0) {
		goto L651
	} else {
		goto L730
	}
L684:
	;
	v4680 = *(*int32)(unsafe.Add(mBase, uint32(v56)+348))
	v4681 = int32(0)
	if base.B2i32(v313 == v4681)|base.B2i32(v4423 == v313) == v4681 {
		goto L685
	} else {
		goto L686
	}
L685:
	;
	F_SpGistSetLastUsedPage(m, l0, v313)
	mBase = m.M
	v4688 = m.ExcPending
	if v4688 != 0 {
		goto L3
	} else {
		goto L688
	}
L686:
	;
	goto L687
L687:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+428)) = v4680
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+424)) = uint16(v4429)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+420)) = v4424
	*(*int32)(unsafe.Add(mBase, uint32(v56)+416)) = v4423
	*(*int32)(unsafe.Add(mBase, uint32(v56)+412)) = v4432
	v4696 = *(*int32)(unsafe.Add(mBase, uint32(v4541)))
	v4701 = v4541 + int32(base.Ui32(v4696)>>(uint(int32(16))%32)) + int32(8)
	if v4680 == int32(0) {
		goto L691
	} else {
		goto L692
	}
L688:
	;
	F_UnlockReleaseBuffer(m, v313)
	mBase = m.M
	v4690 = m.ExcPending
	if v4690 != 0 {
		goto L3
	} else {
		goto L689
	}
L689:
	;
	goto L687
L690:
	;
	v5021 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4974)+4)))
	if v5021 != 0 {
		goto L710
	} else {
		goto L711
	}
L691:
	;
	if v4680 != 0 {
		goto L32
	} else {
		goto L709
	}
L692:
	;
	v4707 = int32(base.Ui32(v4696)>>(uint(int32(3))%32)) & int32(_a_F_spgdoinsert_22)
	if v4707 == int32(0) {
		goto L691
	} else {
		goto L693
	}
L693:
	;
	v4710 = int32(1)
	v4711 = v4680 - v4710
	v4713 = v4707 - v4710
	if base.Ui32(v4711) < base.Ui32(v4713) {
		goto L694
	} else {
		goto L695
	}
L694:
	;
	v4715 = v4711
	goto L696
L695:
	;
	v4715 = v4713
	goto L696
L696:
	;
	v4717 = v4715 + int32(1)
	v4718 = int32(3)
	v4719 = v4717 & v4718
	if base.Ui32(v4718) <= base.Ui32(v4715) {
		goto L698
	} else {
		goto L699
	}
L697:
	;
	if v4717 == v4680 {
		v4974 = v4919
		goto L690
	} else {
		goto L708
	}
L698:
	;
	v4731 = v4701
	v4732 = int32(0)
	goto L701
L699:
	;
	v4805 = v4701
	goto L700
L700:
	;
	v4859 = v4805
	v4860 = int32(0)
	goto L705
L701:
	;
	v4778 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4731)+6)))
	v4779 = int32(_a_F_spgdoinsert_22)
	v4781 = v4731 + v4778&v4779
	v4782 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4781)+6)))
	v4785 = v4781 + v4782&v4779
	v4786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4785)+6)))
	v4789 = v4785 + v4786&v4779
	v4790 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4789)+6)))
	v4793 = v4789 + v4790&v4779
	v4795 = v4732 + int32(4)
	if v4795 != v4717&int32(-4) {
		v4731 = v4793
		v4732 = v4795
		goto L701
	} else {
		goto L703
	}
L702:
	;
	if v4719 == int32(0) {
		v4919 = v4793
		goto L697
	} else {
		goto L704
	}
L703:
	;
	goto L702
L704:
	;
	v4805 = v4793
	goto L700
L705:
	;
	v4906 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4859)+6)))
	v4909 = v4859 + v4906&int32(_a_F_spgdoinsert_22)
	v4911 = v4860 + int32(1)
	if v4911 != v4719 {
		v4859 = v4909
		v4860 = v4911
		goto L705
	} else {
		goto L707
	}
L706:
	;
	v4919 = v4909
	goto L697
L707:
	;
	goto L706
L708:
	;
	goto L32
L709:
	;
	v4974 = v4701
	goto L690
L710:
	;
	v5022 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4974)+2)))
	v5023 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4974))))
	v5028 = v5022 | v5023<<(uint(int32(16))%32)
	goto L712
L711:
	;
	v5028 = int32(-1)
	goto L712
L712:
	;
	v5029 = *(*int32)(unsafe.Add(mBase, uint32(v56)+352))
	if v60 == int32(0) {
		goto L713
	} else {
		goto L714
	}
L713:
	;
	v5032 = *(*int32)(unsafe.Add(mBase, uint32(v56)+356))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+432)) = v5032
	v5036 = F_SpGistGetLeafTupleSize(m, v58, v56+int32(432), l4)
	mBase = m.M
	v5037 = m.ExcPending
	if v5037 != 0 {
		goto L3
	} else {
		goto L716
	}
L714:
	;
	v5040 = v334
	goto L715
L715:
	;
	if base.Ui32(v5040) < base.Ui32(int32(_a_F_spgdoinsert_20)) {
		goto L717
	} else {
		goto L718
	}
L716:
	;
	v5040 = v5036 + int32(4)
	goto L715
L717:
	;
	v6558 = v312
	v6559 = v315
	goto L641
L718:
	;
	goto L719
L719:
	;
	if v60 != 0 {
		goto L720
	} else {
		goto L721
	}
L720:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5059 = m.ExcPending
	if v5059 != 0 {
		goto L3
	} else {
		goto L725
	}
L721:
	;
	v5043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v5043&int32(1) == int32(0) {
		goto L720
	} else {
		goto L722
	}
L722:
	;
	if v5040 < v312 {
		v6558 = v5040
		v6559 = int32(0)
		goto L641
	} else {
		goto L723
	}
L723:
	;
	v5051 = v315 + int32(1)
	if v5051 < int32(10) {
		v6558 = v312
		v6559 = v5051
		goto L641
	} else {
		goto L724
	}
L724:
	;
	goto L720
L725:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v5062 = m.ExcPending
	if v5062 != 0 {
		goto L3
	} else {
		goto L726
	}
L726:
	;
	v5063 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+36)) = int32(_a_F_spgdoinsert_30)
	v5066 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+32)) = v5040 - v5066
	*(*int32)(unsafe.Add(mBase, uint32(v56)+40)) = v5063 + v5066
	F_errmsg(m, int32(_a_F_spgdoinsert_31), v56+int32(32))
	mBase = m.M
	v5076 = m.ExcPending
	if v5076 != 0 {
		goto L3
	} else {
		goto L727
	}
L727:
	;
	F_errhint(m, int32(_a_F_spgdoinsert_32), int32(0))
	mBase = m.M
	v5080 = m.ExcPending
	if v5080 != 0 {
		goto L3
	} else {
		goto L728
	}
L728:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(2282), int32(_a_F_spgdoinsert_19))
	mBase = m.M
	v5085 = m.ExcPending
	if v5085 != 0 {
		goto L3
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
	v5089 = *(*int32)(unsafe.Add(mBase, uint32(v56)+348))
	v5090 = *(*int32)(unsafe.Add(mBase, uint32(v4541)))
	v5094 = int32(base.Ui32(v5090)>>(uint(int32(3))%32)) & int32(_a_F_spgdoinsert_22)
	v5095 = *(*int32)(unsafe.Add(mBase, uint32(v56)+352))
	if v5095 < int32(0) {
		goto L732
	} else {
		goto L733
	}
L731:
	;
	v5101 = v4541 + int32(8)
	v5108 = F_palloc(m, int32(base.Ui32(v5090)>>(uint(int32(1))%32))&int32(_a_F_spgdoinsert_33)+int32(4))
	mBase = m.M
	v5109 = m.ExcPending
	if v5109 != 0 {
		goto L3
	} else {
		goto L736
	}
L732:
	;
	v5099 = v5094
	goto L731
L733:
	;
	goto L734
L734:
	;
	if base.Ui32(v5094) < base.Ui32(v5095) {
		goto L650
	} else {
		goto L735
	}
L735:
	;
	v5099 = v5095
	goto L731
L736:
	;
	v5110 = *(*int32)(unsafe.Add(mBase, uint32(v4541)))
	if v5110&int32(_a_F_spgdoinsert_21) != 0 {
		goto L737
	} else {
		goto L738
	}
L737:
	;
	v5123 = v5101 + int32(base.Ui32(v5110)>>(uint(int32(16))%32))
	v5124 = int32(0)
	goto L740
L738:
	;
	goto L739
L739:
	;
	v5241 = int32(0)
	v5246 = F_spgFormNodeTuple(m, l1, v5089, v5241)
	mBase = m.M
	v5247 = m.ExcPending
	if v5247 != 0 {
		goto L3
	} else {
		goto L747
	}
L740:
	;
	v5172 = v5108 + v5124<<(uint(int32(2))%32)
	if v5124 < v5099 {
		goto L743
	} else {
		goto L744
	}
L741:
	;
	goto L739
L742:
	;
	v5176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5123)+6)))
	v5177 = int32(_a_F_spgdoinsert_22)
	v5181 = v5124 + int32(1)
	v5182 = *(*int32)(unsafe.Add(mBase, uint32(v4541)))
	if base.Ui32(v5181) < base.Ui32(int32(base.Ui32(v5182)>>(uint(int32(3))%32))&v5177) {
		v5123 = v5123 + v5176&v5177
		v5124 = v5181
		goto L740
	} else {
		goto L746
	}
L743:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5172))) = v5123
	goto L742
L744:
	;
	goto L745
L745:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5172)+4)) = v5123
	goto L742
L746:
	;
	goto L741
L747:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5108+v5099<<(uint(int32(2))%32)))) = v5246
	v5249 = *(*int32)(unsafe.Add(mBase, uint32(v4541)))
	v5251 = int32(base.Ui32(v5249) >> (uint(int32(16)) % 32))
	if v5251 == int32(0) {
		v5258 = v5241
		goto L748
	} else {
		goto L749
	}
L748:
	;
	v5267 = F_spgFormInnerTuple(m, l1, base.B2i32(v5251 != int32(0)), v5258, int32(base.Ui32(v5249)>>(uint(int32(3))%32))&int32(_a_F_spgdoinsert_22)+int32(1), v5108)
	mBase = m.M
	v5268 = m.ExcPending
	if v5268 != 0 {
		goto L3
	} else {
		goto L753
	}
L749:
	;
	v5254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+50)))
	if v5254 == int32(1) {
		goto L750
	} else {
		goto L751
	}
L750:
	;
	v5257 = *(*int32)(unsafe.Add(mBase, uint32(v5101)))
	v5258 = v5257
	goto L748
L751:
	;
	goto L752
L752:
	;
	v5258 = v5101
	goto L748
L753:
	;
	v5269 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+636)) = v5269
	v5271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+624)) = uint16(v4429)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+640)) = uint8(v5271)
	*(*int64)(unsafe.Add(mBase, uint32(v56)+626)) = int64(4278190080)
	v5276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4424)+14)))
	v5277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4424)+12)))
	v5278 = v5276 - v5277
	v5279 = int32(0)
	if v5279 < v5278 {
		goto L756
	} else {
		goto L757
	}
L754:
	;
	v5579 = int32(0)
	v5581 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v5581 == v5579 {
		v4423 = v5572
		v4424 = v5573
		v4429 = v5576
		v4432 = v5577
		goto L652
	} else {
		goto L847
	}
L755:
	;
	v5283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5267)+4)))
	v5284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+4)))
	if base.Ui32(v5283-v5284) <= base.Ui32(v5282) {
		goto L759
	} else {
		goto L760
	}
L756:
	;
	v5282 = v5278
	goto L758
L757:
	;
	v5282 = v5279
	goto L758
L758:
	;
	goto L755
L759:
	;
	v5287 = int32(_a_F_spgdoinsert_4)
	v5289 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v5289 + int32(1)
	F_PageIndexTupleDelete(m, v4424, v4474)
	mBase = m.M
	v5294 = m.ExcPending
	if v5294 != 0 {
		goto L3
	} else {
		goto L762
	}
L760:
	;
	goto L761
L761:
	;
	if base.Ui32(v4468) <= base.Ui32(int32(1)) {
		goto L648
	} else {
		goto L779
	}
L762:
	;
	v5295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5267)+4)))
	v5297 = F_PageAddItemExtended(m, v4424, v5267, v5295, v4474, int32(0))
	mBase = m.M
	v5298 = m.ExcPending
	if v5298 != 0 {
		goto L3
	} else {
		goto L763
	}
L763:
	;
	if v5297 != v4474 {
		goto L649
	} else {
		goto L764
	}
L764:
	;
	F_MarkBufferDirty(m, v4423)
	mBase = m.M
	v5301 = m.ExcPending
	if v5301 != 0 {
		goto L3
	} else {
		goto L765
	}
L765:
	;
	v5302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5302)+118)))
	if v5303 != int32(112) {
		goto L766
	} else {
		goto L767
	}
L766:
	;
	v5334 = int32(_a_F_spgdoinsert_4)
	v5336 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v5336 - int32(1)
	v5572 = v4423
	v5573 = v4424
	v5576 = v4429
	v5577 = v4432
	goto L754
L767:
	;
	v5307 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[6]))
	if v5307 <= int32(0) {
		goto L768
	} else {
		goto L769
	}
L768:
	;
	v5310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v5310 != 0 {
		goto L766
	} else {
		goto L771
	}
L769:
	;
	goto L770
L770:
	;
	v5312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v5312 != 0 {
		goto L766
	} else {
		goto L773
	}
L771:
	;
	v5311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v5311 != 0 {
		goto L766
	} else {
		goto L772
	}
L772:
	;
	goto L770
L773:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v5314 = m.ExcPending
	if v5314 != 0 {
		goto L3
	} else {
		goto L774
	}
L774:
	;
	F_XLogRegisterData(m, v56+int32(624), int32(20))
	mBase = m.M
	v5319 = m.ExcPending
	if v5319 != 0 {
		goto L3
	} else {
		goto L775
	}
L775:
	;
	v5320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5267)+4)))
	F_XLogRegisterData(m, v5267, v5320)
	mBase = m.M
	v5322 = m.ExcPending
	if v5322 != 0 {
		goto L3
	} else {
		goto L776
	}
L776:
	;
	F_XLogRegisterBuffer(m, int32(0), v4423, int32(8))
	mBase = m.M
	v5326 = m.ExcPending
	if v5326 != 0 {
		goto L3
	} else {
		goto L777
	}
L777:
	;
	v5329 = F_XLogInsert(m, int32(16), int32(48))
	mBase = m.M
	v5330 = m.ExcPending
	if v5330 != 0 {
		goto L3
	} else {
		goto L778
	}
L778:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4424))) = base.I64_rotr(v5329, int64(32))
	goto L766
L779:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+632)) = uint16(v330)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+630)) = uint16(v319)
	v5345 = base.I32_rem_u_s(v4432, int32(3))
	v5346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5267)+4)))
	v5349 = F_SpGistGetBuffer(m, l0, v5345, v5346+int32(4), v296)
	mBase = m.M
	v5350 = m.ExcPending
	if v5350 != 0 {
		goto L3
	} else {
		goto L780
	}
L780:
	;
	if v5349 < int32(0) {
		goto L782
	} else {
		goto L783
	}
L781:
	;
	if v5349 < int32(0) {
		goto L786
	} else {
		goto L787
	}
L782:
	;
	v5354 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[1]))
	v5360 = *(*int32)(unsafe.Add(mBase, uint32(v5354+(v5349^int32(-1))<<(uint(int32(6))%32))+16))
	v5369 = v5360
	goto L781
L783:
	;
	goto L784
L784:
	;
	v5362 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[2]))
	v5368 = *(*int32)(unsafe.Add(mBase, uint32(v5362+v5349<<(uint(int32(6))%32)+int32(-64))+16))
	v5369 = v5368
	goto L781
L785:
	;
	if v5369 == v4432 {
		goto L647
	} else {
		goto L789
	}
L786:
	;
	v5373 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v5379 = *(*int32)(unsafe.Add(mBase, uint32(v5373+(v5349^int32(-1))<<(uint(int32(2))%32))))
	v5387 = v5379
	goto L785
L787:
	;
	goto L788
L788:
	;
	v5381 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v5387 = v5381 + v5349<<(uint(int32(13))%32) + int32(-8192)
	goto L785
L789:
	;
	v5389 = int32(_a_F_spgdoinsert_4)
	v5391 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	v5392 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v5391 + v5392
	if v5349 == v313 {
		goto L790
	} else {
		goto L791
	}
L790:
	;
	v5399 = v5392
	goto L792
L791:
	;
	v5399 = int32(2)
	goto L792
L792:
	;
	v5400 = base.B2i32(v4423 == v313)
	if v4423 == v313 {
		goto L793
	} else {
		goto L794
	}
L793:
	;
	v5401 = int32(0)
	goto L795
L794:
	;
	v5401 = v5399
	goto L795
L795:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+629)) = uint8(v5401)
	v5403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5267)+4)))
	v5405 = F_SpGistPageAddNewItem(m, v5387, v5267, v5403, int32(0))
	mBase = m.M
	v5406 = m.ExcPending
	if v5406 != 0 {
		goto L3
	} else {
		goto L796
	}
L796:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+626)) = uint16(v5405)
	F_MarkBufferDirty(m, v5349)
	mBase = m.M
	v5409 = m.ExcPending
	if v5409 != 0 {
		goto L3
	} else {
		goto L797
	}
L797:
	;
	F_saveNodeLink(m, v56+int32(412), v5369, v5405)
	mBase = m.M
	v5413 = m.ExcPending
	if v5413 != 0 {
		goto L3
	} else {
		goto L798
	}
L798:
	;
	v5414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v5414 == int32(1) {
		goto L800
	} else {
		goto L801
	}
L799:
	;
	F_PageIndexTupleDelete(m, v4424, v4474)
	mBase = m.M
	v5475 = m.ExcPending
	if v5475 != 0 {
		goto L3
	} else {
		goto L811
	}
L800:
	;
	v5421 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v5421))) = int32(67)
	v5427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5421)+4)))
	v5429 = v5427 & int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v5421)+4)) = uint16(v5429)
	goto L805
L801:
	;
	goto L802
L802:
	;
	v5448 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v5448))) = int32(65)
	v5454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5448)+4)))
	v5456 = v5454 & int32(_a_F_spgdoinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v5448)+4)) = uint16(v5456)
	goto L808
L803:
	;
	v5473 = v5421
	goto L799
L805:
	;
	goto L806
L806:
	;
	v5440 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5421)+10)) = uint16(v5440)
	*(*int32)(unsafe.Add(mBase, uint32(v5421)+6)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v5421)+12)) = v5440
	goto L803
L807:
	;
	v5473 = v5448
	goto L799
L808:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v5448)+10)) = uint16(v5405)
	*(*uint16)(unsafe.Add(mBase, uint32(v5448)+8)) = uint16(v5369)
	v5463 = int32(base.Ui32(v5369) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v5448)+6)) = uint16(v5463)
	v5465 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v5448)+12)) = v5465
	goto L807
L811:
	;
	v5476 = *(*int32)(unsafe.Add(mBase, uint32(v5473)))
	v5480 = F_PageAddItemExtended(m, v4424, v5473, int32(base.Ui32(v5476)>>(uint(int32(2))%32)), v4474, int32(0))
	mBase = m.M
	v5481 = m.ExcPending
	if v5481 != 0 {
		goto L3
	} else {
		goto L812
	}
L812:
	;
	if v5480 != v4474 {
		goto L646
	} else {
		goto L813
	}
L813:
	;
	v5483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4424)+16)))
	v5487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v5487 != 0 {
		goto L814
	} else {
		goto L815
	}
L814:
	;
	v5488 = int32(4)
	goto L816
L815:
	;
	v5488 = int32(2)
	goto L816
L816:
	;
	v5489 = v4424 + v5483 + v5488
	v5490 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5489))))
	v5492 = v5490 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5489))) = uint16(v5492)
	F_MarkBufferDirty(m, v4423)
	mBase = m.M
	v5495 = m.ExcPending
	if v5495 != 0 {
		goto L3
	} else {
		goto L817
	}
L817:
	;
	v5496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5496)+118)))
	if v5497 != int32(112) {
		goto L818
	} else {
		goto L819
	}
L818:
	;
	v5553 = int32(_a_F_spgdoinsert_4)
	v5555 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v5555 - int32(1)
	if base.B2i32(v5349 == v4423) == int32(0) {
		goto L839
	} else {
		goto L840
	}
L819:
	;
	v5501 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[6]))
	if v5501 <= int32(0) {
		goto L820
	} else {
		goto L821
	}
L820:
	;
	v5504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v5504 != 0 {
		goto L818
	} else {
		goto L823
	}
L821:
	;
	goto L822
L822:
	;
	v5506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v5506 != 0 {
		goto L818
	} else {
		goto L825
	}
L823:
	;
	v5505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v5505 != 0 {
		goto L818
	} else {
		goto L824
	}
L824:
	;
	goto L822
L825:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v5508 = m.ExcPending
	if v5508 != 0 {
		goto L3
	} else {
		goto L826
	}
L826:
	;
	F_XLogRegisterBuffer(m, int32(0), v4423, int32(8))
	mBase = m.M
	v5512 = m.ExcPending
	if v5512 != 0 {
		goto L3
	} else {
		goto L827
	}
L827:
	;
	v5516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+628)))
	if v5516 != 0 {
		goto L828
	} else {
		goto L829
	}
L828:
	;
	v5517 = int32(14)
	goto L830
L829:
	;
	v5517 = int32(8)
	goto L830
L830:
	;
	F_XLogRegisterBuffer(m, int32(1), v5349, v5517)
	mBase = m.M
	v5519 = m.ExcPending
	if v5519 != 0 {
		goto L3
	} else {
		goto L831
	}
L831:
	;
	v5520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+629)))
	if v5520 == int32(2) {
		goto L832
	} else {
		goto L833
	}
L832:
	;
	F_XLogRegisterBuffer(m, int32(2), v313, int32(8))
	mBase = m.M
	v5526 = m.ExcPending
	if v5526 != 0 {
		goto L3
	} else {
		goto L835
	}
L833:
	;
	goto L834
L834:
	;
	F_XLogRegisterData(m, v56+int32(624), int32(20))
	mBase = m.M
	v5531 = m.ExcPending
	if v5531 != 0 {
		goto L3
	} else {
		goto L836
	}
L835:
	;
	goto L834
L836:
	;
	v5532 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5267)+4)))
	F_XLogRegisterData(m, v5267, v5532)
	mBase = m.M
	v5534 = m.ExcPending
	if v5534 != 0 {
		goto L3
	} else {
		goto L837
	}
L837:
	;
	v5537 = F_XLogInsert(m, int32(16), int32(48))
	mBase = m.M
	v5538 = m.ExcPending
	if v5538 != 0 {
		goto L3
	} else {
		goto L838
	}
L838:
	;
	v5539 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v5387))) = base.I64_rotr(v5537, v5539)
	v5542 = base.I32_wrap_i64(v5537)
	*(*int32)(unsafe.Add(mBase, uint32(v314)+4)) = v5542
	v5546 = base.I32_wrap_i64(int64(base.Ui64(v5537) >> (uint(v5539) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v314))) = v5546
	*(*int32)(unsafe.Add(mBase, uint32(v4424)+4)) = v5542
	*(*int32)(unsafe.Add(mBase, uint32(v4424))) = v5546
	goto L818
L839:
	;
	if v5400 == int32(0) {
		goto L842
	} else {
		goto L843
	}
L840:
	;
	v5568 = v4423
	goto L841
L841:
	;
	v5572 = v5568
	v5573 = v5387
	v5576 = v5405
	v5577 = v5369
	goto L754
L842:
	;
	F_SpGistSetLastUsedPage(m, l0, v4423)
	mBase = m.M
	v5565 = m.ExcPending
	if v5565 != 0 {
		goto L3
	} else {
		goto L845
	}
L843:
	;
	goto L844
L844:
	;
	v5568 = v5349
	goto L841
L845:
	;
	F_UnlockReleaseBuffer(m, v4423)
	mBase = m.M
	v5567 = m.ExcPending
	if v5567 != 0 {
		goto L3
	} else {
		goto L846
	}
L846:
	;
	goto L844
L847:
	;
	v6776 = v5579
	v6779 = v5572
	goto L30
L848:
	;
	v5589 = *(*int32)(unsafe.Add(mBase, uint32(v56)+364))
	if base.Ui32(v5584) <= base.Ui32(v5589) {
		goto L644
	} else {
		goto L849
	}
L849:
	;
	v5593 = F_palloc(m, v5584<<(uint(int32(2))%32))
	mBase = m.M
	v5594 = m.ExcPending
	if v5594 != 0 {
		goto L3
	} else {
		goto L850
	}
L850:
	;
	v5595 = int32(0)
	v5596 = *(*int32)(unsafe.Add(mBase, uint32(v56)+356))
	if v5595 < v5596 {
		goto L851
	} else {
		goto L852
	}
L851:
	;
	v5605 = v5595
	goto L854
L852:
	;
	v5676 = v5596
	goto L853
L853:
	;
	v5722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+348)))
	v5723 = *(*int32)(unsafe.Add(mBase, uint32(v56)+352))
	v5724 = F_spgFormInnerTuple(m, l1, v5722, v5723, v5676, v5593)
	mBase = m.M
	v5725 = m.ExcPending
	if v5725 != 0 {
		goto L3
	} else {
		goto L861
	}
L854:
	;
	v5653 = v5605 << (uint(int32(2)) % 32)
	v5655 = *(*int32)(unsafe.Add(mBase, uint32(v56)+360))
	if v5655 != 0 {
		goto L856
	} else {
		goto L857
	}
L855:
	;
	v5676 = v5667
	goto L853
L856:
	;
	v5657 = *(*int32)(unsafe.Add(mBase, uint32(v5655+v5653)))
	v5659 = v5657
	goto L858
L857:
	;
	v5659 = int32(0)
	goto L858
L858:
	;
	v5662 = F_spgFormNodeTuple(m, l1, v5659, base.B2i32(v5655 == int32(0)))
	mBase = m.M
	v5663 = m.ExcPending
	if v5663 != 0 {
		goto L3
	} else {
		goto L859
	}
L859:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5593+v5653))) = v5662
	v5666 = v5605 + int32(1)
	v5667 = *(*int32)(unsafe.Add(mBase, uint32(v56)+356))
	if v5666 < v5667 {
		v5605 = v5666
		goto L854
	} else {
		goto L860
	}
L860:
	;
	goto L855
L861:
	;
	v5726 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5724)+4)))
	v5727 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+4)))
	if base.Ui32(v5727) < base.Ui32(v5726) {
		goto L643
	} else {
		goto L862
	}
L862:
	;
	v5729 = *(*int32)(unsafe.Add(mBase, uint32(v4541)))
	v5734 = F_palloc(m, int32(base.Ui32(v5729)>>(uint(int32(1))%32))&int32(_a_F_spgdoinsert_33))
	mBase = m.M
	v5735 = m.ExcPending
	if v5735 != 0 {
		goto L3
	} else {
		goto L863
	}
L863:
	;
	v5736 = *(*int32)(unsafe.Add(mBase, uint32(v4541)))
	if v5736&int32(_a_F_spgdoinsert_21) == int32(0) {
		goto L865
	} else {
		goto L866
	}
L864:
	;
	v5870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+368)))
	v5871 = *(*int32)(unsafe.Add(mBase, uint32(v56)+372))
	v5872 = F_spgFormInnerTuple(m, l1, v5870, v5871, v5830, v5734)
	mBase = m.M
	v5873 = m.ExcPending
	if v5873 != 0 {
		goto L3
	} else {
		goto L871
	}
L865:
	;
	v5830 = int32(0)
	goto L864
L866:
	;
	goto L867
L867:
	;
	v5754 = v4541 + int32(base.Ui32(v5736)>>(uint(int32(16))%32)) + int32(8)
	v5755 = int32(0)
	goto L868
L868:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5734+v5755<<(uint(int32(2))%32)))) = v5754
	v5805 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5754)+6)))
	v5806 = int32(_a_F_spgdoinsert_22)
	v5810 = v5755 + int32(1)
	v5811 = *(*int32)(unsafe.Add(mBase, uint32(v4541)))
	v5815 = int32(base.Ui32(v5811)>>(uint(int32(3))%32)) & v5806
	if base.Ui32(v5810) < base.Ui32(v5815) {
		v5754 = v5754 + v5805&v5806
		v5755 = v5810
		goto L868
	} else {
		goto L870
	}
L869:
	;
	v5830 = v5815
	goto L864
L870:
	;
	goto L869
L871:
	;
	v5874 = *(*int32)(unsafe.Add(mBase, uint32(v5872)))
	v5877 = *(*int32)(unsafe.Add(mBase, uint32(v4541)))
	*(*int32)(unsafe.Add(mBase, uint32(v5872))) = v5874&int32(-5) | v5877&int32(4)
	v5882 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+628)) = uint8(v5882)
	if base.Ui32(v4468) <= base.Ui32(int32(1)) {
		goto L874
	} else {
		goto L875
	}
L872:
	;
	v5918 = int32(_a_F_spgdoinsert_4)
	v5920 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v5920 + int32(1)
	F_PageIndexTupleDelete(m, v4424, v4474)
	mBase = m.M
	v5925 = m.ExcPending
	if v5925 != 0 {
		goto L3
	} else {
		goto L886
	}
L873:
	;
	v5914 = F_SpGistGetBuffer(m, l0, v4472, v5910+int32(4), v296)
	mBase = m.M
	v5915 = m.ExcPending
	if v5915 != 0 {
		goto L3
	} else {
		goto L885
	}
L874:
	;
	v5886 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5872)+4)))
	v5910 = v5886
	goto L873
L875:
	;
	goto L876
L876:
	;
	v5887 = int32(0)
	v5888 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4424)+14)))
	v5889 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4424)+12)))
	v5890 = v5888 - v5889
	if v5887 < v5890 {
		goto L878
	} else {
		goto L879
	}
L877:
	;
	v5897 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4424)+16)))
	v5899 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4424+v5897)+4)))
	if v5899 != 0 {
		goto L881
	} else {
		goto L882
	}
L878:
	;
	v5894 = v5890
	goto L880
L879:
	;
	v5894 = v5887
	goto L880
L880:
	;
	goto L877
L881:
	;
	v5900 = int32(20)
	goto L883
L882:
	;
	v5900 = int32(0)
	goto L883
L883:
	;
	v5902 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+4)))
	v5904 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5872)+4)))
	v5905 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5724)+4)))
	if base.Ui32(v5904+v5905+int32(4)) <= base.Ui32(v5894+v5900+v5902) {
		v5917 = v5887
		goto L872
	} else {
		goto L884
	}
L884:
	;
	v5910 = v5904
	goto L873
L885:
	;
	v5917 = v5914
	goto L872
L886:
	;
	v5926 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5724)+4)))
	v5928 = F_PageAddItemExtended(m, v4424, v5724, v5926, v4474, int32(0))
	mBase = m.M
	v5929 = m.ExcPending
	if v5929 != 0 {
		goto L3
	} else {
		goto L887
	}
L887:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+624)) = uint16(v5928)
	if v5928 != v4474 {
		goto L642
	} else {
		goto L888
	}
L888:
	;
	if v5917 == int32(0) {
		goto L890
	} else {
		goto L891
	}
L889:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+629)) = uint8(v5985)
	v5989 = *(*int32)(unsafe.Add(mBase, uint32(v56)+364))
	v5990 = *(*int32)(unsafe.Add(mBase, uint32(v5724)))
	v5994 = int32(base.Ui32(v5990)>>(uint(int32(3))%32)) & int32(_a_F_spgdoinsert_22)
	if v5994 != 0 {
		goto L905
	} else {
		goto L906
	}
L890:
	;
	v5934 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5872)+4)))
	v5936 = F_SpGistPageAddNewItem(m, v4424, v5872, v5934, int32(0))
	mBase = m.M
	v5937 = m.ExcPending
	if v5937 != 0 {
		goto L3
	} else {
		goto L893
	}
L891:
	;
	goto L892
L892:
	;
	if v5917 < int32(0) {
		goto L895
	} else {
		goto L896
	}
L893:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+626)) = uint16(v5936)
	v5985 = int32(1)
	v5986 = v4432
	v5987 = v5936
	goto L889
L894:
	;
	v5959 = int32(0)
	if v5917 < v5959 {
		goto L899
	} else {
		goto L900
	}
L895:
	;
	v5943 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[1]))
	v5949 = *(*int32)(unsafe.Add(mBase, uint32(v5943+(v5917^int32(-1))<<(uint(int32(6))%32))+16))
	v5958 = v5949
	goto L894
L896:
	;
	goto L897
L897:
	;
	v5951 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[2]))
	v5957 = *(*int32)(unsafe.Add(mBase, uint32(v5951+v5917<<(uint(int32(6))%32)+int32(-64))+16))
	v5958 = v5957
	goto L894
L898:
	;
	v5978 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5872)+4)))
	v5980 = F_SpGistPageAddNewItem(m, v5977, v5872, v5978, int32(0))
	mBase = m.M
	v5981 = m.ExcPending
	if v5981 != 0 {
		goto L3
	} else {
		goto L902
	}
L899:
	;
	v5963 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v5969 = *(*int32)(unsafe.Add(mBase, uint32(v5963+(v5917^int32(-1))<<(uint(int32(2))%32))))
	v5977 = v5969
	goto L898
L900:
	;
	goto L901
L901:
	;
	v5971 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v5977 = v5971 + v5917<<(uint(int32(13))%32) + int32(-8192)
	goto L898
L902:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+626)) = uint16(v5980)
	F_MarkBufferDirty(m, v5917)
	mBase = m.M
	v5984 = m.ExcPending
	if v5984 != 0 {
		goto L3
	} else {
		goto L903
	}
L903:
	;
	v5985 = v5959
	v5986 = v5958
	v5987 = v5980
	goto L889
L904:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v6007)+4)) = uint16(v5987)
	*(*uint16)(unsafe.Add(mBase, uint32(v6007)+2)) = uint16(v5986)
	v6133 = int32(base.Ui32(v5986) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v6007))) = uint16(v6133)
	v6135 = *(*int32)(unsafe.Add(mBase, uint32(v56)+364))
	v6136 = *(*int32)(unsafe.Add(mBase, uint32(v4479)))
	v6139 = v4424 + v6136&int32(_a_F_spgdoinsert_6)
	v6140 = *(*int32)(unsafe.Add(mBase, uint32(v6139)))
	v6144 = int32(base.Ui32(v6140)>>(uint(int32(3))%32)) & int32(_a_F_spgdoinsert_22)
	if v6144 != 0 {
		goto L916
	} else {
		goto L917
	}
L905:
	;
	v6007 = v5724 + int32(base.Ui32(v5990)>>(uint(int32(16))%32)) + int32(8)
	v6008 = int32(0)
	goto L908
L906:
	;
	goto L907
L907:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6118 = m.ExcPending
	if v6118 != 0 {
		goto L3
	} else {
		goto L912
	}
L908:
	;
	if v6008 == v5989 {
		goto L904
	} else {
		goto L910
	}
L909:
	;
	goto L907
L910:
	;
	v6055 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6007)+6)))
	v6060 = v6008 + int32(1)
	if v6060 != v5994 {
		v6007 = v6007 + v6055&int32(_a_F_spgdoinsert_22)
		v6008 = v6060
		goto L908
	} else {
		goto L911
	}
L911:
	;
	goto L909
L912:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+112)) = v5989
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_34), v56+int32(112))
	mBase = m.M
	v6124 = m.ExcPending
	if v6124 != 0 {
		goto L3
	} else {
		goto L913
	}
L913:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(68), int32(_a_F_spgdoinsert_35))
	mBase = m.M
	v6129 = m.ExcPending
	if v6129 != 0 {
		goto L3
	} else {
		goto L914
	}
L914:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L915:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v6157)+4)) = uint16(v5987)
	*(*uint16)(unsafe.Add(mBase, uint32(v6157)+2)) = uint16(v5986)
	*(*uint16)(unsafe.Add(mBase, uint32(v6157))) = uint16(v6133)
	F_MarkBufferDirty(m, v4423)
	mBase = m.M
	v6284 = m.ExcPending
	if v6284 != 0 {
		goto L3
	} else {
		goto L926
	}
L916:
	;
	v6157 = v6139 + int32(base.Ui32(v6140)>>(uint(int32(16))%32)) + int32(8)
	v6158 = int32(0)
	goto L919
L917:
	;
	goto L918
L918:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6268 = m.ExcPending
	if v6268 != 0 {
		goto L3
	} else {
		goto L923
	}
L919:
	;
	if v6158 == v6135 {
		goto L915
	} else {
		goto L921
	}
L920:
	;
	goto L918
L921:
	;
	v6205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6157)+6)))
	v6210 = v6158 + int32(1)
	if v6210 != v6144 {
		v6157 = v6157 + v6205&int32(_a_F_spgdoinsert_22)
		v6158 = v6210
		goto L919
	} else {
		goto L922
	}
L922:
	;
	goto L920
L923:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+128)) = v6135
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_34), v56+int32(128))
	mBase = m.M
	v6274 = m.ExcPending
	if v6274 != 0 {
		goto L3
	} else {
		goto L924
	}
L924:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(68), int32(_a_F_spgdoinsert_35))
	mBase = m.M
	v6279 = m.ExcPending
	if v6279 != 0 {
		goto L3
	} else {
		goto L925
	}
L925:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L926:
	;
	v6285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v6286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6285)+118)))
	if v6286 != int32(112) {
		goto L929
	} else {
		goto L930
	}
L927:
	;
	v6390 = int32(0)
	v6392 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v6392 == v6390 {
		goto L654
	} else {
		goto L958
	}
L928:
	;
	F_SpGistSetLastUsedPage(m, l0, v5917)
	mBase = m.M
	v6383 = m.ExcPending
	if v6383 != 0 {
		goto L3
	} else {
		goto L956
	}
L929:
	;
	v6370 = int32(_a_F_spgdoinsert_4)
	v6372 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v6372 - int32(1)
	if v5917 == int32(0) {
		goto L927
	} else {
		goto L955
	}
L930:
	;
	v6290 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[6]))
	if v6290 <= int32(0) {
		goto L931
	} else {
		goto L932
	}
L931:
	;
	v6293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v6293 != 0 {
		goto L929
	} else {
		goto L934
	}
L932:
	;
	goto L933
L933:
	;
	v6295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	if v6295 != 0 {
		goto L929
	} else {
		goto L936
	}
L934:
	;
	v6294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v6294 != 0 {
		goto L929
	} else {
		goto L935
	}
L935:
	;
	goto L933
L936:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v6297 = m.ExcPending
	if v6297 != 0 {
		goto L3
	} else {
		goto L937
	}
L937:
	;
	F_XLogRegisterData(m, v56+int32(624), int32(6))
	mBase = m.M
	v6302 = m.ExcPending
	if v6302 != 0 {
		goto L3
	} else {
		goto L938
	}
L938:
	;
	v6303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6139)+4)))
	F_XLogRegisterData(m, v6139, v6303)
	mBase = m.M
	v6305 = m.ExcPending
	if v6305 != 0 {
		goto L3
	} else {
		goto L939
	}
L939:
	;
	v6306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5872)+4)))
	F_XLogRegisterData(m, v5872, v6306)
	mBase = m.M
	v6308 = m.ExcPending
	if v6308 != 0 {
		goto L3
	} else {
		goto L940
	}
L940:
	;
	F_XLogRegisterBuffer(m, int32(0), v4423, int32(8))
	mBase = m.M
	v6312 = m.ExcPending
	if v6312 != 0 {
		goto L3
	} else {
		goto L941
	}
L941:
	;
	if v5917 != 0 {
		goto L942
	} else {
		goto L943
	}
L942:
	;
	v6316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+628)))
	if v6316 != 0 {
		goto L945
	} else {
		goto L946
	}
L943:
	;
	goto L944
L944:
	;
	v6359 = F_XLogInsert(m, int32(16), int32(64))
	mBase = m.M
	v6360 = m.ExcPending
	if v6360 != 0 {
		goto L3
	} else {
		goto L954
	}
L945:
	;
	v6317 = int32(14)
	goto L947
L946:
	;
	v6317 = int32(8)
	goto L947
L947:
	;
	F_XLogRegisterBuffer(m, int32(1), v5917, v6317)
	mBase = m.M
	v6319 = m.ExcPending
	if v6319 != 0 {
		goto L3
	} else {
		goto L948
	}
L948:
	;
	v6322 = F_XLogInsert(m, int32(16), int32(64))
	mBase = m.M
	v6323 = m.ExcPending
	if v6323 != 0 {
		goto L3
	} else {
		goto L949
	}
L949:
	;
	v6324 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v4424))) = base.I64_rotr(v6322, v6324)
	if v5917 < int32(0) {
		goto L951
	} else {
		goto L952
	}
L950:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6348)+4)) = base.I32_wrap_i64(v6322)
	*(*int32)(unsafe.Add(mBase, uint32(v6348))) = base.I32_wrap_i64(int64(base.Ui64(v6322) >> (uint(v6324) % 64)))
	v6351 = int32(_a_F_spgdoinsert_4)
	v6353 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v6353 - int32(1)
	goto L928
L951:
	;
	v6334 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[3]))
	v6340 = *(*int32)(unsafe.Add(mBase, uint32(v6334+(v5917^int32(-1))<<(uint(int32(2))%32))))
	v6348 = v6340
	goto L950
L952:
	;
	goto L953
L953:
	;
	v6342 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[4]))
	v6348 = v6342 + v5917<<(uint(int32(13))%32) + int32(-8192)
	goto L950
L954:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4424))) = base.I64_rotr(v6359, int64(32))
	v6364 = int32(_a_F_spgdoinsert_4)
	v6366 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[5])) = v6366 - int32(1)
	goto L927
L955:
	;
	goto L928
L956:
	;
	F_UnlockReleaseBuffer(m, v5917)
	mBase = m.M
	v6385 = m.ExcPending
	if v6385 != 0 {
		goto L3
	} else {
		goto L957
	}
L957:
	;
	goto L927
L958:
	;
	v6776 = v6390
	v6779 = v4423
	goto L30
L959:
	;
	v6399 = *(*int32)(unsafe.Add(mBase, uint32(v56)+344))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+16)) = v6399
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_36), v56+int32(16))
	mBase = m.M
	v6405 = m.ExcPending
	if v6405 != 0 {
		goto L3
	} else {
		goto L960
	}
L960:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(2318), int32(_a_F_spgdoinsert_19))
	mBase = m.M
	v6410 = m.ExcPending
	if v6410 != 0 {
		goto L3
	} else {
		goto L961
	}
L961:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L962:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_37), int32(0))
	mBase = m.M
	v6418 = m.ExcPending
	if v6418 != 0 {
		goto L3
	} else {
		goto L963
	}
L963:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(2295), int32(_a_F_spgdoinsert_19))
	mBase = m.M
	v6423 = m.ExcPending
	if v6423 != 0 {
		goto L3
	} else {
		goto L964
	}
L964:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L965:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_38), int32(0))
	mBase = m.M
	v6431 = m.ExcPending
	if v6431 != 0 {
		goto L3
	} else {
		goto L966
	}
L966:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(90), int32(_a_F_spgdoinsert_39))
	mBase = m.M
	v6436 = m.ExcPending
	if v6436 != 0 {
		goto L3
	} else {
		goto L967
	}
L967:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L968:
	;
	v6441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5267)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+80)) = v6441
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_8), v56+int32(80))
	mBase = m.M
	v6447 = m.ExcPending
	if v6447 != 0 {
		goto L3
	} else {
		goto L969
	}
L969:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1553), int32(_a_F_spgdoinsert_40))
	mBase = m.M
	v6452 = m.ExcPending
	if v6452 != 0 {
		goto L3
	} else {
		goto L970
	}
L970:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L971:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_41), int32(0))
	mBase = m.M
	v6460 = m.ExcPending
	if v6460 != 0 {
		goto L3
	} else {
		goto L972
	}
L972:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1588), int32(_a_F_spgdoinsert_40))
	mBase = m.M
	v6465 = m.ExcPending
	if v6465 != 0 {
		goto L3
	} else {
		goto L973
	}
L973:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L974:
	;
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_42), int32(0))
	mBase = m.M
	v6473 = m.ExcPending
	if v6473 != 0 {
		goto L3
	} else {
		goto L975
	}
L975:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1616), int32(_a_F_spgdoinsert_40))
	mBase = m.M
	v6478 = m.ExcPending
	if v6478 != 0 {
		goto L3
	} else {
		goto L976
	}
L976:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L977:
	;
	v6483 = *(*int32)(unsafe.Add(mBase, uint32(v5473)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+64)) = int32(base.Ui32(v6483) >> (uint(int32(2)) % 32))
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_8), v56-int32(-64))
	mBase = m.M
	v6491 = m.ExcPending
	if v6491 != 0 {
		goto L3
	} else {
		goto L978
	}
L978:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1661), int32(_a_F_spgdoinsert_40))
	mBase = m.M
	v6496 = m.ExcPending
	if v6496 != 0 {
		goto L3
	} else {
		goto L979
	}
L979:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L980:
	;
	v6501 = *(*int32)(unsafe.Add(mBase, uint32(v56)+356))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+96)) = v6501
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_43), v56+int32(96))
	mBase = m.M
	v6507 = m.ExcPending
	if v6507 != 0 {
		goto L3
	} else {
		goto L981
	}
L981:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1736), int32(_a_F_spgdoinsert_44))
	mBase = m.M
	v6512 = m.ExcPending
	if v6512 != 0 {
		goto L3
	} else {
		goto L982
	}
L982:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L983:
	;
	v6517 = *(*int32)(unsafe.Add(mBase, uint32(v56)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+160)) = v6517
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_45), v56+int32(160))
	mBase = m.M
	v6523 = m.ExcPending
	if v6523 != 0 {
		goto L3
	} else {
		goto L984
	}
L984:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1741), int32(_a_F_spgdoinsert_44))
	mBase = m.M
	v6528 = m.ExcPending
	if v6528 != 0 {
		goto L3
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
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_46), int32(0))
	mBase = m.M
	v6536 = m.ExcPending
	if v6536 != 0 {
		goto L3
	} else {
		goto L987
	}
L987:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1769), int32(_a_F_spgdoinsert_44))
	mBase = m.M
	v6541 = m.ExcPending
	if v6541 != 0 {
		goto L3
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
	v6546 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5724)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+144)) = v6546
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_8), v56+int32(144))
	mBase = m.M
	v6552 = m.ExcPending
	if v6552 != 0 {
		goto L3
	} else {
		goto L990
	}
L990:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1825), int32(_a_F_spgdoinsert_44))
	mBase = m.M
	v6557 = m.ExcPending
	if v6557 != 0 {
		goto L3
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
	goto L63
L993:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v6628 = m.ExcPending
	if v6628 != 0 {
		goto L3
	} else {
		goto L994
	}
L994:
	;
	v6629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = int32(_a_F_spgdoinsert_30)
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v6629 + int32(4)
	F_errmsg(m, int32(_a_F_spgdoinsert_31), v56)
	mBase = m.M
	v6638 = m.ExcPending
	if v6638 != 0 {
		goto L3
	} else {
		goto L995
	}
L995:
	;
	F_errhint(m, int32(_a_F_spgdoinsert_32), int32(0))
	mBase = m.M
	v6642 = m.ExcPending
	if v6642 != 0 {
		goto L3
	} else {
		goto L996
	}
L996:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(2005), int32(_a_F_spgdoinsert_19))
	mBase = m.M
	v6647 = m.ExcPending
	if v6647 != 0 {
		goto L3
	} else {
		goto L997
	}
L997:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L998:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+48)) = v4680
	F_errmsg_internal(m, int32(_a_F_spgdoinsert_34), v56+int32(48))
	mBase = m.M
	v6710 = m.ExcPending
	if v6710 != 0 {
		goto L3
	} else {
		goto L999
	}
L999:
	;
	F_errfinish(m, int32(_a_F_spgdoinsert_9), int32(1490), int32(_a_F_spgdoinsert_47))
	mBase = m.M
	v6715 = m.ExcPending
	if v6715 != 0 {
		goto L3
	} else {
		goto L1000
	}
L1000:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1001:
	;
	v6836 = v6776
	v6839 = int32(0)
	v6846 = v313
	goto L29
L1002:
	;
	goto L1003
L1003:
	;
	F_SpGistSetLastUsedPage(m, l0, v6779)
	mBase = m.M
	v6827 = m.ExcPending
	if v6827 != 0 {
		goto L3
	} else {
		goto L1004
	}
L1004:
	;
	F_UnlockReleaseBuffer(m, v6779)
	mBase = m.M
	v6829 = m.ExcPending
	if v6829 != 0 {
		goto L3
	} else {
		goto L1005
	}
L1005:
	;
	v6836 = v6776
	v6839 = v6779
	v6846 = v313
	goto L29
L1006:
	;
	F_SpGistSetLastUsedPage(m, l0, v6846)
	mBase = m.M
	v6890 = m.ExcPending
	if v6890 != 0 {
		goto L3
	} else {
		goto L1009
	}
L1007:
	;
	goto L1008
L1008:
	;
	v6894 = *(*int32)(unsafe.Add(mBase, _c_F_spgdoinsert[0]))
	if v6894 == int32(0) {
		v6905 = v6836
		goto L28
	} else {
		goto L1011
	}
L1009:
	;
	F_UnlockReleaseBuffer(m, v6846)
	mBase = m.M
	v6892 = m.ExcPending
	if v6892 != 0 {
		goto L3
	} else {
		goto L1010
	}
L1010:
	;
	goto L1008
L1011:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v6898 = m.ExcPending
	if v6898 != 0 {
		goto L3
	} else {
		goto L1012
	}
L1012:
	;
	v6905 = v6836
	goto L28
L1013:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
