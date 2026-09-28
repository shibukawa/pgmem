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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v170 int32
	_ = v170
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v208 int32
	_ = v208
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v249 int32
	_ = v249
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v418 int32
	_ = v418
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v516 int32
	_ = v516
	var v534 int32
	_ = v534
	var v553 int32
	_ = v553
	var v572 int32
	_ = v572
	var v592 int32
	_ = v592
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v617 int64
	_ = v617
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v663 int32
	_ = v663
	var v681 int32
	_ = v681
	var v703 int32
	_ = v703
	var v723 int32
	_ = v723
	var v742 int64
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v776 int64
	_ = v776
	var v778 int32
	_ = v778
	var v782 int64
	_ = v782
	var v784 int32
	_ = v784
	var v788 int64
	_ = v788
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v817 int64
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int64
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v828 int64
	_ = v828
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v885 int32
	_ = v885
	var v903 int32
	_ = v903
	var v922 int32
	_ = v922
	var v942 int32
	_ = v942
	var v961 int32
	_ = v961
	var v979 int32
	_ = v979
	var v998 int32
	_ = v998
	var v1017 int32
	_ = v1017
	var v1037 int32
	_ = v1037
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1096 int32
	_ = v1096
	var v1114 int32
	_ = v1114
	var v1133 int32
	_ = v1133
	var v1152 int32
	_ = v1152
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1229 int32
	_ = v1229
	var v1238 int32
	_ = v1238
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1344 int32
	_ = v1344
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1401 int32
	_ = v1401
	var v1419 int32
	_ = v1419
	var v1438 int32
	_ = v1438
	var v1457 int32
	_ = v1457
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1505 int32
	_ = v1505
	var v1523 int32
	_ = v1523
	var v1542 int32
	_ = v1542
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1580 int32
	_ = v1580
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1613 int32
	_ = v1613
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1646 int32
	_ = v1646
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1704 int32
	_ = v1704
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1764 int32
	_ = v1764
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1840 int32
	_ = v1840
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1848 int32
	_ = v1848
	var v1851 int32
	_ = v1851
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1889 int32
	_ = v1889
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1929 int32
	_ = v1929
	var v1948 int32
	_ = v1948
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2018 int64
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2026 int32
	_ = v2026
	var v2047 int32
	_ = v2047
	var v2065 int32
	_ = v2065
	var v2084 int32
	_ = v2084
	var v2103 int32
	_ = v2103
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2127 int32
	_ = v2127
	var v2150 int32
	_ = v2150
	var v2168 int32
	_ = v2168
	var v2187 int32
	_ = v2187
	var v2206 int32
	_ = v2206
	var v2226 int32
	_ = v2226
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2302 int32
	_ = v2302
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2329 int32
	_ = v2329
	var v2335 int32
	_ = v2335
	var v2347 int32
	_ = v2347
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2358 int32
	_ = v2358
	var v2382 int32
	_ = v2382
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2420 int32
	_ = v2420
	var v2423 int32
	_ = v2423
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2434 int32
	_ = v2434
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2466 int32
	_ = v2466
	var v2484 int32
	_ = v2484
	var v2506 int32
	_ = v2506
	var v2526 int32
	_ = v2526
	var v2528 int32
	_ = v2528
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2602 int32
	_ = v2602
	var v2603 int32
	_ = v2603
	var v2622 int32
	_ = v2622
	var v2640 int32
	_ = v2640
	var v2662 int32
	_ = v2662
	var v2682 int32
	_ = v2682
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2738 int32
	_ = v2738
	var v2740 int32
	_ = v2740
	var v2741 int32
	_ = v2741
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2763 int32
	_ = v2763
	var v2801 int32
	_ = v2801
	var v2819 int32
	_ = v2819
	var v2838 int32
	_ = v2838
	var v2858 int32
	_ = v2858
	var v2875 int32
	_ = v2875
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2883 int32
	_ = v2883
	var v2886 int32
	_ = v2886
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2921 int32
	_ = v2921
	var v2924 int32
	_ = v2924
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2959 int32
	_ = v2959
	var v2962 int32
	_ = v2962
	var v2964 int32
	_ = v2964
	var v2970 int32
	_ = v2970
	var v2971 int32
	_ = v2971
	var v2972 int32
	_ = v2972
	var v2993 int32
	_ = v2993
	var v3011 int32
	_ = v3011
	var v3030 int32
	_ = v3030
	var v3050 int32
	_ = v3050
	var v3056 int32
	_ = v3056
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3080 int32
	_ = v3080
	var v3103 int32
	_ = v3103
	var v3121 int32
	_ = v3121
	var v3140 int32
	_ = v3140
	var v3160 int32
	_ = v3160
	var v3166 int32
	_ = v3166
	var v3186 int32
	_ = v3186
	var v3204 int32
	_ = v3204
	var v3221 int32
	_ = v3221
	var v3223 int32
	_ = v3223
	var v3224 int32
	_ = v3224
	var v3245 int32
	_ = v3245
	var v3263 int32
	_ = v3263
	var v3285 int32
	_ = v3285
	var v3305 int32
	_ = v3305
	var v3306 int32
	_ = v3306
	var v3309 int32
	_ = v3309
	var v3332 int32
	_ = v3332
	var v3350 int32
	_ = v3350
	var v3369 int32
	_ = v3369
	var v3388 int32
	_ = v3388
	var v3408 int32
	_ = v3408
	var v3426 int32
	_ = v3426
	var v3442 int32
	_ = v3442
	var v3445 int32
	_ = v3445
	var v3446 int32
	_ = v3446
	var v3467 int32
	_ = v3467
	var v3485 int32
	_ = v3485
	var v3507 int32
	_ = v3507
	var v3527 int32
	_ = v3527
	var v3544 int32
	_ = v3544
	var v3545 int64
	_ = v3545
	var v3566 int32
	_ = v3566
	var v3568 int32
	_ = v3568
	var v3585 int32
	_ = v3585
	var v3586 int32
	_ = v3586
	var v3603 int64
	_ = v3603
	var v3604 int32
	_ = v3604
	var v3607 int64
	_ = v3607
	var v3628 int32
	_ = v3628
	var v3646 int32
	_ = v3646
	var v3662 int64
	_ = v3662
	var v3664 int64
	_ = v3664
	var v3665 int64
	_ = v3665
	var v3669 int64
	_ = v3669
	var v3675 int32
	_ = v3675
	var v3695 int32
	_ = v3695
	var v3714 int32
	_ = v3714
	var v3715 int32
	_ = v3715
	var v3736 int32
	_ = v3736
	var v3756 int32
	_ = v3756
	var v3758 int64
	_ = v3758
	var v3760 int32
	_ = v3760
	var v3763 int32
	_ = v3763
	var v3780 int32
	_ = v3780
	var v3781 int32
	_ = v3781
	var v3805 int32
	_ = v3805
	var v3820 int32
	_ = v3820
	var v3821 int32
	_ = v3821
	var v3828 int32
	_ = v3828
	var v3845 int32
	_ = v3845
	var v3846 int32
	_ = v3846
	var v3886 int32
	_ = v3886
	var v3890 int32
	_ = v3890
	var v3891 int64
	_ = v3891
	var v3893 int32
	_ = v3893
	var v3914 int32
	_ = v3914
	var v3918 int32
	_ = v3918
	var v3933 int32
	_ = v3933
	var v3952 int32
	_ = v3952
	var v3953 int64
	_ = v3953
	var v3955 int32
	_ = v3955
	var v3957 int32
	_ = v3957
	var v3958 int32
	_ = v3958
	var v3961 int32
	_ = v3961
	var v3963 int32
	_ = v3963
	var v3964 int64
	_ = v3964
	var v3966 int32
	_ = v3966
	var v3968 int32
	_ = v3968
	var v3971 int32
	_ = v3971
	var v3972 int32
	_ = v3972
	var v3993 int32
	_ = v3993
	var v4011 int32
	_ = v4011
	var v4033 int32
	_ = v4033
	var v4053 int32
	_ = v4053
	var v4069 int32
	_ = v4069
	var v4071 int32
	_ = v4071
	var v4072 int32
	_ = v4072
	var v4091 int32
	_ = v4091
	var v4109 int32
	_ = v4109
	var v4131 int32
	_ = v4131
	var v4150 int32
	_ = v4150
	var v4170 int32
	_ = v4170
	var v4187 int32
	_ = v4187
	var v4191 int32
	_ = v4191
	var v4193 int32
	_ = v4193
	var v4194 int32
	_ = v4194
	var v4196 int32
	_ = v4196
	var v4197 int32
	_ = v4197
	var v4198 int32
	_ = v4198
	var v4200 int32
	_ = v4200
	var v4203 int32
	_ = v4203
	var v4205 int32
	_ = v4205
	var v4212 int32
	_ = v4212
	var v4213 int32
	_ = v4213
	var v4229 int32
	_ = v4229
	var v4242 int32
	_ = v4242
	var v4243 int32
	_ = v4243
	var v4244 int32
	_ = v4244
	var v4263 int32
	_ = v4263
	var v4264 int32
	_ = v4264
	var v4265 int32
	_ = v4265
	var v4268 int32
	_ = v4268
	var v4287 int32
	_ = v4287
	var v4291 int32
	_ = v4291
	var v4292 int32
	_ = v4292
	var v4295 int32
	_ = v4295
	var v4296 int32
	_ = v4296
	var v4306 int32
	_ = v4306
	var v4315 int32
	_ = v4315
	var v4318 int32
	_ = v4318
	var v4320 int32
	_ = v4320
	var v4329 int32
	_ = v4329
	var v4333 int32
	_ = v4333
	var v4337 int32
	_ = v4337
	var v4338 int32
	_ = v4338
	var v4339 int32
	_ = v4339
	var v4340 int32
	_ = v4340
	var v4341 int32
	_ = v4341
	var v4343 int32
	_ = v4343
	var v4344 int32
	_ = v4344
	var v4364 int32
	_ = v4364
	var v4365 int32
	_ = v4365
	var v4366 int32
	_ = v4366
	var v4367 int32
	_ = v4367
	var v4384 int32
	_ = v4384
	var v4385 int32
	_ = v4385
	var v4392 int32
	_ = v4392
	var v4393 int32
	_ = v4393
	var v4394 int32
	_ = v4394
	var v4395 int32
	_ = v4395
	var v4397 int32
	_ = v4397
	var v4398 int32
	_ = v4398
	var v4409 int32
	_ = v4409
	var v4431 int32
	_ = v4431
	var v4432 int32
	_ = v4432
	var v4451 int32
	_ = v4451
	var v4468 int32
	_ = v4468
	var v4469 int32
	_ = v4469
	var v4470 int32
	_ = v4470
	var v4471 int32
	_ = v4471
	var v4472 int32
	_ = v4472
	var v4473 int32
	_ = v4473
	var v4474 int32
	_ = v4474
	var v4475 int32
	_ = v4475
	var v4477 int32
	_ = v4477
	var v4478 int32
	_ = v4478
	var v4479 int32
	_ = v4479
	var v4480 int32
	_ = v4480
	var v4481 int32
	_ = v4481
	var v4482 int32
	_ = v4482
	var v4483 int32
	_ = v4483
	var v4484 int32
	_ = v4484
	var v4490 int32
	_ = v4490
	var v4491 int32
	_ = v4491
	var v4492 int32
	_ = v4492
	var v4493 int32
	_ = v4493
	var v4495 int32
	_ = v4495
	var v4496 int32
	_ = v4496
	var v4500 int32
	_ = v4500
	var v4501 int32
	_ = v4501
	var v4507 int32
	_ = v4507
	var v4508 int32
	_ = v4508
	var v4509 int32
	_ = v4509
	var v4528 int32
	_ = v4528
	var v4529 int32
	_ = v4529
	var v4532 int32
	_ = v4532
	var v4537 int32
	_ = v4537
	var v4539 int32
	_ = v4539
	var v4542 int32
	_ = v4542
	var v4545 int32
	_ = v4545
	var v4546 int32
	_ = v4546
	var v4548 int32
	_ = v4548
	var v4549 int32
	_ = v4549
	var v4554 int32
	_ = v4554
	var v4567 int32
	_ = v4567
	var v4568 int32
	_ = v4568
	var v4574 int32
	_ = v4574
	var v4575 int32
	_ = v4575
	var v4596 int32
	_ = v4596
	var v4614 int32
	_ = v4614
	var v4615 int32
	_ = v4615
	var v4631 int32
	_ = v4631
	var v4638 int32
	_ = v4638
	var v4658 int32
	_ = v4658
	var v4661 int32
	_ = v4661
	var v4663 int32
	_ = v4663
	var v4665 int32
	_ = v4665
	var v4671 int32
	_ = v4671
	var v4672 int32
	_ = v4672
	var v4673 int32
	_ = v4673
	var v4674 int32
	_ = v4674
	var v4675 int32
	_ = v4675
	var v4676 int32
	_ = v4676
	var v4677 int32
	_ = v4677
	var v4679 int32
	_ = v4679
	var v4680 int32
	_ = v4680
	var v4681 int32
	_ = v4681
	var v4682 int32
	_ = v4682
	var v4683 int32
	_ = v4683
	var v4684 int32
	_ = v4684
	var v4685 int32
	_ = v4685
	var v4686 int32
	_ = v4686
	var v4691 int32
	_ = v4691
	var v4692 int32
	_ = v4692
	var v4693 int32
	_ = v4693
	var v4694 int32
	_ = v4694
	var v4695 int32
	_ = v4695
	var v4698 int32
	_ = v4698
	var v4702 int32
	_ = v4702
	var v4703 int32
	_ = v4703
	var v4714 int32
	_ = v4714
	var v4726 int32
	_ = v4726
	var v4727 int32
	_ = v4727
	var v4731 int32
	_ = v4731
	var v4734 int32
	_ = v4734
	var v4737 int32
	_ = v4737
	var v4738 int32
	_ = v4738
	var v4739 int32
	_ = v4739
	var v4744 int32
	_ = v4744
	var v4745 int32
	_ = v4745
	var v4750 int32
	_ = v4750
	var v4753 int32
	_ = v4753
	var v4761 int32
	_ = v4761
	var v4765 int32
	_ = v4765
	var v4766 int32
	_ = v4766
	var v4768 int32
	_ = v4768
	var v4769 int32
	_ = v4769
	var v4786 int32
	_ = v4786
	var v4788 int32
	_ = v4788
	var v4790 int32
	_ = v4790
	var v4798 int32
	_ = v4798
	var v4799 int32
	_ = v4799
	var v4816 int32
	_ = v4816
	var v4821 int32
	_ = v4821
	var v4822 int32
	_ = v4822
	var v4823 int32
	_ = v4823
	var v4824 int32
	_ = v4824
	var v4825 int32
	_ = v4825
	var v4826 int32
	_ = v4826
	var v4827 int32
	_ = v4827
	var v4829 int32
	_ = v4829
	var v4830 int32
	_ = v4830
	var v4831 int32
	_ = v4831
	var v4832 int32
	_ = v4832
	var v4833 int32
	_ = v4833
	var v4834 int32
	_ = v4834
	var v4835 int32
	_ = v4835
	var v4836 int32
	_ = v4836
	var v4841 int32
	_ = v4841
	var v4843 int32
	_ = v4843
	var v4844 int32
	_ = v4844
	var v4845 int32
	_ = v4845
	var v4848 int32
	_ = v4848
	var v4852 int32
	_ = v4852
	var v4853 int32
	_ = v4853
	var v4870 int32
	_ = v4870
	var v4871 int32
	_ = v4871
	var v4874 int32
	_ = v4874
	var v4877 int32
	_ = v4877
	var v4880 int32
	_ = v4880
	var v4884 int32
	_ = v4884
	var v4886 int32
	_ = v4886
	var v4903 int32
	_ = v4903
	var v4907 int32
	_ = v4907
	var v4924 int32
	_ = v4924
	var v4933 int32
	_ = v4933
	var v4934 int32
	_ = v4934
	var v4939 int32
	_ = v4939
	var v4940 int32
	_ = v4940
	var v4944 int32
	_ = v4944
	var v4947 int32
	_ = v4947
	var v4950 int32
	_ = v4950
	var v4959 int32
	_ = v4959
	var v4976 int32
	_ = v4976
	var v4977 int32
	_ = v4977
	var v4978 int32
	_ = v4978
	var v4979 int32
	_ = v4979
	var v4980 int32
	_ = v4980
	var v4981 int32
	_ = v4981
	var v5008 int32
	_ = v5008
	var v5009 int32
	_ = v5009
	var v5015 int32
	_ = v5015
	var v5016 int64
	_ = v5016
	var v5020 int32
	_ = v5020
	var v5022 int32
	_ = v5022
	var v5023 int32
	_ = v5023
	var v5026 int32
	_ = v5026
	var v5028 int32
	_ = v5028
	var v5030 int32
	_ = v5030
	var v5031 int32
	_ = v5031
	var v5032 int32
	_ = v5032
	var v5033 int32
	_ = v5033
	var v5034 int32
	_ = v5034
	var v5035 int32
	_ = v5035
	var v5036 int32
	_ = v5036
	var v5037 int32
	_ = v5037
	var v5038 int32
	_ = v5038
	var v5039 int32
	_ = v5039
	var v5040 int32
	_ = v5040
	var v5041 int32
	_ = v5041
	var v5042 int32
	_ = v5042
	var v5043 int32
	_ = v5043
	var v5044 int32
	_ = v5044
	var v5045 int32
	_ = v5045
	var v5047 int32
	_ = v5047
	v5 = int32(0)
	v39 = m.G0
	v41 = v39 - int32(848)
	m.G0 = v41
	v48 = l0
	v49 = l1
	v50 = l2
	v51 = l3
	v52 = v41
	v53 = v5
	v54 = v5
	v55 = v5
	v56 = v5
	v57 = v5
	v58 = v5
	v59 = v5
	v60 = v5
	v61 = v5
	v62 = v5
	v63 = v5
	v64 = int32(-1)
	v68 = v5
	v69 = v5
	v70 = v5
	v72 = v5
	v75 = v5
	v79 = v41 + int32(507)
	v80 = v41 + int32(511)
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
	if v64 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	goto L3
