package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GUCArrayAdd(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_validate_option_array_item(m, l1, l2, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = F_find_option(m, l1, int32(0), int32(1), int32(19))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = v21
	goto L6
L5:
	;
	v22 = l1
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v22
	v26 = F_psprintf(m, int32(_a_F_GUCArrayAdd_0), v9)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v28 = F_cstring_to_text(m, v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v28
	if l0 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	m.G0 = v9 + int32(32)
	return v134
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v32 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v31 + v32
	if v31 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v126 = F_construct_array_builtin(m, v9+int32(28), int32(1), int32(25))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L37
	}
L13:
	;
	v120 = F_array_set(m, l0, v9+int32(24), v28, int32(-1), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L36
	}
L14:
	;
	goto L15
L15:
	;
	v49 = F_array_ref(m, l0, v9+int32(16), v9+int32(23))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L13
L17:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+23)))
	if v51 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v106 = v104 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v106 <= v108 {
		goto L15
	} else {
		goto L35
	}
L19:
	;
	v52 = F_text_to_cstring(m, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v54 = F_strlen(m, v22)
	mBase = m.M
	v56 = v54 + int32(1)
	if v56 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v101 != 0 {
		goto L18
	} else {
		goto L34
	}
L22:
	;
	v101 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v62 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v63 = v52
	v64 = v26
	v65 = v56
	v66 = v62
	goto L29
L26:
	;
	v89 = v26
	v93 = int32(0)
	goto L27
L27:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v101 = v93 - v94
	goto L21
L28:
	;
	v89 = v84
	v93 = v86
	goto L27
L29:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if base.B2i32(v66 != v68)|base.B2i32(v68 == int32(0)) != 0 {
		v84 = v64
		v86 = v66
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v84 = v78
	v86 = int32(0)
	goto L28
L31:
	;
	v74 = v65 - int32(1)
	if v74 == int32(0) {
		v84 = v64
		v86 = v66
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v77 = int32(1)
	v78 = v64 + v77
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v79 != 0 {
		v63 = v63 + v77
		v64 = v78
		v65 = v74
		v66 = v79
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v102
	goto L13
L35:
	;
	goto L16
L36:
	;
	v134 = v120
	goto L9
L37:
	;
	v134 = v126
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
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
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
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
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v718 int32
	_ = v718
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v873 int32
	_ = v873
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
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
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
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
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1144 int32
	_ = v1144
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1264 int32
	_ = v1264
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1288 int32
	_ = v1288
	var v1293 int32
	_ = v1293
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1318 int32
	_ = v1318
	var v1321 int32
	_ = v1321
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v84 = l0
	v86 = v80
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
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_GUC_yylex[0]))
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
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_GUC_yylex[1]))
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
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v65
	v69 = v63 + v61<<(uint(int32(2))%32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v75
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v77)
	v80 = v71
	goto L1
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34+v35<<(uint(int32(2))%32))))
	if v39 != 0 {
		v61 = v35
		v63 = v34
		v64 = v39
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
	v61 = v56
	v63 = v55
	v64 = v60
	goto L14
L22:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v96)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v84)+44))
	v99 = v84
	v100 = v98
	v101 = v86
	v104 = v86
	goto L24
L24:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_GUC_yylex[2]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v100))%64)&int64(1101676544007) == int64(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+68)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v99)+64)) = v100
	goto L28
L27:
	;
	goto L28
L28:
	;
	v122 = int32(1)
	v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100<<(uint(v122)%32))+uint32(_c_F_GUC_yylex[3]))))
	v127 = v126 + v112
	v132 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127<<(uint(v122)%32))+uint32(_c_F_GUC_yylex[4]))))
	if v132 != v100 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v135 = v100
	v138 = v112
	v140 = v112
	goto L32
L30:
	;
	v179 = v127
	goto L31
L31:
	;
	v188 = int32(1)
	v194 = int32(*(*int16)(unsafe.Add(mBase, uint32(v179<<(uint(v188)%32))+uint32(_c_F_GUC_yylex[5]))))
	if v194 != int32(40) {
		v100 = v194
		v104 = v104 + v188
		goto L24
	} else {
		goto L38
	}
