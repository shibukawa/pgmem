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
	v19 = *(*int32)(unsafe.Add(mBase, _consts[433]))
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[433]))
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
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
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v218 int32
	_ = v218
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
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
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
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
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
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
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
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
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1076 int32
	_ = v1076
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1133 int32
	_ = v1133
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
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
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1224 int32
	_ = v1224
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
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
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1330 int32
	_ = v1330
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1385 int32
	_ = v1385
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
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
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1528 int32
	_ = v1528
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1576 int32
	_ = v1576
	var v1581 int32
	_ = v1581
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1611 int32
	_ = v1611
	var v1616 int32
	_ = v1616
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1699 int32
	_ = v1699
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1768 int32
	_ = v1768
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1779 int32
	_ = v1779
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1806 int32
	_ = v1806
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1851 int32
	_ = v1851
	var v1856 int32
	_ = v1856
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1946 int32
	_ = v1946
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1966 int32
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1983 int32
	_ = v1983
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1998 int32
	_ = v1998
	var v2001 int32
	_ = v2001
	var v2009 int32
	_ = v2009
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2030 int32
	_ = v2030
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2062 int32
	_ = v2062
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2086 int32
	_ = v2086
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2108 int32
	_ = v2108
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2119 int32
	_ = v2119
	var v2124 int32
	_ = v2124
	var v2128 int32
	_ = v2128
	var v2130 int32
	_ = v2130
	var v2134 int32
	_ = v2134
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2145 int32
	_ = v2145
	var v2150 int32
	_ = v2150
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2160 int32
	_ = v2160
	var v2162 int32
	_ = v2162
	var v2164 int32
	_ = v2164
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2185 int32
	_ = v2185
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2231 int32
	_ = v2231
	var v2235 int32
	_ = v2235
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2250 int32
	_ = v2250
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2270 int32
	_ = v2270
	var v2273 int32
	_ = v2273
	var v2287 int32
	_ = v2287
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2294 int32
	_ = v2294
	var v2300 int32
	_ = v2300
	var v2305 int32
	_ = v2305
	var v2317 int32
	_ = v2317
	var v2330 int32
	_ = v2330
	var v2342 int32
	_ = v2342
	v12 = m.G0
	v14 = v12 - int32(1184)
	m.G0 = v14
	F_initStringInfo(m, v14+int32(1168))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
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
	v22 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v22
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v22
	goto L5
L4:
	;
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v26 <= int32(2752) {
		goto L56
	} else {
		goto L57
	}
L6:
	;
	m.G0 = v14 + int32(1184)
	return v2342
L7:
	;
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1168))
	v2342 = v2330
	goto L6
L8:
	;
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1172))
	if v2317 != 0 {
		goto L7
	} else {
		goto L749
	}
L9:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L743
	}
L10:
	;
	v2153 = F_table_open(m, int32(826), int32(1))
	mBase = m.M
	v2154 = m.ExcPending
	if v2154 != 0 {
		goto L1
	} else {
		goto L704
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L1
	} else {
		goto L701
	}
L12:
	;
	F_sequence_close(m, v2130, int32(1))
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L1
	} else {
		goto L700
	}
L13:
	;
	F_getRelationIdentity(m, v14+int32(1168), v82, l1, l3)
	mBase = m.M
	v2128 = m.ExcPending
	if v2128 != 0 {
		goto L1
	} else {
		goto L699
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L1
	} else {
		goto L696
	}
L15:
	;
	if v26 == int32(826) {
		goto L10
	} else {
		goto L695
	}
L16:
	;
	v2043 = F_table_open(m, int32(3576), int32(1))
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L1
	} else {
		goto L676
	}
L17:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2021 = F_get_subscription_name(m, v2020, l3)
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L1
	} else {
		goto L670
	}
L18:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1969 = F_SearchSysCache1(m, int32(52), v1968)
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L1
	} else {
		goto L654
	}
L19:
	;
	v1932 = F_getPublicationSchemaInfo(m, l0, l3, v14+int32(1056), v14+int32(1040))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L1
	} else {
		goto L640
	}
L20:
	;
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1908 = F_get_publication_name(m, v1907, l3)
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L1
	} else {
		goto L634
	}
L21:
	;
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v1878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1877)+22)))
	v1879 = v1877 + v1878
	v1881 = v1879 + int32(4)
	v1882 = F_quote_identifier(m, v1881)
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L1
	} else {
		goto L625
	}
L22:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1837 = F_SearchSysCache1(m, int32(44), v1836)
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L1
	} else {
		goto L609
	}
L23:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1792 = F_SearchSysCache1(m, int32(26), v1791)
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L1
	} else {
		goto L593
	}
L24:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1754 = F_get_extension_name(m, v1753)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L1
	} else {
		goto L581
	}
L25:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1685 = F_SearchSysCache1(m, int32(83), v1684)
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L1
	} else {
		goto L558
	}
L26:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1659 = F_GetForeignServerExtended(m, v1658, l3)
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L1
	} else {
		goto L551
	}
L27:
	;
	if v26 != int32(2328) {
		goto L14
	} else {
		goto L543
	}
L28:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1597 = F_get_tablespace_name(m, v1596)
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L1
	} else {
		goto L529
	}
L29:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1562 = F_get_database_name(m, v1561)
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L1
	} else {
		goto L515
	}
L30:
	;
	v1496 = F_table_open(m, int32(1261), int32(1))
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L1
	} else {
		goto L498
	}
L31:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1476 = F_GetUserNameFromId(m, v1475, l3)
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L1
	} else {
		goto L490
	}
L32:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1424 = F_SearchSysCache1(m, int32(74), v1423)
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L1
	} else {
		goto L473
	}
L33:
	;
	if v26 != int32(3764) {
		goto L14
	} else {
		goto L455
	}
L34:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1316 = F_SearchSysCache1(m, int32(76), v1315)
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L1
	} else {
		goto L438
	}
L35:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1263 = F_SearchSysCache1(m, int32(78), v1262)
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L1
	} else {
		goto L421
	}
L36:
	;
	if v26 != int32(3381) {
		goto L14
	} else {
		goto L403
	}
L37:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1170 = F_get_namespace_name_or_temp(m, v1169)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L1
	} else {
		goto L391
	}
L38:
	;
	v1114 = F_table_open(m, int32(2620), int32(1))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L1
	} else {
		goto L373
	}
L39:
	;
	v1057 = F_table_open(m, int32(2618), int32(1))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L355
	}
L40:
	;
	v950 = F_table_open(m, int32(2603), int32(1))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L329
	}
L41:
	;
	v843 = F_table_open(m, int32(2602), int32(1))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L303
	}
L42:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v805 = F_get_am_name(m, v804)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L291
	}
L43:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_getOpFamilyIdentity(m, v14+int32(1168), v801, l1, l3)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L290
	}
L44:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v721 = F_SearchSysCache1(m, int32(14), v720)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L268
	}
L45:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v639 = F_format_operator_extended(m, v637, int32(3))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L241
	}
L46:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v607 = F_LargeObjectExists(m, v606)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L235
	}
L47:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v562 = F_SearchSysCache1(m, int32(36), v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L219
	}
L48:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_GetAttrDefaultColumnAddress(m, v14+int32(1056), v525)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L209
	}
L49:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v472 = F_SearchSysCache1(m, int32(20), v471)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L192
	}
L50:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v381 = F_SearchSysCache1(m, int32(19), v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L166
	}
L51:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v328 = F_SearchSysCache1(m, int32(16), v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L149
	}