L6:
	;
	v5015 = int32(m.ExcTag)
	v5016 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v5015 == int32(0) {
		goto L488
	} else {
		goto L489
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[0])) = v4680
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[1])) = v4681
	v4933 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[2]))
	v4934 = *(*int32)(unsafe.Add(mBase, uint32(v4933)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+796)) = v4680
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+800)) = v4681
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+804)) = v4679
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+808)) = v4683
	v4939 = int32(1)
	v4940 = v4691 & v4939
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+815)) = uint8(v4940)
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+816)) = v4694
	v4944 = v4698 & v4939
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+821)) = uint8(v4944)
	v4947 = v4695 & v4939
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+822)) = uint8(v4947)
	v4950 = v4693 & v4939
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+823)) = uint8(v4950)
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+824)) = v4684
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+828)) = v4685
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+832)) = v4686
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+836)) = v4676
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+840)) = v4677
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+844)) = v4682
	m.T0[v4934].(func(*base.Module, int32))(m, v4679)
	mBase = m.M
	v4959 = m.ExcPending
	if v4959 != 0 {
		v4977 = v4671
		v4978 = v4672
		v4979 = v4673
		v4980 = v4674
		v4981 = v4675
		v5008 = v4702
		v5009 = v4703
		goto L6
	} else {
		goto L486
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+800)) = v4831
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+796)) = v4830
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+804)) = v4829
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+808)) = v4833
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+816)) = v4844
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+824)) = v4834
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+828)) = v4835
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+832)) = v4836
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+836)) = v4826
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+840)) = v4827
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+844)) = v4832
	v4870 = int32(1)
	v4871 = v4841 & v4870
	*(*uint8)(unsafe.Add(mBase, uint32(v4825)+815)) = uint8(v4871)
	v4874 = v4848 & v4870
	*(*uint8)(unsafe.Add(mBase, uint32(v4825)+821)) = uint8(v4874)
	v4877 = v4845 & v4870
	*(*uint8)(unsafe.Add(mBase, uint32(v4825)+822)) = uint8(v4877)
	v4880 = v4843 & v4870
	*(*uint8)(unsafe.Add(mBase, uint32(v4825)+823)) = uint8(v4880)
	F_relation_close(m, v4832, int32(3))
	mBase = m.M
	v4884 = m.ExcPending
	if v4884 != 0 {
		v4977 = v4821
		v4978 = v4822
		v4979 = v4823
		v4980 = v4824
		v4981 = v4825
		v5008 = v4852
		v5009 = v4853
		goto L6
	} else {
		goto L480
	}
L9:
	;
	if v4692 != 0 {
		goto L7
	} else {
		goto L463
	}
L10:
	;
	v4671 = v48
	v4672 = v49
	v4673 = v50
	v4674 = v51
	v4675 = v52
	v4676 = v53
	v4677 = v54
	v4679 = v56
	v4680 = v57
	v4681 = v58
	v4682 = v59
	v4683 = v60
	v4684 = v61
	v4685 = v62
	v4686 = v63
	v4691 = v68
	v4692 = v69
	v4693 = v70
	v4694 = v55
	v4695 = v72
	v4698 = v75
	v4702 = v79
	v4703 = v80
	goto L9
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	v93 = int32(1)
	v94 = v68 & v93
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	v97 = v75 & v93
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	v99 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+536)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v52)+528)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v52)+520)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v52)+512)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v52)+504)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v52)+496)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v52)+488)) = v99
	v114 = v72 & v93
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	v117 = v70 & v93
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v59
	v127 = F_table_open(m, int32(_a_F_AlterSubscription_0), int32(3))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v129 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v50)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v147 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_AlterSubscription[3])))
	v148 = F_SearchSysCacheCopy(m, int32(66), v147, v129)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L14
	}
