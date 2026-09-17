package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_cube_consistent(m *base.Module, l0 int32) int32 {
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 float64
	_ = v79
	var v81 int32
	_ = v81
	var v85 float64
	_ = v85
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v90 float64
	_ = v90
	var v92 int32
	_ = v92
	var v96 float64
	_ = v96
	var v99 float64
	_ = v99
	var v104 float64
	_ = v104
	var v106 float64
	_ = v106
	var v111 float64
	_ = v111
	var v113 float64
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 float64
	_ = v146
	var v148 int32
	_ = v148
	var v152 float64
	_ = v152
	var v155 float64
	_ = v155
	var v156 int32
	_ = v156
	var v157 float64
	_ = v157
	var v159 int32
	_ = v159
	var v163 float64
	_ = v163
	var v166 float64
	_ = v166
	var v171 float64
	_ = v171
	var v173 float64
	_ = v173
	var v178 float64
	_ = v178
	var v180 float64
	_ = v180
	var v184 int32
	_ = v184
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v229 int32
	_ = v229
	var v230 float64
	_ = v230
	var v237 float64
	_ = v237
	var v239 int32
	_ = v239
	var v243 float64
	_ = v243
	var v249 float64
	_ = v249
	var v255 float64
	_ = v255
	var v259 int32
	_ = v259
	var v272 int32
	_ = v272
	var v283 int32
	_ = v283
	var v284 float64
	_ = v284
	var v291 float64
	_ = v291
	var v293 int32
	_ = v293
	var v297 float64
	_ = v297
	var v303 float64
	_ = v303
	var v309 float64
	_ = v309
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v331 int32
	_ = v331
	var v344 int32
	_ = v344
	var v345 float64
	_ = v345
	var v352 float64
	_ = v352
	var v354 int32
	_ = v354
	var v358 float64
	_ = v358
	var v370 float64
	_ = v370
	var v374 float64
	_ = v374
	var v378 int32
	_ = v378
	var v391 int32
	_ = v391
	var v402 int32
	_ = v402
	var v403 float64
	_ = v403
	var v410 float64
	_ = v410
	var v412 int32
	_ = v412
	var v416 float64
	_ = v416
	var v428 float64
	_ = v428
	var v432 float64
	_ = v432
	var v437 int32
	_ = v437
	var v469 int32
	_ = v469
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v560 int32
	_ = v560
	var v561 float64
	_ = v561
	var v569 float64
	_ = v569
	var v573 int32
	_ = v573
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 float64
	_ = v617
	var v619 int32
	_ = v619
	var v623 float64
	_ = v623
	var v626 float64
	_ = v626
	var v627 int32
	_ = v627
	var v628 float64
	_ = v628
	var v630 int32
	_ = v630
	var v634 float64
	_ = v634
	var v637 float64
	_ = v637
	var v642 float64
	_ = v642
	var v644 float64
	_ = v644
	var v649 float64
	_ = v649
	var v651 float64
	_ = v651
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v668 int32
	_ = v668
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v717 int32
	_ = v717
	var v731 int32
	_ = v731
	var v732 float64
	_ = v732
	var v740 float64
	_ = v740
	var v744 int32
	_ = v744
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v772 int32
	_ = v772
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 float64
	_ = v788
	var v790 int32
	_ = v790
	var v794 float64
	_ = v794
	var v797 float64
	_ = v797
	var v798 int32
	_ = v798
	var v799 float64
	_ = v799
	var v801 int32
	_ = v801
	var v805 float64
	_ = v805
	var v808 float64
	_ = v808
	var v813 float64
	_ = v813
	var v815 float64
	_ = v815
	var v820 float64
	_ = v820
	var v822 float64
	_ = v822
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v839 int32
	_ = v839
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v875 int32
	_ = v875
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 float64
	_ = v929
	var v931 int32
	_ = v931
	var v935 float64
	_ = v935
	var v938 float64
	_ = v938
	var v939 int32
	_ = v939
	var v940 float64
	_ = v940
	var v942 int32
	_ = v942
	var v946 float64
	_ = v946
	var v949 float64
	_ = v949
	var v954 float64
	_ = v954
	var v956 float64
	_ = v956
	var v961 float64
	_ = v961
	var v963 float64
	_ = v963
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v985 int32
	_ = v985
	var v997 int32
	_ = v997
	var v1009 int32
	_ = v1009
	var v1010 float64
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1018 float64
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1024 float64
	_ = v1024
	var v1030 float64
	_ = v1030
	var v1036 float64
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1055 int32
	_ = v1055
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1108 int32
	_ = v1108
	var v1122 int32
	_ = v1122
	var v1123 float64
	_ = v1123
	var v1131 float64
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1163 int32
	_ = v1163
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 float64
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1185 float64
	_ = v1185
	var v1188 float64
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 float64
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1196 float64
	_ = v1196
	var v1199 float64
	_ = v1199
	var v1204 float64
	_ = v1204
	var v1206 float64
	_ = v1206
	var v1211 float64
	_ = v1211
	var v1213 float64
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1230 int32
	_ = v1230
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v15 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v15)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
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
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1252 != v9 {
		goto L286
	} else {
		goto L287
	}
