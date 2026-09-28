package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_read_stream_reset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+14)) = v2
	goto L1
L1:
	;
	v20 = F_read_stream_next_buffer(m, l0, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+88)))
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v25 <= v24 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	return
L4:
	;
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_ReleaseBuffer(m, v20)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	goto L2
L8:
	;
	goto L1
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+22)) = int32(_a_F_read_stream_reset_0)
	v79 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v79)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v79
	return
L10:
	;
	v28 = l0 + int32(92)
	v30 = v24
	goto L11
L11:
	;
	v36 = v30 << (uint(int32(2)) % 32)
	v37 = v28 + v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v38 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v43 = v41 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v43)
	F_ReleaseBuffer(m, v38)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(0)
	v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
	if v30 < v49-int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v28+v53<<(uint(int32(2))%32)+v36))) = int32(0)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v61 = v30 + int32(1)
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v63 != v61&int32(_a_F_read_stream_reset_1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v67 = v61
	goto L20
L19:
	;
	v67 = int32(0)
	goto L20
L20:
	;
	v68 = base.I32_extend16_s(v67)
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v68 < v69 {
		v30 = v68
		goto L11
	} else {
		goto L21
	}
L21:
	;
	goto L12
}
func F_read_stream_start_pending_read(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
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
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int64
	_ = v381
	var v383 int64
	_ = v383
	var v414 int64
	_ = v414
	var v425 int64
	_ = v425
	var v463 int32
	_ = v463
	var v464 int64
	_ = v464
	var v467 int64
	_ = v467
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v524 int64
	_ = v524
	var v526 int64
	_ = v526
	var v557 int64
	_ = v557
	var v560 int32
	_ = v560
	var v562 int64
	_ = v562
	var v564 int64
	_ = v564
	var v567 int64
	_ = v567
	var v568 int64
	_ = v568
	var v573 int64
	_ = v573
	var v576 int64
	_ = v576
	var v581 int64
	_ = v581
	var v609 int64
	_ = v609
	var v616 int64
	_ = v616
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v660 int64
	_ = v660
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v725 int64
	_ = v725
	var v727 int32
	_ = v727
	var v728 int64
	_ = v728
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v744 int64
	_ = v744
	var v748 int32
	_ = v748
	var v760 int32
	_ = v760
	var v764 int64
	_ = v764
	var v768 int64
	_ = v768
	var v770 int32
	_ = v770
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int64
	_ = v802
	var v808 int32
	_ = v808
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int64
	_ = v845
	var v851 int32
	_ = v851
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1015 int32
	_ = v1015
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1142 int64
	_ = v1142
	var v1146 int64
	_ = v1146
	var v1150 int64
	_ = v1150
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1269 int32
	_ = v1269
	var v1275 int32
	_ = v1275
	v30 = m.G0
	v32 = v30 - int32(16)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v35 != int32(1) {
		v55 = v34
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	if v56 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v38 == v39 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v43 == int32(-1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v38
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	if int32(0) < v50 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v46 = v34
	goto L8
L7:
	;
	v46 = v34 | int32(2)
	goto L8
L8:
	;
	v55 = v46
	goto L1
L9:
	;
	v53 = v34 | int32(2)
	goto L11
L10:
	;
	v53 = v34
	goto L11
L11:
	;
	v55 = v53
	goto L1
L12:
	;
	v114 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+10)))
	v115 = v113 + v114
	if v115 != 0 {
		goto L35
	} else {
		goto L36
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[0]))
	v65 = base.I32_div_s(v63, int32(4))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[1]))
	v68 = v65 - v67
	if base.Ui32(v68) <= base.Ui32(v65) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[2]))
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[3]))
	v95 = v90 - v92 - int32(8)
	if base.Ui32(v95) <= base.Ui32(v90) {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	if base.Ui32(int32(_a_F_read_stream_start_pending_read_0)) < base.Ui32(v71) {
		v113 = int32(_a_F_read_stream_start_pending_read_1)
		goto L12
	} else {
		goto L20
	}
L17:
	;
	v71 = v68
	goto L19
L18:
	;
	v71 = int32(0)
	goto L19
L19:
	;
	goto L16
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[0]))
	v79 = base.I32_div_s(v77, int32(4))
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[1]))
	v82 = v79 - v81
	if base.Ui32(v82) <= base.Ui32(v79) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v113 = v85
	goto L12
L22:
	;
	v85 = v82
	goto L24
L23:
	;
	v85 = int32(0)
	goto L24
L24:
	;
	goto L21
