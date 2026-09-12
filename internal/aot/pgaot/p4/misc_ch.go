package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckRequiredParameterValues(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[0])))
	if v3 != int32(1) {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[1]))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+172))
		if v8 == int32(0) {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_CheckRequiredParameterValues_0), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						F_errdetail(m, int32(_a_F_CheckRequiredParameterValues_1), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							F_errhint(m, int32(_a_F_CheckRequiredParameterValues_2), int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_CheckRequiredParameterValues_3), int32(_a_F_CheckRequiredParameterValues_4), int32(_a_F_CheckRequiredParameterValues_5))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			}
		} else {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[2])))
			if v12 != int32(1) {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[3]))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+180))
				F_RecoveryRequiresIntParameter(m, int32(_a_F_CheckRequiredParameterValues_6), v17, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[4]))
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[1]))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+184))
					F_RecoveryRequiresIntParameter(m, int32(_a_F_CheckRequiredParameterValues_7), v23, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[5]))
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[1]))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+188))
						F_RecoveryRequiresIntParameter(m, int32(_a_F_CheckRequiredParameterValues_8), v31, v34)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[6]))
							v41 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[1]))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+192))
							F_RecoveryRequiresIntParameter(m, int32(_a_F_CheckRequiredParameterValues_9), v39, v42)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[7]))
								v49 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[1]))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+196))
								F_RecoveryRequiresIntParameter(m, int32(_a_F_CheckRequiredParameterValues_10), v47, v50)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
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
func F_CheckRestrictedOperation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckRestrictedOperation[0])))
	if int32(base.Ui32(v8&int32(2))>>(uint(int32(1))%32)) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			F_errcode(m, int32(16797828))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
				F_errmsg(m, int32(_a_F_CheckRestrictedOperation_0), v5)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CheckRestrictedOperation_1), int32(466), int32(_a_F_CheckRestrictedOperation_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_CheckpointerMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v276 int32
	_ = v276
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v322 int32
	_ = v322
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v356 int32
	_ = v356
	var v368 int32
	_ = v368
	var v379 int32
	_ = v379
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v402 int32
	_ = v402
	var v414 int32
	_ = v414
	var v425 int32
	_ = v425
	var v434 int64
	_ = v434
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v535 int32
	_ = v535
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v619 int32
	_ = v619
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v665 int32
	_ = v665
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v731 int64
	_ = v731
	var v733 int64
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v767 int32
	_ = v767
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v804 int64
	_ = v804
	var v808 int32
	_ = v808
	var v810 int64
	_ = v810
	var v816 int32
	_ = v816
	var v818 int64
	_ = v818
	var v822 int32
	_ = v822
	var v824 int64
	_ = v824
	var v833 int32
	_ = v833
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v855 int32
	_ = v855
	var v864 int32
	_ = v864
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v886 int64
	_ = v886
	var v887 int32
	_ = v887
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v906 int64
	_ = v906
	var v907 int32
	_ = v907
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v945 int32
	_ = v945
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v966 int32
	_ = v966
	var v973 int32
	_ = v973
	var v975 int64
	_ = v975
	var v981 int32
	_ = v981
	var v983 int64
	_ = v983
	var v989 int64
	_ = v989
	var v995 int32
	_ = v995
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1032 int32
	_ = v1032
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1047 int64
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1051 int64
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1080 int64
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1111 int32
	_ = v1111
	var v1119 int32
	_ = v1119
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1153 int64
	_ = v1153
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1174 int32
	_ = v1174
	var v1181 int32
	_ = v1181
	var v1188 int32
	_ = v1188
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1248 int32
	_ = v1248
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1283 int32
	_ = v1283
	var v1295 int32
	_ = v1295
	var v1300 int32
	_ = v1300
	var v1301 int64
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	v3 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v23 = int32(-1)
	v24 = v3
	v25 = v3
	v26 = v3
	v29 = v3
	v30 = v3
	v33 = v19
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v23 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v1300 = int32(m.ExcTag)
	v1301 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1300 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L6:
	;
	v41 = v33 - int32(160)
	m.G0 = v41
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[0])) = int32(11)
	v46 = int32(1)
	v47 = v29 & v46
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	v50 = v30 & v46
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		v1295 = v41
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v467 = v24
	v468 = v25
	v469 = v26
	v470 = v33
	goto L8
L8:
	;
	if v468 != 0 {
		goto L120
	} else {
		goto L121
	}
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	v66 = int32(913)
	v68 = m.G0
	v70 = v68 - int32(144)
	m.G0 = v70
	switch int32(915) {
	case 0, 2:
		v80 = v66
		goto L11
	default:
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	v112 = int32(928)
	v114 = m.G0
	v116 = v114 - int32(144)
	m.G0 = v116
	switch int32(930) {
	case 0, 2:
		v126 = v112
		goto L24
	default:
		goto L25
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v80
	F_sigemptyset(m, v70+int32(8))
	mBase = m.M
	goto L14
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[3])) = v66
	v80 = int32(_a_F_CheckpointerMain_0)
	goto L11
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+136)) = int32(268435456)
	v92 = v70 + int32(4)
	goto L18
L16:
	;
	m.G0 = v70 + int32(144)
	goto L10
L18:
	;
	goto L19
L19:
	;
	if v92 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v103 = F___memcpy(m, int32(_a_F_CheckpointerMain_1), v92, int32(140))
	mBase = m.M
	goto L22
L21:
	;
	goto L22
L22:
	;
	goto L16
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	v158 = int32(-2)
	v160 = m.G0
	v162 = v160 - int32(144)
	m.G0 = v162
	switch int32(0) {
	case 0, 2:
		v172 = v158
		goto L37
	default:
		goto L38
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v126
	F_sigemptyset(m, v116+int32(8))
	mBase = m.M
	goto L27
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[4])) = v112
	v126 = int32(_a_F_CheckpointerMain_0)
	goto L24
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+136)) = int32(268435456)
	v138 = v116 + int32(4)
	goto L31
L29:
	;
	m.G0 = v116 + int32(144)
	goto L23
L31:
	;
	goto L32
L32:
	;
	if v138 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v149 = F___memcpy(m, int32(_a_F_CheckpointerMain_2), v138, int32(140))
	mBase = m.M
	goto L35
L34:
	;
	goto L35
L35:
	;
	goto L29
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	v204 = int32(-2)
	v206 = m.G0
	v208 = v206 - int32(144)
	m.G0 = v208
	switch int32(0) {
	case 0, 2:
		v218 = v204
		goto L50
	default:
		goto L51
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v172
	F_sigemptyset(m, v162+int32(8))
	mBase = m.M
	goto L40
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[5])) = v158
	v172 = int32(_a_F_CheckpointerMain_0)
	goto L37
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+136)) = int32(268435456)
	v184 = v162 + int32(4)
	goto L44
L42:
	;
	m.G0 = v162 + int32(144)
	goto L36
L44:
	;
	goto L45
L45:
	;
	if v184 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v195 = F___memcpy(m, int32(_a_F_CheckpointerMain_3), v184, int32(140))
	mBase = m.M
	goto L48
L47:
	;
	goto L48
L48:
	;
	goto L42
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	v250 = int32(-2)
	v252 = m.G0
	v254 = v252 - int32(144)
	m.G0 = v254
	switch int32(0) {
	case 0, 2:
		v264 = v250
		goto L63
	default:
		goto L64
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+4)) = v218
	F_sigemptyset(m, v208+int32(8))
	mBase = m.M
	goto L53
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[6])) = v204
	v218 = int32(_a_F_CheckpointerMain_0)
	goto L50
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+136)) = int32(268435456)
	v230 = v208 + int32(4)
	goto L57
L55:
	;
	m.G0 = v208 + int32(144)
	goto L49
L57:
	;
	goto L58
L58:
	;
	if v230 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v241 = F___memcpy(m, int32(_a_F_CheckpointerMain_4), v230, int32(140))
	mBase = m.M
	goto L61
L60:
	;
	goto L61
L61:
	;
	goto L55
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	v296 = int32(916)
	v298 = m.G0
	v300 = v298 - int32(144)
	m.G0 = v300
	switch int32(918) {
	case 0, 2:
		v310 = v296
		goto L76
	default:
		goto L77
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+4)) = v264
	F_sigemptyset(m, v254+int32(8))
	mBase = m.M
	goto L66
L64:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[7])) = v250
	v264 = int32(_a_F_CheckpointerMain_0)
	goto L63
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+136)) = int32(268435456)
	v276 = v254 + int32(4)
	goto L70
L68:
	;
	m.G0 = v254 + int32(144)
	goto L62
L70:
	;
	goto L71
L71:
	;
	if v276 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v287 = F___memcpy(m, int32(_a_F_CheckpointerMain_5), v276, int32(140))
	mBase = m.M
	goto L74
