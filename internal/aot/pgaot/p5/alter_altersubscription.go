package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AlterSubscription(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v159 int32
	_ = v159
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v220 int32
	_ = v220
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v261 int32
	_ = v261
	var v274 int32
	_ = v274
	var v288 int32
	_ = v288
	var v302 int32
	_ = v302
	var v317 int32
	_ = v317
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v339 int64
	_ = v339
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v402 int32
	_ = v402
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v474 int32
	_ = v474
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v583 int32
	_ = v583
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v670 int32
	_ = v670
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v740 int32
	_ = v740
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v810 int32
	_ = v810
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v880 int32
	_ = v880
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v950 int32
	_ = v950
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1020 int32
	_ = v1020
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1090 int32
	_ = v1090
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1160 int32
	_ = v1160
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1218 int32
	_ = v1218
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1241 int32
	_ = v1241
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1273 int32
	_ = v1273
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1296 int32
	_ = v1296
	var v1312 int32
	_ = v1312
	var v1325 int32
	_ = v1325
	var v1336 int32
	_ = v1336
	var v1342 int32
	_ = v1342
	var v1357 int32
	_ = v1357
	var v1371 int32
	_ = v1371
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1401 int32
	_ = v1401
	var v1416 int32
	_ = v1416
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1474 int32
	_ = v1474
	var v1490 int32
	_ = v1490
	var v1503 int32
	_ = v1503
	var v1524 int32
	_ = v1524
	var v1539 int32
	_ = v1539
	var v1558 int32
	_ = v1558
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1590 int32
	_ = v1590
	var v1603 int32
	_ = v1603
	var v1624 int32
	_ = v1624
	var v1639 int32
	_ = v1639
	var v1658 int32
	_ = v1658
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1690 int32
	_ = v1690
	var v1703 int32
	_ = v1703
	var v1720 int32
	_ = v1720
	var v1735 int32
	_ = v1735
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1757 int32
	_ = v1757
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1774 int32
	_ = v1774
	var v1777 int32
	_ = v1777
	var v1779 int32
	_ = v1779
	var v1783 int32
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1796 int32
	_ = v1796
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1828 int32
	_ = v1828
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1871 int32
	_ = v1871
	var v1884 int32
	_ = v1884
	var v1898 int32
	_ = v1898
	var v1913 int32
	_ = v1913
	var v1927 int32
	_ = v1927
	var v1940 int32
	_ = v1940
	var v1954 int32
	_ = v1954
	var v1968 int32
	_ = v1968
	var v1983 int32
	_ = v1983
	var v1994 int32
	_ = v1994
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2027 int32
	_ = v2027
	var v2040 int32
	_ = v2040
	var v2054 int32
	_ = v2054
	var v2068 int32
	_ = v2068
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2104 int32
	_ = v2104
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2131 int32
	_ = v2131
	var v2134 int32
	_ = v2134
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2159 int32
	_ = v2159
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2167 int32
	_ = v2167
	var v2169 int32
	_ = v2169
	var v2175 int32
	_ = v2175
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2182 int32
	_ = v2182
	var v2185 int32
	_ = v2185
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2196 int32
	_ = v2196
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2234 int32
	_ = v2234
	var v2247 int32
	_ = v2247
	var v2251 int32
	_ = v2251
	var v2258 int32
	_ = v2258
	var v2261 int32
	_ = v2261
	var v2265 int32
	_ = v2265
	var v2270 int32
	_ = v2270
	var v2286 int32
	_ = v2286
	var v2299 int32
	_ = v2299
	var v2313 int32
	_ = v2313
	var v2327 int32
	_ = v2327
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2344 int64
	_ = v2344
	var v2354 int32
	_ = v2354
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2378 int32
	_ = v2378
	var v2393 int32
	_ = v2393
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2409 int32
	_ = v2409
	var v2412 int32
	_ = v2412
	var v2415 int32
	_ = v2415
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
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2449 int32
	_ = v2449
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2479 int32
	_ = v2479
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2509 int32
	_ = v2509
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2530 int32
	_ = v2530
	var v2546 int32
	_ = v2546
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2551 int32
	_ = v2551
	var v2554 int32
	_ = v2554
	var v2557 int32
	_ = v2557
	var v2558 int32
	_ = v2558
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2589 int32
	_ = v2589
	var v2603 int32
	_ = v2603
	var v2604 int32
	_ = v2604
	var v2615 int32
	_ = v2615
	var v2616 int32
	_ = v2616
	var v2618 int32
	_ = v2618
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2625 int32
	_ = v2625
	var v2641 int32
	_ = v2641
	var v2654 int32
	_ = v2654
	var v2668 int32
	_ = v2668
	var v2682 int32
	_ = v2682
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2701 int32
	_ = v2701
	var v2719 int32
	_ = v2719
	var v2732 int32
	_ = v2732
	var v2746 int32
	_ = v2746
	var v2760 int32
	_ = v2760
	var v2775 int32
	_ = v2775
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2801 int32
	_ = v2801
	var v2803 int32
	_ = v2803
	var v2805 int32
	_ = v2805
	var v2819 int32
	_ = v2819
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2822 int32
	_ = v2822
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2847 int32
	_ = v2847
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2871 int32
	_ = v2871
	var v2874 int32
	_ = v2874
	var v2884 int32
	_ = v2884
	var v2888 int32
	_ = v2888
	var v2889 int32
	_ = v2889
	var v2892 int32
	_ = v2892
	var v2895 int32
	_ = v2895
	var v2912 int32
	_ = v2912
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2944 int32
	_ = v2944
	var v2947 int32
	_ = v2947
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2954 int32
	_ = v2954
	var v2955 int32
	_ = v2955
	var v2958 int32
	_ = v2958
	var v2965 int32
	_ = v2965
	var v2966 int32
	_ = v2966
	var v2985 int32
	_ = v2985
	var v2998 int32
	_ = v2998
	var v3015 int32
	_ = v3015
	var v3030 int32
	_ = v3030
	var v3032 int32
	_ = v3032
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3103 int32
	_ = v3103
	var v3116 int32
	_ = v3116
	var v3133 int32
	_ = v3133
	var v3148 int32
	_ = v3148
	var v3159 int32
	_ = v3159
	var v3160 int32
	_ = v3160
	var v3174 int32
	_ = v3174
	var v3175 int32
	_ = v3175
	var v3191 int32
	_ = v3191
	var v3193 int32
	_ = v3193
	var v3194 int32
	_ = v3194
	var v3209 int32
	_ = v3209
	var v3210 int32
	_ = v3210
	var v3213 int32
	_ = v3213
	var v3241 int32
	_ = v3241
	var v3254 int32
	_ = v3254
	var v3268 int32
	_ = v3268
	var v3283 int32
	_ = v3283
	var v3295 int32
	_ = v3295
	var v3300 int32
	_ = v3300
	var v3301 int32
	_ = v3301
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3306 int32
	_ = v3306
	var v3319 int32
	_ = v3319
	var v3320 int32
	_ = v3320
	var v3332 int32
	_ = v3332
	var v3335 int32
	_ = v3335
	var v3347 int32
	_ = v3347
	var v3348 int32
	_ = v3348
	var v3360 int32
	_ = v3360
	var v3362 int32
	_ = v3362
	var v3364 int32
	_ = v3364
	var v3370 int32
	_ = v3370
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3388 int32
	_ = v3388
	var v3401 int32
	_ = v3401
	var v3415 int32
	_ = v3415
	var v3430 int32
	_ = v3430
	var v3436 int32
	_ = v3436
	var v3451 int32
	_ = v3451
	var v3452 int32
	_ = v3452
	var v3455 int32
	_ = v3455
	var v3473 int32
	_ = v3473
	var v3486 int32
	_ = v3486
	var v3500 int32
	_ = v3500
	var v3515 int32
	_ = v3515
	var v3521 int32
	_ = v3521
	var v3536 int32
	_ = v3536
	var v3549 int32
	_ = v3549
	var v3561 int32
	_ = v3561
	var v3563 int32
	_ = v3563
	var v3564 int32
	_ = v3564
	var v3580 int32
	_ = v3580
	var v3593 int32
	_ = v3593
	var v3607 int32
	_ = v3607
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3624 int64
	_ = v3624
	var v3634 int32
	_ = v3634
	var v3639 int32
	_ = v3639
	var v3640 int32
	_ = v3640
	var v3658 int32
	_ = v3658
	var v3673 int32
	_ = v3673
	var v3677 int32
	_ = v3677
	var v3678 int32
	_ = v3678
	var v3689 int32
	_ = v3689
	var v3692 int32
	_ = v3692
	var v3695 int32
	_ = v3695
	var v3698 int32
	_ = v3698
	var v3699 int32
	_ = v3699
	var v3702 int32
	_ = v3702
	var v3703 int32
	_ = v3703
	var v3706 int32
	_ = v3706
	var v3713 int32
	_ = v3713
	var v3714 int32
	_ = v3714
	var v3716 int32
	_ = v3716
	var v3730 int32
	_ = v3730
	var v3744 int32
	_ = v3744
	var v3745 int32
	_ = v3745
	var v3748 int32
	_ = v3748
	var v3749 int32
	_ = v3749
	var v3751 int32
	_ = v3751
	var v3768 int32
	_ = v3768
	var v3782 int32
	_ = v3782
	var v3798 int32
	_ = v3798
	var v3811 int32
	_ = v3811
	var v3812 int32
	_ = v3812
	var v3828 int32
	_ = v3828
	var v3843 int32
	_ = v3843
	var v3844 int32
	_ = v3844
	var v3860 int32
	_ = v3860
	var v3873 int32
	_ = v3873
	var v3894 int32
	_ = v3894
	var v3909 int32
	_ = v3909
	var v3928 int32
	_ = v3928
	var v3943 int32
	_ = v3943
	var v3944 int32
	_ = v3944
	var v3960 int32
	_ = v3960
	var v3973 int32
	_ = v3973
	var v3994 int32
	_ = v3994
	var v4009 int32
	_ = v4009
	var v4028 int32
	_ = v4028
	var v4043 int32
	_ = v4043
	var v4044 int32
	_ = v4044
	var v4045 int64
	_ = v4045
	var v4058 int32
	_ = v4058
	var v4059 int32
	_ = v4059
	var v4077 int32
	_ = v4077
	var v4092 int32
	_ = v4092
	var v4096 int32
	_ = v4096
	var v4097 int32
	_ = v4097
	var v4108 int32
	_ = v4108
	var v4111 int32
	_ = v4111
	var v4114 int32
	_ = v4114
	var v4117 int32
	_ = v4117
	var v4118 int32
	_ = v4118
	var v4121 int32
	_ = v4121
	var v4122 int32
	_ = v4122
	var v4125 int32
	_ = v4125
	var v4132 int32
	_ = v4132
	var v4133 int32
	_ = v4133
	var v4145 int32
	_ = v4145
	var v4146 int32
	_ = v4146
	var v4147 int32
	_ = v4147
	var v4161 int32
	_ = v4161
	var v4172 int32
	_ = v4172
	var v4175 int32
	_ = v4175
	var v4178 int32
	_ = v4178
	var v4181 int32
	_ = v4181
	var v4182 int32
	_ = v4182
	var v4185 int32
	_ = v4185
	var v4186 int32
	_ = v4186
	var v4189 int32
	_ = v4189
	var v4196 int32
	_ = v4196
	var v4197 int32
	_ = v4197
	var v4214 int32
	_ = v4214
	var v4215 int32
	_ = v4215
	var v4216 int64
	_ = v4216
	var v4219 int32
	_ = v4219
	var v4220 int32
	_ = v4220
	var v4221 int64
	_ = v4221
	var v4227 int32
	_ = v4227
	var v4228 int32
	_ = v4228
	var v4243 int32
	_ = v4243
	var v4256 int32
	_ = v4256
	var v4272 int32
	_ = v4272
	var v4287 int32
	_ = v4287
	var v4301 int32
	_ = v4301
	var v4314 int32
	_ = v4314
	var v4315 int32
	_ = v4315
	var v4331 int32
	_ = v4331
	var v4346 int32
	_ = v4346
	var v4360 int32
	_ = v4360
	var v4361 int32
	_ = v4361
	var v4377 int32
	_ = v4377
	var v4392 int32
	_ = v4392
	var v4393 int32
	_ = v4393
	var v4399 int32
	_ = v4399
	var v4415 int32
	_ = v4415
	var v4428 int32
	_ = v4428
	var v4449 int32
	_ = v4449
	var v4464 int32
	_ = v4464
	var v4483 int32
	_ = v4483
	var v4498 int32
	_ = v4498
	var v4499 int32
	_ = v4499
	var v4515 int32
	_ = v4515
	var v4528 int32
	_ = v4528
	var v4549 int32
	_ = v4549
	var v4564 int32
	_ = v4564
	var v4583 int32
	_ = v4583
	var v4598 int32
	_ = v4598
	var v4614 int32
	_ = v4614
	var v4616 int32
	_ = v4616
	var v4628 int32
	_ = v4628
	var v4629 int32
	_ = v4629
	var v4641 int64
	_ = v4641
	var v4642 int32
	_ = v4642
	var v4645 int64
	_ = v4645
	var v4661 int32
	_ = v4661
	var v4674 int32
	_ = v4674
	var v4685 int64
	_ = v4685
	var v4687 int64
	_ = v4687
	var v4688 int64
	_ = v4688
	var v4692 int64
	_ = v4692
	var v4698 int32
	_ = v4698
	var v4713 int32
	_ = v4713
	var v4742 int64
	_ = v4742
	var v4754 int32
	_ = v4754
	var v4755 int32
	_ = v4755
	var v4756 int32
	_ = v4756
	var v4775 int32
	_ = v4775
	var v4789 int32
	_ = v4789
	var v4791 int32
	_ = v4791
	var v4810 int32
	_ = v4810
	var v4823 int32
	_ = v4823
	var v4837 int32
	_ = v4837
	var v4851 int32
	_ = v4851
	var v4866 int32
	_ = v4866
	var v4879 int32
	_ = v4879
	var v4890 int32
	_ = v4890
	var v4891 int32
	_ = v4891
	var v4894 int32
	_ = v4894
	var v4909 int32
	_ = v4909
	var v4910 int32
	_ = v4910
	var v4926 int32
	_ = v4926
	var v4928 int32
	_ = v4928
	var v4929 int32
	_ = v4929
	var v4951 int32
	_ = v4951
	var v4979 int32
	_ = v4979
	var v4992 int32
	_ = v4992
	var v5013 int32
	_ = v5013
	var v5028 int32
	_ = v5028
	var v5047 int32
	_ = v5047
	var v5062 int32
	_ = v5062
	var v5063 int32
	_ = v5063
	var v5079 int32
	_ = v5079
	var v5092 int32
	_ = v5092
	var v5113 int32
	_ = v5113
	var v5128 int32
	_ = v5128
	var v5147 int32
	_ = v5147
	var v5162 int32
	_ = v5162
	var v5179 int32
	_ = v5179
	var v5193 int32
	_ = v5193
	var v5194 int32
	_ = v5194
	var v5212 int32
	_ = v5212
	var v5225 int32
	_ = v5225
	var v5239 int32
	_ = v5239
	var v5254 int32
	_ = v5254
	var v5255 int32
	_ = v5255
	var v5258 int32
	_ = v5258
	var v5273 int32
	_ = v5273
	var v5277 int32
	_ = v5277
	var v5310 int32
	_ = v5310
	var v5314 int32
	_ = v5314
	var v5315 int32
	_ = v5315
	var v5317 int32
	_ = v5317
	var v5334 int32
	_ = v5334
	var v5335 int32
	_ = v5335
	var v5349 int32
	_ = v5349
	var v5363 int32
	_ = v5363
	var v5364 int32
	_ = v5364
	var v5366 int32
	_ = v5366
	var v5368 int32
	_ = v5368
	var v5369 int32
	_ = v5369
	var v5371 int32
	_ = v5371
	var v5386 int32
	_ = v5386
	var v5387 int32
	_ = v5387
	var v5388 int32
	_ = v5388
	var v5389 int32
	_ = v5389
	var v5405 int32
	_ = v5405
	var v5406 int32
	_ = v5406
	var v5407 int32
	_ = v5407
	var v5409 int32
	_ = v5409
	var v5422 int32
	_ = v5422
	var v5439 int32
	_ = v5439
	var v5440 int32
	_ = v5440
	var v5454 int32
	_ = v5454
	var v5466 int32
	_ = v5466
	var v5480 int32
	_ = v5480
	var v5481 int32
	_ = v5481
	var v5482 int32
	_ = v5482
	var v5484 int32
	_ = v5484
	var v5498 int32
	_ = v5498
	var v5515 int32
	_ = v5515
	var v5517 int32
	_ = v5517
	var v5520 int32
	_ = v5520
	var v5523 int32
	_ = v5523
	var v5524 int32
	_ = v5524
	var v5525 int32
	_ = v5525
	var v5527 int32
	_ = v5527
	var v5528 int32
	_ = v5528
	var v5539 int32
	_ = v5539
	var v5545 int32
	_ = v5545
	var v5546 int32
	_ = v5546
	var v5562 int32
	_ = v5562
	var v5575 int32
	_ = v5575
	var v5576 int32
	_ = v5576
	var v5587 int32
	_ = v5587
	var v5594 int32
	_ = v5594
	var v5609 int32
	_ = v5609
	var v5612 int32
	_ = v5612
	var v5614 int32
	_ = v5614
	var v5616 int32
	_ = v5616
	var v5622 int32
	_ = v5622
	var v5623 int32
	_ = v5623
	var v5624 int32
	_ = v5624
	var v5625 int32
	_ = v5625
	var v5626 int32
	_ = v5626
	var v5627 int32
	_ = v5627
	var v5628 int32
	_ = v5628
	var v5629 int32
	_ = v5629
	var v5630 int32
	_ = v5630
	var v5631 int32
	_ = v5631
	var v5632 int32
	_ = v5632
	var v5633 int32
	_ = v5633
	var v5635 int32
	_ = v5635
	var v5636 int32
	_ = v5636
	var v5637 int32
	_ = v5637
	var v5639 int32
	_ = v5639
	var v5645 int32
	_ = v5645
	var v5646 int32
	_ = v5646
	var v5656 int32
	_ = v5656
	var v5658 int32
	_ = v5658
	var v5659 int32
	_ = v5659
	var v5668 int32
	_ = v5668
	var v5669 int32
	_ = v5669
	var v5672 int32
	_ = v5672
	var v5675 int32
	_ = v5675
	var v5677 int32
	_ = v5677
	var v5679 int32
	_ = v5679
	var v5685 int32
	_ = v5685
	var v5686 int32
	_ = v5686
	var v5698 int32
	_ = v5698
	var v5703 int32
	_ = v5703
	var v5704 int32
	_ = v5704
	var v5705 int32
	_ = v5705
	var v5706 int32
	_ = v5706
	var v5707 int32
	_ = v5707
	var v5708 int32
	_ = v5708
	var v5709 int32
	_ = v5709
	var v5711 int32
	_ = v5711
	var v5712 int32
	_ = v5712
	var v5713 int32
	_ = v5713
	var v5714 int32
	_ = v5714
	var v5716 int32
	_ = v5716
	var v5717 int32
	_ = v5717
	var v5718 int32
	_ = v5718
	var v5720 int32
	_ = v5720
	var v5726 int32
	_ = v5726
	var v5727 int32
	_ = v5727
	var v5741 int32
	_ = v5741
	var v5742 int32
	_ = v5742
	var v5745 int32
	_ = v5745
	var v5749 int32
	_ = v5749
	var v5756 int32
	_ = v5756
	var v5768 int32
	_ = v5768
	var v5772 int32
	_ = v5772
	var v5784 int32
	_ = v5784
	var v5793 int32
	_ = v5793
	var v5794 int32
	_ = v5794
	var v5798 int32
	_ = v5798
	var v5799 int32
	_ = v5799
	var v5802 int32
	_ = v5802
	var v5810 int32
	_ = v5810
	var v5822 int32
	_ = v5822
	var v5823 int32
	_ = v5823
	var v5824 int32
	_ = v5824
	var v5825 int32
	_ = v5825
	var v5826 int32
	_ = v5826
	var v5827 int32
	_ = v5827
	var v5846 int32
	_ = v5846
	var v5847 int32
	_ = v5847
	var v5853 int32
	_ = v5853
	var v5854 int64
	_ = v5854
	var v5858 int32
	_ = v5858
	var v5860 int32
	_ = v5860
	var v5861 int32
	_ = v5861
	var v5864 int32
	_ = v5864
	var v5866 int32
	_ = v5866
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
	var v5880 int32
	_ = v5880
	v5 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(944)
	m.G0 = v33
	v40 = l0
	v41 = l1
	v42 = l2
	v43 = l3
	v44 = v33
	v45 = v5
	v46 = v5
	v47 = v5
	v48 = v5
	v49 = v5
	v50 = v5
	v51 = v5
	v52 = int32(-1)
	v53 = v5
	v54 = v5
	v55 = v5
	v57 = v5
	v63 = v33 + int32(747)
	v64 = v33 + int32(751)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v52 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L3
L6:
	;
	v5853 = int32(m.ExcTag)
	v5854 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v5853 == int32(0) {
		goto L754
	} else {
		goto L755
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[0])) = v5630
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[1])) = v5631
	v5793 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[2]))
	v5794 = *(*int32)(unsafe.Add(mBase, uint32(v5793)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+908)) = v5630
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+912)) = v5631
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+916)) = v5632
	v5798 = int32(1)
	v5799 = v5639 & v5798
	*(*uint8)(unsafe.Add(mBase, uint32(v5626)+922)) = uint8(v5799)
	v5802 = v5637 & v5798
	*(*uint8)(unsafe.Add(mBase, uint32(v5626)+923)) = uint8(v5802)
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+924)) = v5635
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+928)) = v5636
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+932)) = v5627
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+936)) = v5628
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+940)) = v5633
	m.T0[v5794].(func(*base.Module, int32))(m, v5632)
	mBase = m.M
	v5810 = m.ExcPending
	if v5810 != 0 {
		v5823 = v5622
		v5824 = v5623
		v5825 = v5624
		v5826 = v5625
		v5827 = v5626
		v5846 = v5645
		v5847 = v5646
		goto L6
	} else {
		goto L752
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+912)) = v5712
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+908)) = v5711
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+916)) = v5713
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+924)) = v5716
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+928)) = v5717
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+932)) = v5708
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+936)) = v5709
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+940)) = v5714
	v5741 = int32(1)
	v5742 = v5720 & v5741
	*(*uint8)(unsafe.Add(mBase, uint32(v5707)+922)) = uint8(v5742)
	v5745 = v5718 & v5741
	*(*uint8)(unsafe.Add(mBase, uint32(v5707)+923)) = uint8(v5745)
	F_relation_close(m, v5714, int32(3))
	mBase = m.M
	v5749 = m.ExcPending
	if v5749 != 0 {
		v5823 = v5703
		v5824 = v5704
		v5825 = v5705
		v5826 = v5706
		v5827 = v5707
		v5846 = v5726
		v5847 = v5727
		goto L6
	} else {
		goto L746
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v46
	v79 = int32(1)
	v80 = v57 & v79
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	v83 = v55 & v79
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v85 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+760)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v44)+752)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v44)+744)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v44)+736)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v44)+728)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v51
	v98 = F_table_open(m, int32(_a_F_AlterSubscription_0), int32(3))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v5622 = v40
	v5623 = v41
	v5624 = v42
	v5625 = v43
	v5626 = v44
	v5627 = v45
	v5628 = v46
	v5629 = v47
	v5630 = v48
	v5631 = v49
	v5632 = v50
	v5633 = v51
	v5635 = v53
	v5636 = v54
	v5637 = v55
	v5639 = v57
	v5645 = v63
	v5646 = v64
	goto L11
