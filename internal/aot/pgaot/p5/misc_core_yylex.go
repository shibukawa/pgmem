package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_core_yylex(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
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
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
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
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v678 int64
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v704 int64
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v849 int64
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v894 int64
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1228 int32
	_ = v1228
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1446 int32
	_ = v1446
	var v1457 int32
	_ = v1457
	var v1465 int32
	_ = v1465
	var v1480 int32
	_ = v1480
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1514 int32
	_ = v1514
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1632 int32
	_ = v1632
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1666 int32
	_ = v1666
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1687 int32
	_ = v1687
	var v1690 int64
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1730 int64
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
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
	var v1747 int32
	_ = v1747
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1756 int64
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1779 int32
	_ = v1779
	var v1782 int64
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1808 int64
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1884 int32
	_ = v1884
	var v1887 int64
	_ = v1887
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1979 int32
	_ = v1979
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1989 int32
	_ = v1989
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2006 int32
	_ = v2006
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2017 int32
	_ = v2017
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2068 int32
	_ = v2068
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2081 int32
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2087 int32
	_ = v2087
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2100 int32
	_ = v2100
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2113 int32
	_ = v2113
	var v2115 int32
	_ = v2115
	var v2119 int32
	_ = v2119
	var v2122 int32
	_ = v2122
	var v2133 int32
	_ = v2133
	var v2136 int32
	_ = v2136
	var v2139 int32
	_ = v2139
	var v2146 int32
	_ = v2146
	var v2148 int32
	_ = v2148
	var v2153 int32
	_ = v2153
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2173 int32
	_ = v2173
	var v2177 int32
	_ = v2177
	var v2180 int32
	_ = v2180
	var v2183 int32
	_ = v2183
	var v2187 int32
	_ = v2187
	var v2190 int32
	_ = v2190
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2212 int32
	_ = v2212
	var v2216 int32
	_ = v2216
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2227 int32
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2254 int32
	_ = v2254
	var v2257 int32
	_ = v2257
	var v2267 int32
	_ = v2267
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2279 int32
	_ = v2279
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2301 int32
	_ = v2301
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2324 int32
	_ = v2324
	var v2328 int32
	_ = v2328
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2352 int32
	_ = v2352
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2369 int32
	_ = v2369
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2416 int32
	_ = v2416
	var v2432 int32
	_ = v2432
	var v2439 int32
	_ = v2439
	var v2442 int32
	_ = v2442
	var v2446 int32
	_ = v2446
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2454 int32
	_ = v2454
	var v2459 int32
	_ = v2459
	var v2462 int32
	_ = v2462
	var v2465 int32
	_ = v2465
	var v2468 int32
	_ = v2468
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2482 int32
	_ = v2482
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2500 int32
	_ = v2500
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2519 int32
	_ = v2519
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2531 int32
	_ = v2531
	var v2533 int32
	_ = v2533
	var v2535 int32
	_ = v2535
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2542 int32
	_ = v2542
	var v2546 int32
	_ = v2546
	var v2548 int32
	_ = v2548
	var v2551 int32
	_ = v2551
	var v2558 int32
	_ = v2558
	var v2560 int32
	_ = v2560
	var v2563 int32
	_ = v2563
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2569 int32
	_ = v2569
	var v2571 int32
	_ = v2571
	var v2572 int32
	_ = v2572
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2583 int32
	_ = v2583
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2592 int32
	_ = v2592
	var v2600 int32
	_ = v2600
	var v2601 int32
	_ = v2601
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2606 int32
	_ = v2606
	var v2608 int32
	_ = v2608
	var v2609 int32
	_ = v2609
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2615 int32
	_ = v2615
	var v2618 int32
	_ = v2618
	var v2619 int32
	_ = v2619
	var v2620 int32
	_ = v2620
	var v2623 int32
	_ = v2623
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2645 int32
	_ = v2645
	var v2646 int32
	_ = v2646
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2657 int32
	_ = v2657
	var v2659 int32
	_ = v2659
	var v2660 int32
	_ = v2660
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2683 int32
	_ = v2683
	var v2684 int32
	_ = v2684
	var v2686 int32
	_ = v2686
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2695 int32
	_ = v2695
	var v2699 int32
	_ = v2699
	var v2700 int32
	_ = v2700
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2716 int32
	_ = v2716
	var v2721 int32
	_ = v2721
	var v2727 int32
	_ = v2727
	var v2729 int32
	_ = v2729
	var v2732 int32
	_ = v2732
	var v2733 int32
	_ = v2733
	var v2735 int32
	_ = v2735
	var v2746 int32
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2749 int32
	_ = v2749
	var v2750 int32
	_ = v2750
	var v2753 int32
	_ = v2753
	var v2756 int32
	_ = v2756
	var v2758 int32
	_ = v2758
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2768 int32
	_ = v2768
	var v2770 int32
	_ = v2770
	var v2776 int32
	_ = v2776
	var v2778 int32
	_ = v2778
	var v2783 int32
	_ = v2783
	var v2785 int32
	_ = v2785
	var v2788 int32
	_ = v2788
	var v2790 int32
	_ = v2790
	var v2793 int32
	_ = v2793
	var v2796 int32
	_ = v2796
	var v2799 int32
	_ = v2799
	var v2807 int32
	_ = v2807
	var v2808 int32
	_ = v2808
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2814 int32
	_ = v2814
	var v2817 int32
	_ = v2817
	var v2819 int32
	_ = v2819
	var v2823 int32
	_ = v2823
	var v2824 int32
	_ = v2824
	var v2827 int32
	_ = v2827
	var v2839 int32
	_ = v2839
	var v2842 int32
	_ = v2842
	var v2852 int32
	_ = v2852
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2858 int32
	_ = v2858
	var v2861 int32
	_ = v2861
	var v2862 int32
	_ = v2862
	var v2864 int32
	_ = v2864
	var v2865 int32
	_ = v2865
	var v2868 int32
	_ = v2868
	var v2871 int32
	_ = v2871
	var v2872 int32
	_ = v2872
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2878 int32
	_ = v2878
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2884 int32
	_ = v2884
	var v2885 int32
	_ = v2885
	var v2888 int32
	_ = v2888
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2896 int32
	_ = v2896
	var v2898 int32
	_ = v2898
	var v2911 int32
	_ = v2911
	var v2913 int32
	_ = v2913
	var v2923 int32
	_ = v2923
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2929 int32
	_ = v2929
	var v2932 int32
	_ = v2932
	var v2933 int32
	_ = v2933
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2939 int32
	_ = v2939
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2945 int32
	_ = v2945
	var v2946 int32
	_ = v2946
	var v2949 int32
	_ = v2949
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2955 int32
	_ = v2955
	var v2956 int32
	_ = v2956
	var v2959 int32
	_ = v2959
	var v2962 int32
	_ = v2962
	var v2964 int32
	_ = v2964
	var v2969 int32
	_ = v2969
	var v2983 int32
	_ = v2983
	var v2989 int32
	_ = v2989
	var v3000 int32
	_ = v3000
	var v3004 int32
	_ = v3004
	var v3008 int32
	_ = v3008
	var v3010 int32
	_ = v3010
	var v3023 int32
	_ = v3023
	var v3025 int32
	_ = v3025
	var v3035 int32
	_ = v3035
	var v3037 int32
	_ = v3037
	var v3038 int32
	_ = v3038
	var v3041 int32
	_ = v3041
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3051 int32
	_ = v3051
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3057 int32
	_ = v3057
	var v3058 int32
	_ = v3058
	var v3061 int32
	_ = v3061
	var v3064 int32
	_ = v3064
	var v3065 int32
	_ = v3065
	var v3067 int32
	_ = v3067
	var v3068 int32
	_ = v3068
	var v3071 int32
	_ = v3071
	var v3074 int32
	_ = v3074
	var v3076 int32
	_ = v3076
	var v3081 int32
	_ = v3081
	var v3093 int32
	_ = v3093
	var v3096 int32
	_ = v3096
	var v3098 int32
	_ = v3098
	var v3104 int32
	_ = v3104
	var v3106 int32
	_ = v3106
	var v3107 int32
	_ = v3107
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3118 int32
	_ = v3118
	var v3119 int32
	_ = v3119
	var v3125 int32
	_ = v3125
	var v3126 int32
	_ = v3126
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3129 int32
	_ = v3129
	var v3134 int32
	_ = v3134
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(l2)+96)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l2)+92)) = l0
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = int32(1)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if v25 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v91 = l2
	v97 = v16
	goto L22
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+44)) = int32(1)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v30 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v34
	goto L9
L8:
	;
	goto L9
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v36 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v40
	goto L12
L11:
	;
	goto L12
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v72
	v76 = v70 + v69<<(uint(int32(2))%32)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v82
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+24)) = uint8(v84)
	goto L3
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42+v43<<(uint(int32(2))%32))))
	if v47 != 0 {
		v69 = v43
		v70 = v42
		v71 = v47
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_core_yyensure_buffer_stack(m, l2)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v55 = F_core_yy_create_buffer(m, v54, l2)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v59 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v57+v58<<(uint(v59)%32)))) = v55
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63+v64<<(uint(v59)%32))))
	v69 = v64
	v70 = v63
	v71 = v68
	goto L13
L21:
	;
	v3126 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v3127 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v3128 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v3129 = *(*int32)(unsafe.Add(mBase, uint32(v3128)))
	*(*int32)(unsafe.Add(mBase, uint32(v3126))) = v3127 - v3129
	F_scanner_yyerror(m, int32(_a_F_core_yylex_0), v91)
	mBase = m.M
	v3134 = m.ExcPending
	if v3134 != 0 {
		goto L18
	} else {
		goto L654
	}
L22:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v91)+36))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v103)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v91)+44))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v105<<(uint(int32(2))%32))+uint32(_c_F_core_yylex[2])))
	v111 = v103
	v112 = v102
	v114 = v110
	v115 = v102
	goto L25
L23:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_1), v91)
	mBase = m.M
	v3125 = m.ExcPending
	if v3125 != 0 {
		goto L18
	} else {
		goto L653
	}
L24:
	;
	goto L23
L25:
	;
	v125 = v111 & int32(255)
	v128 = v114 + v125<<(uint(int32(3))%32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v129 == v125 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v3117 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v3118 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v3119 = *(*int32)(unsafe.Add(mBase, uint32(v3118)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v3117))) = v3119
	goto L22
L27:
	;
	v132 = v112
	v134 = v114
	v137 = v128
	goto L30
L28:
	;
	v157 = v112
	v159 = v114
	goto L29
L29:
	;
	v170 = v157
	v172 = v159
	v173 = v115
	goto L35
L30:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	v146 = v132 + int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v148 = int32(3)
	v150 = v134 + v147<<(uint(v148)%32)
	v153 = v150 + v144<<(uint(v148)%32)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	if v144 == v154 {
		v132 = v146
		v134 = v150
		v137 = v153
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v157 = v146
	v159 = v150
	goto L29
L32:
	;
	goto L31
L33:
	;
	goto L26
L34:
	;
	v3116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3104))))
	v111 = v3116
	v112 = v3104
	v114 = v3106
	v115 = v3107
	goto L25
L35:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v172-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v170 - v173
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v188)
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v190)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v170
	v193 = v184
	goto L38
L36:
	;
	v3096 = *(*int32)(unsafe.Add(mBase, uint32(v3081)+2052))
	v3098 = v2010 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v3098
	v3104 = v3098
	v3106 = v3081 + v3096<<(uint(int32(3))%32)
	v3107 = v2006
	goto L34