L73:
	;
	goto L74
L74:
	;
	goto L68
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	v342 = int32(915)
	v344 = m.G0
	v346 = v344 - int32(144)
	m.G0 = v346
	switch int32(917) {
	case 0, 2:
		v356 = v342
		goto L89
	default:
		goto L90
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v300)+4)) = v310
	F_sigemptyset(m, v300+int32(8))
	mBase = m.M
	goto L79
L77:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[8])) = v296
	v310 = int32(_a_F_CheckpointerMain_0)
	goto L76
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v300)+136)) = int32(268435456)
	v322 = v300 + int32(4)
	goto L83
L81:
	;
	m.G0 = v300 + int32(144)
	goto L75
L83:
	;
	goto L84
L84:
	;
	if v322 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v333 = F___memcpy(m, int32(_a_F_CheckpointerMain_6), v322, int32(140))
	mBase = m.M
	goto L87
L86:
	;
	goto L87
L87:
	;
	goto L81
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	v388 = int32(0)
	v390 = m.G0
	v392 = v390 - int32(144)
	m.G0 = v392
	switch int32(2) {
	case 0, 2:
		v402 = v388
		goto L102
	default:
		goto L103
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+4)) = v356
	F_sigemptyset(m, v346+int32(8))
	mBase = m.M
	goto L92
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[9])) = v342
	v356 = int32(_a_F_CheckpointerMain_0)
	goto L89
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+136)) = int32(268435456)
	v368 = v346 + int32(4)
	goto L96
L94:
	;
	m.G0 = v346 + int32(144)
	goto L88
L96:
	;
	goto L97
L97:
	;
	if v368 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v379 = F___memcpy(m, int32(_a_F_CheckpointerMain_7), v368, int32(140))
	mBase = m.M
	goto L100
L99:
	;
	goto L100
L100:
	;
	goto L94
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	v434 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10])) = v434
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[11])) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	F_before_shmem_exit(m, int32(929), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		v1295 = v41
		goto L5
	} else {
		goto L114
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v392)+4)) = v402
	F_sigemptyset(m, v392+int32(8))
	mBase = m.M
	goto L104
L103:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[12])) = v388
	v402 = int32(_a_F_CheckpointerMain_0)
	goto L102
L104:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v392)+136)) = int32(268435457)
	v414 = v392 + int32(4)
	goto L109
L107:
	;
	m.G0 = v392 + int32(144)
	goto L101
L109:
	;
	goto L110
L110:
	;
	if v414 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v425 = F___memcpy(m, int32(_a_F_CheckpointerMain_8), v414, int32(140))
	mBase = m.M
	goto L113
L112:
	;
	goto L113
L113:
	;
	goto L107
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	v451 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[13]))
	v456 = F_AllocSetContextCreateInternal(m, v451, int32(_a_F_CheckpointerMain_9), int32(0), int32(_a_F_CheckpointerMain_10), int32(_a_F_CheckpointerMain_11))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		v1295 = v41
		goto L5
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[14])) = v456
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v19 + int32(32)
	goto L119
L117:
	;
	v467 = v41
	v468 = int32(0)
	v469 = v456
	v470 = v41
	goto L8
L119:
	;
	goto L117
L120:
	;
	v472 = int32(_a_F_CheckpointerMain_12)
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[15]))
	v475 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[15])) = v474 + v475
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[16])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	v484 = v29 & v475
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	v487 = v30 & v475
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	F_EmitErrorReport(m)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[17])) = v467
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	v628 = int32(1)
	v629 = v29 & v628
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v629)
	v632 = v30 & v628
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v632)
	F_sigprocmask(m, int32(_a_F_CheckpointerMain_13), int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L143
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	F_LWLockReleaseAll(m)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L125
	}
L125:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	F_UnlockBuffers(m)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	F_smgrdestroyall(m)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L131
	}
L131:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[19])))
	if v551 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v553)+4)) = int32(1)
	if v554 != 0 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[14])) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	F_FlushErrorState(m)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L140
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	v562 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	F_s_lock(m, v562+int32(4), int32(_a_F_CheckpointerMain_14), int32(294), int32(_a_F_CheckpointerMain_15))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v571 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v571)+4)) = int32(0)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v571)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v571)+12)) = v574
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v571)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v571)+16)) = v576 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	F_ConditionVariableBroadcast(m, v571+int32(36))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L139
	}
L138:
	;
	goto L137
L139:
	;
	v589 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[19])) = uint8(v589)
	goto L134
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	F_MemoryContextReset(m, v469)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L141
	}
L141:
	;
	v607 = int32(_a_F_CheckpointerMain_12)
	v609 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[15]))
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[15])) = v609 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	F_pg_usleep(m, int32(_a_F_CheckpointerMain_16))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L142
	}
L142:
	;
	goto L122
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v632)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v629)
	F_SyncRepUpdateSyncStandbysDefined(m)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v632)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v629)
	F_UpdateFullPageWrites(m)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v632)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v629)
	v656 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L146
	}
L146:
	;
	if v656 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v632)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v629)
	F_errmsg_internal(m, int32(_a_F_CheckpointerMain_17), int32(0))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v676 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[20]))
	v678 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v676)+64)) = v678
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v629)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v632)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v685))) = int32(0)
	goto L152
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v632)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v629)
	F_errfinish(m, int32(_a_F_CheckpointerMain_14), int32(1390), int32(_a_F_CheckpointerMain_18))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v632)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v629)
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v632)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v629)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L154
	}
L154:
	;
	v701 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[23]))
	if v701 != 0 {
		v1137 = v29
		v1138 = v30
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v1147 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[24])) = uint8(v1147)
	v1150 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[23]))
	if v1150 != 0 {
		goto L239
	} else {
		goto L240
	}
L156:
	;
	v709 = v29
	v710 = v30
	goto L157
L157:
	;
	v719 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[25]))
	if v719 != 0 {
		v1137 = v709
		v1138 = v710
		goto L155
	} else {
		goto L159
	}
L158:
	;
	v1137 = v1011
	v1138 = v1012
	goto L155
L159:
	;
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v721)+20))
	v723 = int32(1)
	v724 = v710 & v723
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v724)
	v727 = v709 & v723
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v727)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	v731 = F___time(m)
	mBase = m.M
	v733 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10]))
	v735 = base.I32_wrap_i64(v731 - v733)
	v737 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[26]))
	v738 = base.B2i32(v737 <= v735)
	if v738|v722 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v724)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v727)
	v746 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[27])))
	if v746 == int32(1) {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	v1011 = v709
	v1012 = v710
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	v1019 = int32(1)
	v1020 = v1012 & v1019
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1020)
	v1023 = v1011 & v1019
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1023)
	F_CheckArchiveTimeout(m)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L217
	}
L163:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v758)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v758)+4)) = int32(1)
	if v759 != 0 {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	v751 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[28]))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v751)+316))
	v754 = base.B2i32(v752 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[27])) = uint8(v754)
	v756 = v754
	goto L166
L165:
	;
	v756 = int32(0)
	goto L166
L166:
	;
	goto L163
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v727)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v724)
	v767 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	F_s_lock(m, v767+int32(4), int32(_a_F_CheckpointerMain_14), int32(414), int32(_a_F_CheckpointerMain_15))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v776 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v777 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v776)+4)) = v777
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v776)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v776)+20)) = v777
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v776)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v776)+8)) = v782 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v724)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v727)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	F_ConditionVariableBroadcast(m, v776+int32(24))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	v796 = int32(0)
	v798 = base.B2i32(v779&int32(2) == v796) & v756
	if base.B2i32(v722 == v796)&v738 != 0 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	if v798 != 0 {
		goto L183
	} else {
		goto L184
	}
L173:
	;
	if v798 != 0 {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	goto L175
L175:
	;
	if v722 == int32(0) {
		goto L172
	} else {
		goto L179
	}
L176:
	;
	v802 = int32(_a_F_CheckpointerMain_19)
	v804 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[29]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[29])) = v804 + int64(1)
	goto L172
L177:
	;
	goto L178
L178:
	;
	v808 = int32(_a_F_CheckpointerMain_20)
	v810 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[30]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[30])) = v810 + int64(1)
	goto L172
L179:
	;
	if v798 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v816 = int32(_a_F_CheckpointerMain_21)
	v818 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[31]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[31])) = v818 + int64(1)
	goto L172
L181:
	;
	goto L182
L182:
	;
	v822 = int32(_a_F_CheckpointerMain_22)
	v824 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[32]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[32])) = v824 + int64(1)
	goto L172
L183:
	;
	v876 = v779 | v738<<(uint(int32(8))%32)
	v878 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[19])) = uint8(v878)
	if v798 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L184:
	;
	if v779&int32(128) == int32(0) {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v833 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[33]))
	if v833 <= v735 {
		goto L183
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v724)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v727)
	v841 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L187
	}
