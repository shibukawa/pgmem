package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GUCArrayAdd(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
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
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = F_validate_option_array_item(m, l1, l2, int32(0))
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
	v20 = F_find_option(m, l1, int32(0), int32(1), int32(19))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v23 = v22
	goto L6
L5:
	;
	v23 = l1
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v23
	v27 = F_psprintf(m, int32(180447), v10)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v29 = F_cstring_to_text(m, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v29
	if l0 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	m.G0 = v10 + int32(32)
	return v139
L10:
	;
	v33 = l0 + int32(16)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v34 + v35
	if v34 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v130 = F_construct_array_builtin(m, v10+int32(28), int32(1), int32(25))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L38
	}
L13:
	;
	v124 = F_array_set(m, l0, v10+int32(24), v29, int32(-1), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L37
	}
L14:
	;
	goto L15
L15:
	;
	v53 = F_array_ref(m, l0, v10+int32(16), v10+int32(23))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L13
L17:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+23)))
	if v55 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v109 = v107 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v109 <= v111 {
		goto L15
	} else {
		goto L36
	}
L19:
	;
	v56 = F_text_to_cstring(m, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v58 = F_strlen(m, v23)
	mBase = m.M
	v60 = v58 + int32(1)
	if v60 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v104 != 0 {
		goto L18
	} else {
		goto L35
	}
L22:
	;
	v104 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v66 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v67 = v56
	v68 = v27
	v69 = v60
	v70 = v66
	goto L29
L26:
	;
	v92 = v27
	v96 = int32(0)
	goto L27
L27:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v104 = v96 - v97
	goto L21
L28:
	;
	v92 = v87
	v96 = v89
	goto L27
L29:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v70 != v72 {
		v87 = v68
		v89 = v70
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v87 = v81
	v89 = int32(0)
	goto L28
L31:
	;
	if v72 == int32(0) {
		v87 = v68
		v89 = v70
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v77 = v69 - int32(1)
	if v77 == int32(0) {
		v87 = v68
		v89 = v70
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v80 = int32(1)
	v81 = v68 + v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v82 != 0 {
		v67 = v67 + v80
		v68 = v81
		v69 = v77
		v70 = v82
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v105
	goto L13
L36:
	;
	goto L16
L37:
	;
	v139 = v124
	goto L9
L38:
	;
	v139 = v130
	goto L9
}
func F_GUC_yylex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
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
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v515 int32
	_ = v515
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
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
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
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
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1122 int32
	_ = v1122
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1174 int32
	_ = v1174
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1302 int32
	_ = v1302
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1344 int32
	_ = v1344
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1354 int32
	_ = v1354
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v82 = l0
	v85 = v80
	goto L22
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v80 = v14
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v17 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(1)
	goto L7
L6:
	;
	goto L7
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v26
	goto L10
L9:
	;
	goto L10
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v28 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[717]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32
	goto L13
L12:
	;
	goto L13
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v34 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v64
	v68 = v61 + v62<<(uint(int32(2))%32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v74
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v76)
	v80 = v70
	goto L1
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34+v35<<(uint(int32(2))%32))))
	if v39 != 0 {
		v61 = v34
		v62 = v35
		v63 = v39
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_GUC_yyensure_buffer_stack(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
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
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v47 = F_GUC_yy_create_buffer(m, v46, l0)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v51 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v49+v50<<(uint(v51)%32)))) = v47
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55+v56<<(uint(v51)%32))))
	v61 = v55
	v62 = v56
	v63 = v60
	goto L14
L22:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v94)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	v99 = v96
	v100 = v85
	v102 = v85
	goto L24
L24:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+uint32(_consts[1135]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v99&int32(2147483647)))%64)&int64(1101676544007) == int64(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+68)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v82)+64)) = v99
	goto L28
L27:
	;
	goto L28
L28:
	;
	v124 = int32(1)
	v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v99<<(uint(v124)%32))+uint32(_consts[1136]))))
	v129 = v128 + v112
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129<<(uint(v124)%32))+uint32(_consts[1137]))))
	if v134 != v99 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v138 = v99
	v140 = v112
	v142 = v112
	goto L32
L30:
	;
	v183 = v129
	goto L31
L31:
	;
	v194 = int32(1)
	v200 = int32(*(*int16)(unsafe.Add(mBase, uint32(v183<<(uint(v194)%32))+uint32(_consts[1138]))))
	if v200 != int32(40) {
		v99 = v200
		v102 = v102 + v194
		goto L24
	} else {
		goto L38
	}
L32:
	;
	v152 = int32(*(*int16)(unsafe.Add(mBase, uint32(v138<<(uint(int32(1))%32))+uint32(_consts[1139]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v138&int32(2147483647)))%64)&int64(241224598912) != int64(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v183 = v175
	goto L31
L34:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+uint32(_consts[1140]))))
	v165 = v164
	goto L36
