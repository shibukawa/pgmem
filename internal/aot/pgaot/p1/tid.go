package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecTidScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(768), int32(769))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_TidRangeNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
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
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)))
	if v11 == v2 {
		v14 = F_TidRangeEval(m, l0)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v14 == int32(0) {
				v79 = v2
				return v79
			} else {
				if v8 == int32(0) {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					v24 = int32(0)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+188))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
					v30 = m.T0[v29].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v22, v23, v24, v24, v24, int32(272))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+188))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
						m.T0[v38].(func(*base.Module, int32, int32, int32))(m, v30, l0+int32(120), l0+int32(126))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v30
							v61 = v30
							v62 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v62)
							v64 = v61
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+188))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+28))
							v68 = m.T0[v67].(func(*base.Module, int32, int32, int32) int32)(m, v64, v10, v7)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								if v68 == int32(0) {
									v72 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v72)
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									m.T0[v75].(func(*base.Module, int32))(m, v7)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v79 = v7
										return v79
									}
								} else {
									v79 = v7
									return v79
								}
							}
						}
					}
				} else {
					v42 = int32(0)
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+188))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
					m.T0[v49].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v8, v42, v42, v42, v42, v42)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+188))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
						m.T0[v58].(func(*base.Module, int32, int32, int32))(m, v8, l0+int32(120), l0+int32(126))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							v61 = v8
							v62 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v62)
							v64 = v61
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+188))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+28))
							v68 = m.T0[v67].(func(*base.Module, int32, int32, int32) int32)(m, v64, v10, v7)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								if v68 == int32(0) {
									v72 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v72)
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									m.T0[v75].(func(*base.Module, int32))(m, v7)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v79 = v7
										return v79
									}
								} else {
									v79 = v7
									return v79
								}
							}
						}
					}
				}
			}
		}
	} else {
		v64 = v8
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
		v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+188))
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+28))
		v68 = m.T0[v67].(func(*base.Module, int32, int32, int32) int32)(m, v64, v10, v7)
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return int32(0)
		} else {
			if v68 == int32(0) {
				v72 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v72)
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
				m.T0[v75].(func(*base.Module, int32))(m, v7)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					v79 = v7
					return v79
				}
			} else {
				v79 = v7
				return v79
			}
		}
	}
}
func F_TidRangeRecheck(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	v3 = int32(0)
	v4 = F_TidRangeEval(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v69 = v3
		} else {
			v11 = l1 + int32(28)
			v13 = l0 + int32(120)
			v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)))
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
			v19 = int32(16)
			v21 = v17 | v18<<(uint(v19)%32)
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
			v26 = v22 | v23<<(uint(v19)%32)
			if base.Ui32(v21) < base.Ui32(v26) {
				v37 = int32(-1)
			} else {
				if base.Ui32(v26) < base.Ui32(v21) {
					v37 = int32(1)
				} else {
					v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
					v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
					if base.Ui32(v31) < base.Ui32(v32) {
						v37 = int32(-1)
					} else {
						v37 = base.B2i32(base.Ui32(v32) < base.Ui32(v31))
					}
				}
			}
			if v37 < int32(0) {
				v69 = v3
			} else {
				v41 = l0 + int32(126)
				v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)))
				v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
				v47 = int32(16)
				v49 = v45 | v46<<(uint(v47)%32)
				v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+2)))
				v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41))))
				v54 = v50 | v51<<(uint(v47)%32)
				if base.Ui32(v49) < base.Ui32(v54) {
					v65 = int32(-1)
				} else {
					if base.Ui32(v54) < base.Ui32(v49) {
						v65 = int32(1)
					} else {
						v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
						v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)))
						if base.Ui32(v59) < base.Ui32(v60) {
							v65 = int32(-1)
						} else {
							v65 = base.B2i32(base.Ui32(v60) < base.Ui32(v59))
						}
					}
				}
				v69 = base.B2i32(v65 <= int32(0))
			}
		}
		return v69
	}
}
func F_TidStoreSetBlockOffsets(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v152 int32
	_ = v152
	var v165 int32
	_ = v165
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v236 int32
	_ = v236
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
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v275 int32
	_ = v275
	var v276 int64
	_ = v276
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v352 int32
	_ = v352
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v394 int32
	_ = v394
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v492 int32
	_ = v492
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v535 int64
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int64
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v650 int32
	_ = v650
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v705 int32
	_ = v705
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v784 int32
	_ = v784
	var v791 int32
	_ = v791
	var v800 int32
	_ = v800
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int64
	_ = v816
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int64
	_ = v836
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v906 int32
	_ = v906
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v983 int32
	_ = v983
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1075 int32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1129 int32
	_ = v1129
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1216 int32
	_ = v1216
	var v1241 int32
	_ = v1241
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1265 int32
	_ = v1265
	var v1275 int32
	_ = v1275
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1373 int64
	_ = v1373
	var v1375 int64
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1395 int64
	_ = v1395
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1414 int32
	_ = v1414
	var v1419 int32
	_ = v1419
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1471 int32
	_ = v1471
	var v1472 int64
	_ = v1472
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1491 int32
	_ = v1491
	var v1503 int32
	_ = v1503
	var v1512 int32
	_ = v1512
	var v1517 int32
	_ = v1517
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1543 int32
	_ = v1543
	var v1565 int32
	_ = v1565
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1583 int32
	_ = v1583
	var v1605 int32
	_ = v1605
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1676 int32
	_ = v1676
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1731 int64
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1743 int32
	_ = v1743
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1790 int32
	_ = v1790
	var v1793 int32
	_ = v1793
	var v1798 int32
	_ = v1798
	var v1801 int32
	_ = v1801
	var v1806 int32
	_ = v1806
	var v1811 int32
	_ = v1811
	var v1815 int32
	_ = v1815
	var v1820 int32
	_ = v1820
	var v1828 int32
	_ = v1828
	var v1845 int32
	_ = v1845
	var v1848 int32
	_ = v1848
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1882 int32
	_ = v1882
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1915 int32
	_ = v1915
	var v1938 int32
	_ = v1938
	var v1943 int32
	_ = v1943
	var v1950 int32
	_ = v1950
	var v1962 int32
	_ = v1962
	var v1969 int32
	_ = v1969
	var v1977 int32
	_ = v1977
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2009 int64
	_ = v2009
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2022 int32
	_ = v2022
	var v2025 int32
	_ = v2025
	var v2029 int32
	_ = v2029
	var v2034 int32
	_ = v2034
	var v2054 int32
	_ = v2054
	var v2057 int32
	_ = v2057
	var v2080 int32
	_ = v2080
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2108 int32
	_ = v2108
	var v2126 int32
	_ = v2126
	var v2129 int32
	_ = v2129
	var v2131 int32
	_ = v2131
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2145 int32
	_ = v2145
	var v2153 int32
	_ = v2153
	var v2156 int32
	_ = v2156
	var v2161 int32
	_ = v2161
	var v2163 int32
	_ = v2163
	var v2166 int32
	_ = v2166
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2177 int32
	_ = v2177
	var v2185 int32
	_ = v2185
	var v2186 int32
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2195 int32
	_ = v2195
	var v2213 int32
	_ = v2213
	var v2217 int32
	_ = v2217
	var v2219 int32
	_ = v2219
	var v2223 int32
	_ = v2223
	var v2225 int32
	_ = v2225
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2235 int32
	_ = v2235
	var v2243 int32
	_ = v2243
	var v2248 int32
	_ = v2248
	var v2251 int32
	_ = v2251
	var v2255 int32
	_ = v2255
	var v2278 int32
	_ = v2278
	var v2280 int32
	_ = v2280
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2295 int32
	_ = v2295
	var v2312 int32
	_ = v2312
	var v2315 int32
	_ = v2315
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2326 int32
	_ = v2326
	var v2328 int32
	_ = v2328
	var v2334 int32
	_ = v2334
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2349 int32
	_ = v2349
	var v2353 int32
	_ = v2353
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2366 int32
	_ = v2366
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2375 int32
	_ = v2375
	var v2377 int32
	_ = v2377
	var v2382 int32
	_ = v2382
	var v2407 int32
	_ = v2407
	var v2412 int32
	_ = v2412
	var v2415 int32
	_ = v2415
	var v2431 int32
	_ = v2431
	var v2441 int32
	_ = v2441
	var v2448 int32
	_ = v2448
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2464 int32
	_ = v2464
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2472 int32
	_ = v2472
	var v2474 int32
	_ = v2474
	var v2483 int32
	_ = v2483
	var v2485 int32
	_ = v2485
	var v2491 int32
	_ = v2491
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2502 int32
	_ = v2502
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2512 int32
	_ = v2512
	var v2515 int32
	_ = v2515
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2523 int32
	_ = v2523
	var v2528 int32
	_ = v2528
	var v2532 int32
	_ = v2532
	var v2537 int32
	_ = v2537
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2561 int32
	_ = v2561
	var v2564 int32
	_ = v2564
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2571 int32
	_ = v2571
	var v2572 int32
	_ = v2572
	var v2573 int64
	_ = v2573
	var v2575 int32
	_ = v2575
	var v2577 int32
	_ = v2577
	var v2586 int32
	_ = v2586
	var v2589 int32
	_ = v2589
	var v2591 int32
	_ = v2591
	var v2593 int32
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2596 int32
	_ = v2596
	var v2605 int32
	_ = v2605
	var v2616 int32
	_ = v2616
	var v2618 int32
	_ = v2618
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2639 int32
	_ = v2639
	var v2642 int32
	_ = v2642
	var v2648 int32
	_ = v2648
	var v2653 int32
	_ = v2653
	var v2655 int32
	_ = v2655
	var v2657 int32
	_ = v2657
	var v2660 int32
	_ = v2660
	var v2666 int32
	_ = v2666
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2683 int32
	_ = v2683
	var v2688 int32
	_ = v2688
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2695 int32
	_ = v2695
	var v2699 int32
	_ = v2699
	var v2701 int32
	_ = v2701
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2715 int32
	_ = v2715
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2720 int32
	_ = v2720
	var v2732 int32
	_ = v2732
	var v2735 int32
	_ = v2735
	var v2762 int32
	_ = v2762
	var v2765 int32
	_ = v2765
	var v2766 int32
	_ = v2766
	var v2771 int32
	_ = v2771
	var v2787 int32
	_ = v2787
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2796 int32
	_ = v2796
	var v2800 int32
	_ = v2800
	var v2804 int32
	_ = v2804
	var v2805 int32
	_ = v2805
	var v2806 int32
	_ = v2806
	var v2812 int32
	_ = v2812
	var v2813 int64
	_ = v2813
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2840 int64
	_ = v2840
	var v2843 int32
	_ = v2843
	var v2844 int32
	_ = v2844
	var v2847 int32
	_ = v2847
	var v2848 int32
	_ = v2848
	var v2849 int32
	_ = v2849
	var v2852 int32
	_ = v2852
	var v2853 int32
	_ = v2853
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2857 int64
	_ = v2857
	var v2859 int32
	_ = v2859
	var v2861 int32
	_ = v2861
	var v2869 int32
	_ = v2869
	var v2872 int32
	_ = v2872
	var v2874 int32
	_ = v2874
	var v2876 int32
	_ = v2876
	var v2878 int32
	_ = v2878
	var v2880 int32
	_ = v2880
	var v2884 int32
	_ = v2884
	var v2888 int32
	_ = v2888
	var v2901 int32
	_ = v2901
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2906 int32
	_ = v2906
	var v2924 int32
	_ = v2924
	var v2927 int32
	_ = v2927
	var v2933 int32
	_ = v2933
	var v2938 int32
	_ = v2938
	var v2940 int32
	_ = v2940
	var v2942 int32
	_ = v2942
	var v2945 int32
	_ = v2945
	var v2951 int32
	_ = v2951
	var v2956 int32
	_ = v2956
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2960 int32
	_ = v2960
	var v2968 int32
	_ = v2968
	var v2972 int32
	_ = v2972
	var v2973 int32
	_ = v2973
	var v2975 int32
	_ = v2975
	var v2980 int32
	_ = v2980
	var v2981 int32
	_ = v2981
	var v2982 int32
	_ = v2982
	var v2986 int32
	_ = v2986
	var v2988 int32
	_ = v2988
	var v2996 int32
	_ = v2996
	var v2999 int32
	_ = v2999
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3004 int32
	_ = v3004
	var v3007 int32
	_ = v3007
	var v3019 int32
	_ = v3019
	var v3022 int32
	_ = v3022
	var v3049 int32
	_ = v3049
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3058 int32
	_ = v3058
	var v3074 int32
	_ = v3074
	var v3081 int32
	_ = v3081
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3086 int32
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3091 int32
	_ = v3091
	var v3095 int32
	_ = v3095
	var v3096 int32
	_ = v3096
	var v3097 int32
	_ = v3097
	var v3104 int32
	_ = v3104
	var v3105 int64
	_ = v3105
	v5 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(336)
	m.G0 = v23
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v5
	if base.Ui32(int32(2)) <= base.Ui32(l3) {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	m.G0 = v23 + int32(336)
	return
L2:
	;
	v3074 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+65)))
	if base.Ui32(int32(5)) <= base.Ui32(v3074<<(uint(int32(2))%32)+int32(4)) {
		goto L462
	} else {
		goto L463
	}
