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
	var v516 int64
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int64
	_ = v532
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int64
	_ = v548
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v631 int64
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int64
	_ = v647
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
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
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int64
	_ = v769
	var v770 int32
	_ = v770
	var v771 int64
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v804 int64
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int64
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1065 int64
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1100 int64
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int64
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int64
	_ = v1134
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1152 int32
	_ = v1152
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
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
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L20
	} else {
		goto L303
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L20
	} else {
		goto L300
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L20
	} else {
		goto L297
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_btree_redo[0])) = v22
	v1182 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[1]))
	F_MemoryContextReset(m, v1182)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L20
	} else {
		goto L296
	}
L5:
	;
	v1159 = int32(0)
	F_btree_xlog_insert(m, int32(1), v1159, v1159, l0)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L20
	} else {
		goto L295
	}
L6:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L20
	} else {
		goto L292
	}
L7:
	;
	F__bt_restore_meta(m, l0, int32(0))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L20
	} else {
		goto L291
	}
L8:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[2]))
	if base.Ui32(v1126) < base.Ui32(int32(2)) {
		goto L4
	} else {
		goto L289
	}
L9:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v989 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v991 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L20
	} else {
		goto L253
	}
L10:
	;
	v769 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v771 = *(*int64)(unsafe.Add(mBase, uint32(v770)+16))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v770)+8))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v770)+4))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v770)))
	if v774 != 0 {
		goto L199
	} else {
		goto L200
	}
L11:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v647 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v651 = F_XLogReadBufferForRedo(m, l0, int32(1), v17+int32(-4))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L20
	} else {
		goto L174
	}
L12:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v532 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[2]))
	if base.Ui32(int32(2)) <= base.Ui32(v534) {
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
	F_errfinish(m, int32(_a_F_btree_redo_8), int32(526), int32(_a_F_btree_redo_9))
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
	*(*int64)(unsafe.Add(mBase, uint32(v113))) = base.I64_rotr(v57, int64(32))
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
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v526 == int32(0) {
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
	*(*uint32)(unsafe.Add(mBase, uint32(v491)+4)) = uint32(v433)
	v516 = int64(base.Ui64(v433) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v491))) = uint32(v516)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	F_MarkBufferDirty(m, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
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
	F_UnlockReleaseBuffer(m, v526)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L20
	} else {
		goto L138
	}
L138:
	;
	goto L4
L139:
	;
	v537 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v537, v17+int32(-20), v537, v537)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L20
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v556 = F_XLogReadBufferForRedo(m, l0, int32(0), v17+int32(-20))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L20
	} else {
		goto L144
	}
L142:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531)+8)))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v546
	v548 = *(*int64)(unsafe.Add(mBase, uint32(v19)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v548
	F_ResolveRecoveryConflictWithSnapshot(m, v545, v544, v19)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L20
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	if v556 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v560 = int32(0)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)+72))
	if v564 < v560 {
		v586 = v560
		goto L149
	} else {
		goto L150
	}
L146:
	;
	goto L147
L147:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v641 == int32(0) {
		goto L4
	} else {
		goto L172
	}
L148:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v590 < int32(0) {
		goto L160
	} else {
		goto L161
	}
L149:
	;
	goto L148
L150:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563+int32(0))+76)))
	if v569 != int32(1) {
		v586 = v560
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v573 = v563 + int32(76)
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+43)))
	if v574 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v586 = v560
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
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v573)+44))
	v586 = v584
	goto L149
L159:
	;
	v609 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v531)+6)))
	if v609 != 0 {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v594 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v594+(v590^int32(-1))<<(uint(int32(2))%32))))
	v608 = v600
	goto L159
L161:
	;
	goto L162
L162:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v608 = v602 + v590<<(uint(int32(13))%32) + int32(-8192)
	goto L159
L163:
	;
	v610 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v531)+4)))
	v611 = int32(1)
	v613 = v586 + v610<<(uint(v611)%32)
	F_btree_xlog_updates(m, v608, v613, v613+v609<<(uint(v611)%32), v609)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L20
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v620 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v531)+4)))
	if v620 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	goto L165
