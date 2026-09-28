package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RunObjectPostAlterHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v5 = l4
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l3
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_RunObjectPostAlterHook[0]))
	m.T0[v19].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(2), l0, l1, l2, v9+int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		m.G0 = v9 + int32(16)
		return
	}
}
func F_RunObjectPostCreateHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = l3
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v4)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_RunObjectPostCreateHook[0]))
	m.T0[v15].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(0), l0, l1, l2, v8+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
func F_getObjectIdentityParts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v222 int32
	_ = v222
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v332 int64
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v385 int64
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v474 int64
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v564 int64
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v724 int64
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int64
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v855 int64
	_ = v855
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v958 int64
	_ = v958
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
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
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1073 int32
	_ = v1073
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1128 int32
	_ = v1128
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1202 int64
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1217 int32
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1253 int32
	_ = v1253
	var v1255 int64
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1270 int32
	_ = v1270
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1308 int64
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1323 int32
	_ = v1323
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1363 int64
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1378 int32
	_ = v1378
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1416 int64
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
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
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1496 int64
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1519 int32
	_ = v1519
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1544 int32
	_ = v1544
	var v1548 int32
	_ = v1548
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1567 int32
	_ = v1567
	var v1572 int32
	_ = v1572
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1602 int32
	_ = v1602
	var v1607 int32
	_ = v1607
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1675 int64
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1690 int32
	_ = v1690
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1759 int32
	_ = v1759
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1782 int64
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1797 int32
	_ = v1797
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1827 int64
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1842 int32
	_ = v1842
	var v1847 int32
	_ = v1847
	var v1852 int64
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1878 int32
	_ = v1878
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
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
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1907 int32
	_ = v1907
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1926 int32
	_ = v1926
	var v1928 int32
	_ = v1928
	var v1936 int32
	_ = v1936
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1946 int32
	_ = v1946
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1958 int64
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1973 int32
	_ = v1973
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1996 int32
	_ = v1996
	var v2000 int32
	_ = v2000
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2021 int32
	_ = v2021
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2053 int32
	_ = v2053
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2077 int32
	_ = v2077
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2099 int32
	_ = v2099
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2110 int32
	_ = v2110
	var v2115 int32
	_ = v2115
	var v2119 int32
	_ = v2119
	var v2121 int32
	_ = v2121
	var v2125 int32
	_ = v2125
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2136 int32
	_ = v2136
	var v2141 int32
	_ = v2141
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2151 int64
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2155 int32
	_ = v2155
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2174 int32
	_ = v2174
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2218 int32
	_ = v2218
	var v2222 int32
	_ = v2222
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2237 int32
	_ = v2237
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2258 int32
	_ = v2258
	var v2261 int32
	_ = v2261
	var v2276 int32
	_ = v2276
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2283 int32
	_ = v2283
	var v2289 int32
	_ = v2289
	var v2294 int32
	_ = v2294
	var v2307 int32
	_ = v2307
	var v2321 int32
	_ = v2321
	var v2334 int32
	_ = v2334
	v13 = m.G0
	v15 = v13 - int32(1184)
	m.G0 = v15
	F_initStringInfo(m, v15+int32(1168))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v23
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v23
	goto L5
L4:
	;
	goto L5
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v27 <= int32(2752) {
		goto L56
	} else {
		goto L57
	}
L6:
	;
	m.G0 = v15 + int32(1184)
	return v2334
L7:
	;
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1168))
	v2334 = v2321
	goto L6
L8:
	;
	v2307 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1172))
	if v2307 != 0 {
		goto L7
	} else {
		goto L752
	}
L9:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L746
	}
L10:
	;
	v2144 = F_table_open(m, int32(826), int32(1))
	mBase = m.M
	v2145 = m.ExcPending
	if v2145 != 0 {
		goto L1
	} else {
		goto L707
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2129 = m.ExcPending
	if v2129 != 0 {
		goto L1
	} else {
		goto L704
	}
L12:
	;
	F_relation_close(m, v2121, int32(1))
	mBase = m.M
	v2125 = m.ExcPending
	if v2125 != 0 {
		goto L1
	} else {
		goto L703
	}
L13:
	;
	F_getRelationIdentity(m, v15+int32(1168), v83, l1, l3)
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L1
	} else {
		goto L702
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L1
	} else {
		goto L699
	}
L15:
	;
	if v27 == int32(826) {
		goto L10
	} else {
		goto L698
	}
L16:
	;
	v2034 = F_table_open(m, int32(3576), int32(1))
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L1
	} else {
		goto L679
	}
L17:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2012 = F_get_subscription_name(m, v2011, l3)
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L1
	} else {
		goto L673
	}
L18:
	;
	v1958 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v1959 = F_SearchSysCache1(m, int32(52), v1958)
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L1
	} else {
		goto L654
	}
L19:
	;
	v1922 = F_getPublicationSchemaInfo(m, l0, l3, v15+int32(1056), v15+int32(1040))
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L1
	} else {
		goto L640
	}
L20:
	;
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1898 = F_get_publication_name(m, v1897, l3)
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L1
	} else {
		goto L634
	}
L21:
	;
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	v1870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1869)+22)))
	v1871 = v1869 + v1870
	v1873 = v1871 + int32(4)
	v1874 = F_quote_identifier(m, v1873)
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
		goto L1
	} else {
		goto L625
	}
L22:
	;
	v1827 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v1828 = F_SearchSysCache1(m, int32(44), v1827)
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L1
	} else {
		goto L609
	}
L23:
	;
	v1782 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v1783 = F_SearchSysCache1(m, int32(26), v1782)
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L1
	} else {
		goto L593
	}
L24:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1745 = F_get_extension_name(m, v1744)
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L1
	} else {
		goto L581
	}
L25:
	;
	v1675 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v1676 = F_SearchSysCache1(m, int32(83), v1675)
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L1
	} else {
		goto L558
	}
L26:
	;
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1650 = F_GetForeignServerExtended(m, v1649, l3)
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L1
	} else {
		goto L551
	}
L27:
	;
	if v27 != int32(2328) {
		goto L14
	} else {
		goto L543
	}
L28:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1588 = F_get_tablespace_name(m, v1587)
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		goto L1
	} else {
		goto L529
	}
L29:
	;
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1553 = F_get_database_name(m, v1552)
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L1
	} else {
		goto L515
	}
L30:
	;
	v1489 = F_table_open(m, int32(1261), int32(1))
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L1
	} else {
		goto L498
	}
L31:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1469 = F_GetUserNameFromId(m, v1468, l3)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L1
	} else {
		goto L490
	}
L32:
	;
	v1416 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v1417 = F_SearchSysCache1(m, int32(74), v1416)
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L1
	} else {
		goto L473
	}
L33:
	;
	if v27 != int32(3764) {
		goto L14
	} else {
		goto L455
	}
L34:
	;
	v1308 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v1309 = F_SearchSysCache1(m, int32(76), v1308)
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L1
	} else {
		goto L438
	}
L35:
	;
	v1255 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v1256 = F_SearchSysCache1(m, int32(78), v1255)
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L1
	} else {
		goto L421
	}
L36:
	;
	if v27 != int32(3381) {
		goto L14
	} else {
		goto L403
	}
L37:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1163 = F_get_namespace_name_or_temp(m, v1162)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L391
	}
L38:
	;
	v1109 = F_table_open(m, int32(2620), int32(1))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L1
	} else {
		goto L373
	}
L39:
	;
	v1054 = F_table_open(m, int32(2618), int32(1))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L1
	} else {
		goto L355
	}
L40:
	;
	v951 = F_table_open(m, int32(2603), int32(1))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L329
	}
L41:
	;
	v848 = F_table_open(m, int32(2602), int32(1))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L303
	}
L42:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v810 = F_get_am_name(m, v809)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L291
	}
L43:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_getOpFamilyIdentity(m, v15+int32(1168), v806, l1, l3)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L290
	}