L32:
	;
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(v135<<(uint(int32(1))%32))+uint32(_c_F_GUC_yylex[6]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v135))%64)&int64(241224598912) != int64(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v179 = v169
	goto L31
L34:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+uint32(_c_F_GUC_yylex[7]))))
	v159 = v158
	goto L36
L35:
	;
	v159 = v138
	goto L36
L36:
	;
	v163 = v159 & int32(255)
	v164 = int32(1)
	v168 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150<<(uint(v164)%32))+uint32(_c_F_GUC_yylex[3]))))
	v169 = v163 + v168
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169<<(uint(v164)%32))+uint32(_c_F_GUC_yylex[4]))))
	if v150&int32(_a_F_GUC_yylex_0) != v174 {
		v135 = v150
		v138 = v159
		v140 = v163
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v204 = v101
	goto L39
L39:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v99)+64))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v99)+68))
	v212 = v209
	v218 = v204
	v219 = v210
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+80)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v99)+32)) = v219 - v218
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+24)) = uint8(v226)
	v228 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v228)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+36)) = v219
	v235 = int32(*(*int16)(unsafe.Add(mBase, uint32(v212<<(uint(int32(1))%32))+uint32(_c_F_GUC_yylex[8]))))
	v238 = v219
	v239 = v235
	goto L43
L43:
	;
	if v239 != int32(13) {
		goto L61
	} else {
		goto L62
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+36)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v99)+48)) = int32(0)
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v99)+44))
	v1338 = base.I32_div_s(v1334-int32(1), int32(2))
	v238 = v1321
	v239 = v1338 + int32(14)
	goto L43
L46:
	;
	F_GUC_flex_fatal(m, int32(_a_F_GUC_yylex_1))
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		goto L19
	} else {
		goto L247
	}
L47:
	;
	F_GUC_flex_fatal(m, int32(_a_F_GUC_yylex_2))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L19
	} else {
		goto L246
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v1205 = v1196 + v1198
	*(*int32)(unsafe.Add(mBase, uint32(v99)+36)) = v1205
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v99)+44))
	if base.Ui32(v1205) <= base.Ui32(v1195) {
		v212 = v1207
		v218 = v1195
		v219 = v1205
		goto L41
	} else {
		goto L227
	}
L50:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v880)))
	*(*int32)(unsafe.Add(mBase, uint32(v881)+16)) = v873
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v99)+28))
	if v884 != 0 {
		v1006 = int32(0)
		goto L177
	} else {
		goto L178
	}
L51:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v873 = v856
	v880 = v863 + v864<<(uint(int32(2))%32)
	goto L50
L52:
	;
	F_GUC_flex_fatal(m, int32(_a_F_GUC_yylex_3))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L19
	} else {
		goto L176
	}
L53:
	;
	F_GUC_flex_fatal(m, int32(_a_F_GUC_yylex_4))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L19
	} else {
		goto L175
	}
L54:
	;
	return v843
L55:
	;
	v843 = int32(3)
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
	F_GUC_flex_fatal(m, int32(_a_F_GUC_yylex_5))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L19
	} else {
		goto L174
	}
L61:
	;
	switch v239 {
	case 0:
		goto L69
	case 1:
		goto L68
	case 2, 3:
		v84 = v99
		v86 = v238
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
		v843 = int32(4)
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
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v99)+80))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v271)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v277 = v273 + v274<<(uint(int32(2))%32)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+44))
	if v279 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L64:
	;
	return int32(0)
L65:
	;
	F_GUC_flex_fatal(m, int32(_a_F_GUC_yylex_6))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
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
	v253 = int32(_a_F_GUC_yylex_7)
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_GUC_yylex[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_GUC_yylex[9])) = v255 + int32(1)
	return int32(99)
L69:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v251)
	v204 = v218
	goto L39
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+28)) = v282
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v284))) = v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v289 = int32(2)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v287+v288<<(uint(v289)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v292)+44)) = int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v295+v296<<(uint(v289)%32))))
	v301 = v300
	v302 = v295
	v303 = v296
	goto L73
L72:
	;
	v301 = v278
	v302 = v273
	v303 = v274
	goto L73