L35:
	;
	v165 = v142
	goto L36
L36:
	;
	v169 = v165 & int32(255)
	v170 = int32(1)
	v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v152<<(uint(v170)%32))+uint32(_consts[1136]))))
	v175 = v169 + v174
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175<<(uint(v170)%32))+uint32(_consts[1137]))))
	if v152&int32(65535) != v180 {
		v138 = v152
		v140 = v169
		v142 = v165
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v212 = v100
	goto L39
L39:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v82)+64))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v82)+68))
	v219 = v215
	v224 = v216
	v226 = v212
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+80)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v82)+32)) = v224 - v226
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+24)) = uint8(v232)
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v234)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v224
	v241 = int32(*(*int16)(unsafe.Add(mBase, uint32(v219<<(uint(int32(1))%32))+uint32(_consts[1141]))))
	v243 = v241
	v245 = v224
	goto L43
L43:
	;
	if v243 != int32(13) {
		goto L61
	} else {
		goto L62
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v1363
	*(*int32)(unsafe.Add(mBase, uint32(v82)+48)) = int32(0)
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	v1379 = base.I32_div_s(v1375-int32(1), int32(2))
	v243 = v1379 + int32(14)
	v245 = v1363
	goto L43
L46:
	;
	F_GUC_flex_fatal(m, int32(32266))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L19
	} else {
		goto L259
	}
L47:
	;
	F_GUC_flex_fatal(m, int32(697791))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L19
	} else {
		goto L258
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v1239 = v1228 + v1235
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v1239
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	if base.Ui32(v1239) <= base.Ui32(v1230) {
		v219 = v1241
		v224 = v1239
		v226 = v1230
		goto L41
	} else {
		goto L239
	}
L50:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v908)))
	*(*int32)(unsafe.Add(mBase, uint32(v909)+16)) = v897
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	if v912 != 0 {
		v1034 = int32(0)
		goto L189
	} else {
		goto L190
	}
L51:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v897 = v880
	v908 = v891 + v892<<(uint(int32(2))%32)
	goto L50
L52:
	;
	F_GUC_flex_fatal(m, int32(463974))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L19
	} else {
		goto L188
	}
L53:
	;
	F_GUC_flex_fatal(m, int32(459697))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L19
	} else {
		goto L187
	}
L54:
	;
	return v871
L55:
	;
	v871 = int32(3)
	goto L54
L56:
	;
	return int32(6)
L57:
	;
	return int32(2)
L58:
	;
	return int32(7)
L59:
	;
	return int32(1)
L60:
	;
	F_GUC_flex_fatal(m, int32(433516))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L19
	} else {
		goto L186
	}
L61:
	;
	switch v243 {
	case 0:
		goto L69
	case 1:
		goto L68
	case 2, 3:
		v85 = v245
		goto L22
	case 4:
		goto L59
	case 5:
		goto L58
	case 6:
		goto L57
	case 7:
		goto L56
	case 8:
		goto L55
	case 9:
		v871 = int32(4)
		goto L54
	case 10:
		goto L67
	case 11:
		goto L66
	case 12:
		goto L65
	default:
		goto L60
	case 14:
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v82)+80))
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v277)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v283 = v279 + v280<<(uint(int32(2))%32)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+44))
	if v285 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L64:
	;
	return int32(0)
L65:
	;
	F_GUC_flex_fatal(m, int32(463326))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L19
	} else {
		goto L70
	}
L66:
	;
	return int32(100)
L67:
	;
	return int32(5)
L68:
	;
	v259 = int32(4534868)
	v261 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
	*(*int32)(unsafe.Add(mBase, _consts[1142])) = v261 + int32(1)
	return int32(99)
L69:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v257)
	v212 = v226
	goto L39
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+28)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v290))) = v291
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v295 = int32(2)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v293+v294<<(uint(v295)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+44)) = int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v301+v302<<(uint(v295)%32))))
	v307 = v306
	v308 = v301
	v309 = v302
	goto L73
L72:
	;
	v307 = v284
	v308 = v279
	v309 = v280
	goto L73
L73:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v82)+36))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v307)+4))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	v313 = v311 + v312
	if base.Ui32(v310) <= base.Ui32(v313) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v82)+80))
	v319 = v315 + (v276 ^ int32(-1)) + v224
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v319
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	if base.Ui32(v315) < base.Ui32(v319) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(v313+int32(1)) < base.Ui32(v310) {
		goto L53
	} else {
		goto L109
	}
L77:
	;
	v325 = v321
	v328 = v315
	goto L80
L78:
	;
	v434 = v321
	goto L79
L79:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v434&int32(2147483647)))%64)&int64(1101676544007) == int64(0) {
		goto L98
	} else {
		goto L99
	}
L80:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
	if v335 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v434 = v428
	goto L79
L82:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+uint32(_consts[1135]))))
	v340 = v338
	goto L84
L83:
	;
	v340 = int32(1)
	goto L84
