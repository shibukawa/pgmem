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
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v347 int32
	_ = v347
	var v348 int64
	_ = v348
	var v350 int32
	_ = v350
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
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
	var v411 int32
	_ = v411
	var v412 int64
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v484 int64
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v693 int32
	_ = v693
	var v718 int32
	_ = v718
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v765 int32
	_ = v765
	var v768 int64
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int64
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v796 int32
	_ = v796
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int64
	_ = v827
	var v829 int32
	_ = v829
	var v830 int64
	_ = v830
	var v831 int32
	_ = v831
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int64
	_ = v842
	var v843 int32
	_ = v843
	var v844 int64
	_ = v844
	var v847 int64
	_ = v847
	var v848 int64
	_ = v848
	var v852 int64
	_ = v852
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v865 int64
	_ = v865
	var v867 int64
	_ = v867
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int64
	_ = v877
	var v880 int64
	_ = v880
	var v884 int64
	_ = v884
	var v885 int64
	_ = v885
	var v888 int64
	_ = v888
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v904 int64
	_ = v904
	var v907 int32
	_ = v907
	var v926 int32
	_ = v926
	var v950 int64
	_ = v950
	var v952 int64
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v957 int64
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
	var v968 int32
	_ = v968
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1018 int64
	_ = v1018
	var v1019 int64
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1025 int64
	_ = v1025
	var v1029 int64
	_ = v1029
	var v1031 int64
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1037 int32
	_ = v1037
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1092 int64
	_ = v1092
	var v1096 int64
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1129 int32
	_ = v1129
	var v1191 int32
	_ = v1191
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1232 int32
	_ = v1232
	var v1251 int64
	_ = v1251
	var v1259 int32
	_ = v1259
	var v1260 int64
	_ = v1260
	var v1262 int64
	_ = v1262
	var v1266 int64
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1308 int64
	_ = v1308
	var v1358 int32
	_ = v1358
	var v1361 int64
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1369 int64
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1389 int32
	_ = v1389
	var v1390 int64
	_ = v1390
	var v1393 int64
	_ = v1393
	var v1399 int32
	_ = v1399
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1430 int32
	_ = v1430
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1478 int32
	_ = v1478
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1527 int32
	_ = v1527
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1579 int32
	_ = v1579
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1647 int32
	_ = v1647
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1653 int32
	_ = v1653
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1677 int32
	_ = v1677
	var v1719 int32
	_ = v1719
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1924 int32
	_ = v1924
	var v1929 int32
	_ = v1929
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int64
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1943 int64
	_ = v1943
	var v1946 int64
	_ = v1946
	var v1947 int64
	_ = v1947
	var v1951 int64
	_ = v1951
	var v1957 int32
	_ = v1957
	var v1962 int32
	_ = v1962
	var v1966 int32
	_ = v1966
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1971 int64
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1973 int64
	_ = v1973
	var v1976 int64
	_ = v1976
	var v1977 int64
	_ = v1977
	var v1981 int64
	_ = v1981
	var v1987 int32
	_ = v1987
	var v1992 int32
	_ = v1992
	var v1997 int32
	_ = v1997
	var v2000 int32
	_ = v2000
	var v2004 int32
	_ = v2004
	var v2009 int32
	_ = v2009
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2108 int32
	_ = v2108
	var v2112 int32
	_ = v2112
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2123 int32
	_ = v2123
	var v2131 int32
	_ = v2131
	var v2137 int32
	_ = v2137
	var v2141 int32
	_ = v2141
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2152 int32
	_ = v2152
	var v2184 int64
	_ = v2184
	var v2185 int64
	_ = v2185
	var v2191 int32
	_ = v2191
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2217 int32
	_ = v2217
	var v2220 int64
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2237 int32
	_ = v2237
	var v2239 int64
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2242 int64
	_ = v2242
	var v2243 int64
	_ = v2243
	var v2244 int64
	_ = v2244
	var v2246 int64
	_ = v2246
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2289 int64
	_ = v2289
	var v2290 int64
	_ = v2290
	var v2296 int32
	_ = v2296
	var v2334 int64
	_ = v2334
	var v2335 int64
	_ = v2335
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2388 int32
	_ = v2388
	var v2413 int32
	_ = v2413
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2438 int32
	_ = v2438
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2470 int32
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2491 int32
	_ = v2491
	var v2507 int32
	_ = v2507
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2531 int64
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2551 int32
	_ = v2551
	var v2568 int32
	_ = v2568
	var v2583 int32
	_ = v2583
	var v2603 int32
	_ = v2603
	var v2621 int32
	_ = v2621
	var v2635 int32
	_ = v2635
	var v2639 int32
	_ = v2639
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2706 int64
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2711 int32
	_ = v2711
	var v2714 int32
	_ = v2714
	var v2717 int32
	_ = v2717
	var v2719 int32
	_ = v2719
	var v2721 int32
	_ = v2721
	var v2722 int32
	_ = v2722
	var v2738 int32
	_ = v2738
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2755 int32
	_ = v2755
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2781 int32
	_ = v2781
	var v2795 int32
	_ = v2795
	var v2798 int32
	_ = v2798
	var v2813 int32
	_ = v2813
	var v2870 int32
	_ = v2870
	var v2872 int32
	_ = v2872
	var v2880 int32
	_ = v2880
	var v2881 int64
	_ = v2881
	var v2883 int64
	_ = v2883
	var v2897 int32
	_ = v2897
	var v2901 int32
	_ = v2901
	var v2906 int32
	_ = v2906
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2912 int32
	_ = v2912
	var v3007 int32
	_ = v3007
	var v3010 int32
	_ = v3010
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3026 int64
	_ = v3026
	var v3028 int32
	_ = v3028
	var v3031 int32
	_ = v3031
	var v3042 int32
	_ = v3042
	var v3043 int32
	_ = v3043
	var v3046 int32
	_ = v3046
	var v3048 int32
	_ = v3048
	var v3064 int32
	_ = v3064
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3085 int64
	_ = v3085
	var v3100 int32
	_ = v3100
	var v3115 int32
	_ = v3115
	var v3132 int32
	_ = v3132
	var v3137 int32
	_ = v3137
	var v3155 int32
	_ = v3155
	var v3159 int32
	_ = v3159
	var v3164 int32
	_ = v3164
	var v3166 int32
	_ = v3166
	var v3167 int32
	_ = v3167
	var v3170 int32
	_ = v3170
	var v3178 int32
	_ = v3178
	var v3184 int32
	_ = v3184
	var v3188 int32
	_ = v3188
	var v3189 int64
	_ = v3189
	var v3204 int64
	_ = v3204
	var v3206 int64
	_ = v3206
	var v3208 int64
	_ = v3208
	var v3209 int64
	_ = v3209
	var v3212 int64
	_ = v3212
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3236 int64
	_ = v3236
	var v3240 int64
	_ = v3240
	var v3242 int64
	_ = v3242
	var v3243 int64
	_ = v3243
	var v3246 int64
	_ = v3246
	var v3254 int32
	_ = v3254
	var v3255 int32
	_ = v3255
	var v3270 int32
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3290 int32
	_ = v3290
	var v3294 int32
	_ = v3294
	var v3295 int32
	_ = v3295
	var v3299 int32
	_ = v3299
	var v3300 int32
	_ = v3300
	var v3306 int32
	_ = v3306
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3316 int32
	_ = v3316
	var v3318 int32
	_ = v3318
	var v3319 int32
	_ = v3319
	var v3359 int32
	_ = v3359
	var v3360 int32
	_ = v3360
	var v3376 int32
	_ = v3376
	var v3380 int32
	_ = v3380
	var v3382 int32
	_ = v3382
	var v3383 int64
	_ = v3383
	var v3391 int32
	_ = v3391
	var v3395 int32
	_ = v3395
	var v3399 int32
	_ = v3399
	var v3405 int32
	_ = v3405
	var v3409 int32
	_ = v3409
	var v3410 int32
	_ = v3410
	var v3417 int32
	_ = v3417
	var v3418 int32
	_ = v3418
	var v3419 int32
	_ = v3419
	var v3423 int32
	_ = v3423
	var v3426 int32
	_ = v3426
	var v3430 int32
	_ = v3430
	var v3431 int32
	_ = v3431
	var v3439 int32
	_ = v3439
	var v3445 int32
	_ = v3445
	var v3447 int32
	_ = v3447
	var v3449 int32
	_ = v3449
	var v3459 int32
	_ = v3459
	var v3476 int32
	_ = v3476
	var v3479 int32
	_ = v3479
	var v3482 int32
	_ = v3482
	var v3485 int32
	_ = v3485
	var v3486 int32
	_ = v3486
	var v3489 int32
	_ = v3489
	var v3490 int32
	_ = v3490
	var v3493 int32
	_ = v3493
	var v3500 int32
	_ = v3500
	var v3501 int32
	_ = v3501
	var v3520 int32
	_ = v3520
	var v3523 int32
	_ = v3523
	var v3526 int32
	_ = v3526
	var v3527 int32
	_ = v3527
	var v3530 int32
	_ = v3530
	var v3531 int32
	_ = v3531
	var v3534 int32
	_ = v3534
	var v3541 int32
	_ = v3541
	var v3542 int32
	_ = v3542
	var v3559 int32
	_ = v3559
	var v3560 int32
	_ = v3560
	var v3574 int32
	_ = v3574
	var v3575 int32
	_ = v3575
	var v3589 int32
	_ = v3589
	var v3593 int32
	_ = v3593
	var v3595 int32
	_ = v3595
	var v3596 int64
	_ = v3596
	var v3604 int32
	_ = v3604
	var v3608 int32
	_ = v3608
	var v3612 int32
	_ = v3612
	var v3618 int32
	_ = v3618
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3630 int32
	_ = v3630
	var v3631 int32
	_ = v3631
	var v3632 int32
	_ = v3632
	var v3636 int32
	_ = v3636
	var v3639 int32
	_ = v3639
	var v3643 int32
	_ = v3643
	var v3644 int32
	_ = v3644
	var v3652 int32
	_ = v3652
	var v3658 int32
	_ = v3658
	var v3660 int32
	_ = v3660
	var v3662 int32
	_ = v3662
	var v3672 int32
	_ = v3672
	var v3689 int32
	_ = v3689
	var v3690 int32
	_ = v3690
	var v3693 int32
	_ = v3693
	var v3696 int32
	_ = v3696
	var v3699 int32
	_ = v3699
	var v3700 int32
	_ = v3700
	var v3703 int32
	_ = v3703
	var v3704 int32
	_ = v3704
	var v3707 int32
	_ = v3707
	var v3714 int32
	_ = v3714
	var v3715 int32
	_ = v3715
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3745 int32
	_ = v3745
	var v3746 int32
	_ = v3746
	var v3748 int32
	_ = v3748
	var v3749 int32
	_ = v3749
	var v3750 int32
	_ = v3750
	var v3751 int32
	_ = v3751
	var v3766 int32
	_ = v3766
	var v3767 int32
	_ = v3767
	var v3780 int32
	_ = v3780
	var v3781 int32
	_ = v3781
	var v3782 int32
	_ = v3782
	var v3784 int32
	_ = v3784
	var v3785 int32
	_ = v3785
	var v3825 int32
	_ = v3825
	var v3826 int32
	_ = v3826
	var v3841 int32
	_ = v3841
	var v3857 int32
	_ = v3857
	var v3876 int32
	_ = v3876
	var v3893 int32
	_ = v3893
	var v3911 int32
	_ = v3911
	var v3912 int32
	_ = v3912
	var v3913 int32
	_ = v3913
	var v3928 int64
	_ = v3928
	var v3941 int32
	_ = v3941
	var v3942 int32
	_ = v3942
	var v3943 int64
	_ = v3943
	var v3944 int64
	_ = v3944
	var v3946 int64
	_ = v3946
	var v3950 int32
	_ = v3950
	var v3951 int32
	_ = v3951
	var v3955 int32
	_ = v3955
	var v3970 int32
	_ = v3970
	var v3972 int32
	_ = v3972
	var v3974 int32
	_ = v3974
	var v3991 int32
	_ = v3991
	var v4010 int32
	_ = v4010
	var v4028 int32
	_ = v4028
	var v4033 int32
	_ = v4033
	var v4068 int64
	_ = v4068
	var v4072 int32
	_ = v4072
	var v4076 int32
	_ = v4076
	var v4091 int64
	_ = v4091
	var v4104 int32
	_ = v4104
	var v4105 int32
	_ = v4105
	var v4107 int64
	_ = v4107
	var v4108 int64
	_ = v4108
	var v4109 int64
	_ = v4109
	var v4111 int64
	_ = v4111
	var v4113 int64
	_ = v4113
	var v4120 int32
	_ = v4120
	var v4121 int32
	_ = v4121
	var v4137 int32
	_ = v4137
	var v4138 int32
	_ = v4138
	var v4140 int32
	_ = v4140
	var v4142 int32
	_ = v4142
	var v4159 int32
	_ = v4159
	var v4178 int32
	_ = v4178
	var v4196 int32
	_ = v4196
	var v4255 int32
	_ = v4255
	var v4257 int32
	_ = v4257
	var v4259 int32
	_ = v4259
	var v4276 int32
	_ = v4276
	var v4295 int32
	_ = v4295
	var v4313 int32
	_ = v4313
	var v4336 int32
	_ = v4336
	var v4360 int32
	_ = v4360
	var v4364 int32
	_ = v4364
	var v4380 int32
	_ = v4380
	var v4385 int32
	_ = v4385
	var v4386 int32
	_ = v4386
	var v4401 int64
	_ = v4401
	var v4414 int32
	_ = v4414
	var v4415 int32
	_ = v4415
	var v4430 int64
	_ = v4430
	var v4431 int64
	_ = v4431
	var v4432 int64
	_ = v4432
	var v4434 int32
	_ = v4434
	var v4435 int32
	_ = v4435
	var v4437 int64
	_ = v4437
	var v4454 int32
	_ = v4454
	var v4468 int32
	_ = v4468
	var v4470 int32
	_ = v4470
	var v4489 int32
	_ = v4489
	var v4504 int32
	_ = v4504
	var v4521 int32
	_ = v4521
	var v4539 int32
	_ = v4539
	var v4558 int32
	_ = v4558
	var v4561 int32
	_ = v4561
	var v4562 int32
	_ = v4562
	var v4579 int32
	_ = v4579
	var v4594 int32
	_ = v4594
	var v4615 int32
	_ = v4615
	var v4633 int32
	_ = v4633
	var v4634 int64
	_ = v4634
	var v4636 int64
	_ = v4636
	var v4651 int32
	_ = v4651
	var v4653 int32
	_ = v4653
	var v4670 int32
	_ = v4670
	var v4685 int32
	_ = v4685
	var v4704 int32
	_ = v4704
	var v4722 int32
	_ = v4722
	var v4738 int32
	_ = v4738
	var v4743 int32
	_ = v4743
	var v4745 int32
	_ = v4745
	var v4751 int32
	_ = v4751
	var v4786 int64
	_ = v4786
	var v4790 int32
	_ = v4790
	var v4791 int64
	_ = v4791
	var v4793 int32
	_ = v4793
	var v4810 int64
	_ = v4810
	var v4812 int64
	_ = v4812
	var v4814 int32
	_ = v4814
	var v4816 int32
	_ = v4816
	var v4817 int32
	_ = v4817
	var v4837 int32
	_ = v4837
	var v4852 int32
	_ = v4852
	var v4873 int32
	_ = v4873
	var v4891 int32
	_ = v4891
	var v4905 int32
	_ = v4905
	var v4907 int32
	_ = v4907
	var v4908 int32
	_ = v4908
	var v4909 int32
	_ = v4909
	var v4924 int32
	_ = v4924
	var v4926 int64
	_ = v4926
	var v4928 int32
	_ = v4928
	var v4932 int64
	_ = v4932
	var v4947 int32
	_ = v4947
	var v4949 int32
	_ = v4949
	var v4966 int32
	_ = v4966
	var v4981 int32
	_ = v4981
	var v5000 int32
	_ = v5000
	var v5018 int32
	_ = v5018
	var v5035 int32
	_ = v5035
	var v5036 int32
	_ = v5036
	var v5054 int32
	_ = v5054
	var v5059 int32
	_ = v5059
	var v5060 int32
	_ = v5060
	var v5078 int32
	_ = v5078
	var v5080 int32
	_ = v5080
	var v5081 int32
	_ = v5081
	var v5128 int32
	_ = v5128
	var v5129 int32
	_ = v5129
	var v5136 int32
	_ = v5136
	var v5175 int32
	_ = v5175
	var v5179 int32
	_ = v5179
	var v5195 int32
	_ = v5195
	var v5200 int32
	_ = v5200
	var v5201 int32
	_ = v5201
	var v5219 int32
	_ = v5219
	var v5236 int32
	_ = v5236
	var v5251 int32
	_ = v5251
	var v5270 int32
	_ = v5270
	var v5288 int32
	_ = v5288
	var v5303 int32
	_ = v5303
	var v5306 int32
	_ = v5306
	var v5312 int32
	_ = v5312
	var v5316 int32
	_ = v5316
	var v5317 int32
	_ = v5317
	var v5338 int32
	_ = v5338
	var v5339 int32
	_ = v5339
	var v5355 int32
	_ = v5355
	var v5357 int32
	_ = v5357
	var v5358 int32
	_ = v5358
	var v5403 int32
	_ = v5403
	var v5405 int32
	_ = v5405
	var v5407 int32
	_ = v5407
	var v5408 int32
	_ = v5408
	var v5424 int32
	_ = v5424
	var v5425 int32
	_ = v5425
	var v5426 int32
	_ = v5426
	var v5441 int32
	_ = v5441
	var v5454 int32
	_ = v5454
	var v5455 int32
	_ = v5455
	var v5456 int32
	_ = v5456
	var v5485 int32
	_ = v5485
	var v5486 int64
	_ = v5486
	var v5501 int32
	_ = v5501
	var v5503 int32
	_ = v5503
	var v5506 int32
	_ = v5506
	var v5507 int32
	_ = v5507
	var v5510 int32
	_ = v5510
	var v5511 int32
	_ = v5511
	var v5512 int32
	_ = v5512
	var v5515 int32
	_ = v5515
	var v5519 int32
	_ = v5519
	var v5544 int32
	_ = v5544
	var v5545 int32
	_ = v5545
	var v5564 int64
	_ = v5564
	var v5569 int32
	_ = v5569
	var v5573 int32
	_ = v5573
	var v5574 int64
	_ = v5574
	var v5581 int32
	_ = v5581
	var v5585 int64
	_ = v5585
	var v5588 int64
	_ = v5588
	var v5590 int64
	_ = v5590
	var v5591 int64
	_ = v5591
	var v5596 int64
	_ = v5596
	var v5602 int32
	_ = v5602
	var v5607 int32
	_ = v5607
	var v5608 int32
	_ = v5608
	var v5610 int32
	_ = v5610
	var v5612 int32
	_ = v5612
	var v5613 int32
	_ = v5613
	var v5616 int64
	_ = v5616
	var v5617 int32
	_ = v5617
	var v5619 int64
	_ = v5619
	var v5622 int32
	_ = v5622
	var v5623 int32
	_ = v5623
	var v5671 int32
	_ = v5671
	var v5676 int32
	_ = v5676
	var v5681 int32
	_ = v5681
	var v5684 int32
	_ = v5684
	var v5734 int32
	_ = v5734
	var v5735 int32
	_ = v5735
	var v5742 int32
	_ = v5742
	var v5747 int32
	_ = v5747
	var v5751 int32
	_ = v5751
	var v5752 int32
	_ = v5752
	var v5759 int32
	_ = v5759
	var v5764 int32
	_ = v5764
	var v5780 int32
	_ = v5780
	var v5782 int32
	_ = v5782
	var v5785 int32
	_ = v5785
	var v5786 int32
	_ = v5786
	var v5787 int32
	_ = v5787
	var v5789 int32
	_ = v5789
	var v5791 int32
	_ = v5791
	var v5793 int32
	_ = v5793
	var v5798 int32
	_ = v5798
	var v5801 int32
	_ = v5801
	var v5804 int32
	_ = v5804
	var v5806 int32
	_ = v5806
	var v5808 int32
	_ = v5808
	var v5809 int32
	_ = v5809
	var v5811 int32
	_ = v5811
	var v5816 int32
	_ = v5816
	var v5825 int32
	_ = v5825
	var v5828 int32
	_ = v5828
	var v5831 int32
	_ = v5831
	var v5832 int32
	_ = v5832
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
	var v5841 int32
	_ = v5841
	var v5842 int64
	_ = v5842
	var v5863 int32
	_ = v5863
	var v5884 int64
	_ = v5884
	var v5885 int64
	_ = v5885
	var v5888 int32
	_ = v5888
	var v5889 int32
	_ = v5889
	var v5890 int64
	_ = v5890
	var v5891 int64
	_ = v5891
	var v5893 int64
	_ = v5893
	var v5894 int32
	_ = v5894
	var v5896 int32
	_ = v5896
	var v5897 int32
	_ = v5897
	var v5898 int32
	_ = v5898
	var v5900 int32
	_ = v5900
	var v5901 int64
	_ = v5901
	var v5902 int32
	_ = v5902
	var v5903 int64
	_ = v5903
	var v5948 int32
	_ = v5948
	var v5949 int32
	_ = v5949
	var v5951 int32
	_ = v5951
	var v5952 int32
	_ = v5952
	var v5954 int32
	_ = v5954
	var v6004 int32
	_ = v6004
	var v6005 int32
	_ = v6005
	var v6012 int32
	_ = v6012
	var v6015 int32
	_ = v6015
	var v6018 int32
	_ = v6018
	var v6020 int32
	_ = v6020
	var v6024 int32
	_ = v6024
	var v6029 int32
	_ = v6029
	var v6033 int32
	_ = v6033
	var v6035 int32
	_ = v6035
	var v6039 int32
	_ = v6039
	var v6044 int32
	_ = v6044
	var v6045 int32
	_ = v6045
	var v6046 int32
	_ = v6046
	var v6061 int32
	_ = v6061
	var v6063 int64
	_ = v6063
	var v6083 int32
	_ = v6083
	var v6084 int32
	_ = v6084
	var v6101 int64
	_ = v6101
	var v6109 int32
	_ = v6109
	var v6127 int32
	_ = v6127
	var v6145 int32
	_ = v6145
	var v6161 int32
	_ = v6161
	var v6178 int32
	_ = v6178
	var v6196 int32
	_ = v6196
	var v6211 int32
	_ = v6211
	var v6212 int32
	_ = v6212
	var v6214 int32
	_ = v6214
	var v6232 int32
	_ = v6232
	var v6248 int32
	_ = v6248
	var v6252 int32
	_ = v6252
	var v6257 int32
	_ = v6257
	var v6260 int32
	_ = v6260
	var v6262 int32
	_ = v6262
	var v6263 int32
	_ = v6263
	var v6266 int32
	_ = v6266
	var v6270 int32
	_ = v6270
	var v6280 int32
	_ = v6280
	var v6284 int32
	_ = v6284
	var v6285 int32
	_ = v6285
	var v6286 int32
	_ = v6286
	var v6287 int32
	_ = v6287
	var v6311 int32
	_ = v6311
	var v6315 int32
	_ = v6315
	var v6316 int32
	_ = v6316
	var v6327 int32
	_ = v6327
	var v6328 int64
	_ = v6328
	var v6332 int32
	_ = v6332
	var v6334 int32
	_ = v6334
	var v6335 int32
	_ = v6335
	var v6338 int32
	_ = v6338
	var v6340 int32
	_ = v6340
	var v6342 int32
	_ = v6342
	var v6343 int32
	_ = v6343
	var v6344 int32
	_ = v6344
	var v6345 int32
	_ = v6345
	var v6346 int32
	_ = v6346
	var v6347 int32
	_ = v6347
	var v6348 int32
	_ = v6348
	var v6349 int32
	_ = v6349
	var v6350 int32
	_ = v6350
	var v6351 int64
	_ = v6351
	var v6352 int64
	_ = v6352
	var v6353 int32
	_ = v6353
	var v6354 int32
	_ = v6354
	var v6355 int32
	_ = v6355
	var v6357 int32
	_ = v6357
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
	v6327 = int32(m.ExcTag)
	v6328 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v6327 == int32(0) {
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
	v381 = v59
	v382 = v60
	v383 = v61
	v384 = v62
	v385 = v63
	v386 = v64
	v387 = v65
	v388 = v66
	v389 = v70
	goto L9
L9:
	;
	if v381 == int32(0) {
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
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+5)))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
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
	F_do_pg_backup_start(m, v330, v329, v58+int32(2088), v262, v277)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	*(*int64)(unsafe.Add(mBase, uint32(v296+int32(0))+232)) = int64(1)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	*(*int32)(unsafe.Add(mBase, uint32(v296))) = v319 + v308
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v325 - v308
	goto L43
