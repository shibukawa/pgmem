package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_seg_contained(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_seg_contained_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_seg_left(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float32)(unsafe.Add(mBase, uint32(v2)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float32)(unsafe.Add(mBase, uint32(v4)))
	return base.F32_lt(v3, v5)
}
func F_seg_over_right(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float32)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float32)(unsafe.Add(mBase, uint32(v4)))
	return base.F32_ge(v3, v5)
}
func F_seg_overlap(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float32
	_ = v6
	var v7 int32
	_ = v7
	var v8 float32
	_ = v8
	var v12 float32
	_ = v12
	var v23 float32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*float32)(unsafe.Add(mBase, uint32(v7)+4))
	if base.F32_ge(v6, v8) == int32(0) {
		if base.F32_le(v6, v8) == int32(0) {
			return int32(0)
		} else {
			v23 = *(*float32)(unsafe.Add(mBase, uint32(v7)))
			return base.F32_le(v23, v6)
		}
	} else {
		v12 = *(*float32)(unsafe.Add(mBase, uint32(v5)))
		if base.F32_le(v12, v8) == int32(0) {
			if base.F32_le(v6, v8) == int32(0) {
				return int32(0)
			} else {
				v23 = *(*float32)(unsafe.Add(mBase, uint32(v7)))
				return base.F32_le(v23, v6)
			}
		} else {
			return int32(1)
		}
	}
}
func F_seg_yylex(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
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
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
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
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
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
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1032 int32
	_ = v1032
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1152 int32
	_ = v1152
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = l0
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v70 = l1
	v72 = v68
	v81 = int32(0)
	goto L22
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v68 = v16
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v19 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = int32(1)
	goto L7
L6:
	;
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v24 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_seg_yylex[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v28
	goto L10
L9:
	;
	goto L10
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_seg_yylex[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v34
	goto L13
L12:
	;
	goto L13
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v36 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v60
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v63
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v65)
	v68 = v60
	goto L1
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36+v37<<(uint(int32(2))%32))))
	if v41 != 0 {
		v57 = v41
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_seg_yyensure_buffer_stack(m, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	return int32(0)
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v49 = F_seg_yy_create_buffer(m, v47, int32(_a_F_seg_yylex_0), l1)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v51+v52<<(uint(int32(2))%32)))) = v49
	v57 = v49
	goto L14
L22:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v82)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v70)+44))
	v86 = v70
	v87 = v84
	v88 = v72
	v91 = v72
	v97 = v81
	goto L24
L24:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_seg_yylex[2]))))
	if int32(1)<<(uint(v87)%32)&int32(619905031) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v87
	goto L28
L27:
	;
	goto L28
L28:
	;
	v110 = int32(1)
	v114 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87<<(uint(v110)%32))+uint32(_c_F_seg_yylex[3]))))
	v115 = v101 + v114
	v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v115<<(uint(v110)%32))+uint32(_c_F_seg_yylex[4]))))
	if v120 != v87 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v124 = v87
	v126 = v101
	v127 = v101
	goto L32
L30:
	;
	v162 = v115
	goto L31
L31:
	;
	v175 = int32(1)
	v179 = int32(*(*int16)(unsafe.Add(mBase, uint32(v162<<(uint(v175)%32))+uint32(_c_F_seg_yylex[5]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v162))%64)&int64(72054295503044608) == int64(0) {
		v87 = v179
		v91 = v91 + v175
		goto L24
	} else {
		goto L38
	}
L32:
	;
	v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(v124<<(uint(int32(1))%32))+uint32(_c_F_seg_yylex[6]))))
	if v124 == int32(21) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v162 = v153
	goto L31
L34:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+uint32(_c_F_seg_yylex[7]))))
	v145 = v144
	goto L36
L35:
	;
	v145 = v127
	goto L36
L36:
	;
	v147 = v145 & int32(255)
	v148 = int32(1)
	v152 = int32(*(*int16)(unsafe.Add(mBase, uint32(v139<<(uint(v148)%32))+uint32(_c_F_seg_yylex[3]))))
	v153 = v147 + v152
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153<<(uint(v148)%32))+uint32(_c_F_seg_yylex[4]))))
	if v158 != v139&int32(_a_F_seg_yylex_1) {
		v124 = v139
		v126 = v147
		v127 = v145
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v196 = v88
	goto L39
L39:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v86)+64))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v86)+68))
	v206 = v202
	v211 = v196
	v212 = v203
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+80)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v86)+32)) = v212 - v211
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+24)) = uint8(v220)
	v222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v212))) = uint8(v222)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v212
	v229 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206<<(uint(int32(1))%32))+uint32(_c_F_seg_yylex[8]))))
	v230 = v229
	v233 = v212
	goto L43
L43:
	;
	if v230 != int32(10) {
		goto L53
	} else {
		goto L54
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v1222
	*(*int32)(unsafe.Add(mBase, uint32(v86)+48)) = int32(0)
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	v1239 = base.I32_div_s(v1235-int32(1), int32(2))
	v230 = v1239 + int32(11)
	v233 = v1222
	goto L43
L46:
	;
	F_yy_fatal_error_7(m, int32(_a_F_seg_yylex_2))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L19
	} else {
		goto L240
	}
L47:
	;
	F_yy_fatal_error_7(m, int32(_a_F_seg_yylex_3))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L19
	} else {
		goto L239
	}
L48:
	;
	F_yy_fatal_error_7(m, int32(_a_F_seg_yylex_4))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L19
	} else {
		goto L238
	}
L49:
	;
	F_yy_fatal_error_7(m, int32(_a_F_seg_yylex_5))
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L19
	} else {
		goto L237
	}
L50:
	;
	return v1203
L51:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1199))) = v1200
	v1203 = int32(259)
	goto L50
L52:
	;
	F_yy_fatal_error_7(m, int32(_a_F_seg_yylex_6))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L19
	} else {
		goto L236
	}
L53:
	;
	switch v230 {
	case 0:
		goto L63
	case 1:
		goto L51
	case 2:
		goto L62
	case 3:
		goto L61
	case 4:
		goto L60
	case 5:
		goto L59
	case 6:
		goto L58
	case 7:
		v70 = v86
		v72 = v233
		v81 = v97
		goto L22
	case 8:
		goto L57
	case 9:
		goto L56
	default:
		goto L52
	case 11:
		v1203 = v97
		goto L50
	}
L54:
	;
	goto L55
L55:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v212))) = uint8(v279)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v281+v282<<(uint(int32(2))%32))))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+44))
	if v287 != 0 {
		goto L66
	} else {
		goto L67
	}