L3:
	;
	if v321 == int32(0) {
		v3058 = v3049
		goto L2
	} else {
		goto L459
	}
L4:
	;
	v2837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	v2838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	if v2837 == v2838 {
		goto L438
	} else {
		goto L439
	}
L5:
	;
	v2787 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+65)))
	if base.Ui32(int32(5)) <= base.Ui32(v2787<<(uint(int32(2))%32)+int32(4)) {
		goto L426
	} else {
		goto L427
	}
L6:
	;
	if v1517 == int32(0) {
		v2771 = v2762
		goto L5
	} else {
		goto L423
	}
L7:
	;
	v2558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527)+2)))
	v2559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527)+1)))
	if v2558 == v2559 {
		goto L403
	} else {
		goto L404
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2528 = m.ExcPending
	if v2528 != 0 {
		goto L28
	} else {
		goto L400
	}
L9:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+65)) = uint8(v152)
	v165 = base.I32_extend8_s(v152)<<(uint(int32(2))%32) + int32(4)
	v166 = base.I64_extend_i32_u(l1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v169 != 0 {
		goto L34
	} else {
		goto L35
	}
L10:
	;
	v29 = int32(1)
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+l3<<(uint(v29)%32)-int32(2)))))
	v38 = int32(base.Ui32(v34)>>(uint(int32(5))%32)) + v29
	v46 = v5
	v48 = int32(32)
	v55 = v5
	goto L13
L11:
	;
	goto L12
L12:
	;
	if l3 != int32(1) {
		v152 = v5
		goto L9
	} else {
		goto L32
	}
L13:
	;
	if l3 < v46 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L28
	} else {
		goto L29
	}
L15:
	;
	v63 = v46
	goto L17
L16:
	;
	v63 = l3
	goto L17
L17:
	;
	v70 = int32(0)
	v73 = v46
	goto L19
L18:
	;
	goto L14
L19:
	;
	if v63 == v73 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(68)+v55<<(uint(int32(2))%32)))) = v70
	v113 = v55 + int32(1)
	if v113 != v38 {
		v46 = v102
		v48 = v48 + int32(32)
		v55 = v113
		goto L13
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v102 = v63
	goto L21
L23:
	;
	goto L24
L24:
	;
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v73<<(uint(int32(1))%32)))))
	if base.Ui32((v89-int32(2049))&int32(65535)) <= base.Ui32(int32(63487)) {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v96 = int32(1)
	if base.Ui32(v89) < base.Ui32(v48) {
		v70 = v96<<(uint(v89)%32) | v70
		v73 = v73 + v96
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v102 = v73
	goto L21
L27:
	;
	v152 = v38
	goto L9
L28:
	;
	return
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v89
	F_errmsg_internal(m, int32(58754), v23+int32(48))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(498501), int32(396), int32(123995))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	if base.Ui32((v132-int32(2049))&int32(65535)) <= base.Ui32(int32(63487)) {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+66)) = uint16(v132)
	v152 = v5
	goto L9
L34:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v168)+32))
	if base.Ui64(v170) < base.Ui64(v166) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v1373 = *(*int64)(unsafe.Add(mBase, uint32(v168)+8))
	if base.Ui64(v1373) < base.Ui64(v166) {
		goto L224
	} else {
		goto L225
	}
L37:
	;
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v168)+40))
	if v172 == int64(0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v295 = v168
	goto L39
L39:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v295)+48))
	v321 = v309
	v323 = v295 + int32(24)
	goto L56
L40:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v168)+24))
	v177 = F_dsa_get_address(m, v175, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L28
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+48))
	if v166 == int64(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v179 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v177)+2)) = uint8(v179)
	v185 = (base.I32_clz(l1) ^ int32(-1)) & int32(24)
	v186 = int32(base.Ui32(l1) >> (uint(v185) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v177)+3)) = uint8(v186)
	v190 = F_shared_ts_extend_down(m, v167, v177+int32(8), v166, v185)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L28
	} else {
		goto L44
	}
L44:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+48)) = v185
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v195 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v194)+32)) = v195<<(uint(base.I64_extend_i32_u(v185+int32(8)))%64) ^ v195
	v3058 = v190
	goto L2
L45:
	;
	v214 = int32(0)
	goto L47
L46:
	;
	v214 = (base.I32_wrap_i64(base.I64_clz(v166)) ^ int32(-1)) & int32(56)
	goto L47
L47:
	;
	if v204 < v214 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v220 = v204
	goto L51
L49:
	;
	v275 = v203
	goto L50
L50:
	;
	v276 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v275)+32)) = v276<<(uint(base.I64_extend_i32_u(v214+int32(8)))%64) ^ v276
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v284)+48)) = v214
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v295 = v286
	goto L39
L51:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v239 = F_dsa_allocate_extended(m, v236, int32(24), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L28
	} else {
		goto L53
	}
L52:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v275 = v254
	goto L50
L53:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v242 = F_dsa_get_address(m, v241, v239)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L28
	} else {
		goto L54
	}
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v242))) = int64(66560)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = v247
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v249)+24)) = v239
	v252 = v220 + int32(8)
	if v252 < v214 {
		v220 = v252
		goto L51
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	v332 = base.I32_wrap_i64(int64(base.Ui64(v166) >> (uint(base.I64_extend_i32_u(v321)) % 64)))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	v335 = F_dsa_get_address(m, v333, v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L28
	} else {
		goto L73
	}
L57:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v478)))
	v1319 = v1317 & int32(1)
	if base.Ui32(v165) <= base.Ui32(int32(4)) {
		goto L193
	} else {
		goto L194
	}
L58:
	;
	if v321 != 0 {
		goto L190
	} else {
		goto L191
	}
