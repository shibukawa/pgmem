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
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
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
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v690 int64
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v718 int64
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v962 int64
	_ = v962
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1009 int64
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1345 int32
	_ = v1345
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
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
	var v1430 int32
	_ = v1430
	var v1433 int32
	_ = v1433
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1563 int32
	_ = v1563
	var v1574 int32
	_ = v1574
	var v1582 int32
	_ = v1582
	var v1597 int32
	_ = v1597
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1631 int32
	_ = v1631
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1683 int32
	_ = v1683
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1739 int32
	_ = v1739
	var v1741 int32
	_ = v1741
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1773 int32
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1780 int32
	_ = v1780
	var v1783 int32
	_ = v1783
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1804 int32
	_ = v1804
	var v1807 int64
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1847 int64
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1873 int64
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1899 int64
	_ = v1899
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
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
	var v1922 int32
	_ = v1922
	var v1925 int64
	_ = v1925
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1930 int32
	_ = v1930
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
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
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1983 int32
	_ = v1983
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1989 int32
	_ = v1989
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2001 int32
	_ = v2001
	var v2004 int64
	_ = v2004
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2009 int32
	_ = v2009
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2051 int32
	_ = v2051
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2106 int32
	_ = v2106
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2129 int32
	_ = v2129
	var v2134 int32
	_ = v2134
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2148 int32
	_ = v2148
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2166 int32
	_ = v2166
	var v2169 int32
	_ = v2169
	var v2171 int32
	_ = v2171
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2185 int32
	_ = v2185
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2198 int32
	_ = v2198
	var v2201 int32
	_ = v2201
	var v2204 int32
	_ = v2204
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2217 int32
	_ = v2217
	var v2219 int32
	_ = v2219
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2225 int32
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2230 int32
	_ = v2230
	var v2232 int32
	_ = v2232
	var v2236 int32
	_ = v2236
	var v2239 int32
	_ = v2239
	var v2250 int32
	_ = v2250
	var v2253 int32
	_ = v2253
	var v2256 int32
	_ = v2256
	var v2263 int32
	_ = v2263
	var v2265 int32
	_ = v2265
	var v2270 int32
	_ = v2270
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2297 int32
	_ = v2297
	var v2300 int32
	_ = v2300
	var v2304 int32
	_ = v2304
	var v2307 int32
	_ = v2307
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2319 int32
	_ = v2319
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2329 int32
	_ = v2329
	var v2333 int32
	_ = v2333
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2344 int32
	_ = v2344
	var v2346 int32
	_ = v2346
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2359 int32
	_ = v2359
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2371 int32
	_ = v2371
	var v2374 int32
	_ = v2374
	var v2384 int32
	_ = v2384
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2396 int32
	_ = v2396
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2418 int32
	_ = v2418
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2441 int32
	_ = v2441
	var v2445 int32
	_ = v2445
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2469 int32
	_ = v2469
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2486 int32
	_ = v2486
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2533 int32
	_ = v2533
	var v2552 int32
	_ = v2552
	var v2555 int32
	_ = v2555
	var v2559 int32
	_ = v2559
	var v2563 int32
	_ = v2563
	var v2564 int32
	_ = v2564
	var v2565 int32
	_ = v2565
	var v2567 int32
	_ = v2567
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2580 int32
	_ = v2580
	var v2583 int32
	_ = v2583
	var v2587 int32
	_ = v2587
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2600 int32
	_ = v2600
	var v2603 int32
	_ = v2603
	var v2606 int32
	_ = v2606
	var v2609 int32
	_ = v2609
	var v2612 int32
	_ = v2612
	var v2615 int32
	_ = v2615
	var v2623 int32
	_ = v2623
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2641 int32
	_ = v2641
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2651 int32
	_ = v2651
	var v2654 int32
	_ = v2654
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2660 int32
	_ = v2660
	var v2662 int32
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2672 int32
	_ = v2672
	var v2674 int32
	_ = v2674
	var v2676 int32
	_ = v2676
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2683 int32
	_ = v2683
	var v2687 int32
	_ = v2687
	var v2689 int32
	_ = v2689
	var v2692 int32
	_ = v2692
	var v2699 int32
	_ = v2699
	var v2701 int32
	_ = v2701
	var v2704 int32
	_ = v2704
	var v2707 int32
	_ = v2707
	var v2708 int32
	_ = v2708
	var v2710 int32
	_ = v2710
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2719 int32
	_ = v2719
	var v2721 int32
	_ = v2721
	var v2724 int32
	_ = v2724
	var v2728 int32
	_ = v2728
	var v2729 int32
	_ = v2729
	var v2733 int32
	_ = v2733
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2745 int32
	_ = v2745
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
	var v2754 int32
	_ = v2754
	var v2756 int32
	_ = v2756
	var v2759 int32
	_ = v2759
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2764 int32
	_ = v2764
	var v2773 int32
	_ = v2773
	var v2774 int32
	_ = v2774
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2781 int32
	_ = v2781
	var v2782 int32
	_ = v2782
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2791 int32
	_ = v2791
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2798 int32
	_ = v2798
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2805 int32
	_ = v2805
	var v2806 int32
	_ = v2806
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2819 int32
	_ = v2819
	var v2821 int32
	_ = v2821
	var v2824 int32
	_ = v2824
	var v2825 int32
	_ = v2825
	var v2827 int32
	_ = v2827
	var v2829 int32
	_ = v2829
	var v2830 int32
	_ = v2830
	var v2834 int32
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2836 int32
	_ = v2836
	var v2840 int32
	_ = v2840
	var v2841 int32
	_ = v2841
	var v2844 int32
	_ = v2844
	var v2845 int32
	_ = v2845
	var v2846 int32
	_ = v2846
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2857 int32
	_ = v2857
	var v2862 int32
	_ = v2862
	var v2868 int32
	_ = v2868
	var v2870 int32
	_ = v2870
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2876 int32
	_ = v2876
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2894 int32
	_ = v2894
	var v2897 int32
	_ = v2897
	var v2899 int32
	_ = v2899
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2909 int32
	_ = v2909
	var v2911 int32
	_ = v2911
	var v2917 int32
	_ = v2917
	var v2919 int32
	_ = v2919
	var v2924 int32
	_ = v2924
	var v2926 int32
	_ = v2926
	var v2929 int32
	_ = v2929
	var v2931 int32
	_ = v2931
	var v2934 int32
	_ = v2934
	var v2937 int32
	_ = v2937
	var v2940 int32
	_ = v2940
	var v2948 int32
	_ = v2948
	var v2949 int32
	_ = v2949
	var v2951 int32
	_ = v2951
	var v2952 int32
	_ = v2952
	var v2955 int32
	_ = v2955
	var v2958 int32
	_ = v2958
	var v2960 int32
	_ = v2960
	var v2964 int32
	_ = v2964
	var v2965 int32
	_ = v2965
	var v2968 int32
	_ = v2968
	var v2980 int32
	_ = v2980
	var v2983 int32
	_ = v2983
	var v2993 int32
	_ = v2993
	var v2995 int32
	_ = v2995
	var v2996 int32
	_ = v2996
	var v2999 int32
	_ = v2999
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3005 int32
	_ = v3005
	var v3006 int32
	_ = v3006
	var v3009 int32
	_ = v3009
	var v3012 int32
	_ = v3012
	var v3013 int32
	_ = v3013
	var v3015 int32
	_ = v3015
	var v3016 int32
	_ = v3016
	var v3019 int32
	_ = v3019
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3029 int32
	_ = v3029
	var v3032 int32
	_ = v3032
	var v3034 int32
	_ = v3034
	var v3037 int32
	_ = v3037
	var v3039 int32
	_ = v3039
	var v3052 int32
	_ = v3052
	var v3054 int32
	_ = v3054
	var v3064 int32
	_ = v3064
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3070 int32
	_ = v3070
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3080 int32
	_ = v3080
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3086 int32
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3090 int32
	_ = v3090
	var v3093 int32
	_ = v3093
	var v3094 int32
	_ = v3094
	var v3096 int32
	_ = v3096
	var v3097 int32
	_ = v3097
	var v3100 int32
	_ = v3100
	var v3103 int32
	_ = v3103
	var v3105 int32
	_ = v3105
	var v3110 int32
	_ = v3110
	var v3124 int32
	_ = v3124
	var v3130 int32
	_ = v3130
	var v3141 int32
	_ = v3141
	var v3145 int32
	_ = v3145
	var v3149 int32
	_ = v3149
	var v3151 int32
	_ = v3151
	var v3164 int32
	_ = v3164
	var v3166 int32
	_ = v3166
	var v3176 int32
	_ = v3176
	var v3178 int32
	_ = v3178
	var v3179 int32
	_ = v3179
	var v3182 int32
	_ = v3182
	var v3185 int32
	_ = v3185
	var v3186 int32
	_ = v3186
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3192 int32
	_ = v3192
	var v3195 int32
	_ = v3195
	var v3196 int32
	_ = v3196
	var v3198 int32
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3202 int32
	_ = v3202
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3208 int32
	_ = v3208
	var v3209 int32
	_ = v3209
	var v3212 int32
	_ = v3212
	var v3215 int32
	_ = v3215
	var v3217 int32
	_ = v3217
	var v3222 int32
	_ = v3222
	var v3234 int32
	_ = v3234
	var v3237 int32
	_ = v3237
	var v3239 int32
	_ = v3239
	var v3245 int32
	_ = v3245
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3259 int32
	_ = v3259
	var v3260 int32
	_ = v3260
	var v3266 int32
	_ = v3266
	var v3267 int32
	_ = v3267
	var v3268 int32
	_ = v3268
	var v3269 int32
	_ = v3269
	var v3270 int32
	_ = v3270
	var v3275 int32
	_ = v3275
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
	v3267 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v3268 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v3269 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v3270 = *(*int32)(unsafe.Add(mBase, uint32(v3269)))
	*(*int32)(unsafe.Add(mBase, uint32(v3267))) = v3268 - v3270
	F_scanner_yyerror(m, int32(_a_F_core_yylex_0), v91)
	mBase = m.M
	v3275 = m.ExcPending
	if v3275 != 0 {
		goto L18
	} else {
		goto L703
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
	v3266 = m.ExcPending
	if v3266 != 0 {
		goto L18
	} else {
		goto L702
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
	v3258 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v3259 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(v3259)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v3258))) = v3260
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
	v3257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3245))))
	v111 = v3257
	v112 = v3245
	v114 = v3247
	v115 = v3248
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
	v3237 = *(*int32)(unsafe.Add(mBase, uint32(v3222)+2052))
	v3239 = v2127 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v3239
	v3245 = v3239
	v3247 = v3222 + v3237<<(uint(int32(3))%32)
	v3248 = v2123
	goto L34
L37:
	;
	if v3222 == int32(0) {
		v170 = v2127
		v172 = v3222
		v173 = v2123
		goto L35
	} else {
		goto L700
	}
L38:
	;
	switch v193 - int32(1) {
	case 0, 4, 5, 6:
		goto L22
	case 1:
		goto L134
	case 2:
		goto L133
	case 3:
		goto L132
	case 7:
		goto L130
	case 8, 9:
		goto L129
	case 10:
		goto L127
	case 11:
		goto L125
	case 12:
		goto L124
	case 13:
		goto L123
	case 14:
		goto L122
	case 15:
		goto L121
	case 16:
		goto L120
	case 17, 18, 80:
		goto L119
	case 19:
		goto L118
	case 20:
		goto L117
	case 21:
		goto L116
	case 22:
		goto L115
	case 23:
		goto L114
	case 24, 25, 85:
		goto L113
	case 26:
		goto L112
	case 27:
		goto L111
	case 28:
		goto L110
	case 29:
		goto L109
	case 30:
		goto L108
	case 31:
		goto L106
	case 32:
		goto L105
	case 33:
		goto L104
	case 34:
		goto L103
	case 35:
		goto L102
	case 36:
		goto L101
	case 37:
		goto L99
	case 38:
		goto L98
	case 39:
		goto L97
	case 40:
		goto L96
	case 41:
		goto L95
	case 42:
		goto L94
	case 43:
		goto L92
	case 44:
		goto L91
	case 45:
		goto L90
	case 46:
		goto L89
	case 47:
		goto L88
	case 48:
		goto L87
	case 49:
		goto L86
	case 50:
		goto L64
	case 51:
		goto L85
	case 52:
		goto L84
	case 53:
		goto L83
	case 54:
		goto L82
	case 55:
		goto L81
	case 56:
		goto L80
	case 57:
		goto L79
	case 58:
		goto L78
	case 59:
		goto L77
	case 60:
		goto L76
	case 61:
		goto L75
	case 62:
		goto L74
	case 63:
		goto L73
	case 64:
		goto L72
	case 65:
		goto L71
	case 66, 67, 68, 69:
		goto L21
	case 70:
		goto L70
	case 71:
		goto L69
	case 72:
		goto L67
	case 73:
		goto L66
	case 74:
		goto L68
	case 75:
		goto L128
	case 76:
		goto L131
	case 77, 83:
		goto L93
	case 78:
		goto L126
	case 79, 81, 84:
		goto L107
	case 82:
		goto L100
	default:
		goto L65
	}
