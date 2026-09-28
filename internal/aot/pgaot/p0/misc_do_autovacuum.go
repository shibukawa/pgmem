package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_do_autovacuum(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v153 int32
	_ = v153
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v197 int64
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v232 int32
	_ = v232
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 float64
	_ = v738
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v938 int32
	_ = v938
	var v962 int32
	_ = v962
	var v968 int32
	_ = v968
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1047 int32
	_ = v1047
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1145 int32
	_ = v1145
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1204 int32
	_ = v1204
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 float64
	_ = v1250
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1356 int32
	_ = v1356
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1403 int32
	_ = v1403
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1437 int32
	_ = v1437
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1509 int32
	_ = v1509
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1597 int32
	_ = v1597
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1712 int32
	_ = v1712
	var v1716 int32
	_ = v1716
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1758 int32
	_ = v1758
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1804 int32
	_ = v1804
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1867 int32
	_ = v1867
	var v1894 int32
	_ = v1894
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1944 int32
	_ = v1944
	var v1977 int32
	_ = v1977
	var v2001 int32
	_ = v2001
	var v2025 int32
	_ = v2025
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2107 float64
	_ = v2107
	var v2111 float64
	_ = v2111
	var v2115 float64
	_ = v2115
	var v2119 float64
	_ = v2119
	var v2123 float64
	_ = v2123
	var v2150 int32
	_ = v2150
	var v2170 int32
	_ = v2170
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2196 int32
	_ = v2196
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2212 int32
	_ = v2212
	var v2214 int32
	_ = v2214
	var v2218 int32
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2230 int32
	_ = v2230
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
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
	var v2246 int32
	_ = v2246
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2254 int32
	_ = v2254
	var v2259 int32
	_ = v2259
	var v2267 int32
	_ = v2267
	var v2269 int32
	_ = v2269
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2274 int32
	_ = v2274
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2312 int32
	_ = v2312
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2319 int32
	_ = v2319
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2341 int32
	_ = v2341
	var v2347 int32
	_ = v2347
	var v2349 int32
	_ = v2349
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2373 int32
	_ = v2373
	var v2381 int32
	_ = v2381
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2403 int32
	_ = v2403
	var v2409 int64
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2417 int32
	_ = v2417
	var v2441 int32
	_ = v2441
	var v2461 int32
	_ = v2461
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2490 int32
	_ = v2490
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2506 int32
	_ = v2506
	var v2509 int32
	_ = v2509
	var v2511 int32
	_ = v2511
	var v2513 int32
	_ = v2513
	var v2555 int32
	_ = v2555
	var v2558 int32
	_ = v2558
	var v2560 int32
	_ = v2560
	var v2562 int32
	_ = v2562
	var v2625 int32
	_ = v2625
	var v2633 int32
	_ = v2633
	var v2635 int32
	_ = v2635
	var v2657 int32
	_ = v2657
	var v2665 int32
	_ = v2665
	var v2668 int32
	_ = v2668
	var v2695 int32
	_ = v2695
	var v2696 int32
	_ = v2696
	var v2699 int32
	_ = v2699
	var v2700 int32
	_ = v2700
	var v2724 int32
	_ = v2724
	var v2725 int32
	_ = v2725
	var v2726 int32
	_ = v2726
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2781 int32
	_ = v2781
	var v2801 int32
	_ = v2801
	var v2809 int32
	_ = v2809
	var v2826 int32
	_ = v2826
	var v2828 int32
	_ = v2828
	var v2831 int32
	_ = v2831
	var v2834 int32
	_ = v2834
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2846 int32
	_ = v2846
	var v2879 int32
	_ = v2879
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2889 int32
	_ = v2889
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2927 int32
	_ = v2927
	var v2930 int32
	_ = v2930
	var v2935 int32
	_ = v2935
	var v2939 int32
	_ = v2939
	var v2941 int32
	_ = v2941
	var v2943 int32
	_ = v2943
	var v2945 int32
	_ = v2945
	var v2947 int32
	_ = v2947
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2953 int32
	_ = v2953
	var v2955 int32
	_ = v2955
	var v2956 int32
	_ = v2956
	var v2959 int32
	_ = v2959
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
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
	var v2978 int32
	_ = v2978
	var v2981 int32
	_ = v2981
	var v2982 int32
	_ = v2982
	var v2983 int32
	_ = v2983
	var v2984 int32
	_ = v2984
	var v2985 int32
	_ = v2985
	var v2986 int32
	_ = v2986
	var v2989 int32
	_ = v2989
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2993 int32
	_ = v2993
	var v2994 int32
	_ = v2994
	var v2995 int32
	_ = v2995
	var v2996 int32
	_ = v2996
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3024 int32
	_ = v3024
	var v3027 int32
	_ = v3027
	var v3038 int32
	_ = v3038
	var v3044 int32
	_ = v3044
	var v3048 int32
	_ = v3048
	var v3052 int32
	_ = v3052
	var v3055 int32
	_ = v3055
	var v3059 float64
	_ = v3059
	var v3061 int32
	_ = v3061
	var v3063 float64
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3070 int32
	_ = v3070
	var v3073 float64
	_ = v3073
	var v3079 float64
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3093 int32
	_ = v3093
	var v3096 int32
	_ = v3096
	var v3128 int32
	_ = v3128
	var v3152 int32
	_ = v3152
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
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3266 int32
	_ = v3266
	var v3274 int32
	_ = v3274
	var v3275 int32
	_ = v3275
	var v3276 int32
	_ = v3276
	var v3278 int32
	_ = v3278
	var v3279 int32
	_ = v3279
	var v3280 int32
	_ = v3280
	var v3284 int32
	_ = v3284
	var v3287 int32
	_ = v3287
	var v3313 int32
	_ = v3313
	var v3328 int32
	_ = v3328
	var v3329 int32
	_ = v3329
	var v3333 int32
	_ = v3333
	var v3334 int32
	_ = v3334
	var v3363 int32
	_ = v3363
	var v3399 int32
	_ = v3399
	var v3407 int32
	_ = v3407
	var v3431 int32
	_ = v3431
	var v3451 int32
	_ = v3451
	var v3457 int32
	_ = v3457
	var v3458 int32
	_ = v3458
	var v3481 int32
	_ = v3481
	var v3482 int32
	_ = v3482
	var v3484 int32
	_ = v3484
	var v3491 int32
	_ = v3491
	var v3509 int32
	_ = v3509
	var v3510 int32
	_ = v3510
	var v3533 int32
	_ = v3533
	var v3534 int32
	_ = v3534
	var v3541 int32
	_ = v3541
	var v3557 int32
	_ = v3557
	var v3562 int32
	_ = v3562
	var v3563 int32
	_ = v3563
	var v3566 int32
	_ = v3566
	var v3567 int32
	_ = v3567
	var v3570 int32
	_ = v3570
	var v3577 int32
	_ = v3577
	var v3579 int32
	_ = v3579
	var v3581 int32
	_ = v3581
	var v3589 int32
	_ = v3589
	var v3594 int32
	_ = v3594
	var v3600 int32
	_ = v3600
	var v3620 int32
	_ = v3620
	var v3621 int32
	_ = v3621
	var v3624 int32
	_ = v3624
	var v3633 int32
	_ = v3633
	var v3641 int32
	_ = v3641
	var v3642 int32
	_ = v3642
	var v3660 int32
	_ = v3660
	var v3661 int32
	_ = v3661
	var v3664 int32
	_ = v3664
	var v3674 int32
	_ = v3674
	var v3675 int32
	_ = v3675
	var v3693 int32
	_ = v3693
	var v3694 int32
	_ = v3694
	var v3697 int32
	_ = v3697
	var v3703 int32
	_ = v3703
	var v3704 int32
	_ = v3704
	var v3705 int32
	_ = v3705
	var v3706 int32
	_ = v3706
	var v3707 int32
	_ = v3707
	var v3732 int32
	_ = v3732
	var v3742 int32
	_ = v3742
	var v3743 int32
	_ = v3743
	var v3767 int32
	_ = v3767
	var v3771 int64
	_ = v3771
	var v3815 int32
	_ = v3815
	var v3825 int32
	_ = v3825
	var v3826 int32
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3828 int32
	_ = v3828
	var v3831 int32
	_ = v3831
	var v3832 int32
	_ = v3832
	var v3856 int32
	_ = v3856
	var v3857 int32
	_ = v3857
	var v3858 int32
	_ = v3858
	var v3882 int32
	_ = v3882
	var v3883 int32
	_ = v3883
	var v3912 int32
	_ = v3912
	var v3913 int32
	_ = v3913
	var v3940 int32
	_ = v3940
	var v3964 int32
	_ = v3964
	var v3972 int32
	_ = v3972
	var v3974 int32
	_ = v3974
	var v3975 int32
	_ = v3975
	var v3978 int32
	_ = v3978
	var v3987 int32
	_ = v3987
	var v3991 int32
	_ = v3991
	var v4007 int32
	_ = v4007
	var v4008 int32
	_ = v4008
	var v4009 int32
	_ = v4009
	var v4010 int32
	_ = v4010
	var v4040 int32
	_ = v4040
	var v4044 int32
	_ = v4044
	var v4068 int32
	_ = v4068
	var v4092 int32
	_ = v4092
	var v4116 int32
	_ = v4116
	var v4136 int32
	_ = v4136
	var v4142 int32
	_ = v4142
	var v4166 int32
	_ = v4166
	var v4167 int32
	_ = v4167
	var v4169 int32
	_ = v4169
	var v4187 int32
	_ = v4187
	var v4190 int32
	_ = v4190
	var v4194 int32
	_ = v4194
	var v4195 int32
	_ = v4195
	var v4197 int32
	_ = v4197
	var v4202 int32
	_ = v4202
	var v4216 int32
	_ = v4216
	var v4225 int32
	_ = v4225
	var v4233 int32
	_ = v4233
	var v4251 int32
	_ = v4251
	var v4252 int32
	_ = v4252
	var v4255 int32
	_ = v4255
	var v4261 int32
	_ = v4261
	var v4262 int32
	_ = v4262
	var v4280 int32
	_ = v4280
	var v4281 int32
	_ = v4281
	var v4284 int32
	_ = v4284
	var v4290 int32
	_ = v4290
	var v4291 int32
	_ = v4291
	var v4309 int32
	_ = v4309
	var v4310 int32
	_ = v4310
	var v4313 int32
	_ = v4313
	var v4319 int32
	_ = v4319
	var v4337 int32
	_ = v4337
	var v4338 int32
	_ = v4338
	var v4341 int32
	_ = v4341
	var v4347 int32
	_ = v4347
	var v4367 int32
	_ = v4367
	var v4375 int32
	_ = v4375
	var v4376 int32
	_ = v4376
	var v4378 int32
	_ = v4378
	var v4379 int32
	_ = v4379
	var v4402 int32
	_ = v4402
	var v4410 int32
	_ = v4410
	var v4412 int32
	_ = v4412
	var v4415 int32
	_ = v4415
	var v4419 int32
	_ = v4419
	var v4420 int32
	_ = v4420
	var v4422 int32
	_ = v4422
	var v4427 int32
	_ = v4427
	var v4436 int32
	_ = v4436
	var v4441 int32
	_ = v4441
	var v4450 int32
	_ = v4450
	var v4460 int32
	_ = v4460
	var v4461 int32
	_ = v4461
	var v4462 int32
	_ = v4462
	var v4464 int32
	_ = v4464
	var v4469 int32
	_ = v4469
	var v4478 int32
	_ = v4478
	var v4483 int32
	_ = v4483
	var v4492 int32
	_ = v4492
	var v4501 int32
	_ = v4501
	var v4502 int32
	_ = v4502
	var v4509 int32
	_ = v4509
	var v4511 int32
	_ = v4511
	var v4512 int32
	_ = v4512
	var v4513 int32
	_ = v4513
	var v4514 int32
	_ = v4514
	var v4515 int32
	_ = v4515
	var v4516 int32
	_ = v4516
	var v4517 int32
	_ = v4517
	var v4518 int32
	_ = v4518
	var v4519 int32
	_ = v4519
	var v4520 int32
	_ = v4520
	var v4521 int32
	_ = v4521
	var v4522 int32
	_ = v4522
	var v4523 int32
	_ = v4523
	var v4524 int32
	_ = v4524
	var v4525 int32
	_ = v4525
	var v4526 int32
	_ = v4526
	var v4527 int32
	_ = v4527
	var v4528 int32
	_ = v4528
	var v4529 int32
	_ = v4529
	var v4533 int32
	_ = v4533
	var v4536 int32
	_ = v4536
	var v4537 int32
	_ = v4537
	var v4542 int32
	_ = v4542
	var v4550 int32
	_ = v4550
	var v4568 int32
	_ = v4568
	var v4569 int32
	_ = v4569
	var v4572 int32
	_ = v4572
	var v4578 int32
	_ = v4578
	var v4598 int32
	_ = v4598
	var v4606 int32
	_ = v4606
	var v4607 int32
	_ = v4607
	var v4610 int32
	_ = v4610
	var v4613 int32
	_ = v4613
	var v4634 int32
	_ = v4634
	var v4655 int32
	_ = v4655
	var v4656 int32
	_ = v4656
	var v4660 int32
	_ = v4660
	var v4661 int32
	_ = v4661
	var v4662 int32
	_ = v4662
	var v4664 int32
	_ = v4664
	var v4666 int32
	_ = v4666
	var v4687 int32
	_ = v4687
	var v4695 int32
	_ = v4695
	var v4718 int32
	_ = v4718
	var v4719 int32
	_ = v4719
	var v4743 int32
	_ = v4743
	var v4767 int32
	_ = v4767
	var v4791 int32
	_ = v4791
	var v4817 int32
	_ = v4817
	var v4819 int32
	_ = v4819
	var v4843 int32
	_ = v4843
	var v4845 int32
	_ = v4845
	var v4873 int32
	_ = v4873
	var v4897 int32
	_ = v4897
	var v4917 int32
	_ = v4917
	var v4925 int32
	_ = v4925
	var v4926 int32
	_ = v4926
	var v4927 int32
	_ = v4927
	var v4930 int32
	_ = v4930
	var v4931 int32
	_ = v4931
	var v4934 int32
	_ = v4934
	var v4956 int32
	_ = v4956
	var v4964 int32
	_ = v4964
	var v4967 int32
	_ = v4967
	var v4996 int32
	_ = v4996
	var v5020 int32
	_ = v5020
	var v5046 int32
	_ = v5046
	var v5055 int32
	_ = v5055
	var v5063 int32
	_ = v5063
	var v5064 int64
	_ = v5064
	var v5068 int32
	_ = v5068
	var v5070 int32
	_ = v5070
	var v5071 int32
	_ = v5071
	var v5074 int32
	_ = v5074
	var v5076 int32
	_ = v5076
	var v5078 int32
	_ = v5078
	var v5079 int32
	_ = v5079
	var v5080 int32
	_ = v5080
	var v5081 int32
	_ = v5081
	var v5082 int32
	_ = v5082
	var v5083 int32
	_ = v5083
	var v5084 int32
	_ = v5084
	var v5085 int32
	_ = v5085
	var v5086 int32
	_ = v5086
	var v5087 int32
	_ = v5087
	var v5088 int32
	_ = v5088
	var v5089 int32
	_ = v5089
	var v5090 int32
	_ = v5090
	var v5091 int32
	_ = v5091
	var v5092 int32
	_ = v5092
	var v5093 int32
	_ = v5093
	var v5094 int32
	_ = v5094
	var v5095 int32
	_ = v5095
	var v5096 int32
	_ = v5096
	var v5097 int32
	_ = v5097
	var v5098 int32
	_ = v5098
	var v5099 int32
	_ = v5099
	var v5100 int32
	_ = v5100
	var v5102 int32
	_ = v5102
	v1 = int32(0)
	v43 = m.G0
	v45 = v43 - int32(848)
	m.G0 = v45
	v48 = v45
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
	v68 = v1
	v71 = int32(-1)
	v73 = v1
	v76 = v1
	v77 = v1
	v81 = v1
	v82 = v1
	goto L1
L1:
	;
	goto L3
L2:
	;
	m.G0 = v48 + int32(848)
	return
L3:
	;
	if v71 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	goto L2
L5:
	;
	v5063 = int32(m.ExcTag)
	v5064 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v5063 == int32(0) {
		goto L398
	} else {
		goto L399
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	v4568 = int32(1)
	v4569 = v4536 & v4568
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	v4572 = v4537 & v4568
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	F_list_free_deep(m, v4529)
	mBase = m.M
	v4578 = m.ExcPending
	if v4578 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L364
	}
L7:
	;
	v2269 = v2267
	v2271 = v51
	v2272 = v52
	v2274 = v54
	v2277 = v57
	v2278 = v58
	v2279 = v59
	v2288 = v68
	v2291 = v2248
	v2292 = v2249
	v2293 = v2250
	v2301 = v81
	v2302 = v2259
	goto L169
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+452)) = int32(0)
	v2107 = *(*float64)(unsafe.Add(mBase, _c_F_do_autovacuum[0]))
	if base.F64_ne(v2107, float64(0)) != 0 {
		goto L158
	} else {
		goto L159
	}