L14:
	;
	if v148 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v148)+16))
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+22)))
	v231 = v229 + v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v266 = F_object_ownercheck(m, int32(_a_F_AlterSubscription_0), v232, v249)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L22
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(67137668))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v189
	F_errmsg(m, int32(_a_F_AlterSubscription_1), v52)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1563), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L21
	}
L21:
	;
	goto L3
L22:
	;
	if v266 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	F_aclcheck_error(m, int32(2), int32(39), v270)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v298 = int32(0)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v291))|base.B2i32(int32(base.Ui32(int32(889))>>(uint(v291)%32))&int32(1) == v298) == v298 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v291<<(uint(int32(2))%32))+uint32(_c_F_AlterSubscription[5])))
	F_parse_subscription_options(m, v49, v303, v321, v52+int32(488))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v343 = F_GetSubscription(m, v232, int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v345 = int32(0)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	switch v346 {
	case 0:
		goto L34
	default:
		v470 = v63
		v475 = v345
		goto L32
	case 3, 4, 5:
		goto L35
	case 6, 7:
		goto L33
	}
L32:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+48)))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v343)+44))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v343)+68))
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+41)))
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+38)))
	if v480 != 0 {
		goto L63
	} else {
		goto L64
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v468 = F_SubscriptionConninfo(m, v343)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L62
	}
L34:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v52)+488))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v343)+52))
	if v351 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+504)))
	if v347 == int32(0) {
		v470 = v63
		v475 = v345
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+507)))
	v366 = (v352^int32(1))&int32(base.Ui32(v350&int32(512))>>(uint(int32(9))%32)) | int32(base.Ui32(v350&int32(_a_F_AlterSubscription_4))>>(uint(int32(13))%32))
	goto L39
L38:
	;
	v366 = int32(0)
	goto L39
L39:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+512)))
	v369 = v350 & int32(_a_F_AlterSubscription_5)
	v373 = v366 | v367&int32(base.Ui32(v369)>>(uint(int32(14))%32))
	if v350&int32(_a_F_AlterSubscription_6) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v470 = v63
	v475 = v345
	goto L32
L41:
	;
	if v373&int32(1) != 0 {
		goto L33
	} else {
		goto L61
	}
L42:
	;
	if v369 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v379 = v367
	goto L45
L44:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+41)))
	v379 = v378
	goto L45
L45:
	;
	if v379&int32(255) == int32(0) {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v52)+520))
	v403 = v399
	v404 = int32(_a_F_AlterSubscription_7)
	goto L48
L47:
	;
	if (base.B2i32(v441 == int32(0))|v373)&int32(1) != 0 {
		goto L33
	} else {
		goto L60
	}
L48:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	if v407 == v408 {
		v430 = v407
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v441 = int32(0)
	goto L47
L50:
	;
	v432 = int32(1)
	if v430 != 0 {
		v403 = v403 + v432
		v404 = v404 + v432
		goto L48
	} else {
		goto L59
	}
L51:
	;
	if base.Ui32((v407-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v418 = v407 | int32(32)
	goto L54
L53:
	;
	v418 = v407
	goto L54
L54:
	;
	if base.Ui32((v408-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v427 = v408 | int32(32)
	goto L57
L56:
	;
	v427 = v408
	goto L57
L57:
	;
	if v418 == v427 {
		v430 = v418
		goto L50
	} else {
		goto L58
	}
L58:
	;
	v441 = v418 - v427
	goto L47
L59:
	;
	goto L49
L60:
	;
	goto L40
L61:
	;
	goto L40
L62:
	;
	v470 = v468
	v475 = v468
	goto L32
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_LockSharedObject(m, int32(_a_F_AlterSubscription_0), v232, int32(8))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L72
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v496 = F_superuser(m)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L65
	}
L65:
	;
	if v496 != 0 {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(16797828))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errmsg(m, int32(_a_F_AlterSubscription_8), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errhint(m, int32(_a_F_AlterSubscription_9), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1689), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L71
	}
L71:
	;
	goto L3
L72:
	;
	v614 = int32(0)
	base.MemoryFill(m, v52+int32(544), v614, int32(184))
	v617 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+783)) = v617
	*(*int64)(unsafe.Add(mBase, uint32(v52)+776)) = v617
	*(*int64)(unsafe.Add(mBase, uint32(v52)+768)) = v617
	*(*int64)(unsafe.Add(mBase, uint32(v52)+736)) = v617
	*(*int64)(unsafe.Add(mBase, uint32(v52)+744)) = v617
	*(*int64)(unsafe.Add(mBase, uint32(v52)+751)) = v617
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(_a_F_AlterSubscription_0)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	switch v634 {
	case 0:
		goto L90
	case 1:
		goto L88
	case 2:
		goto L87
	case 3:
		goto L86
	case 4, 5:
		goto L85
	case 6:
		goto L84
	case 7:
		goto L83
	case 8:
		goto L89
	case 9:
		goto L82
	default:
		goto L81
	}
L73:
	;
	v4507 = v4491 | v4493
	v4508 = int32(1)
	v4509 = v4507 & v4508
	if v4509|v4495&v4508 == int32(0) {
		v4821 = v4469
		v4822 = v4470
		v4823 = v4471
		v4824 = v4472
		v4825 = v4473
		v4826 = v4474
		v4827 = v4475
		v4829 = v4477
		v4830 = v4478
		v4831 = v4479
		v4832 = v4480
		v4833 = v4481
		v4834 = v4482
		v4835 = v4483
		v4836 = v4484
		v4841 = v4507
		v4843 = v4491
		v4844 = v4492
		v4845 = v4493
		v4848 = v4496
		v4852 = v4500
		v4853 = v4501
		goto L8
	} else {
		goto L443
	}
L74:
	;
	v4409 = *(*int32)(unsafe.Add(mBase, uint32(v127)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v4384
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v4385
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v4431 = F_heap_modify_tuple(m, v148, v4409, v52+int32(544), v52+int32(768), v52+int32(736))
	mBase = m.M
	v4432 = m.ExcPending
	if v4432 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L440
	}
L75:
	;
	v3933 = v3914 & int32(_a_F_AlterSubscription_4)
	if v3933 != 0 {
		goto L381
	} else {
		goto L382
	}
L76:
	;
	v3886 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+744)) = uint8(v3886)
	v3890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+507)))
	if v3890 != 0 {
		goto L378
	} else {
		goto L379
	}
L77:
	;
	v3846 = int32(0)
	v4384 = v3820
	v4385 = v3821
	v4392 = v3828
	v4393 = v3846
	v4394 = v478
	v4395 = v3846
	v4397 = v3845
	v4398 = v479
	goto L74
L78:
	;
	v3805 = int32(0)
	v3820 = v3780
	v3821 = v3781
	v3828 = v3805
	v3845 = v3805
	goto L77
L79:
	;
	v3763 = int32(0)
	v4469 = v48
	v4470 = v49
	v4471 = v50
	v4472 = v51
	v4473 = v52
	v4474 = v343
	v4475 = v232
	v4477 = v56
	v4478 = v57
	v4479 = v58
	v4480 = v127
	v4481 = v60
	v4482 = v61
	v4483 = v62
	v4484 = v470
	v4490 = v3763
	v4491 = v3763
	v4492 = v478
	v4493 = v3763
	v4495 = v3763
	v4496 = v479
	v4500 = v79
	v4501 = v80
	goto L73
L80:
	;
	v3760 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+738)) = uint8(v3760)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+560)) = v3758
	v3780 = v61
	v3781 = v62
	goto L78
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3714 = m.ExcPending
	if v3714 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L375
	}
L82:
	;
	v3545 = *(*int64)(unsafe.Add(mBase, uint32(v52)+528))
	if v3545 == int64(0) {
		goto L364
	} else {
		goto L365
	}
L83:
	;
	v3446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+33)))
	if v3446 == int32(0) {
		goto L356
	} else {
		goto L357
	}
L84:
	;
	v3224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+33)))
	if v3224 == int32(0) {
		goto L339
	} else {
		goto L340
	}
L85:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v2266 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	v2267 = *(*int32)(unsafe.Add(mBase, uint32(v343)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v2283 = F_list_copy(m, v2267)
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L258
	}
L86:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v2018 = F_publicationListToArray(m, v2002)
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L238
	}
L87:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v231)+104))
	if v1902 != 0 {
		goto L228
	} else {
		goto L229
	}
L88:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v231)+104))
	if v1619 != 0 {
		goto L206
	} else {
		goto L207
	}
L89:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v343)+52))
	v1479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+501)))
	v1482 = int32(0)
	if v1478|base.B2i32(v1479&int32(1) == v1482) == v1482 {
		goto L192
	} else {
		goto L193
	}
L90:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+488)))
	if v635&int32(8) != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v52)+492))
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+33)))
	if v638|base.B2i32(v639 != int32(1)) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v52)+496))
	if v750 != 0 {
		goto L106
	} else {
		goto L107
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	if v638 != 0 {
		goto L102
	} else {
		goto L103
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(325))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+80)) = int32(_a_F_AlterSubscription_10)
	F_errmsg(m, int32(_a_F_AlterSubscription_11), v52+int32(80))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1718), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L100
	}
L100:
	;
	goto L3
L101:
	;
	v747 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+754)) = uint8(v747)
	goto L93
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v742 = F_DirectFunctionCall1Coll(m, int32(534), int32(0), base.I64_extend_i32_u(v638))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v745 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+786)) = uint8(v745)
	goto L101
L105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v52)+688)) = v742
	goto L101
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v766 = F_cstring_to_text(m, v750)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v52)+488))
	if v773&int32(128) != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v768 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+755)) = uint8(v768)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+696)) = base.I64_extend_i32_u(v766)
	goto L108
L110:
	;
	v776 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v52)+505)))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+592)) = v776
	v778 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+742)) = uint8(v778)
	goto L112
L111:
	;
	goto L112
L112:
	;
	if v773&int32(256) != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v782 = int64(*(*int8)(unsafe.Add(mBase, uint32(v52)+506)))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+600)) = v782
	v784 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+743)) = uint8(v784)
	goto L115
L114:
	;
	goto L115
L115:
	;
	if v773&int32(1024) != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v788 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v52)+508)))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+616)) = v788
	v790 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+745)) = uint8(v790)
	goto L118
L117:
	;
	goto L118
L118:
	;
	if v773&int32(2048) != 0 {
		goto L124
	} else {
		goto L125
	}