L37:
	;
	if v3081 == int32(0) {
		v170 = v2010
		v172 = v3081
		v173 = v2006
		goto L35
	} else {
		goto L651
	}
L38:
	;
	switch v193 - int32(1) {
	case 0, 4, 5, 6:
		goto L22
	case 1:
		goto L133
	case 2:
		goto L132
	case 3:
		goto L131
	case 7:
		goto L129
	case 8, 9:
		goto L128
	case 10:
		goto L126
	case 11:
		goto L124
	case 12:
		goto L123
	case 13:
		goto L122
	case 14:
		goto L121
	case 15:
		goto L120
	case 16:
		goto L119
	case 17, 18, 80:
		goto L118
	case 19:
		goto L117
	case 20:
		goto L116
	case 21:
		goto L115
	case 22:
		goto L114
	case 23:
		goto L113
	case 24, 25, 85:
		goto L112
	case 26:
		goto L111
	case 27:
		goto L110
	case 28:
		goto L109
	case 29:
		goto L108
	case 30:
		goto L107
	case 31:
		goto L105
	case 32:
		goto L104
	case 33:
		goto L103
	case 34:
		goto L102
	case 35:
		goto L101
	case 36:
		goto L100
	case 37:
		goto L98
	case 38:
		goto L97
	case 39:
		goto L96
	case 40:
		goto L95
	case 41:
		goto L94
	case 42:
		goto L93
	case 43:
		goto L91
	case 44:
		goto L90
	case 45:
		goto L89
	case 46:
		goto L88
	case 47:
		goto L87
	case 48:
		goto L86
	case 49:
		goto L85
	case 50:
		goto L63
	case 51:
		goto L84
	case 52:
		goto L83
	case 53:
		goto L82
	case 54:
		goto L81
	case 55:
		goto L80
	case 56:
		goto L79
	case 57:
		goto L78
	case 58:
		goto L77
	case 59:
		goto L76
	case 60:
		goto L75
	case 61:
		goto L74
	case 62:
		goto L73
	case 63:
		goto L72
	case 64:
		goto L71
	case 65:
		goto L70
	case 66, 67, 68, 69:
		goto L21
	case 70:
		goto L69
	case 71:
		goto L68
	case 72:
		goto L66
	case 73:
		goto L65
	case 74:
		goto L67
	case 75:
		goto L127
	case 76:
		goto L130
	case 77, 83:
		goto L92
	case 78:
		goto L125
	case 79, 81, 84:
		goto L106
	case 82:
		goto L99
	default:
		goto L64
	}
L39:
	;
	if base.Ui32(v170-v1967-int32(2)) < base.Ui32(int32(3)) {
		v3081 = v3010
		goto L37
	} else {
		goto L635
	}
L40:
	;
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v2989
	*(*int32)(unsafe.Add(mBase, uint32(v91)+48)) = int32(0)
	v3000 = *(*int32)(unsafe.Add(mBase, uint32(v91)+44))
	v3004 = base.I32_div_s(v3000-int32(1), int32(2))
	v193 = v3004 + int32(75)
	goto L38
L42:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_2))
	mBase = m.M
	v2983 = m.ExcPending
	if v2983 != 0 {
		goto L18
	} else {
		goto L634
	}
L43:
	;
	v3104 = v2714
	v3106 = v2969
	v3107 = v2705
	goto L34
L44:
	;
	if base.Ui32(v170-v1967-int32(2)) < base.Ui32(int32(3)) {
		v2969 = v2898
		goto L43
	} else {
		goto L618
	}
L45:
	;
	if base.Ui32(v2785-int32(1)) < base.Ui32(int32(3)) {
		v170 = v2776
		v172 = v2827
		v173 = v2768
		goto L35
	} else {
		goto L602
	}
L46:
	;
	v2824 = v2768
	v2827 = v2783
	goto L45
L47:
	;
	v2896 = v2705
	v2898 = v2721
	goto L44
L48:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_3))
	mBase = m.M
	v2823 = m.ExcPending
	if v2823 != 0 {
		goto L18
	} else {
		goto L601
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v2776 = v2763 + v2770
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v2776
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(v91)+44))
	v2783 = *(*int32)(unsafe.Add(mBase, uint32(v2778<<(uint(int32(2))%32))+uint32(_c_F_core_yylex[2])))
	if base.Ui32(v2776) <= base.Ui32(v2768) {
		v170 = v2776
		v172 = v2783
		v173 = v2768
		goto L35
	} else {
		goto L593
	}
L51:
	;
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(v2506)))
	*(*int32)(unsafe.Add(mBase, uint32(v2507)+16)) = v2500
	v2510 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	if v2510 != 0 {
		v2632 = int32(0)
		goto L548
	} else {
		goto L549
	}
L52:
	;
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2489 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2500 = v2482
	v2506 = v2488 + v2489<<(uint(int32(2))%32)
	goto L51
L53:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_4))
	mBase = m.M
	v2474 = m.ExcPending
	if v2474 != 0 {
		goto L18
	} else {
		goto L547
	}
L54:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_5))
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L18
	} else {
		goto L546
	}
L55:
	;
	v3008 = v2006
	v3010 = v2017
	goto L40
L56:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_6), v91)
	mBase = m.M
	v2468 = m.ExcPending
	if v2468 != 0 {
		goto L18
	} else {
		goto L545
	}
L57:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_7), v91)
	mBase = m.M
	v2465 = m.ExcPending
	if v2465 != 0 {
		goto L18
	} else {
		goto L544
	}
L58:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_7), v91)
	mBase = m.M
	v2462 = m.ExcPending
	if v2462 != 0 {
		goto L18
	} else {
		goto L543
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2439 = m.ExcPending
	if v2439 != 0 {
		goto L18
	} else {
		goto L537
	}
L60:
	;
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v2432)+48)) = v679
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(23)
	goto L33
L61:
	;
	m.G0 = v97 + int32(16)
	return v2416
L62:
	;
	v2416 = int32(274)
	goto L61
L63:
	;
	v2395 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v2398 = *(*int32)(unsafe.Add(mBase, uint32(v2397)))
	*(*int32)(unsafe.Add(mBase, uint32(v2395))) = v2396 - v2398
	goto L62
L64:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_8))
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L18
	} else {
		goto L536
	}
L65:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v1968)
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v1974 = v1970 + v1971<<(uint(int32(2))%32)
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(v1974)))
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v1975)+44))
	if v1976 == int32(0) {
		goto L451
	} else {
		goto L452
	}
L66:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_9))
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		goto L18
	} else {
		goto L450
	}
L67:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(v1959)))
	*(*int32)(unsafe.Add(mBase, uint32(v1957))) = v1958 - v1960
	v2416 = int32(0)
	goto L61
L68:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1951)))
	*(*int32)(unsafe.Add(mBase, uint32(v1949))) = v1950 - v1952
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1956 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1955))))
	v2416 = v1956
	goto L61
L69:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v1915)))
	*(*int32)(unsafe.Add(mBase, uint32(v1913))) = v1914 - v1916
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v1920)+8))
	v1922 = F_ScanKeywordLookup(m, v1919, v1921)
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L18
	} else {
		goto L445
	}
L70:
	;
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1903)))
	*(*int32)(unsafe.Add(mBase, uint32(v1901))) = v1902 - v1904
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1908 = F_pstrdup(m, v1907)
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L18
	} else {
		goto L444
	}
L71:
	;
	v1861 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v1862)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v1866 = v1861 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v1866
	v1868 = v1866 + v173
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1868
	v1870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1868))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v1870)
	v1872 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1868))) = uint8(v1872)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1868
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1877)))
	*(*int32)(unsafe.Add(mBase, uint32(v1875))) = v1876 - v1878
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1884 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v1884
	v1887 = *(*int64)(unsafe.Add(mBase, _c_F_core_yylex[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v1887
	v1889 = F_pg_strtoint32_safe(m, v1882, v97)
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L18
	} else {
		goto L439
	}
L72:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1851)))
	*(*int32)(unsafe.Add(mBase, uint32(v1849))) = v1850 - v1852
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1856 = F_pstrdup(m, v1855)
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L18
	} else {
		goto L438
	}
L73:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1842)))
	*(*int32)(unsafe.Add(mBase, uint32(v1840))) = v1841 - v1843
	F_scanner_yyerror(m, int32(_a_F_core_yylex_10), v91)
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L18
	} else {
		goto L437
	}
L74:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1833)))
	*(*int32)(unsafe.Add(mBase, uint32(v1831))) = v1832 - v1834
	F_scanner_yyerror(m, int32(_a_F_core_yylex_11), v91)
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L18
	} else {
		goto L436
	}
L75:
	;
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1824)))
	*(*int32)(unsafe.Add(mBase, uint32(v1822))) = v1823 - v1825
	F_scanner_yyerror(m, int32(_a_F_core_yylex_12), v91)
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L18
	} else {
		goto L435
	}
L76:
	;
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(v1798)))
	*(*int32)(unsafe.Add(mBase, uint32(v1796))) = v1797 - v1799
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1805 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v1805
	v1808 = *(*int64)(unsafe.Add(mBase, _c_F_core_yylex[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v1808
	v1810 = F_pg_strtoint32_safe(m, v1803, v97)
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L18
	} else {
		goto L430
	}
L77:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1772)))
	*(*int32)(unsafe.Add(mBase, uint32(v1770))) = v1771 - v1773
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1779 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v1779
	v1782 = *(*int64)(unsafe.Add(mBase, _c_F_core_yylex[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v1782
	v1784 = F_pg_strtoint32_safe(m, v1777, v97)
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L18
	} else {
		goto L425
	}
L78:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v1746)))
	*(*int32)(unsafe.Add(mBase, uint32(v1744))) = v1745 - v1747
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1753 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v1753
	v1756 = *(*int64)(unsafe.Add(mBase, _c_F_core_yylex[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v1756
	v1758 = F_pg_strtoint32_safe(m, v1751, v97)
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L18
	} else {
		goto L420
	}
L79:
	;
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1720)))
	*(*int32)(unsafe.Add(mBase, uint32(v1718))) = v1719 - v1721
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1727 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v1727
	v1730 = *(*int64)(unsafe.Add(mBase, _c_F_core_yylex[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v1730
	v1732 = F_pg_strtoint32_safe(m, v1725, v97)
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L18
	} else {
		goto L415
	}
L80:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1711)))
	*(*int32)(unsafe.Add(mBase, uint32(v1709))) = v1710 - v1712
	F_scanner_yyerror(m, int32(_a_F_core_yylex_13), v91)
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L18
	} else {
		goto L414
	}
L81:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v1687
	v1690 = *(*int64)(unsafe.Add(mBase, _c_F_core_yylex[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v1690
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1694)))
	*(*int32)(unsafe.Add(mBase, uint32(v1692))) = v1693 - v1695
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1701 = F_pg_strtoint32_safe(m, v1698+int32(1), v97)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L18
	} else {
		goto L412
	}
L82:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1429 = F_strstr(m, v1427, int32(_a_F_core_yylex_14))
	mBase = m.M
	v1431 = F_strstr(m, v1427, int32(_a_F_core_yylex_15))
	mBase = m.M
	if base.B2i32(v1431 != int32(0))&base.B2i32(base.Ui32(v1431) < base.Ui32(v1429)) != 0 {
		goto L347
	} else {
		goto L348
	}
