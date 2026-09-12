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
	v6 = F_DirectFunctionCall2Coll(m, int32(6675), int32(0), v4, v5)
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
	var v74 int32
	_ = v74
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
	var v90 int32
	_ = v90
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
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
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
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
	var v391 int32
	_ = v391
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
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
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1157 int32
	_ = v1157
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1223 int32
	_ = v1223
	var v1229 int32
	_ = v1229
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
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
	v74 = v68
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
	v28 = *(*int32)(unsafe.Add(mBase, _consts[718]))
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
	v34 = *(*int32)(unsafe.Add(mBase, _consts[719]))
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
	v49 = F_seg_yy_create_buffer(m, v47, int32(16384), l1)
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
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v82)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v70)+44))
	v86 = v70
	v87 = v84
	v90 = v74
	v91 = v74
	v97 = v81
	goto L24
L24:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+uint32(_consts[1347]))))
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
	v114 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87<<(uint(v110)%32))+uint32(_consts[1348]))))
	v115 = v114 + v101
	v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v115<<(uint(v110)%32))+uint32(_consts[1349]))))
	if v120 != v87 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v124 = v87
	v125 = v101
	v126 = v101
	goto L32
L30:
	;
	v162 = v115
	goto L31
L31:
	;
	v175 = int32(1)
	v179 = int32(*(*int16)(unsafe.Add(mBase, uint32(v162<<(uint(v175)%32))+uint32(_consts[1350]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v162))%64)&int64(72054295503044608) == int64(0) {
		v87 = v179
		v91 = v91 + v175
		goto L24
	} else {
		goto L38
	}
L32:
	;
	v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(v124<<(uint(int32(1))%32))+uint32(_consts[1351]))))
	if v124 == int32(21) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v162 = v155
	goto L31
L34:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+uint32(_consts[1352]))))
	v145 = v144
	goto L36
L35:
	;
	v145 = v126
	goto L36
L36:
	;
	v149 = v145 & int32(255)
	v150 = int32(1)
	v154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v139<<(uint(v150)%32))+uint32(_consts[1348]))))
	v155 = v149 + v154
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155<<(uint(v150)%32))+uint32(_consts[1349]))))
	if v139&int32(65535) != v160 {
		v124 = v139
		v125 = v149
		v126 = v145
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v198 = v90
	goto L39
L39:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v86)+64))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v86)+68))
	v206 = v202
	v212 = v203
	v213 = v198
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+80)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v86)+32)) = v212 - v213
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+24)) = uint8(v220)
	v222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v212))) = uint8(v222)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v212
	v229 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206<<(uint(int32(1))%32))+uint32(_consts[1353]))))
	v230 = v229
	v235 = v212
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
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v1229
	*(*int32)(unsafe.Add(mBase, uint32(v86)+48)) = int32(0)
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	v1244 = base.I32_div_s(v1240-int32(1), int32(2))
	v230 = v1244 + int32(11)
	v235 = v1229
	goto L43
L46:
	;
	F_yy_fatal_error_7(m, int32(32817))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L19
	} else {
		goto L253
	}
L47:
	;
	F_yy_fatal_error_7(m, int32(715354))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L19
	} else {
		goto L252
	}
L48:
	;
	F_yy_fatal_error_7(m, int32(475694))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L19
	} else {
		goto L251
	}
L49:
	;
	F_yy_fatal_error_7(m, int32(471350))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L19
	} else {
		goto L250
	}
L50:
	;
	return v1208
L51:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1204))) = v1205
	v1208 = int32(259)
	goto L50
L52:
	;
	F_yy_fatal_error_7(m, int32(444628))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L19
	} else {
		goto L249
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
		v74 = v235
		v81 = v97
		goto L22
	case 8:
		goto L57
	case 9:
		goto L56
	default:
		goto L52
	case 11:
		v1208 = v97
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
	F_yy_fatal_error_7(m, int32(475046))
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
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = int32(4101)
	return int32(261)
L59:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = int32(573111)
	return int32(261)
L60:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(573169)
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
	v198 = v213
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
	if v306 != 0 {
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
		goto L104
	}
L72:
	;
	v312 = v309
	v316 = v303
	goto L75
L73:
	;
	v414 = v309
	goto L74
L74:
	;
	if int32(1)<<(uint(v414)%32)&int32(619905031) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L75:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316))))
	if v323 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v414 = v408
	goto L74
L77:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323)+uint32(_consts[1347]))))
	v328 = v326
	goto L79
L78:
	;
	v328 = int32(1)
	goto L79
L79:
	;
	if int32(1)<<(uint(v312)%32)&int32(619905031) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v312
	goto L82
L81:
	;
	goto L82
L82:
	;
	v338 = v328 & int32(255)
	v339 = int32(1)
	v343 = int32(*(*int16)(unsafe.Add(mBase, uint32(v312<<(uint(v339)%32))+uint32(_consts[1348]))))
	v344 = v338 + v343
	v349 = int32(*(*int16)(unsafe.Add(mBase, uint32(v344<<(uint(v339)%32))+uint32(_consts[1349]))))
	if v349 != v312 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v353 = v312
	v354 = v338
	v355 = v328
	goto L86
L84:
	;
	v391 = v344
	goto L85
L85:
	;
	v404 = int32(1)
	v408 = int32(*(*int16)(unsafe.Add(mBase, uint32(v391<<(uint(v404)%32))+uint32(_consts[1350]))))
	v410 = v316 + v404
	if v410 != v307 {
		v312 = v408
		v316 = v410
		goto L75
	} else {
		goto L92
	}
L86:
	;
	v368 = int32(*(*int16)(unsafe.Add(mBase, uint32(v353<<(uint(int32(1))%32))+uint32(_consts[1351]))))
	if v353 == int32(21) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v391 = v384
	goto L85
L88:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+uint32(_consts[1352]))))
	v374 = v373
	goto L90
L89:
	;
	v374 = v355
	goto L90