L25:
	;
	if base.Ui32(int32(_a_F_read_stream_start_pending_read_0)) < base.Ui32(v98) {
		v113 = int32(_a_F_read_stream_start_pending_read_1)
		goto L12
	} else {
		goto L29
	}
L26:
	;
	v98 = v95
	goto L28
L27:
	;
	v98 = int32(0)
	goto L28
L28:
	;
	goto L25
L29:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[2]))
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[3]))
	v109 = v104 - v106 - int32(8)
	if base.Ui32(v109) <= base.Ui32(v104) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v113 = v112
	goto L12
L31:
	;
	v112 = v109
	goto L33
L32:
	;
	v112 = int32(0)
	goto L33
L33:
	;
	goto L30
L34:
	;
	v124 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+64)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v124
	if v123 < v124 {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	v116 = int32(_a_F_read_stream_start_pending_read_1)
	if v116 <= v115 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v123 = base.B2i32(v120 == int32(0))
	goto L34
L38:
	;
	v119 = v116
	goto L40
L39:
	;
	v119 = v115
	goto L40
L40:
	;
	v123 = v119
	goto L34
L41:
	;
	m.G0 = v1275 + int32(16)
	return v1269
L42:
	;
	v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	v129 = base.I32_extend16_s(v127 + v123)
	v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
	if v129 < v130 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v138 = v124
	goto L44
L44:
	;
	v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+82)))
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+88)))
	v141 = v138 + v140
	v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v142 < v141 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v129)
	goto L47
L46:
	;
	goto L47
L47:
	;
	v133 = int32(0)
	if v133 < v127 {
		v1269 = v133
		v1275 = v32
		goto L41
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v123
	v138 = v123
	goto L44
L49:
	;
	v147 = v142
	goto L52
L50:
	;
	goto L51
L51:
	;
	v215 = v139 * int32(84)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v219 = v215 + v216 + int32(4)
	v221 = l0 + int32(92)
	v224 = v221 + v140<<(uint(int32(2))%32)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v226 = m.G0
	v228 = v226 + int32(-64)
	m.G0 = v228
	v231 = v32 + int32(12)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	if v233 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L52:
	;
	v176 = v147 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v176)
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(92)+v147<<(uint(int32(2))%32)))) = int32(0)
	v183 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v183 < v141 {
		v147 = v183
		goto L52
	} else {
		goto L54
	}
L53:
	;
	goto L51
L54:
	;
	goto L53
L55:
	;
	m.G0 = v1066 - int32(-64)
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+12))
	v1096 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1063)+12)))
	v1097 = v1095 + v1096
	*(*uint16)(unsafe.Add(mBase, uint32(v1063)+12)) = uint16(v1097)
	if v1067 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L64
	} else {
		goto L184
	}
L57:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+8)))
	if v245 != int32(116) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v233)+48))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+118)))
	if v237 != int32(116) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+24)))
	if v240 == int32(0) {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	goto L57
L61:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v250 = F_IOContextForStrategy(m, v249)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v254 = int32(3)
	v255 = int32(1)
	goto L63
L63:
	;
	if v232 <= int32(0) {
		v985 = l0
		v986 = v232
		v988 = v228
		v990 = v219
		v993 = v138
		v995 = v32
		v996 = v140
		v997 = v225
		v999 = v55
		v1000 = v231
		v1003 = v221
		v1005 = v224
		v1007 = v215
		goto L67
	} else {
		goto L68
	}
L64:
	;
	return int32(0)
L65:
	;
	v254 = v250
	v255 = int32(0)
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274))) = int32(1)
	v1063 = v259
	v1066 = v262
	v1067 = int32(0)
	v1071 = v267
	v1073 = v269
	v1074 = v270
	v1081 = v277
	v1085 = v281
	goto L55
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1000))) = v986
	v1015 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v990)+32)) = uint16(v1015)
	*(*uint16)(unsafe.Add(mBase, uint32(v990)+30)) = uint16(v986)
	*(*uint16)(unsafe.Add(mBase, uint32(v990)+28)) = uint16(v999)
	*(*int32)(unsafe.Add(mBase, uint32(v990)+24)) = v997
	*(*int32)(unsafe.Add(mBase, uint32(v990)+20)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v990+int32(36)))) = int32(-1)
	goto L177
L68:
	;
	v259 = l0
	v260 = v232
	v262 = v228
	v263 = int32(0)
	v264 = v219
	v267 = v138
	v269 = v32
	v270 = v140
	v271 = v225
	v273 = v55
	v274 = v231
	v275 = v254
	v277 = v221
	v279 = v224
	v280 = v255
	v281 = v215
	goto L69
