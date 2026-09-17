package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_do_autovacuum(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v149 int32
	_ = v149
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v228 int32
	_ = v228
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v822 int32
	_ = v822
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v947 int32
	_ = v947
	var v971 int32
	_ = v971
	var v977 int32
	_ = v977
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1152 int32
	_ = v1152
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1211 int32
	_ = v1211
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1237 int32
	_ = v1237
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1271 int32
	_ = v1271
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1303 int32
	_ = v1303
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1381 int32
	_ = v1381
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1425 int32
	_ = v1425
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1477 int32
	_ = v1477
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1529 int32
	_ = v1529
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1731 int32
	_ = v1731
	var v1735 int32
	_ = v1735
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1777 int32
	_ = v1777
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1823 int32
	_ = v1823
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1886 int32
	_ = v1886
	var v1913 int32
	_ = v1913
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1963 int32
	_ = v1963
	var v1996 int32
	_ = v1996
	var v2020 int32
	_ = v2020
	var v2044 int32
	_ = v2044
	var v2068 int32
	_ = v2068
	var v2071 int32
	_ = v2071
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2141 int32
	_ = v2141
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2168 int32
	_ = v2168
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2189 int32
	_ = v2189
	var v2193 int32
	_ = v2193
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2222 int32
	_ = v2222
	var v2229 int32
	_ = v2229
	var v2234 int32
	_ = v2234
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2263 int32
	_ = v2263
	var v2274 int32
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2280 int32
	_ = v2280
	var v2283 int32
	_ = v2283
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2310 int32
	_ = v2310
	var v2316 int32
	_ = v2316
	var v2318 int32
	_ = v2318
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2342 int32
	_ = v2342
	var v2350 int32
	_ = v2350
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2372 int32
	_ = v2372
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2385 int32
	_ = v2385
	var v2409 int32
	_ = v2409
	var v2429 int32
	_ = v2429
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2458 int32
	_ = v2458
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2474 int32
	_ = v2474
	var v2477 int32
	_ = v2477
	var v2479 int32
	_ = v2479
	var v2481 int32
	_ = v2481
	var v2521 int32
	_ = v2521
	var v2524 int32
	_ = v2524
	var v2526 int32
	_ = v2526
	var v2528 int32
	_ = v2528
	var v2589 int32
	_ = v2589
	var v2597 int32
	_ = v2597
	var v2599 int32
	_ = v2599
	var v2621 int32
	_ = v2621
	var v2629 int32
	_ = v2629
	var v2632 int32
	_ = v2632
	var v2659 int32
	_ = v2659
	var v2660 int32
	_ = v2660
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2745 int32
	_ = v2745
	var v2765 int32
	_ = v2765
	var v2773 int32
	_ = v2773
	var v2790 int32
	_ = v2790
	var v2792 int32
	_ = v2792
	var v2795 int32
	_ = v2795
	var v2798 int32
	_ = v2798
	var v2806 int32
	_ = v2806
	var v2807 int32
	_ = v2807
	var v2810 int32
	_ = v2810
	var v2843 int32
	_ = v2843
	var v2844 int32
	_ = v2844
	var v2845 int32
	_ = v2845
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2853 int32
	_ = v2853
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2858 int32
	_ = v2858
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2912 int32
	_ = v2912
	var v2936 int32
	_ = v2936
	var v2937 int32
	_ = v2937
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2944 int32
	_ = v2944
	var v2945 int32
	_ = v2945
	var v2950 int32
	_ = v2950
	var v2955 int32
	_ = v2955
	var v2959 int32
	_ = v2959
	var v2961 int32
	_ = v2961
	var v2963 int32
	_ = v2963
	var v2965 int32
	_ = v2965
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2971 int32
	_ = v2971
	var v2973 int32
	_ = v2973
	var v2974 int32
	_ = v2974
	var v2977 int32
	_ = v2977
	var v2979 int32
	_ = v2979
	var v2980 int32
	_ = v2980
	var v2983 int32
	_ = v2983
	var v2985 int32
	_ = v2985
	var v2986 int32
	_ = v2986
	var v2989 int32
	_ = v2989
	var v2990 int32
	_ = v2990
	var v2993 int32
	_ = v2993
	var v2994 int32
	_ = v2994
	var v2995 int32
	_ = v2995
	var v2996 int32
	_ = v2996
	var v2998 int32
	_ = v2998
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3004 int32
	_ = v3004
	var v3005 int32
	_ = v3005
	var v3006 int32
	_ = v3006
	var v3030 int32
	_ = v3030
	var v3031 int32
	_ = v3031
	var v3032 int32
	_ = v3032
	var v3034 int32
	_ = v3034
	var v3036 int32
	_ = v3036
	var v3041 int32
	_ = v3041
	var v3049 int32
	_ = v3049
	var v3057 int32
	_ = v3057
	var v3061 int32
	_ = v3061
	var v3065 float64
	_ = v3065
	var v3069 int32
	_ = v3069
	var v3078 int32
	_ = v3078
	var v3080 float64
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3087 int32
	_ = v3087
	var v3090 float64
	_ = v3090
	var v3096 int32
	_ = v3096
	var v3098 int32
	_ = v3098
	var v3129 int32
	_ = v3129
	var v3153 int32
	_ = v3153
	var v3155 int32
	_ = v3155
	var v3156 int32
	_ = v3156
	var v3186 int32
	_ = v3186
	var v3194 int32
	_ = v3194
	var v3195 int32
	_ = v3195
	var v3197 int32
	_ = v3197
	var v3198 int32
	_ = v3198
	var v3221 int32
	_ = v3221
	var v3229 int32
	_ = v3229
	var v3231 float64
	_ = v3231
	var v3234 int32
	_ = v3234
	var v3237 int32
	_ = v3237
	var v3238 int32
	_ = v3238
	var v3264 int32
	_ = v3264
	var v3272 int32
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3277 int32
	_ = v3277
	var v3278 int32
	_ = v3278
	var v3282 int32
	_ = v3282
	var v3285 int32
	_ = v3285
	var v3310 int32
	_ = v3310
	var v3324 int32
	_ = v3324
	var v3325 int32
	_ = v3325
	var v3329 int32
	_ = v3329
	var v3330 int32
	_ = v3330
	var v3358 int32
	_ = v3358
	var v3393 int32
	_ = v3393
	var v3401 int32
	_ = v3401
	var v3425 int32
	_ = v3425
	var v3445 int32
	_ = v3445
	var v3451 int32
	_ = v3451
	var v3452 int32
	_ = v3452
	var v3475 int32
	_ = v3475
	var v3476 int32
	_ = v3476
	var v3478 int32
	_ = v3478
	var v3485 int32
	_ = v3485
	var v3503 int32
	_ = v3503
	var v3504 int32
	_ = v3504
	var v3527 int32
	_ = v3527
	var v3528 int32
	_ = v3528
	var v3535 int32
	_ = v3535
	var v3551 int32
	_ = v3551
	var v3556 int32
	_ = v3556
	var v3557 int32
	_ = v3557
	var v3560 int32
	_ = v3560
	var v3561 int32
	_ = v3561
	var v3564 int32
	_ = v3564
	var v3571 int32
	_ = v3571
	var v3573 int32
	_ = v3573
	var v3575 int32
	_ = v3575
	var v3583 int32
	_ = v3583
	var v3588 int32
	_ = v3588
	var v3594 int32
	_ = v3594
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3618 int32
	_ = v3618
	var v3627 int32
	_ = v3627
	var v3635 int32
	_ = v3635
	var v3636 int32
	_ = v3636
	var v3654 int32
	_ = v3654
	var v3655 int32
	_ = v3655
	var v3658 int32
	_ = v3658
	var v3668 int32
	_ = v3668
	var v3669 int32
	_ = v3669
	var v3687 int32
	_ = v3687
	var v3688 int32
	_ = v3688
	var v3691 int32
	_ = v3691
	var v3697 int32
	_ = v3697
	var v3698 int32
	_ = v3698
	var v3699 int32
	_ = v3699
	var v3700 int32
	_ = v3700
	var v3701 int32
	_ = v3701
	var v3726 int32
	_ = v3726
	var v3736 int32
	_ = v3736
	var v3737 int32
	_ = v3737
	var v3761 int32
	_ = v3761
	var v3765 int64
	_ = v3765
	var v3809 int32
	_ = v3809
	var v3819 int32
	_ = v3819
	var v3820 int32
	_ = v3820
	var v3821 int32
	_ = v3821
	var v3822 int32
	_ = v3822
	var v3825 int32
	_ = v3825
	var v3826 int32
	_ = v3826
	var v3850 int32
	_ = v3850
	var v3851 int32
	_ = v3851
	var v3852 int32
	_ = v3852
	var v3876 int32
	_ = v3876
	var v3877 int32
	_ = v3877
	var v3906 int32
	_ = v3906
	var v3907 int32
	_ = v3907
	var v3934 int32
	_ = v3934
	var v3958 int32
	_ = v3958
	var v3966 int32
	_ = v3966
	var v3968 int32
	_ = v3968
	var v3969 int32
	_ = v3969
	var v3972 int32
	_ = v3972
	var v3981 int32
	_ = v3981
	var v3985 int32
	_ = v3985
	var v4001 int32
	_ = v4001
	var v4002 int32
	_ = v4002
	var v4003 int32
	_ = v4003
	var v4004 int32
	_ = v4004
	var v4034 int32
	_ = v4034
	var v4038 int32
	_ = v4038
	var v4062 int32
	_ = v4062
	var v4086 int32
	_ = v4086
	var v4110 int32
	_ = v4110
	var v4130 int32
	_ = v4130
	var v4136 int32
	_ = v4136
	var v4160 int32
	_ = v4160
	var v4161 int32
	_ = v4161
	var v4163 int32
	_ = v4163
	var v4181 int32
	_ = v4181
	var v4184 int32
	_ = v4184
	var v4188 int32
	_ = v4188
	var v4189 int32
	_ = v4189
	var v4191 int32
	_ = v4191
	var v4196 int32
	_ = v4196
	var v4210 int32
	_ = v4210
	var v4222 int32
	_ = v4222
	var v4225 int32
	_ = v4225
	var v4243 int32
	_ = v4243
	var v4244 int32
	_ = v4244
	var v4247 int32
	_ = v4247
	var v4253 int32
	_ = v4253
	var v4254 int32
	_ = v4254
	var v4272 int32
	_ = v4272
	var v4273 int32
	_ = v4273
	var v4276 int32
	_ = v4276
	var v4282 int32
	_ = v4282
	var v4283 int32
	_ = v4283
	var v4301 int32
	_ = v4301
	var v4302 int32
	_ = v4302
	var v4305 int32
	_ = v4305
	var v4311 int32
	_ = v4311
	var v4329 int32
	_ = v4329
	var v4330 int32
	_ = v4330
	var v4333 int32
	_ = v4333
	var v4339 int32
	_ = v4339
	var v4359 int32
	_ = v4359
	var v4367 int32
	_ = v4367
	var v4368 int32
	_ = v4368
	var v4370 int32
	_ = v4370
	var v4371 int32
	_ = v4371
	var v4394 int32
	_ = v4394
	var v4402 int32
	_ = v4402
	var v4404 int32
	_ = v4404
	var v4410 int32
	_ = v4410
	var v4411 int32
	_ = v4411
	var v4413 int32
	_ = v4413
	var v4418 int32
	_ = v4418
	var v4428 int32
	_ = v4428
	var v4432 int32
	_ = v4432
	var v4444 int32
	_ = v4444
	var v4450 int32
	_ = v4450
	var v4451 int32
	_ = v4451
	var v4453 int32
	_ = v4453
	var v4458 int32
	_ = v4458
	var v4467 int32
	_ = v4467
	var v4468 int32
	_ = v4468
	var v4472 int32
	_ = v4472
	var v4484 int32
	_ = v4484
	var v4488 int32
	_ = v4488
	var v4489 int32
	_ = v4489
	var v4496 int32
	_ = v4496
	var v4498 int32
	_ = v4498
	var v4499 int32
	_ = v4499
	var v4500 int32
	_ = v4500
	var v4501 int32
	_ = v4501
	var v4502 int32
	_ = v4502
	var v4503 int32
	_ = v4503
	var v4504 int32
	_ = v4504
	var v4505 int32
	_ = v4505
	var v4506 int32
	_ = v4506
	var v4507 int32
	_ = v4507
	var v4508 int32
	_ = v4508
	var v4509 int32
	_ = v4509
	var v4510 int32
	_ = v4510
	var v4511 int32
	_ = v4511
	var v4512 int32
	_ = v4512
	var v4513 int32
	_ = v4513
	var v4514 int32
	_ = v4514
	var v4516 int32
	_ = v4516
	var v4517 int32
	_ = v4517
	var v4520 int32
	_ = v4520
	var v4522 int32
	_ = v4522
	var v4527 int32
	_ = v4527
	var v4532 int32
	_ = v4532
	var v4535 int32
	_ = v4535
	var v4554 int32
	_ = v4554
	var v4555 int32
	_ = v4555
	var v4558 int32
	_ = v4558
	var v4565 int32
	_ = v4565
	var v4585 int32
	_ = v4585
	var v4593 int32
	_ = v4593
	var v4594 int32
	_ = v4594
	var v4597 int32
	_ = v4597
	var v4618 int32
	_ = v4618
	var v4621 int32
	_ = v4621
	var v4640 int32
	_ = v4640
	var v4641 int32
	_ = v4641
	var v4645 int32
	_ = v4645
	var v4646 int32
	_ = v4646
	var v4647 int32
	_ = v4647
	var v4649 int32
	_ = v4649
	var v4651 int32
	_ = v4651
	var v4672 int32
	_ = v4672
	var v4680 int32
	_ = v4680
	var v4703 int32
	_ = v4703
	var v4704 int32
	_ = v4704
	var v4728 int32
	_ = v4728
	var v4752 int32
	_ = v4752
	var v4776 int32
	_ = v4776
	var v4802 int32
	_ = v4802
	var v4804 int32
	_ = v4804
	var v4828 int32
	_ = v4828
	var v4830 int32
	_ = v4830
	var v4858 int32
	_ = v4858
	var v4882 int32
	_ = v4882
	var v4902 int32
	_ = v4902
	var v4910 int32
	_ = v4910
	var v4911 int32
	_ = v4911
	var v4912 int32
	_ = v4912
	var v4915 int32
	_ = v4915
	var v4916 int32
	_ = v4916
	var v4919 int32
	_ = v4919
	var v4941 int32
	_ = v4941
	var v4949 int32
	_ = v4949
	var v4977 int32
	_ = v4977
	var v5001 int32
	_ = v5001
	var v5027 int32
	_ = v5027
	var v5039 int32
	_ = v5039
	var v5042 int32
	_ = v5042
	var v5043 int64
	_ = v5043
	var v5047 int32
	_ = v5047
	var v5049 int32
	_ = v5049
	var v5050 int32
	_ = v5050
	var v5053 int32
	_ = v5053
	var v5055 int32
	_ = v5055
	var v5057 int32
	_ = v5057
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
	var v5063 int32
	_ = v5063
	var v5064 int32
	_ = v5064
	var v5065 int32
	_ = v5065
	var v5066 int32
	_ = v5066
	var v5067 int32
	_ = v5067
	var v5068 int32
	_ = v5068
	var v5069 int32
	_ = v5069
	var v5070 int32
	_ = v5070
	var v5071 int32
	_ = v5071
	var v5072 int32
	_ = v5072
	var v5073 int32
	_ = v5073
	var v5074 int32
	_ = v5074
	var v5075 int32
	_ = v5075
	var v5076 int32
	_ = v5076
	var v5077 int32
	_ = v5077
	var v5078 int32
	_ = v5078
	var v5079 int32
	_ = v5079
	var v5081 int32
	_ = v5081
	v1 = int32(0)
	v41 = m.G0
	v43 = v41 - int32(672)
	m.G0 = v43
	v46 = v43
	v47 = v1
	v49 = v1
	v50 = v1
	v51 = v1
	v52 = v1
	v53 = v1
	v54 = v1
	v55 = v1
	v56 = v1
	v57 = v1
	v58 = v1
	v59 = v1
	v60 = v1
	v61 = v1
	v62 = v1
	v63 = v1
	v64 = v1
	v65 = v1
	v66 = v1
	v67 = v1
	v69 = int32(-1)
	v71 = v1
	v73 = v1
	v78 = v1
	v82 = v1
	v83 = v1
	goto L1
L1:
	;
	goto L3
L2:
	;
	m.G0 = v46 + int32(672)
	return
L3:
	;
	if v69 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	goto L2
L5:
	;
	v5042 = int32(m.ExcTag)
	v5043 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v5042 == int32(0) {
		goto L399
	} else {
		goto L400
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+336)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	v4554 = int32(1)
	v4555 = v4522 & v4554
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	v4558 = v4527 & v4554
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	F_list_free(m, v4517)
	mBase = m.M
	v4565 = m.ExcPending
	if v4565 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L365
	}