L73:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v99)+36))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v99)+28))
	v307 = v305 + v306
	if base.Ui32(v304) <= base.Ui32(v307) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v99)+80))
	v312 = v270 ^ int32(-1) + v219
	v313 = v309 + v312
	*(*int32)(unsafe.Add(mBase, uint32(v99)+36)) = v313
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v99)+44))
	if int32(0) < v312 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(v307+int32(1)) < base.Ui32(v304) {
		goto L53
	} else {
		goto L109
	}
L77:
	;
	v319 = v315
	v323 = v309
	goto L80
L78:
	;
	v420 = v315
	goto L79
L79:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v420))%64)&int64(1101676544007) == int64(0) {
		goto L98
	} else {
		goto L99
	}
L80:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	if v331 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v420 = v415
	goto L79
L82:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+uint32(_c_F_GUC_yylex[2]))))
	v333 = v332
	goto L84
L83:
	;
	v333 = int32(1)
	goto L84
L84:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v319))%64)&int64(1101676544007) == int64(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+68)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v99)+64)) = v319
	goto L87
L86:
	;
	goto L87
L87:
	;
	v344 = v333 & int32(255)
	v345 = int32(1)
	v349 = int32(*(*int16)(unsafe.Add(mBase, uint32(v319<<(uint(v345)%32))+uint32(_c_F_GUC_yylex[3]))))
	v350 = v344 + v349
	v355 = int32(*(*int16)(unsafe.Add(mBase, uint32(v350<<(uint(v345)%32))+uint32(_c_F_GUC_yylex[4]))))
	if v355 != v319 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v358 = v319
	v361 = v333
	v363 = v344
	goto L91
L89:
	;
	v402 = v350
	goto L90
L90:
	;
	v411 = int32(1)
	v415 = int32(*(*int16)(unsafe.Add(mBase, uint32(v402<<(uint(v411)%32))+uint32(_c_F_GUC_yylex[5]))))
	v417 = v323 + v411
	if v417 != v313 {
		v319 = v415
		v323 = v417
		goto L80
	} else {
		goto L97
	}
L91:
	;
	v373 = int32(*(*int16)(unsafe.Add(mBase, uint32(v358<<(uint(int32(1))%32))+uint32(_c_F_GUC_yylex[6]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v358))%64)&int64(241224598912) != int64(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v402 = v392
	goto L90
L93:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_GUC_yylex[7]))))
	v382 = v381
	goto L95
L94:
	;
	v382 = v361
	goto L95
L95:
	;
	v386 = v382 & int32(255)
	v387 = int32(1)
	v391 = int32(*(*int16)(unsafe.Add(mBase, uint32(v373<<(uint(v387)%32))+uint32(_c_F_GUC_yylex[3]))))
	v392 = v386 + v391
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v392<<(uint(v387)%32))+uint32(_c_F_GUC_yylex[4]))))
	if v373&int32(_a_F_GUC_yylex_0) != v397 {
		v358 = v373
		v361 = v382
		v363 = v386
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
	*(*int32)(unsafe.Add(mBase, uint32(v99)+68)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v99)+64)) = v420
	goto L100
L99:
	;
	goto L100
L100:
	;
	v440 = int32(1)
	v444 = int32(*(*int16)(unsafe.Add(mBase, uint32(v420<<(uint(v440)%32))+uint32(_c_F_GUC_yylex[3]))))
	v446 = v444 + v440
	v451 = int32(*(*int16)(unsafe.Add(mBase, uint32(v446<<(uint(v440)%32))+uint32(_c_F_GUC_yylex[4]))))
	if v451 != v420 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v454 = v420
	goto L104
L102:
	;
	v487 = v446
	goto L103
L103:
	;
	if v487 == int32(0) {
		v204 = v309
		goto L39
	} else {
		goto L107
	}
L104:
	;
	v465 = int32(1)
	v469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v454<<(uint(v465)%32))+uint32(_c_F_GUC_yylex[6]))))
	v470 = base.I32_extend16_s(v469)
	v475 = int32(*(*int16)(unsafe.Add(mBase, uint32(v470<<(uint(v465)%32))+uint32(_c_F_GUC_yylex[3]))))
	v477 = v475 + v465
	v482 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v477<<(uint(v465)%32))+uint32(_c_F_GUC_yylex[4]))))
	if v469 != v482 {
		v454 = v470
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v487 = v477
	goto L103
L106:
	;
	goto L105
L107:
	;
	v502 = int32(*(*int16)(unsafe.Add(mBase, uint32(v487<<(uint(int32(1))%32))+uint32(_c_F_GUC_yylex[5]))))
	if v502 == int32(40) {
		v204 = v309
		goto L39
	} else {
		goto L108
	}
L108:
	;
	v506 = v313 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+36)) = v506
	v100 = v502
	v101 = v309
	v104 = v506
	goto L24
