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
	var v37 int64
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v110 int32
	_ = v110
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v200 int64
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v355 int32
	_ = v355
	var v356 int64
	_ = v356
	var v358 int32
	_ = v358
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v419 int32
	_ = v419
	var v420 int64
	_ = v420
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v492 int64
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v701 int32
	_ = v701
	var v726 int32
	_ = v726
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v773 int32
	_ = v773
	var v776 int64
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int64
	_ = v779
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v804 int32
	_ = v804
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int64
	_ = v835
	var v837 int32
	_ = v837
	var v838 int64
	_ = v838
	var v839 int32
	_ = v839
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int64
	_ = v850
	var v851 int32
	_ = v851
	var v852 int64
	_ = v852
	var v855 int64
	_ = v855
	var v856 int64
	_ = v856
	var v860 int64
	_ = v860
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v873 int64
	_ = v873
	var v875 int64
	_ = v875
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int64
	_ = v885
	var v888 int64
	_ = v888
	var v892 int64
	_ = v892
	var v893 int64
	_ = v893
	var v896 int64
	_ = v896
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int64
	_ = v912
	var v915 int32
	_ = v915
	var v934 int32
	_ = v934
	var v958 int64
	_ = v958
	var v960 int64
	_ = v960
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v965 int64
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1026 int64
	_ = v1026
	var v1027 int64
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int64
	_ = v1033
	var v1037 int64
	_ = v1037
	var v1039 int64
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1100 int64
	_ = v1100
	var v1104 int64
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1137 int32
	_ = v1137
	var v1199 int32
	_ = v1199
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1240 int32
	_ = v1240
	var v1259 int64
	_ = v1259
	var v1267 int32
	_ = v1267
	var v1268 int64
	_ = v1268
	var v1270 int64
	_ = v1270
	var v1274 int64
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1316 int64
	_ = v1316
	var v1366 int32
	_ = v1366
	var v1369 int64
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1377 int64
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1397 int32
	_ = v1397
	var v1398 int64
	_ = v1398
	var v1401 int64
	_ = v1401
	var v1407 int32
	_ = v1407
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1438 int32
	_ = v1438
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1486 int32
	_ = v1486
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1514 int32
	_ = v1514
	var v1535 int32
	_ = v1535
	var v1561 int32
	_ = v1561
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1587 int32
	_ = v1587
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1685 int32
	_ = v1685
	var v1727 int32
	_ = v1727
	var v1754 int32
	_ = v1754
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1932 int32
	_ = v1932
	var v1937 int32
	_ = v1937
	var v1941 int32
	_ = v1941
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1946 int64
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1951 int64
	_ = v1951
	var v1954 int64
	_ = v1954
	var v1955 int64
	_ = v1955
	var v1959 int64
	_ = v1959
	var v1965 int32
	_ = v1965
	var v1970 int32
	_ = v1970
	var v1974 int32
	_ = v1974
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1979 int64
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1981 int64
	_ = v1981
	var v1984 int64
	_ = v1984
	var v1985 int64
	_ = v1985
	var v1989 int64
	_ = v1989
	var v1995 int32
	_ = v1995
	var v2000 int32
	_ = v2000
	var v2005 int32
	_ = v2005
	var v2008 int32
	_ = v2008
	var v2012 int32
	_ = v2012
	var v2017 int32
	_ = v2017
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2116 int32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2131 int32
	_ = v2131
	var v2135 int32
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2153 int32
	_ = v2153
	var v2157 int32
	_ = v2157
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2168 int32
	_ = v2168
	var v2200 int64
	_ = v2200
	var v2201 int64
	_ = v2201
	var v2207 int32
	_ = v2207
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2233 int32
	_ = v2233
	var v2236 int64
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2253 int32
	_ = v2253
	var v2255 int64
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2258 int64
	_ = v2258
	var v2259 int64
	_ = v2259
	var v2260 int64
	_ = v2260
	var v2262 int64
	_ = v2262
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2305 int64
	_ = v2305
	var v2306 int64
	_ = v2306
	var v2312 int32
	_ = v2312
	var v2350 int64
	_ = v2350
	var v2351 int64
	_ = v2351
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2404 int32
	_ = v2404
	var v2429 int32
	_ = v2429
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2454 int32
	_ = v2454
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2486 int32
	_ = v2486
	var v2488 int32
	_ = v2488
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2507 int32
	_ = v2507
	var v2523 int32
	_ = v2523
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2547 int64
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2567 int32
	_ = v2567
	var v2584 int32
	_ = v2584
	var v2599 int32
	_ = v2599
	var v2619 int32
	_ = v2619
	var v2637 int32
	_ = v2637
	var v2651 int32
	_ = v2651
	var v2655 int32
	_ = v2655
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2722 int64
	_ = v2722
	var v2723 int32
	_ = v2723
	var v2727 int32
	_ = v2727
	var v2730 int32
	_ = v2730
	var v2733 int32
	_ = v2733
	var v2735 int32
	_ = v2735
	var v2737 int32
	_ = v2737
	var v2738 int32
	_ = v2738
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2756 int32
	_ = v2756
	var v2771 int32
	_ = v2771
	var v2774 int32
	_ = v2774
	var v2775 int32
	_ = v2775
	var v2797 int32
	_ = v2797
	var v2811 int32
	_ = v2811
	var v2814 int32
	_ = v2814
	var v2829 int32
	_ = v2829
	var v2886 int32
	_ = v2886
	var v2888 int32
	_ = v2888
	var v2896 int32
	_ = v2896
	var v2897 int64
	_ = v2897
	var v2899 int64
	_ = v2899
	var v2913 int32
	_ = v2913
	var v2917 int32
	_ = v2917
	var v2922 int32
	_ = v2922
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2928 int32
	_ = v2928
	var v2932 int32
	_ = v2932
	var v2935 int32
	_ = v2935
	var v3027 int32
	_ = v3027
	var v3030 int32
	_ = v3030
	var v3039 int32
	_ = v3039
	var v3040 int32
	_ = v3040
	var v3046 int64
	_ = v3046
	var v3048 int32
	_ = v3048
	var v3051 int32
	_ = v3051
	var v3062 int32
	_ = v3062
	var v3065 int32
	_ = v3065
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3070 int32
	_ = v3070
	var v3072 int32
	_ = v3072
	var v3088 int32
	_ = v3088
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3109 int64
	_ = v3109
	var v3124 int32
	_ = v3124
	var v3139 int32
	_ = v3139
	var v3156 int32
	_ = v3156
	var v3161 int32
	_ = v3161
	var v3179 int32
	_ = v3179
	var v3183 int32
	_ = v3183
	var v3188 int32
	_ = v3188
	var v3190 int32
	_ = v3190
	var v3191 int32
	_ = v3191
	var v3194 int32
	_ = v3194
	var v3198 int32
	_ = v3198
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3209 int32
	_ = v3209
	var v3210 int32
	_ = v3210
	var v3216 int32
	_ = v3216
	var v3220 int32
	_ = v3220
	var v3221 int64
	_ = v3221
	var v3236 int64
	_ = v3236
	var v3238 int64
	_ = v3238
	var v3240 int64
	_ = v3240
	var v3241 int64
	_ = v3241
	var v3244 int64
	_ = v3244
	var v3252 int32
	_ = v3252
	var v3253 int32
	_ = v3253
	var v3268 int64
	_ = v3268
	var v3272 int64
	_ = v3272
	var v3274 int64
	_ = v3274
	var v3275 int64
	_ = v3275
	var v3278 int64
	_ = v3278
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3318 int32
	_ = v3318
	var v3319 int32
	_ = v3319
	var v3322 int32
	_ = v3322
	var v3326 int32
	_ = v3326
	var v3327 int32
	_ = v3327
	var v3331 int32
	_ = v3331
	var v3332 int32
	_ = v3332
	var v3338 int32
	_ = v3338
	var v3346 int32
	_ = v3346
	var v3347 int32
	_ = v3347
	var v3348 int32
	_ = v3348
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3391 int32
	_ = v3391
	var v3392 int32
	_ = v3392
	var v3408 int32
	_ = v3408
	var v3412 int32
	_ = v3412
	var v3414 int32
	_ = v3414
	var v3415 int64
	_ = v3415
	var v3423 int32
	_ = v3423
	var v3427 int32
	_ = v3427
	var v3431 int32
	_ = v3431
	var v3437 int32
	_ = v3437
	var v3441 int32
	_ = v3441
	var v3442 int32
	_ = v3442
	var v3449 int32
	_ = v3449
	var v3450 int32
	_ = v3450
	var v3451 int32
	_ = v3451
	var v3455 int32
	_ = v3455
	var v3458 int32
	_ = v3458
	var v3462 int32
	_ = v3462
	var v3463 int32
	_ = v3463
	var v3471 int32
	_ = v3471
	var v3477 int32
	_ = v3477
	var v3479 int32
	_ = v3479
	var v3481 int32
	_ = v3481
	var v3491 int32
	_ = v3491
	var v3508 int32
	_ = v3508
	var v3511 int32
	_ = v3511
	var v3514 int32
	_ = v3514
	var v3517 int32
	_ = v3517
	var v3518 int32
	_ = v3518
	var v3521 int32
	_ = v3521
	var v3522 int32
	_ = v3522
	var v3525 int32
	_ = v3525
	var v3532 int32
	_ = v3532
	var v3533 int32
	_ = v3533
	var v3552 int32
	_ = v3552
	var v3555 int32
	_ = v3555
	var v3558 int32
	_ = v3558
	var v3559 int32
	_ = v3559
	var v3562 int32
	_ = v3562
	var v3563 int32
	_ = v3563
	var v3566 int32
	_ = v3566
	var v3573 int32
	_ = v3573
	var v3574 int32
	_ = v3574
	var v3591 int32
	_ = v3591
	var v3592 int32
	_ = v3592
	var v3606 int32
	_ = v3606
	var v3607 int32
	_ = v3607
	var v3621 int32
	_ = v3621
	var v3625 int32
	_ = v3625
	var v3627 int32
	_ = v3627
	var v3628 int64
	_ = v3628
	var v3636 int32
	_ = v3636
	var v3640 int32
	_ = v3640
	var v3644 int32
	_ = v3644
	var v3650 int32
	_ = v3650
	var v3654 int32
	_ = v3654
	var v3655 int32
	_ = v3655
	var v3662 int32
	_ = v3662
	var v3663 int32
	_ = v3663
	var v3664 int32
	_ = v3664
	var v3668 int32
	_ = v3668
	var v3671 int32
	_ = v3671
	var v3675 int32
	_ = v3675
	var v3676 int32
	_ = v3676
	var v3684 int32
	_ = v3684
	var v3690 int32
	_ = v3690
	var v3692 int32
	_ = v3692
	var v3694 int32
	_ = v3694
	var v3704 int32
	_ = v3704
	var v3721 int32
	_ = v3721
	var v3722 int32
	_ = v3722
	var v3725 int32
	_ = v3725
	var v3728 int32
	_ = v3728
	var v3731 int32
	_ = v3731
	var v3732 int32
	_ = v3732
	var v3735 int32
	_ = v3735
	var v3736 int32
	_ = v3736
	var v3739 int32
	_ = v3739
	var v3746 int32
	_ = v3746
	var v3747 int32
	_ = v3747
	var v3762 int32
	_ = v3762
	var v3763 int32
	_ = v3763
	var v3777 int32
	_ = v3777
	var v3778 int32
	_ = v3778
	var v3780 int32
	_ = v3780
	var v3781 int32
	_ = v3781
	var v3782 int32
	_ = v3782
	var v3783 int32
	_ = v3783
	var v3798 int32
	_ = v3798
	var v3799 int32
	_ = v3799
	var v3812 int32
	_ = v3812
	var v3813 int32
	_ = v3813
	var v3814 int32
	_ = v3814
	var v3816 int32
	_ = v3816
	var v3817 int32
	_ = v3817
	var v3857 int32
	_ = v3857
	var v3858 int32
	_ = v3858
	var v3873 int32
	_ = v3873
	var v3889 int32
	_ = v3889
	var v3908 int32
	_ = v3908
	var v3925 int32
	_ = v3925
	var v3943 int32
	_ = v3943
	var v3944 int32
	_ = v3944
	var v3945 int32
	_ = v3945
	var v3960 int64
	_ = v3960
	var v3973 int32
	_ = v3973
	var v3974 int32
	_ = v3974
	var v3975 int64
	_ = v3975
	var v3976 int64
	_ = v3976
	var v3978 int64
	_ = v3978
	var v3982 int32
	_ = v3982
	var v3983 int32
	_ = v3983
	var v3987 int32
	_ = v3987
	var v4002 int32
	_ = v4002
	var v4004 int32
	_ = v4004
	var v4006 int32
	_ = v4006
	var v4023 int32
	_ = v4023
	var v4042 int32
	_ = v4042
	var v4060 int32
	_ = v4060
	var v4065 int32
	_ = v4065
	var v4100 int64
	_ = v4100
	var v4104 int32
	_ = v4104
	var v4108 int32
	_ = v4108
	var v4123 int64
	_ = v4123
	var v4136 int32
	_ = v4136
	var v4137 int32
	_ = v4137
	var v4139 int64
	_ = v4139
	var v4140 int64
	_ = v4140
	var v4141 int64
	_ = v4141
	var v4143 int64
	_ = v4143
	var v4145 int64
	_ = v4145
	var v4152 int32
	_ = v4152
	var v4153 int32
	_ = v4153
	var v4169 int32
	_ = v4169
	var v4170 int32
	_ = v4170
	var v4172 int32
	_ = v4172
	var v4174 int32
	_ = v4174
	var v4191 int32
	_ = v4191
	var v4210 int32
	_ = v4210
	var v4228 int32
	_ = v4228
	var v4287 int32
	_ = v4287
	var v4289 int32
	_ = v4289
	var v4291 int32
	_ = v4291
	var v4308 int32
	_ = v4308
	var v4327 int32
	_ = v4327
	var v4345 int32
	_ = v4345
	var v4368 int32
	_ = v4368
	var v4392 int32
	_ = v4392
	var v4396 int32
	_ = v4396
	var v4412 int32
	_ = v4412
	var v4417 int32
	_ = v4417
	var v4418 int32
	_ = v4418
	var v4433 int64
	_ = v4433
	var v4446 int32
	_ = v4446
	var v4447 int32
	_ = v4447
	var v4462 int64
	_ = v4462
	var v4463 int64
	_ = v4463
	var v4464 int64
	_ = v4464
	var v4466 int32
	_ = v4466
	var v4467 int32
	_ = v4467
	var v4469 int64
	_ = v4469
	var v4486 int32
	_ = v4486
	var v4500 int32
	_ = v4500
	var v4502 int32
	_ = v4502
	var v4521 int32
	_ = v4521
	var v4536 int32
	_ = v4536
	var v4553 int32
	_ = v4553
	var v4571 int32
	_ = v4571
	var v4590 int32
	_ = v4590
	var v4593 int32
	_ = v4593
	var v4594 int32
	_ = v4594
	var v4611 int32
	_ = v4611
	var v4626 int32
	_ = v4626
	var v4647 int32
	_ = v4647
	var v4665 int32
	_ = v4665
	var v4666 int64
	_ = v4666
	var v4668 int64
	_ = v4668
	var v4683 int32
	_ = v4683
	var v4685 int32
	_ = v4685
	var v4702 int32
	_ = v4702
	var v4717 int32
	_ = v4717
	var v4736 int32
	_ = v4736
	var v4754 int32
	_ = v4754
	var v4770 int32
	_ = v4770
	var v4775 int32
	_ = v4775
	var v4777 int32
	_ = v4777
	var v4783 int32
	_ = v4783
	var v4818 int64
	_ = v4818
	var v4822 int32
	_ = v4822
	var v4823 int64
	_ = v4823
	var v4825 int32
	_ = v4825
	var v4842 int64
	_ = v4842
	var v4844 int64
	_ = v4844
	var v4846 int32
	_ = v4846
	var v4848 int32
	_ = v4848
	var v4849 int32
	_ = v4849
	var v4869 int32
	_ = v4869
	var v4884 int32
	_ = v4884
	var v4905 int32
	_ = v4905
	var v4923 int32
	_ = v4923
	var v4937 int32
	_ = v4937
	var v4939 int32
	_ = v4939
	var v4940 int32
	_ = v4940
	var v4941 int32
	_ = v4941
	var v4956 int32
	_ = v4956
	var v4958 int64
	_ = v4958
	var v4960 int32
	_ = v4960
	var v4964 int64
	_ = v4964
	var v4979 int32
	_ = v4979
	var v4981 int32
	_ = v4981
	var v4998 int32
	_ = v4998
	var v5013 int32
	_ = v5013
	var v5032 int32
	_ = v5032
	var v5050 int32
	_ = v5050
	var v5067 int32
	_ = v5067
	var v5068 int32
	_ = v5068
	var v5086 int32
	_ = v5086
	var v5091 int32
	_ = v5091
	var v5092 int32
	_ = v5092
	var v5110 int32
	_ = v5110
	var v5112 int32
	_ = v5112
	var v5113 int32
	_ = v5113
	var v5160 int32
	_ = v5160
	var v5161 int32
	_ = v5161
	var v5168 int32
	_ = v5168
	var v5207 int32
	_ = v5207
	var v5211 int32
	_ = v5211
	var v5227 int32
	_ = v5227
	var v5232 int32
	_ = v5232
	var v5233 int32
	_ = v5233
	var v5251 int32
	_ = v5251
	var v5268 int32
	_ = v5268
	var v5283 int32
	_ = v5283
	var v5302 int32
	_ = v5302
	var v5320 int32
	_ = v5320
	var v5335 int32
	_ = v5335
	var v5338 int32
	_ = v5338
	var v5344 int32
	_ = v5344
	var v5348 int32
	_ = v5348
	var v5349 int32
	_ = v5349
	var v5370 int32
	_ = v5370
	var v5371 int32
	_ = v5371
	var v5387 int32
	_ = v5387
	var v5389 int32
	_ = v5389
	var v5390 int32
	_ = v5390
	var v5435 int32
	_ = v5435
	var v5437 int32
	_ = v5437
	var v5439 int32
	_ = v5439
	var v5440 int32
	_ = v5440
	var v5456 int32
	_ = v5456
	var v5457 int32
	_ = v5457
	var v5458 int32
	_ = v5458
	var v5473 int32
	_ = v5473
	var v5486 int32
	_ = v5486
	var v5487 int32
	_ = v5487
	var v5488 int32
	_ = v5488
	var v5517 int32
	_ = v5517
	var v5518 int64
	_ = v5518
	var v5533 int32
	_ = v5533
	var v5535 int32
	_ = v5535
	var v5538 int32
	_ = v5538
	var v5539 int32
	_ = v5539
	var v5542 int32
	_ = v5542
	var v5543 int32
	_ = v5543
	var v5544 int32
	_ = v5544
	var v5547 int32
	_ = v5547
	var v5551 int32
	_ = v5551
	var v5576 int32
	_ = v5576
	var v5577 int32
	_ = v5577
	var v5596 int64
	_ = v5596
	var v5601 int32
	_ = v5601
	var v5605 int32
	_ = v5605
	var v5606 int64
	_ = v5606
	var v5613 int32
	_ = v5613
	var v5617 int64
	_ = v5617
	var v5620 int64
	_ = v5620
	var v5622 int64
	_ = v5622
	var v5623 int64
	_ = v5623
	var v5628 int64
	_ = v5628
	var v5634 int32
	_ = v5634
	var v5639 int32
	_ = v5639
	var v5640 int32
	_ = v5640
	var v5642 int32
	_ = v5642
	var v5644 int32
	_ = v5644
	var v5645 int32
	_ = v5645
	var v5648 int64
	_ = v5648
	var v5649 int32
	_ = v5649
	var v5651 int64
	_ = v5651
	var v5654 int32
	_ = v5654
	var v5655 int32
	_ = v5655
	var v5703 int32
	_ = v5703
	var v5708 int32
	_ = v5708
	var v5713 int32
	_ = v5713
	var v5716 int32
	_ = v5716
	var v5766 int32
	_ = v5766
	var v5767 int32
	_ = v5767
	var v5774 int32
	_ = v5774
	var v5779 int32
	_ = v5779
	var v5783 int32
	_ = v5783
	var v5784 int32
	_ = v5784
	var v5791 int32
	_ = v5791
	var v5796 int32
	_ = v5796
	var v5812 int32
	_ = v5812
	var v5814 int32
	_ = v5814
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
	var v5830 int32
	_ = v5830
	var v5833 int32
	_ = v5833
	var v5836 int32
	_ = v5836
	var v5838 int32
	_ = v5838
	var v5840 int32
	_ = v5840
	var v5841 int32
	_ = v5841
	var v5843 int32
	_ = v5843
	var v5848 int32
	_ = v5848
	var v5857 int32
	_ = v5857
	var v5860 int32
	_ = v5860
	var v5863 int32
	_ = v5863
	var v5864 int32
	_ = v5864
	var v5865 int32
	_ = v5865
	var v5868 int32
	_ = v5868
	var v5869 int32
	_ = v5869
	var v5870 int32
	_ = v5870
	var v5871 int32
	_ = v5871
	var v5873 int32
	_ = v5873
	var v5874 int64
	_ = v5874
	var v5895 int32
	_ = v5895
	var v5916 int64
	_ = v5916
	var v5917 int64
	_ = v5917
	var v5920 int32
	_ = v5920
	var v5921 int32
	_ = v5921
	var v5922 int64
	_ = v5922
	var v5923 int64
	_ = v5923
	var v5925 int64
	_ = v5925
	var v5926 int32
	_ = v5926
	var v5928 int32
	_ = v5928
	var v5929 int32
	_ = v5929
	var v5930 int32
	_ = v5930
	var v5932 int32
	_ = v5932
	var v5933 int64
	_ = v5933
	var v5934 int32
	_ = v5934
	var v5935 int64
	_ = v5935
	var v5980 int32
	_ = v5980
	var v5981 int32
	_ = v5981
	var v5983 int32
	_ = v5983
	var v5984 int32
	_ = v5984
	var v5986 int32
	_ = v5986
	var v6036 int32
	_ = v6036
	var v6037 int32
	_ = v6037
	var v6044 int32
	_ = v6044
	var v6047 int32
	_ = v6047
	var v6050 int32
	_ = v6050
	var v6052 int32
	_ = v6052
	var v6056 int32
	_ = v6056
	var v6061 int32
	_ = v6061
	var v6065 int32
	_ = v6065
	var v6067 int32
	_ = v6067
	var v6071 int32
	_ = v6071
	var v6076 int32
	_ = v6076
	var v6077 int32
	_ = v6077
	var v6078 int32
	_ = v6078
	var v6093 int32
	_ = v6093
	var v6095 int64
	_ = v6095
	var v6115 int32
	_ = v6115
	var v6116 int32
	_ = v6116
	var v6133 int64
	_ = v6133
	var v6141 int32
	_ = v6141
	var v6159 int32
	_ = v6159
	var v6177 int32
	_ = v6177
	var v6193 int32
	_ = v6193
	var v6210 int32
	_ = v6210
	var v6228 int32
	_ = v6228
	var v6243 int32
	_ = v6243
	var v6244 int32
	_ = v6244
	var v6246 int32
	_ = v6246
	var v6264 int32
	_ = v6264
	var v6280 int32
	_ = v6280
	var v6284 int32
	_ = v6284
	var v6289 int32
	_ = v6289
	var v6292 int32
	_ = v6292
	var v6294 int32
	_ = v6294
	var v6295 int32
	_ = v6295
	var v6298 int32
	_ = v6298
	var v6302 int32
	_ = v6302
	var v6304 int32
	_ = v6304
	var v6305 int32
	_ = v6305
	var v6313 int32
	_ = v6313
	var v6314 int32
	_ = v6314
	var v6320 int32
	_ = v6320
	var v6324 int32
	_ = v6324
	var v6325 int32
	_ = v6325
	var v6326 int32
	_ = v6326
	var v6327 int32
	_ = v6327
	var v6351 int32
	_ = v6351
	var v6355 int32
	_ = v6355
	var v6356 int32
	_ = v6356
	var v6367 int32
	_ = v6367
	var v6368 int64
	_ = v6368
	var v6372 int32
	_ = v6372
	var v6374 int32
	_ = v6374
	var v6375 int32
	_ = v6375
	var v6378 int32
	_ = v6378
	var v6380 int32
	_ = v6380
	var v6382 int32
	_ = v6382
	var v6383 int32
	_ = v6383
	var v6384 int32
	_ = v6384
	var v6385 int32
	_ = v6385
	var v6386 int32
	_ = v6386
	var v6387 int32
	_ = v6387
	var v6388 int32
	_ = v6388
	var v6389 int32
	_ = v6389
	var v6390 int32
	_ = v6390
	var v6391 int64
	_ = v6391
	var v6392 int64
	_ = v6392
	var v6393 int32
	_ = v6393
	var v6394 int32
	_ = v6394
	var v6395 int32
	_ = v6395
	var v6397 int32
	_ = v6397
	v4 = int32(0)
	v37 = int64(0)
	v44 = m.G0
	v46 = v44 - int32(2224)
	m.G0 = v46
	v55 = l0
	v56 = l1
	v57 = l2
	v58 = v46
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
	v69 = v4
	v70 = v4
	v73 = int32(-1)
	v82 = v46 + int32(2104)
	v86 = v46 + int32(2128)
	v87 = v46 + int32(2120)
	v91 = v37
	v92 = v37
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
	m.G0 = v58 + int32(2224)
	return