L7:
	;
	v2239 = v2237
	v2240 = v2199
	v2241 = v49
	v2242 = v50
	v2244 = v52
	v2247 = v55
	v2248 = v56
	v2249 = v57
	v2259 = v67
	v2261 = v2220
	v2263 = v2222
	v2274 = v82
	v2275 = v2234
	goto L169
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+336)) = int32(0)
	v2141 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v2147 = F_GetAccessStrategyWithSize(m, v2141)
	mBase = m.M
	v2148 = m.ExcPending
	if v2148 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L165
	}
L9:
	;
	v2199 = v47
	v2202 = v51
	v2204 = v53
	v2205 = v54
	v2209 = v58
	v2210 = v59
	v2211 = v60
	v2212 = v61
	v2213 = v62
	v2214 = v63
	v2215 = v64
	v2216 = v65
	v2219 = v66
	v2220 = v73
	v2222 = v67
	v2229 = v78
	v2234 = v83
	v2237 = int32(1)
	goto L7
L10:
	;
	goto L9
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	v101 = int32(1)
	v102 = v73 & v101
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	v105 = v78 & v101
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v54
	v121 = F_AllocSetContextCreateInternal(m, v108, int32(_a_F_do_autovacuum_0), int32(0), int32(_a_F_do_autovacuum_1), int32(_a_F_do_autovacuum_2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v121
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[3])) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v63
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v54
	F_StartTransactionCommand(m)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v65
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v54
	v172 = F_MultiXactMemberFreezeThreshold(m)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v64
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v199 = F_SearchSysCache1(m, int32(21), v193)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L16
	}