L69:
	;
	v290 = v279 + v263<<(uint(int32(2))%32)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	if v291 != 0 {
		goto L79
	} else {
		goto L80
	}
L70:
	;
	if v263 == int32(0) {
		goto L66
	} else {
		goto L176
	}
L71:
	;
	goto L70
L72:
	;
	if base.B2i32(v260 < int32(2))|v263 != 0 {
		v949 = v260
		goto L166
	} else {
		goto L167
	}
L73:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v851)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v290))) = v878 + int32(1)
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+19)))
	if v882 != 0 {
		goto L71
	} else {
		goto L165
	}
L74:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v304)+272))
	if v835 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L75:
	;
	v760 = v280*int32(320) + v275<<(uint(int32(6))%32)
	v764 = *(*int64)(unsafe.Add(mBase, uint32(v760)+uint32(_c_F_read_stream_start_pending_read[4])))
	*(*int64)(unsafe.Add(mBase, uint32(v760)+uint32(_c_F_read_stream_start_pending_read[4]))) = v764 + int64(1)
	v768 = *(*int64)(unsafe.Add(mBase, uint32(v760)+uint32(_c_F_read_stream_start_pending_read[5])))
	*(*int64)(unsafe.Add(mBase, uint32(v760)+uint32(_c_F_read_stream_start_pending_read[5]))) = v768
	v770 = int32(1)
	F_pgstat_count_backend_io_op(m, v280, v275, int32(2), v770, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[6])) = uint8(v770)
	*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[7])) = uint8(v770)
	goto L150
L76:
	;
	v742 = int32(_a_F_read_stream_start_pending_read_2)
	v744 = *(*int64)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[8])) = v744 + int64(1)
	v748 = v737
	goto L75
L77:
	;
	v725 = int64(0)
	v727 = int32(24)
	v728 = base.AtomicRmwCmpxchg64(m, v724, v727, v725, v725)
	v733 = int32(base.Ui32(base.I32_wrap_i64(v728))>>(uint(v727)%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v262)+19)) = uint8(v733)
	if v733 == int32(0) {
		goto L72
	} else {
		goto L149
	}
L78:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[9]))
	v719 = int32(56)
	v724 = v718 + v291*v719 - v719
	goto L77
L79:
	;
	if int32(0) <= v291 {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v301 = v263 + v271
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+8)))
	if v305 != int32(116) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[10]))
	v724 = v295 + (v291^int32(-1))*int32(56)
	goto L77
L83:
	;
	if v304 == int32(0) {
		v851 = v688
		goto L73
	} else {
		goto L148
	}
L84:
	;
	v665 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[11]))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	F_ResourceOwnerForget(m, v665, base.I64_extend_i32_s(v666+int32(1)), int32(_a_F_read_stream_start_pending_read_3))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L64
	} else {
		goto L143
	}
L85:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v264)+16))
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[11]))
	F_ResourceOwnerEnlarge(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L64
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v653 = F_LocalBufferAlloc(m, v303, v302, v301, v262+int32(19))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L64
	} else {
		goto L141
	}
L88:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L64
	} else {
		goto L89
	}
L89:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	*(*int32)(unsafe.Add(mBase, uint32(v262)+20)) = v315
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v262)+24)) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v303)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v262)+36)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v262)+32)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v262)+28)) = v319
	v324 = v262 + int32(20)
	v325 = F_BufTableHashCode(m, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L64
	} else {
		goto L90
	}
L90:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[12]))
	v335 = v328 + v325&int32(127)<<(uint(int32(7))%32) + int32(_a_F_read_stream_start_pending_read_4)
	v337 = F_LWLockAcquire(m, v335, int32(1))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L64
	} else {
		goto L91
	}
L91:
	;
	v339 = F_BufTableLookup(m, v324, v325)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L64
	} else {
		goto L92
	}
L92:
	;
	if int32(0) <= v339 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[9]))
	v347 = v344 + v339*int32(56)
	v349 = F_PinBuffer(m, v347, v308, int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L64
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	F_LWLockRelease(m, v335)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L64
	} else {
		goto L99
	}
L96:
	;
	F_LWLockRelease(m, v335)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L64
	} else {
		goto L97
	}
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v262)+19)) = uint8(v349)
	if v349 == int32(0) {
		v688 = v347
		goto L83
	} else {
		goto L98
	}
L98:
	;
	v737 = v347
	goto L76
L99:
	;
	v358 = F_GetVictimBuffer(m, v308, v275)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L64
	} else {
		goto L100
	}
L100:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[9]))
	v363 = F_LWLockAcquire(m, v335, int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L64
	} else {
		goto L101
	}