L11:
	;
	if v5629 != 0 {
		goto L7
	} else {
		goto L737
	}
L12:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[3]))
	v114 = F_SearchSysCacheCopy(m, int32(66), v113, v100)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v114 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+22)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175+v176)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v202 = F_object_ownercheck(m, int32(_a_F_AlterSubscription_0), v178, v190)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L21
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(67137668))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v145
	F_errmsg(m, int32(_a_F_AlterSubscription_1), v44)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1128), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L20
	}
L20:
	;
	goto L3
L21:
	;
	if v202 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	F_aclcheck_error(m, int32(2), int32(38), v206)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v233 = F_GetSubscription(m, v178, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_LockSharedObject(m, int32(_a_F_AlterSubscription_0), v178, int32(8))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L36
	}
L27:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+30)))
	if v235 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v246 = F_superuser(m)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L29
	}
L29:
	;
	if v246 != 0 {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16797828))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errmsg(m, int32(_a_F_AlterSubscription_4), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errhint(m, int32(_a_F_AlterSubscription_5), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1148), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L3
L36:
	;
	v334 = int32(0)
	base.MemoryFill(m, v44+int32(768), v334, int32(72))
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+896)) = uint16(v334)
	v339 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+888)) = v339
	*(*int64)(unsafe.Add(mBase, uint32(v44)+880)) = v339
	*(*int64)(unsafe.Add(mBase, uint32(v44)+848)) = v339
	*(*int64)(unsafe.Add(mBase, uint32(v44)+856)) = v339
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+864)) = uint16(v334)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	switch v349 {
	case 0:
		goto L56
	case 1:
		goto L54
	case 2:
		goto L53
	case 3, 4:
		goto L52
	case 5:
		goto L51
	case 6:
		goto L55
	case 7:
		goto L50
	default:
		goto L49
	}
L37:
	;
	v5498 = v5484 & int32(1)
	if v5482|v5498 == int32(0) {
		v5703 = v40
		v5704 = v41
		v5705 = v42
		v5706 = v43
		v5707 = v44
		v5708 = v233
		v5709 = v178
		v5711 = v48
		v5712 = v49
		v5713 = v50
		v5714 = v98
		v5716 = v5480
		v5717 = v5481
		v5718 = v5482
		v5720 = v5484
		v5726 = v63
		v5727 = v64
		goto L8
	} else {
		goto L720
	}
L38:
	;
	v5422 = *(*int32)(unsafe.Add(mBase, uint32(v98)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v5405
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v5406
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v5439 = F_heap_modify_tuple(m, v114, v5422, v44+int32(768), v44+int32(880), v44+int32(848))
	mBase = m.M
	v5440 = m.ExcPending
	if v5440 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L717
	}
L39:
	;
	v5349 = v5334 & int32(_a_F_AlterSubscription_6)
	if v5349 != 0 {
		goto L711
	} else {
		goto L712
	}
L40:
	;
	v5310 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+856)) = uint8(v5310)
	v5314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+747)))
	if v5314 != 0 {
		goto L708
	} else {
		goto L709
	}