L44:
	;
	v724 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v725 = F_SearchSysCache1(m, int32(14), v724)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L268
	}
L45:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v642 = F_format_operator_extended(m, v640, int32(3))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L241
	}
L46:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v610 = F_LargeObjectExists(m, v609)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L235
	}
L47:
	;
	v564 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v565 = F_SearchSysCache1(m, int32(36), v564)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L219
	}
L48:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_GetAttrDefaultColumnAddress(m, v15+int32(1056), v528)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L209
	}
L49:
	;
	v474 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v475 = F_SearchSysCache1(m, int32(20), v474)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L192
	}
L50:
	;
	v385 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v386 = F_SearchSysCache1(m, int32(19), v385)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L166
	}
L51:
	;
	v332 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v333 = F_SearchSysCache1(m, int32(16), v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L149
	}
L52:
	;
	v261 = F_table_open(m, int32(2605), int32(1))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L129
	}
L53:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v241 = F_format_type_extended(m, v238, int32(-1), int32(12))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L124
	}
L54:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v120 = F_format_procedure_extended(m, v118, int32(3))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L99
	}
L55:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v84 == int32(0) {
		goto L13
	} else {
		goto L83
	}
L56:
	;
	if v27 <= int32(2327) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	if v27 <= int32(3575) {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	switch v27 - int32(1213) {
	case 0:
		goto L28
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 35, 36, 37, 38, 39, 40, 41, 43, 44, 45:
		goto L14
	case 34:
		goto L53
	case 42:
		goto L54
	case 46:
		goto L55
	case 47:
		goto L31
	case 48:
		goto L30
	case 49:
		goto L29
	default:
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	switch v27 - int32(2601) {
	case 0:
		goto L42
	case 1:
		goto L41
	case 2:
		goto L40
	case 3:
		goto L48
	case 4:
		goto L52
	case 5:
		goto L50
	case 6:
		goto L49
	case 7, 8, 9, 10, 13, 18:
		goto L14
	case 11:
		goto L47
	case 12:
		goto L46
	case 14:
		goto L37
	case 15:
		goto L44
	case 16:
		goto L45
	case 17:
		goto L39
	case 19:
		goto L38
	default:
		goto L27
	}
L62:
	;
	switch v27 - int32(1417) {
	case 0:
		goto L26
	case 1:
		goto L25
	default:
		goto L15
	}
L63:
	;
	if v27 <= int32(3380) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	if v27 <= int32(_a_F_getObjectIdentityParts_0) {
		goto L79
	} else {
		goto L80
	}
L66:
	;
	if v27 == int32(2753) {
		goto L43
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	switch v27 - int32(3456) {
	case 0:
		goto L51
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		goto L14
	case 10:
		goto L23
	default:
		goto L36
	}
L69:
	;
	if v27 == int32(3079) {
		goto L24
	} else {
		goto L70
	}
L70:
	;
	if v27 != int32(3256) {
		goto L14
	} else {
		goto L71
	}
L71:
	;
	v50 = F_table_open(m, int32(3256), int32(1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v55 = F_get_catalog_object_by_oid_extended(m, v50, int32(1), v53, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v55 != 0 {
		goto L21
	} else {
		goto L74
	}
L74:
	;
	if l3 != 0 {
		v2121 = v50
		goto L12
	} else {
		goto L75
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+800)) = v61
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_1), v15+int32(800))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_3), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	switch v27 - int32(3576) {
	case 0:
		goto L16
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:
		goto L14
	case 24:
		goto L34
	case 25:
		goto L35
	case 26:
		goto L32
	default:
		goto L33
	}
L80:
	;
	goto L81
L81:
	;
	switch v27 - int32(_a_F_getObjectIdentityParts_5) {
	case 0:
		goto L17
	case 1, 2, 3, 5:
		goto L14
	case 4:
		goto L20
	case 6:
		goto L18
	default:
		goto L82
	}
L82:
	;
	switch v27 - int32(_a_F_getObjectIdentityParts_6) {
	case 0:
		goto L19
	default:
		goto L14
	case 6:
		goto L22
	}
L83:
	;
	v89 = F_get_attname(m, v83, base.I32_extend16_s(v84), l3)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if v89 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v91 = int32(0)
	goto L87
L86:
	;
	v91 = l3
	goto L87
L87:
	;
	if v91 != 0 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_getRelationIdentity(m, v15+int32(1168), v94, l1, l3)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	if l1 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v97 == int32(0) {
		goto L9
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if v89 == int32(0) {
		goto L9
	} else {
		goto L94
	}
L93:
	;
	goto L92
L94:
	;
	v102 = F_quote_identifier(m, v89)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v102
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_7), v15+int32(32))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L97
	}
L97:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v115 = F_lappend(m, v114, v89)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v115
	goto L9
L99:
	;
	if v120 == int32(0) {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	F_appendStringInfoString(m, v15+int32(1168), v120)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L102
	}
L102:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v131 = m.G0
	v133 = v131 - int32(32)
	m.G0 = v133
	v137 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(v130))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L104
	}
L103:
	;
	m.G0 = v133 + int32(32)
	goto L9
L104:
	;
	if v137 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	if l3 != 0 {
		goto L103
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+22)))
	v156 = v154 + v155
	v157 = int32(*(*int16)(unsafe.Add(mBase, uint32(v156)+104)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v156)+68))
	v159 = F_get_namespace_name_or_temp(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L112
	}
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v130
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_8), v133)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_9), int32(417), int32(_a_F_getObjectIdentityParts_10))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+28)) = v159
	v164 = F_pstrdup(m, v156+int32(4))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+24)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v133)+16)) = v164
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v133)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+20)) = v168
	v174 = F_list_make2_impl(m, v133+int32(20), v133+int32(16))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v174
	v177 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v177
	if v177 < v157 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v190 = v177
	v192 = int32(0)
	goto L118
L116:
	;
	goto L117
L117:
	;
	F_ReleaseCatCache(m, v137)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L123
	}
L118:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v156+int32(136)+v190<<(uint(int32(2))%32))))
	v201 = F_format_type_be_qualified(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L120
	}
L119:
	;
	goto L117
L120:
	;
	v203 = F_lappend(m, v192, v201)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v203
	v207 = v190 + int32(1)
	if v207 != v157 {
		v190 = v207
		v192 = v203
		goto L118
	} else {
		goto L122
	}
L122:
	;
	goto L119
L123:
	;
	goto L103
L124:
	;
	if v241 == int32(0) {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	F_appendStringInfoString(m, v15+int32(1168), v241)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1164)) = v241
	v256 = F_list_make1_impl(m, int32(1), v15+int32(44))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v256
	goto L9
L129:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v266 = F_get_catalog_object_by_oid_extended(m, v261, int32(1), v264, int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	if v266 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if l3 != 0 {
		v2121 = v261
		goto L12
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v266)+16))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+22)))
	v288 = v286 + v287
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	v290 = F_format_type_be_qualified(m, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L138
	}
L134:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v274
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_11), v15+int32(48))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_12), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288)+8))
	v293 = F_format_type_be_qualified(m, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v290
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_13), v15-int32(-64))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	if l1 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	v305 = F_format_type_be_qualified(m, v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	F_relation_close(m, v261, int32(1))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L148
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1160)) = v305
	v312 = F_list_make1_impl(m, int32(1), v15+int32(60))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v312
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v288)+8))
	v316 = F_format_type_be_qualified(m, v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1156)) = v316
	v323 = F_list_make1_impl(m, int32(1), v15+int32(56))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v323
	goto L143
L148:
	;
	goto L9
L149:
	;
	if v333 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v333)+16))
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355)+22)))
	v357 = v355 + v356
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+68))
	v359 = F_get_namespace_name_or_temp(m, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L157
	}