L101:
	;
	v365 = int32(56)
	v367 = v361 + v358*v365
	v369 = v367 - v365
	v373 = v367 - int32(36)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	v375 = F_BufTableInsert(m, v262+int32(20), v325, v374)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L64
	} else {
		goto L102
	}
L102:
	;
	if int32(0) <= v375 {
		goto L84
	} else {
		goto L103
	}
L103:
	;
	v380 = v367 - int32(32)
	v381 = int64(4194304)
	v383 = base.AtomicRmwOr64(m, v380, int32(0), v381)
	if v383&v381 != int64(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v414 = v383
	goto L107
L105:
	;
	v557 = v383
	goto L106
L106:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v262)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+16)) = v560
	v562 = *(*int64)(unsafe.Add(mBase, uint32(v262)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v369)+8)) = v562
	v564 = *(*int64)(unsafe.Add(mBase, uint32(v262)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v369))) = v564
	v567 = v557 | int64(4194304)
	v568 = int64(2181300224)
	if v302 == int32(3) {
		goto L128
	} else {
		goto L129
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+60)) = int32(_a_F_read_stream_start_pending_read_5)
	*(*int32)(unsafe.Add(mBase, uint32(v262)+56)) = int32(_a_F_read_stream_start_pending_read_6)
	*(*int32)(unsafe.Add(mBase, uint32(v262)+52)) = int32(_a_F_read_stream_start_pending_read_7)
	*(*int32)(unsafe.Add(mBase, uint32(v262)+48)) = int32(0)
	v425 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v262)+40)) = v425
	if v414&int64(4194304) != v425 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v557 = v526
	goto L106
L109:
	;
	goto L112
L110:
	;
	goto L111
L111:
	;
	v504 = int32(_a_F_read_stream_start_pending_read_8)
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[13]))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v262+int32(40))+8))
	if v507 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L112:
	;
	F_perform_spin_delay(m, v262+int32(40))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L64
	} else {
		goto L114
	}
L113:
	;
	goto L111
L114:
	;
	v464 = int64(0)
	v467 = base.AtomicRmwCmpxchg64(m, v380, int32(0), v464, v464)
	if v467&int64(4194304) != v464 {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v524 = int64(4194304)
	v526 = base.AtomicRmwOr64(m, v380, int32(0), v524)
	if v526&v524 != int64(0) {
		v414 = v526
		goto L107
	} else {
		goto L127
	}
L117:
	;
	goto L116
L118:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[13])) = v522
	goto L117
L119:
	;
	if int32(999) < v505 {
		goto L117
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	if v505 < int32(11) {
		goto L117
	} else {
		goto L126
	}
L122:
	;
	v512 = int32(900)
	if v512 <= v505 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v515 = v512
	goto L125
L124:
	;
	v515 = v505
	goto L125
L125:
	;
	v522 = v515 + int32(100)
	goto L118
L126:
	;
	v522 = v505 - int32(1)
	goto L118
L127:
	;
	goto L108
L128:
	;
	v573 = v568
	goto L130
L129:
	;
	v573 = int64(33816576)
	goto L130
L130:
	;
	if v305 == int32(112) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v576 = v568
	goto L133
L132:
	;
	v576 = v573
	goto L133
L133:
	;
	v581 = base.AtomicRmwCmpxchg64(m, v380, int32(0), v567, v576|v557&int64(-38010881))
	if v581 != v567 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v609 = v581
	goto L137
L135:
	;
	goto L136
L136:
	;
	F_LWLockRelease(m, v335)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L64
	} else {
		goto L140
	}
L137:
	;
	v616 = base.AtomicRmwCmpxchg64(m, v380, int32(0), v609, v609&int64(-38010881)|v576)
	if v609 != v616 {
		v609 = v616
		goto L137
	} else {
		goto L139
	}
L138:
	;
	goto L136
L139:
	;
	goto L138
L140:
	;
	v649 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v262)+19)) = uint8(v649)
	v688 = v369
	goto L83
L141:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+19)))
	if v655 != int32(1) {
		v688 = v653
		goto L83
	} else {
		goto L142
	}
L142:
	;
	v658 = int32(_a_F_read_stream_start_pending_read_9)
	v660 = *(*int64)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[14]))
	*(*int64)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[14])) = v660 + int64(1)
	v748 = v653
	goto L75
L143:
	;
	F_UnpinBufferNoOwner(m, v369)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L64
	} else {
		goto L144
	}
L144:
	;
	v676 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[9]))
	v679 = v676 + v375*int32(56)
	v681 = F_PinBuffer(m, v679, v308, int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L64
	} else {
		goto L145
	}