L52:
	;
	v256 = F_table_open(m, int32(2605), int32(1))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L129
	}
L53:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v236 = F_format_type_extended(m, v233, int32(-1), int32(12))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L124
	}
L54:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v119 = F_format_procedure_extended(m, v117, int32(3))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L99
	}
L55:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v83 == int32(0) {
		goto L13
	} else {
		goto L83
	}
L56:
	;
	if v26 <= int32(2327) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	if v26 <= int32(3575) {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	switch v26 - int32(1213) {
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
	switch v26 - int32(2601) {
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
	switch v26 - int32(1417) {
	case 0:
		goto L26
	case 1:
		goto L25
	default:
		goto L15
	}
L63:
	;
	if v26 <= int32(3380) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	if v26 <= int32(6099) {
		goto L79
	} else {
		goto L80
	}
L66:
	;
	if v26 == int32(2753) {
		goto L43
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	switch v26 - int32(3456) {
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
	if v26 == int32(3079) {
		goto L24
	} else {
		goto L70
	}
L70:
	;
	if v26 != int32(3256) {
		goto L14
	} else {
		goto L71
	}
L71:
	;
	v49 = F_table_open(m, int32(3256), int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v54 = F_get_catalog_object_by_oid_extended(m, v49, int32(1), v52, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v54 != 0 {
		goto L21
	} else {
		goto L74
	}
L74:
	;
	if l3 != 0 {
		v2130 = v49
		goto L12
	} else {
		goto L75
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+800)) = v60
	F_errmsg_internal(m, int32(38408), v14+int32(800))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(476959), int32(5880), int32(112201))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
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
	switch v26 - int32(3576) {
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
	switch v26 - int32(6100) {
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
	switch v26 - int32(6237) {
	case 0:
		goto L19
	default:
		goto L14
	case 6:
		goto L22
	}
L83:
	;
	v88 = F_get_attname(m, v82, base.I32_extend16_s(v83), l3)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if v88 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v90 = int32(0)
	goto L87
L86:
	;
	v90 = l3
	goto L87
L87:
	;
	if v90 != 0 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_getRelationIdentity(m, v14+int32(1168), v93, l1, l3)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
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
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v96 == int32(0) {
		goto L9
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if v88 == int32(0) {
		goto L9
	} else {
		goto L94
	}
L93:
	;
	goto L92
L94:
	;
	v101 = F_quote_identifier(m, v88)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v101
	F_appendStringInfo(m, v14+int32(1168), int32(169868), v14+int32(32))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
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
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v114 = F_lappend(m, v113, v88)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v114
	goto L9
L99:
	;
	if v119 == int32(0) {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	F_appendStringInfoString(m, v14+int32(1168), v119)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
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
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v130 = m.G0
	v132 = v130 - int32(32)
	m.G0 = v132
	v135 = F_SearchSysCache1(m, int32(47), v129)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L104
	}
L103:
	;
	m.G0 = v132 + int32(32)
	goto L9
L104:
	;
	if v135 == int32(0) {
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
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v135)+16))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+22)))
	v154 = v152 + v153
	v155 = int32(*(*int16)(unsafe.Add(mBase, uint32(v154)+104)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)+68))
	v157 = F_get_namespace_name_or_temp(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L112
	}
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v129
	F_errmsg_internal(m, int32(54944), v132)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(482938), int32(411), int32(112178))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v132)+28)) = v157
	v162 = F_pstrdup(m, v154+int32(4))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+24)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v132)+16)) = v162
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v132)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+20)) = v166
	v172 = F_list_make2_impl(m, v132+int32(20), v132+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v172
	v175 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v175
	if v175 < v155 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v188 = int32(0)
	v191 = v175
	goto L118
L116:
	;
	goto L117
L117:
	;
	F_ReleaseCatCache(m, v135)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L123
	}
L118:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v154+int32(136)+v188<<(uint(int32(2))%32))))
	v198 = F_format_type_be_qualified(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L120
	}
L119:
	;
	goto L117
L120:
	;
	v200 = F_lappend(m, v191, v198)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v200
	v204 = v188 + int32(1)
	if v204 != v155 {
		v188 = v204
		v191 = v200
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
	if v236 == int32(0) {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	F_appendStringInfoString(m, v14+int32(1168), v236)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1164)) = v236
	v251 = F_list_make1_impl(m, int32(1), v14+int32(44))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v251
	goto L9
L129:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v261 = F_get_catalog_object_by_oid_extended(m, v256, int32(1), v259, int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	if v261 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if l3 != 0 {
		v2130 = v256
		goto L12
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+22)))
	v283 = v281 + v282
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	v285 = F_format_type_be_qualified(m, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L138
	}
L134:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v269
	F_errmsg_internal(m, int32(38883), v14+int32(48))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(476959), int32(4942), int32(112201))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
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
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283)+8))
	v288 = F_format_type_be_qualified(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v285
	F_appendStringInfo(m, v14+int32(1168), int32(641128), v14-int32(-64))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
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
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	v300 = F_format_type_be_qualified(m, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	F_sequence_close(m, v256, int32(1))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L148
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1160)) = v300
	v307 = F_list_make1_impl(m, int32(1), v14+int32(60))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v307
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v283)+8))
	v311 = F_format_type_be_qualified(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1156)) = v311
	v318 = F_list_make1_impl(m, int32(1), v14+int32(56))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v318
	goto L143
L148:
	;
	goto L9
L149:
	;
	if v328 == int32(0) {
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
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v328)+16))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+22)))
	v352 = v350 + v351
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+68))
	v354 = F_get_namespace_name_or_temp(m, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L157
	}
L153:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v336
	F_errmsg_internal(m, int32(44384), v14+int32(80))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(476959), int32(4976), int32(112201))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
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
	v357 = v352 + int32(4)
	v358 = F_quote_qualified_identifier(m, v354, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_appendStringInfoString(m, v14+int32(1168), v358)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1152)) = v354
	v363 = F_pstrdup(m, v357)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	F_ReleaseCatCache(m, v328)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L165
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1148)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v363
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1152))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+92)) = v367
	v373 = F_list_make2_impl(m, v14+int32(92), v14+int32(88))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v373
	goto L162
L165:
	;
	goto L9
L166:
	;
	if v381 == int32(0) {
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
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v381)+16))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+22)))
	v403 = v401 + v402
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+80))
	if v404 != 0 {
		goto L175
	} else {
		goto L176
	}
L170:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v389
	F_errmsg_internal(m, int32(39476), v14+int32(96))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(476959), int32(5002), int32(112201))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
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
	F_ReleaseCatCache(m, v381)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L191
	}
L175:
	;
	v406 = v403 + int32(4)
	v407 = F_quote_identifier(m, v406)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1056)) = int32(1247)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v403)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1064)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1060)) = v435
	v440 = v403 + int32(4)
	v441 = F_quote_identifier(m, v440)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L185
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v407
	F_appendStringInfo(m, v14+int32(1168), int32(704715), v14+int32(128))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v403)+80))
	F_getRelationIdentity(m, v14+int32(1168), v419, l1, int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
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
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v426 = F_pstrdup(m, v406)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v428 = F_lappend(m, v425, v426)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v428
	F_ReleaseCatCache(m, v381)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	goto L9
L185:
	;
	v446 = F_getObjectIdentityParts(m, v14+int32(1056), l1, l2, int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+116)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v441
	F_appendStringInfo(m, v14+int32(1168), int32(177027), v14+int32(112))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
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
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v460 = F_pstrdup(m, v440)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v462 = F_lappend(m, v459, v460)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v462
	goto L174
L191:
	;
	goto L9
L192:
	;
	if v472 == int32(0) {
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
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v472)+16))
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+22)))
	v496 = v494 + v495
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+68))
	v498 = F_get_namespace_name_or_temp(m, v497)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L200
	}
