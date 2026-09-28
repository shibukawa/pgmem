package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_cube_consistent(m *base.Module, l0 int32) int64 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 float64
	_ = v77
	var v79 int32
	_ = v79
	var v83 float64
	_ = v83
	var v86 float64
	_ = v86
	var v87 int32
	_ = v87
	var v88 float64
	_ = v88
	var v90 int32
	_ = v90
	var v94 float64
	_ = v94
	var v97 float64
	_ = v97
	var v102 float64
	_ = v102
	var v104 float64
	_ = v104
	var v109 float64
	_ = v109
	var v111 float64
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 float64
	_ = v144
	var v146 int32
	_ = v146
	var v150 float64
	_ = v150
	var v153 float64
	_ = v153
	var v154 int32
	_ = v154
	var v155 float64
	_ = v155
	var v157 int32
	_ = v157
	var v161 float64
	_ = v161
	var v164 float64
	_ = v164
	var v169 float64
	_ = v169
	var v171 float64
	_ = v171
	var v176 float64
	_ = v176
	var v178 float64
	_ = v178
	var v182 int32
	_ = v182
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v227 int32
	_ = v227
	var v228 float64
	_ = v228
	var v235 float64
	_ = v235
	var v237 int32
	_ = v237
	var v241 float64
	_ = v241
	var v247 float64
	_ = v247
	var v253 float64
	_ = v253
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
	var v281 int32
	_ = v281
	var v282 float64
	_ = v282
	var v289 float64
	_ = v289
	var v291 int32
	_ = v291
	var v295 float64
	_ = v295
	var v301 float64
	_ = v301
	var v307 float64
	_ = v307
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v342 int32
	_ = v342
	var v343 float64
	_ = v343
	var v350 float64
	_ = v350
	var v352 int32
	_ = v352
	var v356 float64
	_ = v356
	var v368 float64
	_ = v368
	var v372 float64
	_ = v372
	var v376 int32
	_ = v376
	var v389 int32
	_ = v389
	var v400 int32
	_ = v400
	var v401 float64
	_ = v401
	var v408 float64
	_ = v408
	var v410 int32
	_ = v410
	var v414 float64
	_ = v414
	var v426 float64
	_ = v426
	var v430 float64
	_ = v430
	var v435 int32
	_ = v435
	var v467 int32
	_ = v467
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v558 int32
	_ = v558
	var v559 float64
	_ = v559
	var v567 float64
	_ = v567
	var v571 int32
	_ = v571
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 float64
	_ = v615
	var v617 int32
	_ = v617
	var v621 float64
	_ = v621
	var v624 float64
	_ = v624
	var v625 int32
	_ = v625
	var v626 float64
	_ = v626
	var v628 int32
	_ = v628
	var v632 float64
	_ = v632
	var v635 float64
	_ = v635
	var v640 float64
	_ = v640
	var v642 float64
	_ = v642
	var v647 float64
	_ = v647
	var v649 float64
	_ = v649
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v666 int32
	_ = v666
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v715 int32
	_ = v715
	var v729 int32
	_ = v729
	var v730 float64
	_ = v730
	var v738 float64
	_ = v738
	var v742 int32
	_ = v742
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 float64
	_ = v786
	var v788 int32
	_ = v788
	var v792 float64
	_ = v792
	var v795 float64
	_ = v795
	var v796 int32
	_ = v796
	var v797 float64
	_ = v797
	var v799 int32
	_ = v799
	var v803 float64
	_ = v803
	var v806 float64
	_ = v806
	var v811 float64
	_ = v811
	var v813 float64
	_ = v813
	var v818 float64
	_ = v818
	var v820 float64
	_ = v820
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v837 int32
	_ = v837
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v870 int32
	_ = v870
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 float64
	_ = v924
	var v926 int32
	_ = v926
	var v930 float64
	_ = v930
	var v933 float64
	_ = v933
	var v934 int32
	_ = v934
	var v935 float64
	_ = v935
	var v937 int32
	_ = v937
	var v941 float64
	_ = v941
	var v944 float64
	_ = v944
	var v949 float64
	_ = v949
	var v951 float64
	_ = v951
	var v956 float64
	_ = v956
	var v958 float64
	_ = v958
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v980 int32
	_ = v980
	var v992 int32
	_ = v992
	var v1004 int32
	_ = v1004
	var v1005 float64
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1013 float64
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1019 float64
	_ = v1019
	var v1025 float64
	_ = v1025
	var v1031 float64
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1050 int32
	_ = v1050
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1103 int32
	_ = v1103
	var v1117 int32
	_ = v1117
	var v1118 float64
	_ = v1118
	var v1126 float64
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1158 int32
	_ = v1158
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 float64
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1180 float64
	_ = v1180
	var v1183 float64
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 float64
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1191 float64
	_ = v1191
	var v1194 float64
	_ = v1194
	var v1199 float64
	_ = v1199
	var v1201 float64
	_ = v1201
	var v1206 float64
	_ = v1206
	var v1208 float64
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1225 int32
	_ = v1225
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v15 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v15)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v18)+12)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v20&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1246 != v9 {
		goto L286
	} else {
		goto L287
	}
