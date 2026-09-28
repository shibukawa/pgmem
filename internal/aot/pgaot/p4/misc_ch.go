package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_CheckElement_3(m *base.Module, l0 float32) {
	var v8 int32
	_ = v8
	Fn14209(m, l0, int32(122), int32(_a_F_CheckElement_3_0), int32(_a_F_CheckElement_3_1), int32(117), int32(_a_F_CheckElement_3_2))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
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
	var v67 int32
	_ = v67
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
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+180))
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
						v67 = F_errdetail(m, int32(_a_F_CheckRequiredParameterValues_1), int32(0))
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
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+188))
				F_RecoveryRequiresIntParameter(m, int32(_a_F_CheckRequiredParameterValues_6), v17, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[4]))
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[1]))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+192))
					F_RecoveryRequiresIntParameter(m, int32(_a_F_CheckRequiredParameterValues_7), v23, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[5]))
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[1]))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+196))
						F_RecoveryRequiresIntParameter(m, int32(_a_F_CheckRequiredParameterValues_8), v31, v34)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[6]))
							v41 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[1]))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+200))
							F_RecoveryRequiresIntParameter(m, int32(_a_F_CheckRequiredParameterValues_9), v39, v42)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[7]))
								v49 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRequiredParameterValues[1]))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+204))
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
					F_errfinish(m, int32(_a_F_CheckRestrictedOperation_1), int32(468), int32(_a_F_CheckRestrictedOperation_2))
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v271 int32
	_ = v271
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v316 int32
	_ = v316
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v361 int32
	_ = v361
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v414 int64
	_ = v414
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v619 int32
	_ = v619
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
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
	var v679 int32
	_ = v679
	var v682 int64
	_ = v682
	var v684 int64
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v749 int64
	_ = v749
	var v753 int32
	_ = v753
	var v755 int64
	_ = v755
	var v761 int32
	_ = v761
	var v763 int64
	_ = v763
	var v767 int32
	_ = v767
	var v769 int64
	_ = v769
	var v779 int32
	_ = v779
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v799 int32
	_ = v799
	var v807 int32
	_ = v807
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v828 int64
	_ = v828
	var v829 int32
	_ = v829
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v846 int64
	_ = v846
	var v847 int32
	_ = v847
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v899 int32
	_ = v899
	var v906 int32
	_ = v906
	var v908 int64
	_ = v908
	var v914 int32
	_ = v914
	var v916 int64
	_ = v916
	var v922 int64
	_ = v922
	var v928 int32
	_ = v928
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1015 int64
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1019 int64
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1047 int64
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1061 int32
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1118 int64
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1143 int32
	_ = v1143
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1255 int32
	_ = v1255
	var v1270 int32
	_ = v1270
	var v1271 int64
	_ = v1271
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(208)
	m.G0 = v17
	v21 = int32(-1)
	v22 = v3
	v23 = v3
	v26 = v3
	v27 = v3
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
	if v21 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v1270 = int32(m.ExcTag)
	v1271 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1270 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	v37 = int32(1)
	v38 = v26 & v37
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v38)
	v41 = v27 & v37
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v41)
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v447 = v22
	v448 = v23
	goto L8
L8:
	;
	if v448 != 0 {
		goto L96
	} else {
		goto L97
	}
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[0]))
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v48
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v41)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v38)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	v57 = m.G0
	v59 = v57 - int32(32)
	m.G0 = v59
	v62 = int32(967)
	switch v62 {
	case 0, 2:
		goto L11
	default:
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v41)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v38)
	v102 = m.G0
	v104 = v102 - int32(32)
	m.G0 = v104
	v107 = int32(989)
	switch v107 {
	case 0, 2:
		goto L21
	default:
		goto L22
	}
L11:
	;
	F_sigemptyset(m, v59+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = int32(268435456)
	switch v62 {
	case 0:
		goto L16
	default:
		goto L14
	case 2:
		goto L15
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[2])) = int32(965)
	goto L11
L13:
	;
	goto L18
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = int32(_a_F_CheckpointerMain_0)
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = int32(0)
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = int32(-2)
	goto L13
L18:
	;
	goto L19
L19:
	;
	v91 = F___sigaction(m, int32(1), v59+int32(12), int32(0))
	mBase = m.M
	m.G0 = v59 + int32(32)
	goto L10
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v41)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v38)
	v145 = int32(0)
	v147 = m.G0
	v149 = v147 - int32(32)
	m.G0 = v149
	switch v145 {
	case 0, 2:
		goto L31
	default:
		goto L32
	}
L21:
	;
	F_sigemptyset(m, v104+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v104)+24)) = int32(268435456)
	switch v107 {
	case 0:
		goto L26
	default:
		goto L24
	case 2:
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[3])) = int32(987)
	goto L21
L23:
	;
	goto L28
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = int32(_a_F_CheckpointerMain_0)
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = int32(0)
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = int32(-2)
	goto L23
L28:
	;
	goto L29
L29:
	;
	v136 = F___sigaction(m, int32(2), v104+int32(12), int32(0))
	mBase = m.M
	m.G0 = v104 + int32(32)
	goto L20
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v41)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v38)
	v190 = int32(0)
	v192 = m.G0
	v194 = v192 - int32(32)
	m.G0 = v194
	switch v190 {
	case 0, 2:
		goto L41
	default:
		goto L42
	}
L31:
	;
	F_sigemptyset(m, v149+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v149)+24)) = int32(268435456)
	switch v145 {
	case 0:
		goto L36
	default:
		goto L34
	case 2:
		goto L35
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[4])) = int32(-2)
	goto L31
L33:
	;
	goto L38
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = int32(_a_F_CheckpointerMain_0)
	goto L33
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = int32(0)
	goto L33
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = int32(-2)
	goto L33
L38:
	;
	goto L39
L39:
	;
	v181 = F___sigaction(m, int32(15), v149+int32(12), int32(0))
	mBase = m.M
	m.G0 = v149 + int32(32)
	goto L30
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v41)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v38)
	v235 = int32(0)
	v237 = m.G0
	v239 = v237 - int32(32)
	m.G0 = v239
	switch v235 {
	case 0, 2:
		goto L51
	default:
		goto L52
	}
L41:
	;
	F_sigemptyset(m, v194+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v194)+24)) = int32(268435456)
	switch v190 {
	case 0:
		goto L46
	default:
		goto L44
	case 2:
		goto L45
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[5])) = int32(-2)
	goto L41
L43:
	;
	goto L48
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v194)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v194)+12)) = int32(_a_F_CheckpointerMain_0)
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v194)+12)) = int32(0)
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v194)+12)) = int32(-2)
	goto L43
L48:
	;
	goto L49
L49:
	;
	v226 = F___sigaction(m, int32(14), v194+int32(12), int32(0))
	mBase = m.M
	m.G0 = v194 + int32(32)
	goto L40
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v41)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v38)
	v282 = m.G0
	v284 = v282 - int32(32)
	m.G0 = v284
	v287 = int32(970)
	switch v287 {
	case 0, 2:
		goto L61
	default:
		goto L62
	}
L51:
	;
	F_sigemptyset(m, v239+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v239)+24)) = int32(268435456)
	switch v235 {
	case 0:
		goto L56
	default:
		goto L54
	case 2:
		goto L55
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[6])) = int32(-2)
	goto L51
L53:
	;
	goto L58
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v239)+12)) = int32(_a_F_CheckpointerMain_0)
	goto L53
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+12)) = int32(0)
	goto L53
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+12)) = int32(-2)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v271 = F___sigaction(m, int32(13), v239+int32(12), int32(0))
	mBase = m.M
	m.G0 = v239 + int32(32)
	goto L50
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v41)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v38)
	v327 = m.G0
	v329 = v327 - int32(32)
	m.G0 = v329
	v332 = int32(969)
	switch v332 {
	case 0, 2:
		goto L71
	default:
		goto L72
	}
L61:
	;
	F_sigemptyset(m, v284+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v284)+24)) = int32(268435456)
	switch v287 {
	case 0:
		goto L66
	default:
		goto L64
	case 2:
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[7])) = int32(968)
	goto L61
L63:
	;
	goto L68
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v284)+12)) = int32(_a_F_CheckpointerMain_0)
	goto L63
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+12)) = int32(0)
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+12)) = int32(-2)
	goto L63
L68:
	;
	goto L69
L69:
	;
	v316 = F___sigaction(m, int32(10), v284+int32(12), int32(0))
	mBase = m.M
	m.G0 = v284 + int32(32)
	goto L60
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v41)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v38)
	v372 = m.G0
	v374 = v372 - int32(32)
	m.G0 = v374
	v376 = int32(2)
	switch v376 {
	case 0, 2:
		goto L81
	default:
		goto L82
	}
L71:
	;
	F_sigemptyset(m, v329+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v329)+24)) = int32(268435456)
	switch v332 {
	case 0:
		goto L76
	default:
		goto L74
	case 2:
		goto L75
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[8])) = int32(967)
	goto L71