L145:
	;
	F_LWLockRelease(m, v335)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L64
	} else {
		goto L146
	}
L146:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v262)+19)) = uint8(v681)
	if v681 != 0 {
		v737 = v679
		goto L76
	} else {
		goto L147
	}
L147:
	;
	v688 = v679
	goto L83
L148:
	;
	v808 = v688
	goto L74
L149:
	;
	goto L71
L150:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[15])))
	if v780 == int32(1) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v783 = int32(_a_F_read_stream_start_pending_read_10)
	v785 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[16]))
	v787 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[17]))
	*(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[16])) = v785 + v787
	goto L153
L152:
	;
	goto L153
L153:
	;
	if v304 == int32(0) {
		v851 = v748
		goto L73
	} else {
		goto L154
	}
L154:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v304)+272))
	if v792 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+268)))
	if v795 != int32(1) {
		v808 = v748
		goto L74
	} else {
		goto L158
	}
L156:
	;
	v801 = v792
	goto L157
L157:
	;
	v802 = *(*int64)(unsafe.Add(mBase, uint32(v801)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v801)+120)) = v802 + int64(1)
	v808 = v748
	goto L74
L158:
	;
	F_pgstat_assoc_relation(m, v304)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L64
	} else {
		goto L159
	}
L159:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v304)+272))
	v801 = v800
	goto L157
L160:
	;
	v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+268)))
	if v838 != int32(1) {
		v851 = v808
		goto L73
	} else {
		goto L163
	}
L161:
	;
	v844 = v835
	goto L162
L162:
	;
	v845 = *(*int64)(unsafe.Add(mBase, uint32(v844)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v844)+112)) = v845 + int64(1)
	v851 = v808
	goto L73
L163:
	;
	F_pgstat_assoc_relation(m, v304)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L64
	} else {
		goto L164
	}
L164:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v304)+272))
	v844 = v843
	goto L162
L165:
	;
	goto L72
L166:
	;
	v952 = v263 + int32(1)
	if v952 < v949 {
		v260 = v949
		v263 = v952
		goto L69
	} else {
		goto L175
	}
L167:
	;
	v917 = int32(_a_F_read_stream_start_pending_read_11)
	v919 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[18]))
	v920 = int32(1)
	v921 = v919 + v920
	*(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[18])) = v921
	*(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[18])) = v921 - v920
	v932 = int32(_a_F_read_stream_start_pending_read_12) - v271&int32(_a_F_read_stream_start_pending_read_13)
	if v260 <= v932 {
		v949 = v260
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v936 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L64
	} else {
		goto L169
	}
L169:
	;
	if v936 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+8)) = v932
	*(*int32)(unsafe.Add(mBase, uint32(v262)+4)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v271
	F_errmsg_internal(m, int32(_a_F_read_stream_start_pending_read_14), v262)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L64
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v949 = v932
	goto L166
L173:
	;
	F_errfinish(m, int32(_a_F_read_stream_start_pending_read_7), int32(1511), int32(_a_F_read_stream_start_pending_read_15))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L64
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v985 = v259
	v986 = v949
	v988 = v262
	v990 = v264
	v993 = v267
	v995 = v269
	v996 = v270
	v997 = v271
	v999 = v273
	v1000 = v274
	v1003 = v277
	v1005 = v279
	v1007 = v281
	goto L67
L176:
	;
	v985 = v259
	v986 = v263
	v988 = v262
	v990 = v264
	v993 = v267
	v995 = v269
	v996 = v270
	v997 = v271
	v999 = v273
	v1000 = v274
	v1003 = v277
	v1005 = v279
	v1007 = v281
	goto L67
L177:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[19]))
	if v1026 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v1027 = F_AsyncReadBuffers(m, v990, v1000)
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L64
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	v1031 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v990)+28)))
	v1033 = v1031 | int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v990)+28)) = uint16(v1033)
	v1035 = int32(1)
	if v999&int32(2) == int32(0) {
		v1063 = v985
		v1066 = v988
		v1067 = v1035
		v1071 = v993
		v1073 = v995
		v1074 = v996
		v1081 = v1003
		v1085 = v1007
		goto L55
	} else {
		goto L182
	}
L181:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1000)))
	*(*uint16)(unsafe.Add(mBase, uint32(v990)+30)) = uint16(v1029)
	v1063 = v985
	v1066 = v988
	v1067 = v1027
	v1071 = v993
	v1073 = v995
	v1074 = v996
	v1081 = v1003
	v1085 = v1007
	goto L55