L4:
	;
	if v73 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v6367 = int32(m.ExcTag)
	v6368 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v6367 == int32(0) {
		goto L668
	} else {
		goto L669
	}
L7:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[1])) = v102
	v104 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2088)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2096)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2104)) = v104
	v110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+2112)) = uint8(v110)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v63
	v123 = v58 + int32(2112)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v82
	v130 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[2])))
	if v130 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v389 = v59
	v390 = v60
	v391 = v61
	v392 = v62
	v393 = v63
	v394 = v64
	v395 = v65
	v396 = v66
	v397 = v70
	goto L9
L9:
	;
	if v389 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[3])) = uint8(v140)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v55)+56))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v82
	v157 = m.G0
	v159 = v157 - int32(32)
	m.G0 = v159
	v162 = v58 + int32(2056)
	v163 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v162))) = v163
	*(*int64)(unsafe.Add(mBase, uint32(v162)+24)) = v163
	*(*int64)(unsafe.Add(mBase, uint32(v162)+16)) = v163
	*(*int64)(unsafe.Add(mBase, uint32(v162)+8)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v142
	if v143 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[4]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+316))
	v138 = base.B2i32(v136 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[2])) = uint8(v138)
	v140 = v138
	goto L13
L12:
	;
	v140 = v110
	goto L13
L13:
	;
	goto L10
L14:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_perform_base_backup[5])) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v82
	v262 = F_palloc0(m, int32(1112))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L40
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L27
	}
L16:
	;
	m.G0 = v159 + int32(32)
	goto L14
L17:
	;
	v174 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v162)+26)) = uint8(v174)
	v176 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+24)) = uint16(v176)
	*(*int64)(unsafe.Add(mBase, uint32(v162)+16)) = int64(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v181 = F_BufFileCreateTemp(m, int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v181
	v185 = F_pg_cryptohash_create(m, int32(3))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = v185
	v188 = F_pg_cryptohash_init(m, v185)
	mBase = m.M
	if v188 < int32(0) {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v191 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+25)) = uint16(v191)
	*(*int64)(unsafe.Add(mBase, uint32(v162)+16)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v162)+24)) = uint8(base.B2i32(v143 == int32(2)))
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[6]))
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v199)))
	goto L23
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v159)+16)) = v200
	v205 = F_psprintf(m, int32(_a_F_perform_base_backup_0), v159+int32(16))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L24
	}
L24:
	;
	F_AppendStringToManifest(m, v162, v205)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L25
	}
L25:
	;
	F_pfree(m, v205)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L16
L27:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	if v220 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v235
	F_errmsg_internal(m, int32(_a_F_perform_base_backup_1), v159)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L38
	}
L29:
	;
	v235 = int32(_a_F_perform_base_backup_2)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v227 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v230 = int32(_a_F_perform_base_backup_3)
	goto L34
L33:
	;
	v230 = int32(_a_F_perform_base_backup_4)
	goto L34
L34:
	;
	if v227 == int32(2) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v233 = int32(_a_F_perform_base_backup_2)
	goto L37