L90:
	;
	v378 = v374 & int32(255)
	v379 = int32(1)
	v383 = int32(*(*int16)(unsafe.Add(mBase, uint32(v368<<(uint(v379)%32))+uint32(_consts[1348]))))
	v384 = v378 + v383
	v389 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384<<(uint(v379)%32))+uint32(_consts[1349]))))
	if v368&int32(65535) != v389 {
		v353 = v368
		v354 = v378
		v355 = v374
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
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v414
	goto L95
L94:
	;
	goto L95
L95:
	;
	v433 = int32(1)
	v437 = int32(*(*int16)(unsafe.Add(mBase, uint32(v414<<(uint(v433)%32))+uint32(_consts[1348]))))
	v439 = v437 + v433
	v444 = int32(*(*int16)(unsafe.Add(mBase, uint32(v439<<(uint(v433)%32))+uint32(_consts[1349]))))
	if v444 != v414 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v448 = v414
	goto L99
L97:
	;
	v478 = v439
	goto L98
L98:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v478))%64)&int64(72054295503044608) != int64(0) {
		v198 = v303
		goto L39
	} else {
		goto L102
	}
L99:
	;
	v459 = int32(1)
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v448<<(uint(v459)%32))+uint32(_consts[1351]))))
	v464 = base.I32_extend16_s(v463)
	v469 = int32(*(*int16)(unsafe.Add(mBase, uint32(v464<<(uint(v459)%32))+uint32(_consts[1348]))))
	v471 = v469 + v459
	v476 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v471<<(uint(v459)%32))+uint32(_consts[1349]))))
	if v463 != v476 {
		v448 = v464
		goto L99
	} else {
		goto L101
	}
L100:
	;
	v478 = v471
	goto L98
L101:
	;
	goto L100
L102:
	;
	if v478 == int32(0) {
		v198 = v303
		goto L39
	} else {
		goto L103
	}
L103:
	;
	v500 = int32(1)
	v501 = v307 + v500
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v501
	v507 = int32(*(*int16)(unsafe.Add(mBase, uint32(v478<<(uint(v500)%32))+uint32(_consts[1350]))))
	v87 = v507
	v90 = v303
	v91 = v501
	goto L24
L104:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v286)+40))
	if v512 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	v1095 = v1082 + v1086
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v1095
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	if base.Ui32(v1095) <= base.Ui32(v1087) {
		v206 = v1097
		v212 = v1095
		v213 = v1087
		goto L41
	} else {
		goto L230
	}
L107:
	;
	if v299-v511 != int32(1) {
		v1082 = v300
		v1086 = v298
		v1087 = v511
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v520 = v511 ^ int32(-1) + v299
	if v520 != 0 {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	v1229 = v511
	goto L45
L111:
	;
	v914 = v900 + v520
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v902)+12))
	if base.Ui32(v914) <= base.Ui32(v915) {
		goto L198
	} else {
		goto L199
	}
L112:
	;
	if v520 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L113:
	;
	v521 = int32(7)
	v522 = v520 & v521
	if base.Ui32(v299-v511-int32(2)) < base.Ui32(v521) {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	v625 = v286
	v636 = v297
	goto L115
L115:
	;
	if v636 == int32(2) {
		goto L129
	} else {
		goto L130
	}
L116:
	;
	if v522 != 0 {
		goto L123
	} else {
		goto L124
	}
L117:
	;
	v567 = v300
	v569 = v511
	goto L116
L118:
	;
	goto L119
L119:
	;
	v531 = v300
	v533 = v511
	v534 = int32(0)
	goto L120
L120:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533))))
	*(*uint8)(unsafe.Add(mBase, uint32(v531))) = uint8(v544)
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v531)+1)) = uint8(v546)
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v531)+2)) = uint8(v548)
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v531)+3)) = uint8(v550)
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v531)+4)) = uint8(v552)
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v531)+5)) = uint8(v554)
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v531)+6)) = uint8(v556)
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v531)+7)) = uint8(v558)
	v560 = int32(8)
	v561 = v531 + v560
	v563 = v533 + v560
	v565 = v534 + v560
	if v565 != v520&int32(-8) {
		v531 = v561
		v533 = v563
		v534 = v565
		goto L120
	} else {
		goto L122
	}
L121:
	;
	v567 = v561
	v569 = v563
	goto L116
L122:
	;
	goto L121
L123:
	;
	v581 = v567
	v583 = v569
	v584 = int32(0)
	goto L126
L124:
	;
	goto L125
L125:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v616+v617<<(uint(int32(2))%32))))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v621)+44))
	v625 = v621
	v636 = v622
	goto L115
L126:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583))))
	*(*uint8)(unsafe.Add(mBase, uint32(v581))) = uint8(v594)
	v596 = int32(1)
	v601 = v584 + v596
	if v601 != v522 {
		v581 = v581 + v596
		v583 = v583 + v596
		v584 = v601
		goto L126
	} else {
		goto L128
	}
L127:
	;
	goto L125
L128:
	;
	goto L127
L129:
	;
	v639 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v625)+16)) = v639
	v872 = v625
	goto L112
L130:
	;
	goto L131
L131:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v625)+12))
	v644 = v511 - v299
	v645 = v643 + v644
	if v645 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v86)+36))
	v649 = v643
	v651 = v625
	v653 = v648
	goto L135
L133:
	;
	v700 = v625
	v701 = v645
	goto L134
L134:
	;
	v711 = int32(16777216)
	if base.Ui32(v711) <= base.Ui32(v701) {
		goto L151
	} else {
		goto L152
	}
L135:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v651)+20))
	if v662 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v700 = v693
	v701 = v695
	goto L134
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v651)+4)) = int32(0)
	goto L46
L138:
	;
	goto L139
L139:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v651)+4))
	if base.Ui32(int32(2147483646)) <= base.Ui32(v649) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v673 = int32(-3)
	goto L142
L141:
	;
	v673 = v649 << (uint(int32(1)) % 32)
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v651)+12)) = v673
	v676 = v673 + int32(2)
	if v667 != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v651)+4)) = v681
	if v681 == int32(0) {
		goto L46
	} else {
		goto L149
	}
L144:
	;
	v677 = F_repalloc(m, v667, v676)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L19
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v679 = F_palloc(m, v676)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L19
	} else {
		goto L148
	}
L147:
	;
	v681 = v677
	goto L143
L148:
	;
	v681 = v679
	goto L143
L149:
	;
	v686 = v681 + (v653 - v667)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v686
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v688+v689<<(uint(int32(2))%32))))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v693)+12))
	v695 = v694 + v644
	if v695 == int32(0) {
		v649 = v694
		v651 = v693
		v653 = v686
		goto L135
	} else {
		goto L150
	}
L150:
	;
	goto L136
L151:
	;
	v714 = v711
	goto L153