L167:
	;
	F_PageIndexMultiDelete(m, v608, v586, v620)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L20
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v608)+16)))
	v624 = v608 + v623
	v625 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v624)+12)))
	v627 = v625 & int32(_a_F_btree_redo_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v624)+12)) = uint16(v627)
	*(*uint32)(unsafe.Add(mBase, uint32(v608)+4)) = uint32(v532)
	v631 = int64(base.Ui64(v532) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v608))) = uint32(v631)
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	F_MarkBufferDirty(m, v633)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
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
	F_UnlockReleaseBuffer(m, v641)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L20
	} else {
		goto L173
	}
L173:
	;
	goto L4
L174:
	;
	if v651 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v655 < int32(0) {
		goto L179
	} else {
		goto L180
	}
L176:
	;
	goto L177
L177:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v708 != 0 {
		goto L184
	} else {
		goto L185
	}
L178:
	;
	v675 = v673 + int32(20)
	v676 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v646))))
	v677 = int32(2)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v675+v676<<(uint(v677)%32))))
	v681 = int32(_a_F_btree_redo_1)
	v687 = (v676 + int32(1)) & int32(_a_F_btree_redo_3)
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v675+v687<<(uint(v677)%32))))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v673+v691&v681)))
	*(*int32)(unsafe.Add(mBase, uint32(v680&v681+v673))) = v695
	F_PageIndexTupleDelete(m, v673, v687)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L20
	} else {
		goto L182
	}
L179:
	;
	v659 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v659+(v655^int32(-1))<<(uint(int32(2))%32))))
	v673 = v665
	goto L178
L180:
	;
	goto L181
L181:
	;
	v667 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v673 = v667 + v655<<(uint(int32(13))%32) + int32(-8192)
	goto L178
L182:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v673))) = base.I64_rotr(v647, int64(32))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_MarkBufferDirty(m, v702)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L20
	} else {
		goto L183
	}
L183:
	;
	goto L177
L184:
	;
	F_UnlockReleaseBuffer(m, v708)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L20
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v712 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L20
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v712
	if v712 < int32(0) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	F_PageInit(m, v732, int32(_a_F_btree_redo_4), int32(16))
	mBase = m.M
	goto L193
L190:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v718+(v712^int32(-1))<<(uint(int32(2))%32))))
	v732 = v724
	goto L189
L191:
	;
	goto L192
L192:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v732 = v726 + v712<<(uint(int32(13))%32) + int32(-8192)
	goto L189
L193:
	;
	v736 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v732)+16)))
	v737 = v732 + v736
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v646)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v737))) = v738
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v646)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v737)+8)) = int64(73014444032)
	*(*int32)(unsafe.Add(mBase, uint32(v737)+4)) = v740
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v646)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+46)) = uint16(v744)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(537395200)
	v749 = int32(base.Ui32(v744) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+44)) = uint16(v749)
	v756 = F_PageAddItemExtended(m, v732, v17+int32(-20), int32(8), int32(1), int32(0))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L20
	} else {
		goto L194
	}
L194:
	;
	if v756 == int32(0) {
		goto L2
	} else {
		goto L195
	}
L195:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v732))) = base.I64_rotr(v647, int64(32))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_MarkBufferDirty(m, v763)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L20
	} else {
		goto L196
	}
L196:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_UnlockReleaseBuffer(m, v766)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L20
	} else {
		goto L197
	}
L197:
	;
	goto L4
L198:
	;
	v813 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L20
	} else {
		goto L210
	}
L199:
	;
	v778 = F_XLogReadBufferForRedo(m, l0, int32(1), v17+int32(-4))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
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
	if v778 != 0 {
		goto L198
	} else {
		goto L203
	}
L203:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v780 < int32(0) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v799 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v798)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v799+v798)+4)) = v773
	*(*uint32)(unsafe.Add(mBase, uint32(v798)+4)) = uint32(v769)
	v804 = int64(base.Ui64(v769) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v798))) = uint32(v804)
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_MarkBufferDirty(m, v806)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L20
	} else {
		goto L208
	}
L205:
	;
	v784 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v784+(v780^int32(-1))<<(uint(int32(2))%32))))
	v798 = v790
	goto L204
L206:
	;
	goto L207
