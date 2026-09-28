package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_PostgresMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int64
	_ = v233
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int64
	_ = v286
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v468 int32
	_ = v468
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v510 int32
	_ = v510
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int64
	_ = v619
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int64
	_ = v672
	var v685 int32
	_ = v685
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v728 int32
	_ = v728
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v812 int32
	_ = v812
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v854 int32
	_ = v854
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v889 int32
	_ = v889
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v966 int32
	_ = v966
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1170 int64
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int64
	_ = v1178
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1266 int64
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1279 int64
	_ = v1279
	var v1290 int32
	_ = v1290
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1363 int32
	_ = v1363
	var v1367 int32
	_ = v1367
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1550 int32
	_ = v1550
	var v1554 int32
	_ = v1554
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1566 int32
	_ = v1566
	var v1572 int32
	_ = v1572
	var v1576 int32
	_ = v1576
	var v1608 int32
	_ = v1608
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1652 int32
	_ = v1652
	var v1654 int32
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1698 int32
	_ = v1698
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1757 int32
	_ = v1757
	var v1762 int32
	_ = v1762
	var v1793 int32
	_ = v1793
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1854 int64
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1887 int32
	_ = v1887
	var v1895 int32
	_ = v1895
	var v1936 int32
	_ = v1936
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1988 int32
	_ = v1988
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2048 int32
	_ = v2048
	var v2195 int32
	_ = v2195
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2212 int32
	_ = v2212
	var v2217 int32
	_ = v2217
	var v2219 int32
	_ = v2219
	var v2223 int32
	_ = v2223
	var v2225 int32
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2231 int32
	_ = v2231
	var v2233 int32
	_ = v2233
	var v2236 int32
	_ = v2236
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2244 int32
	_ = v2244
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2272 int32
	_ = v2272
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2285 int32
	_ = v2285
	var v2287 int32
	_ = v2287
	var v2289 int32
	_ = v2289
	var v2292 int32
	_ = v2292
	var v2294 int32
	_ = v2294
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2303 int32
	_ = v2303
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2342 int32
	_ = v2342
	var v2346 int32
	_ = v2346
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2399 int32
	_ = v2399
	var v2401 int32
	_ = v2401
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2408 int32
	_ = v2408
	var v2411 int32
	_ = v2411
	var v2413 int32
	_ = v2413
	var v2415 int32
	_ = v2415
	var v2418 int32
	_ = v2418
	var v2421 int32
	_ = v2421
	var v2424 int32
	_ = v2424
	var v2426 int32
	_ = v2426
	var v2431 int32
	_ = v2431
	var v2435 int32
	_ = v2435
	var v2440 int32
	_ = v2440
	var v2443 int32
	_ = v2443
	var v2447 int32
	_ = v2447
	var v2452 int32
	_ = v2452
	var v2496 int32
	_ = v2496
	var v2500 int32
	_ = v2500
	var v2541 int32
	_ = v2541
	var v2546 int32
	_ = v2546
	var v2548 int32
	_ = v2548
	var v2552 int32
	_ = v2552
	var v2558 int32
	_ = v2558
	var v2562 int32
	_ = v2562
	var v2564 int32
	_ = v2564
	var v2568 int32
	_ = v2568
	var v2570 int32
	_ = v2570
	var v2573 int32
	_ = v2573
	var v2578 int32
	_ = v2578
	var v2583 int32
	_ = v2583
	var v2585 int32
	_ = v2585
	var v2589 int32
	_ = v2589
	var v2592 int32
	_ = v2592
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2601 int32
	_ = v2601
	var v2604 int32
	_ = v2604
	var v2621 int32
	_ = v2621
	var v2625 int32
	_ = v2625
	var v2626 int32
	_ = v2626
	var v2634 int32
	_ = v2634
	var v2637 int32
	_ = v2637
	var v2641 int32
	_ = v2641
	var v2644 int32
	_ = v2644
	var v2646 int32
	_ = v2646
	var v2650 int32
	_ = v2650
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2657 int32
	_ = v2657
	var v2660 int32
	_ = v2660
	var v2664 int32
	_ = v2664
	var v2667 int32
	_ = v2667
	var v2669 int32
	_ = v2669
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2681 int32
	_ = v2681
	var v2685 int32
	_ = v2685
	var v2690 int32
	_ = v2690
	var v2695 int32
	_ = v2695
	var v2697 int32
	_ = v2697
	var v2700 int32
	_ = v2700
	var v2704 int32
	_ = v2704
	var v2708 int32
	_ = v2708
	var v2712 int32
	_ = v2712
	var v2716 int32
	_ = v2716
	var v2721 int32
	_ = v2721
	var v2726 int32
	_ = v2726
	var v2727 int32
	_ = v2727
	var v2729 int32
	_ = v2729
	var v2731 int32
	_ = v2731
	var v2733 int32
	_ = v2733
	var v2736 int32
	_ = v2736
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2745 int32
	_ = v2745
	var v2752 int32
	_ = v2752
	var v2787 int32
	_ = v2787
	var v2791 int32
	_ = v2791
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2840 int64
	_ = v2840
	var v2844 int32
	_ = v2844
	var v2850 int32
	_ = v2850
	var v2857 int32
	_ = v2857
	var v2858 int32
	_ = v2858
	var v2859 int32
	_ = v2859
	var v2862 int64
	_ = v2862
	var v2863 int64
	_ = v2863
	var v2871 int64
	_ = v2871
	var v2874 int64
	_ = v2874
	var v2876 int64
	_ = v2876
	var v2878 int64
	_ = v2878
	var v2880 int64
	_ = v2880
	var v2882 int64
	_ = v2882
	var v2885 int32
	_ = v2885
	var v2886 int32
	_ = v2886
	var v2892 int64
	_ = v2892
	var v2900 int64
	_ = v2900
	var v2908 int64
	_ = v2908
	var v2917 int32
	_ = v2917
	var v2922 int32
	_ = v2922
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2933 int32
	_ = v2933
	var v2935 int32
	_ = v2935
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2944 int32
	_ = v2944
	var v2947 int32
	_ = v2947
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
	var v2971 int32
	_ = v2971
	var v2972 int32
	_ = v2972
	var v2978 int32
	_ = v2978
	var v2979 int32
	_ = v2979
	var v2980 int32
	_ = v2980
	var v2987 int32
	_ = v2987
	var v2989 int32
	_ = v2989
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v3000 int32
	_ = v3000
	var v3042 int32
	_ = v3042
	var v3045 int32
	_ = v3045
	var v3048 int32
	_ = v3048
	var v3050 int32
	_ = v3050
	var v3055 int32
	_ = v3055
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3063 int32
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3067 int32
	_ = v3067
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3076 int32
	_ = v3076
	var v3080 int32
	_ = v3080
	var v3085 int32
	_ = v3085
	var v3087 int32
	_ = v3087
	var v3089 int32
	_ = v3089
	var v3092 int32
	_ = v3092
	var v3093 int32
	_ = v3093
	var v3098 int32
	_ = v3098
	var v3102 int32
	_ = v3102
	var v3107 int32
	_ = v3107
	var v3109 int32
	_ = v3109
	var v3113 int32
	_ = v3113
	var v3120 int32
	_ = v3120
	var v3123 int32
	_ = v3123
	var v3126 int32
	_ = v3126
	var v3132 int32
	_ = v3132
	var v3135 int32
	_ = v3135
	var v3139 int32
	_ = v3139
	var v3144 int32
	_ = v3144
	var v3146 int32
	_ = v3146
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3151 int32
	_ = v3151
	var v3153 int32
	_ = v3153
	var v3155 int32
	_ = v3155
	var v3162 int32
	_ = v3162
	var v3164 int32
	_ = v3164
	var v3165 int32
	_ = v3165
	var v3166 int32
	_ = v3166
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3170 int32
	_ = v3170
	var v3177 int32
	_ = v3177
	var v3218 int32
	_ = v3218
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3224 int32
	_ = v3224
	var v3226 int32
	_ = v3226
	var v3228 int32
	_ = v3228
	var v3230 int32
	_ = v3230
	var v3232 int32
	_ = v3232
	var v3234 int32
	_ = v3234
	var v3236 int32
	_ = v3236
	var v3241 int32
	_ = v3241
	var v3243 int32
	_ = v3243
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3251 int32
	_ = v3251
	var v3252 int32
	_ = v3252
	var v3255 int32
	_ = v3255
	var v3258 int32
	_ = v3258
	var v3259 int32
	_ = v3259
	var v3262 int32
	_ = v3262
	var v3266 int32
	_ = v3266
	var v3268 int32
	_ = v3268
	var v3270 int32
	_ = v3270
	var v3273 int32
	_ = v3273
	var v3276 int32
	_ = v3276
	var v3280 int32
	_ = v3280
	var v3284 int32
	_ = v3284
	var v3288 int32
	_ = v3288
	var v3296 int32
	_ = v3296
	var v3303 int32
	_ = v3303
	var v3305 int32
	_ = v3305
	var v3310 int32
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3314 int32
	_ = v3314
	var v3319 int32
	_ = v3319
	var v3327 int32
	_ = v3327
	var v3330 int32
	_ = v3330
	var v3334 int32
	_ = v3334
	var v3338 int32
	_ = v3338
	var v3341 int32
	_ = v3341
	var v3343 int32
	_ = v3343
	var v3350 int32
	_ = v3350
	var v3357 int32
	_ = v3357
	var v3359 int32
	_ = v3359
	var v3360 int32
	_ = v3360
	var v3366 int32
	_ = v3366
	var v3367 int32
	_ = v3367
	var v3368 int32
	_ = v3368
	var v3374 int32
	_ = v3374
	var v3410 int32
	_ = v3410
	var v3415 int32
	_ = v3415
	var v3417 int32
	_ = v3417
	var v3420 int32
	_ = v3420
	var v3425 int32
	_ = v3425
	var v3427 int32
	_ = v3427
	var v3430 int32
	_ = v3430
	var v3432 int32
	_ = v3432
	var v3434 int32
	_ = v3434
	var v3437 int32
	_ = v3437
	var v3443 int32
	_ = v3443
	var v3447 int32
	_ = v3447
	var v3453 int32
	_ = v3453
	var v3457 int64
	_ = v3457
	var v3460 int32
	_ = v3460
	var v3461 int32
	_ = v3461
	var v3462 int32
	_ = v3462
	var v3464 int32
	_ = v3464
	var v3466 int32
	_ = v3466
	var v3469 int32
	_ = v3469
	var v3479 int32
	_ = v3479
	var v3481 int32
	_ = v3481
	var v3484 int32
	_ = v3484
	var v3486 int32
	_ = v3486
	var v3488 int32
	_ = v3488
	var v3491 int32
	_ = v3491
	var v3496 int32
	_ = v3496
	var v3501 int32
	_ = v3501
	var v3504 int32
	_ = v3504
	var v3508 int32
	_ = v3508
	var v3509 int32
	_ = v3509
	var v3510 int32
	_ = v3510
	var v3514 int32
	_ = v3514
	var v3516 int32
	_ = v3516
	var v3517 int32
	_ = v3517
	var v3523 int32
	_ = v3523
	var v3525 int32
	_ = v3525
	var v3532 int32
	_ = v3532
	var v3536 int32
	_ = v3536
	var v3541 int32
	_ = v3541
	var v3543 int32
	_ = v3543
	var v3545 int32
	_ = v3545
	var v3547 int32
	_ = v3547
	var v3552 int32
	_ = v3552
	var v3557 int32
	_ = v3557
	var v3558 int32
	_ = v3558
	var v3561 int32
	_ = v3561
	var v3563 int32
	_ = v3563
	var v3564 int32
	_ = v3564
	var v3568 int32
	_ = v3568
	var v3570 int32
	_ = v3570
	var v3571 int32
	_ = v3571
	var v3574 int32
	_ = v3574
	var v3575 int32
	_ = v3575
	var v3580 int32
	_ = v3580
	var v3585 int32
	_ = v3585
	var v3589 int32
	_ = v3589
	var v3594 int32
	_ = v3594
	var v3595 int32
	_ = v3595
	var v3598 int32
	_ = v3598
	var v3601 int64
	_ = v3601
	var v3613 int32
	_ = v3613
	var v3615 int32
	_ = v3615
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3619 int32
	_ = v3619
	var v3623 int32
	_ = v3623
	var v3639 int32
	_ = v3639
	var v3653 int32
	_ = v3653
	var v3669 int32
	_ = v3669
	var v3672 int32
	_ = v3672
	var v3675 int32
	_ = v3675
	var v3678 int32
	_ = v3678
	var v3681 int32
	_ = v3681
	var v3684 int32
	_ = v3684
	var v3687 int32
	_ = v3687
	var v3689 int32
	_ = v3689
	var v3690 int32
	_ = v3690
	var v3692 int32
	_ = v3692
	var v3707 int32
	_ = v3707
	var v3746 int32
	_ = v3746
	var v3747 int32
	_ = v3747
	var v3776 int32
	_ = v3776
	var v3778 int32
	_ = v3778
	var v3781 int32
	_ = v3781
	var v3823 int32
	_ = v3823
	var v3829 int32
	_ = v3829
	var v3831 int32
	_ = v3831
	var v3835 int32
	_ = v3835
	var v3837 int32
	_ = v3837
	var v3838 int32
	_ = v3838
	var v3841 int32
	_ = v3841
	var v3854 int32
	_ = v3854
	var v3855 int32
	_ = v3855
	var v3856 int32
	_ = v3856
	var v3860 int32
	_ = v3860
	var v3862 int32
	_ = v3862
	var v3863 int32
	_ = v3863
	var v3865 int32
	_ = v3865
	var v3866 int32
	_ = v3866
	var v3867 int32
	_ = v3867
	var v3870 int32
	_ = v3870
	var v3871 int32
	_ = v3871
	var v3873 int32
	_ = v3873
	var v3874 int32
	_ = v3874
	var v3878 int32
	_ = v3878
	var v3879 int32
	_ = v3879
	var v3881 int32
	_ = v3881
	var v3882 int32
	_ = v3882
	var v3883 int32
	_ = v3883
	var v3884 int32
	_ = v3884
	var v3885 int32
	_ = v3885
	var v3889 int32
	_ = v3889
	var v3890 int32
	_ = v3890
	var v3893 int32
	_ = v3893
	var v3894 int32
	_ = v3894
	var v3895 int32
	_ = v3895
	var v3897 int32
	_ = v3897
	var v3898 int32
	_ = v3898
	var v3901 int32
	_ = v3901
	var v3902 int32
	_ = v3902
	var v3904 int32
	_ = v3904
	var v3911 int32
	_ = v3911
	var v3914 int32
	_ = v3914
	var v3921 int32
	_ = v3921
	var v3924 int32
	_ = v3924
	var v3925 int32
	_ = v3925
	var v3926 int32
	_ = v3926
	var v3928 int32
	_ = v3928
	var v3932 int32
	_ = v3932
	var v3933 int32
	_ = v3933
	var v3941 int32
	_ = v3941
	var v3944 int32
	_ = v3944
	var v3948 int32
	_ = v3948
	var v3952 int32
	_ = v3952
	var v3956 int32
	_ = v3956
	var v3958 int32
	_ = v3958
	var v3960 int32
	_ = v3960
	var v3964 int32
	_ = v3964
	var v3967 int32
	_ = v3967
	var v3971 int32
	_ = v3971
	var v3976 int32
	_ = v3976
	var v3979 int32
	_ = v3979
	var v3981 int32
	_ = v3981
	var v3987 int32
	_ = v3987
	var v3989 int32
	_ = v3989
	var v3994 int32
	_ = v3994
	var v4000 int32
	_ = v4000
	var v4002 int32
	_ = v4002
	var v4004 int32
	_ = v4004
	var v4016 int32
	_ = v4016
	var v4017 int32
	_ = v4017
	var v4021 int32
	_ = v4021
	var v4037 int32
	_ = v4037
	var v4039 int32
	_ = v4039
	var v4042 int32
	_ = v4042
	var v4047 int32
	_ = v4047
	var v4048 int32
	_ = v4048
	var v4051 int32
	_ = v4051
	var v4053 int32
	_ = v4053
	var v4058 int32
	_ = v4058
	var v4059 int32
	_ = v4059
	var v4061 int32
	_ = v4061
	var v4063 int32
	_ = v4063
	var v4069 int32
	_ = v4069
	var v4078 int32
	_ = v4078
	var v4080 int32
	_ = v4080
	var v4081 int32
	_ = v4081
	var v4084 int32
	_ = v4084
	var v4085 int32
	_ = v4085
	var v4091 int32
	_ = v4091
	var v4098 int32
	_ = v4098
	var v4099 int32
	_ = v4099
	var v4100 int32
	_ = v4100
	var v4103 int32
	_ = v4103
	var v4111 int32
	_ = v4111
	var v4112 int32
	_ = v4112
	var v4113 int32
	_ = v4113
	var v4114 int32
	_ = v4114
	var v4117 int32
	_ = v4117
	var v4119 int64
	_ = v4119
	var v4124 int32
	_ = v4124
	var v4127 int32
	_ = v4127
	var v4130 int32
	_ = v4130
	var v4132 int32
	_ = v4132
	var v4138 int32
	_ = v4138
	var v4142 int32
	_ = v4142
	var v4143 int32
	_ = v4143
	var v4145 int32
	_ = v4145
	var v4146 int32
	_ = v4146
	var v4151 int32
	_ = v4151
	var v4154 int32
	_ = v4154
	var v4155 int32
	_ = v4155
	var v4161 int32
	_ = v4161
	var v4162 int32
	_ = v4162
	var v4165 int32
	_ = v4165
	var v4170 int32
	_ = v4170
	var v4171 int32
	_ = v4171
	var v4174 int32
	_ = v4174
	var v4178 int32
	_ = v4178
	var v4183 int32
	_ = v4183
	var v4184 int32
	_ = v4184
	var v4188 int32
	_ = v4188
	var v4189 int32
	_ = v4189
	var v4193 int32
	_ = v4193
	var v4194 int32
	_ = v4194
	var v4197 int32
	_ = v4197
	var v4199 int32
	_ = v4199
	var v4204 int32
	_ = v4204
	var v4205 int32
	_ = v4205
	var v4211 int32
	_ = v4211
	var v4212 int32
	_ = v4212
	var v4219 int32
	_ = v4219
	var v4223 int32
	_ = v4223
	var v4225 int32
	_ = v4225
	var v4230 int32
	_ = v4230
	var v4231 int32
	_ = v4231
	var v4238 int32
	_ = v4238
	var v4242 int32
	_ = v4242
	var v4244 int32
	_ = v4244
	var v4246 int32
	_ = v4246
	var v4248 int32
	_ = v4248
	var v4252 int32
	_ = v4252
	var v4254 int32
	_ = v4254
	var v4257 int32
	_ = v4257
	var v4262 int32
	_ = v4262
	var v4263 int32
	_ = v4263
	var v4264 int32
	_ = v4264
	var v4265 int32
	_ = v4265
	var v4268 int32
	_ = v4268
	var v4272 int32
	_ = v4272
	var v4273 int32
	_ = v4273
	var v4275 int32
	_ = v4275
	var v4276 int32
	_ = v4276
	var v4281 int32
	_ = v4281
	var v4282 int32
	_ = v4282
	var v4284 int32
	_ = v4284
	var v4285 int32
	_ = v4285
	var v4290 int32
	_ = v4290
	var v4291 int32
	_ = v4291
	var v4293 int32
	_ = v4293
	var v4294 int32
	_ = v4294
	var v4299 int32
	_ = v4299
	var v4300 int32
	_ = v4300
	var v4302 int32
	_ = v4302
	var v4303 int32
	_ = v4303
	var v4308 int32
	_ = v4308
	var v4309 int32
	_ = v4309
	var v4311 int32
	_ = v4311
	var v4312 int32
	_ = v4312
	var v4316 int32
	_ = v4316
	var v4317 int32
	_ = v4317
	var v4320 int32
	_ = v4320
	var v4321 int32
	_ = v4321
	var v4327 int32
	_ = v4327
	var v4328 int32
	_ = v4328
	var v4331 int32
	_ = v4331
	var v4333 int32
	_ = v4333
	var v4334 int32
	_ = v4334
	var v4340 int32
	_ = v4340
	var v4341 int32
	_ = v4341
	var v4346 int32
	_ = v4346
	var v4348 int32
	_ = v4348
	var v4350 int32
	_ = v4350
	var v4355 int32
	_ = v4355
	var v4356 int32
	_ = v4356
	var v4361 int32
	_ = v4361
	var v4363 int32
	_ = v4363
	var v4365 int64
	_ = v4365
	var v4367 int32
	_ = v4367
	var v4372 int32
	_ = v4372
	var v4373 int32
	_ = v4373
	var v4378 int32
	_ = v4378
	var v4380 int32
	_ = v4380
	var v4382 int64
	_ = v4382
	var v4384 int32
	_ = v4384
	var v4388 int32
	_ = v4388
	var v4392 int32
	_ = v4392
	var v4393 int32
	_ = v4393
	var v4396 int32
	_ = v4396
	var v4401 int32
	_ = v4401
	var v4402 int32
	_ = v4402
	var v4409 int32
	_ = v4409
	var v4412 int32
	_ = v4412
	var v4417 int32
	_ = v4417
	var v4419 int32
	_ = v4419
	var v4422 int32
	_ = v4422
	var v4428 int32
	_ = v4428
	var v4429 int32
	_ = v4429
	var v4434 int32
	_ = v4434
	var v4435 int32
	_ = v4435
	var v4436 int32
	_ = v4436
	var v4437 int32
	_ = v4437
	var v4442 int32
	_ = v4442
	var v4443 int32
	_ = v4443
	var v4445 int32
	_ = v4445
	var v4446 int32
	_ = v4446
	var v4449 int32
	_ = v4449
	var v4450 int32
	_ = v4450
	var v4451 int32
	_ = v4451
	var v4456 int32
	_ = v4456
	var v4457 int32
	_ = v4457
	var v4458 int32
	_ = v4458
	var v4459 int32
	_ = v4459
	var v4462 int32
	_ = v4462
	var v4468 int32
	_ = v4468
	var v4469 int32
	_ = v4469
	var v4472 int32
	_ = v4472
	var v4475 int32
	_ = v4475
	var v4476 int32
	_ = v4476
	var v4481 int32
	_ = v4481
	var v4482 int32
	_ = v4482
	var v4483 int32
	_ = v4483
	var v4484 int32
	_ = v4484
	var v4486 int32
	_ = v4486
	var v4487 int32
	_ = v4487
	var v4492 int32
	_ = v4492
	var v4493 int32
	_ = v4493
	var v4494 int32
	_ = v4494
	var v4495 int32
	_ = v4495
	var v4497 int32
	_ = v4497
	var v4498 int32
	_ = v4498
	var v4503 int32
	_ = v4503
	var v4504 int32
	_ = v4504
	var v4505 int32
	_ = v4505
	var v4506 int32
	_ = v4506
	var v4508 int32
	_ = v4508
	var v4509 int32
	_ = v4509
	var v4512 int32
	_ = v4512
	var v4556 int32
	_ = v4556
	var v4557 int32
	_ = v4557
	var v4560 int32
	_ = v4560
	var v4564 int32
	_ = v4564
	var v4569 int32
	_ = v4569
	var v4570 int32
	_ = v4570
	var v4571 int32
	_ = v4571
	var v4574 int32
	_ = v4574
	var v4577 int32
	_ = v4577
	var v4578 int32
	_ = v4578
	var v4581 int32
	_ = v4581
	var v4587 int32
	_ = v4587
	var v4590 int32
	_ = v4590
	var v4595 int32
	_ = v4595
	var v4599 int32
	_ = v4599
	var v4600 int32
	_ = v4600
	var v4601 int32
	_ = v4601
	var v4607 int32
	_ = v4607
	var v4612 int32
	_ = v4612
	var v4614 int32
	_ = v4614
	var v4624 int32
	_ = v4624
	var v4629 int32
	_ = v4629
	var v4636 int32
	_ = v4636
	var v4639 int32
	_ = v4639
	var v4640 int32
	_ = v4640
	var v4646 int32
	_ = v4646
	var v4651 int32
	_ = v4651
	var v4655 int32
	_ = v4655
	var v4658 int32
	_ = v4658
	var v4659 int32
	_ = v4659
	var v4665 int32
	_ = v4665
	var v4670 int32
	_ = v4670
	var v4671 int32
	_ = v4671
	var v4673 int32
	_ = v4673
	var v4681 int32
	_ = v4681
	var v4682 int32
	_ = v4682
	var v4684 int32
	_ = v4684
	var v4685 int32
	_ = v4685
	var v4691 int32
	_ = v4691
	var v4696 int32
	_ = v4696
	var v4698 int32
	_ = v4698
	var v4699 int32
	_ = v4699
	var v4707 int32
	_ = v4707
	var v4709 int32
	_ = v4709
	var v4712 int32
	_ = v4712
	var v4715 int32
	_ = v4715
	var v4718 int32
	_ = v4718
	var v4719 int32
	_ = v4719
	var v4720 int32
	_ = v4720
	var v4726 int32
	_ = v4726
	var v4727 int64
	_ = v4727
	var v4731 int32
	_ = v4731
	var v4735 int32
	_ = v4735
	var v4736 int32
	_ = v4736
	var v4740 int32
	_ = v4740
	var v4745 int32
	_ = v4745
	var v4746 int32
	_ = v4746
	var v4748 int32
	_ = v4748
	var v4750 int32
	_ = v4750
	var v4755 int64
	_ = v4755
	var v4756 int32
	_ = v4756
	var v4759 int64
	_ = v4759
	var v4760 int32
	_ = v4760
	var v4761 int32
	_ = v4761
	var v4764 int64
	_ = v4764
	var v4765 int32
	_ = v4765
	var v4767 int64
	_ = v4767
	var v4769 int32
	_ = v4769
	var v4771 int32
	_ = v4771
	var v4772 int32
	_ = v4772
	var v4773 int64
	_ = v4773
	var v4776 int64
	_ = v4776
	var v4778 int32
	_ = v4778
	var v4781 int32
	_ = v4781
	var v4784 int32
	_ = v4784
	var v4788 int64
	_ = v4788
	var v4791 int32
	_ = v4791
	var v4792 int32
	_ = v4792
	var v4795 int64
	_ = v4795
	var v4799 int64
	_ = v4799
	var v4802 int64
	_ = v4802
	var v4810 int32
	_ = v4810
	var v4811 int32
	_ = v4811
	var v4813 int32
	_ = v4813
	var v4815 int32
	_ = v4815
	var v4817 int32
	_ = v4817
	var v4819 int32
	_ = v4819
	var v4820 int32
	_ = v4820
	var v4821 int32
	_ = v4821
	var v4822 int32
	_ = v4822
	var v4823 int32
	_ = v4823
	var v4825 int32
	_ = v4825
	var v4826 int32
	_ = v4826
	var v4828 int32
	_ = v4828
	var v4829 int32
	_ = v4829
	var v4831 int32
	_ = v4831
	var v4832 int32
	_ = v4832
	var v4837 int32
	_ = v4837
	var v4842 int32
	_ = v4842
	var v4847 int32
	_ = v4847
	var v4852 int32
	_ = v4852
	var v4853 int32
	_ = v4853
	var v4862 int32
	_ = v4862
	var v4866 int32
	_ = v4866
	var v4873 int32
	_ = v4873
	var v4874 int32
	_ = v4874
	var v4876 int32
	_ = v4876
	var v4882 int32
	_ = v4882
	var v4885 int32
	_ = v4885
	var v4887 int32
	_ = v4887
	var v4890 int32
	_ = v4890
	var v4893 int32
	_ = v4893
	var v4896 int32
	_ = v4896
	var v4899 int32
	_ = v4899
	var v4903 int32
	_ = v4903
	var v4904 int32
	_ = v4904
	var v4907 int32
	_ = v4907
	var v4910 int32
	_ = v4910
	var v4916 int32
	_ = v4916
	var v4922 int32
	_ = v4922
	var v4924 int32
	_ = v4924
	var v4930 int32
	_ = v4930
	var v4937 int32
	_ = v4937
	var v4941 int32
	_ = v4941
	var v4942 int32
	_ = v4942
	var v4945 int32
	_ = v4945
	var v4946 int32
	_ = v4946
	var v4949 int64
	_ = v4949
	var v4953 int32
	_ = v4953
	var v4954 int32
	_ = v4954
	var v4957 int32
	_ = v4957
	var v4958 int32
	_ = v4958
	var v4961 int32
	_ = v4961
	var v4968 int32
	_ = v4968
	var v4970 int32
	_ = v4970
	var v4972 int32
	_ = v4972
	var v4973 int64
	_ = v4973
	var v4980 int32
	_ = v4980
	var v4981 int32
	_ = v4981
	var v4986 int32
	_ = v4986
	var v4991 int32
	_ = v4991
	var v4996 int32
	_ = v4996
	var v4997 int32
	_ = v4997
	var v5006 int32
	_ = v5006
	var v5010 int32
	_ = v5010
	var v5017 int32
	_ = v5017
	var v5018 int32
	_ = v5018
	var v5020 int32
	_ = v5020
	var v5026 int32
	_ = v5026
	var v5029 int32
	_ = v5029
	var v5031 int32
	_ = v5031
	var v5034 int32
	_ = v5034
	var v5037 int32
	_ = v5037
	var v5040 int32
	_ = v5040
	var v5043 int32
	_ = v5043
	var v5047 int32
	_ = v5047
	var v5048 int32
	_ = v5048
	var v5051 int32
	_ = v5051
	var v5054 int32
	_ = v5054
	var v5060 int32
	_ = v5060
	var v5066 int32
	_ = v5066
	var v5068 int32
	_ = v5068
	var v5074 int32
	_ = v5074
	var v5081 int32
	_ = v5081
	var v5084 int32
	_ = v5084
	var v5086 int32
	_ = v5086
	var v5089 int32
	_ = v5089
	var v5093 int32
	_ = v5093
	var v5094 int32
	_ = v5094
	var v5095 int32
	_ = v5095
	var v5097 int32
	_ = v5097
	var v5098 int32
	_ = v5098
	var v5099 int32
	_ = v5099
	var v5101 int32
	_ = v5101
	var v5105 int32
	_ = v5105
	var v5108 int32
	_ = v5108
	var v5111 int32
	_ = v5111
	var v5116 int32
	_ = v5116
	var v5118 int32
	_ = v5118
	var v5120 int64
	_ = v5120
	var v5122 int64
	_ = v5122
	var v5130 int32
	_ = v5130
	var v5134 int32
	_ = v5134
	var v5138 int32
	_ = v5138
	var v5140 int32
	_ = v5140
	var v5141 int32
	_ = v5141
	var v5142 int32
	_ = v5142
	var v5150 int64
	_ = v5150
	var v5153 int32
	_ = v5153
	var v5158 int32
	_ = v5158
	var v5159 int32
	_ = v5159
	var v5160 int32
	_ = v5160
	var v5161 int32
	_ = v5161
	var v5162 int32
	_ = v5162
	var v5166 int64
	_ = v5166
	var v5171 int32
	_ = v5171
	var v5176 int32
	_ = v5176
	var v5177 int32
	_ = v5177
	var v5179 int32
	_ = v5179
	var v5181 int32
	_ = v5181
	var v5182 int64
	_ = v5182
	var v5183 int32
	_ = v5183
	var v5184 int32
	_ = v5184
	var v5186 int32
	_ = v5186
	var v5187 int32
	_ = v5187
	var v5189 int32
	_ = v5189
	var v5190 int32
	_ = v5190
	var v5191 int32
	_ = v5191
	var v5192 int64
	_ = v5192
	var v5193 int32
	_ = v5193
	var v5194 int32
	_ = v5194
	var v5195 int32
	_ = v5195
	var v5204 int32
	_ = v5204
	var v5205 int32
	_ = v5205
	var v5207 int32
	_ = v5207
	var v5208 int32
	_ = v5208
	var v5214 int32
	_ = v5214
	var v5216 int32
	_ = v5216
	var v5218 int32
	_ = v5218
	var v5219 int32
	_ = v5219
	var v5221 int32
	_ = v5221
	var v5224 int32
	_ = v5224
	var v5230 int32
	_ = v5230
	var v5235 int32
	_ = v5235
	var v5245 int32
	_ = v5245
	var v5246 int32
	_ = v5246
	var v5251 int32
	_ = v5251
	var v5252 int32
	_ = v5252
	var v5256 int32
	_ = v5256
	var v5257 int32
	_ = v5257
	var v5259 int32
	_ = v5259
	var v5267 int32
	_ = v5267
	var v5271 int32
	_ = v5271
	var v5272 int32
	_ = v5272
	var v5273 int32
	_ = v5273
	var v5276 int32
	_ = v5276
	var v5279 int32
	_ = v5279
	var v5282 int32
	_ = v5282
	var v5283 int32
	_ = v5283
	var v5286 int32
	_ = v5286
	var v5287 int32
	_ = v5287
	var v5290 int32
	_ = v5290
	var v5297 int32
	_ = v5297
	var v5298 int32
	_ = v5298
	var v5304 int32
	_ = v5304
	var v5307 int32
	_ = v5307
	var v5308 int32
	_ = v5308
	var v5309 int32
	_ = v5309
	var v5312 int32
	_ = v5312
	var v5315 int32
	_ = v5315
	var v5318 int32
	_ = v5318
	var v5319 int32
	_ = v5319
	var v5322 int32
	_ = v5322
	var v5323 int32
	_ = v5323
	var v5326 int32
	_ = v5326
	var v5333 int32
	_ = v5333
	var v5334 int32
	_ = v5334
	var v5340 int32
	_ = v5340
	var v5341 int32
	_ = v5341
	var v5344 int32
	_ = v5344
	var v5347 int32
	_ = v5347
	var v5350 int32
	_ = v5350
	var v5351 int32
	_ = v5351
	var v5354 int32
	_ = v5354
	var v5355 int32
	_ = v5355
	var v5358 int32
	_ = v5358
	var v5365 int32
	_ = v5365
	var v5366 int32
	_ = v5366
	var v5371 int32
	_ = v5371
	var v5374 int32
	_ = v5374
	var v5377 int32
	_ = v5377
	var v5380 int32
	_ = v5380
	var v5381 int32
	_ = v5381
	var v5384 int32
	_ = v5384
	var v5385 int32
	_ = v5385
	var v5388 int32
	_ = v5388
	var v5395 int32
	_ = v5395
	var v5396 int32
	_ = v5396
	var v5404 int32
	_ = v5404
	var v5407 int32
	_ = v5407
	var v5408 int32
	_ = v5408
	var v5417 int32
	_ = v5417
	var v5422 int32
	_ = v5422
	var v5423 int32
	_ = v5423
	var v5426 int32
	_ = v5426
	var v5429 int32
	_ = v5429
	var v5432 int32
	_ = v5432
	var v5433 int32
	_ = v5433
	var v5436 int32
	_ = v5436
	var v5437 int32
	_ = v5437
	var v5440 int32
	_ = v5440
	var v5447 int32
	_ = v5447
	var v5448 int32
	_ = v5448
	var v5454 int32
	_ = v5454
	var v5456 int32
	_ = v5456
	var v5457 int32
	_ = v5457
	var v5458 int32
	_ = v5458
	var v5461 int32
	_ = v5461
	var v5464 int32
	_ = v5464
	var v5467 int32
	_ = v5467
	var v5468 int32
	_ = v5468
	var v5471 int32
	_ = v5471
	var v5472 int32
	_ = v5472
	var v5475 int32
	_ = v5475
	var v5482 int32
	_ = v5482
	var v5483 int32
	_ = v5483
	var v5487 int32
	_ = v5487
	var v5491 int32
	_ = v5491
	var v5492 int32
	_ = v5492
	var v5493 int32
	_ = v5493
	var v5496 int32
	_ = v5496
	var v5499 int32
	_ = v5499
	var v5502 int32
	_ = v5502
	var v5503 int32
	_ = v5503
	var v5506 int32
	_ = v5506
	var v5507 int32
	_ = v5507
	var v5510 int32
	_ = v5510
	var v5517 int32
	_ = v5517
	var v5518 int32
	_ = v5518
	var v5522 int32
	_ = v5522
	var v5526 int32
	_ = v5526
	var v5527 int32
	_ = v5527
	var v5529 int32
	_ = v5529
	var v5530 int32
	_ = v5530
	var v5531 int32
	_ = v5531
	var v5532 int32
	_ = v5532
	var v5533 int32
	_ = v5533
	var v5534 int32
	_ = v5534
	var v5535 int32
	_ = v5535
	var v5536 int32
	_ = v5536
	var v5538 int32
	_ = v5538
	var v5539 int32
	_ = v5539
	var v5548 int32
	_ = v5548
	var v5559 int32
	_ = v5559
	var v5564 int32
	_ = v5564
	var v5572 int32
	_ = v5572
	var v5580 int32
	_ = v5580
	var v5583 int32
	_ = v5583
	var v5584 int32
	_ = v5584
	var v5586 int32
	_ = v5586
	var v5596 int32
	_ = v5596
	var v5602 int32
	_ = v5602
	var v5604 int32
	_ = v5604
	var v5605 int32
	_ = v5605
	var v5607 int32
	_ = v5607
	var v5608 int32
	_ = v5608
	var v5611 int32
	_ = v5611
	var v5612 int32
	_ = v5612
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
	var v5625 int32
	_ = v5625
	var v5627 int32
	_ = v5627
	var v5629 int32
	_ = v5629
	var v5630 int32
	_ = v5630
	var v5638 int32
	_ = v5638
	var v5645 int32
	_ = v5645
	var v5650 int32
	_ = v5650
	var v5652 int32
	_ = v5652
	var v5653 int32
	_ = v5653
	var v5659 int32
	_ = v5659
	var v5663 int32
	_ = v5663
	var v5666 int32
	_ = v5666
	var v5668 int32
	_ = v5668
	var v5672 int32
	_ = v5672
	var v5673 int32
	_ = v5673
	var v5676 int32
	_ = v5676
	var v5678 int32
	_ = v5678
	var v5679 int32
	_ = v5679
	var v5693 int32
	_ = v5693
	var v5694 int32
	_ = v5694
	var v5699 int32
	_ = v5699
	var v5700 int32
	_ = v5700
	var v5701 int32
	_ = v5701
	var v5703 int32
	_ = v5703
	var v5706 int32
	_ = v5706
	var v5707 int32
	_ = v5707
	var v5713 int32
	_ = v5713
	var v5715 int32
	_ = v5715
	var v5719 int32
	_ = v5719
	var v5722 int32
	_ = v5722
	var v5724 int32
	_ = v5724
	var v5729 int32
	_ = v5729
	var v5730 int32
	_ = v5730
	var v5731 int32
	_ = v5731
	var v5732 int32
	_ = v5732
	var v5735 int32
	_ = v5735
	var v5736 int32
	_ = v5736
	var v5737 int32
	_ = v5737
	var v5743 int32
	_ = v5743
	var v5748 int32
	_ = v5748
	var v5756 int32
	_ = v5756
	var v5760 int32
	_ = v5760
	var v5765 int32
	_ = v5765
	var v5769 int32
	_ = v5769
	var v5773 int32
	_ = v5773
	var v5778 int32
	_ = v5778
	var v5779 int32
	_ = v5779
	var v5780 int32
	_ = v5780
	var v5781 int32
	_ = v5781
	var v5783 int32
	_ = v5783
	var v5785 int32
	_ = v5785
	var v5789 int32
	_ = v5789
	var v5791 int32
	_ = v5791
	var v5792 int32
	_ = v5792
	var v5794 int32
	_ = v5794
	var v5799 int32
	_ = v5799
	var v5801 int32
	_ = v5801
	var v5802 int64
	_ = v5802
	var v5805 int64
	_ = v5805
	var v5808 int32
	_ = v5808
	var v5813 int32
	_ = v5813
	var v5814 int32
	_ = v5814
	var v5816 int32
	_ = v5816
	var v5817 int32
	_ = v5817
	var v5819 int32
	_ = v5819
	var v5820 int32
	_ = v5820
	var v5825 int32
	_ = v5825
	var v5830 int32
	_ = v5830
	var v5835 int32
	_ = v5835
	var v5840 int32
	_ = v5840
	var v5841 int32
	_ = v5841
	var v5850 int32
	_ = v5850
	var v5854 int32
	_ = v5854
	var v5861 int32
	_ = v5861
	var v5862 int32
	_ = v5862
	var v5864 int32
	_ = v5864
	var v5870 int32
	_ = v5870
	var v5873 int32
	_ = v5873
	var v5875 int32
	_ = v5875
	var v5878 int32
	_ = v5878
	var v5881 int32
	_ = v5881
	var v5884 int32
	_ = v5884
	var v5887 int32
	_ = v5887
	var v5891 int32
	_ = v5891
	var v5892 int32
	_ = v5892
	var v5895 int32
	_ = v5895
	var v5898 int32
	_ = v5898
	var v5904 int32
	_ = v5904
	var v5910 int32
	_ = v5910
	var v5912 int32
	_ = v5912
	var v5918 int32
	_ = v5918
	var v5925 int32
	_ = v5925
	var v5929 int32
	_ = v5929
	var v5930 int32
	_ = v5930
	var v5932 int32
	_ = v5932
	var v5935 int32
	_ = v5935
	var v5936 int32
	_ = v5936
	var v5939 int32
	_ = v5939
	var v5940 int32
	_ = v5940
	var v5943 int32
	_ = v5943
	var v5944 int32
	_ = v5944
	var v5947 int32
	_ = v5947
	var v5949 int32
	_ = v5949
	var v5950 int32
	_ = v5950
	var v5951 int32
	_ = v5951
	var v5954 int32
	_ = v5954
	var v5961 int32
	_ = v5961
	var v5963 int32
	_ = v5963
	var v5965 int32
	_ = v5965
	var v5968 int32
	_ = v5968
	var v5969 int32
	_ = v5969
	var v5970 int32
	_ = v5970
	var v5976 int32
	_ = v5976
	var v5978 int32
	_ = v5978
	var v5979 int32
	_ = v5979
	var v5980 int32
	_ = v5980
	var v5983 int32
	_ = v5983
	var v5984 int32
	_ = v5984
	var v5990 int32
	_ = v5990
	var v6011 int32
	_ = v6011
	var v6012 int32
	_ = v6012
	var v6016 int32
	_ = v6016
	var v6017 int32
	_ = v6017
	var v6027 int32
	_ = v6027
	var v6031 int32
	_ = v6031
	var v6032 int32
	_ = v6032
	var v6033 int32
	_ = v6033
	var v6036 int32
	_ = v6036
	var v6039 int32
	_ = v6039
	var v6042 int32
	_ = v6042
	var v6043 int32
	_ = v6043
	var v6046 int32
	_ = v6046
	var v6047 int32
	_ = v6047
	var v6050 int32
	_ = v6050
	var v6057 int32
	_ = v6057
	var v6058 int32
	_ = v6058
	var v6065 int32
	_ = v6065
	var v6066 int32
	_ = v6066
	var v6067 int32
	_ = v6067
	var v6070 int32
	_ = v6070
	var v6073 int32
	_ = v6073
	var v6076 int32
	_ = v6076
	var v6077 int32
	_ = v6077
	var v6080 int32
	_ = v6080
	var v6081 int32
	_ = v6081
	var v6084 int32
	_ = v6084
	var v6091 int32
	_ = v6091
	var v6092 int32
	_ = v6092
	var v6097 int32
	_ = v6097
	var v6098 int32
	_ = v6098
	var v6099 int32
	_ = v6099
	var v6100 int32
	_ = v6100
	var v6101 int32
	_ = v6101
	var v6102 int32
	_ = v6102
	var v6104 int32
	_ = v6104
	var v6105 int32
	_ = v6105
	var v6112 int32
	_ = v6112
	var v6118 int32
	_ = v6118
	var v6143 int32
	_ = v6143
	var v6147 int32
	_ = v6147
	var v6148 int32
	_ = v6148
	var v6158 int32
	_ = v6158
	var v6161 int32
	_ = v6161
	var v6162 int32
	_ = v6162
	var v6163 int32
	_ = v6163
	var v6165 int32
	_ = v6165
	var v6170 int32
	_ = v6170
	var v6172 int32
	_ = v6172
	var v6173 int32
	_ = v6173
	var v6176 int32
	_ = v6176
	var v6181 int32
	_ = v6181
	var v6182 int32
	_ = v6182
	var v6184 int32
	_ = v6184
	var v6186 int32
	_ = v6186
	var v6188 int32
	_ = v6188
	var v6189 int32
	_ = v6189
	var v6195 int32
	_ = v6195
	var v6201 int32
	_ = v6201
	var v6204 int32
	_ = v6204
	var v6208 int32
	_ = v6208
	var v6213 int32
	_ = v6213
	var v6216 int32
	_ = v6216
	var v6218 int32
	_ = v6218
	var v6219 int32
	_ = v6219
	var v6223 int32
	_ = v6223
	var v6226 int32
	_ = v6226
	var v6227 int32
	_ = v6227
	var v6228 int32
	_ = v6228
	var v6230 int32
	_ = v6230
	var v6233 int32
	_ = v6233
	var v6236 int32
	_ = v6236
	var v6238 int32
	_ = v6238
	var v6239 int32
	_ = v6239
	var v6241 int32
	_ = v6241
	var v6246 int32
	_ = v6246
	var v6251 int32
	_ = v6251
	var v6252 int32
	_ = v6252
	var v6253 int32
	_ = v6253
	var v6257 int32
	_ = v6257
	var v6260 int32
	_ = v6260
	var v6262 int32
	_ = v6262
	var v6263 int32
	_ = v6263
	var v6265 int32
	_ = v6265
	var v6273 int32
	_ = v6273
	var v6276 int32
	_ = v6276
	var v6279 int32
	_ = v6279
	var v6280 int32
	_ = v6280
	var v6281 int32
	_ = v6281
	var v6282 int32
	_ = v6282
	var v6284 int32
	_ = v6284
	var v6290 int32
	_ = v6290
	var v6295 int32
	_ = v6295
	var v6299 int32
	_ = v6299
	var v6300 int32
	_ = v6300
	var v6302 int32
	_ = v6302
	var v6305 int32
	_ = v6305
	var v6308 int32
	_ = v6308
	var v6315 int32
	_ = v6315
	var v6318 int32
	_ = v6318
	var v6323 int32
	_ = v6323
	var v6328 int32
	_ = v6328
	var v6332 int32
	_ = v6332
	var v6335 int32
	_ = v6335
	var v6341 int32
	_ = v6341
	var v6344 int32
	_ = v6344
	var v6345 int32
	_ = v6345
	var v6350 int32
	_ = v6350
	var v6354 int32
	_ = v6354
	var v6357 int32
	_ = v6357
	var v6361 int32
	_ = v6361
	var v6366 int32
	_ = v6366
	var v6367 int32
	_ = v6367
	var v6371 int32
	_ = v6371
	var v6372 int32
	_ = v6372
	var v6375 int32
	_ = v6375
	var v6377 int32
	_ = v6377
	var v6383 int32
	_ = v6383
	var v6387 int32
	_ = v6387
	var v6391 int32
	_ = v6391
	var v6392 int32
	_ = v6392
	var v6394 int32
	_ = v6394
	var v6395 int32
	_ = v6395
	var v6398 int32
	_ = v6398
	var v6400 int32
	_ = v6400
	var v6401 int32
	_ = v6401
	var v6405 int32
	_ = v6405
	var v6410 int32
	_ = v6410
	var v6411 int32
	_ = v6411
	var v6413 int32
	_ = v6413
	var v6415 int32
	_ = v6415
	var v6420 int64
	_ = v6420
	var v6421 int32
	_ = v6421
	var v6424 int64
	_ = v6424
	var v6425 int32
	_ = v6425
	var v6426 int32
	_ = v6426
	var v6429 int64
	_ = v6429
	var v6430 int32
	_ = v6430
	var v6432 int64
	_ = v6432
	var v6434 int32
	_ = v6434
	var v6436 int32
	_ = v6436
	var v6437 int32
	_ = v6437
	var v6438 int64
	_ = v6438
	var v6441 int64
	_ = v6441
	var v6443 int32
	_ = v6443
	var v6446 int32
	_ = v6446
	var v6449 int32
	_ = v6449
	var v6453 int64
	_ = v6453
	var v6456 int32
	_ = v6456
	var v6457 int32
	_ = v6457
	var v6460 int64
	_ = v6460
	var v6464 int64
	_ = v6464
	var v6465 int32
	_ = v6465
	var v6468 int32
	_ = v6468
	var v6471 int32
	_ = v6471
	var v6474 int32
	_ = v6474
	var v6476 int32
	_ = v6476
	var v6477 int32
	_ = v6477
	var v6478 int32
	_ = v6478
	var v6480 int64
	_ = v6480
	var v6481 int32
	_ = v6481
	var v6483 int32
	_ = v6483
	var v6486 int64
	_ = v6486
	var v6491 int32
	_ = v6491
	var v6492 int64
	_ = v6492
	var v6493 int32
	_ = v6493
	var v6497 int64
	_ = v6497
	var v6503 int32
	_ = v6503
	var v6504 int32
	_ = v6504
	var v6508 int64
	_ = v6508
	var v6513 int32
	_ = v6513
	var v6514 int32
	_ = v6514
	var v6519 int32
	_ = v6519
	var v6521 int32
	_ = v6521
	var v6527 int32
	_ = v6527
	var v6538 int32
	_ = v6538
	var v6541 int32
	_ = v6541
	var v6545 int32
	_ = v6545
	var v6548 int32
	_ = v6548
	var v6549 int32
	_ = v6549
	var v6554 int32
	_ = v6554
	var v6558 int32
	_ = v6558
	var v6561 int32
	_ = v6561
	var v6565 int32
	_ = v6565
	var v6570 int32
	_ = v6570
	var v6575 int64
	_ = v6575
	var v6579 int32
	_ = v6579
	var v6585 int32
	_ = v6585
	var v6588 int64
	_ = v6588
	var v6593 int32
	_ = v6593
	var v6594 int32
	_ = v6594
	var v6599 int32
	_ = v6599
	var v6604 int32
	_ = v6604
	var v6607 int32
	_ = v6607
	var v6611 int32
	_ = v6611
	var v6614 int32
	_ = v6614
	var v6617 int32
	_ = v6617
	var v6618 int32
	_ = v6618
	var v6619 int32
	_ = v6619
	var v6621 int32
	_ = v6621
	var v6628 int32
	_ = v6628
	var v6629 int32
	_ = v6629
	var v6630 int32
	_ = v6630
	var v6632 int32
	_ = v6632
	var v6638 int32
	_ = v6638
	var v6640 int32
	_ = v6640
	var v6641 int32
	_ = v6641
	var v6642 int32
	_ = v6642
	var v6643 int32
	_ = v6643
	var v6644 int64
	_ = v6644
	var v6649 int32
	_ = v6649
	var v6652 int32
	_ = v6652
	var v6657 int32
	_ = v6657
	var v6659 int32
	_ = v6659
	var v6661 int64
	_ = v6661
	var v6663 int32
	_ = v6663
	var v6667 int32
	_ = v6667
	var v6673 int32
	_ = v6673
	var v6678 int32
	_ = v6678
	var v6680 int32
	_ = v6680
	var v6681 int32
	_ = v6681
	var v6686 int32
	_ = v6686
	var v6691 int32
	_ = v6691
	var v6692 int32
	_ = v6692
	var v6701 int32
	_ = v6701
	var v6703 int32
	_ = v6703
	var v6705 int32
	_ = v6705
	var v6709 int64
	_ = v6709
	var v6711 int32
	_ = v6711
	var v6714 int64
	_ = v6714
	var v6717 int32
	_ = v6717
	var v6722 int32
	_ = v6722
	var v6723 int32
	_ = v6723
	var v6725 int32
	_ = v6725
	var v6726 int32
	_ = v6726
	var v6728 int32
	_ = v6728
	var v6729 int32
	_ = v6729
	var v6734 int32
	_ = v6734
	var v6739 int32
	_ = v6739
	var v6740 int32
	_ = v6740
	var v6749 int32
	_ = v6749
	var v6753 int32
	_ = v6753
	var v6760 int32
	_ = v6760
	var v6761 int32
	_ = v6761
	var v6763 int32
	_ = v6763
	var v6769 int32
	_ = v6769
	var v6772 int32
	_ = v6772
	var v6774 int32
	_ = v6774
	var v6777 int32
	_ = v6777
	var v6780 int32
	_ = v6780
	var v6783 int32
	_ = v6783
	var v6786 int32
	_ = v6786
	var v6790 int32
	_ = v6790
	var v6791 int32
	_ = v6791
	var v6794 int32
	_ = v6794
	var v6797 int32
	_ = v6797
	var v6803 int32
	_ = v6803
	var v6809 int32
	_ = v6809
	var v6811 int32
	_ = v6811
	var v6817 int32
	_ = v6817
	var v6824 int32
	_ = v6824
	var v6828 int32
	_ = v6828
	var v6829 int32
	_ = v6829
	var v6831 int64
	_ = v6831
	var v6833 int32
	_ = v6833
	var v6834 int32
	_ = v6834
	var v6842 int32
	_ = v6842
	var v6844 int32
	_ = v6844
	var v6851 int32
	_ = v6851
	var v6858 int32
	_ = v6858
	var v6859 int64
	_ = v6859
	var v6861 int64
	_ = v6861
	var v6862 int64
	_ = v6862
	var v6866 int64
	_ = v6866
	var v6870 int32
	_ = v6870
	var v6875 int32
	_ = v6875
	var v6878 int32
	_ = v6878
	var v6881 int32
	_ = v6881
	var v6882 int32
	_ = v6882
	var v6883 int32
	_ = v6883
	var v6886 int32
	_ = v6886
	var v6888 int32
	_ = v6888
	var v6893 int32
	_ = v6893
	var v6898 int32
	_ = v6898
	var v6899 int32
	_ = v6899
	var v6901 int32
	_ = v6901
	var v6903 int32
	_ = v6903
	var v6906 int32
	_ = v6906
	var v6907 int32
	_ = v6907
	var v6911 int32
	_ = v6911
	var v6916 int32
	_ = v6916
	var v6920 int32
	_ = v6920
	var v6921 int64
	_ = v6921
	var v6935 int32
	_ = v6935
	var v6936 int32
	_ = v6936
	var v6939 int32
	_ = v6939
	var v6942 int32
	_ = v6942
	var v6943 int32
	_ = v6943
	var v6948 int32
	_ = v6948
	var v6953 int32
	_ = v6953
	var v6956 int32
	_ = v6956
	var v6960 int32
	_ = v6960
	var v6963 int32
	_ = v6963
	var v6966 int32
	_ = v6966
	var v6967 int32
	_ = v6967
	var v6968 int32
	_ = v6968
	var v6970 int32
	_ = v6970
	var v6977 int32
	_ = v6977
	var v6978 int32
	_ = v6978
	var v6979 int32
	_ = v6979
	var v6981 int32
	_ = v6981
	var v6987 int32
	_ = v6987
	var v6989 int32
	_ = v6989
	var v6990 int32
	_ = v6990
	var v6991 int32
	_ = v6991
	var v6992 int32
	_ = v6992
	var v6994 int32
	_ = v6994
	var v6995 int32
	_ = v6995
	var v6997 int32
	_ = v6997
	var v6998 int64
	_ = v6998
	var v7000 int32
	_ = v7000
	var v7003 int32
	_ = v7003
	var v7004 int64
	_ = v7004
	var v7007 int32
	_ = v7007
	var v7010 int32
	_ = v7010
	var v7015 int32
	_ = v7015
	var v7017 int32
	_ = v7017
	var v7019 int32
	_ = v7019
	var v7020 int64
	_ = v7020
	var v7022 int32
	_ = v7022
	var v7029 int32
	_ = v7029
	var v7032 int32
	_ = v7032
	var v7034 int32
	_ = v7034
	var v7036 int32
	_ = v7036
	var v7038 int32
	_ = v7038
	var v7043 int32
	_ = v7043
	var v7045 int32
	_ = v7045
	var v7046 int32
	_ = v7046
	var v7049 int32
	_ = v7049
	var v7054 int32
	_ = v7054
	var v7055 int32
	_ = v7055
	var v7068 int32
	_ = v7068
	var v7072 int32
	_ = v7072
	var v7073 int32
	_ = v7073
	var v7075 int32
	_ = v7075
	var v7076 int32
	_ = v7076
	var v7078 int32
	_ = v7078
	var v7079 int32
	_ = v7079
	var v7084 int32
	_ = v7084
	var v7089 int32
	_ = v7089
	var v7090 int32
	_ = v7090
	var v7099 int32
	_ = v7099
	var v7103 int32
	_ = v7103
	var v7110 int32
	_ = v7110
	var v7111 int32
	_ = v7111
	var v7113 int32
	_ = v7113
	var v7119 int32
	_ = v7119
	var v7122 int32
	_ = v7122
	var v7124 int32
	_ = v7124
	var v7127 int32
	_ = v7127
	var v7130 int32
	_ = v7130
	var v7133 int32
	_ = v7133
	var v7136 int32
	_ = v7136
	var v7140 int32
	_ = v7140
	var v7141 int32
	_ = v7141
	var v7144 int32
	_ = v7144
	var v7147 int32
	_ = v7147
	var v7153 int32
	_ = v7153
	var v7159 int32
	_ = v7159
	var v7161 int32
	_ = v7161
	var v7167 int32
	_ = v7167
	var v7174 int32
	_ = v7174
	var v7177 int32
	_ = v7177
	var v7180 int32
	_ = v7180
	var v7185 int32
	_ = v7185
	var v7186 int32
	_ = v7186
	var v7187 int32
	_ = v7187
	var v7190 int32
	_ = v7190
	var v7195 int32
	_ = v7195
	var v7196 int32
	_ = v7196
	var v7198 int32
	_ = v7198
	var v7200 int32
	_ = v7200
	var v7202 int32
	_ = v7202
	var v7205 int32
	_ = v7205
	var v7208 int32
	_ = v7208
	var v7209 int32
	_ = v7209
	var v7210 int32
	_ = v7210
	var v7212 int32
	_ = v7212
	var v7217 int32
	_ = v7217
	var v7220 int32
	_ = v7220
	var v7221 int32
	_ = v7221
	var v7222 int32
	_ = v7222
	var v7226 int32
	_ = v7226
	var v7238 int32
	_ = v7238
	var v7240 int32
	_ = v7240
	var v7241 int32
	_ = v7241
	var v7244 int64
	_ = v7244
	var v7246 int64
	_ = v7246
	var v7249 int64
	_ = v7249
	var v7251 int64
	_ = v7251
	var v7256 int32
	_ = v7256
	var v7257 int32
	_ = v7257
	var v7258 int32
	_ = v7258
	var v7260 int32
	_ = v7260
	var v7261 int32
	_ = v7261
	var v7310 int64
	_ = v7310
	var v7315 int32
	_ = v7315
	var v7316 int32
	_ = v7316
	var v7320 int32
	_ = v7320
	var v7322 int32
	_ = v7322
	var v7324 int32
	_ = v7324
	var v7325 int32
	_ = v7325
	var v7334 int32
	_ = v7334
	var v7336 int64
	_ = v7336
	var v7378 int32
	_ = v7378
	var v7379 int32
	_ = v7379
	var v7383 int32
	_ = v7383
	var v7388 int32
	_ = v7388
	var v7391 int32
	_ = v7391
	var v7394 int32
	_ = v7394
	var v7399 int32
	_ = v7399
	var v7400 int32
	_ = v7400
	var v7401 int32
	_ = v7401
	var v7402 int32
	_ = v7402
	var v7406 int32
	_ = v7406
	var v7407 int32
	_ = v7407
	var v7412 int32
	_ = v7412
	var v7414 int32
	_ = v7414
	var v7415 int32
	_ = v7415
	var v7421 int32
	_ = v7421
	var v7422 int32
	_ = v7422
	var v7430 int32
	_ = v7430
	var v7431 int32
	_ = v7431
	var v7444 int32
	_ = v7444
	var v7445 int32
	_ = v7445
	var v7447 int32
	_ = v7447
	var v7448 int32
	_ = v7448
	var v7449 int32
	_ = v7449
	var v7457 int32
	_ = v7457
	var v7458 int32
	_ = v7458
	var v7461 int32
	_ = v7461
	var v7468 int32
	_ = v7468
	var v7474 int32
	_ = v7474
	var v7475 int32
	_ = v7475
	var v7478 int32
	_ = v7478
	var v7479 int32
	_ = v7479
	var v7481 int32
	_ = v7481
	var v7482 int32
	_ = v7482
	var v7484 int32
	_ = v7484
	var v7485 int32
	_ = v7485
	var v7487 int32
	_ = v7487
	var v7488 int32
	_ = v7488
	var v7489 int32
	_ = v7489
	var v7493 int32
	_ = v7493
	var v7497 int32
	_ = v7497
	var v7499 int32
	_ = v7499
	var v7501 int32
	_ = v7501
	var v7503 int32
	_ = v7503
	var v7504 int32
	_ = v7504
	var v7507 int32
	_ = v7507
	var v7511 int32
	_ = v7511
	var v7512 int32
	_ = v7512
	var v7514 int32
	_ = v7514
	var v7541 int32
	_ = v7541
	var v7542 int32
	_ = v7542
	var v7547 int32
	_ = v7547
	var v7549 int32
	_ = v7549
	var v7550 int32
	_ = v7550
	var v7555 int32
	_ = v7555
	var v7557 int32
	_ = v7557
	var v7563 int32
	_ = v7563
	var v7566 int32
	_ = v7566
	var v7569 int32
	_ = v7569
	var v7570 int32
	_ = v7570
	var v7571 int32
	_ = v7571
	var v7573 int32
	_ = v7573
	var v7580 int32
	_ = v7580
	var v7581 int32
	_ = v7581
	var v7582 int32
	_ = v7582
	var v7584 int32
	_ = v7584
	var v7590 int32
	_ = v7590
	var v7592 int32
	_ = v7592
	var v7593 int32
	_ = v7593
	var v7594 int32
	_ = v7594
	var v7595 int32
	_ = v7595
	var v7635 int32
	_ = v7635
	var v7637 int32
	_ = v7637
	var v7642 int32
	_ = v7642
	var v7643 int32
	_ = v7643
	var v7644 int32
	_ = v7644
	var v7646 int32
	_ = v7646
	var v7659 int32
	_ = v7659
	var v7662 int32
	_ = v7662
	var v7663 int32
	_ = v7663
	var v7664 int32
	_ = v7664
	var v7666 int32
	_ = v7666
	var v7670 int32
	_ = v7670
	var v7671 int32
	_ = v7671
	var v7672 int32
	_ = v7672
	var v7673 int32
	_ = v7673
	var v7675 int32
	_ = v7675
	var v7677 int32
	_ = v7677
	var v7686 int32
	_ = v7686
	var v7687 int32
	_ = v7687
	var v7692 int32
	_ = v7692
	var v7693 int32
	_ = v7693
	var v7694 int32
	_ = v7694
	var v7696 int32
	_ = v7696
	var v7706 int32
	_ = v7706
	var v7712 int32
	_ = v7712
	var v7715 int32
	_ = v7715
	var v7718 int32
	_ = v7718
	var v7719 int32
	_ = v7719
	var v7725 int32
	_ = v7725
	var v7730 int32
	_ = v7730
	var v7731 int32
	_ = v7731
	var v7732 int32
	_ = v7732
	var v7734 int32
	_ = v7734
	var v7736 int32
	_ = v7736
	var v7737 int32
	_ = v7737
	var v7738 int32
	_ = v7738
	var v7741 int32
	_ = v7741
	var v7742 int32
	_ = v7742
	var v7744 int32
	_ = v7744
	var v7747 int32
	_ = v7747
	var v7748 int32
	_ = v7748
	var v7750 int32
	_ = v7750
	var v7752 int32
	_ = v7752
	var v7754 int32
	_ = v7754
	var v7758 int32
	_ = v7758
	var v7760 int32
	_ = v7760
	var v7762 int32
	_ = v7762
	var v7766 int32
	_ = v7766
	var v7770 int32
	_ = v7770
	var v7771 int32
	_ = v7771
	var v7776 int32
	_ = v7776
	var v7783 int32
	_ = v7783
	var v7801 int32
	_ = v7801
	var v7808 int32
	_ = v7808
	var v7809 int32
	_ = v7809
	var v7810 int32
	_ = v7810
	var v7814 int32
	_ = v7814
	var v7819 int32
	_ = v7819
	var v7820 int32
	_ = v7820
	var v7824 int32
	_ = v7824
	var v7825 int32
	_ = v7825
	var v7827 int32
	_ = v7827
	var v7829 int32
	_ = v7829
	var v7833 int32
	_ = v7833
	var v7836 int32
	_ = v7836
	var v7840 int32
	_ = v7840
	var v7845 int32
	_ = v7845
	var v7849 int32
	_ = v7849
	var v7852 int32
	_ = v7852
	var v7858 int32
	_ = v7858
	var v7863 int32
	_ = v7863
	var v7867 int32
	_ = v7867
	var v7870 int32
	_ = v7870
	var v7874 int32
	_ = v7874
	var v7879 int32
	_ = v7879
	var v7883 int32
	_ = v7883
	var v7886 int32
	_ = v7886
	var v7893 int32
	_ = v7893
	var v7898 int32
	_ = v7898
	var v7902 int32
	_ = v7902
	var v7905 int32
	_ = v7905
	var v7909 int32
	_ = v7909
	var v7914 int32
	_ = v7914
	var v7918 int32
	_ = v7918
	var v7921 int32
	_ = v7921
	var v7925 int32
	_ = v7925
	var v7930 int32
	_ = v7930
	var v7934 int32
	_ = v7934
	var v7937 int32
	_ = v7937
	var v7941 int32
	_ = v7941
	var v7946 int32
	_ = v7946
	var v7950 int32
	_ = v7950
	var v7953 int32
	_ = v7953
	var v7957 int32
	_ = v7957
	var v7962 int32
	_ = v7962
	var v7966 int32
	_ = v7966
	var v7967 int32
	_ = v7967
	var v7973 int32
	_ = v7973
	var v7978 int32
	_ = v7978
	var v7982 int32
	_ = v7982
	var v7989 int32
	_ = v7989
	var v7994 int32
	_ = v7994
	var v7998 int32
	_ = v7998
	var v8005 int32
	_ = v8005
	var v8010 int32
	_ = v8010
	var v8014 int32
	_ = v8014
	var v8021 int32
	_ = v8021
	var v8026 int32
	_ = v8026
	var v8030 int32
	_ = v8030
	var v8037 int32
	_ = v8037
	var v8042 int32
	_ = v8042
	var v8046 int32
	_ = v8046
	var v8053 int32
	_ = v8053
	var v8058 int32
	_ = v8058
	var v8062 int32
	_ = v8062
	var v8065 int32
	_ = v8065
	var v8069 int32
	_ = v8069
	var v8074 int32
	_ = v8074
	var v8078 int32
	_ = v8078
	var v8081 int32
	_ = v8081
	var v8085 int32
	_ = v8085
	var v8090 int32
	_ = v8090
	var v8094 int32
	_ = v8094
	var v8095 int32
	_ = v8095
	var v8101 int32
	_ = v8101
	var v8106 int32
	_ = v8106
	var v8109 int32
	_ = v8109
	var v8113 int32
	_ = v8113
	var v8115 int32
	_ = v8115
	var v8123 int32
	_ = v8123
	var v8128 int32
	_ = v8128
	var v8132 int32
	_ = v8132
	var v8134 int32
	_ = v8134
	var v8142 int32
	_ = v8142
	var v8147 int32
	_ = v8147
	var v8151 int32
	_ = v8151
	var v8153 int32
	_ = v8153
	var v8161 int32
	_ = v8161
	var v8166 int32
	_ = v8166
	var v8170 int32
	_ = v8170
	var v8172 int32
	_ = v8172
	var v8180 int32
	_ = v8180
	var v8185 int32
	_ = v8185
	var v8189 int32
	_ = v8189
	var v8192 int32
	_ = v8192
	var v8203 int32
	_ = v8203
	var v8208 int32
	_ = v8208
	var v8212 int32
	_ = v8212
	var v8214 int32
	_ = v8214
	var v8222 int32
	_ = v8222
	var v8227 int32
	_ = v8227
	var v8231 int32
	_ = v8231
	var v8234 int32
	_ = v8234
	var v8238 int32
	_ = v8238
	var v8243 int32
	_ = v8243
	var v8250 int32
	_ = v8250
	var v8253 int32
	_ = v8253
	var v8257 int32
	_ = v8257
	var v8262 int32
	_ = v8262
	var v8266 int32
	_ = v8266
	var v8269 int32
	_ = v8269
	var v8275 int32
	_ = v8275
	var v8280 int32
	_ = v8280
	var v8282 int32
	_ = v8282
	var v8283 int32
	_ = v8283
	var v8284 int32
	_ = v8284
	var v8287 int32
	_ = v8287
	var v8288 int32
	_ = v8288
	var v8290 int32
	_ = v8290
	var v8292 int32
	_ = v8292
	var v8300 int32
	_ = v8300
	var v8333 int32
	_ = v8333
	var v8337 int32
	_ = v8337
	var v8339 int32
	_ = v8339
	var v8424 int32
	_ = v8424
	var v8426 int32
	_ = v8426
	var v8428 int32
	_ = v8428
	var v8433 int32
	_ = v8433
	var v8435 int32
	_ = v8435
	var v8440 int32
	_ = v8440
	var v8450 int32
	_ = v8450
	var v8454 int32
	_ = v8454
	var v8456 int32
	_ = v8456
	var v8461 int32
	_ = v8461
	var v8462 int32
	_ = v8462
	var v8463 int32
	_ = v8463
	var v8466 int32
	_ = v8466
	var v8469 int32
	_ = v8469
	var v8472 int32
	_ = v8472
	var v8482 int32
	_ = v8482
	var v8486 int32
	_ = v8486
	var v8487 int32
	_ = v8487
	var v8489 int32
	_ = v8489
	var v8494 int32
	_ = v8494
	var v8496 int32
	_ = v8496
	var v8501 int32
	_ = v8501
	var v8503 int32
	_ = v8503
	var v8505 int32
	_ = v8505
	var v8508 int32
	_ = v8508
	var v8513 int32
	_ = v8513
	var v8550 int32
	_ = v8550
	var v8554 int32
	_ = v8554
	var v8555 int32
	_ = v8555
	var v8556 int32
	_ = v8556
	var v8558 int32
	_ = v8558
	var v8561 int32
	_ = v8561
	var v8562 int32
	_ = v8562
	var v8583 int32
	_ = v8583
	var v8644 int32
	_ = v8644
	var v8647 int32
	_ = v8647
	var v8648 int32
	_ = v8648
	var v8656 int32
	_ = v8656
	var v8658 int32
	_ = v8658
	var v8661 int32
	_ = v8661
	var v8667 int32
	_ = v8667
	var v8672 int32
	_ = v8672
	var v8704 int32
	_ = v8704
	var v8708 int32
	_ = v8708
	var v8709 int32
	_ = v8709
	var v8710 int32
	_ = v8710
	var v8713 int32
	_ = v8713
	var v8715 int32
	_ = v8715
	var v8716 int32
	_ = v8716
	var v8717 int32
	_ = v8717
	var v8719 int32
	_ = v8719
	var v8721 int32
	_ = v8721
	var v8725 int32
	_ = v8725
	var v8726 int32
	_ = v8726
	var v8731 int32
	_ = v8731
	var v8732 int32
	_ = v8732
	var v8776 int32
	_ = v8776
	var v8790 int32
	_ = v8790
	var v8820 int32
	_ = v8820
	var v8834 int32
	_ = v8834
	var v8840 int32
	_ = v8840
	var v8860 int32
	_ = v8860
	var v8890 int32
	_ = v8890
	var v8902 int32
	_ = v8902
	var v8905 int32
	_ = v8905
	var v8906 int32
	_ = v8906
	var v8911 int32
	_ = v8911
	var v8915 int32
	_ = v8915
	var v8922 int64
	_ = v8922
	var v8926 int32
	_ = v8926
	var v8928 int32
	_ = v8928
	var v8929 int32
	_ = v8929
	var v8932 int32
	_ = v8932
	var v8936 int32
	_ = v8936
	var v8938 int32
	_ = v8938
	var v8939 int32
	_ = v8939
	var v8944 int32
	_ = v8944
	var v8945 int32
	_ = v8945
	var v8951 int32
	_ = v8951
	var v8959 int32
	_ = v8959
	var v8963 int32
	_ = v8963
	var v8970 int64
	_ = v8970
	var v8974 int32
	_ = v8974
	var v8976 int32
	_ = v8976
	var v8977 int32
	_ = v8977
	var v8980 int32
	_ = v8980
	var v8984 int32
	_ = v8984
	var v8986 int32
	_ = v8986
	var v8987 int32
	_ = v8987
	var v8992 int32
	_ = v8992
	var v8993 int32
	_ = v8993
	var v8999 int32
	_ = v8999
	var v9003 int32
	_ = v9003
	var v9004 int32
	_ = v9004
	var v9005 int32
	_ = v9005
	var v9010 int32
	_ = v9010
	var v9015 int32
	_ = v9015
	var v9016 int32
	_ = v9016
	var v9025 int32
	_ = v9025
	var v9028 int32
	_ = v9028
	var v9031 int32
	_ = v9031
	var v9041 int32
	_ = v9041
	var v9044 int32
	_ = v9044
	var v9048 int32
	_ = v9048
	var v9053 int32
	_ = v9053
	var v9056 int32
	_ = v9056
	var v9058 int32
	_ = v9058
	var v9063 int32
	_ = v9063
	var v9064 int32
	_ = v9064
	var v9070 int32
	_ = v9070
	var v9072 int32
	_ = v9072
	var v9075 int32
	_ = v9075
	var v9076 int32
	_ = v9076
	var v9080 int32
	_ = v9080
	var v9081 int32
	_ = v9081
	var v9082 int32
	_ = v9082
	var v9084 int32
	_ = v9084
	var v9087 int32
	_ = v9087
	var v9089 int32
	_ = v9089
	var v9090 int32
	_ = v9090
	var v9091 int32
	_ = v9091
	var v9100 int32
	_ = v9100
	var v9101 int32
	_ = v9101
	var v9102 int32
	_ = v9102
	var v9103 int32
	_ = v9103
	var v9104 int32
	_ = v9104
	var v9105 int32
	_ = v9105
	var v9109 int32
	_ = v9109
	var v9112 int32
	_ = v9112
	var v9122 int32
	_ = v9122
	var v9125 int32
	_ = v9125
	var v9128 int32
	_ = v9128
	var v9129 int32
	_ = v9129
	var v9131 int32
	_ = v9131
	var v9136 int32
	_ = v9136
	var v9137 int32
	_ = v9137
	var v9138 int32
	_ = v9138
	var v9141 int32
	_ = v9141
	var v9142 int32
	_ = v9142
	var v9144 int32
	_ = v9144
	var v9146 int32
	_ = v9146
	var v9148 int32
	_ = v9148
	var v9150 int32
	_ = v9150
	var v9152 int32
	_ = v9152
	var v9153 int32
	_ = v9153
	var v9154 int32
	_ = v9154
	var v9168 int32
	_ = v9168
	var v9172 int32
	_ = v9172
	var v9173 int32
	_ = v9173
	var v9175 int32
	_ = v9175
	var v9176 int32
	_ = v9176
	var v9179 int32
	_ = v9179
	var v9180 int32
	_ = v9180
	var v9181 int32
	_ = v9181
	var v9182 int32
	_ = v9182
	var v9185 int32
	_ = v9185
	var v9190 int32
	_ = v9190
	var v9197 int32
	_ = v9197
	var v9198 int32
	_ = v9198
	var v9199 int32
	_ = v9199
	var v9209 int32
	_ = v9209
	var v9210 int32
	_ = v9210
	var v9211 int32
	_ = v9211
	var v9213 int32
	_ = v9213
	var v9216 int32
	_ = v9216
	var v9217 int32
	_ = v9217
	var v9218 int32
	_ = v9218
	var v9227 int32
	_ = v9227
	var v9228 int32
	_ = v9228
	var v9236 int32
	_ = v9236
	var v9239 int32
	_ = v9239
	var v9241 int32
	_ = v9241
	var v9244 int32
	_ = v9244
	var v9245 int32
	_ = v9245
	var v9251 int32
	_ = v9251
	var v9254 int32
	_ = v9254
	var v9256 int32
	_ = v9256
	var v9258 int32
	_ = v9258
	var v9262 int32
	_ = v9262
	var v9267 int32
	_ = v9267
	var v9269 int32
	_ = v9269
	var v9271 int32
	_ = v9271
	var v9276 int32
	_ = v9276
	var v9278 int32
	_ = v9278
	var v9280 int32
	_ = v9280
	var v9281 int32
	_ = v9281
	var v9296 int32
	_ = v9296
	var v9325 int32
	_ = v9325
	var v9328 int32
	_ = v9328
	var v9330 int32
	_ = v9330
	var v9332 int32
	_ = v9332
	var v9334 int32
	_ = v9334
	var v9337 int32
	_ = v9337
	var v9381 int32
	_ = v9381
	var v9384 int32
	_ = v9384
	var v9385 int32
	_ = v9385
	var v9387 int32
	_ = v9387
	var v9391 int32
	_ = v9391
	var v9393 int32
	_ = v9393
	var v9409 int32
	_ = v9409
	var v9435 int32
	_ = v9435
	var v9438 int32
	_ = v9438
	var v9439 int32
	_ = v9439
	var v9444 int32
	_ = v9444
	var v9445 int32
	_ = v9445
	var v9453 int32
	_ = v9453
	var v9455 int32
	_ = v9455
	var v9459 int32
	_ = v9459
	var v9460 int32
	_ = v9460
	var v9471 int32
	_ = v9471
	var v9473 int32
	_ = v9473
	var v9474 int32
	_ = v9474
	var v9475 int32
	_ = v9475
	var v9481 int32
	_ = v9481
	var v9486 int32
	_ = v9486
	var v9518 int32
	_ = v9518
	var v9522 int32
	_ = v9522
	var v9523 int32
	_ = v9523
	var v9524 int32
	_ = v9524
	var v9527 int32
	_ = v9527
	var v9529 int32
	_ = v9529
	var v9530 int32
	_ = v9530
	var v9531 int32
	_ = v9531
	var v9533 int32
	_ = v9533
	var v9535 int32
	_ = v9535
	var v9537 int32
	_ = v9537
	var v9538 int32
	_ = v9538
	var v9543 int32
	_ = v9543
	var v9544 int32
	_ = v9544
	var v9568 int32
	_ = v9568
	var v9587 int32
	_ = v9587
	var v9629 int32
	_ = v9629
	var v9676 int32
	_ = v9676
	var v9680 int32
	_ = v9680
	var v9684 int32
	_ = v9684
	var v9688 int64
	_ = v9688
	var v9691 int32
	_ = v9691
	var v9692 int32
	_ = v9692
	var v9693 int32
	_ = v9693
	var v9694 int32
	_ = v9694
	var v9695 int32
	_ = v9695
	var v9697 int32
	_ = v9697
	var v9698 int32
	_ = v9698
	var v9703 int32
	_ = v9703
	var v9704 int32
	_ = v9704
	var v9709 int32
	_ = v9709
	var v9750 int32
	_ = v9750
	var v9751 int32
	_ = v9751
	var v9754 int32
	_ = v9754
	var v9773 int32
	_ = v9773
	var v9798 int32
	_ = v9798
	var v9804 int32
	_ = v9804
	var v9809 int32
	_ = v9809
	var v9819 int32
	_ = v9819
	var v9824 int32
	_ = v9824
	var v9825 int32
	_ = v9825
	var v9826 int32
	_ = v9826
	var v9829 int32
	_ = v9829
	var v9835 int32
	_ = v9835
	var v9840 int32
	_ = v9840
	var v9843 int32
	_ = v9843
	var v9844 int32
	_ = v9844
	var v9846 int32
	_ = v9846
	var v9849 int32
	_ = v9849
	var v9854 int32
	_ = v9854
	var v9856 int32
	_ = v9856
	var v9861 int32
	_ = v9861
	var v9862 int32
	_ = v9862
	var v9864 int32
	_ = v9864
	var v9865 int32
	_ = v9865
	var v9866 int32
	_ = v9866
	var v9867 int32
	_ = v9867
	var v9871 int32
	_ = v9871
	var v9874 int32
	_ = v9874
	var v9884 int32
	_ = v9884
	var v9888 int32
	_ = v9888
	var v9889 int32
	_ = v9889
	var v9891 int32
	_ = v9891
	var v9896 int32
	_ = v9896
	var v9898 int32
	_ = v9898
	var v9903 int32
	_ = v9903
	var v9905 int32
	_ = v9905
	var v9906 int32
	_ = v9906
	var v9909 int32
	_ = v9909
	var v9910 int32
	_ = v9910
	var v9912 int32
	_ = v9912
	var v9913 int32
	_ = v9913
	var v9920 int32
	_ = v9920
	var v9923 int32
	_ = v9923
	var v9926 int32
	_ = v9926
	var v9931 int32
	_ = v9931
	var v9932 int32
	_ = v9932
	var v9933 int32
	_ = v9933
	var v9934 int32
	_ = v9934
	var v9937 int32
	_ = v9937
	var v9938 int32
	_ = v9938
	var v9942 int32
	_ = v9942
	var v9949 int32
	_ = v9949
	var v9950 int32
	_ = v9950
	var v9951 int32
	_ = v9951
	var v9952 int32
	_ = v9952
	var v9954 int32
	_ = v9954
	var v9959 int32
	_ = v9959
	var v9960 int32
	_ = v9960
	var v9962 int32
	_ = v9962
	var v9963 int32
	_ = v9963
	var v9966 int32
	_ = v9966
	var v9967 int32
	_ = v9967
	var v9968 int32
	_ = v9968
	var v9970 int32
	_ = v9970
	var v9971 int32
	_ = v9971
	var v9973 int32
	_ = v9973
	var v9977 int32
	_ = v9977
	var v9981 int32
	_ = v9981
	var v9982 int32
	_ = v9982
	var v9987 int32
	_ = v9987
	var v9994 int32
	_ = v9994
	var v10006 int32
	_ = v10006
	var v10007 int32
	_ = v10007
	var v10008 int32
	_ = v10008
	var v10013 int32
	_ = v10013
	var v10015 int32
	_ = v10015
	var v10017 int32
	_ = v10017
	var v10020 int32
	_ = v10020
	var v10022 int32
	_ = v10022
	var v10028 int32
	_ = v10028
	var v10030 int32
	_ = v10030
	var v10035 int32
	_ = v10035
	var v10039 int32
	_ = v10039
	var v10040 int32
	_ = v10040
	var v10045 int32
	_ = v10045
	var v10046 int32
	_ = v10046
	var v10056 int32
	_ = v10056
	var v10060 int32
	_ = v10060
	var v10061 int32
	_ = v10061
	var v10064 int32
	_ = v10064
	var v10067 int32
	_ = v10067
	var v10076 int32
	_ = v10076
	var v10079 int32
	_ = v10079
	var v10081 int32
	_ = v10081
	var v10085 int32
	_ = v10085
	var v10089 int32
	_ = v10089
	var v10094 int32
	_ = v10094
	var v10098 int32
	_ = v10098
	var v10102 int64
	_ = v10102
	var v10105 int32
	_ = v10105
	var v10107 int32
	_ = v10107
	var v10108 int32
	_ = v10108
	var v10109 int32
	_ = v10109
	var v10110 int32
	_ = v10110
	var v10111 int32
	_ = v10111
	var v10114 int32
	_ = v10114
	var v10115 int32
	_ = v10115
	var v10116 int32
	_ = v10116
	var v10118 int32
	_ = v10118
	var v10119 int32
	_ = v10119
	var v10122 int32
	_ = v10122
	var v10128 int32
	_ = v10128
	var v10133 int32
	_ = v10133
	var v10135 int32
	_ = v10135
	var v10137 int32
	_ = v10137
	var v10138 int32
	_ = v10138
	var v10139 int32
	_ = v10139
	var v10141 int32
	_ = v10141
	var v10144 int32
	_ = v10144
	var v10146 int32
	_ = v10146
	var v10150 int32
	_ = v10150
	var v10153 int32
	_ = v10153
	var v10156 int32
	_ = v10156
	var v10162 int32
	_ = v10162
	var v10200 int32
	_ = v10200
	var v10201 int64
	_ = v10201
	var v10205 int32
	_ = v10205
	var v10210 int32
	_ = v10210
	var v10214 int32
	_ = v10214
	var v10221 int64
	_ = v10221
	var v10225 int32
	_ = v10225
	var v10227 int32
	_ = v10227
	var v10228 int32
	_ = v10228
	var v10231 int32
	_ = v10231
	var v10235 int32
	_ = v10235
	var v10237 int32
	_ = v10237
	var v10238 int32
	_ = v10238
	var v10243 int32
	_ = v10243
	var v10244 int32
	_ = v10244
	var v10250 int32
	_ = v10250
	var v10295 int32
	_ = v10295
	var v10305 int32
	_ = v10305
	var v10309 int32
	_ = v10309
	var v10312 int32
	_ = v10312
	var v10317 int32
	_ = v10317
	var v10318 int32
	_ = v10318
	var v10323 int32
	_ = v10323
	var v10324 int32
	_ = v10324
	var v10329 int32
	_ = v10329
	var v10370 int32
	_ = v10370
	var v10371 int32
	_ = v10371
	var v10374 int32
	_ = v10374
	var v10379 int32
	_ = v10379
	var v10418 int32
	_ = v10418
	var v10419 int32
	_ = v10419
	var v10424 int32
	_ = v10424
	var v10427 int32
	_ = v10427
	var v10428 int32
	_ = v10428
	var v10435 int32
	_ = v10435
	var v10438 int32
	_ = v10438
	var v10441 int32
	_ = v10441
	var v10444 int32
	_ = v10444
	var v10451 int32
	_ = v10451
	var v10454 int32
	_ = v10454
	var v10456 int32
	_ = v10456
	var v10457 int32
	_ = v10457
	var v10458 int32
	_ = v10458
	var v10460 int32
	_ = v10460
	var v10461 int32
	_ = v10461
	var v10462 int32
	_ = v10462
	var v10463 int32
	_ = v10463
	var v10464 int32
	_ = v10464
	var v10466 int32
	_ = v10466
	var v10468 int32
	_ = v10468
	var v10469 int32
	_ = v10469
	var v10470 int32
	_ = v10470
	var v10471 int32
	_ = v10471
	var v10472 int32
	_ = v10472
	var v10473 int32
	_ = v10473
	var v10474 int32
	_ = v10474
	var v10477 int32
	_ = v10477
	var v10478 int32
	_ = v10478
	var v10484 int32
	_ = v10484
	var v10485 int32
	_ = v10485
	var v10489 int32
	_ = v10489
	var v10492 int32
	_ = v10492
	var v10493 int32
	_ = v10493
	var v10495 int32
	_ = v10495
	var v10496 int32
	_ = v10496
	var v10497 int32
	_ = v10497
	var v10499 int32
	_ = v10499
	var v10500 int32
	_ = v10500
	var v10504 int32
	_ = v10504
	var v10505 int32
	_ = v10505
	var v10516 int32
	_ = v10516
	var v10517 int32
	_ = v10517
	var v10520 int32
	_ = v10520
	var v10527 int32
	_ = v10527
	var v10535 int32
	_ = v10535
	var v10563 int32
	_ = v10563
	var v10564 int32
	_ = v10564
	var v10566 int32
	_ = v10566
	var v10573 int32
	_ = v10573
	var v10574 int32
	_ = v10574
	var v10576 int32
	_ = v10576
	var v10577 int32
	_ = v10577
	var v10581 int32
	_ = v10581
	var v10582 int32
	_ = v10582
	var v10583 int32
	_ = v10583
	var v10584 int32
	_ = v10584
	var v10585 int32
	_ = v10585
	var v10590 int32
	_ = v10590
	var v10592 int32
	_ = v10592
	var v10599 int32
	_ = v10599
	var v10600 int32
	_ = v10600
	var v10607 int32
	_ = v10607
	var v10609 int32
	_ = v10609
	var v10610 int32
	_ = v10610
	var v10611 int32
	_ = v10611
	var v10612 int32
	_ = v10612
	var v10614 int32
	_ = v10614
	var v10615 int32
	_ = v10615
	var v10617 int64
	_ = v10617
	var v10618 int32
	_ = v10618
	var v10619 int32
	_ = v10619
	var v10624 int32
	_ = v10624
	var v10625 int32
	_ = v10625
	var v10626 int32
	_ = v10626
	var v10629 int32
	_ = v10629
	var v10634 int32
	_ = v10634
	var v10635 int32
	_ = v10635
	var v10637 int32
	_ = v10637
	var v10638 int32
	_ = v10638
	var v10639 int32
	_ = v10639
	var v10642 int32
	_ = v10642
	var v10643 int32
	_ = v10643
	var v10646 int32
	_ = v10646
	var v10647 int32
	_ = v10647
	var v10648 int32
	_ = v10648
	var v10655 int32
	_ = v10655
	var v10657 int32
	_ = v10657
	var v10660 int32
	_ = v10660
	var v10666 int32
	_ = v10666
	var v10667 int32
	_ = v10667
	var v10671 int32
	_ = v10671
	var v10672 int32
	_ = v10672
	var v10674 int64
	_ = v10674
	var v10675 int32
	_ = v10675
	var v10676 int32
	_ = v10676
	var v10677 int32
	_ = v10677
	var v10681 int32
	_ = v10681
	var v10684 int64
	_ = v10684
	var v10687 int32
	_ = v10687
	var v10692 int32
	_ = v10692
	var v10694 int32
	_ = v10694
	var v10699 int32
	_ = v10699
	var v10701 int32
	_ = v10701
	var v10703 int32
	_ = v10703
	var v10704 int32
	_ = v10704
	var v10707 int32
	_ = v10707
	var v10710 int32
	_ = v10710
	var v10711 int32
	_ = v10711
	var v10740 int32
	_ = v10740
	var v10780 int32
	_ = v10780
	var v10792 int32
	_ = v10792
	var v10795 int32
	_ = v10795
	var v10798 int32
	_ = v10798
	var v10799 int32
	_ = v10799
	var v10814 int32
	_ = v10814
	var v10815 int32
	_ = v10815
	var v10820 int32
	_ = v10820
	var v10821 int32
	_ = v10821
	var v10826 int32
	_ = v10826
	var v10867 int32
	_ = v10867
	var v10868 int32
	_ = v10868
	var v10871 int32
	_ = v10871
	var v10890 int32
	_ = v10890
	var v10915 int32
	_ = v10915
	var v10916 int32
	_ = v10916
	var v10918 int32
	_ = v10918
	var v10919 int32
	_ = v10919
	var v10920 int32
	_ = v10920
	var v10921 int32
	_ = v10921
	var v10932 int32
	_ = v10932
	var v10935 int32
	_ = v10935
	var v10938 int32
	_ = v10938
	var v10944 int32
	_ = v10944
	var v10982 int32
	_ = v10982
	var v10983 int64
	_ = v10983
	var v10987 int32
	_ = v10987
	var v10992 int32
	_ = v10992
	var v10996 int32
	_ = v10996
	var v11003 int64
	_ = v11003
	var v11007 int32
	_ = v11007
	var v11009 int32
	_ = v11009
	var v11010 int32
	_ = v11010
	var v11013 int32
	_ = v11013
	var v11017 int32
	_ = v11017
	var v11019 int32
	_ = v11019
	var v11020 int32
	_ = v11020
	var v11025 int32
	_ = v11025
	var v11026 int32
	_ = v11026
	var v11032 int32
	_ = v11032
	var v11076 int32
	_ = v11076
	var v11077 int32
	_ = v11077
	var v11080 int32
	_ = v11080
	var v11082 int32
	_ = v11082
	var v11083 int32
	_ = v11083
	var v11085 int32
	_ = v11085
	var v11086 int32
	_ = v11086
	var v11089 int32
	_ = v11089
	var v11094 int32
	_ = v11094
	var v11098 int32
	_ = v11098
	var v11099 int32
	_ = v11099
	var v11104 int32
	_ = v11104
	var v11105 int32
	_ = v11105
	var v11115 int32
	_ = v11115
	var v11117 int32
	_ = v11117
	var v11121 int32
	_ = v11121
	var v11122 int32
	_ = v11122
	var v11125 int32
	_ = v11125
	var v11126 int32
	_ = v11126
	var v11127 int32
	_ = v11127
	var v11130 int32
	_ = v11130
	var v11134 int32
	_ = v11134
	var v11137 int32
	_ = v11137
	var v11146 int32
	_ = v11146
	var v11148 int32
	_ = v11148
	var v11149 int32
	_ = v11149
	var v11152 int32
	_ = v11152
	var v11156 int32
	_ = v11156
	var v11160 int32
	_ = v11160
	var v11161 int32
	_ = v11161
	var v11164 int32
	_ = v11164
	var v11171 int32
	_ = v11171
	var v11172 int32
	_ = v11172
	var v11175 int32
	_ = v11175
	var v11179 int32
	_ = v11179
	var v11187 int32
	_ = v11187
	var v11192 int32
	_ = v11192
	var v11196 int32
	_ = v11196
	var v11200 int64
	_ = v11200
	var v11203 int32
	_ = v11203
	var v11204 int32
	_ = v11204
	var v11205 int32
	_ = v11205
	var v11207 int32
	_ = v11207
	var v11208 int32
	_ = v11208
	var v11210 int32
	_ = v11210
	var v11212 int32
	_ = v11212
	var v11214 int32
	_ = v11214
	var v11215 int32
	_ = v11215
	var v11216 int32
	_ = v11216
	var v11222 int32
	_ = v11222
	var v11223 int32
	_ = v11223
	var v11227 int32
	_ = v11227
	var v11228 int32
	_ = v11228
	var v11231 int32
	_ = v11231
	var v11234 int32
	_ = v11234
	var v11235 int32
	_ = v11235
	var v11236 int32
	_ = v11236
	var v11240 int32
	_ = v11240
	var v11241 int32
	_ = v11241
	var v11247 int32
	_ = v11247
	var v11248 int32
	_ = v11248
	var v11249 int32
	_ = v11249
	var v11250 int32
	_ = v11250
	var v11251 int32
	_ = v11251
	var v11252 int32
	_ = v11252
	var v11253 int32
	_ = v11253
	var v11255 int32
	_ = v11255
	var v11256 int32
	_ = v11256
	var v11261 int32
	_ = v11261
	var v11264 int32
	_ = v11264
	var v11267 int32
	_ = v11267
	var v11270 int32
	_ = v11270
	var v11271 int32
	_ = v11271
	var v11277 int32
	_ = v11277
	var v11315 int32
	_ = v11315
	var v11316 int64
	_ = v11316
	var v11320 int32
	_ = v11320
	var v11325 int32
	_ = v11325
	var v11329 int32
	_ = v11329
	var v11336 int64
	_ = v11336
	var v11340 int32
	_ = v11340
	var v11342 int32
	_ = v11342
	var v11343 int32
	_ = v11343
	var v11346 int32
	_ = v11346
	var v11350 int32
	_ = v11350
	var v11352 int32
	_ = v11352
	var v11353 int32
	_ = v11353
	var v11358 int32
	_ = v11358
	var v11359 int32
	_ = v11359
	var v11365 int32
	_ = v11365
	var v11369 int32
	_ = v11369
	var v11372 int32
	_ = v11372
	var v11373 int32
	_ = v11373
	var v11376 int32
	_ = v11376
	var v11377 int32
	_ = v11377
	var v11379 int32
	_ = v11379
	var v11380 int32
	_ = v11380
	var v11383 int32
	_ = v11383
	var v11389 int32
	_ = v11389
	var v11427 int32
	_ = v11427
	var v11428 int64
	_ = v11428
	var v11432 int32
	_ = v11432
	var v11437 int32
	_ = v11437
	var v11441 int32
	_ = v11441
	var v11448 int64
	_ = v11448
	var v11452 int32
	_ = v11452
	var v11454 int32
	_ = v11454
	var v11455 int32
	_ = v11455
	var v11458 int32
	_ = v11458
	var v11462 int32
	_ = v11462
	var v11464 int32
	_ = v11464
	var v11465 int32
	_ = v11465
	var v11470 int32
	_ = v11470
	var v11471 int32
	_ = v11471
	var v11477 int32
	_ = v11477
	var v11520 int32
	_ = v11520
	var v11525 int32
	_ = v11525
	var v11531 int32
	_ = v11531
	var v11541 int32
	_ = v11541
	var v11545 int32
	_ = v11545
	var v11546 int32
	_ = v11546
	var v11551 int32
	_ = v11551
	var v11552 int32
	_ = v11552
	var v11553 int32
	_ = v11553
	var v11555 int32
	_ = v11555
	var v11556 int32
	_ = v11556
	var v11559 int32
	_ = v11559
	var v11566 int32
	_ = v11566
	var v11601 int32
	_ = v11601
	var v11605 int32
	_ = v11605
	var v11606 int32
	_ = v11606
	var v11607 int32
	_ = v11607
	var v11609 int32
	_ = v11609
	var v11612 int32
	_ = v11612
	var v11613 int32
	_ = v11613
	var v11657 int32
	_ = v11657
	var v11658 int32
	_ = v11658
	var v11662 int32
	_ = v11662
	var v11666 int32
	_ = v11666
	var v11670 int32
	_ = v11670
	var v11672 int32
	_ = v11672
	var v11677 int32
	_ = v11677
	var v11683 int32
	_ = v11683
	var v11685 int32
	_ = v11685
	var v11688 int32
	_ = v11688
	var v11692 int32
	_ = v11692
	var v11696 int32
	_ = v11696
	var v11697 int32
	_ = v11697
	var v11700 int32
	_ = v11700
	var v11707 int32
	_ = v11707
	var v11708 int32
	_ = v11708
	var v11714 int32
	_ = v11714
	var v11719 int32
	_ = v11719
	var v11755 int32
	_ = v11755
	var v11756 int32
	_ = v11756
	var v11763 int32
	_ = v11763
	var v11766 int32
	_ = v11766
	var v11769 int32
	_ = v11769
	var v11770 int32
	_ = v11770
	var v11771 int32
	_ = v11771
	var v11774 int32
	_ = v11774
	var v11777 int32
	_ = v11777
	var v11780 int32
	_ = v11780
	var v11787 int32
	_ = v11787
	var v11789 int32
	_ = v11789
	var v11790 int32
	_ = v11790
	var v11793 int32
	_ = v11793
	var v11794 int32
	_ = v11794
	var v11808 int32
	_ = v11808
	var v11812 int32
	_ = v11812
	var v11813 int32
	_ = v11813
	var v11814 int32
	_ = v11814
	var v11816 int32
	_ = v11816
	var v11817 int32
	_ = v11817
	var v11819 int32
	_ = v11819
	var v11820 int32
	_ = v11820
	var v11825 int32
	_ = v11825
	var v11833 int32
	_ = v11833
	var v11836 int32
	_ = v11836
	var v11837 int32
	_ = v11837
	var v11839 int32
	_ = v11839
	var v11843 int32
	_ = v11843
	var v11845 int32
	_ = v11845
	var v11848 int32
	_ = v11848
	var v11849 int32
	_ = v11849
	var v11851 int32
	_ = v11851
	var v11858 int32
	_ = v11858
	var v11863 int32
	_ = v11863
	var v11864 int32
	_ = v11864
	var v11868 int32
	_ = v11868
	var v11870 int32
	_ = v11870
	var v11875 int32
	_ = v11875
	var v11876 int32
	_ = v11876
	var v11878 int32
	_ = v11878
	var v11882 int32
	_ = v11882
	var v11885 int32
	_ = v11885
	var v11886 int32
	_ = v11886
	var v11891 int32
	_ = v11891
	var v11892 int32
	_ = v11892
	var v11902 int32
	_ = v11902
	var v11904 int32
	_ = v11904
	var v11908 int32
	_ = v11908
	var v11909 int32
	_ = v11909
	var v11912 int32
	_ = v11912
	var v11915 int32
	_ = v11915
	var v11920 int32
	_ = v11920
	var v11926 int32
	_ = v11926
	var v11935 int32
	_ = v11935
	var v11937 int32
	_ = v11937
	var v11938 int32
	_ = v11938
	var v11941 int32
	_ = v11941
	var v11945 int32
	_ = v11945
	var v11949 int32
	_ = v11949
	var v11950 int32
	_ = v11950
	var v11953 int32
	_ = v11953
	var v11960 int32
	_ = v11960
	var v11961 int32
	_ = v11961
	var v11963 int32
	_ = v11963
	var v11967 int32
	_ = v11967
	var v11974 int32
	_ = v11974
	var v11979 int32
	_ = v11979
	var v11983 int32
	_ = v11983
	var v11987 int64
	_ = v11987
	var v11993 int32
	_ = v11993
	var v11996 int32
	_ = v11996
	var v11999 int32
	_ = v11999
	var v12000 int32
	_ = v12000
	var v12002 int32
	_ = v12002
	var v12005 int32
	_ = v12005
	var v12006 int32
	_ = v12006
	var v12015 int32
	_ = v12015
	var v12016 int32
	_ = v12016
	var v12018 int32
	_ = v12018
	var v12020 int32
	_ = v12020
	var v12021 int32
	_ = v12021
	var v12029 int32
	_ = v12029
	var v12030 int32
	_ = v12030
	var v12033 int32
	_ = v12033
	var v12034 int32
	_ = v12034
	var v12035 int32
	_ = v12035
	var v12036 int32
	_ = v12036
	var v12039 int32
	_ = v12039
	var v12042 int32
	_ = v12042
	var v12045 int32
	_ = v12045
	var v12047 int32
	_ = v12047
	var v12050 int32
	_ = v12050
	var v12051 int32
	_ = v12051
	var v12053 int32
	_ = v12053
	var v12058 int32
	_ = v12058
	var v12060 int32
	_ = v12060
	var v12062 int32
	_ = v12062
	var v12069 int32
	_ = v12069
	var v12073 int32
	_ = v12073
	var v12085 int32
	_ = v12085
	var v12086 int32
	_ = v12086
	var v12087 int32
	_ = v12087
	var v12089 int32
	_ = v12089
	var v12093 int32
	_ = v12093
	var v12094 int32
	_ = v12094
	var v12096 int32
	_ = v12096
	var v12097 int32
	_ = v12097
	var v12098 int32
	_ = v12098
	var v12100 int32
	_ = v12100
	var v12106 int32
	_ = v12106
	var v12107 int32
	_ = v12107
	var v12108 int32
	_ = v12108
	var v12109 int32
	_ = v12109
	var v12112 int32
	_ = v12112
	var v12119 int32
	_ = v12119
	var v12120 int32
	_ = v12120
	var v12121 int32
	_ = v12121
	var v12124 int32
	_ = v12124
	var v12127 int32
	_ = v12127
	var v12132 int32
	_ = v12132
	var v12133 int32
	_ = v12133
	var v12135 int32
	_ = v12135
	var v12137 int32
	_ = v12137
	var v12141 int32
	_ = v12141
	var v12142 int32
	_ = v12142
	var v12143 int32
	_ = v12143
	var v12148 int32
	_ = v12148
	var v12149 int32
	_ = v12149
	var v12150 int32
	_ = v12150
	var v12153 int32
	_ = v12153
	var v12154 int32
	_ = v12154
	var v12155 int32
	_ = v12155
	var v12157 int32
	_ = v12157
	var v12161 int32
	_ = v12161
	var v12162 int32
	_ = v12162
	var v12164 int32
	_ = v12164
	var v12166 int32
	_ = v12166
	var v12168 int32
	_ = v12168
	var v12169 int32
	_ = v12169
	var v12172 int32
	_ = v12172
	var v12179 int32
	_ = v12179
	var v12183 int32
	_ = v12183
	var v12185 int32
	_ = v12185
	var v12188 int32
	_ = v12188
	var v12193 int32
	_ = v12193
	var v12194 int32
	_ = v12194
	var v12203 int32
	_ = v12203
	var v12208 int32
	_ = v12208
	var v12210 int32
	_ = v12210
	var v12212 int32
	_ = v12212
	var v12214 int32
	_ = v12214
	var v12215 int32
	_ = v12215
	var v12217 int32
	_ = v12217
	var v12218 int32
	_ = v12218
	var v12219 int32
	_ = v12219
	var v12221 int32
	_ = v12221
	var v12223 int32
	_ = v12223
	var v12224 int32
	_ = v12224
	var v12226 int32
	_ = v12226
	var v12227 int32
	_ = v12227
	var v12230 int32
	_ = v12230
	var v12232 int32
	_ = v12232
	var v12233 int32
	_ = v12233
	var v12235 int32
	_ = v12235
	var v12236 int32
	_ = v12236
	var v12238 int32
	_ = v12238
	var v12241 int32
	_ = v12241
	var v12243 int32
	_ = v12243
	var v12244 int64
	_ = v12244
	var v12250 int32
	_ = v12250
	var v12251 int32
	_ = v12251
	var v12257 int32
	_ = v12257
	var v12258 int32
	_ = v12258
	var v12276 int32
	_ = v12276
	var v12302 int32
	_ = v12302
	var v12303 int32
	_ = v12303
	var v12306 int32
	_ = v12306
	var v12312 int32
	_ = v12312
	var v12348 int32
	_ = v12348
	var v12349 int32
	_ = v12349
	var v12352 int32
	_ = v12352
	var v12362 int32
	_ = v12362
	var v12366 int32
	_ = v12366
	var v12387 int32
	_ = v12387
	var v12409 int32
	_ = v12409
	var v12410 int32
	_ = v12410
	var v12413 int32
	_ = v12413
	var v12415 int32
	_ = v12415
	var v12416 int32
	_ = v12416
	var v12419 int32
	_ = v12419
	var v12421 int32
	_ = v12421
	var v12426 int32
	_ = v12426
	var v12427 int32
	_ = v12427
	var v12428 int32
	_ = v12428
	var v12434 int32
	_ = v12434
	var v12435 int32
	_ = v12435
	var v12437 int32
	_ = v12437
	var v12446 int32
	_ = v12446
	var v12447 int32
	_ = v12447
	var v12452 int32
	_ = v12452
	var v12458 int32
	_ = v12458
	var v12462 int32
	_ = v12462
	var v12463 int32
	_ = v12463
	var v12464 int32
	_ = v12464
	var v12465 int32
	_ = v12465
	var v12467 int32
	_ = v12467
	var v12468 int32
	_ = v12468
	var v12470 int64
	_ = v12470
	var v12471 int32
	_ = v12471
	var v12475 int32
	_ = v12475
	var v12478 int32
	_ = v12478
	var v12482 int32
	_ = v12482
	var v12488 int32
	_ = v12488
	var v12490 int32
	_ = v12490
	var v12495 int32
	_ = v12495
	var v12496 int32
	_ = v12496
	var v12497 int32
	_ = v12497
	var v12499 int64
	_ = v12499
	var v12500 int32
	_ = v12500
	var v12502 int32
	_ = v12502
	var v12503 int32
	_ = v12503
	var v12507 int32
	_ = v12507
	var v12549 int32
	_ = v12549
	var v12550 int32
	_ = v12550
	var v12552 int32
	_ = v12552
	var v12553 int32
	_ = v12553
	var v12556 int32
	_ = v12556
	var v12577 int32
	_ = v12577
	var v12604 int32
	_ = v12604
	var v12608 int32
	_ = v12608
	var v12610 int32
	_ = v12610
	var v12616 int32
	_ = v12616
	var v12619 int32
	_ = v12619
	var v12623 int32
	_ = v12623
	var v12628 int32
	_ = v12628
	var v12632 int32
	_ = v12632
	var v12635 int32
	_ = v12635
	var v12639 int32
	_ = v12639
	var v12644 int32
	_ = v12644
	var v12648 int32
	_ = v12648
	var v12651 int32
	_ = v12651
	var v12659 int32
	_ = v12659
	var v12664 int32
	_ = v12664
	var v12668 int32
	_ = v12668
	var v12678 int32
	_ = v12678
	var v12683 int32
	_ = v12683
	var v12687 int32
	_ = v12687
	var v12690 int32
	_ = v12690
	var v12692 int32
	_ = v12692
	var v12698 int32
	_ = v12698
	var v12703 int32
	_ = v12703
	var v12707 int32
	_ = v12707
	var v12710 int32
	_ = v12710
	var v12717 int32
	_ = v12717
	var v12722 int32
	_ = v12722
	var v12726 int32
	_ = v12726
	var v12729 int32
	_ = v12729
	var v12735 int32
	_ = v12735
	var v12740 int32
	_ = v12740
	var v12744 int32
	_ = v12744
	var v12747 int32
	_ = v12747
	var v12755 int32
	_ = v12755
	var v12760 int32
	_ = v12760
	var v12764 int32
	_ = v12764
	var v12767 int32
	_ = v12767
	var v12774 int32
	_ = v12774
	var v12779 int32
	_ = v12779
	var v12821 int32
	_ = v12821
	var v12822 int64
	_ = v12822
	var v12823 int32
	_ = v12823
	var v12863 int64
	_ = v12863
	var v12865 int32
	_ = v12865
	var v12867 int32
	_ = v12867
	var v12868 int32
	_ = v12868
	var v12869 int32
	_ = v12869
	var v12871 int32
	_ = v12871
	var v12874 int32
	_ = v12874
	var v12879 int32
	_ = v12879
	var v12880 int32
	_ = v12880
	var v12881 int32
	_ = v12881
	var v12895 int32
	_ = v12895
	var v12896 int32
	_ = v12896
	var v12897 int32
	_ = v12897
	var v12899 int32
	_ = v12899
	var v12902 int32
	_ = v12902
	var v12904 int32
	_ = v12904
	var v12907 int32
	_ = v12907
	var v12908 int64
	_ = v12908
	var v12912 int32
	_ = v12912
	var v12916 int32
	_ = v12916
	var v12920 int32
	_ = v12920
	var v12921 int64
	_ = v12921
	var v12922 int32
	_ = v12922
	var v12923 int32
	_ = v12923
	var v12926 int32
	_ = v12926
	var v12927 int32
	_ = v12927
	var v12930 int32
	_ = v12930
	var v12931 int32
	_ = v12931
	var v12932 int32
	_ = v12932
	var v12939 int32
	_ = v12939
	var v12940 int32
	_ = v12940
	var v12944 int32
	_ = v12944
	var v12949 int32
	_ = v12949
	var v12950 int32
	_ = v12950
	var v12952 int32
	_ = v12952
	var v12955 int32
	_ = v12955
	var v12956 int32
	_ = v12956
	var v12957 int32
	_ = v12957
	var v12959 int32
	_ = v12959
	var v12961 int32
	_ = v12961
	var v12962 int32
	_ = v12962
	var v12963 int32
	_ = v12963
	var v12978 int32
	_ = v12978
	var v12984 int32
	_ = v12984
	var v12986 int32
	_ = v12986
	var v12990 int32
	_ = v12990
	var v12993 int32
	_ = v12993
	var v13000 int32
	_ = v13000
	var v13005 int32
	_ = v13005
	var v13011 int32
	_ = v13011
	var v13014 int32
	_ = v13014
	var v13015 int32
	_ = v13015
	var v13016 int32
	_ = v13016
	var v13017 int32
	_ = v13017
	var v13019 int32
	_ = v13019
	var v13021 int32
	_ = v13021
	var v13028 int32
	_ = v13028
	var v13030 int32
	_ = v13030
	var v13032 int32
	_ = v13032
	var v13036 int32
	_ = v13036
	var v13037 int32
	_ = v13037
	var v13042 int32
	_ = v13042
	var v13043 int32
	_ = v13043
	var v13053 int32
	_ = v13053
	var v13057 int32
	_ = v13057
	var v13058 int32
	_ = v13058
	var v13070 int32
	_ = v13070
	var v13072 int32
	_ = v13072
	var v13075 int32
	_ = v13075
	var v13082 int32
	_ = v13082
	var v13085 int32
	_ = v13085
	var v13087 int32
	_ = v13087
	var v13089 int32
	_ = v13089
	var v13091 int32
	_ = v13091
	var v13094 int32
	_ = v13094
	var v13097 int32
	_ = v13097
	var v13101 int32
	_ = v13101
	var v13102 int32
	_ = v13102
	var v13103 int32
	_ = v13103
	var v13104 int32
	_ = v13104
	var v13105 int32
	_ = v13105
	var v13107 int32
	_ = v13107
	var v13110 int32
	_ = v13110
	var v13113 int32
	_ = v13113
	var v13115 int32
	_ = v13115
	var v13122 int32
	_ = v13122
	var v13123 int32
	_ = v13123
	var v13124 int32
	_ = v13124
	var v13129 int32
	_ = v13129
	var v13132 int32
	_ = v13132
	var v13137 int32
	_ = v13137
	var v13139 int32
	_ = v13139
	var v13143 int32
	_ = v13143
	var v13147 int64
	_ = v13147
	var v13150 int32
	_ = v13150
	var v13151 int32
	_ = v13151
	var v13152 int32
	_ = v13152
	var v13153 int32
	_ = v13153
	var v13154 int32
	_ = v13154
	var v13156 int32
	_ = v13156
	var v13160 int32
	_ = v13160
	var v13163 int32
	_ = v13163
	var v13165 int32
	_ = v13165
	var v13167 int32
	_ = v13167
	var v13168 int32
	_ = v13168
	var v13169 int32
	_ = v13169
	var v13171 int32
	_ = v13171
	var v13174 int32
	_ = v13174
	var v13176 int32
	_ = v13176
	var v13177 int32
	_ = v13177
	var v13184 int32
	_ = v13184
	var v13186 int32
	_ = v13186
	var v13189 int32
	_ = v13189
	var v13193 int32
	_ = v13193
	var v13197 int32
	_ = v13197
	var v13199 int32
	_ = v13199
	var v13200 int32
	_ = v13200
	var v13201 int32
	_ = v13201
	var v13203 int32
	_ = v13203
	var v13207 int32
	_ = v13207
	var v13211 int32
	_ = v13211
	var v13215 int32
	_ = v13215
	var v13223 int32
	_ = v13223
	var v13258 int32
	_ = v13258
	var v13262 int32
	_ = v13262
	var v13266 int32
	_ = v13266
	var v13267 int32
	_ = v13267
	var v13268 int32
	_ = v13268
	var v13270 int32
	_ = v13270
	var v13274 int32
	_ = v13274
	var v13287 int32
	_ = v13287
	var v13288 int32
	_ = v13288
	var v13331 int32
	_ = v13331
	var v13332 int32
	_ = v13332
	var v13335 int32
	_ = v13335
	var v13336 int32
	_ = v13336
	var v13338 int32
	_ = v13338
	var v13341 int32
	_ = v13341
	var v13343 int32
	_ = v13343
	var v13346 int32
	_ = v13346
	var v13348 int32
	_ = v13348
	var v13349 int32
	_ = v13349
	var v13353 int32
	_ = v13353
	var v13354 int32
	_ = v13354
	var v13361 int32
	_ = v13361
	var v13363 int32
	_ = v13363
	var v13366 int32
	_ = v13366
	var v13368 int32
	_ = v13368
	var v13369 int32
	_ = v13369
	var v13370 int32
	_ = v13370
	var v13372 int32
	_ = v13372
	var v13375 int32
	_ = v13375
	var v13379 int32
	_ = v13379
	var v13382 int32
	_ = v13382
	var v13388 int32
	_ = v13388
	var v13393 int32
	_ = v13393
	var v13397 int32
	_ = v13397
	var v13399 int32
	_ = v13399
	var v13403 int32
	_ = v13403
	var v13404 int32
	_ = v13404
	var v13405 int32
	_ = v13405
	var v13406 int32
	_ = v13406
	var v13410 int32
	_ = v13410
	var v13413 int32
	_ = v13413
	var v13414 int32
	_ = v13414
	var v13422 int32
	_ = v13422
	var v13425 int32
	_ = v13425
	var v13427 int32
	_ = v13427
	var v13429 int32
	_ = v13429
	var v13431 int32
	_ = v13431
	var v13434 int32
	_ = v13434
	var v13440 int32
	_ = v13440
	var v13447 int32
	_ = v13447
	var v13450 int32
	_ = v13450
	var v13454 int32
	_ = v13454
	var v13457 int32
	_ = v13457
	var v13463 int32
	_ = v13463
	var v13468 int32
	_ = v13468
	var v13471 int32
	_ = v13471
	var v13517 int32
	_ = v13517
	var v13520 int32
	_ = v13520
	var v13524 int32
	_ = v13524
	var v13529 int32
	_ = v13529
	var v13533 int32
	_ = v13533
	var v13536 int32
	_ = v13536
	var v13540 int32
	_ = v13540
	var v13545 int32
	_ = v13545
	var v13549 int32
	_ = v13549
	var v13552 int32
	_ = v13552
	var v13556 int32
	_ = v13556
	var v13561 int32
	_ = v13561
	var v13565 int32
	_ = v13565
	var v13568 int32
	_ = v13568
	var v13572 int32
	_ = v13572
	var v13577 int32
	_ = v13577
	var v13581 int32
	_ = v13581
	var v13584 int32
	_ = v13584
	var v13588 int32
	_ = v13588
	var v13593 int32
	_ = v13593
	var v13597 int32
	_ = v13597
	var v13600 int32
	_ = v13600
	var v13607 int32
	_ = v13607
	var v13612 int32
	_ = v13612
	var v13616 int32
	_ = v13616
	var v13619 int32
	_ = v13619
	var v13620 int32
	_ = v13620
	var v13628 int32
	_ = v13628
	var v13633 int32
	_ = v13633
	var v13638 int32
	_ = v13638
	var v13641 int32
	_ = v13641
	var v13645 int32
	_ = v13645
	var v13650 int32
	_ = v13650
	var v13654 int32
	_ = v13654
	var v13657 int32
	_ = v13657
	var v13665 int32
	_ = v13665
	var v13670 int32
	_ = v13670
	var v13674 int32
	_ = v13674
	var v13677 int32
	_ = v13677
	var v13684 int32
	_ = v13684
	var v13689 int32
	_ = v13689
	var v13693 int32
	_ = v13693
	var v13696 int32
	_ = v13696
	var v13700 int32
	_ = v13700
	var v13705 int32
	_ = v13705
	var v13709 int32
	_ = v13709
	var v13712 int32
	_ = v13712
	var v13718 int32
	_ = v13718
	var v13723 int32
	_ = v13723
	var v13728 int32
	_ = v13728
	var v13731 int32
	_ = v13731
	var v13735 int32
	_ = v13735
	var v13740 int32
	_ = v13740
	var v13744 int32
	_ = v13744
	var v13747 int32
	_ = v13747
	var v13751 int32
	_ = v13751
	var v13756 int32
	_ = v13756
	var v13760 int32
	_ = v13760
	var v13763 int32
	_ = v13763
	var v13767 int32
	_ = v13767
	var v13772 int32
	_ = v13772
	var v13776 int32
	_ = v13776
	var v13779 int32
	_ = v13779
	var v13785 int32
	_ = v13785
	var v13790 int32
	_ = v13790
	var v13794 int32
	_ = v13794
	var v13797 int32
	_ = v13797
	var v13801 int32
	_ = v13801
	var v13806 int32
	_ = v13806
	var v13810 int32
	_ = v13810
	var v13813 int32
	_ = v13813
	var v13817 int32
	_ = v13817
	var v13822 int32
	_ = v13822
	var v13826 int32
	_ = v13826
	var v13829 int32
	_ = v13829
	var v13833 int32
	_ = v13833
	var v13838 int32
	_ = v13838
	var v13842 int32
	_ = v13842
	var v13845 int32
	_ = v13845
	var v13851 int32
	_ = v13851
	var v13856 int32
	_ = v13856
	var v13860 int32
	_ = v13860
	var v13863 int32
	_ = v13863
	var v13867 int32
	_ = v13867
	var v13872 int32
	_ = v13872
	var v13879 int32
	_ = v13879
	var v13880 int32
	_ = v13880
	var v13911 int32
	_ = v13911
	var v13918 int32
	_ = v13918
	var v13919 int64
	_ = v13919
	var v13923 int32
	_ = v13923
	var v13925 int32
	_ = v13925
	var v13926 int32
	_ = v13926
	var v13929 int32
	_ = v13929
	var v13931 int32
	_ = v13931
	var v13933 int32
	_ = v13933
	var v13937 int32
	_ = v13937
	v40 = m.G0
	v42 = v40 - int32(32)
	m.G0 = v42
	v45 = l0
	v46 = l1
	v47 = int32(-1)
	v48 = int32(0)
	v77 = v42
	goto L1
L1:
	;
	goto L3
L3:
	;
	if v47 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v13918 = int32(m.ExcTag)
	v13919 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v13918 == int32(0) {
		goto L3084
	} else {
		goto L3085
	}
L6:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[0])))
	if v87 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v2040 = v48
	goto L8
L8:
	;
	if v2040 != 0 {
		goto L365
	} else {
		goto L366
	}
L9:
	;
	v862 = m.G0
	v864 = v862 - int32(32)
	m.G0 = v864
	v866 = int32(2)
	switch v866 {
	case 0, 2:
		goto L169
	default:
		goto L170
	}
L10:
	;
	v94 = m.G0
	v96 = v94 - int32(32)
	m.G0 = v96
	v99 = int32(967)
	switch v99 {
	case 0, 2:
		goto L14
	default:
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	v434 = m.G0
	v436 = v434 - int32(32)
	m.G0 = v436
	v439 = int32(967)
	switch v439 {
	case 0, 2:
		goto L80
	default:
		goto L81
	}
L13:
	;
	v136 = m.G0
	v138 = v136 - int32(32)
	m.G0 = v138
	v141 = int32(968)
	switch v141 {
	case 0, 2:
		goto L24
	default:
		goto L25
	}
L14:
	;
	F_sigemptyset(m, v96+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v96)+24)) = int32(268435456)
	switch v99 {
	case 0:
		goto L19
	default:
		goto L17
	case 2:
		goto L18
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[1])) = int32(965)
	goto L14
L16:
	;
	goto L21
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = int32(_a_F_PostgresMain_0)
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = int32(0)
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = int32(-2)
	goto L16
L21:
	;
	goto L22
L22:
	;
	v128 = F___sigaction(m, int32(1), v96+int32(12), int32(0))
	mBase = m.M
	m.G0 = v96 + int32(32)
	goto L13
L23:
	;
	v178 = m.G0
	v180 = v178 - int32(32)
	m.G0 = v180
	v183 = int32(974)
	switch v183 {
	case 0, 2:
		goto L34
	default:
		goto L35
	}
L24:
	;
	F_sigemptyset(m, v138+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v138)+24)) = int32(268435456)
	switch v141 {
	case 0:
		goto L29
	default:
		goto L27
	case 2:
		goto L28
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[2])) = int32(966)
	goto L24
L26:
	;
	goto L31
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v138)+12)) = int32(_a_F_PostgresMain_0)
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+12)) = int32(0)
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+12)) = int32(-2)
	goto L26
L31:
	;
	goto L32
L32:
	;
	v170 = F___sigaction(m, int32(2), v138+int32(12), int32(0))
	mBase = m.M
	m.G0 = v138 + int32(32)
	goto L23
L33:
	;
	v216 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[3])) = v216
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[4])) = v216
	v226 = v216
	goto L44
L34:
	;
	F_sigemptyset(m, v180+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v180)+24)) = int32(268435456)
	switch v183 {
	case 0:
		goto L39
	default:
		goto L37
	case 2:
		goto L38
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[5])) = int32(972)
	goto L34
L36:
	;
	goto L41
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v180)+12)) = int32(_a_F_PostgresMain_0)
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+12)) = int32(0)
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+12)) = int32(-2)
	goto L36
L41:
	;
	goto L42
L42:
	;
	v212 = F___sigaction(m, int32(15), v180+int32(12), int32(0))
	mBase = m.M
	m.G0 = v180 + int32(32)
	goto L33
L43:
	;
	v306 = int32(0)
	v308 = m.G0
	v310 = v308 - int32(32)
	m.G0 = v310
	switch v306 {
	case 0, 2:
		goto L50
	default:
		goto L51
	}
L44:
	;
	v228 = int32(40)
	v229 = v226 * v228
	v230 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_PostgresMain[6]))) = uint8(v230)
	*(*int32)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_PostgresMain[7]))) = v226
	v233 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_PostgresMain[8]))) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_PostgresMain[9]))) = v230
	*(*int64)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_PostgresMain[10]))) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_PostgresMain[11]))) = v230
	*(*uint8)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_PostgresMain[12]))) = uint8(v230)
	v244 = v226 | int32(1)
	v246 = v244 * v228
	*(*uint8)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_PostgresMain[6]))) = uint8(v230)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_PostgresMain[7]))) = v244
	*(*int64)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_PostgresMain[8]))) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_PostgresMain[9]))) = v230
	*(*int64)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_PostgresMain[10]))) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_PostgresMain[11]))) = v230
	*(*uint8)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_PostgresMain[12]))) = uint8(v230)
	v261 = v226 | int32(2)
	v263 = v261 * v228
	*(*uint8)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_PostgresMain[6]))) = uint8(v230)
	*(*int32)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_PostgresMain[7]))) = v261
	*(*int64)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_PostgresMain[8]))) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_PostgresMain[9]))) = v230
	*(*int64)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_PostgresMain[10]))) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_PostgresMain[11]))) = v230
	*(*uint8)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_PostgresMain[12]))) = uint8(v230)
	if v226 != int32(20) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v299 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[13])) = uint8(v299)
	F_pqsignal_be(m, int32(14), int32(1992))
	mBase = m.M
	goto L43
L46:
	;
	v280 = v226 | int32(3)
	v282 = v280 * int32(40)
	v283 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_PostgresMain[6]))) = uint8(v283)
	*(*int32)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_PostgresMain[7]))) = v280
	v286 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_PostgresMain[8]))) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_PostgresMain[9]))) = v283
	*(*int64)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_PostgresMain[10]))) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_PostgresMain[11]))) = v283
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_PostgresMain[12]))) = uint8(v283)
	v226 = v226 + int32(4)
	goto L44
L47:
	;
	goto L48
L48:
	;
	goto L45
L49:
	;
	v350 = m.G0
	v352 = v350 - int32(32)
	m.G0 = v352
	v355 = int32(970)
	switch v355 {
	case 0, 2:
		goto L60
	default:
		goto L61
	}
L50:
	;
	F_sigemptyset(m, v310+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v310)+24)) = int32(268435456)
	switch v306 {
	case 0:
		goto L55
	default:
		goto L53
	case 2:
		goto L54
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[14])) = int32(-2)
	goto L50
L52:
	;
	goto L57
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v310)+12)) = int32(_a_F_PostgresMain_0)
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310)+12)) = int32(0)
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310)+12)) = int32(-2)
	goto L52
L57:
	;
	goto L58
L58:
	;
	v342 = F___sigaction(m, int32(13), v310+int32(12), int32(0))
	mBase = m.M
	m.G0 = v310 + int32(32)
	goto L49
L59:
	;
	v392 = m.G0
	v394 = v392 - int32(32)
	m.G0 = v394
	v397 = int32(1116)
	switch v397 {
	case 0, 2:
		goto L70
	default:
		goto L71
	}
L60:
	;
	F_sigemptyset(m, v352+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v352)+24)) = int32(268435456)
	switch v355 {
	case 0:
		goto L65
	default:
		goto L63
	case 2:
		goto L64
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[15])) = int32(968)
	goto L60
L62:
	;
	goto L67
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v352)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v352)+12)) = int32(_a_F_PostgresMain_0)
	goto L62
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v352)+12)) = int32(0)
	goto L62
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v352)+12)) = int32(-2)
	goto L62
L67:
	;
	goto L68
L68:
	;
	v384 = F___sigaction(m, int32(10), v352+int32(12), int32(0))
	mBase = m.M
	m.G0 = v352 + int32(32)
	goto L59
L69:
	;
	goto L9
L70:
	;
	F_sigemptyset(m, v394+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v394)+24)) = int32(268435456)
	switch v397 {
	case 0:
		goto L75
	default:
		goto L73
	case 2:
		goto L74
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[16])) = int32(1114)
	goto L70
L72:
	;
	goto L77
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v394)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v394)+12)) = int32(_a_F_PostgresMain_0)
	goto L72
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v394)+12)) = int32(0)
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v394)+12)) = int32(-2)
	goto L72
L77:
	;
	goto L78
L78:
	;
	v426 = F___sigaction(m, int32(12), v394+int32(12), int32(0))
	mBase = m.M
	m.G0 = v394 + int32(32)
	goto L69
L79:
	;
	v476 = m.G0
	v478 = v476 - int32(32)
	m.G0 = v478
	v481 = int32(968)
	switch v481 {
	case 0, 2:
		goto L90
	default:
		goto L91
	}
L80:
	;
	F_sigemptyset(m, v436+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v436)+24)) = int32(268435456)
	switch v439 {
	case 0:
		goto L85
	default:
		goto L83
	case 2:
		goto L84
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[1])) = int32(965)
	goto L80
L82:
	;
	goto L87
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v436)+12)) = int32(_a_F_PostgresMain_0)
	goto L82
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436)+12)) = int32(0)
	goto L82
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436)+12)) = int32(-2)
	goto L82
L87:
	;
	goto L88
L88:
	;
	v468 = F___sigaction(m, int32(1), v436+int32(12), int32(0))
	mBase = m.M
	m.G0 = v436 + int32(32)
	goto L79
L89:
	;
	v518 = m.G0
	v520 = v518 - int32(32)
	m.G0 = v520
	v523 = int32(974)
	switch v523 {
	case 0, 2:
		goto L100
	default:
		goto L101
	}
L90:
	;
	F_sigemptyset(m, v478+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v478)+24)) = int32(268435456)
	switch v481 {
	case 0:
		goto L95
	default:
		goto L93
	case 2:
		goto L94
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[2])) = int32(966)
	goto L90
L92:
	;
	goto L97
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v478)+12)) = int32(_a_F_PostgresMain_0)
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478)+12)) = int32(0)
	goto L92
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478)+12)) = int32(-2)
	goto L92
L97:
	;
	goto L98
L98:
	;
	v510 = F___sigaction(m, int32(2), v478+int32(12), int32(0))
	mBase = m.M
	m.G0 = v478 + int32(32)
	goto L89
L99:
	;
	v560 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[17])))
	if v560 != 0 {
		goto L109
	} else {
		goto L110
	}
L100:
	;
	F_sigemptyset(m, v520+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v520)+24)) = int32(268435456)
	switch v523 {
	case 0:
		goto L105
	default:
		goto L103
	case 2:
		goto L104
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[5])) = int32(972)
	goto L100
L102:
	;
	goto L107
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v520)+12)) = int32(_a_F_PostgresMain_0)
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520)+12)) = int32(0)
	goto L102
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520)+12)) = int32(-2)
	goto L102
L107:
	;
	goto L108
L108:
	;
	v552 = F___sigaction(m, int32(15), v520+int32(12), int32(0))
	mBase = m.M
	m.G0 = v520 + int32(32)
	goto L99
L109:
	;
	v561 = int32(1258)
	goto L111
L110:
	;
	v561 = int32(972)
	goto L111
L111:
	;
	v564 = m.G0
	v566 = v564 - int32(32)
	m.G0 = v566
	v569 = v561 + int32(2)
	switch v569 {
	case 0, 2:
		goto L113
	default:
		goto L114
	}
L112:
	;
	v602 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[3])) = v602
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[4])) = v602
	v612 = v602
	goto L123
L113:
	;
	F_sigemptyset(m, v566+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v566)+24)) = int32(268435456)
	switch v569 {
	case 0:
		goto L118
	default:
		goto L116
	case 2:
		goto L117
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[18])) = v561
	goto L113
L115:
	;
	goto L120
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v566)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v566)+12)) = int32(_a_F_PostgresMain_0)
	goto L115
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v566)+12)) = int32(0)
	goto L115
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v566)+12)) = int32(-2)
	goto L115
L120:
	;
	goto L121
L121:
	;
	v598 = F___sigaction(m, int32(3), v566+int32(12), int32(0))
	mBase = m.M
	m.G0 = v566 + int32(32)
	goto L112
L122:
	;
	v692 = int32(0)
	v694 = m.G0
	v696 = v694 - int32(32)
	m.G0 = v696
	switch v692 {
	case 0, 2:
		goto L129
	default:
		goto L130
	}
L123:
	;
	v614 = int32(40)
	v615 = v612 * v614
	v616 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v615)+uint32(_c_F_PostgresMain[6]))) = uint8(v616)
	*(*int32)(unsafe.Add(mBase, uint32(v615)+uint32(_c_F_PostgresMain[7]))) = v612
	v619 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v615)+uint32(_c_F_PostgresMain[8]))) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v615)+uint32(_c_F_PostgresMain[9]))) = v616
	*(*int64)(unsafe.Add(mBase, uint32(v615)+uint32(_c_F_PostgresMain[10]))) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v615)+uint32(_c_F_PostgresMain[11]))) = v616
	*(*uint8)(unsafe.Add(mBase, uint32(v615)+uint32(_c_F_PostgresMain[12]))) = uint8(v616)
	v630 = v612 | int32(1)
	v632 = v630 * v614
	*(*uint8)(unsafe.Add(mBase, uint32(v632)+uint32(_c_F_PostgresMain[6]))) = uint8(v616)
	*(*int32)(unsafe.Add(mBase, uint32(v632)+uint32(_c_F_PostgresMain[7]))) = v630
	*(*int64)(unsafe.Add(mBase, uint32(v632)+uint32(_c_F_PostgresMain[8]))) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v632)+uint32(_c_F_PostgresMain[9]))) = v616
	*(*int64)(unsafe.Add(mBase, uint32(v632)+uint32(_c_F_PostgresMain[10]))) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v632)+uint32(_c_F_PostgresMain[11]))) = v616
	*(*uint8)(unsafe.Add(mBase, uint32(v632)+uint32(_c_F_PostgresMain[12]))) = uint8(v616)
	v647 = v612 | int32(2)
	v649 = v647 * v614
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_PostgresMain[6]))) = uint8(v616)
	*(*int32)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_PostgresMain[7]))) = v647
	*(*int64)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_PostgresMain[8]))) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_PostgresMain[9]))) = v616
	*(*int64)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_PostgresMain[10]))) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_PostgresMain[11]))) = v616
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_PostgresMain[12]))) = uint8(v616)
	if v612 != int32(20) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v685 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[13])) = uint8(v685)
	F_pqsignal_be(m, int32(14), int32(1992))
	mBase = m.M
	goto L122
L125:
	;
	v666 = v612 | int32(3)
	v668 = v666 * int32(40)
	v669 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v668)+uint32(_c_F_PostgresMain[6]))) = uint8(v669)
	*(*int32)(unsafe.Add(mBase, uint32(v668)+uint32(_c_F_PostgresMain[7]))) = v666
	v672 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v668)+uint32(_c_F_PostgresMain[8]))) = v672
	*(*int32)(unsafe.Add(mBase, uint32(v668)+uint32(_c_F_PostgresMain[9]))) = v669
	*(*int64)(unsafe.Add(mBase, uint32(v668)+uint32(_c_F_PostgresMain[10]))) = v672
	*(*int32)(unsafe.Add(mBase, uint32(v668)+uint32(_c_F_PostgresMain[11]))) = v669
	*(*uint8)(unsafe.Add(mBase, uint32(v668)+uint32(_c_F_PostgresMain[12]))) = uint8(v669)
	v612 = v612 + int32(4)
	goto L123
L126:
	;
	goto L127
L127:
	;
	goto L124
L128:
	;
	v736 = m.G0
	v738 = v736 - int32(32)
	m.G0 = v738
	v741 = int32(970)
	switch v741 {
	case 0, 2:
		goto L139
	default:
		goto L140
	}
L129:
	;
	F_sigemptyset(m, v696+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v696)+24)) = int32(268435456)
	switch v692 {
	case 0:
		goto L134
	default:
		goto L132
	case 2:
		goto L133
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[14])) = int32(-2)
	goto L129
L131:
	;
	goto L136
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v696)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v696)+12)) = int32(_a_F_PostgresMain_0)
	goto L131
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v696)+12)) = int32(0)
	goto L131
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v696)+12)) = int32(-2)
	goto L131
L136:
	;
	goto L137
L137:
	;
	v728 = F___sigaction(m, int32(13), v696+int32(12), int32(0))
	mBase = m.M
	m.G0 = v696 + int32(32)
	goto L128
L138:
	;
	v776 = int32(0)
	v778 = m.G0
	v780 = v778 - int32(32)
	m.G0 = v780
	switch v776 {
	case 0, 2:
		goto L149
	default:
		goto L150
	}
L139:
	;
	F_sigemptyset(m, v738+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v738)+24)) = int32(268435456)
	switch v741 {
	case 0:
		goto L144
	default:
		goto L142
	case 2:
		goto L143
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[15])) = int32(968)
	goto L139
L141:
	;
	goto L146
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v738)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v738)+12)) = int32(_a_F_PostgresMain_0)
	goto L141
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v738)+12)) = int32(0)
	goto L141
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v738)+12)) = int32(-2)
	goto L141
L146:
	;
	goto L147
L147:
	;
	v770 = F___sigaction(m, int32(10), v738+int32(12), int32(0))
	mBase = m.M
	m.G0 = v738 + int32(32)
	goto L138
L148:
	;
	v820 = m.G0
	v822 = v820 - int32(32)
	m.G0 = v822
	v825 = int32(972)
	switch v825 {
	case 0, 2:
		goto L159
	default:
		goto L160
	}
L149:
	;
	F_sigemptyset(m, v780+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v780)+24)) = int32(268435456)
	switch v776 {
	case 0:
		goto L154
	default:
		goto L152
	case 2:
		goto L153
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[16])) = int32(-2)
	goto L149
L151:
	;
	goto L156
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v780)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v780)+12)) = int32(_a_F_PostgresMain_0)
	goto L151
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v780)+12)) = int32(0)
	goto L151
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v780)+12)) = int32(-2)
	goto L151
L156:
	;
	goto L157
L157:
	;
	v812 = F___sigaction(m, int32(12), v780+int32(12), int32(0))
	mBase = m.M
	m.G0 = v780 + int32(32)
	goto L148
L158:
	;
	goto L9
L159:
	;
	F_sigemptyset(m, v822+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v822)+24)) = int32(268435456)
	switch v825 {
	case 0:
		goto L164
	default:
		goto L162
	case 2:
		goto L163
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[19])) = int32(970)
	goto L159
L161:
	;
	goto L166
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v822)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v822)+12)) = int32(_a_F_PostgresMain_0)
	goto L161
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v822)+12)) = int32(0)
	goto L161
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v822)+12)) = int32(-2)
	goto L161
L166:
	;
	goto L167
L167:
	;
	v854 = F___sigaction(m, int32(8), v822+int32(12), int32(0))
	mBase = m.M
	m.G0 = v822 + int32(32)
	goto L158
L168:
	;
	F_BaseInit(m)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L178
	}
L169:
	;
	F_sigemptyset(m, v864+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v864)+24)) = int32(268435456)
	switch v866 {
	case 0:
		goto L174
	default:
		goto L172
	case 2:
		goto L173
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[20])) = int32(0)
	goto L169
L171:
	;
	goto L175
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v864)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v864)+12)) = int32(_a_F_PostgresMain_0)
	v889 = int32(268435461)
	goto L171
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v864)+12)) = int32(0)
	v889 = int32(268435457)
	goto L171
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v864)+12)) = int32(-2)
	v889 = int32(268435457)
	goto L171
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v864)+24)) = v889
	goto L177
L177:
	;
	v896 = F___sigaction(m, int32(17), v864+int32(12), int32(0))
	mBase = m.M
	m.G0 = v864 + int32(32)
	goto L168
L178:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_PostgresMain_1), int32(0))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L179
	}
L179:
	;
	v907 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	if v907 == int32(2) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v912 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[22]))
	if v912 != 0 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	goto L182
L182:
	;
	v993 = int32(0)
	v996 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[0])))
	F_InitPostgres(m, v45, v993, v46, v993, v996^int32(1), v993)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L209
	}
L183:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v912)+8))
	if base.Ui32(int32(_a_F_PostgresMain_2)) < base.Ui32(v915) {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	v919 = int32(32)
	goto L185
L185:
	;
	v921 = int32(0)
	v925 = m.G0
	v927 = v925 - int32(16)
	m.G0 = v927
	*(*int32)(unsafe.Add(mBase, uint32(v927))) = v921
	v933 = F_open(m, int32(_a_F_PostgresMain_3), v921, v927)
	mBase = m.M
	if v933 != int32(-1) {
		goto L190
	} else {
		goto L191
	}
L186:
	;
	v918 = int32(32)
	goto L188
L187:
	;
	v918 = int32(4)
	goto L188
L188:
	;
	v919 = v918
	goto L185
L189:
	;
	if v966 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L190:
	;
	v936 = int32(1)
	if v919 == int32(0) {
		v959 = v936
		goto L193
	} else {
		goto L194
	}
L191:
	;
	v966 = v921
	goto L192
L192:
	;
	m.G0 = v927 + int32(16)
	goto L189
L193:
	;
	v961 = F_close(m, v933)
	mBase = m.M
	v966 = v959
	goto L192
L194:
	;
	v939 = int32(_a_F_PostgresMain_4)
	v940 = v919
	goto L195
L195:
	;
	v945 = F_read(m, v933, v939, v940)
	mBase = m.M
	if v945 <= int32(0) {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v959 = v936
	goto L193
L197:
	;
	v949 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[23]))
	if v949 == int32(27) {
		goto L195
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v954 = v940 - v945
	if v954 != 0 {
		v939 = v939 + v945
		v940 = v954
		goto L195
	} else {
		goto L201
	}
L200:
	;
	v959 = int32(0)
	goto L193
L201:
	;
	goto L196
L202:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[24])) = v919
	goto L182
L205:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L206
	}
L206:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_5), int32(0))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(_a_F_PostgresMain_7), int32(_a_F_PostgresMain_8))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
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
	v1003 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[25]))
	if v1003 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	F_MemoryContextDelete(m, v1003)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v1010 = int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[26])) = v1010
	v1012 = m.G0
	v1014 = v1012 - int32(32)
	m.G0 = v1014
	v1017 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	if v1017 != v1010 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[25])) = int32(0)
	goto L212
L214:
	;
	m.G0 = v1014 + int32(32)
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[17])))
	if v1150 != int32(1) {
		goto L235
	} else {
		goto L236
	}
L215:
	;
	v1021 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[27])) = uint8(v1021)
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])))
	if v1025 == v1021 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	if v1035 != 0 {
		goto L220
	} else {
		goto L221
	}
L217:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+308))
	v1033 = base.B2i32(v1031 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])) = uint8(v1033)
	v1035 = v1033
	goto L219
L218:
	;
	v1035 = int32(0)
	goto L219
L219:
	;
	goto L216
L220:
	;
	v1037 = int32(0)
	v1040 = int32(10)
	v1046 = F_set_config_with_handle(m, int32(_a_F_PostgresMain_9), v1037, int32(_a_F_PostgresMain_10), v1037, v1040, v1040, v1037, int32(1), v1037, v1037)
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v1049 = v1014 + int32(12)
	v1051 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[30]))
	F_hash_seq_init(m, v1049, v1051)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L224
	}
L223:
	;
	goto L222
L224:
	;
	v1054 = F_hash_seq_search(m, v1049)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L225
	}
L225:
	;
	if v1054 == int32(0) {
		goto L214
	} else {
		goto L226
	}
L226:
	;
	v1060 = v1054
	goto L227
L227:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1060)+4))
	v1098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1097)+20)))
	if v1098&int32(64) != 0 {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	goto L214
L229:
	;
	F_ReportGUCOption(m, v1097)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v1105 = F_hash_seq_search(m, v1014+int32(12))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L233
	}
L232:
	;
	goto L231
L233:
	;
	if v1105 != 0 {
		v1060 = v1105
		goto L227
	} else {
		goto L234
	}
L234:
	;
	goto L228
L235:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[31]))
	v1165 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[32]))
	if v1165 == int32(1) {
		goto L239
	} else {
		goto L240
	}
L236:
	;
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[33])))
	if v1154&int32(1) == int32(0) {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	F_on_proc_exit(m, int32(1259))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L238
	}
L238:
	;
	goto L235
L239:
	;
	v1170 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[34]))
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[35])) = v1170
	v1175 = F_pgstat_prep_pending_entry(m, int32(1), v1163, int64(0), int32(0))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[0])))
	if v1184 == int32(1) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1175)+12))
	v1178 = *(*int64)(unsafe.Add(mBase, uint32(v1177)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v1177)+184)) = v1178 + int64(1)
	goto L241
L243:
	;
	v1187 = int32(0)
	v1191 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])))
	if v1191 == int32(1) {
		goto L247
	} else {
		goto L248
	}
L244:
	;
	goto L245
L245:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	if v1455 == int32(2) {
		goto L274
	} else {
		goto L275
	}
L246:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[36])) = uint8(v1201)
	v1204 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[37]))
	if v1204 <= int32(0) {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1196)+308))
	v1199 = base.B2i32(v1197 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])) = uint8(v1199)
	v1201 = v1199
	goto L249
L248:
	;
	v1201 = v1187
	goto L249
L249:
	;
	goto L246
L250:
	;
	F_on_shmem_exit(m, int32(1106), int64(0))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L262
	}
L251:
	;
	v1209 = v1187
	goto L252
L252:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[38]))
	v1250 = v1247 + v1209*int32(96)
	v1251 = int32(164)
	v1252 = v1250 + v1251
	v1255 = base.AtomicRmwXchg32(m, v1250, v1251, int32(1))
	if v1255 != 0 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	goto L250
L254:
	;
	F_s_lock(m, v1252, int32(_a_F_PostgresMain_11))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v1260 = v1250 + int32(88)
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1260)))
	if v1261 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	goto L256
L258:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[39]))
	v1266 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1260)+24)) = v1266
	v1268 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1260)+16)) = uint8(v1268)
	*(*int64)(unsafe.Add(mBase, uint32(v1260)+8)) = v1266
	*(*int32)(unsafe.Add(mBase, uint32(v1260)+4)) = v1268
	*(*int32)(unsafe.Add(mBase, uint32(v1260))) = v1265
	*(*int64)(unsafe.Add(mBase, uint32(v1260)+32)) = v1266
	*(*int64)(unsafe.Add(mBase, uint32(v1260)+40)) = v1266
	v1279 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1260)+48)) = v1279
	*(*int64)(unsafe.Add(mBase, uint32(v1260)+56)) = v1279
	*(*int64)(unsafe.Add(mBase, uint32(v1260)+64)) = v1279
	*(*int64)(unsafe.Add(mBase, uint32(v1260)+80)) = v1266
	*(*int32)(unsafe.Add(mBase, uint32(v1260)+72)) = v1268
	v1290 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v1260)+88)) = base.B2i32(v1290 != v1268)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1260)+76)), uint32(v1268))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[40])) = v1260
	goto L250
L259:
	;
	goto L260
L260:
	;
	v1299 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1252))), uint32(v1299))
	v1303 = v1209 + int32(1)
	v1305 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[37]))
	if v1303 < v1305 {
		v1209 = v1303
		goto L252
	} else {
		goto L261
	}
L261:
	;
	goto L253
L262:
	;
	F_CreateAuxProcessResourceOwner(m)
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L263
	}
L263:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[41]))
	v1355 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v1353+v1355<<(uint(int32(2))%32))+48)) = int32(3)
	v1363 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[17])))
	if v1363 == int32(1) {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[31]))
	if v1378 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L265:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[41]))
	*(*int32)(unsafe.Add(mBase, uint32(v1367+int32(36)))) = int32(1)
	v1374 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[43]))
	v1376 = F_pgmem_kill(m, v1374, int32(10))
	mBase = m.M
	goto L267
L266:
	;
	goto L267
L267:
	;
	goto L264
L268:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[44]))
	v1386 = F_LWLockAcquire(m, v1382+int32(512), int32(0))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[45]))
	v1412 = F_MemoryContextAllocZero(m, v1410, int32(_a_F_PostgresMain_12))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L273
	}
L271:
	;
	v1389 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[46]))
	v1390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1389)+36)))
	v1392 = v1390 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1389)+36)) = uint8(v1392)
	v1395 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[47]))
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+12))
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v1396+v1397))) = uint8(v1392)
	v1401 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[44]))
	F_LWLockRelease(m, v1401+int32(512))
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L272
	}
L272:
	;
	goto L270
L273:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[48])) = v1412
	goto L245
L274:
	;
	v1459 = v77 + int32(16)
	F_pq_beginmessage(m, v1459, int32(75))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L277
	}
L275:
	;
	v1496 = v1455
	goto L276
L276:
	;
	if v1496 == int32(1) {
		goto L281
	} else {
		goto L282
	}
L277:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[39]))
	F_enlargeStringInfo(m, v1459, int32(4))
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L278
	}
L278:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v1473 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v1468+v1469))) = base.I32_rotr(v1464, int32(24))&v1473 | base.I32_rotr(v1464&v1473, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = v1468 + int32(4)
	v1486 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[24]))
	F_appendBinaryStringInfo(m, v1459, int32(_a_F_PostgresMain_4), v1486)
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L279
	}
L279:
	;
	F_pq_endmessage(m, v1459)
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L280
	}
L280:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	v1496 = v1492
	goto L276
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(_a_F_PostgresMain_13)
	F_pg_printf(m, int32(_a_F_PostgresMain_14), v77)
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v1506 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[45]))
	v1511 = F_AllocSetContextCreateInternal(m, v1506, int32(_a_F_PostgresMain_15), int32(0), int32(_a_F_PostgresMain_16), int32(_a_F_PostgresMain_17))
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L285
	}
L284:
	;
	goto L283
L285:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[49])) = v1511
	v1516 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[45]))
	v1521 = F_AllocSetContextCreateInternal(m, v1516, int32(_a_F_PostgresMain_18), int32(0), int32(_a_F_PostgresMain_16), int32(_a_F_PostgresMain_17))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v1521
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[51])) = v1521
	F_initStringInfo(m, int32(_a_F_PostgresMain_19))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L287
	}
L287:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[45]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v1531
	v1533 = int32(0)
	v1535 = m.G0
	v1537 = v1535 - int32(80)
	m.G0 = v1537
	v1540 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[17])))
	if v1540 != int32(1) {
		goto L290
	} else {
		goto L291
	}
L288:
	;
	goto L361
L289:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L358
	}
L290:
	;
	m.G0 = v1537 + int32(80)
	goto L288
L291:
	;
	v1544 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[52])))
	if v1544&int32(1) == int32(0) {
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[31]))
	if v1550 == int32(0) {
		goto L290
	} else {
		goto L293
	}
L293:
	;
	v1554 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[53])))
	if v1554&int32(1) == int32(0) {
		goto L290
	} else {
		goto L294
	}
L294:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L295
	}
L295:
	;
	v1562 = F_EventCacheLookup(m, int32(4))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L298
	}
L296:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L357
	}
L297:
	;
	v1698 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])))
	if v1698 == int32(1) {
		goto L324
	} else {
		goto L325
	}
L298:
	;
	if v1562 == int32(0) {
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1562)+4))
	if v1566 <= int32(0) {
		goto L297
	} else {
		goto L300
	}
L300:
	;
	v1572 = v1533
	v1576 = v1533
	goto L301
L301:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1562)+12))
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1608+v1572<<(uint(int32(2))%32))))
	v1613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1612)+4)))
	v1615 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[54]))
	if v1615 == int32(1) {
		goto L305
	} else {
		goto L306
	}
L302:
	;
	if v1632 == int32(0) {
		goto L297
	} else {
		goto L317
	}
L303:
	;
	v1634 = v1572 + int32(1)
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1562)+4))
	if v1634 < v1635 {
		v1572 = v1634
		v1576 = v1632
		goto L301
	} else {
		goto L316
	}
L304:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1612)+8))
	if v1622 != 0 {
		goto L310
	} else {
		goto L311
	}
L305:
	;
	if v1613 != int32(79) {
		goto L304
	} else {
		goto L308
	}
L306:
	;
	goto L307
L307:
	;
	if v1613 == int32(82) {
		v1632 = v1576
		goto L303
	} else {
		goto L309
	}
L308:
	;
	v1632 = v1576
	goto L303
L309:
	;
	goto L304
L310:
	;
	v1624 = F_bms_is_member(m, int32(162), v1622)
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1612)))
	v1629 = F_lappend_oid(m, v1576, v1628)
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L315
	}
L313:
	;
	if v1624 == int32(0) {
		v1632 = v1576
		goto L303
	} else {
		goto L314
	}
L314:
	;
	goto L312
L315:
	;
	v1632 = v1629
	goto L303
L316:
	;
	goto L302
L317:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1537)+24)) = int64(695784701952)
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+20)) = int32(_a_F_PostgresMain_20)
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+16)) = int32(447)
	v1645 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L318
	}
L318:
	;
	F_PushActiveSnapshot(m, v1645)
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L319
	}
L319:
	;
	F_EventTriggerInvoke(m, v1632, v1537+int32(16))
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L320
	}
L320:
	;
	F_list_free(m, v1632)
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L321
	}
L321:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L322
	}
L322:
	;
	goto L296
L323:
	;
	if v1708 != 0 {
		goto L296
	} else {
		goto L327
	}
L324:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v1703)+308))
	v1706 = base.B2i32(v1704 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])) = uint8(v1706)
	v1708 = v1706
	goto L326
L325:
	;
	v1708 = int32(0)
	goto L326
L326:
	;
	goto L323
L327:
	;
	v1710 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[31]))
	v1711 = m.G0
	v1713 = v1711 - int32(32)
	m.G0 = v1713
	v1715 = int32(264)
	*(*uint16)(unsafe.Add(mBase, uint32(v1713)+30)) = uint16(v1715)
	v1717 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1713)+28)) = uint16(v1717)
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+24)) = v1710
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+20)) = int32(1262)
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+16)) = v1717
	v1732 = F_LockAcquireExtended(m, v1713+int32(16), int32(8), v1717, int32(1), v1713+int32(12), v1717)
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L330
	}
L328:
	;
	m.G0 = v1713 + int32(32)
	if v1732 == int32(0) {
		goto L296
	} else {
		goto L333
	}
L329:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L331
	}
L330:
	;
	switch v1732 {
	case 0, 3:
		goto L328
	default:
		goto L329
	}
L331:
	;
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+12))
	v1737 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1736)+53)) = uint8(v1737)
	goto L332
L332:
	;
	goto L328
L333:
	;
	v1745 = F_EventCacheLookup(m, int32(4))
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L336
	}
L334:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1537)+24)) = int64(695784701952)
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+20)) = int32(_a_F_PostgresMain_20)
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+16)) = int32(447)
	F_list_free(m, v1799)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L356
	}
L335:
	;
	v1846 = F_table_open(m, int32(1262), int32(3))
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L344
	}
L336:
	;
	if v1745 == int32(0) {
		goto L335
	} else {
		goto L337
	}
L337:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1745)+4))
	if v1749 <= int32(0) {
		goto L335
	} else {
		goto L338
	}
L338:
	;
	v1752 = int32(0)
	v1757 = v1752
	v1762 = v1752
	goto L339
L339:
	;
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(v1745)+12))
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1793+v1757<<(uint(int32(2))%32))))
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1797)))
	v1799 = F_lappend_oid(m, v1762, v1798)
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L341
	}
L340:
	;
	if v1799 != 0 {
		goto L334
	} else {
		goto L343
	}
L341:
	;
	v1802 = v1757 + int32(1)
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1745)+4))
	if v1802 < v1803 {
		v1757 = v1802
		v1762 = v1799
		goto L339
	} else {
		goto L342
	}
L342:
	;
	goto L340
L343:
	;
	goto L335
L344:
	;
	v1849 = v1537 + int32(16)
	v1854 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_PostgresMain[31])))
	F_ScanKeyInit(m, v1849, int32(1), int32(3), int32(184), v1854)
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L345
	}
L345:
	;
	F_systable_inplace_update_begin(m, v1846, int32(2672), v1849, v1537+int32(76), v1537+int32(72))
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L346
	}
L346:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(v1537)+76))
	if v1864 == int32(0) {
		goto L289
	} else {
		goto L347
	}
L347:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v1864)+16))
	v1868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1867)+22)))
	v1869 = v1867 + v1868
	v1870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1869)+79)))
	if v1870 == int32(1) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	F_relation_close(m, v1846, int32(3))
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L354
	}
L349:
	;
	v1873 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1869)+79)) = uint8(v1873)
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v1537)+72))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1537)+76))
	F_systable_inplace_update_finish(m, v1875, v1876)
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1537)+72))
	F_systable_inplace_update_cancel(m, v1879)
	mBase = m.M
	v1881 = m.ExcPending
	if v1881 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L353
	}
L352:
	;
	goto L348
L353:
	;
	goto L348
L354:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1537)+76))
	F_pfree(m, v1885)
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L355
	}
L355:
	;
	goto L296
L356:
	;
	goto L296
L357:
	;
	goto L290
L358:
	;
	v1984 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v1537))) = v1984
	F_errmsg_internal(m, int32(_a_F_PostgresMain_21), v1537)
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L359
	}
L359:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_22), int32(980), int32(_a_F_PostgresMain_23))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L360
	}
L360:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L361:
	;
	v1994 = int32(_a_F_PostgresMain_24)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[55])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[56])) = v77 + int32(12)
	goto L364
L362:
	;
	v2040 = int32(0)
	goto L8
L364:
	;
	goto L362
L365:
	;
	v2041 = int32(_a_F_PostgresMain_25)
	v2043 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[57])) = v2043 + int32(1)
	v2048 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[58])) = v2048
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[3])) = v2048
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[4])) = v2048
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[6])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[12])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[59])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[60])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[61])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[62])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[63])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[64])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[65])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[66])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[67])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[68])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[69])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[70])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[71])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[72])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[73])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[74])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[75])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[76])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[77])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[78])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[79])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[80])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[81])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[82])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[83])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[84])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[85])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[86])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[87])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[88])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[89])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[90])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[91])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[92])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[93])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[94])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[95])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[96])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[97])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[98])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[99])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[100])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[101])) = uint8(v2048)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[102])) = uint8(v2048)
	goto L369
L366:
	;
	goto L367
L367:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[103])) = int32(_a_F_PostgresMain_24)
	v2496 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[104])))
	if v2496 == int32(0) {
		goto L441
	} else {
		goto L442
	}
L368:
	;
	goto L367
L369:
	;
	v2195 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[105])) = v2195
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[106])) = uint8(v2195)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[107])) = uint8(v2195)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[108])) = uint8(v2195)
	v2207 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[109]))
	v2208 = *(*int32)(unsafe.Add(mBase, uint32(v2207)))
	m.T0[v2208].(func(*base.Module))(m)
	mBase = m.M
	v2210 = m.ExcPending
	if v2210 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L370
	}
L370:
	;
	F_EmitErrorReport(m)
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L371
	}
L371:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[110])) = int32(0)
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v2217 = m.ExcPending
	if v2217 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L372
	}
L372:
	;
	v2219 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[0])))
	if v2219 == int32(1) {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	F_LWLockReleaseAll(m)
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L377
	}
L374:
	;
	goto L375
L375:
	;
	v2287 = m.G0
	v2289 = v2287 - int32(32)
	m.G0 = v2289
	v2292 = v2289 + int32(12)
	v2294 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[111]))
	F_hash_seq_init(m, v2292, v2294)
	mBase = m.M
	v2296 = m.ExcPending
	if v2296 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L405
	}
L376:
	;
	goto L375
L377:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L378
	}
L378:
	;
	v2227 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[112]))
	*(*int32)(unsafe.Add(mBase, uint32(v2227))) = int32(0)
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v2231 = m.ExcPending
	if v2231 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L379
	}
L379:
	;
	v2233 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[113]))
	if v2233 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v2244 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	if v2244 != 0 {
		goto L384
	} else {
		goto L385
	}
L381:
	;
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(v2233)+1168))
	if v2236 < int32(0) {
		goto L380
	} else {
		goto L382
	}
L382:
	;
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(v2233)+1168))
	v2240 = F_close(m, v2239)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2233)+1168)) = int32(-1)
	goto L383
L383:
	;
	goto L380
L384:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v2246 = m.ExcPending
	if v2246 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	F_ReplicationSlotCleanup(m, int32(0))
	mBase = m.M
	v2249 = m.ExcPending
	if v2249 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L388
	}
L387:
	;
	goto L386
L388:
	;
	v2251 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[115])) = v2251
	v2254 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v2255 = *(*int32)(unsafe.Add(mBase, uint32(v2254)+24))
	goto L389
L389:
	;
	if base.B2i32(v2255 != v2251) == int32(0) {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v2262 = m.ExcPending
	if v2262 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L393
	}
L391:
	;
	goto L392
L392:
	;
	v2264 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[117]))
	if v2264 != 0 {
		goto L394
	} else {
		goto L395
	}
L393:
	;
	goto L392
L394:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v2285 = m.ExcPending
	if v2285 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L404
	}
L395:
	;
	v2266 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[118]))
	if v2266 != 0 {
		goto L394
	} else {
		goto L396
	}
L396:
	;
	v2268 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[40]))
	v2269 = *(*int32)(unsafe.Add(mBase, uint32(v2268)+4))
	if v2269 != 0 {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v2272 = base.AtomicRmwXchg32(m, v2268, int32(76), int32(1))
	if v2272 != 0 {
		goto L400
	} else {
		goto L401
	}
L398:
	;
	goto L399
L399:
	;
	goto L376
L400:
	;
	F_s_lock(m, v2268+int32(76), int32(_a_F_PostgresMain_11))
	mBase = m.M
	v2277 = m.ExcPending
	if v2277 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L403
	}
L401:
	;
	goto L402
L402:
	;
	v2278 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2268)+4)) = v2278
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2268)+76)), uint32(v2278))
	goto L399
L403:
	;
	goto L402
L404:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L405:
	;
	v2297 = F_hash_seq_search(m, v2292)
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L406
	}
L406:
	;
	if v2297 != 0 {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v2303 = v2297
	goto L410
L408:
	;
	goto L409
L409:
	;
	m.G0 = v2289 + int32(32)
	v2394 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	if v2394 != 0 {
		goto L418
	} else {
		goto L419
	}
L410:
	;
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(v2303)+64))
	v2339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2338)+85)))
	if v2339 == int32(1) {
		goto L412
	} else {
		goto L413
	}
L411:
	;
	goto L409
L412:
	;
	v2342 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2338)+84)) = uint8(v2342)
	F_PortalDrop(m, v2338, v2342)
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L415
	}
L413:
	;
	goto L414
L414:
	;
	v2349 = F_hash_seq_search(m, v2289+int32(12))
	mBase = m.M
	v2350 = m.ExcPending
	if v2350 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L416
	}
L415:
	;
	goto L414
L416:
	;
	if v2349 != 0 {
		v2303 = v2349
		goto L410
	} else {
		goto L417
	}
L417:
	;
	goto L411
L418:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L421
	}
L419:
	;
	goto L420
L420:
	;
	F_ReplicationSlotCleanup(m, int32(0))
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L422
	}
L421:
	;
	goto L420
L422:
	;
	v2401 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[119])))
	if v2401 != 0 {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v2403 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[120]))
	m.T0[v2403].(func(*base.Module))(m)
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L426
	}
L424:
	;
	goto L425
L425:
	;
	v2408 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[49]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v2408
	F_FlushErrorState(m)
	mBase = m.M
	v2411 = m.ExcPending
	if v2411 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L427
	}
L426:
	;
	goto L425
L427:
	;
	v2413 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[121])))
	if v2413 != 0 {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v2415 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[104])) = uint8(v2415)
	goto L430
L429:
	;
	goto L430
L430:
	;
	v2418 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])) = uint8(v2418)
	v2421 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[123])))
	if v2421 == v2418 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v2424 = int32(_a_F_PostgresMain_25)
	v2426 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[57])) = v2426 - int32(1)
	v2431 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[104])))
	if v2431 == int32(0) {
		goto L434
	} else {
		goto L435
	}
L432:
	;
	goto L433
L433:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L437
	}
L434:
	;
	v2435 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[124])) = uint8(v2435)
	goto L436
L435:
	;
	goto L436
L436:
	;
	goto L368
L437:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L438
	}
L438:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_26), int32(0))
	mBase = m.M
	v2447 = m.ExcPending
	if v2447 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L439
	}
L439:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(_a_F_PostgresMain_27), int32(_a_F_PostgresMain_28))
	mBase = m.M
	v2452 = m.ExcPending
	if v2452 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L440
	}
L440:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L441:
	;
	v2500 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[124])) = uint8(v2500)
	goto L443
L442:
	;
	goto L443
L443:
	;
	goto L444
L444:
	;
	v2541 = int32(0)
	v2546 = m.G0
	v2548 = v2546 - int32(528)
	m.G0 = v2548
	v2552 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[49]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v2552
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[121])) = uint8(v2541)
	F_MemoryContextReset(m, v2552)
	mBase = m.M
	v2558 = m.ExcPending
	if v2558 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L446
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[125])) = int32(99)
	m.Env.Emscripten_exit_with_live_runtime(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L446:
	;
	F_initStringInfo(m, v2548+int32(440))
	mBase = m.M
	v2562 = m.ExcPending
	if v2562 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L447
	}
L447:
	;
	v2564 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[126]))
	if v2564 == int32(0) {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v2621 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[124])))
	if v2621 == int32(1) {
		goto L464
	} else {
		goto L465
	}
L449:
	;
	v2568 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[127]))
	if v2568 != 0 {
		goto L448
	} else {
		goto L450
	}
L450:
	;
	v2570 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[128]))
	if v2570 == int32(0) {
		goto L448
	} else {
		goto L451
	}
L451:
	;
	v2573 = *(*int32)(unsafe.Add(mBase, uint32(v2570)))
	if v2573 != 0 {
		goto L448
	} else {
		goto L452
	}
L452:
	;
	F_pairingheap_remove(m, int32(_a_F_PostgresMain_29), v2564+int32(52))
	mBase = m.M
	v2578 = m.ExcPending
	if v2578 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[126])) = int32(0)
	v2583 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[127]))
	if v2583 != 0 {
		goto L448
	} else {
		goto L454
	}
L454:
	;
	v2585 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[128]))
	if v2585 == int32(0) {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v2589 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[129])) = v2589
	v2592 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[46]))
	*(*int32)(unsafe.Add(mBase, uint32(v2592)+52)) = v2589
	goto L448
L456:
	;
	goto L457
L457:
	;
	v2596 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[46]))
	v2597 = *(*int32)(unsafe.Add(mBase, uint32(v2596)+52))
	v2598 = int32(3)
	v2601 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[128]))
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(v2601-int32(48))))
	if base.B2i32(base.Ui32(v2597) < base.Ui32(v2598))|base.B2i32(base.Ui32(v2604) < base.Ui32(v2598)) == int32(0) {
		goto L459
	} else {
		goto L460
	}
L458:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[129])) = v2604
	*(*int32)(unsafe.Add(mBase, uint32(v2596)+52)) = v2604
	goto L448
L459:
	;
	if v2597-v2604 < int32(0) {
		goto L458
	} else {
		goto L462
	}
L460:
	;
	goto L461
L461:
	;
	if base.Ui32(v2604) <= base.Ui32(v2597) {
		goto L448
	} else {
		goto L463
	}
L462:
	;
	goto L448
L463:
	;
	goto L458
L464:
	;
	v2625 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v2626 = *(*int32)(unsafe.Add(mBase, uint32(v2625)+24))
	goto L468
L465:
	;
	goto L466
L466:
	;
	v3042 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[107])) = uint8(v3042)
	v3045 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	if v3045 == int32(2) {
		goto L555
	} else {
		goto L556
	}
L467:
	;
	v2712 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[27])))
	if v2712 == int32(0) {
		goto L504
	} else {
		goto L505
	}
L468:
	;
	if (v2626-int32(7))&int32(-9) == int32(0) {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v2634 = int32(0)
	F_pgstat_report_activity(m, int32(6), v2634)
	mBase = m.M
	v2637 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[130]))
	if v2637 <= v2634 {
		goto L467
	} else {
		goto L472
	}
L470:
	;
	goto L471
L471:
	;
	v2652 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v2653 = *(*int32)(unsafe.Add(mBase, uint32(v2652)+24))
	goto L478
L472:
	;
	v2641 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[131]))
	if v2641 != 0 {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v2644 = base.B2i32(v2641 <= v2637)
	goto L475
L474:
	;
	v2644 = int32(0)
	goto L475
L475:
	;
	if v2644 != 0 {
		goto L467
	} else {
		goto L476
	}
L476:
	;
	v2646 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[106])) = uint8(v2646)
	F_enable_timeout_after(m, int32(7), v2637)
	mBase = m.M
	v2650 = m.ExcPending
	if v2650 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L477
	}
L477:
	;
	goto L467
L478:
	;
	if v2653 != int32(0) {
		goto L479
	} else {
		goto L480
	}
L479:
	;
	v2657 = int32(0)
	F_pgstat_report_activity(m, int32(4), v2657)
	mBase = m.M
	v2660 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[130]))
	if v2660 <= v2657 {
		goto L467
	} else {
		goto L482
	}
L480:
	;
	goto L481
L481:
	;
	v2675 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[132]))
	if v2675 != 0 {
		goto L488
	} else {
		goto L489
	}
L482:
	;
	v2664 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[131]))
	if v2664 != 0 {
		goto L483
	} else {
		goto L484
	}
L483:
	;
	v2667 = base.B2i32(v2664 <= v2660)
	goto L485
L484:
	;
	v2667 = int32(0)
	goto L485
L485:
	;
	if v2667 != 0 {
		goto L467
	} else {
		goto L486
	}
L486:
	;
	v2669 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[106])) = uint8(v2669)
	F_enable_timeout_after(m, int32(7), v2660)
	mBase = m.M
	v2673 = m.ExcPending
	if v2673 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L487
	}
L487:
	;
	goto L467
L488:
	;
	F_ProcessNotifyInterrupt(m, int32(0))
	mBase = m.M
	v2678 = m.ExcPending
	if v2678 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L491
	}
L489:
	;
	goto L490
L490:
	;
	v2680 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v2681 = m.ExcPending
	if v2681 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L492
	}
L491:
	;
	goto L490
L492:
	;
	v2685 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[77])))
	goto L493
L493:
	;
	if int32(0) < v2680 {
		goto L495
	} else {
		goto L496
	}
L494:
	;
	v2697 = int32(0)
	F_pgstat_report_activity(m, int32(2), v2697)
	mBase = m.M
	v2700 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[133]))
	if v2700 <= v2697 {
		goto L467
	} else {
		goto L502
	}
L495:
	;
	if v2685 != 0 {
		goto L494
	} else {
		goto L498
	}
L496:
	;
	goto L497
L497:
	;
	if v2685 == int32(0) {
		goto L494
	} else {
		goto L500
	}
L498:
	;
	F_enable_timeout_after(m, int32(10), v2680)
	mBase = m.M
	v2690 = m.ExcPending
	if v2690 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L499
	}
L499:
	;
	goto L494
L500:
	;
	F_disable_timeout(m, int32(10))
	mBase = m.M
	v2695 = m.ExcPending
	if v2695 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L501
	}
L501:
	;
	goto L494
L502:
	;
	v2704 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[108])) = uint8(v2704)
	F_enable_timeout_after(m, int32(9), v2700)
	mBase = m.M
	v2708 = m.ExcPending
	if v2708 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L503
	}
L503:
	;
	goto L467
L504:
	;
	v2840 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[134]))
	if v2840 != int64(-9223372036854775807-1) {
		goto L519
	} else {
		goto L520
	}
L505:
	;
	v2716 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[135])))
	if v2716 != int32(1) {
		goto L506
	} else {
		goto L507
	}
L506:
	;
	v2745 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[136]))
	if v2745 == int32(0) {
		goto L504
	} else {
		goto L514
	}
L507:
	;
	v2721 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])))
	if v2721 == int32(1) {
		goto L509
	} else {
		goto L510
	}
L508:
	;
	if v2731 != 0 {
		goto L506
	} else {
		goto L512
	}
L509:
	;
	v2726 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v2727 = *(*int32)(unsafe.Add(mBase, uint32(v2726)+308))
	v2729 = base.B2i32(v2727 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])) = uint8(v2729)
	v2731 = v2729
	goto L511
L510:
	;
	v2731 = int32(0)
	goto L511
L511:
	;
	goto L508
L512:
	;
	v2733 = int32(0)
	v2736 = int32(10)
	v2742 = F_set_config_with_handle(m, int32(_a_F_PostgresMain_9), v2733, int32(_a_F_PostgresMain_30), v2733, v2736, v2736, v2733, int32(1), v2733, v2733)
	mBase = m.M
	v2743 = m.ExcPending
	if v2743 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L513
	}
L513:
	;
	goto L506
L514:
	;
	v2752 = v2745
	goto L515
L515:
	;
	v2787 = *(*int32)(unsafe.Add(mBase, uint32(v2752)))
	F_ReportGUCOption(m, v2752-int32(80))
	mBase = m.M
	v2791 = m.ExcPending
	if v2791 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L517
	}
L516:
	;
	goto L504
L517:
	;
	v2793 = v2752 - int32(52)
	v2794 = *(*int32)(unsafe.Add(mBase, uint32(v2793)))
	*(*int32)(unsafe.Add(mBase, uint32(v2793))) = v2794 & int32(-5)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[136])) = v2787
	if v2787 != 0 {
		v2752 = v2787
		goto L515
	} else {
		goto L518
	}
L518:
	;
	goto L516
L519:
	;
	v2930 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	v2931 = m.G0
	v2933 = v2931 - int32(16)
	m.G0 = v2933
	v2935 = int32(2)
	if base.Ui32(v2930-v2935) <= base.Ui32(v2935) {
		goto L537
	} else {
		goto L538
	}
L520:
	;
	v2844 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[137])))
	if v2844&int32(8) == int32(0) {
		goto L519
	} else {
		goto L521
	}
L521:
	;
	v2850 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[32]))
	switch v2850 - int32(1) {
	case 0, 5:
		goto L522
	default:
		goto L519
	}
L522:
	;
	v2857 = m.G0
	v2858 = int32(16)
	v2859 = v2857 - v2858
	m.G0 = v2859
	F_gettimeofday(m, v2859)
	mBase = m.M
	v2862 = *(*int64)(unsafe.Add(mBase, uint32(v2859)))
	v2863 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2859)+8)))
	m.G0 = v2859 + v2858
	v2871 = v2863 + v2862*int64(1000000) - int64(946684800000000)
	goto L523
L523:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[134])) = v2871
	v2874 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[138]))
	v2876 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[139]))
	v2878 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[140]))
	v2880 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[141]))
	v2882 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[142]))
	v2885 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2886 = m.ExcPending
	if v2886 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L524
	}
L524:
	;
	if v2885 == int32(0) {
		goto L519
	} else {
		goto L525
	}
L525:
	;
	if v2876 < v2874 {
		goto L526
	} else {
		goto L527
	}
L526:
	;
	v2892 = v2874 - v2876
	goto L528
L527:
	;
	v2892 = int64(0)
	goto L528
L528:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v2548)+432)) = base.F64_div(base.F64_convert_i64_u(v2892), float64(1000))
	if v2880 < v2878 {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v2900 = v2878 - v2880
	goto L531
L530:
	;
	v2900 = int64(0)
	goto L531
L531:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v2548)+424)) = base.F64_div(base.F64_convert_i64_u(v2900), float64(1000))
	if v2882 < v2871 {
		goto L532
	} else {
		goto L533
	}
L532:
	;
	v2908 = v2871 - v2882
	goto L534
L533:
	;
	v2908 = int64(0)
	goto L534
L534:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v2548)+416)) = base.F64_div(base.F64_convert_i64_u(v2908), float64(1000))
	F_errmsg(m, int32(_a_F_PostgresMain_31), v2548+int32(416))
	mBase = m.M
	v2917 = m.ExcPending
	if v2917 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L535
	}
L535:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(_a_F_PostgresMain_32), int32(_a_F_PostgresMain_33))
	mBase = m.M
	v2922 = m.ExcPending
	if v2922 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L536
	}
L536:
	;
	goto L519
L537:
	;
	F_pq_beginmessage(m, v2933, int32(90))
	mBase = m.M
	v2941 = m.ExcPending
	if v2941 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L540
	}
L538:
	;
	goto L539
L539:
	;
	m.G0 = v2933 + int32(16)
	v3000 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[124])) = uint8(v3000)
	goto L466
L540:
	;
	v2942 = m.G0
	v2944 = v2942 - int32(16)
	m.G0 = v2944
	v2947 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v2948 = *(*int32)(unsafe.Add(mBase, uint32(v2947)+24))
	if base.Ui32(int32(20)) <= base.Ui32(v2948) {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2954 = m.ExcPending
	if v2954 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L544
	}
L542:
	;
	goto L543
L543:
	;
	v2972 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2948)+uint32(_c_F_PostgresMain[143]))))
	m.G0 = v2944 + int32(16)
	F_enlargeStringInfo(m, v2933, int32(1))
	mBase = m.M
	v2978 = m.ExcPending
	if v2978 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L551
	}
L544:
	;
	v2955 = *(*int32)(unsafe.Add(mBase, uint32(v2947)+24))
	if base.Ui32(v2955) <= base.Ui32(int32(19)) {
		goto L546
	} else {
		goto L547
	}
L545:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2944))) = v2962
	F_errmsg_internal(m, int32(_a_F_PostgresMain_34), v2944)
	mBase = m.M
	v2966 = m.ExcPending
	if v2966 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L549
	}
L546:
	;
	v2960 = *(*int32)(unsafe.Add(mBase, uint32(v2955<<(uint(int32(2))%32))+uint32(_c_F_PostgresMain[144])))
	v2962 = v2960
	goto L548
L547:
	;
	v2962 = int32(_a_F_PostgresMain_35)
	goto L548
L548:
	;
	goto L545
L549:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_36), int32(_a_F_PostgresMain_37), int32(_a_F_PostgresMain_38))
	mBase = m.M
	v2971 = m.ExcPending
	if v2971 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L550
	}
L550:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L551:
	;
	v2979 = *(*int32)(unsafe.Add(mBase, uint32(v2933)+4))
	v2980 = *(*int32)(unsafe.Add(mBase, uint32(v2933)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2979+v2980))) = uint8(v2972)
	*(*int32)(unsafe.Add(mBase, uint32(v2933)+4)) = v2979 + int32(1)
	F_pq_endmessage(m, v2933)
	mBase = m.M
	v2987 = m.ExcPending
	if v2987 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L552
	}
L552:
	;
	v2989 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[109]))
	v2990 = *(*int32)(unsafe.Add(mBase, uint32(v2989)+4))
	v2991 = m.T0[v2990].(func(*base.Module) int32)(m)
	mBase = m.M
	v2992 = m.ExcPending
	if v2992 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L553
	}
L553:
	;
	goto L539
L554:
	;
	v3410 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[106])))
	if v3410 == int32(1) {
		goto L652
	} else {
		goto L653
	}
L555:
	;
	v3048 = int32(_a_F_PostgresMain_39)
	v3050 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[145]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[145])) = v3050 + int32(1)
	F_pq_startmsgread(m)
	mBase = m.M
	v3055 = m.ExcPending
	if v3055 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L558
	}
L556:
	;
	goto L557
L557:
	;
	F_pg_printf(m, int32(_a_F_PostgresMain_40), int32(0))
	mBase = m.M
	v3162 = m.ExcPending
	if v3162 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L592
	}
L558:
	;
	v3058 = F_pq_getbyte(m)
	mBase = m.M
	v3059 = m.ExcPending
	if v3059 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L569
	}
L559:
	;
	v3150 = F_pq_getmessage(m, v2548+int32(440), v3149)
	mBase = m.M
	v3151 = m.ExcPending
	if v3151 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L588
	}
L560:
	;
	v3149 = int32(1073741822)
	goto L559
L561:
	;
	v3146 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[121])) = uint8(v3146)
	goto L560
L562:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v3132 = m.ExcPending
	if v3132 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L584
	}
L563:
	;
	v3126 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[121])) = uint8(v3126)
	v3149 = int32(_a_F_PostgresMain_41)
	goto L559
L564:
	;
	v3123 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[104])) = uint8(v3123)
	goto L563
L565:
	;
	v3120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[121])) = uint8(v3120)
	goto L560
L566:
	;
	v3113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[121])) = uint8(v3113)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[104])) = uint8(v3113)
	v3149 = int32(_a_F_PostgresMain_41)
	goto L559
L567:
	;
	v3109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[121])) = uint8(v3109)
	v3149 = int32(_a_F_PostgresMain_41)
	goto L559
L568:
	;
	v3063 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v3064 = *(*int32)(unsafe.Add(mBase, uint32(v3063)+20))
	goto L570
L569:
	;
	switch v3058 + int32(1) {
	case 0:
		goto L568
	default:
		goto L562
	case 67, 81:
		goto L561
	case 68, 69, 70, 73:
		goto L567
	case 71, 82, 101:
		goto L565
	case 84:
		goto L566
	case 89:
		goto L564
	case 100, 103:
		goto L563
	}
L570:
	;
	if v3064 == int32(2) {
		goto L571
	} else {
		goto L572
	}
L571:
	;
	v3067 = int32(-1)
	v3070 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v3071 = m.ExcPending
	if v3071 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L574
	}
L572:
	;
	goto L573
L573:
	;
	v3087 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21])) = v3087
	v3089 = int32(-1)
	v3092 = F_errstart(m, int32(14), v3087)
	mBase = m.M
	v3093 = m.ExcPending
	if v3093 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L579
	}
L574:
	;
	if v3070 == int32(0) {
		v3374 = v3067
		goto L554
	} else {
		goto L575
	}
L575:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v3076 = m.ExcPending
	if v3076 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L576
	}
L576:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_42), int32(0))
	mBase = m.M
	v3080 = m.ExcPending
	if v3080 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L577
	}
L577:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(490), int32(_a_F_PostgresMain_43))
	mBase = m.M
	v3085 = m.ExcPending
	if v3085 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L578
	}
L578:
	;
	v3374 = v3067
	goto L554
L579:
	;
	if v3092 == int32(0) {
		v3374 = v3089
		goto L554
	} else {
		goto L580
	}
L580:
	;
	F_errcode(m, int32(50332160))
	mBase = m.M
	v3098 = m.ExcPending
	if v3098 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L581
	}
L581:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMain_44), int32(0))
	mBase = m.M
	v3102 = m.ExcPending
	if v3102 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L582
	}
L582:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(501), int32(_a_F_PostgresMain_43))
	mBase = m.M
	v3107 = m.ExcPending
	if v3107 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L583
	}
L583:
	;
	v3374 = v3089
	goto L554
L584:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v3135 = m.ExcPending
	if v3135 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L585
	}
L585:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548))) = v3058
	F_errmsg(m, int32(_a_F_PostgresMain_45), v2548)
	mBase = m.M
	v3139 = m.ExcPending
	if v3139 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L586
	}
L586:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(576), int32(_a_F_PostgresMain_43))
	mBase = m.M
	v3144 = m.ExcPending
	if v3144 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L587
	}
L587:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L588:
	;
	if v3150 != 0 {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	v3374 = int32(-1)
	goto L554
L590:
	;
	goto L591
L591:
	;
	v3153 = int32(_a_F_PostgresMain_39)
	v3155 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[145]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[145])) = v3155 - int32(1)
	v3374 = v3058
	goto L554
L592:
	;
	v3164 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[146]))
	v3165 = F_fflush(m, v3164)
	mBase = m.M
	v3166 = m.ExcPending
	if v3166 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L593
	}
L593:
	;
	v3168 = v2548 + int32(440)
	v3169 = *(*int32)(unsafe.Add(mBase, uint32(v3168)))
	v3170 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3169))) = uint8(v3170)
	*(*int32)(unsafe.Add(mBase, uint32(v3168)+12)) = v3170
	*(*int32)(unsafe.Add(mBase, uint32(v3168)+4)) = v3170
	goto L594
L594:
	;
	v3177 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[147]))
	goto L597
L595:
	;
	F_appendStringInfoChar(m, v2548+int32(440), int32(0))
	mBase = m.M
	v3357 = m.ExcPending
	if v3357 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L646
	}
L596:
	;
	F_appendStringInfoChar(m, v2548+int32(440), int32(10))
	mBase = m.M
	v3350 = m.ExcPending
	if v3350 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L645
	}
L597:
	;
	v3218 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[148]))
	if v3218 != 0 {
		goto L599
	} else {
		goto L600
	}
L598:
	;
	v3343 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+444))
	if v3343 != 0 {
		goto L595
	} else {
		goto L644
	}
L599:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3220 = m.ExcPending
	if v3220 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L602
	}
L600:
	;
	goto L601
L601:
	;
	v3221 = F_do_getc(m, v3177)
	mBase = m.M
	v3222 = m.ExcPending
	if v3222 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L603
	}
L602:
	;
	goto L601
L603:
	;
	v3224 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[23]))
	v3226 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[107])))
	if v3226 != 0 {
		goto L605
	} else {
		goto L606
	}
L604:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[23])) = v3224
	switch v3221 + int32(1) {
	case 0:
		goto L633
	default:
		goto L635
	case 11:
		goto L636
	}
L605:
	;
	v3228 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[148]))
	if v3228 != 0 {
		goto L608
	} else {
		goto L609
	}
L606:
	;
	goto L607
L607:
	;
	v3243 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[149]))
	if v3243 == int32(0) {
		goto L604
	} else {
		goto L618
	}
L608:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3230 = m.ExcPending
	if v3230 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L611
	}
L609:
	;
	goto L610
L610:
	;
	v3232 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[150]))
	if v3232 != 0 {
		goto L612
	} else {
		goto L613
	}
L611:
	;
	goto L610
L612:
	;
	F_ProcessCatchupInterrupt(m)
	mBase = m.M
	v3234 = m.ExcPending
	if v3234 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L615
	}
L613:
	;
	goto L614
L614:
	;
	v3236 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[132]))
	if v3236 == int32(0) {
		goto L604
	} else {
		goto L616
	}
L615:
	;
	goto L614
L616:
	;
	F_ProcessNotifyInterrupt(m, int32(1))
	mBase = m.M
	v3241 = m.ExcPending
	if v3241 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L617
	}
L617:
	;
	goto L604
L618:
	;
	v3247 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[151]))
	v3248 = int32(0)
	v3251 = base.AtomicRmwOr32(m, v3248, int32(_a_F_PostgresMain_46), v3248)
	v3252 = *(*int32)(unsafe.Add(mBase, uint32(v3247)))
	if v3252 != 0 {
		goto L620
	} else {
		goto L621
	}
L619:
	;
	goto L604
L620:
	;
	goto L619
L621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3247))) = int32(1)
	v3255 = int32(0)
	v3258 = base.AtomicRmwOr32(m, v3255, int32(_a_F_PostgresMain_46), v3255)
	v3259 = *(*int32)(unsafe.Add(mBase, uint32(v3247)+4))
	if v3259 == v3255 {
		goto L620
	} else {
		goto L622
	}
L622:
	;
	v3262 = *(*int32)(unsafe.Add(mBase, uint32(v3247)+12))
	if v3262 == int32(0) {
		goto L620
	} else {
		goto L623
	}
L623:
	;
	v3266 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[39]))
	if v3266 == v3262 {
		goto L624
	} else {
		goto L625
	}
L624:
	;
	v3268 = m.G0
	v3270 = v3268 - int32(16)
	m.G0 = v3270
	v3273 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[152]))
	if v3273 == int32(0) {
		goto L627
	} else {
		goto L628
	}
L625:
	;
	goto L626
L626:
	;
	v3296 = F_pgmem_kill(m, v3262, int32(23))
	mBase = m.M
	goto L620
L627:
	;
	m.G0 = v3270 + int32(16)
	goto L619
L628:
	;
	v3276 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3270)+15)) = uint8(v3276)
	goto L629
L629:
	;
	v3280 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[153]))
	v3284 = F_write(m, v3280, v3270+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v3284 {
		goto L627
	} else {
		goto L631
	}
L630:
	;
	goto L627
L631:
	;
	v3288 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[23]))
	if v3288 == int32(27) {
		goto L629
	} else {
		goto L632
	}
L632:
	;
	goto L630
L633:
	;
	goto L598
L634:
	;
	if v3303 <= int32(0) {
		goto L596
	} else {
		goto L642
	}
L635:
	;
	F_appendStringInfoChar(m, v2548+int32(440), base.I32_extend8_s(v3221))
	mBase = m.M
	v3327 = m.ExcPending
	if v3327 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L641
	}
L636:
	;
	v3303 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+444))
	v3305 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[154])))
	if v3305 == int32(0) {
		goto L634
	} else {
		goto L637
	}
L637:
	;
	if v3303 < int32(2) {
		goto L635
	} else {
		goto L638
	}
L638:
	;
	v3310 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+440))
	v3311 = v3310 + v3303
	v3314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3311-int32(1)))))
	if v3314 != int32(10) {
		goto L635
	} else {
		goto L639
	}
L639:
	;
	v3319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3311-int32(2)))))
	if v3319 == int32(59) {
		goto L595
	} else {
		goto L640
	}
L640:
	;
	goto L635
L641:
	;
	goto L597
L642:
	;
	v3330 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+440))
	v3334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3330+v3303-int32(1)))))
	if v3334 != int32(92) {
		goto L596
	} else {
		goto L643
	}
L643:
	;
	v3338 = v3303 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+444)) = v3338
	v3341 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3338+v3330))) = uint8(v3341)
	goto L597
L644:
	;
	v3374 = int32(-1)
	goto L554
L645:
	;
	goto L595
L646:
	;
	v3359 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[155])))
	if v3359 != 0 {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v3360 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+400)) = v3360
	F_pg_printf(m, int32(_a_F_PostgresMain_47), v2548+int32(400))
	mBase = m.M
	v3366 = m.ExcPending
	if v3366 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L650
	}
L648:
	;
	goto L649
L649:
	;
	v3367 = F_fflush(m, v3164)
	mBase = m.M
	v3368 = m.ExcPending
	if v3368 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L651
	}
L650:
	;
	goto L649
L651:
	;
	v3374 = int32(81)
	goto L554
L652:
	;
	F_disable_timeout(m, int32(7))
	mBase = m.M
	v3415 = m.ExcPending
	if v3415 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L655
	}
L653:
	;
	goto L654
L654:
	;
	v3420 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[108])))
	if v3420 == int32(1) {
		goto L656
	} else {
		goto L657
	}
L655:
	;
	v3417 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[106])) = uint8(v3417)
	goto L654
L656:
	;
	F_disable_timeout(m, int32(9))
	mBase = m.M
	v3425 = m.ExcPending
	if v3425 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L659
	}
L657:
	;
	goto L658
L658:
	;
	v3430 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[148]))
	if v3430 != 0 {
		goto L660
	} else {
		goto L661
	}
L659:
	;
	v3427 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[108])) = uint8(v3427)
	goto L658
L660:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3432 = m.ExcPending
	if v3432 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L663
	}
L661:
	;
	goto L662
L662:
	;
	v3434 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[107])) = uint8(v3434)
	v3437 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[156]))
	if v3437 != 0 {
		goto L664
	} else {
		goto L665
	}
L663:
	;
	goto L662
L664:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[156])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v3443 = m.ExcPending
	if v3443 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L667
	}
L665:
	;
	goto L666
L666:
	;
	if v3374 != int32(-1) {
		goto L691
	} else {
		goto L692
	}
L667:
	;
	goto L666
L668:
	;
	goto L445
L669:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13860 = m.ExcPending
	if v13860 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3080
	}
L670:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13842 = m.ExcPending
	if v13842 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3076
	}
L671:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13826 = m.ExcPending
	if v13826 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3072
	}
L672:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13810 = m.ExcPending
	if v13810 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3068
	}
L673:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13794 = m.ExcPending
	if v13794 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3064
	}
L674:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13776 = m.ExcPending
	if v13776 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3060
	}
L675:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13760 = m.ExcPending
	if v13760 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3056
	}
L676:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13744 = m.ExcPending
	if v13744 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3052
	}
L677:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13728 = m.ExcPending
	if v13728 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3048
	}
L678:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13709 = m.ExcPending
	if v13709 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3044
	}
L679:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13693 = m.ExcPending
	if v13693 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3040
	}
L680:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13674 = m.ExcPending
	if v13674 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3036
	}
L681:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13654 = m.ExcPending
	if v13654 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3032
	}
L682:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13638 = m.ExcPending
	if v13638 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3028
	}
L683:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13616 = m.ExcPending
	if v13616 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3024
	}
L684:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13597 = m.ExcPending
	if v13597 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3020
	}
L685:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13581 = m.ExcPending
	if v13581 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3016
	}
L686:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13565 = m.ExcPending
	if v13565 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3012
	}
L687:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13549 = m.ExcPending
	if v13549 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3008
	}
L688:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13533 = m.ExcPending
	if v13533 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3004
	}
L689:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13517 = m.ExcPending
	if v13517 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3000
	}
L690:
	;
	m.G0 = v2548 + int32(528)
	goto L444
L691:
	;
	v3447 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[104])))
	if v3447&int32(1) != 0 {
		goto L690
	} else {
		goto L694
	}
L692:
	;
	goto L693
L693:
	;
	switch v3374 + int32(1) {
	case 0:
		goto L698
	default:
		goto L696
	case 67:
		goto L705
	case 68:
		goto L702
	case 69:
		goto L701
	case 70:
		goto L704
	case 71:
		goto L703
	case 73:
		goto L700
	case 81:
		goto L706
	case 82:
		goto L707
	case 84:
		goto L699
	case 89:
		goto L697
	case 100, 101, 103:
		goto L690
	}
L694:
	;
	goto L693
L695:
	;
	F_pq_putemptymessage(m, int32(110))
	mBase = m.M
	v13471 = m.ExcPending
	if v13471 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2999
	}
L696:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v13454 = m.ExcPending
	if v13454 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2995
	}
L697:
	;
	v13440 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	if v13440 == int32(2) {
		goto L2990
	} else {
		goto L2991
	}
L698:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[157])) = int32(2)
	goto L697
L699:
	;
	F_pq_getmsgend(m, v2548+int32(440))
	mBase = m.M
	v13410 = m.ExcPending
	if v13410 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2976
	}
L700:
	;
	F_pq_getmsgend(m, v2548+int32(440))
	mBase = m.M
	v13397 = m.ExcPending
	if v13397 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2973
	}
L701:
	;
	v13139 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[0])))
	if v13139 == int32(1) {
		goto L673
	} else {
		goto L2917
	}
L702:
	;
	v13097 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[0])))
	if v13097 == int32(1) {
		goto L675
	} else {
		goto L2899
	}
L703:
	;
	v11979 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[0])))
	if v11979 == int32(1) {
		goto L676
	} else {
		goto L2649
	}
L704:
	;
	v11192 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[0])))
	if v11192 == int32(1) {
		goto L679
	} else {
		goto L2463
	}
L705:
	;
	v10094 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[0])))
	if v10094 == int32(1) {
		goto L686
	} else {
		goto L2252
	}
L706:
	;
	v9680 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[0])))
	if v9680 == int32(1) {
		goto L689
	} else {
		goto L2117
	}
L707:
	;
	v3453 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[158]))
	if v3453 < int32(0) {
		goto L709
	} else {
		goto L710
	}
L708:
	;
	v3460 = v2548 + int32(440)
	v3461 = F_pq_getmsgstring(m, v3460)
	mBase = m.M
	v3462 = m.ExcPending
	if v3462 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L712
	}
L709:
	;
	v3457 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[159])) = v3457
	goto L711
L710:
	;
	goto L711
L711:
	;
	goto L708
L712:
	;
	F_pq_getmsgend(m, v3460)
	mBase = m.M
	v3464 = m.ExcPending
	if v3464 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L713
	}
L713:
	;
	v3466 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[0])))
	if v3466 == int32(1) {
		goto L715
	} else {
		goto L716
	}
L714:
	;
	v9676 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[124])) = uint8(v9676)
	goto L690
L715:
	;
	v3469 = int32(0)
	v3479 = m.G0
	v3481 = v3479 - int32(_a_F_PostgresMain_48)
	m.G0 = v3481
	v3484 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[40]))
	v3486 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50]))
	v3488 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[117]))
	if v3488 == v3469 {
		v3509 = v3484
		goto L718
	} else {
		goto L719
	}
L716:
	;
	goto L717
L717:
	;
	v8424 = int32(0)
	v8426 = m.G0
	v8428 = v8426 - int32(112)
	m.G0 = v8428
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[110])) = v3461
	v8433 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	v8435 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[160])))
	F_pgstat_report_activity(m, int32(3), v3461)
	mBase = m.M
	if v8435 == int32(1) {
		goto L1888
	} else {
		goto L1889
	}
L718:
	;
	v3510 = *(*int32)(unsafe.Add(mBase, uint32(v3509)+4))
	if v3510 != int32(4) {
		goto L753
	} else {
		goto L754
	}
L719:
	;
	v3491 = *(*int32)(unsafe.Add(mBase, uint32(v3484)+4))
	if v3491 == int32(4) {
		v3509 = v3484
		goto L718
	} else {
		goto L720
	}
L720:
	;
	v3496 = base.AtomicRmwXchg32(m, v3484, int32(76), int32(1))
	if v3496 != 0 {
		goto L721
	} else {
		goto L722
	}
L721:
	;
	F_s_lock(m, v3484+int32(76), int32(_a_F_PostgresMain_11))
	mBase = m.M
	v3501 = m.ExcPending
	if v3501 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L724
	}
L722:
	;
	goto L723
L723:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3484)+4)) = int32(4)
	v3504 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3484)+76)), uint32(v3504))
	v3508 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[40]))
	v3509 = v3508
	goto L718
L724:
	;
	goto L723
L725:
	;
	m.G0 = v3481 + int32(_a_F_PostgresMain_48)
	if v3944 != 0 {
		goto L714
	} else {
		goto L1887
	}
L726:
	;
	F_EndReplicationCommand(m, v8300)
	mBase = m.M
	v8333 = m.ExcPending
	if v8333 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1885
	}
L727:
	;
	v8282 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v8283 = m.ExcPending
	if v8283 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1881
	}
L728:
	;
	if v7643 == int32(-1) {
		goto L1870
	} else {
		goto L1871
	}
L729:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8231 = m.ExcPending
	if v8231 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1866
	}
L730:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8212 = m.ExcPending
	if v8212 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1862
	}
L731:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8189 = m.ExcPending
	if v8189 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1858
	}
L732:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8170 = m.ExcPending
	if v8170 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1854
	}
L733:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8151 = m.ExcPending
	if v8151 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1850
	}
L734:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8132 = m.ExcPending
	if v8132 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1846
	}
L735:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8113 = m.ExcPending
	if v8113 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1842
	}
L736:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v8109 = m.ExcPending
	if v8109 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1841
	}
L737:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8094 = m.ExcPending
	if v8094 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1838
	}
L738:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8078 = m.ExcPending
	if v8078 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1834
	}
L739:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8062 = m.ExcPending
	if v8062 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1830
	}
L740:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8046 = m.ExcPending
	if v8046 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1827
	}
L741:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8030 = m.ExcPending
	if v8030 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1824
	}
L742:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8014 = m.ExcPending
	if v8014 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1821
	}
L743:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7998 = m.ExcPending
	if v7998 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1818
	}
L744:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7982 = m.ExcPending
	if v7982 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1815
	}
L745:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7966 = m.ExcPending
	if v7966 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1812
	}
L746:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7950 = m.ExcPending
	if v7950 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1808
	}
L747:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7934 = m.ExcPending
	if v7934 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1804
	}
L748:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7918 = m.ExcPending
	if v7918 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1800
	}
L749:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7902 = m.ExcPending
	if v7902 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1796
	}
L750:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7883 = m.ExcPending
	if v7883 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1792
	}
L751:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7867 = m.ExcPending
	if v7867 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1788
	}
L752:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7849 = m.ExcPending
	if v7849 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1784
	}
L753:
	;
	v3514 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[161])))
	if v3514 != 0 {
		goto L758
	} else {
		goto L759
	}
L754:
	;
	goto L755
L755:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7833 = m.ExcPending
	if v7833 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1780
	}
L756:
	;
	v3543 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[148]))
	if v3543 != 0 {
		goto L767
	} else {
		goto L768
	}
L757:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3532 = m.ExcPending
	if v3532 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L764
	}
L758:
	;
	v3516 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v3517 = *(*int32)(unsafe.Add(mBase, uint32(v3516)+20))
	goto L761
L759:
	;
	goto L760
L760:
	;
	goto L756
L761:
	;
	if base.B2i32(v3517 == int32(2)) == int32(0) {
		goto L757
	} else {
		goto L762
	}
L762:
	;
	v3523 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[162]))
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v3525 = m.ExcPending
	if v3525 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L763
	}
L763:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[163])) = v3523
	goto L760
L764:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMain_49), int32(0))
	mBase = m.M
	v3536 = m.ExcPending
	if v3536 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L765
	}
L765:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_50), int32(612), int32(_a_F_PostgresMain_51))
	mBase = m.M
	v3541 = m.ExcPending
	if v3541 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L766
	}
L766:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L767:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3545 = m.ExcPending
	if v3545 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L770
	}
L768:
	;
	goto L769
L769:
	;
	v3547 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[164]))
	if v3547 == int32(0) {
		goto L772
	} else {
		goto L773
	}
L770:
	;
	goto L769
L771:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v3564
	v3568 = v3481 + int32(428)
	v3570 = F_palloc0(m, int32(20))
	mBase = m.M
	v3571 = m.ExcPending
	if v3571 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L777
	}
L772:
	;
	v3552 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[45]))
	v3557 = F_AllocSetContextCreateInternal(m, v3552, int32(_a_F_PostgresMain_52), int32(0), int32(_a_F_PostgresMain_16), int32(_a_F_PostgresMain_17))
	mBase = m.M
	v3558 = m.ExcPending
	if v3558 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L775
	}
L773:
	;
	goto L774
L774:
	;
	F_MemoryContextReset(m, v3547)
	mBase = m.M
	v3561 = m.ExcPending
	if v3561 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L776
	}
L775:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[164])) = v3557
	v3564 = v3557
	goto L771
L776:
	;
	v3563 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[164]))
	v3564 = v3563
	goto L771
L777:
	;
	if v3568 != 0 {
		goto L779
	} else {
		goto L780
	}
L778:
	;
	v3595 = int32(0)
	base.MemoryFill(m, v3574, v3595, int32(96))
	v3598 = *(*int32)(unsafe.Add(mBase, uint32(v3568)))
	*(*int32)(unsafe.Add(mBase, uint32(v3598)+60)) = v3595
	v3601 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3598)+52)) = v3601
	*(*int32)(unsafe.Add(mBase, uint32(v3598)+44)) = v3595
	*(*int64)(unsafe.Add(mBase, uint32(v3598)+36)) = v3601
	*(*int64)(unsafe.Add(mBase, uint32(v3598)+4)) = v3601
	*(*int64)(unsafe.Add(mBase, uint32(v3598)+12)) = v3601
	*(*int32)(unsafe.Add(mBase, uint32(v3598)+20)) = v3595
	v3613 = *(*int32)(unsafe.Add(mBase, uint32(v3568)))
	*(*int32)(unsafe.Add(mBase, uint32(v3613))) = v3570
	v3615 = F_strlen(m, v3461)
	mBase = m.M
	v3617 = v3615 + int32(2)
	v3618 = F_palloc(m, v3617)
	mBase = m.M
	v3619 = m.ExcPending
	if v3619 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L789
	}
L779:
	;
	v3574 = F_palloc(m, int32(96))
	mBase = m.M
	v3575 = m.ExcPending
	if v3575 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L782
	}
L780:
	;
	v3580 = int32(28)
	goto L781
L781:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[23])) = v3580
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3585 = m.ExcPending
	if v3585 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L784
	}
L782:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3568))) = v3574
	if v3574 != 0 {
		goto L778
	} else {
		goto L783
	}
L783:
	;
	v3580 = int32(48)
	goto L781
L784:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMain_53), int32(0))
	mBase = m.M
	v3589 = m.ExcPending
	if v3589 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L785
	}
L785:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_54), int32(279), int32(_a_F_PostgresMain_55))
	mBase = m.M
	v3594 = m.ExcPending
	if v3594 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
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
	v3925 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+428))
	v3926 = m.G0
	v3928 = v3926 - int32(16)
	m.G0 = v3928
	v3932 = F_replication_yylex(m, v3928+int32(8), v3925)
	mBase = m.M
	v3933 = m.ExcPending
	if v3933 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L823
	}
L788:
	;
	F_yy_fatal_error_3(m, int32(_a_F_PostgresMain_56))
	mBase = m.M
	v3924 = m.ExcPending
	if v3924 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L821
	}
L789:
	;
	if v3618 != 0 {
		goto L790
	} else {
		goto L791
	}
L790:
	;
	if v3615 <= int32(0) {
		goto L793
	} else {
		goto L794
	}
L791:
	;
	goto L792
L792:
	;
	F_yy_fatal_error_3(m, int32(_a_F_PostgresMain_57))
	mBase = m.M
	v3921 = m.ExcPending
	if v3921 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L820
	}
L793:
	;
	v3823 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3615+v3618))) = uint16(v3823)
	if base.Ui32(v3617) < base.Ui32(int32(2)) {
		v3911 = v3823
		goto L807
	} else {
		goto L808
	}
L794:
	;
	v3623 = v3615 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v3615) {
		goto L795
	} else {
		goto L796
	}
L795:
	;
	v3639 = v3469
	v3653 = v3469
	goto L798
L796:
	;
	v3707 = v3469
	goto L797
L797:
	;
	v3746 = v3707
	v3747 = v2541
	goto L802
L798:
	;
	v3669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3461+v3639))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3639+v3618))) = uint8(v3669)
	v3672 = v3639 | int32(1)
	v3675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3672+v3461))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3618+v3672))) = uint8(v3675)
	v3678 = v3639 | int32(2)
	v3681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3678+v3461))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3618+v3678))) = uint8(v3681)
	v3684 = v3639 | int32(3)
	v3687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3684+v3461))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3618+v3684))) = uint8(v3687)
	v3689 = int32(4)
	v3690 = v3639 + v3689
	v3692 = v3653 + v3689
	if v3692 != v3615&int32(2147483644) {
		v3639 = v3690
		v3653 = v3692
		goto L798
	} else {
		goto L800
	}
L799:
	;
	if v3623 == int32(0) {
		goto L793
	} else {
		goto L801
	}
L800:
	;
	goto L799
L801:
	;
	v3707 = v3690
	goto L797
L802:
	;
	v3776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3461+v3746))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3746+v3618))) = uint8(v3776)
	v3778 = int32(1)
	v3781 = v3747 + v3778
	if v3781 != v3623 {
		v3746 = v3746 + v3778
		v3747 = v3781
		goto L802
	} else {
		goto L804
	}
L803:
	;
	goto L793
L804:
	;
	goto L803
L805:
	;
	if v3911 == int32(0) {
		goto L788
	} else {
		goto L819
	}
L806:
	;
	F_yy_fatal_error_3(m, int32(_a_F_PostgresMain_58))
	mBase = m.M
	v3914 = m.ExcPending
	if v3914 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L818
	}
L807:
	;
	goto L805
L808:
	;
	v3829 = v3617 - int32(2)
	v3831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3618+v3829))))
	if v3831 != 0 {
		v3911 = v3823
		goto L807
	} else {
		goto L809
	}
L809:
	;
	v3835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3617+v3618-int32(1)))))
	if v3835 != 0 {
		v3911 = v3823
		goto L807
	} else {
		goto L810
	}
L810:
	;
	v3837 = F_palloc(m, int32(48))
	mBase = m.M
	v3838 = m.ExcPending
	if v3838 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L811
	}
L811:
	;
	if v3837 == int32(0) {
		goto L806
	} else {
		goto L812
	}
L812:
	;
	v3841 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3837)+20)) = v3841
	*(*int32)(unsafe.Add(mBase, uint32(v3837)+8)) = v3618
	*(*int32)(unsafe.Add(mBase, uint32(v3837)+4)) = v3618
	*(*int32)(unsafe.Add(mBase, uint32(v3837)+12)) = v3829
	*(*int64)(unsafe.Add(mBase, uint32(v3837)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3837)+24)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v3837)+16)) = v3829
	*(*int32)(unsafe.Add(mBase, uint32(v3837))) = v3841
	F_replication_yyensure_buffer_stack(m, v3613)
	mBase = m.M
	v3854 = m.ExcPending
	if v3854 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L813
	}
L813:
	;
	v3855 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+20))
	v3856 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+12))
	v3860 = *(*int32)(unsafe.Add(mBase, uint32(v3855+v3856<<(uint(int32(2))%32))))
	if v3860 == v3837 {
		v3911 = v3837
		goto L807
	} else {
		goto L814
	}
L814:
	;
	if v3860 != 0 {
		goto L815
	} else {
		goto L816
	}
L815:
	;
	v3862 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+36))
	v3863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3613)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3862))) = uint8(v3863)
	v3865 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+20))
	v3866 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+12))
	v3867 = int32(2)
	v3870 = *(*int32)(unsafe.Add(mBase, uint32(v3865+v3866<<(uint(v3867)%32))))
	v3871 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v3870)+8)) = v3871
	v3873 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+20))
	v3874 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+12))
	v3878 = *(*int32)(unsafe.Add(mBase, uint32(v3873+v3874<<(uint(v3867)%32))))
	v3879 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v3878)+16)) = v3879
	v3881 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+20))
	v3882 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+12))
	v3883 = v3881
	v3884 = v3882
	goto L817
L816:
	;
	v3883 = v3855
	v3884 = v3856
	goto L817
L817:
	;
	v3885 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v3884<<(uint(v3885)%32)+v3883))) = v3837
	v3889 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+20))
	v3890 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+12))
	v3893 = v3889 + v3890<<(uint(v3885)%32)
	v3894 = *(*int32)(unsafe.Add(mBase, uint32(v3893)))
	v3895 = *(*int32)(unsafe.Add(mBase, uint32(v3894)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v3613)+28)) = v3895
	v3897 = *(*int32)(unsafe.Add(mBase, uint32(v3893)))
	v3898 = *(*int32)(unsafe.Add(mBase, uint32(v3897)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3613)+36)) = v3898
	*(*int32)(unsafe.Add(mBase, uint32(v3613)+80)) = v3898
	v3901 = *(*int32)(unsafe.Add(mBase, uint32(v3893)))
	v3902 = *(*int32)(unsafe.Add(mBase, uint32(v3901)))
	*(*int32)(unsafe.Add(mBase, uint32(v3613)+4)) = v3902
	v3904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3898))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3613)+24)) = uint8(v3904)
	*(*int32)(unsafe.Add(mBase, uint32(v3613)+48)) = int32(1)
	v3911 = v3837
	goto L807
L818:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L819:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3911)+20)) = int32(1)
	goto L787
L820:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L821:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L822:
	;
	m.G0 = v3928 + int32(16)
	v3948 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+428))
	if v3944 == int32(0) {
		goto L828
	} else {
		goto L829
	}
L823:
	;
	if base.Ui32(int32(9)) <= base.Ui32(v3932-int32(262)) {
		goto L824
	} else {
		goto L825
	}
L824:
	;
	if v3932 != int32(282) {
		v3944 = int32(0)
		goto L822
	} else {
		goto L827
	}
L825:
	;
	goto L826
L826:
	;
	v3941 = *(*int32)(unsafe.Add(mBase, uint32(v3925)))
	*(*int32)(unsafe.Add(mBase, uint32(v3941))) = v3932
	v3944 = int32(1)
	goto L822
L827:
	;
	goto L826
L828:
	;
	F_replication_scanner_finish(m, v3948)
	mBase = m.M
	v3952 = m.ExcPending
	if v3952 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L831
	}
L829:
	;
	goto L830
L830:
	;
	v3979 = m.G0
	v3981 = v3979 - int32(1872)
	m.G0 = v3981
	*(*int64)(unsafe.Add(mBase, uint32(v3981)+1864)) = int64(0)
	v3987 = v3981 - int32(-64)
	v3989 = v3981 + int32(1664)
	v3994 = v3989
	v4000 = v3469
	v4002 = v3987
	v4004 = v3987
	v4016 = v3989
	v4017 = int32(-2)
	v4021 = int32(200)
	goto L844
L831:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v3486
	v3956 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[164]))
	F_MemoryContextReset(m, v3956)
	mBase = m.M
	v3958 = m.ExcPending
	if v3958 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L832
	}
L832:
	;
	v3960 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[31]))
	if v3960 != 0 {
		goto L725
	} else {
		goto L833
	}
L833:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3964 = m.ExcPending
	if v3964 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L834
	}
L834:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3967 = m.ExcPending
	if v3967 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L835
	}
L835:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_59), int32(0))
	mBase = m.M
	v3971 = m.ExcPending
	if v3971 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L836
	}
L836:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(2169), int32(_a_F_PostgresMain_61))
	mBase = m.M
	v3976 = m.ExcPending
	if v3976 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L837
	}
L837:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L838:
	;
	if v4624 != 0 {
		goto L752
	} else {
		goto L1011
	}
L839:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4655 = m.ExcPending
	if v4655 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1007
	}
L840:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4636 = m.ExcPending
	if v4636 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1003
	}
L841:
	;
	if v3981+int32(1664) != v4614 {
		goto L999
	} else {
		goto L1000
	}
L842:
	;
	v4614 = v4047
	v4624 = int32(1)
	goto L841
L843:
	;
	F_replication_yyerror(m, int32(_a_F_PostgresMain_62))
	mBase = m.M
	v4612 = m.ExcPending
	if v4612 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L998
	}
L844:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v4016))) = uint8(v4000)
	if base.Ui32(v3994+v4021-int32(1)) <= base.Ui32(v4016) {
		goto L846
	} else {
		goto L847
	}
L845:
	;
	F_replication_yyerror(m, int32(_a_F_PostgresMain_63))
	mBase = m.M
	v4607 = m.ExcPending
	if v4607 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L997
	}
L846:
	;
	if int32(_a_F_PostgresMain_64) < v4021 {
		goto L843
	} else {
		goto L849
	}
L847:
	;
	v4078 = v3994
	v4080 = v4002
	v4081 = v4004
	v4084 = v4016
	v4085 = v4021
	goto L848
L848:
	;
	if v4000 == int32(34) {
		goto L866
	} else {
		goto L867
	}
L849:
	;
	v4037 = int32(_a_F_PostgresMain_41)
	v4039 = v4021 << (uint(int32(1)) % 32)
	if v4037 <= v4039 {
		goto L850
	} else {
		goto L851
	}
L850:
	;
	v4042 = v4037
	goto L852
L851:
	;
	v4042 = v4039
	goto L852
L852:
	;
	v4047 = F_palloc(m, v4042*int32(9)+int32(7))
	mBase = m.M
	v4048 = m.ExcPending
	if v4048 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L853
	}
L853:
	;
	if v4047 == int32(0) {
		goto L843
	} else {
		goto L854
	}
L854:
	;
	v4051 = v4016 - v3994
	v4053 = v4051 + int32(1)
	if v4053 != 0 {
		goto L855
	} else {
		goto L856
	}
L855:
	;
	base.MemoryCopy(m, v4047, v3994, v4053)
	goto L857
L856:
	;
	goto L857
L857:
	;
	v4058 = base.I32_div_s(v4042+int32(7), int32(8))
	v4059 = int32(3)
	v4061 = v4047 + v4058<<(uint(v4059)%32)
	v4063 = v4053 << (uint(v4059) % 32)
	if v4063 != 0 {
		goto L858
	} else {
		goto L859
	}
L858:
	;
	base.MemoryCopy(m, v4061, v4002, v4063)
	goto L860
L859:
	;
	goto L860
L860:
	;
	if v3981+int32(1664) != v3994 {
		goto L861
	} else {
		goto L862
	}
L861:
	;
	F_pfree(m, v3994)
	mBase = m.M
	v4069 = m.ExcPending
	if v4069 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L864
	}
L862:
	;
	goto L863
L863:
	;
	if v4042-int32(1) <= v4051 {
		goto L842
	} else {
		goto L865
	}
L864:
	;
	goto L863
L865:
	;
	v4078 = v4047
	v4080 = v4061
	v4081 = v4061 + v4063 - int32(8)
	v4084 = v4047 + v4051
	v4085 = v4042
	goto L848
L866:
	;
	v4614 = v4078
	v4624 = int32(0)
	goto L841
L867:
	;
	goto L868
L868:
	;
	v4091 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4000)+uint32(_c_F_PostgresMain[165]))))
	if v4091 == int32(-36) {
		v4127 = v4017
		goto L871
	} else {
		goto L872
	}
L869:
	;
	goto L845
L870:
	;
	v3994 = v4078
	v4000 = base.I32_extend8_s(v4601)
	v4002 = v4080
	v4004 = v4595
	v4016 = v4599 + int32(1)
	v4017 = v4600
	v4021 = v4085
	goto L844
L871:
	;
	v4130 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4000)+uint32(_c_F_PostgresMain[166]))))
	v4132 = v4130 & int32(255)
	if v4132 == int32(0) {
		goto L869
	} else {
		goto L887
	}
L872:
	;
	if v4017 == int32(-2) {
		goto L874
	} else {
		goto L875
	}
L873:
	;
	v4114 = v4091 + v4113
	if base.Ui32(int32(80)) < base.Ui32(v4114) {
		v4127 = v4112
		goto L871
	} else {
		goto L885
	}
L874:
	;
	v4098 = F_replication_yylex(m, v3981+int32(1864), v3948)
	mBase = m.M
	v4099 = m.ExcPending
	if v4099 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L877
	}
L875:
	;
	v4100 = v4017
	goto L876
L876:
	;
	if v4100 <= int32(0) {
		goto L878
	} else {
		goto L879
	}
L877:
	;
	v4100 = v4098
	goto L876
L878:
	;
	v4103 = int32(0)
	v4112 = v4103
	v4113 = v4103
	goto L873
L879:
	;
	goto L880
L880:
	;
	if v4100 == int32(256) {
		goto L881
	} else {
		goto L882
	}
L881:
	;
	v4614 = v4078
	v4624 = int32(1)
	goto L841
L882:
	;
	goto L883
L883:
	;
	if base.Ui32(int32(282)) < base.Ui32(v4100) {
		v4112 = v4100
		v4113 = int32(2)
		goto L873
	} else {
		goto L884
	}
L884:
	;
	v4111 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4100)+uint32(_c_F_PostgresMain[167]))))
	v4112 = v4100
	v4113 = v4111
	goto L873
L885:
	;
	v4117 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4114)+uint32(_c_F_PostgresMain[168]))))
	if v4113 != v4117 {
		v4127 = v4112
		goto L871
	} else {
		goto L886
	}
L886:
	;
	v4119 = *(*int64)(unsafe.Add(mBase, uint32(v3981)+1864))
	*(*int64)(unsafe.Add(mBase, uint32(v4081)+8)) = v4119
	v4124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4114)+uint32(_c_F_PostgresMain[169]))))
	v4595 = v4081 + int32(8)
	v4599 = v4084
	v4600 = int32(-2)
	v4601 = v4124
	goto L870
L887:
	;
	v4138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4130)+uint32(_c_F_PostgresMain[170]))))
	v4142 = v4081 + (int32(1)-v4138)<<(uint(int32(3))%32)
	v4143 = *(*int32)(unsafe.Add(mBase, uint32(v4142)))
	v4145 = int32(base.Ui32(v4143) >> (uint(int32(8)) % 32))
	v4146 = *(*int32)(unsafe.Add(mBase, uint32(v4142)+4))
	switch v4132 - int32(2) {
	case 0:
		goto L950
	default:
		v4556 = v4143
		v4557 = v4145
		goto L888
	case 14:
		goto L949
	case 15:
		goto L948
	case 16:
		goto L947
	case 17:
		goto L946
	case 18:
		goto L945
	case 19:
		goto L944
	case 20:
		goto L943
	case 21:
		goto L942
	case 22:
		goto L941
	case 23:
		goto L940
	case 24:
		goto L939
	case 25:
		goto L938
	case 26, 44, 46, 48, 53:
		goto L937
	case 27:
		goto L936
	case 28:
		goto L935
	case 29:
		goto L934
	case 30:
		goto L933
	case 31:
		goto L932
	case 32:
		goto L931
	case 33:
		goto L930
	case 34:
		goto L929
	case 35:
		goto L928
	case 36:
		goto L927
	case 37:
		goto L926
	case 38:
		goto L925
	case 41:
		goto L924
	case 42:
		goto L923
	case 43:
		goto L922
	case 45:
		goto L921
	case 47:
		goto L920
	case 49:
		goto L919
	case 50:
		goto L918
	case 51:
		goto L917
	case 52:
		goto L916
	case 54:
		goto L915
	case 55:
		goto L914
	case 56:
		goto L913
	case 57:
		goto L912
	case 58:
		goto L911
	case 59:
		goto L910
	case 60:
		goto L909
	case 61:
		goto L908
	case 62:
		goto L907
	case 63:
		goto L906
	case 64:
		goto L905
	case 65:
		goto L904
	case 66:
		goto L903
	case 67:
		goto L902
	case 68:
		goto L901
	case 69:
		goto L900
	case 70:
		goto L899
	case 71:
		goto L898
	case 72:
		goto L897
	case 73:
		goto L896
	case 74:
		goto L895
	case 75:
		goto L894
	case 76:
		goto L893
	case 77:
		goto L892
	case 78:
		goto L891
	case 79:
		goto L890
	case 80:
		goto L889
	}
L888:
	;
	v4560 = v4081 - v4138<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v4560)+12)) = v4146
	v4564 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v4560)+8)) = v4556&int32(255) | v4557<<(uint(v4564)%32)
	v4569 = v4560 + v4564
	v4570 = v4084 - v4138
	v4571 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4570))))
	v4574 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4130)+uint32(_c_F_PostgresMain[171]))))
	v4577 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4574)+uint32(_c_F_PostgresMain[172]))))
	v4578 = v4571 + v4577
	if base.Ui32(v4578) <= base.Ui32(int32(80)) {
		goto L993
	} else {
		goto L994
	}
L889:
	;
	v4556 = int32(_a_F_PostgresMain_65)
	v4557 = int32(336)
	goto L888
L890:
	;
	v4556 = int32(_a_F_PostgresMain_66)
	v4557 = int32(374)
	goto L888
L891:
	;
	v4556 = int32(_a_F_PostgresMain_67)
	v4557 = int32(374)
	goto L888
L892:
	;
	v4556 = int32(_a_F_PostgresMain_68)
	v4557 = int32(374)
	goto L888
L893:
	;
	v4556 = int32(_a_F_PostgresMain_69)
	v4557 = int32(1526)
	goto L888
L894:
	;
	v4556 = int32(_a_F_PostgresMain_70)
	v4557 = int32(73)
	goto L888
L895:
	;
	v4556 = int32(_a_F_PostgresMain_71)
	v4557 = int32(1303)
	goto L888
L896:
	;
	v4556 = int32(_a_F_PostgresMain_72)
	v4557 = int32(372)
	goto L888
L897:
	;
	v4556 = int32(_a_F_PostgresMain_73)
	v4557 = int32(1327)
	goto L888
L898:
	;
	v4556 = int32(_a_F_PostgresMain_74)
	v4557 = int32(1326)
	goto L888
L899:
	;
	v4556 = int32(_a_F_PostgresMain_75)
	v4557 = int32(1575)
	goto L888
L900:
	;
	v4556 = int32(_a_F_PostgresMain_76)
	v4557 = int32(445)
	goto L888
L901:
	;
	v4556 = int32(_a_F_PostgresMain_77)
	v4557 = int32(53)
	goto L888
L902:
	;
	v4556 = int32(_a_F_PostgresMain_78)
	v4557 = int32(367)
	goto L888
L903:
	;
	v4556 = int32(_a_F_PostgresMain_79)
	v4557 = int32(368)
	goto L888
L904:
	;
	v4556 = int32(_a_F_PostgresMain_80)
	v4557 = int32(368)
	goto L888
L905:
	;
	v4556 = int32(_a_F_PostgresMain_81)
	v4557 = int32(1127)
	goto L888
L906:
	;
	v4556 = int32(_a_F_PostgresMain_82)
	v4557 = int32(137)
	goto L888
L907:
	;
	v4556 = int32(_a_F_PostgresMain_83)
	v4557 = int32(1223)
	goto L888
L908:
	;
	v4556 = int32(_a_F_PostgresMain_84)
	v4557 = int32(984)
	goto L888
L909:
	;
	v4512 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4556 = v4512
	v4557 = int32(base.Ui32(v4512) >> (uint(int32(8)) % 32))
	goto L888
L910:
	;
	v4503 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(8))))
	v4504 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4505 = F_makeInteger(m, v4504)
	mBase = m.M
	v4506 = m.ExcPending
	if v4506 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L990
	}
L911:
	;
	v4492 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(8))))
	v4493 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4494 = F_makeString(m, v4493)
	mBase = m.M
	v4495 = m.ExcPending
	if v4495 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L988
	}
L912:
	;
	v4481 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(8))))
	v4482 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4483 = F_makeString(m, v4482)
	mBase = m.M
	v4484 = m.ExcPending
	if v4484 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L986
	}
L913:
	;
	v4472 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4475 = F_makeDefElem(m, v4472, int32(0), int32(-1))
	mBase = m.M
	v4476 = m.ExcPending
	if v4476 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L985
	}
L914:
	;
	v4462 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	*(*int32)(unsafe.Add(mBase, uint32(v3981)+52)) = v4462
	*(*int32)(unsafe.Add(mBase, uint32(v3981)+56)) = v4462
	v4468 = F_list_make1_impl(m, int32(1), v3981+int32(52))
	mBase = m.M
	v4469 = m.ExcPending
	if v4469 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L984
	}
L915:
	;
	v4456 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(16))))
	v4457 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4458 = F_lappend(m, v4456, v4457)
	mBase = m.M
	v4459 = m.ExcPending
	if v4459 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L983
	}
L916:
	;
	v4449 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4450 = F_makeString(m, v4449)
	mBase = m.M
	v4451 = m.ExcPending
	if v4451 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L982
	}
L917:
	;
	v4442 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(8))))
	v4443 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4445 = F_makeDefElem(m, v4442, v4443, int32(-1))
	mBase = m.M
	v4446 = m.ExcPending
	if v4446 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L981
	}
L918:
	;
	v4434 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(16))))
	v4435 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4436 = F_lappend(m, v4434, v4435)
	mBase = m.M
	v4437 = m.ExcPending
	if v4437 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L980
	}
L919:
	;
	v4422 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	*(*int32)(unsafe.Add(mBase, uint32(v3981)+48)) = v4422
	*(*int32)(unsafe.Add(mBase, uint32(v3981)+60)) = v4422
	v4428 = F_list_make1_impl(m, int32(1), v3981+int32(48))
	mBase = m.M
	v4429 = m.ExcPending
	if v4429 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L979
	}
L920:
	;
	v4417 = int32(8)
	v4419 = *(*int32)(unsafe.Add(mBase, uint32(v4081-v4417)))
	v4556 = v4419
	v4557 = int32(base.Ui32(v4419) >> (uint(v4417) % 32))
	goto L888
L921:
	;
	v4412 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	if v4412 == int32(0) {
		goto L839
	} else {
		goto L978
	}
L922:
	;
	v4409 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4556 = v4409
	v4557 = int32(base.Ui32(v4409) >> (uint(int32(8)) % 32))
	goto L888
L923:
	;
	v4556 = int32(0)
	v4557 = v4145
	goto L888
L924:
	;
	v4556 = int32(1)
	v4557 = v4145
	goto L888
L925:
	;
	v4401 = F_palloc0(m, int32(4))
	mBase = m.M
	v4402 = m.ExcPending
	if v4402 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L977
	}
L926:
	;
	v4388 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	if v4388 == int32(0) {
		goto L840
	} else {
		goto L975
	}
L927:
	;
	v4372 = F_palloc0(m, int32(32))
	mBase = m.M
	v4373 = m.ExcPending
	if v4373 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L974
	}
L928:
	;
	v4355 = F_palloc0(m, int32(32))
	mBase = m.M
	v4356 = m.ExcPending
	if v4356 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L973
	}
L929:
	;
	v4340 = F_palloc0(m, int32(12))
	mBase = m.M
	v4341 = m.ExcPending
	if v4341 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L972
	}
L930:
	;
	v4327 = F_palloc0(m, int32(12))
	mBase = m.M
	v4328 = m.ExcPending
	if v4328 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L971
	}
L931:
	;
	v4316 = F_palloc0(m, int32(12))
	mBase = m.M
	v4317 = m.ExcPending
	if v4317 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L970
	}
L932:
	;
	v4308 = F_makeBoolean(m, int32(1))
	mBase = m.M
	v4309 = m.ExcPending
	if v4309 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L968
	}
L933:
	;
	v4299 = F_makeBoolean(m, int32(1))
	mBase = m.M
	v4300 = m.ExcPending
	if v4300 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L966
	}
L934:
	;
	v4290 = F_makeString(m, int32(_a_F_PostgresMain_85))
	mBase = m.M
	v4291 = m.ExcPending
	if v4291 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L964
	}
L935:
	;
	v4281 = F_makeString(m, int32(_a_F_PostgresMain_86))
	mBase = m.M
	v4282 = m.ExcPending
	if v4282 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L962
	}
L936:
	;
	v4272 = F_makeString(m, int32(_a_F_PostgresMain_87))
	mBase = m.M
	v4273 = m.ExcPending
	if v4273 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L960
	}
L937:
	;
	v4268 = int32(0)
	v4556 = v4268
	v4557 = v4268
	goto L888
L938:
	;
	v4262 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(8))))
	v4263 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4264 = F_lappend(m, v4262, v4263)
	mBase = m.M
	v4265 = m.ExcPending
	if v4265 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L959
	}
L939:
	;
	v4257 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4556 = v4257
	v4557 = int32(base.Ui32(v4257) >> (uint(int32(8)) % 32))
	goto L888
L940:
	;
	v4252 = int32(8)
	v4254 = *(*int32)(unsafe.Add(mBase, uint32(v4081-v4252)))
	v4556 = v4254
	v4557 = int32(base.Ui32(v4254) >> (uint(v4252) % 32))
	goto L888
L941:
	;
	v4230 = F_palloc0(m, int32(24))
	mBase = m.M
	v4231 = m.ExcPending
	if v4231 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L958
	}
L942:
	;
	v4211 = F_palloc0(m, int32(24))
	mBase = m.M
	v4212 = m.ExcPending
	if v4212 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L957
	}
L943:
	;
	v4204 = F_palloc0(m, int32(8))
	mBase = m.M
	v4205 = m.ExcPending
	if v4205 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L956
	}
L944:
	;
	v4193 = F_palloc0(m, int32(8))
	mBase = m.M
	v4194 = m.ExcPending
	if v4194 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L955
	}
L945:
	;
	v4183 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(16))))
	v4184 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	*(*int32)(unsafe.Add(mBase, uint32(v3981)+4)) = v4184
	*(*int32)(unsafe.Add(mBase, uint32(v3981))) = v4183
	v4188 = F_psprintf(m, int32(_a_F_PostgresMain_88), v3981)
	mBase = m.M
	v4189 = m.ExcPending
	if v4189 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L954
	}
L946:
	;
	v4178 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4556 = v4178
	v4557 = int32(base.Ui32(v4178) >> (uint(int32(8)) % 32))
	goto L888
L947:
	;
	v4170 = F_palloc0(m, int32(8))
	mBase = m.M
	v4171 = m.ExcPending
	if v4171 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L953
	}
L948:
	;
	v4161 = F_palloc0(m, int32(8))
	mBase = m.M
	v4162 = m.ExcPending
	if v4162 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L952
	}
L949:
	;
	v4154 = F_palloc0(m, int32(4))
	mBase = m.M
	v4155 = m.ExcPending
	if v4155 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L951
	}
L950:
	;
	v4151 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v3481+int32(424)))) = v4151
	v4556 = v4143
	v4557 = v4145
	goto L888
L951:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4154))) = int32(454)
	v4556 = v4154
	v4557 = int32(base.Ui32(v4154) >> (uint(int32(8)) % 32))
	goto L888
L952:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4161))) = int32(460)
	v4165 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	*(*int32)(unsafe.Add(mBase, uint32(v4161)+4)) = v4165
	v4556 = v4161
	v4557 = int32(base.Ui32(v4161) >> (uint(int32(8)) % 32))
	goto L888
L953:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4170))) = int32(159)
	v4174 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	*(*int32)(unsafe.Add(mBase, uint32(v4170)+4)) = v4174
	v4556 = v4170
	v4557 = int32(base.Ui32(v4170) >> (uint(int32(8)) % 32))
	goto L888
L954:
	;
	v4556 = v4188
	v4557 = int32(base.Ui32(v4188) >> (uint(int32(8)) % 32))
	goto L888
L955:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4193))) = int32(455)
	v4197 = int32(8)
	v4199 = *(*int32)(unsafe.Add(mBase, uint32(v4081-v4197)))
	*(*int32)(unsafe.Add(mBase, uint32(v4193)+4)) = v4199
	v4556 = v4193
	v4557 = int32(base.Ui32(v4193) >> (uint(v4197) % 32))
	goto L888
L956:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4204))) = int32(455)
	v4556 = v4204
	v4557 = int32(base.Ui32(v4204) >> (uint(int32(8)) % 32))
	goto L888
L957:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4211)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4211))) = int32(456)
	v4219 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v4211)+4)) = v4219
	v4223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4081-int32(16)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4211)+16)) = uint8(v4223)
	v4225 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	*(*int32)(unsafe.Add(mBase, uint32(v4211)+20)) = v4225
	v4556 = v4211
	v4557 = int32(base.Ui32(v4211) >> (uint(int32(8)) % 32))
	goto L888
L958:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4230)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4230))) = int32(456)
	v4238 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v4230)+4)) = v4238
	v4242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4081-int32(24)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4230)+16)) = uint8(v4242)
	v4244 = int32(8)
	v4246 = *(*int32)(unsafe.Add(mBase, uint32(v4081-v4244)))
	*(*int32)(unsafe.Add(mBase, uint32(v4230)+12)) = v4246
	v4248 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	*(*int32)(unsafe.Add(mBase, uint32(v4230)+20)) = v4248
	v4556 = v4230
	v4557 = int32(base.Ui32(v4230) >> (uint(v4244) % 32))
	goto L888
L959:
	;
	v4556 = v4264
	v4557 = int32(base.Ui32(v4264) >> (uint(int32(8)) % 32))
	goto L888
L960:
	;
	v4275 = F_makeDefElem(m, int32(_a_F_PostgresMain_89), v4272, int32(-1))
	mBase = m.M
	v4276 = m.ExcPending
	if v4276 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L961
	}
L961:
	;
	v4556 = v4275
	v4557 = int32(base.Ui32(v4275) >> (uint(int32(8)) % 32))
	goto L888
L962:
	;
	v4284 = F_makeDefElem(m, int32(_a_F_PostgresMain_89), v4281, int32(-1))
	mBase = m.M
	v4285 = m.ExcPending
	if v4285 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L963
	}
L963:
	;
	v4556 = v4284
	v4557 = int32(base.Ui32(v4284) >> (uint(int32(8)) % 32))
	goto L888
L964:
	;
	v4293 = F_makeDefElem(m, int32(_a_F_PostgresMain_89), v4290, int32(-1))
	mBase = m.M
	v4294 = m.ExcPending
	if v4294 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L965
	}
L965:
	;
	v4556 = v4293
	v4557 = int32(base.Ui32(v4293) >> (uint(int32(8)) % 32))
	goto L888
L966:
	;
	v4302 = F_makeDefElem(m, int32(_a_F_PostgresMain_71), v4299, int32(-1))
	mBase = m.M
	v4303 = m.ExcPending
	if v4303 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L967
	}
L967:
	;
	v4556 = v4302
	v4557 = int32(base.Ui32(v4302) >> (uint(int32(8)) % 32))
	goto L888
L968:
	;
	v4311 = F_makeDefElem(m, int32(_a_F_PostgresMain_69), v4308, int32(-1))
	mBase = m.M
	v4312 = m.ExcPending
	if v4312 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L969
	}
L969:
	;
	v4556 = v4311
	v4557 = int32(base.Ui32(v4311) >> (uint(int32(8)) % 32))
	goto L888
L970:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4316))) = int32(457)
	v4320 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	v4321 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4316)+8)) = uint8(v4321)
	*(*int32)(unsafe.Add(mBase, uint32(v4316)+4)) = v4320
	v4556 = v4316
	v4557 = int32(base.Ui32(v4316) >> (uint(int32(8)) % 32))
	goto L888
L971:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4327))) = int32(457)
	v4331 = int32(8)
	v4333 = *(*int32)(unsafe.Add(mBase, uint32(v4081-v4331)))
	v4334 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4327)+8)) = uint8(v4334)
	*(*int32)(unsafe.Add(mBase, uint32(v4327)+4)) = v4333
	v4556 = v4327
	v4557 = int32(base.Ui32(v4327) >> (uint(v4331) % 32))
	goto L888
L972:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4340))) = int32(458)
	v4346 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v4340)+4)) = v4346
	v4348 = int32(8)
	v4350 = *(*int32)(unsafe.Add(mBase, uint32(v4081-v4348)))
	*(*int32)(unsafe.Add(mBase, uint32(v4340)+8)) = v4350
	v4556 = v4340
	v4557 = int32(base.Ui32(v4340) >> (uint(v4348) % 32))
	goto L888
L973:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4355))) = int64(459)
	v4361 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v4355)+8)) = v4361
	v4363 = int32(8)
	v4365 = *(*int64)(unsafe.Add(mBase, uint32(v4081-v4363)))
	*(*int64)(unsafe.Add(mBase, uint32(v4355)+16)) = v4365
	v4367 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	*(*int32)(unsafe.Add(mBase, uint32(v4355)+12)) = v4367
	v4556 = v4355
	v4557 = int32(base.Ui32(v4355) >> (uint(v4363) % 32))
	goto L888
L974:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4372))) = int64(4294967755)
	v4378 = *(*int32)(unsafe.Add(mBase, uint32(v4081-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v4372)+8)) = v4378
	v4380 = int32(8)
	v4382 = *(*int64)(unsafe.Add(mBase, uint32(v4081-v4380)))
	*(*int64)(unsafe.Add(mBase, uint32(v4372)+16)) = v4382
	v4384 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	*(*int32)(unsafe.Add(mBase, uint32(v4372)+24)) = v4384
	v4556 = v4372
	v4557 = int32(base.Ui32(v4372) >> (uint(v4380) % 32))
	goto L888
L975:
	;
	v4392 = F_palloc0(m, int32(8))
	mBase = m.M
	v4393 = m.ExcPending
	if v4393 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L976
	}
L976:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4392))) = int32(461)
	v4396 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	*(*int32)(unsafe.Add(mBase, uint32(v4392)+4)) = v4396
	v4556 = v4392
	v4557 = int32(base.Ui32(v4392) >> (uint(int32(8)) % 32))
	goto L888
L977:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4401))) = int32(462)
	v4556 = v4401
	v4557 = int32(base.Ui32(v4401) >> (uint(int32(8)) % 32))
	goto L888
L978:
	;
	v4556 = v4412
	v4557 = int32(base.Ui32(v4412) >> (uint(int32(8)) % 32))
	goto L888
L979:
	;
	v4556 = v4428
	v4557 = int32(base.Ui32(v4428) >> (uint(int32(8)) % 32))
	goto L888
L980:
	;
	v4556 = v4436
	v4557 = int32(base.Ui32(v4436) >> (uint(int32(8)) % 32))
	goto L888
L981:
	;
	v4556 = v4445
	v4557 = int32(base.Ui32(v4445) >> (uint(int32(8)) % 32))
	goto L888
L982:
	;
	v4556 = v4450
	v4557 = int32(base.Ui32(v4450) >> (uint(int32(8)) % 32))
	goto L888
L983:
	;
	v4556 = v4458
	v4557 = int32(base.Ui32(v4458) >> (uint(int32(8)) % 32))
	goto L888
L984:
	;
	v4556 = v4468
	v4557 = int32(base.Ui32(v4468) >> (uint(int32(8)) % 32))
	goto L888
L985:
	;
	v4556 = v4475
	v4557 = int32(base.Ui32(v4475) >> (uint(int32(8)) % 32))
	goto L888
L986:
	;
	v4486 = F_makeDefElem(m, v4481, v4483, int32(-1))
	mBase = m.M
	v4487 = m.ExcPending
	if v4487 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L987
	}
L987:
	;
	v4556 = v4486
	v4557 = int32(base.Ui32(v4486) >> (uint(int32(8)) % 32))
	goto L888
L988:
	;
	v4497 = F_makeDefElem(m, v4492, v4494, int32(-1))
	mBase = m.M
	v4498 = m.ExcPending
	if v4498 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L989
	}
L989:
	;
	v4556 = v4497
	v4557 = int32(base.Ui32(v4497) >> (uint(int32(8)) % 32))
	goto L888
L990:
	;
	v4508 = F_makeDefElem(m, v4503, v4505, int32(-1))
	mBase = m.M
	v4509 = m.ExcPending
	if v4509 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L991
	}
L991:
	;
	v4556 = v4508
	v4557 = int32(base.Ui32(v4508) >> (uint(int32(8)) % 32))
	goto L888
L992:
	;
	v4590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4578)+uint32(_c_F_PostgresMain[169]))))
	v4595 = v4569
	v4599 = v4570
	v4600 = v4127
	v4601 = v4590
	goto L870
L993:
	;
	v4581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4578)+uint32(_c_F_PostgresMain[168]))))
	if v4581 == v4571&int32(255) {
		goto L992
	} else {
		goto L996
	}
L994:
	;
	goto L995
L995:
	;
	v4587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4574)+uint32(_c_F_PostgresMain[173]))))
	v4595 = v4569
	v4599 = v4570
	v4600 = v4127
	v4601 = v4587
	goto L870
L996:
	;
	goto L995
L997:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L998:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L999:
	;
	F_pfree(m, v4614)
	mBase = m.M
	v4629 = m.ExcPending
	if v4629 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1002
	}
L1000:
	;
	goto L1001
L1001:
	;
	m.G0 = v3981 + int32(1872)
	goto L838
L1002:
	;
	goto L1001
L1003:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4639 = m.ExcPending
	if v4639 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1004
	}
L1004:
	;
	v4640 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	*(*int32)(unsafe.Add(mBase, uint32(v3981)+16)) = v4640
	F_errmsg(m, int32(_a_F_PostgresMain_90), v3981+int32(16))
	mBase = m.M
	v4646 = m.ExcPending
	if v4646 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1005
	}
L1005:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_91), int32(322), int32(_a_F_PostgresMain_92))
	mBase = m.M
	v4651 = m.ExcPending
	if v4651 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1006
	}
L1006:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1007:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4658 = m.ExcPending
	if v4658 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1008
	}
L1008:
	;
	v4659 = *(*int32)(unsafe.Add(mBase, uint32(v4081)))
	*(*int32)(unsafe.Add(mBase, uint32(v3981)+32)) = v4659
	F_errmsg(m, int32(_a_F_PostgresMain_90), v3981+int32(32))
	mBase = m.M
	v4665 = m.ExcPending
	if v4665 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1009
	}
L1009:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_91), int32(363), int32(_a_F_PostgresMain_92))
	mBase = m.M
	v4670 = m.ExcPending
	if v4670 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
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
	v4671 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+428))
	F_replication_scanner_finish(m, v4671)
	mBase = m.M
	v4673 = m.ExcPending
	if v4673 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1012
	}
L1012:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[110])) = v3461
	F_pgstat_report_activity(m, int32(3), v3461)
	mBase = m.M
	v4681 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[174])))
	if v4681 != 0 {
		goto L1013
	} else {
		goto L1014
	}
L1013:
	;
	v4682 = int32(15)
	goto L1015
L1014:
	;
	v4682 = int32(14)
	goto L1015
L1015:
	;
	v4684 = F_errstart(m, v4682, int32(0))
	mBase = m.M
	v4685 = m.ExcPending
	if v4685 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1016
	}
L1016:
	;
	if v4684 != 0 {
		goto L1017
	} else {
		goto L1018
	}
L1017:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+400)) = v3461
	F_errmsg(m, int32(_a_F_PostgresMain_93), v3481+int32(400))
	mBase = m.M
	v4691 = m.ExcPending
	if v4691 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1020
	}
L1018:
	;
	goto L1019
L1019:
	;
	v4698 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v4699 = *(*int32)(unsafe.Add(mBase, uint32(v4698)+24))
	goto L1022
L1020:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(2200), int32(_a_F_PostgresMain_61))
	mBase = m.M
	v4696 = m.ExcPending
	if v4696 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1021
	}
L1021:
	;
	goto L1019
L1022:
	;
	if (v4699-int32(7))&int32(-9) == int32(0) {
		goto L751
	} else {
		goto L1023
	}
L1023:
	;
	v4707 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[148]))
	if v4707 != 0 {
		goto L1024
	} else {
		goto L1025
	}
L1024:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4709 = m.ExcPending
	if v4709 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1027
	}
L1025:
	;
	goto L1026
L1026:
	;
	F_initStringInfo(m, int32(_a_F_PostgresMain_94))
	mBase = m.M
	v4712 = m.ExcPending
	if v4712 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1028
	}
L1027:
	;
	goto L1026
L1028:
	;
	F_initStringInfo(m, int32(_a_F_PostgresMain_95))
	mBase = m.M
	v4715 = m.ExcPending
	if v4715 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1029
	}
L1029:
	;
	F_initStringInfo(m, int32(_a_F_PostgresMain_96))
	mBase = m.M
	v4718 = m.ExcPending
	if v4718 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1030
	}
L1030:
	;
	v4719 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+424))
	v4720 = *(*int32)(unsafe.Add(mBase, uint32(v4719)))
	switch v4720 - int32(454) {
	case 0:
		goto L1040
	case 1:
		goto L1031
	case 2:
		goto L1038
	case 3:
		goto L1037
	case 4:
		goto L1036
	case 5:
		goto L1035
	case 6:
		goto L1039
	case 7:
		goto L1034
	case 8:
		goto L1033
	default:
		goto L1032
	}
L1031:
	;
	v7820 = int32(_a_F_PostgresMain_97)
	F_PreventInTransactionBlock(m, int32(1), v7820)
	mBase = m.M
	v7824 = m.ExcPending
	if v7824 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1778
	}
L1032:
	;
	if v4720 == int32(159) {
		goto L727
	} else {
		goto L1774
	}
L1033:
	;
	F_PreventInTransactionBlock(m, int32(1), int32(_a_F_PostgresMain_98))
	mBase = m.M
	v7388 = m.ExcPending
	if v7388 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1685
	}
L1034:
	;
	F_PreventInTransactionBlock(m, int32(1), int32(_a_F_PostgresMain_99))
	mBase = m.M
	v7072 = m.ExcPending
	if v7072 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1637
	}
L1035:
	;
	v6367 = int32(_a_F_PostgresMain_100)
	F_PreventInTransactionBlock(m, int32(1), v6367)
	mBase = m.M
	v6371 = m.ExcPending
	if v6371 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1460
	}
L1036:
	;
	v5978 = int32(0)
	v5979 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+424))
	v5980 = *(*int32)(unsafe.Add(mBase, uint32(v5979)+8))
	if v5980 == v5978 {
		v6143 = v3469
		v6147 = v3469
		v6148 = v3469
		v6158 = v5978
		goto L1362
	} else {
		goto L1363
	}
L1037:
	;
	v5968 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+424))
	v5969 = *(*int32)(unsafe.Add(mBase, uint32(v5968)+4))
	v5970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5968)+8)))
	F_ReplicationSlotDrop(m, v5969, (v5970^int32(-1))&int32(1))
	mBase = m.M
	v5976 = m.ExcPending
	if v5976 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1361
	}
L1038:
	;
	v5218 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+424))
	v5219 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+464)) = v5219
	v5221 = *(*int32)(unsafe.Add(mBase, uint32(v5218)+20))
	if v5221 == v5219 {
		v5548 = v3469
		v5559 = v3469
		v5564 = v2541
		v5572 = v3469
		goto L1165
	} else {
		goto L1166
	}
L1039:
	;
	v4972 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+424))
	v4973 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+480)) = v4973
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+472)) = v4973
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+464)) = v4973
	v4980 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v4981 = m.ExcPending
	if v4981 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1106
	}
L1040:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+464)) = int32(0)
	v4726 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[175]))
	v4727 = *(*int64)(unsafe.Add(mBase, uint32(v4726)))
	goto L1041
L1041:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+32)) = v4727
	v4731 = int32(32)
	v4735 = F_pg_snprintf(m, v3481+int32(_a_F_PostgresMain_101), v4731, int32(_a_F_PostgresMain_102), v3481+v4731)
	mBase = m.M
	v4736 = m.ExcPending
	if v4736 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1042
	}
L1042:
	;
	v4740 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])))
	if v4740 == int32(1) {
		goto L1044
	} else {
		goto L1045
	}
L1043:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[36])) = uint8(v4750)
	if v4750 != 0 {
		goto L1048
	} else {
		goto L1049
	}
L1044:
	;
	v4745 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v4746 = *(*int32)(unsafe.Add(mBase, uint32(v4745)+308))
	v4748 = base.B2i32(v4746 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])) = uint8(v4748)
	v4750 = v4748
	goto L1046
L1045:
	;
	v4750 = int32(0)
	goto L1046
L1046:
	;
	goto L1043
L1047:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v3481)+20)) = uint32(v4799)
	v4802 = int64(base.Ui64(v4799) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3481)+16)) = uint32(v4802)
	v4810 = F_pg_snprintf(m, v3481+int32(496), int32(64), int32(_a_F_PostgresMain_103), v3481+int32(16))
	mBase = m.M
	v4811 = m.ExcPending
	if v4811 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1063
	}
L1048:
	;
	v4755 = F_GetWalRcvFlushRecPtr(m, int32(0), v3481+int32(_a_F_PostgresMain_104))
	mBase = m.M
	v4756 = m.ExcPending
	if v4756 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1051
	}
L1049:
	;
	goto L1050
L1050:
	;
	v4769 = v3481 + int32(440)
	v4771 = int32(_a_F_PostgresMain_105)
	v4772 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v4773 = int64(0)
	v4776 = base.AtomicRmwCmpxchg64(m, v4772, int32(272), v4773, v4773)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[176])) = v4776
	v4778 = int32(0)
	v4781 = base.AtomicRmwOr32(m, v4778, int32(_a_F_PostgresMain_106), v4778)
	v4784 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v4788 = base.AtomicRmwCmpxchg64(m, v4784, int32(264), v4773, v4773)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[177])) = v4788
	if v4769 != 0 {
		goto L1060
	} else {
		goto L1061
	}
L1051:
	;
	v4759 = F_GetXLogReplayRecPtr(m, v3481+int32(496))
	mBase = m.M
	v4760 = m.ExcPending
	if v4760 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1052
	}
L1052:
	;
	v4761 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+496))
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+440)) = v4761
	if base.Ui64(v4759) < base.Ui64(v4755) {
		goto L1053
	} else {
		goto L1054
	}
L1053:
	;
	v4764 = v4755
	goto L1055
L1054:
	;
	v4764 = v4759
	goto L1055
L1055:
	;
	v4765 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[178])))
	if v4761 == v4765 {
		goto L1056
	} else {
		goto L1057
	}
L1056:
	;
	v4767 = v4764
	goto L1058
L1057:
	;
	v4767 = v4759
	goto L1058
L1058:
	;
	v4799 = v4767
	goto L1047
L1059:
	;
	v4799 = v4795
	goto L1047
L1060:
	;
	v4791 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v4792 = *(*int32)(unsafe.Add(mBase, uint32(v4791)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v4769))) = v4792
	goto L1062
L1061:
	;
	goto L1062
L1062:
	;
	v4795 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[176]))
	goto L1059
L1063:
	;
	v4813 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[31]))
	if v4813 != 0 {
		goto L1064
	} else {
		goto L1065
	}
L1064:
	;
	v4815 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50]))
	F_StartTransactionCommand(m)
	mBase = m.M
	v4817 = m.ExcPending
	if v4817 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1067
	}
L1065:
	;
	v4826 = v3469
	goto L1066
L1066:
	;
	v4828 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v4829 = m.ExcPending
	if v4829 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1071
	}
L1067:
	;
	v4819 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[31]))
	v4820 = F_get_database_name(m, v4819)
	mBase = m.M
	v4821 = m.ExcPending
	if v4821 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1068
	}
L1068:
	;
	v4822 = F_MemoryContextStrdup(m, v4815, v4820)
	mBase = m.M
	v4823 = m.ExcPending
	if v4823 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1069
	}
L1069:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v4825 = m.ExcPending
	if v4825 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1070
	}
L1070:
	;
	v4826 = v4822
	goto L1066
L1071:
	;
	v4831 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v4832 = m.ExcPending
	if v4832 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1072
	}
L1072:
	;
	F_TupleDescInitBuiltinEntry(m, v4831, int32(1), int32(_a_F_PostgresMain_107), int32(25))
	mBase = m.M
	v4837 = m.ExcPending
	if v4837 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1073
	}
L1073:
	;
	F_TupleDescInitBuiltinEntry(m, v4831, int32(2), int32(_a_F_PostgresMain_75), int32(20))
	mBase = m.M
	v4842 = m.ExcPending
	if v4842 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1074
	}
L1074:
	;
	F_TupleDescInitBuiltinEntry(m, v4831, int32(3), int32(_a_F_PostgresMain_108), int32(25))
	mBase = m.M
	v4847 = m.ExcPending
	if v4847 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1075
	}
L1075:
	;
	F_TupleDescInitBuiltinEntry(m, v4831, int32(4), int32(_a_F_PostgresMain_109), int32(25))
	mBase = m.M
	v4852 = m.ExcPending
	if v4852 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1076
	}
L1076:
	;
	v4853 = int32(0)
	v4862 = *(*int32)(unsafe.Add(mBase, uint32(v4831)))
	if v4853 < v4862 {
		goto L1078
	} else {
		goto L1079
	}
L1077:
	;
	v4941 = F_begin_tup_output_tupdesc(m, v4828, v4831, int32(_a_F_PostgresMain_110))
	mBase = m.M
	v4942 = m.ExcPending
	if v4942 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1096
	}
L1078:
	;
	v4866 = v4831 + int32(28)
	v4873 = v4853
	v4874 = v4862
	v4876 = v4853
	goto L1082
L1079:
	;
	v4930 = v4853
	v4937 = v4862
	goto L1080
L1080:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4831)+20)) = v4937
	*(*int32)(unsafe.Add(mBase, uint32(v4831)+16)) = v4930
	goto L1077
L1081:
	;
	v4930 = v4924
	v4937 = v4903
	goto L1080
L1082:
	;
	v4882 = v4866 + v4862<<(uint(int32(3))%32) + v4873*int32(100)
	v4885 = v4866 + v4873<<(uint(int32(3))%32)
	if v4862 != v4874 {
		v4903 = v4874
		goto L1084
	} else {
		goto L1085
	}
L1083:
	;
	v4924 = v4862
	goto L1081
L1084:
	;
	v4904 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4885)+2)))
	if v4904 <= int32(0) {
		v4924 = v4873
		goto L1081
	} else {
		goto L1092
	}
L1085:
	;
	v4887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4885)+7)))
	if v4887 != int32(118) {
		goto L1086
	} else {
		goto L1087
	}
L1086:
	;
	v4903 = v4873
	goto L1084
L1087:
	;
	v4890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4885)+4)))
	if v4890 != int32(1) {
		goto L1086
	} else {
		goto L1088
	}
L1088:
	;
	v4893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4885)+6)))
	if v4893&int32(6) != 0 {
		goto L1086
	} else {
		goto L1089
	}
L1089:
	;
	v4896 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4885)+2)))
	if v4896 <= int32(0) {
		goto L1086
	} else {
		goto L1090
	}
L1090:
	;
	v4899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4882)+90)))
	if v4899 != int32(118) {
		v4903 = v4862
		goto L1084
	} else {
		goto L1091
	}
L1091:
	;
	goto L1086
L1092:
	;
	v4907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4882)+90)))
	if v4907 == int32(118) {
		v4924 = v4873
		goto L1081
	} else {
		goto L1093
	}
L1093:
	;
	v4910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4885)+5)))
	v4916 = (v4876 + v4910 - int32(1)) & (int32(0) - v4910)
	if int32(_a_F_PostgresMain_111) < v4916 {
		v4924 = v4873
		goto L1081
	} else {
		goto L1094
	}
L1094:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v4885))) = uint16(v4916)
	v4922 = v4873 + int32(1)
	if v4922 != v4862 {
		v4873 = v4922
		v4874 = v4903
		v4876 = v4916 + v4904
		goto L1082
	} else {
		goto L1095
	}
L1095:
	;
	goto L1083
L1096:
	;
	v4945 = F_cstring_to_text(m, v3481+int32(_a_F_PostgresMain_101))
	mBase = m.M
	v4946 = m.ExcPending
	if v4946 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1097
	}
L1097:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[178]))) = base.I64_extend_i32_u(v4945)
	v4949 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3481)+440)))
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[179]))) = v4949
	v4953 = F_cstring_to_text(m, v3481+int32(496))
	mBase = m.M
	v4954 = m.ExcPending
	if v4954 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1098
	}
L1098:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[180]))) = base.I64_extend_i32_u(v4953)
	if v4826 != 0 {
		goto L1100
	} else {
		goto L1101
	}
L1099:
	;
	F_do_tup_output(m, v4941, v3481+int32(_a_F_PostgresMain_104), v3481+int32(464))
	mBase = m.M
	v4968 = m.ExcPending
	if v4968 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1104
	}
L1100:
	;
	v4957 = F_cstring_to_text(m, v4826)
	mBase = m.M
	v4958 = m.ExcPending
	if v4958 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1103
	}
L1101:
	;
	goto L1102
L1102:
	;
	v4961 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3481)+467)) = uint8(v4961)
	goto L1099
L1103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[181]))) = base.I64_extend_i32_u(v4957)
	goto L1099
L1104:
	;
	F_end_tup_output(m, v4941)
	mBase = m.M
	v4970 = m.ExcPending
	if v4970 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1105
	}
L1105:
	;
	v8300 = int32(_a_F_PostgresMain_112)
	goto L726
L1106:
	;
	F_TupleDescInitBuiltinEntry(m, v4980, int32(1), int32(_a_F_PostgresMain_113), int32(25))
	mBase = m.M
	v4986 = m.ExcPending
	if v4986 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1107
	}
L1107:
	;
	F_TupleDescInitBuiltinEntry(m, v4980, int32(2), int32(_a_F_PostgresMain_114), int32(25))
	mBase = m.M
	v4991 = m.ExcPending
	if v4991 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1108
	}
L1108:
	;
	F_TupleDescInitBuiltinEntry(m, v4980, int32(3), int32(_a_F_PostgresMain_115), int32(20))
	mBase = m.M
	v4996 = m.ExcPending
	if v4996 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1109
	}
L1109:
	;
	v4997 = int32(0)
	v5006 = *(*int32)(unsafe.Add(mBase, uint32(v4980)))
	if v4997 < v5006 {
		goto L1111
	} else {
		goto L1112
	}
L1110:
	;
	v5084 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3481)+462)) = uint8(v5084)
	v5086 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v3481)+460)) = uint16(v5086)
	v5089 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[44]))
	v5093 = F_LWLockAcquire(m, v5089+int32(_a_F_PostgresMain_116), v5084)
	mBase = m.M
	v5094 = m.ExcPending
	if v5094 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1129
	}
L1111:
	;
	v5010 = v4980 + int32(28)
	v5017 = v4997
	v5018 = v5006
	v5020 = v4997
	goto L1115
L1112:
	;
	v5074 = v4997
	v5081 = v5006
	goto L1113
L1113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4980)+20)) = v5081
	*(*int32)(unsafe.Add(mBase, uint32(v4980)+16)) = v5074
	goto L1110
L1114:
	;
	v5074 = v5068
	v5081 = v5047
	goto L1113
L1115:
	;
	v5026 = v5010 + v5006<<(uint(int32(3))%32) + v5017*int32(100)
	v5029 = v5010 + v5017<<(uint(int32(3))%32)
	if v5006 != v5018 {
		v5047 = v5018
		goto L1117
	} else {
		goto L1118
	}
L1116:
	;
	v5068 = v5006
	goto L1114
L1117:
	;
	v5048 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5029)+2)))
	if v5048 <= int32(0) {
		v5068 = v5017
		goto L1114
	} else {
		goto L1125
	}
L1118:
	;
	v5031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5029)+7)))
	if v5031 != int32(118) {
		goto L1119
	} else {
		goto L1120
	}
L1119:
	;
	v5047 = v5017
	goto L1117
L1120:
	;
	v5034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5029)+4)))
	if v5034 != int32(1) {
		goto L1119
	} else {
		goto L1121
	}
L1121:
	;
	v5037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5029)+6)))
	if v5037&int32(6) != 0 {
		goto L1119
	} else {
		goto L1122
	}
L1122:
	;
	v5040 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5029)+2)))
	if v5040 <= int32(0) {
		goto L1119
	} else {
		goto L1123
	}
L1123:
	;
	v5043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5026)+90)))
	if v5043 != int32(118) {
		v5047 = v5006
		goto L1117
	} else {
		goto L1124
	}
L1124:
	;
	goto L1119
L1125:
	;
	v5051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5026)+90)))
	if v5051 == int32(118) {
		v5068 = v5017
		goto L1114
	} else {
		goto L1126
	}
L1126:
	;
	v5054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5029)+5)))
	v5060 = (v5020 + v5054 - int32(1)) & (int32(0) - v5054)
	if int32(_a_F_PostgresMain_111) < v5060 {
		v5068 = v5017
		goto L1114
	} else {
		goto L1127
	}
L1127:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v5029))) = uint16(v5060)
	v5066 = v5017 + int32(1)
	if v5066 != v5006 {
		v5017 = v5066
		v5018 = v5047
		v5020 = v5060 + v5048
		goto L1115
	} else {
		goto L1128
	}
L1128:
	;
	goto L1116
L1129:
	;
	v5095 = *(*int32)(unsafe.Add(mBase, uint32(v4972)+4))
	v5097 = F_SearchNamedReplicationSlot(m, v5095, int32(0))
	mBase = m.M
	v5098 = m.ExcPending
	if v5098 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1132
	}
L1130:
	;
	v5204 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v5205 = m.ExcPending
	if v5205 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1161
	}
L1131:
	;
	v5108 = base.AtomicRmwXchg32(m, v5097, int32(0), int32(1))
	if v5108 != 0 {
		goto L1138
	} else {
		goto L1139
	}
L1132:
	;
	if v5097 != 0 {
		goto L1133
	} else {
		goto L1134
	}
L1133:
	;
	v5099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5097)+4)))
	if v5099 != 0 {
		goto L1131
	} else {
		goto L1136
	}
L1134:
	;
	goto L1135
L1135:
	;
	v5101 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[44]))
	F_LWLockRelease(m, v5101+int32(_a_F_PostgresMain_116))
	mBase = m.M
	v5105 = m.ExcPending
	if v5105 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1137
	}
L1136:
	;
	goto L1135
L1137:
	;
	goto L1130
L1138:
	;
	F_s_lock(m, v5097, int32(_a_F_PostgresMain_11))
	mBase = m.M
	v5111 = m.ExcPending
	if v5111 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1141
	}
L1139:
	;
	goto L1140
L1140:
	;
	base.MemoryCopy(m, v3481+int32(_a_F_PostgresMain_101), v5097, int32(88))
	v5116 = *(*int32)(unsafe.Add(mBase, uint32(v5097)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+456)) = v5116
	v5118 = *(*int32)(unsafe.Add(mBase, uint32(v5097)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+448)) = v5118
	v5120 = *(*int64)(unsafe.Add(mBase, uint32(v5097)+92))
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+440)) = v5120
	v5122 = *(*int64)(unsafe.Add(mBase, uint32(v5097)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+432)) = v5122
	base.MemoryCopy(m, v3481+int32(496), v5097+int32(112), int32(184))
	v5130 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v5097))), uint32(v5130))
	v5134 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[44]))
	F_LWLockRelease(m, v5134+int32(_a_F_PostgresMain_116))
	mBase = m.M
	v5138 = m.ExcPending
	if v5138 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1142
	}
L1141:
	;
	goto L1140
L1142:
	;
	if v5116 != 0 {
		goto L750
	} else {
		goto L1143
	}
L1143:
	;
	v5140 = F_cstring_to_text(m, int32(_a_F_PostgresMain_74))
	mBase = m.M
	v5141 = m.ExcPending
	if v5141 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1144
	}
L1144:
	;
	v5142 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3481)+460)) = uint8(v5142)
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+464)) = base.I64_extend_i32_u(v5140)
	if v5122 == int64(0) {
		goto L1130
	} else {
		goto L1145
	}
L1145:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v3481)+52)) = uint32(v5122)
	v5150 = int64(base.Ui64(v5122) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3481)+48)) = uint32(v5150)
	v5153 = v3481 + int32(_a_F_PostgresMain_104)
	v5158 = F_pg_snprintf(m, v5153, int32(64), int32(_a_F_PostgresMain_103), v3481+int32(48))
	mBase = m.M
	v5159 = m.ExcPending
	if v5159 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1146
	}
L1146:
	;
	v5160 = F_cstring_to_text(m, v5153)
	mBase = m.M
	v5161 = m.ExcPending
	if v5161 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1147
	}
L1147:
	;
	v5162 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3481)+461)) = uint8(v5162)
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+472)) = base.I64_extend_i32_u(v5160)
	v5166 = *(*int64)(unsafe.Add(mBase, uint32(v3481)+432))
	if v5166 == int64(0) {
		goto L1130
	} else {
		goto L1148
	}
L1148:
	;
	v5171 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])))
	if v5171 == int32(1) {
		goto L1151
	} else {
		goto L1152
	}
L1149:
	;
	v5190 = F_readTimeLineHistory(m, v5189)
	mBase = m.M
	v5191 = m.ExcPending
	if v5191 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1159
	}
L1150:
	;
	if v5181 != 0 {
		goto L1154
	} else {
		goto L1155
	}
L1151:
	;
	v5176 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v5177 = *(*int32)(unsafe.Add(mBase, uint32(v5176)+308))
	v5179 = base.B2i32(v5177 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])) = uint8(v5179)
	v5181 = v5179
	goto L1153
L1152:
	;
	v5181 = int32(0)
	goto L1153
L1153:
	;
	goto L1150
L1154:
	;
	v5182 = F_GetXLogReplayRecPtr(m, v5153)
	mBase = m.M
	v5183 = m.ExcPending
	if v5183 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1157
	}
L1155:
	;
	goto L1156
L1156:
	;
	v5186 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v5187 = *(*int32)(unsafe.Add(mBase, uint32(v5186)+300))
	goto L1158
L1157:
	;
	v5184 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[178])))
	v5189 = v5184
	goto L1149
L1158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[178]))) = v5187
	v5189 = v5187
	goto L1149
L1159:
	;
	v5192 = *(*int64)(unsafe.Add(mBase, uint32(v3481)+432))
	v5193 = F_tliOfPointInHistory(m, v5192, v5190)
	mBase = m.M
	v5194 = m.ExcPending
	if v5194 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1160
	}
L1160:
	;
	v5195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3481)+462)) = uint8(v5195)
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+480)) = base.I64_extend_i32_u(v5193)
	goto L1130
L1161:
	;
	v5207 = F_begin_tup_output_tupdesc(m, v5204, v4980, int32(_a_F_PostgresMain_110))
	mBase = m.M
	v5208 = m.ExcPending
	if v5208 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1162
	}
L1162:
	;
	F_do_tup_output(m, v5207, v3481+int32(464), v3481+int32(460))
	mBase = m.M
	v5214 = m.ExcPending
	if v5214 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1163
	}
L1163:
	;
	F_end_tup_output(m, v5207)
	mBase = m.M
	v5216 = m.ExcPending
	if v5216 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1164
	}
L1164:
	;
	v8300 = int32(_a_F_PostgresMain_117)
	goto L726
L1165:
	;
	v5580 = *(*int32)(unsafe.Add(mBase, uint32(v5218)+8))
	if v5580 == int32(0) {
		goto L1257
	} else {
		goto L1258
	}
L1166:
	;
	v5224 = *(*int32)(unsafe.Add(mBase, uint32(v5221)+4))
	if v5224 <= int32(0) {
		v5548 = v3469
		v5559 = v3469
		v5564 = v2541
		v5572 = v3469
		goto L1165
	} else {
		goto L1167
	}
L1167:
	;
	v5230 = int32(0)
	v5235 = v3469
	v5245 = v3469
	v5246 = v3469
	v5251 = v2541
	v5252 = v3469
	v5256 = v3469
	v5257 = v3469
	v5259 = v3469
	goto L1168
L1168:
	;
	v5267 = *(*int32)(unsafe.Add(mBase, uint32(v5221)+12))
	v5271 = *(*int32)(unsafe.Add(mBase, uint32(v5267+v5230<<(uint(int32(2))%32))))
	v5272 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+8))
	v5273 = int32(_a_F_PostgresMain_89)
	v5276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5272))))
	v5279 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[182])))
	if base.B2i32(v5276 == int32(0))|base.B2i32(v5276 != v5279) != 0 {
		v5297 = v5276
		v5298 = v5279
		goto L1172
	} else {
		goto L1173
	}
L1169:
	;
	v5548 = v5529
	v5559 = v5531
	v5564 = v5532
	v5572 = v5536
	goto L1165
L1170:
	;
	v5538 = v5230 + int32(1)
	v5539 = *(*int32)(unsafe.Add(mBase, uint32(v5221)+4))
	if v5538 < v5539 {
		v5230 = v5538
		v5235 = v5529
		v5245 = v5530
		v5246 = v5531
		v5251 = v5532
		v5252 = v5533
		v5256 = v5534
		v5257 = v5535
		v5259 = v5536
		goto L1168
	} else {
		goto L1255
	}
L1171:
	;
	if v5297-v5298 == int32(0) {
		goto L1178
	} else {
		goto L1179
	}
L1172:
	;
	goto L1171
L1173:
	;
	v5282 = v5272
	v5283 = v5273
	goto L1174
L1174:
	;
	v5286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5283)+1)))
	v5287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5282)+1)))
	if v5287 == int32(0) {
		v5297 = v5287
		v5298 = v5286
		goto L1172
	} else {
		goto L1176
	}
L1175:
	;
	v5297 = v5287
	v5298 = v5286
	goto L1172
L1176:
	;
	v5290 = int32(1)
	if v5287 == v5286 {
		v5282 = v5282 + v5290
		v5283 = v5283 + v5290
		goto L1174
	} else {
		goto L1177
	}
L1177:
	;
	goto L1175
L1178:
	;
	if v5252&int32(1) != 0 {
		goto L749
	} else {
		goto L1181
	}
L1179:
	;
	goto L1180
L1180:
	;
	v5423 = int32(_a_F_PostgresMain_71)
	v5426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5272))))
	v5429 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[183])))
	if base.B2i32(v5426 == int32(0))|base.B2i32(v5426 != v5429) != 0 {
		v5447 = v5426
		v5448 = v5429
		goto L1219
	} else {
		goto L1220
	}
L1181:
	;
	v5304 = *(*int32)(unsafe.Add(mBase, uint32(v5218)+8))
	if v5304 != int32(1) {
		goto L749
	} else {
		goto L1182
	}
L1182:
	;
	v5307 = F_defGetString(m, v5271)
	mBase = m.M
	v5308 = m.ExcPending
	if v5308 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1183
	}
L1183:
	;
	v5309 = int32(_a_F_PostgresMain_87)
	v5312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5307))))
	v5315 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[184])))
	if base.B2i32(v5312 == int32(0))|base.B2i32(v5312 != v5315) != 0 {
		v5333 = v5312
		v5334 = v5315
		goto L1185
	} else {
		goto L1186
	}
L1184:
	;
	if v5333-v5334 == int32(0) {
		goto L1191
	} else {
		goto L1192
	}
L1185:
	;
	goto L1184
L1186:
	;
	v5318 = v5307
	v5319 = v5309
	goto L1187
L1187:
	;
	v5322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5319)+1)))
	v5323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5318)+1)))
	if v5323 == int32(0) {
		v5333 = v5323
		v5334 = v5322
		goto L1185
	} else {
		goto L1189
	}
L1188:
	;
	v5333 = v5323
	v5334 = v5322
	goto L1185
L1189:
	;
	v5326 = int32(1)
	if v5323 == v5322 {
		v5318 = v5318 + v5326
		v5319 = v5319 + v5326
		goto L1187
	} else {
		goto L1190
	}
L1190:
	;
	goto L1188
L1191:
	;
	v5529 = v5235
	v5530 = v5245
	v5531 = v5246
	v5532 = int32(0)
	v5533 = int32(1)
	v5534 = v5256
	v5535 = v5257
	v5536 = v5259
	goto L1170
L1192:
	;
	goto L1193
L1193:
	;
	v5340 = int32(1)
	v5341 = int32(_a_F_PostgresMain_86)
	v5344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5307))))
	v5347 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[185])))
	if base.B2i32(v5344 == int32(0))|base.B2i32(v5344 != v5347) != 0 {
		v5365 = v5344
		v5366 = v5347
		goto L1195
	} else {
		goto L1196
	}
L1194:
	;
	if v5365-v5366 == int32(0) {
		goto L1201
	} else {
		goto L1202
	}
L1195:
	;
	goto L1194
L1196:
	;
	v5350 = v5307
	v5351 = v5341
	goto L1197
L1197:
	;
	v5354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5351)+1)))
	v5355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5350)+1)))
	if v5355 == int32(0) {
		v5365 = v5355
		v5366 = v5354
		goto L1195
	} else {
		goto L1199
	}
L1198:
	;
	v5365 = v5355
	v5366 = v5354
	goto L1195
L1199:
	;
	v5358 = int32(1)
	if v5355 == v5354 {
		v5350 = v5350 + v5358
		v5351 = v5351 + v5358
		goto L1197
	} else {
		goto L1200
	}
L1200:
	;
	goto L1198
L1201:
	;
	v5529 = v5235
	v5530 = v5245
	v5531 = v5246
	v5532 = int32(1)
	v5533 = v5340
	v5534 = v5256
	v5535 = v5257
	v5536 = v5259
	goto L1170
L1202:
	;
	goto L1203
L1203:
	;
	v5371 = int32(_a_F_PostgresMain_85)
	v5374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5307))))
	v5377 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[186])))
	if base.B2i32(v5374 == int32(0))|base.B2i32(v5374 != v5377) != 0 {
		v5395 = v5374
		v5396 = v5377
		goto L1205
	} else {
		goto L1206
	}
L1204:
	;
	if v5395-v5396 == int32(0) {
		goto L1211
	} else {
		goto L1212
	}
L1205:
	;
	goto L1204
L1206:
	;
	v5380 = v5307
	v5381 = v5371
	goto L1207
L1207:
	;
	v5384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5381)+1)))
	v5385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5380)+1)))
	if v5385 == int32(0) {
		v5395 = v5385
		v5396 = v5384
		goto L1205
	} else {
		goto L1209
	}
L1208:
	;
	v5395 = v5385
	v5396 = v5384
	goto L1205
L1209:
	;
	v5388 = int32(1)
	if v5385 == v5384 {
		v5380 = v5380 + v5388
		v5381 = v5381 + v5388
		goto L1207
	} else {
		goto L1210
	}
L1210:
	;
	goto L1208
L1211:
	;
	v5529 = v5235
	v5530 = v5245
	v5531 = v5246
	v5532 = int32(2)
	v5533 = v5340
	v5534 = v5256
	v5535 = v5257
	v5536 = v5259
	goto L1170
L1212:
	;
	goto L1213
L1213:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5404 = m.ExcPending
	if v5404 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1214
	}
L1214:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v5407 = m.ExcPending
	if v5407 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1215
	}
L1215:
	;
	v5408 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+200)) = v5307
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+196)) = v5408
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+192)) = int32(_a_F_PostgresMain_118)
	F_errmsg(m, int32(_a_F_PostgresMain_119), v3481+int32(192))
	mBase = m.M
	v5417 = m.ExcPending
	if v5417 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1216
	}
L1216:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1210), int32(_a_F_PostgresMain_120))
	mBase = m.M
	v5422 = m.ExcPending
	if v5422 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1217
	}
L1217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1218:
	;
	if v5447-v5448 == int32(0) {
		goto L1225
	} else {
		goto L1226
	}
L1219:
	;
	goto L1218
L1220:
	;
	v5432 = v5272
	v5433 = v5423
	goto L1221
L1221:
	;
	v5436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5433)+1)))
	v5437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5432)+1)))
	if v5437 == int32(0) {
		v5447 = v5437
		v5448 = v5436
		goto L1219
	} else {
		goto L1223
	}
L1222:
	;
	v5447 = v5437
	v5448 = v5436
	goto L1219
L1223:
	;
	v5440 = int32(1)
	if v5437 == v5436 {
		v5432 = v5432 + v5440
		v5433 = v5433 + v5440
		goto L1221
	} else {
		goto L1224
	}
L1224:
	;
	goto L1222
L1225:
	;
	if v5256&int32(1) != 0 {
		goto L748
	} else {
		goto L1228
	}
L1226:
	;
	goto L1227
L1227:
	;
	v5458 = int32(_a_F_PostgresMain_69)
	v5461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5272))))
	v5464 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[187])))
	if base.B2i32(v5461 == int32(0))|base.B2i32(v5461 != v5464) != 0 {
		v5482 = v5461
		v5483 = v5464
		goto L1232
	} else {
		goto L1233
	}
L1228:
	;
	v5454 = *(*int32)(unsafe.Add(mBase, uint32(v5218)+8))
	if v5454 != 0 {
		goto L748
	} else {
		goto L1229
	}
L1229:
	;
	v5456 = F_defGetBoolean(m, v5271)
	mBase = m.M
	v5457 = m.ExcPending
	if v5457 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1230
	}
L1230:
	;
	v5529 = v5235
	v5530 = v5245
	v5531 = v5246
	v5532 = v5251
	v5533 = v5252
	v5534 = int32(1)
	v5535 = v5257
	v5536 = v5456
	goto L1170
L1231:
	;
	if v5482-v5483 == int32(0) {
		goto L1238
	} else {
		goto L1239
	}
L1232:
	;
	goto L1231
L1233:
	;
	v5467 = v5272
	v5468 = v5458
	goto L1234
L1234:
	;
	v5471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5468)+1)))
	v5472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5467)+1)))
	if v5472 == int32(0) {
		v5482 = v5472
		v5483 = v5471
		goto L1232
	} else {
		goto L1236
	}
L1235:
	;
	v5482 = v5472
	v5483 = v5471
	goto L1232
L1236:
	;
	v5475 = int32(1)
	if v5472 == v5471 {
		v5467 = v5467 + v5475
		v5468 = v5468 + v5475
		goto L1234
	} else {
		goto L1237
	}
L1237:
	;
	goto L1235
L1238:
	;
	if v5245 != 0 {
		goto L747
	} else {
		goto L1241
	}
L1239:
	;
	goto L1240
L1240:
	;
	v5493 = int32(_a_F_PostgresMain_121)
	v5496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5272))))
	v5499 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[188])))
	if base.B2i32(v5496 == int32(0))|base.B2i32(v5496 != v5499) != 0 {
		v5517 = v5496
		v5518 = v5499
		goto L1245
	} else {
		goto L1246
	}
L1241:
	;
	v5487 = *(*int32)(unsafe.Add(mBase, uint32(v5218)+8))
	if v5487 != int32(1) {
		goto L747
	} else {
		goto L1242
	}
L1242:
	;
	v5491 = F_defGetBoolean(m, v5271)
	mBase = m.M
	v5492 = m.ExcPending
	if v5492 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1243
	}
L1243:
	;
	v5529 = v5491
	v5530 = int32(1)
	v5531 = v5246
	v5532 = v5251
	v5533 = v5252
	v5534 = v5256
	v5535 = v5257
	v5536 = v5259
	goto L1170
L1244:
	;
	if v5517-v5518 != 0 {
		goto L745
	} else {
		goto L1251
	}
L1245:
	;
	goto L1244
L1246:
	;
	v5502 = v5272
	v5503 = v5493
	goto L1247
L1247:
	;
	v5506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5503)+1)))
	v5507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5502)+1)))
	if v5507 == int32(0) {
		v5517 = v5507
		v5518 = v5506
		goto L1245
	} else {
		goto L1249
	}
L1248:
	;
	v5517 = v5507
	v5518 = v5506
	goto L1245
L1249:
	;
	v5510 = int32(1)
	if v5507 == v5506 {
		v5502 = v5502 + v5510
		v5503 = v5503 + v5510
		goto L1247
	} else {
		goto L1250
	}
L1250:
	;
	goto L1248
L1251:
	;
	if v5257&int32(1) != 0 {
		goto L746
	} else {
		goto L1252
	}
L1252:
	;
	v5522 = *(*int32)(unsafe.Add(mBase, uint32(v5218)+8))
	if v5522 != int32(1) {
		goto L746
	} else {
		goto L1253
	}
L1253:
	;
	v5526 = F_defGetBoolean(m, v5271)
	mBase = m.M
	v5527 = m.ExcPending
	if v5527 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1254
	}
L1254:
	;
	v5529 = v5235
	v5530 = v5245
	v5531 = v5526
	v5532 = v5251
	v5533 = v5252
	v5534 = v5256
	v5535 = int32(1)
	v5536 = v5259
	goto L1170
L1255:
	;
	goto L1169
L1256:
	;
	v5801 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v5802 = *(*int64)(unsafe.Add(mBase, uint32(v5801)+120))
	*(*uint32)(unsafe.Add(mBase, uint32(v3481)+84)) = uint32(v5802)
	v5805 = int64(base.Ui64(v5802) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3481)+80)) = uint32(v5805)
	v5808 = v3481 + int32(496)
	v5813 = F_pg_snprintf(m, v5808, int32(64), int32(_a_F_PostgresMain_103), v3481+int32(80))
	mBase = m.M
	v5814 = m.ExcPending
	if v5814 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1319
	}
L1257:
	;
	v5583 = int32(0)
	v5584 = *(*int32)(unsafe.Add(mBase, uint32(v5218)+4))
	v5586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5218)+16)))
	F_ReplicationSlotCreate(m, v5584, v5583, v5586<<(uint(int32(1))%32)&int32(2), v5583, v5583, v5583, v5583)
	mBase = m.M
	v5596 = m.ExcPending
	if v5596 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1260
	}
L1258:
	;
	goto L1259
L1259:
	;
	v5608 = int32(0)
	F_CheckLogicalDecodingRequirements(m, v5608)
	mBase = m.M
	v5611 = m.ExcPending
	if v5611 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1266
	}
L1260:
	;
	if v5572&int32(1) == int32(0) {
		v5799 = v5583
		goto L1256
	} else {
		goto L1261
	}
L1261:
	;
	F_ReplicationSlotReserveWal(m)
	mBase = m.M
	v5602 = m.ExcPending
	if v5602 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1262
	}
L1262:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v5604 = m.ExcPending
	if v5604 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1263
	}
L1263:
	;
	v5605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5218)+16)))
	if v5605 != 0 {
		v5799 = v5583
		goto L1256
	} else {
		goto L1264
	}
L1264:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v5607 = m.ExcPending
	if v5607 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1265
	}
L1265:
	;
	v5799 = v5583
	goto L1256
L1266:
	;
	v5612 = *(*int32)(unsafe.Add(mBase, uint32(v5218)+4))
	v5613 = int32(1)
	v5616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5218)+16)))
	if v5616 != 0 {
		goto L1267
	} else {
		goto L1268
	}
L1267:
	;
	v5617 = int32(2)
	goto L1269
L1268:
	;
	v5617 = v5613
	goto L1269
L1269:
	;
	v5618 = int32(1)
	v5620 = int32(0)
	F_ReplicationSlotCreate(m, v5612, v5613, v5617, v5548&v5618, v5620, v5559&v5618, v5620)
	mBase = m.M
	v5625 = m.ExcPending
	if v5625 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1270
	}
L1270:
	;
	switch v5564 {
	case 0:
		goto L1273
	default:
		v5676 = int32(0)
		goto L1271
	case 2:
		goto L1272
	}
L1271:
	;
	F_EnsureLogicalDecodingEnabled(m)
	mBase = m.M
	v5678 = m.ExcPending
	if v5678 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1286
	}
L1272:
	;
	v5652 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v5653 = *(*int32)(unsafe.Add(mBase, uint32(v5652)+24))
	goto L1279
L1273:
	;
	v5627 = int32(1)
	v5629 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v5630 = *(*int32)(unsafe.Add(mBase, uint32(v5629)+24))
	goto L1274
L1274:
	;
	if base.B2i32(base.Ui32(v5627) < base.Ui32(v5630)) == int32(0) {
		v5676 = v5627
		goto L1271
	} else {
		goto L1275
	}
L1275:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5638 = m.ExcPending
	if v5638 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1276
	}
L1276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+96)) = int32(_a_F_PostgresMain_122)
	F_errmsg(m, int32(_a_F_PostgresMain_123), v3481+int32(96))
	mBase = m.M
	v5645 = m.ExcPending
	if v5645 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1277
	}
L1277:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1316), int32(_a_F_PostgresMain_124))
	mBase = m.M
	v5650 = m.ExcPending
	if v5650 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
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
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(v5653)) == int32(0) {
		goto L744
	} else {
		goto L1280
	}
L1280:
	;
	v5659 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[189]))
	if v5659 != int32(2) {
		goto L743
	} else {
		goto L1281
	}
L1281:
	;
	v5663 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[190])))
	if v5663 == int32(0) {
		goto L742
	} else {
		goto L1282
	}
L1282:
	;
	v5666 = int32(1)
	v5668 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[191])))
	if v5668 == v5666 {
		goto L741
	} else {
		goto L1283
	}
L1283:
	;
	v5672 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v5673 = *(*int32)(unsafe.Add(mBase, uint32(v5672)+28))
	goto L1284
L1284:
	;
	if int32(1) < v5673 {
		goto L740
	} else {
		goto L1285
	}
L1285:
	;
	v5676 = v5666
	goto L1271
L1286:
	;
	v5679 = *(*int32)(unsafe.Add(mBase, uint32(v5218)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[179]))) = int32(414)
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[192]))) = int32(1107)
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[178]))) = int32(1108)
	v5693 = F_CreateInitDecodingContext(m, v5679, v5676, int32(0), int64(0), v3481+int32(_a_F_PostgresMain_104), int32(1109), int32(1110), int32(1111))
	mBase = m.M
	v5694 = m.ExcPending
	if v5694 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1287
	}
L1287:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[193])) = int64(0)
	F_DecodingContextFindStartpoint(m, v5693)
	mBase = m.M
	v5699 = m.ExcPending
	if v5699 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1288
	}
L1288:
	;
	switch v5564 {
	case 0:
		goto L1291
	default:
		v5789 = v5608
		goto L1289
	case 2:
		goto L1290
	}
L1289:
	;
	F_FreeDecodingContext(m, v5693)
	mBase = m.M
	v5791 = m.ExcPending
	if v5791 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1316
	}
L1290:
	;
	v5779 = *(*int32)(unsafe.Add(mBase, uint32(v5693)+16))
	v5780 = F_SnapBuildInitialSnapshot(m, v5779)
	mBase = m.M
	v5781 = m.ExcPending
	if v5781 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1314
	}
L1291:
	;
	v5700 = *(*int32)(unsafe.Add(mBase, uint32(v5693)+16))
	v5701 = m.G0
	v5703 = v5701 - int32(16)
	m.G0 = v5703
	v5706 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v5707 = *(*int32)(unsafe.Add(mBase, uint32(v5706)+24))
	goto L1294
L1292:
	;
	v5789 = v5731
	goto L1289
L1293:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5769 = m.ExcPending
	if v5769 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1311
	}
L1294:
	;
	if base.B2i32(v5707 != int32(0)) == int32(0) {
		goto L1295
	} else {
		goto L1296
	}
L1295:
	;
	v5713 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[162]))
	if v5713 != 0 {
		goto L1293
	} else {
		goto L1298
	}
L1296:
	;
	goto L1297
L1297:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5756 = m.ExcPending
	if v5756 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1308
	}
L1298:
	;
	v5715 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[161])) = uint8(v5715)
	v5719 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[163]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[162])) = v5719
	F_StartTransactionCommand(m)
	mBase = m.M
	v5722 = m.ExcPending
	if v5722 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1299
	}
L1299:
	;
	v5724 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[190])) = uint8(v5724)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[189])) = int32(2)
	v5729 = F_SnapBuildInitialSnapshot(m, v5700)
	mBase = m.M
	v5730 = m.ExcPending
	if v5730 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1300
	}
L1300:
	;
	v5731 = F_ExportSnapshot(m, v5729)
	mBase = m.M
	v5732 = m.ExcPending
	if v5732 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1301
	}
L1301:
	;
	v5735 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5736 = m.ExcPending
	if v5736 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1302
	}
L1302:
	;
	if v5735 != 0 {
		goto L1303
	} else {
		goto L1304
	}
L1303:
	;
	v5737 = *(*int32)(unsafe.Add(mBase, uint32(v5729)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v5703)+4)) = v5737
	*(*int32)(unsafe.Add(mBase, uint32(v5703))) = v5731
	F_errmsg_plural(m, int32(_a_F_PostgresMain_125), int32(_a_F_PostgresMain_126), v5737, v5703)
	mBase = m.M
	v5743 = m.ExcPending
	if v5743 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1306
	}
L1304:
	;
	goto L1305
L1305:
	;
	m.G0 = v5703 + int32(16)
	goto L1292
L1306:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_50), int32(574), int32(_a_F_PostgresMain_127))
	mBase = m.M
	v5748 = m.ExcPending
	if v5748 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1307
	}
L1307:
	;
	goto L1305
L1308:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMain_128), int32(0))
	mBase = m.M
	v5760 = m.ExcPending
	if v5760 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1309
	}
L1309:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_50), int32(548), int32(_a_F_PostgresMain_127))
	mBase = m.M
	v5765 = m.ExcPending
	if v5765 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1310
	}
L1310:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1311:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMain_129), int32(0))
	mBase = m.M
	v5773 = m.ExcPending
	if v5773 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1312
	}
L1312:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_50), int32(551), int32(_a_F_PostgresMain_127))
	mBase = m.M
	v5778 = m.ExcPending
	if v5778 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1313
	}
L1313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1314:
	;
	v5783 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[46]))
	F_RestoreTransactionSnapshot(m, v5780, v5783)
	mBase = m.M
	v5785 = m.ExcPending
	if v5785 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1315
	}
L1315:
	;
	v5789 = v5608
	goto L1289
L1316:
	;
	v5792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5218)+16)))
	if v5792 != 0 {
		v5799 = v5789
		goto L1256
	} else {
		goto L1317
	}
L1317:
	;
	F_ReplicationSlotPersist(m)
	mBase = m.M
	v5794 = m.ExcPending
	if v5794 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1318
	}
L1318:
	;
	v5799 = v5789
	goto L1256
L1319:
	;
	v5816 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v5817 = m.ExcPending
	if v5817 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1320
	}
L1320:
	;
	v5819 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v5820 = m.ExcPending
	if v5820 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1321
	}
L1321:
	;
	F_TupleDescInitBuiltinEntry(m, v5819, int32(1), int32(_a_F_PostgresMain_130), int32(25))
	mBase = m.M
	v5825 = m.ExcPending
	if v5825 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1322
	}
L1322:
	;
	F_TupleDescInitBuiltinEntry(m, v5819, int32(2), int32(_a_F_PostgresMain_131), int32(25))
	mBase = m.M
	v5830 = m.ExcPending
	if v5830 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1323
	}
L1323:
	;
	F_TupleDescInitBuiltinEntry(m, v5819, int32(3), int32(_a_F_PostgresMain_132), int32(25))
	mBase = m.M
	v5835 = m.ExcPending
	if v5835 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1324
	}
L1324:
	;
	F_TupleDescInitBuiltinEntry(m, v5819, int32(4), int32(_a_F_PostgresMain_133), int32(25))
	mBase = m.M
	v5840 = m.ExcPending
	if v5840 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1325
	}
L1325:
	;
	v5841 = int32(0)
	v5850 = *(*int32)(unsafe.Add(mBase, uint32(v5819)))
	if v5841 < v5850 {
		goto L1327
	} else {
		goto L1328
	}
L1326:
	;
	v5929 = F_begin_tup_output_tupdesc(m, v5816, v5819, int32(_a_F_PostgresMain_110))
	mBase = m.M
	v5930 = m.ExcPending
	if v5930 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1345
	}
L1327:
	;
	v5854 = v5819 + int32(28)
	v5861 = v5841
	v5862 = v5850
	v5864 = v5841
	goto L1331
L1328:
	;
	v5918 = v5841
	v5925 = v5850
	goto L1329
L1329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5819)+20)) = v5925
	*(*int32)(unsafe.Add(mBase, uint32(v5819)+16)) = v5918
	goto L1326
L1330:
	;
	v5918 = v5912
	v5925 = v5891
	goto L1329
L1331:
	;
	v5870 = v5854 + v5850<<(uint(int32(3))%32) + v5861*int32(100)
	v5873 = v5854 + v5861<<(uint(int32(3))%32)
	if v5850 != v5862 {
		v5891 = v5862
		goto L1333
	} else {
		goto L1334
	}
L1332:
	;
	v5912 = v5850
	goto L1330
L1333:
	;
	v5892 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5873)+2)))
	if v5892 <= int32(0) {
		v5912 = v5861
		goto L1330
	} else {
		goto L1341
	}
L1334:
	;
	v5875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5873)+7)))
	if v5875 != int32(118) {
		goto L1335
	} else {
		goto L1336
	}
L1335:
	;
	v5891 = v5861
	goto L1333
L1336:
	;
	v5878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5873)+4)))
	if v5878 != int32(1) {
		goto L1335
	} else {
		goto L1337
	}
L1337:
	;
	v5881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5873)+6)))
	if v5881&int32(6) != 0 {
		goto L1335
	} else {
		goto L1338
	}
L1338:
	;
	v5884 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5873)+2)))
	if v5884 <= int32(0) {
		goto L1335
	} else {
		goto L1339
	}
L1339:
	;
	v5887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5870)+90)))
	if v5887 != int32(118) {
		v5891 = v5850
		goto L1333
	} else {
		goto L1340
	}
L1340:
	;
	goto L1335
L1341:
	;
	v5895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5870)+90)))
	if v5895 == int32(118) {
		v5912 = v5861
		goto L1330
	} else {
		goto L1342
	}
L1342:
	;
	v5898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5873)+5)))
	v5904 = (v5864 + v5898 - int32(1)) & (int32(0) - v5898)
	if int32(_a_F_PostgresMain_111) < v5904 {
		v5912 = v5861
		goto L1330
	} else {
		goto L1343
	}
L1343:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v5873))) = uint16(v5904)
	v5910 = v5861 + int32(1)
	if v5910 != v5850 {
		v5861 = v5910
		v5862 = v5891
		v5864 = v5904 + v5892
		goto L1331
	} else {
		goto L1344
	}
L1344:
	;
	goto L1332
L1345:
	;
	v5932 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v5935 = F_cstring_to_text(m, v5932+int32(24))
	mBase = m.M
	v5936 = m.ExcPending
	if v5936 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1346
	}
L1346:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[194]))) = base.I64_extend_i32_u(v5935)
	v5939 = F_cstring_to_text(m, v5808)
	mBase = m.M
	v5940 = m.ExcPending
	if v5940 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1347
	}
L1347:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[195]))) = base.I64_extend_i32_u(v5939)
	if v5799 != 0 {
		goto L1349
	} else {
		goto L1350
	}
L1348:
	;
	v5949 = *(*int32)(unsafe.Add(mBase, uint32(v5218)+12))
	if v5949 != 0 {
		goto L1354
	} else {
		goto L1355
	}
L1349:
	;
	v5943 = F_cstring_to_text(m, v5799)
	mBase = m.M
	v5944 = m.ExcPending
	if v5944 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1352
	}
L1350:
	;
	goto L1351
L1351:
	;
	v5947 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3481)+466)) = uint8(v5947)
	goto L1348
L1352:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[196]))) = base.I64_extend_i32_u(v5943)
	goto L1348
L1353:
	;
	F_do_tup_output(m, v5929, v3481+int32(_a_F_PostgresMain_101), v3481+int32(464))
	mBase = m.M
	v5961 = m.ExcPending
	if v5961 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1358
	}
L1354:
	;
	v5950 = F_cstring_to_text(m, v5949)
	mBase = m.M
	v5951 = m.ExcPending
	if v5951 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1357
	}
L1355:
	;
	goto L1356
L1356:
	;
	v5954 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3481)+467)) = uint8(v5954)
	goto L1353
L1357:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[197]))) = base.I64_extend_i32_u(v5950)
	goto L1353
L1358:
	;
	F_end_tup_output(m, v5929)
	mBase = m.M
	v5963 = m.ExcPending
	if v5963 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1359
	}
L1359:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v5965 = m.ExcPending
	if v5965 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1360
	}
L1360:
	;
	v8300 = int32(_a_F_PostgresMain_118)
	goto L726
L1361:
	;
	v8300 = int32(_a_F_PostgresMain_134)
	goto L726
L1362:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3481)+496)) = uint8(v6147)
	*(*uint8)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[194]))) = uint8(v6148)
	v6161 = *(*int32)(unsafe.Add(mBase, uint32(v5979)+4))
	v6162 = int32(0)
	v6163 = m.G0
	v6165 = v6163 - int32(1072)
	m.G0 = v6165
	F_ReplicationSlotAcquire(m, v6161, v6162, int32(1))
	mBase = m.M
	v6170 = m.ExcPending
	if v6170 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1397
	}
L1363:
	;
	v5983 = int32(0)
	v5984 = *(*int32)(unsafe.Add(mBase, uint32(v5980)+4))
	if v5984 <= v5983 {
		v6143 = v3469
		v6147 = v3469
		v6148 = v3469
		v6158 = v5983
		goto L1362
	} else {
		goto L1364
	}
L1364:
	;
	v5990 = int32(0)
	v6011 = v2541
	v6012 = v3469
	v6016 = v3469
	v6017 = v3469
	goto L1365
L1365:
	;
	v6027 = *(*int32)(unsafe.Add(mBase, uint32(v5980)+12))
	v6031 = *(*int32)(unsafe.Add(mBase, uint32(v6027+v5990<<(uint(int32(2))%32))))
	v6032 = *(*int32)(unsafe.Add(mBase, uint32(v6031)+8))
	v6033 = int32(_a_F_PostgresMain_121)
	v6036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6032))))
	v6039 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[188])))
	if base.B2i32(v6036 == int32(0))|base.B2i32(v6036 != v6039) != 0 {
		v6057 = v6036
		v6058 = v6039
		goto L1369
	} else {
		goto L1370
	}
L1366:
	;
	if v6100&int32(1) != 0 {
		goto L1391
	} else {
		goto L1392
	}
L1367:
	;
	v6104 = v5990 + int32(1)
	v6105 = *(*int32)(unsafe.Add(mBase, uint32(v5980)+4))
	if v6104 < v6105 {
		v5990 = v6104
		v6011 = v6099
		v6012 = v6100
		v6016 = v6101
		v6017 = v6102
		goto L1365
	} else {
		goto L1390
	}
L1368:
	;
	if v6057-v6058 == int32(0) {
		goto L1375
	} else {
		goto L1376
	}
L1369:
	;
	goto L1368
L1370:
	;
	v6042 = v6032
	v6043 = v6033
	goto L1371
L1371:
	;
	v6046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6043)+1)))
	v6047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6042)+1)))
	if v6047 == int32(0) {
		v6057 = v6047
		v6058 = v6046
		goto L1369
	} else {
		goto L1373
	}
L1372:
	;
	v6057 = v6047
	v6058 = v6046
	goto L1369
L1373:
	;
	v6050 = int32(1)
	if v6047 == v6046 {
		v6042 = v6042 + v6050
		v6043 = v6043 + v6050
		goto L1371
	} else {
		goto L1374
	}
L1374:
	;
	goto L1372
L1375:
	;
	if v6011&int32(1) != 0 {
		goto L739
	} else {
		goto L1378
	}
L1376:
	;
	goto L1377
L1377:
	;
	v6067 = int32(_a_F_PostgresMain_69)
	v6070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6032))))
	v6073 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[187])))
	if base.B2i32(v6070 == int32(0))|base.B2i32(v6070 != v6073) != 0 {
		v6091 = v6070
		v6092 = v6073
		goto L1381
	} else {
		goto L1382
	}
L1378:
	;
	v6065 = F_defGetBoolean(m, v6031)
	mBase = m.M
	v6066 = m.ExcPending
	if v6066 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1379
	}
L1379:
	;
	v6099 = int32(1)
	v6100 = v6012
	v6101 = v6065
	v6102 = v6017
	goto L1367
L1380:
	;
	if v6091-v6092 != 0 {
		goto L737
	} else {
		goto L1387
	}
L1381:
	;
	goto L1380
L1382:
	;
	v6076 = v6032
	v6077 = v6067
	goto L1383
L1383:
	;
	v6080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6077)+1)))
	v6081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6076)+1)))
	if v6081 == int32(0) {
		v6091 = v6081
		v6092 = v6080
		goto L1381
	} else {
		goto L1385
	}
L1384:
	;
	v6091 = v6081
	v6092 = v6080
	goto L1381
L1385:
	;
	v6084 = int32(1)
	if v6081 == v6080 {
		v6076 = v6076 + v6084
		v6077 = v6077 + v6084
		goto L1383
	} else {
		goto L1386
	}
L1386:
	;
	goto L1384
L1387:
	;
	if v6012&int32(1) != 0 {
		goto L738
	} else {
		goto L1388
	}
L1388:
	;
	v6097 = F_defGetBoolean(m, v6031)
	mBase = m.M
	v6098 = m.ExcPending
	if v6098 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1389
	}
L1389:
	;
	v6099 = v6011
	v6100 = int32(1)
	v6101 = v6016
	v6102 = v6097
	goto L1367
L1390:
	;
	goto L1366
L1391:
	;
	v6112 = v3481 + int32(_a_F_PostgresMain_101)
	goto L1393
L1392:
	;
	v6112 = int32(0)
	goto L1393
L1393:
	;
	if v6099&int32(1) != 0 {
		goto L1394
	} else {
		goto L1395
	}
L1394:
	;
	v6118 = v3481 + int32(496)
	goto L1396
L1395:
	;
	v6118 = int32(0)
	goto L1396
L1396:
	;
	v6143 = v6112
	v6147 = v6101
	v6148 = v6102
	v6158 = v6118
	goto L1362
L1397:
	;
	v6172 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v6173 = *(*int32)(unsafe.Add(mBase, uint32(v6172)+88))
	if v6173 != 0 {
		goto L1401
	} else {
		goto L1402
	}
L1398:
	;
	v8300 = int32(_a_F_PostgresMain_135)
	goto L726
L1399:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6354 = m.ExcPending
	if v6354 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1456
	}
L1400:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6332 = m.ExcPending
	if v6332 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1451
	}
L1401:
	;
	v6176 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])))
	if v6176 == int32(1) {
		goto L1407
	} else {
		goto L1408
	}
L1402:
	;
	goto L1403
L1403:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6315 = m.ExcPending
	if v6315 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1447
	}
L1404:
	;
	if v6143 == int32(0) {
		goto L1432
	} else {
		goto L1433
	}
L1405:
	;
	v6228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6226)+202)))
	if v6227 == v6228 {
		v6246 = v6162
		goto L1404
	} else {
		goto L1425
	}
L1406:
	;
	if v6186 != 0 {
		goto L1410
	} else {
		goto L1411
	}
L1407:
	;
	v6181 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v6182 = *(*int32)(unsafe.Add(mBase, uint32(v6181)+308))
	v6184 = base.B2i32(v6182 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])) = uint8(v6184)
	v6186 = v6184
	goto L1409
L1408:
	;
	v6186 = int32(0)
	goto L1409
L1409:
	;
	goto L1406
L1410:
	;
	v6188 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v6189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6188)+201)))
	if v6189 == int32(1) {
		goto L1400
	} else {
		goto L1413
	}
L1411:
	;
	goto L1412
L1412:
	;
	if v6158 == int32(0) {
		v6246 = v6162
		goto L1404
	} else {
		goto L1420
	}
L1413:
	;
	if v6158 == int32(0) {
		v6246 = v6162
		goto L1404
	} else {
		goto L1414
	}
L1414:
	;
	v6195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6158))))
	if v6195 != int32(1) {
		v6226 = v6188
		v6227 = int32(0)
		goto L1405
	} else {
		goto L1415
	}
L1415:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6201 = m.ExcPending
	if v6201 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1416
	}
L1416:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v6204 = m.ExcPending
	if v6204 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1417
	}
L1417:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_136), int32(0))
	mBase = m.M
	v6208 = m.ExcPending
	if v6208 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1418
	}
L1418:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_137), int32(984), int32(_a_F_PostgresMain_138))
	mBase = m.M
	v6213 = m.ExcPending
	if v6213 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1419
	}
L1419:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1420:
	;
	v6216 = int32(1)
	v6218 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v6219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6158))))
	if v6219 != v6216 {
		goto L1421
	} else {
		goto L1422
	}
L1421:
	;
	v6226 = v6218
	v6227 = int32(0)
	goto L1405
L1422:
	;
	goto L1423
L1423:
	;
	v6223 = *(*int32)(unsafe.Add(mBase, uint32(v6218)+92))
	if v6223 == int32(2) {
		goto L1399
	} else {
		goto L1424
	}
L1424:
	;
	v6226 = v6218
	v6227 = v6216
	goto L1405
L1425:
	;
	v6230 = int32(1)
	v6233 = base.AtomicRmwXchg32(m, v6226, int32(0), v6230)
	if v6233 != 0 {
		goto L1426
	} else {
		goto L1427
	}
L1426:
	;
	F_s_lock(m, v6226, int32(_a_F_PostgresMain_11))
	mBase = m.M
	v6236 = m.ExcPending
	if v6236 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1429
	}
L1427:
	;
	goto L1428
L1428:
	;
	v6238 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v6239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6158))))
	*(*uint8)(unsafe.Add(mBase, uint32(v6238)+202)) = uint8(v6239)
	v6241 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6238))), uint32(v6241))
	v6246 = v6230
	goto L1404
L1429:
	;
	goto L1428
L1430:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v6308 = m.ExcPending
	if v6308 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1446
	}
L1431:
	;
	v6273 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v6276 = base.AtomicRmwXchg32(m, v6273, int32(0), int32(1))
	if v6276 != 0 {
		goto L1440
	} else {
		goto L1441
	}
L1432:
	;
	if v6246 == int32(0) {
		goto L1430
	} else {
		goto L1439
	}
L1433:
	;
	v6251 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v6252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6251)+136)))
	v6253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6143))))
	if v6252 == v6253 {
		goto L1432
	} else {
		goto L1434
	}
L1434:
	;
	v6257 = base.AtomicRmwXchg32(m, v6251, int32(0), int32(1))
	if v6257 != 0 {
		goto L1435
	} else {
		goto L1436
	}
L1435:
	;
	F_s_lock(m, v6251, int32(_a_F_PostgresMain_11))
	mBase = m.M
	v6260 = m.ExcPending
	if v6260 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1438
	}
L1436:
	;
	goto L1437
L1437:
	;
	v6262 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v6263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6143))))
	*(*uint8)(unsafe.Add(mBase, uint32(v6262)+136)) = uint8(v6263)
	v6265 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6262))), uint32(v6265))
	goto L1431
L1438:
	;
	goto L1437
L1439:
	;
	goto L1431
L1440:
	;
	F_s_lock(m, v6273, int32(_a_F_PostgresMain_11))
	mBase = m.M
	v6279 = m.ExcPending
	if v6279 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1443
	}
L1441:
	;
	goto L1442
L1442:
	;
	v6280 = int32(_a_F_PostgresMain_139)
	v6281 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v6282 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v6281)+12)) = uint16(v6282)
	v6284 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6273))), uint32(v6284))
	*(*int32)(unsafe.Add(mBase, uint32(v6165)+32)) = int32(_a_F_PostgresMain_140)
	v6290 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	*(*int32)(unsafe.Add(mBase, uint32(v6165)+36)) = v6290 + int32(24)
	v6295 = v6165 + int32(48)
	v6299 = F_pg_sprintf(m, v6295, int32(_a_F_PostgresMain_141), v6165+int32(32))
	mBase = m.M
	v6300 = m.ExcPending
	if v6300 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1444
	}
L1443:
	;
	goto L1442
L1444:
	;
	v6302 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	F_SaveSlotToPath(m, v6302, v6295, int32(21))
	mBase = m.M
	v6305 = m.ExcPending
	if v6305 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1445
	}
L1445:
	;
	goto L1430
L1446:
	;
	m.G0 = v6165 + int32(1072)
	goto L1398
L1447:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v6318 = m.ExcPending
	if v6318 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1448
	}
L1448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6165))) = int32(_a_F_PostgresMain_135)
	F_errmsg(m, int32(_a_F_PostgresMain_142), v6165)
	mBase = m.M
	v6323 = m.ExcPending
	if v6323 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1449
	}
L1449:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_137), int32(962), int32(_a_F_PostgresMain_138))
	mBase = m.M
	v6328 = m.ExcPending
	if v6328 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1450
	}
L1450:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1451:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v6335 = m.ExcPending
	if v6335 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1452
	}
L1452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6165)+16)) = v6161
	F_errmsg(m, int32(_a_F_PostgresMain_143), v6165+int32(16))
	mBase = m.M
	v6341 = m.ExcPending
	if v6341 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1453
	}
L1453:
	;
	v6344 = F_errdetail(m, int32(_a_F_PostgresMain_144), int32(0))
	mBase = m.M
	v6345 = m.ExcPending
	if v6345 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1454
	}
L1454:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_137), int32(974), int32(_a_F_PostgresMain_138))
	mBase = m.M
	v6350 = m.ExcPending
	if v6350 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1455
	}
L1455:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1456:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v6357 = m.ExcPending
	if v6357 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1457
	}
L1457:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_145), int32(0))
	mBase = m.M
	v6361 = m.ExcPending
	if v6361 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1458
	}
L1458:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_137), int32(996), int32(_a_F_PostgresMain_138))
	mBase = m.M
	v6366 = m.ExcPending
	if v6366 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1459
	}
L1459:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1460:
	;
	v6372 = *(*int32)(unsafe.Add(mBase, uint32(v4719)+4))
	if v6372 == int32(0) {
		goto L1461
	} else {
		goto L1462
	}
L1461:
	;
	v6375 = m.G0
	v6377 = v6375 - int32(144)
	m.G0 = v6377
	*(*int32)(unsafe.Add(mBase, uint32(v6377)+120)) = int32(414)
	*(*int32)(unsafe.Add(mBase, uint32(v6377)+116)) = int32(1107)
	v6383 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6377)+112)) = v6383
	v6387 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[198]))
	v6391 = F_XLogReaderAllocate(m, v6387, v6377+int32(112), v6383)
	mBase = m.M
	v6392 = m.ExcPending
	if v6392 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1464
	}
L1462:
	;
	goto L1463
L1463:
	;
	F_CheckLogicalDecodingRequirements(m, int32(0))
	mBase = m.M
	v6881 = m.ExcPending
	if v6881 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1591
	}
L1464:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[113])) = v6391
	if v6391 != 0 {
		goto L1472
	} else {
		goto L1473
	}
L1465:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v6878 = m.ExcPending
	if v6878 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1590
	}
L1466:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6858 = m.ExcPending
	if v6858 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1587
	}
L1467:
	;
	v6701 = *(*int32)(unsafe.Add(mBase, uint32(v4719)+8))
	if v6701 != 0 {
		goto L1551
	} else {
		goto L1552
	}
L1468:
	;
	v6593 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[40]))
	v6594 = *(*int32)(unsafe.Add(mBase, uint32(v6593)+4))
	if v6594 != int32(2) {
		goto L1526
	} else {
		goto L1527
	}
L1469:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[199])) = v6575
	v6579 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[200])) = uint8(v6579)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[201])) = uint8(v6579)
	v6585 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[202])))
	if v6585 != int32(1) {
		goto L1468
	} else {
		goto L1524
	}
L1470:
	;
	v6575 = int64(0)
	goto L1469
L1471:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6558 = m.ExcPending
	if v6558 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1520
	}
L1472:
	;
	v6394 = *(*int32)(unsafe.Add(mBase, uint32(v4719)+8))
	if v6394 != 0 {
		goto L1475
	} else {
		goto L1476
	}
L1473:
	;
	goto L1474
L1474:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6538 = m.ExcPending
	if v6538 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1515
	}
L1475:
	;
	v6395 = int32(1)
	F_ReplicationSlotAcquire(m, v6394, v6395, v6395)
	mBase = m.M
	v6398 = m.ExcPending
	if v6398 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1478
	}
L1476:
	;
	goto L1477
L1477:
	;
	v6405 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])))
	if v6405 == int32(1) {
		goto L1481
	} else {
		goto L1482
	}
L1478:
	;
	v6400 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v6401 = *(*int32)(unsafe.Add(mBase, uint32(v6400)+88))
	if v6401 != 0 {
		goto L1471
	} else {
		goto L1479
	}
L1479:
	;
	goto L1477
L1480:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[36])) = uint8(v6415)
	if v6415 != 0 {
		goto L1485
	} else {
		goto L1486
	}
L1481:
	;
	v6410 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v6411 = *(*int32)(unsafe.Add(mBase, uint32(v6410)+308))
	v6413 = base.B2i32(v6411 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])) = uint8(v6413)
	v6415 = v6413
	goto L1483
L1482:
	;
	v6415 = int32(0)
	goto L1483
L1483:
	;
	goto L1480
L1484:
	;
	v6465 = *(*int32)(unsafe.Add(mBase, uint32(v4719)+12))
	if v6465 != 0 {
		goto L1500
	} else {
		goto L1501
	}
L1485:
	;
	v6420 = F_GetWalRcvFlushRecPtr(m, int32(0), v6377+int32(128))
	mBase = m.M
	v6421 = m.ExcPending
	if v6421 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1488
	}
L1486:
	;
	goto L1487
L1487:
	;
	v6434 = v6377 + int32(124)
	v6436 = int32(_a_F_PostgresMain_105)
	v6437 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v6438 = int64(0)
	v6441 = base.AtomicRmwCmpxchg64(m, v6437, int32(272), v6438, v6438)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[176])) = v6441
	v6443 = int32(0)
	v6446 = base.AtomicRmwOr32(m, v6443, int32(_a_F_PostgresMain_106), v6443)
	v6449 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v6453 = base.AtomicRmwCmpxchg64(m, v6449, int32(264), v6438, v6438)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[177])) = v6453
	if v6434 != 0 {
		goto L1497
	} else {
		goto L1498
	}
L1488:
	;
	v6424 = F_GetXLogReplayRecPtr(m, v6377+int32(80))
	mBase = m.M
	v6425 = m.ExcPending
	if v6425 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1489
	}
L1489:
	;
	v6426 = *(*int32)(unsafe.Add(mBase, uint32(v6377)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v6377)+124)) = v6426
	if base.Ui64(v6424) < base.Ui64(v6420) {
		goto L1490
	} else {
		goto L1491
	}
L1490:
	;
	v6429 = v6420
	goto L1492
L1491:
	;
	v6429 = v6424
	goto L1492
L1492:
	;
	v6430 = *(*int32)(unsafe.Add(mBase, uint32(v6377)+128))
	if v6426 == v6430 {
		goto L1493
	} else {
		goto L1494
	}
L1493:
	;
	v6432 = v6429
	goto L1495
L1494:
	;
	v6432 = v6424
	goto L1495
L1495:
	;
	v6464 = v6432
	goto L1484
L1496:
	;
	v6464 = v6460
	goto L1484
L1497:
	;
	v6456 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v6457 = *(*int32)(unsafe.Add(mBase, uint32(v6456)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v6434))) = v6457
	goto L1499
L1498:
	;
	goto L1499
L1499:
	;
	v6460 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[176]))
	goto L1496
L1500:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[203])) = v6465
	v6468 = *(*int32)(unsafe.Add(mBase, uint32(v6377)+124))
	if v6465 == v6468 {
		goto L1503
	} else {
		goto L1504
	}
L1501:
	;
	goto L1502
L1502:
	;
	v6521 = *(*int32)(unsafe.Add(mBase, uint32(v6377)+124))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[203])) = v6521
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[199])) = int64(0)
	v6527 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[202])) = uint8(v6527)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[200])) = uint8(v6527)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[201])) = uint8(v6527)
	goto L1468
L1503:
	;
	v6471 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[202])) = uint8(v6471)
	goto L1470
L1504:
	;
	goto L1505
L1505:
	;
	v6474 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[202])) = uint8(v6474)
	v6476 = F_readTimeLineHistory(m, v6468)
	mBase = m.M
	v6477 = m.ExcPending
	if v6477 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1506
	}
L1506:
	;
	v6478 = *(*int32)(unsafe.Add(mBase, uint32(v4719)+12))
	v6480 = F_tliSwitchPoint(m, v6478, v6476, int32(_a_F_PostgresMain_146))
	mBase = m.M
	v6481 = m.ExcPending
	if v6481 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1507
	}
L1507:
	;
	F_list_free_deep(m, v6476)
	mBase = m.M
	v6483 = m.ExcPending
	if v6483 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1508
	}
L1508:
	;
	if v6480 == int64(0) {
		goto L1470
	} else {
		goto L1509
	}
L1509:
	;
	v6486 = *(*int64)(unsafe.Add(mBase, uint32(v4719)+16))
	if base.Ui64(v6486) <= base.Ui64(v6480) {
		v6575 = v6480
		goto L1469
	} else {
		goto L1510
	}
L1510:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6491 = m.ExcPending
	if v6491 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1511
	}
L1511:
	;
	v6492 = *(*int64)(unsafe.Add(mBase, uint32(v4719)+16))
	v6493 = *(*int32)(unsafe.Add(mBase, uint32(v4719)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v6377)+56)) = v6493
	*(*uint32)(unsafe.Add(mBase, uint32(v6377)+52)) = uint32(v6492)
	v6497 = int64(base.Ui64(v6492) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6377)+48)) = uint32(v6497)
	F_errmsg(m, int32(_a_F_PostgresMain_147), v6377+int32(48))
	mBase = m.M
	v6503 = m.ExcPending
	if v6503 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1512
	}
L1512:
	;
	v6504 = *(*int32)(unsafe.Add(mBase, uint32(v4719)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v6377)+32)) = v6504
	*(*uint32)(unsafe.Add(mBase, uint32(v6377)+40)) = uint32(v6480)
	v6508 = int64(base.Ui64(v6480) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6377)+36)) = uint32(v6508)
	v6513 = F_errdetail(m, int32(_a_F_PostgresMain_148), v6377+int32(32))
	mBase = m.M
	v6514 = m.ExcPending
	if v6514 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1513
	}
L1513:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(949), int32(_a_F_PostgresMain_149))
	mBase = m.M
	v6519 = m.ExcPending
	if v6519 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1514
	}
L1514:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1515:
	;
	F_errcode(m, int32(_a_F_PostgresMain_150))
	mBase = m.M
	v6541 = m.ExcPending
	if v6541 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1516
	}
L1516:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_151), int32(0))
	mBase = m.M
	v6545 = m.ExcPending
	if v6545 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1517
	}
L1517:
	;
	v6548 = F_errdetail(m, int32(_a_F_PostgresMain_152), int32(0))
	mBase = m.M
	v6549 = m.ExcPending
	if v6549 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1518
	}
L1518:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(861), int32(_a_F_PostgresMain_149))
	mBase = m.M
	v6554 = m.ExcPending
	if v6554 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1519
	}
L1519:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1520:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v6561 = m.ExcPending
	if v6561 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1521
	}
L1521:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_153), int32(0))
	mBase = m.M
	v6565 = m.ExcPending
	if v6565 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1522
	}
L1522:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(878), int32(_a_F_PostgresMain_149))
	mBase = m.M
	v6570 = m.ExcPending
	if v6570 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1523
	}
L1523:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1524:
	;
	v6588 = *(*int64)(unsafe.Add(mBase, uint32(v4719)+16))
	if base.Ui64(v6575) <= base.Ui64(v6588) {
		goto L1467
	} else {
		goto L1525
	}
L1525:
	;
	goto L1468
L1526:
	;
	v6599 = base.AtomicRmwXchg32(m, v6593, int32(76), int32(1))
	if v6599 != 0 {
		goto L1529
	} else {
		goto L1530
	}
L1527:
	;
	goto L1528
L1528:
	;
	v6611 = v6377 + int32(128)
	F_pq_beginmessage(m, v6611, int32(87))
	mBase = m.M
	v6614 = m.ExcPending
	if v6614 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1533
	}
L1529:
	;
	F_s_lock(m, v6593+int32(76), int32(_a_F_PostgresMain_11))
	mBase = m.M
	v6604 = m.ExcPending
	if v6604 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1532
	}
L1530:
	;
	goto L1531
L1531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6593)+4)) = int32(2)
	v6607 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6593)+76)), uint32(v6607))
	goto L1528
L1532:
	;
	goto L1531
L1533:
	;
	F_enlargeStringInfo(m, v6611, int32(1))
	mBase = m.M
	v6617 = m.ExcPending
	if v6617 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1534
	}
L1534:
	;
	v6618 = *(*int32)(unsafe.Add(mBase, uint32(v6377)+132))
	v6619 = *(*int32)(unsafe.Add(mBase, uint32(v6377)+128))
	v6621 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6618+v6619))) = uint8(v6621)
	*(*int32)(unsafe.Add(mBase, uint32(v6377)+132)) = v6618 + int32(1)
	F_enlargeStringInfo(m, v6611, int32(2))
	mBase = m.M
	v6628 = m.ExcPending
	if v6628 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1535
	}
L1535:
	;
	v6629 = *(*int32)(unsafe.Add(mBase, uint32(v6377)+132))
	v6630 = *(*int32)(unsafe.Add(mBase, uint32(v6377)+128))
	v6632 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v6629+v6630))) = uint16(v6632)
	*(*int32)(unsafe.Add(mBase, uint32(v6377)+132)) = v6629 + int32(2)
	F_pq_endmessage(m, v6611)
	mBase = m.M
	v6638 = m.ExcPending
	if v6638 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1536
	}
L1536:
	;
	v6640 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[109]))
	v6641 = *(*int32)(unsafe.Add(mBase, uint32(v6640)+4))
	v6642 = m.T0[v6641].(func(*base.Module) int32)(m)
	mBase = m.M
	v6643 = m.ExcPending
	if v6643 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1537
	}
L1537:
	;
	v6644 = *(*int64)(unsafe.Add(mBase, uint32(v4719)+16))
	if base.Ui64(v6464) < base.Ui64(v6644) {
		goto L1466
	} else {
		goto L1538
	}
L1538:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[204])) = v6644
	v6649 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[40]))
	v6652 = base.AtomicRmwXchg32(m, v6649, int32(76), int32(1))
	if v6652 != 0 {
		goto L1539
	} else {
		goto L1540
	}
L1539:
	;
	F_s_lock(m, v6649+int32(76), int32(_a_F_PostgresMain_11))
	mBase = m.M
	v6657 = m.ExcPending
	if v6657 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1542
	}
L1540:
	;
	goto L1541
L1541:
	;
	v6659 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[40]))
	v6661 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[204]))
	*(*int64)(unsafe.Add(mBase, uint32(v6659)+8)) = v6661
	v6663 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6659)+76)), uint32(v6663))
	F_SyncRepInitConfig(m)
	mBase = m.M
	v6667 = m.ExcPending
	if v6667 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1543
	}
L1542:
	;
	goto L1541
L1543:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[115])) = int32(1)
	F_WalSndLoop(m, int32(1113))
	mBase = m.M
	v6673 = m.ExcPending
	if v6673 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1544
	}
L1544:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[115])) = int32(0)
	v6678 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[117]))
	if v6678 != 0 {
		goto L1465
	} else {
		goto L1545
	}
L1545:
	;
	v6680 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[40]))
	v6681 = *(*int32)(unsafe.Add(mBase, uint32(v6680)+4))
	if v6681 == int32(0) {
		goto L1467
	} else {
		goto L1546
	}
L1546:
	;
	v6686 = base.AtomicRmwXchg32(m, v6680, int32(76), int32(1))
	if v6686 != 0 {
		goto L1547
	} else {
		goto L1548
	}
L1547:
	;
	F_s_lock(m, v6680+int32(76), int32(_a_F_PostgresMain_11))
	mBase = m.M
	v6691 = m.ExcPending
	if v6691 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1550
	}
L1548:
	;
	goto L1549
L1549:
	;
	v6692 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6680)+4)) = v6692
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6680)+76)), uint32(v6692))
	goto L1467
L1550:
	;
	goto L1549
L1551:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v6703 = m.ExcPending
	if v6703 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1554
	}
L1552:
	;
	goto L1553
L1553:
	;
	v6705 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[202])))
	if v6705 == int32(1) {
		goto L1555
	} else {
		goto L1556
	}
L1554:
	;
	goto L1553
L1555:
	;
	v6709 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[199]))
	*(*uint32)(unsafe.Add(mBase, uint32(v6377)+20)) = uint32(v6709)
	v6711 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v6377)+62)) = uint16(v6711)
	v6714 = int64(base.Ui64(v6709) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6377)+16)) = uint32(v6714)
	v6717 = v6377 + int32(80)
	v6722 = F_pg_snprintf(m, v6717, int32(18), int32(_a_F_PostgresMain_103), v6377+int32(16))
	mBase = m.M
	v6723 = m.ExcPending
	if v6723 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1558
	}
L1556:
	;
	goto L1557
L1557:
	;
	F_EndReplicationCommand(m, int32(_a_F_PostgresMain_154))
	mBase = m.M
	v6851 = m.ExcPending
	if v6851 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1586
	}
L1558:
	;
	v6725 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v6726 = m.ExcPending
	if v6726 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1559
	}
L1559:
	;
	v6728 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v6729 = m.ExcPending
	if v6729 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1560
	}
L1560:
	;
	F_TupleDescInitBuiltinEntry(m, v6728, int32(1), int32(_a_F_PostgresMain_155), int32(20))
	mBase = m.M
	v6734 = m.ExcPending
	if v6734 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1561
	}
L1561:
	;
	F_TupleDescInitBuiltinEntry(m, v6728, int32(2), int32(_a_F_PostgresMain_156), int32(25))
	mBase = m.M
	v6739 = m.ExcPending
	if v6739 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1562
	}
L1562:
	;
	v6740 = int32(0)
	v6749 = *(*int32)(unsafe.Add(mBase, uint32(v6728)))
	if v6740 < v6749 {
		goto L1564
	} else {
		goto L1565
	}
L1563:
	;
	v6828 = F_begin_tup_output_tupdesc(m, v6725, v6728, int32(_a_F_PostgresMain_110))
	mBase = m.M
	v6829 = m.ExcPending
	if v6829 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1582
	}
L1564:
	;
	v6753 = v6728 + int32(28)
	v6760 = v6740
	v6761 = v6749
	v6763 = v6740
	goto L1568
L1565:
	;
	v6817 = v6740
	v6824 = v6749
	goto L1566
L1566:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6728)+20)) = v6824
	*(*int32)(unsafe.Add(mBase, uint32(v6728)+16)) = v6817
	goto L1563
L1567:
	;
	v6817 = v6811
	v6824 = v6790
	goto L1566
L1568:
	;
	v6769 = v6753 + v6749<<(uint(int32(3))%32) + v6760*int32(100)
	v6772 = v6753 + v6760<<(uint(int32(3))%32)
	if v6749 != v6761 {
		v6790 = v6761
		goto L1570
	} else {
		goto L1571
	}
L1569:
	;
	v6811 = v6749
	goto L1567
L1570:
	;
	v6791 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6772)+2)))
	if v6791 <= int32(0) {
		v6811 = v6760
		goto L1567
	} else {
		goto L1578
	}
L1571:
	;
	v6774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6772)+7)))
	if v6774 != int32(118) {
		goto L1572
	} else {
		goto L1573
	}
L1572:
	;
	v6790 = v6760
	goto L1570
L1573:
	;
	v6777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6772)+4)))
	if v6777 != int32(1) {
		goto L1572
	} else {
		goto L1574
	}
L1574:
	;
	v6780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6772)+6)))
	if v6780&int32(6) != 0 {
		goto L1572
	} else {
		goto L1575
	}
L1575:
	;
	v6783 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6772)+2)))
	if v6783 <= int32(0) {
		goto L1572
	} else {
		goto L1576
	}
L1576:
	;
	v6786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6769)+90)))
	if v6786 != int32(118) {
		v6790 = v6749
		goto L1570
	} else {
		goto L1577
	}
L1577:
	;
	goto L1572
L1578:
	;
	v6794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6769)+90)))
	if v6794 == int32(118) {
		v6811 = v6760
		goto L1567
	} else {
		goto L1579
	}
L1579:
	;
	v6797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6772)+5)))
	v6803 = (v6763 + v6797 - int32(1)) & (int32(0) - v6797)
	if int32(_a_F_PostgresMain_111) < v6803 {
		v6811 = v6760
		goto L1567
	} else {
		goto L1580
	}
L1580:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v6772))) = uint16(v6803)
	v6809 = v6760 + int32(1)
	if v6809 != v6749 {
		v6760 = v6809
		v6761 = v6790
		v6763 = v6803 + v6791
		goto L1568
	} else {
		goto L1581
	}
L1581:
	;
	goto L1569
L1582:
	;
	v6831 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_PostgresMain[205])))
	*(*int64)(unsafe.Add(mBase, uint32(v6377)+64)) = v6831
	v6833 = F_cstring_to_text(m, v6717)
	mBase = m.M
	v6834 = m.ExcPending
	if v6834 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1583
	}
L1583:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6377)+72)) = base.I64_extend_i32_u(v6833)
	F_do_tup_output(m, v6828, v6377-int32(-64), v6377+int32(62))
	mBase = m.M
	v6842 = m.ExcPending
	if v6842 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1584
	}
L1584:
	;
	F_end_tup_output(m, v6828)
	mBase = m.M
	v6844 = m.ExcPending
	if v6844 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1585
	}
L1585:
	;
	goto L1557
L1586:
	;
	m.G0 = v6377 + int32(144)
	v8300 = v6367
	goto L726
L1587:
	;
	v6859 = *(*int64)(unsafe.Add(mBase, uint32(v4719)+16))
	*(*uint32)(unsafe.Add(mBase, uint32(v6377)+4)) = uint32(v6859)
	v6861 = int64(32)
	v6862 = int64(base.Ui64(v6859) >> (uint(v6861) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6377))) = uint32(v6862)
	*(*uint32)(unsafe.Add(mBase, uint32(v6377)+12)) = uint32(v6464)
	v6866 = int64(base.Ui64(v6464) >> (uint(v6861) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6377)+8)) = uint32(v6866)
	F_errmsg(m, int32(_a_F_PostgresMain_157), v6377)
	mBase = m.M
	v6870 = m.ExcPending
	if v6870 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1588
	}
L1588:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(993), int32(_a_F_PostgresMain_149))
	mBase = m.M
	v6875 = m.ExcPending
	if v6875 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1589
	}
L1589:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1590:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1591:
	;
	v6882 = *(*int32)(unsafe.Add(mBase, uint32(v4719)+8))
	v6883 = int32(1)
	F_ReplicationSlotAcquire(m, v6882, v6883, v6883)
	mBase = m.M
	v6886 = m.ExcPending
	if v6886 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1592
	}
L1592:
	;
	v6888 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[36])))
	if v6888 != int32(1) {
		goto L1593
	} else {
		goto L1594
	}
L1593:
	;
	v6920 = *(*int32)(unsafe.Add(mBase, uint32(v4719)+24))
	v6921 = *(*int64)(unsafe.Add(mBase, uint32(v4719)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[179]))) = int32(414)
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[192]))) = int32(1107)
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[178]))) = int32(1108)
	v6935 = F_CreateDecodingContext(m, v6921, v6920, int32(0), v3481+int32(_a_F_PostgresMain_104), int32(1109), int32(1110), int32(1111))
	mBase = m.M
	v6936 = m.ExcPending
	if v6936 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1606
	}
L1594:
	;
	v6893 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])))
	if v6893 == int32(1) {
		goto L1596
	} else {
		goto L1597
	}
L1595:
	;
	if v6903 != 0 {
		goto L1593
	} else {
		goto L1599
	}
L1596:
	;
	v6898 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[29]))
	v6899 = *(*int32)(unsafe.Add(mBase, uint32(v6898)+308))
	v6901 = base.B2i32(v6899 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[28])) = uint8(v6901)
	v6903 = v6901
	goto L1598
L1597:
	;
	v6903 = int32(0)
	goto L1598
L1598:
	;
	goto L1595
L1599:
	;
	v6906 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6907 = m.ExcPending
	if v6907 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1600
	}
L1600:
	;
	if v6906 != 0 {
		goto L1601
	} else {
		goto L1602
	}
L1601:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_158), int32(0))
	mBase = m.M
	v6911 = m.ExcPending
	if v6911 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1604
	}
L1602:
	;
	goto L1603
L1603:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[117])) = int32(1)
	goto L1593
L1604:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1536), int32(_a_F_PostgresMain_159))
	mBase = m.M
	v6916 = m.ExcPending
	if v6916 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1605
	}
L1605:
	;
	goto L1603
L1606:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[206])) = v6935
	v6939 = *(*int32)(unsafe.Add(mBase, uint32(v6935)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[113])) = v6939
	v6942 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[40]))
	v6943 = *(*int32)(unsafe.Add(mBase, uint32(v6942)+4))
	if v6943 != int32(2) {
		goto L1607
	} else {
		goto L1608
	}
L1607:
	;
	v6948 = base.AtomicRmwXchg32(m, v6942, int32(76), int32(1))
	if v6948 != 0 {
		goto L1610
	} else {
		goto L1611
	}
L1608:
	;
	goto L1609
L1609:
	;
	v6960 = v3481 + int32(496)
	F_pq_beginmessage(m, v6960, int32(87))
	mBase = m.M
	v6963 = m.ExcPending
	if v6963 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1614
	}
L1610:
	;
	F_s_lock(m, v6942+int32(76), int32(_a_F_PostgresMain_11))
	mBase = m.M
	v6953 = m.ExcPending
	if v6953 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1613
	}
L1611:
	;
	goto L1612
L1612:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6942)+4)) = int32(2)
	v6956 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6942)+76)), uint32(v6956))
	goto L1609
L1613:
	;
	goto L1612
L1614:
	;
	F_enlargeStringInfo(m, v6960, int32(1))
	mBase = m.M
	v6966 = m.ExcPending
	if v6966 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1615
	}
L1615:
	;
	v6967 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+500))
	v6968 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+496))
	v6970 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6967+v6968))) = uint8(v6970)
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+500)) = v6967 + int32(1)
	F_enlargeStringInfo(m, v6960, int32(2))
	mBase = m.M
	v6977 = m.ExcPending
	if v6977 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1616
	}
L1616:
	;
	v6978 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+500))
	v6979 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+496))
	v6981 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v6978+v6979))) = uint16(v6981)
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+500)) = v6978 + int32(2)
	F_pq_endmessage(m, v6960)
	mBase = m.M
	v6987 = m.ExcPending
	if v6987 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1617
	}
L1617:
	;
	v6989 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[109]))
	v6990 = *(*int32)(unsafe.Add(mBase, uint32(v6989)+4))
	v6991 = m.T0[v6990].(func(*base.Module) int32)(m)
	mBase = m.M
	v6992 = m.ExcPending
	if v6992 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1618
	}
L1618:
	;
	v6994 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[206]))
	v6995 = *(*int32)(unsafe.Add(mBase, uint32(v6994)+8))
	v6997 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v6998 = *(*int64)(unsafe.Add(mBase, uint32(v6997)+104))
	F_XLogBeginRead(m, v6995, v6998)
	mBase = m.M
	v7000 = m.ExcPending
	if v7000 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1619
	}
L1619:
	;
	v7003 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v7004 = *(*int64)(unsafe.Add(mBase, uint32(v7003)+120))
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[204])) = v7004
	v7007 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[40]))
	v7010 = base.AtomicRmwXchg32(m, v7007, int32(76), int32(1))
	if v7010 != 0 {
		goto L1620
	} else {
		goto L1621
	}
L1620:
	;
	F_s_lock(m, v7007+int32(76), int32(_a_F_PostgresMain_11))
	mBase = m.M
	v7015 = m.ExcPending
	if v7015 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1623
	}
L1621:
	;
	goto L1622
L1622:
	;
	v7017 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[40]))
	v7019 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[114]))
	v7020 = *(*int64)(unsafe.Add(mBase, uint32(v7019)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v7017)+8)) = v7020
	v7022 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7017)+76)), uint32(v7022))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[115])) = int32(1)
	F_SyncRepInitConfig(m)
	mBase = m.M
	v7029 = m.ExcPending
	if v7029 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1624
	}
L1623:
	;
	goto L1622
L1624:
	;
	F_WalSndLoop(m, int32(1112))
	mBase = m.M
	v7032 = m.ExcPending
	if v7032 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1625
	}
L1625:
	;
	v7034 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[206]))
	F_FreeDecodingContext(m, v7034)
	mBase = m.M
	v7036 = m.ExcPending
	if v7036 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1626
	}
L1626:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v7038 = m.ExcPending
	if v7038 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1627
	}
L1627:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[115])) = int32(0)
	v7043 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[117]))
	if v7043 != 0 {
		goto L736
	} else {
		goto L1628
	}
L1628:
	;
	v7045 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[40]))
	v7046 = *(*int32)(unsafe.Add(mBase, uint32(v7045)+4))
	if v7046 != 0 {
		goto L1629
	} else {
		goto L1630
	}
L1629:
	;
	v7049 = base.AtomicRmwXchg32(m, v7045, int32(76), int32(1))
	if v7049 != 0 {
		goto L1632
	} else {
		goto L1633
	}
L1630:
	;
	goto L1631
L1631:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[195]))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+uint32(_c_F_PostgresMain[194]))) = int32(56)
	F_EndCommand(m, v3481+int32(_a_F_PostgresMain_101), int32(2))
	mBase = m.M
	v7068 = m.ExcPending
	if v7068 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1636
	}
L1632:
	;
	F_s_lock(m, v7045+int32(76), int32(_a_F_PostgresMain_11))
	mBase = m.M
	v7054 = m.ExcPending
	if v7054 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1635
	}
L1633:
	;
	goto L1634
L1634:
	;
	v7055 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7045)+4)) = v7055
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7045)+76)), uint32(v7055))
	goto L1631
L1635:
	;
	goto L1634
L1636:
	;
	v8300 = v6367
	goto L726
L1637:
	;
	v7073 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+424))
	v7075 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v7076 = m.ExcPending
	if v7076 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1638
	}
L1638:
	;
	v7078 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v7079 = m.ExcPending
	if v7079 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1639
	}
L1639:
	;
	F_TupleDescInitBuiltinEntry(m, v7078, int32(1), int32(_a_F_PostgresMain_160), int32(25))
	mBase = m.M
	v7084 = m.ExcPending
	if v7084 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1640
	}
L1640:
	;
	F_TupleDescInitBuiltinEntry(m, v7078, int32(2), int32(_a_F_PostgresMain_161), int32(25))
	mBase = m.M
	v7089 = m.ExcPending
	if v7089 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1641
	}
L1641:
	;
	v7090 = int32(0)
	v7099 = *(*int32)(unsafe.Add(mBase, uint32(v7078)))
	if v7090 < v7099 {
		goto L1643
	} else {
		goto L1644
	}
L1642:
	;
	v7177 = *(*int32)(unsafe.Add(mBase, uint32(v7073)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+352)) = v7177
	v7180 = v3481 + int32(_a_F_PostgresMain_104)
	v7185 = F_pg_snprintf(m, v7180, int32(64), int32(_a_F_PostgresMain_162), v3481+int32(352))
	mBase = m.M
	v7186 = m.ExcPending
	if v7186 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1661
	}
L1643:
	;
	v7103 = v7078 + int32(28)
	v7110 = v7090
	v7111 = v7099
	v7113 = v7090
	goto L1647
L1644:
	;
	v7167 = v7090
	v7174 = v7099
	goto L1645
L1645:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7078)+20)) = v7174
	*(*int32)(unsafe.Add(mBase, uint32(v7078)+16)) = v7167
	goto L1642
L1646:
	;
	v7167 = v7161
	v7174 = v7140
	goto L1645
L1647:
	;
	v7119 = v7103 + v7099<<(uint(int32(3))%32) + v7110*int32(100)
	v7122 = v7103 + v7110<<(uint(int32(3))%32)
	if v7099 != v7111 {
		v7140 = v7111
		goto L1649
	} else {
		goto L1650
	}
L1648:
	;
	v7161 = v7099
	goto L1646
L1649:
	;
	v7141 = int32(*(*int16)(unsafe.Add(mBase, uint32(v7122)+2)))
	if v7141 <= int32(0) {
		v7161 = v7110
		goto L1646
	} else {
		goto L1657
	}
L1650:
	;
	v7124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7122)+7)))
	if v7124 != int32(118) {
		goto L1651
	} else {
		goto L1652
	}
L1651:
	;
	v7140 = v7110
	goto L1649
L1652:
	;
	v7127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7122)+4)))
	if v7127 != int32(1) {
		goto L1651
	} else {
		goto L1653
	}
L1653:
	;
	v7130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7122)+6)))
	if v7130&int32(6) != 0 {
		goto L1651
	} else {
		goto L1654
	}
L1654:
	;
	v7133 = int32(*(*int16)(unsafe.Add(mBase, uint32(v7122)+2)))
	if v7133 <= int32(0) {
		goto L1651
	} else {
		goto L1655
	}
L1655:
	;
	v7136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7119)+90)))
	if v7136 != int32(118) {
		v7140 = v7099
		goto L1649
	} else {
		goto L1656
	}
L1656:
	;
	goto L1651
L1657:
	;
	v7144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7119)+90)))
	if v7144 == int32(118) {
		v7161 = v7110
		goto L1646
	} else {
		goto L1658
	}
L1658:
	;
	v7147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7122)+5)))
	v7153 = (v7113 + v7147 - int32(1)) & (int32(0) - v7147)
	if int32(_a_F_PostgresMain_111) < v7153 {
		v7161 = v7110
		goto L1646
	} else {
		goto L1659
	}
L1659:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v7122))) = uint16(v7153)
	v7159 = v7110 + int32(1)
	if v7159 != v7099 {
		v7110 = v7159
		v7111 = v7140
		v7113 = v7153 + v7141
		goto L1647
	} else {
		goto L1660
	}
L1660:
	;
	goto L1648
L1661:
	;
	v7187 = *(*int32)(unsafe.Add(mBase, uint32(v7073)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+336)) = v7187
	v7190 = v3481 + int32(_a_F_PostgresMain_101)
	v7195 = F_pg_snprintf(m, v7190, int32(1024), int32(_a_F_PostgresMain_163), v3481+int32(336))
	mBase = m.M
	v7196 = m.ExcPending
	if v7196 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1662
	}
L1662:
	;
	v7198 = *(*int32)(unsafe.Add(mBase, uint32(v7075)+4))
	m.T0[v7198].(func(*base.Module, int32, int32, int32))(m, v7075, int32(1), v7078)
	mBase = m.M
	v7200 = m.ExcPending
	if v7200 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1663
	}
L1663:
	;
	v7202 = v3481 + int32(464)
	F_pq_beginmessage(m, v7202, int32(68))
	mBase = m.M
	v7205 = m.ExcPending
	if v7205 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1664
	}
L1664:
	;
	F_enlargeStringInfo(m, v7202, int32(2))
	mBase = m.M
	v7208 = m.ExcPending
	if v7208 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1665
	}
L1665:
	;
	v7209 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+468))
	v7210 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+464))
	v7212 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v7209+v7210))) = uint16(v7212)
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+468)) = v7209 + int32(2)
	v7217 = F_strlen(m, v7180)
	mBase = m.M
	F_enlargeStringInfo(m, v7202, int32(4))
	mBase = m.M
	v7220 = m.ExcPending
	if v7220 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1666
	}
L1666:
	;
	v7221 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+468))
	v7222 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+464))
	v7226 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v7221+v7222))) = base.I32_rotr(v7217, int32(24))&v7226 | base.I32_rotr(v7217&v7226, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+468)) = v7221 + int32(4)
	F_appendBinaryStringInfo(m, v7202, v7180, v7217)
	mBase = m.M
	v7238 = m.ExcPending
	if v7238 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1667
	}
L1667:
	;
	v7240 = F_OpenTransientFile(m, v7190, int32(0))
	mBase = m.M
	v7241 = m.ExcPending
	if v7241 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1668
	}
L1668:
	;
	if v7240 < int32(0) {
		goto L735
	} else {
		goto L1669
	}
L1669:
	;
	v7244 = int64(0)
	v7246 = F___lseek(m, v7240, v7244, int32(2))
	mBase = m.M
	if v7246 < v7244 {
		goto L734
	} else {
		goto L1670
	}
L1670:
	;
	v7249 = int64(0)
	v7251 = F___lseek(m, v7240, v7249, int32(0))
	mBase = m.M
	if v7251 != v7249 {
		goto L733
	} else {
		goto L1671
	}
L1671:
	;
	F_enlargeStringInfo(m, v7202, int32(4))
	mBase = m.M
	v7256 = m.ExcPending
	if v7256 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1672
	}
L1672:
	;
	v7257 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+468))
	v7258 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+464))
	v7260 = base.I32_wrap_i64(v7246)
	v7261 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v7257+v7258))) = base.I32_rotr(v7260&v7261, int32(8)) | base.I32_rotr(v7260, int32(24))&v7261
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+468)) = v7257 + int32(4)
	if v7246 != int64(0) {
		goto L1673
	} else {
		goto L1674
	}
L1673:
	;
	v7310 = v7246
	goto L1676
L1674:
	;
	goto L1675
L1675:
	;
	v7378 = F_CloseTransientFile(m, v7240)
	mBase = m.M
	v7379 = m.ExcPending
	if v7379 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1682
	}
L1676:
	;
	v7315 = int32(_a_F_PostgresMain_164)
	v7316 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[112]))
	*(*int32)(unsafe.Add(mBase, uint32(v7316))) = int32(167772229)
	v7320 = v3481 + int32(496)
	v7322 = F_read(m, v7240, v7320, int32(_a_F_PostgresMain_16))
	mBase = m.M
	v7324 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[112]))
	v7325 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7324))) = v7325
	if v7322 < v7325 {
		goto L732
	} else {
		goto L1678
	}
L1677:
	;
	goto L1675
L1678:
	;
	if v7322 == int32(0) {
		goto L731
	} else {
		goto L1679
	}
L1679:
	;
	F_appendBinaryStringInfo(m, v3481+int32(464), v7320, v7322)
	mBase = m.M
	v7334 = m.ExcPending
	if v7334 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1680
	}
L1680:
	;
	v7336 = v7310 - base.I64_extend_i32_u(v7322)
	if int64(0) < v7336 {
		v7310 = v7336
		goto L1676
	} else {
		goto L1681
	}
L1681:
	;
	goto L1677
L1682:
	;
	if v7378 != 0 {
		goto L730
	} else {
		goto L1683
	}
L1683:
	;
	F_pq_endmessage(m, v3481+int32(464))
	mBase = m.M
	v7383 = m.ExcPending
	if v7383 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1684
	}
L1684:
	;
	v8300 = int32(_a_F_PostgresMain_99)
	goto L726
L1685:
	;
	v7391 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[207]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[163])) = v7391
	v7394 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50]))
	v7399 = F_AllocSetContextCreateInternal(m, v7394, int32(_a_F_PostgresMain_165), int32(0), int32(_a_F_PostgresMain_16), int32(_a_F_PostgresMain_17))
	mBase = m.M
	v7400 = m.ExcPending
	if v7400 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1686
	}
L1686:
	;
	v7401 = int32(_a_F_PostgresMain_166)
	v7402 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v7399
	v7406 = F_palloc0(m, int32(36))
	mBase = m.M
	v7407 = m.ExcPending
	if v7407 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1687
	}
L1687:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7406))) = v7399
	F_initStringInfo(m, v7406+int32(4))
	mBase = m.M
	v7412 = m.ExcPending
	if v7412 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1688
	}
L1688:
	;
	v7414 = F_MemoryContextAllocZero(m, v7399, int32(32))
	mBase = m.M
	v7415 = m.ExcPending
	if v7415 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1689
	}
L1689:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7414)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7414)+24)) = v7399
	v7421 = F_MemoryContextAllocExtended(m, v7399, int32(_a_F_PostgresMain_167), int32(5))
	mBase = m.M
	v7422 = m.ExcPending
	if v7422 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1690
	}
L1690:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7414)+12)) = int64(63329292795903)
	*(*int64)(unsafe.Add(mBase, uint32(v7414))) = int64(16384)
	*(*int32)(unsafe.Add(mBase, uint32(v7414)+20)) = v7421
	*(*int32)(unsafe.Add(mBase, uint32(v7406)+24)) = v7414
	v7430 = F_palloc0(m, int32(24))
	mBase = m.M
	v7431 = m.ExcPending
	if v7431 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1691
	}
L1691:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7430)+20)) = int32(461)
	*(*int32)(unsafe.Add(mBase, uint32(v7430)+16)) = int32(462)
	*(*int32)(unsafe.Add(mBase, uint32(v7430)+12)) = int32(463)
	*(*int32)(unsafe.Add(mBase, uint32(v7430)+8)) = int32(464)
	*(*int32)(unsafe.Add(mBase, uint32(v7430)+4)) = int32(465)
	*(*int32)(unsafe.Add(mBase, uint32(v7430))) = v7406
	v7444 = F_palloc(m, int32(112))
	mBase = m.M
	v7445 = m.ExcPending
	if v7445 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1692
	}
L1692:
	;
	v7447 = F_palloc(m, int32(68))
	mBase = m.M
	v7448 = m.ExcPending
	if v7448 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1693
	}
L1693:
	;
	v7449 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7447)+52)) = uint8(v7449)
	*(*int32)(unsafe.Add(mBase, uint32(v7447)+4)) = v7449
	*(*int32)(unsafe.Add(mBase, uint32(v7447))) = v7430
	if v7444 == v7449 {
		goto L1696
	} else {
		goto L1697
	}
L1694:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7444)+104)) = int32(_a_F_PostgresMain_168)
	*(*int32)(unsafe.Add(mBase, uint32(v7444)+100)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7444)+92)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7444)+88)) = int32(_a_F_PostgresMain_169)
	*(*int32)(unsafe.Add(mBase, uint32(v7444)+84)) = int32(_a_F_PostgresMain_170)
	*(*int32)(unsafe.Add(mBase, uint32(v7444)+80)) = int32(_a_F_PostgresMain_171)
	*(*int32)(unsafe.Add(mBase, uint32(v7444)+76)) = int32(_a_F_PostgresMain_172)
	*(*int32)(unsafe.Add(mBase, uint32(v7444)+72)) = int32(_a_F_PostgresMain_173)
	*(*int32)(unsafe.Add(mBase, uint32(v7444)+68)) = v7447
	v7541 = F_pg_cryptohash_create(m, int32(3))
	mBase = m.M
	v7542 = m.ExcPending
	if v7542 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1708
	}
L1695:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7468)+8)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v7468)+40)) = int32(1)
	v7474 = F_palloc0(m, int32(20))
	mBase = m.M
	v7475 = m.ExcPending
	if v7475 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1701
	}
L1696:
	;
	v7457 = F_palloc0(m, int32(68))
	mBase = m.M
	v7458 = m.ExcPending
	if v7458 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1699
	}
L1697:
	;
	goto L1698
L1698:
	;
	base.MemoryFill(m, v7444, int32(0), int32(68))
	v7468 = v7444
	goto L1695
L1699:
	;
	if v7457 == int32(0) {
		goto L1694
	} else {
		goto L1700
	}
L1700:
	;
	v7461 = *(*int32)(unsafe.Add(mBase, uint32(v7457)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v7457)+36)) = v7461 | int32(1)
	v7468 = v7457
	goto L1695
L1701:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7468)+52)) = v7474
	v7478 = F_palloc0(m, int32(28))
	mBase = m.M
	v7479 = m.ExcPending
	if v7479 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1702
	}
L1702:
	;
	v7481 = F_palloc(m, int32(640))
	mBase = m.M
	v7482 = m.ExcPending
	if v7482 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1703
	}
L1703:
	;
	v7484 = F_palloc(m, int32(256))
	mBase = m.M
	v7485 = m.ExcPending
	if v7485 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1704
	}
L1704:
	;
	v7487 = F_palloc(m, int32(64))
	mBase = m.M
	v7488 = m.ExcPending
	if v7488 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1705
	}
L1705:
	;
	v7489 = *(*int32)(unsafe.Add(mBase, uint32(v7468)+52))
	F_initStringInfo(m, v7489+int32(4))
	mBase = m.M
	v7493 = m.ExcPending
	if v7493 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1706
	}
L1706:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7468)+48)) = v7478
	*(*int32)(unsafe.Add(mBase, uint32(v7478))) = int32(64)
	v7497 = *(*int32)(unsafe.Add(mBase, uint32(v7468)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v7497)+4)) = v7481
	v7499 = *(*int32)(unsafe.Add(mBase, uint32(v7468)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v7499)+12)) = v7484
	v7501 = *(*int32)(unsafe.Add(mBase, uint32(v7468)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v7501)+16)) = v7487
	v7503 = *(*int32)(unsafe.Add(mBase, uint32(v7468)+48))
	v7504 = *(*int32)(unsafe.Add(mBase, uint32(v7503)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7504))) = int32(0)
	v7507 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7468)+56)) = uint8(v7507)
	*(*uint8)(unsafe.Add(mBase, uint32(v7468)+24)) = uint8(v7507)
	v7511 = F_makeStringInfo(m)
	mBase = m.M
	v7512 = m.ExcPending
	if v7512 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1707
	}
L1707:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7468)+60)) = v7511
	v7514 = *(*int32)(unsafe.Add(mBase, uint32(v7468)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v7468)+36)) = v7514 | int32(2)
	goto L1694
L1708:
	;
	if v7541 == int32(0) {
		goto L1709
	} else {
		goto L1710
	}
L1709:
	;
	v7547 = *(*int32)(unsafe.Add(mBase, uint32(v7430)+20))
	m.T0[v7547].(func(*base.Module, int32, int32, int32))(m, v7430, int32(_a_F_PostgresMain_151), int32(0))
	mBase = m.M
	v7549 = m.ExcPending
	if v7549 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1712
	}
L1710:
	;
	goto L1711
L1711:
	;
	v7550 = F_pg_cryptohash_init(m, v7541)
	mBase = m.M
	if v7550 < int32(0) {
		goto L1713
	} else {
		goto L1714
	}
L1712:
	;
	goto L1711
L1713:
	;
	v7555 = *(*int32)(unsafe.Add(mBase, uint32(v7430)+20))
	m.T0[v7555].(func(*base.Module, int32, int32, int32))(m, v7430, int32(_a_F_PostgresMain_174), int32(0))
	mBase = m.M
	v7557 = m.ExcPending
	if v7557 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1716
	}
L1714:
	;
	goto L1715
L1715:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7444)+108)) = v7541
	*(*int32)(unsafe.Add(mBase, uint32(v7406)+32)) = v7444
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v7402
	v7563 = v3481 + int32(496)
	F_pq_beginmessage(m, v7563, int32(71))
	mBase = m.M
	v7566 = m.ExcPending
	if v7566 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1717
	}
L1716:
	;
	goto L1715
L1717:
	;
	F_enlargeStringInfo(m, v7563, int32(1))
	mBase = m.M
	v7569 = m.ExcPending
	if v7569 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1718
	}
L1718:
	;
	v7570 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+500))
	v7571 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+496))
	v7573 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7570+v7571))) = uint8(v7573)
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+500)) = v7570 + int32(1)
	F_enlargeStringInfo(m, v7563, int32(2))
	mBase = m.M
	v7580 = m.ExcPending
	if v7580 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1719
	}
L1719:
	;
	v7581 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+500))
	v7582 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+496))
	v7584 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v7581+v7582))) = uint16(v7584)
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+500)) = v7581 + int32(2)
	F_pq_endmessage_reuse(m, v7563)
	mBase = m.M
	v7590 = m.ExcPending
	if v7590 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1720
	}
L1720:
	;
	v7592 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[109]))
	v7593 = *(*int32)(unsafe.Add(mBase, uint32(v7592)+4))
	v7594 = m.T0[v7593].(func(*base.Module) int32)(m)
	mBase = m.M
	v7595 = m.ExcPending
	if v7595 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1721
	}
L1721:
	;
	goto L1723
L1722:
	;
	v7731 = int32(_a_F_PostgresMain_166)
	v7732 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50]))
	v7734 = *(*int32)(unsafe.Add(mBase, uint32(v7406)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v7734
	v7736 = *(*int32)(unsafe.Add(mBase, uint32(v7406)+32))
	v7737 = *(*int32)(unsafe.Add(mBase, uint32(v7406)+4))
	v7738 = *(*int32)(unsafe.Add(mBase, uint32(v7406)+8))
	F_json_parse_manifest_incremental_chunk(m, v7736, v7737, v7738, int32(1))
	mBase = m.M
	v7741 = m.ExcPending
	if v7741 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1747
	}
L1723:
	;
	v7635 = int32(_a_F_PostgresMain_39)
	v7637 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[145]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[145])) = v7637 + int32(1)
	F_pq_startmsgread(m)
	mBase = m.M
	v7642 = m.ExcPending
	if v7642 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1725
	}
L1724:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7712 = m.ExcPending
	if v7712 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1742
	}
L1725:
	;
	v7643 = F_pq_getbyte(m)
	mBase = m.M
	v7644 = m.ExcPending
	if v7644 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1726
	}
L1726:
	;
	v7646 = v7643 - int32(72)
	if base.Ui32(int32(30)) < base.Ui32(v7646) {
		goto L728
	} else {
		goto L1727
	}
L1727:
	;
	if int32(1)<<(uint(v7646)%32)&int32(1207961601) == int32(0) {
		goto L1729
	} else {
		goto L1730
	}
L1728:
	;
	v7662 = F_pq_getmessage(m, v3481+int32(496), v7659)
	mBase = m.M
	v7663 = m.ExcPending
	if v7663 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1733
	}
L1729:
	;
	if v7646 != int32(28) {
		goto L728
	} else {
		goto L1732
	}
L1730:
	;
	goto L1731
L1731:
	;
	v7659 = int32(_a_F_PostgresMain_41)
	goto L1728
L1732:
	;
	v7659 = int32(1073741822)
	goto L1728
L1733:
	;
	if v7662 != 0 {
		goto L729
	} else {
		goto L1734
	}
L1734:
	;
	v7664 = int32(_a_F_PostgresMain_39)
	v7666 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[145]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[145])) = v7666 - int32(1)
	switch v7646 {
	case 0, 11:
		goto L1723
	default:
		goto L1722
	case 28:
		goto L1736
	case 30:
		goto L1735
	}
L1735:
	;
	goto L1724
L1736:
	;
	v7670 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+496))
	v7671 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+500))
	v7672 = int32(_a_F_PostgresMain_166)
	v7673 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50]))
	v7675 = *(*int32)(unsafe.Add(mBase, uint32(v7406)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v7675
	v7677 = *(*int32)(unsafe.Add(mBase, uint32(v7406)+8))
	if base.B2i32(v7677 < int32(1025))|base.B2i32(v7677+v7671 < int32(_a_F_PostgresMain_175)) == int32(0) {
		goto L1737
	} else {
		goto L1738
	}
L1737:
	;
	v7686 = *(*int32)(unsafe.Add(mBase, uint32(v7406)+32))
	v7687 = *(*int32)(unsafe.Add(mBase, uint32(v7406)+4))
	F_json_parse_manifest_incremental_chunk(m, v7686, v7687, v7677-int32(1024), int32(0))
	mBase = m.M
	v7692 = m.ExcPending
	if v7692 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1740
	}
L1738:
	;
	goto L1739
L1739:
	;
	F_appendBinaryStringInfo(m, v7406+int32(4), v7670, v7671)
	mBase = m.M
	v7706 = m.ExcPending
	if v7706 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1741
	}
L1740:
	;
	v7693 = *(*int32)(unsafe.Add(mBase, uint32(v7406)+4))
	v7694 = *(*int32)(unsafe.Add(mBase, uint32(v7406)+8))
	v7696 = int32(1024)
	base.MemoryCopy(m, v7693, v7693+v7694-v7696, int32(1025))
	*(*int32)(unsafe.Add(mBase, uint32(v7406)+8)) = v7696
	goto L1739
L1741:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v7673
	goto L1723
L1742:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v7715 = m.ExcPending
	if v7715 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1743
	}
L1743:
	;
	v7718 = F_pq_getmsgstring(m, v3481+int32(496))
	mBase = m.M
	v7719 = m.ExcPending
	if v7719 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1744
	}
L1744:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+384)) = v7718
	F_errmsg(m, int32(_a_F_PostgresMain_176), v3481+int32(384))
	mBase = m.M
	v7725 = m.ExcPending
	if v7725 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1745
	}
L1745:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(829), int32(_a_F_PostgresMain_177))
	mBase = m.M
	v7730 = m.ExcPending
	if v7730 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1746
	}
L1746:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1747:
	;
	v7742 = *(*int32)(unsafe.Add(mBase, uint32(v7406)+4))
	F_pfree(m, v7742)
	mBase = m.M
	v7744 = m.ExcPending
	if v7744 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1748
	}
L1748:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7406)+4)) = int32(0)
	v7747 = *(*int32)(unsafe.Add(mBase, uint32(v7406)+32))
	v7748 = *(*int32)(unsafe.Add(mBase, uint32(v7747)+68))
	F_pfree(m, v7748)
	mBase = m.M
	v7750 = m.ExcPending
	if v7750 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1749
	}
L1749:
	;
	F_freeJsonLexContext(m, v7747)
	mBase = m.M
	v7752 = m.ExcPending
	if v7752 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1750
	}
L1750:
	;
	F_pfree(m, v7747)
	mBase = m.M
	v7754 = m.ExcPending
	if v7754 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1751
	}
L1751:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v7732
	v7758 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[208]))
	if v7758 != 0 {
		goto L1752
	} else {
		goto L1753
	}
L1752:
	;
	F_MemoryContextDelete(m, v7758)
	mBase = m.M
	v7760 = m.ExcPending
	if v7760 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1755
	}
L1753:
	;
	goto L1754
L1754:
	;
	v7762 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[209]))
	v7766 = *(*int32)(unsafe.Add(mBase, uint32(v7399)+16))
	if v7766 != v7762 {
		goto L1757
	} else {
		goto L1758
	}
L1755:
	;
	goto L1754
L1756:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[208])) = v7399
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[210])) = v7406
	F_ReleaseAuxProcessResources(m, int32(1))
	mBase = m.M
	v7801 = m.ExcPending
	if v7801 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1773
	}
L1757:
	;
	if v7766 == int32(0) {
		goto L1760
	} else {
		goto L1761
	}
L1758:
	;
	goto L1759
L1759:
	;
	goto L1756
L1760:
	;
	if v7762 != 0 {
		goto L1767
	} else {
		goto L1768
	}
L1761:
	;
	v7770 = *(*int32)(unsafe.Add(mBase, uint32(v7399)+28))
	v7771 = *(*int32)(unsafe.Add(mBase, uint32(v7399)+24))
	if v7771 != 0 {
		goto L1763
	} else {
		goto L1764
	}
L1762:
	;
	if v7770 == int32(0) {
		goto L1760
	} else {
		goto L1766
	}
L1763:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7771)+28)) = v7770
	goto L1762
L1764:
	;
	goto L1765
L1765:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7766)+20)) = v7770
	goto L1762
L1766:
	;
	v7776 = *(*int32)(unsafe.Add(mBase, uint32(v7399)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v7770)+24)) = v7776
	goto L1760
L1767:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7399)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7399)+16)) = v7762
	v7783 = *(*int32)(unsafe.Add(mBase, uint32(v7762)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7399)+28)) = v7783
	if v7783 != 0 {
		goto L1770
	} else {
		goto L1771
	}
L1768:
	;
	goto L1769
L1769:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7399)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7399)+16)) = int32(0)
	goto L1759
L1770:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7783)+24)) = v7399
	goto L1772
L1771:
	;
	goto L1772
L1772:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7762)+20)) = v7399
	goto L1756
L1773:
	;
	v8300 = int32(_a_F_PostgresMain_98)
	goto L726
L1774:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7808 = m.ExcPending
	if v7808 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1775
	}
L1775:
	;
	v7809 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+424))
	v7810 = *(*int32)(unsafe.Add(mBase, uint32(v7809)))
	*(*int32)(unsafe.Add(mBase, uint32(v3481))) = v7810
	F_errmsg_internal(m, int32(_a_F_PostgresMain_178), v3481)
	mBase = m.M
	v7814 = m.ExcPending
	if v7814 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1776
	}
L1776:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(2320), int32(_a_F_PostgresMain_61))
	mBase = m.M
	v7819 = m.ExcPending
	if v7819 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1777
	}
L1777:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1778:
	;
	v7825 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+424))
	v7827 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[210]))
	F_SendBaseBackup(m, v7825, v7827)
	mBase = m.M
	v7829 = m.ExcPending
	if v7829 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1779
	}
L1779:
	;
	v8300 = v7820
	goto L726
L1780:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v7836 = m.ExcPending
	if v7836 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1781
	}
L1781:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_179), int32(0))
	mBase = m.M
	v7840 = m.ExcPending
	if v7840 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1782
	}
L1782:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(2115), int32(_a_F_PostgresMain_61))
	mBase = m.M
	v7845 = m.ExcPending
	if v7845 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1783
	}
L1783:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1784:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v7852 = m.ExcPending
	if v7852 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1785
	}
L1785:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+416)) = v4624
	F_errmsg_internal(m, int32(_a_F_PostgresMain_180), v3481+int32(416))
	mBase = m.M
	v7858 = m.ExcPending
	if v7858 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1786
	}
L1786:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(2183), int32(_a_F_PostgresMain_61))
	mBase = m.M
	v7863 = m.ExcPending
	if v7863 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1787
	}
L1787:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1788:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v7870 = m.ExcPending
	if v7870 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1789
	}
L1789:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_181), int32(0))
	mBase = m.M
	v7874 = m.ExcPending
	if v7874 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1790
	}
L1790:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(2209), int32(_a_F_PostgresMain_61))
	mBase = m.M
	v7879 = m.ExcPending
	if v7879 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1791
	}
L1791:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1792:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v7886 = m.ExcPending
	if v7886 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1793
	}
L1793:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+64)) = int32(_a_F_PostgresMain_117)
	F_errmsg(m, int32(_a_F_PostgresMain_182), v3481-int32(-64))
	mBase = m.M
	v7893 = m.ExcPending
	if v7893 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1794
	}
L1794:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(554), int32(_a_F_PostgresMain_183))
	mBase = m.M
	v7898 = m.ExcPending
	if v7898 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1795
	}
L1795:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1796:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v7905 = m.ExcPending
	if v7905 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1797
	}
L1797:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_184), int32(0))
	mBase = m.M
	v7909 = m.ExcPending
	if v7909 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1798
	}
L1798:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1195), int32(_a_F_PostgresMain_120))
	mBase = m.M
	v7914 = m.ExcPending
	if v7914 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1799
	}
L1799:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1800:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v7921 = m.ExcPending
	if v7921 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1801
	}
L1801:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_184), int32(0))
	mBase = m.M
	v7925 = m.ExcPending
	if v7925 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1802
	}
L1802:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1217), int32(_a_F_PostgresMain_120))
	mBase = m.M
	v7930 = m.ExcPending
	if v7930 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1803
	}
L1803:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1804:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v7937 = m.ExcPending
	if v7937 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1805
	}
L1805:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_184), int32(0))
	mBase = m.M
	v7941 = m.ExcPending
	if v7941 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1806
	}
L1806:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1227), int32(_a_F_PostgresMain_120))
	mBase = m.M
	v7946 = m.ExcPending
	if v7946 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1807
	}
L1807:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1808:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v7953 = m.ExcPending
	if v7953 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1809
	}
L1809:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_184), int32(0))
	mBase = m.M
	v7957 = m.ExcPending
	if v7957 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1810
	}
L1810:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1236), int32(_a_F_PostgresMain_120))
	mBase = m.M
	v7962 = m.ExcPending
	if v7962 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1811
	}
L1811:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1812:
	;
	v7967 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+208)) = v7967
	F_errmsg_internal(m, int32(_a_F_PostgresMain_185), v3481+int32(208))
	mBase = m.M
	v7973 = m.ExcPending
	if v7973 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1813
	}
L1813:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1241), int32(_a_F_PostgresMain_120))
	mBase = m.M
	v7978 = m.ExcPending
	if v7978 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1814
	}
L1814:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1815:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+176)) = int32(_a_F_PostgresMain_186)
	F_errmsg(m, int32(_a_F_PostgresMain_187), v3481+int32(176))
	mBase = m.M
	v7989 = m.ExcPending
	if v7989 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1816
	}
L1816:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1326), int32(_a_F_PostgresMain_124))
	mBase = m.M
	v7994 = m.ExcPending
	if v7994 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1817
	}
L1817:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1818:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+160)) = int32(_a_F_PostgresMain_186)
	F_errmsg(m, int32(_a_F_PostgresMain_188), v3481+int32(160))
	mBase = m.M
	v8005 = m.ExcPending
	if v8005 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1819
	}
L1819:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1332), int32(_a_F_PostgresMain_124))
	mBase = m.M
	v8010 = m.ExcPending
	if v8010 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1820
	}
L1820:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1821:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+144)) = int32(_a_F_PostgresMain_186)
	F_errmsg(m, int32(_a_F_PostgresMain_189), v3481+int32(144))
	mBase = m.M
	v8021 = m.ExcPending
	if v8021 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1822
	}
L1822:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1337), int32(_a_F_PostgresMain_124))
	mBase = m.M
	v8026 = m.ExcPending
	if v8026 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1823
	}
L1823:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1824:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+112)) = int32(_a_F_PostgresMain_186)
	F_errmsg(m, int32(_a_F_PostgresMain_190), v3481+int32(112))
	mBase = m.M
	v8037 = m.ExcPending
	if v8037 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1825
	}
L1825:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1343), int32(_a_F_PostgresMain_124))
	mBase = m.M
	v8042 = m.ExcPending
	if v8042 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1826
	}
L1826:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1827:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+128)) = int32(_a_F_PostgresMain_186)
	F_errmsg(m, int32(_a_F_PostgresMain_191), v3481+int32(128))
	mBase = m.M
	v8053 = m.ExcPending
	if v8053 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1828
	}
L1828:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1349), int32(_a_F_PostgresMain_124))
	mBase = m.M
	v8058 = m.ExcPending
	if v8058 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1829
	}
L1829:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1830:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v8065 = m.ExcPending
	if v8065 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1831
	}
L1831:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_184), int32(0))
	mBase = m.M
	v8069 = m.ExcPending
	if v8069 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1832
	}
L1832:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1489), int32(_a_F_PostgresMain_192))
	mBase = m.M
	v8074 = m.ExcPending
	if v8074 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1833
	}
L1833:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1834:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v8081 = m.ExcPending
	if v8081 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1835
	}
L1835:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_184), int32(0))
	mBase = m.M
	v8085 = m.ExcPending
	if v8085 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1836
	}
L1836:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1498), int32(_a_F_PostgresMain_192))
	mBase = m.M
	v8090 = m.ExcPending
	if v8090 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1837
	}
L1837:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1838:
	;
	v8095 = *(*int32)(unsafe.Add(mBase, uint32(v6031)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+224)) = v8095
	F_errmsg_internal(m, int32(_a_F_PostgresMain_185), v3481+int32(224))
	mBase = m.M
	v8101 = m.ExcPending
	if v8101 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1839
	}
L1839:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(1503), int32(_a_F_PostgresMain_192))
	mBase = m.M
	v8106 = m.ExcPending
	if v8106 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1840
	}
L1840:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1841:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1842:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v8115 = m.ExcPending
	if v8115 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1843
	}
L1843:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+240)) = v3481 + int32(_a_F_PostgresMain_101)
	F_errmsg(m, int32(_a_F_PostgresMain_193), v3481+int32(240))
	mBase = m.M
	v8123 = m.ExcPending
	if v8123 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1844
	}
L1844:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(651), int32(_a_F_PostgresMain_194))
	mBase = m.M
	v8128 = m.ExcPending
	if v8128 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1845
	}
L1845:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1846:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v8134 = m.ExcPending
	if v8134 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1847
	}
L1847:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+256)) = v3481 + int32(_a_F_PostgresMain_101)
	F_errmsg(m, int32(_a_F_PostgresMain_195), v3481+int32(256))
	mBase = m.M
	v8142 = m.ExcPending
	if v8142 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1848
	}
L1848:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(658), int32(_a_F_PostgresMain_194))
	mBase = m.M
	v8147 = m.ExcPending
	if v8147 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1849
	}
L1849:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1850:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v8153 = m.ExcPending
	if v8153 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1851
	}
L1851:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+320)) = v3481 + int32(_a_F_PostgresMain_101)
	F_errmsg(m, int32(_a_F_PostgresMain_196), v3481+int32(320))
	mBase = m.M
	v8161 = m.ExcPending
	if v8161 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1852
	}
L1852:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(662), int32(_a_F_PostgresMain_194))
	mBase = m.M
	v8166 = m.ExcPending
	if v8166 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1853
	}
L1853:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1854:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v8172 = m.ExcPending
	if v8172 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1855
	}
L1855:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+288)) = v3481 + int32(_a_F_PostgresMain_101)
	F_errmsg(m, int32(_a_F_PostgresMain_197), v3481+int32(288))
	mBase = m.M
	v8180 = m.ExcPending
	if v8180 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1856
	}
L1856:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(679), int32(_a_F_PostgresMain_194))
	mBase = m.M
	v8185 = m.ExcPending
	if v8185 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1857
	}
L1857:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1858:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v8192 = m.ExcPending
	if v8192 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1859
	}
L1859:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v3481)+312)) = uint32(v7310)
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+308)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+304)) = v3481 + int32(_a_F_PostgresMain_101)
	F_errmsg(m, int32(_a_F_PostgresMain_198), v3481+int32(304))
	mBase = m.M
	v8203 = m.ExcPending
	if v8203 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1860
	}
L1860:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(684), int32(_a_F_PostgresMain_194))
	mBase = m.M
	v8208 = m.ExcPending
	if v8208 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1861
	}
L1861:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1862:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v8214 = m.ExcPending
	if v8214 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1863
	}
L1863:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+272)) = v3481 + int32(_a_F_PostgresMain_101)
	F_errmsg(m, int32(_a_F_PostgresMain_199), v3481+int32(272))
	mBase = m.M
	v8222 = m.ExcPending
	if v8222 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1864
	}
L1864:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(693), int32(_a_F_PostgresMain_194))
	mBase = m.M
	v8227 = m.ExcPending
	if v8227 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1865
	}
L1865:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1866:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v8234 = m.ExcPending
	if v8234 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1867
	}
L1867:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_42), int32(0))
	mBase = m.M
	v8238 = m.ExcPending
	if v8238 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1868
	}
L1868:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(807), int32(_a_F_PostgresMain_177))
	mBase = m.M
	v8243 = m.ExcPending
	if v8243 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1869
	}
L1869:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1870:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8250 = m.ExcPending
	if v8250 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1873
	}
L1871:
	;
	goto L1872
L1872:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8266 = m.ExcPending
	if v8266 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1877
	}
L1873:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v8253 = m.ExcPending
	if v8253 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1874
	}
L1874:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_42), int32(0))
	mBase = m.M
	v8257 = m.ExcPending
	if v8257 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1875
	}
L1875:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(781), int32(_a_F_PostgresMain_177))
	mBase = m.M
	v8262 = m.ExcPending
	if v8262 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1876
	}
L1876:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1877:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v8269 = m.ExcPending
	if v8269 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1878
	}
L1878:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3481)+368)) = v7643
	F_errmsg(m, int32(_a_F_PostgresMain_200), v3481+int32(368))
	mBase = m.M
	v8275 = m.ExcPending
	if v8275 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1879
	}
L1879:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_60), int32(798), int32(_a_F_PostgresMain_177))
	mBase = m.M
	v8280 = m.ExcPending
	if v8280 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1880
	}
L1880:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1881:
	;
	v8284 = *(*int32)(unsafe.Add(mBase, uint32(v3481)+424))
	F_StartTransactionCommand(m)
	mBase = m.M
	v8287 = m.ExcPending
	if v8287 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1882
	}
L1882:
	;
	v8288 = *(*int32)(unsafe.Add(mBase, uint32(v8284)+4))
	F_GetPGVariable(m, v8288, v8282)
	mBase = m.M
	v8290 = m.ExcPending
	if v8290 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1883
	}
L1883:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v8292 = m.ExcPending
	if v8292 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1884
	}
L1884:
	;
	v8300 = int32(_a_F_PostgresMain_201)
	goto L726
L1885:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v3486
	v8337 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[164]))
	F_MemoryContextReset(m, v8337)
	mBase = m.M
	v8339 = m.ExcPending
	if v8339 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1886
	}
L1886:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[110])) = int32(0)
	goto L725
L1887:
	;
	goto L717
L1888:
	;
	v8440 = int32(_a_F_PostgresMain_202)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[211])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[212])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[213])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[214])) = int64(1)
	v8450 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1891
L1889:
	;
	goto L1890
L1890:
	;
	F_start_xact_command(m)
	mBase = m.M
	v8454 = m.ExcPending
	if v8454 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1892
	}
L1891:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMain_203))
	mBase = m.M
	goto L1890
L1892:
	;
	v8456 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[215]))
	if v8456 != 0 {
		goto L1893
	} else {
		goto L1894
	}
L1893:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[215])) = int32(0)
	F_DropCachedPlan(m, v8456)
	mBase = m.M
	v8461 = m.ExcPending
	if v8461 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1896
	}
L1894:
	;
	goto L1895
L1895:
	;
	v8462 = int32(_a_F_PostgresMain_166)
	v8463 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50]))
	v8466 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[49]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v8466
	v8469 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[216])))
	if v8469 == int32(1) {
		goto L1897
	} else {
		goto L1898
	}
L1896:
	;
	goto L1895
L1897:
	;
	v8472 = int32(_a_F_PostgresMain_202)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[211])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[212])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[213])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[214])) = int64(1)
	v8482 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1900
L1898:
	;
	goto L1899
L1899:
	;
	v8486 = F_raw_parser(m, v3461, int32(0))
	mBase = m.M
	v8487 = m.ExcPending
	if v8487 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1901
	}
L1900:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMain_203))
	mBase = m.M
	goto L1899
L1901:
	;
	v8489 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[216])))
	if v8489 == int32(1) {
		goto L1902
	} else {
		goto L1903
	}
L1902:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMain_204))
	mBase = m.M
	v8494 = m.ExcPending
	if v8494 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1905
	}
L1903:
	;
	goto L1904
L1904:
	;
	v8496 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[217])))
	if v8496 == int32(1) {
		goto L1906
	} else {
		goto L1907
	}
L1905:
	;
	goto L1904
L1906:
	;
	v8501 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[218])))
	F_elog_node_display(m, int32(_a_F_PostgresMain_205), v8486, v8501)
	mBase = m.M
	v8503 = m.ExcPending
	if v8503 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1909
	}
L1907:
	;
	goto L1908
L1908:
	;
	v8505 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[219]))
	switch v8505 {
	case 0:
		v8790 = v8424
		goto L1914
	default:
		goto L1919
	case 3:
		goto L1918
	}
L1909:
	;
	goto L1908
L1910:
	;
	v9438 = F_check_log_duration(m, v8428+int32(80), v9409)
	mBase = m.M
	v9439 = m.ExcPending
	if v9439 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2091
	}
L1911:
	;
	v9381 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[63])))
	goto L2080
L1912:
	;
	v9325 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[63])))
	goto L2070
L1913:
	;
	v8860 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+4))
	if v8860 <= int32(0) {
		goto L1911
	} else {
		goto L1946
	}
L1914:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v8463
	if v8486 == int32(0) {
		v9296 = v8790
		goto L1912
	} else {
		goto L1945
	}
L1915:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(1198), int32(_a_F_PostgresMain_206))
	mBase = m.M
	v8776 = m.ExcPending
	if v8776 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1944
	}
L1916:
	;
	v8725 = *(*int32)(unsafe.Add(mBase, uint32(v8715)+64))
	v8726 = *(*int32)(unsafe.Add(mBase, uint32(v8725)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8428)+48)) = v8726
	v8731 = F_errdetail(m, int32(_a_F_PostgresMain_207), v8428+int32(48))
	mBase = m.M
	v8732 = m.ExcPending
	if v8732 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1943
	}
L1917:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v8463
	v9296 = v8424
	goto L1912
L1918:
	;
	v8644 = int32(1)
	v8647 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8648 = m.ExcPending
	if v8648 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1929
	}
L1919:
	;
	if v8486 == int32(0) {
		goto L1917
	} else {
		goto L1920
	}
L1920:
	;
	v8508 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+4))
	if int32(0) < v8508 {
		goto L1921
	} else {
		goto L1922
	}
L1921:
	;
	v8513 = v8424
	goto L1924
L1922:
	;
	v8583 = v8508
	goto L1923
L1923:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v8463
	v8834 = v8424
	v8840 = v8583
	goto L1913
L1924:
	;
	v8550 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+12))
	v8554 = *(*int32)(unsafe.Add(mBase, uint32(v8550+v8513<<(uint(int32(2))%32))))
	v8555 = F_GetCommandLogLevel(m, v8554)
	mBase = m.M
	v8556 = m.ExcPending
	if v8556 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1926
	}
L1925:
	;
	v8583 = v8562
	goto L1923
L1926:
	;
	v8558 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[219]))
	if base.Ui32(v8555) <= base.Ui32(v8558) {
		goto L1918
	} else {
		goto L1927
	}
L1927:
	;
	v8561 = v8513 + int32(1)
	v8562 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+4))
	if v8561 < v8562 {
		v8513 = v8561
		goto L1924
	} else {
		goto L1928
	}
L1928:
	;
	goto L1925
L1929:
	;
	if v8647 == int32(0) {
		v8790 = v8644
		goto L1914
	} else {
		goto L1930
	}
L1930:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8428)+64)) = v3461
	F_errmsg(m, int32(_a_F_PostgresMain_208), v8428-int32(-64))
	mBase = m.M
	v8656 = m.ExcPending
	if v8656 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1931
	}
L1931:
	;
	F_errhidestmt(m)
	mBase = m.M
	v8658 = m.ExcPending
	if v8658 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1932
	}
L1932:
	;
	if v8486 == int32(0) {
		goto L1915
	} else {
		goto L1933
	}
L1933:
	;
	v8661 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+4))
	if v8661 <= int32(0) {
		goto L1915
	} else {
		goto L1934
	}
L1934:
	;
	v8667 = int32(0)
	v8672 = v8661
	goto L1935
L1935:
	;
	v8704 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+12))
	v8708 = *(*int32)(unsafe.Add(mBase, uint32(v8704+v8667<<(uint(int32(2))%32))))
	v8709 = *(*int32)(unsafe.Add(mBase, uint32(v8708)+4))
	v8710 = *(*int32)(unsafe.Add(mBase, uint32(v8709)))
	if v8710 == int32(253) {
		goto L1937
	} else {
		goto L1938
	}
L1936:
	;
	goto L1915
L1937:
	;
	v8713 = *(*int32)(unsafe.Add(mBase, uint32(v8709)+4))
	v8715 = F_FetchPreparedStatement(m, v8713, int32(0))
	mBase = m.M
	v8716 = m.ExcPending
	if v8716 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1940
	}
L1938:
	;
	v8719 = v8672
	goto L1939
L1939:
	;
	v8721 = v8667 + int32(1)
	if v8721 < v8719 {
		v8667 = v8721
		v8672 = v8719
		goto L1935
	} else {
		goto L1942
	}
L1940:
	;
	if v8715 != 0 {
		goto L1916
	} else {
		goto L1941
	}
L1941:
	;
	v8717 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+4))
	v8719 = v8717
	goto L1939
L1942:
	;
	goto L1936
L1943:
	;
	goto L1915
L1944:
	;
	v8790 = v8644
	goto L1914
L1945:
	;
	v8820 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+4))
	v8834 = v8790
	v8840 = v8820
	goto L1913
L1946:
	;
	v8890 = v2541
	goto L1947
L1947:
	;
	v8902 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+12))
	v8905 = v8902 + v8890<<(uint(int32(2))%32)
	v8906 = *(*int32)(unsafe.Add(mBase, uint32(v8905)))
	v8911 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[220]))
	if v8911 == int32(0) {
		goto L1950
	} else {
		goto L1951
	}
L1948:
	;
	goto L1911
L1949:
	;
	v8959 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[220]))
	if v8959 == int32(0) {
		goto L1955
	} else {
		goto L1956
	}
L1950:
	;
	goto L1949
L1951:
	;
	v8915 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[221])))
	if v8915&int32(1) == int32(0) {
		goto L1950
	} else {
		goto L1952
	}
L1952:
	;
	v8922 = *(*int64)(unsafe.Add(mBase, uint32(v8911)+392))
	if int32(0)&base.B2i32(v8922 != int64(0)) != 0 {
		goto L1950
	} else {
		goto L1953
	}
L1953:
	;
	v8926 = int32(_a_F_PostgresMain_209)
	v8928 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222]))
	v8929 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222])) = v8928 + v8929
	v8932 = *(*int32)(unsafe.Add(mBase, uint32(v8911)))
	*(*int32)(unsafe.Add(mBase, uint32(v8911))) = v8932 + v8929
	v8936 = int32(0)
	v8938 = int32(_a_F_PostgresMain_210)
	v8939 = base.AtomicRmwOr32(m, v8936, v8938, v8936)
	*(*int64)(unsafe.Add(mBase, uint32(v8911)+392)) = int64(0)
	v8944 = base.AtomicRmwOr32(m, v8936, v8938, v8936)
	v8945 = *(*int32)(unsafe.Add(mBase, uint32(v8911)))
	*(*int32)(unsafe.Add(mBase, uint32(v8911))) = v8945 + v8929
	v8951 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222])) = v8951 - v8929
	goto L1950
L1954:
	;
	v9003 = *(*int32)(unsafe.Add(mBase, uint32(v8906)+4))
	v9004 = F_CreateCommandTag(m, v9003)
	mBase = m.M
	v9005 = m.ExcPending
	if v9005 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1959
	}
L1955:
	;
	goto L1954
L1956:
	;
	v8963 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[221])))
	if v8963&int32(1) == int32(0) {
		goto L1955
	} else {
		goto L1957
	}
L1957:
	;
	v8970 = *(*int64)(unsafe.Add(mBase, uint32(v8959)+400))
	if int32(0)&base.B2i32(v8970 != int64(0)) != 0 {
		goto L1955
	} else {
		goto L1958
	}
L1958:
	;
	v8974 = int32(_a_F_PostgresMain_209)
	v8976 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222]))
	v8977 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222])) = v8976 + v8977
	v8980 = *(*int32)(unsafe.Add(mBase, uint32(v8959)))
	*(*int32)(unsafe.Add(mBase, uint32(v8959))) = v8980 + v8977
	v8984 = int32(0)
	v8986 = int32(_a_F_PostgresMain_210)
	v8987 = base.AtomicRmwOr32(m, v8984, v8986, v8984)
	*(*int64)(unsafe.Add(mBase, uint32(v8959)+400)) = int64(0)
	v8992 = base.AtomicRmwOr32(m, v8984, v8986, v8984)
	v8993 = *(*int32)(unsafe.Add(mBase, uint32(v8959)))
	*(*int32)(unsafe.Add(mBase, uint32(v8959))) = v8993 + v8977
	v8999 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222])) = v8999 - v8977
	goto L1955
L1959:
	;
	v9010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9004<<(uint(int32(3))%32))+uint32(_c_F_PostgresMain[223]))))
	*(*int32)(unsafe.Add(mBase, uint32(v8428+int32(72)))) = v9010
	goto L1960
L1960:
	;
	v9015 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v9016 = *(*int32)(unsafe.Add(mBase, uint32(v9015)+24))
	goto L1962
L1961:
	;
	F_start_xact_command(m)
	mBase = m.M
	v9056 = m.ExcPending
	if v9056 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1972
	}
L1962:
	;
	if base.B2i32((v9016-int32(7))&int32(-9) == int32(0)) == int32(0) {
		goto L1961
	} else {
		goto L1963
	}
L1963:
	;
	v9025 = *(*int32)(unsafe.Add(mBase, uint32(v8906)+4))
	if v9025 == int32(0) {
		goto L1964
	} else {
		goto L1965
	}
L1964:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9041 = m.ExcPending
	if v9041 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1968
	}
L1965:
	;
	v9028 = *(*int32)(unsafe.Add(mBase, uint32(v9025)))
	if v9028 != int32(225) {
		goto L1964
	} else {
		goto L1966
	}
L1966:
	;
	v9031 = *(*int32)(unsafe.Add(mBase, uint32(v9025)+4))
	if (v9031-int32(2))&int32(-6) == int32(0) {
		goto L1961
	} else {
		goto L1967
	}
L1967:
	;
	goto L1964
L1968:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v9044 = m.ExcPending
	if v9044 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1969
	}
L1969:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_181), int32(0))
	mBase = m.M
	v9048 = m.ExcPending
	if v9048 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1970
	}
L1970:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(1264), int32(_a_F_PostgresMain_206))
	mBase = m.M
	v9053 = m.ExcPending
	if v9053 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1971
	}
L1971:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1972:
	;
	v9058 = base.B2i32(v8840 < int32(2))
	if v9058 == int32(0) {
		goto L1973
	} else {
		goto L1974
	}
L1973:
	;
	v9063 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v9064 = *(*int32)(unsafe.Add(mBase, uint32(v9063)+24))
	if v9064 == int32(1) {
		goto L1977
	} else {
		goto L1978
	}
L1974:
	;
	goto L1975
L1975:
	;
	v9070 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[148]))
	if v9070 != 0 {
		goto L1980
	} else {
		goto L1981
	}
L1976:
	;
	goto L1975
L1977:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9063)+24)) = int32(4)
	goto L1979
L1978:
	;
	goto L1979
L1979:
	;
	goto L1976
L1980:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v9072 = m.ExcPending
	if v9072 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1983
	}
L1981:
	;
	goto L1982
L1982:
	;
	v9075 = *(*int32)(unsafe.Add(mBase, uint32(v8906)+4))
	v9076 = *(*int32)(unsafe.Add(mBase, uint32(v9075)))
	switch v9076 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v9080 = int32(1)
		goto L1985
	default:
		goto L1986
	}
L1983:
	;
	goto L1982
L1984:
	;
	if v9080 != 0 {
		goto L1987
	} else {
		goto L1988
	}
L1985:
	;
	goto L1984
L1986:
	;
	v9080 = int32(0)
	goto L1985
L1987:
	;
	v9081 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v9082 = m.ExcPending
	if v9082 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1990
	}
L1988:
	;
	goto L1989
L1989:
	;
	v9087 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[49]))
	v9089 = v8905 + int32(4)
	v9090 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+12))
	v9091 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+4))
	if base.Ui32(v9089) < base.Ui32(v9090+v9091<<(uint(int32(2))%32)) {
		goto L1992
	} else {
		goto L1993
	}
L1990:
	;
	F_PushActiveSnapshot(m, v9081)
	mBase = m.M
	v9084 = m.ExcPending
	if v9084 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1991
	}
L1991:
	;
	goto L1989
L1992:
	;
	v9100 = F_AllocSetContextCreateInternal(m, v9087, int32(_a_F_PostgresMain_211), int32(0), int32(_a_F_PostgresMain_16), int32(_a_F_PostgresMain_17))
	mBase = m.M
	v9101 = m.ExcPending
	if v9101 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L1995
	}
L1993:
	;
	v9102 = v9087
	v9103 = int32(0)
	goto L1994
L1994:
	;
	v9104 = int32(_a_F_PostgresMain_166)
	v9105 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v9102
	v9109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[216])))
	if v9109 == int32(1) {
		goto L1996
	} else {
		goto L1997
	}
L1995:
	;
	v9102 = v9100
	v9103 = v9100
	goto L1994
L1996:
	;
	v9112 = int32(_a_F_PostgresMain_202)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[211])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[212])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[213])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[214])) = int64(1)
	v9122 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1999
L1997:
	;
	goto L1998
L1998:
	;
	v9125 = int32(0)
	v9128 = F_parse_analyze_fixedparams(m, v8906, v3461, v9125, v9125, v9125)
	mBase = m.M
	v9129 = m.ExcPending
	if v9129 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2000
	}
L1999:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMain_203))
	mBase = m.M
	goto L1998
L2000:
	;
	v9131 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[216])))
	if v9131 == int32(1) {
		goto L2001
	} else {
		goto L2002
	}
L2001:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMain_212))
	mBase = m.M
	v9136 = m.ExcPending
	if v9136 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2004
	}
L2002:
	;
	goto L2003
L2003:
	;
	v9137 = F_pg_rewrite_query(m, v9128)
	mBase = m.M
	v9138 = m.ExcPending
	if v9138 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2005
	}
L2004:
	;
	goto L2003
L2005:
	;
	v9141 = F_pg_plan_queries(m, v9137, v3461, int32(2048), int32(0))
	mBase = m.M
	v9142 = m.ExcPending
	if v9142 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2006
	}
L2006:
	;
	if v9080 != 0 {
		goto L2007
	} else {
		goto L2008
	}
L2007:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v9144 = m.ExcPending
	if v9144 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2010
	}
L2008:
	;
	goto L2009
L2009:
	;
	v9146 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[148]))
	if v9146 != 0 {
		goto L2011
	} else {
		goto L2012
	}
L2010:
	;
	goto L2009
L2011:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v9148 = m.ExcPending
	if v9148 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2014
	}
L2012:
	;
	goto L2013
L2013:
	;
	v9150 = int32(1)
	v9152 = F_CreatePortal(m, int32(_a_F_PostgresMain_213), v9150, v9150)
	mBase = m.M
	v9153 = m.ExcPending
	if v9153 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2015
	}
L2014:
	;
	goto L2013
L2015:
	;
	v9154 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9152)+136)) = uint8(v9154)
	*(*int32)(unsafe.Add(mBase, uint32(v9152)+80)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9152)+60)) = v9154
	*(*int32)(unsafe.Add(mBase, uint32(v9152)+56)) = v9141
	*(*int64)(unsafe.Add(mBase, uint32(v9152)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9152)+40)) = v9004
	*(*int32)(unsafe.Add(mBase, uint32(v9152)+36)) = v9004
	*(*int32)(unsafe.Add(mBase, uint32(v9152)+32)) = v3461
	*(*int32)(unsafe.Add(mBase, uint32(v9152)+4)) = v9154
	goto L2016
L2016:
	;
	v9168 = int32(0)
	F_PortalStart(m, v9152, v9168, v9168, v9168)
	mBase = m.M
	v9172 = m.ExcPending
	if v9172 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2017
	}
L2017:
	;
	v9173 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v8428)+78)) = uint16(v9173)
	v9175 = *(*int32)(unsafe.Add(mBase, uint32(v8906)+4))
	v9176 = *(*int32)(unsafe.Add(mBase, uint32(v9175)))
	if v9176 != int32(203) {
		goto L2018
	} else {
		goto L2019
	}
L2018:
	;
	F_PortalSetResultFormat(m, v9152, int32(1), v8428+int32(78))
	mBase = m.M
	v9197 = m.ExcPending
	if v9197 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2024
	}
L2019:
	;
	v9179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9175)+16)))
	if v9179 != 0 {
		goto L2018
	} else {
		goto L2020
	}
L2020:
	;
	v9180 = *(*int32)(unsafe.Add(mBase, uint32(v9175)+12))
	v9181 = F_GetPortalByName(m, v9180)
	mBase = m.M
	v9182 = m.ExcPending
	if v9182 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2021
	}
L2021:
	;
	if v9181 == int32(0) {
		goto L2018
	} else {
		goto L2022
	}
L2022:
	;
	v9185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9181)+76)))
	if v9185&int32(1) == int32(0) {
		goto L2018
	} else {
		goto L2023
	}
L2023:
	;
	v9190 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v8428)+78)) = uint16(v9190)
	goto L2018
L2024:
	;
	v9198 = F_CreateDestReceiver(m, v8433)
	mBase = m.M
	v9199 = m.ExcPending
	if v9199 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2025
	}
L2025:
	;
	if v8433 == int32(2) {
		goto L2026
	} else {
		goto L2027
	}
L2026:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9198)+20)) = v9152
	goto L2028
L2027:
	;
	goto L2028
L2028:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v9105
	v9209 = F_PortalRun(m, v9152, int32(2147483647), int32(1), v9198, v9198, v8428+int32(80))
	mBase = m.M
	v9210 = m.ExcPending
	if v9210 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2029
	}
L2029:
	;
	v9211 = *(*int32)(unsafe.Add(mBase, uint32(v9198)+12))
	m.T0[v9211].(func(*base.Module, int32))(m, v9198)
	mBase = m.M
	v9213 = m.ExcPending
	if v9213 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2030
	}
L2030:
	;
	F_PortalDrop(m, v9152, int32(0))
	mBase = m.M
	v9216 = m.ExcPending
	if v9216 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2031
	}
L2031:
	;
	v9217 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+12))
	v9218 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+4))
	if base.Ui32(v9217+v9218<<(uint(int32(2))%32)) <= base.Ui32(v9089) {
		goto L2034
	} else {
		goto L2035
	}
L2032:
	;
	F_EndCommand(m, v8428+int32(80), v8433)
	mBase = m.M
	v9276 = m.ExcPending
	if v9276 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2064
	}
L2033:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v9269 = m.ExcPending
	if v9269 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2063
	}
L2034:
	;
	if v9058 == int32(0) {
		goto L2037
	} else {
		goto L2038
	}
L2035:
	;
	goto L2036
L2036:
	;
	v9244 = *(*int32)(unsafe.Add(mBase, uint32(v8906)+4))
	v9245 = *(*int32)(unsafe.Add(mBase, uint32(v9244)))
	if v9245 == int32(225) {
		goto L2050
	} else {
		goto L2051
	}
L2037:
	;
	v9227 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v9228 = *(*int32)(unsafe.Add(mBase, uint32(v9227)+24))
	if v9228 == int32(4) {
		goto L2041
	} else {
		goto L2042
	}
L2038:
	;
	goto L2039
L2039:
	;
	v9236 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[63])))
	goto L2044
L2040:
	;
	goto L2039
L2041:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9227)+24)) = int32(1)
	goto L2043
L2042:
	;
	goto L2043
L2043:
	;
	goto L2040
L2044:
	;
	if v9236 != 0 {
		goto L2045
	} else {
		goto L2046
	}
L2045:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v9239 = m.ExcPending
	if v9239 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2048
	}
L2046:
	;
	goto L2047
L2047:
	;
	v9241 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])))
	if v9241 == int32(0) {
		goto L2032
	} else {
		goto L2049
	}
L2048:
	;
	goto L2047
L2049:
	;
	goto L2033
L2050:
	;
	v9251 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[63])))
	goto L2053
L2051:
	;
	goto L2052
L2052:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v9258 = m.ExcPending
	if v9258 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2059
	}
L2053:
	;
	if v9251 != 0 {
		goto L2054
	} else {
		goto L2055
	}
L2054:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v9254 = m.ExcPending
	if v9254 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2057
	}
L2055:
	;
	goto L2056
L2056:
	;
	v9256 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])))
	if v9256 != 0 {
		goto L2033
	} else {
		goto L2058
	}
L2057:
	;
	goto L2056
L2058:
	;
	goto L2032
L2059:
	;
	v9262 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[63])))
	goto L2060
L2060:
	;
	if v9262 == int32(0) {
		goto L2032
	} else {
		goto L2061
	}
L2061:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v9267 = m.ExcPending
	if v9267 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2062
	}
L2062:
	;
	goto L2032
L2063:
	;
	v9271 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])) = uint8(v9271)
	goto L2032
L2064:
	;
	if v9103 != 0 {
		goto L2065
	} else {
		goto L2066
	}
L2065:
	;
	F_MemoryContextDelete(m, v9103)
	mBase = m.M
	v9278 = m.ExcPending
	if v9278 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2068
	}
L2066:
	;
	goto L2067
L2067:
	;
	v9280 = v8890 + int32(1)
	v9281 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+4))
	if v9280 < v9281 {
		v8890 = v9280
		goto L1947
	} else {
		goto L2069
	}
L2068:
	;
	goto L2067
L2069:
	;
	goto L1948
L2070:
	;
	if v9325 != 0 {
		goto L2071
	} else {
		goto L2072
	}
L2071:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v9328 = m.ExcPending
	if v9328 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2074
	}
L2072:
	;
	goto L2073
L2073:
	;
	v9330 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])))
	if v9330 != 0 {
		goto L2075
	} else {
		goto L2076
	}
L2074:
	;
	goto L2073
L2075:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v9332 = m.ExcPending
	if v9332 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2078
	}
L2076:
	;
	goto L2077
L2077:
	;
	F_NullCommand(m, v8433)
	mBase = m.M
	v9337 = m.ExcPending
	if v9337 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2079
	}
L2078:
	;
	v9334 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])) = uint8(v9334)
	goto L2077
L2079:
	;
	v9409 = v9296
	v9435 = int32(1)
	goto L1910
L2080:
	;
	if v9381 != 0 {
		goto L2081
	} else {
		goto L2082
	}
L2081:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v9384 = m.ExcPending
	if v9384 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2084
	}
L2082:
	;
	goto L2083
L2083:
	;
	v9385 = int32(0)
	v9387 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])))
	if v9387 == v9385 {
		v9409 = v8834
		v9435 = v9385
		goto L1910
	} else {
		goto L2085
	}
L2084:
	;
	goto L2083
L2085:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v9391 = m.ExcPending
	if v9391 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2086
	}
L2086:
	;
	v9393 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])) = uint8(v9393)
	v9409 = v8834
	v9435 = v9393
	goto L1910
L2087:
	;
	if v8435 != 0 {
		goto L2113
	} else {
		goto L2114
	}
L2088:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), v9568, int32(_a_F_PostgresMain_206))
	mBase = m.M
	v9587 = m.ExcPending
	if v9587 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2112
	}
L2089:
	;
	v9459 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9460 = m.ExcPending
	if v9460 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2096
	}
L2090:
	;
	v9444 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9445 = m.ExcPending
	if v9445 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2092
	}
L2091:
	;
	switch v9438 - int32(1) {
	case 0:
		goto L2090
	case 1:
		goto L2089
	default:
		goto L2087
	}
L2092:
	;
	if v9444 == int32(0) {
		goto L2087
	} else {
		goto L2093
	}
L2093:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8428))) = v8428 + int32(80)
	F_errmsg(m, int32(_a_F_PostgresMain_214), v8428)
	mBase = m.M
	v9453 = m.ExcPending
	if v9453 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2094
	}
L2094:
	;
	F_errhidestmt(m)
	mBase = m.M
	v9455 = m.ExcPending
	if v9455 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2095
	}
L2095:
	;
	v9568 = int32(1489)
	goto L2088
L2096:
	;
	if v9459 == int32(0) {
		goto L2087
	} else {
		goto L2097
	}
L2097:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8428)+36)) = v3461
	*(*int32)(unsafe.Add(mBase, uint32(v8428)+32)) = v8428 + int32(80)
	F_errmsg(m, int32(_a_F_PostgresMain_215), v8428+int32(32))
	mBase = m.M
	v9471 = m.ExcPending
	if v9471 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2098
	}
L2098:
	;
	F_errhidestmt(m)
	mBase = m.M
	v9473 = m.ExcPending
	if v9473 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2099
	}
L2099:
	;
	v9474 = int32(1496)
	if v9435 != 0 {
		v9568 = v9474
		goto L2088
	} else {
		goto L2100
	}
L2100:
	;
	v9475 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+4))
	if v9475 <= int32(0) {
		v9568 = v9474
		goto L2088
	} else {
		goto L2101
	}
L2101:
	;
	v9481 = int32(0)
	v9486 = v9475
	goto L2102
L2102:
	;
	v9518 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+12))
	v9522 = *(*int32)(unsafe.Add(mBase, uint32(v9518+v9481<<(uint(int32(2))%32))))
	v9523 = *(*int32)(unsafe.Add(mBase, uint32(v9522)+4))
	v9524 = *(*int32)(unsafe.Add(mBase, uint32(v9523)))
	if v9524 == int32(253) {
		goto L2105
	} else {
		goto L2106
	}
L2103:
	;
	v9537 = *(*int32)(unsafe.Add(mBase, uint32(v9529)+64))
	v9538 = *(*int32)(unsafe.Add(mBase, uint32(v9537)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8428)+16)) = v9538
	v9543 = F_errdetail(m, int32(_a_F_PostgresMain_207), v8428+int32(16))
	mBase = m.M
	v9544 = m.ExcPending
	if v9544 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2111
	}
L2104:
	;
	goto L2103
L2105:
	;
	v9527 = *(*int32)(unsafe.Add(mBase, uint32(v9523)+4))
	v9529 = F_FetchPreparedStatement(m, v9527, int32(0))
	mBase = m.M
	v9530 = m.ExcPending
	if v9530 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2108
	}
L2106:
	;
	v9533 = v9486
	goto L2107
L2107:
	;
	v9535 = v9481 + int32(1)
	if v9535 < v9533 {
		v9481 = v9535
		v9486 = v9533
		goto L2102
	} else {
		goto L2110
	}
L2108:
	;
	if v9529 != 0 {
		goto L2104
	} else {
		goto L2109
	}
L2109:
	;
	v9531 = *(*int32)(unsafe.Add(mBase, uint32(v8486)+4))
	v9533 = v9531
	goto L2107
L2110:
	;
	v9568 = v9474
	goto L2088
L2111:
	;
	v9568 = v9474
	goto L2088
L2112:
	;
	goto L2087
L2113:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMain_216))
	mBase = m.M
	v9629 = m.ExcPending
	if v9629 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2116
	}
L2114:
	;
	goto L2115
L2115:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[110])) = int32(0)
	m.G0 = v8428 + int32(112)
	goto L714
L2116:
	;
	goto L2115
L2117:
	;
	v9684 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[158]))
	if v9684 < int32(0) {
		goto L2119
	} else {
		goto L2120
	}
L2118:
	;
	v9691 = v2548 + int32(440)
	v9692 = F_pq_getmsgstring(m, v9691)
	mBase = m.M
	v9693 = m.ExcPending
	if v9693 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2122
	}
L2119:
	;
	v9688 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[159])) = v9688
	goto L2121
L2120:
	;
	goto L2121
L2121:
	;
	goto L2118
L2122:
	;
	v9694 = F_pq_getmsgstring(m, v9691)
	mBase = m.M
	v9695 = m.ExcPending
	if v9695 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2123
	}
L2123:
	;
	v9697 = F_pq_getmsgint(m, v9691, int32(2))
	mBase = m.M
	v9698 = m.ExcPending
	if v9698 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2124
	}
L2124:
	;
	if int32(0) < v9697 {
		goto L2125
	} else {
		goto L2126
	}
L2125:
	;
	v9703 = F_palloc_mul(m, int32(4), v9697)
	mBase = m.M
	v9704 = m.ExcPending
	if v9704 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2128
	}
L2126:
	;
	v9773 = int32(0)
	goto L2127
L2127:
	;
	F_pq_getmsgend(m, v2548+int32(440))
	mBase = m.M
	v9798 = m.ExcPending
	if v9798 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2133
	}
L2128:
	;
	v9709 = int32(0)
	goto L2129
L2129:
	;
	v9750 = F_pq_getmsgint(m, v2548+int32(440), int32(4))
	mBase = m.M
	v9751 = m.ExcPending
	if v9751 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2131
	}
L2130:
	;
	v9773 = v9703
	goto L2127
L2131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9703+v9709<<(uint(int32(2))%32)))) = v9750
	v9754 = v9709 + int32(1)
	if v9754 != v9697 {
		v9709 = v9754
		goto L2129
	} else {
		goto L2132
	}
L2132:
	;
	goto L2130
L2133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[110])) = v9694
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+512)) = v9773
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+460)) = v9697
	v9804 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[160])))
	F_pgstat_report_activity(m, int32(3), v9694)
	mBase = m.M
	if v9804 == int32(1) {
		goto L2134
	} else {
		goto L2135
	}
L2134:
	;
	v9809 = int32(_a_F_PostgresMain_202)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[211])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[212])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[213])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[214])) = int64(1)
	v9819 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L2137
L2135:
	;
	goto L2136
L2136:
	;
	v9824 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v9825 = m.ExcPending
	if v9825 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2138
	}
L2137:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMain_203))
	mBase = m.M
	goto L2136
L2138:
	;
	if v9824 != 0 {
		goto L2139
	} else {
		goto L2140
	}
L2139:
	;
	v9826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9692))))
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+68)) = v9694
	if v9826 != 0 {
		goto L2142
	} else {
		goto L2143
	}
L2140:
	;
	goto L2141
L2141:
	;
	F_start_xact_command(m)
	mBase = m.M
	v9843 = m.ExcPending
	if v9843 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2147
	}
L2142:
	;
	v9829 = v9692
	goto L2144
L2143:
	;
	v9829 = int32(_a_F_PostgresMain_217)
	goto L2144
L2144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+64)) = v9829
	F_errmsg_internal(m, int32(_a_F_PostgresMain_218), v2548-int32(-64))
	mBase = m.M
	v9835 = m.ExcPending
	if v9835 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2145
	}
L2145:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(1544), int32(_a_F_PostgresMain_219))
	mBase = m.M
	v9840 = m.ExcPending
	if v9840 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2146
	}
L2146:
	;
	goto L2141
L2147:
	;
	v9844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9692))))
	if v9844 != 0 {
		goto L2149
	} else {
		goto L2150
	}
L2148:
	;
	v9866 = int32(_a_F_PostgresMain_166)
	v9867 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v9864
	v9871 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[216])))
	if v9871 == int32(1) {
		goto L2157
	} else {
		goto L2158
	}
L2149:
	;
	v9846 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[49]))
	v9864 = v9846
	v9865 = int32(0)
	goto L2148
L2150:
	;
	goto L2151
L2151:
	;
	v9849 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[215]))
	if v9849 != 0 {
		goto L2152
	} else {
		goto L2153
	}
L2152:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[215])) = int32(0)
	F_DropCachedPlan(m, v9849)
	mBase = m.M
	v9854 = m.ExcPending
	if v9854 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2155
	}
L2153:
	;
	goto L2154
L2154:
	;
	v9856 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[49]))
	v9861 = F_AllocSetContextCreateInternal(m, v9856, int32(_a_F_PostgresMain_220), int32(0), int32(_a_F_PostgresMain_16), int32(_a_F_PostgresMain_17))
	mBase = m.M
	v9862 = m.ExcPending
	if v9862 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2156
	}
L2155:
	;
	goto L2154
L2156:
	;
	v9864 = v9861
	v9865 = v9861
	goto L2148
L2157:
	;
	v9874 = int32(_a_F_PostgresMain_202)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[211])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[212])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[213])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[214])) = int64(1)
	v9884 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L2160
L2158:
	;
	goto L2159
L2159:
	;
	v9888 = F_raw_parser(m, v9694, int32(0))
	mBase = m.M
	v9889 = m.ExcPending
	if v9889 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2161
	}
L2160:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMain_203))
	mBase = m.M
	goto L2159
L2161:
	;
	v9891 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[216])))
	if v9891 == int32(1) {
		goto L2162
	} else {
		goto L2163
	}
L2162:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMain_204))
	mBase = m.M
	v9896 = m.ExcPending
	if v9896 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2165
	}
L2163:
	;
	goto L2164
L2164:
	;
	v9898 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[217])))
	if v9898 == int32(1) {
		goto L2166
	} else {
		goto L2167
	}
L2165:
	;
	goto L2164
L2166:
	;
	v9903 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[218])))
	F_elog_node_display(m, int32(_a_F_PostgresMain_205), v9888, v9903)
	mBase = m.M
	v9905 = m.ExcPending
	if v9905 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2169
	}
L2167:
	;
	goto L2168
L2168:
	;
	if v9888 != 0 {
		goto L2171
	} else {
		goto L2172
	}
L2169:
	;
	goto L2168
L2170:
	;
	if v9865 != 0 {
		goto L2196
	} else {
		goto L2197
	}
L2171:
	;
	v9906 = *(*int32)(unsafe.Add(mBase, uint32(v9888)+4))
	if int32(2) <= v9906 {
		goto L688
	} else {
		goto L2174
	}
L2172:
	;
	goto L2173
L2173:
	;
	v9963 = int32(0)
	v9966 = F_CreateCachedPlan(m, v9963, v9694, v9963)
	mBase = m.M
	v9967 = m.ExcPending
	if v9967 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2195
	}
L2174:
	;
	v9909 = *(*int32)(unsafe.Add(mBase, uint32(v9888)+12))
	v9910 = *(*int32)(unsafe.Add(mBase, uint32(v9909)))
	v9912 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v9913 = *(*int32)(unsafe.Add(mBase, uint32(v9912)+24))
	goto L2175
L2175:
	;
	v9920 = *(*int32)(unsafe.Add(mBase, uint32(v9910)+4))
	if (v9913-int32(7))&int32(-9) == int32(0) {
		goto L2176
	} else {
		goto L2177
	}
L2176:
	;
	if v9920 == int32(0) {
		goto L687
	} else {
		goto L2179
	}
L2177:
	;
	goto L2178
L2178:
	;
	v9931 = F_CreateCommandTag(m, v9920)
	mBase = m.M
	v9932 = m.ExcPending
	if v9932 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2182
	}
L2179:
	;
	v9923 = *(*int32)(unsafe.Add(mBase, uint32(v9920)))
	if v9923 != int32(225) {
		goto L687
	} else {
		goto L2180
	}
L2180:
	;
	v9926 = *(*int32)(unsafe.Add(mBase, uint32(v9920)+4))
	if (v9926-int32(2))&int32(-6) != 0 {
		goto L687
	} else {
		goto L2181
	}
L2181:
	;
	goto L2178
L2182:
	;
	v9933 = F_CreateCachedPlan(m, v9910, v9694, v9931)
	mBase = m.M
	v9934 = m.ExcPending
	if v9934 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2183
	}
L2183:
	;
	v9937 = *(*int32)(unsafe.Add(mBase, uint32(v9910)+4))
	v9938 = *(*int32)(unsafe.Add(mBase, uint32(v9937)))
	switch v9938 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v9942 = int32(1)
		goto L2185
	default:
		goto L2186
	}
L2184:
	;
	if v9942 == int32(0) {
		goto L2187
	} else {
		goto L2188
	}
L2185:
	;
	goto L2184
L2186:
	;
	v9942 = int32(0)
	goto L2185
L2187:
	;
	v9949 = F_pg_analyze_and_rewrite_varparams(m, v9910, v9694, v2548+int32(512), v2548+int32(460))
	mBase = m.M
	v9950 = m.ExcPending
	if v9950 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2190
	}
L2188:
	;
	goto L2189
L2189:
	;
	v9951 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v9952 = m.ExcPending
	if v9952 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2191
	}
L2190:
	;
	v9968 = v9933
	v9970 = v9949
	goto L2170
L2191:
	;
	F_PushActiveSnapshot(m, v9951)
	mBase = m.M
	v9954 = m.ExcPending
	if v9954 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2192
	}
L2192:
	;
	v9959 = F_pg_analyze_and_rewrite_varparams(m, v9910, v9694, v2548+int32(512), v2548+int32(460))
	mBase = m.M
	v9960 = m.ExcPending
	if v9960 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2193
	}
L2193:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v9962 = m.ExcPending
	if v9962 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2194
	}
L2194:
	;
	v9968 = v9933
	v9970 = v9959
	goto L2170
L2195:
	;
	v9968 = v9966
	v9970 = v9963
	goto L2170
L2196:
	;
	v9971 = *(*int32)(unsafe.Add(mBase, uint32(v9968)+56))
	v9973 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[49]))
	v9977 = *(*int32)(unsafe.Add(mBase, uint32(v9971)+16))
	if v9977 != v9973 {
		goto L2200
	} else {
		goto L2201
	}
L2197:
	;
	goto L2198
L2198:
	;
	v10006 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+512))
	v10007 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+460))
	v10008 = int32(0)
	F_CompleteCachedPlan(m, v9968, v9970, v9865, v10006, v10007, v10008, v10008, int32(2048), int32(1))
	mBase = m.M
	v10013 = m.ExcPending
	if v10013 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2216
	}
L2199:
	;
	goto L2198
L2200:
	;
	if v9977 == int32(0) {
		goto L2203
	} else {
		goto L2204
	}
L2201:
	;
	goto L2202
L2202:
	;
	goto L2199
L2203:
	;
	if v9973 != 0 {
		goto L2210
	} else {
		goto L2211
	}
L2204:
	;
	v9981 = *(*int32)(unsafe.Add(mBase, uint32(v9971)+28))
	v9982 = *(*int32)(unsafe.Add(mBase, uint32(v9971)+24))
	if v9982 != 0 {
		goto L2206
	} else {
		goto L2207
	}
L2205:
	;
	if v9981 == int32(0) {
		goto L2203
	} else {
		goto L2209
	}
L2206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9982)+28)) = v9981
	goto L2205
L2207:
	;
	goto L2208
L2208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9977)+20)) = v9981
	goto L2205
L2209:
	;
	v9987 = *(*int32)(unsafe.Add(mBase, uint32(v9971)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9981)+24)) = v9987
	goto L2203
L2210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9971)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9971)+16)) = v9973
	v9994 = *(*int32)(unsafe.Add(mBase, uint32(v9973)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v9971)+28)) = v9994
	if v9994 != 0 {
		goto L2213
	} else {
		goto L2214
	}
L2211:
	;
	goto L2212
L2212:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9971)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9971)+16)) = int32(0)
	goto L2202
L2213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9994)+24)) = v9971
	goto L2215
L2214:
	;
	goto L2215
L2215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9973)+20)) = v9971
	goto L2199
L2216:
	;
	v10015 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[148]))
	if v10015 != 0 {
		goto L2217
	} else {
		goto L2218
	}
L2217:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v10017 = m.ExcPending
	if v10017 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2220
	}
L2218:
	;
	goto L2219
L2219:
	;
	if v9844 != 0 {
		goto L2222
	} else {
		goto L2223
	}
L2220:
	;
	goto L2219
L2221:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v9867
	F_CommandCounterIncrement(m)
	mBase = m.M
	v10028 = m.ExcPending
	if v10028 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2227
	}
L2222:
	;
	F_StorePreparedStatement(m, v9692, v9968, int32(0))
	mBase = m.M
	v10020 = m.ExcPending
	if v10020 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2225
	}
L2223:
	;
	goto L2224
L2224:
	;
	F_SaveCachedPlan(m, v9968)
	mBase = m.M
	v10022 = m.ExcPending
	if v10022 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2226
	}
L2225:
	;
	goto L2221
L2226:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[215])) = v9968
	goto L2221
L2227:
	;
	v10030 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	if v10030 == int32(2) {
		goto L2228
	} else {
		goto L2229
	}
L2228:
	;
	F_pq_putemptymessage(m, int32(49))
	mBase = m.M
	v10035 = m.ExcPending
	if v10035 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2231
	}
L2229:
	;
	goto L2230
L2230:
	;
	v10039 = F_check_log_duration(m, v2548+int32(480), int32(0))
	mBase = m.M
	v10040 = m.ExcPending
	if v10040 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2236
	}
L2231:
	;
	goto L2230
L2232:
	;
	if v9804 != 0 {
		goto L2248
	} else {
		goto L2249
	}
L2233:
	;
	F_errhidestmt(m)
	mBase = m.M
	v10081 = m.ExcPending
	if v10081 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2246
	}
L2234:
	;
	v10060 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v10061 = m.ExcPending
	if v10061 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2240
	}
L2235:
	;
	v10045 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v10046 = m.ExcPending
	if v10046 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2237
	}
L2236:
	;
	switch v10039 - int32(1) {
	case 0:
		goto L2235
	case 1:
		goto L2234
	default:
		goto L2232
	}
L2237:
	;
	if v10045 == int32(0) {
		goto L2232
	} else {
		goto L2238
	}
L2238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+32)) = v2548 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMain_214), v2548+int32(32))
	mBase = m.M
	v10056 = m.ExcPending
	if v10056 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2239
	}
L2239:
	;
	v10079 = int32(1724)
	goto L2233
L2240:
	;
	if v10060 == int32(0) {
		goto L2232
	} else {
		goto L2241
	}
L2241:
	;
	v10064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9692))))
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+56)) = v9694
	if v10064 != 0 {
		goto L2242
	} else {
		goto L2243
	}
L2242:
	;
	v10067 = v9692
	goto L2244
L2243:
	;
	v10067 = int32(_a_F_PostgresMain_217)
	goto L2244
L2244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+52)) = v10067
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+48)) = v2548 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMain_221), v2548+int32(48))
	mBase = m.M
	v10076 = m.ExcPending
	if v10076 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2245
	}
L2245:
	;
	v10079 = int32(1732)
	goto L2233
L2246:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), v10079, int32(_a_F_PostgresMain_219))
	mBase = m.M
	v10085 = m.ExcPending
	if v10085 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2247
	}
L2247:
	;
	goto L2232
L2248:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMain_222))
	mBase = m.M
	v10089 = m.ExcPending
	if v10089 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2251
	}
L2249:
	;
	goto L2250
L2250:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[110])) = int32(0)
	goto L690
L2251:
	;
	goto L2250
L2252:
	;
	v10098 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[158]))
	if v10098 < int32(0) {
		goto L2254
	} else {
		goto L2255
	}
L2253:
	;
	v10105 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[160])))
	v10107 = v2548 + int32(440)
	v10108 = F_pq_getmsgstring(m, v10107)
	mBase = m.M
	v10109 = m.ExcPending
	if v10109 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2257
	}
L2254:
	;
	v10102 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[159])) = v10102
	goto L2256
L2255:
	;
	goto L2256
L2256:
	;
	goto L2253
L2257:
	;
	v10110 = F_pq_getmsgstring(m, v10107)
	mBase = m.M
	v10111 = m.ExcPending
	if v10111 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2258
	}
L2258:
	;
	v10114 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v10115 = m.ExcPending
	if v10115 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2259
	}
L2259:
	;
	if v10114 != 0 {
		goto L2260
	} else {
		goto L2261
	}
L2260:
	;
	v10116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10108))))
	v10118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10110))))
	if v10118 != 0 {
		goto L2263
	} else {
		goto L2264
	}
L2261:
	;
	goto L2262
L2262:
	;
	v10135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10110))))
	if v10135 != 0 {
		goto L2272
	} else {
		goto L2273
	}
L2263:
	;
	v10119 = v10110
	goto L2265
L2264:
	;
	v10119 = int32(_a_F_PostgresMain_217)
	goto L2265
L2265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+212)) = v10119
	if v10116 != 0 {
		goto L2266
	} else {
		goto L2267
	}
L2266:
	;
	v10122 = v10108
	goto L2268
L2267:
	;
	v10122 = int32(_a_F_PostgresMain_217)
	goto L2268
L2268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+208)) = v10122
	F_errmsg_internal(m, int32(_a_F_PostgresMain_223), v2548+int32(208))
	mBase = m.M
	v10128 = m.ExcPending
	if v10128 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2269
	}
L2269:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(1778), int32(_a_F_PostgresMain_224))
	mBase = m.M
	v10133 = m.ExcPending
	if v10133 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2270
	}
L2270:
	;
	goto L2262
L2271:
	;
	v10146 = *(*int32)(unsafe.Add(mBase, uint32(v10144)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[110])) = v10146
	F_pgstat_report_activity(m, int32(3), v10146)
	mBase = m.M
	v10150 = *(*int32)(unsafe.Add(mBase, uint32(v10144)+60))
	if v10150 == int32(0) {
		goto L2277
	} else {
		goto L2278
	}
L2272:
	;
	v10137 = F_FetchPreparedStatement(m, v10110, int32(1))
	mBase = m.M
	v10138 = m.ExcPending
	if v10138 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2275
	}
L2273:
	;
	goto L2274
L2274:
	;
	v10141 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[215]))
	if v10141 == int32(0) {
		goto L685
	} else {
		goto L2276
	}
L2275:
	;
	v10139 = *(*int32)(unsafe.Add(mBase, uint32(v10137)+64))
	v10144 = v10139
	goto L2271
L2276:
	;
	v10144 = v10141
	goto L2271
L2277:
	;
	if v10105&int32(1) != 0 {
		goto L2291
	} else {
		goto L2292
	}
L2278:
	;
	v10153 = *(*int32)(unsafe.Add(mBase, uint32(v10150)+4))
	if v10153 <= int32(0) {
		goto L2277
	} else {
		goto L2279
	}
L2279:
	;
	v10156 = *(*int32)(unsafe.Add(mBase, uint32(v10150)+12))
	v10162 = int32(0)
	goto L2280
L2280:
	;
	v10200 = *(*int32)(unsafe.Add(mBase, uint32(v10156+v10162<<(uint(int32(2))%32))))
	v10201 = *(*int64)(unsafe.Add(mBase, uint32(v10200)+16))
	if v10201 == int64(0) {
		goto L2282
	} else {
		goto L2283
	}
L2281:
	;
	v10210 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[220]))
	if v10210 == int32(0) {
		goto L2287
	} else {
		goto L2288
	}
L2282:
	;
	v10205 = v10162 + int32(1)
	if v10205 != v10153 {
		v10162 = v10205
		goto L2280
	} else {
		goto L2285
	}
L2283:
	;
	goto L2284
L2284:
	;
	goto L2281
L2285:
	;
	goto L2277
L2286:
	;
	goto L2277
L2287:
	;
	goto L2286
L2288:
	;
	v10214 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[221])))
	if v10214&int32(1) == int32(0) {
		goto L2287
	} else {
		goto L2289
	}
L2289:
	;
	v10221 = *(*int64)(unsafe.Add(mBase, uint32(v10210)+392))
	if int32(1)&base.B2i32(v10221 != int64(0)) != 0 {
		goto L2287
	} else {
		goto L2290
	}
L2290:
	;
	v10225 = int32(_a_F_PostgresMain_209)
	v10227 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222]))
	v10228 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222])) = v10227 + v10228
	v10231 = *(*int32)(unsafe.Add(mBase, uint32(v10210)))
	*(*int32)(unsafe.Add(mBase, uint32(v10210))) = v10231 + v10228
	v10235 = int32(0)
	v10237 = int32(_a_F_PostgresMain_210)
	v10238 = base.AtomicRmwOr32(m, v10235, v10237, v10235)
	*(*int64)(unsafe.Add(mBase, uint32(v10210)+392)) = v10201
	v10243 = base.AtomicRmwOr32(m, v10235, v10237, v10235)
	v10244 = *(*int32)(unsafe.Add(mBase, uint32(v10210)))
	*(*int32)(unsafe.Add(mBase, uint32(v10210))) = v10244 + v10228
	v10250 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222])) = v10250 - v10228
	goto L2287
L2291:
	;
	v10295 = int32(_a_F_PostgresMain_202)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[211])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[212])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[213])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[214])) = int64(1)
	v10305 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L2294
L2292:
	;
	goto L2293
L2293:
	;
	F_start_xact_command(m)
	mBase = m.M
	v10309 = m.ExcPending
	if v10309 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2295
	}
L2294:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMain_203))
	mBase = m.M
	goto L2293
L2295:
	;
	v10312 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[49]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v10312
	v10317 = F_pq_getmsgint(m, v2548+int32(440), int32(2))
	mBase = m.M
	v10318 = m.ExcPending
	if v10318 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2296
	}
L2296:
	;
	if int32(0) < v10317 {
		goto L2297
	} else {
		goto L2298
	}
L2297:
	;
	v10323 = F_palloc_mul(m, int32(2), v10317)
	mBase = m.M
	v10324 = m.ExcPending
	if v10324 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2300
	}
L2298:
	;
	v10379 = v2541
	goto L2299
L2299:
	;
	v10418 = F_pq_getmsgint(m, v2548+int32(440), int32(2))
	mBase = m.M
	v10419 = m.ExcPending
	if v10419 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2305
	}
L2300:
	;
	v10329 = int32(0)
	goto L2301
L2301:
	;
	v10370 = F_pq_getmsgint(m, v2548+int32(440), int32(2))
	mBase = m.M
	v10371 = m.ExcPending
	if v10371 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2303
	}
L2302:
	;
	v10379 = v10323
	goto L2299
L2303:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v10323+v10329<<(uint(int32(1))%32)))) = uint16(v10370)
	v10374 = v10329 + int32(1)
	if v10374 != v10317 {
		v10329 = v10374
		goto L2301
	} else {
		goto L2304
	}
L2304:
	;
	goto L2302
L2305:
	;
	if base.B2i32(v10418 != v10317)&base.B2i32(int32(2) <= v10317) != 0 {
		goto L684
	} else {
		goto L2306
	}
L2306:
	;
	v10424 = *(*int32)(unsafe.Add(mBase, uint32(v10144)+24))
	if v10418 != v10424 {
		goto L683
	} else {
		goto L2307
	}
L2307:
	;
	v10427 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v10428 = *(*int32)(unsafe.Add(mBase, uint32(v10427)+24))
	goto L2308
L2308:
	;
	if (v10428-int32(7))&int32(-9) == int32(0) {
		goto L2309
	} else {
		goto L2310
	}
L2309:
	;
	v10435 = *(*int32)(unsafe.Add(mBase, uint32(v10144)+4))
	if v10435 == int32(0) {
		goto L682
	} else {
		goto L2312
	}
L2310:
	;
	goto L2311
L2311:
	;
	v10451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10108))))
	if v10451 == int32(0) {
		goto L2317
	} else {
		goto L2318
	}
L2312:
	;
	v10438 = *(*int32)(unsafe.Add(mBase, uint32(v10435)+4))
	if v10438 == int32(0) {
		goto L682
	} else {
		goto L2313
	}
L2313:
	;
	v10441 = *(*int32)(unsafe.Add(mBase, uint32(v10438)))
	if v10441 != int32(225) {
		goto L682
	} else {
		goto L2314
	}
L2314:
	;
	v10444 = *(*int32)(unsafe.Add(mBase, uint32(v10438)+4))
	if (v10444-int32(2))&int32(-6)|v10418 != 0 {
		goto L682
	} else {
		goto L2315
	}
L2315:
	;
	goto L2311
L2316:
	;
	v10463 = int32(_a_F_PostgresMain_166)
	v10464 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50]))
	v10466 = *(*int32)(unsafe.Add(mBase, uint32(v10462)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v10466
	v10468 = *(*int32)(unsafe.Add(mBase, uint32(v10144)+12))
	v10469 = F_pstrdup(m, v10468)
	mBase = m.M
	v10470 = m.ExcPending
	if v10470 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2322
	}
L2317:
	;
	v10454 = int32(1)
	v10456 = F_CreatePortal(m, v10108, v10454, v10454)
	mBase = m.M
	v10457 = m.ExcPending
	if v10457 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2320
	}
L2318:
	;
	goto L2319
L2319:
	;
	v10458 = int32(0)
	v10460 = F_CreatePortal(m, v10108, v10458, v10458)
	mBase = m.M
	v10461 = m.ExcPending
	if v10461 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2321
	}
L2320:
	;
	v10462 = v10456
	goto L2316
L2321:
	;
	v10462 = v10460
	goto L2316
L2322:
	;
	v10471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10110))))
	if v10471 != 0 {
		goto L2323
	} else {
		goto L2324
	}
L2323:
	;
	v10472 = F_pstrdup(m, v10110)
	mBase = m.M
	v10473 = m.ExcPending
	if v10473 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2326
	}
L2324:
	;
	v10474 = v2541
	goto L2325
L2325:
	;
	if v10418 <= int32(0) {
		goto L2329
	} else {
		goto L2330
	}
L2326:
	;
	v10474 = v10472
	goto L2325
L2327:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v10464
	v10795 = *(*int32)(unsafe.Add(mBase, uint32(v10462)))
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+464)) = v10780
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+460)) = v10795
	v10798 = int32(_a_F_PostgresMain_225)
	v10799 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[58]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[58])) = v2548 + int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+516)) = int32(1261)
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+512)) = v10799
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+520)) = v2548 + int32(460)
	v10814 = F_pq_getmsgint(m, v2548+int32(440), int32(2))
	mBase = m.M
	v10815 = m.ExcPending
	if v10815 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2393
	}
L2328:
	;
	v10780 = v10740
	v10792 = int32(1)
	goto L2327
L2329:
	;
	v10477 = int32(0)
	v10478 = *(*int32)(unsafe.Add(mBase, uint32(v10144)+4))
	if v10478 == v10477 {
		v10780 = v2541
		v10792 = v10477
		goto L2327
	} else {
		goto L2332
	}
L2330:
	;
	goto L2331
L2331:
	;
	v10496 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v10497 = m.ExcPending
	if v10497 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2339
	}
L2332:
	;
	v10484 = *(*int32)(unsafe.Add(mBase, uint32(v10478)+4))
	v10485 = *(*int32)(unsafe.Add(mBase, uint32(v10484)))
	switch v10485 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v10489 = int32(1)
		goto L2334
	default:
		goto L2335
	}
L2333:
	;
	if v10489 == int32(0) {
		v10780 = v2541
		v10792 = int32(0)
		goto L2327
	} else {
		goto L2336
	}
L2334:
	;
	goto L2333
L2335:
	;
	v10489 = int32(0)
	goto L2334
L2336:
	;
	v10492 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v10493 = m.ExcPending
	if v10493 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2337
	}
L2337:
	;
	F_PushActiveSnapshot(m, v10492)
	mBase = m.M
	v10495 = m.ExcPending
	if v10495 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2338
	}
L2338:
	;
	v10740 = v2541
	goto L2328
L2339:
	;
	F_PushActiveSnapshot(m, v10496)
	mBase = m.M
	v10499 = m.ExcPending
	if v10499 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2340
	}
L2340:
	;
	v10500 = *(*int32)(unsafe.Add(mBase, uint32(v10462)))
	*(*int64)(unsafe.Add(mBase, uint32(v2548)+464)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+460)) = v10500
	v10504 = int32(_a_F_PostgresMain_225)
	v10505 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[58]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[58])) = v2548 + int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+516)) = int32(1260)
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+512)) = v10505
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+520)) = v2548 + int32(460)
	v10516 = F_makeParamList(m, v10418)
	mBase = m.M
	v10517 = m.ExcPending
	if v10517 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2341
	}
L2341:
	;
	v10520 = int32(0)
	v10527 = v10520
	v10535 = v2541
	goto L2342
L2342:
	;
	v10563 = v10527 << (uint(int32(2)) % 32)
	v10564 = *(*int32)(unsafe.Add(mBase, uint32(v10144)+20))
	v10566 = *(*int32)(unsafe.Add(mBase, uint32(v10563+v10564)))
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+468)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+464)) = v10527
	v10573 = F_pq_getmsgint(m, v2548+int32(440), int32(4))
	mBase = m.M
	v10574 = m.ExcPending
	if v10574 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2345
	}
L2343:
	;
	v10701 = int32(_a_F_PostgresMain_225)
	v10703 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[58]))
	v10704 = *(*int32)(unsafe.Add(mBase, uint32(v10703)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[58])) = v10704
	v10707 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[224]))
	if v10707 == int32(0) {
		v10740 = v10516
		goto L2328
	} else {
		goto L2391
	}
L2344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+480)) = v10590
	if int32(2) <= v10317 {
		goto L2354
	} else {
		goto L2355
	}
L2345:
	;
	v10576 = base.B2i32(v10573 == int32(-1))
	if v10573 == int32(-1) {
		goto L2346
	} else {
		goto L2347
	}
L2346:
	;
	v10577 = int32(0)
	v10590 = v10577
	v10592 = v10577
	goto L2344
L2347:
	;
	goto L2348
L2348:
	;
	v10581 = F_pq_getmsgbytes(m, v2548+int32(440), v10573)
	mBase = m.M
	v10582 = m.ExcPending
	if v10582 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2349
	}
L2349:
	;
	v10583 = v10581 + v10573
	v10584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10583))))
	v10585 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10583))) = uint8(v10585)
	*(*int64)(unsafe.Add(mBase, uint32(v2548)+488)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+484)) = v10573
	v10590 = v10581
	v10592 = v10584
	goto L2344
L2350:
	;
	if v10576 == int32(0) {
		goto L2387
	} else {
		goto L2388
	}
L2351:
	;
	F_getTypeBinaryInputInfo(m, v10566, v2548+int32(472), v2548+int32(456))
	mBase = m.M
	v10666 = m.ExcPending
	if v10666 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2380
	}
L2352:
	;
	F_getTypeInputInfo(m, v10566, v2548+int32(472), v2548+int32(456))
	mBase = m.M
	v10607 = m.ExcPending
	if v10607 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2358
	}
L2353:
	;
	v10600 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10599))))
	switch v10600 {
	case 0:
		goto L2352
	case 1:
		goto L2351
	default:
		goto L680
	}
L2354:
	;
	v10599 = v10379 + v10527<<(uint(int32(1))%32)
	goto L2353
L2355:
	;
	goto L2356
L2356:
	;
	if v10317 <= v10520 {
		goto L2352
	} else {
		goto L2357
	}
L2357:
	;
	v10599 = v10379
	goto L2353
L2358:
	;
	if v10573 == int32(-1) {
		goto L2359
	} else {
		goto L2360
	}
L2359:
	;
	v10612 = int32(0)
	goto L2361
L2360:
	;
	v10609 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+480))
	v10610 = F_pg_client_to_server(m, v10609, v10573)
	mBase = m.M
	v10611 = m.ExcPending
	if v10611 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2362
	}
L2361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+468)) = v10612
	v10614 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+472))
	v10615 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+456))
	v10617 = F_OidInputFunctionCall(m, v10614, v10612, v10615, int32(-1))
	mBase = m.M
	v10618 = m.ExcPending
	if v10618 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2363
	}
L2362:
	;
	v10612 = v10610
	goto L2361
L2363:
	;
	v10619 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+468)) = v10619
	if v10612 == v10619 {
		v10681 = v10535
		v10684 = v10617
		goto L2350
	} else {
		goto L2364
	}
L2364:
	;
	v10624 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[224]))
	if v10624 != 0 {
		goto L2365
	} else {
		goto L2366
	}
L2365:
	;
	v10625 = int32(_a_F_PostgresMain_166)
	v10626 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50]))
	v10629 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[49]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v10629
	if v10535 == int32(0) {
		goto L2369
	} else {
		goto L2370
	}
L2366:
	;
	v10655 = v10535
	goto L2367
L2367:
	;
	v10657 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+480))
	if v10612 == v10657 {
		v10681 = v10655
		v10684 = v10617
		goto L2350
	} else {
		goto L2378
	}
L2368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10638+v10563))) = v10648
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v10626
	v10655 = v10638
	goto L2367
L2369:
	;
	v10634 = F_palloc0_mul(m, int32(4), v10418)
	mBase = m.M
	v10635 = m.ExcPending
	if v10635 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2372
	}
L2370:
	;
	v10638 = v10535
	v10639 = v10624
	goto L2371
L2371:
	;
	if v10639 < int32(0) {
		goto L2373
	} else {
		goto L2374
	}
L2372:
	;
	v10637 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[224]))
	v10638 = v10634
	v10639 = v10637
	goto L2371
L2373:
	;
	v10642 = F_pstrdup(m, v10612)
	mBase = m.M
	v10643 = m.ExcPending
	if v10643 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2376
	}
L2374:
	;
	goto L2375
L2375:
	;
	v10646 = F_pnstrdup(m, v10612, v10639+int32(8))
	mBase = m.M
	v10647 = m.ExcPending
	if v10647 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2377
	}
L2376:
	;
	v10648 = v10642
	goto L2368
L2377:
	;
	v10648 = v10646
	goto L2368
L2378:
	;
	F_pfree(m, v10612)
	mBase = m.M
	v10660 = m.ExcPending
	if v10660 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2379
	}
L2379:
	;
	v10681 = v10655
	v10684 = v10617
	goto L2350
L2380:
	;
	v10667 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+472))
	if v10573 == int32(-1) {
		goto L2381
	} else {
		goto L2382
	}
L2381:
	;
	v10671 = int32(0)
	goto L2383
L2382:
	;
	v10671 = v2548 + int32(480)
	goto L2383
L2383:
	;
	v10672 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+456))
	v10674 = F_OidReceiveFunctionCall(m, v10667, v10671, v10672, int32(-1))
	mBase = m.M
	v10675 = m.ExcPending
	if v10675 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2384
	}
L2384:
	;
	if v10573 == int32(-1) {
		v10681 = v10535
		v10684 = v10674
		goto L2350
	} else {
		goto L2385
	}
L2385:
	;
	v10676 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+492))
	v10677 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+484))
	if v10676 != v10677 {
		goto L681
	} else {
		goto L2386
	}
L2386:
	;
	v10681 = v10535
	v10684 = v10674
	goto L2350
L2387:
	;
	v10687 = *(*int32)(unsafe.Add(mBase, uint32(v2548)+480))
	*(*uint8)(unsafe.Add(mBase, uint32(v10687+v10573))) = uint8(v10592)
	goto L2389
L2388:
	;
	goto L2389
L2389:
	;
	v10692 = v10516 + int32(32) + v10527<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v10692)+12)) = v10566
	v10694 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v10692)+10)) = uint16(v10694)
	*(*uint8)(unsafe.Add(mBase, uint32(v10692)+8)) = uint8(v10576)
	*(*int64)(unsafe.Add(mBase, uint32(v10692))) = v10684
	v10699 = v10527 + v10694
	if v10699 != v10418 {
		v10527 = v10699
		v10535 = v10681
		goto L2342
	} else {
		goto L2390
	}
L2390:
	;
	goto L2343
L2391:
	;
	v10710 = F_BuildParamLogString(m, v10516, v10681, v10707)
	mBase = m.M
	v10711 = m.ExcPending
	if v10711 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2392
	}
L2392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10516)+24)) = v10710
	v10740 = v10516
	goto L2328
L2393:
	;
	if int32(0) < v10814 {
		goto L2394
	} else {
		goto L2395
	}
L2394:
	;
	v10820 = F_palloc_mul(m, int32(2), v10814)
	mBase = m.M
	v10821 = m.ExcPending
	if v10821 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2397
	}
L2395:
	;
	v10890 = int32(0)
	goto L2396
L2396:
	;
	F_pq_getmsgend(m, v2548+int32(440))
	mBase = m.M
	v10915 = m.ExcPending
	if v10915 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2402
	}
L2397:
	;
	v10826 = int32(0)
	goto L2398
L2398:
	;
	v10867 = F_pq_getmsgint(m, v2548+int32(440), int32(2))
	mBase = m.M
	v10868 = m.ExcPending
	if v10868 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2400
	}
L2399:
	;
	v10890 = v10820
	goto L2396
L2400:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v10820+v10826<<(uint(int32(1))%32)))) = uint16(v10867)
	v10871 = v10826 + int32(1)
	if v10871 != v10814 {
		v10826 = v10871
		goto L2398
	} else {
		goto L2401
	}
L2401:
	;
	goto L2399
L2402:
	;
	v10916 = int32(0)
	v10918 = F_GetCachedPlan(m, v10144, v10780, v10916, v10916)
	mBase = m.M
	v10919 = m.ExcPending
	if v10919 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2403
	}
L2403:
	;
	v10920 = *(*int32)(unsafe.Add(mBase, uint32(v10144)+16))
	v10921 = *(*int32)(unsafe.Add(mBase, uint32(v10918)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10462)+80)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10462)+60)) = v10918
	*(*int32)(unsafe.Add(mBase, uint32(v10462)+56)) = v10921
	*(*int64)(unsafe.Add(mBase, uint32(v10462)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10462)+40)) = v10920
	*(*int32)(unsafe.Add(mBase, uint32(v10462)+36)) = v10920
	*(*int32)(unsafe.Add(mBase, uint32(v10462)+32)) = v10469
	*(*int32)(unsafe.Add(mBase, uint32(v10462)+4)) = v10474
	goto L2404
L2404:
	;
	v10932 = *(*int32)(unsafe.Add(mBase, uint32(v10462)+56))
	if v10932 == int32(0) {
		goto L2405
	} else {
		goto L2406
	}
L2405:
	;
	if v10792 != 0 {
		goto L2419
	} else {
		goto L2420
	}
L2406:
	;
	v10935 = *(*int32)(unsafe.Add(mBase, uint32(v10932)+4))
	if v10935 <= int32(0) {
		goto L2405
	} else {
		goto L2407
	}
L2407:
	;
	v10938 = *(*int32)(unsafe.Add(mBase, uint32(v10932)+12))
	v10944 = int32(0)
	goto L2408
L2408:
	;
	v10982 = *(*int32)(unsafe.Add(mBase, uint32(v10938+v10944<<(uint(int32(2))%32))))
	v10983 = *(*int64)(unsafe.Add(mBase, uint32(v10982)+16))
	if v10983 == int64(0) {
		goto L2410
	} else {
		goto L2411
	}
L2409:
	;
	v10992 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[220]))
	if v10992 == int32(0) {
		goto L2415
	} else {
		goto L2416
	}
L2410:
	;
	v10987 = v10944 + int32(1)
	if v10987 != v10935 {
		v10944 = v10987
		goto L2408
	} else {
		goto L2413
	}
L2411:
	;
	goto L2412
L2412:
	;
	goto L2409
L2413:
	;
	goto L2405
L2414:
	;
	goto L2405
L2415:
	;
	goto L2414
L2416:
	;
	v10996 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[221])))
	if v10996&int32(1) == int32(0) {
		goto L2415
	} else {
		goto L2417
	}
L2417:
	;
	v11003 = *(*int64)(unsafe.Add(mBase, uint32(v10992)+400))
	if int32(1)&base.B2i32(v11003 != int64(0)) != 0 {
		goto L2415
	} else {
		goto L2418
	}
L2418:
	;
	v11007 = int32(_a_F_PostgresMain_209)
	v11009 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222]))
	v11010 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222])) = v11009 + v11010
	v11013 = *(*int32)(unsafe.Add(mBase, uint32(v10992)))
	*(*int32)(unsafe.Add(mBase, uint32(v10992))) = v11013 + v11010
	v11017 = int32(0)
	v11019 = int32(_a_F_PostgresMain_210)
	v11020 = base.AtomicRmwOr32(m, v11017, v11019, v11017)
	*(*int64)(unsafe.Add(mBase, uint32(v10992)+400)) = v10983
	v11025 = base.AtomicRmwOr32(m, v11017, v11019, v11017)
	v11026 = *(*int32)(unsafe.Add(mBase, uint32(v10992)))
	*(*int32)(unsafe.Add(mBase, uint32(v10992))) = v11026 + v11010
	v11032 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222])) = v11032 - v11010
	goto L2415
L2419:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v11076 = m.ExcPending
	if v11076 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2422
	}
L2420:
	;
	goto L2421
L2421:
	;
	v11077 = int32(0)
	F_PortalStart(m, v10462, v10780, v11077, v11077)
	mBase = m.M
	v11080 = m.ExcPending
	if v11080 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2423
	}
L2422:
	;
	goto L2421
L2423:
	;
	F_PortalSetResultFormat(m, v10462, v10814, v10890)
	mBase = m.M
	v11082 = m.ExcPending
	if v11082 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2424
	}
L2424:
	;
	v11083 = int32(_a_F_PostgresMain_225)
	v11085 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[58]))
	v11086 = *(*int32)(unsafe.Add(mBase, uint32(v11085)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[58])) = v11086
	v11089 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	if v11089 == int32(2) {
		goto L2425
	} else {
		goto L2426
	}
L2425:
	;
	F_pq_putemptymessage(m, int32(50))
	mBase = m.M
	v11094 = m.ExcPending
	if v11094 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2428
	}
L2426:
	;
	goto L2427
L2427:
	;
	v11098 = F_check_log_duration(m, v2548+int32(480), int32(0))
	mBase = m.M
	v11099 = m.ExcPending
	if v11099 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2433
	}
L2428:
	;
	goto L2427
L2429:
	;
	if v10105&int32(1) != 0 {
		goto L2459
	} else {
		goto L2460
	}
L2430:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), v11175, int32(_a_F_PostgresMain_224))
	mBase = m.M
	v11179 = m.ExcPending
	if v11179 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2458
	}
L2431:
	;
	v11121 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v11122 = m.ExcPending
	if v11122 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2438
	}
L2432:
	;
	v11104 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v11105 = m.ExcPending
	if v11105 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2434
	}
L2433:
	;
	switch v11098 - int32(1) {
	case 0:
		goto L2432
	case 1:
		goto L2431
	default:
		goto L2429
	}
L2434:
	;
	if v11104 == int32(0) {
		goto L2429
	} else {
		goto L2435
	}
L2435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+96)) = v2548 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMain_214), v2548+int32(96))
	mBase = m.M
	v11115 = m.ExcPending
	if v11115 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2436
	}
L2436:
	;
	F_errhidestmt(m)
	mBase = m.M
	v11117 = m.ExcPending
	if v11117 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2437
	}
L2437:
	;
	v11175 = int32(2201)
	goto L2430
L2438:
	;
	if v11121 == int32(0) {
		goto L2429
	} else {
		goto L2439
	}
L2439:
	;
	v11125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10110))))
	v11126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10108))))
	v11127 = *(*int32)(unsafe.Add(mBase, uint32(v10144)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+144)) = v11127
	if v11126 != 0 {
		goto L2440
	} else {
		goto L2441
	}
L2440:
	;
	v11130 = v10108
	goto L2442
L2441:
	;
	v11130 = int32(_a_F_PostgresMain_213)
	goto L2442
L2442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+140)) = v11130
	if v11126 != 0 {
		goto L2443
	} else {
		goto L2444
	}
L2443:
	;
	v11134 = int32(_a_F_PostgresMain_226)
	goto L2445
L2444:
	;
	v11134 = int32(_a_F_PostgresMain_213)
	goto L2445
L2445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+136)) = v11134
	if v11125 != 0 {
		goto L2446
	} else {
		goto L2447
	}
L2446:
	;
	v11137 = v10110
	goto L2448
L2447:
	;
	v11137 = int32(_a_F_PostgresMain_217)
	goto L2448
L2448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+132)) = v11137
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+128)) = v2548 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMain_227), v2548+int32(128))
	mBase = m.M
	v11146 = m.ExcPending
	if v11146 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2449
	}
L2449:
	;
	F_errhidestmt(m)
	mBase = m.M
	v11148 = m.ExcPending
	if v11148 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2450
	}
L2450:
	;
	v11149 = int32(2212)
	if v10780 == int32(0) {
		v11175 = v11149
		goto L2430
	} else {
		goto L2451
	}
L2451:
	;
	v11152 = *(*int32)(unsafe.Add(mBase, uint32(v10780)+28))
	if v11152 <= int32(0) {
		v11175 = v11149
		goto L2430
	} else {
		goto L2452
	}
L2452:
	;
	v11156 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[225]))
	if v11156 == int32(0) {
		v11175 = v11149
		goto L2430
	} else {
		goto L2453
	}
L2453:
	;
	v11160 = F_BuildParamLogString(m, v10780, int32(0), v11156)
	mBase = m.M
	v11161 = m.ExcPending
	if v11161 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2454
	}
L2454:
	;
	if v11160 == int32(0) {
		v11175 = v11149
		goto L2430
	} else {
		goto L2455
	}
L2455:
	;
	v11164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11160))))
	if v11164 == int32(0) {
		v11175 = v11149
		goto L2430
	} else {
		goto L2456
	}
L2456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+112)) = v11160
	v11171 = F_errdetail(m, int32(_a_F_PostgresMain_228), v2548+int32(112))
	mBase = m.M
	v11172 = m.ExcPending
	if v11172 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2457
	}
L2457:
	;
	v11175 = v11149
	goto L2430
L2458:
	;
	goto L2429
L2459:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMain_229))
	mBase = m.M
	v11187 = m.ExcPending
	if v11187 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2462
	}
L2460:
	;
	goto L2461
L2461:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[110])) = int32(0)
	goto L690
L2462:
	;
	goto L2461
L2463:
	;
	v11196 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[158]))
	if v11196 < int32(0) {
		goto L2465
	} else {
		goto L2466
	}
L2464:
	;
	v11203 = v2548 + int32(440)
	v11204 = F_pq_getmsgstring(m, v11203)
	mBase = m.M
	v11205 = m.ExcPending
	if v11205 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2468
	}
L2465:
	;
	v11200 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[159])) = v11200
	goto L2467
L2466:
	;
	goto L2467
L2467:
	;
	goto L2464
L2468:
	;
	v11207 = F_pq_getmsgint(m, v11203, int32(4))
	mBase = m.M
	v11208 = m.ExcPending
	if v11208 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2469
	}
L2469:
	;
	F_pq_getmsgend(m, v11203)
	mBase = m.M
	v11210 = m.ExcPending
	if v11210 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2470
	}
L2470:
	;
	v11212 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	v11214 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[160])))
	v11215 = F_GetPortalByName(m, v11204)
	mBase = m.M
	v11216 = m.ExcPending
	if v11216 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2471
	}
L2471:
	;
	if v11215 == int32(0) {
		goto L678
	} else {
		goto L2472
	}
L2472:
	;
	if v11212 == int32(2) {
		goto L2473
	} else {
		goto L2474
	}
L2473:
	;
	v11222 = int32(3)
	goto L2475
L2474:
	;
	v11222 = v11212
	goto L2475
L2475:
	;
	v11223 = *(*int32)(unsafe.Add(mBase, uint32(v11215)+36))
	if v11223 == int32(0) {
		goto L2476
	} else {
		goto L2477
	}
L2476:
	;
	F_NullCommand(m, v11222)
	mBase = m.M
	v11227 = m.ExcPending
	if v11227 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2479
	}
L2477:
	;
	goto L2478
L2478:
	;
	v11228 = *(*int32)(unsafe.Add(mBase, uint32(v11215)+56))
	if v11228 == int32(0) {
		v11247 = v2541
		goto L2480
	} else {
		goto L2481
	}
L2479:
	;
	goto L690
L2480:
	;
	v11248 = *(*int32)(unsafe.Add(mBase, uint32(v11215)+32))
	v11249 = F_pstrdup(m, v11248)
	mBase = m.M
	v11250 = m.ExcPending
	if v11250 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2487
	}
L2481:
	;
	v11231 = *(*int32)(unsafe.Add(mBase, uint32(v11228)+4))
	if v11231 != int32(1) {
		v11247 = v2541
		goto L2480
	} else {
		goto L2482
	}
L2482:
	;
	v11234 = *(*int32)(unsafe.Add(mBase, uint32(v11228)+12))
	v11235 = *(*int32)(unsafe.Add(mBase, uint32(v11234)))
	v11236 = *(*int32)(unsafe.Add(mBase, uint32(v11235)+4))
	if v11236 == int32(6) {
		goto L2483
	} else {
		goto L2484
	}
L2483:
	;
	v11240 = *(*int32)(unsafe.Add(mBase, uint32(v11235)+100))
	v11241 = *(*int32)(unsafe.Add(mBase, uint32(v11240)))
	if v11241 == int32(225) {
		v11247 = int32(1)
		goto L2480
	} else {
		goto L2486
	}
L2484:
	;
	goto L2485
L2485:
	;
	v11247 = int32(0)
	goto L2480
L2486:
	;
	goto L2485
L2487:
	;
	v11251 = *(*int32)(unsafe.Add(mBase, uint32(v11215)+4))
	if v11251 != 0 {
		goto L2488
	} else {
		goto L2489
	}
L2488:
	;
	v11252 = F_pstrdup(m, v11251)
	mBase = m.M
	v11253 = m.ExcPending
	if v11253 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2491
	}
L2489:
	;
	v11255 = int32(_a_F_PostgresMain_217)
	goto L2490
L2490:
	;
	v11256 = *(*int32)(unsafe.Add(mBase, uint32(v11215)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[110])) = v11249
	F_pgstat_report_activity(m, int32(3), v11249)
	mBase = m.M
	v11261 = *(*int32)(unsafe.Add(mBase, uint32(v11215)+56))
	if v11261 == int32(0) {
		goto L2492
	} else {
		goto L2493
	}
L2491:
	;
	v11255 = v11252
	goto L2490
L2492:
	;
	v11520 = *(*int32)(unsafe.Add(mBase, uint32(v11215)+36))
	v11525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11520<<(uint(int32(3))%32))+uint32(_c_F_PostgresMain[223]))))
	*(*int32)(unsafe.Add(mBase, uint32(v2548+int32(456)))) = v11525
	goto L2526
L2493:
	;
	v11264 = *(*int32)(unsafe.Add(mBase, uint32(v11261)+4))
	if v11264 <= int32(0) {
		goto L2492
	} else {
		goto L2494
	}
L2494:
	;
	v11267 = int32(0)
	if v11267 < v11264 {
		goto L2495
	} else {
		goto L2496
	}
L2495:
	;
	v11270 = v11264
	goto L2497
L2496:
	;
	v11270 = v11267
	goto L2497
L2497:
	;
	v11271 = *(*int32)(unsafe.Add(mBase, uint32(v11261)+12))
	v11277 = int32(0)
	goto L2499
L2498:
	;
	if v11379 <= int32(0) {
		goto L2492
	} else {
		goto L2514
	}
L2499:
	;
	v11315 = *(*int32)(unsafe.Add(mBase, uint32(v11271+v11277<<(uint(int32(2))%32))))
	v11316 = *(*int64)(unsafe.Add(mBase, uint32(v11315)+8))
	if v11316 == int64(0) {
		goto L2501
	} else {
		goto L2502
	}
L2500:
	;
	v11325 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[220]))
	if v11325 == int32(0) {
		goto L2506
	} else {
		goto L2507
	}
L2501:
	;
	v11320 = v11277 + int32(1)
	if v11320 != v11264 {
		v11277 = v11320
		goto L2499
	} else {
		goto L2504
	}
L2502:
	;
	goto L2503
L2503:
	;
	goto L2500
L2504:
	;
	v11377 = v11270
	v11379 = v11264
	v11380 = v11261
	goto L2498
L2505:
	;
	v11369 = *(*int32)(unsafe.Add(mBase, uint32(v11215)+56))
	if v11369 == int32(0) {
		goto L2492
	} else {
		goto L2510
	}
L2506:
	;
	goto L2505
L2507:
	;
	v11329 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[221])))
	if v11329&int32(1) == int32(0) {
		goto L2506
	} else {
		goto L2508
	}
L2508:
	;
	v11336 = *(*int64)(unsafe.Add(mBase, uint32(v11325)+392))
	if int32(1)&base.B2i32(v11336 != int64(0)) != 0 {
		goto L2506
	} else {
		goto L2509
	}
L2509:
	;
	v11340 = int32(_a_F_PostgresMain_209)
	v11342 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222]))
	v11343 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222])) = v11342 + v11343
	v11346 = *(*int32)(unsafe.Add(mBase, uint32(v11325)))
	*(*int32)(unsafe.Add(mBase, uint32(v11325))) = v11346 + v11343
	v11350 = int32(0)
	v11352 = int32(_a_F_PostgresMain_210)
	v11353 = base.AtomicRmwOr32(m, v11350, v11352, v11350)
	*(*int64)(unsafe.Add(mBase, uint32(v11325)+392)) = v11316
	v11358 = base.AtomicRmwOr32(m, v11350, v11352, v11350)
	v11359 = *(*int32)(unsafe.Add(mBase, uint32(v11325)))
	*(*int32)(unsafe.Add(mBase, uint32(v11325))) = v11359 + v11343
	v11365 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222])) = v11365 - v11343
	goto L2506
L2510:
	;
	v11372 = *(*int32)(unsafe.Add(mBase, uint32(v11369)+4))
	v11373 = int32(0)
	if v11373 < v11372 {
		goto L2511
	} else {
		goto L2512
	}
L2511:
	;
	v11376 = v11372
	goto L2513
L2512:
	;
	v11376 = v11373
	goto L2513
L2513:
	;
	v11377 = v11376
	v11379 = v11372
	v11380 = v11369
	goto L2498
L2514:
	;
	v11383 = *(*int32)(unsafe.Add(mBase, uint32(v11380)+12))
	v11389 = int32(0)
	goto L2515
L2515:
	;
	v11427 = *(*int32)(unsafe.Add(mBase, uint32(v11383+v11389<<(uint(int32(2))%32))))
	v11428 = *(*int64)(unsafe.Add(mBase, uint32(v11427)+16))
	if v11428 == int64(0) {
		goto L2517
	} else {
		goto L2518
	}
L2516:
	;
	v11437 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[220]))
	if v11437 == int32(0) {
		goto L2522
	} else {
		goto L2523
	}
L2517:
	;
	v11432 = v11389 + int32(1)
	if v11377 != v11432 {
		v11389 = v11432
		goto L2515
	} else {
		goto L2520
	}
L2518:
	;
	goto L2519
L2519:
	;
	goto L2516
L2520:
	;
	goto L2492
L2521:
	;
	goto L2492
L2522:
	;
	goto L2521
L2523:
	;
	v11441 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[221])))
	if v11441&int32(1) == int32(0) {
		goto L2522
	} else {
		goto L2524
	}
L2524:
	;
	v11448 = *(*int64)(unsafe.Add(mBase, uint32(v11437)+400))
	if int32(1)&base.B2i32(v11448 != int64(0)) != 0 {
		goto L2522
	} else {
		goto L2525
	}
L2525:
	;
	v11452 = int32(_a_F_PostgresMain_209)
	v11454 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222]))
	v11455 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222])) = v11454 + v11455
	v11458 = *(*int32)(unsafe.Add(mBase, uint32(v11437)))
	*(*int32)(unsafe.Add(mBase, uint32(v11437))) = v11458 + v11455
	v11462 = int32(0)
	v11464 = int32(_a_F_PostgresMain_210)
	v11465 = base.AtomicRmwOr32(m, v11462, v11464, v11462)
	*(*int64)(unsafe.Add(mBase, uint32(v11437)+400)) = v11428
	v11470 = base.AtomicRmwOr32(m, v11462, v11464, v11462)
	v11471 = *(*int32)(unsafe.Add(mBase, uint32(v11437)))
	*(*int32)(unsafe.Add(mBase, uint32(v11437))) = v11471 + v11455
	v11477 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[222])) = v11477 - v11455
	goto L2522
L2526:
	;
	if v11214&int32(1) != 0 {
		goto L2527
	} else {
		goto L2528
	}
L2527:
	;
	v11531 = int32(_a_F_PostgresMain_202)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[211])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[212])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[213])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[214])) = int64(1)
	v11541 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L2530
L2528:
	;
	goto L2529
L2529:
	;
	v11545 = F_CreateDestReceiver(m, v11222)
	mBase = m.M
	v11546 = m.ExcPending
	if v11546 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2531
	}
L2530:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMain_203))
	mBase = m.M
	goto L2529
L2531:
	;
	if v11222 == int32(3) {
		goto L2532
	} else {
		goto L2533
	}
L2532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11545)+20)) = v11215
	goto L2534
L2533:
	;
	goto L2534
L2534:
	;
	F_start_xact_command(m)
	mBase = m.M
	v11551 = m.ExcPending
	if v11551 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2535
	}
L2535:
	;
	v11552 = int32(0)
	v11553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11215)+116)))
	v11555 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[219]))
	switch v11555 {
	case 0:
		v11719 = v11552
		goto L2536
	default:
		goto L2538
	case 3:
		goto L2537
	}
L2536:
	;
	v11755 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v11756 = *(*int32)(unsafe.Add(mBase, uint32(v11755)+24))
	goto L2570
L2537:
	;
	v11657 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v11658 = m.ExcPending
	if v11658 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2546
	}
L2538:
	;
	v11556 = *(*int32)(unsafe.Add(mBase, uint32(v11215)+56))
	if v11556 == int32(0) {
		v11719 = v11552
		goto L2536
	} else {
		goto L2539
	}
L2539:
	;
	v11559 = *(*int32)(unsafe.Add(mBase, uint32(v11556)+4))
	if v11559 <= int32(0) {
		v11719 = v11552
		goto L2536
	} else {
		goto L2540
	}
L2540:
	;
	v11566 = v11552
	goto L2541
L2541:
	;
	v11601 = *(*int32)(unsafe.Add(mBase, uint32(v11556)+12))
	v11605 = *(*int32)(unsafe.Add(mBase, uint32(v11601+v11566<<(uint(int32(2))%32))))
	v11606 = F_GetCommandLogLevel(m, v11605)
	mBase = m.M
	v11607 = m.ExcPending
	if v11607 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2543
	}
L2542:
	;
	v11719 = int32(0)
	goto L2536
L2543:
	;
	v11609 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[219]))
	if base.Ui32(v11606) <= base.Ui32(v11609) {
		goto L2537
	} else {
		goto L2544
	}
L2544:
	;
	v11612 = v11566 + int32(1)
	v11613 = *(*int32)(unsafe.Add(mBase, uint32(v11556)+4))
	if v11612 < v11613 {
		v11566 = v11612
		goto L2541
	} else {
		goto L2545
	}
L2545:
	;
	goto L2542
L2546:
	;
	if v11657 == int32(0) {
		goto L2547
	} else {
		goto L2548
	}
L2547:
	;
	v11719 = int32(1)
	goto L2536
L2548:
	;
	goto L2549
L2549:
	;
	v11662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11204))))
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+336)) = v11249
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+324)) = v11255
	if v11662 != 0 {
		goto L2550
	} else {
		goto L2551
	}
L2550:
	;
	v11666 = v11204
	goto L2552
L2551:
	;
	v11666 = int32(_a_F_PostgresMain_213)
	goto L2552
L2552:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+332)) = v11666
	if v11662 != 0 {
		goto L2553
	} else {
		goto L2554
	}
L2553:
	;
	v11670 = int32(_a_F_PostgresMain_226)
	goto L2555
L2554:
	;
	v11670 = int32(_a_F_PostgresMain_213)
	goto L2555
L2555:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+328)) = v11670
	v11672 = int32(1)
	if v11553&v11672 != 0 {
		goto L2556
	} else {
		goto L2557
	}
L2556:
	;
	v11677 = int32(_a_F_PostgresMain_230)
	goto L2558
L2557:
	;
	v11677 = int32(_a_F_PostgresMain_231)
	goto L2558
L2558:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+320)) = v11677
	F_errmsg(m, int32(_a_F_PostgresMain_232), v2548+int32(320))
	mBase = m.M
	v11683 = m.ExcPending
	if v11683 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2559
	}
L2559:
	;
	F_errhidestmt(m)
	mBase = m.M
	v11685 = m.ExcPending
	if v11685 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2560
	}
L2560:
	;
	if v11256 == int32(0) {
		goto L2561
	} else {
		goto L2562
	}
L2561:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(2362), int32(_a_F_PostgresMain_233))
	mBase = m.M
	v11714 = m.ExcPending
	if v11714 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2569
	}
L2562:
	;
	v11688 = *(*int32)(unsafe.Add(mBase, uint32(v11256)+28))
	if v11688 <= int32(0) {
		goto L2561
	} else {
		goto L2563
	}
L2563:
	;
	v11692 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[225]))
	if v11692 == int32(0) {
		goto L2561
	} else {
		goto L2564
	}
L2564:
	;
	v11696 = F_BuildParamLogString(m, v11256, int32(0), v11692)
	mBase = m.M
	v11697 = m.ExcPending
	if v11697 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2565
	}
L2565:
	;
	if v11696 == int32(0) {
		goto L2561
	} else {
		goto L2566
	}
L2566:
	;
	v11700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11696))))
	if v11700 == int32(0) {
		goto L2561
	} else {
		goto L2567
	}
L2567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+304)) = v11696
	v11707 = F_errdetail(m, int32(_a_F_PostgresMain_228), v2548+int32(304))
	mBase = m.M
	v11708 = m.ExcPending
	if v11708 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2568
	}
L2568:
	;
	goto L2561
L2569:
	;
	v11719 = v11672
	goto L2536
L2570:
	;
	if (v11756-int32(7))&int32(-9) == int32(0) {
		goto L2571
	} else {
		goto L2572
	}
L2571:
	;
	v11763 = *(*int32)(unsafe.Add(mBase, uint32(v11215)+56))
	if v11763 == int32(0) {
		goto L677
	} else {
		goto L2574
	}
L2572:
	;
	goto L2573
L2573:
	;
	v11787 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[148]))
	if v11787 != 0 {
		goto L2580
	} else {
		goto L2581
	}
L2574:
	;
	v11766 = *(*int32)(unsafe.Add(mBase, uint32(v11763)+4))
	if v11766 != int32(1) {
		goto L677
	} else {
		goto L2575
	}
L2575:
	;
	v11769 = *(*int32)(unsafe.Add(mBase, uint32(v11763)+12))
	v11770 = *(*int32)(unsafe.Add(mBase, uint32(v11769)))
	v11771 = *(*int32)(unsafe.Add(mBase, uint32(v11770)+4))
	if v11771 != int32(6) {
		goto L677
	} else {
		goto L2576
	}
L2576:
	;
	v11774 = *(*int32)(unsafe.Add(mBase, uint32(v11770)+100))
	if v11774 == int32(0) {
		goto L677
	} else {
		goto L2577
	}
L2577:
	;
	v11777 = *(*int32)(unsafe.Add(mBase, uint32(v11774)))
	if v11777 != int32(225) {
		goto L677
	} else {
		goto L2578
	}
L2578:
	;
	v11780 = *(*int32)(unsafe.Add(mBase, uint32(v11774)+4))
	if (v11780-int32(2))&int32(-6) != 0 {
		goto L677
	} else {
		goto L2579
	}
L2579:
	;
	goto L2573
L2580:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v11789 = m.ExcPending
	if v11789 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2583
	}
L2581:
	;
	goto L2582
L2582:
	;
	v11790 = *(*int32)(unsafe.Add(mBase, uint32(v11215)))
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+476)) = v11256
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+472)) = v11790
	v11793 = int32(_a_F_PostgresMain_225)
	v11794 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[58]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[58])) = v2548 + int32(460)
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+464)) = int32(1261)
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+460)) = v11794
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+468)) = v2548 + int32(472)
	if v11207 <= int32(0) {
		goto L2584
	} else {
		goto L2585
	}
L2583:
	;
	goto L2582
L2584:
	;
	v11808 = int32(2147483647)
	goto L2586
L2585:
	;
	v11808 = v11207
	goto L2586
L2586:
	;
	v11812 = F_PortalRun(m, v11215, v11808, int32(1), v11545, v11545, v2548+int32(512))
	mBase = m.M
	v11813 = m.ExcPending
	if v11813 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2587
	}
L2587:
	;
	v11814 = *(*int32)(unsafe.Add(mBase, uint32(v11545)+12))
	m.T0[v11814].(func(*base.Module, int32))(m, v11545)
	mBase = m.M
	v11816 = m.ExcPending
	if v11816 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2588
	}
L2588:
	;
	v11817 = int32(_a_F_PostgresMain_225)
	v11819 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[58]))
	v11820 = *(*int32)(unsafe.Add(mBase, uint32(v11819)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[58])) = v11820
	if v11812 != 0 {
		goto L2590
	} else {
		goto L2591
	}
L2589:
	;
	v11885 = F_check_log_duration(m, v2548+int32(480), v11719)
	mBase = m.M
	v11886 = m.ExcPending
	if v11886 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2619
	}
L2590:
	;
	if v11247 == int32(0) {
		goto L2595
	} else {
		goto L2596
	}
L2591:
	;
	goto L2592
L2592:
	;
	v11870 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	if v11870 == int32(2) {
		goto L2611
	} else {
		goto L2612
	}
L2593:
	;
	F_EndCommand(m, v2548+int32(512), v11222)
	mBase = m.M
	v11868 = m.ExcPending
	if v11868 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2610
	}
L2594:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v11848 = m.ExcPending
	if v11848 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2606
	}
L2595:
	;
	v11825 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[226])))
	if v11825&int32(4) == int32(0) {
		goto L2594
	} else {
		goto L2598
	}
L2596:
	;
	goto L2597
L2597:
	;
	v11833 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[63])))
	goto L2599
L2598:
	;
	goto L2597
L2599:
	;
	if v11833 != 0 {
		goto L2600
	} else {
		goto L2601
	}
L2600:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v11836 = m.ExcPending
	if v11836 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2603
	}
L2601:
	;
	goto L2602
L2602:
	;
	v11837 = int32(0)
	v11839 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])))
	if v11839 == v11837 {
		v11864 = v11837
		goto L2593
	} else {
		goto L2604
	}
L2603:
	;
	goto L2602
L2604:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v11843 = m.ExcPending
	if v11843 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2605
	}
L2605:
	;
	v11845 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])) = uint8(v11845)
	v11864 = v11837
	goto L2593
L2606:
	;
	v11849 = int32(_a_F_PostgresMain_234)
	v11851 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[226]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[226])) = v11851 | int32(8)
	v11858 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[63])))
	goto L2607
L2607:
	;
	if v11858 == int32(0) {
		v11864 = v11256
		goto L2593
	} else {
		goto L2608
	}
L2608:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v11863 = m.ExcPending
	if v11863 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2609
	}
L2609:
	;
	v11864 = v11256
	goto L2593
L2610:
	;
	v11882 = v11864
	goto L2589
L2611:
	;
	F_pq_putemptymessage(m, int32(115))
	mBase = m.M
	v11875 = m.ExcPending
	if v11875 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2614
	}
L2612:
	;
	goto L2613
L2613:
	;
	v11876 = int32(_a_F_PostgresMain_234)
	v11878 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[226]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[226])) = v11878 | int32(8)
	v11882 = v11256
	goto L2589
L2614:
	;
	goto L2613
L2615:
	;
	if v11214&int32(1) != 0 {
		goto L2645
	} else {
		goto L2646
	}
L2616:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), v11963, int32(_a_F_PostgresMain_233))
	mBase = m.M
	v11967 = m.ExcPending
	if v11967 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2644
	}
L2617:
	;
	v11908 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v11909 = m.ExcPending
	if v11909 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2624
	}
L2618:
	;
	v11891 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v11892 = m.ExcPending
	if v11892 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2620
	}
L2619:
	;
	switch v11885 - int32(1) {
	case 0:
		goto L2618
	case 1:
		goto L2617
	default:
		goto L2615
	}
L2620:
	;
	if v11891 == int32(0) {
		goto L2615
	} else {
		goto L2621
	}
L2621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+240)) = v2548 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMain_214), v2548+int32(240))
	mBase = m.M
	v11902 = m.ExcPending
	if v11902 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2622
	}
L2622:
	;
	F_errhidestmt(m)
	mBase = m.M
	v11904 = m.ExcPending
	if v11904 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2623
	}
L2623:
	;
	v11963 = int32(2472)
	goto L2616
L2624:
	;
	if v11908 == int32(0) {
		goto L2615
	} else {
		goto L2625
	}
L2625:
	;
	v11912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11204))))
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+292)) = v11249
	if v11912 != 0 {
		goto L2626
	} else {
		goto L2627
	}
L2626:
	;
	v11915 = v11204
	goto L2628
L2627:
	;
	v11915 = int32(_a_F_PostgresMain_213)
	goto L2628
L2628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+288)) = v11915
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+280)) = v11255
	if v11912 != 0 {
		goto L2629
	} else {
		goto L2630
	}
L2629:
	;
	v11920 = int32(_a_F_PostgresMain_226)
	goto L2631
L2630:
	;
	v11920 = int32(_a_F_PostgresMain_213)
	goto L2631
L2631:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+284)) = v11920
	if v11553&int32(1) != 0 {
		goto L2632
	} else {
		goto L2633
	}
L2632:
	;
	v11926 = int32(_a_F_PostgresMain_230)
	goto L2634
L2633:
	;
	v11926 = int32(_a_F_PostgresMain_231)
	goto L2634
L2634:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+276)) = v11926
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+272)) = v2548 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMain_235), v2548+int32(272))
	mBase = m.M
	v11935 = m.ExcPending
	if v11935 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2635
	}
L2635:
	;
	F_errhidestmt(m)
	mBase = m.M
	v11937 = m.ExcPending
	if v11937 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2636
	}
L2636:
	;
	v11938 = int32(2486)
	if v11882 == int32(0) {
		v11963 = v11938
		goto L2616
	} else {
		goto L2637
	}
L2637:
	;
	v11941 = *(*int32)(unsafe.Add(mBase, uint32(v11882)+28))
	if v11941 <= int32(0) {
		v11963 = v11938
		goto L2616
	} else {
		goto L2638
	}
L2638:
	;
	v11945 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[225]))
	if v11945 == int32(0) {
		v11963 = v11938
		goto L2616
	} else {
		goto L2639
	}
L2639:
	;
	v11949 = F_BuildParamLogString(m, v11882, int32(0), v11945)
	mBase = m.M
	v11950 = m.ExcPending
	if v11950 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2640
	}
L2640:
	;
	if v11949 == int32(0) {
		v11963 = v11938
		goto L2616
	} else {
		goto L2641
	}
L2641:
	;
	v11953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11949))))
	if v11953 == int32(0) {
		v11963 = v11938
		goto L2616
	} else {
		goto L2642
	}
L2642:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+256)) = v11949
	v11960 = F_errdetail(m, int32(_a_F_PostgresMain_228), v2548+int32(256))
	mBase = m.M
	v11961 = m.ExcPending
	if v11961 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2643
	}
L2643:
	;
	v11963 = v11938
	goto L2616
L2644:
	;
	goto L2615
L2645:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMain_236))
	mBase = m.M
	v11974 = m.ExcPending
	if v11974 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2648
	}
L2646:
	;
	goto L2647
L2647:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[110])) = int32(0)
	goto L690
L2648:
	;
	goto L2647
L2649:
	;
	v11983 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[158]))
	if v11983 < int32(0) {
		goto L2651
	} else {
		goto L2652
	}
L2650:
	;
	F_pgstat_report_activity(m, int32(5), int32(0))
	mBase = m.M
	F_start_xact_command(m)
	mBase = m.M
	v11993 = m.ExcPending
	if v11993 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2654
	}
L2651:
	;
	v11987 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[159])) = v11987
	goto L2653
L2652:
	;
	goto L2653
L2653:
	;
	goto L2650
L2654:
	;
	v11996 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[49]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v11996
	v11999 = v2548 + int32(440)
	v12000 = m.G0
	v12002 = v12000 - int32(2368)
	m.G0 = v12002
	v12005 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v12006 = *(*int32)(unsafe.Add(mBase, uint32(v12005)+24))
	goto L2665
L2655:
	;
	v12865 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[148]))
	if v12865 != 0 {
		goto L2838
	} else {
		goto L2839
	}
L2656:
	;
	v12821 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+236))
	v12822 = m.T0[v12821].(func(*base.Module, int32) int64)(m, v12002+int32(736))
	mBase = m.M
	v12823 = m.ExcPending
	if v12823 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2837
	}
L2657:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12764 = m.ExcPending
	if v12764 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2833
	}
L2658:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12744 = m.ExcPending
	if v12744 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2829
	}
L2659:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12726 = m.ExcPending
	if v12726 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2825
	}
L2660:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12707 = m.ExcPending
	if v12707 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2821
	}
L2661:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12687 = m.ExcPending
	if v12687 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2817
	}
L2662:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12668 = m.ExcPending
	if v12668 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2814
	}
L2663:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12648 = m.ExcPending
	if v12648 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2810
	}
L2664:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12632 = m.ExcPending
	if v12632 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2806
	}
L2665:
	;
	if base.B2i32((v12006-int32(7))&int32(-9) == int32(0)) == int32(0) {
		goto L2666
	} else {
		goto L2667
	}
L2666:
	;
	v12015 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v12016 = m.ExcPending
	if v12016 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2669
	}
L2667:
	;
	goto L2668
L2668:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12616 = m.ExcPending
	if v12616 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2802
	}
L2669:
	;
	F_PushActiveSnapshot(m, v12015)
	mBase = m.M
	v12018 = m.ExcPending
	if v12018 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2670
	}
L2670:
	;
	v12020 = F_pq_getmsgint(m, v11999, int32(4))
	mBase = m.M
	v12021 = m.ExcPending
	if v12021 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2671
	}
L2671:
	;
	base.MemoryFill(m, v12002+int32(232), int32(0), int32(504))
	v12029 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(v12020))
	mBase = m.M
	v12030 = m.ExcPending
	if v12030 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2672
	}
L2672:
	;
	if v12029 == int32(0) {
		goto L2664
	} else {
		goto L2673
	}
L2673:
	;
	v12033 = *(*int32)(unsafe.Add(mBase, uint32(v12029)+16))
	v12034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12033)+22)))
	v12035 = v12033 + v12034
	v12036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12035)+96)))
	if v12036 != int32(102) {
		goto L2663
	} else {
		goto L2674
	}
L2674:
	;
	v12039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12035)+100)))
	if v12039 == int32(1) {
		goto L2663
	} else {
		goto L2675
	}
L2675:
	;
	v12042 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12035)+104)))
	if int32(101) <= v12042 {
		goto L2662
	} else {
		goto L2676
	}
L2676:
	;
	v12045 = *(*int32)(unsafe.Add(mBase, uint32(v12035)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+264)) = v12045
	v12047 = *(*int32)(unsafe.Add(mBase, uint32(v12035)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+268)) = v12047
	v12050 = v12002 + int32(272)
	v12051 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12035)+104)))
	v12053 = v12051 << (uint(int32(2)) % 32)
	if v12053 != 0 {
		goto L2677
	} else {
		goto L2678
	}
L2677:
	;
	base.MemoryCopy(m, v12050, v12035+int32(136), v12053)
	goto L2679
L2678:
	;
	goto L2679
L2679:
	;
	v12058 = v12002 + int32(236)
	v12060 = v12002 + int32(672)
	v12062 = v12035 + int32(4)
	goto L2683
L2680:
	;
	F_ReleaseCatCache(m, v12029)
	mBase = m.M
	v12183 = m.ExcPending
	if v12183 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2711
	}
L2681:
	;
	v12179 = F_strlen(m, v12168)
	mBase = m.M
	goto L2680
L2683:
	;
	goto L2684
L2684:
	;
	v12069 = int32(63)
	if (v12060^v12062)&int32(3) != 0 {
		goto L2688
	} else {
		goto L2689
	}
L2685:
	;
	v12172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12169))) = uint8(v12172)
	goto L2681
L2686:
	;
	v12153 = v12148
	v12154 = v12149
	v12155 = v12150
	goto L2707
L2687:
	;
	if v12143 == int32(0) {
		v12168 = v12141
		v12169 = v12142
		goto L2685
	} else {
		goto L2706
	}
L2688:
	;
	v12141 = v12062
	v12142 = v12060
	v12143 = v12069
	goto L2687
L2689:
	;
	goto L2690
L2690:
	;
	v12073 = int32(0)
	if base.B2i32(v12062&int32(3) == v12073)|int32(0) == v12073 {
		goto L2692
	} else {
		goto L2693
	}
L2691:
	;
	if v12109 == int32(0) {
		v12168 = v12106
		v12169 = v12107
		goto L2685
	} else {
		goto L2700
	}
L2692:
	;
	v12085 = v12062
	v12086 = v12060
	v12087 = v12069
	goto L2695
L2693:
	;
	goto L2694
L2694:
	;
	v12106 = v12062
	v12107 = v12060
	v12108 = v12069
	v12109 = int32(1)
	goto L2691
L2695:
	;
	v12089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12085))))
	*(*uint8)(unsafe.Add(mBase, uint32(v12086))) = uint8(v12089)
	if v12089 == int32(0) {
		v12148 = v12085
		v12149 = v12086
		v12150 = v12087
		goto L2686
	} else {
		goto L2697
	}
L2696:
	;
	v12106 = v12100
	v12107 = v12094
	v12108 = v12096
	v12109 = v12098
	goto L2691
L2697:
	;
	v12093 = int32(1)
	v12094 = v12086 + v12093
	v12096 = v12087 - v12093
	v12097 = int32(0)
	v12098 = base.B2i32(v12096 != v12097)
	v12100 = v12085 + v12093
	if v12100&int32(3) == v12097 {
		v12106 = v12100
		v12107 = v12094
		v12108 = v12096
		v12109 = v12098
		goto L2691
	} else {
		goto L2698
	}
L2698:
	;
	if v12096 != 0 {
		v12085 = v12100
		v12086 = v12094
		v12087 = v12096
		goto L2695
	} else {
		goto L2699
	}
L2699:
	;
	goto L2696
L2700:
	;
	v12112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12106))))
	if base.B2i32(v12112 == int32(0))|base.B2i32(base.Ui32(v12108) < base.Ui32(int32(4))) != 0 {
		v12141 = v12106
		v12142 = v12107
		v12143 = v12108
		goto L2687
	} else {
		goto L2701
	}
L2701:
	;
	v12119 = v12106
	v12120 = v12107
	v12121 = v12108
	goto L2702
L2702:
	;
	v12124 = *(*int32)(unsafe.Add(mBase, uint32(v12119)))
	v12127 = int32(-2139062144)
	if (int32(16843008)-v12124|v12124)&v12127 != v12127 {
		v12148 = v12119
		v12149 = v12120
		v12150 = v12121
		goto L2686
	} else {
		goto L2704
	}
L2703:
	;
	v12141 = v12135
	v12142 = v12133
	v12143 = v12137
	goto L2687
L2704:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12120))) = v12124
	v12132 = int32(4)
	v12133 = v12120 + v12132
	v12135 = v12119 + v12132
	v12137 = v12121 - v12132
	if base.Ui32(int32(3)) < base.Ui32(v12137) {
		v12119 = v12135
		v12120 = v12133
		v12121 = v12137
		goto L2702
	} else {
		goto L2705
	}
L2705:
	;
	goto L2703
L2706:
	;
	v12148 = v12141
	v12149 = v12142
	v12150 = v12143
	goto L2686
L2707:
	;
	v12157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12153))))
	*(*uint8)(unsafe.Add(mBase, uint32(v12154))) = uint8(v12157)
	if v12157 == int32(0) {
		v12168 = v12153
		v12169 = v12154
		goto L2685
	} else {
		goto L2709
	}
L2708:
	;
	v12168 = v12164
	v12169 = v12162
	goto L2685
L2709:
	;
	v12161 = int32(1)
	v12162 = v12154 + v12161
	v12164 = v12153 + v12161
	v12166 = v12155 - v12161
	if v12166 != 0 {
		v12153 = v12164
		v12154 = v12162
		v12155 = v12166
		goto L2707
	} else {
		goto L2710
	}
L2710:
	;
	goto L2708
L2711:
	;
	F_fmgr_info(m, v12020, v12058)
	mBase = m.M
	v12185 = m.ExcPending
	if v12185 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2712
	}
L2712:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+232)) = v12020
	v12188 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[219]))
	if v12188 != int32(3) {
		goto L2713
	} else {
		goto L2714
	}
L2713:
	;
	v12210 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+264))
	v12212 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[227]))
	v12214 = F_object_aclcheck(m, int32(2615), v12210, v12212, int64(256))
	mBase = m.M
	v12215 = m.ExcPending
	if v12215 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2719
	}
L2714:
	;
	v12193 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v12194 = m.ExcPending
	if v12194 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2715
	}
L2715:
	;
	if v12193 == int32(0) {
		goto L2713
	} else {
		goto L2716
	}
L2716:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+180)) = v12020
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+176)) = v12060
	F_errmsg(m, int32(_a_F_PostgresMain_237), v12002+int32(176))
	mBase = m.M
	v12203 = m.ExcPending
	if v12203 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2717
	}
L2717:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_238), int32(234), int32(_a_F_PostgresMain_239))
	mBase = m.M
	v12208 = m.ExcPending
	if v12208 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2718
	}
L2718:
	;
	goto L2713
L2719:
	;
	if v12214 != 0 {
		goto L2720
	} else {
		goto L2721
	}
L2720:
	;
	v12217 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+264))
	v12218 = F_get_namespace_name(m, v12217)
	mBase = m.M
	v12219 = m.ExcPending
	if v12219 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2723
	}
L2721:
	;
	goto L2722
L2722:
	;
	v12223 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[228]))
	if v12223 != 0 {
		goto L2725
	} else {
		goto L2726
	}
L2723:
	;
	F_aclcheck_error(m, v12214, int32(37), v12218)
	mBase = m.M
	v12221 = m.ExcPending
	if v12221 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2724
	}
L2724:
	;
	goto L2722
L2725:
	;
	v12224 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+264))
	v12226 = F_RunNamespaceSearchHook(m, v12224, int32(1))
	mBase = m.M
	v12227 = m.ExcPending
	if v12227 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2728
	}
L2726:
	;
	goto L2727
L2727:
	;
	v12230 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[227]))
	v12232 = F_object_aclcheck(m, int32(1255), v12020, v12230, int64(128))
	mBase = m.M
	v12233 = m.ExcPending
	if v12233 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2729
	}
L2728:
	;
	goto L2727
L2729:
	;
	if v12232 != 0 {
		goto L2730
	} else {
		goto L2731
	}
L2730:
	;
	v12235 = F_get_func_name(m, v12020)
	mBase = m.M
	v12236 = m.ExcPending
	if v12236 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2733
	}
L2731:
	;
	goto L2732
L2732:
	;
	v12241 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[228]))
	if v12241 != 0 {
		goto L2735
	} else {
		goto L2736
	}
L2733:
	;
	F_aclcheck_error(m, v12232, int32(19), v12235)
	mBase = m.M
	v12238 = m.ExcPending
	if v12238 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2734
	}
L2734:
	;
	goto L2732
L2735:
	;
	F_RunFunctionExecuteHook(m, v12020)
	mBase = m.M
	v12243 = m.ExcPending
	if v12243 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2738
	}
L2736:
	;
	goto L2737
L2737:
	;
	v12244 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12002)+740)) = v12244
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+736)) = v12058
	*(*int64)(unsafe.Add(mBase, uint32(v12002)+745)) = v12244
	v12250 = F_pq_getmsgint(m, v11999, int32(2))
	mBase = m.M
	v12251 = m.ExcPending
	if v12251 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2739
	}
L2738:
	;
	goto L2737
L2739:
	;
	if int32(0) < v12250 {
		goto L2740
	} else {
		goto L2741
	}
L2740:
	;
	v12257 = F_palloc(m, v12250<<(uint(int32(1))%32))
	mBase = m.M
	v12258 = m.ExcPending
	if v12258 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2743
	}
L2741:
	;
	v12312 = int32(0)
	goto L2742
L2742:
	;
	v12348 = F_pq_getmsgint(m, v11999, int32(2))
	mBase = m.M
	v12349 = m.ExcPending
	if v12349 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2748
	}
L2743:
	;
	v12276 = int32(0)
	goto L2744
L2744:
	;
	v12302 = F_pq_getmsgint(m, v11999, int32(2))
	mBase = m.M
	v12303 = m.ExcPending
	if v12303 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2746
	}
L2745:
	;
	v12312 = v12257
	goto L2742
L2746:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v12257+v12276<<(uint(int32(1))%32)))) = uint16(v12302)
	v12306 = v12276 + int32(1)
	if v12306 != v12250 {
		v12276 = v12306
		goto L2744
	} else {
		goto L2747
	}
L2747:
	;
	goto L2745
L2748:
	;
	if int32(100) < v12348 {
		goto L2661
	} else {
		goto L2749
	}
L2749:
	;
	v12352 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12002)+244)))
	if v12348 != v12352 {
		goto L2661
	} else {
		goto L2750
	}
L2750:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v12002)+754)) = uint16(v12348)
	if base.B2i32(v12348 != v12250)&base.B2i32(int32(2) <= v12250) != 0 {
		goto L2660
	} else {
		goto L2751
	}
L2751:
	;
	F_initStringInfo(m, v12002+int32(192))
	mBase = m.M
	v12362 = m.ExcPending
	if v12362 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2752
	}
L2752:
	;
	if int32(0) < v12348 {
		goto L2753
	} else {
		goto L2754
	}
L2753:
	;
	v12366 = v12002 + int32(760)
	v12387 = int32(0)
	goto L2756
L2754:
	;
	goto L2755
L2755:
	;
	v12549 = F_pq_getmsgint(m, v11999, int32(2))
	mBase = m.M
	v12550 = m.ExcPending
	if v12550 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2792
	}
L2756:
	;
	v12409 = int32(4)
	v12410 = v12387 << (uint(v12409) % 32)
	v12413 = v12410 + (v12002 + int32(736))
	v12415 = F_pq_getmsgint(m, v11999, v12409)
	mBase = m.M
	v12416 = m.ExcPending
	if v12416 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2759
	}
L2757:
	;
	goto L2755
L2758:
	;
	if base.B2i32(v12250 < int32(2)) == int32(0) {
		goto L2771
	} else {
		goto L2772
	}
L2759:
	;
	if v12415 == int32(-1) {
		goto L2760
	} else {
		goto L2761
	}
L2760:
	;
	v12419 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12413)+32)) = uint8(v12419)
	goto L2758
L2761:
	;
	goto L2762
L2762:
	;
	v12421 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12413)+32)) = uint8(v12421)
	if v12415 < v12421 {
		goto L2659
	} else {
		goto L2763
	}
L2763:
	;
	v12426 = v12002 + int32(192)
	v12427 = *(*int32)(unsafe.Add(mBase, uint32(v12426)))
	v12428 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12427))) = uint8(v12428)
	*(*int32)(unsafe.Add(mBase, uint32(v12426)+12)) = v12428
	*(*int32)(unsafe.Add(mBase, uint32(v12426)+4)) = v12428
	goto L2764
L2764:
	;
	v12434 = F_pq_getmsgbytes(m, v11999, v12415)
	mBase = m.M
	v12435 = m.ExcPending
	if v12435 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2765
	}
L2765:
	;
	F_appendBinaryStringInfo(m, v12426, v12434, v12415)
	mBase = m.M
	v12437 = m.ExcPending
	if v12437 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2766
	}
L2766:
	;
	goto L2758
L2767:
	;
	v12507 = v12387 + int32(1)
	if v12507 != v12348 {
		v12387 = v12507
		goto L2756
	} else {
		goto L2791
	}
L2768:
	;
	v12482 = *(*int32)(unsafe.Add(mBase, uint32(v12050+v12387<<(uint(int32(2))%32))))
	F_getTypeBinaryInputInfo(m, v12482, v12002+int32(2364), v12002+int32(2360))
	mBase = m.M
	v12488 = m.ExcPending
	if v12488 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2784
	}
L2769:
	;
	v12452 = *(*int32)(unsafe.Add(mBase, uint32(v12050+v12387<<(uint(int32(2))%32))))
	F_getTypeInputInfo(m, v12452, v12002+int32(2364), v12002+int32(2360))
	mBase = m.M
	v12458 = m.ExcPending
	if v12458 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2775
	}
L2770:
	;
	v12447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12446))))
	switch v12447 {
	case 0:
		goto L2769
	case 1:
		goto L2768
	default:
		goto L2657
	}
L2771:
	;
	v12446 = v12312 + v12387<<(uint(int32(1))%32)
	goto L2770
L2772:
	;
	goto L2773
L2773:
	;
	if v12250 <= int32(0) {
		goto L2769
	} else {
		goto L2774
	}
L2774:
	;
	v12446 = v12312
	goto L2770
L2775:
	;
	if v12415 == int32(-1) {
		goto L2776
	} else {
		goto L2777
	}
L2776:
	;
	v12465 = int32(0)
	goto L2778
L2777:
	;
	v12462 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+192))
	v12463 = F_pg_client_to_server(m, v12462, v12415)
	mBase = m.M
	v12464 = m.ExcPending
	if v12464 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2779
	}
L2778:
	;
	v12467 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+2364))
	v12468 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+2360))
	v12470 = F_OidInputFunctionCall(m, v12467, v12465, v12468, int32(-1))
	mBase = m.M
	v12471 = m.ExcPending
	if v12471 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2780
	}
L2779:
	;
	v12465 = v12463
	goto L2778
L2780:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12366+v12410))) = v12470
	if v12465 == int32(0) {
		goto L2767
	} else {
		goto L2781
	}
L2781:
	;
	v12475 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+192))
	if v12465 == v12475 {
		goto L2767
	} else {
		goto L2782
	}
L2782:
	;
	F_pfree(m, v12465)
	mBase = m.M
	v12478 = m.ExcPending
	if v12478 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2783
	}
L2783:
	;
	goto L2767
L2784:
	;
	v12490 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+2364))
	v12495 = base.B2i32(v12415 == int32(-1))
	if v12415 == int32(-1) {
		goto L2785
	} else {
		goto L2786
	}
L2785:
	;
	v12496 = int32(0)
	goto L2787
L2786:
	;
	v12496 = v12002 + int32(192)
	goto L2787
L2787:
	;
	v12497 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+2360))
	v12499 = F_OidReceiveFunctionCall(m, v12490, v12496, v12497, int32(-1))
	mBase = m.M
	v12500 = m.ExcPending
	if v12500 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2788
	}
L2788:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12366+v12410))) = v12499
	if v12415 == int32(-1) {
		goto L2767
	} else {
		goto L2789
	}
L2789:
	;
	v12502 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+204))
	v12503 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+196))
	if v12502 != v12503 {
		goto L2658
	} else {
		goto L2790
	}
L2790:
	;
	goto L2767
L2791:
	;
	goto L2757
L2792:
	;
	F_pq_getmsgend(m, v11999)
	mBase = m.M
	v12552 = m.ExcPending
	if v12552 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2793
	}
L2793:
	;
	v12553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12002)+246)))
	if v12553 != int32(1) {
		goto L2656
	} else {
		goto L2794
	}
L2794:
	;
	v12556 = int32(0)
	if v12348 <= v12556 {
		goto L2656
	} else {
		goto L2795
	}
L2795:
	;
	v12577 = v12556
	goto L2796
L2796:
	;
	v12604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12002+int32(736)+v12577<<(uint(int32(4))%32))+32)))
	if v12604 == int32(0) {
		goto L2798
	} else {
		goto L2799
	}
L2797:
	;
	v12610 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12002)+752)) = uint8(v12610)
	v12863 = int64(0)
	goto L2655
L2798:
	;
	v12608 = v12577 + int32(1)
	if base.I32_extend16_s(v12348) != v12608 {
		v12577 = v12608
		goto L2796
	} else {
		goto L2801
	}
L2799:
	;
	goto L2800
L2800:
	;
	goto L2797
L2801:
	;
	goto L2656
L2802:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v12619 = m.ExcPending
	if v12619 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2803
	}
L2803:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_181), int32(0))
	mBase = m.M
	v12623 = m.ExcPending
	if v12623 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2804
	}
L2804:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_238), int32(209), int32(_a_F_PostgresMain_239))
	mBase = m.M
	v12628 = m.ExcPending
	if v12628 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2805
	}
L2805:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2806:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v12635 = m.ExcPending
	if v12635 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2807
	}
L2807:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12002))) = v12020
	F_errmsg(m, int32(_a_F_PostgresMain_240), v12002)
	mBase = m.M
	v12639 = m.ExcPending
	if v12639 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2808
	}
L2808:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_238), int32(141), int32(_a_F_PostgresMain_241))
	mBase = m.M
	v12644 = m.ExcPending
	if v12644 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2809
	}
L2809:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2810:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v12651 = m.ExcPending
	if v12651 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2811
	}
L2811:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+16)) = v12035 + int32(4)
	F_errmsg(m, int32(_a_F_PostgresMain_242), v12002+int32(16))
	mBase = m.M
	v12659 = m.ExcPending
	if v12659 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2812
	}
L2812:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_238), int32(149), int32(_a_F_PostgresMain_241))
	mBase = m.M
	v12664 = m.ExcPending
	if v12664 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2813
	}
L2813:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2814:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+36)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+32)) = v12035 + int32(4)
	F_errmsg_internal(m, int32(_a_F_PostgresMain_243), v12002+int32(32))
	mBase = m.M
	v12678 = m.ExcPending
	if v12678 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2815
	}
L2815:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_238), int32(154), int32(_a_F_PostgresMain_241))
	mBase = m.M
	v12683 = m.ExcPending
	if v12683 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2816
	}
L2816:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2817:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v12690 = m.ExcPending
	if v12690 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2818
	}
L2818:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+48)) = v12348
	v12692 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12002)+244)))
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+52)) = v12692
	F_errmsg(m, int32(_a_F_PostgresMain_244), v12002+int32(48))
	mBase = m.M
	v12698 = m.ExcPending
	if v12698 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2819
	}
L2819:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_238), int32(353), int32(_a_F_PostgresMain_245))
	mBase = m.M
	v12703 = m.ExcPending
	if v12703 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2820
	}
L2820:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2821:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v12710 = m.ExcPending
	if v12710 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2822
	}
L2822:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+164)) = v12348
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+160)) = v12250
	F_errmsg(m, int32(_a_F_PostgresMain_246), v12002+int32(160))
	mBase = m.M
	v12717 = m.ExcPending
	if v12717 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2823
	}
L2823:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_238), int32(361), int32(_a_F_PostgresMain_245))
	mBase = m.M
	v12722 = m.ExcPending
	if v12722 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2824
	}
L2824:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2825:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v12729 = m.ExcPending
	if v12729 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2826
	}
L2826:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+144)) = v12415
	F_errmsg(m, int32(_a_F_PostgresMain_247), v12002+int32(144))
	mBase = m.M
	v12735 = m.ExcPending
	if v12735 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2827
	}
L2827:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_238), int32(385), int32(_a_F_PostgresMain_245))
	mBase = m.M
	v12740 = m.ExcPending
	if v12740 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2828
	}
L2828:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2829:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v12747 = m.ExcPending
	if v12747 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2830
	}
L2830:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+128)) = v12387 + int32(1)
	F_errmsg(m, int32(_a_F_PostgresMain_248), v12002+int32(128))
	mBase = m.M
	v12755 = m.ExcPending
	if v12755 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2831
	}
L2831:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_238), int32(448), int32(_a_F_PostgresMain_245))
	mBase = m.M
	v12760 = m.ExcPending
	if v12760 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2832
	}
L2832:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2833:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v12767 = m.ExcPending
	if v12767 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2834
	}
L2834:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+64)) = base.I32_extend16_s(v12447)
	F_errmsg(m, int32(_a_F_PostgresMain_249), v12002-int32(-64))
	mBase = m.M
	v12774 = m.ExcPending
	if v12774 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2835
	}
L2835:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_238), int32(453), int32(_a_F_PostgresMain_245))
	mBase = m.M
	v12779 = m.ExcPending
	if v12779 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2836
	}
L2836:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2837:
	;
	v12863 = v12822
	goto L2655
L2838:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v12867 = m.ExcPending
	if v12867 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2841
	}
L2839:
	;
	goto L2840
L2840:
	;
	v12868 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+268))
	v12869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12002)+752)))
	v12871 = v12002 + int32(192)
	F_pq_beginmessage(m, v12871, int32(86))
	mBase = m.M
	v12874 = m.ExcPending
	if v12874 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2842
	}
L2841:
	;
	goto L2840
L2842:
	;
	if v12869 == int32(1) {
		goto L2844
	} else {
		goto L2845
	}
L2843:
	;
	v13028 = v12002 + int32(192)
	F_pq_endmessage(m, v13028)
	mBase = m.M
	v13030 = m.ExcPending
	if v13030 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2876
	}
L2844:
	;
	F_enlargeStringInfo(m, v12871, int32(4))
	mBase = m.M
	v12879 = m.ExcPending
	if v12879 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2847
	}
L2845:
	;
	goto L2846
L2846:
	;
	switch v12549 & int32(_a_F_PostgresMain_250) {
	case 0:
		goto L2848
	case 1:
		goto L2850
	default:
		goto L2849
	}
L2847:
	;
	v12880 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+196))
	v12881 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v12880+v12881))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+196)) = v12880 + int32(4)
	goto L2843
L2848:
	;
	F_getTypeOutputInfo(m, v12868, v12002+int32(2364), v12002+int32(2360))
	mBase = m.M
	v13011 = m.ExcPending
	if v13011 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2872
	}
L2849:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12990 = m.ExcPending
	if v12990 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2868
	}
L2850:
	;
	F_getTypeBinaryOutputInfo(m, v12868, v12002+int32(2364), v12002+int32(2360))
	mBase = m.M
	v12895 = m.ExcPending
	if v12895 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2851
	}
L2851:
	;
	v12896 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+2364))
	v12897 = m.G0
	v12899 = v12897 - int32(80)
	m.G0 = v12899
	v12902 = v12899 + int32(12)
	v12904 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50]))
	F_fmgr_info_cxt_security(m, v12896, v12902, v12904, int32(0))
	mBase = m.M
	v12907 = m.ExcPending
	if v12907 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2853
	}
L2852:
	;
	v12950 = *(*int32)(unsafe.Add(mBase, uint32(v12932)))
	v12952 = v12002 + int32(192)
	F_enlargeStringInfo(m, v12952, int32(4))
	mBase = m.M
	v12955 = m.ExcPending
	if v12955 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2865
	}
L2853:
	;
	v12908 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12899)+44)) = v12908
	*(*int64)(unsafe.Add(mBase, uint32(v12899)+49)) = v12908
	v12912 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12899)+72)) = uint8(v12912)
	*(*int64)(unsafe.Add(mBase, uint32(v12899)+64)) = v12863
	*(*int32)(unsafe.Add(mBase, uint32(v12899)+40)) = v12902
	v12916 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v12899)+58)) = uint16(v12916)
	v12920 = *(*int32)(unsafe.Add(mBase, uint32(v12899)+12))
	v12921 = m.T0[v12920].(func(*base.Module, int32) int64)(m, v12899+int32(40))
	mBase = m.M
	v12922 = m.ExcPending
	if v12922 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2854
	}
L2854:
	;
	v12923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12899)+56)))
	if v12923 != int32(1) {
		goto L2855
	} else {
		goto L2856
	}
L2855:
	;
	v12926 = base.I32_wrap_i64(v12921)
	v12927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12926))))
	if v12927&int32(3) != 0 {
		goto L2858
	} else {
		goto L2859
	}
L2856:
	;
	goto L2857
L2857:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12939 = m.ExcPending
	if v12939 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2862
	}
L2858:
	;
	v12930 = F_detoast_attr(m, v12926)
	mBase = m.M
	v12931 = m.ExcPending
	if v12931 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2861
	}
L2859:
	;
	v12932 = v12926
	goto L2860
L2860:
	;
	m.G0 = v12899 + int32(80)
	goto L2852
L2861:
	;
	v12932 = v12930
	goto L2860
L2862:
	;
	v12940 = *(*int32)(unsafe.Add(mBase, uint32(v12899)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12899))) = v12940
	F_errmsg_internal(m, int32(_a_F_PostgresMain_251), v12899)
	mBase = m.M
	v12944 = m.ExcPending
	if v12944 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2863
	}
L2863:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_252), int32(1145), int32(_a_F_PostgresMain_253))
	mBase = m.M
	v12949 = m.ExcPending
	if v12949 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2864
	}
L2864:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2865:
	;
	v12956 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+196))
	v12957 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+192))
	v12959 = int32(2)
	v12961 = int32(4)
	v12962 = int32(base.Ui32(v12950)>>(uint(v12959)%32)) - v12961
	v12963 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v12956+v12957))) = base.I32_rotr(v12962&v12963, int32(8)) | base.I32_rotr(v12962, int32(24))&v12963
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+196)) = v12956 + v12961
	v12978 = *(*int32)(unsafe.Add(mBase, uint32(v12932)))
	F_appendBinaryStringInfo(m, v12952, v12932+v12961, int32(base.Ui32(v12978)>>(uint(v12959)%32))-v12961)
	mBase = m.M
	v12984 = m.ExcPending
	if v12984 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2866
	}
L2866:
	;
	F_pfree(m, v12932)
	mBase = m.M
	v12986 = m.ExcPending
	if v12986 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2867
	}
L2867:
	;
	goto L2843
L2868:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v12993 = m.ExcPending
	if v12993 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2869
	}
L2869:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+112)) = base.I32_extend16_s(v12549)
	F_errmsg(m, int32(_a_F_PostgresMain_249), v12002+int32(112))
	mBase = m.M
	v13000 = m.ExcPending
	if v13000 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2870
	}
L2870:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_238), int32(106), int32(_a_F_PostgresMain_254))
	mBase = m.M
	v13005 = m.ExcPending
	if v13005 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2871
	}
L2871:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2872:
	;
	v13014 = *(*int32)(unsafe.Add(mBase, uint32(v12002)+2364))
	v13015 = F_OidOutputFunctionCall(m, v13014, v12863)
	mBase = m.M
	v13016 = m.ExcPending
	if v13016 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2873
	}
L2873:
	;
	v13017 = F_strlen(m, v13015)
	mBase = m.M
	F_pq_sendcountedtext(m, v12002+int32(192), v13015, v13017)
	mBase = m.M
	v13019 = m.ExcPending
	if v13019 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2874
	}
L2874:
	;
	F_pfree(m, v13015)
	mBase = m.M
	v13021 = m.ExcPending
	if v13021 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2875
	}
L2875:
	;
	goto L2843
L2876:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v13032 = m.ExcPending
	if v13032 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2877
	}
L2877:
	;
	v13036 = F_check_log_duration(m, v13028, base.B2i32(v12188 == int32(3)))
	mBase = m.M
	v13037 = m.ExcPending
	if v13037 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2882
	}
L2878:
	;
	m.G0 = v12002 + int32(2368)
	v13082 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[63])))
	goto L2890
L2879:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_238), v13072, int32(_a_F_PostgresMain_239))
	mBase = m.M
	v13075 = m.ExcPending
	if v13075 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2889
	}
L2880:
	;
	v13057 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v13058 = m.ExcPending
	if v13058 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2886
	}
L2881:
	;
	v13042 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v13043 = m.ExcPending
	if v13043 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2883
	}
L2882:
	;
	switch v13036 - int32(1) {
	case 0:
		goto L2881
	case 1:
		goto L2880
	default:
		goto L2878
	}
L2883:
	;
	if v13042 == int32(0) {
		goto L2878
	} else {
		goto L2884
	}
L2884:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+80)) = v12002 + int32(192)
	F_errmsg(m, int32(_a_F_PostgresMain_214), v12002+int32(80))
	mBase = m.M
	v13053 = m.ExcPending
	if v13053 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2885
	}
L2885:
	;
	v13072 = int32(312)
	goto L2879
L2886:
	;
	if v13057 == int32(0) {
		goto L2878
	} else {
		goto L2887
	}
L2887:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+104)) = v12020
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+100)) = v12060
	*(*int32)(unsafe.Add(mBase, uint32(v12002)+96)) = v12002 + int32(192)
	F_errmsg(m, int32(_a_F_PostgresMain_255), v12002+int32(96))
	mBase = m.M
	v13070 = m.ExcPending
	if v13070 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2888
	}
L2888:
	;
	v13072 = int32(317)
	goto L2879
L2889:
	;
	goto L2878
L2890:
	;
	if v13082 != 0 {
		goto L2891
	} else {
		goto L2892
	}
L2891:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v13085 = m.ExcPending
	if v13085 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2894
	}
L2892:
	;
	goto L2893
L2893:
	;
	v13087 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])))
	if v13087 != 0 {
		goto L2895
	} else {
		goto L2896
	}
L2894:
	;
	goto L2893
L2895:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v13089 = m.ExcPending
	if v13089 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2898
	}
L2896:
	;
	goto L2897
L2897:
	;
	v13094 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[124])) = uint8(v13094)
	goto L690
L2898:
	;
	v13091 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])) = uint8(v13091)
	goto L2897
L2899:
	;
	v13101 = v2548 + int32(440)
	v13102 = F_pq_getmsgbyte(m, v13101)
	mBase = m.M
	v13103 = m.ExcPending
	if v13103 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2900
	}
L2900:
	;
	v13104 = F_pq_getmsgstring(m, v13101)
	mBase = m.M
	v13105 = m.ExcPending
	if v13105 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2901
	}
L2901:
	;
	F_pq_getmsgend(m, v13101)
	mBase = m.M
	v13107 = m.ExcPending
	if v13107 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2902
	}
L2902:
	;
	switch v13102 - int32(80) {
	case 0:
		goto L2904
	default:
		goto L674
	case 3:
		goto L2905
	}
L2903:
	;
	v13132 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	if v13132 != int32(2) {
		goto L690
	} else {
		goto L2915
	}
L2904:
	;
	v13123 = F_GetPortalByName(m, v13104)
	mBase = m.M
	v13124 = m.ExcPending
	if v13124 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2912
	}
L2905:
	;
	v13110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13104))))
	if v13110 != 0 {
		goto L2906
	} else {
		goto L2907
	}
L2906:
	;
	F_DropPreparedStatement(m, v13104, int32(0))
	mBase = m.M
	v13113 = m.ExcPending
	if v13113 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2909
	}
L2907:
	;
	goto L2908
L2908:
	;
	v13115 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[215]))
	if v13115 == int32(0) {
		goto L2903
	} else {
		goto L2910
	}
L2909:
	;
	goto L2903
L2910:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[215])) = int32(0)
	F_DropCachedPlan(m, v13115)
	mBase = m.M
	v13122 = m.ExcPending
	if v13122 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2911
	}
L2911:
	;
	goto L2903
L2912:
	;
	if v13123 == int32(0) {
		goto L2903
	} else {
		goto L2913
	}
L2913:
	;
	F_PortalDrop(m, v13123, int32(0))
	mBase = m.M
	v13129 = m.ExcPending
	if v13129 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2914
	}
L2914:
	;
	goto L2903
L2915:
	;
	F_pq_putemptymessage(m, int32(51))
	mBase = m.M
	v13137 = m.ExcPending
	if v13137 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2916
	}
L2916:
	;
	goto L690
L2917:
	;
	v13143 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[158]))
	if v13143 < int32(0) {
		goto L2919
	} else {
		goto L2920
	}
L2918:
	;
	v13150 = v2548 + int32(440)
	v13151 = F_pq_getmsgbyte(m, v13150)
	mBase = m.M
	v13152 = m.ExcPending
	if v13152 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2922
	}
L2919:
	;
	v13147 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMain[159])) = v13147
	goto L2921
L2920:
	;
	goto L2921
L2921:
	;
	goto L2918
L2922:
	;
	v13153 = F_pq_getmsgstring(m, v13150)
	mBase = m.M
	v13154 = m.ExcPending
	if v13154 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2923
	}
L2923:
	;
	F_pq_getmsgend(m, v13150)
	mBase = m.M
	v13156 = m.ExcPending
	if v13156 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2924
	}
L2924:
	;
	switch v13151 - int32(80) {
	case 0:
		goto L2926
	default:
		goto L2925
	case 3:
		goto L2927
	}
L2925:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13379 = m.ExcPending
	if v13379 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2969
	}
L2926:
	;
	F_start_xact_command(m)
	mBase = m.M
	v13343 = m.ExcPending
	if v13343 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2954
	}
L2927:
	;
	F_start_xact_command(m)
	mBase = m.M
	v13160 = m.ExcPending
	if v13160 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2928
	}
L2928:
	;
	v13163 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[49]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v13163
	v13165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13153))))
	if v13165 != 0 {
		goto L2930
	} else {
		goto L2931
	}
L2929:
	;
	v13176 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v13177 = *(*int32)(unsafe.Add(mBase, uint32(v13176)+24))
	goto L2935
L2930:
	;
	v13167 = F_FetchPreparedStatement(m, v13153, int32(1))
	mBase = m.M
	v13168 = m.ExcPending
	if v13168 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2933
	}
L2931:
	;
	goto L2932
L2932:
	;
	v13171 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[215]))
	if v13171 == int32(0) {
		goto L672
	} else {
		goto L2934
	}
L2933:
	;
	v13169 = *(*int32)(unsafe.Add(mBase, uint32(v13167)+64))
	v13174 = v13169
	goto L2929
L2934:
	;
	v13174 = v13171
	goto L2929
L2935:
	;
	if (v13177-int32(7))&int32(-9) == int32(0) {
		goto L2936
	} else {
		goto L2937
	}
L2936:
	;
	v13184 = *(*int32)(unsafe.Add(mBase, uint32(v13174)+52))
	if v13184 != 0 {
		goto L671
	} else {
		goto L2939
	}
L2937:
	;
	goto L2938
L2938:
	;
	v13186 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	if v13186 != int32(2) {
		goto L690
	} else {
		goto L2940
	}
L2939:
	;
	goto L2938
L2940:
	;
	v13189 = int32(_a_F_PostgresMain_19)
	F_resetStringInfo(m, v13189)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[229])) = int32(116)
	goto L2941
L2941:
	;
	v13193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13174)+24)))
	F_enlargeStringInfo(m, int32(_a_F_PostgresMain_19), int32(2))
	mBase = m.M
	v13197 = m.ExcPending
	if v13197 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2942
	}
L2942:
	;
	v13199 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[230]))
	v13200 = int32(_a_F_PostgresMain_256)
	v13201 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[231]))
	v13203 = int32(8)
	v13207 = v13193<<(uint(v13203)%32) | int32(base.Ui32(v13193)>>(uint(v13203)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v13199+v13201))) = uint16(v13207)
	v13211 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[231]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[231])) = v13211 + int32(2)
	v13215 = *(*int32)(unsafe.Add(mBase, uint32(v13174)+24))
	if int32(0) < v13215 {
		goto L2943
	} else {
		goto L2944
	}
L2943:
	;
	v13223 = int32(0)
	goto L2946
L2944:
	;
	goto L2945
L2945:
	;
	F_pq_endmessage_reuse(m, int32(_a_F_PostgresMain_19))
	mBase = m.M
	v13331 = m.ExcPending
	if v13331 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2950
	}
L2946:
	;
	v13258 = *(*int32)(unsafe.Add(mBase, uint32(v13174)+20))
	v13262 = *(*int32)(unsafe.Add(mBase, uint32(v13258+v13223<<(uint(int32(2))%32))))
	F_enlargeStringInfo(m, int32(_a_F_PostgresMain_19), int32(4))
	mBase = m.M
	v13266 = m.ExcPending
	if v13266 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2948
	}
L2947:
	;
	goto L2945
L2948:
	;
	v13267 = int32(_a_F_PostgresMain_256)
	v13268 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[231]))
	v13270 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[230]))
	v13274 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v13268+v13270))) = base.I32_rotr(v13262, int32(24))&v13274 | base.I32_rotr(v13262&v13274, int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[231])) = v13268 + int32(4)
	v13287 = v13223 + int32(1)
	v13288 = *(*int32)(unsafe.Add(mBase, uint32(v13174)+24))
	if v13287 < v13288 {
		v13223 = v13287
		goto L2946
	} else {
		goto L2949
	}
L2949:
	;
	goto L2947
L2950:
	;
	v13332 = *(*int32)(unsafe.Add(mBase, uint32(v13174)+52))
	if v13332 == int32(0) {
		goto L695
	} else {
		goto L2951
	}
L2951:
	;
	v13335 = F_CachedPlanGetTargetList(m, v13174)
	mBase = m.M
	v13336 = m.ExcPending
	if v13336 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2952
	}
L2952:
	;
	v13338 = *(*int32)(unsafe.Add(mBase, uint32(v13174)+52))
	F_SendRowDescriptionMessage(m, int32(_a_F_PostgresMain_19), v13338, v13335, int32(0))
	mBase = m.M
	v13341 = m.ExcPending
	if v13341 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2953
	}
L2953:
	;
	goto L690
L2954:
	;
	v13346 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[49]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[50])) = v13346
	v13348 = F_GetPortalByName(m, v13153)
	mBase = m.M
	v13349 = m.ExcPending
	if v13349 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2955
	}
L2955:
	;
	if v13348 == int32(0) {
		goto L670
	} else {
		goto L2956
	}
L2956:
	;
	v13353 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v13354 = *(*int32)(unsafe.Add(mBase, uint32(v13353)+24))
	goto L2957
L2957:
	;
	if (v13354-int32(7))&int32(-9) == int32(0) {
		goto L2958
	} else {
		goto L2959
	}
L2958:
	;
	v13361 = *(*int32)(unsafe.Add(mBase, uint32(v13348)+92))
	if v13361 != 0 {
		goto L669
	} else {
		goto L2961
	}
L2959:
	;
	goto L2960
L2960:
	;
	v13363 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	if v13363 != int32(2) {
		goto L690
	} else {
		goto L2962
	}
L2961:
	;
	goto L2960
L2962:
	;
	v13366 = *(*int32)(unsafe.Add(mBase, uint32(v13348)+92))
	if v13366 != 0 {
		goto L2963
	} else {
		goto L2964
	}
L2963:
	;
	v13368 = F_FetchPortalTargetList(m, v13348)
	mBase = m.M
	v13369 = m.ExcPending
	if v13369 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2966
	}
L2964:
	;
	goto L2965
L2965:
	;
	F_pq_putemptymessage(m, int32(110))
	mBase = m.M
	v13375 = m.ExcPending
	if v13375 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2968
	}
L2966:
	;
	v13370 = *(*int32)(unsafe.Add(mBase, uint32(v13348)+96))
	F_SendRowDescriptionMessage(m, int32(_a_F_PostgresMain_19), v13366, v13368, v13370)
	mBase = m.M
	v13372 = m.ExcPending
	if v13372 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2967
	}
L2967:
	;
	goto L690
L2968:
	;
	goto L690
L2969:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v13382 = m.ExcPending
	if v13382 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2970
	}
L2970:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+368)) = v13151
	F_errmsg(m, int32(_a_F_PostgresMain_257), v2548+int32(368))
	mBase = m.M
	v13388 = m.ExcPending
	if v13388 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2971
	}
L2971:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(_a_F_PostgresMain_258), int32(_a_F_PostgresMain_33))
	mBase = m.M
	v13393 = m.ExcPending
	if v13393 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2972
	}
L2972:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2973:
	;
	v13399 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21]))
	if v13399 != int32(2) {
		goto L690
	} else {
		goto L2974
	}
L2974:
	;
	v13403 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[109]))
	v13404 = *(*int32)(unsafe.Add(mBase, uint32(v13403)+4))
	v13405 = m.T0[v13404].(func(*base.Module) int32)(m)
	mBase = m.M
	v13406 = m.ExcPending
	if v13406 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2975
	}
L2975:
	;
	goto L690
L2976:
	;
	v13413 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[116]))
	v13414 = *(*int32)(unsafe.Add(mBase, uint32(v13413)+24))
	if v13414 == int32(4) {
		goto L2978
	} else {
		goto L2979
	}
L2977:
	;
	v13422 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[63])))
	goto L2981
L2978:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13413)+24)) = int32(1)
	goto L2980
L2979:
	;
	goto L2980
L2980:
	;
	goto L2977
L2981:
	;
	if v13422 != 0 {
		goto L2982
	} else {
		goto L2983
	}
L2982:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v13425 = m.ExcPending
	if v13425 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2985
	}
L2983:
	;
	goto L2984
L2984:
	;
	v13427 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])))
	if v13427 != 0 {
		goto L2986
	} else {
		goto L2987
	}
L2985:
	;
	goto L2984
L2986:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v13429 = m.ExcPending
	if v13429 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2989
	}
L2987:
	;
	goto L2988
L2988:
	;
	v13434 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[124])) = uint8(v13434)
	goto L690
L2989:
	;
	v13431 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMain[122])) = uint8(v13431)
	goto L2988
L2990:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[21])) = int32(0)
	goto L2992
L2991:
	;
	goto L2992
L2992:
	;
	v13447 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMain[232]))
	if v13447 != 0 {
		goto L668
	} else {
		goto L2993
	}
L2993:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v13450 = m.ExcPending
	if v13450 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2994
	}
L2994:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2995:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v13457 = m.ExcPending
	if v13457 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2996
	}
L2996:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+16)) = v3374
	F_errmsg(m, int32(_a_F_PostgresMain_45), v2548+int32(16))
	mBase = m.M
	v13463 = m.ExcPending
	if v13463 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2997
	}
L2997:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(_a_F_PostgresMain_259), int32(_a_F_PostgresMain_33))
	mBase = m.M
	v13468 = m.ExcPending
	if v13468 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L2998
	}
L2998:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2999:
	;
	goto L690
L3000:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v13520 = m.ExcPending
	if v13520 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3001
	}
L3001:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_260), int32(0))
	mBase = m.M
	v13524 = m.ExcPending
	if v13524 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3002
	}
L3002:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(_a_F_PostgresMain_261), int32(_a_F_PostgresMain_262))
	mBase = m.M
	v13529 = m.ExcPending
	if v13529 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3003
	}
L3003:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3004:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v13536 = m.ExcPending
	if v13536 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3005
	}
L3005:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_263), int32(0))
	mBase = m.M
	v13540 = m.ExcPending
	if v13540 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3006
	}
L3006:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(1599), int32(_a_F_PostgresMain_219))
	mBase = m.M
	v13545 = m.ExcPending
	if v13545 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3007
	}
L3007:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3008:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v13552 = m.ExcPending
	if v13552 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3009
	}
L3009:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_181), int32(0))
	mBase = m.M
	v13556 = m.ExcPending
	if v13556 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3010
	}
L3010:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(1620), int32(_a_F_PostgresMain_219))
	mBase = m.M
	v13561 = m.ExcPending
	if v13561 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3011
	}
L3011:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3012:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v13568 = m.ExcPending
	if v13568 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3013
	}
L3013:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_260), int32(0))
	mBase = m.M
	v13572 = m.ExcPending
	if v13572 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3014
	}
L3014:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(_a_F_PostgresMain_261), int32(_a_F_PostgresMain_262))
	mBase = m.M
	v13577 = m.ExcPending
	if v13577 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3015
	}
L3015:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3016:
	;
	F_errcode(m, int32(386))
	mBase = m.M
	v13584 = m.ExcPending
	if v13584 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3017
	}
L3017:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_264), int32(0))
	mBase = m.M
	v13588 = m.ExcPending
	if v13588 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3018
	}
L3018:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(1795), int32(_a_F_PostgresMain_224))
	mBase = m.M
	v13593 = m.ExcPending
	if v13593 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3019
	}
L3019:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3020:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v13600 = m.ExcPending
	if v13600 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3021
	}
L3021:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+196)) = v10418
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+192)) = v10317
	F_errmsg(m, int32(_a_F_PostgresMain_265), v2548+int32(192))
	mBase = m.M
	v13607 = m.ExcPending
	if v13607 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3022
	}
L3022:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(1848), int32(_a_F_PostgresMain_224))
	mBase = m.M
	v13612 = m.ExcPending
	if v13612 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3023
	}
L3023:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3024:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v13619 = m.ExcPending
	if v13619 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3025
	}
L3025:
	;
	v13620 = *(*int32)(unsafe.Add(mBase, uint32(v10144)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+184)) = v13620
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+180)) = v10110
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+176)) = v10418
	F_errmsg(m, int32(_a_F_PostgresMain_266), v2548+int32(176))
	mBase = m.M
	v13628 = m.ExcPending
	if v13628 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3026
	}
L3026:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(1854), int32(_a_F_PostgresMain_224))
	mBase = m.M
	v13633 = m.ExcPending
	if v13633 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3027
	}
L3027:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3028:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v13641 = m.ExcPending
	if v13641 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3029
	}
L3029:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_181), int32(0))
	mBase = m.M
	v13645 = m.ExcPending
	if v13645 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3030
	}
L3030:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(1871), int32(_a_F_PostgresMain_224))
	mBase = m.M
	v13650 = m.ExcPending
	if v13650 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3031
	}
L3031:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3032:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v13657 = m.ExcPending
	if v13657 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3033
	}
L3033:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+160)) = v10527 + int32(1)
	F_errmsg(m, int32(_a_F_PostgresMain_267), v2548+int32(160))
	mBase = m.M
	v13665 = m.ExcPending
	if v13665 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3034
	}
L3034:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(2067), int32(_a_F_PostgresMain_224))
	mBase = m.M
	v13670 = m.ExcPending
	if v13670 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3035
	}
L3035:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3036:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v13677 = m.ExcPending
	if v13677 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3037
	}
L3037:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+80)) = base.I32_extend16_s(v10600)
	F_errmsg(m, int32(_a_F_PostgresMain_249), v2548+int32(80))
	mBase = m.M
	v13684 = m.ExcPending
	if v13684 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3038
	}
L3038:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(2074), int32(_a_F_PostgresMain_224))
	mBase = m.M
	v13689 = m.ExcPending
	if v13689 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3039
	}
L3039:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3040:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v13696 = m.ExcPending
	if v13696 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3041
	}
L3041:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_260), int32(0))
	mBase = m.M
	v13700 = m.ExcPending
	if v13700 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3042
	}
L3042:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(_a_F_PostgresMain_261), int32(_a_F_PostgresMain_262))
	mBase = m.M
	v13705 = m.ExcPending
	if v13705 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3043
	}
L3043:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3044:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v13712 = m.ExcPending
	if v13712 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3045
	}
L3045:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+224)) = v11204
	F_errmsg(m, int32(_a_F_PostgresMain_268), v2548+int32(224))
	mBase = m.M
	v13718 = m.ExcPending
	if v13718 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3046
	}
L3046:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(2260), int32(_a_F_PostgresMain_233))
	mBase = m.M
	v13723 = m.ExcPending
	if v13723 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3047
	}
L3047:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3048:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v13731 = m.ExcPending
	if v13731 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3049
	}
L3049:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_181), int32(0))
	mBase = m.M
	v13735 = m.ExcPending
	if v13735 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3050
	}
L3050:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(2375), int32(_a_F_PostgresMain_233))
	mBase = m.M
	v13740 = m.ExcPending
	if v13740 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3051
	}
L3051:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3052:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v13747 = m.ExcPending
	if v13747 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3053
	}
L3053:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_269), int32(0))
	mBase = m.M
	v13751 = m.ExcPending
	if v13751 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3054
	}
L3054:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(_a_F_PostgresMain_270), int32(_a_F_PostgresMain_262))
	mBase = m.M
	v13756 = m.ExcPending
	if v13756 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3055
	}
L3055:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3056:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v13763 = m.ExcPending
	if v13763 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3057
	}
L3057:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_260), int32(0))
	mBase = m.M
	v13767 = m.ExcPending
	if v13767 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3058
	}
L3058:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(_a_F_PostgresMain_261), int32(_a_F_PostgresMain_262))
	mBase = m.M
	v13772 = m.ExcPending
	if v13772 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3059
	}
L3059:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3060:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v13779 = m.ExcPending
	if v13779 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3061
	}
L3061:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+352)) = v13102
	F_errmsg(m, int32(_a_F_PostgresMain_271), v2548+int32(352))
	mBase = m.M
	v13785 = m.ExcPending
	if v13785 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3062
	}
L3062:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(_a_F_PostgresMain_272), int32(_a_F_PostgresMain_33))
	mBase = m.M
	v13790 = m.ExcPending
	if v13790 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3063
	}
L3063:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3064:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v13797 = m.ExcPending
	if v13797 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3065
	}
L3065:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_260), int32(0))
	mBase = m.M
	v13801 = m.ExcPending
	if v13801 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3066
	}
L3066:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(_a_F_PostgresMain_261), int32(_a_F_PostgresMain_262))
	mBase = m.M
	v13806 = m.ExcPending
	if v13806 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3067
	}
L3067:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3068:
	;
	F_errcode(m, int32(386))
	mBase = m.M
	v13813 = m.ExcPending
	if v13813 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3069
	}
L3069:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_264), int32(0))
	mBase = m.M
	v13817 = m.ExcPending
	if v13817 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3070
	}
L3070:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(2777), int32(_a_F_PostgresMain_273))
	mBase = m.M
	v13822 = m.ExcPending
	if v13822 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3071
	}
L3071:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3072:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v13829 = m.ExcPending
	if v13829 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3073
	}
L3073:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_181), int32(0))
	mBase = m.M
	v13833 = m.ExcPending
	if v13833 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3074
	}
L3074:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(2797), int32(_a_F_PostgresMain_273))
	mBase = m.M
	v13838 = m.ExcPending
	if v13838 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3075
	}
L3075:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3076:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v13845 = m.ExcPending
	if v13845 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3077
	}
L3077:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2548)+384)) = v13153
	F_errmsg(m, int32(_a_F_PostgresMain_268), v2548+int32(384))
	mBase = m.M
	v13851 = m.ExcPending
	if v13851 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3078
	}
L3078:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(2858), int32(_a_F_PostgresMain_274))
	mBase = m.M
	v13856 = m.ExcPending
	if v13856 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3079
	}
L3079:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3080:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v13863 = m.ExcPending
	if v13863 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3081
	}
L3081:
	;
	F_errmsg(m, int32(_a_F_PostgresMain_181), int32(0))
	mBase = m.M
	v13867 = m.ExcPending
	if v13867 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3082
	}
L3082:
	;
	F_errfinish(m, int32(_a_F_PostgresMain_6), int32(2873), int32(_a_F_PostgresMain_274))
	mBase = m.M
	v13872 = m.ExcPending
	if v13872 != 0 {
		v13879 = v45
		v13880 = v46
		v13911 = v77
		goto L5
	} else {
		goto L3083
	}
L3083:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3084:
	;
	v13923 = int32(v13919)
	m.G0 = v13911
	v13925 = *(*int32)(unsafe.Add(mBase, uint32(v13923)+4))
	v13926 = *(*int32)(unsafe.Add(mBase, uint32(v13923)))
	v13929 = *(*int32)(unsafe.Add(mBase, uint32(v13926)))
	if v13911+int32(12) == v13929 {
		goto L3087
	} else {
		goto L3088
	}
L3085:
	;
	m.ExcPending = 1
	goto L3093
L3086:
	;
	if v13933 == int32(0) {
		goto L3090
	} else {
		goto L3091
	}
L3087:
	;
	v13931 = *(*int32)(unsafe.Add(mBase, uint32(v13926)+4))
	v13933 = v13931
	goto L3089
L3088:
	;
	v13933 = int32(0)
	goto L3089
L3089:
	;
	goto L3086
L3090:
	;
	F___wasm_longjmp(m, v13926, v13925)
	mBase = m.M
	v13937 = m.ExcPending
	if v13937 != 0 {
		goto L3093
	} else {
		goto L3094
	}
L3091:
	;
	goto L3092
L3092:
	;
	v45 = v13879
	v46 = v13880
	v47 = v13933
	v48 = v13925
	v77 = v13911
	goto L1
L3093:
	;
	return
L3094:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
