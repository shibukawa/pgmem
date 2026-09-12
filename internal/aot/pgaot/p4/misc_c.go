package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CheckElement_3(m *base.Module, l0 float32) {
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	if base.Ui32(base.I32_reinterpret_f32(l0)&int32(2147483647)) < base.Ui32(int32(2139095041)) {
		if base.F32_eq(base.F32_abs(l0), math.Float32frombits(uint32(0x7f800000))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(130))
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_errmsg(m, int32(219041), int32(0))
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_errfinish(m, int32(517975), int32(122), int32(102636))
						v41 = m.ExcPending
						if v41 != 0 {
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
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errcode(m, int32(130))
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errmsg(m, int32(219078), int32(0))
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errfinish(m, int32(517975), int32(117), int32(102636))
					v25 = m.ExcPending
					if v25 != 0 {
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
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[283])))
	if v3 != int32(1) {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[284]))
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
					F_errmsg(m, int32(347031), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						F_errdetail(m, int32(635706), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							F_errhint(m, int32(697280), int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								F_errfinish(m, int32(521598), int32(5437), int32(168235))
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
			v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[285])))
			if v12 != int32(1) {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _consts[286]))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+180))
				F_RecoveryRequiresIntParameter(m, int32(149967), v17, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, _consts[287]))
					v25 = *(*int32)(unsafe.Add(mBase, _consts[284]))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+184))
					F_RecoveryRequiresIntParameter(m, int32(170744), v23, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, _consts[282]))
						v33 = *(*int32)(unsafe.Add(mBase, _consts[284]))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+188))
						F_RecoveryRequiresIntParameter(m, int32(143836), v31, v34)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, _consts[288]))
							v41 = *(*int32)(unsafe.Add(mBase, _consts[284]))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+192))
							F_RecoveryRequiresIntParameter(m, int32(150434), v39, v42)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, _consts[289]))
								v49 = *(*int32)(unsafe.Add(mBase, _consts[284]))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+196))
								F_RecoveryRequiresIntParameter(m, int32(268286), v47, v50)
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
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[196])))
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
				F_errmsg(m, int32(273101), v5)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errfinish(m, int32(514409), int32(466), int32(273488))
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
	*(*int32)(unsafe.Add(mBase, _consts[273])) = int32(11)
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
	v57 = *(*int32)(unsafe.Add(mBase, _consts[756]))
	v59 = *(*int32)(unsafe.Add(mBase, _consts[712]))
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	v66 = int32(914)
	v68 = m.G0
	v70 = v68 - int32(144)
	m.G0 = v70
	switch int32(916) {
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
	v112 = int32(929)
	v114 = m.G0
	v116 = v114 - int32(144)
	m.G0 = v116
	switch int32(931) {
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
	*(*int32)(unsafe.Add(mBase, _consts[330])) = v66
	v80 = int32(4730)
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
	v103 = F___memcpy(m, int32(4719836), v92, int32(140))
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
	*(*int32)(unsafe.Add(mBase, _consts[331])) = v112
	v126 = int32(4730)
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
	v149 = F___memcpy(m, int32(4719976), v138, int32(140))
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
	*(*int32)(unsafe.Add(mBase, _consts[332])) = v158
	v172 = int32(4730)
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
	v195 = F___memcpy(m, int32(4721796), v184, int32(140))
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
	*(*int32)(unsafe.Add(mBase, _consts[734])) = v204
	v218 = int32(4730)
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
	v241 = F___memcpy(m, int32(4721656), v230, int32(140))
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
	v296 = int32(917)
	v298 = m.G0
	v300 = v298 - int32(144)
	m.G0 = v300
	switch int32(919) {
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
	*(*int32)(unsafe.Add(mBase, _consts[656])) = v250
	v264 = int32(4730)
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
	v287 = F___memcpy(m, int32(4721516), v276, int32(140))
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
	v342 = int32(916)
	v344 = m.G0
	v346 = v344 - int32(144)
	m.G0 = v346
	switch int32(918) {
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
	*(*int32)(unsafe.Add(mBase, _consts[657])) = v296
	v310 = int32(4730)
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
	v333 = F___memcpy(m, int32(4721096), v322, int32(140))
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
	*(*int32)(unsafe.Add(mBase, _consts[658])) = v342
	v356 = int32(4730)
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
	v379 = F___memcpy(m, int32(4721376), v368, int32(140))
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
	*(*int64)(unsafe.Add(mBase, _consts[759])) = v434
	*(*int64)(unsafe.Add(mBase, _consts[760])) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v47)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v50)
	F_before_shmem_exit(m, int32(930), int32(0))
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
	*(*int32)(unsafe.Add(mBase, _consts[660])) = v388
	v402 = int32(4730)
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
	v425 = F___memcpy(m, int32(4722076), v414, int32(140))
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
	v451 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v456 = F_AllocSetContextCreateInternal(m, v451, int32(226050), int32(0), int32(8192), int32(8388608))
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
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v456
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
	v472 = int32(4548780)
	v474 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	v475 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[171])) = v474 + v475
	*(*int32)(unsafe.Add(mBase, _consts[394])) = int32(0)
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
	*(*int32)(unsafe.Add(mBase, _consts[425])) = v467
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	v628 = int32(1)
	v629 = v29 & v628
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v629)
	v632 = v30 & v628
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v632)
	F_sigprocmask(m, int32(4461352), int32(0))
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
	v504 = *(*int32)(unsafe.Add(mBase, _consts[157]))
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
	v551 = int32(*(*uint8)(unsafe.Add(mBase, _consts[761])))
	if v551 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _consts[756]))
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
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v469
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
	v562 = *(*int32)(unsafe.Add(mBase, _consts[756]))
	F_s_lock(m, v562+int32(4), int32(518255), int32(294), int32(291229))
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
	v571 = *(*int32)(unsafe.Add(mBase, _consts[756]))
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
	*(*uint8)(unsafe.Add(mBase, _consts[761])) = uint8(v589)
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
	v607 = int32(4548780)
	v609 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	*(*int32)(unsafe.Add(mBase, _consts[171])) = v609 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v484)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v487)
	F_pg_usleep(m, int32(1000000))
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
	F_errmsg_internal(m, int32(167546), int32(0))
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
	v676 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	v678 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	*(*int32)(unsafe.Add(mBase, uint32(v676)+64)) = v678
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v629)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v632)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	v685 = *(*int32)(unsafe.Add(mBase, _consts[506]))
	*(*int32)(unsafe.Add(mBase, uint32(v685))) = int32(0)
	goto L152
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+39)) = uint8(v632)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+38)) = uint8(v629)
	F_errfinish(m, int32(518255), int32(1390), int32(353625))
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
	v701 = *(*int32)(unsafe.Add(mBase, _consts[762]))
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
	*(*uint8)(unsafe.Add(mBase, _consts[763])) = uint8(v1147)
	v1150 = *(*int32)(unsafe.Add(mBase, _consts[762]))
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
	v719 = *(*int32)(unsafe.Add(mBase, _consts[710]))
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
	v721 = *(*int32)(unsafe.Add(mBase, _consts[756]))
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
	v733 = *(*int64)(unsafe.Add(mBase, _consts[759]))
	v735 = base.I32_wrap_i64(v731 - v733)
	v737 = *(*int32)(unsafe.Add(mBase, _consts[764]))
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
	v746 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
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
	v758 = *(*int32)(unsafe.Add(mBase, _consts[756]))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v758)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v758)+4)) = int32(1)
	if v759 != 0 {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	v751 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v751)+316))
	v754 = base.B2i32(v752 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v754)
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
	v767 = *(*int32)(unsafe.Add(mBase, _consts[756]))
	F_s_lock(m, v767+int32(4), int32(518255), int32(414), int32(291229))
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
	v776 = *(*int32)(unsafe.Add(mBase, _consts[756]))
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
	v802 = int32(4533784)
	v804 = *(*int64)(unsafe.Add(mBase, _consts[765]))
	*(*int64)(unsafe.Add(mBase, _consts[765])) = v804 + int64(1)
	goto L172
L177:
	;
	goto L178
L178:
	;
	v808 = int32(4533760)
	v810 = *(*int64)(unsafe.Add(mBase, _consts[766]))
	*(*int64)(unsafe.Add(mBase, _consts[766])) = v810 + int64(1)
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
	v816 = int32(4533792)
	v818 = *(*int64)(unsafe.Add(mBase, _consts[767]))
	*(*int64)(unsafe.Add(mBase, _consts[767])) = v818 + int64(1)
	goto L172
L181:
	;
	goto L182
L182:
	;
	v822 = int32(4533768)
	v824 = *(*int64)(unsafe.Add(mBase, _consts[768]))
	*(*int64)(unsafe.Add(mBase, _consts[768])) = v824 + int64(1)
	goto L172
L183:
	;
	v876 = v779 | v738<<(uint(int32(8))%32)
	v878 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[761])) = uint8(v878)
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
	v833 = *(*int32)(unsafe.Add(mBase, _consts[769]))
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
	F_errmsg_plural(m, int32(701686), int32(701626), v735, v19+int32(16))
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
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(357372)
	F_errhint(m, int32(694147), v19)
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
	F_errfinish(m, int32(518255), int32(462), int32(291229))
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
	*(*int64)(unsafe.Add(mBase, _consts[770])) = v731
	*(*int64)(unsafe.Add(mBase, _consts[771])) = v886
	*(*int64)(unsafe.Add(mBase, _consts[772])) = int64(0)
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
	*(*int64)(unsafe.Add(mBase, _consts[770])) = v731
	*(*int64)(unsafe.Add(mBase, _consts[771])) = v906
	*(*int64)(unsafe.Add(mBase, _consts[772])) = int64(0)
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
	v936 = *(*int32)(unsafe.Add(mBase, _consts[756]))
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
	v945 = *(*int32)(unsafe.Add(mBase, _consts[756]))
	F_s_lock(m, v945+int32(4), int32(518255), int32(495), int32(291229))
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
	v954 = *(*int32)(unsafe.Add(mBase, _consts[756]))
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
	*(*uint8)(unsafe.Add(mBase, _consts[761])) = uint8(v995)
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
	*(*int64)(unsafe.Add(mBase, _consts[759])) = v731
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
	v973 = int32(4533776)
	v975 = *(*int64)(unsafe.Add(mBase, _consts[773]))
	*(*int64)(unsafe.Add(mBase, _consts[773])) = v975 + int64(1)
	goto L206
L211:
	;
	*(*int64)(unsafe.Add(mBase, _consts[759])) = v731
	v981 = int32(4533800)
	v983 = *(*int64)(unsafe.Add(mBase, _consts[774]))
	*(*int64)(unsafe.Add(mBase, _consts[774])) = v983 + int64(1)
	goto L206
L212:
	;
	goto L213
L213:
	;
	v989 = int64(*(*int32)(unsafe.Add(mBase, _consts[764])))
	*(*int64)(unsafe.Add(mBase, _consts[759])) = v731 - v989 + int64(15)
	goto L206
L214:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, _consts[762]))
	if v1004 != 0 {
		v1137 = v921
		v1138 = v922
		goto L155
	} else {
		goto L215
	}
L215:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, _consts[710]))
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
	v1041 = *(*int32)(unsafe.Add(mBase, _consts[756]))
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
	v1111 = *(*int32)(unsafe.Add(mBase, _consts[506]))
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
	v1049 = *(*int32)(unsafe.Add(mBase, _consts[764]))
	v1051 = *(*int64)(unsafe.Add(mBase, _consts[759]))
	v1053 = base.I32_wrap_i64(v1047 - v1051)
	if v1049 <= v1053 {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v1055 = v1049 - v1053
	v1057 = *(*int32)(unsafe.Add(mBase, _consts[775]))
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
	v1095 = *(*int32)(unsafe.Add(mBase, _consts[506]))
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
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
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
	v1071 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+316))
	v1074 = base.B2i32(v1072 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v1074)
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
	v1078 = *(*int32)(unsafe.Add(mBase, _consts[775]))
	v1080 = *(*int64)(unsafe.Add(mBase, _consts[760]))
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
	v1127 = *(*int32)(unsafe.Add(mBase, _consts[762]))
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
	v1151 = int32(4533768)
	v1153 = *(*int64)(unsafe.Add(mBase, _consts[768]))
	*(*int64)(unsafe.Add(mBase, _consts[768])) = v1153 + int64(1)
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
	v1203 = *(*int32)(unsafe.Add(mBase, _consts[506]))
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
	*(*int32)(unsafe.Add(mBase, _consts[762])) = int32(0)
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
	v1213 = *(*int32)(unsafe.Add(mBase, _consts[710]))
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
	v1237 = *(*int32)(unsafe.Add(mBase, _consts[506]))
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
	v1248 = *(*int32)(unsafe.Add(mBase, _consts[506]))
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
	v1258 = *(*int32)(unsafe.Add(mBase, _consts[710]))
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
func F_ClosePipeStream(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, _consts[952]))
	v8 = v6 - int32(1)
	if int32(0) <= v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[954]))
	v14 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	v40 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L12
	}
L4:
	;
	v19 = v12 + v14*int32(12)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v20 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	if int32(0) < v14 {
		v14 = v14 - int32(1)
		goto L4
	} else {
		goto L11
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v23 != l0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v25 = F_FreeDesc(m, v19)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	return v25
L11:
	;
	goto L5
L12:
	;
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_errmsg_internal(m, int32(305590), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v51 = F_pgl_pclose(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	F_errfinish(m, int32(523542), int32(3076), int32(305658))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	return v51
}
func F_CommentObject(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = *(*int64)(unsafe.Add(mBase, _consts[400]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v14
	v17 = l0 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[401]))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v21 == int32(9) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return
L2:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	if v236 == int32(0) {
		goto L1
	} else {
		goto L58
	}
L3:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_CreateComments(m, v223, v224, v225, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L10
	} else {
		goto L57
	}
L4:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v118 = int32(0)
	v120 = m.G0
	v122 = v120 - int32(128)
	m.G0 = v122
	v124 = int32(1)
	if v117 == v118 {
		v145 = v118
		v146 = v124
		goto L29
	} else {
		goto L30
	}
L5:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+48))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+119)))
	v79 = v77 - int32(99)
	if int32(1)<<(uint(v79)%32)&int32(566281) != 0 {
		goto L20
	} else {
		goto L21
	}
L6:
	;
	v59 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L15
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v27 = F_get_database_oid(m, v25, int32(1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v33 = v21
	goto L9
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_get_object_address(m, l0, v33, v34, v11+int32(44), int32(4), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L13
	}
L10:
	;
	return
L11:
	;
	if v27 == int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v33 = v31
	goto L9
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v45
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v47
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	F_check_object_ownership(m, v42, v44, v11+int32(32), v43, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v54 - int32(6) {
	case 0:
		goto L5
	default:
		goto L3
	case 3, 27, 36:
		goto L4
	}
L15:
	;
	if v59 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v25
	F_errmsg(m, int32(78129), v11)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(515509), int32(61), int32(118817))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	v87 = base.B2i32(base.Ui32(v79) <= base.Ui32(int32(19)))
	goto L22
L21:
	;
	v87 = int32(0)
	goto L22
L22:
	;
	if v87 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v96 + int32(4)
	F_errmsg(m, int32(737830), v11+int32(16))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+48))
	v107 = int32(*(*int8)(unsafe.Add(mBase, uint32(v106)+119)))
	F_errdetail_relkind_not_supported(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(515509), int32(103), int32(118817))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_ScanKeyInit(m, v122+int32(32), int32(1), int32(3), int32(184), v115)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L10
	} else {
		goto L33
	}
L30:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v127 == int32(0) {
		v145 = v118
		v146 = v124
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+18)) = uint8(v130)
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+14)) = uint8(v133)
	*(*uint16)(unsafe.Add(mBase, uint32(v122)+16)) = uint16(v130)
	v138 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v122)+12)) = uint16(v138)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+24)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = v115
	v142 = F_cstring_to_text(m, v117)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+28)) = v142
	v145 = v133
	v146 = v130
	goto L29
L33:
	;
	F_ScanKeyInit(m, v122+int32(80), int32(2), int32(3), int32(184), v116)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	v163 = F_table_open(m, int32(2396), int32(3))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L10
	} else {
		goto L36
	}
L35:
	;
	F_systable_endscan(m, v171)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L10
	} else {
		goto L46
	}
L36:
	;
	v171 = F_systable_beginscan(m, v163, int32(2397), int32(1), int32(0), int32(2), v122+int32(32))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	v173 = F_systable_getnext(m, v171)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	if v173 == int32(0) {
		v194 = v118
		goto L35
	} else {
		goto L39
	}
L39:
	;
	if v146 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_CatalogTupleDelete(m, v163, v173+int32(4))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L10
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v163)+52))
	v190 = F_heap_modify_tuple(m, v173, v183, v122+int32(20), v122+int32(16), v122+int32(12))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L10
	} else {
		goto L44
	}
L43:
	;
	v194 = v118
	goto L35
L44:
	;
	F_CatalogTupleUpdate(m, v163, v173+int32(4), v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	v194 = v190
	goto L35
L46:
	;
	v197 = int32(0)
	if base.B2i32(v145 == v197)|base.B2i32(v194 != v197) == v197 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v163)+52))
	v209 = F_heap_form_tuple(m, v204, v122+int32(20), v122+int32(16))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L10
	} else {
		goto L50
	}
L48:
	;
	v213 = v194
	goto L49
L49:
	;
	if v213 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	F_CatalogTupleInsert(m, v163, v209)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	v213 = v209
	goto L49
L52:
	;
	F_pfree(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L10
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	F_sequence_close(m, v163, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L10
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	m.G0 = v122 + int32(128)
	goto L2
L57:
	;
	goto L2
L58:
	;
	F_relation_close(m, v236, int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	goto L1
}
func F_CompareFurthestCandidates(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v11 int32
	_ = v11
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)+20))
	if base.F64_lt(v7, v8) != 0 {
		v11 = int32(-1)
	} else {
		v11 = base.F64_gt(v7, v8)
	}
	return v11
}
func F_CompareLists(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_lt(v9, v10) != 0 {
		v12 = int32(-1)
	} else {
		v12 = int32(0)
	}
	if base.F64_gt(v9, v10) != 0 {
		v14 = int32(1)
	} else {
		v14 = v12
	}
	return v14
}
func F_CompareNearestDiscardedCandidates(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l1)+20))
	if base.F64_gt(v9, v10) != 0 {
		v12 = int32(-1)
	} else {
		v12 = int32(0)
	}
	if base.F64_lt(v9, v10) != 0 {
		v14 = int32(1)
	} else {
		v14 = v12
	}
	return v14
}
func F_CompleteCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v193 int32
	_ = v193
	v9 = l8
	v14 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v63
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v66 != 0 {
		goto L28
	} else {
		goto L29
	}
L2:
	;
	v62 = l1
	v63 = v14
	goto L1
L3:
	;
	goto L4
L4:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v20 != v15 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v57 = F_AllocSetContextCreateInternal(m, v15, int32(17288), int32(0), int32(1024), int32(8388608))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L25
	} else {
		goto L26
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = l2
	v62 = l1
	v63 = l2
	goto L1
L9:
	;
	if v20 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	if v15 != 0 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v25 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v24 == int32(0) {
		goto L12
	} else {
		goto L18
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v24
	goto L14
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v24
	goto L14
L18:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v30
	goto L12
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v15
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v37
	if v37 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = int32(0)
	goto L11
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = l2
	goto L24
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l2
	goto L8
L25:
	;
	return
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v57
	v60 = F_copyObjectImpl(m, l1)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v62 = v60
	v63 = v57
	goto L1
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v15
	if int32(0) < l4 {
		goto L47
	} else {
		goto L48
	}
L29:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v67 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	F_extract_query_dependencies(m, v62, l0-int32(-64), l0+int32(68), l0+int32(85))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L25
	} else {
		goto L44
	}
L31:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	switch v71 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v75 = int32(1)
		goto L35
	default:
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v76 == int32(0) {
		goto L28
	} else {
		goto L38
	}
L34:
	;
	if v75 != 0 {
		goto L30
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v75 = int32(0)
	goto L35
L37:
	;
	goto L28
L38:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v80 != int32(6) {
		v95 = int32(1)
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v95&int32(1) == int32(0) {
		goto L28
	} else {
		goto L43
	}
L40:
	;
	goto L39
L41:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v76)+28))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v87 = v85 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v87) {
		v95 = int32(0)
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v95 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v87)) % 64)))
	goto L40
L43:
	;
	goto L30
L44:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v110
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1373])))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v113)
	v115 = F_GetSearchPathMatcher(m, v63)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v115
	goto L28
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v9)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l4
	v140 = F_ChoosePortalStrategy(m, v62)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L25
	} else {
		goto L59
	}
L47:
	;
	v124 = l4 << (uint(int32(2)) % 32)
	v125 = F_palloc(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L25
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	goto L46
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v125
	if v124 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L46
L52:
	;
	v128 = F__emscripten_memcpy_bulkmem(m, v125, l3, v124)
	mBase = m.M
	goto L54
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v180
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v14
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v193)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+93)) = uint8(v193)
	return
L56:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+28))
	v176 = F_UtilityTupleDescriptor(m, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L25
	} else {
		goto L65
	}
L57:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v150 = int32(0)
	goto L61
L58:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+76))
	v145 = F_ExecCleanTypeFromTL(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L25
	} else {
		goto L60
	}
L59:
	;
	switch v140 {
	case 0, 2:
		goto L58
	case 1:
		goto L57
	case 3:
		goto L56
	default:
		v180 = int32(0)
		goto L55
	}
L60:
	;
	v180 = v145
	goto L55
L61:
	;
	v163 = int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v150<<(uint(int32(2))%32)+v147)))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+24)))
	if v167 != v163 {
		v150 = v150 + v163
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166)+96))
	v171 = F_ExecCleanTypeFromTL(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L25
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	v180 = v171
	goto L55
L65:
	;
	v180 = v176
	goto L55
}
func F_ConditionVariableSleep(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = F_ConditionVariableTimedSleep(m, l0, int32(-1), l1)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_CopyReadAttributesCSV(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v21 <= v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L7
	} else {
		goto L102
	}
L2:
	;
	m.G0 = v19 + int32(16)
	return v411
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v24 == int32(0) {
		v411 = v2
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v52 = l0 + int32(264)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v54
	goto L12
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	F_errmsg(m, int32(287025), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(522176), int32(1837), int32(541033))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v61 <= v60 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_enlargeStringInfo(m, v52, v60)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	v66 = v60
	goto L15
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v68 = v66 + v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v76 = v69
	v77 = v2
	v78 = v67
	goto L18
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v66 = v65
	goto L15
L17:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+268)) = v390 - v403
	v411 = v392
	goto L2
L18:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v89 <= v77 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+52))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L97
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v89 << (uint(int32(1)) % 32)
	v96 = F_repalloc(m, v88, v89<<(uint(int32(3))%32))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	v99 = v88
	goto L22
L22:
	;
	v101 = v77 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v99+v101))) = v76
	if base.Ui32(v78) < base.Ui32(v68) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v96
	v99 = v96
	goto L22
L24:
	;
	goto L19
L25:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v350+v101))) = int32(0)
	v355 = v77 + int32(1)
	if v206 != 0 {
		v76 = v197
		v77 = v355
		v78 = v198
		goto L18
	} else {
		goto L96
	}
L26:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v277 != 0 {
		goto L73
	} else {
		goto L74
	}
L27:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v212 == v203 {
		goto L54
	} else {
		goto L55
	}
L28:
	;
	v107 = v78
	v109 = v76
	v114 = int32(0)
	goto L31
L29:
	;
	goto L30
L30:
	;
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v76))) = uint8(v190)
	v197 = v76 + int32(1)
	v198 = v78
	v199 = v76
	v203 = v190
	v206 = v190
	goto L27
L31:
	;
	v123 = v107 + int32(1)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	v125 = base.B2i32(v124 == v50&int32(255))
	if v125 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v184)
	v263 = v163
	v264 = v149
	v265 = v149 + int32(1)
	v268 = v163 - v78
	v271 = v184
	goto L26
L33:
	;
	v148 = v123
	v149 = v109
	goto L43
L34:
	;
	if v48 == v124 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v134 = v107
	v135 = v109
	goto L36
L36:
	;
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v137)
	v139 = v134 - v78
	v140 = int32(1)
	v141 = v135 + v140
	if v114&v140 == v137 {
		v197 = v141
		v198 = v123
		v199 = v135
		v203 = v139
		v206 = v125
		goto L27
	} else {
		goto L42
	}
L37:
	;
	if base.Ui32(v123) < base.Ui32(v68) {
		goto L33
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v124)
	v132 = v109 + int32(1)
	if base.Ui32(v123) < base.Ui32(v68) {
		v107 = v123
		v109 = v132
		goto L31
	} else {
		goto L41
	}
L40:
	;
	goto L1
L41:
	;
	v134 = v123
	v135 = v132
	goto L36
L42:
	;
	v263 = v123
	v264 = v135
	v265 = v141
	v268 = v139
	v271 = v125
	goto L26
L43:
	;
	v163 = v148 + int32(1)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v164 != v46 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if base.Ui32(v163) < base.Ui32(v68) {
		v107 = v163
		v109 = v149
		v114 = int32(1)
		goto L31
	} else {
		goto L53
	}
L45:
	;
	goto L44
L46:
	;
	if base.Ui32(v178) < base.Ui32(v68) {
		v148 = v178
		v149 = v149 + int32(1)
		goto L43
	} else {
		goto L52
	}
L47:
	;
	if v164 == v48 {
		goto L45
	} else {
		goto L51
	}
L48:
	;
	if base.Ui32(v68) <= base.Ui32(v163) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if base.B2i32(v46 != v167)&base.B2i32(v48 != v167) != 0 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v167)
	v178 = v148 + int32(2)
	goto L46
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v164)
	v178 = v163
	goto L46
L52:
	;
	goto L1
L53:
	;
	goto L32
L54:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v203 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	goto L56
L56:
	;
	v263 = v198
	v264 = v199
	v265 = v197
	v268 = v203
	v271 = v206
	goto L26
L57:
	;
	if v258 == int32(0) {
		goto L25
	} else {
		goto L71
	}
L58:
	;
	v258 = int32(0)
	goto L57
L59:
	;
	goto L60
L60:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v220 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v221 = v78
	v222 = v214
	v223 = v203
	v224 = v220
	goto L65
L62:
	;
	v246 = v214
	v250 = int32(0)
	goto L63
L63:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	v258 = v250 - v251
	goto L57
L64:
	;
	v246 = v241
	v250 = v243
	goto L63
L65:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v224 != v226 {
		v241 = v222
		v243 = v224
		goto L64
	} else {
		goto L67
	}
L66:
	;
	v241 = v235
	v243 = int32(0)
	goto L64
L67:
	;
	if v226 == int32(0) {
		v241 = v222
		v243 = v224
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v231 = v223 - int32(1)
	if v231 == int32(0) {
		v241 = v222
		v243 = v224
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v234 = int32(1)
	v235 = v222 + v234
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	if v236 != 0 {
		v221 = v221 + v234
		v222 = v235
		v223 = v231
		v224 = v236
		goto L65
	} else {
		goto L70
	}
L70:
	;
	goto L66
L71:
	;
	goto L56
L72:
	;
	v349 = v77 + int32(1)
	if v271 != 0 {
		v76 = v265
		v77 = v349
		v78 = v263
		goto L18
	} else {
		goto L95
	}
L73:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v280 = v278
	goto L75
L74:
	;
	v280 = int32(0)
	goto L75
L75:
	;
	if v280 <= v77 {
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v282 == int32(0) {
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v268 != v285 {
		goto L72
	} else {
		goto L78
	}
L78:
	;
	if v268 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v330 != 0 {
		goto L72
	} else {
		goto L93
	}
L80:
	;
	v330 = int32(0)
	goto L79
L81:
	;
	goto L82
L82:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v292 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v293 = v78
	v294 = v282
	v295 = v268
	v296 = v292
	goto L87
L84:
	;
	v318 = v282
	v322 = int32(0)
	goto L85
L85:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	v330 = v322 - v323
	goto L79
L86:
	;
	v318 = v313
	v322 = v315
	goto L85
L87:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v296 != v298 {
		v313 = v294
		v315 = v296
		goto L86
	} else {
		goto L89
	}
L88:
	;
	v313 = v307
	v315 = int32(0)
	goto L86
L89:
	;
	if v298 == int32(0) {
		v313 = v294
		v315 = v296
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v303 = v295 - int32(1)
	if v303 == int32(0) {
		v313 = v294
		v315 = v296
		goto L86
	} else {
		goto L91
	}
L91:
	;
	v306 = int32(1)
	v307 = v294 + v306
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+1)))
	if v308 != 0 {
		v293 = v293 + v306
		v294 = v307
		v295 = v303
		v296 = v308
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v277)+12))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v332+v101)))
	v336 = v334 - int32(1)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v331+v336<<(uint(int32(2))%32))))
	if v340 == int32(0) {
		goto L24
	} else {
		goto L94
	}
L94:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v345 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v343+v336))) = uint8(v345)
	goto L72
L95:
	;
	v390 = v264
	v392 = v349
	goto L17
L96:
	;
	v390 = v199
	v392 = v355
	goto L17
L97:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(529579), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v357 + v358<<(uint(int32(4))%32) + v336*int32(100) + int32(24)
	F_errdetail(m, int32(656061), v19)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L7
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(522176), int32(1990), int32(541033))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L7
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(451673), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(522176), int32(1921), int32(541033))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CreateAuxProcessResourceOwner(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v6 = F_MemoryContextAllocZero(m, v4, int32(360))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(138311)
		v11 = v6 + int32(352)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+356)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v6)+352)) = v11
		*(*int32)(unsafe.Add(mBase, _consts[182])) = v6
		*(*int32)(unsafe.Add(mBase, _consts[292])) = v6
		F_on_shmem_exit(m, int32(1834), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			return
		}
	}
}
func F_CreateInheritance(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
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
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
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
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v542 int32
	_ = v542
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
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v930 int32
	_ = v930
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v969 int32
	_ = v969
	var v974 int32
	_ = v974
	v20 = m.G0
	v22 = v20 - int32(384)
	m.G0 = v22
	v26 = F_table_open(m, int32(2611), int32(3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v22+int32(240), int32(1), int32(3), int32(184), v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = int32(1)
	v43 = F_systable_beginscan(m, v26, int32(2680), v36, int32(0), v36, v22+int32(240))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L1
	} else {
		goto L237
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L1
	} else {
		goto L233
	}
L6:
	;
	v45 = F_systable_getnext(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v45 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v50 = int32(0)
	v52 = v45
	goto L11
L9:
	;
	v94 = v36
	goto L10
L10:
	;
	F_systable_endscan(m, v43)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L19
	}
L11:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+22)))
	v68 = v66 + v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v69 == v70 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v94 = v74 + int32(1)
	goto L10
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	if v50 < v72 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v74 = v72
	goto L16
L15:
	;
	v74 = v50
	goto L16
L16:
	;
	v75 = F_systable_getnext(m, v43)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v75 != 0 {
		v50 = v74
		v52 = v75
		goto L11
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	v102 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if int32(0) < v105 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L229
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L225
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L1
	} else {
		goto L221
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L1
	} else {
		goto L217
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L213
	}
L26:
	;
	v110 = int32(1)
	v115 = v110
	v117 = v105
	v120 = v110
	goto L29
L27:
	;
	goto L28
L28:
	;
	F_sequence_close(m, v102, int32(3))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L79
	}
L29:
	;
	v136 = v104 - int32(80) + v117<<(uint(int32(4))%32) + v115*int32(100)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+91)))
	if v137 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v142 = v136 + int32(4)
	v143 = F_SearchSysCacheCopyAttName(m, v140, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	v263 = v117
	goto L33
L33:
	;
	v269 = v120 + int32(1)
	v270 = base.I32_extend16_s(v269)
	if v270 <= v263 {
		v115 = v270
		v117 = v263
		v120 = v269
		goto L29
	} else {
		goto L78
	}
L34:
	;
	if v143 == int32(0) {
		goto L25
	} else {
		goto L35
	}
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v136)+68))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+22)))
	v150 = v148 + v149
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+68))
	if v147 != v151 {
		goto L21
	} else {
		goto L36
	}
L36:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v136)+76))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150)+76))
	if v153 != v154 {
		goto L21
	} else {
		goto L37
	}
L37:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v136)+96))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v150)+96))
	if v156 != v157 {
		goto L22
	} else {
		goto L38
	}
L38:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+86)))
	if v159 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+90)))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+90)))
	if v177 != 0 {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+86)))
	if v162 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v164 = int32(*(*int16)(unsafe.Add(mBase, uint32(v136)+74)))
	v165 = F_findNotNullConstraintAttnum(m, v163, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v165 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+22)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v170)+106)))
	if v172 == int32(0) {
		goto L23
	} else {
		goto L44
	}
L44:
	;
	goto L39
L45:
	;
	if l2 != 0 {
		goto L69
	} else {
		goto L70
	}
L46:
	;
	v179 = v176 & int32(255)
	if v179 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	if v176&int32(255) != 0 {
		goto L4
	} else {
		goto L68
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v179 == v177 {
		goto L45
	} else {
		goto L56
	}
L52:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v142
	F_errmsg(m, int32(287393), v22+int32(128))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(517311), int32(17563), int32(344873))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = v142
	F_errmsg(m, int32(445121), v22+int32(160))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+90)))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+90)))
	if v217 == int32(115) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v220 = int32(568965)
	goto L62
L61:
	;
	v220 = int32(558604)
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+148)) = v220
	if v214 == int32(115) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v226 = int32(568965)
	goto L65
L64:
	;
	v226 = int32(558604)
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v226
	F_errdetail(m, int32(631006), v22+int32(144))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(517311), int32(17575), int32(344873))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	goto L45
L69:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+89)))
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+89)) = uint8(v241)
	goto L71
L70:
	;
	goto L71
L71:
	;
	v243 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+94)))
	v245 = v243 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v150)+94)) = uint16(v245)
	if base.I32_extend16_s(v245) != v245 {
		goto L24
	} else {
		goto L72
	}
L72:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+119)))
	if v250 == int32(112) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v253 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+92)) = uint8(v253)
	goto L75
L74:
	;
	goto L75
L75:
	;
	F_CatalogTupleUpdate(m, v102, v143+int32(4), v143)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_pfree(m, v143)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v263 = v261
	goto L33
L78:
	;
	goto L30
L79:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v297 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_ScanKeyInit(m, v22+int32(336), int32(9), int32(3), int32(184), v294)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v307 = int32(1)
	v312 = F_systable_beginscan(m, v297, int32(2665), v307, int32(0), v307, v22+int32(336))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v317 = F_build_attrmap_by_name(m, v314, v315, int32(1))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v319 = F_systable_getnext(m, v312)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L89
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L209
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L205
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L201
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L197
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L193
	}
L89:
	;
	if v319 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v328 = v319
	goto L93
L91:
	;
	goto L92
L92:
	;
	F_systable_endscan(m, v312)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L180
	}
L93:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v328)+16))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+22)))
	v342 = v340 + v341
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+72)))
	switch v343 - int32(99) {
	case 0, 11:
		goto L96
	default:
		goto L95
	}
L94:
	;
	goto L92
L95:
	;
	v656 = F_systable_getnext(m, v312)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L178
	}
L96:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+106)))
	if v346 != 0 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	if v343 == int32(110) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v350 = F_extractNotNullColumn(m, v328)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	v352 = int32(0)
	goto L100
L100:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v22+int32(288), int32(9), int32(3), int32(184), v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L102
	}
L101:
	;
	v352 = v350
	goto L100
L102:
	;
	v362 = int32(1)
	v367 = F_systable_beginscan(m, v297, int32(2665), v362, int32(0), v362, v22+int32(288))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L107
	}
L103:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+106)))
	if v594 == int32(1) {
		goto L87
	} else {
		goto L161
	}
L104:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v328)+16))
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+22)))
	v553 = v551 + v552
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+73)))
	v555 = v548 + v550
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+73)))
	if v554 != v556 {
		goto L88
	} else {
		goto L148
	}
L105:
	;
	v530 = F_extractNotNullColumn(m, v328)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L144
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L141
	}
L107:
	;
	v369 = F_systable_getnext(m, v367)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v369 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v372 = v342 + int32(4)
	v376 = v369
	goto L112
L110:
	;
	goto L111
L111:
	;
	F_systable_endscan(m, v367)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L135
	}
L112:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v376)+16))
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+22)))
	v394 = v392 + v393
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+72)))
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+72)))
	if v395 != v396 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L111
L114:
	;
	v470 = F_systable_getnext(m, v367)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L133
	}
L115:
	;
	switch v395 - int32(99) {
	case 0:
		goto L116
	default:
		v430 = v395
		goto L117
	case 11:
		goto L118
	}
L116:
	;
	v439 = v394 + int32(4)
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439))))
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	if v443 == int32(0) {
		v462 = v442
		v463 = v443
		goto L125
	} else {
		goto L126
	}
L117:
	;
	if v430&int32(255) != int32(99) {
		goto L103
	} else {
		goto L123
	}
L118:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	v402 = F_extractNotNullColumn(m, v376)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	v405 = int32(1)
	v406 = v402 - v405
	v410 = int32(*(*int16)(unsafe.Add(mBase, uint32(v404+v406<<(uint(v405)%32)))))
	if v352 != v410 {
		goto L114
	} else {
		goto L120
	}
L120:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400+v401<<(uint(int32(4))%32)+v352*int32(100))+11)))
	if v418 != 0 {
		goto L106
	} else {
		goto L121
	}
L121:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419+v420<<(uint(int32(4))%32)+v406*int32(100))+111)))
	if v427 != 0 {
		goto L106
	} else {
		goto L122
	}
L122:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+72)))
	v430 = v428
	goto L117
L123:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v376)+16))
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436)+22)))
	v548 = v436
	v550 = v437
	goto L104
L124:
	;
	if v463-v462 == int32(0) {
		v548 = v392
		v550 = v393
		goto L104
	} else {
		goto L132
	}
L125:
	;
	goto L124
L126:
	;
	if v442 != v443 {
		v462 = v442
		v463 = v443
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v447 = v372
	v448 = v439
	goto L128
L128:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+1)))
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+1)))
	if v452 == int32(0) {
		v462 = v451
		v463 = v452
		goto L125
	} else {
		goto L130
	}
L129:
	;
	v462 = v451
	v463 = v452
	goto L125
L130:
	;
	v455 = int32(1)
	if v451 == v452 {
		v447 = v447 + v455
		v448 = v448 + v455
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	goto L114
L133:
	;
	if v470 != 0 {
		v376 = v470
		goto L112
	} else {
		goto L134
	}
L134:
	;
	goto L113
L135:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+72)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	if v493 == int32(110) {
		goto L105
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v342 + int32(4)
	F_errmsg(m, int32(731287), v22+int32(16))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(517311), int32(17811), int32(344844))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_errmsg_internal(m, int32(157180), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(517311), int32(17723), int32(344844))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	v533 = F_get_attname(m, v294, v530, int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v535 + int32(4)
	F_errmsg(m, int32(557515), v22)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(517311), int32(17806), int32(344844))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+74)))
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+74)))
	if v558 != v559 {
		goto L88
	} else {
		goto L149
	}
L149:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v297)+52))
	v562 = F_decompile_conbin(m, v328, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v564 = F_decompile_conbin(m, v376, v561)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564))))
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562))))
	if v569 == int32(0) {
		v588 = v568
		v589 = v569
		goto L153
	} else {
		goto L154
	}
L152:
	;
	if v589-v588 != 0 {
		goto L88
	} else {
		goto L160
	}
L153:
	;
	goto L152
L154:
	;
	if v568 != v569 {
		v588 = v568
		v589 = v569
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v573 = v562
	v574 = v564
	goto L156
L156:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+1)))
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+1)))
	if v578 == int32(0) {
		v588 = v577
		v589 = v578
		goto L153
	} else {
		goto L158
	}
L157:
	;
	v588 = v577
	v589 = v578
	goto L153
L158:
	;
	v581 = int32(1)
	if v577 == v578 {
		v573 = v573 + v581
		v574 = v574 + v581
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	goto L103
L161:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+76)))
	if v597 != int32(1) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+75)))
	if v606 == int32(1) {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+75)))
	if v600 != int32(1) {
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+76)))
	if v603 == int32(0) {
		goto L86
	} else {
		goto L165
	}
L165:
	;
	goto L162
L166:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+75)))
	if v609 == int32(0) {
		goto L85
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v612 = F_heap_copytuple(m, v376)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v612)+16))
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+22)))
	v616 = v614 + v615
	v617 = int32(*(*int16)(unsafe.Add(mBase, uint32(v616)+104)))
	v619 = v617 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v616)+104)) = uint16(v619)
	if base.I32_extend16_s(v619) != v619 {
		goto L84
	} else {
		goto L171
	}
L171:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+119)))
	if v624 == int32(112) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v627 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v616)+103)) = uint8(v627)
	goto L174
L173:
	;
	goto L174
L174:
	;
	F_CatalogTupleUpdate(m, v297, v612+int32(4), v612)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_pfree(m, v612)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F_systable_endscan(m, v367)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	goto L95
L178:
	;
	if v656 != 0 {
		v328 = v656
		goto L93
	} else {
		goto L179
	}
L179:
	;
	goto L94
L180:
	;
	F_sequence_close(m, v297, int32(3))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682)+119)))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_StoreSingleInheritance(m, v684, v685, v94)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v688 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+296)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v22)+292)) = v685
	v691 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+288)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v22)+344)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v22)+340)) = v684
	*(*int32)(unsafe.Add(mBase, uint32(v22)+336)) = v691
	if v683 == int32(112) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v706 = int32(97)
	goto L185
L184:
	;
	v706 = int32(110)
	goto L185
L185:
	;
	F_recordDependencyOn(m, v22+int32(336), v22+int32(288), v706)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v710 = *(*int32)(unsafe.Add(mBase, _consts[389]))
	if v710 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v712 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2611), v684, v712, v685, v712)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	F_SetRelationHasSubclass(m, v685, int32(1))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L191
	}
L190:
	;
	goto L189
L191:
	;
	F_sequence_close(m, v26, int32(3))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	m.G0 = v22 + int32(384)
	return
L193:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v733 + int32(4)
	F_errmsg(m, int32(731219), v22+int32(80))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(517311), int32(17731), int32(344844))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v756 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v394 + v756
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v755 + v756
	F_errmsg(m, int32(752915), v22+int32(32))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(517311), int32(17740), int32(344844))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v780 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v394 + v780
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v779 + v780
	F_errmsg(m, int32(752991), v22-int32(-64))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(517311), int32(17751), int32(344844))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v804 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v394 + v804
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v803 + v804
	F_errmsg(m, int32(753063), v22+int32(48))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(517311), int32(17762), int32(344844))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	F_errmsg(m, int32(128411), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(517311), int32(17775), int32(344844))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v142
	F_errmsg(m, int32(743910), v22+int32(96))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(517311), int32(17613), int32(344873))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
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
	F_errcode(m, int32(261))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	F_errmsg(m, int32(128411), int32(0))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	F_errfinish(m, int32(517311), int32(17593), int32(344873))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L221:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v877 + int32(4)
	F_errmsg(m, int32(557515), v22+int32(176))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(517311), int32(17554), int32(344873))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L225:
	;
	F_errcode(m, int32(17432708))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v899 + int32(4)
	F_errmsg(m, int32(742797), v22+int32(192))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(517311), int32(17535), int32(344873))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+212)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = v921 + int32(4)
	F_errmsg(m, int32(742854), v22+int32(208))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	F_errfinish(m, int32(517311), int32(17529), int32(344873))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L233:
	;
	F_errcode(m, int32(117571716))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+224)) = v943 + int32(4)
	F_errmsg(m, int32(434131), v22+int32(224))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(517311), int32(17411), int32(437438))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v142
	F_errmsg(m, int32(287447), v22+int32(112))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(517311), int32(17567), int32(344873))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CreateIntoRelDestReceiver(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(52))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(7)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(540)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(541)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(542)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(543)
		return v4
	}
}
func F_calc_hist_selectivity_contained(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) float64 {
	mBase := m.M
	_ = mBase
	var v14 float64
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 float64
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 float64
	_ = v114
	var v115 int32
	_ = v115
	var v118 float64
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v131 float64
	_ = v131
	var v134 float64
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 float64
	_ = v144
	var v146 float64
	_ = v146
	var v150 float64
	_ = v150
	var v151 int32
	_ = v151
	var v152 float64
	_ = v152
	var v155 float64
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 float64
	_ = v162
	var v163 int32
	_ = v163
	var v166 float64
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 float64
	_ = v173
	var v179 float64
	_ = v179
	var v182 float64
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 float64
	_ = v192
	var v194 float64
	_ = v194
	var v195 float64
	_ = v195
	var v202 float64
	_ = v202
	var v213 float64
	_ = v213
	var v218 float64
	_ = v218
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 float64
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 float64
	_ = v276
	var v277 float64
	_ = v277
	var v278 float64
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 float64
	_ = v281
	var v282 float64
	_ = v282
	var v284 int32
	_ = v284
	var v301 float64
	_ = v301
	var v305 float64
	_ = v305
	var v307 int32
	_ = v307
	var v313 float64
	_ = v313
	var v316 float64
	_ = v316
	var v317 float64
	_ = v317
	var v325 int32
	_ = v325
	var v331 float64
	_ = v331
	var v333 float64
	_ = v333
	var v334 float64
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 float64
	_ = v345
	var v352 int32
	_ = v352
	var v355 float64
	_ = v355
	var v358 float64
	_ = v358
	var v368 float64
	_ = v368
	var v373 int32
	_ = v373
	var v374 float64
	_ = v374
	var v377 float64
	_ = v377
	var v378 float64
	_ = v378
	var v379 int32
	_ = v379
	var v380 float64
	_ = v380
	var v382 int32
	_ = v382
	var v399 float64
	_ = v399
	var v403 float64
	_ = v403
	var v404 float64
	_ = v404
	var v409 float64
	_ = v409
	var v416 int32
	_ = v416
	var v420 float64
	_ = v420
	var v422 float64
	_ = v422
	var v423 float64
	_ = v423
	var v425 float64
	_ = v425
	var v429 float64
	_ = v429
	var v433 float64
	_ = v433
	var v443 float64
	_ = v443
	var v464 float64
	_ = v464
	var v474 float64
	_ = v474
	v14 = float64(0)
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v19)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	v23 = v21 ^ v19
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v23)
	v27 = l4 - v19
	v35 = int32(-1)
	v37 = v27
	goto L1
L1:
	;
	v50 = base.I32_div_s(v35+v37+int32(1), int32(2))
	v54 = F_range_cmp_bounds(m, l0, l3+v50<<(uint(int32(3))%32), l2)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	if v60 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	return float64(0)
L4:
	;
	v59 = base.B2i32(v54 < int32(0))
	if v54 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v60 = v50
	goto L7
L6:
	;
	v60 = v35
	goto L7
L7:
	;
	if v54 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v63 = v37
	goto L10
L9:
	;
	v63 = v50 - int32(1)
	goto L10
L10:
	;
	if v60 < v63 {
		v35 = v60
		v37 = v63
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L2
L12:
	;
	return float64(0)
L13:
	;
	goto L14
L14:
	;
	v70 = l0 + int32(268)
	v73 = l4 - int32(2)
	if v60 < v73 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v75 = v60
	goto L17
L16:
	;
	v75 = v73
	goto L17
L17:
	;
	v78 = l3 + v75<<(uint(int32(3))%32)
	v81 = F_get_position(m, l0, l2, v78, v78+int32(8))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v91 = v75
	v97 = v81
	v98 = v14
	v99 = v14
	goto L19
L19:
	;
	v103 = l3 + v91<<(uint(int32(3))%32)
	v104 = F_range_cmp_bounds(m, l0, v103, l1)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L22
	}
L20:
	;
	return v474
L21:
	;
	v202 = float64(0)
	if base.F64_lt(v194, v202) != 0 {
		v464 = v202
		goto L70
	} else {
		goto L71
	}
L22:
	;
	if v104 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v108 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
	if v156 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L26:
	;
	v150 = F_get_position(m, l0, l1, v103, v103+int32(8))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L46
	}
L27:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v113 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v136 != int32(1) {
		v146 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L26
	} else {
		goto L42
	}
L30:
	;
	v114 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L32
L31:
	;
	v114 = float64(1)
	goto L32
L32:
	;
	if v113 != 0 {
		v146 = v114
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v115 == int32(0) {
		v146 = v114
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v118 = float64(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v123 = F_FunctionCall2Coll(m, v70, v120, v121, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v125)&int64(9223372036854775807)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v131 = v118
	goto L38
L37:
	;
	v131 = v125
	goto L38
L38:
	;
	if base.F64_lt(v125, float64(0)) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v134 = v118
	goto L41
L40:
	;
	v134 = v131
	goto L41
L41:
	;
	v146 = v134
	goto L26
L42:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v141 == v142 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v144 = float64(0)
	goto L45
L44:
	;
	v144 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L45
L45:
	;
	v146 = v144
	goto L26
L46:
	;
	v152 = base.F64_sub(v97, v150)
	if base.F64_lt(v152, float64(0)) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v155 = float64(0)
	goto L49
L48:
	;
	v155 = v152
	goto L49
L49:
	;
	v194 = v146
	v195 = v155
	goto L21
L50:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v161 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v184 != int32(1) {
		v194 = math.Float64frombits(uint64(0x7ff0000000000000))
		v195 = v97
		goto L21
	} else {
		goto L65
	}
L53:
	;
	v162 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L55
L54:
	;
	v162 = float64(1)
	goto L55
L55:
	;
	if v161 != 0 {
		v194 = v162
		v195 = v97
		goto L21
	} else {
		goto L56
	}
L56:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v163 == int32(0) {
		v194 = v162
		v195 = v97
		goto L21
	} else {
		goto L57
	}
L57:
	;
	v166 = float64(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v171 = F_FunctionCall2Coll(m, v70, v168, v169, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v173 = *(*float64)(unsafe.Add(mBase, uint32(v171)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v173)&int64(9223372036854775807)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v179 = v166
	goto L61
L60:
	;
	v179 = v173
	goto L61
L61:
	;
	if base.F64_lt(v173, float64(0)) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v182 = v166
	goto L64
L63:
	;
	v182 = v179
	goto L64
L64:
	;
	v194 = v182
	v195 = v97
	goto L21
L65:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+6)))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v189 == v190 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v192 = float64(0)
	goto L68
L67:
	;
	v192 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L68
L68:
	;
	v194 = v192
	v195 = v97
	goto L21
L69:
	;
	v474 = base.F64_add(v98, base.F64_div(base.F64_mul(v195, v464), base.F64_convert_i32_u(v27)))
	if int32(0) <= v104 {
		goto L145
	} else {
		goto L146
	}
L70:
	;
	goto L69
L71:
	;
	v213 = float64(1)
	v218 = base.F64_abs(v194)
	if int32(1)&base.F64_eq(v218, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v464 = v213
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v224 = l6 - int32(1)
	if v224 < int32(0) {
		v464 = v213
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v228 = v224
	v232 = int32(-1)
	goto L74
L74:
	;
	v249 = int32(2)
	v250 = base.I32_div_s(v228+v232+int32(1), v249)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l5+v250<<(uint(v249)%32))))
	v255 = *(*float64)(unsafe.Add(mBase, uint32(v254)))
	if base.F64_gt(v99, v255) != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	if v224 <= v265 {
		v464 = v213
		goto L70
	} else {
		goto L87
	}
L76:
	;
	if v265 < v263 {
		v228 = v263
		v232 = v265
		goto L74
	} else {
		goto L86
	}
L77:
	;
	v263 = v228
	v265 = v250
	goto L76
L78:
	;
	goto L79
L79:
	;
	v260 = int32(1) & base.F64_ge(v99, v255)
	if v260 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v261 = v228
	goto L82
L81:
	;
	v261 = v250 - int32(1)
	goto L82
L82:
	;
	if v260 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v262 = v250
	goto L85
L84:
	;
	v262 = v232
	goto L85
L85:
	;
	v263 = v261
	v265 = v262
	goto L76
L86:
	;
	goto L75
L87:
	;
	if v265 < int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v316 = base.F64_convert_i32_s(v224)
	v317 = base.F64_div(base.F64_add(v313, base.F64_convert_i32_u(v307)), v316)
	if base.F64_eq(v99, v194) != 0 {
		v464 = v317
		goto L70
	} else {
		goto L105
	}
L89:
	;
	v307 = int32(0)
	v313 = float64(0)
	goto L88
L90:
	;
	goto L91
L91:
	;
	v274 = l5 + v265<<(uint(int32(2))%32)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	v276 = *(*float64)(unsafe.Add(mBase, uint32(v275)))
	v277 = base.F64_abs(v276)
	v278 = math.Float64frombits(uint64(0x7ff0000000000000))
	v279 = base.F64_eq(v277, v278)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	v281 = *(*float64)(unsafe.Add(mBase, uint32(v280)))
	v282 = base.F64_abs(v281)
	v284 = base.F64_eq(v282, v278)
	if v284 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v279 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	if v279 != 0 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	if base.F64_eq(base.F64_abs(v99), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v307 = v265
		v313 = float64(0.5)
		goto L88
	} else {
		goto L95
	}
L95:
	;
	v307 = v265
	v313 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v276, v99), base.F64_sub(v276, v281)))
	goto L88
L96:
	;
	if base.F64_eq(v277, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	if v284 == int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v307 = v265
	v313 = float64(1)
	goto L88
L99:
	;
	v301 = float64(0)
	goto L101
L100:
	;
	v301 = float64(0.5)
	goto L101
L101:
	;
	if base.F64_eq(v282, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v305 = v301
	goto L104
L103:
	;
	v305 = float64(0.5)
	goto L104
L104:
	;
	v307 = v265
	v313 = v305
	goto L88
L105:
	;
	if v224 <= v307 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v429 = float64(0)
	v433 = base.F64_div(base.F64_add(v425, base.F64_convert_i32_u(v416)), v316)
	if base.F64_gt(v420, v429)|base.F64_gt(v433, v429) != 0 {
		goto L138
	} else {
		goto L139
	}
L107:
	;
	v416 = v307
	v420 = v317
	v422 = v99
	v423 = v202
	v425 = v202
	goto L106
L108:
	;
	goto L109
L109:
	;
	v325 = v307
	v331 = v317
	v333 = v202
	v334 = v99
	goto L111
L110:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l5+v325<<(uint(int32(2))%32))))
	v374 = *(*float64)(unsafe.Add(mBase, uint32(v373)))
	if base.F64_eq(v345, v374) != 0 {
		goto L121
	} else {
		goto L122
	}
L111:
	;
	v339 = int32(1)
	v340 = v325 + v339
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l5+v340<<(uint(int32(2))%32))))
	v345 = *(*float64)(unsafe.Add(mBase, uint32(v344)))
	if base.B2i32(base.F64_ge(v194, v345) == int32(0))|int32(0) != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v416 = v224
	v420 = v358
	v422 = v345
	v423 = v368
	v425 = v202
	goto L106
L113:
	;
	v352 = base.F64_lt(v345, v194)
	goto L115
L114:
	;
	v352 = v339
	goto L115
L115:
	;
	if v352 == int32(0) {
		goto L110
	} else {
		goto L116
	}
L116:
	;
	v355 = float64(0)
	v358 = base.F64_div(base.F64_convert_i32_u(v325), v316)
	if base.F64_gt(v331, v355)|base.F64_gt(v358, v355) != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v368 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v331, v358), float64(0.5)), base.F64_sub(v345, v334)), v333)
	goto L119
L118:
	;
	v368 = v333
	goto L119
L119:
	;
	if v224 != v340 {
		v325 = v340
		v331 = v358
		v333 = v368
		v334 = v345
		goto L111
	} else {
		goto L120
	}
L120:
	;
	goto L112
L121:
	;
	v409 = float64(0)
	goto L123
L122:
	;
	v377 = base.F64_abs(v345)
	v378 = math.Float64frombits(uint64(0x7ff0000000000000))
	v379 = base.F64_eq(v377, v378)
	v380 = base.F64_abs(v374)
	v382 = base.F64_eq(v380, v378)
	if v382 != 0 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	v416 = v325
	v420 = v331
	v422 = v334
	v423 = v333
	v425 = v409
	goto L106
L124:
	;
	v409 = v404
	goto L123
L125:
	;
	if v379 != 0 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	if v379 != 0 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	if base.F64_eq(base.F64_abs(v194), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v404 = float64(0.5)
		goto L124
	} else {
		goto L128
	}
L128:
	;
	v404 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v345, v194), base.F64_sub(v345, v374)))
	goto L124
L129:
	;
	if base.F64_eq(v377, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	if v382 == int32(0) {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v404 = float64(1)
	goto L124
L132:
	;
	v399 = float64(0)
	goto L134
L133:
	;
	v399 = float64(0.5)
	goto L134
L134:
	;
	if base.F64_eq(v380, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v403 = v399
	goto L137
L136:
	;
	v403 = float64(0.5)
	goto L137
L137:
	;
	v404 = v403
	goto L124
L138:
	;
	v443 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v420, v433), float64(0.5)), base.F64_sub(v194, v422)), v423)
	goto L140
L139:
	;
	v443 = v423
	goto L140
L140:
	;
	if base.F64_eq(v218, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	if base.F64_eq(base.F64_abs(v443), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v464 = float64(0.5)
		goto L70
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v464 = base.F64_div(v443, base.F64_sub(v194, v99))
	goto L70
L144:
	;
	goto L143
L145:
	;
	if int32(0) < v91 {
		v91 = v91 - int32(1)
		v97 = float64(1)
		v98 = v474
		v99 = v194
		goto L19
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	goto L20
L148:
	;
	goto L147
}
func F_calc_length_hist_frac(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 int32) float64 {
	mBase := m.M
	_ = mBase
	var v11 float64
	_ = v11
	var v22 float64
	_ = v22
	var v24 int32
	_ = v24
	var v27 float64
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v93 int32
	_ = v93
	var v110 float64
	_ = v110
	var v114 float64
	_ = v114
	var v116 int32
	_ = v116
	var v122 float64
	_ = v122
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v134 int32
	_ = v134
	var v140 float64
	_ = v140
	var v142 float64
	_ = v142
	var v143 float64
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 float64
	_ = v154
	var v161 int32
	_ = v161
	var v164 float64
	_ = v164
	var v167 float64
	_ = v167
	var v177 float64
	_ = v177
	var v182 int32
	_ = v182
	var v183 float64
	_ = v183
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v188 int32
	_ = v188
	var v189 float64
	_ = v189
	var v191 int32
	_ = v191
	var v208 float64
	_ = v208
	var v212 float64
	_ = v212
	var v213 float64
	_ = v213
	var v218 float64
	_ = v218
	var v225 int32
	_ = v225
	var v229 float64
	_ = v229
	var v231 float64
	_ = v231
	var v232 float64
	_ = v232
	var v234 float64
	_ = v234
	var v238 float64
	_ = v238
	var v242 float64
	_ = v242
	var v252 float64
	_ = v252
	var v273 float64
	_ = v273
	v11 = float64(0)
	if base.F64_lt(l3, v11) != 0 {
		v273 = v11
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v273
L2:
	;
	v22 = float64(1)
	v24 = l4 ^ int32(1)
	v27 = base.F64_abs(l3)
	if base.B2i32(v24 == int32(0))&base.F64_eq(v27, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v273 = v22
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = l1 - int32(1)
	if v33 < int32(0) {
		v273 = v22
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = v33
	v41 = int32(-1)
	goto L5
L5:
	;
	v58 = int32(2)
	v59 = base.I32_div_s(v37+v41+int32(1), v58)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0+v59<<(uint(v58)%32))))
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v63)))
	if base.F64_gt(l2, v64) != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v33 <= v74 {
		v273 = v22
		goto L1
	} else {
		goto L18
	}
L7:
	;
	if v74 < v72 {
		v37 = v72
		v41 = v74
		goto L5
	} else {
		goto L17
	}
L8:
	;
	v72 = v37
	v74 = v59
	goto L7
L9:
	;
	goto L10
L10:
	;
	v69 = l4 & base.F64_ge(l2, v64)
	if v69 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v70 = v37
	goto L13
L12:
	;
	v70 = v59 - int32(1)
	goto L13
L13:
	;
	if v69 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = v59
	goto L16
L15:
	;
	v71 = v41
	goto L16
L16:
	;
	v72 = v70
	v74 = v71
	goto L7
L17:
	;
	goto L6
L18:
	;
	if v74 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v125 = base.F64_convert_i32_s(v33)
	v126 = base.F64_div(base.F64_add(v122, base.F64_convert_i32_u(v116)), v125)
	if base.F64_eq(l2, l3) != 0 {
		v273 = v126
		goto L1
	} else {
		goto L36
	}
L20:
	;
	v116 = int32(0)
	v122 = float64(0)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v83 = l0 + v74<<(uint(int32(2))%32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v85 = *(*float64)(unsafe.Add(mBase, uint32(v84)))
	v86 = base.F64_abs(v85)
	v87 = math.Float64frombits(uint64(0x7ff0000000000000))
	v88 = base.F64_eq(v86, v87)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v89)))
	v91 = base.F64_abs(v90)
	v93 = base.F64_eq(v91, v87)
	if v93 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v88 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	if v88 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if base.F64_eq(base.F64_abs(l2), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v116 = v74
		v122 = float64(0.5)
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v116 = v74
	v122 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v85, l2), base.F64_sub(v85, v90)))
	goto L19
L27:
	;
	if base.F64_eq(v86, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v93 == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v116 = v74
	v122 = float64(1)
	goto L19
L30:
	;
	v110 = float64(0)
	goto L32
L31:
	;
	v110 = float64(0.5)
	goto L32
L32:
	;
	if base.F64_eq(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v114 = v110
	goto L35
L34:
	;
	v114 = float64(0.5)
	goto L35
L35:
	;
	v116 = v74
	v122 = v114
	goto L19
L36:
	;
	if v33 <= v116 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v238 = float64(0)
	v242 = base.F64_div(base.F64_add(v234, base.F64_convert_i32_u(v225)), v125)
	if base.F64_gt(v229, v238)|base.F64_gt(v242, v238) != 0 {
		goto L69
	} else {
		goto L70
	}
L38:
	;
	v225 = v116
	v229 = v126
	v231 = l2
	v232 = v11
	v234 = v11
	goto L37
L39:
	;
	goto L40
L40:
	;
	v134 = v116
	v140 = v126
	v142 = v11
	v143 = l2
	goto L42
L41:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0+v134<<(uint(int32(2))%32))))
	v183 = *(*float64)(unsafe.Add(mBase, uint32(v182)))
	if base.F64_eq(v154, v183) != 0 {
		goto L52
	} else {
		goto L53
	}
L42:
	;
	v148 = int32(1)
	v149 = v134 + v148
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0+v149<<(uint(int32(2))%32))))
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v153)))
	if base.B2i32(base.F64_ge(l3, v154) == int32(0))|v24 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v225 = v33
	v229 = v167
	v231 = v154
	v232 = v177
	v234 = v11
	goto L37
L44:
	;
	v161 = base.F64_lt(v154, l3)
	goto L46
L45:
	;
	v161 = v148
	goto L46
L46:
	;
	if v161 == int32(0) {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v164 = float64(0)
	v167 = base.F64_div(base.F64_convert_i32_u(v134), v125)
	if base.F64_gt(v140, v164)|base.F64_gt(v167, v164) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v177 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v140, v167), float64(0.5)), base.F64_sub(v154, v143)), v142)
	goto L50
L49:
	;
	v177 = v142
	goto L50
L50:
	;
	if v33 != v149 {
		v134 = v149
		v140 = v167
		v142 = v177
		v143 = v154
		goto L42
	} else {
		goto L51
	}
L51:
	;
	goto L43
L52:
	;
	v218 = float64(0)
	goto L54
L53:
	;
	v186 = base.F64_abs(v154)
	v187 = math.Float64frombits(uint64(0x7ff0000000000000))
	v188 = base.F64_eq(v186, v187)
	v189 = base.F64_abs(v183)
	v191 = base.F64_eq(v189, v187)
	if v191 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v225 = v134
	v229 = v140
	v231 = v143
	v232 = v142
	v234 = v218
	goto L37
L55:
	;
	v218 = v213
	goto L54
L56:
	;
	if v188 != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	if v188 != 0 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	if base.F64_eq(base.F64_abs(l3), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v213 = float64(0.5)
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v213 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v154, l3), base.F64_sub(v154, v183)))
	goto L55
L60:
	;
	if base.F64_eq(v186, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	if v191 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v213 = float64(1)
	goto L55
L63:
	;
	v208 = float64(0)
	goto L65
L64:
	;
	v208 = float64(0.5)
	goto L65
L65:
	;
	if base.F64_eq(v189, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v212 = v208
	goto L68
L67:
	;
	v212 = float64(0.5)
	goto L68
L68:
	;
	v213 = v212
	goto L55
L69:
	;
	v252 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v229, v242), float64(0.5)), base.F64_sub(l3, v231)), v232)
	goto L71
L70:
	;
	v252 = v232
	goto L71
L71:
	;
	if base.F64_eq(v27, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if base.F64_eq(base.F64_abs(v252), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v273 = float64(0.5)
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v273 = base.F64_div(v252, base.F64_sub(l3, l2))
	goto L1
L75:
	;
	goto L74
}
func F_calc_non_nestloop_required_outer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v6 = v5
	} else {
		v6 = v3
	}
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v7 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		v9 = v8
	} else {
		v9 = v3
	}
	v10 = F_bms_union(m, v6, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_calc_rank(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v29 float32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v96 int32
	_ = v96
	var v109 float32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v305 float32
	_ = v305
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v372 float32
	_ = v372
	var v374 float32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 float32
	_ = v384
	var v385 int32
	_ = v385
	var v386 float32
	_ = v386
	var v388 int32
	_ = v388
	var v392 float32
	_ = v392
	var v393 int32
	_ = v393
	var v426 float32
	_ = v426
	var v428 float32
	_ = v428
	var v429 int32
	_ = v429
	var v440 float32
	_ = v440
	var v442 int32
	_ = v442
	var v475 float32
	_ = v475
	var v480 int32
	_ = v480
	var v516 float32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v560 float32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v739 int32
	_ = v739
	var v757 float32
	_ = v757
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v795 int32
	_ = v795
	var v816 float32
	_ = v816
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v851 int32
	_ = v851
	var v867 float32
	_ = v867
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v887 int32
	_ = v887
	var v911 float32
	_ = v911
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v925 float32
	_ = v925
	var v926 int32
	_ = v926
	var v931 float32
	_ = v931
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v948 float64
	_ = v948
	var v956 float32
	_ = v956
	var v958 float32
	_ = v958
	var v961 float64
	_ = v961
	var v972 float32
	_ = v972
	var v975 int32
	_ = v975
	var v1005 float32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1040 float32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1075 float32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1113 float32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1132 int32
	_ = v1132
	var v1152 float32
	_ = v1152
	var v1159 float32
	_ = v1159
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1270 float64
	_ = v1270
	var v1271 float64
	_ = v1271
	var v1304 float32
	_ = v1304
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1411 float32
	_ = v1411
	var v1419 int32
	_ = v1419
	var v1425 float32
	_ = v1425
	var v1430 int32
	_ = v1430
	var v1437 float64
	_ = v1437
	var v1443 float32
	_ = v1443
	var v1459 int32
	_ = v1459
	var v1479 float32
	_ = v1479
	v5 = int32(0)
	v29 = float32(0)
	v33 = m.G0
	v35 = v33 - int32(16)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v37 == v5 {
		v1459 = v35
		v1479 = v29
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v1459 + int32(16)
	return v1479
L2:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v40 == int32(0) {
		v1459 = v35
		v1479 = v29
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v43 != int32(2) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if base.F32_lt(v1152, float32(0)) != 0 {
		goto L147
	} else {
		goto L148
	}
L5:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v525 = F_palloc0(m, v522<<(uint(int32(2))%32))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L9
	} else {
		goto L73
	}
L6:
	;
	v63 = m.G0
	v65 = v63 - int32(16)
	m.G0 = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = int32(1)
	v73 = F_SortAndUniqItems(m, l2, v65+int32(4))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L13
	}
L7:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	switch v46 - int32(2) {
	case 0, 2:
		goto L8
	default:
		goto L6
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v40
	v52 = F_SortAndUniqItems(m, l2, v35+int32(4))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return float32(0)
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if int32(1) < v56 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	F_pfree(m, v52)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v75 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v516 = float32(0)
	goto L16
L15:
	;
	v96 = v5
	v109 = v29
	goto L17
L16:
	;
	F_pfree(m, v73)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L9
	} else {
		goto L72
	}
L17:
	;
	v113 = int32(2)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v73+v96<<(uint(v113)%32))))
	v117 = int32(8)
	v118 = v65 + v117
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = int32(0)
	v129 = l1 + v117
	v132 = v129 + v125<<(uint(v113)%32)
	if base.Ui32(v132) <= base.Ui32(v129) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v516 = base.F32_div(v475, base.F32_convert_i32_u(v75))
	goto L16
L19:
	;
	v480 = v96 + int32(1)
	if v480 != v75 {
		v96 = v480
		v109 = v475
		goto L17
	} else {
		goto L71
	}
L20:
	;
	if v271 == int32(0) {
		v475 = v109
		goto L19
	} else {
		goto L50
	}
L21:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+2)))
	if v199 != int32(1) {
		goto L37
	} else {
		goto L38
	}
L22:
	;
	v193 = v129
	v194 = v132
	v195 = v132
	goto L21
L23:
	;
	goto L24
L24:
	;
	v140 = v129
	v142 = v132
	goto L25
L25:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v147 = int32(12)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v157 = int32(2)
	v164 = base.I32_div_s((v142-v140)>>(uint(v157)%32), v157)
	v167 = v140 + v164<<(uint(v157)%32)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v176 = int32(0)
	v177 = F_tsCompareString(m, l2+int32(8)+v146*v147+int32(base.Ui32(v150)>>(uint(v147)%32)), v150&int32(4095), v129+v156<<(uint(v157)%32)+int32(base.Ui32(v168)>>(uint(v147)%32)), int32(base.Ui32(v168)>>(uint(int32(1))%32))&int32(2047), v176)
	mBase = m.M
	if v177 == v176 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v193 = v186
	v194 = v167
	v195 = v187
	goto L21
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = int32(1)
	v193 = v140
	v194 = v167
	v195 = v167
	goto L21
L28:
	;
	goto L29
L29:
	;
	v185 = base.B2i32(int32(0) < v177)
	if int32(0) < v177 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v186 = v167 + int32(4)
	goto L32
L31:
	;
	v186 = v140
	goto L32
L32:
	;
	if int32(0) < v177 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v187 = v142
	goto L35
L34:
	;
	v187 = v167
	goto L35
L35:
	;
	if base.Ui32(v186) < base.Ui32(v187) {
		v140 = v186
		v142 = v187
		goto L25
	} else {
		goto L36
	}
L36:
	;
	goto L26
L37:
	;
	v267 = int32(0)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v267 < v268 {
		goto L47
	} else {
		goto L48
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = int32(0)
	if base.Ui32(v193) < base.Ui32(v195) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v205 = v194
	goto L41
L40:
	;
	v205 = v195
	goto L41
L41:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v129+v206<<(uint(int32(2))%32)) <= base.Ui32(v205) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v217 = v205
	v218 = v206
	goto L43
L43:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v224 = int32(12)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v240 = int32(1)
	v245 = F_tsCompareString(m, l2+int32(8)+v223*v224+int32(base.Ui32(v227)>>(uint(v224)%32)), v227&int32(4095), v129+v218<<(uint(int32(2))%32)+int32(base.Ui32(v236)>>(uint(v224)%32)), int32(base.Ui32(v236)>>(uint(v240)%32))&int32(2047), v240)
	mBase = m.M
	if v245 != 0 {
		goto L37
	} else {
		goto L45
	}
L44:
	;
	goto L37
L45:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v246 + int32(1)
	v251 = v217 + int32(4)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v251) < base.Ui32(v129+v252<<(uint(int32(2))%32)) {
		v217 = v251
		v218 = v252
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v271 = v195
	goto L49
L48:
	;
	v271 = v267
	goto L49
L49:
	;
	goto L20
L50:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if v274 <= int32(0) {
		v475 = v109
		goto L19
	} else {
		goto L51
	}
L51:
	;
	v287 = v271
	v305 = v109
	goto L52
L52:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	if v311&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v475 = v440
	goto L19
L54:
	;
	v440 = base.F32_demote_f64(base.F64_add(base.F64_div(base.F64_promote_f32(base.F32_sub(base.F32_add(v428, v426), base.F32_div(v426, base.F32_convert_i32_s(v429*v429)))), float64(1.64493406685)), base.F64_promote_f32(v305)))
	v442 = v287 + int32(4)
	if (v442-v271)>>(uint(int32(2))%32) < v274 {
		v287 = v442
		v305 = v440
		goto L52
	} else {
		goto L70
	}
L55:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v318 = int32(1)
	v330 = l1 + int32(8) + v314<<(uint(int32(2))%32) + (int32(base.Ui32(v311)>>(uint(v318)%32))&int32(2047)+int32(base.Ui32(v311)>>(uint(int32(12))%32))+v318)&int32(4194302)
	goto L57
L56:
	;
	v330 = v65 + int32(12)
	goto L57
L57:
	;
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v330))))
	if v331 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v426 = float32(-1)
	v428 = float32(0)
	v429 = int32(1)
	goto L54
L59:
	;
	goto L60
L60:
	;
	v341 = int32(0)
	v349 = v341
	v355 = v341
	v372 = float32(-1)
	v374 = float32(0)
	goto L61
L61:
	;
	v378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v330+int32(2)+v349<<(uint(int32(1))%32)))))
	v379 = int32(12)
	v384 = *(*float32)(unsafe.Add(mBase, uint32(l0+int32(base.Ui32(v378)>>(uint(v379)%32))&v379)))
	v385 = base.F32_lt(v372, v384)
	if v385 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v426 = v386
	v428 = v392
	v429 = v393 + int32(1)
	goto L54
L63:
	;
	v386 = v384
	goto L65
L64:
	;
	v386 = v372
	goto L65
L65:
	;
	v388 = v349 + int32(1)
	v392 = base.F32_add(v374, base.F32_div(v384, base.F32_convert_i32_s(v388*v388)))
	if v385 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v393 = v349
	goto L68
L67:
	;
	v393 = v355
	goto L68
L68:
	;
	if v331 != v388 {
		v349 = v388
		v355 = v393
		v372 = v386
		v374 = v392
		goto L61
	} else {
		goto L69
	}
L69:
	;
	goto L62
L70:
	;
	goto L53
L71:
	;
	goto L18
L72:
	;
	m.G0 = v65 + int32(16)
	v1125 = l1
	v1127 = l3
	v1132 = v35
	v1152 = v516
	goto L4
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = int32(1073676289)
	v532 = l0
	v533 = l1
	v534 = l2
	v535 = l3
	v540 = v35
	v551 = v56
	v552 = v5
	v553 = v52
	v554 = l1 + int32(8)
	v559 = v525
	v560 = float32(-1)
	goto L74
L74:
	;
	v564 = int32(2)
	v565 = v552 << (uint(v564) % 32)
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v553+v565)))
	v568 = int32(8)
	v569 = v540 + v568
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = int32(0)
	v580 = v533 + v568
	v583 = v580 + v576<<(uint(v564)%32)
	if base.Ui32(v583) <= base.Ui32(v580) {
		goto L79
	} else {
		goto L80
	}
L75:
	;
	F_pfree(m, v559)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L9
	} else {
		goto L145
	}
L76:
	;
	v1118 = v552 + int32(1)
	if v1118 != v551 {
		v552 = v1118
		v560 = v1113
		goto L74
	} else {
		goto L144
	}
L77:
	;
	if v722 == int32(0) {
		v1113 = v560
		goto L76
	} else {
		goto L107
	}
L78:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+2)))
	if v650 != int32(1) {
		goto L94
	} else {
		goto L95
	}
L79:
	;
	v644 = v580
	v645 = v583
	v646 = v583
	goto L78
L80:
	;
	goto L81
L81:
	;
	v591 = v580
	v593 = v583
	goto L82
L82:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	v598 = int32(12)
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	v608 = int32(2)
	v615 = base.I32_div_s((v593-v591)>>(uint(v608)%32), v608)
	v618 = v591 + v615<<(uint(v608)%32)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	v627 = int32(0)
	v628 = F_tsCompareString(m, v534+int32(8)+v597*v598+int32(base.Ui32(v601)>>(uint(v598)%32)), v601&int32(4095), v580+v607<<(uint(v608)%32)+int32(base.Ui32(v619)>>(uint(v598)%32)), int32(base.Ui32(v619)>>(uint(int32(1))%32))&int32(2047), v627)
	mBase = m.M
	if v628 == v627 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v644 = v637
	v645 = v618
	v646 = v638
	goto L78
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = int32(1)
	v644 = v591
	v645 = v618
	v646 = v618
	goto L78
L85:
	;
	goto L86
L86:
	;
	v636 = base.B2i32(int32(0) < v628)
	if int32(0) < v628 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v637 = v618 + int32(4)
	goto L89
L88:
	;
	v637 = v591
	goto L89
L89:
	;
	if int32(0) < v628 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v638 = v593
	goto L92
L91:
	;
	v638 = v618
	goto L92
L92:
	;
	if base.Ui32(v637) < base.Ui32(v638) {
		v591 = v637
		v593 = v638
		goto L82
	} else {
		goto L93
	}
L93:
	;
	goto L83
L94:
	;
	v718 = int32(0)
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	if v718 < v719 {
		goto L104
	} else {
		goto L105
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = int32(0)
	if base.Ui32(v644) < base.Ui32(v646) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v656 = v645
	goto L98
L97:
	;
	v656 = v646
	goto L98
L98:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if base.Ui32(v580+v657<<(uint(int32(2))%32)) <= base.Ui32(v656) {
		goto L94
	} else {
		goto L99
	}
L99:
	;
	v668 = v656
	v669 = v657
	goto L100
L100:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	v675 = int32(12)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v668)))
	v691 = int32(1)
	v696 = F_tsCompareString(m, v534+int32(8)+v674*v675+int32(base.Ui32(v678)>>(uint(v675)%32)), v678&int32(4095), v580+v669<<(uint(int32(2))%32)+int32(base.Ui32(v687)>>(uint(v675)%32)), int32(base.Ui32(v687)>>(uint(v691)%32))&int32(2047), v691)
	mBase = m.M
	if v696 != 0 {
		goto L94
	} else {
		goto L102
	}
L101:
	;
	goto L94
L102:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = v697 + int32(1)
	v702 = v668 + int32(4)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if base.Ui32(v702) < base.Ui32(v580+v703<<(uint(int32(2))%32)) {
		v668 = v702
		v669 = v703
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v722 = v646
	goto L106
L105:
	;
	v722 = v718
	goto L106
L106:
	;
	goto L77
L107:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v540)+8))
	if v725 <= int32(0) {
		v1113 = v560
		goto L76
	} else {
		goto L108
	}
L108:
	;
	v739 = v722
	v757 = v560
	goto L109
L109:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v739)))
	if v763&int32(1) != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v1113 = v1075
	goto L76
L111:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	v770 = int32(1)
	v782 = v554 + v766<<(uint(int32(2))%32) + (int32(base.Ui32(v763)>>(uint(v770)%32))&int32(2047)+int32(base.Ui32(v763)>>(uint(int32(12))%32))+v770)&int32(4194302)
	goto L113
L112:
	;
	v782 = v540 + int32(12)
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v565+v559))) = v782
	if v552 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v782))))
	v795 = int32(0)
	v816 = v757
	goto L117
L115:
	;
	v1075 = v757
	goto L116
L116:
	;
	v1080 = v739 + int32(4)
	if (v1080-v722)>>(uint(int32(2))%32) < v725 {
		v739 = v1080
		v757 = v1075
		goto L109
	} else {
		goto L143
	}
L117:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v559+v795<<(uint(int32(2))%32))))
	if v823 == int32(0) {
		v1040 = v816
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v1075 = v1040
	goto L116
L119:
	;
	v1045 = v795 + int32(1)
	if v1045 != v552 {
		v795 = v1045
		v816 = v1040
		goto L117
	} else {
		goto L142
	}
L120:
	;
	if v786 == int32(0) {
		v1040 = v816
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v830 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v823))))
	v832 = v540 + int32(12)
	v851 = int32(0)
	v867 = v816
	goto L122
L122:
	;
	if v830 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v1040 = v1005
	goto L119
L124:
	;
	v874 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v782+int32(2)+v851<<(uint(int32(1))%32)))))
	v876 = v874 & int32(16383)
	v877 = int32(12)
	v887 = int32(0)
	v911 = v867
	goto L127
L125:
	;
	v1005 = v867
	goto L126
L126:
	;
	v1010 = v851 + int32(1)
	if v1010 != v786 {
		v851 = v1010
		v867 = v1005
		goto L122
	} else {
		goto L141
	}
L127:
	;
	v918 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v823+int32(2)+v887<<(uint(int32(1))%32)))))
	v920 = v918 & int32(16383)
	v921 = base.B2i32(v876 != v920)
	if base.B2i32(v782 == v832)|base.B2i32(v823 == v832)|v921 == int32(0) {
		v972 = v911
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v1005 = v972
	goto L126
L129:
	;
	v975 = v887 + int32(1)
	if v975 != v830 {
		v887 = v975
		v911 = v972
		goto L127
	} else {
		goto L140
	}
L130:
	;
	v925 = *(*float32)(unsafe.Add(mBase, uint32(v532+int32(base.Ui32(v874)>>(uint(v877)%32))&v877)))
	v926 = int32(12)
	v931 = *(*float32)(unsafe.Add(mBase, uint32(v532+int32(base.Ui32(v918)>>(uint(v926)%32))&v926)))
	v933 = v876 - v920
	v935 = v933 >> (uint(int32(31)) % 32)
	if v876 != v920 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v939 = v933 ^ v935 - v935
	goto L133
L132:
	;
	v939 = int32(16384)
	goto L133
L133:
	;
	if base.Ui32(v939) <= base.Ui32(int32(100)) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v948 = F_exp(m, base.F64_add(base.F64_div(base.F64_convert_i32_u(v939), float64(1.5)), float64(-2)))
	mBase = m.M
	v956 = base.F32_demote_f64(base.F64_div(float64(1), base.F64_add(base.F64_mul(v948, float64(0.05)), float64(1.005))))
	goto L136
L135:
	;
	v956 = float32(1e-30)
	goto L136
L136:
	;
	v958 = base.F32_sqrt(base.F32_mul(base.F32_mul(v925, v931), v956))
	if base.F32_lt(v911, float32(0)) != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v972 = v958
	goto L129
L138:
	;
	goto L139
L139:
	;
	v961 = float64(1)
	v972 = base.F32_demote_f64(base.F64_sub(v961, base.F64_mul(base.F64_sub(v961, base.F64_promote_f32(v911)), base.F64_sub(v961, base.F64_promote_f32(v958)))))
	goto L129
L140:
	;
	goto L128
L141:
	;
	goto L123
L142:
	;
	goto L118
L143:
	;
	goto L110
L144:
	;
	goto L75
L145:
	;
	F_pfree(m, v553)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L9
	} else {
		goto L146
	}
L146:
	;
	v1125 = v533
	v1127 = v535
	v1132 = v540
	v1152 = v1113
	goto L4
L147:
	;
	v1159 = float32(1e-20)
	goto L149
L148:
	;
	v1159 = v1152
	goto L149
L149:
	;
	if v1127&int32(1) == int32(0) {
		v1304 = v1159
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if v1127&int32(2) == int32(0) {
		v1411 = v1304
		goto L165
	} else {
		goto L166
	}
L151:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+4))
	if v1164 <= int32(0) {
		v1304 = v1159
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v1169 = v1125 + int32(8)
	v1172 = v1169 + v1164<<(uint(int32(2))%32)
	if base.Ui32(v1169) < base.Ui32(v1172) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v1179 = v1169
	v1180 = int32(0)
	goto L156
L154:
	;
	v1270 = float64(1)
	goto L155
L155:
	;
	v1271 = F_log(m, v1270)
	mBase = m.M
	v1304 = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(v1159), base.F64_div(v1271, float64(0.6931471805599453))))
	goto L150
L156:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1179)))
	if v1207&int32(1) != 0 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v1270 = base.F64_convert_i32_s(v1230 + int32(1))
	goto L155
L158:
	;
	v1210 = int32(1)
	v1223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1172+(int32(base.Ui32(v1207)>>(uint(v1210)%32))&int32(2047)+int32(base.Ui32(v1207)>>(uint(int32(12))%32))+v1210)&int32(4194302)))))
	if base.Ui32(v1223) <= base.Ui32(v1210) {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	v1229 = int32(1)
	goto L160
L160:
	;
	v1230 = v1229 + v1180
	v1232 = v1179 + int32(4)
	if base.Ui32(v1232) < base.Ui32(v1172) {
		v1179 = v1232
		v1180 = v1230
		goto L156
	} else {
		goto L164
	}
L161:
	;
	v1226 = v1210
	goto L163
L162:
	;
	v1226 = v1223
	goto L163
L163:
	;
	v1229 = v1226
	goto L160
L164:
	;
	goto L157
L165:
	;
	if v1127&int32(8) == int32(0) {
		v1425 = v1411
		goto L178
	} else {
		goto L179
	}
L166:
	;
	v1313 = v1125 + int32(8)
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+4))
	v1317 = v1313 + v1314<<(uint(int32(2))%32)
	if base.Ui32(v1317) <= base.Ui32(v1313) {
		v1411 = v1304
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v1324 = v1313
	v1325 = int32(0)
	goto L168
L168:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1324)))
	if v1352&int32(1) != 0 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v1375 <= int32(0) {
		v1411 = v1304
		goto L165
	} else {
		goto L177
	}
L170:
	;
	v1355 = int32(1)
	v1368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1317+(int32(base.Ui32(v1352)>>(uint(v1355)%32))&int32(2047)+int32(base.Ui32(v1352)>>(uint(int32(12))%32))+v1355)&int32(4194302)))))
	if base.Ui32(v1368) <= base.Ui32(v1355) {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	v1374 = int32(1)
	goto L172
L172:
	;
	v1375 = v1374 + v1325
	v1377 = v1324 + int32(4)
	if base.Ui32(v1377) < base.Ui32(v1317) {
		v1324 = v1377
		v1325 = v1375
		goto L168
	} else {
		goto L176
	}
L173:
	;
	v1371 = v1355
	goto L175
L174:
	;
	v1371 = v1368
	goto L175
L175:
	;
	v1374 = v1371
	goto L172
L176:
	;
	goto L169
L177:
	;
	v1411 = base.F32_div(v1304, base.F32_convert_i32_u(v1375))
	goto L165
L178:
	;
	if v1127&int32(16) == int32(0) {
		v1443 = v1425
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+4))
	if v1419 <= int32(0) {
		v1425 = v1411
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v1425 = base.F32_div(v1411, base.F32_convert_i32_u(v1419))
	goto L178
L181:
	;
	if v1127&int32(32) == int32(0) {
		v1459 = v1132
		v1479 = v1443
		goto L1
	} else {
		goto L184
	}
L182:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+4))
	if v1430 <= int32(0) {
		v1443 = v1425
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v1437 = F_log(m, base.F64_convert_i32_s(v1430+int32(1)))
	mBase = m.M
	v1443 = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(v1425), base.F64_div(v1437, float64(0.6931471805599453))))
	goto L181
L184:
	;
	v1459 = v1132
	v1479 = base.F32_div(v1443, base.F32_add(v1443, float32(1)))
	goto L1
}
func F_call_real_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v14 == int32(0) {
		v81 = v13
		m.G0 = v11 - int32(-64)
		return v81
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[492])) = int32(50856066)
		v21 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[493])) = v21
		*(*int32)(unsafe.Add(mBase, _consts[142])) = v21
		*(*int32)(unsafe.Add(mBase, _consts[1438])) = v21
		v29 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			if v29 != 0 {
				v81 = v13
				m.G0 = v11 - int32(-64)
				return v81
			} else {
				v34 = F_errstart(m, l4, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					if v34 != 0 {
						v37 = *(*int32)(unsafe.Add(mBase, _consts[492]))
						F_errcode(m, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, _consts[493]))
							if v41 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v41
								F_errmsg_internal(m, int32(216470), v9+int32(-16))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _consts[142]))
									if v59 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v59
										F_errdetail_internal(m, int32(216470), v9+int32(-48))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, _consts[1438]))
											if v67 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
												F_errhint(m, int32(216470), v11)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(523571), int32(6904), int32(329944))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_FlushErrorState(m)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															v81 = int32(0)
															m.G0 = v11 - int32(-64)
															return v81
														}
													}
												}
											} else {
												F_errfinish(m, int32(523571), int32(6904), int32(329944))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, _consts[1438]))
										if v67 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
											F_errhint(m, int32(216470), v11)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(523571), int32(6904), int32(329944))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										} else {
											F_errfinish(m, int32(523571), int32(6904), int32(329944))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_FlushErrorState(m)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													v81 = int32(0)
													m.G0 = v11 - int32(-64)
													return v81
												}
											}
										}
									}
								}
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+40)) = v49
								*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v48
								F_errmsg(m, int32(354349), v9+int32(-32))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _consts[142]))
									if v59 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v59
										F_errdetail_internal(m, int32(216470), v9+int32(-48))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, _consts[1438]))
											if v67 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
												F_errhint(m, int32(216470), v11)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(523571), int32(6904), int32(329944))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_FlushErrorState(m)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															v81 = int32(0)
															m.G0 = v11 - int32(-64)
															return v81
														}
													}
												}
											} else {
												F_errfinish(m, int32(523571), int32(6904), int32(329944))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, _consts[1438]))
										if v67 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
											F_errhint(m, int32(216470), v11)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(523571), int32(6904), int32(329944))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										} else {
											F_errfinish(m, int32(523571), int32(6904), int32(329944))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_FlushErrorState(m)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													v81 = int32(0)
													m.G0 = v11 - int32(-64)
													return v81
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_FlushErrorState(m)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							v81 = int32(0)
							m.G0 = v11 - int32(-64)
							return v81
						}
					}
				}
			}
		}
	}
}
func F_cfunc_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
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
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
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
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v615 int32
	_ = v615
	v4 = int32(24)
	v10 = int32(-1636608408)
	if l0&int32(3) != 0 {
		v119 = l0
		v120 = v4
		v121 = v10
		v122 = v10
		v123 = v10
		for {
			v125 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
			v126 = v125 + v122
			v127 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
			v129 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
			v130 = v129 + v123
			v132 = int32(4)
			v134 = v127 + v121 - v130 ^ base.I32_rotl(v130, v132)
			v138 = v126 - v134 ^ base.I32_rotl(v134, int32(6))
			v139 = v130 + v126
			v140 = v134 + v139
			v141 = v138 + v140
			v145 = v139 - v138 ^ base.I32_rotl(v138, int32(8))
			v149 = v140 - v145 ^ base.I32_rotl(v145, int32(16))
			v153 = v141 - v149 ^ base.I32_rotl(v149, int32(19))
			v154 = v145 + v141
			v155 = v149 + v154
			v156 = v153 + v155
			v160 = v154 - v153 ^ base.I32_rotl(v153, v132)
			v161 = int32(12)
			v162 = v119 + v161
			v164 = v120 - v161
			if base.Ui32(int32(11)) < base.Ui32(v164) {
				v119 = v162
				v120 = v164
				v121 = v155
				v122 = v156
				v123 = v160
				continue
			} else {
				break
			}
			break
		}
		switch v164 - int32(1) {
		case 0:
			v230 = v155
			v231 = v156
			v232 = v160
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 1:
			v223 = v155
			v224 = v156
			v225 = v160
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 2:
			v216 = v155
			v217 = v156
			v218 = v160
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 3:
			v210 = v156
			v211 = v160
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 4:
			v206 = v156
			v207 = v160
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 5:
			v200 = v156
			v201 = v160
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 6:
			v194 = v156
			v195 = v160
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 7:
			v189 = v160
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+7)))
			v194 = v190<<(uint(int32(24))%32) + v156
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 8:
			v184 = v160
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+7)))
			v194 = v190<<(uint(int32(24))%32) + v156
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 9:
			v179 = v160
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+7)))
			v194 = v190<<(uint(int32(24))%32) + v156
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 10:
			v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+10)))
			v179 = v175<<(uint(int32(24))%32) + v160
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+7)))
			v194 = v190<<(uint(int32(24))%32) + v156
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
			v216 = v212<<(uint(int32(24))%32) + v155
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		default:
			v237 = v155
			v238 = v156
			v239 = v160
		}
	} else {
		v17 = l0
		v18 = v4
		v19 = v10
		v20 = v10
		v21 = v10
		for {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v24 = v23 + v20
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
			v28 = v27 + v21
			v30 = int32(4)
			v32 = v25 + v19 - v28 ^ base.I32_rotl(v28, v30)
			v36 = v24 - v32 ^ base.I32_rotl(v32, int32(6))
			v37 = v28 + v24
			v38 = v32 + v37
			v39 = v36 + v38
			v43 = v37 - v36 ^ base.I32_rotl(v36, int32(8))
			v47 = v38 - v43 ^ base.I32_rotl(v43, int32(16))
			v51 = v39 - v47 ^ base.I32_rotl(v47, int32(19))
			v52 = v43 + v39
			v53 = v47 + v52
			v54 = v51 + v53
			v58 = v52 - v51 ^ base.I32_rotl(v51, v30)
			v59 = int32(12)
			v60 = v17 + v59
			v62 = v18 - v59
			if base.Ui32(int32(11)) < base.Ui32(v62) {
				v17 = v60
				v18 = v62
				v19 = v53
				v20 = v54
				v21 = v58
				continue
			} else {
				break
			}
			break
		}
		switch v62 - int32(1) {
		case 0:
			v116 = v53
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
			v237 = v116 + v117
			v238 = v54
			v239 = v58
		case 1:
			v111 = v53
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
			v116 = v112<<(uint(int32(8))%32) + v111
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
			v237 = v116 + v117
			v238 = v54
			v239 = v58
		case 2:
			v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
			v111 = v107<<(uint(int32(16))%32) + v53
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
			v116 = v112<<(uint(int32(8))%32) + v111
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
			v237 = v116 + v117
			v238 = v54
			v239 = v58
		case 3:
			v104 = v54
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v237 = v105 + v53
			v238 = v104
			v239 = v58
		case 4:
			v101 = v54
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)))
			v104 = v101 + v102
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v237 = v105 + v53
			v238 = v104
			v239 = v58
		case 5:
			v96 = v54
			v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+5)))
			v101 = v97<<(uint(int32(8))%32) + v96
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)))
			v104 = v101 + v102
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v237 = v105 + v53
			v238 = v104
			v239 = v58
		case 6:
			v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+6)))
			v96 = v92<<(uint(int32(16))%32) + v54
			v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+5)))
			v101 = v97<<(uint(int32(8))%32) + v96
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)))
			v104 = v101 + v102
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v237 = v105 + v53
			v238 = v104
			v239 = v58
		case 7:
			v87 = v58
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
			v237 = v88 + v53
			v238 = v90 + v54
			v239 = v87
		case 8:
			v82 = v58
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+8)))
			v87 = v83<<(uint(int32(8))%32) + v82
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
			v237 = v88 + v53
			v238 = v90 + v54
			v239 = v87
		case 9:
			v77 = v58
			v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+9)))
			v82 = v78<<(uint(int32(16))%32) + v77
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+8)))
			v87 = v83<<(uint(int32(8))%32) + v82
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
			v237 = v88 + v53
			v238 = v90 + v54
			v239 = v87
		case 10:
			v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+10)))
			v77 = v73<<(uint(int32(24))%32) + v58
			v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+9)))
			v82 = v78<<(uint(int32(16))%32) + v77
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+8)))
			v87 = v83<<(uint(int32(8))%32) + v82
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
			v237 = v88 + v53
			v238 = v90 + v54
			v239 = v87
		default:
			v237 = v53
			v238 = v54
			v239 = v58
		}
	}
	v242 = int32(14)
	v244 = v238 ^ v239 - base.I32_rotl(v238, v242)
	v248 = v244 ^ v237 - base.I32_rotl(v244, int32(11))
	v252 = v248 ^ v238 - base.I32_rotl(v248, int32(25))
	v256 = v252 ^ v244 - base.I32_rotl(v252, int32(16))
	v260 = v256 ^ v248 - base.I32_rotl(v256, int32(4))
	v264 = v260 ^ v252 - base.I32_rotl(v260, v242)
	v268 = v264 ^ v256 - base.I32_rotl(v264, int32(24))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v269 {
		v273 = l0 + int32(28)
		v275 = v269 << (uint(int32(2)) % 32)
		v281 = v275 - int32(1636608432)
		if v273&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v275) {
				v390 = v273
				v391 = v275
				v392 = v281
				v393 = v281
				v394 = v281
				for {
					v396 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
					v397 = v396 + v393
					v398 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
					v400 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
					v401 = v400 + v394
					v403 = int32(4)
					v405 = v398 + v392 - v401 ^ base.I32_rotl(v401, v403)
					v409 = v397 - v405 ^ base.I32_rotl(v405, int32(6))
					v410 = v401 + v397
					v411 = v405 + v410
					v412 = v409 + v411
					v416 = v410 - v409 ^ base.I32_rotl(v409, int32(8))
					v420 = v411 - v416 ^ base.I32_rotl(v416, int32(16))
					v424 = v412 - v420 ^ base.I32_rotl(v420, int32(19))
					v425 = v416 + v412
					v426 = v420 + v425
					v427 = v424 + v426
					v431 = v425 - v424 ^ base.I32_rotl(v424, v403)
					v432 = int32(12)
					v433 = v390 + v432
					v435 = v391 - v432
					if base.Ui32(int32(11)) < base.Ui32(v435) {
						v390 = v433
						v391 = v435
						v392 = v426
						v393 = v427
						v394 = v431
						continue
					} else {
						break
					}
					break
				}
				v438 = v433
				v439 = v435
				v440 = v426
				v441 = v427
				v442 = v431
			} else {
				v438 = v273
				v439 = v275
				v440 = v281
				v441 = v281
				v442 = v281
			}
			switch v439 - int32(1) {
			case 0:
				v501 = v440
				v502 = v441
				v503 = v442
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 1:
				v494 = v440
				v495 = v441
				v496 = v442
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 2:
				v487 = v440
				v488 = v441
				v489 = v442
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 3:
				v481 = v441
				v482 = v442
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 4:
				v477 = v441
				v478 = v442
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 5:
				v471 = v441
				v472 = v442
				v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
				v477 = v473<<(uint(int32(8))%32) + v471
				v478 = v472
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 6:
				v465 = v441
				v466 = v442
				v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+6)))
				v471 = v467<<(uint(int32(16))%32) + v465
				v472 = v466
				v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
				v477 = v473<<(uint(int32(8))%32) + v471
				v478 = v472
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 7:
				v460 = v442
				v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+7)))
				v465 = v461<<(uint(int32(24))%32) + v441
				v466 = v460
				v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+6)))
				v471 = v467<<(uint(int32(16))%32) + v465
				v472 = v466
				v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
				v477 = v473<<(uint(int32(8))%32) + v471
				v478 = v472
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 8:
				v455 = v442
				v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+8)))
				v460 = v456<<(uint(int32(8))%32) + v455
				v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+7)))
				v465 = v461<<(uint(int32(24))%32) + v441
				v466 = v460
				v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+6)))
				v471 = v467<<(uint(int32(16))%32) + v465
				v472 = v466
				v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
				v477 = v473<<(uint(int32(8))%32) + v471
				v478 = v472
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 9:
				v450 = v442
				v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+9)))
				v455 = v451<<(uint(int32(16))%32) + v450
				v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+8)))
				v460 = v456<<(uint(int32(8))%32) + v455
				v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+7)))
				v465 = v461<<(uint(int32(24))%32) + v441
				v466 = v460
				v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+6)))
				v471 = v467<<(uint(int32(16))%32) + v465
				v472 = v466
				v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
				v477 = v473<<(uint(int32(8))%32) + v471
				v478 = v472
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			case 10:
				v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+10)))
				v450 = v446<<(uint(int32(24))%32) + v442
				v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+9)))
				v455 = v451<<(uint(int32(16))%32) + v450
				v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+8)))
				v460 = v456<<(uint(int32(8))%32) + v455
				v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+7)))
				v465 = v461<<(uint(int32(24))%32) + v441
				v466 = v460
				v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+6)))
				v471 = v467<<(uint(int32(16))%32) + v465
				v472 = v466
				v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
				v477 = v473<<(uint(int32(8))%32) + v471
				v478 = v472
				v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
				v481 = v477 + v479
				v482 = v478
				v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+3)))
				v487 = v483<<(uint(int32(24))%32) + v440
				v488 = v481
				v489 = v482
				v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+2)))
				v494 = v490<<(uint(int32(16))%32) + v487
				v495 = v488
				v496 = v489
				v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
				v501 = v497<<(uint(int32(8))%32) + v494
				v502 = v495
				v503 = v496
				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
				v508 = v501 + v504
				v509 = v502
				v510 = v503
			default:
				v508 = v440
				v509 = v441
				v510 = v442
			}
		} else {
			if base.Ui32(v275) < base.Ui32(int32(12)) {
				v336 = v273
				v337 = v275
				v338 = v281
				v339 = v281
				v340 = v281
			} else {
				v288 = v273
				v289 = v275
				v290 = v281
				v291 = v281
				v292 = v281
				for {
					v294 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
					v295 = v294 + v291
					v296 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
					v298 = *(*int32)(unsafe.Add(mBase, uint32(v288)+8))
					v299 = v298 + v292
					v301 = int32(4)
					v303 = v296 + v290 - v299 ^ base.I32_rotl(v299, v301)
					v307 = v295 - v303 ^ base.I32_rotl(v303, int32(6))
					v308 = v299 + v295
					v309 = v303 + v308
					v310 = v307 + v309
					v314 = v308 - v307 ^ base.I32_rotl(v307, int32(8))
					v318 = v309 - v314 ^ base.I32_rotl(v314, int32(16))
					v322 = v310 - v318 ^ base.I32_rotl(v318, int32(19))
					v323 = v314 + v310
					v324 = v318 + v323
					v325 = v322 + v324
					v329 = v323 - v322 ^ base.I32_rotl(v322, v301)
					v330 = int32(12)
					v331 = v288 + v330
					v333 = v289 - v330
					if base.Ui32(int32(11)) < base.Ui32(v333) {
						v288 = v331
						v289 = v333
						v290 = v324
						v291 = v325
						v292 = v329
						continue
					} else {
						break
					}
					break
				}
				v336 = v331
				v337 = v333
				v338 = v324
				v339 = v325
				v340 = v329
			}
			switch v337 - int32(1) {
			case 0:
				v387 = v338
				v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
				v508 = v387 + v388
				v509 = v339
				v510 = v340
			case 1:
				v382 = v338
				v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+1)))
				v387 = v383<<(uint(int32(8))%32) + v382
				v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
				v508 = v387 + v388
				v509 = v339
				v510 = v340
			case 2:
				v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+2)))
				v382 = v378<<(uint(int32(16))%32) + v338
				v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+1)))
				v387 = v383<<(uint(int32(8))%32) + v382
				v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
				v508 = v387 + v388
				v509 = v339
				v510 = v340
			case 3:
				v375 = v339
				v376 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v508 = v376 + v338
				v509 = v375
				v510 = v340
			case 4:
				v372 = v339
				v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+4)))
				v375 = v372 + v373
				v376 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v508 = v376 + v338
				v509 = v375
				v510 = v340
			case 5:
				v367 = v339
				v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+5)))
				v372 = v368<<(uint(int32(8))%32) + v367
				v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+4)))
				v375 = v372 + v373
				v376 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v508 = v376 + v338
				v509 = v375
				v510 = v340
			case 6:
				v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+6)))
				v367 = v363<<(uint(int32(16))%32) + v339
				v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+5)))
				v372 = v368<<(uint(int32(8))%32) + v367
				v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+4)))
				v375 = v372 + v373
				v376 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v508 = v376 + v338
				v509 = v375
				v510 = v340
			case 7:
				v358 = v340
				v359 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v361 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
				v508 = v359 + v338
				v509 = v361 + v339
				v510 = v358
			case 8:
				v353 = v340
				v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+8)))
				v358 = v354<<(uint(int32(8))%32) + v353
				v359 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v361 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
				v508 = v359 + v338
				v509 = v361 + v339
				v510 = v358
			case 9:
				v348 = v340
				v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+9)))
				v353 = v349<<(uint(int32(16))%32) + v348
				v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+8)))
				v358 = v354<<(uint(int32(8))%32) + v353
				v359 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v361 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
				v508 = v359 + v338
				v509 = v361 + v339
				v510 = v358
			case 10:
				v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+10)))
				v348 = v344<<(uint(int32(24))%32) + v340
				v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+9)))
				v353 = v349<<(uint(int32(16))%32) + v348
				v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+8)))
				v358 = v354<<(uint(int32(8))%32) + v353
				v359 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
				v361 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
				v508 = v359 + v338
				v509 = v361 + v339
				v510 = v358
			default:
				v508 = v338
				v509 = v339
				v510 = v340
			}
		}
		v513 = int32(14)
		v515 = v509 ^ v510 - base.I32_rotl(v509, v513)
		v519 = v515 ^ v508 - base.I32_rotl(v515, int32(11))
		v523 = v519 ^ v509 - base.I32_rotl(v519, int32(25))
		v527 = v523 ^ v515 - base.I32_rotl(v523, int32(16))
		v531 = v527 ^ v519 - base.I32_rotl(v527, int32(4))
		v535 = v531 ^ v523 - base.I32_rotl(v531, v513)
		v549 = v535 ^ v527 - base.I32_rotl(v535, int32(24)) + (v268<<(uint(int32(6))%32) + int32(base.Ui32(v268)>>(uint(int32(2))%32))) - int32(1640531527) ^ v268
	} else {
		v549 = v268
	}
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v550 != 0 {
		v551 = int32(0)
		v555 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
		v556 = F_hash_bytes_uint32(m, v555)
		mBase = m.M
		v557 = *(*int32)(unsafe.Add(mBase, uint32(v550)+4))
		v558 = F_hash_bytes_uint32(m, v557)
		mBase = m.M
		v559 = int32(1640531527)
		v560 = v556 - v559
		v569 = v558 + v560<<(uint(int32(6))%32) + int32(base.Ui32(v560)>>(uint(int32(2))%32)) - v559 ^ v560
		v570 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
		if v551 < v570 {
			v576 = v569
			v577 = v551
			v578 = v570
			for {
				v586 = *(*int32)(unsafe.Add(mBase, uint32(v550+int32(88)+v578<<(uint(int32(4))%32)+v577*int32(100))))
				v587 = F_hash_bytes_uint32(m, v586)
				mBase = m.M
				v596 = v587 + (v576<<(uint(int32(6))%32) + int32(base.Ui32(v576)>>(uint(int32(2))%32))) - int32(1640531527) ^ v576
				v598 = v577 + int32(1)
				v599 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
				if v598 < v599 {
					v576 = v596
					v577 = v598
					v578 = v599
					continue
				} else {
					break
				}
				break
			}
			v602 = v596
		} else {
			v602 = v569
		}
		v615 = v602 + (v549<<(uint(int32(6))%32) + int32(base.Ui32(v549)>>(uint(int32(2))%32))) - int32(1640531527) ^ v549
	} else {
		v615 = v549
	}
	return v615
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
	v24 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
	v72 = v29 + (v67+int32(7))&int32(131064)
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
	v216 = v78 + (v211+int32(9))&int32(131064)
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
	v133 = int32(5646)
	goto L39
L38:
	;
	v133 = int32(5645)
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
		v114 = v114 + (v159+int32(7))&int32(131064) + int32(8)
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
	v12 = *(*int32)(unsafe.Add(mBase, _consts[1463]))
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
			F_errstart_cold(m, int32(21), int32(581856))
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
					F_errmsg(m, int32(542953), v9+int32(16))
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
							F_errfinish(m, int32(27512), int32(3569), int32(413491))
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
		F_errstart_cold(m, int32(21), int32(581856))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v62
			F_errmsg_internal(m, int32(505879), v9)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				F_errfinish(m, int32(27512), int32(3580), int32(413491))
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
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v963 int32
	_ = v963
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v994 int32
	_ = v994
	v4 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _consts[477]))
	v26 = *(*int32)(unsafe.Add(mBase, _consts[478]))
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
	return v994
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
	v40 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v40
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
	v46 = F_format_elog_string(m, int32(673732), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v46
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
	v994 = v4
	goto L1
L12:
	;
	v963 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v963
	goto L277
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
	v897 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v897
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
	v89 = int32(551611)
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
	v141 = int32(556918)
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
	v194 = int32(548013)
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
	v255 = int32(555477)
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
	v309 = int32(567577)
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
	v361 = int32(533718)
	goto L111
L109:
	;
	v462 = v84
	v463 = int32(534174)
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
	v404 = int32(551618)
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
	v611 = int32(545160)
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
	v507 = int32(547075)
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
	v552 = int32(551615)
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
	v661 = F_find_option(m, int32(400081), v649, v649, int32(21))
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
	v676 = int32(4552448)
	v682 = F_pg_snprintf(m, v676, int32(256), int32(354383), v655+int32(16))
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
	v668 = int32(4552448)
	v672 = F_pg_snprintf(m, v668, int32(256), int32(509851), v655)
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
	v686 = int32(790160)
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
	F_errmsg_internal(m, int32(191935), v655+int32(32))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L2
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(523571), int32(3036), int32(361582))
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
	v754 = int32(285807)
	goto L233
L232:
	;
	v754 = int32(355204)
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
	*(*int32)(unsafe.Add(mBase, uint32(v655)+64)) = int32(400081)
	F_errmsg(m, int32(747070), v655-int32(-64))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L2
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v655)+48)) = int32(165147)
	F_errdetail(m, int32(635986), v655+int32(48))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L2
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(523571), int32(4419), int32(346753))
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
	v903 = F_format_elog_string(m, int32(618423), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L2
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v903
	v994 = v4
	goto L1
L263:
	;
	if v907 == int32(0) {
		v994 = v4
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
	v929 = F_strlen(m, v907)
	mBase = m.M
	v930 = v929 + v907
	switch v862 {
	case 0:
		goto L273
	case 1:
		goto L272
	default:
		goto L271
	}
L266:
	;
	v924 = int32(*(*uint8)(unsafe.Add(mBase, _consts[479])))
	*(*uint8)(unsafe.Add(mBase, uint32(v907)+8)) = uint8(v924)
	v927 = *(*int64)(unsafe.Add(mBase, _consts[480]))
	*(*int64)(unsafe.Add(mBase, uint32(v907))) = v927
	goto L265
L267:
	;
	v918 = *(*int32)(unsafe.Add(mBase, _consts[481]))
	*(*int32)(unsafe.Add(mBase, uint32(v907)+3)) = v918
	v921 = *(*int32)(unsafe.Add(mBase, _consts[482]))
	*(*int32)(unsafe.Add(mBase, uint32(v907))) = v921
	goto L265
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v907))) = int32(5001555)
	goto L265
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v907))) = int32(5198665)
	goto L265
L270:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_bms_free(m, v949)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L2
	} else {
		goto L274
	}
L271:
	;
	v944 = int32(*(*uint16)(unsafe.Add(mBase, _consts[483])))
	*(*uint16)(unsafe.Add(mBase, uint32(v930)+4)) = uint16(v944)
	v947 = *(*int32)(unsafe.Add(mBase, _consts[484]))
	*(*int32)(unsafe.Add(mBase, uint32(v930))) = v947
	goto L270
L272:
	;
	v938 = int32(*(*uint16)(unsafe.Add(mBase, _consts[485])))
	*(*uint16)(unsafe.Add(mBase, uint32(v930)+4)) = uint16(v938)
	v941 = *(*int32)(unsafe.Add(mBase, _consts[486]))
	*(*int32)(unsafe.Add(mBase, uint32(v930))) = v941
	goto L270
L273:
	;
	v932 = int32(*(*uint16)(unsafe.Add(mBase, _consts[487])))
	*(*uint16)(unsafe.Add(mBase, uint32(v930)+4)) = uint16(v932)
	v935 = *(*int32)(unsafe.Add(mBase, _consts[488]))
	*(*int32)(unsafe.Add(mBase, uint32(v930))) = v935
	goto L270
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v907
	v954 = F_guc_malloc(m, int32(8))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L2
	} else {
		goto L275
	}
L275:
	;
	if v954 == int32(0) {
		v994 = v4
		goto L1
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v954))) = v863
	*(*int32)(unsafe.Add(mBase, uint32(v954)+4)) = v862
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v954
	v994 = int32(1)
	goto L1
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v84
	v969 = F_format_elog_string(m, int32(696825), v21)
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L2
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v969
	F_pfree(m, v28)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L2
	} else {
		goto L279
	}
L279:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_list_free(m, v974)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L2
	} else {
		goto L280
	}
L280:
	;
	v994 = v4
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
	v25 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v25
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
	v31 = F_format_elog_string(m, int32(673732), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v31
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
	v171 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v171
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
	v62 = int32(33883)
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
	v104 = int32(410304)
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
	v177 = F_format_elog_string(m, int32(696825), v10)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v177
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
							F_errmsg(m, int32(266154), int32(0))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_errdetail(m, int32(675084), int32(0))
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
										F_errfinish(m, int32(519077), int32(3282), int32(232229))
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
						if base.Ui32(int32(12000)) <= base.Ui32(v16) {
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
									F_errmsg(m, int32(388045), int32(0))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										F_errdetail(m, int32(675173), int32(0))
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
												F_errfinish(m, int32(519077), int32(3298), int32(232229))
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
	v25 = v13 + (int32(base.Ui32(v10)>>(uint(v14)%32))&int32(2047)+int32(base.Ui32(v10)>>(uint(int32(12))%32))+v14)&int32(4194302)
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
	v66 = v59 & int32(16383)
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
	if base.Ui32(v74) < base.Ui32(v43+v72&int32(65535)<<(uint(int32(1))%32)) {
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
	v22 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
	if v51 != int32(65534) {
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
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
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
	return v99
L2:
	;
	v87 = F_range_(m, l0, v81, v81, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L25
	} else {
		goto L26
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v81 = v14
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
	v27 = int32(556701)
	v30 = int32(1652320)
	goto L7
L6:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	if v76 != 0 {
		v99 = l3
		goto L1
	} else {
		goto L24
	}
L7:
	;
	v32 = F_strlen(m, v27)
	mBase = m.M
	if v32 == v17 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	return l3
L9:
	;
	if v17 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v70 = v30 + int32(8)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v71 != 0 {
		v27 = v71
		v30 = v70
		goto L7
	} else {
		goto L23
	}
L12:
	;
	if v66 == int32(0) {
		goto L6
	} else {
		goto L22
	}
L13:
	;
	v66 = int32(0)
	goto L12
L14:
	;
	v38 = v27
	v39 = l1
	v40 = v17
	goto L15
L15:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v43 != v44 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L13
L17:
	;
	v66 = v43 - v44
	goto L12
L18:
	;
	goto L19
L19:
	;
	if v43 == int32(0) {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v49 = int32(1)
	v54 = v40 - v49
	if v54 != 0 {
		v38 = v38 + v49
		v39 = v39 + int32(4)
		v40 = v54
		goto L15
	} else {
		goto L21
	}
L21:
	;
	goto L16
L22:
	;
	goto L11
L23:
	;
	goto L8
L24:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
	v81 = v78
	goto L2
L25:
	;
	return int32(0)
L26:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v91 == int32(0) {
		v99 = l3
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v99 = v95
	goto L1
}
func F_cidr_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_network_in(m, v2, int32(1), v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_clause_is_strict_for(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	v4 = int32(0)
	if l0 == v4 {
		v163 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v163 & int32(1)
L2:
	;
	if l1 == int32(0) {
		v163 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = l0
	v12 = l1
	v13 = l2
	v15 = v4
	goto L4
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v17 == int32(27) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v163 = v158
	goto L1
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v21 = v20
	goto L8
L7:
	;
	v21 = v12
	goto L8
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v22 != int32(27) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v26 = v11
	goto L11
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v26 = v25
	goto L11
L11:
	;
	v27 = F_equal(m, v26, v21)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v27 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v163 = int32(1)
	goto L1
L15:
	;
	goto L16
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v32 == int32(17) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v36 = F_op_strict(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L20
	}
L18:
	;
	v64 = v32
	goto L19
L19:
	;
	if v64 == int32(15) {
		goto L40
	} else {
		goto L41
	}
L20:
	;
	if v36 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	if v38 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v64 = v63
	goto L19
L24:
	;
	v163 = v15
	goto L1
L25:
	;
	goto L26
L26:
	;
	v41 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v42 <= v41 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v163 = v15
	goto L1
L28:
	;
	goto L29
L29:
	;
	v48 = v41
	goto L30
L30:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v48<<(uint(int32(2))%32))))
	v57 = F_clause_is_strict_for(m, v55, v21, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L32
	}
L31:
	;
	v163 = v57
	goto L1
L32:
	;
	if v57 != 0 {
		v163 = v57
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v60 = v48 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v60 < v61 {
		v48 = v60
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v154 = int32(0)
	if v150 == v154 {
		goto L77
	} else {
		goto L78
	}
L36:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
	v163 = v149
	goto L1
L37:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v105 = F_clause_is_strict_for(m, v103, v21, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L12
	} else {
		goto L56
	}
L38:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v150 = v99
	goto L35
L39:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	if v74 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v68 = F_func_strict(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L43
	}
L41:
	;
	v71 = v64
	goto L42
L42:
	;
	switch v71 - int32(7) {
	case 0:
		goto L36
	default:
		v163 = v15
		goto L1
	case 13:
		goto L37
	case 21, 22, 23, 48:
		goto L38
	}
L43:
	;
	if v68 != 0 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v71 = v70
	goto L42
L45:
	;
	v163 = v15
	goto L1
L46:
	;
	goto L47
L47:
	;
	v77 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v78 <= v77 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v163 = v15
	goto L1
L49:
	;
	goto L50
L50:
	;
	v84 = v77
	goto L51
L51:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v84<<(uint(int32(2))%32))))
	v93 = F_clause_is_strict_for(m, v91, v21, int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L53
	}
L52:
	;
	v163 = v93
	goto L1
L53:
	;
	if v93 != 0 {
		v163 = v93
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v96 = v84 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v96 < v97 {
		v84 = v96
		goto L51
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	if v105 == int32(0) {
		v150 = v102
		goto L35
	} else {
		goto L57
	}
L57:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v110 = F_op_strict(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	if v110 == int32(0) {
		v150 = v102
		goto L35
	} else {
		goto L59
	}
L59:
	;
	if v13&int32(1) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v102 == int32(0) {
		v163 = v15
		goto L1
	} else {
		goto L63
	}
L61:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+20)))
	if v118 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v163 = int32(1)
	goto L1
L63:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v124 != int32(35) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v145 <= int32(0) {
		v150 = v102
		goto L35
	} else {
		goto L76
	}
L65:
	;
	if v124 != int32(7) {
		v150 = v102
		goto L35
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+20)))
	if v139 != 0 {
		v150 = v102
		goto L35
	} else {
		goto L74
	}
L68:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+24)))
	if v129 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v163 = int32(1)
	goto L1
L70:
	;
	goto L71
L71:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v102)+20))
	v132 = F_pg_detoast_datum(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L72
	}
L72:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v137 = F_ArrayGetNItems(m, v134, v132+int32(16))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	v145 = v137
	goto L64
L74:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	if v140 == int32(0) {
		v150 = v102
		goto L35
	} else {
		goto L75
	}
L75:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	v145 = v143
	goto L64
L76:
	;
	v163 = int32(1)
	goto L1
L77:
	;
	v163 = int32(0)
	goto L1
L78:
	;
	goto L79
L79:
	;
	v158 = int32(0)
	if v21 != 0 {
		v11 = v150
		v12 = v21
		v13 = v154
		v15 = v158
		goto L4
	} else {
		goto L80
	}
L80:
	;
	goto L5
}
func F_clean_stopword_intree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	F_check_stack_depth(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v17
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
		switch v22 - int32(1) {
		case 0:
			v122 = l0
			m.G0 = v11 + int32(16)
			return v122
		default:
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
			if v28 == int32(1) {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v32 = F_clean_stopword_intree(m, v31, l1, l2)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v32
					if v32 != 0 {
						v122 = l0
						m.G0 = v11 + int32(16)
						return v122
					} else {
						F_freetree(m, l0)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							v122 = int32(0)
							m.G0 = v11 + int32(16)
							return v122
						}
					}
				}
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v40 = F_clean_stopword_intree(m, v35, v11+int32(12), v11+int32(8))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v46 = F_clean_stopword_intree(m, v43, v11+int32(4), v11)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
						if v50 == int32(4) {
							v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+2)))
							v54 = v53
						} else {
							v54 = int32(0)
						}
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v55 == int32(0) {
							if v46 == int32(0) {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
								if v50 == int32(4) {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v70 = v63 + (v60 + v54)
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									if v66 < v60 {
										v68 = v60
									} else {
										v68 = v66
									}
									v70 = v68
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v70
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v70
								F_freetree(m, l0)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v122 = int32(0)
									m.G0 = v11 + int32(16)
									return v122
								}
							} else {
								if v50 == int32(4) {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
									v80 = v75 + (v76 + v54)
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v80 = v79
								}
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v80
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v82
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								F_pfree(m, l0)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									v122 = v84
									m.G0 = v11 + int32(16)
									return v122
								}
							}
						} else {
							if v46 == int32(0) {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v91
								if v50 == int32(4) {
									v97 = v89 + (v90 + v54)
								} else {
									v97 = v90
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v97
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_pfree(m, l0)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									v122 = v99
									m.G0 = v11 + int32(16)
									return v122
								}
							} else {
								if v50 != int32(4) {
									v122 = l0
								} else {
									v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+2)))
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
									v108 = v104 + (v105 + v106)
									*(*uint16)(unsafe.Add(mBase, uint32(v49)+2)) = uint16(v108)
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v110
									v112 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v112
									v122 = l0
								}
								m.G0 = v11 + int32(16)
								return v122
							}
						}
					}
				}
			}
		case 2:
			F_pfree(m, l0)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v122 = int32(0)
				m.G0 = v11 + int32(16)
				return v122
			}
		}
	}
}
func F_clogsyncfiletag(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_SlruSyncFileTag(m, int32(4448656), l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_cmpNodePtr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+8)))
	return v3 - v4
}
func F_cmp_fxid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(base.Ui64(v6) < base.Ui64(v5)) - base.B2i32(base.Ui64(v5) < base.Ui64(v6))
}
func F_cmp_lsn(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	return base.B2i32(base.Ui64(v6) < base.Ui64(v5)) - base.B2i32(base.Ui64(v5) < base.Ui64(v6))
}
func F_cmpaffix(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v84 int32
	_ = v84
	v8 = int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = v9 & v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = v12 & v8
	if base.Ui32(v11) < base.Ui32(v14) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	goto L3
L3:
	;
	if base.Ui32(v14) < base.Ui32(v11) {
		v84 = v8
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v84
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v26 == int32(0) {
		v45 = v25
		v46 = v26
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	v49 = F_strlen(m, v20)
	mBase = m.M
	v50 = F_strlen(m, v19)
	mBase = m.M
	v51 = v49
	v52 = v50
	goto L17
L9:
	;
	return v46 - v45
L10:
	;
	goto L9
L11:
	;
	if v25 != v26 {
		v45 = v25
		v46 = v26
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v30 = v20
	v31 = v19
	goto L13
L13:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v35 == int32(0) {
		v45 = v34
		v46 = v35
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v45 = v34
	v46 = v35
	goto L10
L15:
	;
	v38 = int32(1)
	if v34 == v35 {
		v30 = v30 + v38
		v31 = v31 + v38
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v58 = int32(1)
	v59 = v52 - v58
	v61 = v51 - v58
	if v61 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v61 < v59 {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	goto L18
L20:
	;
	if v59 < int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v20))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v19))))
	if base.Ui32(v67) < base.Ui32(v69) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(-1)
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(v67) <= base.Ui32(v69) {
		v51 = v61
		v52 = v59
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v84 = v8
	goto L4
L26:
	;
	return int32(-1)
L27:
	;
	goto L28
L28:
	;
	v84 = base.B2i32(v59 < v61)
	goto L4
}
func F_cnt_sml(m *base.Module, l0 int32, l1 int32, l2 int32) float32 {
	mBase := m.M
	_ = mBase
	var v10 float32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v91 float32
	_ = v91
	v10 = float32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = int32(2)
	v14 = int32(5)
	v15 = int32(base.Ui32(v11)>>(uint(v12)%32)) - v14
	v16 = int32(3)
	v17 = base.I32_div_u_s(v15, v16)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = int32(base.Ui32(v18)>>(uint(v12)%32)) - v14
	v24 = base.I32_div_u_s(v22, v16)
	if base.Ui32(v22) < base.Ui32(v16) {
		v91 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v91
L2:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v91 = v10
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = int32(5)
	v30 = l0 + v29
	v32 = l1 + v29
	v33 = v30
	v34 = v32
	v36 = int32(0)
	goto L4
L4:
	;
	v45 = base.I32_div_s(v34-v32, int32(3))
	if v45 < v17 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if l2 != 0 {
		goto L19
	} else {
		goto L20
	}
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[1559]))
	v49 = m.T0[v48].(func(*base.Module, int32, int32) int32)(m, v33, v34)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v74 = v36
	goto L8
L8:
	;
	goto L5
L9:
	;
	v70 = base.I32_div_s(v65-v30, int32(3))
	if v70 < v24 {
		v33 = v65
		v34 = v66
		v36 = v67
		goto L4
	} else {
		goto L18
	}
L10:
	;
	return float32(0)
L11:
	;
	if int32(0) <= v49 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v49 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v61 = v34
	v62 = v36
	goto L14
L14:
	;
	v65 = v33 + int32(3)
	v66 = v61
	v67 = v62
	goto L9
L15:
	;
	v65 = v33
	v66 = v34 + int32(3)
	v67 = v36
	goto L9
L16:
	;
	goto L17
L17:
	;
	v61 = v34 + int32(3)
	v62 = v36 + int32(1)
	goto L14
L18:
	;
	v74 = v67
	goto L8
L19:
	;
	v78 = v74
	goto L21
L20:
	;
	v78 = v17
	goto L21
L21:
	;
	v91 = base.F32_div(base.F32_convert_i32_s(v74), base.F32_convert_i32_s(v24-v74+v78))
	goto L1
}
func F_collect_visibility_data(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v3
	v18 = F_GetAccessStrategy(m, int32(1))
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
	v23 = F_relation_open(m, l0, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L52
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+119)))
	v28 = v26 - int32(109)
	if base.Ui32(int32(7)) < base.Ui32(v28) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if int32(1)<<(uint(v28)%32)&int32(161) == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v38 = F_RelationGetNumberOfBlocksInFork(m, v23, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v42 = F_palloc0(m, v38+int32(8))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(0)
	if l1 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v38
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v48
	v55 = F_read_stream_begin_relation(m, int32(12), v18, v23, int32(120), v13+int32(4), v48)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v57 = v3
	goto L11
L11:
	;
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v57 = v55
	goto L11
L13:
	;
	v59 = v42 + int32(8)
	v61 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	if l1 != 0 {
		goto L43
	} else {
		goto L44
	}
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v77 = F_visibilitymap_get_status(m, v23, v61, v13+int32(12))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	if v77&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v81 = v61 + v59
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v84 = v82 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v84)
	goto L25
L24:
	;
	goto L25
L25:
	;
	if v77&int32(2) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v89 = v61 + v59
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v92 = v90 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v92)
	goto L28
L27:
	;
	goto L28
L28:
	;
	if l1 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v96 = F_read_stream_next_buffer(m, v57, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v133 = v61 + int32(1)
	if v133 != v38 {
		v61 = v133
		goto L16
	} else {
		goto L42
	}
L32:
	;
	F_LockBuffer(m, v96, int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v96 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+10)))
	if v119&int32(4) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v104+(v96^int32(-1))<<(uint(int32(2))%32))))
	v118 = v110
	goto L34
L36:
	;
	goto L37
L37:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v118 = v112 + v96<<(uint(int32(13))%32) + int32(-8192)
	goto L34
L38:
	;
	v122 = v61 + v59
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v125 = v123 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v125)
	goto L40
L39:
	;
	goto L40
L40:
	;
	F_UnlockReleaseBuffer(m, v96)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	goto L31
L42:
	;
	goto L17
L43:
	;
	F_read_stream_end(m, v57)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v147 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	F_ReleaseBuffer(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_relation_close(m, v23, int32(1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	m.G0 = v13 + int32(16)
	return v42
L52:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v164 + int32(4)
	F_errmsg(m, int32(445182), v13)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v174 = int32(*(*int8)(unsafe.Add(mBase, uint32(v173)+119)))
	F_errdetail_relkind_not_supported(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(514419), int32(951), int32(444998))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_compare3(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	v6 = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v8) < base.Ui32(v7) {
		v24 = v6
	} else {
		v10 = base.B2i32(v7 != v8)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if base.B2i32(v10 == int32(0))&base.B2i32(base.Ui32(v14) < base.Ui32(v13)) != 0 {
			v24 = v6
		} else {
			v24 = int32(0) - (v10 | base.B2i32(v13 != v14))
		}
	}
	return v24
}
func F_compare_fractional_path_costs(m *base.Module, l0 int32, l1 int32, l2 float64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v45 float64
	_ = v45
	var v51 int32
	_ = v51
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v8 != v9 {
		if v8 < v9 {
			v14 = int32(-1)
		} else {
			v14 = int32(1)
		}
		return v14
	} else {
		if base.F64_le(l2, float64(0))|base.F64_ge(l2, float64(1)) != 0 {
			v21 = int32(-1)
			v22 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
			v23 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			if base.F64_lt(v22, v23) != 0 {
				v51 = v21
				return v51
			} else {
				if base.F64_gt(v22, v23) != 0 {
					return int32(1)
				} else {
					v28 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
					v29 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
					if base.F64_lt(v28, v29) != 0 {
						v51 = v21
						return v51
					} else {
						if base.F64_gt(v28, v29) != 0 {
							v51 = int32(1)
							return v51
						} else {
							return int32(0)
						}
					}
				}
			}
		} else {
			v36 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
			v37 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
			v40 = base.F64_add(base.F64_mul(l2, base.F64_sub(v36, v37)), v37)
			v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			v42 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
			v45 = base.F64_add(base.F64_mul(l2, base.F64_sub(v41, v42)), v42)
			if base.F64_lt(v40, v45) != 0 {
				v51 = int32(-1)
			} else {
				v51 = base.F64_lt(v45, v40)
			}
			return v51
		}
	}
}
func F_compare_subnode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	if l2 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	v15 = l1 + l2
	v17 = l0 + int32(2)
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v19 = v17 + v18
	v21 = l1
	goto L3
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v32 != int32(95) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L1
L5:
	;
	if v18 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L6:
	;
	v49 = v21
	goto L14
L7:
	;
	if base.Ui32(v21) < base.Ui32(v15) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v36 = F_pg_mblen_range(m, v21, v15)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v68 = v21
	goto L5
L11:
	;
	return int32(0)
L12:
	;
	v40 = v36 + v21
	if base.Ui32(v40) < base.Ui32(v15) {
		v21 = v40
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L1
L14:
	;
	v54 = F_pg_mblen_range(m, v49, v15)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v68 = v56
	goto L5
L16:
	;
	v56 = v54 + v49
	if base.Ui32(v15) <= base.Ui32(v56) {
		v68 = v56
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v58 != int32(95) {
		v49 = v56
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	return int32(0)
L20:
	;
	goto L21
L21:
	;
	v77 = v68 - v21
	v83 = v17
	goto L23
L22:
	;
	if base.Ui32(v68) < base.Ui32(v15) {
		v21 = v21 + v77
		goto L3
	} else {
		goto L46
	}
L23:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v90 != int32(95) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	return int32(0)
L25:
	;
	v135 = int32(0)
	if base.B2i32(l4&base.B2i32(v77 < v129) == v135)&base.B2i32(v129 != v77) == v135 {
		goto L40
	} else {
		goto L41
	}
L26:
	;
	v103 = v83
	goto L33
L27:
	;
	if base.Ui32(v83) < base.Ui32(v19) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v95 = F_pg_mblen_range(m, v83, v19)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L11
	} else {
		goto L31
	}
L30:
	;
	v129 = int32(0)
	goto L25
L31:
	;
	v97 = v95 + v83
	if base.Ui32(v97) < base.Ui32(v19) {
		v83 = v97
		goto L23
	} else {
		goto L32
	}
L32:
	;
	return int32(0)
L33:
	;
	v113 = F_pg_mblen_range(m, v103, v19)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L11
	} else {
		goto L35
	}
L34:
	;
	v129 = v115 - v83
	goto L25
L35:
	;
	v115 = v113 + v103
	if base.Ui32(v115) < base.Ui32(v19) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v117 != int32(95) {
		v103 = v115
		goto L33
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L34
L39:
	;
	goto L38
L40:
	;
	v141 = m.T0[l3].(func(*base.Module, int32, int32, int32, int32) int32)(m, v21, v77, v83, v129)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L11
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v143 = v83 + v129
	if base.Ui32(v143) < base.Ui32(v19) {
		v83 = v143
		goto L23
	} else {
		goto L45
	}
L43:
	;
	if v141 != 0 {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	goto L24
L46:
	;
	goto L4
}
func F_comparecost_1(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	return base.B2i32(v4 < v3) - base.B2i32(v3 < v4)
}
func F_compute_new_xmax_infomask(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
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
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v746 int32
	_ = v746
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v903 int32
	_ = v903
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1044 int32
	_ = v1044
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	v10 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	if l1&int32(2048) != 0 {
		v1027 = l4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l7))) = uint16(v1068)
	*(*uint16)(unsafe.Add(mBase, uint32(l8))) = uint16(v1064)
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1066
	m.G0 = v19 + int32(80)
	return
L2:
	;
	if l5 != 0 {
		goto L321
	} else {
		goto L322
	}
L3:
	;
	if l1&int32(4096) != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L12
	} else {
		goto L315
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L12
	} else {
		goto L309
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L12
	} else {
		goto L303
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L12
	} else {
		goto L297
	}
L8:
	;
	if l1&int32(4304) == int32(4224) {
		v1027 = l4
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v432 = l2 & int32(8192)
	if v432 != 0 {
		goto L115
	} else {
		goto L116
	}
L11:
	;
	v37 = int32(base.Ui32(l1&int32(128))>>(uint(int32(7))%32)) | base.B2i32(l1&int32(4176) == int32(64))
	v38 = F_MultiXactIdIsRunning(m, l0, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	if v38 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v37 != 0 {
		v1027 = l4
		goto L2
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if l5 != 0 {
		goto L31
	} else {
		goto L32
	}
L17:
	;
	v42 = int32(0)
	v46 = F_GetMultiXactIdMembers(m, l0, v19+int32(76), v42)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	if int32(0) < v46 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	v52 = v42
	goto L24
L20:
	;
	v83 = v42
	goto L21
L21:
	;
	v98 = F_TransactionIdDidCommit(m, v83)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L12
	} else {
		goto L29
	}
L22:
	;
	F_pfree(m, v50)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L12
	} else {
		goto L28
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v79 = v77
	goto L22
L24:
	;
	v69 = v50 + v52<<(uint(int32(3))%32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v70) {
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v79 = int32(0)
	goto L22
L26:
	;
	v74 = v52 + int32(1)
	if v74 != v46 {
		v52 = v74
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v83 = v79
	goto L21
L29:
	;
	if v98 == int32(0) {
		v1027 = l4
		goto L2
	} else {
		goto L30
	}
L30:
	;
	goto L16
L31:
	;
	v122 = int32(8)
	goto L33
L32:
	;
	v122 = int32(4)
	goto L33
L33:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v122)+uint32(_consts[95])))
	if v126 == int32(-1) {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v129 = int32(0)
	v131 = m.G0
	v133 = v131 - int32(16)
	m.G0 = v133
	v138 = F_GetMultiXactIdMembers(m, l0, v133+int32(12), v129)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L12
	} else {
		goto L38
	}
L35:
	;
	m.G0 = v133 + int32(16)
	v293 = F_GetMultiXactIdMembers(m, v271, v19+int32(76), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L12
	} else {
		goto L70
	}
L36:
	;
	v260 = v252 + v242<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v260)+4)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = l3
	v265 = F_MultiXactIdCreateFromMembers(m, v242+int32(1), v252)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L12
	} else {
		goto L67
	}
L37:
	;
	v240 = F_palloc(m, int32(8))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L12
	} else {
		goto L66
	}
L38:
	;
	if int32(0) <= v138 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	if v138 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+8)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = l3
	v236 = F_MultiXactIdCreateFromMembers(m, int32(1), v133+int32(4))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L12
	} else {
		goto L65
	}
L42:
	;
	v149 = v129
	goto L43
L43:
	;
	v163 = v142 + v149<<(uint(int32(3))%32)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v164 != l3 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v173 = int32(1)
	if v138 <= v173 {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	v171 = v149 + int32(1)
	if v171 != v138 {
		v149 = v171
		goto L43
	} else {
		goto L49
	}
L46:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v166 != v126 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	F_pfree(m, v142)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	v271 = l0
	goto L35
L49:
	;
	goto L44
L50:
	;
	v176 = v173
	goto L52
L51:
	;
	v176 = v138
	goto L52
L52:
	;
	v182 = F_palloc(m, v138<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	v185 = int32(0)
	v189 = int32(0)
	goto L54
L54:
	;
	v203 = v142 + v189<<(uint(int32(3))%32)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v205 = F_TransactionIdIsInProgress(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L12
	} else {
		goto L57
	}
L55:
	;
	v242 = v226
	v252 = v182
	goto L36
L56:
	;
	v229 = v189 + int32(1)
	if v229 != v176 {
		v185 = v226
		v189 = v229
		goto L54
	} else {
		goto L64
	}
L57:
	;
	if v205 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	if base.Ui32(v209) < base.Ui32(int32(4)) {
		v226 = v185
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v219 = v182 + v185<<(uint(int32(3))%32)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v219)+4)) = v222
	v226 = v185 + int32(1)
	goto L56
L61:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v213 = F_TransactionIdDidCommit(m, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	if v213 == int32(0) {
		v226 = v185
		goto L56
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	goto L55
L65:
	;
	v271 = v236
	goto L35
L66:
	;
	v242 = int32(0)
	v252 = v240
	goto L36
L67:
	;
	F_pfree(m, v142)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	F_pfree(m, v252)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	v271 = v265
	goto L35
L70:
	;
	if v293 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v1064 = v129
	v1066 = v271
	v1068 = int32(4240)
	goto L1
L72:
	;
	goto L73
L73:
	;
	v298 = int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v293 == v298 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v293&v298 == int32(0) {
		v408 = v371
		v410 = v375
		v411 = v383
		goto L95
	} else {
		goto L96
	}
L75:
	;
	v303 = int32(0)
	v371 = v129
	v374 = v303
	v375 = v303
	v383 = v10
	goto L74
L76:
	;
	goto L77
L77:
	;
	v307 = int32(0)
	v310 = v307
	v311 = v129
	v314 = v307
	v315 = v307
	v323 = v10
	goto L78
L78:
	;
	v328 = v300 + v314<<(uint(int32(3))%32)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v329<<(uint(int32(2))%32))+uint32(_consts[96])))
	if base.Ui32(v315) < base.Ui32(v334) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v371 = v362
	v374 = v366
	v375 = v364
	v383 = v363
	goto L74
L80:
	;
	v336 = v334
	goto L82
L81:
	;
	v336 = v315
	goto L82
L82:
	;
	switch v329 - int32(3) {
	case 0:
		goto L86
	case 1:
		v343 = v311
		goto L84
	case 2:
		goto L85
	default:
		v345 = v311
		v346 = v323
		goto L83
	}
L83:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v347<<(uint(int32(2))%32))+uint32(_consts[96])))
	switch v347 - int32(3) {
	case 0:
		goto L90
	case 1:
		v360 = v345
		goto L88
	case 2:
		goto L89
	default:
		v362 = v345
		v363 = v346
		goto L87
	}
L84:
	;
	v345 = v343
	v346 = int32(1)
	goto L83
L85:
	;
	v343 = v311 | int32(8192)
	goto L84
L86:
	;
	v345 = v311 | int32(8192)
	v346 = v323
	goto L83
L87:
	;
	if base.Ui32(v336) < base.Ui32(v352) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v362 = v360
	v363 = int32(1)
	goto L87
L89:
	;
	v360 = v345 | int32(8192)
	goto L88
L90:
	;
	v362 = v345 | int32(8192)
	v363 = v346
	goto L87
L91:
	;
	v364 = v352
	goto L93
L92:
	;
	v364 = v336
	goto L93
L93:
	;
	v365 = int32(2)
	v366 = v314 + v365
	v368 = v310 + v365
	if v368 != v293&int32(2147483646) {
		v310 = v368
		v311 = v362
		v314 = v366
		v315 = v364
		v323 = v363
		goto L78
	} else {
		goto L94
	}
L94:
	;
	goto L79
L95:
	;
	F_pfree(m, v300)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L12
	} else {
		goto L103
	}
L96:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v300+v374<<(uint(int32(3))%32))+4))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v391<<(uint(int32(2))%32))+uint32(_consts[96])))
	if base.Ui32(v375) < base.Ui32(v396) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v398 = v396
	goto L99
L98:
	;
	v398 = v375
	goto L99
L99:
	;
	switch v391 - int32(3) {
	case 0:
		goto L102
	case 1:
		v405 = v371
		goto L100
	case 2:
		goto L101
	default:
		v408 = v371
		v410 = v398
		v411 = v383
		goto L95
	}
L100:
	;
	v408 = v405
	v410 = v398
	v411 = int32(1)
	goto L95
L101:
	;
	v405 = v371 | int32(8192)
	goto L100
L102:
	;
	v408 = v371 | int32(8192)
	v410 = v398
	v411 = v383
	goto L95
L103:
	;
	if v410&int32(-2) == int32(2) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if v411 != 0 {
		v1064 = v408
		v1066 = v271
		v1068 = int32(4160)
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v410 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v1064 = v408
	v1066 = v271
	v1068 = int32(4288)
	goto L1
L108:
	;
	v423 = int32(4096)
	goto L110
L109:
	;
	v423 = int32(4112)
	goto L110
L110:
	;
	if v410 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v426 = int32(4176)
	goto L113
L112:
	;
	v426 = v423
	goto L113
L113:
	;
	if v411 != 0 {
		v1064 = v408
		v1066 = v271
		v1068 = v426
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v1064 = v408
	v1066 = v271
	v1068 = v426 | int32(128)
	goto L1
L115:
	;
	v433 = int32(5)
	goto L117
L116:
	;
	v433 = int32(4)
	goto L117
L117:
	;
	if l1&int32(1024) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	if l5 != 0 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v597 = int32(base.Ui32(l1&int32(128))>>(uint(int32(7))%32)) | base.B2i32(l1&int32(80) == int32(64))
	v598 = F_TransactionIdIsInProgress(m, l0)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L12
	} else {
		goto L171
	}
L121:
	;
	v440 = int32(8)
	goto L123
L122:
	;
	v440 = int32(4)
	goto L123
L123:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v440)+uint32(_consts[95])))
	if v444 == int32(-1) {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	v447 = int32(0)
	v448 = F_MultiXactIdCreate(m, l0, v433, l3, v444)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L12
	} else {
		goto L125
	}
L125:
	;
	v453 = F_GetMultiXactIdMembers(m, v448, v19+int32(76), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L12
	} else {
		goto L126
	}
L126:
	;
	if v453 <= int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v1064 = v447
	v1066 = v448
	v1068 = int32(4240)
	goto L1
L128:
	;
	goto L129
L129:
	;
	v458 = int32(1)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v453 == v458 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v453&v458 == int32(0) {
		v568 = v531
		v570 = v535
		v571 = v543
		goto L151
	} else {
		goto L152
	}
L131:
	;
	v463 = int32(0)
	v531 = v447
	v534 = v463
	v535 = v463
	v543 = v10
	goto L130
L132:
	;
	goto L133
L133:
	;
	v467 = int32(0)
	v470 = v467
	v471 = v447
	v474 = v467
	v475 = v467
	v483 = v10
	goto L134
L134:
	;
	v488 = v460 + v474<<(uint(int32(3))%32)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v489<<(uint(int32(2))%32))+uint32(_consts[96])))
	if base.Ui32(v475) < base.Ui32(v494) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v531 = v522
	v534 = v526
	v535 = v524
	v543 = v523
	goto L130
L136:
	;
	v496 = v494
	goto L138
L137:
	;
	v496 = v475
	goto L138
L138:
	;
	switch v489 - int32(3) {
	case 0:
		goto L142
	case 1:
		v503 = v471
		goto L140
	case 2:
		goto L141
	default:
		v505 = v471
		v506 = v483
		goto L139
	}
L139:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v507<<(uint(int32(2))%32))+uint32(_consts[96])))
	switch v507 - int32(3) {
	case 0:
		goto L146
	case 1:
		v520 = v505
		goto L144
	case 2:
		goto L145
	default:
		v522 = v505
		v523 = v506
		goto L143
	}
L140:
	;
	v505 = v503
	v506 = int32(1)
	goto L139
L141:
	;
	v503 = v471 | int32(8192)
	goto L140
L142:
	;
	v505 = v471 | int32(8192)
	v506 = v483
	goto L139
L143:
	;
	if base.Ui32(v496) < base.Ui32(v512) {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	v522 = v520
	v523 = int32(1)
	goto L143
L145:
	;
	v520 = v505 | int32(8192)
	goto L144
L146:
	;
	v522 = v505 | int32(8192)
	v523 = v506
	goto L143
L147:
	;
	v524 = v512
	goto L149
L148:
	;
	v524 = v496
	goto L149
L149:
	;
	v525 = int32(2)
	v526 = v474 + v525
	v528 = v470 + v525
	if v528 != v453&int32(2147483646) {
		v470 = v528
		v471 = v522
		v474 = v526
		v475 = v524
		v483 = v523
		goto L134
	} else {
		goto L150
	}
L150:
	;
	goto L135
L151:
	;
	F_pfree(m, v460)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L12
	} else {
		goto L159
	}
L152:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v460+v534<<(uint(int32(3))%32))+4))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v551<<(uint(int32(2))%32))+uint32(_consts[96])))
	if base.Ui32(v535) < base.Ui32(v556) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v558 = v556
	goto L155
L154:
	;
	v558 = v535
	goto L155
L155:
	;
	switch v551 - int32(3) {
	case 0:
		goto L158
	case 1:
		v565 = v531
		goto L156
	case 2:
		goto L157
	default:
		v568 = v531
		v570 = v558
		v571 = v543
		goto L151
	}
L156:
	;
	v568 = v565
	v570 = v558
	v571 = int32(1)
	goto L151
L157:
	;
	v565 = v531 | int32(8192)
	goto L156
L158:
	;
	v568 = v531 | int32(8192)
	v570 = v558
	v571 = v543
	goto L151
L159:
	;
	if v570&int32(-2) == int32(2) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	if v571 != 0 {
		v1064 = v568
		v1066 = v448
		v1068 = int32(4160)
		goto L1
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if v570 != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v1064 = v568
	v1066 = v448
	v1068 = int32(4288)
	goto L1
L164:
	;
	v583 = int32(4096)
	goto L166
L165:
	;
	v583 = int32(4112)
	goto L166
L166:
	;
	if v570 == int32(1) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v586 = int32(4176)
	goto L169
L168:
	;
	v586 = v583
	goto L169
L169:
	;
	if v571 != 0 {
		v1064 = v568
		v1066 = v448
		v1068 = v586
		goto L1
	} else {
		goto L170
	}
L170:
	;
	v1064 = v568
	v1066 = v448
	v1068 = v586 | int32(128)
	goto L1
L171:
	;
	if v598 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	if v597 == int32(0) {
		v630 = v433
		goto L175
	} else {
		goto L176
	}
L173:
	;
	goto L174
L174:
	;
	if v597 != 0 {
		v1027 = l4
		goto L2
	} else {
		goto L244
	}
L175:
	;
	if l0 == l3 {
		goto L188
	} else {
		goto L189
	}
L176:
	;
	switch int32(base.Ui32(l1)>>(uint(int32(4))%32))&int32(5) - int32(1) {
	case 0:
		v630 = int32(0)
		goto L175
	case 1, 2:
		goto L179
	case 3:
		goto L180
	case 4:
		goto L177
	default:
		goto L178
	}
L177:
	;
	v630 = int32(1)
	goto L175
L178:
	;
	v614 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L12
	} else {
		goto L184
	}
L179:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	if v432 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v611 = int32(3)
	goto L183
L182:
	;
	v611 = int32(2)
	goto L183
L183:
	;
	v630 = v611
	goto L175
L184:
	;
	if v614 == int32(0) {
		v1027 = l4
		goto L2
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l0
	F_errmsg_internal(m, int32(45820), v19+int32(16))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L12
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(520720), int32(5499), int32(329578))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L12
	} else {
		goto L187
	}
L187:
	;
	v1027 = l4
	goto L2
L188:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v630<<(uint(int32(2))%32))+uint32(_consts[96])))
	if base.Ui32(v636) < base.Ui32(l4) {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	goto L190
L190:
	;
	if l5 != 0 {
		goto L194
	} else {
		goto L195
	}
L191:
	;
	v638 = l4
	goto L193
L192:
	;
	v638 = v636
	goto L193
L193:
	;
	v1027 = v638
	goto L2
L194:
	;
	v643 = int32(8)
	goto L196
L195:
	;
	v643 = int32(4)
	goto L196
L196:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v643)+uint32(_consts[95])))
	if v647 == int32(-1) {
		goto L5
	} else {
		goto L197
	}
L197:
	;
	v650 = int32(0)
	v651 = F_MultiXactIdCreate(m, l0, v630, l3, v647)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L12
	} else {
		goto L198
	}
L198:
	;
	v656 = F_GetMultiXactIdMembers(m, v651, v19+int32(76), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L12
	} else {
		goto L199
	}
L199:
	;
	if v656 <= int32(0) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1064 = v650
	v1066 = v651
	v1068 = int32(4240)
	goto L1
L201:
	;
	goto L202
L202:
	;
	v661 = int32(1)
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v656 == v661 {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	if v656&v661 == int32(0) {
		v771 = v734
		v773 = v738
		v774 = v746
		goto L224
	} else {
		goto L225
	}
L204:
	;
	v666 = int32(0)
	v734 = v650
	v737 = v666
	v738 = v666
	v746 = v10
	goto L203
L205:
	;
	goto L206
L206:
	;
	v670 = int32(0)
	v673 = v670
	v674 = v650
	v677 = v670
	v678 = v670
	v686 = v10
	goto L207
L207:
	;
	v691 = v663 + v677<<(uint(int32(3))%32)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v692<<(uint(int32(2))%32))+uint32(_consts[96])))
	if base.Ui32(v678) < base.Ui32(v697) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v734 = v725
	v737 = v729
	v738 = v727
	v746 = v726
	goto L203
L209:
	;
	v699 = v697
	goto L211
L210:
	;
	v699 = v678
	goto L211
L211:
	;
	switch v692 - int32(3) {
	case 0:
		goto L215
	case 1:
		v706 = v674
		goto L213
	case 2:
		goto L214
	default:
		v708 = v674
		v709 = v686
		goto L212
	}
L212:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v691)+12))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v710<<(uint(int32(2))%32))+uint32(_consts[96])))
	switch v710 - int32(3) {
	case 0:
		goto L219
	case 1:
		v723 = v708
		goto L217
	case 2:
		goto L218
	default:
		v725 = v708
		v726 = v709
		goto L216
	}
L213:
	;
	v708 = v706
	v709 = int32(1)
	goto L212
L214:
	;
	v706 = v674 | int32(8192)
	goto L213
L215:
	;
	v708 = v674 | int32(8192)
	v709 = v686
	goto L212
L216:
	;
	if base.Ui32(v699) < base.Ui32(v715) {
		goto L220
	} else {
		goto L221
	}
L217:
	;
	v725 = v723
	v726 = int32(1)
	goto L216
L218:
	;
	v723 = v708 | int32(8192)
	goto L217
L219:
	;
	v725 = v708 | int32(8192)
	v726 = v709
	goto L216
L220:
	;
	v727 = v715
	goto L222
L221:
	;
	v727 = v699
	goto L222
L222:
	;
	v728 = int32(2)
	v729 = v677 + v728
	v731 = v673 + v728
	if v731 != v656&int32(2147483646) {
		v673 = v731
		v674 = v725
		v677 = v729
		v678 = v727
		v686 = v726
		goto L207
	} else {
		goto L223
	}
L223:
	;
	goto L208
L224:
	;
	F_pfree(m, v663)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L12
	} else {
		goto L232
	}
L225:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v663+v737<<(uint(int32(3))%32))+4))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v754<<(uint(int32(2))%32))+uint32(_consts[96])))
	if base.Ui32(v738) < base.Ui32(v759) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v761 = v759
	goto L228
L227:
	;
	v761 = v738
	goto L228
L228:
	;
	switch v754 - int32(3) {
	case 0:
		goto L231
	case 1:
		v768 = v734
		goto L229
	case 2:
		goto L230
	default:
		v771 = v734
		v773 = v761
		v774 = v746
		goto L224
	}
L229:
	;
	v771 = v768
	v773 = v761
	v774 = int32(1)
	goto L224
L230:
	;
	v768 = v734 | int32(8192)
	goto L229
L231:
	;
	v771 = v734 | int32(8192)
	v773 = v761
	v774 = v746
	goto L224
L232:
	;
	if v773&int32(-2) == int32(2) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if v774 != 0 {
		v1064 = v771
		v1066 = v651
		v1068 = int32(4160)
		goto L1
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	if v773 != 0 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v1064 = v771
	v1066 = v651
	v1068 = int32(4288)
	goto L1
L237:
	;
	v786 = int32(4096)
	goto L239
L238:
	;
	v786 = int32(4112)
	goto L239
L239:
	;
	if v773 == int32(1) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v789 = int32(4176)
	goto L242
L241:
	;
	v789 = v786
	goto L242
L242:
	;
	if v774 != 0 {
		v1064 = v771
		v1066 = v651
		v1068 = v789
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v1064 = v771
	v1066 = v651
	v1068 = v789 | int32(128)
	goto L1
L244:
	;
	v792 = F_TransactionIdDidCommit(m, l0)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L12
	} else {
		goto L245
	}
L245:
	;
	if v792 == int32(0) {
		v1027 = l4
		goto L2
	} else {
		goto L246
	}
L246:
	;
	if l5 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v800 = int32(8)
	goto L249
L248:
	;
	v800 = int32(4)
	goto L249
L249:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v800)+uint32(_consts[95])))
	if v804 == int32(-1) {
		goto L4
	} else {
		goto L250
	}
L250:
	;
	v807 = int32(0)
	v808 = F_MultiXactIdCreate(m, l0, v433, l3, v804)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L12
	} else {
		goto L251
	}
L251:
	;
	v813 = F_GetMultiXactIdMembers(m, v808, v19+int32(76), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L12
	} else {
		goto L252
	}
L252:
	;
	if v813 <= int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1064 = v807
	v1066 = v808
	v1068 = int32(4240)
	goto L1
L254:
	;
	goto L255
L255:
	;
	v818 = int32(1)
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v813 == v818 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if v813&v818 == int32(0) {
		v928 = v891
		v930 = v895
		v931 = v903
		goto L277
	} else {
		goto L278
	}
L257:
	;
	v823 = int32(0)
	v891 = v807
	v894 = v823
	v895 = v823
	v903 = v10
	goto L256
L258:
	;
	goto L259
L259:
	;
	v827 = int32(0)
	v830 = v827
	v831 = v807
	v834 = v827
	v835 = v827
	v843 = v10
	goto L260
L260:
	;
	v848 = v820 + v834<<(uint(int32(3))%32)
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v848)+4))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v849<<(uint(int32(2))%32))+uint32(_consts[96])))
	if base.Ui32(v835) < base.Ui32(v854) {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v891 = v882
	v894 = v886
	v895 = v884
	v903 = v883
	goto L256
L262:
	;
	v856 = v854
	goto L264
L263:
	;
	v856 = v835
	goto L264
L264:
	;
	switch v849 - int32(3) {
	case 0:
		goto L268
	case 1:
		v863 = v831
		goto L266
	case 2:
		goto L267
	default:
		v865 = v831
		v866 = v843
		goto L265
	}
L265:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v848)+12))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v867<<(uint(int32(2))%32))+uint32(_consts[96])))
	switch v867 - int32(3) {
	case 0:
		goto L272
	case 1:
		v880 = v865
		goto L270
	case 2:
		goto L271
	default:
		v882 = v865
		v883 = v866
		goto L269
	}
L266:
	;
	v865 = v863
	v866 = int32(1)
	goto L265
L267:
	;
	v863 = v831 | int32(8192)
	goto L266
L268:
	;
	v865 = v831 | int32(8192)
	v866 = v843
	goto L265
L269:
	;
	if base.Ui32(v856) < base.Ui32(v872) {
		goto L273
	} else {
		goto L274
	}
L270:
	;
	v882 = v880
	v883 = int32(1)
	goto L269
L271:
	;
	v880 = v865 | int32(8192)
	goto L270
L272:
	;
	v882 = v865 | int32(8192)
	v883 = v866
	goto L269
L273:
	;
	v884 = v872
	goto L275
L274:
	;
	v884 = v856
	goto L275
L275:
	;
	v885 = int32(2)
	v886 = v834 + v885
	v888 = v830 + v885
	if v888 != v813&int32(2147483646) {
		v830 = v888
		v831 = v882
		v834 = v886
		v835 = v884
		v843 = v883
		goto L260
	} else {
		goto L276
	}
L276:
	;
	goto L261
L277:
	;
	F_pfree(m, v820)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L12
	} else {
		goto L285
	}
L278:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v820+v894<<(uint(int32(3))%32))+4))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v911<<(uint(int32(2))%32))+uint32(_consts[96])))
	if base.Ui32(v895) < base.Ui32(v916) {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v918 = v916
	goto L281
L280:
	;
	v918 = v895
	goto L281
L281:
	;
	switch v911 - int32(3) {
	case 0:
		goto L284
	case 1:
		v925 = v891
		goto L282
	case 2:
		goto L283
	default:
		v928 = v891
		v930 = v918
		v931 = v903
		goto L277
	}
L282:
	;
	v928 = v925
	v930 = v918
	v931 = int32(1)
	goto L277
L283:
	;
	v925 = v891 | int32(8192)
	goto L282
L284:
	;
	v928 = v891 | int32(8192)
	v930 = v918
	v931 = v903
	goto L277
L285:
	;
	if v930&int32(-2) == int32(2) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	if v931 != 0 {
		v1064 = v928
		v1066 = v808
		v1068 = int32(4160)
		goto L1
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	if v930 != 0 {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	v1064 = v928
	v1066 = v808
	v1068 = int32(4288)
	goto L1
L290:
	;
	v943 = int32(4096)
	goto L292
L291:
	;
	v943 = int32(4112)
	goto L292
L292:
	;
	if v930 == int32(1) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v946 = int32(4176)
	goto L295
L294:
	;
	v946 = v943
	goto L295
L295:
	;
	if v931 != 0 {
		v1064 = v928
		v1066 = v808
		v1068 = v946
		goto L1
	} else {
		goto L296
	}
L296:
	;
	v1064 = v928
	v1066 = v808
	v1068 = v946 | int32(128)
	goto L1
L297:
	;
	if l5 != 0 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v955 = int32(361051)
	goto L300
L299:
	;
	v955 = int32(378570)
	goto L300
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = l4
	F_errmsg_internal(m, int32(187369), v19-int32(-64))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L12
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(520720), int32(4538), int32(331587))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L12
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
	if l5 != 0 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v974 = int32(361051)
	goto L306
L305:
	;
	v974 = int32(378570)
	goto L306
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v974
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = l4
	F_errmsg_internal(m, int32(187369), v19+int32(48))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L12
	} else {
		goto L307
	}
L307:
	;
	F_errfinish(m, int32(520720), int32(4538), int32(331587))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L12
	} else {
		goto L308
	}
L308:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L309:
	;
	if l5 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v993 = int32(361051)
	goto L312
L311:
	;
	v993 = int32(378570)
	goto L312
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l4
	F_errmsg_internal(m, int32(187369), v19)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L12
	} else {
		goto L313
	}
L313:
	;
	F_errfinish(m, int32(520720), int32(4538), int32(331587))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L12
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
	if l5 != 0 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1010 = int32(361051)
	goto L318
L317:
	;
	v1010 = int32(378570)
	goto L318
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v1010
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = l4
	F_errmsg_internal(m, int32(187369), v19+int32(32))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L12
	} else {
		goto L319
	}
L319:
	;
	F_errfinish(m, int32(520720), int32(4538), int32(331587))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L12
	} else {
		goto L320
	}
L320:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L321:
	;
	v1064 = base.B2i32(v1027 == int32(3)) << (uint(int32(13)) % 32)
	v1066 = l3
	v1068 = int32(0)
	goto L1
L322:
	;
	goto L323
L323:
	;
	v1044 = int32(0)
	switch v1027 {
	case 0:
		v1064 = v1044
		v1066 = l3
		v1068 = int32(144)
		goto L1
	case 1:
		goto L327
	case 2:
		goto L326
	case 3:
		goto L325
	default:
		goto L324
	}
L324:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L12
	} else {
		goto L328
	}
L325:
	;
	v1064 = int32(8192)
	v1066 = l3
	v1068 = int32(192)
	goto L1
L326:
	;
	v1064 = v1044
	v1066 = l3
	v1068 = int32(192)
	goto L1
L327:
	;
	v1064 = v1044
	v1066 = l3
	v1068 = int32(208)
	goto L1
L328:
	;
	F_errmsg_internal(m, int32(431411), int32(0))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L12
	} else {
		goto L329
	}
L329:
	;
	F_errfinish(m, int32(520720), int32(5381), int32(329578))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L12
	} else {
		goto L330
	}
L330:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_concat_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
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
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v14 == v4 {
		v26 = v4
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L12
	} else {
		goto L51
	}
L2:
	;
	m.G0 = v12 + int32(32)
	return v181
L3:
	;
	if v26&int32(1) != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L3
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v18 == int32(0) {
		v26 = v4
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v21 != int32(15) {
		v26 = v4
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+13)))
	v26 = v24
	goto L4
L8:
	;
	v31 = l2 + l1<<(uint(int32(3))%32)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+24)))
	if v32 != 0 {
		v181 = v4
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_initStringInfo(m, v12+int32(8))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L15
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	v34 = F_pg_detoast_datum(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v39 = F_array_to_text_internal(m, l2, v34, l0, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v181 = v39
	goto L2
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v46 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+18)))
	v53 = F_MemoryContextAlloc(m, v49, v50*int32(28))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	v106 = v46
	goto L18
L18:
	;
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+18)))
	if l1 < v109 {
		goto L30
	} else {
		goto L31
	}
L19:
	;
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+18)))
	if l1 < v55 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v60 = l1
	goto L23
L21:
	;
	goto L22
L22:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+16)) = v53
	v106 = v53
	goto L18
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v67 = F_get_fn_expr_argtype(m, v66, v60)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	if v67 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_getTypeOutputInfo(m, v67, v12+int32(28), v12+int32(27))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	F_fmgr_info_cxt(m, v77, v53+v60*int32(28), v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v86 = v60 + int32(1)
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+18)))
	if v86 < v87 {
		v60 = v86
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	v115 = l1
	v117 = v109
	v121 = int32(1)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v165 = v163 + int32(4)
	v166 = F_palloc(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L12
	} else {
		goto L45
	}
L33:
	;
	v125 = l2 + int32(20) + v115<<(uint(int32(3))%32)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+4)))
	if v126 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v121 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v147 = v117
	v148 = v121
	goto L37
L37:
	;
	v150 = v115 + int32(1)
	if v150 < base.I32_extend16_s(v147) {
		v115 = v150
		v117 = v147
		v121 = v148
		goto L33
	} else {
		goto L44
	}
L38:
	;
	F_appendStringInfoString(m, v12+int32(8), l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L12
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v141 = F_OutputFunctionCall(m, v106+v115*int32(28), v129)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L12
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	F_appendStringInfoString(m, v12+int32(8), v141)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)))
	v147 = v146
	v148 = int32(0)
	goto L37
L44:
	;
	goto L34
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v165 << (uint(int32(2)) % 32)
	if v163 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	F_pfree(m, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L12
	} else {
		goto L50
	}
L47:
	;
	v173 = F__emscripten_memcpy_bulkmem(m, v166+int32(4), v162, v163)
	mBase = m.M
	goto L49
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	v181 = v166
	goto L2
L51:
	;
	F_errmsg_internal(m, int32(70917), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(524313), int32(5662), int32(417088))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_convert_saop_to_hashed_saop_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == v3 {
		v94 = v3
		m.G0 = v9 + int32(16)
		return v94
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v13 != int32(20) {
			v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v94 = v90
				m.G0 = v9 + int32(16)
				return v94
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			if v18 == int32(0) {
				v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					v94 = v90
					m.G0 = v9 + int32(16)
					return v94
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				if v21 != int32(7) {
					v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v94 = v90
						m.G0 = v9 + int32(16)
						return v94
					}
				} else {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
					if v24 != 0 {
						v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v94 = v90
							m.G0 = v9 + int32(16)
							return v94
						}
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v26 == int32(1) {
							v33 = F_get_op_hash_functions(m, v25, v9+int32(12), v9+int32(8))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								if v33 == int32(0) {
									v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										v94 = v90
										m.G0 = v9 + int32(16)
										return v94
									}
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
									if v39 != v40 {
										v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											v94 = v90
											m.G0 = v9 + int32(16)
											return v94
										}
									} else {
										v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
										v46 = F_ArrayGetNItems(m, v43, v42+int32(16))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int32(0)
										} else {
											if v46 < int32(9) {
												v94 = v3
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												v82 = int32(12)
												v83 = v51
												*(*int32)(unsafe.Add(mBase, uint32(l0+v82))) = v83
												v94 = v3
											}
											m.G0 = v9 + int32(16)
											return v94
										}
									}
								}
							}
						} else {
							v52 = F_get_negator(m, v25)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if v52 == int32(0) {
									v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										v94 = v90
										m.G0 = v9 + int32(16)
										return v94
									}
								} else {
									v60 = F_get_op_hash_functions(m, v52, v9+int32(12), v9+int32(8))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										if v60 == int32(0) {
											v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int32(0)
											} else {
												v94 = v90
												m.G0 = v9 + int32(16)
												return v94
											}
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
											if v64 != v65 {
												v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int32(0)
												} else {
													v94 = v90
													m.G0 = v9 + int32(16)
													return v94
												}
											} else {
												v67 = int32(16)
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
												v72 = F_ArrayGetNItems(m, v69, v68+v67)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													if v72 < int32(9) {
														v94 = v3
														m.G0 = v9 + int32(16)
														return v94
													} else {
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v76
														v78 = F_get_opcode(m, v52)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															v82 = v67
															v83 = v78
															*(*int32)(unsafe.Add(mBase, uint32(l0+v82))) = v83
															v94 = v3
															m.G0 = v9 + int32(16)
															return v94
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_convert_testexpr_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == int32(0) {
		v40 = int32(0)
		m.G0 = v7 + int32(16)
		return v40
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v12 - int32(8) {
		case 0:
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v15 != int32(2) {
				v38 = F_expression_tree_mutator_impl(m, l0, int32(845), l1)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = v38
					m.G0 = v7 + int32(16)
					return v40
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v18 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v52
						F_errmsg_internal(m, int32(509575), v7)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(515808), int32(669), int32(219645))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v21 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v52
							F_errmsg_internal(m, int32(509575), v7)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(515808), int32(669), int32(219645))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
						if v24 < v18 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v52
								F_errmsg_internal(m, int32(509575), v7)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(515808), int32(669), int32(219645))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+v18<<(uint(int32(2))%32)-int32(4))))
							v33 = F_copyObjectImpl(m, v32)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v40 = v33
								m.G0 = v7 + int32(16)
								return v40
							}
						}
					}
				}
			}
		default:
			v38 = F_expression_tree_mutator_impl(m, l0, int32(845), l1)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = v38
				m.G0 = v7 + int32(16)
				return v40
			}
		case 14:
			v40 = l0
			m.G0 = v7 + int32(16)
			return v40
		}
	}
}
func F_core_yyensure_buffer_stack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		v8 = F_palloc(m, int32(4))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
			if v8 == int32(0) {
				F_yy_fatal_error_2(m, int32(714054))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(4294967296)
				return
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(v18-int32(1)) <= base.Ui32(v17) {
			v23 = v18 + int32(8)
			v26 = F_repalloc(m, v4, v23<<(uint(int32(2))%32))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v26
				if v26 == int32(0) {
					F_yy_fatal_error_2(m, int32(714054))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = v26 + v31<<(uint(int32(2))%32)
					v35 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					return
				}
			}
		} else {
			return
		}
	}
}
func F_cost_append(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 float64
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 float64
	_ = v180
	var v181 float64
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 float64
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 float64
	_ = v193
	var v198 float64
	_ = v198
	var v201 float64
	_ = v201
	var v203 float64
	_ = v203
	var v206 float64
	_ = v206
	var v213 float64
	_ = v213
	var v215 float64
	_ = v215
	var v219 int32
	_ = v219
	var v220 float64
	_ = v220
	var v223 int64
	_ = v223
	var v224 float64
	_ = v224
	var v226 float64
	_ = v226
	var v227 int32
	_ = v227
	var v230 float64
	_ = v230
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	var v238 float64
	_ = v238
	var v240 float64
	_ = v240
	var v241 float64
	_ = v241
	var v244 float64
	_ = v244
	var v247 float64
	_ = v247
	var v251 float64
	_ = v251
	var v258 float64
	_ = v258
	var v259 float64
	_ = v259
	var v262 float64
	_ = v262
	var v266 float64
	_ = v266
	var v276 float64
	_ = v276
	var v279 float64
	_ = v279
	var v283 int32
	_ = v283
	var v288 float64
	_ = v288
	var v289 float64
	_ = v289
	var v291 float64
	_ = v291
	var v296 int32
	_ = v296
	var v299 float64
	_ = v299
	var v300 float64
	_ = v300
	var v301 float64
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 float64
	_ = v307
	var v308 float64
	_ = v308
	var v311 float64
	_ = v311
	var v312 float64
	_ = v312
	var v313 float64
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 float64
	_ = v320
	var v322 int32
	_ = v322
	var v328 float64
	_ = v328
	var v332 float64
	_ = v332
	var v335 float64
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 float64
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 float64
	_ = v349
	var v350 float64
	_ = v350
	var v354 float64
	_ = v354
	var v358 float64
	_ = v358
	var v361 float64
	_ = v361
	var v364 float64
	_ = v364
	var v365 float64
	_ = v365
	var v367 float64
	_ = v367
	var v369 float64
	_ = v369
	var v371 float64
	_ = v371
	var v374 float64
	_ = v374
	var v376 float64
	_ = v376
	var v378 float64
	_ = v378
	var v379 int32
	_ = v379
	var v381 float64
	_ = v381
	var v389 float64
	_ = v389
	var v393 float64
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v418 float64
	_ = v418
	var v420 float64
	_ = v420
	var v421 float64
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v429 float64
	_ = v429
	var v431 float64
	_ = v431
	var v434 float64
	_ = v434
	var v436 float64
	_ = v436
	var v438 float64
	_ = v438
	var v440 int32
	_ = v440
	var v441 float64
	_ = v441
	var v442 float64
	_ = v442
	var v446 float64
	_ = v446
	var v450 float64
	_ = v450
	var v453 float64
	_ = v453
	var v456 float64
	_ = v456
	var v458 float64
	_ = v458
	var v459 float64
	_ = v459
	var v461 float64
	_ = v461
	var v462 float64
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 float64
	_ = v467
	var v475 float64
	_ = v475
	var v479 float64
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v495 int32
	_ = v495
	var v504 float64
	_ = v504
	var v506 float64
	_ = v506
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v551 int32
	_ = v551
	var v552 float64
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v630 int32
	_ = v630
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 float64
	_ = v660
	var v661 float64
	_ = v661
	var v667 int32
	_ = v667
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v706 float64
	_ = v706
	var v710 float64
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 float64
	_ = v716
	var v720 float64
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 float64
	_ = v726
	var v730 float64
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 float64
	_ = v736
	var v740 float64
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v798 int32
	_ = v798
	var v801 float64
	_ = v801
	var v805 float64
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v939 float64
	_ = v939
	var v943 float64
	_ = v943
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 float64
	_ = v949
	var v953 float64
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 float64
	_ = v959
	var v963 float64
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 float64
	_ = v969
	var v973 float64
	_ = v973
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1031 int32
	_ = v1031
	var v1034 float64
	_ = v1034
	var v1038 float64
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1070 float64
	_ = v1070
	var v1074 float64
	_ = v1074
	var v1075 float64
	_ = v1075
	var v1095 float64
	_ = v1095
	var v1096 float64
	_ = v1096
	var v1102 float64
	_ = v1102
	v2 = int32(0)
	v19 = float64(0)
	v25 = m.G0
	v27 = v25 - int32(96)
	m.G0 = v27
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v29
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v39 = l0 + int32(48)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v40 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v27 + int32(96)
	return
L4:
	;
	v1102 = *(*float64)(unsafe.Add(mBase, _consts[611]))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(base.F64_mul(base.F64_mul(v1102, float64(0.5)), v1096), v1095)
	goto L3
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v43 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v320 = base.F64_convert_i32_s(v319)
	v322 = int32(*(*uint8)(unsafe.Add(mBase, _consts[635])))
	if v322 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L8:
	;
	v99 = v2
	goto L19
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if int32(0) < v44 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v49 = *(*float64)(unsafe.Add(mBase, uint32(v48)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v39))) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v51 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v1095 = v19
	v1096 = v19
	goto L4
L13:
	;
	v1095 = v19
	v1096 = v19
	goto L4
L14:
	;
	goto L15
L15:
	;
	v56 = v2
	v58 = int32(0)
	v73 = v19
	v74 = v19
	goto L16
L16:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v56<<(uint(int32(2))%32))))
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v83)+32))
	v85 = base.F64_add(v84, v74)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
	v88 = v58 + v87
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v88
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v83)+56))
	v91 = base.F64_add(v90, v73)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v91
	v94 = v56 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v94 < v95 {
		v56 = v94
		v58 = v88
		v73 = v91
		v74 = v85
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v1095 = v91
	v1096 = v85
	goto L4
L18:
	;
	goto L17
L19:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v99<<(uint(int32(2))%32))))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+64))
	if v43 == v126 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v1095 = v313
	v1096 = v301
	goto L4
L21:
	;
	v300 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v301 = base.F64_add(v299, v300)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v301
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v296)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v303 + v304
	v307 = *(*float64)(unsafe.Add(mBase, uint32(v296)+48))
	v308 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = base.F64_add(v307, v308)
	v311 = *(*float64)(unsafe.Add(mBase, uint32(v296)+56))
	v312 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
	v313 = base.F64_add(v311, v312)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v313
	v316 = v99 + int32(1)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v316 < v317 {
		v99 = v316
		goto L19
	} else {
		goto L63
	}
L22:
	;
	if v179 != 0 {
		goto L40
	} else {
		goto L41
	}
L23:
	;
	v179 = int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v135 = int32(0)
	goto L27
L26:
	;
	v179 = v171
	goto L22
L27:
	;
	v139 = int32(0)
	if v43 == v139 {
		v149 = v139
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v171 = int32(0)
	goto L26
L29:
	;
	if v126 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v143 <= v135 {
		v149 = int32(0)
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v149 = v145 + v135<<(uint(int32(2))%32)
	goto L29
L32:
	;
	v155 = base.B2i32(v149 == int32(0))
	if v149 == int32(0) {
		v171 = v155
		goto L26
	} else {
		goto L37
	}
L33:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v135 < v150 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v179 = base.B2i32(v149 == int32(0))
	goto L22
L36:
	;
	goto L35
L37:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v161 = v158 + v135<<(uint(int32(2))%32)
	if v161 == int32(0) {
		v171 = v155
		goto L26
	} else {
		goto L38
	}
L38:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v166 == v167 {
		v135 = v135 + int32(1)
		goto L27
	} else {
		goto L39
	}
L39:
	;
	goto L28
L40:
	;
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v125)+32))
	v296 = v125
	v299 = v180
	goto L21
L41:
	;
	goto L42
L42:
	;
	v181 = *(*float64)(unsafe.Add(mBase, uint32(v125)+56))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v125)+40))
	v184 = v27 + int32(88)
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v125)+32))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+32))
	v192 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	v193 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
	v198 = float64(2)
	if base.F64_lt(v187, v198) != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, _consts[636])))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v182 + (v283 ^ int32(1))
	v288 = *(*float64)(unsafe.Add(mBase, uint32(v27)+88))
	v289 = base.F64_add(v181, v288)
	*(*float64)(unsafe.Add(mBase, uint32(v27)+56)) = v289
	v291 = *(*float64)(unsafe.Add(mBase, uint32(v27)+80))
	*(*float64)(unsafe.Add(mBase, uint32(v27)+64)) = base.F64_add(v289, v291)
	v296 = v27 + int32(8)
	v299 = v187
	goto L21
L44:
	;
	v201 = v198
	goto L46
L45:
	;
	v201 = v187
	goto L46
L46:
	;
	v203 = *(*float64)(unsafe.Add(mBase, _consts[613]))
	v206 = base.F64_mul(v201, base.F64_add(base.F64_add(v203, v203), float64(0)))
	v213 = base.F64_convert_i32_u((v189+int32(7))&int32(-8) + int32(24))
	v215 = base.F64_mul(v187, v213)
	v219 = base.F64_lt(v193, v201) & base.F64_gt(v193, float64(0))
	if v219 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v184))) = v276
	v279 = *(*float64)(unsafe.Add(mBase, _consts[613]))
	*(*float64)(unsafe.Add(mBase, uint32(v27+int32(80)))) = base.F64_mul(v201, v279)
	goto L43
L48:
	;
	v220 = base.F64_mul(v193, v213)
	goto L50
L49:
	;
	v220 = v215
	goto L50
L50:
	;
	v223 = base.I64_extend_i32_s(v192) << (uint(int64(10)) % 64)
	v224 = base.F64_convert_i64_s(v223)
	if base.F64_gt(v220, v224) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v226 = F_log(m, v201)
	mBase = m.M
	v227 = F_tuplesort_merge_order(m, v223)
	mBase = m.M
	v230 = base.F64_mul(base.F64_div(v226, float64(0.693147180559945)), v206)
	*(*float64)(unsafe.Add(mBase, uint32(v184))) = v230
	v235 = base.F64_ceil(base.F64_mul(v215, float64(0.0001220703125)))
	v237 = base.F64_div(v215, v224)
	v238 = base.F64_convert_i32_s(v227)
	if base.F64_gt(v237, v238) != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	if v219 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v240 = F_log(m, v237)
	mBase = m.M
	v241 = F_log(m, v238)
	mBase = m.M
	v244 = base.F64_ceil(base.F64_div(v240, v241))
	goto L56
L55:
	;
	v244 = float64(1)
	goto L56
L56:
	;
	v247 = *(*float64)(unsafe.Add(mBase, _consts[637]))
	v251 = *(*float64)(unsafe.Add(mBase, _consts[638]))
	v276 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v235, v235), v244), base.F64_add(base.F64_mul(v247, float64(0.75)), base.F64_mul(v251, float64(0.25)))), v230)
	goto L47
L57:
	;
	v258 = v193
	goto L59
L58:
	;
	v258 = v201
	goto L59
L59:
	;
	v259 = base.F64_add(v258, v258)
	if base.F64_gt(v215, v224)|base.F64_gt(v201, v259) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v262 = F_log(m, v259)
	mBase = m.M
	v276 = base.F64_mul(base.F64_div(v262, float64(0.693147180559945)), v206)
	goto L47
L61:
	;
	goto L62
L62:
	;
	v266 = F_log(m, v201)
	mBase = m.M
	v276 = base.F64_mul(base.F64_div(v266, float64(0.693147180559945)), v206)
	goto L47
L63:
	;
	goto L20
L64:
	;
	v328 = base.F64_add(base.F64_mul(v320, float64(-0.3)), float64(1))
	if base.F64_gt(v328, float64(0)) != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v335 = v320
	goto L66
L66:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v336 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	v332 = v328
	goto L69
L68:
	;
	v332 = math.Float64frombits(uint64(0x8000000000000000))
	goto L69
L69:
	;
	v335 = base.F64_add(v332, v320)
	goto L66
L70:
	;
	if v495 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L71:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v495 = v339
	v504 = float64(0)
	v506 = v19
	goto L70
L72:
	;
	goto L73
L73:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v342)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v343
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v345 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v342)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v379
	v381 = float64(1e+100)
	if base.F64_gt(v376, v381) != 0 {
		v393 = v381
		goto L84
	} else {
		goto L85
	}
L75:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v342)+24))
	v349 = base.F64_convert_i32_s(v348)
	v350 = *(*float64)(unsafe.Add(mBase, uint32(v342)+32))
	if v322 != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	v371 = *(*float64)(unsafe.Add(mBase, uint32(v342)+32))
	v374 = base.F64_add(base.F64_div(v371, v335), float64(0))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v374
	v376 = v374
	v378 = v19
	goto L74
L78:
	;
	v354 = base.F64_add(base.F64_mul(v349, float64(-0.3)), float64(1))
	if base.F64_gt(v354, float64(0)) != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v361 = v349
	goto L80
L80:
	;
	v364 = float64(0)
	v365 = base.F64_add(base.F64_mul(v350, base.F64_div(v361, v335)), v364)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v365
	v367 = *(*float64)(unsafe.Add(mBase, uint32(v342)+56))
	v369 = base.F64_add(v367, v364)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v369
	v376 = v365
	v378 = v369
	goto L74
L81:
	;
	v358 = v354
	goto L83
L82:
	;
	v358 = math.Float64frombits(uint64(0x8000000000000000))
	goto L83
L83:
	;
	v361 = base.F64_add(v358, v349)
	goto L80
L84:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v393
	v395 = int32(1)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v396 <= v395 {
		v495 = v345
		v504 = v393
		v506 = v378
		goto L70
	} else {
		goto L88
	}
L85:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v376)&int64(9223372036854775807)) {
		v393 = v381
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v389 = float64(1)
	if base.F64_le(v376, v389) != 0 {
		v393 = v389
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v393 = base.F64_nearest(v376)
	goto L84
L88:
	;
	v400 = v395
	v402 = v379
	v418 = v393
	v420 = v378
	v421 = v343
	goto L89
L89:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v423+v400<<(uint(int32(2))%32))))
	if v400 < v319 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v495 = v345
	v504 = v479
	v506 = v462
	goto L70
L91:
	;
	v429 = *(*float64)(unsafe.Add(mBase, uint32(v427)+48))
	if base.F64_gt(v429, v421) != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v434 = v421
	goto L93
L93:
	;
	if v400 < v345 {
		goto L98
	} else {
		goto L99
	}
L94:
	;
	v431 = v421
	goto L96
L95:
	;
	v431 = v429
	goto L96
L96:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v39))) = v431
	v434 = v431
	goto L93
L97:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v427)+40))
	v465 = v402 + v464
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v465
	v467 = float64(1e+100)
	if base.F64_gt(v461, v467) != 0 {
		v479 = v467
		goto L107
	} else {
		goto L108
	}
L98:
	;
	v436 = *(*float64)(unsafe.Add(mBase, uint32(v427)+32))
	v438 = base.F64_add(v418, base.F64_div(v436, v335))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v438
	v461 = v438
	v462 = v420
	goto L97
L99:
	;
	goto L100
L100:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v427)+24))
	v441 = base.F64_convert_i32_s(v440)
	v442 = *(*float64)(unsafe.Add(mBase, uint32(v427)+32))
	if v322 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v446 = base.F64_add(base.F64_mul(v441, float64(-0.3)), float64(1))
	if base.F64_gt(v446, float64(0)) != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v453 = v441
	goto L103
L103:
	;
	v456 = base.F64_add(base.F64_mul(v442, base.F64_div(v453, v335)), v418)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v456
	v458 = *(*float64)(unsafe.Add(mBase, uint32(v427)+56))
	v459 = base.F64_add(v458, v420)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v459
	v461 = v456
	v462 = v459
	goto L97
L104:
	;
	v450 = v446
	goto L106
L105:
	;
	v450 = math.Float64frombits(uint64(0x8000000000000000))
	goto L106
L106:
	;
	v453 = base.F64_add(v450, v441)
	goto L103
L107:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v479
	v482 = v400 + int32(1)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v482 < v483 {
		v400 = v482
		v402 = v465
		v418 = v479
		v420 = v462
		v421 = v434
		goto L89
	} else {
		goto L111
	}
L108:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v461)&int64(9223372036854775807)) {
		v479 = v467
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v475 = float64(1)
	if base.F64_le(v461, v475) != 0 {
		v479 = v475
		goto L107
	} else {
		goto L110
	}
L110:
	;
	v479 = base.F64_nearest(v461)
	goto L107
L111:
	;
	goto L90
L112:
	;
	v1095 = base.F64_add(v506, float64(0))
	v1096 = v504
	goto L4
L113:
	;
	goto L114
L114:
	;
	if v319 < v495 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v514 = v319
	goto L117
L116:
	;
	v514 = v495
	goto L117
L117:
	;
	v517 = F_palloc(m, v514<<(uint(int32(3))%32))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	return
L119:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v519 <= int32(0) {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	v1070 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v1074 = *(*float64)(unsafe.Add(mBase, uint32(v517+v1049<<(uint(int32(3))%32))))
	v1075 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
	v1095 = base.F64_add(v1074, v1075)
	v1096 = v1070
	goto L4
L121:
	;
	v894 = int32(3)
	v895 = v514 & v894
	v896 = int32(0)
	if base.Ui32(v894) <= base.Ui32(v514-int32(1)) {
		goto L170
	} else {
		goto L171
	}
L122:
	;
	v614 = int32(2)
	v617 = v514 << (uint(v614) % 32) >> (uint(v614) % 32)
	if v568 <= v617 {
		goto L134
	} else {
		goto L135
	}
L123:
	;
	v611 = int32(0)
	if v514 <= v611 {
		v1049 = v611
		goto L120
	} else {
		goto L133
	}
L124:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v514 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v526 = int32(0)
	v527 = v522
	goto L128
L126:
	;
	v566 = v522
	v568 = v519
	goto L127
L127:
	;
	if v566 != 0 {
		goto L122
	} else {
		goto L132
	}
L128:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	v552 = *(*float64)(unsafe.Add(mBase, uint32(v551)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v517+v526<<(uint(int32(3))%32)))) = v552
	v555 = v526 + int32(1)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v556 <= v555 {
		goto L123
	} else {
		goto L130
	}
L129:
	;
	v566 = v561
	v568 = v556
	goto L127
L130:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v561 = v558 + v555<<(uint(int32(2))%32)
	if v555 != v514 {
		v526 = v555
		v527 = v561
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	goto L123
L133:
	;
	goto L121
L134:
	;
	if int32(0) < v514 {
		goto L121
	} else {
		goto L169
	}
L135:
	;
	v622 = v514 & int32(3)
	v630 = v514 - int32(1)
	v640 = v617
	v641 = v514
	goto L136
L136:
	;
	if v495 == v641 {
		goto L134
	} else {
		goto L138
	}
L137:
	;
	goto L134
L138:
	;
	v654 = v517 + v630<<(uint(int32(3))%32)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v655+v640<<(uint(int32(2))%32))))
	v660 = *(*float64)(unsafe.Add(mBase, uint32(v659)+56))
	v661 = *(*float64)(unsafe.Add(mBase, uint32(v654)))
	*(*float64)(unsafe.Add(mBase, uint32(v654))) = base.F64_add(v660, v661)
	if v514 <= int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v837 = int32(1)
	v840 = v640 + v837
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v840 < v841 {
		v630 = v816
		v640 = v840
		v641 = v641 + v837
		goto L136
	} else {
		goto L168
	}
L140:
	;
	v816 = int32(0)
	goto L139
L141:
	;
	goto L142
L142:
	;
	v667 = int32(0)
	if base.B2i32(base.Ui32(v514) < base.Ui32(int32(4))) == v667 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v675 = v667
	v676 = v667
	v681 = v667
	goto L146
L144:
	;
	v750 = v667
	v751 = v667
	goto L145
L145:
	;
	if v622 == int32(0) {
		v816 = v751
		goto L139
	} else {
		goto L161
	}
L146:
	;
	v697 = int32(3)
	v698 = v675 | v697
	v700 = v675 | int32(2)
	v702 = v675 | int32(1)
	v706 = *(*float64)(unsafe.Add(mBase, uint32(v517+v675<<(uint(v697)%32))))
	v710 = *(*float64)(unsafe.Add(mBase, uint32(v517+v676<<(uint(v697)%32))))
	if base.F64_lt(v706, v710) != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v750 = v744
	v751 = v742
	goto L145
L148:
	;
	v712 = v675
	goto L150
L149:
	;
	v712 = v676
	goto L150
L150:
	;
	v713 = int32(3)
	v716 = *(*float64)(unsafe.Add(mBase, uint32(v517+v702<<(uint(v713)%32))))
	v720 = *(*float64)(unsafe.Add(mBase, uint32(v517+v712<<(uint(v713)%32))))
	if base.F64_lt(v716, v720) != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v722 = v702
	goto L153
L152:
	;
	v722 = v712
	goto L153
L153:
	;
	v723 = int32(3)
	v726 = *(*float64)(unsafe.Add(mBase, uint32(v517+v700<<(uint(v723)%32))))
	v730 = *(*float64)(unsafe.Add(mBase, uint32(v517+v722<<(uint(v723)%32))))
	if base.F64_lt(v726, v730) != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v732 = v700
	goto L156
L155:
	;
	v732 = v722
	goto L156
L156:
	;
	v733 = int32(3)
	v736 = *(*float64)(unsafe.Add(mBase, uint32(v517+v698<<(uint(v733)%32))))
	v740 = *(*float64)(unsafe.Add(mBase, uint32(v517+v732<<(uint(v733)%32))))
	if base.F64_lt(v736, v740) != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v742 = v698
	goto L159
L158:
	;
	v742 = v732
	goto L159
L159:
	;
	v743 = int32(4)
	v744 = v675 + v743
	v746 = v681 + v743
	if v746 != v514&int32(2147483644) {
		v675 = v744
		v676 = v742
		v681 = v746
		goto L146
	} else {
		goto L160
	}
L160:
	;
	goto L147
L161:
	;
	v776 = v750
	v777 = v751
	v779 = v667
	goto L162
L162:
	;
	v798 = int32(3)
	v801 = *(*float64)(unsafe.Add(mBase, uint32(v517+v776<<(uint(v798)%32))))
	v805 = *(*float64)(unsafe.Add(mBase, uint32(v517+v777<<(uint(v798)%32))))
	if base.F64_lt(v801, v805) != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v816 = v807
	goto L139
L164:
	;
	v807 = v776
	goto L166
L165:
	;
	v807 = v777
	goto L166
L166:
	;
	v808 = int32(1)
	v811 = v779 + v808
	if v811 != v622 {
		v776 = v776 + v808
		v777 = v807
		v779 = v811
		goto L162
	} else {
		goto L167
	}
L167:
	;
	goto L163
L168:
	;
	goto L137
L169:
	;
	v1049 = int32(0)
	goto L120
L170:
	;
	v908 = v896
	v909 = v896
	v911 = int32(0)
	goto L173
L171:
	;
	v983 = v896
	v984 = v896
	goto L172
L172:
	;
	if v895 == int32(0) {
		v1049 = v984
		goto L120
	} else {
		goto L188
	}
L173:
	;
	v930 = int32(3)
	v931 = v908 | v930
	v933 = v908 | int32(2)
	v935 = v908 | int32(1)
	v939 = *(*float64)(unsafe.Add(mBase, uint32(v517+v908<<(uint(v930)%32))))
	v943 = *(*float64)(unsafe.Add(mBase, uint32(v517+v909<<(uint(v930)%32))))
	if base.F64_gt(v939, v943) != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v983 = v977
	v984 = v975
	goto L172
L175:
	;
	v945 = v908
	goto L177
L176:
	;
	v945 = v909
	goto L177
L177:
	;
	v946 = int32(3)
	v949 = *(*float64)(unsafe.Add(mBase, uint32(v517+v935<<(uint(v946)%32))))
	v953 = *(*float64)(unsafe.Add(mBase, uint32(v517+v945<<(uint(v946)%32))))
	if base.F64_gt(v949, v953) != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v955 = v935
	goto L180
L179:
	;
	v955 = v945
	goto L180
L180:
	;
	v956 = int32(3)
	v959 = *(*float64)(unsafe.Add(mBase, uint32(v517+v933<<(uint(v956)%32))))
	v963 = *(*float64)(unsafe.Add(mBase, uint32(v517+v955<<(uint(v956)%32))))
	if base.F64_gt(v959, v963) != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v965 = v933
	goto L183
L182:
	;
	v965 = v955
	goto L183
L183:
	;
	v966 = int32(3)
	v969 = *(*float64)(unsafe.Add(mBase, uint32(v517+v931<<(uint(v966)%32))))
	v973 = *(*float64)(unsafe.Add(mBase, uint32(v517+v965<<(uint(v966)%32))))
	if base.F64_gt(v969, v973) != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v975 = v931
	goto L186
L185:
	;
	v975 = v965
	goto L186
L186:
	;
	v976 = int32(4)
	v977 = v908 + v976
	v979 = v911 + v976
	if v979 != v514&int32(-4) {
		v908 = v977
		v909 = v975
		v911 = v979
		goto L173
	} else {
		goto L187
	}
L187:
	;
	goto L174
L188:
	;
	v1009 = v983
	v1010 = v984
	v1011 = v896
	goto L189
L189:
	;
	v1031 = int32(3)
	v1034 = *(*float64)(unsafe.Add(mBase, uint32(v517+v1009<<(uint(v1031)%32))))
	v1038 = *(*float64)(unsafe.Add(mBase, uint32(v517+v1010<<(uint(v1031)%32))))
	if base.F64_gt(v1034, v1038) != 0 {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v1049 = v1040
	goto L120
L191:
	;
	v1040 = v1009
	goto L193
L192:
	;
	v1040 = v1010
	goto L193
L193:
	;
	v1041 = int32(1)
	v1044 = v1011 + v1041
	if v1044 != v895 {
		v1009 = v1009 + v1041
		v1010 = v1040
		v1011 = v1044
		goto L189
	} else {
		goto L194
	}
L194:
	;
	goto L190
}
func F_countVariablesFromJsonb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	if l0 == int32(0) {
		return base.B2i32(l0 != int32(0))
	} else {
		v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
		if v4&int32(32) != 0 {
			return base.B2i32(l0 != int32(0))
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(118096), int32(0))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						F_errdetail(m, int32(609364), int32(0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(523842), int32(3211), int32(527085))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
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
	}
}
func F_countitem_compare_count(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	return base.B2i32(v7 < v5) - base.B2i32(v5 < v7)
}
func F_cr_circle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v15 float64
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_palloc(m, int32(24))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		*(*float64)(unsafe.Add(mBase, uint32(v9))) = v13
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v6
		*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v15
		return v9
	}
}
func F_createPostingTree(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int64
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int64
	_ = v240
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	v20 = F_palloc(m, int32(8192))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(131)
	F_PageInit(m, v20, int32(8192), int32(8))
	mBase = m.M
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v29 = v20 + v28
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+6)) = uint16(v24)
	goto L3
L3:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20+v33))) = int32(-1)
	if l2 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v97 = v91 + int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+12)) = uint16(v97)
	v99 = F_GinNewBuffer(m, l0)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L21
	}
L5:
	;
	v89 = v6
	v91 = v6
	goto L4
L6:
	;
	goto L7
L7:
	;
	v48 = v6
	v49 = v6
	v53 = v20 + int32(32)
	goto L8
L8:
	;
	v62 = F_ginCompressPostingList(m, l1+v48*int32(6), l2-v48, int32(384), v17+int32(28))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v89 = v80
	v91 = v71
	goto L4
L10:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
	v70 = (v64+int32(1))&int32(131070) + int32(8)
	v71 = v70 + v49
	if base.Ui32(int32(8153)) <= base.Ui32(v71) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v89 = v48
	v91 = v49
	goto L4
L12:
	;
	goto L13
L13:
	;
	if v70 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	F_pfree(m, v62)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L18
	}
L15:
	;
	v74 = F__emscripten_memcpy_bulkmem(m, v53, v62, v70)
	mBase = m.M
	v75 = v74
	goto L17
L16:
	;
	v75 = v53
	goto L17
L17:
	;
	goto L14
L18:
	;
	v80 = v48 + v76
	if base.Ui32(v80) < base.Ui32(l2) {
		v48 = v80
		v49 = v71
		v53 = v75 + v70
		goto L8
	} else {
		goto L19
	}
L19:
	;
	goto L9
L20:
	;
	if v99 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	if v99 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v104+(v99^int32(-1))<<(uint(int32(2))%32))))
	v118 = v110
	goto L20
L23:
	;
	goto L24
L24:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v118 = v112 + v99<<(uint(int32(13))%32) + int32(-8192)
	goto L20
L25:
	;
	if l4 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122+(v99^int32(-1))<<(uint(int32(6))%32))+16))
	v137 = v128
	goto L25
L27:
	;
	goto L28
L28:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130+v99<<(uint(int32(6))%32)+int32(-64))+16))
	v137 = v136
	goto L25
L29:
	;
	F_PredicateLockPageSplit(m, l0, v156, v137)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L33
	}
L30:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141+(l4^int32(-1))<<(uint(int32(6))%32))+16))
	v156 = v147
	goto L29
L31:
	;
	goto L32
L32:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v149+l4<<(uint(int32(6))%32)+int32(-64))+16))
	v156 = v155
	goto L29
L33:
	;
	v159 = int32(4548788)
	v161 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v161 + int32(1)
	F_PageRestoreTempPage(m, v20, v118)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_MarkBufferDirty(m, v99)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+118)))
	if v170 != int32(112) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v226 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L57
	}
L37:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v220 + int32(1)
	goto L36
L38:
	;
	F_UnlockReleaseBuffer(m, v99)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L56
	}
L39:
	;
	F_UnlockReleaseBuffer(m, v99)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L54
	}
L40:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v174 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v91
	F_XLogBeginInsert(m)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L49
	}
L42:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v177 != 0 {
		goto L39
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if l3 != 0 {
		goto L38
	} else {
		goto L48
	}
L45:
	;
	if l3 != 0 {
		goto L39
	} else {
		goto L46
	}
L46:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v178 == int32(0) {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	goto L39
L48:
	;
	goto L41
L49:
	;
	F_XLogRegisterData(m, v17+int32(28), int32(4))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_XLogRegisterData(m, v118+int32(32), v91)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_XLogRegisterBuffer(m, int32(0), v99, int32(6))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v199 = F_XLogInsert(m, int32(13), int32(16))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v118))) = base.I64_rotr(v199, int64(32))
	goto L39
L54:
	;
	v206 = int32(4548788)
	v208 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v208 - int32(1)
	if l3 != 0 {
		goto L37
	} else {
		goto L55
	}
L55:
	;
	goto L36
L56:
	;
	v214 = int32(4548788)
	v216 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v216 - int32(1)
	goto L37
L57:
	;
	if v226 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v89
	F_errmsg_internal(m, int32(160149), v17)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if base.Ui32(v89) < base.Ui32(l2) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	F_errfinish(m, int32(522896), int32(1865), int32(429692))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v239 = v17 + int32(72)
	v240 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v239))) = v240
	*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = v240
	*(*int64)(unsafe.Add(mBase, uint32(v17)+88)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v239))) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = int32(35)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = int32(37)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = int32(38)
	v257 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = int32(39)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = int32(42)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = l0
	v268 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+64)) = uint8(v268)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = l2 - v89
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+81)) = uint8(base.B2i32(l3 != v257))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l1 + v89*int32(6)
	v282 = v17 + int32(90)
	v290 = v257
	goto L66
L64:
	;
	goto L65
L65:
	;
	m.G0 = v17 + int32(96)
	return v137
L66:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v301 = v298 + v290*int32(6)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v302
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v301)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v282)+4)) = uint16(v304)
	v307 = v17 + int32(28)
	v312 = F_ginFindLeafPage(m, v307, int32(0), int32(1))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	goto L65
L68:
	;
	F_ginInsertValue(m, v307, v312, v17+int32(16), l3)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if base.Ui32(v318) < base.Ui32(v319) {
		v290 = v318
		goto L66
	} else {
		goto L70
	}
L70:
	;
	goto L67
}
func F_create_ctas_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = F_palloc0(m, int32(56))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(160)
		v18 = *(*int64)(unsafe.Add(mBase, _consts[421]))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(0)
		v24 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v21
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v28
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v30
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		*(*uint8)(unsafe.Add(mBase, uint32(v13)+52)) = uint8(v24)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v32
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v36
		if v20 != 0 {
			v40 = int32(109)
		} else {
			v40 = int32(114)
		}
		v41 = int32(0)
		F_DefineRelation(m, l0, v13, v40, v41, v41, v41)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			F_CommandCounterIncrement(m)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				v49 = int32(0)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
				v56 = F_transformRelOptions(m, v49, v50, int32(83935), v10+int32(8), int32(1), v49)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_heap_reloptions(m, int32(116), v56)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_NewRelationCreateToastTable(m, v60, v56)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							if v20 != 0 {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
								v64 = F_copyObjectImpl(m, v63)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									v66 = m.G0
									v68 = v66 - int32(32)
									m.G0 = v68
									v71 = F_pstrdup(m, int32(552929))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = v64
										*(*int32)(unsafe.Add(mBase, uint32(v68)+28)) = v64
										v77 = int32(0)
										v78 = int32(1)
										v84 = F_list_make1_impl(m, v78, v68+int32(12))
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											F_DefineQueryRewrite(m, v68+int32(16), v71, v60, v77, v78, v78, v77, v84)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												m.G0 = v68 + int32(32)
												F_CommandCounterIncrement(m)
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return
												} else {
													m.G0 = v10 + int32(16)
													return
												}
											}
										}
									}
								}
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_create_material_path(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 float64
	_ = v45
	var v49 float64
	_ = v49
	var v57 float64
	_ = v57
	var v63 float64
	_ = v63
	var v69 float64
	_ = v69
	var v71 int32
	_ = v71
	v7 = F_palloc0(m, int32(80))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(1546188226853)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v14
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v17 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)) = uint8(v17)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v16
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if v20 == int32(1) {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
			v24 = v23
		} else {
			v24 = int32(0)
		}
		v26 = v24 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+21)) = uint8(v26)
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v28
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+72)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v30
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		v34 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
		v35 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
		v36 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
		v42 = *(*int32)(unsafe.Add(mBase, _consts[135]))
		*(*float64)(unsafe.Add(mBase, uint32(v7)+32)) = v36
		v45 = *(*float64)(unsafe.Add(mBase, _consts[613]))
		v49 = base.F64_add(base.F64_mul(base.F64_add(v45, v45), v36), base.F64_sub(v35, v34))
		v57 = base.F64_mul(v36, base.F64_convert_i32_u((v38+int32(7))&int32(-8)+int32(24)))
		if base.F64_gt(v57, base.F64_convert_i32_u(v42<<(uint(int32(10))%32))) != 0 {
			v63 = *(*float64)(unsafe.Add(mBase, _consts[637]))
			v69 = base.F64_add(base.F64_mul(v63, base.F64_ceil(base.F64_mul(v57, float64(0.0001220703125)))), v49)
		} else {
			v69 = v49
		}
		v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[641])))
		*(*float64)(unsafe.Add(mBase, uint32(v7)+56)) = base.F64_add(v34, v69)
		*(*float64)(unsafe.Add(mBase, uint32(v7)+48)) = v34
		*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v33 + (v71 ^ int32(1))
		return v7
	}
}
func F_create_ordinary_grouping_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v27 float64
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v106 int32
	_ = v106
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v265 int32
	_ = v265
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v311 int32
	_ = v311
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
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
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v500 int32
	_ = v500
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v653 int32
	_ = v653
	var v654 int64
	_ = v654
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v686 int32
	_ = v686
	var v687 int64
	_ = v687
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v713 float64
	_ = v713
	var v714 int32
	_ = v714
	var v715 float64
	_ = v715
	var v716 int32
	_ = v716
	var v717 float64
	_ = v717
	var v718 float64
	_ = v718
	var v719 int32
	_ = v719
	var v720 float64
	_ = v720
	var v721 int32
	_ = v721
	var v722 float64
	_ = v722
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v753 int32
	_ = v753
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v784 int32
	_ = v784
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v831 int32
	_ = v831
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
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1031 int32
	_ = v1031
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1062 int32
	_ = v1062
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1157 int32
	_ = v1157
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1313 int32
	_ = v1313
	var v1323 int32
	_ = v1323
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1402 int32
	_ = v1402
	var v1420 int32
	_ = v1420
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1459 int32
	_ = v1459
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1627 int32
	_ = v1627
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1722 float64
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1724 float64
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1755 int32
	_ = v1755
	var v1771 int32
	_ = v1771
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1791 int32
	_ = v1791
	var v1812 int32
	_ = v1812
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1820 int32
	_ = v1820
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	var v1858 int32
	_ = v1858
	var v1862 int32
	_ = v1862
	var v1868 int32
	_ = v1868
	var v1872 int32
	_ = v1872
	var v1875 int32
	_ = v1875
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1886 int32
	_ = v1886
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1904 int32
	_ = v1904
	var v1908 int32
	_ = v1908
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1948 int32
	_ = v1948
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v2017 int32
	_ = v2017
	var v2020 int32
	_ = v2020
	var v2044 int32
	_ = v2044
	var v2052 int32
	_ = v2052
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2072 int32
	_ = v2072
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2102 int32
	_ = v2102
	var v2120 int32
	_ = v2120
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2132 int32
	_ = v2132
	var v2140 int32
	_ = v2140
	var v2144 int32
	_ = v2144
	var v2150 int32
	_ = v2150
	var v2154 int32
	_ = v2154
	var v2157 int32
	_ = v2157
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2168 int32
	_ = v2168
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2186 int32
	_ = v2186
	var v2190 int32
	_ = v2190
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2221 int32
	_ = v2221
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2302 int32
	_ = v2302
	var v2305 int32
	_ = v2305
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2331 int32
	_ = v2331
	var v2334 int32
	_ = v2334
	var v2339 int32
	_ = v2339
	var v2374 int32
	_ = v2374
	var v2377 int32
	_ = v2377
	var v2381 int32
	_ = v2381
	var v2385 int32
	_ = v2385
	var v2390 int32
	_ = v2390
	v8 = int32(0)
	v27 = float64(0)
	v29 = m.G0
	v31 = v29 - int32(128)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v34 = int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l5)+100))
	if v35 == v8 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v332&int32(4) == int32(0) {
		v1323 = v8
		goto L51
	} else {
		goto L52
	}
L2:
	;
	v311 = v8
	v325 = v8
	v326 = v34
	goto L1
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	if v38 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+240))
	if v41 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	if v44 <= int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	if v47 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v50 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v52 == v50 {
		v72 = v50
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v72 != 0 {
		v311 = v8
		v325 = v8
		v326 = v34
		goto L1
	} else {
		goto L18
	}
L9:
	;
	goto L8
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v56 = v55
	goto L11
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if base.Ui32(int32(2)) <= base.Ui32(v60-int32(301)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v72 = int32(1)
	goto L9
L13:
	;
	if v60 != int32(290) {
		v72 = v50
		goto L9
	} else {
		goto L16
	}
L14:
	;
	v56 = v59 + int32(72)
	goto L11
L15:
	;
	goto L12
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v59)+72))
	if v67 != 0 {
		v72 = v50
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l5)+100))
	if v74 != int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v297 = v295 & int32(4)
	v311 = int32(base.Ui32(v297) >> (uint(int32(2)) % 32))
	v325 = int32(base.Ui32(v297) >> (uint(int32(1)) % 32))
	v326 = base.B2i32(v297 == int32(0))
	goto L1
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+100))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v80 = F_get_sortgrouplist_exprs(m, v78, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	if v82 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v85 = int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v86)+2)))
	if v87 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v311 = v8
	v325 = v85
	v326 = int32(0)
	goto L1
L25:
	;
	goto L26
L26:
	;
	v106 = v8
	goto L27
L27:
	;
	v120 = v106 << (uint(int32(2)) % 32)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120+v121)))
	if v123 == int32(0) {
		goto L19
	} else {
		goto L29
	}
L28:
	;
	goto L19
L29:
	;
	v126 = int32(0)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v127 <= v126 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v137 = v127
	v144 = v126
	goto L31
L31:
	;
	if v80 == int32(0) {
		v243 = v137
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L28
L33:
	;
	v265 = v144 + int32(1)
	if v265 < v243 {
		v137 = v243
		v144 = v265
		goto L31
	} else {
		goto L50
	}
L34:
	;
	v160 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v161 <= v160 {
		v243 = v137
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v164+v144<<(uint(int32(2))%32))))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v170+v120)))
	v181 = v160
	goto L37
L36:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v243 = v235
	goto L33
L37:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201+v181<<(uint(int32(2))%32))))
	v206 = F_exprCollation(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L21
	} else {
		goto L39
	}
L38:
	;
	v221 = int32(0)
	if base.B2i32(v172 == v221)|base.B2i32(v206 == v221)|base.B2i32(v172 == v206) == v221 {
		goto L19
	} else {
		goto L48
	}
L39:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	if v208 == int32(27) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	v212 = v211
	goto L42
L41:
	;
	v212 = v205
	goto L42
L42:
	;
	v213 = F_equal(m, v212, v168)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L21
	} else {
		goto L43
	}
L43:
	;
	if v213 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v218 = v181 + int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v219 <= v218 {
		goto L36
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	goto L38
L47:
	;
	v181 = v218
	goto L37
L48:
	;
	v230 = int32(0)
	v233 = v106 + int32(1)
	if v233 == v87 {
		v311 = v230
		v325 = v85
		v326 = v230
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v106 = v233
	goto L27
L50:
	;
	goto L32
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1323
	if v326 != 0 {
		goto L311
	} else {
		goto L312
	}
L52:
	;
	v337 = int32(0)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v339 == v337 {
		v347 = v337
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+26)))
	if v348 != int32(1) {
		v357 = v337
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l5)+100))
	if v343 != int32(2) {
		v347 = int32(0)
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v347 = v346
	goto L53
L56:
	;
	v358 = int32(0)
	if v311|base.B2i32(v347|v357 != v358) != int32(1) {
		v1323 = v358
		goto L51
	} else {
		goto L59
	}
L57:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v351 == int32(0) {
		v357 = v337
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	v357 = v355
	goto L56
L59:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v368 = F_fetch_upper_rel(m, l0, int32(1), v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L21
	} else {
		goto L60
	}
L60:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v368)+26)) = uint8(v370)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v368)+4)) = v372
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l2)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v368)+156)) = v374
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l2)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v368)+160)) = v376
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+164)))
	*(*uint8)(unsafe.Add(mBase, uint32(v368)+164)) = uint8(v378)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l2)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v368)+168)) = v380
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l5)+92))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v384 = F_create_empty_pathtarget(m)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L21
	} else {
		goto L61
	}
L61:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v386 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	if v382 != 0 {
		goto L92
	} else {
		goto L93
	}
L63:
	;
	v401 = v387
	v405 = int32(0)
	goto L68
L64:
	;
	v387 = int32(0)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	if v387 < v388 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v500 = int32(0)
	goto L62
L67:
	;
	goto L66
L68:
	;
	v423 = v401 << (uint(int32(2)) % 32)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v423+v424)))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v383)+8))
	if v427 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v500 = v482
	goto L62
L70:
	;
	v486 = v401 + int32(1)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	if v486 < v487 {
		v401 = v486
		v405 = v482
		goto L68
	} else {
		goto L91
	}
L71:
	;
	v480 = F_lappend(m, v405, v426)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L21
	} else {
		goto L90
	}
L72:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v427+v423)))
	if v431 == int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v434 == int32(0) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	if v434 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	if v472 == int32(0) {
		goto L71
	} else {
		goto L88
	}
L76:
	;
	goto L75
L77:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	if v440 <= int32(0) {
		v472 = int32(0)
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v472 = int32(0)
	goto L76
L80:
	;
	v443 = int32(0)
	if v443 < v440 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v446 = v440
	goto L83
L82:
	;
	v446 = v443
	goto L83
L83:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v434)+12))
	v450 = int32(0)
	goto L84
L84:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v447+v450<<(uint(int32(2))%32))))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)+4))
	if v458 == v431 {
		v472 = v457
		goto L76
	} else {
		goto L86
	}
L85:
	;
	goto L79
L86:
	;
	v461 = v450 + int32(1)
	if v461 != v446 {
		v450 = v461
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	F_add_column_to_pathtarget(m, v384, v426, v431)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L21
	} else {
		goto L89
	}
L89:
	;
	v482 = v405
	goto L70
L90:
	;
	v482 = v480
	goto L70
L91:
	;
	goto L69
L92:
	;
	v517 = F_lappend(m, v500, v382)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L21
	} else {
		goto L95
	}
L93:
	;
	v519 = v500
	goto L94
L94:
	;
	v521 = F_pull_var_clause(m, v519, int32(25))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L21
	} else {
		goto L96
	}
L95:
	;
	v519 = v517
	goto L94
L96:
	;
	F_add_new_columns_to_pathtarget(m, v384, v521)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L21
	} else {
		goto L97
	}
L97:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v384)+4))
	if v525 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v620 = l5 + int32(8)
	F_list_free(m, v521)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L21
	} else {
		goto L115
	}
L99:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	if v528 <= int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v539 = int32(0)
	v540 = v528
	goto L101
L101:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v525)+12))
	v563 = v560 + v539<<(uint(int32(2))%32)
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	if v565 == int32(9) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L98
L103:
	;
	v569 = F_palloc0(m, int32(72))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L21
	} else {
		goto L106
	}
L104:
	;
	v586 = v540
	goto L105
L105:
	;
	v589 = v539 + int32(1)
	if v589 < v586 {
		v539 = v589
		v540 = v586
		goto L101
	} else {
		goto L114
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = int32(9)
	goto L108
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v574)+56)) = int32(6)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v574)+20))
	if v579 == int32(2281) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v574 = F__emscripten_memcpy_bulkmem(m, v569, v564, int32(72))
	mBase = m.M
	goto L110
L110:
	;
	goto L107
L111:
	;
	v582 = int32(17)
	goto L113
L112:
	;
	v582 = v579
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v574)+8)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v563))) = v574
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	v586 = v585
	goto L105
L114:
	;
	goto L102
L115:
	;
	F_list_free(m, v519)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L21
	} else {
		goto L116
	}
L116:
	;
	v625 = F_set_pathtarget_cost_width(m, l0, v384)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L21
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368)+28)) = v625
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+4)))
	if v628 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v632 = l5 + int32(48)
	if v620&int32(3) == int32(0) {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	goto L120
L120:
	;
	if v347 != 0 {
		goto L144
	} else {
		goto L145
	}
L121:
	;
	if v632&int32(3) == int32(0) {
		goto L131
	} else {
		goto L132
	}
L122:
	;
	if base.Ui32(v632) <= base.Ui32(v620) {
		goto L121
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v654 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v620))) = v654
	*(*int64)(unsafe.Add(mBase, uint32(v620)+32)) = v654
	*(*int64)(unsafe.Add(mBase, uint32(v620)+24)) = v654
	*(*int64)(unsafe.Add(mBase, uint32(v620)+16)) = v654
	*(*int64)(unsafe.Add(mBase, uint32(v620)+8)) = v654
	goto L121
L125:
	;
	v640 = l5 + int32(48)
	v642 = l5 + int32(12)
	if base.Ui32(v642) < base.Ui32(v640) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v644 = v640
	goto L128
L127:
	;
	v644 = v642
	goto L128
L128:
	;
	v653 = F__emscripten_memset_bulkmem(m, v620, base.I32_extend8_s(int32(0)), (v644-l5-int32(9))&int32(-4)+int32(4))
	mBase = m.M
	goto L129
L129:
	;
	goto L121
L130:
	;
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+36)))
	if v699 == int32(1) {
		goto L139
	} else {
		goto L140
	}
L131:
	;
	v671 = l5 + int32(88)
	if base.Ui32(v671) <= base.Ui32(v632) {
		goto L130
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v687 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v632))) = v687
	*(*int64)(unsafe.Add(mBase, uint32(v632)+32)) = v687
	*(*int64)(unsafe.Add(mBase, uint32(v632)+24)) = v687
	*(*int64)(unsafe.Add(mBase, uint32(v632)+16)) = v687
	*(*int64)(unsafe.Add(mBase, uint32(v632)+8)) = v687
	goto L130
L134:
	;
	v675 = l5 + int32(52)
	if base.Ui32(v675) < base.Ui32(v671) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v677 = v671
	goto L137
L136:
	;
	v677 = v675
	goto L137
L137:
	;
	v686 = F__emscripten_memset_bulkmem(m, v632, base.I32_extend8_s(int32(0)), (v677-l5-int32(49))&int32(-4)+int32(4))
	mBase = m.M
	goto L138
L138:
	;
	goto L130
L139:
	;
	F_get_agg_clause_costs(m, l0, int32(6), v620)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L21
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v708 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+4)) = uint8(v708)
	goto L120
L142:
	;
	F_get_agg_clause_costs(m, l0, int32(9), v632)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L21
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v713 = *(*float64)(unsafe.Add(mBase, uint32(v347)+32))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v715 = F_get_number_of_groups(m, l0, v713, l4, v714)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L21
	} else {
		goto L147
	}
L145:
	;
	v717 = v27
	goto L146
L146:
	;
	if v357 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v717 = v715
	goto L146
L148:
	;
	v718 = *(*float64)(unsafe.Add(mBase, uint32(v357)+32))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v720 = F_get_number_of_groups(m, l0, v718, l4, v719)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L21
	} else {
		goto L151
	}
L149:
	;
	v722 = v27
	goto L150
L150:
	;
	v724 = v332 & int32(1)
	if v724 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v722 = v720
	goto L150
L152:
	;
	v1002 = v332 & int32(2)
	if v724 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L153:
	;
	if v347 == int32(0) {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v729 == int32(0) {
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v729)+4))
	if v732 <= int32(0) {
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v753 = int32(0)
	goto L157
L157:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v729)+12))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v764+v753<<(uint(int32(2))%32))))
	v769 = F_get_useful_group_keys_orderings(m, l0, v768)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L21
	} else {
		goto L160
	}
L158:
	;
	goto L152
L159:
	;
	v970 = v753 + int32(1)
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v729)+4))
	if v970 < v971 {
		v753 = v970
		goto L157
	} else {
		goto L224
	}
L160:
	;
	if v769 == int32(0) {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v773 = int32(0)
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v769)+4))
	if v774 <= v773 {
		goto L159
	} else {
		goto L162
	}
L162:
	;
	v784 = v773
	goto L163
L163:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v769)+12))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v805+v784<<(uint(int32(2))%32))))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v809)+4))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v768)+64))
	v813 = v31 + int32(16)
	if v810 == v811 {
		goto L169
	} else {
		goto L170
	}
L164:
	;
	goto L159
L165:
	;
	v938 = v784 + int32(1)
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v769)+4))
	if v938 < v939 {
		v784 = v938
		goto L163
	} else {
		goto L223
	}
L166:
	;
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+36)))
	if v915 == int32(1) {
		goto L216
	} else {
		goto L217
	}
L167:
	;
	if v891 != 0 {
		goto L199
	} else {
		goto L200
	}
L168:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v810)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v813))) = v879
	v891 = int32(1)
	goto L167
L169:
	;
	if v810 != 0 {
		goto L168
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	if v810 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v813))) = int32(0)
	v891 = int32(1)
	goto L167
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v813))) = int32(0)
	v891 = int32(1)
	goto L167
L174:
	;
	goto L175
L175:
	;
	if v811 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v831 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v813))) = v831
	v891 = v831
	goto L167
L177:
	;
	goto L178
L178:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v811)+4))
	v835 = int32(0)
	if v835 < v834 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v838 = v834
	goto L181
L180:
	;
	v838 = v835
	goto L181
L181:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v810)+4))
	v843 = int32(0)
	goto L182
L182:
	;
	if v843 < v839 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v810)+12))
	v855 = v851 + v843<<(uint(int32(2))%32)
	goto L186
L185:
	;
	v855 = int32(0)
	goto L186
L186:
	;
	if v843 == v838 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v813))) = v838
	v891 = base.B2i32(v855 == int32(0))
	goto L167
L188:
	;
	goto L189
L189:
	;
	v861 = base.B2i32(v855 == int32(0))
	if v855 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v813))) = v843
	v891 = v861
	goto L167
L191:
	;
	goto L192
L192:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v811)+12))
	v868 = v865 + v843<<(uint(int32(2))%32)
	if v868 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v813))) = v843
	v891 = v861
	goto L167
L194:
	;
	goto L195
L195:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v855)))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	if v872 != v873 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v813))) = v843
	v891 = int32(0)
	goto L167
L197:
	;
	v843 = v843 + int32(1)
	goto L182
L199:
	;
	v913 = v768
	goto L166
L200:
	;
	goto L201
L201:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	if v768 != v347 {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	if v910 == int32(0) {
		goto L165
	} else {
		goto L215
	}
L203:
	;
	v908 = F_create_incremental_sort_path(m, l0, v368, v768, v810, v892, float64(-1))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L21
	} else {
		goto L214
	}
L204:
	;
	if v892 == int32(0) {
		goto L165
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	if v892 != 0 {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	v897 = int32(*(*uint8)(unsafe.Add(mBase, _consts[639])))
	if v897 == int32(0) {
		goto L165
	} else {
		goto L208
	}
L208:
	;
	goto L203
L209:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, _consts[639])))
	if v901&int32(1) != 0 {
		goto L203
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v905 = F_create_sort_path(m, v368, v768, v810, float64(-1))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L21
	} else {
		goto L213
	}
L212:
	;
	goto L211
L213:
	;
	v910 = v905
	goto L202
L214:
	;
	v910 = v908
	goto L202
L215:
	;
	v913 = v910
	goto L166
L216:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v368)+28))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v365)+100))
	v920 = int32(0)
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v809)+8))
	v925 = F_create_agg_path(m, l0, v368, v913, v918, base.B2i32(v919 != v920), int32(6), v923, v920, v620, v717)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L21
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v809)+8))
	v931 = F_create_group_path(m, l0, v368, v913, v929, int32(0), v717)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L21
	} else {
		goto L221
	}
L219:
	;
	F_add_path(m, v368, v925)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L21
	} else {
		goto L220
	}
L220:
	;
	goto L165
L221:
	;
	F_add_path(m, v368, v931)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L21
	} else {
		goto L222
	}
L222:
	;
	goto L165
L223:
	;
	goto L164
L224:
	;
	goto L158
L225:
	;
	if v1002 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L226:
	;
	if v357 == int32(0) {
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v1007 == int32(0) {
		goto L225
	} else {
		goto L228
	}
L228:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1007)+4))
	if v1010 <= int32(0) {
		goto L225
	} else {
		goto L229
	}
L229:
	;
	v1031 = int32(0)
	goto L230
L230:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1007)+12))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1042+v1031<<(uint(int32(2))%32))))
	v1047 = F_get_useful_group_keys_orderings(m, l0, v1046)
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L21
	} else {
		goto L233
	}
L231:
	;
	goto L225
L232:
	;
	v1248 = v1031 + int32(1)
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1007)+4))
	if v1248 < v1249 {
		v1031 = v1248
		goto L230
	} else {
		goto L297
	}
L233:
	;
	if v1047 == int32(0) {
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v1051 = int32(0)
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+4))
	if v1052 <= v1051 {
		goto L232
	} else {
		goto L235
	}
L235:
	;
	v1062 = v1051
	goto L236
L236:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+12))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1083+v1062<<(uint(int32(2))%32))))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1087)+4))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1046)+64))
	v1091 = v31 + int32(16)
	if v1088 == v1089 {
		goto L242
	} else {
		goto L243
	}
L237:
	;
	goto L232
L238:
	;
	v1216 = v1062 + int32(1)
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+4))
	if v1216 < v1217 {
		v1062 = v1216
		goto L236
	} else {
		goto L296
	}
L239:
	;
	v1193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+36)))
	if v1193 == int32(1) {
		goto L289
	} else {
		goto L290
	}
L240:
	;
	if v1169 != 0 {
		goto L272
	} else {
		goto L273
	}
L241:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1091))) = v1157
	v1169 = int32(1)
	goto L240
L242:
	;
	if v1088 != 0 {
		goto L241
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	if v1088 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1091))) = int32(0)
	v1169 = int32(1)
	goto L240
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1091))) = int32(0)
	v1169 = int32(1)
	goto L240
L247:
	;
	goto L248
L248:
	;
	if v1089 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1109 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1091))) = v1109
	v1169 = v1109
	goto L240
L250:
	;
	goto L251
L251:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+4))
	v1113 = int32(0)
	if v1113 < v1112 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1116 = v1112
	goto L254
L253:
	;
	v1116 = v1113
	goto L254
L254:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+4))
	v1121 = int32(0)
	goto L255
L255:
	;
	if v1121 < v1117 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+12))
	v1133 = v1129 + v1121<<(uint(int32(2))%32)
	goto L259
L258:
	;
	v1133 = int32(0)
	goto L259
L259:
	;
	if v1121 == v1116 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1091))) = v1116
	v1169 = base.B2i32(v1133 == int32(0))
	goto L240
L261:
	;
	goto L262
L262:
	;
	v1139 = base.B2i32(v1133 == int32(0))
	if v1133 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1091))) = v1121
	v1169 = v1139
	goto L240
L264:
	;
	goto L265
L265:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+12))
	v1146 = v1143 + v1121<<(uint(int32(2))%32)
	if v1146 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1091))) = v1121
	v1169 = v1139
	goto L240
L267:
	;
	goto L268
L268:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1133)))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1146)))
	if v1150 != v1151 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1091))) = v1121
	v1169 = int32(0)
	goto L240
L270:
	;
	v1121 = v1121 + int32(1)
	goto L255
L272:
	;
	v1191 = v1046
	goto L239
L273:
	;
	goto L274
L274:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	if v1046 != v357 {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	if v1188 == int32(0) {
		goto L238
	} else {
		goto L288
	}
L276:
	;
	v1186 = F_create_incremental_sort_path(m, l0, v368, v1046, v1088, v1170, float64(-1))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L21
	} else {
		goto L287
	}
L277:
	;
	if v1170 == int32(0) {
		goto L238
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	if v1170 != 0 {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	v1175 = int32(*(*uint8)(unsafe.Add(mBase, _consts[639])))
	if v1175 == int32(0) {
		goto L238
	} else {
		goto L281
	}
L281:
	;
	goto L276
L282:
	;
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, _consts[639])))
	if v1179&int32(1) != 0 {
		goto L276
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v1183 = F_create_sort_path(m, v368, v1046, v1088, float64(-1))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L21
	} else {
		goto L286
	}
L285:
	;
	goto L284
L286:
	;
	v1188 = v1183
	goto L275
L287:
	;
	v1188 = v1186
	goto L275
L288:
	;
	v1191 = v1188
	goto L239
L289:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v368)+28))
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v365)+100))
	v1198 = int32(0)
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1087)+8))
	v1203 = F_create_agg_path(m, l0, v368, v1191, v1196, base.B2i32(v1197 != v1198), int32(6), v1201, v1198, v620, v722)
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L21
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1087)+8))
	v1209 = F_create_group_path(m, l0, v368, v1191, v1207, int32(0), v722)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L21
	} else {
		goto L294
	}
L292:
	;
	F_add_partial_path(m, v368, v1203)
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L21
	} else {
		goto L293
	}
L293:
	;
	goto L238
L294:
	;
	F_add_partial_path(m, v368, v1209)
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L21
	} else {
		goto L295
	}
L295:
	;
	goto L238
L296:
	;
	goto L237
L297:
	;
	goto L231
L298:
	;
	if v1002 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L299:
	;
	if v347 == int32(0) {
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v368)+28))
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v1288 = F_create_agg_path(m, l0, v368, v347, v1283, int32(2), int32(6), v1286, int32(0), v620, v717)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L21
	} else {
		goto L301
	}
L301:
	;
	F_add_path(m, v368, v1288)
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L21
	} else {
		goto L302
	}
L302:
	;
	goto L298
L303:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v368)+168))
	if v1305 == int32(0) {
		v1323 = v368
		goto L51
	} else {
		goto L308
	}
L304:
	;
	if v357 == int32(0) {
		goto L303
	} else {
		goto L305
	}
L305:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v368)+28))
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v1301 = F_create_agg_path(m, l0, v368, v357, v1296, int32(2), int32(6), v1299, int32(0), v620, v722)
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L21
	} else {
		goto L306
	}
L306:
	;
	F_add_partial_path(m, v368, v1301)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L21
	} else {
		goto L307
	}
L307:
	;
	goto L303
L308:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+36))
	if v1308 == int32(0) {
		v1323 = v368
		goto L51
	} else {
		goto L309
	}
L309:
	;
	m.T0[v1308].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, int32(1), l1, v368, l5)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L21
	} else {
		goto L310
	}
L310:
	;
	v1323 = v368
	goto L51
L311:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(l5)+100))
	if v1703 == int32(2) {
		goto L396
	} else {
		goto L397
	}
L312:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v1344 = int32(0)
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v1345 == v1344 {
		goto L316
	} else {
		goto L317
	}
L313:
	;
	if v1323 == int32(0) {
		goto L386
	} else {
		goto L387
	}
L314:
	;
	if v1402 < int32(0) {
		goto L325
	} else {
		goto L326
	}
L315:
	;
	v1402 = base.I32_ctz(v1388) | v1389<<(uint(int32(5))%32)
	goto L314
L316:
	;
	v1402 = int32(-2)
	goto L314
L317:
	;
	v1355 = base.I32_div_s(int32(0), int32(32))
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1345)+4))
	if v1356 <= v1355 {
		goto L316
	} else {
		goto L318
	}
L318:
	;
	v1359 = v1345 + int32(8)
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1359+v1355<<(uint(int32(2))%32))))
	v1366 = v1363 & int32(-1)
	if v1366 != 0 {
		v1388 = v1366
		v1389 = v1355
		goto L315
	} else {
		goto L319
	}
L319:
	;
	v1368 = v1355 + int32(1)
	if v1368 == v1356 {
		goto L316
	} else {
		goto L320
	}
L320:
	;
	v1371 = v1368
	goto L321
L321:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1359+v1371<<(uint(int32(2))%32))))
	if v1378 != 0 {
		v1388 = v1378
		v1389 = v1371
		goto L315
	} else {
		goto L323
	}
L322:
	;
	goto L316
L323:
	;
	v1380 = v1371 + int32(1)
	if v1380 != v1356 {
		v1371 = v1380
		goto L321
	} else {
		goto L324
	}
L324:
	;
	goto L322
L325:
	;
	v1645 = v1344
	v1647 = int32(1)
	v1649 = int32(0)
	goto L313
L326:
	;
	goto L327
L327:
	;
	v1420 = v1402
	v1424 = v1344
	v1426 = int32(1)
	v1428 = int32(0)
	goto L328
L328:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1437+v1420<<(uint(int32(2))%32))))
	v1442 = int32(0)
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+32))
	if v1444 == v1442 {
		v1464 = v1442
		goto L331
	} else {
		goto L332
	}
L329:
	;
	v1645 = v1567
	v1647 = v1568
	v1649 = v1570
	goto L313
L330:
	;
	if v1464 == int32(0) {
		goto L340
	} else {
		goto L341
	}
L331:
	;
	goto L330
L332:
	;
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1444)+12))
	v1448 = v1447
	goto L333
L333:
	;
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1448)))
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1451)))
	if base.Ui32(int32(2)) <= base.Ui32(v1452-int32(301)) {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	v1464 = int32(1)
	goto L331
L335:
	;
	if v1452 != int32(290) {
		v1464 = v1442
		goto L331
	} else {
		goto L338
	}
L336:
	;
	v1448 = v1451 + int32(72)
	goto L333
L337:
	;
	goto L334
L338:
	;
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+72))
	if v1459 != 0 {
		v1464 = v1442
		goto L331
	} else {
		goto L339
	}
L339:
	;
	goto L337
L340:
	;
	v1468 = F_copy_pathtarget(m, v1343)
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L21
	} else {
		goto L343
	}
L341:
	;
	v1567 = v1424
	v1568 = v1426
	v1570 = v1428
	goto L342
L342:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v1571 == int32(0) {
		goto L376
	} else {
		goto L377
	}
L343:
	;
	goto L345
L344:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+8))
	v1478 = F_find_appinfos_by_relids(m, l0, v1475, v31+int32(124))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L21
	} else {
		goto L348
	}
L345:
	;
	v1473 = F__emscripten_memcpy_bulkmem(m, v31+int32(16), l5, int32(104))
	mBase = m.M
	goto L347
L347:
	;
	goto L344
L348:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+4))
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v31)+124))
	v1482 = F_adjust_appendrel_attrs(m, l0, v1480, v1481, v1478)
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L21
	} else {
		goto L349
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1468)+4)) = v1482
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(l5)+92))
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v31)+124))
	v1487 = F_adjust_appendrel_attrs(m, l0, v1485, v1486, v1478)
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L21
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+108)) = v1487
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v31)+124))
	v1492 = F_adjust_appendrel_attrs(m, l0, v1490, v1491, v1478)
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L21
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v325
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1492
	v1496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+88)))
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+4))
	if base.Ui32(int32(5)) < base.Ui32(v1498) {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1517)+28)) = v1468
	v1519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1441)+26)))
	if v1519&v1496&int32(1) == int32(0) {
		goto L358
	} else {
		goto L359
	}
L353:
	;
	v1515 = F_fetch_upper_rel(m, l0, int32(2), int32(0))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L21
	} else {
		goto L357
	}
L354:
	;
	if int32(1)<<(uint(v1498)%32)&int32(44) == int32(0) {
		goto L353
	} else {
		goto L355
	}
L355:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+8))
	v1509 = F_fetch_upper_rel(m, l0, int32(2), v1508)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L21
	} else {
		goto L356
	}
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+4)) = int32(5)
	v1517 = v1509
	goto L352
L357:
	;
	v1517 = v1515
	goto L352
L358:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v1517)+156)) = v1531
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v1517)+160)) = v1533
	v1535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1441)+164)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1517)+164)) = uint8(v1535)
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v1517)+168)) = v1537
	F_create_ordinary_grouping_paths(m, l0, v1441, v1517, l3, l4, v31+int32(16), v31+int32(12))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L21
	} else {
		goto L362
	}
L359:
	;
	v1525 = F_is_parallel_safe(m, l0, v1497)
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L21
	} else {
		goto L360
	}
L360:
	;
	if v1525 == int32(0) {
		goto L358
	} else {
		goto L361
	}
L361:
	;
	v1529 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1517)+26)) = uint8(v1529)
	goto L358
L362:
	;
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v1545 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L363:
	;
	if v325 == int32(1) {
		goto L368
	} else {
		goto L369
	}
L364:
	;
	v1551 = v1424
	v1552 = int32(0)
	goto L363
L365:
	;
	goto L366
L366:
	;
	v1549 = F_lappend(m, v1424, v1545)
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L21
	} else {
		goto L367
	}
L367:
	;
	v1551 = v1549
	v1552 = v1426
	goto L363
L368:
	;
	F_set_cheapest(m, v1517)
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L21
	} else {
		goto L371
	}
L369:
	;
	v1559 = v1428
	goto L370
L370:
	;
	F_pfree(m, v1478)
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L21
	} else {
		goto L373
	}
L371:
	;
	v1557 = F_lappend(m, v1428, v1517)
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L21
	} else {
		goto L372
	}
L372:
	;
	v1559 = v1557
	goto L370
L373:
	;
	v1567 = v1551
	v1568 = v1552
	v1570 = v1559
	goto L342
L374:
	;
	if int32(0) <= v1627 {
		v1420 = v1627
		v1424 = v1567
		v1426 = v1568
		v1428 = v1570
		goto L328
	} else {
		goto L385
	}
L375:
	;
	v1627 = base.I32_ctz(v1613) | v1614<<(uint(int32(5))%32)
	goto L374
L376:
	;
	v1627 = int32(-2)
	goto L374
L377:
	;
	v1578 = v1420 + int32(1)
	v1580 = base.I32_div_s(v1578, int32(32))
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1571)+4))
	if v1581 <= v1580 {
		goto L376
	} else {
		goto L378
	}
L378:
	;
	v1584 = v1571 + int32(8)
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v1584+v1580<<(uint(int32(2))%32))))
	v1591 = v1588 & (int32(-1) << (uint(v1578) % 32))
	if v1591 != 0 {
		v1613 = v1591
		v1614 = v1580
		goto L375
	} else {
		goto L379
	}
L379:
	;
	v1593 = v1580 + int32(1)
	if v1593 == v1581 {
		goto L376
	} else {
		goto L380
	}
L380:
	;
	v1596 = v1593
	goto L381
L381:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v1584+v1596<<(uint(int32(2))%32))))
	if v1603 != 0 {
		v1613 = v1603
		v1614 = v1596
		goto L375
	} else {
		goto L383
	}
L382:
	;
	goto L376
L383:
	;
	v1605 = v1596 + int32(1)
	if v1605 != v1581 {
		v1596 = v1605
		goto L381
	} else {
		goto L384
	}
L384:
	;
	goto L382
L385:
	;
	goto L329
L386:
	;
	if v325 != int32(1) {
		goto L311
	} else {
		goto L392
	}
L387:
	;
	if v1647&int32(1) == int32(0) {
		goto L386
	} else {
		goto L388
	}
L388:
	;
	F_add_paths_to_append_rel(m, l0, v1323, v1645)
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L21
	} else {
		goto L389
	}
L389:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1323)+32))
	if v1666 == int32(0) {
		goto L386
	} else {
		goto L390
	}
L390:
	;
	F_set_cheapest(m, v1323)
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L21
	} else {
		goto L391
	}
L391:
	;
	goto L386
L392:
	;
	F_add_paths_to_append_rel(m, l0, l2, v1649)
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L21
	} else {
		goto L393
	}
L393:
	;
	goto L311
L394:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L21
	} else {
		goto L580
	}
L395:
	;
	m.G0 = v31 + int32(128)
	return
L396:
	;
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1323)+32))
	if v1706 == int32(0) {
		goto L395
	} else {
		goto L399
	}
L397:
	;
	goto L398
L398:
	;
	if v1323 == int32(0) {
		goto L401
	} else {
		goto L402
	}
L399:
	;
	F_set_cheapest(m, v1323)
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L21
	} else {
		goto L400
	}
L400:
	;
	goto L395
L401:
	;
	v1721 = l5 + int32(48)
	v1722 = *(*float64)(unsafe.Add(mBase, uint32(v33)+32))
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v1724 = F_get_number_of_groups(m, l0, v1722, l4, v1723)
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L21
	} else {
		goto L406
	}
L402:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1323)+40))
	if v1713 == int32(0) {
		goto L401
	} else {
		goto L403
	}
L403:
	;
	F_gather_grouping_paths(m, l0, v1323)
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L21
	} else {
		goto L404
	}
L404:
	;
	F_set_cheapest(m, v1323)
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L21
	} else {
		goto L405
	}
L405:
	;
	goto L401
L406:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1728 = v1726 & int32(2)
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(l5)+92))
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1726&int32(1) == int32(0) {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	if v1728 == int32(0) {
		goto L556
	} else {
		goto L557
	}
L408:
	;
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v1736 == int32(0) {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	if v1323 == int32(0) {
		goto L407
	} else {
		goto L485
	}
L410:
	;
	v1739 = int32(0)
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1736)+4))
	if v1740 <= v1739 {
		goto L409
	} else {
		goto L411
	}
L411:
	;
	v1755 = v1739
	goto L412
L412:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1736)+12))
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(v1771+v1755<<(uint(int32(2))%32))))
	v1776 = F_get_useful_group_keys_orderings(m, l0, v1775)
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L21
	} else {
		goto L415
	}
L413:
	;
	goto L409
L414:
	;
	v1984 = v1755 + int32(1)
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v1736)+4))
	if v1984 < v1985 {
		v1755 = v1984
		goto L412
	} else {
		goto L484
	}
L415:
	;
	if v1776 == int32(0) {
		goto L414
	} else {
		goto L416
	}
L416:
	;
	v1780 = int32(0)
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+4))
	if v1781 <= v1780 {
		goto L414
	} else {
		goto L417
	}
L417:
	;
	v1791 = v1780
	goto L418
L418:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+12))
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1812+v1791<<(uint(int32(2))%32))))
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+4))
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1775)+64))
	v1820 = v31 + int32(16)
	if v1817 == v1818 {
		goto L424
	} else {
		goto L425
	}
L419:
	;
	goto L414
L420:
	;
	v1952 = v1791 + int32(1)
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+4))
	if v1952 < v1953 {
		v1791 = v1952
		goto L418
	} else {
		goto L483
	}
L421:
	;
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v1731)+108))
	if v1922 != 0 {
		goto L471
	} else {
		goto L472
	}
L422:
	;
	if v1898 != 0 {
		goto L454
	} else {
		goto L455
	}
L423:
	;
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(v1817)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1820))) = v1886
	v1898 = int32(1)
	goto L422
L424:
	;
	if v1817 != 0 {
		goto L423
	} else {
		goto L427
	}
L425:
	;
	goto L426
L426:
	;
	if v1817 == int32(0) {
		goto L428
	} else {
		goto L429
	}
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1820))) = int32(0)
	v1898 = int32(1)
	goto L422
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1820))) = int32(0)
	v1898 = int32(1)
	goto L422
L429:
	;
	goto L430
L430:
	;
	if v1818 == int32(0) {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1838 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1820))) = v1838
	v1898 = v1838
	goto L422
L432:
	;
	goto L433
L433:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+4))
	v1842 = int32(0)
	if v1842 < v1841 {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1845 = v1841
	goto L436
L435:
	;
	v1845 = v1842
	goto L436
L436:
	;
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1817)+4))
	v1850 = int32(0)
	goto L437
L437:
	;
	if v1850 < v1846 {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1817)+12))
	v1862 = v1858 + v1850<<(uint(int32(2))%32)
	goto L441
L440:
	;
	v1862 = int32(0)
	goto L441
L441:
	;
	if v1850 == v1845 {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1820))) = v1845
	v1898 = base.B2i32(v1862 == int32(0))
	goto L422
L443:
	;
	goto L444
L444:
	;
	v1868 = base.B2i32(v1862 == int32(0))
	if v1862 == int32(0) {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1820))) = v1850
	v1898 = v1868
	goto L422
L446:
	;
	goto L447
L447:
	;
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+12))
	v1875 = v1872 + v1850<<(uint(int32(2))%32)
	if v1875 == int32(0) {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1820))) = v1850
	v1898 = v1868
	goto L422
L449:
	;
	goto L450
L450:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1862)))
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1875)))
	if v1879 != v1880 {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1820))) = v1850
	v1898 = int32(0)
	goto L422
L452:
	;
	v1850 = v1850 + int32(1)
	goto L437
L454:
	;
	v1920 = v1775
	goto L421
L455:
	;
	goto L456
L456:
	;
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	if v1775 != v1730 {
		goto L459
	} else {
		goto L460
	}
L457:
	;
	if v1917 == int32(0) {
		goto L420
	} else {
		goto L470
	}
L458:
	;
	v1915 = F_create_incremental_sort_path(m, l0, l2, v1775, v1817, v1899, float64(-1))
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L21
	} else {
		goto L469
	}
L459:
	;
	if v1899 == int32(0) {
		goto L420
	} else {
		goto L462
	}
L460:
	;
	goto L461
L461:
	;
	if v1899 != 0 {
		goto L464
	} else {
		goto L465
	}
L462:
	;
	v1904 = int32(*(*uint8)(unsafe.Add(mBase, _consts[639])))
	if v1904 == int32(0) {
		goto L420
	} else {
		goto L463
	}
L463:
	;
	goto L458
L464:
	;
	v1908 = int32(*(*uint8)(unsafe.Add(mBase, _consts[639])))
	if v1908&int32(1) != 0 {
		goto L458
	} else {
		goto L467
	}
L465:
	;
	goto L466
L466:
	;
	v1912 = F_create_sort_path(m, l2, v1775, v1817, float64(-1))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L21
	} else {
		goto L468
	}
L467:
	;
	goto L466
L468:
	;
	v1917 = v1912
	goto L457
L469:
	;
	v1917 = v1915
	goto L457
L470:
	;
	v1920 = v1917
	goto L421
L471:
	;
	F_consider_groupingsets_paths(m, l0, l2, v1920, int32(1), base.B2i32(v1728 != int32(0)), l4, l3, v1724)
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L21
	} else {
		goto L474
	}
L472:
	;
	goto L473
L473:
	;
	v1928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1731)+36)))
	if v1928 == int32(1) {
		goto L475
	} else {
		goto L476
	}
L474:
	;
	goto L420
L475:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1731)+100))
	v1933 = int32(0)
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+8))
	v1937 = F_create_agg_path(m, l0, l2, v1920, v1931, base.B2i32(v1932 != v1933), v1933, v1936, v1729, l3, v1724)
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L21
	} else {
		goto L478
	}
L476:
	;
	goto L477
L477:
	;
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1731)+100))
	if v1941 == int32(0) {
		goto L420
	} else {
		goto L480
	}
L478:
	;
	F_add_path(m, l2, v1937)
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L21
	} else {
		goto L479
	}
L479:
	;
	goto L420
L480:
	;
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+8))
	v1945 = F_create_group_path(m, l0, l2, v1920, v1944, v1729, v1724)
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L21
	} else {
		goto L481
	}
L481:
	;
	F_add_path(m, l2, v1945)
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L21
	} else {
		goto L482
	}
L482:
	;
	goto L420
L483:
	;
	goto L419
L484:
	;
	goto L413
L485:
	;
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v1323)+32))
	if v2017 == int32(0) {
		goto L407
	} else {
		goto L486
	}
L486:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v2017)+4))
	if v2020 <= int32(0) {
		goto L407
	} else {
		goto L487
	}
L487:
	;
	v2044 = int32(0)
	goto L488
L488:
	;
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v2017)+12))
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v2052+v2044<<(uint(int32(2))%32))))
	v2057 = F_get_useful_group_keys_orderings(m, l0, v2056)
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L21
	} else {
		goto L491
	}
L489:
	;
	goto L407
L490:
	;
	v2257 = v2044 + int32(1)
	v2258 = *(*int32)(unsafe.Add(mBase, uint32(v2017)+4))
	if v2257 < v2258 {
		v2044 = v2257
		goto L488
	} else {
		goto L555
	}
L491:
	;
	if v2057 == int32(0) {
		goto L490
	} else {
		goto L492
	}
L492:
	;
	v2061 = int32(0)
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(v2057)+4))
	if v2062 <= v2061 {
		goto L490
	} else {
		goto L493
	}
L493:
	;
	v2072 = v2061
	goto L494
L494:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v1323)+48))
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v2057)+12))
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v2094+v2072<<(uint(int32(2))%32))))
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v2098)+4))
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+64))
	v2102 = v31 + int32(16)
	if v2099 == v2100 {
		goto L500
	} else {
		goto L501
	}
L495:
	;
	goto L490
L496:
	;
	v2225 = v2072 + int32(1)
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v2057)+4))
	if v2225 < v2226 {
		v2072 = v2225
		goto L494
	} else {
		goto L554
	}
L497:
	;
	v2204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1731)+36)))
	if v2204 == int32(1) {
		goto L547
	} else {
		goto L548
	}
L498:
	;
	if v2180 != 0 {
		goto L530
	} else {
		goto L531
	}
L499:
	;
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v2099)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2102))) = v2168
	v2180 = int32(1)
	goto L498
L500:
	;
	if v2099 != 0 {
		goto L499
	} else {
		goto L503
	}
L501:
	;
	goto L502
L502:
	;
	if v2099 == int32(0) {
		goto L504
	} else {
		goto L505
	}
L503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2102))) = int32(0)
	v2180 = int32(1)
	goto L498
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2102))) = int32(0)
	v2180 = int32(1)
	goto L498
L505:
	;
	goto L506
L506:
	;
	if v2100 == int32(0) {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v2120 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2102))) = v2120
	v2180 = v2120
	goto L498
L508:
	;
	goto L509
L509:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v2100)+4))
	v2124 = int32(0)
	if v2124 < v2123 {
		goto L510
	} else {
		goto L511
	}
L510:
	;
	v2127 = v2123
	goto L512
L511:
	;
	v2127 = v2124
	goto L512
L512:
	;
	v2128 = *(*int32)(unsafe.Add(mBase, uint32(v2099)+4))
	v2132 = int32(0)
	goto L513
L513:
	;
	if v2132 < v2128 {
		goto L515
	} else {
		goto L516
	}
L515:
	;
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v2099)+12))
	v2144 = v2140 + v2132<<(uint(int32(2))%32)
	goto L517
L516:
	;
	v2144 = int32(0)
	goto L517
L517:
	;
	if v2132 == v2127 {
		goto L518
	} else {
		goto L519
	}
L518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2102))) = v2127
	v2180 = base.B2i32(v2144 == int32(0))
	goto L498
L519:
	;
	goto L520
L520:
	;
	v2150 = base.B2i32(v2144 == int32(0))
	if v2144 == int32(0) {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2102))) = v2132
	v2180 = v2150
	goto L498
L522:
	;
	goto L523
L523:
	;
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v2100)+12))
	v2157 = v2154 + v2132<<(uint(int32(2))%32)
	if v2157 == int32(0) {
		goto L524
	} else {
		goto L525
	}
L524:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2102))) = v2132
	v2180 = v2150
	goto L498
L525:
	;
	goto L526
L526:
	;
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v2144)))
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v2157)))
	if v2161 != v2162 {
		goto L527
	} else {
		goto L528
	}
L527:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2102))) = v2132
	v2180 = int32(0)
	goto L498
L528:
	;
	v2132 = v2132 + int32(1)
	goto L513
L530:
	;
	v2202 = v2056
	goto L497
L531:
	;
	goto L532
L532:
	;
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	if v2056 != v2093 {
		goto L535
	} else {
		goto L536
	}
L533:
	;
	if v2199 == int32(0) {
		goto L496
	} else {
		goto L546
	}
L534:
	;
	v2197 = F_create_incremental_sort_path(m, l0, l2, v2056, v2099, v2181, float64(-1))
	mBase = m.M
	v2198 = m.ExcPending
	if v2198 != 0 {
		goto L21
	} else {
		goto L545
	}
L535:
	;
	if v2181 == int32(0) {
		goto L496
	} else {
		goto L538
	}
L536:
	;
	goto L537
L537:
	;
	if v2181 != 0 {
		goto L540
	} else {
		goto L541
	}
L538:
	;
	v2186 = int32(*(*uint8)(unsafe.Add(mBase, _consts[639])))
	if v2186 == int32(0) {
		goto L496
	} else {
		goto L539
	}
L539:
	;
	goto L534
L540:
	;
	v2190 = int32(*(*uint8)(unsafe.Add(mBase, _consts[639])))
	if v2190&int32(1) != 0 {
		goto L534
	} else {
		goto L543
	}
L541:
	;
	goto L542
L542:
	;
	v2194 = F_create_sort_path(m, l2, v2056, v2099, float64(-1))
	mBase = m.M
	v2195 = m.ExcPending
	if v2195 != 0 {
		goto L21
	} else {
		goto L544
	}
L543:
	;
	goto L542
L544:
	;
	v2199 = v2194
	goto L533
L545:
	;
	v2199 = v2197
	goto L533
L546:
	;
	v2202 = v2199
	goto L497
L547:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v2208 = *(*int32)(unsafe.Add(mBase, uint32(v1731)+100))
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(v2098)+8))
	v2213 = F_create_agg_path(m, l0, l2, v2202, v2207, base.B2i32(v2208 != int32(0)), int32(9), v2212, v1729, v1721, v1724)
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L21
	} else {
		goto L550
	}
L548:
	;
	goto L549
L549:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v2098)+8))
	v2218 = F_create_group_path(m, l0, l2, v2202, v2217, v1729, v1724)
	mBase = m.M
	v2219 = m.ExcPending
	if v2219 != 0 {
		goto L21
	} else {
		goto L552
	}
L550:
	;
	F_add_path(m, l2, v2213)
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L21
	} else {
		goto L551
	}
L551:
	;
	goto L496
L552:
	;
	F_add_path(m, l2, v2218)
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		goto L21
	} else {
		goto L553
	}
L553:
	;
	goto L496
L554:
	;
	goto L495
L555:
	;
	goto L489
L556:
	;
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v2317 != 0 {
		goto L569
	} else {
		goto L570
	}
L557:
	;
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v1731)+108))
	if v2290 != 0 {
		goto L559
	} else {
		goto L560
	}
L558:
	;
	if v1323 == int32(0) {
		goto L556
	} else {
		goto L565
	}
L559:
	;
	F_consider_groupingsets_paths(m, l0, l2, v1730, int32(0), int32(1), l4, l3, v1724)
	mBase = m.M
	v2294 = m.ExcPending
	if v2294 != 0 {
		goto L21
	} else {
		goto L562
	}
L560:
	;
	goto L561
L561:
	;
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v2298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v2299 = F_create_agg_path(m, l0, l2, v1730, v2295, int32(2), int32(0), v2298, v1729, l3, v1724)
	mBase = m.M
	v2300 = m.ExcPending
	if v2300 != 0 {
		goto L21
	} else {
		goto L563
	}
L562:
	;
	goto L558
L563:
	;
	F_add_path(m, l2, v2299)
	mBase = m.M
	v2302 = m.ExcPending
	if v2302 != 0 {
		goto L21
	} else {
		goto L564
	}
L564:
	;
	goto L558
L565:
	;
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(v1323)+32))
	if v2305 == int32(0) {
		goto L556
	} else {
		goto L566
	}
L566:
	;
	v2308 = *(*int32)(unsafe.Add(mBase, uint32(v1323)+48))
	v2309 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v2313 = F_create_agg_path(m, l0, l2, v2308, v2309, int32(2), int32(9), v2312, v1729, v1721, v1724)
	mBase = m.M
	v2314 = m.ExcPending
	if v2314 != 0 {
		goto L21
	} else {
		goto L567
	}
L567:
	;
	F_add_path(m, l2, v2313)
	mBase = m.M
	v2316 = m.ExcPending
	if v2316 != 0 {
		goto L21
	} else {
		goto L568
	}
L568:
	;
	goto L556
L569:
	;
	F_gather_grouping_paths(m, l0, l2)
	mBase = m.M
	v2319 = m.ExcPending
	if v2319 != 0 {
		goto L21
	} else {
		goto L572
	}
L570:
	;
	goto L571
L571:
	;
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v2320 == int32(0) {
		goto L394
	} else {
		goto L573
	}
L572:
	;
	goto L571
L573:
	;
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(l2)+168))
	if v2323 == int32(0) {
		goto L574
	} else {
		goto L575
	}
L574:
	;
	v2334 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	if v2334 == int32(0) {
		goto L395
	} else {
		goto L578
	}
L575:
	;
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(v2323)+36))
	if v2326 == int32(0) {
		goto L574
	} else {
		goto L576
	}
L576:
	;
	m.T0[v2326].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, int32(2), l1, l2, l5)
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L21
	} else {
		goto L577
	}
L577:
	;
	goto L574
L578:
	;
	m.T0[v2334].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, int32(2), l1, l2, l5)
	mBase = m.M
	v2339 = m.ExcPending
	if v2339 != 0 {
		goto L21
	} else {
		goto L579
	}
L579:
	;
	goto L395
L580:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2377 = m.ExcPending
	if v2377 != 0 {
		goto L21
	} else {
		goto L581
	}
L581:
	;
	F_errmsg(m, int32(534319), int32(0))
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L21
	} else {
		goto L582
	}
L582:
	;
	F_errdetail(m, int32(653709), int32(0))
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L21
	} else {
		goto L583
	}
L583:
	;
	F_errfinish(m, int32(518349), int32(4145), int32(164488))
	mBase = m.M
	v2390 = m.ExcPending
	if v2390 != 0 {
		goto L21
	} else {
		goto L584
	}
L584:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_create_scan_plan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
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
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
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
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
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
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v459 int32
	_ = v459
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
	var v474 int32
	_ = v474
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
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v492 float64
	_ = v492
	var v494 float64
	_ = v494
	var v496 float64
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
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
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v545 int32
	_ = v545
	var v547 float64
	_ = v547
	var v549 float64
	_ = v549
	var v551 float64
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v836 int32
	_ = v836
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v896 int32
	_ = v896
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v951 int32
	_ = v951
	var v953 float64
	_ = v953
	var v955 float64
	_ = v955
	var v957 float64
	_ = v957
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
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
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
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
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1086 int32
	_ = v1086
	var v1088 float64
	_ = v1088
	var v1090 float64
	_ = v1090
	var v1092 float64
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
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
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
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
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1216 int32
	_ = v1216
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1267 int32
	_ = v1267
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1301 int32
	_ = v1301
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1377 int32
	_ = v1377
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1425 int32
	_ = v1425
	var v1427 float64
	_ = v1427
	var v1429 float64
	_ = v1429
	var v1431 float64
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
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
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1482 int32
	_ = v1482
	var v1484 float64
	_ = v1484
	var v1486 float64
	_ = v1486
	var v1488 float64
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1537 int32
	_ = v1537
	var v1539 float64
	_ = v1539
	var v1541 float64
	_ = v1541
	var v1543 float64
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1592 int32
	_ = v1592
	var v1594 float64
	_ = v1594
	var v1596 float64
	_ = v1596
	var v1598 float64
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1644 int32
	_ = v1644
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1655 int32
	_ = v1655
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1674 int32
	_ = v1674
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1707 int32
	_ = v1707
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1720 int32
	_ = v1720
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1748 int32
	_ = v1748
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1769 int32
	_ = v1769
	var v1772 int32
	_ = v1772
	var v1777 int32
	_ = v1777
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1823 int32
	_ = v1823
	var v1825 float64
	_ = v1825
	var v1827 float64
	_ = v1827
	var v1829 float64
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1875 int32
	_ = v1875
	var v1877 float64
	_ = v1877
	var v1879 float64
	_ = v1879
	var v1881 float64
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1910 int32
	_ = v1910
	var v1912 float64
	_ = v1912
	var v1914 float64
	_ = v1914
	var v1916 float64
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1975 int32
	_ = v1975
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
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
	var v1993 int32
	_ = v1993
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v2005 int32
	_ = v2005
	var v2007 float64
	_ = v2007
	var v2009 float64
	_ = v2009
	var v2011 float64
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2054 float64
	_ = v2054
	var v2056 float64
	_ = v2056
	var v2058 float64
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2065 int32
	_ = v2065
	var v2067 int32
	_ = v2067
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
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
	var v2105 int32
	_ = v2105
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2126 int32
	_ = v2126
	var v2140 int32
	_ = v2140
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
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
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2196 int32
	_ = v2196
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2226 int32
	_ = v2226
	var v2230 int32
	_ = v2230
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2248 int32
	_ = v2248
	var v2253 int32
	_ = v2253
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2267 int32
	_ = v2267
	var v2272 int32
	_ = v2272
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2283 int32
	_ = v2283
	var v2288 int32
	_ = v2288
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2315 int32
	_ = v2315
	var v2320 int32
	_ = v2320
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2331 int32
	_ = v2331
	var v2336 int32
	_ = v2336
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2347 int32
	_ = v2347
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2377 float64
	_ = v2377
	var v2379 float64
	_ = v2379
	var v2381 float64
	_ = v2381
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2386 int32
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2393 int32
	_ = v2393
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2415 int32
	_ = v2415
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2436 int32
	_ = v2436
	var v2437 int32
	_ = v2437
	var v2447 int32
	_ = v2447
	var v2449 float64
	_ = v2449
	var v2451 float64
	_ = v2451
	var v2453 float64
	_ = v2453
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2460 int32
	_ = v2460
	var v2465 int32
	_ = v2465
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(176)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v22-int32(341)) <= base.Ui32(int32(1)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v34 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v32 = v27 + int32(96)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v32 = v21 + int32(184)
	goto L1
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v36 = F_list_concat_copy(m, v33, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v40 = v33
	goto L7
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	switch v41 - int32(1) {
	case 0, 2:
		goto L14
	default:
		goto L13
	}
L8:
	;
	return int32(0)
L9:
	;
	v40 = v36
	goto L7
L10:
	;
	v67 = int32(0)
	if v66 != 0 {
		goto L23
	} else {
		goto L24
	}
L11:
	;
	v61 = F_order_qual_clauses(m, l0, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L8
	} else {
		goto L20
	}
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v54 == int32(354) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+319)))
	if v48 == int32(1) {
		v60 = v40
		goto L11
	} else {
		goto L16
	}
L14:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+319)))
	if v44 == int32(1) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v66 = int32(0)
	goto L10
L16:
	;
	v66 = int32(0)
	goto L10
L17:
	;
	v57 = int32(76)
	goto L19
L18:
	;
	v57 = int32(80)
	goto L19
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1+v57)))
	v60 = v59
	goto L11
L20:
	;
	v64 = F_extract_actual_clauses(m, v61, int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v66 = v64
	goto L10
L22:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v466 - int32(331) {
	case 0:
		goto L139
	default:
		goto L135
	case 8:
		goto L151
	case 9:
		goto L150
	case 10:
		goto L134
	case 11:
		goto L149
	case 13:
		goto L148
	case 14:
		goto L147
	case 15:
		goto L146
	case 16:
		goto L145
	case 17:
		goto L144
	case 18:
		goto L142
	case 19:
		goto L143
	case 20:
		goto L141
	case 21:
		goto L140
	case 22:
		goto L138
	case 23:
		goto L137
	case 24:
		goto L136
	}
L23:
	;
	v69 = v67
	goto L25
L24:
	;
	v69 = l2
	goto L25
L25:
	;
	if v69 == int32(8) {
		v459 = v67
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v72 = F_use_physical_tlist(m, l0, l1, v69)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L8
	} else {
		goto L29
	}
L27:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_apply_pathtarget_labeling_to_tlist(m, v440, v447)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L8
	} else {
		goto L125
	}
L28:
	;
	if v69&int32(4) == int32(0) {
		v459 = v289
		goto L22
	} else {
		goto L124
	}
L29:
	;
	if v72 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v74 == int32(342) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v374 == int32(0) {
		v459 = v67
		goto L22
	} else {
		goto L110
	}
L33:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+92))
	v79 = F_copyObjectImpl(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v83 = m.G0
	v85 = v83 - int32(16)
	m.G0 = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v88 != 0 {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	if v69&int32(4) != 0 {
		v440 = v79
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v459 = v79
	goto L22
L38:
	;
	if v289 != 0 {
		goto L28
	} else {
		goto L95
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L8
	} else {
		goto L92
	}
L40:
	;
	m.G0 = v85 + int32(16)
	goto L38
L41:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	if base.Ui32(int32(6)) <= base.Ui32(v102-int32(3)) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v100 = v88 + v87<<(uint(int32(2))%32)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+52))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v100 = v94 + v87<<(uint(int32(2))%32) - int32(4)
	goto L41
L45:
	;
	switch v102 {
	case 0:
		goto L49
	case 1:
		goto L48
	default:
		goto L39
	}
L46:
	;
	goto L47
L47:
	;
	v233 = int32(0)
	F_expandRTE(m, v101, v87, v233, v233, int32(-1), int32(1), v233, v85+int32(12))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L8
	} else {
		goto L81
	}
L48:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v101)+36))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+76))
	if v189 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L49:
	;
	v107 = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	v110 = F_table_open(m, v108, v107)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)+48))
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(v112)+120)))
	if int32(0) < v113 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v120 = v107
	v125 = int32(1)
	goto L54
L52:
	;
	v172 = v107
	goto L53
L53:
	;
	F_sequence_close(m, v110, int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L8
	} else {
		goto L68
	}
L54:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v140 = v133 + v134<<(uint(int32(4))%32) + v125*int32(100)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+11)))
	if v141 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v172 = v165
	goto L53
L56:
	;
	v142 = int32(0)
	F_sequence_close(m, v110, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L8
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v147 = v140 - int32(80)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+88)))
	if v148 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v289 = v142
	goto L40
L60:
	;
	v149 = int32(0)
	F_sequence_close(m, v110, v149)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L8
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v154 = base.I32_extend16_s(v125)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v147)+68))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v147)+76))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v147)+96))
	v159 = F_makeVar(m, v87, v154, v155, v156, v157, int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L8
	} else {
		goto L64
	}
L63:
	;
	v289 = v149
	goto L40
L64:
	;
	v161 = int32(0)
	v163 = F_makeTargetEntry(m, v159, v154, v161, v161)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L8
	} else {
		goto L65
	}
L65:
	;
	v165 = F_lappend(m, v120, v163)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	if v125 != v113 {
		v120 = v165
		v125 = v125 + int32(1)
		goto L54
	} else {
		goto L67
	}
L67:
	;
	goto L55
L68:
	;
	v289 = v172
	goto L40
L69:
	;
	v289 = int32(0)
	goto L40
L70:
	;
	goto L71
L71:
	;
	v193 = int32(0)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	if v194 <= v193 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v289 = int32(0)
	goto L40
L73:
	;
	goto L74
L74:
	;
	v202 = int32(0)
	v207 = v193
	goto L75
L75:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v215+v207<<(uint(int32(2))%32))))
	v220 = F_makeVarFromTargetEntry(m, v87, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L77
	}
L76:
	;
	v289 = v227
	goto L40
L77:
	;
	v222 = int32(*(*int16)(unsafe.Add(mBase, uint32(v219)+8)))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+26)))
	v225 = F_makeTargetEntry(m, v220, v222, int32(0), v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	v227 = F_lappend(m, v202, v225)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v230 = v207 + int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	if v230 < v231 {
		v202 = v227
		v207 = v230
		goto L75
	} else {
		goto L80
	}
L80:
	;
	goto L76
L81:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	if v243 == int32(0) {
		v289 = v233
		goto L40
	} else {
		goto L82
	}
L82:
	;
	v246 = int32(0)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if v247 <= v246 {
		v289 = v233
		goto L40
	} else {
		goto L83
	}
L83:
	;
	v253 = v233
	v255 = v246
	goto L84
L84:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v266+v255<<(uint(int32(2))%32))))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	if v271 != int32(6) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v289 = v280
	goto L40
L86:
	;
	v289 = int32(0)
	goto L40
L87:
	;
	goto L88
L88:
	;
	v275 = int32(*(*int16)(unsafe.Add(mBase, uint32(v270)+8)))
	v276 = int32(0)
	v278 = F_makeTargetEntry(m, v270, v275, v276, v276)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	v280 = F_lappend(m, v253, v278)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	v283 = v255 + int32(1)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if v283 < v284 {
		v253 = v280
		v255 = v283
		goto L84
	} else {
		goto L91
	}
L91:
	;
	goto L85
L92:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v309
	F_errmsg_internal(m, int32(79391), v85)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(515923), int32(1903), int32(79418))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	if v320 == int32(0) {
		v459 = v67
		goto L22
	} else {
		goto L96
	}
L96:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	if v324 <= int32(0) {
		v459 = v67
		goto L22
	} else {
		goto L97
	}
L97:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v319)+8))
	v333 = int32(1)
	v334 = v4
	v339 = v67
	goto L98
L98:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v320)+12))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v346+v334<<(uint(int32(2))%32))))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v351 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v459 = v367
	goto L22
L100:
	;
	v352 = F_replace_nestloop_params_mutator(m, v350, l0)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L8
	} else {
		goto L103
	}
L101:
	;
	v354 = v350
	goto L102
L102:
	;
	v356 = int32(0)
	v358 = F_makeTargetEntry(m, v354, base.I32_extend16_s(v333), v356, v356)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L8
	} else {
		goto L104
	}
L103:
	;
	v354 = v352
	goto L102
L104:
	;
	if v327 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v327-int32(4)+v333<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v358)+16)) = v363
	goto L107
L106:
	;
	goto L107
L107:
	;
	v367 = F_lappend(m, v339, v358)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	v370 = v334 + int32(1)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	if v370 < v371 {
		v333 = v333 + int32(1)
		v334 = v370
		v339 = v367
		goto L98
	} else {
		goto L109
	}
L109:
	;
	goto L99
L110:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if v378 <= int32(0) {
		v459 = v67
		goto L22
	} else {
		goto L111
	}
L111:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	v387 = int32(1)
	v388 = v4
	v393 = v67
	goto L112
L112:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v374)+12))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v400+v388<<(uint(int32(2))%32))))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v405 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v459 = v421
	goto L22
L114:
	;
	v406 = F_replace_nestloop_params_mutator(m, v404, l0)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L8
	} else {
		goto L117
	}
L115:
	;
	v408 = v404
	goto L116
L116:
	;
	v410 = int32(0)
	v412 = F_makeTargetEntry(m, v408, base.I32_extend16_s(v387), v410, v410)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L8
	} else {
		goto L118
	}
L117:
	;
	v408 = v406
	goto L116
L118:
	;
	if v381 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v381-int32(4)+v387<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v412)+16)) = v417
	goto L121
L120:
	;
	goto L121
L121:
	;
	v421 = F_lappend(m, v393, v412)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L8
	} else {
		goto L122
	}
L122:
	;
	v424 = v388 + int32(1)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if v424 < v425 {
		v387 = v387 + int32(1)
		v388 = v424
		v393 = v421
		goto L112
	} else {
		goto L123
	}
L123:
	;
	goto L113
L124:
	;
	v440 = v289
	goto L27
L125:
	;
	v459 = v440
	goto L22
L126:
	;
	if v66 != 0 {
		goto L599
	} else {
		goto L600
	}
L127:
	;
	v2420 = F_order_qual_clauses(m, l0, v2415)
	mBase = m.M
	v2421 = m.ExcPending
	if v2421 != 0 {
		goto L8
	} else {
		goto L590
	}
L128:
	;
	v2369 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v2370 = m.ExcPending
	if v2370 != 0 {
		goto L8
	} else {
		goto L585
	}
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		goto L8
	} else {
		goto L582
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2324 = m.ExcPending
	if v2324 != 0 {
		goto L8
	} else {
		goto L579
	}
L131:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
		goto L8
	} else {
		goto L576
	}
L132:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L8
	} else {
		goto L573
	}
L133:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L8
	} else {
		goto L570
	}
L134:
	;
	v2255 = F_create_indexscan_plan(m, l0, l1, v459, v40, int32(0))
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L8
	} else {
		goto L569
	}
L135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L8
	} else {
		goto L566
	}
L136:
	;
	v2201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v2202 != 0 {
		goto L557
	} else {
		goto L558
	}
L137:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+68))
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v2023 != 0 {
		goto L506
	} else {
		goto L507
	}
L138:
	;
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+68))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1927 != 0 {
		goto L486
	} else {
		goto L487
	}
L139:
	;
	v1890 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L8
	} else {
		goto L478
	}
L140:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1838)+68))
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1840 != 0 {
		goto L468
	} else {
		goto L469
	}
L141:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1607)+68))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1609 != 0 {
		goto L417
	} else {
		goto L418
	}
L142:
	;
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+68))
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1554 != 0 {
		goto L405
	} else {
		goto L406
	}
L143:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1497)+68))
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1499 != 0 {
		goto L393
	} else {
		goto L394
	}
L144:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+68))
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1442 != 0 {
		goto L381
	} else {
		goto L382
	}
L145:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+68))
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+140))
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1105 = F_create_plan(m, v1103, v1104)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L8
	} else {
		goto L309
	}
L146:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v967)+68))
	v969 = int32(0)
	if v40 == v969 {
		v1048 = v969
		goto L277
	} else {
		goto L278
	}
L147:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v736)+68))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v738 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L148:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)+68))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v572 = F_create_bitmap_subplan(m, l0, v565, v19+int32(172), v19+int32(168), v19+int32(164))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L8
	} else {
		goto L172
	}
L149:
	;
	v561 = F_create_indexscan_plan(m, l0, l1, v459, v40, int32(1))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L8
	} else {
		goto L171
	}
L150:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v505)+68))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v507 != 0 {
		goto L160
	} else {
		goto L161
	}
L151:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+68))
	v471 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L8
	} else {
		goto L152
	}
L152:
	;
	v474 = F_extract_actual_clauses(m, v471, int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L8
	} else {
		goto L153
	}
L153:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v476 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v477 = F_replace_nestloop_params_mutator(m, v474, l0)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L8
	} else {
		goto L157
	}
L155:
	;
	v479 = v474
	goto L156
L156:
	;
	v481 = F_palloc0(m, int32(80))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L8
	} else {
		goto L158
	}
L157:
	;
	v479 = v477
	goto L156
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v481)+72)) = v470
	*(*int64)(unsafe.Add(mBase, uint32(v481)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v481)+48)) = v479
	*(*int32)(unsafe.Add(mBase, uint32(v481)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v481))) = int32(339)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v481)+4)) = v490
	v492 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v481)+8)) = v492
	v494 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v481)+16)) = v494
	v496 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v481)+24)) = v496
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v481)+32)) = v499
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v481)+36)) = uint8(v501)
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v481)+37)) = uint8(v503)
	v2465 = v481
	goto L126
L159:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v520)+32))
	v522 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L8
	} else {
		goto L163
	}
L160:
	;
	v519 = v507 + v506<<(uint(int32(2))%32)
	goto L159
L161:
	;
	goto L162
L162:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)+52))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)+12))
	v519 = v513 + v506<<(uint(int32(2))%32) - int32(4)
	goto L159
L163:
	;
	v525 = F_extract_actual_clauses(m, v522, int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L8
	} else {
		goto L164
	}
L164:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v527 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v528 = F_replace_nestloop_params_mutator(m, v525, l0)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L8
	} else {
		goto L168
	}
L166:
	;
	v532 = v521
	v533 = v525
	goto L167
L167:
	;
	v535 = F_palloc0(m, int32(88))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L8
	} else {
		goto L170
	}
L168:
	;
	v530 = F_replace_nestloop_params_mutator(m, v521, l0)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L8
	} else {
		goto L169
	}
L169:
	;
	v532 = v530
	v533 = v528
	goto L167
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v535)+80)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v535)+72)) = v506
	*(*int64)(unsafe.Add(mBase, uint32(v535)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v535)+48)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v535)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v535))) = int32(340)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v535)+4)) = v545
	v547 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v535)+8)) = v547
	v549 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v535)+16)) = v549
	v551 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v535)+24)) = v551
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v535)+32)) = v554
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v535)+36)) = uint8(v556)
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v535)+37)) = uint8(v558)
	v2465 = v535
	goto L126
L171:
	;
	v2465 = v561
	goto L126
L172:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v574 == int32(1) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v580 = v572
	goto L177
L174:
	;
	goto L175
L175:
	;
	v639 = int32(0)
	if v40 == v639 {
		v2415 = v639
		goto L127
	} else {
		goto L185
	}
L176:
	;
	v621 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v580)+84)) = uint8(v621)
	goto L175
L177:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v580)))
	switch v594 - int32(337) {
	case 0:
		v600 = int32(72)
		goto L180
	case 1:
		goto L181
	default:
		goto L179
	case 6:
		goto L176
	}
L178:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L8
	} else {
		goto L182
	}
L179:
	;
	goto L178
L180:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v600+v580)))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v602)+12))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	v580 = v604
	goto L177
L181:
	;
	v597 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v580)+72)) = uint8(v597)
	v600 = int32(76)
	goto L180
L182:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v580)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v609
	F_errmsg_internal(m, int32(507592), v19+int32(16))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L8
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(519788), int32(5625), int32(471826))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L8
	} else {
		goto L184
	}
L184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L185:
	;
	v642 = int32(0)
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v643 <= v642 {
		v2415 = v639
		goto L127
	} else {
		goto L186
	}
L186:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v19)+164))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v19)+168))
	v651 = v642
	v659 = v639
	goto L187
L187:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v664+v651<<(uint(int32(2))%32))))
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668)+10)))
	if v669 != 0 {
		v731 = v659
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v2415 = v731
	goto L127
L189:
	;
	v733 = v651 + int32(1)
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v733 < v734 {
		v651 = v733
		v659 = v731
		goto L187
	} else {
		goto L218
	}
L190:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v668)+4))
	v671 = F_list_member(m, v647, v670)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L8
	} else {
		goto L191
	}
L191:
	;
	if v671 != 0 {
		v731 = v659
		goto L189
	} else {
		goto L192
	}
L192:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v668)+60))
	if v673 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v674 = int32(0)
	if v646 == v674 {
		goto L197
	} else {
		goto L198
	}
L194:
	;
	goto L195
L195:
	;
	v713 = F_contain_mutable_functions(m, v670)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L8
	} else {
		goto L210
	}
L196:
	;
	if v712 != 0 {
		v731 = v659
		goto L189
	} else {
		goto L209
	}
L197:
	;
	v712 = int32(0)
	goto L196
L198:
	;
	goto L199
L199:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v646)+4))
	if v680 <= int32(0) {
		v705 = v674
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v712 = v705
	goto L196
L201:
	;
	v683 = int32(0)
	if v683 < v680 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v686 = v680
	goto L204
L203:
	;
	v686 = v683
	goto L204
L204:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v646)+12))
	v689 = int32(0)
	goto L205
L205:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v687+v689<<(uint(int32(2))%32))))
	v698 = base.B2i32(v697 == v673)
	if v697 == v673 {
		v705 = v698
		goto L200
	} else {
		goto L207
	}
L206:
	;
	v705 = v698
	goto L200
L207:
	;
	v700 = v689 + int32(1)
	if v700 != v686 {
		v689 = v700
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	goto L195
L210:
	;
	if v713 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v670
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = v670
	v722 = F_list_make1_impl(m, int32(1), v19+int32(24))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L8
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v727 = F_lappend(m, v659, v668)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L8
	} else {
		goto L217
	}
L214:
	;
	v725 = F_predicate_implied_by(m, v722, v647, int32(0))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L8
	} else {
		goto L215
	}
L215:
	;
	if v725 != 0 {
		v731 = v659
		goto L189
	} else {
		goto L216
	}
L216:
	;
	goto L213
L217:
	;
	v731 = v727
	goto L189
L218:
	;
	goto L188
L219:
	;
	v907 = F_order_qual_clauses(m, l0, v896)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L8
	} else {
		goto L262
	}
L220:
	;
	v896 = v40
	goto L219
L221:
	;
	goto L222
L222:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v738)+4))
	if v741 != int32(1) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v896 = v40
	goto L219
L224:
	;
	goto L225
L225:
	;
	if v40 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v896 = int32(0)
	goto L219
L227:
	;
	goto L228
L228:
	;
	v747 = int32(0)
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v748 <= v747 {
		v896 = v747
		goto L219
	} else {
		goto L229
	}
L229:
	;
	v755 = int32(0)
	v757 = v747
	goto L230
L230:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v768+v755<<(uint(int32(2))%32))))
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v772)+10)))
	if v773 != 0 {
		v876 = v757
		goto L232
	} else {
		goto L233
	}
L231:
	;
	v896 = v876
	goto L219
L232:
	;
	v888 = v755 + int32(1)
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v888 < v889 {
		v755 = v888
		v757 = v876
		goto L230
	} else {
		goto L261
	}
L233:
	;
	v774 = int32(0)
	if v738 == v774 {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	if v812 != 0 {
		v876 = v757
		goto L232
	} else {
		goto L247
	}
L235:
	;
	v812 = int32(0)
	goto L234
L236:
	;
	goto L237
L237:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v738)+4))
	if v780 <= int32(0) {
		v805 = v774
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v812 = v805
	goto L234
L239:
	;
	v783 = int32(0)
	if v783 < v780 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v786 = v780
	goto L242
L241:
	;
	v786 = v783
	goto L242
L242:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v738)+12))
	v789 = int32(0)
	goto L243
L243:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v787+v789<<(uint(int32(2))%32))))
	v798 = base.B2i32(v797 == v772)
	if v797 == v772 {
		v805 = v798
		goto L238
	} else {
		goto L245
	}
L244:
	;
	v805 = v798
	goto L238
L245:
	;
	v800 = v789 + int32(1)
	if v800 != v786 {
		v789 = v800
		goto L243
	} else {
		goto L246
	}
L246:
	;
	goto L244
L247:
	;
	v813 = int32(0)
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v772)+60))
	if v814 == v813 {
		v866 = v813
		goto L248
	} else {
		goto L249
	}
L248:
	;
	if v866 != 0 {
		v876 = v757
		goto L232
	} else {
		goto L259
	}
L249:
	;
	if v738 == int32(0) {
		v866 = v813
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v738)+4))
	if v819 <= int32(0) {
		v866 = v813
		goto L248
	} else {
		goto L251
	}
L251:
	;
	v822 = int32(0)
	if v822 < v819 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v825 = v819
	goto L254
L253:
	;
	v825 = v822
	goto L254
L254:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v738)+12))
	v836 = int32(0)
	goto L255
L255:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v826+v836<<(uint(int32(2))%32))))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v847)+60))
	v849 = base.B2i32(v848 == v814)
	if v848 == v814 {
		v866 = v849
		goto L248
	} else {
		goto L257
	}
L256:
	;
	v866 = v849
	goto L248
L257:
	;
	v851 = v836 + int32(1)
	if v851 != v825 {
		v836 = v851
		goto L255
	} else {
		goto L258
	}
L258:
	;
	goto L256
L259:
	;
	v869 = F_lappend(m, v757, v772)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L8
	} else {
		goto L260
	}
L260:
	;
	v876 = v869
	goto L232
L261:
	;
	goto L231
L262:
	;
	v910 = F_extract_actual_clauses(m, v738, int32(0))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L8
	} else {
		goto L263
	}
L263:
	;
	v913 = F_extract_actual_clauses(m, v907, int32(0))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L8
	} else {
		goto L264
	}
L264:
	;
	if v910 == int32(0) {
		v932 = v913
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v933 != 0 {
		goto L271
	} else {
		goto L272
	}
L266:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v910)+4))
	if v917 < int32(2) {
		v932 = v913
		goto L265
	} else {
		goto L267
	}
L267:
	;
	v920 = F_make_orclause(m, v910)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L8
	} else {
		goto L268
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v19)+172)) = v920
	v927 = F_list_make1_impl(m, int32(1), v19+int32(28))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L8
	} else {
		goto L269
	}
L269:
	;
	v929 = F_list_difference(m, v913, v927)
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L8
	} else {
		goto L270
	}
L270:
	;
	v932 = v929
	goto L265
L271:
	;
	v934 = F_replace_nestloop_params_mutator(m, v910, l0)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L8
	} else {
		goto L274
	}
L272:
	;
	v938 = v910
	v939 = v932
	goto L273
L273:
	;
	v941 = F_palloc0(m, int32(88))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L8
	} else {
		goto L276
	}
L274:
	;
	v936 = F_replace_nestloop_params_mutator(m, v932, l0)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L8
	} else {
		goto L275
	}
L275:
	;
	v938 = v934
	v939 = v936
	goto L273
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v941)+80)) = v938
	*(*int32)(unsafe.Add(mBase, uint32(v941)+72)) = v737
	*(*int64)(unsafe.Add(mBase, uint32(v941)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v941)+48)) = v939
	*(*int32)(unsafe.Add(mBase, uint32(v941)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v941))) = int32(345)
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v941)+4)) = v951
	v953 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v941)+8)) = v953
	v955 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v941)+16)) = v955
	v957 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v941)+24)) = v957
	v959 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v959)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v941)+32)) = v960
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v941)+36)) = uint8(v962)
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v941)+37)) = uint8(v964)
	v2465 = v941
	goto L126
L277:
	;
	v1060 = F_order_qual_clauses(m, l0, v1048)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L8
	} else {
		goto L300
	}
L278:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v972 <= int32(0) {
		v1048 = v969
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v979 = int32(0)
	v980 = v969
	goto L280
L280:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v992+v979<<(uint(int32(2))%32))))
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996)+10)))
	if v997 != 0 {
		v1039 = v980
		goto L282
	} else {
		goto L283
	}
L281:
	;
	v1048 = v1039
	goto L277
L282:
	;
	v1041 = v979 + int32(1)
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v1041 < v1042 {
		v979 = v1041
		v980 = v1039
		goto L280
	} else {
		goto L299
	}
L283:
	;
	v998 = int32(0)
	if v966 == v998 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	if v1036 != 0 {
		v1039 = v980
		goto L282
	} else {
		goto L297
	}
L285:
	;
	v1036 = int32(0)
	goto L284
L286:
	;
	goto L287
L287:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v966)+4))
	if v1004 <= int32(0) {
		v1029 = v998
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1036 = v1029
	goto L284
L289:
	;
	v1007 = int32(0)
	if v1007 < v1004 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1010 = v1004
	goto L292
L291:
	;
	v1010 = v1007
	goto L292
L292:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v966)+12))
	v1013 = int32(0)
	goto L293
L293:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1011+v1013<<(uint(int32(2))%32))))
	v1022 = base.B2i32(v1021 == v996)
	if v1021 == v996 {
		v1029 = v1022
		goto L288
	} else {
		goto L295
	}
L294:
	;
	v1029 = v1022
	goto L288
L295:
	;
	v1024 = v1013 + int32(1)
	if v1024 != v1010 {
		v1013 = v1024
		goto L293
	} else {
		goto L296
	}
L296:
	;
	goto L294
L297:
	;
	v1037 = F_lappend(m, v980, v996)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L8
	} else {
		goto L298
	}
L298:
	;
	v1039 = v1037
	goto L282
L299:
	;
	goto L281
L300:
	;
	v1063 = F_extract_actual_clauses(m, v966, int32(0))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L8
	} else {
		goto L301
	}
L301:
	;
	v1066 = F_extract_actual_clauses(m, v1060, int32(0))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L8
	} else {
		goto L302
	}
L302:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1068 != 0 {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1069 = F_replace_nestloop_params_mutator(m, v1063, l0)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L8
	} else {
		goto L306
	}
L304:
	;
	v1073 = v1063
	v1074 = v1066
	goto L305
L305:
	;
	v1076 = F_palloc0(m, int32(88))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L8
	} else {
		goto L308
	}
L306:
	;
	v1071 = F_replace_nestloop_params_mutator(m, v1066, l0)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L8
	} else {
		goto L307
	}
L307:
	;
	v1073 = v1069
	v1074 = v1071
	goto L305
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+80)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+72)) = v968
	*(*int64)(unsafe.Add(mBase, uint32(v1076)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+48)) = v1074
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1076))) = int32(346)
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+4)) = v1086
	v1088 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1076)+8)) = v1088
	v1090 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1076)+16)) = v1090
	v1092 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1076)+24)) = v1092
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+32)) = v1095
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1076)+36)) = uint8(v1097)
	v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1076)+37)) = uint8(v1099)
	v2465 = v1076
	goto L126
L309:
	;
	v1107 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L8
	} else {
		goto L310
	}
L310:
	;
	v1110 = F_extract_actual_clauses(m, v1107, int32(0))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L8
	} else {
		goto L311
	}
L311:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1112 != 0 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1113 = int32(0)
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+144))
	if v1114 == v1113 {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	v1398 = v1110
	goto L314
L314:
	;
	v1413 = F_palloc0(m, int32(88))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L8
	} else {
		goto L379
	}
L315:
	;
	v1394 = F_replace_nestloop_params_mutator(m, v1110, l0)
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L8
	} else {
		goto L378
	}
L316:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+4))
	if v1117 <= int32(0) {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	v1133 = v1113
	goto L320
L318:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L8
	} else {
		goto L375
	}
L319:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L8
	} else {
		goto L372
	}
L320:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+12))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1136+v1133<<(uint(int32(2))%32))))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1140)+4))
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1141)))
	if v1142 != int32(319) {
		goto L325
	} else {
		goto L326
	}
L321:
	;
	goto L315
L322:
	;
	v1349 = v1133 + int32(1)
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+4))
	if v1349 < v1350 {
		v1133 = v1349
		goto L320
	} else {
		goto L371
	}
L323:
	;
	v1319 = F_palloc0(m, int32(12))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L8
	} else {
		goto L368
	}
L324:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L8
	} else {
		goto L365
	}
L325:
	;
	if v1142 != int32(6) {
		goto L324
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	v1191 = F_find_placeholder_info(m, l0, v1141)
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L8
	} else {
		goto L340
	}
L328:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1141)+4))
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	v1149 = F_bms_is_member(m, v1147, v1148)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L8
	} else {
		goto L329
	}
L329:
	;
	if v1149 == int32(0) {
		goto L319
	} else {
		goto L330
	}
L330:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v1153 == int32(0) {
		goto L323
	} else {
		goto L331
	}
L331:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1153)+4))
	if v1156 <= int32(0) {
		goto L323
	} else {
		goto L332
	}
L332:
	;
	v1159 = int32(0)
	if v1159 < v1156 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v1162 = v1156
	goto L335
L334:
	;
	v1162 = v1159
	goto L335
L335:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1140)+8))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1153)+12))
	v1169 = int32(0)
	goto L336
L336:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v1164+v1169<<(uint(int32(2))%32))))
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+4))
	if v1186 == v1163 {
		goto L322
	} else {
		goto L338
	}
L337:
	;
	goto L323
L338:
	;
	v1189 = v1169 + int32(1)
	if v1162 != v1189 {
		v1169 = v1189
		goto L336
	} else {
		goto L339
	}
L339:
	;
	goto L337
L340:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+12))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	v1195 = int32(0)
	if v1193 == v1195 {
		goto L342
	} else {
		goto L343
	}
L341:
	;
	if v1248 == int32(0) {
		goto L318
	} else {
		goto L355
	}
L342:
	;
	v1248 = int32(1)
	goto L341
L343:
	;
	goto L344
L344:
	;
	if v1194 == int32(0) {
		v1239 = v1195
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1248 = v1239
	goto L341
L346:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+4))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1194)+4))
	if v1205 < v1204 {
		v1239 = v1195
		goto L345
	} else {
		goto L347
	}
L347:
	;
	v1207 = int32(1)
	if v1204 <= v1207 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1210 = v1207
	goto L350
L349:
	;
	v1210 = v1204
	goto L350
L350:
	;
	v1211 = int32(8)
	v1216 = int32(0)
	goto L351
L351:
	;
	v1223 = v1216 << (uint(int32(2)) % 32)
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1193+v1211+v1223)))
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1223+(v1194+v1211))))
	v1230 = v1225 & (v1227 ^ int32(-1))
	v1232 = base.B2i32(v1230 == int32(0))
	if v1230 != 0 {
		v1239 = v1232
		goto L345
	} else {
		goto L353
	}
L352:
	;
	v1239 = v1232
	goto L345
L353:
	;
	v1234 = v1216 + int32(1)
	if v1234 != v1210 {
		v1216 = v1234
		goto L351
	} else {
		goto L354
	}
L354:
	;
	goto L352
L355:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v1251 == int32(0) {
		goto L323
	} else {
		goto L356
	}
L356:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1251)+4))
	if v1254 <= int32(0) {
		goto L323
	} else {
		goto L357
	}
L357:
	;
	v1257 = int32(0)
	if v1257 < v1254 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1260 = v1254
	goto L360
L359:
	;
	v1260 = v1257
	goto L360
L360:
	;
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1140)+8))
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1251)+12))
	v1267 = int32(0)
	goto L361
L361:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1262+v1267<<(uint(int32(2))%32))))
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1283)+4))
	if v1284 == v1261 {
		goto L322
	} else {
		goto L363
	}
L362:
	;
	goto L323
L363:
	;
	v1287 = v1267 + int32(1)
	if v1260 != v1287 {
		v1267 = v1287
		goto L361
	} else {
		goto L364
	}
L364:
	;
	goto L362
L365:
	;
	F_errmsg_internal(m, int32(227175), int32(0))
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L8
	} else {
		goto L366
	}
L366:
	;
	F_errfinish(m, int32(519752), int32(597), int32(160335))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L8
	} else {
		goto L367
	}
L367:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1319))) = int32(357)
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1140)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1319)+4)) = v1323
	v1325 = F_copyObjectImpl(m, v1141)
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L8
	} else {
		goto L369
	}
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1319)+8)) = v1325
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	v1329 = F_lappend(m, v1328, v1319)
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L8
	} else {
		goto L370
	}
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+356)) = v1329
	goto L322
L371:
	;
	goto L321
L372:
	;
	F_errmsg_internal(m, int32(15573), int32(0))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L8
	} else {
		goto L373
	}
L373:
	;
	F_errfinish(m, int32(519752), int32(543), int32(160335))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L8
	} else {
		goto L374
	}
L374:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L375:
	;
	F_errmsg_internal(m, int32(15573), int32(0))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L8
	} else {
		goto L376
	}
L376:
	;
	F_errfinish(m, int32(519752), int32(574), int32(160335))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L8
	} else {
		goto L377
	}
L377:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L378:
	;
	v1398 = v1394
	goto L314
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1413)+84)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1413)+80)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v1413)+72)) = v1102
	*(*int64)(unsafe.Add(mBase, uint32(v1413)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1413)+48)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v1413)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1413))) = int32(347)
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1413)+4)) = v1425
	v1427 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1413)+8)) = v1427
	v1429 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1413)+16)) = v1429
	v1431 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1413)+24)) = v1431
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1413)+32)) = v1434
	v1436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1413)+36)) = uint8(v1436)
	v1438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1413)+37)) = uint8(v1438)
	v2465 = v1413
	goto L126
L380:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1454)))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1455)+68))
	v1457 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L8
	} else {
		goto L384
	}
L381:
	;
	v1454 = v1442 + v1441<<(uint(int32(2))%32)
	goto L380
L382:
	;
	goto L383
L383:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+52))
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+12))
	v1454 = v1448 + v1441<<(uint(int32(2))%32) - int32(4)
	goto L380
L384:
	;
	v1460 = F_extract_actual_clauses(m, v1457, int32(0))
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L8
	} else {
		goto L385
	}
L385:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1462 != 0 {
		goto L386
	} else {
		goto L387
	}
L386:
	;
	v1463 = F_replace_nestloop_params_mutator(m, v1460, l0)
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L8
	} else {
		goto L389
	}
L387:
	;
	v1467 = v1456
	v1468 = v1460
	goto L388
L388:
	;
	v1469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1455)+72)))
	v1471 = F_palloc0(m, int32(88))
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L8
	} else {
		goto L391
	}
L389:
	;
	v1465 = F_replace_nestloop_params_mutator(m, v1456, l0)
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L8
	} else {
		goto L390
	}
L390:
	;
	v1467 = v1465
	v1468 = v1463
	goto L388
L391:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1471)+84)) = uint8(v1469)
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+80)) = v1467
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+72)) = v1441
	*(*int64)(unsafe.Add(mBase, uint32(v1471)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+48)) = v1468
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1471))) = int32(348)
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+4)) = v1482
	v1484 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1471)+8)) = v1484
	v1486 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1471)+16)) = v1486
	v1488 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1471)+24)) = v1488
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+32)) = v1491
	v1493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1471)+36)) = uint8(v1493)
	v1495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1471)+37)) = uint8(v1495)
	v2465 = v1471
	goto L126
L392:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1511)))
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1512)+76))
	v1514 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L8
	} else {
		goto L396
	}
L393:
	;
	v1511 = v1499 + v1498<<(uint(int32(2))%32)
	goto L392
L394:
	;
	goto L395
L395:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1503)+52))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1504)+12))
	v1511 = v1505 + v1498<<(uint(int32(2))%32) - int32(4)
	goto L392
L396:
	;
	v1517 = F_extract_actual_clauses(m, v1514, int32(0))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L8
	} else {
		goto L397
	}
L397:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1519 != 0 {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	v1520 = F_replace_nestloop_params_mutator(m, v1517, l0)
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L8
	} else {
		goto L401
	}
L399:
	;
	v1524 = v1513
	v1525 = v1517
	goto L400
L400:
	;
	v1527 = F_palloc0(m, int32(88))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L8
	} else {
		goto L403
	}
L401:
	;
	v1522 = F_replace_nestloop_params_mutator(m, v1513, l0)
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L8
	} else {
		goto L402
	}
L402:
	;
	v1524 = v1522
	v1525 = v1520
	goto L400
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1527)+80)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v1527)+72)) = v1498
	*(*int64)(unsafe.Add(mBase, uint32(v1527)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1527)+48)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v1527)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1527))) = int32(350)
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1527)+4)) = v1537
	v1539 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1527)+8)) = v1539
	v1541 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1527)+16)) = v1541
	v1543 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1527)+24)) = v1543
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1545)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1527)+32)) = v1546
	v1548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1527)+36)) = uint8(v1548)
	v1550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1527)+37)) = uint8(v1550)
	v2465 = v1527
	goto L126
L404:
	;
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1566)))
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1567)+80))
	v1569 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L8
	} else {
		goto L408
	}
L405:
	;
	v1566 = v1554 + v1553<<(uint(int32(2))%32)
	goto L404
L406:
	;
	goto L407
L407:
	;
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1558)+52))
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v1559)+12))
	v1566 = v1560 + v1553<<(uint(int32(2))%32) - int32(4)
	goto L404
L408:
	;
	v1572 = F_extract_actual_clauses(m, v1569, int32(0))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L8
	} else {
		goto L409
	}
L409:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1574 != 0 {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v1575 = F_replace_nestloop_params_mutator(m, v1572, l0)
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L8
	} else {
		goto L413
	}
L411:
	;
	v1579 = v1568
	v1580 = v1572
	goto L412
L412:
	;
	v1582 = F_palloc0(m, int32(88))
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L8
	} else {
		goto L415
	}
L413:
	;
	v1577 = F_replace_nestloop_params_mutator(m, v1568, l0)
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L8
	} else {
		goto L414
	}
L414:
	;
	v1579 = v1577
	v1580 = v1575
	goto L412
L415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1582)+80)) = v1579
	*(*int32)(unsafe.Add(mBase, uint32(v1582)+72)) = v1553
	*(*int64)(unsafe.Add(mBase, uint32(v1582)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1582)+48)) = v1580
	*(*int32)(unsafe.Add(mBase, uint32(v1582)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1582))) = int32(349)
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1582)+4)) = v1592
	v1594 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1582)+8)) = v1594
	v1596 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1582)+16)) = v1596
	v1598 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1582)+24)) = v1598
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1600)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1582)+32)) = v1601
	v1603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1582)+36)) = uint8(v1603)
	v1605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1582)+37)) = uint8(v1605)
	v2465 = v1582
	goto L126
L416:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1621)))
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+88))
	v1626 = l0
	v1627 = v1623
	goto L421
L417:
	;
	v1621 = v1609 + v1608<<(uint(int32(2))%32)
	goto L416
L418:
	;
	goto L419
L419:
	;
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1613)+52))
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v1614)+12))
	v1621 = v1615 + v1608<<(uint(int32(2))%32) - int32(4)
	goto L416
L420:
	;
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+4))
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+48))
	if v1662 == int32(0) {
		goto L429
	} else {
		goto L430
	}
L421:
	;
	if v1627 == int32(0) {
		goto L420
	} else {
		goto L423
	}
L422:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L8
	} else {
		goto L425
	}
L423:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+16))
	if v1644 != 0 {
		v1626 = v1644
		v1627 = v1627 - int32(1)
		goto L421
	} else {
		goto L424
	}
L424:
	;
	goto L422
L425:
	;
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v1649
	F_errmsg_internal(m, int32(756844), v19+int32(96))
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L8
	} else {
		goto L426
	}
L426:
	;
	F_errfinish(m, int32(519788), int32(3912), int32(296007))
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L8
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
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+76))
	if v1754 != 0 {
		goto L447
	} else {
		goto L448
	}
L429:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L8
	} else {
		goto L444
	}
L430:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1662)+4))
	if v1665 <= int32(0) {
		goto L429
	} else {
		goto L431
	}
L431:
	;
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+84))
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1662)+12))
	v1674 = int32(0)
	goto L432
L432:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1669+v1674<<(uint(int32(2))%32))))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+4))
	v1694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1668))))
	v1695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1691))))
	if v1695 == int32(0) {
		v1714 = v1694
		v1715 = v1695
		goto L435
	} else {
		goto L436
	}
L433:
	;
	goto L429
L434:
	;
	if v1715-v1714 == int32(0) {
		goto L428
	} else {
		goto L442
	}
L435:
	;
	goto L434
L436:
	;
	if v1694 != v1695 {
		v1714 = v1694
		v1715 = v1695
		goto L435
	} else {
		goto L437
	}
L437:
	;
	v1699 = v1691
	v1700 = v1668
	goto L438
L438:
	;
	v1703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1700)+1)))
	v1704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1699)+1)))
	if v1704 == int32(0) {
		v1714 = v1703
		v1715 = v1704
		goto L435
	} else {
		goto L440
	}
L439:
	;
	v1714 = v1703
	v1715 = v1704
	goto L435
L440:
	;
	v1707 = int32(1)
	if v1703 == v1704 {
		v1699 = v1699 + v1707
		v1700 = v1700 + v1707
		goto L438
	} else {
		goto L441
	}
L441:
	;
	goto L439
L442:
	;
	v1720 = v1674 + int32(1)
	if v1665 != v1720 {
		v1674 = v1720
		goto L432
	} else {
		goto L443
	}
L443:
	;
	goto L433
L444:
	;
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v1742
	F_errmsg_internal(m, int32(757003), v19+int32(32))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L8
	} else {
		goto L445
	}
L445:
	;
	F_errfinish(m, int32(519788), int32(3930), int32(296007))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L8
	} else {
		goto L446
	}
L446:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L447:
	;
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+4))
	v1757 = v1755
	goto L449
L448:
	;
	v1757 = int32(0)
	goto L449
L449:
	;
	if v1757 <= v1674 {
		goto L133
	} else {
		goto L450
	}
L450:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+12))
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1759+v1674<<(uint(int32(2))%32))))
	if v1763 <= int32(0) {
		goto L132
	} else {
		goto L451
	}
L451:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+72))
	if v1766 == int32(0) {
		goto L131
	} else {
		goto L452
	}
L452:
	;
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(v1766)+4))
	if v1769 <= int32(0) {
		goto L131
	} else {
		goto L453
	}
L453:
	;
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1766)+12))
	v1777 = int32(0)
	goto L454
L454:
	;
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(v1772+v1777<<(uint(int32(2))%32))))
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(v1793)+16))
	if v1763 != v1794 {
		goto L456
	} else {
		goto L457
	}
L455:
	;
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(v1793)+40))
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1799)+12))
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1800)))
	v1802 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L8
	} else {
		goto L460
	}
L456:
	;
	v1797 = v1777 + int32(1)
	if v1797 != v1769 {
		v1777 = v1797
		goto L454
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	goto L455
L459:
	;
	goto L131
L460:
	;
	v1805 = F_extract_actual_clauses(m, v1802, int32(0))
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L8
	} else {
		goto L461
	}
L461:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1807 != 0 {
		goto L462
	} else {
		goto L463
	}
L462:
	;
	v1808 = F_replace_nestloop_params_mutator(m, v1805, l0)
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L8
	} else {
		goto L465
	}
L463:
	;
	v1810 = v1805
	goto L464
L464:
	;
	v1812 = F_palloc0(m, int32(88))
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L8
	} else {
		goto L466
	}
L465:
	;
	v1810 = v1808
	goto L464
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1812)+84)) = v1801
	*(*int32)(unsafe.Add(mBase, uint32(v1812)+80)) = v1763
	*(*int32)(unsafe.Add(mBase, uint32(v1812)+72)) = v1608
	*(*int64)(unsafe.Add(mBase, uint32(v1812)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1812)+48)) = v1810
	*(*int32)(unsafe.Add(mBase, uint32(v1812)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1812))) = int32(351)
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1812)+4)) = v1823
	v1825 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1812)+8)) = v1825
	v1827 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1812)+16)) = v1827
	v1829 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1812)+24)) = v1829
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1831)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1812)+32)) = v1832
	v1834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1812)+36)) = uint8(v1834)
	v1836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1812)+37)) = uint8(v1836)
	v2465 = v1812
	goto L126
L467:
	;
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1852)))
	v1854 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L8
	} else {
		goto L471
	}
L468:
	;
	v1852 = v1840 + v1839<<(uint(int32(2))%32)
	goto L467
L469:
	;
	goto L470
L470:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1844)+52))
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1845)+12))
	v1852 = v1846 + v1839<<(uint(int32(2))%32) - int32(4)
	goto L467
L471:
	;
	v1857 = F_extract_actual_clauses(m, v1854, int32(0))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L8
	} else {
		goto L472
	}
L472:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1859 != 0 {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v1860 = F_replace_nestloop_params_mutator(m, v1857, l0)
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L8
	} else {
		goto L476
	}
L474:
	;
	v1862 = v1857
	goto L475
L475:
	;
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v1853)+108))
	v1865 = F_palloc0(m, int32(88))
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L8
	} else {
		goto L477
	}
L476:
	;
	v1862 = v1860
	goto L475
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1865)+80)) = v1863
	*(*int32)(unsafe.Add(mBase, uint32(v1865)+72)) = v1839
	*(*int64)(unsafe.Add(mBase, uint32(v1865)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1865)+48)) = v1862
	*(*int32)(unsafe.Add(mBase, uint32(v1865)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1865))) = int32(352)
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1865)+4)) = v1875
	v1877 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1865)+8)) = v1877
	v1879 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1865)+16)) = v1879
	v1881 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1865)+24)) = v1881
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1883)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1865)+32)) = v1884
	v1886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1865)+36)) = uint8(v1886)
	v1888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1865)+37)) = uint8(v1888)
	v2465 = v1865
	goto L126
L478:
	;
	v1893 = F_extract_actual_clauses(m, v1890, int32(0))
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L8
	} else {
		goto L479
	}
L479:
	;
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1895 != 0 {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	v1896 = F_replace_nestloop_params_mutator(m, v1893, l0)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L8
	} else {
		goto L483
	}
L481:
	;
	v1898 = v1893
	goto L482
L482:
	;
	v1900 = F_palloc0(m, int32(80))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L8
	} else {
		goto L484
	}
L483:
	;
	v1898 = v1896
	goto L482
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1900)+72)) = v1898
	*(*int32)(unsafe.Add(mBase, uint32(v1900)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1900)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1900)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1900))) = int32(331)
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1900)+4)) = v1910
	v1912 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1900)+8)) = v1912
	v1914 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1900)+16)) = v1914
	v1916 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1900)+24)) = v1916
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v1918)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1900)+32)) = v1919
	v1921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1900)+36)) = uint8(v1921)
	v1923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1900)+37)) = uint8(v1923)
	v2465 = v1900
	goto L126
L485:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1939)))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1940)+88))
	if v1941 == int32(0) {
		goto L130
	} else {
		goto L489
	}
L486:
	;
	v1939 = v1927 + v1926<<(uint(int32(2))%32)
	goto L485
L487:
	;
	goto L488
L488:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1931)+52))
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v1932)+12))
	v1939 = v1933 + v1926<<(uint(int32(2))%32) - int32(4)
	goto L485
L489:
	;
	v1946 = l0
	v1947 = v1941
	goto L491
L490:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(v1946)+344))
	if v1981 < int32(0) {
		goto L129
	} else {
		goto L498
	}
L491:
	;
	v1961 = v1947 - int32(1)
	if v1961 == int32(0) {
		goto L490
	} else {
		goto L493
	}
L492:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L8
	} else {
		goto L495
	}
L493:
	;
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1946)+16))
	if v1964 != 0 {
		v1946 = v1964
		v1947 = v1961
		goto L491
	} else {
		goto L494
	}
L494:
	;
	goto L492
L495:
	;
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v1940)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v1969
	F_errmsg_internal(m, int32(756844), v19+int32(144))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L8
	} else {
		goto L496
	}
L496:
	;
	F_errfinish(m, int32(519788), int32(4083), int32(296027))
	mBase = m.M
	v1980 = m.ExcPending
	if v1980 != 0 {
		goto L8
	} else {
		goto L497
	}
L497:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L498:
	;
	v1984 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L8
	} else {
		goto L499
	}
L499:
	;
	v1987 = F_extract_actual_clauses(m, v1984, int32(0))
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L8
	} else {
		goto L500
	}
L500:
	;
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1989 != 0 {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	v1990 = F_replace_nestloop_params_mutator(m, v1987, l0)
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L8
	} else {
		goto L504
	}
L502:
	;
	v1992 = v1987
	goto L503
L503:
	;
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1946)+344))
	v1995 = F_palloc0(m, int32(88))
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L8
	} else {
		goto L505
	}
L504:
	;
	v1992 = v1990
	goto L503
L505:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1995)+80)) = v1993
	*(*int32)(unsafe.Add(mBase, uint32(v1995)+72)) = v1926
	*(*int64)(unsafe.Add(mBase, uint32(v1995)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1995)+48)) = v1992
	*(*int32)(unsafe.Add(mBase, uint32(v1995)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1995))) = int32(353)
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1995)+4)) = v2005
	v2007 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1995)+8)) = v2007
	v2009 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1995)+16)) = v2009
	v2011 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1995)+24)) = v2011
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v2013)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1995)+32)) = v2014
	v2016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1995)+36)) = uint8(v2016)
	v2018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1995)+37)) = uint8(v2018)
	v2465 = v1995
	goto L126
L506:
	;
	v2025 = F_create_plan_recurse(m, l0, v2023, int32(1))
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L8
	} else {
		goto L509
	}
L507:
	;
	v2027 = int32(0)
	goto L508
L508:
	;
	if v2021 != 0 {
		goto L510
	} else {
		goto L511
	}
L509:
	;
	v2027 = v2025
	goto L508
L510:
	;
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2028 != 0 {
		goto L514
	} else {
		goto L515
	}
L511:
	;
	v2045 = int32(0)
	goto L512
L512:
	;
	v2046 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L8
	} else {
		goto L517
	}
L513:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(v2040)))
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v2041)+16))
	v2045 = v2042
	goto L512
L514:
	;
	v2040 = v2028 + v2021<<(uint(int32(2))%32)
	goto L513
L515:
	;
	goto L516
L516:
	;
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v2032)+52))
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v2033)+12))
	v2040 = v2034 + v2021<<(uint(int32(2))%32) - int32(4)
	goto L513
L517:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+168))
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v2048)+12))
	v2050 = m.T0[v2049].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, v2020, v2045, l1, v459, v2046, v2027)
	mBase = m.M
	v2051 = m.ExcPending
	if v2051 != 0 {
		goto L8
	} else {
		goto L518
	}
L518:
	;
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2050)+4)) = v2052
	v2054 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v2050)+8)) = v2054
	v2056 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v2050)+16)) = v2056
	v2058 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v2050)+24)) = v2058
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(v2060)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2050)+32)) = v2061
	v2063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2050)+36)) = uint8(v2063)
	v2065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2050)+37)) = uint8(v2065)
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v2050)+88)) = v2067
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v2050)+92)) = v2069
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+4))
	if v2071 == int32(4) {
		goto L519
	} else {
		goto L520
	}
L519:
	;
	v2079 = l0 + int32(52)
	goto L521
L520:
	;
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2079 = v2076 + int32(8)
	goto L521
L521:
	;
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v2079)))
	*(*int32)(unsafe.Add(mBase, uint32(v2050)+112)) = v2080
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2083 = F_bms_difference(m, v2080, v2082)
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L8
	} else {
		goto L522
	}
L522:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2050)+116)) = v2083
	v2086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2020)+164)))
	if v2086 == int32(1) {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2090 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2089)+81)) = uint8(v2090)
	goto L525
L524:
	;
	goto L525
L525:
	;
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v2092 != 0 {
		goto L526
	} else {
		goto L527
	}
L526:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v2050)+48))
	v2094 = F_replace_nestloop_params_mutator(m, v2093, l0)
	mBase = m.M
	v2095 = m.ExcPending
	if v2095 != 0 {
		goto L8
	} else {
		goto L529
	}
L527:
	;
	goto L528
L528:
	;
	v2105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2050)+120)) = uint8(v2105)
	if v2021 == v2105 {
		v2465 = v2050
		goto L126
	} else {
		goto L532
	}
L529:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2050)+48)) = v2094
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v2050)+96))
	v2098 = F_replace_nestloop_params_mutator(m, v2097, l0)
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L8
	} else {
		goto L530
	}
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2050)+96)) = v2098
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v2050)+108))
	v2102 = F_replace_nestloop_params_mutator(m, v2101, l0)
	mBase = m.M
	v2103 = m.ExcPending
	if v2103 != 0 {
		goto L8
	} else {
		goto L531
	}
L531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2050)+108)) = v2102
	goto L528
L532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+172)) = int32(0)
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+28))
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+4))
	F_pull_varattnos(m, v2112, v2021, v19+int32(172))
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L8
	} else {
		goto L533
	}
L533:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+184))
	if v2117 == int32(0) {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2172 = F_bms_is_member(m, int32(1), v2171)
	mBase = m.M
	v2173 = m.ExcPending
	if v2173 != 0 {
		goto L8
	} else {
		goto L543
	}
L535:
	;
	v2120 = int32(0)
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v2117)+4))
	if v2121 <= v2120 {
		goto L534
	} else {
		goto L536
	}
L536:
	;
	v2126 = v2120
	goto L537
L537:
	;
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v2117)+12))
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(v2140+v2126<<(uint(int32(2))%32))))
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v2144)+4))
	F_pull_varattnos(m, v2145, v2021, v19+int32(172))
	mBase = m.M
	v2149 = m.ExcPending
	if v2149 != 0 {
		goto L8
	} else {
		goto L539
	}
L538:
	;
	goto L534
L539:
	;
	v2151 = v2126 + int32(1)
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v2117)+4))
	if v2151 < v2152 {
		v2126 = v2151
		goto L537
	} else {
		goto L540
	}
L540:
	;
	goto L538
L541:
	;
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	F_bms_free(m, v2198)
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L8
	} else {
		goto L555
	}
L542:
	;
	v2196 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2050)+120)) = uint8(v2196)
	goto L541
L543:
	;
	if v2172 != 0 {
		goto L542
	} else {
		goto L544
	}
L544:
	;
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2176 = F_bms_is_member(m, int32(2), v2175)
	mBase = m.M
	v2177 = m.ExcPending
	if v2177 != 0 {
		goto L8
	} else {
		goto L545
	}
L545:
	;
	if v2176 != 0 {
		goto L542
	} else {
		goto L546
	}
L546:
	;
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2180 = F_bms_is_member(m, int32(3), v2179)
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L8
	} else {
		goto L547
	}
L547:
	;
	if v2180 != 0 {
		goto L542
	} else {
		goto L548
	}
L548:
	;
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2184 = F_bms_is_member(m, int32(4), v2183)
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L8
	} else {
		goto L549
	}
L549:
	;
	if v2184 != 0 {
		goto L542
	} else {
		goto L550
	}
L550:
	;
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2188 = F_bms_is_member(m, int32(5), v2187)
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
		goto L8
	} else {
		goto L551
	}
L551:
	;
	if v2188 != 0 {
		goto L542
	} else {
		goto L552
	}
L552:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2192 = F_bms_is_member(m, int32(6), v2191)
	mBase = m.M
	v2193 = m.ExcPending
	if v2193 != 0 {
		goto L8
	} else {
		goto L553
	}
L553:
	;
	if v2192 == int32(0) {
		goto L541
	} else {
		goto L554
	}
L554:
	;
	goto L542
L555:
	;
	v2465 = v2050
	goto L126
L556:
	;
	v2212 = int32(0)
	v2213 = v2203
	goto L561
L557:
	;
	v2203 = int32(0)
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(v2202)+4))
	if v2203 < v2204 {
		goto L556
	} else {
		goto L560
	}
L558:
	;
	goto L559
L559:
	;
	v2355 = int32(0)
	goto L128
L560:
	;
	goto L559
L561:
	;
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v2202)+12))
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v2226+v2213<<(uint(int32(2))%32))))
	v2232 = F_create_plan_recurse(m, l0, v2230, int32(1))
	mBase = m.M
	v2233 = m.ExcPending
	if v2233 != 0 {
		goto L8
	} else {
		goto L563
	}
L562:
	;
	v2355 = v2234
	goto L128
L563:
	;
	v2234 = F_lappend(m, v2212, v2232)
	mBase = m.M
	v2235 = m.ExcPending
	if v2235 != 0 {
		goto L8
	} else {
		goto L564
	}
L564:
	;
	v2237 = v2213 + int32(1)
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v2202)+4))
	if v2237 < v2238 {
		v2212 = v2234
		v2213 = v2237
		goto L561
	} else {
		goto L565
	}
L565:
	;
	goto L562
L566:
	;
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v2244
	F_errmsg_internal(m, int32(507592), v19)
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L8
	} else {
		goto L567
	}
L567:
	;
	F_errfinish(m, int32(519788), int32(802), int32(296053))
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L8
	} else {
		goto L568
	}
L568:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L569:
	;
	v2465 = v2255
	goto L126
L570:
	;
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v2261
	F_errmsg_internal(m, int32(756870), v19+int32(80))
	mBase = m.M
	v2267 = m.ExcPending
	if v2267 != 0 {
		goto L8
	} else {
		goto L571
	}
L571:
	;
	F_errfinish(m, int32(519788), int32(3932), int32(296007))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L8
	} else {
		goto L572
	}
L572:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L573:
	;
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v2277
	F_errmsg_internal(m, int32(756936), v19+int32(48))
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L8
	} else {
		goto L574
	}
L574:
	;
	F_errfinish(m, int32(519788), int32(3935), int32(296007))
	mBase = m.M
	v2288 = m.ExcPending
	if v2288 != 0 {
		goto L8
	} else {
		goto L575
	}
L575:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L576:
	;
	v2309 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v2309
	F_errmsg_internal(m, int32(756870), v19-int32(-64))
	mBase = m.M
	v2315 = m.ExcPending
	if v2315 != 0 {
		goto L8
	} else {
		goto L577
	}
L577:
	;
	F_errfinish(m, int32(519788), int32(3943), int32(296007))
	mBase = m.M
	v2320 = m.ExcPending
	if v2320 != 0 {
		goto L8
	} else {
		goto L578
	}
L578:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L579:
	;
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(v1940)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v2325
	F_errmsg_internal(m, int32(756844), v19+int32(112))
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L8
	} else {
		goto L580
	}
L580:
	;
	F_errfinish(m, int32(519788), int32(4076), int32(296027))
	mBase = m.M
	v2336 = m.ExcPending
	if v2336 != 0 {
		goto L8
	} else {
		goto L581
	}
L581:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L582:
	;
	v2341 = *(*int32)(unsafe.Add(mBase, uint32(v1940)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v2341
	F_errmsg_internal(m, int32(756966), v19+int32(128))
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L8
	} else {
		goto L583
	}
L583:
	;
	F_errfinish(m, int32(519788), int32(4086), int32(296027))
	mBase = m.M
	v2352 = m.ExcPending
	if v2352 != 0 {
		goto L8
	} else {
		goto L584
	}
L584:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L585:
	;
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v2372 = *(*int32)(unsafe.Add(mBase, uint32(v2371)+4))
	v2373 = m.T0[v2372].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v2201, l1, v459, v2369, v2355)
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L8
	} else {
		goto L586
	}
L586:
	;
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2373)+4)) = v2375
	v2377 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v2373)+8)) = v2377
	v2379 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v2373)+16)) = v2379
	v2381 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v2373)+24)) = v2381
	v2383 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(v2383)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2373)+32)) = v2384
	v2386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2373)+36)) = uint8(v2386)
	v2388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2373)+37)) = uint8(v2388)
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v2390)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2373)+100)) = v2391
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v2393 == int32(0) {
		v2465 = v2373
		goto L126
	} else {
		goto L587
	}
L587:
	;
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v2373)+48))
	v2397 = F_replace_nestloop_params_mutator(m, v2396, l0)
	mBase = m.M
	v2398 = m.ExcPending
	if v2398 != 0 {
		goto L8
	} else {
		goto L588
	}
L588:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2373)+48)) = v2397
	v2400 = *(*int32)(unsafe.Add(mBase, uint32(v2373)+88))
	v2401 = F_replace_nestloop_params_mutator(m, v2400, l0)
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L8
	} else {
		goto L589
	}
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2373)+88)) = v2401
	v2465 = v2373
	goto L126
L590:
	;
	v2423 = F_extract_actual_clauses(m, v2420, int32(0))
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L8
	} else {
		goto L591
	}
L591:
	;
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2426 = F_list_difference_ptr(m, v2425, v2423)
	mBase = m.M
	v2427 = m.ExcPending
	if v2427 != 0 {
		goto L8
	} else {
		goto L592
	}
L592:
	;
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v2428 != 0 {
		goto L593
	} else {
		goto L594
	}
L593:
	;
	v2429 = F_replace_nestloop_params_mutator(m, v2423, l0)
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L8
	} else {
		goto L596
	}
L594:
	;
	v2433 = v2423
	v2434 = v2426
	goto L595
L595:
	;
	v2436 = F_palloc0(m, int32(88))
	mBase = m.M
	v2437 = m.ExcPending
	if v2437 != 0 {
		goto L8
	} else {
		goto L598
	}
L596:
	;
	v2431 = F_replace_nestloop_params_mutator(m, v2426, l0)
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L8
	} else {
		goto L597
	}
L597:
	;
	v2433 = v2429
	v2434 = v2431
	goto L595
L598:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2436)+80)) = v2434
	*(*int32)(unsafe.Add(mBase, uint32(v2436)+72)) = v564
	*(*int32)(unsafe.Add(mBase, uint32(v2436)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2436)+52)) = v572
	*(*int32)(unsafe.Add(mBase, uint32(v2436)+48)) = v2433
	*(*int32)(unsafe.Add(mBase, uint32(v2436)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v2436))) = int32(344)
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2436)+4)) = v2447
	v2449 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v2436)+8)) = v2449
	v2451 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v2436)+16)) = v2451
	v2453 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v2436)+24)) = v2453
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v2455)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2436)+32)) = v2456
	v2458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2436)+36)) = uint8(v2458)
	v2460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2436)+37)) = uint8(v2460)
	v2465 = v2436
	goto L126
L599:
	;
	v2478 = F_create_gating_plan(m, l0, l1, v2465, v66)
	mBase = m.M
	v2479 = m.ExcPending
	if v2479 != 0 {
		goto L8
	} else {
		goto L602
	}
L600:
	;
	v2480 = v2465
	goto L601
L601:
	;
	m.G0 = v19 + int32(176)
	return v2480
L602:
	;
	v2480 = v2478
	goto L601
}
func F_create_unique_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
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
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 float64
	_ = v403
	var v405 int32
	_ = v405
	var v407 float64
	_ = v407
	var v409 float64
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v507 float64
	_ = v507
	var v509 int32
	_ = v509
	var v511 float64
	_ = v511
	var v513 float64
	_ = v513
	var v515 int32
	_ = v515
	var v536 int32
	_ = v536
	var v537 float64
	_ = v537
	var v540 float64
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 float64
	_ = v551
	var v552 float64
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v567 float64
	_ = v567
	var v570 int32
	_ = v570
	var v571 float64
	_ = v571
	var v577 float64
	_ = v577
	var v584 float64
	_ = v584
	var v585 float64
	_ = v585
	var v589 float64
	_ = v589
	var v592 int32
	_ = v592
	var v595 float64
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v604 float64
	_ = v604
	var v606 int32
	_ = v606
	var v610 float64
	_ = v610
	var v611 float64
	_ = v611
	var v614 float64
	_ = v614
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v631 float64
	_ = v631
	var v633 int32
	_ = v633
	var v634 float64
	_ = v634
	var v635 float64
	_ = v635
	var v636 float64
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 float64
	_ = v661
	var v662 float64
	_ = v662
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v687 float64
	_ = v687
	var v689 float64
	_ = v689
	var v716 int32
	_ = v716
	var v737 int32
	_ = v737
	v5 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(176)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v24 != 0 {
		v737 = v24
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v22 + int32(176)
	return v737
L2:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+45)))
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(int32(5)) < base.Ui32(v30) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+46)))
	if v26 == int32(1) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v737 = int32(0)
	goto L1
L6:
	;
	v42 = F_GetMemoryChunkContext(m, l1)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	if int32(1)<<(uint(v30)%32)&int32(44) == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+224))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	if v40 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v737 = int32(0)
	goto L1
L10:
	;
	return int32(0)
L11:
	;
	v46 = int32(4554128)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v42
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(int32(5)) < base.Ui32(v50) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v47
	v737 = v716
	goto L1
L13:
	;
	v366 = F_palloc0(m, int32(88))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L10
	} else {
		goto L104
	}
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v74 = int32(0)
	v77 = v5
	v78 = v5
	v81 = v5
	v83 = v5
	goto L20
L15:
	;
	if int32(1)<<(uint(v50)%32)&int32(44) == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+224))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+84))
	v62 = F_adjust_appendrel_attrs_multilevel(m, l0, v61, l1, v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+80))
	v65 = F_copyObjectImpl(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v353 = v62
	v359 = v65
	goto L13
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L10
	} else {
		goto L101
	}
L20:
	;
	v89 = int32(0)
	if v68 == v89 {
		v100 = v89
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L10
	} else {
		goto L98
	}
L22:
	;
	if v67 == int32(0) {
		v716 = v89
		goto L12
	} else {
		goto L25
	}
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v94 <= v74 {
		v100 = int32(0)
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v100 = v96 + v74<<(uint(int32(2))%32)
	goto L22
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v103 <= v74 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v114 = F_get_ordering_op_for_equality_op(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L10
	} else {
		goto L34
	}
L27:
	;
	if v77 != 0 {
		v353 = v77
		v359 = v83
		goto L13
	} else {
		goto L31
	}
L28:
	;
	if v100 == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v110 = v107 + v74<<(uint(int32(2))%32)
	if v110 != 0 {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v716 = v89
	goto L12
L32:
	;
	goto L21
L33:
	;
	v310 = F_copyObjectImpl(m, v112)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L10
	} else {
		goto L95
	}
L34:
	;
	if v114 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v117 = F_get_equality_op_for_ordering_op(m, v114, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+45)))
	if v302 == int32(1) {
		goto L19
	} else {
		goto L94
	}
L38:
	;
	if v117 == int32(0) {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	if v78 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+4)))
	v125 = v121 + int32(1)
	goto L42
L41:
	;
	v125 = int32(1)
	goto L42
L42:
	;
	v127 = int32(0)
	v129 = F_makeTargetEntry(m, v112, base.I32_extend16_s(v125), v127, v127)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	v131 = F_lappend(m, v78, v129)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	v134 = F_palloc0(m, int32(20))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = int32(106)
	v138 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	if v147 == v138 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v278 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v134)+18)) = uint8(v278)
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+16)) = uint16(v278)
	*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v134)+8)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v134)+4)) = v269
	v285 = F_lappend(m, v81, v134)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L10
	} else {
		goto L83
	}
L47:
	;
	if v131 == int32(0) {
		v265 = int32(1)
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v269 = v147
	goto L49
L49:
	;
	goto L46
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+16)) = v265
	v269 = v265
	goto L49
L51:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if v154 <= int32(0) {
		v265 = int32(1)
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v157 = int32(0)
	if v157 < v154 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v160 = v154
	goto L55
L54:
	;
	v160 = v157
	goto L55
L55:
	;
	v162 = v160 & int32(3)
	v163 = int32(0)
	if int32(4) <= v154 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v172 = v163
	v173 = v138
	v174 = int32(0)
	goto L59
L57:
	;
	v207 = v163
	v208 = v138
	goto L58
L58:
	;
	if v162 != 0 {
		goto L74
	} else {
		goto L75
	}
L59:
	;
	v183 = v168 + v173<<(uint(int32(2))%32)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+16))
	if base.Ui32(v172) < base.Ui32(v191) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v207 = v199
	v208 = v201
	goto L58
L61:
	;
	v193 = v191
	goto L63
L62:
	;
	v193 = v172
	goto L63
L63:
	;
	if base.Ui32(v193) < base.Ui32(v189) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v195 = v189
	goto L66
L65:
	;
	v195 = v193
	goto L66
L66:
	;
	if base.Ui32(v195) < base.Ui32(v187) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v197 = v187
	goto L69
L68:
	;
	v197 = v195
	goto L69
L69:
	;
	if base.Ui32(v197) < base.Ui32(v185) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v199 = v185
	goto L72
L71:
	;
	v199 = v197
	goto L72
L72:
	;
	v200 = int32(4)
	v201 = v173 + v200
	v203 = v174 + v200
	if v203 != v160&int32(2147483644) {
		v172 = v199
		v173 = v201
		v174 = v203
		goto L59
	} else {
		goto L73
	}
L73:
	;
	goto L60
L74:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v218 = int32(0)
	v220 = v207
	v221 = v208
	goto L77
L75:
	;
	v243 = v207
	goto L76
L76:
	;
	v265 = v243 + int32(1)
	goto L50
L77:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v216+v221<<(uint(int32(2))%32))))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
	if base.Ui32(v220) < base.Ui32(v233) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v243 = v235
	goto L76
L79:
	;
	v235 = v233
	goto L81
L80:
	;
	v235 = v220
	goto L81
L81:
	;
	v236 = int32(1)
	v239 = v218 + v236
	if v239 != v162 {
		v218 = v239
		v220 = v235
		v221 = v221 + v236
		goto L77
	} else {
		goto L82
	}
L82:
	;
	goto L78
L83:
	;
	v287 = F_make_pathkeys_for_sortclauses(m, l0, v285, v131)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	if v287 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	v291 = v289
	goto L87
L86:
	;
	v291 = int32(0)
	goto L87
L87:
	;
	if v285 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	v294 = v292
	goto L90
L89:
	;
	v294 = int32(0)
	goto L90
L90:
	;
	if v291 == v294 {
		v306 = v131
		v308 = v285
		goto L33
	} else {
		goto L91
	}
L91:
	;
	v296 = F_list_delete_last(m, v285)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	v298 = F_list_delete_last(m, v131)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	v74 = v74 + int32(1)
	v78 = v298
	v81 = v296
	goto L20
L94:
	;
	v306 = v78
	v308 = v81
	goto L33
L95:
	;
	v312 = F_lappend(m, v77, v310)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L10
	} else {
		goto L96
	}
L96:
	;
	v314 = F_lappend_oid(m, v83, v113)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L10
	} else {
		goto L97
	}
L97:
	;
	v74 = v74 + int32(1)
	v77 = v312
	v78 = v306
	v81 = v308
	v83 = v314
	goto L20
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v114
	F_errmsg_internal(m, int32(46859), v22+int32(16))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(522980), int32(1841), int32(336656))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L10
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v113
	F_errmsg_internal(m, int32(46453), v22)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(522980), int32(1878), int32(336656))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v366))) = int64(1576252997927)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+12)) = v371
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v374 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v366)+20)) = uint8(v374)
	*(*int32)(unsafe.Add(mBase, uint32(v366)+16)) = v373
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v377 == int32(1) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	v382 = v380
	goto L107
L106:
	;
	v382 = int32(0)
	goto L107
L107:
	;
	v384 = v382 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v366)+21)) = uint8(v384)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+84)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v366)+80)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v366)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v366)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v366)+24)) = v386
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v393 != 0 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v366
	v716 = v366
	goto L12
L109:
	;
	v536 = int32(0)
	v537 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v540 = F_estimate_num_groups(m, l0, v353, v537, v536, v536)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L10
	} else {
		goto L143
	}
L110:
	;
	v414 = v393
	goto L112
L111:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+45)))
	if v394 != int32(1) {
		goto L109
	} else {
		goto L113
	}
L112:
	;
	if v414 != int32(1) {
		goto L109
	} else {
		goto L118
	}
L113:
	;
	v397 = int32(0)
	v399 = F_relation_has_unique_index_ext(m, l0, l1, v397, v353, v359, v397)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L10
	} else {
		goto L114
	}
L114:
	;
	if v399 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+76)) = int32(0)
	v403 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v366)+32)) = v403
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+40)) = v405
	v407 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v366)+48)) = v407
	v409 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v366)+56)) = v409
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+64)) = v411
	goto L108
L116:
	;
	goto L117
L117:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v414 = v413
	goto L112
L118:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v417 != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)+36))
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+38)))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v433)+120))
	v440 = v434 ^ int32(1) | base.B2i32(v437 != int32(0))
	if v434 != 0 {
		v449 = v440
		goto L123
	} else {
		goto L124
	}
L120:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v431 = v417 + v418<<(uint(int32(2))%32)
	goto L119
L121:
	;
	goto L122
L122:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)+52))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+12))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v431 = v424 + v425<<(uint(int32(2))%32) - int32(4)
	goto L119
L123:
	;
	if v449 == int32(0) {
		goto L109
	} else {
		goto L130
	}
L124:
	;
	if v437 != 0 {
		v449 = v440
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v441 = int32(1)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v433)+100))
	if v442 != 0 {
		v449 = v441
		goto L123
	} else {
		goto L126
	}
L126:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v433)+108))
	if v443 != 0 {
		v449 = v441
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+36)))
	if v444 != 0 {
		v449 = v441
		goto L123
	} else {
		goto L128
	}
L128:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v433)+112))
	if v445 != 0 {
		v449 = v441
		goto L123
	} else {
		goto L129
	}
L129:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v433)+144))
	v449 = base.B2i32(v446 != int32(0))
	goto L123
L130:
	;
	if v353 == int32(0) {
		goto L109
	} else {
		goto L131
	}
L131:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	if v454 <= int32(0) {
		goto L109
	} else {
		goto L132
	}
L132:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v458 = int32(0)
	v468 = v458
	v469 = v458
	goto L133
L133:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v479+v468<<(uint(int32(2))%32))))
	if v483 == int32(0) {
		goto L109
	} else {
		goto L135
	}
L134:
	;
	if v492 == int32(0) {
		goto L109
	} else {
		goto L140
	}
L135:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	if v486 != int32(6) {
		goto L109
	} else {
		goto L136
	}
L136:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v483)+4))
	if v489 != v457 {
		goto L109
	} else {
		goto L137
	}
L137:
	;
	v491 = int32(*(*int16)(unsafe.Add(mBase, uint32(v483)+8)))
	v492 = F_lappend_int(m, v469, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L10
	} else {
		goto L138
	}
L138:
	;
	v495 = v468 + int32(1)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	if v495 < v496 {
		v468 = v495
		v469 = v492
		goto L133
	} else {
		goto L139
	}
L139:
	;
	goto L134
L140:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v432)+36))
	v501 = F_query_is_distinct_for(m, v500, v492, v359)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L10
	} else {
		goto L141
	}
L141:
	;
	if v501 == int32(0) {
		goto L109
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+76)) = int32(0)
	v507 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v366)+32)) = v507
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+40)) = v509
	v511 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v366)+48)) = v511
	v513 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v366)+56)) = v513
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+64)) = v515
	goto L108
L143:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v366)+32)) = v540
	if v353 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	v544 = v543
	goto L146
L145:
	;
	v544 = v536
	goto L146
L146:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+45)))
	if v545 == int32(1) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v549 = v22 + int32(104)
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v551 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v552 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)+32))
	v557 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	v560 = m.G0
	v561 = int32(16)
	v562 = v560 - v561
	m.G0 = v562
	F_cost_tuplesort(m, v562+int32(8), v562, v552, v554, float64(0), v557, float64(-1))
	mBase = m.M
	v567 = *(*float64)(unsafe.Add(mBase, uint32(v562)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v549)+32)) = v552
	v570 = int32(*(*uint8)(unsafe.Add(mBase, _consts[636])))
	v571 = base.F64_add(v551, v567)
	*(*float64)(unsafe.Add(mBase, uint32(v549)+48)) = v571
	*(*int32)(unsafe.Add(mBase, uint32(v549)+40)) = v550 + (v570 ^ int32(1))
	v577 = *(*float64)(unsafe.Add(mBase, uint32(v562)))
	*(*float64)(unsafe.Add(mBase, uint32(v549)+56)) = base.F64_add(v571, v577)
	m.G0 = v562 + v561
	goto L150
L148:
	;
	goto L149
L149:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+46)))
	if v592 != int32(1) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v584 = *(*float64)(unsafe.Add(mBase, _consts[613]))
	v585 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v589 = *(*float64)(unsafe.Add(mBase, uint32(v22)+160))
	*(*float64)(unsafe.Add(mBase, uint32(v22)+160)) = base.F64_add(base.F64_mul(base.F64_mul(v584, v585), base.F64_convert_i32_s(v544)), v589)
	goto L149
L151:
	;
	v643 = v22 + int32(88)
	v645 = v22 + int32(80)
	v647 = v22 + int32(72)
	v648 = int32(1)
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+46)))
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+45)))
	if v650 == v648 {
		goto L165
	} else {
		goto L166
	}
L152:
	;
	v595 = *(*float64)(unsafe.Add(mBase, uint32(v366)+32))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v596)+32))
	v604 = *(*float64)(unsafe.Add(mBase, _consts[498]))
	v606 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	v610 = base.F64_mul(base.F64_mul(v604, base.F64_convert_i32_s(v606)), float64(1024))
	v611 = float64(4.294967295e+09)
	if base.F64_lt(v610, v611) != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if base.F64_gt(base.F64_mul(v595, base.F64_convert_i32_s(v597-int32(-64))), base.F64_convert_i32_u(v622)) != 0 {
		goto L160
	} else {
		goto L161
	}
L154:
	;
	v614 = v610
	goto L156
L155:
	;
	v614 = v611
	goto L156
L156:
	;
	if base.F64_lt(v614, float64(4.294967296e+09))&base.F64_ge(v614, float64(0)) != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v620 = base.I32_trunc_f64_u(v614)
	v622 = v620
	goto L153
L158:
	;
	goto L159
L159:
	;
	v622 = int32(0)
	goto L153
L160:
	;
	v625 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+46)) = uint8(v625)
	goto L151
L161:
	;
	goto L162
L162:
	;
	v630 = int32(0)
	v631 = *(*float64)(unsafe.Add(mBase, uint32(v366)+32))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v634 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v635 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v636 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v637)+32))
	F_cost_agg(m, v22+int32(32), l0, int32(2), v630, v544, v631, v630, v633, v634, v635, v636, base.F64_convert_i32_s(v638))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L10
	} else {
		goto L163
	}
L163:
	;
	goto L151
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+76)) = v682
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+40)) = v685
	v687 = *(*float64)(unsafe.Add(mBase, uint32(v683)))
	*(*float64)(unsafe.Add(mBase, uint32(v366)+48)) = v687
	v689 = *(*float64)(unsafe.Add(mBase, uint32(v679)))
	*(*float64)(unsafe.Add(mBase, uint32(v366)+56)) = v689
	goto L108
L165:
	;
	if v649&int32(1) == int32(0) {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	goto L167
L167:
	;
	v673 = int32(0)
	if v649&int32(1) == v673 {
		v716 = v673
		goto L12
	} else {
		goto L173
	}
L168:
	;
	v678 = v22 + int32(144)
	v679 = v22 + int32(160)
	v682 = int32(2)
	v683 = v22 + int32(152)
	goto L164
L169:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v22)+144))
	if v657 < v658 {
		v678 = v647
		v679 = v643
		v682 = v648
		v683 = v645
		goto L164
	} else {
		goto L170
	}
L170:
	;
	if v658 != v657 {
		goto L168
	} else {
		goto L171
	}
L171:
	;
	v661 = *(*float64)(unsafe.Add(mBase, uint32(v22)+88))
	v662 = *(*float64)(unsafe.Add(mBase, uint32(v22)+160))
	if base.F64_lt(v661, v662) != 0 {
		v678 = v647
		v679 = v643
		v682 = v648
		v683 = v645
		goto L164
	} else {
		goto L172
	}
L172:
	;
	goto L168
L173:
	;
	v678 = v647
	v679 = v643
	v682 = v648
	v683 = v645
	goto L164
}
func F_crosstab_hash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int64
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int64
	_ = v161
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int64
	_ = v200
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
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
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int64
	_ = v329
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v401 int64
	_ = v401
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v528 int32
	_ = v528
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int64
	_ = v601
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int64
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	v31 = m.G0
	v33 = v31 - int32(160)
	m.G0 = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = F_pg_detoast_datum_packed(m, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v40 = F_text_to_cstring(m, v36)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v43 = F_pg_detoast_datum_packed(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v45 = F_text_to_cstring(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v47 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L1
	} else {
		goto L154
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L149
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L144
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L140
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L1
	} else {
		goto L137
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L133
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L1
	} else {
		goto L129
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L124
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L119
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L115
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L111
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v50 != int32(383) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+12)))
	if v53&int32(2) == int32(0) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	if v58 == int32(0) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v61 = int32(4554128)
	v62 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v65
	v67 = F_CreateTupleDescCopy(m, v58)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v69 <= int32(1) {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v65
	*(*int64)(unsafe.Add(mBase, uint32(v33)+64)) = int64(292057776192)
	v80 = F_hash_create(m, int32(338591), int32(64), v33+int32(48), int32(1048))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v87 = F_SPI_execute(m, v45, int32(1), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v232 = F_SPI_finish(m)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L39
	}
L26:
	;
	if v87 != int32(5) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v92 = *(*int64)(unsafe.Add(mBase, _consts[510]))
	if v92 == int64(0) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v98 != int32(1) {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v139 = int64(0)
	goto L30
L30:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141+base.I32_wrap_i64(v139)<<(uint(int32(2))%32))))
	v148 = F_SPI_getvalue(m, v146, v97, int32(1))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L25
L32:
	;
	if v148 == int32(0) {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v152 = int32(4554128)
	v153 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v65
	v157 = F_palloc(m, int32(16))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v157)+8)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v148
	v161 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(152)))) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(144)))) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(136)))) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(128)))) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(120)))) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v33)+112)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v33)+104)) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v33)+96)) = v161
	v184 = F_pg_snprintf(m, v33+int32(96), int32(63), int32(216470), v33+int32(32))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v191 = F_hash_search(m, v80, v33+int32(96), int32(1), v33+int32(47))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+47)))
	if v193 == int32(1) {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+64)) = v157
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v153
	v200 = v139 + int64(1)
	if v200 != v92 {
		v139 = v200
		goto L30
	} else {
		goto L38
	}
L38:
	;
	goto L31
L39:
	;
	if v232 != int32(2) {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = int32(2)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v240)+412))
	if v242 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v308 = F_TupleDescGetAttInMetadata(m, v67)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L45
	}
L42:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v240)+376))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240)+364))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v240)+352))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v240)+340))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v240)+328))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v240)+316))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v240)+304))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v240)+292))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v240)+280))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v240)+268))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v240)+256))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v240)+244))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v240)+232))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v240)+220))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v240)+208))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v240)+196))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v240)+184))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v240)+172))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v240)+160))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v240)+148))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v240)+136))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v240)+124))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v240)+112))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v240)+100))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v240)+88))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v240)+76))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v240-int32(-64))))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v240)+52))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v240)+40))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v240)+28))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v240)+16))
	v307 = v243 + (v244 + (v245 + (v246 + (v247 + (v248 + (v249 + (v250 + (v251 + (v252 + (v253 + (v254 + (v255 + (v256 + (v257 + (v258 + (v259 + (v260 + (v261 + (v262 + (v263 + (v264 + (v265 + (v266 + (v267 + (v268 + (v271 + (v272 + (v273 + (v274 + (v275 + v241))))))))))))))))))))))))))))))
	goto L44
L43:
	;
	v307 = v241
	goto L44
L44:
	;
	goto L41
L45:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	v317 = F_tuplestore_begin_heap(m, int32(base.Ui32(v238&int32(4))>>(uint(int32(2))%32)), int32(0), v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v324 = F_SPI_execute(m, v40, int32(1), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	v688 = F_SPI_finish(m)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L109
	}
L49:
	;
	if v324 != int32(5) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v329 = *(*int64)(unsafe.Add(mBase, _consts[510]))
	if v329 == int64(0) {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	if v307 == int32(0) {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	if v337 <= int32(2) {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v341 = v337 - int32(2)
	v342 = v341 + v307
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v342 != v343 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	v347 = F_palloc0(m, v342<<(uint(int32(2))%32))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v354 = int32(1)
	v373 = v354
	v377 = int32(0)
	v401 = int64(0)
	goto L56
L56:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v403+base.I32_wrap_i64(v401)<<(uint(int32(2))%32))))
	v410 = F_SPI_getvalue(m, v408, v336, int32(1))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v654 = F_BuildTupleFromCStrings(m, v308, v347)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L107
	}
L58:
	;
	if v373&int32(1) != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v597 = F_SPI_getvalue(m, v408, v336, v337-v354)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L91
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v410
	if v337 == int32(3) {
		goto L59
	} else {
		goto L85
	}
L61:
	;
	if v377|v410 == int32(0) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	if v377 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v448 = F_BuildTupleFromCStrings(m, v308, v347)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L75
	}
L64:
	;
	if v410 == int32(0) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377))))
	if v424 == int32(0) {
		v443 = v423
		v444 = v424
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v444-v443 == int32(0) {
		goto L59
	} else {
		goto L74
	}
L67:
	;
	goto L66
L68:
	;
	if v423 != v424 {
		v443 = v423
		v444 = v424
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v428 = v377
	v429 = v410
	goto L70
L70:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+1)))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428)+1)))
	if v433 == int32(0) {
		v443 = v432
		v444 = v433
		goto L67
	} else {
		goto L72
	}
L71:
	;
	v443 = v432
	v444 = v433
	goto L67
L72:
	;
	v436 = int32(1)
	if v432 == v433 {
		v428 = v428 + v436
		v429 = v429 + v436
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	goto L63
L75:
	;
	F_tuplestore_puttuple(m, v317, v448)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v452 = int32(0)
	if v342 <= v452 {
		goto L60
	} else {
		goto L77
	}
L77:
	;
	v455 = v452
	goto L78
L78:
	;
	v487 = v347 + v455<<(uint(int32(2))%32)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	if v488 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L60
L80:
	;
	F_pfree(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v494 = v455 + int32(1)
	if v494 != v342 {
		v455 = v494
		goto L78
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = int32(0)
	goto L82
L84:
	;
	goto L79
L85:
	;
	v528 = int32(1)
	goto L86
L86:
	;
	v562 = v528 + int32(1)
	v563 = F_SPI_getvalue(m, v408, v336, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L88
	}
L87:
	;
	goto L59
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347+v528<<(uint(int32(2))%32)))) = v563
	if v562 != v341 {
		v528 = v562
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	if v377 != 0 {
		goto L98
	} else {
		goto L99
	}
L91:
	;
	if v597 == int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v601 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(152)))) = v601
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(144)))) = v601
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(136)))) = v601
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(128)))) = v601
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(120)))) = v601
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(112)))) = v601
	*(*int64)(unsafe.Add(mBase, uint32(v33)+104)) = v601
	*(*int64)(unsafe.Add(mBase, uint32(v33)+96)) = v601
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v597
	v622 = F_pg_snprintf(m, v33+int32(96), int32(63), int32(216470), v33)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v626 = int32(0)
	v628 = F_hash_search(m, v80, v33+int32(96), v626, v626)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v628 == int32(0) {
		goto L90
	} else {
		goto L95
	}
L95:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v628)+64))
	if v632 == int32(0) {
		goto L90
	} else {
		goto L96
	}
L96:
	;
	v635 = F_SPI_getvalue(m, v408, v336, v337)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v632)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v347+v337<<(uint(int32(2))%32)-int32(8)+v637<<(uint(int32(2))%32)))) = v635
	goto L90
L98:
	;
	F_pfree(m, v377)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v646 = int32(0)
	if v410 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L100
L102:
	;
	v648 = F_pstrdup(m, v410)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	v650 = v646
	goto L104
L104:
	;
	v652 = v401 + int64(1)
	if v652 != v329 {
		v373 = v646
		v377 = v650
		v401 = v652
		goto L56
	} else {
		goto L106
	}
L105:
	;
	v650 = v648
	goto L104
L106:
	;
	goto L57
L107:
	;
	F_tuplestore_puttuple(m, v317, v654)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	goto L48
L109:
	;
	if v688 != int32(2) {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v317
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v62
	m.G0 = v33 + int32(160)
	return int32(0)
L111:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(113795), int32(0))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(523797), int32(650), int32(338347))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(66095), int32(0))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(523797), int32(655), int32(338347))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errmsg(m, int32(386063), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errdetail(m, int32(619219), int32(0))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(523797), int32(674), int32(338347))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(16491), int32(0))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errdetail(m, int32(646130), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(523797), int32(749), int32(338106))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(317120), int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(523797), int32(765), int32(338106))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errmsg(m, int32(398660), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(523797), int32(774), int32(338106))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	F_errmsg_internal(m, int32(475086), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(523797), int32(782), int32(338106))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(31798), int32(0))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(523797), int32(832), int32(381471))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(17045), int32(0))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errdetail(m, int32(656459), int32(0))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(523797), int32(852), int32(381471))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(386063), int32(0))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v866
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v342
	F_errdetail(m, int32(680420), v33+int32(16))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(523797), int32(862), int32(381471))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_errmsg_internal(m, int32(475128), int32(0))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(523797), int32(932), int32(381471))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cursor_to_xml(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v82 int64
	_ = v82
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = F_text_to_cstring(m, v13)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v22 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = F_text_to_cstring(m, v22)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_initStringInfo(m, v10-int32(-64))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v19 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(413344)
	F_appendStringInfo(m, v10-int32(-64), int32(187179), v10+int32(48))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L10:
	;
	F_appendStringInfoString(m, v10-int32(-64), int32(761945))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v46 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v24
	F_appendStringInfo(m, v10-int32(-64), int32(722014), v10+int32(32))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_appendStringInfoString(m, v10-int32(-64), int32(786838))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	F_appendStringInfoChar(m, v10-int32(-64), int32(10))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L9
L18:
	;
	v68 = F_GetPortalByName(m, v17)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v68 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_SPI_cursor_fetch(m, v68, v20)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L37
	}
L23:
	;
	v73 = *(*int64)(unsafe.Add(mBase, _consts[510]))
	if v73 != int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v82 = int64(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	v101 = F_SPI_finish(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L31
	}
L27:
	;
	F_SPI_sql_row_to_xmlelement(m, v10-int32(-64), base.B2i32(v19 != int32(0)), v24)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v90 = v82 + int64(1)
	v92 = *(*int64)(unsafe.Add(mBase, _consts[510]))
	if base.Ui64(v90) < base.Ui64(v92) {
		v82 = v90
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if v19 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(413344)
	F_appendStringInfo(m, v10-int32(-64), int32(785394), v10+int32(16))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
	v116 = F_cstring_to_text_with_len(m, v114, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	m.G0 = v10 + int32(80)
	return v116
L37:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v17
	F_errmsg(m, int32(76741), v10)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(520832), int32(2937), int32(315694))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