L84:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v325&int32(2147483647)))%64)&int64(1101676544007) == int64(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+68)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v82)+64)) = v325
	goto L87
L86:
	;
	goto L87
L87:
	;
	v353 = v340 & int32(255)
	v354 = int32(1)
	v358 = int32(*(*int16)(unsafe.Add(mBase, uint32(v325<<(uint(v354)%32))+uint32(_consts[1136]))))
	v359 = v353 + v358
	v364 = int32(*(*int16)(unsafe.Add(mBase, uint32(v359<<(uint(v354)%32))+uint32(_consts[1137]))))
	if v364 != v325 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v368 = v325
	v370 = v340
	v372 = v353
	goto L91
L89:
	;
	v413 = v359
	goto L90
L90:
	;
	v424 = int32(1)
	v428 = int32(*(*int16)(unsafe.Add(mBase, uint32(v413<<(uint(v424)%32))+uint32(_consts[1138]))))
	v430 = v328 + v424
	if v430 != v319 {
		v325 = v428
		v328 = v430
		goto L80
	} else {
		goto L97
	}
L91:
	;
	v382 = int32(*(*int16)(unsafe.Add(mBase, uint32(v368<<(uint(int32(1))%32))+uint32(_consts[1139]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v368&int32(2147483647)))%64)&int64(241224598912) != int64(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v413 = v405
	goto L90
L93:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+uint32(_consts[1140]))))
	v395 = v394
	goto L95
L94:
	;
	v395 = v370
	goto L95
L95:
	;
	v399 = v395 & int32(255)
	v400 = int32(1)
	v404 = int32(*(*int16)(unsafe.Add(mBase, uint32(v382<<(uint(v400)%32))+uint32(_consts[1136]))))
	v405 = v399 + v404
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v405<<(uint(v400)%32))+uint32(_consts[1137]))))
	if v382&int32(65535) != v410 {
		v368 = v382
		v370 = v395
		v372 = v399
		goto L91
	} else {
		goto L96
	}
L96:
	;
	goto L92
L97:
	;
	goto L81
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+68)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v82)+64)) = v434
	goto L100
L99:
	;
	goto L100
L100:
	;
	v455 = int32(1)
	v459 = int32(*(*int16)(unsafe.Add(mBase, uint32(v434<<(uint(v455)%32))+uint32(_consts[1136]))))
	v461 = v459 + v455
	v466 = int32(*(*int16)(unsafe.Add(mBase, uint32(v461<<(uint(v455)%32))+uint32(_consts[1137]))))
	if v466 != v434 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v470 = v434
	goto L104
L102:
	;
	v500 = v461
	goto L103
L103:
	;
	v515 = int32(*(*int16)(unsafe.Add(mBase, uint32(v500<<(uint(int32(1))%32))+uint32(_consts[1138]))))
	if v515 == int32(40) {
		v212 = v315
		goto L39
	} else {
		goto L107
	}
L104:
	;
	v480 = int32(1)
	v484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v470<<(uint(v480)%32))+uint32(_consts[1139]))))
	v485 = base.I32_extend16_s(v484)
	v490 = int32(*(*int16)(unsafe.Add(mBase, uint32(v485<<(uint(v480)%32))+uint32(_consts[1136]))))
	v492 = v490 + v480
	v497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v492<<(uint(v480)%32))+uint32(_consts[1137]))))
	if v484 != v497 {
		v470 = v485
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v500 = v492
	goto L103
L106:
	;
	goto L105
L107:
	;
	if v500&int32(2147483647) == int32(0) {
		v212 = v315
		goto L39
	} else {
		goto L108
	}
L108:
	;
	v523 = v319 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v523
	v99 = v515
	v100 = v315
	v102 = v523
	goto L24
L109:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v82)+80))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v307)+40))
	if v529 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	if v310-v528 != int32(1) {
		v1228 = v311
		v1230 = v528
		v1235 = v312
		goto L49
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v537 = v528 ^ int32(-1) + v310
	if v537 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v1363 = v528
	goto L45
L114:
	;
	v538 = int32(7)
	v539 = v537 & v538
	if base.Ui32(v310-v528-int32(2)) < base.Ui32(v538) {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v637 = v307
	v639 = v308
	v641 = v309
	goto L116
L116:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v637)+44))
	if v647 == int32(2) {
		goto L130
	} else {
		goto L131
	}
L117:
	;
	if v539 != 0 {
		goto L124
	} else {
		goto L125
	}
L118:
	;
	v584 = v311
	v585 = v528
	goto L117
L119:
	;
	goto L120
L120:
	;
	v549 = v311
	v550 = v528
	v552 = int32(0)
	goto L121