L73:
	;
	goto L78
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v329)+12)) = int32(_a_F_CheckpointerMain_0)
	goto L73
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+12)) = int32(0)
	goto L73
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+12)) = int32(-2)
	goto L73
L78:
	;
	goto L79
L79:
	;
	v361 = F___sigaction(m, int32(12), v329+int32(12), int32(0))
	mBase = m.M
	m.G0 = v329 + int32(32)
	goto L70
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v41)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v38)
	v414 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[9])) = v414
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10])) = v414
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v38)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v41)
	F_before_shmem_exit(m, int32(988), int64(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L5
	} else {
		goto L90
	}
L81:
	;
	F_sigemptyset(m, v374+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v374)+24)) = int32(268435456)
	switch v376 {
	case 0:
		goto L86
	default:
		goto L84
	case 2:
		goto L85
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[11])) = int32(0)
	goto L81
L83:
	;
	goto L87
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+12)) = int32(_a_F_CheckpointerMain_0)
	v399 = int32(268435461)
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374)+12)) = int32(0)
	v399 = int32(268435457)
	goto L83
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374)+12)) = int32(-2)
	v399 = int32(268435457)
	goto L83
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374)+24)) = v399
	goto L89
L89:
	;
	v406 = F___sigaction(m, int32(17), v374+int32(12), int32(0))
	mBase = m.M
	m.G0 = v374 + int32(32)
	goto L80
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v38)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v41)
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[12]))
	v434 = F_AllocSetContextCreateInternal(m, v429, int32(_a_F_CheckpointerMain_1), int32(0), int32(_a_F_CheckpointerMain_2), int32(_a_F_CheckpointerMain_3))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[13])) = v434
	goto L92
L92:
	;
	v440 = v17 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v440)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = v17 + int32(28)
	goto L95
L93:
	;
	v447 = v434
	v448 = int32(0)
	goto L8
L95:
	;
	goto L93
L96:
	;
	v450 = int32(_a_F_CheckpointerMain_4)
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[14]))
	v453 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[14])) = v452 + v453
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[15])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	v461 = v26 & v453
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	v464 = v27 & v453
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	F_EmitErrorReport(m)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L5
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[16])) = v17 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	v586 = int32(1)
	v587 = v26 & v586
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v587)
	v590 = v27 & v586
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v590)
	F_pgmem_sigprocmask(m, int32(_a_F_CheckpointerMain_5), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L5
	} else {
		goto L118
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	F_LWLockReleaseAll(m)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	v479 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v479))) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	F_UnlockBuffers(m)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	F_smgrdestroyall(m)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[18])))
	if v519 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v521 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[0]))
	v524 = base.AtomicRmwXchg32(m, v521, int32(4), int32(1))
	if v524 != 0 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[13])) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	F_FlushErrorState(m)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L5
	} else {
		goto L116
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	F_s_lock(m, v521+int32(4), int32(_a_F_CheckpointerMain_6))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L5
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[0]))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v534)+12)) = v535
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v534)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v534)+16)) = v537 + int32(1)
	v541 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v534)+4)), uint32(v541))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	F_ConditionVariableBroadcast(m, v534+int32(36))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L5
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	v552 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[18])) = uint8(v552)
	goto L110
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	F_MemoryContextReset(m, v447)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v567 = int32(_a_F_CheckpointerMain_4)
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[14])) = v569 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v461)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v464)
	F_pg_usleep(m, int32(_a_F_CheckpointerMain_7))
	mBase = m.M
	goto L98
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v590)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v587)
	F_SyncRepUpdateSyncStandbysDefined(m)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v590)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v587)
	F_UpdateFullPageWrites(m)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v590)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v587)
	v611 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	if v611 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v590)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v587)
	F_errmsg_internal(m, int32(_a_F_CheckpointerMain_8), int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L5
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v629 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[19]))
	v631 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v629)+68)) = v631
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v587)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v590)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[21]))
	v638 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v637))) = v638
	v643 = base.AtomicRmwOr32(m, v638, int32(_a_F_CheckpointerMain_9), v638)
	goto L127
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v590)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v587)
	F_errfinish(m, int32(_a_F_CheckpointerMain_10), int32(1515), int32(_a_F_CheckpointerMain_11))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L5
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v590)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v587)
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v590)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v587)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	v655 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	if v655 != 0 {
		v1103 = v26
		v1104 = v27
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v1112 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[23])) = uint8(v1112)
	v1115 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	if v1115 != 0 {
		goto L227
	} else {
		goto L228
	}
L131:
	;
	v662 = v26
	v663 = v27
	goto L132
L132:
	;
	v671 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[24]))
	if v671 != 0 {
		v1103 = v662
		v1104 = v663
		goto L130
	} else {
		goto L134
	}
L133:
	;
	v1103 = v943
	v1104 = v944
	goto L130
L134:
	;
	v673 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[0]))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v673)+20))
	v675 = int32(1)
	v676 = v663 & v675
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v676)
	v679 = v662 & v675
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v679)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	v682 = F_time(m)
	mBase = m.M
	v684 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[9]))
	v686 = base.I32_wrap_i64(v682 - v684)
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[25]))
	v689 = base.B2i32(v688 <= v686)
	if v689|v674 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v679)
	v696 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[26])))
	if v696 == int32(1) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	v943 = v662
	v944 = v663
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	v950 = int32(1)
	v951 = v944 & v950
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v951)
	v954 = v943 & v950
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v954)
	v957 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[27]))
	if v957 != v950 {
		goto L194
	} else {
		goto L195
	}
L138:
	;
	v708 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[0]))
	v711 = base.AtomicRmwXchg32(m, v708, int32(4), int32(1))
	if v711 != 0 {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v701 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[28]))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)+308))
	v704 = base.B2i32(v702 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[26])) = uint8(v704)
	v706 = v704
	goto L141
L140:
	;
	v706 = int32(0)
	goto L141
L141:
	;
	goto L138
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v679)
	F_s_lock(m, v708+int32(4), int32(_a_F_CheckpointerMain_6))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L5
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[0]))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v721)+20))
	v723 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v721)+20)) = v723
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v721)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v721)+8)) = v725 + int32(1)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v721)+4)), uint32(v723))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v679)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	F_ConditionVariableBroadcast(m, v721+int32(24))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L5
	} else {
		goto L146
	}
L145:
	;
	goto L144
L146:
	;
	v741 = int32(0)
	v743 = base.B2i32(v722&int32(2) == v741) & v706
	if base.B2i32(v674 == v741)&v689 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	if v743|base.B2i32(v722&int32(128) == int32(0)) != 0 {
		goto L158
	} else {
		goto L159
	}
L148:
	;
	if v743 != 0 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	if v674 == int32(0) {
		goto L147
	} else {
		goto L154
	}
L151:
	;
	v747 = int32(_a_F_CheckpointerMain_12)
	v749 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[29]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[29])) = v749 + int64(1)
	goto L147
L152:
	;
	goto L153
L153:
	;
	v753 = int32(_a_F_CheckpointerMain_13)
	v755 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[30]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[30])) = v755 + int64(1)
	goto L147
L154:
	;
	if v743 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v761 = int32(_a_F_CheckpointerMain_14)
	v763 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[31]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[31])) = v763 + int64(1)
	goto L147
L156:
	;
	goto L157
L157:
	;
	v767 = int32(_a_F_CheckpointerMain_15)
	v769 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[32]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[32])) = v769 + int64(1)
	goto L147
L158:
	;
	if v688 <= v686 {
		goto L166
	} else {
		goto L167
	}
L159:
	;
	v779 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[33]))
	if v779 <= v686 {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v679)
	v786 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L5
	} else {
		goto L161
	}
L161:
	;
	if v786 == int32(0) {
		goto L158
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v679)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v686
	F_errmsg_plural(m, int32(_a_F_CheckpointerMain_16), int32(_a_F_CheckpointerMain_17), v686, v17+int32(16))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L5
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v679)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(_a_F_CheckpointerMain_18)
	F_errhint(m, int32(_a_F_CheckpointerMain_19), v17)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L5
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v679)
	F_errfinish(m, int32(_a_F_CheckpointerMain_10), int32(484), int32(_a_F_CheckpointerMain_20))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L5
	} else {
		goto L165
	}
L165:
	;
	goto L158
L166:
	;
	v818 = int32(256)
	goto L168
L167:
	;
	v818 = int32(0)
	goto L168
L168:
	;
	v819 = v722 | v818
	v821 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[18])) = uint8(v821)
	if v743 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	v865 = int32(1)
	v866 = v861 & v865
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v866)
	v869 = v860 & v865
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v869)
	F_smgrdestroyall(m)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L5
	} else {
		goto L177
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v679)
	v828 = F_GetInsertRecPtr(m)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L5
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v679)
	v846 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L5
	} else {
		goto L175
	}
L173:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[34])) = v682
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[35])) = v828
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[36])) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v679)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v676)
	v840 = F_CreateCheckPoint(m, v819)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L5
	} else {
		goto L174
	}
L174:
	;
	v860 = v662
	v861 = v840
	v863 = v840
	goto L169