L83:
	;
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v1421)))
	*(*int32)(unsafe.Add(mBase, uint32(v1419))) = v1420 - v1422
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1426 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1425))))
	v2416 = v1426
	goto L61
L84:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1415)))
	*(*int32)(unsafe.Add(mBase, uint32(v1413))) = v1414 - v1416
	goto L62
L85:
	;
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1408)))
	*(*int32)(unsafe.Add(mBase, uint32(v1406))) = v1407 - v1409
	v2416 = int32(273)
	goto L61
L86:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1401)))
	*(*int32)(unsafe.Add(mBase, uint32(v1399))) = v1400 - v1402
	v2416 = int32(272)
	goto L61
L87:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1394)))
	*(*int32)(unsafe.Add(mBase, uint32(v1392))) = v1393 - v1395
	v2416 = int32(271)
	goto L61
L88:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	*(*int32)(unsafe.Add(mBase, uint32(v1385))) = v1386 - v1388
	v2416 = int32(270)
	goto L61
L89:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1380)))
	*(*int32)(unsafe.Add(mBase, uint32(v1378))) = v1379 - v1381
	v2416 = int32(269)
	goto L61
L90:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1373)))
	*(*int32)(unsafe.Add(mBase, uint32(v1371))) = v1372 - v1374
	v2416 = int32(268)
	goto L61
L91:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1346)))
	*(*int32)(unsafe.Add(mBase, uint32(v1344))) = v1345 - v1347
	v1350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v1350)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v1353 = int32(1)
	v1354 = v173 + v1353
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v1353
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v1358)
	v1360 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)) = uint8(v1360)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1354
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1366 = F_downcase_truncate_identifier(m, v1363, v1364, v1353)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L18
	} else {
		goto L345
	}
L92:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_16), v91)
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L18
	} else {
		goto L344
	}
L93:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1307)+24))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1310 = v1308 + v1309
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1307)+28))
	if v1311 <= v1310 {
		goto L334
	} else {
		goto L335
	}
L94:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1276)+24))
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1276)+28))
	if v1280 <= v1277+int32(1) {
		goto L330
	} else {
		goto L331
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(1)
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1259)+24))
	if v1260 == int32(0) {
		goto L57
	} else {
		goto L325
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(1)
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+24))
	if v1234 == int32(0) {
		goto L58
	} else {
		goto L316
	}
L97:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1222)))
	*(*int32)(unsafe.Add(mBase, uint32(v1220))) = v1221 - v1223
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(19)
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1228)+24)) = int32(0)
	goto L22
L98:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1211)))
	*(*int32)(unsafe.Add(mBase, uint32(v1209))) = v1210 - v1212
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(7)
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+24)) = int32(0)
	goto L22
L99:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_17), v91)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L18
	} else {
		goto L315
	}
L100:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1176))))
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1178)+24))
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1178)+28))
	if v1182 <= v1179+int32(1) {
		goto L311
	} else {
		goto L312
	}
L101:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+24))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1145 = v1143 + v1144
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+28))
	if v1146 <= v1145 {
		goto L301
	} else {
		goto L302
	}
L102:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+24))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1110 = v1108 + v1109
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+28))
	if v1111 <= v1110 {
		goto L291
	} else {
		goto L292
	}
L103:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+40))
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1003))))
	v1011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1005))))
	if base.B2i32(v1008 == int32(0))|base.B2i32(v1008 != v1011) != 0 {
		v1029 = v1008
		v1030 = v1011
		goto L267
	} else {
		goto L268
	}
L104:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v984)))
	*(*int32)(unsafe.Add(mBase, uint32(v982))) = v983 - v985
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v988)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v991 = int32(1)
	v992 = v173 + v991
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v992
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v991
	v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v996)
	v998 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)) = uint8(v998)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v992
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1002 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1001))))
	v2416 = v1002
	goto L61
L105:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v968)))
	*(*int32)(unsafe.Add(mBase, uint32(v966))) = v967 - v969
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v973 = F_pstrdup(m, v972)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L18
	} else {
		goto L265
	}
L106:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_18), v91)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L18
	} else {
		goto L264
	}
L107:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933))))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v935)+24))
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v935)+28))
	if v939 <= v936+int32(1) {
		goto L260
	} else {
		goto L261
	}
L108:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v894 = F_strtox_2(m, v888+int32(2), int32(0), int32(16), int64(4294967295))
	mBase = m.M
	v895 = base.I32_wrap_i64(v894)
	goto L251
L109:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v849 = F_strtox_2(m, v843+int32(1), int32(0), int32(8), int64(4294967295))
	mBase = m.M
	v850 = base.I32_wrap_i64(v849)
	goto L242
L110:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+1)))
	if v776 != int32(39) {
		goto L226
	} else {
		goto L227
	}
L111:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	*(*int32)(unsafe.Add(mBase, uint32(v745))) = v746 - v748
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L18
	} else {
		goto L220
	}
L112:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v741)))
	*(*int32)(unsafe.Add(mBase, uint32(v739))) = v740 - v742
	goto L24
L113:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v704 = F_strtox_2(m, v698+int32(2), int32(0), int32(16), int64(4294967295))
	mBase = m.M
	v705 = base.I32_wrap_i64(v704)
	goto L217
L114:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v678 = F_strtox_2(m, v672+int32(2), int32(0), int32(16), int64(4294967295))
	mBase = m.M
	v679 = base.I32_wrap_i64(v678)
	goto L213
L115:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)+24))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v641 = v639 + v640
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v638)+28))
	if v642 <= v641 {
		goto L203
	} else {
		goto L204
	}
L116:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)+24))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v606 = v604 + v605
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	if v607 <= v606 {
		goto L193
	} else {
		goto L194
	}
L117:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+24))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	if v576 <= v573+int32(1) {
		goto L189
	} else {
		goto L190
	}
L118:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v487)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v173
	v491 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v491
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v493)
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v491)
	v497 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v173
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)+32))
	switch v501 - v497 {
	case 0:
		goto L167
	default:
		goto L163
	case 3:
		goto L166
	case 4, 6:
		goto L165
	case 9:
		goto L164
	}
L119:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)+32))
	v482 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = v481<<(uint(v482)%32) | v482
	goto L22
L120:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v91)+44))
	v476 = base.I32_div_s(v472-int32(1), int32(2))
	*(*int32)(unsafe.Add(mBase, uint32(v471)+32)) = v476
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(13)
	goto L22
L121:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	*(*int32)(unsafe.Add(mBase, uint32(v460))) = v461 - v463
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(21)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+24)) = int32(0)
	goto L22
L122:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v447 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v446)+52)) = uint8(v447)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	*(*int32)(unsafe.Add(mBase, uint32(v449))) = v450 - v452
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(15)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v457)+24)) = v447
	goto L22
L123:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v433 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+52)) = uint8(v433)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	*(*int32)(unsafe.Add(mBase, uint32(v435))) = v436 - v438
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(11)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v443)+24)) = v433
	goto L22
L124:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	*(*int32)(unsafe.Add(mBase, uint32(v385))) = v386 - v388
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v391)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v394 = int32(1)
	v395 = v173 + v394
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v394
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v399)
	v401 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)) = uint8(v401)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v395
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+8))
	v407 = F_ScanKeywordLookup(m, int32(_a_F_core_yylex_19), v406)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L18
	} else {
		goto L158
	}
L125:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_20), v91)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L18
	} else {
		goto L157
	}
L126:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = v342 - v344
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(9)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+24)) = int32(0)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+24))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v352)+28))
	if v356 <= v353+int32(1) {
		goto L153
	} else {
		goto L154
	}
L127:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_21), v91)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L18
	} else {
		goto L152
	}
L128:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+24))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v307 = v305 + v306
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v304)+28))
	if v308 <= v307 {
		goto L142
	} else {
		goto L143
	}
L129:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v263 - v265
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(3)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v270)+24)) = int32(0)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+24))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v273)+28))
	if v277 <= v274+int32(1) {
		goto L138
	} else {
		goto L139
	}
L130:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_22), v91)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L18
	} else {
		goto L137
	}
L131:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+36))
	if v251 <= int32(0) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v232)+36)) = v233 + int32(1)
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v237)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v240 = int32(2)
	v241 = v173 + v240
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v240
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v245)
	v247 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)) = uint8(v247)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v241
	goto L22
L133:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v209 - v211
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v215 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v214)+36)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(5)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v219)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v222 = int32(2)
	v223 = v173 + v222
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v222
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v227)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)) = uint8(v215)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v223
	goto L22
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(1)
	goto L22
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+36)) = v251 - int32(1)
	goto L22
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+28)) = v277 << (uint(int32(1)) % 32)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+20))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v282)+28))
	v285 = F_repalloc(m, v283, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L18
	} else {
		goto L141
	}
L139:
	;
	v292 = v273
	v293 = v274
	goto L140
L140:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v292)+20))
	v296 = int32(98)
	*(*uint8)(unsafe.Add(mBase, uint32(v293+v294))) = uint8(v296)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+24)) = v299 + int32(1)
	goto L22
L141:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v287)+20)) = v285
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+24))
	v292 = v289
	v293 = v290
	goto L140
L142:
	;
	v310 = int32(1)
	v313 = v307 + v310
	if v313&v307 != 0 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v329 = v304
	v330 = v305
	goto L144
L144:
	;
	if v306 != 0 {
		goto L149
	} else {
		goto L150
	}
L145:
	;
	v318 = v310 << (uint(int32(32)-base.I32_clz(v313)) % 32)
	goto L147
L146:
	;
	v318 = v313
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v304)+28)) = v318
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+20))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v320)+28))
	v323 = F_repalloc(m, v321, v322)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L18
	} else {
		goto L148
	}
L148:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v325)+20)) = v323
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+24))
	v329 = v327
	v330 = v328
	goto L144
L149:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v329)+20))
	base.MemoryCopy(m, v331+v330, v303, v306)
	goto L151
L150:
	;
	goto L151
L151:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v334)+24)) = v335 + v306
	goto L22
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v352)+28)) = v356 << (uint(int32(1)) % 32)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v361)+28))
	v364 = F_repalloc(m, v362, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L18
	} else {
		goto L156
	}
L154:
	;
	v371 = v352
	v372 = v353
	goto L155
L155:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v371)+20))
	v375 = int32(120)
	*(*uint8)(unsafe.Add(mBase, uint32(v372+v373))) = uint8(v375)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v377)+24)) = v378 + int32(1)
	goto L22
L156:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+20)) = v364
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+24))
	v371 = v368
	v372 = v369
	goto L155
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	if int32(0) <= v407 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)+8))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	v416 = v407 << (uint(int32(1)) % 32)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v413)+4))
	v419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v416+v417))))
	*(*int32)(unsafe.Add(mBase, uint32(v411))) = v414 + v419
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)+12))
	v425 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v423+v416))))
	v2416 = v425
	goto L61
L160:
	;
	goto L161
L161:
	;
	v427 = F_pstrdup(m, int32(_a_F_core_yylex_23))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L18
	} else {
		goto L162
	}
L162:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = v427
	v2416 = int32(258)
	goto L61
L163:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_24), v91)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L18
	} else {
		goto L188
	}
L164:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v500)+24))
	v558 = F_palloc(m, v555+int32(1))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L18
	} else {
		goto L184
	}
L165:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+52)))
	if v532 == int32(1) {
		goto L176
	} else {
		goto L177
	}
L166:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v500)+24))
	v521 = F_palloc(m, v518+int32(1))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L18
	} else {
		goto L172
	}