L56:
	;
	F_yy_fatal_error_7(m, int32(_a_F_seg_yylex_7))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L19
	} else {
		goto L64
	}
L57:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	v273 = int32(*(*int8)(unsafe.Add(mBase, uint32(v272))))
	return v273
L58:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = int32(_a_F_seg_yylex_8)
	return int32(261)
L59:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = int32(_a_F_seg_yylex_9)
	return int32(261)
L60:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(_a_F_seg_yylex_10)
	return int32(261)
L61:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v253
	return int32(258)
L62:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v248
	return int32(260)
L63:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v212))) = uint8(v245)
	v196 = v211
	goto L39
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v86)+36))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	v301 = v300 + v298
	if base.Ui32(v299) <= base.Ui32(v301) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v297 = v287
	v298 = v288
	goto L65
L67:
	;
	goto L68
L68:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v286)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v289
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v292 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v286)+44)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v286))) = v291
	v297 = v292
	v298 = v289
	goto L65
L69:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	v306 = v278 ^ int32(-1) + v212
	v307 = v303 + v306
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v307
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	if int32(0) < v306 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	if base.Ui32(v301+int32(1)) < base.Ui32(v299) {
		goto L49
	} else {
		goto L103
	}
L72:
	;
	v314 = v309
	v318 = v303
	goto L75
L73:
	;
	v416 = v309
	goto L74
L74:
	;
	if int32(1)<<(uint(v416)%32)&int32(619905031) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L75:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	if v325 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v416 = v410
	goto L74
L77:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+uint32(_c_F_seg_yylex[2]))))
	v330 = v328
	goto L79
L78:
	;
	v330 = int32(1)
	goto L79
L79:
	;
	if int32(1)<<(uint(v314)%32)&int32(619905031) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v314
	goto L82
L81:
	;
	goto L82
L82:
	;
	v340 = v330 & int32(255)
	v341 = int32(1)
	v345 = int32(*(*int16)(unsafe.Add(mBase, uint32(v314<<(uint(v341)%32))+uint32(_c_F_seg_yylex[3]))))
	v346 = v340 + v345
	v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v346<<(uint(v341)%32))+uint32(_c_F_seg_yylex[4]))))
	if v351 != v314 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v355 = v314
	v357 = v340
	v358 = v330
	goto L86
L84:
	;
	v393 = v346
	goto L85
L85:
	;
	v406 = int32(1)
	v410 = int32(*(*int16)(unsafe.Add(mBase, uint32(v393<<(uint(v406)%32))+uint32(_c_F_seg_yylex[5]))))
	v412 = v318 + v406
	if v412 != v307 {
		v314 = v410
		v318 = v412
		goto L75
	} else {
		goto L92
	}
L86:
	;
	v370 = int32(*(*int16)(unsafe.Add(mBase, uint32(v355<<(uint(int32(1))%32))+uint32(_c_F_seg_yylex[6]))))
	if v355 == int32(21) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v393 = v384
	goto L85
L88:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357)+uint32(_c_F_seg_yylex[7]))))
	v376 = v375
	goto L90
L89:
	;
	v376 = v358
	goto L90
L90:
	;
	v378 = v376 & int32(255)
	v379 = int32(1)
	v383 = int32(*(*int16)(unsafe.Add(mBase, uint32(v370<<(uint(v379)%32))+uint32(_c_F_seg_yylex[3]))))
	v384 = v378 + v383
	v389 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384<<(uint(v379)%32))+uint32(_c_F_seg_yylex[4]))))
	if v389 != v370&int32(_a_F_seg_yylex_1) {
		v355 = v370
		v357 = v378
		v358 = v376
		goto L86
	} else {
		goto L91
	}
L91:
	;
	goto L87
L92:
	;
	goto L76
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v416
	goto L95
L94:
	;
	goto L95
L95:
	;
	v435 = int32(1)
	v439 = int32(*(*int16)(unsafe.Add(mBase, uint32(v416<<(uint(v435)%32))+uint32(_c_F_seg_yylex[3]))))
	v441 = v439 + v435
	v446 = int32(*(*int16)(unsafe.Add(mBase, uint32(v441<<(uint(v435)%32))+uint32(_c_F_seg_yylex[4]))))
	if v446 != v416 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v450 = v416
	goto L99
L97:
	;
	v480 = v441
	goto L98
L98:
	;
	if base.B2i32(v480 == int32(0))|base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v480))%64)&int64(72054295503044608) != int64(0)) != 0 {
		v196 = v303
		goto L39
	} else {
		goto L102
	}
L99:
	;
	v461 = int32(1)
	v465 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450<<(uint(v461)%32))+uint32(_c_F_seg_yylex[6]))))
	v466 = base.I32_extend16_s(v465)
	v471 = int32(*(*int16)(unsafe.Add(mBase, uint32(v466<<(uint(v461)%32))+uint32(_c_F_seg_yylex[3]))))
	v473 = v471 + v461
	v478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v473<<(uint(v461)%32))+uint32(_c_F_seg_yylex[4]))))
	if v465 != v478 {
		v450 = v466
		goto L99
	} else {
		goto L101
	}
L100:
	;
	v480 = v473
	goto L98
L101:
	;
	goto L100
L102:
	;
	v503 = int32(1)
	v504 = v307 + v503
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v504
	v510 = int32(*(*int16)(unsafe.Add(mBase, uint32(v480<<(uint(v503)%32))+uint32(_c_F_seg_yylex[5]))))
	v87 = v510
	v88 = v303
	v91 = v504
	goto L24
L103:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v286)+40))
	if v515 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	v1090 = v1077 + v1082
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v1090
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	if base.Ui32(v1090) <= base.Ui32(v1080) {
		v206 = v1092
		v211 = v1080
		v212 = v1090
		goto L41
	} else {
		goto L217
	}
L106:
	;
	if v299-v514 != int32(1) {
		v1077 = v300
		v1080 = v514
		v1082 = v298
		goto L105
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v523 = v514 ^ int32(-1) + v299
	if int32(0) < v523 {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v1222 = v514
	goto L45
L110:
	;
	v908 = v894 + v523
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v896)+12))
	if v908 <= v909 {
		goto L185
	} else {
		goto L186
	}
L111:
	;
	if v523 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L112:
	;
	v526 = int32(7)
	v527 = v523 & v526
	if base.Ui32(v299-v514-int32(2)) < base.Ui32(v526) {
		goto L117
	} else {
		goto L118
	}
L113:
	;
	v632 = v286
	v643 = v297
	goto L114
L114:
	;
	if v643 == int32(2) {
		goto L127
	} else {
		goto L128
	}