L153:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v341
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_14), v15+int32(80))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_15), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	v362 = v357 + int32(4)
	v363 = F_quote_qualified_identifier(m, v359, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_appendStringInfoString(m, v15+int32(1168), v363)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	if l1 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1152)) = v359
	v368 = F_pstrdup(m, v362)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	F_ReleaseCatCache(m, v333)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L165
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1148)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = v368
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1152))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+92)) = v372
	v378 = F_list_make2_impl(m, v15+int32(92), v15+int32(88))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v378
	goto L162
L165:
	;
	goto L9
L166:
	;
	if v386 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v386)+16))
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+22)))
	v408 = v406 + v407
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)+80))
	if v409 != 0 {
		goto L175
	} else {
		goto L176
	}
L170:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v394
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_16), v15+int32(96))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_17), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	F_ReleaseCatCache(m, v386)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L191
	}
L175:
	;
	v411 = v408 + int32(4)
	v412 = F_quote_identifier(m, v411)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1056)) = int32(1247)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v408)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1064)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1060)) = v438
	v443 = v408 + int32(4)
	v444 = F_quote_identifier(m, v443)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L185
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v412
	v416 = v15 + int32(1168)
	F_appendStringInfo(m, v416, int32(_a_F_getObjectIdentityParts_18), v15+int32(128))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v408)+80))
	F_getRelationIdentity(m, v416, v422, l1, int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	if l1 == int32(0) {
		goto L174
	} else {
		goto L181
	}
L181:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v429 = F_pstrdup(m, v411)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v431 = F_lappend(m, v428, v429)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v431
	F_ReleaseCatCache(m, v386)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	goto L9
L185:
	;
	v449 = F_getObjectIdentityParts(m, v15+int32(1056), l1, l2, int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v444
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_19), v15+int32(112))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	if l1 == int32(0) {
		goto L174
	} else {
		goto L188
	}
L188:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v463 = F_pstrdup(m, v443)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v465 = F_lappend(m, v462, v463)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v465
	goto L174
L191:
	;
	goto L9
L192:
	;
	if v475 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v475)+16))
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+22)))
	v499 = v497 + v498
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)+68))
	v501 = F_get_namespace_name_or_temp(m, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L200
	}
L196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v483
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_20), v15+int32(144))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_21), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	v504 = v499 + int32(4)
	v505 = F_quote_qualified_identifier(m, v501, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_appendStringInfoString(m, v15+int32(1168), v505)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	if l1 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1144)) = v501
	v510 = F_pstrdup(m, v504)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	F_ReleaseCatCache(m, v475)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L208
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1140)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v510
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1144))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+156)) = v514
	v520 = F_list_make2_impl(m, v15+int32(156), v15+int32(152))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v520
	goto L205
L208:
	;
	goto L9
L209:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1060))
	if v531 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	if l3 != 0 {
		goto L9
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v553 = F_getObjectIdentityParts(m, v15+int32(1056), l1, l2, int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L217
	}
L213:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v538
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_22), v15+int32(160))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_23), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v553
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_24), v15+int32(176))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	goto L9
L219:
	;
	if v565 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v565)+16))
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587)+22)))
	v591 = v587 + v588 + int32(4)
	v592 = F_quote_identifier(m, v591)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L227
	}
L223:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = v573
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_25), v15+int32(192))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_26), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L227:
	;
	F_appendStringInfoString(m, v15+int32(1168), v592)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	if l1 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v596 = F_pstrdup(m, v591)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	F_ReleaseCatCache(m, v565)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L234
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+200)) = v596
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1136)) = v596
	v603 = F_list_make1_impl(m, int32(1), v15+int32(200))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v603
	goto L231
L234:
	;
	goto L9
L235:
	;
	if v610 == int32(0) {
		goto L9
	} else {
		goto L236
	}
L236:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = v614
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_27), v15+int32(224))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L238
	}
L238:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v625
	v630 = F_psprintf(m, int32(_a_F_getObjectIdentityParts_27), v15+int32(208))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+204)) = v630
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1132)) = v630
	v637 = F_list_make1_impl(m, int32(1), v15+int32(204))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v637
	goto L9
L241:
	;
	if v642 == int32(0) {
		goto L9
	} else {
		goto L242
	}
L242:
	;
	F_appendStringInfoString(m, v15+int32(1168), v642)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L244
	}
L244:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v653 = m.G0
	v655 = v653 - int32(32)
	m.G0 = v655
	v659 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(v652))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L246
	}
L245:
	;
	m.G0 = v655 + int32(32)
	goto L9
L246:
	;
	if v659 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	if l3 != 0 {
		goto L245
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v659)+16))
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676)+22)))
	v678 = v676 + v677
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)+68))
	v680 = F_get_namespace_name_or_temp(m, v679)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L254
	}
L250:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v655))) = v652
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_28), v655)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_9), int32(825), int32(_a_F_getObjectIdentityParts_29))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v655)+28)) = v680
	v685 = F_pstrdup(m, v678+int32(4))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v655)+24)) = v685
	*(*int32)(unsafe.Add(mBase, uint32(v655)+16)) = v685
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v655)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v655)+20)) = v689
	v695 = F_list_make2_impl(m, v655+int32(20), v655+int32(16))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v695
	v698 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v698
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v678)+80))
	if v701 != 0 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v703 = F_format_type_be_qualified(m, v701)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L260
	}
L258:
	;
	v708 = v698
	goto L259
L259:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v678)+84))
	if v709 != 0 {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	v705 = F_lappend(m, int32(0), v703)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v705
	v708 = v705
	goto L259
L262:
	;
	v710 = F_format_type_be_qualified(m, v709)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	F_ReleaseCatCache(m, v659)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L267
	}
L265:
	;
	v712 = F_lappend(m, v708, v710)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v712
	goto L264
L267:
	;
	goto L245
L268:
	;
	if v725 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v725)+16))
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v745)+22)))
	v747 = v745 + v746
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)+72))
	v749 = F_get_namespace_name_or_temp(m, v748)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L276
	}
L272:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v733
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_30), v15+int32(240))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_31), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	v752 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v747)+4)))
	v753 = F_SearchSysCache1(m, int32(2), v752)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	if v753 == int32(0) {
		goto L11
	} else {
		goto L278
	}
L278:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v753)+16))
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757)+22)))
	v760 = v747 + int32(8)
	v761 = F_quote_qualified_identifier(m, v749, v760)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	v765 = v757 + v758 + int32(4)
	v766 = F_quote_identifier(m, v765)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+276)) = v766
	*(*int32)(unsafe.Add(mBase, uint32(v15)+272)) = v761
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_32), v15+int32(272))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	if l1 != 0 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v777 = F_pstrdup(m, v765)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	F_ReleaseCatCache(m, v753)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L288
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1124)) = v749
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1128)) = v777
	v781 = F_pstrdup(m, v760)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1120)) = v781
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v781
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1128))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v785
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1124))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v787
	v795 = F_list_make3_impl(m, v15+int32(268), v15+int32(264), v15+int32(260))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v795
	goto L284
L288:
	;
	F_ReleaseCatCache(m, v725)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	goto L9
L290:
	;
	goto L9
L291:
	;
	if v810 == int32(0) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	v832 = F_quote_identifier(m, v810)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L299
	}
L295:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+288)) = v818
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_33), v15+int32(288))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_34), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L299:
	;
	F_appendStringInfoString(m, v15+int32(1168), v832)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+300)) = v810
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1116)) = v810
	v843 = F_list_make1_impl(m, int32(1), v15+int32(300))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v843
	goto L9
L303:
	;
	v851 = v15 + int32(1056)
	v855 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	F_ScanKeyInit(m, v851, int32(1), int32(3), int32(184), v855)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	v859 = int32(1)
	v862 = F_systable_beginscan(m, v848, int32(2756), v859, int32(0), v859, v851)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L306
	}
L305:
	;
	F_systable_endscan(m, v862)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L327
	}