L121:
	;
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	*(*uint8)(unsafe.Add(mBase, uint32(v549))) = uint8(v560)
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v549)+1)) = uint8(v562)
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v549)+2)) = uint8(v564)
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v549)+3)) = uint8(v566)
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v549)+4)) = uint8(v568)
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v549)+5)) = uint8(v570)
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v549)+6)) = uint8(v572)
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v549)+7)) = uint8(v574)
	v576 = int32(8)
	v577 = v549 + v576
	v579 = v550 + v576
	v581 = v552 + v576
	if v581 != v537&int32(-8) {
		v549 = v577
		v550 = v579
		v552 = v581
		goto L121
	} else {
		goto L123
	}
L122:
	;
	v584 = v577
	v585 = v579
	goto L117
L123:
	;
	goto L122
L124:
	;
	v597 = v584
	v598 = v585
	v600 = int32(0)
	goto L127
L125:
	;
	goto L126
L126:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v629+v630<<(uint(int32(2))%32))))
	v637 = v634
	v639 = v629
	v641 = v630
	goto L116
L127:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598))))
	*(*uint8)(unsafe.Add(mBase, uint32(v597))) = uint8(v608)
	v610 = int32(1)
	v615 = v600 + v610
	if v615 != v539 {
		v597 = v597 + v610
		v598 = v598 + v610
		v600 = v615
		goto L127
	} else {
		goto L129
	}
L128:
	;
	goto L126
L129:
	;
	goto L128
L130:
	;
	v650 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+28)) = v650
	v897 = v650
	v908 = v639 + v641<<(uint(int32(2))%32)
	goto L50
L131:
	;
	goto L132
L132:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v637)+12))
	v657 = v528 - v310
	v658 = v656 + v657
	if v658 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v82)+36))
	v663 = v656
	v664 = v637
	v668 = v661
	goto L136
L134:
	;
	v708 = v637
	v710 = v658
	goto L135
L135:
	;
	v718 = int32(8192)
	if base.Ui32(v718) <= base.Ui32(v710) {
		goto L146
	} else {
		goto L147
	}
L136:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v664)+20))
	if v674 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v708 = v701
	v710 = v703
	goto L135
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v664)+4)) = int32(0)
	goto L46
L139:
	;
	goto L140
L140:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v664)+4))
	if base.Ui32(int32(2147483646)) <= base.Ui32(v663) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v685 = int32(-3)
	goto L143
L142:
	;
	v685 = v663 << (uint(int32(1)) % 32)
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v664)+12)) = v685
	v689 = F_emscripten_builtin_realloc(m, v679, v685+int32(2))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v664)+4)) = v689
	if v689 == int32(0) {
		goto L46
	} else {
		goto L144
	}
L144:
	;
	v694 = v689 + (v668 - v679)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v694
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v696+v697<<(uint(int32(2))%32))))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)+12))
	v703 = v702 + v657
	if v703 == int32(0) {
		v663 = v702
		v664 = v701
		v668 = v694
		goto L136
	} else {
		goto L145
	}
L145:
	;
	goto L137
L146:
	;
	v721 = v718
	goto L148
L147:
	;
	v721 = v710
	goto L148
L148:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v708)+24))
	if v723 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v725 = int32(0)
	goto L153
L150:
	;
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0)
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v789+v790<<(uint(int32(2))%32))))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v794)+4))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v799 = F_fread(m, v795+v537, int32(1), v721, v798)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L19
	} else {
		goto L168
	}
L152:
	;
	switch v740 {
	case 0:
		goto L160
	default:
		v784 = v754
		goto L158
	case 11:
		goto L159
	}
L153:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v737 = F_do_getc(m, v736)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L19
	} else {
		goto L156
	}
L154:
	;
	v754 = v721
	goto L152
L155:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v741+v742<<(uint(int32(2))%32))))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v746)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v747+v537+v725))) = uint8(v737)
	v752 = v725 + int32(1)
	if v752 != v721 {
		v725 = v752
		goto L153
	} else {
		goto L157
	}
L156:
	;
	v740 = v737 + int32(1)
	switch v740 {
	case 0, 11:
		v754 = v725
		goto L152
	default:
		goto L155
	}
L157:
	;
	goto L154
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+28)) = v784
	v880 = v784
	goto L51
L159:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v771+v772<<(uint(int32(2))%32))))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v776)+4))
	v780 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v777+v537+v754))) = uint8(v780)
	v784 = v754 + int32(1)
	goto L158
L160:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v755)+76))
	if v756 < int32(0) {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	if int32(base.Ui32(v761)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v784 = v754
		goto L158
	} else {
		goto L166
	}
L162:
	;
	goto L161
L163:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	v761 = v759
	goto L162
L164:
	;
	goto L165
L165:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	v761 = v760
	goto L162
L166:
	;
	F_GUC_flex_fatal(m, int32(463974))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L19
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	v802 = v799
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+28)) = v802
	if v802 != 0 {
		v880 = v802
		goto L51
	} else {
		goto L171
	}