L119:
	;
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+36)))
	if v1173 != int32(101) {
		goto L76
	} else {
		goto L157
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L152
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v1072 = int32(1)
	v1074 = F_logicalrep_workers_find(m, v232, v1072, v1072)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L150
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v1053 = int32(1)
	v1055 = F_logicalrep_workers_find(m, v232, v1053, v1053)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L148
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L143
	}
L124:
	;
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+509)))
	if v795 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v824 = v773
	goto L126
L126:
	;
	if v824&int32(_a_F_AlterSubscription_12) != 0 {
		goto L132
	} else {
		goto L133
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v813 = F_superuser(m)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L130
	}
L128:
	;
	v819 = v773
	v820 = int64(1)
	goto L129
L129:
	;
	v821 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+746)) = uint8(v821)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+624)) = v820
	v824 = v819
	goto L126
L130:
	;
	if v813 == int32(0) {
		goto L123
	} else {
		goto L131
	}
L131:
	;
	v817 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v52)+509)))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v52)+488))
	v819 = v818
	v820 = v817
	goto L129
L132:
	;
	v828 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v52)+510)))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+632)) = v828
	v830 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+747)) = uint8(v830)
	goto L134
L133:
	;
	goto L134
L134:
	;
	v832 = int32(0)
	if v824&int32(512) == v832 {
		v3914 = v824
		v3918 = v832
		goto L75
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+507)))
	v854 = int32(1)
	v855 = v853 ^ v854
	F_CheckAlterSubOption(m, v343, int32(_a_F_AlterSubscription_13), v855&v854, v51)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L136
	}
L136:
	;
	if v853&int32(1) != 0 {
		goto L122
	} else {
		goto L137
	}
L137:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+488)))
	if v862&int32(8) == int32(0) {
		goto L121
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errmsg(m, int32(_a_F_AlterSubscription_14), int32(0))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1802), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L142
	}
L142:
	;
	goto L3
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(16797828))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errmsg(m, int32(_a_F_AlterSubscription_8), int32(0))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errhint(m, int32(_a_F_AlterSubscription_9), int32(0))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1764), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L147
	}
L147:
	;
	goto L3
L148:
	;
	if v1055 != 0 {
		goto L120
	} else {
		goto L149
	}
L149:
	;
	goto L76
L150:
	;
	if v1074 == int32(0) {
		goto L119
	} else {
		goto L151
	}
L151:
	;
	goto L120
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(325))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errmsg(m, int32(_a_F_AlterSubscription_15), int32(0))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errhint(m, int32(_a_F_AlterSubscription_16), int32(0))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1819), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L156
	}
L156:
	;
	goto L3
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v1191 = int32(0)
	v1193 = m.G0
	v1195 = v1193 - int32(240)
	m.G0 = v1195
	v1198 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[6]))
	v1202 = F_LWLockAcquire(m, v1198+int32(2304), int32(1))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L159
	}
L158:
	;
	if v1344 == int32(0) {
		goto L76
	} else {
		goto L186
	}
L159:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[7]))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+4))
	if v1206 <= int32(0) {
		v1344 = v1191
		goto L161
	} else {
		goto L162
	}
L160:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L182
	}
L161:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[6]))
	F_LWLockRelease(m, v1357+int32(2304))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L181
	}
L162:
	;
	v1229 = v1205
	v1238 = v1191
	goto L163
L163:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1229+v1238<<(uint(int32(2))%32))+8))
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1250)+48)))
	if v1251 != int32(1) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v1344 = int32(0)
	goto L161
L165:
	;
	v1313 = v1238 + int32(1)
	v1315 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[7]))
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+4))
	if v1313 < v1316 {
		v1229 = v1315
		v1238 = v1313
		goto L163
	} else {
		goto L180
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1195)+16)) = v1195 + int32(236)
	*(*int32)(unsafe.Add(mBase, uint32(v1195)+20)) = v1195 + int32(232)
	v1261 = v1250 + int32(51)
	v1265 = F_sscanf(m, v1261, int32(_a_F_AlterSubscription_17), v1195+int32(16))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L167
	}
L167:
	;
	if v1265 != int32(2) {
		goto L165
	} else {
		goto L168
	}
L168:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1195)+236))
	if v232 != v1269 {
		goto L165
	} else {
		goto L169
	}
L169:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1195)+232))
	if v1271 == int32(0) {
		goto L160
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1195)+4)) = v1271
	*(*int32)(unsafe.Add(mBase, uint32(v1195))) = v232
	v1277 = v1195 + int32(32)
	v1280 = F_pg_snprintf(m, v1277, int32(200), int32(_a_F_AlterSubscription_17), v1195)
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L171
	}
L171:
	;
	v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1261))))
	v1287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1277))))
	if base.B2i32(v1284 == int32(0))|base.B2i32(v1284 != v1287) != 0 {
		v1305 = v1284
		v1306 = v1287
		goto L173
	} else {
		goto L174
	}
L172:
	;
	if v1305-v1306 != 0 {
		goto L165
	} else {
		goto L179
	}
L173:
	;
	goto L172
L174:
	;
	v1290 = v1261
	v1291 = v1277
	goto L175
L175:
	;
	v1294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+1)))
	v1295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1290)+1)))
	if v1295 == int32(0) {
		v1305 = v1295
		v1306 = v1294
		goto L173
	} else {
		goto L177
	}
L176:
	;
	v1305 = v1295
	v1306 = v1294
	goto L173
L177:
	;
	v1298 = int32(1)
	if v1295 == v1294 {
		v1290 = v1290 + v1298
		v1291 = v1291 + v1298
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	v1344 = int32(1)
	goto L161
L180:
	;
	goto L164
L181:
	;
	m.G0 = v1195 + int32(240)
	goto L158
L182:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L183
	}
L183:
	;
	F_errmsg_internal(m, int32(_a_F_AlterSubscription_18), int32(0))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_AlterSubscription_19), int32(2760), int32(_a_F_AlterSubscription_20))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L185
	}
L185:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(325))
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errmsg(m, int32(_a_F_AlterSubscription_21), int32(0))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errhint(m, int32(_a_F_AlterSubscription_22), int32(0))
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1833), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L191
	}
L191:
	;
	goto L3
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v1563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+48)))
	v1564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+41)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v1580 = int32(1)
	F_CheckSubDeadTupleRetention(m, v1479&v1580, (v1479^int32(-1))&v1580, int32(19), v1564, v1563, int32(0))
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L199
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(325))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errmsg(m, int32(_a_F_AlterSubscription_23), int32(0))
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1976), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L198
	}
L198:
	;
	goto L3
L199:
	;
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+501)))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+584)) = base.I64_extend_i32_u(v1590)
	v1593 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+741)) = uint8(v1593)
	if v1590 != v1593 {
		v3780 = v61
		v3781 = v62
		goto L78
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v1613 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[8])))
	if v1613 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v3780 = v61
	v3781 = v62
	goto L78
L202:
	;
	v1617 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[8])) = uint8(v1617)
	goto L204
L203:
	;
	goto L204
L204:
	;
	goto L201
L205:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v1663 = F_GetForeignServerByName(m, v1646, int32(0))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L210
	}
L206:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	F_deleteDependencyRecordsForSpecific(m, int32(_a_F_AlterSubscription_0), v1620, int32(110), int32(1417), v1619)
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	v1641 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+753)) = uint8(v1641)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+785)) = uint8(v1641)
	goto L205
L209:
	;
	goto L205
L210:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v231)+80))
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1663)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v1684 = F_object_aclcheck(m, int32(1417), v1666, v1665, int64(256))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L211
	}
L211:
	;
	if v1684 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v1663)))
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(v231)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	F_GetUserMappingExtended(m, v1786, v1785)
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L220
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(16797828))
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L216
	}
L216:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v231)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v1740 = F_GetUserNameFromId(m, v1723, int32(0))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L217
	}
L217:
	;
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v52)+100)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v52)+96)) = v1740
	F_errmsg(m, int32(_a_F_AlterSubscription_24), v52+int32(96))
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2033), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L219
	}
L219:
	;
	goto L3
L220:
	;
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v231)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v1820 = F_ForeignServerConnectionString(m, v1804, v1663)
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_load_file(m, int32(_a_F_AlterSubscription_25), int32(0))
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L222
	}
L222:
	;
	v1843 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[2]))
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v1843)+4))
	v1845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+38)))
	if v1845 == int32(1) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+32)))
	v1851 = v1848 ^ int32(1)
	goto L225
L224:
	;
	v1851 = int32(0)
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	m.T0[v1844].(func(*base.Module, int32, int32))(m, v1820, v1851&int32(1))
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L226
	}
L226:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1663)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+484)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+480)) = v1871
	*(*int32)(unsafe.Add(mBase, uint32(v52)+476)) = int32(1417)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int64)(unsafe.Add(mBase, uint32(v52)+672)) = base.I64_extend_i32_u(v1871)
	v1889 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+752)) = uint8(v1889)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	F_recordDependencyOn(m, v48, v52+int32(476), int32(110))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L227
	}
L227:
	;
	v1901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+41)))
	v3820 = v61
	v3821 = v62
	v3828 = v1820
	v3845 = v1901
	goto L77
L228:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	F_deleteDependencyRecordsForSpecific(m, int32(_a_F_AlterSubscription_0), v1903, int32(110), int32(1417), v1902)
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	F_load_file(m, int32(_a_F_AlterSubscription_25), int32(0))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L232
	}
L231:
	;
	v1924 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+752)) = uint8(v1924)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+672)) = int64(0)
	goto L230
L232:
	;
	v1951 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[2]))
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1951)+4))
	v1953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+38)))
	if v1953 == int32(1) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+32)))
	v1959 = v1956 ^ int32(1)
	goto L235
L234:
	;
	v1959 = int32(0)
	goto L235
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	m.T0[v1952].(func(*base.Module, int32, int32))(m, v1929, v1959&int32(1))
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L236
	}
L236:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v1995 = F_cstring_to_text(m, v1979)
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L237
	}
L237:
	;
	v1997 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+753)) = uint8(v1997)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+680)) = base.I64_extend_i32_u(v1995)
	v2001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+41)))
	v3820 = v61
	v3821 = v62
	v3828 = v1929
	v3845 = v2001
	goto L77
L238:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v52)+712)) = v2018
	v2021 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+757)) = uint8(v2021)
	v2023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+504)))
	if v2023 != v2021 {
		v3780 = v61
		v3781 = v62
		goto L78
	} else {
		goto L239
	}
L239:
	;
	v2026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+33)))
	if v2026 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v2124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+36)))
	if v2124 != int32(101) {
		goto L248
	} else {
		goto L249
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(325))
	mBase = m.M
	v2065 = m.ExcPending
	if v2065 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errmsg(m, int32(_a_F_AlterSubscription_26), int32(0))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errhint(m, int32(_a_F_AlterSubscription_27), int32(0))
	mBase = m.M
	v2103 = m.ExcPending
	if v2103 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2112), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L247
	}