L41:
	;
	v5193 = *(*int32)(unsafe.Add(mBase, uint32(v233)+40))
	v5194 = int32(0)
	if v5193|base.B2i32(v5179 == v5194) == v5194 {
		goto L694
	} else {
		goto L695
	}
L42:
	;
	v5063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+742)))
	if v5063 != int32(1) {
		v5179 = v2523
		goto L41
	} else {
		goto L684
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4979 = m.ExcPending
	if v4979 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L675
	}
L44:
	;
	v4928 = *(*int32)(unsafe.Add(mBase, uint32(v44)+732))
	if v4928 != 0 {
		v5179 = v2523
		goto L41
	} else {
		goto L672
	}
L45:
	;
	v4926 = int32(0)
	v5405 = v4909
	v5406 = v4910
	v5407 = v4926
	v5409 = v4926
	goto L38
L46:
	;
	v4789 = int32(0)
	v4791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+28)))
	if base.B2i32(v4775 == v4789)|base.B2i32(v4791 != int32(101)) == v4789 {
		goto L662
	} else {
		goto L663
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v4754 = F_Int64GetDatum(m, v4742)
	mBase = m.M
	v4755 = m.ExcPending
	if v4755 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L661
	}
L48:
	;
	v4393 = *(*int32)(unsafe.Add(mBase, uint32(v44)+732))
	if v4393|base.B2i32(v4220&int32(8) == int32(0)) != 0 {
		goto L626
	} else {
		goto L627
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4360 = m.ExcPending
	if v4360 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L623
	}
L50:
	;
	v4044 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	v4045 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+760)) = v4045
	*(*int64)(unsafe.Add(mBase, uint32(v44)+752)) = v4045
	*(*int64)(unsafe.Add(mBase, uint32(v44)+744)) = v4045
	*(*int64)(unsafe.Add(mBase, uint32(v44)+736)) = v4045
	*(*int64)(unsafe.Add(mBase, uint32(v44)+728)) = v4045
	if v4044 == int32(0) {
		v4742 = v4045
		goto L47
	} else {
		goto L582
	}
L51:
	;
	v3564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+25)))
	if v3564 == int32(0) {
		goto L525
	} else {
		goto L526
	}
L52:
	;
	v2805 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	F_parse_subscription_options_x2especialized_x2e1(m, v41, v2805, v44+int32(728))
	mBase = m.M
	v2819 = m.ExcPending
	if v2819 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L443
	}
L53:
	;
	v2589 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	F_parse_subscription_options_x2especialized_x2e1(m, v41, v2589, v44+int32(728))
	mBase = m.M
	v2603 = m.ExcPending
	if v2603 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L420
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_load_file(m, int32(_a_F_AlterSubscription_7), int32(0))
	mBase = m.M
	v2546 = m.ExcPending
	if v2546 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L414
	}
L55:
	;
	v2343 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	v2344 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+736)) = v2344
	*(*int64)(unsafe.Add(mBase, uint32(v44)+760)) = v2344
	*(*int64)(unsafe.Add(mBase, uint32(v44)+752)) = v2344
	*(*int64)(unsafe.Add(mBase, uint32(v44)+744)) = v2344
	*(*int64)(unsafe.Add(mBase, uint32(v44)+728)) = v2344
	v2354 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+741)) = uint8(v2354)
	if v2343 == int32(0) {
		v5179 = v2354
		goto L41
	} else {
		goto L387
	}
L56:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	v351 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+744)) = v351
	*(*int64)(unsafe.Add(mBase, uint32(v44)+760)) = v351
	*(*int64)(unsafe.Add(mBase, uint32(v44)+752)) = v351
	*(*int64)(unsafe.Add(mBase, uint32(v44)+736)) = v351
	*(*int64)(unsafe.Add(mBase, uint32(v44)+728)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v369 = int32(112)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+746)) = uint8(v369)
	v371 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+749)) = uint8(v371)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v376 = F_pstrdup(m, int32(_a_F_AlterSubscription_8))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+752)) = v376
	if v350 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	v1468 = v1466 & int32(8)
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v44)+732))
	if v1469 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L59:
	;
	v381 = int32(0)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	if v382 <= v381 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v402 = v381
	goto L61
L61:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v415+v402<<(uint(int32(2))%32))))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v431 = int32(_a_F_AlterSubscription_9)
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v437 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[5])))
	if base.B2i32(v434 == int32(0))|base.B2i32(v434 != v437) != 0 {
		v455 = v434
		v456 = v437
		goto L66
	} else {
		goto L67
	}
L62:
	;
	goto L58
L63:
	;
	v1433 = v402 + int32(1)
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	if v1433 < v1434 {
		v402 = v1433
		goto L61
	} else {
		goto L262
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1428 = F_ReplicationSlotValidateName(m, v488, int32(21))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L261
	}
L65:
	;
	if v455-v456 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L66:
	;
	goto L65
L67:
	;
	v440 = v420
	v441 = v431
	goto L68
L68:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441)+1)))
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440)+1)))
	if v445 == int32(0) {
		v455 = v445
		v456 = v444
		goto L66
	} else {
		goto L70
	}
L69:
	;
	v455 = v445
	v456 = v444
	goto L66
L70:
	;
	v448 = int32(1)
	if v445 == v444 {
		v440 = v440 + v448
		v441 = v441 + v448
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v460&int32(8) != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v540 = int32(_a_F_AlterSubscription_10)
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v546 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[6])))
	if base.B2i32(v543 == int32(0))|base.B2i32(v543 != v546) != 0 {
		v564 = v543
		v565 = v546
		goto L89
	} else {
		goto L90
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errorConflictingDefElem(m, v419, v41)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v460 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v488 = F_defGetString(m, v419)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L79
	}
L78:
	;
	goto L3
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v501 = int32(_a_F_AlterSubscription_11)
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[7])))
	if base.B2i32(v504 == int32(0))|base.B2i32(v504 != v507) != 0 {
		v525 = v504
		v526 = v507
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v525-v526 != 0 {
		goto L64
	} else {
		goto L87
	}
L81:
	;
	goto L80
L82:
	;
	v510 = v488
	v511 = v501
	goto L83
L83:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511)+1)))
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+1)))
	if v515 == int32(0) {
		v525 = v515
		v526 = v514
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v525 = v515
	v526 = v514
	goto L81
L85:
	;
	v518 = int32(1)
	if v515 == v514 {
		v510 = v510 + v518
		v511 = v511 + v518
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = int32(0)
	goto L63
L88:
	;
	if v564-v565 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	goto L88
L90:
	;
	v549 = v420
	v550 = v540
	goto L91
L91:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+1)))
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549)+1)))
	if v554 == int32(0) {
		v564 = v554
		v565 = v553
		goto L89
	} else {
		goto L93
	}
L92:
	;
	v564 = v554
	v565 = v553
	goto L89
L93:
	;
	v557 = int32(1)
	if v554 == v553 {
		v549 = v549 + v557
		v550 = v550 + v557
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v569&int32(32) != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v627 = int32(_a_F_AlterSubscription_12)
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v633 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[8])))
	if base.B2i32(v630 == int32(0))|base.B2i32(v630 != v633) != 0 {
		v651 = v630
		v652 = v633
		goto L105
	} else {
		goto L106
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errorConflictingDefElem(m, v419, v41)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v569 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v597 = F_defGetString(m, v419)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L102
	}
L101:
	;
	goto L3
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v613 = int32(0)
	F_set_config_option(m, int32(_a_F_AlterSubscription_10), v597, int32(4), int32(12), v613, v613)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L103
	}
L103:
	;
	goto L63
L104:
	;
	if v651-v652 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L105:
	;
	goto L104
L106:
	;
	v636 = v420
	v637 = v627
	goto L107
L107:
	;
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637)+1)))
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636)+1)))
	if v641 == int32(0) {
		v651 = v641
		v652 = v640
		goto L105
	} else {
		goto L109
	}
L108:
	;
	v651 = v641
	v652 = v640
	goto L105
L109:
	;
	v644 = int32(1)
	if v641 == v640 {
		v636 = v636 + v644
		v637 = v637 + v644
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v656&int32(128) != 0 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v697 = int32(_a_F_AlterSubscription_13)
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v703 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[9])))
	if base.B2i32(v700 == int32(0))|base.B2i32(v700 != v703) != 0 {
		v721 = v700
		v722 = v703
		goto L120
	} else {
		goto L121
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errorConflictingDefElem(m, v419, v41)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v656 | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v684 = F_defGetBoolean(m, v419)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L118
	}
L117:
	;
	goto L3
L118:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+745)) = uint8(v684)
	goto L63
L119:
	;
	if v721-v722 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L120:
	;
	goto L119
L121:
	;
	v706 = v420
	v707 = v697
	goto L122
L122:
	;
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707)+1)))
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+1)))
	if v711 == int32(0) {
		v721 = v711
		v722 = v710
		goto L120
	} else {
		goto L124
	}
L123:
	;
	v721 = v711
	v722 = v710
	goto L120
L124:
	;
	v714 = int32(1)
	if v711 == v710 {
		v706 = v706 + v714
		v707 = v707 + v714
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v726&int32(256) != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v767 = int32(_a_F_AlterSubscription_14)
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v773 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[10])))
	if base.B2i32(v770 == int32(0))|base.B2i32(v770 != v773) != 0 {
		v791 = v770
		v792 = v773
		goto L135
	} else {
		goto L136
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errorConflictingDefElem(m, v419, v41)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v726 | int32(256)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v754 = F_defGetStreamingMode(m, v419)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L133
	}
L132:
	;
	goto L3
L133:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+746)) = uint8(v754)
	goto L63
L134:
	;
	if v791-v792 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L135:
	;
	goto L134
L136:
	;
	v776 = v420
	v777 = v767
	goto L137
L137:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777)+1)))
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776)+1)))
	if v781 == int32(0) {
		v791 = v781
		v792 = v780
		goto L135
	} else {
		goto L139
	}
L138:
	;
	v791 = v781
	v792 = v780
	goto L135
L139:
	;
	v784 = int32(1)
	if v781 == v780 {
		v776 = v776 + v784
		v777 = v777 + v784
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v796&int32(512) != 0 {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v837 = int32(_a_F_AlterSubscription_15)
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v843 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[11])))
	if base.B2i32(v840 == int32(0))|base.B2i32(v840 != v843) != 0 {
		v861 = v840
		v862 = v843
		goto L150
	} else {
		goto L151
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errorConflictingDefElem(m, v419, v41)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v796 | int32(512)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v824 = F_defGetBoolean(m, v419)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L148
	}
L147:
	;
	goto L3
L148:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+747)) = uint8(v824)
	goto L63
L149:
	;
	if v861-v862 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L150:
	;
	goto L149
L151:
	;
	v846 = v420
	v847 = v837
	goto L152
L152:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847)+1)))
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846)+1)))
	if v851 == int32(0) {
		v861 = v851
		v862 = v850
		goto L150
	} else {
		goto L154
	}
L153:
	;
	v861 = v851
	v862 = v850
	goto L150
L154:
	;
	v854 = int32(1)
	if v851 == v850 {
		v846 = v846 + v854
		v847 = v847 + v854
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v866&int32(1024) != 0 {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v907 = int32(_a_F_AlterSubscription_16)
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v913 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[12])))
	if base.B2i32(v910 == int32(0))|base.B2i32(v910 != v913) != 0 {
		v931 = v910
		v932 = v913
		goto L165
	} else {
		goto L166
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errorConflictingDefElem(m, v419, v41)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v866 | int32(1024)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v894 = F_defGetBoolean(m, v419)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L163
	}
L162:
	;
	goto L3
L163:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+748)) = uint8(v894)
	goto L63
L164:
	;
	if v931-v932 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L165:
	;
	goto L164
L166:
	;
	v916 = v420
	v917 = v907
	goto L167
L167:
	;
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v917)+1)))
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v916)+1)))
	if v921 == int32(0) {
		v931 = v921
		v932 = v920
		goto L165
	} else {
		goto L169
	}
L168:
	;
	v931 = v921
	v932 = v920
	goto L165
L169:
	;
	v924 = int32(1)
	if v921 == v920 {
		v916 = v916 + v924
		v917 = v917 + v924
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v936&int32(2048) != 0 {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v977 = int32(_a_F_AlterSubscription_17)
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v983 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[13])))
	if base.B2i32(v980 == int32(0))|base.B2i32(v980 != v983) != 0 {
		v1001 = v980
		v1002 = v983
		goto L180
	} else {
		goto L181
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errorConflictingDefElem(m, v419, v41)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v936 | int32(2048)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v964 = F_defGetBoolean(m, v419)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L178
	}
L177:
	;
	goto L3
L178:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+749)) = uint8(v964)
	goto L63
L179:
	;
	if v1001-v1002 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L180:
	;
	goto L179
L181:
	;
	v986 = v420
	v987 = v977
	goto L182
L182:
	;
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987)+1)))
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986)+1)))
	if v991 == int32(0) {
		v1001 = v991
		v1002 = v990
		goto L180
	} else {
		goto L184
	}
L183:
	;
	v1001 = v991
	v1002 = v990
	goto L180
L184:
	;
	v994 = int32(1)
	if v991 == v990 {
		v986 = v986 + v994
		v987 = v987 + v994
		goto L182
	} else {
		goto L185
	}
L185:
	;
	goto L183