L109:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v99)+80))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v301)+40))
	if v512 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	if v304-v511 != int32(1) {
		v1195 = v511
		v1196 = v305
		v1198 = v306
		goto L49
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v520 = v511 ^ int32(-1) + v304
	if int32(0) < v520 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v1321 = v511
	goto L45
L114:
	;
	v523 = int32(7)
	v524 = v520 & v523
	if base.Ui32(v304-v511-int32(2)) < base.Ui32(v523) {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v623 = v301
	v626 = v302
	v628 = v303
	goto L116
L116:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v623)+44))
	if v634 == int32(2) {
		goto L129
	} else {
		goto L130
	}
L117:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v616+v617<<(uint(int32(2))%32))))
	v623 = v621
	v626 = v616
	v628 = v617
	goto L116
L118:
	;
	v584 = v571
	v586 = v573
	v587 = int32(0)
	goto L126
L119:
	;
	v571 = v511
	v573 = v305
	goto L118
L120:
	;
	goto L121
L121:
	;
	v534 = v511
	v536 = v305
	v537 = int32(0)
	goto L122
L122:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	*(*uint8)(unsafe.Add(mBase, uint32(v536))) = uint8(v545)
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+1)) = uint8(v547)
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+2)) = uint8(v549)
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+3)) = uint8(v551)
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+4)) = uint8(v553)
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+5)) = uint8(v555)
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+6)) = uint8(v557)
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v536)+7)) = uint8(v559)
	v561 = int32(8)
	v562 = v536 + v561
	v564 = v534 + v561
	v566 = v537 + v561
	if v566 != v520&int32(2147483640) {
		v534 = v564
		v536 = v562
		v537 = v566
		goto L122
	} else {
		goto L124
	}
L123:
	;
	if v524 == int32(0) {
		goto L117
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	v571 = v564
	v573 = v562
	goto L118
L126:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584))))
	*(*uint8)(unsafe.Add(mBase, uint32(v586))) = uint8(v595)
	v597 = int32(1)
	v602 = v587 + v597
	if v602 != v524 {
		v584 = v584 + v597
		v586 = v586 + v597
		v587 = v602
		goto L126
	} else {
		goto L128
	}
L127:
	;
	goto L117
L128:
	;
	goto L127
L129:
	;
	v637 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+28)) = v637
	v873 = v637
	v880 = v626 + v628<<(uint(int32(2))%32)
	goto L50
L130:
	;
	goto L131
L131:
	;
	v643 = int32(0)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v623)+12))
	v645 = v511 - v304
	v646 = v644 + v645
	if v646 <= v643 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v99)+36))
	v651 = v623
	v652 = v649
	v656 = v644
	goto L135
L133:
	;
	v697 = v623
	v699 = v646
	goto L134
L134:
	;
	v708 = int32(_a_F_GUC_yylex_8)
	if base.Ui32(v708) <= base.Ui32(v699) {
		goto L145
	} else {
		goto L146
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
	v697 = v691
	v699 = v693
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
	v669 = v656 << (uint(int32(1)) % 32)
	if v669 <= int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v673 = base.I32_div_s(v656, int32(8))
	v675 = v673 + v656
	goto L142
L141:
	;
	v675 = v669
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v651)+12)) = v675
	v679 = F_emscripten_builtin_realloc(m, v667, v675+int32(2))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v651)+4)) = v679
	if v679 == int32(0) {
		goto L46
	} else {
		goto L143
	}
L143:
	;
	v684 = v679 + (v652 - v667)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+36)) = v684
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v686+v687<<(uint(int32(2))%32))))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v691)+12))
	v693 = v692 + v645
	if v693 <= int32(0) {
		v651 = v691
		v652 = v684
		v656 = v692
		goto L135
	} else {
		goto L144
	}
L144:
	;
	goto L136
L145:
	;
	v711 = v708
	goto L147
L146:
	;
	v711 = v699
	goto L147