L196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v480
	F_errmsg_internal(m, int32(45229), v14+int32(144))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(476959), int32(5050), int32(112201))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
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
	v501 = v496 + int32(4)
	v502 = F_quote_qualified_identifier(m, v498, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_appendStringInfoString(m, v14+int32(1168), v502)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1144)) = v498
	v507 = F_pstrdup(m, v501)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	F_ReleaseCatCache(m, v472)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L208
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1140)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v14)+152)) = v507
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1144))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+156)) = v511
	v517 = F_list_make2_impl(m, v14+int32(156), v14+int32(152))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v517
	goto L205
L208:
	;
	goto L9
L209:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1060))
	if v528 == int32(0) {
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
	v550 = F_getObjectIdentityParts(m, v14+int32(1056), l1, l2, int32(0))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L217
	}
L213:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v535
	F_errmsg_internal(m, int32(47474), v14+int32(160))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(476959), int32(5075), int32(112201))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = v550
	F_appendStringInfo(m, v14+int32(1168), int32(174301), v14+int32(176))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	goto L9
L219:
	;
	if v562 == int32(0) {
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
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v562)+16))
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584)+22)))
	v588 = v584 + v585 + int32(4)
	v589 = F_quote_identifier(m, v588)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L227
	}
L223:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = v570
	F_errmsg_internal(m, int32(50409), v14+int32(192))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(476959), int32(5097), int32(112201))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
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
	F_appendStringInfoString(m, v14+int32(1168), v589)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
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
	v593 = F_pstrdup(m, v588)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	F_ReleaseCatCache(m, v562)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L234
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+200)) = v593
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1136)) = v593
	v600 = F_list_make1_impl(m, int32(1), v14+int32(200))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v600
	goto L231
L234:
	;
	goto L9
L235:
	;
	if v607 == int32(0) {
		goto L9
	} else {
		goto L236
	}
L236:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v611
	F_appendStringInfo(m, v14+int32(1168), int32(57825), v14+int32(224))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
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
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v622
	v627 = F_psprintf(m, int32(57825), v14+int32(208))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+204)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1132)) = v627
	v634 = F_list_make1_impl(m, int32(1), v14+int32(204))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v634
	goto L9
L241:
	;
	if v639 == int32(0) {
		goto L9
	} else {
		goto L242
	}
L242:
	;
	F_appendStringInfoString(m, v14+int32(1168), v639)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
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
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v650 = m.G0
	v652 = v650 - int32(32)
	m.G0 = v652
	v655 = F_SearchSysCache1(m, int32(40), v649)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L246
	}
L245:
	;
	m.G0 = v652 + int32(32)
	goto L9
L246:
	;
	if v655 == int32(0) {
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
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v655)+16))
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+22)))
	v674 = v672 + v673
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v674)+68))
	v676 = F_get_namespace_name_or_temp(m, v675)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L254
	}
L250:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v652))) = v649
	F_errmsg_internal(m, int32(54443), v652)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(482938), int32(817), int32(112156))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v652)+28)) = v676
	v681 = F_pstrdup(m, v674+int32(4))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v652)+24)) = v681
	*(*int32)(unsafe.Add(mBase, uint32(v652)+16)) = v681
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v652)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v652)+20)) = v685
	v691 = F_list_make2_impl(m, v652+int32(20), v652+int32(16))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v691
	v694 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v694
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v674)+80))
	if v697 != 0 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v699 = F_format_type_be_qualified(m, v697)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L260
	}
L258:
	;
	v704 = v694
	goto L259
L259:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v674)+84))
	if v705 != 0 {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	v701 = F_lappend(m, int32(0), v699)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v701
	v704 = v701
	goto L259
L262:
	;
	v706 = F_format_type_be_qualified(m, v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	F_ReleaseCatCache(m, v655)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L267
	}
L265:
	;
	v708 = F_lappend(m, v704, v706)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v708
	goto L264
L267:
	;
	goto L245
L268:
	;
	if v721 == int32(0) {
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
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v721)+16))
	v742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v741)+22)))
	v743 = v741 + v742
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v743)+72))
	v745 = F_get_namespace_name_or_temp(m, v744)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L276
	}
L272:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v729
	F_errmsg_internal(m, int32(40727), v14+int32(240))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(476959), int32(5147), int32(112201))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
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
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v743)+4))
	v749 = F_SearchSysCache1(m, int32(2), v748)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	if v749 == int32(0) {
		goto L11
	} else {
		goto L278
	}
L278:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v749)+16))
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753)+22)))
	v756 = v743 + int32(8)
	v757 = F_quote_qualified_identifier(m, v745, v756)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	v761 = v753 + v754 + int32(4)
	v762 = F_quote_identifier(m, v761)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+276)) = v762
	*(*int32)(unsafe.Add(mBase, uint32(v14)+272)) = v757
	F_appendStringInfo(m, v14+int32(1168), int32(190617), v14+int32(272))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
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
	v773 = F_pstrdup(m, v761)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	F_ReleaseCatCache(m, v749)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L288
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1124)) = v745
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1128)) = v773
	v777 = F_pstrdup(m, v756)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1120)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v14)+260)) = v777
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1128))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+268)) = v781
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1124))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+264)) = v783
	v791 = F_list_make3_impl(m, v14+int32(268), v14+int32(264), v14+int32(260))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v791
	goto L284
L288:
	;
	F_ReleaseCatCache(m, v721)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
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
	if v805 == int32(0) {
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
	v827 = F_quote_identifier(m, v805)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L1
	} else {
		goto L299
	}
L295:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+288)) = v813
	F_errmsg_internal(m, int32(51623), v14+int32(288))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(476959), int32(5188), int32(112201))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
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
	F_appendStringInfoString(m, v14+int32(1168), v827)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+300)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1116)) = v805
	v838 = F_list_make1_impl(m, int32(1), v14+int32(300))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v838
	goto L9
L303:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v14+int32(1056), int32(1), int32(3), int32(184), v850)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	v854 = int32(1)
	v859 = F_systable_beginscan(m, v843, int32(2756), v854, int32(0), v854, v14+int32(1056))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L1
	} else {
		goto L306
	}
L305:
	;
	F_systable_endscan(m, v859)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L1
	} else {
		goto L327
	}
L306:
	;
	v861 = F_systable_getnext(m, v859)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	if v861 == int32(0) {
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
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v861)+16))
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v881)+22)))
	F_initStringInfo(m, v14+int32(1040))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L315
	}
L311:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+304)) = v869
	F_errmsg_internal(m, int32(37889), v14+int32(304))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	F_errfinish(m, int32(476959), int32(5225), int32(112201))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
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
	v889 = v881 + v882
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v889)+4))
	F_getOpFamilyIdentity(m, v14+int32(1040), v890, l1, int32(0))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v889)+8))
	v895 = F_format_type_be_qualified(m, v894)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v889)+12))
	v898 = F_format_type_be_qualified(m, v897)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
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
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v901 = int32(*(*int16)(unsafe.Add(mBase, uint32(v889)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+352)) = v901
	v906 = F_psprintf(m, int32(471827), v14+int32(352))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L1
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	v923 = int32(*(*int16)(unsafe.Add(mBase, uint32(v889)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+320)) = v923
	*(*int32)(unsafe.Add(mBase, uint32(v14)+324)) = v895
	*(*int32)(unsafe.Add(mBase, uint32(v14)+328)) = v898
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1040))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+332)) = v927
	F_appendStringInfo(m, v14+int32(1168), int32(179281), v14+int32(320))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L1
	} else {
		goto L325
	}
