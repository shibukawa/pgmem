package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_perform_base_backup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v36 int64
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v352 int32
	_ = v352
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v410 int32
	_ = v410
	var v411 int64
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v482 int64
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v687 int32
	_ = v687
	var v712 int32
	_ = v712
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v758 int32
	_ = v758
	var v761 int64
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int64
	_ = v764
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v788 int32
	_ = v788
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int64
	_ = v819
	var v821 int32
	_ = v821
	var v822 int64
	_ = v822
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int64
	_ = v834
	var v835 int32
	_ = v835
	var v836 int64
	_ = v836
	var v839 int64
	_ = v839
	var v840 int64
	_ = v840
	var v844 int64
	_ = v844
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v857 int64
	_ = v857
	var v859 int64
	_ = v859
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int64
	_ = v869
	var v872 int64
	_ = v872
	var v876 int64
	_ = v876
	var v877 int64
	_ = v877
	var v880 int64
	_ = v880
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v896 int64
	_ = v896
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int64
	_ = v909
	var v910 int32
	_ = v910
	var v911 int64
	_ = v911
	var v914 int64
	_ = v914
	var v915 int64
	_ = v915
	var v919 int64
	_ = v919
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int64
	_ = v939
	var v943 int32
	_ = v943
	var v944 int64
	_ = v944
	var v947 int64
	_ = v947
	var v948 int64
	_ = v948
	var v952 int64
	_ = v952
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v997 int32
	_ = v997
	var v1021 int64
	_ = v1021
	var v1023 int64
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1028 int64
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1088 int64
	_ = v1088
	var v1089 int64
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1095 int64
	_ = v1095
	var v1099 int64
	_ = v1099
	var v1101 int64
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1107 int32
	_ = v1107
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1161 int64
	_ = v1161
	var v1165 int64
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1197 int32
	_ = v1197
	var v1258 int32
	_ = v1258
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1298 int32
	_ = v1298
	var v1317 int64
	_ = v1317
	var v1325 int32
	_ = v1325
	var v1326 int64
	_ = v1326
	var v1328 int64
	_ = v1328
	var v1332 int64
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1373 int64
	_ = v1373
	var v1422 int32
	_ = v1422
	var v1425 int64
	_ = v1425
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1433 int64
	_ = v1433
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1453 int32
	_ = v1453
	var v1454 int64
	_ = v1454
	var v1457 int64
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1493 int32
	_ = v1493
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1540 int32
	_ = v1540
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1588 int32
	_ = v1588
	var v1614 int32
	_ = v1614
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1640 int32
	_ = v1640
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1736 int32
	_ = v1736
	var v1777 int32
	_ = v1777
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1979 int32
	_ = v1979
	var v1984 int32
	_ = v1984
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2079 int32
	_ = v2079
	var v2083 int32
	_ = v2083
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2094 int32
	_ = v2094
	var v2098 int32
	_ = v2098
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2116 int32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2131 int32
	_ = v2131
	var v2162 int64
	_ = v2162
	var v2163 int64
	_ = v2163
	var v2169 int32
	_ = v2169
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2190 int32
	_ = v2190
	var v2192 int32
	_ = v2192
	var v2194 int32
	_ = v2194
	var v2197 int64
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2213 int32
	_ = v2213
	var v2215 int64
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2218 int64
	_ = v2218
	var v2219 int64
	_ = v2219
	var v2220 int64
	_ = v2220
	var v2222 int64
	_ = v2222
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2264 int64
	_ = v2264
	var v2265 int64
	_ = v2265
	var v2271 int32
	_ = v2271
	var v2308 int64
	_ = v2308
	var v2309 int64
	_ = v2309
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2360 int32
	_ = v2360
	var v2385 int32
	_ = v2385
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2409 int32
	_ = v2409
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2439 int32
	_ = v2439
	var v2441 int32
	_ = v2441
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2472 int32
	_ = v2472
	var v2474 int32
	_ = v2474
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2496 int64
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2515 int32
	_ = v2515
	var v2531 int32
	_ = v2531
	var v2545 int32
	_ = v2545
	var v2564 int32
	_ = v2564
	var v2581 int32
	_ = v2581
	var v2594 int32
	_ = v2594
	var v2598 int32
	_ = v2598
	var v2608 int32
	_ = v2608
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2646 int32
	_ = v2646
	var v2662 int64
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2667 int32
	_ = v2667
	var v2670 int32
	_ = v2670
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2695 int32
	_ = v2695
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2734 int32
	_ = v2734
	var v2750 int32
	_ = v2750
	var v2764 int32
	_ = v2764
	var v2819 int32
	_ = v2819
	var v2821 int32
	_ = v2821
	var v2829 int32
	_ = v2829
	var v2830 int64
	_ = v2830
	var v2832 int64
	_ = v2832
	var v2846 int32
	_ = v2846
	var v2850 int32
	_ = v2850
	var v2855 int32
	_ = v2855
	var v2857 int32
	_ = v2857
	var v2858 int32
	_ = v2858
	var v2861 int32
	_ = v2861
	var v2865 int32
	_ = v2865
	var v2868 int32
	_ = v2868
	var v2960 int32
	_ = v2960
	var v2963 int32
	_ = v2963
	var v2972 int32
	_ = v2972
	var v2973 int32
	_ = v2973
	var v2979 int64
	_ = v2979
	var v2981 int32
	_ = v2981
	var v2984 int32
	_ = v2984
	var v2995 int32
	_ = v2995
	var v2998 int32
	_ = v2998
	var v2999 int32
	_ = v2999
	var v3000 int32
	_ = v3000
	var v3003 int32
	_ = v3003
	var v3005 int32
	_ = v3005
	var v3021 int32
	_ = v3021
	var v3039 int32
	_ = v3039
	var v3040 int32
	_ = v3040
	var v3041 int64
	_ = v3041
	var v3054 int32
	_ = v3054
	var v3056 int32
	_ = v3056
	var v3070 int32
	_ = v3070
	var v3086 int32
	_ = v3086
	var v3091 int32
	_ = v3091
	var v3108 int32
	_ = v3108
	var v3112 int32
	_ = v3112
	var v3117 int32
	_ = v3117
	var v3119 int32
	_ = v3119
	var v3120 int32
	_ = v3120
	var v3123 int32
	_ = v3123
	var v3127 int32
	_ = v3127
	var v3129 int32
	_ = v3129
	var v3130 int32
	_ = v3130
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3145 int32
	_ = v3145
	var v3149 int32
	_ = v3149
	var v3150 int64
	_ = v3150
	var v3164 int64
	_ = v3164
	var v3166 int64
	_ = v3166
	var v3168 int64
	_ = v3168
	var v3169 int64
	_ = v3169
	var v3172 int64
	_ = v3172
	var v3180 int32
	_ = v3180
	var v3181 int32
	_ = v3181
	var v3195 int64
	_ = v3195
	var v3199 int64
	_ = v3199
	var v3201 int64
	_ = v3201
	var v3202 int64
	_ = v3202
	var v3205 int64
	_ = v3205
	var v3213 int32
	_ = v3213
	var v3214 int32
	_ = v3214
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3247 int32
	_ = v3247
	var v3251 int32
	_ = v3251
	var v3252 int32
	_ = v3252
	var v3256 int32
	_ = v3256
	var v3257 int32
	_ = v3257
	var v3263 int32
	_ = v3263
	var v3270 int32
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3274 int32
	_ = v3274
	var v3275 int32
	_ = v3275
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3330 int32
	_ = v3330
	var v3334 int32
	_ = v3334
	var v3336 int32
	_ = v3336
	var v3337 int64
	_ = v3337
	var v3345 int32
	_ = v3345
	var v3349 int32
	_ = v3349
	var v3353 int32
	_ = v3353
	var v3359 int32
	_ = v3359
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3377 int32
	_ = v3377
	var v3380 int32
	_ = v3380
	var v3384 int32
	_ = v3384
	var v3385 int32
	_ = v3385
	var v3393 int32
	_ = v3393
	var v3399 int32
	_ = v3399
	var v3401 int32
	_ = v3401
	var v3403 int32
	_ = v3403
	var v3413 int32
	_ = v3413
	var v3429 int32
	_ = v3429
	var v3432 int32
	_ = v3432
	var v3435 int32
	_ = v3435
	var v3438 int32
	_ = v3438
	var v3439 int32
	_ = v3439
	var v3442 int32
	_ = v3442
	var v3443 int32
	_ = v3443
	var v3446 int32
	_ = v3446
	var v3453 int32
	_ = v3453
	var v3454 int32
	_ = v3454
	var v3472 int32
	_ = v3472
	var v3475 int32
	_ = v3475
	var v3478 int32
	_ = v3478
	var v3479 int32
	_ = v3479
	var v3482 int32
	_ = v3482
	var v3483 int32
	_ = v3483
	var v3486 int32
	_ = v3486
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3510 int32
	_ = v3510
	var v3511 int32
	_ = v3511
	var v3524 int32
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3538 int32
	_ = v3538
	var v3542 int32
	_ = v3542
	var v3544 int32
	_ = v3544
	var v3545 int64
	_ = v3545
	var v3553 int32
	_ = v3553
	var v3557 int32
	_ = v3557
	var v3561 int32
	_ = v3561
	var v3567 int32
	_ = v3567
	var v3571 int32
	_ = v3571
	var v3572 int32
	_ = v3572
	var v3579 int32
	_ = v3579
	var v3580 int32
	_ = v3580
	var v3581 int32
	_ = v3581
	var v3585 int32
	_ = v3585
	var v3588 int32
	_ = v3588
	var v3592 int32
	_ = v3592
	var v3593 int32
	_ = v3593
	var v3601 int32
	_ = v3601
	var v3607 int32
	_ = v3607
	var v3609 int32
	_ = v3609
	var v3611 int32
	_ = v3611
	var v3621 int32
	_ = v3621
	var v3637 int32
	_ = v3637
	var v3638 int32
	_ = v3638
	var v3641 int32
	_ = v3641
	var v3644 int32
	_ = v3644
	var v3647 int32
	_ = v3647
	var v3648 int32
	_ = v3648
	var v3651 int32
	_ = v3651
	var v3652 int32
	_ = v3652
	var v3655 int32
	_ = v3655
	var v3662 int32
	_ = v3662
	var v3663 int32
	_ = v3663
	var v3677 int32
	_ = v3677
	var v3678 int32
	_ = v3678
	var v3691 int32
	_ = v3691
	var v3692 int32
	_ = v3692
	var v3694 int32
	_ = v3694
	var v3695 int32
	_ = v3695
	var v3696 int32
	_ = v3696
	var v3697 int32
	_ = v3697
	var v3711 int32
	_ = v3711
	var v3712 int32
	_ = v3712
	var v3724 int32
	_ = v3724
	var v3725 int32
	_ = v3725
	var v3726 int32
	_ = v3726
	var v3728 int32
	_ = v3728
	var v3729 int32
	_ = v3729
	var v3768 int32
	_ = v3768
	var v3769 int32
	_ = v3769
	var v3783 int32
	_ = v3783
	var v3798 int32
	_ = v3798
	var v3816 int32
	_ = v3816
	var v3832 int32
	_ = v3832
	var v3849 int32
	_ = v3849
	var v3850 int32
	_ = v3850
	var v3851 int32
	_ = v3851
	var v3865 int64
	_ = v3865
	var v3878 int32
	_ = v3878
	var v3879 int32
	_ = v3879
	var v3880 int64
	_ = v3880
	var v3881 int64
	_ = v3881
	var v3883 int64
	_ = v3883
	var v3887 int32
	_ = v3887
	var v3888 int32
	_ = v3888
	var v3892 int32
	_ = v3892
	var v3906 int32
	_ = v3906
	var v3908 int32
	_ = v3908
	var v3910 int32
	_ = v3910
	var v3926 int32
	_ = v3926
	var v3944 int32
	_ = v3944
	var v3961 int32
	_ = v3961
	var v3966 int32
	_ = v3966
	var v4000 int64
	_ = v4000
	var v4004 int32
	_ = v4004
	var v4008 int32
	_ = v4008
	var v4022 int64
	_ = v4022
	var v4035 int32
	_ = v4035
	var v4036 int32
	_ = v4036
	var v4038 int64
	_ = v4038
	var v4039 int64
	_ = v4039
	var v4040 int64
	_ = v4040
	var v4042 int64
	_ = v4042
	var v4044 int64
	_ = v4044
	var v4051 int32
	_ = v4051
	var v4052 int32
	_ = v4052
	var v4067 int32
	_ = v4067
	var v4068 int32
	_ = v4068
	var v4070 int32
	_ = v4070
	var v4072 int32
	_ = v4072
	var v4088 int32
	_ = v4088
	var v4106 int32
	_ = v4106
	var v4123 int32
	_ = v4123
	var v4180 int32
	_ = v4180
	var v4182 int32
	_ = v4182
	var v4184 int32
	_ = v4184
	var v4200 int32
	_ = v4200
	var v4218 int32
	_ = v4218
	var v4235 int32
	_ = v4235
	var v4257 int32
	_ = v4257
	var v4281 int32
	_ = v4281
	var v4285 int32
	_ = v4285
	var v4300 int32
	_ = v4300
	var v4305 int32
	_ = v4305
	var v4306 int32
	_ = v4306
	var v4320 int64
	_ = v4320
	var v4333 int32
	_ = v4333
	var v4334 int32
	_ = v4334
	var v4348 int64
	_ = v4348
	var v4349 int64
	_ = v4349
	var v4350 int64
	_ = v4350
	var v4352 int32
	_ = v4352
	var v4353 int32
	_ = v4353
	var v4355 int64
	_ = v4355
	var v4371 int32
	_ = v4371
	var v4384 int32
	_ = v4384
	var v4386 int32
	_ = v4386
	var v4404 int32
	_ = v4404
	var v4418 int32
	_ = v4418
	var v4434 int32
	_ = v4434
	var v4451 int32
	_ = v4451
	var v4469 int32
	_ = v4469
	var v4472 int32
	_ = v4472
	var v4473 int32
	_ = v4473
	var v4489 int32
	_ = v4489
	var v4503 int32
	_ = v4503
	var v4523 int32
	_ = v4523
	var v4540 int32
	_ = v4540
	var v4541 int64
	_ = v4541
	var v4543 int64
	_ = v4543
	var v4557 int32
	_ = v4557
	var v4559 int32
	_ = v4559
	var v4575 int32
	_ = v4575
	var v4589 int32
	_ = v4589
	var v4607 int32
	_ = v4607
	var v4624 int32
	_ = v4624
	var v4639 int32
	_ = v4639
	var v4644 int32
	_ = v4644
	var v4646 int32
	_ = v4646
	var v4652 int32
	_ = v4652
	var v4686 int64
	_ = v4686
	var v4690 int32
	_ = v4690
	var v4691 int64
	_ = v4691
	var v4693 int32
	_ = v4693
	var v4709 int64
	_ = v4709
	var v4711 int64
	_ = v4711
	var v4713 int32
	_ = v4713
	var v4715 int32
	_ = v4715
	var v4716 int32
	_ = v4716
	var v4735 int32
	_ = v4735
	var v4749 int32
	_ = v4749
	var v4769 int32
	_ = v4769
	var v4786 int32
	_ = v4786
	var v4799 int32
	_ = v4799
	var v4801 int32
	_ = v4801
	var v4802 int32
	_ = v4802
	var v4803 int32
	_ = v4803
	var v4817 int32
	_ = v4817
	var v4819 int64
	_ = v4819
	var v4821 int32
	_ = v4821
	var v4825 int64
	_ = v4825
	var v4839 int32
	_ = v4839
	var v4841 int32
	_ = v4841
	var v4857 int32
	_ = v4857
	var v4871 int32
	_ = v4871
	var v4889 int32
	_ = v4889
	var v4906 int32
	_ = v4906
	var v4922 int32
	_ = v4922
	var v4923 int32
	_ = v4923
	var v4940 int32
	_ = v4940
	var v4945 int32
	_ = v4945
	var v4946 int32
	_ = v4946
	var v4963 int32
	_ = v4963
	var v4965 int32
	_ = v4965
	var v4966 int32
	_ = v4966
	var v5012 int32
	_ = v5012
	var v5013 int32
	_ = v5013
	var v5020 int32
	_ = v5020
	var v5058 int32
	_ = v5058
	var v5062 int32
	_ = v5062
	var v5077 int32
	_ = v5077
	var v5082 int32
	_ = v5082
	var v5083 int32
	_ = v5083
	var v5100 int32
	_ = v5100
	var v5116 int32
	_ = v5116
	var v5130 int32
	_ = v5130
	var v5148 int32
	_ = v5148
	var v5165 int32
	_ = v5165
	var v5179 int32
	_ = v5179
	var v5182 int32
	_ = v5182
	var v5188 int32
	_ = v5188
	var v5192 int32
	_ = v5192
	var v5193 int32
	_ = v5193
	var v5213 int32
	_ = v5213
	var v5214 int32
	_ = v5214
	var v5229 int32
	_ = v5229
	var v5231 int32
	_ = v5231
	var v5232 int32
	_ = v5232
	var v5276 int32
	_ = v5276
	var v5278 int32
	_ = v5278
	var v5280 int32
	_ = v5280
	var v5281 int32
	_ = v5281
	var v5296 int32
	_ = v5296
	var v5297 int32
	_ = v5297
	var v5298 int32
	_ = v5298
	var v5312 int32
	_ = v5312
	var v5324 int32
	_ = v5324
	var v5325 int32
	_ = v5325
	var v5326 int32
	_ = v5326
	var v5355 int32
	_ = v5355
	var v5356 int64
	_ = v5356
	var v5370 int32
	_ = v5370
	var v5372 int32
	_ = v5372
	var v5375 int32
	_ = v5375
	var v5376 int32
	_ = v5376
	var v5379 int32
	_ = v5379
	var v5380 int32
	_ = v5380
	var v5381 int32
	_ = v5381
	var v5384 int32
	_ = v5384
	var v5388 int32
	_ = v5388
	var v5412 int32
	_ = v5412
	var v5413 int32
	_ = v5413
	var v5432 int64
	_ = v5432
	var v5437 int32
	_ = v5437
	var v5441 int32
	_ = v5441
	var v5442 int64
	_ = v5442
	var v5449 int32
	_ = v5449
	var v5453 int64
	_ = v5453
	var v5456 int64
	_ = v5456
	var v5458 int64
	_ = v5458
	var v5459 int64
	_ = v5459
	var v5464 int64
	_ = v5464
	var v5470 int32
	_ = v5470
	var v5475 int32
	_ = v5475
	var v5476 int32
	_ = v5476
	var v5478 int32
	_ = v5478
	var v5480 int32
	_ = v5480
	var v5481 int32
	_ = v5481
	var v5484 int64
	_ = v5484
	var v5485 int32
	_ = v5485
	var v5487 int64
	_ = v5487
	var v5490 int32
	_ = v5490
	var v5491 int32
	_ = v5491
	var v5538 int32
	_ = v5538
	var v5543 int32
	_ = v5543
	var v5548 int32
	_ = v5548
	var v5551 int32
	_ = v5551
	var v5600 int32
	_ = v5600
	var v5601 int32
	_ = v5601
	var v5608 int32
	_ = v5608
	var v5613 int32
	_ = v5613
	var v5617 int32
	_ = v5617
	var v5618 int32
	_ = v5618
	var v5625 int32
	_ = v5625
	var v5630 int32
	_ = v5630
	var v5645 int32
	_ = v5645
	var v5647 int32
	_ = v5647
	var v5650 int32
	_ = v5650
	var v5651 int32
	_ = v5651
	var v5652 int32
	_ = v5652
	var v5654 int32
	_ = v5654
	var v5656 int32
	_ = v5656
	var v5658 int32
	_ = v5658
	var v5663 int32
	_ = v5663
	var v5666 int32
	_ = v5666
	var v5669 int32
	_ = v5669
	var v5671 int32
	_ = v5671
	var v5673 int32
	_ = v5673
	var v5674 int32
	_ = v5674
	var v5676 int32
	_ = v5676
	var v5681 int32
	_ = v5681
	var v5690 int32
	_ = v5690
	var v5693 int32
	_ = v5693
	var v5696 int32
	_ = v5696
	var v5697 int32
	_ = v5697
	var v5698 int32
	_ = v5698
	var v5701 int32
	_ = v5701
	var v5702 int32
	_ = v5702
	var v5703 int32
	_ = v5703
	var v5704 int32
	_ = v5704
	var v5706 int32
	_ = v5706
	var v5707 int64
	_ = v5707
	var v5727 int32
	_ = v5727
	var v5748 int64
	_ = v5748
	var v5749 int64
	_ = v5749
	var v5752 int32
	_ = v5752
	var v5753 int32
	_ = v5753
	var v5754 int64
	_ = v5754
	var v5755 int64
	_ = v5755
	var v5757 int64
	_ = v5757
	var v5758 int32
	_ = v5758
	var v5760 int32
	_ = v5760
	var v5761 int32
	_ = v5761
	var v5762 int32
	_ = v5762
	var v5764 int32
	_ = v5764
	var v5765 int64
	_ = v5765
	var v5766 int32
	_ = v5766
	var v5767 int64
	_ = v5767
	var v5811 int32
	_ = v5811
	var v5812 int32
	_ = v5812
	var v5814 int32
	_ = v5814
	var v5815 int32
	_ = v5815
	var v5817 int32
	_ = v5817
	var v5866 int32
	_ = v5866
	var v5867 int32
	_ = v5867
	var v5874 int32
	_ = v5874
	var v5877 int32
	_ = v5877
	var v5880 int32
	_ = v5880
	var v5882 int32
	_ = v5882
	var v5886 int32
	_ = v5886
	var v5891 int32
	_ = v5891
	var v5895 int32
	_ = v5895
	var v5897 int32
	_ = v5897
	var v5901 int32
	_ = v5901
	var v5906 int32
	_ = v5906
	var v5907 int32
	_ = v5907
	var v5908 int32
	_ = v5908
	var v5922 int32
	_ = v5922
	var v5924 int64
	_ = v5924
	var v5941 int32
	_ = v5941
	var v5942 int32
	_ = v5942
	var v5956 int64
	_ = v5956
	var v5964 int32
	_ = v5964
	var v5981 int32
	_ = v5981
	var v5998 int32
	_ = v5998
	var v6013 int32
	_ = v6013
	var v6029 int32
	_ = v6029
	var v6046 int32
	_ = v6046
	var v6060 int32
	_ = v6060
	var v6061 int32
	_ = v6061
	var v6063 int32
	_ = v6063
	var v6080 int32
	_ = v6080
	var v6081 int32
	_ = v6081
	var v6082 int32
	_ = v6082
	var v6083 int32
	_ = v6083
	var v6084 int32
	_ = v6084
	var v6107 int32
	_ = v6107
	var v6111 int32
	_ = v6111
	var v6112 int32
	_ = v6112
	var v6123 int32
	_ = v6123
	var v6124 int64
	_ = v6124
	var v6128 int32
	_ = v6128
	var v6130 int32
	_ = v6130
	var v6131 int32
	_ = v6131
	var v6134 int32
	_ = v6134
	var v6136 int32
	_ = v6136
	var v6138 int32
	_ = v6138
	var v6139 int32
	_ = v6139
	var v6140 int32
	_ = v6140
	var v6141 int32
	_ = v6141
	var v6142 int32
	_ = v6142
	var v6143 int32
	_ = v6143
	var v6144 int32
	_ = v6144
	var v6145 int32
	_ = v6145
	var v6146 int64
	_ = v6146
	var v6147 int64
	_ = v6147
	var v6148 int32
	_ = v6148
	var v6149 int32
	_ = v6149
	var v6150 int32
	_ = v6150
	var v6152 int32
	_ = v6152
	v4 = int32(0)
	v36 = int64(0)
	v43 = m.G0
	v45 = v43 - int32(2240)
	m.G0 = v45
	v54 = l0
	v55 = l1
	v56 = l2
	v57 = v45
	v58 = v4
	v59 = v4
	v60 = v4
	v61 = v4
	v62 = v4
	v63 = v4
	v64 = v4
	v65 = v4
	v66 = v4
	v67 = v4
	v68 = v4
	v71 = int32(-1)
	v80 = v45 + int32(2120)
	v84 = v45 + int32(2144)
	v85 = v45 + int32(2136)
	v89 = v36
	v90 = v36
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L4
L3:
	;
	m.G0 = v57 + int32(2240)
	return