L152:
	;
	v714 = v701
	goto L153
L153:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v700)+24))
	if v716 != 0 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v860
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v862+v863<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v867)+16)) = v860
	if v860 != 0 {
		v900 = v860
		v902 = v867
		v913 = int32(0)
		goto L111
	} else {
		goto L192
	}
L155:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v847+v848<<(uint(int32(2))%32))))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v852)+4))
	v856 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v853+v520+v748))) = uint8(v856)
	v860 = v748 + int32(1)
	goto L154
L156:
	;
	v717 = int32(0)
	goto L160
L157:
	;
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v700)+4))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v772 = F_fread(m, v768+v520, int32(1), v714, v771)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L19
	} else {
		goto L173
	}
L159:
	;
	switch v734 {
	case 0:
		goto L165
	default:
		v860 = v748
		goto L154
	case 11:
		goto L155
	}
L160:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v731 = F_do_getc(m, v730)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L19
	} else {
		goto L163
	}
L161:
	;
	v748 = v714
	goto L159
L162:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v735+v736<<(uint(int32(2))%32))))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v740)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v741+v520+v717))) = uint8(v731)
	v746 = v717 + int32(1)
	if v746 != v714 {
		v717 = v746
		goto L160
	} else {
		goto L164
	}
L163:
	;
	v734 = v731 + int32(1)
	switch v734 {
	case 0, 11:
		v748 = v717
		goto L159
	default:
		goto L162
	}
L164:
	;
	goto L161
L165:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v749)+76))
	if v750 < int32(0) {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	if int32(base.Ui32(v755)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v860 = v748
		goto L154
	} else {
		goto L171
	}
L167:
	;
	goto L166
L168:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v749)))
	v755 = v753
	goto L167
L169:
	;
	goto L170
L170:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v749)))
	v755 = v754
	goto L167
L171:
	;
	F_yy_fatal_error_7(m, int32(475694))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L19
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	v774 = v772
	goto L174
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v774
	if v774 != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v839+v840<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v844)+16)) = v774
	v900 = v774
	v902 = v844
	v913 = int32(0)
	goto L111
L176:
	;
	goto L175
L177:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v788)+76))
	if v789 < int32(0) {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	if int32(base.Ui32(v794)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L183
	} else {
		goto L184
	}
L179:
	;
	goto L178
L180:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v788)))
	v794 = v792
	goto L179
L181:
	;
	goto L182
L182:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v788)))
	v794 = v793
	goto L179
L183:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v801+v802<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v806)+16)) = int32(0)
	v872 = v806
	goto L112
L184:
	;
	goto L185
L185:
	;
	v810 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v810 != int32(27) {
		goto L48
	} else {
		goto L186
	}
L186:
	;
	v814 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v814
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v788)+76))
	if v814 <= v816 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v827+v828<<(uint(int32(2))%32))))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v832)+4))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v837 = F_fread(m, v833+v520, int32(1), v714, v836)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L19
	} else {
		goto L191
	}
L188:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v788)))
	*(*int32)(unsafe.Add(mBase, uint32(v788))) = v819 & int32(-49)
	goto L187
L189:
	;
	goto L190
L190:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v788)))
	*(*int32)(unsafe.Add(mBase, uint32(v788))) = v823 & int32(-49)
	goto L187
L191:
	;
	v774 = v837
	goto L174
L192:
	;
	v872 = v867
	goto L112
L193:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	F_seg_yyrestart(m, v885, v86)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L19
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v896 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v872)+44)) = v896
	v900 = int32(0)
	v902 = v872
	v913 = v896
	goto L111
L196:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v888+v889<<(uint(int32(2))%32))))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v900 = v894
	v902 = v893
	v913 = int32(1)
	goto L111
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v943
	v946 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v941+v943))) = uint8(v946)
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v950 = int32(2)
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v948+v949<<(uint(v950)%32))))
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v953)+4))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v954+v955)+1)) = uint8(v946)
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v959+v960<<(uint(v950)%32))))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v964)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+80)) = v965
	if v913 == int32(1) {
		v1229 = v965
		goto L45
	} else {
		goto L208
	}
L198:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v902)+4))
	v941 = v917
	v943 = v914
	goto L197
L199:
	;
	goto L200
L200:
	;
	v920 = v914 + int32(base.Ui32(v900)>>(uint(int32(1))%32))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v902)+4))
	if v921 != 0 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v927+v928<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v932)+4)) = v926
	if v926 == int32(0) {
		goto L47
	} else {
		goto L207
	}
L202:
	;
	v922 = F_repalloc(m, v921, v920)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L19
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v924 = F_palloc(m, v920)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L19
	} else {
		goto L206
	}
L205:
	;
	v926 = v922
	goto L201
L206:
	;
	v926 = v924
	goto L201
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v932)+12)) = v920 - int32(2)
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v941 = v926
	v943 = v939 + v520
	goto L197
L208:
	;
	switch v913 - int32(1) {
	case 0:
		goto L105
	case 1:
		goto L209
	default:
		goto L210
	}
L209:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v1082 = v965
	v1086 = v1081
	v1087 = v965
	goto L106
L210:
	;
	v973 = v278 ^ int32(-1) + v212
	v974 = v965 + v973
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v974
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	if v973 == int32(0) {
		v87 = v976
		v90 = v965
		v91 = v974
		goto L24
	} else {
		goto L211
	}
L211:
	;
	v981 = v976
	v986 = v965
	goto L212
L212:
	;
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986))))
	if v992 != 0 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v87 = v1077
	v90 = v965
	v91 = v974
	goto L24
L214:
	;
	v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v992)+uint32(_consts[1347]))))
	v997 = v995
	goto L216
L215:
	;
	v997 = int32(1)
	goto L216
L216:
	;
	if int32(1)<<(uint(v981)%32)&int32(619905031) == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v981
	goto L219
L218:
	;
	goto L219
L219:
	;
	v1007 = v997 & int32(255)
	v1008 = int32(1)
	v1012 = int32(*(*int16)(unsafe.Add(mBase, uint32(v981<<(uint(v1008)%32))+uint32(_consts[1348]))))
	v1013 = v1007 + v1012
	v1018 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1013<<(uint(v1008)%32))+uint32(_consts[1349]))))
	if v1018 != v981 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v1022 = v981
	v1023 = v1007
	v1024 = v997
	goto L223