L39:
	;
	if base.Ui32(v170-v2084-int32(2)) < base.Ui32(int32(3)) {
		v3222 = v3151
		goto L37
	} else {
		goto L684
	}
L40:
	;
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v3130
	*(*int32)(unsafe.Add(mBase, uint32(v91)+48)) = int32(0)
	v3141 = *(*int32)(unsafe.Add(mBase, uint32(v91)+44))
	v3145 = base.I32_div_s(v3141-int32(1), int32(2))
	v193 = v3145 + int32(75)
	goto L38
L42:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_2))
	mBase = m.M
	v3124 = m.ExcPending
	if v3124 != 0 {
		goto L18
	} else {
		goto L683
	}
L43:
	;
	v3245 = v2855
	v3247 = v3110
	v3248 = v2846
	goto L34
L44:
	;
	if base.Ui32(v170-v2084-int32(2)) < base.Ui32(int32(3)) {
		v3110 = v3039
		goto L43
	} else {
		goto L667
	}
L45:
	;
	if base.Ui32(v2926-int32(1)) < base.Ui32(int32(3)) {
		v170 = v2917
		v172 = v2968
		v173 = v2909
		goto L35
	} else {
		goto L651
	}
L46:
	;
	v2965 = v2909
	v2968 = v2924
	goto L45
L47:
	;
	v3037 = v2846
	v3039 = v2862
	goto L44
L48:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_3))
	mBase = m.M
	v2964 = m.ExcPending
	if v2964 != 0 {
		goto L18
	} else {
		goto L650
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v2917 = v2904 + v2911
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v2917
	v2919 = *(*int32)(unsafe.Add(mBase, uint32(v91)+44))
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v2919<<(uint(int32(2))%32))+uint32(_c_F_core_yylex[2])))
	if base.Ui32(v2917) <= base.Ui32(v2909) {
		v170 = v2917
		v172 = v2924
		v173 = v2909
		goto L35
	} else {
		goto L642
	}
L51:
	;
	v2648 = *(*int32)(unsafe.Add(mBase, uint32(v2647)))
	*(*int32)(unsafe.Add(mBase, uint32(v2648)+16)) = v2641
	v2651 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	if v2651 != 0 {
		v2773 = int32(0)
		goto L597
	} else {
		goto L598
	}
L52:
	;
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2641 = v2623
	v2647 = v2629 + v2630<<(uint(int32(2))%32)
	goto L51
L53:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_4))
	mBase = m.M
	v2615 = m.ExcPending
	if v2615 != 0 {
		goto L18
	} else {
		goto L596
	}
L54:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_5))
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L18
	} else {
		goto L595
	}
L55:
	;
	v3149 = v2123
	v3151 = v2134
	goto L40
L56:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_6), v91)
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		goto L18
	} else {
		goto L594
	}
L57:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_7), v91)
	mBase = m.M
	v2606 = m.ExcPending
	if v2606 != 0 {
		goto L18
	} else {
		goto L593
	}
L58:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_7), v91)
	mBase = m.M
	v2603 = m.ExcPending
	if v2603 != 0 {
		goto L18
	} else {
		goto L592
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2580 = m.ExcPending
	if v2580 != 0 {
		goto L18
	} else {
		goto L586
	}
L60:
	;
	v2573 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v2573)+52)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(23)
	goto L33
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2552 = m.ExcPending
	if v2552 != 0 {
		goto L18
	} else {
		goto L580
	}
L62:
	;
	m.G0 = v97 + int32(16)
	return v2533
L63:
	;
	v2533 = int32(274)
	goto L62
L64:
	;
	v2512 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(v2514)))
	*(*int32)(unsafe.Add(mBase, uint32(v2512))) = v2513 - v2515
	goto L63
L65:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_8))
	mBase = m.M
	v2511 = m.ExcPending
	if v2511 != 0 {
		goto L18
	} else {
		goto L579
	}
L66:
	;
	v2084 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v2085)
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2091 = v2087 + v2088<<(uint(int32(2))%32)
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v2091)))
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v2092)+44))
	if v2093 == int32(0) {
		goto L494
	} else {
		goto L495
	}
L67:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_9))
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		goto L18
	} else {
		goto L493
	}
L68:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v2076)))
	*(*int32)(unsafe.Add(mBase, uint32(v2074))) = v2075 - v2077
	v2533 = int32(0)
	goto L62
L69:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2068)))
	*(*int32)(unsafe.Add(mBase, uint32(v2066))) = v2067 - v2069
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2073 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2072))))
	v2533 = v2073
	goto L62
L70:
	;
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v2032)))
	*(*int32)(unsafe.Add(mBase, uint32(v2030))) = v2031 - v2033
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v2037)+8))
	v2039 = F_ScanKeywordLookup(m, v2036, v2038)
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L18
	} else {
		goto L488
	}
L71:
	;
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v2020)))
	*(*int32)(unsafe.Add(mBase, uint32(v2018))) = v2019 - v2021
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2025 = F_pstrdup(m, v2024)
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L18
	} else {
		goto L487
	}
L72:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v1979)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v1983 = v1978 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v1983
	v1985 = v1983 + v173
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1985
	v1987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1985))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v1987)
	v1989 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1985))) = uint8(v1989)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1985
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v1994)))
	*(*int32)(unsafe.Add(mBase, uint32(v1992))) = v1993 - v1995
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2001 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v2001
	v2004 = *(*int64)(unsafe.Add(mBase, _c_F_core_yylex[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v2004
	v2006 = F_pg_strtoint32_safe(m, v1999, v97)
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L18
	} else {
		goto L482
	}
L73:
	;
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v1968)))
	*(*int32)(unsafe.Add(mBase, uint32(v1966))) = v1967 - v1969
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1973 = F_pstrdup(m, v1972)
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L18
	} else {
		goto L481
	}
L74:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(v1959)))
	*(*int32)(unsafe.Add(mBase, uint32(v1957))) = v1958 - v1960
	F_scanner_yyerror(m, int32(_a_F_core_yylex_10), v91)
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L18
	} else {
		goto L480
	}
L75:
	;
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1950)))
	*(*int32)(unsafe.Add(mBase, uint32(v1948))) = v1949 - v1951
	F_scanner_yyerror(m, int32(_a_F_core_yylex_11), v91)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L18
	} else {
		goto L479
	}
L76:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1941)))
	*(*int32)(unsafe.Add(mBase, uint32(v1939))) = v1940 - v1942
	F_scanner_yyerror(m, int32(_a_F_core_yylex_12), v91)
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L18
	} else {
		goto L478
	}
L77:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v1915)))
	*(*int32)(unsafe.Add(mBase, uint32(v1913))) = v1914 - v1916
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1922 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v1922
	v1925 = *(*int64)(unsafe.Add(mBase, _c_F_core_yylex[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v1925
	v1927 = F_pg_strtoint32_safe(m, v1920, v97)
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L18
	} else {
		goto L473
	}
L78:
	;
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v1889)))
	*(*int32)(unsafe.Add(mBase, uint32(v1887))) = v1888 - v1890
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1896 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v1896
	v1899 = *(*int64)(unsafe.Add(mBase, _c_F_core_yylex[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v1899
	v1901 = F_pg_strtoint32_safe(m, v1894, v97)
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L18
	} else {
		goto L468
	}
L79:
	;
	v1861 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(v1863)))
	*(*int32)(unsafe.Add(mBase, uint32(v1861))) = v1862 - v1864
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1870 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v1870
	v1873 = *(*int64)(unsafe.Add(mBase, _c_F_core_yylex[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v1873
	v1875 = F_pg_strtoint32_safe(m, v1868, v97)
	mBase = m.M
	v1876 = m.ExcPending
	if v1876 != 0 {
		goto L18
	} else {
		goto L463
	}
L80:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1837)))
	*(*int32)(unsafe.Add(mBase, uint32(v1835))) = v1836 - v1838
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1844 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v1844
	v1847 = *(*int64)(unsafe.Add(mBase, _c_F_core_yylex[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v1847
	v1849 = F_pg_strtoint32_safe(m, v1842, v97)
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L18
	} else {
		goto L458
	}
L81:
	;
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v1828)))
	*(*int32)(unsafe.Add(mBase, uint32(v1826))) = v1827 - v1829
	F_scanner_yyerror(m, int32(_a_F_core_yylex_13), v91)
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L18
	} else {
		goto L457
	}
L82:
	;
	v1804 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v1804
	v1807 = *(*int64)(unsafe.Add(mBase, _c_F_core_yylex[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v1807
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1811)))
	*(*int32)(unsafe.Add(mBase, uint32(v1809))) = v1810 - v1812
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1818 = F_pg_strtoint32_safe(m, v1815+int32(1), v97)
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L18
	} else {
		goto L455
	}
L83:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1546 = F_strstr(m, v1544, int32(_a_F_core_yylex_14))
	mBase = m.M
	v1548 = F_strstr(m, v1544, int32(_a_F_core_yylex_15))
	mBase = m.M
	if base.B2i32(v1548 != int32(0))&base.B2i32(base.Ui32(v1548) < base.Ui32(v1546)) != 0 {
		goto L390
	} else {
		goto L391
	}
L84:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1538)))
	*(*int32)(unsafe.Add(mBase, uint32(v1536))) = v1537 - v1539
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1543 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1542))))
	v2533 = v1543
	goto L62
L85:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1532)))
	*(*int32)(unsafe.Add(mBase, uint32(v1530))) = v1531 - v1533
	goto L63
L86:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1525)))
	*(*int32)(unsafe.Add(mBase, uint32(v1523))) = v1524 - v1526
	v2533 = int32(273)
	goto L62
L87:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1518)))
	*(*int32)(unsafe.Add(mBase, uint32(v1516))) = v1517 - v1519
	v2533 = int32(272)
	goto L62
L88:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1511)))
	*(*int32)(unsafe.Add(mBase, uint32(v1509))) = v1510 - v1512
	v2533 = int32(271)
	goto L62
L89:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1504)))
	*(*int32)(unsafe.Add(mBase, uint32(v1502))) = v1503 - v1505
	v2533 = int32(270)
	goto L62
L90:
	;
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1497)))
	*(*int32)(unsafe.Add(mBase, uint32(v1495))) = v1496 - v1498
	v2533 = int32(269)
	goto L62
L91:
	;
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1490)))
	*(*int32)(unsafe.Add(mBase, uint32(v1488))) = v1489 - v1491
	v2533 = int32(268)
	goto L62
L92:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1463)))
	*(*int32)(unsafe.Add(mBase, uint32(v1461))) = v1462 - v1464
	v1467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v1467)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v1470 = int32(1)
	v1471 = v173 + v1470
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v1470
	v1475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v1475)
	v1477 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)) = uint8(v1477)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1471
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1483 = F_downcase_truncate_identifier(m, v1480, v1481, v1470)
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L18
	} else {
		goto L388
	}
L93:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_16), v91)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L18
	} else {
		goto L387
	}
L94:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+28))
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1427 = v1425 + v1426
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+32))
	if v1428 <= v1427 {
		goto L377
	} else {
		goto L378
	}