L5:
	;
	v1073 = int32(0)
	if base.B2i32(v22 == v1073)|base.B2i32(v9 == v1073) != 0 {
		v1225 = v1073
		goto L250
	} else {
		goto L251
	}
L6:
	;
	v870 = int32(0)
	if base.B2i32(v22 == v870)|base.B2i32(v9 == v870) != 0 {
		v1050 = v870
		goto L202
	} else {
		goto L203
	}
L7:
	;
	switch v13 - int32(3) {
	case 0:
		goto L6
	default:
		v1245 = v2
		goto L4
	case 3:
		goto L12
	case 4, 10:
		goto L11
	case 5, 11:
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if base.Ui32(int32(14)) < base.Ui32(v13) {
		v1245 = v2
		goto L4
	} else {
		goto L197
	}
L10:
	;
	v685 = int32(0)
	if base.B2i32(v9 == v685)|base.B2i32(v22 == v685) != 0 {
		v837 = v685
		goto L161
	} else {
		goto L162
	}
L11:
	;
	v514 = int32(0)
	if base.B2i32(v22 == v514)|base.B2i32(v9 == v514) != 0 {
		v666 = v514
		goto L124
	} else {
		goto L125
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v45 = int32(2147483647)
	v46 = v44 & v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v49 = v47 & v45
	if base.Ui32(v46) < base.Ui32(v49) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v1245 = base.B2i32(v511 == int32(0))
	goto L4
L14:
	;
	v511 = int32(-1)
	goto L13
L15:
	;
	v511 = v467
	goto L13
L16:
	;
	v467 = int32(1)
	goto L15
L17:
	;
	v51 = v46
	goto L19
L18:
	;
	v51 = v49
	goto L19
L19:
	;
	if v51 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v52 = int32(8)
	v69 = int32(0)
	goto L23
L21:
	;
	goto L22
L22:
	;
	if base.Ui32(v49) < base.Ui32(v46) {
		goto L57
	} else {
		goto L58
	}
L23:
	;
	v75 = v69 << (uint(int32(3)) % 32)
	v76 = v22 + v52 + v75
	v77 = *(*float64)(unsafe.Add(mBase, uint32(v76)))
	v79 = base.B2i32(v44 < int32(0))
	if v44 < int32(0) {
		v86 = v77
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v118 = int32(8)
	v136 = int32(0)
	goto L40
L25:
	;
	v87 = v75 + (v9 + v52)
	v88 = *(*float64)(unsafe.Add(mBase, uint32(v87)))
	v90 = base.B2i32(v47 < int32(0))
	if v47 < int32(0) {
		v97 = v88
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v83 = *(*float64)(unsafe.Add(mBase, uint32(v76+v44<<(uint(int32(3))%32))))
	if base.F64_lt(v77, v83) != 0 {
		v86 = v77
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v86 = v83
	goto L25
L28:
	;
	if base.F64_gt(v86, v97) != 0 {
		goto L16
	} else {
		goto L31
	}
L29:
	;
	v94 = *(*float64)(unsafe.Add(mBase, uint32(v87+v47<<(uint(int32(3))%32))))
	if base.F64_lt(v88, v94) != 0 {
		v97 = v88
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v97 = v94
	goto L28
L31:
	;
	if v44 < int32(0) {
		v104 = v77
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v47 < int32(0) {
		v111 = v88
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v76+v44<<(uint(int32(3))%32))))
	if base.F64_lt(v77, v102) != 0 {
		v104 = v77
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v104 = v102
	goto L32
L35:
	;
	v113 = int32(-1)
	if base.F64_lt(v104, v111) != 0 {
		v467 = v113
		goto L15
	} else {
		goto L38
	}
L36:
	;
	v109 = *(*float64)(unsafe.Add(mBase, uint32(v87+v47<<(uint(int32(3))%32))))
	if base.F64_lt(v88, v109) != 0 {
		v111 = v88
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v111 = v109
	goto L35
L38:
	;
	v116 = v69 + int32(1)
	if v116 != v51 {
		v69 = v116
		goto L23
	} else {
		goto L39
	}
L39:
	;
	goto L24
L40:
	;
	v142 = v136 << (uint(int32(3)) % 32)
	v143 = v22 + v118 + v142
	v144 = *(*float64)(unsafe.Add(mBase, uint32(v143)))
	v146 = base.B2i32(v44 < int32(0))
	if v44 < int32(0) {
		v153 = v144
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L22
L42:
	;
	v154 = v142 + (v9 + v118)
	v155 = *(*float64)(unsafe.Add(mBase, uint32(v154)))
	v157 = base.B2i32(v47 < int32(0))
	if v47 < int32(0) {
		v164 = v155
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v150 = *(*float64)(unsafe.Add(mBase, uint32(v143+v44<<(uint(int32(3))%32))))
	if base.F64_gt(v144, v150) != 0 {
		v153 = v144
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v153 = v150
	goto L42
L45:
	;
	if base.F64_gt(v153, v164) != 0 {
		goto L16
	} else {
		goto L48
	}
L46:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v154+v47<<(uint(int32(3))%32))))
	if base.F64_gt(v155, v161) != 0 {
		v164 = v155
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v164 = v161
	goto L45
L48:
	;
	if v44 < int32(0) {
		v171 = v144
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v47 < int32(0) {
		v178 = v155
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v143+v44<<(uint(int32(3))%32))))
	if base.F64_gt(v144, v169) != 0 {
		v171 = v144
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v171 = v169
	goto L49
L52:
	;
	if base.F64_lt(v171, v178) != 0 {
		v467 = v113
		goto L15
	} else {
		goto L55
	}
L53:
	;
	v176 = *(*float64)(unsafe.Add(mBase, uint32(v154+v47<<(uint(int32(3))%32))))
	if base.F64_gt(v155, v176) != 0 {
		v178 = v155
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v178 = v176
	goto L52
L55:
	;
	v182 = v136 + int32(1)
	if v182 != v51 {
		v136 = v182
		goto L40
	} else {
		goto L56
	}
L56:
	;
	goto L41
L57:
	;
	v204 = v22 + int32(8)
	v207 = v51
	goto L60
L58:
	;
	goto L59
L59:
	;
	if base.Ui32(v49) <= base.Ui32(v46) {
		goto L90
	} else {
		goto L91
	}
L60:
	;
	v227 = v204 + v207<<(uint(int32(3))%32)
	v228 = *(*float64)(unsafe.Add(mBase, uint32(v227)))
	if base.B2i32(v44 < int32(0)) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v270 = v51
	goto L74
L62:
	;
	if base.F64_lt(v253, float64(0)) != 0 {
		goto L14
	} else {
		goto L72
	}
L63:
	;
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v227+v46<<(uint(int32(3))%32))))
	if base.F64_lt(v228, v235) != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	if base.F64_gt(v228, float64(0)) != 0 {
		goto L16
	} else {
		goto L71
	}
L66:
	;
	v237 = int32(0)
	goto L68
L67:
	;
	v237 = v44
	goto L68
L68:
	;
	v241 = *(*float64)(unsafe.Add(mBase, uint32(v227+v237<<(uint(int32(3))%32))))
	if base.F64_gt(v241, float64(0)) != 0 {
		goto L16
	} else {
		goto L69
	}
L69:
	;
	v247 = *(*float64)(unsafe.Add(mBase, uint32(v227+v44<<(uint(int32(3))%32))))
	if base.F64_lt(v228, v247) == int32(0) {
		v253 = v247
		goto L62
	} else {
		goto L70
	}
L70:
	;
	v253 = v228
	goto L62
L71:
	;
	v253 = v228
	goto L62
L72:
	;
	v257 = v207 + int32(1)
	if v257 != v46 {
		v207 = v257
		goto L60
	} else {
		goto L73
	}
L73:
	;
	goto L61
L74:
	;
	v281 = v204 + v270<<(uint(int32(3))%32)
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v281)))
	if base.B2i32(v44 < int32(0)) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L14
L76:
	;
	if base.F64_lt(v307, float64(0)) == int32(0) {
		goto L86
	} else {
		goto L87
	}
L77:
	;
	v289 = *(*float64)(unsafe.Add(mBase, uint32(v281+v46<<(uint(int32(3))%32))))
	if base.F64_gt(v282, v289) != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	if base.F64_gt(v282, float64(0)) != 0 {
		goto L16
	} else {
		goto L85
	}
L80:
	;
	v291 = int32(0)
	goto L82
L81:
	;
	v291 = v44
	goto L82
L82:
	;
	v295 = *(*float64)(unsafe.Add(mBase, uint32(v281+v291<<(uint(int32(3))%32))))
	if base.F64_gt(v295, float64(0)) != 0 {
		goto L16
	} else {
		goto L83
	}
L83:
	;
	v301 = *(*float64)(unsafe.Add(mBase, uint32(v281+v44<<(uint(int32(3))%32))))
	if base.F64_gt(v282, v301) == int32(0) {
		v307 = v301
		goto L76
	} else {
		goto L84
	}
L84:
	;
	v307 = v282
	goto L76
L85:
	;
	v307 = v282
	goto L76
L86:
	;
	v312 = int32(1)
	v314 = v270 + v312
	if v314 == v46 {
		v467 = v312
		goto L15
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	goto L75
L89:
	;
	v270 = v314
	goto L74
L90:
	;
	v511 = int32(0)
	goto L13
L91:
	;
	goto L92
L92:
	;
	v319 = v9 + int32(8)
	v329 = v46
	goto L93
L93:
	;
	v342 = v319 + v329<<(uint(int32(3))%32)
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v342)))
	if base.B2i32(v47 < int32(0)) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v389 = v51
	goto L108
L95:
	;
	if base.F64_lt(v372, float64(0)) != 0 {
		goto L16
	} else {
		goto L106
	}
L96:
	;
	v368 = *(*float64)(unsafe.Add(mBase, uint32(v342+v47<<(uint(int32(3))%32))))
	if base.F64_lt(v343, v368) == int32(0) {
		v372 = v368
		goto L95
	} else {
		goto L105
	}
L97:
	;
	v350 = *(*float64)(unsafe.Add(mBase, uint32(v342+v49<<(uint(int32(3))%32))))
	if base.F64_lt(v343, v350) != 0 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	if base.F64_gt(v343, float64(0)) == int32(0) {
		v372 = v343
		goto L95
	} else {
		goto L104
	}
L100:
	;
	v352 = int32(0)
	goto L102
L101:
	;
	v352 = v47
	goto L102
L102:
	;
	v356 = *(*float64)(unsafe.Add(mBase, uint32(v342+v352<<(uint(int32(3))%32))))
	if base.F64_gt(v356, float64(0)) == int32(0) {
		goto L96
	} else {
		goto L103
	}
L103:
	;
	goto L14
L104:
	;
	goto L14
L105:
	;
	v372 = v343
	goto L95
L106:
	;
	v376 = v329 + int32(1)
	if v376 != v49 {
		v329 = v376
		goto L93
	} else {
		goto L107
	}
L107:
	;
	goto L94
L108:
	;
	v400 = v319 + v389<<(uint(int32(3))%32)
	v401 = *(*float64)(unsafe.Add(mBase, uint32(v400)))
	if base.B2i32(v47 < int32(0)) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v467 = int32(-1)
	goto L15
L110:
	;
	if base.F64_lt(v430, float64(0)) != 0 {
		goto L16
	} else {
		goto L121
	}
L111:
	;
	v426 = *(*float64)(unsafe.Add(mBase, uint32(v400+v47<<(uint(int32(3))%32))))
	if base.F64_gt(v401, v426) == int32(0) {
		v430 = v426
		goto L110
	} else {
		goto L120
	}
L112:
	;
	v408 = *(*float64)(unsafe.Add(mBase, uint32(v400+v49<<(uint(int32(3))%32))))
	if base.F64_gt(v401, v408) != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L114
L114:
	;
	if base.F64_gt(v401, float64(0)) == int32(0) {
		v430 = v401
		goto L110
	} else {
		goto L119
	}
L115:
	;
	v410 = int32(0)
	goto L117
L116:
	;
	v410 = v47
	goto L117
L117:
	;
	v414 = *(*float64)(unsafe.Add(mBase, uint32(v400+v410<<(uint(int32(3))%32))))
	if base.F64_gt(v414, float64(0)) == int32(0) {
		goto L111
	} else {
		goto L118
	}
L118:
	;
	goto L14
L119:
	;
	goto L14
L120:
	;
	v430 = v401
	goto L110
L121:
	;
	v435 = v389 + int32(1)
	if v49 != v435 {
		v389 = v435
		goto L108
	} else {
		goto L122
	}
L122:
	;
	goto L109
L123:
	;
	v1245 = v684
	goto L4
L124:
	;
	v684 = v666
	goto L123
L125:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v532 = int32(2147483647)
	v533 = v531 & v532
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v536 = v534 & v532
	if base.Ui32(v533) < base.Ui32(v536) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v544 = v533
	goto L129
L127:
	;
	goto L128
L128:
	;
	if base.Ui32(v533) < base.Ui32(v536) {
		goto L137
	} else {
		goto L138
	}
L129:
	;
	v558 = v9 + int32(8) + v544<<(uint(int32(3))%32)
	v559 = *(*float64)(unsafe.Add(mBase, uint32(v558)))
	if base.F64_ne(v559, float64(0)) != 0 {
		v666 = v514
		goto L124
	} else {
		goto L131
	}
L130:
	;
	goto L128
L131:
	;
	if base.B2i32(v534 < int32(0)) == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v567 = *(*float64)(unsafe.Add(mBase, uint32(v558+v536<<(uint(int32(3))%32))))
	if base.F64_ne(v567, float64(0)) != 0 {
		v666 = v514
		goto L124
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v571 = v544 + int32(1)
	if v571 != v536 {
		v544 = v571
		goto L129
	} else {
		goto L136
	}
L135:
	;
	goto L134
L136:
	;
	goto L130
L137:
	;
	v588 = v533
	goto L139
L138:
	;
	v588 = v536
	goto L139
L139:
	;
	if v588 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v684 = int32(1)
	goto L123
L141:
	;
	goto L142
L142:
	;
	v592 = int32(8)
	v599 = int32(0)
	goto L143
L143:
	;
	v611 = int32(0)
	v613 = v599 << (uint(int32(3)) % 32)
	v614 = v22 + v592 + v613
	v615 = *(*float64)(unsafe.Add(mBase, uint32(v614)))
	v617 = base.B2i32(v531 < v611)
	if v531 < v611 {
		v624 = v615
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v666 = v652
	goto L124
L145:
	;
	v625 = v613 + (v9 + v592)
	v626 = *(*float64)(unsafe.Add(mBase, uint32(v625)))
	v628 = base.B2i32(v534 < int32(0))
	if v534 < int32(0) {
		v635 = v626
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v621 = *(*float64)(unsafe.Add(mBase, uint32(v614+v531<<(uint(int32(3))%32))))
	if base.F64_lt(v615, v621) != 0 {
		v624 = v615
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v624 = v621
	goto L145
L148:
	;
	if base.F64_gt(v624, v635) != 0 {
		v666 = v611
		goto L124
	} else {
		goto L151
	}
L149:
	;
	v632 = *(*float64)(unsafe.Add(mBase, uint32(v625+v534<<(uint(int32(3))%32))))
	if base.F64_lt(v626, v632) != 0 {
		v635 = v626
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v635 = v632
	goto L148
L151:
	;
	if v531 < v611 {
		v642 = v615
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if v534 < int32(0) {
		v649 = v626
		goto L155
	} else {
		goto L156
	}
L153:
	;
	v640 = *(*float64)(unsafe.Add(mBase, uint32(v614+v531<<(uint(int32(3))%32))))
	if base.F64_gt(v615, v640) != 0 {
		v642 = v615
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v642 = v640
	goto L152
L155:
	;
	if base.F64_gt(v649, v642) != 0 {
		v666 = v611
		goto L124
	} else {
		goto L158
	}
L156:
	;
	v647 = *(*float64)(unsafe.Add(mBase, uint32(v625+v534<<(uint(int32(3))%32))))
	if base.F64_gt(v626, v647) != 0 {
		v649 = v626
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v649 = v647
	goto L155
L158:
	;
	v652 = int32(1)
	v654 = v599 + v652
	if v654 != v588 {
		v599 = v654
		goto L143
	} else {
		goto L159
	}
L159:
	;
	goto L144
L160:
	;
	v1245 = v855
	goto L4
L161:
	;
	v855 = v837
	goto L160
L162:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v703 = int32(2147483647)
	v704 = v702 & v703
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v707 = v705 & v703
	if base.Ui32(v704) < base.Ui32(v707) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v715 = v704
	goto L166
L164:
	;
	goto L165
L165:
	;
	if base.Ui32(v704) < base.Ui32(v707) {
		goto L174
	} else {
		goto L175
	}
L166:
	;
	v729 = v22 + int32(8) + v715<<(uint(int32(3))%32)
	v730 = *(*float64)(unsafe.Add(mBase, uint32(v729)))
	if base.F64_ne(v730, float64(0)) != 0 {
		v837 = v685
		goto L161
	} else {
		goto L168
	}
L167:
	;
	goto L165
L168:
	;
	if base.B2i32(v705 < int32(0)) == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v738 = *(*float64)(unsafe.Add(mBase, uint32(v729+v707<<(uint(int32(3))%32))))
	if base.F64_ne(v738, float64(0)) != 0 {
		v837 = v685
		goto L161
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v742 = v715 + int32(1)
	if v742 != v707 {
		v715 = v742
		goto L166
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	goto L167
L174:
	;
	v759 = v704
	goto L176
L175:
	;
	v759 = v707
	goto L176
L176:
	;
	if v759 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v855 = int32(1)
	goto L160
L178:
	;
	goto L179
L179:
	;
	v763 = int32(8)
	v770 = int32(0)
	goto L180
L180:
	;
	v782 = int32(0)
	v784 = v770 << (uint(int32(3)) % 32)
	v785 = v9 + v763 + v784
	v786 = *(*float64)(unsafe.Add(mBase, uint32(v785)))
	v788 = base.B2i32(v702 < v782)
	if v702 < v782 {
		v795 = v786
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v837 = v823
	goto L161
L182:
	;
	v796 = v784 + (v22 + v763)
	v797 = *(*float64)(unsafe.Add(mBase, uint32(v796)))
	v799 = base.B2i32(v705 < int32(0))
	if v705 < int32(0) {
		v806 = v797
		goto L185
	} else {
		goto L186
	}
L183:
	;
	v792 = *(*float64)(unsafe.Add(mBase, uint32(v785+v702<<(uint(int32(3))%32))))
	if base.F64_lt(v786, v792) != 0 {
		v795 = v786
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v795 = v792
	goto L182
L185:
	;
	if base.F64_gt(v795, v806) != 0 {
		v837 = v782
		goto L161
	} else {
		goto L188
	}
L186:
	;
	v803 = *(*float64)(unsafe.Add(mBase, uint32(v796+v705<<(uint(int32(3))%32))))
	if base.F64_lt(v797, v803) != 0 {
		v806 = v797
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v806 = v803
	goto L185
L188:
	;
	if v702 < v782 {
		v813 = v786
		goto L189
	} else {
		goto L190
	}
L189:
	;
	if v705 < int32(0) {
		v820 = v797
		goto L192
	} else {
		goto L193
	}
L190:
	;
	v811 = *(*float64)(unsafe.Add(mBase, uint32(v785+v702<<(uint(int32(3))%32))))
	if base.F64_gt(v786, v811) != 0 {
		v813 = v786
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v813 = v811
	goto L189
L192:
	;
	if base.F64_gt(v820, v813) != 0 {
		v837 = v782
		goto L161
	} else {
		goto L195
	}
L193:
	;
	v818 = *(*float64)(unsafe.Add(mBase, uint32(v796+v705<<(uint(int32(3))%32))))
	if base.F64_gt(v797, v818) != 0 {
		v820 = v797
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v820 = v818
	goto L192
L195:
	;
	v823 = int32(1)
	v825 = v770 + v823
	if v825 != v759 {
		v770 = v825
		goto L180
	} else {
		goto L196
	}
L196:
	;
	goto L181
L197:
	;
	v859 = int32(1) << (uint(v13) % 32)
	if v859&int32(_a_F_g_cube_consistent_0) != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	if v859&int32(_a_F_g_cube_consistent_1) != 0 {
		goto L6
	} else {
		goto L199
	}
L199:
	;
	if v13 != int32(3) {
		v1245 = v2
		goto L4
	} else {
		goto L200
	}
L200:
	;
	goto L6
L201:
	;
	v1245 = v1072
	goto L4
L202:
	;
	v1072 = v1050
	goto L201
L203:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v886 = int32(2147483647)
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v891 = base.B2i32(base.Ui32(v885&v886) < base.Ui32(v888&v886))
	if base.Ui32(v885&v886) < base.Ui32(v888&v886) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v892 = v9
	goto L206
L205:
	;
	v892 = v22
	goto L206
L206:
	;
	if base.Ui32(v885&v886) < base.Ui32(v888&v886) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v980 = v964 & int32(2147483647)
	if base.Ui32(v980) <= base.Ui32(v896) {
		goto L231
	} else {
		goto L232
	}
L208:
	;
	v893 = v22
	goto L210
L209:
	;
	v893 = v9
	goto L210
L210:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v893)+4))
	v896 = v894 & int32(2147483647)
	if v896 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v892)+4))
	v964 = v899
	goto L207
L212:
	;
	goto L213
L213:
	;
	v900 = int32(8)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v892)+4))
	v907 = int32(0)
	goto L214
L214:
	;
	v922 = v907 << (uint(int32(3)) % 32)
	v923 = v892 + v900 + v922
	v924 = *(*float64)(unsafe.Add(mBase, uint32(v923)))
	v926 = base.B2i32(v904 < int32(0))
	if v904 < int32(0) {
		v933 = v924
		goto L216
	} else {
		goto L217
	}
L215:
	;
	v964 = v904
	goto L207
L216:
	;
	v934 = v922 + (v893 + v900)
	v935 = *(*float64)(unsafe.Add(mBase, uint32(v934)))
	v937 = base.B2i32(v894 < int32(0))
	if v894 < int32(0) {
		v944 = v935
		goto L219
	} else {
		goto L220
	}
L217:
	;
	v930 = *(*float64)(unsafe.Add(mBase, uint32(v923+v904<<(uint(int32(3))%32))))
	if base.F64_lt(v924, v930) != 0 {
		v933 = v924
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v933 = v930
	goto L216
L219:
	;
	if base.F64_gt(v933, v944) != 0 {
		v1050 = v870
		goto L202
	} else {
		goto L222
	}
L220:
	;
	v941 = *(*float64)(unsafe.Add(mBase, uint32(v934+v894<<(uint(int32(3))%32))))
	if base.F64_gt(v935, v941) != 0 {
		v944 = v935
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v944 = v941
	goto L219
L222:
	;
	if v904 < int32(0) {
		v951 = v924
		goto L223
	} else {
		goto L224
	}
L223:
	;
	if v894 < int32(0) {
		v958 = v935
		goto L226
	} else {
		goto L227
	}
L224:
	;
	v949 = *(*float64)(unsafe.Add(mBase, uint32(v923+v904<<(uint(int32(3))%32))))
	if base.F64_gt(v924, v949) != 0 {
		v951 = v924
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v951 = v949
	goto L223
L226:
	;
	if base.F64_gt(v958, v951) != 0 {
		v1050 = v870
		goto L202
	} else {
		goto L229
	}
L227:
	;
	v956 = *(*float64)(unsafe.Add(mBase, uint32(v934+v894<<(uint(int32(3))%32))))
	if base.F64_lt(v935, v956) != 0 {
		v958 = v935
		goto L226
	} else {
		goto L228
	}
L228:
	;
	v958 = v956
	goto L226
L229:
	;
	v962 = v907 + int32(1)
	if v962 != v896 {
		v907 = v962
		goto L214
	} else {
		goto L230
	}
L230:
	;
	goto L215
L231:
	;
	v1072 = int32(1)
	goto L201
L232:
	;
	goto L233
L233:
	;
	v992 = v896
	goto L234
L234:
	;
	v1004 = v892 + int32(8) + v992<<(uint(int32(3))%32)
	v1005 = *(*float64)(unsafe.Add(mBase, uint32(v1004)))
	if base.B2i32(v964 < int32(0)) == int32(0) {
		goto L238
	} else {
		goto L239
	}
L235:
	;
	v1050 = int32(0)
	goto L202
L236:
	;
	goto L235
L237:
	;
	if base.F64_lt(v1031, float64(0)) != 0 {
		goto L236
	} else {
		goto L247
	}
L238:
	;
	v1008 = int32(0)
	v1013 = *(*float64)(unsafe.Add(mBase, uint32(v1004+v980<<(uint(int32(3))%32))))
	if base.F64_lt(v1005, v1013) != 0 {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	goto L240
L240:
	;
	if base.F64_gt(v1005, float64(0)) != 0 {
		goto L236
	} else {
		goto L246
	}
L241:
	;
	v1015 = v1008
	goto L243
L242:
	;
	v1015 = v964
	goto L243
L243:
	;
	v1019 = *(*float64)(unsafe.Add(mBase, uint32(v1004+v1015<<(uint(int32(3))%32))))
	if base.F64_gt(v1019, float64(0)) != 0 {
		v1050 = v1008
		goto L202
	} else {
		goto L244
	}
L244:
	;
	v1025 = *(*float64)(unsafe.Add(mBase, uint32(v1004+v964<<(uint(int32(3))%32))))
	if base.F64_gt(v1005, v1025) == int32(0) {
		v1031 = v1025
		goto L237
	} else {
		goto L245
	}
L245:
	;
	v1031 = v1005
	goto L237
L246:
	;
	v1031 = v1005
	goto L237
L247:
	;
	v1035 = int32(1)
	v1037 = v992 + v1035
	if v980 != v1037 {
		v992 = v1037
		goto L234
	} else {
		goto L248
	}
L248:
	;
	v1050 = v1035
	goto L202
L249:
	;
	v1245 = v1243
	goto L4
L250:
	;
	v1243 = v1225
	goto L249
L251:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v1091 = int32(2147483647)
	v1092 = v1090 & v1091
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v1095 = v1093 & v1091
	if base.Ui32(v1092) < base.Ui32(v1095) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1103 = v1092
	goto L255
L253:
	;
	goto L254
L254:
	;
	if base.Ui32(v1092) < base.Ui32(v1095) {
		goto L263
	} else {
		goto L264
	}
L255:
	;
	v1117 = v9 + int32(8) + v1103<<(uint(int32(3))%32)
	v1118 = *(*float64)(unsafe.Add(mBase, uint32(v1117)))
	if base.F64_ne(v1118, float64(0)) != 0 {
		v1225 = v1073
		goto L250
	} else {
		goto L257
	}
L256:
	;
	goto L254
L257:
	;
	if base.B2i32(v1093 < int32(0)) == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1126 = *(*float64)(unsafe.Add(mBase, uint32(v1117+v1095<<(uint(int32(3))%32))))
	if base.F64_ne(v1126, float64(0)) != 0 {
		v1225 = v1073
		goto L250
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v1130 = v1103 + int32(1)
	if v1130 != v1095 {
		v1103 = v1130
		goto L255
	} else {
		goto L262
	}
L261:
	;
	goto L260
L262:
	;
	goto L256
L263:
	;
	v1147 = v1092
	goto L265
L264:
	;
	v1147 = v1095
	goto L265
L265:
	;
	if v1147 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1243 = int32(1)
	goto L249
L267:
	;
	goto L268
L268:
	;
	v1151 = int32(8)
	v1158 = int32(0)
	goto L269
L269:
	;
	v1170 = int32(0)
	v1172 = v1158 << (uint(int32(3)) % 32)
	v1173 = v22 + v1151 + v1172
	v1174 = *(*float64)(unsafe.Add(mBase, uint32(v1173)))
	v1176 = base.B2i32(v1090 < v1170)
	if v1090 < v1170 {
		v1183 = v1174
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1225 = v1211
	goto L250
L271:
	;
	v1184 = v1172 + (v9 + v1151)
	v1185 = *(*float64)(unsafe.Add(mBase, uint32(v1184)))
	v1187 = base.B2i32(v1093 < int32(0))
	if v1093 < int32(0) {
		v1194 = v1185
		goto L274
	} else {
		goto L275
	}
L272:
	;
	v1180 = *(*float64)(unsafe.Add(mBase, uint32(v1173+v1090<<(uint(int32(3))%32))))
	if base.F64_lt(v1174, v1180) != 0 {
		v1183 = v1174
		goto L271
	} else {
		goto L273
	}
L273:
	;
	v1183 = v1180
	goto L271
L274:
	;
	if base.F64_gt(v1183, v1194) != 0 {
		v1225 = v1170
		goto L250
	} else {
		goto L277
	}
L275:
	;
	v1191 = *(*float64)(unsafe.Add(mBase, uint32(v1184+v1093<<(uint(int32(3))%32))))
	if base.F64_lt(v1185, v1191) != 0 {
		v1194 = v1185
		goto L274
	} else {
		goto L276
	}
L276:
	;
	v1194 = v1191
	goto L274
L277:
	;
	if v1090 < v1170 {
		v1201 = v1174
		goto L278
	} else {
		goto L279
	}
L278:
	;
	if v1093 < int32(0) {
		v1208 = v1185
		goto L281
	} else {
		goto L282
	}
L279:
	;
	v1199 = *(*float64)(unsafe.Add(mBase, uint32(v1173+v1090<<(uint(int32(3))%32))))
	if base.F64_gt(v1174, v1199) != 0 {
		v1201 = v1174
		goto L278
	} else {
		goto L280
	}
L280:
	;
	v1201 = v1199
	goto L278
L281:
	;
	if base.F64_gt(v1208, v1201) != 0 {
		v1225 = v1170
		goto L250
	} else {
		goto L284
	}
L282:
	;
	v1206 = *(*float64)(unsafe.Add(mBase, uint32(v1184+v1093<<(uint(int32(3))%32))))
	if base.F64_gt(v1185, v1206) != 0 {
		v1208 = v1185
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1208 = v1206
	goto L281
L284:
	;
	v1211 = int32(1)
	v1213 = v1158 + v1211
	if v1213 != v1147 {
		v1158 = v1213
		goto L269
	} else {
		goto L285
	}
L285:
	;
	goto L270
L286:
	;
	F_pfree(m, v9)
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L1
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	return base.I64_extend_i32_u(v1245)
L289:
	;
	goto L288
}