L46:
	;
	v348 = *(*int64)(unsafe.Add(mBase, uint32(v262)+1032))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2120)) = v348
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v262)+1040))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2128)) = v350
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
	v368 = m.ExcPending
	if v368 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L47
	}
L47:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[10]))
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[11]))
	goto L48
L48:
	;
	v375 = v58 + int32(1888)
	*(*int32)(unsafe.Add(mBase, uint32(v375)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v375))) = v58 + int32(332)
	goto L51
L49:
	;
	v381 = int32(0)
	v382 = v86
	v383 = v262
	v384 = v82
	v385 = v277
	v386 = v87
	v387 = v371
	v388 = v373
	v389 = v123
	goto L9
L51:
	;
	goto L49
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v2870 = m.G0
	v2872 = v2870 - int32(32)
	m.G0 = v2872
	*(*int64)(unsafe.Add(mBase, uint32(v2872)+24)) = int64(17179869184)
	*(*int64)(unsafe.Add(mBase, uint32(v2872))) = int64(4)
	v2880 = *(*int32)(unsafe.Add(mBase, uint32(v58+int32(2088))))
	if v2880 != 0 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[10])) = v387
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[11])) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_cancel_before_shmem_exit(m, int32(407), int32(0))
	mBase = m.M
	v2781 = m.ExcPending
	if v2781 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v411 = int32(0)
	v412 = int64(0)
	v415 = m.G0
	v417 = v415 - int32(2336)
	m.G0 = v417
	v419 = int32(_a_F_perform_base_backup_8)
	v420 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[12]))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[12])) = v422
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v424 == v411 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v2067 = F_palloc0(m, int32(24))
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v1997 = m.ExcPending
	if v1997 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L253
	}
L61:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	if v427 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v383)+1040))
	v431 = F_readTimeLineHistory(m, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v435 = F_palloc0(m, v427<<(uint(int32(2))%32))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L64
	}
L64:
	;
	if v427 <= int32(0) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L249
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L245
	}
L67:
	;
	v952 = *(*int64)(unsafe.Add(mBase, uint32(v383)+1032))
	F_WaitForWalSummarization(m, v952)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L130
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+1080)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v383)+1072)) = int64(0)
	v926 = v411
	v950 = v412
	goto L67
L69:
	;
	goto L70
L70:
	;
	v460 = v411
	v462 = v411
	v469 = v411
	v484 = v412
	goto L71