L175:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[34])) = v682
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[35])) = v846
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[36])) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v679)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v676)
	v858 = F_CreateRestartPoint(m, v819)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L5
	} else {
		goto L176
	}
L176:
	;
	v860 = v858
	v861 = v663
	v863 = v858
	goto L169
L177:
	;
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[0]))
	v877 = base.AtomicRmwXchg32(m, v874, int32(4), int32(1))
	if v877 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v866)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v869)
	F_s_lock(m, v874+int32(4), int32(_a_F_CheckpointerMain_6))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L5
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	v887 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[0]))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v887)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v887)+12)) = v888
	v890 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v887)+4)), uint32(v890))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v866)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v869)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	F_ConditionVariableBroadcast(m, v887+int32(36))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L5
	} else {
		goto L182
	}
L181:
	;
	goto L180
L182:
	;
	if v743 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v928 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[18])) = uint8(v928)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v869)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v866)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L5
	} else {
		goto L191
	}
L184:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[9])) = v682
	if v863 == int32(0) {
		goto L183
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	if v863 != 0 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v906 = int32(_a_F_CheckpointerMain_21)
	v908 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[37]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[37])) = v908 + int64(1)
	goto L183
L188:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[9])) = v682
	v914 = int32(_a_F_CheckpointerMain_22)
	v916 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[38]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[38])) = v916 + int64(1)
	goto L183
L189:
	;
	goto L190
L190:
	;
	v922 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[25])))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[9])) = v682 - v922 + int64(15)
	goto L183
L191:
	;
	v936 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	if v936 != 0 {
		v1103 = v860
		v1104 = v861
		goto L130
	} else {
		goto L192
	}
L192:
	;
	v938 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[24]))
	if v938 != 0 {
		v1103 = v860
		v1104 = v861
		goto L130
	} else {
		goto L193
	}
L193:
	;
	v943 = v860
	v944 = v861
	goto L137
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v951)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v954)
	F_CheckArchiveTimeout(m)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L5
	} else {
		goto L205
	}
L195:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[26])))
	if v962 == int32(1) {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	if v972 != 0 {
		goto L194
	} else {
		goto L200
	}
L197:
	;
	v967 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[28]))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v967)+308))
	v970 = base.B2i32(v968 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[26])) = uint8(v970)
	v972 = v970
	goto L199
L198:
	;
	v972 = int32(0)
	goto L199
L199:
	;
	goto L196
L200:
	;
	v974 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[39]))
	v978 = F_LWLockAcquire(m, v974+int32(_a_F_CheckpointerMain_23), int32(1))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L5
	} else {
		goto L201
	}
L201:
	;
	v981 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[40]))
	v982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v981)+2)))
	v984 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[39]))
	F_LWLockRelease(m, v984+int32(_a_F_CheckpointerMain_23))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L5
	} else {
		goto L202
	}
L202:
	;
	if v982 != int32(1) {
		goto L194
	} else {
		goto L203
	}
L203:
	;
	F_DisableLogicalDecoding(m)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L5
	} else {
		goto L204
	}
L204:
	;
	goto L194
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v951)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v954)
	F_pgstat_report_checkpointer(m)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L5
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v951)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v954)
	F_pgstat_report_wal(m, int32(1))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L5
	} else {
		goto L207
	}
L207:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[0]))
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v1010)+20))
	if v1011 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v954)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v951)
	v1076 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[21]))
	v1077 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1076))) = v1077
	v1082 = base.AtomicRmwOr32(m, v1077, int32(_a_F_CheckpointerMain_9), v1077)
	goto L223
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v951)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v954)
	v1015 = F_time(m)
	mBase = m.M
	v1017 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[25]))
	v1019 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[9]))
	v1021 = base.I32_wrap_i64(v1015 - v1019)
	if v1017 <= v1021 {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v1023 = v1017 - v1021
	v1025 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[41]))
	if v1025 <= int32(0) {
		v1054 = v1023
		goto L211
	} else {
		goto L212
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v954)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v951)
	v1061 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[21]))
	v1066 = F_WaitLatch(m, v1061, int32(41), v1054*int32(1000), int32(83886084))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L5
	} else {
		goto L222
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v951)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v954)
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[26])))
	if v1033 == int32(1) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	if v1043 != 0 {
		v1054 = v1023
		goto L211
	} else {
		goto L217
	}
L214:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[28]))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+308))
	v1041 = base.B2i32(v1039 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[26])) = uint8(v1041)
	v1043 = v1041
	goto L216
L215:
	;
	v1043 = int32(0)
	goto L216
L216:
	;
	goto L213
L217:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[41]))
	v1047 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10]))
	v1049 = base.I32_wrap_i64(v1015 - v1047)
	if v1045 <= v1049 {
		goto L208
	} else {
		goto L218
	}
L218:
	;
	v1051 = v1045 - v1049
	if v1023 < v1051 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v1053 = v1023
	goto L221
L220:
	;
	v1053 = v1051
	goto L221
L221:
	;
	v1054 = v1053
	goto L211
L222:
	;
	goto L208
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v951)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v954)
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L5
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v951)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v954)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L5
	} else {
		goto L225
	}
L225:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	if v1094 == int32(0) {
		v662 = v943
		v663 = v944
		goto L132
	} else {
		goto L226
	}
L226:
	;
	goto L133
L227:
	;
	v1116 = int32(_a_F_CheckpointerMain_15)
	v1118 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[32]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[32])) = v1118 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	v1123 = int32(1)
	v1124 = v1103 & v1123
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1124)
	v1127 = v1104 & v1123
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1127)
	F_ShutdownXLOG(m, int32(0), int64(0))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L5
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	v1169 = int32(1)
	v1170 = v1103 & v1169
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1170)
	v1173 = v1104 & v1169
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1173)
	v1176 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[21]))
	v1177 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1176))) = v1177
	v1182 = base.AtomicRmwOr32(m, v1177, int32(_a_F_CheckpointerMain_9), v1177)
	goto L237
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1127)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1124)
	F_pgstat_report_checkpointer(m)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L5
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1127)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1124)
	F_pgstat_report_wal(m, int32(1))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L5
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1127)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1124)
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[42])))
	if v1149 == int32(1) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22])) = int32(0)
	goto L229
L234:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v1153+int32(40)))) = int32(1)
	v1160 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[44]))
	v1162 = F_pgmem_kill(m, v1160, int32(10))
	mBase = m.M
	goto L236
L235:
	;
	goto L236
L236:
	;
	goto L233
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1173)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1170)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L5
	} else {
		goto L238
	}
L238:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[24]))
	if v1189 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	goto L242
L240:
	;
	goto L241
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1173)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1170)
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L5
	} else {
		goto L248
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1170)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1173)
	v1210 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[21]))
	v1214 = F_WaitLatch(m, v1210, int32(33), int32(0), int32(83886085))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L5
	} else {
		goto L244
	}
L243:
	;
	goto L241
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1170)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1173)
	v1220 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[21]))
	v1221 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1220))) = v1221
	v1226 = base.AtomicRmwOr32(m, v1221, int32(_a_F_CheckpointerMain_9), v1221)
	goto L245
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v447
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1173)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1170)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L5
	} else {
		goto L246
	}
L246:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[24]))
	if v1233 == int32(0) {
		goto L242
	} else {
		goto L247
	}
L247:
	;
	goto L243
L248:
	;
	goto L4
L249:
	;
	v1275 = int32(v1271)
	m.G0 = v17
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+4))
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1275)))
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1278)))
	if v17+int32(28) == v1281 {
		goto L252
	} else {
		goto L253
	}
L250:
	;
	m.ExcPending = 1
	goto L258
L251:
	;
	if v1285 != 0 {
		goto L255
	} else {
		goto L256
	}
L252:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+4))
	v1285 = v1283
	goto L254
L253:
	;
	v1285 = int32(0)
	goto L254
L254:
	;
	goto L251
L255:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v17)+204))
	v1287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)))
	v1288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)))
	v21 = v1285
	v22 = v1286
	v23 = v1277
	v26 = v1288
	v27 = v1287
	goto L1
L256:
	;
	goto L257
