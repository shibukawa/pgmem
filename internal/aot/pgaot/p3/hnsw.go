package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HnswBuildAppendPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v6 = F_HnswNewBuffer(m, l0, l3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		if v6 < int32(0) {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[0]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v11+(v6^int32(-1))<<(uint(int32(6))%32))+16))
			v26 = v17
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[1]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v19+v6<<(uint(int32(6))%32)+int32(-64))+16))
			v26 = v25
		}
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+16)))
		*(*int32)(unsafe.Add(mBase, uint32(v27+v28))) = v26
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		F_MarkBufferDirty(m, v31)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			F_UnlockReleaseBuffer(m, v34)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_LockBuffer(m, v6, int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[2]))
					if v41 != 0 {
						F_ProcessInterrupts(m)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_LockBuffer(m, v6, int32(2))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
								if v6 < int32(0) {
									v51 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[3]))
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(v6^int32(-1))<<(uint(int32(2))%32))))
									v65 = v57
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[4]))
									v65 = v59 + v6<<(uint(int32(13))%32) + int32(-8192)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65
								F_PageInit(m, v65, int32(_a_F_HnswBuildAppendPage_0), int32(8))
								mBase = m.M
								v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+16)))
								v71 = v65 + v70
								v72 = int32(_a_F_HnswBuildAppendPage_1)
								*(*uint16)(unsafe.Add(mBase, uint32(v71)+6)) = uint16(v72)
								*(*int32)(unsafe.Add(mBase, uint32(v71))) = int32(-1)
								return
							}
						}
					} else {
						F_LockBuffer(m, v6, int32(2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
							if v6 < int32(0) {
								v51 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[3]))
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(v6^int32(-1))<<(uint(int32(2))%32))))
								v65 = v57
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[4]))
								v65 = v59 + v6<<(uint(int32(13))%32) + int32(-8192)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65
							F_PageInit(m, v65, int32(_a_F_HnswBuildAppendPage_0), int32(8))
							mBase = m.M
							v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+16)))
							v71 = v65 + v70
							v72 = int32(_a_F_HnswBuildAppendPage_1)
							*(*uint16)(unsafe.Add(mBase, uint32(v71)+6)) = uint16(v72)
							*(*int32)(unsafe.Add(mBase, uint32(v71))) = int32(-1)
							return
						}
					}
				}
			}
		}
	}
}
func F_HnswInitElementFromBlock(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = l1
	v5 = F_palloc(m, int32(108))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*uint16)(unsafe.Add(mBase, uint32(v5)+80)) = uint16(v2)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+76)) = l0
		v11 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+88)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v5)+72)) = v11
		return v5
	}
}
func F_HnswInsertTupleOnDisk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 float64
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
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
	var v360 int32
	_ = v360
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
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v499 int32
	_ = v499
	var v513 int32
	_ = v513
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v533 int32
	_ = v533
	var v541 int32
	_ = v541
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
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
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v623 int32
	_ = v623
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v778 int32
	_ = v778
	var v795 int32
	_ = v795
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
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
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1030 int32
	_ = v1030
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1074 int32
	_ = v1074
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1195 int32
	_ = v1195
	var v1229 int32
	_ = v1229
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	v32 = m.G0
	v34 = v32 - int32(80)
	m.G0 = v34
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v37 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_LockPage(m, l0, int32(0), int32(5))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v42 = int32(64)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v42 = v41
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	F_HnswGetMetaPageInfo(m, l0, v34+int32(56), v34+int32(60))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	v59 = F_log(m, base.F64_convert_i32_s(v56))
	mBase = m.M
	v61 = int32(63)
	v63 = base.I32_div_u_s(int32(1358), v56)
	v65 = v63 - int32(2)
	if base.Ui32(v61) <= base.Ui32(v65) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v68 = v61
	goto L10
L9:
	;
	v68 = v65
	goto L10
L10:
	;
	v70 = F_HnswInitElement(m, int32(0), l3, v56, base.F64_div(float64(1), v59), v68, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+88)) = l2
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
	if v73 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	F_HnswFindElementNeighbors(m, v91, v70, v89, l0, l1, v92, v42, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L20
	}
L13:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+65)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+65)))
	if base.Ui32(v74) <= base.Ui32(v75) {
		v89 = v73
		v90 = int32(5)
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_UnlockPage(m, l0, int32(0), int32(5))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v81 = int32(7)
	F_LockPage(m, l0, int32(0), v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v86 = F_HnswGetEntryPoint(m, l0)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v86
	v89 = v86
	v90 = v81
	goto L12
L20:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v70)+72))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v100 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L5
	} else {
		goto L320
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L5
	} else {
		goto L317
	}