L221:
	;
	v1060 = v1013
	goto L222
L222:
	;
	v1073 = int32(1)
	v1077 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1060<<(uint(v1073)%32))+uint32(_consts[1350]))))
	v1079 = v986 + v1073
	if v974 != v1079 {
		v981 = v1077
		v986 = v1079
		goto L212
	} else {
		goto L229
	}
L223:
	;
	v1037 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1022<<(uint(int32(1))%32))+uint32(_consts[1351]))))
	if v1022 == int32(21) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v1060 = v1053
	goto L222
L225:
	;
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1023)+uint32(_consts[1352]))))
	v1043 = v1042
	goto L227
L226:
	;
	v1043 = v1024
	goto L227
L227:
	;
	v1047 = v1043 & int32(255)
	v1048 = int32(1)
	v1052 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1037<<(uint(v1048)%32))+uint32(_consts[1348]))))
	v1053 = v1047 + v1052
	v1058 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1053<<(uint(v1048)%32))+uint32(_consts[1349]))))
	if v1037&int32(65535) != v1058 {
		v1022 = v1037
		v1023 = v1047
		v1024 = v1043
		goto L223
	} else {
		goto L228
	}
L228:
	;
	goto L224
L229:
	;
	goto L213
L230:
	;
	v1101 = v1097
	v1105 = v1087
	goto L231
L231:
	;
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1105))))
	if v1112 != 0 {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v206 = v1197
	v212 = v1095
	v213 = v1087
	goto L41
L233:
	;
	v1115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1112)+uint32(_consts[1347]))))
	v1117 = v1115
	goto L235
L234:
	;
	v1117 = int32(1)
	goto L235
L235:
	;
	if int32(1)<<(uint(v1101)%32)&int32(619905031) == int32(0) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v1101
	goto L238
L237:
	;
	goto L238
L238:
	;
	v1127 = v1117 & int32(255)
	v1128 = int32(1)
	v1132 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1101<<(uint(v1128)%32))+uint32(_consts[1348]))))
	v1133 = v1127 + v1132
	v1138 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1133<<(uint(v1128)%32))+uint32(_consts[1349]))))
	if v1138 != v1101 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1142 = v1101
	v1143 = v1127
	v1144 = v1117
	goto L242
L240:
	;
	v1180 = v1133
	goto L241
L241:
	;
	v1193 = int32(1)
	v1197 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1180<<(uint(v1193)%32))+uint32(_consts[1350]))))
	v1199 = v1105 + v1193
	if v1199 != v1095 {
		v1101 = v1197
		v1105 = v1199
		goto L231
	} else {
		goto L248
	}
L242:
	;
	v1157 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1142<<(uint(int32(1))%32))+uint32(_consts[1351]))))
	if v1142 == int32(21) {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	v1180 = v1173
	goto L241
L244:
	;
	v1162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1143)+uint32(_consts[1352]))))
	v1163 = v1162
	goto L246
L245:
	;
	v1163 = v1144
	goto L246
L246:
	;
	v1167 = v1163 & int32(255)
	v1168 = int32(1)
	v1172 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1157<<(uint(v1168)%32))+uint32(_consts[1348]))))
	v1173 = v1167 + v1172
	v1178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1173<<(uint(v1168)%32))+uint32(_consts[1349]))))
	if v1157&int32(65535) != v1178 {
		v1142 = v1157
		v1143 = v1167
		v1144 = v1163
		goto L242
	} else {
		goto L247
	}
L247:
	;
	goto L243
L248:
	;
	goto L232
L249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_seg_yyparse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
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
	var v157 int64
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
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
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 float32
	_ = v364
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 float32
	_ = v526
	var v528 float32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v539 float32
	_ = v539
	var v540 float32
	_ = v540
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v573 float32
	_ = v573
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v591 float32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v601 float32
	_ = v601
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v612 float32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 float32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 float32
	_ = v782
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
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
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v865 float32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v874 float32
	_ = v874
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v927 int32
	_ = v927
	var v936 int32
	_ = v936
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v963 int32
	_ = v963
	v22 = m.G0
	v24 = v22 - int32(2096)
	m.G0 = v24
	v28 = v24 + int32(1680)
	v30 = v24 + int32(80)
	v35 = int32(0)
	v36 = v30
	v39 = v28
	v42 = v28
	v43 = int32(-2)
	v45 = int32(200)
	v47 = v30
	goto L3
L1:
	;
	if v24+int32(1680) != v950 {
		goto L291
	} else {
		goto L292
	}
L2:
	;
	F_seg_yyerror(m, l0, l1, l2, int32(462179))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L12
	} else {
		goto L290
	}
L3:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v39))) = uint16(v35)
	v55 = v45 << (uint(int32(1)) % 32)
	if base.Ui32(v42+v55-int32(2)) <= base.Ui32(v39) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_seg_yyerror(m, l0, l1, l2, int32(223326))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L12
	} else {
		goto L289
	}
L5:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v45) {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	v111 = v36
	v113 = v39
	v116 = v42
	v117 = v45
	v118 = v47
	goto L7
L7:
	;
	v121 = int32(1) << (uint(v35) % 32)
	if v121&int32(14786) != 0 {
		v169 = v43
		goto L32
	} else {
		goto L33
	}
L8:
	;
	v62 = int32(10000)
	if base.Ui32(v62) <= base.Ui32(v55) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v65 = v62
	goto L11
L10:
	;
	v65 = v55
	goto L11
L11:
	;
	v70 = F_palloc(m, v65*int32(10)+int32(7))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v70 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v77 = int32(1)
	v80 = (v39-v42)>>(uint(v77)%32) + v77
	v82 = v80 << (uint(v77) % 32)
	if v82 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v91 = v84 + (v65<<(uint(int32(1))%32)+int32(7))&int32(2147483640)
	v93 = v80 << (uint(int32(3)) % 32)
	if v93 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v83 = F__emscripten_memcpy_bulkmem(m, v70, v42, v82)
	mBase = m.M
	v84 = v83
	goto L18
L17:
	;
	v84 = v70
	goto L18
L18:
	;
	goto L15
L19:
	;
	if v24+int32(1680) != v42 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v94 = F__emscripten_memcpy_bulkmem(m, v91, v47, v93)
	mBase = m.M
	v95 = v94
	goto L22
L21:
	;
	v95 = v91
	goto L22
L22:
	;
	goto L19
L23:
	;
	F_pfree(m, v42)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(v65) <= base.Ui32(v80) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	v946 = int32(1)
	v950 = v84
	goto L1
L28:
	;
	goto L29
L29:
	;
	v111 = v93 + v95 - int32(8)
	v113 = v84 + v80<<(uint(int32(1))%32) - int32(2)
	v116 = v84
	v117 = v65
	v118 = v95
	goto L7
L30:
	;
	goto L4
L31:
	;
	v35 = v927
	v36 = v914
	v39 = v916 + int32(2)
	v42 = v116
	v43 = v919
	v45 = v117
	v47 = v118
	goto L3
L32:
	;
	if v121&int32(1053) != 0 {
		goto L30
	} else {
		goto L49
	}
L33:
	;
	v127 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[1338]))))
	if v43 == int32(-2) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v147 = v127 + v146
	if base.Ui32(int32(12)) < base.Ui32(v147) {
		v169 = v145
		goto L32
	} else {
		goto L43
	}