L182:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v990)+4))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v990)+12))
	v1042 = F_smgrprefetch(m, v1040, v1041, v997, v986)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L64
	} else {
		goto L183
	}
L183:
	;
	v1063 = v985
	v1066 = v988
	v1067 = v1035
	v1071 = v993
	v1073 = v995
	v1074 = v996
	v1081 = v1003
	v1085 = v1007
	goto L55
L184:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L64
	} else {
		goto L185
	}
L185:
	;
	F_errmsg(m, int32(_a_F_read_stream_start_pending_read_16), int32(0))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L64
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_read_stream_start_pending_read_7), int32(1392), int32(_a_F_read_stream_start_pending_read_15))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L64
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	v1158 = int32(0)
	if v1071 <= v1156 {
		v1206 = v1158
		goto L204
	} else {
		goto L205
	}
L189:
	;
	v1101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1063)+4)))
	if v1101 != 0 {
		v1156 = v1095
		goto L188
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+76))
	*(*uint16)(unsafe.Add(mBase, uint32(v1118+v1085))) = uint16(v1074)
	v1121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1063)+4)))
	v1122 = int32(1)
	v1123 = v1121 + v1122
	*(*uint16)(unsafe.Add(mBase, uint32(v1063)+4)) = uint16(v1123)
	v1125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1063)+82)))
	v1127 = v1125 + v1122
	v1129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1063))))
	if v1129 != v1127&int32(_a_F_read_stream_start_pending_read_17) {
		goto L200
	} else {
		goto L201
	}
L192:
	;
	v1102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1063)+18)))
	if v1102 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1104 = v1102 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1063)+18)) = uint16(v1104)
	v1156 = v1095
	goto L188
L194:
	;
	goto L195
L195:
	;
	v1106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1063)+16)))
	if int32(2) <= v1106 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1110 = v1106 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1063)+16)) = uint16(v1110)
	goto L198
L197:
	;
	goto L198
L198:
	;
	v1112 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1063)+14)))
	if v1112 < int32(2) {
		v1156 = v1095
		goto L188
	} else {
		goto L199
	}
L199:
	;
	v1116 = v1112 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1063)+14)) = uint16(v1116)
	v1156 = v1095
	goto L188
L200:
	;
	v1133 = v1127
	goto L202
L201:
	;
	v1133 = int32(0)
	goto L202
L202:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1063)+82)) = uint16(v1133)
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+12))
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+52)) = v1135 + v1136
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+36))
	if v1139 == int32(0) {
		v1156 = v1135
		goto L188
	} else {
		goto L203
	}
L203:
	;
	v1142 = *(*int64)(unsafe.Add(mBase, uint32(v1139)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1139)+32)) = v1142 + int64(1)
	v1146 = *(*int64)(unsafe.Add(mBase, uint32(v1139)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1139)+40)) = v1146 + base.I64_extend_i32_s(v1135)
	v1150 = *(*int64)(unsafe.Add(mBase, uint32(v1139)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1139)+48)) = v1150 + base.I64_extend16_s(base.I64_extend_i32_u(v1123))
	v1156 = v1135
	goto L188
L204:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1063)+10)) = uint16(v1206)
	v1238 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1063)+6)))
	v1240 = base.I32_extend16_s(v1074 + (v1206 + v1156) - v1238)
	if v1240 <= int32(0) {
		goto L210
	} else {
		goto L211
	}
L205:
	;
	v1160 = int32(2)
	v1166 = v1071 - v1156
	v1168 = v1158
	goto L206
L206:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1081+v1156<<(uint(v1160)%32)+v1074<<(uint(v1160)%32)+v1168<<(uint(int32(2))%32))))
	if v1199 == int32(0) {
		v1206 = v1168
		goto L204
	} else {
		goto L208
	}
L207:
	;
	v1206 = v1166
	goto L204