L187:
	;
	if v841 == int32(0) {
		goto L183
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v724)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v727)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v735
	F_errmsg_plural(m, int32(_a_F_CheckpointerMain_23), int32(_a_F_CheckpointerMain_24), v735, v19+int32(16))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v724)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v727)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(_a_F_CheckpointerMain_25)
	F_errhint(m, int32(_a_F_CheckpointerMain_26), v19)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v724)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v727)
	F_errfinish(m, int32(_a_F_CheckpointerMain_14), int32(462), int32(_a_F_CheckpointerMain_15))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L191
	}
L191:
	;
	goto L183
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	v927 = int32(1)
	v928 = v922 & v927
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v928)
	v931 = v921 & v927
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v931)
	F_smgrdestroyall(m)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L200
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v724)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v727)
	v886 = F_GetInsertRecPtr(m)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v724)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v727)
	v906 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L198
	}
L196:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[34])) = v731
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[35])) = v886
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[36])) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v727)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v724)
	v899 = F_CreateCheckPoint(m, v876)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L197
	}
L197:
	;
	v921 = v709
	v922 = v899
	v924 = v899
	goto L192
L198:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[34])) = v731
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[35])) = v906
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[36])) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v727)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v724)
	v919 = F_CreateRestartPoint(m, v876)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L199
	}
L199:
	;
	v921 = v919
	v922 = v710
	v924 = v919
	goto L192
L200:
	;
	v936 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v936)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v936)+4)) = int32(1)
	if v937 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v931)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v928)
	v945 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	F_s_lock(m, v945+int32(4), int32(_a_F_CheckpointerMain_14), int32(495), int32(_a_F_CheckpointerMain_15))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v954 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v954)+4)) = int32(0)
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v954)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v954)+12)) = v957
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v928)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v931)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	F_ConditionVariableBroadcast(m, v954+int32(36))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	if v798 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v995 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[19])) = uint8(v995)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v931)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v928)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L214
	}
L207:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10])) = v731
	if v924 == int32(0) {
		goto L206
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	if v924 != 0 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v973 = int32(_a_F_CheckpointerMain_27)
	v975 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[37]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[37])) = v975 + int64(1)
	goto L206
L211:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10])) = v731
	v981 = int32(_a_F_CheckpointerMain_28)
	v983 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[38]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[38])) = v983 + int64(1)
	goto L206
L212:
	;
	goto L213
L213:
	;
	v989 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[26])))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10])) = v731 - v989 + int64(15)
	goto L206
L214:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[23]))
	if v1004 != 0 {
		v1137 = v921
		v1138 = v922
		goto L155
	} else {
		goto L215
	}
L215:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[25]))
	if v1006 != 0 {
		v1137 = v921
		v1138 = v922
		goto L155
	} else {
		goto L216
	}
L216:
	;
	v1011 = v921
	v1012 = v922
	goto L162
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1020)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1023)
	F_pgstat_report_checkpointer(m)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1020)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1023)
	F_pgstat_report_wal(m, int32(1))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L219
	}
L219:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+20))
	if v1042 != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1023)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1020)
	v1111 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1111))) = int32(0)
	goto L235
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1020)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1023)
	v1047 = F___time(m)
	mBase = m.M
	v1049 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[26]))
	v1051 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10]))
	v1053 = base.I32_wrap_i64(v1047 - v1051)
	if v1049 <= v1053 {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v1055 = v1049 - v1053
	v1057 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[39]))
	if v1057 <= int32(0) {
		v1087 = v1055
		goto L223
	} else {
		goto L224
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1023)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1020)
	v1095 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	v1100 = F_WaitLatch(m, v1095, int32(41), v1087*int32(1000), int32(83886084))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L234
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1020)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1023)
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[27])))
	if v1066 == int32(1) {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	if v1076 != 0 {
		v1087 = v1055
		goto L223
	} else {
		goto L229
	}
L226:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[28]))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+316))
	v1074 = base.B2i32(v1072 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[27])) = uint8(v1074)
	v1076 = v1074
	goto L228
L227:
	;
	v1076 = int32(0)
	goto L228
L228:
	;
	goto L225
L229:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[39]))
	v1080 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[11]))
	v1082 = base.I32_wrap_i64(v1047 - v1080)
	if v1078 <= v1082 {
		goto L220
	} else {
		goto L230
	}
L230:
	;
	v1084 = v1078 - v1082
	if v1055 < v1084 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1086 = v1055
	goto L233
L232:
	;
	v1086 = v1084
	goto L233
L233:
	;
	v1087 = v1086
	goto L223
L234:
	;
	goto L220
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1020)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1023)
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1020)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1023)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L237
	}
L237:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[23]))
	if v1127 == int32(0) {
		v709 = v1011
		v710 = v1012
		goto L157
	} else {
		goto L238
	}
L238:
	;
	goto L158
L239:
	;
	v1151 = int32(_a_F_CheckpointerMain_22)
	v1153 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[32]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[32])) = v1153 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	v1159 = int32(1)
	v1160 = v1137 & v1159
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1160)
	v1163 = v1138 & v1159
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1163)
	v1165 = int32(0)
	F_ShutdownXLOG(m, v1165, v1165)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	v1196 = int32(1)
	v1197 = v1137 & v1196
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1197)
	v1200 = v1138 & v1196
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1200)
	v1203 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1203))) = int32(0)
	goto L246
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1163)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1160)
	F_pgstat_report_checkpointer(m)
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1163)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1160)
	F_pgstat_report_wal(m, int32(1))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1163)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1160)
	F_SendPostmasterSignal(m, int32(9))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[23])) = int32(0)
	goto L241
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1200)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1197)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L247
	}
L247:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[25]))
	if v1213 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	goto L251
L249:
	;
	goto L250
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1200)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1197)
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L257
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1197)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1200)
	v1237 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	v1241 = F_WaitLatch(m, v1237, int32(33), int32(0), int32(83886085))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L253
	}
L252:
	;
	goto L250
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1197)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1200)
	v1248 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1248))) = int32(0)
	goto L254
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v1200)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v1197)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		v1295 = v470
		goto L5
	} else {
		goto L255
	}
L255:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[25]))
	if v1258 == int32(0) {
		goto L251
	} else {
		goto L256
	}
L256:
	;
	goto L252
L257:
	;
	goto L4
L258:
	;
	v1305 = int32(v1301)
	m.G0 = v1295
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+4))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1305)))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1308)))
	if v19+int32(32) == v1312 {
		goto L261
	} else {
		goto L262
	}
L259:
	;
	m.ExcPending = 1
	goto L267
L260:
	;
	if v1315 != 0 {
		goto L264
	} else {
		goto L265
	}
L261:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+4))
	v1315 = v1314
	goto L263
L262:
	;
	v1315 = int32(0)
	goto L263
L263:
	;
	goto L260
L264:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v1318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)))
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)))
	v23 = v1315
	v24 = v1316
	v25 = v1307
	v26 = v1317
	v29 = v1319
	v30 = v1318
	v33 = v1295
	goto L1
L265:
	;
	goto L266
L266:
	;
	F___wasm_longjmp(m, v1308, v1307)
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	return
L268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_char_increment(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	v4 = int32(255)
	v7 = base.B2i32(l1&v4 == v4)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v7)
	v10 = int32(24)
	if l1&v4 == v4 {
		v16 = int32(0)
	} else {
		v16 = (l1<<(uint(v10)%32) + int32(16777216)) >> (uint(v10) % 32)
	}
	return v16
}
func F_char_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if base.I32_extend8_s(v3) < int32(0) {
			v12 = int32(92)
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v12)
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(32)
			v16 = int32(7)
			v18 = int32(48)
			v19 = v3&v16 | v18
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+7)) = uint8(v19)
			v26 = int32(base.Ui32(v3)>>(uint(int32(3))%32))&v16 | v18
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)) = uint8(v26)
			v33 = int32(base.Ui32(v3&int32(192))>>(uint(int32(6))%32)) | v18
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)) = uint8(v33)
			return v5
		} else {
			if v3&int32(255) != 0 {
				*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v3)
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(20)
				return v5
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(16)
				return v5
			}
		}
	}
}
func F_chareq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 == v3)
}
func F_chareqfast(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	v3 = int32(255)
	return base.B2i32(l0&v3 == l1&v3)
}
func F_charrecv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pq_getmsgbyte(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return base.I32_extend8_s(v3)
	}
}
func F_charsend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_enlargeStringInfo(m, v6, int32(1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*uint8)(unsafe.Add(mBase, uint32(v16+v17))) = uint8(v8)
			v21 = v16 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v21
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = v21 << (uint(int32(2)) % 32)
			m.G0 = v6 + int32(16)
			return v24
		}
	}
}
func F_checkCond(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v245 int32
	_ = v245
	F_check_stack_depth(m)
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
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_checkCond[0]))
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if l1 == int32(0) {
		v245 = l3
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return base.B2i32(v245 == int32(0))
L8:
	;
	v29 = l0
	v30 = l1
	v31 = l2
	v32 = l3
	goto L9
L9:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
	if v47&int32(32) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return int32(0)
L11:
	;
	if v59 < v32 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v52 = int32(1)
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+4)))
	if v54 != 0 {
		v58 = v52
		v59 = v52
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+6)))
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+8)))
	v58 = v56
	v59 = v57
	goto L11
