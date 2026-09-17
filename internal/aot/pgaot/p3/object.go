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
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
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
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v221 int32
	_ = v221
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
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
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
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
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
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
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
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
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
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
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
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
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
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
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
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
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1071 int32
	_ = v1071
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
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1175 int32
	_ = v1175
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1215 int32
	_ = v1215
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
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
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
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1321 int32
	_ = v1321
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
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
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1376 int32
	_ = v1376
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1429 int32
	_ = v1429
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1565 int32
	_ = v1565
	var v1570 int32
	_ = v1570
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1600 int32
	_ = v1600
	var v1605 int32
	_ = v1605
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1688 int32
	_ = v1688
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
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
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
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
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1757 int32
	_ = v1757
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1795 int32
	_ = v1795
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1840 int32
	_ = v1840
	var v1845 int32
	_ = v1845
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1875 int32
	_ = v1875
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1933 int32
	_ = v1933
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1943 int32
	_ = v1943
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1970 int32
	_ = v1970
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
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
	var v1988 int32
	_ = v1988
	var v1994 int32
	_ = v1994
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2047 int32
	_ = v2047
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
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
	var v2071 int32
	_ = v2071
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2093 int32
	_ = v2093
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2104 int32
	_ = v2104
	var v2109 int32
	_ = v2109
	var v2113 int32
	_ = v2113
	var v2115 int32
	_ = v2115
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2130 int32
	_ = v2130
	var v2135 int32
	_ = v2135
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2141 int32
	_ = v2141
	var v2145 int32
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2149 int32
	_ = v2149
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2168 int32
	_ = v2168
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
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
	var v2195 int32
	_ = v2195
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2204 int32
	_ = v2204
	var v2212 int32
	_ = v2212
	var v2216 int32
	_ = v2216
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2252 int32
	_ = v2252
	var v2255 int32
	_ = v2255
	var v2270 int32
	_ = v2270
	var v2274 int32
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2277 int32
	_ = v2277
	var v2283 int32
	_ = v2283
	var v2288 int32
	_ = v2288
	var v2301 int32
	_ = v2301
	var v2315 int32
	_ = v2315
	var v2328 int32
	_ = v2328
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
	return v2328
L7:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1168))
	v2328 = v2315
	goto L6
L8:
	;
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1172))
	if v2301 != 0 {
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
	v2138 = F_table_open(m, int32(826), int32(1))
	mBase = m.M
	v2139 = m.ExcPending
	if v2139 != 0 {
		goto L1
	} else {
		goto L704
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L1
	} else {
		goto L701
	}
L12:
	;
	F_relation_close(m, v2115, int32(1))
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L1
	} else {
		goto L700
	}
L13:
	;
	F_getRelationIdentity(m, v15+int32(1168), v83, l1, l3)
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L1
	} else {
		goto L699
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L1
	} else {
		goto L696
	}
L15:
	;
	if v27 == int32(826) {
		goto L10
	} else {
		goto L695
	}
L16:
	;
	v2028 = F_table_open(m, int32(3576), int32(1))
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L1
	} else {
		goto L676
	}
L17:
	;
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2006 = F_get_subscription_name(m, v2005, l3)
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L1
	} else {
		goto L670
	}
L18:
	;
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1956 = F_SearchSysCache1(m, int32(52), v1955)
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L1
	} else {
		goto L654
	}
L19:
	;
	v1919 = F_getPublicationSchemaInfo(m, l0, l3, v15+int32(1056), v15+int32(1040))
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L1
	} else {
		goto L640
	}
L20:
	;
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1895 = F_get_publication_name(m, v1894, l3)
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L1
	} else {
		goto L634
	}
L21:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	v1867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1866)+22)))
	v1868 = v1866 + v1867
	v1870 = v1868 + int32(4)
	v1871 = F_quote_identifier(m, v1870)
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L1
	} else {
		goto L625
	}
L22:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1826 = F_SearchSysCache1(m, int32(44), v1825)
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L1
	} else {
		goto L609
	}
L23:
	;
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1781 = F_SearchSysCache1(m, int32(26), v1780)
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L1
	} else {
		goto L593
	}
L24:
	;
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1743 = F_get_extension_name(m, v1742)
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L1
	} else {
		goto L581
	}
L25:
	;
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1674 = F_SearchSysCache1(m, int32(83), v1673)
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L1
	} else {
		goto L558
	}
L26:
	;
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1648 = F_GetForeignServerExtended(m, v1647, l3)
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
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
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1586 = F_get_tablespace_name(m, v1585)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L1
	} else {
		goto L529
	}
L29:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1551 = F_get_database_name(m, v1550)
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L1
	} else {
		goto L515
	}
L30:
	;
	v1487 = F_table_open(m, int32(1261), int32(1))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L1
	} else {
		goto L498
	}
L31:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1467 = F_GetUserNameFromId(m, v1466, l3)
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L1
	} else {
		goto L490
	}
L32:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1415 = F_SearchSysCache1(m, int32(74), v1414)
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
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
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1307 = F_SearchSysCache1(m, int32(76), v1306)
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L1
	} else {
		goto L438
	}
L35:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1254 = F_SearchSysCache1(m, int32(78), v1253)
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
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
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1161 = F_get_namespace_name_or_temp(m, v1160)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L1
	} else {
		goto L391
	}
L38:
	;
	v1107 = F_table_open(m, int32(2620), int32(1))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L1
	} else {
		goto L373
	}
L39:
	;
	v1052 = F_table_open(m, int32(2618), int32(1))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L355
	}
L40:
	;
	v949 = F_table_open(m, int32(2603), int32(1))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L329
	}
L41:
	;
	v846 = F_table_open(m, int32(2602), int32(1))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L303
	}
L42:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v808 = F_get_am_name(m, v807)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L1
	} else {
		goto L291
	}
L43:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_getOpFamilyIdentity(m, v15+int32(1168), v804, l1, l3)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L290
	}
L44:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v723 = F_SearchSysCache1(m, int32(14), v722)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L268
	}
L45:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v641 = F_format_operator_extended(m, v639, int32(3))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L241
	}
L46:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v609 = F_LargeObjectExists(m, v608)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L235
	}
L47:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v564 = F_SearchSysCache1(m, int32(36), v563)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L219
	}