L171:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v814)+76))
	if v815 < int32(0) {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	if int32(base.Ui32(v820)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L177
	} else {
		goto L178
	}
L173:
	;
	goto L172
L174:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v814)))
	v820 = v818
	goto L173
L175:
	;
	goto L176
L176:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v814)))
	v820 = v819
	goto L173
L177:
	;
	v880 = int32(0)
	goto L51
L178:
	;
	goto L179
L179:
	;
	v829 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v829 != int32(27) {
		goto L52
	} else {
		goto L180
	}
L180:
	;
	v833 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v833
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v835)+76))
	if v833 <= v836 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v847+v848<<(uint(int32(2))%32))))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v852)+4))
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v857 = F_fread(m, v853+v537, int32(1), v721, v856)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L19
	} else {
		goto L185
	}
L182:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v835)))
	*(*int32)(unsafe.Add(mBase, uint32(v835))) = v839 & int32(-49)
	goto L181
L183:
	;
	goto L184
L184:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v835)))
	*(*int32)(unsafe.Add(mBase, uint32(v835))) = v843 & int32(-49)
	goto L181
L185:
	;
	v802 = v857
	goto L169
L186:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	v1036 = v1035 + v537
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1037+v1038<<(uint(int32(2))%32))))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1042)+12))
	if base.Ui32(v1043) < base.Ui32(v1036) {
		goto L213
	} else {
		goto L214
	}
L190:
	;
	if v537 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v916 != 0 {
		goto L196
	} else {
		goto L197
	}
L192:
	;
	goto L193
L193:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v1022 = int32(2)
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1020+v1021<<(uint(v1022)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1025)+44)) = v1022
	v1034 = v1022
	goto L189
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v983)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v983))) = v915
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v989 != 0 {
		goto L209
	} else {
		goto L210
	}
L195:
	;
	v939 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v937+v940<<(uint(int32(2))%32))))
	if v944 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L196:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v916+v917<<(uint(int32(2))%32))))
	if v921 != 0 {
		v937 = v916
		goto L195
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	F_GUC_yyensure_buffer_stack(m, v82)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L19
	} else {
		goto L200
	}
L199:
	;
	goto L198
L200:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v925 = F_GUC_yy_create_buffer(m, v924, v82)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L19
	} else {
		goto L201
	}
L201:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v927+v928<<(uint(int32(2))%32)))) = v925
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v933 != 0 {
		v937 = v933
		goto L195
	} else {
		goto L202
	}
L202:
	;
	v935 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v983 = int32(0)
	v985 = v935
	goto L194
L203:
	;
	v983 = int32(0)
	v985 = v939
	goto L194
L204:
	;
	goto L205
L205:
	;
	v948 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v944)+16)) = v948
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v944)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v950))) = uint8(v948)
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v944)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v953)+1)) = uint8(v948)
	*(*int32)(unsafe.Add(mBase, uint32(v944)+44)) = v948
	*(*int32)(unsafe.Add(mBase, uint32(v944)+28)) = int32(1)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v944)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v944)+8)) = v960
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v962 == v948 {
		v983 = v944
		v985 = v939
		goto L194
	} else {
		goto L206
	}
L206:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v968 = v962 + v965<<(uint(int32(2))%32)
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v968)))
	if v944 != v969 {
		v983 = v944
		v985 = v939
		goto L194
	} else {
		goto L207
	}
L207:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v969)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+28)) = v971
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v968)))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v973)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+80)) = v974
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v974
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v968)))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v977)))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v978
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v974))))
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+24)) = uint8(v980)
	v983 = v944
	v985 = v939
	goto L194
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v983)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v985
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v1006 = v1002 + v1003<<(uint(int32(2))%32)
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v1006)))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1007)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+28)) = v1008
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1006)))
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v1010)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v82)+80)) = v1011
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1006)))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v1015
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1011))))
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+24)) = uint8(v1017)
	v1034 = int32(1)
	goto L189
L209:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v989+v990<<(uint(int32(2))%32))))
	if v983 == v994 {
		goto L208
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v983)+32)) = int64(1)
	goto L208
L212:
	;
	goto L211
L213:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v1042)+4))
	v1048 = v1036 + int32(base.Ui32(v1035)>>(uint(int32(1))%32))
	v1049 = F_emscripten_builtin_realloc(m, v1045, v1048)
	mBase = m.M
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v1052 = int32(2)
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1050+v1051<<(uint(v1052)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1055)+4)) = v1049
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1057+v1058<<(uint(v1052)%32))))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1062)+4))
	if v1063 == int32(0) {
		goto L47
	} else {
		goto L216
	}