L9:
	;
	v2230 = v53
	v2232 = v55
	v2233 = v56
	v2237 = v60
	v2238 = v61
	v2239 = v62
	v2240 = v63
	v2241 = v64
	v2242 = v65
	v2243 = v66
	v2244 = v67
	v2246 = v50
	v2248 = v76
	v2249 = v49
	v2250 = v68
	v2254 = v77
	v2259 = v82
	v2267 = int32(1)
	goto L7
L10:
	;
	goto L9
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	v105 = int32(1)
	v106 = v76 & v105
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	v109 = v77 & v105
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v56
	v125 = F_AllocSetContextCreateInternal(m, v112, int32(_a_F_do_autovacuum_0), int32(0), int32(_a_F_do_autovacuum_1), int32(_a_F_do_autovacuum_2))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v125
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[3])) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v64
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v56
	F_StartTransactionCommand(m)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v56
	v176 = F_MultiXactMemberFreezeThreshold(m)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	v197 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_do_autovacuum[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v203 = F_SearchSysCache1(m, int32(21), v197)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L16
	}
L16:
	;
	if v203 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v203)+16))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+22)))
	v293 = v291 + v292
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+77)))
	if v294 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v48)+64)) = v252
	F_errmsg_internal(m, int32(_a_F_do_autovacuum_3), v48-int32(-64))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_errfinish(m, int32(_a_F_do_autovacuum_4), int32(1978), int32(_a_F_do_autovacuum_5))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		v5046 = v73
		v5055 = v82
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
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[5])) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_ReleaseCatCache(m, v203)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L29
	}
