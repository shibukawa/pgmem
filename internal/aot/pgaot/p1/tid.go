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
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
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
				v81 = v2
				return v81
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
							v63 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v63)
							v65 = v61
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+188))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
							v70 = m.T0[v69].(func(*base.Module, int32, int32, int32) int32)(m, v65, v10, v7)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								if v70 == int32(0) {
									v74 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v74)
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
									m.T0[v77].(func(*base.Module, int32))(m, v7)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										v81 = v7
										return v81
									}
								} else {
									v81 = v7
									return v81
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
							v63 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v63)
							v65 = v61
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+188))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
							v70 = m.T0[v69].(func(*base.Module, int32, int32, int32) int32)(m, v65, v10, v7)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								if v70 == int32(0) {
									v74 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v74)
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
									m.T0[v77].(func(*base.Module, int32))(m, v7)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										v81 = v7
										return v81
									}
								} else {
									v81 = v7
									return v81
								}
							}
						}
					}
				}
			}
		}
	} else {
		v65 = v8
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
		v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+188))
		v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
		v70 = m.T0[v69].(func(*base.Module, int32, int32, int32) int32)(m, v65, v10, v7)
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			if v70 == int32(0) {
				v74 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v74)
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
				m.T0[v77].(func(*base.Module, int32))(m, v7)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					v81 = v7
					return v81
				}
			} else {
				v81 = v7
				return v81
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
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
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
	var v148 int32
	_ = v148
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
	var v218 int32
	_ = v218
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
	var v292 int32
	_ = v292
	var v309 int32
	_ = v309
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
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
	var v350 int32
	_ = v350
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v390 int32
	_ = v390
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int64
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
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
	var v502 int32
	_ = v502
	var v504 int64
	_ = v504
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v741 int64
	_ = v741
	var v743 int64
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int64
	_ = v763
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v839 int32
	_ = v839
	var v840 int64
	_ = v840
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v871 int32
	_ = v871
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v909 int32
	_ = v909
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v948 int32
	_ = v948
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1028 int32
	_ = v1028
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int64
	_ = v1056
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1086 int32
	_ = v1086
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1128 int32
	_ = v1128
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1250 int32
	_ = v1250
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1318 int32
	_ = v1318
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1346 int32
	_ = v1346
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1382 int64
	_ = v1382
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1404 int32
	_ = v1404
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1450 int32
	_ = v1450
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1550 int32
	_ = v1550
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1581 int32
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1622 int32
	_ = v1622
	var v1645 int32
	_ = v1645
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1680 int64
	_ = v1680
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1728 int32
	_ = v1728
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1739 int32
	_ = v1739
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1752 int32
	_ = v1752
	var v1757 int32
	_ = v1757
	var v1761 int32
	_ = v1761
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1773 int32
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1777 int32
	_ = v1777
	var v1783 int32
	_ = v1783
	var v1786 int32
	_ = v1786
	var v1793 int32
	_ = v1793
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1852 int32
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1876 int32
	_ = v1876
	var v1884 int32
	_ = v1884
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1897 int32
	_ = v1897
	var v1916 int32
	_ = v1916
	var v1919 int32
	_ = v1919
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1928 int32
	_ = v1928
	var v1930 int32
	_ = v1930
	var v1932 int32
	_ = v1932
	var v1938 int32
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1948 int32
	_ = v1948
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1970 int32
	_ = v1970
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1979 int32
	_ = v1979
	var v1981 int32
	_ = v1981
	var v1986 int32
	_ = v1986
	var v2009 int32
	_ = v2009
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2033 int32
	_ = v2033
	var v2060 int32
	_ = v2060
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2069 int32
	_ = v2069
	var v2085 int32
	_ = v2085
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2104 int32
	_ = v2104
	var v2109 int32
	_ = v2109
	var v2110 int64
	_ = v2110
	var v2118 int32
	_ = v2118
	var v2136 int32
	_ = v2136
	var v2139 int32
	_ = v2139
	var v2145 int32
	_ = v2145
	var v2164 int32
	_ = v2164
	var v2167 int32
	_ = v2167
	var v2173 int32
	_ = v2173
	var v2189 int32
	_ = v2189
	var v2191 int64
	_ = v2191
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2206 int32
	_ = v2206
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
	var v2213 int64
	_ = v2213
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2226 int32
	_ = v2226
	var v2230 int32
	_ = v2230
	var v2234 int32
	_ = v2234
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2280 int32
	_ = v2280
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2305 int32
	_ = v2305
	var v2307 int32
	_ = v2307
	var v2310 int32
	_ = v2310
	var v2329 int32
	_ = v2329
	var v2332 int32
	_ = v2332
	var v2334 int32
	_ = v2334
	var v2340 int32
	_ = v2340
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2350 int32
	_ = v2350
	var v2352 int32
	_ = v2352
	var v2358 int32
	_ = v2358
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2374 int32
	_ = v2374
	var v2378 int32
	_ = v2378
	var v2380 int32
	_ = v2380
	var v2383 int32
	_ = v2383
	var v2386 int32
	_ = v2386
	var v2405 int32
	_ = v2405
	var v2409 int32
	_ = v2409
	var v2411 int32
	_ = v2411
	var v2415 int32
	_ = v2415
	var v2417 int32
	_ = v2417
	var v2423 int32
	_ = v2423
	var v2425 int32
	_ = v2425
	var v2433 int32
	_ = v2433
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2438 int32
	_ = v2438
	var v2452 int32
	_ = v2452
	var v2455 int32
	_ = v2455
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2461 int32
	_ = v2461
	var v2484 int32
	_ = v2484
	var v2505 int64
	_ = v2505
	var v2508 int32
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2523 int64
	_ = v2523
	var v2531 int32
	_ = v2531
	var v2533 int32
	_ = v2533
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2553 int32
	_ = v2553
	var v2555 int32
	_ = v2555
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2564 int32
	_ = v2564
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2572 int32
	_ = v2572
	var v2575 int32
	_ = v2575
	var v2580 int32
	_ = v2580
	var v2583 int32
	_ = v2583
	var v2588 int32
	_ = v2588
	var v2591 int32
	_ = v2591
	var v2596 int32
	_ = v2596
	var v2601 int32
	_ = v2601
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2607 int32
	_ = v2607
	var v2613 int32
	_ = v2613
	var v2614 int32
	_ = v2614
	var v2619 int32
	_ = v2619
	var v2621 int32
	_ = v2621
	var v2623 int32
	_ = v2623
	var v2629 int32
	_ = v2629
	var v2632 int32
	_ = v2632
	var v2639 int32
	_ = v2639
	var v2657 int32
	_ = v2657
	var v2658 int32
	_ = v2658
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2662 int32
	_ = v2662
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2670 int32
	_ = v2670
	var v2672 int32
	_ = v2672
	var v2675 int32
	_ = v2675
	var v2698 int32
	_ = v2698
	var v2701 int32
	_ = v2701
	var v2703 int32
	_ = v2703
	var v2722 int32
	_ = v2722
	var v2730 int32
	_ = v2730
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2738 int32
	_ = v2738
	var v2743 int32
	_ = v2743
	var v2762 int32
	_ = v2762
	var v2765 int32
	_ = v2765
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2774 int32
	_ = v2774
	var v2776 int32
	_ = v2776
	var v2778 int32
	_ = v2778
	var v2784 int32
	_ = v2784
	var v2787 int32
	_ = v2787
	var v2794 int32
	_ = v2794
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2816 int32
	_ = v2816
	var v2821 int32
	_ = v2821
	var v2822 int32
	_ = v2822
	var v2825 int32
	_ = v2825
	var v2827 int32
	_ = v2827
	var v2832 int32
	_ = v2832
	var v2855 int32
	_ = v2855
	var v2858 int32
	_ = v2858
	var v2860 int32
	_ = v2860
	var v2879 int32
	_ = v2879
	var v2906 int32
	_ = v2906
	var v2909 int32
	_ = v2909
	var v2910 int32
	_ = v2910
	var v2915 int32
	_ = v2915
	var v2931 int32
	_ = v2931
	var v2938 int32
	_ = v2938
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2944 int32
	_ = v2944
	var v2954 int32
	_ = v2954
	var v2959 int32
	_ = v2959
	var v2960 int64
	_ = v2960
	v5 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(336)
	m.G0 = v23
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v5
	if base.Ui32(int32(2)) <= base.Ui32(l3) {
		goto L20
	} else {
		goto L21
	}
L1:
	;
	m.G0 = v23 + int32(336)
	return
L2:
	;
	v2931 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+65)))
	if base.Ui32(int32(5)) <= base.Ui32(v2931<<(uint(int32(2))%32)+int32(4)) {
		goto L404
	} else {
		goto L405
	}