L48:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_GetAttrDefaultColumnAddress(m, v15+int32(1056), v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L209
	}
L49:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v474 = F_SearchSysCache1(m, int32(20), v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L192
	}
L50:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v385 = F_SearchSysCache1(m, int32(19), v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L166
	}
L51:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v332 = F_SearchSysCache1(m, int32(16), v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L149
	}
L52:
	;
	v260 = F_table_open(m, int32(2605), int32(1))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L129
	}
L53:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v240 = F_format_type_extended(m, v237, int32(-1), int32(12))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
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
		v2115 = v50
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
	v136 = F_SearchSysCache1(m, int32(47), v130)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
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
	if v136 == int32(0) {
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
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+22)))
	v155 = v153 + v154
	v156 = int32(*(*int16)(unsafe.Add(mBase, uint32(v155)+104)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v155)+68))
	v158 = F_get_namespace_name_or_temp(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L112
	}
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v130
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_8), v133)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_9), int32(411), int32(_a_F_getObjectIdentityParts_10))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v133)+28)) = v158
	v163 = F_pstrdup(m, v155+int32(4))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+24)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v133)+16)) = v163
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v133)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+20)) = v167
	v173 = F_list_make2_impl(m, v133+int32(20), v133+int32(16))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v173
	v176 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v176
	if v176 < v156 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v189 = v176
	v191 = int32(0)
	goto L118
L116:
	;
	goto L117
L117:
	;
	F_ReleaseCatCache(m, v136)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L123
	}
L118:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v155+int32(136)+v189<<(uint(int32(2))%32))))
	v200 = F_format_type_be_qualified(m, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L120
	}
L119:
	;
	goto L117
L120:
	;
	v202 = F_lappend(m, v191, v200)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v202
	v206 = v189 + int32(1)
	if v206 != v156 {
		v189 = v206
		v191 = v202
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
	if v240 == int32(0) {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	F_appendStringInfoString(m, v15+int32(1168), v240)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1164)) = v240
	v255 = F_list_make1_impl(m, int32(1), v15+int32(44))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v255
	goto L9
L129:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v265 = F_get_catalog_object_by_oid_extended(m, v260, int32(1), v263, int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	if v265 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if l3 != 0 {
		v2115 = v260
		goto L12
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v265)+16))
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+22)))
	v287 = v285 + v286
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	v289 = F_format_type_be_qualified(m, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L138
	}
L134:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v273
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_11), v15+int32(48))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_12), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
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
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v287)+8))
	v292 = F_format_type_be_qualified(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v289
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_13), v15-int32(-64))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
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
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	v304 = F_format_type_be_qualified(m, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	F_relation_close(m, v260, int32(1))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L148
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1160)) = v304
	v311 = F_list_make1_impl(m, int32(1), v15+int32(60))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v311
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v287)+8))
	v315 = F_format_type_be_qualified(m, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1156)) = v315
	v322 = F_list_make1_impl(m, int32(1), v15+int32(56))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v322
	goto L143
L148:
	;
	goto L9
L149:
	;
	if v332 == int32(0) {
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
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v332)+16))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+22)))
	v356 = v354 + v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+68))
	v358 = F_get_namespace_name_or_temp(m, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L157
	}
L153:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v340
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_14), v15+int32(80))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_15), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
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
	v361 = v356 + int32(4)
	v362 = F_quote_qualified_identifier(m, v358, v361)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_appendStringInfoString(m, v15+int32(1168), v362)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1152)) = v358
	v367 = F_pstrdup(m, v361)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	F_ReleaseCatCache(m, v332)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L165
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1148)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = v367
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1152))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+92)) = v371
	v377 = F_list_make2_impl(m, v15+int32(92), v15+int32(88))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v377
	goto L162
L165:
	;
	goto L9
L166:
	;
	if v385 == int32(0) {
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
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v385)+16))
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+22)))
	v407 = v405 + v406
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)+80))
	if v408 != 0 {
		goto L175
	} else {
		goto L176
	}
L170:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v393
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_16), v15+int32(96))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_17), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
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
	F_ReleaseCatCache(m, v385)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L191
	}
L175:
	;
	v410 = v407 + int32(4)
	v411 = F_quote_identifier(m, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
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
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v407)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1064)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1060)) = v437
	v442 = v407 + int32(4)
	v443 = F_quote_identifier(m, v442)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L185
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v411
	v415 = v15 + int32(1168)
	F_appendStringInfo(m, v415, int32(_a_F_getObjectIdentityParts_18), v15+int32(128))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v407)+80))
	F_getRelationIdentity(m, v415, v421, l1, int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
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
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v428 = F_pstrdup(m, v410)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v430 = F_lappend(m, v427, v428)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v430
	F_ReleaseCatCache(m, v385)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	goto L9
L185:
	;
	v448 = F_getObjectIdentityParts(m, v15+int32(1056), l1, l2, int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v443
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_19), v15+int32(112))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
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
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v462 = F_pstrdup(m, v442)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v464 = F_lappend(m, v461, v462)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v464
	goto L174
L191:
	;
	goto L9
L192:
	;
	if v474 == int32(0) {
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
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v474)+16))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+22)))
	v498 = v496 + v497
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+68))
	v500 = F_get_namespace_name_or_temp(m, v499)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L200
	}
L196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v482
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_20), v15+int32(144))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_21), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
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
	v503 = v498 + int32(4)
	v504 = F_quote_qualified_identifier(m, v500, v503)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_appendStringInfoString(m, v15+int32(1168), v504)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1144)) = v500
	v509 = F_pstrdup(m, v503)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	F_ReleaseCatCache(m, v474)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L208
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1140)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v509
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1144))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+156)) = v513
	v519 = F_list_make2_impl(m, v15+int32(156), v15+int32(152))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v519
	goto L205
L208:
	;
	goto L9
L209:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1060))
	if v530 == int32(0) {
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
	v552 = F_getObjectIdentityParts(m, v15+int32(1056), l1, l2, int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L217
	}
L213:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v537
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_22), v15+int32(160))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_23), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v552
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_24), v15+int32(176))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	goto L9
L219:
	;
	if v564 == int32(0) {
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
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v564)+16))
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+22)))
	v590 = v586 + v587 + int32(4)
	v591 = F_quote_identifier(m, v590)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L227
	}