L24:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[7])) = v310
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[9])) = v314
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[11])) = v318
	v321 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[12]))
	v322 = v321
	goto L23
L25:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+78)))
	if v297 != 0 {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v299 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[9])) = v299
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[7])) = v299
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[11])) = v299
	v322 = v299
	goto L23
L28:
	;
	goto L27
L29:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v64
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v376 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v376)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v401 = F_CreateTupleDescCopy(m, v378)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v48)+464)) = int64(481036337156)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v64
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v432 = F_hash_create(m, int32(_a_F_do_autovacuum_6), int64(100), v48+int32(456), int32(40))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v456 = int32(0)
	v458 = F_table_beginscan_catalog(m, v376, v456, v456)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v482 = F_heap_getnext(m, v458)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v484 = int32(0)
	if v482 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v502 = v64
	v504 = v66
	v505 = v67
	v507 = v484
	v510 = v484
	v513 = v482
	goto L38
L36:
	;
	v886 = v64
	v888 = v66
	v889 = v67
	v891 = v484
	v894 = v484
	goto L37
L37:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v458)))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v912)+188))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v913)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	m.T0[v914].(func(*base.Module, int32))(m, v458)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L72
	}
L38:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v513)+16))
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+22)))
	v530 = v528 + v529
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530)+119)))
	switch v531 - int32(109) {
	case 0, 5:
		goto L41
	default:
		v838 = v504
		v839 = v505
		v840 = v507
		v841 = v510
		goto L40
	}
L39:
	;
	v886 = v868
	v888 = v838
	v889 = v839
	v891 = v840
	v894 = v841
	goto L37
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v839
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v868 = F_heap_getnext(m, v458)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L70
	}
L41:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530)+118)))
	if v535 == int32(116) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v530)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v561 = F_checkTempNamespaceStatus(m, v538)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v505
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v612 = F_extractRelOptions(m, v513, v401, int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L49
	}
L45:
	;
	if v561 != int32(1) {
		v838 = v504
		v839 = v505
		v840 = v507
		v841 = v510
		goto L40
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v505
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v587 = F_lappend_oid(m, v510, v534)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v838 = v504
	v839 = v587
	v840 = v507
	v841 = v587
	goto L40
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v505
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_relation_needs_vacanalyze(m, v534, v670, v530, v176, int32(12), v48+int32(391), v48+int32(390), v48+int32(389), v48+int32(336))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L55
	}
L49:
	;
	if v612 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v670 = int32(0)
	goto L48
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v505
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v640 = F_palloc(m, int32(96))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L53
	}
L53:
	;
	base.MemoryCopy(m, v640, v612+int32(16), int32(96))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_pfree(m, v612)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v670 = v640
	goto L48
L55:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+391)))
	if v704 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v530)+112))
	if v767 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+390)))
	if v707&int32(1) == int32(0) {
		v764 = v504
		v765 = v507
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v505
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v735 = F_palloc(m, int32(16))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v735))) = v534
	v738 = *(*float64)(unsafe.Add(mBase, uint32(v48)+336))
	*(*float64)(unsafe.Add(mBase, uint32(v735)+8)) = v738
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v762 = F_lappend(m, v507, v735)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v764 = v762
	v765 = v762
	goto L56
L63:
	;
	if v670 == int32(0) {
		v838 = v764
		v839 = v505
		v840 = v765
		v841 = v510
		goto L40
	} else {
		goto L68
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v764
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v797 = F_hash_search(m, v432, v530+int32(112), int32(1), v48+int32(335))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+335)))
	if v799 != 0 {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v800 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v797)+8)) = uint8(v800)
	*(*int32)(unsafe.Add(mBase, uint32(v797)+4)) = v534
	if v670 == v800 {
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v805 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v797)+8)) = uint8(v805)
	base.MemoryCopy(m, v797+int32(16), v670, int32(96))
	goto L63
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v764
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v505
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_pfree(m, v670)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L69
	}
L69:
	;
	v838 = v764
	v839 = v505
	v840 = v765
	v841 = v510
	goto L40
L70:
	;
	if v868 != 0 {
		v502 = v868
		v504 = v838
		v505 = v839
		v507 = v840
		v510 = v841
		v513 = v868
		goto L38
	} else {
		goto L71
	}
L71:
	;
	goto L39
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v962 = v48 + int32(392)
	F_ScanKeyInit(m, v962, int32(18), int32(3), int32(61), int64(116))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v992 = F_table_beginscan_catalog(m, v376, int32(1), v962)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1016 = F_heap_getnext(m, v992)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L75
	}
L75:
	;
	if v1016 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v1033 = v63
	v1035 = v65
	v1039 = v891
	v1047 = v1016
	goto L79
L77:
	;
	v1350 = v63
	v1352 = v65
	v1356 = v891
	goto L78
L78:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+188))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	m.T0[v1379].(func(*base.Module, int32))(m, v992)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L105
	}
L79:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+16))
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1060)+22)))
	v1062 = v1060 + v1061
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062)+118)))
	if v1063 == int32(116) {
		v1305 = v1035
		v1306 = v1039
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v1350 = v1333
	v1352 = v1305
	v1356 = v1306
	goto L78
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1033
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1305
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1333 = F_heap_getnext(m, v992)
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L103
	}
L82:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1062)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+328)) = v1066
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1033
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1091 = F_extractRelOptions(m, v1047, v401, int32(0))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1033
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v48)+328))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_relation_needs_vacanalyze(m, v1204, v1184, v1062, v176, int32(12), v48+int32(327), v48+int32(326), v48+int32(325), v48+int32(272))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L95
	}
L84:
	;
	if v1091 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1033
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1116 = F_palloc(m, int32(96))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1033
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1173 = F_hash_search(m, v432, v48+int32(328), int32(0), v48+int32(271))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L90
	}
L88:
	;
	base.MemoryCopy(m, v1116, v1091+int32(16), int32(96))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1033
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_pfree(m, v1091)
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L89
	}
L89:
	;
	v1184 = v1116
	goto L83
L90:
	;
	v1176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+271)))
	if v1176 != int32(1) {
		v1184 = int32(0)
		goto L83
	} else {
		goto L91
	}
L91:
	;
	v1182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1173)+8)))
	if v1182 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v1183 = v1173 + int32(16)
	goto L94
L93:
	;
	v1183 = int32(0)
	goto L94
L94:
	;
	v1184 = v1183
	goto L83
L95:
	;
	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+327)))
	if v1220 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1033
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1246 = F_palloc(m, int32(16))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L99
	}
L97:
	;
	v1276 = v1035
	v1277 = v1039
	goto L98
L98:
	;
	if v1091 == int32(0) {
		v1305 = v1276
		v1306 = v1277
		goto L81
	} else {
		goto L101
	}
L99:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v48)+328))
	*(*int32)(unsafe.Add(mBase, uint32(v1246))) = v1248
	v1250 = *(*float64)(unsafe.Add(mBase, uint32(v48)+272))
	*(*float64)(unsafe.Add(mBase, uint32(v1246)+8)) = v1250
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1033
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1274 = F_lappend(m, v1039, v1246)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L100
	}
L100:
	;
	v1276 = v1274
	v1277 = v1274
	goto L98
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1033
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1276
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_pfree(m, v1184)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L102
	}
L102:
	;
	v1305 = v1276
	v1306 = v1277
	goto L81
L103:
	;
	if v1333 != 0 {
		v1033 = v1333
		v1035 = v1305
		v1039 = v1306
		v1047 = v1333
		goto L79
	} else {
		goto L104
	}
L104:
	;
	goto L80
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_relation_close(m, v376, int32(1))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L106
	}
L106:
	;
	if v894 == int32(0) {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	v1431 = int32(0)
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v894)+4))
	if v1432 <= v1431 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	v1437 = v1431
	goto L109
L109:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v894)+12))
	v1480 = v1477 + v1437<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+452)) = v1480
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1480)))
	v1485 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[13]))
	if v1485 != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L8
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_ProcessInterrupts(m)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1533 = F_ConditionalLockRelationOid(m, v1483, int32(8))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L116
	}
L114:
	;
	goto L113
L115:
	;
	v2058 = v1437 + int32(1)
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v894)+4))
	if v2058 < v2059 {
		v1437 = v2058
		goto L109
	} else {
		goto L156
	}
L116:
	;
	if v1533 == int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1562 = F_SearchSysCacheCopy(m, int32(57), base.I64_extend_i32_u(v1483), int64(0))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L118
	}
L118:
	;
	if v1562 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_UnlockRelationOid(m, v1483, int32(8))
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v1562)+16))
	v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1591)+22)))
	v1593 = v1591 + v1592
	v1594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593)+119)))
	switch v1594 - int32(109) {
	case 0, 5:
		goto L125
	default:
		goto L124
	}
L122:
	;
	goto L115
L123:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1648 = F_checkTempNamespaceStatus(m, v1625)
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L128
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_UnlockRelationOid(m, v1483, int32(8))
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L127
	}