L16:
	;
	if v199 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v65
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v199)+16))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+22)))
	v289 = v287 + v288
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+77)))
	if v290 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v64
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v46)+64)) = v248
	F_errmsg_internal(m, int32(_a_F_do_autovacuum_3), v46-int32(-64))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v65
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_errfinish(m, int32(_a_F_do_autovacuum_4), int32(1937), int32(_a_F_do_autovacuum_5))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[5])) = v318
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v64
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_ReleaseCatCache(m, v199)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L29
	}
L24:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[7])) = v306
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[9])) = v310
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[11])) = v314
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[12]))
	v318 = v317
	goto L23
L25:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+78)))
	if v293 != 0 {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v295 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[9])) = v295
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[7])) = v295
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[11])) = v295
	v318 = v295
	goto L23
L28:
	;
	goto L27
L29:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v63
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v372 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v397 = F_CreateTupleDescCopy(m, v374)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v46)+356)) = int64(446676598788)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v63
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v428 = F_hash_create(m, int32(_a_F_do_autovacuum_6), int32(100), v46+int32(340), int32(40))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v65
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v452 = int32(0)
	v454 = F_table_beginscan_catalog(m, v372, v452, v452)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v65
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v478 = F_heap_getnext(m, v454)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v480 = int32(0)
	if v478 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v499 = v63
	v500 = v64
	v501 = v65
	v504 = v480
	v506 = v478
	v510 = v480
	goto L38
L36:
	;
	v898 = v63
	v899 = v64
	v900 = v65
	v903 = v480
	v909 = v480
	goto L37
L37:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v921)+188))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v922)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	m.T0[v923].(func(*base.Module, int32))(m, v454)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L76
	}
L38:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v506)+16))
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+22)))
	v524 = v522 + v523
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524)+119)))
	switch v525 - int32(109) {
	case 0, 5:
		goto L41
	default:
		v849 = v500
		v850 = v501
		v851 = v504
		v854 = v510
		goto L40
	}
L39:
	;
	v898 = v879
	v899 = v849
	v900 = v850
	v903 = v851
	v909 = v854
	goto L37
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v850
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v879 = F_heap_getnext(m, v454)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L74
	}
L41:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524)+118)))
	if v529 == int32(116) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v524)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v555 = F_checkTempNamespaceStatus(m, v532)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v501
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v606 = F_extractRelOptions(m, v506, v397, int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L49
	}
L45:
	;
	if v555 != int32(1) {
		v849 = v500
		v850 = v501
		v851 = v504
		v854 = v510
		goto L40
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v501
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v581 = F_lappend_oid(m, v510, v528)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v849 = v500
	v850 = v581
	v851 = v504
	v854 = v581
	goto L40
L48:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524)+117)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v688 = F_pgstat_fetch_stat_tabentry_ext(m, v665, v528)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L55
	}
L49:
	;
	if v606 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v664 = int32(0)
	goto L48
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v501
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v634 = F_palloc(m, int32(88))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L53
	}
L53:
	;
	base.MemoryCopy(m, v634, v606+int32(16), int32(88))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_pfree(m, v606)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v664 = v634
	goto L48
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v501
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_relation_needs_vacanalyze(m, v528, v664, v524, v688, v172, v46+int32(287), v46+int32(286), v46+int32(285))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L56
	}
L56:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+287)))
	if v720 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v524)+112))
	if v754 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+286)))
	if v723&int32(1) == int32(0) {
		v752 = v500
		v753 = v504
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v501
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v750 = F_lappend_oid(m, v504, v528)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	v752 = v750
	v753 = v750
	goto L57
L63:
	;
	if v664 != 0 {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v752
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v784 = F_hash_search(m, v428, v524+int32(112), int32(1), v46+int32(284))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+284)))
	if v786 != 0 {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v787 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v784)+8)) = uint8(v787)
	*(*int32)(unsafe.Add(mBase, uint32(v784)+4)) = v528
	if v664 == v787 {
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v792 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v784)+8)) = uint8(v792)
	base.MemoryCopy(m, v784+int32(16), v664, int32(88))
	goto L63
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v752
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v501
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_pfree(m, v664)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if v688 == int32(0) {
		v849 = v752
		v850 = v501
		v851 = v753
		v854 = v510
		goto L40
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v752
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v501
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_pfree(m, v688)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L73
	}
L73:
	;
	v849 = v752
	v850 = v501
	v851 = v753
	v854 = v510
	goto L40
L74:
	;
	if v879 != 0 {
		v499 = v879
		v500 = v849
		v501 = v850
		v504 = v851
		v506 = v879
		v510 = v854
		goto L38
	} else {
		goto L75
	}
L75:
	;
	goto L39
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v971 = v46 + int32(288)
	F_ScanKeyInit(m, v971, int32(18), int32(3), int32(61), int32(116))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1001 = F_table_beginscan_catalog(m, v372, int32(1), v971)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1025 = F_heap_getnext(m, v1001)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L79
	}
L79:
	;
	if v1025 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v1039 = v58
	v1043 = v62
	v1049 = v903
	v1053 = v1025
	goto L83
L81:
	;
	v1371 = v58
	v1375 = v62
	v1381 = v903
	goto L82
L82:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1001)))
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1399)+188))
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1400)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	m.T0[v1401].(func(*base.Module, int32))(m, v1001)
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L113
	}
L83:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+16))
	v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+22)))
	v1069 = v1067 + v1068
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1069)+118)))
	if v1070 == int32(116) {
		v1330 = v1039
		v1331 = v1049
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v1371 = v1330
	v1375 = v1357
	v1381 = v1331
	goto L82
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1043
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1357 = F_heap_getnext(m, v1001)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L111
	}
L86:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1069)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+280)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1039
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1098 = F_extractRelOptions(m, v1053, v397, int32(0))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L88
	}
L87:
	;
	v1193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1069)+117)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1039
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v46)+280))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1217 = F_pgstat_fetch_stat_tabentry_ext(m, v1193, v1211)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L99
	}
L88:
	;
	if v1098 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1039
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1123 = F_palloc(m, int32(88))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1039
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1180 = F_hash_search(m, v428, v46+int32(280), int32(0), v46+int32(276))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L94
	}
L92:
	;
	base.MemoryCopy(m, v1123, v1098+int32(16), int32(88))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1039
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_pfree(m, v1098)
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L93
	}
L93:
	;
	v1191 = v1123
	goto L87
L94:
	;
	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+276)))
	if v1183 != int32(1) {
		v1191 = int32(0)
		goto L87
	} else {
		goto L95
	}
L95:
	;
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180)+8)))
	if v1189 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v1190 = v1180 + int32(16)
	goto L98
L97:
	;
	v1190 = int32(0)
	goto L98
L98:
	;
	v1191 = v1190
	goto L87
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1039
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v46)+280))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_relation_needs_vacanalyze(m, v1237, v1191, v1069, v1217, v172, v46+int32(279), v46+int32(278), v46+int32(277))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L100
	}
L100:
	;
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+279)))
	if v1250 == int32(1) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1039
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v46)+280))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1276 = F_lappend_oid(m, v1049, v1271)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L104
	}
L102:
	;
	v1278 = v1039
	v1279 = v1049
	goto L103
L103:
	;
	if v1098 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v1278 = v1276
	v1279 = v1276
	goto L103
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_pfree(m, v1191)
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v1217 == int32(0) {
		v1330 = v1278
		v1331 = v1279
		goto L85
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_pfree(m, v1217)
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v1330 = v1278
	v1331 = v1279
	goto L85
L111:
	;
	if v1357 != 0 {
		v1039 = v1330
		v1043 = v1357
		v1049 = v1331
		v1053 = v1357
		goto L83
	} else {
		goto L112
	}
L112:
	;
	goto L84
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_relation_close(m, v372, int32(1))
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L114
	}
L114:
	;
	if v909 == int32(0) {
		goto L8
	} else {
		goto L115
	}
L115:
	;
	v1453 = int32(0)
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v909)+4))
	if v1454 <= v1453 {
		goto L8
	} else {
		goto L116
	}
L116:
	;
	v1477 = v1453
	goto L117
L117:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v909)+12))
	v1500 = v1497 + v1477<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+336)) = v1500
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1500)))
	v1505 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[13]))
	if v1505 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L8
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_ProcessInterrupts(m)
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1553 = F_ConditionalLockRelationOid(m, v1503, int32(8))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L124
	}
L122:
	;
	goto L121
L123:
	;
	v2077 = v1477 + int32(1)
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v909)+4))
	if v2077 < v2078 {
		v1477 = v2077
		goto L117
	} else {
		goto L164
	}
L124:
	;
	if v1553 == int32(0) {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1581 = F_SearchSysCacheCopy(m, int32(57), v1503, int32(0))
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L126
	}