L71:
	;
	v487 = v469 << (uint(int32(2)) % 32)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v487+v489)))
	if v431 != 0 {
		goto L76
	} else {
		goto L77
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+1080)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v383)+1072)) = v771
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v777)+12))
	v796 = int32(0)
	goto L106
L73:
	;
	if v737&int32(1) != 0 {
		goto L99
	} else {
		goto L100
	}
L74:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	v737 = v693
	v738 = v718
	goto L73
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L95
	}
L76:
	;
	v492 = int32(0)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v493 <= v492 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	goto L78
L78:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v435+v487)))
	if v614 != 0 {
		v693 = int32(0)
		goto L74
	} else {
		goto L94
	}
L79:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v435+v487)))
	if v607 == int32(0) {
		goto L75
	} else {
		goto L92
	}
L80:
	;
	v581 = v492
	v583 = int32(0)
	goto L79
L81:
	;
	goto L82
L82:
	;
	v497 = int32(0)
	if v497 < v493 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v500 = v493
	goto L85
L84:
	;
	v500 = v497
	goto L85
L85:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v431)+12))
	v503 = int32(0)
	v521 = v503
	v523 = v492
	v525 = v503
	goto L86
L86:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v502+v521<<(uint(int32(2))%32))))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)))
	if v501 == v552 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v581 = v559
	v583 = v557
	goto L79
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435+v487))) = v551
	v581 = v523
	v583 = v525
	goto L79
L89:
	;
	goto L90
L90:
	;
	v557 = base.B2i32(v462 == v552) | v525
	v559 = base.B2i32(v460 == v552) | v523
	v561 = v521 + int32(1)
	if v561 != v500 {
		v521 = v561
		v523 = v559
		v525 = v557
		goto L86
	} else {
		goto L91
	}
L91:
	;
	goto L87
L92:
	;
	if v583&int32(1) != 0 {
		v737 = v581
		v738 = v462
		goto L73
	} else {
		goto L93
	}
L93:
	;
	v693 = v581
	goto L74
L94:
	;
	goto L75
L95:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L96
	}
L96:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	*(*int32)(unsafe.Add(mBase, uint32(v417))) = v665
	F_errmsg(m, int32(_a_F_perform_base_backup_9), v417)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_10), int32(348), int32(_a_F_perform_base_backup_11))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v765 = int32(0)
	goto L101
L100:
	;
	v765 = v460
	goto L101
L101:
	;
	if v765 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v768 = *(*int64)(unsafe.Add(mBase, uint32(v491)+8))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	v770 = v769
	v771 = v768
	goto L104
L103:
	;
	v770 = v460
	v771 = v484
	goto L104
L104:
	;
	v773 = v469 + int32(1)
	if v773 != v427 {
		v460 = v770
		v462 = v738
		v469 = v773
		v484 = v771
		goto L71
	} else {
		goto L105
	}
L105:
	;
	goto L72
L106:
	;
	v824 = v796 << (uint(int32(2)) % 32)
	v825 = v435 + v824
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v825)))
	v827 = *(*int64)(unsafe.Add(mBase, uint32(v826)+8))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v824+v778)))
	v830 = *(*int64)(unsafe.Add(mBase, uint32(v829)+8))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	if v770 == v831 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v926 = v770
	v950 = v771
	goto L67
L108:
	;
	v865 = *(*int64)(unsafe.Add(mBase, uint32(v829)+16))
	if v738 == v831 {
		goto L119
	} else {
		goto L120
	}
L109:
	;
	if base.Ui64(v827) <= base.Ui64(v830) {
		goto L108
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	if v827 != v830 {
		goto L65
	} else {
		goto L117
	}
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L113
	}
L113:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L114
	}
L114:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	v842 = *(*int64)(unsafe.Add(mBase, uint32(v829)+8))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v825)))
	v844 = *(*int64)(unsafe.Add(mBase, uint32(v843)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+32)) = uint32(v844)
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+24)) = uint32(v842)
	v847 = int64(32)
	v848 = int64(base.Ui64(v842) >> (uint(v847) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+20)) = uint32(v848)
	*(*int32)(unsafe.Add(mBase, uint32(v417)+16)) = v841
	v852 = int64(base.Ui64(v844) >> (uint(v847) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+28)) = uint32(v852)
	F_errmsg(m, int32(_a_F_perform_base_backup_12), v417+int32(16))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_10), int32(415), int32(_a_F_perform_base_backup_11))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v907 = v796 + int32(1)
	if v907 != v427 {
		v796 = v907
		goto L106
	} else {
		goto L129
	}
L119:
	;
	v867 = *(*int64)(unsafe.Add(mBase, uint32(v383)+1032))
	if base.Ui64(v865) <= base.Ui64(v867) {
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v904 = *(*int64)(unsafe.Add(mBase, uint32(v826)+16))
	if v865 != v904 {
		goto L66
	} else {
		goto L128
	}
L122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L123
	}
L123:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L124
	}
L124:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	v877 = *(*int64)(unsafe.Add(mBase, uint32(v829)+16))
	v880 = *(*int64)(unsafe.Add(mBase, uint32(v383)+1032))
	*(*uint32)(unsafe.Add(mBase, uint32(v417-int32(-64)))) = uint32(v880)
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+56)) = uint32(v877)
	*(*int32)(unsafe.Add(mBase, uint32(v417)+48)) = v876
	v884 = int64(32)
	v885 = int64(base.Ui64(v880) >> (uint(v884) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+60)) = uint32(v885)
	v888 = int64(base.Ui64(v877) >> (uint(v884) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+52)) = uint32(v888)
	F_errmsg(m, int32(_a_F_perform_base_backup_13), v417+int32(48))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L125
	}
L125:
	;
	F_errhint(m, int32(_a_F_perform_base_backup_14), int32(0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_10), int32(437), int32(_a_F_perform_base_backup_11))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v955 = int32(0)
	v957 = *(*int64)(unsafe.Add(mBase, uint32(v383)+1032))
	v958 = F_GetWalSummaries(m, v955, v950, v957)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L131
	}
L131:
	;
	if v431 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v417)+112)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v417)+108)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v417)+104)) = v1376
	*(*int32)(unsafe.Add(mBase, uint32(v417)+100)) = v1374
	*(*int32)(unsafe.Add(mBase, uint32(v417)+96)) = v1377
	F_errmsg(m, int32(_a_F_perform_base_backup_15), v417+int32(96))
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L243
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[12])) = v420
	m.G0 = v417 + int32(2336)
	goto L59
L134:
	;
	v962 = F_CreateEmptyBlockRefTable(m)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v965 <= int32(0) {
		v1478 = v955
		goto L138
	} else {
		goto L139
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+28)) = v962
	goto L133
L138:
	;
	v1501 = F_CreateEmptyBlockRefTable(m)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L204
	}
L139:
	;
	v968 = int32(0)
	v988 = v968
	v989 = v968
	v990 = v955
	goto L140
L140:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v431)+12))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1013+v988<<(uint(int32(2))%32))))
	v1018 = *(*int64)(unsafe.Add(mBase, uint32(v1017)+16))
	v1019 = *(*int64)(unsafe.Add(mBase, uint32(v1017)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v417)+240)) = int64(0)
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v1017)))
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v383)+1040))
	if v1022 == v1023 {
		goto L144
	} else {
		goto L145
	}
L141:
	;
	v1478 = v1430
	goto L138
L142:
	;
	v1455 = v988 + int32(1)
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v1455 < v1456 {
		v988 = v1455
		v989 = v1453
		v990 = v1430
		goto L140
	} else {
		goto L203
	}
L143:
	;
	if v926 == v1022 {
		goto L150
	} else {
		goto L151
	}
L144:
	;
	v1025 = *(*int64)(unsafe.Add(mBase, uint32(v383)+1032))
	v1029 = v1025
	goto L143
L145:
	;
	goto L146
L146:
	;
	if v989&int32(1) != 0 {
		v1029 = v1018
		goto L143
	} else {
		goto L147
	}
L147:
	;
	v1430 = v990
	v1453 = int32(0)
	goto L142
L148:
	;
	if v1358 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L149:
	;
	if v1191 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L150:
	;
	v1031 = v950
	goto L152
L151:
	;
	v1031 = v1019
	goto L152
L152:
	;
	v1032 = int32(0)
	if v958 == v1032 {
		v1191 = v1032
		goto L149
	} else {
		goto L153
	}
L153:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v958)+4))
	if int32(0) < v1037 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v1064 = v1032
	v1066 = v1032
	goto L157
L155:
	;
	v1129 = v1032
	goto L156
L156:
	;
	v1191 = v1129
	goto L149
L157:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v958)+12))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1083+v1066<<(uint(int32(2))%32))))
	if v1022 != 0 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	v1129 = v1100
	goto L156
L159:
	;
	v1102 = v1066 + int32(1)
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v958)+4))
	if v1102 < v1103 {
		v1064 = v1100
		v1066 = v1102
		goto L157
	} else {
		goto L173
	}
L160:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1087)+16))
	if v1022 != v1088 {
		v1100 = v1064
		goto L159
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if v1031 != int64(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L162
L164:
	;
	v1092 = *(*int64)(unsafe.Add(mBase, uint32(v1087)+8))
	if base.Ui64(v1092) < base.Ui64(v1031) {
		v1100 = v1064
		goto L159
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	if v1029 != int64(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	goto L166
L168:
	;
	v1096 = *(*int64)(unsafe.Add(mBase, uint32(v1087)))
	if base.Ui64(v1029) < base.Ui64(v1096) {
		v1100 = v1064
		goto L159
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v1098 = F_lappend(m, v1064, v1087)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L172
	}
L171:
	;
	goto L170
L172:
	;
	v1100 = v1098
	goto L159
L173:
	;
	goto L158
L174:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v417+int32(240)))) = v1308
	v1358 = int32(0)
	goto L148
L175:
	;
	v1308 = int64(0)
	goto L174
L176:
	;
	goto L177
L177:
	;
	v1197 = F_list_copy(m, v1191)
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L178
	}
L178:
	;
	F_list_sort(m, v1197, int32(459))
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L179
	}
L179:
	;
	if v1197 == int32(0) {
		v1308 = v1031
		goto L174
	} else {
		goto L180
	}
L180:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+4))
	if v1204 <= int32(0) {
		v1308 = v1031
		goto L174
	} else {
		goto L181
	}
L181:
	;
	v1207 = int32(0)
	if v1207 < v1204 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v1211 = v1204
	goto L184
L183:
	;
	v1211 = v1207
	goto L184
L184:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+12))
	v1232 = v1207
	v1251 = v1031
	goto L185
L185:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1212+v1232<<(uint(int32(2))%32))))
	v1260 = *(*int64)(unsafe.Add(mBase, uint32(v1259)))
	if base.Ui64(v1251) < base.Ui64(v1260) {
		v1308 = v1251
		goto L174
	} else {
		goto L187
	}
L186:
	;
	v1308 = v1266
	goto L174
L187:
	;
	v1262 = *(*int64)(unsafe.Add(mBase, uint32(v1259)+8))
	if base.Ui64(v1262) <= base.Ui64(v1251) {
		v1266 = v1251
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1268 = v1232 + int32(1)
	if v1268 != v1211 {
		v1232 = v1268
		v1251 = v1266
		goto L185
	} else {
		goto L191
	}
L189:
	;
	if base.Ui64(v1262) < base.Ui64(v1029) {
		v1266 = v1262
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v1358 = int32(1)
	goto L148
L191:
	;
	goto L186
L192:
	;
	v1361 = *(*int64)(unsafe.Add(mBase, uint32(v417)+240))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v1405 = F_list_concat(m, v990, v1191)
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L201
	}
L195:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L196
	}
L196:
	;
	v1369 = int64(32)
	v1371 = base.I32_wrap_i64(int64(base.Ui64(v1029) >> (uint(v1369) % 64)))
	v1374 = base.I32_wrap_i64(int64(base.Ui64(v1031) >> (uint(v1369) % 64)))
	v1375 = base.I32_wrap_i64(v1029)
	v1376 = base.I32_wrap_i64(v1031)
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1017)))
	if v1361 == int64(0) {
		goto L132
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v417)+160)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v417)+156)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v417)+152)) = v1376
	*(*int32)(unsafe.Add(mBase, uint32(v417)+148)) = v1374
	*(*int32)(unsafe.Add(mBase, uint32(v417)+144)) = v1377
	F_errmsg(m, int32(_a_F_perform_base_backup_16), v417+int32(144))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L198
	}
L198:
	;
	v1390 = *(*int64)(unsafe.Add(mBase, uint32(v417)+240))
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+132)) = uint32(v1390)
	v1393 = int64(base.Ui64(v1390) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+128)) = uint32(v1393)
	F_errdetail(m, int32(_a_F_perform_base_backup_17), v417+int32(128))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_10), int32(537), int32(_a_F_perform_base_backup_11))
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1017)))
	if v1407 == v926 {
		v1478 = v1405
		goto L138
	} else {
		goto L202
	}
L202:
	;
	v1430 = v1405
	v1453 = int32(1)
	goto L142
L203:
	;
	goto L141
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+28)) = v1501
	if v1478 == int32(0) {
		goto L133
	} else {
		goto L205
	}
L205:
	;
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1478)+4))
	if v1506 <= int32(0) {
		goto L133
	} else {
		goto L206
	}
L206:
	;
	v1527 = int32(0)
	goto L207
L207:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1478)+12))
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1553+v1527<<(uint(int32(2))%32))))
	v1558 = F_OpenWalSummaryFile(m, v1557)
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L209
	}
L208:
	;
	goto L133
L209:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v417)+2328)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v417)+2320)) = v1558
	v1565 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L210
	}
L210:
	;
	if v1565 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v417)+2320))
	v1569 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[13]))
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1569+v1567*int32(48))+32))
	goto L214
L212:
	;
	goto L213
