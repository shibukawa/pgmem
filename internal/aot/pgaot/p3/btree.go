package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_btree_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int64
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int64
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int64
	_ = v547
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
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
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int64
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v767 int64
	_ = v767
	var v768 int32
	_ = v768
	var v769 int64
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v860 int64
	_ = v860
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int64
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int64
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int64
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1125 int64
	_ = v1125
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1217 int32
	_ = v1217
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	v21 = int32(_a_F_btree_redo_0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[0]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+48)))
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_btree_redo[0])) = v27
	v30 = v24 & int32(240)
	switch int32(base.Ui32(v30)>>(uint(int32(4))%32)) - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	case 5:
		goto L14
	case 6:
		goto L12
	case 7, 8:
		goto L10
	case 9:
		goto L9
	case 10:
		goto L11
	case 11:
		goto L13
	case 12:
		goto L8
	case 13:
		goto L7
	case 14:
		goto L6
	default:
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L20
	} else {
		goto L303
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L20
	} else {
		goto L300
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L20
	} else {
		goto L297
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_btree_redo[0])) = v22
	v1173 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[1]))
	F_MemoryContextReset(m, v1173)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L20
	} else {
		goto L296
	}
L5:
	;
	v1150 = int32(0)
	F_btree_xlog_insert(m, int32(1), v1150, v1150, l0)
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L20
	} else {
		goto L295
	}
L6:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L20
	} else {
		goto L292
	}
L7:
	;
	F__bt_restore_meta(m, l0, int32(0))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L20
	} else {
		goto L291
	}
L8:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[2]))
	if base.Ui32(v1117) < base.Ui32(int32(2)) {
		goto L4
	} else {
		goto L289
	}
L9:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v981 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v983 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L20
	} else {
		goto L253
	}
L10:
	;
	v767 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v769 = *(*int64)(unsafe.Add(mBase, uint32(v768)+16))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v768)+8))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v768)+4))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v768)))
	if v772 != 0 {
		goto L199
	} else {
		goto L200
	}
L11:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v645 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v649 = F_XLogReadBufferForRedo(m, l0, int32(1), v17+int32(-4))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L20
	} else {
		goto L174
	}
L12:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v531 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[2]))
	if base.Ui32(int32(2)) <= base.Ui32(v533) {
		goto L139
	} else {
		goto L140
	}
L13:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v433 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v434 = int32(0)
	v439 = F_XLogReadBufferForRedoExtended(m, l0, v434, v434, int32(1), v17+int32(-20))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L20
	} else {
		goto L109
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v61 = F_XLogReadBufferForRedo(m, l0, int32(0), v17+int32(-20))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L26
	}
L15:
	;
	v51 = int32(1)
	F_btree_xlog_insert(m, v51, int32(0), v51, l0)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L20
	} else {
		goto L25
	}
L16:
	;
	F_btree_xlog_split(m, int32(0), l0)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L20
	} else {
		goto L24
	}
L17:
	;
	F_btree_xlog_split(m, int32(1), l0)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L20
	} else {
		goto L23
	}
L18:
	;
	v40 = int32(0)
	F_btree_xlog_insert(m, v40, int32(1), v40, l0)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L20
	} else {
		goto L22
	}
L19:
	;
	v35 = int32(0)
	F_btree_xlog_insert(m, v35, v35, v35, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	goto L4
L22:
	;
	goto L4
L23:
	;
	goto L4
L24:
	;
	goto L4
L25:
	;
	goto L4
L26:
	;
	if v61 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v65 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+72))
	if v69 < v65 {
		v91 = v65
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L29
L29:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v427 == int32(0) {
		goto L4
	} else {
		goto L107
	}
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v95 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L31:
	;
	goto L30
L32:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+int32(0))+76)))
	if v74 != int32(1) {
		v91 = v65
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v78 = v68 + int32(76)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+43)))
	if v79 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v91 = v65
	goto L31
L35:
	;
	goto L36
L36:
	;
	goto L39
L39:
	;
	goto L40
L40:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	v91 = v89
	goto L31
L41:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113)+16)))
	v116 = F_palloc(m, int32(1676))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L20
	} else {
		goto L45
	}