L167:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v500)+24))
	v507 = F_palloc(m, v504+int32(1))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L18
	} else {
		goto L168
	}
L168:
	;
	if v504 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+20))
	base.MemoryCopy(m, v507, v510, v504)
	goto L171
L170:
	;
	goto L171
L171:
	;
	v513 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v504+v507))) = uint8(v513)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v515))) = v507
	v2416 = int32(263)
	goto L61
L172:
	;
	if v518 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+20))
	base.MemoryCopy(m, v521, v524, v518)
	goto L175
L174:
	;
	goto L175
L175:
	;
	v527 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v518+v521))) = uint8(v527)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v529))) = v521
	v2416 = int32(264)
	goto L61
L176:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v500)+20))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v500)+24))
	F_pg_verifymbstr(m, v535, v536)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L18
	} else {
		goto L179
	}
L177:
	;
	v540 = v500
	goto L178
L178:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)+24))
	v544 = F_palloc(m, v541+int32(1))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L18
	} else {
		goto L180
	}
L179:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v540 = v539
	goto L178
L180:
	;
	if v541 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)+20))
	base.MemoryCopy(m, v544, v547, v541)
	goto L183
L182:
	;
	goto L183
L183:
	;
	v550 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v541+v544))) = uint8(v550)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v552))) = v544
	v2416 = int32(261)
	goto L61
L184:
	;
	if v555 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)+20))
	base.MemoryCopy(m, v558, v561, v555)
	goto L187
L186:
	;
	goto L187
L187:
	;
	v564 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v555+v558))) = uint8(v564)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v566))) = v558
	v2416 = int32(262)
	goto L61
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v576 << (uint(int32(1)) % 32)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+20))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v581)+28))
	v584 = F_repalloc(m, v582, v583)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L18
	} else {
		goto L192
	}
L190:
	;
	v591 = v572
	v592 = v573
	goto L191
L191:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v591)+20))
	v595 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v592+v593))) = uint8(v595)
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v597)+24)) = v598 + int32(1)
	goto L22
L192:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v586)+20)) = v584
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v588)+24))
	v591 = v588
	v592 = v589
	goto L191
L193:
	;
	v609 = int32(1)
	v612 = v606 + v609
	if v612&v606 != 0 {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	v628 = v603
	v629 = v604
	goto L195
L195:
	;
	if v605 != 0 {
		goto L200
	} else {
		goto L201
	}
L196:
	;
	v617 = v609 << (uint(int32(32)-base.I32_clz(v612)) % 32)
	goto L198
L197:
	;
	v617 = v612
	goto L198
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v617
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)+20))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v619)+28))
	v622 = F_repalloc(m, v620, v621)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L18
	} else {
		goto L199
	}
L199:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v624)+20)) = v622
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v626)+24))
	v628 = v626
	v629 = v627
	goto L195
L200:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v628)+20))
	base.MemoryCopy(m, v630+v629, v602, v605)
	goto L202
L201:
	;
	goto L202
L202:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v633)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v633)+24)) = v634 + v605
	goto L22
L203:
	;
	v644 = int32(1)
	v647 = v641 + v644
	if v647&v641 != 0 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	v663 = v638
	v664 = v639
	goto L205
L205:
	;
	if v640 != 0 {
		goto L210
	} else {
		goto L211
	}
L206:
	;
	v652 = v644 << (uint(int32(32)-base.I32_clz(v647)) % 32)
	goto L208
L207:
	;
	v652 = v647
	goto L208
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v638)+28)) = v652
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v654)+20))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v654)+28))
	v657 = F_repalloc(m, v655, v656)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L18
	} else {
		goto L209
	}
L209:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v659)+20)) = v657
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v661)+24))
	v663 = v661
	v664 = v662
	goto L205
L210:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v663)+20))
	base.MemoryCopy(m, v665+v664, v637, v640)
	goto L212
L211:
	;
	goto L212
L212:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v668)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v668)+24)) = v669 + v640
	goto L22
L213:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	*(*int32)(unsafe.Add(mBase, uint32(v680)+44)) = v682
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v686)))
	*(*int32)(unsafe.Add(mBase, uint32(v684))) = v685 - v687
	v691 = v679 & int32(-1024)
	if v691 == int32(_a_F_core_yylex_25) {
		goto L60
	} else {
		goto L214
	}
L214:
	;
	if v691 == int32(_a_F_core_yylex_26) {
		goto L24
	} else {
		goto L215
	}
L215:
	;
	F_addunicode(m, v679, v91)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L18
	} else {
		goto L216
	}
L216:
	;
	goto L33
L217:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+44)) = v708
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)))
	*(*int32)(unsafe.Add(mBase, uint32(v710))) = v711 - v713
	if v705&int32(-1024) != int32(_a_F_core_yylex_26) {
		goto L24
	} else {
		goto L218
	}
L218:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)+48))
	F_addunicode(m, v721<<(uint(int32(10))%32)&int32(_a_F_core_yylex_27)|v705&int32(1023)+int32(_a_F_core_yylex_28), v91)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L18
	} else {
		goto L219
	}
L219:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v734)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v733))) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(15)
	goto L22
L220:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L18
	} else {
		goto L221
	}
L221:
	;
	F_errmsg(m, int32(_a_F_core_yylex_29), int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L18
	} else {
		goto L222
	}
L222:
	;
	F_errhint(m, int32(_a_F_core_yylex_30), int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L18
	} else {
		goto L223
	}
L223:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v766)))
	F_scanner_errposition(m, v767, v91)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L18
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(_a_F_core_yylex_31), int32(682), int32(_a_F_core_yylex_32))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L18
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v791)+1)))
	v796 = base.I32_rotl(v792-int32(98), int32(31))
	if base.B2i32(base.Ui32(v796) <= base.Ui32(int32(10)))&(int32(base.Ui32(int32(1861))>>(uint(v796)%32))&int32(1)) == int32(0) {
		goto L234
	} else {
		goto L235
	}
L227:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v779)+16))
	switch v780 {
	case 0:
		goto L59
	default:
		goto L226
	case 2:
		goto L228
	}
L228:
	;
	v782 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[5]))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v782)+4))
	goto L229
L229:
	;
	if v783 < int32(35) {
		goto L226
	} else {
		goto L230
	}
L230:
	;
	v787 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[5]))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v787)+4))
	goto L231
L231:
	;
	if v788 <= int32(41) {
		goto L59
	} else {
		goto L232
	}
L232:
	;
	goto L226
L233:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v814)+24))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v814)+28))
	if v818 <= v815+int32(1) {
		goto L238
	} else {
		goto L239
	}
L234:
	;
	v806 = base.I32_extend8_s(v792)
	if int32(0) < v806 {
		v813 = v806
		goto L233
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796)+uint32(_c_F_core_yylex[6]))))
	v813 = v812
	goto L233
L237:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v810 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v809)+52)) = uint8(v810)
	v813 = v806
	goto L233
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v814)+28)) = v818 << (uint(int32(1)) % 32)
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v823)+20))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v823)+28))
	v826 = F_repalloc(m, v824, v825)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L18
	} else {
		goto L241
	}
L239:
	;
	v833 = v814
	v834 = v815
	goto L240
L240:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v833)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v834+v835))) = uint8(v813)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v838)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v838)+24)) = v839 + int32(1)
	goto L22
L241:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+20)) = v826
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v830)+24))
	v833 = v830
	v834 = v831
	goto L240
L242:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v851)+24))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v851)+28))
	if v855 <= v852+int32(1) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v851)+28)) = v855 << (uint(int32(1)) % 32)
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v860)+20))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v860)+28))
	v863 = F_repalloc(m, v861, v862)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L18
	} else {
		goto L246
	}
L244:
	;
	v869 = v851
	v870 = v852
	goto L245
L245:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v869)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v870+v871))) = uint8(v850)
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v874)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v874)+24)) = v875 + int32(1)
	if v850&int32(128) != 0 {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v865)+20)) = v863
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v867)+24))
	v869 = v867
	v870 = v868
	goto L245
L247:
	;
	v884 = int32(0)
	goto L249
L248:
	;
	v884 = v850 & int32(255)
	goto L249
L249:
	;
	if v884 != 0 {
		goto L22
	} else {
		goto L250
	}
L250:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v886 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v885)+52)) = uint8(v886)
	goto L22
L251:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v896)+24))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v896)+28))
	if v900 <= v897+int32(1) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v896)+28)) = v900 << (uint(int32(1)) % 32)
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v905)+20))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v905)+28))
	v908 = F_repalloc(m, v906, v907)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L18
	} else {
		goto L255
	}
L253:
	;
	v914 = v896
	v915 = v897
	goto L254
L254:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v914)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v915+v916))) = uint8(v895)
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v919)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v919)+24)) = v920 + int32(1)
	if v895&int32(128) != 0 {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v910)+20)) = v908
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v912)+24))
	v914 = v912
	v915 = v913
	goto L254
L256:
	;
	v929 = int32(0)
	goto L258
L257:
	;
	v929 = v895 & int32(255)
	goto L258
L258:
	;
	if v929 != 0 {
		goto L22
	} else {
		goto L259
	}
L259:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v931 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v930)+52)) = uint8(v931)
	goto L22
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v935)+28)) = v939 << (uint(int32(1)) % 32)
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v944)+20))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v944)+28))
	v947 = F_repalloc(m, v945, v946)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L18
	} else {
		goto L263
	}
L261:
	;
	v953 = v935
	v954 = v936
	goto L262
L262:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v953)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v954+v955))) = uint8(v934)
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v958)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v958)+24)) = v959 + int32(1)
	goto L22
L263:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v949)+20)) = v947
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v951)+24))
	v953 = v951
	v954 = v952
	goto L262
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v975)+40)) = v973
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(17)
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v979)+24)) = int32(0)
	goto L22
L266:
	;
	if v1029-v1030 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L267:
	;
	goto L266
L268:
	;
	v1014 = v1003
	v1015 = v1005
	goto L269
L269:
	;
	v1018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1015)+1)))
	v1019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014)+1)))
	if v1019 == int32(0) {
		v1029 = v1019
		v1030 = v1018
		goto L267
	} else {
		goto L271
	}
L270:
	;
	v1029 = v1019
	v1030 = v1018
	goto L267
L271:
	;
	v1022 = int32(1)
	if v1019 == v1018 {
		v1014 = v1014 + v1022
		v1015 = v1015 + v1022
		goto L269
	} else {
		goto L272
	}
L272:
	;
	goto L270
L273:
	;
	F_pfree(m, v1005)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L18
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+28))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+24))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1060 = v1058 - int32(1)
	if v1056 <= v1057+v1060 {
		goto L281
	} else {
		goto L282
	}
L276:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1036)+40)) = int32(0)
	v1039 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = v1039
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+24))
	v1045 = F_palloc(m, v1042+v1039)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L18
	} else {
		goto L277
	}
L277:
	;
	if v1042 != 0 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+20))
	base.MemoryCopy(m, v1045, v1048, v1042)
	goto L280
L279:
	;
	goto L280
L280:
	;
	v1051 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1042+v1045))) = uint8(v1051)
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1053))) = v1045
	v2416 = int32(261)
	goto L61
L281:
	;
	v1063 = int32(1)
	v1065 = v1057 + v1058
	if v1065&(v1065-v1063) != 0 {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	v1083 = v1004
	v1084 = v1057
	goto L283
L283:
	;
	if v1060 != 0 {
		goto L288
	} else {
		goto L289
	}
L284:
	;
	v1072 = v1063 << (uint(int32(32)-base.I32_clz(v1065)) % 32)
	goto L286
L285:
	;
	v1072 = v1065
	goto L286
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1004)+28)) = v1072
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1074)+20))
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1074)+28))
	v1077 = F_repalloc(m, v1075, v1076)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L18
	} else {
		goto L287
	}