L213:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v417)+2320))
	v1589 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[13]))
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1589+v1587*int32(48))+32))
	goto L217
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v417)+80)) = v1573
	F_errmsg_internal(m, int32(_a_F_perform_base_backup_18), v417+int32(80))
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_10), int32(586), int32(_a_F_perform_base_backup_11))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L216
	}
L216:
	;
	goto L213
L217:
	;
	v1594 = F_CreateBlockRefTableReader(m, v417+int32(2320), v1593)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L218
	}
L218:
	;
	v1602 = F_BlockRefTableReaderNextRelation(m, v1594, v417+int32(2308), v417+int32(2304), v417+int32(2300))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L219
	}
L219:
	;
	if v1602 != 0 {
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
	F_DestroyBlockRefTableReader(m, v1594)
	mBase = m.M
	v1859 = m.ExcPending
	if v1859 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L240
	}
L223:
	;
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v417)+2304))
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v417)+2300))
	F_BlockRefTableSetLimitBlock(m, v1647, v417+int32(2308), v1650, v1651)
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L225
	}
L224:
	;
	goto L222
L225:
	;
	v1657 = F_BlockRefTableReaderGetBlocks(m, v1594, v417+int32(240), int32(512))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L226
	}
L226:
	;
	if v1657 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v1677 = v1657
	goto L230
L228:
	;
	goto L229
L229:
	;
	v1813 = F_BlockRefTableReaderNextRelation(m, v1594, v417+int32(2308), v417+int32(2304), v417+int32(2300))
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L238
	}
L230:
	;
	v1719 = int32(0)
	goto L232
L231:
	;
	goto L229
L232:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v417)+2304))
	v1751 = v417 + int32(240)
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1751+v1719<<(uint(int32(2))%32))))
	F_BlockRefTableMarkBlockModified(m, v1746, v417+int32(2308), v1749, v1755)
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L234
	}
L233:
	;
	v1762 = F_BlockRefTableReaderGetBlocks(m, v1594, v1751, int32(512))
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L236
	}
L234:
	;
	v1759 = v1719 + int32(1)
	if v1759 != v1677 {
		v1719 = v1759
		goto L232
	} else {
		goto L235
	}
L235:
	;
	goto L233
L236:
	;
	if v1762 != 0 {
		v1677 = v1762
		goto L230
	} else {
		goto L237
	}
L237:
	;
	goto L231
L238:
	;
	if v1813 != 0 {
		goto L223
	} else {
		goto L239
	}
L239:
	;
	goto L224
L240:
	;
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v417)+2320))
	F_FileClose(m, v1860)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L241
	}
L241:
	;
	v1864 = v1527 + int32(1)
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v1478)+4))
	if v1864 < v1865 {
		v1527 = v1864
		goto L207
	} else {
		goto L242
	}
L242:
	;
	goto L208
L243:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_10), int32(528), int32(_a_F_perform_base_backup_11))
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v1936 = m.ExcPending
	if v1936 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L246
	}
L246:
	;
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	v1938 = *(*int64)(unsafe.Add(mBase, uint32(v829)+16))
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v435+v796<<(uint(int32(2))%32))))
	v1943 = *(*int64)(unsafe.Add(mBase, uint32(v1942)+16))
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+192)) = uint32(v1943)
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+184)) = uint32(v1938)
	v1946 = int64(32)
	v1947 = int64(base.Ui64(v1938) >> (uint(v1946) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+180)) = uint32(v1947)
	*(*int32)(unsafe.Add(mBase, uint32(v417)+176)) = v1937
	v1951 = int64(base.Ui64(v1943) >> (uint(v1946) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+188)) = uint32(v1951)
	F_errmsg(m, int32(_a_F_perform_base_backup_19), v417+int32(176))
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_10), int32(447), int32(_a_F_perform_base_backup_11))
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v1969 = m.ExcPending
	if v1969 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L250
	}
L250:
	;
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	v1971 = *(*int64)(unsafe.Add(mBase, uint32(v829)+8))
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v825)))
	v1973 = *(*int64)(unsafe.Add(mBase, uint32(v1972)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+224)) = uint32(v1973)
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+216)) = uint32(v1971)
	v1976 = int64(32)
	v1977 = int64(base.Ui64(v1971) >> (uint(v1976) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+212)) = uint32(v1977)
	*(*int32)(unsafe.Add(mBase, uint32(v417)+208)) = v1970
	v1981 = int64(base.Ui64(v1973) >> (uint(v1976) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v417)+220)) = uint32(v1981)
	F_errmsg(m, int32(_a_F_perform_base_backup_20), v417+int32(208))
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_10), int32(425), int32(_a_F_perform_base_backup_11))
	mBase = m.M
	v1992 = m.ExcPending
	if v1992 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v2000 = m.ExcPending
	if v2000 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L254
	}
L254:
	;
	F_errmsg(m, int32(_a_F_perform_base_backup_21), int32(0))
	mBase = m.M
	v2004 = m.ExcPending
	if v2004 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_10), int32(292), int32(_a_F_perform_base_backup_11))
	mBase = m.M
	v2009 = m.ExcPending
	if v2009 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	*(*int64)(unsafe.Add(mBase, uint32(v2067)+16)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v2084 = *(*int32)(unsafe.Add(mBase, uint32(v58)+2088))
	v2085 = F_lappend(m, v2084, v2067)
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2088)) = v2085
	v2088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+4)))
	if v2088 == int32(1) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v2108 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[7]))
	if v2108 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	v2334 = v91
	v2335 = v92
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = int32(_a_F_perform_base_backup_22)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+16)) = v58 + int32(2088)
	v2346 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(v2346)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	m.T0[v2347].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L278
	}
L262:
	;
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v58)+2088))
	if v2141 == int32(0) {
		v2289 = v91
		v2290 = v92
		goto L266
	} else {
		goto L267
	}
L263:
	;
	goto L262
L264:
	;
	v2112 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[8])))
	if v2112&int32(1) == int32(0) {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v2117 = int32(_a_F_perform_base_backup_7)
	v2119 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	v2120 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v2119 + v2120
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v2108)))
	*(*int32)(unsafe.Add(mBase, uint32(v2108))) = v2123 + v2120
	*(*int64)(unsafe.Add(mBase, uint32(v2108+int32(0))+232)) = int64(2)
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2108)))
	*(*int32)(unsafe.Add(mBase, uint32(v2108))) = v2131 + v2120
	v2137 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v2137 - v2120
	goto L263
L266:
	;
	v2296 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v389))) = uint8(v2296)
	v2334 = v2289
	v2335 = v2290
	goto L261
L267:
	;
	v2144 = int32(0)
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	if v2145 <= v2144 {
		v2289 = v91
		v2290 = v92
		goto L266
	} else {
		goto L268
	}
L268:
	;
	v2152 = v2144
	v2184 = v91
	v2185 = v92
	goto L269
L269:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+12))
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v2191+v2152<<(uint(int32(2))%32))))
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(v2195)+4))
	if v2196 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	v2289 = v2242
	v2290 = v2243
	goto L266
L271:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2195)+16)) = v2244
	v2246 = *(*int64)(unsafe.Add(mBase, uint32(v384)))
	*(*int64)(unsafe.Add(mBase, uint32(v384))) = v2246 + v2244
	v2250 = v2152 + int32(1)
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	if v2250 < v2251 {
		v2152 = v2250
		v2184 = v2242
		v2185 = v2243
		goto L269
	} else {
		goto L277
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2184
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2185
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v2213 = int32(1)
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(v58)+2088))
	v2217 = int32(0)
	v2220 = F_sendDir(m, v56, int32(_a_F_perform_base_backup_23), v2213, v2213, v2215, v2213, v2217, v2217, v2217)
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(v2195)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2184
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2185
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v2237 = int32(0)
	v2239 = F_sendTablespace(m, v56, v2196, v2222, int32(1), v2237, v2237)
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L276
	}
L275:
	;
	v2242 = v2184
	v2243 = v2220
	v2244 = v2220
	goto L271
L276:
	;
	v2242 = v2239
	v2243 = v2185
	v2244 = v2239
	goto L271
L277:
	;
	goto L270
L278:
	;
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(v58)+2088))
	if v2363 == int32(0) {
		goto L52
	} else {
		goto L279
	}
L279:
	;
	v2366 = int32(0)
	v2367 = *(*int32)(unsafe.Add(mBase, uint32(v2363)+4))
	if v2367 <= v2366 {
		goto L52
	} else {
		goto L280
	}
L280:
	;
	v2388 = v2366
	goto L281
L281:
	;
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v2363)+12))
	v2417 = *(*int32)(unsafe.Add(mBase, uint32(v2413+v2388<<(uint(int32(2))%32))))
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(v2417)+4))
	if v2418 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	goto L52
L283:
	;
	v2711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+7)))
	if v2711 == int32(1) {
		goto L309
	} else {
		goto L310
	}
L284:
	;
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v2422 = *(*int32)(unsafe.Add(mBase, uint32(v2421)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	m.T0[v2422].(func(*base.Module, int32, int32))(m, v56, int32(_a_F_perform_base_backup_24))
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v2651 = *(*int32)(unsafe.Add(mBase, uint32(v2417)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+320)) = v2651
	v2669 = F_psprintf(m, int32(_a_F_perform_base_backup_25), v58+int32(320))
	mBase = m.M
	v2670 = m.ExcPending
	if v2670 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L305
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v2453 = F_build_backup_content(m, v383, int32(0))
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v2470 = v58 + int32(2056)
	F_sendFileWithContent(m, v56, int32(_a_F_perform_base_backup_26), v2453, v2470)
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L289
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_pfree(m, v2453)
	mBase = m.M
	v2487 = m.ExcPending
	if v2487 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L290
	}
L290:
	;
	v2488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+16)))
	if v2488 == int32(1) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v2491 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_sendFileWithContent(m, v56, int32(_a_F_perform_base_backup_27), v2491, v2470)
	mBase = m.M
	v2507 = m.ExcPending
	if v2507 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v2523 = int32(1)
	v2524 = int32(0)
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(v58)+2088))
	v2531 = F_sendDir(m, v56, int32(_a_F_perform_base_backup_23), v2523, v2524, v2525, v2488^v2523, v58+int32(2056), v2524, v57)
	mBase = m.M
	v2532 = m.ExcPending
	if v2532 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v2551 = F___fstatat(m, int32(-100), int32(_a_F_perform_base_backup_28), v58+int32(1792), int32(256))
	mBase = m.M
	goto L296
L296:
	;
	if v2551 != 0 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2568 = m.ExcPending
	if v2568 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v2635 = int32(_a_F_perform_base_backup_28)
	v2639 = int32(0)
	v2649 = F_sendFile(m, v56, v2635, v2635, v58+int32(1792), v2639, v2639, v2639, v2639, v2639, v58+int32(2056), v2639, v2639, v2639)
	mBase = m.M
	v2650 = m.ExcPending
	if v2650 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L304
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errcode_for_file_access(m)
	mBase = m.M
	v2583 = m.ExcPending
	if v2583 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+304)) = int32(_a_F_perform_base_backup_28)
	F_errmsg(m, int32(_a_F_perform_base_backup_29), v58+int32(304))
	mBase = m.M
	v2603 = m.ExcPending
	if v2603 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errfinish(m, int32(_a_F_perform_base_backup_30), int32(358), int32(_a_F_perform_base_backup_31))
	mBase = m.M
	v2621 = m.ExcPending
	if v2621 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v2671 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v2671)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	m.T0[v2672].(func(*base.Module, int32, int32))(m, v56, v2669)
	mBase = m.M
	v2687 = m.ExcPending
	if v2687 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L306
	}
L306:
	;
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(v2417)))
	v2689 = *(*int32)(unsafe.Add(mBase, uint32(v2417)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v2706 = F_sendTablespace(m, v56, v2689, v2688, int32(0), v58+int32(2056), v57)
	mBase = m.M
	v2707 = m.ExcPending
	if v2707 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L307
	}
L307:
	;
	goto L283
L308:
	;
	v2758 = v2388 + int32(1)
	v2759 = *(*int32)(unsafe.Add(mBase, uint32(v2363)+4))
	if v2758 < v2759 {
		v2388 = v2758
		goto L281
	} else {
		goto L315
	}
L309:
	;
	v2714 = *(*int32)(unsafe.Add(mBase, uint32(v2417)+4))
	if v2714 == int32(0) {
		goto L308
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v2719 = int32(1024)
	base.MemoryFill(m, v2717, int32(0), v2719)
	v2721 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v2722 = *(*int32)(unsafe.Add(mBase, uint32(v2721)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	m.T0[v2722].(func(*base.Module, int32, int32))(m, v56, v2719)
	mBase = m.M
	v2738 = m.ExcPending
	if v2738 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L313
	}
L312:
	;
	goto L311
L313:
	;
	v2739 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v2739)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	m.T0[v2740].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v2755 = m.ExcPending
	if v2755 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v2795 = int32(0)
	F_do_pg_abort_backup(m, v2795, v2795)
	mBase = m.M
	v2798 = m.ExcPending
	if v2798 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_pg_re_throw(m)
	mBase = m.M
	v2813 = m.ExcPending
	if v2813 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L318
	}
L318:
	;
	goto L1
L319:
	;
	v2881 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2880)+4)))
	v2883 = v2881
	goto L321
L320:
	;
	v2883 = int64(0)
	goto L321
L321:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2872)+8)) = v2883
	goto L324
L322:
	;
	m.G0 = v2872 + int32(32)
	v3064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_do_pg_backup_stop(m, v383, (v3064^int32(-1))&int32(1))
	mBase = m.M
	v3083 = m.ExcPending
	if v3083 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L339
	}
L323:
	;
	goto L322
L324:
	;
	v2897 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[7]))
	if v2897 == int32(0) {
		goto L323
	} else {
		goto L325
	}