L15:
	;
	goto L14
L16:
	;
	v61 = v59
	goto L18
L17:
	;
	v61 = v32
	goto L18
L18:
	;
	if v61 < v58 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	goto L21
L21:
	;
	v66 = v30 - int32(1)
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29))))
	v72 = v29 + (v67+int32(7))&int32(_a_F_checkCond_0)
	if v61 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L10
L23:
	;
	v78 = v31
	v79 = v32
	v83 = int32(0)
	goto L26
L24:
	;
	v222 = v31
	v223 = v32
	goto L25
L25:
	;
	if v30 < int32(2) {
		v245 = v223
		goto L7
	} else {
		goto L57
	}
L26:
	;
	if base.Ui32(v83) < base.Ui32(v58) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v222 = v216
	v223 = v210
	goto L25
L28:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+2)))
	v103 = v101 & int32(16)
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+4)))
	if v104 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v95 = F_checkCond(m, v72, v66, v78, v79)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v95 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	return int32(1)
L32:
	;
	v209 = int32(1)
	v210 = v79 - v209
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78))))
	v216 = v78 + (v211+int32(9))&int32(_a_F_checkCond_0)
	v218 = v83 + v209
	if v218 != v61 {
		v78 = v216
		v79 = v210
		v83 = v218
		goto L26
	} else {
		goto L56
	}
L33:
	;
	if v103 != 0 {
		goto L22
	} else {
		goto L55
	}
L34:
	;
	v114 = v29 + int32(16)
	v121 = int32(0)
	goto L35
L35:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
	if v130&int32(2) != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v103 != 0 {
		goto L32
	} else {
		goto L54
	}
L37:
	;
	v133 = int32(_a_F_checkCond_1)
	goto L39
L38:
	;
	v133 = int32(_a_F_checkCond_2)
	goto L39
L39:
	;
	if v130&int32(4) != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)))
	v168 = v121 + int32(1)
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+4)))
	if base.Ui32(v168) < base.Ui32(v169) {
		v114 = v114 + (v159+int32(7))&int32(_a_F_checkCond_0) + int32(8)
		v121 = v168
		goto L35
	} else {
		goto L53
	}
L41:
	;
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)))
	v141 = F_compare_subnode(m, v78, v114+int32(7), v138, v133, v130&int32(1))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)))
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78))))
	if v145 != v146 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v141 == int32(0) {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L33
L46:
	;
	if v130&int32(1) == int32(0) {
		goto L40
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v155 = m.T0[v133].(func(*base.Module, int32, int32, int32, int32) int32)(m, v114+int32(7), v145, v78+int32(2), v146)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	if base.Ui32(v146) <= base.Ui32(v145) {
		goto L40
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	if v155 != 0 {
		goto L33
	} else {
		goto L52
	}
L52:
	;
	goto L40
L53:
	;
	goto L36
L54:
	;
	return int32(0)
L55:
	;
	goto L32
L56:
	;
	goto L27
L57:
	;
	v29 = v72
	v30 = v66
	v31 = v222
	v32 = v223
	goto L9
}
func F_check_assignable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_check_assignable[0]))
	v13 = l0
	for {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v19 != int32(3) {
			break
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v12+v51<<(uint(int32(2))%32))))
			v13 = v55
			continue
		}
		break
	}
	switch v19 {
	case 0, 2, 4:
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+16)))
		if v22 != int32(1) {
			m.G0 = v9 + int32(32)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(_a_F_check_assignable_0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				F_errcode(m, int32(83886210))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v34
					F_errmsg(m, int32(_a_F_check_assignable_1), v9+int32(16))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						v42 = F_plpgsql_scanner_errposition(m, l1, l2)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_assignable_2), int32(3569), int32(_a_F_check_assignable_3))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	case 1:
		m.G0 = v9 + int32(32)
		return
	default:
		F_errstart_cold(m, int32(21), int32(_a_F_check_assignable_0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v62
			F_errmsg_internal(m, int32(_a_F_check_assignable_4), v9)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_check_assignable_2), int32(3580), int32(_a_F_check_assignable_3))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_check_datestyle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v477 int32
	_ = v477
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v521 int32
	_ = v521
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v544 int32
	_ = v544
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
	var v567 int32
	_ = v567
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v625 int32
	_ = v625
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
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
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 float64
	_ = v674
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v767 int32
	_ = v767
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v897 int32
	_ = v897
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v927 int64
	_ = v927
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1019 int32
	_ = v1019
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1050 int32
	_ = v1050
	v4 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[0]))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[1]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_pstrdup(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v21 + int32(16)
	return v1050
L2:
	;
	return int32(0)
L3:
	;
	v35 = F_SplitIdentifierString(m, v28, int32(44), v21+int32(12))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v35 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[3])) = v40
	goto L8
L6:
	;
	goto L7
L7:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v54 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	v46 = F_format_elog_string(m, int32(_a_F_check_datestyle_0), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[4])) = v46
	F_pfree(m, v28)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_list_free(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v1050 = v4
	goto L1
L12:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[3])) = v1019
	goto L294
L13:
	;
	v907 = F_guc_malloc(m, int32(32))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L2
	} else {
		goto L263
	}
L14:
	;
	v897 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[3])) = v897
	goto L261
L15:
	;
	F_pfree(m, v28)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L2
	} else {
		goto L258
	}
L16:
	;
	v860 = int32(1)
	v862 = v24
	v863 = v26
	goto L15
L17:
	;
	goto L18
L18:
	;
	v58 = int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v59 <= int32(0) {
		v860 = v58
		v862 = v24
		v863 = v26
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v67 = v58
	v68 = v4
	v69 = v24
	v70 = v26
	v71 = v4
	v73 = v4
	goto L20
L20:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v68<<(uint(int32(2))%32))))
	v88 = v84
	v89 = int32(_a_F_check_datestyle_1)
	goto L24
L21:
	;
	v860 = v838
	v862 = v840
	v863 = v841
	goto L15
L22:
	;
	v852 = v68 + int32(1)
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v852 < v853 {
		v67 = v838
		v68 = v852
		v69 = v840
		v70 = v841
		v71 = v842
		v73 = v844
		goto L20
	} else {
		goto L257
	}
