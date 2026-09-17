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
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
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
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v236 int64
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int64
	_ = v240
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v246 int64
	_ = v246
	var v250 int64
	_ = v250
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v270 int64
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v286 int64
	_ = v286
	var v287 int32
	_ = v287
	var v288 int64
	_ = v288
	var v289 int64
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
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
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v425 int64
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int64
	_ = v429
	var v430 int32
	_ = v430
	var v431 int64
	_ = v431
	var v432 int32
	_ = v432
	var v433 int64
	_ = v433
	var v434 int32
	_ = v434
	var v435 int64
	_ = v435
	var v439 int64
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v459 int64
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v475 int64
	_ = v475
	var v476 int32
	_ = v476
	var v477 int64
	_ = v477
	var v478 int64
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v518 int64
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int64
	_ = v522
	var v523 int32
	_ = v523
	var v524 int64
	_ = v524
	var v525 int32
	_ = v525
	var v526 int64
	_ = v526
	var v527 int32
	_ = v527
	var v528 int64
	_ = v528
	var v532 int64
	_ = v532
	var v534 int32
	_ = v534
	var v542 int32
	_ = v542
	var v552 int64
	_ = v552
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v568 int64
	_ = v568
	var v569 int32
	_ = v569
	var v570 int64
	_ = v570
	var v571 int64
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v654 int64
	_ = v654
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v681 int64
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v701 int64
	_ = v701
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v713 int64
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v721 int64
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int64
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v739 int64
	_ = v739
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int64
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int64
	_ = v759
	var v760 int64
	_ = v760
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v770 int64
	_ = v770
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v779 int64
	_ = v779
	var v780 int32
	_ = v780
	var v781 int64
	_ = v781
	var v782 int32
	_ = v782
	var v783 int64
	_ = v783
	var v784 int32
	_ = v784
	var v785 int64
	_ = v785
	var v786 int32
	_ = v786
	var v787 int64
	_ = v787
	var v791 int64
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v802 int64
	_ = v802
	var v807 int64
	_ = v807
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v834 int64
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v854 int64
	_ = v854
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v866 int64
	_ = v866
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v874 int64
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v882 int64
	_ = v882
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v892 int64
	_ = v892
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v906 int64
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int64
	_ = v912
	var v913 int64
	_ = v913
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v923 int64
	_ = v923
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v932 int64
	_ = v932
	var v933 int32
	_ = v933
	var v934 int64
	_ = v934
	var v935 int32
	_ = v935
	var v936 int64
	_ = v936
	var v937 int32
	_ = v937
	var v938 int64
	_ = v938
	var v939 int32
	_ = v939
	var v940 int64
	_ = v940
	var v944 int64
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v955 int64
	_ = v955
	var v960 int64
	_ = v960
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v987 int64
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1007 int64
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1019 int64
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1027 int64
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1035 int64
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1045 int64
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1059 int64
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int64
	_ = v1065
	var v1066 int64
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1076 int64
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1085 int64
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int64
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int64
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int64
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int64
	_ = v1093
	var v1097 int64
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1108 int64
	_ = v1108
	var v1123 int64
	_ = v1123
	var v1140 int64
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1175 int64
	_ = v1175
	var v1198 float32
	_ = v1198
	var v1201 int32
	_ = v1201
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
					v59 = int32(2)
					v64 = int32(base.Ui32(int32(base.Ui32(v58)>>(uint(v59)%32))-int32(8)) >> (uint(v59) % 32))
					v65 = int32(3)
					if v46&v65|(v56&v65|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v46))) == int32(0) {
						if v46 == int32(0) {
						} else {
							v79 = v46 + v56
							v81 = v56 + int32(4)
							if base.Ui32(v81) < base.Ui32(v79) {
								v83 = v79
							} else {
								v83 = v81
							}
							v91 = (v56^int32(-1)+v83)&int32(-4) + int32(4)
							if v91 == int32(0) {
							} else {
								base.MemoryFill(m, v56, int32(0), v91)
							}
						}
					} else {
						v91 = v46
						if v91 == int32(0) {
						} else {
							base.MemoryFill(m, v56, int32(0), v91)
						}
					}
					if v64 == int32(0) {
					} else {
						v102 = v47 + int32(8)
						v104 = v46 << (uint(int32(3)) % 32)
						v105 = int32(0)
						if v64 != int32(1) {
							v113 = v105
							v121 = v2
							for {
								v127 = int32(2)
								v129 = v102 + v113<<(uint(v127)%32)
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
								v131 = base.I32_rem_u_s(v130, v104)
								v132 = int32(3)
								v134 = v56 + int32(base.Ui32(v131)>>(uint(v132)%32))
								v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
								v136 = int32(1)
								v137 = int32(7)
								v140 = v135 | v136<<(uint(v131&v137)%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v140)
								v142 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
								v143 = base.I32_rem_u_s(v142, v104)
								v146 = v56 + int32(base.Ui32(v143)>>(uint(v132)%32))
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
								v152 = v147 | v136<<(uint(v143&v137)%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v146))) = uint8(v152)
								v155 = v113 + v127
								v157 = v121 + v127
								if v157 != v64&int32(1073741822) {
									v113 = v155
									v121 = v157
									continue
								} else {
									break
								}
								break
							}
							if v64&int32(1) == int32(0) {
							} else {
								v162 = v155
								v179 = *(*int32)(unsafe.Add(mBase, uint32(v102+v162<<(uint(int32(2))%32))))
								v180 = base.I32_rem_u_s(v179, v104)
								v183 = v56 + int32(base.Ui32(v180)>>(uint(int32(3))%32))
								v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
								v189 = v184 | int32(1)<<(uint(v180&int32(7))%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v189)
							}
						} else {
							v162 = v105
							v179 = *(*int32)(unsafe.Add(mBase, uint32(v102+v162<<(uint(int32(2))%32))))
							v180 = base.I32_rem_u_s(v179, v104)
							v183 = v56 + int32(base.Ui32(v180)>>(uint(int32(3))%32))
							v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
							v189 = v184 | int32(1)<<(uint(v180&int32(7))%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v189)
						}
					}
					v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
					if v206&int32(4) != 0 {
						v209 = int32(3)
						v210 = v46 << (uint(v209) % 32)
						if v209 < v46 {
							v654 = int64(0)
							if base.B2i32(v56 != (v56+int32(3))&int32(-4))|base.B2i32(v46 < int32(4)) != 0 {
								v733 = v56
								v734 = v46
								v739 = v654
							} else {
								v664 = v46 - int32(4)
								v668 = int32(base.Ui32(v664)>>(uint(int32(2))%32)) + int32(1)
								v670 = v668 & int32(3)
								if base.Ui32(int32(12)) <= base.Ui32(v664) {
									v675 = v56
									v676 = v46
									v679 = int32(0)
									v681 = v654
									for {
										v682 = int32(16)
										v683 = v676 - v682
										v685 = v675 + v682
										v686 = *(*int32)(unsafe.Add(mBase, uint32(v675)+12))
										v689 = *(*int32)(unsafe.Add(mBase, uint32(v675)+8))
										v692 = *(*int32)(unsafe.Add(mBase, uint32(v675)+4))
										v695 = *(*int32)(unsafe.Add(mBase, uint32(v675)))
										v701 = base.I64_extend_i32_u(base.I32_popcnt(v686)) + (base.I64_extend_i32_u(base.I32_popcnt(v689)) + (base.I64_extend_i32_u(base.I32_popcnt(v692)) + (v681 + base.I64_extend_i32_u(base.I32_popcnt(v695)))))
										v703 = v679 + int32(4)
										if v703 != v668&int32(2147483644) {
											v675 = v685
											v676 = v683
											v679 = v703
											v681 = v701
											continue
										} else {
											break
										}
										break
									}
									if v670 == int32(0) {
										v733 = v685
										v734 = v683
										v739 = v701
									} else {
										v707 = v685
										v708 = v683
										v713 = v701
										v715 = v707
										v716 = v708
										v717 = int32(0)
										v721 = v713
										for {
											v722 = int32(4)
											v723 = v716 - v722
											v725 = v715 + v722
											v726 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
											v729 = v721 + base.I64_extend_i32_u(base.I32_popcnt(v726))
											v731 = v717 + int32(1)
											if v731 != v670 {
												v715 = v725
												v716 = v723
												v717 = v731
												v721 = v729
												continue
											} else {
												break
											}
											break
										}
										v733 = v725
										v734 = v723
										v739 = v729
									}
								} else {
									v707 = v56
									v708 = v46
									v713 = v654
									v715 = v707
									v716 = v708
									v717 = int32(0)
									v721 = v713
									for {
										v722 = int32(4)
										v723 = v716 - v722
										v725 = v715 + v722
										v726 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
										v729 = v721 + base.I64_extend_i32_u(base.I32_popcnt(v726))
										v731 = v717 + int32(1)
										if v731 != v670 {
											v715 = v725
											v716 = v723
											v717 = v731
											v721 = v729
											continue
										} else {
											break
										}
										break
									}
									v733 = v725
									v734 = v723
									v739 = v729
								}
							}
							if v734 == int32(0) {
								v802 = v739
							} else {
								v743 = v734 & int32(3)
								if v743 == int32(0) {
									v764 = v733
									v766 = v734
									v770 = v739
								} else {
									v747 = v733
									v749 = v734
									v751 = int32(0)
									v753 = v739
									for {
										v754 = int32(1)
										v755 = v747 + v754
										v757 = v749 - v754
										v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747))))
										v759 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v758)+uint32(_c_F_gtsvector_penalty[0]))))
										v760 = v753 + v759
										v762 = v751 + v754
										if v762 != v743 {
											v747 = v755
											v749 = v757
											v751 = v762
											v753 = v760
											continue
										} else {
											break
										}
										break
									}
									v764 = v755
									v766 = v757
									v770 = v760
								}
								if base.Ui32(v734) < base.Ui32(int32(4)) {
									v802 = v770
								} else {
									v773 = v764
									v775 = v766
									v779 = v770
									for {
										v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773)+3)))
										v781 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v780)+uint32(_c_F_gtsvector_penalty[0]))))
										v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773)+2)))
										v783 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v782)+uint32(_c_F_gtsvector_penalty[0]))))
										v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773)+1)))
										v785 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v784)+uint32(_c_F_gtsvector_penalty[0]))))
										v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773))))
										v787 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v786)+uint32(_c_F_gtsvector_penalty[0]))))
										v791 = v781 + (v783 + (v785 + (v779 + v787)))
										v792 = int32(4)
										v795 = v775 - v792
										if v795 != 0 {
											v773 = v773 + v792
											v775 = v795
											v779 = v791
											continue
										} else {
											break
										}
										break
									}
									v802 = v791
								}
							}
							v1175 = v802
						} else {
							if v46 == int32(0) {
								v1175 = v15
							} else {
								v216 = v46 & int32(3)
								if base.Ui32(int32(4)) <= base.Ui32(v46) {
									v223 = v56
									v224 = int32(0)
									v236 = v15
									for {
										v237 = int32(4)
										v238 = v223 + v237
										v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+3)))
										v240 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v239)+uint32(_c_F_gtsvector_penalty[0]))))
										v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+2)))
										v242 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v241)+uint32(_c_F_gtsvector_penalty[0]))))
										v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
										v244 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v243)+uint32(_c_F_gtsvector_penalty[0]))))
										v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
										v246 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v245)+uint32(_c_F_gtsvector_penalty[0]))))
										v250 = v240 + (v242 + (v244 + (v236 + v246)))
										v252 = v224 + v237
										if v252 != v46&int32(-4) {
											v223 = v238
											v224 = v252
											v236 = v250
											continue
										} else {
											break
										}
										break
									}
									if v216 == int32(0) {
										v1175 = v250
									} else {
										v257 = v238
										v270 = v250
										v273 = v257
										v276 = int32(0)
										v286 = v270
										for {
											v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
											v288 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v287)+uint32(_c_F_gtsvector_penalty[0]))))
											v289 = v286 + v288
											v290 = int32(1)
											v293 = v276 + v290
											if v293 != v216 {
												v273 = v273 + v290
												v276 = v293
												v286 = v289
												continue
											} else {
												break
											}
											break
										}
										v1175 = v289
									}
								} else {
									v257 = v56
									v270 = v15
									v273 = v257
									v276 = int32(0)
									v286 = v270
									for {
										v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
										v288 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v287)+uint32(_c_F_gtsvector_penalty[0]))))
										v289 = v286 + v288
										v290 = int32(1)
										v293 = v276 + v290
										if v293 != v216 {
											v273 = v273 + v290
											v276 = v293
											v286 = v289
											continue
										} else {
											break
										}
										break
									}
									v1175 = v289
								}
							}
						}
						v1198 = base.F32_div(base.F32_convert_i32_s(v210-base.I32_wrap_i64(v1175)), base.F32_convert_i32_s(v210|int32(1)))
					} else {
						if v46 <= int32(0) {
							v1198 = float32(0)
						} else {
							v298 = int32(0)
							if v46 != int32(1) {
								v308 = v298
								v309 = v298
								v316 = int32(0)
								for {
									v323 = v308 | int32(1)
									v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v323))))
									v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323+v56))))
									v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325^v327)+uint32(_c_F_gtsvector_penalty[0]))))
									v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308+v52))))
									v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308+v56))))
									v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331^v333)+uint32(_c_F_gtsvector_penalty[0]))))
									v337 = v329 + (v309 + v335)
									v338 = int32(2)
									v339 = v308 + v338
									v341 = v316 + v338
									if v341 != v46&int32(2147483646) {
										v308 = v339
										v309 = v337
										v316 = v341
										continue
									} else {
										break
									}
									break
								}
								if v46&int32(1) == int32(0) {
									v369 = v337
								} else {
									v346 = v339
									v347 = v337
									v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346+v52))))
									v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346+v56))))
									v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361^v363)+uint32(_c_F_gtsvector_penalty[0]))))
									v369 = v347 + v365
								}
							} else {
								v346 = v298
								v347 = v298
								v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346+v52))))
								v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346+v56))))
								v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361^v363)+uint32(_c_F_gtsvector_penalty[0]))))
								v369 = v347 + v365
							}
							v1198 = base.F32_convert_i32_s(v369)
						}
					}
					*(*float32)(unsafe.Add(mBase, uint32(v16))) = v1198
					F_pfree(m, v56)
					mBase = m.M
					v1201 = m.ExcPending
					if v1201 != 0 {
						return int32(0)
					} else {
						return v16
					}
				}
			} else {
				v383 = int32(4)
				v384 = v53 & v383
				v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
				if v385&v383 != 0 {
					if v384 != 0 {
						v1144 = int32(0)
					} else {
						v389 = int32(8)
						v390 = v47 + v389
						v391 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
						v393 = int32(base.Ui32(v391) >> (uint(int32(2)) % 32))
						v395 = v393 - v389
						if base.Ui32(int32(47)) < base.Ui32(v391) {
							v807 = int64(0)
							if base.B2i32(v390 != (v47+int32(11))&int32(-4))|base.B2i32(v395 < int32(4)) != 0 {
								v886 = v390
								v887 = v395
								v892 = v807
							} else {
								v817 = v395 - int32(4)
								v821 = int32(base.Ui32(v817)>>(uint(int32(2))%32)) + int32(1)
								v823 = v821 & int32(3)
								if base.Ui32(int32(12)) <= base.Ui32(v817) {
									v828 = v390
									v829 = v395
									v832 = int32(0)
									v834 = v807
									for {
										v835 = int32(16)
										v836 = v829 - v835
										v838 = v828 + v835
										v839 = *(*int32)(unsafe.Add(mBase, uint32(v828)+12))
										v842 = *(*int32)(unsafe.Add(mBase, uint32(v828)+8))
										v845 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
										v848 = *(*int32)(unsafe.Add(mBase, uint32(v828)))
										v854 = base.I64_extend_i32_u(base.I32_popcnt(v839)) + (base.I64_extend_i32_u(base.I32_popcnt(v842)) + (base.I64_extend_i32_u(base.I32_popcnt(v845)) + (v834 + base.I64_extend_i32_u(base.I32_popcnt(v848)))))
										v856 = v832 + int32(4)
										if v856 != v821&int32(2147483644) {
											v828 = v838
											v829 = v836
											v832 = v856
											v834 = v854
											continue
										} else {
											break
										}
										break
									}
									if v823 == int32(0) {
										v886 = v838
										v887 = v836
										v892 = v854
									} else {
										v860 = v838
										v861 = v836
										v866 = v854
										v868 = v860
										v869 = v861
										v870 = int32(0)
										v874 = v866
										for {
											v875 = int32(4)
											v876 = v869 - v875
											v878 = v868 + v875
											v879 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
											v882 = v874 + base.I64_extend_i32_u(base.I32_popcnt(v879))
											v884 = v870 + int32(1)
											if v884 != v823 {
												v868 = v878
												v869 = v876
												v870 = v884
												v874 = v882
												continue
											} else {
												break
											}
											break
										}
										v886 = v878
										v887 = v876
										v892 = v882
									}
								} else {
									v860 = v390
									v861 = v395
									v866 = v807
									v868 = v860
									v869 = v861
									v870 = int32(0)
									v874 = v866
									for {
										v875 = int32(4)
										v876 = v869 - v875
										v878 = v868 + v875
										v879 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
										v882 = v874 + base.I64_extend_i32_u(base.I32_popcnt(v879))
										v884 = v870 + int32(1)
										if v884 != v823 {
											v868 = v878
											v869 = v876
											v870 = v884
											v874 = v882
											continue
										} else {
											break
										}
										break
									}
									v886 = v878
									v887 = v876
									v892 = v882
								}
							}
							if v887 == int32(0) {
								v955 = v892
							} else {
								v896 = v887 & int32(3)
								if v896 == int32(0) {
									v917 = v886
									v919 = v887
									v923 = v892
								} else {
									v900 = v886
									v902 = v887
									v904 = int32(0)
									v906 = v892
									for {
										v907 = int32(1)
										v908 = v900 + v907
										v910 = v902 - v907
										v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v900))))
										v912 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v911)+uint32(_c_F_gtsvector_penalty[0]))))
										v913 = v906 + v912
										v915 = v904 + v907
										if v915 != v896 {
											v900 = v908
											v902 = v910
											v904 = v915
											v906 = v913
											continue
										} else {
											break
										}
										break
									}
									v917 = v908
									v919 = v910
									v923 = v913
								}
								if base.Ui32(v887) < base.Ui32(int32(4)) {
									v955 = v923
								} else {
									v926 = v917
									v928 = v919
									v932 = v923
									for {
										v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926)+3)))
										v934 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v933)+uint32(_c_F_gtsvector_penalty[0]))))
										v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926)+2)))
										v936 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v935)+uint32(_c_F_gtsvector_penalty[0]))))
										v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926)+1)))
										v938 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v937)+uint32(_c_F_gtsvector_penalty[0]))))
										v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926))))
										v940 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v939)+uint32(_c_F_gtsvector_penalty[0]))))
										v944 = v934 + (v936 + (v938 + (v932 + v940)))
										v945 = int32(4)
										v948 = v928 - v945
										if v948 != 0 {
											v926 = v926 + v945
											v928 = v948
											v932 = v944
											continue
										} else {
											break
										}
										break
									}
									v955 = v944
								}
							}
							v1140 = v955
						} else {
							if v395 == int32(0) {
								v1140 = v15
							} else {
								v402 = int32(3)
								v403 = v393 & v402
								if base.Ui32(v402) <= base.Ui32(v393-int32(9)) {
									v411 = v390
									v415 = int32(0)
									v425 = v15
									for {
										v426 = int32(4)
										v427 = v411 + v426
										v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+3)))
										v429 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v428)+uint32(_c_F_gtsvector_penalty[0]))))
										v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+2)))
										v431 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v430)+uint32(_c_F_gtsvector_penalty[0]))))
										v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+1)))
										v433 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v432)+uint32(_c_F_gtsvector_penalty[0]))))
										v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
										v435 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v434)+uint32(_c_F_gtsvector_penalty[0]))))
										v439 = v429 + (v431 + (v433 + (v425 + v435)))
										v441 = v415 + v426
										if v441 != v395&int32(-4) {
											v411 = v427
											v415 = v441
											v425 = v439
											continue
										} else {
											break
										}
										break
									}
									if v403 == int32(0) {
										v1140 = v439
									} else {
										v445 = v427
										v459 = v439
										v461 = v445
										v462 = int32(0)
										v475 = v459
										for {
											v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
											v477 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v476)+uint32(_c_F_gtsvector_penalty[0]))))
											v478 = v475 + v477
											v479 = int32(1)
											v482 = v462 + v479
											if v482 != v403 {
												v461 = v461 + v479
												v462 = v482
												v475 = v478
												continue
											} else {
												break
											}
											break
										}
										v1140 = v478
									}
								} else {
									v445 = v390
									v459 = v15
									v461 = v445
									v462 = int32(0)
									v475 = v459
									for {
										v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
										v477 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v476)+uint32(_c_F_gtsvector_penalty[0]))))
										v478 = v475 + v477
										v479 = int32(1)
										v482 = v462 + v479
										if v482 != v403 {
											v461 = v461 + v479
											v462 = v482
											v475 = v478
											continue
										} else {
											break
										}
										break
									}
									v1140 = v478
								}
							}
						}
						v1144 = v395<<(uint(int32(3))%32) - base.I32_wrap_i64(v1140)
					}
				} else {
					v484 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
					v486 = int32(base.Ui32(v484) >> (uint(int32(2)) % 32))
					v488 = v486 - int32(8)
					if v384 != 0 {
						if base.Ui32(int32(47)) < base.Ui32(v484) {
							v960 = int64(0)
							if base.B2i32(v52 != (v48+int32(11))&int32(-4))|base.B2i32(v488 < int32(4)) != 0 {
								v1039 = v52
								v1040 = v488
								v1045 = v960
							} else {
								v970 = v488 - int32(4)
								v974 = int32(base.Ui32(v970)>>(uint(int32(2))%32)) + int32(1)
								v976 = v974 & int32(3)
								if base.Ui32(int32(12)) <= base.Ui32(v970) {
									v981 = v52
									v982 = v488
									v985 = int32(0)
									v987 = v960
									for {
										v988 = int32(16)
										v989 = v982 - v988
										v991 = v981 + v988
										v992 = *(*int32)(unsafe.Add(mBase, uint32(v981)+12))
										v995 = *(*int32)(unsafe.Add(mBase, uint32(v981)+8))
										v998 = *(*int32)(unsafe.Add(mBase, uint32(v981)+4))
										v1001 = *(*int32)(unsafe.Add(mBase, uint32(v981)))
										v1007 = base.I64_extend_i32_u(base.I32_popcnt(v992)) + (base.I64_extend_i32_u(base.I32_popcnt(v995)) + (base.I64_extend_i32_u(base.I32_popcnt(v998)) + (v987 + base.I64_extend_i32_u(base.I32_popcnt(v1001)))))
										v1009 = v985 + int32(4)
										if v1009 != v974&int32(2147483644) {
											v981 = v991
											v982 = v989
											v985 = v1009
											v987 = v1007
											continue
										} else {
											break
										}
										break
									}
									if v976 == int32(0) {
										v1039 = v991
										v1040 = v989
										v1045 = v1007
									} else {
										v1013 = v991
										v1014 = v989
										v1019 = v1007
										v1021 = v1013
										v1022 = v1014
										v1023 = int32(0)
										v1027 = v1019
										for {
											v1028 = int32(4)
											v1029 = v1022 - v1028
											v1031 = v1021 + v1028
											v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1021)))
											v1035 = v1027 + base.I64_extend_i32_u(base.I32_popcnt(v1032))
											v1037 = v1023 + int32(1)
											if v1037 != v976 {
												v1021 = v1031
												v1022 = v1029
												v1023 = v1037
												v1027 = v1035
												continue
											} else {
												break
											}
											break
										}
										v1039 = v1031
										v1040 = v1029
										v1045 = v1035
									}
								} else {
									v1013 = v52
									v1014 = v488
									v1019 = v960
									v1021 = v1013
									v1022 = v1014
									v1023 = int32(0)
									v1027 = v1019
									for {
										v1028 = int32(4)
										v1029 = v1022 - v1028
										v1031 = v1021 + v1028
										v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1021)))
										v1035 = v1027 + base.I64_extend_i32_u(base.I32_popcnt(v1032))
										v1037 = v1023 + int32(1)
										if v1037 != v976 {
											v1021 = v1031
											v1022 = v1029
											v1023 = v1037
											v1027 = v1035
											continue
										} else {
											break
										}
										break
									}
									v1039 = v1031
									v1040 = v1029
									v1045 = v1035
								}
							}
							if v1040 == int32(0) {
								v1108 = v1045
							} else {
								v1049 = v1040 & int32(3)
								if v1049 == int32(0) {
									v1070 = v1039
									v1072 = v1040
									v1076 = v1045
								} else {
									v1053 = v1039
									v1055 = v1040
									v1057 = int32(0)
									v1059 = v1045
									for {
										v1060 = int32(1)
										v1061 = v1053 + v1060
										v1063 = v1055 - v1060
										v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053))))
										v1065 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1064)+uint32(_c_F_gtsvector_penalty[0]))))
										v1066 = v1059 + v1065
										v1068 = v1057 + v1060
										if v1068 != v1049 {
											v1053 = v1061
											v1055 = v1063
											v1057 = v1068
											v1059 = v1066
											continue
										} else {
											break
										}
										break
									}
									v1070 = v1061
									v1072 = v1063
									v1076 = v1066
								}
								if base.Ui32(v1040) < base.Ui32(int32(4)) {
									v1108 = v1076
								} else {
									v1079 = v1070
									v1081 = v1072
									v1085 = v1076
									for {
										v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+3)))
										v1087 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1086)+uint32(_c_F_gtsvector_penalty[0]))))
										v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+2)))
										v1089 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1088)+uint32(_c_F_gtsvector_penalty[0]))))
										v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+1)))
										v1091 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1090)+uint32(_c_F_gtsvector_penalty[0]))))
										v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079))))
										v1093 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1092)+uint32(_c_F_gtsvector_penalty[0]))))
										v1097 = v1087 + (v1089 + (v1091 + (v1085 + v1093)))
										v1098 = int32(4)
										v1101 = v1081 - v1098
										if v1101 != 0 {
											v1079 = v1079 + v1098
											v1081 = v1101
											v1085 = v1097
											continue
										} else {
											break
										}
										break
									}
									v1108 = v1097
								}
							}
							v1123 = v1108
						} else {
							if v488 == int32(0) {
								v1123 = v15
							} else {
								v495 = int32(3)
								v496 = v486 & v495
								if base.Ui32(v495) <= base.Ui32(v486-int32(9)) {
									v505 = int32(0)
									v508 = v52
									v518 = v15
									for {
										v519 = int32(4)
										v520 = v508 + v519
										v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+3)))
										v522 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v521)+uint32(_c_F_gtsvector_penalty[0]))))
										v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+2)))
										v524 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v523)+uint32(_c_F_gtsvector_penalty[0]))))
										v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+1)))
										v526 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v525)+uint32(_c_F_gtsvector_penalty[0]))))
										v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
										v528 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v527)+uint32(_c_F_gtsvector_penalty[0]))))
										v532 = v522 + (v524 + (v526 + (v518 + v528)))
										v534 = v505 + v519
										if v534 != v488&int32(-4) {
											v505 = v534
											v508 = v520
											v518 = v532
											continue
										} else {
											break
										}
										break
									}
									if v496 == int32(0) {
										v1123 = v532
									} else {
										v542 = v520
										v552 = v532
										v554 = int32(0)
										v558 = v542
										v568 = v552
										for {
											v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
											v570 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v569)+uint32(_c_F_gtsvector_penalty[0]))))
											v571 = v568 + v570
											v572 = int32(1)
											v575 = v554 + v572
											if v575 != v496 {
												v554 = v575
												v558 = v558 + v572
												v568 = v571
												continue
											} else {
												break
											}
											break
										}
										v1123 = v571
									}
								} else {
									v542 = v52
									v552 = v15
									v554 = int32(0)
									v558 = v542
									v568 = v552
									for {
										v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
										v570 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v569)+uint32(_c_F_gtsvector_penalty[0]))))
										v571 = v568 + v570
										v572 = int32(1)
										v575 = v554 + v572
										if v575 != v496 {
											v554 = v575
											v558 = v558 + v572
											v568 = v571
											continue
										} else {
											break
										}
										break
									}
									v1123 = v571
								}
							}
						}
						v1144 = v488<<(uint(int32(3))%32) - base.I32_wrap_i64(v1123)
					} else {
						if base.Ui32(v484) < base.Ui32(int32(36)) {
							v1144 = int32(0)
						} else {
							v581 = v47 + int32(8)
							v582 = int32(0)
							if v486 != int32(9) {
								v590 = v582
								v591 = v582
								v599 = v2
								for {
									v606 = v590 | int32(1)
									v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581+v606))))
									v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v606+v52))))
									v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608^v610)+uint32(_c_F_gtsvector_penalty[0]))))
									v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590+v581))))
									v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590+v52))))
									v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614^v616)+uint32(_c_F_gtsvector_penalty[0]))))
									v620 = v612 + (v591 + v618)
									v621 = int32(2)
									v622 = v590 + v621
									v624 = v599 + v621
									if v624 != v488&int32(-2) {
										v590 = v622
										v591 = v620
										v599 = v624
										continue
									} else {
										break
									}
									break
								}
								if v486&int32(1) == int32(0) {
									v1144 = v620
								} else {
									v628 = v622
									v629 = v620
									v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628+v581))))
									v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628+v52))))
									v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644^v646)+uint32(_c_F_gtsvector_penalty[0]))))
									v1144 = v629 + v648
								}
							} else {
								v628 = v582
								v629 = v582
								v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628+v581))))
								v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628+v52))))
								v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644^v646)+uint32(_c_F_gtsvector_penalty[0]))))
								v1144 = v629 + v648
							}
						}
					}
				}
				*(*float32)(unsafe.Add(mBase, uint32(v16))) = base.F32_convert_i32_s(v1144)
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
				v59 = int32(2)
				v64 = int32(base.Ui32(int32(base.Ui32(v58)>>(uint(v59)%32))-int32(8)) >> (uint(v59) % 32))
				v65 = int32(3)
				if v46&v65|(v56&v65|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v46))) == int32(0) {
					if v46 == int32(0) {
					} else {
						v79 = v46 + v56
						v81 = v56 + int32(4)
						if base.Ui32(v81) < base.Ui32(v79) {
							v83 = v79
						} else {
							v83 = v81
						}
						v91 = (v56^int32(-1)+v83)&int32(-4) + int32(4)
						if v91 == int32(0) {
						} else {
							base.MemoryFill(m, v56, int32(0), v91)
						}
					}
				} else {
					v91 = v46
					if v91 == int32(0) {
					} else {
						base.MemoryFill(m, v56, int32(0), v91)
					}
				}
				if v64 == int32(0) {
				} else {
					v102 = v47 + int32(8)
					v104 = v46 << (uint(int32(3)) % 32)
					v105 = int32(0)
					if v64 != int32(1) {
						v113 = v105
						v121 = v2
						for {
							v127 = int32(2)
							v129 = v102 + v113<<(uint(v127)%32)
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
							v131 = base.I32_rem_u_s(v130, v104)
							v132 = int32(3)
							v134 = v56 + int32(base.Ui32(v131)>>(uint(v132)%32))
							v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
							v136 = int32(1)
							v137 = int32(7)
							v140 = v135 | v136<<(uint(v131&v137)%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v140)
							v142 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
							v143 = base.I32_rem_u_s(v142, v104)
							v146 = v56 + int32(base.Ui32(v143)>>(uint(v132)%32))
							v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
							v152 = v147 | v136<<(uint(v143&v137)%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v146))) = uint8(v152)
							v155 = v113 + v127
							v157 = v121 + v127
							if v157 != v64&int32(1073741822) {
								v113 = v155
								v121 = v157
								continue
							} else {
								break
							}
							break
						}
						if v64&int32(1) == int32(0) {
						} else {
							v162 = v155
							v179 = *(*int32)(unsafe.Add(mBase, uint32(v102+v162<<(uint(int32(2))%32))))
							v180 = base.I32_rem_u_s(v179, v104)
							v183 = v56 + int32(base.Ui32(v180)>>(uint(int32(3))%32))
							v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
							v189 = v184 | int32(1)<<(uint(v180&int32(7))%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v189)
						}
					} else {
						v162 = v105
						v179 = *(*int32)(unsafe.Add(mBase, uint32(v102+v162<<(uint(int32(2))%32))))
						v180 = base.I32_rem_u_s(v179, v104)
						v183 = v56 + int32(base.Ui32(v180)>>(uint(int32(3))%32))
						v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
						v189 = v184 | int32(1)<<(uint(v180&int32(7))%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v189)
					}
				}
				v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
				if v206&int32(4) != 0 {
					v209 = int32(3)
					v210 = v46 << (uint(v209) % 32)
					if v209 < v46 {
						v654 = int64(0)
						if base.B2i32(v56 != (v56+int32(3))&int32(-4))|base.B2i32(v46 < int32(4)) != 0 {
							v733 = v56
							v734 = v46
							v739 = v654
						} else {
							v664 = v46 - int32(4)
							v668 = int32(base.Ui32(v664)>>(uint(int32(2))%32)) + int32(1)
							v670 = v668 & int32(3)
							if base.Ui32(int32(12)) <= base.Ui32(v664) {
								v675 = v56
								v676 = v46
								v679 = int32(0)
								v681 = v654
								for {
									v682 = int32(16)
									v683 = v676 - v682
									v685 = v675 + v682
									v686 = *(*int32)(unsafe.Add(mBase, uint32(v675)+12))
									v689 = *(*int32)(unsafe.Add(mBase, uint32(v675)+8))
									v692 = *(*int32)(unsafe.Add(mBase, uint32(v675)+4))
									v695 = *(*int32)(unsafe.Add(mBase, uint32(v675)))
									v701 = base.I64_extend_i32_u(base.I32_popcnt(v686)) + (base.I64_extend_i32_u(base.I32_popcnt(v689)) + (base.I64_extend_i32_u(base.I32_popcnt(v692)) + (v681 + base.I64_extend_i32_u(base.I32_popcnt(v695)))))
									v703 = v679 + int32(4)
									if v703 != v668&int32(2147483644) {
										v675 = v685
										v676 = v683
										v679 = v703
										v681 = v701
										continue
									} else {
										break
									}
									break
								}
								if v670 == int32(0) {
									v733 = v685
									v734 = v683
									v739 = v701
								} else {
									v707 = v685
									v708 = v683
									v713 = v701
									v715 = v707
									v716 = v708
									v717 = int32(0)
									v721 = v713
									for {
										v722 = int32(4)
										v723 = v716 - v722
										v725 = v715 + v722
										v726 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
										v729 = v721 + base.I64_extend_i32_u(base.I32_popcnt(v726))
										v731 = v717 + int32(1)
										if v731 != v670 {
											v715 = v725
											v716 = v723
											v717 = v731
											v721 = v729
											continue
										} else {
											break
										}
										break
									}
									v733 = v725
									v734 = v723
									v739 = v729
								}
							} else {
								v707 = v56
								v708 = v46
								v713 = v654
								v715 = v707
								v716 = v708
								v717 = int32(0)
								v721 = v713
								for {
									v722 = int32(4)
									v723 = v716 - v722
									v725 = v715 + v722
									v726 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
									v729 = v721 + base.I64_extend_i32_u(base.I32_popcnt(v726))
									v731 = v717 + int32(1)
									if v731 != v670 {
										v715 = v725
										v716 = v723
										v717 = v731
										v721 = v729
										continue
									} else {
										break
									}
									break
								}
								v733 = v725
								v734 = v723
								v739 = v729
							}
						}
						if v734 == int32(0) {
							v802 = v739
						} else {
							v743 = v734 & int32(3)
							if v743 == int32(0) {
								v764 = v733
								v766 = v734
								v770 = v739
							} else {
								v747 = v733
								v749 = v734
								v751 = int32(0)
								v753 = v739
								for {
									v754 = int32(1)
									v755 = v747 + v754
									v757 = v749 - v754
									v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747))))
									v759 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v758)+uint32(_c_F_gtsvector_penalty[0]))))
									v760 = v753 + v759
									v762 = v751 + v754
									if v762 != v743 {
										v747 = v755
										v749 = v757
										v751 = v762
										v753 = v760
										continue
									} else {
										break
									}
									break
								}
								v764 = v755
								v766 = v757
								v770 = v760
							}
							if base.Ui32(v734) < base.Ui32(int32(4)) {
								v802 = v770
							} else {
								v773 = v764
								v775 = v766
								v779 = v770
								for {
									v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773)+3)))
									v781 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v780)+uint32(_c_F_gtsvector_penalty[0]))))
									v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773)+2)))
									v783 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v782)+uint32(_c_F_gtsvector_penalty[0]))))
									v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773)+1)))
									v785 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v784)+uint32(_c_F_gtsvector_penalty[0]))))
									v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773))))
									v787 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v786)+uint32(_c_F_gtsvector_penalty[0]))))
									v791 = v781 + (v783 + (v785 + (v779 + v787)))
									v792 = int32(4)
									v795 = v775 - v792
									if v795 != 0 {
										v773 = v773 + v792
										v775 = v795
										v779 = v791
										continue
									} else {
										break
									}
									break
								}
								v802 = v791
							}
						}
						v1175 = v802
					} else {
						if v46 == int32(0) {
							v1175 = v15
						} else {
							v216 = v46 & int32(3)
							if base.Ui32(int32(4)) <= base.Ui32(v46) {
								v223 = v56
								v224 = int32(0)
								v236 = v15
								for {
									v237 = int32(4)
									v238 = v223 + v237
									v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+3)))
									v240 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v239)+uint32(_c_F_gtsvector_penalty[0]))))
									v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+2)))
									v242 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v241)+uint32(_c_F_gtsvector_penalty[0]))))
									v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
									v244 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v243)+uint32(_c_F_gtsvector_penalty[0]))))
									v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
									v246 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v245)+uint32(_c_F_gtsvector_penalty[0]))))
									v250 = v240 + (v242 + (v244 + (v236 + v246)))
									v252 = v224 + v237
									if v252 != v46&int32(-4) {
										v223 = v238
										v224 = v252
										v236 = v250
										continue
									} else {
										break
									}
									break
								}
								if v216 == int32(0) {
									v1175 = v250
								} else {
									v257 = v238
									v270 = v250
									v273 = v257
									v276 = int32(0)
									v286 = v270
									for {
										v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
										v288 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v287)+uint32(_c_F_gtsvector_penalty[0]))))
										v289 = v286 + v288
										v290 = int32(1)
										v293 = v276 + v290
										if v293 != v216 {
											v273 = v273 + v290
											v276 = v293
											v286 = v289
											continue
										} else {
											break
										}
										break
									}
									v1175 = v289
								}
							} else {
								v257 = v56
								v270 = v15
								v273 = v257
								v276 = int32(0)
								v286 = v270
								for {
									v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
									v288 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v287)+uint32(_c_F_gtsvector_penalty[0]))))
									v289 = v286 + v288
									v290 = int32(1)
									v293 = v276 + v290
									if v293 != v216 {
										v273 = v273 + v290
										v276 = v293
										v286 = v289
										continue
									} else {
										break
									}
									break
								}
								v1175 = v289
							}
						}
					}
					v1198 = base.F32_div(base.F32_convert_i32_s(v210-base.I32_wrap_i64(v1175)), base.F32_convert_i32_s(v210|int32(1)))
				} else {
					if v46 <= int32(0) {
						v1198 = float32(0)
					} else {
						v298 = int32(0)
						if v46 != int32(1) {
							v308 = v298
							v309 = v298
							v316 = int32(0)
							for {
								v323 = v308 | int32(1)
								v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v323))))
								v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323+v56))))
								v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325^v327)+uint32(_c_F_gtsvector_penalty[0]))))
								v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308+v52))))
								v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308+v56))))
								v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331^v333)+uint32(_c_F_gtsvector_penalty[0]))))
								v337 = v329 + (v309 + v335)
								v338 = int32(2)
								v339 = v308 + v338
								v341 = v316 + v338
								if v341 != v46&int32(2147483646) {
									v308 = v339
									v309 = v337
									v316 = v341
									continue
								} else {
									break
								}
								break
							}
							if v46&int32(1) == int32(0) {
								v369 = v337
							} else {
								v346 = v339
								v347 = v337
								v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346+v52))))
								v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346+v56))))
								v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361^v363)+uint32(_c_F_gtsvector_penalty[0]))))
								v369 = v347 + v365
							}
						} else {
							v346 = v298
							v347 = v298
							v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346+v52))))
							v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346+v56))))
							v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361^v363)+uint32(_c_F_gtsvector_penalty[0]))))
							v369 = v347 + v365
						}
						v1198 = base.F32_convert_i32_s(v369)
					}
				}
				*(*float32)(unsafe.Add(mBase, uint32(v16))) = v1198
				F_pfree(m, v56)
				mBase = m.M
				v1201 = m.ExcPending
				if v1201 != 0 {
					return int32(0)
				} else {
					return v16
				}
			}
		} else {
			v383 = int32(4)
			v384 = v53 & v383
			v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
			if v385&v383 != 0 {
				if v384 != 0 {
					v1144 = int32(0)
				} else {
					v389 = int32(8)
					v390 = v47 + v389
					v391 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
					v393 = int32(base.Ui32(v391) >> (uint(int32(2)) % 32))
					v395 = v393 - v389
					if base.Ui32(int32(47)) < base.Ui32(v391) {
						v807 = int64(0)
						if base.B2i32(v390 != (v47+int32(11))&int32(-4))|base.B2i32(v395 < int32(4)) != 0 {
							v886 = v390
							v887 = v395
							v892 = v807
						} else {
							v817 = v395 - int32(4)
							v821 = int32(base.Ui32(v817)>>(uint(int32(2))%32)) + int32(1)
							v823 = v821 & int32(3)
							if base.Ui32(int32(12)) <= base.Ui32(v817) {
								v828 = v390
								v829 = v395
								v832 = int32(0)
								v834 = v807
								for {
									v835 = int32(16)
									v836 = v829 - v835
									v838 = v828 + v835
									v839 = *(*int32)(unsafe.Add(mBase, uint32(v828)+12))
									v842 = *(*int32)(unsafe.Add(mBase, uint32(v828)+8))
									v845 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
									v848 = *(*int32)(unsafe.Add(mBase, uint32(v828)))
									v854 = base.I64_extend_i32_u(base.I32_popcnt(v839)) + (base.I64_extend_i32_u(base.I32_popcnt(v842)) + (base.I64_extend_i32_u(base.I32_popcnt(v845)) + (v834 + base.I64_extend_i32_u(base.I32_popcnt(v848)))))
									v856 = v832 + int32(4)
									if v856 != v821&int32(2147483644) {
										v828 = v838
										v829 = v836
										v832 = v856
										v834 = v854
										continue
									} else {
										break
									}
									break
								}
								if v823 == int32(0) {
									v886 = v838
									v887 = v836
									v892 = v854
								} else {
									v860 = v838
									v861 = v836
									v866 = v854
									v868 = v860
									v869 = v861
									v870 = int32(0)
									v874 = v866
									for {
										v875 = int32(4)
										v876 = v869 - v875
										v878 = v868 + v875
										v879 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
										v882 = v874 + base.I64_extend_i32_u(base.I32_popcnt(v879))
										v884 = v870 + int32(1)
										if v884 != v823 {
											v868 = v878
											v869 = v876
											v870 = v884
											v874 = v882
											continue
										} else {
											break
										}
										break
									}
									v886 = v878
									v887 = v876
									v892 = v882
								}
							} else {
								v860 = v390
								v861 = v395
								v866 = v807
								v868 = v860
								v869 = v861
								v870 = int32(0)
								v874 = v866
								for {
									v875 = int32(4)
									v876 = v869 - v875
									v878 = v868 + v875
									v879 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
									v882 = v874 + base.I64_extend_i32_u(base.I32_popcnt(v879))
									v884 = v870 + int32(1)
									if v884 != v823 {
										v868 = v878
										v869 = v876
										v870 = v884
										v874 = v882
										continue
									} else {
										break
									}
									break
								}
								v886 = v878
								v887 = v876
								v892 = v882
							}
						}
						if v887 == int32(0) {
							v955 = v892
						} else {
							v896 = v887 & int32(3)
							if v896 == int32(0) {
								v917 = v886
								v919 = v887
								v923 = v892
							} else {
								v900 = v886
								v902 = v887
								v904 = int32(0)
								v906 = v892
								for {
									v907 = int32(1)
									v908 = v900 + v907
									v910 = v902 - v907
									v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v900))))
									v912 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v911)+uint32(_c_F_gtsvector_penalty[0]))))
									v913 = v906 + v912
									v915 = v904 + v907
									if v915 != v896 {
										v900 = v908
										v902 = v910
										v904 = v915
										v906 = v913
										continue
									} else {
										break
									}
									break
								}
								v917 = v908
								v919 = v910
								v923 = v913
							}
							if base.Ui32(v887) < base.Ui32(int32(4)) {
								v955 = v923
							} else {
								v926 = v917
								v928 = v919
								v932 = v923
								for {
									v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926)+3)))
									v934 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v933)+uint32(_c_F_gtsvector_penalty[0]))))
									v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926)+2)))
									v936 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v935)+uint32(_c_F_gtsvector_penalty[0]))))
									v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926)+1)))
									v938 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v937)+uint32(_c_F_gtsvector_penalty[0]))))
									v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926))))
									v940 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v939)+uint32(_c_F_gtsvector_penalty[0]))))
									v944 = v934 + (v936 + (v938 + (v932 + v940)))
									v945 = int32(4)
									v948 = v928 - v945
									if v948 != 0 {
										v926 = v926 + v945
										v928 = v948
										v932 = v944
										continue
									} else {
										break
									}
									break
								}
								v955 = v944
							}
						}
						v1140 = v955
					} else {
						if v395 == int32(0) {
							v1140 = v15
						} else {
							v402 = int32(3)
							v403 = v393 & v402
							if base.Ui32(v402) <= base.Ui32(v393-int32(9)) {
								v411 = v390
								v415 = int32(0)
								v425 = v15
								for {
									v426 = int32(4)
									v427 = v411 + v426
									v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+3)))
									v429 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v428)+uint32(_c_F_gtsvector_penalty[0]))))
									v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+2)))
									v431 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v430)+uint32(_c_F_gtsvector_penalty[0]))))
									v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+1)))
									v433 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v432)+uint32(_c_F_gtsvector_penalty[0]))))
									v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
									v435 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v434)+uint32(_c_F_gtsvector_penalty[0]))))
									v439 = v429 + (v431 + (v433 + (v425 + v435)))
									v441 = v415 + v426
									if v441 != v395&int32(-4) {
										v411 = v427
										v415 = v441
										v425 = v439
										continue
									} else {
										break
									}
									break
								}
								if v403 == int32(0) {
									v1140 = v439
								} else {
									v445 = v427
									v459 = v439
									v461 = v445
									v462 = int32(0)
									v475 = v459
									for {
										v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
										v477 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v476)+uint32(_c_F_gtsvector_penalty[0]))))
										v478 = v475 + v477
										v479 = int32(1)
										v482 = v462 + v479
										if v482 != v403 {
											v461 = v461 + v479
											v462 = v482
											v475 = v478
											continue
										} else {
											break
										}
										break
									}
									v1140 = v478
								}
							} else {
								v445 = v390
								v459 = v15
								v461 = v445
								v462 = int32(0)
								v475 = v459
								for {
									v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
									v477 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v476)+uint32(_c_F_gtsvector_penalty[0]))))
									v478 = v475 + v477
									v479 = int32(1)
									v482 = v462 + v479
									if v482 != v403 {
										v461 = v461 + v479
										v462 = v482
										v475 = v478
										continue
									} else {
										break
									}
									break
								}
								v1140 = v478
							}
						}
					}
					v1144 = v395<<(uint(int32(3))%32) - base.I32_wrap_i64(v1140)
				}
			} else {
				v484 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
				v486 = int32(base.Ui32(v484) >> (uint(int32(2)) % 32))
				v488 = v486 - int32(8)
				if v384 != 0 {
					if base.Ui32(int32(47)) < base.Ui32(v484) {
						v960 = int64(0)
						if base.B2i32(v52 != (v48+int32(11))&int32(-4))|base.B2i32(v488 < int32(4)) != 0 {
							v1039 = v52
							v1040 = v488
							v1045 = v960
						} else {
							v970 = v488 - int32(4)
							v974 = int32(base.Ui32(v970)>>(uint(int32(2))%32)) + int32(1)
							v976 = v974 & int32(3)
							if base.Ui32(int32(12)) <= base.Ui32(v970) {
								v981 = v52
								v982 = v488
								v985 = int32(0)
								v987 = v960
								for {
									v988 = int32(16)
									v989 = v982 - v988
									v991 = v981 + v988
									v992 = *(*int32)(unsafe.Add(mBase, uint32(v981)+12))
									v995 = *(*int32)(unsafe.Add(mBase, uint32(v981)+8))
									v998 = *(*int32)(unsafe.Add(mBase, uint32(v981)+4))
									v1001 = *(*int32)(unsafe.Add(mBase, uint32(v981)))
									v1007 = base.I64_extend_i32_u(base.I32_popcnt(v992)) + (base.I64_extend_i32_u(base.I32_popcnt(v995)) + (base.I64_extend_i32_u(base.I32_popcnt(v998)) + (v987 + base.I64_extend_i32_u(base.I32_popcnt(v1001)))))
									v1009 = v985 + int32(4)
									if v1009 != v974&int32(2147483644) {
										v981 = v991
										v982 = v989
										v985 = v1009
										v987 = v1007
										continue
									} else {
										break
									}
									break
								}
								if v976 == int32(0) {
									v1039 = v991
									v1040 = v989
									v1045 = v1007
								} else {
									v1013 = v991
									v1014 = v989
									v1019 = v1007
									v1021 = v1013
									v1022 = v1014
									v1023 = int32(0)
									v1027 = v1019
									for {
										v1028 = int32(4)
										v1029 = v1022 - v1028
										v1031 = v1021 + v1028
										v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1021)))
										v1035 = v1027 + base.I64_extend_i32_u(base.I32_popcnt(v1032))
										v1037 = v1023 + int32(1)
										if v1037 != v976 {
											v1021 = v1031
											v1022 = v1029
											v1023 = v1037
											v1027 = v1035
											continue
										} else {
											break
										}
										break
									}
									v1039 = v1031
									v1040 = v1029
									v1045 = v1035
								}
							} else {
								v1013 = v52
								v1014 = v488
								v1019 = v960
								v1021 = v1013
								v1022 = v1014
								v1023 = int32(0)
								v1027 = v1019
								for {
									v1028 = int32(4)
									v1029 = v1022 - v1028
									v1031 = v1021 + v1028
									v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1021)))
									v1035 = v1027 + base.I64_extend_i32_u(base.I32_popcnt(v1032))
									v1037 = v1023 + int32(1)
									if v1037 != v976 {
										v1021 = v1031
										v1022 = v1029
										v1023 = v1037
										v1027 = v1035
										continue
									} else {
										break
									}
									break
								}
								v1039 = v1031
								v1040 = v1029
								v1045 = v1035
							}
						}
						if v1040 == int32(0) {
							v1108 = v1045
						} else {
							v1049 = v1040 & int32(3)
							if v1049 == int32(0) {
								v1070 = v1039
								v1072 = v1040
								v1076 = v1045
							} else {
								v1053 = v1039
								v1055 = v1040
								v1057 = int32(0)
								v1059 = v1045
								for {
									v1060 = int32(1)
									v1061 = v1053 + v1060
									v1063 = v1055 - v1060
									v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053))))
									v1065 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1064)+uint32(_c_F_gtsvector_penalty[0]))))
									v1066 = v1059 + v1065
									v1068 = v1057 + v1060
									if v1068 != v1049 {
										v1053 = v1061
										v1055 = v1063
										v1057 = v1068
										v1059 = v1066
										continue
									} else {
										break
									}
									break
								}
								v1070 = v1061
								v1072 = v1063
								v1076 = v1066
							}
							if base.Ui32(v1040) < base.Ui32(int32(4)) {
								v1108 = v1076
							} else {
								v1079 = v1070
								v1081 = v1072
								v1085 = v1076
								for {
									v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+3)))
									v1087 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1086)+uint32(_c_F_gtsvector_penalty[0]))))
									v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+2)))
									v1089 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1088)+uint32(_c_F_gtsvector_penalty[0]))))
									v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+1)))
									v1091 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1090)+uint32(_c_F_gtsvector_penalty[0]))))
									v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079))))
									v1093 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1092)+uint32(_c_F_gtsvector_penalty[0]))))
									v1097 = v1087 + (v1089 + (v1091 + (v1085 + v1093)))
									v1098 = int32(4)
									v1101 = v1081 - v1098
									if v1101 != 0 {
										v1079 = v1079 + v1098
										v1081 = v1101
										v1085 = v1097
										continue
									} else {
										break
									}
									break
								}
								v1108 = v1097
							}
						}
						v1123 = v1108
					} else {
						if v488 == int32(0) {
							v1123 = v15
						} else {
							v495 = int32(3)
							v496 = v486 & v495
							if base.Ui32(v495) <= base.Ui32(v486-int32(9)) {
								v505 = int32(0)
								v508 = v52
								v518 = v15
								for {
									v519 = int32(4)
									v520 = v508 + v519
									v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+3)))
									v522 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v521)+uint32(_c_F_gtsvector_penalty[0]))))
									v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+2)))
									v524 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v523)+uint32(_c_F_gtsvector_penalty[0]))))
									v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+1)))
									v526 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v525)+uint32(_c_F_gtsvector_penalty[0]))))
									v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
									v528 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v527)+uint32(_c_F_gtsvector_penalty[0]))))
									v532 = v522 + (v524 + (v526 + (v518 + v528)))
									v534 = v505 + v519
									if v534 != v488&int32(-4) {
										v505 = v534
										v508 = v520
										v518 = v532
										continue
									} else {
										break
									}
									break
								}
								if v496 == int32(0) {
									v1123 = v532
								} else {
									v542 = v520
									v552 = v532
									v554 = int32(0)
									v558 = v542
									v568 = v552
									for {
										v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
										v570 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v569)+uint32(_c_F_gtsvector_penalty[0]))))
										v571 = v568 + v570
										v572 = int32(1)
										v575 = v554 + v572
										if v575 != v496 {
											v554 = v575
											v558 = v558 + v572
											v568 = v571
											continue
										} else {
											break
										}
										break
									}
									v1123 = v571
								}
							} else {
								v542 = v52
								v552 = v15
								v554 = int32(0)
								v558 = v542
								v568 = v552
								for {
									v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
									v570 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v569)+uint32(_c_F_gtsvector_penalty[0]))))
									v571 = v568 + v570
									v572 = int32(1)
									v575 = v554 + v572
									if v575 != v496 {
										v554 = v575
										v558 = v558 + v572
										v568 = v571
										continue
									} else {
										break
									}
									break
								}
								v1123 = v571
							}
						}
					}
					v1144 = v488<<(uint(int32(3))%32) - base.I32_wrap_i64(v1123)
				} else {
					if base.Ui32(v484) < base.Ui32(int32(36)) {
						v1144 = int32(0)
					} else {
						v581 = v47 + int32(8)
						v582 = int32(0)
						if v486 != int32(9) {
							v590 = v582
							v591 = v582
							v599 = v2
							for {
								v606 = v590 | int32(1)
								v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581+v606))))
								v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v606+v52))))
								v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608^v610)+uint32(_c_F_gtsvector_penalty[0]))))
								v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590+v581))))
								v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590+v52))))
								v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614^v616)+uint32(_c_F_gtsvector_penalty[0]))))
								v620 = v612 + (v591 + v618)
								v621 = int32(2)
								v622 = v590 + v621
								v624 = v599 + v621
								if v624 != v488&int32(-2) {
									v590 = v622
									v591 = v620
									v599 = v624
									continue
								} else {
									break
								}
								break
							}
							if v486&int32(1) == int32(0) {
								v1144 = v620
							} else {
								v628 = v622
								v629 = v620
								v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628+v581))))
								v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628+v52))))
								v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644^v646)+uint32(_c_F_gtsvector_penalty[0]))))
								v1144 = v629 + v648
							}
						} else {
							v628 = v582
							v629 = v582
							v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628+v581))))
							v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628+v52))))
							v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644^v646)+uint32(_c_F_gtsvector_penalty[0]))))
							v1144 = v629 + v648
						}
					}
				}
			}
			*(*float32)(unsafe.Add(mBase, uint32(v16))) = base.F32_convert_i32_s(v1144)
			return v16
		}
	}
}