L287:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+20)) = v1077
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+24))
	v1083 = v1081
	v1084 = v1082
	goto L283
L288:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+20))
	base.MemoryCopy(m, v1085+v1084, v1003, v1060)
	goto L290
L289:
	;
	goto L290
L290:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1088)+24)) = v1089 + v1060
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v1093)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v1097 = v1092 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v1097
	v1099 = v1097 + v173
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1099
	v1101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v1101)
	v1103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1099))) = uint8(v1103)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1099
	goto L22
L291:
	;
	v1113 = int32(1)
	v1116 = v1110 + v1113
	if v1116&v1110 != 0 {
		goto L294
	} else {
		goto L295
	}
L292:
	;
	v1132 = v1107
	v1133 = v1108
	goto L293
L293:
	;
	if v1109 != 0 {
		goto L298
	} else {
		goto L299
	}
L294:
	;
	v1121 = v1113 << (uint(int32(32)-base.I32_clz(v1116)) % 32)
	goto L296
L295:
	;
	v1121 = v1116
	goto L296
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1107)+28)) = v1121
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1123)+20))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1123)+28))
	v1126 = F_repalloc(m, v1124, v1125)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L18
	} else {
		goto L297
	}
L297:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1128)+20)) = v1126
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+24))
	v1132 = v1130
	v1133 = v1131
	goto L293
L298:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+20))
	base.MemoryCopy(m, v1134+v1133, v1106, v1109)
	goto L300
L299:
	;
	goto L300
L300:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1137)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1137)+24)) = v1138 + v1109
	goto L22
L301:
	;
	v1148 = int32(1)
	v1151 = v1145 + v1148
	if v1151&v1145 != 0 {
		goto L304
	} else {
		goto L305
	}
L302:
	;
	v1167 = v1142
	v1168 = v1143
	goto L303
L303:
	;
	if v1144 != 0 {
		goto L308
	} else {
		goto L309
	}
L304:
	;
	v1156 = v1148 << (uint(int32(32)-base.I32_clz(v1151)) % 32)
	goto L306
L305:
	;
	v1156 = v1151
	goto L306
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1142)+28)) = v1156
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+20))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+28))
	v1161 = F_repalloc(m, v1159, v1160)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L18
	} else {
		goto L307
	}
L307:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1163)+20)) = v1161
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+24))
	v1167 = v1165
	v1168 = v1166
	goto L303
L308:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1167)+20))
	base.MemoryCopy(m, v1169+v1168, v1141, v1144)
	goto L310
L309:
	;
	goto L310
L310:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1172)+24)) = v1173 + v1144
	goto L22
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1178)+28)) = v1182 << (uint(int32(1)) % 32)
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1187)+20))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1187)+28))
	v1190 = F_repalloc(m, v1188, v1189)
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L18
	} else {
		goto L314
	}
L312:
	;
	v1196 = v1178
	v1197 = v1179
	goto L313
L313:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1196)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v1197+v1198))) = uint8(v1177)
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1201)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1201)+24)) = v1202 + int32(1)
	goto L22
L314:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1192)+20)) = v1190
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1194)+24))
	v1196 = v1194
	v1197 = v1195
	goto L313
L315:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L316:
	;
	v1239 = F_palloc(m, v1234+int32(1))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L18
	} else {
		goto L317
	}
L317:
	;
	if v1234 != 0 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+20))
	base.MemoryCopy(m, v1239, v1242, v1234)
	goto L320
L319:
	;
	goto L320
L320:
	;
	v1245 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1234+v1239))) = uint8(v1245)
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+24))
	if int32(64) <= v1248 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	F_truncate_identifier(m, v1239, v1248, int32(1))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L18
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1254))) = v1239
	v2416 = int32(258)
	goto L61
L324:
	;
	goto L323
L325:
	;
	v1265 = F_palloc(m, v1260+int32(1))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L18
	} else {
		goto L326
	}
L326:
	;
	if v1260 != 0 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1267)+20))
	base.MemoryCopy(m, v1265, v1268, v1260)
	goto L329
L328:
	;
	goto L329
L329:
	;
	v1271 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1260+v1265))) = uint8(v1271)
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1273))) = v1265
	v2416 = int32(259)
	goto L61
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+28)) = v1280 << (uint(int32(1)) % 32)
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1285)+20))
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1285)+28))
	v1288 = F_repalloc(m, v1286, v1287)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L18
	} else {
		goto L333
	}
L331:
	;
	v1295 = v1276
	v1296 = v1277
	goto L332
L332:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+20))
	v1299 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v1296+v1297))) = uint8(v1299)
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1301)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1301)+24)) = v1302 + int32(1)
	goto L22
L333:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1290)+20)) = v1288
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+24))
	v1295 = v1292
	v1296 = v1293
	goto L332
L334:
	;
	v1313 = int32(1)
	v1316 = v1310 + v1313
	if v1316&v1310 != 0 {
		goto L337
	} else {
		goto L338
	}
L335:
	;
	v1332 = v1307
	v1333 = v1308
	goto L336
L336:
	;
	if v1309 != 0 {
		goto L341
	} else {
		goto L342
	}
L337:
	;
	v1321 = v1313 << (uint(int32(32)-base.I32_clz(v1316)) % 32)
	goto L339
L338:
	;
	v1321 = v1316
	goto L339
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1307)+28)) = v1321
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1323)+20))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1323)+28))
	v1326 = F_repalloc(m, v1324, v1325)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L18
	} else {
		goto L340
	}
L340:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1328)+20)) = v1326
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1330)+24))
	v1332 = v1330
	v1333 = v1331
	goto L336
L341:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+20))
	base.MemoryCopy(m, v1334+v1333, v1306, v1309)
	goto L343
L342:
	;
	goto L343
L343:
	;
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1337)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1337)+24)) = v1338 + v1309
	goto L22
L344:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L345:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1368))) = v1366
	v2416 = int32(258)
	goto L61
L346:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1525)))
	*(*int32)(unsafe.Add(mBase, uint32(v1524))) = v1427 - v1526
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	if v1529 <= v1514 {
		goto L371
	} else {
		goto L372
	}
L347:
	;
	v1436 = v1431
	goto L349
L348:
	;
	v1436 = v1429
	goto L349
L349:
	;
	if v1429 != 0 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1437 = v1436
	goto L352
L351:
	;
	v1437 = v1431
	goto L352
L352:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	if v1437 != 0 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1440 = v1437 - v1427
	goto L355
L354:
	;
	v1440 = v1439
	goto L355
L355:
	;
	if v1440 < int32(2) {
		v1514 = v1440
		goto L346
	} else {
		goto L356
	}
L356:
	;
	v1446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1427+v1440-int32(1)))))
	switch v1446 - int32(43) {
	case 0, 2:
		goto L357
	default:
		v1514 = v1440
		goto L346
	}
L357:
	;
	v1457 = v1440 - int32(2)
	goto L358
L358:
	;
	v1465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1427+v1457))))
	switch v1465 - int32(33) {
	case 0, 2, 4, 5, 30, 31, 61, 63, 91, 93:
		v1514 = v1440
		goto L346
	default:
		goto L360
	}
L359:
	;
	if v1440 != int32(2) {
		goto L362
	} else {
		goto L363
	}
L360:
	;
	if int32(0) < v1457 {
		v1457 = v1457 - int32(1)
		goto L358
	} else {
		goto L361
	}
L361:
	;
	goto L359
L362:
	;
	v1480 = v1440
	goto L365
L363:
	;
	goto L364
L364:
	;
	v1514 = int32(1)
	goto L346
L365:
	;
	v1488 = v1480 - int32(1)
	v1492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1427+v1480-int32(2)))))
	switch v1492 - int32(43) {
	case 0, 2:
		goto L367
	default:
		v1514 = v1488
		goto L346
	}
L366:
	;
	goto L364
L367:
	;
	if int32(3) < v1480 {
		v1480 = v1488
		goto L365
	} else {
		goto L368
	}
L368:
	;
	goto L366
L369:
	;
	v1681 = F_pstrdup(m, v1679)
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L18
	} else {
		goto L411
	}
L370:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1572 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1571))))
	goto L389
L371:
	;
	if v1514 <= int32(63) {
		goto L381
	} else {
		goto L382
	}
L372:
	;
	v1531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v1531)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v1534 = v1514 + v173
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1534
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v1514
	v1537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1534))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v1537)
	v1539 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1534))) = uint8(v1539)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1534
	switch v1514 - int32(1) {
	case 0:
		goto L370
	case 1:
		goto L373
	default:
		goto L371
	}
L373:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1544))))
	switch v1545 - int32(33) {
	case 0:
		goto L374
	default:
		v1679 = v1544
		goto L369
	case 27:
		goto L375
	case 28:
		goto L377
	case 29:
		goto L376
	}
L374:
	;
	v1560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1544)+1)))
	if v1560 != int32(61) {
		v1679 = v1544
		goto L369
	} else {
		goto L380
	}
L375:
	;
	v1557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1544)+1)))
	switch v1557 - int32(61) {
	case 0:
		v2416 = int32(272)
		goto L61
	case 1:
		goto L62
	default:
		v1679 = v1544
		goto L369
	}
L376:
	;
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1544)+1)))
	if v1552 != int32(61) {
		v1679 = v1544
		goto L369
	} else {
		goto L379
	}
L377:
	;
	v1548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1544)+1)))
	if v1548 != int32(62) {
		v1679 = v1544
		goto L369
	} else {
		goto L378
	}
L378:
	;
	v2416 = int32(271)
	goto L61
L379:
	;
	v2416 = int32(273)
	goto L61
L380:
	;
	goto L62
L381:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1679 = v1566
	goto L369
L382:
	;
	goto L383
L383:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_33), v91)
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L18
	} else {
		goto L384
	}
L384:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L385:
	;
	if v1678 != 0 {
		v2416 = v1572
		goto L61
	} else {
		goto L410
	}
L386:
	;
	v1678 = int32(0)
	goto L385
L387:
	;
	v1656 = v1649
	v1658 = v1651
	goto L404
L388:
	;
	if base.B2i32(v1595 != v1596) == int32(0) {
		goto L386
	} else {
		goto L395
	}
L389:
	;
	v1587 = int32(_a_F_core_yylex_34)
	v1589 = int32(18)
	goto L390
L390:
	;
	v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1587))))
	if v1592 == v1572&int32(255) {
		v1649 = v1587
		v1651 = v1589
		goto L387
	} else {
		goto L392
	}
L391:
	;
	goto L388
L392:
	;
	v1594 = int32(1)
	v1595 = v1589 - v1594
	v1596 = int32(0)
	v1599 = v1587 + v1594
	if v1599&int32(3) == v1596 {
		goto L388
	} else {
		goto L393
	}
L393:
	;
	if v1595 != 0 {
		v1587 = v1599
		v1589 = v1595
		goto L390
	} else {
		goto L394
	}
L394:
	;
	goto L391