L4:
	;
	if v71 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v6123 = int32(m.ExcTag)
	v6124 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v6123 == int32(0) {
		goto L663
	} else {
		goto L664
	}
L7:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[1])) = v100
	v102 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2104)) = v102
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2112)) = v102
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2120)) = v102
	v108 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+2128)) = uint8(v108)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v60
	v121 = v57 + int32(2128)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v80
	v127 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[2])))
	if v127 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v382 = v58
	v383 = v59
	v384 = v60
	v385 = v61
	v386 = v62
	v387 = v63
	v388 = v64
	v389 = v68
	goto L9
L9:
	;
	if v382 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[3])) = uint8(v137)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v54)+56))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v80
	v153 = m.G0
	v155 = v153 - int32(32)
	m.G0 = v155
	v158 = v57 + int32(2072)
	v159 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v158))) = v159
	*(*int64)(unsafe.Add(mBase, uint32(v158)+24)) = v159
	*(*int64)(unsafe.Add(mBase, uint32(v158)+16)) = v159
	*(*int64)(unsafe.Add(mBase, uint32(v158)+8)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v158)+4)) = v139
	if v140 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[4]))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+308))
	v135 = base.B2i32(v133 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[2])) = uint8(v135)
	v137 = v135
	goto L13
L12:
	;
	v137 = v108
	goto L13
L13:
	;
	goto L10
L14:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_perform_base_backup[5])) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v80
	v257 = F_palloc0(m, int32(1112))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L40
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L27
	}
L16:
	;
	m.G0 = v155 + int32(32)
	goto L14
L17:
	;
	v170 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v158)+26)) = uint8(v170)
	v172 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v158)+24)) = uint16(v172)
	*(*int64)(unsafe.Add(mBase, uint32(v158)+16)) = int64(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v177 = F_BufFileCreateTemp(m, int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v177
	v181 = F_pg_cryptohash_create(m, int32(3))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+8)) = v181
	v184 = F_pg_cryptohash_init(m, v181)
	mBase = m.M
	if v184 < int32(0) {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v187 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v158)+25)) = uint16(v187)
	*(*int64)(unsafe.Add(mBase, uint32(v158)+16)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v158)+24)) = uint8(base.B2i32(v140 == int32(2)))
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[6]))
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v195)))
	goto L23
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v155)+16)) = v196
	v201 = F_psprintf(m, int32(_a_F_perform_base_backup_0), v155+int32(16))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L24
	}
L24:
	;
	F_AppendStringToManifest(m, v158, v201)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L25
	}
L25:
	;
	F_pfree(m, v201)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L16
L27:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	if v216 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v231
	F_errmsg_internal(m, int32(_a_F_perform_base_backup_1), v155)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L38
	}
L29:
	;
	v231 = int32(_a_F_perform_base_backup_2)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v223 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v226 = int32(_a_F_perform_base_backup_3)
	goto L34
L33:
	;
	v226 = int32(_a_F_perform_base_backup_4)
	goto L34
L34:
	;
	if v223 == int32(2) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v229 = int32(_a_F_perform_base_backup_2)
	goto L37
L36:
	;
	v229 = v226
	goto L37
L37:
	;
	v231 = v229
	goto L28
L38:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(72), int32(_a_F_perform_base_backup_6))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v80
	v272 = v57 + int32(2056)
	F_initStringInfo(m, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v80
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[7]))
	if v291 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+5)))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v80
	F_do_pg_backup_start(m, v333, v332, v57+int32(2104), v257, v272)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L46
	}
L43:
	;
	goto L42