L115:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v623+v624<<(uint(int32(2))%32))))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)+44))
	v632 = v628
	v643 = v629
	goto L114
L116:
	;
	v588 = v574
	v590 = v576
	v592 = int32(0)
	goto L124
L117:
	;
	v574 = v300
	v576 = v514
	goto L116
L118:
	;
	goto L119
L119:
	;
	v536 = v300
	v538 = v514
	v540 = int32(0)
	goto L120
L120:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538))))
	*(*uint8)(unsafe.Add(mBase, uint32(v536))) = uint8(v549)
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+1)) = uint8(v551)
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+2)) = uint8(v553)
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+3)) = uint8(v555)
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+4)) = uint8(v557)
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+5)) = uint8(v559)
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+6)) = uint8(v561)
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+7)) = uint8(v563)
	v565 = int32(8)
	v566 = v536 + v565
	v568 = v538 + v565
	v570 = v540 + v565
	if v570 != v523&int32(2147483640) {
		v536 = v566
		v538 = v568
		v540 = v570
		goto L120
	} else {
		goto L122
	}
L121:
	;
	if v527 == int32(0) {
		goto L115
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	v574 = v566
	v576 = v568
	goto L116
L124:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590))))
	*(*uint8)(unsafe.Add(mBase, uint32(v588))) = uint8(v601)
	v603 = int32(1)
	v608 = v592 + v603
	if v608 != v527 {
		v588 = v588 + v603
		v590 = v590 + v603
		v592 = v608
		goto L124
	} else {
		goto L126
	}
L125:
	;
	goto L115
L126:
	;
	goto L125
L127:
	;
	v646 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v632)+16)) = v646
	v866 = v632
	goto L111
L128:
	;
	goto L129
L129:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v632)+12))
	v651 = v514 - v299
	v652 = v650 + v651
	if v652 <= int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v86)+36))
	v656 = v650
	v658 = v632
	v661 = v655
	goto L133
L131:
	;
	v711 = v632
	v713 = v652
	goto L132
L132:
	;
	v722 = int32(16777216)
	if base.Ui32(v722) <= base.Ui32(v713) {
		goto L149
	} else {
		goto L150
	}
L133:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v658)+20))
	if v669 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v711 = v704
	v713 = v706
	goto L132
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v658)+4)) = int32(0)
	goto L46
L136:
	;
	goto L137
L137:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v658)+4))
	v675 = int32(0)
	if v656 <= v675 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v684 = v656 - int32(base.Ui32(v675-v656)>>(uint(int32(3))%32))
	goto L140
L139:
	;
	v684 = v656 << (uint(int32(1)) % 32)
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v658)+12)) = v684
	v687 = v684 + int32(2)
	if v674 != 0 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v658)+4)) = v692
	if v692 == int32(0) {
		goto L46
	} else {
		goto L147
	}
L142:
	;
	v688 = F_repalloc(m, v674, v687)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L19
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v690 = F_palloc(m, v687)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L19
	} else {
		goto L146
	}
L145:
	;
	v692 = v688
	goto L141
L146:
	;
	v692 = v690
	goto L141
L147:
	;
	v697 = v692 + (v661 - v674)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v697
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v699+v700<<(uint(int32(2))%32))))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v704)+12))
	v706 = v705 + v651
	if v706 <= int32(0) {
		v656 = v705
		v658 = v704
		v661 = v697
		goto L133
	} else {
		goto L148
	}
L148:
	;
	goto L134
L149:
	;
	v725 = v722
	goto L151
L150:
	;
	v725 = v713
	goto L151
L151:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v711)+24))
	if v727 != 0 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v854
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v856+v857<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v861)+16)) = v854
	if v854 != 0 {
		v894 = v854
		v896 = v861
		v907 = int32(0)
		goto L110
	} else {
		goto L179
	}
L153:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v841+v842<<(uint(int32(2))%32))))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v846)+4))
	v850 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v847+v523+v759))) = uint8(v850)
	v854 = v759 + int32(1)
	goto L152
L154:
	;
	v728 = int32(0)
	goto L158
L155:
	;
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_seg_yylex[9])) = int32(0)
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v711)+4))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v778 = F_fread(m, v774+v523, int32(1), v725, v777)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L19
	} else {
		goto L167
	}
L157:
	;
	switch v745 {
	case 0:
		goto L163
	default:
		v854 = v759
		goto L152
	case 11:
		goto L153
	}
L158:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v742 = F_do_getc(m, v741)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L19
	} else {
		goto L161
	}
L159:
	;
	v759 = v725
	goto L157
L160:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v746+v747<<(uint(int32(2))%32))))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v751)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v752+v523+v728))) = uint8(v742)
	v757 = v728 + int32(1)
	if v757 != v725 {
		v728 = v757
		goto L158
	} else {
		goto L162
	}
L161:
	;
	v745 = v742 + int32(1)
	switch v745 {
	case 0, 11:
		v759 = v728
		goto L157
	default:
		goto L160
	}
L162:
	;
	goto L159
L163:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	goto L164
L164:
	;
	if int32(base.Ui32(v761)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v854 = v759
		goto L152
	} else {
		goto L165
	}
L165:
	;
	F_yy_fatal_error_7(m, int32(_a_F_seg_yylex_4))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L19
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	v780 = v778
	goto L168
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v780
	if v780 != 0 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v833+v834<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v838)+16)) = v780
	v894 = v780
	v896 = v838
	v907 = int32(0)
	goto L110
L170:
	;
	goto L169
L171:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	goto L172
L172:
	;
	if int32(base.Ui32(v795)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v802+v803<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v807)+16)) = int32(0)
	v866 = v807
	goto L111
L174:
	;
	goto L175
L175:
	;
	v811 = *(*int32)(unsafe.Add(mBase, _c_F_seg_yylex[9]))
	if v811 != int32(27) {
		goto L48
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_seg_yylex[9])) = int32(0)
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	*(*int32)(unsafe.Add(mBase, uint32(v794))) = v817 & int32(-49)
	goto L177
L177:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v821+v822<<(uint(int32(2))%32))))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v826)+4))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v831 = F_fread(m, v827+v523, int32(1), v725, v830)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L19
	} else {
		goto L178
	}
L178:
	;
	v780 = v831
	goto L168
L179:
	;
	v866 = v861
	goto L111
L180:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	F_seg_yyrestart(m, v879, v86)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L19
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v890 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v866)+44)) = v890
	v894 = int32(0)
	v896 = v866
	v907 = v890
	goto L110