L207:
	;
	v792 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v798 = v792 + v780<<(uint(int32(13))%32) + int32(-8192)
	goto L204
L208:
	;
	goto L198
L209:
	;
	F_PageInit(m, v832, int32(_a_F_btree_redo_4), int32(16))
	mBase = m.M
	goto L214
L210:
	;
	if v813 < int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v818 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v818+(v813^int32(-1))<<(uint(int32(2))%32))))
	v832 = v824
	goto L209
L212:
	;
	goto L213
L213:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v832 = v826 + v813<<(uint(int32(13))%32) + int32(-8192)
	goto L209
L214:
	;
	v836 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v832)+16)))
	v837 = v832 + v836
	*(*int32)(unsafe.Add(mBase, uint32(v837)+8)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v837)+4)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v837))) = v774
	v841 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v832)+16)))
	v842 = v832 + v841
	v843 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v842)+12)))
	v847 = v843&int32(_a_F_btree_redo_11) | int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v842)+12)) = uint16(v847)
	v849 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v832)+12)) = uint16(v849)
	*(*int64)(unsafe.Add(mBase, uint32(v832)+24)) = v771
	v852 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v832)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v832)+14)) = uint16(v852)
	if v772 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v856 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v837)+12)))
	v858 = v856 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v837)+12)) = uint16(v858)
	goto L217
L216:
	;
	goto L217
L217:
	;
	v860 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v837)+14)) = uint16(v860)
	v862 = base.I32_wrap_i64(v769)
	*(*int32)(unsafe.Add(mBase, uint32(v832)+4)) = v862
	v866 = base.I32_wrap_i64(int64(base.Ui64(v769) >> (uint(int64(32)) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v832))) = v866
	F_MarkBufferDirty(m, v813)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L20
	} else {
		goto L218
	}
L218:
	;
	v873 = F_XLogReadBufferForRedo(m, l0, int32(2), v17+int32(-8))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L20
	} else {
		goto L219
	}
L219:
	;
	if v873 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	if v877 < int32(0) {
		goto L224
	} else {
		goto L225
	}
L221:
	;
	goto L222
L222:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v905 != 0 {
		goto L228
	} else {
		goto L229
	}
L223:
	;
	v896 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v895)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v896+v895))) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v895)+4)) = v862
	*(*int32)(unsafe.Add(mBase, uint32(v895))) = v866
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	F_MarkBufferDirty(m, v901)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L20
	} else {
		goto L227
	}
L224:
	;
	v881 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v881+(v877^int32(-1))<<(uint(int32(2))%32))))
	v895 = v887
	goto L223
L225:
	;
	goto L226
L226:
	;
	v889 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v895 = v889 + v877<<(uint(int32(13))%32) + int32(-8192)
	goto L223
L227:
	;
	goto L222
L228:
	;
	F_UnlockReleaseBuffer(m, v905)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L20
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	if v908 != 0 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L230
L232:
	;
	F_UnlockReleaseBuffer(m, v908)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L20
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	F_UnlockReleaseBuffer(m, v813)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L20
	} else {
		goto L236
	}
L235:
	;
	goto L234
L236:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v913)+72))
	if v914 < int32(3) {
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
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v913)+232)))
	if v917 != int32(1) {
		goto L237
	} else {
		goto L239
	}
L239:
	;
	v921 = F_XLogInitBufferForRedo(m, l0, int32(3))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L20
	} else {
		goto L241
	}
L240:
	;
	F_PageInit(m, v940, int32(_a_F_btree_redo_4), int32(16))
	mBase = m.M
	goto L245
L241:
	;
	if v921 < int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v926 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v926+(v921^int32(-1))<<(uint(int32(2))%32))))
	v940 = v932
	goto L240
L243:
	;
	goto L244
L244:
	;
	v934 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v940 = v934 + v921<<(uint(int32(13))%32) + int32(-8192)
	goto L240