L186:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v1006&int32(_a_F_AlterSubscription_18) != 0 {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	goto L188
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1047 = int32(_a_F_AlterSubscription_19)
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v1053 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[14])))
	if base.B2i32(v1050 == int32(0))|base.B2i32(v1050 != v1053) != 0 {
		v1071 = v1050
		v1072 = v1053
		goto L195
	} else {
		goto L196
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errorConflictingDefElem(m, v419, v41)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1006 | int32(_a_F_AlterSubscription_18)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1034 = F_defGetBoolean(m, v419)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L193
	}
L192:
	;
	goto L3
L193:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+750)) = uint8(v1034)
	goto L63
L194:
	;
	if v1071-v1072 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L195:
	;
	goto L194
L196:
	;
	v1056 = v420
	v1057 = v1047
	goto L197
L197:
	;
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1057)+1)))
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056)+1)))
	if v1061 == int32(0) {
		v1071 = v1061
		v1072 = v1060
		goto L195
	} else {
		goto L199
	}
L198:
	;
	v1071 = v1061
	v1072 = v1060
	goto L195
L199:
	;
	v1064 = int32(1)
	if v1061 == v1060 {
		v1056 = v1056 + v1064
		v1057 = v1057 + v1064
		goto L197
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v1076&int32(_a_F_AlterSubscription_6) != 0 {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	goto L203
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1117 = int32(_a_F_AlterSubscription_20)
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[15])))
	if base.B2i32(v1120 == int32(0))|base.B2i32(v1120 != v1123) != 0 {
		v1141 = v1120
		v1142 = v1123
		goto L210
	} else {
		goto L211
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errorConflictingDefElem(m, v419, v41)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1076 | int32(_a_F_AlterSubscription_6)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1104 = F_defGetBoolean(m, v419)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L208
	}
L207:
	;
	goto L3
L208:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+751)) = uint8(v1104)
	goto L63
L209:
	;
	if v1141-v1142 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L210:
	;
	goto L209
L211:
	;
	v1126 = v420
	v1127 = v1117
	goto L212
L212:
	;
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1127)+1)))
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1126)+1)))
	if v1131 == int32(0) {
		v1141 = v1131
		v1142 = v1130
		goto L210
	} else {
		goto L214
	}
L213:
	;
	v1141 = v1131
	v1142 = v1130
	goto L210
L214:
	;
	v1134 = int32(1)
	if v1131 == v1130 {
		v1126 = v1126 + v1134
		v1127 = v1127 + v1134
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v1146&int32(_a_F_AlterSubscription_21) != 0 {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	goto L218
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L257
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errorConflictingDefElem(m, v419, v41)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1146 | int32(_a_F_AlterSubscription_21)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v44)+752))
	F_pfree(m, v1174)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L223
	}
L222:
	;
	goto L3
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1187 = F_defGetString(m, v419)
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+752)) = v1187
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1203 = v1187
	v1204 = int32(_a_F_AlterSubscription_11)
	goto L226
L225:
	;
	if v1241 == int32(0) {
		goto L63
	} else {
		goto L238
	}
L226:
	;
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203))))
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1204))))
	if v1207 == v1208 {
		v1230 = v1207
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v1241 = int32(0)
	goto L225
L228:
	;
	v1232 = int32(1)
	if v1230 != 0 {
		v1203 = v1203 + v1232
		v1204 = v1204 + v1232
		goto L226
	} else {
		goto L237
	}
L229:
	;
	if base.Ui32((v1207-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1218 = v1207 | int32(32)
	goto L232
L231:
	;
	v1218 = v1207
	goto L232
L232:
	;
	if base.Ui32((v1208-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1227 = v1208 | int32(32)
	goto L235
L234:
	;
	v1227 = v1208
	goto L235
L235:
	;
	if v1218 == v1227 {
		v1230 = v1218
		goto L228
	} else {
		goto L236
	}
L236:
	;
	v1241 = v1218 - v1227
	goto L225
L237:
	;
	goto L227
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v44)+752))
	v1258 = v1254
	v1259 = int32(_a_F_AlterSubscription_8)
	goto L240
L239:
	;
	if v1296 == int32(0) {
		goto L63
	} else {
		goto L252
	}
L240:
	;
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1258))))
	v1263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1259))))
	if v1262 == v1263 {
		v1285 = v1262
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v1296 = int32(0)
	goto L239
L242:
	;
	v1287 = int32(1)
	if v1285 != 0 {
		v1258 = v1258 + v1287
		v1259 = v1259 + v1287
		goto L240
	} else {
		goto L251
	}
L243:
	;
	if base.Ui32((v1262-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1273 = v1262 | int32(32)
	goto L246
L245:
	;
	v1273 = v1262
	goto L246
L246:
	;
	if base.Ui32((v1263-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1282 = v1263 | int32(32)
	goto L249
L248:
	;
	v1282 = v1263
	goto L249
L249:
	;
	if v1273 == v1282 {
		v1285 = v1273
		goto L242
	} else {
		goto L250
	}
L250:
	;
	v1296 = v1273 - v1282
	goto L239
L251:
	;
	goto L241
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v44)+752))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+128)) = v1336
	F_errmsg(m, int32(_a_F_AlterSubscription_22), v44+int32(128))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(331), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L256
	}
L256:
	;
	goto L3
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L258
	}
L258:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+144)) = v1385
	F_errmsg(m, int32(_a_F_AlterSubscription_24), v44+int32(144))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(363), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L260
	}
L260:
	;
	goto L3
L261:
	;
	goto L63
L262:
	;
	goto L62
L263:
	;
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v44)+736))
	if v1757 != 0 {
		goto L301
	} else {
		goto L302
	}
L264:
	;
	v1755 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+862)) = uint8(v1755)
	goto L263
L265:
	;
	v1753 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+894)) = uint8(v1753)
	goto L264
L266:
	;
	if v1468 == int32(0) {
		goto L263
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	if v1468 == int32(0) {
		goto L263
	} else {
		goto L299
	}
L269:
	;
	v1474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+741)))
	if v1474 == int32(1) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v1574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+742)))
	if v1574 == int32(1) {
		goto L282
	} else {
		goto L283
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L274
	}
L274:
	;
	if v1466&int32(2) != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+68)) = int32(_a_F_AlterSubscription_25)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+64)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_27), v44-int32(-64))
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+52)) = int32(_a_F_AlterSubscription_28)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+48)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_29), v44+int32(48))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L280
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(415), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L279
	}
L279:
	;
	goto L3
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(421), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L281
	}
L281:
	;
	goto L3
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v1674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+25)))
	if v1674 != int32(1) {
		goto L265
	} else {
		goto L294
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L286
	}
L286:
	;
	if v1466&int32(4) != 0 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+100)) = int32(_a_F_AlterSubscription_30)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+96)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_27), v44+int32(96))
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+84)) = int32(_a_F_AlterSubscription_31)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_29), v44+int32(80))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L292
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(431), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L291
	}
L291:
	;
	goto L3
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(437), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L293
	}
L293:
	;
	goto L3
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(325))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+112)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_32), v44+int32(112))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L297
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1186), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L298
	}
L298:
	;
	goto L3
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1750 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v1469)
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L300
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+824)) = v1750
	goto L264
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1768 = F_cstring_to_text(m, v1757)
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v1774&int32(128) != 0 {
		goto L305
	} else {
		goto L306
	}
L304:
	;
	v1770 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+863)) = uint8(v1770)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+828)) = v1768
	goto L303
L305:
	;
	v1777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+745)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+792)) = v1777
	v1779 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+854)) = uint8(v1779)
	goto L307
L306:
	;
	goto L307
L307:
	;
	if v1774&int32(256) != 0 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1783 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+746)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+796)) = v1783
	v1785 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+855)) = uint8(v1785)
	goto L310
L309:
	;
	goto L310
L310:
	;
	if v1774&int32(1024) != 0 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+748)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+804)) = v1789
	v1791 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+857)) = uint8(v1791)
	goto L313
L312:
	;
	goto L313
L313:
	;
	if v1774&int32(2048) != 0 {
		goto L319
	} else {
		goto L320
	}
L314:
	;
	v2084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+28)))
	if v2084 != int32(101) {
		goto L40
	} else {
		goto L352
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L347
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v2008 = int32(1)
	v2010 = F_logicalrep_workers_find(m, v178, v2008, v2008)
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L345
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1994 = int32(1)
	v1996 = F_logicalrep_workers_find(m, v178, v1994, v1994)
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L343
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L338
	}
L319:
	;
	v1796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+749)))
	if v1796 == int32(0) {
		goto L322
	} else {
		goto L323
	}
L320:
	;
	v1821 = v1774
	goto L321
L321:
	;
	if v1821&int32(_a_F_AlterSubscription_18) != 0 {
		goto L327
	} else {
		goto L328
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1809 = F_superuser(m)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L325
	}
L323:
	;
	v1815 = int32(1)
	v1816 = v1774
	goto L324
L324:
	;
	v1817 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+858)) = uint8(v1817)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+808)) = v1815
	v1821 = v1816
	goto L321
L325:
	;
	if v1809 == int32(0) {
		goto L318
	} else {
		goto L326
	}
L326:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	v1814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+749)))
	v1815 = v1814
	v1816 = v1813
	goto L324
L327:
	;
	v1824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+750)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+812)) = v1824
	v1826 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+859)) = uint8(v1826)
	goto L329
L328:
	;
	goto L329
L329:
	;
	v1828 = int32(0)
	if v1821&int32(512) == v1828 {
		v5334 = v1821
		v5335 = v1828
		goto L39
	} else {
		goto L330
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v1844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+747)))
	v1845 = int32(1)
	v1846 = v1844 ^ v1845
	F_CheckAlterSubOption(m, v233, int32(_a_F_AlterSubscription_14), v1846&v1845, v43)
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L331
	}
L331:
	;
	if v1844&int32(1) != 0 {
		goto L317
	} else {
		goto L332
	}
L332:
	;
	v1853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+728)))
	if v1853&int32(8) == int32(0) {
		goto L316
	} else {
		goto L333
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errmsg(m, int32(_a_F_AlterSubscription_33), int32(0))
	mBase = m.M
	v1898 = m.ExcPending
	if v1898 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1270), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L337
	}
L337:
	;
	goto L3
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16797828))
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L339
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errmsg(m, int32(_a_F_AlterSubscription_4), int32(0))
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L340
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errhint(m, int32(_a_F_AlterSubscription_5), int32(0))
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1232), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L342
	}
L342:
	;
	goto L3
L343:
	;
	if v1996 != 0 {
		goto L315
	} else {
		goto L344
	}
L344:
	;
	goto L40
L345:
	;
	if v2010 == int32(0) {
		goto L314
	} else {
		goto L346
	}
L346:
	;
	goto L315
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(325))
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L348
	}
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errmsg(m, int32(_a_F_AlterSubscription_34), int32(0))
	mBase = m.M
	v2054 = m.ExcPending
	if v2054 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L349
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errhint(m, int32(_a_F_AlterSubscription_35), int32(0))
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1287), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L351
	}
L351:
	;
	goto L3
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v2097 = int32(0)
	v2099 = m.G0
	v2101 = v2099 - int32(240)
	m.G0 = v2101
	v2104 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[16]))
	v2108 = F_LWLockAcquire(m, v2104+int32(2304), int32(1))
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L354
	}
L353:
	;
	if v2234 == int32(0) {
		goto L40
	} else {
		goto L381
	}
L354:
	;
	v2111 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[17]))
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+4))
	if v2112 <= int32(0) {
		v2234 = v2097
		goto L356
	} else {
		goto L357
	}
L355:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L377
	}
L356:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[16]))
	F_LWLockRelease(m, v2247+int32(2304))
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L376
	}
L357:
	;
	v2131 = v2111
	v2134 = v2097
	goto L358
L358:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v2131+v2134<<(uint(int32(2))%32))+8))
	v2149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2148)+44)))
	if v2149 != int32(1) {
		goto L360
	} else {
		goto L361
	}
L359:
	;
	v2234 = int32(0)
	goto L356
L360:
	;
	v2211 = v2134 + int32(1)
	v2213 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[17]))
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(v2213)+4))
	if v2211 < v2214 {
		v2131 = v2213
		v2134 = v2211
		goto L358
	} else {
		goto L375
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+16)) = v2101 + int32(236)
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+20)) = v2101 + int32(232)
	v2159 = v2148 + int32(47)
	v2163 = F_sscanf(m, v2159, int32(_a_F_AlterSubscription_36), v2101+int32(16))
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L362
	}
L362:
	;
	if v2163 != int32(2) {
		goto L360
	} else {
		goto L363
	}
L363:
	;
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+236))
	if v178 != v2167 {
		goto L360
	} else {
		goto L364
	}
L364:
	;
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+232))
	if v2169 == int32(0) {
		goto L355
	} else {
		goto L365
	}
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+4)) = v2169
	*(*int32)(unsafe.Add(mBase, uint32(v2101))) = v178
	v2175 = v2101 + int32(32)
	v2178 = F_pg_snprintf(m, v2175, int32(200), int32(_a_F_AlterSubscription_36), v2101)
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L366
	}
L366:
	;
	v2182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2159))))
	v2185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2175))))
	if base.B2i32(v2182 == int32(0))|base.B2i32(v2182 != v2185) != 0 {
		v2203 = v2182
		v2204 = v2185
		goto L368
	} else {
		goto L369
	}
L367:
	;
	if v2203-v2204 != 0 {
		goto L360
	} else {
		goto L374
	}
L368:
	;
	goto L367
L369:
	;
	v2188 = v2159
	v2189 = v2175
	goto L370