L36:
	;
	v233 = v230
	goto L37
L37:
	;
	v235 = v233
	goto L28
L38:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(72), int32(_a_F_perform_base_backup_6))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
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
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v82
	v277 = F_makeStringInfo(m)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v82
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[7]))
	if v296 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+5)))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v82
	F_do_pg_backup_start(m, v338, v337, v58+int32(2088), v262, v277)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L46
	}
L43:
	;
	goto L42
L44:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[8])))
	if v300&int32(1) == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v305 = int32(_a_F_perform_base_backup_7)
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	v308 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v307 + v308
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	*(*int32)(unsafe.Add(mBase, uint32(v296))) = v311 + v308
	v315 = int32(0)
	v317 = int32(_a_F_perform_base_backup_8)
	v318 = base.AtomicRmwOr32(m, v315, v317, v315)
	*(*int64)(unsafe.Add(mBase, uint32(v296+v315)+232)) = int64(1)
	v326 = base.AtomicRmwOr32(m, v315, v317, v315)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	*(*int32)(unsafe.Add(mBase, uint32(v296))) = v327 + v308
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v333 - v308
	goto L43
L46:
	;
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v262)+1032))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2120)) = v356
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v262)+1040))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2128)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v82
	F_before_shmem_exit(m, int32(407), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L47
	}
L47:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[10]))
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[11]))
	goto L48
L48:
	;
	v383 = v58 + int32(1888)
	*(*int32)(unsafe.Add(mBase, uint32(v383)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v383))) = v58 + int32(332)
	goto L51
L49:
	;
	v389 = int32(0)
	v390 = v86
	v391 = v262
	v392 = v82
	v393 = v277
	v394 = v87
	v395 = v379
	v396 = v381
	v397 = v123
	goto L9
L51:
	;
	goto L49
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v2886 = m.G0
	v2888 = v2886 - int32(32)
	m.G0 = v2888
	*(*int64)(unsafe.Add(mBase, uint32(v2888)+24)) = int64(17179869184)
	*(*int64)(unsafe.Add(mBase, uint32(v2888))) = int64(4)
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(v58+int32(2088))))
	if v2896 != 0 {
		goto L319
	} else {
		goto L320
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[11])) = v58 + int32(1888)
	if v57 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[10])) = v395
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[11])) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_cancel_before_shmem_exit(m, int32(407), int32(0))
	mBase = m.M
	v2797 = m.ExcPending
	if v2797 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L316
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v419 = int32(0)
	v420 = int64(0)
	v423 = m.G0
	v425 = v423 - int32(2336)
	m.G0 = v425
	v427 = int32(_a_F_perform_base_backup_9)
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[12]))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[12])) = v430
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v432 == v419 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v2075 = F_palloc0(m, int32(24))
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L257
	}
L59:
	;
	goto L58
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L253
	}
L61:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	if v435 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v391)+1040))
	v439 = F_readTimeLineHistory(m, v438)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v443 = F_palloc0(m, v435<<(uint(int32(2))%32))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L64
	}
L64:
	;
	if v435 <= int32(0) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L249
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L245
	}
L67:
	;
	v960 = *(*int64)(unsafe.Add(mBase, uint32(v391)+1032))
	F_WaitForWalSummarization(m, v960)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L130
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+1080)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v391)+1072)) = int64(0)
	v934 = v419
	v958 = v420
	goto L67
L69:
	;
	goto L70
L70:
	;
	v468 = v419
	v470 = v419
	v477 = v419
	v492 = v420
	goto L71
L71:
	;
	v495 = v477 << (uint(int32(2)) % 32)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+12))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v495+v497)))
	if v439 != 0 {
		goto L76
	} else {
		goto L77
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+1080)) = v778
	*(*int64)(unsafe.Add(mBase, uint32(v391)+1072)) = v779
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v785)+12))
	v804 = int32(0)
	goto L106
L73:
	;
	if v745&int32(1) != 0 {
		goto L99
	} else {
		goto L100
	}
L74:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v745 = v701
	v746 = v726
	goto L73
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L95
	}
L76:
	;
	v500 = int32(0)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	if v501 <= v500 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	goto L78
L78:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v443+v495)))
	if v622 != 0 {
		v701 = int32(0)
		goto L74
	} else {
		goto L94
	}
L79:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v443+v495)))
	if v615 == int32(0) {
		goto L75
	} else {
		goto L92
	}
L80:
	;
	v589 = v500
	v591 = int32(0)
	goto L79
L81:
	;
	goto L82
L82:
	;
	v505 = int32(0)
	if v505 < v501 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v508 = v501
	goto L85
L84:
	;
	v508 = v505
	goto L85
L85:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v439)+12))
	v511 = int32(0)
	v529 = v511
	v531 = v500
	v533 = v511
	goto L86
L86:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v510+v529<<(uint(int32(2))%32))))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	if v509 == v560 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v589 = v567
	v591 = v565
	goto L79
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v443+v495))) = v559
	v589 = v531
	v591 = v533
	goto L79
L89:
	;
	goto L90
L90:
	;
	v565 = base.B2i32(v470 == v560) | v533
	v567 = base.B2i32(v468 == v560) | v531
	v569 = v529 + int32(1)
	if v569 != v508 {
		v529 = v569
		v531 = v567
		v533 = v565
		goto L86
	} else {
		goto L91
	}
L91:
	;
	goto L87
L92:
	;
	if v591&int32(1) != 0 {
		v745 = v589
		v746 = v470
		goto L73
	} else {
		goto L93
	}
L93:
	;
	v701 = v589
	goto L74
L94:
	;
	goto L75
L95:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L96
	}
L96:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	*(*int32)(unsafe.Add(mBase, uint32(v425))) = v673
	F_errmsg(m, int32(_a_F_perform_base_backup_10), v425)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(348), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	v773 = int32(0)
	goto L101
L100:
	;
	v773 = v468
	goto L101
L101:
	;
	if v773 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v776 = *(*int64)(unsafe.Add(mBase, uint32(v499)+8))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v778 = v777
	v779 = v776
	goto L104
L103:
	;
	v778 = v468
	v779 = v492
	goto L104
L104:
	;
	v781 = v477 + int32(1)
	if v781 != v435 {
		v468 = v778
		v470 = v746
		v477 = v781
		v492 = v779
		goto L71
	} else {
		goto L105
	}
L105:
	;
	goto L72
L106:
	;
	v832 = v804 << (uint(int32(2)) % 32)
	v833 = v443 + v832
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v833)))
	v835 = *(*int64)(unsafe.Add(mBase, uint32(v834)+8))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v832+v786)))
	v838 = *(*int64)(unsafe.Add(mBase, uint32(v837)+8))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v837)))
	if v778 == v839 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v934 = v778
	v958 = v779
	goto L67
L108:
	;
	v873 = *(*int64)(unsafe.Add(mBase, uint32(v837)+16))
	if v746 == v839 {
		goto L119
	} else {
		goto L120
	}
L109:
	;
	if base.Ui64(v835) <= base.Ui64(v838) {
		goto L108
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	if v835 != v838 {
		goto L65
	} else {
		goto L117
	}
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L113
	}
L113:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L114
	}
L114:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v837)))
	v850 = *(*int64)(unsafe.Add(mBase, uint32(v837)+8))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v833)))
	v852 = *(*int64)(unsafe.Add(mBase, uint32(v851)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+32)) = uint32(v852)
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+24)) = uint32(v850)
	v855 = int64(32)
	v856 = int64(base.Ui64(v850) >> (uint(v855) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+20)) = uint32(v856)
	*(*int32)(unsafe.Add(mBase, uint32(v425)+16)) = v849
	v860 = int64(base.Ui64(v852) >> (uint(v855) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+28)) = uint32(v860)
	F_errmsg(m, int32(_a_F_perform_base_backup_13), v425+int32(16))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(415), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	goto L108
L118:
	;
	v915 = v804 + int32(1)
	if v915 != v435 {
		v804 = v915
		goto L106
	} else {
		goto L129
	}
L119:
	;
	v875 = *(*int64)(unsafe.Add(mBase, uint32(v391)+1032))
	if base.Ui64(v873) <= base.Ui64(v875) {
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v912 = *(*int64)(unsafe.Add(mBase, uint32(v834)+16))
	if v873 != v912 {
		goto L66
	} else {
		goto L128
	}
L122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L123
	}
L123:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L124
	}
L124:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v837)))
	v885 = *(*int64)(unsafe.Add(mBase, uint32(v837)+16))
	v888 = *(*int64)(unsafe.Add(mBase, uint32(v391)+1032))
	*(*uint32)(unsafe.Add(mBase, uint32(v425-int32(-64)))) = uint32(v888)
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+56)) = uint32(v885)
	*(*int32)(unsafe.Add(mBase, uint32(v425)+48)) = v884
	v892 = int64(32)
	v893 = int64(base.Ui64(v888) >> (uint(v892) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+60)) = uint32(v893)
	v896 = int64(base.Ui64(v885) >> (uint(v892) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+52)) = uint32(v896)
	F_errmsg(m, int32(_a_F_perform_base_backup_14), v425+int32(48))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L125
	}
L125:
	;
	F_errhint(m, int32(_a_F_perform_base_backup_15), int32(0))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(437), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	goto L118
L129:
	;
	goto L107
L130:
	;
	v963 = int32(0)
	v965 = *(*int64)(unsafe.Add(mBase, uint32(v391)+1032))
	v966 = F_GetWalSummaries(m, v963, v958, v965)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L131
	}
L131:
	;
	if v439 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v425)+112)) = v1383
	*(*int32)(unsafe.Add(mBase, uint32(v425)+108)) = v1379
	*(*int32)(unsafe.Add(mBase, uint32(v425)+104)) = v1384
	*(*int32)(unsafe.Add(mBase, uint32(v425)+100)) = v1382
	*(*int32)(unsafe.Add(mBase, uint32(v425)+96)) = v1385
	F_errmsg(m, int32(_a_F_perform_base_backup_16), v425+int32(96))
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L243
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[12])) = v428
	m.G0 = v425 + int32(2336)
	goto L59
L134:
	;
	v970 = F_CreateEmptyBlockRefTable(m)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	if v973 <= int32(0) {
		v1486 = v963
		goto L138
	} else {
		goto L139
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+28)) = v970
	goto L133
L138:
	;
	v1509 = F_CreateEmptyBlockRefTable(m)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L204
	}
L139:
	;
	v976 = int32(0)
	v996 = v976
	v997 = v976
	v998 = v963
	goto L140
L140:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v439)+12))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1021+v996<<(uint(int32(2))%32))))
	v1026 = *(*int64)(unsafe.Add(mBase, uint32(v1025)+16))
	v1027 = *(*int64)(unsafe.Add(mBase, uint32(v1025)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v425)+240)) = int64(0)
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1025)))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v391)+1040))
	if v1030 == v1031 {
		goto L144
	} else {
		goto L145
	}
L141:
	;
	v1486 = v1438
	goto L138
L142:
	;
	v1463 = v996 + int32(1)
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	if v1463 < v1464 {
		v996 = v1463
		v997 = v1461
		v998 = v1438
		goto L140
	} else {
		goto L203
	}
L143:
	;
	if v934 == v1030 {
		goto L150
	} else {
		goto L151
	}
L144:
	;
	v1033 = *(*int64)(unsafe.Add(mBase, uint32(v391)+1032))
	v1037 = v1033
	goto L143
L145:
	;
	goto L146
L146:
	;
	if v997&int32(1) != 0 {
		v1037 = v1026
		goto L143
	} else {
		goto L147
	}
L147:
	;
	v1438 = v998
	v1461 = int32(0)
	goto L142
L148:
	;
	if v1366 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L149:
	;
	if v1199 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L150:
	;
	v1039 = v958
	goto L152
L151:
	;
	v1039 = v1027
	goto L152
L152:
	;
	v1040 = int32(0)
	if v966 == v1040 {
		v1199 = v1040
		goto L149
	} else {
		goto L153
	}
L153:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v966)+4))
	if int32(0) < v1045 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v1072 = v1040
	v1074 = v1040
	goto L157
L155:
	;
	v1137 = v1040
	goto L156
L156:
	;
	v1199 = v1137
	goto L149
L157:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v966)+12))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1091+v1074<<(uint(int32(2))%32))))
	if v1030 != 0 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	v1137 = v1108
	goto L156
L159:
	;
	v1110 = v1074 + int32(1)
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v966)+4))
	if v1110 < v1111 {
		v1072 = v1108
		v1074 = v1110
		goto L157
	} else {
		goto L173
	}
L160:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+16))
	if v1030 != v1096 {
		v1108 = v1072
		goto L159
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if v1039 != int64(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L162
L164:
	;
	v1100 = *(*int64)(unsafe.Add(mBase, uint32(v1095)+8))
	if base.Ui64(v1100) < base.Ui64(v1039) {
		v1108 = v1072
		goto L159
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	if v1037 != int64(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	goto L166
L168:
	;
	v1104 = *(*int64)(unsafe.Add(mBase, uint32(v1095)))
	if base.Ui64(v1037) < base.Ui64(v1104) {
		v1108 = v1072
		goto L159
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v1106 = F_lappend(m, v1072, v1095)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L172
	}
L171:
	;
	goto L170
L172:
	;
	v1108 = v1106
	goto L159
L173:
	;
	goto L158
L174:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v425+int32(240)))) = v1316
	v1366 = int32(0)
	goto L148
L175:
	;
	v1316 = int64(0)
	goto L174
L176:
	;
	goto L177
L177:
	;
	v1205 = F_list_copy(m, v1199)
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L178
	}
L178:
	;
	F_list_sort(m, v1205, int32(459))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L179
	}
L179:
	;
	if v1205 == int32(0) {
		v1316 = v1039
		goto L174
	} else {
		goto L180
	}
L180:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+4))
	if v1212 <= int32(0) {
		v1316 = v1039
		goto L174
	} else {
		goto L181
	}
L181:
	;
	v1215 = int32(0)
	if v1215 < v1212 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v1219 = v1212
	goto L184
L183:
	;
	v1219 = v1215
	goto L184
L184:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+12))
	v1240 = v1215
	v1259 = v1039
	goto L185
L185:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1220+v1240<<(uint(int32(2))%32))))
	v1268 = *(*int64)(unsafe.Add(mBase, uint32(v1267)))
	if base.Ui64(v1259) < base.Ui64(v1268) {
		v1316 = v1259
		goto L174
	} else {
		goto L187
	}