L247:
	;
	goto L3
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_PreventInTransactionBlock(m, v51, int32(_a_F_AlterSubscription_28))
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L256
	}
L249:
	;
	v2127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+503)))
	if v2127&int32(1) == int32(0) {
		goto L248
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(325))
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errmsg(m, int32(_a_F_AlterSubscription_29), int32(0))
	mBase = m.M
	v2187 = m.ExcPending
	if v2187 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errhint(m, int32(_a_F_AlterSubscription_30), int32(0))
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2122), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v2226 = m.ExcPending
	if v2226 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L255
	}
L255:
	;
	goto L3
L256:
	;
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v343)+64)) = v2245
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v2262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+503)))
	F_AlterSubscription_refresh(m, v343, v2262, v2245, v475)
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L257
	}
L257:
	;
	v3780 = v61
	v3781 = v62
	goto L78
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_check_duplicates_in_publist(m, v2266, int32(0))
	mBase = m.M
	v2302 = m.ExcPending
	if v2302 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L259
	}
L259:
	;
	if v2266 == int32(0) {
		v2756 = v61
		v2757 = v62
		v2763 = v2283
		goto L260
	} else {
		goto L261
	}
L260:
	;
	if v2763 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L261:
	;
	v2305 = int32(0)
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(v2266)+4))
	if v2306 <= v2305 {
		v2756 = v61
		v2757 = v62
		v2763 = v2283
		goto L260
	} else {
		goto L262
	}
L262:
	;
	v2322 = v61
	v2323 = v62
	v2329 = v2283
	v2335 = v2305
	goto L263
L263:
	;
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(v2266)+12))
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(v2347+v2335<<(uint(int32(2))%32))))
	v2352 = *(*int32)(unsafe.Add(mBase, uint32(v2351)+4))
	if v2329 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L264:
	;
	v2756 = v2713
	v2757 = v2714
	v2763 = v2738
	goto L260
L265:
	;
	v2740 = v2335 + int32(1)
	v2741 = *(*int32)(unsafe.Add(mBase, uint32(v2266)+4))
	if v2740 < v2741 {
		v2322 = v2713
		v2323 = v2714
		v2329 = v2738
		v2335 = v2740
		goto L263
	} else {
		goto L298
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v2698 = F_list_delete_nth_cell(m, v2329, v2382)
	mBase = m.M
	v2699 = m.ExcPending
	if v2699 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L297
	}
L267:
	;
	if v634 == int32(4) {
		goto L288
	} else {
		goto L289
	}
L268:
	;
	v2355 = *(*int32)(unsafe.Add(mBase, uint32(v2329)+4))
	if v2355 <= int32(0) {
		goto L267
	} else {
		goto L269
	}
L269:
	;
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v2329)+12))
	v2382 = int32(0)
	goto L270
L270:
	;
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v2358+v2382<<(uint(int32(2))%32))))
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v2401)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v2420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2352))))
	v2423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2402))))
	if base.B2i32(v2420 == int32(0))|base.B2i32(v2420 != v2423) != 0 {
		v2441 = v2420
		v2442 = v2423
		goto L273
	} else {
		goto L274
	}
L271:
	;
	goto L267
L272:
	;
	if v2441-v2442 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L273:
	;
	goto L272
L274:
	;
	v2426 = v2352
	v2427 = v2402
	goto L275
L275:
	;
	v2430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2427)+1)))
	v2431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2426)+1)))
	if v2431 == int32(0) {
		v2441 = v2431
		v2442 = v2430
		goto L273
	} else {
		goto L277
	}
L276:
	;
	v2441 = v2431
	v2442 = v2430
	goto L273
L277:
	;
	v2434 = int32(1)
	if v2431 == v2430 {
		v2426 = v2426 + v2434
		v2427 = v2427 + v2434
		goto L275
	} else {
		goto L278
	}
L278:
	;
	goto L276
L279:
	;
	if v634 != int32(4) {
		goto L266
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v2528 = v2382 + int32(1)
	if v2355 != v2528 {
		v2382 = v2528
		goto L270
	} else {
		goto L287
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2466 = m.ExcPending
	if v2466 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L283
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(_a_F_AlterSubscription_31))
	mBase = m.M
	v2484 = m.ExcPending
	if v2484 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L284
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+164)) = v2265
	*(*int32)(unsafe.Add(mBase, uint32(v52)+160)) = v2352
	F_errmsg(m, int32(_a_F_AlterSubscription_32), v52+int32(160))
	mBase = m.M
	v2506 = m.ExcPending
	if v2506 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(3576), int32(_a_F_AlterSubscription_33))
	mBase = m.M
	v2526 = m.ExcPending
	if v2526 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L286
	}
L286:
	;
	goto L3
L287:
	;
	goto L271
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v2585 = F_makeString(m, v2352)
	mBase = m.M
	v2586 = m.ExcPending
	if v2586 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2622 = m.ExcPending
	if v2622 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L293
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v2602 = F_lappend(m, v2329, v2585)
	mBase = m.M
	v2603 = m.ExcPending
	if v2603 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L292
	}
L292:
	;
	v2713 = v2602
	v2714 = v2323
	v2738 = v2602
	goto L265
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(117833860))
	mBase = m.M
	v2640 = m.ExcPending
	if v2640 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L294
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+148)) = v2265
	*(*int32)(unsafe.Add(mBase, uint32(v52)+144)) = v2352
	F_errmsg(m, int32(_a_F_AlterSubscription_34), v52+int32(144))
	mBase = m.M
	v2662 = m.ExcPending
	if v2662 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(3590), int32(_a_F_AlterSubscription_33))
	mBase = m.M
	v2682 = m.ExcPending
	if v2682 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L296
	}
L296:
	;
	goto L3
L297:
	;
	v2713 = v2322
	v2714 = v2698
	v2738 = v2698
	goto L265
L298:
	;
	goto L264
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2801 = m.ExcPending
	if v2801 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v2875 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[9]))
	v2880 = F_AllocSetContextCreateInternal(m, v2875, int32(_a_F_AlterSubscription_35), int32(0), int32(_a_F_AlterSubscription_4), int32(_a_F_AlterSubscription_36))
	mBase = m.M
	v2881 = m.ExcPending
	if v2881 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L306
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(117833860))
	mBase = m.M
	v2819 = m.ExcPending
	if v2819 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L303
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errmsg(m, int32(_a_F_AlterSubscription_37), int32(0))
	mBase = m.M
	v2838 = m.ExcPending
	if v2838 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L304
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(3600), int32(_a_F_AlterSubscription_33))
	mBase = m.M
	v2858 = m.ExcPending
	if v2858 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L305
	}
L305:
	;
	goto L3
L306:
	;
	v2882 = int32(_a_F_AlterSubscription_38)
	v2883 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[9])) = v2880
	v2886 = *(*int32)(unsafe.Add(mBase, uint32(v2763)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v2903 = F_palloc_mul(m, int32(8), v2886)
	mBase = m.M
	v2904 = m.ExcPending
	if v2904 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L307
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_check_duplicates_in_publist(m, v2763, v2903)
	mBase = m.M
	v2921 = m.ExcPending
	if v2921 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L308
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[9])) = v2883
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v2763)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v2941 = F_construct_array_builtin(m, v2903, v2924, int32(25))
	mBase = m.M
	v2942 = m.ExcPending
	if v2942 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L309
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_MemoryContextDelete(m, v2880)
	mBase = m.M
	v2959 = m.ExcPending
	if v2959 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L310
	}
L310:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v52)+712)) = base.I64_extend_i32_u(v2941)
	v2962 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+757)) = uint8(v2962)
	v2964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+504)))
	if v2964 != v2962 {
		v3780 = v2756
		v3781 = v2757
		goto L78
	} else {
		goto L311
	}
L311:
	;
	if v634 == int32(4) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v2970 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	v2971 = v2970
	goto L314
L313:
	;
	v2971 = int32(0)
	goto L314
L314:
	;
	v2972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+33)))
	if v2972 == int32(0) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2993 = m.ExcPending
	if v2993 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	v3077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+36)))
	if v3077 != int32(101) {
		goto L326
	} else {
		goto L327
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(325))
	mBase = m.M
	v3011 = m.ExcPending
	if v3011 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errmsg(m, int32(_a_F_AlterSubscription_26), int32(0))
	mBase = m.M
	v3030 = m.ExcPending
	if v3030 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L320
	}
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	if v634 == int32(4) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v3050 = int32(_a_F_AlterSubscription_39)
	goto L323
L322:
	;
	v3050 = int32(_a_F_AlterSubscription_40)
	goto L323
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+128)) = v3050
	F_errhint(m, int32(_a_F_AlterSubscription_41), v52+int32(128))
	mBase = m.M
	v3056 = m.ExcPending
	if v3056 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L324
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2164), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v3076 = m.ExcPending
	if v3076 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L325
	}
L325:
	;
	goto L3
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_PreventInTransactionBlock(m, v51, int32(_a_F_AlterSubscription_28))
	mBase = m.M
	v3204 = m.ExcPending
	if v3204 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L337
	}
L327:
	;
	v3080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+503)))
	if v3080&int32(1) == int32(0) {
		goto L326
	} else {
		goto L328
	}
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3103 = m.ExcPending
	if v3103 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L329
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(325))
	mBase = m.M
	v3121 = m.ExcPending
	if v3121 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L330
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errmsg(m, int32(_a_F_AlterSubscription_29), int32(0))
	mBase = m.M
	v3140 = m.ExcPending
	if v3140 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	if v634 == int32(4) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v3160 = int32(_a_F_AlterSubscription_42)
	goto L334
L333:
	;
	v3160 = int32(_a_F_AlterSubscription_43)
	goto L334
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+112)) = v3160
	F_errhint(m, int32(_a_F_AlterSubscription_44), v52+int32(112))
	mBase = m.M
	v3166 = m.ExcPending
	if v3166 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2178), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v3186 = m.ExcPending
	if v3186 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L336
	}
L336:
	;
	goto L3
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v343)+64)) = v2763
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v2756
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v2757
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v3221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+503)))
	F_AlterSubscription_refresh(m, v343, v3221, v2971, v475)
	mBase = m.M
	v3223 = m.ExcPending
	if v3223 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L338
	}
L338:
	;
	v3780 = v2756
	v3781 = v2757
	goto L78
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3245 = m.ExcPending
	if v3245 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L342
	}
L340:
	;
	goto L341
L341:
	;
	v3306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+36)))
	if v3306 != int32(101) {
		goto L346
	} else {
		goto L347
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(325))
	mBase = m.M
	v3263 = m.ExcPending
	if v3263 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L343
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+176)) = int32(_a_F_AlterSubscription_45)
	F_errmsg(m, int32(_a_F_AlterSubscription_46), v52+int32(176))
	mBase = m.M
	v3285 = m.ExcPending
	if v3285 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L344
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2199), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v3305 = m.ExcPending
	if v3305 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L345
	}