L147:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v697)+24))
	if v712 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v718 = v643
	goto L152
L149:
	;
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GUC_yylex[10])) = int32(0)
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v773+v774<<(uint(int32(2))%32))))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v778)+4))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v783 = F_fread(m, v779+v520, int32(1), v711, v782)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L19
	} else {
		goto L163
	}
L151:
	;
	switch v729 {
	case 0:
		goto L159
	default:
		v768 = v743
		goto L157
	case 11:
		goto L158
	}
L152:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v726 = F_do_getc(m, v725)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L19
	} else {
		goto L155
	}
L153:
	;
	v743 = v711
	goto L151
L154:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v730+v731<<(uint(int32(2))%32))))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v735)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v736+v520+v718))) = uint8(v726)
	v741 = v718 + int32(1)
	if v741 != v711 {
		v718 = v741
		goto L152
	} else {
		goto L156
	}
L155:
	;
	v729 = v726 + int32(1)
	switch v729 {
	case 0, 11:
		v743 = v718
		goto L151
	default:
		goto L154
	}
L156:
	;
	goto L153
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+28)) = v768
	v856 = v768
	goto L51
L158:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v755+v756<<(uint(int32(2))%32))))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	v764 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v761+v520+v743))) = uint8(v764)
	v768 = v743 + int32(1)
	goto L157
L159:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v744)))
	goto L160
L160:
	;
	if int32(base.Ui32(v745)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v768 = v743
		goto L157
	} else {
		goto L161
	}
L161:
	;
	F_GUC_flex_fatal(m, int32(_a_F_GUC_yylex_3))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L19
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	v790 = v783
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+28)) = v790
	if v790 != 0 {
		v856 = v790
		goto L51
	} else {
		goto L166
	}
L166:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v798)))
	goto L167
L167:
	;
	if int32(base.Ui32(v799)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v856 = int32(0)
	goto L51
L169:
	;
	goto L170
L170:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _c_F_GUC_yylex[10]))
	if v808 != int32(27) {
		goto L52
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GUC_yylex[10])) = int32(0)
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v814)))
	*(*int32)(unsafe.Add(mBase, uint32(v814))) = v815 & int32(-49)
	goto L172
L172:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v819+v820<<(uint(int32(2))%32))))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v824)+4))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v829 = F_fread(m, v825+v520, int32(1), v711, v828)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L19
	} else {
		goto L173
	}
L173:
	;
	v790 = v829
	goto L164
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v99)+28))
	v1008 = v1007 + v520
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1009+v1010<<(uint(int32(2))%32))))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+12))
	if v1015 < v1008 {
		goto L201
	} else {
		goto L202
	}
L178:
	;
	if v520 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	if v888 != 0 {
		goto L184
	} else {
		goto L185
	}
L180:
	;
	goto L181
L181:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v994 = int32(2)
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v992+v993<<(uint(v994)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v997)+44)) = v994
	v1006 = v994
	goto L177
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v955)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v955))) = v887
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	if v961 != 0 {
		goto L197
	} else {
		goto L198
	}
L183:
	;
	v911 = *(*int32)(unsafe.Add(mBase, _c_F_GUC_yylex[10]))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v909+v912<<(uint(int32(2))%32))))
	if v916 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L184:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v888+v889<<(uint(int32(2))%32))))
	if v893 != 0 {
		v909 = v888
		goto L183
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	F_GUC_yyensure_buffer_stack(m, v99)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L19
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v897 = F_GUC_yy_create_buffer(m, v896, v99)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L19
	} else {
		goto L189
	}
L189:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v899+v900<<(uint(int32(2))%32)))) = v897
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	if v905 != 0 {
		v909 = v905
		goto L183
	} else {
		goto L190
	}
L190:
	;
	v907 = *(*int32)(unsafe.Add(mBase, _c_F_GUC_yylex[10]))
	v955 = int32(0)
	v957 = v907
	goto L182
L191:
	;
	v955 = int32(0)
	v957 = v911
	goto L182
L192:
	;
	goto L193
L193:
	;
	v920 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v916)+16)) = v920
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v916)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v922))) = uint8(v920)
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v916)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v925)+1)) = uint8(v920)
	*(*int32)(unsafe.Add(mBase, uint32(v916)+44)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v916)+28)) = int32(1)
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v916)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v916)+8)) = v932
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	if v934 == v920 {
		v955 = v916
		v957 = v911
		goto L182
	} else {
		goto L194
	}