L23:
	;
	if v126 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L24:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v92 == v93 {
		v115 = v92
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v126 = int32(0)
	goto L23
L26:
	;
	v117 = int32(1)
	if v115 != 0 {
		v88 = v88 + v117
		v89 = v89 + v117
		goto L24
	} else {
		goto L35
	}
L27:
	;
	if base.Ui32((v92-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v103 = v92 | int32(32)
	goto L30
L29:
	;
	v103 = v92
	goto L30
L30:
	;
	if base.Ui32((v93-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v112 = v93 | int32(32)
	goto L33
L32:
	;
	v112 = v93
	goto L33
L33:
	;
	if v103 == v112 {
		v115 = v103
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v126 = v103 - v112
	goto L23
L35:
	;
	goto L25
L36:
	;
	v131 = int32(1)
	v838 = (base.B2i32(v71 == int32(0)) | base.B2i32(v70 == v131)) & v67
	v840 = v69
	v841 = v131
	v842 = v131
	v844 = v73
	goto L22
L37:
	;
	goto L38
L38:
	;
	v140 = v84
	v141 = int32(_a_F_check_datestyle_2)
	goto L40
L39:
	;
	if v178 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L40:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v144 == v145 {
		v167 = v144
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v178 = int32(0)
	goto L39
L42:
	;
	v169 = int32(1)
	if v167 != 0 {
		v140 = v140 + v169
		v141 = v141 + v169
		goto L40
	} else {
		goto L51
	}
L43:
	;
	if base.Ui32((v144-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v155 = v144 | int32(32)
	goto L46
L45:
	;
	v155 = v144
	goto L46
L46:
	;
	if base.Ui32((v145-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v164 = v145 | int32(32)
	goto L49
L48:
	;
	v164 = v145
	goto L49
L49:
	;
	if v155 == v164 {
		v167 = v155
		goto L42
	} else {
		goto L50
	}
L50:
	;
	v178 = v155 - v164
	goto L39
L51:
	;
	goto L41
L52:
	;
	v183 = int32(2)
	v838 = (base.B2i32(v71 == int32(0)) | base.B2i32(v70 == v183)) & v67
	v840 = v69
	v841 = v183
	v842 = int32(1)
	v844 = v73
	goto L22
L53:
	;
	goto L54
L54:
	;
	v193 = v84
	v194 = int32(_a_F_check_datestyle_3)
	v195 = int32(8)
	goto L56
L55:
	;
	if v240 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L56:
	;
	if v195 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v240 = int32(0)
	goto L55
L58:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v198 == v199 {
		v221 = v198
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	goto L57
L61:
	;
	v223 = int32(1)
	if v221 != 0 {
		v193 = v193 + v223
		v194 = v194 + v223
		v195 = v195 - v223
		goto L56
	} else {
		goto L70
	}
L62:
	;
	if base.Ui32((v198-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v209 = v198 | int32(32)
	goto L65
L64:
	;
	v209 = v198
	goto L65
L65:
	;
	if base.Ui32((v199-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v218 = v199 | int32(32)
	goto L68
L67:
	;
	v218 = v199
	goto L68
L68:
	;
	if v209 == v218 {
		v221 = v209
		goto L61
	} else {
		goto L69
	}
L69:
	;
	v240 = v209 - v218
	goto L55
L70:
	;
	goto L60
L71:
	;
	v243 = int32(0)
	v838 = (base.B2i32(v71 == v243) | base.B2i32(v70 == v243)) & v67
	v840 = v69
	v841 = v243
	v842 = int32(1)
	v844 = v73
	goto L22
L72:
	;
	goto L73
L73:
	;
	v254 = v84
	v255 = int32(_a_F_check_datestyle_4)
	goto L75
L74:
	;
	if v292 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L75:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if v258 == v259 {
		v281 = v258
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v292 = int32(0)
	goto L74
L77:
	;
	v283 = int32(1)
	if v281 != 0 {
		v254 = v254 + v283
		v255 = v255 + v283
		goto L75
	} else {
		goto L86
	}
L78:
	;
	if base.Ui32((v258-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v269 = v258 | int32(32)
	goto L81
L80:
	;
	v269 = v258
	goto L81
L81:
	;
	if base.Ui32((v259-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v278 = v259 | int32(32)
	goto L84
L83:
	;
	v278 = v259
	goto L84
L84:
	;
	if v269 == v278 {
		v281 = v269
		goto L77
	} else {
		goto L85
	}
L85:
	;
	v292 = v269 - v278
	goto L74
L86:
	;
	goto L76
L87:
	;
	if v73 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v308 = v84
	v309 = int32(_a_F_check_datestyle_5)
	goto L94
L90:
	;
	v302 = v69
	goto L92
L91:
	;
	v302 = int32(1)
	goto L92
L92:
	;
	v838 = (base.B2i32(v71 == int32(0)) | base.B2i32(v70 == int32(3))) & v67
	v840 = v302
	v841 = int32(3)
	v842 = int32(1)
	v844 = v73
	goto L22
L93:
	;
	if v346 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L94:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if v312 == v313 {
		v335 = v312
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v346 = int32(0)
	goto L93
L96:
	;
	v337 = int32(1)
	if v335 != 0 {
		v308 = v308 + v337
		v309 = v309 + v337
		goto L94
	} else {
		goto L105
	}
L97:
	;
	if base.Ui32((v312-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v323 = v312 | int32(32)
	goto L100
L99:
	;
	v323 = v312
	goto L100
L100:
	;
	if base.Ui32((v313-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v332 = v313 | int32(32)
	goto L103
L102:
	;
	v332 = v313
	goto L103
L103:
	;
	if v323 == v332 {
		v335 = v323
		goto L96
	} else {
		goto L104
	}
L104:
	;
	v346 = v323 - v332
	goto L93
L105:
	;
	goto L95
L106:
	;
	v349 = int32(0)
	v838 = (base.B2i32(v73 == v349) | base.B2i32(v69 == v349)) & v67
	v840 = v349
	v841 = v70
	v842 = v71
	v844 = int32(1)
	goto L22
L107:
	;
	goto L108
L108:
	;
	v360 = v84
	v361 = int32(_a_F_check_datestyle_6)
	goto L111
L109:
	;
	v462 = v84
	v463 = int32(_a_F_check_datestyle_7)
	goto L146
L110:
	;
	if v398 != 0 {
		goto L123
	} else {
		goto L124
	}
L111:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	if v364 == v365 {
		v387 = v364
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v398 = int32(0)
	goto L110
L113:
	;
	v389 = int32(1)
	if v387 != 0 {
		v360 = v360 + v389
		v361 = v361 + v389
		goto L111
	} else {
		goto L122
	}
L114:
	;
	if base.Ui32((v364-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v375 = v364 | int32(32)
	goto L117
L116:
	;
	v375 = v364
	goto L117
L117:
	;
	if base.Ui32((v365-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v384 = v365 | int32(32)
	goto L120
L119:
	;
	v384 = v365
	goto L120
L120:
	;
	if v375 == v384 {
		v387 = v375
		goto L113
	} else {
		goto L121
	}
L121:
	;
	v398 = v375 - v384
	goto L110
L122:
	;
	goto L112
L123:
	;
	v403 = v84
	v404 = int32(_a_F_check_datestyle_8)
	v405 = int32(4)
	goto L127
L124:
	;
	goto L125
L125:
	;
	v453 = int32(1)
	v838 = (base.B2i32(v73 == int32(0)) | base.B2i32(v69 == v453)) & v67
	v840 = v453
	v841 = v70
	v842 = v71
	v844 = v453
	goto L22
L126:
	;
	if v450 != 0 {
		goto L109
	} else {
		goto L142
	}
L127:
	;
	if v405 != 0 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v450 = int32(0)
	goto L126
L129:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	if v408 == v409 {
		v431 = v408
		goto L132
	} else {
		goto L133
	}
L130:
	;
	goto L131
L131:
	;
	goto L128
L132:
	;
	v433 = int32(1)
	if v431 != 0 {
		v403 = v403 + v433
		v404 = v404 + v433
		v405 = v405 - v433
		goto L127
	} else {
		goto L141
	}
L133:
	;
	if base.Ui32((v408-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v419 = v408 | int32(32)
	goto L136
L135:
	;
	v419 = v408
	goto L136
L136:
	;
	if base.Ui32((v409-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v428 = v409 | int32(32)
	goto L139
L138:
	;
	v428 = v409
	goto L139
L139:
	;
	if v419 == v428 {
		v431 = v419
		goto L132
	} else {
		goto L140
	}
L140:
	;
	v450 = v419 - v428
	goto L126
L141:
	;
	goto L131
L142:
	;
	goto L125
L143:
	;
	v610 = v84
	v611 = int32(_a_F_check_datestyle_9)
	goto L191
L144:
	;
	v601 = int32(2)
	v838 = (base.B2i32(v73 == int32(0)) | base.B2i32(v69 == v601)) & v67
	v840 = v601
	v841 = v70
	v842 = v71
	v844 = int32(1)
	goto L22
L145:
	;
	if v500 == int32(0) {
		goto L144
	} else {
		goto L158
	}
L146:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462))))
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	if v466 == v467 {
		v489 = v466
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v500 = int32(0)
	goto L145
L148:
	;
	v491 = int32(1)
	if v489 != 0 {
		v462 = v462 + v491
		v463 = v463 + v491
		goto L146
	} else {
		goto L157
	}
L149:
	;
	if base.Ui32((v466-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v477 = v466 | int32(32)
	goto L152
L151:
	;
	v477 = v466
	goto L152
L152:
	;
	if base.Ui32((v467-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v486 = v467 | int32(32)
	goto L155
L154:
	;
	v486 = v467
	goto L155
L155:
	;
	if v477 == v486 {
		v489 = v477
		goto L148
	} else {
		goto L156
	}
L156:
	;
	v500 = v477 - v486
	goto L145
L157:
	;
	goto L147
L158:
	;
	v506 = v84
	v507 = int32(_a_F_check_datestyle_10)
	goto L160
L159:
	;
	if v544 == int32(0) {
		goto L144
	} else {
		goto L172
	}
L160:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507))))
	if v510 == v511 {
		v533 = v510
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v544 = int32(0)
	goto L159
L162:
	;
	v535 = int32(1)
	if v533 != 0 {
		v506 = v506 + v535
		v507 = v507 + v535
		goto L160
	} else {
		goto L171
	}
L163:
	;
	if base.Ui32((v510-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v521 = v510 | int32(32)
	goto L166
L165:
	;
	v521 = v510
	goto L166
L166:
	;
	if base.Ui32((v511-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v530 = v511 | int32(32)
	goto L169
L168:
	;
	v530 = v511
	goto L169
L169:
	;
	if v521 == v530 {
		v533 = v521
		goto L162
	} else {
		goto L170
	}
L170:
	;
	v544 = v521 - v530
	goto L159
L171:
	;
	goto L161
L172:
	;
	v551 = v84
	v552 = int32(_a_F_check_datestyle_11)
	v553 = int32(7)
	goto L174
L173:
	;
	if v598 != 0 {
		goto L143
	} else {
		goto L189
	}
L174:
	;
	if v553 != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v598 = int32(0)
	goto L173
L176:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552))))
	if v556 == v557 {
		v579 = v556
		goto L179
	} else {
		goto L180
	}
L177:
	;
	goto L178
L178:
	;
	goto L175
L179:
	;
	v581 = int32(1)
	if v579 != 0 {
		v551 = v551 + v581
		v552 = v552 + v581
		v553 = v553 - v581
		goto L174
	} else {
		goto L188
	}
L180:
	;
	if base.Ui32((v556-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v567 = v556 | int32(32)
	goto L183
L182:
	;
	v567 = v556
	goto L183
L183:
	;
	if base.Ui32((v557-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v576 = v557 | int32(32)
	goto L186
L185:
	;
	v576 = v557
	goto L186
L186:
	;
	if v567 == v576 {
		v579 = v567
		goto L179
	} else {
		goto L187
	}
L187:
	;
	v598 = v567 - v576
	goto L173
L188:
	;
	goto L178
L189:
	;
	goto L144
L190:
	;
	if v648 != 0 {
		goto L12
	} else {
		goto L203
	}
L191:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610))))
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
	if v614 == v615 {
		v637 = v614
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v648 = int32(0)
	goto L190
L193:
	;
	v639 = int32(1)
	if v637 != 0 {
		v610 = v610 + v639
		v611 = v611 + v639
		goto L191
	} else {
		goto L202
	}
L194:
	;
	if base.Ui32((v614-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v625 = v614 | int32(32)
	goto L197
L196:
	;
	v625 = v614
	goto L197
L197:
	;
	if base.Ui32((v615-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v634 = v615 | int32(32)
	goto L200
L199:
	;
	v634 = v615
	goto L200
L200:
	;
	if v625 == v634 {
		v637 = v625
		goto L193
	} else {
		goto L201
	}
L201:
	;
	v648 = v625 - v634
	goto L190
L202:
	;
	goto L192
L203:
	;
	v649 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v649
	v653 = m.G0
	v655 = v653 - int32(80)
	m.G0 = v655
	v661 = F_find_option(m, int32(_a_F_check_datestyle_12), v649, v649, int32(21))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L2
	} else {
		goto L205
	}
L204:
	;
	v802 = F_guc_strdup(m, int32(15), v767)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L2
	} else {
		goto L239
	}
L205:
	;
	v663 = F_ConfigOptionIsVisible(m, v661)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L2
	} else {
		goto L206
	}
L206:
	;
	if v663 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v661)+24))
	switch v665 {
	case 0:
		goto L211
	case 1:
		goto L215
	case 2:
		goto L214
	case 3:
		goto L213
	case 4:
		goto L212
	default:
		v767 = v649
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L2
	} else {
		goto L234
	}
L210:
	;
	m.G0 = v655 + int32(80)
	goto L204
L211:
	;
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661)+112)))
	if v753 != 0 {
		goto L231
	} else {
		goto L232
	}
L212:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v661)+116))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v661)+100))
	if v688 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L213:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v661)+112))
	if v684 != 0 {
		goto L218
	} else {
		goto L219
	}
L214:
	;
	v674 = *(*float64)(unsafe.Add(mBase, uint32(v661)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v655)+16)) = v674
	v676 = int32(_a_F_check_datestyle_13)
	v682 = F_pg_snprintf(m, v676, int32(256), int32(_a_F_check_datestyle_14), v655+int32(16))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L2
	} else {
		goto L217
	}
L215:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v661)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v655))) = v666
	v668 = int32(_a_F_check_datestyle_13)
	v672 = F_pg_snprintf(m, v668, int32(256), int32(_a_F_check_datestyle_15), v655)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L2
	} else {
		goto L216
	}
L216:
	;
	v767 = v668
	goto L210
L217:
	;
	v767 = v676
	goto L210
L218:
	;
	v686 = v684
	goto L220
L219:
	;
	v686 = int32(_a_F_check_datestyle_16)
	goto L220
L220:
	;
	v767 = v686
	goto L210
L221:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L2
	} else {
		goto L228
	}
L222:
	;
	v704 = v688
	goto L223
L223:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	if v709 == int32(0) {
		goto L221
	} else {
		goto L225
	}
L224:
	;
	goto L221
L225:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v704)+4))
	if v712 == v687 {
		v767 = v709
		goto L210
	} else {
		goto L226
	}
L226:
	;
	v715 = v704 + int32(12)
	if v715 != 0 {
		v704 = v715
		goto L223
	} else {
		goto L227
	}
L227:
	;
	goto L224
L228:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v661)))
	*(*int32)(unsafe.Add(mBase, uint32(v655)+36)) = v738
	*(*int32)(unsafe.Add(mBase, uint32(v655)+32)) = v687
	F_errmsg_internal(m, int32(_a_F_check_datestyle_17), v655+int32(32))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L2
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(_a_F_check_datestyle_18), int32(3036), int32(_a_F_check_datestyle_19))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L2
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	v754 = int32(_a_F_check_datestyle_20)
	goto L233
L232:
	;
	v754 = int32(_a_F_check_datestyle_21)
	goto L233
L233:
	;
	v767 = v754
	goto L210
L234:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L2
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v655)+64)) = int32(_a_F_check_datestyle_12)
	F_errmsg(m, int32(_a_F_check_datestyle_22), v655-int32(-64))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L2
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v655)+48)) = int32(_a_F_check_datestyle_23)
	F_errdetail(m, int32(_a_F_check_datestyle_24), v655+int32(48))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L2
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_check_datestyle_18), int32(_a_F_check_datestyle_25), int32(_a_F_check_datestyle_26))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L2
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v802
	if v802 != 0 {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v71 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L241:
	;
	v809 = F_check_datestyle(m, v21+int32(8), v21+int32(4), l2)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L2
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	F_pfree(m, v28)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L2
	} else {
		goto L247
	}