L223:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = v572
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_25), v15+int32(192))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_26), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
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
	F_appendStringInfoString(m, v15+int32(1168), v591)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
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
	v595 = F_pstrdup(m, v590)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	F_ReleaseCatCache(m, v564)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L234
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+200)) = v595
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1136)) = v595
	v602 = F_list_make1_impl(m, int32(1), v15+int32(200))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v602
	goto L231
L234:
	;
	goto L9
L235:
	;
	if v609 == int32(0) {
		goto L9
	} else {
		goto L236
	}
L236:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = v613
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_27), v15+int32(224))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
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
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v624
	v629 = F_psprintf(m, int32(_a_F_getObjectIdentityParts_27), v15+int32(208))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+204)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1132)) = v629
	v636 = F_list_make1_impl(m, int32(1), v15+int32(204))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v636
	goto L9
L241:
	;
	if v641 == int32(0) {
		goto L9
	} else {
		goto L242
	}
L242:
	;
	F_appendStringInfoString(m, v15+int32(1168), v641)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
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
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v652 = m.G0
	v654 = v652 - int32(32)
	m.G0 = v654
	v657 = F_SearchSysCache1(m, int32(40), v651)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L246
	}
L245:
	;
	m.G0 = v654 + int32(32)
	goto L9
L246:
	;
	if v657 == int32(0) {
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
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v657)+16))
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674)+22)))
	v676 = v674 + v675
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v676)+68))
	v678 = F_get_namespace_name_or_temp(m, v677)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L254
	}
L250:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v654))) = v651
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_28), v654)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_9), int32(817), int32(_a_F_getObjectIdentityParts_29))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v654)+28)) = v678
	v683 = F_pstrdup(m, v676+int32(4))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v654)+24)) = v683
	*(*int32)(unsafe.Add(mBase, uint32(v654)+16)) = v683
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v654)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v654)+20)) = v687
	v693 = F_list_make2_impl(m, v654+int32(20), v654+int32(16))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v693
	v696 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v696
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v676)+80))
	if v699 != 0 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v701 = F_format_type_be_qualified(m, v699)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L260
	}
L258:
	;
	v706 = v696
	goto L259
L259:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v676)+84))
	if v707 != 0 {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	v703 = F_lappend(m, int32(0), v701)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v703
	v706 = v703
	goto L259
L262:
	;
	v708 = F_format_type_be_qualified(m, v707)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	F_ReleaseCatCache(m, v657)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L267
	}
L265:
	;
	v710 = F_lappend(m, v706, v708)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v710
	goto L264
L267:
	;
	goto L245
L268:
	;
	if v723 == int32(0) {
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
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v723)+16))
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+22)))
	v745 = v743 + v744
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v745)+72))
	v747 = F_get_namespace_name_or_temp(m, v746)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L276
	}
L272:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v731
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_30), v15+int32(240))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_31), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
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
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v745)+4))
	v751 = F_SearchSysCache1(m, int32(2), v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	if v751 == int32(0) {
		goto L11
	} else {
		goto L278
	}
L278:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v751)+16))
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+22)))
	v758 = v745 + int32(8)
	v759 = F_quote_qualified_identifier(m, v747, v758)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	v763 = v755 + v756 + int32(4)
	v764 = F_quote_identifier(m, v763)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+276)) = v764
	*(*int32)(unsafe.Add(mBase, uint32(v15)+272)) = v759
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_32), v15+int32(272))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
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
	v775 = F_pstrdup(m, v763)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	F_ReleaseCatCache(m, v751)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L288
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1124)) = v747
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1128)) = v775
	v779 = F_pstrdup(m, v758)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1120)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v779
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1128))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v783
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1124))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v785
	v793 = F_list_make3_impl(m, v15+int32(268), v15+int32(264), v15+int32(260))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v793
	goto L284
L288:
	;
	F_ReleaseCatCache(m, v723)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
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
	if v808 == int32(0) {
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
	v830 = F_quote_identifier(m, v808)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L299
	}
L295:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+288)) = v816
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_33), v15+int32(288))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_34), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
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
	F_appendStringInfoString(m, v15+int32(1168), v830)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+300)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1116)) = v808
	v841 = F_list_make1_impl(m, int32(1), v15+int32(300))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v841
	goto L9
L303:
	;
	v849 = v15 + int32(1056)
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v849, int32(1), int32(3), int32(184), v853)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	v857 = int32(1)
	v860 = F_systable_beginscan(m, v846, int32(2756), v857, int32(0), v857, v849)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L1
	} else {
		goto L306
	}
L305:
	;
	F_systable_endscan(m, v860)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L1
	} else {
		goto L327
	}
L306:
	;
	v862 = F_systable_getnext(m, v860)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	if v862 == int32(0) {
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
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v862)+16))
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v882)+22)))
	v885 = v15 + int32(1040)
	F_initStringInfo(m, v885)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L1
	} else {
		goto L315
	}
L311:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+304)) = v870
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_35), v15+int32(304))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_36), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
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
	v888 = v882 + v883
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v888)+4))
	F_getOpFamilyIdentity(m, v885, v889, l1, int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v888)+8))
	v894 = F_format_type_be_qualified(m, v893)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v888)+12))
	v897 = F_format_type_be_qualified(m, v896)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
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
	v899 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v900 = int32(*(*int16)(unsafe.Add(mBase, uint32(v888)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+352)) = v900
	v905 = F_psprintf(m, int32(_a_F_getObjectIdentityParts_37), v15+int32(352))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L1
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	v922 = int32(*(*int16)(unsafe.Add(mBase, uint32(v888)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+320)) = v922
	*(*int32)(unsafe.Add(mBase, uint32(v15)+324)) = v894
	*(*int32)(unsafe.Add(mBase, uint32(v15)+328)) = v897
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1040))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+332)) = v926
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_38), v15+int32(320))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L325
	}
L322:
	;
	v907 = F_lappend(m, v899, v905)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1032)) = v897
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1036)) = v894
	*(*int32)(unsafe.Add(mBase, uint32(v15)+348)) = v894
	*(*int32)(unsafe.Add(mBase, uint32(v15)+344)) = v897
	v918 = F_list_make2_impl(m, v15+int32(348), v15+int32(344))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v918
	goto L321
L325:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1040))
	F_pfree(m, v935)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	goto L305