L306:
	;
	v864 = F_systable_getnext(m, v862)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	if v864 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	if l3 != 0 {
		goto L305
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v864)+16))
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v884)+22)))
	v887 = v15 + int32(1040)
	F_initStringInfo(m, v887)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L1
	} else {
		goto L315
	}
L311:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+304)) = v872
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_35), v15+int32(304))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_36), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L315:
	;
	v890 = v884 + v885
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v890)+4))
	F_getOpFamilyIdentity(m, v887, v891, l1, int32(0))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v890)+8))
	v896 = F_format_type_be_qualified(m, v895)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v890)+12))
	v899 = F_format_type_be_qualified(m, v898)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	if l1 != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v902 = int32(*(*int16)(unsafe.Add(mBase, uint32(v890)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+352)) = v902
	v907 = F_psprintf(m, int32(_a_F_getObjectIdentityParts_37), v15+int32(352))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	v924 = int32(*(*int16)(unsafe.Add(mBase, uint32(v890)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+320)) = v924
	*(*int32)(unsafe.Add(mBase, uint32(v15)+324)) = v896
	*(*int32)(unsafe.Add(mBase, uint32(v15)+328)) = v899
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1040))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+332)) = v928
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_38), v15+int32(320))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L325
	}
L322:
	;
	v909 = F_lappend(m, v901, v907)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v909
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1032)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1036)) = v896
	*(*int32)(unsafe.Add(mBase, uint32(v15)+348)) = v896
	*(*int32)(unsafe.Add(mBase, uint32(v15)+344)) = v899
	v920 = F_list_make2_impl(m, v15+int32(348), v15+int32(344))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v920
	goto L321
L325:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1040))
	F_pfree(m, v937)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	goto L305
L327:
	;
	F_relation_close(m, v848, int32(1))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	goto L9
L329:
	;
	v954 = v15 + int32(1056)
	v958 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	F_ScanKeyInit(m, v954, int32(1), int32(3), int32(184), v958)
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	v962 = int32(1)
	v965 = F_systable_beginscan(m, v951, int32(2757), v962, int32(0), v962, v954)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L332
	}
L331:
	;
	F_systable_endscan(m, v965)
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L1
	} else {
		goto L353
	}
L332:
	;
	v967 = F_systable_getnext(m, v965)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	if v967 == int32(0) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	if l3 != 0 {
		goto L331
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v967)+16))
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987)+22)))
	v990 = v15 + int32(1040)
	F_initStringInfo(m, v990)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L1
	} else {
		goto L341
	}
L337:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+368)) = v975
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_39), v15+int32(368))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_40), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L341:
	;
	v993 = v987 + v988
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v993)+4))
	F_getOpFamilyIdentity(m, v990, v994, l1, int32(0))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v993)+8))
	v999 = F_format_type_be_qualified(m, v998)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v993)+12))
	v1002 = F_format_type_be_qualified(m, v1001)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L1
	} else {
		goto L344
	}
L344:
	;
	if l1 != 0 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1005 = int32(*(*int16)(unsafe.Add(mBase, uint32(v993)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+416)) = v1005
	v1010 = F_psprintf(m, int32(_a_F_getObjectIdentityParts_37), v15+int32(416))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	v1027 = int32(*(*int16)(unsafe.Add(mBase, uint32(v993)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+384)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v15)+388)) = v999
	*(*int32)(unsafe.Add(mBase, uint32(v15)+392)) = v1002
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1040))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+396)) = v1031
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_41), v15+int32(384))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L351
	}
L348:
	;
	v1012 = F_lappend(m, v1004, v1010)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L349
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1012
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1024)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1028)) = v999
	*(*int32)(unsafe.Add(mBase, uint32(v15)+412)) = v999
	*(*int32)(unsafe.Add(mBase, uint32(v15)+408)) = v1002
	v1023 = F_list_make2_impl(m, v15+int32(412), v15+int32(408))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L1
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1023
	goto L347
L351:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1040))
	F_pfree(m, v1040)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	goto L331
L353:
	;
	F_relation_close(m, v951, int32(1))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	goto L9
L355:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1059 = F_get_catalog_object_by_oid_extended(m, v1054, int32(1), v1057, int32(0))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	if v1059 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	if l3 != 0 {
		v2121 = v1054
		goto L12
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+16))
	v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+22)))
	v1081 = v1079 + v1080
	v1083 = v1081 + int32(4)
	v1084 = F_quote_identifier(m, v1083)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L1
	} else {
		goto L364
	}
L360:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+432)) = v1067
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_42), v15+int32(432))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_43), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+448)) = v1084
	v1088 = v15 + int32(1168)
	F_appendStringInfo(m, v1088, int32(_a_F_getObjectIdentityParts_18), v15+int32(448))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+68))
	F_getRelationIdentity(m, v1088, v1094, l1, int32(0))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L1
	} else {
		goto L366
	}
L366:
	;
	if l1 != 0 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1099 = F_pstrdup(m, v1083)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	F_relation_close(m, v1054, int32(1))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L1
	} else {
		goto L372
	}
L370:
	;
	v1101 = F_lappend(m, v1098, v1099)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1101
	goto L369
L372:
	;
	goto L9
L373:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1114 = F_get_catalog_object_by_oid_extended(m, v1109, int32(1), v1112, int32(0))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L1
	} else {
		goto L374
	}
L374:
	;
	if v1114 == int32(0) {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	if l3 != 0 {
		v2121 = v1109
		goto L12
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+16))
	v1135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1134)+22)))
	v1136 = v1134 + v1135
	v1138 = v1136 + int32(12)
	v1139 = F_quote_identifier(m, v1138)
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L1
	} else {
		goto L382
	}
L378:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+464)) = v1122
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_44), v15+int32(464))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_45), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+480)) = v1139
	v1143 = v15 + int32(1168)
	F_appendStringInfo(m, v1143, int32(_a_F_getObjectIdentityParts_18), v15+int32(480))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+4))
	F_getRelationIdentity(m, v1143, v1149, l1, int32(0))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L1
	} else {
		goto L384
	}
L384:
	;
	if l1 != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1154 = F_pstrdup(m, v1138)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L1
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	F_relation_close(m, v1109, int32(1))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L1
	} else {
		goto L390
	}
L388:
	;
	v1156 = F_lappend(m, v1153, v1154)
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1156
	goto L387
L390:
	;
	goto L9
L391:
	;
	if v1163 == int32(0) {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	v1185 = F_quote_identifier(m, v1163)
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L1
	} else {
		goto L399
	}
L395:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+496)) = v1171
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_46), v15+int32(496))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_47), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L399:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1185)
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L401
	}
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+508)) = v1163
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1020)) = v1163
	v1196 = F_list_make1_impl(m, int32(1), v15+int32(508))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1196
	goto L9
L403:
	;
	v1202 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v1203 = F_SearchSysCache1(m, int32(64), v1202)
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	if v1203 == int32(0) {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L408
	}
L406:
	;
	goto L407
L407:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1203)+16))
	v1226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1225)+22)))
	v1227 = v1225 + v1226
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+72))
	v1229 = F_get_namespace_name_or_temp(m, v1228)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L1
	} else {
		goto L412
	}
L408:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L1
	} else {
		goto L409
	}
L409:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+512)) = v1211
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_48), v15+int32(512))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_49), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L412:
	;
	v1232 = v1227 + int32(8)
	v1233 = F_quote_qualified_identifier(m, v1229, v1232)
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1233)
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	if l1 != 0 {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1016)) = v1229
	v1238 = F_pstrdup(m, v1232)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L1
	} else {
		goto L418
	}
L416:
	;
	goto L417
L417:
	;
	F_ReleaseCatCache(m, v1203)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L1
	} else {
		goto L420
	}
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1012)) = v1238
	*(*int32)(unsafe.Add(mBase, uint32(v15)+520)) = v1238
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1016))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+524)) = v1242
	v1248 = F_list_make2_impl(m, v15+int32(524), v15+int32(520))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L1
	} else {
		goto L419
	}
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1248
	goto L417