L35:
	;
	v132 = F_seg_yylex(m, v24+int32(2088), l2)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L38
	}
L36:
	;
	v134 = v43
	goto L37
L37:
	;
	if v134 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v134 = v132
	goto L37
L39:
	;
	v137 = int32(0)
	v145 = v137
	v146 = v137
	goto L34
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(int32(261)) < base.Ui32(v134) {
		v145 = v134
		v146 = int32(2)
		goto L34
	} else {
		goto L42
	}
L42:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+uint32(_consts[1339]))))
	v145 = v134
	v146 = v144
	goto L34
L43:
	;
	v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v147)+uint32(_consts[1340]))))
	if v146 != v152 {
		v169 = v145
		goto L32
	} else {
		goto L44
	}
L44:
	;
	switch v147 - int32(1) {
	case 0:
		goto L30
	default:
		goto L45
	case 7:
		v946 = int32(0)
		v950 = v116
		goto L1
	}
L45:
	;
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v24)+2088))
	*(*int64)(unsafe.Add(mBase, uint32(v111)+8)) = v157
	if v145 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v161 = int32(-2)
	goto L48
L47:
	;
	v161 = int32(0)
	goto L48
L48:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+uint32(_consts[1341]))))
	v914 = v111 + int32(8)
	v916 = v113
	v919 = v161
	v927 = v166
	goto L31
L49:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[1342]))))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+uint32(_consts[1343]))))
	v185 = v111 + (int32(1)-v181)<<(uint(int32(3))%32)
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+6)))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+5)))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+4)))
	v189 = *(*float32)(unsafe.Add(mBase, uint32(v185)))
	switch v178 - int32(2) {
	case 0:
		goto L59
	case 1:
		goto L58
	case 2:
		goto L57
	case 3:
		goto L56
	case 4:
		goto L55
	case 5:
		goto L54
	case 6:
		goto L53
	case 7:
		goto L52
	default:
		v869 = v187
		v870 = v188
		v874 = v189
		goto L50
	}
L50:
	;
	v878 = v111 - v181<<(uint(int32(3))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v878)+14)) = uint16(v186)
	*(*uint8)(unsafe.Add(mBase, uint32(v878)+13)) = uint8(v869)
	*(*uint8)(unsafe.Add(mBase, uint32(v878)+12)) = uint8(v870)
	*(*float32)(unsafe.Add(mBase, uint32(v878)+8)) = v874
	v884 = v878 + int32(8)
	v887 = v113 - v181<<(uint(int32(1))%32)
	v888 = int32(*(*int16)(unsafe.Add(mBase, uint32(v887))))
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+uint32(_consts[1344]))))
	v897 = v895 - int32(7)
	v899 = int32(*(*int8)(unsafe.Add(mBase, uint32(v897)+uint32(_consts[1345]))))
	v900 = v888 + v899
	if base.Ui32(int32(12)) < base.Ui32(v900) {
		goto L286
	} else {
		goto L287
	}
L51:
	;
	v869 = v866
	v870 = int32(0)
	v874 = v865
	goto L50
L52:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v782 = F_float4in_internal(m, v780, int32(354856), v780, l1)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L12
	} else {
		goto L254
	}
L53:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v695 = F_float4in_internal(m, v693, int32(354856), v693, l1)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L12
	} else {
		goto L222
	}
L54:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v612 = F_float4in_internal(m, v610, int32(354856), v610, l1)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L12
	} else {
		goto L190
	}
L55:
	;
	v601 = *(*float32)(unsafe.Add(mBase, uint32(v111)))
	*(*float32)(unsafe.Add(mBase, uint32(l0))) = v601
	*(*float32)(unsafe.Add(mBase, uint32(l0)+4)) = v601
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v604)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v604)
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v607)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v607)
	v869 = v187
	v870 = v188
	v874 = v189
	goto L50
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(-8388608)
	v591 = *(*float32)(unsafe.Add(mBase, uint32(v111)))
	v592 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v592)
	*(*float32)(unsafe.Add(mBase, uint32(l0)+4)) = v591
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+5)))
	v596 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v596)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v595)
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v599)
	v869 = v187
	v870 = v188
	v874 = v189
	goto L50
L57:
	;
	v573 = *(*float32)(unsafe.Add(mBase, uint32(v111-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2139095040)
	*(*float32)(unsafe.Add(mBase, uint32(l0))) = v573
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111-int32(3)))))
	v580 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v580)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v579)
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111-int32(4)))))
	v586 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v586)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v585)
	v869 = v187
	v870 = v188
	v874 = v189
	goto L50
L58:
	;
	v526 = *(*float32)(unsafe.Add(mBase, uint32(v111-int32(16))))
	*(*float32)(unsafe.Add(mBase, uint32(l0))) = v526
	v528 = *(*float32)(unsafe.Add(mBase, uint32(v111)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+4)) = v528
	if base.F32_gt(v526, v528) != 0 {
		goto L182
	} else {
		goto L183
	}
L59:
	;
	v192 = int32(16)
	v193 = v111 - v192
	v194 = *(*float32)(unsafe.Add(mBase, uint32(v193)))
	v195 = *(*float32)(unsafe.Add(mBase, uint32(v111)))
	v196 = base.F32_sub(v194, v195)
	*(*float32)(unsafe.Add(mBase, uint32(l0))) = v196
	v198 = *(*float32)(unsafe.Add(mBase, uint32(v193)))
	v199 = *(*float32)(unsafe.Add(mBase, uint32(v111)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+4)) = base.F32_add(v198, v199)
	*(*float64)(unsafe.Add(mBase, uint32(v24)+16)) = base.F64_promote_f32(v196)
	v210 = F_pg_snprintf(m, v24+int32(48), int32(25), int32(355281), v24+v192)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	v219 = v24 + int32(48)
	goto L62