L186:
	;
	v1316 = v1274
	goto L174
L187:
	;
	v1270 = *(*int64)(unsafe.Add(mBase, uint32(v1267)+8))
	if base.Ui64(v1270) <= base.Ui64(v1259) {
		v1274 = v1259
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1276 = v1240 + int32(1)
	if v1276 != v1219 {
		v1240 = v1276
		v1259 = v1274
		goto L185
	} else {
		goto L191
	}
L189:
	;
	if base.Ui64(v1270) < base.Ui64(v1037) {
		v1274 = v1270
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v1366 = int32(1)
	goto L148
L191:
	;
	goto L186
L192:
	;
	v1369 = *(*int64)(unsafe.Add(mBase, uint32(v425)+240))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v1413 = F_list_concat(m, v998, v1199)
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L201
	}
L195:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L196
	}
L196:
	;
	v1377 = int64(32)
	v1379 = base.I32_wrap_i64(int64(base.Ui64(v1037) >> (uint(v1377) % 64)))
	v1382 = base.I32_wrap_i64(int64(base.Ui64(v1039) >> (uint(v1377) % 64)))
	v1383 = base.I32_wrap_i64(v1037)
	v1384 = base.I32_wrap_i64(v1039)
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1025)))
	if v1369 == int64(0) {
		goto L132
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v425)+160)) = v1383
	*(*int32)(unsafe.Add(mBase, uint32(v425)+156)) = v1379
	*(*int32)(unsafe.Add(mBase, uint32(v425)+152)) = v1384
	*(*int32)(unsafe.Add(mBase, uint32(v425)+148)) = v1382
	*(*int32)(unsafe.Add(mBase, uint32(v425)+144)) = v1385
	F_errmsg(m, int32(_a_F_perform_base_backup_17), v425+int32(144))
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L198
	}
L198:
	;
	v1398 = *(*int64)(unsafe.Add(mBase, uint32(v425)+240))
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+132)) = uint32(v1398)
	v1401 = int64(base.Ui64(v1398) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+128)) = uint32(v1401)
	F_errdetail(m, int32(_a_F_perform_base_backup_18), v425+int32(128))
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(537), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
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
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1025)))
	if v1415 == v934 {
		v1486 = v1413
		goto L138
	} else {
		goto L202
	}
L202:
	;
	v1438 = v1413
	v1461 = int32(1)
	goto L142
L203:
	;
	goto L141
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+28)) = v1509
	if v1486 == int32(0) {
		goto L133
	} else {
		goto L205
	}
L205:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+4))
	if v1514 <= int32(0) {
		goto L133
	} else {
		goto L206
	}
L206:
	;
	v1535 = int32(0)
	goto L207
L207:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+12))
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1561+v1535<<(uint(int32(2))%32))))
	v1566 = F_OpenWalSummaryFile(m, v1565)
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L209
	}
L208:
	;
	goto L133
L209:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v425)+2328)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v425)+2320)) = v1566
	v1573 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L210
	}
L210:
	;
	if v1573 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v425)+2320))
	v1577 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[13]))
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1577+v1575*int32(48))+32))
	goto L214
L212:
	;
	goto L213
L213:
	;
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v425)+2320))
	v1597 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[13]))
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1597+v1595*int32(48))+32))
	goto L217
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v425)+80)) = v1581
	F_errmsg_internal(m, int32(_a_F_perform_base_backup_19), v425+int32(80))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(586), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L216
	}
L216:
	;
	goto L213
L217:
	;
	v1602 = F_CreateBlockRefTableReader(m, v425+int32(2320), v1601)
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L218
	}
L218:
	;
	v1610 = F_BlockRefTableReaderNextRelation(m, v1602, v425+int32(2308), v425+int32(2304), v425+int32(2300))
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L219
	}
L219:
	;
	if v1610 != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	goto L223
L221:
	;
	goto L222
L222:
	;
	F_DestroyBlockRefTableReader(m, v1602)
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L240
	}
L223:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v425)+2304))
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v425)+2300))
	F_BlockRefTableSetLimitBlock(m, v1655, v425+int32(2308), v1658, v1659)
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L225
	}
L224:
	;
	goto L222
L225:
	;
	v1665 = F_BlockRefTableReaderGetBlocks(m, v1602, v425+int32(240), int32(512))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L226
	}
L226:
	;
	if v1665 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v1685 = v1665
	goto L230
L228:
	;
	goto L229
L229:
	;
	v1821 = F_BlockRefTableReaderNextRelation(m, v1602, v425+int32(2308), v425+int32(2304), v425+int32(2300))
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L238
	}
L230:
	;
	v1727 = int32(0)
	goto L232
L231:
	;
	goto L229
L232:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v425)+2304))
	v1759 = v425 + int32(240)
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1759+v1727<<(uint(int32(2))%32))))
	F_BlockRefTableMarkBlockModified(m, v1754, v425+int32(2308), v1757, v1763)
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L234
	}
L233:
	;
	v1770 = F_BlockRefTableReaderGetBlocks(m, v1602, v1759, int32(512))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L236
	}
L234:
	;
	v1767 = v1727 + int32(1)
	if v1767 != v1685 {
		v1727 = v1767
		goto L232
	} else {
		goto L235
	}
L235:
	;
	goto L233
L236:
	;
	if v1770 != 0 {
		v1685 = v1770
		goto L230
	} else {
		goto L237
	}
L237:
	;
	goto L231
L238:
	;
	if v1821 != 0 {
		goto L223
	} else {
		goto L239
	}
L239:
	;
	goto L224
L240:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v425)+2320))
	F_FileClose(m, v1868)
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L241
	}
L241:
	;
	v1872 = v1535 + int32(1)
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+4))
	if v1872 < v1873 {
		v1535 = v1872
		goto L207
	} else {
		goto L242
	}
L242:
	;
	goto L208
L243:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(528), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L244
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L245:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L246
	}
L246:
	;
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v837)))
	v1946 = *(*int64)(unsafe.Add(mBase, uint32(v837)+16))
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v443+v804<<(uint(int32(2))%32))))
	v1951 = *(*int64)(unsafe.Add(mBase, uint32(v1950)+16))
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+192)) = uint32(v1951)
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+184)) = uint32(v1946)
	v1954 = int64(32)
	v1955 = int64(base.Ui64(v1946) >> (uint(v1954) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+180)) = uint32(v1955)
	*(*int32)(unsafe.Add(mBase, uint32(v425)+176)) = v1945
	v1959 = int64(base.Ui64(v1951) >> (uint(v1954) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+188)) = uint32(v1959)
	F_errmsg(m, int32(_a_F_perform_base_backup_20), v425+int32(176))
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(447), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L250
	}
L250:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v837)))
	v1979 = *(*int64)(unsafe.Add(mBase, uint32(v837)+8))
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v833)))
	v1981 = *(*int64)(unsafe.Add(mBase, uint32(v1980)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+224)) = uint32(v1981)
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+216)) = uint32(v1979)
	v1984 = int64(32)
	v1985 = int64(base.Ui64(v1979) >> (uint(v1984) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+212)) = uint32(v1985)
	*(*int32)(unsafe.Add(mBase, uint32(v425)+208)) = v1978
	v1989 = int64(base.Ui64(v1981) >> (uint(v1984) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v425)+220)) = uint32(v1989)
	F_errmsg(m, int32(_a_F_perform_base_backup_21), v425+int32(208))
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(425), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2008 = m.ExcPending
	if v2008 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L254
	}
L254:
	;
	F_errmsg(m, int32(_a_F_perform_base_backup_22), int32(0))
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_11), int32(292), int32(_a_F_perform_base_backup_12))
	mBase = m.M
	v2017 = m.ExcPending
	if v2017 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
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
	*(*int64)(unsafe.Add(mBase, uint32(v2075)+16)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v58)+2088))
	v2093 = F_lappend(m, v2092, v2075)
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2088)) = v2093
	v2096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+4)))
	if v2096 == int32(1) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v2116 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[7]))
	if v2116 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	v2350 = v91
	v2351 = v92
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = int32(_a_F_perform_base_backup_23)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+16)) = v58 + int32(2088)
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(v2362)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	m.T0[v2363].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v2378 = m.ExcPending
	if v2378 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L278
	}
L262:
	;
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v58)+2088))
	if v2157 == int32(0) {
		v2305 = v91
		v2306 = v92
		goto L266
	} else {
		goto L267
	}
L263:
	;
	goto L262
L264:
	;
	v2120 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[8])))
	if v2120&int32(1) == int32(0) {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v2125 = int32(_a_F_perform_base_backup_7)
	v2127 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	v2128 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v2127 + v2128
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2116)))
	*(*int32)(unsafe.Add(mBase, uint32(v2116))) = v2131 + v2128
	v2135 = int32(0)
	v2137 = int32(_a_F_perform_base_backup_8)
	v2138 = base.AtomicRmwOr32(m, v2135, v2137, v2135)
	*(*int64)(unsafe.Add(mBase, uint32(v2116+v2135)+232)) = int64(2)
	v2146 = base.AtomicRmwOr32(m, v2135, v2137, v2135)
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v2116)))
	*(*int32)(unsafe.Add(mBase, uint32(v2116))) = v2147 + v2128
	v2153 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v2153 - v2128
	goto L263
L266:
	;
	v2312 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v397))) = uint8(v2312)
	v2350 = v2305
	v2351 = v2306
	goto L261
L267:
	;
	v2160 = int32(0)
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v2157)+4))
	if v2161 <= v2160 {
		v2305 = v91
		v2306 = v92
		goto L266
	} else {
		goto L268
	}
L268:
	;
	v2168 = v2160
	v2200 = v91
	v2201 = v92
	goto L269
L269:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(v2157)+12))
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(v2207+v2168<<(uint(int32(2))%32))))
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(v2211)+4))
	if v2212 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	v2305 = v2258
	v2306 = v2259
	goto L266
L271:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2211)+16)) = v2260
	v2262 = *(*int64)(unsafe.Add(mBase, uint32(v392)))
	*(*int64)(unsafe.Add(mBase, uint32(v392))) = v2262 + v2260
	v2266 = v2168 + int32(1)
	v2267 = *(*int32)(unsafe.Add(mBase, uint32(v2157)+4))
	if v2266 < v2267 {
		v2168 = v2266
		v2200 = v2258
		v2201 = v2259
		goto L269
	} else {
		goto L277
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2200
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2201
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v2229 = int32(1)
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v58)+2088))
	v2233 = int32(0)
	v2236 = F_sendDir(m, v56, int32(_a_F_perform_base_backup_24), v2229, v2229, v2231, v2229, v2233, v2233, v2233)
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v2211)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2200
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2201
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v2253 = int32(0)
	v2255 = F_sendTablespace(m, v56, v2212, v2238, int32(1), v2253, v2253)
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L276
	}
L275:
	;
	v2258 = v2200
	v2259 = v2236
	v2260 = v2236
	goto L271
L276:
	;
	v2258 = v2255
	v2259 = v2201
	v2260 = v2255
	goto L271
L277:
	;
	goto L270
L278:
	;
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(v58)+2088))
	if v2379 == int32(0) {
		goto L52
	} else {
		goto L279
	}
L279:
	;
	v2382 = int32(0)
	v2383 = *(*int32)(unsafe.Add(mBase, uint32(v2379)+4))
	if v2383 <= v2382 {
		goto L52
	} else {
		goto L280
	}
L280:
	;
	v2404 = v2382
	goto L281
L281:
	;
	v2429 = *(*int32)(unsafe.Add(mBase, uint32(v2379)+12))
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(v2429+v2404<<(uint(int32(2))%32))))
	v2434 = *(*int32)(unsafe.Add(mBase, uint32(v2433)+4))
	if v2434 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	goto L52
L283:
	;
	v2727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+7)))
	if v2727 == int32(1) {
		goto L309
	} else {
		goto L310
	}
L284:
	;
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(v2437)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	m.T0[v2438].(func(*base.Module, int32, int32))(m, v56, int32(_a_F_perform_base_backup_25))
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v2667 = *(*int32)(unsafe.Add(mBase, uint32(v2433)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+320)) = v2667
	v2685 = F_psprintf(m, int32(_a_F_perform_base_backup_26), v58+int32(320))
	mBase = m.M
	v2686 = m.ExcPending
	if v2686 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L305
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v2469 = F_build_backup_content(m, v391, int32(0))
	mBase = m.M
	v2470 = m.ExcPending
	if v2470 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v2486 = v58 + int32(2056)
	F_sendFileWithContent(m, v56, int32(_a_F_perform_base_backup_27), v2469, v2486)
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L289
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_pfree(m, v2469)
	mBase = m.M
	v2503 = m.ExcPending
	if v2503 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L290
	}
L290:
	;
	v2504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+16)))
	if v2504 == int32(1) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_sendFileWithContent(m, v56, int32(_a_F_perform_base_backup_28), v2507, v2486)
	mBase = m.M
	v2523 = m.ExcPending
	if v2523 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v2539 = int32(1)
	v2540 = int32(0)
	v2541 = *(*int32)(unsafe.Add(mBase, uint32(v58)+2088))
	v2547 = F_sendDir(m, v56, int32(_a_F_perform_base_backup_24), v2539, v2540, v2541, v2504^v2539, v58+int32(2056), v2540, v57)
	mBase = m.M
	v2548 = m.ExcPending
	if v2548 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L295
	}
L294:
	;
	goto L293
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v2567 = F___fstatat(m, int32(-100), int32(_a_F_perform_base_backup_29), v58+int32(1792), int32(256))
	mBase = m.M
	goto L296
L296:
	;
	if v2567 != 0 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2584 = m.ExcPending
	if v2584 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v2651 = int32(_a_F_perform_base_backup_29)
	v2655 = int32(0)
	v2665 = F_sendFile(m, v56, v2651, v2651, v58+int32(1792), v2655, v2655, v2655, v2655, v2655, v58+int32(2056), v2655, v2655, v2655)
	mBase = m.M
	v2666 = m.ExcPending
	if v2666 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L304
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errcode_for_file_access(m)
	mBase = m.M
	v2599 = m.ExcPending
	if v2599 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+304)) = int32(_a_F_perform_base_backup_29)
	F_errmsg(m, int32(_a_F_perform_base_backup_30), v58+int32(304))
	mBase = m.M
	v2619 = m.ExcPending
	if v2619 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(358), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v2637 = m.ExcPending
	if v2637 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
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
	v2687 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(v2687)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	m.T0[v2688].(func(*base.Module, int32, int32))(m, v56, v2685)
	mBase = m.M
	v2703 = m.ExcPending
	if v2703 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L306
	}