L44:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[8])))
	if v295&int32(1) == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v300 = int32(_a_F_perform_base_backup_7)
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	v303 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v302 + v303
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = v306 + v303
	v310 = int32(0)
	v312 = int32(_a_F_perform_base_backup_8)
	v313 = base.AtomicRmwOr32(m, v310, v312, v310)
	*(*int64)(unsafe.Add(mBase, uint32(v291+v310)+232)) = int64(1)
	v321 = base.AtomicRmwOr32(m, v310, v312, v310)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = v322 + v303
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v328 - v303
	goto L43
L46:
	;
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v257)+1032))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2136)) = v350
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v257)+1040))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2144)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v80
	F_before_shmem_exit(m, int32(430), int64(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L47
	}
L47:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[10]))
	v374 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[11]))
	goto L48
L48:
	;
	v376 = v57 + int32(1888)
	*(*int32)(unsafe.Add(mBase, uint32(v376)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v376))) = v57 + int32(332)
	goto L51
L49:
	;
	v382 = int32(0)
	v383 = v84
	v384 = v257
	v385 = v80
	v386 = v372
	v387 = v85
	v388 = v374
	v389 = v121
	goto L9
L51:
	;
	goto L49
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v2819 = m.G0
	v2821 = v2819 - int32(32)
	m.G0 = v2821
	*(*int64)(unsafe.Add(mBase, uint32(v2821)+24)) = int64(17179869184)
	*(*int64)(unsafe.Add(mBase, uint32(v2821))) = int64(4)
	v2829 = *(*int32)(unsafe.Add(mBase, uint32(v57+int32(2104))))
	if v2829 != 0 {
		goto L319
	} else {
		goto L320
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[11])) = v57 + int32(1888)
	if v56 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[10])) = v386
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[11])) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_cancel_before_shmem_exit(m, int32(430), int64(0))
	mBase = m.M
	v2734 = m.ExcPending
	if v2734 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L316
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v410 = int32(0)
	v411 = int64(0)
	v414 = m.G0
	v416 = v414 - int32(2336)
	m.G0 = v416
	v418 = int32(_a_F_perform_base_backup_9)
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[12]))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[12])) = v421
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	if v423 == v410 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v2040 = F_palloc0(m, int32(24))
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L257
	}
L59:
	;
	v1023 = *(*int64)(unsafe.Add(mBase, uint32(v384)+1032))
	F_WaitForWalSummarization(m, v1023)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L141
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L137
	}
L61:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	if v426 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v384)+1040))
	v430 = F_readTimeLineHistory(m, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v434 = F_palloc0(m, v426<<(uint(int32(2))%32))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L64
	}
L64:
	;
	if v426 <= int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384)+1080)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v384)+1072)) = int64(0)
	v997 = v410
	v1021 = v411
	goto L59
L66:
	;
	goto L67
L67:
	;
	v458 = v410
	v460 = v410
	v467 = v410
	v482 = v411
	goto L68
L68:
	;
	v485 = v467 << (uint(int32(2)) % 32)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v486)+12))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v485+v487)))
	if v430 != 0 {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384)+1080)) = v763
	*(*int64)(unsafe.Add(mBase, uint32(v384)+1072)) = v764
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v770)+12))
	v788 = int32(0)
	goto L104
L70:
	;
	if v730&int32(1) != 0 {
		goto L96
	} else {
		goto L97
	}
L71:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	v730 = v687
	v731 = v712
	goto L70
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L92
	}
L73:
	;
	v490 = int32(0)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	if v491 <= v490 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	goto L75
L75:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v434+v485)))
	if v610 != 0 {
		v687 = int32(0)
		goto L71
	} else {
		goto L91
	}
L76:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v434+v485)))
	if v603 == int32(0) {
		goto L72
	} else {
		goto L89
	}
L77:
	;
	v577 = v490
	v579 = int32(0)
	goto L76
L78:
	;
	goto L79
L79:
	;
	v495 = int32(0)
	if v495 < v491 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v498 = v491
	goto L82
L81:
	;
	v498 = v495
	goto L82
L82:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v430)+12))
	v501 = int32(0)
	v518 = v501
	v520 = v490
	v522 = v501
	goto L83
L83:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v500+v518<<(uint(int32(2))%32))))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	if v499 == v549 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v577 = v556
	v579 = v554
	goto L76
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v434+v485))) = v548
	v577 = v520
	v579 = v522
	goto L76
L86:
	;
	goto L87
L87:
	;
	v554 = base.B2i32(v460 == v549) | v522
	v556 = base.B2i32(v458 == v549) | v520
	v558 = v518 + int32(1)
	if v558 != v498 {
		v518 = v558
		v520 = v556
		v522 = v554
		goto L83
	} else {
		goto L88
	}
L88:
	;
	goto L84
L89:
	;
	if v579&int32(1) != 0 {
		v730 = v577
		v731 = v460
		goto L70
	} else {
		goto L90
	}
L90:
	;
	v687 = v577
	goto L71
L91:
	;
	goto L72
L92:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L93
	}
L93:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	*(*int32)(unsafe.Add(mBase, uint32(v416))) = v660
	F_errmsg(m, int32(_a_F_perform_base_backup_10), v416)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(347), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	v758 = int32(0)
	goto L98
L97:
	;
	v758 = v458
	goto L98
L98:
	;
	if v758 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v761 = *(*int64)(unsafe.Add(mBase, uint32(v489)+8))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	v763 = v762
	v764 = v761
	goto L101
L100:
	;
	v763 = v458
	v764 = v482
	goto L101
L101:
	;
	v766 = v467 + int32(1)
	if v766 != v426 {
		v458 = v763
		v460 = v731
		v467 = v766
		v482 = v764
		goto L68
	} else {
		goto L102
	}
L102:
	;
	goto L69
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L133
	}
L104:
	;
	v816 = v788 << (uint(int32(2)) % 32)
	v817 = v434 + v816
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	v819 = *(*int64)(unsafe.Add(mBase, uint32(v818)+8))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v816+v771)))
	v822 = *(*int64)(unsafe.Add(mBase, uint32(v821)+8))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	if v763 == v823 {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L129
	}
L106:
	;
	goto L105
L107:
	;
	v857 = *(*int64)(unsafe.Add(mBase, uint32(v821)+16))
	if v731 == v823 {
		goto L118
	} else {
		goto L119
	}
L108:
	;
	if base.Ui64(v819) <= base.Ui64(v822) {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	if v819 != v822 {
		goto L106
	} else {
		goto L116
	}
L111:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L112
	}
L112:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L113
	}
L113:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	v834 = *(*int64)(unsafe.Add(mBase, uint32(v821)+8))
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	v836 = *(*int64)(unsafe.Add(mBase, uint32(v835)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+32)) = uint32(v836)
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+24)) = uint32(v834)
	v839 = int64(32)
	v840 = int64(base.Ui64(v834) >> (uint(v839) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+20)) = uint32(v840)
	*(*int32)(unsafe.Add(mBase, uint32(v416)+16)) = v833
	v844 = int64(base.Ui64(v836) >> (uint(v839) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+28)) = uint32(v844)
	F_errmsg(m, int32(_a_F_perform_base_backup_13), v416+int32(16))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(414), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	goto L107
L117:
	;
	v899 = v788 + int32(1)
	if v426 != v899 {
		v788 = v899
		goto L104
	} else {
		goto L128
	}
L118:
	;
	v859 = *(*int64)(unsafe.Add(mBase, uint32(v384)+1032))
	if base.Ui64(v857) <= base.Ui64(v859) {
		goto L117
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v896 = *(*int64)(unsafe.Add(mBase, uint32(v818)+16))
	if v857 != v896 {
		goto L103
	} else {
		goto L127
	}
L121:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L122
	}
L122:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L123
	}
L123:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	v869 = *(*int64)(unsafe.Add(mBase, uint32(v821)+16))
	v872 = *(*int64)(unsafe.Add(mBase, uint32(v384)+1032))
	*(*uint32)(unsafe.Add(mBase, uint32(v416-int32(-64)))) = uint32(v872)
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+56)) = uint32(v869)
	*(*int32)(unsafe.Add(mBase, uint32(v416)+48)) = v868
	v876 = int64(32)
	v877 = int64(base.Ui64(v872) >> (uint(v876) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+60)) = uint32(v877)
	v880 = int64(base.Ui64(v869) >> (uint(v876) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+52)) = uint32(v880)
	F_errmsg(m, int32(_a_F_perform_base_backup_14), v416+int32(48))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L124
	}
L124:
	;
	F_errhint(m, int32(_a_F_perform_base_backup_15), int32(0))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(436), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	goto L117
L128:
	;
	v997 = v763
	v1021 = v764
	goto L59
L129:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L130
	}
L130:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	v909 = *(*int64)(unsafe.Add(mBase, uint32(v821)+8))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	v911 = *(*int64)(unsafe.Add(mBase, uint32(v910)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+224)) = uint32(v911)
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+216)) = uint32(v909)
	v914 = int64(32)
	v915 = int64(base.Ui64(v909) >> (uint(v914) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+212)) = uint32(v915)
	*(*int32)(unsafe.Add(mBase, uint32(v416)+208)) = v908
	v919 = int64(base.Ui64(v911) >> (uint(v914) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+220)) = uint32(v919)
	F_errmsg(m, int32(_a_F_perform_base_backup_16), v416+int32(208))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(424), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
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
	F_errcode(m, int32(325))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L134
	}
L134:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	v939 = *(*int64)(unsafe.Add(mBase, uint32(v821)+16))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v434+v788<<(uint(int32(2))%32))))
	v944 = *(*int64)(unsafe.Add(mBase, uint32(v943)+16))
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+192)) = uint32(v944)
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+184)) = uint32(v939)
	v947 = int64(32)
	v948 = int64(base.Ui64(v939) >> (uint(v947) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+180)) = uint32(v948)
	*(*int32)(unsafe.Add(mBase, uint32(v416)+176)) = v938
	v952 = int64(base.Ui64(v944) >> (uint(v947) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+188)) = uint32(v952)
	F_errmsg(m, int32(_a_F_perform_base_backup_17), v416+int32(176))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(446), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L138
	}
L138:
	;
	F_errmsg(m, int32(_a_F_perform_base_backup_18), int32(0))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(291), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	v1026 = int32(0)
	v1028 = *(*int64)(unsafe.Add(mBase, uint32(v384)+1032))
	v1029 = F_GetWalSummaries(m, v1026, v1021, v1028)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L142
	}
L142:
	;
	if v430 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	goto L58
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416)+112)) = v1439
	*(*int32)(unsafe.Add(mBase, uint32(v416)+108)) = v1435
	*(*int32)(unsafe.Add(mBase, uint32(v416)+104)) = v1440
	*(*int32)(unsafe.Add(mBase, uint32(v416)+100)) = v1438
	*(*int32)(unsafe.Add(mBase, uint32(v416)+96)) = v1441
	F_errmsg(m, int32(_a_F_perform_base_backup_19), v416+int32(96))
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L255
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[12])) = v419
	m.G0 = v416 + int32(2336)
	goto L143
L146:
	;
	v1033 = F_CreateEmptyBlockRefTable(m)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	if v1036 <= int32(0) {
		v1540 = v1026
		goto L150
	} else {
		goto L151
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+28)) = v1033
	goto L145
L150:
	;
	v1563 = F_CreateEmptyBlockRefTable(m)
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L216
	}
L151:
	;
	v1039 = int32(0)
	v1058 = v1039
	v1059 = v1039
	v1060 = v1026
	goto L152
L152:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v430)+12))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1083+v1058<<(uint(int32(2))%32))))
	v1088 = *(*int64)(unsafe.Add(mBase, uint32(v1087)+16))
	v1089 = *(*int64)(unsafe.Add(mBase, uint32(v1087)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v416)+240)) = int64(0)
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1087)))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v384)+1040))
	if v1092 == v1093 {
		goto L156
	} else {
		goto L157
	}
L153:
	;
	v1540 = v1493
	goto L150
L154:
	;
	v1518 = v1058 + int32(1)
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	if v1518 < v1519 {
		v1058 = v1518
		v1059 = v1516
		v1060 = v1493
		goto L152
	} else {
		goto L215
	}
L155:
	;
	if v997 == v1092 {
		goto L162
	} else {
		goto L163
	}
L156:
	;
	v1095 = *(*int64)(unsafe.Add(mBase, uint32(v384)+1032))
	v1099 = v1095
	goto L155
L157:
	;
	goto L158
L158:
	;
	if v1059&int32(1) != 0 {
		v1099 = v1088
		goto L155
	} else {
		goto L159
	}
L159:
	;
	v1493 = v1060
	v1516 = int32(0)
	goto L154
L160:
	;
	if v1422 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L161:
	;
	if v1258 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L162:
	;
	v1101 = v1021
	goto L164
L163:
	;
	v1101 = v1089
	goto L164
L164:
	;
	v1102 = int32(0)
	if v1029 == v1102 {
		v1258 = v1102
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+4))
	if int32(0) < v1107 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v1133 = v1102
	v1135 = v1102
	goto L169
L167:
	;
	v1197 = v1102
	goto L168
L168:
	;
	v1258 = v1197
	goto L161
L169:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+12))
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1152+v1135<<(uint(int32(2))%32))))
	if v1092 != 0 {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	v1197 = v1169
	goto L168
L171:
	;
	v1171 = v1135 + int32(1)
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+4))
	if v1171 < v1172 {
		v1133 = v1169
		v1135 = v1171
		goto L169
	} else {
		goto L185
	}
L172:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+16))
	if v1092 != v1157 {
		v1169 = v1133
		goto L171
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	if v1101 != int64(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	goto L174
L176:
	;
	v1161 = *(*int64)(unsafe.Add(mBase, uint32(v1156)+8))
	if base.Ui64(v1161) < base.Ui64(v1101) {
		v1169 = v1133
		goto L171
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	if v1099 != int64(0) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L178
L180:
	;
	v1165 = *(*int64)(unsafe.Add(mBase, uint32(v1156)))
	if base.Ui64(v1099) < base.Ui64(v1165) {
		v1169 = v1133
		goto L171
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v1167 = F_lappend(m, v1133, v1156)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L184
	}
L183:
	;
	goto L182
L184:
	;
	v1169 = v1167
	goto L171
L185:
	;
	goto L170
L186:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v416+int32(240)))) = v1373
	v1422 = int32(0)
	goto L160