L194:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v940 = v934 + v937<<(uint(int32(2))%32)
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v940)))
	if v916 != v941 {
		v955 = v916
		v957 = v911
		goto L182
	} else {
		goto L195
	}
L195:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v941)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+28)) = v943
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v940)))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v945)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+80)) = v946
	*(*int32)(unsafe.Add(mBase, uint32(v99)+36)) = v946
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v940)))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v949)))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v950
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v946))))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+24)) = uint8(v952)
	v955 = v916
	v957 = v911
	goto L182
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v955)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_GUC_yylex[10])) = v957
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v978 = v974 + v975<<(uint(int32(2))%32)
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v978)))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v979)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+28)) = v980
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v978)))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v982)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+36)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(v99)+80)) = v983
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v978)))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v987
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983))))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+24)) = uint8(v989)
	v1006 = int32(1)
	goto L177
L197:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v961+v962<<(uint(int32(2))%32))))
	if v955 == v966 {
		goto L196
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v955)+32)) = int64(1)
	goto L196
L200:
	;
	goto L199
L201:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+4))
	v1020 = v1008 + v1007>>(uint(int32(1))%32)
	v1021 = F_emscripten_builtin_realloc(m, v1017, v1020)
	mBase = m.M
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v1024 = int32(2)
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1022+v1023<<(uint(v1024)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1027)+4)) = v1021
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1029+v1030<<(uint(v1024)%32))))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+4))
	if v1035 == int32(0) {
		goto L47
	} else {
		goto L204
	}
L202:
	;
	v1045 = v1009
	v1047 = v1008
	v1048 = v1010
	goto L203
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+28)) = v1047
	v1050 = int32(2)
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1045+v1048<<(uint(v1050)%32))))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+4))
	v1056 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1054+v1047))) = uint8(v1056)
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1058+v1059<<(uint(v1050)%32))))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+4))
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v99)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1064+v1065)+1)) = uint8(v1056)
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v1073 = v1069 + v1070<<(uint(v1050)%32)
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1073)))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1074)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+80)) = v1075
	if v1006 == int32(1) {
		v1321 = v1075
		goto L45
	} else {
		goto L205
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1034)+12)) = v1020 - int32(2)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v99)+28))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v1045 = v1044
	v1047 = v1041 + v520
	v1048 = v1043
	goto L203
L205:
	;
	switch v1006 - int32(1) {
	case 0:
		goto L48
	case 1:
		goto L206
	default:
		goto L207
	}
L206:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v99)+28))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1073)))
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+4))
	v1195 = v1075
	v1196 = v1192
	v1198 = v1190
	goto L49
L207:
	;
	v1083 = v270 ^ int32(-1) + v219
	v1084 = v1075 + v1083
	*(*int32)(unsafe.Add(mBase, uint32(v99)+36)) = v1084
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v99)+44))
	if v1083 <= int32(0) {
		v100 = v1086
		v101 = v1075
		v104 = v1084
		goto L24
	} else {
		goto L208
	}
L208:
	;
	v1090 = v1086
	v1097 = v1075
	goto L209
L209:
	;
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1097))))
	if v1102 != 0 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v100 = v1186
	v101 = v1075
	v104 = v1084
	goto L24
L211:
	;
	v1103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1102)+uint32(_c_F_GUC_yylex[2]))))
	v1104 = v1103
	goto L213
L212:
	;
	v1104 = int32(1)
	goto L213
L213:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v1090))%64)&int64(1101676544007) == int64(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+68)) = v1097
	*(*int32)(unsafe.Add(mBase, uint32(v99)+64)) = v1090
	goto L216
L215:
	;
	goto L216
L216:
	;
	v1115 = v1104 & int32(255)
	v1116 = int32(1)
	v1120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1090<<(uint(v1116)%32))+uint32(_c_F_GUC_yylex[3]))))
	v1121 = v1115 + v1120
	v1126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1121<<(uint(v1116)%32))+uint32(_c_F_GUC_yylex[4]))))
	if v1126 != v1090 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1129 = v1090
	v1132 = v1104
	v1134 = v1115
	goto L220