L59:
	;
	v1298 = int32(2)
	v1302 = v335 + v1286<<(uint(v1298)%32) + int32(4)
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1302)))
	*(*int32)(unsafe.Add(mBase, uint32(v1302))) = v1303 | v1285
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	v1308 = v1306 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)) = uint8(v1308)
	v3049 = v335 + v1282<<(uint(v1298)%32) + int32(36)
	goto L3
L60:
	;
	v1275 = v332 & int32(255)
	v1282 = v1275
	v1285 = int32(1) << (uint(v332) % 32)
	v1286 = int32(base.Ui32(v1275) >> (uint(int32(5)) % 32))
	goto L59
L61:
	;
	v813 = v800 & int32(255)
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	if v813 == v814 {
		goto L132
	} else {
		goto L133
	}
L62:
	;
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	v800 = v791
	goto L61
L63:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v765+v335)+3)) = uint8(v332)
	v784 = v768 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)) = uint8(v784)
	v3049 = v335 + v765<<(uint(int32(2))%32) + int32(8)
	goto L3
L64:
	;
	v642 = v335 + int32(3)
	v650 = int32(0)
	goto L115
L65:
	;
	v535 = *(*int64)(unsafe.Add(mBase, uint32(v23)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v535
	v538 = v332 & int32(255)
	v539 = int32(0)
	v541 = v23 + int32(24)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v546 = F_dsa_allocate_extended(m, v543, int32(96), v539)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L28
	} else {
		goto L97
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v23)+332)) = v335
	v512 = int32(0)
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	if v514 != 0 {
		v765 = v512
		v768 = v512
		goto L63
	} else {
		goto L96
	}
L67:
	;
	v505 = v492 & int32(255)
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	if v505 == v506 {
		goto L65
	} else {
		goto L94
	}
L68:
	;
	if v478 != 0 {
		goto L58
	} else {
		goto L92
	}
L69:
	;
	v439 = int32(1) << (uint(v332) % 32)
	v441 = v332 & int32(255)
	v443 = int32(base.Ui32(v441) >> (uint(int32(5)) % 32))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v335+v443<<(uint(int32(2))%32))+4))
	if v439&v447 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L70:
	;
	v425 = int32(255)
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335+v332&v425)+12)))
	if v428 == v425 {
		goto L86
	} else {
		goto L87
	}
L71:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	if v381 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	if v340 == int32(0) {
		goto L66
	} else {
		goto L74
	}
L73:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	switch v337 - int32(1) {
	case 0:
		goto L71
	case 1:
		goto L70
	case 2:
		goto L69
	default:
		goto L72
	}
L74:
	;
	v352 = int32(0)
	goto L75
L75:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352+(v335+int32(3))))))
	if v332&int32(255) == v374 {
		v478 = v335 + v352<<(uint(int32(2))%32) + int32(8)
		goto L68
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+332)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v334
	v492 = v340
	goto L67
L77:
	;
	v377 = v352 + int32(1)
	if v377 != v340 {
		v352 = v377
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+332)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v334
	v800 = int32(0)
	goto L61
L80:
	;
	goto L81
L81:
	;
	v394 = int32(0)
	goto L82
L82:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394+(v335+int32(3))))))
	if v416 == v332&int32(255) {
		v478 = v335 + v394<<(uint(int32(2))%32) + int32(36)
		goto L68
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+332)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v334
	v800 = v381
	goto L61
L84:
	;
	v421 = v394 + int32(1)
	if v421 != v381 {
		v394 = v421
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+332)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v334
	goto L4
L87:
	;
	goto L88
L88:
	;
	v478 = v335 + v428<<(uint(int32(2))%32) + int32(268)
	goto L68
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+332)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v334
	v1282 = v441
	v1285 = v439
	v1286 = v443
	goto L59
L90:
	;
	goto L91
L91:
	;
	v478 = v335 + v441<<(uint(int32(2))%32) + int32(36)
	goto L68
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+332)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v334
	switch v337 - int32(1) {
	case 0:
		goto L62
	case 1:
		goto L4
	case 2:
		goto L60
	default:
		goto L93
	}
L93:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	v492 = v483
	goto L67
L94:
	;
	if v505 != 0 {
		goto L64
	} else {
		goto L95
	}
L95:
	;
	v765 = int32(0)
	v768 = int32(0)
	goto L63
L96:
	;
	goto L65
L97:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v549 = F_dsa_get_address(m, v548, v546)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L28
	} else {
		goto L98
	}
L98:
	;
	v551 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v549)+2)) = v551
	v553 = int32(3841)
	*(*uint16)(unsafe.Add(mBase, uint32(v549))) = uint16(v553)
	v555 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v549)+34)) = uint16(v555)
	*(*int64)(unsafe.Add(mBase, uint32(v549)+26)) = v551
	*(*int64)(unsafe.Add(mBase, uint32(v549)+18)) = v551
	*(*int64)(unsafe.Add(mBase, uint32(v549)+10)) = v551
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v549)+2)) = uint8(v563)
	v565 = int32(2)
	v566 = int32(4)
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+3)))
	v569 = base.B2i32(base.Ui32(v538) <= base.Ui32(v568))
	if base.Ui32(v538) <= base.Ui32(v568) {
		v589 = v566
		v590 = v565
		v591 = v539
		v592 = int32(3)
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v594 = v549 + int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v594+v569))) = uint8(v568)
	v598 = v549 + int32(36)
	v599 = int32(2)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v542)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v598+v569<<(uint(v599)%32)))) = v602
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v594+v590))) = uint8(v605)
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v542)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v598+v590<<(uint(v599)%32)))) = v610
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v594+v592))) = uint8(v613)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v542)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v598+v592<<(uint(v599)%32)))) = v618
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v594+v589))) = uint8(v621)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v542)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v598+v589<<(uint(v599)%32)))) = v626
	*(*uint8)(unsafe.Add(mBase, uint32(v594+v591))) = uint8(v538)
	v631 = v563 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v549)+2)) = uint8(v631)
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v546
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v541)))
	F_dsa_free(m, v634, v635)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L28
	} else {
		goto L113
	}
L100:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+4)))
	if base.Ui32(v538) <= base.Ui32(v570) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v589 = v566
	v590 = v565
	v591 = int32(1)
	v592 = int32(3)
	goto L99
L102:
	;
	goto L103
L103:
	;
	v574 = int32(1)
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+5)))
	if base.Ui32(v538) <= base.Ui32(v575) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v589 = v566
	v590 = v574
	v591 = int32(2)
	v592 = int32(3)
	goto L99
L105:
	;
	goto L106
L106:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+6)))
	v582 = base.B2i32(base.Ui32(v581) < base.Ui32(v538))
	if base.Ui32(v581) < base.Ui32(v538) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v583 = int32(4)
	goto L109
L108:
	;
	v583 = int32(3)
	goto L109
L109:
	;
	if base.Ui32(v581) < base.Ui32(v538) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v586 = int32(3)
	goto L112
L111:
	;
	v586 = int32(4)
	goto L112
L112:
	;
	v589 = v586
	v590 = v574
	v591 = v583
	v592 = int32(2)
	goto L99
L113:
	;
	v3049 = v598 + v591<<(uint(int32(2))%32)
	goto L3
L114:
	;
	if base.Ui32(v505) <= base.Ui32(v650) {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650+v642))))
	if base.Ui32(v332&int32(255)) <= base.Ui32(v667) {
		goto L114
	} else {
		goto L117
	}
L116:
	;
	v765 = v505
	v768 = v492
	goto L63
L117:
	;
	v670 = v650 + int32(1)
	if v670 != v505 {
		v650 = v670
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v765 = v650
	v768 = v492
	goto L63
L120:
	;
	goto L121
L121:
	;
	v674 = v335 + int32(8)
	v675 = v505 - v650
	if v675&int32(1) == int32(0) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	if v675 != int32(1) {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v694 = v505
	goto L122
L124:
	;
	goto L125
L125:
	;
	v682 = v505 - int32(1)
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642+v682))))
	*(*uint8)(unsafe.Add(mBase, uint32(v505+v642))) = uint8(v684)
	v686 = int32(2)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v674+v682<<(uint(v686)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v674+v505<<(uint(v686)%32)))) = v692
	v694 = v682
	goto L122
L126:
	;
	v705 = v694
	goto L129
L127:
	;
	goto L128
L128:
	;
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	v765 = v650
	v768 = v760
	goto L63
L129:
	;
	v719 = v705 - int32(1)
	v720 = v642 + v719
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v720))))
	*(*uint8)(unsafe.Add(mBase, uint32(v642+v705))) = uint8(v721)
	v723 = int32(2)
	v724 = v705 << (uint(v723) % 32)
	v728 = v674 + v719<<(uint(v723)%32)
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v728)))
	*(*int32)(unsafe.Add(mBase, uint32(v674+v724))) = v729
	v732 = v705 - v723
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642+v732))))
	*(*uint8)(unsafe.Add(mBase, uint32(v720))) = uint8(v734)
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v335+v724)))
	*(*int32)(unsafe.Add(mBase, uint32(v728))) = v737
	if v650 < v732 {
		v705 = v732
		goto L129
	} else {
		goto L131
	}
L130:
	;
	goto L128
L131:
	;
	goto L130
L132:
	;
	v816 = *(*int64)(unsafe.Add(mBase, uint32(v23)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v816
	v819 = v332 & int32(255)
	v820 = int32(0)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v824 = v23 + int32(32)
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v824)+4))
	v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v825)+1)))
	if base.Ui32(v826) <= base.Ui32(int32(30)) {
		goto L138
	} else {
		goto L139
	}
L133:
	;
	goto L134
L134:
	;
	v1114 = v335 + int32(3)
	if v813 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L135:
	;
	v3049 = v1112
	goto L3