L322:
	;
	v908 = F_lappend(m, v900, v906)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v908
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1032)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1036)) = v895
	*(*int32)(unsafe.Add(mBase, uint32(v14)+348)) = v895
	*(*int32)(unsafe.Add(mBase, uint32(v14)+344)) = v898
	v919 = F_list_make2_impl(m, v14+int32(348), v14+int32(344))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v919
	goto L321
L325:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1040))
	F_pfree(m, v936)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	goto L305
L327:
	;
	F_sequence_close(m, v843, int32(1))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	goto L9
L329:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v14+int32(1056), int32(1), int32(3), int32(184), v957)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	v961 = int32(1)
	v966 = F_systable_beginscan(m, v950, int32(2757), v961, int32(0), v961, v14+int32(1056))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L332
	}
L331:
	;
	F_systable_endscan(m, v966)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L1
	} else {
		goto L353
	}
L332:
	;
	v968 = F_systable_getnext(m, v966)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	if v968 == int32(0) {
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
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v968)+16))
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v988)+22)))
	F_initStringInfo(m, v14+int32(1040))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L1
	} else {
		goto L341
	}
L337:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+368)) = v976
	F_errmsg_internal(m, int32(37928), v14+int32(368))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	F_errfinish(m, int32(476959), int32(5287), int32(112201))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
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
	v996 = v988 + v989
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v996)+4))
	F_getOpFamilyIdentity(m, v14+int32(1040), v997, l1, int32(0))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v996)+8))
	v1002 = F_format_type_be_qualified(m, v1001)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v996)+12))
	v1005 = F_format_type_be_qualified(m, v1004)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
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
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1008 = int32(*(*int16)(unsafe.Add(mBase, uint32(v996)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+416)) = v1008
	v1013 = F_psprintf(m, int32(471827), v14+int32(416))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	v1030 = int32(*(*int16)(unsafe.Add(mBase, uint32(v996)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+384)) = v1030
	*(*int32)(unsafe.Add(mBase, uint32(v14)+388)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v14)+392)) = v1005
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1040))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+396)) = v1034
	F_appendStringInfo(m, v14+int32(1168), int32(179308), v14+int32(384))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L1
	} else {
		goto L351
	}
L348:
	;
	v1015 = F_lappend(m, v1007, v1013)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L1
	} else {
		goto L349
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1015
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1024)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1028)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v14)+412)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v14)+408)) = v1005
	v1026 = F_list_make2_impl(m, v14+int32(412), v14+int32(408))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L1
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1026
	goto L347
L351:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1040))
	F_pfree(m, v1043)
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	goto L331
L353:
	;
	F_sequence_close(m, v950, int32(1))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	goto L9
L355:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1062 = F_get_catalog_object_by_oid_extended(m, v1057, int32(1), v1060, int32(0))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	if v1062 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	if l3 != 0 {
		v2130 = v1057
		goto L12
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1062)+16))
	v1083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+22)))
	v1084 = v1082 + v1083
	v1086 = v1084 + int32(4)
	v1087 = F_quote_identifier(m, v1086)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L1
	} else {
		goto L364
	}
L360:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+432)) = v1070
	F_errmsg_internal(m, int32(50154), v14+int32(432))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	F_errfinish(m, int32(476959), int32(5336), int32(112201))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+448)) = v1087
	F_appendStringInfo(m, v14+int32(1168), int32(704715), v14+int32(448))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+68))
	F_getRelationIdentity(m, v14+int32(1168), v1099, l1, int32(0))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
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
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1104 = F_pstrdup(m, v1086)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L1
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	F_sequence_close(m, v1057, int32(1))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L372
	}
L370:
	;
	v1106 = F_lappend(m, v1103, v1104)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1106
	goto L369
L372:
	;
	goto L9
L373:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1119 = F_get_catalog_object_by_oid_extended(m, v1114, int32(1), v1117, int32(0))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L374
	}
L374:
	;
	if v1119 == int32(0) {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	if l3 != 0 {
		v2130 = v1114
		goto L12
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+16))
	v1140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1139)+22)))
	v1141 = v1139 + v1140
	v1143 = v1141 + int32(12)
	v1144 = F_quote_identifier(m, v1143)
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L1
	} else {
		goto L382
	}
L378:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+464)) = v1127
	F_errmsg_internal(m, int32(42329), v14+int32(464))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	F_errfinish(m, int32(476959), int32(5369), int32(112201))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+480)) = v1144
	F_appendStringInfo(m, v14+int32(1168), int32(704715), v14+int32(480))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1141)+4))
	F_getRelationIdentity(m, v14+int32(1168), v1156, l1, int32(0))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
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
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1161 = F_pstrdup(m, v1143)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L1
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	F_sequence_close(m, v1114, int32(1))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L390
	}
L388:
	;
	v1163 = F_lappend(m, v1160, v1161)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1163
	goto L387
L390:
	;
	goto L9
L391:
	;
	if v1170 == int32(0) {
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
	v1192 = F_quote_identifier(m, v1170)
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L1
	} else {
		goto L399
	}
L395:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+496)) = v1178
	F_errmsg_internal(m, int32(51227), v14+int32(496))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(476959), int32(5396), int32(112201))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
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
	F_appendStringInfoString(m, v14+int32(1168), v1192)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+508)) = v1170
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1020)) = v1170
	v1203 = F_list_make1_impl(m, int32(1), v14+int32(508))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1203
	goto L9
L403:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1210 = F_SearchSysCache1(m, int32(64), v1209)
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	if v1210 == int32(0) {
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
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+16))
	v1233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1232)+22)))
	v1234 = v1232 + v1233
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+72))
	v1236 = F_get_namespace_name_or_temp(m, v1235)
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L1
	} else {
		goto L412
	}
L408:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L1
	} else {
		goto L409
	}
L409:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+512)) = v1218
	F_errmsg_internal(m, int32(40289), v14+int32(512))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	F_errfinish(m, int32(476959), int32(5418), int32(112201))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
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
	v1239 = v1234 + int32(8)
	v1240 = F_quote_qualified_identifier(m, v1236, v1239)
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	F_appendStringInfoString(m, v14+int32(1168), v1240)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1016)) = v1236
	v1245 = F_pstrdup(m, v1239)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L1
	} else {
		goto L418
	}
L416:
	;
	goto L417
L417:
	;
	F_ReleaseCatCache(m, v1210)
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L1
	} else {
		goto L420
	}
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1012)) = v1245
	*(*int32)(unsafe.Add(mBase, uint32(v14)+520)) = v1245
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1016))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+524)) = v1249
	v1255 = F_list_make2_impl(m, v14+int32(524), v14+int32(520))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L1
	} else {
		goto L419
	}
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1255
	goto L417
L420:
	;
	goto L9
L421:
	;
	if v1263 == int32(0) {
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
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1263)+16))
	v1286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1285)+22)))
	v1287 = v1285 + v1286
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+68))
	v1289 = F_get_namespace_name_or_temp(m, v1288)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L1
	} else {
		goto L429
	}
L425:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+528)) = v1271
	F_errmsg_internal(m, int32(42131), v14+int32(528))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L1
	} else {
		goto L427
	}