L95:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+28))
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+32))
	if v1397 <= v1394+int32(1) {
		goto L373
	} else {
		goto L374
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(1)
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+28))
	if v1377 == int32(0) {
		goto L57
	} else {
		goto L368
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(1)
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1350)+28))
	if v1351 == int32(0) {
		goto L58
	} else {
		goto L359
	}
L98:
	;
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1339)))
	*(*int32)(unsafe.Add(mBase, uint32(v1337))) = v1338 - v1340
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(19)
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1345)+28)) = int32(0)
	goto L22
L99:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1328)))
	*(*int32)(unsafe.Add(mBase, uint32(v1326))) = v1327 - v1329
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(7)
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1334)+28)) = int32(0)
	goto L22
L100:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_17), v91)
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L18
	} else {
		goto L358
	}
L101:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1293))))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+28))
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+32))
	if v1299 <= v1296+int32(1) {
		goto L354
	} else {
		goto L355
	}
L102:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1259)+28))
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1262 = v1260 + v1261
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1259)+32))
	if v1263 <= v1262 {
		goto L344
	} else {
		goto L345
	}
L103:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+28))
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1227 = v1225 + v1226
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+32))
	if v1228 <= v1227 {
		goto L334
	} else {
		goto L335
	}
L104:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+44))
	v1125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1120))))
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122))))
	if base.B2i32(v1125 == int32(0))|base.B2i32(v1125 != v1128) != 0 {
		v1146 = v1125
		v1147 = v1128
		goto L310
	} else {
		goto L311
	}
L105:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1101)))
	*(*int32)(unsafe.Add(mBase, uint32(v1099))) = v1100 - v1102
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v1105)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v1108 = int32(1)
	v1109 = v173 + v1108
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1109
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v1108
	v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v1113)
	v1115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)) = uint8(v1115)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1109
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1119 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1118))))
	v2533 = v1119
	goto L62
L106:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1085)))
	*(*int32)(unsafe.Add(mBase, uint32(v1083))) = v1084 - v1086
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1090 = F_pstrdup(m, v1089)
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L18
	} else {
		goto L308
	}
L107:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_18), v91)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L18
	} else {
		goto L307
	}
L108:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050))))
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+28))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+32))
	if v1056 <= v1053+int32(1) {
		goto L303
	} else {
		goto L304
	}
L109:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1009 = F_strtox_2(m, v1003+int32(2), int32(0), int32(16), int64(4294967295))
	mBase = m.M
	v1010 = base.I32_wrap_i64(v1009)
	goto L293
L110:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v962 = F_strtox_2(m, v956+int32(1), int32(0), int32(8), int64(4294967295))
	mBase = m.M
	v963 = base.I32_wrap_i64(v962)
	goto L283
L111:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v789)+1)))
	if v790 != int32(39) {
		goto L232
	} else {
		goto L233
	}
L112:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v761)))
	*(*int32)(unsafe.Add(mBase, uint32(v759))) = v760 - v762
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L18
	} else {
		goto L226
	}
L113:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	*(*int32)(unsafe.Add(mBase, uint32(v753))) = v754 - v756
	goto L24
L114:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v718 = F_strtox_2(m, v712+int32(2), int32(0), int32(16), int64(4294967295))
	mBase = m.M
	v719 = base.I32_wrap_i64(v718)
	goto L223
L115:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v690 = F_strtox_2(m, v684+int32(2), int32(0), int32(16), int64(4294967295))
	mBase = m.M
	v691 = base.I32_wrap_i64(v690)
	goto L218
L116:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v650)+28))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v653 = v651 + v652
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v650)+32))
	if v654 <= v653 {
		goto L208
	} else {
		goto L209
	}
L117:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v615)+28))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v618 = v616 + v617
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v615)+32))
	if v619 <= v618 {
		goto L198
	} else {
		goto L199
	}
L118:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v584)+28))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v584)+32))
	if v588 <= v585+int32(1) {
		goto L194
	} else {
		goto L195
	}
L119:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v499)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v173
	v503 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v503
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v505)
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v503)
	v509 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v173
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)+36))
	switch v513 - v509 {
	case 0:
		goto L172
	default:
		goto L168
	case 3:
		goto L171
	case 4, 6:
		goto L170
	case 9:
		goto L169
	}
L120:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+36))
	v494 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = v493<<(uint(v494)%32) | v494
	goto L22
L121:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v91)+44))
	v488 = base.I32_div_s(v484-int32(1), int32(2))
	*(*int32)(unsafe.Add(mBase, uint32(v483)+36)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(13)
	goto L22
L122:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	*(*int32)(unsafe.Add(mBase, uint32(v469))) = v470 - v472
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+21)))
	if v476 == int32(0) {
		goto L61
	} else {
		goto L167
	}
L123:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v453 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v452)+56)) = uint8(v453)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+57)) = uint8(v453)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	*(*int32)(unsafe.Add(mBase, uint32(v458))) = v459 - v461
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(15)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v466)+28)) = v453
	goto L22
L124:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v433 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+56)) = uint8(v433)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v436 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v435)+57)) = uint8(v436)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	*(*int32)(unsafe.Add(mBase, uint32(v438))) = v439 - v441
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+21)))
	if v447 != 0 {
		goto L164
	} else {
		goto L165
	}
L125:
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
		goto L159
	}
L126:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_20), v91)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L18
	} else {
		goto L158
	}
L127:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = v342 - v344
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(9)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+28)) = int32(0)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+28))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v352)+32))
	if v356 <= v353+int32(1) {
		goto L154
	} else {
		goto L155
	}
L128:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_21), v91)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L18
	} else {
		goto L153
	}
L129:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+28))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v307 = v305 + v306
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v304)+32))
	if v308 <= v307 {
		goto L143
	} else {
		goto L144
	}
L130:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v263 - v265
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(3)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v270)+28)) = int32(0)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+28))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v273)+32))
	if v277 <= v274+int32(1) {
		goto L139
	} else {
		goto L140
	}
L131:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_22), v91)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L18
	} else {
		goto L138
	}
L132:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+40))
	if v251 <= int32(0) {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v232)+40)) = v233 + int32(1)
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
L134:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v209 - v211
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v215 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v214)+40)) = v215
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
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(1)
	goto L22
L136:
	;
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+40)) = v251 - int32(1)
	goto L22
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+32)) = v277 << (uint(int32(1)) % 32)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+24))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v282)+32))
	v285 = F_repalloc(m, v283, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L18
	} else {
		goto L142
	}
L140:
	;
	v292 = v273
	v293 = v274
	goto L141
L141:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v292)+24))
	v296 = int32(98)
	*(*uint8)(unsafe.Add(mBase, uint32(v293+v294))) = uint8(v296)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+28)) = v299 + int32(1)
	goto L22
L142:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v287)+24)) = v285
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+28))
	v292 = v289
	v293 = v290
	goto L141
L143:
	;
	v310 = int32(1)
	v313 = v307 + v310
	if v313&v307 != 0 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v329 = v304
	v330 = v305
	goto L145
L145:
	;
	if v306 != 0 {
		goto L150
	} else {
		goto L151
	}
L146:
	;
	v318 = v310 << (uint(int32(32)-base.I32_clz(v313)) % 32)
	goto L148
L147:
	;
	v318 = v313
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v304)+32)) = v318
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+24))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v320)+32))
	v323 = F_repalloc(m, v321, v322)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v325)+24)) = v323
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+28))
	v329 = v327
	v330 = v328
	goto L145
L150:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v329)+24))
	base.MemoryCopy(m, v331+v330, v303, v306)
	goto L152
L151:
	;
	goto L152
L152:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v334)+28)) = v335 + v306
	goto L22
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v352)+32)) = v356 << (uint(int32(1)) % 32)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+24))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v361)+32))
	v364 = F_repalloc(m, v362, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L18
	} else {
		goto L157
	}
L155:
	;
	v371 = v352
	v372 = v353
	goto L156
L156:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v371)+24))
	v375 = int32(120)
	*(*uint8)(unsafe.Add(mBase, uint32(v372+v373))) = uint8(v375)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v377)+28)) = v378 + int32(1)
	goto L22
L157:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+24)) = v364
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+28))
	v371 = v368
	v372 = v369
	goto L156
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	if int32(0) <= v407 {
		goto L160
	} else {
		goto L161
	}
L160:
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
	v2533 = v425
	goto L62
L161:
	;
	goto L162
L162:
	;
	v427 = F_pstrdup(m, int32(_a_F_core_yylex_23))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L18
	} else {
		goto L163
	}
L163:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = v427
	v2533 = int32(258)
	goto L62
L164:
	;
	v448 = int32(11)
	goto L166
L165:
	;
	v448 = int32(15)
	goto L166
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v446)+28)) = int32(0)
	goto L22
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(21)
	*(*int32)(unsafe.Add(mBase, uint32(v475)+28)) = int32(0)
	goto L22
L168:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_24), v91)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L18
	} else {
		goto L193
	}
L169:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v512)+28))
	v570 = F_palloc(m, v567+int32(1))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L18
	} else {
		goto L189
	}
L170:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+57)))
	if v544 == int32(1) {
		goto L181
	} else {
		goto L182
	}
L171:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v512)+28))
	v533 = F_palloc(m, v530+int32(1))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L18
	} else {
		goto L177
	}
L172:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v512)+28))
	v519 = F_palloc(m, v516+int32(1))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L18
	} else {
		goto L173
	}
L173:
	;
	if v516 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)+24))
	base.MemoryCopy(m, v519, v522, v516)
	goto L176
L175:
	;
	goto L176
L176:
	;
	v525 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v516+v519))) = uint8(v525)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v527))) = v519
	v2533 = int32(263)
	goto L62
L177:
	;
	if v530 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+24))
	base.MemoryCopy(m, v533, v536, v530)
	goto L180
L179:
	;
	goto L180
L180:
	;
	v539 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v530+v533))) = uint8(v539)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v541))) = v533
	v2533 = int32(264)
	goto L62
L181:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v512)+24))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v512)+28))
	F_pg_verifymbstr(m, v547, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L18
	} else {
		goto L184
	}
L182:
	;
	v552 = v512
	goto L183
L183:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)+28))
	v556 = F_palloc(m, v553+int32(1))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L18
	} else {
		goto L185
	}
L184:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v552 = v551
	goto L183
L185:
	;
	if v553 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)+24))
	base.MemoryCopy(m, v556, v559, v553)
	goto L188
L187:
	;
	goto L188
L188:
	;
	v562 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v553+v556))) = uint8(v562)
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v564))) = v556
	v2533 = int32(261)
	goto L62
L189:
	;
	if v567 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+24))
	base.MemoryCopy(m, v570, v573, v567)
	goto L192
L191:
	;
	goto L192
L192:
	;
	v576 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v567+v570))) = uint8(v576)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v578))) = v570
	v2533 = int32(262)
	goto L62
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v584)+32)) = v588 << (uint(int32(1)) % 32)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v593)+24))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v593)+32))
	v596 = F_repalloc(m, v594, v595)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L18
	} else {
		goto L197
	}
L195:
	;
	v603 = v584
	v604 = v585
	goto L196
L196:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v603)+24))
	v607 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v604+v605))) = uint8(v607)
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v609)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v609)+28)) = v610 + int32(1)
	goto L22
L197:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v598)+24)) = v596
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v600)+28))
	v603 = v600
	v604 = v601
	goto L196
L198:
	;
	v621 = int32(1)
	v624 = v618 + v621
	if v624&v618 != 0 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v640 = v615
	v641 = v616
	goto L200
L200:
	;
	if v617 != 0 {
		goto L205
	} else {
		goto L206
	}
L201:
	;
	v629 = v621 << (uint(int32(32)-base.I32_clz(v624)) % 32)
	goto L203
L202:
	;
	v629 = v624
	goto L203
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615)+32)) = v629
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)+24))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v631)+32))
	v634 = F_repalloc(m, v632, v633)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L18
	} else {
		goto L204
	}
L204:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v636)+24)) = v634
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)+28))
	v640 = v638
	v641 = v639
	goto L200
L205:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v640)+24))
	base.MemoryCopy(m, v642+v641, v614, v617)
	goto L207