L126:
	;
	if v1581 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_UnlockRelationOid(m, v1503, int32(8))
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1581)+16))
	v1611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1610)+22)))
	v1612 = v1610 + v1611
	v1613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1612)+119)))
	switch v1613 - int32(109) {
	case 0, 5:
		goto L133
	default:
		goto L132
	}
L130:
	;
	goto L123
L131:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1612)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1667 = F_checkTempNamespaceStatus(m, v1644)
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L136
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_UnlockRelationOid(m, v1503, int32(8))
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L135
	}
L133:
	;
	v1616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1612)+118)))
	if v1616 == int32(116) {
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	goto L123
L136:
	;
	if v1667 != int32(1) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_UnlockRelationOid(m, v1503, int32(8))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v1612)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1719 = m.G0
	v1721 = v1719 - int32(32)
	m.G0 = v1721
	v1723 = int32(264)
	*(*uint16)(unsafe.Add(mBase, uint32(v1721)+30)) = uint16(v1723)
	v1725 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1721)+28)) = uint16(v1725)
	*(*int32)(unsafe.Add(mBase, uint32(v1721)+24)) = v1696
	*(*int32)(unsafe.Add(mBase, uint32(v1721)+20)) = int32(2615)
	v1731 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v1721)+16)) = v1731
	v1735 = int32(1)
	v1741 = F_LockAcquireExtended(m, v1721+int32(16), v1735, v1725, v1735, v1721+int32(12), v1725)
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L143
	}
L140:
	;
	goto L123
L141:
	;
	m.G0 = v1721 + int32(32)
	if v1741 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L142:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L144
	}
L143:
	;
	switch v1741 {
	case 0, 3:
		goto L141
	default:
		goto L142
	}
L144:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1721)+12))
	v1746 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1745)+53)) = uint8(v1746)
	goto L145
L145:
	;
	goto L141
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_UnlockRelationOid(m, v1503, int32(8))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1802 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L150
	}
L149:
	;
	goto L123
L150:
	;
	if v1802 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	v1823 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1828 = F_get_database_name(m, v1823)
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1938 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L158
	}
L154:
	;
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v1612)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v1853 = F_get_namespace_name(m, v1830)
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v46)+88)) = v1612 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+84)) = v1853
	*(*int32)(unsafe.Add(mBase, uint32(v46)+80)) = v1828
	F_errmsg(m, int32(_a_F_do_autovacuum_7), v46+int32(80))
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_errfinish(m, int32(_a_F_do_autovacuum_4), int32(2235), int32(_a_F_do_autovacuum_5))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L157
	}
L157:
	;
	goto L153
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_PushActiveSnapshot(m, v1938)
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+272)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+268)) = v1503
	*(*int32)(unsafe.Add(mBase, uint32(v46)+264)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_performDeletion(m, v46+int32(264), int32(1), int32(21))
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	F_StartTransactionCommand(m)
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L163
	}
L163:
	;
	v2071 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v2071
	goto L123
L164:
	;
	goto L118
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2147
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v899
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v105)
	v2168 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v172
	v2178 = F_AllocSetContextCreateInternal(m, v2168, int32(_a_F_do_autovacuum_8), int32(0), int32(_a_F_do_autovacuum_1), int32(_a_F_do_autovacuum_2))
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		v5027 = v71
		v5039 = v83
		goto L5
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[14])) = v2178
	v2182 = v1381 + int32(4)
	v2183 = int32(0)
	if v1381 == v2183 {
		v4496 = v47
		v4498 = v49
		v4499 = v50
		v4500 = v428
		v4501 = v52
		v4502 = v397
		v4503 = v172
		v4504 = v55
		v4505 = v56
		v4506 = v57
		v4507 = v1371
		v4508 = v59
		v4509 = v2182
		v4510 = v2147
		v4511 = v1375
		v4512 = v898
		v4513 = v899
		v4514 = v900
		v4516 = v67
		v4517 = v1381
		v4520 = v71
		v4522 = v73
		v4527 = v78
		v4532 = v83
		v4535 = v2183
		goto L6
	} else {
		goto L167
	}
L167:
	;
	v2186 = int32(0)
	v2187 = base.B2i32(v428 != v2186)
	v2189 = v1381 + int32(12)
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+4))
	if v2193 <= v2186 {
		v4496 = v47
		v4498 = v49
		v4499 = v50
		v4500 = v428
		v4501 = v52
		v4502 = v397
		v4503 = v172
		v4504 = v55
		v4505 = v56
		v4506 = v57
		v4507 = v1371
		v4508 = v2189
		v4509 = v2182
		v4510 = v2147
		v4511 = v1375
		v4512 = v898
		v4513 = v899
		v4514 = v900
		v4516 = v67
		v4517 = v1381
		v4520 = v71
		v4522 = v73
		v4527 = v2187
		v4532 = v2186
		v4535 = v2186
		goto L6
	} else {
		goto L168
	}
L168:
	;
	v2199 = v2186
	v2202 = v428
	v2204 = v397
	v2205 = v172
	v2209 = v1371
	v2210 = v2189
	v2211 = v2182
	v2212 = v2147
	v2213 = v1375
	v2214 = v898
	v2215 = v899
	v2216 = v900
	v2219 = v1381
	v2220 = v2186
	v2222 = v71
	v2229 = v2187
	v2234 = v2186
	v2237 = int32(0)
	goto L7
L169:
	;
	if v2239 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L170:
	;
	v4496 = v2240
	v4498 = v4450
	v4499 = v4451
	v4500 = v2202
	v4501 = v4453
	v4502 = v2204
	v4503 = v2205
	v4504 = v2247
	v4505 = v2248
	v4506 = v4458
	v4507 = v2209
	v4508 = v2210
	v4509 = v2211
	v4510 = v2212
	v4511 = v2213
	v4512 = v2214
	v4513 = v2215
	v4514 = v2216
	v4516 = v4468
	v4517 = v2219
	v4520 = v4472
	v4522 = v2261
	v4527 = v2229
	v4532 = v4484
	v4535 = base.B2i32(v4484 == int32(0)) & v4467
	goto L6
L171:
	;
	v4488 = v2240 + int32(1)
	v4489 = *(*int32)(unsafe.Add(mBase, uint32(v2211)))
	if v4488 < v4489 {
		goto L362
	} else {
		goto L363
	}
L172:
	;
	v4450 = v4410
	v4451 = v4411
	v4453 = v4413
	v4458 = v4418
	v4467 = v2261
	v4468 = v4428
	v4472 = v4432
	v4484 = v4444
	goto L171
L173:
	;
	if v4225 != 0 {
		goto L347
	} else {
		goto L348
	}
L174:
	;
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(v2210)))
	v2283 = v2280 + v2240<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+336)) = v2283
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(v2283)))
	v2288 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[13]))
	if v2288 != 0 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	goto L176
L176:
	;
	v3583 = v2263 + int32(8)
	if v2274 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	v2306 = int32(1)
	v2307 = v2261 & v2306
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2307)
	v2310 = v2229 & v2306
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2310)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_ProcessInterrupts(m)
	mBase = m.M
	v2316 = m.ExcPending
	if v2316 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v2318 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[15]))
	if v2318 != 0 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	goto L179
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[15])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	v2338 = int32(1)
	v2339 = v2261 & v2338
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2339)
	v2342 = v2229 & v2338
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2342)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v2350 = m.ExcPending
	if v2350 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	v2368 = int32(1)
	v2369 = v2261 & v2368
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	v2372 = v2229 & v2368
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v2378 = F_SearchSysCache1(m, int32(57), v2286)
	mBase = m.M
	v2379 = m.ExcPending
	if v2379 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L185
	}
L184:
	;
	goto L183
L185:
	;
	if v2378 == int32(0) {
		v4410 = v2241
		v4411 = v2242
		v4413 = v2244
		v4418 = v2249
		v4428 = v2259
		v4432 = v2263
		v4444 = v2275
		goto L172
	} else {
		goto L186
	}
L186:
	;
	v2382 = *(*int32)(unsafe.Add(mBase, uint32(v2378)+16))
	v2383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2382)+22)))
	v2385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2382+v2383)+117)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_ReleaseCatCache(m, v2378)
	mBase = m.M
	v2409 = m.ExcPending
	if v2409 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	v2429 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v2437 = F_LWLockAcquire(m, v2429+int32(2944), int32(0))
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	v2458 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v2466 = F_LWLockAcquire(m, v2458+int32(2816), int32(1))
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L189
	}
L189:
	;
	v2469 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[17]))
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(v2469)+28))
	if v2470 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L190:
	;
	v3231 = *(*float64)(unsafe.Add(mBase, uint32(v3098)+64))
	*(*float64)(unsafe.Add(mBase, _c_F_do_autovacuum[18])) = v3231
	v3234 = *(*int32)(unsafe.Add(mBase, uint32(v3098)+72))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[19])) = v3234
	v3237 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[20]))
	v3238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3098)+76)))
	if v3238 == int32(1) {
		goto L280
	} else {
		goto L281
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v3155
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	v3186 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v3194 = F_LWLockAcquire(m, v3186+int32(2944), int32(0))
	mBase = m.M
	v3195 = m.ExcPending
	if v3195 != 0 {
		v5027 = v3156
		v5039 = v2275
		goto L5
	} else {
		goto L277
	}
L192:
	;
	v2858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2690)+117)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v2881 = F_pgstat_fetch_stat_tabentry_ext(m, v2858, v2856)
	mBase = m.M
	v2882 = m.ExcPending
	if v2882 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L226
	}