L327:
	;
	F_relation_close(m, v846, int32(1))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	goto L9
L329:
	;
	v952 = v15 + int32(1056)
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v952, int32(1), int32(3), int32(184), v956)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	v960 = int32(1)
	v963 = F_systable_beginscan(m, v949, int32(2757), v960, int32(0), v960, v952)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L332
	}
L331:
	;
	F_systable_endscan(m, v963)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L353
	}
L332:
	;
	v965 = F_systable_getnext(m, v963)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	if v965 == int32(0) {
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
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v965)+16))
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985)+22)))
	v988 = v15 + int32(1040)
	F_initStringInfo(m, v988)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L1
	} else {
		goto L341
	}
L337:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+368)) = v973
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_39), v15+int32(368))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_40), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
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
	v991 = v985 + v986
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v991)+4))
	F_getOpFamilyIdentity(m, v988, v992, l1, int32(0))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v991)+8))
	v997 = F_format_type_be_qualified(m, v996)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v991)+12))
	v1000 = F_format_type_be_qualified(m, v999)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
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
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1003 = int32(*(*int16)(unsafe.Add(mBase, uint32(v991)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+416)) = v1003
	v1008 = F_psprintf(m, int32(_a_F_getObjectIdentityParts_37), v15+int32(416))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L1
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	v1025 = int32(*(*int16)(unsafe.Add(mBase, uint32(v991)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+384)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v15)+388)) = v997
	*(*int32)(unsafe.Add(mBase, uint32(v15)+392)) = v1000
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1040))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+396)) = v1029
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_41), v15+int32(384))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L351
	}
L348:
	;
	v1010 = F_lappend(m, v1002, v1008)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L349
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1010
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1024)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1028)) = v997
	*(*int32)(unsafe.Add(mBase, uint32(v15)+412)) = v997
	*(*int32)(unsafe.Add(mBase, uint32(v15)+408)) = v1000
	v1021 = F_list_make2_impl(m, v15+int32(412), v15+int32(408))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1021
	goto L347
L351:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1040))
	F_pfree(m, v1038)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	goto L331
L353:
	;
	F_relation_close(m, v949, int32(1))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	goto L9
L355:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1057 = F_get_catalog_object_by_oid_extended(m, v1052, int32(1), v1055, int32(0))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	if v1057 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	if l3 != 0 {
		v2115 = v1052
		goto L12
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+16))
	v1078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1077)+22)))
	v1079 = v1077 + v1078
	v1081 = v1079 + int32(4)
	v1082 = F_quote_identifier(m, v1081)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L364
	}
L360:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+432)) = v1065
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_42), v15+int32(432))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_43), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+448)) = v1082
	v1086 = v15 + int32(1168)
	F_appendStringInfo(m, v1086, int32(_a_F_getObjectIdentityParts_18), v15+int32(448))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+68))
	F_getRelationIdentity(m, v1086, v1092, l1, int32(0))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
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
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1097 = F_pstrdup(m, v1081)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L1
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	F_relation_close(m, v1052, int32(1))
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L1
	} else {
		goto L372
	}
L370:
	;
	v1099 = F_lappend(m, v1096, v1097)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1099
	goto L369
L372:
	;
	goto L9
L373:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1112 = F_get_catalog_object_by_oid_extended(m, v1107, int32(1), v1110, int32(0))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L1
	} else {
		goto L374
	}
L374:
	;
	if v1112 == int32(0) {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	if l3 != 0 {
		v2115 = v1107
		goto L12
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+16))
	v1133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1132)+22)))
	v1134 = v1132 + v1133
	v1136 = v1134 + int32(12)
	v1137 = F_quote_identifier(m, v1136)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L1
	} else {
		goto L382
	}
L378:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+464)) = v1120
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_44), v15+int32(464))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_45), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+480)) = v1137
	v1141 = v15 + int32(1168)
	F_appendStringInfo(m, v1141, int32(_a_F_getObjectIdentityParts_18), v15+int32(480))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+4))
	F_getRelationIdentity(m, v1141, v1147, l1, int32(0))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
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
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1152 = F_pstrdup(m, v1136)
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L1
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	F_relation_close(m, v1107, int32(1))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L1
	} else {
		goto L390
	}
L388:
	;
	v1154 = F_lappend(m, v1151, v1152)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1154
	goto L387
L390:
	;
	goto L9
L391:
	;
	if v1161 == int32(0) {
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
	v1183 = F_quote_identifier(m, v1161)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L1
	} else {
		goto L399
	}
L395:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+496)) = v1169
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_46), v15+int32(496))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_47), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
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
	F_appendStringInfoString(m, v15+int32(1168), v1183)
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+508)) = v1161
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1020)) = v1161
	v1194 = F_list_make1_impl(m, int32(1), v15+int32(508))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1194
	goto L9
L403:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1201 = F_SearchSysCache1(m, int32(64), v1200)
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	if v1201 == int32(0) {
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
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1201)+16))
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1223)+22)))
	v1225 = v1223 + v1224
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+72))
	v1227 = F_get_namespace_name_or_temp(m, v1226)
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L1
	} else {
		goto L412
	}
L408:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L1
	} else {
		goto L409
	}
L409:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+512)) = v1209
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_48), v15+int32(512))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_49), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
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
	v1230 = v1225 + int32(8)
	v1231 = F_quote_qualified_identifier(m, v1227, v1230)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1231)
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1016)) = v1227
	v1236 = F_pstrdup(m, v1230)
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L1
	} else {
		goto L418
	}
L416:
	;
	goto L417
L417:
	;
	F_ReleaseCatCache(m, v1201)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L1
	} else {
		goto L420
	}
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1012)) = v1236
	*(*int32)(unsafe.Add(mBase, uint32(v15)+520)) = v1236
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1016))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+524)) = v1240
	v1246 = F_list_make2_impl(m, v15+int32(524), v15+int32(520))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L1
	} else {
		goto L419
	}
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1246
	goto L417
L420:
	;
	goto L9
L421:
	;
	if v1254 == int32(0) {
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
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+16))
	v1277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1276)+22)))
	v1278 = v1276 + v1277
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+68))
	v1280 = F_get_namespace_name_or_temp(m, v1279)
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L1
	} else {
		goto L429
	}
L425:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+528)) = v1262
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_50), v15+int32(528))
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L1
	} else {
		goto L427
	}