L420:
	;
	goto L9
L421:
	;
	if v1256 == int32(0) {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L425
	}
L423:
	;
	goto L424
L424:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+16))
	v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1278)+22)))
	v1280 = v1278 + v1279
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+68))
	v1282 = F_get_namespace_name_or_temp(m, v1281)
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L1
	} else {
		goto L429
	}
L425:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+528)) = v1264
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_50), v15+int32(528))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L1
	} else {
		goto L427
	}
L427:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_51), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L429:
	;
	v1285 = v1280 + int32(4)
	v1286 = F_quote_qualified_identifier(m, v1282, v1285)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1286)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	if l1 != 0 {
		goto L432
	} else {
		goto L433
	}
L432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1008)) = v1282
	v1291 = F_pstrdup(m, v1285)
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L1
	} else {
		goto L435
	}
L433:
	;
	goto L434
L434:
	;
	F_ReleaseCatCache(m, v1256)
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L1
	} else {
		goto L437
	}
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1004)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v15)+536)) = v1291
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1008))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+540)) = v1295
	v1301 = F_list_make2_impl(m, v15+int32(540), v15+int32(536))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L1
	} else {
		goto L436
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1301
	goto L434
L437:
	;
	goto L9
L438:
	;
	if v1309 == int32(0) {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L442
	}
L440:
	;
	goto L441
L441:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1309)+16))
	v1332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1331)+22)))
	v1333 = v1331 + v1332
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+68))
	v1335 = F_get_namespace_name_or_temp(m, v1334)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L1
	} else {
		goto L446
	}
L442:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L1
	} else {
		goto L443
	}
L443:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+544)) = v1317
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_52), v15+int32(544))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L1
	} else {
		goto L444
	}
L444:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_53), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L446:
	;
	v1338 = v1333 + int32(4)
	v1339 = F_quote_qualified_identifier(m, v1335, v1338)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L1
	} else {
		goto L447
	}
L447:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1339)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	if l1 != 0 {
		goto L449
	} else {
		goto L450
	}
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1000)) = v1335
	v1344 = F_pstrdup(m, v1338)
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L1
	} else {
		goto L452
	}
L450:
	;
	goto L451
L451:
	;
	F_ReleaseCatCache(m, v1309)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L1
	} else {
		goto L454
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+996)) = v1344
	*(*int32)(unsafe.Add(mBase, uint32(v15)+552)) = v1344
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1000))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+556)) = v1348
	v1354 = F_list_make2_impl(m, v15+int32(556), v15+int32(552))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1354
	goto L451
L454:
	;
	goto L9
L455:
	;
	v1363 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v1364 = F_SearchSysCache1(m, int32(80), v1363)
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	if v1364 == int32(0) {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L460
	}
L458:
	;
	goto L459
L459:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1364)+16))
	v1387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1386)+22)))
	v1388 = v1386 + v1387
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1388)+68))
	v1390 = F_get_namespace_name_or_temp(m, v1389)
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L1
	} else {
		goto L464
	}
L460:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L1
	} else {
		goto L461
	}
L461:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+560)) = v1372
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_54), v15+int32(560))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L1
	} else {
		goto L462
	}
L462:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_55), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L1
	} else {
		goto L463
	}
L463:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L464:
	;
	v1393 = v1388 + int32(4)
	v1394 = F_quote_qualified_identifier(m, v1390, v1393)
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1394)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L1
	} else {
		goto L466
	}
L466:
	;
	if l1 != 0 {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+992)) = v1390
	v1399 = F_pstrdup(m, v1393)
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L1
	} else {
		goto L470
	}
L468:
	;
	goto L469
L469:
	;
	F_ReleaseCatCache(m, v1364)
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L1
	} else {
		goto L472
	}
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+988)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v15)+568)) = v1399
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v15)+992))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+572)) = v1403
	v1409 = F_list_make2_impl(m, v15+int32(572), v15+int32(568))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L1
	} else {
		goto L471
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1409
	goto L469
L472:
	;
	goto L9
L473:
	;
	if v1417 == int32(0) {
		goto L474
	} else {
		goto L475
	}
L474:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L477
	}
L475:
	;
	goto L476
L476:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+16))
	v1440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1439)+22)))
	v1441 = v1439 + v1440
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+68))
	v1443 = F_get_namespace_name_or_temp(m, v1442)
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L1
	} else {
		goto L481
	}
L477:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L1
	} else {
		goto L478
	}
L478:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+576)) = v1425
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_56), v15+int32(576))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L479
	}
L479:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_57), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L481:
	;
	v1446 = v1441 + int32(4)
	v1447 = F_quote_qualified_identifier(m, v1443, v1446)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1447)
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L1
	} else {
		goto L483
	}
L483:
	;
	if l1 != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+984)) = v1443
	v1452 = F_pstrdup(m, v1446)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L1
	} else {
		goto L487
	}
L485:
	;
	goto L486
L486:
	;
	F_ReleaseCatCache(m, v1417)
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L1
	} else {
		goto L489
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+980)) = v1452
	*(*int32)(unsafe.Add(mBase, uint32(v15)+580)) = v1452
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v15)+984))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+584)) = v1456
	v1462 = F_list_make2_impl(m, v15+int32(584), v15+int32(580))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1462
	goto L486
L489:
	;
	goto L9
L490:
	;
	if v1469 == int32(0) {
		goto L9
	} else {
		goto L491
	}
L491:
	;
	if l1 != 0 {
		goto L492
	} else {
		goto L493
	}
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+588)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v15)+976)) = v1469
	v1478 = F_list_make1_impl(m, int32(1), v15+int32(588))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L1
	} else {
		goto L495
	}
L493:
	;
	goto L494
L494:
	;
	v1483 = F_quote_identifier(m, v1469)
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L1
	} else {
		goto L496
	}
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1478
	goto L494
L496:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1483)
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	goto L9
L498:
	;
	v1492 = v15 + int32(1056)
	v1496 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	F_ScanKeyInit(m, v1492, int32(1), int32(3), int32(184), v1496)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	v1500 = int32(1)
	v1503 = F_systable_beginscan(m, v1489, int32(_a_F_getObjectIdentityParts_58), v1500, int32(0), v1500, v1492)
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L1
	} else {
		goto L501
	}
L500:
	;
	F_systable_endscan(m, v1503)
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L1
	} else {
		goto L513
	}
L501:
	;
	v1505 = F_systable_getnext(m, v1503)
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L1
	} else {
		goto L502
	}
L502:
	;
	if v1505 == int32(0) {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	if l3 != 0 {
		goto L500
	} else {
		goto L506
	}
L504:
	;
	goto L505
L505:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v1505)+16))
	v1526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1525)+22)))
	v1527 = v1525 + v1526
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+8))
	v1530 = F_GetUserNameFromId(m, v1528, int32(0))
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L1
	} else {
		goto L510
	}
L506:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L1
	} else {
		goto L507
	}
L507:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+592)) = v1513
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_59), v15+int32(592))
	mBase = m.M
	v1519 = m.ExcPending
	if v1519 != 0 {
		goto L1
	} else {
		goto L508
	}
L508:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_60), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		goto L1
	} else {
		goto L509
	}
L509:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L510:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+4))
	v1534 = F_GetUserNameFromId(m, v1532, int32(0))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+612)) = v1534
	*(*int32)(unsafe.Add(mBase, uint32(v15)+608)) = v1530
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_61), v15+int32(608))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	goto L500
L513:
	;
	F_relation_close(m, v1489, int32(1))
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	goto L9
L515:
	;
	if v1553 == int32(0) {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L519
	}
L517:
	;
	goto L518
L518:
	;
	if l1 != 0 {
		goto L523
	} else {
		goto L524
	}
L519:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+624)) = v1561
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_62), v15+int32(624))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_63), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L1
	} else {
		goto L522
	}