L214:
	;
	v1073 = v1036
	v1074 = v1037
	v1075 = v1038
	goto L215
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+28)) = v1073
	v1077 = int32(2)
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1074+v1075<<(uint(v1077)%32))))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1080)+4))
	v1083 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1081+v1073))) = uint8(v1083)
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1085+v1086<<(uint(v1077)%32))))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+4))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1091+v1092)+1)) = uint8(v1083)
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v1100 = v1096 + v1097<<(uint(v1077)%32)
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1100)))
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+80)) = v1102
	if v1034 == int32(1) {
		v1363 = v1102
		goto L45
	} else {
		goto L217
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1062)+12)) = v1048 - int32(2)
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v1073 = v1069 + v537
	v1074 = v1072
	v1075 = v1071
	goto L215
L217:
	;
	switch v1034 - int32(1) {
	case 0:
		goto L48
	case 1:
		goto L218
	default:
		goto L219
	}
L218:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1100)))
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+4))
	v1228 = v1226
	v1230 = v1102
	v1235 = v1224
	goto L49
L219:
	;
	v1111 = v1102 + (v276 ^ int32(-1)) + v224
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v1111
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	if base.Ui32(v1111) <= base.Ui32(v1102) {
		v99 = v1113
		v100 = v1102
		v102 = v1111
		goto L24
	} else {
		goto L220
	}
L220:
	;
	v1117 = v1113
	v1122 = v1102
	goto L221
L221:
	;
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122))))
	if v1127 != 0 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	v99 = v1220
	v100 = v1102
	v102 = v1111
	goto L24
L223:
	;
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1127)+uint32(_consts[1135]))))
	v1132 = v1130
	goto L225
L224:
	;
	v1132 = int32(1)
	goto L225
L225:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v1117&int32(2147483647)))%64)&int64(1101676544007) == int64(0) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+68)) = v1122
	*(*int32)(unsafe.Add(mBase, uint32(v82)+64)) = v1117
	goto L228
L227:
	;
	goto L228
L228:
	;
	v1145 = v1132 & int32(255)
	v1146 = int32(1)
	v1150 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1117<<(uint(v1146)%32))+uint32(_consts[1136]))))
	v1151 = v1145 + v1150
	v1156 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1151<<(uint(v1146)%32))+uint32(_consts[1137]))))
	if v1156 != v1117 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1160 = v1117
	v1162 = v1132
	v1164 = v1145
	goto L232
L230:
	;
	v1205 = v1151
	goto L231
L231:
	;
	v1216 = int32(1)
	v1220 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1205<<(uint(v1216)%32))+uint32(_consts[1138]))))
	v1222 = v1122 + v1216
	if v1111 != v1222 {
		v1117 = v1220
		v1122 = v1222
		goto L221
	} else {
		goto L238
	}
L232:
	;
	v1174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1160<<(uint(int32(1))%32))+uint32(_consts[1139]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v1160&int32(2147483647)))%64)&int64(241224598912) != int64(0) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	v1205 = v1197
	goto L231
L234:
	;
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1164)+uint32(_consts[1140]))))
	v1187 = v1186
	goto L236
L235:
	;
	v1187 = v1162
	goto L236
L236:
	;
	v1191 = v1187 & int32(255)
	v1192 = int32(1)
	v1196 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1174<<(uint(v1192)%32))+uint32(_consts[1136]))))
	v1197 = v1191 + v1196
	v1202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1197<<(uint(v1192)%32))+uint32(_consts[1137]))))
	if v1174&int32(65535) != v1202 {
		v1160 = v1174
		v1162 = v1187
		v1164 = v1191
		goto L232
	} else {
		goto L237
	}
L237:
	;
	goto L233
L238:
	;
	goto L222
L239:
	;
	v1245 = v1241
	v1248 = v1230
	goto L240
L240:
	;
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248))))
	if v1255 != 0 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v219 = v1348
	v224 = v1239
	v226 = v1230
	goto L41
L242:
	;
	v1258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1255)+uint32(_consts[1135]))))
	v1260 = v1258
	goto L244
L243:
	;
	v1260 = int32(1)
	goto L244
L244:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v1245&int32(2147483647)))%64)&int64(1101676544007) == int64(0) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+68)) = v1248
	*(*int32)(unsafe.Add(mBase, uint32(v82)+64)) = v1245
	goto L247
L246:
	;
	goto L247
L247:
	;
	v1273 = v1260 & int32(255)
	v1274 = int32(1)
	v1278 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1245<<(uint(v1274)%32))+uint32(_consts[1136]))))
	v1279 = v1273 + v1278
	v1284 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1279<<(uint(v1274)%32))+uint32(_consts[1137]))))
	if v1284 != v1245 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1288 = v1245
	v1290 = v1260
	v1292 = v1273
	goto L251
L249:
	;
	v1333 = v1279
	goto L250
L250:
	;
	v1344 = int32(1)
	v1348 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1333<<(uint(v1344)%32))+uint32(_consts[1138]))))
	v1350 = v1248 + v1344
	if v1350 != v1239 {
		v1245 = v1348
		v1248 = v1350
		goto L240
	} else {
		goto L257
	}