L206:
	;
	goto L207
L207:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v645)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v645)+28)) = v646 + v617
	goto L22
L208:
	;
	v656 = int32(1)
	v659 = v653 + v656
	if v659&v653 != 0 {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	v675 = v650
	v676 = v651
	goto L210
L210:
	;
	if v652 != 0 {
		goto L215
	} else {
		goto L216
	}
L211:
	;
	v664 = v656 << (uint(int32(32)-base.I32_clz(v659)) % 32)
	goto L213
L212:
	;
	v664 = v659
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v650)+32)) = v664
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v666)+24))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v666)+32))
	v669 = F_repalloc(m, v667, v668)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L18
	} else {
		goto L214
	}
L214:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v671)+24)) = v669
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v673)+28))
	v675 = v673
	v676 = v674
	goto L210
L215:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v675)+24))
	base.MemoryCopy(m, v677+v676, v649, v652)
	goto L217
L216:
	;
	goto L217
L217:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v680)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v680)+28)) = v681 + v652
	goto L22
L218:
	;
	F_check_escape_warning(m, v91)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L18
	} else {
		goto L219
	}
L219:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)))
	*(*int32)(unsafe.Add(mBase, uint32(v694)+48)) = v696
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v700)))
	*(*int32)(unsafe.Add(mBase, uint32(v698))) = v699 - v701
	v705 = v691 & int32(-1024)
	if v705 == int32(_a_F_core_yylex_25) {
		goto L60
	} else {
		goto L220
	}
L220:
	;
	if v705 == int32(_a_F_core_yylex_26) {
		goto L24
	} else {
		goto L221
	}
L221:
	;
	F_addunicode(m, v691, v91)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L18
	} else {
		goto L222
	}
L222:
	;
	goto L33
L223:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	*(*int32)(unsafe.Add(mBase, uint32(v720)+48)) = v722
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	*(*int32)(unsafe.Add(mBase, uint32(v724))) = v725 - v727
	if v719&int32(-1024) != int32(_a_F_core_yylex_26) {
		goto L24
	} else {
		goto L224
	}
L224:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v734)+52))
	F_addunicode(m, v735<<(uint(int32(10))%32)&int32(_a_F_core_yylex_27)|v719&int32(1023)+int32(_a_F_core_yylex_28), v91)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L18
	} else {
		goto L225
	}
L225:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v748)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v747))) = v749
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(15)
	goto L22
L226:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L18
	} else {
		goto L227
	}
L227:
	;
	F_errmsg(m, int32(_a_F_core_yylex_29), int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L18
	} else {
		goto L228
	}
L228:
	;
	F_errhint(m, int32(_a_F_core_yylex_30), int32(0))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L18
	} else {
		goto L229
	}
L229:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	F_scanner_errposition(m, v781, v91)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L18
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_core_yylex_31), int32(704), int32(_a_F_core_yylex_32))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L18
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v805)+1)))
	if v806 != int32(92) {
		goto L241
	} else {
		goto L242
	}
L233:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v793)+16))
	switch v794 {
	case 0:
		goto L59
	default:
		goto L232
	case 2:
		goto L234
	}
L234:
	;
	v796 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[5]))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v796)+4))
	goto L235
L235:
	;
	if v797 < int32(35) {
		goto L232
	} else {
		goto L236
	}
L236:
	;
	v801 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[5]))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v801)+4))
	goto L237
L237:
	;
	if v802 <= int32(41) {
		goto L59
	} else {
		goto L238
	}
L238:
	;
	goto L232
L239:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904)+1)))
	v909 = base.I32_rotl(v905-int32(98), int32(31))
	if base.B2i32(base.Ui32(v909) <= base.Ui32(int32(10)))&(int32(base.Ui32(int32(1861))>>(uint(v909)%32))&int32(1)) == int32(0) {
		goto L275
	} else {
		goto L276
	}
L240:
	;
	F_check_escape_warning(m, v91)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L18
	} else {
		goto L273
	}
L241:
	;
	if v806 != int32(39) {
		goto L240
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856)+56)))
	if v857 != int32(1) {
		goto L259
	} else {
		goto L260
	}
L244:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v811)+56)))
	if v812 != int32(1) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v854 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v853)+56)) = uint8(v854)
	goto L239
L246:
	;
	v815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v811)+20)))
	if v815 != int32(1) {
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v820 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L18
	} else {
		goto L248
	}
L248:
	;
	if v820 == int32(0) {
		goto L245
	} else {
		goto L249
	}
L249:
	;
	F_errcode(m, int32(100794498))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L18
	} else {
		goto L250
	}
L250:
	;
	F_errmsg(m, int32(_a_F_core_yylex_33), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L18
	} else {
		goto L251
	}
L251:
	;
	F_errhint(m, int32(_a_F_core_yylex_34), int32(0))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L18
	} else {
		goto L252
	}
L252:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v835)))
	if int32(0) <= v836 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v839)))
	v841 = F_pg_mbstrlen_with_len(m, v840, v836)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L18
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	F_errfinish(m, int32(_a_F_core_yylex_31), int32(1433), int32(_a_F_core_yylex_35))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L18
	} else {
		goto L258
	}
L256:
	;
	v845 = F_errposition(m, v841+int32(1))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L18
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	goto L245
L259:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v899 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v898)+56)) = uint8(v899)
	goto L239
L260:
	;
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856)+20)))
	if v860 != int32(1) {
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v865 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L18
	} else {
		goto L262
	}
L262:
	;
	if v865 == int32(0) {
		goto L259
	} else {
		goto L263
	}
L263:
	;
	F_errcode(m, int32(100794498))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L18
	} else {
		goto L264
	}
L264:
	;
	F_errmsg(m, int32(_a_F_core_yylex_36), int32(0))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L18
	} else {
		goto L265
	}
L265:
	;
	F_errhint(m, int32(_a_F_core_yylex_37), int32(0))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L18
	} else {
		goto L266
	}
L266:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v880)))
	if int32(0) <= v881 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v884)))
	v886 = F_pg_mbstrlen_with_len(m, v885, v881)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L18
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	F_errfinish(m, int32(_a_F_core_yylex_31), int32(1443), int32(_a_F_core_yylex_35))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L18
	} else {
		goto L272
	}
L270:
	;
	v890 = F_errposition(m, v886+int32(1))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L18
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	goto L259
L273:
	;
	goto L239
L274:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v927)+28))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v927)+32))
	if v931 <= v928+int32(1) {
		goto L279
	} else {
		goto L280
	}
L275:
	;
	v919 = base.I32_extend8_s(v905)
	if int32(0) < v919 {
		v926 = v919
		goto L274
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909)+uint32(_c_F_core_yylex[6]))))
	v926 = v925
	goto L274
L278:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v923 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v922)+57)) = uint8(v923)
	v926 = v919
	goto L274
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v927)+32)) = v931 << (uint(int32(1)) % 32)
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v936)+24))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v936)+32))
	v939 = F_repalloc(m, v937, v938)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L18
	} else {
		goto L282
	}
L280:
	;
	v946 = v927
	v947 = v928
	goto L281
L281:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v946)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v947+v948))) = uint8(v926)
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v951)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v951)+28)) = v952 + int32(1)
	goto L22
L282:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v941)+24)) = v939
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v943)+28))
	v946 = v943
	v947 = v944
	goto L281
L283:
	;
	F_check_escape_warning(m, v91)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L18
	} else {
		goto L284
	}
L284:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v966)+28))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v966)+32))
	if v970 <= v967+int32(1) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v966)+32)) = v970 << (uint(int32(1)) % 32)
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v975)+24))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v975)+32))
	v978 = F_repalloc(m, v976, v977)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L18
	} else {
		goto L288
	}
L286:
	;
	v984 = v966
	v985 = v967
	goto L287
L287:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v984)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v985+v986))) = uint8(v963)
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v989)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+28)) = v990 + int32(1)
	if v963&int32(128) != 0 {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v980)+24)) = v978
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v982)+28))
	v984 = v982
	v985 = v983
	goto L287
L289:
	;
	v999 = int32(0)
	goto L291
L290:
	;
	v999 = v963 & int32(255)
	goto L291
L291:
	;
	if v999 != 0 {
		goto L22
	} else {
		goto L292
	}
L292:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1001 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1000)+57)) = uint8(v1001)
	goto L22
L293:
	;
	F_check_escape_warning(m, v91)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L18
	} else {
		goto L294
	}
L294:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1013)+28))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1013)+32))
	if v1017 <= v1014+int32(1) {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1013)+32)) = v1017 << (uint(int32(1)) % 32)
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+24))
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+32))
	v1025 = F_repalloc(m, v1023, v1024)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L18
	} else {
		goto L298
	}
L296:
	;
	v1031 = v1013
	v1032 = v1014
	goto L297
L297:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1031)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v1032+v1033))) = uint8(v1010)
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1036)+28)) = v1037 + int32(1)
	if v1010&int32(128) != 0 {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1027)+24)) = v1025
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+28))
	v1031 = v1029
	v1032 = v1030
	goto L297
L299:
	;
	v1046 = int32(0)
	goto L301
L300:
	;
	v1046 = v1010 & int32(255)
	goto L301
L301:
	;
	if v1046 != 0 {
		goto L22
	} else {
		goto L302
	}
L302:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1048 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1047)+57)) = uint8(v1048)
	goto L22
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1052)+32)) = v1056 << (uint(int32(1)) % 32)
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+24))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+32))
	v1064 = F_repalloc(m, v1062, v1063)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L18
	} else {
		goto L306
	}
L304:
	;
	v1070 = v1052
	v1071 = v1053
	goto L305
L305:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v1071+v1072))) = uint8(v1051)
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+28)) = v1076 + int32(1)
	goto L22
L306:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1066)+24)) = v1064
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1068)+28))
	v1070 = v1068
	v1071 = v1069
	goto L305
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L308:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1092)+44)) = v1090
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = int32(17)
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1096)+28)) = int32(0)
	goto L22
L309:
	;
	if v1146-v1147 == int32(0) {
		goto L316
	} else {
		goto L317
	}
L310:
	;
	goto L309
L311:
	;
	v1131 = v1120
	v1132 = v1122
	goto L312
L312:
	;
	v1135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1132)+1)))
	v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131)+1)))
	if v1136 == int32(0) {
		v1146 = v1136
		v1147 = v1135
		goto L310
	} else {
		goto L314
	}
L313:
	;
	v1146 = v1136
	v1147 = v1135
	goto L310
L314:
	;
	v1139 = int32(1)
	if v1136 == v1135 {
		v1131 = v1131 + v1139
		v1132 = v1132 + v1139
		goto L312
	} else {
		goto L315
	}
L315:
	;
	goto L313
L316:
	;
	F_pfree(m, v1122)
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L18
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+32))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+28))
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1177 = v1175 - int32(1)
	if v1173 <= v1174+v1177 {
		goto L324
	} else {
		goto L325
	}
L319:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1153)+44)) = int32(0)
	v1156 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = v1156
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+28))
	v1162 = F_palloc(m, v1159+v1156)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L18
	} else {
		goto L320
	}
L320:
	;
	if v1159 != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1164)+24))
	base.MemoryCopy(m, v1162, v1165, v1159)
	goto L323
L322:
	;
	goto L323
L323:
	;
	v1168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1159+v1162))) = uint8(v1168)
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1170))) = v1162
	v2533 = int32(261)
	goto L62
L324:
	;
	v1180 = int32(1)
	v1182 = v1174 + v1175
	if v1182&(v1182-v1180) != 0 {
		goto L327
	} else {
		goto L328
	}
L325:
	;
	v1200 = v1121
	v1201 = v1174
	goto L326
L326:
	;
	if v1177 != 0 {
		goto L331
	} else {
		goto L332
	}
L327:
	;
	v1189 = v1180 << (uint(int32(32)-base.I32_clz(v1182)) % 32)
	goto L329
L328:
	;
	v1189 = v1182
	goto L329
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1121)+32)) = v1189
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+24))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+32))
	v1194 = F_repalloc(m, v1192, v1193)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L18
	} else {
		goto L330
	}