L522:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L523:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+636)) = v1553
	*(*int32)(unsafe.Add(mBase, uint32(v15)+972)) = v1553
	v1578 = F_list_make1_impl(m, int32(1), v15+int32(636))
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L1
	} else {
		goto L526
	}
L524:
	;
	goto L525
L525:
	;
	v1583 = F_quote_identifier(m, v1553)
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L1
	} else {
		goto L527
	}
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1578
	goto L525
L527:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1583)
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L1
	} else {
		goto L528
	}
L528:
	;
	goto L9
L529:
	;
	if v1588 == int32(0) {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L533
	}
L531:
	;
	goto L532
L532:
	;
	if l1 != 0 {
		goto L537
	} else {
		goto L538
	}
L533:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L1
	} else {
		goto L534
	}
L534:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+640)) = v1596
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_64), v15+int32(640))
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L1
	} else {
		goto L535
	}
L535:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_65), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L537:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+644)) = v1588
	*(*int32)(unsafe.Add(mBase, uint32(v15)+968)) = v1588
	v1613 = F_list_make1_impl(m, int32(1), v15+int32(644))
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L1
	} else {
		goto L540
	}
L538:
	;
	goto L539
L539:
	;
	v1618 = F_quote_identifier(m, v1588)
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L1
	} else {
		goto L541
	}
L540:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1613
	goto L539
L541:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1618)
	mBase = m.M
	v1621 = m.ExcPending
	if v1621 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	goto L9
L543:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1625 = F_GetForeignDataWrapperExtended(m, v1624, l3)
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L1
	} else {
		goto L544
	}
L544:
	;
	if v1625 == int32(0) {
		goto L9
	} else {
		goto L545
	}
L545:
	;
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+8))
	v1632 = F_quote_identifier(m, v1631)
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L1
	} else {
		goto L546
	}
L546:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1632)
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L1
	} else {
		goto L547
	}
L547:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L548
	}
L548:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+8))
	v1639 = F_pstrdup(m, v1638)
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L1
	} else {
		goto L549
	}
L549:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+648)) = v1639
	*(*int32)(unsafe.Add(mBase, uint32(v15)+964)) = v1639
	v1646 = F_list_make1_impl(m, int32(1), v15+int32(648))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1646
	goto L9
L551:
	;
	if v1650 == int32(0) {
		goto L9
	} else {
		goto L552
	}
L552:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1650)+12))
	v1657 = F_quote_identifier(m, v1656)
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1657)
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L1
	} else {
		goto L554
	}
L554:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L555
	}
L555:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1650)+12))
	v1664 = F_pstrdup(m, v1663)
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L1
	} else {
		goto L556
	}
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+652)) = v1664
	*(*int32)(unsafe.Add(mBase, uint32(v15)+960)) = v1664
	v1671 = F_list_make1_impl(m, int32(1), v15+int32(652))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L1
	} else {
		goto L557
	}
L557:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1671
	goto L9
L558:
	;
	if v1676 == int32(0) {
		goto L559
	} else {
		goto L560
	}
L559:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L562
	}
L560:
	;
	goto L561
L561:
	;
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v1676)+16))
	v1697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1696)+22)))
	v1698 = v1696 + v1697
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v1698)+4))
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1698)+8))
	v1701 = F_GetForeignServer(m, v1700)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L1
	} else {
		goto L566
	}
L562:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+656)) = v1684
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_66), v15+int32(656))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_67), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L1
	} else {
		goto L565
	}
L565:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L566:
	;
	F_ReleaseCatCache(m, v1676)
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L1
	} else {
		goto L567
	}
L567:
	;
	if v1699 != 0 {
		goto L568
	} else {
		goto L569
	}
L568:
	;
	v1706 = F_GetUserNameFromId(m, v1699, int32(0))
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L1
	} else {
		goto L571
	}
L569:
	;
	v1709 = int32(_a_F_getObjectIdentityParts_68)
	goto L570
L570:
	;
	if l1 != 0 {
		goto L572
	} else {
		goto L573
	}
L571:
	;
	v1709 = v1706
	goto L570
L572:
	;
	v1710 = F_pstrdup(m, v1709)
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L1
	} else {
		goto L575
	}
L573:
	;
	goto L574
L574:
	;
	v1732 = F_quote_identifier(m, v1709)
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L1
	} else {
		goto L579
	}
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+684)) = v1710
	*(*int32)(unsafe.Add(mBase, uint32(v15)+956)) = v1710
	v1717 = F_list_make1_impl(m, int32(1), v15+int32(684))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L1
	} else {
		goto L576
	}
L576:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1717
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1701)+12))
	v1721 = F_pstrdup(m, v1720)
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+680)) = v1721
	*(*int32)(unsafe.Add(mBase, uint32(v15)+952)) = v1721
	v1728 = F_list_make1_impl(m, int32(1), v15+int32(680))
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1728
	goto L574
L579:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1701)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+676)) = v1734
	*(*int32)(unsafe.Add(mBase, uint32(v15)+672)) = v1732
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_69), v15+int32(672))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	goto L9
L581:
	;
	if v1745 == int32(0) {
		goto L582
	} else {
		goto L583
	}
L582:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L585
	}
L583:
	;
	goto L584
L584:
	;
	v1767 = F_quote_identifier(m, v1745)
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L1
	} else {
		goto L589
	}
L585:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1752 = m.ExcPending
	if v1752 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+752)) = v1753
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_70), v15+int32(752))
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L1
	} else {
		goto L587
	}
L587:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_71), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L589:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1767)
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L1
	} else {
		goto L590
	}
L590:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L591
	}
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+764)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v15)+940)) = v1745
	v1778 = F_list_make1_impl(m, int32(1), v15+int32(764))
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1778
	goto L9
L593:
	;
	if v1783 == int32(0) {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L597
	}
L595:
	;
	goto L596
L596:
	;
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+16))
	v1806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1805)+22)))
	v1810 = F_pstrdup(m, v1805+v1806+int32(4))
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L1
	} else {
		goto L601
	}
L597:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+768)) = v1791
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_72), v15+int32(768))
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_73), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L601:
	;
	v1812 = F_quote_identifier(m, v1810)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L1
	} else {
		goto L602
	}
L602:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1812)
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L1
	} else {
		goto L603
	}
L603:
	;
	if l1 != 0 {
		goto L604
	} else {
		goto L605
	}
L604:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+780)) = v1810
	*(*int32)(unsafe.Add(mBase, uint32(v15)+936)) = v1810
	v1821 = F_list_make1_impl(m, int32(1), v15+int32(780))
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L1
	} else {
		goto L607
	}
L605:
	;
	goto L606
L606:
	;
	F_ReleaseCatCache(m, v1783)
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L1
	} else {
		goto L608
	}
L607:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1821
	goto L606
L608:
	;
	goto L9
L609:
	;
	if v1828 == int32(0) {
		goto L610
	} else {
		goto L611
	}
L610:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L613
	}
L611:
	;
	goto L612
L612:
	;
	v1852 = F_SysCacheGetAttrNotNull(m, int32(44), v1828, int32(2))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L1
	} else {
		goto L617
	}
L613:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L1
	} else {
		goto L614
	}
L614:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+784)) = v1836
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_74), v15+int32(784))
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_75), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L617:
	;
	v1855 = F_text_to_cstring(m, base.I32_wrap_i64(v1852))
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L1
	} else {
		goto L618
	}
L618:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1855)
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L1
	} else {
		goto L619
	}
L619:
	;
	if l1 != 0 {
		goto L620
	} else {
		goto L621
	}
L620:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+796)) = v1855
	*(*int32)(unsafe.Add(mBase, uint32(v15)+932)) = v1855
	v1864 = F_list_make1_impl(m, int32(1), v15+int32(796))
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L1
	} else {
		goto L623
	}
L621:
	;
	goto L622
L622:
	;
	F_ReleaseCatCache(m, v1828)
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L1
	} else {
		goto L624
	}
L623:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1864
	goto L622