L395:
	;
	v1612 = v1572 & int32(255)
	v1613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1599))))
	if base.B2i32(v1612 == v1613)|base.B2i32(base.Ui32(v1595) < base.Ui32(int32(4))) == int32(0) {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v1622 = v1599
	v1624 = v1595
	goto L399
L397:
	;
	v1642 = v1599
	v1644 = v1595
	goto L398
L398:
	;
	if v1644 == int32(0) {
		goto L386
	} else {
		goto L403
	}
L399:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1622)))
	v1629 = v1628 ^ v1612*int32(16843009)
	v1632 = int32(-2139062144)
	if (int32(16843008)-v1629|v1629)&v1632 != v1632 {
		v1649 = v1622
		v1651 = v1624
		goto L387
	} else {
		goto L401
	}
L400:
	;
	v1642 = v1637
	v1644 = v1639
	goto L398
L401:
	;
	v1636 = int32(4)
	v1637 = v1622 + v1636
	v1639 = v1624 - v1636
	if base.Ui32(int32(3)) < base.Ui32(v1639) {
		v1622 = v1637
		v1624 = v1639
		goto L399
	} else {
		goto L402
	}
L402:
	;
	goto L400
L403:
	;
	v1649 = v1642
	v1651 = v1644
	goto L387
L404:
	;
	v1661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1656))))
	if v1572&int32(255) == v1661 {
		goto L406
	} else {
		goto L407
	}
L405:
	;
	goto L386
L406:
	;
	v1678 = v1656
	goto L385
L407:
	;
	goto L408
L408:
	;
	v1663 = int32(1)
	v1666 = v1658 - v1663
	if v1666 != 0 {
		v1656 = v1656 + v1663
		v1658 = v1666
		goto L404
	} else {
		goto L409
	}
L409:
	;
	goto L405
L410:
	;
	v1679 = v1571
	goto L369
L411:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1683))) = v1681
	v2416 = int32(265)
	goto L61
L412:
	;
	v1703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v1703 == int32(1) {
		goto L56
	} else {
		goto L413
	}
L413:
	;
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1706))) = v1701
	v2416 = int32(267)
	goto L61
L414:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L415:
	;
	v1735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v1735 == int32(1) {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v1739 = F_pstrdup(m, v1725)
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L18
	} else {
		goto L419
	}
L417:
	;
	v1741 = int32(266)
	v1742 = v1732
	goto L418
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1724))) = v1742
	v2416 = v1741
	goto L61
L419:
	;
	v1741 = int32(260)
	v1742 = v1739
	goto L418
L420:
	;
	v1761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v1761 == int32(1) {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v1765 = F_pstrdup(m, v1751)
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L18
	} else {
		goto L424
	}
L422:
	;
	v1767 = int32(266)
	v1768 = v1758
	goto L423
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1750))) = v1768
	v2416 = v1767
	goto L61
L424:
	;
	v1767 = int32(260)
	v1768 = v1765
	goto L423
L425:
	;
	v1787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v1787 == int32(1) {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v1791 = F_pstrdup(m, v1777)
	mBase = m.M
	v1792 = m.ExcPending
	if v1792 != 0 {
		goto L18
	} else {
		goto L429
	}
L427:
	;
	v1793 = int32(266)
	v1794 = v1784
	goto L428
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1776))) = v1794
	v2416 = v1793
	goto L61
L429:
	;
	v1793 = int32(260)
	v1794 = v1791
	goto L428
L430:
	;
	v1813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v1813 == int32(1) {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1817 = F_pstrdup(m, v1803)
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L18
	} else {
		goto L434
	}
L432:
	;
	v1819 = int32(266)
	v1820 = v1810
	goto L433
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1802))) = v1820
	v2416 = v1819
	goto L61
L434:
	;
	v1819 = int32(260)
	v1820 = v1817
	goto L433
L435:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L436:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L437:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L438:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1858))) = v1856
	v2416 = int32(260)
	goto L61
L439:
	;
	v1892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v1892 == int32(1) {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v1896 = F_pstrdup(m, v1882)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L18
	} else {
		goto L443
	}
L441:
	;
	v1898 = int32(266)
	v1899 = v1889
	goto L442
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1881))) = v1899
	v2416 = v1898
	goto L61
L443:
	;
	v1898 = int32(260)
	v1899 = v1896
	goto L442
L444:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1910))) = v1908
	v2416 = int32(260)
	goto L61
L445:
	;
	if int32(0) <= v1922 {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v1927)+8))
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1928)))
	v1931 = v1922 << (uint(int32(1)) % 32)
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+4))
	v1934 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1931+v1932))))
	*(*int32)(unsafe.Add(mBase, uint32(v1926))) = v1929 + v1934
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1937)+12))
	v1940 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1938+v1931))))
	v2416 = v1940
	goto L61
L447:
	;
	goto L448
L448:
	;
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1944 = F_downcase_truncate_identifier(m, v1941, v1942, int32(1))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L18
	} else {
		goto L449
	}
L449:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1946))) = v1944
	v2416 = int32(258)
	goto L61
L450:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L451:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v1975)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v1979
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(v1974)))
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1981))) = v1982
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v1986 = int32(2)
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(v1984+v1985<<(uint(v1986)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1989)+44)) = int32(1)
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1992+v1993<<(uint(v1986)%32))))
	v1998 = v1997
	v1999 = v1992
	v2000 = v1993
	goto L453
L452:
	;
	v1998 = v1975
	v1999 = v1970
	v2000 = v1971
	goto L453
L453:
	;
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v91)+36))
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v1998)+4))
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v2004 = v2002 + v2003
	if base.Ui32(v2001) <= base.Ui32(v2004) {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2009 = v1967 ^ int32(-1) + v170
	v2010 = v2006 + v2009
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v2010
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v91)+44))
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v2012<<(uint(int32(2))%32))+uint32(_c_F_core_yylex[2])))
	if v2009 <= int32(0) {
		v3081 = v2017
		goto L37
	} else {
		goto L457
	}
L455:
	;
	goto L456
L456:
	;
	if base.Ui32(v2004+int32(1)) < base.Ui32(v2001) {
		goto L54
	} else {
		goto L465
	}
L457:
	;
	v2023 = int32(0)
	v2025 = v2009 & int32(3)
	if v2025 == v2023 {
		goto L55
	} else {
		goto L458
	}
L458:
	;
	v2028 = v2023
	v2029 = v2006
	v2031 = v2017
	goto L459
L459:
	;
	v2042 = v2029 + int32(1)
	v2043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2029))))
	if v2043 != 0 {
		goto L461
	} else {
		goto L462
	}
L460:
	;
	v3008 = v2042
	v3010 = v2052
	goto L40
L461:
	;
	v2045 = v2043
	goto L463
L462:
	;
	v2045 = int32(256)
	goto L463
L463:
	;
	v2046 = int32(3)
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v2031+v2045<<(uint(v2046)%32))+4))
	v2052 = v2031 + v2049<<(uint(v2046)%32)
	v2054 = v2028 + int32(1)
	if v2054 != v2025 {
		v2028 = v2054
		v2029 = v2042
		v2031 = v2052
		goto L459
	} else {
		goto L464
	}
L464:
	;
	goto L460
L465:
	;
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(v1998)+40))
	if v2060 == int32(0) {
		goto L466
	} else {
		goto L467
	}
L466:
	;
	if v2001-v2059 != int32(1) {
		v2763 = v2002
		v2768 = v2059
		v2770 = v2003
		goto L50
	} else {
		goto L469
	}
L467:
	;
	goto L468
L468:
	;
	v2068 = v2059 ^ int32(-1) + v2001
	if int32(0) < v2068 {
		goto L470
	} else {
		goto L471
	}
L469:
	;
	v2989 = v2059
	goto L41
L470:
	;
	v2071 = int32(7)
	v2072 = v2068 & v2071
	if base.Ui32(v2001-v2059-int32(2)) < base.Ui32(v2071) {
		goto L475
	} else {
		goto L476
	}
L471:
	;
	v2177 = v1998
	v2180 = v1999
	v2183 = v2000
	goto L472
L472:
	;
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v2177)+44))
	if v2187 == int32(2) {
		goto L485
	} else {
		goto L486
	}
L473:
	;
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v2168+v2169<<(uint(int32(2))%32))))
	v2177 = v2173
	v2180 = v2168
	v2183 = v2169
	goto L472
L474:
	;
	v2133 = v2119
	v2136 = v2122
	v2139 = int32(0)
	goto L482
L475:
	;
	v2119 = v2002
	v2122 = v2059
	goto L474
L476:
	;
	goto L477
L477:
	;
	v2081 = v2002
	v2084 = v2059
	v2087 = int32(0)
	goto L478
L478:
	;
	v2094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2084))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2081))) = uint8(v2094)
	v2096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2084)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2081)+1)) = uint8(v2096)
	v2098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2084)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2081)+2)) = uint8(v2098)
	v2100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2084)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2081)+3)) = uint8(v2100)
	v2102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2084)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2081)+4)) = uint8(v2102)
	v2104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2084)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2081)+5)) = uint8(v2104)
	v2106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2084)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2081)+6)) = uint8(v2106)
	v2108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2084)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2081)+7)) = uint8(v2108)
	v2110 = int32(8)
	v2111 = v2081 + v2110
	v2113 = v2084 + v2110
	v2115 = v2087 + v2110
	if v2115 != v2068&int32(2147483640) {
		v2081 = v2111
		v2084 = v2113
		v2087 = v2115
		goto L478
	} else {
		goto L480
	}
L479:
	;
	if v2072 == int32(0) {
		goto L473
	} else {
		goto L481
	}
L480:
	;
	goto L479
L481:
	;
	v2119 = v2111
	v2122 = v2113
	goto L474
L482:
	;
	v2146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2136))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2133))) = uint8(v2146)
	v2148 = int32(1)
	v2153 = v2139 + v2148
	if v2153 != v2072 {
		v2133 = v2133 + v2148
		v2136 = v2136 + v2148
		v2139 = v2153
		goto L482
	} else {
		goto L484
	}
L483:
	;
	goto L473
L484:
	;
	goto L483
L485:
	;
	v2190 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v2190
	v2500 = v2190
	v2506 = v2180 + v2183<<(uint(int32(2))%32)
	goto L51
L486:
	;
	goto L487
L487:
	;
	v2196 = int32(0)
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v2177)+12))
	v2198 = v2059 - v2001
	v2199 = v2197 + v2198
	if v2199 <= v2196 {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(v91)+36))
	v2206 = v2177
	v2208 = v2202
	v2212 = v2197
	goto L491
L489:
	;
	v2254 = v2199
	v2257 = v2177
	goto L490
L490:
	;
	v2267 = int32(_a_F_core_yylex_35)
	if base.Ui32(v2267) <= base.Ui32(v2254) {
		goto L507
	} else {
		goto L508
	}
L491:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v2206)+20))
	if v2216 == int32(0) {
		goto L493
	} else {
		goto L494
	}
L492:
	;
	v2254 = v2251
	v2257 = v2249
	goto L490
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2206)+4)) = int32(0)
	goto L42
L494:
	;
	goto L495
L495:
	;
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2206)+4))
	v2223 = v2212 << (uint(int32(1)) % 32)
	if v2223 <= int32(0) {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	v2227 = base.I32_div_s(v2212, int32(8))
	v2229 = v2227 + v2212
	goto L498
L497:
	;
	v2229 = v2223
	goto L498
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2206)+12)) = v2229
	v2232 = v2229 + int32(2)
	if v2221 != 0 {
		goto L500
	} else {
		goto L501
	}
L499:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2206)+4)) = v2237
	if v2237 == int32(0) {
		goto L42
	} else {
		goto L505
	}