L325:
	;
	v2901 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[8])))
	if v2901&int32(1) == int32(0) {
		goto L323
	} else {
		goto L326
	}
L326:
	;
	v2906 = int32(_a_F_perform_base_backup_7)
	v2908 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	v2909 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v2908 + v2909
	v2912 = *(*int32)(unsafe.Add(mBase, uint32(v2897)))
	*(*int32)(unsafe.Add(mBase, uint32(v2897))) = v2912 + v2909
	goto L328
L327:
	;
	v3042 = *(*int32)(unsafe.Add(mBase, uint32(v2897)))
	v3043 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2897))) = v3042 + v3043
	v3046 = int32(_a_F_perform_base_backup_7)
	v3048 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v3048 - v3043
	goto L323
L328:
	;
	goto L330
L330:
	;
	goto L331
L331:
	;
	v3007 = int32(0)
	v3010 = int32(0)
	goto L336
L336:
	;
	v3019 = *(*int32)(unsafe.Add(mBase, uint32(v2872+int32(24)+v3010<<(uint(int32(2))%32))))
	v3020 = int32(3)
	v3026 = *(*int64)(unsafe.Add(mBase, uint32(v2872+v3010<<(uint(v3020)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2897+int32(232)+v3019<<(uint(v3020)%32)))) = v3026
	v3028 = int32(1)
	v3031 = v3007 + v3028
	if v3031 != int32(2) {
		v3007 = v3031
		v3010 = v3010 + v3028
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
	v3084 = *(*int32)(unsafe.Add(mBase, uint32(v383)+1096))
	v3085 = *(*int64)(unsafe.Add(mBase, uint32(v383)+1088))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_free_attrmap(m, v385)
	mBase = m.M
	v3100 = m.ExcPending
	if v3100 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L340
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_pfree(m, v383)
	mBase = m.M
	v3115 = m.ExcPending
	if v3115 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_cancel_before_shmem_exit(m, int32(407), int32(0))
	mBase = m.M
	v3132 = m.ExcPending
	if v3132 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L342
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[10])) = v387
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[11])) = v388
	v3137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+7)))
	if v3137 != 0 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3155 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[7]))
	if v3155 == int32(0) {
		goto L347
	} else {
		goto L348
	}
L344:
	;
	v5454 = v67
	v5455 = v68
	v5456 = v69
	goto L345
L345:
	;
	v5485 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v5486 = *(*int64)(unsafe.Add(mBase, uint32(v386)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5456
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v5501 = m.G0
	v5503 = v5501 - int32(80)
	m.G0 = v5503
	v5506 = v58 + int32(2056)
	v5507 = *(*int32)(unsafe.Add(mBase, uint32(v5506)))
	if v5507 != 0 {
		goto L559
	} else {
		goto L560
	}
L346:
	;
	v3188 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v3189 = *(*int64)(unsafe.Add(mBase, uint32(v386)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3204 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+288)) = v3188
	v3206 = base.I64_div_u_s(v3189, v3204)
	v3208 = base.I64_div_u_s(int64(4294967296), v3204)
	v3209 = base.I64_div_u_s(v3206, v3208)
	*(*uint32)(unsafe.Add(mBase, uint32(v58)+292)) = uint32(v3209)
	v3212 = v3206 - v3208*v3209
	*(*uint32)(unsafe.Add(mBase, uint32(v58)+296)) = uint32(v3212)
	v3220 = F_pg_snprintf(m, v58+int32(608), int32(64), int32(_a_F_perform_base_backup_32), v58+int32(288))
	mBase = m.M
	v3221 = m.ExcPending
	if v3221 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L350
	}
L347:
	;
	goto L346
L348:
	;
	v3159 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[8])))
	if v3159&int32(1) == int32(0) {
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v3164 = int32(_a_F_perform_base_backup_7)
	v3166 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	v3167 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v3166 + v3167
	v3170 = *(*int32)(unsafe.Add(mBase, uint32(v3155)))
	*(*int32)(unsafe.Add(mBase, uint32(v3155))) = v3170 + v3167
	*(*int64)(unsafe.Add(mBase, uint32(v3155+int32(0))+232)) = int64(5)
	v3178 = *(*int32)(unsafe.Add(mBase, uint32(v3155)))
	*(*int32)(unsafe.Add(mBase, uint32(v3155))) = v3178 + v3167
	v3184 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v3184 - v3167
	goto L347
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3236 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+272)) = v3084
	v3240 = base.I64_div_u_s(v3085-int64(1), v3236)
	v3242 = base.I64_div_u_s(int64(4294967296), v3236)
	v3243 = base.I64_div_u_s(v3240, v3242)
	*(*uint32)(unsafe.Add(mBase, uint32(v58)+276)) = uint32(v3243)
	v3246 = v3240 - v3242*v3243
	*(*uint32)(unsafe.Add(mBase, uint32(v58)+280)) = uint32(v3246)
	v3254 = F_pg_snprintf(m, v58+int32(544), int32(64), int32(_a_F_perform_base_backup_32), v58+int32(272))
	mBase = m.M
	v3255 = m.ExcPending
	if v3255 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3270 = F_AllocateDir(m, int32(_a_F_perform_base_backup_33))
	mBase = m.M
	v3271 = m.ExcPending
	if v3271 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L352
	}
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3286 = F_ReadDir(m, v3270, int32(_a_F_perform_base_backup_33))
	mBase = m.M
	v3287 = m.ExcPending
	if v3287 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L354
	}
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_FreeDir(m, v3270)
	mBase = m.M
	v3825 = m.ExcPending
	if v3825 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L433
	}
L354:
	;
	if v3286 == int32(0) {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v3290 = int32(0)
	v3780 = v67
	v3781 = v68
	v3782 = v69
	v3784 = v3290
	v3785 = v3290
	goto L353
L356:
	;
	goto L357
L357:
	;
	v3294 = int32(8)
	v3295 = v58 + int32(544) | v3294
	v3299 = v58 + int32(608) | v3294
	v3300 = int32(0)
	v3306 = v3286
	v3314 = v67
	v3315 = v68
	v3316 = v69
	v3318 = v3300
	v3319 = v3300
	goto L358
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3315
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3314
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3316
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3359 = v3306 + int32(19)
	v3360 = F_strlen(m, v3359)
	mBase = m.M
	switch v3360 - int32(16) {
	case 0:
		goto L361
	default:
		v3748 = v3315
		v3749 = v3316
		v3750 = v3318
		v3751 = v3319
		goto L360
	case 8:
		goto L362
	}
L359:
	;
	v3780 = v3766
	v3781 = v3748
	v3782 = v3749
	v3784 = v3750
	v3785 = v3751
	goto L353
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3748
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3314
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3749
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3766 = F_ReadDir(m, v3270, int32(_a_F_perform_base_backup_33))
	mBase = m.M
	v3767 = m.ExcPending
	if v3767 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L431
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3315
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3314
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3316
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3589 = int32(_a_F_perform_base_backup_34)
	v3593 = m.G0
	v3595 = v3593 - int32(32)
	v3596 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3595)+24)) = v3596
	*(*int64)(unsafe.Add(mBase, uint32(v3595)+16)) = v3596
	*(*int64)(unsafe.Add(mBase, uint32(v3595)+8)) = v3596
	*(*int64)(unsafe.Add(mBase, uint32(v3595))) = v3596
	v3604 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[15])))
	if v3604 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3315
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3314
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3316
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3376 = int32(_a_F_perform_base_backup_34)
	v3380 = m.G0
	v3382 = v3380 - int32(32)
	v3383 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3382)+24)) = v3383
	*(*int64)(unsafe.Add(mBase, uint32(v3382)+16)) = v3383
	*(*int64)(unsafe.Add(mBase, uint32(v3382)+8)) = v3383
	*(*int64)(unsafe.Add(mBase, uint32(v3382))) = v3383
	v3391 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[15])))
	if v3391 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L363:
	;
	if v3459 != int32(24) {
		v3748 = v3315
		v3749 = v3316
		v3750 = v3318
		v3751 = v3319
		goto L360
	} else {
		goto L382
	}
L364:
	;
	v3459 = int32(0)
	goto L363
L365:
	;
	goto L366
L366:
	;
	v3395 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[16])))
	if v3395 == int32(0) {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v3399 = v3359
	goto L370
L368:
	;
	goto L369
L369:
	;
	v3409 = v3376
	v3410 = v3391
	goto L373
L370:
	;
	v3405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3399))))
	if v3405 == v3391 {
		v3399 = v3399 + int32(1)
		goto L370
	} else {
		goto L372
	}
L371:
	;
	v3459 = v3399 - v3359
	goto L363
L372:
	;
	goto L371
L373:
	;
	v3417 = v3382 + int32(base.Ui32(v3410)>>(uint(int32(3))%32))&int32(28)
	v3418 = *(*int32)(unsafe.Add(mBase, uint32(v3417)))
	v3419 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3417))) = v3418 | v3419<<(uint(v3410)%32)
	v3423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3409)+1)))
	if v3423 != 0 {
		v3409 = v3409 + v3419
		v3410 = v3423
		goto L373
	} else {
		goto L375
	}
L374:
	;
	v3426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3359))))
	if v3426 == int32(0) {
		v3449 = v3359
		goto L376
	} else {
		goto L377
	}
L375:
	;
	goto L374
L376:
	;
	v3459 = v3449 - v3359
	goto L363
L377:
	;
	v3430 = v3359
	v3431 = v3426
	goto L378
L378:
	;
	v3439 = *(*int32)(unsafe.Add(mBase, uint32(v3382+int32(base.Ui32(v3431)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3439)>>(uint(v3431)%32))&int32(1) == int32(0) {
		v3449 = v3430
		goto L376
	} else {
		goto L380
	}
L379:
	;
	v3449 = v3447
	goto L376
L380:
	;
	v3445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3430)+1)))
	v3447 = v3430 + int32(1)
	if v3445 != 0 {
		v3430 = v3447
		v3431 = v3445
		goto L378
	} else {
		goto L381
	}
L381:
	;
	goto L379
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3315
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3314
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3316
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3476 = v3306 + int32(27)
	v3479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476))))
	v3482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3299))))
	if base.B2i32(v3479 == int32(0))|base.B2i32(v3479 != v3482) != 0 {
		v3500 = v3479
		v3501 = v3482
		goto L384
	} else {
		goto L385
	}
L383:
	;
	if v3500-v3501 < int32(0) {
		v3748 = v3315
		v3749 = v3316
		v3750 = v3318
		v3751 = v3319
		goto L360
	} else {
		goto L390
	}
L384:
	;
	goto L383
L385:
	;
	v3485 = v3476
	v3486 = v3299
	goto L386
L386:
	;
	v3489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3486)+1)))
	v3490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3485)+1)))
	if v3490 == int32(0) {
		v3500 = v3490
		v3501 = v3489
		goto L384
	} else {
		goto L388
	}
L387:
	;
	v3500 = v3490
	v3501 = v3489
	goto L384
L388:
	;
	v3493 = int32(1)
	if v3490 == v3489 {
		v3485 = v3485 + v3493
		v3486 = v3486 + v3493
		goto L386
	} else {
		goto L389
	}
L389:
	;
	goto L387
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3315
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3314
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3316
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476))))
	v3523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3295))))
	if base.B2i32(v3520 == int32(0))|base.B2i32(v3520 != v3523) != 0 {
		v3541 = v3520
		v3542 = v3523
		goto L392
	} else {
		goto L393
	}
L391:
	;
	if int32(0) < v3541-v3542 {
		v3748 = v3315
		v3749 = v3316
		v3750 = v3318
		v3751 = v3319
		goto L360
	} else {
		goto L398
	}
L392:
	;
	goto L391
L393:
	;
	v3526 = v3476
	v3527 = v3295
	goto L394
L394:
	;
	v3530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3527)+1)))
	v3531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3526)+1)))
	if v3531 == int32(0) {
		v3541 = v3531
		v3542 = v3530
		goto L392
	} else {
		goto L396
	}
L395:
	;
	v3541 = v3531
	v3542 = v3530
	goto L392
L396:
	;
	v3534 = int32(1)
	if v3531 == v3530 {
		v3526 = v3526 + v3534
		v3527 = v3527 + v3534
		goto L394
	} else {
		goto L397
	}
L397:
	;
	goto L395
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3315
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3314
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3316
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3559 = F_pstrdup(m, v3359)
	mBase = m.M
	v3560 = m.ExcPending
	if v3560 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L399
	}
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3315
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3314
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3316
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3574 = F_lappend(m, v3318, v3559)
	mBase = m.M
	v3575 = m.ExcPending
	if v3575 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L400
	}
L400:
	;
	v3748 = v3315
	v3749 = v3574
	v3750 = v3574
	v3751 = v3319
	goto L360
L401:
	;
	if v3672 != int32(8) {
		v3748 = v3315
		v3749 = v3316
		v3750 = v3318
		v3751 = v3319
		goto L360
	} else {
		goto L420
	}
L402:
	;
	v3672 = int32(0)
	goto L401
L403:
	;
	goto L404
L404:
	;
	v3608 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[16])))
	if v3608 == int32(0) {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v3612 = v3359
	goto L408
L406:
	;
	goto L407
L407:
	;
	v3622 = v3589
	v3623 = v3604
	goto L411
L408:
	;
	v3618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3612))))
	if v3618 == v3604 {
		v3612 = v3612 + int32(1)
		goto L408
	} else {
		goto L410
	}
L409:
	;
	v3672 = v3612 - v3359
	goto L401
L410:
	;
	goto L409
L411:
	;
	v3630 = v3595 + int32(base.Ui32(v3623)>>(uint(int32(3))%32))&int32(28)
	v3631 = *(*int32)(unsafe.Add(mBase, uint32(v3630)))
	v3632 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3630))) = v3631 | v3632<<(uint(v3623)%32)
	v3636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3622)+1)))
	if v3636 != 0 {
		v3622 = v3622 + v3632
		v3623 = v3636
		goto L411
	} else {
		goto L413
	}