L125:
	;
	v1597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593)+118)))
	if v1597 == int32(116) {
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	goto L115
L128:
	;
	if v1648 != int32(1) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_UnlockRelationOid(m, v1483, int32(8))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1700 = m.G0
	v1702 = v1700 - int32(32)
	m.G0 = v1702
	v1704 = int32(264)
	*(*uint16)(unsafe.Add(mBase, uint32(v1702)+30)) = uint16(v1704)
	v1706 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1702)+28)) = uint16(v1706)
	*(*int32)(unsafe.Add(mBase, uint32(v1702)+24)) = v1677
	*(*int32)(unsafe.Add(mBase, uint32(v1702)+20)) = int32(2615)
	v1712 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v1702)+16)) = v1712
	v1716 = int32(1)
	v1722 = F_LockAcquireExtended(m, v1702+int32(16), v1716, v1706, v1716, v1702+int32(12), v1706)
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L135
	}
L132:
	;
	goto L115
L133:
	;
	m.G0 = v1702 + int32(32)
	if v1722 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L134:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L136
	}
L135:
	;
	switch v1722 {
	case 0, 3:
		goto L133
	default:
		goto L134
	}
L136:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1702)+12))
	v1727 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1726)+53)) = uint8(v1727)
	goto L137
L137:
	;
	goto L133
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_UnlockRelationOid(m, v1483, int32(8))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1783 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L142
	}
L141:
	;
	goto L115
L142:
	;
	if v1783 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	v1804 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1809 = F_get_database_name(m, v1804)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1919 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L150
	}
L146:
	;
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v1834 = F_get_namespace_name(m, v1811)
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v48)+88)) = v1593 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+84)) = v1834
	*(*int32)(unsafe.Add(mBase, uint32(v48)+80)) = v1809
	F_errmsg(m, int32(_a_F_do_autovacuum_7), v48+int32(80))
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_errfinish(m, int32(_a_F_do_autovacuum_4), int32(2282), int32(_a_F_do_autovacuum_5))
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L149
	}
L149:
	;
	goto L145
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_PushActiveSnapshot(m, v1919)
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+264)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+260)) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v48)+256)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_performDeletion(m, v48+int32(256), int32(1), int32(21))
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_StartTransactionCommand(m)
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L155
	}
L155:
	;
	v2052 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v2052
	goto L115
L156:
	;
	goto L110
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	v2170 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v2175 = F_GetAccessStrategyWithSize(m, v2170)
	mBase = m.M
	v2176 = m.ExcPending
	if v2176 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L165
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	F_list_sort(m, v1356, int32(974))
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L164
	}
L159:
	;
	v2111 = *(*float64)(unsafe.Add(mBase, _c_F_do_autovacuum[15]))
	if base.F64_ne(v2111, float64(0)) != 0 {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v2115 = *(*float64)(unsafe.Add(mBase, _c_F_do_autovacuum[16]))
	if base.F64_ne(v2115, float64(0)) != 0 {
		goto L158
	} else {
		goto L161
	}
L161:
	;
	v2119 = *(*float64)(unsafe.Add(mBase, _c_F_do_autovacuum[17]))
	if base.F64_ne(v2119, float64(0)) != 0 {
		goto L158
	} else {
		goto L162
	}
L162:
	;
	v2123 = *(*float64)(unsafe.Add(mBase, _c_F_do_autovacuum[18]))
	if base.F64_eq(v2123, float64(0)) != 0 {
		goto L157
	} else {
		goto L163
	}
L163:
	;
	goto L158
L164:
	;
	goto L157
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2175
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v888
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v106)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v109)
	v2196 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v176
	v2206 = F_AllocSetContextCreateInternal(m, v2196, int32(_a_F_do_autovacuum_8), int32(0), int32(_a_F_do_autovacuum_1), int32(_a_F_do_autovacuum_2))
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		v5046 = v73
		v5055 = v82
		goto L5
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[19])) = v2206
	v2209 = int32(0)
	v2210 = base.B2i32(v432 != v2209)
	v2212 = v1356 + int32(12)
	v2214 = v1356 + int32(4)
	if v1356 == v2209 {
		v4509 = v49
		v4511 = v51
		v4512 = v52
		v4513 = v432
		v4514 = v54
		v4515 = v401
		v4516 = v176
		v4517 = v57
		v4518 = v58
		v4519 = v59
		v4520 = v2214
		v4521 = v2212
		v4522 = v2175
		v4523 = v1350
		v4524 = v886
		v4525 = v1352
		v4526 = v888
		v4527 = v889
		v4528 = v68
		v4529 = v1356
		v4533 = v73
		v4536 = v76
		v4537 = v2210
		v4542 = v82
		v4550 = v2209
		goto L6
	} else {
		goto L167
	}
L167:
	;
	v2218 = int32(0)
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2214)))
	if v2221 <= v2218 {
		v4509 = v49
		v4511 = v51
		v4512 = v52
		v4513 = v432
		v4514 = v54
		v4515 = v401
		v4516 = v176
		v4517 = v57
		v4518 = v58
		v4519 = v59
		v4520 = v2214
		v4521 = v2212
		v4522 = v2175
		v4523 = v1350
		v4524 = v886
		v4525 = v1352
		v4526 = v888
		v4527 = v889
		v4528 = v68
		v4529 = v1356
		v4533 = v73
		v4536 = v76
		v4537 = v2210
		v4542 = v2218
		v4550 = v2218
		goto L6
	} else {
		goto L168
	}
L168:
	;
	v2230 = v432
	v2232 = v401
	v2233 = v176
	v2237 = v2214
	v2238 = v2212
	v2239 = v2175
	v2240 = v1350
	v2241 = v886
	v2242 = v1352
	v2243 = v888
	v2244 = v889
	v2246 = v1356
	v2248 = v2218
	v2249 = v2218
	v2250 = v73
	v2254 = v2210
	v2259 = v2218
	v2267 = int32(0)
	goto L7
L169:
	;
	if v2269 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L170:
	;
	v4509 = v2292
	v4511 = v4461
	v4512 = v4462
	v4513 = v2230
	v4514 = v4464
	v4515 = v2232
	v4516 = v2233
	v4517 = v2277
	v4518 = v2278
	v4519 = v4469
	v4520 = v2237
	v4521 = v2238
	v4522 = v2239
	v4523 = v2240
	v4524 = v2241
	v4525 = v2242
	v4526 = v2243
	v4527 = v2244
	v4528 = v4478
	v4529 = v2246
	v4533 = v4483
	v4536 = v2291
	v4537 = v2254
	v4542 = v4492
	v4550 = base.B2i32(v4492 == int32(0)) & v4460
	goto L6
L171:
	;
	v4501 = v2292 + int32(1)
	v4502 = *(*int32)(unsafe.Add(mBase, uint32(v2237)))
	if v4501 < v4502 {
		goto L361
	} else {
		goto L362
	}
L172:
	;
	v4460 = v2291
	v4461 = v4419
	v4462 = v4420
	v4464 = v4422
	v4469 = v4427
	v4478 = v4436
	v4483 = v4441
	v4492 = v4450
	goto L171
L173:
	;
	if v4233 != 0 {
		goto L346
	} else {
		goto L347
	}
L174:
	;
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(v2238)))
	v2316 = *(*int32)(unsafe.Add(mBase, uint32(v2312+v2292<<(uint(int32(2))%32))))
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(v2316)))
	v2319 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[13]))
	if v2319 != 0 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	goto L176
L176:
	;
	v3589 = v2293 + int32(8)
	if v2301 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	v2337 = int32(1)
	v2338 = v2291 & v2337
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2338)
	v2341 = v2254 & v2337
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2341)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_ProcessInterrupts(m)
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[20]))
	if v2349 != 0 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	goto L179
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[20])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	v2369 = int32(1)
	v2370 = v2291 & v2369
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2370)
	v2373 = v2254 & v2369
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2373)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	v2399 = int32(1)
	v2400 = v2291 & v2399
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	v2403 = v2254 & v2399
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v2409 = base.I64_extend_i32_u(v2317)
	v2410 = F_SearchSysCache1(m, int32(57), v2409)
	mBase = m.M
	v2411 = m.ExcPending
	if v2411 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L185
	}
L184:
	;
	goto L183
L185:
	;
	if v2410 == int32(0) {
		v4419 = v2271
		v4420 = v2272
		v4422 = v2274
		v4427 = v2279
		v4436 = v2288
		v4441 = v2293
		v4450 = v2302
		goto L172
	} else {
		goto L186
	}
L186:
	;
	v2414 = *(*int32)(unsafe.Add(mBase, uint32(v2410)+16))
	v2415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2414)+22)))
	v2417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2414+v2415)+117)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_ReleaseCatCache(m, v2410)
	mBase = m.M
	v2441 = m.ExcPending
	if v2441 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	v2461 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v2469 = F_LWLockAcquire(m, v2461+int32(2944), int32(0))
	mBase = m.M
	v2470 = m.ExcPending
	if v2470 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	v2490 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v2498 = F_LWLockAcquire(m, v2490+int32(2816), int32(1))
	mBase = m.M
	v2499 = m.ExcPending
	if v2499 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L189
	}