L42:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99+(v95^int32(-1))<<(uint(int32(2))%32))))
	v113 = v105
	goto L41
L43:
	;
	goto L44
L44:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v113 = v107 + v95<<(uint(int32(13))%32) + int32(-8192)
	goto L41
L45:
	;
	v118 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = v118
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+16)) = uint16(v118)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+12)) = v118
	*(*int64)(unsafe.Add(mBase, uint32(v116)+4)) = int64(11613591568384)
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v116))) = uint8(v126)
	v129 = F_palloc(m, int32(2704))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	v131 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v116)+28)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v116)+24)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v116)+36)) = v131
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113)+12)))
	v137 = v114 + v113
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v139 = F_PageGetTempPageCopySpecial(m, v113)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L20
	} else {
		goto L47
	}
L47:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v141 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v113)+24))
	v150 = F_PageAddItemExtended(m, v139, v113+v142&int32(_a_F_btree_redo_1), int32(base.Ui32(v142)>>(uint(int32(17))%32)), int32(1), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L20
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v138 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	if v150 == int32(0) {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v157 = int32(2)
	goto L55
L54:
	;
	v157 = int32(1)
	goto L55
L55:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v136) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v165 = int32(base.Ui32(v136+int32(_a_F_btree_redo_2)) >> (uint(int32(2)) % 32))
	goto L58
L57:
	;
	v165 = int32(0)
	goto L58
L58:
	;
	v167 = v165 & int32(_a_F_btree_redo_3)
	if base.Ui32(v157) <= base.Ui32(v167) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v171 = v157
	goto L62
L60:
	;
	goto L61
L61:
	;
	F__bt_dedup_finish_pending(m, v139, v116)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L20
	} else {
		goto L101
	}
L62:
	;
	v188 = v171 & int32(_a_F_btree_redo_3)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v113+int32(20)+v188<<(uint(int32(2))%32))))
	v195 = v113 + v192&int32(_a_F_btree_redo_1)
	if v157 != v188 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L61
L64:
	;
	v371 = v171 + int32(1)
	if base.Ui32(v371&int32(_a_F_btree_redo_3)) <= base.Ui32(v167) {
		v171 = v371
		goto L62
	} else {
		goto L100
	}
L65:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v116)+40))
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56))))
	if v198 <= v197 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v306 = v157
	goto L67
L67:
	;
	v308 = v306 & int32(_a_F_btree_redo_3)
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+7)))
	if v311&int32(32) != 0 {
		goto L93
	} else {
		goto L94
	}
L68:
	;
	F__bt_dedup_finish_pending(m, v139, v116)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L20
	} else {
		goto L89
	}
L69:
	;
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+16)))
	v203 = v91 + v197<<(uint(int32(2))%32)
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v203))))
	if v200 != v204 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v203)+2)))
	if v207 <= v206 {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v214 = int32(1)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+7)))
	if v215&int32(32) == int32(0) {
		v233 = v214
		v235 = v195
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if base.Ui32((v237+(v238+v233)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v236) {
		goto L64
	} else {
		goto L85
	}
L73:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v116)+28))
	v247 = base.B2i32(base.Ui32((v237+(v238+v233)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v236))
	if v247 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+4)))
	if v220&int32(_a_F_btree_redo_4) == int32(0) {
		v233 = v214
		v235 = v195
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+2)))
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195))))
	v233 = v220 & int32(4095)
	v235 = v227 + (v195 + v228<<(uint(int32(16))%32))
	goto L73
L76:
	;
	goto L72
L77:
	;
	v281 = v116 + v278
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	*(*int32)(unsafe.Add(mBase, uint32(v281))) = v282 + v280
	goto L76
L78:
	;
	if v238 <= int32(50) {
		goto L76
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+32)) = v254 + int32(1)
	v259 = v233 * int32(6)
	if v259 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v278 = int32(4)
	v280 = int32(1)
	goto L77
L82:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v116)+24))
	base.MemoryCopy(m, v260+v238*int32(6), v235, v259)
	goto L84
L83:
	;
	goto L84