L427:
	;
	F_errfinish(m, int32(476959), int32(5445), int32(112201))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
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
	v1292 = v1287 + int32(4)
	v1293 = F_quote_qualified_identifier(m, v1289, v1292)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	F_appendStringInfoString(m, v14+int32(1168), v1293)
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1008)) = v1289
	v1298 = F_pstrdup(m, v1292)
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L1
	} else {
		goto L435
	}
L433:
	;
	goto L434
L434:
	;
	F_ReleaseCatCache(m, v1263)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L1
	} else {
		goto L437
	}
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1004)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v14)+536)) = v1298
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1008))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+540)) = v1302
	v1308 = F_list_make2_impl(m, v14+int32(540), v14+int32(536))
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L1
	} else {
		goto L436
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1308
	goto L434
L437:
	;
	goto L9
L438:
	;
	if v1316 == int32(0) {
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
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1316)+16))
	v1339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1338)+22)))
	v1340 = v1338 + v1339
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1340)+68))
	v1342 = F_get_namespace_name_or_temp(m, v1341)
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L1
	} else {
		goto L446
	}
L442:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L1
	} else {
		goto L443
	}
L443:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+544)) = v1324
	F_errmsg_internal(m, int32(37969), v14+int32(544))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L1
	} else {
		goto L444
	}
L444:
	;
	F_errfinish(m, int32(476959), int32(5472), int32(112201))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
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
	v1345 = v1340 + int32(4)
	v1346 = F_quote_qualified_identifier(m, v1342, v1345)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L1
	} else {
		goto L447
	}
L447:
	;
	F_appendStringInfoString(m, v14+int32(1168), v1346)
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1000)) = v1342
	v1351 = F_pstrdup(m, v1345)
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L1
	} else {
		goto L452
	}
L450:
	;
	goto L451
L451:
	;
	F_ReleaseCatCache(m, v1316)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L1
	} else {
		goto L454
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+996)) = v1351
	*(*int32)(unsafe.Add(mBase, uint32(v14)+552)) = v1351
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1000))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+556)) = v1355
	v1361 = F_list_make2_impl(m, v14+int32(556), v14+int32(552))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1361
	goto L451
L454:
	;
	goto L9
L455:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1371 = F_SearchSysCache1(m, int32(80), v1370)
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	if v1371 == int32(0) {
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
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1371)+16))
	v1394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1393)+22)))
	v1395 = v1393 + v1394
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+68))
	v1397 = F_get_namespace_name_or_temp(m, v1396)
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L1
	} else {
		goto L464
	}
L460:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L1
	} else {
		goto L461
	}
L461:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+560)) = v1379
	F_errmsg_internal(m, int32(47617), v14+int32(560))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L1
	} else {
		goto L462
	}
L462:
	;
	F_errfinish(m, int32(476959), int32(5499), int32(112201))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
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
	v1400 = v1395 + int32(4)
	v1401 = F_quote_qualified_identifier(m, v1397, v1400)
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	F_appendStringInfoString(m, v14+int32(1168), v1401)
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+992)) = v1397
	v1406 = F_pstrdup(m, v1400)
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L1
	} else {
		goto L470
	}
L468:
	;
	goto L469
L469:
	;
	F_ReleaseCatCache(m, v1371)
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L1
	} else {
		goto L472
	}
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+988)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v14)+568)) = v1406
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v14)+992))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+572)) = v1410
	v1416 = F_list_make2_impl(m, v14+int32(572), v14+int32(568))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L1
	} else {
		goto L471
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1416
	goto L469
L472:
	;
	goto L9
L473:
	;
	if v1424 == int32(0) {
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
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+16))
	v1447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1446)+22)))
	v1448 = v1446 + v1447
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v1448)+68))
	v1450 = F_get_namespace_name_or_temp(m, v1449)
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L1
	} else {
		goto L481
	}
L477:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L478
	}
L478:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+576)) = v1432
	F_errmsg_internal(m, int32(44215), v14+int32(576))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L1
	} else {
		goto L479
	}
L479:
	;
	F_errfinish(m, int32(476959), int32(5526), int32(112201))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
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
	v1453 = v1448 + int32(4)
	v1454 = F_quote_qualified_identifier(m, v1450, v1453)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	F_appendStringInfoString(m, v14+int32(1168), v1454)
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+984)) = v1450
	v1459 = F_pstrdup(m, v1453)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L1
	} else {
		goto L487
	}
L485:
	;
	goto L486
L486:
	;
	F_ReleaseCatCache(m, v1424)
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L1
	} else {
		goto L489
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+980)) = v1459
	*(*int32)(unsafe.Add(mBase, uint32(v14)+580)) = v1459
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v14)+984))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+584)) = v1463
	v1469 = F_list_make2_impl(m, v14+int32(584), v14+int32(580))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1469
	goto L486
L489:
	;
	goto L9
L490:
	;
	if v1476 == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+588)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v14)+976)) = v1476
	v1485 = F_list_make1_impl(m, int32(1), v14+int32(588))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L1
	} else {
		goto L495
	}
L493:
	;
	goto L494
L494:
	;
	v1490 = F_quote_identifier(m, v1476)
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L1
	} else {
		goto L496
	}
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1485
	goto L494
L496:
	;
	F_appendStringInfoString(m, v14+int32(1168), v1490)
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	goto L9
L498:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v14+int32(1056), int32(1), int32(3), int32(184), v1503)
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	v1507 = int32(1)
	v1512 = F_systable_beginscan(m, v1496, int32(6303), v1507, int32(0), v1507, v14+int32(1056))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L1
	} else {
		goto L501
	}
L500:
	;
	F_systable_endscan(m, v1512)
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L1
	} else {
		goto L513
	}
L501:
	;
	v1514 = F_systable_getnext(m, v1512)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L1
	} else {
		goto L502
	}
L502:
	;
	if v1514 == int32(0) {
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
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1514)+16))
	v1535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1534)+22)))
	v1536 = v1534 + v1535
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+8))
	v1539 = F_GetUserNameFromId(m, v1537, int32(0))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L1
	} else {
		goto L510
	}
L506:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L1
	} else {
		goto L507
	}
L507:
	;
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+592)) = v1522
	F_errmsg_internal(m, int32(37839), v14+int32(592))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L1
	} else {
		goto L508
	}
L508:
	;
	F_errfinish(m, int32(476959), int32(5580), int32(112201))
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
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
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+4))
	v1543 = F_GetUserNameFromId(m, v1541, int32(0))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+612)) = v1543
	*(*int32)(unsafe.Add(mBase, uint32(v14)+608)) = v1539
	F_appendStringInfo(m, v14+int32(1168), int32(187395), v14+int32(608))
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	goto L500
L513:
	;
	F_sequence_close(m, v1496, int32(1))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	goto L9
L515:
	;
	if v1562 == int32(0) {
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
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+624)) = v1570
	F_errmsg_internal(m, int32(47838), v14+int32(624))
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	F_errfinish(m, int32(476959), int32(5607), int32(112201))
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+636)) = v1562
	*(*int32)(unsafe.Add(mBase, uint32(v14)+972)) = v1562
	v1587 = F_list_make1_impl(m, int32(1), v14+int32(636))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L1
	} else {
		goto L526
	}
L524:
	;
	goto L525
L525:
	;
	v1592 = F_quote_identifier(m, v1562)
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L1
	} else {
		goto L527
	}
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1587
	goto L525
L527:
	;
	F_appendStringInfoString(m, v14+int32(1168), v1592)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L1
	} else {
		goto L528
	}