L330:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1196)+24)) = v1194
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+28))
	v1200 = v1198
	v1201 = v1199
	goto L326
L331:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+24))
	base.MemoryCopy(m, v1202+v1201, v1120, v1177)
	goto L333
L332:
	;
	goto L333
L333:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1205)+28)) = v1206 + v1177
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v1210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v1210)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v1214 = v1209 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v1214
	v1216 = v1214 + v173
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1216
	v1218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1216))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v1218)
	v1220 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1216))) = uint8(v1220)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1216
	goto L22
L334:
	;
	v1230 = int32(1)
	v1233 = v1227 + v1230
	if v1233&v1227 != 0 {
		goto L337
	} else {
		goto L338
	}
L335:
	;
	v1249 = v1224
	v1250 = v1225
	goto L336
L336:
	;
	if v1226 != 0 {
		goto L341
	} else {
		goto L342
	}
L337:
	;
	v1238 = v1230 << (uint(int32(32)-base.I32_clz(v1233)) % 32)
	goto L339
L338:
	;
	v1238 = v1233
	goto L339
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1224)+32)) = v1238
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+24))
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+32))
	v1243 = F_repalloc(m, v1241, v1242)
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L18
	} else {
		goto L340
	}
L340:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1245)+24)) = v1243
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+28))
	v1249 = v1247
	v1250 = v1248
	goto L336
L341:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1249)+24))
	base.MemoryCopy(m, v1251+v1250, v1223, v1226)
	goto L343
L342:
	;
	goto L343
L343:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1254)+28)) = v1255 + v1226
	goto L22
L344:
	;
	v1265 = int32(1)
	v1268 = v1262 + v1265
	if v1268&v1262 != 0 {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	v1284 = v1259
	v1285 = v1260
	goto L346
L346:
	;
	if v1261 != 0 {
		goto L351
	} else {
		goto L352
	}
L347:
	;
	v1273 = v1265 << (uint(int32(32)-base.I32_clz(v1268)) % 32)
	goto L349
L348:
	;
	v1273 = v1268
	goto L349
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1259)+32)) = v1273
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+24))
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+32))
	v1278 = F_repalloc(m, v1276, v1277)
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L18
	} else {
		goto L350
	}
L350:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1280)+24)) = v1278
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1282)+28))
	v1284 = v1282
	v1285 = v1283
	goto L346
L351:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+24))
	base.MemoryCopy(m, v1286+v1285, v1258, v1261)
	goto L353
L352:
	;
	goto L353
L353:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1289)+28)) = v1290 + v1261
	goto L22
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+32)) = v1299 << (uint(int32(1)) % 32)
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+24))
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+32))
	v1307 = F_repalloc(m, v1305, v1306)
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L18
	} else {
		goto L357
	}
L355:
	;
	v1313 = v1295
	v1314 = v1296
	goto L356
L356:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v1314+v1315))) = uint8(v1294)
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1318)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1318)+28)) = v1319 + int32(1)
	goto L22
L357:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1309)+24)) = v1307
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+28))
	v1313 = v1311
	v1314 = v1312
	goto L356
L358:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L359:
	;
	v1356 = F_palloc(m, v1351+int32(1))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L18
	} else {
		goto L360
	}
L360:
	;
	if v1351 != 0 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1358)+24))
	base.MemoryCopy(m, v1356, v1359, v1351)
	goto L363
L362:
	;
	goto L363
L363:
	;
	v1362 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1351+v1356))) = uint8(v1362)
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1364)+28))
	if int32(64) <= v1365 {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	F_truncate_identifier(m, v1356, v1365, int32(1))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L18
	} else {
		goto L367
	}
L365:
	;
	goto L366
L366:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1371))) = v1356
	v2533 = int32(258)
	goto L62
L367:
	;
	goto L366
L368:
	;
	v1382 = F_palloc(m, v1377+int32(1))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L18
	} else {
		goto L369
	}
L369:
	;
	if v1377 != 0 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1384)+24))
	base.MemoryCopy(m, v1382, v1385, v1377)
	goto L372
L371:
	;
	goto L372
L372:
	;
	v1388 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1377+v1382))) = uint8(v1388)
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1390))) = v1382
	v2533 = int32(259)
	goto L62
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1393)+32)) = v1397 << (uint(int32(1)) % 32)
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1402)+24))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1402)+32))
	v1405 = F_repalloc(m, v1403, v1404)
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L18
	} else {
		goto L376
	}
L374:
	;
	v1412 = v1393
	v1413 = v1394
	goto L375
L375:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+24))
	v1416 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v1413+v1414))) = uint8(v1416)
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1418)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1418)+28)) = v1419 + int32(1)
	goto L22
L376:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1407)+24)) = v1405
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+28))
	v1412 = v1409
	v1413 = v1410
	goto L375
L377:
	;
	v1430 = int32(1)
	v1433 = v1427 + v1430
	if v1433&v1427 != 0 {
		goto L380
	} else {
		goto L381
	}
L378:
	;
	v1449 = v1424
	v1450 = v1425
	goto L379
L379:
	;
	if v1426 != 0 {
		goto L384
	} else {
		goto L385
	}
L380:
	;
	v1438 = v1430 << (uint(int32(32)-base.I32_clz(v1433)) % 32)
	goto L382
L381:
	;
	v1438 = v1433
	goto L382
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1424)+32)) = v1438
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+24))
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+32))
	v1443 = F_repalloc(m, v1441, v1442)
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L18
	} else {
		goto L383
	}
L383:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v1445)+24)) = v1443
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+28))
	v1449 = v1447
	v1450 = v1448
	goto L379
L384:
	;
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1449)+24))
	base.MemoryCopy(m, v1451+v1450, v1423, v1426)
	goto L386
L385:
	;
	goto L386
L386:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1454)+28)) = v1455 + v1426
	goto L22
L387:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L388:
	;
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1485))) = v1483
	v2533 = int32(258)
	goto L62
L389:
	;
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1642)))
	*(*int32)(unsafe.Add(mBase, uint32(v1641))) = v1544 - v1643
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	if v1646 <= v1631 {
		goto L414
	} else {
		goto L415
	}
L390:
	;
	v1553 = v1548
	goto L392
L391:
	;
	v1553 = v1546
	goto L392
L392:
	;
	if v1546 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1554 = v1553
	goto L395
L394:
	;
	v1554 = v1548
	goto L395
L395:
	;
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	if v1554 != 0 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v1557 = v1554 - v1544
	goto L398
L397:
	;
	v1557 = v1556
	goto L398
L398:
	;
	if v1557 < int32(2) {
		v1631 = v1557
		goto L389
	} else {
		goto L399
	}
L399:
	;
	v1563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1544+v1557-int32(1)))))
	switch v1563 - int32(43) {
	case 0, 2:
		goto L400
	default:
		v1631 = v1557
		goto L389
	}
L400:
	;
	v1574 = v1557 - int32(2)
	goto L401
L401:
	;
	v1582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1544+v1574))))
	switch v1582 - int32(33) {
	case 0, 2, 4, 5, 30, 31, 61, 63, 91, 93:
		v1631 = v1557
		goto L389
	default:
		goto L403
	}
L402:
	;
	if v1557 != int32(2) {
		goto L405
	} else {
		goto L406
	}
L403:
	;
	if int32(0) < v1574 {
		v1574 = v1574 - int32(1)
		goto L401
	} else {
		goto L404
	}
L404:
	;
	goto L402
L405:
	;
	v1597 = v1557
	goto L408
L406:
	;
	goto L407
L407:
	;
	v1631 = int32(1)
	goto L389
L408:
	;
	v1605 = v1597 - int32(1)
	v1609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1544+v1597-int32(2)))))
	switch v1609 - int32(43) {
	case 0, 2:
		goto L410
	default:
		v1631 = v1605
		goto L389
	}
L409:
	;
	goto L407
L410:
	;
	if int32(3) < v1597 {
		v1597 = v1605
		goto L408
	} else {
		goto L411
	}
L411:
	;
	goto L409
L412:
	;
	v1798 = F_pstrdup(m, v1796)
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L18
	} else {
		goto L454
	}
L413:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1689 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1688))))
	goto L432
L414:
	;
	if v1631 <= int32(63) {
		goto L424
	} else {
		goto L425
	}
L415:
	;
	v1648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v1648)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v173
	v1651 = v1631 + v173
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1651
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v1631
	v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1651))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v1654)
	v1656 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1651))) = uint8(v1656)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v1651
	switch v1631 - int32(1) {
	case 0:
		goto L413
	case 1:
		goto L416
	default:
		goto L414
	}
L416:
	;
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1661))))
	switch v1662 - int32(33) {
	case 0:
		goto L417
	default:
		v1796 = v1661
		goto L412
	case 27:
		goto L418
	case 28:
		goto L420
	case 29:
		goto L419
	}
L417:
	;
	v1677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1661)+1)))
	if v1677 != int32(61) {
		v1796 = v1661
		goto L412
	} else {
		goto L423
	}
L418:
	;
	v1674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1661)+1)))
	switch v1674 - int32(61) {
	case 0:
		v2533 = int32(272)
		goto L62
	case 1:
		goto L63
	default:
		v1796 = v1661
		goto L412
	}
L419:
	;
	v1669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1661)+1)))
	if v1669 != int32(61) {
		v1796 = v1661
		goto L412
	} else {
		goto L422
	}
L420:
	;
	v1665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1661)+1)))
	if v1665 != int32(62) {
		v1796 = v1661
		goto L412
	} else {
		goto L421
	}
L421:
	;
	v2533 = int32(271)
	goto L62
L422:
	;
	v2533 = int32(273)
	goto L62
L423:
	;
	goto L63
L424:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v1796 = v1683
	goto L412
L425:
	;
	goto L426
L426:
	;
	F_scanner_yyerror(m, int32(_a_F_core_yylex_38), v91)
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L18
	} else {
		goto L427
	}
L427:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L428:
	;
	if v1795 != 0 {
		v2533 = v1689
		goto L62
	} else {
		goto L453
	}
L429:
	;
	v1795 = int32(0)
	goto L428
L430:
	;
	v1773 = v1766
	v1775 = v1768
	goto L447
L431:
	;
	if base.B2i32(v1712 != v1713) == int32(0) {
		goto L429
	} else {
		goto L438
	}
L432:
	;
	v1704 = int32(_a_F_core_yylex_39)
	v1706 = int32(18)
	goto L433
L433:
	;
	v1709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1704))))
	if v1709 == v1689&int32(255) {
		v1766 = v1704
		v1768 = v1706
		goto L430
	} else {
		goto L435
	}
L434:
	;
	goto L431
L435:
	;
	v1711 = int32(1)
	v1712 = v1706 - v1711
	v1713 = int32(0)
	v1716 = v1704 + v1711
	if v1716&int32(3) == v1713 {
		goto L431
	} else {
		goto L436
	}
L436:
	;
	if v1712 != 0 {
		v1704 = v1716
		v1706 = v1712
		goto L433
	} else {
		goto L437
	}
L437:
	;
	goto L434
L438:
	;
	v1729 = v1689 & int32(255)
	v1730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1716))))
	if base.B2i32(v1729 == v1730)|base.B2i32(base.Ui32(v1712) < base.Ui32(int32(4))) == int32(0) {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v1739 = v1716
	v1741 = v1712
	goto L442
L440:
	;
	v1759 = v1716
	v1761 = v1712
	goto L441
L441:
	;
	if v1761 == int32(0) {
		goto L429
	} else {
		goto L446
	}
L442:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1739)))
	v1746 = v1745 ^ v1729*int32(16843009)
	v1749 = int32(-2139062144)
	if (int32(16843008)-v1746|v1746)&v1749 != v1749 {
		v1766 = v1739
		v1768 = v1741
		goto L430
	} else {
		goto L444
	}
L443:
	;
	v1759 = v1754
	v1761 = v1756
	goto L441
L444:
	;
	v1753 = int32(4)
	v1754 = v1739 + v1753
	v1756 = v1741 - v1753
	if base.Ui32(int32(3)) < base.Ui32(v1756) {
		v1739 = v1754
		v1741 = v1756
		goto L442
	} else {
		goto L445
	}