L3:
	;
	if v320 == int32(0) {
		v2915 = v2906
		goto L2
	} else {
		goto L401
	}
L4:
	;
	v2730 = v335 + int32(3)
	if v2173 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L5:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v335+v2701)+3)) = uint8(v332)
	v2722 = v2703 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)) = uint8(v2722)
	v2906 = v335 + v2701<<(uint(int32(2))%32) + int32(8)
	goto L3
L6:
	;
	if base.Ui32(v2118) <= base.Ui32(v2145) {
		v2701 = v2145
		v2703 = v2118
		goto L5
	} else {
		goto L369
	}
L7:
	;
	v2505 = *(*int64)(unsafe.Add(mBase, uint32(v23)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v2505
	v2508 = v332 & int32(255)
	v2509 = int32(0)
	v2511 = v23 + int32(24)
	v2512 = *(*int32)(unsafe.Add(mBase, uint32(v2511)+4))
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v2516 = F_dsa_allocate_extended(m, v2513, int32(96), v2509)
	mBase = m.M
	v2517 = m.ExcPending
	if v2517 != 0 {
		goto L38
	} else {
		goto L351
	}
L8:
	;
	v2189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	if v2173 != v2189 {
		goto L4
	} else {
		goto L323
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+332)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v334
	if v337 != 0 {
		v2173 = v2118
		goto L8
	} else {
		goto L317
	}
L10:
	;
	v2085 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+65)))
	if base.Ui32(int32(5)) <= base.Ui32(v2085<<(uint(int32(2))%32)+int32(4)) {
		goto L309
	} else {
		goto L310
	}
L11:
	;
	if v883 == int32(0) {
		v2069 = v2060
		goto L10
	} else {
		goto L306
	}
L12:
	;
	v1884 = v895 + int32(3)
	if v1346 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v895+v1855)+3)) = uint8(v894)
	v1876 = v1857 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v895)+2)) = uint8(v1876)
	v2060 = v895 + v1855<<(uint(int32(2))%32) + int32(8)
	goto L11
L14:
	;
	if base.Ui32(v1292) <= base.Ui32(v1318) {
		v1855 = v1318
		v1857 = v1292
		goto L13
	} else {
		goto L274
	}
L15:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v23)+328))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v1666
	v1669 = v894 & int32(255)
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(12))))
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v1676 = F_MemoryContextAlloc(m, v1674, int32(100))
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L38
	} else {
		goto L257
	}
L16:
	;
	v1362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+1)))
	if v1346 != v1362 {
		goto L12
	} else {
		goto L233
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v895
	if v896 != 0 {
		v1346 = v1292
		goto L16
	} else {
		goto L227
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L38
	} else {
		goto L224
	}
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+65)) = uint8(v148)
	v165 = base.I32_extend8_s(v148)<<(uint(int32(2))%32) + int32(4)
	v166 = base.I64_extend_i32_u(l1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v169 != 0 {
		goto L44
	} else {
		goto L45
	}
L20:
	;
	v29 = int32(1)
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+l3<<(uint(v29)%32)-int32(2)))))
	v38 = int32(base.Ui32(v34)>>(uint(int32(5))%32)) + v29
	v46 = v5
	v49 = int32(32)
	v53 = v5
	goto L23
L21:
	;
	goto L22
L22:
	;
	if l3 != int32(1) {
		v148 = v5
		goto L19
	} else {
		goto L42
	}
L23:
	;
	if l3 < v46 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L38
	} else {
		goto L39
	}
L25:
	;
	v63 = v46
	goto L27
L26:
	;
	v63 = l3
	goto L27
L27:
	;
	v70 = v46
	v75 = int32(0)
	goto L29
L28:
	;
	goto L24
L29:
	;
	if v63 == v70 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(68)+v53<<(uint(int32(2))%32)))) = v75
	v113 = v53 + int32(1)
	if v113 != v38 {
		v46 = v102
		v49 = v49 + int32(32)
		v53 = v113
		goto L23
	} else {
		goto L37
	}
L31:
	;
	goto L30
L32:
	;
	v102 = v63
	goto L31
L33:
	;
	goto L34
L34:
	;
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v70<<(uint(int32(1))%32)))))
	if base.Ui32((v89-int32(2049))&int32(_a_F_TidStoreSetBlockOffsets_0)) <= base.Ui32(int32(_a_F_TidStoreSetBlockOffsets_1)) {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	v96 = int32(1)
	if base.Ui32(v89) < base.Ui32(v49) {
		v70 = v70 + v96
		v75 = v96<<(uint(v89)%32) | v75
		goto L29
	} else {
		goto L36
	}
L36:
	;
	v102 = v70
	goto L31
L37:
	;
	v148 = v38
	goto L19
L38:
	;
	return
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v89
	F_errmsg_internal(m, int32(_a_F_TidStoreSetBlockOffsets_2), v23+int32(48))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_TidStoreSetBlockOffsets_3), int32(396), int32(_a_F_TidStoreSetBlockOffsets_4))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	if base.Ui32((v132-int32(2049))&int32(_a_F_TidStoreSetBlockOffsets_0)) <= base.Ui32(int32(_a_F_TidStoreSetBlockOffsets_1)) {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+66)) = uint16(v132)
	v148 = v5
	goto L19
L44:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v168)+32))
	if base.Ui64(v170) < base.Ui64(v166) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v741 = *(*int64)(unsafe.Add(mBase, uint32(v168)+8))
	if base.Ui64(v741) < base.Ui64(v166) {
		goto L139
	} else {
		goto L140
	}
L47:
	;
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v168)+40))
	if v172 == int64(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v292 = v168
	goto L49
L49:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v292)+48))
	v320 = v309
	v321 = v292 + int32(24)
	goto L69
L50:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v168)+24))
	v177 = F_dsa_get_address(m, v175, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L38
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+48))
	if v166 == int64(0) {
		goto L55
	} else {
		goto L56
	}
L53:
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
		goto L38
	} else {
		goto L54
	}
L54:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+48)) = v185
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v195 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v194)+32)) = v195<<(uint(base.I64_extend_i32_u(v185+int32(8)))%64) ^ v195
	v2915 = v190
	goto L2
L55:
	;
	v214 = int32(0)
	goto L57
L56:
	;
	v214 = (base.I32_wrap_i64(base.I64_clz(v166)) ^ int32(-1)) & int32(56)
	goto L57
L57:
	;
	if v204 < v214 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v218 = v204
	goto L61
L59:
	;
	v275 = v203
	goto L60
L60:
	;
	v276 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v275)+32)) = v276<<(uint(base.I64_extend_i32_u(v214+int32(8)))%64) ^ v276
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v284)+48)) = v214
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v292 = v286
	goto L49
L61:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v239 = F_dsa_allocate_extended(m, v236, int32(24), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L38
	} else {
		goto L63
	}
L62:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v275 = v254
	goto L60
L63:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v242 = F_dsa_get_address(m, v241, v239)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L38
	} else {
		goto L64
	}
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v242))) = int64(66560)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = v247
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v249)+24)) = v239
	v252 = v218 + int32(8)
	if v252 < v214 {
		v218 = v252
		goto L61
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	v686 = v684 & int32(1)
	if base.Ui32(v165) <= base.Ui32(int32(4)) {
		goto L115
	} else {
		goto L116
	}
L67:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+4)) = v439 | v441
	v677 = v673 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)) = uint8(v677)
	v2906 = v335 + v433<<(uint(int32(2))%32) + int32(36)
	goto L3
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v23)+332)) = v335
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	if v482 == v483 {
		goto L94
	} else {
		goto L95
	}