L187:
	;
	v1373 = int64(0)
	goto L186
L188:
	;
	goto L189
L189:
	;
	v1264 = F_list_copy(m, v1258)
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L190
	}
L190:
	;
	F_list_sort(m, v1264, int32(494))
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L191
	}
L191:
	;
	if v1264 == int32(0) {
		v1373 = v1101
		goto L186
	} else {
		goto L192
	}
L192:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+4))
	if v1271 <= int32(0) {
		v1373 = v1101
		goto L186
	} else {
		goto L193
	}
L193:
	;
	v1274 = int32(0)
	if v1274 < v1271 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1278 = v1271
	goto L196
L195:
	;
	v1278 = v1274
	goto L196
L196:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+12))
	v1298 = v1274
	v1317 = v1101
	goto L197
L197:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1279+v1298<<(uint(int32(2))%32))))
	v1326 = *(*int64)(unsafe.Add(mBase, uint32(v1325)))
	if base.Ui64(v1317) < base.Ui64(v1326) {
		v1373 = v1317
		goto L186
	} else {
		goto L199
	}
L198:
	;
	v1373 = v1332
	goto L186
L199:
	;
	v1328 = *(*int64)(unsafe.Add(mBase, uint32(v1325)+8))
	if base.Ui64(v1328) <= base.Ui64(v1317) {
		v1332 = v1317
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1334 = v1298 + int32(1)
	if v1334 != v1278 {
		v1298 = v1334
		v1317 = v1332
		goto L197
	} else {
		goto L203
	}
L201:
	;
	if base.Ui64(v1328) < base.Ui64(v1099) {
		v1332 = v1328
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v1422 = int32(1)
	goto L160
L203:
	;
	goto L198
L204:
	;
	v1425 = *(*int64)(unsafe.Add(mBase, uint32(v416)+240))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v1469 = F_list_concat(m, v1060, v1258)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L213
	}
L207:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L208
	}
L208:
	;
	v1433 = int64(32)
	v1435 = base.I32_wrap_i64(int64(base.Ui64(v1099) >> (uint(v1433) % 64)))
	v1438 = base.I32_wrap_i64(int64(base.Ui64(v1101) >> (uint(v1433) % 64)))
	v1439 = base.I32_wrap_i64(v1099)
	v1440 = base.I32_wrap_i64(v1101)
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1087)))
	if v1425 == int64(0) {
		goto L144
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416)+160)) = v1439
	*(*int32)(unsafe.Add(mBase, uint32(v416)+156)) = v1435
	*(*int32)(unsafe.Add(mBase, uint32(v416)+152)) = v1440
	*(*int32)(unsafe.Add(mBase, uint32(v416)+148)) = v1438
	*(*int32)(unsafe.Add(mBase, uint32(v416)+144)) = v1441
	F_errmsg(m, int32(_a_F_perform_base_backup_20), v416+int32(144))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L210
	}
L210:
	;
	v1454 = *(*int64)(unsafe.Add(mBase, uint32(v416)+240))
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+132)) = uint32(v1454)
	v1457 = int64(base.Ui64(v1454) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v416)+128)) = uint32(v1457)
	v1462 = F_errdetail(m, int32(_a_F_perform_base_backup_21), v416+int32(128))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(536), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
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
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1087)))
	if v1471 == v997 {
		v1540 = v1469
		goto L150
	} else {
		goto L214
	}
L214:
	;
	v1493 = v1469
	v1516 = int32(1)
	goto L154
L215:
	;
	goto L153
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+28)) = v1563
	if v1540 == int32(0) {
		goto L145
	} else {
		goto L217
	}
L217:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1540)+4))
	if v1568 <= int32(0) {
		goto L145
	} else {
		goto L218
	}
L218:
	;
	v1588 = int32(0)
	goto L219
L219:
	;
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1540)+12))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1614+v1588<<(uint(int32(2))%32))))
	v1619 = F_OpenWalSummaryFile(m, v1618)
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L221
	}
L220:
	;
	goto L145
L221:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v416)+2328)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v416)+2320)) = v1619
	v1626 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L222
	}
L222:
	;
	if v1626 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v416)+2320))
	v1630 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[13]))
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v1630+v1628*int32(48))+32))
	goto L226
L224:
	;
	goto L225
L225:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v416)+2320))
	v1650 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[13]))
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v1650+v1648*int32(48))+32))
	goto L229
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416)+80)) = v1634
	F_errmsg_internal(m, int32(_a_F_perform_base_backup_22), v416+int32(80))
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(585), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L228
	}
L228:
	;
	goto L225
L229:
	;
	v1655 = F_CreateBlockRefTableReader(m, v416+int32(2320), v1654)
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L230
	}
L230:
	;
	v1663 = F_BlockRefTableReaderNextRelation(m, v1655, v416+int32(2308), v416+int32(2304), v416+int32(2300))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L231
	}
L231:
	;
	if v1663 != 0 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	goto L235
L233:
	;
	goto L234
L234:
	;
	F_DestroyBlockRefTableReader(m, v1655)
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L252
	}
L235:
	;
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v416)+2304))
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v416)+2300))
	F_BlockRefTableSetLimitBlock(m, v1707, v416+int32(2308), v1710, v1711)
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L237
	}
L236:
	;
	goto L234
L237:
	;
	v1717 = F_BlockRefTableReaderGetBlocks(m, v1655, v416+int32(240), int32(512))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L238
	}
L238:
	;
	if v1717 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1736 = v1717
	goto L242
L240:
	;
	goto L241
L241:
	;
	v1870 = F_BlockRefTableReaderNextRelation(m, v1655, v416+int32(2308), v416+int32(2304), v416+int32(2300))
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L250
	}
L242:
	;
	v1777 = int32(0)
	goto L244
L243:
	;
	goto L241
L244:
	;
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v416)+2304))
	v1809 = v416 + int32(240)
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1809+v1777<<(uint(int32(2))%32))))
	F_BlockRefTableMarkBlockModified(m, v1804, v416+int32(2308), v1807, v1813)
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L246
	}
L245:
	;
	v1820 = F_BlockRefTableReaderGetBlocks(m, v1655, v1809, int32(512))
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L248
	}
L246:
	;
	v1817 = v1777 + int32(1)
	if v1817 != v1736 {
		v1777 = v1817
		goto L244
	} else {
		goto L247
	}
L247:
	;
	goto L245
L248:
	;
	if v1820 != 0 {
		v1736 = v1820
		goto L242
	} else {
		goto L249
	}
L249:
	;
	goto L243
L250:
	;
	if v1870 != 0 {
		goto L235
	} else {
		goto L251
	}
L251:
	;
	goto L236
L252:
	;
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v416)+2320))
	F_FileClose(m, v1916)
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L253
	}
L253:
	;
	v1920 = v1588 + int32(1)
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v1540)+4))
	if v1920 < v1921 {
		v1588 = v1920
		goto L219
	} else {
		goto L254
	}
L254:
	;
	goto L220
L255:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(527), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2040)+16)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v57)+2104))
	v2057 = F_lappend(m, v2056, v2040)
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2104)) = v2057
	v2060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+4)))
	if v2060 == int32(1) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v2079 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[7]))
	if v2079 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	v2308 = v89
	v2309 = v90
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = int32(_a_F_perform_base_backup_23)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = v57 + int32(2104)
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(v2320)))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	m.T0[v2321].(func(*base.Module, int32))(m, v55)
	mBase = m.M
	v2335 = m.ExcPending
	if v2335 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L278
	}
L262:
	;
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v57)+2104))
	if v2120 == int32(0) {
		v2264 = v89
		v2265 = v90
		goto L266
	} else {
		goto L267
	}
L263:
	;
	goto L262
L264:
	;
	v2083 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[8])))
	if v2083&int32(1) == int32(0) {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v2088 = int32(_a_F_perform_base_backup_7)
	v2090 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	v2091 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v2090 + v2091
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v2079)))
	*(*int32)(unsafe.Add(mBase, uint32(v2079))) = v2094 + v2091
	v2098 = int32(0)
	v2100 = int32(_a_F_perform_base_backup_8)
	v2101 = base.AtomicRmwOr32(m, v2098, v2100, v2098)
	*(*int64)(unsafe.Add(mBase, uint32(v2079+v2098)+232)) = int64(2)
	v2109 = base.AtomicRmwOr32(m, v2098, v2100, v2098)
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v2079)))
	*(*int32)(unsafe.Add(mBase, uint32(v2079))) = v2110 + v2091
	v2116 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v2116 - v2091
	goto L263
L266:
	;
	v2271 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v389))) = uint8(v2271)
	v2308 = v2264
	v2309 = v2265
	goto L261
L267:
	;
	v2123 = int32(0)
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v2120)+4))
	if v2124 <= v2123 {
		v2264 = v89
		v2265 = v90
		goto L266
	} else {
		goto L268
	}
L268:
	;
	v2131 = v2123
	v2162 = v89
	v2163 = v90
	goto L269
L269:
	;
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v2120)+12))
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v2169+v2131<<(uint(int32(2))%32))))
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+4))
	if v2174 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	v2264 = v2218
	v2265 = v2219
	goto L266
L271:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2173)+16)) = v2220
	v2222 = *(*int64)(unsafe.Add(mBase, uint32(v385)))
	*(*int64)(unsafe.Add(mBase, uint32(v385))) = v2222 + v2220
	v2226 = v2131 + int32(1)
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v2120)+4))
	if v2226 < v2227 {
		v2131 = v2226
		v2162 = v2218
		v2163 = v2219
		goto L269
	} else {
		goto L277
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2162
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2163
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v2190 = int32(1)
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v57)+2104))
	v2194 = int32(0)
	v2197 = F_sendDir(m, v55, int32(_a_F_perform_base_backup_24), v2190, v2190, v2192, v2190, v2194, v2194, v2194)
	mBase = m.M
	v2198 = m.ExcPending
	if v2198 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v2173)))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2162
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2163
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v2213 = int32(0)
	v2215 = F_sendTablespace(m, v55, v2174, v2199, int32(1), v2213, v2213)
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L276
	}
L275:
	;
	v2218 = v2162
	v2219 = v2197
	v2220 = v2197
	goto L271
L276:
	;
	v2218 = v2215
	v2219 = v2163
	v2220 = v2215
	goto L271
L277:
	;
	goto L270
L278:
	;
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(v57)+2104))
	if v2336 == int32(0) {
		goto L52
	} else {
		goto L279
	}
L279:
	;
	v2339 = int32(0)
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v2336)+4))
	if v2340 <= v2339 {
		goto L52
	} else {
		goto L280
	}
L280:
	;
	v2360 = v2339
	goto L281
L281:
	;
	v2385 = *(*int32)(unsafe.Add(mBase, uint32(v2336)+12))
	v2389 = *(*int32)(unsafe.Add(mBase, uint32(v2385+v2360<<(uint(int32(2))%32))))
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(v2389)+4))
	if v2390 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	goto L52
L283:
	;
	v2667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+7)))
	if v2667 == int32(1) {
		goto L309
	} else {
		goto L310
	}
L284:
	;
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	m.T0[v2394].(func(*base.Module, int32, int32))(m, v55, int32(_a_F_perform_base_backup_25))
	mBase = m.M
	v2409 = m.ExcPending
	if v2409 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v2610 = *(*int32)(unsafe.Add(mBase, uint32(v2389)))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+320)) = v2610
	v2627 = F_psprintf(m, int32(_a_F_perform_base_backup_26), v57+int32(320))
	mBase = m.M
	v2628 = m.ExcPending
	if v2628 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L305
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v2423 = F_build_backup_content(m, v384, int32(0))
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v2439 = v57 + int32(2072)
	F_sendFileWithContent(m, v55, int32(_a_F_perform_base_backup_27), v2423, v2439)
	mBase = m.M
	v2441 = m.ExcPending
	if v2441 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L289
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_pfree(m, v2423)
	mBase = m.M
	v2455 = m.ExcPending
	if v2455 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L290
	}
L290:
	;
	v2456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+16)))
	if v2456 == int32(1) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v2472 = *(*int32)(unsafe.Add(mBase, uint32(v57)+2056))
	F_sendFileWithContent(m, v55, int32(_a_F_perform_base_backup_28), v2472, v2439)
	mBase = m.M
	v2474 = m.ExcPending
	if v2474 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v2488 = int32(1)
	v2489 = int32(0)
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(v57)+2104))
	v2496 = F_sendDir(m, v55, int32(_a_F_perform_base_backup_24), v2488, v2489, v2490, v2456^v2488, v57+int32(2072), v2489, v56)
	mBase = m.M
	v2497 = m.ExcPending
	if v2497 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L295
	}
L294:
	;
	goto L293
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v2515 = F___fstatat(m, int32(-100), int32(_a_F_perform_base_backup_29), v57+int32(1792), int32(256))
	mBase = m.M
	goto L296
L296:
	;
	if v2515 != 0 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2531 = m.ExcPending
	if v2531 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v2594 = int32(_a_F_perform_base_backup_29)
	v2598 = int32(0)
	v2608 = F_sendFile(m, v55, v2594, v2594, v57+int32(1792), v2598, v2598, v2598, v2598, v2598, v57+int32(2072), v2598, v2598, v2598)
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L304
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errcode_for_file_access(m)
	mBase = m.M
	v2545 = m.ExcPending
	if v2545 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+304)) = int32(_a_F_perform_base_backup_29)
	F_errmsg(m, int32(_a_F_perform_base_backup_30), v57+int32(304))
	mBase = m.M
	v2564 = m.ExcPending
	if v2564 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(364), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v2581 = m.ExcPending
	if v2581 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L303
	}
L303:
	;
	goto L1
L304:
	;
	goto L283
L305:
	;
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v2629)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	m.T0[v2630].(func(*base.Module, int32, int32))(m, v55, v2627)
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L306
	}