L193:
	;
	v2807 = int32(0)
	v2810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2690)+119)))
	if (v2229^int32(-1)|base.B2i32(v2810 != int32(116)))&int32(1) != 0 {
		v2855 = v2807
		v2856 = v2286
		goto L192
	} else {
		goto L218
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	v2765 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_LWLockRelease(m, v2765+int32(2816))
	mBase = m.M
	v2773 = m.ExcPending
	if v2773 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L216
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	v2589 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_LWLockRelease(m, v2589+int32(2816))
	mBase = m.M
	v2597 = m.ExcPending
	if v2597 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L208
	}
L196:
	;
	v2474 = v2469 + int32(24)
	if v2470 == v2474 {
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v2477 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[4]))
	v2479 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[20]))
	v2481 = v2470
	goto L198
L198:
	;
	if v2481 == v2479 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	goto L195
L200:
	;
	v2528 = *(*int32)(unsafe.Add(mBase, uint32(v2481)+4))
	if v2528 != v2474 {
		v2481 = v2528
		goto L198
	} else {
		goto L207
	}
L201:
	;
	v2521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2481)+36)))
	if v2521 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v2524 = *(*int32)(unsafe.Add(mBase, uint32(v2481)+8))
	if v2524 != v2477 {
		goto L200
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v2526 = *(*int32)(unsafe.Add(mBase, uint32(v2481)+12))
	if v2526 == v2286 {
		goto L194
	} else {
		goto L206
	}
L205:
	;
	goto L204
L206:
	;
	goto L200
L207:
	;
	goto L199
L208:
	;
	v2599 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[20]))
	*(*uint8)(unsafe.Add(mBase, uint32(v2599)+36)) = uint8(v2385)
	*(*int32)(unsafe.Add(mBase, uint32(v2599)+12)) = v2286
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	v2621 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_LWLockRelease(m, v2621+int32(2944))
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L209
	}
L209:
	;
	v2632 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v2632
	*(*int32)(unsafe.Add(mBase, uint32(v46)+392)) = v2286
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v2659 = F_SearchSysCacheCopy(m, int32(57), v2286, int32(0))
	mBase = m.M
	v2660 = m.ExcPending
	if v2660 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L210
	}
L210:
	;
	if v2659 == int32(0) {
		v3155 = v2259
		v3156 = v2263
		goto L191
	} else {
		goto L211
	}
L211:
	;
	v2663 = *(*int32)(unsafe.Add(mBase, uint32(v2659)+16))
	v2664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2663)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v2688 = F_extractRelOptions(m, v2659, v2204, int32(0))
	mBase = m.M
	v2689 = m.ExcPending
	if v2689 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L212
	}
L212:
	;
	v2690 = v2663 + v2664
	if v2688 == int32(0) {
		goto L193
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v2716 = F_palloc(m, int32(88))
	mBase = m.M
	v2717 = m.ExcPending
	if v2717 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L214
	}
L214:
	;
	base.MemoryCopy(m, v2716, v2688+int32(16), int32(88))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_pfree(m, v2688)
	mBase = m.M
	v2745 = m.ExcPending
	if v2745 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L215
	}
L215:
	;
	v2855 = v2716
	v2856 = v2286
	goto L192
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	v2790 = int32(1)
	v2792 = v2261 & v2790
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2792)
	v2795 = v2229 & v2790
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2795)
	v2798 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_LWLockRelease(m, v2798+int32(2944))
	mBase = m.M
	v2806 = m.ExcPending
	if v2806 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L217
	}
L217:
	;
	v4450 = v2241
	v4451 = v2242
	v4453 = v2244
	v4458 = v2249
	v4467 = v2790
	v4468 = v2259
	v4472 = v2263
	v4484 = v2275
	goto L171
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v2843 = F_hash_search(m, v2202, v46+int32(392), int32(0), v46+int32(388))
	mBase = m.M
	v2844 = m.ExcPending
	if v2844 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L219
	}
L219:
	;
	v2845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+388)))
	if v2845 == int32(1) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v2851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2843)+8)))
	if v2851 != 0 {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	v2853 = v2807
	goto L222
L222:
	;
	v2854 = *(*int32)(unsafe.Add(mBase, uint32(v46)+392))
	v2855 = v2853
	v2856 = v2854
	goto L192
L223:
	;
	v2852 = v2843 + int32(16)
	goto L225
L224:
	;
	v2852 = int32(0)
	goto L225
L225:
	;
	v2853 = v2852
	goto L222
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_relation_needs_vacanalyze(m, v2856, v2855, v2690, v2881, v2205, v46+int32(391), v46+int32(390), v46+int32(389))
	mBase = m.M
	v2912 = m.ExcPending
	if v2912 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L227
	}
L227:
	;
	if v2881 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_pfree(m, v2881)
	mBase = m.M
	v2936 = m.ExcPending
	if v2936 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v2937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2690)+119)))
	if v2937 != int32(116) {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	goto L230
L232:
	;
	v2945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+391)))
	if v2944&int32(1) == int32(0) {
		goto L237
	} else {
		goto L238
	}
L233:
	;
	v2940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+390)))
	v2944 = v2940
	goto L232
L234:
	;
	goto L235
L235:
	;
	v2941 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+390)) = uint8(v2941)
	v2944 = v2941
	goto L232
L236:
	;
	if v2688 != 0 {
		goto L271
	} else {
		goto L272
	}
L237:
	;
	v2950 = int32(0)
	if v2945&int32(1) == v2950 {
		v3098 = v2950
		goto L236
	} else {
		goto L240
	}
L238:
	;
	v2955 = v2263
	goto L239
L239:
	;
	if v2855 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L240:
	;
	v2955 = v2950
	goto L239
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v3030 = F_palloc(m, int32(96))
	mBase = m.M
	v3031 = m.ExcPending
	if v3031 != 0 {
		v5027 = v3001
		v5039 = v2275
		goto L5
	} else {
		goto L259
	}
L242:
	;
	v3000 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[5]))
	v3001 = v2993
	v3002 = v2994
	v3003 = v2995
	v3004 = v2996
	v3005 = v3000
	v3006 = v2998
	goto L241
L243:
	;
	v2959 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[11]))
	v2961 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[9]))
	v2963 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[7]))
	v2965 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	v2993 = v2955
	v2994 = v2959
	v2995 = v2961
	v2996 = v2963
	v2998 = v2965
	goto L242
L244:
	;
	goto L245
L245:
	;
	v2967 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[11]))
	v2968 = *(*int32)(unsafe.Add(mBase, uint32(v2855)+36))
	if v2968 < int32(0) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v2971 = v2967
	goto L248
L247:
	;
	v2971 = v2968
	goto L248
L248:
	;
	v2973 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[9]))
	v2974 = *(*int32)(unsafe.Add(mBase, uint32(v2855)+32))
	if v2974 < int32(0) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v2977 = v2973
	goto L251
L250:
	;
	v2977 = v2974
	goto L251
L251:
	;
	v2979 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[7]))
	v2980 = *(*int32)(unsafe.Add(mBase, uint32(v2855)+24))
	if v2980 < int32(0) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v2983 = v2979
	goto L254
L253:
	;
	v2983 = v2980
	goto L254
L254:
	;
	v2985 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	v2986 = *(*int32)(unsafe.Add(mBase, uint32(v2855)+48))
	if v2986 < int32(0) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v2989 = v2985
	goto L257
L256:
	;
	v2989 = v2986
	goto L257
L257:
	;
	v2990 = *(*int32)(unsafe.Add(mBase, uint32(v2855)+44))
	if int32(0) <= v2990 {
		v3001 = v2986
		v3002 = v2971
		v3003 = v2977
		v3004 = v2983
		v3005 = v2990
		v3006 = v2989
		goto L241
	} else {
		goto L258
	}
L258:
	;
	v2993 = v2986
	v2994 = v2971
	v2995 = v2977
	v2996 = v2983
	v2998 = v2989
	goto L242
L259:
	;
	v3032 = *(*int32)(unsafe.Add(mBase, uint32(v46)+392))
	*(*int32)(unsafe.Add(mBase, uint32(v3030))) = v3032
	v3034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2690)+117)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3030)+77)) = uint8(v3034)
	v3036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+389)))
	*(*int32)(unsafe.Add(mBase, uint32(v3030)+56)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v3030)+36)) = int64(0)
	v3041 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3030)+44)) = v3041
	*(*int32)(unsafe.Add(mBase, uint32(v3030)+32)) = v3006
	*(*uint8)(unsafe.Add(mBase, uint32(v3030)+28)) = uint8(v3036)
	*(*int32)(unsafe.Add(mBase, uint32(v3030)+24)) = v3005
	*(*int32)(unsafe.Add(mBase, uint32(v3030)+20)) = v3002
	*(*int32)(unsafe.Add(mBase, uint32(v3030)+16)) = v3003
	*(*int32)(unsafe.Add(mBase, uint32(v3030)+12)) = v3004
	v3049 = int32(1)
	if v2945&v3049 != 0 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v3057 = int32(577)
	goto L262