L69:
	;
	v332 = base.I32_wrap_i64(int64(base.Ui64(v166) >> (uint(base.I64_extend_i32_u(v320)) % 64)))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v335 = F_dsa_get_address(m, v333, v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L38
	} else {
		goto L77
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v23)+332)) = v335
	v477 = int32(0)
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	if v479 != 0 {
		v2701 = v477
		v2703 = v477
		goto L5
	} else {
		goto L93
	}
L71:
	;
	goto L70
L72:
	;
	if v320 == int32(0) {
		goto L66
	} else {
		goto L92
	}
L73:
	;
	v433 = v332 & int32(255)
	v438 = v335 + int32(base.Ui32(v332)>>(uint(int32(3))%32))&int32(28)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	v441 = int32(1) << (uint(v332) % 32)
	if v439&v441 == int32(0) {
		goto L67
	} else {
		goto L91
	}
L74:
	;
	v421 = int32(255)
	v423 = v335 + v332&v421
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+12)))
	if v424 == v421 {
		goto L68
	} else {
		goto L90
	}
L75:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	if v379 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	if v340 == int32(0) {
		goto L71
	} else {
		goto L78
	}
L77:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	switch v337 - int32(1) {
	case 0:
		goto L75
	case 1:
		goto L74
	case 2:
		goto L73
	default:
		goto L76
	}
L78:
	;
	v350 = int32(0)
	goto L79
L79:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350+(v335+int32(3))))))
	if v332&int32(255) == v374 {
		v470 = v335 + v350<<(uint(int32(2))%32) + int32(8)
		goto L72
	} else {
		goto L81
	}
L80:
	;
	v2118 = v340
	goto L9
L81:
	;
	v377 = v350 + int32(1)
	if v377 != v340 {
		v350 = v377
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+332)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v334
	v2173 = int32(0)
	goto L8
L84:
	;
	goto L85
L85:
	;
	v390 = int32(0)
	goto L86
L86:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390+(v335+int32(3))))))
	if v414 == v332&int32(255) {
		v470 = v335 + v390<<(uint(int32(2))%32) + int32(36)
		goto L72
	} else {
		goto L88
	}
L87:
	;
	v2118 = v379
	goto L9
L88:
	;
	v419 = v390 + int32(1)
	if v379 != v419 {
		v390 = v419
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v470 = v335 + v424<<(uint(int32(2))%32) + int32(268)
	goto L72
L91:
	;
	v470 = v335 + v433<<(uint(int32(2))%32) + int32(36)
	goto L72
L92:
	;
	v320 = v320 - int32(8)
	v321 = v470
	goto L69
L93:
	;
	goto L7
L94:
	;
	v485 = *(*int64)(unsafe.Add(mBase, uint32(v23)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v485
	v488 = v332 & int32(255)
	v489 = int32(0)
	v492 = v23 + int32(40)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v497 = F_dsa_allocate_extended(m, v494, int32(1060), v489)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L38
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	if v639 == int32(-1) {
		goto L112
	} else {
		goto L113
	}
L97:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v500 = F_dsa_get_address(m, v499, v497)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L38
	} else {
		goto L98
	}
L98:
	;
	v502 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v500)+34)) = uint16(v502)
	v504 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v500)+26)) = v504
	*(*int64)(unsafe.Add(mBase, uint32(v500)+18)) = v504
	*(*int64)(unsafe.Add(mBase, uint32(v500)+10)) = v504
	*(*int64)(unsafe.Add(mBase, uint32(v500)+2)) = v504
	v512 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v500))) = uint16(v512)
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v500)+2)) = uint8(v514)
	v517 = v500 + int32(4)
	v519 = v500 + int32(36)
	v521 = v500 + int32(2)
	v523 = v493 + int32(268)
	v525 = v493 + int32(12)
	v530 = v489
	v534 = v489
	goto L99
L99:
	;
	v546 = int32(0)
	v548 = v546
	v549 = v546
	v552 = v530
	goto L101
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v497
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	F_dsa_free(m, v617, v618)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L38
	} else {
		goto L111
	}
L101:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552+v525))))
	if v569 != int32(255) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517+v534<<(uint(int32(2))%32)))) = v601
	v613 = v534 + int32(1)
	if v613 != int32(8) {
		v530 = v603
		v534 = v613
		goto L99
	} else {
		goto L110
	}
L103:
	;
	v572 = int32(2)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v523+v569<<(uint(v572)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v519+v552<<(uint(v572)%32)))) = v578
	v583 = int32(1)<<(uint(v548)%32) | v549
	goto L105
L104:
	;
	v583 = v549
	goto L105
L105:
	;
	v585 = v552 + int32(1)
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525+v585))))
	if v587 != int32(255) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v590 = int32(2)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v523+v587<<(uint(v590)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v519+v585<<(uint(v590)%32)))) = v596
	v601 = v590<<(uint(v548)%32) | v583
	goto L108
L107:
	;
	v601 = v583
	goto L108
L108:
	;
	v602 = int32(2)
	v603 = v552 + v602
	v605 = v548 + v602
	if v605 != int32(32) {
		v548 = v605
		v549 = v601
		v552 = v603
		goto L101
	} else {
		goto L109
	}
L109:
	;
	goto L102
L110:
	;
	goto L100
L111:
	;
	v625 = v517 + int32(base.Ui32(v488)>>(uint(int32(3))%32))&int32(28)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v625)))
	v627 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v625))) = v626 | v627<<(uint(v488)%32)
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521))))
	v633 = v631 + v627
	*(*uint8)(unsafe.Add(mBase, uint32(v521))) = uint8(v633)
	v2906 = v519 + v488<<(uint(int32(2))%32)
	goto L3
L112:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v335)+8))
	v645 = base.B2i32(v642 != int32(-1))
	v646 = v642
	goto L114
L113:
	;
	v645 = int32(0)
	v646 = v639
	goto L114
L114:
	;
	v649 = int32(2)
	v652 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v335+int32(4)+v645<<(uint(v649)%32)))) = v646 + v652 | v646
	v663 = base.I32_ctz(v646^int32(-1)) + v645<<(uint(int32(5))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v423+int32(12)))) = uint8(v663)
	v666 = v482 + v652
	*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)) = uint8(v666)
	v2906 = v335 + v663<<(uint(v649)%32) + int32(268)
	goto L3
L115:
	;
	if v686 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L117
L117:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v686 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L118:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	F_dsa_free(m, v691, v684)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L38
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	if v165 != 0 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L120
L122:
	;
	base.MemoryCopy(m, v470, v23-int32(-64), v165)
	goto L124
L123:
	;
	goto L124
L124:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = v697 | int32(1)
	goto L1
L125:
	;
	v704 = F_dsa_get_address(m, v701, v684)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L38
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v730 = F_dsa_allocate_extended(m, v701, v165, int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L38
	} else {
		goto L136
	}
L128:
	;
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704)+1)))
	if v706 != v148&int32(255) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	F_dsa_free(m, v710, v711)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L38
	} else {
		goto L132
	}
L130:
	;
	v723 = v704
	goto L131
L131:
	;
	if v165 == int32(0) {
		goto L1
	} else {
		goto L135
	}
L132:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v716 = F_dsa_allocate_extended(m, v714, v165, int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L38
	} else {
		goto L133
	}
L133:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v719 = F_dsa_get_address(m, v718, v716)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L38
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = v716
	v723 = v719
	goto L131
L135:
	;
	base.MemoryCopy(m, v723, v23-int32(-64), v165)
	goto L1
L136:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v733 = F_dsa_get_address(m, v732, v730)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L38
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = v730
	if v165 == int32(0) {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	base.MemoryCopy(m, v733, v23-int32(-64), v165)
	goto L1
L139:
	;
	v743 = *(*int64)(unsafe.Add(mBase, uint32(v168)+16))
	if v743 == int64(0) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v856 = v168
	goto L141
L141:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v856)+24))
	v877 = v856
	v883 = v871
	goto L159