L189:
	;
	v2501 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[22]))
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(v2501)+28))
	if v2502 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L190:
	;
	v3231 = *(*float64)(unsafe.Add(mBase, uint32(v3096)+72))
	*(*float64)(unsafe.Add(mBase, _c_F_do_autovacuum[23])) = v3231
	v3234 = *(*int32)(unsafe.Add(mBase, uint32(v3096)+80))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[24])) = v3234
	v3237 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[25]))
	v3238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3096)+84)))
	if v3238 == int32(1) {
		goto L279
	} else {
		goto L280
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v3155
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	v3186 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v3194 = F_LWLockAcquire(m, v3186+int32(2944), int32(0))
	mBase = m.M
	v3195 = m.ExcPending
	if v3195 != 0 {
		v5046 = v3156
		v5055 = v2302
		goto L5
	} else {
		goto L276
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_relation_needs_vacanalyze(m, v2892, v2891, v2726, v2233, int32(12), v48+int32(567), v48+int32(566), v48+int32(565), v48+int32(512))
	mBase = m.M
	v2925 = m.ExcPending
	if v2925 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L226
	}
L193:
	;
	v2843 = int32(0)
	v2846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2726)+119)))
	if (v2254^int32(-1)|base.B2i32(v2846 != int32(116)))&int32(1) != 0 {
		v2891 = v2843
		v2892 = v2317
		goto L192
	} else {
		goto L218
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	v2801 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_LWLockRelease(m, v2801+int32(2816))
	mBase = m.M
	v2809 = m.ExcPending
	if v2809 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L216
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	v2625 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_LWLockRelease(m, v2625+int32(2816))
	mBase = m.M
	v2633 = m.ExcPending
	if v2633 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L208
	}
L196:
	;
	v2506 = v2501 + int32(24)
	if v2502 == v2506 {
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v2509 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[4]))
	v2511 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[25]))
	v2513 = v2502
	goto L198
L198:
	;
	if v2513 == v2511 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	goto L195
L200:
	;
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(v2513)+4))
	if v2562 != v2506 {
		v2513 = v2562
		goto L198
	} else {
		goto L207
	}
L201:
	;
	v2555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2513)+36)))
	if v2555 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(v2513)+8))
	if v2558 != v2509 {
		goto L200
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v2560 = *(*int32)(unsafe.Add(mBase, uint32(v2513)+12))
	if v2560 == v2317 {
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
	v2635 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[25]))
	*(*uint8)(unsafe.Add(mBase, uint32(v2635)+36)) = uint8(v2417)
	*(*int32)(unsafe.Add(mBase, uint32(v2635)+12)) = v2317
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	v2657 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_LWLockRelease(m, v2657+int32(2944))
	mBase = m.M
	v2665 = m.ExcPending
	if v2665 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L209
	}
L209:
	;
	v2668 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v2668
	*(*int32)(unsafe.Add(mBase, uint32(v48)+568)) = v2317
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v2695 = F_SearchSysCacheCopy(m, int32(57), v2409, int64(0))
	mBase = m.M
	v2696 = m.ExcPending
	if v2696 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L210
	}
L210:
	;
	if v2695 == int32(0) {
		v3155 = v2288
		v3156 = v2293
		goto L191
	} else {
		goto L211
	}
L211:
	;
	v2699 = *(*int32)(unsafe.Add(mBase, uint32(v2695)+16))
	v2700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2699)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v2724 = F_extractRelOptions(m, v2695, v2232, int32(0))
	mBase = m.M
	v2725 = m.ExcPending
	if v2725 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L212
	}
L212:
	;
	v2726 = v2699 + v2700
	if v2724 == int32(0) {
		goto L193
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v2752 = F_palloc(m, int32(96))
	mBase = m.M
	v2753 = m.ExcPending
	if v2753 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L214
	}
L214:
	;
	base.MemoryCopy(m, v2752, v2724+int32(16), int32(96))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_pfree(m, v2724)
	mBase = m.M
	v2781 = m.ExcPending
	if v2781 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L215
	}
L215:
	;
	v2891 = v2752
	v2892 = v2317
	goto L192
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	v2826 = int32(1)
	v2828 = v2291 & v2826
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2828)
	v2831 = v2254 & v2826
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2831)
	v2834 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_LWLockRelease(m, v2834+int32(2944))
	mBase = m.M
	v2842 = m.ExcPending
	if v2842 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L217
	}
L217:
	;
	v4460 = v2826
	v4461 = v2271
	v4462 = v2272
	v4464 = v2274
	v4469 = v2279
	v4478 = v2288
	v4483 = v2293
	v4492 = v2302
	goto L171
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v2879 = F_hash_search(m, v2230, v48+int32(568), int32(0), v48+int32(511))
	mBase = m.M
	v2880 = m.ExcPending
	if v2880 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L219
	}
L219:
	;
	v2881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+511)))
	if v2881 == int32(1) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v2887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2879)+8)))
	if v2887 != 0 {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	v2889 = v2843
	goto L222
L222:
	;
	v2890 = *(*int32)(unsafe.Add(mBase, uint32(v48)+568))
	v2891 = v2889
	v2892 = v2890
	goto L192
L223:
	;
	v2888 = v2879 + int32(16)
	goto L225
L224:
	;
	v2888 = int32(0)
	goto L225
L225:
	;
	v2889 = v2888
	goto L222
L226:
	;
	v2926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+567)))
	v2927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+566)))
	if v2927 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	if v2724 != 0 {
		goto L270
	} else {
		goto L271
	}
L228:
	;
	v2930 = int32(0)
	if v2926&int32(1) == v2930 {
		v3096 = v2930
		goto L227
	} else {
		goto L231
	}
L229:
	;
	v2935 = v2293
	goto L230
L230:
	;
	if v2891 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L231:
	;
	v2935 = v2930
	goto L230
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v3020 = F_palloc(m, int32(104))
	mBase = m.M
	v3021 = m.ExcPending
	if v3021 != 0 {
		v5046 = v2990
		v5055 = v2302
		goto L5
	} else {
		goto L253
	}
L233:
	;
	v2989 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[5]))
	v2990 = v2981
	v2991 = v2982
	v2992 = v2983
	v2993 = v2984
	v2994 = v2985
	v2995 = v2986
	v2996 = v2989
	goto L232
L234:
	;
	v2939 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[11]))
	v2941 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[9]))
	v2943 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[7]))
	v2945 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[26]))
	v2947 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[27]))
	v2981 = v2935
	v2982 = v2939
	v2983 = v2943
	v2984 = v2941
	v2985 = v2945
	v2986 = v2947
	goto L233
L235:
	;
	goto L236
L236:
	;
	v2949 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[11]))
	v2950 = *(*int32)(unsafe.Add(mBase, uint32(v2891)+40))
	if v2950 < int32(0) {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v2953 = v2949
	goto L239
L238:
	;
	v2953 = v2950
	goto L239
L239:
	;
	v2955 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[9]))
	v2956 = *(*int32)(unsafe.Add(mBase, uint32(v2891)+36))
	if v2956 < int32(0) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v2959 = v2955
	goto L242
L241:
	;
	v2959 = v2956
	goto L242
L242:
	;
	v2961 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[7]))
	v2962 = *(*int32)(unsafe.Add(mBase, uint32(v2891)+28))
	if v2962 < int32(0) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v2965 = v2961
	goto L245
L244:
	;
	v2965 = v2962
	goto L245
L245:
	;
	v2967 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[26]))
	v2968 = *(*int32)(unsafe.Add(mBase, uint32(v2891)+56))
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
	v2973 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[27]))
	v2974 = *(*int32)(unsafe.Add(mBase, uint32(v2891)+52))
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
	v2978 = *(*int32)(unsafe.Add(mBase, uint32(v2891)+48))
	if int32(0) <= v2978 {
		v2990 = v2974
		v2991 = v2953
		v2992 = v2965
		v2993 = v2959
		v2994 = v2971
		v2995 = v2977
		v2996 = v2978
		goto L232
	} else {
		goto L252
	}
L252:
	;
	v2981 = v2974
	v2982 = v2953
	v2983 = v2965
	v2984 = v2959
	v2985 = v2971
	v2986 = v2977
	goto L233
L253:
	;
	v3022 = *(*int32)(unsafe.Add(mBase, uint32(v48)+568))
	*(*int32)(unsafe.Add(mBase, uint32(v3020))) = v3022
	v3024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+565)))
	*(*int64)(unsafe.Add(mBase, uint32(v3020)+40)) = int64(0)
	v3027 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+64)) = v3027
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+48)) = v3027
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+36)) = v2994
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+32)) = v2995
	*(*uint8)(unsafe.Add(mBase, uint32(v3020)+28)) = uint8(v3024)
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+24)) = v2996
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+20)) = v2991
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+16)) = v2993
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+12)) = v2992
	v3038 = int32(1)
	if v2926&v3038 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v3044 = int32(577)
	goto L256
L255:
	;
	v3044 = v3027
	goto L256
L256:
	;
	if v3024 != 0 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v3048 = int32(0)
	goto L259
L258:
	;
	v3048 = int32(32)
	goto L259
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+8)) = v2927<<(uint(v3038)%32) | v3044 | v3048
	if v2891 != 0 {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3020)+84)) = uint8(v3093)
	v3096 = v3020
	goto L227
L261:
	;
	v3052 = *(*int32)(unsafe.Add(mBase, uint32(v2891)+4))
	if v3052 != 0 {
		goto L265
	} else {
		goto L266
	}
L262:
	;
	goto L263