L306:
	;
	v2704 = *(*int32)(unsafe.Add(mBase, uint32(v2433)))
	v2705 = *(*int32)(unsafe.Add(mBase, uint32(v2433)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v2722 = F_sendTablespace(m, v56, v2705, v2704, int32(0), v58+int32(2056), v57)
	mBase = m.M
	v2723 = m.ExcPending
	if v2723 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L307
	}
L307:
	;
	goto L283
L308:
	;
	v2774 = v2404 + int32(1)
	v2775 = *(*int32)(unsafe.Add(mBase, uint32(v2379)+4))
	if v2774 < v2775 {
		v2404 = v2774
		goto L281
	} else {
		goto L315
	}
L309:
	;
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v2433)+4))
	if v2730 == int32(0) {
		goto L308
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v2735 = int32(1024)
	base.MemoryFill(m, v2733, int32(0), v2735)
	v2737 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v2738 = *(*int32)(unsafe.Add(mBase, uint32(v2737)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	m.T0[v2738].(func(*base.Module, int32, int32))(m, v56, v2735)
	mBase = m.M
	v2754 = m.ExcPending
	if v2754 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L313
	}
L312:
	;
	goto L311
L313:
	;
	v2755 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v2755)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	m.T0[v2756].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v2771 = m.ExcPending
	if v2771 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
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
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v2811 = int32(0)
	F_do_pg_abort_backup(m, v2811, v2811)
	mBase = m.M
	v2814 = m.ExcPending
	if v2814 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_pg_re_throw(m)
	mBase = m.M
	v2829 = m.ExcPending
	if v2829 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L318
	}
L318:
	;
	goto L1
L319:
	;
	v2897 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2896)+4)))
	v2899 = v2897
	goto L321
L320:
	;
	v2899 = int64(0)
	goto L321
L321:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2888)+8)) = v2899
	goto L324
L322:
	;
	m.G0 = v2888 + int32(32)
	v3088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_do_pg_backup_stop(m, v391, (v3088^int32(-1))&int32(1))
	mBase = m.M
	v3107 = m.ExcPending
	if v3107 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L339
	}
L323:
	;
	goto L322
L324:
	;
	v2913 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[7]))
	if v2913 == int32(0) {
		goto L323
	} else {
		goto L325
	}
L325:
	;
	v2917 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[8])))
	if v2917&int32(1) == int32(0) {
		goto L323
	} else {
		goto L326
	}
L326:
	;
	v2922 = int32(_a_F_perform_base_backup_7)
	v2924 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	v2925 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v2924 + v2925
	v2928 = *(*int32)(unsafe.Add(mBase, uint32(v2913)))
	*(*int32)(unsafe.Add(mBase, uint32(v2913))) = v2928 + v2925
	v2932 = int32(0)
	v2935 = base.AtomicRmwOr32(m, v2932, int32(_a_F_perform_base_backup_8), v2932)
	goto L328
L327:
	;
	v3062 = int32(0)
	v3065 = base.AtomicRmwOr32(m, v3062, int32(_a_F_perform_base_backup_8), v3062)
	v3066 = *(*int32)(unsafe.Add(mBase, uint32(v2913)))
	v3067 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2913))) = v3066 + v3067
	v3070 = int32(_a_F_perform_base_backup_7)
	v3072 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v3072 - v3067
	goto L323
L328:
	;
	goto L330
L330:
	;
	goto L331
L331:
	;
	v3027 = int32(0)
	v3030 = int32(0)
	goto L336
L336:
	;
	v3039 = *(*int32)(unsafe.Add(mBase, uint32(v2888+int32(24)+v3030<<(uint(int32(2))%32))))
	v3040 = int32(3)
	v3046 = *(*int64)(unsafe.Add(mBase, uint32(v2888+v3030<<(uint(v3040)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2913+int32(232)+v3039<<(uint(v3040)%32)))) = v3046
	v3048 = int32(1)
	v3051 = v3027 + v3048
	if v3051 != int32(2) {
		v3027 = v3051
		v3030 = v3030 + v3048
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
	v3108 = *(*int32)(unsafe.Add(mBase, uint32(v391)+1096))
	v3109 = *(*int64)(unsafe.Add(mBase, uint32(v391)+1088))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_free_attrmap(m, v393)
	mBase = m.M
	v3124 = m.ExcPending
	if v3124 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L340
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_pfree(m, v391)
	mBase = m.M
	v3139 = m.ExcPending
	if v3139 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_cancel_before_shmem_exit(m, int32(407), int32(0))
	mBase = m.M
	v3156 = m.ExcPending
	if v3156 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L342
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[10])) = v395
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[11])) = v396
	v3161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+7)))
	if v3161 != 0 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3179 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[7]))
	if v3179 == int32(0) {
		goto L347
	} else {
		goto L348
	}
L344:
	;
	v5486 = v67
	v5487 = v68
	v5488 = v69
	goto L345
L345:
	;
	v5517 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v5518 = *(*int64)(unsafe.Add(mBase, uint32(v394)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5487
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5488
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v5533 = m.G0
	v5535 = v5533 - int32(80)
	m.G0 = v5535
	v5538 = v58 + int32(2056)
	v5539 = *(*int32)(unsafe.Add(mBase, uint32(v5538)))
	if v5539 != 0 {
		goto L559
	} else {
		goto L560
	}
L346:
	;
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v3221 = *(*int64)(unsafe.Add(mBase, uint32(v394)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3236 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+288)) = v3220
	v3238 = base.I64_div_u_s(v3221, v3236)
	v3240 = base.I64_div_u_s(int64(4294967296), v3236)
	v3241 = base.I64_div_u_s(v3238, v3240)
	*(*uint32)(unsafe.Add(mBase, uint32(v58)+292)) = uint32(v3241)
	v3244 = v3238 - v3240*v3241
	*(*uint32)(unsafe.Add(mBase, uint32(v58)+296)) = uint32(v3244)
	v3252 = F_pg_snprintf(m, v58+int32(608), int32(64), int32(_a_F_perform_base_backup_33), v58+int32(288))
	mBase = m.M
	v3253 = m.ExcPending
	if v3253 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L350
	}
L347:
	;
	goto L346
L348:
	;
	v3183 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[8])))
	if v3183&int32(1) == int32(0) {
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v3188 = int32(_a_F_perform_base_backup_7)
	v3190 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	v3191 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v3190 + v3191
	v3194 = *(*int32)(unsafe.Add(mBase, uint32(v3179)))
	*(*int32)(unsafe.Add(mBase, uint32(v3179))) = v3194 + v3191
	v3198 = int32(0)
	v3200 = int32(_a_F_perform_base_backup_8)
	v3201 = base.AtomicRmwOr32(m, v3198, v3200, v3198)
	*(*int64)(unsafe.Add(mBase, uint32(v3179+v3198)+232)) = int64(5)
	v3209 = base.AtomicRmwOr32(m, v3198, v3200, v3198)
	v3210 = *(*int32)(unsafe.Add(mBase, uint32(v3179)))
	*(*int32)(unsafe.Add(mBase, uint32(v3179))) = v3210 + v3191
	v3216 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v3216 - v3191
	goto L347
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3268 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+272)) = v3108
	v3272 = base.I64_div_u_s(v3109-int64(1), v3268)
	v3274 = base.I64_div_u_s(int64(4294967296), v3268)
	v3275 = base.I64_div_u_s(v3272, v3274)
	*(*uint32)(unsafe.Add(mBase, uint32(v58)+276)) = uint32(v3275)
	v3278 = v3272 - v3274*v3275
	*(*uint32)(unsafe.Add(mBase, uint32(v58)+280)) = uint32(v3278)
	v3286 = F_pg_snprintf(m, v58+int32(544), int32(64), int32(_a_F_perform_base_backup_33), v58+int32(272))
	mBase = m.M
	v3287 = m.ExcPending
	if v3287 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3302 = F_AllocateDir(m, int32(_a_F_perform_base_backup_34))
	mBase = m.M
	v3303 = m.ExcPending
	if v3303 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L352
	}
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3318 = F_ReadDir(m, v3302, int32(_a_F_perform_base_backup_34))
	mBase = m.M
	v3319 = m.ExcPending
	if v3319 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L354
	}
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_FreeDir(m, v3302)
	mBase = m.M
	v3857 = m.ExcPending
	if v3857 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L433
	}
L354:
	;
	if v3318 == int32(0) {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v3322 = int32(0)
	v3812 = v67
	v3813 = v68
	v3814 = v69
	v3816 = v3322
	v3817 = v3322
	goto L353
L356:
	;
	goto L357
L357:
	;
	v3326 = int32(8)
	v3327 = v58 + int32(544) | v3326
	v3331 = v58 + int32(608) | v3326
	v3332 = int32(0)
	v3338 = v3318
	v3346 = v67
	v3347 = v68
	v3348 = v69
	v3350 = v3332
	v3351 = v3332
	goto L358
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3347
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3346
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3348
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3391 = v3338 + int32(19)
	v3392 = F_strlen(m, v3391)
	mBase = m.M
	switch v3392 - int32(16) {
	case 0:
		goto L361
	default:
		v3780 = v3347
		v3781 = v3348
		v3782 = v3350
		v3783 = v3351
		goto L360
	case 8:
		goto L362
	}
L359:
	;
	v3812 = v3798
	v3813 = v3780
	v3814 = v3781
	v3816 = v3782
	v3817 = v3783
	goto L353
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3346
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3781
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3798 = F_ReadDir(m, v3302, int32(_a_F_perform_base_backup_34))
	mBase = m.M
	v3799 = m.ExcPending
	if v3799 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L431
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3347
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3346
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3348
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3621 = int32(_a_F_perform_base_backup_35)
	v3625 = m.G0
	v3627 = v3625 - int32(32)
	v3628 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3627)+24)) = v3628
	*(*int64)(unsafe.Add(mBase, uint32(v3627)+16)) = v3628
	*(*int64)(unsafe.Add(mBase, uint32(v3627)+8)) = v3628
	*(*int64)(unsafe.Add(mBase, uint32(v3627))) = v3628
	v3636 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[15])))
	if v3636 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3347
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3346
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3348
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3408 = int32(_a_F_perform_base_backup_35)
	v3412 = m.G0
	v3414 = v3412 - int32(32)
	v3415 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3414)+24)) = v3415
	*(*int64)(unsafe.Add(mBase, uint32(v3414)+16)) = v3415
	*(*int64)(unsafe.Add(mBase, uint32(v3414)+8)) = v3415
	*(*int64)(unsafe.Add(mBase, uint32(v3414))) = v3415
	v3423 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[15])))
	if v3423 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L363:
	;
	if v3491 != int32(24) {
		v3780 = v3347
		v3781 = v3348
		v3782 = v3350
		v3783 = v3351
		goto L360
	} else {
		goto L382
	}
L364:
	;
	v3491 = int32(0)
	goto L363
L365:
	;
	goto L366
L366:
	;
	v3427 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[16])))
	if v3427 == int32(0) {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v3431 = v3391
	goto L370
L368:
	;
	goto L369
L369:
	;
	v3441 = v3408
	v3442 = v3423
	goto L373
L370:
	;
	v3437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3431))))
	if v3437 == v3423 {
		v3431 = v3431 + int32(1)
		goto L370
	} else {
		goto L372
	}
L371:
	;
	v3491 = v3431 - v3391
	goto L363
L372:
	;
	goto L371
L373:
	;
	v3449 = v3414 + int32(base.Ui32(v3442)>>(uint(int32(3))%32))&int32(28)
	v3450 = *(*int32)(unsafe.Add(mBase, uint32(v3449)))
	v3451 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3449))) = v3450 | v3451<<(uint(v3442)%32)
	v3455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3441)+1)))
	if v3455 != 0 {
		v3441 = v3441 + v3451
		v3442 = v3455
		goto L373
	} else {
		goto L375
	}
L374:
	;
	v3458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3391))))
	if v3458 == int32(0) {
		v3481 = v3391
		goto L376
	} else {
		goto L377
	}
L375:
	;
	goto L374
L376:
	;
	v3491 = v3481 - v3391
	goto L363
L377:
	;
	v3462 = v3391
	v3463 = v3458
	goto L378
L378:
	;
	v3471 = *(*int32)(unsafe.Add(mBase, uint32(v3414+int32(base.Ui32(v3463)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3471)>>(uint(v3463)%32))&int32(1) == int32(0) {
		v3481 = v3462
		goto L376
	} else {
		goto L380
	}
L379:
	;
	v3481 = v3479
	goto L376
L380:
	;
	v3477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3462)+1)))
	v3479 = v3462 + int32(1)
	if v3477 != 0 {
		v3462 = v3479
		v3463 = v3477
		goto L378
	} else {
		goto L381
	}
L381:
	;
	goto L379
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3347
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3346
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3348
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3508 = v3338 + int32(27)
	v3511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3508))))
	v3514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3331))))
	if base.B2i32(v3511 == int32(0))|base.B2i32(v3511 != v3514) != 0 {
		v3532 = v3511
		v3533 = v3514
		goto L384
	} else {
		goto L385
	}
L383:
	;
	if v3532-v3533 < int32(0) {
		v3780 = v3347
		v3781 = v3348
		v3782 = v3350
		v3783 = v3351
		goto L360
	} else {
		goto L390
	}
L384:
	;
	goto L383
L385:
	;
	v3517 = v3508
	v3518 = v3331
	goto L386
L386:
	;
	v3521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3518)+1)))
	v3522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3517)+1)))
	if v3522 == int32(0) {
		v3532 = v3522
		v3533 = v3521
		goto L384
	} else {
		goto L388
	}
L387:
	;
	v3532 = v3522
	v3533 = v3521
	goto L384
L388:
	;
	v3525 = int32(1)
	if v3522 == v3521 {
		v3517 = v3517 + v3525
		v3518 = v3518 + v3525
		goto L386
	} else {
		goto L389
	}