L445:
	;
	goto L443
L446:
	;
	v1766 = v1759
	v1768 = v1761
	goto L430
L447:
	;
	v1778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1773))))
	if v1689&int32(255) == v1778 {
		goto L449
	} else {
		goto L450
	}
L448:
	;
	goto L429
L449:
	;
	v1795 = v1773
	goto L428
L450:
	;
	goto L451
L451:
	;
	v1780 = int32(1)
	v1783 = v1775 - v1780
	if v1783 != 0 {
		v1773 = v1773 + v1780
		v1775 = v1783
		goto L447
	} else {
		goto L452
	}
L452:
	;
	goto L448
L453:
	;
	v1796 = v1688
	goto L412
L454:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1800))) = v1798
	v2533 = int32(265)
	goto L62
L455:
	;
	v1820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v1820 == int32(1) {
		goto L56
	} else {
		goto L456
	}
L456:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1823))) = v1818
	v2533 = int32(267)
	goto L62
L457:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L458:
	;
	v1852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v1852 == int32(1) {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v1856 = F_pstrdup(m, v1842)
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L18
	} else {
		goto L462
	}
L460:
	;
	v1858 = int32(266)
	v1859 = v1849
	goto L461
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1841))) = v1859
	v2533 = v1858
	goto L62
L462:
	;
	v1858 = int32(260)
	v1859 = v1856
	goto L461
L463:
	;
	v1878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v1878 == int32(1) {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v1882 = F_pstrdup(m, v1868)
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L18
	} else {
		goto L467
	}
L465:
	;
	v1884 = int32(266)
	v1885 = v1875
	goto L466
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1867))) = v1885
	v2533 = v1884
	goto L62
L467:
	;
	v1884 = int32(260)
	v1885 = v1882
	goto L466
L468:
	;
	v1904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v1904 == int32(1) {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v1908 = F_pstrdup(m, v1894)
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L18
	} else {
		goto L472
	}
L470:
	;
	v1910 = int32(266)
	v1911 = v1901
	goto L471
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1893))) = v1911
	v2533 = v1910
	goto L62
L472:
	;
	v1910 = int32(260)
	v1911 = v1908
	goto L471
L473:
	;
	v1930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v1930 == int32(1) {
		goto L474
	} else {
		goto L475
	}
L474:
	;
	v1934 = F_pstrdup(m, v1920)
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L18
	} else {
		goto L477
	}
L475:
	;
	v1936 = int32(266)
	v1937 = v1927
	goto L476
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1919))) = v1937
	v2533 = v1936
	goto L62
L477:
	;
	v1936 = int32(260)
	v1937 = v1934
	goto L476
L478:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L479:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L480:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L481:
	;
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1975))) = v1973
	v2533 = int32(260)
	goto L62
L482:
	;
	v2009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v2009 == int32(1) {
		goto L483
	} else {
		goto L484
	}
L483:
	;
	v2013 = F_pstrdup(m, v1999)
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L18
	} else {
		goto L486
	}
L484:
	;
	v2015 = int32(266)
	v2016 = v2006
	goto L485
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1998))) = v2016
	v2533 = v2015
	goto L62
L486:
	;
	v2015 = int32(260)
	v2016 = v2013
	goto L485
L487:
	;
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2027))) = v2025
	v2533 = int32(260)
	goto L62
L488:
	;
	if int32(0) <= v2039 {
		goto L489
	} else {
		goto L490
	}
L489:
	;
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v2045 = *(*int32)(unsafe.Add(mBase, uint32(v2044)+8))
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(v2045)))
	v2048 = v2039 << (uint(int32(1)) % 32)
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v2045)+4))
	v2051 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2048+v2049))))
	*(*int32)(unsafe.Add(mBase, uint32(v2043))) = v2046 + v2051
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v2054)+12))
	v2057 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2055+v2048))))
	v2533 = v2057
	goto L62
L490:
	;
	goto L491
L491:
	;
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v2061 = F_downcase_truncate_identifier(m, v2058, v2059, int32(1))
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L18
	} else {
		goto L492
	}
L492:
	;
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2063))) = v2061
	v2533 = int32(258)
	goto L62
L493:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L494:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v2092)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v2096
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v2091)))
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2098))) = v2099
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2102 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2103 = int32(2)
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2101+v2102<<(uint(v2103)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2106)+44)) = int32(1)
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v2109+v2110<<(uint(v2103)%32))))
	v2115 = v2114
	v2116 = v2109
	v2117 = v2110
	goto L496
L495:
	;
	v2115 = v2092
	v2116 = v2087
	v2117 = v2088
	goto L496
L496:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v91)+36))
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v2115)+4))
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v2121 = v2119 + v2120
	if base.Ui32(v2118) <= base.Ui32(v2121) {
		goto L497
	} else {
		goto L498
	}
L497:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2126 = v2084 ^ int32(-1) + v170
	v2127 = v2123 + v2126
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v2127
	v2129 = *(*int32)(unsafe.Add(mBase, uint32(v91)+44))
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2129<<(uint(int32(2))%32))+uint32(_c_F_core_yylex[2])))
	if v2126 <= int32(0) {
		v3222 = v2134
		goto L37
	} else {
		goto L500
	}
L498:
	;
	goto L499
L499:
	;
	if base.Ui32(v2121+int32(1)) < base.Ui32(v2118) {
		goto L54
	} else {
		goto L508
	}
L500:
	;
	v2140 = int32(0)
	v2142 = v2126 & int32(3)
	if v2142 == v2140 {
		goto L55
	} else {
		goto L501
	}
L501:
	;
	v2145 = v2140
	v2146 = v2123
	v2148 = v2134
	goto L502
L502:
	;
	v2159 = v2146 + int32(1)
	v2160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2146))))
	if v2160 != 0 {
		goto L504
	} else {
		goto L505
	}
L503:
	;
	v3149 = v2159
	v3151 = v2169
	goto L40
L504:
	;
	v2162 = v2160
	goto L506
L505:
	;
	v2162 = int32(256)
	goto L506
L506:
	;
	v2163 = int32(3)
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(v2148+v2162<<(uint(v2163)%32))+4))
	v2169 = v2148 + v2166<<(uint(v2163)%32)
	v2171 = v2145 + int32(1)
	if v2171 != v2142 {
		v2145 = v2171
		v2146 = v2159
		v2148 = v2169
		goto L502
	} else {
		goto L507
	}
L507:
	;
	goto L503
L508:
	;
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v2115)+40))
	if v2177 == int32(0) {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	if v2118-v2176 != int32(1) {
		v2904 = v2119
		v2909 = v2176
		v2911 = v2120
		goto L50
	} else {
		goto L512
	}
L510:
	;
	goto L511
L511:
	;
	v2185 = v2176 ^ int32(-1) + v2118
	if int32(0) < v2185 {
		goto L513
	} else {
		goto L514
	}
L512:
	;
	v3130 = v2176
	goto L41
L513:
	;
	v2188 = int32(7)
	v2189 = v2185 & v2188
	if base.Ui32(v2118-v2176-int32(2)) < base.Ui32(v2188) {
		goto L518
	} else {
		goto L519
	}
L514:
	;
	v2294 = v2115
	v2297 = v2116
	v2300 = v2117
	goto L515
L515:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v2294)+44))
	if v2304 == int32(2) {
		goto L528
	} else {
		goto L529
	}
L516:
	;
	v2285 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v2285+v2286<<(uint(int32(2))%32))))
	v2294 = v2290
	v2297 = v2285
	v2300 = v2286
	goto L515
L517:
	;
	v2250 = v2236
	v2253 = v2239
	v2256 = int32(0)
	goto L525
L518:
	;
	v2236 = v2119
	v2239 = v2176
	goto L517
L519:
	;
	goto L520
L520:
	;
	v2198 = v2119
	v2201 = v2176
	v2204 = int32(0)
	goto L521
L521:
	;
	v2211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2201))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2198))) = uint8(v2211)
	v2213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2201)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2198)+1)) = uint8(v2213)
	v2215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2201)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2198)+2)) = uint8(v2215)
	v2217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2201)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2198)+3)) = uint8(v2217)
	v2219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2201)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2198)+4)) = uint8(v2219)
	v2221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2201)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2198)+5)) = uint8(v2221)
	v2223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2201)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2198)+6)) = uint8(v2223)
	v2225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2201)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2198)+7)) = uint8(v2225)
	v2227 = int32(8)
	v2228 = v2198 + v2227
	v2230 = v2201 + v2227
	v2232 = v2204 + v2227
	if v2232 != v2185&int32(2147483640) {
		v2198 = v2228
		v2201 = v2230
		v2204 = v2232
		goto L521
	} else {
		goto L523
	}
L522:
	;
	if v2189 == int32(0) {
		goto L516
	} else {
		goto L524
	}
L523:
	;
	goto L522
L524:
	;
	v2236 = v2228
	v2239 = v2230
	goto L517
L525:
	;
	v2263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2253))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2250))) = uint8(v2263)
	v2265 = int32(1)
	v2270 = v2256 + v2265
	if v2270 != v2189 {
		v2250 = v2250 + v2265
		v2253 = v2253 + v2265
		v2256 = v2270
		goto L525
	} else {
		goto L527
	}
L526:
	;
	goto L516
L527:
	;
	goto L526
L528:
	;
	v2307 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v2307
	v2641 = v2307
	v2647 = v2297 + v2300<<(uint(int32(2))%32)
	goto L51
L529:
	;
	goto L530
L530:
	;
	v2313 = int32(0)
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v2294)+12))
	v2315 = v2176 - v2118
	v2316 = v2314 + v2315
	if v2316 <= v2313 {
		goto L531
	} else {
		goto L532
	}
L531:
	;
	v2319 = *(*int32)(unsafe.Add(mBase, uint32(v91)+36))
	v2323 = v2294
	v2325 = v2319
	v2329 = v2314
	goto L534
L532:
	;
	v2371 = v2316
	v2374 = v2294
	goto L533
L533:
	;
	v2384 = int32(_a_F_core_yylex_40)
	if base.Ui32(v2384) <= base.Ui32(v2371) {
		goto L550
	} else {
		goto L551
	}
L534:
	;
	v2333 = *(*int32)(unsafe.Add(mBase, uint32(v2323)+20))
	if v2333 == int32(0) {
		goto L536
	} else {
		goto L537
	}
L535:
	;
	v2371 = v2368
	v2374 = v2366
	goto L533
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2323)+4)) = int32(0)
	goto L42
L537:
	;
	goto L538
L538:
	;
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(v2323)+4))
	v2340 = v2329 << (uint(int32(1)) % 32)
	if v2340 <= int32(0) {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	v2344 = base.I32_div_s(v2329, int32(8))
	v2346 = v2344 + v2329
	goto L541
L540:
	;
	v2346 = v2340
	goto L541
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2323)+12)) = v2346
	v2349 = v2346 + int32(2)
	if v2338 != 0 {
		goto L543
	} else {
		goto L544
	}
L542:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2323)+4)) = v2354
	if v2354 == int32(0) {
		goto L42
	} else {
		goto L548
	}
L543:
	;
	v2350 = F_repalloc(m, v2338, v2349)
	mBase = m.M
	v2351 = m.ExcPending
	if v2351 != 0 {
		goto L18
	} else {
		goto L546
	}
L544:
	;
	goto L545
L545:
	;
	v2352 = F_palloc(m, v2349)
	mBase = m.M
	v2353 = m.ExcPending
	if v2353 != 0 {
		goto L18
	} else {
		goto L547
	}
L546:
	;
	v2354 = v2350
	goto L542
L547:
	;
	v2354 = v2352
	goto L542
L548:
	;
	v2359 = v2354 + (v2325 - v2338)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v2359
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2366 = *(*int32)(unsafe.Add(mBase, uint32(v2361+v2362<<(uint(int32(2))%32))))
	v2367 = *(*int32)(unsafe.Add(mBase, uint32(v2366)+12))
	v2368 = v2367 + v2315
	if v2368 <= int32(0) {
		v2323 = v2366
		v2325 = v2359
		v2329 = v2367
		goto L534
	} else {
		goto L549
	}