L263:
	;
	v3079 = *(*float64)(unsafe.Add(mBase, _c_F_do_autovacuum[28]))
	v3080 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+80)) = v3080
	*(*float64)(unsafe.Add(mBase, uint32(v3020)+56)) = v3079
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+96)) = v3080
	*(*int64)(unsafe.Add(mBase, uint32(v3020)+88)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3020)+72)) = int64(-4616189618054758400)
	v3093 = int32(1)
	goto L260
L264:
	;
	v3059 = *(*float64)(unsafe.Add(mBase, _c_F_do_autovacuum[28]))
	*(*float64)(unsafe.Add(mBase, uint32(v3020)+56)) = v3059
	v3061 = *(*int32)(unsafe.Add(mBase, uint32(v2891)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+80)) = v3061
	v3063 = *(*float64)(unsafe.Add(mBase, uint32(v2891)+64))
	v3064 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+96)) = v3064
	*(*int64)(unsafe.Add(mBase, uint32(v3020)+88)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v3020)+72)) = v3063
	v3070 = *(*int32)(unsafe.Add(mBase, uint32(v2891)+24))
	if v3064 < v3070 {
		v3093 = v3064
		goto L260
	} else {
		goto L269
	}
L265:
	;
	if v3052 <= int32(0) {
		goto L264
	} else {
		goto L268
	}
L266:
	;
	v3055 = int32(-1)
	goto L267
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+64)) = v3055
	goto L264
L268:
	;
	v3055 = v3052
	goto L267
L269:
	;
	v3073 = *(*float64)(unsafe.Add(mBase, uint32(v2891)+64))
	v3093 = base.B2i32(base.F64_ge(v3073, float64(0)) == int32(0))
	goto L260
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v3096
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_pfree(m, v2891)
	mBase = m.M
	v3128 = m.ExcPending
	if v3128 != 0 {
		v5046 = v3096
		v5055 = v2302
		goto L5
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v3096
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_pfree(m, v2695)
	mBase = m.M
	v3152 = m.ExcPending
	if v3152 != 0 {
		v5046 = v3096
		v5055 = v2302
		goto L5
	} else {
		goto L274
	}
L273:
	;
	goto L272
L274:
	;
	if v3096 != 0 {
		goto L190
	} else {
		goto L275
	}
L275:
	;
	v3155 = v3096
	v3156 = v3096
	goto L191
L276:
	;
	v3197 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[25]))
	v3198 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3197)+36)) = uint8(v3198)
	*(*int32)(unsafe.Add(mBase, uint32(v3197)+12)) = v3198
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v3155
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	v3221 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_LWLockRelease(m, v3221+int32(2944))
	mBase = m.M
	v3229 = m.ExcPending
	if v3229 != 0 {
		v5046 = v3156
		v5055 = v2302
		goto L5
	} else {
		goto L277
	}
L277:
	;
	v4460 = v2291
	v4461 = v2271
	v4462 = v2272
	v4464 = v2274
	v4469 = v2279
	v4478 = v3155
	v4483 = v3156
	v4492 = v2302
	goto L171
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v3096
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	v3266 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v3274 = F_LWLockAcquire(m, v3266+int32(2816), int32(1))
	mBase = m.M
	v3275 = m.ExcPending
	if v3275 != 0 {
		v5046 = v3096
		v5055 = v2302
		goto L5
	} else {
		goto L282
	}
L279:
	;
	v3243 = base.AtomicRmwXchg32(m, v3237, int32(32), int32(1))
	goto L278
L280:
	;
	goto L281
L281:
	;
	v3244 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3237)+32)), uint32(v3244))
	goto L278
L282:
	;
	v3276 = int32(0)
	v3278 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[22]))
	v3279 = *(*int32)(unsafe.Add(mBase, uint32(v3278)+28))
	v3280 = *(*int32)(unsafe.Add(mBase, uint32(v3278)+uint32(_c_F_do_autovacuum[29])))
	if v3279 == v3276 {
		v3363 = v3276
		goto L283
	} else {
		goto L284
	}
L283:
	;
	if v3363 != v3280 {
		goto L292
	} else {
		goto L293
	}
L284:
	;
	v3284 = v3278 + int32(24)
	if v3279 == v3284 {
		v3363 = v3276
		goto L283
	} else {
		goto L285
	}
L285:
	;
	v3287 = v3279
	v3313 = v3276
	goto L286
L286:
	;
	v3328 = *(*int32)(unsafe.Add(mBase, uint32(v3287)+16))
	if v3328 != 0 {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	v3363 = v3333
	goto L283
L288:
	;
	v3329 = *(*int32)(unsafe.Add(mBase, uint32(v3287)+32))
	v3333 = v3313 + base.B2i32(v3329 != int32(0))
	goto L290
L289:
	;
	v3333 = v3313
	goto L290
L290:
	;
	v3334 = *(*int32)(unsafe.Add(mBase, uint32(v3287)+4))
	if v3334 != v3284 {
		v3287 = v3334
		v3313 = v3333
		goto L286
	} else {
		goto L291
	}
L291:
	;
	goto L287
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3278)+uint32(_c_F_do_autovacuum[29]))) = v3363
	goto L294
L293:
	;
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v3096
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	v3399 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_LWLockRelease(m, v3399+int32(2816))
	mBase = m.M
	v3407 = m.ExcPending
	if v3407 != 0 {
		v5046 = v3096
		v5055 = v2302
		goto L5
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v3096
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_VacuumUpdateCosts(m)
	mBase = m.M
	v3431 = m.ExcPending
	if v3431 != 0 {
		v5046 = v3096
		v5055 = v2302
		goto L5
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v3096
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	v3451 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_MemoryContextReset(m, v3451)
	mBase = m.M
	v3457 = m.ExcPending
	if v3457 != 0 {
		v5046 = v3096
		v5055 = v2302
		goto L5
	} else {
		goto L297
	}
L297:
	;
	v3458 = *(*int32)(unsafe.Add(mBase, uint32(v3096)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v3096
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v3481 = F_get_rel_name(m, v3458)
	mBase = m.M
	v3482 = m.ExcPending
	if v3482 != 0 {
		v5046 = v3096
		v5055 = v2302
		goto L5
	} else {
		goto L298
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3096)+88)) = v3481
	v3484 = *(*int32)(unsafe.Add(mBase, uint32(v3096)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	v3491 = v3096 + int32(88)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v3491
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v3096
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v3509 = F_get_rel_namespace(m, v3484)
	mBase = m.M
	v3510 = m.ExcPending
	if v3510 != 0 {
		v5046 = v3096
		v5055 = v2302
		goto L5
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v3491
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v3096
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v3533 = F_get_namespace_name(m, v3509)
	mBase = m.M
	v3534 = m.ExcPending
	if v3534 != 0 {
		v5046 = v3096
		v5055 = v2302
		goto L5
	} else {
		goto L300
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3096)+92)) = v3533
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	v3541 = v3096 + int32(92)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v3541
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v3491
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v3096
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v2400)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v2403)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	v3557 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v3562 = F_get_database_name(m, v3557)
	mBase = m.M
	v3563 = m.ExcPending
	if v3563 != 0 {
		v5046 = v3096
		v5055 = v2302
		goto L5
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3096)+96)) = v3562
	v3566 = v3096 + int32(96)
	v3567 = *(*int32)(unsafe.Add(mBase, uint32(v3096)+88))
	if v3567 == int32(0) {
		v4194 = v3562
		v4195 = v3541
		v4197 = v3491
		v4202 = v3566
		v4216 = v3096
		v4225 = v2302
		v4233 = v3562
		goto L173
	} else {
		goto L302
	}
L302:
	;
	v3570 = *(*int32)(unsafe.Add(mBase, uint32(v3541)))
	if v3570 == int32(0) {
		v4194 = v3562
		v4195 = v3541
		v4197 = v3491
		v4202 = v3566
		v4216 = v3096
		v4225 = v2302
		v4233 = v3562
		goto L173
	} else {
		goto L303
	}
L303:
	;
	if v3562 == int32(0) {
		v4194 = v3562
		v4195 = v3541
		v4197 = v3491
		v4202 = v3566
		v4216 = v3096
		v4225 = v2302
		v4233 = v3562
		goto L173
	} else {
		goto L304
	}
L304:
	;
	v3577 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[30]))
	v3579 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[31]))
	goto L305
L305:
	;
	v3581 = v48 + int32(96)
	*(*int32)(unsafe.Add(mBase, uint32(v3581)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3581))) = v48 + int32(92)
	goto L308
L306:
	;
	v2269 = int32(1)
	v2271 = v3562
	v2272 = v3541
	v2274 = v3491
	v2277 = v3577
	v2278 = v3579
	v2279 = v3566
	v2293 = v3096
	v2301 = int32(0)
	goto L169
L308:
	;
	goto L306
L309:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[30])) = v2277
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[31])) = v2278
	v4187 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v4187
	v4190 = *(*int32)(unsafe.Add(mBase, uint32(v2279)))
	v4194 = v2271
	v4195 = v2272
	v4197 = v2274
	v4202 = v2279
	v4216 = v2293
	v4225 = int32(1)
	v4233 = v4190
	goto L173
L310:
	;
	v3594 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[19]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v3594
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[31])) = v48 + int32(96)
	v3600 = *(*int32)(unsafe.Add(mBase, uint32(v3589)))
	if v3600&int32(1) != 0 {
		goto L314
	} else {
		goto L315
	}
L311:
	;
	goto L312