L500:
	;
	v2233 = F_repalloc(m, v2221, v2232)
	mBase = m.M
	v2234 = m.ExcPending
	if v2234 != 0 {
		goto L18
	} else {
		goto L503
	}
L501:
	;
	goto L502
L502:
	;
	v2235 = F_palloc(m, v2232)
	mBase = m.M
	v2236 = m.ExcPending
	if v2236 != 0 {
		goto L18
	} else {
		goto L504
	}
L503:
	;
	v2237 = v2233
	goto L499
L504:
	;
	v2237 = v2235
	goto L499
L505:
	;
	v2242 = v2237 + (v2208 - v2221)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v2242
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v2244+v2245<<(uint(int32(2))%32))))
	v2250 = *(*int32)(unsafe.Add(mBase, uint32(v2249)+12))
	v2251 = v2250 + v2198
	if v2251 <= int32(0) {
		v2206 = v2249
		v2208 = v2242
		v2212 = v2250
		goto L491
	} else {
		goto L506
	}
L506:
	;
	goto L492
L507:
	;
	v2270 = v2267
	goto L509
L508:
	;
	v2270 = v2254
	goto L509
L509:
	;
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(v2257)+24))
	if v2271 != 0 {
		goto L510
	} else {
		goto L511
	}
L510:
	;
	v2279 = v2196
	goto L514
L511:
	;
	goto L512
L512:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_core_yylex[7])) = int32(0)
	v2333 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(v2333+v2334<<(uint(int32(2))%32))))
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(v2338)+4))
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2343 = F_fread(m, v2339+v2068, int32(1), v2270, v2342)
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		goto L18
	} else {
		goto L525
	}
L513:
	;
	switch v2289 {
	case 0:
		goto L521
	default:
		v2328 = v2303
		goto L519
	case 11:
		goto L520
	}
L514:
	;
	v2285 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2286 = F_do_getc(m, v2285)
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L18
	} else {
		goto L517
	}
L515:
	;
	v2303 = v2270
	goto L513
L516:
	;
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2291 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(v2290+v2291<<(uint(int32(2))%32))))
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2295)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v2296+v2068+v2279))) = uint8(v2286)
	v2301 = v2279 + int32(1)
	if v2301 != v2270 {
		v2279 = v2301
		goto L514
	} else {
		goto L518
	}
L517:
	;
	v2289 = v2286 + int32(1)
	switch v2289 {
	case 0, 11:
		v2303 = v2279
		goto L513
	default:
		goto L516
	}
L518:
	;
	goto L515
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v2328
	v2482 = v2328
	goto L52
L520:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2316 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v2315+v2316<<(uint(int32(2))%32))))
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(v2320)+4))
	v2324 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v2321+v2068+v2303))) = uint8(v2324)
	v2328 = v2303 + int32(1)
	goto L519
L521:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(v2304)))
	goto L522
L522:
	;
	if int32(base.Ui32(v2305)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v2328 = v2303
		goto L519
	} else {
		goto L523
	}
L523:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_4))
	mBase = m.M
	v2314 = m.ExcPending
	if v2314 != 0 {
		goto L18
	} else {
		goto L524
	}
L524:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L525:
	;
	v2352 = v2343
	goto L526
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v2352
	if v2352 != 0 {
		v2482 = v2352
		goto L52
	} else {
		goto L528
	}
L528:
	;
	v2359 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(v2359)))
	goto L529
L529:
	;
	if int32(base.Ui32(v2360)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	v2482 = int32(0)
	goto L52
L531:
	;
	goto L532
L532:
	;
	v2369 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[7]))
	if v2369 != int32(27) {
		goto L53
	} else {
		goto L533
	}
L533:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_core_yylex[7])) = int32(0)
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2376 = *(*int32)(unsafe.Add(mBase, uint32(v2375)))
	*(*int32)(unsafe.Add(mBase, uint32(v2375))) = v2376 & int32(-49)
	goto L534
L534:
	;
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2385 = *(*int32)(unsafe.Add(mBase, uint32(v2380+v2381<<(uint(int32(2))%32))))
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+4))
	v2389 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2390 = F_fread(m, v2386+v2068, int32(1), v2270, v2389)
	mBase = m.M
	v2391 = m.ExcPending
	if v2391 != 0 {
		goto L18
	} else {
		goto L535
	}
L535:
	;
	v2352 = v2390
	goto L526
L536:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L537:
	;
	F_errcode(m, int32(100794498))
	mBase = m.M
	v2442 = m.ExcPending
	if v2442 != 0 {
		goto L18
	} else {
		goto L538
	}
L538:
	;
	F_errmsg(m, int32(_a_F_core_yylex_36), int32(0))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L18
	} else {
		goto L539
	}
L539:
	;
	F_errhint(m, int32(_a_F_core_yylex_37), int32(0))
	mBase = m.M
	v2450 = m.ExcPending
	if v2450 != 0 {
		goto L18
	} else {
		goto L540
	}
L540:
	;
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v2452 = *(*int32)(unsafe.Add(mBase, uint32(v2451)))
	F_scanner_errposition(m, v2452, v91)
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		goto L18
	} else {
		goto L541
	}
L541:
	;
	F_errfinish(m, int32(_a_F_core_yylex_31), int32(694), int32(_a_F_core_yylex_32))
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L18
	} else {
		goto L542
	}
L542:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L543:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L544:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L545:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L546:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L547:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L548:
	;
	v2633 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v2634 = v2633 + v2068
	v2635 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2636 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2640 = *(*int32)(unsafe.Add(mBase, uint32(v2635+v2636<<(uint(int32(2))%32))))
	v2641 = *(*int32)(unsafe.Add(mBase, uint32(v2640)+12))
	if v2641 < v2634 {
		goto L572
	} else {
		goto L573
	}
L549:
	;
	if v2068 == int32(0) {
		goto L550
	} else {
		goto L551
	}
L550:
	;
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v2514 != 0 {
		goto L555
	} else {
		goto L556
	}
L551:
	;
	goto L552
L552:
	;
	v2618 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2619 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2620 = int32(2)
	v2623 = *(*int32)(unsafe.Add(mBase, uint32(v2618+v2619<<(uint(v2620)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2623)+44)) = v2620
	v2632 = v2620
	goto L548
L553:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2580)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2580))) = v2513
	v2587 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v2587 != 0 {
		goto L568
	} else {
		goto L569
	}
L554:
	;
	v2537 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[7]))
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2542 = *(*int32)(unsafe.Add(mBase, uint32(v2535+v2538<<(uint(int32(2))%32))))
	if v2542 == int32(0) {
		goto L562
	} else {
		goto L563
	}
L555:
	;
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2519 = *(*int32)(unsafe.Add(mBase, uint32(v2514+v2515<<(uint(int32(2))%32))))
	if v2519 != 0 {
		v2535 = v2514
		goto L554
	} else {
		goto L558
	}
L556:
	;
	goto L557
L557:
	;
	F_core_yyensure_buffer_stack(m, v91)
	mBase = m.M
	v2521 = m.ExcPending
	if v2521 != 0 {
		goto L18
	} else {
		goto L559
	}
L558:
	;
	goto L557
L559:
	;
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2523 = F_core_yy_create_buffer(m, v2522, v91)
	mBase = m.M
	v2524 = m.ExcPending
	if v2524 != 0 {
		goto L18
	} else {
		goto L560
	}
L560:
	;
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2526 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2525+v2526<<(uint(int32(2))%32)))) = v2523
	v2531 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v2531 != 0 {
		v2535 = v2531
		goto L554
	} else {
		goto L561
	}
L561:
	;
	v2533 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[7]))
	v2580 = int32(0)
	v2583 = v2533
	goto L553
L562:
	;
	v2580 = int32(0)
	v2583 = v2537
	goto L553
L563:
	;
	goto L564
L564:
	;
	v2546 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2542)+16)) = v2546
	v2548 = *(*int32)(unsafe.Add(mBase, uint32(v2542)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v2548))) = uint8(v2546)
	v2551 = *(*int32)(unsafe.Add(mBase, uint32(v2542)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v2551)+1)) = uint8(v2546)
	*(*int32)(unsafe.Add(mBase, uint32(v2542)+44)) = v2546
	*(*int32)(unsafe.Add(mBase, uint32(v2542)+28)) = int32(1)
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(v2542)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2542)+8)) = v2558
	v2560 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v2560 == v2546 {
		v2580 = v2542
		v2583 = v2537
		goto L553
	} else {
		goto L565
	}
L565:
	;
	v2563 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2566 = v2560 + v2563<<(uint(int32(2))%32)
	v2567 = *(*int32)(unsafe.Add(mBase, uint32(v2566)))
	if v2542 != v2567 {
		v2580 = v2542
		v2583 = v2537
		goto L553
	} else {
		goto L566
	}
L566:
	;
	v2569 = *(*int32)(unsafe.Add(mBase, uint32(v2567)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v2569
	v2571 = *(*int32)(unsafe.Add(mBase, uint32(v2566)))
	v2572 = *(*int32)(unsafe.Add(mBase, uint32(v2571)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v2572
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v2572
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(v2566)))
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(v2575)))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v2576
	v2578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2572))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v2578)
	v2580 = v2542
	v2583 = v2537
	goto L553
L567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2580)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_core_yylex[7])) = v2583
	v2600 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2601 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2604 = v2600 + v2601<<(uint(int32(2))%32)
	v2605 = *(*int32)(unsafe.Add(mBase, uint32(v2604)))
	v2606 = *(*int32)(unsafe.Add(mBase, uint32(v2605)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v2606
	v2608 = *(*int32)(unsafe.Add(mBase, uint32(v2604)))
	v2609 = *(*int32)(unsafe.Add(mBase, uint32(v2608)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v2609
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v2609
	v2612 = *(*int32)(unsafe.Add(mBase, uint32(v2604)))
	v2613 = *(*int32)(unsafe.Add(mBase, uint32(v2612)))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v2613
	v2615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2609))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v2615)
	v2632 = int32(1)
	goto L548
L568:
	;
	v2588 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2592 = *(*int32)(unsafe.Add(mBase, uint32(v2587+v2588<<(uint(int32(2))%32))))
	if v2580 == v2592 {
		goto L567
	} else {
		goto L571
	}
L569:
	;
	goto L570
L570:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2580)+32)) = int64(1)
	goto L567
L571:
	;
	goto L570
L572:
	;
	v2645 = v2634 + v2633>>(uint(int32(1))%32)
	v2646 = *(*int32)(unsafe.Add(mBase, uint32(v2640)+4))
	if v2646 != 0 {
		goto L576
	} else {
		goto L577
	}
L573:
	;
	v2675 = v2634
	v2676 = v2635
	v2678 = v2636
	goto L574
L574:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v2675
	v2680 = int32(2)
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v2676+v2678<<(uint(v2680)%32))))
	v2684 = *(*int32)(unsafe.Add(mBase, uint32(v2683)+4))
	v2686 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2684+v2675))) = uint8(v2686)
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2689 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2693 = *(*int32)(unsafe.Add(mBase, uint32(v2688+v2689<<(uint(v2680)%32))))
	v2694 = *(*int32)(unsafe.Add(mBase, uint32(v2693)+4))
	v2695 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v2694+v2695)+1)) = uint8(v2686)
	v2699 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2700 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2703 = v2699 + v2700<<(uint(v2680)%32)
	v2704 = *(*int32)(unsafe.Add(mBase, uint32(v2703)))
	v2705 = *(*int32)(unsafe.Add(mBase, uint32(v2704)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v2705
	if v2632 == int32(1) {
		v2989 = v2705
		goto L41
	} else {
		goto L582
	}
L575:
	;
	v2652 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2653 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2654 = int32(2)
	v2657 = *(*int32)(unsafe.Add(mBase, uint32(v2652+v2653<<(uint(v2654)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2657)+4)) = v2651
	v2659 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2660 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2664 = *(*int32)(unsafe.Add(mBase, uint32(v2659+v2660<<(uint(v2654)%32))))
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v2664)+4))
	if v2665 == int32(0) {
		goto L48
	} else {
		goto L581
	}