L257:
	;
	F___wasm_longjmp(m, v1278, v1277)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	return
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_char_increment(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v16 int64
	_ = v16
	v4 = int64(255)
	v7 = base.B2i32(l1&v4 == v4)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v7)
	v10 = int64(56)
	if l1&v4 == v4 {
		v16 = int64(0)
	} else {
		v16 = (l1<<(uint(v10)%64) + int64(72057594037927936)) >> (uint(v10) % 64)
	}
	return v16
}
func F_char_text(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	v3 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+24)))
	v5 = F_palloc(m, int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v3 < int32(0) {
			v11 = int32(92)
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v11)
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(32)
			v15 = int32(7)
			v17 = int32(48)
			v18 = v3&v15 | v17
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+7)) = uint8(v18)
			v25 = int32(base.Ui32(v3)>>(uint(int32(3))%32))&v15 | v17
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)) = uint8(v25)
			v32 = int32(base.Ui32(v3&int32(192))>>(uint(int32(6))%32)) | v17
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)) = uint8(v32)
			return base.I64_extend_i32_u(v5)
		} else {
			if v3 != 0 {
				*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v3)
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(20)
				return base.I64_extend_i32_u(v5)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(16)
				return base.I64_extend_i32_u(v5)
			}
		}
	}
}
func F_chareq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3))
}
func F_chareqfast(m *base.Module, l0 int64, l1 int64) int32 {
	var v4 int32
	_ = v4
	v4 = int32(255)
	return base.B2i32(base.I32_wrap_i64(l0)&v4 == base.I32_wrap_i64(l1)&v4)
}
func F_charrecv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pq_getmsgbyte(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_extend8_s(base.I64_extend_i32_u(v3))
	}
}
func F_charsend(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
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
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_pq_begintypsend(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		F_enlargeStringInfo(m, v6, int32(1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*uint8)(unsafe.Add(mBase, uint32(v16+v17))) = uint8(v8)
			v21 = v16 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v21
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = v21 << (uint(int32(2)) % 32)
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v24)
		}
	}
}
func F_checkCond(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v242 int32
	_ = v242
	F_check_stack_depth(m)
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
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_checkCond[0]))
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
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
		v242 = l3
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return base.B2i32(v242 == int32(0))
L8:
	;
	v30 = l0
	v31 = l1
	v32 = l2
	v33 = l3
	goto L9
L9:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+2)))
	if v49&int32(32) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return int32(0)
L11:
	;
	if v61 < v33 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v54 = int32(1)
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+4)))
	if v56 != 0 {
		v60 = v54
		v61 = v54
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+6)))
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+8)))
	v60 = v58
	v61 = v59
	goto L11
L15:
	;
	goto L14
L16:
	;
	v63 = v61
	goto L18
L17:
	;
	v63 = v33
	goto L18
L18:
	;
	if v63 < v60 {
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
	v68 = v31 - int32(1)
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30))))
	v74 = v30 + (v69+int32(7))&int32(_a_F_checkCond_0)
	if v63 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L10
L23:
	;
	v80 = v32
	v81 = v33
	v85 = int32(0)
	goto L26
L24:
	;
	v218 = v32
	v219 = v33
	goto L25
L25:
	;
	if v31 < int32(2) {
		v242 = v219
		goto L7
	} else {
		goto L49
	}
L26:
	;
	if base.Ui32(v85) < base.Ui32(v60) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v218 = v212
	v219 = v206
	goto L25
L28:
	;
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+2)))
	v106 = v104 & int32(16)
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+4)))
	if v107 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v98 = F_checkCond(m, v74, v68, v80, v81)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v98 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	return int32(1)
L32:
	;
	v205 = int32(1)
	v206 = v81 - v205
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80))))
	v212 = v80 + (v207+int32(9))&int32(_a_F_checkCond_0)
	v214 = v85 + v205
	if v214 != v63 {
		v80 = v212
		v81 = v206
		v85 = v214
		goto L26
	} else {
		goto L48
	}
L33:
	;
	if v106 != 0 {
		goto L22
	} else {
		goto L47
	}
L34:
	;
	v118 = v30 + int32(16)
	v123 = int32(0)
	goto L35
L35:
	;
	v133 = v118 + int32(7)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+6)))
	v136 = v134 & int32(2)
	v138 = v134 & int32(1)
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+4)))
	if v134&int32(4) != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if v106 != 0 {
		goto L32
	} else {
		goto L46
	}
L37:
	;
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+4)))
	v162 = v123 + int32(1)
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+4)))
	if base.Ui32(v162) < base.Ui32(v163) {
		v118 = v118 + (v153+int32(7))&int32(_a_F_checkCond_0) + int32(8)
		v123 = v162
		goto L35
	} else {
		goto L45
	}
L38:
	;
	v144 = F_compare_subnode(m, v80, v133, v139, v138, base.B2i32(v136 != int32(0)))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80))))
	v151 = F_ltree_label_match(m, v133, v139, v80+int32(2), v148, v138, base.B2i32(v136 != int32(0)))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	if v144 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	goto L33
L43:
	;
	if v151 != 0 {
		goto L33
	} else {
		goto L44
	}
L44:
	;
	goto L37
L45:
	;
	goto L36
L46:
	;
	return int32(0)
L47:
	;
	goto L32
L48:
	;
	goto L27
L49:
	;
	v30 = v74
	v31 = v68
	v32 = v218
	v33 = v219
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
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
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v12+v46<<(uint(int32(2))%32))))
			v13 = v50
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
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_errcode(m, int32(83886210))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v32
					F_errmsg(m, int32(_a_F_check_assignable_1), v9+int32(16))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v39 = F_plpgsql_scanner_errposition(m, l1, l2)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_assignable_2), int32(3574), int32(_a_F_check_assignable_3))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
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
		v54 = m.ExcPending
		if v54 != 0 {
			return
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v55
			F_errmsg_internal(m, int32(_a_F_check_assignable_4), v9)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_check_assignable_2), int32(3585), int32(_a_F_check_assignable_3))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v377 int32
	_ = v377
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v421 int32
	_ = v421
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v479 int32
	_ = v479
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v523 int32
	_ = v523
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v569 int32
	_ = v569
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v627 int32
	_ = v627
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 float64
	_ = v676
	var v678 int32
	_ = v678
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v775 int32
	_ = v775
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v910 int32
	_ = v910
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v940 int64
	_ = v940
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v1008 int32
	_ = v1008
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[0]))
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[1]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = F_pstrdup(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v22 + int32(16)
	return v1008
L2:
	;
	return int32(0)
L3:
	;
	v36 = F_SplitIdentifierString(m, v29, int32(44), v22+int32(12))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v36 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[3])) = v41
	goto L8
L6:
	;
	goto L7
L7:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v55 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	v47 = F_format_elog_string(m, int32(_a_F_check_datestyle_0), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[4])) = v47
	F_pfree(m, v29)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	F_list_free(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v1008 = v4
	goto L1
L12:
	;
	v976 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[3])) = v976
	goto L278
L13:
	;
	v920 = F_guc_malloc(m, int32(32))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L2
	} else {
		goto L264
	}
L14:
	;
	v910 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[3])) = v910
	goto L262
L15:
	;
	F_pfree(m, v29)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L2
	} else {
		goto L259
	}
L16:
	;
	v870 = int32(1)
	v872 = v25
	v873 = v27
	goto L15
L17:
	;
	goto L18
L18:
	;
	v59 = int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v60 <= int32(0) {
		v870 = v59
		v872 = v25
		v873 = v27
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v67 = v59
	v69 = v25
	v70 = v27
	v72 = v4
	v73 = v4
	v79 = v4
	goto L20
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v79<<(uint(int32(2))%32))))
	v90 = v86
	v91 = int32(_a_F_check_datestyle_1)
	goto L24
L21:
	;
	v870 = v847
	v872 = v849
	v873 = v850
	goto L15
L22:
	;
	v863 = v79 + int32(1)
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v863 < v864 {
		v67 = v847
		v69 = v849
		v70 = v850
		v72 = v852
		v73 = v853
		v79 = v863
		goto L20
	} else {
		goto L258
	}