L345:
	;
	goto L3
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_PreventInTransactionBlock(m, v51, int32(_a_F_AlterSubscription_45))
	mBase = m.M
	v3426 = m.ExcPending
	if v3426 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L354
	}
L347:
	;
	v3309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+503)))
	if v3309&int32(1) == int32(0) {
		goto L346
	} else {
		goto L348
	}
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3332 = m.ExcPending
	if v3332 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L349
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3350 = m.ExcPending
	if v3350 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errmsg(m, int32(_a_F_AlterSubscription_47), int32(0))
	mBase = m.M
	v3369 = m.ExcPending
	if v3369 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errhint(m, int32(_a_F_AlterSubscription_48), int32(0))
	mBase = m.M
	v3388 = m.ExcPending
	if v3388 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L352
	}
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2222), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v3408 = m.ExcPending
	if v3408 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L353
	}
L353:
	;
	goto L3
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v3442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+503)))
	F_AlterSubscription_refresh(m, v343, v3442, int32(0), v475)
	mBase = m.M
	v3445 = m.ExcPending
	if v3445 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L355
	}
L355:
	;
	goto L79
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3467 = m.ExcPending
	if v3467 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L359
	}
L357:
	;
	goto L358
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_AlterSubscription_refresh_seq(m, v343, v475)
	mBase = m.M
	v3544 = m.ExcPending
	if v3544 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L363
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(325))
	mBase = m.M
	v3485 = m.ExcPending
	if v3485 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L360
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+192)) = int32(_a_F_AlterSubscription_49)
	F_errmsg(m, int32(_a_F_AlterSubscription_46), v52+int32(192))
	mBase = m.M
	v3507 = m.ExcPending
	if v3507 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L361
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2238), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v3527 = m.ExcPending
	if v3527 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L362
	}
L362:
	;
	goto L3
L363:
	;
	goto L79
L364:
	;
	v3758 = int64(0)
	goto L80
L365:
	;
	goto L366
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v3566 = v52 + int32(400)
	F_ReplicationOriginNameForLogicalRep(m, v232, int32(0), v3566)
	mBase = m.M
	v3568 = m.ExcPending
	if v3568 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L367
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v3585 = F_replorigin_by_name(m, v3566, int32(0))
	mBase = m.M
	v3586 = m.ExcPending
	if v3586 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L368
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v3603 = F_replorigin_get_progress(m, v3585, int32(0))
	mBase = m.M
	v3604 = m.ExcPending
	if v3604 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L369
	}
L369:
	;
	v3607 = *(*int64)(unsafe.Add(mBase, uint32(v52)+528))
	if base.B2i32(v3603 == int64(0))|base.B2i32(base.Ui64(v3603) <= base.Ui64(v3607)) != 0 {
		v3758 = v3607
		goto L80
	} else {
		goto L370
	}
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3628 = m.ExcPending
	if v3628 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L371
	}
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3646 = m.ExcPending
	if v3646 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L372
	}
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v3662 = *(*int64)(unsafe.Add(mBase, uint32(v52)+528))
	*(*uint32)(unsafe.Add(mBase, uint32(v52)+212)) = uint32(v3662)
	v3664 = int64(32)
	v3665 = int64(base.Ui64(v3662) >> (uint(v3664) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v52)+208)) = uint32(v3665)
	*(*uint32)(unsafe.Add(mBase, uint32(v52)+220)) = uint32(v3603)
	v3669 = int64(base.Ui64(v3603) >> (uint(v3664) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v52)+216)) = uint32(v3669)
	F_errmsg(m, int32(_a_F_AlterSubscription_50), v52+int32(208))
	mBase = m.M
	v3675 = m.ExcPending
	if v3675 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L373
	}
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2271), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v3695 = m.ExcPending
	if v3695 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L374
	}
L374:
	;
	goto L3
L375:
	;
	v3715 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v3715
	F_errmsg_internal(m, int32(_a_F_AlterSubscription_51), v52+int32(16))
	mBase = m.M
	v3736 = m.ExcPending
	if v3736 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L376
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2283), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v3756 = m.ExcPending
	if v3756 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L377
	}
L377:
	;
	goto L3
L378:
	;
	v3891 = int64(112)
	goto L380
L379:
	;
	v3891 = int64(100)
	goto L380
L380:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v52)+608)) = v3891
	v3893 = *(*int32)(unsafe.Add(mBase, uint32(v52)+488))
	v3914 = v3893
	v3918 = v855
	goto L75
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_CheckAlterSubOption(m, v343, int32(_a_F_AlterSubscription_52), int32(1), v51)
	mBase = m.M
	v3952 = m.ExcPending
	if v3952 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L384
	}
L382:
	;
	v3958 = v3914
	goto L383
L383:
	;
	if v3958&int32(_a_F_AlterSubscription_5) != 0 {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	v3953 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v52)+511)))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+640)) = v3953
	v3955 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+748)) = uint8(v3955)
	v3957 = *(*int32)(unsafe.Add(mBase, uint32(v52)+488))
	v3958 = v3957
	goto L383
L385:
	;
	v3961 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+749)) = uint8(v3961)
	v3963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+512)))
	v3964 = base.I64_extend_i32_u(v3963)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+648)) = v3964
	v3966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+41)))
	if v3966 != v3963 {
		goto L388
	} else {
		goto L389
	}
L386:
	;
	v4196 = v3958
	v4197 = v476
	v4198 = v479
	v4200 = int32(0)
	goto L387
L387:
	;
	if v4196&int32(_a_F_AlterSubscription_53) != 0 {
		goto L413
	} else {
		goto L414
	}
L388:
	;
	v3968 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+751)) = uint8(v3968)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+664)) = v3964
	v3971 = v3963
	goto L390
L389:
	;
	v3971 = v476
	goto L390
L390:
	;
	v3972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+33)))
	if v3972 == int32(1) {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3993 = m.ExcPending
	if v3993 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L394
	}
L392:
	;
	goto L393
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v4069 = int32(1)
	v4071 = F_logicalrep_workers_find(m, v232, v4069, v4069)
	mBase = m.M
	v4072 = m.ExcPending
	if v4072 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L398
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(325))
	mBase = m.M
	v4011 = m.ExcPending
	if v4011 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L395
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+48)) = int32(_a_F_AlterSubscription_54)
	F_errmsg(m, int32(_a_F_AlterSubscription_55), v52+int32(48))
	mBase = m.M
	v4033 = m.ExcPending
	if v4033 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L396
	}
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1498), int32(_a_F_AlterSubscription_56))
	mBase = m.M
	v4053 = m.ExcPending
	if v4053 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L397
	}
L397:
	;
	goto L3
L398:
	;
	if v4071 != 0 {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4091 = m.ExcPending
	if v4091 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L402
	}
L400:
	;
	goto L401
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v4187 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[8])))
	if v4187 == int32(0) {
		goto L408
	} else {
		goto L409
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errcode(m, int32(325))
	mBase = m.M
	v4109 = m.ExcPending
	if v4109 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L403
	}
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+64)) = int32(_a_F_AlterSubscription_54)
	F_errmsg(m, int32(_a_F_AlterSubscription_57), v52-int32(-64))
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L404
	}
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errhint(m, int32(_a_F_AlterSubscription_16), int32(0))
	mBase = m.M
	v4150 = m.ExcPending
	if v4150 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(1904), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v4170 = m.ExcPending
	if v4170 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L406
	}
L406:
	;
	goto L3
L407:
	;
	v4193 = *(*int32)(unsafe.Add(mBase, uint32(v52)+488))
	v4194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+512)))
	v4196 = v4193
	v4197 = v3971
	v4198 = v4194
	v4200 = v4194
	goto L387
L408:
	;
	v4191 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscription[8])) = uint8(v4191)
	goto L410
L409:
	;
	goto L410
L410:
	;
	goto L407
L411:
	;
	if v4244&int32(_a_F_AlterSubscription_6) != 0 {
		goto L418
	} else {
		goto L419
	}
L412:
	;
	v4213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+33)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	v4229 = int32(1)
	F_CheckSubDeadTupleRetention(m, v4229, (v4213^int32(-1))&v4229, int32(18), v4198&v4229, v4197&v4229, base.B2i32(int32(0) < v4212))
	mBase = m.M
	v4242 = m.ExcPending
	if v4242 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L417
	}
L413:
	;
	v4203 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+750)) = uint8(v4203)
	v4205 = *(*int32)(unsafe.Add(mBase, uint32(v52)+516))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+656)) = base.I64_extend_i32_s(v4205)
	v4212 = v4205
	goto L412
L414:
	;
	goto L415
L415:
	;
	if v4196&int32(_a_F_AlterSubscription_5) == int32(0) {
		v4244 = v4196
		goto L411
	} else {
		goto L416
	}
L416:
	;
	v4212 = v477
	goto L412
L417:
	;
	v4243 = *(*int32)(unsafe.Add(mBase, uint32(v52)+488))
	v4244 = v4243
	goto L411
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v4263 = *(*int32)(unsafe.Add(mBase, uint32(v52)+520))
	v4264 = F_cstring_to_text(m, v4263)
	mBase = m.M
	v4265 = m.ExcPending
	if v4265 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L421
	}
L419:
	;
	v4339 = v4244
	v4340 = v478
	v4341 = v4200
	goto L420
L420:
	;
	v4343 = int32(base.Ui32(v3933) >> (uint(int32(13)) % 32))
	v4344 = int32(0)
	if v4339&int32(_a_F_AlterSubscription_58) == v4344 {
		v4384 = v61
		v4385 = v62
		v4392 = v4344
		v4393 = v4343
		v4394 = v4340
		v4395 = v3918
		v4397 = v4341
		v4398 = v4198
		goto L74
	} else {
		goto L438
	}
L421:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v52)+720)) = base.I64_extend_i32_u(v4264)
	v4268 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+758)) = uint8(v4268)
	if v4198&v4268 != 0 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v4287 = *(*int32)(unsafe.Add(mBase, uint32(v52)+520))
	v4291 = v4287
	v4292 = int32(_a_F_AlterSubscription_7)
	goto L426
L423:
	;
	v4333 = int32(0)
	goto L424
L424:
	;
	v4337 = *(*int32)(unsafe.Add(mBase, uint32(v52)+488))
	v4338 = *(*int32)(unsafe.Add(mBase, uint32(v52)+520))
	v4339 = v4337
	v4340 = v4338
	v4341 = v4333 | v4200&int32(1)
	goto L420
L425:
	;
	v4333 = base.B2i32(v4329 == int32(0))
	goto L424