L136:
	;
	goto L163
L137:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v906+v932))) = uint8(v819)
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926))))
	v1060 = v1058 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v926))) = uint8(v1060)
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v824)))
	F_dsa_free(m, v1062, v1063)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L28
	} else {
		goto L161
	}
L138:
	;
	v831 = F_dsa_allocate_extended(m, v822, int32(160), int32(0))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L28
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v989 = F_dsa_allocate_extended(m, v822, int32(512), int32(0))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L28
	} else {
		goto L155
	}
L141:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v834 = F_dsa_get_address(m, v833, v831)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L28
	} else {
		goto L142
	}
L142:
	;
	v836 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v834)+2)) = v836
	v838 = int32(7937)
	*(*uint16)(unsafe.Add(mBase, uint32(v834))) = uint16(v838)
	v840 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v834)+34)) = uint16(v840)
	*(*int64)(unsafe.Add(mBase, uint32(v834)+26)) = v836
	*(*int64)(unsafe.Add(mBase, uint32(v834)+18)) = v836
	*(*int64)(unsafe.Add(mBase, uint32(v834)+10)) = v836
	v848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v825)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v834)+2)) = uint8(v848)
	v851 = v825 + int32(3)
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848+v851-int32(1)))))
	if base.Ui32(v819) <= base.Ui32(v855) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v926 = v834 + int32(2)
	v927 = int32(36)
	v928 = v825 + v927
	v930 = v834 + v927
	v932 = v834 + int32(3)
	v937 = v820
	goto L152
L144:
	;
	if v848 == int32(0) {
		v906 = v820
		goto L143
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v906 = v848
	goto L143
L147:
	;
	v860 = v820
	goto L148
L148:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860+v851))))
	if base.Ui32(v819) < base.Ui32(v880) {
		v906 = v860
		goto L143
	} else {
		goto L150
	}
L149:
	;
	goto L146
L150:
	;
	v883 = v860 + int32(1)
	if v883 != v848 {
		v860 = v883
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	v954 = v937 | base.B2i32(base.Ui32(v906) <= base.Ui32(v937))
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937+v851))))
	*(*uint8)(unsafe.Add(mBase, uint32(v932+v954))) = uint8(v957)
	v959 = int32(2)
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v928+v937<<(uint(v959)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v930+v954<<(uint(v959)%32)))) = v965
	if v937 == int32(14) {
		goto L137
	} else {
		goto L154
	}
L154:
	;
	v970 = v937 | int32(1)
	v972 = v970 + base.B2i32(base.Ui32(v906) <= base.Ui32(v970))
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970+v851))))
	*(*uint8)(unsafe.Add(mBase, uint32(v932+v972))) = uint8(v975)
	v977 = int32(2)
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v928+v970<<(uint(v977)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v930+v972<<(uint(v977)%32)))) = v983
	v937 = v937 + v977
	goto L152
L155:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v992 = F_dsa_get_address(m, v991, v989)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L28
	} else {
		goto L156
	}
L156:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v992))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v992)+8)) = int32(0)
	v1003 = F__emscripten_memset_bulkmem(m, v992+int32(12), base.I32_extend8_s(int32(255)), int32(256))
	mBase = m.M
	goto L157
L157:
	;
	v1004 = int32(15618)
	*(*uint16)(unsafe.Add(mBase, uint32(v992))) = uint16(v1004)
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v825)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v992)+2)) = uint8(v1006)
	v1009 = v825 + int32(3)
	v1014 = v820
	goto L158
L158:
	;
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009+v1014))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1003+v1031))) = uint8(v1014)
	v1035 = v1014 | int32(1)
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009+v1035))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1003+v1037))) = uint8(v1035)
	v1041 = v1014 | int32(2)
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009+v1041))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1003+v1043))) = uint8(v1041)
	if v1014 == int32(28) {
		goto L136
	} else {
		goto L160
	}
L160:
	;
	v1049 = v1014 | int32(3)
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009+v1049))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1003+v1051))) = uint8(v1049)
	v1014 = v1014 + int32(4)
	goto L158
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v831
	v1112 = v930 + v906<<(uint(int32(2))%32)
	goto L135
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v992)+4)) = int32(-1)
	v1080 = int32(31)
	*(*uint8)(unsafe.Add(mBase, uint32(v1003+v819))) = uint8(v1080)
	v1083 = v1006 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v992)+2)) = uint8(v1083)
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v989
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v824)))
	F_dsa_free(m, v1086, v1087)
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L28
	} else {
		goto L166
	}
L163:
	;
	v1075 = F__emscripten_memcpy_bulkmem(m, v992+int32(268), v825+int32(36), int32(124))
	mBase = m.M
	goto L165
L165:
	;
	goto L162
L166:
	;
	v1112 = v992 + int32(392)
	goto L135
L167:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1246+v1114))) = uint8(v332)
	v1265 = v1250 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)) = uint8(v1265)
	v3049 = v335 + v1246<<(uint(int32(2))%32) + int32(36)
	goto L3
L168:
	;
	v1246 = v813
	v1250 = v800
	goto L167
L169:
	;
	goto L170
L170:
	;
	v1118 = v332 & int32(255)
	v1120 = v813 - int32(1)
	v1121 = v1114 + v1120
	v1122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1121))))
	if base.Ui32(v1122) < base.Ui32(v1118) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v1246 = v813
	v1250 = v800
	goto L167
L172:
	;
	goto L173
L173:
	;
	v1129 = int32(0)
	goto L175
L174:
	;
	if base.Ui32(v813) <= base.Ui32(v1129) {
		v1246 = v1129
		v1250 = v800
		goto L167
	} else {
		goto L179
	}
L175:
	;
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129+v1114))))
	if base.Ui32(v1118) < base.Ui32(v1146) {
		goto L174
	} else {
		goto L177
	}
L176:
	;
	v1246 = v813
	v1250 = v800
	goto L167
L177:
	;
	v1149 = v1129 + int32(1)
	if v1149 != v813 {
		v1129 = v1149
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	v1153 = v335 + int32(36)
	v1154 = v813 - v1129
	if v1154&int32(1) == int32(0) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	if v1154 != int32(1) {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	v1172 = v1120
	v1173 = v813
	goto L180
L182:
	;
	goto L183
L183:
	;
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1121))))
	*(*uint8)(unsafe.Add(mBase, uint32(v813+v1114))) = uint8(v1160)
	v1162 = int32(2)
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1153+v1120<<(uint(v1162)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1153+v813<<(uint(v1162)%32)))) = v1168
	v1172 = v813 - v1162
	v1173 = v1120
	goto L180
L184:
	;
	v1184 = v1172
	v1189 = v1173
	goto L187
L185:
	;
	goto L186
L186:
	;
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	v1246 = v1129
	v1250 = v1241
	goto L167
L187:
	;
	v1197 = v1114 + v1184
	v1198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1197))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1114+v1189))) = uint8(v1198)
	v1200 = int32(2)
	v1205 = v1153 + v1184<<(uint(v1200)%32)
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1205)))
	*(*int32)(unsafe.Add(mBase, uint32(v1153+v1189<<(uint(v1200)%32)))) = v1206
	v1209 = v1184 - int32(1)
	v1211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1114+v1209))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1197))) = uint8(v1211)
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1153+v1209<<(uint(v1200)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1205))) = v1216
	if v1129 < v1209 {
		v1184 = v1184 - v1200
		v1189 = v1209
		goto L187
	} else {
		goto L189
	}
L188:
	;
	goto L186
L189:
	;
	goto L188
L190:
	;
	v321 = v321 - int32(8)
	v323 = v478
	goto L56
L191:
	;
	goto L192
L192:
	;
	goto L57
L193:
	;
	if v1319 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	goto L195
L195:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v1319 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L196:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	F_dsa_free(m, v1324, v1317)
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L28
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	if v165 != 0 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	goto L198
L200:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1330)))
	*(*int32)(unsafe.Add(mBase, uint32(v1330))) = v1331 | int32(1)
	goto L1
L201:
	;
	v1329 = F__emscripten_memcpy_bulkmem(m, v478, v23-int32(-64), v165)
	mBase = m.M
	v1330 = v1329
	goto L203
L202:
	;
	v1330 = v478
	goto L203
L203:
	;
	goto L200
L204:
	;
	v1338 = F_dsa_get_address(m, v1335, v1317)
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L28
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v1363 = F_dsa_allocate_extended(m, v1335, v165, int32(0))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L28
	} else {
		goto L218
	}
L207:
	;
	v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1338)+1)))
	if v1340 != v152&int32(255) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v478)))
	F_dsa_free(m, v1344, v1345)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L28
	} else {
		goto L211
	}
L209:
	;
	v1357 = v1338
	goto L210
L210:
	;
	if v165 != 0 {
		goto L215
	} else {
		goto L216
	}
L211:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v1350 = F_dsa_allocate_extended(m, v1348, v165, int32(0))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L28
	} else {
		goto L212
	}
L212:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v1353 = F_dsa_get_address(m, v1352, v1350)
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L28
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478))) = v1350
	v1357 = v1353
	goto L210
L214:
	;
	goto L1
L215:
	;
	v1360 = F__emscripten_memcpy_bulkmem(m, v1357, v23-int32(-64), v165)
	mBase = m.M
	goto L217
L216:
	;
	goto L217
L217:
	;
	goto L214
L218:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v1366 = F_dsa_get_address(m, v1365, v1363)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L28
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478))) = v1363
	if v165 != 0 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	goto L1