L183:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v882+v883<<(uint(int32(2))%32))))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v894 = v888
	v896 = v887
	v907 = int32(1)
	goto L110
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v938
	v941 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v935+v938))) = uint8(v941)
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v945 = int32(2)
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v943+v944<<(uint(v945)%32))))
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v949+v950)+1)) = uint8(v941)
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v954+v955<<(uint(v945)%32))))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v959)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+80)) = v960
	if v907 == int32(1) {
		v1222 = v960
		goto L45
	} else {
		goto L195
	}
L185:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v896)+4))
	v935 = v911
	v938 = v908
	goto L184
L186:
	;
	goto L187
L187:
	;
	v914 = v908 + v894>>(uint(int32(1))%32)
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v896)+4))
	if v915 != 0 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v921+v922<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v926)+4)) = v920
	if v920 == int32(0) {
		goto L47
	} else {
		goto L194
	}
L189:
	;
	v916 = F_repalloc(m, v915, v914)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L19
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v918 = F_palloc(m, v914)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L19
	} else {
		goto L193
	}
L192:
	;
	v920 = v916
	goto L188
L193:
	;
	v920 = v918
	goto L188
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v926)+12)) = v914 - int32(2)
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v935 = v920
	v938 = v933 + v523
	goto L184
L195:
	;
	switch v907 - int32(1) {
	case 0:
		goto L104
	case 1:
		goto L196
	default:
		goto L197
	}
L196:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v1077 = v960
	v1080 = v960
	v1082 = v1076
	goto L105
L197:
	;
	v968 = v278 ^ int32(-1) + v212
	v969 = v960 + v968
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v969
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	if v968 <= int32(0) {
		v87 = v971
		v88 = v960
		v91 = v969
		goto L24
	} else {
		goto L198
	}
L198:
	;
	v976 = v971
	v981 = v960
	goto L199
L199:
	;
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v981))))
	if v987 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v87 = v1072
	v88 = v960
	v91 = v969
	goto L24
L201:
	;
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987)+uint32(_c_F_seg_yylex[2]))))
	v992 = v990
	goto L203
L202:
	;
	v992 = int32(1)
	goto L203
L203:
	;
	if int32(1)<<(uint(v976)%32)&int32(619905031) == int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v976
	goto L206
L205:
	;
	goto L206
L206:
	;
	v1002 = v992 & int32(255)
	v1003 = int32(1)
	v1007 = int32(*(*int16)(unsafe.Add(mBase, uint32(v976<<(uint(v1003)%32))+uint32(_c_F_seg_yylex[3]))))
	v1008 = v1002 + v1007
	v1013 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1008<<(uint(v1003)%32))+uint32(_c_F_seg_yylex[4]))))
	if v1013 != v976 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1017 = v976
	v1019 = v1002
	v1020 = v992
	goto L210
L208:
	;
	v1055 = v1008
	goto L209
L209:
	;
	v1068 = int32(1)
	v1072 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1055<<(uint(v1068)%32))+uint32(_c_F_seg_yylex[5]))))
	v1074 = v981 + v1068
	if v969 != v1074 {
		v976 = v1072
		v981 = v1074
		goto L199
	} else {
		goto L216
	}
L210:
	;
	v1032 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1017<<(uint(int32(1))%32))+uint32(_c_F_seg_yylex[6]))))
	if v1017 == int32(21) {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	v1055 = v1046
	goto L209
L212:
	;
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019)+uint32(_c_F_seg_yylex[7]))))
	v1038 = v1037
	goto L214
L213:
	;
	v1038 = v1020
	goto L214
L214:
	;
	v1040 = v1038 & int32(255)
	v1041 = int32(1)
	v1045 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1032<<(uint(v1041)%32))+uint32(_c_F_seg_yylex[3]))))
	v1046 = v1040 + v1045
	v1051 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1046<<(uint(v1041)%32))+uint32(_c_F_seg_yylex[4]))))
	if v1051 != v1032&int32(_a_F_seg_yylex_1) {
		v1017 = v1032
		v1019 = v1040
		v1020 = v1038
		goto L210
	} else {
		goto L215
	}
L215:
	;
	goto L211
L216:
	;
	goto L200
L217:
	;
	v1096 = v1092
	v1100 = v1080
	goto L218
L218:
	;
	v1107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1100))))
	if v1107 != 0 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v206 = v1192
	v211 = v1080
	v212 = v1090
	goto L41
L220:
	;
	v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1107)+uint32(_c_F_seg_yylex[2]))))
	v1112 = v1110
	goto L222
L221:
	;
	v1112 = int32(1)
	goto L222
L222:
	;
	if int32(1)<<(uint(v1096)%32)&int32(619905031) == int32(0) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v1096
	goto L225
L224:
	;
	goto L225
L225:
	;
	v1122 = v1112 & int32(255)
	v1123 = int32(1)
	v1127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1096<<(uint(v1123)%32))+uint32(_c_F_seg_yylex[3]))))
	v1128 = v1122 + v1127
	v1133 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1128<<(uint(v1123)%32))+uint32(_c_F_seg_yylex[4]))))
	if v1133 != v1096 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v1137 = v1096
	v1139 = v1122
	v1140 = v1112
	goto L229
L227:
	;
	v1175 = v1128
	goto L228
L228:
	;
	v1188 = int32(1)
	v1192 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1175<<(uint(v1188)%32))+uint32(_c_F_seg_yylex[5]))))
	v1194 = v1100 + v1188
	if v1194 != v1090 {
		v1096 = v1192
		v1100 = v1194
		goto L218
	} else {
		goto L235
	}
L229:
	;
	v1152 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1137<<(uint(int32(1))%32))+uint32(_c_F_seg_yylex[6]))))
	if v1137 == int32(21) {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	v1175 = v1166
	goto L228
L231:
	;
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1139)+uint32(_c_F_seg_yylex[7]))))
	v1158 = v1157
	goto L233
L232:
	;
	v1158 = v1140
	goto L233
L233:
	;
	v1160 = v1158 & int32(255)
	v1161 = int32(1)
	v1165 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1152<<(uint(v1161)%32))+uint32(_c_F_seg_yylex[3]))))
	v1166 = v1160 + v1165
	v1171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1166<<(uint(v1161)%32))+uint32(_c_F_seg_yylex[4]))))
	if v1171 != v1152&int32(_a_F_seg_yylex_1) {
		v1137 = v1152
		v1139 = v1160
		v1140 = v1158
		goto L229
	} else {
		goto L234
	}
L234:
	;
	goto L230