L426:
	;
	v4295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4291))))
	v4296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4292))))
	if v4295 == v4296 {
		v4318 = v4295
		goto L428
	} else {
		goto L429
	}
L427:
	;
	v4329 = int32(0)
	goto L425
L428:
	;
	v4320 = int32(1)
	if v4318 != 0 {
		v4291 = v4291 + v4320
		v4292 = v4292 + v4320
		goto L426
	} else {
		goto L437
	}
L429:
	;
	if base.Ui32((v4295-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v4306 = v4295 | int32(32)
	goto L432
L431:
	;
	v4306 = v4295
	goto L432
L432:
	;
	if base.Ui32((v4296-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v4315 = v4296 | int32(32)
	goto L435
L434:
	;
	v4315 = v4296
	goto L435
L435:
	;
	if v4306 == v4315 {
		v4318 = v4306
		goto L428
	} else {
		goto L436
	}
L436:
	;
	v4329 = v4306 - v4315
	goto L425
L437:
	;
	goto L427
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	v4364 = *(*int32)(unsafe.Add(mBase, uint32(v52)+536))
	v4365 = F_cstring_to_text(m, v4364)
	mBase = m.M
	v4366 = m.ExcPending
	if v4366 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L439
	}
L439:
	;
	v4367 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+756)) = uint8(v4367)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+704)) = base.I64_extend_i32_u(v4365)
	v4384 = v61
	v4385 = v62
	v4392 = v4344
	v4393 = v4343
	v4394 = v4340
	v4395 = v3918
	v4397 = v4341
	v4398 = v4198
	goto L74
L440:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v4384
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v4385
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_CatalogTupleUpdate(m, v127, v4431+int32(4), v4431)
	mBase = m.M
	v4451 = m.ExcPending
	if v4451 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L441
	}
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+800)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v52)+796)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v52)+804)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v52)+808)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v52)+816)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+824)) = v4384
	*(*int32)(unsafe.Add(mBase, uint32(v52)+828)) = v4385
	*(*int32)(unsafe.Add(mBase, uint32(v52)+832)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v52)+836)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v52)+840)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v52)+844)) = v127
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+815)) = uint8(v94)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+821)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+822)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+823)) = uint8(v117)
	F_pfree(m, v4431)
	mBase = m.M
	v4468 = m.ExcPending
	if v4468 != 0 {
		v4977 = v48
		v4978 = v49
		v4979 = v50
		v4980 = v51
		v4981 = v52
		v5008 = v79
		v5009 = v80
		goto L6
	} else {
		goto L442
	}
L442:
	;
	v4469 = v48
	v4470 = v49
	v4471 = v50
	v4472 = v51
	v4473 = v52
	v4474 = v343
	v4475 = v232
	v4477 = v56
	v4478 = v57
	v4479 = v58
	v4480 = v127
	v4481 = v60
	v4482 = v4384
	v4483 = v4385
	v4484 = v470
	v4490 = v4392
	v4491 = v4393
	v4492 = v4394
	v4493 = v4395
	v4495 = v4397
	v4496 = v4398
	v4500 = v79
	v4501 = v80
	goto L73
L443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+796)) = v4478
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+800)) = v4479
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+804)) = v4477
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+808)) = v4481
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+816)) = v4492
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+823)) = uint8(v4491)
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+824)) = v4482
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+828)) = v4483
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+832)) = v4484
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+836)) = v4474
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+840)) = v4475
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+844)) = v4480
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+815)) = uint8(v4509)
	v4528 = int32(1)
	v4529 = v4496 & v4528
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+821)) = uint8(v4529)
	v4532 = v4493 & v4528
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+822)) = uint8(v4532)
	F_load_file(m, int32(_a_F_AlterSubscription_25), int32(0))
	mBase = m.M
	v4537 = m.ExcPending
	if v4537 != 0 {
		v4977 = v4469
		v4978 = v4470
		v4979 = v4471
		v4980 = v4472
		v4981 = v4473
		v5008 = v4500
		v5009 = v4501
		goto L6
	} else {
		goto L444
	}
L444:
	;
	v4539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4474)+38)))
	if v4539 == int32(1) {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v4542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4474)+32)))
	v4545 = v4542 ^ int32(1)
	goto L447
L446:
	;
	v4545 = int32(0)
	goto L447
L447:
	;
	v4546 = *(*int32)(unsafe.Add(mBase, uint32(v4474)+24))
	v4548 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[2]))
	v4549 = *(*int32)(unsafe.Add(mBase, uint32(v4548)))
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+796)) = v4478
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+800)) = v4479
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+804)) = v4477
	v4554 = v4474 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+808)) = v4554
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+815)) = uint8(v4509)
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+816)) = v4492
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+821)) = uint8(v4529)
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+822)) = uint8(v4532)
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+823)) = uint8(v4491)
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+824)) = v4482
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+828)) = v4483
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+832)) = v4484
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+836)) = v4474
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+840)) = v4475
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+844)) = v4480
	if v4490 != 0 {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v4567 = v4490
	goto L450
L449:
	;
	v4567 = v475
	goto L450
L450:
	;
	v4568 = int32(1)
	v4574 = m.T0[v4549].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v4567, v4568, v4568, v4545&v4568, v4546, v4473+int32(396))
	mBase = m.M
	v4575 = m.ExcPending
	if v4575 != 0 {
		v4977 = v4469
		v4978 = v4470
		v4979 = v4471
		v4980 = v4472
		v4981 = v4473
		v5008 = v4500
		v5009 = v4501
		goto L6
	} else {
		goto L451
	}
L451:
	;
	if v4574 == int32(0) {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+800)) = v4479
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+796)) = v4478
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+804)) = v4574
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+808)) = v4554
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+816)) = v4492
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+823)) = uint8(v4491)
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+824)) = v4482
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+828)) = v4483
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+832)) = v4484
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+836)) = v4474
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+840)) = v4475
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+844)) = v4480
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+815)) = uint8(v4509)
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+821)) = uint8(v4529)
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+822)) = uint8(v4532)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4596 = m.ExcPending
	if v4596 != 0 {
		v4977 = v4469
		v4978 = v4470
		v4979 = v4471
		v4980 = v4472
		v4981 = v4473
		v5008 = v4500
		v5009 = v4501
		goto L6
	} else {
		goto L455
	}
L453:
	;
	goto L454
L454:
	;
	v4661 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[0]))
	v4663 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[1]))
	goto L459
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+800)) = v4479
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+796)) = v4478
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+804)) = v4574
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+808)) = v4554
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+816)) = v4492
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+823)) = uint8(v4491)
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+824)) = v4482
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+828)) = v4483
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+832)) = v4484
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+836)) = v4474
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+840)) = v4475
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+844)) = v4480
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+815)) = uint8(v4509)
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+821)) = uint8(v4529)
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+822)) = uint8(v4532)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v4614 = m.ExcPending
	if v4614 != 0 {
		v4977 = v4469
		v4978 = v4470
		v4979 = v4471
		v4980 = v4472
		v4981 = v4473
		v5008 = v4500
		v5009 = v4501
		goto L6
	} else {
		goto L456
	}
L456:
	;
	v4615 = *(*int32)(unsafe.Add(mBase, uint32(v4554)))
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+796)) = v4478
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+800)) = v4479
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+804)) = v4574
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+808)) = v4554
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+815)) = uint8(v4509)
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+816)) = v4492
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+821)) = uint8(v4529)
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+822)) = uint8(v4532)
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+823)) = uint8(v4491)
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+824)) = v4482
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+828)) = v4483
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+832)) = v4484
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+836)) = v4474
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+840)) = v4475
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+844)) = v4480
	v4631 = *(*int32)(unsafe.Add(mBase, uint32(v4473)+396))
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+36)) = v4631
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+32)) = v4615
	F_errmsg(m, int32(_a_F_AlterSubscription_59), v4473+int32(32))
	mBase = m.M
	v4638 = m.ExcPending
	if v4638 != 0 {
		v4977 = v4469
		v4978 = v4470
		v4979 = v4471
		v4980 = v4472
		v4981 = v4473
		v5008 = v4500
		v5009 = v4501
		goto L6
	} else {
		goto L457
	}
L457:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+800)) = v4479
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+796)) = v4478
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+804)) = v4574
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+808)) = v4554
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+816)) = v4492
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+823)) = uint8(v4491)
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+824)) = v4482
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+828)) = v4483
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+832)) = v4484
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+836)) = v4474
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+840)) = v4475
	*(*int32)(unsafe.Add(mBase, uint32(v4473)+844)) = v4480
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+815)) = uint8(v4509)
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+821)) = uint8(v4529)
	*(*uint8)(unsafe.Add(mBase, uint32(v4473)+822)) = uint8(v4532)
	F_errfinish(m, int32(_a_F_AlterSubscription_2), int32(2329), int32(_a_F_AlterSubscription_3))
	mBase = m.M
	v4658 = m.ExcPending
	if v4658 != 0 {
		v4977 = v4469
		v4978 = v4470
		v4979 = v4471
		v4980 = v4472
		v4981 = v4473
		v5008 = v4500
		v5009 = v4501
		goto L6
	} else {
		goto L458
	}
L458:
	;
	goto L3
L459:
	;
	v4665 = v4473 + int32(240)
	*(*int32)(unsafe.Add(mBase, uint32(v4665)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4665))) = v4473 + int32(236)
	goto L462
L460:
	;
	v4671 = v4469
	v4672 = v4470
	v4673 = v4471
	v4674 = v4472
	v4675 = v4473
	v4676 = v4474
	v4677 = v4475
	v4679 = v4574
	v4680 = v4661
	v4681 = v4663
	v4682 = v4480
	v4683 = v4554
	v4684 = v4482
	v4685 = v4483
	v4686 = v4484
	v4691 = v4507
	v4692 = int32(0)
	v4693 = v4491
	v4694 = v4492
	v4695 = v4493
	v4698 = v4496
	v4702 = v4500
	v4703 = v4501
	goto L9
L462:
	;
	goto L460
L463:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[1])) = v4675 + int32(240)
	v4714 = v4698 & int32(1)
	if v4714 != 0 {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+800)) = v4681
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+796)) = v4680
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+804)) = v4679
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+808)) = v4683
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+816)) = v4694
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+824)) = v4684
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+828)) = v4685
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+832)) = v4686
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+836)) = v4676
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+840)) = v4677
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+844)) = v4682
	v4726 = int32(1)
	v4727 = v4691 & v4726
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+815)) = uint8(v4727)
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+821)) = uint8(v4714)
	v4731 = v4695 & v4726
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+822)) = uint8(v4731)
	v4734 = v4693 & v4726
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+823)) = uint8(v4734)
	F_CheckPubDeadTupleRetention(m, v4679)
	mBase = m.M
	v4737 = m.ExcPending
	if v4737 != 0 {
		v4977 = v4671
		v4978 = v4672
		v4979 = v4673
		v4980 = v4674
		v4981 = v4675
		v5008 = v4702
		v5009 = v4703
		goto L6
	} else {
		goto L467
	}