L84:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v116)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = v265 + v233
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+6)))
	v278 = int32(36)
	v280 = (v269&int32(_a_F_btree_redo_5)+int32(7))&int32(_a_F_btree_redo_6) | int32(4)
	goto L77
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L20
	} else {
		goto L86
	}
L86:
	;
	F_errmsg_internal(m, int32(_a_F_btree_redo_7), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L20
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_btree_redo_8), int32(515), int32(_a_F_btree_redo_9))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L20
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v306 = v171
	goto L67
L90:
	;
	goto L64
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = v348
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+16)) = uint16(v308)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+12)) = v195
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+36)) = (v354&int32(_a_F_btree_redo_5)+int32(7))&int32(_a_F_btree_redo_6) | int32(4)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v116)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v116+v364<<(uint(int32(2))%32))+44)) = uint16(v308)
	goto L90
L92:
	;
	v329 = v314 & int32(4095)
	v331 = v329 * int32(6)
	if v331 != 0 {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+4)))
	if v314&int32(_a_F_btree_redo_4) != 0 {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v116)+24))
	v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v318)+4)) = uint16(v319)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	*(*int32)(unsafe.Add(mBase, uint32(v318))) = v321
	*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = int32(1)
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+6)))
	v348 = v325 & int32(_a_F_btree_redo_5)
	goto L91
L96:
	;
	goto L95
L97:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v116)+24))
	v333 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+2)))
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195))))
	base.MemoryCopy(m, v332, v333+(v195+v334<<(uint(int32(16))%32)), v331)
	goto L99
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = v329
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+2)))
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195))))
	v348 = v341 | v342<<(uint(int32(16))%32)
	goto L91
L100:
	;
	goto L63
L101:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+12)))
	if v393&int32(64) != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+16)))
	v397 = v139 + v396
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+12)))
	v400 = v398 & int32(_a_F_btree_redo_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v397)+12)) = uint16(v400)
	goto L104
L103:
	;
	goto L104
L104:
	;
	F_PageRestoreTempPage(m, v139, v113)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L20
	} else {
		goto L105
	}
L105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v113))) = base.I64_rotl(v57, int64(32))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	F_MarkBufferDirty(m, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L20
	} else {
		goto L106
	}
L106:
	;
	goto L29
L107:
	;
	F_UnlockReleaseBuffer(m, v427)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L20
	} else {
		goto L108
	}
L108:
	;
	goto L4
L109:
	;
	if v439 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v443 = int32(0)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)+72))
	if v447 < v443 {
		v469 = v443
		goto L114
	} else {
		goto L115
	}
L111:
	;
	goto L112
L112:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v525 == int32(0) {
		goto L4
	} else {
		goto L137
	}
L113:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v473 < int32(0) {
		goto L125
	} else {
		goto L126
	}
L114:
	;
	goto L113
L115:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446+int32(0))+76)))
	if v452 != int32(1) {
		v469 = v443
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v456 = v446 + int32(76)
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456)+43)))
	if v457 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v469 = v443
	goto L114
L118:
	;
	goto L119
L119:
	;
	goto L122
L122:
	;
	goto L123
L123:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v456)+44))
	v469 = v467
	goto L114
L124:
	;
	v492 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v432)+2)))
	if v492 != 0 {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v477+(v473^int32(-1))<<(uint(int32(2))%32))))
	v491 = v483
	goto L124
L126:
	;
	goto L127
L127:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v491 = v485 + v473<<(uint(int32(13))%32) + int32(-8192)
	goto L124
L128:
	;
	v493 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v432))))
	v494 = int32(1)
	v496 = v469 + v493<<(uint(v494)%32)
	F_btree_xlog_updates(m, v491, v496, v496+v492<<(uint(v494)%32), v492)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L20
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v432))))
	if v503 != 0 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	goto L130
L132:
	;
	F_PageIndexMultiDelete(m, v491, v469, v503)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L20
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v506 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v491)+16)))
	v507 = v491 + v506
	v508 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v507)+14)) = uint16(v508)
	v510 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v507)+12)))
	v512 = v510 & int32(_a_F_btree_redo_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v507)+12)) = uint16(v512)
	*(*int64)(unsafe.Add(mBase, uint32(v491))) = base.I64_rotl(v433, int64(32))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	F_MarkBufferDirty(m, v517)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L20
	} else {
		goto L136
	}