L306:
	;
	v2645 = *(*int32)(unsafe.Add(mBase, uint32(v2389)))
	v2646 = *(*int32)(unsafe.Add(mBase, uint32(v2389)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v2662 = F_sendTablespace(m, v55, v2646, v2645, int32(0), v57+int32(2072), v56)
	mBase = m.M
	v2663 = m.ExcPending
	if v2663 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L307
	}
L307:
	;
	goto L283
L308:
	;
	v2712 = v2360 + int32(1)
	v2713 = *(*int32)(unsafe.Add(mBase, uint32(v2336)+4))
	if v2712 < v2713 {
		v2360 = v2712
		goto L281
	} else {
		goto L315
	}
L309:
	;
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(v2389)+4))
	if v2670 == int32(0) {
		goto L308
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v2673 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v2675 = int32(1024)
	base.MemoryFill(m, v2673, int32(0), v2675)
	v2677 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v2678 = *(*int32)(unsafe.Add(mBase, uint32(v2677)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	m.T0[v2678].(func(*base.Module, int32, int32))(m, v55, v2675)
	mBase = m.M
	v2693 = m.ExcPending
	if v2693 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L313
	}
L312:
	;
	goto L311
L313:
	;
	v2694 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v2695 = *(*int32)(unsafe.Add(mBase, uint32(v2694)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	m.T0[v2695].(func(*base.Module, int32))(m, v55)
	mBase = m.M
	v2709 = m.ExcPending
	if v2709 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L314
	}
L314:
	;
	goto L308
L315:
	;
	goto L282
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_do_pg_abort_backup(m, int32(0), int64(0))
	mBase = m.M
	v2750 = m.ExcPending
	if v2750 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_pg_re_throw(m)
	mBase = m.M
	v2764 = m.ExcPending
	if v2764 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L318
	}
L318:
	;
	goto L1
L319:
	;
	v2830 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2829)+4)))
	v2832 = v2830
	goto L321
L320:
	;
	v2832 = int64(0)
	goto L321
L321:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2821)+8)) = v2832
	goto L324
L322:
	;
	m.G0 = v2821 + int32(32)
	v3021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_do_pg_backup_stop(m, v384, (v3021^int32(-1))&int32(1))
	mBase = m.M
	v3039 = m.ExcPending
	if v3039 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L339
	}
L323:
	;
	goto L322
L324:
	;
	v2846 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[7]))
	if v2846 == int32(0) {
		goto L323
	} else {
		goto L325
	}
L325:
	;
	v2850 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[8])))
	if v2850&int32(1) == int32(0) {
		goto L323
	} else {
		goto L326
	}
L326:
	;
	v2855 = int32(_a_F_perform_base_backup_7)
	v2857 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	v2858 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v2857 + v2858
	v2861 = *(*int32)(unsafe.Add(mBase, uint32(v2846)))
	*(*int32)(unsafe.Add(mBase, uint32(v2846))) = v2861 + v2858
	v2865 = int32(0)
	v2868 = base.AtomicRmwOr32(m, v2865, int32(_a_F_perform_base_backup_8), v2865)
	goto L328
L327:
	;
	v2995 = int32(0)
	v2998 = base.AtomicRmwOr32(m, v2995, int32(_a_F_perform_base_backup_8), v2995)
	v2999 = *(*int32)(unsafe.Add(mBase, uint32(v2846)))
	v3000 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2846))) = v2999 + v3000
	v3003 = int32(_a_F_perform_base_backup_7)
	v3005 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v3005 - v3000
	goto L323
L328:
	;
	goto L330
L330:
	;
	goto L331
L331:
	;
	v2960 = int32(0)
	v2963 = int32(0)
	goto L336
L336:
	;
	v2972 = *(*int32)(unsafe.Add(mBase, uint32(v2821+int32(24)+v2963<<(uint(int32(2))%32))))
	v2973 = int32(3)
	v2979 = *(*int64)(unsafe.Add(mBase, uint32(v2821+v2963<<(uint(v2973)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2846+int32(232)+v2972<<(uint(v2973)%32)))) = v2979
	v2981 = int32(1)
	v2984 = v2960 + v2981
	if v2984 != int32(2) {
		v2960 = v2984
		v2963 = v2963 + v2981
		goto L336
	} else {
		goto L338
	}
L337:
	;
	goto L327
L338:
	;
	goto L337
L339:
	;
	v3040 = *(*int32)(unsafe.Add(mBase, uint32(v384)+1096))
	v3041 = *(*int64)(unsafe.Add(mBase, uint32(v384)+1088))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3054 = *(*int32)(unsafe.Add(mBase, uint32(v57)+2056))
	F_pfree(m, v3054)
	mBase = m.M
	v3056 = m.ExcPending
	if v3056 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L340
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_pfree(m, v384)
	mBase = m.M
	v3070 = m.ExcPending
	if v3070 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_cancel_before_shmem_exit(m, int32(430), int64(0))
	mBase = m.M
	v3086 = m.ExcPending
	if v3086 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L342
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[10])) = v386
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[11])) = v388
	v3091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+7)))
	if v3091 != 0 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3108 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[7]))
	if v3108 == int32(0) {
		goto L347
	} else {
		goto L348
	}
L344:
	;
	v5324 = v65
	v5325 = v66
	v5326 = v67
	goto L345
L345:
	;
	v5355 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v5356 = *(*int64)(unsafe.Add(mBase, uint32(v387)))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v5325
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v5324
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v5326
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v5370 = m.G0
	v5372 = v5370 - int32(80)
	m.G0 = v5372
	v5375 = v57 + int32(2072)
	v5376 = *(*int32)(unsafe.Add(mBase, uint32(v5375)))
	if v5376 != 0 {
		goto L559
	} else {
		goto L560
	}
L346:
	;
	v3149 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v3150 = *(*int64)(unsafe.Add(mBase, uint32(v387)))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3164 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+288)) = v3149
	v3166 = base.I64_div_u_s(v3150, v3164)
	v3168 = base.I64_div_u_s(int64(4294967296), v3164)
	v3169 = base.I64_div_u_s(v3166, v3168)
	*(*uint32)(unsafe.Add(mBase, uint32(v57)+292)) = uint32(v3169)
	v3172 = v3166 - v3168*v3169
	*(*uint32)(unsafe.Add(mBase, uint32(v57)+296)) = uint32(v3172)
	v3180 = F_pg_snprintf(m, v57+int32(608), int32(64), int32(_a_F_perform_base_backup_33), v57+int32(288))
	mBase = m.M
	v3181 = m.ExcPending
	if v3181 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L350
	}
L347:
	;
	goto L346
L348:
	;
	v3112 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[8])))
	if v3112&int32(1) == int32(0) {
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v3117 = int32(_a_F_perform_base_backup_7)
	v3119 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	v3120 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v3119 + v3120
	v3123 = *(*int32)(unsafe.Add(mBase, uint32(v3108)))
	*(*int32)(unsafe.Add(mBase, uint32(v3108))) = v3123 + v3120
	v3127 = int32(0)
	v3129 = int32(_a_F_perform_base_backup_8)
	v3130 = base.AtomicRmwOr32(m, v3127, v3129, v3127)
	*(*int64)(unsafe.Add(mBase, uint32(v3108+v3127)+232)) = int64(5)
	v3138 = base.AtomicRmwOr32(m, v3127, v3129, v3127)
	v3139 = *(*int32)(unsafe.Add(mBase, uint32(v3108)))
	*(*int32)(unsafe.Add(mBase, uint32(v3108))) = v3139 + v3120
	v3145 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v3145 - v3120
	goto L347
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3195 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+272)) = v3040
	v3199 = base.I64_div_u_s(v3041-int64(1), v3195)
	v3201 = base.I64_div_u_s(int64(4294967296), v3195)
	v3202 = base.I64_div_u_s(v3199, v3201)
	*(*uint32)(unsafe.Add(mBase, uint32(v57)+276)) = uint32(v3202)
	v3205 = v3199 - v3201*v3202
	*(*uint32)(unsafe.Add(mBase, uint32(v57)+280)) = uint32(v3205)
	v3213 = F_pg_snprintf(m, v57+int32(544), int32(64), int32(_a_F_perform_base_backup_33), v57+int32(272))
	mBase = m.M
	v3214 = m.ExcPending
	if v3214 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3228 = F_AllocateDir(m, int32(_a_F_perform_base_backup_34))
	mBase = m.M
	v3229 = m.ExcPending
	if v3229 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L352
	}
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3243 = F_ReadDir(m, v3228, int32(_a_F_perform_base_backup_34))
	mBase = m.M
	v3244 = m.ExcPending
	if v3244 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L354
	}
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_FreeDir(m, v3228)
	mBase = m.M
	v3768 = m.ExcPending
	if v3768 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L433
	}
L354:
	;
	if v3243 == int32(0) {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v3247 = int32(0)
	v3724 = v65
	v3725 = v66
	v3726 = v67
	v3728 = v3247
	v3729 = v3247
	goto L353
L356:
	;
	goto L357
L357:
	;
	v3251 = int32(8)
	v3252 = v57 + int32(544) | v3251
	v3256 = v57 + int32(608) | v3251
	v3257 = int32(0)
	v3263 = v3243
	v3270 = v65
	v3271 = v66
	v3272 = v67
	v3274 = v3257
	v3275 = v3257
	goto L358
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3271
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3270
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3272
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3314 = v3263 + int32(19)
	v3315 = F_strlen(m, v3314)
	mBase = m.M
	switch v3315 - int32(16) {
	case 0:
		goto L361
	default:
		v3694 = v3271
		v3695 = v3272
		v3696 = v3274
		v3697 = v3275
		goto L360
	case 8:
		goto L362
	}
L359:
	;
	v3724 = v3711
	v3725 = v3694
	v3726 = v3695
	v3728 = v3696
	v3729 = v3697
	goto L353
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3694
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3270
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3695
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3711 = F_ReadDir(m, v3228, int32(_a_F_perform_base_backup_34))
	mBase = m.M
	v3712 = m.ExcPending
	if v3712 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L431
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3271
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3270
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3272
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3538 = int32(_a_F_perform_base_backup_35)
	v3542 = m.G0
	v3544 = v3542 - int32(32)
	v3545 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3544)+24)) = v3545
	*(*int64)(unsafe.Add(mBase, uint32(v3544)+16)) = v3545
	*(*int64)(unsafe.Add(mBase, uint32(v3544)+8)) = v3545
	*(*int64)(unsafe.Add(mBase, uint32(v3544))) = v3545
	v3553 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[15])))
	if v3553 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3271
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3270
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3272
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3330 = int32(_a_F_perform_base_backup_35)
	v3334 = m.G0
	v3336 = v3334 - int32(32)
	v3337 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3336)+24)) = v3337
	*(*int64)(unsafe.Add(mBase, uint32(v3336)+16)) = v3337
	*(*int64)(unsafe.Add(mBase, uint32(v3336)+8)) = v3337
	*(*int64)(unsafe.Add(mBase, uint32(v3336))) = v3337
	v3345 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[15])))
	if v3345 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L363:
	;
	if v3413 != int32(24) {
		v3694 = v3271
		v3695 = v3272
		v3696 = v3274
		v3697 = v3275
		goto L360
	} else {
		goto L382
	}
L364:
	;
	v3413 = int32(0)
	goto L363
L365:
	;
	goto L366
L366:
	;
	v3349 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[16])))
	if v3349 == int32(0) {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v3353 = v3314
	goto L370
L368:
	;
	goto L369
L369:
	;
	v3363 = v3330
	v3364 = v3345
	goto L373
L370:
	;
	v3359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3353))))
	if v3359 == v3345 {
		v3353 = v3353 + int32(1)
		goto L370
	} else {
		goto L372
	}
L371:
	;
	v3413 = v3353 - v3314
	goto L363
L372:
	;
	goto L371
L373:
	;
	v3371 = v3336 + int32(base.Ui32(v3364)>>(uint(int32(3))%32))&int32(28)
	v3372 = *(*int32)(unsafe.Add(mBase, uint32(v3371)))
	v3373 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3371))) = v3372 | v3373<<(uint(v3364)%32)
	v3377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3363)+1)))
	if v3377 != 0 {
		v3363 = v3363 + v3373
		v3364 = v3377
		goto L373
	} else {
		goto L375
	}
L374:
	;
	v3380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3314))))
	if v3380 == int32(0) {
		v3403 = v3314
		goto L376
	} else {
		goto L377
	}
L375:
	;
	goto L374
L376:
	;
	v3413 = v3403 - v3314
	goto L363
L377:
	;
	v3384 = v3314
	v3385 = v3380
	goto L378
L378:
	;
	v3393 = *(*int32)(unsafe.Add(mBase, uint32(v3336+int32(base.Ui32(v3385)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3393)>>(uint(v3385)%32))&int32(1) == int32(0) {
		v3403 = v3384
		goto L376
	} else {
		goto L380
	}
L379:
	;
	v3403 = v3401
	goto L376
L380:
	;
	v3399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3384)+1)))
	v3401 = v3384 + int32(1)
	if v3399 != 0 {
		v3384 = v3401
		v3385 = v3399
		goto L378
	} else {
		goto L381
	}
L381:
	;
	goto L379
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3271
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3270
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3272
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3429 = v3263 + int32(27)
	v3432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3429))))
	v3435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3256))))
	if base.B2i32(v3432 == int32(0))|base.B2i32(v3432 != v3435) != 0 {
		v3453 = v3432
		v3454 = v3435
		goto L384
	} else {
		goto L385
	}
L383:
	;
	if v3453-v3454 < int32(0) {
		v3694 = v3271
		v3695 = v3272
		v3696 = v3274
		v3697 = v3275
		goto L360
	} else {
		goto L390
	}
L384:
	;
	goto L383
L385:
	;
	v3438 = v3429
	v3439 = v3256
	goto L386
L386:
	;
	v3442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3439)+1)))
	v3443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3438)+1)))
	if v3443 == int32(0) {
		v3453 = v3443
		v3454 = v3442
		goto L384
	} else {
		goto L388
	}
L387:
	;
	v3453 = v3443
	v3454 = v3442
	goto L384
L388:
	;
	v3446 = int32(1)
	if v3443 == v3442 {
		v3438 = v3438 + v3446
		v3439 = v3439 + v3446
		goto L386
	} else {
		goto L389
	}