L465:
	;
	goto L466
L466:
	;
	v4738 = *(*int32)(unsafe.Add(mBase, uint32(v4683)))
	v4739 = *(*int32)(unsafe.Add(mBase, uint32(v4676)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+800)) = v4681
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+796)) = v4680
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+804)) = v4679
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+808)) = v4683
	v4744 = int32(1)
	v4745 = v4691 & v4744
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+815)) = uint8(v4745)
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+816)) = v4694
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+821)) = uint8(v4714)
	v4750 = v4695 & v4744
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+822)) = uint8(v4750)
	v4753 = v4693 & v4744
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+823)) = uint8(v4753)
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+824)) = v4684
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+828)) = v4685
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+832)) = v4686
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+836)) = v4676
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+840)) = v4677
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+844)) = v4682
	v4761 = int32(0)
	F_check_publications_origin_tables(m, v4679, v4739, v4761, v4714, v4694, v4761, v4761, v4738)
	mBase = m.M
	v4765 = m.ExcPending
	if v4765 != 0 {
		v4977 = v4671
		v4978 = v4672
		v4979 = v4673
		v4980 = v4674
		v4981 = v4675
		v5008 = v4702
		v5009 = v4703
		goto L6
	} else {
		goto L468
	}
L467:
	;
	goto L466
L468:
	;
	if v4745 != 0 {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v4766 = *(*int32)(unsafe.Add(mBase, uint32(v4676)+52))
	v4768 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[2]))
	v4769 = *(*int32)(unsafe.Add(mBase, uint32(v4768)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+796)) = v4680
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+800)) = v4681
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+804)) = v4679
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+808)) = v4683
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+815)) = uint8(v4745)
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+816)) = v4694
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+821)) = uint8(v4714)
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+824)) = v4684
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+828)) = v4685
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+832)) = v4686
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+836)) = v4676
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+840)) = v4677
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+844)) = v4682
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+822)) = uint8(v4750)
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+823)) = uint8(v4753)
	if v4753 != 0 {
		goto L472
	} else {
		goto L473
	}
L470:
	;
	goto L471
L471:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[0])) = v4680
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[1])) = v4681
	v4798 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[2]))
	v4799 = *(*int32)(unsafe.Add(mBase, uint32(v4798)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+796)) = v4680
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+800)) = v4681
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+804)) = v4679
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+808)) = v4683
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+815)) = uint8(v4745)
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+816)) = v4694
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+821)) = uint8(v4714)
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+822)) = uint8(v4750)
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+823)) = uint8(v4753)
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+824)) = v4684
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+828)) = v4685
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+832)) = v4686
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+836)) = v4676
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+840)) = v4677
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+844)) = v4682
	m.T0[v4799].(func(*base.Module, int32))(m, v4679)
	mBase = m.M
	v4816 = m.ExcPending
	if v4816 != 0 {
		v4977 = v4671
		v4978 = v4672
		v4979 = v4673
		v4980 = v4674
		v4981 = v4675
		v5008 = v4702
		v5009 = v4703
		goto L6
	} else {
		goto L479
	}
L472:
	;
	v4786 = v4703
	goto L474
L473:
	;
	v4786 = int32(0)
	goto L474
L474:
	;
	if v4750 != 0 {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v4788 = v4702
	goto L477
L476:
	;
	v4788 = int32(0)
	goto L477
L477:
	;
	m.T0[v4769].(func(*base.Module, int32, int32, int32, int32))(m, v4679, v4766, v4786, v4788)
	mBase = m.M
	v4790 = m.ExcPending
	if v4790 != 0 {
		v4977 = v4671
		v4978 = v4672
		v4979 = v4673
		v4980 = v4674
		v4981 = v4675
		v5008 = v4702
		v5009 = v4703
		goto L6
	} else {
		goto L478
	}
L478:
	;
	goto L471
L479:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[0])) = v4680
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[1])) = v4681
	v4821 = v4671
	v4822 = v4672
	v4823 = v4673
	v4824 = v4674
	v4825 = v4675
	v4826 = v4676
	v4827 = v4677
	v4829 = v4679
	v4830 = v4680
	v4831 = v4681
	v4832 = v4682
	v4833 = v4683
	v4834 = v4684
	v4835 = v4685
	v4836 = v4686
	v4841 = v4691
	v4843 = v4693
	v4844 = v4694
	v4845 = v4695
	v4848 = v4698
	v4852 = v4702
	v4853 = v4703
	goto L8
L480:
	;
	v4886 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription[10]))
	if v4886 != 0 {
		goto L481
	} else {
		goto L482
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+800)) = v4831
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+796)) = v4830
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+804)) = v4829
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+808)) = v4833
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+816)) = v4844
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+824)) = v4834
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+828)) = v4835
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+832)) = v4836
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+836)) = v4826
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+840)) = v4827
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+844)) = v4832
	*(*uint8)(unsafe.Add(mBase, uint32(v4825)+815)) = uint8(v4871)
	*(*uint8)(unsafe.Add(mBase, uint32(v4825)+821)) = uint8(v4874)
	*(*uint8)(unsafe.Add(mBase, uint32(v4825)+822)) = uint8(v4877)
	*(*uint8)(unsafe.Add(mBase, uint32(v4825)+823)) = uint8(v4880)
	v4903 = int32(0)
	F_RunObjectPostAlterHook(m, int32(_a_F_AlterSubscription_0), v4827, v4903, v4903, v4903)
	mBase = m.M
	v4907 = m.ExcPending
	if v4907 != 0 {
		v4977 = v4821
		v4978 = v4822
		v4979 = v4823
		v4980 = v4824
		v4981 = v4825
		v5008 = v4852
		v5009 = v4853
		goto L6
	} else {
		goto L484
	}
L482:
	;
	goto L483
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+800)) = v4831
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+796)) = v4830
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+804)) = v4829
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+808)) = v4833
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+816)) = v4844
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+824)) = v4834
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+828)) = v4835
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+832)) = v4836
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+836)) = v4826
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+840)) = v4827
	*(*int32)(unsafe.Add(mBase, uint32(v4825)+844)) = v4832
	*(*uint8)(unsafe.Add(mBase, uint32(v4825)+815)) = uint8(v4871)
	*(*uint8)(unsafe.Add(mBase, uint32(v4825)+821)) = uint8(v4874)
	*(*uint8)(unsafe.Add(mBase, uint32(v4825)+822)) = uint8(v4877)
	*(*uint8)(unsafe.Add(mBase, uint32(v4825)+823)) = uint8(v4880)
	F_LogicalRepWorkersWakeupAtCommit(m, v4827)
	mBase = m.M
	v4924 = m.ExcPending
	if v4924 != 0 {
		v4977 = v4821
		v4978 = v4822
		v4979 = v4823
		v4980 = v4824
		v4981 = v4825
		v5008 = v4852
		v5009 = v4853
		goto L6
	} else {
		goto L485
	}
L484:
	;
	goto L483
L485:
	;
	m.G0 = v4825 + int32(848)
	return
L486:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+800)) = v4681
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+796)) = v4680
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+804)) = v4679
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+808)) = v4683
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+816)) = v4694
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+824)) = v4684
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+828)) = v4685
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+832)) = v4686
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+836)) = v4676
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+840)) = v4677
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+844)) = v4682
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+815)) = uint8(v4940)
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+821)) = uint8(v4944)
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+822)) = uint8(v4947)
	*(*uint8)(unsafe.Add(mBase, uint32(v4675)+823)) = uint8(v4950)
	F_pg_re_throw(m)
	mBase = m.M
	v4976 = m.ExcPending
	if v4976 != 0 {
		v4977 = v4671
		v4978 = v4672
		v4979 = v4673
		v4980 = v4674
		v4981 = v4675
		v5008 = v4702
		v5009 = v4703
		goto L6
	} else {
		goto L487
	}
L487:
	;
	goto L5
L488:
	;
	v5020 = int32(v5016)
	m.G0 = v4981
	v5022 = *(*int32)(unsafe.Add(mBase, uint32(v5020)+4))
	v5023 = *(*int32)(unsafe.Add(mBase, uint32(v5020)))
	v5026 = *(*int32)(unsafe.Add(mBase, uint32(v5023)))
	if v4981+int32(236) == v5026 {
		goto L491
	} else {
		goto L492
	}
L489:
	;
	m.ExcPending = 1
	goto L497
L490:
	;
	if v5030 != 0 {
		goto L494
	} else {
		goto L495
	}
L491:
	;
	v5028 = *(*int32)(unsafe.Add(mBase, uint32(v5023)+4))
	v5030 = v5028
	goto L493
L492:
	;
	v5030 = int32(0)
	goto L493
L493:
	;
	goto L490
L494:
	;
	v5031 = *(*int32)(unsafe.Add(mBase, uint32(v4981)+844))
	v5032 = *(*int32)(unsafe.Add(mBase, uint32(v4981)+840))
	v5033 = *(*int32)(unsafe.Add(mBase, uint32(v4981)+836))
	v5034 = *(*int32)(unsafe.Add(mBase, uint32(v4981)+832))
	v5035 = *(*int32)(unsafe.Add(mBase, uint32(v4981)+828))
	v5036 = *(*int32)(unsafe.Add(mBase, uint32(v4981)+824))
	v5037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4981)+823)))
	v5038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4981)+822)))
	v5039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4981)+821)))
	v5040 = *(*int32)(unsafe.Add(mBase, uint32(v4981)+816))
	v5041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4981)+815)))
	v5042 = *(*int32)(unsafe.Add(mBase, uint32(v4981)+808))
	v5043 = *(*int32)(unsafe.Add(mBase, uint32(v4981)+804))
	v5044 = *(*int32)(unsafe.Add(mBase, uint32(v4981)+800))
	v5045 = *(*int32)(unsafe.Add(mBase, uint32(v4981)+796))
	v48 = v4977
	v49 = v4978
	v50 = v4979
	v51 = v4980
	v52 = v4981
	v53 = v5033
	v54 = v5032
	v55 = v5040
	v56 = v5043
	v57 = v5045
	v58 = v5044
	v59 = v5031
	v60 = v5042
	v61 = v5036
	v62 = v5035
	v63 = v5034
	v64 = v5030
	v68 = v5041
	v69 = v5022
	v70 = v5037
	v72 = v5038
	v75 = v5039
	v79 = v5008
	v80 = v5009
	goto L1
L495:
	;
	goto L496
L496:
	;
	F___wasm_longjmp(m, v5023, v5022)
	mBase = m.M
	v5047 = m.ExcPending
	if v5047 != 0 {
		goto L497
	} else {
		goto L498
	}
L497:
	;
	return
L498:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