L23:
	;
	F_UnlockPage(m, l0, int32(0), v90)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L5
	} else {
		goto L316
	}
L24:
	;
	v287 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L5
	} else {
		goto L84
	}
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v70)+88))
	v105 = v70 + int32(4)
	v112 = int32(0)
	goto L26
L26:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v99+int32(8)+v112*int32(12))))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+88))
	v147 = F_datumIsEqual(m, v103, v144, int32(0), int32(-1))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L28
	}
L27:
	;
	goto L24
L28:
	;
	if v147 == int32(0) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v143)+76))
	v152 = F_ReadBuffer(m, l0, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_LockBuffer(m, v152, int32(2))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	if l4 != 0 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	F_UnlockReleaseBuffer(m, v152)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L5
	} else {
		goto L78
	}
L33:
	;
	v235 = v190 + v232*int32(6)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	*(*int32)(unsafe.Add(mBase, uint32(v235)+4)) = v236
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v235)+8)) = uint16(v238)
	if l4 != 0 {
		goto L73
	} else {
		goto L74
	}
L34:
	;
	v229 = base.B2i32(v191 == int32(0))
	if l4 != 0 {
		v245 = v229
		goto L32
	} else {
		goto L70
	}
L35:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+80)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v182+v183<<(uint(int32(2))%32))+20))
	v190 = v187&int32(_a_F_HnswInsertTupleOnDisk_0) + v182
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+8)))
	if v191 == int32(0) {
		goto L34
	} else {
		goto L44
	}
L36:
	;
	if v152 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v176 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L5
	} else {
		goto L42
	}
L39:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v161+(v152^int32(-1))<<(uint(int32(2))%32))))
	v181 = int32(0)
	v182 = v167
	goto L35
L40:
	;
	goto L41
L41:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v181 = int32(0)
	v182 = v170 + v152<<(uint(int32(13))%32) + int32(-8192)
	goto L35
L42:
	;
	v179 = F_GenericXLogRegisterBuffer(m, v176, v152, int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v181 = v176
	v182 = v179
	goto L35
L44:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+14)))
	if v194 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v232 = int32(1)
	goto L33
L46:
	;
	goto L47
L47:
	;
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+20)))
	if v198 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v232 = int32(2)
	goto L33
L49:
	;
	goto L50
L50:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+26)))
	if v202 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v232 = int32(3)
	goto L33
L52:
	;
	goto L53
L53:
	;
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+32)))
	if v206 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v232 = int32(4)
	goto L33
L55:
	;
	goto L56
L56:
	;
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+38)))
	if v210 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v232 = int32(5)
	goto L33
L58:
	;
	goto L59
L59:
	;
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+44)))
	if v214 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v232 = int32(6)
	goto L33
L61:
	;
	goto L62
L62:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+50)))
	if v218 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v232 = int32(7)
	goto L33
L64:
	;
	goto L65
L65:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+56)))
	if v222 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v232 = int32(8)
	goto L33
L67:
	;
	goto L68
L68:
	;
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+62)))
	if v226 != 0 {
		goto L34
	} else {
		goto L69
	}
L69:
	;
	v232 = int32(9)
	goto L33
L70:
	;
	F_pfree(m, v181)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	v245 = v229
	goto L32
L72:
	;
	v245 = int32(1)
	goto L32
L73:
	;
	F_MarkBufferDirty(m, v152)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_GenericXLogFinish(m, v181)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L5
	} else {
		goto L77
	}
L76:
	;
	goto L72
L77:
	;
	goto L72
L78:
	;
	if v245 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v250 = v191
	goto L81
L80:
	;
	v250 = int32(0)
	goto L81
L81:
	;
	if v250 != 0 {
		goto L23
	} else {
		goto L82
	}
L82:
	;
	v252 = v112 + int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v252 < v253 {
		v112 = v252
		goto L26
	} else {
		goto L83
	}
L83:
	;
	goto L27
L84:
	;
	F_LockBuffer(m, v287, int32(1))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	if v287 < int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+48))
	F_UnlockReleaseBuffer(m, v287)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L5
	} else {
		goto L90
	}
L87:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v295+(v287^int32(-1))<<(uint(int32(2))%32))))
	v309 = v301
	goto L86
L88:
	;
	goto L89
L89:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v309 = v303 + v287<<(uint(int32(13))%32) + int32(-8192)
	goto L86