L549:
	;
	goto L535
L550:
	;
	v2387 = v2384
	goto L552
L551:
	;
	v2387 = v2371
	goto L552
L552:
	;
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v2374)+24))
	if v2388 != 0 {
		goto L553
	} else {
		goto L554
	}
L553:
	;
	v2396 = v2313
	goto L557
L554:
	;
	goto L555
L555:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_core_yylex[7])) = int32(0)
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(v2450+v2451<<(uint(int32(2))%32))))
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v2455)+4))
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2460 = F_fread(m, v2456+v2185, int32(1), v2387, v2459)
	mBase = m.M
	v2461 = m.ExcPending
	if v2461 != 0 {
		goto L18
	} else {
		goto L568
	}
L556:
	;
	switch v2406 {
	case 0:
		goto L564
	default:
		v2445 = v2420
		goto L562
	case 11:
		goto L563
	}
L557:
	;
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2403 = F_do_getc(m, v2402)
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L18
	} else {
		goto L560
	}
L558:
	;
	v2420 = v2387
	goto L556
L559:
	;
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2408 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2412 = *(*int32)(unsafe.Add(mBase, uint32(v2407+v2408<<(uint(int32(2))%32))))
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v2412)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v2413+v2185+v2396))) = uint8(v2403)
	v2418 = v2396 + int32(1)
	if v2418 != v2387 {
		v2396 = v2418
		goto L557
	} else {
		goto L561
	}
L560:
	;
	v2406 = v2403 + int32(1)
	switch v2406 {
	case 0, 11:
		v2420 = v2396
		goto L556
	default:
		goto L559
	}
L561:
	;
	goto L558
L562:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v2445
	v2623 = v2445
	goto L52
L563:
	;
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(v2432+v2433<<(uint(int32(2))%32))))
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(v2437)+4))
	v2441 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v2438+v2185+v2420))) = uint8(v2441)
	v2445 = v2420 + int32(1)
	goto L562
L564:
	;
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2422 = *(*int32)(unsafe.Add(mBase, uint32(v2421)))
	goto L565
L565:
	;
	if int32(base.Ui32(v2422)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v2445 = v2420
		goto L562
	} else {
		goto L566
	}
L566:
	;
	F_yy_fatal_error_2(m, int32(_a_F_core_yylex_4))
	mBase = m.M
	v2431 = m.ExcPending
	if v2431 != 0 {
		goto L18
	} else {
		goto L567
	}
L567:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L568:
	;
	v2469 = v2460
	goto L569
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v2469
	if v2469 != 0 {
		v2623 = v2469
		goto L52
	} else {
		goto L571
	}
L571:
	;
	v2476 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2477 = *(*int32)(unsafe.Add(mBase, uint32(v2476)))
	goto L572
L572:
	;
	if int32(base.Ui32(v2477)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v2623 = int32(0)
	goto L52
L574:
	;
	goto L575
L575:
	;
	v2486 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[7]))
	if v2486 != int32(27) {
		goto L53
	} else {
		goto L576
	}
L576:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_core_yylex[7])) = int32(0)
	v2492 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2493 = *(*int32)(unsafe.Add(mBase, uint32(v2492)))
	*(*int32)(unsafe.Add(mBase, uint32(v2492))) = v2493 & int32(-49)
	goto L577
L577:
	;
	v2497 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2498 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(v2497+v2498<<(uint(int32(2))%32))))
	v2503 = *(*int32)(unsafe.Add(mBase, uint32(v2502)+4))
	v2506 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2507 = F_fread(m, v2503+v2185, int32(1), v2387, v2506)
	mBase = m.M
	v2508 = m.ExcPending
	if v2508 != 0 {
		goto L18
	} else {
		goto L578
	}
L578:
	;
	v2469 = v2507
	goto L569
L579:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L580:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2555 = m.ExcPending
	if v2555 != 0 {
		goto L18
	} else {
		goto L581
	}
L581:
	;
	F_errmsg(m, int32(_a_F_core_yylex_41), int32(0))
	mBase = m.M
	v2559 = m.ExcPending
	if v2559 != 0 {
		goto L18
	} else {
		goto L582
	}
L582:
	;
	F_errdetail(m, int32(_a_F_core_yylex_42), int32(0))
	mBase = m.M
	v2563 = m.ExcPending
	if v2563 != 0 {
		goto L18
	} else {
		goto L583
	}
L583:
	;
	v2564 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v2565 = *(*int32)(unsafe.Add(mBase, uint32(v2564)))
	F_scanner_errposition(m, v2565, v91)
	mBase = m.M
	v2567 = m.ExcPending
	if v2567 != 0 {
		goto L18
	} else {
		goto L584
	}
L584:
	;
	F_errfinish(m, int32(_a_F_core_yylex_31), int32(569), int32(_a_F_core_yylex_32))
	mBase = m.M
	v2572 = m.ExcPending
	if v2572 != 0 {
		goto L18
	} else {
		goto L585
	}
L585:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L586:
	;
	F_errcode(m, int32(100794498))
	mBase = m.M
	v2583 = m.ExcPending
	if v2583 != 0 {
		goto L18
	} else {
		goto L587
	}
L587:
	;
	F_errmsg(m, int32(_a_F_core_yylex_43), int32(0))
	mBase = m.M
	v2587 = m.ExcPending
	if v2587 != 0 {
		goto L18
	} else {
		goto L588
	}
L588:
	;
	F_errhint(m, int32(_a_F_core_yylex_44), int32(0))
	mBase = m.M
	v2591 = m.ExcPending
	if v2591 != 0 {
		goto L18
	} else {
		goto L589
	}
L589:
	;
	v2592 = *(*int32)(unsafe.Add(mBase, uint32(v91)+96))
	v2593 = *(*int32)(unsafe.Add(mBase, uint32(v2592)))
	F_scanner_errposition(m, v2593, v91)
	mBase = m.M
	v2595 = m.ExcPending
	if v2595 != 0 {
		goto L18
	} else {
		goto L590
	}
L590:
	;
	F_errfinish(m, int32(_a_F_core_yylex_31), int32(716), int32(_a_F_core_yylex_32))
	mBase = m.M
	v2600 = m.ExcPending
	if v2600 != 0 {
		goto L18
	} else {
		goto L591
	}
L591:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L592:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L593:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L594:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L595:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L596:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L597:
	;
	v2774 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v2775 = v2774 + v2185
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2777 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2781 = *(*int32)(unsafe.Add(mBase, uint32(v2776+v2777<<(uint(int32(2))%32))))
	v2782 = *(*int32)(unsafe.Add(mBase, uint32(v2781)+12))
	if v2782 < v2775 {
		goto L621
	} else {
		goto L622
	}
L598:
	;
	if v2185 == int32(0) {
		goto L599
	} else {
		goto L600
	}
L599:
	;
	v2654 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2655 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v2655 != 0 {
		goto L604
	} else {
		goto L605
	}
L600:
	;
	goto L601
L601:
	;
	v2759 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2760 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2761 = int32(2)
	v2764 = *(*int32)(unsafe.Add(mBase, uint32(v2759+v2760<<(uint(v2761)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2764)+44)) = v2761
	v2773 = v2761
	goto L597
L602:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2721)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2721))) = v2654
	v2728 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v2728 != 0 {
		goto L617
	} else {
		goto L618
	}
L603:
	;
	v2678 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[7]))
	v2679 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v2676+v2679<<(uint(int32(2))%32))))
	if v2683 == int32(0) {
		goto L611
	} else {
		goto L612
	}
L604:
	;
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2660 = *(*int32)(unsafe.Add(mBase, uint32(v2655+v2656<<(uint(int32(2))%32))))
	if v2660 != 0 {
		v2676 = v2655
		goto L603
	} else {
		goto L607
	}
L605:
	;
	goto L606
L606:
	;
	F_core_yyensure_buffer_stack(m, v91)
	mBase = m.M
	v2662 = m.ExcPending
	if v2662 != 0 {
		goto L18
	} else {
		goto L608
	}
L607:
	;
	goto L606
L608:
	;
	v2663 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v2664 = F_core_yy_create_buffer(m, v2663, v91)
	mBase = m.M
	v2665 = m.ExcPending
	if v2665 != 0 {
		goto L18
	} else {
		goto L609
	}
L609:
	;
	v2666 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2667 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2666+v2667<<(uint(int32(2))%32)))) = v2664
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v2672 != 0 {
		v2676 = v2672
		goto L603
	} else {
		goto L610
	}
L610:
	;
	v2674 = *(*int32)(unsafe.Add(mBase, _c_F_core_yylex[7]))
	v2721 = int32(0)
	v2724 = v2674
	goto L602
L611:
	;
	v2721 = int32(0)
	v2724 = v2678
	goto L602
L612:
	;
	goto L613
L613:
	;
	v2687 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2683)+16)) = v2687
	v2689 = *(*int32)(unsafe.Add(mBase, uint32(v2683)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v2689))) = uint8(v2687)
	v2692 = *(*int32)(unsafe.Add(mBase, uint32(v2683)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v2692)+1)) = uint8(v2687)
	*(*int32)(unsafe.Add(mBase, uint32(v2683)+44)) = v2687
	*(*int32)(unsafe.Add(mBase, uint32(v2683)+28)) = int32(1)
	v2699 = *(*int32)(unsafe.Add(mBase, uint32(v2683)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2683)+8)) = v2699
	v2701 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v2701 == v2687 {
		v2721 = v2683
		v2724 = v2678
		goto L602
	} else {
		goto L614
	}
L614:
	;
	v2704 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2707 = v2701 + v2704<<(uint(int32(2))%32)
	v2708 = *(*int32)(unsafe.Add(mBase, uint32(v2707)))
	if v2683 != v2708 {
		v2721 = v2683
		v2724 = v2678
		goto L602
	} else {
		goto L615
	}
L615:
	;
	v2710 = *(*int32)(unsafe.Add(mBase, uint32(v2708)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v2710
	v2712 = *(*int32)(unsafe.Add(mBase, uint32(v2707)))
	v2713 = *(*int32)(unsafe.Add(mBase, uint32(v2712)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v2713
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v2713
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(v2707)))
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v2716)))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v2717
	v2719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2713))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v2719)
	v2721 = v2683
	v2724 = v2678
	goto L602
L616:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2721)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_core_yylex[7])) = v2724
	v2741 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2742 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2745 = v2741 + v2742<<(uint(int32(2))%32)
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v2745)))
	v2747 = *(*int32)(unsafe.Add(mBase, uint32(v2746)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v2747
	v2749 = *(*int32)(unsafe.Add(mBase, uint32(v2745)))
	v2750 = *(*int32)(unsafe.Add(mBase, uint32(v2749)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v2750
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v2750
	v2753 = *(*int32)(unsafe.Add(mBase, uint32(v2745)))
	v2754 = *(*int32)(unsafe.Add(mBase, uint32(v2753)))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v2754
	v2756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2750))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v2756)
	v2773 = int32(1)
	goto L597
L617:
	;
	v2729 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(v2728+v2729<<(uint(int32(2))%32))))
	if v2721 == v2733 {
		goto L616
	} else {
		goto L620
	}
L618:
	;
	goto L619
L619:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2721)+32)) = int64(1)
	goto L616
L620:
	;
	goto L619
L621:
	;
	v2786 = v2775 + v2774>>(uint(int32(1))%32)
	v2787 = *(*int32)(unsafe.Add(mBase, uint32(v2781)+4))
	if v2787 != 0 {
		goto L625
	} else {
		goto L626
	}
L622:
	;
	v2816 = v2775
	v2817 = v2776
	v2819 = v2777
	goto L623
L623:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v2816
	v2821 = int32(2)
	v2824 = *(*int32)(unsafe.Add(mBase, uint32(v2817+v2819<<(uint(v2821)%32))))
	v2825 = *(*int32)(unsafe.Add(mBase, uint32(v2824)+4))
	v2827 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2825+v2816))) = uint8(v2827)
	v2829 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2830 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2834 = *(*int32)(unsafe.Add(mBase, uint32(v2829+v2830<<(uint(v2821)%32))))
	v2835 = *(*int32)(unsafe.Add(mBase, uint32(v2834)+4))
	v2836 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v2835+v2836)+1)) = uint8(v2827)
	v2840 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2841 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2844 = v2840 + v2841<<(uint(v2821)%32)
	v2845 = *(*int32)(unsafe.Add(mBase, uint32(v2844)))
	v2846 = *(*int32)(unsafe.Add(mBase, uint32(v2845)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+80)) = v2846
	if v2773 == int32(1) {
		v3130 = v2846
		goto L41
	} else {
		goto L631
	}