L389:
	;
	goto L387
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3271
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3270
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3272
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3429))))
	v3475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3252))))
	if base.B2i32(v3472 == int32(0))|base.B2i32(v3472 != v3475) != 0 {
		v3493 = v3472
		v3494 = v3475
		goto L392
	} else {
		goto L393
	}
L391:
	;
	if int32(0) < v3493-v3494 {
		v3694 = v3271
		v3695 = v3272
		v3696 = v3274
		v3697 = v3275
		goto L360
	} else {
		goto L398
	}
L392:
	;
	goto L391
L393:
	;
	v3478 = v3429
	v3479 = v3252
	goto L394
L394:
	;
	v3482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3479)+1)))
	v3483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3478)+1)))
	if v3483 == int32(0) {
		v3493 = v3483
		v3494 = v3482
		goto L392
	} else {
		goto L396
	}
L395:
	;
	v3493 = v3483
	v3494 = v3482
	goto L392
L396:
	;
	v3486 = int32(1)
	if v3483 == v3482 {
		v3478 = v3478 + v3486
		v3479 = v3479 + v3486
		goto L394
	} else {
		goto L397
	}
L397:
	;
	goto L395
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3271
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3270
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3272
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3510 = F_pstrdup(m, v3314)
	mBase = m.M
	v3511 = m.ExcPending
	if v3511 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L399
	}
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3271
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3270
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3272
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3524 = F_lappend(m, v3274, v3510)
	mBase = m.M
	v3525 = m.ExcPending
	if v3525 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L400
	}
L400:
	;
	v3694 = v3271
	v3695 = v3524
	v3696 = v3524
	v3697 = v3275
	goto L360
L401:
	;
	if v3621 != int32(8) {
		v3694 = v3271
		v3695 = v3272
		v3696 = v3274
		v3697 = v3275
		goto L360
	} else {
		goto L420
	}
L402:
	;
	v3621 = int32(0)
	goto L401
L403:
	;
	goto L404
L404:
	;
	v3557 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[16])))
	if v3557 == int32(0) {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v3561 = v3314
	goto L408
L406:
	;
	goto L407
L407:
	;
	v3571 = v3538
	v3572 = v3553
	goto L411
L408:
	;
	v3567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3561))))
	if v3567 == v3553 {
		v3561 = v3561 + int32(1)
		goto L408
	} else {
		goto L410
	}
L409:
	;
	v3621 = v3561 - v3314
	goto L401
L410:
	;
	goto L409
L411:
	;
	v3579 = v3544 + int32(base.Ui32(v3572)>>(uint(int32(3))%32))&int32(28)
	v3580 = *(*int32)(unsafe.Add(mBase, uint32(v3579)))
	v3581 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3579))) = v3580 | v3581<<(uint(v3572)%32)
	v3585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3571)+1)))
	if v3585 != 0 {
		v3571 = v3571 + v3581
		v3572 = v3585
		goto L411
	} else {
		goto L413
	}
L412:
	;
	v3588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3314))))
	if v3588 == int32(0) {
		v3611 = v3314
		goto L414
	} else {
		goto L415
	}
L413:
	;
	goto L412
L414:
	;
	v3621 = v3611 - v3314
	goto L401
L415:
	;
	v3592 = v3314
	v3593 = v3588
	goto L416
L416:
	;
	v3601 = *(*int32)(unsafe.Add(mBase, uint32(v3544+int32(base.Ui32(v3593)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3601)>>(uint(v3593)%32))&int32(1) == int32(0) {
		v3611 = v3592
		goto L414
	} else {
		goto L418
	}
L417:
	;
	v3611 = v3609
	goto L414
L418:
	;
	v3607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3592)+1)))
	v3609 = v3592 + int32(1)
	if v3607 != 0 {
		v3592 = v3609
		v3593 = v3607
		goto L416
	} else {
		goto L419
	}
L419:
	;
	goto L417
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3271
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3270
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3272
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3637 = v3263 + int32(27)
	v3638 = int32(_a_F_perform_base_backup_36)
	v3641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3637))))
	v3644 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[17])))
	if base.B2i32(v3641 == int32(0))|base.B2i32(v3641 != v3644) != 0 {
		v3662 = v3641
		v3663 = v3644
		goto L422
	} else {
		goto L423
	}
L421:
	;
	if v3662-v3663 != 0 {
		v3694 = v3271
		v3695 = v3272
		v3696 = v3274
		v3697 = v3275
		goto L360
	} else {
		goto L428
	}
L422:
	;
	goto L421
L423:
	;
	v3647 = v3637
	v3648 = v3638
	goto L424
L424:
	;
	v3651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3648)+1)))
	v3652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3647)+1)))
	if v3652 == int32(0) {
		v3662 = v3652
		v3663 = v3651
		goto L422
	} else {
		goto L426
	}
L425:
	;
	v3662 = v3652
	v3663 = v3651
	goto L422
L426:
	;
	v3655 = int32(1)
	if v3652 == v3651 {
		v3647 = v3647 + v3655
		v3648 = v3648 + v3655
		goto L424
	} else {
		goto L427
	}
L427:
	;
	goto L425
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3271
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3270
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3272
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3677 = F_pstrdup(m, v3314)
	mBase = m.M
	v3678 = m.ExcPending
	if v3678 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3271
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3270
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3272
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3691 = F_lappend(m, v3275, v3677)
	mBase = m.M
	v3692 = m.ExcPending
	if v3692 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L430
	}
L430:
	;
	v3694 = v3691
	v3695 = v3272
	v3696 = v3274
	v3697 = v3691
	goto L360
L431:
	;
	if v3711 != 0 {
		v3263 = v3711
		v3270 = v3711
		v3271 = v3694
		v3272 = v3695
		v3274 = v3696
		v3275 = v3697
		goto L358
	} else {
		goto L432
	}
L432:
	;
	goto L359
L433:
	;
	v3769 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_CheckXLogRemoved(m, v3166, v3769)
	mBase = m.M
	v3783 = m.ExcPending
	if v3783 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L434
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_list_sort(m, v3728, int32(451))
	mBase = m.M
	v3798 = m.ExcPending
	if v3798 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L435
	}
L435:
	;
	if v3728 == int32(0) {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3816 = m.ExcPending
	if v3816 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L439
	}
L437:
	;
	goto L438
L438:
	;
	v3850 = *(*int32)(unsafe.Add(mBase, uint32(v3728)+12))
	v3851 = *(*int32)(unsafe.Add(mBase, uint32(v3850)))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3865 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+264)) = v57 + int32(2156)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+260)) = v57 + int32(2160)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+256)) = v57 + int32(540)
	v3878 = F_sscanf(m, v3851, int32(_a_F_perform_base_backup_33), v57+int32(256))
	mBase = m.M
	v3879 = m.ExcPending
	if v3879 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L442
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errmsg(m, int32(_a_F_perform_base_backup_37), int32(0))
	mBase = m.M
	v3832 = m.ExcPending
	if v3832 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L440
	}
L440:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(485), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v3849 = m.ExcPending
	if v3849 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L441
	}
L441:
	;
	goto L1
L442:
	;
	v3880 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57)+2156)))
	v3881 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57)+2160)))
	v3883 = base.I64_div_u_s(int64(4294967296), v3865)
	if v3166 == v3880+v3881*v3883 {
		goto L447
	} else {
		goto L448
	}
L443:
	;
	if v3729 == int32(0) {
		goto L536
	} else {
		goto L537
	}
L444:
	;
	if v4052 <= int32(0) {
		goto L443
	} else {
		goto L473
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v4180 = v57 + int32(336)
	v4182 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	F_XLogFileName(m, v4180, v3040, v3199, v4182)
	mBase = m.M
	v4184 = m.ExcPending
	if v4184 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L469
	}
L446:
	;
	v3966 = v3887
	v4000 = v3166
	goto L457
L447:
	;
	v3887 = int32(0)
	v3888 = *(*int32)(unsafe.Add(mBase, uint32(v3728)+4))
	if v3887 < v3888 {
		goto L446
	} else {
		goto L450
	}
L448:
	;
	goto L449
L449:
	;
	v3892 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v3906 = v57 + int32(464)
	v3908 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	F_XLogFileName(m, v3906, v3892, v3166, v3908)
	mBase = m.M
	v3910 = m.ExcPending
	if v3910 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L452
	}
L450:
	;
	if v3166 != v3199 {
		goto L445
	} else {
		goto L451
	}
L451:
	;
	goto L443
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3926 = m.ExcPending
	if v3926 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+240)) = v3906
	F_errmsg(m, int32(_a_F_perform_base_backup_38), v57+int32(240))
	mBase = m.M
	v3944 = m.ExcPending
	if v3944 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L454
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(500), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v3961 = m.ExcPending
	if v3961 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L455
	}
L455:
	;
	goto L1
L456:
	;
	if v4044 == v3199 {
		goto L444
	} else {
		goto L468
	}
L457:
	;
	v4004 = *(*int32)(unsafe.Add(mBase, uint32(v3728)+12))
	v4008 = *(*int32)(unsafe.Add(mBase, uint32(v4004+v3966<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v4022 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+232)) = v57 + int32(2164)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+228)) = v57 + int32(2168)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+224)) = v57 + int32(540)
	v4035 = F_sscanf(m, v4008, int32(_a_F_perform_base_backup_33), v57+int32(224))
	mBase = m.M
	v4036 = m.ExcPending
	if v4036 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L459
	}
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v4067 = v57 + int32(400)
	v4068 = *(*int32)(unsafe.Add(mBase, uint32(v57)+540))
	v4070 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	F_XLogFileName(m, v4067, v4068, v4038, v4070)
	mBase = m.M
	v4072 = m.ExcPending
	if v4072 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L464
	}
L459:
	;
	v4038 = v4000 + int64(1)
	v4039 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57)+2164)))
	v4040 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57)+2168)))
	v4042 = base.I64_div_u_s(int64(4294967296), v4022)
	v4044 = v4039 + v4040*v4042
	if base.B2i32(v4038 != v4044)&base.B2i32(v4000 != v4044) == int32(0) {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v4051 = v3966 + int32(1)
	v4052 = *(*int32)(unsafe.Add(mBase, uint32(v3728)+4))
	if v4052 <= v4051 {
		goto L456
	} else {
		goto L463
	}
L461:
	;
	goto L462
L462:
	;
	goto L458
L463:
	;
	v3966 = v4051
	v4000 = v4044
	goto L457
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4088 = m.ExcPending
	if v4088 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L465
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+208)) = v4067
	F_errmsg(m, int32(_a_F_perform_base_backup_38), v57+int32(208))
	mBase = m.M
	v4106 = m.ExcPending
	if v4106 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L466
	}
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(515), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v4123 = m.ExcPending
	if v4123 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L467
	}
L467:
	;
	goto L1
L468:
	;
	goto L445
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4200 = m.ExcPending
	if v4200 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L470
	}
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+192)) = v4180
	F_errmsg(m, int32(_a_F_perform_base_backup_38), v57+int32(192))
	mBase = m.M
	v4218 = m.ExcPending
	if v4218 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L471
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(524), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v4235 = m.ExcPending
	if v4235 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L472
	}
L472:
	;
	goto L1
L473:
	;
	v4257 = int32(0)
	goto L474
L474:
	;
	v4281 = *(*int32)(unsafe.Add(mBase, uint32(v3728)+12))
	v4285 = *(*int32)(unsafe.Add(mBase, uint32(v4281+v4257<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+176)) = v4285
	v4300 = v57 + int32(768)
	v4305 = F_pg_snprintf(m, v4300, int32(1024), int32(_a_F_perform_base_backup_39), v57+int32(176))
	mBase = m.M
	v4306 = m.ExcPending
	if v4306 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L476
	}
L475:
	;
	goto L443
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v4320 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+168)) = v57 + int32(2172)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+164)) = v57 + int32(2176)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+160)) = v57 + int32(540)
	v4333 = F_sscanf(m, v4285, int32(_a_F_perform_base_backup_33), v57+int32(160))
	mBase = m.M
	v4334 = m.ExcPending
	if v4334 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L477
	}
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v4348 = base.I64_div_u_s(int64(4294967296), v4320)
	v4349 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57)+2176)))
	v4350 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57)+2172)))
	v4352 = F_OpenTransientFile(m, v4300, int32(0))
	mBase = m.M
	v4353 = m.ExcPending
	if v4353 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L478
	}
L478:
	;
	v4355 = v4348*v4349 + v4350
	if v4352 < int32(0) {
		goto L479
	} else {
		goto L480
	}
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v4371 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v4384 = *(*int32)(unsafe.Add(mBase, uint32(v57)+540))
	F_CheckXLogRemoved(m, v4355, v4384)
	mBase = m.M
	v4386 = m.ExcPending
	if v4386 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L482
	}
L480:
	;
	goto L481
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	if v4352 < int32(0) {
		goto L488
	} else {
		goto L489
	}
L482:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[18])) = v4371
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4404 = m.ExcPending
	if v4404 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L483
	}
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errcode_for_file_access(m)
	mBase = m.M
	v4418 = m.ExcPending
	if v4418 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L484
	}
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v4300
	F_errmsg(m, int32(_a_F_perform_base_backup_40), v57)
	mBase = m.M
	v4434 = m.ExcPending
	if v4434 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L485
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(553), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v4451 = m.ExcPending
	if v4451 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L486
	}
L486:
	;
	goto L1
L487:
	;
	if v4473 != 0 {
		goto L491
	} else {
		goto L492
	}
L488:
	;
	v4469 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v4473 = v4469
	goto L487
L489:
	;
	goto L490
L490:
	;
	v4472 = F___fstatat(m, v4352, int32(_a_F_perform_base_backup_41), v57+int32(672), int32(_a_F_perform_base_backup_42))
	mBase = m.M
	v4473 = v4472
	goto L487
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4489 = m.ExcPending
	if v4489 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L494
	}
L492:
	;
	goto L493