L90:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v70)+88))
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	if v315 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v342 = F_add_size(m, int32(72), v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L5
	} else {
		goto L100
	}
L92:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	if base.Ui32((v319-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v341 = int32(6)
		goto L91
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v333 = int32(1)
	if v315&v333 != 0 {
		v341 = int32(base.Ui32(v315) >> (uint(v333) % 32))
		goto L91
	} else {
		goto L99
	}
L95:
	;
	v326 = int32(18)
	if v319&int32(255) == v326 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v332 = v326
	goto L98
L97:
	;
	v332 = int32(2)
	goto L98
L98:
	;
	v341 = v332
	goto L91
L99:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v341 = int32(base.Ui32(v337) >> (uint(int32(2)) % 32))
	goto L91
L100:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+65)))
	v348 = F_add_size(m, v346, int32(2))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	v350 = F_mul_size(m, v348, v97)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	v352 = F_mul_size(m, int32(6), v350)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	v354 = F_add_size(m, int32(4), v352)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	v360 = F_add_size(m, int32(0), int32(2))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	v362 = F_mul_size(m, v360, v97)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	v364 = F_mul_size(m, int32(6), v362)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v366 = F_add_size(m, int32(4), v364)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v372 = (v342 + int32(7)) & int32(-8)
	v373 = F_palloc0(m, v372)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	F_HnswSetElementTuple(m, int32(0), v373, v70)
	mBase = m.M
	v380 = (v354 + int32(7)) & int32(-8)
	v381 = F_palloc0(m, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v383 = int32(0)
	v395 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v381))) = uint8(v395)
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+65)))
	v409 = v401
	v411 = v383
	goto L112
L111:
	;
	v520 = int32(4)
	v524 = v372 + v380 | v520
	v533 = v310
	v541 = int32(-1)
	goto L139
L112:
	;
	goto L116
L113:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v381)+2)) = uint16(v499)
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+67)))
	*(*uint8)(unsafe.Add(mBase, uint32(v381)+1)) = uint8(v513)
	goto L111
L114:
	;
	if base.B2i32(v97 <= v383) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L116:
	;
	goto L117
L117:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v70)+72))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v429+v409<<(uint(int32(2))%32))))
	goto L114
L121:
	;
	v440 = int32(0)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	v452 = v411
	v453 = v440
	goto L124
L122:
	;
	v499 = v411
	goto L123
L123:
	;
	if int32(0) < v409 {
		v409 = v409 - int32(1)
		v411 = v499
		goto L112
	} else {
		goto L135
	}
L124:
	;
	v463 = v381 + int32(4) + v452*int32(6)
	if v453 < v443 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v499 = v486
	goto L123
L126:
	;
	v485 = int32(1)
	v486 = v452 + v485
	*(*uint16)(unsafe.Add(mBase, uint32(v463)+4)) = uint16(v483)
	*(*uint16)(unsafe.Add(mBase, uint32(v463)+2)) = uint16(v484)
	v490 = v453 + v485
	if v490 != v97<<(uint(base.B2i32(v409 == v440))%32) {
		v452 = v486
		v453 = v490
		goto L124
	} else {
		goto L134
	}
L127:
	;
	goto L131
L128:
	;
	goto L129
L129:
	;
	v479 = int32(_a_F_HnswInsertTupleOnDisk_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v463))) = uint16(v479)
	v483 = int32(0)
	v484 = v479
	goto L126
L130:
	;
	v474 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v470)+80)))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v470)+76))
	v477 = int32(base.Ui32(v475) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v463))) = uint16(v477)
	v483 = v474
	v484 = v475
	goto L126
L131:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v433+int32(8)+v453*int32(12))))
	goto L130
L134:
	;
	goto L125
L135:
	;
	goto L113
L136:
	;
	if v1003 < int32(0) {
		goto L258
	} else {
		goto L259
	}
L137:
	;
	v994 = int32(0)
	v999 = v966
	v1001 = v994
	v1003 = v970
	v1008 = v975
	v1009 = v976
	v1015 = v994
	goto L136
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v933
	*(*int32)(unsafe.Add(mBase, uint32(v34)+76)) = v937
	v966 = v933
	v970 = v937
	v975 = v942
	v976 = v943
	goto L137
L139:
	;
	v559 = F_ReadBuffer(m, l0, v533)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L5
	} else {
		goto L141
	}
L140:
	;
	F_HnswInsertAppendPage(m, l0, v34+int32(68), v34-int32(-64), v588, v589, l4)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L5
	} else {
		goto L237
	}