L412:
	;
	v3639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3359))))
	if v3639 == int32(0) {
		v3662 = v3359
		goto L414
	} else {
		goto L415
	}
L413:
	;
	goto L412
L414:
	;
	v3672 = v3662 - v3359
	goto L401
L415:
	;
	v3643 = v3359
	v3644 = v3639
	goto L416
L416:
	;
	v3652 = *(*int32)(unsafe.Add(mBase, uint32(v3595+int32(base.Ui32(v3644)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3652)>>(uint(v3644)%32))&int32(1) == int32(0) {
		v3662 = v3643
		goto L414
	} else {
		goto L418
	}
L417:
	;
	v3662 = v3660
	goto L414
L418:
	;
	v3658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3643)+1)))
	v3660 = v3643 + int32(1)
	if v3658 != 0 {
		v3643 = v3660
		v3644 = v3658
		goto L416
	} else {
		goto L419
	}
L419:
	;
	goto L417
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3315
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3314
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3316
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3689 = v3306 + int32(27)
	v3690 = int32(_a_F_perform_base_backup_35)
	v3693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3689))))
	v3696 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[17])))
	if base.B2i32(v3693 == int32(0))|base.B2i32(v3693 != v3696) != 0 {
		v3714 = v3693
		v3715 = v3696
		goto L422
	} else {
		goto L423
	}
L421:
	;
	if v3714-v3715 != 0 {
		v3748 = v3315
		v3749 = v3316
		v3750 = v3318
		v3751 = v3319
		goto L360
	} else {
		goto L428
	}
L422:
	;
	goto L421
L423:
	;
	v3699 = v3689
	v3700 = v3690
	goto L424
L424:
	;
	v3703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3700)+1)))
	v3704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3699)+1)))
	if v3704 == int32(0) {
		v3714 = v3704
		v3715 = v3703
		goto L422
	} else {
		goto L426
	}
L425:
	;
	v3714 = v3704
	v3715 = v3703
	goto L422
L426:
	;
	v3707 = int32(1)
	if v3704 == v3703 {
		v3699 = v3699 + v3707
		v3700 = v3700 + v3707
		goto L424
	} else {
		goto L427
	}
L427:
	;
	goto L425
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3315
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3314
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3316
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3730 = F_pstrdup(m, v3359)
	mBase = m.M
	v3731 = m.ExcPending
	if v3731 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3315
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3314
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3316
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3745 = F_lappend(m, v3319, v3730)
	mBase = m.M
	v3746 = m.ExcPending
	if v3746 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L430
	}
L430:
	;
	v3748 = v3745
	v3749 = v3316
	v3750 = v3318
	v3751 = v3745
	goto L360
L431:
	;
	if v3766 != 0 {
		v3306 = v3766
		v3314 = v3766
		v3315 = v3748
		v3316 = v3749
		v3318 = v3750
		v3319 = v3751
		goto L358
	} else {
		goto L432
	}
L432:
	;
	goto L359
L433:
	;
	v3826 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_CheckXLogRemoved(m, v3206, v3826)
	mBase = m.M
	v3841 = m.ExcPending
	if v3841 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L434
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_list_sort(m, v3784, int32(417))
	mBase = m.M
	v3857 = m.ExcPending
	if v3857 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L435
	}
L435:
	;
	if v3784 == int32(0) {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3876 = m.ExcPending
	if v3876 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L439
	}
L437:
	;
	goto L438
L438:
	;
	v3912 = *(*int32)(unsafe.Add(mBase, uint32(v3784)+12))
	v3913 = *(*int32)(unsafe.Add(mBase, uint32(v3912)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3928 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+264)) = v58 + int32(2140)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+260)) = v58 + int32(2144)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+256)) = v58 + int32(540)
	v3941 = F_sscanf(m, v3913, int32(_a_F_perform_base_backup_32), v58+int32(256))
	mBase = m.M
	v3942 = m.ExcPending
	if v3942 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L442
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errmsg(m, int32(_a_F_perform_base_backup_36), int32(0))
	mBase = m.M
	v3893 = m.ExcPending
	if v3893 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L440
	}
L440:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errfinish(m, int32(_a_F_perform_base_backup_30), int32(481), int32(_a_F_perform_base_backup_31))
	mBase = m.M
	v3911 = m.ExcPending
	if v3911 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L441
	}
L441:
	;
	goto L1
L442:
	;
	v3943 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+2140)))
	v3944 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+2144)))
	v3946 = base.I64_div_u_s(int64(4294967296), v3928)
	if v3206 == v3943+v3944*v3946 {
		goto L447
	} else {
		goto L448
	}
L443:
	;
	if v3785 == int32(0) {
		goto L536
	} else {
		goto L537
	}
L444:
	;
	if v4121 <= int32(0) {
		goto L443
	} else {
		goto L473
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v4255 = v58 + int32(336)
	v4257 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	F_XLogFileName(m, v4255, v3084, v3240, v4257)
	mBase = m.M
	v4259 = m.ExcPending
	if v4259 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L469
	}
L446:
	;
	v4033 = v3950
	v4068 = v3206
	goto L457
L447:
	;
	v3950 = int32(0)
	v3951 = *(*int32)(unsafe.Add(mBase, uint32(v3784)+4))
	if v3950 < v3951 {
		goto L446
	} else {
		goto L450
	}
L448:
	;
	goto L449
L449:
	;
	v3955 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v3970 = v58 + int32(464)
	v3972 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	F_XLogFileName(m, v3970, v3955, v3206, v3972)
	mBase = m.M
	v3974 = m.ExcPending
	if v3974 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L452
	}
L450:
	;
	if v3206 != v3240 {
		goto L445
	} else {
		goto L451
	}
L451:
	;
	goto L443
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3991 = m.ExcPending
	if v3991 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+240)) = v3970
	F_errmsg(m, int32(_a_F_perform_base_backup_37), v58+int32(240))
	mBase = m.M
	v4010 = m.ExcPending
	if v4010 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L454
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errfinish(m, int32(_a_F_perform_base_backup_30), int32(496), int32(_a_F_perform_base_backup_31))
	mBase = m.M
	v4028 = m.ExcPending
	if v4028 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L455
	}
L455:
	;
	goto L1
L456:
	;
	if v4113 == v3240 {
		goto L444
	} else {
		goto L468
	}
L457:
	;
	v4072 = *(*int32)(unsafe.Add(mBase, uint32(v3784)+12))
	v4076 = *(*int32)(unsafe.Add(mBase, uint32(v4072+v4033<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v4091 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+232)) = v58 + int32(2148)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+228)) = v58 + int32(2152)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+224)) = v58 + int32(540)
	v4104 = F_sscanf(m, v4076, int32(_a_F_perform_base_backup_32), v58+int32(224))
	mBase = m.M
	v4105 = m.ExcPending
	if v4105 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L459
	}
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v4137 = v58 + int32(400)
	v4138 = *(*int32)(unsafe.Add(mBase, uint32(v58)+540))
	v4140 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	F_XLogFileName(m, v4137, v4138, v4107, v4140)
	mBase = m.M
	v4142 = m.ExcPending
	if v4142 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L464
	}
L459:
	;
	v4107 = v4068 + int64(1)
	v4108 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+2148)))
	v4109 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+2152)))
	v4111 = base.I64_div_u_s(int64(4294967296), v4091)
	v4113 = v4108 + v4109*v4111
	if base.B2i32(v4107 != v4113)&base.B2i32(v4068 != v4113) == int32(0) {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v4120 = v4033 + int32(1)
	v4121 = *(*int32)(unsafe.Add(mBase, uint32(v3784)+4))
	if v4121 <= v4120 {
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
	v4033 = v4120
	v4068 = v4113
	goto L457
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4159 = m.ExcPending
	if v4159 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L465
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+208)) = v4137
	F_errmsg(m, int32(_a_F_perform_base_backup_37), v58+int32(208))
	mBase = m.M
	v4178 = m.ExcPending
	if v4178 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L466
	}
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errfinish(m, int32(_a_F_perform_base_backup_30), int32(511), int32(_a_F_perform_base_backup_31))
	mBase = m.M
	v4196 = m.ExcPending
	if v4196 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4276 = m.ExcPending
	if v4276 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L470
	}
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+192)) = v4255
	F_errmsg(m, int32(_a_F_perform_base_backup_37), v58+int32(192))
	mBase = m.M
	v4295 = m.ExcPending
	if v4295 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L471
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errfinish(m, int32(_a_F_perform_base_backup_30), int32(520), int32(_a_F_perform_base_backup_31))
	mBase = m.M
	v4313 = m.ExcPending
	if v4313 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L472
	}
L472:
	;
	goto L1
L473:
	;
	v4336 = int32(0)
	goto L474
L474:
	;
	v4360 = *(*int32)(unsafe.Add(mBase, uint32(v3784)+12))
	v4364 = *(*int32)(unsafe.Add(mBase, uint32(v4360+v4336<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+176)) = v4364
	v4380 = v58 + int32(768)
	v4385 = F_pg_snprintf(m, v4380, int32(1024), int32(_a_F_perform_base_backup_38), v58+int32(176))
	mBase = m.M
	v4386 = m.ExcPending
	if v4386 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L476
	}
L475:
	;
	goto L443
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v4401 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+168)) = v58 + int32(2156)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+164)) = v58 + int32(2160)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+160)) = v58 + int32(540)
	v4414 = F_sscanf(m, v4364, int32(_a_F_perform_base_backup_32), v58+int32(160))
	mBase = m.M
	v4415 = m.ExcPending
	if v4415 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L477
	}
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v4430 = base.I64_div_u_s(int64(4294967296), v4401)
	v4431 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+2160)))
	v4432 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+2156)))
	v4434 = F_OpenTransientFile(m, v4380, int32(0))
	mBase = m.M
	v4435 = m.ExcPending
	if v4435 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L478
	}
L478:
	;
	v4437 = v4430*v4431 + v4432
	if v4434 < int32(0) {
		goto L479
	} else {
		goto L480
	}
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v4454 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v4468 = *(*int32)(unsafe.Add(mBase, uint32(v58)+540))
	F_CheckXLogRemoved(m, v4437, v4468)
	mBase = m.M
	v4470 = m.ExcPending
	if v4470 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L482
	}
L480:
	;
	goto L481
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	if v4434 < int32(0) {
		goto L488
	} else {
		goto L489
	}
L482:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[18])) = v4454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4489 = m.ExcPending
	if v4489 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L483
	}
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errcode_for_file_access(m)
	mBase = m.M
	v4504 = m.ExcPending
	if v4504 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L484
	}
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v4380
	F_errmsg(m, int32(_a_F_perform_base_backup_39), v58)
	mBase = m.M
	v4521 = m.ExcPending
	if v4521 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L485
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errfinish(m, int32(_a_F_perform_base_backup_30), int32(549), int32(_a_F_perform_base_backup_31))
	mBase = m.M
	v4539 = m.ExcPending
	if v4539 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L486
	}
L486:
	;
	goto L1
L487:
	;
	if v4562 != 0 {
		goto L491
	} else {
		goto L492
	}
L488:
	;
	v4558 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v4562 = v4558
	goto L487
L489:
	;
	goto L490
L490:
	;
	v4561 = F___fstatat(m, v4434, int32(_a_F_perform_base_backup_40), v58+int32(672), int32(_a_F_perform_base_backup_41))
	mBase = m.M
	v4562 = v4561
	goto L487
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4579 = m.ExcPending
	if v4579 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L494
	}
L492:
	;
	goto L493
L493:
	;
	v4634 = *(*int64)(unsafe.Add(mBase, uint32(v58)+696))
	v4636 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	if v4634 != v4636 {
		goto L498
	} else {
		goto L499
	}
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errcode_for_file_access(m)
	mBase = m.M
	v4594 = m.ExcPending
	if v4594 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L495
	}
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+144)) = v58 + int32(768)
	F_errmsg(m, int32(_a_F_perform_base_backup_29), v58+int32(144))
	mBase = m.M
	v4615 = m.ExcPending
	if v4615 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L496
	}
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errfinish(m, int32(_a_F_perform_base_backup_30), int32(556), int32(_a_F_perform_base_backup_31))
	mBase = m.M
	v4633 = m.ExcPending
	if v4633 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L497
	}
L497:
	;
	goto L1
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v4651 = *(*int32)(unsafe.Add(mBase, uint32(v58)+540))
	F_CheckXLogRemoved(m, v4437, v4651)
	mBase = m.M
	v4653 = m.ExcPending
	if v4653 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L501
	}
L499:
	;
	goto L500
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v4738 = int32(0)
	F__tarWriteHeader(m, v56, v58+int32(768), v4738, v58+int32(672), v4738)
	mBase = m.M
	v4743 = m.ExcPending
	if v4743 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L506
	}
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4670 = m.ExcPending
	if v4670 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L502
	}
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errcode_for_file_access(m)
	mBase = m.M
	v4685 = m.ExcPending
	if v4685 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L503
	}
L503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+128)) = v4364
	F_errmsg(m, int32(_a_F_perform_base_backup_42), v58+int32(128))
	mBase = m.M
	v4704 = m.ExcPending
	if v4704 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L504
	}
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errfinish(m, int32(_a_F_perform_base_backup_30), int32(562), int32(_a_F_perform_base_backup_31))
	mBase = m.M
	v4722 = m.ExcPending
	if v4722 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L505
	}
L505:
	;
	goto L1
L506:
	;
	v4745 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	v4751 = v4745
	v4786 = int64(0)
	goto L508
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v5035 = F_CloseTransientFile(m, v4434)
	mBase = m.M
	v5036 = m.ExcPending
	if v5036 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L532
	}