L5:
	;
	v1078 = int32(0)
	if base.B2i32(v22 == v1078)|base.B2i32(v9 == v1078) != 0 {
		v1230 = v1078
		goto L250
	} else {
		goto L251
	}
L6:
	;
	v875 = int32(0)
	if base.B2i32(v22 == v875)|base.B2i32(v9 == v875) != 0 {
		v1055 = v875
		goto L202
	} else {
		goto L203
	}
L7:
	;
	switch v13&int32(_a_F_g_cube_consistent_0) - int32(3) {
	case 0:
		goto L6
	default:
		v1251 = v2
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
	v859 = v13 & int32(_a_F_g_cube_consistent_0)
	if base.Ui32(int32(14)) < base.Ui32(v859) {
		v1251 = v2
		goto L4
	} else {
		goto L197
	}
L10:
	;
	v687 = int32(0)
	if base.B2i32(v9 == v687)|base.B2i32(v22 == v687) != 0 {
		v839 = v687
		goto L161
	} else {
		goto L162
	}
L11:
	;
	v516 = int32(0)
	if base.B2i32(v22 == v516)|base.B2i32(v9 == v516) != 0 {
		v668 = v516
		goto L124
	} else {
		goto L125
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v47 = int32(2147483647)
	v48 = v46 & v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v51 = v49 & v47
	if base.Ui32(v48) < base.Ui32(v51) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v1251 = base.B2i32(v513 == int32(0))
	goto L4
L14:
	;
	v513 = int32(-1)
	goto L13
L15:
	;
	v513 = v469
	goto L13
L16:
	;
	v469 = int32(1)
	goto L15
L17:
	;
	v53 = v48
	goto L19
L18:
	;
	v53 = v51
	goto L19
L19:
	;
	if v53 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v54 = int32(8)
	v71 = int32(0)
	goto L23
L21:
	;
	goto L22
L22:
	;
	if base.Ui32(v51) < base.Ui32(v48) {
		goto L57
	} else {
		goto L58
	}
L23:
	;
	v77 = v71 << (uint(int32(3)) % 32)
	v78 = v22 + v54 + v77
	v79 = *(*float64)(unsafe.Add(mBase, uint32(v78)))
	v81 = base.B2i32(v46 < int32(0))
	if v46 < int32(0) {
		v88 = v79
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v120 = int32(8)
	v138 = int32(0)
	goto L40
L25:
	;
	v89 = v77 + (v9 + v54)
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v89)))
	v92 = base.B2i32(v49 < int32(0))
	if v49 < int32(0) {
		v99 = v90
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v85 = *(*float64)(unsafe.Add(mBase, uint32(v78+v46<<(uint(int32(3))%32))))
	if base.F64_lt(v79, v85) != 0 {
		v88 = v79
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v88 = v85
	goto L25
L28:
	;
	if base.F64_gt(v88, v99) != 0 {
		goto L16
	} else {
		goto L31
	}
L29:
	;
	v96 = *(*float64)(unsafe.Add(mBase, uint32(v89+v49<<(uint(int32(3))%32))))
	if base.F64_lt(v90, v96) != 0 {
		v99 = v90
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v99 = v96
	goto L28
L31:
	;
	if v46 < int32(0) {
		v106 = v79
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v49 < int32(0) {
		v113 = v90
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v78+v46<<(uint(int32(3))%32))))
	if base.F64_lt(v79, v104) != 0 {
		v106 = v79
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v106 = v104
	goto L32
L35:
	;
	v115 = int32(-1)
	if base.F64_lt(v106, v113) != 0 {
		v469 = v115
		goto L15
	} else {
		goto L38
	}
L36:
	;
	v111 = *(*float64)(unsafe.Add(mBase, uint32(v89+v49<<(uint(int32(3))%32))))
	if base.F64_lt(v90, v111) != 0 {
		v113 = v90
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v113 = v111
	goto L35
L38:
	;
	v118 = v71 + int32(1)
	if v118 != v53 {
		v71 = v118
		goto L23
	} else {
		goto L39
	}
L39:
	;
	goto L24
L40:
	;
	v144 = v138 << (uint(int32(3)) % 32)
	v145 = v22 + v120 + v144
	v146 = *(*float64)(unsafe.Add(mBase, uint32(v145)))
	v148 = base.B2i32(v46 < int32(0))
	if v46 < int32(0) {
		v155 = v146
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L22
L42:
	;
	v156 = v144 + (v9 + v120)
	v157 = *(*float64)(unsafe.Add(mBase, uint32(v156)))
	v159 = base.B2i32(v49 < int32(0))
	if v49 < int32(0) {
		v166 = v157
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v152 = *(*float64)(unsafe.Add(mBase, uint32(v145+v46<<(uint(int32(3))%32))))
	if base.F64_gt(v146, v152) != 0 {
		v155 = v146
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v155 = v152
	goto L42
L45:
	;
	if base.F64_gt(v155, v166) != 0 {
		goto L16
	} else {
		goto L48
	}
L46:
	;
	v163 = *(*float64)(unsafe.Add(mBase, uint32(v156+v49<<(uint(int32(3))%32))))
	if base.F64_gt(v157, v163) != 0 {
		v166 = v157
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v166 = v163
	goto L45
L48:
	;
	if v46 < int32(0) {
		v173 = v146
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v49 < int32(0) {
		v180 = v157
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v171 = *(*float64)(unsafe.Add(mBase, uint32(v145+v46<<(uint(int32(3))%32))))
	if base.F64_gt(v146, v171) != 0 {
		v173 = v146
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v173 = v171
	goto L49
L52:
	;
	if base.F64_lt(v173, v180) != 0 {
		v469 = v115
		goto L15
	} else {
		goto L55
	}
L53:
	;
	v178 = *(*float64)(unsafe.Add(mBase, uint32(v156+v49<<(uint(int32(3))%32))))
	if base.F64_gt(v157, v178) != 0 {
		v180 = v157
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v180 = v178
	goto L52
L55:
	;
	v184 = v138 + int32(1)
	if v184 != v53 {
		v138 = v184
		goto L40
	} else {
		goto L56
	}
L56:
	;
	goto L41
L57:
	;
	v206 = v22 + int32(8)
	v209 = v53
	goto L60
L58:
	;
	goto L59
L59:
	;
	if base.Ui32(v51) <= base.Ui32(v48) {
		goto L90
	} else {
		goto L91
	}
L60:
	;
	v229 = v206 + v209<<(uint(int32(3))%32)
	v230 = *(*float64)(unsafe.Add(mBase, uint32(v229)))
	if base.B2i32(v46 < int32(0)) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v272 = v53
	goto L74
L62:
	;
	if base.F64_lt(v255, float64(0)) != 0 {
		goto L14
	} else {
		goto L72
	}
L63:
	;
	v237 = *(*float64)(unsafe.Add(mBase, uint32(v229+v48<<(uint(int32(3))%32))))
	if base.F64_lt(v230, v237) != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	if base.F64_gt(v230, float64(0)) != 0 {
		goto L16
	} else {
		goto L71
	}
L66:
	;
	v239 = int32(0)
	goto L68
L67:
	;
	v239 = v46
	goto L68
L68:
	;
	v243 = *(*float64)(unsafe.Add(mBase, uint32(v229+v239<<(uint(int32(3))%32))))
	if base.F64_gt(v243, float64(0)) != 0 {
		goto L16
	} else {
		goto L69
	}
L69:
	;
	v249 = *(*float64)(unsafe.Add(mBase, uint32(v229+v46<<(uint(int32(3))%32))))
	if base.F64_lt(v230, v249) == int32(0) {
		v255 = v249
		goto L62
	} else {
		goto L70
	}
L70:
	;
	v255 = v230
	goto L62
L71:
	;
	v255 = v230
	goto L62
L72:
	;
	v259 = v209 + int32(1)
	if v259 != v48 {
		v209 = v259
		goto L60
	} else {
		goto L73
	}
L73:
	;
	goto L61
L74:
	;
	v283 = v206 + v272<<(uint(int32(3))%32)
	v284 = *(*float64)(unsafe.Add(mBase, uint32(v283)))
	if base.B2i32(v46 < int32(0)) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L14
L76:
	;
	if base.F64_lt(v309, float64(0)) == int32(0) {
		goto L86
	} else {
		goto L87
	}
L77:
	;
	v291 = *(*float64)(unsafe.Add(mBase, uint32(v283+v48<<(uint(int32(3))%32))))
	if base.F64_gt(v284, v291) != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	if base.F64_gt(v284, float64(0)) != 0 {
		goto L16
	} else {
		goto L85
	}
L80:
	;
	v293 = int32(0)
	goto L82
L81:
	;
	v293 = v46
	goto L82
L82:
	;
	v297 = *(*float64)(unsafe.Add(mBase, uint32(v283+v293<<(uint(int32(3))%32))))
	if base.F64_gt(v297, float64(0)) != 0 {
		goto L16
	} else {
		goto L83
	}
L83:
	;
	v303 = *(*float64)(unsafe.Add(mBase, uint32(v283+v46<<(uint(int32(3))%32))))
	if base.F64_gt(v284, v303) == int32(0) {
		v309 = v303
		goto L76
	} else {
		goto L84
	}
L84:
	;
	v309 = v284
	goto L76
L85:
	;
	v309 = v284
	goto L76
L86:
	;
	v314 = int32(1)
	v316 = v272 + v314
	if v316 == v48 {
		v469 = v314
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
	v272 = v316
	goto L74
L90:
	;
	v513 = int32(0)
	goto L13
L91:
	;
	goto L92
L92:
	;
	v321 = v9 + int32(8)
	v331 = v48
	goto L93
L93:
	;
	v344 = v321 + v331<<(uint(int32(3))%32)
	v345 = *(*float64)(unsafe.Add(mBase, uint32(v344)))
	if base.B2i32(v49 < int32(0)) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v391 = v53
	goto L108
L95:
	;
	if base.F64_lt(v374, float64(0)) != 0 {
		goto L16
	} else {
		goto L106
	}
L96:
	;
	v370 = *(*float64)(unsafe.Add(mBase, uint32(v344+v49<<(uint(int32(3))%32))))
	if base.F64_lt(v345, v370) == int32(0) {
		v374 = v370
		goto L95
	} else {
		goto L105
	}
L97:
	;
	v352 = *(*float64)(unsafe.Add(mBase, uint32(v344+v51<<(uint(int32(3))%32))))
	if base.F64_lt(v345, v352) != 0 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	if base.F64_gt(v345, float64(0)) == int32(0) {
		v374 = v345
		goto L95
	} else {
		goto L104
	}
L100:
	;
	v354 = int32(0)
	goto L102
L101:
	;
	v354 = v49
	goto L102
L102:
	;
	v358 = *(*float64)(unsafe.Add(mBase, uint32(v344+v354<<(uint(int32(3))%32))))
	if base.F64_gt(v358, float64(0)) == int32(0) {
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
	v374 = v345
	goto L95
L106:
	;
	v378 = v331 + int32(1)
	if v378 != v51 {
		v331 = v378
		goto L93
	} else {
		goto L107
	}
L107:
	;
	goto L94
L108:
	;
	v402 = v321 + v391<<(uint(int32(3))%32)
	v403 = *(*float64)(unsafe.Add(mBase, uint32(v402)))
	if base.B2i32(v49 < int32(0)) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v469 = int32(-1)
	goto L15
L110:
	;
	if base.F64_lt(v432, float64(0)) != 0 {
		goto L16
	} else {
		goto L121
	}
L111:
	;
	v428 = *(*float64)(unsafe.Add(mBase, uint32(v402+v49<<(uint(int32(3))%32))))
	if base.F64_gt(v403, v428) == int32(0) {
		v432 = v428
		goto L110
	} else {
		goto L120
	}
L112:
	;
	v410 = *(*float64)(unsafe.Add(mBase, uint32(v402+v51<<(uint(int32(3))%32))))
	if base.F64_gt(v403, v410) != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L114
L114:
	;
	if base.F64_gt(v403, float64(0)) == int32(0) {
		v432 = v403
		goto L110
	} else {
		goto L119
	}
L115:
	;
	v412 = int32(0)
	goto L117
L116:
	;
	v412 = v49
	goto L117
L117:
	;
	v416 = *(*float64)(unsafe.Add(mBase, uint32(v402+v412<<(uint(int32(3))%32))))
	if base.F64_gt(v416, float64(0)) == int32(0) {
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
	v432 = v403
	goto L110
L121:
	;
	v437 = v391 + int32(1)
	if v51 != v437 {
		v391 = v437
		goto L108
	} else {
		goto L122
	}
L122:
	;
	goto L109
L123:
	;
	v1251 = v686
	goto L4
L124:
	;
	v686 = v668
	goto L123
L125:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v534 = int32(2147483647)
	v535 = v533 & v534
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v538 = v536 & v534
	if base.Ui32(v535) < base.Ui32(v538) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v546 = v535
	goto L129
L127:
	;
	goto L128
L128:
	;
	if base.Ui32(v535) < base.Ui32(v538) {
		goto L137
	} else {
		goto L138
	}
L129:
	;
	v560 = v9 + int32(8) + v546<<(uint(int32(3))%32)
	v561 = *(*float64)(unsafe.Add(mBase, uint32(v560)))
	if base.F64_ne(v561, float64(0)) != 0 {
		v668 = v516
		goto L124
	} else {
		goto L131
	}
L130:
	;
	goto L128
L131:
	;
	if base.B2i32(v536 < int32(0)) == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v569 = *(*float64)(unsafe.Add(mBase, uint32(v560+v538<<(uint(int32(3))%32))))
	if base.F64_ne(v569, float64(0)) != 0 {
		v668 = v516
		goto L124
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v573 = v546 + int32(1)
	if v573 != v538 {
		v546 = v573
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
	v590 = v535
	goto L139
L138:
	;
	v590 = v538
	goto L139
L139:
	;
	if v590 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v686 = int32(1)
	goto L123
L141:
	;
	goto L142
L142:
	;
	v594 = int32(8)
	v601 = int32(0)
	goto L143
L143:
	;
	v613 = int32(0)
	v615 = v601 << (uint(int32(3)) % 32)
	v616 = v22 + v594 + v615
	v617 = *(*float64)(unsafe.Add(mBase, uint32(v616)))
	v619 = base.B2i32(v533 < v613)
	if v533 < v613 {
		v626 = v617
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v668 = v654
	goto L124
L145:
	;
	v627 = v615 + (v9 + v594)
	v628 = *(*float64)(unsafe.Add(mBase, uint32(v627)))
	v630 = base.B2i32(v536 < int32(0))
	if v536 < int32(0) {
		v637 = v628
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v623 = *(*float64)(unsafe.Add(mBase, uint32(v616+v533<<(uint(int32(3))%32))))
	if base.F64_lt(v617, v623) != 0 {
		v626 = v617
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v626 = v623
	goto L145
L148:
	;
	if base.F64_gt(v626, v637) != 0 {
		v668 = v613
		goto L124
	} else {
		goto L151
	}
L149:
	;
	v634 = *(*float64)(unsafe.Add(mBase, uint32(v627+v536<<(uint(int32(3))%32))))
	if base.F64_lt(v628, v634) != 0 {
		v637 = v628
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v637 = v634
	goto L148
L151:
	;
	if v533 < v613 {
		v644 = v617
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if v536 < int32(0) {
		v651 = v628
		goto L155
	} else {
		goto L156
	}
L153:
	;
	v642 = *(*float64)(unsafe.Add(mBase, uint32(v616+v533<<(uint(int32(3))%32))))
	if base.F64_gt(v617, v642) != 0 {
		v644 = v617
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v644 = v642
	goto L152
L155:
	;
	if base.F64_gt(v651, v644) != 0 {
		v668 = v613
		goto L124
	} else {
		goto L158
	}
L156:
	;
	v649 = *(*float64)(unsafe.Add(mBase, uint32(v627+v536<<(uint(int32(3))%32))))
	if base.F64_gt(v628, v649) != 0 {
		v651 = v628
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v651 = v649
	goto L155
L158:
	;
	v654 = int32(1)
	v656 = v601 + v654
	if v656 != v590 {
		v601 = v656
		goto L143
	} else {
		goto L159
	}
L159:
	;
	goto L144
L160:
	;
	v1251 = v857
	goto L4
L161:
	;
	v857 = v839
	goto L160
L162:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v705 = int32(2147483647)
	v706 = v704 & v705
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v709 = v707 & v705
	if base.Ui32(v706) < base.Ui32(v709) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v717 = v706
	goto L166
L164:
	;
	goto L165
L165:
	;
	if base.Ui32(v706) < base.Ui32(v709) {
		goto L174
	} else {
		goto L175
	}
L166:
	;
	v731 = v22 + int32(8) + v717<<(uint(int32(3))%32)
	v732 = *(*float64)(unsafe.Add(mBase, uint32(v731)))
	if base.F64_ne(v732, float64(0)) != 0 {
		v839 = v687
		goto L161
	} else {
		goto L168
	}
L167:
	;
	goto L165
L168:
	;
	if base.B2i32(v707 < int32(0)) == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v740 = *(*float64)(unsafe.Add(mBase, uint32(v731+v709<<(uint(int32(3))%32))))
	if base.F64_ne(v740, float64(0)) != 0 {
		v839 = v687
		goto L161
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v744 = v717 + int32(1)
	if v744 != v709 {
		v717 = v744
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
	v761 = v706
	goto L176
L175:
	;
	v761 = v709
	goto L176
L176:
	;
	if v761 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v857 = int32(1)
	goto L160
L178:
	;
	goto L179
L179:
	;
	v765 = int32(8)
	v772 = int32(0)
	goto L180
L180:
	;
	v784 = int32(0)
	v786 = v772 << (uint(int32(3)) % 32)
	v787 = v9 + v765 + v786
	v788 = *(*float64)(unsafe.Add(mBase, uint32(v787)))
	v790 = base.B2i32(v704 < v784)
	if v704 < v784 {
		v797 = v788
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v839 = v825
	goto L161
L182:
	;
	v798 = v786 + (v22 + v765)
	v799 = *(*float64)(unsafe.Add(mBase, uint32(v798)))
	v801 = base.B2i32(v707 < int32(0))
	if v707 < int32(0) {
		v808 = v799
		goto L185
	} else {
		goto L186
	}
L183:
	;
	v794 = *(*float64)(unsafe.Add(mBase, uint32(v787+v704<<(uint(int32(3))%32))))
	if base.F64_lt(v788, v794) != 0 {
		v797 = v788
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v797 = v794
	goto L182
L185:
	;
	if base.F64_gt(v797, v808) != 0 {
		v839 = v784
		goto L161
	} else {
		goto L188
	}
L186:
	;
	v805 = *(*float64)(unsafe.Add(mBase, uint32(v798+v707<<(uint(int32(3))%32))))
	if base.F64_lt(v799, v805) != 0 {
		v808 = v799
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v808 = v805
	goto L185
L188:
	;
	if v704 < v784 {
		v815 = v788
		goto L189
	} else {
		goto L190
	}
L189:
	;
	if v707 < int32(0) {
		v822 = v799
		goto L192
	} else {
		goto L193
	}
L190:
	;
	v813 = *(*float64)(unsafe.Add(mBase, uint32(v787+v704<<(uint(int32(3))%32))))
	if base.F64_gt(v788, v813) != 0 {
		v815 = v788
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v815 = v813
	goto L189
L192:
	;
	if base.F64_gt(v822, v815) != 0 {
		v839 = v784
		goto L161
	} else {
		goto L195
	}
L193:
	;
	v820 = *(*float64)(unsafe.Add(mBase, uint32(v798+v707<<(uint(int32(3))%32))))
	if base.F64_gt(v799, v820) != 0 {
		v822 = v799
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v822 = v820
	goto L192
L195:
	;
	v825 = int32(1)
	v827 = v772 + v825
	if v827 != v761 {
		v772 = v827
		goto L180
	} else {
		goto L196
	}
L196:
	;
	goto L181
L197:
	;
	v863 = int32(1) << (uint(v859) % 32)
	if v863&int32(_a_F_g_cube_consistent_1) != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	if v863&int32(_a_F_g_cube_consistent_2) != 0 {
		goto L6
	} else {
		goto L199
	}
L199:
	;
	if v859 != int32(3) {
		v1251 = v2
		goto L4
	} else {
		goto L200
	}
L200:
	;
	goto L6
L201:
	;
	v1251 = v1077
	goto L4
L202:
	;
	v1077 = v1055
	goto L201
L203:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v891 = int32(2147483647)
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v896 = base.B2i32(base.Ui32(v890&v891) < base.Ui32(v893&v891))
	if base.Ui32(v890&v891) < base.Ui32(v893&v891) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v897 = v9
	goto L206
L205:
	;
	v897 = v22
	goto L206
L206:
	;
	if base.Ui32(v890&v891) < base.Ui32(v893&v891) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v985 = v969 & int32(2147483647)
	if base.Ui32(v985) <= base.Ui32(v901) {
		goto L231
	} else {
		goto L232
	}
L208:
	;
	v898 = v22
	goto L210
L209:
	;
	v898 = v9
	goto L210
L210:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v898)+4))
	v901 = v899 & int32(2147483647)
	if v901 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v897)+4))
	v969 = v904
	goto L207
L212:
	;
	goto L213
L213:
	;
	v905 = int32(8)
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v897)+4))
	v912 = int32(0)
	goto L214
L214:
	;
	v927 = v912 << (uint(int32(3)) % 32)
	v928 = v897 + v905 + v927
	v929 = *(*float64)(unsafe.Add(mBase, uint32(v928)))
	v931 = base.B2i32(v909 < int32(0))
	if v909 < int32(0) {
		v938 = v929
		goto L216
	} else {
		goto L217
	}
L215:
	;
	v969 = v909
	goto L207
L216:
	;
	v939 = v927 + (v898 + v905)
	v940 = *(*float64)(unsafe.Add(mBase, uint32(v939)))
	v942 = base.B2i32(v899 < int32(0))
	if v899 < int32(0) {
		v949 = v940
		goto L219
	} else {
		goto L220
	}
L217:
	;
	v935 = *(*float64)(unsafe.Add(mBase, uint32(v928+v909<<(uint(int32(3))%32))))
	if base.F64_lt(v929, v935) != 0 {
		v938 = v929
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v938 = v935
	goto L216
L219:
	;
	if base.F64_gt(v938, v949) != 0 {
		v1055 = v875
		goto L202
	} else {
		goto L222
	}
L220:
	;
	v946 = *(*float64)(unsafe.Add(mBase, uint32(v939+v899<<(uint(int32(3))%32))))
	if base.F64_gt(v940, v946) != 0 {
		v949 = v940
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v949 = v946
	goto L219
L222:
	;
	if v909 < int32(0) {
		v956 = v929
		goto L223
	} else {
		goto L224
	}
L223:
	;
	if v899 < int32(0) {
		v963 = v940
		goto L226
	} else {
		goto L227
	}
L224:
	;
	v954 = *(*float64)(unsafe.Add(mBase, uint32(v928+v909<<(uint(int32(3))%32))))
	if base.F64_gt(v929, v954) != 0 {
		v956 = v929
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v956 = v954
	goto L223
L226:
	;
	if base.F64_gt(v963, v956) != 0 {
		v1055 = v875
		goto L202
	} else {
		goto L229
	}
L227:
	;
	v961 = *(*float64)(unsafe.Add(mBase, uint32(v939+v899<<(uint(int32(3))%32))))
	if base.F64_lt(v940, v961) != 0 {
		v963 = v940
		goto L226
	} else {
		goto L228
	}
L228:
	;
	v963 = v961
	goto L226
L229:
	;
	v967 = v912 + int32(1)
	if v967 != v901 {
		v912 = v967
		goto L214
	} else {
		goto L230
	}
L230:
	;
	goto L215
L231:
	;
	v1077 = int32(1)
	goto L201
L232:
	;
	goto L233
L233:
	;
	v997 = v901
	goto L234
L234:
	;
	v1009 = v897 + int32(8) + v997<<(uint(int32(3))%32)
	v1010 = *(*float64)(unsafe.Add(mBase, uint32(v1009)))
	if base.B2i32(v969 < int32(0)) == int32(0) {
		goto L238
	} else {
		goto L239
	}
L235:
	;
	v1055 = int32(0)
	goto L202
L236:
	;
	goto L235
L237:
	;
	if base.F64_lt(v1036, float64(0)) != 0 {
		goto L236
	} else {
		goto L247
	}
L238:
	;
	v1013 = int32(0)
	v1018 = *(*float64)(unsafe.Add(mBase, uint32(v1009+v985<<(uint(int32(3))%32))))
	if base.F64_lt(v1010, v1018) != 0 {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	goto L240
L240:
	;
	if base.F64_gt(v1010, float64(0)) != 0 {
		goto L236
	} else {
		goto L246
	}
L241:
	;
	v1020 = v1013
	goto L243
L242:
	;
	v1020 = v969
	goto L243
L243:
	;
	v1024 = *(*float64)(unsafe.Add(mBase, uint32(v1009+v1020<<(uint(int32(3))%32))))
	if base.F64_gt(v1024, float64(0)) != 0 {
		v1055 = v1013
		goto L202
	} else {
		goto L244
	}
L244:
	;
	v1030 = *(*float64)(unsafe.Add(mBase, uint32(v1009+v969<<(uint(int32(3))%32))))
	if base.F64_gt(v1010, v1030) == int32(0) {
		v1036 = v1030
		goto L237
	} else {
		goto L245
	}
L245:
	;
	v1036 = v1010
	goto L237
L246:
	;
	v1036 = v1010
	goto L237
L247:
	;
	v1040 = int32(1)
	v1042 = v997 + v1040
	if v985 != v1042 {
		v997 = v1042
		goto L234
	} else {
		goto L248
	}
L248:
	;
	v1055 = v1040
	goto L202
L249:
	;
	v1251 = v1248
	goto L4
L250:
	;
	v1248 = v1230
	goto L249
L251:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v1096 = int32(2147483647)
	v1097 = v1095 & v1096
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v1100 = v1098 & v1096
	if base.Ui32(v1097) < base.Ui32(v1100) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1108 = v1097
	goto L255
L253:
	;
	goto L254
L254:
	;
	if base.Ui32(v1097) < base.Ui32(v1100) {
		goto L263
	} else {
		goto L264
	}
L255:
	;
	v1122 = v9 + int32(8) + v1108<<(uint(int32(3))%32)
	v1123 = *(*float64)(unsafe.Add(mBase, uint32(v1122)))
	if base.F64_ne(v1123, float64(0)) != 0 {
		v1230 = v1078
		goto L250
	} else {
		goto L257
	}
L256:
	;
	goto L254
L257:
	;
	if base.B2i32(v1098 < int32(0)) == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1131 = *(*float64)(unsafe.Add(mBase, uint32(v1122+v1100<<(uint(int32(3))%32))))
	if base.F64_ne(v1131, float64(0)) != 0 {
		v1230 = v1078
		goto L250
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v1135 = v1108 + int32(1)
	if v1135 != v1100 {
		v1108 = v1135
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
	v1152 = v1097
	goto L265
L264:
	;
	v1152 = v1100
	goto L265
L265:
	;
	if v1152 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1248 = int32(1)
	goto L249
L267:
	;
	goto L268
L268:
	;
	v1156 = int32(8)
	v1163 = int32(0)
	goto L269
L269:
	;
	v1175 = int32(0)
	v1177 = v1163 << (uint(int32(3)) % 32)
	v1178 = v22 + v1156 + v1177
	v1179 = *(*float64)(unsafe.Add(mBase, uint32(v1178)))
	v1181 = base.B2i32(v1095 < v1175)
	if v1095 < v1175 {
		v1188 = v1179
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1230 = v1216
	goto L250
L271:
	;
	v1189 = v1177 + (v9 + v1156)
	v1190 = *(*float64)(unsafe.Add(mBase, uint32(v1189)))
	v1192 = base.B2i32(v1098 < int32(0))
	if v1098 < int32(0) {
		v1199 = v1190
		goto L274
	} else {
		goto L275
	}
L272:
	;
	v1185 = *(*float64)(unsafe.Add(mBase, uint32(v1178+v1095<<(uint(int32(3))%32))))
	if base.F64_lt(v1179, v1185) != 0 {
		v1188 = v1179
		goto L271
	} else {
		goto L273
	}
L273:
	;
	v1188 = v1185
	goto L271
L274:
	;
	if base.F64_gt(v1188, v1199) != 0 {
		v1230 = v1175
		goto L250
	} else {
		goto L277
	}
L275:
	;
	v1196 = *(*float64)(unsafe.Add(mBase, uint32(v1189+v1098<<(uint(int32(3))%32))))
	if base.F64_lt(v1190, v1196) != 0 {
		v1199 = v1190
		goto L274
	} else {
		goto L276
	}
L276:
	;
	v1199 = v1196
	goto L274
L277:
	;
	if v1095 < v1175 {
		v1206 = v1179
		goto L278
	} else {
		goto L279
	}
L278:
	;
	if v1098 < int32(0) {
		v1213 = v1190
		goto L281
	} else {
		goto L282
	}
L279:
	;
	v1204 = *(*float64)(unsafe.Add(mBase, uint32(v1178+v1095<<(uint(int32(3))%32))))
	if base.F64_gt(v1179, v1204) != 0 {
		v1206 = v1179
		goto L278
	} else {
		goto L280
	}
L280:
	;
	v1206 = v1204
	goto L278
L281:
	;
	if base.F64_gt(v1213, v1206) != 0 {
		v1230 = v1175
		goto L250
	} else {
		goto L284
	}
L282:
	;
	v1211 = *(*float64)(unsafe.Add(mBase, uint32(v1189+v1098<<(uint(int32(3))%32))))
	if base.F64_gt(v1190, v1211) != 0 {
		v1213 = v1190
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1213 = v1211
	goto L281
L284:
	;
	v1216 = int32(1)
	v1218 = v1163 + v1216
	if v1218 != v1152 {
		v1163 = v1218
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
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L1
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	return v1251
L289:
	;
	goto L288
}