L245:
	;
	v944 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v940)+16)))
	v945 = v940 + v944
	v946 = int32(17)
	*(*uint16)(unsafe.Add(mBase, uint32(v945)+12)) = uint16(v946)
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v770)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v945))) = v948
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v770)+28))
	v951 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v945)+14)) = uint16(v951)
	*(*int32)(unsafe.Add(mBase, uint32(v945)+8)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v945)+4)) = v950
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v770)+32))
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+46)) = uint16(v956)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(537395200)
	v961 = int32(base.Ui32(v956) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+44)) = uint16(v961)
	v968 = F_PageAddItemExtended(m, v940, v17+int32(-20), int32(8), int32(1), v951)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L20
	} else {
		goto L246
	}
L246:
	;
	if v968 == int32(0) {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v940)+4)) = v862
	*(*int32)(unsafe.Add(mBase, uint32(v940))) = v866
	F_MarkBufferDirty(m, v921)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L20
	} else {
		goto L248
	}
L248:
	;
	F_UnlockReleaseBuffer(m, v921)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
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
	v987 = m.ExcPending
	if v987 != 0 {
		goto L20
	} else {
		goto L251
	}
L251:
	;
	goto L4
L252:
	;
	F_PageInit(m, v1010, int32(_a_F_btree_redo_4), int32(16))
	mBase = m.M
	goto L257
L253:
	;
	if v991 < int32(0) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v996 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v996+(v991^int32(-1))<<(uint(int32(2))%32))))
	v1010 = v1002
	goto L252
L255:
	;
	goto L256
L256:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v1010 = v1004 + v991<<(uint(int32(13))%32) + int32(-8192)
	goto L252
L257:
	;
	v1014 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1010)+16)))
	v1015 = v1010 + v1014
	v1016 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1015)+12)) = uint16(v1016)
	*(*int64)(unsafe.Add(mBase, uint32(v1015))) = int64(0)
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v988)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1015)+8)) = v1020
	if v1020 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1024 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v1015)+12)) = uint16(v1024)
	goto L260
L259:
	;
	goto L260
L260:
	;
	v1026 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1015)+14)) = uint16(v1026)
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v988)+4))
	if v1028 == v1026 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1010))) = base.I64_rotr(v989, int64(32))
	F_MarkBufferDirty(m, v991)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L20
	} else {
		goto L286
	}
L262:
	;
	v1031 = int32(0)
	v1033 = v17 + int32(-4)
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+72))
	if v1036 < v1031 {
		v1058 = v1031
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F__bt_restore_page(m, v1010, v1061, v1062)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L20
	} else {
		goto L274
	}
L264:
	;
	v1061 = v1058
	goto L263
L265:
	;
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035+int32(0))+76)))
	if v1041 != int32(1) {
		v1058 = v1031
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v1045 = v1035 + int32(76)
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+43)))
	if v1046 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	if v1033 == int32(0) {
		v1058 = v1031
		goto L264
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	if v1033 != 0 {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1051 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1033))) = v1051
	v1061 = v1051
	goto L263
L271:
	;
	v1054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1045)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v1033))) = v1054
	goto L273
L272:
	;
	goto L273
L273:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+44))
	v1058 = v1056
	goto L264
L274:
	;
	v1065 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1069 = F_XLogReadBufferForRedo(m, l0, int32(1), v17+int32(-20))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L20
	} else {
		goto L275
	}
L275:
	;
	if v1069 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v1073 < int32(0) {
		goto L280
	} else {
		goto L281
	}
L277:
	;
	goto L278
L278:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v1107 == int32(0) {
		goto L261
	} else {
		goto L284
	}
L279:
	;
	v1092 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1091)+16)))
	v1093 = v1092 + v1091
	v1094 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1093)+12)))
	v1096 = v1094 & int32(_a_F_btree_redo_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v1093)+12)) = uint16(v1096)
	*(*uint32)(unsafe.Add(mBase, uint32(v1091)+4)) = uint32(v1065)
	v1100 = int64(base.Ui64(v1065) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1091))) = uint32(v1100)
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	F_MarkBufferDirty(m, v1102)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L20
	} else {
		goto L283
	}
L280:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[3]))
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1077+(v1073^int32(-1))<<(uint(int32(2))%32))))
	v1091 = v1083
	goto L279
L281:
	;
	goto L282
L282:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, _c_F_btree_redo[4]))
	v1091 = v1085 + v1073<<(uint(int32(13))%32) + int32(-8192)
	goto L279
L283:
	;
	goto L278