L528:
	;
	goto L9
L529:
	;
	if v1597 == int32(0) {
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
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L1
	} else {
		goto L534
	}
L534:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+640)) = v1605
	F_errmsg_internal(m, int32(51283), v14+int32(640))
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L1
	} else {
		goto L535
	}
L535:
	;
	F_errfinish(m, int32(476959), int32(5626), int32(112201))
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+644)) = v1597
	*(*int32)(unsafe.Add(mBase, uint32(v14)+968)) = v1597
	v1622 = F_list_make1_impl(m, int32(1), v14+int32(644))
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L1
	} else {
		goto L540
	}
L538:
	;
	goto L539
L539:
	;
	v1627 = F_quote_identifier(m, v1597)
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L1
	} else {
		goto L541
	}
L540:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1622
	goto L539
L541:
	;
	F_appendStringInfoString(m, v14+int32(1168), v1627)
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	goto L9
L543:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1634 = F_GetForeignDataWrapperExtended(m, v1633, l3)
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L1
	} else {
		goto L544
	}
L544:
	;
	if v1634 == int32(0) {
		goto L9
	} else {
		goto L545
	}
L545:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+8))
	v1641 = F_quote_identifier(m, v1640)
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L1
	} else {
		goto L546
	}
L546:
	;
	F_appendStringInfoString(m, v14+int32(1168), v1641)
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
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
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+8))
	v1648 = F_pstrdup(m, v1647)
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L1
	} else {
		goto L549
	}
L549:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+648)) = v1648
	*(*int32)(unsafe.Add(mBase, uint32(v14)+964)) = v1648
	v1655 = F_list_make1_impl(m, int32(1), v14+int32(648))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1655
	goto L9
L551:
	;
	if v1659 == int32(0) {
		goto L9
	} else {
		goto L552
	}
L552:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1659)+12))
	v1666 = F_quote_identifier(m, v1665)
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	F_appendStringInfoString(m, v14+int32(1168), v1666)
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
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
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1659)+12))
	v1673 = F_pstrdup(m, v1672)
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L1
	} else {
		goto L556
	}
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+652)) = v1673
	*(*int32)(unsafe.Add(mBase, uint32(v14)+960)) = v1673
	v1680 = F_list_make1_impl(m, int32(1), v14+int32(652))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L1
	} else {
		goto L557
	}
L557:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1680
	goto L9
L558:
	;
	if v1685 == int32(0) {
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
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v1685)+16))
	v1706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1705)+22)))
	v1707 = v1705 + v1706
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1707)+4))
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1707)+8))
	v1710 = F_GetForeignServer(m, v1709)
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L1
	} else {
		goto L566
	}
L562:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+656)) = v1693
	F_errmsg_internal(m, int32(46671), v14+int32(656))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	F_errfinish(m, int32(476959), int32(5681), int32(112201))
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
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
	F_ReleaseCatCache(m, v1685)
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L1
	} else {
		goto L567
	}
L567:
	;
	if v1708 != 0 {
		goto L568
	} else {
		goto L569
	}
L568:
	;
	v1715 = F_GetUserNameFromId(m, v1708, int32(0))
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L1
	} else {
		goto L571
	}
L569:
	;
	v1718 = int32(474756)
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
	v1718 = v1715
	goto L570
L572:
	;
	v1719 = F_pstrdup(m, v1718)
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L1
	} else {
		goto L575
	}
L573:
	;
	goto L574
L574:
	;
	v1741 = F_quote_identifier(m, v1718)
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L1
	} else {
		goto L579
	}
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+684)) = v1719
	*(*int32)(unsafe.Add(mBase, uint32(v14)+956)) = v1719
	v1726 = F_list_make1_impl(m, int32(1), v14+int32(684))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L1
	} else {
		goto L576
	}
L576:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1726
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1710)+12))
	v1730 = F_pstrdup(m, v1729)
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+680)) = v1730
	*(*int32)(unsafe.Add(mBase, uint32(v14)+952)) = v1730
	v1737 = F_list_make1_impl(m, int32(1), v14+int32(680))
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1737
	goto L574
L579:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v1710)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+676)) = v1743
	*(*int32)(unsafe.Add(mBase, uint32(v14)+672)) = v1741
	F_appendStringInfo(m, v14+int32(1168), int32(174325), v14+int32(672))
	mBase = m.M
	v1752 = m.ExcPending
	if v1752 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	goto L9
L581:
	;
	if v1754 == int32(0) {
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
	v1776 = F_quote_identifier(m, v1754)
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L1
	} else {
		goto L589
	}
L585:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+752)) = v1762
	F_errmsg_internal(m, int32(45359), v14+int32(752))
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L1
	} else {
		goto L587
	}
L587:
	;
	F_errfinish(m, int32(476959), int32(5807), int32(112201))
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
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
	F_appendStringInfoString(m, v14+int32(1168), v1776)
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+764)) = v1754
	*(*int32)(unsafe.Add(mBase, uint32(v14)+940)) = v1754
	v1787 = F_list_make1_impl(m, int32(1), v14+int32(764))
	mBase = m.M
	v1788 = m.ExcPending
	if v1788 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1787
	goto L9
L593:
	;
	if v1792 == int32(0) {
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
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v1792)+16))
	v1815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1814)+22)))
	v1819 = F_pstrdup(m, v1814+v1815+int32(4))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L1
	} else {
		goto L601
	}
L597:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+768)) = v1800
	F_errmsg_internal(m, int32(42258), v14+int32(768))
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	F_errfinish(m, int32(476959), int32(5828), int32(112201))
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
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
	v1821 = F_quote_identifier(m, v1819)
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L1
	} else {
		goto L602
	}
L602:
	;
	F_appendStringInfoString(m, v14+int32(1168), v1821)
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+780)) = v1819
	*(*int32)(unsafe.Add(mBase, uint32(v14)+936)) = v1819
	v1830 = F_list_make1_impl(m, int32(1), v14+int32(780))
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L1
	} else {
		goto L607
	}
L605:
	;
	goto L606
L606:
	;
	F_ReleaseCatCache(m, v1792)
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L1
	} else {
		goto L608
	}
L607:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1830
	goto L606
L608:
	;
	goto L9
L609:
	;
	if v1837 == int32(0) {
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
	v1861 = F_SysCacheGetAttrNotNull(m, int32(44), v1837, int32(2))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L1
	} else {
		goto L617
	}
L613:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L1
	} else {
		goto L614
	}
L614:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+784)) = v1845
	F_errmsg_internal(m, int32(54070), v14+int32(784))
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	F_errfinish(m, int32(476959), int32(5852), int32(112201))
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
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
	v1863 = F_text_to_cstring(m, v1861)
	mBase = m.M
	v1864 = m.ExcPending
	if v1864 != 0 {
		goto L1
	} else {
		goto L618
	}
L618:
	;
	F_appendStringInfoString(m, v14+int32(1168), v1863)
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+796)) = v1863
	*(*int32)(unsafe.Add(mBase, uint32(v14)+932)) = v1863
	v1872 = F_list_make1_impl(m, int32(1), v14+int32(796))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L1
	} else {
		goto L623
	}
L621:
	;
	goto L622
L622:
	;
	F_ReleaseCatCache(m, v1837)
	mBase = m.M
	v1876 = m.ExcPending
	if v1876 != 0 {
		goto L1
	} else {
		goto L624
	}
L623:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1872
	goto L622
L624:
	;
	goto L9