L61:
	;
	if int32(6) <= v278 {
		goto L85
	} else {
		goto L86
	}
L62:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	v225 = v223 - int32(43)
	if base.Ui32(int32(5)) < base.Ui32(v225) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v236 = v219
	v237 = v223
	v239 = int32(1)
	goto L70
L64:
	;
	goto L63
L65:
	;
	if int32(1)<<(uint(v225)%32)&int32(37) == int32(0) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v219 = v219 + int32(1)
	goto L62
L67:
	;
	goto L61
L68:
	;
	v278 = v239
	goto L67
L69:
	;
	v251 = v236
	v252 = v237
	v253 = int32(0)
	goto L76
L70:
	;
	switch v237 - int32(46) {
	case 0:
		v245 = v239
		goto L73
	case 1:
		goto L69
	case 2:
		goto L74
	default:
		goto L72
	}
L71:
	;
	if v237 == int32(0) {
		goto L68
	} else {
		goto L75
	}
L72:
	;
	goto L71
L73:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+1)))
	v236 = v236 + int32(1)
	v237 = v246
	v239 = v245
	goto L70
L74:
	;
	v245 = v239 + int32(1)
	goto L73
L75:
	;
	goto L69
L76:
	;
	if v252 != int32(46) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if v271 != 0 {
		v278 = v271
		goto L67
	} else {
		goto L84
	}
L78:
	;
	goto L77
L79:
	;
	if base.Ui32(int32(9)) < base.Ui32((v252-int32(48))&int32(255)) {
		v271 = v253
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v265 = v253 + base.B2i32(v252 != int32(46))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+1)))
	if v266 != 0 {
		v251 = v251 + int32(1)
		v252 = v266
		v253 = v265
		goto L76
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v271 = v265
	goto L78
L84:
	;
	goto L68
L85:
	;
	v282 = int32(6)
	goto L87
L86:
	;
	v282 = v278
	goto L87
L87:
	;
	v284 = v111 - int32(11)
	v285 = int32(*(*int8)(unsafe.Add(mBase, uint32(v284))))
	v286 = int32(*(*int8)(unsafe.Add(mBase, uint32(v111)+5)))
	if v286 < v285 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v288 = v285
	goto L90
L89:
	;
	v288 = v286
	goto L90
L90:
	;
	if v288 < v282 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v297 = v24 + int32(48)
	goto L95
L92:
	;
	v362 = v288
	goto L93
L93:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v362)
	v364 = *(*float32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*float64)(unsafe.Add(mBase, uint32(v24))) = base.F64_promote_f32(v364)
	v371 = F_pg_snprintf(m, v24+int32(48), int32(25), int32(355281), v24)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L12
	} else {
		goto L121
	}
L94:
	;
	if int32(6) <= v356 {
		goto L118
	} else {
		goto L119
	}
L95:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	v303 = v301 - int32(43)
	if base.Ui32(int32(5)) < base.Ui32(v303) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v314 = v297
	v315 = v301
	v317 = int32(1)
	goto L103
L97:
	;
	goto L96
L98:
	;
	if int32(1)<<(uint(v303)%32)&int32(37) == int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v297 = v297 + int32(1)
	goto L95
L100:
	;
	goto L94
L101:
	;
	v356 = v317
	goto L100
L102:
	;
	v329 = v314
	v330 = v315
	v331 = int32(0)
	goto L109
L103:
	;
	switch v315 - int32(46) {
	case 0:
		v323 = v317
		goto L106
	case 1:
		goto L102
	case 2:
		goto L107
	default:
		goto L105
	}
L104:
	;
	if v315 == int32(0) {
		goto L101
	} else {
		goto L108
	}
L105:
	;
	goto L104
L106:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	v314 = v314 + int32(1)
	v315 = v324
	v317 = v323
	goto L103
L107:
	;
	v323 = v317 + int32(1)
	goto L106
L108:
	;
	goto L102
L109:
	;
	if v330 != int32(46) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	if v349 != 0 {
		v356 = v349
		goto L100
	} else {
		goto L117
	}
L111:
	;
	goto L110
L112:
	;
	if base.Ui32(int32(9)) < base.Ui32((v330-int32(48))&int32(255)) {
		v349 = v331
		goto L111
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v343 = v331 + base.B2i32(v330 != int32(46))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+1)))
	if v344 != 0 {
		v329 = v329 + int32(1)
		v330 = v344
		v331 = v343
		goto L109
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	v349 = v343
	goto L111
L117:
	;
	goto L101
L118:
	;
	v360 = int32(6)
	goto L120
L119:
	;
	v360 = v356
	goto L120
L120:
	;
	v362 = v360
	goto L93
L121:
	;
	v380 = v24 + int32(48)
	goto L123
L122:
	;
	if int32(6) <= v439 {
		goto L146
	} else {
		goto L147
	}
L123:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	v386 = v384 - int32(43)
	if base.Ui32(int32(5)) < base.Ui32(v386) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v397 = v380
	v398 = v384
	v400 = int32(1)
	goto L131
L125:
	;
	goto L124
L126:
	;
	if int32(1)<<(uint(v386)%32)&int32(37) == int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v380 = v380 + int32(1)
	goto L123
L128:
	;
	goto L122
L129:
	;
	v439 = v400
	goto L128
L130:
	;
	v412 = v397
	v413 = v398
	v414 = int32(0)
	goto L137
L131:
	;
	switch v398 - int32(46) {
	case 0:
		v406 = v400
		goto L134
	case 1:
		goto L130
	case 2:
		goto L135
	default:
		goto L133
	}
L132:
	;
	if v398 == int32(0) {
		goto L129
	} else {
		goto L136
	}
L133:
	;
	goto L132
L134:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+1)))
	v397 = v397 + int32(1)
	v398 = v407
	v400 = v406
	goto L131
L135:
	;
	v406 = v400 + int32(1)
	goto L134
L136:
	;
	goto L130
L137:
	;
	if v413 != int32(46) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	if v432 != 0 {
		v439 = v432
		goto L128
	} else {
		goto L145
	}