L221:
	;
	v1371 = F__emscripten_memcpy_bulkmem(m, v1366, v23-int32(-64), v165)
	mBase = m.M
	goto L223
L222:
	;
	goto L223
L223:
	;
	goto L220
L224:
	;
	v1375 = *(*int64)(unsafe.Add(mBase, uint32(v168)+16))
	if v1375 == int64(0) {
		goto L227
	} else {
		goto L228
	}
L225:
	;
	v1491 = v168
	goto L226
L226:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1491)+24))
	v1512 = v1491
	v1517 = v1503
	goto L241
L227:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v1379 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1378)+2)) = uint8(v1379)
	v1385 = (base.I32_clz(l1) ^ int32(-1)) & int32(24)
	v1386 = int32(base.Ui32(l1) >> (uint(v1385) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v1378)+3)) = uint8(v1386)
	v1390 = F_local_ts_extend_down(m, v167, v1378+int32(8), v166, v1385)
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L28
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+24))
	if v166 == int64(0) {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v1392)+24)) = v1385
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v1395 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1394)+8)) = v1395<<(uint(base.I64_extend_i32_u(v1385+int32(8)))%64) ^ v1395
	v2771 = v1390
	goto L5
L231:
	;
	v1414 = int32(0)
	goto L233
L232:
	;
	v1414 = (base.I32_wrap_i64(base.I64_clz(v166)) ^ int32(-1)) & int32(56)
	goto L233
L233:
	;
	if v1404 < v1414 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1419 = v1404
	goto L237
L235:
	;
	v1471 = v1403
	goto L236
L236:
	;
	v1472 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1471)+8)) = v1472<<(uint(base.I64_extend_i32_u(v1414+int32(8)))%64) ^ v1472
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v1480)+24)) = v1414
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v1491 = v1482
	goto L226
L237:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v1438 = F_MemoryContextAlloc(m, v1436, int32(24))
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L28
	} else {
		goto L239
	}
L238:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v1471 = v1450
	goto L236
L239:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1438))) = int64(66560)
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1442)))
	*(*int32)(unsafe.Add(mBase, uint32(v1438)+8)) = v1443
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v1445))) = v1438
	v1448 = v1419 + int32(8)
	if v1448 < v1414 {
		v1419 = v1448
		goto L237
	} else {
		goto L240
	}
L240:
	;
	goto L238
L241:
	;
	v1526 = base.I32_wrap_i64(int64(base.Ui64(v166) >> (uint(base.I64_extend_i32_u(v1517)) % 64)))
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1512)))
	v1528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527))))
	switch v1528 - int32(1) {
	case 0:
		goto L256
	case 1:
		goto L255
	case 2:
		goto L254
	default:
		goto L257
	}
L242:
	;
	v2483 = *(*int32)(unsafe.Add(mBase, uint32(v1664)))
	v2485 = v2483 & int32(1)
	if base.Ui32(v165) <= base.Ui32(int32(4)) {
		goto L372
	} else {
		goto L373
	}
L243:
	;
	if v1517 != 0 {
		goto L369
	} else {
		goto L370
	}
L244:
	;
	v2464 = int32(2)
	v2468 = v1527 + v2451<<(uint(v2464)%32) + int32(4)
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v2468)))
	*(*int32)(unsafe.Add(mBase, uint32(v2468))) = v2469 | v2450
	v2472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527)+2)))
	v2474 = v2472 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1527)+2)) = uint8(v2474)
	v2762 = v1527 + v2448<<(uint(v2464)%32) + int32(36)
	goto L6
L245:
	;
	v2441 = v1526 & int32(255)
	v2448 = v2441
	v2450 = int32(1) << (uint(v1526) % 32)
	v2451 = int32(base.Ui32(v2441) >> (uint(int32(5)) % 32))
	goto L244
L246:
	;
	v1991 = v1977 & int32(255)
	v1992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527)+1)))
	if v1991 == v1992 {
		goto L315
	} else {
		goto L316
	}
L247:
	;
	v1969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527)+2)))
	v1977 = v1969
	goto L246
L248:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1943+v1527)+3)) = uint8(v1526)
	v1962 = v1950 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1527)+2)) = uint8(v1962)
	v2762 = v1527 + v1943<<(uint(int32(2))%32) + int32(8)
	goto L6
L249:
	;
	v1820 = v1527 + int32(3)
	v1828 = int32(0)
	goto L298
L250:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v23)+328))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v1719
	v1722 = v1526 & int32(255)
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(12))))
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v1729 = F_MemoryContextAlloc(m, v1727, int32(100))
	mBase = m.M
	v1730 = m.ExcPending
	if v1730 != 0 {
		goto L28
	} else {
		goto L281
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v1527
	v1696 = int32(0)
	v1698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527)+1)))
	if v1698 != 0 {
		v1943 = v1696
		v1950 = v1696
		goto L248
	} else {
		goto L280
	}
L252:
	;
	v1690 = v1676 & int32(255)
	v1691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527)+1)))
	if v1690 == v1691 {
		goto L250
	} else {
		goto L278
	}
L253:
	;
	if v1664 != 0 {
		goto L243
	} else {
		goto L276
	}
L254:
	;
	v1626 = int32(1) << (uint(v1526) % 32)
	v1628 = v1526 & int32(255)
	v1630 = int32(base.Ui32(v1628) >> (uint(int32(5)) % 32))
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v1527+v1630<<(uint(int32(2))%32))+4))
	if v1626&v1634 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L255:
	;
	v1613 = int32(255)
	v1616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527+v1526&v1613)+12)))
	if v1616 == v1613 {
		goto L270
	} else {
		goto L271
	}
L256:
	;
	v1571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527)+2)))
	if v1571 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L257:
	;
	v1531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527)+2)))
	if v1531 == int32(0) {
		goto L251
	} else {
		goto L258
	}
L258:
	;
	v1543 = int32(0)
	goto L259
L259:
	;
	v1565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1543+(v1527+int32(3))))))
	if v1526&int32(255) == v1565 {
		v1664 = v1527 + v1543<<(uint(int32(2))%32) + int32(8)
		goto L253
	} else {
		goto L261
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v1527
	v1676 = v1531
	goto L252
L261:
	;
	v1568 = v1543 + int32(1)
	if v1568 != v1531 {
		v1543 = v1568
		goto L259
	} else {
		goto L262
	}
L262:
	;
	goto L260
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v1527
	v1977 = int32(0)
	goto L246
L264:
	;
	goto L265
L265:
	;
	v1583 = int32(0)
	goto L266
L266:
	;
	v1605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1583+(v1527+int32(3))))))
	if v1605 == v1526&int32(255) {
		v1664 = v1527 + v1583<<(uint(int32(2))%32) + int32(36)
		goto L253
	} else {
		goto L268
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v1527
	v1977 = v1571
	goto L246
L268:
	;
	v1610 = v1583 + int32(1)
	if v1610 != v1571 {
		v1583 = v1610
		goto L266
	} else {
		goto L269
	}
L269:
	;
	goto L267
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v1527
	goto L7
L271:
	;
	goto L272
L272:
	;
	v1664 = v1527 + v1616<<(uint(int32(2))%32) + int32(268)
	goto L253
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v1527
	v2448 = v1628
	v2450 = v1626
	v2451 = v1630
	goto L244
L274:
	;
	goto L275
L275:
	;
	v1664 = v1527 + v1628<<(uint(int32(2))%32) + int32(36)
	goto L253
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v1527
	switch v1528 - int32(1) {
	case 0:
		goto L247
	case 1:
		goto L7
	case 2:
		goto L245
	default:
		goto L277
	}
L277:
	;
	v1668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527)+2)))
	v1676 = v1668
	goto L252
L278:
	;
	if v1690 != 0 {
		goto L249
	} else {
		goto L279
	}
L279:
	;
	v1943 = int32(0)
	v1950 = int32(0)
	goto L248
L280:
	;
	goto L250
L281:
	;
	v1731 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1729)+2)) = v1731
	v1733 = int32(4097)
	*(*uint16)(unsafe.Add(mBase, uint32(v1729))) = uint16(v1733)
	v1735 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1729)+34)) = uint16(v1735)
	*(*int64)(unsafe.Add(mBase, uint32(v1729)+26)) = v1731
	*(*int64)(unsafe.Add(mBase, uint32(v1729)+18)) = v1731
	*(*int64)(unsafe.Add(mBase, uint32(v1729)+10)) = v1731
	v1743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1729)+2)) = uint8(v1743)
	v1745 = int32(2)
	v1746 = int32(4)
	v1748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726)+3)))
	v1749 = base.B2i32(base.Ui32(v1722) <= base.Ui32(v1748))
	if base.Ui32(v1722) <= base.Ui32(v1748) {
		v1769 = v1746
		v1770 = int32(0)
		v1771 = v1745
		v1772 = int32(3)
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1774 = v1729 + int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v1774+v1749))) = uint8(v1748)
	v1778 = v1729 + int32(36)
	v1779 = int32(2)
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1778+v1749<<(uint(v1779)%32)))) = v1782
	v1785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1774+v1771))) = uint8(v1785)
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1778+v1771<<(uint(v1779)%32)))) = v1790
	v1793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1774+v1772))) = uint8(v1793)
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1778+v1772<<(uint(v1779)%32)))) = v1798
	v1801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1774+v1769))) = uint8(v1801)
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1778+v1769<<(uint(v1779)%32)))) = v1806
	*(*uint8)(unsafe.Add(mBase, uint32(v1774+v1770))) = uint8(v1722)
	v1811 = v1743 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1729)+2)) = uint8(v1811)
	*(*int32)(unsafe.Add(mBase, uint32(v1512))) = v1729
	F_pfree(m, v1726)
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L28
	} else {
		goto L296
	}