L141:
	;
	F_LockBuffer(m, v559, int32(2))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L5
	} else {
		goto L142
	}
L142:
	;
	if l4 != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	if v541 == int32(-1) {
		goto L152
	} else {
		goto L153
	}
L144:
	;
	if v559 < int32(0) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	goto L146
L146:
	;
	v583 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L5
	} else {
		goto L150
	}
L147:
	;
	v568 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v568+(v559^int32(-1))<<(uint(int32(2))%32))))
	v588 = int32(0)
	v589 = v574
	goto L143
L148:
	;
	goto L149
L149:
	;
	v577 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v588 = int32(0)
	v589 = v577 + v559<<(uint(int32(13))%32) + int32(-8192)
	goto L143
L150:
	;
	v586 = F_GenericXLogRegisterBuffer(m, v583, v559, int32(0))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L5
	} else {
		goto L151
	}
L151:
	;
	v588 = v583
	v589 = v586
	goto L143
L152:
	;
	v593 = int32(4)
	v594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+14)))
	v595 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+12)))
	v596 = v594 - v595
	if v596 <= v593 {
		goto L156
	} else {
		goto L157
	}
L153:
	;
	v604 = v541
	goto L154
L154:
	;
	v605 = int32(4)
	v606 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+14)))
	v607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+12)))
	v608 = v606 - v607
	if v608 <= v605 {
		goto L163
	} else {
		goto L164
	}
L155:
	;
	if base.Ui32(v599-int32(4)) < base.Ui32((v366+int32(7))&int32(-8)+v372|v520) {
		goto L159
	} else {
		goto L160
	}
L156:
	;
	v599 = v593
	goto L158
L157:
	;
	v599 = v596
	goto L158
L158:
	;
	goto L155
L159:
	;
	v603 = int32(-1)
	goto L161
L160:
	;
	v603 = v533
	goto L161
L161:
	;
	v604 = v603
	goto L154
L162:
	;
	if base.Ui32(v524) <= base.Ui32(v611-int32(4)) {
		v933 = v589
		v937 = v559
		v942 = v588
		v943 = v604
		goto L138
	} else {
		goto L166
	}
L163:
	;
	v611 = v605
	goto L165
L164:
	;
	v611 = v608
	goto L165
L165:
	;
	goto L162
L166:
	;
	v615 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+12)))
	if base.Ui32(v615) < base.Ui32(int32(25)) {
		v795 = v604
		goto L169
	} else {
		goto L170
	}
L167:
	;
	v860 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+16)))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v589+v860)))
	if v862 != int32(-1) {
		goto L229
	} else {
		goto L230
	}
L168:
	;
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670)+3)))
	if v559 != v727 {
		goto L218
	} else {
		goto L219
	}
L169:
	;
	if base.Ui32(v524) < base.Ui32(int32(_a_F_HnswInsertTupleOnDisk_2)) {
		goto L167
	} else {
		goto L210
	}
L170:
	;
	v623 = int32(base.Ui32(v615+int32(_a_F_HnswInsertTupleOnDisk_3))>>(uint(int32(2))%32)) & int32(_a_F_HnswInsertTupleOnDisk_1)
	if v623 == int32(0) {
		v795 = v604
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v634 = int32(1)
	v642 = v604
	goto L172
L172:
	;
	v666 = v634&int32(_a_F_HnswInsertTupleOnDisk_1)<<(uint(int32(2))%32) + (v589 + int32(24)) - int32(4)
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v666)))
	v670 = v589 + v667&int32(_a_F_HnswInsertTupleOnDisk_0)
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670))))
	if v671 != int32(1) {
		v772 = v642
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v795 = v772
	goto L169
L174:
	;
	v778 = v634 + int32(1)
	if base.Ui32(v778&int32(_a_F_HnswInsertTupleOnDisk_1)) <= base.Ui32(v623) {
		v634 = v778
		v642 = v772
		goto L172
	} else {
		goto L209
	}
L175:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670)+2)))
	if v674 == int32(0) {
		v772 = v642
		goto L174
	} else {
		goto L176
	}
L176:
	;
	if v559 < int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v696 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v670)+68)))
	v697 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v670)+66)))
	v698 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v670)+64)))
	v701 = v697 | v698<<(uint(int32(16))%32)
	if v701 == v695 {
		goto L182
	} else {
		goto L183
	}
L178:
	;
	v680 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[2]))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v680+(v559^int32(-1))<<(uint(int32(6))%32))+16))
	v695 = v686
	goto L177