L142:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v747 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v746)+2)) = uint8(v747)
	v753 = (base.I32_clz(l1) ^ int32(-1)) & int32(24)
	v754 = int32(base.Ui32(l1) >> (uint(v753) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v746)+3)) = uint8(v754)
	v758 = F_local_ts_extend_down(m, v167, v746+int32(8), v166, v753)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L38
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v771)+24))
	if v166 == int64(0) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v760)+24)) = v753
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v763 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v762)+8)) = v763<<(uint(base.I64_extend_i32_u(v753+int32(8)))%64) ^ v763
	v2069 = v758
	goto L10
L146:
	;
	v782 = int32(0)
	goto L148
L147:
	;
	v782 = (base.I32_wrap_i64(base.I64_clz(v166)) ^ int32(-1)) & int32(56)
	goto L148
L148:
	;
	if v772 < v782 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v788 = v772
	goto L152
L150:
	;
	v839 = v771
	goto L151
L151:
	;
	v840 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v839)+8)) = v840<<(uint(base.I64_extend_i32_u(v782+int32(8)))%64) ^ v840
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v848)+24)) = v782
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v856 = v850
	goto L141
L152:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v806 = F_MemoryContextAlloc(m, v804, int32(24))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L38
	} else {
		goto L154
	}
L153:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v839 = v818
	goto L151
L154:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v806))) = int64(66560)
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v810)))
	*(*int32)(unsafe.Add(mBase, uint32(v806)+8)) = v811
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v813))) = v806
	v816 = v788 + int32(8)
	if v816 < v782 {
		v788 = v816
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1028)))
	v1234 = v1232 & int32(1)
	if base.Ui32(v165) <= base.Ui32(int32(4)) {
		goto L203
	} else {
		goto L204
	}
L157:
	;
	v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v996)+4)) = v997 | v999
	v1225 = v1221 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v895)+2)) = uint8(v1225)
	v2060 = v895 + v991<<(uint(int32(2))%32) + int32(36)
	goto L11
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v895
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+2)))
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+1)))
	if v1038 == v1039 {
		goto L183
	} else {
		goto L184
	}
L159:
	;
	v894 = base.I32_wrap_i64(int64(base.Ui64(v166) >> (uint(base.I64_extend_i32_u(v883)) % 64)))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v877)))
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895))))
	switch v896 - int32(1) {
	case 0:
		goto L165
	case 1:
		goto L164
	case 2:
		goto L163
	default:
		goto L166
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v895
	v1034 = int32(0)
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+1)))
	if v1036 != 0 {
		v1855 = v1034
		v1857 = v1034
		goto L13
	} else {
		goto L182
	}
L161:
	;
	goto L160
L162:
	;
	if v883 == int32(0) {
		goto L156
	} else {
		goto L181
	}
L163:
	;
	v991 = v894 & int32(255)
	v996 = v895 + int32(base.Ui32(v894)>>(uint(int32(3))%32))&int32(28)
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v996)+4))
	v999 = int32(1) << (uint(v894) % 32)
	if v997&v999 == int32(0) {
		goto L157
	} else {
		goto L180
	}
L164:
	;
	v979 = int32(255)
	v981 = v895 + v894&v979
	v982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v981)+12)))
	if v982 == v979 {
		goto L158
	} else {
		goto L179
	}
L165:
	;
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+2)))
	if v938 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L166:
	;
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+2)))
	if v899 == int32(0) {
		goto L161
	} else {
		goto L167
	}
L167:
	;
	v909 = int32(0)
	goto L168
L168:
	;
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909+(v895+int32(3))))))
	if v894&int32(255) == v933 {
		v1028 = v895 + v909<<(uint(int32(2))%32) + int32(8)
		goto L162
	} else {
		goto L170
	}
L169:
	;
	v1292 = v899
	goto L17
L170:
	;
	v936 = v909 + int32(1)
	if v936 != v899 {
		v909 = v936
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+328)) = v895
	v1346 = int32(0)
	goto L16
L173:
	;
	goto L174
L174:
	;
	v948 = int32(0)
	goto L175
L175:
	;
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948+(v895+int32(3))))))
	if v972 == v894&int32(255) {
		v1028 = v895 + v948<<(uint(int32(2))%32) + int32(36)
		goto L162
	} else {
		goto L177
	}
L176:
	;
	v1292 = v938
	goto L17
L177:
	;
	v977 = v948 + int32(1)
	if v938 != v977 {
		v948 = v977
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	v1028 = v895 + v982<<(uint(int32(2))%32) + int32(268)
	goto L162
L180:
	;
	v1028 = v895 + v991<<(uint(int32(2))%32) + int32(36)
	goto L162
L181:
	;
	v877 = v1028
	v883 = v883 - int32(8)
	goto L159
L182:
	;
	goto L15
L183:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v23)+328))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v1041
	v1044 = v894 & int32(255)
	v1045 = int32(0)
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(20))))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	v1052 = F_MemoryContextAlloc(m, v1050, int32(1060))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L38
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v895)+4))
	if v1187 == int32(-1) {
		goto L200
	} else {
		goto L201
	}
L186:
	;
	v1054 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1052)+34)) = uint16(v1054)
	v1056 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1052)+26)) = v1056
	*(*int64)(unsafe.Add(mBase, uint32(v1052)+18)) = v1056
	*(*int64)(unsafe.Add(mBase, uint32(v1052)+10)) = v1056
	*(*int64)(unsafe.Add(mBase, uint32(v1052)+2)) = v1056
	v1064 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v1052))) = uint16(v1064)
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1048)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1052)+2)) = uint8(v1066)
	v1069 = v1048 + int32(268)
	v1071 = v1048 + int32(12)
	v1073 = v1052 + int32(4)
	v1075 = v1052 + int32(36)
	v1080 = v1045
	v1086 = v1045
	goto L187
L187:
	;
	v1096 = int32(0)
	v1098 = v1096
	v1100 = v1096
	v1102 = v1080
	goto L189
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v877))) = v1052
	F_pfree(m, v1048)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L38
	} else {
		goto L199
	}
L189:
	;
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1102+v1071))))
	if v1119 != int32(255) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1073+v1086<<(uint(int32(2))%32)))) = v1151
	v1163 = v1086 + int32(1)
	if v1163 != int32(8) {
		v1080 = v1153
		v1086 = v1163
		goto L187
	} else {
		goto L198
	}
L191:
	;
	v1122 = int32(2)
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1069+v1119<<(uint(v1122)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1075+v1102<<(uint(v1122)%32)))) = v1128
	v1133 = int32(1)<<(uint(v1098)%32) | v1100
	goto L193
L192:
	;
	v1133 = v1100
	goto L193
L193:
	;
	v1135 = v1102 + int32(1)
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1071+v1135))))
	if v1137 != int32(255) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1140 = int32(2)
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1069+v1137<<(uint(v1140)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1075+v1135<<(uint(v1140)%32)))) = v1146
	v1151 = v1140<<(uint(v1098)%32) | v1133
	goto L196
L195:
	;
	v1151 = v1133
	goto L196
L196:
	;
	v1152 = int32(2)
	v1153 = v1102 + v1152
	v1155 = v1098 + v1152
	if v1155 != int32(32) {
		v1098 = v1155
		v1100 = v1151
		v1102 = v1153
		goto L189
	} else {
		goto L197
	}
L197:
	;
	goto L190
L198:
	;
	goto L188
L199:
	;
	v1173 = v1073 + int32(base.Ui32(v1044)>>(uint(int32(3))%32))&int32(28)
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1173)))
	v1175 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1173))) = v1174 | v1175<<(uint(v1044)%32)
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+2)))
	v1181 = v1179 + v1175
	*(*uint8)(unsafe.Add(mBase, uint32(v1052)+2)) = uint8(v1181)
	v2060 = v1075 + v1044<<(uint(int32(2))%32)
	goto L11
L200:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v895)+8))
	v1193 = base.B2i32(v1190 != int32(-1))
	v1194 = v1190
	goto L202
L201:
	;
	v1193 = int32(0)
	v1194 = v1187
	goto L202
L202:
	;
	v1197 = int32(2)
	v1200 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v895+int32(4)+v1193<<(uint(v1197)%32)))) = v1194 + v1200 | v1194
	v1211 = base.I32_ctz(v1194^int32(-1)) + v1193<<(uint(int32(5))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v981+int32(12)))) = uint8(v1211)
	v1214 = v1038 + v1200
	*(*uint8)(unsafe.Add(mBase, uint32(v895)+2)) = uint8(v1214)
	v2060 = v895 + v1211<<(uint(v1197)%32) + int32(268)
	goto L11