L251:
	;
	v1302 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1288<<(uint(int32(1))%32))+uint32(_consts[1139]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v1288&int32(2147483647)))%64)&int64(241224598912) != int64(0) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v1333 = v1325
	goto L250
L253:
	;
	v1314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292)+uint32(_consts[1140]))))
	v1315 = v1314
	goto L255
L254:
	;
	v1315 = v1290
	goto L255
L255:
	;
	v1319 = v1315 & int32(255)
	v1320 = int32(1)
	v1324 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1302<<(uint(v1320)%32))+uint32(_consts[1136]))))
	v1325 = v1319 + v1324
	v1330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1325<<(uint(v1320)%32))+uint32(_consts[1137]))))
	if v1302&int32(65535) != v1330 {
		v1288 = v1302
		v1290 = v1315
		v1292 = v1319
		goto L251
	} else {
		goto L256
	}
L256:
	;
	goto L252
L257:
	;
	goto L241
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ProcessGUCArray(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	F_TransformGUCArray(m, l0, v13+int32(12), v13+int32(8))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v29 = int32(0)
	goto L3
L3:
	;
	v39 = int32(0)
	if v27 == v39 {
		v49 = v39
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_list_free(m, v27)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L19
	}
L5:
	;
	if v26 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v43 <= v29 {
		v49 = int32(0)
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v49 = v45 + v29<<(uint(int32(2))%32)
	goto L5
L8:
	;
	goto L4
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v52 <= v29 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if v49 == int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v59 = v56 + v29<<(uint(int32(2))%32)
	if v59 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if base.B2i32(l2 == int32(9))|base.B2i32(base.Ui32(int32(10)) < base.Ui32(l2)) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v68 = v66
	goto L15
L14:
	;
	v68 = int32(10)
	goto L15
L15:
	;
	v70 = int32(0)
	v72 = F_set_config_with_handle(m, v62, int32(0), v64, l1, l2, v68, l3, int32(1), v70, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_pfree(m, v62)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_pfree(m, v64)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v29 = v29 + int32(1)
	goto L3
L19:
	;
	F_list_free(m, v26)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	m.G0 = v13 + int32(16)
	return
}
func F_SplitGUCList(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v8 = l0
	goto L1
L1:
	;
	v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8))))
	goto L3
L2:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v25 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if base.B2i32(v15 == int32(32))|base.B2i32(base.Ui32((v15-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v8 = v8 + int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	return int32(1)
L6:
	;
	v30 = v25
	v32 = v8
	goto L7
L7:
	;
	v34 = v30 & int32(255)
	if v34 != int32(34) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v300 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v246))) = uint8(v300)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v303 = F_lappend(m, v302, v248)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L96
	} else {
		goto L97
	}
L10:
	;
	return int32(0)
L11:
	;
	v249 = v244
	goto L84