L370:
	;
	v2192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2189)+1)))
	v2193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2188)+1)))
	if v2193 == int32(0) {
		v2203 = v2193
		v2204 = v2192
		goto L368
	} else {
		goto L372
	}
L371:
	;
	v2203 = v2193
	v2204 = v2192
	goto L368
L372:
	;
	v2196 = int32(1)
	if v2193 == v2192 {
		v2188 = v2188 + v2196
		v2189 = v2189 + v2196
		goto L370
	} else {
		goto L373
	}
L373:
	;
	goto L371
L374:
	;
	v2234 = int32(1)
	goto L356
L375:
	;
	goto L359
L376:
	;
	m.G0 = v2101 + int32(240)
	goto L353
L377:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L378
	}
L378:
	;
	F_errmsg_internal(m, int32(_a_F_AlterSubscription_37), int32(0))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L379
	}
L379:
	;
	F_errfinish(m, int32(_a_F_AlterSubscription_38), int32(2689), int32(_a_F_AlterSubscription_39))
	mBase = m.M
	v2270 = m.ExcPending
	if v2270 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L380
	}
L380:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L382
	}
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(325))
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L383
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errmsg(m, int32(_a_F_AlterSubscription_40), int32(0))
	mBase = m.M
	v2313 = m.ExcPending
	if v2313 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L384
	}
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errhint(m, int32(_a_F_AlterSubscription_41), int32(0))
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L385
	}
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1301), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v2342 = m.ExcPending
	if v2342 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L386
	}
L386:
	;
	goto L3
L387:
	;
	v2359 = int32(0)
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(v2343)+4))
	if v2359 < v2360 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v2378 = v2359
	goto L391
L389:
	;
	goto L390
L390:
	;
	v2530 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v2530&int32(8) != 0 {
		v4951 = v2530
		goto L43
	} else {
		goto L413
	}
L391:
	;
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(v2343)+12))
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v2393+v2378<<(uint(int32(2))%32))))
	v2398 = *(*int32)(unsafe.Add(mBase, uint32(v2397)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v2409 = int32(_a_F_AlterSubscription_42)
	v2412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2398))))
	v2415 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[18])))
	if base.B2i32(v2412 == int32(0))|base.B2i32(v2412 != v2415) != 0 {
		v2433 = v2412
		v2434 = v2415
		goto L394
	} else {
		goto L395
	}
L393:
	;
	if v2433-v2434 != 0 {
		goto L400
	} else {
		goto L401
	}
L394:
	;
	goto L393
L395:
	;
	v2418 = v2398
	v2419 = v2409
	goto L396
L396:
	;
	v2422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2419)+1)))
	v2423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2418)+1)))
	if v2423 == int32(0) {
		v2433 = v2423
		v2434 = v2422
		goto L394
	} else {
		goto L398
	}
L397:
	;
	v2433 = v2423
	v2434 = v2422
	goto L394
L398:
	;
	v2426 = int32(1)
	if v2423 == v2422 {
		v2418 = v2418 + v2426
		v2419 = v2419 + v2426
		goto L396
	} else {
		goto L399
	}
L399:
	;
	goto L397
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2449 = m.ExcPending
	if v2449 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L403
	}
L401:
	;
	goto L402
L402:
	;
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v2495&int32(2) != 0 {
		goto L407
	} else {
		goto L408
	}
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2462 = m.ExcPending
	if v2462 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L404
	}
L404:
	;
	v2463 = *(*int32)(unsafe.Add(mBase, uint32(v2397)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+224)) = v2463
	F_errmsg(m, int32(_a_F_AlterSubscription_24), v44+int32(224))
	mBase = m.M
	v2479 = m.ExcPending
	if v2479 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(363), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v2494 = m.ExcPending
	if v2494 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L406
	}
L406:
	;
	goto L3
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errorConflictingDefElem(m, v2397, v41)
	mBase = m.M
	v2509 = m.ExcPending
	if v2509 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L410
	}
L408:
	;
	goto L409
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v2495 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v2523 = F_defGetBoolean(m, v2397)
	mBase = m.M
	v2524 = m.ExcPending
	if v2524 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L411
	}
L410:
	;
	goto L3
L411:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+741)) = uint8(v2523)
	v2527 = v2378 + int32(1)
	v2528 = *(*int32)(unsafe.Add(mBase, uint32(v2343)+4))
	if v2527 < v2528 {
		v2378 = v2527
		goto L391
	} else {
		goto L412
	}
L412:
	;
	goto L44
L413:
	;
	v5179 = v2354
	goto L41
L414:
	;
	v2549 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[2]))
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(v2549)+4))
	v2551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+30)))
	if v2551 == int32(1) {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v2554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+24)))
	v2557 = v2554 ^ int32(1)
	goto L417
L416:
	;
	v2557 = int32(0)
	goto L417
L417:
	;
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	m.T0[v2550].(func(*base.Module, int32, int32))(m, v2558, v2557&int32(1))
	mBase = m.M
	v2572 = m.ExcPending
	if v2572 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L418
	}
L418:
	;
	v2573 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v2584 = F_cstring_to_text(m, v2573)
	mBase = m.M
	v2585 = m.ExcPending
	if v2585 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L419
	}
L419:
	;
	v2586 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+861)) = uint8(v2586)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+820)) = v2584
	v4909 = v53
	v4910 = v54
	goto L45
L420:
	;
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v2615 = F_publicationListToArray(m, v2604)
	mBase = m.M
	v2616 = m.ExcPending
	if v2616 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L421
	}
L421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+832)) = v2615
	v2618 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+864)) = uint8(v2618)
	v2620 = int32(0)
	v2621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+744)))
	if v2621 != v2618 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v5405 = v53
	v5406 = v54
	v5407 = int32(0)
	v5409 = v2620
	goto L38
L423:
	;
	goto L424
L424:
	;
	v2625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+25)))
	if v2625 == int32(0) {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2641 = m.ExcPending
	if v2641 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L428
	}
L426:
	;
	goto L427
L427:
	;
	v2698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+28)))
	if v2698 != int32(101) {
		goto L433
	} else {
		goto L434
	}
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(325))
	mBase = m.M
	v2654 = m.ExcPending
	if v2654 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errmsg(m, int32(_a_F_AlterSubscription_43), int32(0))
	mBase = m.M
	v2668 = m.ExcPending
	if v2668 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L430
	}
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errhint(m, int32(_a_F_AlterSubscription_44), int32(0))
	mBase = m.M
	v2682 = m.ExcPending
	if v2682 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L431
	}
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1393), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v2697 = m.ExcPending
	if v2697 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L432
	}
L432:
	;
	goto L3
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_PreventInTransactionBlock(m, v43, int32(_a_F_AlterSubscription_45))
	mBase = m.M
	v2788 = m.ExcPending
	if v2788 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L441
	}
L434:
	;
	v2701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+743)))
	if v2701&int32(1) == int32(0) {
		goto L433
	} else {
		goto L435
	}
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2719 = m.ExcPending
	if v2719 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L436
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(325))
	mBase = m.M
	v2732 = m.ExcPending
	if v2732 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L437
	}
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errmsg(m, int32(_a_F_AlterSubscription_46), int32(0))
	mBase = m.M
	v2746 = m.ExcPending
	if v2746 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L438
	}
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errhint(m, int32(_a_F_AlterSubscription_47), int32(0))
	mBase = m.M
	v2760 = m.ExcPending
	if v2760 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L439
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1403), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v2775 = m.ExcPending
	if v2775 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L440
	}
L440:
	;
	goto L3
L441:
	;
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v233)+48)) = v2789
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v2801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+743)))
	F_AlterSubscription_refresh(m, v233, v2801, v2789)
	mBase = m.M
	v2803 = m.ExcPending
	if v2803 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L442
	}
L442:
	;
	v5405 = v53
	v5406 = v54
	v5407 = int32(0)
	v5409 = v2620
	goto L38
L443:
	;
	v2820 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v2821 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v2822 = *(*int32)(unsafe.Add(mBase, uint32(v233)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v2833 = F_list_copy(m, v2822)
	mBase = m.M
	v2834 = m.ExcPending
	if v2834 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L444
	}
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_check_duplicates_in_publist(m, v2821, int32(0))
	mBase = m.M
	v2847 = m.ExcPending
	if v2847 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L445
	}
L445:
	;
	if v2821 == int32(0) {
		v3209 = v53
		v3210 = v54
		v3213 = v2833
		goto L446
	} else {
		goto L447
	}
L446:
	;
	if v3213 == int32(0) {
		goto L485
	} else {
		goto L486
	}
L447:
	;
	v2850 = int32(0)
	v2851 = *(*int32)(unsafe.Add(mBase, uint32(v2821)+4))
	if v2851 <= v2850 {
		v3209 = v53
		v3210 = v54
		v3213 = v2833
		goto L446
	} else {
		goto L448
	}
L448:
	;
	v2867 = v53
	v2868 = v54
	v2871 = v2833
	v2874 = v2850
	goto L449
L449:
	;
	v2884 = *(*int32)(unsafe.Add(mBase, uint32(v2821)+12))
	v2888 = *(*int32)(unsafe.Add(mBase, uint32(v2884+v2874<<(uint(int32(2))%32))))
	v2889 = *(*int32)(unsafe.Add(mBase, uint32(v2888)+4))
	if v2871 == int32(0) {
		goto L453
	} else {
		goto L454
	}
L450:
	;
	v3209 = v3174
	v3210 = v3175
	v3213 = v3191
	goto L446
L451:
	;
	v3193 = v2874 + int32(1)
	v3194 = *(*int32)(unsafe.Add(mBase, uint32(v2821)+4))
	if v3193 < v3194 {
		v2867 = v3174
		v2868 = v3175
		v2871 = v3191
		v2874 = v3193
		goto L449
	} else {
		goto L484
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v2867
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v2868
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v3159 = F_list_delete_nth_cell(m, v2871, v2912)
	mBase = m.M
	v3160 = m.ExcPending
	if v3160 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L483
	}
L453:
	;
	if v349 == int32(3) {
		goto L474
	} else {
		goto L475
	}
L454:
	;
	v2892 = *(*int32)(unsafe.Add(mBase, uint32(v2871)+4))
	if v2892 <= int32(0) {
		goto L453
	} else {
		goto L455
	}
L455:
	;
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v2871)+12))
	v2912 = int32(0)
	goto L456
L456:
	;
	v2930 = *(*int32)(unsafe.Add(mBase, uint32(v2895+v2912<<(uint(int32(2))%32))))
	v2931 = *(*int32)(unsafe.Add(mBase, uint32(v2930)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v2867
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v2868
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v2944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2889))))
	v2947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2931))))
	if base.B2i32(v2944 == int32(0))|base.B2i32(v2944 != v2947) != 0 {
		v2965 = v2944
		v2966 = v2947
		goto L459
	} else {
		goto L460
	}
L457:
	;
	goto L453
L458:
	;
	if v2965-v2966 == int32(0) {
		goto L465
	} else {
		goto L466
	}
L459:
	;
	goto L458
L460:
	;
	v2950 = v2889
	v2951 = v2931
	goto L461
L461:
	;
	v2954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2951)+1)))
	v2955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2950)+1)))
	if v2955 == int32(0) {
		v2965 = v2955
		v2966 = v2954
		goto L459
	} else {
		goto L463
	}
L462:
	;
	v2965 = v2955
	v2966 = v2954
	goto L459
L463:
	;
	v2958 = int32(1)
	if v2955 == v2954 {
		v2950 = v2950 + v2958
		v2951 = v2951 + v2958
		goto L461
	} else {
		goto L464
	}
L464:
	;
	goto L462
L465:
	;
	if v349 != int32(3) {
		goto L452
	} else {
		goto L468
	}
L466:
	;
	goto L467
L467:
	;
	v3032 = v2912 + int32(1)
	if v2892 != v3032 {
		v2912 = v3032
		goto L456
	} else {
		goto L473
	}
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v2867
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v2868
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2985 = m.ExcPending
	if v2985 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L469
	}
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v2867
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v2868
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(_a_F_AlterSubscription_48))
	mBase = m.M
	v2998 = m.ExcPending
	if v2998 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L470
	}
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v2867
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v2868
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+292)) = v2820
	*(*int32)(unsafe.Add(mBase, uint32(v44)+288)) = v2889
	F_errmsg(m, int32(_a_F_AlterSubscription_49), v44+int32(288))
	mBase = m.M
	v3015 = m.ExcPending
	if v3015 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L471
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v2867
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v2868
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2427), int32(_a_F_AlterSubscription_50))
	mBase = m.M
	v3030 = m.ExcPending
	if v3030 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L472
	}
L472:
	;
	goto L3
L473:
	;
	goto L457
L474:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v2867
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v2868
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v3076 = F_makeString(m, v2889)
	mBase = m.M
	v3077 = m.ExcPending
	if v3077 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L477
	}
L475:
	;
	goto L476
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v2867
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v2868
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3103 = m.ExcPending
	if v3103 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L479
	}
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v2867
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v2868
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v3088 = F_lappend(m, v2871, v3076)
	mBase = m.M
	v3089 = m.ExcPending
	if v3089 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L478
	}
L478:
	;
	v3174 = v3088
	v3175 = v2868
	v3191 = v3088
	goto L451
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v2867
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v2868
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(117833860))
	mBase = m.M
	v3116 = m.ExcPending
	if v3116 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L480
	}
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v2867
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v2868
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+276)) = v2820
	*(*int32)(unsafe.Add(mBase, uint32(v44)+272)) = v2889
	F_errmsg(m, int32(_a_F_AlterSubscription_51), v44+int32(272))
	mBase = m.M
	v3133 = m.ExcPending
	if v3133 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L481
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v2867
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v2868
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2441), int32(_a_F_AlterSubscription_50))
	mBase = m.M
	v3148 = m.ExcPending
	if v3148 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L482
	}