L179:
	;
	goto L180
L180:
	;
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[3]))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v688+v559<<(uint(int32(6))%32)+int32(-64))+16))
	v695 = v694
	goto L177
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v728
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v666)))
	v733 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+14)))
	v734 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+12)))
	v735 = v733 - v734
	v736 = int32(0)
	if v736 < v735 {
		goto L191
	} else {
		goto L192
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+76)) = v559
	v727 = v559
	v728 = v589
	goto L181
L183:
	;
	goto L184
L184:
	;
	v704 = F_ReadBuffer(m, l0, v701)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L5
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+76)) = v704
	F_LockBuffer(m, v704, int32(2))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L5
	} else {
		goto L186
	}
L186:
	;
	if v704 < int32(0) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v713 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v713+(v704^int32(-1))<<(uint(int32(2))%32))))
	v727 = v704
	v728 = v719
	goto L181
L188:
	;
	goto L189
L189:
	;
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v727 = v704
	v728 = v721 + v704<<(uint(int32(13))%32) + int32(-8192)
	goto L181
L190:
	;
	v740 = int32(base.Ui32(v730)>>(uint(int32(17))%32)) + v739
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v696<<(uint(int32(2))%32)+v728)+20))
	v746 = int32(base.Ui32(v744) >> (uint(int32(17)) % 32))
	if v695 != v701 {
		goto L195
	} else {
		goto L196
	}
L191:
	;
	v739 = v735
	goto L193
L192:
	;
	v739 = v736
	goto L193
L193:
	;
	goto L190
L194:
	;
	if v642 == int32(-1) {
		goto L203
	} else {
		goto L204
	}
L195:
	;
	v748 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v728)+14)))
	v749 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v728)+12)))
	v750 = v748 - v749
	v751 = int32(0)
	if v751 < v750 {
		goto L199
	} else {
		goto L200
	}
L196:
	;
	goto L197
L197:
	;
	if base.Ui32(v740) < base.Ui32(v372) {
		v759 = v746
		goto L194
	} else {
		goto L202
	}
L198:
	;
	v759 = v754 + v746
	goto L194
L199:
	;
	v754 = v750
	goto L201
L200:
	;
	v754 = v751
	goto L201
L201:
	;
	goto L198
L202:
	;
	v759 = v740 - v372 + v746
	goto L194
L203:
	;
	v762 = v695
	goto L205
L204:
	;
	v762 = v642
	goto L205
L205:
	;
	if base.B2i32(base.Ui32(v372) <= base.Ui32(v740))&base.B2i32(base.Ui32(v380) <= base.Ui32(v759)) != 0 {
		goto L168
	} else {
		goto L206
	}
L206:
	;
	if v559 == v727 {
		v772 = v762
		goto L174
	} else {
		goto L207
	}
L207:
	;
	F_UnlockReleaseBuffer(m, v727)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L5
	} else {
		goto L208
	}
L208:
	;
	v772 = v762
	goto L174
L209:
	;
	goto L173
L210:
	;
	v813 = int32(4)
	v814 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+14)))
	v815 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+12)))
	v816 = v814 - v815
	if v816 <= v813 {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	if base.Ui32(v819-int32(4)) < base.Ui32(v372) {
		goto L167
	} else {
		goto L215
	}
L212:
	;
	v819 = v813
	goto L214
L213:
	;
	v819 = v816
	goto L214
L214:
	;
	goto L211
L215:
	;
	v823 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+16)))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v589+v823)))
	if v825 != int32(-1) {
		goto L167
	} else {
		goto L216
	}
L216:
	;
	F_HnswInsertAppendPage(m, l0, v34+int32(76), v34+int32(72), v588, v589, l4)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L5
	} else {
		goto L217
	}
L217:
	;
	v966 = v589
	v970 = v559
	v975 = v588
	v976 = v795
	goto L137
L218:
	;
	if l4 != 0 {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	goto L220
L220:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v373)+3)) = uint8(v834)
	*(*uint8)(unsafe.Add(mBase, uint32(v381)+1)) = uint8(v834)
	v999 = v589
	v1001 = v634
	v1003 = v559
	v1008 = v588
	v1009 = v762
	v1015 = v696
	goto L136
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v856
	goto L220
L222:
	;
	if v727 < int32(0) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L224
L224:
	;
	v854 = F_GenericXLogRegisterBuffer(m, v588, v727, int32(0))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L5
	} else {
		goto L228
	}
L225:
	;
	v839 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v839+(v727^int32(-1))<<(uint(int32(2))%32))))
	v856 = v845
	goto L221