L235:
	;
	goto L219
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_seg_yyparse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int64
	_ = v159
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 float32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 float32
	_ = v194
	var v195 float32
	_ = v195
	var v196 float32
	_ = v196
	var v198 float32
	_ = v198
	var v199 float32
	_ = v199
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 float32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 float32
	_ = v559
	var v561 float32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v572 float32
	_ = v572
	var v573 float32
	_ = v573
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v602 float32
	_ = v602
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v620 float32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v630 float32
	_ = v630
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 float32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 float32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v837 float32
	_ = v837
	var v841 int32
	_ = v841
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v921 int32
	_ = v921
	v23 = m.G0
	v25 = v23 - int32(1888)
	m.G0 = v25
	*(*int64)(unsafe.Add(mBase, uint32(v25)+1880)) = int64(0)
	v31 = v25 + int32(1680)
	v33 = v25 + int32(80)
	v39 = v33
	v40 = int32(0)
	v41 = v31
	v44 = v31
	v45 = int32(-2)
	v47 = int32(200)
	v49 = v33
	goto L4
L1:
	;
	if v25+int32(1680) != v906 {
		goto L248
	} else {
		goto L249
	}
L2:
	;
	v904 = int32(1)
	v906 = v74
	goto L1
L3:
	;
	F_seg_yyerror(m, l0, l1, l2, int32(_a_F_seg_yyparse_0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L13
	} else {
		goto L247
	}
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v40)
	if base.Ui32(v44+v47-int32(1)) <= base.Ui32(v41) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_seg_yyerror(m, l0, l1, l2, int32(_a_F_seg_yyparse_1))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L13
	} else {
		goto L246
	}
L6:
	;
	if int32(_a_F_seg_yyparse_2) < v47 {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	v107 = v39
	v108 = v41
	v110 = v44
	v112 = v47
	v113 = v49
	goto L8
L8:
	;
	if v40 == int32(8) {
		goto L27
	} else {
		goto L28
	}
L9:
	;
	v64 = int32(_a_F_seg_yyparse_3)
	v66 = v47 << (uint(int32(1)) % 32)
	if v64 <= v66 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v69 = v64
	goto L12
L11:
	;
	v69 = v66
	goto L12
L12:
	;
	v74 = F_palloc(m, v69*int32(9)+int32(7))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v74 == int32(0) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v80 = v41 - v44
	v82 = v80 + int32(1)
	if v82 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	base.MemoryCopy(m, v74, v44, v82)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v87 = base.I32_div_s(v69+int32(7), int32(8))
	v88 = int32(3)
	v90 = v74 + v87<<(uint(v88)%32)
	v92 = v82 << (uint(v88) % 32)
	if v92 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	base.MemoryCopy(m, v90, v49, v92)
	goto L21
L20:
	;
	goto L21
L21:
	;
	if v25+int32(1680) != v44 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_pfree(m, v44)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L13
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v69-int32(1) <= v80 {
		goto L2
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	v107 = v90 + v92 - int32(8)
	v108 = v74 + v80
	v110 = v74
	v112 = v69
	v113 = v90
	goto L8
L27:
	;
	v904 = int32(0)
	v906 = v110
	goto L1
L28:
	;
	goto L29
L29:
	;
	v117 = int32(1)
	v119 = v117 << (uint(v40) % 32)
	if v119&int32(_a_F_seg_yyparse_4) != 0 {
		v167 = v45
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L5
L31:
	;
	v39 = v874
	v40 = v875
	v41 = v876 + int32(1)
	v44 = v110
	v45 = v878
	v47 = v112
	v49 = v113
	goto L4
L32:
	;
	v181 = int32(*(*int8)(unsafe.Add(mBase, uint32(v178)+uint32(_c_F_seg_yyparse[0]))))
	v185 = (v117-v181)<<(uint(int32(3))%32) + v107
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+6)))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+5)))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+4)))
	v189 = *(*float32)(unsafe.Add(mBase, uint32(v185)))
	switch v178 - int32(2) {
	case 0:
		goto L60
	case 1:
		goto L59
	case 2:
		goto L58
	case 3:
		goto L57
	case 4:
		goto L56
	case 5, 7:
		goto L54
	case 6:
		goto L55
	default:
		v832 = v187
		v833 = v188
		v837 = v189
		goto L53
	}
L33:
	;
	if v119&int32(1053) != 0 {
		goto L30
	} else {
		goto L52
	}
L34:
	;
	v124 = int32(*(*int8)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_seg_yyparse[1]))))
	if v45 == int32(-2) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v147 = v124 + v146
	if base.Ui32(int32(12)) < base.Ui32(v147) {
		v167 = v145
		goto L33
	} else {
		goto L47
	}
L36:
	;
	v129 = F_seg_yylex(m, v25+int32(1880), l2)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L13
	} else {
		goto L39
	}
L37:
	;
	v131 = v45
	goto L38
L38:
	;
	if v131 <= int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v131 = v129
	goto L38
L40:
	;
	v134 = int32(0)
	v145 = v134
	v146 = v134
	goto L35
L41:
	;
	goto L42
L42:
	;
	if v131 == int32(256) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v904 = int32(1)
	v906 = v110
	goto L1
L44:
	;
	goto L45
L45:
	;
	if base.Ui32(int32(261)) < base.Ui32(v131) {
		v145 = v131
		v146 = int32(2)
		goto L35
	} else {
		goto L46
	}
L46:
	;
	v144 = int32(*(*int8)(unsafe.Add(mBase, uint32(v131)+uint32(_c_F_seg_yyparse[2]))))
	v145 = v131
	v146 = v144
	goto L35
L47:
	;
	v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v147)+uint32(_c_F_seg_yyparse[3]))))
	if v146 != v152 {
		v167 = v145
		goto L33
	} else {
		goto L48
	}
L48:
	;
	v156 = int32(*(*int8)(unsafe.Add(mBase, uint32(v147)+uint32(_c_F_seg_yyparse[4]))))
	if v147 != int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v25)+1880))
	*(*int64)(unsafe.Add(mBase, uint32(v107)+8)) = v159
	v874 = v107 + int32(8)
	v875 = v156
	v876 = v108
	v878 = int32(-2)
	goto L31
L50:
	;
	goto L51
L51:
	;
	v176 = v145
	v178 = int32(0) - v156
	goto L32
L52:
	;
	v173 = int32(*(*int8)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_seg_yyparse[5]))))
	v176 = v167
	v178 = v173
	goto L32