L624:
	;
	goto L9
L625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+816)) = v1874
	v1878 = v15 + int32(1168)
	F_appendStringInfo(m, v1878, int32(_a_F_getObjectIdentityParts_18), v15+int32(816))
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1871)+68))
	F_getRelationIdentity(m, v1878, v1884, l1, int32(0))
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		goto L1
	} else {
		goto L627
	}
L627:
	;
	if l1 != 0 {
		goto L628
	} else {
		goto L629
	}
L628:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1889 = F_pstrdup(m, v1873)
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L1
	} else {
		goto L631
	}
L629:
	;
	goto L630
L630:
	;
	F_relation_close(m, v50, int32(1))
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L1
	} else {
		goto L633
	}
L631:
	;
	v1891 = F_lappend(m, v1888, v1889)
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L1
	} else {
		goto L632
	}
L632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1891
	goto L630
L633:
	;
	goto L9
L634:
	;
	if v1898 == int32(0) {
		goto L9
	} else {
		goto L635
	}
L635:
	;
	v1904 = F_quote_identifier(m, v1898)
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1904)
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L1
	} else {
		goto L637
	}
L637:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L638
	}
L638:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+820)) = v1898
	*(*int32)(unsafe.Add(mBase, uint32(v15)+928)) = v1898
	v1915 = F_list_make1_impl(m, int32(1), v15+int32(820))
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1915
	goto L9
L640:
	;
	if v1922 == int32(0) {
		goto L9
	} else {
		goto L641
	}
L641:
	;
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1040))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+832)) = v1926
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1056))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+836)) = v1928
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_76), v15+int32(832))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L1
	} else {
		goto L642
	}
L642:
	;
	if l2 != 0 {
		goto L644
	} else {
		goto L645
	}
L643:
	;
	if l1 != 0 {
		goto L649
	} else {
		goto L650
	}
L644:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+828)) = v1928
	*(*int32)(unsafe.Add(mBase, uint32(v15)+924)) = v1928
	v1942 = F_list_make1_impl(m, int32(1), v15+int32(828))
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L1
	} else {
		goto L647
	}
L645:
	;
	goto L646
L646:
	;
	F_pfree(m, v1928)
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L1
	} else {
		goto L648
	}
L647:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1942
	goto L643
L648:
	;
	goto L643
L649:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+824)) = v1926
	*(*int32)(unsafe.Add(mBase, uint32(v15)+920)) = v1926
	v1952 = F_list_make1_impl(m, int32(1), v15+int32(824))
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L1
	} else {
		goto L652
	}
L650:
	;
	goto L651
L651:
	;
	F_pfree(m, v1926)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L1
	} else {
		goto L653
	}
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1952
	goto L9
L653:
	;
	goto L9
L654:
	;
	if v1959 == int32(0) {
		goto L655
	} else {
		goto L656
	}
L655:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L658
	}
L656:
	;
	goto L657
L657:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+16))
	v1980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1979)+22)))
	v1981 = v1979 + v1980
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v1981)+4))
	v1984 = F_get_publication_name(m, v1982, int32(0))
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L1
	} else {
		goto L662
	}
L658:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+848)) = v1967
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_77), v15+int32(848))
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L1
	} else {
		goto L660
	}
L660:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_78), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L1
	} else {
		goto L661
	}
L661:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L662:
	;
	v1987 = v15 + int32(1168)
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v1981)+8))
	F_getRelationIdentity(m, v1987, v1988, l1, int32(0))
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L1
	} else {
		goto L663
	}
L663:
	;
	v1992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+864)) = v1984
	if v1992 != 0 {
		goto L664
	} else {
		goto L665
	}
L664:
	;
	v1996 = int32(_a_F_getObjectIdentityParts_79)
	goto L666
L665:
	;
	v1996 = int32(_a_F_getObjectIdentityParts_80)
	goto L666
L666:
	;
	F_appendStringInfo(m, v1987, v1996, v15+int32(864))
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L1
	} else {
		goto L667
	}
L667:
	;
	if l2 != 0 {
		goto L668
	} else {
		goto L669
	}
L668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+860)) = v1984
	*(*int32)(unsafe.Add(mBase, uint32(v15)+916)) = v1984
	v2006 = F_list_make1_impl(m, int32(1), v15+int32(860))
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L1
	} else {
		goto L671
	}
L669:
	;
	goto L670
L670:
	;
	F_ReleaseCatCache(m, v1959)
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L1
	} else {
		goto L672
	}
L671:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2006
	goto L670
L672:
	;
	goto L9
L673:
	;
	if v2012 == int32(0) {
		goto L9
	} else {
		goto L674
	}
L674:
	;
	v2018 = F_quote_identifier(m, v2012)
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L1
	} else {
		goto L675
	}
L675:
	;
	F_appendStringInfoString(m, v15+int32(1168), v2018)
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L1
	} else {
		goto L676
	}
L676:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L677
	}
L677:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+876)) = v2012
	*(*int32)(unsafe.Add(mBase, uint32(v15)+912)) = v2012
	v2029 = F_list_make1_impl(m, int32(1), v15+int32(876))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L1
	} else {
		goto L678
	}
L678:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2029
	goto L9
L679:
	;
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2039 = F_get_catalog_object_by_oid_extended(m, v2034, int32(1), v2037, int32(0))
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L1
	} else {
		goto L680
	}
L680:
	;
	if v2039 == int32(0) {
		goto L681
	} else {
		goto L682
	}
L681:
	;
	if l3 != 0 {
		v2121 = v2034
		goto L12
	} else {
		goto L684
	}
L682:
	;
	goto L683
L683:
	;
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v2039)+16))
	v2060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2059)+22)))
	v2061 = v2059 + v2060
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(v2061)+4))
	v2063 = F_format_type_be_qualified(m, v2062)
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L1
	} else {
		goto L688
	}
L684:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L1
	} else {
		goto L685
	}
L685:
	;
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+880)) = v2047
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_81), v15+int32(880))
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_82), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L1
	} else {
		goto L687
	}
L687:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L688:
	;
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v2061)+8))
	v2067 = F_get_language_name(m, v2065, int32(0))
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		goto L1
	} else {
		goto L689
	}
L689:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+900)) = v2067
	*(*int32)(unsafe.Add(mBase, uint32(v15)+896)) = v2063
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_83), v15+int32(896))
	mBase = m.M
	v2077 = m.ExcPending
	if v2077 != 0 {
		goto L1
	} else {
		goto L690
	}
L690:
	;
	if l1 != 0 {
		goto L691
	} else {
		goto L692
	}
L691:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+892)) = v2063
	*(*int32)(unsafe.Add(mBase, uint32(v15)+908)) = v2063
	v2083 = F_list_make1_impl(m, int32(1), v15+int32(892))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L1
	} else {
		goto L694
	}
L692:
	;
	goto L693
L693:
	;
	F_relation_close(m, v2034, int32(1))
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L1
	} else {
		goto L697
	}
L694:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2083
	v2086 = F_pstrdup(m, v2067)
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L1
	} else {
		goto L695
	}
L695:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+888)) = v2086
	*(*int32)(unsafe.Add(mBase, uint32(v15)+904)) = v2086
	v2093 = F_list_make1_impl(m, int32(1), v15+int32(888))
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L1
	} else {
		goto L696
	}
L696:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2093
	goto L693
L697:
	;
	goto L9
L698:
	;
	goto L14
L699:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v2106
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_84), v15)
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L1
	} else {
		goto L700
	}
L700:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_85), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v2115 = m.ExcPending
	if v2115 != 0 {
		goto L1
	} else {
		goto L701
	}
L701:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L702:
	;
	goto L9
L703:
	;
	goto L8
L704:
	;
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v747)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v2130
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_33), v15+int32(256))
	mBase = m.M
	v2136 = m.ExcPending
	if v2136 != 0 {
		goto L1
	} else {
		goto L705
	}