L23:
	;
	if v128 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L24:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v94 == v95 {
		v117 = v94
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v128 = int32(0)
	goto L23
L26:
	;
	v119 = int32(1)
	if v117 != 0 {
		v90 = v90 + v119
		v91 = v91 + v119
		goto L24
	} else {
		goto L35
	}
L27:
	;
	if base.Ui32((v94-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v105 = v94 | int32(32)
	goto L30
L29:
	;
	v105 = v94
	goto L30
L30:
	;
	if base.Ui32((v95-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v114 = v95 | int32(32)
	goto L33
L32:
	;
	v114 = v95
	goto L33
L33:
	;
	if v105 == v114 {
		v117 = v105
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v128 = v105 - v114
	goto L23
L35:
	;
	goto L25
L36:
	;
	v133 = int32(1)
	v847 = (base.B2i32(v72 == int32(0)) | base.B2i32(v70 == v133)) & v67
	v849 = v69
	v850 = v133
	v852 = v133
	v853 = v73
	goto L22
L37:
	;
	goto L38
L38:
	;
	v142 = v86
	v143 = int32(_a_F_check_datestyle_2)
	goto L40
L39:
	;
	if v180 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L40:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v146 == v147 {
		v169 = v146
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v180 = int32(0)
	goto L39
L42:
	;
	v171 = int32(1)
	if v169 != 0 {
		v142 = v142 + v171
		v143 = v143 + v171
		goto L40
	} else {
		goto L51
	}
L43:
	;
	if base.Ui32((v146-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v157 = v146 | int32(32)
	goto L46
L45:
	;
	v157 = v146
	goto L46
L46:
	;
	if base.Ui32((v147-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v166 = v147 | int32(32)
	goto L49
L48:
	;
	v166 = v147
	goto L49
L49:
	;
	if v157 == v166 {
		v169 = v157
		goto L42
	} else {
		goto L50
	}
L50:
	;
	v180 = v157 - v166
	goto L39
L51:
	;
	goto L41
L52:
	;
	v185 = int32(2)
	v847 = (base.B2i32(v72 == int32(0)) | base.B2i32(v70 == v185)) & v67
	v849 = v69
	v850 = v185
	v852 = int32(1)
	v853 = v73
	goto L22
L53:
	;
	goto L54
L54:
	;
	v195 = v86
	v196 = int32(_a_F_check_datestyle_3)
	v197 = int32(8)
	goto L56
L55:
	;
	if v242 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L56:
	;
	if v197 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v242 = int32(0)
	goto L55
L58:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	if v200 == v201 {
		v223 = v200
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
	v225 = int32(1)
	if v223 != 0 {
		v195 = v195 + v225
		v196 = v196 + v225
		v197 = v197 - v225
		goto L56
	} else {
		goto L70
	}
L62:
	;
	if base.Ui32((v200-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v211 = v200 | int32(32)
	goto L65
L64:
	;
	v211 = v200
	goto L65
L65:
	;
	if base.Ui32((v201-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v220 = v201 | int32(32)
	goto L68
L67:
	;
	v220 = v201
	goto L68
L68:
	;
	if v211 == v220 {
		v223 = v211
		goto L61
	} else {
		goto L69
	}
L69:
	;
	v242 = v211 - v220
	goto L55
L70:
	;
	goto L60
L71:
	;
	v245 = int32(0)
	v847 = (base.B2i32(v72 == v245) | base.B2i32(v70 == v245)) & v67
	v849 = v69
	v850 = v245
	v852 = int32(1)
	v853 = v73
	goto L22
L72:
	;
	goto L73
L73:
	;
	v256 = v86
	v257 = int32(_a_F_check_datestyle_4)
	goto L75
L74:
	;
	if v294 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L75:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
	if v260 == v261 {
		v283 = v260
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v294 = int32(0)
	goto L74
L77:
	;
	v285 = int32(1)
	if v283 != 0 {
		v256 = v256 + v285
		v257 = v257 + v285
		goto L75
	} else {
		goto L86
	}
L78:
	;
	if base.Ui32((v260-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v271 = v260 | int32(32)
	goto L81
L80:
	;
	v271 = v260
	goto L81
L81:
	;
	if base.Ui32((v261-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v280 = v261 | int32(32)
	goto L84
L83:
	;
	v280 = v261
	goto L84
L84:
	;
	if v271 == v280 {
		v283 = v271
		goto L77
	} else {
		goto L85
	}
L85:
	;
	v294 = v271 - v280
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
	v310 = v86
	v311 = int32(_a_F_check_datestyle_5)
	goto L94
L90:
	;
	v304 = v69
	goto L92
L91:
	;
	v304 = int32(1)
	goto L92
L92:
	;
	v847 = (base.B2i32(v72 == int32(0)) | base.B2i32(v70 == int32(3))) & v67
	v849 = v304
	v850 = int32(3)
	v852 = int32(1)
	v853 = v73
	goto L22
L93:
	;
	if v348 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L94:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	if v314 == v315 {
		v337 = v314
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v348 = int32(0)
	goto L93
L96:
	;
	v339 = int32(1)
	if v337 != 0 {
		v310 = v310 + v339
		v311 = v311 + v339
		goto L94
	} else {
		goto L105
	}
L97:
	;
	if base.Ui32((v314-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v325 = v314 | int32(32)
	goto L100
L99:
	;
	v325 = v314
	goto L100
L100:
	;
	if base.Ui32((v315-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v334 = v315 | int32(32)
	goto L103
L102:
	;
	v334 = v315
	goto L103
L103:
	;
	if v325 == v334 {
		v337 = v325
		goto L96
	} else {
		goto L104
	}
L104:
	;
	v348 = v325 - v334
	goto L93
L105:
	;
	goto L95
L106:
	;
	v351 = int32(0)
	v847 = (base.B2i32(v73 == v351) | base.B2i32(v69 == v351)) & v67
	v849 = v351
	v850 = v70
	v852 = v72
	v853 = int32(1)
	goto L22
L107:
	;
	goto L108
L108:
	;
	v362 = v86
	v363 = int32(_a_F_check_datestyle_6)
	goto L111
L109:
	;
	v464 = v86
	v465 = int32(_a_F_check_datestyle_7)
	goto L146
L110:
	;
	if v400 != 0 {
		goto L123
	} else {
		goto L124
	}
L111:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
	if v366 == v367 {
		v389 = v366
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v400 = int32(0)
	goto L110
L113:
	;
	v391 = int32(1)
	if v389 != 0 {
		v362 = v362 + v391
		v363 = v363 + v391
		goto L111
	} else {
		goto L122
	}
L114:
	;
	if base.Ui32((v366-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v377 = v366 | int32(32)
	goto L117
L116:
	;
	v377 = v366
	goto L117
L117:
	;
	if base.Ui32((v367-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v386 = v367 | int32(32)
	goto L120
L119:
	;
	v386 = v367
	goto L120
L120:
	;
	if v377 == v386 {
		v389 = v377
		goto L113
	} else {
		goto L121
	}
L121:
	;
	v400 = v377 - v386
	goto L110
L122:
	;
	goto L112
L123:
	;
	v405 = v86
	v406 = int32(_a_F_check_datestyle_8)
	v407 = int32(4)
	goto L127
L124:
	;
	goto L125
L125:
	;
	v455 = int32(1)
	v847 = (base.B2i32(v73 == int32(0)) | base.B2i32(v69 == v455)) & v67
	v849 = v455
	v850 = v70
	v852 = v72
	v853 = v455
	goto L22
L126:
	;
	if v452 != 0 {
		goto L109
	} else {
		goto L142
	}
L127:
	;
	if v407 != 0 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v452 = int32(0)
	goto L126
L129:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405))))
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	if v410 == v411 {
		v433 = v410
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
	v435 = int32(1)
	if v433 != 0 {
		v405 = v405 + v435
		v406 = v406 + v435
		v407 = v407 - v435
		goto L127
	} else {
		goto L141
	}
L133:
	;
	if base.Ui32((v410-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v421 = v410 | int32(32)
	goto L136
L135:
	;
	v421 = v410
	goto L136
L136:
	;
	if base.Ui32((v411-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v430 = v411 | int32(32)
	goto L139
L138:
	;
	v430 = v411
	goto L139
L139:
	;
	if v421 == v430 {
		v433 = v421
		goto L132
	} else {
		goto L140
	}
L140:
	;
	v452 = v421 - v430
	goto L126
L141:
	;
	goto L131
L142:
	;
	goto L125
L143:
	;
	v612 = v86
	v613 = int32(_a_F_check_datestyle_9)
	goto L191
L144:
	;
	v603 = int32(2)
	v847 = (base.B2i32(v73 == int32(0)) | base.B2i32(v69 == v603)) & v67
	v849 = v603
	v850 = v70
	v852 = v72
	v853 = int32(1)
	goto L22
L145:
	;
	if v502 == int32(0) {
		goto L144
	} else {
		goto L158
	}
L146:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464))))
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465))))
	if v468 == v469 {
		v491 = v468
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v502 = int32(0)
	goto L145
L148:
	;
	v493 = int32(1)
	if v491 != 0 {
		v464 = v464 + v493
		v465 = v465 + v493
		goto L146
	} else {
		goto L157
	}
L149:
	;
	if base.Ui32((v468-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v479 = v468 | int32(32)
	goto L152
L151:
	;
	v479 = v468
	goto L152
L152:
	;
	if base.Ui32((v469-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v488 = v469 | int32(32)
	goto L155
L154:
	;
	v488 = v469
	goto L155
L155:
	;
	if v479 == v488 {
		v491 = v479
		goto L148
	} else {
		goto L156
	}
L156:
	;
	v502 = v479 - v488
	goto L145
L157:
	;
	goto L147
L158:
	;
	v508 = v86
	v509 = int32(_a_F_check_datestyle_10)
	goto L160
L159:
	;
	if v546 == int32(0) {
		goto L144
	} else {
		goto L172
	}
L160:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509))))
	if v512 == v513 {
		v535 = v512
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v546 = int32(0)
	goto L159
L162:
	;
	v537 = int32(1)
	if v535 != 0 {
		v508 = v508 + v537
		v509 = v509 + v537
		goto L160
	} else {
		goto L171
	}
L163:
	;
	if base.Ui32((v512-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v523 = v512 | int32(32)
	goto L166
L165:
	;
	v523 = v512
	goto L166
L166:
	;
	if base.Ui32((v513-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v532 = v513 | int32(32)
	goto L169
L168:
	;
	v532 = v513
	goto L169
L169:
	;
	if v523 == v532 {
		v535 = v523
		goto L162
	} else {
		goto L170
	}
L170:
	;
	v546 = v523 - v532
	goto L159
L171:
	;
	goto L161
L172:
	;
	v553 = v86
	v554 = int32(_a_F_check_datestyle_11)
	v555 = int32(7)
	goto L174
L173:
	;
	if v600 != 0 {
		goto L143
	} else {
		goto L189
	}
L174:
	;
	if v555 != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v600 = int32(0)
	goto L173
L176:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553))))
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v554))))
	if v558 == v559 {
		v581 = v558
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
	v583 = int32(1)
	if v581 != 0 {
		v553 = v553 + v583
		v554 = v554 + v583
		v555 = v555 - v583
		goto L174
	} else {
		goto L188
	}
L180:
	;
	if base.Ui32((v558-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v569 = v558 | int32(32)
	goto L183
L182:
	;
	v569 = v558
	goto L183
L183:
	;
	if base.Ui32((v559-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v578 = v559 | int32(32)
	goto L186
L185:
	;
	v578 = v559
	goto L186
L186:
	;
	if v569 == v578 {
		v581 = v569
		goto L179
	} else {
		goto L187
	}
L187:
	;
	v600 = v569 - v578
	goto L173
L188:
	;
	goto L178
L189:
	;
	goto L144
L190:
	;
	if v650 != 0 {
		goto L12
	} else {
		goto L203
	}
L191:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612))))
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613))))
	if v616 == v617 {
		v639 = v616
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v650 = int32(0)
	goto L190
L193:
	;
	v641 = int32(1)
	if v639 != 0 {
		v612 = v612 + v641
		v613 = v613 + v641
		goto L191
	} else {
		goto L202
	}
L194:
	;
	if base.Ui32((v616-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v627 = v616 | int32(32)
	goto L197
L196:
	;
	v627 = v616
	goto L197
L197:
	;
	if base.Ui32((v617-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v636 = v617 | int32(32)
	goto L200
L199:
	;
	v636 = v617
	goto L200
L200:
	;
	if v627 == v636 {
		v639 = v627
		goto L193
	} else {
		goto L201
	}
L201:
	;
	v650 = v627 - v636
	goto L190
L202:
	;
	goto L192
L203:
	;
	v651 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v651
	v655 = m.G0
	v657 = v655 - int32(80)
	m.G0 = v657
	v663 = F_find_option(m, int32(_a_F_check_datestyle_12), v651, v651, int32(21))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L2
	} else {
		goto L205
	}
L204:
	;
	v812 = F_guc_strdup(m, int32(15), v775)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L2
	} else {
		goto L240
	}
L205:
	;
	v665 = F_ConfigOptionIsVisible(m, v663)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L2
	} else {
		goto L206
	}
L206:
	;
	if v665 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v663)+24))
	switch v667 {
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
		v775 = v651
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L2
	} else {
		goto L235
	}
L210:
	;
	m.G0 = v657 + int32(80)
	goto L204
L211:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663)+116)))
	if v762 != 0 {
		goto L232
	} else {
		goto L233
	}
L212:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v663)+120))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v663)+104))
	if v690 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L213:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v663)+116))
	if v686 != 0 {
		goto L218
	} else {
		goto L219
	}
L214:
	;
	v676 = *(*float64)(unsafe.Add(mBase, uint32(v663)+144))
	*(*float64)(unsafe.Add(mBase, uint32(v657)+16)) = v676
	v678 = int32(_a_F_check_datestyle_13)
	v684 = F_pg_snprintf(m, v678, int32(256), int32(_a_F_check_datestyle_14), v657+int32(16))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L2
	} else {
		goto L217
	}
L215:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v663)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v657))) = v668
	v670 = int32(_a_F_check_datestyle_13)
	v674 = F_pg_snprintf(m, v670, int32(256), int32(_a_F_check_datestyle_15), v657)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L2
	} else {
		goto L216
	}