L139:
	;
	goto L138
L140:
	;
	if base.Ui32(int32(9)) < base.Ui32((v413-int32(48))&int32(255)) {
		v432 = v414
		goto L139
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v426 = v414 + base.B2i32(v413 != int32(46))
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+1)))
	if v427 != 0 {
		v412 = v412 + int32(1)
		v413 = v427
		v414 = v426
		goto L137
	} else {
		goto L144
	}
L143:
	;
	goto L142
L144:
	;
	v432 = v426
	goto L139
L145:
	;
	goto L129
L146:
	;
	v443 = int32(6)
	goto L148
L147:
	;
	v443 = v439
	goto L148
L148:
	;
	v444 = int32(*(*int8)(unsafe.Add(mBase, uint32(v284))))
	v445 = int32(*(*int8)(unsafe.Add(mBase, uint32(v111)+5)))
	if v445 < v444 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v447 = v444
	goto L151
L150:
	;
	v447 = v445
	goto L151
L151:
	;
	if v447 < v443 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v456 = v24 + int32(48)
	goto L156
L153:
	;
	v520 = v447
	goto L154
L154:
	;
	v521 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v521)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v520)
	v869 = v187
	v870 = v188
	v874 = v189
	goto L50
L155:
	;
	if int32(6) <= v515 {
		goto L179
	} else {
		goto L180
	}
L156:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	v462 = v460 - int32(43)
	if base.Ui32(int32(5)) < base.Ui32(v462) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v473 = v456
	v474 = v460
	v476 = int32(1)
	goto L164
L158:
	;
	goto L157
L159:
	;
	if int32(1)<<(uint(v462)%32)&int32(37) == int32(0) {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v456 = v456 + int32(1)
	goto L156
L161:
	;
	goto L155
L162:
	;
	v515 = v476
	goto L161
L163:
	;
	v488 = v473
	v489 = v474
	v490 = int32(0)
	goto L170
L164:
	;
	switch v474 - int32(46) {
	case 0:
		v482 = v476
		goto L167
	case 1:
		goto L163
	case 2:
		goto L168
	default:
		goto L166
	}
L165:
	;
	if v474 == int32(0) {
		goto L162
	} else {
		goto L169
	}
L166:
	;
	goto L165
L167:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473)+1)))
	v473 = v473 + int32(1)
	v474 = v483
	v476 = v482
	goto L164
L168:
	;
	v482 = v476 + int32(1)
	goto L167
L169:
	;
	goto L163
L170:
	;
	if v489 != int32(46) {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	if v508 != 0 {
		v515 = v508
		goto L161
	} else {
		goto L178
	}
L172:
	;
	goto L171
L173:
	;
	if base.Ui32(int32(9)) < base.Ui32((v489-int32(48))&int32(255)) {
		v508 = v490
		goto L172
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v502 = v490 + base.B2i32(v489 != int32(46))
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
	if v503 != 0 {
		v488 = v488 + int32(1)
		v489 = v503
		v490 = v502
		goto L170
	} else {
		goto L177
	}
L176:
	;
	goto L175
L177:
	;
	v508 = v502
	goto L172
L178:
	;
	goto L162
L179:
	;
	v519 = int32(6)
	goto L181
L180:
	;
	v519 = v515
	goto L181
L181:
	;
	v520 = v519
	goto L154
L182:
	;
	v531 = int32(1)
	v532 = F_errsave_start(m, l1)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L12
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111-int32(11)))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v561)
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v563)
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111-int32(12)))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v567)
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v569)
	v869 = v187
	v870 = v188
	v874 = v189
	goto L50
L185:
	;
	if v532 == int32(0) {
		v946 = v531
		v950 = v116
		goto L1
	} else {
		goto L186
	}
L186:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L12
	} else {
		goto L187
	}
L187:
	;
	v539 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v540 = *(*float32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*float64)(unsafe.Add(mBase, uint32(v24)+40)) = base.F64_promote_f32(v540)
	*(*float64)(unsafe.Add(mBase, uint32(v24)+32)) = base.F64_promote_f32(v539)
	F_errmsg(m, int32(355187), v24+int32(32))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L12
	} else {
		goto L188
	}
L188:
	;
	F_errsave_finish(m, l1, int32(27583), int32(86), int32(378915))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L12
	} else {
		goto L189
	}
L189:
	;
	v946 = v531
	v950 = v116
	goto L1
L190:
	;
	if l1 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v629 = v624
	goto L196
L192:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v616 != int32(447) {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v619 == int32(0) {
		goto L191
	} else {
		goto L194
	}
L194:
	;
	v946 = int32(1)
	v950 = v116
	goto L1
L195:
	;
	if int32(6) <= v688 {
		goto L219
	} else {
		goto L220
	}
L196:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629))))
	v635 = v633 - int32(43)
	if base.Ui32(int32(5)) < base.Ui32(v635) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v646 = v629
	v647 = v633
	v649 = int32(1)
	goto L204
L198:
	;
	goto L197
L199:
	;
	if int32(1)<<(uint(v635)%32)&int32(37) == int32(0) {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v629 = v629 + int32(1)
	goto L196
L201:
	;
	goto L195
L202:
	;
	v688 = v649
	goto L201
L203:
	;
	v661 = v646
	v662 = v647
	v663 = int32(0)
	goto L210
L204:
	;
	switch v647 - int32(46) {
	case 0:
		v655 = v649
		goto L207
	case 1:
		goto L203
	case 2:
		goto L208
	default:
		goto L206
	}
L205:
	;
	if v647 == int32(0) {
		goto L202
	} else {
		goto L209
	}
L206:
	;
	goto L205
L207:
	;
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646)+1)))
	v646 = v646 + int32(1)
	v647 = v656
	v649 = v655
	goto L204
L208:
	;
	v655 = v649 + int32(1)
	goto L207
L209:
	;
	goto L203
L210:
	;
	if v662 != int32(46) {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	if v681 != 0 {
		v688 = v681
		goto L201
	} else {
		goto L218
	}
L212:
	;
	goto L211
L213:
	;
	if base.Ui32(int32(9)) < base.Ui32((v662-int32(48))&int32(255)) {
		v681 = v663
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v675 = v663 + base.B2i32(v662 != int32(46))
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661)+1)))
	if v676 != 0 {
		v661 = v661 + int32(1)
		v662 = v676
		v663 = v675
		goto L210
	} else {
		goto L217
	}