L53:
	;
	v841 = v107 - v181<<(uint(int32(3))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v841)+14)) = uint16(v186)
	*(*uint8)(unsafe.Add(mBase, uint32(v841)+13)) = uint8(v832)
	*(*uint8)(unsafe.Add(mBase, uint32(v841)+12)) = uint8(v833)
	*(*float32)(unsafe.Add(mBase, uint32(v841)+8)) = v837
	v847 = v841 + int32(8)
	v848 = v108 - v181
	v849 = int32(*(*int8)(unsafe.Add(mBase, uint32(v848))))
	v852 = int32(*(*int8)(unsafe.Add(mBase, uint32(v178)+uint32(_c_F_seg_yyparse[6]))))
	v854 = v852 - int32(7)
	v857 = int32(*(*int8)(unsafe.Add(mBase, uint32(v854)+uint32(_c_F_seg_yyparse[7]))))
	v858 = v849 + v857
	if base.Ui32(int32(12)) < base.Ui32(v858) {
		goto L243
	} else {
		goto L244
	}
L54:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v738 = F_float4in_internal(m, v736, int32(_a_F_seg_yyparse_5), v736, l1)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L13
	} else {
		goto L213
	}
L55:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v641 = F_float4in_internal(m, v639, int32(_a_F_seg_yyparse_5), v639, l1)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L13
	} else {
		goto L183
	}
L56:
	;
	v630 = *(*float32)(unsafe.Add(mBase, uint32(v107)))
	*(*float32)(unsafe.Add(mBase, uint32(l0))) = v630
	*(*float32)(unsafe.Add(mBase, uint32(l0)+4)) = v630
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v633)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v633)
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v636)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v636)
	v832 = v187
	v833 = v188
	v837 = v189
	goto L53
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(-8388608)
	v620 = *(*float32)(unsafe.Add(mBase, uint32(v107)))
	v621 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v621)
	*(*float32)(unsafe.Add(mBase, uint32(l0)+4)) = v620
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+5)))
	v625 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v625)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v624)
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v628)
	v832 = v187
	v833 = v188
	v837 = v189
	goto L53
L58:
	;
	v602 = *(*float32)(unsafe.Add(mBase, uint32(v107-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2139095040)
	*(*float32)(unsafe.Add(mBase, uint32(l0))) = v602
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107-int32(3)))))
	v609 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v609)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v608)
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107-int32(4)))))
	v615 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v615)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v614)
	v832 = v187
	v833 = v188
	v837 = v189
	goto L53
L59:
	;
	v559 = *(*float32)(unsafe.Add(mBase, uint32(v107-int32(16))))
	*(*float32)(unsafe.Add(mBase, uint32(l0))) = v559
	v561 = *(*float32)(unsafe.Add(mBase, uint32(v107)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+4)) = v561
	if base.F32_gt(v559, v561) != 0 {
		goto L175
	} else {
		goto L176
	}
L60:
	;
	v192 = int32(16)
	v193 = v107 - v192
	v194 = *(*float32)(unsafe.Add(mBase, uint32(v193)))
	v195 = *(*float32)(unsafe.Add(mBase, uint32(v107)))
	v196 = base.F32_sub(v194, v195)
	*(*float32)(unsafe.Add(mBase, uint32(l0))) = v196
	v198 = *(*float32)(unsafe.Add(mBase, uint32(v193)))
	v199 = *(*float32)(unsafe.Add(mBase, uint32(v107)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+4)) = base.F32_add(v198, v199)
	*(*float64)(unsafe.Add(mBase, uint32(v25)+16)) = base.F64_promote_f32(v196)
	v205 = v25 + int32(48)
	v210 = F_pg_snprintf(m, v205, int32(25), int32(_a_F_seg_yyparse_6), v25+v192)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	v218 = v205
	goto L63
L62:
	;
	if int32(6) <= v285 {
		goto L84
	} else {
		goto L85
	}
L63:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	v225 = v223 - int32(43)
	v232 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v225))|base.B2i32(int32(1)<<(uint(v225)%32)&int32(37) == v232) == v232 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v239 = v218
	v240 = v223
	v242 = int32(1)
	goto L71
L65:
	;
	v218 = v218 + int32(1)
	goto L63
L66:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	goto L62
L69:
	;
	v285 = v242
	goto L68
L70:
	;
	v255 = v239
	v256 = v240
	v257 = int32(0)
	goto L77
L71:
	;
	switch v240 - int32(46) {
	case 0:
		v249 = v242
		goto L74
	case 1:
		goto L70
	case 2:
		goto L75
	default:
		goto L73
	}
L72:
	;
	if v240 == int32(0) {
		goto L69
	} else {
		goto L76
	}
L73:
	;
	goto L72
L74:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+1)))
	v239 = v239 + int32(1)
	v240 = v250
	v242 = v249
	goto L71
L75:
	;
	v249 = v242 + int32(1)
	goto L74
L76:
	;
	goto L70
L77:
	;
	v261 = base.B2i32(v256 != int32(46))
	if v261&base.B2i32(base.Ui32(int32(9)) < base.Ui32((v256-int32(48))&int32(255))) == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v277 != 0 {
		v285 = v277
		goto L68
	} else {
		goto L83
	}
L79:
	;
	v271 = v257 + v261
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+1)))
	if v272 != 0 {
		v255 = v255 + int32(1)
		v256 = v272
		v257 = v271
		goto L77
	} else {
		goto L82
	}
L80:
	;
	v277 = v257
	goto L81
L81:
	;
	goto L78
L82:
	;
	v277 = v271
	goto L81
L83:
	;
	goto L69
L84:
	;
	v290 = int32(6)
	goto L86
L85:
	;
	v290 = v285
	goto L86
L86:
	;
	v292 = v107 - int32(11)
	v293 = int32(*(*int8)(unsafe.Add(mBase, uint32(v292))))
	v294 = int32(*(*int8)(unsafe.Add(mBase, uint32(v107)+5)))
	if v294 < v293 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v296 = v293
	goto L89
L88:
	;
	v296 = v294
	goto L89
L89:
	;
	if v296 < v290 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v304 = v205
	goto L94
L91:
	;
	v378 = v296
	goto L92
L92:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v378)
	v380 = *(*float32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*float64)(unsafe.Add(mBase, uint32(v25))) = base.F64_promote_f32(v380)
	v384 = v25 + int32(48)
	v387 = F_pg_snprintf(m, v384, int32(25), int32(_a_F_seg_yyparse_6), v25)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L13
	} else {
		goto L118
	}
L93:
	;
	if int32(6) <= v371 {
		goto L115
	} else {
		goto L116
	}
L94:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	v311 = v309 - int32(43)
	v318 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v311))|base.B2i32(int32(1)<<(uint(v311)%32)&int32(37) == v318) == v318 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v325 = v304
	v326 = v309
	v328 = int32(1)
	goto L102
L96:
	;
	v304 = v304 + int32(1)
	goto L94
L97:
	;
	goto L98
L98:
	;
	goto L95