L493:
	;
	v4541 = *(*int64)(unsafe.Add(mBase, uint32(v57)+696))
	v4543 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	if v4541 != v4543 {
		goto L498
	} else {
		goto L499
	}
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errcode_for_file_access(m)
	mBase = m.M
	v4503 = m.ExcPending
	if v4503 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L495
	}
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+144)) = v57 + int32(768)
	F_errmsg(m, int32(_a_F_perform_base_backup_30), v57+int32(144))
	mBase = m.M
	v4523 = m.ExcPending
	if v4523 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L496
	}
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(560), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v4540 = m.ExcPending
	if v4540 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L497
	}
L497:
	;
	goto L1
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v4557 = *(*int32)(unsafe.Add(mBase, uint32(v57)+540))
	F_CheckXLogRemoved(m, v4355, v4557)
	mBase = m.M
	v4559 = m.ExcPending
	if v4559 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L501
	}
L499:
	;
	goto L500
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v4639 = int32(0)
	F__tarWriteHeader(m, v55, v57+int32(768), v4639, v57+int32(672), v4639)
	mBase = m.M
	v4644 = m.ExcPending
	if v4644 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L506
	}
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4575 = m.ExcPending
	if v4575 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L502
	}
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errcode_for_file_access(m)
	mBase = m.M
	v4589 = m.ExcPending
	if v4589 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L503
	}
L503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+128)) = v4285
	F_errmsg(m, int32(_a_F_perform_base_backup_43), v57+int32(128))
	mBase = m.M
	v4607 = m.ExcPending
	if v4607 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L504
	}
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(566), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v4624 = m.ExcPending
	if v4624 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L505
	}
L505:
	;
	goto L1
L506:
	;
	v4646 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	v4652 = v4646
	v4686 = int64(0)
	goto L508
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v4922 = F_CloseTransientFile(m, v4352)
	mBase = m.M
	v4923 = m.ExcPending
	if v4923 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L532
	}
L508:
	;
	v4690 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v4691 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v55)+8)))
	v4693 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v4693))) = int32(167772163)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v4709 = base.I64_extend_i32_s(v4652) - v4686
	if v4709 < v4691 {
		goto L510
	} else {
		goto L511
	}
L509:
	;
	v4825 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	if v4686 == v4825 {
		goto L507
	} else {
		goto L526
	}
L510:
	;
	v4711 = v4709
	goto L512
L511:
	;
	v4711 = v4691
	goto L512
L512:
	;
	v4713 = F_pread(m, v4352, v4690, base.I32_wrap_i64(v4711), v4686)
	mBase = m.M
	v4715 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[19]))
	v4716 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4715))) = v4716
	if v4713 < v4716 {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4735 = m.ExcPending
	if v4735 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L516
	}
L514:
	;
	goto L515
L515:
	;
	if v4713 != 0 {
		goto L520
	} else {
		goto L521
	}
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errcode_for_file_access(m)
	mBase = m.M
	v4749 = m.ExcPending
	if v4749 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L517
	}
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v57 + int32(768)
	F_errmsg(m, int32(_a_F_perform_base_backup_44), v57+int32(16))
	mBase = m.M
	v4769 = m.ExcPending
	if v4769 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L518
	}
L518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(2127), int32(_a_F_perform_base_backup_45))
	mBase = m.M
	v4786 = m.ExcPending
	if v4786 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L519
	}
L519:
	;
	goto L1
L520:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v4799 = *(*int32)(unsafe.Add(mBase, uint32(v57)+540))
	F_CheckXLogRemoved(m, v4355, v4799)
	mBase = m.M
	v4801 = m.ExcPending
	if v4801 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L523
	}
L521:
	;
	goto L522
L522:
	;
	goto L509
L523:
	;
	v4802 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v4803 = *(*int32)(unsafe.Add(mBase, uint32(v4802)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	m.T0[v4803].(func(*base.Module, int32, int32))(m, v55, v4713)
	mBase = m.M
	v4817 = m.ExcPending
	if v4817 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L524
	}
L524:
	;
	v4819 = v4686 + base.I64_extend_i32_u(v4713)
	v4821 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	if v4819 != base.I64_extend_i32_s(v4821) {
		v4652 = v4821
		v4686 = v4819
		goto L508
	} else {
		goto L525
	}
L525:
	;
	goto L507
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v4839 = *(*int32)(unsafe.Add(mBase, uint32(v57)+540))
	F_CheckXLogRemoved(m, v4355, v4839)
	mBase = m.M
	v4841 = m.ExcPending
	if v4841 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L527
	}
L527:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4857 = m.ExcPending
	if v4857 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L528
	}
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errcode_for_file_access(m)
	mBase = m.M
	v4871 = m.ExcPending
	if v4871 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L529
	}
L529:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+112)) = v4285
	F_errmsg(m, int32(_a_F_perform_base_backup_43), v57+int32(112))
	mBase = m.M
	v4889 = m.ExcPending
	if v4889 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L530
	}
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(591), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v4906 = m.ExcPending
	if v4906 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L531
	}
L531:
	;
	goto L1
L532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+100)) = int32(_a_F_perform_base_backup_46)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+96)) = v4285
	v4940 = v57 + int32(768)
	v4945 = F_pg_snprintf(m, v4940, int32(1024), int32(_a_F_perform_base_backup_47), v57+int32(96))
	mBase = m.M
	v4946 = m.ExcPending
	if v4946 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L533
	}
L533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_sendFileWithContent(m, v55, v4940, int32(_a_F_perform_base_backup_41), v57+int32(2072))
	mBase = m.M
	v4963 = m.ExcPending
	if v4963 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L534
	}
L534:
	;
	v4965 = v4257 + int32(1)
	v4966 = *(*int32)(unsafe.Add(mBase, uint32(v3728)+4))
	if v4965 < v4966 {
		v4257 = v4965
		goto L474
	} else {
		goto L535
	}
L535:
	;
	goto L475
L536:
	;
	v5276 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v5278 = int32(1024)
	base.MemoryFill(m, v5276, int32(0), v5278)
	v5280 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v5281 = *(*int32)(unsafe.Add(mBase, uint32(v5280)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	m.T0[v5281].(func(*base.Module, int32, int32))(m, v55, v5278)
	mBase = m.M
	v5296 = m.ExcPending
	if v5296 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L554
	}
L537:
	;
	v5012 = int32(0)
	v5013 = *(*int32)(unsafe.Add(mBase, uint32(v3729)+4))
	if v5013 <= v5012 {
		goto L536
	} else {
		goto L538
	}
L538:
	;
	v5020 = v5012
	goto L539
L539:
	;
	v5058 = *(*int32)(unsafe.Add(mBase, uint32(v3729)+12))
	v5062 = *(*int32)(unsafe.Add(mBase, uint32(v5058+v5020<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+80)) = v5062
	v5077 = v57 + int32(768)
	v5082 = F_pg_snprintf(m, v5077, int32(1024), int32(_a_F_perform_base_backup_39), v57+int32(80))
	mBase = m.M
	v5083 = m.ExcPending
	if v5083 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L541
	}
L540:
	;
	goto L536
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v5100 = F___fstatat(m, int32(-100), v5077, v57+int32(672), int32(256))
	mBase = m.M
	goto L542
L542:
	;
	if v5100 != 0 {
		goto L543
	} else {
		goto L544
	}
L543:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5116 = m.ExcPending
	if v5116 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L546
	}
L544:
	;
	goto L545
L545:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v5179 = v57 + int32(768)
	v5182 = int32(0)
	v5188 = v57 + int32(2072)
	v5192 = F_sendFile(m, v55, v5179, v5179, v57+int32(672), v5182, v5182, v5182, v5182, v5182, v5188, v5182, v5182, v5182)
	mBase = m.M
	v5193 = m.ExcPending
	if v5193 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L550
	}
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errcode_for_file_access(m)
	mBase = m.M
	v5130 = m.ExcPending
	if v5130 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L547
	}
L547:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+64)) = v5077
	F_errmsg(m, int32(_a_F_perform_base_backup_30), v57-int32(-64))
	mBase = m.M
	v5148 = m.ExcPending
	if v5148 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L548
	}
L548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(630), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v5165 = m.ExcPending
	if v5165 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L549
	}
L549:
	;
	goto L1
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v57)+52)) = int32(_a_F_perform_base_backup_46)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+48)) = v5062
	v5213 = F_pg_snprintf(m, v5179, int32(1024), int32(_a_F_perform_base_backup_47), v57+int32(48))
	mBase = m.M
	v5214 = m.ExcPending
	if v5214 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L551
	}
L551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_sendFileWithContent(m, v55, v5179, int32(_a_F_perform_base_backup_41), v5188)
	mBase = m.M
	v5229 = m.ExcPending
	if v5229 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L552
	}
L552:
	;
	v5231 = v5020 + int32(1)
	v5232 = *(*int32)(unsafe.Add(mBase, uint32(v3729)+4))
	if v5231 < v5232 {
		v5020 = v5231
		goto L539
	} else {
		goto L553
	}
L553:
	;
	goto L540
L554:
	;
	v5297 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v5298 = *(*int32)(unsafe.Add(mBase, uint32(v5297)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v3725
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v3726
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	m.T0[v5298].(func(*base.Module, int32))(m, v55)
	mBase = m.M
	v5312 = m.ExcPending
	if v5312 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L555
	}
L555:
	;
	v5324 = v3724
	v5325 = v3725
	v5326 = v3726
	goto L345
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v5325
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v5324
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v5326
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v5645 = m.G0
	v5647 = v5645 - int32(128)
	m.G0 = v5647
	v5650 = v57 + int32(2072)
	v5651 = *(*int32)(unsafe.Add(mBase, uint32(v5650)))
	if v5651 != 0 {
		goto L600
	} else {
		goto L601
	}
L557:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5617 = m.ExcPending
	if v5617 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L594
	}
L558:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5600 = m.ExcPending
	if v5600 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L591
	}
L559:
	;
	F_AppendStringToManifest(m, v5375, int32(_a_F_perform_base_backup_48))
	mBase = m.M
	v5379 = m.ExcPending
	if v5379 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L562
	}
L560:
	;
	goto L561
L561:
	;
	m.G0 = v5372 + int32(80)
	goto L556
L562:
	;
	v5380 = F_readTimeLineHistory(m, v3040)
	mBase = m.M
	v5381 = m.ExcPending
	if v5381 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L563
	}
L563:
	;
	F_AppendStringToManifest(m, v5375, int32(_a_F_perform_base_backup_49))
	mBase = m.M
	v5384 = m.ExcPending
	if v5384 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L564
	}
L564:
	;
	if v5380 == int32(0) {
		goto L566
	} else {
		goto L567
	}
L565:
	;
	F_AppendStringToManifest(m, v5375, int32(_a_F_perform_base_backup_48))
	mBase = m.M
	v5551 = m.ExcPending
	if v5551 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L590
	}
L566:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5538 = m.ExcPending
	if v5538 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L587
	}
L567:
	;
	v5388 = *(*int32)(unsafe.Add(mBase, uint32(v5380)+4))
	if v5388 <= int32(0) {
		goto L566
	} else {
		goto L568
	}
L568:
	;
	v5412 = int32(1)
	v5413 = int32(0)
	v5432 = v3041
	goto L569
L569:
	;
	v5437 = *(*int32)(unsafe.Add(mBase, uint32(v5380)+12))
	v5441 = *(*int32)(unsafe.Add(mBase, uint32(v5437+v5413<<(uint(int32(2))%32))))
	v5442 = *(*int64)(unsafe.Add(mBase, uint32(v5441)+16))
	if base.B2i32(v5442 != int64(0))&base.B2i32(base.Ui64(v5442) < base.Ui64(v5356)) == int32(0) {
		goto L571
	} else {
		goto L572
	}
L570:
	;
	goto L566
L571:
	;
	v5449 = *(*int32)(unsafe.Add(mBase, uint32(v5441)))
	if base.B2i32(v3040 != v5449)&v5412 != 0 {
		goto L558
	} else {
		goto L574
	}
L572:
	;
	v5485 = v5412
	v5487 = v5432
	goto L573
L573:
	;
	v5490 = v5413 + int32(1)
	v5491 = *(*int32)(unsafe.Add(mBase, uint32(v5380)+4))
	if v5490 < v5491 {
		v5412 = v5485
		v5413 = v5490
		v5432 = v5487
		goto L569
	} else {
		goto L586
	}
L574:
	;
	if v5355 != v5449 {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	v5453 = *(*int64)(unsafe.Add(mBase, uint32(v5441)+8))
	if v5453 == int64(0) {
		goto L557
	} else {
		goto L578
	}
L576:
	;
	v5456 = v5356
	goto L577
L577:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v5372+int32(52)))) = uint32(v5432)
	v5458 = int64(32)
	v5459 = int64(base.Ui64(v5432) >> (uint(v5458) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v5372+int32(48)))) = uint32(v5459)
	*(*int32)(unsafe.Add(mBase, uint32(v5372)+36)) = v5449
	*(*uint32)(unsafe.Add(mBase, uint32(v5372)+44)) = uint32(v5456)
	v5464 = int64(base.Ui64(v5456) >> (uint(v5458) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v5372)+40)) = uint32(v5464)
	if v5412&int32(1) != 0 {
		goto L579
	} else {
		goto L580
	}
L578:
	;
	v5456 = v5453
	goto L577
L579:
	;
	v5470 = int32(_a_F_perform_base_backup_41)
	goto L581
L580:
	;
	v5470 = int32(_a_F_perform_base_backup_50)
	goto L581
L581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5372)+32)) = v5470
	v5475 = F_psprintf(m, int32(_a_F_perform_base_backup_51), v5372+int32(32))
	mBase = m.M
	v5476 = m.ExcPending
	if v5476 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L582
	}
L582:
	;
	F_AppendStringToManifest(m, v5375, v5475)
	mBase = m.M
	v5478 = m.ExcPending
	if v5478 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L583
	}
L583:
	;
	F_pfree(m, v5475)
	mBase = m.M
	v5480 = m.ExcPending
	if v5480 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L584
	}
L584:
	;
	v5481 = *(*int32)(unsafe.Add(mBase, uint32(v5441)))
	if v5355 == v5481 {
		goto L565
	} else {
		goto L585
	}
L585:
	;
	v5484 = *(*int64)(unsafe.Add(mBase, uint32(v5441)+8))
	v5485 = int32(0)
	v5487 = v5484
	goto L573