L135:
	;
	goto L134
L136:
	;
	goto L112
L137:
	;
	F_UnlockReleaseBuffer(m, v525)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L20
	} else {
		goto L138
	}
L138:
	;
	goto L4
L139:
	;
	v536 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v536, v17+int32(-20), v536, v536)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L20
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v555 = F_XLogReadBufferForRedo(m, l0, int32(0), v17+int32(-20))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L20
	} else {
		goto L144
	}
L142:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530)+8)))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v545
	v547 = *(*int64)(unsafe.Add(mBase, uint32(v19)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v547
	F_ResolveRecoveryConflictWithSnapshot(m, v544, v543, v19)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L20
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	if v555 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v559 = int32(0)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)+72))
	if v563 < v559 {
		v585 = v559
		goto L149
	} else {
		goto L150
	}
L146:
	;
	goto L147
L147:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v639 == int32(0) {
		goto L4
	} else {
		goto L172
	}
L148:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v589 < int32(0) {
		goto L160
	} else {
		goto L161
	}
L149:
	;
	goto L148
L150:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562+int32(0))+76)))
	if v568 != int32(1) {
		v585 = v559
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v572 = v562 + int32(76)
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+43)))
	if v573 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v585 = v559
	goto L149
L153:
	;
	goto L154
L154:
	;
	goto L157
L157:
	;
	goto L158
L158:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v572)+44))
	v585 = v583
	goto L149
L159:
	;
	v608 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v530)+6)))
	if v608 != 0 {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v593+(v589^int32(-1))<<(uint(int32(2))%32))))
	v607 = v599
	goto L159
L161:
	;
	goto L162
L162:
	;
	v601 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v607 = v601 + v589<<(uint(int32(13))%32) + int32(-8192)
	goto L159
L163:
	;
	v609 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v530)+4)))
	v610 = int32(1)
	v612 = v585 + v609<<(uint(v610)%32)
	F_btree_xlog_updates(m, v607, v612, v612+v608<<(uint(v610)%32), v608)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L20
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v619 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v530)+4)))
	if v619 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	goto L165
L167:
	;
	F_PageIndexMultiDelete(m, v607, v585, v619)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L20
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v622 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v607)+16)))
	v623 = v607 + v622
	v624 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v623)+12)))
	v626 = v624 & int32(_a_F_btree_redo_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v623)+12)) = uint16(v626)
	*(*int64)(unsafe.Add(mBase, uint32(v607))) = base.I64_rotl(v531, int64(32))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	F_MarkBufferDirty(m, v631)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L20
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	goto L147
L172:
	;
	F_UnlockReleaseBuffer(m, v639)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L20
	} else {
		goto L173
	}
L173:
	;
	goto L4
L174:
	;
	if v649 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v653 < int32(0) {
		goto L179
	} else {
		goto L180
	}
L176:
	;
	goto L177
L177:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v706 != 0 {
		goto L184
	} else {
		goto L185
	}
L178:
	;
	v673 = v671 + int32(20)
	v674 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v644))))
	v675 = int32(2)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v673+v674<<(uint(v675)%32))))
	v679 = int32(_a_F_btree_redo_1)
	v685 = (v674 + int32(1)) & int32(_a_F_btree_redo_3)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v673+v685<<(uint(v675)%32))))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v671+v689&v679)))
	*(*int32)(unsafe.Add(mBase, uint32(v678&v679+v671))) = v693
	F_PageIndexTupleDelete(m, v671, v685)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L20
	} else {
		goto L182
	}
L179:
	;
	v657 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v657+(v653^int32(-1))<<(uint(int32(2))%32))))
	v671 = v663
	goto L178
L180:
	;
	goto L181
L181:
	;
	v665 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v671 = v665 + v653<<(uint(int32(13))%32) + int32(-8192)
	goto L178
L182:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v671))) = base.I64_rotl(v645, int64(32))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_MarkBufferDirty(m, v700)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L20
	} else {
		goto L183
	}
L183:
	;
	goto L177