L244:
	;
	if v809 != 0 {
		goto L240
	} else {
		goto L245
	}
L245:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	F_bms_free(m, v811)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L2
	} else {
		goto L246
	}
L246:
	;
	goto L243
L247:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_list_free(m, v816)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L2
	} else {
		goto L248
	}
L248:
	;
	goto L14
L249:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v819)))
	v823 = v822
	goto L251
L250:
	;
	v823 = v70
	goto L251
L251:
	;
	if v73 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v819)+4))
	v827 = v826
	goto L254
L253:
	;
	v827 = v69
	goto L254
L254:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	F_bms_free(m, v828)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L2
	} else {
		goto L255
	}
L255:
	;
	F_bms_free(m, v819)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L2
	} else {
		goto L256
	}
L256:
	;
	v838 = v67
	v840 = v827
	v841 = v823
	v842 = v71
	v844 = v73
	goto L22
L257:
	;
	goto L21
L258:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_list_free(m, v875)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L2
	} else {
		goto L259
	}
L259:
	;
	if v860 != 0 {
		goto L13
	} else {
		goto L260
	}
L260:
	;
	goto L14
L261:
	;
	v903 = F_format_elog_string(m, int32(_a_F_check_datestyle_27), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L2
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[4])) = v903
	v1050 = v4
	goto L1
L263:
	;
	if v907 == int32(0) {
		v1050 = v4
		goto L1
	} else {
		goto L264
	}
L264:
	;
	switch v863 - int32(1) {
	case 0:
		goto L269
	case 1:
		goto L268
	case 2:
		goto L267
	default:
		goto L266
	}
L265:
	;
	if v907&int32(3) == int32(0) {
		v952 = v907
		goto L272
	} else {
		goto L273
	}
L266:
	;
	v924 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_datestyle[5])))
	*(*uint8)(unsafe.Add(mBase, uint32(v907)+8)) = uint8(v924)
	v927 = *(*int64)(unsafe.Add(mBase, _c_F_check_datestyle[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v907))) = v927
	goto L265
L267:
	;
	v918 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v907)+3)) = v918
	v921 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v907))) = v921
	goto L265
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v907))) = int32(_a_F_check_datestyle_28)
	goto L265
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v907))) = int32(_a_F_check_datestyle_29)
	goto L265
L270:
	;
	v986 = v985 + v907
	switch v862 {
	case 0:
		goto L290
	case 1:
		goto L289
	default:
		goto L288
	}
L271:
	;
	v985 = v977 - v907
	goto L270
L272:
	;
	v956 = v952
	goto L281