L261:
	;
	v3057 = v3041
	goto L262
L262:
	;
	if v3036 != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v3061 = int32(0)
	goto L265
L264:
	;
	v3061 = int32(32)
	goto L265
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3030)+8)) = v2944<<(uint(v3049)%32)&int32(254) | v3057 | v3061
	v3065 = *(*float64)(unsafe.Add(mBase, _c_F_do_autovacuum[22]))
	*(*float64)(unsafe.Add(mBase, uint32(v3030)+48)) = v3065
	if v2855 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3030)+76)) = uint8(v3096)
	v3098 = v3030
	goto L236
L267:
	;
	v3069 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3030)+88)) = v3069
	*(*int64)(unsafe.Add(mBase, uint32(v3030)+80)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3030)+64)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v3030)+72)) = v3069
	v3096 = int32(1)
	goto L266
L268:
	;
	goto L269
L269:
	;
	v3078 = *(*int32)(unsafe.Add(mBase, uint32(v2855)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v3030)+72)) = v3078
	v3080 = *(*float64)(unsafe.Add(mBase, uint32(v2855)+56))
	v3081 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3030)+88)) = v3081
	*(*int64)(unsafe.Add(mBase, uint32(v3030)+80)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v3030)+64)) = v3080
	v3087 = *(*int32)(unsafe.Add(mBase, uint32(v2855)+20))
	if v3081 < v3087 {
		v3096 = v3081
		goto L266
	} else {
		goto L270
	}
L270:
	;
	v3090 = *(*float64)(unsafe.Add(mBase, uint32(v2855)+56))
	v3096 = base.B2i32(base.F64_ge(v3090, float64(0)) == int32(0))
	goto L266
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v3098
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_pfree(m, v2855)
	mBase = m.M
	v3129 = m.ExcPending
	if v3129 != 0 {
		v5027 = v3098
		v5039 = v2275
		goto L5
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v3098
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_pfree(m, v2659)
	mBase = m.M
	v3153 = m.ExcPending
	if v3153 != 0 {
		v5027 = v3098
		v5039 = v2275
		goto L5
	} else {
		goto L275
	}
L274:
	;
	goto L273
L275:
	;
	if v3098 != 0 {
		goto L190
	} else {
		goto L276
	}
L276:
	;
	v3155 = v3098
	v3156 = v3098
	goto L191
L277:
	;
	v3197 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[20]))
	v3198 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3197)+36)) = uint8(v3198)
	*(*int32)(unsafe.Add(mBase, uint32(v3197)+12)) = v3198
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v3155
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	v3221 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_LWLockRelease(m, v3221+int32(2944))
	mBase = m.M
	v3229 = m.ExcPending
	if v3229 != 0 {
		v5027 = v3156
		v5039 = v2275
		goto L5
	} else {
		goto L278
	}
L278:
	;
	v4450 = v2241
	v4451 = v2242
	v4453 = v2244
	v4458 = v2249
	v4467 = v2261
	v4468 = v3155
	v4472 = v3156
	v4484 = v2275
	goto L171
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v3098
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	v3264 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v3272 = F_LWLockAcquire(m, v3264+int32(2816), int32(1))
	mBase = m.M
	v3273 = m.ExcPending
	if v3273 != 0 {
		v5027 = v3098
		v5039 = v2275
		goto L5
	} else {
		goto L283
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3237)+32)) = int32(1)
	goto L279
L281:
	;
	goto L282
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3237)+32)) = int32(0)
	goto L279
L283:
	;
	v3274 = int32(0)
	v3276 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[17]))
	v3277 = *(*int32)(unsafe.Add(mBase, uint32(v3276)+28))
	v3278 = *(*int32)(unsafe.Add(mBase, uint32(v3276)+uint32(_c_F_do_autovacuum[23])))
	if v3277 == v3274 {
		v3358 = v3274
		goto L284
	} else {
		goto L285
	}
L284:
	;
	if v3358 != v3278 {
		goto L293
	} else {
		goto L294
	}
L285:
	;
	v3282 = v3276 + int32(24)
	if v3277 == v3282 {
		v3358 = v3274
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v3285 = v3277
	v3310 = v3274
	goto L287
L287:
	;
	v3324 = *(*int32)(unsafe.Add(mBase, uint32(v3285)+16))
	if v3324 != 0 {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	v3358 = v3329
	goto L284
L289:
	;
	v3325 = *(*int32)(unsafe.Add(mBase, uint32(v3285)+32))
	v3329 = v3310 + base.B2i32(v3325 != int32(0))
	goto L291
L290:
	;
	v3329 = v3310
	goto L291
L291:
	;
	v3330 = *(*int32)(unsafe.Add(mBase, uint32(v3285)+4))
	if v3330 != v3282 {
		v3285 = v3330
		v3310 = v3329
		goto L287
	} else {
		goto L292
	}
L292:
	;
	goto L288
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3276)+uint32(_c_F_do_autovacuum[23]))) = v3358
	goto L295
L294:
	;
	goto L295
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v3098
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	v3393 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_LWLockRelease(m, v3393+int32(2816))
	mBase = m.M
	v3401 = m.ExcPending
	if v3401 != 0 {
		v5027 = v3098
		v5039 = v2275
		goto L5
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v3098
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_VacuumUpdateCosts(m)
	mBase = m.M
	v3425 = m.ExcPending
	if v3425 != 0 {
		v5027 = v3098
		v5039 = v2275
		goto L5
	} else {
		goto L297
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v3098
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	v3445 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_MemoryContextReset(m, v3445)
	mBase = m.M
	v3451 = m.ExcPending
	if v3451 != 0 {
		v5027 = v3098
		v5039 = v2275
		goto L5
	} else {
		goto L298
	}
L298:
	;
	v3452 = *(*int32)(unsafe.Add(mBase, uint32(v3098)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v3098
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v3475 = F_get_rel_name(m, v3452)
	mBase = m.M
	v3476 = m.ExcPending
	if v3476 != 0 {
		v5027 = v3098
		v5039 = v2275
		goto L5
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3098)+80)) = v3475
	v3478 = *(*int32)(unsafe.Add(mBase, uint32(v3098)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	v3485 = v3098 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v3485
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v3098
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v3503 = F_get_rel_namespace(m, v3478)
	mBase = m.M
	v3504 = m.ExcPending
	if v3504 != 0 {
		v5027 = v3098
		v5039 = v2275
		goto L5
	} else {
		goto L300
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v3485
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v3098
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v3527 = F_get_namespace_name(m, v3503)
	mBase = m.M
	v3528 = m.ExcPending
	if v3528 != 0 {
		v5027 = v3098
		v5039 = v2275
		goto L5
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3098)+84)) = v3527
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	v3535 = v3098 + int32(84)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v3535
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v3485
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v3098
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v2369)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v2372)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	v3551 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v3556 = F_get_database_name(m, v3551)
	mBase = m.M
	v3557 = m.ExcPending
	if v3557 != 0 {
		v5027 = v3098
		v5039 = v2275
		goto L5
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3098)+88)) = v3556
	v3560 = v3098 + int32(88)
	v3561 = *(*int32)(unsafe.Add(mBase, uint32(v3098)+80))
	if v3561 == int32(0) {
		v4188 = v3556
		v4189 = v3535
		v4191 = v3485
		v4196 = v3560
		v4210 = v3098
		v4222 = v2275
		v4225 = v3556
		goto L173
	} else {
		goto L303
	}
L303:
	;
	v3564 = *(*int32)(unsafe.Add(mBase, uint32(v3535)))
	if v3564 == int32(0) {
		v4188 = v3556
		v4189 = v3535
		v4191 = v3485
		v4196 = v3560
		v4210 = v3098
		v4222 = v2275
		v4225 = v3556
		goto L173
	} else {
		goto L304
	}
L304:
	;
	if v3556 == int32(0) {
		v4188 = v3556
		v4189 = v3535
		v4191 = v3485
		v4196 = v3560
		v4210 = v3098
		v4222 = v2275
		v4225 = v3556
		goto L173
	} else {
		goto L305
	}
L305:
	;
	v3571 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[24]))
	v3573 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[25]))
	goto L306
L306:
	;
	v3575 = v46 + int32(96)
	*(*int32)(unsafe.Add(mBase, uint32(v3575)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3575))) = v46 + int32(92)
	goto L309
L307:
	;
	v2239 = int32(1)
	v2241 = v3556
	v2242 = v3535
	v2244 = v3485
	v2247 = v3571
	v2248 = v3573
	v2249 = v3560
	v2263 = v3098
	v2274 = int32(0)
	goto L169
L309:
	;
	goto L307
L310:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[24])) = v2247
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[25])) = v2248
	v4181 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v4181
	v4184 = *(*int32)(unsafe.Add(mBase, uint32(v2249)))
	v4188 = v2241
	v4189 = v2242
	v4191 = v2244
	v4196 = v2249
	v4210 = v2263
	v4222 = int32(1)
	v4225 = v4184
	goto L173
L311:
	;
	v3588 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v3588
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[25])) = v46 + int32(96)
	v3594 = *(*int32)(unsafe.Add(mBase, uint32(v3583)))
	if v3594&int32(1) != 0 {
		goto L315
	} else {
		goto L316
	}