L283:
	;
	v1750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726)+4)))
	if base.Ui32(v1722) <= base.Ui32(v1750) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1769 = v1746
	v1770 = int32(1)
	v1771 = v1745
	v1772 = int32(3)
	goto L282
L285:
	;
	goto L286
L286:
	;
	v1754 = int32(1)
	v1755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726)+5)))
	if base.Ui32(v1722) <= base.Ui32(v1755) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1769 = v1746
	v1770 = int32(2)
	v1771 = v1754
	v1772 = int32(3)
	goto L282
L288:
	;
	goto L289
L289:
	;
	v1761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726)+6)))
	v1762 = base.B2i32(base.Ui32(v1761) < base.Ui32(v1722))
	if base.Ui32(v1761) < base.Ui32(v1722) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1763 = int32(4)
	goto L292
L291:
	;
	v1763 = int32(3)
	goto L292
L292:
	;
	if base.Ui32(v1761) < base.Ui32(v1722) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1766 = int32(3)
	goto L295
L294:
	;
	v1766 = int32(4)
	goto L295
L295:
	;
	v1769 = v1766
	v1770 = v1763
	v1771 = v1754
	v1772 = int32(2)
	goto L282
L296:
	;
	v2762 = v1778 + v1770<<(uint(int32(2))%32)
	goto L6
L297:
	;
	if base.Ui32(v1690) <= base.Ui32(v1828) {
		goto L302
	} else {
		goto L303
	}
L298:
	;
	v1845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1828+v1820))))
	if base.Ui32(v1526&int32(255)) <= base.Ui32(v1845) {
		goto L297
	} else {
		goto L300
	}
L299:
	;
	v1943 = v1690
	v1950 = v1676
	goto L248
L300:
	;
	v1848 = v1828 + int32(1)
	if v1848 != v1690 {
		v1828 = v1848
		goto L298
	} else {
		goto L301
	}
L301:
	;
	goto L299
L302:
	;
	v1943 = v1828
	v1950 = v1676
	goto L248
L303:
	;
	goto L304
L304:
	;
	v1852 = v1527 + int32(8)
	v1853 = v1690 - v1828
	if v1853&int32(1) == int32(0) {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	if v1853 != int32(1) {
		goto L309
	} else {
		goto L310
	}
L306:
	;
	v1872 = v1690
	goto L305
L307:
	;
	goto L308
L308:
	;
	v1860 = v1690 - int32(1)
	v1862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1820+v1860))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1690+v1820))) = uint8(v1862)
	v1864 = int32(2)
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(v1852+v1860<<(uint(v1864)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1852+v1690<<(uint(v1864)%32)))) = v1870
	v1872 = v1860
	goto L305
L309:
	;
	v1882 = v1872
	goto L312
L310:
	;
	goto L311
L311:
	;
	v1938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527)+2)))
	v1943 = v1828
	v1950 = v1938
	goto L248
L312:
	;
	v1897 = v1882 - int32(1)
	v1898 = v1820 + v1897
	v1899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1898))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1882+v1820))) = uint8(v1899)
	v1901 = int32(2)
	v1902 = v1882 << (uint(v1901) % 32)
	v1906 = v1852 + v1897<<(uint(v1901)%32)
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1906)))
	*(*int32)(unsafe.Add(mBase, uint32(v1852+v1902))) = v1907
	v1910 = v1882 - v1901
	v1912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1820+v1910))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1898))) = uint8(v1912)
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v1527+v1902)))
	*(*int32)(unsafe.Add(mBase, uint32(v1906))) = v1915
	if v1828 < v1910 {
		v1882 = v1910
		goto L312
	} else {
		goto L314
	}
L313:
	;
	goto L311
L314:
	;
	goto L313
L315:
	;
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v23)+328))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v1994
	v1997 = v1526 & int32(255)
	v1998 = int32(0)
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(16))))
	v2002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2001)+1)))
	if base.Ui32(v2002) <= base.Ui32(int32(31)) {
		goto L319
	} else {
		goto L320
	}
L316:
	;
	goto L317
L317:
	;
	v2280 = v1527 + int32(3)
	if v1991 == int32(0) {
		goto L347
	} else {
		goto L348
	}
L318:
	;
	v2762 = v2278
	goto L6
L319:
	;
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v2007 = F_MemoryContextAlloc(m, v2005, int32(164))
	mBase = m.M
	v2008 = m.ExcPending
	if v2008 != 0 {
		goto L28
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v2173 = F_MemoryContextAlloc(m, v2171, int32(524))
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L28
	} else {
		goto L336
	}
L322:
	;
	v2009 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2007)+2)) = v2009
	v2011 = int32(8193)
	*(*uint16)(unsafe.Add(mBase, uint32(v2007))) = uint16(v2011)
	v2013 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2007)+34)) = uint16(v2013)
	*(*int64)(unsafe.Add(mBase, uint32(v2007)+26)) = v2009
	*(*int64)(unsafe.Add(mBase, uint32(v2007)+18)) = v2009
	*(*int64)(unsafe.Add(mBase, uint32(v2007)+10)) = v2009
	v2022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2001)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2007)+2)) = uint8(v2022)
	v2025 = v2001 + int32(3)
	v2029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2022+v2025-int32(1)))))
	if base.Ui32(v1997) <= base.Ui32(v2029) {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v2099 = int32(36)
	v2100 = v2001 + v2099
	v2102 = v2007 + v2099
	v2104 = v2007 + int32(3)
	v2108 = v2013
	goto L332
L324:
	;
	if v2022 == int32(0) {
		v2080 = v1998
		goto L323
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	v2080 = v2022
	goto L323
L327:
	;
	v2034 = v1998
	goto L328
L328:
	;
	v2054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2034+v2025))))
	if base.Ui32(v1997) < base.Ui32(v2054) {
		v2080 = v2034
		goto L323
	} else {
		goto L330
	}
L329:
	;
	goto L326
L330:
	;
	v2057 = v2034 + int32(1)
	if v2057 != v2022 {
		v2034 = v2057
		goto L328
	} else {
		goto L331
	}
L331:
	;
	goto L329
L332:
	;
	v2126 = v2108 | base.B2i32(base.Ui32(v2080) <= base.Ui32(v2108))
	v2129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2108+v2025))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2104+v2126))) = uint8(v2129)
	v2131 = int32(2)
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(v2100+v2108<<(uint(v2131)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2102+v2126<<(uint(v2131)%32)))) = v2137
	v2140 = v2108 | int32(1)
	v2142 = v2140 + base.B2i32(base.Ui32(v2080) <= base.Ui32(v2140))
	v2145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2140+v2025))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2104+v2142))) = uint8(v2145)
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v2100+v2140<<(uint(v2131)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2102+v2142<<(uint(v2131)%32)))) = v2153
	v2156 = v2108 + v2131
	if v2156 != int32(16) {
		v2108 = v2156
		goto L332
	} else {
		goto L334
	}
L333:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2080+v2104))) = uint8(v1997)
	v2161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2007)+2)))
	v2163 = v2161 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2007)+2)) = uint8(v2163)
	F_pfree(m, v2001)
	mBase = m.M
	v2166 = m.ExcPending
	if v2166 != 0 {
		goto L28
	} else {
		goto L335
	}
L334:
	;
	goto L333
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1512))) = v2007
	v2278 = v2102 + v2080<<(uint(int32(2))%32)
	goto L318
L336:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = int64(0)
	v2177 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+8)) = v2177
	v2185 = F__emscripten_memset_bulkmem(m, v2173+int32(12), base.I32_extend8_s(int32(255)), int32(256))
	mBase = m.M
	goto L337
L337:
	;
	v2186 = int32(16386)
	*(*uint16)(unsafe.Add(mBase, uint32(v2173))) = uint16(v2186)
	v2188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2001)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2173)+2)) = uint8(v2188)
	v2191 = v2001 + int32(3)
	v2195 = v2177
	goto L338
L338:
	;
	v2213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2191+v2195))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2185+v2213))) = uint8(v2195)
	v2217 = v2195 | int32(1)
	v2219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2191+v2217))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2185+v2219))) = uint8(v2217)
	v2223 = v2195 | int32(2)
	v2225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2191+v2223))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2185+v2225))) = uint8(v2223)
	v2229 = v2195 | int32(3)
	v2231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2191+v2229))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2185+v2231))) = uint8(v2229)
	v2235 = v2195 + int32(4)
	if v2235 != int32(32) {
		v2195 = v2235
		goto L338
	} else {
		goto L340
	}
L339:
	;
	goto L342
L340:
	;
	goto L339
L341:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2173)+4)) = int64(8589934591)
	v2248 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v2185+v1997))) = uint8(v2248)
	v2251 = v2188 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2173)+2)) = uint8(v2251)
	*(*int32)(unsafe.Add(mBase, uint32(v1512))) = v2173
	F_pfree(m, v2001)
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L28
	} else {
		goto L345
	}
L342:
	;
	v2243 = F__emscripten_memcpy_bulkmem(m, v2173+int32(268), v2001+int32(36), int32(128))
	mBase = m.M
	goto L344
L344:
	;
	goto L341
L345:
	;
	v2278 = v2173 + int32(396)
	goto L318
L346:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2412+v2280))) = uint8(v1526)
	v2431 = v2415 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1527)+2)) = uint8(v2431)
	v2762 = v1527 + v2412<<(uint(int32(2))%32) + int32(36)
	goto L6
L347:
	;
	v2412 = v1991
	v2415 = v1977
	goto L346