L208:
	;
	v1203 = v1168 + int32(1)
	if v1203 != v1166 {
		v1168 = v1203
		goto L206
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+60)) = v1252 + v1156
	v1255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1063)+64)))
	v1256 = v1255 - v1156
	*(*uint16)(unsafe.Add(mBase, uint32(v1063)+64)) = uint16(v1256)
	v1258 = v1156 + v1074
	if v1238 <= base.I32_extend16_s(v1258) {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	v1244 = v1240 << (uint(int32(2)) % 32)
	if v1244 == int32(0) {
		goto L210
	} else {
		goto L212
	}
L212:
	;
	base.MemoryCopy(m, v1081, v1081+v1238<<(uint(int32(2))%32), v1244)
	goto L210
L213:
	;
	v1262 = v1238
	goto L215
L214:
	;
	v1262 = int32(0)
	goto L215
L215:
	;
	v1263 = v1258 - v1262
	*(*uint16)(unsafe.Add(mBase, uint32(v1063)+88)) = uint16(v1263)
	v1269 = int32(1)
	v1275 = v1073
	goto L41
}
func F_stream_cleanup_files(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v5 = m.G0
	v7 = v5 - int32(1056)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
	v12 = v7 + int32(32)
	v17 = F_pg_snprintf(m, v12, int32(1024), int32(_a_F_stream_cleanup_files_0), v7+int32(16))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_stream_cleanup_files[0]))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
		F_BufFileDeleteFileSet(m, v21, v12, int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			v29 = F_pg_snprintf(m, v12, int32(1024), int32(_a_F_stream_cleanup_files_1), v7)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_stream_cleanup_files[0]))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
				F_BufFileDeleteFileSet(m, v33, v12, int32(1))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					m.G0 = v7 + int32(1056)
					return
				}
			}
		}
	}
}
func F_stream_start_internal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	v6 = m.G0
	v8 = v6 - int32(1056)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[0]))
	if v11 < int32(0) {
		v15 = F_GetCurrentTimestamp(m)
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, _c_F_stream_start_internal[1])) = v15
	} else {
	}
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[2]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if base.B2i32(v19 == int32(2)) == int32(0) {
		F_StartTransactionCommand(m)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			F_maybe_reread_subscription(m)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v28 = F_GetTransactionSnapshot(m)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_PushActiveSnapshot(m, v28)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v34
						v37 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
						if v38 != 0 {
							v56 = v37
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v57
							*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l0
							v61 = v8 + int32(32)
							v66 = F_pg_snprintf(m, v61, int32(1024), int32(_a_F_stream_start_internal_0), v8+int32(16))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								v70 = F_errstart(m, int32(14), int32(0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									if v70 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
										F_errmsg_internal(m, int32(_a_F_stream_start_internal_1), v8)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_stream_start_internal_2), int32(_a_F_stream_start_internal_3), int32(_a_F_stream_start_internal_4))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												v81 = int32(_a_F_stream_start_internal_5)
												v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
												v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
												v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
												if l1 != 0 {
													v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
														F_PopActiveSnapshot(m)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												} else {
													v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
														v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
															v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
															F_subxact_info_read(m, v115, l0)
															mBase = m.M
															v117 = m.ExcPending
															if v117 != 0 {
																return
															} else {
																F_PopActiveSnapshot(m)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v123 = m.ExcPending
																	if v123 != 0 {
																		return
																	} else {
																		m.G0 = v8 + int32(1056)
																		return
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v81 = int32(_a_F_stream_start_internal_5)
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
										v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
										*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
										v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
										if l1 != 0 {
											v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
												F_PopActiveSnapshot(m)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														m.G0 = v8 + int32(1056)
														return
													}
												}
											}
										} else {
											v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
												v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
													v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
													F_subxact_info_read(m, v115, l0)
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[8]))
							*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v41
							v44 = F_palloc(m, int32(44))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
								*(*int32)(unsafe.Add(mBase, uint32(v47)+60)) = v44
								F_FileSetInit(m, v44)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v34
									v54 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
									v56 = v54
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v57
									*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l0
									v61 = v8 + int32(32)
									v66 = F_pg_snprintf(m, v61, int32(1024), int32(_a_F_stream_start_internal_0), v8+int32(16))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										v70 = F_errstart(m, int32(14), int32(0))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											if v70 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
												F_errmsg_internal(m, int32(_a_F_stream_start_internal_1), v8)
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_stream_start_internal_2), int32(_a_F_stream_start_internal_3), int32(_a_F_stream_start_internal_4))
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return
													} else {
														v81 = int32(_a_F_stream_start_internal_5)
														v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
														v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
														v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
														if l1 != 0 {
															v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
																*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
																F_PopActiveSnapshot(m)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v123 = m.ExcPending
																	if v123 != 0 {
																		return
																	} else {
																		m.G0 = v8 + int32(1056)
																		return
																	}
																}
															}
														} else {
															v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
																v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
																mBase = m.M
																v110 = m.ExcPending
																if v110 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
																	v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
																	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
																	F_subxact_info_read(m, v115, l0)
																	mBase = m.M
																	v117 = m.ExcPending
																	if v117 != 0 {
																		return
																	} else {
																		F_PopActiveSnapshot(m)
																		mBase = m.M
																		v121 = m.ExcPending
																		if v121 != 0 {
																			return
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v123 = m.ExcPending
																			if v123 != 0 {
																				return
																			} else {
																				m.G0 = v8 + int32(1056)
																				return
																			}
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v81 = int32(_a_F_stream_start_internal_5)
												v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
												v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
												v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
												if l1 != 0 {
													v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
														F_PopActiveSnapshot(m)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												} else {
													v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
														v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
															v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
															F_subxact_info_read(m, v115, l0)
															mBase = m.M
															v117 = m.ExcPending
															if v117 != 0 {
																return
															} else {
																F_PopActiveSnapshot(m)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v123 = m.ExcPending
																	if v123 != 0 {
																		return
																	} else {
																		m.G0 = v8 + int32(1056)
																		return
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v28 = F_GetTransactionSnapshot(m)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			F_PushActiveSnapshot(m, v28)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[3]))
				*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v34
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
				if v38 != 0 {
					v56 = v37
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v57
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l0
					v61 = v8 + int32(32)
					v66 = F_pg_snprintf(m, v61, int32(1024), int32(_a_F_stream_start_internal_0), v8+int32(16))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						v70 = F_errstart(m, int32(14), int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							if v70 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
								F_errmsg_internal(m, int32(_a_F_stream_start_internal_1), v8)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_stream_start_internal_2), int32(_a_F_stream_start_internal_3), int32(_a_F_stream_start_internal_4))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										v81 = int32(_a_F_stream_start_internal_5)
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
										v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
										*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
										v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
										if l1 != 0 {
											v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
												F_PopActiveSnapshot(m)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														m.G0 = v8 + int32(1056)
														return
													}
												}
											}
										} else {
											v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
												v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
													v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
													F_subxact_info_read(m, v115, l0)
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v81 = int32(_a_F_stream_start_internal_5)
								v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
								v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
								*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
								v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
								if l1 != 0 {
									v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
										*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
										F_PopActiveSnapshot(m)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												m.G0 = v8 + int32(1056)
												return
											}
										}
									}
								} else {
									v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
										v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
											v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
											F_subxact_info_read(m, v115, l0)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return
											} else {
												F_PopActiveSnapshot(m)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														m.G0 = v8 + int32(1056)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[8]))
					*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v41
					v44 = F_palloc(m, int32(44))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
						*(*int32)(unsafe.Add(mBase, uint32(v47)+60)) = v44
						F_FileSetInit(m, v44)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v34
							v54 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
							v56 = v54
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v57
							*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l0
							v61 = v8 + int32(32)
							v66 = F_pg_snprintf(m, v61, int32(1024), int32(_a_F_stream_start_internal_0), v8+int32(16))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								v70 = F_errstart(m, int32(14), int32(0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									if v70 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
										F_errmsg_internal(m, int32(_a_F_stream_start_internal_1), v8)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_stream_start_internal_2), int32(_a_F_stream_start_internal_3), int32(_a_F_stream_start_internal_4))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												v81 = int32(_a_F_stream_start_internal_5)
												v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
												v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
												v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
												if l1 != 0 {
													v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
														F_PopActiveSnapshot(m)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												} else {
													v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
														v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
															v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
															F_subxact_info_read(m, v115, l0)
															mBase = m.M
															v117 = m.ExcPending
															if v117 != 0 {
																return
															} else {
																F_PopActiveSnapshot(m)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v123 = m.ExcPending
																	if v123 != 0 {
																		return
																	} else {
																		m.G0 = v8 + int32(1056)
																		return
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v81 = int32(_a_F_stream_start_internal_5)
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
										v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
										*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
										v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
										if l1 != 0 {
											v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
												F_PopActiveSnapshot(m)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														m.G0 = v8 + int32(1056)
														return
													}
												}
											}
										} else {
											v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
												v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
													v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
													F_subxact_info_read(m, v115, l0)
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_stream_stop_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(_a_F_stream_stop_cb_wrapper_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_stream_stop_cb_wrapper[0]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, _c_F_stream_stop_cb_wrapper[0])) = v8 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_stream_stop_cb_wrapper_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(1058)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v8 + int32(16)
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+147)) = uint8(v27)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+164)) = uint8(v4)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+152)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v29
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if v34 == v4 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_stream_stop_cb_wrapper_2)
				F_errmsg(m, int32(_a_F_stream_stop_cb_wrapper_3), v8)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_stream_stop_cb_wrapper_4), int32(1438), int32(_a_F_stream_stop_cb_wrapper_5))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.T0[v34].(func(*base.Module, int32, int32))(m, v12, l1)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_stream_stop_cb_wrapper[0])) = v57
			m.G0 = v8 + int32(32)
			return
		}
	}
}