L312:
	;
	goto L313
L313:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[24])) = v2247
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[25])) = v2248
	v3966 = int32(_a_F_do_autovacuum_9)
	v3968 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[26]))
	v3969 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[26])) = v3968 + v3969
	v3972 = *(*int32)(unsafe.Add(mBase, uint32(v3583)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	v3981 = v2261 & v3969
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3981)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	v3985 = v2229 & v3969
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3985)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v4001 = m.ExcPending
	if v4001 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L337
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	v3687 = int32(1)
	v3688 = v2261 & v3687
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3688)
	v3691 = v2229 & v3687
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3691)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v3697 = v46 + int32(400)
	v3698 = F_strlen(m, v3697)
	mBase = m.M
	v3699 = *(*int32)(unsafe.Add(mBase, uint32(v2242)))
	v3700 = *(*int32)(unsafe.Add(mBase, uint32(v2244)))
	v3701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2263)+28)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3688)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3691)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	if v3701 != 0 {
		goto L323
	} else {
		goto L324
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	v3614 = int32(1)
	v3615 = v2261 & v3614
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3615)
	v3618 = v2229 & v3614
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3618)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	if v3594&int32(2) != 0 {
		goto L318
	} else {
		goto L319
	}
L316:
	;
	goto L317
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	v3654 = int32(1)
	v3655 = v2261 & v3654
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3655)
	v3658 = v2229 & v3654
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3658)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v3668 = F_pg_snprintf(m, v46+int32(400), int32(184), int32(_a_F_do_autovacuum_10), int32(0))
	mBase = m.M
	v3669 = m.ExcPending
	if v3669 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L322
	}
L318:
	;
	v3627 = int32(_a_F_do_autovacuum_11)
	goto L320
L319:
	;
	v3627 = int32(_a_F_do_autovacuum_12)
	goto L320
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+32)) = v3627
	v3635 = F_pg_snprintf(m, v46+int32(400), int32(184), int32(_a_F_do_autovacuum_13), v46+int32(32))
	mBase = m.M
	v3636 = m.ExcPending
	if v3636 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L321
	}
L321:
	;
	goto L314
L322:
	;
	goto L314
L323:
	;
	v3726 = int32(_a_F_do_autovacuum_14)
	goto L325
L324:
	;
	v3726 = int32(_a_F_do_autovacuum_12)
	goto L325
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v3726
	*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v3700
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v3699
	v3736 = F_pg_snprintf(m, v3698+v3697, int32(184)-v3698, int32(_a_F_do_autovacuum_15), v46+int32(16))
	mBase = m.M
	v3737 = m.ExcPending
	if v3737 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L326
	}
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3688)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3691)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v3761 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[27]))
	if v3761 < int32(0) {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3688)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3691)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_pgstat_report_activity(m, int32(3), v3697)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3688)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3691)
	v3809 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v3819 = F_AllocSetContextCreateInternal(m, v3809, int32(_a_F_do_autovacuum_16), int32(0), int32(_a_F_do_autovacuum_1), int32(_a_F_do_autovacuum_2))
	mBase = m.M
	v3820 = m.ExcPending
	if v3820 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L331
	}
L328:
	;
	v3765 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_do_autovacuum[28])) = v3765
	goto L330
L329:
	;
	goto L330
L330:
	;
	goto L327
L331:
	;
	v3821 = int32(_a_F_do_autovacuum_17)
	v3822 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v3819
	v3825 = *(*int32)(unsafe.Add(mBase, uint32(v2244)))
	v3826 = *(*int32)(unsafe.Add(mBase, uint32(v2242)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3688)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3691)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v3850 = F_makeRangeVar(m, v3826, v3825, int32(-1))
	mBase = m.M
	v3851 = m.ExcPending
	if v3851 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L332
	}
L332:
	;
	v3852 = *(*int32)(unsafe.Add(mBase, uint32(v2263)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3688)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3691)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v3876 = F_makeVacuumRelation(m, v3850, v3852, int32(0))
	mBase = m.M
	v3877 = m.ExcPending
	if v3877 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L333
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+396)) = v3876
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3688)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3691)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v3876
	v3906 = F_list_make1_impl(m, int32(1), v46+int32(12))
	mBase = m.M
	v3907 = m.ExcPending
	if v3907 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v3822
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3688)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3691)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_vacuum(m, v3906, v3583, v2212, v3819, int32(1))
	mBase = m.M
	v3934 = m.ExcPending
	if v3934 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3688)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3691)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_MemoryContextDelete(m, v3819)
	mBase = m.M
	v3958 = m.ExcPending
	if v3958 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[29])) = int32(0)
	goto L310
L337:
	;
	v4002 = *(*int32)(unsafe.Add(mBase, uint32(v2249)))
	v4003 = *(*int32)(unsafe.Add(mBase, uint32(v2242)))
	v4004 = *(*int32)(unsafe.Add(mBase, uint32(v2244)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3981)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3985)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	*(*int32)(unsafe.Add(mBase, uint32(v46)+56)) = v4004
	*(*int32)(unsafe.Add(mBase, uint32(v46)+52)) = v4003
	*(*int32)(unsafe.Add(mBase, uint32(v46)+48)) = v4002
	if v3972&int32(1) != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v4034 = int32(_a_F_do_autovacuum_18)
	goto L340
L339:
	;
	v4034 = int32(_a_F_do_autovacuum_19)
	goto L340
L340:
	;
	F_errcontext_msg(m, v4034, v46+int32(48))
	mBase = m.M
	v4038 = m.ExcPending
	if v4038 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3981)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3985)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_EmitErrorReport(m)
	mBase = m.M
	v4062 = m.ExcPending
	if v4062 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L342
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3981)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3985)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_AbortOutOfAnyTransaction(m)
	mBase = m.M
	v4086 = m.ExcPending
	if v4086 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L343
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3981)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3985)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_FlushErrorState(m)
	mBase = m.M
	v4110 = m.ExcPending
	if v4110 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L344
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3981)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3985)
	v4130 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_MemoryContextReset(m, v4130)
	mBase = m.M
	v4136 = m.ExcPending
	if v4136 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L345
	}
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v2249
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v2263
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v3981)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v3985)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_StartTransactionCommand(m)
	mBase = m.M
	v4160 = m.ExcPending
	if v4160 != 0 {
		v5027 = v2263
		v5039 = v2275
		goto L5
	} else {
		goto L346
	}
L346:
	;
	v4161 = int32(_a_F_do_autovacuum_9)
	v4163 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[26]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[26])) = v4163 - int32(1)
	goto L310
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4196
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4188
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4189
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4191
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	v4243 = int32(1)
	v4244 = v2261 & v4243
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4244)
	v4247 = v2229 & v4243
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4247)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_pfree(m, v4225)
	mBase = m.M
	v4253 = m.ExcPending
	if v4253 != 0 {
		v5027 = v4210
		v5039 = v4222
		goto L5
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	v4254 = *(*int32)(unsafe.Add(mBase, uint32(v4189)))
	if v4254 != 0 {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	goto L349
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4196
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4188
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4189
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4191
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	v4272 = int32(1)
	v4273 = v2261 & v4272
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4273)
	v4276 = v2229 & v4272
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4276)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_pfree(m, v4254)
	mBase = m.M
	v4282 = m.ExcPending
	if v4282 != 0 {
		v5027 = v4210
		v5039 = v4222
		goto L5
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	v4283 = *(*int32)(unsafe.Add(mBase, uint32(v4191)))
	if v4283 != 0 {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	goto L353
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4196
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4188
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4189
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4191
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	v4301 = int32(1)
	v4302 = v2261 & v4301
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4302)
	v4305 = v2229 & v4301
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4305)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_pfree(m, v4283)
	mBase = m.M
	v4311 = m.ExcPending
	if v4311 != 0 {
		v5027 = v4210
		v5039 = v4222
		goto L5
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4196
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4188
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4189
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4191
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	v4329 = int32(1)
	v4330 = v2261 & v4329
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4330)
	v4333 = v2229 & v4329
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4333)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_pfree(m, v4210)
	mBase = m.M
	v4339 = m.ExcPending
	if v4339 != 0 {
		v5027 = v4210
		v5039 = v4222
		goto L5
	} else {
		goto L359
	}
L358:
	;
	goto L357
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4196
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4188
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4189
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4191
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4330)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4333)
	v4359 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	v4367 = F_LWLockAcquire(m, v4359+int32(2944), int32(0))
	mBase = m.M
	v4368 = m.ExcPending
	if v4368 != 0 {
		v5027 = v4210
		v5039 = v4222
		goto L5
	} else {
		goto L360
	}
L360:
	;
	v4370 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[20]))
	v4371 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4370)+36)) = uint8(v4371)
	*(*int32)(unsafe.Add(mBase, uint32(v4370)+12)) = v4371
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v2247
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v2248
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4196
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4188
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4189
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4191
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4210
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4330)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v2240
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4333)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v2210
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v2219
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v2214
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v2215
	v4394 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v2202
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v2216
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v2205
	F_LWLockRelease(m, v4394+int32(2944))
	mBase = m.M
	v4402 = m.ExcPending
	if v4402 != 0 {
		v5027 = v4210
		v5039 = v4222
		goto L5
	} else {
		goto L361
	}