L586:
	;
	goto L570
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5372)+4)) = v3040
	*(*int32)(unsafe.Add(mBase, uint32(v5372))) = v5355
	F_errmsg(m, int32(_a_F_perform_base_backup_52), v5372)
	mBase = m.M
	v5543 = m.ExcPending
	if v5543 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L588
	}
L588:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(307), int32(_a_F_perform_base_backup_53))
	mBase = m.M
	v5548 = m.ExcPending
	if v5548 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
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
	goto L561
L591:
	;
	v5601 = *(*int32)(unsafe.Add(mBase, uint32(v5441)))
	*(*int32)(unsafe.Add(mBase, uint32(v5372)+20)) = v5601
	*(*int32)(unsafe.Add(mBase, uint32(v5372)+16)) = v3040
	F_errmsg(m, int32(_a_F_perform_base_backup_54), v5372+int32(16))
	mBase = m.M
	v5608 = m.ExcPending
	if v5608 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L592
	}
L592:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(256), int32(_a_F_perform_base_backup_53))
	mBase = m.M
	v5613 = m.ExcPending
	if v5613 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L593
	}
L593:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L594:
	;
	v5618 = *(*int32)(unsafe.Add(mBase, uint32(v5441)))
	*(*int32)(unsafe.Add(mBase, uint32(v5372)+68)) = v5618
	*(*int32)(unsafe.Add(mBase, uint32(v5372)+64)) = v5355
	F_errmsg(m, int32(_a_F_perform_base_backup_55), v5372-int32(-64))
	mBase = m.M
	v5625 = m.ExcPending
	if v5625 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L595
	}
L595:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(280), int32(_a_F_perform_base_backup_53))
	mBase = m.M
	v5630 = m.ExcPending
	if v5630 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L596
	}
L596:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L597:
	;
	v5907 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v5908 = *(*int32)(unsafe.Add(mBase, uint32(v5907)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v5325
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v5324
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v5326
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	m.T0[v5908].(func(*base.Module, int32, int64, int32))(m, v55, v3041, v3040)
	mBase = m.M
	v5922 = m.ExcPending
	if v5922 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L647
	}
L598:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5895 = m.ExcPending
	if v5895 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L643
	}
L599:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5866 = m.ExcPending
	if v5866 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L630
	}
L600:
	;
	v5652 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5650)+26)) = uint8(v5652)
	v5654 = *(*int32)(unsafe.Add(mBase, uint32(v5650)+8))
	v5656 = v5647 + int32(96)
	v5658 = F_pg_cryptohash_final(m, v5654, v5656, int32(32))
	mBase = m.M
	if v5658 < v5652 {
		goto L599
	} else {
		goto L603
	}
L601:
	;
	goto L602
L602:
	;
	m.G0 = v5647 + int32(128)
	goto L597
L603:
	;
	F_AppendStringToManifest(m, v5650, int32(_a_F_perform_base_backup_56))
	mBase = m.M
	v5663 = m.ExcPending
	if v5663 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L604
	}
L604:
	;
	v5666 = v5647 + int32(16)
	goto L606
L605:
	;
	v5690 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5647)+80)) = uint8(v5690)
	F_AppendStringToManifest(m, v5650, v5666)
	mBase = m.M
	v5693 = m.ExcPending
	if v5693 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L612
	}
L606:
	;
	v5669 = v5656
	v5671 = v5666
	goto L609
L608:
	;
	goto L605
L609:
	;
	v5673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5669))))
	v5674 = int32(1)
	v5676 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5673<<(uint(v5674)%32))+uint32(_c_F_perform_base_backup[20]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5671))) = uint16(v5676)
	v5681 = v5669 + v5674
	if base.Ui32(v5681) < base.Ui32(v5647+int32(128)) {
		v5669 = v5681
		v5671 = v5671 + int32(2)
		goto L609
	} else {
		goto L611
	}
L610:
	;
	goto L608
L611:
	;
	goto L610
L612:
	;
	F_AppendStringToManifest(m, v5650, int32(_a_F_perform_base_backup_57))
	mBase = m.M
	v5696 = m.ExcPending
	if v5696 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L613
	}
L613:
	;
	v5697 = *(*int32)(unsafe.Add(mBase, uint32(v5650)))
	v5698 = int32(0)
	v5701 = F_BufFileSeek(m, v5697, v5698, int64(0), v5698)
	mBase = m.M
	v5702 = m.ExcPending
	if v5702 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L614
	}
L614:
	;
	if v5701 != 0 {
		goto L598
	} else {
		goto L615
	}
L615:
	;
	v5703 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v5704 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+16))
	m.T0[v5704].(func(*base.Module, int32))(m, v55)
	mBase = m.M
	v5706 = m.ExcPending
	if v5706 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L616
	}
L616:
	;
	v5707 = *(*int64)(unsafe.Add(mBase, uint32(v5650)+16))
	if v5707 != int64(0) {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	v5727 = int32(0)
	v5748 = v5707
	v5749 = int64(0)
	goto L620
L618:
	;
	goto L619
L619:
	;
	v5811 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v5812 = *(*int32)(unsafe.Add(mBase, uint32(v5811)+24))
	m.T0[v5812].(func(*base.Module, int32))(m, v55)
	mBase = m.M
	v5814 = m.ExcPending
	if v5814 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L628
	}
L620:
	;
	v5752 = *(*int32)(unsafe.Add(mBase, uint32(v5650)))
	v5753 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v5754 = v5748 - v5749
	v5755 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v55)+8)))
	if base.Ui64(v5754) < base.Ui64(v5755) {
		goto L622
	} else {
		goto L623
	}
L621:
	;
	goto L619
L622:
	;
	v5757 = v5754
	goto L624
L623:
	;
	v5757 = v5755
	goto L624
L624:
	;
	v5758 = base.I32_wrap_i64(v5757)
	F_BufFileReadExact(m, v5752, v5753, v5758)
	mBase = m.M
	v5760 = m.ExcPending
	if v5760 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L625
	}
L625:
	;
	v5761 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v5762 = *(*int32)(unsafe.Add(mBase, uint32(v5761)+20))
	m.T0[v5762].(func(*base.Module, int32, int32))(m, v55, v5758)
	mBase = m.M
	v5764 = m.ExcPending
	if v5764 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L626
	}
L626:
	;
	v5765 = *(*int64)(unsafe.Add(mBase, uint32(v5650)+16))
	v5766 = v5758 + v5727
	v5767 = base.I64_extend_i32_u(v5766)
	if base.Ui64(v5767) < base.Ui64(v5765) {
		v5727 = v5766
		v5748 = v5765
		v5749 = v5767
		goto L620
	} else {
		goto L627
	}
L627:
	;
	goto L621
L628:
	;
	v5815 = *(*int32)(unsafe.Add(mBase, uint32(v5650)))
	F_BufFileClose(m, v5815)
	mBase = m.M
	v5817 = m.ExcPending
	if v5817 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L629
	}
L629:
	;
	goto L602
L630:
	;
	v5867 = *(*int32)(unsafe.Add(mBase, uint32(v5650)+8))
	if v5867 == int32(0) {
		goto L632
	} else {
		goto L633
	}
L631:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5647))) = v5882
	F_errmsg_internal(m, int32(_a_F_perform_base_backup_58), v5647)
	mBase = m.M
	v5886 = m.ExcPending
	if v5886 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L641
	}
L632:
	;
	v5882 = int32(_a_F_perform_base_backup_2)
	goto L631
L633:
	;
	goto L634
L634:
	;
	v5874 = *(*int32)(unsafe.Add(mBase, uint32(v5867)+4))
	if v5874 == int32(1) {
		goto L635
	} else {
		goto L636
	}
L635:
	;
	v5877 = int32(_a_F_perform_base_backup_3)
	goto L637
L636:
	;
	v5877 = int32(_a_F_perform_base_backup_4)
	goto L637
L637:
	;
	if v5874 == int32(2) {
		goto L638
	} else {
		goto L639
	}
L638:
	;
	v5880 = int32(_a_F_perform_base_backup_2)
	goto L640
L639:
	;
	v5880 = v5877
	goto L640
L640:
	;
	v5882 = v5880
	goto L631
L641:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(341), int32(_a_F_perform_base_backup_59))
	mBase = m.M
	v5891 = m.ExcPending
	if v5891 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L642
	}
L642:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L643:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v5897 = m.ExcPending
	if v5897 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L644
	}
L644:
	;
	F_errmsg(m, int32(_a_F_perform_base_backup_60), int32(0))
	mBase = m.M
	v5901 = m.ExcPending
	if v5901 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L645
	}
L645:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(357), int32(_a_F_perform_base_backup_59))
	mBase = m.M
	v5906 = m.ExcPending
	if v5906 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L646
	}
L646:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L647:
	;
	v5924 = *(*int64)(unsafe.Add(mBase, _c_F_perform_base_backup[5]))
	if v5924 != int64(0) {
		goto L648
	} else {
		goto L649
	}
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v5325
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v5324
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v5326
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v5941 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5942 = m.ExcPending
	if v5942 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L651
	}
L649:
	;
	goto L650
L650:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v5325
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v5324
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v5326
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v6060 = v57 + int32(2072)
	v6061 = *(*int32)(unsafe.Add(mBase, uint32(v6060)+8))
	F_pg_cryptohash_free(m, v6061)
	mBase = m.M
	v6063 = m.ExcPending
	if v6063 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L661
	}
L651:
	;
	if v5941 != 0 {
		goto L652
	} else {
		goto L653
	}
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v5324
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v5325
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v5326
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	v5956 = *(*int64)(unsafe.Add(mBase, _c_F_perform_base_backup[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+32)) = v5956
	F_errmsg_plural(m, int32(_a_F_perform_base_backup_61), int32(_a_F_perform_base_backup_62), base.I32_wrap_i64(v5956), v57+int32(32))
	mBase = m.M
	v5964 = m.ExcPending
	if v5964 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L655
	}
L653:
	;
	goto L654
L654:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v5325
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v5324
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v5326
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5998 = m.ExcPending
	if v5998 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L657
	}
L655:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v5325
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v5324
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v5326
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(662), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v5981 = m.ExcPending
	if v5981 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L656
	}
L656:
	;
	goto L654
L657:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v5325
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v5324
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v5326
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errcode(m, int32(16779816))
	mBase = m.M
	v6013 = m.ExcPending
	if v6013 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L658
	}
L658:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v5325
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v5324
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v5326
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errmsg(m, int32(_a_F_perform_base_backup_63), int32(0))
	mBase = m.M
	v6029 = m.ExcPending
	if v6029 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L659
	}
L659:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v5325
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v5324
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v5326
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(666), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v6046 = m.ExcPending
	if v6046 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L660
	}
L660:
	;
	goto L1
L661:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6060)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2184)) = v5325
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2180)) = v5324
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2188)) = v5326
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2192)) = v2308
	*(*int64)(unsafe.Add(mBase, uint32(v57)+2200)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2212)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2216)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2220)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2224)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2228)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2232)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v57)+2236)) = v385
	F_ReleaseAuxProcessResources(m, int32(1))
	mBase = m.M
	v6080 = m.ExcPending
	if v6080 != 0 {
		v6081 = v54
		v6082 = v55
		v6083 = v56
		v6084 = v57
		v6107 = v80
		v6111 = v84
		v6112 = v85
		goto L6
	} else {
		goto L662
	}
L662:
	;
	goto L5
L663:
	;
	v6128 = int32(v6124)
	m.G0 = v6084
	v6130 = *(*int32)(unsafe.Add(mBase, uint32(v6128)+4))
	v6131 = *(*int32)(unsafe.Add(mBase, uint32(v6128)))
	v6134 = *(*int32)(unsafe.Add(mBase, uint32(v6131)))
	if v6084+int32(332) == v6134 {
		goto L666
	} else {
		goto L667
	}
L664:
	;
	m.ExcPending = 1
	goto L672
L665:
	;
	if v6138 != 0 {
		goto L669
	} else {
		goto L670
	}
L666:
	;
	v6136 = *(*int32)(unsafe.Add(mBase, uint32(v6131)+4))
	v6138 = v6136
	goto L668
L667:
	;
	v6138 = int32(0)
	goto L668
L668:
	;
	goto L665
L669:
	;
	v6139 = *(*int32)(unsafe.Add(mBase, uint32(v6084)+2236))
	v6140 = *(*int32)(unsafe.Add(mBase, uint32(v6084)+2232))
	v6141 = *(*int32)(unsafe.Add(mBase, uint32(v6084)+2228))
	v6142 = *(*int32)(unsafe.Add(mBase, uint32(v6084)+2224))
	v6143 = *(*int32)(unsafe.Add(mBase, uint32(v6084)+2220))
	v6144 = *(*int32)(unsafe.Add(mBase, uint32(v6084)+2216))
	v6145 = *(*int32)(unsafe.Add(mBase, uint32(v6084)+2212))
	v6146 = *(*int64)(unsafe.Add(mBase, uint32(v6084)+2200))
	v6147 = *(*int64)(unsafe.Add(mBase, uint32(v6084)+2192))
	v6148 = *(*int32)(unsafe.Add(mBase, uint32(v6084)+2188))
	v6149 = *(*int32)(unsafe.Add(mBase, uint32(v6084)+2184))
	v6150 = *(*int32)(unsafe.Add(mBase, uint32(v6084)+2180))
	v54 = v6081
	v55 = v6082
	v56 = v6083
	v57 = v6084
	v58 = v6130
	v59 = v6143
	v60 = v6141
	v61 = v6139
	v62 = v6145
	v63 = v6142
	v64 = v6144
	v65 = v6150
	v66 = v6149
	v67 = v6148
	v68 = v6140
	v71 = v6138
	v80 = v6107
	v84 = v6111
	v85 = v6112
	v89 = v6147
	v90 = v6146
	goto L2
L670:
	;
	goto L671
L671:
	;
	F___wasm_longjmp(m, v6131, v6130)
	mBase = m.M
	v6152 = m.ExcPending
	if v6152 != 0 {
		goto L672
	} else {
		goto L673
	}
L672:
	;
	return
L673:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