L389:
	;
	goto L387
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3347
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3346
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3348
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3508))))
	v3555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3327))))
	if base.B2i32(v3552 == int32(0))|base.B2i32(v3552 != v3555) != 0 {
		v3573 = v3552
		v3574 = v3555
		goto L392
	} else {
		goto L393
	}
L391:
	;
	if int32(0) < v3573-v3574 {
		v3780 = v3347
		v3781 = v3348
		v3782 = v3350
		v3783 = v3351
		goto L360
	} else {
		goto L398
	}
L392:
	;
	goto L391
L393:
	;
	v3558 = v3508
	v3559 = v3327
	goto L394
L394:
	;
	v3562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3559)+1)))
	v3563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3558)+1)))
	if v3563 == int32(0) {
		v3573 = v3563
		v3574 = v3562
		goto L392
	} else {
		goto L396
	}
L395:
	;
	v3573 = v3563
	v3574 = v3562
	goto L392
L396:
	;
	v3566 = int32(1)
	if v3563 == v3562 {
		v3558 = v3558 + v3566
		v3559 = v3559 + v3566
		goto L394
	} else {
		goto L397
	}
L397:
	;
	goto L395
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3347
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3346
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3348
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3591 = F_pstrdup(m, v3391)
	mBase = m.M
	v3592 = m.ExcPending
	if v3592 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L399
	}
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3347
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3346
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3348
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3606 = F_lappend(m, v3350, v3591)
	mBase = m.M
	v3607 = m.ExcPending
	if v3607 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L400
	}
L400:
	;
	v3780 = v3347
	v3781 = v3606
	v3782 = v3606
	v3783 = v3351
	goto L360
L401:
	;
	if v3704 != int32(8) {
		v3780 = v3347
		v3781 = v3348
		v3782 = v3350
		v3783 = v3351
		goto L360
	} else {
		goto L420
	}
L402:
	;
	v3704 = int32(0)
	goto L401
L403:
	;
	goto L404
L404:
	;
	v3640 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[16])))
	if v3640 == int32(0) {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v3644 = v3391
	goto L408
L406:
	;
	goto L407
L407:
	;
	v3654 = v3621
	v3655 = v3636
	goto L411
L408:
	;
	v3650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3644))))
	if v3650 == v3636 {
		v3644 = v3644 + int32(1)
		goto L408
	} else {
		goto L410
	}
L409:
	;
	v3704 = v3644 - v3391
	goto L401
L410:
	;
	goto L409
L411:
	;
	v3662 = v3627 + int32(base.Ui32(v3655)>>(uint(int32(3))%32))&int32(28)
	v3663 = *(*int32)(unsafe.Add(mBase, uint32(v3662)))
	v3664 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3662))) = v3663 | v3664<<(uint(v3655)%32)
	v3668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3654)+1)))
	if v3668 != 0 {
		v3654 = v3654 + v3664
		v3655 = v3668
		goto L411
	} else {
		goto L413
	}
L412:
	;
	v3671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3391))))
	if v3671 == int32(0) {
		v3694 = v3391
		goto L414
	} else {
		goto L415
	}
L413:
	;
	goto L412
L414:
	;
	v3704 = v3694 - v3391
	goto L401
L415:
	;
	v3675 = v3391
	v3676 = v3671
	goto L416
L416:
	;
	v3684 = *(*int32)(unsafe.Add(mBase, uint32(v3627+int32(base.Ui32(v3676)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3684)>>(uint(v3676)%32))&int32(1) == int32(0) {
		v3694 = v3675
		goto L414
	} else {
		goto L418
	}
L417:
	;
	v3694 = v3692
	goto L414
L418:
	;
	v3690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3675)+1)))
	v3692 = v3675 + int32(1)
	if v3690 != 0 {
		v3675 = v3692
		v3676 = v3690
		goto L416
	} else {
		goto L419
	}
L419:
	;
	goto L417
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3347
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3346
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3348
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3721 = v3338 + int32(27)
	v3722 = int32(_a_F_perform_base_backup_36)
	v3725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3721))))
	v3728 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[17])))
	if base.B2i32(v3725 == int32(0))|base.B2i32(v3725 != v3728) != 0 {
		v3746 = v3725
		v3747 = v3728
		goto L422
	} else {
		goto L423
	}
L421:
	;
	if v3746-v3747 != 0 {
		v3780 = v3347
		v3781 = v3348
		v3782 = v3350
		v3783 = v3351
		goto L360
	} else {
		goto L428
	}
L422:
	;
	goto L421
L423:
	;
	v3731 = v3721
	v3732 = v3722
	goto L424
L424:
	;
	v3735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3732)+1)))
	v3736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3731)+1)))
	if v3736 == int32(0) {
		v3746 = v3736
		v3747 = v3735
		goto L422
	} else {
		goto L426
	}
L425:
	;
	v3746 = v3736
	v3747 = v3735
	goto L422
L426:
	;
	v3739 = int32(1)
	if v3736 == v3735 {
		v3731 = v3731 + v3739
		v3732 = v3732 + v3739
		goto L424
	} else {
		goto L427
	}
L427:
	;
	goto L425
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3347
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3346
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3348
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3762 = F_pstrdup(m, v3391)
	mBase = m.M
	v3763 = m.ExcPending
	if v3763 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3347
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3346
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3348
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3777 = F_lappend(m, v3351, v3762)
	mBase = m.M
	v3778 = m.ExcPending
	if v3778 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L430
	}
L430:
	;
	v3780 = v3777
	v3781 = v3348
	v3782 = v3350
	v3783 = v3777
	goto L360
L431:
	;
	if v3798 != 0 {
		v3338 = v3798
		v3346 = v3798
		v3347 = v3780
		v3348 = v3781
		v3350 = v3782
		v3351 = v3783
		goto L358
	} else {
		goto L432
	}
L432:
	;
	goto L359
L433:
	;
	v3858 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_CheckXLogRemoved(m, v3238, v3858)
	mBase = m.M
	v3873 = m.ExcPending
	if v3873 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L434
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_list_sort(m, v3816, int32(417))
	mBase = m.M
	v3889 = m.ExcPending
	if v3889 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L435
	}
L435:
	;
	if v3816 == int32(0) {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3908 = m.ExcPending
	if v3908 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L439
	}
L437:
	;
	goto L438
L438:
	;
	v3944 = *(*int32)(unsafe.Add(mBase, uint32(v3816)+12))
	v3945 = *(*int32)(unsafe.Add(mBase, uint32(v3944)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v3960 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+264)) = v58 + int32(2140)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+260)) = v58 + int32(2144)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+256)) = v58 + int32(540)
	v3973 = F_sscanf(m, v3945, int32(_a_F_perform_base_backup_33), v58+int32(256))
	mBase = m.M
	v3974 = m.ExcPending
	if v3974 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L442
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errmsg(m, int32(_a_F_perform_base_backup_37), int32(0))
	mBase = m.M
	v3925 = m.ExcPending
	if v3925 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L440
	}
L440:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(481), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v3943 = m.ExcPending
	if v3943 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L441
	}
L441:
	;
	goto L1
L442:
	;
	v3975 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+2140)))
	v3976 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+2144)))
	v3978 = base.I64_div_u_s(int64(4294967296), v3960)
	if v3238 == v3975+v3976*v3978 {
		goto L447
	} else {
		goto L448
	}
L443:
	;
	if v3817 == int32(0) {
		goto L536
	} else {
		goto L537
	}
L444:
	;
	if v4153 <= int32(0) {
		goto L443
	} else {
		goto L473
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v4287 = v58 + int32(336)
	v4289 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	F_XLogFileName(m, v4287, v3108, v3272, v4289)
	mBase = m.M
	v4291 = m.ExcPending
	if v4291 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L469
	}
L446:
	;
	v4065 = v3982
	v4100 = v3238
	goto L457
L447:
	;
	v3982 = int32(0)
	v3983 = *(*int32)(unsafe.Add(mBase, uint32(v3816)+4))
	if v3982 < v3983 {
		goto L446
	} else {
		goto L450
	}
L448:
	;
	goto L449
L449:
	;
	v3987 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v4002 = v58 + int32(464)
	v4004 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	F_XLogFileName(m, v4002, v3987, v3238, v4004)
	mBase = m.M
	v4006 = m.ExcPending
	if v4006 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L452
	}
L450:
	;
	if v3238 != v3272 {
		goto L445
	} else {
		goto L451
	}
L451:
	;
	goto L443
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4023 = m.ExcPending
	if v4023 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+240)) = v4002
	F_errmsg(m, int32(_a_F_perform_base_backup_38), v58+int32(240))
	mBase = m.M
	v4042 = m.ExcPending
	if v4042 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L454
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(496), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v4060 = m.ExcPending
	if v4060 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L455
	}
L455:
	;
	goto L1
L456:
	;
	if v4145 == v3272 {
		goto L444
	} else {
		goto L468
	}
L457:
	;
	v4104 = *(*int32)(unsafe.Add(mBase, uint32(v3816)+12))
	v4108 = *(*int32)(unsafe.Add(mBase, uint32(v4104+v4065<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v4123 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+232)) = v58 + int32(2148)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+228)) = v58 + int32(2152)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+224)) = v58 + int32(540)
	v4136 = F_sscanf(m, v4108, int32(_a_F_perform_base_backup_33), v58+int32(224))
	mBase = m.M
	v4137 = m.ExcPending
	if v4137 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L459
	}
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v4169 = v58 + int32(400)
	v4170 = *(*int32)(unsafe.Add(mBase, uint32(v58)+540))
	v4172 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	F_XLogFileName(m, v4169, v4170, v4139, v4172)
	mBase = m.M
	v4174 = m.ExcPending
	if v4174 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L464
	}
L459:
	;
	v4139 = v4100 + int64(1)
	v4140 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+2148)))
	v4141 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+2152)))
	v4143 = base.I64_div_u_s(int64(4294967296), v4123)
	v4145 = v4140 + v4141*v4143
	if base.B2i32(v4139 != v4145)&base.B2i32(v4100 != v4145) == int32(0) {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v4152 = v4065 + int32(1)
	v4153 = *(*int32)(unsafe.Add(mBase, uint32(v3816)+4))
	if v4153 <= v4152 {
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
	v4065 = v4152
	v4100 = v4145
	goto L457
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4191 = m.ExcPending
	if v4191 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L465
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+208)) = v4169
	F_errmsg(m, int32(_a_F_perform_base_backup_38), v58+int32(208))
	mBase = m.M
	v4210 = m.ExcPending
	if v4210 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L466
	}
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(511), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v4228 = m.ExcPending
	if v4228 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
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
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4308 = m.ExcPending
	if v4308 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L470
	}
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+192)) = v4287
	F_errmsg(m, int32(_a_F_perform_base_backup_38), v58+int32(192))
	mBase = m.M
	v4327 = m.ExcPending
	if v4327 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L471
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(520), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v4345 = m.ExcPending
	if v4345 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L472
	}
L472:
	;
	goto L1
L473:
	;
	v4368 = int32(0)
	goto L474
L474:
	;
	v4392 = *(*int32)(unsafe.Add(mBase, uint32(v3816)+12))
	v4396 = *(*int32)(unsafe.Add(mBase, uint32(v4392+v4368<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+176)) = v4396
	v4412 = v58 + int32(768)
	v4417 = F_pg_snprintf(m, v4412, int32(1024), int32(_a_F_perform_base_backup_39), v58+int32(176))
	mBase = m.M
	v4418 = m.ExcPending
	if v4418 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L476
	}
L475:
	;
	goto L443
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v4433 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+168)) = v58 + int32(2156)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+164)) = v58 + int32(2160)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+160)) = v58 + int32(540)
	v4446 = F_sscanf(m, v4396, int32(_a_F_perform_base_backup_33), v58+int32(160))
	mBase = m.M
	v4447 = m.ExcPending
	if v4447 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L477
	}
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v4462 = base.I64_div_u_s(int64(4294967296), v4433)
	v4463 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+2160)))
	v4464 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+2156)))
	v4466 = F_OpenTransientFile(m, v4412, int32(0))
	mBase = m.M
	v4467 = m.ExcPending
	if v4467 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L478
	}
L478:
	;
	v4469 = v4462*v4463 + v4464
	if v4466 < int32(0) {
		goto L479
	} else {
		goto L480
	}
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v4486 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v4500 = *(*int32)(unsafe.Add(mBase, uint32(v58)+540))
	F_CheckXLogRemoved(m, v4469, v4500)
	mBase = m.M
	v4502 = m.ExcPending
	if v4502 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L482
	}
L480:
	;
	goto L481
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	if v4466 < int32(0) {
		goto L488
	} else {
		goto L489
	}
L482:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[18])) = v4486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4521 = m.ExcPending
	if v4521 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L483
	}
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errcode_for_file_access(m)
	mBase = m.M
	v4536 = m.ExcPending
	if v4536 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L484
	}
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v4412
	F_errmsg(m, int32(_a_F_perform_base_backup_40), v58)
	mBase = m.M
	v4553 = m.ExcPending
	if v4553 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L485
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(549), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v4571 = m.ExcPending
	if v4571 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L486
	}
L486:
	;
	goto L1
L487:
	;
	if v4594 != 0 {
		goto L491
	} else {
		goto L492
	}
L488:
	;
	v4590 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v4594 = v4590
	goto L487
L489:
	;
	goto L490
L490:
	;
	v4593 = F___fstatat(m, v4466, int32(_a_F_perform_base_backup_41), v58+int32(672), int32(_a_F_perform_base_backup_42))
	mBase = m.M
	v4594 = v4593
	goto L487
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4611 = m.ExcPending
	if v4611 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L494
	}
L492:
	;
	goto L493
L493:
	;
	v4666 = *(*int64)(unsafe.Add(mBase, uint32(v58)+696))
	v4668 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	if v4666 != v4668 {
		goto L498
	} else {
		goto L499
	}
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errcode_for_file_access(m)
	mBase = m.M
	v4626 = m.ExcPending
	if v4626 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L495
	}
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+144)) = v58 + int32(768)
	F_errmsg(m, int32(_a_F_perform_base_backup_30), v58+int32(144))
	mBase = m.M
	v4647 = m.ExcPending
	if v4647 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L496
	}
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(556), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v4665 = m.ExcPending
	if v4665 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L497
	}