L705:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_86), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v2141 = m.ExcPending
	if v2141 != 0 {
		goto L1
	} else {
		goto L706
	}
L706:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L707:
	;
	v2147 = v15 + int32(1056)
	v2151 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	F_ScanKeyInit(m, v2147, int32(1), int32(3), int32(184), v2151)
	mBase = m.M
	v2153 = m.ExcPending
	if v2153 != 0 {
		goto L1
	} else {
		goto L708
	}
L708:
	;
	v2155 = int32(1)
	v2158 = F_systable_beginscan(m, v2144, int32(828), v2155, int32(0), v2155, v2147)
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L1
	} else {
		goto L710
	}
L709:
	;
	F_systable_endscan(m, v2158)
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L1
	} else {
		goto L744
	}
L710:
	;
	v2160 = F_systable_getnext(m, v2158)
	mBase = m.M
	v2161 = m.ExcPending
	if v2161 != 0 {
		goto L1
	} else {
		goto L711
	}
L711:
	;
	if v2160 == int32(0) {
		goto L712
	} else {
		goto L713
	}
L712:
	;
	if l3 != 0 {
		goto L709
	} else {
		goto L715
	}
L713:
	;
	goto L714
L714:
	;
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v2160)+16))
	v2181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2180)+22)))
	v2182 = v2180 + v2181
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v2182)+4))
	v2185 = F_GetUserNameFromId(m, v2183, int32(0))
	mBase = m.M
	v2186 = m.ExcPending
	if v2186 != 0 {
		goto L1
	} else {
		goto L719
	}
L715:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L1
	} else {
		goto L716
	}
L716:
	;
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+688)) = v2168
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_87), v15+int32(688))
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_88), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L1
	} else {
		goto L718
	}
L718:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L719:
	;
	v2187 = F_quote_identifier(m, v2185)
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L1
	} else {
		goto L720
	}
L720:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+736)) = v2187
	v2191 = v15 + int32(1168)
	F_appendStringInfo(m, v2191, int32(_a_F_getObjectIdentityParts_89), v15+int32(736))
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L1
	} else {
		goto L721
	}
L721:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v2182)+8))
	if v2197 != 0 {
		goto L722
	} else {
		goto L723
	}
L722:
	;
	v2198 = F_get_namespace_name_or_temp(m, v2197)
	mBase = m.M
	v2199 = m.ExcPending
	if v2199 != 0 {
		goto L1
	} else {
		goto L725
	}
L723:
	;
	v2208 = int32(0)
	goto L724
L724:
	;
	v2210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2182)+12)))
	switch v2210 - int32(76) {
	case 0:
		goto L730
	default:
		goto L728
	case 7:
		goto L734
	case 8:
		goto L732
	case 26:
		goto L733
	case 34:
		goto L731
	case 38:
		v2218 = int32(_a_F_getObjectIdentityParts_90)
		goto L729
	}
L725:
	;
	v2200 = F_quote_identifier(m, v2198)
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L1
	} else {
		goto L726
	}
L726:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+720)) = v2200
	F_appendStringInfo(m, v2191, int32(_a_F_getObjectIdentityParts_91), v15+int32(720))
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L1
	} else {
		goto L727
	}
L727:
	;
	v2208 = v2198
	goto L724
L728:
	;
	if l1 == int32(0) {
		goto L709
	} else {
		goto L736
	}
L729:
	;
	F_appendStringInfoString(m, v15+int32(1168), v2218)
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L1
	} else {
		goto L735
	}
L730:
	;
	v2218 = int32(_a_F_getObjectIdentityParts_92)
	goto L729
L731:
	;
	v2218 = int32(_a_F_getObjectIdentityParts_93)
	goto L729
L732:
	;
	v2218 = int32(_a_F_getObjectIdentityParts_94)
	goto L729
L733:
	;
	v2218 = int32(_a_F_getObjectIdentityParts_95)
	goto L729
L734:
	;
	v2218 = int32(_a_F_getObjectIdentityParts_96)
	goto L729
L735:
	;
	goto L728
L736:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+716)) = v2185
	*(*int32)(unsafe.Add(mBase, uint32(v15)+948)) = v2185
	v2231 = F_list_make1_impl(m, int32(1), v15+int32(716))
	mBase = m.M
	v2232 = m.ExcPending
	if v2232 != 0 {
		goto L1
	} else {
		goto L737
	}
L737:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2231
	if v2208 != 0 {
		goto L738
	} else {
		goto L739
	}
L738:
	;
	v2234 = F_lappend(m, v2231, v2208)
	mBase = m.M
	v2235 = m.ExcPending
	if v2235 != 0 {
		goto L1
	} else {
		goto L741
	}
L739:
	;
	goto L740
L740:
	;
	v2237 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2182)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+704)) = v2237
	v2242 = F_psprintf(m, int32(_a_F_getObjectIdentityParts_97), v15+int32(704))
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L1
	} else {
		goto L742
	}
L741:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2234
	goto L740
L742:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+700)) = v2242
	*(*int32)(unsafe.Add(mBase, uint32(v15)+944)) = v2242
	v2249 = F_list_make1_impl(m, int32(1), v15+int32(700))
	mBase = m.M
	v2250 = m.ExcPending
	if v2250 != 0 {
		goto L1
	} else {
		goto L743
	}
L743:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2249
	goto L709
L744:
	;
	F_relation_close(m, v2144, int32(1))
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L1
	} else {
		goto L745
	}
L745:
	;
	goto L9
L746:
	;
	if l1 == int32(0) {
		goto L7
	} else {
		goto L747
	}
L747:
	;
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v2276 != 0 {
		goto L7
	} else {
		goto L748
	}
L748:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L1
	} else {
		goto L749
	}
L749:
	;
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v2281
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1168))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v2283
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_98), v15+int32(16))
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
		goto L1
	} else {
		goto L750
	}
L750:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_99), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v2294 = m.ExcPending
	if v2294 != 0 {
		goto L1
	} else {
		goto L751
	}
L751:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L752:
	;
	v2334 = int32(0)
	goto L6
}
func F_get_object_class_descr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_class_descr[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	m.G0 = v8 + int32(16)
	return v52
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == l0 {
		v48 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_object_class_descr[0])) = v45
	v48 = v45
	goto L2
L8:
	;
	v45 = v21 + int32(_a_F_get_object_class_descr_0)
	goto L7
L9:
	;
	v45 = v21 + int32(_a_F_get_object_class_descr_1)
	goto L7
L10:
	;
	v45 = v21 + int32(_a_F_get_object_class_descr_2)
	goto L7
L11:
	;
	v21 = v18 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_class_descr[1])))
	if l0 != v22 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v45 = v21 + int32(_a_F_get_object_class_descr_3)
	goto L7
L13:
	;
	if v18 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_class_descr[2])))
	if v28 == l0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_class_descr[3])))
	if v30 == l0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_class_descr[4])))
	if v32 == l0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v18 = v18 + int32(4)
	goto L11
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_get_object_class_descr_4), v8)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_get_object_class_descr_5), int32(2827), int32(_a_F_get_object_class_descr_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_new_object_addresses(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v4 = F_palloc(m, int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = int64(137438953472)
		v12 = F_palloc_mul(m, int32(12), int32(32))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = v12
			return v4
		}
	}
}
func F_object_address_comparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v7) < base.Ui32(v6) {
		return int32(-1)
	} else {
		v11 = int32(1)
		if base.Ui32(v6) < base.Ui32(v7) {
			v28 = v11
			return v28
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if base.Ui32(v13) < base.Ui32(v14) {
				return int32(-1)
			} else {
				if base.Ui32(v14) < base.Ui32(v13) {
					v28 = v11
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if base.Ui32(v20) < base.Ui32(v21) {
						v28 = int32(-1)
					} else {
						v28 = base.B2i32(base.Ui32(v21) < base.Ui32(v20))
					}
				}
				return v28
			}
		}
	}
}