L216:
	;
	goto L215
L217:
	;
	v681 = v675
	goto L212
L218:
	;
	goto L202
L219:
	;
	v692 = int32(6)
	goto L221
L220:
	;
	v692 = v688
	goto L221
L221:
	;
	v865 = v612
	v866 = v692
	goto L51
L222:
	;
	if l1 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v111-int32(8))))
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708))))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v716 = v711
	goto L228
L224:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v699 != int32(447) {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v702 == int32(0) {
		goto L223
	} else {
		goto L226
	}
L226:
	;
	v946 = int32(1)
	v950 = v116
	goto L1
L227:
	;
	if int32(6) <= v775 {
		goto L251
	} else {
		goto L252
	}
L228:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v716))))
	v722 = v720 - int32(43)
	if base.Ui32(int32(5)) < base.Ui32(v722) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v733 = v716
	v734 = v720
	v736 = int32(1)
	goto L236
L230:
	;
	goto L229
L231:
	;
	if int32(1)<<(uint(v722)%32)&int32(37) == int32(0) {
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v716 = v716 + int32(1)
	goto L228
L233:
	;
	goto L227
L234:
	;
	v775 = v736
	goto L233
L235:
	;
	v748 = v733
	v749 = v734
	v750 = int32(0)
	goto L242
L236:
	;
	switch v734 - int32(46) {
	case 0:
		v742 = v736
		goto L239
	case 1:
		goto L235
	case 2:
		goto L240
	default:
		goto L238
	}
L237:
	;
	if v734 == int32(0) {
		goto L234
	} else {
		goto L241
	}
L238:
	;
	goto L237
L239:
	;
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733)+1)))
	v733 = v733 + int32(1)
	v734 = v743
	v736 = v742
	goto L236
L240:
	;
	v742 = v736 + int32(1)
	goto L239
L241:
	;
	goto L235
L242:
	;
	if v749 != int32(46) {
		goto L245
	} else {
		goto L246
	}
L243:
	;
	if v768 != 0 {
		v775 = v768
		goto L233
	} else {
		goto L250
	}
L244:
	;
	goto L243
L245:
	;
	if base.Ui32(int32(9)) < base.Ui32((v749-int32(48))&int32(255)) {
		v768 = v750
		goto L244
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v762 = v750 + base.B2i32(v749 != int32(46))
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v748)+1)))
	if v763 != 0 {
		v748 = v748 + int32(1)
		v749 = v763
		v750 = v762
		goto L242
	} else {
		goto L249
	}
L248:
	;
	goto L247
L249:
	;
	v768 = v762
	goto L244
L250:
	;
	goto L234
L251:
	;
	v779 = int32(6)
	goto L253
L252:
	;
	v779 = v775
	goto L253
L253:
	;
	v869 = v779
	v870 = v709
	v874 = v695
	goto L50
L254:
	;
	if l1 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v799 = v794
	goto L260
L256:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v786 != int32(447) {
		goto L255
	} else {
		goto L257
	}
L257:
	;
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v789 == int32(0) {
		goto L255
	} else {
		goto L258
	}
L258:
	;
	v946 = int32(1)
	v950 = v116
	goto L1
L259:
	;
	if int32(6) <= v858 {
		goto L283
	} else {
		goto L284
	}
L260:
	;
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v799))))
	v805 = v803 - int32(43)
	if base.Ui32(int32(5)) < base.Ui32(v805) {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v816 = v799
	v817 = v803
	v819 = int32(1)
	goto L268
L262:
	;
	goto L261
L263:
	;
	if int32(1)<<(uint(v805)%32)&int32(37) == int32(0) {
		goto L262
	} else {
		goto L264
	}
L264:
	;
	v799 = v799 + int32(1)
	goto L260
L265:
	;
	goto L259
L266:
	;
	v858 = v819
	goto L265
L267:
	;
	v831 = v816
	v832 = v817
	v833 = int32(0)
	goto L274
L268:
	;
	switch v817 - int32(46) {
	case 0:
		v825 = v819
		goto L271
	case 1:
		goto L267
	case 2:
		goto L272
	default:
		goto L270
	}
L269:
	;
	if v817 == int32(0) {
		goto L266
	} else {
		goto L273
	}
L270:
	;
	goto L269
L271:
	;
	v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816)+1)))
	v816 = v816 + int32(1)
	v817 = v826
	v819 = v825
	goto L268
L272:
	;
	v825 = v819 + int32(1)
	goto L271
L273:
	;
	goto L267
L274:
	;
	if v832 != int32(46) {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	if v851 != 0 {
		v858 = v851
		goto L265
	} else {
		goto L282
	}
L276:
	;
	goto L275
L277:
	;
	if base.Ui32(int32(9)) < base.Ui32((v832-int32(48))&int32(255)) {
		v851 = v833
		goto L276
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v845 = v833 + base.B2i32(v832 != int32(46))
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v831)+1)))
	if v846 != 0 {
		v831 = v831 + int32(1)
		v832 = v846
		v833 = v845
		goto L274
	} else {
		goto L281
	}
L280:
	;
	goto L279
L281:
	;
	v851 = v845
	goto L276
L282:
	;
	goto L266
L283:
	;
	v862 = int32(6)
	goto L285
L284:
	;
	v862 = v858
	goto L285
L285:
	;
	v865 = v782
	v866 = v862
	goto L51
L286:
	;
	v912 = int32(*(*int8)(unsafe.Add(mBase, uint32(v897)+uint32(_consts[1346]))))
	v914 = v884
	v916 = v887
	v919 = v169
	v927 = v912
	goto L31
L287:
	;
	v905 = int32(*(*int8)(unsafe.Add(mBase, uint32(v900)+uint32(_consts[1340]))))
	if v888 != v905 {
		goto L286
	} else {
		goto L288
	}
L288:
	;
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v900)+uint32(_consts[1341]))))
	v914 = v884
	v916 = v887
	v919 = v169
	v927 = v909
	goto L31
L289:
	;
	v946 = int32(1)
	v950 = v116
	goto L1
L290:
	;
	v946 = int32(2)
	v950 = v42
	goto L1
L291:
	;
	F_pfree(m, v950)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L12
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	m.G0 = v24 + int32(2096)
	return v946
L294:
	;
	goto L293
}