L312:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[30])) = v2277
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[31])) = v2278
	v3972 = int32(_a_F_do_autovacuum_9)
	v3974 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[32]))
	v3975 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[32])) = v3974 + v3975
	v3978 = *(*int32)(unsafe.Add(mBase, uint32(v3589)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	v3987 = v2291 & v3975
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3987)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	v3991 = v2254 & v3975
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3991)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v4007 = m.ExcPending
	if v4007 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L336
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	v3693 = int32(1)
	v3694 = v2291 & v3693
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3694)
	v3697 = v2254 & v3693
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3697)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v3703 = v48 + int32(576)
	v3704 = F_strlen(m, v3703)
	mBase = m.M
	v3705 = *(*int32)(unsafe.Add(mBase, uint32(v2272)))
	v3706 = *(*int32)(unsafe.Add(mBase, uint32(v2274)))
	v3707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2293)+28)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3694)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3697)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	if v3707 != 0 {
		goto L322
	} else {
		goto L323
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	v3620 = int32(1)
	v3621 = v2291 & v3620
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3621)
	v3624 = v2254 & v3620
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3624)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	if v3600&int32(2) != 0 {
		goto L317
	} else {
		goto L318
	}
L315:
	;
	goto L316
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	v3660 = int32(1)
	v3661 = v2291 & v3660
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3661)
	v3664 = v2254 & v3660
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3664)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v3674 = F_pg_snprintf(m, v48+int32(576), int32(184), int32(_a_F_do_autovacuum_10), int32(0))
	mBase = m.M
	v3675 = m.ExcPending
	if v3675 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L321
	}
L317:
	;
	v3633 = int32(_a_F_do_autovacuum_11)
	goto L319
L318:
	;
	v3633 = int32(_a_F_do_autovacuum_12)
	goto L319
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = v3633
	v3641 = F_pg_snprintf(m, v48+int32(576), int32(184), int32(_a_F_do_autovacuum_13), v48+int32(32))
	mBase = m.M
	v3642 = m.ExcPending
	if v3642 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L320
	}
L320:
	;
	goto L313
L321:
	;
	goto L313
L322:
	;
	v3732 = int32(_a_F_do_autovacuum_14)
	goto L324
L323:
	;
	v3732 = int32(_a_F_do_autovacuum_12)
	goto L324
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v3732
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v3706
	*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v3705
	v3742 = F_pg_snprintf(m, v3704+v3703, int32(184)-v3704, int32(_a_F_do_autovacuum_15), v48+int32(16))
	mBase = m.M
	v3743 = m.ExcPending
	if v3743 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L325
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3694)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3697)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v3767 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[33]))
	if v3767 < int32(0) {
		goto L327
	} else {
		goto L328
	}
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3694)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3697)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_pgstat_report_activity(m, int32(3), v3703)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3694)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3697)
	v3815 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v3825 = F_AllocSetContextCreateInternal(m, v3815, int32(_a_F_do_autovacuum_16), int32(0), int32(_a_F_do_autovacuum_1), int32(_a_F_do_autovacuum_2))
	mBase = m.M
	v3826 = m.ExcPending
	if v3826 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L330
	}
L327:
	;
	v3771 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_do_autovacuum[34])) = v3771
	goto L329
L328:
	;
	goto L329
L329:
	;
	goto L326
L330:
	;
	v3827 = int32(_a_F_do_autovacuum_17)
	v3828 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v3825
	v3831 = *(*int32)(unsafe.Add(mBase, uint32(v2274)))
	v3832 = *(*int32)(unsafe.Add(mBase, uint32(v2272)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3694)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3697)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v3856 = F_makeRangeVar(m, v3832, v3831, int32(-1))
	mBase = m.M
	v3857 = m.ExcPending
	if v3857 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L331
	}
L331:
	;
	v3858 = *(*int32)(unsafe.Add(mBase, uint32(v2293)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3694)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3697)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v3882 = F_makeVacuumRelation(m, v3856, v3858, int32(0))
	mBase = m.M
	v3883 = m.ExcPending
	if v3883 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L332
	}
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+572)) = v3882
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3694)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3697)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v3882
	v3912 = F_list_make1_impl(m, int32(1), v48+int32(12))
	mBase = m.M
	v3913 = m.ExcPending
	if v3913 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L333
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v3828
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3694)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3697)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_vacuum(m, v3912, v3589, v2239, v3825, int32(1))
	mBase = m.M
	v3940 = m.ExcPending
	if v3940 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3694)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3697)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_MemoryContextDelete(m, v3825)
	mBase = m.M
	v3964 = m.ExcPending
	if v3964 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[35])) = int32(0)
	goto L309
L336:
	;
	v4008 = *(*int32)(unsafe.Add(mBase, uint32(v2279)))
	v4009 = *(*int32)(unsafe.Add(mBase, uint32(v2272)))
	v4010 = *(*int32)(unsafe.Add(mBase, uint32(v2274)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3987)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3991)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	*(*int32)(unsafe.Add(mBase, uint32(v48)+56)) = v4010
	*(*int32)(unsafe.Add(mBase, uint32(v48)+52)) = v4009
	*(*int32)(unsafe.Add(mBase, uint32(v48)+48)) = v4008
	if v3978&int32(1) != 0 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v4040 = int32(_a_F_do_autovacuum_18)
	goto L339
L338:
	;
	v4040 = int32(_a_F_do_autovacuum_19)
	goto L339
L339:
	;
	F_errcontext_msg(m, v4040, v48+int32(48))
	mBase = m.M
	v4044 = m.ExcPending
	if v4044 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L340
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3987)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3991)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_EmitErrorReport(m)
	mBase = m.M
	v4068 = m.ExcPending
	if v4068 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3987)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3991)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_AbortOutOfAnyTransaction(m)
	mBase = m.M
	v4092 = m.ExcPending
	if v4092 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L342
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3987)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3991)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_FlushErrorState(m)
	mBase = m.M
	v4116 = m.ExcPending
	if v4116 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L343
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3987)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3991)
	v4136 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_MemoryContextReset(m, v4136)
	mBase = m.M
	v4142 = m.ExcPending
	if v4142 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L344
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v2279
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v2271
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v2272
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v3987)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v3991)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_StartTransactionCommand(m)
	mBase = m.M
	v4166 = m.ExcPending
	if v4166 != 0 {
		v5046 = v2293
		v5055 = v2302
		goto L5
	} else {
		goto L345
	}
L345:
	;
	v4167 = int32(_a_F_do_autovacuum_9)
	v4169 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[32]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[32])) = v4169 - int32(1)
	goto L309
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4202
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4194
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4195
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4197
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4216
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	v4251 = int32(1)
	v4252 = v2291 & v4251
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4252)
	v4255 = v2254 & v4251
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4255)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_pfree(m, v4233)
	mBase = m.M
	v4261 = m.ExcPending
	if v4261 != 0 {
		v5046 = v4216
		v5055 = v4225
		goto L5
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	v4262 = *(*int32)(unsafe.Add(mBase, uint32(v4195)))
	if v4262 != 0 {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	goto L348
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4202
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4194
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4195
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4197
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4216
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	v4280 = int32(1)
	v4281 = v2291 & v4280
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4281)
	v4284 = v2254 & v4280
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4284)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_pfree(m, v4262)
	mBase = m.M
	v4290 = m.ExcPending
	if v4290 != 0 {
		v5046 = v4216
		v5055 = v4225
		goto L5
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	v4291 = *(*int32)(unsafe.Add(mBase, uint32(v4197)))
	if v4291 != 0 {
		goto L354
	} else {
		goto L355
	}
L353:
	;
	goto L352
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4202
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4194
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4195
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4197
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4216
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	v4309 = int32(1)
	v4310 = v2291 & v4309
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4310)
	v4313 = v2254 & v4309
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4313)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_pfree(m, v4291)
	mBase = m.M
	v4319 = m.ExcPending
	if v4319 != 0 {
		v5046 = v4216
		v5055 = v4225
		goto L5
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4202
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4194
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4195
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4197
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4216
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	v4337 = int32(1)
	v4338 = v2291 & v4337
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4338)
	v4341 = v2254 & v4337
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4341)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_pfree(m, v4216)
	mBase = m.M
	v4347 = m.ExcPending
	if v4347 != 0 {
		v5046 = v4216
		v5055 = v4225
		goto L5
	} else {
		goto L358
	}
L357:
	;
	goto L356
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4202
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4194
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4195
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4197
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4216
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4338)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4341)
	v4367 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	v4375 = F_LWLockAcquire(m, v4367+int32(2944), int32(0))
	mBase = m.M
	v4376 = m.ExcPending
	if v4376 != 0 {
		v5046 = v4216
		v5055 = v4225
		goto L5
	} else {
		goto L359
	}
L359:
	;
	v4378 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[25]))
	v4379 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4378)+36)) = uint8(v4379)
	*(*int32)(unsafe.Add(mBase, uint32(v4378)+12)) = v4379
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v2278
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4202
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4194
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4195
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4197
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4216
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4338)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v2292
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4341)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v2239
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v2246
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v2241
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v2243
	v4402 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v2244
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v2233
	F_LWLockRelease(m, v4402+int32(2944))
	mBase = m.M
	v4410 = m.ExcPending
	if v4410 != 0 {
		v5046 = v4216
		v5055 = v4225
		goto L5
	} else {
		goto L360
	}