L348:
	;
	goto L349
L349:
	;
	v2284 = v1526 & int32(255)
	v2286 = v1991 - int32(1)
	v2287 = v2280 + v2286
	v2288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2287))))
	if base.Ui32(v2288) < base.Ui32(v2284) {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v2412 = v1991
	v2415 = v1977
	goto L346
L351:
	;
	goto L352
L352:
	;
	v2295 = int32(0)
	goto L354
L353:
	;
	if base.Ui32(v1991) <= base.Ui32(v2295) {
		v2412 = v2295
		v2415 = v1977
		goto L346
	} else {
		goto L358
	}
L354:
	;
	v2312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2295+v2280))))
	if base.Ui32(v2284) < base.Ui32(v2312) {
		goto L353
	} else {
		goto L356
	}
L355:
	;
	v2412 = v1991
	v2415 = v1977
	goto L346
L356:
	;
	v2315 = v2295 + int32(1)
	if v2315 != v1991 {
		v2295 = v2315
		goto L354
	} else {
		goto L357
	}
L357:
	;
	goto L355
L358:
	;
	v2319 = v1527 + int32(36)
	v2320 = v1991 - v2295
	if v2320&int32(1) == int32(0) {
		goto L360
	} else {
		goto L361
	}
L359:
	;
	if v2320 != int32(1) {
		goto L363
	} else {
		goto L364
	}
L360:
	;
	v2338 = v2286
	v2339 = v1991
	goto L359
L361:
	;
	goto L362
L362:
	;
	v2326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2287))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1991+v2280))) = uint8(v2326)
	v2328 = int32(2)
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(v2319+v2286<<(uint(v2328)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2319+v1991<<(uint(v2328)%32)))) = v2334
	v2338 = v1991 - v2328
	v2339 = v2286
	goto L359
L363:
	;
	v2349 = v2338
	v2353 = v2339
	goto L366
L364:
	;
	goto L365
L365:
	;
	v2407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527)+2)))
	v2412 = v2295
	v2415 = v2407
	goto L346
L366:
	;
	v2363 = v2349 + v2280
	v2364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2363))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2280+v2353))) = uint8(v2364)
	v2366 = int32(2)
	v2371 = v2319 + v2349<<(uint(v2366)%32)
	v2372 = *(*int32)(unsafe.Add(mBase, uint32(v2371)))
	*(*int32)(unsafe.Add(mBase, uint32(v2319+v2353<<(uint(v2366)%32)))) = v2372
	v2375 = v2349 - int32(1)
	v2377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2280+v2375))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2363))) = uint8(v2377)
	v2382 = *(*int32)(unsafe.Add(mBase, uint32(v2319+v2375<<(uint(v2366)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2371))) = v2382
	if v2295 < v2375 {
		v2349 = v2349 - v2366
		v2353 = v2375
		goto L366
	} else {
		goto L368
	}
L367:
	;
	goto L365
L368:
	;
	goto L367
L369:
	;
	v1512 = v1664
	v1517 = v1517 - int32(8)
	goto L241
L370:
	;
	goto L371
L371:
	;
	goto L242
L372:
	;
	if v2485 == int32(0) {
		goto L375
	} else {
		goto L376
	}
L373:
	;
	goto L374
L374:
	;
	if v2485 == int32(0) {
		goto L383
	} else {
		goto L384
	}
L375:
	;
	F_pfree(m, v2483)
	mBase = m.M
	v2491 = m.ExcPending
	if v2491 != 0 {
		goto L28
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	if v165 != 0 {
		goto L380
	} else {
		goto L381
	}
L378:
	;
	goto L377
L379:
	;
	v2496 = *(*int32)(unsafe.Add(mBase, uint32(v2495)))
	*(*int32)(unsafe.Add(mBase, uint32(v2495))) = v2496 | int32(1)
	goto L1
L380:
	;
	v2494 = F__emscripten_memcpy_bulkmem(m, v1664, v23-int32(-64), v165)
	mBase = m.M
	v2495 = v2494
	goto L382
L381:
	;
	v2495 = v1664
	goto L382
L382:
	;
	goto L379
L383:
	;
	v2502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2483)+1)))
	if v2502 != v152&int32(255) {
		goto L386
	} else {
		goto L387
	}
L384:
	;
	goto L385
L385:
	;
	v2517 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v2518 = F_MemoryContextAlloc(m, v2517, v165)
	mBase = m.M
	v2519 = m.ExcPending
	if v2519 != 0 {
		goto L28
	} else {
		goto L395
	}
L386:
	;
	F_pfree(m, v2483)
	mBase = m.M
	v2507 = m.ExcPending
	if v2507 != 0 {
		goto L28
	} else {
		goto L389
	}
L387:
	;
	v2512 = v2483
	goto L388
L388:
	;
	if v165 != 0 {
		goto L392
	} else {
		goto L393
	}
L389:
	;
	v2508 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v2509 = F_MemoryContextAlloc(m, v2508, v165)
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L28
	} else {
		goto L390
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1664))) = v2509
	v2512 = v2509
	goto L388
L391:
	;
	goto L1
L392:
	;
	v2515 = F__emscripten_memcpy_bulkmem(m, v2512, v23-int32(-64), v165)
	mBase = m.M
	goto L394
L393:
	;
	goto L394
L394:
	;
	goto L391
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1664))) = v2518
	if v165 != 0 {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	goto L1
L397:
	;
	v2523 = F__emscripten_memcpy_bulkmem(m, v2518, v23-int32(-64), v165)
	mBase = m.M
	goto L399
L398:
	;
	goto L399
L399:
	;
	goto L396
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v132
	F_errmsg_internal(m, int32(58754), v23)
	mBase = m.M
	v2532 = m.ExcPending
	if v2532 != 0 {
		goto L28
	} else {
		goto L401
	}
L401:
	;
	F_errfinish(m, int32(498501), int32(375), int32(123995))
	mBase = m.M
	v2537 = m.ExcPending
	if v2537 != 0 {
		goto L28
	} else {
		goto L402
	}
L402:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L403:
	;
	v2561 = *(*int32)(unsafe.Add(mBase, uint32(v23)+328))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v2561
	v2564 = v1526 & int32(255)
	v2568 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(20))))
	v2569 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	v2571 = F_MemoryContextAlloc(m, v2569, int32(1060))
	mBase = m.M
	v2572 = m.ExcPending
	if v2572 != 0 {
		goto L28
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	v2709 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+4))
	if v2709 == int32(-1) {
		goto L420
	} else {
		goto L421
	}
L406:
	;
	v2573 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2571)+2)) = v2573
	v2575 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v2571))) = uint16(v2575)
	v2577 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2571)+34)) = uint16(v2577)
	*(*int64)(unsafe.Add(mBase, uint32(v2571)+26)) = v2573
	*(*int64)(unsafe.Add(mBase, uint32(v2571)+18)) = v2573
	*(*int64)(unsafe.Add(mBase, uint32(v2571)+10)) = v2573
	v2586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2568)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2571)+2)) = uint8(v2586)
	v2589 = v2568 + int32(268)
	v2591 = v2568 + int32(12)
	v2593 = v2571 + int32(4)
	v2595 = v2571 + int32(36)
	v2596 = v2577
	v2605 = int32(0)
	goto L407
L407:
	;
	v2616 = int32(0)
	v2618 = v2596
	v2621 = v2616
	v2622 = v2616
	goto L409
L408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1512))) = v2571
	F_pfree(m, v2568)
	mBase = m.M
	v2688 = m.ExcPending
	if v2688 != 0 {
		goto L28
	} else {
		goto L419
	}
L409:
	;
	v2639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2618+v2591))))
	if v2639 != int32(255) {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2593+v2605<<(uint(int32(2))%32)))) = v2671
	v2683 = v2605 + int32(1)
	if v2683 != int32(8) {
		v2596 = v2673
		v2605 = v2683
		goto L407
	} else {
		goto L418
	}
L411:
	;
	v2642 = int32(2)
	v2648 = *(*int32)(unsafe.Add(mBase, uint32(v2589+v2639<<(uint(v2642)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2595+v2618<<(uint(v2642)%32)))) = v2648
	v2653 = int32(1)<<(uint(v2622)%32) | v2621
	goto L413
L412:
	;
	v2653 = v2621
	goto L413
L413:
	;
	v2655 = v2618 + int32(1)
	v2657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2591+v2655))))
	if v2657 != int32(255) {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v2660 = int32(2)
	v2666 = *(*int32)(unsafe.Add(mBase, uint32(v2589+v2657<<(uint(v2660)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2595+v2655<<(uint(v2660)%32)))) = v2666
	v2671 = v2660<<(uint(v2622)%32) | v2653
	goto L416
L415:
	;
	v2671 = v2653
	goto L416
L416:
	;
	v2672 = int32(2)
	v2673 = v2618 + v2672
	v2675 = v2622 + v2672
	if v2675 != int32(32) {
		v2618 = v2673
		v2621 = v2671
		v2622 = v2675
		goto L409
	} else {
		goto L417
	}
L417:
	;
	goto L410
L418:
	;
	goto L408
L419:
	;
	v2693 = v2593 + int32(base.Ui32(v2564)>>(uint(int32(3))%32))&int32(28)
	v2694 = *(*int32)(unsafe.Add(mBase, uint32(v2693)))
	v2695 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2693))) = v2694 | v2695<<(uint(v2564)%32)
	v2699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2571)+2)))
	v2701 = v2699 + v2695
	*(*uint8)(unsafe.Add(mBase, uint32(v2571)+2)) = uint8(v2701)
	v2762 = v2595 + v2564<<(uint(int32(2))%32)
	goto L6