L12:
	;
	if v34 == int32(0) {
		v63 = v32
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v70 = v32 + int32(1)
	v71 = int32(34)
	v72 = F___strchrnul(m, v70, v71)
	mBase = m.M
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v74 == v71 {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	if v63 == v32 {
		goto L10
	} else {
		goto L24
	}
L16:
	;
	if v34 == int32(44) {
		v63 = v32
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v41 = v32
	v43 = v30
	goto L18
L18:
	;
	v46 = base.I32_extend8_s(v43)
	goto L20
L19:
	;
	v63 = v57
	goto L15
L20:
	;
	if base.B2i32(v46 == int32(32))|base.B2i32(base.Ui32((v46-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v63 = v41
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v57 = v41 + int32(1)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v58 == int32(0) {
		v63 = v57
		goto L15
	} else {
		goto L22
	}
L22:
	;
	if v58 != int32(44) {
		v41 = v57
		v43 = v58
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v244 = v63
	v246 = v63
	v248 = v32
	goto L11
L25:
	;
	if v78 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L26:
	;
	v78 = v72
	goto L28
L27:
	;
	v78 = int32(0)
	goto L28
L28:
	;
	goto L25
L29:
	;
	v83 = v78
	goto L30
L30:
	;
	v87 = v83 + int32(1)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v88 != int32(34) {
		v244 = v87
		v246 = v83
		v248 = v70
		goto L11
	} else {
		goto L32
	}
L31:
	;
	goto L10
L32:
	;
	v91 = F_strlen(m, v83)
	mBase = m.M
	if v83 == v87 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v236 = int32(34)
	v237 = F___strchrnul(m, v87, v236)
	mBase = m.M
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	if v239 == v236 {
		goto L80
	} else {
		goto L81
	}
L34:
	;
	goto L33
L35:
	;
	v95 = v83 + v91
	if base.Ui32(v87-v95) <= base.Ui32(int32(0)-v91<<(uint(int32(1))%32)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v102 = F___memcpy(m, v83, v87, v91)
	mBase = m.M
	goto L33
L37:
	;
	goto L38
L38:
	;
	v105 = (v83 ^ v87) & int32(3)
	if base.Ui32(v83) < base.Ui32(v87) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if v207 == int32(0) {
		goto L34
	} else {
		goto L75
	}
L40:
	;
	if base.Ui32(v185) <= base.Ui32(int32(3)) {
		v206 = v184
		v207 = v185
		v208 = v186
		goto L39
	} else {
		goto L71
	}
L41:
	;
	if v105 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	if v105 != 0 {
		v167 = v91
		goto L54
	} else {
		goto L55
	}
L44:
	;
	v206 = v87
	v207 = v91
	v208 = v83
	goto L39
L45:
	;
	goto L46
L46:
	;
	if v83&int32(3) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v184 = v87
	v185 = v91
	v186 = v83
	goto L40
L48:
	;
	goto L49
L49:
	;
	v112 = v87
	v113 = v91
	v114 = v83
	goto L50
L50:
	;
	if v113 == int32(0) {
		goto L34
	} else {
		goto L52
	}
L51:
	;
	v184 = v121
	v185 = v123
	v186 = v125
	goto L40
L52:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v118)
	v120 = int32(1)
	v121 = v112 + v120
	v123 = v113 - v120
	v125 = v114 + v120
	if v125&int32(3) != 0 {
		v112 = v121
		v113 = v123
		v114 = v125
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	if v167 == int32(0) {
		goto L34
	} else {
		goto L67
	}
L55:
	;
	if v95&int32(3) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v132 = v91
	goto L59
L57:
	;
	v147 = v91
	goto L58
L58:
	;
	if base.Ui32(v147) <= base.Ui32(int32(3)) {
		v167 = v147
		goto L54
	} else {
		goto L63
	}
L59:
	;
	if v132 == int32(0) {
		goto L34
	} else {
		goto L61
	}
L60:
	;
	v147 = v138
	goto L58
L61:
	;
	v138 = v132 - int32(1)
	v139 = v83 + v138
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v138))))
	*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v141)
	if v139&int32(3) != 0 {
		v132 = v138
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v154 = v147
	goto L64
L64:
	;
	v158 = v154 - int32(4)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v87+v158)))
	*(*int32)(unsafe.Add(mBase, uint32(v83+v158))) = v161
	if base.Ui32(int32(3)) < base.Ui32(v158) {
		v154 = v158
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v167 = v158
	goto L54
L66:
	;
	goto L65
L67:
	;
	v174 = v167
	goto L68
L68:
	;
	v178 = v174 - int32(1)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v178))))
	*(*uint8)(unsafe.Add(mBase, uint32(v83+v178))) = uint8(v181)
	if v178 != 0 {
		v174 = v178
		goto L68
	} else {
		goto L70
	}
L69:
	;
	goto L34
L70:
	;
	goto L69
L71:
	;
	v191 = v184
	v192 = v185
	v193 = v186
	goto L72
L72:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v195
	v197 = int32(4)
	v198 = v191 + v197
	v200 = v193 + v197
	v202 = v192 - v197
	if base.Ui32(int32(3)) < base.Ui32(v202) {
		v191 = v198
		v192 = v202
		v193 = v200
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v206 = v198
	v207 = v202
	v208 = v200
	goto L39
L74:
	;
	goto L73
L75:
	;
	v213 = v206
	v214 = v207
	v215 = v208
	goto L76
L76:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v217)
	v219 = int32(1)
	v224 = v214 - v219
	if v224 != 0 {
		v213 = v213 + v219
		v214 = v224
		v215 = v215 + v219
		goto L76
	} else {
		goto L78
	}
L77:
	;
	goto L34
L78:
	;
	goto L77
L79:
	;
	if v243 != 0 {
		v83 = v243
		goto L30
	} else {
		goto L83
	}
L80:
	;
	v243 = v237
	goto L82
L81:
	;
	v243 = int32(0)
	goto L82
L82:
	;
	goto L79
L83:
	;
	goto L31
L84:
	;
	v256 = int32(*(*int8)(unsafe.Add(mBase, uint32(v249))))
	goto L86
L85:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v266 == int32(44) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	if base.B2i32(v256 == int32(32))|base.B2i32(base.Ui32((v256-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v249 = v249 + int32(1)
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v272 = v249
	goto L91
L89:
	;
	goto L90
L90:
	;
	if v266 == int32(0) {
		v298 = v249
		goto L9
	} else {
		goto L95
	}
L91:
	;
	v275 = v272 + int32(1)
	v276 = int32(*(*int8)(unsafe.Add(mBase, uint32(v275))))
	goto L93
L93:
	;
	if base.B2i32(v276 == int32(32))|base.B2i32(base.Ui32((v276-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v272 = v275
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v298 = v275
	goto L9
L95:
	;
	goto L10
L96:
	;
	return int32(0)
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v303
	if v266 != int32(44) {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	v30 = v310
	v32 = v298
	goto L7
}