L203:
	;
	if v1234 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	if v1234 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L206:
	;
	F_pfree(m, v1232)
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L38
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	if v165 != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	goto L208
L210:
	;
	base.MemoryCopy(m, v1028, v23-int32(-64), v165)
	goto L212
L211:
	;
	goto L212
L212:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1028)))
	*(*int32)(unsafe.Add(mBase, uint32(v1028))) = v1244 | int32(1)
	goto L1
L213:
	;
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1232)+1)))
	if v1250 != v148&int32(255) {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	goto L215
L215:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v1267 = F_MemoryContextAlloc(m, v1266, v165)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L38
	} else {
		goto L222
	}
L216:
	;
	F_pfree(m, v1232)
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L38
	} else {
		goto L219
	}
L217:
	;
	v1260 = v1232
	goto L218
L218:
	;
	if v165 == int32(0) {
		goto L1
	} else {
		goto L221
	}
L219:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v1257 = F_MemoryContextAlloc(m, v1256, v165)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L38
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1028))) = v1257
	v1260 = v1257
	goto L218
L221:
	;
	base.MemoryCopy(m, v1260, v23-int32(-64), v165)
	goto L1
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1028))) = v1267
	if v165 == int32(0) {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	base.MemoryCopy(m, v1267, v23-int32(-64), v165)
	goto L1
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v132
	F_errmsg_internal(m, int32(_a_F_TidStoreSetBlockOffsets_2), v23)
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L38
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(_a_F_TidStoreSetBlockOffsets_3), int32(375), int32(_a_F_TidStoreSetBlockOffsets_4))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L38
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
	v1309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+1)))
	if v1292 == v1309 {
		goto L15
	} else {
		goto L228
	}
L228:
	;
	v1312 = v895 + int32(3)
	v1318 = int32(0)
	goto L229
L229:
	;
	v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1318+v1312))))
	if base.Ui32(v894&int32(255)) <= base.Ui32(v1337) {
		goto L14
	} else {
		goto L231
	}
L230:
	;
	v1855 = v1292
	v1857 = v1292
	goto L13
L231:
	;
	v1340 = v1318 + int32(1)
	if v1340 != v1292 {
		v1318 = v1340
		goto L229
	} else {
		goto L232
	}
L232:
	;
	goto L230
L233:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v23)+328))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v1364
	v1367 = v894 & int32(255)
	v1368 = int32(0)
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(16))))
	v1372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1371)+1)))
	if base.Ui32(v1372) <= base.Ui32(int32(31)) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v2060 = v1645
	goto L11
L235:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v1378 = F_MemoryContextAlloc(m, v1376, int32(164))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L38
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v1543 = F_MemoryContextAlloc(m, v1541, int32(524))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L38
	} else {
		goto L252
	}
L238:
	;
	v1380 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1378)+34)) = uint16(v1380)
	v1382 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1378)+26)) = v1382
	*(*int64)(unsafe.Add(mBase, uint32(v1378)+18)) = v1382
	*(*int64)(unsafe.Add(mBase, uint32(v1378)+10)) = v1382
	*(*int64)(unsafe.Add(mBase, uint32(v1378)+2)) = v1382
	v1390 = int32(_a_F_TidStoreSetBlockOffsets_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v1378))) = uint16(v1390)
	v1392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1371)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1378)+2)) = uint8(v1392)
	v1395 = v1371 + int32(3)
	v1399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1392+v1395-int32(1)))))
	if base.Ui32(v1367) <= base.Ui32(v1399) {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	v1469 = int32(36)
	v1470 = v1371 + v1469
	v1472 = v1378 + v1469
	v1474 = v1378 + int32(3)
	v1477 = int32(0)
	goto L248
L240:
	;
	if v1392 == int32(0) {
		v1450 = v1368
		goto L239
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v1450 = v1392
	goto L239
L243:
	;
	v1404 = v1368
	goto L244
L244:
	;
	v1424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1404+v1395))))
	if base.Ui32(v1367) < base.Ui32(v1424) {
		v1450 = v1404
		goto L239
	} else {
		goto L246
	}
L245:
	;
	goto L242
L246:
	;
	v1427 = v1404 + int32(1)
	if v1427 != v1392 {
		v1404 = v1427
		goto L244
	} else {
		goto L247
	}
L247:
	;
	goto L245
L248:
	;
	v1496 = v1477 | base.B2i32(base.Ui32(v1450) <= base.Ui32(v1477))
	v1499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1477+v1395))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1474+v1496))) = uint8(v1499)
	v1501 = int32(2)
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1470+v1477<<(uint(v1501)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1472+v1496<<(uint(v1501)%32)))) = v1507
	v1510 = v1477 | int32(1)
	v1512 = v1510 + base.B2i32(base.Ui32(v1450) <= base.Ui32(v1510))
	v1515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1395+v1510))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1474+v1512))) = uint8(v1515)
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1470+v1510<<(uint(v1501)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1472+v1512<<(uint(v1501)%32)))) = v1523
	v1526 = v1477 + v1501
	if v1526 != int32(16) {
		v1477 = v1526
		goto L248
	} else {
		goto L250
	}
L249:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1450+v1474))) = uint8(v1367)
	v1532 = v1392 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1378)+2)) = uint8(v1532)
	F_pfree(m, v1371)
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L38
	} else {
		goto L251
	}
L250:
	;
	goto L249
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v877))) = v1378
	v1645 = v1472 + v1450<<(uint(int32(2))%32)
	goto L234
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1543)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1543))) = int64(0)
	v1550 = v1543 + int32(12)
	base.MemoryFill(m, v1550, int32(255), int32(256))
	v1554 = int32(_a_F_TidStoreSetBlockOffsets_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v1543))) = uint16(v1554)
	v1556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1371)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1543)+2)) = uint8(v1556)
	v1559 = v1371 + int32(3)
	v1562 = int32(0)
	goto L253
L253:
	;
	v1581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1562+v1559))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1550+v1581))) = uint8(v1562)
	v1585 = v1562 | int32(1)
	v1587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1559+v1585))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1550+v1587))) = uint8(v1585)
	v1591 = v1562 | int32(2)
	v1593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1559+v1591))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1550+v1593))) = uint8(v1591)
	v1597 = v1562 | int32(3)
	v1599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1559+v1597))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1550+v1599))) = uint8(v1597)
	v1603 = v1562 + int32(4)
	if v1603 != int32(32) {
		v1562 = v1603
		goto L253
	} else {
		goto L255
	}
L254:
	;
	base.MemoryCopy(m, v1543+int32(268), v1371+int32(36), int32(128))
	*(*int64)(unsafe.Add(mBase, uint32(v1543)+4)) = int64(8589934591)
	v1615 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1550+v1367))) = uint8(v1615)
	v1618 = v1556 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1543)+2)) = uint8(v1618)
	*(*int32)(unsafe.Add(mBase, uint32(v877))) = v1543
	F_pfree(m, v1371)
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L38
	} else {
		goto L256
	}
L255:
	;
	goto L254
L256:
	;
	v1645 = v1543 + int32(396)
	goto L234
L257:
	;
	v1678 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1676)+34)) = uint16(v1678)
	v1680 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1676)+26)) = v1680
	*(*int64)(unsafe.Add(mBase, uint32(v1676)+18)) = v1680
	*(*int64)(unsafe.Add(mBase, uint32(v1676)+10)) = v1680
	*(*int64)(unsafe.Add(mBase, uint32(v1676)+2)) = v1680
	v1688 = int32(_a_F_TidStoreSetBlockOffsets_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v1676))) = uint16(v1688)
	v1690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1673)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1676)+2)) = uint8(v1690)
	v1692 = int32(4)
	v1693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1673)+3)))
	v1694 = base.B2i32(base.Ui32(v1669) <= base.Ui32(v1693))
	if base.Ui32(v1669) <= base.Ui32(v1693) {
		v1712 = int32(0)
		goto L259
	} else {
		goto L260
	}