L361:
	;
	v4404 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v4404)+32)) = int32(1)
	v4410 = v4188
	v4411 = v4189
	v4413 = v4191
	v4418 = v4196
	v4428 = v4210
	v4432 = v4210
	v4444 = v4222
	goto L172
L362:
	;
	v2239 = int32(0)
	v2240 = v4488
	v2241 = v4450
	v2242 = v4451
	v2244 = v4453
	v2249 = v4458
	v2259 = v4468
	v2261 = v4467
	v2263 = v4472
	v2275 = v4484
	goto L169
L363:
	;
	goto L364
L364:
	;
	goto L170
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	v4585 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	v4593 = F_LWLockAcquire(m, v4585+int32(2816), int32(0))
	mBase = m.M
	v4594 = m.ExcPending
	if v4594 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L366
	}
L366:
	;
	v4597 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[17]))
	v4618 = v4597
	v4621 = int32(0)
	goto L367
L367:
	;
	v4640 = v4618 + v4621*int32(20)
	v4641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4640)+40)))
	if v4641 != int32(1) {
		v4916 = v4618
		goto L369
	} else {
		goto L370
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	v4941 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	F_LWLockRelease(m, v4941+int32(2816))
	mBase = m.M
	v4949 = m.ExcPending
	if v4949 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L393
	}
L369:
	;
	v4919 = v4621 + int32(1)
	if v4919 != int32(256) {
		v4618 = v4916
		v4621 = v4919
		goto L367
	} else {
		goto L392
	}
L370:
	;
	v4645 = v4640 + int32(36)
	v4646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4645)+5)))
	if v4646 != 0 {
		v4916 = v4618
		goto L369
	} else {
		goto L371
	}
L371:
	;
	v4647 = *(*int32)(unsafe.Add(mBase, uint32(v4645)+8))
	v4649 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[4]))
	if v4647 != v4649 {
		v4916 = v4618
		goto L369
	} else {
		goto L372
	}
L372:
	;
	v4651 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4645)+5)) = uint8(v4651)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	v4672 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	F_LWLockRelease(m, v4672+int32(2816))
	mBase = m.M
	v4680 = m.ExcPending
	if v4680 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L373
	}
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	v4703 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v4704 = m.ExcPending
	if v4704 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L374
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	F_PushActiveSnapshot(m, v4703)
	mBase = m.M
	v4728 = m.ExcPending
	if v4728 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L375
	}
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	F_perform_work_item(m, v4645)
	mBase = m.M
	v4752 = m.ExcPending
	if v4752 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L376
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	v4776 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[30]))
	goto L377
L377:
	;
	if v4776 != int32(0) {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	F_PopActiveSnapshot(m)
	mBase = m.M
	v4802 = m.ExcPending
	if v4802 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	v4804 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[13]))
	if v4804 != 0 {
		goto L382
	} else {
		goto L383
	}
L381:
	;
	goto L380
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	F_ProcessInterrupts(m)
	mBase = m.M
	v4828 = m.ExcPending
	if v4828 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L385
	}
L383:
	;
	goto L384
L384:
	;
	v4830 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[15]))
	if v4830 != 0 {
		goto L386
	} else {
		goto L387
	}
L385:
	;
	goto L384
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[15])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v4858 = m.ExcPending
	if v4858 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L389
	}
L387:
	;
	goto L388
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	v4902 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	v4910 = F_LWLockAcquire(m, v4902+int32(2816), int32(0))
	mBase = m.M
	v4911 = m.ExcPending
	if v4911 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L391
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	F_VacuumUpdateCosts(m)
	mBase = m.M
	v4882 = m.ExcPending
	if v4882 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L390
	}
L390:
	;
	goto L388
L391:
	;
	v4912 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v4645)+4)) = uint16(v4912)
	v4915 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[17]))
	v4916 = v4915
	goto L369
L392:
	;
	goto L368
L393:
	;
	if v4535&int32(1) == int32(0) {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	F_vac_update_datfrozenxid(m)
	mBase = m.M
	v4977 = m.ExcPending
	if v4977 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L397
	}
L395:
	;
	goto L396
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+588)) = v4505
	*(*int32)(unsafe.Add(mBase, uint32(v46)+584)) = v4504
	*(*int32)(unsafe.Add(mBase, uint32(v46)+592)) = v4506
	*(*int32)(unsafe.Add(mBase, uint32(v46)+596)) = v4498
	*(*int32)(unsafe.Add(mBase, uint32(v46)+600)) = v4499
	*(*int32)(unsafe.Add(mBase, uint32(v46)+604)) = v4501
	*(*int32)(unsafe.Add(mBase, uint32(v46)+608)) = v4516
	*(*int32)(unsafe.Add(mBase, uint32(v46)+616)) = v4496
	*(*int32)(unsafe.Add(mBase, uint32(v46)+624)) = v4508
	*(*int32)(unsafe.Add(mBase, uint32(v46)+628)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v46)+632)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v46)+636)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v46)+640)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v46)+644)) = v4507
	*(*int32)(unsafe.Add(mBase, uint32(v46)+648)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v46)+652)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v46)+656)) = v4514
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)) = uint8(v4555)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)) = uint8(v4558)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+664)) = v4502
	*(*int32)(unsafe.Add(mBase, uint32(v46)+660)) = v4500
	*(*int32)(unsafe.Add(mBase, uint32(v46)+668)) = v4503
	F_CommitTransactionCommand(m)
	mBase = m.M
	v5001 = m.ExcPending
	if v5001 != 0 {
		v5027 = v4520
		v5039 = v4532
		goto L5
	} else {
		goto L398
	}
L397:
	;
	goto L396
L398:
	;
	goto L4
L399:
	;
	v5047 = int32(v5043)
	m.G0 = v46
	v5049 = *(*int32)(unsafe.Add(mBase, uint32(v5047)+4))
	v5050 = *(*int32)(unsafe.Add(mBase, uint32(v5047)))
	v5053 = *(*int32)(unsafe.Add(mBase, uint32(v5050)))
	if v46+int32(92) == v5053 {
		goto L402
	} else {
		goto L403
	}
L400:
	;
	m.ExcPending = 1
	goto L408
L401:
	;
	if v5057 != 0 {
		goto L405
	} else {
		goto L406
	}
L402:
	;
	v5055 = *(*int32)(unsafe.Add(mBase, uint32(v5050)+4))
	v5057 = v5055
	goto L404
L403:
	;
	v5057 = int32(0)
	goto L404
L404:
	;
	goto L401
L405:
	;
	v5058 = *(*int32)(unsafe.Add(mBase, uint32(v46)+668))
	v5059 = *(*int32)(unsafe.Add(mBase, uint32(v46)+664))
	v5060 = *(*int32)(unsafe.Add(mBase, uint32(v46)+660))
	v5061 = *(*int32)(unsafe.Add(mBase, uint32(v46)+656))
	v5062 = *(*int32)(unsafe.Add(mBase, uint32(v46)+652))
	v5063 = *(*int32)(unsafe.Add(mBase, uint32(v46)+648))
	v5064 = *(*int32)(unsafe.Add(mBase, uint32(v46)+644))
	v5065 = *(*int32)(unsafe.Add(mBase, uint32(v46)+640))
	v5066 = *(*int32)(unsafe.Add(mBase, uint32(v46)+636))
	v5067 = *(*int32)(unsafe.Add(mBase, uint32(v46)+632))
	v5068 = *(*int32)(unsafe.Add(mBase, uint32(v46)+628))
	v5069 = *(*int32)(unsafe.Add(mBase, uint32(v46)+624))
	v5070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+623)))
	v5071 = *(*int32)(unsafe.Add(mBase, uint32(v46)+616))
	v5072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+615)))
	v5073 = *(*int32)(unsafe.Add(mBase, uint32(v46)+608))
	v5074 = *(*int32)(unsafe.Add(mBase, uint32(v46)+604))
	v5075 = *(*int32)(unsafe.Add(mBase, uint32(v46)+600))
	v5076 = *(*int32)(unsafe.Add(mBase, uint32(v46)+596))
	v5077 = *(*int32)(unsafe.Add(mBase, uint32(v46)+592))
	v5078 = *(*int32)(unsafe.Add(mBase, uint32(v46)+588))
	v5079 = *(*int32)(unsafe.Add(mBase, uint32(v46)+584))
	v47 = v5071
	v49 = v5076
	v50 = v5075
	v51 = v5060
	v52 = v5074
	v53 = v5059
	v54 = v5058
	v55 = v5079
	v56 = v5078
	v57 = v5077
	v58 = v5064
	v59 = v5069
	v60 = v5068
	v61 = v5067
	v62 = v5065
	v63 = v5063
	v64 = v5062
	v65 = v5061
	v66 = v5066
	v67 = v5073
	v69 = v5057
	v71 = v5027
	v73 = v5072
	v78 = v5070
	v82 = v5049
	v83 = v5039
	goto L1
L406:
	;
	goto L407
L407:
	;
	F___wasm_longjmp(m, v5050, v5049)
	mBase = m.M
	v5081 = m.ExcPending
	if v5081 != 0 {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	return
L409:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