L427:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_51), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
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
	v1283 = v1278 + int32(4)
	v1284 = F_quote_qualified_identifier(m, v1280, v1283)
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1284)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1008)) = v1280
	v1289 = F_pstrdup(m, v1283)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L1
	} else {
		goto L435
	}
L433:
	;
	goto L434
L434:
	;
	F_ReleaseCatCache(m, v1254)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L1
	} else {
		goto L437
	}
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1004)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v15)+536)) = v1289
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1008))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+540)) = v1293
	v1299 = F_list_make2_impl(m, v15+int32(540), v15+int32(536))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L1
	} else {
		goto L436
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1299
	goto L434
L437:
	;
	goto L9
L438:
	;
	if v1307 == int32(0) {
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
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1307)+16))
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+22)))
	v1331 = v1329 + v1330
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+68))
	v1333 = F_get_namespace_name_or_temp(m, v1332)
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L1
	} else {
		goto L446
	}
L442:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L1
	} else {
		goto L443
	}
L443:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+544)) = v1315
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_52), v15+int32(544))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L1
	} else {
		goto L444
	}
L444:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_53), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
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
	v1336 = v1331 + int32(4)
	v1337 = F_quote_qualified_identifier(m, v1333, v1336)
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L1
	} else {
		goto L447
	}
L447:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1337)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1000)) = v1333
	v1342 = F_pstrdup(m, v1336)
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L1
	} else {
		goto L452
	}
L450:
	;
	goto L451
L451:
	;
	F_ReleaseCatCache(m, v1307)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L1
	} else {
		goto L454
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+996)) = v1342
	*(*int32)(unsafe.Add(mBase, uint32(v15)+552)) = v1342
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1000))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+556)) = v1346
	v1352 = F_list_make2_impl(m, v15+int32(556), v15+int32(552))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1352
	goto L451
L454:
	;
	goto L9
L455:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1362 = F_SearchSysCache1(m, int32(80), v1361)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	if v1362 == int32(0) {
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
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1362)+16))
	v1385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1384)+22)))
	v1386 = v1384 + v1385
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+68))
	v1388 = F_get_namespace_name_or_temp(m, v1387)
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L1
	} else {
		goto L464
	}
L460:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L1
	} else {
		goto L461
	}
L461:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+560)) = v1370
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_54), v15+int32(560))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L1
	} else {
		goto L462
	}
L462:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_55), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
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
	v1391 = v1386 + int32(4)
	v1392 = F_quote_qualified_identifier(m, v1388, v1391)
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1392)
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+992)) = v1388
	v1397 = F_pstrdup(m, v1391)
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L1
	} else {
		goto L470
	}
L468:
	;
	goto L469
L469:
	;
	F_ReleaseCatCache(m, v1362)
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L1
	} else {
		goto L472
	}
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+988)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v15)+568)) = v1397
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v15)+992))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+572)) = v1401
	v1407 = F_list_make2_impl(m, v15+int32(572), v15+int32(568))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L1
	} else {
		goto L471
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1407
	goto L469
L472:
	;
	goto L9
L473:
	;
	if v1415 == int32(0) {
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
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v1415)+16))
	v1438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1437)+22)))
	v1439 = v1437 + v1438
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+68))
	v1441 = F_get_namespace_name_or_temp(m, v1440)
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L1
	} else {
		goto L481
	}
L477:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L1
	} else {
		goto L478
	}
L478:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+576)) = v1423
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_56), v15+int32(576))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L1
	} else {
		goto L479
	}
L479:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_57), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
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
	v1444 = v1439 + int32(4)
	v1445 = F_quote_qualified_identifier(m, v1441, v1444)
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1445)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+984)) = v1441
	v1450 = F_pstrdup(m, v1444)
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L1
	} else {
		goto L487
	}
L485:
	;
	goto L486
L486:
	;
	F_ReleaseCatCache(m, v1415)
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L1
	} else {
		goto L489
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+980)) = v1450
	*(*int32)(unsafe.Add(mBase, uint32(v15)+580)) = v1450
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v15)+984))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+584)) = v1454
	v1460 = F_list_make2_impl(m, v15+int32(584), v15+int32(580))
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1460
	goto L486
L489:
	;
	goto L9
L490:
	;
	if v1467 == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+588)) = v1467
	*(*int32)(unsafe.Add(mBase, uint32(v15)+976)) = v1467
	v1476 = F_list_make1_impl(m, int32(1), v15+int32(588))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L1
	} else {
		goto L495
	}
L493:
	;
	goto L494
L494:
	;
	v1481 = F_quote_identifier(m, v1467)
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L1
	} else {
		goto L496
	}
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1476
	goto L494
L496:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1481)
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	goto L9
L498:
	;
	v1490 = v15 + int32(1056)
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v1490, int32(1), int32(3), int32(184), v1494)
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	v1498 = int32(1)
	v1501 = F_systable_beginscan(m, v1487, int32(_a_F_getObjectIdentityParts_58), v1498, int32(0), v1498, v1490)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L1
	} else {
		goto L501
	}
L500:
	;
	F_systable_endscan(m, v1501)
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L1
	} else {
		goto L513
	}
L501:
	;
	v1503 = F_systable_getnext(m, v1501)
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L1
	} else {
		goto L502
	}
L502:
	;
	if v1503 == int32(0) {
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
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1503)+16))
	v1524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1523)+22)))
	v1525 = v1523 + v1524
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1525)+8))
	v1528 = F_GetUserNameFromId(m, v1526, int32(0))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L1
	} else {
		goto L510
	}
L506:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L1
	} else {
		goto L507
	}
L507:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+592)) = v1511
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_59), v15+int32(592))
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L1
	} else {
		goto L508
	}
L508:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_60), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
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
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v1525)+4))
	v1532 = F_GetUserNameFromId(m, v1530, int32(0))
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+612)) = v1532
	*(*int32)(unsafe.Add(mBase, uint32(v15)+608)) = v1528
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_61), v15+int32(608))
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	goto L500
L513:
	;
	F_relation_close(m, v1487, int32(1))
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	goto L9
L515:
	;
	if v1551 == int32(0) {
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
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+624)) = v1559
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_62), v15+int32(624))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_63), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+636)) = v1551
	*(*int32)(unsafe.Add(mBase, uint32(v15)+972)) = v1551
	v1576 = F_list_make1_impl(m, int32(1), v15+int32(636))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L1
	} else {
		goto L526
	}