L284:
	;
	F_UnlockReleaseBuffer(m, v1107)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L20
	} else {
		goto L285
	}
L285:
	;
	goto L261
L286:
	;
	F_UnlockReleaseBuffer(m, v991)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L20
	} else {
		goto L287
	}
L287:
	;
	F__bt_restore_meta(m, l0, int32(2))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L20
	} else {
		goto L288
	}
L288:
	;
	goto L4
L289:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v1130 = *(*int64)(unsafe.Add(mBase, uint32(v1129)+16))
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129)+24)))
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v1132
	v1134 = *(*int64)(unsafe.Add(mBase, uint32(v1129)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v1134
	F_ResolveRecoveryConflictWithSnapshotFullXid(m, v1130, v1131, v17+int32(-48))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
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
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L20
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(_a_F_btree_redo_8), int32(1070), int32(_a_F_btree_redo_14))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
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
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L20
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(_a_F_btree_redo_8), int32(508), int32(_a_F_btree_redo_9))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
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
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L20
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(_a_F_btree_redo_8), int32(793), int32(_a_F_btree_redo_17))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
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
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L20
	} else {
		goto L304
	}
L304:
	;
	F_errfinish(m, int32(_a_F_btree_redo_8), int32(928), int32(_a_F_btree_redo_18))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
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
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
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
	v69 = F_XLogReadBufferForRedo(m, l3, int32(0), v13+int32(8))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
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
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v59 == int32(0) {
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
	*(*uint32)(unsafe.Add(mBase, uint32(v43)+4)) = uint32(v15)
	v52 = int64(base.Ui64(v15) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v43))) = uint32(v52)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_MarkBufferDirty(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
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
	F_UnlockReleaseBuffer(m, v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L1
L15:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L60
	}
L16:
	;
	if v69 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v73 = int32(0)
	v75 = v13 + int32(4)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l3)+96))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+72))
	if v78 < v73 {
		v100 = v73
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v196 != 0 {
		goto L52
	} else {
		goto L53
	}
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v104 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	v103 = v100
	goto L20
L22:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+int32(0))+76)))
	if v83 != int32(1) {
		v100 = v73
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v87 = v77 + int32(76)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+43)))
	if v88 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v75 == int32(0) {
		v100 = v73
		goto L21
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v75 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v93 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v93
	v103 = v93
	goto L20
L28:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v96
	goto L30
L29:
	;
	goto L30
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v87)+44))
	v100 = v98
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
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_insert[0]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108+(v104^int32(-1))<<(uint(int32(2))%32))))
	v122 = v114
	goto L31
L33:
	;
	goto L34
L34:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_insert[1]))
	v122 = v116 + v104<<(uint(int32(13))%32) + int32(-8192)
	goto L31
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v122))) = base.I64_rotr(v15, int64(32))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	F_MarkBufferDirty(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L3
	} else {
		goto L51
	}
L36:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17))))
	v128 = F_PageAddItemExtended(m, v122, v103, v125, v126, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103))))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v145 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v144 - v145
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17))))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v122+(v148-int32(1))&int32(_a_F_btree_xlog_insert_3)<<(uint(v145)%32))+20))
	v159 = F_CopyIndexTuple(m, v103+v145)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L44
	}
L39:
	;
	if v128 != 0 {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	F_errmsg_internal(m, int32(_a_F_btree_xlog_insert_0), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_insert_1), int32(191), int32(_a_F_btree_xlog_insert_2))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
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
	v163 = v122 + v156&int32(_a_F_btree_xlog_insert_4)
	v164 = F__bt_swap_posting(m, v159, v163, v143)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+6)))
	v172 = (v166&int32(_a_F_btree_xlog_insert_5) + int32(7)) & int32(_a_F_btree_xlog_insert_6)
	if v172 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	base.MemoryCopy(m, v163, v164, v172)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17))))
	v177 = F_PageAddItemExtended(m, v122, v159, v174, v175, int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	if v177 == int32(0) {
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
	F_UnlockReleaseBuffer(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
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
	v201 = m.ExcPending
	if v201 != 0 {
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
	v212 = m.ExcPending
	if v212 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_insert_1), int32(230), int32(_a_F_btree_xlog_insert_2))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
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