L184:
	;
	F_UnlockReleaseBuffer(m, v706)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L20
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v710 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L20
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v710
	if v710 < int32(0) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	F_PageInit(m, v730, int32(_a_F_btree_redo_4), int32(16))
	mBase = m.M
	goto L193
L190:
	;
	v716 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v716+(v710^int32(-1))<<(uint(int32(2))%32))))
	v730 = v722
	goto L189
L191:
	;
	goto L192
L192:
	;
	v724 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v730 = v724 + v710<<(uint(int32(13))%32) + int32(-8192)
	goto L189
L193:
	;
	v734 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v730)+16)))
	v735 = v730 + v734
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v644)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v735))) = v736
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v644)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v735)+8)) = int64(73014444032)
	*(*int32)(unsafe.Add(mBase, uint32(v735)+4)) = v738
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v644)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+46)) = uint16(v742)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(537395200)
	v747 = int32(base.Ui32(v742) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+44)) = uint16(v747)
	v754 = F_PageAddItemExtended(m, v730, v17+int32(-20), int32(8), int32(1), int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L20
	} else {
		goto L194
	}
L194:
	;
	if v754 == int32(0) {
		goto L2
	} else {
		goto L195
	}
L195:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v730))) = base.I64_rotl(v645, int64(32))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_MarkBufferDirty(m, v761)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L20
	} else {
		goto L196
	}
L196:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_UnlockReleaseBuffer(m, v764)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L20
	} else {
		goto L197
	}
L197:
	;
	goto L4
L198:
	;
	v810 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L20
	} else {
		goto L210
	}
L199:
	;
	v776 = F_XLogReadBufferForRedo(m, l0, int32(1), v17+int32(-4))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L20
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = int32(0)
	goto L198
L202:
	;
	if v776 != 0 {
		goto L198
	} else {
		goto L203
	}
L203:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v778 < int32(0) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v797 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v796)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v797+v796)+4)) = v771
	*(*int64)(unsafe.Add(mBase, uint32(v796))) = base.I64_rotl(v767, int64(32))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_MarkBufferDirty(m, v803)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L20
	} else {
		goto L208
	}
L205:
	;
	v782 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v782+(v778^int32(-1))<<(uint(int32(2))%32))))
	v796 = v788
	goto L204
L206:
	;
	goto L207
L207:
	;
	v790 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v796 = v790 + v778<<(uint(int32(13))%32) + int32(-8192)
	goto L204
L208:
	;
	goto L198
L209:
	;
	F_PageInit(m, v829, int32(_a_F_btree_redo_4), int32(16))
	mBase = m.M
	goto L214
L210:
	;
	if v810 < int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v815 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v815+(v810^int32(-1))<<(uint(int32(2))%32))))
	v829 = v821
	goto L209
L212:
	;
	goto L213
L213:
	;
	v823 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v829 = v823 + v810<<(uint(int32(13))%32) + int32(-8192)
	goto L209
L214:
	;
	v833 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v829)+16)))
	v834 = v829 + v833
	*(*int32)(unsafe.Add(mBase, uint32(v834)+8)) = v770
	*(*int32)(unsafe.Add(mBase, uint32(v834)+4)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v834))) = v772
	v838 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v829)+16)))
	v839 = v829 + v838
	v840 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v839)+12)))
	v844 = v840&int32(_a_F_btree_redo_11) | int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v839)+12)) = uint16(v844)
	v846 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v829)+12)) = uint16(v846)
	*(*int64)(unsafe.Add(mBase, uint32(v829)+24)) = v769
	v849 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v829)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v829)+14)) = uint16(v849)
	if v770 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v853 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v834)+12)))
	v855 = v853 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v834)+12)) = uint16(v855)
	goto L217
L216:
	;
	goto L217
L217:
	;
	v857 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v834)+14)) = uint16(v857)
	v860 = base.I64_rotl(v767, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v829))) = v860
	F_MarkBufferDirty(m, v810)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L20
	} else {
		goto L218
	}
L218:
	;
	v867 = F_XLogReadBufferForRedo(m, l0, int32(2), v17+int32(-8))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L20
	} else {
		goto L219
	}
L219:
	;
	if v867 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	if v871 < int32(0) {
		goto L224
	} else {
		goto L225
	}