L625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+816)) = v1882
	F_appendStringInfo(m, v14+int32(1168), int32(704715), v14+int32(816))
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1879)+68))
	F_getRelationIdentity(m, v14+int32(1168), v1894, l1, int32(0))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
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
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1899 = F_pstrdup(m, v1881)
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L1
	} else {
		goto L631
	}
L629:
	;
	goto L630
L630:
	;
	F_sequence_close(m, v49, int32(1))
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L1
	} else {
		goto L633
	}
L631:
	;
	v1901 = F_lappend(m, v1898, v1899)
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L1
	} else {
		goto L632
	}
L632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1901
	goto L630
L633:
	;
	goto L9
L634:
	;
	if v1908 == int32(0) {
		goto L9
	} else {
		goto L635
	}
L635:
	;
	v1914 = F_quote_identifier(m, v1908)
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	F_appendStringInfoString(m, v14+int32(1168), v1914)
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+820)) = v1908
	*(*int32)(unsafe.Add(mBase, uint32(v14)+928)) = v1908
	v1925 = F_list_make1_impl(m, int32(1), v14+int32(820))
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1925
	goto L9
L640:
	;
	if v1932 == int32(0) {
		goto L9
	} else {
		goto L641
	}
L641:
	;
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1040))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+832)) = v1936
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1056))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+836)) = v1938
	F_appendStringInfo(m, v14+int32(1168), int32(176599), v14+int32(832))
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+828)) = v1938
	*(*int32)(unsafe.Add(mBase, uint32(v14)+924)) = v1938
	v1952 = F_list_make1_impl(m, int32(1), v14+int32(828))
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L1
	} else {
		goto L647
	}
L645:
	;
	goto L646
L646:
	;
	F_pfree(m, v1938)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L1
	} else {
		goto L648
	}
L647:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1952
	goto L643
L648:
	;
	goto L643
L649:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+824)) = v1936
	*(*int32)(unsafe.Add(mBase, uint32(v14)+920)) = v1936
	v1962 = F_list_make1_impl(m, int32(1), v14+int32(824))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L1
	} else {
		goto L652
	}
L650:
	;
	goto L651
L651:
	;
	F_pfree(m, v1936)
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		goto L1
	} else {
		goto L653
	}
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1962
	goto L9
L653:
	;
	goto L9
L654:
	;
	if v1969 == int32(0) {
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
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(v1969)+16))
	v1990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1989)+22)))
	v1991 = v1989 + v1990
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v1991)+4))
	v1994 = F_get_publication_name(m, v1992, int32(0))
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L1
	} else {
		goto L662
	}
L658:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
	;
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+848)) = v1977
	F_errmsg_internal(m, int32(50252), v14+int32(848))
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L1
	} else {
		goto L660
	}
L660:
	;
	F_errfinish(m, int32(476959), int32(5949), int32(112201))
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
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
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v1991)+8))
	F_getRelationIdentity(m, v14+int32(1168), v1998, l1, int32(0))
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L1
	} else {
		goto L663
	}
L663:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+864)) = v1994
	F_appendStringInfo(m, v14+int32(1168), int32(176601), v14+int32(864))
	mBase = m.M
	v2009 = m.ExcPending
	if v2009 != 0 {
		goto L1
	} else {
		goto L664
	}
L664:
	;
	if l2 != 0 {
		goto L665
	} else {
		goto L666
	}
L665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+860)) = v1994
	*(*int32)(unsafe.Add(mBase, uint32(v14)+916)) = v1994
	v2015 = F_list_make1_impl(m, int32(1), v14+int32(860))
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		goto L1
	} else {
		goto L668
	}
L666:
	;
	goto L667
L667:
	;
	F_ReleaseCatCache(m, v1969)
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L1
	} else {
		goto L669
	}
L668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2015
	goto L667
L669:
	;
	goto L9
L670:
	;
	if v2021 == int32(0) {
		goto L9
	} else {
		goto L671
	}
L671:
	;
	v2027 = F_quote_identifier(m, v2021)
	mBase = m.M
	v2028 = m.ExcPending
	if v2028 != 0 {
		goto L1
	} else {
		goto L672
	}
L672:
	;
	F_appendStringInfoString(m, v14+int32(1168), v2027)
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L1
	} else {
		goto L673
	}
L673:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L674
	}
L674:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+876)) = v2021
	*(*int32)(unsafe.Add(mBase, uint32(v14)+912)) = v2021
	v2038 = F_list_make1_impl(m, int32(1), v14+int32(876))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L1
	} else {
		goto L675
	}
L675:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2038
	goto L9
L676:
	;
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2048 = F_get_catalog_object_by_oid_extended(m, v2043, int32(1), v2046, int32(0))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L1
	} else {
		goto L677
	}
L677:
	;
	if v2048 == int32(0) {
		goto L678
	} else {
		goto L679
	}
L678:
	;
	if l3 != 0 {
		v2130 = v2043
		goto L12
	} else {
		goto L681
	}
L679:
	;
	goto L680
L680:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v2048)+16))
	v2069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2068)+22)))
	v2070 = v2068 + v2069
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v2070)+4))
	v2072 = F_format_type_be_qualified(m, v2071)
	mBase = m.M
	v2073 = m.ExcPending
	if v2073 != 0 {
		goto L1
	} else {
		goto L685
	}
L681:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L1
	} else {
		goto L682
	}
L682:
	;
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+880)) = v2056
	F_errmsg_internal(m, int32(45607), v14+int32(880))
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L1
	} else {
		goto L683
	}
L683:
	;
	F_errfinish(m, int32(476959), int32(5999), int32(112201))
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L1
	} else {
		goto L684
	}
L684:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L685:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v2070)+8))
	v2076 = F_get_language_name(m, v2074, int32(0))
	mBase = m.M
	v2077 = m.ExcPending
	if v2077 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+900)) = v2076
	*(*int32)(unsafe.Add(mBase, uint32(v14)+896)) = v2072
	F_appendStringInfo(m, v14+int32(1168), int32(187939), v14+int32(896))
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		goto L1
	} else {
		goto L687
	}
L687:
	;
	if l1 != 0 {
		goto L688
	} else {
		goto L689
	}
L688:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+892)) = v2072
	*(*int32)(unsafe.Add(mBase, uint32(v14)+908)) = v2072
	v2092 = F_list_make1_impl(m, int32(1), v14+int32(892))
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L1
	} else {
		goto L691
	}
L689:
	;
	goto L690
L690:
	;
	F_sequence_close(m, v2043, int32(1))
	mBase = m.M
	v2108 = m.ExcPending
	if v2108 != 0 {
		goto L1
	} else {
		goto L694
	}
L691:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2092
	v2095 = F_pstrdup(m, v2076)
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L1
	} else {
		goto L692
	}
L692:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+888)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v14)+904)) = v2095
	v2102 = F_list_make1_impl(m, int32(1), v14+int32(888))
	mBase = m.M
	v2103 = m.ExcPending
	if v2103 != 0 {
		goto L1
	} else {
		goto L693
	}
L693:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2102
	goto L690
L694:
	;
	goto L9
L695:
	;
	goto L14
L696:
	;
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v2115
	F_errmsg_internal(m, int32(56216), v14)
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L1
	} else {
		goto L697
	}
L697:
	;
	F_errfinish(m, int32(476959), int32(6024), int32(112201))
	mBase = m.M
	v2124 = m.ExcPending
	if v2124 != 0 {
		goto L1
	} else {
		goto L698
	}
L698:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L699:
	;
	goto L9