L273:
	;
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907))))
	if v936 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v985 = int32(0)
	goto L270
L275:
	;
	goto L276
L276:
	;
	v941 = v907
	goto L277
L277:
	;
	v945 = v941 + int32(1)
	if v945&int32(3) == int32(0) {
		v952 = v945
		goto L272
	} else {
		goto L279
	}
L278:
	;
	v977 = v945
	goto L271
L279:
	;
	v950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945))))
	if v950 != 0 {
		v941 = v945
		goto L277
	} else {
		goto L280
	}
L280:
	;
	goto L278
L281:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v956)))
	v965 = int32(-2139062144)
	if (int32(16843008)-v962|v962)&v965 == v965 {
		v956 = v956 + int32(4)
		goto L281
	} else {
		goto L283
	}
L282:
	;
	v971 = v956
	goto L284
L283:
	;
	goto L282
L284:
	;
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971))))
	if v975 != 0 {
		v971 = v971 + int32(1)
		goto L284
	} else {
		goto L286
	}
L285:
	;
	v977 = v971
	goto L271
L286:
	;
	goto L285
L287:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_bms_free(m, v1005)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L2
	} else {
		goto L291
	}
L288:
	;
	v1000 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_check_datestyle[9])))
	*(*uint16)(unsafe.Add(mBase, uint32(v986)+4)) = uint16(v1000)
	v1003 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v986))) = v1003
	goto L287
L289:
	;
	v994 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_check_datestyle[11])))
	*(*uint16)(unsafe.Add(mBase, uint32(v986)+4)) = uint16(v994)
	v997 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v986))) = v997
	goto L287
L290:
	;
	v988 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_check_datestyle[13])))
	*(*uint16)(unsafe.Add(mBase, uint32(v986)+4)) = uint16(v988)
	v991 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v986))) = v991
	goto L287
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v907
	v1010 = F_guc_malloc(m, int32(8))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L2
	} else {
		goto L292
	}
L292:
	;
	if v1010 == int32(0) {
		v1050 = v4
		goto L1
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1010))) = v863
	*(*int32)(unsafe.Add(mBase, uint32(v1010)+4)) = v862
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1010
	v1050 = int32(1)
	goto L1
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v84
	v1025 = F_format_elog_string(m, int32(_a_F_check_datestyle_30), v21)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L2
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[4])) = v1025
	F_pfree(m, v28)
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L2
	} else {
		goto L296
	}
L296:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_list_free(m, v1030)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L2
	} else {
		goto L297
	}
L297:
	;
	v1050 = v4
	goto L1
}
func F_check_restrict_nonsystem_relation_kind(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = F_pstrdup(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v191
L2:
	;
	return int32(0)
L3:
	;
	v20 = F_SplitIdentifierString(m, v13, int32(44), v10+int32(12))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v20 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_check_restrict_nonsystem_relation_kind[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_restrict_nonsystem_relation_kind[1])) = v25
	goto L8
L6:
	;
	goto L7
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v39 == int32(0) {
		v153 = v4
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v31 = F_format_elog_string(m, int32(_a_F_check_restrict_nonsystem_relation_kind_0), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_restrict_nonsystem_relation_kind[2])) = v31
	F_pfree(m, v13)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_list_free(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v191 = v4
	goto L1
L12:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_check_restrict_nonsystem_relation_kind[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_restrict_nonsystem_relation_kind[1])) = v171
	goto L53
L13:
	;
	F_pfree(m, v13)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L49
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v42 <= int32(0) {
		v153 = v4
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v46 = int32(0)
	v49 = v4
	goto L16
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v46<<(uint(int32(2))%32))))
	v61 = v57
	v62 = int32(_a_F_check_restrict_nonsystem_relation_kind_1)
	goto L19
L17:
	;
	v153 = v145
	goto L13
L18:
	;
	if v99 != 0 {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v65 == v66 {
		v88 = v65
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v99 = int32(0)
	goto L18
L21:
	;
	v90 = int32(1)
	if v88 != 0 {
		v61 = v61 + v90
		v62 = v62 + v90
		goto L19
	} else {
		goto L30
	}
L22:
	;
	if base.Ui32((v65-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v76 = v65 | int32(32)
	goto L25
L24:
	;
	v76 = v65
	goto L25
L25:
	;
	if base.Ui32((v66-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v85 = v66 | int32(32)
	goto L28
L27:
	;
	v85 = v66
	goto L28
L28:
	;
	if v76 == v85 {
		v88 = v76
		goto L21
	} else {
		goto L29
	}
L29:
	;
	v99 = v76 - v85
	goto L18
L30:
	;
	goto L20
L31:
	;
	v103 = v57
	v104 = int32(_a_F_check_restrict_nonsystem_relation_kind_2)
	goto L35
L32:
	;
	v144 = int32(1)
	goto L33
L33:
	;
	v145 = v49 | v144
	v147 = v46 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v147 < v148 {
		v46 = v147
		v49 = v145
		goto L16
	} else {
		goto L48
	}
L34:
	;
	if v141 != 0 {
		goto L12
	} else {
		goto L47
	}
L35:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v107 == v108 {
		v130 = v107
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v141 = int32(0)
	goto L34
L37:
	;
	v132 = int32(1)
	if v130 != 0 {
		v103 = v103 + v132
		v104 = v104 + v132
		goto L35
	} else {
		goto L46
	}
L38:
	;
	if base.Ui32((v107-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v118 = v107 | int32(32)
	goto L41
L40:
	;
	v118 = v107
	goto L41
L41:
	;
	if base.Ui32((v108-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v127 = v108 | int32(32)
	goto L44
L43:
	;
	v127 = v108
	goto L44
L44:
	;
	if v118 == v127 {
		v130 = v118
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v141 = v118 - v127
	goto L34
L46:
	;
	goto L36
L47:
	;
	v144 = int32(2)
	goto L33
L48:
	;
	goto L17
L49:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_list_free(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v163 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v163
	if v163 == int32(0) {
		v191 = v4
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v153
	v191 = int32(1)
	goto L1
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v57
	v177 = F_format_elog_string(m, int32(_a_F_check_restrict_nonsystem_relation_kind_3), v10)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_restrict_nonsystem_relation_kind[2])) = v177
	F_pfree(m, v13)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_list_free(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v191 = v4
	goto L1
}
func F_check_virtual_generated_security_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 != int32(1) {
			v12 = F_check_functions_in_node(m, l0, int32(465), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v12 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_check_virtual_generated_security_walker_0), int32(0))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_errdetail(m, int32(_a_F_check_virtual_generated_security_walker_1), int32(0))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									v39 = F_exprLocation(m, l0)
									mBase = m.M
									F_parser_errposition(m, l1, v39)
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_check_virtual_generated_security_walker_2), int32(3282), int32(_a_F_check_virtual_generated_security_walker_3))
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
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
					}
				} else {
					v16 = F_exprType(m, l0)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return int32(0)
					} else {
						if base.Ui32(int32(_a_F_check_virtual_generated_security_walker_4)) <= base.Ui32(v16) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_check_virtual_generated_security_walker_5), int32(0))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										F_errdetail(m, int32(_a_F_check_virtual_generated_security_walker_6), int32(0))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v62 = F_exprLocation(m, l0)
											mBase = m.M
											F_parser_errposition(m, l1, v62)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_check_virtual_generated_security_walker_2), int32(3298), int32(_a_F_check_virtual_generated_security_walker_3))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
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
							}
						} else {
							v21 = F_expression_tree_walker_impl(m, l0, int32(466), l1)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return v21
							}
						}
					}
				}
			}
		} else {
			v21 = F_expression_tree_walker_impl(m, l0, int32(466), l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return v21
			}
		}
	}
}
func F_checkclass_str(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v10&int32(1) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v156
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = int32(1)
	v25 = v13 + (int32(base.Ui32(v10)>>(uint(v14)%32))&int32(2047)+int32(base.Ui32(v10)>>(uint(int32(12))%32))+v14)&int32(_a_F_checkclass_str_0)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if l3 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if l3 != 0 {
		goto L32
	} else {
		goto L33
	}
L5:
	;
	if v26 != 0 {
		goto L21
	} else {
		goto L22
	}
L6:
	;
	if v26 == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v31 = int32(1)
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v35 = F_palloc(m, v32<<(uint(v31)%32))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v39 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v39)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v35
	v43 = v25 + int32(2)
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	if base.Ui32(v43+v44<<(uint(v39)%32)) <= base.Ui32(v43) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v82 = v35
	v91 = v35
	goto L12
L11:
	;
	v49 = v35
	v50 = v43
	v55 = v44
	goto L13