L497:
	;
	goto L1
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v4683 = *(*int32)(unsafe.Add(mBase, uint32(v58)+540))
	F_CheckXLogRemoved(m, v4469, v4683)
	mBase = m.M
	v4685 = m.ExcPending
	if v4685 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L501
	}
L499:
	;
	goto L500
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v4770 = int32(0)
	F__tarWriteHeader(m, v56, v58+int32(768), v4770, v58+int32(672), v4770)
	mBase = m.M
	v4775 = m.ExcPending
	if v4775 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L506
	}
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4702 = m.ExcPending
	if v4702 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L502
	}
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errcode_for_file_access(m)
	mBase = m.M
	v4717 = m.ExcPending
	if v4717 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L503
	}
L503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+128)) = v4396
	F_errmsg(m, int32(_a_F_perform_base_backup_43), v58+int32(128))
	mBase = m.M
	v4736 = m.ExcPending
	if v4736 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L504
	}
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(562), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v4754 = m.ExcPending
	if v4754 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L505
	}
L505:
	;
	goto L1
L506:
	;
	v4777 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	v4783 = v4777
	v4818 = int64(0)
	goto L508
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v5067 = F_CloseTransientFile(m, v4466)
	mBase = m.M
	v5068 = m.ExcPending
	if v5068 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L532
	}
L508:
	;
	v4822 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v4823 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v56)+8)))
	v4825 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v4825))) = int32(167772163)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v4842 = base.I64_extend_i32_s(v4783) - v4818
	if v4842 < v4823 {
		goto L510
	} else {
		goto L511
	}
L509:
	;
	v4964 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	if v4818 == v4964 {
		goto L507
	} else {
		goto L526
	}
L510:
	;
	v4844 = v4842
	goto L512
L511:
	;
	v4844 = v4823
	goto L512
L512:
	;
	v4846 = F_pread(m, v4466, v4822, base.I32_wrap_i64(v4844), v4818)
	mBase = m.M
	v4848 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[19]))
	v4849 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4848))) = v4849
	if v4846 < v4849 {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4869 = m.ExcPending
	if v4869 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L516
	}
L514:
	;
	goto L515
L515:
	;
	if v4846 != 0 {
		goto L520
	} else {
		goto L521
	}
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errcode_for_file_access(m)
	mBase = m.M
	v4884 = m.ExcPending
	if v4884 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L517
	}
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v58 + int32(768)
	F_errmsg(m, int32(_a_F_perform_base_backup_44), v58+int32(16))
	mBase = m.M
	v4905 = m.ExcPending
	if v4905 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L518
	}
L518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(2128), int32(_a_F_perform_base_backup_45))
	mBase = m.M
	v4923 = m.ExcPending
	if v4923 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L519
	}
L519:
	;
	goto L1
L520:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v4937 = *(*int32)(unsafe.Add(mBase, uint32(v58)+540))
	F_CheckXLogRemoved(m, v4469, v4937)
	mBase = m.M
	v4939 = m.ExcPending
	if v4939 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
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
	v4940 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v4941 = *(*int32)(unsafe.Add(mBase, uint32(v4940)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	m.T0[v4941].(func(*base.Module, int32, int32))(m, v56, v4846)
	mBase = m.M
	v4956 = m.ExcPending
	if v4956 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L524
	}
L524:
	;
	v4958 = v4818 + base.I64_extend_i32_u(v4846)
	v4960 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	if v4958 != base.I64_extend_i32_s(v4960) {
		v4783 = v4960
		v4818 = v4958
		goto L508
	} else {
		goto L525
	}
L525:
	;
	goto L507
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v4979 = *(*int32)(unsafe.Add(mBase, uint32(v58)+540))
	F_CheckXLogRemoved(m, v4469, v4979)
	mBase = m.M
	v4981 = m.ExcPending
	if v4981 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L527
	}
L527:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4998 = m.ExcPending
	if v4998 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L528
	}
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errcode_for_file_access(m)
	mBase = m.M
	v5013 = m.ExcPending
	if v5013 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L529
	}
L529:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+112)) = v4396
	F_errmsg(m, int32(_a_F_perform_base_backup_43), v58+int32(112))
	mBase = m.M
	v5032 = m.ExcPending
	if v5032 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L530
	}
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(587), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v5050 = m.ExcPending
	if v5050 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L531
	}
L531:
	;
	goto L1
L532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+100)) = int32(_a_F_perform_base_backup_46)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+96)) = v4396
	v5086 = v58 + int32(768)
	v5091 = F_pg_snprintf(m, v5086, int32(1024), int32(_a_F_perform_base_backup_47), v58+int32(96))
	mBase = m.M
	v5092 = m.ExcPending
	if v5092 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L533
	}
L533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_sendFileWithContent(m, v56, v5086, int32(_a_F_perform_base_backup_41), v58+int32(2056))
	mBase = m.M
	v5110 = m.ExcPending
	if v5110 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L534
	}
L534:
	;
	v5112 = v4368 + int32(1)
	v5113 = *(*int32)(unsafe.Add(mBase, uint32(v3816)+4))
	if v5112 < v5113 {
		v4368 = v5112
		goto L474
	} else {
		goto L535
	}
L535:
	;
	goto L475
L536:
	;
	v5435 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v5437 = int32(1024)
	base.MemoryFill(m, v5435, int32(0), v5437)
	v5439 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v5440 = *(*int32)(unsafe.Add(mBase, uint32(v5439)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	m.T0[v5440].(func(*base.Module, int32, int32))(m, v56, v5437)
	mBase = m.M
	v5456 = m.ExcPending
	if v5456 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L554
	}
L537:
	;
	v5160 = int32(0)
	v5161 = *(*int32)(unsafe.Add(mBase, uint32(v3817)+4))
	if v5161 <= v5160 {
		goto L536
	} else {
		goto L538
	}
L538:
	;
	v5168 = v5160
	goto L539
L539:
	;
	v5207 = *(*int32)(unsafe.Add(mBase, uint32(v3817)+12))
	v5211 = *(*int32)(unsafe.Add(mBase, uint32(v5207+v5168<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+80)) = v5211
	v5227 = v58 + int32(768)
	v5232 = F_pg_snprintf(m, v5227, int32(1024), int32(_a_F_perform_base_backup_39), v58+int32(80))
	mBase = m.M
	v5233 = m.ExcPending
	if v5233 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L541
	}
L540:
	;
	goto L536
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v5251 = F___fstatat(m, int32(-100), v5227, v58+int32(672), int32(256))
	mBase = m.M
	goto L542
L542:
	;
	if v5251 != 0 {
		goto L543
	} else {
		goto L544
	}
L543:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5268 = m.ExcPending
	if v5268 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L546
	}
L544:
	;
	goto L545
L545:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v5335 = v58 + int32(768)
	v5338 = int32(0)
	v5344 = v58 + int32(2056)
	v5348 = F_sendFile(m, v56, v5335, v5335, v58+int32(672), v5338, v5338, v5338, v5338, v5338, v5344, v5338, v5338, v5338)
	mBase = m.M
	v5349 = m.ExcPending
	if v5349 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L550
	}
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errcode_for_file_access(m)
	mBase = m.M
	v5283 = m.ExcPending
	if v5283 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L547
	}
L547:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+64)) = v5227
	F_errmsg(m, int32(_a_F_perform_base_backup_30), v58-int32(-64))
	mBase = m.M
	v5302 = m.ExcPending
	if v5302 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L548
	}
L548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(626), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v5320 = m.ExcPending
	if v5320 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L549
	}
L549:
	;
	goto L1
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v58)+52)) = int32(_a_F_perform_base_backup_46)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+48)) = v5211
	v5370 = F_pg_snprintf(m, v5335, int32(1024), int32(_a_F_perform_base_backup_47), v58+int32(48))
	mBase = m.M
	v5371 = m.ExcPending
	if v5371 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L551
	}
L551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_sendFileWithContent(m, v56, v5335, int32(_a_F_perform_base_backup_41), v5344)
	mBase = m.M
	v5387 = m.ExcPending
	if v5387 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L552
	}
L552:
	;
	v5389 = v5168 + int32(1)
	v5390 = *(*int32)(unsafe.Add(mBase, uint32(v3817)+4))
	if v5389 < v5390 {
		v5168 = v5389
		goto L539
	} else {
		goto L553
	}
L553:
	;
	goto L540
L554:
	;
	v5457 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v5458 = *(*int32)(unsafe.Add(mBase, uint32(v5457)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3813
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3812
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3814
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	m.T0[v5458].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v5473 = m.ExcPending
	if v5473 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L555
	}
L555:
	;
	v5486 = v3812
	v5487 = v3813
	v5488 = v3814
	goto L345
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5487
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5488
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v5812 = m.G0
	v5814 = v5812 - int32(128)
	m.G0 = v5814
	v5817 = v58 + int32(2056)
	v5818 = *(*int32)(unsafe.Add(mBase, uint32(v5817)))
	if v5818 != 0 {
		goto L600
	} else {
		goto L601
	}
L557:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5783 = m.ExcPending
	if v5783 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L594
	}
L558:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5766 = m.ExcPending
	if v5766 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L591
	}
L559:
	;
	F_AppendStringToManifest(m, v5538, int32(_a_F_perform_base_backup_48))
	mBase = m.M
	v5542 = m.ExcPending
	if v5542 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L562
	}
L560:
	;
	goto L561
L561:
	;
	m.G0 = v5535 + int32(80)
	goto L556
L562:
	;
	v5543 = F_readTimeLineHistory(m, v3108)
	mBase = m.M
	v5544 = m.ExcPending
	if v5544 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L563
	}
L563:
	;
	F_AppendStringToManifest(m, v5538, int32(_a_F_perform_base_backup_49))
	mBase = m.M
	v5547 = m.ExcPending
	if v5547 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L564
	}
L564:
	;
	if v5543 == int32(0) {
		goto L566
	} else {
		goto L567
	}
L565:
	;
	F_AppendStringToManifest(m, v5538, int32(_a_F_perform_base_backup_48))
	mBase = m.M
	v5716 = m.ExcPending
	if v5716 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L590
	}
L566:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5703 = m.ExcPending
	if v5703 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L587
	}
L567:
	;
	v5551 = *(*int32)(unsafe.Add(mBase, uint32(v5543)+4))
	if v5551 <= int32(0) {
		goto L566
	} else {
		goto L568
	}
L568:
	;
	v5576 = int32(1)
	v5577 = int32(0)
	v5596 = v3109
	goto L569
L569:
	;
	v5601 = *(*int32)(unsafe.Add(mBase, uint32(v5543)+12))
	v5605 = *(*int32)(unsafe.Add(mBase, uint32(v5601+v5577<<(uint(int32(2))%32))))
	v5606 = *(*int64)(unsafe.Add(mBase, uint32(v5605)+16))
	if base.B2i32(v5606 != int64(0))&base.B2i32(base.Ui64(v5606) < base.Ui64(v5518)) == int32(0) {
		goto L571
	} else {
		goto L572
	}
L570:
	;
	goto L566
L571:
	;
	v5613 = *(*int32)(unsafe.Add(mBase, uint32(v5605)))
	if base.B2i32(v3108 != v5613)&v5576 != 0 {
		goto L558
	} else {
		goto L574
	}
L572:
	;
	v5649 = v5576
	v5651 = v5596
	goto L573
L573:
	;
	v5654 = v5577 + int32(1)
	v5655 = *(*int32)(unsafe.Add(mBase, uint32(v5543)+4))
	if v5654 < v5655 {
		v5576 = v5649
		v5577 = v5654
		v5596 = v5651
		goto L569
	} else {
		goto L586
	}
L574:
	;
	if v5517 != v5613 {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	v5617 = *(*int64)(unsafe.Add(mBase, uint32(v5605)+8))
	if v5617 == int64(0) {
		goto L557
	} else {
		goto L578
	}
L576:
	;
	v5620 = v5518
	goto L577
L577:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v5535+int32(52)))) = uint32(v5596)
	v5622 = int64(32)
	v5623 = int64(base.Ui64(v5596) >> (uint(v5622) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v5535+int32(48)))) = uint32(v5623)
	*(*int32)(unsafe.Add(mBase, uint32(v5535)+36)) = v5613
	*(*uint32)(unsafe.Add(mBase, uint32(v5535)+44)) = uint32(v5620)
	v5628 = int64(base.Ui64(v5620) >> (uint(v5622) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v5535)+40)) = uint32(v5628)
	if v5576&int32(1) != 0 {
		goto L579
	} else {
		goto L580
	}
L578:
	;
	v5620 = v5617
	goto L577
L579:
	;
	v5634 = int32(_a_F_perform_base_backup_41)
	goto L581
L580:
	;
	v5634 = int32(_a_F_perform_base_backup_50)
	goto L581
L581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5535)+32)) = v5634
	v5639 = F_psprintf(m, int32(_a_F_perform_base_backup_51), v5535+int32(32))
	mBase = m.M
	v5640 = m.ExcPending
	if v5640 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L582
	}
L582:
	;
	F_AppendStringToManifest(m, v5538, v5639)
	mBase = m.M
	v5642 = m.ExcPending
	if v5642 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L583
	}
L583:
	;
	F_pfree(m, v5639)
	mBase = m.M
	v5644 = m.ExcPending
	if v5644 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L584
	}
L584:
	;
	v5645 = *(*int32)(unsafe.Add(mBase, uint32(v5605)))
	if v5517 == v5645 {
		goto L565
	} else {
		goto L585
	}
L585:
	;
	v5648 = *(*int64)(unsafe.Add(mBase, uint32(v5605)+8))
	v5649 = int32(0)
	v5651 = v5648
	goto L573
L586:
	;
	goto L570
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5535)+4)) = v3108
	*(*int32)(unsafe.Add(mBase, uint32(v5535))) = v5517
	F_errmsg(m, int32(_a_F_perform_base_backup_52), v5535)
	mBase = m.M
	v5708 = m.ExcPending
	if v5708 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L588
	}
L588:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(307), int32(_a_F_perform_base_backup_53))
	mBase = m.M
	v5713 = m.ExcPending
	if v5713 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
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
	v5767 = *(*int32)(unsafe.Add(mBase, uint32(v5605)))
	*(*int32)(unsafe.Add(mBase, uint32(v5535)+20)) = v5767
	*(*int32)(unsafe.Add(mBase, uint32(v5535)+16)) = v3108
	F_errmsg(m, int32(_a_F_perform_base_backup_54), v5535+int32(16))
	mBase = m.M
	v5774 = m.ExcPending
	if v5774 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L592
	}
