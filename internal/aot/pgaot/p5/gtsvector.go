package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtsvector_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v233 int64
	_ = v233
	var v234 int32
	_ = v234
	var v237 int64
	_ = v237
	var v238 int32
	_ = v238
	var v241 int64
	_ = v241
	var v242 int32
	_ = v242
	var v245 int64
	_ = v245
	var v246 int32
	_ = v246
	var v249 int64
	_ = v249
	var v253 int64
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v273 int64
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v291 int64
	_ = v291
	var v292 int32
	_ = v292
	var v295 int64
	_ = v295
	var v296 int64
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
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
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
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
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
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
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v424 int64
	_ = v424
	var v425 int32
	_ = v425
	var v428 int64
	_ = v428
	var v429 int32
	_ = v429
	var v432 int64
	_ = v432
	var v433 int32
	_ = v433
	var v436 int64
	_ = v436
	var v437 int32
	_ = v437
	var v440 int64
	_ = v440
	var v444 int64
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v464 int64
	_ = v464
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v482 int64
	_ = v482
	var v483 int32
	_ = v483
	var v486 int64
	_ = v486
	var v487 int64
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v527 int64
	_ = v527
	var v528 int32
	_ = v528
	var v531 int64
	_ = v531
	var v532 int32
	_ = v532
	var v535 int64
	_ = v535
	var v536 int32
	_ = v536
	var v539 int64
	_ = v539
	var v540 int32
	_ = v540
	var v543 int64
	_ = v543
	var v547 int64
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v567 int64
	_ = v567
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v585 int64
	_ = v585
	var v586 int32
	_ = v586
	var v589 int64
	_ = v589
	var v590 int64
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v682 int64
	_ = v682
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v709 int64
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v725 int64
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v739 int64
	_ = v739
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int64
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int64
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v767 int64
	_ = v767
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int64
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int64
	_ = v787
	var v788 int64
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v800 int64
	_ = v800
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v809 int64
	_ = v809
	var v810 int32
	_ = v810
	var v813 int64
	_ = v813
	var v814 int32
	_ = v814
	var v817 int64
	_ = v817
	var v818 int32
	_ = v818
	var v821 int64
	_ = v821
	var v822 int32
	_ = v822
	var v825 int64
	_ = v825
	var v829 int64
	_ = v829
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v840 int64
	_ = v840
	var v845 int64
	_ = v845
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v872 int64
	_ = v872
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v888 int64
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v902 int64
	_ = v902
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int64
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int64
	_ = v918
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v930 int64
	_ = v930
	var v934 int32
	_ = v934
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int64
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v950 int64
	_ = v950
	var v951 int64
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v963 int64
	_ = v963
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v972 int64
	_ = v972
	var v973 int32
	_ = v973
	var v976 int64
	_ = v976
	var v977 int32
	_ = v977
	var v980 int64
	_ = v980
	var v981 int32
	_ = v981
	var v984 int64
	_ = v984
	var v985 int32
	_ = v985
	var v988 int64
	_ = v988
	var v992 int64
	_ = v992
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v1003 int64
	_ = v1003
	var v1008 int64
	_ = v1008
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1035 int64
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1051 int64
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1065 int64
	_ = v1065
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int64
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1081 int64
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1093 int64
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1107 int64
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int64
	_ = v1113
	var v1114 int64
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1126 int64
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1135 int64
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int64
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1143 int64
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 int64
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int64
	_ = v1151
	var v1155 int64
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1166 int64
	_ = v1166
	var v1181 int64
	_ = v1181
	var v1198 int64
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1233 int64
	_ = v1233
	var v1256 float32
	_ = v1256
	var v1259 int32
	_ = v1259
	v2 = int32(0)
	v15 = int64(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20 == v2 {
		v37 = v2
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
		if v24 == int32(0) {
			v37 = v2
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
			if v27 != int32(7) {
				v37 = v2
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				if v30 != int32(17) {
					v37 = v2
				} else {
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+24)))
					v37 = v33 ^ int32(1)
				}
			}
		}
	}
	if v37&int32(1) != 0 {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v41 = F_get_fn_opclass_options(m, v40)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
			v46 = v45
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(0)
			v52 = v48 + int32(8)
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
			if v53&int32(1) != 0 {
				v56 = F_palloc(m, v46)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
					v62 = int32(base.Ui32(v58)>>(uint(int32(2))%32)) - int32(8)
					if v56&int32(3) != 0 {
						v82 = v46
						v86 = F__emscripten_memset_bulkmem(m, v56, base.I32_extend8_s(int32(0)), v82)
						mBase = m.M
					} else {
						if base.Ui32(int32(1024)) < base.Ui32(v46) {
							v82 = v46
							v86 = F__emscripten_memset_bulkmem(m, v56, base.I32_extend8_s(int32(0)), v82)
							mBase = m.M
						} else {
							if v46&int32(3) != 0 {
								v82 = v46
								v86 = F__emscripten_memset_bulkmem(m, v56, base.I32_extend8_s(int32(0)), v82)
								mBase = m.M
							} else {
								v69 = v56 + v46
								if base.Ui32(v69) <= base.Ui32(v56) {
								} else {
									v74 = v56 + int32(4)
									if base.Ui32(v74) < base.Ui32(v69) {
										v76 = v69
									} else {
										v76 = v74
									}
									v82 = (v56^int32(-1)+v76)&int32(-4) + int32(4)
									v86 = F__emscripten_memset_bulkmem(m, v56, base.I32_extend8_s(int32(0)), v82)
									mBase = m.M
								}
							}
						}
					}
					if base.Ui32(v62) < base.Ui32(int32(4)) {
					} else {
						v92 = v47 + int32(8)
						v94 = v46 << (uint(int32(3)) % 32)
						v95 = int32(1)
						v97 = int32(base.Ui32(v62) >> (uint(int32(2)) % 32))
						if base.Ui32(v97) <= base.Ui32(v95) {
							v100 = v95
						} else {
							v100 = v97
						}
						v103 = int32(0)
						if base.Ui32(int32(8)) <= base.Ui32(v62) {
							v110 = v103
							v113 = int32(0)
							for {
								v124 = int32(2)
								v126 = v92 + v110<<(uint(v124)%32)
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
								v128 = base.I32_rem_u_s(v127, v94)
								v129 = int32(3)
								v131 = v56 + int32(base.Ui32(v128)>>(uint(v129)%32))
								v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
								v133 = int32(1)
								v134 = int32(7)
								v137 = v132 | v133<<(uint(v128&v134)%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v137)
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
								v140 = base.I32_rem_u_s(v139, v94)
								v143 = v56 + int32(base.Ui32(v140)>>(uint(v129)%32))
								v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
								v149 = v144 | v133<<(uint(v140&v134)%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v149)
								v152 = v110 + v124
								v154 = v113 + v124
								if v154 != v100&int32(1073741822) {
									v110 = v152
									v113 = v154
									continue
								} else {
									break
								}
								break
							}
							v157 = v152
						} else {
							v157 = v103
						}
						if v100&int32(1) == int32(0) {
						} else {
							v176 = *(*int32)(unsafe.Add(mBase, uint32(v92+v157<<(uint(int32(2))%32))))
							v177 = base.I32_rem_u_s(v176, v94)
							v180 = v56 + int32(base.Ui32(v177)>>(uint(int32(3))%32))
							v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
							v186 = v181 | int32(1)<<(uint(v177&int32(7))%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v186)
						}
					}
					v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
					if v203&int32(4) != 0 {
						v206 = int32(3)
						v207 = v46 << (uint(v206) % 32)
						if v206 < v46 {
							v682 = int64(0)
							if v46 < int32(4) {
								v761 = v56
								v762 = v46
								v767 = v682
							} else {
								if v56 != (v56+int32(3))&int32(-4) {
									v761 = v56
									v762 = v46
									v767 = v682
								} else {
									v691 = v46 - int32(4)
									v695 = int32(base.Ui32(v691)>>(uint(int32(2))%32)) + int32(1)
									v697 = v695 & int32(3)
									if base.Ui32(v691) < base.Ui32(int32(12)) {
										v733 = v56
										v734 = v46
										v739 = v682
									} else {
										v703 = v56
										v704 = v46
										v705 = int32(0)
										v709 = v682
										for {
											v710 = *(*int32)(unsafe.Add(mBase, uint32(v703)+12))
											v713 = *(*int32)(unsafe.Add(mBase, uint32(v703)+8))
											v716 = *(*int32)(unsafe.Add(mBase, uint32(v703)+4))
											v719 = *(*int32)(unsafe.Add(mBase, uint32(v703)))
											v725 = base.I64_extend_i32_u(base.I32_popcnt(v710)) + (base.I64_extend_i32_u(base.I32_popcnt(v713)) + (base.I64_extend_i32_u(base.I32_popcnt(v716)) + (v709 + base.I64_extend_i32_u(base.I32_popcnt(v719)))))
											v726 = int32(16)
											v727 = v704 - v726
											v729 = v703 + v726
											v731 = v705 + int32(4)
											if v731 != v695&int32(2147483644) {
												v703 = v729
												v704 = v727
												v705 = v731
												v709 = v725
												continue
											} else {
												break
											}
											break
										}
										v733 = v729
										v734 = v727
										v739 = v725
									}
									if v697 == int32(0) {
										v761 = v733
										v762 = v734
										v767 = v739
									} else {
										v744 = v734
										v745 = v733
										v746 = int32(0)
										v749 = v739
										for {
											v750 = int32(4)
											v751 = v744 - v750
											v752 = *(*int32)(unsafe.Add(mBase, uint32(v745)))
											v755 = v749 + base.I64_extend_i32_u(base.I32_popcnt(v752))
											v757 = v745 + v750
											v759 = v746 + int32(1)
											if v759 != v697 {
												v744 = v751
												v745 = v757
												v746 = v759
												v749 = v755
												continue
											} else {
												break
											}
											break
										}
										v761 = v757
										v762 = v751
										v767 = v755
									}
								}
							}
							if v762 == int32(0) {
								v840 = v767
							} else {
								v771 = v762 & int32(3)
								if v771 == int32(0) {
									v794 = v761
									v796 = v762
									v800 = v767
								} else {
									v777 = v762
									v778 = v761
									v779 = int32(0)
									v781 = v767
									for {
										v782 = int32(1)
										v783 = v777 - v782
										v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778))))
										v787 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v784)+uint32(_consts[1104]))))
										v788 = v781 + v787
										v790 = v778 + v782
										v792 = v779 + v782
										if v792 != v771 {
											v777 = v783
											v778 = v790
											v779 = v792
											v781 = v788
											continue
										} else {
											break
										}
										break
									}
									v794 = v790
									v796 = v783
									v800 = v788
								}
								if base.Ui32(v762) < base.Ui32(int32(4)) {
									v840 = v800
								} else {
									v803 = v794
									v805 = v796
									v809 = v800
									for {
										v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+3)))
										v813 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v810)+uint32(_consts[1104]))))
										v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+2)))
										v817 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v814)+uint32(_consts[1104]))))
										v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+1)))
										v821 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v818)+uint32(_consts[1104]))))
										v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803))))
										v825 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v822)+uint32(_consts[1104]))))
										v829 = v813 + (v817 + (v821 + (v809 + v825)))
										v830 = int32(4)
										v833 = v805 - v830
										if v833 != 0 {
											v803 = v803 + v830
											v805 = v833
											v809 = v829
											continue
										} else {
											break
										}
										break
									}
									v840 = v829
								}
							}
							v1233 = v840
						} else {
							if v46 == int32(0) {
								v1233 = v15
							} else {
								v213 = v46 & int32(3)
								if base.Ui32(v46) < base.Ui32(int32(4)) {
									v261 = v56
									v273 = v15
								} else {
									v220 = int32(0)
									v221 = v56
									v233 = v15
									for {
										v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+3)))
										v237 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v234)+uint32(_consts[1104]))))
										v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+2)))
										v241 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v238)+uint32(_consts[1104]))))
										v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
										v245 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v242)+uint32(_consts[1104]))))
										v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
										v249 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v246)+uint32(_consts[1104]))))
										v253 = v237 + (v241 + (v245 + (v233 + v249)))
										v254 = int32(4)
										v255 = v221 + v254
										v257 = v220 + v254
										if v257 != v46&int32(-4) {
											v220 = v257
											v221 = v255
											v233 = v253
											continue
										} else {
											break
										}
										break
									}
									v261 = v255
									v273 = v253
								}
								if v213 == int32(0) {
									v1233 = v273
								} else {
									v278 = int32(0)
									v279 = v261
									v291 = v273
									for {
										v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
										v295 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v292)+uint32(_consts[1104]))))
										v296 = v291 + v295
										v297 = int32(1)
										v300 = v278 + v297
										if v300 != v213 {
											v278 = v300
											v279 = v279 + v297
											v291 = v296
											continue
										} else {
											break
										}
										break
									}
									v1233 = v296
								}
							}
						}
						v1256 = base.F32_div(base.F32_convert_i32_s(v207-base.I32_wrap_i64(v1233)), base.F32_convert_i32_s(v207|int32(1)))
					} else {
						if v46 <= int32(0) {
							v1256 = float32(0)
						} else {
							v305 = int32(1)
							if v46 == v305 {
								v309 = int32(0)
								v357 = v309
								v359 = v309
							} else {
								v313 = int32(0)
								v317 = v313
								v319 = v313
								v320 = v313
								for {
									v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317+v52))))
									v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v317))))
									v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332^v334)+uint32(_consts[1104]))))
									v341 = v317 | int32(1)
									v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v341))))
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v341))))
									v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343^v345)+uint32(_consts[1104]))))
									v350 = v319 + v338 + v349
									v351 = int32(2)
									v352 = v317 + v351
									v354 = v320 + v351
									if v354 != v46&int32(2147483646) {
										v317 = v352
										v319 = v350
										v320 = v354
										continue
									} else {
										break
									}
									break
								}
								v357 = v352
								v359 = v350
							}
							if v46&v305 != 0 {
								v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357+v52))))
								v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v357))))
								v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372^v374)+uint32(_consts[1104]))))
								v380 = v359 + v378
							} else {
								v380 = v359
							}
							v1256 = base.F32_convert_i32_s(v380)
						}
					}
					*(*float32)(unsafe.Add(mBase, uint32(v16))) = v1256
					F_pfree(m, v56)
					mBase = m.M
					v1259 = m.ExcPending
					if v1259 != 0 {
						return int32(0)
					} else {
						return v16
					}
				}
			} else {
				v382 = int32(4)
				v383 = v53 & v382
				v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
				if v384&v382 != 0 {
					if v383 != 0 {
						v1202 = int32(0)
					} else {
						v388 = int32(8)
						v389 = v47 + v388
						v390 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
						v392 = int32(base.Ui32(v390) >> (uint(int32(2)) % 32))
						v394 = v392 - v388
						if base.Ui32(int32(47)) < base.Ui32(v390) {
							v845 = int64(0)
							if v394 < int32(4) {
								v924 = v389
								v925 = v394
								v930 = v845
							} else {
								if v389 != (v47+int32(11))&int32(-4) {
									v924 = v389
									v925 = v394
									v930 = v845
								} else {
									v854 = v394 - int32(4)
									v858 = int32(base.Ui32(v854)>>(uint(int32(2))%32)) + int32(1)
									v860 = v858 & int32(3)
									if base.Ui32(v854) < base.Ui32(int32(12)) {
										v896 = v389
										v897 = v394
										v902 = v845
									} else {
										v866 = v389
										v867 = v394
										v868 = int32(0)
										v872 = v845
										for {
											v873 = *(*int32)(unsafe.Add(mBase, uint32(v866)+12))
											v876 = *(*int32)(unsafe.Add(mBase, uint32(v866)+8))
											v879 = *(*int32)(unsafe.Add(mBase, uint32(v866)+4))
											v882 = *(*int32)(unsafe.Add(mBase, uint32(v866)))
											v888 = base.I64_extend_i32_u(base.I32_popcnt(v873)) + (base.I64_extend_i32_u(base.I32_popcnt(v876)) + (base.I64_extend_i32_u(base.I32_popcnt(v879)) + (v872 + base.I64_extend_i32_u(base.I32_popcnt(v882)))))
											v889 = int32(16)
											v890 = v867 - v889
											v892 = v866 + v889
											v894 = v868 + int32(4)
											if v894 != v858&int32(2147483644) {
												v866 = v892
												v867 = v890
												v868 = v894
												v872 = v888
												continue
											} else {
												break
											}
											break
										}
										v896 = v892
										v897 = v890
										v902 = v888
									}
									if v860 == int32(0) {
										v924 = v896
										v925 = v897
										v930 = v902
									} else {
										v907 = v897
										v908 = v896
										v909 = int32(0)
										v912 = v902
										for {
											v913 = int32(4)
											v914 = v907 - v913
											v915 = *(*int32)(unsafe.Add(mBase, uint32(v908)))
											v918 = v912 + base.I64_extend_i32_u(base.I32_popcnt(v915))
											v920 = v908 + v913
											v922 = v909 + int32(1)
											if v922 != v860 {
												v907 = v914
												v908 = v920
												v909 = v922
												v912 = v918
												continue
											} else {
												break
											}
											break
										}
										v924 = v920
										v925 = v914
										v930 = v918
									}
								}
							}
							if v925 == int32(0) {
								v1003 = v930
							} else {
								v934 = v925 & int32(3)
								if v934 == int32(0) {
									v957 = v924
									v959 = v925
									v963 = v930
								} else {
									v940 = v925
									v941 = v924
									v942 = int32(0)
									v944 = v930
									for {
										v945 = int32(1)
										v946 = v940 - v945
										v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v941))))
										v950 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v947)+uint32(_consts[1104]))))
										v951 = v944 + v950
										v953 = v941 + v945
										v955 = v942 + v945
										if v955 != v934 {
											v940 = v946
											v941 = v953
											v942 = v955
											v944 = v951
											continue
										} else {
											break
										}
										break
									}
									v957 = v953
									v959 = v946
									v963 = v951
								}
								if base.Ui32(v925) < base.Ui32(int32(4)) {
									v1003 = v963
								} else {
									v966 = v957
									v968 = v959
									v972 = v963
									for {
										v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v966)+3)))
										v976 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v973)+uint32(_consts[1104]))))
										v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v966)+2)))
										v980 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v977)+uint32(_consts[1104]))))
										v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v966)+1)))
										v984 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v981)+uint32(_consts[1104]))))
										v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v966))))
										v988 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v985)+uint32(_consts[1104]))))
										v992 = v976 + (v980 + (v984 + (v972 + v988)))
										v993 = int32(4)
										v996 = v968 - v993
										if v996 != 0 {
											v966 = v966 + v993
											v968 = v996
											v972 = v992
											continue
										} else {
											break
										}
										break
									}
									v1003 = v992
								}
							}
							v1198 = v1003
						} else {
							if v394 == int32(0) {
								v1198 = v15
							} else {
								v401 = int32(3)
								v402 = v392 & v401
								if base.Ui32(v392-int32(9)) < base.Ui32(v401) {
									v450 = v389
									v464 = v15
								} else {
									v410 = v389
									v412 = int32(0)
									v424 = v15
									for {
										v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+3)))
										v428 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v425)+uint32(_consts[1104]))))
										v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+2)))
										v432 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v429)+uint32(_consts[1104]))))
										v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+1)))
										v436 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v433)+uint32(_consts[1104]))))
										v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
										v440 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v437)+uint32(_consts[1104]))))
										v444 = v428 + (v432 + (v436 + (v424 + v440)))
										v445 = int32(4)
										v446 = v410 + v445
										v448 = v412 + v445
										if v448 != v394&int32(-4) {
											v410 = v446
											v412 = v448
											v424 = v444
											continue
										} else {
											break
										}
										break
									}
									v450 = v446
									v464 = v444
								}
								if v402 == int32(0) {
									v1198 = v464
								} else {
									v468 = v450
									v470 = int32(0)
									v482 = v464
									for {
										v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
										v486 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v483)+uint32(_consts[1104]))))
										v487 = v482 + v486
										v488 = int32(1)
										v491 = v470 + v488
										if v491 != v402 {
											v468 = v468 + v488
											v470 = v491
											v482 = v487
											continue
										} else {
											break
										}
										break
									}
									v1198 = v487
								}
							}
						}
						v1202 = v394<<(uint(int32(3))%32) - base.I32_wrap_i64(v1198)
					}
				} else {
					v493 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
					v495 = int32(base.Ui32(v493) >> (uint(int32(2)) % 32))
					v497 = v495 - int32(8)
					if v383 != 0 {
						if base.Ui32(int32(47)) < base.Ui32(v493) {
							v1008 = int64(0)
							if v497 < int32(4) {
								v1087 = v52
								v1088 = v497
								v1093 = v1008
							} else {
								if v52 != (v48+int32(11))&int32(-4) {
									v1087 = v52
									v1088 = v497
									v1093 = v1008
								} else {
									v1017 = v497 - int32(4)
									v1021 = int32(base.Ui32(v1017)>>(uint(int32(2))%32)) + int32(1)
									v1023 = v1021 & int32(3)
									if base.Ui32(v1017) < base.Ui32(int32(12)) {
										v1059 = v52
										v1060 = v497
										v1065 = v1008
									} else {
										v1029 = v52
										v1030 = v497
										v1031 = int32(0)
										v1035 = v1008
										for {
											v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+12))
											v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+8))
											v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+4))
											v1045 = *(*int32)(unsafe.Add(mBase, uint32(v1029)))
											v1051 = base.I64_extend_i32_u(base.I32_popcnt(v1036)) + (base.I64_extend_i32_u(base.I32_popcnt(v1039)) + (base.I64_extend_i32_u(base.I32_popcnt(v1042)) + (v1035 + base.I64_extend_i32_u(base.I32_popcnt(v1045)))))
											v1052 = int32(16)
											v1053 = v1030 - v1052
											v1055 = v1029 + v1052
											v1057 = v1031 + int32(4)
											if v1057 != v1021&int32(2147483644) {
												v1029 = v1055
												v1030 = v1053
												v1031 = v1057
												v1035 = v1051
												continue
											} else {
												break
											}
											break
										}
										v1059 = v1055
										v1060 = v1053
										v1065 = v1051
									}
									if v1023 == int32(0) {
										v1087 = v1059
										v1088 = v1060
										v1093 = v1065
									} else {
										v1070 = v1060
										v1071 = v1059
										v1072 = int32(0)
										v1075 = v1065
										for {
											v1076 = int32(4)
											v1077 = v1070 - v1076
											v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1071)))
											v1081 = v1075 + base.I64_extend_i32_u(base.I32_popcnt(v1078))
											v1083 = v1071 + v1076
											v1085 = v1072 + int32(1)
											if v1085 != v1023 {
												v1070 = v1077
												v1071 = v1083
												v1072 = v1085
												v1075 = v1081
												continue
											} else {
												break
											}
											break
										}
										v1087 = v1083
										v1088 = v1077
										v1093 = v1081
									}
								}
							}
							if v1088 == int32(0) {
								v1166 = v1093
							} else {
								v1097 = v1088 & int32(3)
								if v1097 == int32(0) {
									v1120 = v1087
									v1122 = v1088
									v1126 = v1093
								} else {
									v1103 = v1088
									v1104 = v1087
									v1105 = int32(0)
									v1107 = v1093
									for {
										v1108 = int32(1)
										v1109 = v1103 - v1108
										v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1104))))
										v1113 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1110)+uint32(_consts[1104]))))
										v1114 = v1107 + v1113
										v1116 = v1104 + v1108
										v1118 = v1105 + v1108
										if v1118 != v1097 {
											v1103 = v1109
											v1104 = v1116
											v1105 = v1118
											v1107 = v1114
											continue
										} else {
											break
										}
										break
									}
									v1120 = v1116
									v1122 = v1109
									v1126 = v1114
								}
								if base.Ui32(v1088) < base.Ui32(int32(4)) {
									v1166 = v1126
								} else {
									v1129 = v1120
									v1131 = v1122
									v1135 = v1126
									for {
										v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129)+3)))
										v1139 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1136)+uint32(_consts[1104]))))
										v1140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129)+2)))
										v1143 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1140)+uint32(_consts[1104]))))
										v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129)+1)))
										v1147 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+uint32(_consts[1104]))))
										v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129))))
										v1151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1148)+uint32(_consts[1104]))))
										v1155 = v1139 + (v1143 + (v1147 + (v1135 + v1151)))
										v1156 = int32(4)
										v1159 = v1131 - v1156
										if v1159 != 0 {
											v1129 = v1129 + v1156
											v1131 = v1159
											v1135 = v1155
											continue
										} else {
											break
										}
										break
									}
									v1166 = v1155
								}
							}
							v1181 = v1166
						} else {
							if v497 == int32(0) {
								v1181 = v15
							} else {
								v504 = int32(3)
								v505 = v495 & v504
								if base.Ui32(v495-int32(9)) < base.Ui32(v504) {
									v555 = v52
									v567 = v15
								} else {
									v513 = int32(0)
									v515 = v52
									v527 = v15
									for {
										v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+3)))
										v531 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v528)+uint32(_consts[1104]))))
										v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+2)))
										v535 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v532)+uint32(_consts[1104]))))
										v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
										v539 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v536)+uint32(_consts[1104]))))
										v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
										v543 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v540)+uint32(_consts[1104]))))
										v547 = v531 + (v535 + (v539 + (v527 + v543)))
										v548 = int32(4)
										v549 = v515 + v548
										v551 = v513 + v548
										if v551 != v497&int32(-4) {
											v513 = v551
											v515 = v549
											v527 = v547
											continue
										} else {
											break
										}
										break
									}
									v555 = v549
									v567 = v547
								}
								if v505 == int32(0) {
									v1181 = v567
								} else {
									v571 = int32(0)
									v573 = v555
									v585 = v567
									for {
										v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573))))
										v589 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v586)+uint32(_consts[1104]))))
										v590 = v585 + v589
										v591 = int32(1)
										v594 = v571 + v591
										if v594 != v505 {
											v571 = v594
											v573 = v573 + v591
											v585 = v590
											continue
										} else {
											break
										}
										break
									}
									v1181 = v590
								}
							}
						}
						v1202 = v497<<(uint(int32(3))%32) - base.I32_wrap_i64(v1181)
					} else {
						if base.Ui32(v493) < base.Ui32(int32(36)) {
							v1202 = int32(0)
						} else {
							v600 = v47 + int32(8)
							if v495 == int32(9) {
								v603 = int32(0)
								v650 = v603
								v651 = v603
							} else {
								v607 = int32(0)
								v610 = v607
								v611 = v607
								v614 = v607
								for {
									v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610+v600))))
									v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610+v52))))
									v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626^v628)+uint32(_consts[1104]))))
									v635 = v610 | int32(1)
									v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600+v635))))
									v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635+v52))))
									v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637^v639)+uint32(_consts[1104]))))
									v644 = v611 + v632 + v643
									v645 = int32(2)
									v646 = v610 + v645
									v648 = v614 + v645
									if v648 != v497&int32(-2) {
										v610 = v646
										v611 = v644
										v614 = v648
										continue
									} else {
										break
									}
									break
								}
								v650 = v646
								v651 = v644
							}
							if v493&int32(4) == int32(0) {
								v1202 = v651
							} else {
								v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650+v600))))
								v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650+v52))))
								v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670^v672)+uint32(_consts[1104]))))
								v1202 = v651 + v676
							}
						}
					}
				}
				*(*float32)(unsafe.Add(mBase, uint32(v16))) = base.F32_convert_i32_s(v1202)
				return v16
			}
		}
	} else {
		v46 = int32(124)
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(0)
		v52 = v48 + int32(8)
		v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
		if v53&int32(1) != 0 {
			v56 = F_palloc(m, v46)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
				v62 = int32(base.Ui32(v58)>>(uint(int32(2))%32)) - int32(8)
				if v56&int32(3) != 0 {
					v82 = v46
					v86 = F__emscripten_memset_bulkmem(m, v56, base.I32_extend8_s(int32(0)), v82)
					mBase = m.M
				} else {
					if base.Ui32(int32(1024)) < base.Ui32(v46) {
						v82 = v46
						v86 = F__emscripten_memset_bulkmem(m, v56, base.I32_extend8_s(int32(0)), v82)
						mBase = m.M
					} else {
						if v46&int32(3) != 0 {
							v82 = v46
							v86 = F__emscripten_memset_bulkmem(m, v56, base.I32_extend8_s(int32(0)), v82)
							mBase = m.M
						} else {
							v69 = v56 + v46
							if base.Ui32(v69) <= base.Ui32(v56) {
							} else {
								v74 = v56 + int32(4)
								if base.Ui32(v74) < base.Ui32(v69) {
									v76 = v69
								} else {
									v76 = v74
								}
								v82 = (v56^int32(-1)+v76)&int32(-4) + int32(4)
								v86 = F__emscripten_memset_bulkmem(m, v56, base.I32_extend8_s(int32(0)), v82)
								mBase = m.M
							}
						}
					}
				}
				if base.Ui32(v62) < base.Ui32(int32(4)) {
				} else {
					v92 = v47 + int32(8)
					v94 = v46 << (uint(int32(3)) % 32)
					v95 = int32(1)
					v97 = int32(base.Ui32(v62) >> (uint(int32(2)) % 32))
					if base.Ui32(v97) <= base.Ui32(v95) {
						v100 = v95
					} else {
						v100 = v97
					}
					v103 = int32(0)
					if base.Ui32(int32(8)) <= base.Ui32(v62) {
						v110 = v103
						v113 = int32(0)
						for {
							v124 = int32(2)
							v126 = v92 + v110<<(uint(v124)%32)
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
							v128 = base.I32_rem_u_s(v127, v94)
							v129 = int32(3)
							v131 = v56 + int32(base.Ui32(v128)>>(uint(v129)%32))
							v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
							v133 = int32(1)
							v134 = int32(7)
							v137 = v132 | v133<<(uint(v128&v134)%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v137)
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
							v140 = base.I32_rem_u_s(v139, v94)
							v143 = v56 + int32(base.Ui32(v140)>>(uint(v129)%32))
							v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
							v149 = v144 | v133<<(uint(v140&v134)%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v149)
							v152 = v110 + v124
							v154 = v113 + v124
							if v154 != v100&int32(1073741822) {
								v110 = v152
								v113 = v154
								continue
							} else {
								break
							}
							break
						}
						v157 = v152
					} else {
						v157 = v103
					}
					if v100&int32(1) == int32(0) {
					} else {
						v176 = *(*int32)(unsafe.Add(mBase, uint32(v92+v157<<(uint(int32(2))%32))))
						v177 = base.I32_rem_u_s(v176, v94)
						v180 = v56 + int32(base.Ui32(v177)>>(uint(int32(3))%32))
						v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
						v186 = v181 | int32(1)<<(uint(v177&int32(7))%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v186)
					}
				}
				v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
				if v203&int32(4) != 0 {
					v206 = int32(3)
					v207 = v46 << (uint(v206) % 32)
					if v206 < v46 {
						v682 = int64(0)
						if v46 < int32(4) {
							v761 = v56
							v762 = v46
							v767 = v682
						} else {
							if v56 != (v56+int32(3))&int32(-4) {
								v761 = v56
								v762 = v46
								v767 = v682
							} else {
								v691 = v46 - int32(4)
								v695 = int32(base.Ui32(v691)>>(uint(int32(2))%32)) + int32(1)
								v697 = v695 & int32(3)
								if base.Ui32(v691) < base.Ui32(int32(12)) {
									v733 = v56
									v734 = v46
									v739 = v682
								} else {
									v703 = v56
									v704 = v46
									v705 = int32(0)
									v709 = v682
									for {
										v710 = *(*int32)(unsafe.Add(mBase, uint32(v703)+12))
										v713 = *(*int32)(unsafe.Add(mBase, uint32(v703)+8))
										v716 = *(*int32)(unsafe.Add(mBase, uint32(v703)+4))
										v719 = *(*int32)(unsafe.Add(mBase, uint32(v703)))
										v725 = base.I64_extend_i32_u(base.I32_popcnt(v710)) + (base.I64_extend_i32_u(base.I32_popcnt(v713)) + (base.I64_extend_i32_u(base.I32_popcnt(v716)) + (v709 + base.I64_extend_i32_u(base.I32_popcnt(v719)))))
										v726 = int32(16)
										v727 = v704 - v726
										v729 = v703 + v726
										v731 = v705 + int32(4)
										if v731 != v695&int32(2147483644) {
											v703 = v729
											v704 = v727
											v705 = v731
											v709 = v725
											continue
										} else {
											break
										}
										break
									}
									v733 = v729
									v734 = v727
									v739 = v725
								}
								if v697 == int32(0) {
									v761 = v733
									v762 = v734
									v767 = v739
								} else {
									v744 = v734
									v745 = v733
									v746 = int32(0)
									v749 = v739
									for {
										v750 = int32(4)
										v751 = v744 - v750
										v752 = *(*int32)(unsafe.Add(mBase, uint32(v745)))
										v755 = v749 + base.I64_extend_i32_u(base.I32_popcnt(v752))
										v757 = v745 + v750
										v759 = v746 + int32(1)
										if v759 != v697 {
											v744 = v751
											v745 = v757
											v746 = v759
											v749 = v755
											continue
										} else {
											break
										}
										break
									}
									v761 = v757
									v762 = v751
									v767 = v755
								}
							}
						}
						if v762 == int32(0) {
							v840 = v767
						} else {
							v771 = v762 & int32(3)
							if v771 == int32(0) {
								v794 = v761
								v796 = v762
								v800 = v767
							} else {
								v777 = v762
								v778 = v761
								v779 = int32(0)
								v781 = v767
								for {
									v782 = int32(1)
									v783 = v777 - v782
									v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778))))
									v787 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v784)+uint32(_consts[1104]))))
									v788 = v781 + v787
									v790 = v778 + v782
									v792 = v779 + v782
									if v792 != v771 {
										v777 = v783
										v778 = v790
										v779 = v792
										v781 = v788
										continue
									} else {
										break
									}
									break
								}
								v794 = v790
								v796 = v783
								v800 = v788
							}
							if base.Ui32(v762) < base.Ui32(int32(4)) {
								v840 = v800
							} else {
								v803 = v794
								v805 = v796
								v809 = v800
								for {
									v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+3)))
									v813 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v810)+uint32(_consts[1104]))))
									v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+2)))
									v817 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v814)+uint32(_consts[1104]))))
									v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+1)))
									v821 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v818)+uint32(_consts[1104]))))
									v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803))))
									v825 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v822)+uint32(_consts[1104]))))
									v829 = v813 + (v817 + (v821 + (v809 + v825)))
									v830 = int32(4)
									v833 = v805 - v830
									if v833 != 0 {
										v803 = v803 + v830
										v805 = v833
										v809 = v829
										continue
									} else {
										break
									}
									break
								}
								v840 = v829
							}
						}
						v1233 = v840
					} else {
						if v46 == int32(0) {
							v1233 = v15
						} else {
							v213 = v46 & int32(3)
							if base.Ui32(v46) < base.Ui32(int32(4)) {
								v261 = v56
								v273 = v15
							} else {
								v220 = int32(0)
								v221 = v56
								v233 = v15
								for {
									v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+3)))
									v237 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v234)+uint32(_consts[1104]))))
									v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+2)))
									v241 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v238)+uint32(_consts[1104]))))
									v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
									v245 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v242)+uint32(_consts[1104]))))
									v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
									v249 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v246)+uint32(_consts[1104]))))
									v253 = v237 + (v241 + (v245 + (v233 + v249)))
									v254 = int32(4)
									v255 = v221 + v254
									v257 = v220 + v254
									if v257 != v46&int32(-4) {
										v220 = v257
										v221 = v255
										v233 = v253
										continue
									} else {
										break
									}
									break
								}
								v261 = v255
								v273 = v253
							}
							if v213 == int32(0) {
								v1233 = v273
							} else {
								v278 = int32(0)
								v279 = v261
								v291 = v273
								for {
									v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
									v295 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v292)+uint32(_consts[1104]))))
									v296 = v291 + v295
									v297 = int32(1)
									v300 = v278 + v297
									if v300 != v213 {
										v278 = v300
										v279 = v279 + v297
										v291 = v296
										continue
									} else {
										break
									}
									break
								}
								v1233 = v296
							}
						}
					}
					v1256 = base.F32_div(base.F32_convert_i32_s(v207-base.I32_wrap_i64(v1233)), base.F32_convert_i32_s(v207|int32(1)))
				} else {
					if v46 <= int32(0) {
						v1256 = float32(0)
					} else {
						v305 = int32(1)
						if v46 == v305 {
							v309 = int32(0)
							v357 = v309
							v359 = v309
						} else {
							v313 = int32(0)
							v317 = v313
							v319 = v313
							v320 = v313
							for {
								v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317+v52))))
								v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v317))))
								v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332^v334)+uint32(_consts[1104]))))
								v341 = v317 | int32(1)
								v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v341))))
								v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v341))))
								v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343^v345)+uint32(_consts[1104]))))
								v350 = v319 + v338 + v349
								v351 = int32(2)
								v352 = v317 + v351
								v354 = v320 + v351
								if v354 != v46&int32(2147483646) {
									v317 = v352
									v319 = v350
									v320 = v354
									continue
								} else {
									break
								}
								break
							}
							v357 = v352
							v359 = v350
						}
						if v46&v305 != 0 {
							v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357+v52))))
							v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v357))))
							v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372^v374)+uint32(_consts[1104]))))
							v380 = v359 + v378
						} else {
							v380 = v359
						}
						v1256 = base.F32_convert_i32_s(v380)
					}
				}
				*(*float32)(unsafe.Add(mBase, uint32(v16))) = v1256
				F_pfree(m, v56)
				mBase = m.M
				v1259 = m.ExcPending
				if v1259 != 0 {
					return int32(0)
				} else {
					return v16
				}
			}
		} else {
			v382 = int32(4)
			v383 = v53 & v382
			v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
			if v384&v382 != 0 {
				if v383 != 0 {
					v1202 = int32(0)
				} else {
					v388 = int32(8)
					v389 = v47 + v388
					v390 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
					v392 = int32(base.Ui32(v390) >> (uint(int32(2)) % 32))
					v394 = v392 - v388
					if base.Ui32(int32(47)) < base.Ui32(v390) {
						v845 = int64(0)
						if v394 < int32(4) {
							v924 = v389
							v925 = v394
							v930 = v845
						} else {
							if v389 != (v47+int32(11))&int32(-4) {
								v924 = v389
								v925 = v394
								v930 = v845
							} else {
								v854 = v394 - int32(4)
								v858 = int32(base.Ui32(v854)>>(uint(int32(2))%32)) + int32(1)
								v860 = v858 & int32(3)
								if base.Ui32(v854) < base.Ui32(int32(12)) {
									v896 = v389
									v897 = v394
									v902 = v845
								} else {
									v866 = v389
									v867 = v394
									v868 = int32(0)
									v872 = v845
									for {
										v873 = *(*int32)(unsafe.Add(mBase, uint32(v866)+12))
										v876 = *(*int32)(unsafe.Add(mBase, uint32(v866)+8))
										v879 = *(*int32)(unsafe.Add(mBase, uint32(v866)+4))
										v882 = *(*int32)(unsafe.Add(mBase, uint32(v866)))
										v888 = base.I64_extend_i32_u(base.I32_popcnt(v873)) + (base.I64_extend_i32_u(base.I32_popcnt(v876)) + (base.I64_extend_i32_u(base.I32_popcnt(v879)) + (v872 + base.I64_extend_i32_u(base.I32_popcnt(v882)))))
										v889 = int32(16)
										v890 = v867 - v889
										v892 = v866 + v889
										v894 = v868 + int32(4)
										if v894 != v858&int32(2147483644) {
											v866 = v892
											v867 = v890
											v868 = v894
											v872 = v888
											continue
										} else {
											break
										}
										break
									}
									v896 = v892
									v897 = v890
									v902 = v888
								}
								if v860 == int32(0) {
									v924 = v896
									v925 = v897
									v930 = v902
								} else {
									v907 = v897
									v908 = v896
									v909 = int32(0)
									v912 = v902
									for {
										v913 = int32(4)
										v914 = v907 - v913
										v915 = *(*int32)(unsafe.Add(mBase, uint32(v908)))
										v918 = v912 + base.I64_extend_i32_u(base.I32_popcnt(v915))
										v920 = v908 + v913
										v922 = v909 + int32(1)
										if v922 != v860 {
											v907 = v914
											v908 = v920
											v909 = v922
											v912 = v918
											continue
										} else {
											break
										}
										break
									}
									v924 = v920
									v925 = v914
									v930 = v918
								}
							}
						}
						if v925 == int32(0) {
							v1003 = v930
						} else {
							v934 = v925 & int32(3)
							if v934 == int32(0) {
								v957 = v924
								v959 = v925
								v963 = v930
							} else {
								v940 = v925
								v941 = v924
								v942 = int32(0)
								v944 = v930
								for {
									v945 = int32(1)
									v946 = v940 - v945
									v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v941))))
									v950 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v947)+uint32(_consts[1104]))))
									v951 = v944 + v950
									v953 = v941 + v945
									v955 = v942 + v945
									if v955 != v934 {
										v940 = v946
										v941 = v953
										v942 = v955
										v944 = v951
										continue
									} else {
										break
									}
									break
								}
								v957 = v953
								v959 = v946
								v963 = v951
							}
							if base.Ui32(v925) < base.Ui32(int32(4)) {
								v1003 = v963
							} else {
								v966 = v957
								v968 = v959
								v972 = v963
								for {
									v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v966)+3)))
									v976 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v973)+uint32(_consts[1104]))))
									v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v966)+2)))
									v980 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v977)+uint32(_consts[1104]))))
									v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v966)+1)))
									v984 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v981)+uint32(_consts[1104]))))
									v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v966))))
									v988 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v985)+uint32(_consts[1104]))))
									v992 = v976 + (v980 + (v984 + (v972 + v988)))
									v993 = int32(4)
									v996 = v968 - v993
									if v996 != 0 {
										v966 = v966 + v993
										v968 = v996
										v972 = v992
										continue
									} else {
										break
									}
									break
								}
								v1003 = v992
							}
						}
						v1198 = v1003
					} else {
						if v394 == int32(0) {
							v1198 = v15
						} else {
							v401 = int32(3)
							v402 = v392 & v401
							if base.Ui32(v392-int32(9)) < base.Ui32(v401) {
								v450 = v389
								v464 = v15
							} else {
								v410 = v389
								v412 = int32(0)
								v424 = v15
								for {
									v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+3)))
									v428 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v425)+uint32(_consts[1104]))))
									v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+2)))
									v432 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v429)+uint32(_consts[1104]))))
									v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+1)))
									v436 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v433)+uint32(_consts[1104]))))
									v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
									v440 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v437)+uint32(_consts[1104]))))
									v444 = v428 + (v432 + (v436 + (v424 + v440)))
									v445 = int32(4)
									v446 = v410 + v445
									v448 = v412 + v445
									if v448 != v394&int32(-4) {
										v410 = v446
										v412 = v448
										v424 = v444
										continue
									} else {
										break
									}
									break
								}
								v450 = v446
								v464 = v444
							}
							if v402 == int32(0) {
								v1198 = v464
							} else {
								v468 = v450
								v470 = int32(0)
								v482 = v464
								for {
									v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
									v486 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v483)+uint32(_consts[1104]))))
									v487 = v482 + v486
									v488 = int32(1)
									v491 = v470 + v488
									if v491 != v402 {
										v468 = v468 + v488
										v470 = v491
										v482 = v487
										continue
									} else {
										break
									}
									break
								}
								v1198 = v487
							}
						}
					}
					v1202 = v394<<(uint(int32(3))%32) - base.I32_wrap_i64(v1198)
				}
			} else {
				v493 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
				v495 = int32(base.Ui32(v493) >> (uint(int32(2)) % 32))
				v497 = v495 - int32(8)
				if v383 != 0 {
					if base.Ui32(int32(47)) < base.Ui32(v493) {
						v1008 = int64(0)
						if v497 < int32(4) {
							v1087 = v52
							v1088 = v497
							v1093 = v1008
						} else {
							if v52 != (v48+int32(11))&int32(-4) {
								v1087 = v52
								v1088 = v497
								v1093 = v1008
							} else {
								v1017 = v497 - int32(4)
								v1021 = int32(base.Ui32(v1017)>>(uint(int32(2))%32)) + int32(1)
								v1023 = v1021 & int32(3)
								if base.Ui32(v1017) < base.Ui32(int32(12)) {
									v1059 = v52
									v1060 = v497
									v1065 = v1008
								} else {
									v1029 = v52
									v1030 = v497
									v1031 = int32(0)
									v1035 = v1008
									for {
										v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+12))
										v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+8))
										v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+4))
										v1045 = *(*int32)(unsafe.Add(mBase, uint32(v1029)))
										v1051 = base.I64_extend_i32_u(base.I32_popcnt(v1036)) + (base.I64_extend_i32_u(base.I32_popcnt(v1039)) + (base.I64_extend_i32_u(base.I32_popcnt(v1042)) + (v1035 + base.I64_extend_i32_u(base.I32_popcnt(v1045)))))
										v1052 = int32(16)
										v1053 = v1030 - v1052
										v1055 = v1029 + v1052
										v1057 = v1031 + int32(4)
										if v1057 != v1021&int32(2147483644) {
											v1029 = v1055
											v1030 = v1053
											v1031 = v1057
											v1035 = v1051
											continue
										} else {
											break
										}
										break
									}
									v1059 = v1055
									v1060 = v1053
									v1065 = v1051
								}
								if v1023 == int32(0) {
									v1087 = v1059
									v1088 = v1060
									v1093 = v1065
								} else {
									v1070 = v1060
									v1071 = v1059
									v1072 = int32(0)
									v1075 = v1065
									for {
										v1076 = int32(4)
										v1077 = v1070 - v1076
										v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1071)))
										v1081 = v1075 + base.I64_extend_i32_u(base.I32_popcnt(v1078))
										v1083 = v1071 + v1076
										v1085 = v1072 + int32(1)
										if v1085 != v1023 {
											v1070 = v1077
											v1071 = v1083
											v1072 = v1085
											v1075 = v1081
											continue
										} else {
											break
										}
										break
									}
									v1087 = v1083
									v1088 = v1077
									v1093 = v1081
								}
							}
						}
						if v1088 == int32(0) {
							v1166 = v1093
						} else {
							v1097 = v1088 & int32(3)
							if v1097 == int32(0) {
								v1120 = v1087
								v1122 = v1088
								v1126 = v1093
							} else {
								v1103 = v1088
								v1104 = v1087
								v1105 = int32(0)
								v1107 = v1093
								for {
									v1108 = int32(1)
									v1109 = v1103 - v1108
									v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1104))))
									v1113 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1110)+uint32(_consts[1104]))))
									v1114 = v1107 + v1113
									v1116 = v1104 + v1108
									v1118 = v1105 + v1108
									if v1118 != v1097 {
										v1103 = v1109
										v1104 = v1116
										v1105 = v1118
										v1107 = v1114
										continue
									} else {
										break
									}
									break
								}
								v1120 = v1116
								v1122 = v1109
								v1126 = v1114
							}
							if base.Ui32(v1088) < base.Ui32(int32(4)) {
								v1166 = v1126
							} else {
								v1129 = v1120
								v1131 = v1122
								v1135 = v1126
								for {
									v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129)+3)))
									v1139 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1136)+uint32(_consts[1104]))))
									v1140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129)+2)))
									v1143 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1140)+uint32(_consts[1104]))))
									v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129)+1)))
									v1147 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+uint32(_consts[1104]))))
									v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129))))
									v1151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1148)+uint32(_consts[1104]))))
									v1155 = v1139 + (v1143 + (v1147 + (v1135 + v1151)))
									v1156 = int32(4)
									v1159 = v1131 - v1156
									if v1159 != 0 {
										v1129 = v1129 + v1156
										v1131 = v1159
										v1135 = v1155
										continue
									} else {
										break
									}
									break
								}
								v1166 = v1155
							}
						}
						v1181 = v1166
					} else {
						if v497 == int32(0) {
							v1181 = v15
						} else {
							v504 = int32(3)
							v505 = v495 & v504
							if base.Ui32(v495-int32(9)) < base.Ui32(v504) {
								v555 = v52
								v567 = v15
							} else {
								v513 = int32(0)
								v515 = v52
								v527 = v15
								for {
									v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+3)))
									v531 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v528)+uint32(_consts[1104]))))
									v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+2)))
									v535 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v532)+uint32(_consts[1104]))))
									v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
									v539 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v536)+uint32(_consts[1104]))))
									v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
									v543 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v540)+uint32(_consts[1104]))))
									v547 = v531 + (v535 + (v539 + (v527 + v543)))
									v548 = int32(4)
									v549 = v515 + v548
									v551 = v513 + v548
									if v551 != v497&int32(-4) {
										v513 = v551
										v515 = v549
										v527 = v547
										continue
									} else {
										break
									}
									break
								}
								v555 = v549
								v567 = v547
							}
							if v505 == int32(0) {
								v1181 = v567
							} else {
								v571 = int32(0)
								v573 = v555
								v585 = v567
								for {
									v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573))))
									v589 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v586)+uint32(_consts[1104]))))
									v590 = v585 + v589
									v591 = int32(1)
									v594 = v571 + v591
									if v594 != v505 {
										v571 = v594
										v573 = v573 + v591
										v585 = v590
										continue
									} else {
										break
									}
									break
								}
								v1181 = v590
							}
						}
					}
					v1202 = v497<<(uint(int32(3))%32) - base.I32_wrap_i64(v1181)
				} else {
					if base.Ui32(v493) < base.Ui32(int32(36)) {
						v1202 = int32(0)
					} else {
						v600 = v47 + int32(8)
						if v495 == int32(9) {
							v603 = int32(0)
							v650 = v603
							v651 = v603
						} else {
							v607 = int32(0)
							v610 = v607
							v611 = v607
							v614 = v607
							for {
								v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610+v600))))
								v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610+v52))))
								v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626^v628)+uint32(_consts[1104]))))
								v635 = v610 | int32(1)
								v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600+v635))))
								v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635+v52))))
								v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637^v639)+uint32(_consts[1104]))))
								v644 = v611 + v632 + v643
								v645 = int32(2)
								v646 = v610 + v645
								v648 = v614 + v645
								if v648 != v497&int32(-2) {
									v610 = v646
									v611 = v644
									v614 = v648
									continue
								} else {
									break
								}
								break
							}
							v650 = v646
							v651 = v644
						}
						if v493&int32(4) == int32(0) {
							v1202 = v651
						} else {
							v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650+v600))))
							v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650+v52))))
							v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670^v672)+uint32(_consts[1104]))))
							v1202 = v651 + v676
						}
					}
				}
			}
			*(*float32)(unsafe.Add(mBase, uint32(v16))) = base.F32_convert_i32_s(v1202)
			return v16
		}
	}
}