L99:
	;
	goto L93
L100:
	;
	v371 = v328
	goto L99
L101:
	;
	v341 = v325
	v342 = v326
	v343 = int32(0)
	goto L108
L102:
	;
	switch v326 - int32(46) {
	case 0:
		v335 = v328
		goto L105
	case 1:
		goto L101
	case 2:
		goto L106
	default:
		goto L104
	}
L103:
	;
	if v326 == int32(0) {
		goto L100
	} else {
		goto L107
	}
L104:
	;
	goto L103
L105:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+1)))
	v325 = v325 + int32(1)
	v326 = v336
	v328 = v335
	goto L102
L106:
	;
	v335 = v328 + int32(1)
	goto L105
L107:
	;
	goto L101
L108:
	;
	v347 = base.B2i32(v342 != int32(46))
	if v347&base.B2i32(base.Ui32(int32(9)) < base.Ui32((v342-int32(48))&int32(255))) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if v363 != 0 {
		v371 = v363
		goto L99
	} else {
		goto L114
	}
L110:
	;
	v357 = v343 + v347
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+1)))
	if v358 != 0 {
		v341 = v341 + int32(1)
		v342 = v358
		v343 = v357
		goto L108
	} else {
		goto L113
	}
L111:
	;
	v363 = v343
	goto L112
L112:
	;
	goto L109
L113:
	;
	v363 = v357
	goto L112
L114:
	;
	goto L100
L115:
	;
	v376 = int32(6)
	goto L117
L116:
	;
	v376 = v371
	goto L117
L117:
	;
	v378 = v376
	goto L92
L118:
	;
	v395 = v384
	goto L120
L119:
	;
	if int32(6) <= v462 {
		goto L141
	} else {
		goto L142
	}
L120:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	v402 = v400 - int32(43)
	v409 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v402))|base.B2i32(int32(1)<<(uint(v402)%32)&int32(37) == v409) == v409 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v416 = v395
	v417 = v400
	v419 = int32(1)
	goto L128
L122:
	;
	v395 = v395 + int32(1)
	goto L120
L123:
	;
	goto L124
L124:
	;
	goto L121
L125:
	;
	goto L119
L126:
	;
	v462 = v419
	goto L125
L127:
	;
	v432 = v416
	v433 = v417
	v434 = int32(0)
	goto L134
L128:
	;
	switch v417 - int32(46) {
	case 0:
		v426 = v419
		goto L131
	case 1:
		goto L127
	case 2:
		goto L132
	default:
		goto L130
	}
L129:
	;
	if v417 == int32(0) {
		goto L126
	} else {
		goto L133
	}
L130:
	;
	goto L129
L131:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+1)))
	v416 = v416 + int32(1)
	v417 = v427
	v419 = v426
	goto L128
L132:
	;
	v426 = v419 + int32(1)
	goto L131
L133:
	;
	goto L127
L134:
	;
	v438 = base.B2i32(v433 != int32(46))
	if v438&base.B2i32(base.Ui32(int32(9)) < base.Ui32((v433-int32(48))&int32(255))) == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	if v454 != 0 {
		v462 = v454
		goto L125
	} else {
		goto L140
	}
L136:
	;
	v448 = v434 + v438
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432)+1)))
	if v449 != 0 {
		v432 = v432 + int32(1)
		v433 = v449
		v434 = v448
		goto L134
	} else {
		goto L139
	}
L137:
	;
	v454 = v434
	goto L138
L138:
	;
	goto L135
L139:
	;
	v454 = v448
	goto L138
L140:
	;
	goto L126
L141:
	;
	v467 = int32(6)
	goto L143
L142:
	;
	v467 = v462
	goto L143
L143:
	;
	v468 = int32(*(*int8)(unsafe.Add(mBase, uint32(v292))))
	v469 = int32(*(*int8)(unsafe.Add(mBase, uint32(v107)+5)))
	if v469 < v468 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v471 = v468
	goto L146
L145:
	;
	v471 = v469
	goto L146
L146:
	;
	if v471 < v467 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v479 = v384
	goto L151
L148:
	;
	v553 = v471
	goto L149
L149:
	;
	v554 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v554)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v553)
	v832 = v187
	v833 = v188
	v837 = v189
	goto L53
L150:
	;
	if int32(6) <= v546 {
		goto L172
	} else {
		goto L173
	}
L151:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	v486 = v484 - int32(43)
	v493 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v486))|base.B2i32(int32(1)<<(uint(v486)%32)&int32(37) == v493) == v493 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v500 = v479
	v501 = v484
	v503 = int32(1)
	goto L159
L153:
	;
	v479 = v479 + int32(1)
	goto L151
L154:
	;
	goto L155
L155:
	;
	goto L152
L156:
	;
	goto L150
L157:
	;
	v546 = v503
	goto L156
L158:
	;
	v516 = v500
	v517 = v501
	v518 = int32(0)
	goto L165
L159:
	;
	switch v501 - int32(46) {
	case 0:
		v510 = v503
		goto L162
	case 1:
		goto L158
	case 2:
		goto L163
	default:
		goto L161
	}
L160:
	;
	if v501 == int32(0) {
		goto L157
	} else {
		goto L164
	}
L161:
	;
	goto L160
L162:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+1)))
	v500 = v500 + int32(1)
	v501 = v511
	v503 = v510
	goto L159
L163:
	;
	v510 = v503 + int32(1)
	goto L162
L164:
	;
	goto L158
L165:
	;
	v522 = base.B2i32(v517 != int32(46))
	if v522&base.B2i32(base.Ui32(int32(9)) < base.Ui32((v517-int32(48))&int32(255))) == int32(0) {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	if v538 != 0 {
		v546 = v538
		goto L156
	} else {
		goto L171
	}
L167:
	;
	v532 = v518 + v522
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
	if v533 != 0 {
		v516 = v516 + int32(1)
		v517 = v533
		v518 = v532
		goto L165
	} else {
		goto L170
	}
L168:
	;
	v538 = v518
	goto L169
L169:
	;
	goto L166
L170:
	;
	v538 = v532
	goto L169
L171:
	;
	goto L157
L172:
	;
	v551 = int32(6)
	goto L174
L173:
	;
	v551 = v546
	goto L174
L174:
	;
	v553 = v551
	goto L149
L175:
	;
	v564 = int32(1)
	v565 = F_errsave_start(m, l1)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L13
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107-int32(11)))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v590)
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v592)
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107-int32(12)))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v596)
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v598)
	v832 = v187
	v833 = v188
	v837 = v189
	goto L53
L178:
	;
	if v565 == int32(0) {
		v904 = v564
		v906 = v110
		goto L1
	} else {
		goto L179
	}