L221:
	;
	goto L222
L222:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v898 != 0 {
		goto L228
	} else {
		goto L229
	}
L223:
	;
	v890 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v889)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v890+v889))) = v772
	*(*int64)(unsafe.Add(mBase, uint32(v889))) = v860
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	F_MarkBufferDirty(m, v894)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L20
	} else {
		goto L227
	}
L224:
	;
	v875 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v875+(v871^int32(-1))<<(uint(int32(2))%32))))
	v889 = v881
	goto L223
L225:
	;
	goto L226
L226:
	;
	v883 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v889 = v883 + v871<<(uint(int32(13))%32) + int32(-8192)
	goto L223
L227:
	;
	goto L222
L228:
	;
	F_UnlockReleaseBuffer(m, v898)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L20
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	if v901 != 0 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L230
L232:
	;
	F_UnlockReleaseBuffer(m, v901)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L20
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	F_UnlockReleaseBuffer(m, v810)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L20
	} else {
		goto L236
	}
L235:
	;
	goto L234
L236:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v906)+72))
	if v907 < int32(3) {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	if v30 != int32(144) {
		goto L4
	} else {
		goto L250
	}
L238:
	;
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v906)+232)))
	if v910 != int32(1) {
		goto L237
	} else {
		goto L239
	}
L239:
	;
	v914 = F_XLogInitBufferForRedo(m, l0, int32(3))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L20
	} else {
		goto L241
	}
L240:
	;
	F_PageInit(m, v933, int32(_a_F_btree_redo_4), int32(16))
	mBase = m.M
	goto L245
L241:
	;
	if v914 < int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v919 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v919+(v914^int32(-1))<<(uint(int32(2))%32))))
	v933 = v925
	goto L240
L243:
	;
	goto L244
L244:
	;
	v927 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v933 = v927 + v914<<(uint(int32(13))%32) + int32(-8192)
	goto L240
L245:
	;
	v937 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v933)+16)))
	v938 = v933 + v937
	v939 = int32(17)
	*(*uint16)(unsafe.Add(mBase, uint32(v938)+12)) = uint16(v939)
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v768)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v938))) = v941
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v768)+28))
	v944 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v938)+14)) = uint16(v944)
	*(*int32)(unsafe.Add(mBase, uint32(v938)+8)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v938)+4)) = v943
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v768)+32))
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+46)) = uint16(v949)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(537395200)
	v954 = int32(base.Ui32(v949) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+44)) = uint16(v954)
	v961 = F_PageAddItemExtended(m, v933, v17+int32(-20), int32(8), int32(1), v944)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L20
	} else {
		goto L246
	}
L246:
	;
	if v961 == int32(0) {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v933))) = v860
	F_MarkBufferDirty(m, v914)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L20
	} else {
		goto L248
	}
L248:
	;
	F_UnlockReleaseBuffer(m, v914)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L20
	} else {
		goto L249
	}
L249:
	;
	goto L237
L250:
	;
	F__bt_restore_meta(m, l0, int32(4))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L20
	} else {
		goto L251
	}
L251:
	;
	goto L4
L252:
	;
	F_PageInit(m, v1002, int32(_a_F_btree_redo_4), int32(16))
	mBase = m.M
	goto L257
L253:
	;
	if v983 < int32(0) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v988 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v988+(v983^int32(-1))<<(uint(int32(2))%32))))
	v1002 = v994
	goto L252
L255:
	;
	goto L256
L256:
	;
	v996 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v1002 = v996 + v983<<(uint(int32(13))%32) + int32(-8192)
	goto L252
L257:
	;
	v1006 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1002)+16)))
	v1007 = v1002 + v1006
	v1008 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1007)+12)) = uint16(v1008)
	*(*int64)(unsafe.Add(mBase, uint32(v1007))) = int64(0)
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v980)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1007)+8)) = v1012
	if v1012 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1016 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v1007)+12)) = uint16(v1016)
	goto L260
L259:
	;
	goto L260
L260:
	;
	v1018 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1007)+14)) = uint16(v1018)
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v980)+4))
	if v1020 == v1018 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1002))) = base.I64_rotl(v981, int64(32))
	F_MarkBufferDirty(m, v983)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L20
	} else {
		goto L286
	}