L482:
	;
	goto L3
L483:
	;
	v3174 = v2867
	v3175 = v3159
	v3191 = v3159
	goto L451
L484:
	;
	goto L450
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3241 = m.ExcPending
	if v3241 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L488
	}
L486:
	;
	goto L487
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v3295 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[19]))
	v3300 = F_AllocSetContextCreateInternal(m, v3295, int32(_a_F_AlterSubscription_52), int32(0), int32(_a_F_AlterSubscription_6), int32(_a_F_AlterSubscription_53))
	mBase = m.M
	v3301 = m.ExcPending
	if v3301 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L492
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(117833860))
	mBase = m.M
	v3254 = m.ExcPending
	if v3254 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L489
	}
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errmsg(m, int32(_a_F_AlterSubscription_54), int32(0))
	mBase = m.M
	v3268 = m.ExcPending
	if v3268 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L490
	}
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2451), int32(_a_F_AlterSubscription_50))
	mBase = m.M
	v3283 = m.ExcPending
	if v3283 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L491
	}
L491:
	;
	goto L3
L492:
	;
	v3302 = int32(_a_F_AlterSubscription_55)
	v3303 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[19]))
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[19])) = v3300
	v3306 = *(*int32)(unsafe.Add(mBase, uint32(v3213)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v3319 = F_palloc(m, v3306<<(uint(int32(2))%32))
	mBase = m.M
	v3320 = m.ExcPending
	if v3320 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L493
	}
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_check_duplicates_in_publist(m, v3213, v3319)
	mBase = m.M
	v3332 = m.ExcPending
	if v3332 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L494
	}
L494:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[19])) = v3303
	v3335 = *(*int32)(unsafe.Add(mBase, uint32(v3213)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v3347 = F_construct_array_builtin(m, v3319, v3335, int32(25))
	mBase = m.M
	v3348 = m.ExcPending
	if v3348 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L495
	}
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_MemoryContextDelete(m, v3300)
	mBase = m.M
	v3360 = m.ExcPending
	if v3360 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L496
	}
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+832)) = v3347
	v3362 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+864)) = uint8(v3362)
	v3364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+744)))
	if v3364 != v3362 {
		v4909 = v3209
		v4910 = v3210
		goto L45
	} else {
		goto L497
	}
L497:
	;
	if v349 == int32(3) {
		goto L498
	} else {
		goto L499
	}
L498:
	;
	v3370 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v3371 = v3370
	goto L500
L499:
	;
	v3371 = int32(0)
	goto L500
L500:
	;
	v3372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+25)))
	if v3372 == int32(0) {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3388 = m.ExcPending
	if v3388 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L504
	}
L502:
	;
	goto L503
L503:
	;
	v3452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+28)))
	if v3452 != int32(101) {
		goto L512
	} else {
		goto L513
	}
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(325))
	mBase = m.M
	v3401 = m.ExcPending
	if v3401 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L505
	}
L505:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errmsg(m, int32(_a_F_AlterSubscription_43), int32(0))
	mBase = m.M
	v3415 = m.ExcPending
	if v3415 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L506
	}
L506:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	if v349 == int32(3) {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v3430 = int32(_a_F_AlterSubscription_56)
	goto L509
L508:
	;
	v3430 = int32(_a_F_AlterSubscription_57)
	goto L509
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+256)) = v3430
	F_errhint(m, int32(_a_F_AlterSubscription_58), v44+int32(256))
	mBase = m.M
	v3436 = m.ExcPending
	if v3436 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L510
	}
L510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1448), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v3451 = m.ExcPending
	if v3451 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L511
	}
L511:
	;
	goto L3
L512:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_PreventInTransactionBlock(m, v43, int32(_a_F_AlterSubscription_45))
	mBase = m.M
	v3549 = m.ExcPending
	if v3549 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L523
	}
L513:
	;
	v3455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+743)))
	if v3455&int32(1) == int32(0) {
		goto L512
	} else {
		goto L514
	}
L514:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3473 = m.ExcPending
	if v3473 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L515
	}
L515:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(325))
	mBase = m.M
	v3486 = m.ExcPending
	if v3486 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L516
	}
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errmsg(m, int32(_a_F_AlterSubscription_46), int32(0))
	mBase = m.M
	v3500 = m.ExcPending
	if v3500 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L517
	}
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	if v349 == int32(3) {
		goto L518
	} else {
		goto L519
	}
L518:
	;
	v3515 = int32(_a_F_AlterSubscription_59)
	goto L520
L519:
	;
	v3515 = int32(_a_F_AlterSubscription_60)
	goto L520
L520:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+240)) = v3515
	F_errhint(m, int32(_a_F_AlterSubscription_61), v44+int32(240))
	mBase = m.M
	v3521 = m.ExcPending
	if v3521 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L521
	}
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1462), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v3536 = m.ExcPending
	if v3536 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L522
	}
L522:
	;
	goto L3
L523:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v233)+48)) = v3213
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v3561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+743)))
	F_AlterSubscription_refresh(m, v233, v3561, v3371)
	mBase = m.M
	v3563 = m.ExcPending
	if v3563 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L524
	}
L524:
	;
	v4909 = v3209
	v4910 = v3210
	goto L45
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3580 = m.ExcPending
	if v3580 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L528
	}
L526:
	;
	goto L527
L527:
	;
	v3623 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	v3624 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+736)) = v3624
	*(*int64)(unsafe.Add(mBase, uint32(v44)+760)) = v3624
	*(*int64)(unsafe.Add(mBase, uint32(v44)+752)) = v3624
	*(*int64)(unsafe.Add(mBase, uint32(v44)+744)) = v3624
	*(*int64)(unsafe.Add(mBase, uint32(v44)+728)) = v3624
	v3634 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+743)) = uint8(v3634)
	if v3623 == int32(0) {
		v4775 = v3634
		goto L46
	} else {
		goto L532
	}
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(325))
	mBase = m.M
	v3593 = m.ExcPending
	if v3593 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L529
	}
L529:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errmsg(m, int32(_a_F_AlterSubscription_62), int32(0))
	mBase = m.M
	v3607 = m.ExcPending
	if v3607 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L530
	}
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1481), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v3622 = m.ExcPending
	if v3622 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L531
	}
L531:
	;
	goto L3
L532:
	;
	v3639 = int32(0)
	v3640 = *(*int32)(unsafe.Add(mBase, uint32(v3623)+4))
	if v3639 < v3640 {
		goto L535
	} else {
		goto L536
	}
L533:
	;
	v3844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+741)))
	if v3844 == int32(1) {
		goto L560
	} else {
		goto L561
	}
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3798 = m.ExcPending
	if v3798 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L556
	}
L535:
	;
	v3658 = v3639
	goto L538
L536:
	;
	v3768 = v3634
	goto L537
L537:
	;
	v3782 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v3782&int32(8) != 0 {
		goto L533
	} else {
		goto L555
	}
L538:
	;
	v3673 = *(*int32)(unsafe.Add(mBase, uint32(v3623)+12))
	v3677 = *(*int32)(unsafe.Add(mBase, uint32(v3673+v3658<<(uint(int32(2))%32))))
	v3678 = *(*int32)(unsafe.Add(mBase, uint32(v3677)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v3689 = int32(_a_F_AlterSubscription_63)
	v3692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3678))))
	v3695 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[20])))
	if base.B2i32(v3692 == int32(0))|base.B2i32(v3692 != v3695) != 0 {
		v3713 = v3692
		v3714 = v3695
		goto L541
	} else {
		goto L542
	}
L539:
	;
	v3751 = *(*int32)(unsafe.Add(mBase, uint32(v44)+732))
	if v3751 != 0 {
		v4775 = v3744
		goto L46
	} else {
		goto L554
	}
L540:
	;
	if v3713-v3714 != 0 {
		goto L534
	} else {
		goto L547
	}
L541:
	;
	goto L540
L542:
	;
	v3698 = v3678
	v3699 = v3689
	goto L543
L543:
	;
	v3702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3699)+1)))
	v3703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3698)+1)))
	if v3703 == int32(0) {
		v3713 = v3703
		v3714 = v3702
		goto L541
	} else {
		goto L545
	}
L544:
	;
	v3713 = v3703
	v3714 = v3702
	goto L541
L545:
	;
	v3706 = int32(1)
	if v3703 == v3702 {
		v3698 = v3698 + v3706
		v3699 = v3699 + v3706
		goto L543
	} else {
		goto L546
	}
L546:
	;
	goto L544
L547:
	;
	v3716 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v3716&int32(16) != 0 {
		goto L548
	} else {
		goto L549
	}
L548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errorConflictingDefElem(m, v3677, v41)
	mBase = m.M
	v3730 = m.ExcPending
	if v3730 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L551
	}
L549:
	;
	goto L550
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v3716 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v3744 = F_defGetBoolean(m, v3677)
	mBase = m.M
	v3745 = m.ExcPending
	if v3745 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L552
	}
L551:
	;
	goto L3
L552:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+743)) = uint8(v3744)
	v3748 = v3658 + int32(1)
	v3749 = *(*int32)(unsafe.Add(mBase, uint32(v3623)+4))
	if v3748 < v3749 {
		v3658 = v3748
		goto L538
	} else {
		goto L553
	}
L553:
	;
	goto L539
L554:
	;
	v3768 = v3744
	goto L537
L555:
	;
	v4775 = v3768
	goto L46
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3811 = m.ExcPending
	if v3811 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L557
	}
L557:
	;
	v3812 = *(*int32)(unsafe.Add(mBase, uint32(v3677)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+368)) = v3812
	F_errmsg(m, int32(_a_F_AlterSubscription_24), v44+int32(368))
	mBase = m.M
	v3828 = m.ExcPending
	if v3828 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L558
	}
L558:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(363), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v3843 = m.ExcPending
	if v3843 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L559
	}
L559:
	;
	goto L3
L560:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3860 = m.ExcPending
	if v3860 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L563
	}
L561:
	;
	goto L562
L562:
	;
	v3944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+742)))
	if v3944 != int32(1) {
		v4775 = v3768
		goto L46
	} else {
		goto L572
	}
L563:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3873 = m.ExcPending
	if v3873 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L564
	}
L564:
	;
	if v3782&int32(2) != 0 {
		goto L565
	} else {
		goto L566
	}
L565:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+324)) = int32(_a_F_AlterSubscription_25)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+320)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_27), v44+int32(320))
	mBase = m.M
	v3894 = m.ExcPending
	if v3894 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L568
	}
L566:
	;
	goto L567
L567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+308)) = int32(_a_F_AlterSubscription_28)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+304)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_29), v44+int32(304))
	mBase = m.M
	v3928 = m.ExcPending
	if v3928 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L570
	}
L568:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(415), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v3909 = m.ExcPending
	if v3909 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L569
	}
L569:
	;
	goto L3
L570:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(421), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v3943 = m.ExcPending
	if v3943 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L571
	}
L571:
	;
	goto L3
L572:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3960 = m.ExcPending
	if v3960 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L573
	}
L573:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3973 = m.ExcPending
	if v3973 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L574
	}
L574:
	;
	if v3782&int32(4) != 0 {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+356)) = int32(_a_F_AlterSubscription_30)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+352)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_27), v44+int32(352))
	mBase = m.M
	v3994 = m.ExcPending
	if v3994 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L578
	}
L576:
	;
	goto L577
L577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+340)) = int32(_a_F_AlterSubscription_31)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+336)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_29), v44+int32(336))
	mBase = m.M
	v4028 = m.ExcPending
	if v4028 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L580
	}
L578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(431), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v4009 = m.ExcPending
	if v4009 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L579
	}
L579:
	;
	goto L3
L580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(437), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v4043 = m.ExcPending
	if v4043 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L581
	}
L581:
	;
	goto L3
L582:
	;
	v4058 = int32(0)
	v4059 = *(*int32)(unsafe.Add(mBase, uint32(v4044)+4))
	if v4059 <= v4058 {
		v4742 = v4045
		goto L47
	} else {
		goto L583
	}
L583:
	;
	v4077 = v4058
	goto L585
L584:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4301 = m.ExcPending
	if v4301 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L619
	}
L585:
	;
	v4092 = *(*int32)(unsafe.Add(mBase, uint32(v4044)+12))
	v4096 = *(*int32)(unsafe.Add(mBase, uint32(v4092+v4077<<(uint(int32(2))%32))))
	v4097 = *(*int32)(unsafe.Add(mBase, uint32(v4096)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v4108 = int32(_a_F_AlterSubscription_64)
	v4111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4097))))
	v4114 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[21])))
	if base.B2i32(v4111 == int32(0))|base.B2i32(v4111 != v4114) != 0 {
		v4132 = v4111
		v4133 = v4114
		goto L588
	} else {
		goto L589
	}
L586:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4243 = m.ExcPending
	if v4243 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L615
	}
L587:
	;
	if v4132-v4133 != 0 {
		goto L584
	} else {
		goto L594
	}
L588:
	;
	goto L587
L589:
	;
	v4117 = v4097
	v4118 = v4108
	goto L590
L590:
	;
	v4121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4118)+1)))
	v4122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4117)+1)))
	if v4122 == int32(0) {
		v4132 = v4122
		v4133 = v4121
		goto L588
	} else {
		goto L592
	}