L524:
	;
	goto L525
L525:
	;
	v1581 = F_quote_identifier(m, v1551)
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L1
	} else {
		goto L527
	}
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1576
	goto L525
L527:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1581)
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L1
	} else {
		goto L528
	}
L528:
	;
	goto L9
L529:
	;
	if v1586 == int32(0) {
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
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L1
	} else {
		goto L534
	}
L534:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+640)) = v1594
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_64), v15+int32(640))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L1
	} else {
		goto L535
	}
L535:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_65), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+644)) = v1586
	*(*int32)(unsafe.Add(mBase, uint32(v15)+968)) = v1586
	v1611 = F_list_make1_impl(m, int32(1), v15+int32(644))
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L1
	} else {
		goto L540
	}
L538:
	;
	goto L539
L539:
	;
	v1616 = F_quote_identifier(m, v1586)
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L1
	} else {
		goto L541
	}
L540:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1611
	goto L539
L541:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1616)
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	goto L9
L543:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1623 = F_GetForeignDataWrapperExtended(m, v1622, l3)
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L1
	} else {
		goto L544
	}
L544:
	;
	if v1623 == int32(0) {
		goto L9
	} else {
		goto L545
	}
L545:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+8))
	v1630 = F_quote_identifier(m, v1629)
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L1
	} else {
		goto L546
	}
L546:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1630)
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
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
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+8))
	v1637 = F_pstrdup(m, v1636)
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L1
	} else {
		goto L549
	}
L549:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+648)) = v1637
	*(*int32)(unsafe.Add(mBase, uint32(v15)+964)) = v1637
	v1644 = F_list_make1_impl(m, int32(1), v15+int32(648))
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1644
	goto L9
L551:
	;
	if v1648 == int32(0) {
		goto L9
	} else {
		goto L552
	}
L552:
	;
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+12))
	v1655 = F_quote_identifier(m, v1654)
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1655)
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
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
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+12))
	v1662 = F_pstrdup(m, v1661)
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L1
	} else {
		goto L556
	}
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+652)) = v1662
	*(*int32)(unsafe.Add(mBase, uint32(v15)+960)) = v1662
	v1669 = F_list_make1_impl(m, int32(1), v15+int32(652))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L1
	} else {
		goto L557
	}
L557:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1669
	goto L9
L558:
	;
	if v1674 == int32(0) {
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
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1674)+16))
	v1695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1694)+22)))
	v1696 = v1694 + v1695
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v1696)+4))
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v1696)+8))
	v1699 = F_GetForeignServer(m, v1698)
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		goto L1
	} else {
		goto L566
	}
L562:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+656)) = v1682
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_66), v15+int32(656))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_67), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
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
	F_ReleaseCatCache(m, v1674)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L1
	} else {
		goto L567
	}
L567:
	;
	if v1697 != 0 {
		goto L568
	} else {
		goto L569
	}
L568:
	;
	v1704 = F_GetUserNameFromId(m, v1697, int32(0))
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L1
	} else {
		goto L571
	}
L569:
	;
	v1707 = int32(_a_F_getObjectIdentityParts_68)
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
	v1707 = v1704
	goto L570
L572:
	;
	v1708 = F_pstrdup(m, v1707)
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L1
	} else {
		goto L575
	}
L573:
	;
	goto L574
L574:
	;
	v1730 = F_quote_identifier(m, v1707)
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L1
	} else {
		goto L579
	}
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+684)) = v1708
	*(*int32)(unsafe.Add(mBase, uint32(v15)+956)) = v1708
	v1715 = F_list_make1_impl(m, int32(1), v15+int32(684))
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L1
	} else {
		goto L576
	}
L576:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1715
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1699)+12))
	v1719 = F_pstrdup(m, v1718)
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+680)) = v1719
	*(*int32)(unsafe.Add(mBase, uint32(v15)+952)) = v1719
	v1726 = F_list_make1_impl(m, int32(1), v15+int32(680))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1726
	goto L574
L579:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1699)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+676)) = v1732
	*(*int32)(unsafe.Add(mBase, uint32(v15)+672)) = v1730
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_69), v15+int32(672))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	goto L9
L581:
	;
	if v1743 == int32(0) {
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
	v1765 = F_quote_identifier(m, v1743)
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L1
	} else {
		goto L589
	}
L585:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+752)) = v1751
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_70), v15+int32(752))
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L1
	} else {
		goto L587
	}
L587:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_71), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
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
	F_appendStringInfoString(m, v15+int32(1168), v1765)
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+764)) = v1743
	*(*int32)(unsafe.Add(mBase, uint32(v15)+940)) = v1743
	v1776 = F_list_make1_impl(m, int32(1), v15+int32(764))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1776
	goto L9
L593:
	;
	if v1781 == int32(0) {
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
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1781)+16))
	v1804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1803)+22)))
	v1808 = F_pstrdup(m, v1803+v1804+int32(4))
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L1
	} else {
		goto L601
	}
L597:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1788 = m.ExcPending
	if v1788 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+768)) = v1789
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_72), v15+int32(768))
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_73), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
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
	v1810 = F_quote_identifier(m, v1808)
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L1
	} else {
		goto L602
	}
L602:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1810)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+780)) = v1808
	*(*int32)(unsafe.Add(mBase, uint32(v15)+936)) = v1808
	v1819 = F_list_make1_impl(m, int32(1), v15+int32(780))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L1
	} else {
		goto L607
	}
L605:
	;
	goto L606
L606:
	;
	F_ReleaseCatCache(m, v1781)
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L1
	} else {
		goto L608
	}
L607:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1819
	goto L606
L608:
	;
	goto L9
L609:
	;
	if v1826 == int32(0) {
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
	v1850 = F_SysCacheGetAttrNotNull(m, int32(44), v1826, int32(2))
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L1
	} else {
		goto L617
	}
L613:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L1
	} else {
		goto L614
	}
L614:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+784)) = v1834
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_74), v15+int32(784))
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_75), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
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
	v1852 = F_text_to_cstring(m, v1850)
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L1
	} else {
		goto L618
	}