L700:
	;
	goto L8
L701:
	;
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v743)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+256)) = v2139
	F_errmsg_internal(m, int32(51623), v14+int32(256))
	mBase = m.M
	v2145 = m.ExcPending
	if v2145 != 0 {
		goto L1
	} else {
		goto L702
	}
L702:
	;
	F_errfinish(m, int32(476959), int32(5157), int32(112201))
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		goto L1
	} else {
		goto L703
	}
L703:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L704:
	;
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v14+int32(1056), int32(1), int32(3), int32(184), v2160)
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L1
	} else {
		goto L705
	}
L705:
	;
	v2164 = int32(1)
	v2169 = F_systable_beginscan(m, v2153, int32(828), v2164, int32(0), v2164, v14+int32(1056))
	mBase = m.M
	v2170 = m.ExcPending
	if v2170 != 0 {
		goto L1
	} else {
		goto L707
	}
L706:
	;
	F_systable_endscan(m, v2169)
	mBase = m.M
	v2270 = m.ExcPending
	if v2270 != 0 {
		goto L1
	} else {
		goto L741
	}
L707:
	;
	v2171 = F_systable_getnext(m, v2169)
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L1
	} else {
		goto L708
	}
L708:
	;
	if v2171 == int32(0) {
		goto L709
	} else {
		goto L710
	}
L709:
	;
	if l3 != 0 {
		goto L706
	} else {
		goto L712
	}
L710:
	;
	goto L711
L711:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v2171)+16))
	v2192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2191)+22)))
	v2193 = v2191 + v2192
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v2193)+4))
	v2196 = F_GetUserNameFromId(m, v2194, int32(0))
	mBase = m.M
	v2197 = m.ExcPending
	if v2197 != 0 {
		goto L1
	} else {
		goto L716
	}
L712:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L1
	} else {
		goto L713
	}
L713:
	;
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+688)) = v2179
	F_errmsg_internal(m, int32(54030), v14+int32(688))
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L1
	} else {
		goto L714
	}
L714:
	;
	F_errfinish(m, int32(476959), int32(5733), int32(112201))
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
		goto L1
	} else {
		goto L715
	}
L715:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L716:
	;
	v2198 = F_quote_identifier(m, v2196)
	mBase = m.M
	v2199 = m.ExcPending
	if v2199 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+736)) = v2198
	F_appendStringInfo(m, v14+int32(1168), int32(187003), v14+int32(736))
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L1
	} else {
		goto L718
	}
L718:
	;
	v2208 = *(*int32)(unsafe.Add(mBase, uint32(v2193)+8))
	if v2208 != 0 {
		goto L719
	} else {
		goto L720
	}
L719:
	;
	v2209 = F_get_namespace_name_or_temp(m, v2208)
	mBase = m.M
	v2210 = m.ExcPending
	if v2210 != 0 {
		goto L1
	} else {
		goto L722
	}
L720:
	;
	v2221 = int32(0)
	goto L721
L721:
	;
	v2223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2193)+12)))
	switch v2223 - int32(76) {
	case 0:
		goto L727
	default:
		goto L725
	case 7:
		goto L731
	case 8:
		goto L729
	case 26:
		goto L730
	case 34:
		goto L728
	case 38:
		v2231 = int32(158274)
		goto L726
	}
L722:
	;
	v2211 = F_quote_identifier(m, v2209)
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L1
	} else {
		goto L723
	}
L723:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+720)) = v2211
	F_appendStringInfo(m, v14+int32(1168), int32(190288), v14+int32(720))
	mBase = m.M
	v2220 = m.ExcPending
	if v2220 != 0 {
		goto L1
	} else {
		goto L724
	}
L724:
	;
	v2221 = v2209
	goto L721
L725:
	;
	if l1 == int32(0) {
		goto L706
	} else {
		goto L733
	}
L726:
	;
	F_appendStringInfoString(m, v14+int32(1168), v2231)
	mBase = m.M
	v2235 = m.ExcPending
	if v2235 != 0 {
		goto L1
	} else {
		goto L732
	}
L727:
	;
	v2231 = int32(119621)
	goto L726
L728:
	;
	v2231 = int32(166984)
	goto L726
L729:
	;
	v2231 = int32(155139)
	goto L726
L730:
	;
	v2231 = int32(134002)
	goto L726
L731:
	;
	v2231 = int32(164195)
	goto L726
L732:
	;
	goto L725
L733:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+716)) = v2196
	*(*int32)(unsafe.Add(mBase, uint32(v14)+948)) = v2196
	v2244 = F_list_make1_impl(m, int32(1), v14+int32(716))
	mBase = m.M
	v2245 = m.ExcPending
	if v2245 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2244
	if v2221 != 0 {
		goto L735
	} else {
		goto L736
	}
L735:
	;
	v2247 = F_lappend(m, v2244, v2221)
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L1
	} else {
		goto L738
	}
L736:
	;
	goto L737
L737:
	;
	v2250 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2193)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+704)) = v2250
	v2255 = F_psprintf(m, int32(485555), v14+int32(704))
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L1
	} else {
		goto L739
	}
L738:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2247
	goto L737
L739:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+700)) = v2255
	*(*int32)(unsafe.Add(mBase, uint32(v14)+944)) = v2255
	v2262 = F_list_make1_impl(m, int32(1), v14+int32(700))
	mBase = m.M
	v2263 = m.ExcPending
	if v2263 != 0 {
		goto L1
	} else {
		goto L740
	}
L740:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2262
	goto L706
L741:
	;
	F_sequence_close(m, v2153, int32(1))
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L1
	} else {
		goto L742
	}
L742:
	;
	goto L9
L743:
	;
	if l1 == int32(0) {
		goto L7
	} else {
		goto L744
	}
L744:
	;
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v2287 != 0 {
		goto L7
	} else {
		goto L745
	}
L745:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2291 = m.ExcPending
	if v2291 != 0 {
		goto L1
	} else {
		goto L746
	}
L746:
	;
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v2292
	v2294 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1168))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v2294
	F_errmsg_internal(m, int32(664816), v14+int32(16))
	mBase = m.M
	v2300 = m.ExcPending
	if v2300 != 0 {
		goto L1
	} else {
		goto L747
	}
L747:
	;
	F_errfinish(m, int32(476959), int32(6036), int32(112201))
	mBase = m.M
	v2305 = m.ExcPending
	if v2305 != 0 {
		goto L1
	} else {
		goto L748
	}
L748:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L749:
	;
	v2342 = int32(0)
	goto L6
}
func F_get_object_class_descr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[436]))
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	m.G0 = v7 + int32(16)
	return v67
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == l0 {
		v64 = v10
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(0)
	goto L8
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[436])) = v61
	v64 = v61
	goto L2
L8:
	;
	v19 = v16 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[437])))
	if l0 != v22 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = v19 + int32(731648)
	goto L7
L10:
	;
	if v16 == int32(36) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v29 = (v16 | int32(1)) * int32(40)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[437])))
	if l0 == v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v29 + int32(731648)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v39 = (v16 | int32(2)) * int32(40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[437])))
	if l0 == v42 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v39 + int32(731648)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v49 = (v16 | int32(3)) * int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[437])))
	if l0 == v52 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = v49 + int32(731648)
	goto L7
L21:
	;
	v16 = v16 + int32(4)
	goto L8
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg_internal(m, int32(57773), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(476959), int32(2777), int32(486912))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = F_palloc(m, int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = int64(137438953472)
		v11 = F_palloc(m, int32(384))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = v11
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