L591:
	;
	v4132 = v4122
	v4133 = v4121
	goto L588
L592:
	;
	v4125 = int32(1)
	if v4122 == v4121 {
		v4117 = v4117 + v4125
		v4118 = v4118 + v4125
		goto L590
	} else {
		goto L593
	}
L593:
	;
	goto L591
L594:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v4145 = F_defGetString(m, v4096)
	mBase = m.M
	v4146 = m.ExcPending
	if v4146 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L595
	}
L595:
	;
	v4147 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v4147&int32(_a_F_AlterSubscription_65) != 0 {
		goto L596
	} else {
		goto L597
	}
L596:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errorConflictingDefElem(m, v4096, v41)
	mBase = m.M
	v4161 = m.ExcPending
	if v4161 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L599
	}
L597:
	;
	goto L598
L598:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v4172 = int32(_a_F_AlterSubscription_11)
	v4175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4145))))
	v4178 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[7])))
	if base.B2i32(v4175 == int32(0))|base.B2i32(v4175 != v4178) != 0 {
		v4196 = v4175
		v4197 = v4178
		goto L603
	} else {
		goto L604
	}
L599:
	;
	goto L3
L600:
	;
	goto L586
L601:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v44)+760)) = v4221
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v4220 | int32(_a_F_AlterSubscription_65)
	v4227 = v4077 + int32(1)
	v4228 = *(*int32)(unsafe.Add(mBase, uint32(v4044)+4))
	if v4228 <= v4227 {
		goto L48
	} else {
		goto L614
	}
L602:
	;
	if v4196-v4197 == int32(0) {
		goto L609
	} else {
		goto L610
	}
L603:
	;
	goto L602
L604:
	;
	v4181 = v4145
	v4182 = v4172
	goto L605
L605:
	;
	v4185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4182)+1)))
	v4186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4181)+1)))
	if v4186 == int32(0) {
		v4196 = v4186
		v4197 = v4185
		goto L603
	} else {
		goto L607
	}
L606:
	;
	v4196 = v4186
	v4197 = v4185
	goto L603
L607:
	;
	v4189 = int32(1)
	if v4186 == v4185 {
		v4181 = v4181 + v4189
		v4182 = v4182 + v4189
		goto L605
	} else {
		goto L608
	}
L608:
	;
	goto L606
L609:
	;
	v4220 = v4147
	v4221 = int64(0)
	goto L601
L610:
	;
	goto L611
L611:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v4214 = F_DirectFunctionCall1Coll(m, int32(572), int32(0), v4145)
	mBase = m.M
	v4215 = m.ExcPending
	if v4215 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L612
	}
L612:
	;
	v4216 = *(*int64)(unsafe.Add(mBase, uint32(v4214)))
	if v4216 == int64(0) {
		goto L600
	} else {
		goto L613
	}
L613:
	;
	v4219 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	v4220 = v4219
	v4221 = v4216
	goto L601
L614:
	;
	v4077 = v4227
	goto L585
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(50856066))
	mBase = m.M
	v4256 = m.ExcPending
	if v4256 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L616
	}
L616:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+464)) = v4145
	F_errmsg(m, int32(_a_F_AlterSubscription_66), v44+int32(464))
	mBase = m.M
	v4272 = m.ExcPending
	if v4272 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L617
	}
L617:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(354), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v4287 = m.ExcPending
	if v4287 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L618
	}
L618:
	;
	goto L3
L619:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4314 = m.ExcPending
	if v4314 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L620
	}
L620:
	;
	v4315 = *(*int32)(unsafe.Add(mBase, uint32(v4096)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+480)) = v4315
	F_errmsg(m, int32(_a_F_AlterSubscription_24), v44+int32(480))
	mBase = m.M
	v4331 = m.ExcPending
	if v4331 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L621
	}
L621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(363), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v4346 = m.ExcPending
	if v4346 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L622
	}
L622:
	;
	goto L3
L623:
	;
	v4361 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v4361
	F_errmsg_internal(m, int32(_a_F_AlterSubscription_67), v44+int32(16))
	mBase = m.M
	v4377 = m.ExcPending
	if v4377 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L624
	}
L624:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1556), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v4392 = m.ExcPending
	if v4392 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L625
	}
L625:
	;
	goto L3
L626:
	;
	if v4221 == int64(0) {
		goto L650
	} else {
		goto L651
	}
L627:
	;
	v4399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+741)))
	if v4399 == int32(1) {
		goto L628
	} else {
		goto L629
	}
L628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4415 = m.ExcPending
	if v4415 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L631
	}
L629:
	;
	goto L630
L630:
	;
	v4499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+742)))
	if v4499 != int32(1) {
		goto L626
	} else {
		goto L640
	}
L631:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4428 = m.ExcPending
	if v4428 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L632
	}
L632:
	;
	if v4220&int32(2) != 0 {
		goto L633
	} else {
		goto L634
	}
L633:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+420)) = int32(_a_F_AlterSubscription_25)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+416)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_27), v44+int32(416))
	mBase = m.M
	v4449 = m.ExcPending
	if v4449 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L636
	}
L634:
	;
	goto L635
L635:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+404)) = int32(_a_F_AlterSubscription_28)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+400)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_29), v44+int32(400))
	mBase = m.M
	v4483 = m.ExcPending
	if v4483 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L638
	}
L636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(415), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v4464 = m.ExcPending
	if v4464 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L637
	}
L637:
	;
	goto L3
L638:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(421), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v4498 = m.ExcPending
	if v4498 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L639
	}
L639:
	;
	goto L3
L640:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4515 = m.ExcPending
	if v4515 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L641
	}
L641:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4528 = m.ExcPending
	if v4528 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L642
	}
L642:
	;
	if v4220&int32(4) != 0 {
		goto L643
	} else {
		goto L644
	}
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+452)) = int32(_a_F_AlterSubscription_30)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+448)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_27), v44+int32(448))
	mBase = m.M
	v4549 = m.ExcPending
	if v4549 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L646
	}
L644:
	;
	goto L645
L645:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+436)) = int32(_a_F_AlterSubscription_31)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+432)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_29), v44+int32(432))
	mBase = m.M
	v4583 = m.ExcPending
	if v4583 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L648
	}
L646:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(431), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v4564 = m.ExcPending
	if v4564 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L647
	}
L647:
	;
	goto L3
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(437), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v4598 = m.ExcPending
	if v4598 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L649
	}
L649:
	;
	goto L3
L650:
	;
	v4742 = int64(0)
	goto L47
L651:
	;
	goto L652
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v4614 = v44 + int32(656)
	F_ReplicationOriginNameForLogicalRep(m, v178, int32(0), v4614)
	mBase = m.M
	v4616 = m.ExcPending
	if v4616 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L653
	}
L653:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v4628 = F_replorigin_by_name(m, v4614, int32(0))
	mBase = m.M
	v4629 = m.ExcPending
	if v4629 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L654
	}
L654:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v4641 = F_replorigin_get_progress(m, v4628, int32(0))
	mBase = m.M
	v4642 = m.ExcPending
	if v4642 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L655
	}
L655:
	;
	v4645 = *(*int64)(unsafe.Add(mBase, uint32(v44)+760))
	if base.B2i32(v4641 == int64(0))|base.B2i32(base.Ui64(v4641) <= base.Ui64(v4645)) != 0 {
		v4742 = v4645
		goto L47
	} else {
		goto L656
	}
L656:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4661 = m.ExcPending
	if v4661 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L657
	}
L657:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(50856066))
	mBase = m.M
	v4674 = m.ExcPending
	if v4674 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L658
	}
L658:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v4685 = *(*int64)(unsafe.Add(mBase, uint32(v44)+760))
	*(*uint32)(unsafe.Add(mBase, uint32(v44)+388)) = uint32(v4685)
	v4687 = int64(32)
	v4688 = int64(base.Ui64(v4685) >> (uint(v4687) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v44)+384)) = uint32(v4688)
	*(*uint32)(unsafe.Add(mBase, uint32(v44)+396)) = uint32(v4641)
	v4692 = int64(base.Ui64(v4641) >> (uint(v4687) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v44)+392)) = uint32(v4692)
	F_errmsg(m, int32(_a_F_AlterSubscription_68), v44+int32(384))
	mBase = m.M
	v4698 = m.ExcPending
	if v4698 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L659
	}
L659:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1544), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v4713 = m.ExcPending
	if v4713 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L660
	}
L660:
	;
	goto L3
L661:
	;
	v4756 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+850)) = uint8(v4756)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+776)) = v4754
	v4909 = v53
	v4910 = v54
	goto L45
L662:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4810 = m.ExcPending
	if v4810 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L665
	}
L663:
	;
	goto L664
L664:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_PreventInTransactionBlock(m, v43, int32(_a_F_AlterSubscription_69))
	mBase = m.M
	v4879 = m.ExcPending
	if v4879 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L670
	}
L665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4823 = m.ExcPending
	if v4823 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L666
	}
L666:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errmsg(m, int32(_a_F_AlterSubscription_70), int32(0))
	mBase = m.M
	v4837 = m.ExcPending
	if v4837 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L667
	}
L667:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errhint(m, int32(_a_F_AlterSubscription_71), int32(0))
	mBase = m.M
	v4851 = m.ExcPending
	if v4851 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L668
	}
L668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1507), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v4866 = m.ExcPending
	if v4866 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L669
	}
L669:
	;
	goto L3
L670:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v4890 = int32(0)
	v4891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+743)))
	F_AlterSubscription_refresh(m, v233, v4891, v4890)
	mBase = m.M
	v4894 = m.ExcPending
	if v4894 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L671
	}
L671:
	;
	v5480 = v53
	v5481 = v54
	v5482 = v4890
	v5484 = int32(0)
	goto L37
L672:
	;
	v4929 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	if v4929&int32(8) == int32(0) {
		v5179 = v2523
		goto L41
	} else {
		goto L673
	}
L673:
	;
	if v2523 == int32(0) {
		goto L42
	} else {
		goto L674
	}
L674:
	;
	v4951 = v4929
	goto L43
L675:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4992 = m.ExcPending
	if v4992 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L676
	}
L676:
	;
	if v4951&int32(2) != 0 {
		goto L677
	} else {
		goto L678
	}
L677:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+180)) = int32(_a_F_AlterSubscription_25)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+176)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_27), v44+int32(176))
	mBase = m.M
	v5013 = m.ExcPending
	if v5013 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L680
	}
L678:
	;
	goto L679
L679:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+164)) = int32(_a_F_AlterSubscription_28)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+160)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_29), v44+int32(160))
	mBase = m.M
	v5047 = m.ExcPending
	if v5047 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L682
	}
L680:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(415), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v5028 = m.ExcPending
	if v5028 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L681
	}
L681:
	;
	goto L3
L682:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(421), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v5062 = m.ExcPending
	if v5062 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L683
	}
L683:
	;
	goto L3
L684:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5079 = m.ExcPending
	if v5079 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L685
	}
L685:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v5092 = m.ExcPending
	if v5092 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L686
	}
L686:
	;
	if v4929&int32(4) != 0 {
		goto L687
	} else {
		goto L688
	}
L687:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+212)) = int32(_a_F_AlterSubscription_30)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+208)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_27), v44+int32(208))
	mBase = m.M
	v5113 = m.ExcPending
	if v5113 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L690
	}
L688:
	;
	goto L689
L689:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+196)) = int32(_a_F_AlterSubscription_31)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+192)) = int32(_a_F_AlterSubscription_26)
	F_errmsg(m, int32(_a_F_AlterSubscription_29), v44+int32(192))
	mBase = m.M
	v5147 = m.ExcPending
	if v5147 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L692
	}
L690:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(431), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v5128 = m.ExcPending
	if v5128 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L691
	}
L691:
	;
	goto L3
L692:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(437), int32(_a_F_AlterSubscription_23))
	mBase = m.M
	v5162 = m.ExcPending
	if v5162 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L693
	}
L693:
	;
	goto L3
L694:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5212 = m.ExcPending
	if v5212 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L697
	}
L695:
	;
	goto L696
L696:
	;
	v5255 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+853)) = uint8(v5255)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+788)) = v5179
	v5258 = int32(0)
	if v5179 == v5258 {
		goto L701
	} else {
		goto L702
	}
L697:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errcode(m, int32(325))
	mBase = m.M
	v5225 = m.ExcPending
	if v5225 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L698
	}
L698:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errmsg(m, int32(_a_F_AlterSubscription_72), int32(0))
	mBase = m.M
	v5239 = m.ExcPending
	if v5239 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L699
	}
L699:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1348), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v5254 = m.ExcPending
	if v5254 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L700
	}
L700:
	;
	goto L3
L701:
	;
	v5405 = v53
	v5406 = v54
	v5407 = int32(0)
	v5409 = v5258
	goto L38
L702:
	;
	goto L703
L703:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v5273 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[22])))
	if v5273 == int32(0) {
		goto L705
	} else {
		goto L706
	}
L704:
	;
	v5405 = v53
	v5406 = v54
	v5407 = int32(0)
	v5409 = v5258
	goto L38
L705:
	;
	v5277 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[22])) = uint8(v5277)
	goto L707
L706:
	;
	goto L707
L707:
	;
	goto L704
L708:
	;
	v5315 = int32(112)
	goto L710
L709:
	;
	v5315 = int32(100)
	goto L710
L710:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+800)) = v5315
	v5317 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	v5334 = v5317
	v5335 = v1846
	goto L39