L216:
	;
	v775 = v670
	goto L210
L217:
	;
	v775 = v678
	goto L210
L218:
	;
	v688 = v686
	goto L220
L219:
	;
	v688 = int32(_a_F_check_datestyle_16)
	goto L220
L220:
	;
	v775 = v688
	goto L210
L221:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L2
	} else {
		goto L229
	}
L222:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v690)))
	if v693 == int32(0) {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v690)+4))
	if v696 == v689 {
		v775 = v693
		goto L210
	} else {
		goto L224
	}
L224:
	;
	v710 = v690
	goto L225
L225:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v710)+12))
	if v717 == int32(0) {
		goto L221
	} else {
		goto L227
	}
L226:
	;
	v775 = v717
	goto L210
L227:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v710)+16))
	if v689 != v720 {
		v710 = v710 + int32(12)
		goto L225
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	*(*int32)(unsafe.Add(mBase, uint32(v657)+36)) = v747
	*(*int32)(unsafe.Add(mBase, uint32(v657)+32)) = v689
	F_errmsg_internal(m, int32(_a_F_check_datestyle_17), v657+int32(32))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L2
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_check_datestyle_18), int32(2938), int32(_a_F_check_datestyle_19))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L2
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
	v763 = int32(_a_F_check_datestyle_20)
	goto L234
L233:
	;
	v763 = int32(_a_F_check_datestyle_21)
	goto L234
L234:
	;
	v775 = v763
	goto L210
L235:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L2
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v657)+64)) = int32(_a_F_check_datestyle_12)
	F_errmsg(m, int32(_a_F_check_datestyle_22), v657-int32(-64))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L2
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v657)+48)) = int32(_a_F_check_datestyle_23)
	v805 = F_errdetail(m, int32(_a_F_check_datestyle_24), v657+int32(48))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L2
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_check_datestyle_18), int32(_a_F_check_datestyle_25), int32(_a_F_check_datestyle_26))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L2
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v812
	if v812 != 0 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v72 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L242:
	;
	v819 = F_check_datestyle(m, v22+int32(8), v22+int32(4), l2)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L2
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	F_pfree(m, v29)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L2
	} else {
		goto L248
	}
L245:
	;
	if v819 != 0 {
		goto L241
	} else {
		goto L246
	}
L246:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	F_bms_free(m, v821)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L2
	} else {
		goto L247
	}
L247:
	;
	goto L244
L248:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	F_list_free(m, v826)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L2
	} else {
		goto L249
	}
L249:
	;
	goto L14
L250:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	v833 = v832
	goto L252
L251:
	;
	v833 = v70
	goto L252
L252:
	;
	if v73 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v829)+4))
	v837 = v836
	goto L255
L254:
	;
	v837 = v69
	goto L255
L255:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	F_bms_free(m, v838)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L2
	} else {
		goto L256
	}
L256:
	;
	F_bms_free(m, v829)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L2
	} else {
		goto L257
	}
L257:
	;
	v847 = v67
	v849 = v837
	v850 = v833
	v852 = v72
	v853 = v73
	goto L22
L258:
	;
	goto L21
L259:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	F_list_free(m, v887)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L2
	} else {
		goto L260
	}
L260:
	;
	if v870 != 0 {
		goto L13
	} else {
		goto L261
	}
L261:
	;
	goto L14
L262:
	;
	v916 = F_format_elog_string(m, int32(_a_F_check_datestyle_27), int32(0))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L2
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[4])) = v916
	v1008 = v4
	goto L1
L264:
	;
	if v920 == int32(0) {
		v1008 = v4
		goto L1
	} else {
		goto L265
	}
L265:
	;
	switch v873 - int32(1) {
	case 0:
		goto L270
	case 1:
		goto L269
	case 2:
		goto L268
	default:
		goto L267
	}
L266:
	;
	v942 = F_strlen(m, v920)
	mBase = m.M
	v943 = v942 + v920
	switch v872 {
	case 0:
		goto L274
	case 1:
		goto L273
	default:
		goto L272
	}
L267:
	;
	v937 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_datestyle[5])))
	*(*uint8)(unsafe.Add(mBase, uint32(v920)+8)) = uint8(v937)
	v940 = *(*int64)(unsafe.Add(mBase, _c_F_check_datestyle[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v920))) = v940
	goto L266
L268:
	;
	v931 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v920)+3)) = v931
	v934 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v920))) = v934
	goto L266
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v920))) = int32(_a_F_check_datestyle_28)
	goto L266
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v920))) = int32(_a_F_check_datestyle_29)
	goto L266
L271:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_bms_free(m, v962)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L2
	} else {
		goto L275
	}
L272:
	;
	v957 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_check_datestyle[9])))
	*(*uint16)(unsafe.Add(mBase, uint32(v943)+4)) = uint16(v957)
	v960 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v943))) = v960
	goto L271
L273:
	;
	v951 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_check_datestyle[11])))
	*(*uint16)(unsafe.Add(mBase, uint32(v943)+4)) = uint16(v951)
	v954 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v943))) = v954
	goto L271
L274:
	;
	v945 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_check_datestyle[13])))
	*(*uint16)(unsafe.Add(mBase, uint32(v943)+4)) = uint16(v945)
	v948 = *(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v943))) = v948
	goto L271
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v920
	v967 = F_guc_malloc(m, int32(8))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L2
	} else {
		goto L276
	}