L618:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1852)
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+796)) = v1852
	*(*int32)(unsafe.Add(mBase, uint32(v15)+932)) = v1852
	v1861 = F_list_make1_impl(m, int32(1), v15+int32(796))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L1
	} else {
		goto L623
	}
L621:
	;
	goto L622
L622:
	;
	F_ReleaseCatCache(m, v1826)
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L1
	} else {
		goto L624
	}
L623:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1861
	goto L622
L624:
	;
	goto L9
L625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+816)) = v1871
	v1875 = v15 + int32(1168)
	F_appendStringInfo(m, v1875, int32(_a_F_getObjectIdentityParts_18), v15+int32(816))
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1868)+68))
	F_getRelationIdentity(m, v1875, v1881, l1, int32(0))
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
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
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1886 = F_pstrdup(m, v1870)
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
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
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L1
	} else {
		goto L633
	}
L631:
	;
	v1888 = F_lappend(m, v1885, v1886)
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L1
	} else {
		goto L632
	}
L632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1888
	goto L630
L633:
	;
	goto L9
L634:
	;
	if v1895 == int32(0) {
		goto L9
	} else {
		goto L635
	}
L635:
	;
	v1901 = F_quote_identifier(m, v1895)
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	F_appendStringInfoString(m, v15+int32(1168), v1901)
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+820)) = v1895
	*(*int32)(unsafe.Add(mBase, uint32(v15)+928)) = v1895
	v1912 = F_list_make1_impl(m, int32(1), v15+int32(820))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1912
	goto L9
L640:
	;
	if v1919 == int32(0) {
		goto L9
	} else {
		goto L641
	}
L641:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1040))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+832)) = v1923
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1056))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+836)) = v1925
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_76), v15+int32(832))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+828)) = v1925
	*(*int32)(unsafe.Add(mBase, uint32(v15)+924)) = v1925
	v1939 = F_list_make1_impl(m, int32(1), v15+int32(828))
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L1
	} else {
		goto L647
	}
L645:
	;
	goto L646
L646:
	;
	F_pfree(m, v1925)
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L1
	} else {
		goto L648
	}
L647:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1939
	goto L643
L648:
	;
	goto L643
L649:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+824)) = v1923
	*(*int32)(unsafe.Add(mBase, uint32(v15)+920)) = v1923
	v1949 = F_list_make1_impl(m, int32(1), v15+int32(824))
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L1
	} else {
		goto L652
	}
L650:
	;
	goto L651
L651:
	;
	F_pfree(m, v1923)
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L1
	} else {
		goto L653
	}
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1949
	goto L9
L653:
	;
	goto L9
L654:
	;
	if v1956 == int32(0) {
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
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+16))
	v1977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1976)+22)))
	v1978 = v1976 + v1977
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v1978)+4))
	v1981 = F_get_publication_name(m, v1979, int32(0))
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L1
	} else {
		goto L662
	}
L658:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
	;
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+848)) = v1964
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_77), v15+int32(848))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L1
	} else {
		goto L660
	}
L660:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_78), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
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
	v1984 = v15 + int32(1168)
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v1978)+8))
	F_getRelationIdentity(m, v1984, v1985, l1, int32(0))
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L1
	} else {
		goto L663
	}
L663:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+864)) = v1981
	F_appendStringInfo(m, v1984, int32(_a_F_getObjectIdentityParts_79), v15+int32(864))
	mBase = m.M
	v1994 = m.ExcPending
	if v1994 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+860)) = v1981
	*(*int32)(unsafe.Add(mBase, uint32(v15)+916)) = v1981
	v2000 = F_list_make1_impl(m, int32(1), v15+int32(860))
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L1
	} else {
		goto L668
	}
L666:
	;
	goto L667
L667:
	;
	F_ReleaseCatCache(m, v1956)
	mBase = m.M
	v2004 = m.ExcPending
	if v2004 != 0 {
		goto L1
	} else {
		goto L669
	}
L668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2000
	goto L667
L669:
	;
	goto L9
L670:
	;
	if v2006 == int32(0) {
		goto L9
	} else {
		goto L671
	}
L671:
	;
	v2012 = F_quote_identifier(m, v2006)
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L1
	} else {
		goto L672
	}
L672:
	;
	F_appendStringInfoString(m, v15+int32(1168), v2012)
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+876)) = v2006
	*(*int32)(unsafe.Add(mBase, uint32(v15)+912)) = v2006
	v2023 = F_list_make1_impl(m, int32(1), v15+int32(876))
	mBase = m.M
	v2024 = m.ExcPending
	if v2024 != 0 {
		goto L1
	} else {
		goto L675
	}
L675:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2023
	goto L9
L676:
	;
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2033 = F_get_catalog_object_by_oid_extended(m, v2028, int32(1), v2031, int32(0))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L1
	} else {
		goto L677
	}
L677:
	;
	if v2033 == int32(0) {
		goto L678
	} else {
		goto L679
	}
L678:
	;
	if l3 != 0 {
		v2115 = v2028
		goto L12
	} else {
		goto L681
	}
L679:
	;
	goto L680
L680:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v2033)+16))
	v2054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2053)+22)))
	v2055 = v2053 + v2054
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v2055)+4))
	v2057 = F_format_type_be_qualified(m, v2056)
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L1
	} else {
		goto L685
	}
L681:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L1
	} else {
		goto L682
	}
L682:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+880)) = v2041
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_80), v15+int32(880))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L1
	} else {
		goto L683
	}
L683:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_81), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
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
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v2055)+8))
	v2061 = F_get_language_name(m, v2059, int32(0))
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+900)) = v2061
	*(*int32)(unsafe.Add(mBase, uint32(v15)+896)) = v2057
	F_appendStringInfo(m, v15+int32(1168), int32(_a_F_getObjectIdentityParts_82), v15+int32(896))
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+892)) = v2057
	*(*int32)(unsafe.Add(mBase, uint32(v15)+908)) = v2057
	v2077 = F_list_make1_impl(m, int32(1), v15+int32(892))
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L1
	} else {
		goto L691
	}
L689:
	;
	goto L690
L690:
	;
	F_relation_close(m, v2028, int32(1))
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L1
	} else {
		goto L694
	}
L691:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2077
	v2080 = F_pstrdup(m, v2061)
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L1
	} else {
		goto L692
	}
