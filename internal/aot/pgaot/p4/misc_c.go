package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_CheckElement_3(m *base.Module, l0 float32) {
	var v8 int32
	_ = v8
	Fn13825(m, l0, int32(122), int32(_a_F_CheckElement_3_0), int32(_a_F_CheckElement_3_1), int32(117), int32(_a_F_CheckElement_3_2))
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v287 int32
	_ = v287
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v337 int32
	_ = v337
	var v345 int64
	_ = v345
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int64
	_ = v613
	var v615 int64
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v684 int64
	_ = v684
	var v688 int32
	_ = v688
	var v690 int64
	_ = v690
	var v696 int32
	_ = v696
	var v698 int64
	_ = v698
	var v702 int32
	_ = v702
	var v704 int64
	_ = v704
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v734 int32
	_ = v734
	var v742 int32
	_ = v742
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v763 int64
	_ = v763
	var v764 int32
	_ = v764
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v781 int64
	_ = v781
	var v782 int32
	_ = v782
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v847 int64
	_ = v847
	var v853 int32
	_ = v853
	var v855 int64
	_ = v855
	var v861 int64
	_ = v861
	var v867 int32
	_ = v867
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v914 int64
	_ = v914
	var v916 int32
	_ = v916
	var v918 int64
	_ = v918
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v946 int64
	_ = v946
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v975 int32
	_ = v975
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1013 int64
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1038 int32
	_ = v1038
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1142 int32
	_ = v1142
	var v1157 int32
	_ = v1157
	var v1158 int64
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
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
	v1157 = int32(m.ExcTag)
	v1158 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1157 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[0])) = int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	v40 = int32(1)
	v41 = v26 & v40
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v41)
	v44 = v27 & v40
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v44)
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v378 = v22
	v379 = v23
	goto L8
L8:
	;
	if v379 != 0 {
		goto L64
	} else {
		goto L65
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v51
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v44)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v41)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	v57 = int32(914)
	v59 = m.G0
	v61 = v59 - int32(32)
	m.G0 = v61
	switch int32(916) {
	case 0, 2:
		v71 = v57
		goto L11
	default:
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v44)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v41)
	v93 = int32(929)
	v95 = m.G0
	v97 = v95 - int32(32)
	m.G0 = v97
	switch int32(931) {
	case 0, 2:
		v107 = v93
		goto L17
	default:
		goto L18
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+12)) = v71
	F_sigemptyset(m, v61+int32(16))
	mBase = m.M
	goto L14
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[3])) = v57
	v71 = int32(_a_F_CheckpointerMain_0)
	goto L11
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = int32(268435456)
	v85 = F___sigaction(m, int32(1), v61+int32(12), int32(0))
	mBase = m.M
	m.G0 = v61 + int32(32)
	goto L10
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v44)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v41)
	v129 = int32(-2)
	v131 = m.G0
	v133 = v131 - int32(32)
	m.G0 = v133
	switch int32(0) {
	case 0, 2:
		v143 = v129
		goto L23
	default:
		goto L24
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v107
	F_sigemptyset(m, v97+int32(16))
	mBase = m.M
	goto L20
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[4])) = v93
	v107 = int32(_a_F_CheckpointerMain_0)
	goto L17
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+24)) = int32(268435456)
	v121 = F___sigaction(m, int32(2), v97+int32(12), int32(0))
	mBase = m.M
	m.G0 = v97 + int32(32)
	goto L16
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v44)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v41)
	v165 = int32(-2)
	v167 = m.G0
	v169 = v167 - int32(32)
	m.G0 = v169
	switch int32(0) {
	case 0, 2:
		v179 = v165
		goto L29
	default:
		goto L30
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+12)) = v143
	F_sigemptyset(m, v133+int32(16))
	mBase = m.M
	goto L26
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[5])) = v129
	v143 = int32(_a_F_CheckpointerMain_0)
	goto L23
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+24)) = int32(268435456)
	v157 = F___sigaction(m, int32(15), v133+int32(12), int32(0))
	mBase = m.M
	m.G0 = v133 + int32(32)
	goto L22
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v44)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v41)
	v201 = int32(-2)
	v203 = m.G0
	v205 = v203 - int32(32)
	m.G0 = v205
	switch int32(0) {
	case 0, 2:
		v215 = v201
		goto L35
	default:
		goto L36
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v169)+12)) = v179
	F_sigemptyset(m, v169+int32(16))
	mBase = m.M
	goto L32
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[6])) = v165
	v179 = int32(_a_F_CheckpointerMain_0)
	goto L29
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v169)+24)) = int32(268435456)
	v193 = F___sigaction(m, int32(14), v169+int32(12), int32(0))
	mBase = m.M
	m.G0 = v169 + int32(32)
	goto L28
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v44)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v41)
	v237 = int32(917)
	v239 = m.G0
	v241 = v239 - int32(32)
	m.G0 = v241
	switch int32(919) {
	case 0, 2:
		v251 = v237
		goto L41
	default:
		goto L42
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v205)+12)) = v215
	F_sigemptyset(m, v205+int32(16))
	mBase = m.M
	goto L38
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[7])) = v201
	v215 = int32(_a_F_CheckpointerMain_0)
	goto L35
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v205)+24)) = int32(268435456)
	v229 = F___sigaction(m, int32(13), v205+int32(12), int32(0))
	mBase = m.M
	m.G0 = v205 + int32(32)
	goto L34
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v44)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v41)
	v273 = int32(916)
	v275 = m.G0
	v277 = v275 - int32(32)
	m.G0 = v277
	switch int32(918) {
	case 0, 2:
		v287 = v273
		goto L47
	default:
		goto L48
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v251
	F_sigemptyset(m, v241+int32(16))
	mBase = m.M
	goto L44
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[8])) = v237
	v251 = int32(_a_F_CheckpointerMain_0)
	goto L41
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v241)+24)) = int32(268435456)
	v265 = F___sigaction(m, int32(10), v241+int32(12), int32(0))
	mBase = m.M
	m.G0 = v241 + int32(32)
	goto L40
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v44)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v41)
	v309 = int32(0)
	v311 = m.G0
	v313 = v311 - int32(32)
	m.G0 = v313
	switch int32(2) {
	case 0, 2:
		v323 = v309
		goto L53
	default:
		goto L54
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+12)) = v287
	F_sigemptyset(m, v277+int32(16))
	mBase = m.M
	goto L50
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[9])) = v273
	v287 = int32(_a_F_CheckpointerMain_0)
	goto L47
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+24)) = int32(268435456)
	v301 = F___sigaction(m, int32(12), v277+int32(12), int32(0))
	mBase = m.M
	m.G0 = v277 + int32(32)
	goto L46
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v44)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v41)
	v345 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10])) = v345
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[11])) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v41)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v44)
	F_before_shmem_exit(m, int32(930), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L5
	} else {
		goto L58
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313)+12)) = v323
	F_sigemptyset(m, v313+int32(16))
	mBase = m.M
	goto L55
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[12])) = v309
	v323 = int32(_a_F_CheckpointerMain_0)
	goto L53
L55:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313)+24)) = int32(268435457)
	v337 = F___sigaction(m, int32(17), v313+int32(12), int32(0))
	mBase = m.M
	m.G0 = v313 + int32(32)
	goto L52
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v41)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v44)
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[13]))
	v365 = F_AllocSetContextCreateInternal(m, v360, int32(_a_F_CheckpointerMain_1), int32(0), int32(_a_F_CheckpointerMain_2), int32(_a_F_CheckpointerMain_3))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[14])) = v365
	goto L60
L60:
	;
	v371 = v17 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = v17 + int32(28)
	goto L63
L61:
	;
	v378 = v365
	v379 = int32(0)
	goto L8
L63:
	;
	goto L61
L64:
	;
	v381 = int32(_a_F_CheckpointerMain_4)
	v383 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[15]))
	v384 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[15])) = v383 + v384
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[16])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	v392 = v26 & v384
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	v395 = v27 & v384
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	F_EmitErrorReport(m)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L5
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[17])) = v17 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	v521 = int32(1)
	v522 = v26 & v521
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v522)
	v525 = v27 & v521
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v525)
	F_pgmem_sigprocmask(m, int32(_a_F_CheckpointerMain_5), int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L5
	} else {
		goto L86
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	F_LWLockReleaseAll(m)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v410))) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	F_UnlockBuffers(m)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	F_smgrdestroyall(m)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[19])))
	if v450 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v455 = base.AtomicRmwXchg32(m, v452, int32(4), int32(1))
	if v455 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[14])) = v378
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	F_FlushErrorState(m)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L5
	} else {
		goto L84
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	v460 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	F_s_lock(m, v460+int32(4), int32(_a_F_CheckpointerMain_6), int32(294), int32(_a_F_CheckpointerMain_7))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L5
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v469)+12)) = v470
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v469)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v469)+16)) = v472 + int32(1)
	v476 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v469)+4)), uint32(v476))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	F_ConditionVariableBroadcast(m, v469+int32(36))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L5
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v487 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[19])) = uint8(v487)
	goto L78
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	F_MemoryContextReset(m, v378)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	v502 = int32(_a_F_CheckpointerMain_4)
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[15]))
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[15])) = v504 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v392)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v395)
	F_pg_usleep(m, int32(_a_F_CheckpointerMain_8))
	mBase = m.M
	goto L66
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v522)
	F_SyncRepUpdateSyncStandbysDefined(m)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v522)
	F_UpdateFullPageWrites(m)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v522)
	v546 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	if v546 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v522)
	F_errmsg_internal(m, int32(_a_F_CheckpointerMain_9), int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L5
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[20]))
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v564)+64)) = v566
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v522)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v525)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	v572 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v572))) = int32(0)
	goto L95
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v522)
	F_errfinish(m, int32(_a_F_CheckpointerMain_6), int32(1390), int32(_a_F_CheckpointerMain_10))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v522)
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v522)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	v586 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[23]))
	if v586 != 0 {
		v998 = v26
		v999 = v27
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v1007 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[24])) = uint8(v1007)
	v1010 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[23]))
	if v1010 != 0 {
		goto L184
	} else {
		goto L185
	}
L99:
	;
	v593 = v26
	v594 = v27
	goto L100
L100:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[25]))
	if v602 != 0 {
		v998 = v593
		v999 = v594
		goto L98
	} else {
		goto L102
	}
L101:
	;
	v998 = v882
	v999 = v883
	goto L98
L102:
	;
	v604 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v604)+20))
	v606 = int32(1)
	v607 = v594 & v606
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v607)
	v610 = v593 & v606
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v610)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	v613 = F_time(m)
	mBase = m.M
	v615 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10]))
	v617 = base.I32_wrap_i64(v613 - v615)
	v619 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[26]))
	v620 = base.B2i32(v619 <= v617)
	if v620|v605 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v607)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v610)
	v627 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[27])))
	if v627 == int32(1) {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v882 = v593
	v883 = v594
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	v889 = int32(1)
	v890 = v883 & v889
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v890)
	v893 = v882 & v889
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v893)
	F_CheckArchiveTimeout(m)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L5
	} else {
		goto L162
	}
L106:
	;
	v639 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v642 = base.AtomicRmwXchg32(m, v639, int32(4), int32(1))
	if v642 != 0 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v632 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[28]))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v632)+316))
	v635 = base.B2i32(v633 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[27])) = uint8(v635)
	v637 = v635
	goto L109
L108:
	;
	v637 = int32(0)
	goto L109
L109:
	;
	goto L106
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v610)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v607)
	v647 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	F_s_lock(m, v647+int32(4), int32(_a_F_CheckpointerMain_6), int32(414), int32(_a_F_CheckpointerMain_7))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L5
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v656 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v656)+20))
	v658 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v656)+20)) = v658
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v656)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v656)+8)) = v660 + int32(1)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v656)+4)), uint32(v658))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v607)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v610)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	F_ConditionVariableBroadcast(m, v656+int32(24))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L5
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	v676 = int32(0)
	v678 = base.B2i32(v657&int32(2) == v676) & v637
	if base.B2i32(v605 == v676)&v620 != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v678|base.B2i32(v657&int32(128) == int32(0)) != 0 {
		goto L126
	} else {
		goto L127
	}
L116:
	;
	if v678 != 0 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	if v605 == int32(0) {
		goto L115
	} else {
		goto L122
	}
L119:
	;
	v682 = int32(_a_F_CheckpointerMain_11)
	v684 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[29]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[29])) = v684 + int64(1)
	goto L115
L120:
	;
	goto L121
L121:
	;
	v688 = int32(_a_F_CheckpointerMain_12)
	v690 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[30]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[30])) = v690 + int64(1)
	goto L115
L122:
	;
	if v678 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v696 = int32(_a_F_CheckpointerMain_13)
	v698 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[31]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[31])) = v698 + int64(1)
	goto L115
L124:
	;
	goto L125
L125:
	;
	v702 = int32(_a_F_CheckpointerMain_14)
	v704 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[32]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[32])) = v704 + int64(1)
	goto L115
L126:
	;
	if v619 <= v617 {
		goto L134
	} else {
		goto L135
	}
L127:
	;
	v714 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[33]))
	if v714 <= v617 {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v607)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v610)
	v721 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	if v721 == int32(0) {
		goto L126
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v607)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v610)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v617
	F_errmsg_plural(m, int32(_a_F_CheckpointerMain_15), int32(_a_F_CheckpointerMain_16), v617, v17+int32(16))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L5
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v607)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v610)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(_a_F_CheckpointerMain_17)
	F_errhint(m, int32(_a_F_CheckpointerMain_18), v17)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v607)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v610)
	F_errfinish(m, int32(_a_F_CheckpointerMain_6), int32(462), int32(_a_F_CheckpointerMain_7))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	goto L126
L134:
	;
	v753 = int32(256)
	goto L136
L135:
	;
	v753 = int32(0)
	goto L136
L136:
	;
	v754 = v657 | v753
	v756 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[19])) = uint8(v756)
	if v678 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	v800 = int32(1)
	v801 = v796 & v800
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v801)
	v804 = v795 & v800
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v804)
	F_smgrdestroyall(m)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L5
	} else {
		goto L145
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v607)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v610)
	v763 = F_GetInsertRecPtr(m)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L5
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v607)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v610)
	v781 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L5
	} else {
		goto L143
	}
L141:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[34])) = v613
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[35])) = v763
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[36])) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v610)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v607)
	v775 = F_CreateCheckPoint(m, v754)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L5
	} else {
		goto L142
	}
L142:
	;
	v795 = v593
	v796 = v775
	v798 = v775
	goto L137
L143:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[34])) = v613
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[35])) = v781
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[36])) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v610)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v607)
	v793 = F_CreateRestartPoint(m, v754)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L5
	} else {
		goto L144
	}
L144:
	;
	v795 = v793
	v796 = v594
	v798 = v793
	goto L137
L145:
	;
	v809 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v812 = base.AtomicRmwXchg32(m, v809, int32(4), int32(1))
	if v812 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v804)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v801)
	v817 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	F_s_lock(m, v817+int32(4), int32(_a_F_CheckpointerMain_6), int32(495), int32(_a_F_CheckpointerMain_7))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L5
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v826)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v826)+12)) = v827
	v829 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v826)+4)), uint32(v829))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v801)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v804)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	F_ConditionVariableBroadcast(m, v826+int32(36))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L5
	} else {
		goto L150
	}
L149:
	;
	goto L148
L150:
	;
	if v678 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v867 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[19])) = uint8(v867)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v804)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v801)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L5
	} else {
		goto L159
	}
L152:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10])) = v613
	if v798 == int32(0) {
		goto L151
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	if v798 != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v845 = int32(_a_F_CheckpointerMain_19)
	v847 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[37]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[37])) = v847 + int64(1)
	goto L151
L156:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10])) = v613
	v853 = int32(_a_F_CheckpointerMain_20)
	v855 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[38]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[38])) = v855 + int64(1)
	goto L151
L157:
	;
	goto L158
L158:
	;
	v861 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[26])))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10])) = v613 - v861 + int64(15)
	goto L151
L159:
	;
	v875 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[23]))
	if v875 != 0 {
		v998 = v795
		v999 = v796
		goto L98
	} else {
		goto L160
	}
L160:
	;
	v877 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[25]))
	if v877 != 0 {
		v998 = v795
		v999 = v796
		goto L98
	} else {
		goto L161
	}
L161:
	;
	v882 = v795
	v883 = v796
	goto L105
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v890)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v893)
	F_pgstat_report_checkpointer(m)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L5
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v890)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v893)
	F_pgstat_report_wal(m, int32(1))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L5
	} else {
		goto L164
	}
L164:
	;
	v909 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[1]))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v909)+20))
	if v910 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v893)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v890)
	v975 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v975))) = int32(0)
	goto L180
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v890)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v893)
	v914 = F_time(m)
	mBase = m.M
	v916 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[26]))
	v918 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[10]))
	v920 = base.I32_wrap_i64(v914 - v918)
	if v916 <= v920 {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v922 = v916 - v920
	v924 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[39]))
	if v924 <= int32(0) {
		v953 = v922
		goto L168
	} else {
		goto L169
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v893)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v890)
	v960 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	v965 = F_WaitLatch(m, v960, int32(41), v953*int32(1000), int32(83886084))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L5
	} else {
		goto L179
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v890)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v893)
	v932 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[27])))
	if v932 == int32(1) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v942 != 0 {
		v953 = v922
		goto L168
	} else {
		goto L174
	}
L171:
	;
	v937 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[28]))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v937)+316))
	v940 = base.B2i32(v938 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[27])) = uint8(v940)
	v942 = v940
	goto L173
L172:
	;
	v942 = int32(0)
	goto L173
L173:
	;
	goto L170
L174:
	;
	v944 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[39]))
	v946 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[11]))
	v948 = base.I32_wrap_i64(v914 - v946)
	if v944 <= v948 {
		goto L165
	} else {
		goto L175
	}
L175:
	;
	v950 = v944 - v948
	if v922 < v950 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v952 = v922
	goto L178
L177:
	;
	v952 = v950
	goto L178
L178:
	;
	v953 = v952
	goto L168
L179:
	;
	goto L165
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v890)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v893)
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L5
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v890)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v893)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L5
	} else {
		goto L182
	}
L182:
	;
	v989 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[23]))
	if v989 == int32(0) {
		v593 = v882
		v594 = v883
		goto L100
	} else {
		goto L183
	}
L183:
	;
	goto L101
L184:
	;
	v1011 = int32(_a_F_CheckpointerMain_14)
	v1013 = *(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[32]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckpointerMain[32])) = v1013 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	v1018 = int32(1)
	v1019 = v998 & v1018
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1019)
	v1022 = v999 & v1018
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1022)
	v1024 = int32(0)
	F_ShutdownXLOG(m, v1024, v1024)
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L5
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	v1064 = int32(1)
	v1065 = v998 & v1064
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1065)
	v1068 = v999 & v1064
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1068)
	v1071 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1071))) = int32(0)
	goto L194
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1022)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1019)
	F_pgstat_report_checkpointer(m)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1022)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1019)
	F_pgstat_report_wal(m, int32(1))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L5
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1022)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1019)
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckpointerMain[40])))
	if v1044 == int32(1) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[23])) = int32(0)
	goto L186
L191:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[41]))
	*(*int32)(unsafe.Add(mBase, uint32(v1048+int32(36)))) = int32(1)
	v1055 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[42]))
	v1057 = F_pgmem_kill(m, v1055, int32(10))
	mBase = m.M
	goto L193
L192:
	;
	goto L193
L193:
	;
	goto L190
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1068)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1065)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L5
	} else {
		goto L195
	}
L195:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[25]))
	if v1080 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	goto L199
L197:
	;
	goto L198
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1068)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1065)
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L5
	} else {
		goto L205
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1065)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1068)
	v1101 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	v1105 = F_WaitLatch(m, v1101, int32(33), int32(0), int32(83886085))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L5
	} else {
		goto L201
	}
L200:
	;
	goto L198
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1065)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1068)
	v1111 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1111))) = int32(0)
	goto L202
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v378
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)) = uint8(v1068)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)) = uint8(v1065)
	F_ProcessCheckpointerInterrupts(m)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L5
	} else {
		goto L203
	}
L203:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerMain[25]))
	if v1120 == int32(0) {
		goto L199
	} else {
		goto L204
	}
L204:
	;
	goto L200
L205:
	;
	goto L4
L206:
	;
	v1162 = int32(v1158)
	m.G0 = v17
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1162)+4))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1162)))
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1165)))
	if v17+int32(28) == v1168 {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	m.ExcPending = 1
	goto L215
L208:
	;
	if v1172 != 0 {
		goto L212
	} else {
		goto L213
	}
L209:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+4))
	v1172 = v1170
	goto L211
L210:
	;
	v1172 = int32(0)
	goto L211
L211:
	;
	goto L208
L212:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v17)+204))
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+203)))
	v1175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+202)))
	v21 = v1172
	v22 = v1173
	v23 = v1164
	v26 = v1175
	v27 = v1174
	goto L1
L213:
	;
	goto L214
L214:
	;
	F___wasm_longjmp(m, v1165, v1164)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	return
L216:
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
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ClosePipeStream[0]))
	v8 = v6 - int32(1)
	if int32(0) <= v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ClosePipeStream[1]))
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
	F_errmsg_internal(m, int32(_a_F_ClosePipeStream_0), int32(0))
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
	F_errfinish(m, int32(_a_F_ClosePipeStream_1), int32(3076), int32(_a_F_ClosePipeStream_2))
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
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
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
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
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
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_CommentObject[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15
	v18 = *(*int64)(unsafe.Add(mBase, _c_F_CommentObject[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 == int32(9) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v12 + int32(48)
	return
L2:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v233 == int32(0) {
		goto L1
	} else {
		goto L58
	}
L3:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_CreateComments(m, v219, v220, v221, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L10
	} else {
		goto L57
	}
L4:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v117 = int32(0)
	v118 = m.G0
	v120 = v118 - int32(128)
	m.G0 = v120
	v122 = int32(1)
	if v116 == v117 {
		v143 = v122
		v144 = v3
		goto L29
	} else {
		goto L30
	}
L5:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+48))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+119)))
	v78 = v76 - int32(99)
	if int32(1)<<(uint(v78)%32)&int32(_a_F_CommentObject_3) != 0 {
		goto L20
	} else {
		goto L21
	}
L6:
	;
	v58 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L15
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v26 = F_get_database_oid(m, v24, int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v32 = v20
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_get_object_address(m, l0, v32, v33, v12+int32(44), int32(4), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L13
	}
L10:
	;
	return
L11:
	;
	if v26 == int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v32 = v30
	goto L9
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_CommentObject[2]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v46
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	F_check_object_ownership(m, v41, v43, v12+int32(32), v42, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v53 - int32(6) {
	case 0:
		goto L5
	default:
		goto L3
	case 3, 27, 36:
		goto L4
	}
L15:
	;
	if v58 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v24
	F_errmsg(m, int32(_a_F_CommentObject_0), v12)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_CommentObject_1), int32(61), int32(_a_F_CommentObject_2))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	v86 = base.B2i32(base.Ui32(v78) <= base.Ui32(int32(19)))
	goto L22
L21:
	;
	v86 = int32(0)
	goto L22
L22:
	;
	if v86 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v95 + int32(4)
	F_errmsg(m, int32(_a_F_CommentObject_4), v12+int32(16))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+48))
	v106 = int32(*(*int8)(unsafe.Add(mBase, uint32(v105)+119)))
	F_errdetail_relkind_not_supported(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_CommentObject_1), int32(103), int32(_a_F_CommentObject_2))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
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
	v146 = v120 + int32(32)
	F_ScanKeyInit(m, v146, int32(1), int32(3), int32(184), v114)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L10
	} else {
		goto L33
	}
L30:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v125 == int32(0) {
		v143 = v122
		v144 = v3
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+18)) = uint8(v128)
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+16)) = uint16(v128)
	v133 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+12)) = uint16(v133)
	v135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+14)) = uint8(v135)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+24)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v120)+20)) = v114
	v140 = F_cstring_to_text(m, v116)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+28)) = v140
	v143 = v128
	v144 = v135
	goto L29
L33:
	;
	F_ScanKeyInit(m, v120+int32(80), int32(2), int32(3), int32(184), v115)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	v161 = F_table_open(m, int32(2396), int32(3))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L10
	} else {
		goto L36
	}
L35:
	;
	F_systable_endscan(m, v167)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L10
	} else {
		goto L46
	}
L36:
	;
	v167 = F_systable_beginscan(m, v161, int32(2397), int32(1), int32(0), int32(2), v146)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	v169 = F_systable_getnext(m, v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	if v169 == int32(0) {
		v190 = v117
		goto L35
	} else {
		goto L39
	}
L39:
	;
	if v143 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_simple_heap_delete(m, v161, v169+int32(4))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L10
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v161)+52))
	v186 = F_heap_modify_tuple(m, v169, v179, v120+int32(20), v120+int32(16), v120+int32(12))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L10
	} else {
		goto L44
	}
L43:
	;
	v190 = v117
	goto L35
L44:
	;
	F_CatalogTupleUpdate(m, v161, v169+int32(4), v186)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	v190 = v186
	goto L35
L46:
	;
	v193 = int32(0)
	if base.B2i32(v144 == v193)|base.B2i32(v190 != v193) == v193 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v161)+52))
	v205 = F_heap_form_tuple(m, v200, v120+int32(20), v120+int32(16))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L10
	} else {
		goto L50
	}
L48:
	;
	v209 = v190
	goto L49
L49:
	;
	if v209 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	F_CatalogTupleInsert(m, v161, v205)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	v209 = v205
	goto L49
L52:
	;
	F_pfree(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L10
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	F_relation_close(m, v161, int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L10
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	m.G0 = v120 + int32(128)
	goto L2
L57:
	;
	goto L2
L58:
	;
	F_relation_close(m, v233, int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
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
	var v56 int32
	_ = v56
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
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
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v62
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v65 != 0 {
		goto L28
	} else {
		goto L29
	}
L2:
	;
	v61 = l1
	v62 = v14
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
	v56 = F_AllocSetContextCreateInternal(m, v15, int32(_a_F_CompleteCachedPlan_0), int32(0), int32(1024), int32(_a_F_CompleteCachedPlan_1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L25
	} else {
		goto L26
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0])) = l2
	v61 = l1
	v62 = l2
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
	*(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0])) = v56
	v59 = F_copyObjectImpl(m, l1)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v61 = v59
	v62 = v56
	goto L1
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0])) = v15
	if int32(0) < l4 {
		goto L47
	} else {
		goto L48
	}
L29:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v66 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	F_extract_query_dependencies(m, v61, l0-int32(-64), l0+int32(68), l0+int32(85))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L25
	} else {
		goto L44
	}
L31:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	switch v70 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v74 = int32(1)
		goto L35
	default:
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v75 == int32(0) {
		goto L28
	} else {
		goto L38
	}
L34:
	;
	if v74 != 0 {
		goto L30
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v74 = int32(0)
	goto L35
L37:
	;
	goto L28
L38:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v79 != int32(6) {
		v94 = int32(1)
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v94&int32(1) == int32(0) {
		goto L28
	} else {
		goto L43
	}
L40:
	;
	goto L39
L41:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+28))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v86 = v84 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v86) {
		v94 = int32(0)
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v94 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v86)) % 64)))
	goto L40
L43:
	;
	goto L30
L44:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v109
	v112 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[2])))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v112)
	v114 = F_GetSearchPathMatcher(m, v62)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v114
	goto L28
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v9)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l4
	v140 = F_ChoosePortalStrategy(m, v61)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L25
	} else {
		goto L56
	}
L47:
	;
	v123 = l4 << (uint(int32(2)) % 32)
	v124 = F_palloc(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v124
	if v123 == int32(0) {
		goto L46
	} else {
		goto L51
	}
L51:
	;
	base.MemoryCopy(m, v124, l3, v123)
	goto L46
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v180
	*(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0])) = v14
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v193)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+93)) = uint8(v193)
	return
L53:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+28))
	v176 = F_UtilityTupleDescriptor(m, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L25
	} else {
		goto L62
	}
L54:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v150 = int32(0)
	goto L58
L55:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+76))
	v145 = F_ExecCleanTypeFromTL(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L25
	} else {
		goto L57
	}
L56:
	;
	switch v140 {
	case 0, 2:
		goto L55
	case 1:
		goto L54
	case 3:
		goto L53
	default:
		v180 = int32(0)
		goto L52
	}
L57:
	;
	v180 = v145
	goto L52
L58:
	;
	v163 = int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v147+v150<<(uint(int32(2))%32))))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+24)))
	if v167 != v163 {
		v150 = v150 + v163
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166)+96))
	v171 = F_ExecCleanTypeFromTL(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L25
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v180 = v171
	goto L52
L62:
	;
	v180 = v176
	goto L52
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
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
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v402 int32
	_ = v402
	var v412 int32
	_ = v412
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
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
	v444 = m.ExcPending
	if v444 != 0 {
		goto L7
	} else {
		goto L99
	}
L2:
	;
	m.G0 = v19 + int32(16)
	return v412
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v24 == int32(0) {
		v412 = v2
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
	F_errmsg(m, int32(_a_F_CopyReadAttributesCSV_0), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesCSV_1), int32(1837), int32(_a_F_CopyReadAttributesCSV_2))
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
	v73 = v69
	v75 = v67
	v77 = v2
	goto L18
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v66 = v65
	goto L15
L17:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+268)) = v390 - v402
	v412 = v393
	goto L2
L18:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v87 <= v77 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+52))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L7
	} else {
		goto L94
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v87 << (uint(int32(1)) % 32)
	v94 = F_repalloc(m, v86, v87<<(uint(int32(3))%32))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	v97 = v86
	goto L22
L22:
	;
	v99 = v77 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v97+v99))) = v73
	if base.Ui32(v75) < base.Ui32(v68) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v94
	v97 = v94
	goto L22
L24:
	;
	goto L19
L25:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v349+v99))) = int32(0)
	v354 = v77 + int32(1)
	if v202 != 0 {
		v73 = v193
		v75 = v194
		v77 = v354
		goto L18
	} else {
		goto L93
	}
L26:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v274 != 0 {
		goto L71
	} else {
		goto L72
	}
L27:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v208 == v201 {
		goto L53
	} else {
		goto L54
	}
L28:
	;
	v105 = v75
	v108 = v73
	v110 = int32(0)
	goto L31
L29:
	;
	goto L30
L30:
	;
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v186)
	v193 = v73 + int32(1)
	v194 = v75
	v196 = v73
	v201 = v186
	v202 = v186
	goto L27
L31:
	;
	v121 = v105 + int32(1)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v123 = base.B2i32(v122 == v50)
	if v123 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v180 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v180)
	v260 = v160
	v261 = v145 + int32(1)
	v262 = v145
	v267 = v160 - v75
	v268 = v180
	goto L26
L33:
	;
	v143 = v121
	v145 = v108
	goto L43
L34:
	;
	if v122 == v48 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v132 = v105
	v133 = v108
	goto L36
L36:
	;
	v134 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v134)
	v136 = v132 - v75
	v138 = v133 + int32(1)
	if v110 == v134 {
		v193 = v138
		v194 = v121
		v196 = v133
		v201 = v136
		v202 = v123
		goto L27
	} else {
		goto L42
	}
L37:
	;
	if base.Ui32(v121) < base.Ui32(v68) {
		goto L33
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v122)
	v130 = v108 + int32(1)
	if base.Ui32(v121) < base.Ui32(v68) {
		v105 = v121
		v108 = v130
		goto L31
	} else {
		goto L41
	}
L40:
	;
	goto L1
L41:
	;
	v132 = v121
	v133 = v130
	goto L36
L42:
	;
	v260 = v121
	v261 = v138
	v262 = v133
	v267 = v136
	v268 = v123
	goto L26
L43:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v160 = v143 + int32(1)
	if base.B2i32(v46 != v157)|base.B2i32(base.Ui32(v68) <= base.Ui32(v160)) != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if base.Ui32(v160) < base.Ui32(v68) {
		v105 = v160
		v108 = v145
		v110 = int32(1)
		goto L31
	} else {
		goto L52
	}
L45:
	;
	goto L44
L46:
	;
	if base.Ui32(v174) < base.Ui32(v68) {
		v143 = v174
		v145 = v145 + int32(1)
		goto L43
	} else {
		goto L51
	}
L47:
	;
	if v157 == v48 {
		goto L45
	} else {
		goto L50
	}
L48:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if base.B2i32(v46 != v163)&base.B2i32(v163 != v48) != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v163)
	v174 = v143 + int32(2)
	goto L46
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v157)
	v174 = v160
	goto L46
L51:
	;
	goto L1
L52:
	;
	goto L32
L53:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v201 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	goto L55
L55:
	;
	v260 = v194
	v261 = v193
	v262 = v196
	v267 = v201
	v268 = v202
	goto L26
L56:
	;
	if v255 == int32(0) {
		goto L25
	} else {
		goto L69
	}
L57:
	;
	v255 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v216 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v217 = v75
	v218 = v210
	v219 = v201
	v220 = v216
	goto L64
L61:
	;
	v243 = v210
	v247 = int32(0)
	goto L62
L62:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v255 = v247 - v248
	goto L56
L63:
	;
	v243 = v238
	v247 = v240
	goto L62
L64:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if base.B2i32(v220 != v222)|base.B2i32(v222 == int32(0)) != 0 {
		v238 = v218
		v240 = v220
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v238 = v232
	v240 = int32(0)
	goto L63
L66:
	;
	v228 = v219 - int32(1)
	if v228 == int32(0) {
		v238 = v218
		v240 = v220
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v231 = int32(1)
	v232 = v218 + v231
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	if v233 != 0 {
		v217 = v217 + v231
		v218 = v232
		v219 = v228
		v220 = v233
		goto L64
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	goto L55
L70:
	;
	v348 = v77 + int32(1)
	if v268 != 0 {
		v73 = v261
		v75 = v260
		v77 = v348
		goto L18
	} else {
		goto L92
	}
L71:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	v277 = v275
	goto L73
L72:
	;
	v277 = int32(0)
	goto L73
L73:
	;
	if v277 <= v77 {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v279 == int32(0) {
		goto L70
	} else {
		goto L75
	}
L75:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v267 != v282 {
		goto L70
	} else {
		goto L76
	}
L76:
	;
	if v267 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v328 != 0 {
		goto L70
	} else {
		goto L90
	}
L78:
	;
	v328 = int32(0)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v289 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v290 = v75
	v291 = v279
	v292 = v267
	v293 = v289
	goto L85
L82:
	;
	v316 = v279
	v320 = int32(0)
	goto L83
L83:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316))))
	v328 = v320 - v321
	goto L77
L84:
	;
	v316 = v311
	v320 = v313
	goto L83
L85:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	if base.B2i32(v293 != v295)|base.B2i32(v295 == int32(0)) != 0 {
		v311 = v291
		v313 = v293
		goto L84
	} else {
		goto L87
	}
L86:
	;
	v311 = v305
	v313 = int32(0)
	goto L84
L87:
	;
	v301 = v292 - int32(1)
	if v301 == int32(0) {
		v311 = v291
		v313 = v293
		goto L84
	} else {
		goto L88
	}
L88:
	;
	v304 = int32(1)
	v305 = v291 + v304
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)))
	if v306 != 0 {
		v290 = v290 + v304
		v291 = v305
		v292 = v301
		v293 = v306
		goto L85
	} else {
		goto L89
	}
L89:
	;
	goto L86
L90:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v330+v99)))
	v334 = v332 - int32(1)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v329+v334<<(uint(int32(2))%32))))
	if v338 == int32(0) {
		goto L24
	} else {
		goto L91
	}
L91:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v343 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v341+v334))) = uint8(v343)
	goto L70
L92:
	;
	v390 = v262
	v393 = v348
	goto L17
L93:
	;
	v390 = v196
	v393 = v354
	goto L17
L94:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesCSV_3), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v356 + v357<<(uint(int32(4))%32) + v334*int32(100) + int32(24)
	F_errdetail(m, int32(_a_F_CopyReadAttributesCSV_4), v19)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesCSV_1), int32(1990), int32(_a_F_CopyReadAttributesCSV_2))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L7
	} else {
		goto L100
	}
L100:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesCSV_5), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L7
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesCSV_1), int32(1921), int32(_a_F_CopyReadAttributesCSV_2))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L7
	} else {
		goto L102
	}
L102:
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_CreateAuxProcessResourceOwner[0]))
	v6 = F_MemoryContextAllocZero(m, v4, int32(360))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(_a_F_CreateAuxProcessResourceOwner_0)
		v11 = v6 + int32(352)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+356)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v6)+352)) = v11
		*(*int32)(unsafe.Add(mBase, _c_F_CreateAuxProcessResourceOwner[1])) = v6
		*(*int32)(unsafe.Add(mBase, _c_F_CreateAuxProcessResourceOwner[2])) = v6
		F_on_shmem_exit(m, int32(1818), int32(0))
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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
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
	var v367 int32
	_ = v367
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
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
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
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
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
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
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
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
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v668 int32
	_ = v668
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
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v920 int32
	_ = v920
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
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
	v29 = v22 + int32(240)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v29, int32(1), int32(3), int32(184), v33)
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
	v41 = F_systable_beginscan(m, v26, int32(2680), v36, int32(0), v36, v29)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L235
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L231
	}
L6:
	;
	v43 = F_systable_getnext(m, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v49 = v43
	v50 = int32(0)
	goto L11
L9:
	;
	v91 = v36
	goto L10
L10:
	;
	F_systable_endscan(m, v41)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L19
	}
L11:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
	v66 = v64 + v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v67 == v68 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v91 = v72 + int32(1)
	goto L10
L13:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if v50 < v70 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v72 = v70
	goto L16
L15:
	;
	v72 = v50
	goto L16
L16:
	;
	v73 = F_systable_getnext(m, v41)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v73 != 0 {
		v49 = v73
		v50 = v72
		goto L11
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	v100 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if int32(0) < v103 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L1
	} else {
		goto L227
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L1
	} else {
		goto L223
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L219
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L215
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L211
	}
L26:
	;
	v106 = int32(1)
	v112 = v106
	v113 = v103
	v120 = v106
	goto L29
L27:
	;
	goto L28
L28:
	;
	F_relation_close(m, v100, int32(3))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L79
	}
L29:
	;
	v132 = v102 + v113<<(uint(int32(4))%32) + v112*int32(100)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+11)))
	if v133 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v138 = v132 - int32(76)
	v139 = F_SearchSysCacheCopyAttName(m, v136, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	v256 = v113
	goto L33
L33:
	;
	v262 = v120 + int32(1)
	v263 = base.I32_extend16_s(v262)
	if v263 <= v256 {
		v112 = v263
		v113 = v256
		v120 = v262
		goto L29
	} else {
		goto L78
	}
L34:
	;
	if v139 == int32(0) {
		goto L25
	} else {
		goto L35
	}
L35:
	;
	v144 = v132 - int32(80)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+68))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+22)))
	v148 = v146 + v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+68))
	if v145 != v149 {
		goto L21
	} else {
		goto L36
	}
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v144)+76))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+76))
	if v151 != v152 {
		goto L21
	} else {
		goto L37
	}
L37:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v144)+96))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v148)+96))
	if v154 != v155 {
		goto L22
	} else {
		goto L38
	}
L38:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+86)))
	if v157 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+90)))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+90)))
	if v175 != 0 {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+86)))
	if v160 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v162 = int32(*(*int16)(unsafe.Add(mBase, uint32(v144)+74)))
	v163 = F_findNotNullConstraintAttnum(m, v161, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v163 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+22)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v168)+106)))
	if v170 == int32(0) {
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
	if v174 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	if v174 != 0 {
		goto L4
	} else {
		goto L68
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v174 == v175 {
		goto L45
	} else {
		goto L56
	}
L52:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v138
	F_errmsg(m, int32(_a_F_CreateInheritance_0), v22+int32(128))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_2), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
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
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = v138
	F_errmsg(m, int32(_a_F_CreateInheritance_4), v22+int32(160))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+90)))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+90)))
	if v213 == int32(115) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v216 = int32(_a_F_CreateInheritance_5)
	goto L62
L61:
	;
	v216 = int32(_a_F_CreateInheritance_6)
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+148)) = v216
	if v210 == int32(115) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v222 = int32(_a_F_CreateInheritance_5)
	goto L65
L64:
	;
	v222 = int32(_a_F_CreateInheritance_6)
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v222
	F_errdetail(m, int32(_a_F_CreateInheritance_7), v22+int32(144))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_8), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
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
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+89)))
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+89)) = uint8(v234)
	goto L71
L70:
	;
	goto L71
L71:
	;
	v236 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148)+94)))
	v238 = v236 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v148)+94)) = uint16(v238)
	if base.I32_extend16_s(v238) != v238 {
		goto L24
	} else {
		goto L72
	}
L72:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+119)))
	if v243 == int32(112) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v246 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+92)) = uint8(v246)
	goto L75
L74:
	;
	goto L75
L75:
	;
	F_CatalogTupleUpdate(m, v100, v139+int32(4), v139)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_pfree(m, v139)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v256 = v254
	goto L33
L78:
	;
	goto L30
L79:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v290 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v293 = v22 + int32(336)
	F_ScanKeyInit(m, v293, int32(9), int32(3), int32(184), v287)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v300 = int32(1)
	v303 = F_systable_beginscan(m, v290, int32(2665), v300, int32(0), v300, v293)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v308 = F_build_attrmap_by_name(m, v305, v306, int32(1))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v310 = F_systable_getnext(m, v303)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L89
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L207
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L203
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L199
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L195
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L191
	}
L89:
	;
	if v310 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v319 = v310
	goto L93
L91:
	;
	goto L92
L92:
	;
	F_systable_endscan(m, v303)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L178
	}
L93:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v319)+16))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+22)))
	v333 = v331 + v332
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+72)))
	switch v334 - int32(99) {
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
	v646 = F_systable_getnext(m, v303)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L176
	}
L96:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+106)))
	if v337 != 0 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	if v334 == int32(110) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v341 = F_extractNotNullColumn(m, v319)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	v343 = int32(0)
	goto L100
L100:
	;
	v345 = v22 + int32(288)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v345, int32(9), int32(3), int32(184), v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L102
	}
L101:
	;
	v343 = v341
	goto L100
L102:
	;
	v353 = int32(1)
	v356 = F_systable_beginscan(m, v290, int32(2665), v353, int32(0), v353, v345)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L107
	}
L103:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+106)))
	if v584 == int32(1) {
		goto L87
	} else {
		goto L159
	}
L104:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v319)+16))
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539)+22)))
	v541 = v539 + v540
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+73)))
	v543 = v536 + v538
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543)+73)))
	if v542 != v544 {
		goto L88
	} else {
		goto L147
	}
L105:
	;
	v518 = F_extractNotNullColumn(m, v319)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L143
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L140
	}
L107:
	;
	v358 = F_systable_getnext(m, v356)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v358 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v361 = v333 + int32(4)
	v367 = v358
	goto L112
L110:
	;
	goto L111
L111:
	;
	F_systable_endscan(m, v356)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L134
	}
L112:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v367)+16))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+22)))
	v383 = v381 + v382
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+72)))
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+72)))
	if v384 != v385 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L111
L114:
	;
	v458 = F_systable_getnext(m, v356)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L132
	}
L115:
	;
	switch v384 - int32(99) {
	case 0:
		goto L116
	default:
		v419 = v384
		goto L117
	case 11:
		goto L118
	}
L116:
	;
	v426 = v383 + int32(4)
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	if base.B2i32(v429 == int32(0))|base.B2i32(v429 != v432) != 0 {
		v450 = v429
		v451 = v432
		goto L125
	} else {
		goto L126
	}
L117:
	;
	if v419 != int32(99) {
		goto L103
	} else {
		goto L123
	}
L118:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	v391 = F_extractNotNullColumn(m, v367)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	v394 = int32(1)
	v395 = v391 - v394
	v399 = int32(*(*int16)(unsafe.Add(mBase, uint32(v393+v395<<(uint(v394)%32)))))
	if v343 != v399 {
		goto L114
	} else {
		goto L120
	}
L120:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389+v390<<(uint(int32(4))%32)+v343*int32(100))+11)))
	if v407 != 0 {
		goto L106
	} else {
		goto L121
	}
L121:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408+v409<<(uint(int32(4))%32)+v395*int32(100))+111)))
	if v416 != 0 {
		goto L106
	} else {
		goto L122
	}
L122:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+72)))
	v419 = v417
	goto L117
L123:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v367)+16))
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+22)))
	v536 = v423
	v538 = v424
	goto L104
L124:
	;
	if v450-v451 == int32(0) {
		v536 = v381
		v538 = v382
		goto L104
	} else {
		goto L131
	}
L125:
	;
	goto L124
L126:
	;
	v435 = v361
	v436 = v426
	goto L127
L127:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436)+1)))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435)+1)))
	if v440 == int32(0) {
		v450 = v440
		v451 = v439
		goto L125
	} else {
		goto L129
	}
L128:
	;
	v450 = v440
	v451 = v439
	goto L125
L129:
	;
	v443 = int32(1)
	if v440 == v439 {
		v435 = v435 + v443
		v436 = v436 + v443
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	goto L114
L132:
	;
	if v458 != 0 {
		v367 = v458
		goto L112
	} else {
		goto L133
	}
L133:
	;
	goto L113
L134:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+72)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	if v481 == int32(110) {
		goto L105
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v333 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_9), v22+int32(16))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_10), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_CreateInheritance_12), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_13), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	v521 = F_get_attname(m, v287, v518, int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v523 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_14), v22)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_15), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+74)))
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543)+74)))
	if v546 != v547 {
		goto L88
	} else {
		goto L148
	}
L148:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v290)+52))
	v550 = F_decompile_conbin(m, v319, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v552 = F_decompile_conbin(m, v367, v549)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552))))
	if base.B2i32(v556 == int32(0))|base.B2i32(v556 != v559) != 0 {
		v577 = v556
		v578 = v559
		goto L152
	} else {
		goto L153
	}
L151:
	;
	if v577-v578 != 0 {
		goto L88
	} else {
		goto L158
	}
L152:
	;
	goto L151
L153:
	;
	v562 = v550
	v563 = v552
	goto L154
L154:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+1)))
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+1)))
	if v567 == int32(0) {
		v577 = v567
		v578 = v566
		goto L152
	} else {
		goto L156
	}
L155:
	;
	v577 = v567
	v578 = v566
	goto L152
L156:
	;
	v570 = int32(1)
	if v567 == v566 {
		v562 = v562 + v570
		v563 = v563 + v570
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	goto L103
L159:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+76)))
	if v587 != int32(1) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+75)))
	if v596 == int32(1) {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+75)))
	if v590 != int32(1) {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+76)))
	if v593 == int32(0) {
		goto L86
	} else {
		goto L163
	}
L163:
	;
	goto L160
L164:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+75)))
	if v599 == int32(0) {
		goto L85
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v602 = F_heap_copytuple(m, v367)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L168
	}
L167:
	;
	goto L166
L168:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v602)+16))
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+22)))
	v606 = v604 + v605
	v607 = int32(*(*int16)(unsafe.Add(mBase, uint32(v606)+104)))
	v609 = v607 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v606)+104)) = uint16(v609)
	if base.I32_extend16_s(v609) != v609 {
		goto L84
	} else {
		goto L169
	}
L169:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613)+119)))
	if v614 == int32(112) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v617 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v606)+103)) = uint8(v617)
	goto L172
L171:
	;
	goto L172
L172:
	;
	F_CatalogTupleUpdate(m, v290, v602+int32(4), v602)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	F_pfree(m, v602)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	F_systable_endscan(m, v356)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	goto L95
L176:
	;
	if v646 != 0 {
		v319 = v646
		goto L93
	} else {
		goto L177
	}
L177:
	;
	goto L94
L178:
	;
	F_relation_close(m, v290, int32(3))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+119)))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_StoreSingleInheritance(m, v674, v675, v91)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v678 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+296)) = v678
	*(*int32)(unsafe.Add(mBase, uint32(v22)+292)) = v675
	v681 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+288)) = v681
	*(*int32)(unsafe.Add(mBase, uint32(v22)+344)) = v678
	*(*int32)(unsafe.Add(mBase, uint32(v22)+340)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v22)+336)) = v681
	if v673 == int32(112) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v696 = int32(97)
	goto L183
L182:
	;
	v696 = int32(110)
	goto L183
L183:
	;
	F_recordDependencyOn(m, v22+int32(336), v22+int32(288), v696)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v700 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInheritance[0]))
	if v700 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v702 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2611), v674, v702, v675, v702)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	F_SetRelationHasSubclass(m, v675, int32(1))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L189
	}
L188:
	;
	goto L187
L189:
	;
	F_relation_close(m, v26, int32(3))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	m.G0 = v22 + int32(384)
	return
L191:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v723 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_16), v22+int32(80))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_17), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v746 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v383 + v746
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v745 + v746
	F_errmsg(m, int32(_a_F_CreateInheritance_18), v22+int32(32))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_19), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v770 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v383 + v770
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v769 + v770
	F_errmsg(m, int32(_a_F_CreateInheritance_20), v22-int32(-64))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_21), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v794 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v383 + v794
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v793 + v794
	F_errmsg(m, int32(_a_F_CreateInheritance_22), v22+int32(48))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_23), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	F_errmsg(m, int32(_a_F_CreateInheritance_24), int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_25), int32(_a_F_CreateInheritance_11))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L211:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v138
	F_errmsg(m, int32(_a_F_CreateInheritance_26), v22+int32(96))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_27), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L215:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	F_errmsg(m, int32(_a_F_CreateInheritance_24), int32(0))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_28), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L219:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v867 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_14), v22+int32(176))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_29), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L223:
	;
	F_errcode(m, int32(17432708))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+196)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v889 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_30), v22+int32(192))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_31), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+212)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = v911 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_32), v22+int32(208))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_33), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L1
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
	F_errcode(m, int32(117571716))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+224)) = v933 + int32(4)
	F_errmsg(m, int32(_a_F_CreateInheritance_34), v22+int32(224))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_35), int32(_a_F_CreateInheritance_36))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v138
	F_errmsg(m, int32(_a_F_CreateInheritance_37), v22+int32(112))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_CreateInheritance_1), int32(_a_F_CreateInheritance_38), int32(_a_F_CreateInheritance_3))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
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
	var v8 float64
	_ = v8
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 float64
	_ = v79
	var v80 int32
	_ = v80
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 float64
	_ = v111
	var v112 int32
	_ = v112
	var v115 float64
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 float64
	_ = v122
	var v128 float64
	_ = v128
	var v131 float64
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 float64
	_ = v141
	var v143 float64
	_ = v143
	var v147 float64
	_ = v147
	var v148 int32
	_ = v148
	var v149 float64
	_ = v149
	var v152 float64
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 float64
	_ = v159
	var v160 int32
	_ = v160
	var v163 float64
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 float64
	_ = v170
	var v176 float64
	_ = v176
	var v179 float64
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 float64
	_ = v189
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v193 int32
	_ = v193
	var v198 float64
	_ = v198
	var v209 float64
	_ = v209
	var v211 float64
	_ = v211
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 float64
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 float64
	_ = v268
	var v269 float64
	_ = v269
	var v270 float64
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 float64
	_ = v273
	var v274 float64
	_ = v274
	var v276 int32
	_ = v276
	var v290 int32
	_ = v290
	var v299 float64
	_ = v299
	var v303 float64
	_ = v303
	var v305 int32
	_ = v305
	var v311 float64
	_ = v311
	var v314 float64
	_ = v314
	var v315 float64
	_ = v315
	var v323 int32
	_ = v323
	var v328 float64
	_ = v328
	var v330 float64
	_ = v330
	var v331 float64
	_ = v331
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 float64
	_ = v342
	var v349 float64
	_ = v349
	var v352 float64
	_ = v352
	var v362 float64
	_ = v362
	var v367 int32
	_ = v367
	var v368 float64
	_ = v368
	var v371 float64
	_ = v371
	var v372 float64
	_ = v372
	var v373 int32
	_ = v373
	var v374 float64
	_ = v374
	var v376 int32
	_ = v376
	var v390 int32
	_ = v390
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
	var v419 float64
	_ = v419
	var v421 float64
	_ = v421
	var v422 float64
	_ = v422
	var v424 float64
	_ = v424
	var v428 float64
	_ = v428
	var v432 float64
	_ = v432
	var v442 float64
	_ = v442
	var v462 float64
	_ = v462
	var v472 float64
	_ = v472
	v8 = float64(0)
	v18 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v18)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	v22 = v20 ^ v18
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v22)
	v26 = l4 - v18
	v39 = v26
	v40 = int32(-1)
	goto L1
L1:
	;
	v48 = base.I32_div_s(v39+v40+int32(1), int32(2))
	v52 = F_range_cmp_bounds(m, l0, l3+v48<<(uint(int32(3))%32), l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	if v58 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	return float64(0)
L4:
	;
	v57 = base.B2i32(v52 < int32(0))
	if v52 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v58 = v48
	goto L7
L6:
	;
	v58 = v40
	goto L7
L7:
	;
	if v52 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v61 = v39
	goto L10
L9:
	;
	v61 = v48 - int32(1)
	goto L10
L10:
	;
	if v58 < v61 {
		v39 = v61
		v40 = v58
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
	v68 = l0 + int32(268)
	v71 = l4 - int32(2)
	if base.Ui32(v58) < base.Ui32(v71) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v73 = v58
	goto L17
L16:
	;
	v73 = v71
	goto L17
L17:
	;
	v76 = l3 + v73<<(uint(int32(3))%32)
	v79 = F_get_position(m, l0, l2, v76, v76+int32(8))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v89 = v79
	v90 = v8
	v92 = v8
	v95 = v73
	goto L19
L19:
	;
	v100 = l3 + v95<<(uint(int32(3))%32)
	v101 = F_range_cmp_bounds(m, l0, v100, l1)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L3
	} else {
		goto L22
	}
L20:
	;
	return v472
L21:
	;
	v193 = int32(1)
	v198 = float64(0)
	if base.F64_lt(v191, v198) != 0 {
		v462 = v198
		goto L70
	} else {
		goto L71
	}
L22:
	;
	if v101 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v105 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)))
	if v153 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L26:
	;
	v147 = F_get_position(m, l0, l1, v100, v100+int32(8))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L46
	}
L27:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v110 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v133 != int32(1) {
		v143 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L26
	} else {
		goto L42
	}
L30:
	;
	v111 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L32
L31:
	;
	v111 = float64(1)
	goto L32
L32:
	;
	if v110 != 0 {
		v143 = v111
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v112 == int32(0) {
		v143 = v111
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v115 = float64(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v120 = F_FunctionCall2Coll(m, v68, v117, v118, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v122 = *(*float64)(unsafe.Add(mBase, uint32(v120)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v122)&int64(9223372036854775807)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v128 = v115
	goto L38
L37:
	;
	v128 = v122
	goto L38
L38:
	;
	if base.F64_lt(v122, float64(0)) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v131 = v115
	goto L41
L40:
	;
	v131 = v128
	goto L41
L41:
	;
	v143 = v131
	goto L26
L42:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v138 == v139 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v141 = float64(0)
	goto L45
L44:
	;
	v141 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L45
L45:
	;
	v143 = v141
	goto L26
L46:
	;
	v149 = base.F64_sub(v89, v147)
	if base.F64_lt(v149, float64(0)) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v152 = float64(0)
	goto L49
L48:
	;
	v152 = v149
	goto L49
L49:
	;
	v191 = v143
	v192 = v152
	goto L21
L50:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v158 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v181 != int32(1) {
		v191 = math.Float64frombits(uint64(0x7ff0000000000000))
		v192 = v89
		goto L21
	} else {
		goto L65
	}
L53:
	;
	v159 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L55
L54:
	;
	v159 = float64(1)
	goto L55
L55:
	;
	if v158 != 0 {
		v191 = v159
		v192 = v89
		goto L21
	} else {
		goto L56
	}
L56:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v160 == int32(0) {
		v191 = v159
		v192 = v89
		goto L21
	} else {
		goto L57
	}
L57:
	;
	v163 = float64(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v168 = F_FunctionCall2Coll(m, v68, v165, v166, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v170 = *(*float64)(unsafe.Add(mBase, uint32(v168)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v170)&int64(9223372036854775807)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v176 = v163
	goto L61
L60:
	;
	v176 = v170
	goto L61
L61:
	;
	if base.F64_lt(v170, float64(0)) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v179 = v163
	goto L64
L63:
	;
	v179 = v176
	goto L64
L64:
	;
	v191 = v179
	v192 = v89
	goto L21
L65:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+6)))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v186 == v187 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v189 = float64(0)
	goto L68
L67:
	;
	v189 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L68
L68:
	;
	v191 = v189
	v192 = v89
	goto L21
L69:
	;
	v472 = base.F64_add(v90, base.F64_div(base.F64_mul(v192, v462), base.F64_convert_i32_u(v26)))
	if int32(0) <= v101 {
		goto L141
	} else {
		goto L142
	}
L70:
	;
	goto L69
L71:
	;
	v209 = float64(1)
	v211 = base.F64_abs(v191)
	goto L72
L72:
	;
	goto L74
L74:
	;
	if base.F64_eq(v211, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v462 = v209
		goto L70
	} else {
		goto L75
	}
L75:
	;
	v217 = l6 - int32(1)
	if v217 < int32(0) {
		v462 = v209
		goto L70
	} else {
		goto L76
	}
L76:
	;
	v221 = v217
	v225 = int32(-1)
	goto L77
L77:
	;
	v241 = int32(2)
	v242 = base.I32_div_s(v221+v225+int32(1), v241)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l5+v242<<(uint(v241)%32))))
	v247 = *(*float64)(unsafe.Add(mBase, uint32(v246)))
	if base.F64_gt(v92, v247) != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	if v217 <= v257 {
		v462 = v209
		goto L70
	} else {
		goto L90
	}
L79:
	;
	if v257 < v255 {
		v221 = v255
		v225 = v257
		goto L77
	} else {
		goto L89
	}
L80:
	;
	v255 = v221
	v257 = v242
	goto L79
L81:
	;
	goto L82
L82:
	;
	v252 = v193 & base.F64_ge(v92, v247)
	if v252 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v253 = v221
	goto L85
L84:
	;
	v253 = v242 - int32(1)
	goto L85
L85:
	;
	if v252 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v254 = v242
	goto L88
L87:
	;
	v254 = v225
	goto L88
L88:
	;
	v255 = v253
	v257 = v254
	goto L79
L89:
	;
	goto L78
L90:
	;
	if v257 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v314 = base.F64_convert_i32_u(v217)
	v315 = base.F64_div(base.F64_add(v311, base.F64_convert_i32_u(v305)), v314)
	if base.F64_eq(v92, v191) != 0 {
		v462 = v315
		goto L70
	} else {
		goto L106
	}
L92:
	;
	v305 = int32(0)
	v311 = float64(0)
	goto L91
L93:
	;
	goto L94
L94:
	;
	v266 = l5 + v257<<(uint(int32(2))%32)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v268 = *(*float64)(unsafe.Add(mBase, uint32(v267)))
	v269 = base.F64_abs(v268)
	v270 = math.Float64frombits(uint64(0x7ff0000000000000))
	v271 = base.F64_eq(v269, v270)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v273 = *(*float64)(unsafe.Add(mBase, uint32(v272)))
	v274 = base.F64_abs(v273)
	v276 = base.F64_eq(v274, v270)
	if v271|v276 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if base.F64_eq(base.F64_abs(v92), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v305 = v257
		v311 = float64(0.5)
		goto L91
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v290 = int32(0)
	if v271|base.B2i32(v276 == v290) == v290 {
		v305 = v257
		v311 = float64(1)
		goto L91
	} else {
		goto L99
	}
L98:
	;
	v305 = v257
	v311 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v268, v92), base.F64_sub(v268, v273)))
	goto L91
L99:
	;
	if base.F64_eq(v269, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v299 = float64(0)
	goto L102
L101:
	;
	v299 = float64(0.5)
	goto L102
L102:
	;
	if base.F64_eq(v274, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v303 = v299
	goto L105
L104:
	;
	v303 = float64(0.5)
	goto L105
L105:
	;
	v305 = v257
	v311 = v303
	goto L91
L106:
	;
	if v217 <= v305 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v428 = float64(0)
	v432 = base.F64_div(base.F64_add(v424, base.F64_convert_i32_u(v416)), v314)
	if base.F64_gt(v421, v428)|base.F64_gt(v432, v428) != 0 {
		goto L134
	} else {
		goto L135
	}
L108:
	;
	v416 = v305
	v419 = v92
	v421 = v315
	v422 = v198
	v424 = v198
	goto L107
L109:
	;
	goto L110
L110:
	;
	v323 = v305
	v328 = v315
	v330 = v198
	v331 = v92
	goto L112
L111:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l5+v323<<(uint(int32(2))%32))))
	v368 = *(*float64)(unsafe.Add(mBase, uint32(v367)))
	if base.F64_eq(v342, v368) != 0 {
		goto L119
	} else {
		goto L120
	}
L112:
	;
	v337 = v323 + int32(1)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l5+v337<<(uint(int32(2))%32))))
	v342 = *(*float64)(unsafe.Add(mBase, uint32(v341)))
	if base.F64_lt(v342, v191)|v193&base.F64_ge(v191, v342) == int32(0) {
		goto L111
	} else {
		goto L114
	}
L113:
	;
	v416 = v217
	v419 = v342
	v421 = v352
	v422 = v362
	v424 = v198
	goto L107
L114:
	;
	v349 = float64(0)
	v352 = base.F64_div(base.F64_convert_i32_u(v323), v314)
	if base.F64_gt(v328, v349)|base.F64_gt(v352, v349) != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v362 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v328, v352), float64(0.5)), base.F64_sub(v342, v331)), v330)
	goto L117
L116:
	;
	v362 = v330
	goto L117
L117:
	;
	if v217 != v337 {
		v323 = v337
		v328 = v352
		v330 = v362
		v331 = v342
		goto L112
	} else {
		goto L118
	}
L118:
	;
	goto L113
L119:
	;
	v409 = float64(0)
	goto L121
L120:
	;
	v371 = base.F64_abs(v342)
	v372 = math.Float64frombits(uint64(0x7ff0000000000000))
	v373 = base.F64_eq(v371, v372)
	v374 = base.F64_abs(v368)
	v376 = base.F64_eq(v374, v372)
	if v373|v376 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v416 = v323
	v419 = v331
	v421 = v328
	v422 = v330
	v424 = v409
	goto L107
L122:
	;
	v409 = v404
	goto L121
L123:
	;
	if base.F64_eq(base.F64_abs(v191), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v404 = float64(0.5)
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v390 = int32(0)
	if v373|base.B2i32(v376 == v390) == v390 {
		v404 = float64(1)
		goto L122
	} else {
		goto L127
	}
L126:
	;
	v404 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v342, v191), base.F64_sub(v342, v368)))
	goto L122
L127:
	;
	if base.F64_eq(v371, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v399 = float64(0)
	goto L130
L129:
	;
	v399 = float64(0.5)
	goto L130
L130:
	;
	if base.F64_eq(v374, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v403 = v399
	goto L133
L132:
	;
	v403 = float64(0.5)
	goto L133
L133:
	;
	v404 = v403
	goto L122
L134:
	;
	v442 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v421, v432), float64(0.5)), base.F64_sub(v191, v419)), v422)
	goto L136
L135:
	;
	v442 = v422
	goto L136
L136:
	;
	if base.F64_eq(v211, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	if base.F64_eq(base.F64_abs(v442), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v462 = float64(0.5)
		goto L70
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v462 = base.F64_div(v442, base.F64_sub(v191, v92))
	goto L70
L140:
	;
	goto L139
L141:
	;
	if int32(0) < v95 {
		v89 = float64(1)
		v90 = v472
		v92 = v191
		v95 = v95 - int32(1)
		goto L19
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	goto L20
L144:
	;
	goto L143
}
func F_calc_length_hist_frac(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 int32) float64 {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
	var v21 float64
	_ = v21
	var v23 float64
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v82 float64
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v88 int32
	_ = v88
	var v102 int32
	_ = v102
	var v111 float64
	_ = v111
	var v115 float64
	_ = v115
	var v117 int32
	_ = v117
	var v123 float64
	_ = v123
	var v126 float64
	_ = v126
	var v127 float64
	_ = v127
	var v135 int32
	_ = v135
	var v140 float64
	_ = v140
	var v142 float64
	_ = v142
	var v143 float64
	_ = v143
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 float64
	_ = v154
	var v161 float64
	_ = v161
	var v164 float64
	_ = v164
	var v174 float64
	_ = v174
	var v179 int32
	_ = v179
	var v180 float64
	_ = v180
	var v183 float64
	_ = v183
	var v184 float64
	_ = v184
	var v185 int32
	_ = v185
	var v186 float64
	_ = v186
	var v188 int32
	_ = v188
	var v202 int32
	_ = v202
	var v211 float64
	_ = v211
	var v215 float64
	_ = v215
	var v216 float64
	_ = v216
	var v221 float64
	_ = v221
	var v228 int32
	_ = v228
	var v231 float64
	_ = v231
	var v233 float64
	_ = v233
	var v234 float64
	_ = v234
	var v236 float64
	_ = v236
	var v240 float64
	_ = v240
	var v244 float64
	_ = v244
	var v254 float64
	_ = v254
	var v274 float64
	_ = v274
	v10 = float64(0)
	if base.F64_lt(l3, v10) != 0 {
		v274 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v274
L2:
	;
	v21 = float64(1)
	v23 = base.F64_abs(l3)
	if base.F64_eq(v23, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = l4
	goto L5
L4:
	;
	v26 = int32(0)
	goto L5
L5:
	;
	if v26 != 0 {
		v274 = v21
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v29 = l1 - int32(1)
	if v29 < int32(0) {
		v274 = v21
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v33 = v29
	v37 = int32(-1)
	goto L8
L8:
	;
	v53 = int32(2)
	v54 = base.I32_div_s(v33+v37+int32(1), v53)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0+v54<<(uint(v53)%32))))
	v59 = *(*float64)(unsafe.Add(mBase, uint32(v58)))
	if base.F64_gt(l2, v59) != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v29 <= v69 {
		v274 = v21
		goto L1
	} else {
		goto L21
	}
L10:
	;
	if v69 < v67 {
		v33 = v67
		v37 = v69
		goto L8
	} else {
		goto L20
	}
L11:
	;
	v67 = v33
	v69 = v54
	goto L10
L12:
	;
	goto L13
L13:
	;
	v64 = l4 & base.F64_ge(l2, v59)
	if v64 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v65 = v33
	goto L16
L15:
	;
	v65 = v54 - int32(1)
	goto L16
L16:
	;
	if v64 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v66 = v54
	goto L19
L18:
	;
	v66 = v37
	goto L19
L19:
	;
	v67 = v65
	v69 = v66
	goto L10
L20:
	;
	goto L9
L21:
	;
	if v69 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v126 = base.F64_convert_i32_u(v29)
	v127 = base.F64_div(base.F64_add(v123, base.F64_convert_i32_u(v117)), v126)
	if base.F64_eq(l2, l3) != 0 {
		v274 = v127
		goto L1
	} else {
		goto L37
	}
L23:
	;
	v117 = int32(0)
	v123 = float64(0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v78 = l0 + v69<<(uint(int32(2))%32)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v80 = *(*float64)(unsafe.Add(mBase, uint32(v79)))
	v81 = base.F64_abs(v80)
	v82 = math.Float64frombits(uint64(0x7ff0000000000000))
	v83 = base.F64_eq(v81, v82)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v85 = *(*float64)(unsafe.Add(mBase, uint32(v84)))
	v86 = base.F64_abs(v85)
	v88 = base.F64_eq(v86, v82)
	if v83|v88 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if base.F64_eq(base.F64_abs(l2), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v117 = v69
		v123 = float64(0.5)
		goto L22
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v102 = int32(0)
	if v83|base.B2i32(v88 == v102) == v102 {
		v117 = v69
		v123 = float64(1)
		goto L22
	} else {
		goto L30
	}
L29:
	;
	v117 = v69
	v123 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v80, l2), base.F64_sub(v80, v85)))
	goto L22
L30:
	;
	if base.F64_eq(v81, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v111 = float64(0)
	goto L33
L32:
	;
	v111 = float64(0.5)
	goto L33
L33:
	;
	if base.F64_eq(v86, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v115 = v111
	goto L36
L35:
	;
	v115 = float64(0.5)
	goto L36
L36:
	;
	v117 = v69
	v123 = v115
	goto L22
L37:
	;
	if v29 <= v117 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v240 = float64(0)
	v244 = base.F64_div(base.F64_add(v236, base.F64_convert_i32_u(v228)), v126)
	if base.F64_gt(v233, v240)|base.F64_gt(v244, v240) != 0 {
		goto L65
	} else {
		goto L66
	}
L39:
	;
	v228 = v117
	v231 = l2
	v233 = v127
	v234 = v10
	v236 = v10
	goto L38
L40:
	;
	goto L41
L41:
	;
	v135 = v117
	v140 = v127
	v142 = v10
	v143 = l2
	goto L43
L42:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0+v135<<(uint(int32(2))%32))))
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v179)))
	if base.F64_eq(v154, v180) != 0 {
		goto L50
	} else {
		goto L51
	}
L43:
	;
	v149 = v135 + int32(1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0+v149<<(uint(int32(2))%32))))
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v153)))
	if base.F64_lt(v154, l3)|l4&base.F64_ge(l3, v154) == int32(0) {
		goto L42
	} else {
		goto L45
	}
L44:
	;
	v228 = v29
	v231 = v154
	v233 = v164
	v234 = v174
	v236 = v10
	goto L38
L45:
	;
	v161 = float64(0)
	v164 = base.F64_div(base.F64_convert_i32_u(v135), v126)
	if base.F64_gt(v140, v161)|base.F64_gt(v164, v161) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v174 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v140, v164), float64(0.5)), base.F64_sub(v154, v143)), v142)
	goto L48
L47:
	;
	v174 = v142
	goto L48
L48:
	;
	if v29 != v149 {
		v135 = v149
		v140 = v164
		v142 = v174
		v143 = v154
		goto L43
	} else {
		goto L49
	}
L49:
	;
	goto L44
L50:
	;
	v221 = float64(0)
	goto L52
L51:
	;
	v183 = base.F64_abs(v154)
	v184 = math.Float64frombits(uint64(0x7ff0000000000000))
	v185 = base.F64_eq(v183, v184)
	v186 = base.F64_abs(v180)
	v188 = base.F64_eq(v186, v184)
	if v185|v188 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v228 = v135
	v231 = v143
	v233 = v140
	v234 = v142
	v236 = v221
	goto L38
L53:
	;
	v221 = v216
	goto L52
L54:
	;
	if base.F64_eq(base.F64_abs(l3), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v216 = float64(0.5)
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v202 = int32(0)
	if v185|base.B2i32(v188 == v202) == v202 {
		v216 = float64(1)
		goto L53
	} else {
		goto L58
	}
L57:
	;
	v216 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v154, l3), base.F64_sub(v154, v180)))
	goto L53
L58:
	;
	if base.F64_eq(v183, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v211 = float64(0)
	goto L61
L60:
	;
	v211 = float64(0.5)
	goto L61
L61:
	;
	if base.F64_eq(v186, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v215 = v211
	goto L64
L63:
	;
	v215 = float64(0.5)
	goto L64
L64:
	;
	v216 = v215
	goto L53
L65:
	;
	v254 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v233, v244), float64(0.5)), base.F64_sub(l3, v231)), v234)
	goto L67
L66:
	;
	v254 = v234
	goto L67
L67:
	;
	if base.F64_eq(v23, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if base.F64_eq(base.F64_abs(v254), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v274 = float64(0.5)
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v274 = base.F64_div(v254, base.F64_sub(l3, l2))
	goto L1
L71:
	;
	goto L70
}
func F_calc_non_nestloop_required_outer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		v6 = v4
	} else {
		v6 = int32(0)
	}
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v7 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		v10 = v8
	} else {
		v10 = int32(0)
	}
	v11 = F_bms_union(m, v6, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
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
	var v95 int32
	_ = v95
	var v111 float32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v293 int32
	_ = v293
	var v308 float32
	_ = v308
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v374 float32
	_ = v374
	var v375 float32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 float32
	_ = v386
	var v387 int32
	_ = v387
	var v388 float32
	_ = v388
	var v390 int32
	_ = v390
	var v394 float32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v438 float64
	_ = v438
	var v443 float32
	_ = v443
	var v445 int32
	_ = v445
	var v478 float32
	_ = v478
	var v483 int32
	_ = v483
	var v519 float32
	_ = v519
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v563 float32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v737 int32
	_ = v737
	var v761 float32
	_ = v761
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v802 int32
	_ = v802
	var v820 float32
	_ = v820
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v856 int32
	_ = v856
	var v872 float32
	_ = v872
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v893 int32
	_ = v893
	var v916 float32
	_ = v916
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v930 float32
	_ = v930
	var v931 int32
	_ = v931
	var v936 float32
	_ = v936
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v953 float64
	_ = v953
	var v961 float32
	_ = v961
	var v963 float32
	_ = v963
	var v966 float64
	_ = v966
	var v977 float32
	_ = v977
	var v980 int32
	_ = v980
	var v1010 float32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1045 float32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1080 float32
	_ = v1080
	var v1085 int32
	_ = v1085
	var v1118 float32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1138 int32
	_ = v1138
	var v1157 float32
	_ = v1157
	var v1164 float32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1183 int32
	_ = v1183
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1241 float64
	_ = v1241
	var v1274 float32
	_ = v1274
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1382 float32
	_ = v1382
	var v1390 int32
	_ = v1390
	var v1396 float32
	_ = v1396
	var v1401 int32
	_ = v1401
	var v1408 float64
	_ = v1408
	var v1414 float32
	_ = v1414
	var v1431 int32
	_ = v1431
	var v1450 float32
	_ = v1450
	v5 = int32(0)
	v29 = float32(0)
	v33 = m.G0
	v35 = v33 - int32(16)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v37 == v5 {
		v1431 = v35
		v1450 = v29
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v1431 + int32(16)
	return v1450
L2:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v40 == int32(0) {
		v1431 = v35
		v1450 = v29
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
	if base.F32_lt(v1157, float32(0)) != 0 {
		goto L147
	} else {
		goto L148
	}
L5:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v528 = F_palloc0(m, v525<<(uint(int32(2))%32))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
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
	v519 = float32(0)
	goto L16
L15:
	;
	v95 = v5
	v111 = v29
	goto L17
L16:
	;
	F_pfree(m, v73)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L9
	} else {
		goto L72
	}
L17:
	;
	v115 = int32(2)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v73+v95<<(uint(v115)%32))))
	v119 = int32(8)
	v120 = v65 + v119
	v121 = int32(0)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v121
	v131 = l1 + v119
	v134 = v131 + v127<<(uint(v115)%32)
	if v127 <= v121 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v519 = base.F32_div(v478, base.F32_convert_i32_u(v75))
	goto L16
L19:
	;
	v483 = v95 + int32(1)
	if v483 != v75 {
		v95 = v483
		v111 = v478
		goto L17
	} else {
		goto L71
	}
L20:
	;
	if v274 == int32(0) {
		v478 = v111
		goto L19
	} else {
		goto L50
	}
L21:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
	if v202 != int32(1) {
		goto L37
	} else {
		goto L38
	}
L22:
	;
	v196 = v134
	v197 = v131
	v198 = v134
	goto L21
L23:
	;
	goto L24
L24:
	;
	v144 = v131
	v145 = v134
	goto L25
L25:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v150 = int32(12)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v160 = int32(2)
	v167 = base.I32_div_s((v145-v144)>>(uint(v160)%32), v160)
	v170 = v144 + v167<<(uint(v160)%32)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v179 = int32(0)
	v180 = F_tsCompareString(m, l2+int32(8)+v149*v150+int32(base.Ui32(v153)>>(uint(v150)%32)), v153&int32(4095), v131+v159<<(uint(v160)%32)+int32(base.Ui32(v171)>>(uint(v150)%32)), int32(base.Ui32(v171)>>(uint(int32(1))%32))&int32(2047), v179)
	mBase = m.M
	if v180 == v179 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v196 = v170
	v197 = v189
	v198 = v190
	goto L21
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = int32(1)
	v196 = v170
	v197 = v144
	v198 = v170
	goto L21
L28:
	;
	goto L29
L29:
	;
	v188 = base.B2i32(int32(0) < v180)
	if int32(0) < v180 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v189 = v170 + int32(4)
	goto L32
L31:
	;
	v189 = v144
	goto L32
L32:
	;
	if int32(0) < v180 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v190 = v145
	goto L35
L34:
	;
	v190 = v170
	goto L35
L35:
	;
	if base.Ui32(v189) < base.Ui32(v190) {
		v144 = v189
		v145 = v190
		goto L25
	} else {
		goto L36
	}
L36:
	;
	goto L26
L37:
	;
	v270 = int32(0)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v270 < v271 {
		goto L47
	} else {
		goto L48
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = int32(0)
	if base.Ui32(v197) < base.Ui32(v198) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v208 = v196
	goto L41
L40:
	;
	v208 = v198
	goto L41
L41:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v131+v209<<(uint(int32(2))%32)) <= base.Ui32(v208) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v220 = v209
	v221 = v208
	goto L43
L43:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v227 = int32(12)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v243 = int32(1)
	v248 = F_tsCompareString(m, l2+int32(8)+v226*v227+int32(base.Ui32(v230)>>(uint(v227)%32)), v230&int32(4095), v131+v220<<(uint(int32(2))%32)+int32(base.Ui32(v239)>>(uint(v227)%32)), int32(base.Ui32(v239)>>(uint(v243)%32))&int32(2047), v243)
	mBase = m.M
	if v248 != 0 {
		goto L37
	} else {
		goto L45
	}
L44:
	;
	goto L37
L45:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v249 + int32(1)
	v254 = v221 + int32(4)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v254) < base.Ui32(v131+v255<<(uint(int32(2))%32)) {
		v220 = v255
		v221 = v254
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v274 = v198
	goto L49
L48:
	;
	v274 = v270
	goto L49
L49:
	;
	goto L20
L50:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if v277 <= int32(0) {
		v478 = v111
		goto L19
	} else {
		goto L51
	}
L51:
	;
	v293 = v274
	v308 = v111
	goto L52
L52:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	if v314&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v478 = v443
	goto L19
L54:
	;
	v317 = int32(1)
	v327 = (int32(base.Ui32(v314)>>(uint(v317)%32))&int32(2047) + int32(base.Ui32(v314)>>(uint(int32(12))%32)) + v317) & int32(_a_F_calc_rank_0)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v330 = v328 << (uint(int32(2)) % 32)
	v337 = v327 + (l1 + int32(8) + v330)
	v339 = l1 + v330 + v327 + int32(10)
	goto L56
L55:
	;
	v337 = v65 + int32(12)
	v339 = v65 + int32(14)
	goto L56
L56:
	;
	v340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337))))
	if v340 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v443 = base.F32_demote_f64(base.F64_add(base.F64_div(v438, float64(1.64493406685)), base.F64_promote_f32(v308)))
	v445 = v293 + int32(4)
	if (v445-v274)>>(uint(int32(2))%32) < v277 {
		v293 = v445
		v308 = v443
		goto L52
	} else {
		goto L70
	}
L58:
	;
	v343 = int32(0)
	v349 = v343
	v355 = v343
	v374 = float32(-1)
	v375 = float32(0)
	goto L61
L59:
	;
	goto L60
L60:
	;
	v438 = float64(0)
	goto L57
L61:
	;
	v380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v339+v349<<(uint(int32(1))%32)))))
	v381 = int32(12)
	v386 = *(*float32)(unsafe.Add(mBase, uint32(l0+int32(base.Ui32(v380)>>(uint(v381)%32))&v381)))
	v387 = base.F32_lt(v374, v386)
	if v387 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v399 = v395 + int32(1)
	v438 = base.F64_promote_f32(base.F32_sub(base.F32_add(v394, v388), base.F32_div(v388, base.F32_convert_i32_s(v399*v399))))
	goto L57
L63:
	;
	v388 = v386
	goto L65
L64:
	;
	v388 = v374
	goto L65
L65:
	;
	v390 = v349 + int32(1)
	v394 = base.F32_add(v375, base.F32_div(v386, base.F32_convert_i32_s(v390*v390)))
	if v387 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v395 = v349
	goto L68
L67:
	;
	v395 = v355
	goto L68
L68:
	;
	if v390 != v340 {
		v349 = v390
		v355 = v395
		v374 = v388
		v375 = v394
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
	v1130 = l1
	v1132 = l3
	v1138 = v35
	v1157 = v519
	goto L4
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = int32(1073676289)
	v535 = l0
	v536 = l1
	v537 = l2
	v538 = l3
	v541 = v528
	v542 = v5
	v544 = v35
	v549 = v52
	v553 = v56
	v554 = l1 + int32(8)
	v563 = float32(-1)
	goto L74
L74:
	;
	v567 = int32(2)
	v568 = v542 << (uint(v567) % 32)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v549+v568)))
	v571 = int32(8)
	v572 = v544 + v571
	v573 = int32(0)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v572))) = v573
	v583 = v536 + v571
	v586 = v583 + v579<<(uint(v567)%32)
	if v579 <= v573 {
		goto L79
	} else {
		goto L80
	}
L75:
	;
	F_pfree(m, v541)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L9
	} else {
		goto L145
	}
L76:
	;
	v1123 = v542 + int32(1)
	if v1123 != v553 {
		v542 = v1123
		v563 = v1118
		goto L74
	} else {
		goto L144
	}
L77:
	;
	if v726 == int32(0) {
		v1118 = v563
		goto L76
	} else {
		goto L107
	}
L78:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570)+2)))
	if v654 != int32(1) {
		goto L94
	} else {
		goto L95
	}
L79:
	;
	v648 = v586
	v649 = v583
	v650 = v586
	goto L78
L80:
	;
	goto L81
L81:
	;
	v596 = v583
	v597 = v586
	goto L82
L82:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v537)+4))
	v602 = int32(12)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v570)+8))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	v612 = int32(2)
	v619 = base.I32_div_s((v597-v596)>>(uint(v612)%32), v612)
	v622 = v596 + v619<<(uint(v612)%32)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	v631 = int32(0)
	v632 = F_tsCompareString(m, v537+int32(8)+v601*v602+int32(base.Ui32(v605)>>(uint(v602)%32)), v605&int32(4095), v583+v611<<(uint(v612)%32)+int32(base.Ui32(v623)>>(uint(v602)%32)), int32(base.Ui32(v623)>>(uint(int32(1))%32))&int32(2047), v631)
	mBase = m.M
	if v632 == v631 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v648 = v622
	v649 = v641
	v650 = v642
	goto L78
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572))) = int32(1)
	v648 = v622
	v649 = v596
	v650 = v622
	goto L78
L85:
	;
	goto L86
L86:
	;
	v640 = base.B2i32(int32(0) < v632)
	if int32(0) < v632 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v641 = v622 + int32(4)
	goto L89
L88:
	;
	v641 = v596
	goto L89
L89:
	;
	if int32(0) < v632 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v642 = v597
	goto L92
L91:
	;
	v642 = v622
	goto L92
L92:
	;
	if base.Ui32(v641) < base.Ui32(v642) {
		v596 = v641
		v597 = v642
		goto L82
	} else {
		goto L93
	}
L93:
	;
	goto L83
L94:
	;
	v722 = int32(0)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	if v722 < v723 {
		goto L104
	} else {
		goto L105
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572))) = int32(0)
	if base.Ui32(v649) < base.Ui32(v650) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v660 = v648
	goto L98
L97:
	;
	v660 = v650
	goto L98
L98:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	if base.Ui32(v583+v661<<(uint(int32(2))%32)) <= base.Ui32(v660) {
		goto L94
	} else {
		goto L99
	}
L99:
	;
	v672 = v661
	v673 = v660
	goto L100
L100:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v537)+4))
	v679 = int32(12)
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v570)+8))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v673)))
	v695 = int32(1)
	v700 = F_tsCompareString(m, v537+int32(8)+v678*v679+int32(base.Ui32(v682)>>(uint(v679)%32)), v682&int32(4095), v583+v672<<(uint(int32(2))%32)+int32(base.Ui32(v691)>>(uint(v679)%32)), int32(base.Ui32(v691)>>(uint(v695)%32))&int32(2047), v695)
	mBase = m.M
	if v700 != 0 {
		goto L94
	} else {
		goto L102
	}
L101:
	;
	goto L94
L102:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	*(*int32)(unsafe.Add(mBase, uint32(v572))) = v701 + int32(1)
	v706 = v673 + int32(4)
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	if base.Ui32(v706) < base.Ui32(v583+v707<<(uint(int32(2))%32)) {
		v672 = v707
		v673 = v706
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v726 = v650
	goto L106
L105:
	;
	v726 = v722
	goto L106
L106:
	;
	goto L77
L107:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v544)+8))
	if v729 <= int32(0) {
		v1118 = v563
		goto L76
	} else {
		goto L108
	}
L108:
	;
	v737 = v726
	v761 = v563
	goto L109
L109:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v737)))
	if v767&int32(1) != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v1118 = v1080
	goto L76
L111:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	v774 = int32(1)
	v786 = v554 + v770<<(uint(int32(2))%32) + (int32(base.Ui32(v767)>>(uint(v774)%32))&int32(2047)+int32(base.Ui32(v767)>>(uint(int32(12))%32))+v774)&int32(_a_F_calc_rank_0)
	goto L113
L112:
	;
	v786 = v544 + int32(12)
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v568+v541))) = v786
	if v542 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v790 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v786))))
	v802 = int32(0)
	v820 = v761
	goto L117
L115:
	;
	v1080 = v761
	goto L116
L116:
	;
	v1085 = v737 + int32(4)
	if (v1085-v726)>>(uint(int32(2))%32) < v729 {
		v737 = v1085
		v761 = v1080
		goto L109
	} else {
		goto L143
	}
L117:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v541+v802<<(uint(int32(2))%32))))
	v828 = int32(0)
	if base.B2i32(v827 == v828)|base.B2i32(v790 == v828) == v828 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v1080 = v1045
	goto L116
L119:
	;
	v837 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v827))))
	v839 = v544 + int32(12)
	v856 = int32(0)
	v872 = v820
	goto L122
L120:
	;
	v1045 = v820
	goto L121
L121:
	;
	v1050 = v802 + int32(1)
	if v1050 != v542 {
		v802 = v1050
		v820 = v1045
		goto L117
	} else {
		goto L142
	}
L122:
	;
	if v837 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v1045 = v1010
	goto L121
L124:
	;
	v879 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v786+int32(2)+v856<<(uint(int32(1))%32)))))
	v881 = v879 & int32(_a_F_calc_rank_1)
	v882 = int32(12)
	v893 = int32(0)
	v916 = v872
	goto L127
L125:
	;
	v1010 = v872
	goto L126
L126:
	;
	v1015 = v856 + int32(1)
	if v1015 != v790 {
		v856 = v1015
		v872 = v1010
		goto L122
	} else {
		goto L141
	}
L127:
	;
	v923 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v827+int32(2)+v893<<(uint(int32(1))%32)))))
	v925 = v923 & int32(_a_F_calc_rank_1)
	v926 = base.B2i32(v881 != v925)
	if base.B2i32(v786 == v839)|base.B2i32(v827 == v839)|v926 == int32(0) {
		v977 = v916
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v1010 = v977
	goto L126
L129:
	;
	v980 = v893 + int32(1)
	if v980 != v837 {
		v893 = v980
		v916 = v977
		goto L127
	} else {
		goto L140
	}
L130:
	;
	v930 = *(*float32)(unsafe.Add(mBase, uint32(v535+int32(base.Ui32(v879)>>(uint(v882)%32))&v882)))
	v931 = int32(12)
	v936 = *(*float32)(unsafe.Add(mBase, uint32(v535+int32(base.Ui32(v923)>>(uint(v931)%32))&v931)))
	v938 = v881 - v925
	v940 = v938 >> (uint(int32(31)) % 32)
	if v881 != v925 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v944 = v938 ^ v940 - v940
	goto L133
L132:
	;
	v944 = int32(_a_F_calc_rank_2)
	goto L133
L133:
	;
	if base.Ui32(v944) <= base.Ui32(int32(100)) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v953 = F_exp(m, base.F64_add(base.F64_div(base.F64_convert_i32_u(v944), float64(1.5)), float64(-2)))
	mBase = m.M
	v961 = base.F32_demote_f64(base.F64_div(float64(1), base.F64_add(base.F64_mul(v953, float64(0.05)), float64(1.005))))
	goto L136
L135:
	;
	v961 = float32(1e-30)
	goto L136
L136:
	;
	v963 = base.F32_sqrt(base.F32_mul(base.F32_mul(v930, v936), v961))
	if base.F32_lt(v916, float32(0)) != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v977 = v963
	goto L129
L138:
	;
	goto L139
L139:
	;
	v966 = float64(1)
	v977 = base.F32_demote_f64(base.F64_sub(v966, base.F64_mul(base.F64_sub(v966, base.F64_promote_f32(v916)), base.F64_sub(v966, base.F64_promote_f32(v963)))))
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
	F_pfree(m, v549)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L9
	} else {
		goto L146
	}
L146:
	;
	v1130 = v536
	v1132 = v538
	v1138 = v544
	v1157 = v1118
	goto L4
L147:
	;
	v1164 = float32(1e-20)
	goto L149
L148:
	;
	v1164 = v1157
	goto L149
L149:
	;
	if v1132&int32(1) == int32(0) {
		v1274 = v1164
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if v1132&int32(2) == int32(0) {
		v1382 = v1274
		goto L162
	} else {
		goto L163
	}
L151:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+4))
	if v1169 <= int32(0) {
		v1274 = v1164
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v1173 = v1130 + int32(8)
	v1176 = v1173 + v1169<<(uint(int32(2))%32)
	v1178 = int32(0)
	v1183 = v1173
	goto L153
L153:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1183)))
	if v1210&int32(1) != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v1241 = F_log(m, base.F64_convert_i32_s(v1233+int32(1)))
	mBase = m.M
	v1274 = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(v1164), base.F64_div(v1241, float64(0.6931471805599453))))
	goto L150
L155:
	;
	v1213 = int32(1)
	v1226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1176+(int32(base.Ui32(v1210)>>(uint(v1213)%32))&int32(2047)+int32(base.Ui32(v1210)>>(uint(int32(12))%32))+v1213)&int32(_a_F_calc_rank_0)))))
	if base.Ui32(v1226) <= base.Ui32(v1213) {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v1232 = int32(1)
	goto L157
L157:
	;
	v1233 = v1232 + v1178
	v1235 = v1183 + int32(4)
	if base.Ui32(v1235) < base.Ui32(v1176) {
		v1178 = v1233
		v1183 = v1235
		goto L153
	} else {
		goto L161
	}
L158:
	;
	v1229 = v1213
	goto L160
L159:
	;
	v1229 = v1226
	goto L160
L160:
	;
	v1232 = v1229
	goto L157
L161:
	;
	goto L154
L162:
	;
	if v1132&int32(8) == int32(0) {
		v1396 = v1382
		goto L175
	} else {
		goto L176
	}
L163:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+4))
	if v1282 <= int32(0) {
		v1382 = v1274
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v1286 = v1130 + int32(8)
	v1289 = v1286 + v1282<<(uint(int32(2))%32)
	v1291 = int32(0)
	v1296 = v1286
	goto L165
L165:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1296)))
	if v1323&int32(1) != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	if v1346 <= int32(0) {
		v1382 = v1274
		goto L162
	} else {
		goto L174
	}
L167:
	;
	v1326 = int32(1)
	v1339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1289+(int32(base.Ui32(v1323)>>(uint(v1326)%32))&int32(2047)+int32(base.Ui32(v1323)>>(uint(int32(12))%32))+v1326)&int32(_a_F_calc_rank_0)))))
	if base.Ui32(v1339) <= base.Ui32(v1326) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v1345 = int32(1)
	goto L169
L169:
	;
	v1346 = v1345 + v1291
	v1348 = v1296 + int32(4)
	if base.Ui32(v1348) < base.Ui32(v1289) {
		v1291 = v1346
		v1296 = v1348
		goto L165
	} else {
		goto L173
	}
L170:
	;
	v1342 = v1326
	goto L172
L171:
	;
	v1342 = v1339
	goto L172
L172:
	;
	v1345 = v1342
	goto L169
L173:
	;
	goto L166
L174:
	;
	v1382 = base.F32_div(v1274, base.F32_convert_i32_u(v1346))
	goto L162
L175:
	;
	if v1132&int32(16) == int32(0) {
		v1414 = v1396
		goto L178
	} else {
		goto L179
	}
L176:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+4))
	if v1390 <= int32(0) {
		v1396 = v1382
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v1396 = base.F32_div(v1382, base.F32_convert_i32_u(v1390))
	goto L175
L178:
	;
	if v1132&int32(32) == int32(0) {
		v1431 = v1138
		v1450 = v1414
		goto L1
	} else {
		goto L181
	}
L179:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+4))
	if v1401 <= int32(0) {
		v1414 = v1396
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v1408 = F_log(m, base.F64_convert_i32_s(v1401+int32(1)))
	mBase = m.M
	v1414 = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(v1396), base.F64_div(v1408, float64(0.6931471805599453))))
	goto L178
L181:
	;
	v1431 = v1138
	v1450 = base.F32_div(v1414, base.F32_add(v1414, float32(1)))
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
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v14 == int32(0) {
		v84 = v13
		m.G0 = v11 - int32(-64)
		return v84
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[0])) = int32(50856066)
		v21 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[1])) = v21
		*(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[2])) = v21
		*(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[3])) = v21
		v29 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			if v29 != 0 {
				v84 = v13
				m.G0 = v11 - int32(-64)
				return v84
			} else {
				v34 = F_errstart(m, l4, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					if v34 != 0 {
						v37 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[0]))
						F_errcode(m, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[1]))
							if v41 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v41
								F_errmsg_internal(m, int32(_a_F_call_real_check_hook_0), v9+int32(-16))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[2]))
									if v59 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v59
										F_errdetail_internal(m, int32(_a_F_call_real_check_hook_0), v9+int32(-48))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[3]))
											if v67 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
												F_errhint(m, int32(_a_F_call_real_check_hook_0), v11)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_FlushErrorState(m)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															v84 = int32(0)
															m.G0 = v11 - int32(-64)
															return v84
														}
													}
												}
											} else {
												F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														v84 = int32(0)
														m.G0 = v11 - int32(-64)
														return v84
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[3]))
										if v67 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
											F_errhint(m, int32(_a_F_call_real_check_hook_0), v11)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														v84 = int32(0)
														m.G0 = v11 - int32(-64)
														return v84
													}
												}
											}
										} else {
											F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_FlushErrorState(m)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													v84 = int32(0)
													m.G0 = v11 - int32(-64)
													return v84
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
								F_errmsg(m, int32(_a_F_call_real_check_hook_4), v9+int32(-32))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[2]))
									if v59 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v59
										F_errdetail_internal(m, int32(_a_F_call_real_check_hook_0), v9+int32(-48))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[3]))
											if v67 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
												F_errhint(m, int32(_a_F_call_real_check_hook_0), v11)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_FlushErrorState(m)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															v84 = int32(0)
															m.G0 = v11 - int32(-64)
															return v84
														}
													}
												}
											} else {
												F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														v84 = int32(0)
														m.G0 = v11 - int32(-64)
														return v84
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[3]))
										if v67 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
											F_errhint(m, int32(_a_F_call_real_check_hook_0), v11)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														v84 = int32(0)
														m.G0 = v11 - int32(-64)
														return v84
													}
												}
											}
										} else {
											F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_FlushErrorState(m)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													v84 = int32(0)
													m.G0 = v11 - int32(-64)
													return v84
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
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v84 = int32(0)
							m.G0 = v11 - int32(-64)
							return v84
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
	var v554 int32
	_ = v554
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
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v610 int32
	_ = v610
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
		v554 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
		v555 = F_hash_bytes_uint32(m, v554)
		mBase = m.M
		v556 = *(*int32)(unsafe.Add(mBase, uint32(v550)+4))
		v557 = F_hash_bytes_uint32(m, v556)
		mBase = m.M
		v558 = int32(1640531527)
		v559 = v555 - v558
		v568 = v557 + v559<<(uint(int32(6))%32) + int32(base.Ui32(v559)>>(uint(int32(2))%32)) - v558 ^ v559
		v569 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
		if v551 < v569 {
			v573 = v568
			v574 = v569
			v575 = v551
			for {
				v582 = *(*int32)(unsafe.Add(mBase, uint32(v550+v574<<(uint(int32(4))%32)+v575*int32(100))+88))
				v583 = F_hash_bytes_uint32(m, v582)
				mBase = m.M
				v592 = v583 + (v573<<(uint(int32(6))%32) + int32(base.Ui32(v573)>>(uint(int32(2))%32))) - int32(1640531527) ^ v573
				v594 = v575 + int32(1)
				v595 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
				if v594 < v595 {
					v573 = v592
					v574 = v595
					v575 = v594
					continue
				} else {
					break
				}
				break
			}
			v598 = v592
		} else {
			v598 = v568
		}
		v610 = v598 + (v549<<(uint(int32(6))%32) + int32(base.Ui32(v549)>>(uint(int32(2))%32))) - int32(1640531527) ^ v549
	} else {
		v610 = v549
	}
	return v610
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
	var v115 int32
	_ = v115
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
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v247 int32
	_ = v247
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
		v247 = l3
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return base.B2i32(v247 == int32(0))
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
	v224 = v31
	v225 = v32
	goto L25
L25:
	;
	if v30 < int32(2) {
		v247 = v225
		goto L7
	} else {
		goto L53
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
	v224 = v218
	v225 = v212
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
	v211 = int32(1)
	v212 = v79 - v211
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78))))
	v218 = v78 + (v213+int32(9))&int32(_a_F_checkCond_0)
	v220 = v83 + v211
	if v220 != v61 {
		v78 = v218
		v79 = v212
		v83 = v220
		goto L26
	} else {
		goto L52
	}
L33:
	;
	if v103 != 0 {
		goto L22
	} else {
		goto L51
	}
L34:
	;
	v115 = v29 + int32(16)
	v121 = int32(0)
	goto L35
L35:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+6)))
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
		goto L50
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
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+4)))
	v170 = v121 + int32(1)
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+4)))
	if base.Ui32(v170) < base.Ui32(v171) {
		v115 = v115 + (v161+int32(7))&int32(_a_F_checkCond_0) + int32(8)
		v121 = v170
		goto L35
	} else {
		goto L49
	}
L41:
	;
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+4)))
	v141 = F_compare_subnode(m, v78, v115+int32(7), v138, v133, v130&int32(1))
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
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+4)))
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78))))
	if base.B2i32(v145 != v146)&(base.B2i32(v130&int32(1) == int32(0))|base.B2i32(base.Ui32(v146) <= base.Ui32(v145))) != 0 {
		goto L40
	} else {
		goto L46
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
	v157 = m.T0[v133].(func(*base.Module, int32, int32, int32, int32) int32)(m, v115+int32(7), v145, v78+int32(2), v146)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	if v157 != 0 {
		goto L33
	} else {
		goto L48
	}
L48:
	;
	goto L40
L49:
	;
	goto L36
L50:
	;
	return int32(0)
L51:
	;
	goto L32
L52:
	;
	goto L27
L53:
	;
	v29 = v72
	v30 = v66
	v31 = v224
	v32 = v225
	goto L9
}
func F_checkNameSpaceConflicts(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v134 int32
	_ = v134
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L36
	} else {
		goto L37
	}
L2:
	;
	m.G0 = v15 + int32(16)
	return
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = int32(0)
	if v22 < v19 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = v19
	goto L7
L6:
	;
	v25 = v22
	goto L7
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = v3
	goto L8
L8:
	;
	if l1 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v134 = v32 + int32(1)
	if v134 != v25 {
		v32 = v134
		goto L8
	} else {
		goto L35
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v26+v32<<(uint(int32(2))%32))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+20)))
	if v45&int32(1) == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v50 <= int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v55 = int32(0)
	if v55 < v50 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v58 = v50
	goto L16
L15:
	;
	v58 = v55
	goto L16
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v62 = int32(0)
	goto L17
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v60+v62<<(uint(int32(2))%32))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+20)))
	if v78 != int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L10
L19:
	;
	v119 = v62 + int32(1)
	if v119 != v58 {
		v62 = v119
		goto L17
	} else {
		goto L34
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if base.B2i32(v86 == int32(0))|base.B2i32(v86 != v89) != 0 {
		v107 = v86
		v108 = v89
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v107-v108 != 0 {
		goto L19
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v92 = v83
	v93 = v54
	goto L24
L24:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	if v97 == int32(0) {
		v107 = v97
		v108 = v96
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v107 = v97
	v108 = v96
	goto L22
L26:
	;
	v100 = int32(1)
	if v97 == v96 {
		v92 = v92 + v100
		v93 = v93 + v100
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	if v110 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v111 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if v112 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v113 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	if v114 == v115 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L19
L34:
	;
	goto L18
L35:
	;
	goto L9
L36:
	;
	return
L37:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v54
	F_errmsg(m, int32(_a_F_checkNameSpaceConflicts_0), v15)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_checkNameSpaceConflicts_1), int32(474), int32(_a_F_checkNameSpaceConflicts_2))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
							F_errfinish(m, int32(_a_F_check_assignable_2), int32(3569), int32(_a_F_check_assignable_3))
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
				F_errfinish(m, int32(_a_F_check_assignable_2), int32(3580), int32(_a_F_check_assignable_3))
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
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663)+112)))
	if v762 != 0 {
		goto L232
	} else {
		goto L233
	}
L212:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v663)+116))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v663)+100))
	if v690 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L213:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v663)+112))
	if v686 != 0 {
		goto L218
	} else {
		goto L219
	}
L214:
	;
	v676 = *(*float64)(unsafe.Add(mBase, uint32(v663)+136))
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
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v663)+120))
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
	F_errfinish(m, int32(_a_F_check_datestyle_18), int32(3036), int32(_a_F_check_datestyle_19))
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
	F_errdetail(m, int32(_a_F_check_datestyle_24), v657+int32(48))
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
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v10&int32(1) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v154
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
	v34 = int32(1)
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v38 = F_palloc(m, v35<<(uint(v34)%32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
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
	v42 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v42)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v38
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = v25 + int32(2)
	v48 = v38
	v49 = v47
	v54 = v45
	goto L13
L11:
	;
	v79 = v38
	v88 = v38
	goto L12
L12:
	;
	v91 = (v79 - v88) >> (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v91
	if int32(0) < v91 {
		v154 = v34
		goto L1
	} else {
		goto L19
	}
L13:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49))))
	if int32(base.Ui32(v57)>>(uint(int32(base.Ui32(v58)>>(uint(int32(14))%32)))%32))&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v79 = v70
	v88 = v78
	goto L12
L15:
	;
	v65 = v58 & int32(_a_F_checkclass_str_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v48))) = uint16(v65)
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v70 = v48 + int32(2)
	v71 = v67
	goto L17
L16:
	;
	v70 = v48
	v71 = v54
	goto L17
L17:
	;
	v73 = v49 + int32(2)
	if base.Ui32(v73) < base.Ui32(v47+v71<<(uint(int32(1))%32)) {
		v48 = v70
		v49 = v73
		v54 = v71
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	F_pfree(m, v88)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v97)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v97
	return v97
L21:
	;
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	if v103 == int32(0) {
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
		v154 = int32(1)
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
	v109 = v25 + int32(2)
	v114 = v109
	goto L27
L27:
	;
	v122 = int32(1)
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
	if int32(base.Ui32(v28)>>(uint(int32(base.Ui32(v123)>>(uint(int32(14))%32)))%32))&v122 != 0 {
		v154 = v122
		goto L1
	} else {
		goto L29
	}
L28:
	;
	return int32(0)
L29:
	;
	v130 = v114 + int32(2)
	if base.Ui32(v130) < base.Ui32(v109+v103<<(uint(int32(1))%32)) {
		v114 = v130
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v25 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v137
	v142 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v142)
	return int32(1)
L32:
	;
	v148 = int32(2)
	goto L34
L33:
	;
	v148 = int32(1)
	goto L34
L34:
	;
	v154 = v148
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	v4 = int32(0)
	if base.B2i32(l0 == v4)|base.B2i32(l1 == v4) != 0 {
		v164 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v164 & int32(1)
L2:
	;
	v12 = l0
	v13 = l1
	v14 = l2
	v16 = v4
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v18 == int32(27) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v164 = v159
	goto L1
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v22 = v21
	goto L7
L6:
	;
	v22 = v13
	goto L7
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v23 != int32(27) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v27 = v12
	goto L10
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v27 = v26
	goto L10
L10:
	;
	v28 = F_equal(m, v27, v22)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	if v28 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v164 = int32(1)
	goto L1
L14:
	;
	goto L15
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v33 == int32(17) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v37 = F_op_strict(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L19
	}
L17:
	;
	v65 = v33
	goto L18
L18:
	;
	if v65 == int32(15) {
		goto L35
	} else {
		goto L36
	}
L19:
	;
	if v37 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v39 == int32(0) {
		v164 = v16
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v65 = v64
	goto L18
L23:
	;
	v42 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v43 <= v42 {
		v164 = v16
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v46 = v42
	goto L25
L25:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v46<<(uint(int32(2))%32))))
	v58 = F_clause_is_strict_for(m, v56, v22, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L27
	}
L26:
	;
	v164 = v58
	goto L1
L27:
	;
	if v58 != 0 {
		v164 = v58
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v61 = v46 + int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v61 < v62 {
		v46 = v61
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	v155 = int32(0)
	if v151 == v155 {
		goto L68
	} else {
		goto L69
	}
L31:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+24)))
	v164 = v150
	goto L1
L32:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v106 = F_clause_is_strict_for(m, v104, v22, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L11
	} else {
		goto L47
	}
L33:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v151 = v100
	goto L30
L34:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v75 == int32(0) {
		v164 = v16
		goto L1
	} else {
		goto L40
	}
L35:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v69 = F_func_strict(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L11
	} else {
		goto L38
	}
L36:
	;
	v72 = v65
	goto L37
L37:
	;
	switch v72 - int32(7) {
	case 0:
		goto L31
	default:
		v164 = v16
		goto L1
	case 13:
		goto L32
	case 21, 22, 23, 48:
		goto L33
	}
L38:
	;
	if v69 != 0 {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v72 = v71
	goto L37
L40:
	;
	v78 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v79 <= v78 {
		v164 = v16
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v82 = v78
	goto L42
L42:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v82<<(uint(int32(2))%32))))
	v94 = F_clause_is_strict_for(m, v92, v22, int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L11
	} else {
		goto L44
	}
L43:
	;
	v164 = v94
	goto L1
L44:
	;
	if v94 != 0 {
		v164 = v94
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v97 = v82 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v97 < v98 {
		v82 = v97
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	if v106 == int32(0) {
		v151 = v103
		goto L30
	} else {
		goto L48
	}
L48:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v111 = F_op_strict(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	if v111 == int32(0) {
		v151 = v103
		goto L30
	} else {
		goto L50
	}
L50:
	;
	if v14&int32(1) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v103 == int32(0) {
		v164 = v16
		goto L1
	} else {
		goto L54
	}
L52:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+20)))
	if v119 == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v164 = int32(1)
	goto L1
L54:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v125 != int32(35) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v146 <= int32(0) {
		v151 = v103
		goto L30
	} else {
		goto L67
	}
L56:
	;
	if v125 != int32(7) {
		v151 = v103
		goto L30
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)))
	if v140 != 0 {
		v151 = v103
		goto L30
	} else {
		goto L65
	}
L59:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+24)))
	if v130 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v164 = int32(1)
	goto L1
L61:
	;
	goto L62
L62:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v103)+20))
	v133 = F_pg_detoast_datum(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v138 = F_ArrayGetNItemsSafe(m, v135, v133+int32(16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	v146 = v138
	goto L55
L65:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	if v141 == int32(0) {
		v151 = v103
		goto L30
	} else {
		goto L66
	}
L66:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v146 = v144
	goto L55
L67:
	;
	v164 = int32(1)
	goto L1
L68:
	;
	v164 = int32(0)
	goto L1
L69:
	;
	goto L70
L70:
	;
	v159 = int32(0)
	if v22 != 0 {
		v12 = v151
		v13 = v22
		v14 = v155
		v16 = v159
		goto L3
	} else {
		goto L71
	}
L71:
	;
	goto L4
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
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
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
			v123 = l0
			m.G0 = v11 + int32(16)
			return v123
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
						v123 = l0
						m.G0 = v11 + int32(16)
						return v123
					} else {
						F_freetree(m, l0)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							v123 = int32(0)
							m.G0 = v11 + int32(16)
							return v123
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
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									v123 = int32(0)
									m.G0 = v11 + int32(16)
									return v123
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
									v123 = v84
									m.G0 = v11 + int32(16)
									return v123
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
									v123 = v99
									m.G0 = v11 + int32(16)
									return v123
								}
							} else {
								if v50 != int32(4) {
									v123 = l0
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
									v123 = l0
								}
								m.G0 = v11 + int32(16)
								return v123
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
				v123 = int32(0)
				m.G0 = v11 + int32(16)
				return v123
			}
		}
	}
}
func F_clogsyncfiletag(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_SlruSyncFileTag(m, int32(_a_F_clogsyncfiletag_0), l0, l1)
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v83 int32
	_ = v83
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
		v83 = v8
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v83
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
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if base.B2i32(v25 == int32(0))|base.B2i32(v25 != v28) != 0 {
		v46 = v25
		v47 = v28
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	v50 = F_strlen(m, v20)
	mBase = m.M
	v51 = F_strlen(m, v19)
	mBase = m.M
	v52 = v50
	v53 = v51
	goto L16
L9:
	;
	return v46 - v47
L10:
	;
	goto L9
L11:
	;
	v31 = v20
	v32 = v19
	goto L12
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	if v36 == int32(0) {
		v46 = v36
		v47 = v35
		goto L10
	} else {
		goto L14
	}
L13:
	;
	v46 = v36
	v47 = v35
	goto L10
L14:
	;
	v39 = int32(1)
	if v36 == v35 {
		v31 = v31 + v39
		v32 = v32 + v39
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v59 = int32(1)
	v60 = v52 - v59
	v62 = v53 - v59
	if int32(0) <= v60|v62 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v60 < v62 {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v20))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v19))))
	if base.Ui32(v67) < base.Ui32(v69) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	return int32(-1)
L22:
	;
	goto L23
L23:
	;
	if base.Ui32(v67) <= base.Ui32(v69) {
		v52 = v60
		v53 = v62
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v83 = v8
	goto L4
L25:
	;
	return int32(-1)
L26:
	;
	goto L27
L27:
	;
	v83 = base.B2i32(v62 < v60)
	goto L4
}
func F_cnt_sml(m *base.Module, l0 int32, l1 int32, l2 int32) float32 {
	mBase := m.M
	_ = mBase
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
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
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v94 float32
	_ = v94
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = int32(2)
	v14 = int32(5)
	v15 = int32(base.Ui32(v11)>>(uint(v12)%32)) - v14
	v16 = int32(3)
	v17 = base.I32_div_u_s(v15, v16)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = int32(base.Ui32(v18)>>(uint(v12)%32)) - v14
	v24 = base.I32_div_u_s(v22, v16)
	if base.B2i32(base.Ui32(v22) < base.Ui32(v16))|base.B2i32(base.Ui32(v15) < base.Ui32(v16)) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v94 = float32(0)
	goto L3
L2:
	;
	v31 = int32(5)
	v32 = l0 + v31
	v34 = l1 + v31
	v35 = v32
	v36 = v34
	v38 = int32(0)
	goto L4
L3:
	;
	return v94
L4:
	;
	v47 = base.I32_div_s(v36-v34, int32(3))
	if v47 < v17 {
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
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_cnt_sml[0]))
	v51 = m.T0[v50].(func(*base.Module, int32, int32) int32)(m, v35, v36)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v76 = v38
	goto L8
L8:
	;
	goto L5
L9:
	;
	v72 = base.I32_div_s(v67-v32, int32(3))
	if v72 < v24 {
		v35 = v67
		v36 = v68
		v38 = v69
		goto L4
	} else {
		goto L18
	}
L10:
	;
	return float32(0)
L11:
	;
	if int32(0) <= v51 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v51 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v63 = v36
	v64 = v38
	goto L14
L14:
	;
	v67 = v35 + int32(3)
	v68 = v63
	v69 = v64
	goto L9
L15:
	;
	v67 = v35
	v68 = v36 + int32(3)
	v69 = v38
	goto L9
L16:
	;
	goto L17
L17:
	;
	v63 = v36 + int32(3)
	v64 = v38 + int32(1)
	goto L14
L18:
	;
	v76 = v69
	goto L8
L19:
	;
	v80 = v76
	goto L21
L20:
	;
	v80 = v17
	goto L21
L21:
	;
	v94 = base.F32_div(base.F32_convert_i32_s(v76), base.F32_convert_i32_s(v24-v76+v80))
	goto L3
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
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
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
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+119)))
	v28 = v26 - int32(109)
	v35 = int32(0)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v28))|base.B2i32(int32(1)<<(uint(v28)%32)&int32(161) == v35) == v35 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v41 = F_RelationGetNumberOfBlocksInFork(m, v23, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L52
	}
L7:
	;
	v45 = F_palloc0(m, v41+int32(8))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = int32(0)
	if l1 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v41
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v51
	v59 = F_read_stream_begin_relation(m, int32(12), v18, v23, v51, int32(120), v13+int32(4), v51)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v61 = v3
	goto L11
L11:
	;
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v61 = v59
	goto L11
L13:
	;
	v63 = v45 + int32(8)
	v65 = int32(0)
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
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_collect_visibility_data[0]))
	if v76 != 0 {
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
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v81 = F_visibilitymap_get_status(m, v23, v65, v13+int32(12))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	if v81&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v85 = v65 + v63
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v88 = v86 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v88)
	goto L25
L24:
	;
	goto L25
L25:
	;
	if v81&int32(2) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v93 = v65 + v63
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	v96 = v94 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v96)
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
	v100 = F_read_stream_next_buffer(m, v61, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v137 = v65 + int32(1)
	if v137 != v41 {
		v65 = v137
		goto L16
	} else {
		goto L42
	}
L32:
	;
	F_LockBuffer(m, v100, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v100 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+10)))
	if v123&int32(4) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_collect_visibility_data[1]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108+(v100^int32(-1))<<(uint(int32(2))%32))))
	v122 = v114
	goto L34
L36:
	;
	goto L37
L37:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_collect_visibility_data[2]))
	v122 = v116 + v100<<(uint(int32(13))%32) + int32(-8192)
	goto L34
L38:
	;
	v126 = v65 + v63
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v129 = v127 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v129)
	goto L40
L39:
	;
	goto L40
L40:
	;
	F_UnlockReleaseBuffer(m, v100)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
	F_read_stream_end(m, v61)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v151 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	F_ReleaseBuffer(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
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
	v156 = m.ExcPending
	if v156 != 0 {
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
	return v45
L52:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v168 + int32(4)
	F_errmsg(m, int32(_a_F_collect_visibility_data_0), v13)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v176 = int32(*(*int8)(unsafe.Add(mBase, uint32(v175)+119)))
	F_errdetail_relkind_not_supported(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_collect_visibility_data_1), int32(951), int32(_a_F_collect_visibility_data_2))
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
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v49 int32
	_ = v49
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v7 != v8 {
		if v7 < v8 {
			v13 = int32(-1)
		} else {
			v13 = int32(1)
		}
		return v13
	} else {
		if base.F64_le(l2, float64(0))|base.F64_ge(l2, float64(1)) != 0 {
			v20 = int32(-1)
			v21 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			if base.F64_lt(v21, v22) != 0 {
				v49 = v20
				return v49
			} else {
				if base.F64_gt(v21, v22) != 0 {
					return int32(1)
				} else {
					v27 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
					v28 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
					if base.F64_lt(v27, v28) != 0 {
						v49 = v20
						return v49
					} else {
						if base.F64_gt(v27, v28) != 0 {
							v49 = int32(1)
							return v49
						} else {
							return int32(0)
						}
					}
				}
			}
		} else {
			v35 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
			v36 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
			v39 = base.F64_add(base.F64_mul(l2, base.F64_sub(v35, v36)), v36)
			v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
			v44 = base.F64_add(base.F64_mul(l2, base.F64_sub(v40, v41)), v41)
			if base.F64_lt(v39, v44) != 0 {
				v49 = int32(-1)
			} else {
				v49 = base.F64_lt(v44, v39)
			}
			return v49
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
	var v78 int32
	_ = v78
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
	v78 = v17
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
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
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
	v103 = v78
	goto L33
L27:
	;
	if base.Ui32(v78) < base.Ui32(v19) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v95 = F_pg_mblen_range(m, v78, v19)
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
	v97 = v95 + v78
	if base.Ui32(v97) < base.Ui32(v19) {
		v78 = v97
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
	v129 = v115 - v78
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
	v141 = m.T0[l3].(func(*base.Module, int32, int32, int32, int32) int32)(m, v21, v77, v78, v129)
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
	v143 = v78 + v129
	if base.Ui32(v143) < base.Ui32(v19) {
		v78 = v143
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
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
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
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
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v716 int32
	_ = v716
	var v718 int32
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
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
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
	var v821 int32
	_ = v821
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
	var v842 int32
	_ = v842
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v900 int32
	_ = v900
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1091 int32
	_ = v1091
	v10 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	if l1&int32(2048) != 0 {
		v1012 = l4
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errfinish(m, int32(_a_F_compute_new_xmax_infomask_0), int32(_a_F_compute_new_xmax_infomask_1), int32(_a_F_compute_new_xmax_infomask_2))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L13
	} else {
		goto L331
	}
L2:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l7))) = uint16(v1054)
	*(*uint16)(unsafe.Add(mBase, uint32(l8))) = uint16(v1050)
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1052
	m.G0 = v19 + int32(80)
	return
L3:
	;
	if l5 != 0 {
		goto L318
	} else {
		goto L319
	}
L4:
	;
	if l1&int32(_a_F_compute_new_xmax_infomask_3) != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L13
	} else {
		goto L313
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L13
	} else {
		goto L308
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L13
	} else {
		goto L303
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L13
	} else {
		goto L298
	}
L9:
	;
	if l1&int32(_a_F_compute_new_xmax_infomask_4) == int32(_a_F_compute_new_xmax_infomask_5) {
		v1012 = l4
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v430 = l2 & int32(_a_F_compute_new_xmax_infomask_6)
	if v430 != 0 {
		goto L116
	} else {
		goto L117
	}
L12:
	;
	v30 = l1 & int32(128)
	v33 = F_MultiXactIdIsRunning(m, l0, int32(base.Ui32(v30)>>(uint(int32(7))%32)))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v33 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v30 != 0 {
		v1012 = l4
		goto L3
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if l5 != 0 {
		goto L32
	} else {
		goto L33
	}
L18:
	;
	v37 = int32(0)
	v41 = F_GetMultiXactIdMembers(m, l0, v19+int32(76), v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	if int32(0) < v41 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	v47 = v37
	goto L25
L21:
	;
	v78 = v37
	goto L22
L22:
	;
	v93 = F_TransactionIdDidCommit(m, v78)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L13
	} else {
		goto L30
	}
L23:
	;
	F_pfree(m, v45)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L13
	} else {
		goto L29
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v74 = v72
	goto L23
L25:
	;
	v64 = v45 + v47<<(uint(int32(3))%32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v65) {
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v74 = int32(0)
	goto L23
L27:
	;
	v69 = v47 + int32(1)
	if v69 != v41 {
		v47 = v69
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v78 = v74
	goto L22
L30:
	;
	if v93 == int32(0) {
		v1012 = l4
		goto L3
	} else {
		goto L31
	}
L31:
	;
	goto L17
L32:
	;
	v117 = int32(8)
	goto L34
L33:
	;
	v117 = int32(4)
	goto L34
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v117)+uint32(_c_F_compute_new_xmax_infomask[0])))
	if v119 == int32(-1) {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	v122 = int32(0)
	v124 = m.G0
	v126 = v124 - int32(16)
	m.G0 = v126
	v131 = F_GetMultiXactIdMembers(m, l0, v126+int32(12), v122)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L13
	} else {
		goto L39
	}
L36:
	;
	m.G0 = v126 + int32(16)
	v286 = F_GetMultiXactIdMembers(m, v264, v19+int32(76), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L13
	} else {
		goto L71
	}
L37:
	;
	v253 = v240 + v235<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = l3
	v258 = F_MultiXactIdCreateFromMembers(m, v235+int32(1), v240)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L13
	} else {
		goto L68
	}
L38:
	;
	v233 = F_palloc(m, int32(8))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L13
	} else {
		goto L67
	}
L39:
	;
	if int32(0) <= v131 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	if v131 == int32(0) {
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = l3
	v229 = F_MultiXactIdCreateFromMembers(m, int32(1), v126+int32(4))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L13
	} else {
		goto L66
	}
L43:
	;
	v142 = v122
	goto L44
L44:
	;
	v156 = v135 + v142<<(uint(int32(3))%32)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v157 != l3 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v166 = int32(1)
	if v131 <= v166 {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	v164 = v142 + int32(1)
	if v164 != v131 {
		v142 = v164
		goto L44
	} else {
		goto L50
	}
L47:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v159 != v119 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	F_pfree(m, v135)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	v264 = l0
	goto L36
L50:
	;
	goto L45
L51:
	;
	v169 = v166
	goto L53
L52:
	;
	v169 = v131
	goto L53
L53:
	;
	v175 = F_palloc(m, v131<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L13
	} else {
		goto L54
	}
L54:
	;
	v178 = int32(0)
	v182 = int32(0)
	goto L55
L55:
	;
	v196 = v135 + v182<<(uint(int32(3))%32)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v198 = F_TransactionIdIsInProgress(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L13
	} else {
		goto L58
	}
L56:
	;
	v235 = v219
	v240 = v175
	goto L37
L57:
	;
	v222 = v182 + int32(1)
	if v222 != v169 {
		v178 = v219
		v182 = v222
		goto L55
	} else {
		goto L65
	}
L58:
	;
	if v198 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if base.Ui32(v202) < base.Ui32(int32(4)) {
		v219 = v178
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v212 = v175 + v178<<(uint(int32(3))%32)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v213
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v212)+4)) = v215
	v219 = v178 + int32(1)
	goto L57
L62:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v206 = F_TransactionIdDidCommit(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	if v206 == int32(0) {
		v219 = v178
		goto L57
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	goto L56
L66:
	;
	v264 = v229
	goto L36
L67:
	;
	v235 = int32(0)
	v240 = v233
	goto L37
L68:
	;
	F_pfree(m, v135)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	F_pfree(m, v240)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	v264 = v258
	goto L36
L71:
	;
	if v286 <= int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v1050 = v122
	v1052 = v264
	v1054 = int32(_a_F_compute_new_xmax_infomask_7)
	goto L2
L73:
	;
	goto L74
L74:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v286 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	F_pfree(m, v291)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L13
	} else {
		goto L104
	}
L76:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v291+v365<<(uint(int32(3))%32))+4))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v380<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v366) < base.Ui32(v383) {
		goto L98
	} else {
		goto L99
	}
L77:
	;
	v294 = int32(0)
	v362 = v122
	v365 = v294
	v366 = v294
	v373 = v10
	goto L76
L78:
	;
	goto L79
L79:
	;
	v300 = int32(0)
	v303 = v300
	v304 = v122
	v307 = v300
	v308 = v300
	v315 = v10
	goto L80
L80:
	;
	v321 = v291 + v307<<(uint(int32(3))%32)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v322<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v308) < base.Ui32(v325) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v286&int32(1) == int32(0) {
		v395 = v351
		v399 = v353
		v406 = v352
		goto L75
	} else {
		goto L97
	}
L82:
	;
	v327 = v325
	goto L84
L83:
	;
	v327 = v308
	goto L84
L84:
	;
	switch v322 - int32(3) {
	case 0:
		goto L88
	case 1:
		v334 = v304
		goto L86
	case 2:
		goto L87
	default:
		v336 = v304
		v337 = v315
		goto L85
	}
L85:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v321)+12))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v338<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	switch v338 - int32(3) {
	case 0:
		goto L92
	case 1:
		v349 = v336
		goto L90
	case 2:
		goto L91
	default:
		v351 = v336
		v352 = v337
		goto L89
	}
L86:
	;
	v336 = v334
	v337 = int32(1)
	goto L85
L87:
	;
	v334 = v304 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L86
L88:
	;
	v336 = v304 | int32(_a_F_compute_new_xmax_infomask_6)
	v337 = v315
	goto L85
L89:
	;
	if base.Ui32(v327) < base.Ui32(v341) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v351 = v349
	v352 = int32(1)
	goto L89
L91:
	;
	v349 = v336 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L90
L92:
	;
	v351 = v336 | int32(_a_F_compute_new_xmax_infomask_6)
	v352 = v337
	goto L89
L93:
	;
	v353 = v341
	goto L95
L94:
	;
	v353 = v327
	goto L95
L95:
	;
	v354 = int32(2)
	v355 = v307 + v354
	v357 = v303 + v354
	if v357 != v286&int32(2147483646) {
		v303 = v357
		v304 = v351
		v307 = v355
		v308 = v353
		v315 = v352
		goto L80
	} else {
		goto L96
	}
L96:
	;
	goto L81
L97:
	;
	v362 = v351
	v365 = v355
	v366 = v353
	v373 = v352
	goto L76
L98:
	;
	v385 = v383
	goto L100
L99:
	;
	v385 = v366
	goto L100
L100:
	;
	switch v380 - int32(3) {
	case 0:
		goto L103
	case 1:
		v392 = v362
		goto L101
	case 2:
		goto L102
	default:
		v395 = v362
		v399 = v385
		v406 = v373
		goto L75
	}
L101:
	;
	v395 = v392
	v399 = v385
	v406 = int32(1)
	goto L75
L102:
	;
	v392 = v362 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L101
L103:
	;
	v395 = v362 | int32(_a_F_compute_new_xmax_infomask_6)
	v399 = v385
	v406 = v373
	goto L75
L104:
	;
	if v399&int32(-2) == int32(2) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	if v406 != 0 {
		v1050 = v395
		v1052 = v264
		v1054 = int32(_a_F_compute_new_xmax_infomask_8)
		goto L2
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v399 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v1050 = v395
	v1052 = v264
	v1054 = int32(_a_F_compute_new_xmax_infomask_9)
	goto L2
L109:
	;
	v421 = int32(_a_F_compute_new_xmax_infomask_3)
	goto L111
L110:
	;
	v421 = int32(_a_F_compute_new_xmax_infomask_10)
	goto L111
L111:
	;
	if v399 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v424 = int32(_a_F_compute_new_xmax_infomask_11)
	goto L114
L113:
	;
	v424 = v421
	goto L114
L114:
	;
	if v406 != 0 {
		v1050 = v395
		v1052 = v264
		v1054 = v424
		goto L2
	} else {
		goto L115
	}
L115:
	;
	v1050 = v395
	v1052 = v264
	v1054 = v424 | int32(128)
	goto L2
L116:
	;
	v431 = int32(5)
	goto L118
L117:
	;
	v431 = int32(4)
	goto L118
L118:
	;
	if l1&int32(1024) != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	if l5 != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	v598 = int32(base.Ui32(l1&int32(128))>>(uint(int32(7))%32)) | base.B2i32(l1&int32(80) == int32(64))
	v599 = F_TransactionIdIsInProgress(m, l0)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L13
	} else {
		goto L172
	}
L122:
	;
	v438 = int32(8)
	goto L124
L123:
	;
	v438 = int32(4)
	goto L124
L124:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v438)+uint32(_c_F_compute_new_xmax_infomask[0])))
	if v440 == int32(-1) {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	v443 = int32(0)
	v444 = F_MultiXactIdCreate(m, l0, v431, l3, v440)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L13
	} else {
		goto L126
	}
L126:
	;
	v449 = F_GetMultiXactIdMembers(m, v444, v19+int32(76), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L13
	} else {
		goto L127
	}
L127:
	;
	if v449 <= int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v1050 = v443
	v1052 = v444
	v1054 = int32(_a_F_compute_new_xmax_infomask_7)
	goto L2
L129:
	;
	goto L130
L130:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v449 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	F_pfree(m, v454)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L13
	} else {
		goto L160
	}
L132:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v454+v528<<(uint(int32(3))%32))+4))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v543<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v529) < base.Ui32(v546) {
		goto L154
	} else {
		goto L155
	}
L133:
	;
	v457 = int32(0)
	v525 = v443
	v528 = v457
	v529 = v457
	v536 = v10
	goto L132
L134:
	;
	goto L135
L135:
	;
	v463 = int32(0)
	v466 = v463
	v467 = v443
	v470 = v463
	v471 = v463
	v478 = v10
	goto L136
L136:
	;
	v484 = v454 + v470<<(uint(int32(3))%32)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+4))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v485<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v471) < base.Ui32(v488) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v449&int32(1) == int32(0) {
		v558 = v514
		v562 = v516
		v569 = v515
		goto L131
	} else {
		goto L153
	}
L138:
	;
	v490 = v488
	goto L140
L139:
	;
	v490 = v471
	goto L140
L140:
	;
	switch v485 - int32(3) {
	case 0:
		goto L144
	case 1:
		v497 = v467
		goto L142
	case 2:
		goto L143
	default:
		v499 = v467
		v500 = v478
		goto L141
	}
L141:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v484)+12))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v501<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	switch v501 - int32(3) {
	case 0:
		goto L148
	case 1:
		v512 = v499
		goto L146
	case 2:
		goto L147
	default:
		v514 = v499
		v515 = v500
		goto L145
	}
L142:
	;
	v499 = v497
	v500 = int32(1)
	goto L141
L143:
	;
	v497 = v467 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L142
L144:
	;
	v499 = v467 | int32(_a_F_compute_new_xmax_infomask_6)
	v500 = v478
	goto L141
L145:
	;
	if base.Ui32(v490) < base.Ui32(v504) {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v514 = v512
	v515 = int32(1)
	goto L145
L147:
	;
	v512 = v499 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L146
L148:
	;
	v514 = v499 | int32(_a_F_compute_new_xmax_infomask_6)
	v515 = v500
	goto L145
L149:
	;
	v516 = v504
	goto L151
L150:
	;
	v516 = v490
	goto L151
L151:
	;
	v517 = int32(2)
	v518 = v470 + v517
	v520 = v466 + v517
	if v520 != v449&int32(2147483646) {
		v466 = v520
		v467 = v514
		v470 = v518
		v471 = v516
		v478 = v515
		goto L136
	} else {
		goto L152
	}
L152:
	;
	goto L137
L153:
	;
	v525 = v514
	v528 = v518
	v529 = v516
	v536 = v515
	goto L132
L154:
	;
	v548 = v546
	goto L156
L155:
	;
	v548 = v529
	goto L156
L156:
	;
	switch v543 - int32(3) {
	case 0:
		goto L159
	case 1:
		v555 = v525
		goto L157
	case 2:
		goto L158
	default:
		v558 = v525
		v562 = v548
		v569 = v536
		goto L131
	}
L157:
	;
	v558 = v555
	v562 = v548
	v569 = int32(1)
	goto L131
L158:
	;
	v555 = v525 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L157
L159:
	;
	v558 = v525 | int32(_a_F_compute_new_xmax_infomask_6)
	v562 = v548
	v569 = v536
	goto L131
L160:
	;
	if v562&int32(-2) == int32(2) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	if v569 != 0 {
		v1050 = v558
		v1052 = v444
		v1054 = int32(_a_F_compute_new_xmax_infomask_8)
		goto L2
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	if v562 != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v1050 = v558
	v1052 = v444
	v1054 = int32(_a_F_compute_new_xmax_infomask_9)
	goto L2
L165:
	;
	v584 = int32(_a_F_compute_new_xmax_infomask_3)
	goto L167
L166:
	;
	v584 = int32(_a_F_compute_new_xmax_infomask_10)
	goto L167
L167:
	;
	if v562 == int32(1) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v587 = int32(_a_F_compute_new_xmax_infomask_11)
	goto L170
L169:
	;
	v587 = v584
	goto L170
L170:
	;
	if v569 != 0 {
		v1050 = v558
		v1052 = v444
		v1054 = v587
		goto L2
	} else {
		goto L171
	}
L171:
	;
	v1050 = v558
	v1052 = v444
	v1054 = v587 | int32(128)
	goto L2
L172:
	;
	if v599 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	if v598 == int32(0) {
		v631 = v431
		goto L176
	} else {
		goto L177
	}
L174:
	;
	goto L175
L175:
	;
	if v598 != 0 {
		v1012 = l4
		goto L3
	} else {
		goto L245
	}
L176:
	;
	if l0 == l3 {
		goto L189
	} else {
		goto L190
	}
L177:
	;
	switch int32(base.Ui32(l1)>>(uint(int32(4))%32))&int32(5) - int32(1) {
	case 0:
		v631 = int32(0)
		goto L176
	case 1, 2:
		goto L180
	case 3:
		goto L181
	case 4:
		goto L178
	default:
		goto L179
	}
L178:
	;
	v631 = int32(1)
	goto L176
L179:
	;
	v615 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L13
	} else {
		goto L185
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	if v430 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v612 = int32(3)
	goto L184
L183:
	;
	v612 = int32(2)
	goto L184
L184:
	;
	v631 = v612
	goto L176
L185:
	;
	if v615 == int32(0) {
		v1012 = l4
		goto L3
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_12), v19+int32(16))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L13
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_compute_new_xmax_infomask_0), int32(_a_F_compute_new_xmax_infomask_13), int32(_a_F_compute_new_xmax_infomask_14))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L13
	} else {
		goto L188
	}
L188:
	;
	v1012 = l4
	goto L3
L189:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v631<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v635) < base.Ui32(l4) {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	goto L191
L191:
	;
	if l5 != 0 {
		goto L195
	} else {
		goto L196
	}
L192:
	;
	v637 = l4
	goto L194
L193:
	;
	v637 = v635
	goto L194
L194:
	;
	v1012 = v637
	goto L3
L195:
	;
	v642 = int32(8)
	goto L197
L196:
	;
	v642 = int32(4)
	goto L197
L197:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v642)+uint32(_c_F_compute_new_xmax_infomask[0])))
	if v644 == int32(-1) {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	v647 = int32(0)
	v648 = F_MultiXactIdCreate(m, l0, v631, l3, v644)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L13
	} else {
		goto L199
	}
L199:
	;
	v653 = F_GetMultiXactIdMembers(m, v648, v19+int32(76), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L13
	} else {
		goto L200
	}
L200:
	;
	if v653 <= int32(0) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1050 = v647
	v1052 = v648
	v1054 = int32(_a_F_compute_new_xmax_infomask_7)
	goto L2
L202:
	;
	goto L203
L203:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v653 == int32(1) {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	F_pfree(m, v658)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L13
	} else {
		goto L233
	}
L205:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v658+v732<<(uint(int32(3))%32))+4))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v747<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v733) < base.Ui32(v750) {
		goto L227
	} else {
		goto L228
	}
L206:
	;
	v661 = int32(0)
	v729 = v647
	v732 = v661
	v733 = v661
	v740 = v10
	goto L205
L207:
	;
	goto L208
L208:
	;
	v667 = int32(0)
	v670 = v667
	v671 = v647
	v674 = v667
	v675 = v667
	v682 = v10
	goto L209
L209:
	;
	v688 = v658 + v674<<(uint(int32(3))%32)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v688)+4))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v689<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v675) < base.Ui32(v692) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	if v653&int32(1) == int32(0) {
		v762 = v718
		v766 = v720
		v773 = v719
		goto L204
	} else {
		goto L226
	}
L211:
	;
	v694 = v692
	goto L213
L212:
	;
	v694 = v675
	goto L213
L213:
	;
	switch v689 - int32(3) {
	case 0:
		goto L217
	case 1:
		v701 = v671
		goto L215
	case 2:
		goto L216
	default:
		v703 = v671
		v704 = v682
		goto L214
	}
L214:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v688)+12))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v705<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	switch v705 - int32(3) {
	case 0:
		goto L221
	case 1:
		v716 = v703
		goto L219
	case 2:
		goto L220
	default:
		v718 = v703
		v719 = v704
		goto L218
	}
L215:
	;
	v703 = v701
	v704 = int32(1)
	goto L214
L216:
	;
	v701 = v671 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L215
L217:
	;
	v703 = v671 | int32(_a_F_compute_new_xmax_infomask_6)
	v704 = v682
	goto L214
L218:
	;
	if base.Ui32(v694) < base.Ui32(v708) {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	v718 = v716
	v719 = int32(1)
	goto L218
L220:
	;
	v716 = v703 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L219
L221:
	;
	v718 = v703 | int32(_a_F_compute_new_xmax_infomask_6)
	v719 = v704
	goto L218
L222:
	;
	v720 = v708
	goto L224
L223:
	;
	v720 = v694
	goto L224
L224:
	;
	v721 = int32(2)
	v722 = v674 + v721
	v724 = v670 + v721
	if v724 != v653&int32(2147483646) {
		v670 = v724
		v671 = v718
		v674 = v722
		v675 = v720
		v682 = v719
		goto L209
	} else {
		goto L225
	}
L225:
	;
	goto L210
L226:
	;
	v729 = v718
	v732 = v722
	v733 = v720
	v740 = v719
	goto L205
L227:
	;
	v752 = v750
	goto L229
L228:
	;
	v752 = v733
	goto L229
L229:
	;
	switch v747 - int32(3) {
	case 0:
		goto L232
	case 1:
		v759 = v729
		goto L230
	case 2:
		goto L231
	default:
		v762 = v729
		v766 = v752
		v773 = v740
		goto L204
	}
L230:
	;
	v762 = v759
	v766 = v752
	v773 = int32(1)
	goto L204
L231:
	;
	v759 = v729 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L230
L232:
	;
	v762 = v729 | int32(_a_F_compute_new_xmax_infomask_6)
	v766 = v752
	v773 = v740
	goto L204
L233:
	;
	if v766&int32(-2) == int32(2) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	if v773 != 0 {
		v1050 = v762
		v1052 = v648
		v1054 = int32(_a_F_compute_new_xmax_infomask_8)
		goto L2
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	if v766 != 0 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v1050 = v762
	v1052 = v648
	v1054 = int32(_a_F_compute_new_xmax_infomask_9)
	goto L2
L238:
	;
	v788 = int32(_a_F_compute_new_xmax_infomask_3)
	goto L240
L239:
	;
	v788 = int32(_a_F_compute_new_xmax_infomask_10)
	goto L240
L240:
	;
	if v766 == int32(1) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v791 = int32(_a_F_compute_new_xmax_infomask_11)
	goto L243
L242:
	;
	v791 = v788
	goto L243
L243:
	;
	if v773 != 0 {
		v1050 = v762
		v1052 = v648
		v1054 = v791
		goto L2
	} else {
		goto L244
	}
L244:
	;
	v1050 = v762
	v1052 = v648
	v1054 = v791 | int32(128)
	goto L2
L245:
	;
	v794 = F_TransactionIdDidCommit(m, l0)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L13
	} else {
		goto L246
	}
L246:
	;
	if v794 == int32(0) {
		v1012 = l4
		goto L3
	} else {
		goto L247
	}
L247:
	;
	if l5 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v802 = int32(8)
	goto L250
L249:
	;
	v802 = int32(4)
	goto L250
L250:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v802)+uint32(_c_F_compute_new_xmax_infomask[0])))
	if v804 == int32(-1) {
		goto L5
	} else {
		goto L251
	}
L251:
	;
	v807 = int32(0)
	v808 = F_MultiXactIdCreate(m, l0, v431, l3, v804)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L13
	} else {
		goto L252
	}
L252:
	;
	v813 = F_GetMultiXactIdMembers(m, v808, v19+int32(76), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L13
	} else {
		goto L253
	}
L253:
	;
	if v813 <= int32(0) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1050 = v807
	v1052 = v808
	v1054 = int32(_a_F_compute_new_xmax_infomask_7)
	goto L2
L255:
	;
	goto L256
L256:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v813 == int32(1) {
		goto L259
	} else {
		goto L260
	}
L257:
	;
	F_pfree(m, v818)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L13
	} else {
		goto L286
	}
L258:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v818+v892<<(uint(int32(3))%32))+4))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v907<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v893) < base.Ui32(v910) {
		goto L280
	} else {
		goto L281
	}
L259:
	;
	v821 = int32(0)
	v889 = v807
	v892 = v821
	v893 = v821
	v900 = v10
	goto L258
L260:
	;
	goto L261
L261:
	;
	v827 = int32(0)
	v830 = v827
	v831 = v807
	v834 = v827
	v835 = v827
	v842 = v10
	goto L262
L262:
	;
	v848 = v818 + v834<<(uint(int32(3))%32)
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v848)+4))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v849<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v835) < base.Ui32(v852) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	if v813&int32(1) == int32(0) {
		v922 = v878
		v926 = v880
		v933 = v879
		goto L257
	} else {
		goto L279
	}
L264:
	;
	v854 = v852
	goto L266
L265:
	;
	v854 = v835
	goto L266
L266:
	;
	switch v849 - int32(3) {
	case 0:
		goto L270
	case 1:
		v861 = v831
		goto L268
	case 2:
		goto L269
	default:
		v863 = v831
		v864 = v842
		goto L267
	}
L267:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v848)+12))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v865<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	switch v865 - int32(3) {
	case 0:
		goto L274
	case 1:
		v876 = v863
		goto L272
	case 2:
		goto L273
	default:
		v878 = v863
		v879 = v864
		goto L271
	}
L268:
	;
	v863 = v861
	v864 = int32(1)
	goto L267
L269:
	;
	v861 = v831 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L268
L270:
	;
	v863 = v831 | int32(_a_F_compute_new_xmax_infomask_6)
	v864 = v842
	goto L267
L271:
	;
	if base.Ui32(v854) < base.Ui32(v868) {
		goto L275
	} else {
		goto L276
	}
L272:
	;
	v878 = v876
	v879 = int32(1)
	goto L271
L273:
	;
	v876 = v863 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L272
L274:
	;
	v878 = v863 | int32(_a_F_compute_new_xmax_infomask_6)
	v879 = v864
	goto L271
L275:
	;
	v880 = v868
	goto L277
L276:
	;
	v880 = v854
	goto L277
L277:
	;
	v881 = int32(2)
	v882 = v834 + v881
	v884 = v830 + v881
	if v884 != v813&int32(2147483646) {
		v830 = v884
		v831 = v878
		v834 = v882
		v835 = v880
		v842 = v879
		goto L262
	} else {
		goto L278
	}
L278:
	;
	goto L263
L279:
	;
	v889 = v878
	v892 = v882
	v893 = v880
	v900 = v879
	goto L258
L280:
	;
	v912 = v910
	goto L282
L281:
	;
	v912 = v893
	goto L282
L282:
	;
	switch v907 - int32(3) {
	case 0:
		goto L285
	case 1:
		v919 = v889
		goto L283
	case 2:
		goto L284
	default:
		v922 = v889
		v926 = v912
		v933 = v900
		goto L257
	}
L283:
	;
	v922 = v919
	v926 = v912
	v933 = int32(1)
	goto L257
L284:
	;
	v919 = v889 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L283
L285:
	;
	v922 = v889 | int32(_a_F_compute_new_xmax_infomask_6)
	v926 = v912
	v933 = v900
	goto L257
L286:
	;
	if v926&int32(-2) == int32(2) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	if v933 != 0 {
		v1050 = v922
		v1052 = v808
		v1054 = int32(_a_F_compute_new_xmax_infomask_8)
		goto L2
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	if v926 != 0 {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	v1050 = v922
	v1052 = v808
	v1054 = int32(_a_F_compute_new_xmax_infomask_9)
	goto L2
L291:
	;
	v948 = int32(_a_F_compute_new_xmax_infomask_3)
	goto L293
L292:
	;
	v948 = int32(_a_F_compute_new_xmax_infomask_10)
	goto L293
L293:
	;
	if v926 == int32(1) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v951 = int32(_a_F_compute_new_xmax_infomask_11)
	goto L296
L295:
	;
	v951 = v948
	goto L296
L296:
	;
	if v933 != 0 {
		v1050 = v922
		v1052 = v808
		v1054 = v951
		goto L2
	} else {
		goto L297
	}
L297:
	;
	v1050 = v922
	v1052 = v808
	v1054 = v951 | int32(128)
	goto L2
L298:
	;
	if l5 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v960 = int32(_a_F_compute_new_xmax_infomask_15)
	goto L301
L300:
	;
	v960 = int32(_a_F_compute_new_xmax_infomask_16)
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = l4
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_17), v19-int32(-64))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L13
	} else {
		goto L302
	}
L302:
	;
	goto L1
L303:
	;
	if l5 != 0 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v974 = int32(_a_F_compute_new_xmax_infomask_15)
	goto L306
L305:
	;
	v974 = int32(_a_F_compute_new_xmax_infomask_16)
	goto L306
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v974
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = l4
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_17), v19+int32(48))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L13
	} else {
		goto L307
	}
L307:
	;
	goto L1
L308:
	;
	if l5 != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v988 = int32(_a_F_compute_new_xmax_infomask_15)
	goto L311
L310:
	;
	v988 = int32(_a_F_compute_new_xmax_infomask_16)
	goto L311
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v988
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l4
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_17), v19)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L13
	} else {
		goto L312
	}
L312:
	;
	goto L1
L313:
	;
	if l5 != 0 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1000 = int32(_a_F_compute_new_xmax_infomask_15)
	goto L316
L315:
	;
	v1000 = int32(_a_F_compute_new_xmax_infomask_16)
	goto L316
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = l4
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_17), v19+int32(32))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L13
	} else {
		goto L317
	}
L317:
	;
	goto L1
L318:
	;
	v1024 = int32(0)
	if v1012 == int32(3) {
		goto L321
	} else {
		goto L322
	}
L319:
	;
	goto L320
L320:
	;
	v1030 = int32(0)
	switch v1012 {
	case 0:
		v1050 = v1030
		v1052 = l3
		v1054 = int32(144)
		goto L2
	case 1:
		goto L327
	case 2:
		goto L326
	case 3:
		goto L325
	default:
		goto L324
	}
L321:
	;
	v1029 = int32(_a_F_compute_new_xmax_infomask_6)
	goto L323
L322:
	;
	v1029 = v1024
	goto L323
L323:
	;
	v1050 = v1029
	v1052 = l3
	v1054 = v1024
	goto L2
L324:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L13
	} else {
		goto L328
	}
L325:
	;
	v1050 = int32(_a_F_compute_new_xmax_infomask_6)
	v1052 = l3
	v1054 = int32(192)
	goto L2
L326:
	;
	v1050 = v1030
	v1052 = l3
	v1054 = int32(192)
	goto L2
L327:
	;
	v1050 = v1030
	v1052 = l3
	v1054 = int32(208)
	goto L2
L328:
	;
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_18), int32(0))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L13
	} else {
		goto L329
	}
L329:
	;
	F_errfinish(m, int32(_a_F_compute_new_xmax_infomask_0), int32(_a_F_compute_new_xmax_infomask_19), int32(_a_F_compute_new_xmax_infomask_14))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L13
	} else {
		goto L330
	}
L330:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L331:
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
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
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
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
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
	v193 = m.ExcPending
	if v193 != 0 {
		goto L12
	} else {
		goto L50
	}
L2:
	;
	m.G0 = v12 + int32(32)
	return v180
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
		v180 = v4
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
	v180 = v39
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
	v105 = v46
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
	v105 = v53
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
	v120 = int32(1)
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
	if v120 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v147 = v117
	v148 = v120
	goto L37
L37:
	;
	v150 = v115 + int32(1)
	if v150 < base.I32_extend16_s(v147) {
		v115 = v150
		v117 = v147
		v120 = v148
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
	v141 = F_OutputFunctionCall(m, v105+v115*int32(28), v129)
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
		goto L46
	} else {
		goto L47
	}
L46:
	;
	base.MemoryCopy(m, v166+int32(4), v162, v163)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	F_pfree(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	v180 = v166
	goto L2
L50:
	;
	F_errmsg_internal(m, int32(_a_F_concat_internal_0), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_concat_internal_1), int32(_a_F_concat_internal_2), int32(_a_F_concat_internal_3))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == v3 {
		v96 = v3
		m.G0 = v9 + int32(16)
		return v96
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v13 != int32(20) {
			v91 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return int32(0)
			} else {
				v96 = v91
				m.G0 = v9 + int32(16)
				return v96
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			if v18 == int32(0) {
				v91 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					v96 = v91
					m.G0 = v9 + int32(16)
					return v96
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				if v21 != int32(7) {
					v91 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						v96 = v91
						m.G0 = v9 + int32(16)
						return v96
					}
				} else {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
					if v24 != 0 {
						v91 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							v96 = v91
							m.G0 = v9 + int32(16)
							return v96
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
									v91 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										v96 = v91
										m.G0 = v9 + int32(16)
										return v96
									}
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
									if v39 != v40 {
										v91 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											v96 = v91
											m.G0 = v9 + int32(16)
											return v96
										}
									} else {
										v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
										v46 = F_ArrayGetNItemsSafe(m, v43, v42+int32(16))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int32(0)
										} else {
											if v46 < int32(9) {
												v96 = v3
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												v81 = int32(12)
												v83 = v51
												*(*int32)(unsafe.Add(mBase, uint32(l0+v81))) = v83
												v96 = v3
											}
											m.G0 = v9 + int32(16)
											return v96
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
									v91 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										v96 = v91
										m.G0 = v9 + int32(16)
										return v96
									}
								} else {
									v60 = F_get_op_hash_functions(m, v52, v9+int32(12), v9+int32(8))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										if v60 == int32(0) {
											v91 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int32(0)
											} else {
												v96 = v91
												m.G0 = v9 + int32(16)
												return v96
											}
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
											if v64 != v65 {
												v91 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return int32(0)
												} else {
													v96 = v91
													m.G0 = v9 + int32(16)
													return v96
												}
											} else {
												v67 = int32(16)
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
												v72 = F_ArrayGetNItemsSafe(m, v69, v68+v67)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													if v72 < int32(9) {
														v96 = v3
														m.G0 = v9 + int32(16)
														return v96
													} else {
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v76
														v78 = F_get_opcode(m, v52)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															v81 = v67
															v83 = v78
															*(*int32)(unsafe.Add(mBase, uint32(l0+v81))) = v83
															v96 = v3
															m.G0 = v9 + int32(16)
															return v96
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == int32(0) {
		v42 = int32(0)
		m.G0 = v7 + int32(16)
		return v42
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = v12 - int32(8)
		if v14 != 0 {
			if v14 == int32(14) {
				v42 = l0
				m.G0 = v7 + int32(16)
				return v42
			} else {
				v40 = F_expression_tree_mutator_impl(m, l0, int32(845), l1)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = v40
					m.G0 = v7 + int32(16)
					return v42
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v17 != int32(2) {
				v40 = F_expression_tree_mutator_impl(m, l0, int32(845), l1)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = v40
					m.G0 = v7 + int32(16)
					return v42
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v20 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v54
						F_errmsg_internal(m, int32(_a_F_convert_testexpr_mutator_0), v7)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_convert_testexpr_mutator_1), int32(669), int32(_a_F_convert_testexpr_mutator_2))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v23 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v54
							F_errmsg_internal(m, int32(_a_F_convert_testexpr_mutator_0), v7)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_convert_testexpr_mutator_1), int32(669), int32(_a_F_convert_testexpr_mutator_2))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
						if v26 < v20 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v54
								F_errmsg_internal(m, int32(_a_F_convert_testexpr_mutator_0), v7)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_convert_testexpr_mutator_1), int32(669), int32(_a_F_convert_testexpr_mutator_2))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+v20<<(uint(int32(2))%32)-int32(4))))
							v35 = F_copyObjectImpl(m, v34)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v42 = v35
								m.G0 = v7 + int32(16)
								return v42
							}
						}
					}
				}
			}
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
				F_yy_fatal_error_2(m, int32(_a_F_core_yyensure_buffer_stack_0))
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
					F_yy_fatal_error_2(m, int32(_a_F_core_yyensure_buffer_stack_0))
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
					*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v179 float64
	_ = v179
	var v180 float64
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 float64
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 float64
	_ = v192
	var v197 float64
	_ = v197
	var v200 float64
	_ = v200
	var v202 float64
	_ = v202
	var v205 float64
	_ = v205
	var v208 int64
	_ = v208
	var v209 float64
	_ = v209
	var v216 float64
	_ = v216
	var v218 float64
	_ = v218
	var v222 int32
	_ = v222
	var v223 float64
	_ = v223
	var v225 float64
	_ = v225
	var v226 int32
	_ = v226
	var v229 float64
	_ = v229
	var v233 float64
	_ = v233
	var v235 float64
	_ = v235
	var v236 float64
	_ = v236
	var v238 float64
	_ = v238
	var v239 float64
	_ = v239
	var v243 float64
	_ = v243
	var v246 float64
	_ = v246
	var v250 float64
	_ = v250
	var v256 float64
	_ = v256
	var v257 float64
	_ = v257
	var v261 float64
	_ = v261
	var v265 float64
	_ = v265
	var v273 float64
	_ = v273
	var v276 float64
	_ = v276
	var v280 int32
	_ = v280
	var v285 float64
	_ = v285
	var v286 float64
	_ = v286
	var v288 float64
	_ = v288
	var v293 int32
	_ = v293
	var v296 float64
	_ = v296
	var v297 float64
	_ = v297
	var v298 float64
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 float64
	_ = v304
	var v305 float64
	_ = v305
	var v308 float64
	_ = v308
	var v309 float64
	_ = v309
	var v310 float64
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 float64
	_ = v317
	var v319 int32
	_ = v319
	var v325 float64
	_ = v325
	var v329 float64
	_ = v329
	var v332 float64
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 float64
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 float64
	_ = v345
	var v346 float64
	_ = v346
	var v350 float64
	_ = v350
	var v354 float64
	_ = v354
	var v357 float64
	_ = v357
	var v360 float64
	_ = v360
	var v361 float64
	_ = v361
	var v363 float64
	_ = v363
	var v365 float64
	_ = v365
	var v367 float64
	_ = v367
	var v370 float64
	_ = v370
	var v372 float64
	_ = v372
	var v374 float64
	_ = v374
	var v375 int32
	_ = v375
	var v377 float64
	_ = v377
	var v386 float64
	_ = v386
	var v390 float64
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v415 float64
	_ = v415
	var v418 float64
	_ = v418
	var v419 float64
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v426 float64
	_ = v426
	var v428 float64
	_ = v428
	var v431 float64
	_ = v431
	var v433 float64
	_ = v433
	var v435 float64
	_ = v435
	var v437 int32
	_ = v437
	var v438 float64
	_ = v438
	var v439 float64
	_ = v439
	var v443 float64
	_ = v443
	var v447 float64
	_ = v447
	var v450 float64
	_ = v450
	var v453 float64
	_ = v453
	var v455 float64
	_ = v455
	var v456 float64
	_ = v456
	var v458 float64
	_ = v458
	var v460 float64
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 float64
	_ = v464
	var v473 float64
	_ = v473
	var v477 float64
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v493 int32
	_ = v493
	var v502 float64
	_ = v502
	var v506 float64
	_ = v506
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v554 int32
	_ = v554
	var v555 float64
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 float64
	_ = v657
	var v658 float64
	_ = v658
	var v664 int32
	_ = v664
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v682 int32
	_ = v682
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v704 float64
	_ = v704
	var v708 float64
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 float64
	_ = v714
	var v718 float64
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 float64
	_ = v724
	var v728 float64
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 float64
	_ = v734
	var v738 float64
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v799 float64
	_ = v799
	var v803 float64
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v916 int32
	_ = v916
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v937 float64
	_ = v937
	var v941 float64
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 float64
	_ = v947
	var v951 float64
	_ = v951
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 float64
	_ = v957
	var v961 float64
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v967 float64
	_ = v967
	var v971 float64
	_ = v971
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1029 int32
	_ = v1029
	var v1032 float64
	_ = v1032
	var v1036 float64
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1047 int32
	_ = v1047
	var v1068 float64
	_ = v1068
	var v1072 float64
	_ = v1072
	var v1073 float64
	_ = v1073
	var v1093 float64
	_ = v1093
	var v1094 float64
	_ = v1094
	var v1100 float64
	_ = v1100
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
	v1100 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[0]))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(base.F64_mul(base.F64_mul(v1100, float64(0.5)), v1094), v1093)
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
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v317 = base.F64_convert_i32_s(v316)
	v319 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_append[1])))
	if v319 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L8:
	;
	v97 = v2
	goto L17
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
		v1093 = v19
		v1094 = v19
		goto L4
	} else {
		goto L13
	}
L12:
	;
	v1093 = v19
	v1094 = v19
	goto L4
L13:
	;
	v57 = v2
	v59 = v2
	v72 = v19
	v73 = v19
	goto L14
L14:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v59<<(uint(int32(2))%32))))
	v83 = *(*float64)(unsafe.Add(mBase, uint32(v82)+32))
	v84 = base.F64_add(v83, v73)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	v87 = v57 + v86
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v87
	v89 = *(*float64)(unsafe.Add(mBase, uint32(v82)+56))
	v90 = base.F64_add(v89, v72)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v90
	v93 = v59 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v93 < v94 {
		v57 = v87
		v59 = v93
		v72 = v90
		v73 = v84
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v1093 = v90
	v1094 = v84
	goto L4
L16:
	;
	goto L15
L17:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120+v97<<(uint(int32(2))%32))))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+64))
	if v43 == v125 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v1093 = v310
	v1094 = v298
	goto L4
L19:
	;
	v297 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v298 = base.F64_add(v296, v297)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v298
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v293)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v300 + v301
	v304 = *(*float64)(unsafe.Add(mBase, uint32(v293)+48))
	v305 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = base.F64_add(v304, v305)
	v308 = *(*float64)(unsafe.Add(mBase, uint32(v293)+56))
	v309 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
	v310 = base.F64_add(v308, v309)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v310
	v313 = v97 + int32(1)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v313 < v314 {
		v97 = v313
		goto L17
	} else {
		goto L61
	}
L20:
	;
	if v178 != 0 {
		goto L38
	} else {
		goto L39
	}
L21:
	;
	v178 = int32(1)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v134 = int32(0)
	goto L25
L24:
	;
	v178 = v170
	goto L20
L25:
	;
	v138 = int32(0)
	if v43 == v138 {
		v148 = v138
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v170 = int32(0)
	goto L24
L27:
	;
	if v125 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v142 <= v134 {
		v148 = int32(0)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v148 = v144 + v134<<(uint(int32(2))%32)
	goto L27
L30:
	;
	v154 = base.B2i32(v148 == int32(0))
	if v148 == int32(0) {
		v170 = v154
		goto L24
	} else {
		goto L35
	}
L31:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v134 < v149 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v178 = base.B2i32(v148 == int32(0))
	goto L20
L34:
	;
	goto L33
L35:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	if v157 == int32(0) {
		v170 = v154
		goto L24
	} else {
		goto L36
	}
L36:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v134<<(uint(int32(2))%32)+v157)))
	if v164 == v166 {
		v134 = v134 + int32(1)
		goto L25
	} else {
		goto L37
	}
L37:
	;
	goto L26
L38:
	;
	v179 = *(*float64)(unsafe.Add(mBase, uint32(v124)+32))
	v293 = v124
	v296 = v179
	goto L19
L39:
	;
	goto L40
L40:
	;
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v124)+56))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v124)+40))
	v183 = v27 + int32(88)
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v124)+32))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+32))
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_cost_append[2]))
	v192 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
	v197 = float64(2)
	if base.F64_lt(v186, v197) != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_append[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v181 + (v280 ^ int32(1))
	v285 = *(*float64)(unsafe.Add(mBase, uint32(v27)+88))
	v286 = base.F64_add(v180, v285)
	*(*float64)(unsafe.Add(mBase, uint32(v27)+56)) = v286
	v288 = *(*float64)(unsafe.Add(mBase, uint32(v27)+80))
	*(*float64)(unsafe.Add(mBase, uint32(v27)+64)) = base.F64_add(v286, v288)
	v293 = v27 + int32(8)
	v296 = v186
	goto L19
L42:
	;
	v200 = v197
	goto L44
L43:
	;
	v200 = v186
	goto L44
L44:
	;
	v202 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[4]))
	v205 = base.F64_mul(v200, base.F64_add(base.F64_add(v202, v202), float64(0)))
	v208 = base.I64_extend_i32_s(v191) << (uint(int64(10)) % 64)
	v209 = base.F64_convert_i64_s(v208)
	v216 = base.F64_convert_i32_u((v188+int32(7))&int32(-8) + int32(24))
	v218 = base.F64_mul(v186, v216)
	v222 = base.F64_lt(v192, v200) & base.F64_gt(v192, float64(0))
	if v222 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v183))) = v273
	v276 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[4]))
	*(*float64)(unsafe.Add(mBase, uint32(v27+int32(80)))) = base.F64_mul(v200, v276)
	goto L41
L46:
	;
	v223 = base.F64_mul(v192, v216)
	goto L48
L47:
	;
	v223 = v218
	goto L48
L48:
	;
	if base.F64_lt(v209, v223) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v225 = F_log(m, v200)
	mBase = m.M
	v226 = F_tuplesort_merge_order(m, v208)
	mBase = m.M
	v229 = base.F64_mul(base.F64_div(v225, float64(0.693147180559945)), v205)
	*(*float64)(unsafe.Add(mBase, uint32(v183))) = v229
	v233 = base.F64_ceil(base.F64_mul(v218, float64(0.0001220703125)))
	v235 = base.F64_div(v218, v209)
	v236 = base.F64_convert_i32_s(v226)
	if base.F64_gt(v235, v236) != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	if v222 != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v238 = F_log(m, v235)
	mBase = m.M
	v239 = F_log(m, v236)
	mBase = m.M
	v243 = base.F64_ceil(base.F64_div(v238, v239))
	goto L54
L53:
	;
	v243 = float64(1)
	goto L54
L54:
	;
	v246 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[5]))
	v250 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[6]))
	v273 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v233, v233), v243), base.F64_add(base.F64_mul(v246, float64(0.75)), base.F64_mul(v250, float64(0.25)))), v229)
	goto L45
L55:
	;
	v256 = v192
	goto L57
L56:
	;
	v256 = v200
	goto L57
L57:
	;
	v257 = base.F64_add(v256, v256)
	if base.F64_gt(v200, v257)|base.F64_gt(v218, v209) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v261 = F_log(m, v257)
	mBase = m.M
	v273 = base.F64_mul(base.F64_div(v261, float64(0.693147180559945)), v205)
	goto L45
L59:
	;
	goto L60
L60:
	;
	v265 = F_log(m, v200)
	mBase = m.M
	v273 = base.F64_mul(base.F64_div(v265, float64(0.693147180559945)), v205)
	goto L45
L61:
	;
	goto L18
L62:
	;
	v325 = base.F64_add(base.F64_mul(v317, float64(-0.3)), float64(1))
	if base.F64_gt(v325, float64(0)) != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v332 = v317
	goto L64
L64:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v333 <= int32(0) {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v329 = v325
	goto L67
L66:
	;
	v329 = math.Float64frombits(uint64(0x8000000000000000))
	goto L67
L67:
	;
	v332 = base.F64_add(v329, v317)
	goto L64
L68:
	;
	if v493 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L69:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v493 = v336
	v502 = v19
	v506 = v19
	goto L68
L70:
	;
	goto L71
L71:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	v339 = *(*float64)(unsafe.Add(mBase, uint32(v338)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v339
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v341 <= int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v338)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v375
	v377 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v372)&int64(9223372036854775807)))|base.F64_gt(v372, v377) != 0 {
		v390 = v377
		goto L82
	} else {
		goto L83
	}
L73:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v338)+24))
	v345 = base.F64_convert_i32_s(v344)
	v346 = *(*float64)(unsafe.Add(mBase, uint32(v338)+32))
	if v319 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v367 = *(*float64)(unsafe.Add(mBase, uint32(v338)+32))
	v370 = base.F64_add(base.F64_div(v367, v332), float64(0))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v370
	v372 = v370
	v374 = v19
	goto L72
L76:
	;
	v350 = base.F64_add(base.F64_mul(v345, float64(-0.3)), float64(1))
	if base.F64_gt(v350, float64(0)) != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v357 = v345
	goto L78
L78:
	;
	v360 = float64(0)
	v361 = base.F64_add(base.F64_mul(v346, base.F64_div(v357, v332)), v360)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v361
	v363 = *(*float64)(unsafe.Add(mBase, uint32(v338)+56))
	v365 = base.F64_add(v363, v360)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v365
	v372 = v361
	v374 = v365
	goto L72
L79:
	;
	v354 = v350
	goto L81
L80:
	;
	v354 = math.Float64frombits(uint64(0x8000000000000000))
	goto L81
L81:
	;
	v357 = base.F64_add(v354, v345)
	goto L78
L82:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v390
	v392 = int32(1)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v393 <= v392 {
		v493 = v341
		v502 = v390
		v506 = v374
		goto L68
	} else {
		goto L85
	}
L83:
	;
	v386 = float64(1)
	if base.F64_le(v372, v386) != 0 {
		v390 = v386
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v390 = base.F64_nearest(v372)
	goto L82
L85:
	;
	v399 = v375
	v401 = v392
	v415 = v390
	v418 = v339
	v419 = v374
	goto L86
L86:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v420+v401<<(uint(int32(2))%32))))
	if v401 < v316 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v493 = v341
	v502 = v477
	v506 = v460
	goto L68
L88:
	;
	v426 = *(*float64)(unsafe.Add(mBase, uint32(v424)+48))
	if base.F64_gt(v426, v418) != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v431 = v418
	goto L90
L90:
	;
	if v401 < v341 {
		goto L95
	} else {
		goto L96
	}
L91:
	;
	v428 = v418
	goto L93
L92:
	;
	v428 = v426
	goto L93
L93:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v39))) = v428
	v431 = v428
	goto L90
L94:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v424)+40))
	v462 = v399 + v461
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v462
	v464 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v458)&int64(9223372036854775807)))|base.F64_gt(v458, v464) != 0 {
		v477 = v464
		goto L104
	} else {
		goto L105
	}
L95:
	;
	v433 = *(*float64)(unsafe.Add(mBase, uint32(v424)+32))
	v435 = base.F64_add(v415, base.F64_div(v433, v332))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v435
	v458 = v435
	v460 = v419
	goto L94
L96:
	;
	goto L97
L97:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v424)+24))
	v438 = base.F64_convert_i32_s(v437)
	v439 = *(*float64)(unsafe.Add(mBase, uint32(v424)+32))
	if v319 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v443 = base.F64_add(base.F64_mul(v438, float64(-0.3)), float64(1))
	if base.F64_gt(v443, float64(0)) != 0 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v450 = v438
	goto L100
L100:
	;
	v453 = base.F64_add(base.F64_mul(v439, base.F64_div(v450, v332)), v415)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v453
	v455 = *(*float64)(unsafe.Add(mBase, uint32(v424)+56))
	v456 = base.F64_add(v455, v419)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v456
	v458 = v453
	v460 = v456
	goto L94
L101:
	;
	v447 = v443
	goto L103
L102:
	;
	v447 = math.Float64frombits(uint64(0x8000000000000000))
	goto L103
L103:
	;
	v450 = base.F64_add(v447, v438)
	goto L100
L104:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v477
	v480 = v401 + int32(1)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v480 < v481 {
		v399 = v462
		v401 = v480
		v415 = v477
		v418 = v431
		v419 = v460
		goto L86
	} else {
		goto L107
	}
L105:
	;
	v473 = float64(1)
	if base.F64_le(v458, v473) != 0 {
		v477 = v473
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v477 = base.F64_nearest(v458)
	goto L104
L107:
	;
	goto L87
L108:
	;
	v1093 = base.F64_add(v506, float64(0))
	v1094 = v502
	goto L4
L109:
	;
	goto L110
L110:
	;
	if v316 < v493 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v512 = v316
	goto L113
L112:
	;
	v512 = v493
	goto L113
L113:
	;
	v515 = F_palloc(m, v512<<(uint(int32(3))%32))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	return
L115:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v517 <= int32(0) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v1068 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v1072 = *(*float64)(unsafe.Add(mBase, uint32(v515+v1047<<(uint(int32(3))%32))))
	v1073 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
	v1093 = base.F64_add(v1072, v1073)
	v1094 = v1068
	goto L4
L117:
	;
	v892 = int32(3)
	v893 = v512 & v892
	v894 = int32(0)
	if base.Ui32(v892) <= base.Ui32(v512-int32(1)) {
		goto L166
	} else {
		goto L167
	}
L118:
	;
	if v593 != 0 {
		goto L128
	} else {
		goto L129
	}
L119:
	;
	v587 = int32(0)
	if v512 <= v587 {
		v1047 = v587
		goto L116
	} else {
		goto L126
	}
L120:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v512 == int32(0) {
		v593 = v520
		v596 = v517
		goto L118
	} else {
		goto L121
	}
L121:
	;
	v525 = int32(0)
	v527 = v520
	goto L122
L122:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v527+v525<<(uint(int32(2))%32))))
	v555 = *(*float64)(unsafe.Add(mBase, uint32(v554)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v515+v525<<(uint(int32(3))%32)))) = v555
	v558 = v525 + int32(1)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v559 <= v558 {
		goto L119
	} else {
		goto L124
	}
L123:
	;
	v593 = v561
	v596 = v559
	goto L118
L124:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v558 != v512 {
		v525 = v558
		v527 = v561
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	goto L117
L127:
	;
	if int32(0) < v512 {
		goto L117
	} else {
		goto L165
	}
L128:
	;
	v614 = v512
	goto L130
L129:
	;
	v614 = v596
	goto L130
L130:
	;
	if v596 <= v614 {
		goto L127
	} else {
		goto L131
	}
L131:
	;
	v619 = v512 & int32(3)
	v625 = v512 - int32(1)
	v632 = v512
	v637 = v614
	goto L132
L132:
	;
	if v632 == v493 {
		goto L127
	} else {
		goto L134
	}
L133:
	;
	goto L127
L134:
	;
	v651 = v515 + v625<<(uint(int32(3))%32)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v652+v637<<(uint(int32(2))%32))))
	v657 = *(*float64)(unsafe.Add(mBase, uint32(v656)+56))
	v658 = *(*float64)(unsafe.Add(mBase, uint32(v651)))
	*(*float64)(unsafe.Add(mBase, uint32(v651))) = base.F64_add(v657, v658)
	if v512 <= int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v835 = int32(1)
	v838 = v637 + v835
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v838 < v839 {
		v625 = v812
		v632 = v632 + v835
		v637 = v838
		goto L132
	} else {
		goto L164
	}
L136:
	;
	v812 = int32(0)
	goto L135
L137:
	;
	goto L138
L138:
	;
	v664 = int32(0)
	if base.B2i32(base.Ui32(v512) < base.Ui32(int32(4))) == v664 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v672 = v664
	v674 = v664
	v682 = v664
	goto L142
L140:
	;
	v749 = v664
	v754 = v664
	goto L141
L141:
	;
	v773 = v749
	v778 = v754
	v789 = v664
	goto L158
L142:
	;
	v695 = int32(3)
	v696 = v674 | v695
	v698 = v674 | int32(2)
	v700 = v674 | int32(1)
	v704 = *(*float64)(unsafe.Add(mBase, uint32(v515+v674<<(uint(v695)%32))))
	v708 = *(*float64)(unsafe.Add(mBase, uint32(v515+v672<<(uint(v695)%32))))
	if base.F64_lt(v704, v708) != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	if v619 == int32(0) {
		v812 = v740
		goto L135
	} else {
		goto L157
	}
L144:
	;
	v710 = v674
	goto L146
L145:
	;
	v710 = v672
	goto L146
L146:
	;
	v711 = int32(3)
	v714 = *(*float64)(unsafe.Add(mBase, uint32(v515+v700<<(uint(v711)%32))))
	v718 = *(*float64)(unsafe.Add(mBase, uint32(v515+v710<<(uint(v711)%32))))
	if base.F64_lt(v714, v718) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v720 = v700
	goto L149
L148:
	;
	v720 = v710
	goto L149
L149:
	;
	v721 = int32(3)
	v724 = *(*float64)(unsafe.Add(mBase, uint32(v515+v698<<(uint(v721)%32))))
	v728 = *(*float64)(unsafe.Add(mBase, uint32(v515+v720<<(uint(v721)%32))))
	if base.F64_lt(v724, v728) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v730 = v698
	goto L152
L151:
	;
	v730 = v720
	goto L152
L152:
	;
	v731 = int32(3)
	v734 = *(*float64)(unsafe.Add(mBase, uint32(v515+v696<<(uint(v731)%32))))
	v738 = *(*float64)(unsafe.Add(mBase, uint32(v515+v730<<(uint(v731)%32))))
	if base.F64_lt(v734, v738) != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v740 = v696
	goto L155
L154:
	;
	v740 = v730
	goto L155
L155:
	;
	v741 = int32(4)
	v742 = v674 + v741
	v744 = v682 + v741
	if v744 != v512&int32(2147483644) {
		v672 = v740
		v674 = v742
		v682 = v744
		goto L142
	} else {
		goto L156
	}
L156:
	;
	goto L143
L157:
	;
	v749 = v740
	v754 = v742
	goto L141
L158:
	;
	v796 = int32(3)
	v799 = *(*float64)(unsafe.Add(mBase, uint32(v515+v778<<(uint(v796)%32))))
	v803 = *(*float64)(unsafe.Add(mBase, uint32(v515+v773<<(uint(v796)%32))))
	if base.F64_lt(v799, v803) != 0 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v812 = v805
	goto L135
L160:
	;
	v805 = v778
	goto L162
L161:
	;
	v805 = v773
	goto L162
L162:
	;
	v806 = int32(1)
	v809 = v789 + v806
	if v809 != v619 {
		v773 = v805
		v778 = v778 + v806
		v789 = v809
		goto L158
	} else {
		goto L163
	}
L163:
	;
	goto L159
L164:
	;
	goto L133
L165:
	;
	v1047 = int32(0)
	goto L116
L166:
	;
	v905 = v894
	v907 = v894
	v916 = int32(0)
	goto L169
L167:
	;
	v982 = v894
	v984 = v894
	goto L168
L168:
	;
	v1006 = v982
	v1008 = v984
	v1011 = v894
	goto L185
L169:
	;
	v928 = int32(3)
	v929 = v905 | v928
	v931 = v905 | int32(2)
	v933 = v905 | int32(1)
	v937 = *(*float64)(unsafe.Add(mBase, uint32(v515+v905<<(uint(v928)%32))))
	v941 = *(*float64)(unsafe.Add(mBase, uint32(v515+v907<<(uint(v928)%32))))
	if base.F64_gt(v937, v941) != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v893 == int32(0) {
		v1047 = v973
		goto L116
	} else {
		goto L184
	}
L171:
	;
	v943 = v905
	goto L173
L172:
	;
	v943 = v907
	goto L173
L173:
	;
	v944 = int32(3)
	v947 = *(*float64)(unsafe.Add(mBase, uint32(v515+v933<<(uint(v944)%32))))
	v951 = *(*float64)(unsafe.Add(mBase, uint32(v515+v943<<(uint(v944)%32))))
	if base.F64_gt(v947, v951) != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v953 = v933
	goto L176
L175:
	;
	v953 = v943
	goto L176
L176:
	;
	v954 = int32(3)
	v957 = *(*float64)(unsafe.Add(mBase, uint32(v515+v931<<(uint(v954)%32))))
	v961 = *(*float64)(unsafe.Add(mBase, uint32(v515+v953<<(uint(v954)%32))))
	if base.F64_gt(v957, v961) != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v963 = v931
	goto L179
L178:
	;
	v963 = v953
	goto L179
L179:
	;
	v964 = int32(3)
	v967 = *(*float64)(unsafe.Add(mBase, uint32(v515+v929<<(uint(v964)%32))))
	v971 = *(*float64)(unsafe.Add(mBase, uint32(v515+v963<<(uint(v964)%32))))
	if base.F64_gt(v967, v971) != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v973 = v929
	goto L182
L181:
	;
	v973 = v963
	goto L182
L182:
	;
	v974 = int32(4)
	v975 = v905 + v974
	v977 = v916 + v974
	if v977 != v512&int32(-4) {
		v905 = v975
		v907 = v973
		v916 = v977
		goto L169
	} else {
		goto L183
	}
L183:
	;
	goto L170
L184:
	;
	v982 = v975
	v984 = v973
	goto L168
L185:
	;
	v1029 = int32(3)
	v1032 = *(*float64)(unsafe.Add(mBase, uint32(v515+v1006<<(uint(v1029)%32))))
	v1036 = *(*float64)(unsafe.Add(mBase, uint32(v515+v1008<<(uint(v1029)%32))))
	if base.F64_gt(v1032, v1036) != 0 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v1047 = v1038
	goto L116
L187:
	;
	v1038 = v1006
	goto L189
L188:
	;
	v1038 = v1008
	goto L189
L189:
	;
	v1039 = int32(1)
	v1042 = v1011 + v1039
	if v1042 != v893 {
		v1006 = v1006 + v1039
		v1008 = v1038
		v1011 = v1042
		goto L185
	} else {
		goto L190
	}
L190:
	;
	goto L186
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
					F_errmsg(m, int32(_a_F_countVariablesFromJsonb_0), int32(0))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						F_errdetail(m, int32(_a_F_countVariablesFromJsonb_1), int32(0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_countVariablesFromJsonb_2), int32(3211), int32(_a_F_countVariablesFromJsonb_3))
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
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
	var v238 int64
	_ = v238
	var v254 int32
	_ = v254
	var v265 int32
	_ = v265
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	v20 = F_palloc(m, int32(_a_F_createPostingTree_0))
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
	F_PageInit(m, v20, int32(_a_F_createPostingTree_0), int32(8))
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
		v89 = v6
		v92 = v6
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v96 = v92 + int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+12)) = uint16(v96)
	v98 = F_GinNewBuffer(m, l0)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L18
	}
L5:
	;
	v47 = v6
	v48 = v20 + int32(32)
	v49 = v6
	goto L6
L6:
	;
	v62 = F_ginCompressPostingList(m, l1+v49*int32(6), l2-v49, int32(384), v17+int32(28))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v89 = v79
	v92 = v71
	goto L4
L8:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
	v70 = (v64+int32(1))&int32(_a_F_createPostingTree_1) + int32(8)
	v71 = v70 + v47
	if base.Ui32(int32(_a_F_createPostingTree_2)) <= base.Ui32(v71) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v89 = v49
	v92 = v47
	goto L4
L10:
	;
	goto L11
L11:
	;
	if v70 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	base.MemoryCopy(m, v48, v62, v70)
	goto L14
L13:
	;
	goto L14
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	F_pfree(m, v62)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v79 = v49 + v75
	if base.Ui32(v79) < base.Ui32(l2) {
		v47 = v71
		v48 = v48 + v70
		v49 = v79
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L7
L17:
	;
	if v98 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	if v98 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[0]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v103+(v98^int32(-1))<<(uint(int32(2))%32))))
	v117 = v109
	goto L17
L20:
	;
	goto L21
L21:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[1]))
	v117 = v111 + v98<<(uint(int32(13))%32) + int32(-8192)
	goto L17
L22:
	;
	if l4 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[2]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121+(v98^int32(-1))<<(uint(int32(6))%32))+16))
	v136 = v127
	goto L22
L24:
	;
	goto L25
L25:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[3]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129+v98<<(uint(int32(6))%32)+int32(-64))+16))
	v136 = v135
	goto L22
L26:
	;
	F_PredicateLockPageSplit(m, l0, v155, v136)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L30
	}
L27:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[2]))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140+(l4^int32(-1))<<(uint(int32(6))%32))+16))
	v155 = v146
	goto L26
L28:
	;
	goto L29
L29:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[3]))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v148+l4<<(uint(int32(6))%32)+int32(-64))+16))
	v155 = v154
	goto L26
L30:
	;
	v158 = int32(_a_F_createPostingTree_3)
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[4])) = v160 + int32(1)
	F_PageRestoreTempPage(m, v20, v117)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_MarkBufferDirty(m, v98)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+118)))
	if v169 != int32(112) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v226 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L53
	}
L34:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v220 + int32(1)
	goto L33
L35:
	;
	F_UnlockReleaseBuffer(m, v98)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L52
	}
L36:
	;
	F_UnlockReleaseBuffer(m, v98)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L50
	}
L37:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[5]))
	if v173 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v92
	F_XLogBeginInsert(m)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L45
	}
L39:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v176|l3 != 0 {
		goto L36
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if l3 != 0 {
		goto L35
	} else {
		goto L44
	}
L42:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v178 == int32(0) {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L36
L44:
	;
	goto L38
L45:
	;
	F_XLogRegisterData(m, v17+int32(28), int32(4))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_XLogRegisterData(m, v117+int32(32), v92)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_XLogRegisterBuffer(m, int32(0), v98, int32(6))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v199 = F_XLogInsert(m, int32(13), int32(16))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = base.I64_rotr(v199, int64(32))
	goto L36
L50:
	;
	v206 = int32(_a_F_createPostingTree_3)
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[4])) = v208 - int32(1)
	if l3 != 0 {
		goto L34
	} else {
		goto L51
	}
L51:
	;
	goto L33
L52:
	;
	v214 = int32(_a_F_createPostingTree_3)
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_createPostingTree[4])) = v216 - int32(1)
	goto L34
L53:
	;
	if v226 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v89
	F_errmsg_internal(m, int32(_a_F_createPostingTree_4), v17)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if base.Ui32(v89) < base.Ui32(l2) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	F_errfinish(m, int32(_a_F_createPostingTree_5), int32(1865), int32(_a_F_createPostingTree_6))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v238 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v238
	*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = v238
	*(*int64)(unsafe.Add(mBase, uint32(v17)+72)) = v238
	*(*int64)(unsafe.Add(mBase, uint32(v17)+88)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = int32(35)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = int32(37)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = int32(38)
	v254 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = int32(39)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = int32(42)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = l0
	v265 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+64)) = uint8(v265)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+81)) = uint8(base.B2i32(l3 != v254))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = l2 - v89
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l1 + v89*int32(6)
	v280 = v17 + int32(90)
	v283 = v254
	goto L62
L60:
	;
	goto L61
L61:
	;
	m.G0 = v17 + int32(96)
	return v136
L62:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v299 = v296 + v283*int32(6)
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v299)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v280)+4)) = uint16(v300)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = v302
	v305 = v17 + int32(28)
	v308 = F_ginFindLeafPage(m, v305, int32(0), int32(1))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	goto L61
L64:
	;
	F_ginInsertValue(m, v305, v308, v17+int32(16), l3)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if base.Ui32(v314) < base.Ui32(v315) {
		v283 = v314
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
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
		v18 = *(*int64)(unsafe.Add(mBase, _c_F_create_ctas_internal[0]))
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
				v56 = F_transformRelOptions(m, v49, v50, int32(_a_F_create_ctas_internal_0), v10+int32(8), int32(1), v49)
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
									v71 = F_pstrdup(m, int32(_a_F_create_ctas_internal_1))
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
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	v6 = F_palloc0(m, int32(80))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(1546188226853)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+20)) = uint8(v16)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v15
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if v19 == int32(1) {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
			v24 = v22
		} else {
			v24 = int32(0)
		}
		v26 = v24 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+21)) = uint8(v26)
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v28
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+72)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v30
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		v34 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
		v35 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
		v36 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
		v42 = *(*int32)(unsafe.Add(mBase, _c_F_create_material_path[0]))
		*(*float64)(unsafe.Add(mBase, uint32(v6)+32)) = v36
		v45 = *(*float64)(unsafe.Add(mBase, _c_F_create_material_path[1]))
		v49 = base.F64_add(base.F64_mul(base.F64_add(v45, v45), v36), base.F64_sub(v35, v34))
		v57 = base.F64_mul(v36, base.F64_convert_i32_u((v38+int32(7))&int32(-8)+int32(24)))
		if base.F64_gt(v57, base.F64_convert_i32_u(v42<<(uint(int32(10))%32))) != 0 {
			v63 = *(*float64)(unsafe.Add(mBase, _c_F_create_material_path[2]))
			v69 = base.F64_add(base.F64_mul(v63, base.F64_ceil(base.F64_mul(v57, float64(0.0001220703125)))), v49)
		} else {
			v69 = v49
		}
		v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_material_path[3])))
		*(*float64)(unsafe.Add(mBase, uint32(v6)+56)) = base.F64_add(v34, v69)
		*(*float64)(unsafe.Add(mBase, uint32(v6)+48)) = v34
		*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v33 + (v71 ^ int32(1))
		return v6
	}
}
func F_create_ordinary_grouping_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v29 float64
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v105 int32
	_ = v105
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v187 int32
	_ = v187
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v254 int32
	_ = v254
	var v275 int32
	_ = v275
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v326 int32
	_ = v326
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
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
	var v511 int32
	_ = v511
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 float64
	_ = v667
	var v668 int32
	_ = v668
	var v669 float64
	_ = v669
	var v670 int32
	_ = v670
	var v671 float64
	_ = v671
	var v672 float64
	_ = v672
	var v673 int32
	_ = v673
	var v674 float64
	_ = v674
	var v675 int32
	_ = v675
	var v676 float64
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v709 int32
	_ = v709
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v743 int32
	_ = v743
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v837 int32
	_ = v837
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v1003 int32
	_ = v1003
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1037 int32
	_ = v1037
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1131 int32
	_ = v1131
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1264 int32
	_ = v1264
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1304 int32
	_ = v1304
	var v1314 int32
	_ = v1314
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1352 int32
	_ = v1352
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1395 int32
	_ = v1395
	var v1409 int32
	_ = v1409
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1422 int32
	_ = v1422
	var v1432 int32
	_ = v1432
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1454 int32
	_ = v1454
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1499 int32
	_ = v1499
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1622 int32
	_ = v1622
	var v1632 int32
	_ = v1632
	var v1637 int32
	_ = v1637
	var v1645 int32
	_ = v1645
	var v1655 int32
	_ = v1655
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1701 int32
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1720 float64
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1722 float64
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1763 int32
	_ = v1763
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
	var v1794 int32
	_ = v1794
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1840 int32
	_ = v1840
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1853 int32
	_ = v1853
	var v1860 int32
	_ = v1860
	var v1864 int32
	_ = v1864
	var v1870 int32
	_ = v1870
	var v1874 int32
	_ = v1874
	var v1878 int32
	_ = v1878
	var v1882 int32
	_ = v1882
	var v1888 int32
	_ = v1888
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1912 int32
	_ = v1912
	var v1916 int32
	_ = v1916
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1927 int32
	_ = v1927
	var v1930 int32
	_ = v1930
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1939 int32
	_ = v1939
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
	var v1949 int32
	_ = v1949
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v2030 int32
	_ = v2030
	var v2033 int32
	_ = v2033
	var v2059 int32
	_ = v2059
	var v2067 int32
	_ = v2067
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2090 int32
	_ = v2090
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2150 int32
	_ = v2150
	var v2157 int32
	_ = v2157
	var v2161 int32
	_ = v2161
	var v2167 int32
	_ = v2167
	var v2171 int32
	_ = v2171
	var v2175 int32
	_ = v2175
	var v2179 int32
	_ = v2179
	var v2185 int32
	_ = v2185
	var v2197 int32
	_ = v2197
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2204 int32
	_ = v2204
	var v2209 int32
	_ = v2209
	var v2213 int32
	_ = v2213
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2227 int32
	_ = v2227
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2318 int32
	_ = v2318
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2351 int32
	_ = v2351
	var v2354 int32
	_ = v2354
	var v2359 int32
	_ = v2359
	var v2362 int32
	_ = v2362
	var v2367 int32
	_ = v2367
	var v2404 int32
	_ = v2404
	var v2407 int32
	_ = v2407
	var v2411 int32
	_ = v2411
	var v2415 int32
	_ = v2415
	var v2420 int32
	_ = v2420
	v8 = int32(0)
	v29 = float64(0)
	v31 = m.G0
	v33 = v31 - int32(128)
	m.G0 = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v36 = int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l5)+100))
	if v37 == v8 {
		v326 = v8
		v337 = v36
		v338 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v346&int32(4) == int32(0) {
		v1314 = v8
		goto L50
	} else {
		goto L51
	}
L2:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	if v40 == int32(0) {
		v326 = v8
		v337 = v36
		v338 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+240))
	if v43 == int32(0) {
		v326 = v8
		v337 = v36
		v338 = v8
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	if v46 <= int32(0) {
		v326 = v8
		v337 = v36
		v338 = v8
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	if v49 == int32(0) {
		v326 = v8
		v337 = v36
		v338 = v8
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v52 = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v54 == v52 {
		v75 = v52
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v75 != 0 {
		v326 = v8
		v337 = v36
		v338 = v8
		goto L1
	} else {
		goto L17
	}
L8:
	;
	goto L7
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v58 = v57
	goto L10
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if base.Ui32(int32(2)) <= base.Ui32(v62-int32(301)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v75 = int32(1)
	goto L8
L12:
	;
	if v62 != int32(290) {
		v75 = v52
		goto L8
	} else {
		goto L15
	}
L13:
	;
	v58 = v61 + int32(72)
	goto L10
L14:
	;
	goto L11
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v61)+72))
	if v69 != 0 {
		v75 = v52
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l5)+100))
	if v76 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v309 = v307 & int32(4)
	v326 = int32(base.Ui32(v309) >> (uint(int32(2)) % 32))
	v337 = base.B2i32(v309 == int32(0))
	v338 = int32(base.Ui32(v309) >> (uint(int32(1)) % 32))
	goto L1
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+100))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v82 = F_get_sortgrouplist_exprs(m, v80, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	if v84 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v87 = int32(1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+2)))
	if v89 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v326 = v8
	v337 = int32(0)
	v338 = v87
	goto L1
L24:
	;
	goto L25
L25:
	;
	v105 = v8
	goto L26
L26:
	;
	v124 = v105 << (uint(int32(2)) % 32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124+v125)))
	if v127 == int32(0) {
		goto L18
	} else {
		goto L28
	}
L27:
	;
	goto L18
L28:
	;
	v130 = int32(0)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v131 <= v130 {
		goto L18
	} else {
		goto L29
	}
L29:
	;
	v141 = v130
	v144 = v131
	goto L30
L30:
	;
	if v82 == int32(0) {
		v254 = v144
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L27
L32:
	;
	v275 = v141 + int32(1)
	if v275 < v254 {
		v141 = v275
		v144 = v254
		goto L30
	} else {
		goto L49
	}
L33:
	;
	v166 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v167 <= v166 {
		v254 = v144
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v170+v141<<(uint(int32(2))%32))))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v176+v124)))
	v187 = v166
	goto L36
L35:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v254 = v243
	goto L32
L36:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209+v187<<(uint(int32(2))%32))))
	v214 = F_exprCollation(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L20
	} else {
		goto L38
	}
L37:
	;
	v229 = int32(0)
	if base.B2i32(v178 == v229)|base.B2i32(v214 == v229)|base.B2i32(v178 == v214) == v229 {
		goto L18
	} else {
		goto L47
	}
L38:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	if v216 == int32(27) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v220 = v219
	goto L41
L40:
	;
	v220 = v213
	goto L41
L41:
	;
	v221 = F_equal(m, v220, v174)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L20
	} else {
		goto L42
	}
L42:
	;
	if v221 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v226 = v187 + int32(1)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v227 <= v226 {
		goto L35
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	goto L37
L46:
	;
	v187 = v226
	goto L36
L47:
	;
	v238 = int32(0)
	v241 = v105 + int32(1)
	if v241 == v89 {
		v326 = v238
		v337 = v238
		v338 = v87
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v105 = v241
	goto L26
L49:
	;
	goto L31
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1314
	if v337 != 0 {
		goto L289
	} else {
		goto L290
	}
L51:
	;
	v351 = int32(0)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v352 == v351 {
		v360 = v351
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+26)))
	if v361 != int32(1) {
		v370 = v8
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l5)+100))
	if v356 != int32(2) {
		v360 = int32(0)
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v360 = v359
	goto L52
L55:
	;
	if v326|base.B2i32(v360|v370 != int32(0)) != int32(1) {
		v1314 = v8
		goto L50
	} else {
		goto L58
	}
L56:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v364 == int32(0) {
		v370 = v8
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	v370 = v368
	goto L55
L58:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v380 = F_fetch_upper_rel(m, l0, int32(1), v379)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L20
	} else {
		goto L59
	}
L59:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v380)+26)) = uint8(v382)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v380)+4)) = v384
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l2)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v380)+156)) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l2)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v380)+160)) = v388
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+164)))
	*(*uint8)(unsafe.Add(mBase, uint32(v380)+164)) = uint8(v390)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l2)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v380)+168)) = v392
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l5)+92))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v396 = F_create_empty_pathtarget(m)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L20
	} else {
		goto L60
	}
L60:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	if v398 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v394 != 0 {
		goto L90
	} else {
		goto L91
	}
L62:
	;
	v511 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v402 = int32(0)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	if v403 <= v402 {
		v511 = v402
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v414 = v402
	v417 = int32(0)
	goto L66
L66:
	;
	v438 = v417 << (uint(int32(2)) % 32)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v398)+12))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v438+v439)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v395)+8))
	if v442 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v511 = v497
	goto L61
L68:
	;
	v501 = v417 + int32(1)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	if v501 < v502 {
		v414 = v497
		v417 = v501
		goto L66
	} else {
		goto L89
	}
L69:
	;
	v495 = F_lappend(m, v414, v441)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L20
	} else {
		goto L88
	}
L70:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v438+v442)))
	if v446 == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v449 == int32(0) {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	if v449 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	if v487 == int32(0) {
		goto L69
	} else {
		goto L86
	}
L74:
	;
	goto L73
L75:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v455 <= int32(0) {
		v487 = int32(0)
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v487 = int32(0)
	goto L74
L78:
	;
	v458 = int32(0)
	if v458 < v455 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v461 = v455
	goto L81
L80:
	;
	v461 = v458
	goto L81
L81:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v449)+12))
	v465 = int32(0)
	goto L82
L82:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v462+v465<<(uint(int32(2))%32))))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	if v473 == v446 {
		v487 = v472
		goto L74
	} else {
		goto L84
	}
L83:
	;
	goto L77
L84:
	;
	v476 = v465 + int32(1)
	if v476 != v461 {
		v465 = v476
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	F_add_column_to_pathtarget(m, v396, v441, v446)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L20
	} else {
		goto L87
	}
L87:
	;
	v497 = v414
	goto L68
L88:
	;
	v497 = v495
	goto L68
L89:
	;
	goto L67
L90:
	;
	v534 = F_lappend(m, v511, v394)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L20
	} else {
		goto L93
	}
L91:
	;
	v536 = v511
	goto L92
L92:
	;
	v538 = F_pull_var_clause(m, v536, int32(25))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L20
	} else {
		goto L94
	}
L93:
	;
	v536 = v534
	goto L92
L94:
	;
	F_add_new_columns_to_pathtarget(m, v396, v538)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L20
	} else {
		goto L95
	}
L95:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v542 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v640 = l5 + int32(8)
	F_list_free(m, v538)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L20
	} else {
		goto L109
	}
L97:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	if v545 <= int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v557 = v545
	v559 = int32(0)
	goto L99
L99:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v542)+12))
	v582 = v579 + v559<<(uint(int32(2))%32)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)))
	if v584 == int32(9) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L96
L101:
	;
	v588 = F_palloc0(m, int32(72))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L20
	} else {
		goto L104
	}
L102:
	;
	v604 = v557
	goto L103
L103:
	;
	v607 = v559 + int32(1)
	if v607 < v604 {
		v557 = v604
		v559 = v607
		goto L99
	} else {
		goto L108
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v588))) = int32(9)
	base.MemoryCopy(m, v588, v583, int32(72))
	*(*int32)(unsafe.Add(mBase, uint32(v588)+56)) = int32(6)
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v588)+20))
	if v597 == int32(2281) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v600 = int32(17)
	goto L107
L106:
	;
	v600 = v597
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v588)+8)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v582))) = v588
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	v604 = v603
	goto L103
L108:
	;
	goto L100
L109:
	;
	F_list_free(m, v536)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L20
	} else {
		goto L110
	}
L110:
	;
	v645 = F_set_pathtarget_cost_width(m, l0, v396)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L20
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+28)) = v645
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+4)))
	if v648 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	base.MemoryFill(m, v640, int32(0), int32(80))
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+36)))
	if v654 == int32(1) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L114
L114:
	;
	if v360 != 0 {
		goto L120
	} else {
		goto L121
	}
L115:
	;
	F_get_agg_clause_costs(m, l0, int32(6), v640)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L20
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v665 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+4)) = uint8(v665)
	goto L114
L118:
	;
	F_get_agg_clause_costs(m, l0, int32(9), l5+int32(48))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L20
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v667 = *(*float64)(unsafe.Add(mBase, uint32(v360)+32))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v669 = F_get_number_of_groups(m, l0, v667, l4, v668)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L20
	} else {
		goto L123
	}
L121:
	;
	v671 = v29
	goto L122
L122:
	;
	if v370 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v671 = v669
	goto L122
L124:
	;
	v672 = *(*float64)(unsafe.Add(mBase, uint32(v370)+32))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v674 = F_get_number_of_groups(m, l0, v672, l4, v673)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L20
	} else {
		goto L127
	}
L125:
	;
	v676 = v29
	goto L126
L126:
	;
	v678 = v346 & int32(1)
	v679 = int32(0)
	if base.B2i32(v678 == v679)|base.B2i32(v360 == v679) != 0 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v676 = v674
	goto L126
L128:
	;
	v971 = v346 & int32(2)
	v972 = int32(0)
	if base.B2i32(v678 == v972)|base.B2i32(v370 == v972) != 0 {
		goto L202
	} else {
		goto L203
	}
L129:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v684 == int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v684)+4))
	if v687 <= int32(0) {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v709 = v8
	goto L132
L132:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v684)+12))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v720+v709<<(uint(int32(2))%32))))
	v725 = F_get_useful_group_keys_orderings(m, l0, v724)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L20
	} else {
		goto L135
	}
L133:
	;
	goto L128
L134:
	;
	v937 = v709 + int32(1)
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v684)+4))
	if v937 < v938 {
		v709 = v937
		goto L132
	} else {
		goto L201
	}
L135:
	;
	if v725 == int32(0) {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v729 = int32(0)
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v725)+4))
	if v730 <= v729 {
		goto L134
	} else {
		goto L137
	}
L137:
	;
	v743 = v729
	goto L138
L138:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v725)+12))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v763+v743<<(uint(int32(2))%32))))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v767)+4))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v724)+64))
	v771 = v33 + int32(16)
	if v768 == v769 {
		goto L144
	} else {
		goto L145
	}
L139:
	;
	goto L134
L140:
	;
	v903 = v743 + int32(1)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v725)+4))
	if v903 < v904 {
		v743 = v903
		goto L138
	} else {
		goto L200
	}
L141:
	;
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+36)))
	if v879 == int32(1) {
		goto L193
	} else {
		goto L194
	}
L142:
	;
	if v849 != 0 {
		goto L174
	} else {
		goto L175
	}
L143:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v768)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v771))) = v837
	v849 = int32(1)
	goto L142
L144:
	;
	if v768 != 0 {
		goto L143
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	if v768 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771))) = int32(0)
	v849 = int32(1)
	goto L142
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771))) = int32(0)
	v849 = int32(1)
	goto L142
L149:
	;
	goto L150
L150:
	;
	if v769 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v789 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v771))) = v789
	v849 = v789
	goto L142
L152:
	;
	goto L153
L153:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v769)+4))
	v793 = int32(0)
	if v793 < v792 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v796 = v792
	goto L156
L155:
	;
	v796 = v793
	goto L156
L156:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v768)+4))
	v802 = int32(0)
	goto L157
L157:
	;
	if v802 < v797 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v768)+12))
	v813 = v809 + v802<<(uint(int32(2))%32)
	goto L161
L160:
	;
	v813 = int32(0)
	goto L161
L161:
	;
	if v802 == v796 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771))) = v796
	v849 = base.B2i32(v813 == int32(0))
	goto L142
L163:
	;
	goto L164
L164:
	;
	v819 = base.B2i32(v813 == int32(0))
	if v813 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771))) = v802
	v849 = v819
	goto L142
L166:
	;
	goto L167
L167:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v769)+12))
	if v823 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771))) = v802
	v849 = v819
	goto L142
L169:
	;
	goto L170
L170:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v813)))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v823+v802<<(uint(int32(2))%32))))
	if v827 != v831 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771))) = v802
	v849 = int32(0)
	goto L142
L172:
	;
	v802 = v802 + int32(1)
	goto L157
L174:
	;
	v876 = v724
	goto L141
L175:
	;
	goto L176
L176:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_ordinary_grouping_paths[0])))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v724 == v360 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	if v861&int32(1) != 0 {
		goto L183
	} else {
		goto L184
	}
L178:
	;
	v861 = v851
	goto L177
L179:
	;
	goto L180
L180:
	;
	if v852 == int32(0) {
		goto L140
	} else {
		goto L181
	}
L181:
	;
	v856 = int32(1)
	if v851&v856 == int32(0) {
		goto L140
	} else {
		goto L182
	}
L182:
	;
	v861 = v856
	goto L177
L183:
	;
	v865 = v852
	goto L185
L184:
	;
	v865 = int32(0)
	goto L185
L185:
	;
	if v865 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v869 = F_create_sort_path(m, v380, v724, v768, float64(-1))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L20
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v872 = F_create_incremental_sort_path(m, l0, v380, v724, v768, v852, float64(-1))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L20
	} else {
		goto L191
	}
L189:
	;
	if v869 != 0 {
		v876 = v869
		goto L141
	} else {
		goto L190
	}
L190:
	;
	goto L140
L191:
	;
	if v872 == int32(0) {
		goto L140
	} else {
		goto L192
	}
L192:
	;
	v876 = v872
	goto L141
L193:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v380)+28))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v377)+100))
	v884 = int32(0)
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v767)+8))
	v889 = F_create_agg_path(m, l0, v380, v876, v882, base.B2i32(v883 != v884), int32(6), v887, v884, v640, v671)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L20
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v767)+8))
	v895 = F_create_group_path(m, l0, v380, v876, v893, int32(0), v671)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L20
	} else {
		goto L198
	}
L196:
	;
	F_add_path(m, v380, v889)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L20
	} else {
		goto L197
	}
L197:
	;
	goto L140
L198:
	;
	F_add_path(m, v380, v895)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L20
	} else {
		goto L199
	}
L199:
	;
	goto L140
L200:
	;
	goto L139
L201:
	;
	goto L133
L202:
	;
	v1264 = int32(0)
	if base.B2i32(v971 == v1264)|base.B2i32(v360 == v1264) == v1264 {
		goto L276
	} else {
		goto L277
	}
L203:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v977 == int32(0) {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v977)+4))
	if v980 <= int32(0) {
		goto L202
	} else {
		goto L205
	}
L205:
	;
	v1003 = int32(0)
	goto L206
L206:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v977)+12))
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1014+v1003<<(uint(int32(2))%32))))
	v1019 = F_get_useful_group_keys_orderings(m, l0, v1018)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L20
	} else {
		goto L209
	}
L207:
	;
	goto L202
L208:
	;
	v1231 = v1003 + int32(1)
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v977)+4))
	if v1231 < v1232 {
		v1003 = v1231
		goto L206
	} else {
		goto L275
	}
L209:
	;
	if v1019 == int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v1023 = int32(0)
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+4))
	if v1024 <= v1023 {
		goto L208
	} else {
		goto L211
	}
L211:
	;
	v1037 = v1023
	goto L212
L212:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+12))
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1057+v1037<<(uint(int32(2))%32))))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+4))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1018)+64))
	v1065 = v33 + int32(16)
	if v1062 == v1063 {
		goto L218
	} else {
		goto L219
	}
L213:
	;
	goto L208
L214:
	;
	v1197 = v1037 + int32(1)
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+4))
	if v1197 < v1198 {
		v1037 = v1197
		goto L212
	} else {
		goto L274
	}
L215:
	;
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+36)))
	if v1173 == int32(1) {
		goto L267
	} else {
		goto L268
	}
L216:
	;
	if v1143 != 0 {
		goto L248
	} else {
		goto L249
	}
L217:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1062)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1065))) = v1131
	v1143 = int32(1)
	goto L216
L218:
	;
	if v1062 != 0 {
		goto L217
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	if v1062 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1065))) = int32(0)
	v1143 = int32(1)
	goto L216
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1065))) = int32(0)
	v1143 = int32(1)
	goto L216
L223:
	;
	goto L224
L224:
	;
	if v1063 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1083 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1065))) = v1083
	v1143 = v1083
	goto L216
L226:
	;
	goto L227
L227:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+4))
	v1087 = int32(0)
	if v1087 < v1086 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1090 = v1086
	goto L230
L229:
	;
	v1090 = v1087
	goto L230
L230:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1062)+4))
	v1096 = int32(0)
	goto L231
L231:
	;
	if v1096 < v1091 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1062)+12))
	v1107 = v1103 + v1096<<(uint(int32(2))%32)
	goto L235
L234:
	;
	v1107 = int32(0)
	goto L235
L235:
	;
	if v1096 == v1090 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1065))) = v1090
	v1143 = base.B2i32(v1107 == int32(0))
	goto L216
L237:
	;
	goto L238
L238:
	;
	v1113 = base.B2i32(v1107 == int32(0))
	if v1107 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1065))) = v1096
	v1143 = v1113
	goto L216
L240:
	;
	goto L241
L241:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+12))
	if v1117 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1065))) = v1096
	v1143 = v1113
	goto L216
L243:
	;
	goto L244
L244:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1107)))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1117+v1096<<(uint(int32(2))%32))))
	if v1121 != v1125 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1065))) = v1096
	v1143 = int32(0)
	goto L216
L246:
	;
	v1096 = v1096 + int32(1)
	goto L231
L248:
	;
	v1170 = v1018
	goto L215
L249:
	;
	goto L250
L250:
	;
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_ordinary_grouping_paths[0])))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v1018 == v370 {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	if v1155&int32(1) != 0 {
		goto L257
	} else {
		goto L258
	}
L252:
	;
	v1155 = v1145
	goto L251
L253:
	;
	goto L254
L254:
	;
	if v1146 == int32(0) {
		goto L214
	} else {
		goto L255
	}
L255:
	;
	v1150 = int32(1)
	if v1145&v1150 == int32(0) {
		goto L214
	} else {
		goto L256
	}
L256:
	;
	v1155 = v1150
	goto L251
L257:
	;
	v1159 = v1146
	goto L259
L258:
	;
	v1159 = int32(0)
	goto L259
L259:
	;
	if v1159 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1163 = F_create_sort_path(m, v380, v1018, v1062, float64(-1))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L20
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	v1166 = F_create_incremental_sort_path(m, l0, v380, v1018, v1062, v1146, float64(-1))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L20
	} else {
		goto L265
	}
L263:
	;
	if v1163 != 0 {
		v1170 = v1163
		goto L215
	} else {
		goto L264
	}
L264:
	;
	goto L214
L265:
	;
	if v1166 == int32(0) {
		goto L214
	} else {
		goto L266
	}
L266:
	;
	v1170 = v1166
	goto L215
L267:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v380)+28))
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v377)+100))
	v1178 = int32(0)
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+8))
	v1183 = F_create_agg_path(m, l0, v380, v1170, v1176, base.B2i32(v1177 != v1178), int32(6), v1181, v1178, v640, v676)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L20
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+8))
	v1189 = F_create_group_path(m, l0, v380, v1170, v1187, int32(0), v676)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L20
	} else {
		goto L272
	}
L270:
	;
	F_add_partial_path(m, v380, v1183)
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L20
	} else {
		goto L271
	}
L271:
	;
	goto L214
L272:
	;
	F_add_partial_path(m, v380, v1189)
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L20
	} else {
		goto L273
	}
L273:
	;
	goto L214
L274:
	;
	goto L213
L275:
	;
	goto L207
L276:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v380)+28))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v1276 = F_create_agg_path(m, l0, v380, v360, v1271, int32(2), int32(6), v1274, int32(0), v640, v671)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L20
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	v1280 = int32(0)
	if base.B2i32(v971 == v1280)|base.B2i32(v370 == v1280) == v1280 {
		goto L281
	} else {
		goto L282
	}
L279:
	;
	F_add_path(m, v380, v1276)
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L20
	} else {
		goto L280
	}
L280:
	;
	goto L278
L281:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v380)+28))
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v1292 = F_create_agg_path(m, l0, v380, v370, v1287, int32(2), int32(6), v1290, int32(0), v640, v676)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L20
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v380)+168))
	if v1296 == int32(0) {
		v1314 = v380
		goto L50
	} else {
		goto L286
	}
L284:
	;
	F_add_partial_path(m, v380, v1292)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L20
	} else {
		goto L285
	}
L285:
	;
	goto L283
L286:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+36))
	if v1299 == int32(0) {
		v1314 = v380
		goto L50
	} else {
		goto L287
	}
L287:
	;
	m.T0[v1299].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, int32(1), l1, v380, l5)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L20
	} else {
		goto L288
	}
L288:
	;
	v1314 = v380
	goto L50
L289:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(l5)+100))
	if v1701 == int32(2) {
		goto L369
	} else {
		goto L370
	}
L290:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v1337 = int32(0)
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v1338 == v1337 {
		goto L294
	} else {
		goto L295
	}
L291:
	;
	v1655 = int32(0)
	if base.B2i32(v1314 == v1655)|base.B2i32(v1645 == v1655) != 0 {
		goto L360
	} else {
		goto L361
	}
L292:
	;
	if v1395 < int32(0) {
		goto L303
	} else {
		goto L304
	}
L293:
	;
	v1395 = base.I32_ctz(v1381) | v1382<<(uint(int32(5))%32)
	goto L292
L294:
	;
	v1395 = int32(-2)
	goto L292
L295:
	;
	v1348 = base.I32_div_s(int32(0), int32(32))
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+4))
	if v1349 <= v1348 {
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1352 = v1338 + int32(8)
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1352+v1348<<(uint(int32(2))%32))))
	v1359 = v1356 & int32(-1)
	if v1359 != 0 {
		v1381 = v1359
		v1382 = v1348
		goto L293
	} else {
		goto L297
	}
L297:
	;
	v1361 = v1348 + int32(1)
	if v1361 == v1349 {
		goto L294
	} else {
		goto L298
	}
L298:
	;
	v1364 = v1361
	goto L299
L299:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1352+v1364<<(uint(int32(2))%32))))
	if v1371 != 0 {
		v1381 = v1371
		v1382 = v1364
		goto L293
	} else {
		goto L301
	}
L300:
	;
	goto L294
L301:
	;
	v1373 = v1364 + int32(1)
	if v1373 != v1349 {
		v1364 = v1373
		goto L299
	} else {
		goto L302
	}
L302:
	;
	goto L300
L303:
	;
	v1632 = int32(0)
	v1637 = v1337
	v1645 = int32(1)
	goto L291
L304:
	;
	goto L305
L305:
	;
	v1409 = int32(0)
	v1414 = v1337
	v1416 = v1395
	v1422 = int32(1)
	goto L306
L306:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1432+v1416<<(uint(int32(2))%32))))
	v1437 = int32(0)
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+32))
	if v1439 == v1437 {
		v1460 = v1437
		goto L309
	} else {
		goto L310
	}
L307:
	;
	v1632 = v1558
	v1637 = v1561
	v1645 = v1564
	goto L291
L308:
	;
	if v1460 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L309:
	;
	goto L308
L310:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+12))
	v1443 = v1442
	goto L311
L311:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1443)))
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1446)))
	if base.Ui32(int32(2)) <= base.Ui32(v1447-int32(301)) {
		goto L313
	} else {
		goto L314
	}
L312:
	;
	v1460 = int32(1)
	goto L309
L313:
	;
	if v1447 != int32(290) {
		v1460 = v1437
		goto L309
	} else {
		goto L316
	}
L314:
	;
	v1443 = v1446 + int32(72)
	goto L311
L315:
	;
	goto L312
L316:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+72))
	if v1454 != 0 {
		v1460 = v1437
		goto L309
	} else {
		goto L317
	}
L317:
	;
	goto L315
L318:
	;
	v1463 = F_copy_pathtarget(m, v1336)
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L20
	} else {
		goto L321
	}
L319:
	;
	v1558 = v1409
	v1561 = v1414
	v1564 = v1422
	goto L320
L320:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v1566 == int32(0) {
		goto L350
	} else {
		goto L351
	}
L321:
	;
	base.MemoryCopy(m, v33+int32(16), l5, int32(104))
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+8))
	v1472 = F_find_appinfos_by_relids(m, l0, v1469, v33+int32(124))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L20
	} else {
		goto L322
	}
L322:
	;
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1336)+4))
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
	v1476 = F_adjust_appendrel_attrs(m, l0, v1474, v1475, v1472)
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L20
	} else {
		goto L323
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1463)+4)) = v1476
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(l5)+92))
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
	v1481 = F_adjust_appendrel_attrs(m, l0, v1479, v1480, v1472)
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L20
	} else {
		goto L324
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1481
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
	v1486 = F_adjust_appendrel_attrs(m, l0, v1484, v1485, v1472)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L20
	} else {
		goto L325
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1486
	v1490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+88)))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v33)+108))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+4))
	v1499 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v1492))|base.B2i32(int32(1)<<(uint(v1492)%32)&int32(44) == v1499) == v1499 {
		goto L327
	} else {
		goto L328
	}
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1514)+28)) = v1463
	v1516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1436)+26)))
	if v1490&v1516 == int32(0) {
		goto L332
	} else {
		goto L333
	}
L327:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+8))
	v1506 = F_fetch_upper_rel(m, l0, int32(2), v1505)
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L20
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	v1512 = F_fetch_upper_rel(m, l0, int32(2), int32(0))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L20
	} else {
		goto L331
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+4)) = int32(5)
	v1514 = v1506
	goto L326
L331:
	;
	v1514 = v1512
	goto L326
L332:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v1514)+156)) = v1526
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v1514)+160)) = v1528
	v1530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1436)+164)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1514)+164)) = uint8(v1530)
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v1514)+168)) = v1532
	F_create_ordinary_grouping_paths(m, l0, v1436, v1514, l3, l4, v33+int32(16), v33+int32(12))
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L20
	} else {
		goto L336
	}
L333:
	;
	v1520 = F_is_parallel_safe(m, l0, v1491)
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L20
	} else {
		goto L334
	}
L334:
	;
	if v1520 == int32(0) {
		goto L332
	} else {
		goto L335
	}
L335:
	;
	v1524 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1514)+26)) = uint8(v1524)
	goto L332
L336:
	;
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	if v1540 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	if v338 == int32(1) {
		goto L342
	} else {
		goto L343
	}
L338:
	;
	v1546 = v1414
	v1547 = int32(0)
	goto L337
L339:
	;
	goto L340
L340:
	;
	v1544 = F_lappend(m, v1414, v1540)
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L20
	} else {
		goto L341
	}
L341:
	;
	v1546 = v1544
	v1547 = v1422
	goto L337
L342:
	;
	F_set_cheapest(m, v1514)
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L20
	} else {
		goto L345
	}
L343:
	;
	v1554 = v1409
	goto L344
L344:
	;
	F_pfree(m, v1472)
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L20
	} else {
		goto L347
	}
L345:
	;
	v1552 = F_lappend(m, v1409, v1514)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L20
	} else {
		goto L346
	}
L346:
	;
	v1554 = v1552
	goto L344
L347:
	;
	v1558 = v1554
	v1561 = v1546
	v1564 = v1547
	goto L320
L348:
	;
	if int32(0) <= v1622 {
		v1409 = v1558
		v1414 = v1561
		v1416 = v1622
		v1422 = v1564
		goto L306
	} else {
		goto L359
	}
L349:
	;
	v1622 = base.I32_ctz(v1608) | v1609<<(uint(int32(5))%32)
	goto L348
L350:
	;
	v1622 = int32(-2)
	goto L348
L351:
	;
	v1573 = v1416 + int32(1)
	v1575 = base.I32_div_s(v1573, int32(32))
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1566)+4))
	if v1576 <= v1575 {
		goto L350
	} else {
		goto L352
	}
L352:
	;
	v1579 = v1566 + int32(8)
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1579+v1575<<(uint(int32(2))%32))))
	v1586 = v1583 & (int32(-1) << (uint(v1573) % 32))
	if v1586 != 0 {
		v1608 = v1586
		v1609 = v1575
		goto L349
	} else {
		goto L353
	}
L353:
	;
	v1588 = v1575 + int32(1)
	if v1588 == v1576 {
		goto L350
	} else {
		goto L354
	}
L354:
	;
	v1591 = v1588
	goto L355
L355:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v1579+v1591<<(uint(int32(2))%32))))
	if v1598 != 0 {
		v1608 = v1598
		v1609 = v1591
		goto L349
	} else {
		goto L357
	}
L356:
	;
	goto L350
L357:
	;
	v1600 = v1591 + int32(1)
	if v1600 != v1576 {
		v1591 = v1600
		goto L355
	} else {
		goto L358
	}
L358:
	;
	goto L356
L359:
	;
	goto L307
L360:
	;
	if v338 != int32(1) {
		goto L289
	} else {
		goto L365
	}
L361:
	;
	F_add_paths_to_append_rel(m, l0, v1314, v1637)
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L20
	} else {
		goto L362
	}
L362:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1314)+32))
	if v1662 == int32(0) {
		goto L360
	} else {
		goto L363
	}
L363:
	;
	F_set_cheapest(m, v1314)
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L20
	} else {
		goto L364
	}
L364:
	;
	goto L360
L365:
	;
	F_add_paths_to_append_rel(m, l0, l2, v1632)
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L20
	} else {
		goto L366
	}
L366:
	;
	goto L289
L367:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L20
	} else {
		goto L557
	}
L368:
	;
	m.G0 = v33 + int32(128)
	return
L369:
	;
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v1314)+32))
	if v1704 == int32(0) {
		goto L368
	} else {
		goto L372
	}
L370:
	;
	goto L371
L371:
	;
	if v1314 == int32(0) {
		goto L374
	} else {
		goto L375
	}
L372:
	;
	F_set_cheapest(m, v1314)
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L20
	} else {
		goto L373
	}
L373:
	;
	goto L368
L374:
	;
	v1719 = l5 + int32(48)
	v1720 = *(*float64)(unsafe.Add(mBase, uint32(v35)+32))
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	v1722 = F_get_number_of_groups(m, l0, v1720, l4, v1721)
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		goto L20
	} else {
		goto L379
	}
L375:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1314)+40))
	if v1711 == int32(0) {
		goto L374
	} else {
		goto L376
	}
L376:
	;
	F_gather_grouping_paths(m, l0, v1314)
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L20
	} else {
		goto L377
	}
L377:
	;
	F_set_cheapest(m, v1314)
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L20
	} else {
		goto L378
	}
L378:
	;
	goto L374
L379:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1726 = v1724 & int32(2)
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(l5)+92))
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1724&int32(1) == int32(0) {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	if v1726 == int32(0) {
		goto L533
	} else {
		goto L534
	}
L381:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v1734 == int32(0) {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	if v1314 == int32(0) {
		goto L380
	} else {
		goto L460
	}
L383:
	;
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v1734)+4))
	if v1737 <= int32(0) {
		goto L382
	} else {
		goto L384
	}
L384:
	;
	v1763 = int32(0)
	goto L385
L385:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1734)+12))
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(v1771+v1763<<(uint(int32(2))%32))))
	v1776 = F_get_useful_group_keys_orderings(m, l0, v1775)
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L20
	} else {
		goto L388
	}
L386:
	;
	goto L382
L387:
	;
	v1995 = v1763 + int32(1)
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(v1734)+4))
	if v1995 < v1996 {
		v1763 = v1995
		goto L385
	} else {
		goto L459
	}
L388:
	;
	if v1776 == int32(0) {
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v1780 = int32(0)
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+4))
	if v1781 <= v1780 {
		goto L387
	} else {
		goto L390
	}
L390:
	;
	v1794 = v1780
	goto L391
L391:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+12))
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1814+v1794<<(uint(int32(2))%32))))
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+4))
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v1775)+64))
	v1822 = v33 + int32(16)
	if v1819 == v1820 {
		goto L397
	} else {
		goto L398
	}
L392:
	;
	goto L387
L393:
	;
	v1961 = v1794 + int32(1)
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+4))
	if v1961 < v1962 {
		v1794 = v1961
		goto L391
	} else {
		goto L458
	}
L394:
	;
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1729)+108))
	if v1930 != 0 {
		goto L446
	} else {
		goto L447
	}
L395:
	;
	if v1900 != 0 {
		goto L427
	} else {
		goto L428
	}
L396:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1819)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1822))) = v1888
	v1900 = int32(1)
	goto L395
L397:
	;
	if v1819 != 0 {
		goto L396
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	if v1819 == int32(0) {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1822))) = int32(0)
	v1900 = int32(1)
	goto L395
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1822))) = int32(0)
	v1900 = int32(1)
	goto L395
L402:
	;
	goto L403
L403:
	;
	if v1820 == int32(0) {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1840 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1822))) = v1840
	v1900 = v1840
	goto L395
L405:
	;
	goto L406
L406:
	;
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1820)+4))
	v1844 = int32(0)
	if v1844 < v1843 {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v1847 = v1843
	goto L409
L408:
	;
	v1847 = v1844
	goto L409
L409:
	;
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1819)+4))
	v1853 = int32(0)
	goto L410
L410:
	;
	if v1853 < v1848 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1819)+12))
	v1864 = v1860 + v1853<<(uint(int32(2))%32)
	goto L414
L413:
	;
	v1864 = int32(0)
	goto L414
L414:
	;
	if v1853 == v1847 {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1822))) = v1847
	v1900 = base.B2i32(v1864 == int32(0))
	goto L395
L416:
	;
	goto L417
L417:
	;
	v1870 = base.B2i32(v1864 == int32(0))
	if v1864 == int32(0) {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1822))) = v1853
	v1900 = v1870
	goto L395
L419:
	;
	goto L420
L420:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1820)+12))
	if v1874 == int32(0) {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1822))) = v1853
	v1900 = v1870
	goto L395
L422:
	;
	goto L423
L423:
	;
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1864)))
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v1874+v1853<<(uint(int32(2))%32))))
	if v1878 != v1882 {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1822))) = v1853
	v1900 = int32(0)
	goto L395
L425:
	;
	v1853 = v1853 + int32(1)
	goto L410
L427:
	;
	v1927 = v1775
	goto L394
L428:
	;
	goto L429
L429:
	;
	v1902 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_ordinary_grouping_paths[0])))
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v1775 == v1728 {
		goto L431
	} else {
		goto L432
	}
L430:
	;
	if v1912&int32(1) != 0 {
		goto L436
	} else {
		goto L437
	}
L431:
	;
	v1912 = v1902
	goto L430
L432:
	;
	goto L433
L433:
	;
	if v1903 == int32(0) {
		goto L393
	} else {
		goto L434
	}
L434:
	;
	v1907 = int32(1)
	if v1902&v1907 == int32(0) {
		goto L393
	} else {
		goto L435
	}
L435:
	;
	v1912 = v1907
	goto L430
L436:
	;
	v1916 = v1903
	goto L438
L437:
	;
	v1916 = int32(0)
	goto L438
L438:
	;
	if v1916 == int32(0) {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v1920 = F_create_sort_path(m, l2, v1775, v1819, float64(-1))
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L20
	} else {
		goto L442
	}
L440:
	;
	goto L441
L441:
	;
	v1923 = F_create_incremental_sort_path(m, l0, l2, v1775, v1819, v1903, float64(-1))
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L20
	} else {
		goto L444
	}
L442:
	;
	if v1920 != 0 {
		v1927 = v1920
		goto L394
	} else {
		goto L443
	}
L443:
	;
	goto L393
L444:
	;
	if v1923 == int32(0) {
		goto L393
	} else {
		goto L445
	}
L445:
	;
	v1927 = v1923
	goto L394
L446:
	;
	F_consider_groupingsets_paths(m, l0, l2, v1927, int32(1), base.B2i32(v1726 != int32(0)), l4, l3, v1722)
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L20
	} else {
		goto L449
	}
L447:
	;
	goto L448
L448:
	;
	v1936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1729)+36)))
	if v1936 == int32(1) {
		goto L450
	} else {
		goto L451
	}
L449:
	;
	goto L393
L450:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1729)+100))
	v1941 = int32(0)
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+8))
	v1945 = F_create_agg_path(m, l0, l2, v1927, v1939, base.B2i32(v1940 != v1941), v1941, v1944, v1727, l3, v1722)
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L20
	} else {
		goto L453
	}
L451:
	;
	goto L452
L452:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v1729)+100))
	if v1949 == int32(0) {
		goto L393
	} else {
		goto L455
	}
L453:
	;
	F_add_path(m, l2, v1945)
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L20
	} else {
		goto L454
	}
L454:
	;
	goto L393
L455:
	;
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+8))
	v1953 = F_create_group_path(m, l0, l2, v1927, v1952, v1727, v1722)
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L20
	} else {
		goto L456
	}
L456:
	;
	F_add_path(m, l2, v1953)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L20
	} else {
		goto L457
	}
L457:
	;
	goto L393
L458:
	;
	goto L392
L459:
	;
	goto L386
L460:
	;
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v1314)+32))
	if v2030 == int32(0) {
		goto L380
	} else {
		goto L461
	}
L461:
	;
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v2030)+4))
	if v2033 <= int32(0) {
		goto L380
	} else {
		goto L462
	}
L462:
	;
	v2059 = int32(0)
	goto L463
L463:
	;
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v2030)+12))
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v2067+v2059<<(uint(int32(2))%32))))
	v2072 = F_get_useful_group_keys_orderings(m, l0, v2071)
	mBase = m.M
	v2073 = m.ExcPending
	if v2073 != 0 {
		goto L20
	} else {
		goto L466
	}
L464:
	;
	goto L380
L465:
	;
	v2283 = v2059 + int32(1)
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(v2030)+4))
	if v2283 < v2284 {
		v2059 = v2283
		goto L463
	} else {
		goto L532
	}
L466:
	;
	if v2072 == int32(0) {
		goto L465
	} else {
		goto L467
	}
L467:
	;
	v2076 = int32(0)
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v2072)+4))
	if v2077 <= v2076 {
		goto L465
	} else {
		goto L468
	}
L468:
	;
	v2090 = v2076
	goto L469
L469:
	;
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v1314)+48))
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(v2072)+12))
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v2111+v2090<<(uint(int32(2))%32))))
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v2115)+4))
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v2071)+64))
	v2119 = v33 + int32(16)
	if v2116 == v2117 {
		goto L475
	} else {
		goto L476
	}
L470:
	;
	goto L465
L471:
	;
	v2249 = v2090 + int32(1)
	v2250 = *(*int32)(unsafe.Add(mBase, uint32(v2072)+4))
	if v2249 < v2250 {
		v2090 = v2249
		goto L469
	} else {
		goto L531
	}
L472:
	;
	v2227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1729)+36)))
	if v2227 == int32(1) {
		goto L524
	} else {
		goto L525
	}
L473:
	;
	if v2197 != 0 {
		goto L505
	} else {
		goto L506
	}
L474:
	;
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(v2116)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2119))) = v2185
	v2197 = int32(1)
	goto L473
L475:
	;
	if v2116 != 0 {
		goto L474
	} else {
		goto L478
	}
L476:
	;
	goto L477
L477:
	;
	if v2116 == int32(0) {
		goto L479
	} else {
		goto L480
	}
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2119))) = int32(0)
	v2197 = int32(1)
	goto L473
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2119))) = int32(0)
	v2197 = int32(1)
	goto L473
L480:
	;
	goto L481
L481:
	;
	if v2117 == int32(0) {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	v2137 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2119))) = v2137
	v2197 = v2137
	goto L473
L483:
	;
	goto L484
L484:
	;
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v2117)+4))
	v2141 = int32(0)
	if v2141 < v2140 {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v2144 = v2140
	goto L487
L486:
	;
	v2144 = v2141
	goto L487
L487:
	;
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v2116)+4))
	v2150 = int32(0)
	goto L488
L488:
	;
	if v2150 < v2145 {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v2116)+12))
	v2161 = v2157 + v2150<<(uint(int32(2))%32)
	goto L492
L491:
	;
	v2161 = int32(0)
	goto L492
L492:
	;
	if v2150 == v2144 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2119))) = v2144
	v2197 = base.B2i32(v2161 == int32(0))
	goto L473
L494:
	;
	goto L495
L495:
	;
	v2167 = base.B2i32(v2161 == int32(0))
	if v2161 == int32(0) {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2119))) = v2150
	v2197 = v2167
	goto L473
L497:
	;
	goto L498
L498:
	;
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v2117)+12))
	if v2171 == int32(0) {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2119))) = v2150
	v2197 = v2167
	goto L473
L500:
	;
	goto L501
L501:
	;
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v2161)))
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v2171+v2150<<(uint(int32(2))%32))))
	if v2175 != v2179 {
		goto L502
	} else {
		goto L503
	}
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2119))) = v2150
	v2197 = int32(0)
	goto L473
L503:
	;
	v2150 = v2150 + int32(1)
	goto L488
L505:
	;
	v2224 = v2071
	goto L472
L506:
	;
	goto L507
L507:
	;
	v2199 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_ordinary_grouping_paths[0])))
	v2200 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v2071 == v2110 {
		goto L509
	} else {
		goto L510
	}
L508:
	;
	if v2209&int32(1) != 0 {
		goto L514
	} else {
		goto L515
	}
L509:
	;
	v2209 = v2199
	goto L508
L510:
	;
	goto L511
L511:
	;
	if v2200 == int32(0) {
		goto L471
	} else {
		goto L512
	}
L512:
	;
	v2204 = int32(1)
	if v2199&v2204 == int32(0) {
		goto L471
	} else {
		goto L513
	}
L513:
	;
	v2209 = v2204
	goto L508
L514:
	;
	v2213 = v2200
	goto L516
L515:
	;
	v2213 = int32(0)
	goto L516
L516:
	;
	if v2213 == int32(0) {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v2217 = F_create_sort_path(m, l2, v2071, v2116, float64(-1))
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L20
	} else {
		goto L520
	}
L518:
	;
	goto L519
L519:
	;
	v2220 = F_create_incremental_sort_path(m, l0, l2, v2071, v2116, v2200, float64(-1))
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		goto L20
	} else {
		goto L522
	}
L520:
	;
	if v2217 != 0 {
		v2224 = v2217
		goto L472
	} else {
		goto L521
	}
L521:
	;
	goto L471
L522:
	;
	if v2220 == int32(0) {
		goto L471
	} else {
		goto L523
	}
L523:
	;
	v2224 = v2220
	goto L472
L524:
	;
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v1729)+100))
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v2115)+8))
	v2236 = F_create_agg_path(m, l0, l2, v2224, v2230, base.B2i32(v2231 != int32(0)), int32(9), v2235, v1727, v1719, v1722)
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L20
	} else {
		goto L527
	}
L525:
	;
	goto L526
L526:
	;
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2115)+8))
	v2241 = F_create_group_path(m, l0, l2, v2224, v2240, v1727, v1722)
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L20
	} else {
		goto L529
	}
L527:
	;
	F_add_path(m, l2, v2236)
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L20
	} else {
		goto L528
	}
L528:
	;
	goto L471
L529:
	;
	F_add_path(m, l2, v2241)
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		goto L20
	} else {
		goto L530
	}
L530:
	;
	goto L471
L531:
	;
	goto L470
L532:
	;
	goto L464
L533:
	;
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v2345 != 0 {
		goto L546
	} else {
		goto L547
	}
L534:
	;
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(v1729)+108))
	if v2318 != 0 {
		goto L536
	} else {
		goto L537
	}
L535:
	;
	if v1314 == int32(0) {
		goto L533
	} else {
		goto L542
	}
L536:
	;
	F_consider_groupingsets_paths(m, l0, l2, v1728, int32(0), int32(1), l4, l3, v1722)
	mBase = m.M
	v2322 = m.ExcPending
	if v2322 != 0 {
		goto L20
	} else {
		goto L539
	}
L537:
	;
	goto L538
L538:
	;
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v2327 = F_create_agg_path(m, l0, l2, v1728, v2323, int32(2), int32(0), v2326, v1727, l3, v1722)
	mBase = m.M
	v2328 = m.ExcPending
	if v2328 != 0 {
		goto L20
	} else {
		goto L540
	}
L539:
	;
	goto L535
L540:
	;
	F_add_path(m, l2, v2327)
	mBase = m.M
	v2330 = m.ExcPending
	if v2330 != 0 {
		goto L20
	} else {
		goto L541
	}
L541:
	;
	goto L535
L542:
	;
	v2333 = *(*int32)(unsafe.Add(mBase, uint32(v1314)+32))
	if v2333 == int32(0) {
		goto L533
	} else {
		goto L543
	}
L543:
	;
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(v1314)+48))
	v2337 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v2341 = F_create_agg_path(m, l0, l2, v2336, v2337, int32(2), int32(9), v2340, v1727, v1719, v1722)
	mBase = m.M
	v2342 = m.ExcPending
	if v2342 != 0 {
		goto L20
	} else {
		goto L544
	}
L544:
	;
	F_add_path(m, l2, v2341)
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		goto L20
	} else {
		goto L545
	}
L545:
	;
	goto L533
L546:
	;
	F_gather_grouping_paths(m, l0, l2)
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L20
	} else {
		goto L549
	}
L547:
	;
	goto L548
L548:
	;
	v2348 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v2348 == int32(0) {
		goto L367
	} else {
		goto L550
	}
L549:
	;
	goto L548
L550:
	;
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(l2)+168))
	if v2351 == int32(0) {
		goto L551
	} else {
		goto L552
	}
L551:
	;
	v2362 = *(*int32)(unsafe.Add(mBase, _c_F_create_ordinary_grouping_paths[1]))
	if v2362 == int32(0) {
		goto L368
	} else {
		goto L555
	}
L552:
	;
	v2354 = *(*int32)(unsafe.Add(mBase, uint32(v2351)+36))
	if v2354 == int32(0) {
		goto L551
	} else {
		goto L553
	}
L553:
	;
	m.T0[v2354].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, int32(2), l1, l2, l5)
	mBase = m.M
	v2359 = m.ExcPending
	if v2359 != 0 {
		goto L20
	} else {
		goto L554
	}
L554:
	;
	goto L551
L555:
	;
	m.T0[v2362].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, int32(2), l1, l2, l5)
	mBase = m.M
	v2367 = m.ExcPending
	if v2367 != 0 {
		goto L20
	} else {
		goto L556
	}
L556:
	;
	goto L368
L557:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
		goto L20
	} else {
		goto L558
	}
L558:
	;
	F_errmsg(m, int32(_a_F_create_ordinary_grouping_paths_0), int32(0))
	mBase = m.M
	v2411 = m.ExcPending
	if v2411 != 0 {
		goto L20
	} else {
		goto L559
	}
L559:
	;
	F_errdetail(m, int32(_a_F_create_ordinary_grouping_paths_1), int32(0))
	mBase = m.M
	v2415 = m.ExcPending
	if v2415 != 0 {
		goto L20
	} else {
		goto L560
	}
L560:
	;
	F_errfinish(m, int32(_a_F_create_ordinary_grouping_paths_2), int32(_a_F_create_ordinary_grouping_paths_3), int32(_a_F_create_ordinary_grouping_paths_4))
	mBase = m.M
	v2420 = m.ExcPending
	if v2420 != 0 {
		goto L20
	} else {
		goto L561
	}
L561:
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
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
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
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
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
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
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
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v489 int32
	_ = v489
	var v491 float64
	_ = v491
	var v493 float64
	_ = v493
	var v495 float64
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
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
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v544 int32
	_ = v544
	var v546 float64
	_ = v546
	var v548 float64
	_ = v548
	var v550 float64
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v663 int32
	_ = v663
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
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
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
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v778 int32
	_ = v778
	var v780 float64
	_ = v780
	var v782 float64
	_ = v782
	var v784 float64
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v890 int32
	_ = v890
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v919 int32
	_ = v919
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1009 int32
	_ = v1009
	var v1011 float64
	_ = v1011
	var v1013 float64
	_ = v1013
	var v1015 float64
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1107 int32
	_ = v1107
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1144 int32
	_ = v1144
	var v1146 float64
	_ = v1146
	var v1148 float64
	_ = v1148
	var v1150 float64
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
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
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1186 int32
	_ = v1186
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1240 int32
	_ = v1240
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1287 int32
	_ = v1287
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1312 int32
	_ = v1312
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1338 int32
	_ = v1338
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1435 int32
	_ = v1435
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1483 int32
	_ = v1483
	var v1485 float64
	_ = v1485
	var v1487 float64
	_ = v1487
	var v1489 float64
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
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
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
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
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1540 int32
	_ = v1540
	var v1542 float64
	_ = v1542
	var v1544 float64
	_ = v1544
	var v1546 float64
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
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
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
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
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1595 int32
	_ = v1595
	var v1597 float64
	_ = v1597
	var v1599 float64
	_ = v1599
	var v1601 float64
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1650 int32
	_ = v1650
	var v1652 float64
	_ = v1652
	var v1654 float64
	_ = v1654
	var v1656 float64
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1702 int32
	_ = v1702
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1713 int32
	_ = v1713
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1732 int32
	_ = v1732
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1779 int32
	_ = v1779
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1807 int32
	_ = v1807
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1836 int32
	_ = v1836
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
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
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1882 int32
	_ = v1882
	var v1884 float64
	_ = v1884
	var v1886 float64
	_ = v1886
	var v1888 float64
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1934 int32
	_ = v1934
	var v1936 float64
	_ = v1936
	var v1938 float64
	_ = v1938
	var v1940 float64
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1969 int32
	_ = v1969
	var v1971 float64
	_ = v1971
	var v1973 float64
	_ = v1973
	var v1975 float64
	_ = v1975
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2020 int32
	_ = v2020
	var v2023 int32
	_ = v2023
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2034 int32
	_ = v2034
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
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
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2064 int32
	_ = v2064
	var v2066 float64
	_ = v2066
	var v2068 float64
	_ = v2068
	var v2070 float64
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2082 int32
	_ = v2082
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2113 float64
	_ = v2113
	var v2115 float64
	_ = v2115
	var v2117 float64
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2122 int32
	_ = v2122
	var v2124 int32
	_ = v2124
	var v2126 int32
	_ = v2126
	var v2128 int32
	_ = v2128
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
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2164 int32
	_ = v2164
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2179 int32
	_ = v2179
	var v2185 int32
	_ = v2185
	var v2199 int32
	_ = v2199
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2255 int32
	_ = v2255
	var v2257 int32
	_ = v2257
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2287 int32
	_ = v2287
	var v2291 int32
	_ = v2291
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2303 int32
	_ = v2303
	var v2317 int32
	_ = v2317
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2325 float64
	_ = v2325
	var v2327 float64
	_ = v2327
	var v2329 float64
	_ = v2329
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2334 int32
	_ = v2334
	var v2336 int32
	_ = v2336
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2360 int32
	_ = v2360
	var v2365 int32
	_ = v2365
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2372 int32
	_ = v2372
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2402 int32
	_ = v2402
	var v2407 int32
	_ = v2407
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2418 int32
	_ = v2418
	var v2423 int32
	_ = v2423
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2450 int32
	_ = v2450
	var v2455 int32
	_ = v2455
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2466 int32
	_ = v2466
	var v2471 int32
	_ = v2471
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2482 int32
	_ = v2482
	var v2487 int32
	_ = v2487
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
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+319)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	switch v42 - int32(1) {
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
	v66 = int32(0)
	if v65 != 0 {
		goto L23
	} else {
		goto L24
	}
L11:
	;
	v60 = F_order_qual_clauses(m, l0, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L20
	}
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v53 == int32(354) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	if v41&int32(1) != 0 {
		v59 = v40
		goto L11
	} else {
		goto L16
	}
L14:
	;
	if v41&int32(1) != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v65 = int32(0)
	goto L10
L16:
	;
	v65 = int32(0)
	goto L10
L17:
	;
	v56 = int32(76)
	goto L19
L18:
	;
	v56 = int32(80)
	goto L19
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1+v56)))
	v59 = v58
	goto L11
L20:
	;
	v63 = F_extract_actual_clauses(m, v60, int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v65 = v63
	goto L10
L22:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v465 - int32(331) {
	case 0:
		goto L137
	default:
		goto L133
	case 8:
		goto L149
	case 9:
		goto L148
	case 10:
		goto L132
	case 11:
		goto L147
	case 13:
		goto L146
	case 14:
		goto L145
	case 15:
		goto L144
	case 16:
		goto L143
	case 17:
		goto L142
	case 18:
		goto L140
	case 19:
		goto L141
	case 20:
		goto L139
	case 21:
		goto L138
	case 22:
		goto L136
	case 23:
		goto L135
	case 24:
		goto L134
	}
L23:
	;
	v68 = v66
	goto L25
L24:
	;
	v68 = l2
	goto L25
L25:
	;
	if v68 == int32(8) {
		v459 = v66
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v71 = F_use_physical_tlist(m, l0, l1, v68)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L29
	}
L27:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_apply_pathtarget_labeling_to_tlist(m, v440, v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L8
	} else {
		goto L125
	}
L28:
	;
	if v68&int32(4) == int32(0) {
		v459 = v288
		goto L22
	} else {
		goto L124
	}
L29:
	;
	if v71 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v73 == int32(342) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	if v373 == int32(0) {
		v459 = v66
		goto L22
	} else {
		goto L110
	}
L33:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+92))
	v78 = F_copyObjectImpl(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v82 = m.G0
	v84 = v82 - int32(16)
	m.G0 = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v87 != 0 {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	if v68&int32(4) != 0 {
		v440 = v78
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v459 = v78
	goto L22
L38:
	;
	if v288 != 0 {
		goto L28
	} else {
		goto L95
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L8
	} else {
		goto L92
	}
L40:
	;
	m.G0 = v84 + int32(16)
	goto L38
L41:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	if base.Ui32(int32(6)) <= base.Ui32(v101-int32(3)) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v99 = v87 + v86<<(uint(int32(2))%32)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+52))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v99 = v93 + v86<<(uint(int32(2))%32) - int32(4)
	goto L41
L45:
	;
	switch v101 {
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
	v232 = int32(0)
	F_expandRTE(m, v100, v86, v232, v232, int32(-1), int32(1), v232, v84+int32(12))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L8
	} else {
		goto L81
	}
L48:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v100)+36))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+76))
	if v188 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L49:
	;
	v106 = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	v109 = F_table_open(m, v107, v106)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+48))
	v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(v111)+120)))
	if int32(0) < v112 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v119 = v106
	v120 = int32(1)
	goto L54
L52:
	;
	v171 = v106
	goto L53
L53:
	;
	F_relation_close(m, v109, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L8
	} else {
		goto L68
	}
L54:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v109)+52))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v139 = v132 + v133<<(uint(int32(4))%32) + v120*int32(100)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+11)))
	if v140 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v171 = v164
	goto L53
L56:
	;
	v141 = int32(0)
	F_relation_close(m, v109, v141)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v146 = v139 - int32(80)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+88)))
	if v147 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v288 = v141
	goto L40
L60:
	;
	v148 = int32(0)
	F_relation_close(m, v109, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L8
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v153 = base.I32_extend16_s(v120)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v146)+68))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v146)+76))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v146)+96))
	v158 = F_makeVar(m, v86, v153, v154, v155, v156, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L8
	} else {
		goto L64
	}
L63:
	;
	v288 = v148
	goto L40
L64:
	;
	v160 = int32(0)
	v162 = F_makeTargetEntry(m, v158, v153, v160, v160)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L65
	}
L65:
	;
	v164 = F_lappend(m, v119, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	if v120 != v112 {
		v119 = v164
		v120 = v120 + int32(1)
		goto L54
	} else {
		goto L67
	}
L67:
	;
	goto L55
L68:
	;
	v288 = v171
	goto L40
L69:
	;
	v288 = int32(0)
	goto L40
L70:
	;
	goto L71
L71:
	;
	v192 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v193 <= v192 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v288 = int32(0)
	goto L40
L73:
	;
	goto L74
L74:
	;
	v201 = int32(0)
	v202 = v192
	goto L75
L75:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214+v202<<(uint(int32(2))%32))))
	v219 = F_makeVarFromTargetEntry(m, v86, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
	} else {
		goto L77
	}
L76:
	;
	v288 = v226
	goto L40
L77:
	;
	v221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v218)+8)))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+26)))
	v224 = F_makeTargetEntry(m, v219, v221, int32(0), v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	v226 = F_lappend(m, v201, v224)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v229 = v202 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v229 < v230 {
		v201 = v226
		v202 = v229
		goto L75
	} else {
		goto L80
	}
L80:
	;
	goto L76
L81:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	if v242 == int32(0) {
		v288 = v232
		goto L40
	} else {
		goto L82
	}
L82:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if v245 <= int32(0) {
		v288 = v232
		goto L40
	} else {
		goto L83
	}
L83:
	;
	v251 = int32(0)
	v252 = v232
	goto L84
L84:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v265+v251<<(uint(int32(2))%32))))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	if v270 != int32(6) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v288 = v279
	goto L40
L86:
	;
	v288 = int32(0)
	goto L40
L87:
	;
	goto L88
L88:
	;
	v274 = int32(*(*int16)(unsafe.Add(mBase, uint32(v269)+8)))
	v275 = int32(0)
	v277 = F_makeTargetEntry(m, v269, v274, v275, v275)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	v279 = F_lappend(m, v252, v277)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	v282 = v251 + int32(1)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if v282 < v283 {
		v251 = v282
		v252 = v279
		goto L84
	} else {
		goto L91
	}
L91:
	;
	goto L85
L92:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v308
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_0), v84)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_1), int32(1903), int32(_a_F_create_scan_plan_2))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
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
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	if v319 == int32(0) {
		v459 = v66
		goto L22
	} else {
		goto L96
	}
L96:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	if v323 <= int32(0) {
		v459 = v66
		goto L22
	} else {
		goto L97
	}
L97:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	v330 = int32(1)
	v332 = v4
	v337 = v66
	goto L98
L98:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v343+v332<<(uint(int32(2))%32))))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v348 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v459 = v366
	goto L22
L100:
	;
	v349 = F_replace_nestloop_params_mutator(m, v347, l0)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L8
	} else {
		goto L103
	}
L101:
	;
	v351 = v347
	goto L102
L102:
	;
	v353 = int32(0)
	v355 = F_makeTargetEntry(m, v351, base.I32_extend16_s(v330), v353, v353)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L8
	} else {
		goto L104
	}
L103:
	;
	v351 = v349
	goto L102
L104:
	;
	if v326 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v326+v330<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v355)+16)) = v362
	goto L107
L106:
	;
	goto L107
L107:
	;
	v366 = F_lappend(m, v337, v355)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	v369 = v332 + int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	if v369 < v370 {
		v330 = v330 + int32(1)
		v332 = v369
		v337 = v366
		goto L98
	} else {
		goto L109
	}
L109:
	;
	goto L99
L110:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v377 <= int32(0) {
		v459 = v66
		goto L22
	} else {
		goto L111
	}
L111:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	v384 = int32(1)
	v386 = v4
	v391 = v66
	goto L112
L112:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v397+v386<<(uint(int32(2))%32))))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v402 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v459 = v420
	goto L22
L114:
	;
	v403 = F_replace_nestloop_params_mutator(m, v401, l0)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L8
	} else {
		goto L117
	}
L115:
	;
	v405 = v401
	goto L116
L116:
	;
	v407 = int32(0)
	v409 = F_makeTargetEntry(m, v405, base.I32_extend16_s(v384), v407, v407)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L8
	} else {
		goto L118
	}
L117:
	;
	v405 = v403
	goto L116
L118:
	;
	if v380 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v380+v384<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v409)+16)) = v416
	goto L121
L120:
	;
	goto L121
L121:
	;
	v420 = F_lappend(m, v391, v409)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L8
	} else {
		goto L122
	}
L122:
	;
	v423 = v386 + int32(1)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v423 < v424 {
		v384 = v384 + int32(1)
		v386 = v423
		v391 = v420
		goto L112
	} else {
		goto L123
	}
L123:
	;
	goto L113
L124:
	;
	v440 = v288
	goto L27
L125:
	;
	v459 = v440
	goto L22
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L8
	} else {
		goto L597
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L8
	} else {
		goto L594
	}
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L8
	} else {
		goto L591
	}
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2411 = m.ExcPending
	if v2411 != 0 {
		goto L8
	} else {
		goto L588
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L8
	} else {
		goto L585
	}
L131:
	;
	if v65 != 0 {
		goto L581
	} else {
		goto L582
	}
L132:
	;
	v2367 = F_create_indexscan_plan(m, l0, l1, v459, v40, int32(0))
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L8
	} else {
		goto L580
	}
L133:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2355 = m.ExcPending
	if v2355 != 0 {
		goto L8
	} else {
		goto L577
	}
L134:
	;
	v2260 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v2261 == int32(0) {
		goto L561
	} else {
		goto L562
	}
L135:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v2079)+68))
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v2082 != 0 {
		goto L510
	} else {
		goto L511
	}
L136:
	;
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v1984)+68))
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1986 != 0 {
		goto L490
	} else {
		goto L491
	}
L137:
	;
	v1949 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L8
	} else {
		goto L482
	}
L138:
	;
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1897)+68))
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1899 != 0 {
		goto L472
	} else {
		goto L473
	}
L139:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+68))
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1667 != 0 {
		goto L424
	} else {
		goto L425
	}
L140:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1610)+68))
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1612 != 0 {
		goto L412
	} else {
		goto L413
	}
L141:
	;
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1555)+68))
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1557 != 0 {
		goto L400
	} else {
		goto L401
	}
L142:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v1498)+68))
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1500 != 0 {
		goto L388
	} else {
		goto L389
	}
L143:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+68))
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+140))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1163 = F_create_plan(m, v1161, v1162)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L8
	} else {
		goto L316
	}
L144:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+68))
	v1027 = int32(0)
	if v40 == v1027 {
		v1107 = v1027
		goto L284
	} else {
		goto L285
	}
L145:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v793)+68))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v795 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L146:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)+68))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v571 = F_create_bitmap_subplan(m, l0, v564, v19+int32(172), v19+int32(168), v19+int32(164))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L8
	} else {
		goto L170
	}
L147:
	;
	v560 = F_create_indexscan_plan(m, l0, l1, v459, v40, int32(1))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L8
	} else {
		goto L169
	}
L148:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)+68))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v506 != 0 {
		goto L158
	} else {
		goto L159
	}
L149:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)+68))
	v470 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L8
	} else {
		goto L150
	}
L150:
	;
	v473 = F_extract_actual_clauses(m, v470, int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L8
	} else {
		goto L151
	}
L151:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v475 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v476 = F_replace_nestloop_params_mutator(m, v473, l0)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L8
	} else {
		goto L155
	}
L153:
	;
	v478 = v473
	goto L154
L154:
	;
	v480 = F_palloc0(m, int32(80))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L8
	} else {
		goto L156
	}
L155:
	;
	v478 = v476
	goto L154
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v480)+72)) = v469
	*(*int64)(unsafe.Add(mBase, uint32(v480)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v480)+48)) = v478
	*(*int32)(unsafe.Add(mBase, uint32(v480)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v480))) = int32(339)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v480)+4)) = v489
	v491 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v480)+8)) = v491
	v493 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v480)+16)) = v493
	v495 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v480)+24)) = v495
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v480)+32)) = v498
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v480)+36)) = uint8(v500)
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v480)+37)) = uint8(v502)
	v2372 = v480
	goto L131
L157:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)+32))
	v521 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L8
	} else {
		goto L161
	}
L158:
	;
	v518 = v506 + v505<<(uint(int32(2))%32)
	goto L157
L159:
	;
	goto L160
L160:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v510)+52))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)+12))
	v518 = v512 + v505<<(uint(int32(2))%32) - int32(4)
	goto L157
L161:
	;
	v524 = F_extract_actual_clauses(m, v521, int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L8
	} else {
		goto L162
	}
L162:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v526 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v527 = F_replace_nestloop_params_mutator(m, v524, l0)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L8
	} else {
		goto L166
	}
L164:
	;
	v531 = v520
	v532 = v524
	goto L165
L165:
	;
	v534 = F_palloc0(m, int32(88))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L8
	} else {
		goto L168
	}
L166:
	;
	v529 = F_replace_nestloop_params_mutator(m, v520, l0)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L8
	} else {
		goto L167
	}
L167:
	;
	v531 = v529
	v532 = v527
	goto L165
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v534)+80)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v534)+72)) = v505
	*(*int64)(unsafe.Add(mBase, uint32(v534)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v534)+48)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v534)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v534))) = int32(340)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v534)+4)) = v544
	v546 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+8)) = v546
	v548 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+16)) = v548
	v550 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+24)) = v550
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v534)+32)) = v553
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v534)+36)) = uint8(v555)
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v534)+37)) = uint8(v557)
	v2372 = v534
	goto L131
L169:
	;
	v2372 = v560
	goto L131
L170:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v573 == int32(1) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v579 = v571
	goto L175
L172:
	;
	goto L173
L173:
	;
	v638 = int32(0)
	if v40 == v638 {
		v737 = v638
		goto L183
	} else {
		goto L184
	}
L174:
	;
	v620 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v579)+84)) = uint8(v620)
	goto L173
L175:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	switch v593 - int32(337) {
	case 0:
		v599 = int32(72)
		goto L178
	case 1:
		goto L179
	default:
		goto L177
	case 6:
		goto L174
	}
L176:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L8
	} else {
		goto L180
	}
L177:
	;
	goto L176
L178:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v599+v579)))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)+12))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v602)))
	v579 = v603
	goto L175
L179:
	;
	v596 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v579)+72)) = uint8(v596)
	v599 = int32(76)
	goto L178
L180:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v608
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_3), v19+int32(16))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L8
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(_a_F_create_scan_plan_5), int32(_a_F_create_scan_plan_6))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L8
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	v751 = F_order_qual_clauses(m, l0, v737)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L8
	} else {
		goto L218
	}
L184:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v641 <= int32(0) {
		v737 = v638
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v19)+164))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v19)+168))
	v649 = v638
	v650 = int32(0)
	goto L186
L186:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v663+v650<<(uint(int32(2))%32))))
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+10)))
	if v668 != 0 {
		v728 = v649
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v737 = v728
	goto L183
L188:
	;
	v732 = v650 + int32(1)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v732 < v733 {
		v649 = v728
		v650 = v732
		goto L186
	} else {
		goto L217
	}
L189:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v667)+4))
	v670 = F_list_member(m, v645, v669)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L8
	} else {
		goto L190
	}
L190:
	;
	if v670 != 0 {
		v728 = v649
		goto L188
	} else {
		goto L191
	}
L191:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v667)+60))
	if v672 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v673 = int32(0)
	if v644 == v673 {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	goto L194
L194:
	;
	v712 = F_contain_mutable_functions(m, v669)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L8
	} else {
		goto L209
	}
L195:
	;
	if v711 != 0 {
		v728 = v649
		goto L188
	} else {
		goto L208
	}
L196:
	;
	v711 = int32(0)
	goto L195
L197:
	;
	goto L198
L198:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v644)+4))
	if v679 <= int32(0) {
		v705 = v673
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v711 = v705
	goto L195
L200:
	;
	v682 = int32(0)
	if v682 < v679 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v685 = v679
	goto L203
L202:
	;
	v685 = v682
	goto L203
L203:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v644)+12))
	v688 = int32(0)
	goto L204
L204:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v686+v688<<(uint(int32(2))%32))))
	v697 = base.B2i32(v696 == v672)
	if v696 == v672 {
		v705 = v697
		goto L199
	} else {
		goto L206
	}
L205:
	;
	v705 = v697
	goto L199
L206:
	;
	v699 = v688 + int32(1)
	if v699 != v685 {
		v688 = v699
		goto L204
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	goto L194
L209:
	;
	if v712 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v669
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = v669
	v721 = F_list_make1_impl(m, int32(1), v19+int32(24))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L8
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v726 = F_lappend(m, v649, v667)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L8
	} else {
		goto L216
	}
L213:
	;
	v724 = F_predicate_implied_by(m, v721, v645, int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L8
	} else {
		goto L214
	}
L214:
	;
	if v724 != 0 {
		v728 = v649
		goto L188
	} else {
		goto L215
	}
L215:
	;
	goto L212
L216:
	;
	v728 = v726
	goto L188
L217:
	;
	goto L187
L218:
	;
	v754 = F_extract_actual_clauses(m, v751, int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L8
	} else {
		goto L219
	}
L219:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v757 = F_list_difference_ptr(m, v756, v754)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L8
	} else {
		goto L220
	}
L220:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v759 != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v760 = F_replace_nestloop_params_mutator(m, v754, l0)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L8
	} else {
		goto L224
	}
L222:
	;
	v764 = v754
	v765 = v757
	goto L223
L223:
	;
	v767 = F_palloc0(m, int32(88))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L8
	} else {
		goto L226
	}
L224:
	;
	v762 = F_replace_nestloop_params_mutator(m, v757, l0)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L8
	} else {
		goto L225
	}
L225:
	;
	v764 = v760
	v765 = v762
	goto L223
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v767)+80)) = v765
	*(*int32)(unsafe.Add(mBase, uint32(v767)+72)) = v563
	*(*int32)(unsafe.Add(mBase, uint32(v767)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v767)+52)) = v571
	*(*int32)(unsafe.Add(mBase, uint32(v767)+48)) = v764
	*(*int32)(unsafe.Add(mBase, uint32(v767)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v767))) = int32(344)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v767)+4)) = v778
	v780 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v767)+8)) = v780
	v782 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v767)+16)) = v782
	v784 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v767)+24)) = v784
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v786)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v767)+32)) = v787
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v767)+36)) = uint8(v789)
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v767)+37)) = uint8(v791)
	v2372 = v767
	goto L131
L227:
	;
	v965 = F_order_qual_clauses(m, l0, v951)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L8
	} else {
		goto L269
	}
L228:
	;
	v951 = v40
	goto L227
L229:
	;
	goto L230
L230:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	if v798 != int32(1) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v951 = v40
	goto L227
L232:
	;
	goto L233
L233:
	;
	if v40 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v951 = int32(0)
	goto L227
L235:
	;
	goto L236
L236:
	;
	v804 = int32(0)
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v805 <= v804 {
		v951 = v804
		goto L227
	} else {
		goto L237
	}
L237:
	;
	v811 = v804
	v812 = int32(0)
	goto L238
L238:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v825+v812<<(uint(int32(2))%32))))
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829)+10)))
	if v830 != 0 {
		v931 = v811
		goto L240
	} else {
		goto L241
	}
L239:
	;
	v951 = v931
	goto L227
L240:
	;
	v946 = v812 + int32(1)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v946 < v947 {
		v811 = v931
		v812 = v946
		goto L238
	} else {
		goto L268
	}
L241:
	;
	v831 = int32(0)
	if v795 == v831 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	if v869 != 0 {
		v931 = v811
		goto L240
	} else {
		goto L255
	}
L243:
	;
	v869 = int32(0)
	goto L242
L244:
	;
	goto L245
L245:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	if v837 <= int32(0) {
		v863 = v831
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v869 = v863
	goto L242
L247:
	;
	v840 = int32(0)
	if v840 < v837 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v843 = v837
	goto L250
L249:
	;
	v843 = v840
	goto L250
L250:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v795)+12))
	v846 = int32(0)
	goto L251
L251:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v844+v846<<(uint(int32(2))%32))))
	v855 = base.B2i32(v854 == v829)
	if v854 == v829 {
		v863 = v855
		goto L246
	} else {
		goto L253
	}
L252:
	;
	v863 = v855
	goto L246
L253:
	;
	v857 = v846 + int32(1)
	if v857 != v843 {
		v846 = v857
		goto L251
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v870 = int32(0)
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v829)+60))
	if base.B2i32(v871 == v870)|base.B2i32(v795 == v870) != 0 {
		v919 = v870
		goto L256
	} else {
		goto L257
	}
L256:
	;
	if v919 != 0 {
		v931 = v811
		goto L240
	} else {
		goto L266
	}
L257:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	if v877 <= int32(0) {
		v919 = v870
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v880 = int32(0)
	if v880 < v877 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v883 = v877
	goto L261
L260:
	;
	v883 = v880
	goto L261
L261:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v795)+12))
	v890 = int32(0)
	goto L262
L262:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v884+v890<<(uint(int32(2))%32))))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v905)+60))
	v907 = base.B2i32(v906 == v871)
	if v906 == v871 {
		v919 = v907
		goto L256
	} else {
		goto L264
	}
L263:
	;
	v919 = v907
	goto L256
L264:
	;
	v909 = v890 + int32(1)
	if v909 != v883 {
		v890 = v909
		goto L262
	} else {
		goto L265
	}
L265:
	;
	goto L263
L266:
	;
	v927 = F_lappend(m, v811, v829)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L8
	} else {
		goto L267
	}
L267:
	;
	v931 = v927
	goto L240
L268:
	;
	goto L239
L269:
	;
	v968 = F_extract_actual_clauses(m, v795, int32(0))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L8
	} else {
		goto L270
	}
L270:
	;
	v971 = F_extract_actual_clauses(m, v965, int32(0))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L8
	} else {
		goto L271
	}
L271:
	;
	if v968 == int32(0) {
		v989 = v971
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v991 != 0 {
		goto L278
	} else {
		goto L279
	}
L273:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v968)+4))
	if v975 < int32(2) {
		v989 = v971
		goto L272
	} else {
		goto L274
	}
L274:
	;
	v978 = F_make_orclause(m, v968)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L8
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v978
	*(*int32)(unsafe.Add(mBase, uint32(v19)+172)) = v978
	v985 = F_list_make1_impl(m, int32(1), v19+int32(28))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L8
	} else {
		goto L276
	}
L276:
	;
	v987 = F_list_difference(m, v971, v985)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L8
	} else {
		goto L277
	}
L277:
	;
	v989 = v987
	goto L272
L278:
	;
	v992 = F_replace_nestloop_params_mutator(m, v968, l0)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L8
	} else {
		goto L281
	}
L279:
	;
	v996 = v968
	v997 = v989
	goto L280
L280:
	;
	v999 = F_palloc0(m, int32(88))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L8
	} else {
		goto L283
	}
L281:
	;
	v994 = F_replace_nestloop_params_mutator(m, v989, l0)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L8
	} else {
		goto L282
	}
L282:
	;
	v996 = v992
	v997 = v994
	goto L280
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v999)+80)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v999)+72)) = v794
	*(*int64)(unsafe.Add(mBase, uint32(v999)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v999)+48)) = v997
	*(*int32)(unsafe.Add(mBase, uint32(v999)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v999))) = int32(345)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v999)+4)) = v1009
	v1011 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v999)+8)) = v1011
	v1013 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v999)+16)) = v1013
	v1015 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v999)+24)) = v1015
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1017)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v999)+32)) = v1018
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+36)) = uint8(v1020)
	v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+37)) = uint8(v1022)
	v2372 = v999
	goto L131
L284:
	;
	v1118 = F_order_qual_clauses(m, l0, v1107)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L8
	} else {
		goto L307
	}
L285:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v1030 <= int32(0) {
		v1107 = v1027
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v1037 = int32(0)
	v1039 = v1027
	goto L287
L287:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1050+v1037<<(uint(int32(2))%32))))
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1054)+10)))
	if v1055 != 0 {
		v1097 = v1039
		goto L289
	} else {
		goto L290
	}
L288:
	;
	v1107 = v1097
	goto L284
L289:
	;
	v1099 = v1037 + int32(1)
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v1099 < v1100 {
		v1037 = v1099
		v1039 = v1097
		goto L287
	} else {
		goto L306
	}
L290:
	;
	v1056 = int32(0)
	if v1024 == v1056 {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	if v1094 != 0 {
		v1097 = v1039
		goto L289
	} else {
		goto L304
	}
L292:
	;
	v1094 = int32(0)
	goto L291
L293:
	;
	goto L294
L294:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+4))
	if v1062 <= int32(0) {
		v1088 = v1056
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1094 = v1088
	goto L291
L296:
	;
	v1065 = int32(0)
	if v1065 < v1062 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1068 = v1062
	goto L299
L298:
	;
	v1068 = v1065
	goto L299
L299:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+12))
	v1071 = int32(0)
	goto L300
L300:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1069+v1071<<(uint(int32(2))%32))))
	v1080 = base.B2i32(v1079 == v1054)
	if v1079 == v1054 {
		v1088 = v1080
		goto L295
	} else {
		goto L302
	}
L301:
	;
	v1088 = v1080
	goto L295
L302:
	;
	v1082 = v1071 + int32(1)
	if v1082 != v1068 {
		v1071 = v1082
		goto L300
	} else {
		goto L303
	}
L303:
	;
	goto L301
L304:
	;
	v1095 = F_lappend(m, v1039, v1054)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L8
	} else {
		goto L305
	}
L305:
	;
	v1097 = v1095
	goto L289
L306:
	;
	goto L288
L307:
	;
	v1121 = F_extract_actual_clauses(m, v1024, int32(0))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L8
	} else {
		goto L308
	}
L308:
	;
	v1124 = F_extract_actual_clauses(m, v1118, int32(0))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L8
	} else {
		goto L309
	}
L309:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1126 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1127 = F_replace_nestloop_params_mutator(m, v1121, l0)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L8
	} else {
		goto L313
	}
L311:
	;
	v1131 = v1121
	v1132 = v1124
	goto L312
L312:
	;
	v1134 = F_palloc0(m, int32(88))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L8
	} else {
		goto L315
	}
L313:
	;
	v1129 = F_replace_nestloop_params_mutator(m, v1124, l0)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L8
	} else {
		goto L314
	}
L314:
	;
	v1131 = v1127
	v1132 = v1129
	goto L312
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+80)) = v1131
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+72)) = v1026
	*(*int64)(unsafe.Add(mBase, uint32(v1134)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+48)) = v1132
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1134))) = int32(346)
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+4)) = v1144
	v1146 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1134)+8)) = v1146
	v1148 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1134)+16)) = v1148
	v1150 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1134)+24)) = v1150
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1152)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+32)) = v1153
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1134)+36)) = uint8(v1155)
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1134)+37)) = uint8(v1157)
	v2372 = v1134
	goto L131
L316:
	;
	v1165 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L8
	} else {
		goto L317
	}
L317:
	;
	v1168 = F_extract_actual_clauses(m, v1165, int32(0))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L8
	} else {
		goto L318
	}
L318:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1170 != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1171 = int32(0)
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+144))
	if v1172 == v1171 {
		goto L322
	} else {
		goto L323
	}
L320:
	;
	v1456 = v1168
	goto L321
L321:
	;
	v1471 = F_palloc0(m, int32(88))
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L8
	} else {
		goto L386
	}
L322:
	;
	v1452 = F_replace_nestloop_params_mutator(m, v1168, l0)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L8
	} else {
		goto L385
	}
L323:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+4))
	if v1175 <= int32(0) {
		goto L322
	} else {
		goto L324
	}
L324:
	;
	v1186 = v1171
	goto L327
L325:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L8
	} else {
		goto L382
	}
L326:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L8
	} else {
		goto L379
	}
L327:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+12))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1194+v1186<<(uint(int32(2))%32))))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+4))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1199)))
	if v1200 != int32(6) {
		goto L332
	} else {
		goto L333
	}
L328:
	;
	goto L322
L329:
	;
	v1407 = v1186 + int32(1)
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+4))
	if v1407 < v1408 {
		v1186 = v1407
		goto L327
	} else {
		goto L378
	}
L330:
	;
	v1377 = F_palloc0(m, int32(12))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L8
	} else {
		goto L375
	}
L331:
	;
	v1262 = F_find_placeholder_info(m, l0, v1199)
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L8
	} else {
		goto L350
	}
L332:
	;
	if v1200 == int32(319) {
		goto L331
	} else {
		goto L335
	}
L333:
	;
	goto L334
L334:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+4))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	v1220 = F_bms_is_member(m, v1218, v1219)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L8
	} else {
		goto L339
	}
L335:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L8
	} else {
		goto L336
	}
L336:
	;
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_7), int32(0))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L8
	} else {
		goto L337
	}
L337:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_8), int32(597), int32(_a_F_create_scan_plan_9))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L8
	} else {
		goto L338
	}
L338:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L339:
	;
	if v1220 == int32(0) {
		goto L326
	} else {
		goto L340
	}
L340:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v1224 == int32(0) {
		goto L330
	} else {
		goto L341
	}
L341:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+4))
	if v1227 <= int32(0) {
		goto L330
	} else {
		goto L342
	}
L342:
	;
	v1230 = int32(0)
	if v1230 < v1227 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1233 = v1227
	goto L345
L344:
	;
	v1233 = v1230
	goto L345
L345:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+8))
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+12))
	v1240 = int32(0)
	goto L346
L346:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1235+v1240<<(uint(int32(2))%32))))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+4))
	if v1257 == v1234 {
		goto L329
	} else {
		goto L348
	}
L347:
	;
	goto L330
L348:
	;
	v1260 = v1240 + int32(1)
	if v1233 != v1260 {
		v1240 = v1260
		goto L346
	} else {
		goto L349
	}
L349:
	;
	goto L347
L350:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1262)+12))
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	v1266 = int32(0)
	if v1264 == v1266 {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	if v1319 == int32(0) {
		goto L325
	} else {
		goto L365
	}
L352:
	;
	v1319 = int32(1)
	goto L351
L353:
	;
	goto L354
L354:
	;
	if v1265 == int32(0) {
		v1312 = v1266
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1319 = v1312
	goto L351
L356:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+4))
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+4))
	if v1276 < v1275 {
		v1312 = v1266
		goto L355
	} else {
		goto L357
	}
L357:
	;
	v1278 = int32(1)
	if v1275 <= v1278 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1281 = v1278
	goto L360
L359:
	;
	v1281 = v1275
	goto L360
L360:
	;
	v1282 = int32(8)
	v1287 = int32(0)
	goto L361
L361:
	;
	v1294 = v1287 << (uint(int32(2)) % 32)
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1264+v1282+v1294)))
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1265+v1282+v1294)))
	v1301 = v1296 & (v1298 ^ int32(-1))
	v1303 = base.B2i32(v1301 == int32(0))
	if v1301 != 0 {
		v1312 = v1303
		goto L355
	} else {
		goto L363
	}
L362:
	;
	v1312 = v1303
	goto L355
L363:
	;
	v1305 = v1287 + int32(1)
	if v1305 != v1281 {
		v1287 = v1305
		goto L361
	} else {
		goto L364
	}
L364:
	;
	goto L362
L365:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v1322 == int32(0) {
		goto L330
	} else {
		goto L366
	}
L366:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1322)+4))
	if v1325 <= int32(0) {
		goto L330
	} else {
		goto L367
	}
L367:
	;
	v1328 = int32(0)
	if v1328 < v1325 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1331 = v1325
	goto L370
L369:
	;
	v1331 = v1328
	goto L370
L370:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+8))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1322)+12))
	v1338 = int32(0)
	goto L371
L371:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1333+v1338<<(uint(int32(2))%32))))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1354)+4))
	if v1355 == v1332 {
		goto L329
	} else {
		goto L373
	}
L372:
	;
	goto L330
L373:
	;
	v1358 = v1338 + int32(1)
	if v1358 != v1331 {
		v1338 = v1358
		goto L371
	} else {
		goto L374
	}
L374:
	;
	goto L372
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377))) = int32(357)
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+4)) = v1381
	v1383 = F_copyObjectImpl(m, v1199)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L8
	} else {
		goto L376
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+8)) = v1383
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	v1387 = F_lappend(m, v1386, v1377)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L8
	} else {
		goto L377
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+356)) = v1387
	goto L329
L378:
	;
	goto L328
L379:
	;
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_10), int32(0))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L8
	} else {
		goto L380
	}
L380:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_8), int32(543), int32(_a_F_create_scan_plan_9))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L8
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
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_10), int32(0))
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L8
	} else {
		goto L383
	}
L383:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_8), int32(574), int32(_a_F_create_scan_plan_9))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L8
	} else {
		goto L384
	}
L384:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L385:
	;
	v1456 = v1452
	goto L321
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+84)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+80)) = v1163
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+72)) = v1160
	*(*int64)(unsafe.Add(mBase, uint32(v1471)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+48)) = v1456
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1471))) = int32(347)
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+4)) = v1483
	v1485 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1471)+8)) = v1485
	v1487 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1471)+16)) = v1487
	v1489 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1471)+24)) = v1489
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1491)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+32)) = v1492
	v1494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1471)+36)) = uint8(v1494)
	v1496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1471)+37)) = uint8(v1496)
	v2372 = v1471
	goto L131
L387:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1512)))
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1513)+68))
	v1515 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L8
	} else {
		goto L391
	}
L388:
	;
	v1512 = v1500 + v1499<<(uint(int32(2))%32)
	goto L387
L389:
	;
	goto L390
L390:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1504)+52))
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1505)+12))
	v1512 = v1506 + v1499<<(uint(int32(2))%32) - int32(4)
	goto L387
L391:
	;
	v1518 = F_extract_actual_clauses(m, v1515, int32(0))
	mBase = m.M
	v1519 = m.ExcPending
	if v1519 != 0 {
		goto L8
	} else {
		goto L392
	}
L392:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1520 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1521 = F_replace_nestloop_params_mutator(m, v1518, l0)
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L8
	} else {
		goto L396
	}
L394:
	;
	v1525 = v1514
	v1526 = v1518
	goto L395
L395:
	;
	v1527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1513)+72)))
	v1529 = F_palloc0(m, int32(88))
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L8
	} else {
		goto L398
	}
L396:
	;
	v1523 = F_replace_nestloop_params_mutator(m, v1514, l0)
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		goto L8
	} else {
		goto L397
	}
L397:
	;
	v1525 = v1523
	v1526 = v1521
	goto L395
L398:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1529)+84)) = uint8(v1527)
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+80)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+72)) = v1499
	*(*int64)(unsafe.Add(mBase, uint32(v1529)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+48)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1529))) = int32(348)
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+4)) = v1540
	v1542 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1529)+8)) = v1542
	v1544 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1529)+16)) = v1544
	v1546 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1529)+24)) = v1546
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v1548)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+32)) = v1549
	v1551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1529)+36)) = uint8(v1551)
	v1553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1529)+37)) = uint8(v1553)
	v2372 = v1529
	goto L131
L399:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v1569)))
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1570)+76))
	v1572 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L8
	} else {
		goto L403
	}
L400:
	;
	v1569 = v1557 + v1556<<(uint(int32(2))%32)
	goto L399
L401:
	;
	goto L402
L402:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+52))
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v1562)+12))
	v1569 = v1563 + v1556<<(uint(int32(2))%32) - int32(4)
	goto L399
L403:
	;
	v1575 = F_extract_actual_clauses(m, v1572, int32(0))
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L8
	} else {
		goto L404
	}
L404:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1577 != 0 {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v1578 = F_replace_nestloop_params_mutator(m, v1575, l0)
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L8
	} else {
		goto L408
	}
L406:
	;
	v1582 = v1571
	v1583 = v1575
	goto L407
L407:
	;
	v1585 = F_palloc0(m, int32(88))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L8
	} else {
		goto L410
	}
L408:
	;
	v1580 = F_replace_nestloop_params_mutator(m, v1571, l0)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L8
	} else {
		goto L409
	}
L409:
	;
	v1582 = v1580
	v1583 = v1578
	goto L407
L410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1585)+80)) = v1582
	*(*int32)(unsafe.Add(mBase, uint32(v1585)+72)) = v1556
	*(*int64)(unsafe.Add(mBase, uint32(v1585)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1585)+48)) = v1583
	*(*int32)(unsafe.Add(mBase, uint32(v1585)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1585))) = int32(350)
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1585)+4)) = v1595
	v1597 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1585)+8)) = v1597
	v1599 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1585)+16)) = v1599
	v1601 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1585)+24)) = v1601
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1603)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1585)+32)) = v1604
	v1606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1585)+36)) = uint8(v1606)
	v1608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1585)+37)) = uint8(v1608)
	v2372 = v1585
	goto L131
L411:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1624)))
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+80))
	v1627 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L8
	} else {
		goto L415
	}
L412:
	;
	v1624 = v1612 + v1611<<(uint(int32(2))%32)
	goto L411
L413:
	;
	goto L414
L414:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1616)+52))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1617)+12))
	v1624 = v1618 + v1611<<(uint(int32(2))%32) - int32(4)
	goto L411
L415:
	;
	v1630 = F_extract_actual_clauses(m, v1627, int32(0))
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L8
	} else {
		goto L416
	}
L416:
	;
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1632 != 0 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v1633 = F_replace_nestloop_params_mutator(m, v1630, l0)
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		goto L8
	} else {
		goto L420
	}
L418:
	;
	v1637 = v1626
	v1638 = v1630
	goto L419
L419:
	;
	v1640 = F_palloc0(m, int32(88))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L8
	} else {
		goto L422
	}
L420:
	;
	v1635 = F_replace_nestloop_params_mutator(m, v1626, l0)
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L8
	} else {
		goto L421
	}
L421:
	;
	v1637 = v1635
	v1638 = v1633
	goto L419
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1640)+80)) = v1637
	*(*int32)(unsafe.Add(mBase, uint32(v1640)+72)) = v1611
	*(*int64)(unsafe.Add(mBase, uint32(v1640)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1640)+48)) = v1638
	*(*int32)(unsafe.Add(mBase, uint32(v1640)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1640))) = int32(349)
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1640)+4)) = v1650
	v1652 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1640)+8)) = v1652
	v1654 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1640)+16)) = v1654
	v1656 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1640)+24)) = v1656
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1658)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1640)+32)) = v1659
	v1661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1640)+36)) = uint8(v1661)
	v1663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1640)+37)) = uint8(v1663)
	v2372 = v1640
	goto L131
L423:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1679)))
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+88))
	v1684 = l0
	v1685 = v1681
	goto L428
L424:
	;
	v1679 = v1667 + v1666<<(uint(int32(2))%32)
	goto L423
L425:
	;
	goto L426
L426:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1671)+52))
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v1672)+12))
	v1679 = v1673 + v1666<<(uint(int32(2))%32) - int32(4)
	goto L423
L427:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+4))
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+48))
	if v1720 == int32(0) {
		goto L436
	} else {
		goto L437
	}
L428:
	;
	if v1685 == int32(0) {
		goto L427
	} else {
		goto L430
	}
L429:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L8
	} else {
		goto L432
	}
L430:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+16))
	if v1702 != 0 {
		v1684 = v1702
		v1685 = v1685 - int32(1)
		goto L428
	} else {
		goto L431
	}
L431:
	;
	goto L429
L432:
	;
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v1707
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_11), v19+int32(96))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L8
	} else {
		goto L433
	}
L433:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(3912), int32(_a_F_create_scan_plan_12))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L8
	} else {
		goto L434
	}
L434:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L435:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+76))
	if v1813 == int32(0) {
		goto L130
	} else {
		goto L453
	}
L436:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L8
	} else {
		goto L450
	}
L437:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1720)+4))
	if v1723 <= int32(0) {
		goto L436
	} else {
		goto L438
	}
L438:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+84))
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v1720)+12))
	v1732 = int32(0)
	goto L439
L439:
	;
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1727+v1732<<(uint(int32(2))%32))))
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+4))
	v1752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1749))))
	v1755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726))))
	if base.B2i32(v1752 == int32(0))|base.B2i32(v1752 != v1755) != 0 {
		v1773 = v1752
		v1774 = v1755
		goto L442
	} else {
		goto L443
	}
L440:
	;
	goto L436
L441:
	;
	if v1773-v1774 == int32(0) {
		goto L435
	} else {
		goto L448
	}
L442:
	;
	goto L441
L443:
	;
	v1758 = v1749
	v1759 = v1726
	goto L444
L444:
	;
	v1762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1759)+1)))
	v1763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1758)+1)))
	if v1763 == int32(0) {
		v1773 = v1763
		v1774 = v1762
		goto L442
	} else {
		goto L446
	}
L445:
	;
	v1773 = v1763
	v1774 = v1762
	goto L442
L446:
	;
	v1766 = int32(1)
	if v1763 == v1762 {
		v1758 = v1758 + v1766
		v1759 = v1759 + v1766
		goto L444
	} else {
		goto L447
	}
L447:
	;
	goto L445
L448:
	;
	v1779 = v1732 + int32(1)
	if v1723 != v1779 {
		v1732 = v1779
		goto L439
	} else {
		goto L449
	}
L449:
	;
	goto L440
L450:
	;
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v1801
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_13), v19+int32(32))
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L8
	} else {
		goto L451
	}
L451:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(3930), int32(_a_F_create_scan_plan_12))
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L8
	} else {
		goto L452
	}
L452:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L453:
	;
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1813)+4))
	if v1816 <= v1732 {
		goto L130
	} else {
		goto L454
	}
L454:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1813)+12))
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1818+v1732<<(uint(int32(2))%32))))
	if v1822 <= int32(0) {
		goto L129
	} else {
		goto L455
	}
L455:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+72))
	if v1825 == int32(0) {
		goto L128
	} else {
		goto L456
	}
L456:
	;
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v1825)+4))
	if v1828 <= int32(0) {
		goto L128
	} else {
		goto L457
	}
L457:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1825)+12))
	v1836 = int32(0)
	goto L458
L458:
	;
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1831+v1836<<(uint(int32(2))%32))))
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1852)+16))
	if v1822 != v1853 {
		goto L460
	} else {
		goto L461
	}
L459:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1852)+40))
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1858)+12))
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1859)))
	v1861 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L8
	} else {
		goto L464
	}
L460:
	;
	v1856 = v1836 + int32(1)
	if v1856 != v1828 {
		v1836 = v1856
		goto L458
	} else {
		goto L463
	}
L461:
	;
	goto L462
L462:
	;
	goto L459
L463:
	;
	goto L128
L464:
	;
	v1864 = F_extract_actual_clauses(m, v1861, int32(0))
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L8
	} else {
		goto L465
	}
L465:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1866 != 0 {
		goto L466
	} else {
		goto L467
	}
L466:
	;
	v1867 = F_replace_nestloop_params_mutator(m, v1864, l0)
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L8
	} else {
		goto L469
	}
L467:
	;
	v1869 = v1864
	goto L468
L468:
	;
	v1871 = F_palloc0(m, int32(88))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L8
	} else {
		goto L470
	}
L469:
	;
	v1869 = v1867
	goto L468
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+84)) = v1860
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+80)) = v1822
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+72)) = v1666
	*(*int64)(unsafe.Add(mBase, uint32(v1871)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+48)) = v1869
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1871))) = int32(351)
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+4)) = v1882
	v1884 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1871)+8)) = v1884
	v1886 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1871)+16)) = v1886
	v1888 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1871)+24)) = v1888
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1890)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+32)) = v1891
	v1893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1871)+36)) = uint8(v1893)
	v1895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1871)+37)) = uint8(v1895)
	v2372 = v1871
	goto L131
L471:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1911)))
	v1913 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L8
	} else {
		goto L475
	}
L472:
	;
	v1911 = v1899 + v1898<<(uint(int32(2))%32)
	goto L471
L473:
	;
	goto L474
L474:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1903)+52))
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v1904)+12))
	v1911 = v1905 + v1898<<(uint(int32(2))%32) - int32(4)
	goto L471
L475:
	;
	v1916 = F_extract_actual_clauses(m, v1913, int32(0))
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L8
	} else {
		goto L476
	}
L476:
	;
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1918 != 0 {
		goto L477
	} else {
		goto L478
	}
L477:
	;
	v1919 = F_replace_nestloop_params_mutator(m, v1916, l0)
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L8
	} else {
		goto L480
	}
L478:
	;
	v1921 = v1916
	goto L479
L479:
	;
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v1912)+108))
	v1924 = F_palloc0(m, int32(88))
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L8
	} else {
		goto L481
	}
L480:
	;
	v1921 = v1919
	goto L479
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1924)+80)) = v1922
	*(*int32)(unsafe.Add(mBase, uint32(v1924)+72)) = v1898
	*(*int64)(unsafe.Add(mBase, uint32(v1924)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1924)+48)) = v1921
	*(*int32)(unsafe.Add(mBase, uint32(v1924)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1924))) = int32(352)
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1924)+4)) = v1934
	v1936 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1924)+8)) = v1936
	v1938 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1924)+16)) = v1938
	v1940 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1924)+24)) = v1940
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v1942)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1924)+32)) = v1943
	v1945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1924)+36)) = uint8(v1945)
	v1947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1924)+37)) = uint8(v1947)
	v2372 = v1924
	goto L131
L482:
	;
	v1952 = F_extract_actual_clauses(m, v1949, int32(0))
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L8
	} else {
		goto L483
	}
L483:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1954 != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v1955 = F_replace_nestloop_params_mutator(m, v1952, l0)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L8
	} else {
		goto L487
	}
L485:
	;
	v1957 = v1952
	goto L486
L486:
	;
	v1959 = F_palloc0(m, int32(80))
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L8
	} else {
		goto L488
	}
L487:
	;
	v1957 = v1955
	goto L486
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+72)) = v1957
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1959)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v1959))) = int32(331)
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+4)) = v1969
	v1971 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1959)+8)) = v1971
	v1973 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1959)+16)) = v1973
	v1975 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1959)+24)) = v1975
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v1977)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+32)) = v1978
	v1980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1959)+36)) = uint8(v1980)
	v1982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1959)+37)) = uint8(v1982)
	v2372 = v1959
	goto L131
L489:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v1998)))
	v2000 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+88))
	if v2000 == int32(0) {
		goto L127
	} else {
		goto L493
	}
L490:
	;
	v1998 = v1986 + v1985<<(uint(int32(2))%32)
	goto L489
L491:
	;
	goto L492
L492:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+52))
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v1991)+12))
	v1998 = v1992 + v1985<<(uint(int32(2))%32) - int32(4)
	goto L489
L493:
	;
	v2005 = l0
	v2006 = v2000
	goto L495
L494:
	;
	v2040 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+344))
	if v2040 < int32(0) {
		goto L126
	} else {
		goto L502
	}
L495:
	;
	v2020 = v2006 - int32(1)
	if v2020 == int32(0) {
		goto L494
	} else {
		goto L497
	}
L496:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L8
	} else {
		goto L499
	}
L497:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+16))
	if v2023 != 0 {
		v2005 = v2023
		v2006 = v2020
		goto L495
	} else {
		goto L498
	}
L498:
	;
	goto L496
L499:
	;
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v2028
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_11), v19+int32(144))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L8
	} else {
		goto L500
	}
L500:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(4083), int32(_a_F_create_scan_plan_14))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L8
	} else {
		goto L501
	}
L501:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L502:
	;
	v2043 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L8
	} else {
		goto L503
	}
L503:
	;
	v2046 = F_extract_actual_clauses(m, v2043, int32(0))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L8
	} else {
		goto L504
	}
L504:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v2048 != 0 {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v2049 = F_replace_nestloop_params_mutator(m, v2046, l0)
	mBase = m.M
	v2050 = m.ExcPending
	if v2050 != 0 {
		goto L8
	} else {
		goto L508
	}
L506:
	;
	v2051 = v2046
	goto L507
L507:
	;
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+344))
	v2054 = F_palloc0(m, int32(88))
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L8
	} else {
		goto L509
	}
L508:
	;
	v2051 = v2049
	goto L507
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2054)+80)) = v2052
	*(*int32)(unsafe.Add(mBase, uint32(v2054)+72)) = v1985
	*(*int64)(unsafe.Add(mBase, uint32(v2054)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2054)+48)) = v2051
	*(*int32)(unsafe.Add(mBase, uint32(v2054)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v2054))) = int32(353)
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2054)+4)) = v2064
	v2066 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v2054)+8)) = v2066
	v2068 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v2054)+16)) = v2068
	v2070 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v2054)+24)) = v2070
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2072)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2054)+32)) = v2073
	v2075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2054)+36)) = uint8(v2075)
	v2077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2054)+37)) = uint8(v2077)
	v2372 = v2054
	goto L131
L510:
	;
	v2084 = F_create_plan_recurse(m, l0, v2082, int32(1))
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L8
	} else {
		goto L513
	}
L511:
	;
	v2086 = int32(0)
	goto L512
L512:
	;
	if v2080 != 0 {
		goto L514
	} else {
		goto L515
	}
L513:
	;
	v2086 = v2084
	goto L512
L514:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2087 != 0 {
		goto L518
	} else {
		goto L519
	}
L515:
	;
	v2104 = int32(0)
	goto L516
L516:
	;
	v2105 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v2106 = m.ExcPending
	if v2106 != 0 {
		goto L8
	} else {
		goto L521
	}
L517:
	;
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v2099)))
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v2100)+16))
	v2104 = v2101
	goto L516
L518:
	;
	v2099 = v2087 + v2080<<(uint(int32(2))%32)
	goto L517
L519:
	;
	goto L520
L520:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v2091)+52))
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v2092)+12))
	v2099 = v2093 + v2080<<(uint(int32(2))%32) - int32(4)
	goto L517
L521:
	;
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v2079)+168))
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v2107)+12))
	v2109 = m.T0[v2108].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, v2079, v2104, l1, v459, v2105, v2086)
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L8
	} else {
		goto L522
	}
L522:
	;
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+4)) = v2111
	v2113 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v2109)+8)) = v2113
	v2115 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v2109)+16)) = v2115
	v2117 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v2109)+24)) = v2117
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v2119)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+32)) = v2120
	v2122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2109)+36)) = uint8(v2122)
	v2124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2109)+37)) = uint8(v2124)
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(v2079)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+88)) = v2126
	v2128 = *(*int32)(unsafe.Add(mBase, uint32(v2079)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+92)) = v2128
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v2079)+4))
	if v2130 == int32(4) {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	v2138 = l0 + int32(52)
	goto L525
L524:
	;
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2138 = v2135 + int32(8)
	goto L525
L525:
	;
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v2138)))
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+112)) = v2139
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2142 = F_bms_difference(m, v2139, v2141)
	mBase = m.M
	v2143 = m.ExcPending
	if v2143 != 0 {
		goto L8
	} else {
		goto L526
	}
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+116)) = v2142
	v2145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2079)+164)))
	if v2145 == int32(1) {
		goto L527
	} else {
		goto L528
	}
L527:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2149 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2148)+81)) = uint8(v2149)
	goto L529
L528:
	;
	goto L529
L529:
	;
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v2151 != 0 {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+48))
	v2153 = F_replace_nestloop_params_mutator(m, v2152, l0)
	mBase = m.M
	v2154 = m.ExcPending
	if v2154 != 0 {
		goto L8
	} else {
		goto L533
	}
L531:
	;
	goto L532
L532:
	;
	v2164 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2109)+120)) = uint8(v2164)
	if v2080 == v2164 {
		v2372 = v2109
		goto L131
	} else {
		goto L536
	}
L533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+48)) = v2153
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+96))
	v2157 = F_replace_nestloop_params_mutator(m, v2156, l0)
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L8
	} else {
		goto L534
	}
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+96)) = v2157
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+108))
	v2161 = F_replace_nestloop_params_mutator(m, v2160, l0)
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L8
	} else {
		goto L535
	}
L535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+108)) = v2161
	goto L532
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+172)) = int32(0)
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v2079)+28))
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v2170)+4))
	F_pull_varattnos(m, v2171, v2080, v19+int32(172))
	mBase = m.M
	v2175 = m.ExcPending
	if v2175 != 0 {
		goto L8
	} else {
		goto L537
	}
L537:
	;
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v2079)+184))
	if v2176 == int32(0) {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2231 = F_bms_is_member(m, int32(1), v2230)
	mBase = m.M
	v2232 = m.ExcPending
	if v2232 != 0 {
		goto L8
	} else {
		goto L547
	}
L539:
	;
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v2176)+4))
	if v2179 <= int32(0) {
		goto L538
	} else {
		goto L540
	}
L540:
	;
	v2185 = int32(0)
	goto L541
L541:
	;
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v2176)+12))
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v2199+v2185<<(uint(int32(2))%32))))
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(v2203)+4))
	F_pull_varattnos(m, v2204, v2080, v19+int32(172))
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L8
	} else {
		goto L543
	}
L542:
	;
	goto L538
L543:
	;
	v2210 = v2185 + int32(1)
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(v2176)+4))
	if v2210 < v2211 {
		v2185 = v2210
		goto L541
	} else {
		goto L544
	}
L544:
	;
	goto L542
L545:
	;
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	F_bms_free(m, v2257)
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L8
	} else {
		goto L559
	}
L546:
	;
	v2255 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2109)+120)) = uint8(v2255)
	goto L545
L547:
	;
	if v2231 != 0 {
		goto L546
	} else {
		goto L548
	}
L548:
	;
	v2234 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2235 = F_bms_is_member(m, int32(2), v2234)
	mBase = m.M
	v2236 = m.ExcPending
	if v2236 != 0 {
		goto L8
	} else {
		goto L549
	}
L549:
	;
	if v2235 != 0 {
		goto L546
	} else {
		goto L550
	}
L550:
	;
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2239 = F_bms_is_member(m, int32(3), v2238)
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L8
	} else {
		goto L551
	}
L551:
	;
	if v2239 != 0 {
		goto L546
	} else {
		goto L552
	}
L552:
	;
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2243 = F_bms_is_member(m, int32(4), v2242)
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		goto L8
	} else {
		goto L553
	}
L553:
	;
	if v2243 != 0 {
		goto L546
	} else {
		goto L554
	}
L554:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2247 = F_bms_is_member(m, int32(5), v2246)
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L8
	} else {
		goto L555
	}
L555:
	;
	if v2247 != 0 {
		goto L546
	} else {
		goto L556
	}
L556:
	;
	v2250 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v2251 = F_bms_is_member(m, int32(6), v2250)
	mBase = m.M
	v2252 = m.ExcPending
	if v2252 != 0 {
		goto L8
	} else {
		goto L557
	}
L557:
	;
	if v2251 == int32(0) {
		goto L545
	} else {
		goto L558
	}
L558:
	;
	goto L546
L559:
	;
	v2372 = v2109
	goto L131
L560:
	;
	v2317 = F_order_qual_clauses(m, l0, v40)
	mBase = m.M
	v2318 = m.ExcPending
	if v2318 != 0 {
		goto L8
	} else {
		goto L572
	}
L561:
	;
	v2303 = int32(0)
	goto L560
L562:
	;
	goto L563
L563:
	;
	v2265 = int32(0)
	v2266 = *(*int32)(unsafe.Add(mBase, uint32(v2261)+4))
	if v2266 <= v2265 {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v2303 = int32(0)
	goto L560
L565:
	;
	goto L566
L566:
	;
	v2273 = int32(0)
	v2274 = v2265
	goto L567
L567:
	;
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(v2261)+12))
	v2291 = *(*int32)(unsafe.Add(mBase, uint32(v2287+v2274<<(uint(int32(2))%32))))
	v2293 = F_create_plan_recurse(m, l0, v2291, int32(1))
	mBase = m.M
	v2294 = m.ExcPending
	if v2294 != 0 {
		goto L8
	} else {
		goto L569
	}
L568:
	;
	v2303 = v2295
	goto L560
L569:
	;
	v2295 = F_lappend(m, v2273, v2293)
	mBase = m.M
	v2296 = m.ExcPending
	if v2296 != 0 {
		goto L8
	} else {
		goto L570
	}
L570:
	;
	v2298 = v2274 + int32(1)
	v2299 = *(*int32)(unsafe.Add(mBase, uint32(v2261)+4))
	if v2298 < v2299 {
		v2273 = v2295
		v2274 = v2298
		goto L567
	} else {
		goto L571
	}
L571:
	;
	goto L568
L572:
	;
	v2319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v2319)+4))
	v2321 = m.T0[v2320].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v2260, l1, v459, v2317, v2303)
	mBase = m.M
	v2322 = m.ExcPending
	if v2322 != 0 {
		goto L8
	} else {
		goto L573
	}
L573:
	;
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2321)+4)) = v2323
	v2325 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v2321)+8)) = v2325
	v2327 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v2321)+16)) = v2327
	v2329 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v2321)+24)) = v2329
	v2331 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v2331)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2321)+32)) = v2332
	v2334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2321)+36)) = uint8(v2334)
	v2336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2321)+37)) = uint8(v2336)
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(v2338)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2321)+100)) = v2339
	v2341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v2341 == int32(0) {
		v2372 = v2321
		goto L131
	} else {
		goto L574
	}
L574:
	;
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v2321)+48))
	v2345 = F_replace_nestloop_params_mutator(m, v2344, l0)
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L8
	} else {
		goto L575
	}
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2321)+48)) = v2345
	v2348 = *(*int32)(unsafe.Add(mBase, uint32(v2321)+88))
	v2349 = F_replace_nestloop_params_mutator(m, v2348, l0)
	mBase = m.M
	v2350 = m.ExcPending
	if v2350 != 0 {
		goto L8
	} else {
		goto L576
	}
L576:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2321)+88)) = v2349
	v2372 = v2321
	goto L131
L577:
	;
	v2356 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v2356
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_3), v19)
	mBase = m.M
	v2360 = m.ExcPending
	if v2360 != 0 {
		goto L8
	} else {
		goto L578
	}
L578:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(802), int32(_a_F_create_scan_plan_15))
	mBase = m.M
	v2365 = m.ExcPending
	if v2365 != 0 {
		goto L8
	} else {
		goto L579
	}
L579:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L580:
	;
	v2372 = v2367
	goto L131
L581:
	;
	v2385 = F_create_gating_plan(m, l0, l1, v2372, v65)
	mBase = m.M
	v2386 = m.ExcPending
	if v2386 != 0 {
		goto L8
	} else {
		goto L584
	}
L582:
	;
	v2387 = v2372
	goto L583
L583:
	;
	m.G0 = v19 + int32(176)
	return v2387
L584:
	;
	v2387 = v2385
	goto L583
L585:
	;
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v2396
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_16), v19+int32(48))
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L8
	} else {
		goto L586
	}
L586:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(3932), int32(_a_F_create_scan_plan_12))
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
		goto L8
	} else {
		goto L587
	}
L587:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L588:
	;
	v2412 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v2412
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_17), v19-int32(-64))
	mBase = m.M
	v2418 = m.ExcPending
	if v2418 != 0 {
		goto L8
	} else {
		goto L589
	}
L589:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(3935), int32(_a_F_create_scan_plan_12))
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L8
	} else {
		goto L590
	}
L590:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L591:
	;
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v2444
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_16), v19+int32(80))
	mBase = m.M
	v2450 = m.ExcPending
	if v2450 != 0 {
		goto L8
	} else {
		goto L592
	}
L592:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(3943), int32(_a_F_create_scan_plan_12))
	mBase = m.M
	v2455 = m.ExcPending
	if v2455 != 0 {
		goto L8
	} else {
		goto L593
	}
L593:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L594:
	;
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v2460
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_11), v19+int32(112))
	mBase = m.M
	v2466 = m.ExcPending
	if v2466 != 0 {
		goto L8
	} else {
		goto L595
	}
L595:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(4076), int32(_a_F_create_scan_plan_14))
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L8
	} else {
		goto L596
	}
L596:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L597:
	;
	v2476 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v2476
	F_errmsg_internal(m, int32(_a_F_create_scan_plan_18), v19+int32(128))
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L8
	} else {
		goto L598
	}
L598:
	;
	F_errfinish(m, int32(_a_F_create_scan_plan_4), int32(4086), int32(_a_F_create_scan_plan_14))
	mBase = m.M
	v2487 = m.ExcPending
	if v2487 != 0 {
		goto L8
	} else {
		goto L599
	}
L599:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
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
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
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
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 float64
	_ = v411
	var v413 int32
	_ = v413
	var v415 float64
	_ = v415
	var v417 float64
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
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
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v517 float64
	_ = v517
	var v519 int32
	_ = v519
	var v521 float64
	_ = v521
	var v523 float64
	_ = v523
	var v525 int32
	_ = v525
	var v546 int32
	_ = v546
	var v547 float64
	_ = v547
	var v550 float64
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 float64
	_ = v561
	var v562 float64
	_ = v562
	var v563 int32
	_ = v563
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
	var v577 float64
	_ = v577
	var v580 int32
	_ = v580
	var v581 float64
	_ = v581
	var v587 float64
	_ = v587
	var v594 float64
	_ = v594
	var v595 float64
	_ = v595
	var v599 float64
	_ = v599
	var v602 int32
	_ = v602
	var v605 float64
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v614 float64
	_ = v614
	var v616 int32
	_ = v616
	var v620 float64
	_ = v620
	var v621 float64
	_ = v621
	var v624 float64
	_ = v624
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v634 float64
	_ = v634
	var v636 int32
	_ = v636
	var v637 float64
	_ = v637
	var v638 float64
	_ = v638
	var v639 float64
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 float64
	_ = v664
	var v665 float64
	_ = v665
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v690 float64
	_ = v690
	var v692 float64
	_ = v692
	var v718 int32
	_ = v718
	var v739 int32
	_ = v739
	v5 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(176)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v24 != 0 {
		v739 = v24
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v22 + int32(176)
	return v739
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
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v30))|base.B2i32(int32(1)<<(uint(v30)%32)&int32(44) == int32(0)) != 0 {
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
	v739 = int32(0)
	goto L1
L6:
	;
	v43 = F_GetMemoryChunkContext(m, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+224))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
	if v41 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v739 = int32(0)
	goto L1
L9:
	;
	return int32(0)
L10:
	;
	v47 = int32(_a_F_create_unique_path_0)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_create_unique_path[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_create_unique_path[0])) = v43
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v58 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v51))|base.B2i32(int32(1)<<(uint(v51)%32)&int32(44) == v58) == v58 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_create_unique_path[0])) = v48
	v739 = v718
	goto L1
L12:
	;
	v374 = F_palloc0(m, int32(88))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L9
	} else {
		goto L102
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+224))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+84))
	v66 = F_adjust_appendrel_attrs_multilevel(m, l0, v65, l1, v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v82 = v5
	v83 = v5
	v85 = v5
	v87 = v5
	v88 = v5
	goto L19
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+80))
	v69 = F_copyObjectImpl(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v364 = v66
	v368 = v69
	goto L12
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L9
	} else {
		goto L99
	}
L19:
	;
	v92 = int32(0)
	if v72 == v92 {
		v103 = v92
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L9
	} else {
		goto L96
	}
L21:
	;
	if v71 == int32(0) {
		v718 = v92
		goto L11
	} else {
		goto L24
	}
L22:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v97 <= v85 {
		v103 = int32(0)
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v103 = v99 + v85<<(uint(int32(2))%32)
	goto L21
L24:
	;
	v106 = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if base.B2i32(v103 == v106)|base.B2i32(v108 <= v85) == v106 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113+v85<<(uint(int32(2))%32))))
	v120 = F_get_ordering_op_for_equality_op(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L9
	} else {
		goto L33
	}
L26:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	if v113 != 0 {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v83 != 0 {
		v364 = v83
		v368 = v87
		goto L12
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v718 = v92
	goto L11
L31:
	;
	goto L20
L32:
	;
	v318 = F_copyObjectImpl(m, v115)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L9
	} else {
		goto L93
	}
L33:
	;
	if v120 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v123 = F_get_equality_op_for_ordering_op(m, v120, int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L9
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+45)))
	if v310 == int32(1) {
		goto L18
	} else {
		goto L92
	}
L37:
	;
	if v123 == int32(0) {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	if v82 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+4)))
	v131 = v127 + int32(1)
	goto L41
L40:
	;
	v131 = int32(1)
	goto L41
L41:
	;
	v133 = int32(0)
	v135 = F_makeTargetEntry(m, v115, base.I32_extend16_s(v131), v133, v133)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	v137 = F_lappend(m, v82, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	v140 = F_palloc0(m, int32(20))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = int32(106)
	v144 = int32(0)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v135)+16))
	if v153 == v144 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v286 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v140)+18)) = uint8(v286)
	*(*uint16)(unsafe.Add(mBase, uint32(v140)+16)) = uint16(v286)
	*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v277
	v293 = F_lappend(m, v88, v140)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L9
	} else {
		goto L81
	}
L46:
	;
	if v137 == int32(0) {
		v273 = int32(1)
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v277 = v153
	goto L48
L48:
	;
	goto L45
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+16)) = v273
	v277 = v273
	goto L48
L50:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v160 <= int32(0) {
		v273 = int32(1)
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v163 = int32(0)
	if v163 < v160 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v166 = v160
	goto L54
L53:
	;
	v166 = v163
	goto L54
L54:
	;
	v168 = v166 & int32(3)
	v169 = int32(0)
	if int32(4) <= v160 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v273 = v251 + int32(1)
	goto L49
L56:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v178 = v169
	v179 = int32(0)
	v180 = v144
	goto L59
L57:
	;
	v215 = v169
	v217 = v144
	goto L58
L58:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v226 = int32(0)
	v228 = v215
	v230 = v217
	goto L75
L59:
	;
	v189 = v174 + v180<<(uint(int32(2))%32)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+16))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+16))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+16))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+16))
	if base.Ui32(v178) < base.Ui32(v197) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v168 == int32(0) {
		v251 = v205
		goto L55
	} else {
		goto L74
	}
L61:
	;
	v199 = v197
	goto L63
L62:
	;
	v199 = v178
	goto L63
L63:
	;
	if base.Ui32(v199) < base.Ui32(v195) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v201 = v195
	goto L66
L65:
	;
	v201 = v199
	goto L66
L66:
	;
	if base.Ui32(v201) < base.Ui32(v193) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v203 = v193
	goto L69
L68:
	;
	v203 = v201
	goto L69
L69:
	;
	if base.Ui32(v203) < base.Ui32(v191) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v205 = v191
	goto L72
L71:
	;
	v205 = v203
	goto L72
L72:
	;
	v206 = int32(4)
	v207 = v180 + v206
	v209 = v179 + v206
	if v209 != v166&int32(2147483644) {
		v178 = v205
		v179 = v209
		v180 = v207
		goto L59
	} else {
		goto L73
	}
L73:
	;
	goto L60
L74:
	;
	v215 = v205
	v217 = v207
	goto L58
L75:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v224+v230<<(uint(int32(2))%32))))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+16))
	if base.Ui32(v228) < base.Ui32(v241) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v251 = v243
	goto L55
L77:
	;
	v243 = v241
	goto L79
L78:
	;
	v243 = v228
	goto L79
L79:
	;
	v244 = int32(1)
	v247 = v226 + v244
	if v247 != v168 {
		v226 = v247
		v228 = v243
		v230 = v230 + v244
		goto L75
	} else {
		goto L80
	}
L80:
	;
	goto L76
L81:
	;
	v295 = F_make_pathkeys_for_sortclauses(m, l0, v293, v137)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L9
	} else {
		goto L82
	}
L82:
	;
	if v295 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	v299 = v297
	goto L85
L84:
	;
	v299 = int32(0)
	goto L85
L85:
	;
	if v293 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	v302 = v300
	goto L88
L87:
	;
	v302 = int32(0)
	goto L88
L88:
	;
	if v299 == v302 {
		v315 = v137
		v317 = v293
		goto L32
	} else {
		goto L89
	}
L89:
	;
	v304 = F_list_delete_last(m, v293)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L9
	} else {
		goto L90
	}
L90:
	;
	v306 = F_list_delete_last(m, v137)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L9
	} else {
		goto L91
	}
L91:
	;
	v82 = v306
	v85 = v85 + int32(1)
	v88 = v304
	goto L19
L92:
	;
	v315 = v82
	v317 = v88
	goto L32
L93:
	;
	v320 = F_lappend(m, v83, v318)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	v322 = F_lappend_oid(m, v87, v119)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	v82 = v315
	v83 = v320
	v85 = v85 + int32(1)
	v87 = v322
	v88 = v317
	goto L19
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v120
	F_errmsg_internal(m, int32(_a_F_create_unique_path_1), v22+int32(16))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L9
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_create_unique_path_2), int32(1841), int32(_a_F_create_unique_path_3))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L9
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v119
	F_errmsg_internal(m, int32(_a_F_create_unique_path_4), v22)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_create_unique_path_2), int32(1878), int32(_a_F_create_unique_path_3))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L9
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
	*(*int32)(unsafe.Add(mBase, uint32(v374)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v374))) = int64(1576252997927)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+12)) = v379
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v382 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v374)+20)) = uint8(v382)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+16)) = v381
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v385 == int32(1) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	v390 = v388
	goto L105
L104:
	;
	v390 = int32(0)
	goto L105
L105:
	;
	v392 = v390 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v374)+21)) = uint8(v392)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+84)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v374)+80)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v374)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v374)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+24)) = v394
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v401 != 0 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v374
	v718 = v374
	goto L11
L107:
	;
	v546 = int32(0)
	v547 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v550 = F_estimate_num_groups(m, l0, v364, v547, v546, v546)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L9
	} else {
		goto L139
	}
L108:
	;
	v422 = v401
	goto L110
L109:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+45)))
	if v402 != int32(1) {
		goto L107
	} else {
		goto L111
	}
L110:
	;
	if v422 != int32(1) {
		goto L107
	} else {
		goto L116
	}
L111:
	;
	v405 = int32(0)
	v407 = F_relation_has_unique_index_ext(m, l0, l1, v405, v364, v368, v405)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L9
	} else {
		goto L112
	}
L112:
	;
	if v407 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374)+76)) = int32(0)
	v411 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v374)+32)) = v411
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+40)) = v413
	v415 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v374)+48)) = v415
	v417 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v374)+56)) = v417
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+64)) = v419
	goto L106
L114:
	;
	goto L115
L115:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v422 = v421
	goto L110
L116:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v425 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v439)))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)+36))
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441)+38)))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v441)+120))
	if v445|v442 != 0 {
		v458 = v442 ^ int32(1) | base.B2i32(v445 != int32(0))
		goto L121
	} else {
		goto L122
	}
L118:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v439 = v425 + v426<<(uint(int32(2))%32)
	goto L117
L119:
	;
	goto L120
L120:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+52))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+12))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v439 = v432 + v433<<(uint(int32(2))%32) - int32(4)
	goto L117
L121:
	;
	v459 = int32(0)
	if base.B2i32(v458 == v459)|base.B2i32(v364 == v459) != 0 {
		goto L107
	} else {
		goto L127
	}
L122:
	;
	v450 = int32(1)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v441)+100))
	if v451 != 0 {
		v458 = v450
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v441)+108))
	if v452 != 0 {
		v458 = v450
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441)+36)))
	if v453 != 0 {
		v458 = v450
		goto L121
	} else {
		goto L125
	}
L125:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v441)+112))
	if v454 != 0 {
		v458 = v450
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v441)+144))
	v458 = base.B2i32(v455 != int32(0))
	goto L121
L127:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	if v464 <= int32(0) {
		goto L107
	} else {
		goto L128
	}
L128:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v468 = int32(0)
	v477 = v468
	v479 = v468
	goto L129
L129:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v489+v479<<(uint(int32(2))%32))))
	if v493 == int32(0) {
		goto L107
	} else {
		goto L131
	}
L130:
	;
	if v502 == int32(0) {
		goto L107
	} else {
		goto L136
	}
L131:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v493)))
	if v496 != int32(6) {
		goto L107
	} else {
		goto L132
	}
L132:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v493)+4))
	if v499 != v467 {
		goto L107
	} else {
		goto L133
	}
L133:
	;
	v501 = int32(*(*int16)(unsafe.Add(mBase, uint32(v493)+8)))
	v502 = F_lappend_int(m, v477, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L9
	} else {
		goto L134
	}
L134:
	;
	v505 = v479 + int32(1)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	if v505 < v506 {
		v477 = v502
		v479 = v505
		goto L129
	} else {
		goto L135
	}
L135:
	;
	goto L130
L136:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v440)+36))
	v511 = F_query_is_distinct_for(m, v510, v502, v368)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L9
	} else {
		goto L137
	}
L137:
	;
	if v511 == int32(0) {
		goto L107
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374)+76)) = int32(0)
	v517 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v374)+32)) = v517
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+40)) = v519
	v521 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v374)+48)) = v521
	v523 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v374)+56)) = v523
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+64)) = v525
	goto L106
L139:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v374)+32)) = v550
	if v364 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	v554 = v553
	goto L142
L141:
	;
	v554 = v546
	goto L142
L142:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+45)))
	if v555 == int32(1) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v559 = v22 + int32(104)
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v561 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v562 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)+32))
	v567 = *(*int32)(unsafe.Add(mBase, _c_F_create_unique_path[1]))
	v570 = m.G0
	v571 = int32(16)
	v572 = v570 - v571
	m.G0 = v572
	F_cost_tuplesort(m, v572+int32(8), v572, v562, v564, float64(0), v567, float64(-1))
	mBase = m.M
	v577 = *(*float64)(unsafe.Add(mBase, uint32(v572)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v559)+32)) = v562
	v580 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_unique_path[2])))
	v581 = base.F64_add(v561, v577)
	*(*float64)(unsafe.Add(mBase, uint32(v559)+48)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v559)+40)) = v560 + (v580 ^ int32(1))
	v587 = *(*float64)(unsafe.Add(mBase, uint32(v572)))
	*(*float64)(unsafe.Add(mBase, uint32(v559)+56)) = base.F64_add(v581, v587)
	m.G0 = v572 + v571
	goto L146
L144:
	;
	goto L145
L145:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+46)))
	if v602 != int32(1) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v594 = *(*float64)(unsafe.Add(mBase, _c_F_create_unique_path[3]))
	v595 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v599 = *(*float64)(unsafe.Add(mBase, uint32(v22)+160))
	*(*float64)(unsafe.Add(mBase, uint32(v22)+160)) = base.F64_add(base.F64_mul(base.F64_mul(v594, v595), base.F64_convert_i32_s(v554)), v599)
	goto L145
L147:
	;
	v646 = v22 + int32(88)
	v648 = v22 + int32(80)
	v650 = v22 + int32(72)
	v651 = int32(1)
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+46)))
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+45)))
	if v653 == v651 {
		goto L158
	} else {
		goto L159
	}
L148:
	;
	v605 = *(*float64)(unsafe.Add(mBase, uint32(v374)+32))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)+32))
	v614 = *(*float64)(unsafe.Add(mBase, _c_F_create_unique_path[4]))
	v616 = *(*int32)(unsafe.Add(mBase, _c_F_create_unique_path[1]))
	v620 = base.F64_mul(base.F64_mul(v614, base.F64_convert_i32_s(v616)), float64(1024))
	v621 = float64(4.294967295e+09)
	if base.F64_lt(v620, v621) != 0 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	if base.F64_gt(base.F64_mul(v605, base.F64_convert_i32_s(v607-int32(-64))), base.F64_convert_i32_u(base.I32_trunc_sat_f64_u(v624))) != 0 {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v624 = v620
	goto L152
L151:
	;
	v624 = v621
	goto L152
L152:
	;
	goto L149
L153:
	;
	v628 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+46)) = uint8(v628)
	goto L147
L154:
	;
	goto L155
L155:
	;
	v633 = int32(0)
	v634 = *(*float64)(unsafe.Add(mBase, uint32(v374)+32))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v637 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v638 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v639 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)+32))
	F_cost_agg(m, v22+int32(32), l0, int32(2), v633, v554, v634, v633, v636, v637, v638, v639, base.F64_convert_i32_s(v641))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L9
	} else {
		goto L156
	}
L156:
	;
	goto L147
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374)+76)) = v686
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+40)) = v688
	v690 = *(*float64)(unsafe.Add(mBase, uint32(v685)))
	*(*float64)(unsafe.Add(mBase, uint32(v374)+48)) = v690
	v692 = *(*float64)(unsafe.Add(mBase, uint32(v682)))
	*(*float64)(unsafe.Add(mBase, uint32(v374)+56)) = v692
	goto L106
L158:
	;
	if v652&int32(1) == int32(0) {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	goto L160
L160:
	;
	v676 = int32(0)
	if v652&int32(1) == v676 {
		v718 = v676
		goto L11
	} else {
		goto L166
	}
L161:
	;
	v681 = v22 + int32(144)
	v682 = v22 + int32(160)
	v685 = v22 + int32(152)
	v686 = int32(2)
	goto L157
L162:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v22)+144))
	if v660 < v661 {
		v681 = v650
		v682 = v646
		v685 = v648
		v686 = v651
		goto L157
	} else {
		goto L163
	}
L163:
	;
	if v661 != v660 {
		goto L161
	} else {
		goto L164
	}
L164:
	;
	v664 = *(*float64)(unsafe.Add(mBase, uint32(v22)+88))
	v665 = *(*float64)(unsafe.Add(mBase, uint32(v22)+160))
	if base.F64_lt(v664, v665) != 0 {
		v681 = v650
		v682 = v646
		v685 = v648
		v686 = v651
		goto L157
	} else {
		goto L165
	}
L165:
	;
	goto L161
L166:
	;
	v681 = v650
	v682 = v646
	v685 = v648
	v686 = v651
	goto L157
}
func F_crosstab_hash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int64
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int64
	_ = v140
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int64
	_ = v176
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
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
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int64
	_ = v297
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v348 int64
	_ = v348
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v467 int32
	_ = v467
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int64
	_ = v528
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
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
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int64
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	v25 = m.G0
	v27 = v25 - int32(160)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = F_pg_detoast_datum_packed(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v34 = F_text_to_cstring(m, v30)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v37 = F_pg_detoast_datum_packed(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v39 = F_text_to_cstring(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v41 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L153
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L148
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L143
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L1
	} else {
		goto L139
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L136
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L132
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L128
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L123
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L118
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L114
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L110
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v44 != int32(383) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+12)))
	if v47&int32(2) == int32(0) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	if v52 == int32(0) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v55 = int32(_a_F_crosstab_hash_0)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[0])) = v59
	v61 = F_CreateTupleDescCopy(m, v52)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v63 <= int32(1) {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v27)+64)) = int64(292057776192)
	v74 = F_hash_create(m, int32(_a_F_crosstab_hash_1), int32(64), v27+int32(48), int32(1048))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v81 = F_SPI_execute(m, v39, int32(1), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v202 = F_SPI_finish(m)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L39
	}
L26:
	;
	if v81 != int32(5) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v86 = *(*int64)(unsafe.Add(mBase, _c_F_crosstab_hash[1]))
	if v86 == int64(0) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[2]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v92 != int32(1) {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v117 = int64(0)
	goto L30
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119+base.I32_wrap_i64(v117)<<(uint(int32(2))%32))))
	v126 = F_SPI_getvalue(m, v124, v91, int32(1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L25
L32:
	;
	if v126 == int32(0) {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v130 = int32(_a_F_crosstab_hash_0)
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[0])) = v59
	v135 = F_palloc(m, int32(16))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v135)+8)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v126
	v140 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+152)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+144)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+136)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+128)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+120)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+112)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+104)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v27)+96)) = v140
	v157 = v27 + int32(96)
	v162 = F_pg_snprintf(m, v157, int32(63), int32(_a_F_crosstab_hash_2), v27+int32(32))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v167 = F_hash_search(m, v74, v157, int32(1), v27+int32(47))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+47)))
	if v169 == int32(1) {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+64)) = v135
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[0])) = v131
	v176 = v117 + int64(1)
	if v176 != v86 {
		v117 = v176
		goto L30
	} else {
		goto L38
	}
L38:
	;
	goto L31
L39:
	;
	if v202 != int32(2) {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = int32(2)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)+412))
	if v212 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v276 = F_TupleDescGetAttInMetadata(m, v61)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L45
	}
L42:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v210)+376))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)+364))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v210)+352))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v210)+340))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v210)+328))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v210)+316))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v210)+304))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v210)+292))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v210)+280))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v210)+268))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v210)+256))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v210)+244))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v210)+232))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v210)+220))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v210)+208))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v210)+196))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v210)+184))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v210)+172))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v210)+160))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v210)+148))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v210)+136))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v210)+124))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v210)+112))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v210)+100))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v210)+88))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v210)+76))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v210)+64))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v210)+52))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v210)+40))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v210)+28))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v210)+16))
	v275 = v213 + (v214 + (v215 + (v216 + (v217 + (v218 + (v219 + (v220 + (v221 + (v222 + (v223 + (v224 + (v225 + (v226 + (v227 + (v228 + (v229 + (v230 + (v231 + (v232 + (v233 + (v234 + (v235 + (v236 + (v237 + (v238 + (v239 + (v240 + (v241 + (v242 + (v243 + v211))))))))))))))))))))))))))))))
	goto L44
L43:
	;
	v275 = v211
	goto L44
L44:
	;
	goto L41
L45:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[3]))
	v285 = F_tuplestore_begin_heap(m, int32(base.Ui32(v208&int32(4))>>(uint(int32(2))%32)), int32(0), v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v292 = F_SPI_execute(m, v34, int32(1), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	v609 = F_SPI_finish(m)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L108
	}
L49:
	;
	if v292 != int32(5) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v297 = *(*int64)(unsafe.Add(mBase, _c_F_crosstab_hash[1]))
	if v297 == int64(0) {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	if v275 == int32(0) {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[2]))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	if v305 <= int32(2) {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v309 = v305 - int32(2)
	v310 = v309 + v275
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v310 != v311 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	v315 = F_palloc0(m, v310<<(uint(int32(2))%32))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v320 = int32(1)
	v326 = v320
	v332 = int32(0)
	v348 = int64(0)
	goto L56
L56:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v350+base.I32_wrap_i64(v348)<<(uint(int32(2))%32))))
	v357 = F_SPI_getvalue(m, v355, v304, int32(1))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v581 = F_BuildTupleFromCStrings(m, v276, v315)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L106
	}
L58:
	;
	if v326&int32(1) != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v524 = F_SPI_getvalue(m, v355, v304, v305-v320)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L90
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315))) = v357
	if v305 == int32(3) {
		goto L59
	} else {
		goto L84
	}
L61:
	;
	if v357|v332 == int32(0) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v364 = int32(0)
	if base.B2i32(v332 == v364)|base.B2i32(v357 == v364) == v364 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	if base.B2i32(v373 == int32(0))|base.B2i32(v373 != v376) != 0 {
		v394 = v373
		v395 = v376
		goto L67
	} else {
		goto L68
	}
L64:
	;
	goto L65
L65:
	;
	v399 = F_BuildTupleFromCStrings(m, v276, v315)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L74
	}
L66:
	;
	if v394-v395 == int32(0) {
		goto L59
	} else {
		goto L73
	}
L67:
	;
	goto L66
L68:
	;
	v379 = v332
	v380 = v357
	goto L69
L69:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+1)))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+1)))
	if v384 == int32(0) {
		v394 = v384
		v395 = v383
		goto L67
	} else {
		goto L71
	}
L70:
	;
	v394 = v384
	v395 = v383
	goto L67
L71:
	;
	v387 = int32(1)
	if v384 == v383 {
		v379 = v379 + v387
		v380 = v380 + v387
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	goto L65
L74:
	;
	F_tuplestore_puttuple(m, v285, v399)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v403 = int32(0)
	if v310 <= v403 {
		goto L60
	} else {
		goto L76
	}
L76:
	;
	v406 = v403
	goto L77
L77:
	;
	v432 = v315 + v406<<(uint(int32(2))%32)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)))
	if v433 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L60
L79:
	;
	F_pfree(m, v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v439 = v406 + int32(1)
	if v439 != v310 {
		v406 = v439
		goto L77
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432))) = int32(0)
	goto L81
L83:
	;
	goto L78
L84:
	;
	v467 = int32(1)
	goto L85
L85:
	;
	v495 = v467 + int32(1)
	v496 = F_SPI_getvalue(m, v355, v304, v495)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	goto L59
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315+v467<<(uint(int32(2))%32)))) = v496
	if v495 != v309 {
		v467 = v495
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	if v332 != 0 {
		goto L97
	} else {
		goto L98
	}
L90:
	;
	if v524 == int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v528 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+152)) = v528
	*(*int64)(unsafe.Add(mBase, uint32(v27)+144)) = v528
	*(*int64)(unsafe.Add(mBase, uint32(v27)+136)) = v528
	*(*int64)(unsafe.Add(mBase, uint32(v27)+128)) = v528
	*(*int64)(unsafe.Add(mBase, uint32(v27)+120)) = v528
	*(*int64)(unsafe.Add(mBase, uint32(v27)+112)) = v528
	*(*int64)(unsafe.Add(mBase, uint32(v27)+104)) = v528
	*(*int64)(unsafe.Add(mBase, uint32(v27)+96)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v524
	v546 = v27 + int32(96)
	v549 = F_pg_snprintf(m, v546, int32(63), int32(_a_F_crosstab_hash_2), v27)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v551 = int32(0)
	v553 = F_hash_search(m, v74, v546, v551, v551)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v553 == int32(0) {
		goto L89
	} else {
		goto L94
	}
L94:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v553)+64))
	if v557 == int32(0) {
		goto L89
	} else {
		goto L95
	}
L95:
	;
	v560 = F_SPI_getvalue(m, v355, v304, v305)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v557)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v315+v305<<(uint(int32(2))%32)+v562<<(uint(int32(2))%32)-int32(8)))) = v560
	goto L89
L97:
	;
	F_pfree(m, v332)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v573 = int32(0)
	if v357 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L99
L101:
	;
	v575 = F_pstrdup(m, v357)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	v577 = v573
	goto L103
L103:
	;
	v579 = v348 + int64(1)
	if v579 != v297 {
		v326 = v573
		v332 = v577
		v348 = v579
		goto L56
	} else {
		goto L105
	}
L104:
	;
	v577 = v575
	goto L103
L105:
	;
	goto L57
L106:
	;
	F_tuplestore_puttuple(m, v285, v581)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	goto L48
L108:
	;
	if v609 != int32(2) {
		goto L6
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v285
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab_hash[0])) = v56
	m.G0 = v27 + int32(160)
	return int32(0)
L110:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_3), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(650), int32(_a_F_crosstab_hash_5))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_6), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(655), int32(_a_F_crosstab_hash_5))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_7), int32(0))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errdetail(m, int32(_a_F_crosstab_hash_8), int32(0))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(674), int32(_a_F_crosstab_hash_5))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_9), int32(0))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errdetail(m, int32(_a_F_crosstab_hash_10), int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(749), int32(_a_F_crosstab_hash_11))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_12), int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(765), int32(_a_F_crosstab_hash_11))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(_a_F_crosstab_hash_13))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_14), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(774), int32(_a_F_crosstab_hash_11))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errmsg_internal(m, int32(_a_F_crosstab_hash_15), int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(782), int32(_a_F_crosstab_hash_11))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_16), int32(0))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(832), int32(_a_F_crosstab_hash_17))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_18), int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errdetail(m, int32(_a_F_crosstab_hash_19), int32(0))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(852), int32(_a_F_crosstab_hash_17))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(_a_F_crosstab_hash_7), int32(0))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v310
	F_errdetail(m, int32(_a_F_crosstab_hash_20), v27+int32(16))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(862), int32(_a_F_crosstab_hash_17))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errmsg_internal(m, int32(_a_F_crosstab_hash_21), int32(0))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_crosstab_hash_4), int32(932), int32(_a_F_crosstab_hash_17))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v75 int64
	_ = v75
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
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
	v27 = v10 - int32(-64)
	F_initStringInfo(m, v27)
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(_a_F_cursor_to_xml_0)
	F_appendStringInfo(m, v27, int32(_a_F_cursor_to_xml_1), v10+int32(48))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
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
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L18
	}
L10:
	;
	F_appendStringInfoString(m, v27, int32(_a_F_cursor_to_xml_2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v42 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v24
	F_appendStringInfo(m, v27, int32(_a_F_cursor_to_xml_3), v10+int32(32))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v50 = v10 - int32(-64)
	F_appendStringInfoString(m, v50, int32(_a_F_cursor_to_xml_4))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	F_appendStringInfoChar(m, v50, int32(10))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L9
L18:
	;
	v61 = F_GetPortalByName(m, v17)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v61 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_SPI_cursor_fetch(m, v61, v20)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
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
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L37
	}
L23:
	;
	v66 = *(*int64)(unsafe.Add(mBase, _c_F_cursor_to_xml[0]))
	if v66 != int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v75 = int64(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	v94 = F_SPI_finish(m)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L31
	}
L27:
	;
	F_SPI_sql_row_to_xmlelement(m, v10-int32(-64), base.B2i32(v19 != int32(0)), v24)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v83 = v75 + int64(1)
	v85 = *(*int64)(unsafe.Add(mBase, _c_F_cursor_to_xml[0]))
	if base.Ui64(v83) < base.Ui64(v85) {
		v75 = v83
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_cursor_to_xml_0)
	F_appendStringInfo(m, v10-int32(-64), int32(_a_F_cursor_to_xml_5), v10+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
	v109 = F_cstring_to_text_with_len(m, v107, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
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
	return v109
L37:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v17
	F_errmsg(m, int32(_a_F_cursor_to_xml_6), v10)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_cursor_to_xml_7), int32(2937), int32(_a_F_cursor_to_xml_8))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
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