L179:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L13
	} else {
		goto L180
	}
L180:
	;
	v572 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v573 = *(*float32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+40)) = base.F64_promote_f32(v573)
	*(*float64)(unsafe.Add(mBase, uint32(v25)+32)) = base.F64_promote_f32(v572)
	F_errmsg(m, int32(_a_F_seg_yyparse_7), v25+int32(32))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L13
	} else {
		goto L181
	}
L181:
	;
	F_errsave_finish(m, l1, int32(_a_F_seg_yyparse_8), int32(86), int32(_a_F_seg_yyparse_9))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L13
	} else {
		goto L182
	}
L182:
	;
	v904 = v564
	v906 = v110
	goto L1
L183:
	;
	if l1 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v107-int32(8))))
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654))))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v663 = v657
	goto L189
L185:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v645 != int32(447) {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v648 == int32(0) {
		goto L184
	} else {
		goto L187
	}
L187:
	;
	v904 = int32(1)
	v906 = v110
	goto L1
L188:
	;
	if int32(6) <= v730 {
		goto L210
	} else {
		goto L211
	}
L189:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663))))
	v670 = v668 - int32(43)
	v677 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v670))|base.B2i32(int32(1)<<(uint(v670)%32)&int32(37) == v677) == v677 {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v684 = v663
	v685 = v668
	v687 = int32(1)
	goto L197
L191:
	;
	v663 = v663 + int32(1)
	goto L189
L192:
	;
	goto L193
L193:
	;
	goto L190
L194:
	;
	goto L188
L195:
	;
	v730 = v687
	goto L194
L196:
	;
	v700 = v684
	v701 = v685
	v702 = int32(0)
	goto L203
L197:
	;
	switch v685 - int32(46) {
	case 0:
		v694 = v687
		goto L200
	case 1:
		goto L196
	case 2:
		goto L201
	default:
		goto L199
	}
L198:
	;
	if v685 == int32(0) {
		goto L195
	} else {
		goto L202
	}
L199:
	;
	goto L198
L200:
	;
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v684)+1)))
	v684 = v684 + int32(1)
	v685 = v695
	v687 = v694
	goto L197
L201:
	;
	v694 = v687 + int32(1)
	goto L200
L202:
	;
	goto L196
L203:
	;
	v706 = base.B2i32(v701 != int32(46))
	if v706&base.B2i32(base.Ui32(int32(9)) < base.Ui32((v701-int32(48))&int32(255))) == int32(0) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	if v722 != 0 {
		v730 = v722
		goto L194
	} else {
		goto L209
	}
L205:
	;
	v716 = v702 + v706
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700)+1)))
	if v717 != 0 {
		v700 = v700 + int32(1)
		v701 = v717
		v702 = v716
		goto L203
	} else {
		goto L208
	}
L206:
	;
	v722 = v702
	goto L207
L207:
	;
	goto L204
L208:
	;
	v722 = v716
	goto L207
L209:
	;
	goto L195
L210:
	;
	v735 = int32(6)
	goto L212
L211:
	;
	v735 = v730
	goto L212
L212:
	;
	v832 = v735
	v833 = v655
	v837 = v641
	goto L53
L213:
	;
	if l1 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v756 = v750
	goto L219
L215:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v742 != int32(447) {
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v745 == int32(0) {
		goto L214
	} else {
		goto L217
	}
L217:
	;
	v904 = int32(1)
	v906 = v110
	goto L1
L218:
	;
	if int32(6) <= v823 {
		goto L240
	} else {
		goto L241
	}
L219:
	;
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756))))
	v763 = v761 - int32(43)
	v770 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v763))|base.B2i32(int32(1)<<(uint(v763)%32)&int32(37) == v770) == v770 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v777 = v756
	v778 = v761
	v780 = int32(1)
	goto L227
L221:
	;
	v756 = v756 + int32(1)
	goto L219
L222:
	;
	goto L223
L223:
	;
	goto L220
L224:
	;
	goto L218
L225:
	;
	v823 = v780
	goto L224
L226:
	;
	v793 = v777
	v794 = v778
	v795 = int32(0)
	goto L233
L227:
	;
	switch v778 - int32(46) {
	case 0:
		v787 = v780
		goto L230
	case 1:
		goto L226
	case 2:
		goto L231
	default:
		goto L229
	}
L228:
	;
	if v778 == int32(0) {
		goto L225
	} else {
		goto L232
	}
L229:
	;
	goto L228
L230:
	;
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777)+1)))
	v777 = v777 + int32(1)
	v778 = v788
	v780 = v787
	goto L227
L231:
	;
	v787 = v780 + int32(1)
	goto L230
L232:
	;
	goto L226
L233:
	;
	v799 = base.B2i32(v794 != int32(46))
	if v799&base.B2i32(base.Ui32(int32(9)) < base.Ui32((v794-int32(48))&int32(255))) == int32(0) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	if v815 != 0 {
		v823 = v815
		goto L224
	} else {
		goto L239
	}
L235:
	;
	v809 = v795 + v799
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793)+1)))
	if v810 != 0 {
		v793 = v793 + int32(1)
		v794 = v810
		v795 = v809
		goto L233
	} else {
		goto L238
	}
L236:
	;
	v815 = v795
	goto L237
L237:
	;
	goto L234
L238:
	;
	v815 = v809
	goto L237
L239:
	;
	goto L225
L240:
	;
	v828 = int32(6)
	goto L242
L241:
	;
	v828 = v823
	goto L242
L242:
	;
	v832 = v828
	v833 = int32(0)
	v837 = v738
	goto L53
L243:
	;
	v872 = int32(*(*int8)(unsafe.Add(mBase, uint32(v854)+uint32(_c_F_seg_yyparse[8]))))
	v874 = v847
	v875 = v872
	v876 = v848
	v878 = v176
	goto L31
L244:
	;
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v858)+uint32(_c_F_seg_yyparse[3]))))
	if v863 != v849&int32(255) {
		goto L243
	} else {
		goto L245
	}
L245:
	;
	v869 = int32(*(*int8)(unsafe.Add(mBase, uint32(v858)+uint32(_c_F_seg_yyparse[4]))))
	v874 = v847
	v875 = v869
	v876 = v848
	v878 = v176
	goto L31
L246:
	;
	v904 = int32(1)
	v906 = v110
	goto L1
L247:
	;
	v904 = int32(2)
	v906 = v44
	goto L1
L248:
	;
	F_pfree(m, v906)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L13
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	m.G0 = v25 + int32(1888)
	return v904
L251:
	;
	goto L250
}