L508:
	;
	v4790 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v4791 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v56)+8)))
	v4793 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v4793))) = int32(167772163)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v4810 = base.I64_extend_i32_s(v4751) - v4786
	if v4810 < v4791 {
		goto L510
	} else {
		goto L511
	}
L509:
	;
	v4932 = int64(*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14])))
	if v4786 == v4932 {
		goto L507
	} else {
		goto L526
	}
L510:
	;
	v4812 = v4810
	goto L512
L511:
	;
	v4812 = v4791
	goto L512
L512:
	;
	v4814 = F_pread(m, v4434, v4790, base.I32_wrap_i64(v4812), v4786)
	mBase = m.M
	v4816 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[19]))
	v4817 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4816))) = v4817
	if v4814 < v4817 {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4837 = m.ExcPending
	if v4837 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L516
	}
L514:
	;
	goto L515
L515:
	;
	if v4814 != 0 {
		goto L520
	} else {
		goto L521
	}
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errcode_for_file_access(m)
	mBase = m.M
	v4852 = m.ExcPending
	if v4852 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L517
	}
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v58 + int32(768)
	F_errmsg(m, int32(_a_F_perform_base_backup_43), v58+int32(16))
	mBase = m.M
	v4873 = m.ExcPending
	if v4873 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L518
	}
L518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errfinish(m, int32(_a_F_perform_base_backup_30), int32(2128), int32(_a_F_perform_base_backup_44))
	mBase = m.M
	v4891 = m.ExcPending
	if v4891 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L519
	}
L519:
	;
	goto L1
L520:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v4905 = *(*int32)(unsafe.Add(mBase, uint32(v58)+540))
	F_CheckXLogRemoved(m, v4437, v4905)
	mBase = m.M
	v4907 = m.ExcPending
	if v4907 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v4908 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v4909 = *(*int32)(unsafe.Add(mBase, uint32(v4908)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	m.T0[v4909].(func(*base.Module, int32, int32))(m, v56, v4814)
	mBase = m.M
	v4924 = m.ExcPending
	if v4924 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L524
	}
L524:
	;
	v4926 = v4786 + base.I64_extend_i32_u(v4814)
	v4928 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[14]))
	if v4926 != base.I64_extend_i32_s(v4928) {
		v4751 = v4928
		v4786 = v4926
		goto L508
	} else {
		goto L525
	}
L525:
	;
	goto L507
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v4947 = *(*int32)(unsafe.Add(mBase, uint32(v58)+540))
	F_CheckXLogRemoved(m, v4437, v4947)
	mBase = m.M
	v4949 = m.ExcPending
	if v4949 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L527
	}
L527:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4966 = m.ExcPending
	if v4966 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L528
	}
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errcode_for_file_access(m)
	mBase = m.M
	v4981 = m.ExcPending
	if v4981 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L529
	}
L529:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+112)) = v4364
	F_errmsg(m, int32(_a_F_perform_base_backup_42), v58+int32(112))
	mBase = m.M
	v5000 = m.ExcPending
	if v5000 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L530
	}
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errfinish(m, int32(_a_F_perform_base_backup_30), int32(587), int32(_a_F_perform_base_backup_31))
	mBase = m.M
	v5018 = m.ExcPending
	if v5018 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L531
	}
L531:
	;
	goto L1
L532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+100)) = int32(_a_F_perform_base_backup_45)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+96)) = v4364
	v5054 = v58 + int32(768)
	v5059 = F_pg_snprintf(m, v5054, int32(1024), int32(_a_F_perform_base_backup_46), v58+int32(96))
	mBase = m.M
	v5060 = m.ExcPending
	if v5060 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L533
	}
L533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_sendFileWithContent(m, v56, v5054, int32(_a_F_perform_base_backup_40), v58+int32(2056))
	mBase = m.M
	v5078 = m.ExcPending
	if v5078 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L534
	}
L534:
	;
	v5080 = v4336 + int32(1)
	v5081 = *(*int32)(unsafe.Add(mBase, uint32(v3784)+4))
	if v5080 < v5081 {
		v4336 = v5080
		goto L474
	} else {
		goto L535
	}
L535:
	;
	goto L475
L536:
	;
	v5403 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v5405 = int32(1024)
	base.MemoryFill(m, v5403, int32(0), v5405)
	v5407 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v5408 = *(*int32)(unsafe.Add(mBase, uint32(v5407)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	m.T0[v5408].(func(*base.Module, int32, int32))(m, v56, v5405)
	mBase = m.M
	v5424 = m.ExcPending
	if v5424 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L554
	}
L537:
	;
	v5128 = int32(0)
	v5129 = *(*int32)(unsafe.Add(mBase, uint32(v3785)+4))
	if v5129 <= v5128 {
		goto L536
	} else {
		goto L538
	}
L538:
	;
	v5136 = v5128
	goto L539
L539:
	;
	v5175 = *(*int32)(unsafe.Add(mBase, uint32(v3785)+12))
	v5179 = *(*int32)(unsafe.Add(mBase, uint32(v5175+v5136<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+80)) = v5179
	v5195 = v58 + int32(768)
	v5200 = F_pg_snprintf(m, v5195, int32(1024), int32(_a_F_perform_base_backup_38), v58+int32(80))
	mBase = m.M
	v5201 = m.ExcPending
	if v5201 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L541
	}
L540:
	;
	goto L536
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v5219 = F___fstatat(m, int32(-100), v5195, v58+int32(672), int32(256))
	mBase = m.M
	goto L542
L542:
	;
	if v5219 != 0 {
		goto L543
	} else {
		goto L544
	}
L543:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5236 = m.ExcPending
	if v5236 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L546
	}
L544:
	;
	goto L545
L545:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v5303 = v58 + int32(768)
	v5306 = int32(0)
	v5312 = v58 + int32(2056)
	v5316 = F_sendFile(m, v56, v5303, v5303, v58+int32(672), v5306, v5306, v5306, v5306, v5306, v5312, v5306, v5306, v5306)
	mBase = m.M
	v5317 = m.ExcPending
	if v5317 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L550
	}
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errcode_for_file_access(m)
	mBase = m.M
	v5251 = m.ExcPending
	if v5251 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L547
	}
L547:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+64)) = v5195
	F_errmsg(m, int32(_a_F_perform_base_backup_29), v58-int32(-64))
	mBase = m.M
	v5270 = m.ExcPending
	if v5270 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L548
	}
L548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errfinish(m, int32(_a_F_perform_base_backup_30), int32(626), int32(_a_F_perform_base_backup_31))
	mBase = m.M
	v5288 = m.ExcPending
	if v5288 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L549
	}
L549:
	;
	goto L1
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v58)+52)) = int32(_a_F_perform_base_backup_45)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+48)) = v5179
	v5338 = F_pg_snprintf(m, v5303, int32(1024), int32(_a_F_perform_base_backup_46), v58+int32(48))
	mBase = m.M
	v5339 = m.ExcPending
	if v5339 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L551
	}
L551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_sendFileWithContent(m, v56, v5303, int32(_a_F_perform_base_backup_40), v5312)
	mBase = m.M
	v5355 = m.ExcPending
	if v5355 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L552
	}
L552:
	;
	v5357 = v5136 + int32(1)
	v5358 = *(*int32)(unsafe.Add(mBase, uint32(v3785)+4))
	if v5357 < v5358 {
		v5136 = v5357
		goto L539
	} else {
		goto L553
	}
L553:
	;
	goto L540
L554:
	;
	v5425 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v5426 = *(*int32)(unsafe.Add(mBase, uint32(v5425)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v3781
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v3780
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v3782
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	m.T0[v5426].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v5441 = m.ExcPending
	if v5441 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L555
	}
L555:
	;
	v5454 = v3780
	v5455 = v3781
	v5456 = v3782
	goto L345
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5456
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v5780 = m.G0
	v5782 = v5780 - int32(128)
	m.G0 = v5782
	v5785 = v58 + int32(2056)
	v5786 = *(*int32)(unsafe.Add(mBase, uint32(v5785)))
	if v5786 != 0 {
		goto L600
	} else {
		goto L601
	}
L557:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5751 = m.ExcPending
	if v5751 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L594
	}
L558:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5734 = m.ExcPending
	if v5734 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L591
	}
L559:
	;
	F_AppendStringToManifest(m, v5506, int32(_a_F_perform_base_backup_47))
	mBase = m.M
	v5510 = m.ExcPending
	if v5510 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L562
	}
L560:
	;
	goto L561
L561:
	;
	m.G0 = v5503 + int32(80)
	goto L556
L562:
	;
	v5511 = F_readTimeLineHistory(m, v3084)
	mBase = m.M
	v5512 = m.ExcPending
	if v5512 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L563
	}
L563:
	;
	F_AppendStringToManifest(m, v5506, int32(_a_F_perform_base_backup_48))
	mBase = m.M
	v5515 = m.ExcPending
	if v5515 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L564
	}
L564:
	;
	if v5511 == int32(0) {
		goto L566
	} else {
		goto L567
	}
L565:
	;
	F_AppendStringToManifest(m, v5506, int32(_a_F_perform_base_backup_47))
	mBase = m.M
	v5684 = m.ExcPending
	if v5684 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L590
	}
L566:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5671 = m.ExcPending
	if v5671 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L587
	}
L567:
	;
	v5519 = *(*int32)(unsafe.Add(mBase, uint32(v5511)+4))
	if v5519 <= int32(0) {
		goto L566
	} else {
		goto L568
	}
L568:
	;
	v5544 = int32(1)
	v5545 = int32(0)
	v5564 = v3085
	goto L569
L569:
	;
	v5569 = *(*int32)(unsafe.Add(mBase, uint32(v5511)+12))
	v5573 = *(*int32)(unsafe.Add(mBase, uint32(v5569+v5545<<(uint(int32(2))%32))))
	v5574 = *(*int64)(unsafe.Add(mBase, uint32(v5573)+16))
	if base.B2i32(v5574 != int64(0))&base.B2i32(base.Ui64(v5574) < base.Ui64(v5486)) == int32(0) {
		goto L571
	} else {
		goto L572
	}
L570:
	;
	goto L566
L571:
	;
	v5581 = *(*int32)(unsafe.Add(mBase, uint32(v5573)))
	if base.B2i32(v3084 != v5581)&v5544 != 0 {
		goto L558
	} else {
		goto L574
	}
L572:
	;
	v5617 = v5544
	v5619 = v5564
	goto L573
L573:
	;
	v5622 = v5545 + int32(1)
	v5623 = *(*int32)(unsafe.Add(mBase, uint32(v5511)+4))
	if v5622 < v5623 {
		v5544 = v5617
		v5545 = v5622
		v5564 = v5619
		goto L569
	} else {
		goto L586
	}
L574:
	;
	if v5485 != v5581 {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	v5585 = *(*int64)(unsafe.Add(mBase, uint32(v5573)+8))
	if v5585 == int64(0) {
		goto L557
	} else {
		goto L578
	}
L576:
	;
	v5588 = v5486
	goto L577
L577:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v5503+int32(52)))) = uint32(v5564)
	v5590 = int64(32)
	v5591 = int64(base.Ui64(v5564) >> (uint(v5590) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v5503+int32(48)))) = uint32(v5591)
	*(*int32)(unsafe.Add(mBase, uint32(v5503)+36)) = v5581
	*(*uint32)(unsafe.Add(mBase, uint32(v5503)+44)) = uint32(v5588)
	v5596 = int64(base.Ui64(v5588) >> (uint(v5590) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v5503)+40)) = uint32(v5596)
	if v5544&int32(1) != 0 {
		goto L579
	} else {
		goto L580
	}
L578:
	;
	v5588 = v5585
	goto L577
L579:
	;
	v5602 = int32(_a_F_perform_base_backup_40)
	goto L581
L580:
	;
	v5602 = int32(_a_F_perform_base_backup_49)
	goto L581
L581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5503)+32)) = v5602
	v5607 = F_psprintf(m, int32(_a_F_perform_base_backup_50), v5503+int32(32))
	mBase = m.M
	v5608 = m.ExcPending
	if v5608 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L582
	}
L582:
	;
	F_AppendStringToManifest(m, v5506, v5607)
	mBase = m.M
	v5610 = m.ExcPending
	if v5610 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L583
	}
L583:
	;
	F_pfree(m, v5607)
	mBase = m.M
	v5612 = m.ExcPending
	if v5612 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L584
	}
L584:
	;
	v5613 = *(*int32)(unsafe.Add(mBase, uint32(v5573)))
	if v5485 == v5613 {
		goto L565
	} else {
		goto L585
	}
L585:
	;
	v5616 = *(*int64)(unsafe.Add(mBase, uint32(v5573)+8))
	v5617 = int32(0)
	v5619 = v5616
	goto L573
L586:
	;
	goto L570
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5503)+4)) = v3084
	*(*int32)(unsafe.Add(mBase, uint32(v5503))) = v5485
	F_errmsg(m, int32(_a_F_perform_base_backup_51), v5503)
	mBase = m.M
	v5676 = m.ExcPending
	if v5676 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L588
	}
L588:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(307), int32(_a_F_perform_base_backup_52))
	mBase = m.M
	v5681 = m.ExcPending
	if v5681 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v5735 = *(*int32)(unsafe.Add(mBase, uint32(v5573)))
	*(*int32)(unsafe.Add(mBase, uint32(v5503)+20)) = v5735
	*(*int32)(unsafe.Add(mBase, uint32(v5503)+16)) = v3084
	F_errmsg(m, int32(_a_F_perform_base_backup_53), v5503+int32(16))
	mBase = m.M
	v5742 = m.ExcPending
	if v5742 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L592
	}