L576:
	;
	v2647 = F_repalloc(m, v2646, v2645)
	mBase = m.M
	v2648 = m.ExcPending
	if v2648 != 0 {
		goto L18
	} else {
		goto L579
	}
L577:
	;
	goto L578
L578:
	;
	v2649 = F_palloc(m, v2645)
	mBase = m.M
	v2650 = m.ExcPending
	if v2650 != 0 {
		goto L18
	} else {
		goto L580
	}
L579:
	;
	v2651 = v2647
	goto L575
L580:
	;
	v2651 = v2649
	goto L575
L581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2664)+12)) = v2645 - int32(2)
	v2671 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2673 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v2675 = v2673 + v2068
	v2676 = v2672
	v2678 = v2671
	goto L574
L582:
	;
	switch v2632 - int32(1) {
	case 0:
		goto L49
	case 1:
		goto L583
	default:
		goto L584
	}
L583:
	;
	v2760 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v2761 = *(*int32)(unsafe.Add(mBase, uint32(v2703)))
	v2762 = *(*int32)(unsafe.Add(mBase, uint32(v2761)+4))
	v2763 = v2762
	v2768 = v2705
	v2770 = v2760
	goto L50
L584:
	;
	v2713 = v1967 ^ int32(-1) + v170
	v2714 = v2705 + v2713
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v2714
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(v91)+44))
	v2721 = *(*int32)(unsafe.Add(mBase, uint32(v2716<<(uint(int32(2))%32))+uint32(_c_F_core_yylex[2])))
	if v2713 <= int32(0) {
		v2969 = v2721
		goto L43
	} else {
		goto L585
	}
L585:
	;
	v2727 = int32(0)
	v2729 = v2713 & int32(3)
	if v2729 == v2727 {
		goto L47
	} else {
		goto L586
	}
L586:
	;
	v2732 = v2727
	v2733 = v2705
	v2735 = v2721
	goto L587
L587:
	;
	v2746 = v2733 + int32(1)
	v2747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2733))))
	if v2747 != 0 {
		goto L589
	} else {
		goto L590
	}
L588:
	;
	v2896 = v2746
	v2898 = v2756
	goto L44
L589:
	;
	v2749 = v2747
	goto L591
L590:
	;
	v2749 = int32(256)
	goto L591
L591:
	;
	v2750 = int32(3)
	v2753 = *(*int32)(unsafe.Add(mBase, uint32(v2735+v2749<<(uint(v2750)%32))+4))
	v2756 = v2735 + v2753<<(uint(v2750)%32)
	v2758 = v2732 + int32(1)
	if v2758 != v2729 {
		v2732 = v2758
		v2733 = v2746
		v2735 = v2756
		goto L587
	} else {
		goto L592
	}
L592:
	;
	goto L588
L593:
	;
	v2785 = v2776 - v2768
	v2788 = int32(0)
	v2790 = v2785 & int32(3)
	if v2790 == v2788 {
		goto L46
	} else {
		goto L594
	}
L594:
	;
	v2793 = v2768
	v2796 = v2783
	v2799 = v2788
	goto L595
L595:
	;
	v2807 = v2793 + int32(1)
	v2808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2793))))
	if v2808 != 0 {
		goto L597
	} else {
		goto L598
	}
L596:
	;
	v2824 = v2807
	v2827 = v2817
	goto L45
L597:
	;
	v2810 = v2808
	goto L599
L598:
	;
	v2810 = int32(256)
	goto L599
L599:
	;
	v2811 = int32(3)
	v2814 = *(*int32)(unsafe.Add(mBase, uint32(v2796+v2810<<(uint(v2811)%32))+4))
	v2817 = v2796 + v2814<<(uint(v2811)%32)
	v2819 = v2799 + int32(1)
	if v2819 != v2790 {
		v2793 = v2807
		v2796 = v2817
		v2799 = v2819
		goto L595
	} else {
		goto L600
	}
L600:
	;
	goto L596
L601:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L602:
	;
	v2839 = v2824
	v2842 = v2827
	goto L603
L603:
	;
	v2852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2839))))
	if v2852 != 0 {
		goto L605
	} else {
		goto L606
	}
L604:
	;
	v170 = v2776
	v172 = v2891
	v173 = v2768
	goto L35
L605:
	;
	v2854 = v2852
	goto L607
L606:
	;
	v2854 = int32(256)
	goto L607
L607:
	;
	v2855 = int32(3)
	v2858 = *(*int32)(unsafe.Add(mBase, uint32(v2842+v2854<<(uint(v2855)%32))+4))
	v2861 = v2842 + v2858<<(uint(v2855)%32)
	v2862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2839)+1)))
	if v2862 != 0 {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	v2864 = v2862
	goto L610
L609:
	;
	v2864 = int32(256)
	goto L610
L610:
	;
	v2865 = int32(3)
	v2868 = *(*int32)(unsafe.Add(mBase, uint32(v2861+v2864<<(uint(v2865)%32))+4))
	v2871 = v2861 + v2868<<(uint(v2865)%32)
	v2872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2839)+2)))
	if v2872 != 0 {
		goto L611
	} else {
		goto L612
	}
L611:
	;
	v2874 = v2872
	goto L613
L612:
	;
	v2874 = int32(256)
	goto L613
L613:
	;
	v2875 = int32(3)
	v2878 = *(*int32)(unsafe.Add(mBase, uint32(v2871+v2874<<(uint(v2875)%32))+4))
	v2881 = v2871 + v2878<<(uint(v2875)%32)
	v2882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2839)+3)))
	if v2882 != 0 {
		goto L614
	} else {
		goto L615
	}
L614:
	;
	v2884 = v2882
	goto L616
L615:
	;
	v2884 = int32(256)
	goto L616
L616:
	;
	v2885 = int32(3)
	v2888 = *(*int32)(unsafe.Add(mBase, uint32(v2881+v2884<<(uint(v2885)%32))+4))
	v2891 = v2881 + v2888<<(uint(v2885)%32)
	v2893 = v2839 + int32(4)
	if v2893 != v2776 {
		v2839 = v2893
		v2842 = v2891
		goto L603
	} else {
		goto L617
	}
L617:
	;
	goto L604
L618:
	;
	v2911 = v2896
	v2913 = v2898
	goto L619
L619:
	;
	v2923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2911))))
	if v2923 != 0 {
		goto L621
	} else {
		goto L622
	}
L620:
	;
	v2969 = v2962
	goto L43
L621:
	;
	v2925 = v2923
	goto L623
L622:
	;
	v2925 = int32(256)
	goto L623
L623:
	;
	v2926 = int32(3)
	v2929 = *(*int32)(unsafe.Add(mBase, uint32(v2913+v2925<<(uint(v2926)%32))+4))
	v2932 = v2913 + v2929<<(uint(v2926)%32)
	v2933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2911)+1)))
	if v2933 != 0 {
		goto L624
	} else {
		goto L625
	}
L624:
	;
	v2935 = v2933
	goto L626
L625:
	;
	v2935 = int32(256)
	goto L626
L626:
	;
	v2936 = int32(3)
	v2939 = *(*int32)(unsafe.Add(mBase, uint32(v2932+v2935<<(uint(v2936)%32))+4))
	v2942 = v2932 + v2939<<(uint(v2936)%32)
	v2943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2911)+2)))
	if v2943 != 0 {
		goto L627
	} else {
		goto L628
	}
L627:
	;
	v2945 = v2943
	goto L629
L628:
	;
	v2945 = int32(256)
	goto L629
L629:
	;
	v2946 = int32(3)
	v2949 = *(*int32)(unsafe.Add(mBase, uint32(v2942+v2945<<(uint(v2946)%32))+4))
	v2952 = v2942 + v2949<<(uint(v2946)%32)
	v2953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2911)+3)))
	if v2953 != 0 {
		goto L630
	} else {
		goto L631
	}
L630:
	;
	v2955 = v2953
	goto L632
L631:
	;
	v2955 = int32(256)
	goto L632
L632:
	;
	v2956 = int32(3)
	v2959 = *(*int32)(unsafe.Add(mBase, uint32(v2952+v2955<<(uint(v2956)%32))+4))
	v2962 = v2952 + v2959<<(uint(v2956)%32)
	v2964 = v2911 + int32(4)
	if v2964 != v2714 {
		v2911 = v2964
		v2913 = v2962
		goto L619
	} else {
		goto L633
	}
L633:
	;
	goto L620
L634:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L635:
	;
	v3023 = v3008
	v3025 = v3010
	goto L636
L636:
	;
	v3035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3023))))
	if v3035 != 0 {
		goto L638
	} else {
		goto L639
	}
L637:
	;
	v3081 = v3074
	goto L37
L638:
	;
	v3037 = v3035
	goto L640
L639:
	;
	v3037 = int32(256)
	goto L640
L640:
	;
	v3038 = int32(3)
	v3041 = *(*int32)(unsafe.Add(mBase, uint32(v3025+v3037<<(uint(v3038)%32))+4))
	v3044 = v3025 + v3041<<(uint(v3038)%32)
	v3045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3023)+1)))
	if v3045 != 0 {
		goto L641
	} else {
		goto L642
	}
L641:
	;
	v3047 = v3045
	goto L643
L642:
	;
	v3047 = int32(256)
	goto L643
L643:
	;
	v3048 = int32(3)
	v3051 = *(*int32)(unsafe.Add(mBase, uint32(v3044+v3047<<(uint(v3048)%32))+4))
	v3054 = v3044 + v3051<<(uint(v3048)%32)
	v3055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3023)+2)))
	if v3055 != 0 {
		goto L644
	} else {
		goto L645
	}
L644:
	;
	v3057 = v3055
	goto L646
L645:
	;
	v3057 = int32(256)
	goto L646
L646:
	;
	v3058 = int32(3)
	v3061 = *(*int32)(unsafe.Add(mBase, uint32(v3054+v3057<<(uint(v3058)%32))+4))
	v3064 = v3054 + v3061<<(uint(v3058)%32)
	v3065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3023)+3)))
	if v3065 != 0 {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v3067 = v3065
	goto L649
L648:
	;
	v3067 = int32(256)
	goto L649
L649:
	;
	v3068 = int32(3)
	v3071 = *(*int32)(unsafe.Add(mBase, uint32(v3064+v3067<<(uint(v3068)%32))+4))
	v3074 = v3064 + v3071<<(uint(v3068)%32)
	v3076 = v3023 + int32(4)
	if v3076 != v2010 {
		v3023 = v3076
		v3025 = v3074
		goto L636
	} else {
		goto L650
	}
L650:
	;
	goto L637
L651:
	;
	v3093 = *(*int32)(unsafe.Add(mBase, uint32(v3081)+2048))
	if v3093 != int32(256) {
		v170 = v2010
		v172 = v3081
		v173 = v2006
		goto L35
	} else {
		goto L652
	}
L652:
	;
	goto L36
L653:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L654:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