L258:
	;
	v1720 = v1676 + int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v1720+v1694))) = uint8(v1693)
	v1724 = v1676 + int32(36)
	v1725 = int32(2)
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1724+v1694<<(uint(v1725)%32)))) = v1728
	v1731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1673)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1720+v1717))) = uint8(v1731)
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1724+v1717<<(uint(v1725)%32)))) = v1736
	v1739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1673)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1720+v1718))) = uint8(v1739)
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1724+v1718<<(uint(v1725)%32)))) = v1744
	v1747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1673)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1715+v1720))) = uint8(v1747)
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1724+v1715<<(uint(v1725)%32)))) = v1752
	*(*uint8)(unsafe.Add(mBase, uint32(v1720+v1716))) = uint8(v1669)
	v1757 = v1690 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1676)+2)) = uint8(v1757)
	*(*int32)(unsafe.Add(mBase, uint32(v877))) = v1676
	F_pfree(m, v1673)
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L38
	} else {
		goto L273
	}
L259:
	;
	v1715 = v1692
	v1716 = v1712
	v1717 = int32(2)
	v1718 = int32(3)
	goto L258
L260:
	;
	v1695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1673)+4)))
	if base.Ui32(v1669) <= base.Ui32(v1695) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1712 = int32(1)
	goto L259
L262:
	;
	goto L263
L263:
	;
	v1698 = int32(1)
	v1699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1673)+5)))
	if base.Ui32(v1669) <= base.Ui32(v1699) {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1715 = v1692
	v1716 = int32(2)
	v1717 = v1698
	v1718 = int32(3)
	goto L258
L265:
	;
	goto L266
L266:
	;
	v1705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1673)+6)))
	v1706 = base.B2i32(base.Ui32(v1705) < base.Ui32(v1669))
	if base.Ui32(v1705) < base.Ui32(v1669) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1707 = int32(4)
	goto L269
L268:
	;
	v1707 = int32(3)
	goto L269
L269:
	;
	if base.Ui32(v1705) < base.Ui32(v1669) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1710 = int32(3)
	goto L272
L271:
	;
	v1710 = int32(4)
	goto L272
L272:
	;
	v1715 = v1710
	v1716 = v1707
	v1717 = v1698
	v1718 = int32(2)
	goto L258
L273:
	;
	v2060 = v1724 + v1716<<(uint(int32(2))%32)
	goto L11
L274:
	;
	v1767 = v895 + int32(8)
	v1768 = v1292 - v1318
	if v1768&int32(1) != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1773 = v1292 - int32(1)
	v1775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1312+v1773))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1292+v1312))) = uint8(v1775)
	v1777 = int32(2)
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v1767+v1773<<(uint(v1777)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1767+v1292<<(uint(v1777)%32)))) = v1783
	v1786 = v1773
	goto L277
L276:
	;
	v1786 = v1292
	goto L277
L277:
	;
	if v1768 != int32(1) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1793 = v1786
	goto L281
L279:
	;
	goto L280
L280:
	;
	v1852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+2)))
	v1855 = v1318
	v1857 = v1852
	goto L13
L281:
	;
	v1811 = v1793 - int32(1)
	v1812 = v1312 + v1811
	v1813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1812))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1793+v1312))) = uint8(v1813)
	v1815 = int32(2)
	v1816 = v1793 << (uint(v1815) % 32)
	v1820 = v1767 + v1811<<(uint(v1815)%32)
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1820)))
	*(*int32)(unsafe.Add(mBase, uint32(v1767+v1816))) = v1821
	v1824 = v1793 - v1815
	v1826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1312+v1824))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1812))) = uint8(v1826)
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v895+v1816)))
	*(*int32)(unsafe.Add(mBase, uint32(v1820))) = v1829
	if v1318 < v1824 {
		v1793 = v1824
		goto L281
	} else {
		goto L283
	}
L282:
	;
	goto L280
L283:
	;
	goto L282
L284:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2012+v1884))) = uint8(v894)
	v2033 = v2014 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v895)+2)) = uint8(v2033)
	v2060 = v895 + v2012<<(uint(int32(2))%32) + int32(36)
	goto L11
L285:
	;
	v2012 = v1346
	v2014 = v1346
	goto L284
L286:
	;
	goto L287
L287:
	;
	v1888 = v894 & int32(255)
	v1889 = v1346 + v1884
	v1892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1889-int32(1)))))
	if base.Ui32(v1892) < base.Ui32(v1888) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v2012 = v1346
	v2014 = v1346
	goto L284
L289:
	;
	goto L290
L290:
	;
	v1897 = int32(0)
	goto L292
L291:
	;
	if base.Ui32(v1346) <= base.Ui32(v1897) {
		v2012 = v1897
		v2014 = v1346
		goto L284
	} else {
		goto L296
	}
L292:
	;
	v1916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1897+v1884))))
	if base.Ui32(v1888) < base.Ui32(v1916) {
		goto L291
	} else {
		goto L294
	}
L293:
	;
	v2012 = v1346
	v2014 = v1346
	goto L284
L294:
	;
	v1919 = v1897 + int32(1)
	if v1919 != v1346 {
		v1897 = v1919
		goto L292
	} else {
		goto L295
	}
L295:
	;
	goto L293
L296:
	;
	v1923 = v895 + int32(36)
	v1924 = v1346 - v1897
	if v1924&int32(1) != 0 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1928 = v1346 - int32(1)
	v1930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1884+v1928))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1889))) = uint8(v1930)
	v1932 = int32(2)
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1923+v1928<<(uint(v1932)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1923+v1346<<(uint(v1932)%32)))) = v1938
	v1941 = v1928
	goto L299
L298:
	;
	v1941 = v1346
	goto L299
L299:
	;
	if v1924 != int32(1) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1948 = v1941
	goto L303
L301:
	;
	goto L302
L302:
	;
	v2009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+2)))
	v2012 = v1897
	v2014 = v2009
	goto L284
L303:
	;
	v1966 = v1948 - int32(1)
	v1967 = v1884 + v1966
	v1968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1967))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1948+v1884))) = uint8(v1968)
	v1970 = int32(2)
	v1975 = v1923 + v1966<<(uint(v1970)%32)
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v1975)))
	*(*int32)(unsafe.Add(mBase, uint32(v1923+v1948<<(uint(v1970)%32)))) = v1976
	v1979 = v1948 - v1970
	v1981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1884+v1979))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1967))) = uint8(v1981)
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v1923+v1979<<(uint(v1970)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1975))) = v1986
	if v1897 < v1979 {
		v1948 = v1979
		goto L303
	} else {
		goto L305
	}
L304:
	;
	goto L302
L305:
	;
	goto L304
L306:
	;
	v2063 = F_local_ts_extend_down(m, v167, v2060, v166, v883)
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L38
	} else {
		goto L307
	}
L307:
	;
	v2069 = v2063
	goto L10
L308:
	;
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v2110 = *(*int64)(unsafe.Add(mBase, uint32(v2109)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2109)+16)) = v2110 + int64(1)
	goto L1
L309:
	;
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v2093 = F_MemoryContextAlloc(m, v2092, v165)
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L38
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	if v165 != 0 {
		goto L314
	} else {
		goto L315
	}
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2069))) = v2093
	if v165 == int32(0) {
		goto L308
	} else {
		goto L313
	}
L313:
	;
	base.MemoryCopy(m, v2093, v23-int32(-64), v165)
	goto L308
L314:
	;
	base.MemoryCopy(m, v2069, v23-int32(-64), v165)
	goto L316
L315:
	;
	goto L316
L316:
	;
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(v2069)))
	*(*int32)(unsafe.Add(mBase, uint32(v2069))) = v2104 | int32(1)
	goto L308
L317:
	;
	v2136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	if v2118 == v2136 {
		goto L7
	} else {
		goto L318
	}
L318:
	;
	v2139 = v335 + int32(3)
	v2145 = int32(0)
	goto L319
L319:
	;
	v2164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2145+v2139))))
	if base.Ui32(v332&int32(255)) <= base.Ui32(v2164) {
		goto L6
	} else {
		goto L321
	}
L320:
	;
	v2701 = v2118
	v2703 = v2118
	goto L5