L218:
	;
	v1173 = v1121
	goto L219
L219:
	;
	v1182 = int32(1)
	v1186 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1173<<(uint(v1182)%32))+uint32(_c_F_GUC_yylex[5]))))
	v1188 = v1097 + v1182
	if v1084 != v1188 {
		v1090 = v1186
		v1097 = v1188
		goto L209
	} else {
		goto L226
	}
L220:
	;
	v1144 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1129<<(uint(int32(1))%32))+uint32(_c_F_GUC_yylex[6]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v1129))%64)&int64(241224598912) != int64(0) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v1173 = v1163
	goto L219
L222:
	;
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1134)+uint32(_c_F_GUC_yylex[7]))))
	v1153 = v1152
	goto L224
L223:
	;
	v1153 = v1132
	goto L224
L224:
	;
	v1157 = v1153 & int32(255)
	v1158 = int32(1)
	v1162 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1144<<(uint(v1158)%32))+uint32(_c_F_GUC_yylex[3]))))
	v1163 = v1157 + v1162
	v1168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1163<<(uint(v1158)%32))+uint32(_c_F_GUC_yylex[4]))))
	if v1144&int32(_a_F_GUC_yylex_0) != v1168 {
		v1129 = v1144
		v1132 = v1153
		v1134 = v1157
		goto L220
	} else {
		goto L225
	}
L225:
	;
	goto L221
L226:
	;
	goto L210
L227:
	;
	v1210 = v1207
	v1214 = v1195
	goto L228
L228:
	;
	v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1214))))
	if v1221 != 0 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v212 = v1306
	v218 = v1195
	v219 = v1205
	goto L41
L230:
	;
	v1222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1221)+uint32(_c_F_GUC_yylex[2]))))
	v1224 = v1222
	goto L232
L231:
	;
	v1224 = int32(1)
	goto L232
L232:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v1210))%64)&int64(1101676544007) == int64(0) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+68)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v99)+64)) = v1210
	goto L235
L234:
	;
	goto L235
L235:
	;
	v1235 = v1224 & int32(255)
	v1236 = int32(1)
	v1240 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1210<<(uint(v1236)%32))+uint32(_c_F_GUC_yylex[3]))))
	v1241 = v1235 + v1240
	v1246 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1241<<(uint(v1236)%32))+uint32(_c_F_GUC_yylex[4]))))
	if v1246 != v1210 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1249 = v1210
	v1252 = v1224
	v1254 = v1235
	goto L239
L237:
	;
	v1293 = v1241
	goto L238
L238:
	;
	v1302 = int32(1)
	v1306 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1293<<(uint(v1302)%32))+uint32(_c_F_GUC_yylex[5]))))
	v1308 = v1214 + v1302
	if v1308 != v1205 {
		v1210 = v1306
		v1214 = v1308
		goto L228
	} else {
		goto L245
	}
L239:
	;
	v1264 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1249<<(uint(int32(1))%32))+uint32(_c_F_GUC_yylex[6]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v1249))%64)&int64(241224598912) != int64(0) {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v1293 = v1283
	goto L238
L241:
	;
	v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1254)+uint32(_c_F_GUC_yylex[7]))))
	v1273 = v1272
	goto L243
L242:
	;
	v1273 = v1252
	goto L243
L243:
	;
	v1277 = v1273 & int32(255)
	v1278 = int32(1)
	v1282 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1264<<(uint(v1278)%32))+uint32(_c_F_GUC_yylex[3]))))
	v1283 = v1277 + v1282
	v1288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1283<<(uint(v1278)%32))+uint32(_c_F_GUC_yylex[4]))))
	if v1264&int32(_a_F_GUC_yylex_0) != v1288 {
		v1249 = v1264
		v1252 = v1273
		v1254 = v1277
		goto L239
	} else {
		goto L244
	}
L244:
	;
	goto L240
L245:
	;
	goto L229
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L247:
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
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
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
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L18
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
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if base.B2i32(v49 == int32(0))|base.B2i32(v54 <= v29) != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v57 == int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57+v29<<(uint(int32(2))%32))))
	if base.B2i32(l2 == int32(9))|base.B2i32(base.Ui32(int32(10)) < base.Ui32(l2)) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessGUCArray[0]))
	v69 = v67
	goto L14