L262:
	;
	v1023 = int32(0)
	v1025 = v17 + int32(-4)
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+72))
	if v1028 < v1023 {
		v1050 = v1023
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F__bt_restore_page(m, v1002, v1053, v1054)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L20
	} else {
		goto L274
	}
L264:
	;
	v1053 = v1050
	goto L263
L265:
	;
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1027+int32(0))+76)))
	if v1033 != int32(1) {
		v1050 = v1023
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v1037 = v1027 + int32(76)
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1037)+43)))
	if v1038 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	if v1025 == int32(0) {
		v1050 = v1023
		goto L264
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	if v1025 != 0 {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1043 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1025))) = v1043
	v1053 = v1043
	goto L263
L271:
	;
	v1046 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1037)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v1025))) = v1046
	goto L273
L272:
	;
	goto L273
L273:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+44))
	v1050 = v1048
	goto L264
L274:
	;
	v1057 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1061 = F_XLogReadBufferForRedo(m, l0, int32(1), v17+int32(-20))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L20
	} else {
		goto L275
	}
L275:
	;
	if v1061 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v1065 < int32(0) {
		goto L280
	} else {
		goto L281
	}
L277:
	;
	goto L278
L278:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v1098 == int32(0) {
		goto L261
	} else {
		goto L284
	}
L279:
	;
	v1084 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1083)+16)))
	v1085 = v1084 + v1083
	v1086 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1085)+12)))
	v1088 = v1086 & int32(_a_F_btree_redo_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v1085)+12)) = uint16(v1088)
	*(*int64)(unsafe.Add(mBase, uint32(v1083))) = base.I64_rotl(v1057, int64(32))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	F_MarkBufferDirty(m, v1093)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L20
	} else {
		goto L283
	}
L280:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1069+(v1065^int32(-1))<<(uint(int32(2))%32))))
	v1083 = v1075
	goto L279
L281:
	;
	goto L282
L282:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v1083 = v1077 + v1065<<(uint(int32(13))%32) + int32(-8192)
	goto L279
L283:
	;
	goto L278
L284:
	;
	F_UnlockReleaseBuffer(m, v1098)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L20
	} else {
		goto L285
	}
L285:
	;
	goto L261
L286:
	;
	F_UnlockReleaseBuffer(m, v983)
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L20
	} else {
		goto L287
	}
L287:
	;
	F__bt_restore_meta(m, l0, int32(2))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L20
	} else {
		goto L288
	}
L288:
	;
	goto L4
L289:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v1121 = *(*int64)(unsafe.Add(mBase, uint32(v1120)+16))
	v1122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1120)+24)))
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v1123
	v1125 = *(*int64)(unsafe.Add(mBase, uint32(v1120)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v1125
	F_ResolveRecoveryConflictWithSnapshotFullXid(m, v1121, v1122, v17+int32(-48))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L20
	} else {
		goto L290
	}
L290:
	;
	goto L4
L291:
	;
	goto L4
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v30
	F_errmsg_internal(m, int32(_a_F_btree_redo_13), v17+int32(-32))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L20
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(_a_F_btree_redo_8), int32(1056), int32(_a_F_btree_redo_14))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L20
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	goto L4
L296:
	;
	m.G0 = v19 - int32(-64)
	return
L297:
	;
	F_errmsg_internal(m, int32(_a_F_btree_redo_15), int32(0))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L20
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(_a_F_btree_redo_8), int32(497), int32(_a_F_btree_redo_9))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L20
	} else {
		goto L299
	}
L299:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L300:
	;
	F_errmsg_internal(m, int32(_a_F_btree_redo_16), int32(0))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L20
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(_a_F_btree_redo_8), int32(780), int32(_a_F_btree_redo_17))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L20
	} else {
		goto L302
	}
L302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L303:
	;
	F_errmsg_internal(m, int32(_a_F_btree_redo_16), int32(0))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L20
	} else {
		goto L304
	}
L304:
	;
	F_errfinish(m, int32(_a_F_btree_redo_8), int32(914), int32(_a_F_btree_redo_18))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L20
	} else {
		goto L305
	}