L226:
	;
	goto L227
L227:
	;
	v847 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v856 = v847 + v727<<(uint(int32(13))%32) + int32(-8192)
	goto L221
L228:
	;
	v856 = v854
	goto L221
L229:
	;
	if l4 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L230:
	;
	goto L231
L231:
	;
	goto L140
L232:
	;
	F_pfree(m, v588)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L5
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	F_UnlockReleaseBuffer(m, v559)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L5
	} else {
		goto L236
	}
L235:
	;
	goto L234
L236:
	;
	v533 = v862
	v541 = v795
	goto L139
L237:
	;
	if l4 != 0 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v914 = int32(4)
	v915 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v913)+14)))
	v916 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v913)+12)))
	v917 = v915 - v916
	if v917 <= v914 {
		goto L252
	} else {
		goto L253
	}
L239:
	;
	F_MarkBufferDirty(m, v559)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L5
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	F_GenericXLogFinish(m, v588)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L5
	} else {
		goto L247
	}
L242:
	;
	F_UnlockReleaseBuffer(m, v559)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L5
	} else {
		goto L243
	}
L243:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	if v881 < int32(0) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v886 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v886+(v881^int32(-1))<<(uint(int32(2))%32))))
	v911 = v881
	v912 = int32(0)
	v913 = v892
	goto L238
L245:
	;
	goto L246
L246:
	;
	v895 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v911 = v881
	v912 = int32(0)
	v913 = v895 + v881<<(uint(int32(13))%32) + int32(-8192)
	goto L238
L247:
	;
	F_UnlockReleaseBuffer(m, v559)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L5
	} else {
		goto L248
	}
L248:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	v906 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L5
	} else {
		goto L249
	}
L249:
	;
	v909 = F_GenericXLogRegisterBuffer(m, v906, v905, int32(0))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L5
	} else {
		goto L250
	}
L250:
	;
	v911 = v905
	v912 = v906
	v913 = v909
	goto L238
L251:
	;
	if base.Ui32(v524) <= base.Ui32(v920-int32(4)) {
		v933 = v913
		v937 = v911
		v942 = v912
		v943 = v795
		goto L138
	} else {
		goto L255
	}
L252:
	;
	v920 = v914
	goto L254
L253:
	;
	v920 = v917
	goto L254
L254:
	;
	goto L251
L255:
	;
	F_HnswInsertAppendPage(m, l0, v34+int32(76), v34+int32(72), v912, v913, l4)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L5
	} else {
		goto L256
	}
L256:
	;
	v966 = v913
	v970 = v911
	v975 = v912
	v976 = v795
	goto L137
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+76)) = v1045
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	if v1047 < int32(0) {
		goto L262
	} else {
		goto L263
	}
L258:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[2]))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1030+(v1003^int32(-1))<<(uint(int32(6))%32))+16))
	v1045 = v1036
	goto L257
L259:
	;
	goto L260
L260:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[3]))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1038+v1003<<(uint(int32(6))%32)+int32(-64))+16))
	v1045 = v1044
	goto L257
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+84)) = v1066
	if base.Ui32(int32(2048)) <= base.Ui32((v1001-int32(1))&int32(_a_F_HnswInsertTupleOnDisk_1)) {
		goto L267
	} else {
		goto L268
	}
L262:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[2]))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1051+(v1047^int32(-1))<<(uint(int32(6))%32))+16))
	v1066 = v1057
	goto L261
L263:
	;
	goto L264
L264:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[3]))
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1059+v1047<<(uint(int32(6))%32)+int32(-64))+16))
	v1066 = v1065
	goto L261
L265:
	;
	if l4 != 0 {
		goto L291
	} else {
		goto L292
	}
L266:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v1155 = int32(0)
	v1157 = F_PageAddItemExtended(m, v1154, v381, v380, v1155, v1155)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L5
	} else {
		goto L288
	}
L267:
	;
	v1074 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v999)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1074) {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	goto L269
L269:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+82)) = uint16(v1015)
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+80)) = uint16(v1001)
	*(*uint16)(unsafe.Add(mBase, uint32(v373)+68)) = uint16(v1015)
	*(*uint16)(unsafe.Add(mBase, uint32(v373)+66)) = uint16(v1066)
	v1126 = int32(base.Ui32(v1066) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v373)+64)) = uint16(v1126)
	v1130 = F_PageIndexTupleOverwrite(m, v999, v1001&int32(_a_F_HnswInsertTupleOnDisk_1), v373, v372)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L5
	} else {
		goto L281
	}