L624:
	;
	v2793 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2794 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2795 = int32(2)
	v2798 = *(*int32)(unsafe.Add(mBase, uint32(v2793+v2794<<(uint(v2795)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2798)+4)) = v2792
	v2800 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2801 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2805 = *(*int32)(unsafe.Add(mBase, uint32(v2800+v2801<<(uint(v2795)%32))))
	v2806 = *(*int32)(unsafe.Add(mBase, uint32(v2805)+4))
	if v2806 == int32(0) {
		goto L48
	} else {
		goto L630
	}
L625:
	;
	v2788 = F_repalloc(m, v2787, v2786)
	mBase = m.M
	v2789 = m.ExcPending
	if v2789 != 0 {
		goto L18
	} else {
		goto L628
	}
L626:
	;
	goto L627
L627:
	;
	v2790 = F_palloc(m, v2786)
	mBase = m.M
	v2791 = m.ExcPending
	if v2791 != 0 {
		goto L18
	} else {
		goto L629
	}
L628:
	;
	v2792 = v2788
	goto L624
L629:
	;
	v2792 = v2790
	goto L624
L630:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2805)+12)) = v2786 - int32(2)
	v2812 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v2813 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2814 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v2816 = v2814 + v2185
	v2817 = v2813
	v2819 = v2812
	goto L623
L631:
	;
	switch v2773 - int32(1) {
	case 0:
		goto L49
	case 1:
		goto L632
	default:
		goto L633
	}
L632:
	;
	v2901 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v2902 = *(*int32)(unsafe.Add(mBase, uint32(v2844)))
	v2903 = *(*int32)(unsafe.Add(mBase, uint32(v2902)+4))
	v2904 = v2903
	v2909 = v2846
	v2911 = v2901
	goto L50
L633:
	;
	v2854 = v2084 ^ int32(-1) + v170
	v2855 = v2846 + v2854
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v2855
	v2857 = *(*int32)(unsafe.Add(mBase, uint32(v91)+44))
	v2862 = *(*int32)(unsafe.Add(mBase, uint32(v2857<<(uint(int32(2))%32))+uint32(_c_F_core_yylex[2])))
	if v2854 <= int32(0) {
		v3110 = v2862
		goto L43
	} else {
		goto L634
	}
L634:
	;
	v2868 = int32(0)
	v2870 = v2854 & int32(3)
	if v2870 == v2868 {
		goto L47
	} else {
		goto L635
	}
L635:
	;
	v2873 = v2868
	v2874 = v2846
	v2876 = v2862
	goto L636
L636:
	;
	v2887 = v2874 + int32(1)
	v2888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2874))))
	if v2888 != 0 {
		goto L638
	} else {
		goto L639
	}
L637:
	;
	v3037 = v2887
	v3039 = v2897
	goto L44
L638:
	;
	v2890 = v2888
	goto L640
L639:
	;
	v2890 = int32(256)
	goto L640
L640:
	;
	v2891 = int32(3)
	v2894 = *(*int32)(unsafe.Add(mBase, uint32(v2876+v2890<<(uint(v2891)%32))+4))
	v2897 = v2876 + v2894<<(uint(v2891)%32)
	v2899 = v2873 + int32(1)
	if v2899 != v2870 {
		v2873 = v2899
		v2874 = v2887
		v2876 = v2897
		goto L636
	} else {
		goto L641
	}
L641:
	;
	goto L637
L642:
	;
	v2926 = v2917 - v2909
	v2929 = int32(0)
	v2931 = v2926 & int32(3)
	if v2931 == v2929 {
		goto L46
	} else {
		goto L643
	}
L643:
	;
	v2934 = v2909
	v2937 = v2924
	v2940 = v2929
	goto L644
L644:
	;
	v2948 = v2934 + int32(1)
	v2949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2934))))
	if v2949 != 0 {
		goto L646
	} else {
		goto L647
	}
L645:
	;
	v2965 = v2948
	v2968 = v2958
	goto L45
L646:
	;
	v2951 = v2949
	goto L648
L647:
	;
	v2951 = int32(256)
	goto L648
L648:
	;
	v2952 = int32(3)
	v2955 = *(*int32)(unsafe.Add(mBase, uint32(v2937+v2951<<(uint(v2952)%32))+4))
	v2958 = v2937 + v2955<<(uint(v2952)%32)
	v2960 = v2940 + int32(1)
	if v2960 != v2931 {
		v2934 = v2948
		v2937 = v2958
		v2940 = v2960
		goto L644
	} else {
		goto L649
	}
L649:
	;
	goto L645
L650:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L651:
	;
	v2980 = v2965
	v2983 = v2968
	goto L652
L652:
	;
	v2993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2980))))
	if v2993 != 0 {
		goto L654
	} else {
		goto L655
	}
L653:
	;
	v170 = v2917
	v172 = v3032
	v173 = v2909
	goto L35
L654:
	;
	v2995 = v2993
	goto L656
L655:
	;
	v2995 = int32(256)
	goto L656
L656:
	;
	v2996 = int32(3)
	v2999 = *(*int32)(unsafe.Add(mBase, uint32(v2983+v2995<<(uint(v2996)%32))+4))
	v3002 = v2983 + v2999<<(uint(v2996)%32)
	v3003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2980)+1)))
	if v3003 != 0 {
		goto L657
	} else {
		goto L658
	}
L657:
	;
	v3005 = v3003
	goto L659
L658:
	;
	v3005 = int32(256)
	goto L659
L659:
	;
	v3006 = int32(3)
	v3009 = *(*int32)(unsafe.Add(mBase, uint32(v3002+v3005<<(uint(v3006)%32))+4))
	v3012 = v3002 + v3009<<(uint(v3006)%32)
	v3013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2980)+2)))
	if v3013 != 0 {
		goto L660
	} else {
		goto L661
	}
L660:
	;
	v3015 = v3013
	goto L662
L661:
	;
	v3015 = int32(256)
	goto L662
L662:
	;
	v3016 = int32(3)
	v3019 = *(*int32)(unsafe.Add(mBase, uint32(v3012+v3015<<(uint(v3016)%32))+4))
	v3022 = v3012 + v3019<<(uint(v3016)%32)
	v3023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2980)+3)))
	if v3023 != 0 {
		goto L663
	} else {
		goto L664
	}
L663:
	;
	v3025 = v3023
	goto L665
L664:
	;
	v3025 = int32(256)
	goto L665
L665:
	;
	v3026 = int32(3)
	v3029 = *(*int32)(unsafe.Add(mBase, uint32(v3022+v3025<<(uint(v3026)%32))+4))
	v3032 = v3022 + v3029<<(uint(v3026)%32)
	v3034 = v2980 + int32(4)
	if v3034 != v2917 {
		v2980 = v3034
		v2983 = v3032
		goto L652
	} else {
		goto L666
	}
L666:
	;
	goto L653
L667:
	;
	v3052 = v3037
	v3054 = v3039
	goto L668
L668:
	;
	v3064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3052))))
	if v3064 != 0 {
		goto L670
	} else {
		goto L671
	}
L669:
	;
	v3110 = v3103
	goto L43
L670:
	;
	v3066 = v3064
	goto L672
L671:
	;
	v3066 = int32(256)
	goto L672
L672:
	;
	v3067 = int32(3)
	v3070 = *(*int32)(unsafe.Add(mBase, uint32(v3054+v3066<<(uint(v3067)%32))+4))
	v3073 = v3054 + v3070<<(uint(v3067)%32)
	v3074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3052)+1)))
	if v3074 != 0 {
		goto L673
	} else {
		goto L674
	}
L673:
	;
	v3076 = v3074
	goto L675
L674:
	;
	v3076 = int32(256)
	goto L675
L675:
	;
	v3077 = int32(3)
	v3080 = *(*int32)(unsafe.Add(mBase, uint32(v3073+v3076<<(uint(v3077)%32))+4))
	v3083 = v3073 + v3080<<(uint(v3077)%32)
	v3084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3052)+2)))
	if v3084 != 0 {
		goto L676
	} else {
		goto L677
	}
L676:
	;
	v3086 = v3084
	goto L678
L677:
	;
	v3086 = int32(256)
	goto L678
L678:
	;
	v3087 = int32(3)
	v3090 = *(*int32)(unsafe.Add(mBase, uint32(v3083+v3086<<(uint(v3087)%32))+4))
	v3093 = v3083 + v3090<<(uint(v3087)%32)
	v3094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3052)+3)))
	if v3094 != 0 {
		goto L679
	} else {
		goto L680
	}
L679:
	;
	v3096 = v3094
	goto L681
L680:
	;
	v3096 = int32(256)
	goto L681
L681:
	;
	v3097 = int32(3)
	v3100 = *(*int32)(unsafe.Add(mBase, uint32(v3093+v3096<<(uint(v3097)%32))+4))
	v3103 = v3093 + v3100<<(uint(v3097)%32)
	v3105 = v3052 + int32(4)
	if v3105 != v2855 {
		v3052 = v3105
		v3054 = v3103
		goto L668
	} else {
		goto L682
	}
L682:
	;
	goto L669
L683:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L684:
	;
	v3164 = v3149
	v3166 = v3151
	goto L685
L685:
	;
	v3176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3164))))
	if v3176 != 0 {
		goto L687
	} else {
		goto L688
	}
L686:
	;
	v3222 = v3215
	goto L37
L687:
	;
	v3178 = v3176
	goto L689
L688:
	;
	v3178 = int32(256)
	goto L689
L689:
	;
	v3179 = int32(3)
	v3182 = *(*int32)(unsafe.Add(mBase, uint32(v3166+v3178<<(uint(v3179)%32))+4))
	v3185 = v3166 + v3182<<(uint(v3179)%32)
	v3186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3164)+1)))
	if v3186 != 0 {
		goto L690
	} else {
		goto L691
	}
L690:
	;
	v3188 = v3186
	goto L692
L691:
	;
	v3188 = int32(256)
	goto L692
L692:
	;
	v3189 = int32(3)
	v3192 = *(*int32)(unsafe.Add(mBase, uint32(v3185+v3188<<(uint(v3189)%32))+4))
	v3195 = v3185 + v3192<<(uint(v3189)%32)
	v3196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3164)+2)))
	if v3196 != 0 {
		goto L693
	} else {
		goto L694
	}
L693:
	;
	v3198 = v3196
	goto L695
L694:
	;
	v3198 = int32(256)
	goto L695
L695:
	;
	v3199 = int32(3)
	v3202 = *(*int32)(unsafe.Add(mBase, uint32(v3195+v3198<<(uint(v3199)%32))+4))
	v3205 = v3195 + v3202<<(uint(v3199)%32)
	v3206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3164)+3)))
	if v3206 != 0 {
		goto L696
	} else {
		goto L697
	}
L696:
	;
	v3208 = v3206
	goto L698
L697:
	;
	v3208 = int32(256)
	goto L698
L698:
	;
	v3209 = int32(3)
	v3212 = *(*int32)(unsafe.Add(mBase, uint32(v3205+v3208<<(uint(v3209)%32))+4))
	v3215 = v3205 + v3212<<(uint(v3209)%32)
	v3217 = v3164 + int32(4)
	if v3217 != v2127 {
		v3164 = v3217
		v3166 = v3215
		goto L685
	} else {
		goto L699
	}
L699:
	;
	goto L686
L700:
	;
	v3234 = *(*int32)(unsafe.Add(mBase, uint32(v3222)+2048))
	if v3234 != int32(256) {
		v170 = v2127
		v172 = v3222
		v173 = v2123
		goto L35
	} else {
		goto L701
	}
L701:
	;
	goto L36
L702:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L703:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