L305:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_btree_xlog_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l3)+40))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)+96))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v68 = F_XLogReadBufferForRedo(m, l3, int32(0), v13+int32(8))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L16
	}
L2:
	;
	v21 = F_XLogReadBufferForRedo(m, l3, int32(1), v13+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v25 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v58 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+16)))
	v45 = v44 + v43
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
	v48 = v46 & int32(_a_F_btree_xlog_insert_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)) = uint16(v48)
	*(*int64)(unsafe.Add(mBase, uint32(v43))) = base.I64_rotl(v15, int64(32))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_MarkBufferDirty(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L12
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_insert[0]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v25^int32(-1))<<(uint(int32(2))%32))))
	v43 = v35
	goto L8
L10:
	;
	goto L11
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_insert[1]))
	v43 = v37 + v25<<(uint(int32(13))%32) + int32(-8192)
	goto L8
L12:
	;
	goto L7
L13:
	;
	F_UnlockReleaseBuffer(m, v58)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L1
L15:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L60
	}
L16:
	;
	if v68 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v72 = int32(0)
	v74 = v13 + int32(4)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)+96))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+72))
	if v77 < v72 {
		v99 = v72
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v195 != 0 {
		goto L52
	} else {
		goto L53
	}
L20:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v103 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	v102 = v99
	goto L20
L22:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+int32(0))+76)))
	if v82 != int32(1) {
		v99 = v72
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v86 = v76 + int32(76)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+43)))
	if v87 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v74 == int32(0) {
		v99 = v72
		goto L21
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v74 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v92 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v92
	v102 = v92
	goto L20
L28:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v95
	goto L30
L29:
	;
	goto L30
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	v99 = v97
	goto L21
L31:
	;
	if l2 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_insert[0]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107+(v103^int32(-1))<<(uint(int32(2))%32))))
	v121 = v113
	goto L31
L33:
	;
	goto L34
L34:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_insert[1]))
	v121 = v115 + v103<<(uint(int32(13))%32) + int32(-8192)
	goto L31
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v121))) = base.I64_rotl(v15, int64(32))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	F_MarkBufferDirty(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L3
	} else {
		goto L51
	}
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17))))
	v127 = F_PageAddItemExtended(m, v121, v102, v124, v125, int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102))))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v144 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v143 - v144
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17))))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v121+(v147-int32(1))&int32(_a_F_btree_xlog_insert_3)<<(uint(v144)%32))+20))
	v158 = F_CopyIndexTuple(m, v102+v144)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L44
	}
L39:
	;
	if v127 != 0 {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	F_errmsg_internal(m, int32(_a_F_btree_xlog_insert_0), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_insert_1), int32(188), int32(_a_F_btree_xlog_insert_2))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v162 = v121 + v155&int32(_a_F_btree_xlog_insert_4)
	v163 = F__bt_swap_posting(m, v158, v162, v142)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+6)))
	v171 = (v165&int32(_a_F_btree_xlog_insert_5) + int32(7)) & int32(_a_F_btree_xlog_insert_6)
	if v171 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	base.MemoryCopy(m, v162, v163, v171)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17))))
	v176 = F_PageAddItemExtended(m, v121, v158, v173, v174, int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	if v176 == int32(0) {
		goto L15
	} else {
		goto L50
	}
L50:
	;
	goto L35
L51:
	;
	goto L19
L52:
	;
	F_UnlockReleaseBuffer(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L3
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if l1 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	F__bt_restore_meta(m, l3, int32(2))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	m.G0 = v13 + int32(16)
	return
L59:
	;
	goto L58
L60:
	;
	F_errmsg_internal(m, int32(_a_F_btree_xlog_insert_7), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_insert_1), int32(226), int32(_a_F_btree_xlog_insert_2))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_btree_xlog_startup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_startup[0]))
	v8 = F_AllocSetContextCreateInternal(m, v3, int32(_a_F_btree_xlog_startup_0), int32(0), int32(_a_F_btree_xlog_startup_1), int32(_a_F_btree_xlog_startup_2))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_startup[1])) = v8
		return
	}
}