L270:
	;
	v1082 = int32(base.Ui32(v1074+int32(_a_F_HnswInsertTupleOnDisk_3)) >> (uint(int32(2)) % 32))
	goto L272
L271:
	;
	v1082 = int32(0)
	goto L272
L272:
	;
	v1083 = int32(1)
	v1084 = v1082 + v1083
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+80)) = uint16(v1084)
	if v1003 != v1047 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1090 = v1083
	goto L275
L274:
	;
	v1090 = v1082 + int32(2)
	goto L275
L275:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+82)) = uint16(v1090)
	*(*uint16)(unsafe.Add(mBase, uint32(v373)+68)) = uint16(v1090)
	*(*uint16)(unsafe.Add(mBase, uint32(v373)+66)) = uint16(v1066)
	v1095 = int32(base.Ui32(v1066) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v373)+64)) = uint16(v1095)
	v1097 = int32(0)
	v1099 = F_PageAddItemExtended(m, v999, v373, v372, v1097, v1097)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L5
	} else {
		goto L276
	}
L276:
	;
	v1101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+80)))
	if v1099 == v1101 {
		goto L266
	} else {
		goto L277
	}
L277:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L5
	} else {
		goto L278
	}
L278:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v1107 + int32(4)
	F_errmsg_internal(m, int32(_a_F_HnswInsertTupleOnDisk_4), v34+int32(48))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L5
	} else {
		goto L279
	}
L279:
	;
	F_errfinish(m, int32(_a_F_HnswInsertTupleOnDisk_5), int32(325), int32(_a_F_HnswInsertTupleOnDisk_6))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L5
	} else {
		goto L280
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	if v1130 == int32(0) {
		goto L22
	} else {
		goto L282
	}
L282:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v1135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+82)))
	v1136 = F_PageIndexTupleOverwrite(m, v1134, v1135, v381, v380)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L5
	} else {
		goto L283
	}
L283:
	;
	if v1136 != 0 {
		goto L265
	} else {
		goto L284
	}
L284:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L5
	} else {
		goto L285
	}
L285:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v1142 + int32(4)
	F_errmsg_internal(m, int32(_a_F_HnswInsertTupleOnDisk_4), v34)
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L5
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(_a_F_HnswInsertTupleOnDisk_5), int32(320), int32(_a_F_HnswInsertTupleOnDisk_6))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L5
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	v1159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+82)))
	if v1157 != v1159 {
		goto L21
	} else {
		goto L289
	}
L289:
	;
	goto L265
L290:
	;
	if v1009 == int32(-1) {
		goto L298
	} else {
		goto L299
	}
L291:
	;
	F_MarkBufferDirty(m, v1003)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L5
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	F_GenericXLogFinish(m, v1008)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L5
	} else {
		goto L297
	}
L294:
	;
	if v1003 == v1047 {
		goto L290
	} else {
		goto L295
	}
L295:
	;
	F_MarkBufferDirty(m, v1047)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L5
	} else {
		goto L296
	}
L296:
	;
	goto L290
L297:
	;
	goto L290
L298:
	;
	v1171 = v1066
	goto L300
L299:
	;
	v1171 = v1009
	goto L300
L300:
	;
	F_UnlockReleaseBuffer(m, v1003)
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L5
	} else {
		goto L301
	}
L301:
	;
	if v1003 != v1047 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	F_UnlockReleaseBuffer(m, v1047)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L5
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	if v1171 == int32(-1) {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	goto L304
L306:
	;
	F_HnswUpdateNeighborsOnDisk(m, l0, l1, v70, v97, int32(0), l4)
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L5
	} else {
		goto L310
	}
L307:
	;
	if v1171 == v310 {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v1180 = int32(0)
	F_HnswUpdateMetaPage(m, l0, v1180, v1180, v1171, v1180, l4)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L5
	} else {
		goto L309
	}
L309:
	;
	goto L306
L310:
	;
	if v96 != 0 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+65)))
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+65)))
	if base.Ui32(v1188) <= base.Ui32(v1189) {
		goto L23
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	F_HnswUpdateMetaPage(m, l0, int32(1), v70, int32(-1), int32(0), l4)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L5
	} else {
		goto L315
	}
L314:
	;
	goto L313
L315:
	;
	goto L23
L316:
	;
	m.G0 = v34 + int32(80)
	return int32(1)
L317:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v1239 + int32(4)
	F_errmsg_internal(m, int32(_a_F_HnswInsertTupleOnDisk_4), v34+int32(16))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L5
	} else {
		goto L318
	}