L592:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(256), int32(_a_F_perform_base_backup_52))
	mBase = m.M
	v5747 = m.ExcPending
	if v5747 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v5752 = *(*int32)(unsafe.Add(mBase, uint32(v5573)))
	*(*int32)(unsafe.Add(mBase, uint32(v5503)+68)) = v5752
	*(*int32)(unsafe.Add(mBase, uint32(v5503)+64)) = v5485
	F_errmsg(m, int32(_a_F_perform_base_backup_54), v5503-int32(-64))
	mBase = m.M
	v5759 = m.ExcPending
	if v5759 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L595
	}
L595:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(280), int32(_a_F_perform_base_backup_52))
	mBase = m.M
	v5764 = m.ExcPending
	if v5764 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v6045 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v6046 = *(*int32)(unsafe.Add(mBase, uint32(v6045)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5456
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	m.T0[v6046].(func(*base.Module, int32, int64, int32))(m, v56, v3085, v3084)
	mBase = m.M
	v6061 = m.ExcPending
	if v6061 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L647
	}
L598:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6033 = m.ExcPending
	if v6033 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L643
	}
L599:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6004 = m.ExcPending
	if v6004 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L630
	}
L600:
	;
	v5787 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5785)+26)) = uint8(v5787)
	v5789 = *(*int32)(unsafe.Add(mBase, uint32(v5785)+8))
	v5791 = v5782 + int32(96)
	v5793 = F_pg_cryptohash_final(m, v5789, v5791, int32(32))
	mBase = m.M
	if v5793 < v5787 {
		goto L599
	} else {
		goto L603
	}
L601:
	;
	goto L602
L602:
	;
	m.G0 = v5782 + int32(128)
	goto L597
L603:
	;
	F_AppendStringToManifest(m, v5785, int32(_a_F_perform_base_backup_55))
	mBase = m.M
	v5798 = m.ExcPending
	if v5798 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L604
	}
L604:
	;
	v5801 = v5782 + int32(16)
	goto L606
L605:
	;
	v5825 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5782)+80)) = uint8(v5825)
	F_AppendStringToManifest(m, v5785, v5801)
	mBase = m.M
	v5828 = m.ExcPending
	if v5828 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L612
	}
L606:
	;
	v5804 = v5791
	v5806 = v5801
	goto L609
L608:
	;
	goto L605
L609:
	;
	v5808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5804))))
	v5809 = int32(1)
	v5811 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5808<<(uint(v5809)%32))+uint32(_c_F_perform_base_backup[20]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5806))) = uint16(v5811)
	v5816 = v5804 + v5809
	if base.Ui32(v5816) < base.Ui32(v5782+int32(128)) {
		v5804 = v5816
		v5806 = v5806 + int32(2)
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
	F_AppendStringToManifest(m, v5785, int32(_a_F_perform_base_backup_56))
	mBase = m.M
	v5831 = m.ExcPending
	if v5831 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L613
	}
L613:
	;
	v5832 = *(*int32)(unsafe.Add(mBase, uint32(v5785)))
	v5833 = int32(0)
	v5836 = F_BufFileSeek(m, v5832, v5833, int64(0), v5833)
	mBase = m.M
	v5837 = m.ExcPending
	if v5837 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L614
	}
L614:
	;
	if v5836 != 0 {
		goto L598
	} else {
		goto L615
	}
L615:
	;
	v5838 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v5839 = *(*int32)(unsafe.Add(mBase, uint32(v5838)+16))
	m.T0[v5839].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v5841 = m.ExcPending
	if v5841 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L616
	}
L616:
	;
	v5842 = *(*int64)(unsafe.Add(mBase, uint32(v5785)+16))
	if v5842 != int64(0) {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	v5863 = int32(0)
	v5884 = v5842
	v5885 = int64(0)
	goto L620
L618:
	;
	goto L619
L619:
	;
	v5948 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v5949 = *(*int32)(unsafe.Add(mBase, uint32(v5948)+24))
	m.T0[v5949].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v5951 = m.ExcPending
	if v5951 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L628
	}
L620:
	;
	v5888 = *(*int32)(unsafe.Add(mBase, uint32(v5785)))
	v5889 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v5890 = v5884 - v5885
	v5891 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v56)+8)))
	if base.Ui64(v5890) < base.Ui64(v5891) {
		goto L622
	} else {
		goto L623
	}
L621:
	;
	goto L619
L622:
	;
	v5893 = v5890
	goto L624
L623:
	;
	v5893 = v5891
	goto L624
L624:
	;
	v5894 = base.I32_wrap_i64(v5893)
	F_BufFileReadExact(m, v5888, v5889, v5894)
	mBase = m.M
	v5896 = m.ExcPending
	if v5896 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L625
	}
L625:
	;
	v5897 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v5898 = *(*int32)(unsafe.Add(mBase, uint32(v5897)+20))
	m.T0[v5898].(func(*base.Module, int32, int32))(m, v56, v5894)
	mBase = m.M
	v5900 = m.ExcPending
	if v5900 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L626
	}
L626:
	;
	v5901 = *(*int64)(unsafe.Add(mBase, uint32(v5785)+16))
	v5902 = v5894 + v5863
	v5903 = base.I64_extend_i32_u(v5902)
	if base.Ui64(v5903) < base.Ui64(v5901) {
		v5863 = v5902
		v5884 = v5901
		v5885 = v5903
		goto L620
	} else {
		goto L627
	}
L627:
	;
	goto L621
L628:
	;
	v5952 = *(*int32)(unsafe.Add(mBase, uint32(v5785)))
	F_BufFileClose(m, v5952)
	mBase = m.M
	v5954 = m.ExcPending
	if v5954 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L629
	}
L629:
	;
	goto L602
L630:
	;
	v6005 = *(*int32)(unsafe.Add(mBase, uint32(v5785)+8))
	if v6005 == int32(0) {
		goto L632
	} else {
		goto L633
	}
L631:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5782))) = v6020
	F_errmsg_internal(m, int32(_a_F_perform_base_backup_57), v5782)
	mBase = m.M
	v6024 = m.ExcPending
	if v6024 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L641
	}
L632:
	;
	v6020 = int32(_a_F_perform_base_backup_2)
	goto L631
L633:
	;
	goto L634
L634:
	;
	v6012 = *(*int32)(unsafe.Add(mBase, uint32(v6005)+4))
	if v6012 == int32(1) {
		goto L635
	} else {
		goto L636
	}
L635:
	;
	v6015 = int32(_a_F_perform_base_backup_3)
	goto L637
L636:
	;
	v6015 = int32(_a_F_perform_base_backup_4)
	goto L637
L637:
	;
	if v6012 == int32(2) {
		goto L638
	} else {
		goto L639
	}
L638:
	;
	v6018 = int32(_a_F_perform_base_backup_2)
	goto L640
L639:
	;
	v6018 = v6015
	goto L640
L640:
	;
	v6020 = v6018
	goto L631
L641:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(341), int32(_a_F_perform_base_backup_58))
	mBase = m.M
	v6029 = m.ExcPending
	if v6029 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v6035 = m.ExcPending
	if v6035 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L644
	}
L644:
	;
	F_errmsg(m, int32(_a_F_perform_base_backup_59), int32(0))
	mBase = m.M
	v6039 = m.ExcPending
	if v6039 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L645
	}
L645:
	;
	F_errfinish(m, int32(_a_F_perform_base_backup_5), int32(357), int32(_a_F_perform_base_backup_58))
	mBase = m.M
	v6044 = m.ExcPending
	if v6044 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
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
	v6063 = *(*int64)(unsafe.Add(mBase, _c_F_perform_base_backup[5]))
	if v6063 != int64(0) {
		goto L648
	} else {
		goto L649
	}
L648:
	;
	if v6063 < int64(2) {
		goto L651
	} else {
		goto L652
	}
L649:
	;
	goto L650
L650:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5456
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v6211 = v58 + int32(2056)
	v6212 = *(*int32)(unsafe.Add(mBase, uint32(v6211)+8))
	F_pg_cryptohash_free(m, v6212)
	mBase = m.M
	v6214 = m.ExcPending
	if v6214 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L661
	}
L651:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5456
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6145 = m.ExcPending
	if v6145 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L657
	}
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5456
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v6083 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v6084 = m.ExcPending
	if v6084 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L653
	}
L653:
	;
	if v6083 == int32(0) {
		goto L651
	} else {
		goto L654
	}
L654:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5456
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v6101 = *(*int64)(unsafe.Add(mBase, _c_F_perform_base_backup[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+32)) = v6101
	F_errmsg_plural(m, int32(_a_F_perform_base_backup_60), int32(_a_F_perform_base_backup_61), base.I32_wrap_i64(v6101), v58+int32(32))
	mBase = m.M
	v6109 = m.ExcPending
	if v6109 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L655
	}
L655:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5456
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errfinish(m, int32(_a_F_perform_base_backup_30), int32(661), int32(_a_F_perform_base_backup_31))
	mBase = m.M
	v6127 = m.ExcPending
	if v6127 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L656
	}
L656:
	;
	goto L651
L657:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5456
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errcode(m, int32(16779816))
	mBase = m.M
	v6161 = m.ExcPending
	if v6161 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L658
	}
L658:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5456
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errmsg(m, int32(_a_F_perform_base_backup_62), int32(0))
	mBase = m.M
	v6178 = m.ExcPending
	if v6178 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L659
	}
L659:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5456
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_errfinish(m, int32(_a_F_perform_base_backup_30), int32(665), int32(_a_F_perform_base_backup_31))
	mBase = m.M
	v6196 = m.ExcPending
	if v6196 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L660
	}
L660:
	;
	goto L1
L661:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6211)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5456
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	F_ReleaseAuxProcessResources(m, int32(1))
	mBase = m.M
	v6232 = m.ExcPending
	if v6232 != 0 {
		v6284 = v55
		v6285 = v56
		v6286 = v57
		v6287 = v58
		v6311 = v82
		v6315 = v86
		v6316 = v87
		goto L6
	} else {
		goto L662
	}
L662:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2168)) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2164)) = v5454
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2172)) = v5456
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2176)) = v2334
	*(*int64)(unsafe.Add(mBase, uint32(v58)+2184)) = v2335
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2192)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2196)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2200)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2204)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2208)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2212)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2216)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2220)) = v384
	v6248 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[7]))
	if v6248 == int32(0) {
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
	v6252 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_perform_base_backup[8])))
	if v6252&int32(1) == int32(0) {
		goto L664
	} else {
		goto L666
	}
L666:
	;
	v6257 = *(*int32)(unsafe.Add(mBase, uint32(v6248)+220))
	if v6257 == int32(0) {
		goto L664
	} else {
		goto L667
	}
L667:
	;
	v6260 = int32(_a_F_perform_base_backup_7)
	v6262 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	v6263 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v6262 + v6263
	v6266 = *(*int32)(unsafe.Add(mBase, uint32(v6248)))
	*(*int32)(unsafe.Add(mBase, uint32(v6248))) = v6266 + v6263
	v6270 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6248)+220)) = v6270
	*(*int32)(unsafe.Add(mBase, uint32(v6248)+224)) = v6270
	*(*int32)(unsafe.Add(mBase, uint32(v6248))) = v6266 + int32(2)
	v6280 = *(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_base_backup[9])) = v6280 - v6263
	goto L664
L668:
	;
	v6332 = int32(v6328)
	m.G0 = v6287
	v6334 = *(*int32)(unsafe.Add(mBase, uint32(v6332)+4))
	v6335 = *(*int32)(unsafe.Add(mBase, uint32(v6332)))
	v6338 = *(*int32)(unsafe.Add(mBase, uint32(v6335)))
	if v6287+int32(332) == v6338 {
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
	if v6342 != 0 {
		goto L674
	} else {
		goto L675
	}
L671:
	;
	v6340 = *(*int32)(unsafe.Add(mBase, uint32(v6335)+4))
	v6342 = v6340
	goto L673
L672:
	;
	v6342 = int32(0)
	goto L673
L673:
	;
	goto L670
L674:
	;
	v6343 = *(*int32)(unsafe.Add(mBase, uint32(v6287)+2220))
	v6344 = *(*int32)(unsafe.Add(mBase, uint32(v6287)+2216))
	v6345 = *(*int32)(unsafe.Add(mBase, uint32(v6287)+2212))
	v6346 = *(*int32)(unsafe.Add(mBase, uint32(v6287)+2208))
	v6347 = *(*int32)(unsafe.Add(mBase, uint32(v6287)+2204))
	v6348 = *(*int32)(unsafe.Add(mBase, uint32(v6287)+2200))
	v6349 = *(*int32)(unsafe.Add(mBase, uint32(v6287)+2196))
	v6350 = *(*int32)(unsafe.Add(mBase, uint32(v6287)+2192))
	v6351 = *(*int64)(unsafe.Add(mBase, uint32(v6287)+2184))
	v6352 = *(*int64)(unsafe.Add(mBase, uint32(v6287)+2176))
	v6353 = *(*int32)(unsafe.Add(mBase, uint32(v6287)+2172))
	v6354 = *(*int32)(unsafe.Add(mBase, uint32(v6287)+2168))
	v6355 = *(*int32)(unsafe.Add(mBase, uint32(v6287)+2164))
	v55 = v6284
	v56 = v6285
	v57 = v6286
	v58 = v6287
	v59 = v6334
	v60 = v6348
	v61 = v6345
	v62 = v6343
	v63 = v6346
	v64 = v6347
	v65 = v6350
	v66 = v6349
	v67 = v6355
	v68 = v6354
	v69 = v6353
	v70 = v6344
	v73 = v6342
	v82 = v6311
	v86 = v6315
	v87 = v6316
	v91 = v6352
	v92 = v6351
	goto L2
L675:
	;
	goto L676
L676:
	;
	F___wasm_longjmp(m, v6335, v6334)
	mBase = m.M
	v6357 = m.ExcPending
	if v6357 != 0 {
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