L13:
	;
	v69 = int32(10)
	goto L14
L14:
	;
	v71 = int32(0)
	v73 = F_set_config_with_handle(m, v60, int32(0), v65, l1, l2, v69, l3, int32(1), v71, v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_pfree(m, v60)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_pfree(m, v65)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v29 = v29 + int32(1)
	goto L3
L18:
	;
	F_list_free(m, v26)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	m.G0 = v13 + int32(16)
	return
}
func F_SplitGUCList(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v9 = l0
	goto L1
L1:
	;
	v17 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9))))
	goto L3
L2:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v27 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if base.B2i32(v17 == int32(32))|base.B2i32(base.Ui32((v17-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v9 = v9 + int32(1)
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
	v32 = v27
	v34 = v9
	goto L7
L7:
	;
	v37 = v32 & int32(255)
	if v37 != int32(34) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v168)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v171 = F_lappend(m, v170, v111)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L55
	} else {
		goto L56
	}
L10:
	;
	return int32(0)
L11:
	;
	v113 = v107
	goto L43
L12:
	;
	if v37 == int32(0) {
		v67 = v34
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v75 = v34 + int32(1)
	v76 = int32(34)
	v77 = F___strchrnul(m, v75, v76)
	mBase = m.M
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v79 == v76 {
		goto L28
	} else {
		goto L29
	}
L15:
	;
	if v34 == v67 {
		goto L10
	} else {
		goto L26
	}
L16:
	;
	if v37 == int32(44) {
		v67 = v34
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v46 = v32
	v47 = v34
	goto L18
L18:
	;
	v50 = base.I32_extend8_s(v46)
	goto L20
L19:
	;
	v67 = v61
	goto L15
L20:
	;
	if base.B2i32(v50 == int32(32))|base.B2i32(base.Ui32((v50-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v67 = v47
	goto L15
L22:
	;
	goto L23
L23:
	;
	v61 = v47 + int32(1)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	if v62 == int32(0) {
		v67 = v61
		goto L15
	} else {
		goto L24
	}
L24:
	;
	if v62 != int32(44) {
		v46 = v62
		v47 = v61
		goto L18
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	v107 = v67
	v109 = v67
	v111 = v34
	goto L11
L27:
	;
	if v83 == int32(0) {
		goto L10
	} else {
		goto L31
	}
L28:
	;
	v83 = v77
	goto L30
L29:
	;
	v83 = int32(0)
	goto L30
L30:
	;
	goto L27
L31:
	;
	v88 = v83
	goto L32
L32:
	;
	v93 = v88 + int32(1)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	if v94 != int32(34) {
		v107 = v93
		v109 = v88
		v111 = v75
		goto L11
	} else {
		goto L34
	}
L33:
	;
	goto L10
L34:
	;
	v97 = F_strlen(m, v88)
	mBase = m.M
	if v97 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	base.MemoryCopy(m, v88, v93, v97)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v99 = int32(34)
	v100 = F___strchrnul(m, v93, v99)
	mBase = m.M
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v102 == v99 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v106 != 0 {
		v88 = v106
		goto L32
	} else {
		goto L42
	}
L39:
	;
	v106 = v100
	goto L41
L40:
	;
	v106 = int32(0)
	goto L41
L41:
	;
	goto L38
L42:
	;
	goto L33
L43:
	;
	v121 = int32(*(*int8)(unsafe.Add(mBase, uint32(v113))))
	goto L45
L44:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v131 == int32(44) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if base.B2i32(v121 == int32(32))|base.B2i32(base.Ui32((v121-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v113 = v113 + int32(1)
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v137 = v113
	goto L50
L48:
	;
	goto L49
L49:
	;
	if v131 == int32(0) {
		v165 = v113
		goto L9
	} else {
		goto L54
	}
L50:
	;
	v141 = v137 + int32(1)
	v142 = int32(*(*int8)(unsafe.Add(mBase, uint32(v137)+1)))
	goto L52
L52:
	;
	if base.B2i32(v142 == int32(32))|base.B2i32(base.Ui32((v142-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v137 = v141
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v165 = v141
	goto L9
L54:
	;
	goto L10
L55:
	;
	return int32(0)
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v171
	if v131 != int32(44) {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	v32 = v178
	v34 = v165
	goto L7
}