L276:
	;
	if v967 == int32(0) {
		v1008 = v4
		goto L1
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v967)+4)) = v872
	*(*int32)(unsafe.Add(mBase, uint32(v967))) = v873
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v967
	v1008 = int32(1)
	goto L1
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v86
	v982 = F_format_elog_string(m, int32(_a_F_check_datestyle_30), v22)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L2
	} else {
		goto L279
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_datestyle[4])) = v982
	F_pfree(m, v29)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L2
	} else {
		goto L280
	}
L280:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	F_list_free(m, v987)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L2
	} else {
		goto L281
	}
L281:
	;
	v1008 = v4
	goto L1
}
func F_check_publications_origin_sequences(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = int32(25)
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L18
	} else {
		goto L85
	}
L2:
	;
	m.G0 = v11 + int32(96)
	return
L3:
	;
	v20 = l3
	v21 = int32(_a_F_check_publications_origin_sequences_0)
	goto L5
L4:
	;
	if v58 != 0 {
		goto L2
	} else {
		goto L17
	}
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v24 == v25 {
		v47 = v24
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v58 = int32(0)
	goto L4
L7:
	;
	v49 = int32(1)
	if v47 != 0 {
		v20 = v20 + v49
		v21 = v21 + v49
		goto L5
	} else {
		goto L16
	}
L8:
	;
	if base.Ui32((v24-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v35 = v24 | int32(32)
	goto L11
L10:
	;
	v35 = v24
	goto L11
L11:
	;
	if base.Ui32((v25-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = v25 | int32(32)
	goto L14
L13:
	;
	v44 = v25
	goto L14
L14:
	;
	if v35 == v44 {
		v47 = v35
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v58 = v35 - v44
	goto L4
L16:
	;
	goto L6
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_check_publications_origin_sequences[0]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
	v63 = m.T0[v62].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	if v63 < int32(_a_F_check_publications_origin_sequences_1) {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v68 = v11 + int32(80)
	F_initStringInfo(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_appendStringInfoString(m, v68, int32(_a_F_check_publications_origin_sequences_2))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	F_GetPublicationsStr(m, l1, v68, int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	F_appendStringInfoString(m, v68, int32(_a_F_check_publications_origin_sequences_3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	if int32(0) < l5 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v84 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_check_publications_origin_sequences[0]))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+60))
	v141 = m.T0[v140].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v134, int32(1), v11+int32(76))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L18
	} else {
		goto L42
	}
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l4+v84<<(uint(int32(2))%32))))
	v94 = F_get_rel_name(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L18
	} else {
		goto L31
	}
L29:
	;
	goto L27
L30:
	;
	v124 = v84 + int32(1)
	if v124 != l5 {
		v84 = v124
		goto L28
	} else {
		goto L41
	}
L31:
	;
	if v94 == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v98 = F_get_rel_namespace(m, v93)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v100 = F_get_namespace_name(m, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	if v100 == int32(0) {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	v104 = F_quote_literal_cstr(m, v100)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	v106 = F_quote_literal_cstr(m, v94)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v104
	F_appendStringInfo(m, v11+int32(80), int32(_a_F_check_publications_origin_sequences_4), v11+int32(48))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L18
	} else {
		goto L38
	}
L38:
	;
	F_pfree(m, v104)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	F_pfree(m, v106)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L18
	} else {
		goto L40
	}
L40:
	;
	goto L30
L41:
	;
	goto L29
L42:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	F_pfree(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v146 != int32(2) {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	v151 = F_MakeSingleTupleTableSlot(m, v149, int32(_a_F_check_publications_origin_sequences_5))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v156 = F_tuplestore_gettupleslot(m, v153, int32(1), int32(0), v151)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L18
	} else {
		goto L47
	}
L46:
	;
	F_ExecDropSingleTupleTableSlot(m, v151)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L18
	} else {
		goto L71
	}
L47:
	;
	if v156 == int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v164 = int32(0)
	goto L49
L49:
	;
	v169 = int32(*(*int16)(unsafe.Add(mBase, uint32(v151)+6)))
	if v169 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v187 == int32(0) {
		goto L46
	} else {
		goto L61
	}
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+16))
	m.T0[v174].(func(*base.Module, int32, int32))(m, v151, int32(1))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L18
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v151)+16))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v179 = F_text_to_cstring(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L18
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	m.T0[v182].(func(*base.Module, int32))(m, v151)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L18
	} else {
		goto L56
	}
L56:
	;
	v185 = F_makeString(m, v179)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L18
	} else {
		goto L57
	}
L57:
	;
	v187 = F_list_append_unique(m, v164, v185)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L18
	} else {
		goto L58
	}
L58:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v192 = F_tuplestore_gettupleslot(m, v189, int32(1), int32(0), v151)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L18
	} else {
		goto L59
	}
L59:
	;
	if v192 != 0 {
		v164 = v187
		goto L49
	} else {
		goto L60
	}
L60:
	;
	goto L50
L61:
	;
	v197 = v11 + int32(60)
	F_initStringInfo(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L18
	} else {
		goto L62
	}
L62:
	;
	F_GetPublicationsStr(m, v187, v197, int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L18
	} else {
		goto L63
	}
L63:
	;
	v205 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L18
	} else {
		goto L64
	}
L64:
	;
	if v205 == int32(0) {
		goto L46
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L18
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l6
	F_errmsg(m, int32(_a_F_check_publications_origin_sequences_6), v11+int32(16))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L18
	} else {
		goto L67
	}
L67:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v219
	F_errdetail_plural(m, int32(_a_F_check_publications_origin_sequences_7), int32(_a_F_check_publications_origin_sequences_8), v218, v11)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	F_errhint(m, int32(_a_F_check_publications_origin_sequences_9), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L18
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_check_publications_origin_sequences_10), int32(3194), int32(_a_F_check_publications_origin_sequences_11))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L18
	} else {
		goto L70
	}
L70:
	;
	goto L46
L71:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	if v244 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_pfree(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L18
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	if v247 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	F_tuplestore_end(m, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L18
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	if v250 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L78
L80:
	;
	F_FreeTupleDesc(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L18
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_pfree(m, v141)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L18
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	goto L2
L85:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L18
	} else {
		goto L86
	}
L86:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v273
	F_errmsg(m, int32(_a_F_check_publications_origin_sequences_12), v11+int32(32))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L18
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_check_publications_origin_sequences_10), int32(3158), int32(_a_F_check_publications_origin_sequences_11))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L18
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_restrict_nonsystem_relation_kind(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = F_pstrdup(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v195
L2:
	;
	return int32(0)
L3:
	;
	v21 = F_SplitIdentifierString(m, v14, int32(44), v11+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_check_restrict_nonsystem_relation_kind[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_restrict_nonsystem_relation_kind[1])) = v26
	goto L8
L6:
	;
	goto L7
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v40 == int32(0) {
		v158 = v4
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v32 = F_format_elog_string(m, int32(_a_F_check_restrict_nonsystem_relation_kind_0), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_restrict_nonsystem_relation_kind[2])) = v32
	F_pfree(m, v14)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_list_free(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v195 = v4
	goto L1
L12:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_check_restrict_nonsystem_relation_kind[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_restrict_nonsystem_relation_kind[1])) = v174
	goto L53
L13:
	;
	F_pfree(m, v14)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L49
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v43 <= int32(0) {
		v158 = v4
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v47 = int32(0)
	v53 = v4
	goto L16
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v47<<(uint(int32(2))%32))))
	v63 = v59
	v64 = int32(_a_F_check_restrict_nonsystem_relation_kind_1)
	goto L19
L17:
	;
	v158 = v147
	goto L13