L420:
	;
	v2712 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+8))
	v2715 = v2712
	v2716 = base.B2i32(v2712 != int32(-1))
	goto L422
L421:
	;
	v2715 = v2709
	v2716 = int32(0)
	goto L422
L422:
	;
	v2717 = int32(2)
	v2720 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1527+int32(4)+v2716<<(uint(v2717)%32)))) = v2715 + v2720 | v2715
	v2732 = base.I32_ctz(v2715^int32(-1)) + v2716<<(uint(int32(5))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1527+v1526&int32(255))+12)) = uint8(v2732)
	v2735 = v2558 + v2720
	*(*uint8)(unsafe.Add(mBase, uint32(v1527)+2)) = uint8(v2735)
	v2762 = v1527 + v2732<<(uint(v2717)%32) + int32(268)
	goto L6
L423:
	;
	v2765 = F_local_ts_extend_down(m, v167, v2762, v166, v1517)
	mBase = m.M
	v2766 = m.ExcPending
	if v2766 != 0 {
		goto L28
	} else {
		goto L424
	}
L424:
	;
	v2771 = v2765
	goto L5
L425:
	;
	v2812 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v2813 = *(*int64)(unsafe.Add(mBase, uint32(v2812)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2812)+16)) = v2813 + int64(1)
	goto L1
L426:
	;
	v2794 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v2795 = F_MemoryContextAlloc(m, v2794, v165)
	mBase = m.M
	v2796 = m.ExcPending
	if v2796 != 0 {
		goto L28
	} else {
		goto L429
	}
L427:
	;
	goto L428
L428:
	;
	if v165 != 0 {
		goto L435
	} else {
		goto L436
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2771))) = v2795
	if v165 != 0 {
		goto L431
	} else {
		goto L432
	}
L430:
	;
	goto L425
L431:
	;
	v2800 = F__emscripten_memcpy_bulkmem(m, v2795, v23-int32(-64), v165)
	mBase = m.M
	goto L433
L432:
	;
	goto L433
L433:
	;
	goto L430
L434:
	;
	v2806 = *(*int32)(unsafe.Add(mBase, uint32(v2805)))
	*(*int32)(unsafe.Add(mBase, uint32(v2805))) = v2806 | int32(1)
	goto L425
L435:
	;
	v2804 = F__emscripten_memcpy_bulkmem(m, v2771, v23-int32(-64), v165)
	mBase = m.M
	v2805 = v2804
	goto L437
L436:
	;
	v2805 = v2771
	goto L437
L437:
	;
	goto L434
L438:
	;
	v2840 = *(*int64)(unsafe.Add(mBase, uint32(v23)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v2840
	v2843 = v332 & int32(255)
	v2844 = int32(0)
	v2847 = v23 + int32(40)
	v2848 = *(*int32)(unsafe.Add(mBase, uint32(v2847)+4))
	v2849 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v2852 = F_dsa_allocate_extended(m, v2849, int32(1060), v2844)
	mBase = m.M
	v2853 = m.ExcPending
	if v2853 != 0 {
		goto L28
	} else {
		goto L441
	}
L439:
	;
	goto L440
L440:
	;
	v2996 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	if v2996 == int32(-1) {
		goto L456
	} else {
		goto L457
	}
L441:
	;
	v2854 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v2855 = F_dsa_get_address(m, v2854, v2852)
	mBase = m.M
	v2856 = m.ExcPending
	if v2856 != 0 {
		goto L28
	} else {
		goto L442
	}
L442:
	;
	v2857 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2855)+2)) = v2857
	v2859 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v2855))) = uint16(v2859)
	v2861 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2855)+34)) = uint16(v2861)
	*(*int64)(unsafe.Add(mBase, uint32(v2855)+26)) = v2857
	*(*int64)(unsafe.Add(mBase, uint32(v2855)+18)) = v2857
	*(*int64)(unsafe.Add(mBase, uint32(v2855)+10)) = v2857
	v2869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2848)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2855)+2)) = uint8(v2869)
	v2872 = v2855 + int32(4)
	v2874 = v2855 + int32(36)
	v2876 = v2855 + int32(2)
	v2878 = v2848 + int32(268)
	v2880 = v2848 + int32(12)
	v2884 = v2844
	v2888 = v2844
	goto L443
L443:
	;
	v2901 = int32(0)
	v2903 = v2901
	v2904 = v2901
	v2906 = v2884
	goto L445
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v2852
	v2972 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v2973 = *(*int32)(unsafe.Add(mBase, uint32(v2847)))
	F_dsa_free(m, v2972, v2973)
	mBase = m.M
	v2975 = m.ExcPending
	if v2975 != 0 {
		goto L28
	} else {
		goto L455
	}
L445:
	;
	v2924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2906+v2880))))
	if v2924 != int32(255) {
		goto L447
	} else {
		goto L448
	}
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2872+v2888<<(uint(int32(2))%32)))) = v2956
	v2968 = v2888 + int32(1)
	if v2968 != int32(8) {
		v2884 = v2958
		v2888 = v2968
		goto L443
	} else {
		goto L454
	}
L447:
	;
	v2927 = int32(2)
	v2933 = *(*int32)(unsafe.Add(mBase, uint32(v2878+v2924<<(uint(v2927)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2874+v2906<<(uint(v2927)%32)))) = v2933
	v2938 = int32(1)<<(uint(v2903)%32) | v2904
	goto L449
L448:
	;
	v2938 = v2904
	goto L449
L449:
	;
	v2940 = v2906 + int32(1)
	v2942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2880+v2940))))
	if v2942 != int32(255) {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v2945 = int32(2)
	v2951 = *(*int32)(unsafe.Add(mBase, uint32(v2878+v2942<<(uint(v2945)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2874+v2940<<(uint(v2945)%32)))) = v2951
	v2956 = v2945<<(uint(v2903)%32) | v2938
	goto L452
L451:
	;
	v2956 = v2938
	goto L452
L452:
	;
	v2957 = int32(2)
	v2958 = v2906 + v2957
	v2960 = v2903 + v2957
	if v2960 != int32(32) {
		v2903 = v2960
		v2904 = v2956
		v2906 = v2958
		goto L445
	} else {
		goto L453
	}
L453:
	;
	goto L446
L454:
	;
	goto L444
L455:
	;
	v2980 = v2872 + int32(base.Ui32(v2843)>>(uint(int32(3))%32))&int32(28)
	v2981 = *(*int32)(unsafe.Add(mBase, uint32(v2980)))
	v2982 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2980))) = v2981 | v2982<<(uint(v2843)%32)
	v2986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2876))))
	v2988 = v2986 + v2982
	*(*uint8)(unsafe.Add(mBase, uint32(v2876))) = uint8(v2988)
	v3049 = v2874 + v2843<<(uint(int32(2))%32)
	goto L3
L456:
	;
	v2999 = *(*int32)(unsafe.Add(mBase, uint32(v335)+8))
	v3002 = v2999
	v3003 = base.B2i32(v2999 != int32(-1))
	goto L458
L457:
	;
	v3002 = v2996
	v3003 = int32(0)
	goto L458
L458:
	;
	v3004 = int32(2)
	v3007 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v335+int32(4)+v3003<<(uint(v3004)%32)))) = v3002 + v3007 | v3002
	v3019 = base.I32_ctz(v3002^int32(-1)) + v3003<<(uint(int32(5))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v335+v332&int32(255))+12)) = uint8(v3019)
	v3022 = v2837 + v3007
	*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)) = uint8(v3022)
	v3049 = v335 + v3019<<(uint(v3004)%32) + int32(268)
	goto L3
L459:
	;
	v3052 = F_shared_ts_extend_down(m, v167, v3049, v166, v321)
	mBase = m.M
	v3053 = m.ExcPending
	if v3053 != 0 {
		goto L28
	} else {
		goto L460
	}
L460:
	;
	v3058 = v3052
	goto L2
L461:
	;
	v3104 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v3105 = *(*int64)(unsafe.Add(mBase, uint32(v3104)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v3104)+40)) = v3105 + int64(1)
	goto L1
L462:
	;
	v3081 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v3083 = F_dsa_allocate_extended(m, v3081, v165, int32(0))
	mBase = m.M
	v3084 = m.ExcPending
	if v3084 != 0 {
		goto L28
	} else {
		goto L465
	}
L463:
	;
	goto L464
L464:
	;
	if v165 != 0 {
		goto L472
	} else {
		goto L473
	}
L465:
	;
	v3085 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v3086 = F_dsa_get_address(m, v3085, v3083)
	mBase = m.M
	v3087 = m.ExcPending
	if v3087 != 0 {
		goto L28
	} else {
		goto L466
	}
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3058))) = v3083
	if v165 != 0 {
		goto L468
	} else {
		goto L469
	}
L467:
	;
	goto L461
L468:
	;
	v3091 = F__emscripten_memcpy_bulkmem(m, v3086, v23-int32(-64), v165)
	mBase = m.M
	goto L470
L469:
	;
	goto L470
L470:
	;
	goto L467
L471:
	;
	v3097 = *(*int32)(unsafe.Add(mBase, uint32(v3096)))
	*(*int32)(unsafe.Add(mBase, uint32(v3096))) = v3097 | int32(1)
	goto L461
L472:
	;
	v3095 = F__emscripten_memcpy_bulkmem(m, v3058, v23-int32(-64), v165)
	mBase = m.M
	v3096 = v3095
	goto L474
L473:
	;
	v3096 = v3058
	goto L474
L474:
	;
	goto L471
}