L360:
	;
	v4412 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[25]))
	v4415 = base.AtomicRmwXchg32(m, v4412, int32(32), int32(1))
	v4419 = v4194
	v4420 = v4195
	v4422 = v4197
	v4427 = v4202
	v4436 = v4216
	v4441 = v4216
	v4450 = v4225
	goto L172
L361:
	;
	v2269 = int32(0)
	v2271 = v4461
	v2272 = v4462
	v2274 = v4464
	v2279 = v4469
	v2288 = v4478
	v2291 = v4460
	v2292 = v4501
	v2293 = v4483
	v2302 = v4492
	goto L169
L362:
	;
	goto L363
L363:
	;
	goto L170
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	v4598 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	v4606 = F_LWLockAcquire(m, v4598+int32(2816), int32(0))
	mBase = m.M
	v4607 = m.ExcPending
	if v4607 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L365
	}
L365:
	;
	v4610 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[22]))
	v4613 = v4610
	v4634 = int32(0)
	goto L366
L366:
	;
	v4655 = v4613 + v4634*int32(20)
	v4656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4655)+40)))
	if v4656 != int32(1) {
		v4931 = v4613
		goto L368
	} else {
		goto L369
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	v4956 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	F_LWLockRelease(m, v4956+int32(2816))
	mBase = m.M
	v4964 = m.ExcPending
	if v4964 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L392
	}
L368:
	;
	v4934 = v4634 + int32(1)
	if v4934 != int32(256) {
		v4613 = v4931
		v4634 = v4934
		goto L366
	} else {
		goto L391
	}
L369:
	;
	v4660 = v4655 + int32(36)
	v4661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4660)+5)))
	if v4661 != 0 {
		v4931 = v4613
		goto L368
	} else {
		goto L370
	}
L370:
	;
	v4662 = *(*int32)(unsafe.Add(mBase, uint32(v4660)+8))
	v4664 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[4]))
	if v4662 != v4664 {
		v4931 = v4613
		goto L368
	} else {
		goto L371
	}
L371:
	;
	v4666 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4660)+5)) = uint8(v4666)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	v4687 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	F_LWLockRelease(m, v4687+int32(2816))
	mBase = m.M
	v4695 = m.ExcPending
	if v4695 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L372
	}
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	v4718 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v4719 = m.ExcPending
	if v4719 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L373
	}
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	F_PushActiveSnapshot(m, v4718)
	mBase = m.M
	v4743 = m.ExcPending
	if v4743 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L374
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	F_perform_work_item(m, v4660)
	mBase = m.M
	v4767 = m.ExcPending
	if v4767 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L375
	}
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	v4791 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[36]))
	goto L376
L376:
	;
	if v4791 != int32(0) {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	F_PopActiveSnapshot(m)
	mBase = m.M
	v4817 = m.ExcPending
	if v4817 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	v4819 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[13]))
	if v4819 != 0 {
		goto L381
	} else {
		goto L382
	}
L380:
	;
	goto L379
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	F_ProcessInterrupts(m)
	mBase = m.M
	v4843 = m.ExcPending
	if v4843 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L384
	}
L382:
	;
	goto L383
L383:
	;
	v4845 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[20]))
	if v4845 != 0 {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	goto L383
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[20])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v4873 = m.ExcPending
	if v4873 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	v4917 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	v4925 = F_LWLockAcquire(m, v4917+int32(2816), int32(0))
	mBase = m.M
	v4926 = m.ExcPending
	if v4926 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L390
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	F_VacuumUpdateCosts(m)
	mBase = m.M
	v4897 = m.ExcPending
	if v4897 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L389
	}
L389:
	;
	goto L387
L390:
	;
	v4927 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v4660)+4)) = uint16(v4927)
	v4930 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[22]))
	v4931 = v4930
	goto L368
L391:
	;
	goto L367
L392:
	;
	v4967 = *(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[37]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_autovacuum[2])) = v4967
	if v4550&int32(1) == int32(0) {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	F_vac_update_datfrozenxid(m)
	mBase = m.M
	v4996 = m.ExcPending
	if v4996 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+764)) = v4518
	*(*int32)(unsafe.Add(mBase, uint32(v48)+760)) = v4517
	*(*int32)(unsafe.Add(mBase, uint32(v48)+768)) = v4519
	*(*int32)(unsafe.Add(mBase, uint32(v48)+772)) = v4511
	*(*int32)(unsafe.Add(mBase, uint32(v48)+776)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+780)) = v4514
	*(*int32)(unsafe.Add(mBase, uint32(v48)+784)) = v4528
	*(*int32)(unsafe.Add(mBase, uint32(v48)+792)) = v4509
	*(*int32)(unsafe.Add(mBase, uint32(v48)+800)) = v4521
	*(*int32)(unsafe.Add(mBase, uint32(v48)+804)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+808)) = v4522
	*(*int32)(unsafe.Add(mBase, uint32(v48)+812)) = v4529
	*(*int32)(unsafe.Add(mBase, uint32(v48)+816)) = v4523
	*(*int32)(unsafe.Add(mBase, uint32(v48)+820)) = v4525
	*(*int32)(unsafe.Add(mBase, uint32(v48)+824)) = v4524
	*(*int32)(unsafe.Add(mBase, uint32(v48)+828)) = v4526
	*(*int32)(unsafe.Add(mBase, uint32(v48)+832)) = v4527
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)) = uint8(v4569)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)) = uint8(v4572)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+840)) = v4515
	*(*int32)(unsafe.Add(mBase, uint32(v48)+836)) = v4513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+844)) = v4516
	F_CommitTransactionCommand(m)
	mBase = m.M
	v5020 = m.ExcPending
	if v5020 != 0 {
		v5046 = v4533
		v5055 = v4542
		goto L5
	} else {
		goto L397
	}
L396:
	;
	goto L395
L397:
	;
	goto L4
L398:
	;
	v5068 = int32(v5064)
	m.G0 = v48
	v5070 = *(*int32)(unsafe.Add(mBase, uint32(v5068)+4))
	v5071 = *(*int32)(unsafe.Add(mBase, uint32(v5068)))
	v5074 = *(*int32)(unsafe.Add(mBase, uint32(v5071)))
	if v48+int32(92) == v5074 {
		goto L401
	} else {
		goto L402
	}
L399:
	;
	m.ExcPending = 1
	goto L407
L400:
	;
	if v5078 != 0 {
		goto L404
	} else {
		goto L405
	}
L401:
	;
	v5076 = *(*int32)(unsafe.Add(mBase, uint32(v5071)+4))
	v5078 = v5076
	goto L403
L402:
	;
	v5078 = int32(0)
	goto L403
L403:
	;
	goto L400
L404:
	;
	v5079 = *(*int32)(unsafe.Add(mBase, uint32(v48)+844))
	v5080 = *(*int32)(unsafe.Add(mBase, uint32(v48)+840))
	v5081 = *(*int32)(unsafe.Add(mBase, uint32(v48)+836))
	v5082 = *(*int32)(unsafe.Add(mBase, uint32(v48)+832))
	v5083 = *(*int32)(unsafe.Add(mBase, uint32(v48)+828))
	v5084 = *(*int32)(unsafe.Add(mBase, uint32(v48)+824))
	v5085 = *(*int32)(unsafe.Add(mBase, uint32(v48)+820))
	v5086 = *(*int32)(unsafe.Add(mBase, uint32(v48)+816))
	v5087 = *(*int32)(unsafe.Add(mBase, uint32(v48)+812))
	v5088 = *(*int32)(unsafe.Add(mBase, uint32(v48)+808))
	v5089 = *(*int32)(unsafe.Add(mBase, uint32(v48)+804))
	v5090 = *(*int32)(unsafe.Add(mBase, uint32(v48)+800))
	v5091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+799)))
	v5092 = *(*int32)(unsafe.Add(mBase, uint32(v48)+792))
	v5093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+791)))
	v5094 = *(*int32)(unsafe.Add(mBase, uint32(v48)+784))
	v5095 = *(*int32)(unsafe.Add(mBase, uint32(v48)+780))
	v5096 = *(*int32)(unsafe.Add(mBase, uint32(v48)+776))
	v5097 = *(*int32)(unsafe.Add(mBase, uint32(v48)+772))
	v5098 = *(*int32)(unsafe.Add(mBase, uint32(v48)+768))
	v5099 = *(*int32)(unsafe.Add(mBase, uint32(v48)+764))
	v5100 = *(*int32)(unsafe.Add(mBase, uint32(v48)+760))
	v49 = v5092
	v50 = v5087
	v51 = v5097
	v52 = v5096
	v53 = v5081
	v54 = v5095
	v55 = v5080
	v56 = v5079
	v57 = v5100
	v58 = v5099
	v59 = v5098
	v60 = v5089
	v61 = v5090
	v62 = v5088
	v63 = v5086
	v64 = v5084
	v65 = v5085
	v66 = v5083
	v67 = v5082
	v68 = v5094
	v71 = v5078
	v73 = v5046
	v76 = v5093
	v77 = v5091
	v81 = v5070
	v82 = v5055
	goto L1
L405:
	;
	goto L406
L406:
	;
	F___wasm_longjmp(m, v5071, v5070)
	mBase = m.M
	v5102 = m.ExcPending
	if v5102 != 0 {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	return
L408:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