L711:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_CheckAlterSubOption(m, v233, int32(_a_F_AlterSubscription_19), int32(1), v43)
	mBase = m.M
	v5363 = m.ExcPending
	if v5363 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L714
	}
L712:
	;
	v5369 = v5334
	goto L713
L713:
	;
	v5371 = int32(base.Ui32(v5349) >> (uint(int32(13)) % 32))
	if v5369&int32(_a_F_AlterSubscription_21) == int32(0) {
		v5405 = v53
		v5406 = v54
		v5407 = v5371
		v5409 = v5335
		goto L38
	} else {
		goto L715
	}
L714:
	;
	v5364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+751)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+816)) = v5364
	v5366 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+860)) = uint8(v5366)
	v5368 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	v5369 = v5368
	goto L713
L715:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	v5386 = *(*int32)(unsafe.Add(mBase, uint32(v44)+752))
	v5387 = F_cstring_to_text(m, v5386)
	mBase = m.M
	v5388 = m.ExcPending
	if v5388 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L716
	}
L716:
	;
	v5389 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+865)) = uint8(v5389)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+836)) = v5387
	v5405 = v53
	v5406 = v54
	v5407 = v5371
	v5409 = v5335
	goto L38
L717:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v5405
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v5406
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_CatalogTupleUpdate(m, v98, v5439+int32(4), v5439)
	mBase = m.M
	v5454 = m.ExcPending
	if v5454 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L718
	}
L718:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v5405
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v5406
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v83)
	F_pfree(m, v5439)
	mBase = m.M
	v5466 = m.ExcPending
	if v5466 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L719
	}
L719:
	;
	v5480 = v5405
	v5481 = v5406
	v5482 = v5407
	v5484 = v5409
	goto L37
L720:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v5482)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v5480
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v5481
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v5498)
	F_load_file(m, int32(_a_F_AlterSubscription_7), int32(0))
	mBase = m.M
	v5515 = m.ExcPending
	if v5515 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L721
	}
L721:
	;
	v5517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+30)))
	if v5517 == int32(1) {
		goto L722
	} else {
		goto L723
	}
L722:
	;
	v5520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+24)))
	v5523 = v5520 ^ int32(1)
	goto L724
L723:
	;
	v5523 = int32(0)
	goto L724
L724:
	;
	v5524 = *(*int32)(unsafe.Add(mBase, uint32(v233)+16))
	v5525 = *(*int32)(unsafe.Add(mBase, uint32(v233)+36))
	v5527 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[2]))
	v5528 = *(*int32)(unsafe.Add(mBase, uint32(v5527)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v5498)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v5482)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v5480
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v5481
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v5539 = int32(1)
	v5545 = m.T0[v5528].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v5525, v5539, v5539, v5523&v5539, v5524, v44+int32(652))
	mBase = m.M
	v5546 = m.ExcPending
	if v5546 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L725
	}
L725:
	;
	if v5545 == int32(0) {
		goto L726
	} else {
		goto L727
	}
L726:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v5545
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v5482)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v5480
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v5481
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v5498)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5562 = m.ExcPending
	if v5562 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L729
	}
L727:
	;
	goto L728
L728:
	;
	v5612 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[0]))
	v5614 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[1]))
	goto L733
L729:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v5545
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v5482)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v5480
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v5481
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v5498)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v5575 = m.ExcPending
	if v5575 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L730
	}
L730:
	;
	v5576 = *(*int32)(unsafe.Add(mBase, uint32(v233)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v5545
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v5498)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v5482)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v5480
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v5481
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	v5587 = *(*int32)(unsafe.Add(mBase, uint32(v44)+652))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+36)) = v5587
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v5576
	F_errmsg(m, int32(_a_F_AlterSubscription_73), v44+int32(32))
	mBase = m.M
	v5594 = m.ExcPending
	if v5594 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L731
	}
L731:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+912)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+908)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v44)+916)) = v5545
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+923)) = uint8(v5482)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+924)) = v5480
	*(*int32)(unsafe.Add(mBase, uint32(v44)+928)) = v5481
	*(*int32)(unsafe.Add(mBase, uint32(v44)+932)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v44)+936)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v44)+940)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+922)) = uint8(v5498)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1595), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v5609 = m.ExcPending
	if v5609 != 0 {
		v5823 = v40
		v5824 = v41
		v5825 = v42
		v5826 = v43
		v5827 = v44
		v5846 = v63
		v5847 = v64
		goto L6
	} else {
		goto L732
	}
L732:
	;
	goto L3
L733:
	;
	v5616 = v44 + int32(496)
	*(*int32)(unsafe.Add(mBase, uint32(v5616)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5616))) = v44 + int32(492)
	goto L736
L734:
	;
	v5622 = v40
	v5623 = v41
	v5624 = v42
	v5625 = v43
	v5626 = v44
	v5627 = v233
	v5628 = v178
	v5629 = int32(0)
	v5630 = v5612
	v5631 = v5614
	v5632 = v5545
	v5633 = v98
	v5635 = v5480
	v5636 = v5481
	v5637 = v5482
	v5639 = v5484
	v5645 = v63
	v5646 = v64
	goto L11
L736:
	;
	goto L734
L737:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[1])) = v5626 + int32(496)
	v5656 = *(*int32)(unsafe.Add(mBase, uint32(v5627)+40))
	v5658 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[2]))
	v5659 = *(*int32)(unsafe.Add(mBase, uint32(v5658)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+908)) = v5630
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+912)) = v5631
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+916)) = v5632
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+924)) = v5635
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+928)) = v5636
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+932)) = v5627
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+936)) = v5628
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+940)) = v5633
	v5668 = int32(1)
	v5669 = v5639 & v5668
	*(*uint8)(unsafe.Add(mBase, uint32(v5626)+922)) = uint8(v5669)
	v5672 = v5637 & v5668
	*(*uint8)(unsafe.Add(mBase, uint32(v5626)+923)) = uint8(v5672)
	if v5672 != 0 {
		goto L738
	} else {
		goto L739
	}
L738:
	;
	v5675 = v5646
	goto L740
L739:
	;
	v5675 = int32(0)
	goto L740
L740:
	;
	if v5669 != 0 {
		goto L741
	} else {
		goto L742
	}
L741:
	;
	v5677 = v5645
	goto L743
L742:
	;
	v5677 = int32(0)
	goto L743
L743:
	;
	m.T0[v5659].(func(*base.Module, int32, int32, int32, int32))(m, v5632, v5656, v5675, v5677)
	mBase = m.M
	v5679 = m.ExcPending
	if v5679 != 0 {
		v5823 = v5622
		v5824 = v5623
		v5825 = v5624
		v5826 = v5625
		v5827 = v5626
		v5846 = v5645
		v5847 = v5646
		goto L6
	} else {
		goto L744
	}
L744:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[0])) = v5630
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[1])) = v5631
	v5685 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[2]))
	v5686 = *(*int32)(unsafe.Add(mBase, uint32(v5685)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+908)) = v5630
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+912)) = v5631
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+916)) = v5632
	*(*uint8)(unsafe.Add(mBase, uint32(v5626)+922)) = uint8(v5669)
	*(*uint8)(unsafe.Add(mBase, uint32(v5626)+923)) = uint8(v5672)
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+924)) = v5635
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+928)) = v5636
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+932)) = v5627
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+936)) = v5628
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+940)) = v5633
	m.T0[v5686].(func(*base.Module, int32))(m, v5632)
	mBase = m.M
	v5698 = m.ExcPending
	if v5698 != 0 {
		v5823 = v5622
		v5824 = v5623
		v5825 = v5624
		v5826 = v5625
		v5827 = v5626
		v5846 = v5645
		v5847 = v5646
		goto L6
	} else {
		goto L745
	}
L745:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[0])) = v5630
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[1])) = v5631
	v5703 = v5622
	v5704 = v5623
	v5705 = v5624
	v5706 = v5625
	v5707 = v5626
	v5708 = v5627
	v5709 = v5628
	v5711 = v5630
	v5712 = v5631
	v5713 = v5632
	v5714 = v5633
	v5716 = v5635
	v5717 = v5636
	v5718 = v5637
	v5720 = v5639
	v5726 = v5645
	v5727 = v5646
	goto L8
L746:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5703)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5703)+4)) = v5709
	*(*int32)(unsafe.Add(mBase, uint32(v5703))) = int32(_a_F_AlterSubscription_0)
	v5756 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[23]))
	if v5756 != 0 {
		goto L747
	} else {
		goto L748
	}
L747:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+912)) = v5712
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+908)) = v5711
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+916)) = v5713
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+924)) = v5716
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+928)) = v5717
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+932)) = v5708
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+936)) = v5709
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+940)) = v5714
	*(*uint8)(unsafe.Add(mBase, uint32(v5707)+922)) = uint8(v5742)
	*(*uint8)(unsafe.Add(mBase, uint32(v5707)+923)) = uint8(v5745)
	v5768 = int32(0)
	F_RunObjectPostAlterHook(m, int32(_a_F_AlterSubscription_0), v5709, v5768, v5768, v5768)
	mBase = m.M
	v5772 = m.ExcPending
	if v5772 != 0 {
		v5823 = v5703
		v5824 = v5704
		v5825 = v5705
		v5826 = v5706
		v5827 = v5707
		v5846 = v5726
		v5847 = v5727
		goto L6
	} else {
		goto L750
	}
L748:
	;
	goto L749
L749:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+912)) = v5712
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+908)) = v5711
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+916)) = v5713
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+924)) = v5716
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+928)) = v5717
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+932)) = v5708
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+936)) = v5709
	*(*int32)(unsafe.Add(mBase, uint32(v5707)+940)) = v5714
	*(*uint8)(unsafe.Add(mBase, uint32(v5707)+922)) = uint8(v5742)
	*(*uint8)(unsafe.Add(mBase, uint32(v5707)+923)) = uint8(v5745)
	F_LogicalRepWorkersWakeupAtCommit(m, v5709)
	mBase = m.M
	v5784 = m.ExcPending
	if v5784 != 0 {
		v5823 = v5703
		v5824 = v5704
		v5825 = v5705
		v5826 = v5706
		v5827 = v5707
		v5846 = v5726
		v5847 = v5727
		goto L6
	} else {
		goto L751
	}
L750:
	;
	goto L749
L751:
	;
	m.G0 = v5707 + int32(944)
	return
L752:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+912)) = v5631
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+908)) = v5630
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+916)) = v5632
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+924)) = v5635
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+928)) = v5636
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+932)) = v5627
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+936)) = v5628
	*(*int32)(unsafe.Add(mBase, uint32(v5626)+940)) = v5633
	*(*uint8)(unsafe.Add(mBase, uint32(v5626)+922)) = uint8(v5799)
	*(*uint8)(unsafe.Add(mBase, uint32(v5626)+923)) = uint8(v5802)
	F_pg_re_throw(m)
	mBase = m.M
	v5822 = m.ExcPending
	if v5822 != 0 {
		v5823 = v5622
		v5824 = v5623
		v5825 = v5624
		v5826 = v5625
		v5827 = v5626
		v5846 = v5645
		v5847 = v5646
		goto L6
	} else {
		goto L753
	}
L753:
	;
	goto L5
L754:
	;
	v5858 = int32(v5854)
	m.G0 = v5827
	v5860 = *(*int32)(unsafe.Add(mBase, uint32(v5858)+4))
	v5861 = *(*int32)(unsafe.Add(mBase, uint32(v5858)))
	v5864 = *(*int32)(unsafe.Add(mBase, uint32(v5861)))
	if v5827+int32(492) == v5864 {
		goto L757
	} else {
		goto L758
	}
L755:
	;
	m.ExcPending = 1
	goto L763
L756:
	;
	if v5868 != 0 {
		goto L760
	} else {
		goto L761
	}
L757:
	;
	v5866 = *(*int32)(unsafe.Add(mBase, uint32(v5861)+4))
	v5868 = v5866
	goto L759
L758:
	;
	v5868 = int32(0)
	goto L759
L759:
	;
	goto L756
L760:
	;
	v5869 = *(*int32)(unsafe.Add(mBase, uint32(v5827)+940))
	v5870 = *(*int32)(unsafe.Add(mBase, uint32(v5827)+936))
	v5871 = *(*int32)(unsafe.Add(mBase, uint32(v5827)+932))
	v5872 = *(*int32)(unsafe.Add(mBase, uint32(v5827)+928))
	v5873 = *(*int32)(unsafe.Add(mBase, uint32(v5827)+924))
	v5874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5827)+923)))
	v5875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5827)+922)))
	v5876 = *(*int32)(unsafe.Add(mBase, uint32(v5827)+916))
	v5877 = *(*int32)(unsafe.Add(mBase, uint32(v5827)+912))
	v5878 = *(*int32)(unsafe.Add(mBase, uint32(v5827)+908))
	v40 = v5823
	v41 = v5824
	v42 = v5825
	v43 = v5826
	v44 = v5827
	v45 = v5871
	v46 = v5870
	v47 = v5860
	v48 = v5878
	v49 = v5877
	v50 = v5876
	v51 = v5869
	v52 = v5868
	v53 = v5873
	v54 = v5872
	v55 = v5874
	v57 = v5875
	v63 = v5846
	v64 = v5847
	goto L1
L761:
	;
	goto L762
L762:
	;
	F___wasm_longjmp(m, v5861, v5860)
	mBase = m.M
	v5880 = m.ExcPending
	if v5880 != 0 {
		goto L763
	} else {
		goto L764
	}
L763:
	;
	return
L764:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