L692:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+888)) = v2080
	*(*int32)(unsafe.Add(mBase, uint32(v15)+904)) = v2080
	v2087 = F_list_make1_impl(m, int32(1), v15+int32(888))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L1
	} else {
		goto L693
	}
L693:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2087
	goto L690
L694:
	;
	goto L9
L695:
	;
	goto L14
L696:
	;
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v2100
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_83), v15)
	mBase = m.M
	v2104 = m.ExcPending
	if v2104 != 0 {
		goto L1
	} else {
		goto L697
	}
L697:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_84), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
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
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v745)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v2124
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_33), v15+int32(256))
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L1
	} else {
		goto L702
	}
L702:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_85), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
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
	v2141 = v15 + int32(1056)
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v2141, int32(1), int32(3), int32(184), v2145)
	mBase = m.M
	v2147 = m.ExcPending
	if v2147 != 0 {
		goto L1
	} else {
		goto L705
	}
L705:
	;
	v2149 = int32(1)
	v2152 = F_systable_beginscan(m, v2138, int32(828), v2149, int32(0), v2149, v2141)
	mBase = m.M
	v2153 = m.ExcPending
	if v2153 != 0 {
		goto L1
	} else {
		goto L707
	}
L706:
	;
	F_systable_endscan(m, v2152)
	mBase = m.M
	v2252 = m.ExcPending
	if v2252 != 0 {
		goto L1
	} else {
		goto L741
	}
L707:
	;
	v2154 = F_systable_getnext(m, v2152)
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L1
	} else {
		goto L708
	}
L708:
	;
	if v2154 == int32(0) {
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
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v2154)+16))
	v2175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2174)+22)))
	v2176 = v2174 + v2175
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v2176)+4))
	v2179 = F_GetUserNameFromId(m, v2177, int32(0))
	mBase = m.M
	v2180 = m.ExcPending
	if v2180 != 0 {
		goto L1
	} else {
		goto L716
	}
L712:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2161 = m.ExcPending
	if v2161 != 0 {
		goto L1
	} else {
		goto L713
	}
L713:
	;
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+688)) = v2162
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_86), v15+int32(688))
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L1
	} else {
		goto L714
	}
L714:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_87), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v2173 = m.ExcPending
	if v2173 != 0 {
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
	v2181 = F_quote_identifier(m, v2179)
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+736)) = v2181
	v2185 = v15 + int32(1168)
	F_appendStringInfo(m, v2185, int32(_a_F_getObjectIdentityParts_88), v15+int32(736))
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
		goto L1
	} else {
		goto L718
	}
L718:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v2176)+8))
	if v2191 != 0 {
		goto L719
	} else {
		goto L720
	}
L719:
	;
	v2192 = F_get_namespace_name_or_temp(m, v2191)
	mBase = m.M
	v2193 = m.ExcPending
	if v2193 != 0 {
		goto L1
	} else {
		goto L722
	}
L720:
	;
	v2202 = int32(0)
	goto L721
L721:
	;
	v2204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2176)+12)))
	switch v2204 - int32(76) {
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
		v2212 = int32(_a_F_getObjectIdentityParts_89)
		goto L726
	}
L722:
	;
	v2194 = F_quote_identifier(m, v2192)
	mBase = m.M
	v2195 = m.ExcPending
	if v2195 != 0 {
		goto L1
	} else {
		goto L723
	}
L723:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+720)) = v2194
	F_appendStringInfo(m, v2185, int32(_a_F_getObjectIdentityParts_90), v15+int32(720))
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L1
	} else {
		goto L724
	}
L724:
	;
	v2202 = v2192
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
	F_appendStringInfoString(m, v15+int32(1168), v2212)
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L1
	} else {
		goto L732
	}
L727:
	;
	v2212 = int32(_a_F_getObjectIdentityParts_91)
	goto L726
L728:
	;
	v2212 = int32(_a_F_getObjectIdentityParts_92)
	goto L726
L729:
	;
	v2212 = int32(_a_F_getObjectIdentityParts_93)
	goto L726
L730:
	;
	v2212 = int32(_a_F_getObjectIdentityParts_94)
	goto L726
L731:
	;
	v2212 = int32(_a_F_getObjectIdentityParts_95)
	goto L726
L732:
	;
	goto L725
L733:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+716)) = v2179
	*(*int32)(unsafe.Add(mBase, uint32(v15)+948)) = v2179
	v2225 = F_list_make1_impl(m, int32(1), v15+int32(716))
	mBase = m.M
	v2226 = m.ExcPending
	if v2226 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2225
	if v2202 != 0 {
		goto L735
	} else {
		goto L736
	}
L735:
	;
	v2228 = F_lappend(m, v2225, v2202)
	mBase = m.M
	v2229 = m.ExcPending
	if v2229 != 0 {
		goto L1
	} else {
		goto L738
	}
L736:
	;
	goto L737
L737:
	;
	v2231 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2176)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+704)) = v2231
	v2236 = F_psprintf(m, int32(_a_F_getObjectIdentityParts_96), v15+int32(704))
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L1
	} else {
		goto L739
	}
L738:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2228
	goto L737
L739:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+700)) = v2236
	*(*int32)(unsafe.Add(mBase, uint32(v15)+944)) = v2236
	v2243 = F_list_make1_impl(m, int32(1), v15+int32(700))
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		goto L1
	} else {
		goto L740
	}
L740:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2243
	goto L706
L741:
	;
	F_relation_close(m, v2138, int32(1))
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
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
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v2270 != 0 {
		goto L7
	} else {
		goto L745
	}
L745:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2274 = m.ExcPending
	if v2274 != 0 {
		goto L1
	} else {
		goto L746
	}
L746:
	;
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v2275
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1168))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v2277
	F_errmsg_internal(m, int32(_a_F_getObjectIdentityParts_97), v15+int32(16))
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L1
	} else {
		goto L747
	}
L747:
	;
	F_errfinish(m, int32(_a_F_getObjectIdentityParts_2), int32(_a_F_getObjectIdentityParts_98), int32(_a_F_getObjectIdentityParts_4))
	mBase = m.M
	v2288 = m.ExcPending
	if v2288 != 0 {
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
	v2328 = int32(0)
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
	F_errfinish(m, int32(_a_F_get_object_class_descr_5), int32(2777), int32(_a_F_get_object_class_descr_6))
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