L321:
	;
	v2167 = v2145 + int32(1)
	if v2167 != v2118 {
		v2145 = v2167
		goto L319
	} else {
		goto L322
	}
L322:
	;
	goto L320
L323:
	;
	v2191 = *(*int64)(unsafe.Add(mBase, uint32(v23)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v2191
	v2194 = v332 & int32(255)
	v2195 = int32(0)
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v2199 = v23 + int32(32)
	v2200 = *(*int32)(unsafe.Add(mBase, uint32(v2199)+4))
	v2201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2200)+1)))
	if base.Ui32(v2201) <= base.Ui32(int32(30)) {
		goto L327
	} else {
		goto L328
	}
L324:
	;
	v2906 = v2484
	goto L3
L325:
	;
	base.MemoryCopy(m, v2367+int32(268), v2200+int32(36), int32(124))
	*(*int32)(unsafe.Add(mBase, uint32(v2367)+4)) = int32(-1)
	v2452 = int32(31)
	*(*uint8)(unsafe.Add(mBase, uint32(v2374+v2194))) = uint8(v2452)
	v2455 = v2380 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2367)+2)) = uint8(v2455)
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v2364
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v2199)))
	F_dsa_free(m, v2458, v2459)
	mBase = m.M
	v2461 = m.ExcPending
	if v2461 != 0 {
		goto L38
	} else {
		goto L350
	}
L326:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2280+v2307))) = uint8(v2194)
	v2433 = v2223 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2209+int32(2)))) = uint8(v2433)
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v2436 = *(*int32)(unsafe.Add(mBase, uint32(v2199)))
	F_dsa_free(m, v2435, v2436)
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L38
	} else {
		goto L349
	}
L327:
	;
	v2206 = F_dsa_allocate_extended(m, v2197, int32(160), int32(0))
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L38
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	v2364 = F_dsa_allocate_extended(m, v2197, int32(512), int32(0))
	mBase = m.M
	v2365 = m.ExcPending
	if v2365 != 0 {
		goto L38
	} else {
		goto L344
	}
L330:
	;
	v2208 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v2209 = F_dsa_get_address(m, v2208, v2206)
	mBase = m.M
	v2210 = m.ExcPending
	if v2210 != 0 {
		goto L38
	} else {
		goto L331
	}
L331:
	;
	v2211 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2209)+34)) = uint16(v2211)
	v2213 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2209)+26)) = v2213
	*(*int64)(unsafe.Add(mBase, uint32(v2209)+18)) = v2213
	*(*int64)(unsafe.Add(mBase, uint32(v2209)+10)) = v2213
	*(*int64)(unsafe.Add(mBase, uint32(v2209)+2)) = v2213
	v2221 = int32(_a_F_TidStoreSetBlockOffsets_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v2209))) = uint16(v2221)
	v2223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2200)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2209)+2)) = uint8(v2223)
	v2226 = v2200 + int32(3)
	v2230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2223+v2226-int32(1)))))
	if base.Ui32(v2194) <= base.Ui32(v2230) {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	v2302 = int32(36)
	v2303 = v2200 + v2302
	v2305 = v2209 + v2302
	v2307 = v2209 + int32(3)
	v2310 = v2195
	goto L341
L333:
	;
	if v2223 == int32(0) {
		v2280 = v2195
		goto L332
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	v2280 = v2223
	goto L332
L336:
	;
	v2234 = v2195
	goto L337
L337:
	;
	v2255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2234+v2226))))
	if base.Ui32(v2194) < base.Ui32(v2255) {
		v2280 = v2234
		goto L332
	} else {
		goto L339
	}
L338:
	;
	goto L335
L339:
	;
	v2258 = v2234 + int32(1)
	if v2258 != v2223 {
		v2234 = v2258
		goto L337
	} else {
		goto L340
	}
L340:
	;
	goto L338
L341:
	;
	v2329 = v2310 | base.B2i32(base.Ui32(v2280) <= base.Ui32(v2310))
	v2332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2310+v2226))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2307+v2329))) = uint8(v2332)
	v2334 = int32(2)
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v2303+v2310<<(uint(v2334)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2305+v2329<<(uint(v2334)%32)))) = v2340
	if v2310 == int32(14) {
		goto L326
	} else {
		goto L343
	}
L343:
	;
	v2345 = v2310 | int32(1)
	v2347 = v2345 + base.B2i32(base.Ui32(v2280) <= base.Ui32(v2345))
	v2350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2345+v2226))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2307+v2347))) = uint8(v2350)
	v2352 = int32(2)
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v2303+v2345<<(uint(v2352)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2305+v2347<<(uint(v2352)%32)))) = v2358
	v2310 = v2310 + v2352
	goto L341
L344:
	;
	v2366 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v2367 = F_dsa_get_address(m, v2366, v2364)
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L38
	} else {
		goto L345
	}
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2367)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2367))) = int64(0)
	v2374 = v2367 + int32(12)
	base.MemoryFill(m, v2374, int32(255), int32(256))
	v2378 = int32(_a_F_TidStoreSetBlockOffsets_9)
	*(*uint16)(unsafe.Add(mBase, uint32(v2367))) = uint16(v2378)
	v2380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2200)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2367)+2)) = uint8(v2380)
	v2383 = v2200 + int32(3)
	v2386 = v2195
	goto L346
L346:
	;
	v2405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2386+v2383))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2374+v2405))) = uint8(v2386)
	v2409 = v2386 | int32(1)
	v2411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2383+v2409))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2374+v2411))) = uint8(v2409)
	v2415 = v2386 | int32(2)
	v2417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2383+v2415))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2374+v2417))) = uint8(v2415)
	if v2386 == int32(28) {
		goto L325
	} else {
		goto L348
	}
L348:
	;
	v2423 = v2386 | int32(3)
	v2425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2383+v2423))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2374+v2425))) = uint8(v2423)
	v2386 = v2386 + int32(4)
	goto L346
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v2206
	v2484 = v2305 + v2280<<(uint(int32(2))%32)
	goto L324
L350:
	;
	v2484 = v2367 + int32(392)
	goto L324
L351:
	;
	v2518 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v2519 = F_dsa_get_address(m, v2518, v2516)
	mBase = m.M
	v2520 = m.ExcPending
	if v2520 != 0 {
		goto L38
	} else {
		goto L352
	}
L352:
	;
	v2521 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2519)+34)) = uint16(v2521)
	v2523 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2519)+26)) = v2523
	*(*int64)(unsafe.Add(mBase, uint32(v2519)+18)) = v2523
	*(*int64)(unsafe.Add(mBase, uint32(v2519)+10)) = v2523
	*(*int64)(unsafe.Add(mBase, uint32(v2519)+2)) = v2523
	v2531 = int32(3841)
	*(*uint16)(unsafe.Add(mBase, uint32(v2519))) = uint16(v2531)
	v2533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2512)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2519)+2)) = uint8(v2533)
	v2535 = int32(4)
	v2536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2512)+3)))
	v2537 = base.B2i32(base.Ui32(v2508) <= base.Ui32(v2536))
	if base.Ui32(v2508) <= base.Ui32(v2536) {
		v2555 = v2509
		goto L354
	} else {
		goto L355
	}
L353:
	;
	v2564 = v2519 + int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v2564+v2537))) = uint8(v2536)
	v2568 = v2519 + int32(36)
	v2569 = int32(2)
	v2572 = *(*int32)(unsafe.Add(mBase, uint32(v2512)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2568+v2537<<(uint(v2569)%32)))) = v2572
	v2575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2512)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2558+v2564))) = uint8(v2575)
	v2580 = *(*int32)(unsafe.Add(mBase, uint32(v2512)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2568+v2558<<(uint(v2569)%32)))) = v2580
	v2583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2512)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2564+v2562))) = uint8(v2583)
	v2588 = *(*int32)(unsafe.Add(mBase, uint32(v2512)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2568+v2562<<(uint(v2569)%32)))) = v2588
	v2591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2512)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2564+v2561))) = uint8(v2591)
	v2596 = *(*int32)(unsafe.Add(mBase, uint32(v2512)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2568+v2561<<(uint(v2569)%32)))) = v2596
	*(*uint8)(unsafe.Add(mBase, uint32(v2559+v2564))) = uint8(v2508)
	v2601 = v2533 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2519)+2)) = uint8(v2601)
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v2516
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v2605 = *(*int32)(unsafe.Add(mBase, uint32(v2511)))
	F_dsa_free(m, v2604, v2605)
	mBase = m.M
	v2607 = m.ExcPending
	if v2607 != 0 {
		goto L38
	} else {
		goto L368
	}