L12:
	;
	v94 = (v82 - v91) >> (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v94
	if int32(0) < v94 {
		v156 = v31
		goto L1
	} else {
		goto L19
	}
L13:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50))))
	if int32(base.Ui32(v58)>>(uint(int32(base.Ui32(v59)>>(uint(int32(14))%32)))%32))&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v82 = v71
	v91 = v81
	goto L12
L15:
	;
	v66 = v59 & int32(_a_F_checkclass_str_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v49))) = uint16(v66)
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v71 = v49 + int32(2)
	v72 = v68
	goto L17
L16:
	;
	v71 = v49
	v72 = v55
	goto L17
L17:
	;
	v74 = v50 + int32(2)
	if base.Ui32(v74) < base.Ui32(v43+v72&int32(_a_F_checkclass_str_2)<<(uint(int32(1))%32)) {
		v49 = v71
		v50 = v74
		v55 = v72
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	F_pfree(m, v91)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v100 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v100)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v100
	return v100
L21:
	;
	v107 = v25 + int32(2)
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v111 = v107 + v108<<(uint(int32(1))%32)
	if base.Ui32(v111) <= base.Ui32(v107) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	if l3 == int32(0) {
		v156 = int32(1)
		goto L1
	} else {
		goto L31
	}
L24:
	;
	return int32(0)
L25:
	;
	goto L26
L26:
	;
	v116 = v107
	goto L27
L27:
	;
	v124 = int32(1)
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116))))
	if int32(base.Ui32(v26)>>(uint(int32(base.Ui32(v125)>>(uint(int32(14))%32)))%32))&v124 != 0 {
		v156 = v124
		goto L1
	} else {
		goto L29
	}
L28:
	;
	return int32(0)
L29:
	;
	v132 = v116 + int32(2)
	if base.Ui32(v132) < base.Ui32(v111) {
		v116 = v132
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v25 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v139
	v144 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v144)
	return int32(1)
L32:
	;
	v150 = int32(2)
	goto L34
L33:
	;
	v150 = int32(1)
	goto L34
L34:
	;
	v156 = v150
	goto L1
}
func F_checkmatchall_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	v4 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = m.T0[v16].(func(*base.Module) int32)(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v223 & int32(1)
L2:
	;
	return int32(0)
L3:
	;
	if v17 != 0 {
		v223 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_checkmatchall_recurse[0]))
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v27 = F_palloc_extended(m, int32(258), int32(2))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	if v27 == int32(0) {
		v223 = v4
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v34 = F__emscripten_memset_bulkmem(m, v27, base.I32_extend8_s(int32(0)), int32(258))
	mBase = m.M
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l1
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v36 == int32(0) {
		v204 = v4
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l2+v213<<(uint(int32(2))%32)))) = v34
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
	v223 = v204
	goto L1
L13:
	;
	v42 = v4
	v46 = v36
	v49 = v4
	goto L14
L14:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)))
	if v51 != int32(_a_F_checkmatchall_recurse_0) {
		v143 = v42
		v150 = v49
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v143&v150 == int32(0) {
		v204 = v143
		goto L12
	} else {
		goto L35
	}
L16:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	if v152 != 0 {
		v42 = v143
		v46 = v152
		v49 = v150
		goto L14
	} else {
		goto L34
	}
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v54 == v55 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v57 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v57)
	v143 = v57
	v150 = v49
	goto L16
L19:
	;
	goto L20
L20:
	;
	if l1 == v54 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v143 = v42
	v150 = int32(1)
	goto L16
L22:
	;
	goto L23
L23:
	;
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	if v63 != 0 {
		v204 = v62
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2+v64<<(uint(int32(2))%32))))
	if v68 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v71 = F_checkmatchall_recurse(m, l0, v54, l2)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L28
	}
L26:
	;
	v81 = v68
	goto L27
L27:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+256)))
	v84 = v81 + int32(257)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v82 != v85 {
		v204 = v62
		goto L12
	} else {
		goto L30
	}
L28:
	;
	if v71 == int32(0) {
		v204 = v62
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l2+v76<<(uint(int32(2))%32))))
	v81 = v80
	goto L27
L30:
	;
	v90 = v62
	goto L31
L31:
	;
	v100 = v90 | int32(1)
	v101 = v34 + v100
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v81))))
	v105 = v102 | v104
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v105)
	v108 = v90 | int32(2)
	v109 = v34 + v108
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v100))))
	v113 = v110 | v112
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v113)
	v116 = v90 | int32(3)
	v117 = v34 + v116
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v108))))
	v121 = v118 | v120
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v121)
	v124 = v90 + int32(4)
	v125 = v34 + v124
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v116))))
	v129 = v126 | v128
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v129)
	if v124 != int32(256) {
		v90 = v124
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v134 = v34 + int32(257)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	v137 = v135 | v136
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v137)
	v143 = int32(1)
	v150 = v49
	goto L16
L33:
	;
	goto L32
L34:
	;
	goto L15
L35:
	;
	v160 = int32(0)
	goto L36
L36:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v34))))
	if v170 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	if base.Ui32(int32(256)) < base.Ui32(v188) {
		goto L48
	} else {
		goto L49
	}
L38:
	;
	goto L37
L39:
	;
	v188 = v160
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v160 == int32(256) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v204 = int32(1)
	goto L12
L43:
	;
	goto L44
L44:
	;
	v175 = v160 | int32(1)
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v175))))
	if v177 != 0 {
		v188 = v175
		goto L38
	} else {
		goto L45
	}
L45:
	;
	v179 = v160 | int32(2)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v179))))
	if v181 != 0 {
		v188 = v179
		goto L38
	} else {
		goto L46
	}
L46:
	;
	v183 = v160 | int32(3)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v183))))
	if v185 != 0 {
		v188 = v183
		goto L38
	} else {
		goto L47
	}
L47:
	;
	v160 = v160 + int32(4)
	goto L36
L48:
	;
	v204 = int32(1)
	goto L12
L49:
	;
	goto L50
L50:
	;
	v192 = int32(1)
	v200 = F__emscripten_memset_bulkmem(m, v188+v34+v192, base.I32_extend8_s(v192), int32(257)-v188)
	mBase = m.M
	goto L51
L51:
	;
	v204 = v192
	goto L12
}
func F_chrnamed(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v11 = l2 - l1
	if v11 == int32(4) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v155
L2:
	;
	v143 = F_range_(m, l0, v137, v137, int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L42
	} else {
		goto L43
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v137 = v14
	goto L2
L4:
	;
	goto L5
L5:
	;
	v17 = v11 >> (uint(int32(2)) % 32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v19 | int32(1024)
	v27 = int32(_a_F_chrnamed_0)
	v30 = int32(_a_F_chrnamed_1)
	goto L7
L6:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	if v132 != 0 {
		v155 = l3
		goto L1
	} else {
		goto L41
	}
L7:
	;
	if v27&int32(3) == int32(0) {
		v55 = v27
		goto L11
	} else {
		goto L12
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	return l3
L9:
	;
	if v88 == v17 {
		goto L26
	} else {
		goto L27
	}
L10:
	;
	v88 = v80 - v27
	goto L9
L11:
	;
	v59 = v55
	goto L20
L12:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v39 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v88 = int32(0)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v44 = v27
	goto L16
L16:
	;
	v48 = v44 + int32(1)
	if v48&int32(3) == int32(0) {
		v55 = v48
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v80 = v48
	goto L10
L18:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v53 != 0 {
		v44 = v48
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v68 = int32(-2139062144)
	if (int32(16843008)-v65|v65)&v68 == v68 {
		v59 = v59 + int32(4)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v74 = v59
	goto L23
L22:
	;
	goto L21
L23:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v78 != 0 {
		v74 = v74 + int32(1)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v80 = v74
	goto L10
L25:
	;
	goto L24
L26:
	;
	if v17 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L28
L28:
	;
	v126 = v30 + int32(8)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v127 != 0 {
		v27 = v127
		v30 = v126
		goto L7
	} else {
		goto L40
	}
L29:
	;
	if v122 == int32(0) {
		goto L6
	} else {
		goto L39
	}
L30:
	;
	v122 = int32(0)
	goto L29
L31:
	;
	v94 = v27
	v95 = l1
	v96 = v17
	goto L32
L32:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v99 != v100 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L30
L34:
	;
	v122 = v99 - v100
	goto L29
L35:
	;
	goto L36
L36:
	;
	if v99 == int32(0) {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v105 = int32(1)
	v110 = v96 - v105
	if v110 != 0 {
		v94 = v94 + v105
		v95 = v95 + int32(4)
		v96 = v110
		goto L32
	} else {
		goto L38
	}
L38:
	;
	goto L33
L39:
	;
	goto L28
L40:
	;
	goto L8
L41:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
	v137 = v134
	goto L2
L42:
	;
	return int32(0)
L43:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v147 == int32(0) {
		v155 = l3
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v155 = v151
	goto L1
}