L592:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(256), int32(_a_F_perform_base_backup_53))
	mBase = m.M
	v5779 = m.ExcPending
	if v5779 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
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
	v5784 = *(*int32)(unsafe.Add(mBase, uint32(v5605)))
	*(*int32)(unsafe.Add(mBase, uint32(v5535)+68)) = v5784
	*(*int32)(unsafe.Add(mBase, uint32(v5535)+64)) = v5517
	F_errmsg(m, int32(_a_F_perform_base_backup_55), v5535-int32(-64))
	mBase = m.M
	v5791 = m.ExcPending
	if v5791 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L595
	}
L595:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(280), int32(_a_F_perform_base_backup_53))
	mBase = m.M
	v5796 = m.ExcPending
	if v5796 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
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
	v6077 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v6078 = *(*int32)(unsafe.Add(mBase, uint32(v6077)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5487
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5488
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	m.T0[v6078].(func(*base.Module, int32, int64, int32))(m, v56, v3109, v3108)
	mBase = m.M
	v6093 = m.ExcPending
	if v6093 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L647
	}
L598:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6065 = m.ExcPending
	if v6065 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L643
	}
L599:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6036 = m.ExcPending
	if v6036 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L630
	}
L600:
	;
	v5819 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5817)+26)) = uint8(v5819)
	v5821 = *(*int32)(unsafe.Add(mBase, uint32(v5817)+8))
	v5823 = v5814 + int32(96)
	v5825 = F_pg_cryptohash_final(m, v5821, v5823, int32(32))
	mBase = m.M
	if v5825 < v5819 {
		goto L599
	} else {
		goto L603
	}
L601:
	;
	goto L602
L602:
	;
	m.G0 = v5814 + int32(128)
	goto L597
L603:
	;
	F_AppendStringToManifest(m, v5817, int32(_a_F_perform_base_backup_56))
	mBase = m.M
	v5830 = m.ExcPending
	if v5830 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L604
	}
L604:
	;
	v5833 = v5814 + int32(16)
	goto L606
L605:
	;
	v5857 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5814)+80)) = uint8(v5857)
	F_AppendStringToManifest(m, v5817, v5833)
	mBase = m.M
	v5860 = m.ExcPending
	if v5860 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L612
	}
L606:
	;
	v5836 = v5823
	v5838 = v5833
	goto L609
L608:
	;
	goto L605
L609:
	;
	v5840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5836))))
	v5841 = int32(1)
	v5843 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5840<<(uint(v5841)%32))+uint32(_c_F_perform_base_backup[20]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5838))) = uint16(v5843)
	v5848 = v5836 + v5841
	if base.Ui32(v5848) < base.Ui32(v5814+int32(128)) {
		v5836 = v5848
		v5838 = v5838 + int32(2)
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
	F_AppendStringToManifest(m, v5817, int32(_a_F_perform_base_backup_57))
	mBase = m.M
	v5863 = m.ExcPending
	if v5863 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L613
	}
L613:
	;
	v5864 = *(*int32)(unsafe.Add(mBase, uint32(v5817)))
	v5865 = int32(0)
	v5868 = F_BufFileSeek(m, v5864, v5865, int64(0), v5865)
	mBase = m.M
	v5869 = m.ExcPending
	if v5869 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L614
	}
L614:
	;
	if v5868 != 0 {
		goto L598
	} else {
		goto L615
	}
L615:
	;
	v5870 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v5871 = *(*int32)(unsafe.Add(mBase, uint32(v5870)+16))
	m.T0[v5871].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v5873 = m.ExcPending
	if v5873 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L616
	}
L616:
	;
	v5874 = *(*int64)(unsafe.Add(mBase, uint32(v5817)+16))
	if v5874 != int64(0) {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	v5895 = int32(0)
	v5916 = v5874
	v5917 = int64(0)
	goto L620
L618:
	;
	goto L619
L619:
	;
	v5980 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v5981 = *(*int32)(unsafe.Add(mBase, uint32(v5980)+24))
	m.T0[v5981].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v5983 = m.ExcPending
	if v5983 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L628
	}
L620:
	;
	v5920 = *(*int32)(unsafe.Add(mBase, uint32(v5817)))
	v5921 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v5922 = v5916 - v5917
	v5923 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v56)+8)))
	if base.Ui64(v5922) < base.Ui64(v5923) {
		goto L622
	} else {
		goto L623
	}
L621:
	;
	goto L619
L622:
	;
	v5925 = v5922
	goto L624
L623:
	;
	v5925 = v5923
	goto L624
L624:
	;
	v5926 = base.I32_wrap_i64(v5925)
	F_BufFileReadExact(m, v5920, v5921, v5926)
	mBase = m.M
	v5928 = m.ExcPending
	if v5928 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L625
	}
L625:
	;
	v5929 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v5930 = *(*int32)(unsafe.Add(mBase, uint32(v5929)+20))
	m.T0[v5930].(func(*base.Module, int32, int32))(m, v56, v5926)
	mBase = m.M
	v5932 = m.ExcPending
	if v5932 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L626
	}
L626:
	;
	v5933 = *(*int64)(unsafe.Add(mBase, uint32(v5817)+16))
	v5934 = v5926 + v5895
	v5935 = base.I64_extend_i32_u(v5934)
	if base.Ui64(v5935) < base.Ui64(v5933) {
		v5895 = v5934
		v5916 = v5933
		v5917 = v5935
		goto L620
	} else {
		goto L627
	}
L627:
	;
	goto L621
L628:
	;
	v5984 = *(*int32)(unsafe.Add(mBase, uint32(v5817)))
	F_BufFileClose(m, v5984)
	mBase = m.M
	v5986 = m.ExcPending
	if v5986 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L629
	}
L629:
	;
	goto L602
L630:
	;
	v6037 = *(*int32)(unsafe.Add(mBase, uint32(v5817)+8))
	if v6037 == int32(0) {
		goto L632
	} else {
		goto L633
	}
L631:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5814))) = v6052
	F_errmsg_internal(m, int32(_a_F_perform_base_backup_58), v5814)
	mBase = m.M
	v6056 = m.ExcPending
	if v6056 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L641
	}
L632:
	;
	v6052 = int32(_a_F_perform_base_backup_2)
	goto L631
L633:
	;
	goto L634
L634:
	;
	v6044 = *(*int32)(unsafe.Add(mBase, uint32(v6037)+4))
	if v6044 == int32(1) {
		goto L635
	} else {
		goto L636
	}
L635:
	;
	v6047 = int32(_a_F_perform_base_backup_3)
	goto L637
L636:
	;
	v6047 = int32(_a_F_perform_base_backup_4)
	goto L637
L637:
	;
	if v6044 == int32(2) {
		goto L638
	} else {
		goto L639
	}
L638:
	;
	v6050 = int32(_a_F_perform_base_backup_2)
	goto L640
L639:
	;
	v6050 = v6047
	goto L640
L640:
	;
	v6052 = v6050
	goto L631
L641:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(341), int32(_a_F_perform_base_backup_59))
	mBase = m.M
	v6061 = m.ExcPending
	if v6061 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
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
	v6067 = m.ExcPending
	if v6067 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L644
	}
L644:
	;
	F_errmsg(m, int32(_a_F_perform_base_backup_60), int32(0))
	mBase = m.M
	v6071 = m.ExcPending
	if v6071 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L645
	}
L645:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(357), int32(_a_F_perform_base_backup_59))
	mBase = m.M
	v6076 = m.ExcPending
	if v6076 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
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
	v6095 = *(*int64)(unsafe.Add(mBase, _c_F_perform_base_backup[5]))
	if v6095 != int64(0) {
		goto L648
	} else {
		goto L649
	}
L648:
	;
	if v6095 < int64(2) {
		goto L651
	} else {
		goto L652
	}
L649:
	;
	goto L650
L650:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5487
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5488
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v6243 = v58 + int32(2056)
	v6244 = *(*int32)(unsafe.Add(mBase, uint32(v6243)+8))
	F_pg_cryptohash_free(m, v6244)
	mBase = m.M
	v6246 = m.ExcPending
	if v6246 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L661
	}
L651:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5487
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5488
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6177 = m.ExcPending
	if v6177 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L657
	}
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5487
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5488
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v6115 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v6116 = m.ExcPending
	if v6116 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L653
	}
L653:
	;
	if v6115 == int32(0) {
		goto L651
	} else {
		goto L654
	}
L654:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5487
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5488
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v6133 = *(*int64)(unsafe.Add(mBase, _c_F_perform_base_backup[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+32)) = v6133
	F_errmsg_plural(m, int32(_a_F_perform_base_backup_61), int32(_a_F_perform_base_backup_62), base.I32_wrap_i64(v6133), v58+int32(32))
	mBase = m.M
	v6141 = m.ExcPending
	if v6141 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L655
	}
L655:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5487
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5488
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(661), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v6159 = m.ExcPending
	if v6159 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L656
	}
L656:
	;
	goto L651
L657:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5487
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5488
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errcode(m, int32(16779816))
	mBase = m.M
	v6193 = m.ExcPending
	if v6193 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L658
	}
L658:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5487
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5488
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errmsg(m, int32(_a_F_perform_base_backup_63), int32(0))
	mBase = m.M
	v6210 = m.ExcPending
	if v6210 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L659
	}
L659:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5487
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5488
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_errfinish(m, int32(_a_F_perform_base_backup_31), int32(665), int32(_a_F_perform_base_backup_32))
	mBase = m.M
	v6228 = m.ExcPending
	if v6228 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L660
	}
L660:
	;
	goto L1
L661:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6243)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5487
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5488
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	F_ReleaseAuxProcessResources(m, int32(1))
	mBase = m.M
	v6264 = m.ExcPending
	if v6264 != 0 {
		v6324 = v55
		v6325 = v56
		v6326 = v57
		v6327 = v58
		v6351 = v82
		v6355 = v86
		v6356 = v87
		goto L6
	} else {
		goto L662
	}
L662:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5487
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5486
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5488
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v392
	v6280 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[7]))
	if v6280 == int32(0) {
		goto L664
	} else {
		goto L665
	}
L663:
	;
	goto L5
L664:
	;
	goto L663
L665:
	;
	v6284 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[8])))
	if v6284&int32(1) == int32(0) {
		goto L664
	} else {
		goto L666
	}
L666:
	;
	v6289 = *(*int32)(unsafe.Add(mBase, uint32(v6280)+220))
	if v6289 == int32(0) {
		goto L664
	} else {
		goto L667
	}
L667:
	;
	v6292 = int32(_a_F_perform_base_backup_7)
	v6294 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	v6295 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v6294 + v6295
	v6298 = *(*int32)(unsafe.Add(mBase, uint32(v6280)))
	*(*int32)(unsafe.Add(mBase, uint32(v6280))) = v6298 + v6295
	v6302 = int32(0)
	v6304 = int32(_a_F_perform_base_backup_8)
	v6305 = base.AtomicRmwOr32(m, v6302, v6304, v6302)
	*(*int32)(unsafe.Add(mBase, uint32(v6280)+220)) = v6302
	*(*int32)(unsafe.Add(mBase, uint32(v6280)+224)) = v6302
	v6313 = base.AtomicRmwOr32(m, v6302, v6304, v6302)
	v6314 = *(*int32)(unsafe.Add(mBase, uint32(v6280)))
	*(*int32)(unsafe.Add(mBase, uint32(v6280))) = v6314 + v6295
	v6320 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v6320 - v6295
	goto L664
L668:
	;
	v6372 = int32(v6368)
	m.G0 = v6327
	v6374 = *(*int32)(unsafe.Add(mBase, uint32(v6372)+4))
	v6375 = *(*int32)(unsafe.Add(mBase, uint32(v6372)))
	v6378 = *(*int32)(unsafe.Add(mBase, uint32(v6375)))
	if v6327+int32(332) == v6378 {
		goto L671
	} else {
		goto L672
	}
L669:
	;
	m.ExcPending = 1
	goto L677
L670:
	;
	if v6382 != 0 {
		goto L674
	} else {
		goto L675
	}
L671:
	;
	v6380 = *(*int32)(unsafe.Add(mBase, uint32(v6375)+4))
	v6382 = v6380
	goto L673
L672:
	;
	v6382 = int32(0)
	goto L673
L673:
	;
	goto L670
L674:
	;
	v6383 = *(*int32)(unsafe.Add(mBase, uint32(v6327)+2220))
	v6384 = *(*int32)(unsafe.Add(mBase, uint32(v6327)+2216))
	v6385 = *(*int32)(unsafe.Add(mBase, uint32(v6327)+2212))
	v6386 = *(*int32)(unsafe.Add(mBase, uint32(v6327)+2208))
	v6387 = *(*int32)(unsafe.Add(mBase, uint32(v6327)+2204))
	v6388 = *(*int32)(unsafe.Add(mBase, uint32(v6327)+2200))
	v6389 = *(*int32)(unsafe.Add(mBase, uint32(v6327)+2196))
	v6390 = *(*int32)(unsafe.Add(mBase, uint32(v6327)+2192))
	v6391 = *(*int64)(unsafe.Add(mBase, uint32(v6327)+2184))
	v6392 = *(*int64)(unsafe.Add(mBase, uint32(v6327)+2176))
	v6393 = *(*int32)(unsafe.Add(mBase, uint32(v6327)+2172))
	v6394 = *(*int32)(unsafe.Add(mBase, uint32(v6327)+2168))
	v6395 = *(*int32)(unsafe.Add(mBase, uint32(v6327)+2164))
	v55 = v6324
	v56 = v6325
	v57 = v6326
	v58 = v6327
	v59 = v6374
	v60 = v6388
	v61 = v6385
	v62 = v6383
	v63 = v6386
	v64 = v6387
	v65 = v6390
	v66 = v6389
	v67 = v6395
	v68 = v6394
	v69 = v6393
	v70 = v6384
	v73 = v6382
	v82 = v6351
	v86 = v6355
	v87 = v6356
	v91 = v6392
	v92 = v6391
	goto L2
L675:
	;
	goto L676
L676:
	;
	F___wasm_longjmp(m, v6375, v6374)
	mBase = m.M
	v6397 = m.ExcPending
	if v6397 != 0 {
		goto L677
	} else {
		goto L678
	}
L677:
	;
	return
L678:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