L318:
	;
	F_errfinish(m, int32(_a_F_HnswInsertTupleOnDisk_5), int32(317), int32(_a_F_HnswInsertTupleOnDisk_6))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L5
	} else {
		goto L319
	}
L319:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L320:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v1257 + int32(4)
	F_errmsg_internal(m, int32(_a_F_HnswInsertTupleOnDisk_4), v34+int32(32))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L5
	} else {
		goto L321
	}
L321:
	;
	F_errfinish(m, int32(_a_F_HnswInsertTupleOnDisk_5), int32(328), int32(_a_F_HnswInsertTupleOnDisk_6))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L5
	} else {
		goto L322
	}
L322:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_HnswMemoryContextAlloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	v4 = F_MemoryContextAlloc(m, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+132)) = v12
		return v4
	}
}
func F_HnswNormValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_DirectFunctionCall1Coll(m, v4, l1, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_HnswSharedMemoryAlloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v8 = (l0 + int32(7)) & int32(-8)
	if base.Ui32(v8) < base.Ui32(int32(_a_F_HnswSharedMemoryAlloc_0)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+160))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
		v13 = F_add_size(m, v12, v8)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+160))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
			if base.Ui32(v18) < base.Ui32(v13) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_HnswSharedMemoryAlloc_1), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_HnswSharedMemoryAlloc_2), int32(673), int32(_a_F_HnswSharedMemoryAlloc_3))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v13
				return v21 + v20
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_HnswSharedMemoryAlloc_4), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_HnswSharedMemoryAlloc_2), int32(669), int32(_a_F_HnswSharedMemoryAlloc_3))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_HnswUpdateMetaPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v7 = int32(0)
	v12 = F_ReadBufferExtended(m, l0, l4, v7, v7, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		F_LockBuffer(m, v12, int32(2))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if l5 != 0 {
				if v12 < int32(0) {
					v20 = *(*int32)(unsafe.Add(mBase, _c_F_HnswUpdateMetaPage[0]))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v20+(v12^int32(-1))<<(uint(int32(2))%32))))
					v39 = v7
					v40 = v26
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_HnswUpdateMetaPage[1]))
					v39 = v7
					v40 = v28 + v12<<(uint(int32(13))%32) + int32(-8192)
				}
				if l1 == int32(0) {
				} else {
					if l2 == int32(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = int64(-281470681743361)
					} else {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+65)))
						if l1 != int32(2) {
							v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+46)))
							if v47 <= v50 {
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
								*(*int32)(unsafe.Add(mBase, uint32(v40)+40)) = v52
								v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+80)))
								*(*uint16)(unsafe.Add(mBase, uint32(v40)+46)) = uint16(v47)
								*(*uint16)(unsafe.Add(mBase, uint32(v40)+44)) = uint16(v54)
							}
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+40)) = v52
							v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+80)))
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+46)) = uint16(v47)
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+44)) = uint16(v54)
						}
					}
				}
				if l3 != int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = l3
				} else {
				}
				if l5 != 0 {
					F_MarkBufferDirty(m, v12)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						F_UnlockReleaseBuffer(m, v12)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					F_GenericXLogFinish(m, v39)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						F_UnlockReleaseBuffer(m, v12)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v34 = F_GenericXLogStart(m, l0)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v37 = F_GenericXLogRegisterBuffer(m, v34, v12, int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v39 = v34
						v40 = v37
						if l1 == int32(0) {
						} else {
							if l2 == int32(0) {
								*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = int64(-281470681743361)
							} else {
								v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+65)))
								if l1 != int32(2) {
									v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+46)))
									if v47 <= v50 {
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
										*(*int32)(unsafe.Add(mBase, uint32(v40)+40)) = v52
										v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+80)))
										*(*uint16)(unsafe.Add(mBase, uint32(v40)+46)) = uint16(v47)
										*(*uint16)(unsafe.Add(mBase, uint32(v40)+44)) = uint16(v54)
									}
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
									*(*int32)(unsafe.Add(mBase, uint32(v40)+40)) = v52
									v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+80)))
									*(*uint16)(unsafe.Add(mBase, uint32(v40)+46)) = uint16(v47)
									*(*uint16)(unsafe.Add(mBase, uint32(v40)+44)) = uint16(v54)
								}
							}
						}
						if l3 != int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = l3
						} else {
						}
						if l5 != 0 {
							F_MarkBufferDirty(m, v12)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							F_GenericXLogFinish(m, v39)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