L18:
	;
	if v101 != 0 {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v67 == v68 {
		v90 = v67
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v101 = int32(0)
	goto L18
L21:
	;
	v92 = int32(1)
	if v90 != 0 {
		v63 = v63 + v92
		v64 = v64 + v92
		goto L19
	} else {
		goto L30
	}
L22:
	;
	if base.Ui32((v67-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v78 = v67 | int32(32)
	goto L25
L24:
	;
	v78 = v67
	goto L25
L25:
	;
	if base.Ui32((v68-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v87 = v68 | int32(32)
	goto L28
L27:
	;
	v87 = v68
	goto L28
L28:
	;
	if v78 == v87 {
		v90 = v78
		goto L21
	} else {
		goto L29
	}
L29:
	;
	v101 = v78 - v87
	goto L18
L30:
	;
	goto L20
L31:
	;
	v105 = v59
	v106 = int32(_a_F_check_restrict_nonsystem_relation_kind_2)
	goto L35
L32:
	;
	v146 = int32(1)
	goto L33
L33:
	;
	v147 = v146 | v53
	v149 = v47 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v149 < v150 {
		v47 = v149
		v53 = v147
		goto L16
	} else {
		goto L48
	}
L34:
	;
	if v143 != 0 {
		goto L12
	} else {
		goto L47
	}
L35:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v109 == v110 {
		v132 = v109
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v143 = int32(0)
	goto L34
L37:
	;
	v134 = int32(1)
	if v132 != 0 {
		v105 = v105 + v134
		v106 = v106 + v134
		goto L35
	} else {
		goto L46
	}
L38:
	;
	if base.Ui32((v109-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v120 = v109 | int32(32)
	goto L41
L40:
	;
	v120 = v109
	goto L41
L41:
	;
	if base.Ui32((v110-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v129 = v110 | int32(32)
	goto L44
L43:
	;
	v129 = v110
	goto L44
L44:
	;
	if v120 == v129 {
		v132 = v120
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v143 = v120 - v129
	goto L34
L46:
	;
	goto L36
L47:
	;
	v146 = int32(2)
	goto L33
L48:
	;
	goto L17
L49:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_list_free(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v166 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v166
	if v166 == int32(0) {
		v195 = v4
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v158
	v195 = int32(1)
	goto L1
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v59
	v180 = F_format_elog_string(m, int32(_a_F_check_restrict_nonsystem_relation_kind_3), v11)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_restrict_nonsystem_relation_kind[2])) = v180
	F_pfree(m, v14)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_list_free(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v195 = v4
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
	var v37 int32
	_ = v37
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
	var v60 int32
	_ = v60
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
			v12 = F_check_functions_in_node(m, l0, int32(500), int32(0))
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
								v37 = F_errdetail(m, int32(_a_F_check_virtual_generated_security_walker_1), int32(0))
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
										F_errfinish(m, int32(_a_F_check_virtual_generated_security_walker_2), int32(3318), int32(_a_F_check_virtual_generated_security_walker_3))
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
										v60 = F_errdetail(m, int32(_a_F_check_virtual_generated_security_walker_6), int32(0))
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
												F_errfinish(m, int32(_a_F_check_virtual_generated_security_walker_2), int32(3334), int32(_a_F_check_virtual_generated_security_walker_3))
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
							v21 = F_expression_tree_walker_impl(m, l0, int32(501), l1)
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
			v21 = F_expression_tree_walker_impl(m, l0, int32(501), l1)
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
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v10&int32(1) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v153
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = int32(1)
	v25 = v13 + (int32(base.Ui32(v10)>>(uint(v14)%32))&int32(2047)+int32(base.Ui32(v10)>>(uint(int32(12))%32))+v14)&int32(_a_F_checkclass_str_0)
	v26 = int32(0)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if base.B2i32(l3 == v26)|base.B2i32(v28 == v26) == v26 {
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
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v37 = F_palloc_mul(m, int32(2), v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	if v28 != 0 {
		goto L21
	} else {
		goto L22
	}
L8:
	;
	return int32(0)
L9:
	;
	v41 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v41)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v37
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v46 = v25 + int32(2)
	v47 = v37
	v48 = v46
	v53 = v44
	goto L13
L11:
	;
	v78 = v37
	v87 = v37
	goto L12
L12:
	;
	v90 = (v78 - v87) >> (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v90
	if int32(0) < v90 {
		v153 = int32(1)
		goto L1
	} else {
		goto L19
	}
L13:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48))))
	if int32(base.Ui32(v56)>>(uint(int32(base.Ui32(v57)>>(uint(int32(14))%32)))%32))&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v78 = v69
	v87 = v77
	goto L12
L15:
	;
	v64 = v57 & int32(_a_F_checkclass_str_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v47))) = uint16(v64)
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v69 = v47 + int32(2)
	v70 = v66
	goto L17
L16:
	;
	v69 = v47
	v70 = v53
	goto L17
L17:
	;
	v72 = v48 + int32(2)
	if base.Ui32(v72) < base.Ui32(v46+v70<<(uint(int32(1))%32)) {
		v47 = v69
		v48 = v72
		v53 = v70
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	F_pfree(m, v87)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v96 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v96
	return v96
L21:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	if v102 == int32(0) {
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
		v153 = int32(1)
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
	v108 = v25 + int32(2)
	v113 = v108
	goto L27
L27:
	;
	v121 = int32(1)
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113))))
	if int32(base.Ui32(v28)>>(uint(int32(base.Ui32(v122)>>(uint(int32(14))%32)))%32))&v121 != 0 {
		v153 = v121
		goto L1
	} else {
		goto L29
	}
L28:
	;
	return int32(0)
L29:
	;
	v129 = v113 + int32(2)
	if base.Ui32(v129) < base.Ui32(v108+v102<<(uint(int32(1))%32)) {
		v113 = v129
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v25 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v136
	v141 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v141)
	return int32(1)
L32:
	;
	v147 = int32(2)
	goto L34
L33:
	;
	v147 = int32(1)
	goto L34
L34:
	;
	v153 = v147
	goto L1
}
func F_checkmatchall_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v213 int32
	_ = v213
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = m.T0[v15].(func(*base.Module) int32)(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v213 & int32(1)
L2:
	;
	return int32(0)
L3:
	;
	if v16 != 0 {
		v213 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_checkmatchall_recurse[0]))
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v26 = F_palloc_extended(m, int32(258), int32(2))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	if v26 == int32(0) {
		v213 = v4
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v30 = int32(0)
	base.MemoryFill(m, v26, v30, int32(258))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l1
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v34 == v30 {
		v195 = v4
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l2+v203<<(uint(int32(2))%32)))) = v26
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
	v213 = v195
	goto L1
L12:
	;
	v40 = v4
	v44 = v34
	v46 = v4
	goto L13
L13:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)))
	if v48 != int32(_a_F_checkmatchall_recurse_0) {
		v135 = v40
		v141 = v46
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v135&v141 == int32(0) {
		v195 = v135
		goto L11
	} else {
		goto L34
	}
L15:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	if v143 != 0 {
		v40 = v135
		v44 = v143
		v46 = v141
		goto L13
	} else {
		goto L33
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v51 == v52 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v54 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v54)
	v135 = v54
	v141 = v46
	goto L15
L18:
	;
	goto L19
L19:
	;
	if l1 == v51 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v135 = v40
	v141 = int32(1)
	goto L15
L21:
	;
	goto L22
L22:
	;
	v59 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v51)+24))
	if v60 != 0 {
		v195 = v59
		goto L11
	} else {
		goto L23
	}
L23:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l2+v61<<(uint(int32(2))%32))))
	if v65 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v68 = F_checkmatchall_recurse(m, l0, v51, l2)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L27
	}
L25:
	;
	v78 = v65
	goto L26
L26:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+256)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+257)))
	if v79 != v80 {
		v195 = v59
		goto L11
	} else {
		goto L29
	}
L27:
	;
	if v68 == int32(0) {
		v195 = v59
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2+v73<<(uint(int32(2))%32))))
	v78 = v77
	goto L26
L29:
	;
	v85 = v59
	goto L30
L30:
	;
	v94 = v85 | int32(1)
	v95 = v26 + v94
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v78))))
	v99 = v96 | v98
	*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v99)
	v102 = v85 | int32(2)
	v103 = v26 + v102
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v94))))
	v107 = v104 | v106
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v107)
	v110 = v85 | int32(3)
	v111 = v26 + v110
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v102))))
	v115 = v112 | v114
	*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v115)
	v118 = v85 + int32(4)
	v119 = v26 + v118
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v110))))
	v123 = v120 | v122
	*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v123)
	if v118 != int32(256) {
		v85 = v118
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+257)))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+257)))
	v129 = v127 | v128
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+257)) = uint8(v129)
	v135 = int32(1)
	v141 = v46
	goto L15
L32:
	;
	goto L31
L33:
	;
	goto L14
L34:
	;
	v151 = int32(0)
	goto L35
L35:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v26))))
	if v160 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if base.Ui32(int32(256)) < base.Ui32(v178) {
		goto L47
	} else {
		goto L48
	}
L37:
	;
	goto L36
L38:
	;
	v178 = v151
	goto L37
L39:
	;
	goto L40
L40:
	;
	if v151 == int32(256) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v195 = int32(1)
	goto L11
L42:
	;
	goto L43
L43:
	;
	v165 = v151 | int32(1)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v165))))
	if v167 != 0 {
		v178 = v165
		goto L37
	} else {
		goto L44
	}
L44:
	;
	v169 = v151 | int32(2)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v169))))
	if v171 != 0 {
		v178 = v169
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v173 = v151 | int32(3)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v173))))
	if v175 != 0 {
		v178 = v173
		goto L37
	} else {
		goto L46
	}
L46:
	;
	v151 = v151 + int32(4)
	goto L35
L47:
	;
	v195 = int32(1)
	goto L11
L48:
	;
	goto L49
L49:
	;
	v182 = int32(1)
	v184 = int32(257) - v178
	if v184 == int32(0) {
		v195 = v182
		goto L11
	} else {
		goto L50
	}
L50:
	;
	v188 = int32(1)
	base.MemoryFill(m, v178+v26+v188, v188, v184)
	v195 = v182
	goto L11
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
	var v29 int32
	_ = v29
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
	var v69 int32
	_ = v69
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
	v27 = int32(_a_F_chrnamed_0)
	v29 = int32(_a_F_chrnamed_1)
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
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	if v69 != 0 {
		v27 = v69
		v29 = v29 + int32(8)
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
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
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