L354:
	;
	v2558 = int32(2)
	v2559 = v2555
	v2561 = v2535
	v2562 = int32(3)
	goto L353
L355:
	;
	v2538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2512)+4)))
	if base.Ui32(v2508) <= base.Ui32(v2538) {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v2555 = int32(1)
	goto L354
L357:
	;
	goto L358
L358:
	;
	v2541 = int32(1)
	v2542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2512)+5)))
	if base.Ui32(v2508) <= base.Ui32(v2542) {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v2558 = v2541
	v2559 = int32(2)
	v2561 = v2535
	v2562 = int32(3)
	goto L353
L360:
	;
	goto L361
L361:
	;
	v2548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2512)+6)))
	v2549 = base.B2i32(base.Ui32(v2548) < base.Ui32(v2508))
	if base.Ui32(v2548) < base.Ui32(v2508) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v2550 = int32(4)
	goto L364
L363:
	;
	v2550 = int32(3)
	goto L364
L364:
	;
	if base.Ui32(v2548) < base.Ui32(v2508) {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v2553 = int32(3)
	goto L367
L366:
	;
	v2553 = int32(4)
	goto L367
L367:
	;
	v2558 = v2541
	v2559 = v2550
	v2561 = v2553
	v2562 = int32(2)
	goto L353
L368:
	;
	v2906 = v2568 + v2559<<(uint(int32(2))%32)
	goto L3
L369:
	;
	v2613 = v335 + int32(8)
	v2614 = v2118 - v2145
	if v2614&int32(1) != 0 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v2619 = v2118 - int32(1)
	v2621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2139+v2619))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2118+v2139))) = uint8(v2621)
	v2623 = int32(2)
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v2613+v2619<<(uint(v2623)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2613+v2118<<(uint(v2623)%32)))) = v2629
	v2632 = v2619
	goto L372
L371:
	;
	v2632 = v2118
	goto L372
L372:
	;
	if v2614 != int32(1) {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v2639 = v2632
	goto L376
L374:
	;
	goto L375
L375:
	;
	v2698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	v2701 = v2145
	v2703 = v2698
	goto L5
L376:
	;
	v2657 = v2639 - int32(1)
	v2658 = v2139 + v2657
	v2659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2658))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2639+v2139))) = uint8(v2659)
	v2661 = int32(2)
	v2662 = v2639 << (uint(v2661) % 32)
	v2666 = v2613 + v2657<<(uint(v2661)%32)
	v2667 = *(*int32)(unsafe.Add(mBase, uint32(v2666)))
	*(*int32)(unsafe.Add(mBase, uint32(v2613+v2662))) = v2667
	v2670 = v2639 - v2661
	v2672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2139+v2670))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2658))) = uint8(v2672)
	v2675 = *(*int32)(unsafe.Add(mBase, uint32(v335+v2662)))
	*(*int32)(unsafe.Add(mBase, uint32(v2666))) = v2675
	if v2145 < v2670 {
		v2639 = v2670
		goto L376
	} else {
		goto L378
	}
L377:
	;
	goto L375
L378:
	;
	goto L377
L379:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2858+v2730))) = uint8(v332)
	v2879 = v2860 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)) = uint8(v2879)
	v2906 = v335 + v2858<<(uint(int32(2))%32) + int32(36)
	goto L3
L380:
	;
	v2858 = v2173
	v2860 = v2173
	goto L379
L381:
	;
	goto L382
L382:
	;
	v2734 = v332 & int32(255)
	v2735 = v2173 + v2730
	v2738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2735-int32(1)))))
	if base.Ui32(v2738) < base.Ui32(v2734) {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v2858 = v2173
	v2860 = v2173
	goto L379
L384:
	;
	goto L385
L385:
	;
	v2743 = int32(0)
	goto L387
L386:
	;
	if base.Ui32(v2173) <= base.Ui32(v2743) {
		v2858 = v2743
		v2860 = v2173
		goto L379
	} else {
		goto L391
	}
L387:
	;
	v2762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2743+v2730))))
	if base.Ui32(v2734) < base.Ui32(v2762) {
		goto L386
	} else {
		goto L389
	}
L388:
	;
	v2858 = v2173
	v2860 = v2173
	goto L379
L389:
	;
	v2765 = v2743 + int32(1)
	if v2765 != v2173 {
		v2743 = v2765
		goto L387
	} else {
		goto L390
	}
L390:
	;
	goto L388
L391:
	;
	v2769 = v335 + int32(36)
	v2770 = v2173 - v2743
	if v2770&int32(1) != 0 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v2774 = v2173 - int32(1)
	v2776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2730+v2774))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2735))) = uint8(v2776)
	v2778 = int32(2)
	v2784 = *(*int32)(unsafe.Add(mBase, uint32(v2769+v2774<<(uint(v2778)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2769+v2173<<(uint(v2778)%32)))) = v2784
	v2787 = v2774
	goto L394
L393:
	;
	v2787 = v2173
	goto L394
L394:
	;
	if v2770 != int32(1) {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v2794 = v2787
	goto L398
L396:
	;
	goto L397
L397:
	;
	v2855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	v2858 = v2743
	v2860 = v2855
	goto L379
L398:
	;
	v2812 = v2794 - int32(1)
	v2813 = v2730 + v2812
	v2814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2813))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2794+v2730))) = uint8(v2814)
	v2816 = int32(2)
	v2821 = v2769 + v2812<<(uint(v2816)%32)
	v2822 = *(*int32)(unsafe.Add(mBase, uint32(v2821)))
	*(*int32)(unsafe.Add(mBase, uint32(v2769+v2794<<(uint(v2816)%32)))) = v2822
	v2825 = v2794 - v2816
	v2827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2730+v2825))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2813))) = uint8(v2827)
	v2832 = *(*int32)(unsafe.Add(mBase, uint32(v2769+v2825<<(uint(v2816)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2821))) = v2832
	if v2743 < v2825 {
		v2794 = v2825
		goto L398
	} else {
		goto L400
	}
L399:
	;
	goto L397
L400:
	;
	goto L399
L401:
	;
	v2909 = F_shared_ts_extend_down(m, v167, v2906, v166, v320)
	mBase = m.M
	v2910 = m.ExcPending
	if v2910 != 0 {
		goto L38
	} else {
		goto L402
	}
L402:
	;
	v2915 = v2909
	goto L2
L403:
	;
	v2959 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v2960 = *(*int64)(unsafe.Add(mBase, uint32(v2959)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2959)+40)) = v2960 + int64(1)
	goto L1
L404:
	;
	v2938 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v2940 = F_dsa_allocate_extended(m, v2938, v165, int32(0))
	mBase = m.M
	v2941 = m.ExcPending
	if v2941 != 0 {
		goto L38
	} else {
		goto L407
	}
L405:
	;
	goto L406
L406:
	;
	if v165 != 0 {
		goto L410
	} else {
		goto L411
	}
L407:
	;
	v2942 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v2943 = F_dsa_get_address(m, v2942, v2940)
	mBase = m.M
	v2944 = m.ExcPending
	if v2944 != 0 {
		goto L38
	} else {
		goto L408
	}
L408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2915))) = v2940
	if v165 == int32(0) {
		goto L403
	} else {
		goto L409
	}
L409:
	;
	base.MemoryCopy(m, v2943, v23-int32(-64), v165)
	goto L403
L410:
	;
	base.MemoryCopy(m, v2915, v23-int32(-64), v165)
	goto L412
L411:
	;
	goto L412
L412:
	;
	v2954 = *(*int32)(unsafe.Add(mBase, uint32(v2915)))
	*(*int32)(unsafe.Add(mBase, uint32(v2915))) = v2954 | int32(1)
	goto L403
}
