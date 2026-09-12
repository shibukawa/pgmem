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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v79 float64
	_ = v79
	var v81 int32
	_ = v81
	var v86 float64
	_ = v86
	var v89 float64
	_ = v89
	var v91 float64
	_ = v91
	var v93 int32
	_ = v93
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v107 float64
	_ = v107
	var v109 float64
	_ = v109
	var v115 float64
	_ = v115
	var v117 float64
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v150 float64
	_ = v150
	var v152 int32
	_ = v152
	var v157 float64
	_ = v157
	var v160 float64
	_ = v160
	var v162 float64
	_ = v162
	var v164 int32
	_ = v164
	var v169 float64
	_ = v169
	var v171 float64
	_ = v171
	var v178 float64
	_ = v178
	var v180 float64
	_ = v180
	var v186 float64
	_ = v186
	var v188 float64
	_ = v188
	var v192 int32
	_ = v192
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v237 int32
	_ = v237
	var v238 float64
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v249 float64
	_ = v249
	var v251 int32
	_ = v251
	var v252 float64
	_ = v252
	var v255 float64
	_ = v255
	var v257 float64
	_ = v257
	var v260 float64
	_ = v260
	var v267 int32
	_ = v267
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v292 float64
	_ = v292
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v303 float64
	_ = v303
	var v305 int32
	_ = v305
	var v306 float64
	_ = v306
	var v309 float64
	_ = v309
	var v311 float64
	_ = v311
	var v314 float64
	_ = v314
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v353 float64
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v364 float64
	_ = v364
	var v366 int32
	_ = v366
	var v367 float64
	_ = v367
	var v378 float64
	_ = v378
	var v380 float64
	_ = v380
	var v381 float64
	_ = v381
	var v387 int32
	_ = v387
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v412 float64
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v423 float64
	_ = v423
	var v425 int32
	_ = v425
	var v426 float64
	_ = v426
	var v437 float64
	_ = v437
	var v439 float64
	_ = v439
	var v440 float64
	_ = v440
	var v447 int32
	_ = v447
	var v482 int32
	_ = v482
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v551 float64
	_ = v551
	var v560 float64
	_ = v560
	var v564 int32
	_ = v564
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v605 int32
	_ = v605
	var v607 float64
	_ = v607
	var v609 int32
	_ = v609
	var v614 float64
	_ = v614
	var v617 float64
	_ = v617
	var v619 float64
	_ = v619
	var v621 int32
	_ = v621
	var v626 float64
	_ = v626
	var v629 float64
	_ = v629
	var v630 int32
	_ = v630
	var v636 float64
	_ = v636
	var v638 float64
	_ = v638
	var v644 float64
	_ = v644
	var v646 float64
	_ = v646
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v726 float64
	_ = v726
	var v735 float64
	_ = v735
	var v739 int32
	_ = v739
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v780 int32
	_ = v780
	var v782 float64
	_ = v782
	var v784 int32
	_ = v784
	var v789 float64
	_ = v789
	var v792 float64
	_ = v792
	var v794 float64
	_ = v794
	var v796 int32
	_ = v796
	var v801 float64
	_ = v801
	var v804 float64
	_ = v804
	var v805 int32
	_ = v805
	var v811 float64
	_ = v811
	var v813 float64
	_ = v813
	var v819 float64
	_ = v819
	var v821 float64
	_ = v821
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v875 int32
	_ = v875
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v923 int32
	_ = v923
	var v925 float64
	_ = v925
	var v927 int32
	_ = v927
	var v932 float64
	_ = v932
	var v935 float64
	_ = v935
	var v937 float64
	_ = v937
	var v939 int32
	_ = v939
	var v944 float64
	_ = v944
	var v947 float64
	_ = v947
	var v948 int32
	_ = v948
	var v954 float64
	_ = v954
	var v956 float64
	_ = v956
	var v962 float64
	_ = v962
	var v964 float64
	_ = v964
	var v968 int32
	_ = v968
	var v976 int32
	_ = v976
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v1008 int32
	_ = v1008
	var v1009 float64
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1020 float64
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 float64
	_ = v1023
	var v1026 float64
	_ = v1026
	var v1028 float64
	_ = v1028
	var v1031 float64
	_ = v1031
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1051 int32
	_ = v1051
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1117 float64
	_ = v1117
	var v1126 float64
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1171 int32
	_ = v1171
	var v1173 float64
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1180 float64
	_ = v1180
	var v1183 float64
	_ = v1183
	var v1185 float64
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1192 float64
	_ = v1192
	var v1195 float64
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1202 float64
	_ = v1202
	var v1204 float64
	_ = v1204
	var v1210 float64
	_ = v1210
	var v1212 float64
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1223 int32
	_ = v1223
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
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
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1251 != v9 {
		goto L301
	} else {
		goto L302
	}
L5:
	;
	v1073 = int32(0)
	if v22 == v1073 {
		v1223 = v1073
		goto L264
	} else {
		goto L265
	}
L6:
	;
	v875 = int32(0)
	if v22 == v875 {
		v1051 = v875
		goto L213
	} else {
		goto L214
	}
L7:
	;
	switch v13&int32(65535) - int32(3) {
	case 0:
		goto L6
	default:
		v1248 = v2
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
	v858 = v13 & int32(65535)
	if base.Ui32(int32(14)) < base.Ui32(v858) {
		v1248 = v2
		goto L4
	} else {
		goto L208
	}
L10:
	;
	v682 = int32(0)
	if v9 == v682 {
		v832 = v682
		goto L171
	} else {
		goto L172
	}
L11:
	;
	v507 = int32(0)
	if v22 == v507 {
		v657 = v507
		goto L133
	} else {
		goto L134
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v47 = int32(2147483647)
	v48 = v46 & v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v51 = v49 & v47
	if base.Ui32(v48) < base.Ui32(v51) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v1248 = base.B2i32(v504 == int32(0))
	goto L4
L14:
	;
	v504 = v482
	goto L13
L15:
	;
	v482 = int32(1)
	goto L14
L16:
	;
	v53 = v48
	goto L18
L17:
	;
	v53 = v51
	goto L18
L18:
	;
	if v53 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v54 = int32(8)
	v55 = v9 + v54
	v57 = v22 + v54
	v65 = int32(0)
	goto L22
L20:
	;
	goto L21
L21:
	;
	if base.Ui32(v51) < base.Ui32(v48) {
		goto L56
	} else {
		goto L57
	}
L22:
	;
	v77 = v65 << (uint(int32(3)) % 32)
	v79 = *(*float64)(unsafe.Add(mBase, uint32(v57+v77)))
	v81 = base.B2i32(v46 < int32(0))
	if v46 < int32(0) {
		v89 = v79
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v124 = int32(8)
	v125 = v9 + v124
	v127 = v22 + v124
	v136 = int32(0)
	goto L39
L24:
	;
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v55+v77)))
	v93 = base.B2i32(v49 < int32(0))
	if v49 < int32(0) {
		v100 = v91
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v86 = *(*float64)(unsafe.Add(mBase, uint32(v57+(v65+v46)<<(uint(int32(3))%32))))
	if base.F64_lt(v79, v86) != 0 {
		v89 = v79
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v89 = v86
	goto L24
L27:
	;
	if base.F64_gt(v89, v100) != 0 {
		goto L15
	} else {
		goto L30
	}
L28:
	;
	v98 = *(*float64)(unsafe.Add(mBase, uint32(v55+(v65+v49)<<(uint(int32(3))%32))))
	if base.F64_lt(v91, v98) != 0 {
		v100 = v91
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v100 = v98
	goto L27
L30:
	;
	if v46 < int32(0) {
		v109 = v79
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v49 < int32(0) {
		v117 = v91
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v107 = *(*float64)(unsafe.Add(mBase, uint32(v57+(v65+v46)<<(uint(int32(3))%32))))
	if base.F64_lt(v79, v107) != 0 {
		v109 = v79
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v109 = v107
	goto L31
L34:
	;
	v119 = int32(-1)
	if base.F64_lt(v109, v117) != 0 {
		v482 = v119
		goto L14
	} else {
		goto L37
	}
L35:
	;
	v115 = *(*float64)(unsafe.Add(mBase, uint32(v55+(v65+v49)<<(uint(int32(3))%32))))
	if base.F64_lt(v91, v115) != 0 {
		v117 = v91
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v117 = v115
	goto L34
L37:
	;
	v122 = v65 + int32(1)
	if v122 != v53 {
		v65 = v122
		goto L22
	} else {
		goto L38
	}
L38:
	;
	goto L23
L39:
	;
	v148 = v136 << (uint(int32(3)) % 32)
	v150 = *(*float64)(unsafe.Add(mBase, uint32(v127+v148)))
	v152 = base.B2i32(v46 < int32(0))
	if v46 < int32(0) {
		v160 = v150
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L21
L41:
	;
	v162 = *(*float64)(unsafe.Add(mBase, uint32(v125+v148)))
	v164 = base.B2i32(v49 < int32(0))
	if v49 < int32(0) {
		v171 = v162
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v157 = *(*float64)(unsafe.Add(mBase, uint32(v127+(v136+v46)<<(uint(int32(3))%32))))
	if base.F64_gt(v150, v157) != 0 {
		v160 = v150
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v160 = v157
	goto L41
L44:
	;
	if base.F64_gt(v160, v171) != 0 {
		goto L15
	} else {
		goto L47
	}
L45:
	;
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v125+(v136+v49)<<(uint(int32(3))%32))))
	if base.F64_gt(v162, v169) != 0 {
		v171 = v162
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v171 = v169
	goto L44
L47:
	;
	if v46 < int32(0) {
		v180 = v150
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v49 < int32(0) {
		v188 = v162
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v178 = *(*float64)(unsafe.Add(mBase, uint32(v127+(v136+v46)<<(uint(int32(3))%32))))
	if base.F64_gt(v150, v178) != 0 {
		v180 = v150
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v180 = v178
	goto L48
L51:
	;
	if base.F64_lt(v180, v188) != 0 {
		v482 = v119
		goto L14
	} else {
		goto L54
	}
L52:
	;
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v125+(v136+v49)<<(uint(int32(3))%32))))
	if base.F64_gt(v162, v186) != 0 {
		v188 = v162
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v188 = v186
	goto L51
L54:
	;
	v192 = v136 + int32(1)
	if v192 != v53 {
		v136 = v192
		goto L39
	} else {
		goto L55
	}
L55:
	;
	goto L40
L56:
	;
	v214 = v22 + int32(8)
	v224 = v53
	goto L59
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(v51) <= base.Ui32(v48) {
		goto L95
	} else {
		goto L96
	}
L59:
	;
	v237 = v214 + v224<<(uint(int32(3))%32)
	v238 = *(*float64)(unsafe.Add(mBase, uint32(v237)))
	if base.B2i32(v46 < int32(0)) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v283 = v53
	goto L77
L61:
	;
	if base.F64_lt(v260, float64(0)) != 0 {
		goto L73
	} else {
		goto L74
	}
L62:
	;
	v242 = int32(3)
	v244 = v214 + (v224+v46)<<(uint(v242)%32)
	v249 = *(*float64)(unsafe.Add(mBase, uint32(v214+(v224+v48)<<(uint(v242)%32))))
	if base.F64_lt(v238, v249) != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	if base.F64_gt(v238, float64(0)) != 0 {
		goto L15
	} else {
		goto L72
	}
L65:
	;
	v251 = v237
	goto L67
L66:
	;
	v251 = v244
	goto L67
L67:
	;
	v252 = *(*float64)(unsafe.Add(mBase, uint32(v251)))
	if base.F64_gt(v252, float64(0)) != 0 {
		goto L15
	} else {
		goto L68
	}
L68:
	;
	v255 = *(*float64)(unsafe.Add(mBase, uint32(v244)))
	if base.F64_lt(v238, v255) != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v257 = v238
	goto L71
L70:
	;
	v257 = v255
	goto L71
L71:
	;
	v260 = v257
	goto L61
L72:
	;
	v260 = v238
	goto L61
L73:
	;
	v504 = int32(-1)
	goto L13
L74:
	;
	goto L75
L75:
	;
	v267 = v224 + int32(1)
	if v267 != v48 {
		v224 = v267
		goto L59
	} else {
		goto L76
	}
L76:
	;
	goto L60
L77:
	;
	v291 = v214 + v283<<(uint(int32(3))%32)
	v292 = *(*float64)(unsafe.Add(mBase, uint32(v291)))
	if base.B2i32(v46 < int32(0)) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v504 = int32(-1)
	goto L13
L79:
	;
	if base.F64_lt(v314, float64(0)) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L80:
	;
	v296 = int32(3)
	v298 = v214 + (v46+v283)<<(uint(v296)%32)
	v303 = *(*float64)(unsafe.Add(mBase, uint32(v214+(v283+v48)<<(uint(v296)%32))))
	if base.F64_gt(v292, v303) != 0 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	if base.F64_gt(v292, float64(0)) != 0 {
		goto L15
	} else {
		goto L90
	}
L83:
	;
	v305 = v291
	goto L85
L84:
	;
	v305 = v298
	goto L85
L85:
	;
	v306 = *(*float64)(unsafe.Add(mBase, uint32(v305)))
	if base.F64_gt(v306, float64(0)) != 0 {
		goto L15
	} else {
		goto L86
	}
L86:
	;
	v309 = *(*float64)(unsafe.Add(mBase, uint32(v298)))
	if base.F64_gt(v292, v309) != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v311 = v292
	goto L89
L88:
	;
	v311 = v309
	goto L89
L89:
	;
	v314 = v311
	goto L79
L90:
	;
	v314 = v292
	goto L79
L91:
	;
	v321 = int32(1)
	v323 = v283 + v321
	if v323 == v48 {
		v482 = v321
		goto L14
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	goto L78
L94:
	;
	v283 = v323
	goto L77
L95:
	;
	v504 = int32(0)
	goto L13
L96:
	;
	goto L97
L97:
	;
	v329 = v9 + int32(8)
	v345 = v48
	goto L98
L98:
	;
	v352 = v329 + v345<<(uint(int32(3))%32)
	v353 = *(*float64)(unsafe.Add(mBase, uint32(v352)))
	if base.B2i32(v49 < int32(0)) == int32(0) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	v403 = v53
	goto L115
L100:
	;
	if base.F64_lt(v381, float64(0)) != 0 {
		goto L15
	} else {
		goto L113
	}
L101:
	;
	v378 = *(*float64)(unsafe.Add(mBase, uint32(v359)))
	if base.F64_lt(v353, v378) != 0 {
		goto L110
	} else {
		goto L111
	}
L102:
	;
	v357 = int32(3)
	v359 = v329 + (v49+v345)<<(uint(v357)%32)
	v364 = *(*float64)(unsafe.Add(mBase, uint32(v329+(v345+v51)<<(uint(v357)%32))))
	if base.F64_lt(v353, v364) != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	if base.F64_gt(v353, float64(0)) == int32(0) {
		v381 = v353
		goto L100
	} else {
		goto L109
	}
L105:
	;
	v366 = v352
	goto L107
L106:
	;
	v366 = v359
	goto L107
L107:
	;
	v367 = *(*float64)(unsafe.Add(mBase, uint32(v366)))
	if base.F64_gt(v367, float64(0)) == int32(0) {
		goto L101
	} else {
		goto L108
	}
L108:
	;
	v504 = int32(-1)
	goto L13
L109:
	;
	v504 = int32(-1)
	goto L13
L110:
	;
	v380 = v353
	goto L112
L111:
	;
	v380 = v378
	goto L112
L112:
	;
	v381 = v380
	goto L100
L113:
	;
	v387 = v345 + int32(1)
	if v387 != v51 {
		v345 = v387
		goto L98
	} else {
		goto L114
	}
L114:
	;
	goto L99
L115:
	;
	v411 = v329 + v403<<(uint(int32(3))%32)
	v412 = *(*float64)(unsafe.Add(mBase, uint32(v411)))
	if base.B2i32(v49 < int32(0)) == int32(0) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v482 = int32(-1)
	goto L14
L117:
	;
	if base.F64_lt(v440, float64(0)) != 0 {
		goto L15
	} else {
		goto L130
	}
L118:
	;
	v437 = *(*float64)(unsafe.Add(mBase, uint32(v418)))
	if base.F64_gt(v412, v437) != 0 {
		goto L127
	} else {
		goto L128
	}
L119:
	;
	v416 = int32(3)
	v418 = v329 + (v49+v403)<<(uint(v416)%32)
	v423 = *(*float64)(unsafe.Add(mBase, uint32(v329+(v403+v51)<<(uint(v416)%32))))
	if base.F64_gt(v412, v423) != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	if base.F64_gt(v412, float64(0)) == int32(0) {
		v440 = v412
		goto L117
	} else {
		goto L126
	}
L122:
	;
	v425 = v411
	goto L124
L123:
	;
	v425 = v418
	goto L124
L124:
	;
	v426 = *(*float64)(unsafe.Add(mBase, uint32(v425)))
	if base.F64_gt(v426, float64(0)) == int32(0) {
		goto L118
	} else {
		goto L125
	}
L125:
	;
	v504 = int32(-1)
	goto L13
L126:
	;
	v504 = int32(-1)
	goto L13
L127:
	;
	v439 = v412
	goto L129
L128:
	;
	v439 = v437
	goto L129
L129:
	;
	v440 = v439
	goto L117
L130:
	;
	v447 = v403 + int32(1)
	if v51 != v447 {
		v403 = v447
		goto L115
	} else {
		goto L131
	}
L131:
	;
	goto L116
L132:
	;
	v1248 = v681
	goto L4
L133:
	;
	v681 = v657
	goto L132
L134:
	;
	if v9 == int32(0) {
		v657 = v507
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v524 = int32(2147483647)
	v525 = v523 & v524
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v528 = v526 & v524
	if base.Ui32(v525) < base.Ui32(v528) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v531 = v9 + int32(8)
	v536 = v525
	goto L139
L137:
	;
	goto L138
L138:
	;
	if base.Ui32(v525) < base.Ui32(v528) {
		goto L147
	} else {
		goto L148
	}
L139:
	;
	v551 = *(*float64)(unsafe.Add(mBase, uint32(v531+v536<<(uint(int32(3))%32))))
	if base.F64_ne(v551, float64(0)) != 0 {
		v657 = v507
		goto L133
	} else {
		goto L141
	}
L140:
	;
	goto L138
L141:
	;
	if base.B2i32(v526 < int32(0)) == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v560 = *(*float64)(unsafe.Add(mBase, uint32(v531+(v536+v528)<<(uint(int32(3))%32))))
	if base.F64_ne(v560, float64(0)) != 0 {
		v657 = v507
		goto L133
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v564 = v536 + int32(1)
	if v564 != v528 {
		v536 = v564
		goto L139
	} else {
		goto L146
	}
L145:
	;
	goto L144
L146:
	;
	goto L140
L147:
	;
	v581 = v525
	goto L149
L148:
	;
	v581 = v528
	goto L149
L149:
	;
	if v581 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v681 = int32(1)
	goto L132
L151:
	;
	goto L152
L152:
	;
	v585 = int32(8)
	v586 = v9 + v585
	v588 = v22 + v585
	v592 = int32(0)
	goto L153
L153:
	;
	v605 = v592 << (uint(int32(3)) % 32)
	v607 = *(*float64)(unsafe.Add(mBase, uint32(v588+v605)))
	v609 = base.B2i32(v523 < int32(0))
	if v523 < int32(0) {
		v617 = v607
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v657 = v649
	goto L133
L155:
	;
	v619 = *(*float64)(unsafe.Add(mBase, uint32(v586+v605)))
	v621 = base.B2i32(v526 < int32(0))
	if v526 < int32(0) {
		v629 = v619
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v614 = *(*float64)(unsafe.Add(mBase, uint32(v588+(v592+v523)<<(uint(int32(3))%32))))
	if base.F64_lt(v607, v614) != 0 {
		v617 = v607
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v617 = v614
	goto L155
L158:
	;
	v630 = int32(0)
	if base.F64_gt(v617, v629) != 0 {
		v657 = v630
		goto L133
	} else {
		goto L161
	}
L159:
	;
	v626 = *(*float64)(unsafe.Add(mBase, uint32(v586+(v592+v526)<<(uint(int32(3))%32))))
	if base.F64_lt(v619, v626) != 0 {
		v629 = v619
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v629 = v626
	goto L158
L161:
	;
	if v523 < int32(0) {
		v638 = v607
		goto L162
	} else {
		goto L163
	}
L162:
	;
	if v526 < int32(0) {
		v646 = v619
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v636 = *(*float64)(unsafe.Add(mBase, uint32(v588+(v592+v523)<<(uint(int32(3))%32))))
	if base.F64_gt(v607, v636) != 0 {
		v638 = v607
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v638 = v636
	goto L162
L165:
	;
	if base.F64_gt(v646, v638) != 0 {
		v657 = v630
		goto L133
	} else {
		goto L168
	}
L166:
	;
	v644 = *(*float64)(unsafe.Add(mBase, uint32(v586+(v592+v526)<<(uint(int32(3))%32))))
	if base.F64_gt(v619, v644) != 0 {
		v646 = v619
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v646 = v644
	goto L165
L168:
	;
	v649 = int32(1)
	v651 = v592 + v649
	if v651 != v581 {
		v592 = v651
		goto L153
	} else {
		goto L169
	}
L169:
	;
	goto L154
L170:
	;
	v1248 = v856
	goto L4
L171:
	;
	v856 = v832
	goto L170
L172:
	;
	if v22 == int32(0) {
		v832 = v682
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v699 = int32(2147483647)
	v700 = v698 & v699
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v703 = v701 & v699
	if base.Ui32(v700) < base.Ui32(v703) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v706 = v22 + int32(8)
	v711 = v700
	goto L177
L175:
	;
	goto L176
L176:
	;
	if base.Ui32(v700) < base.Ui32(v703) {
		goto L185
	} else {
		goto L186
	}
L177:
	;
	v726 = *(*float64)(unsafe.Add(mBase, uint32(v706+v711<<(uint(int32(3))%32))))
	if base.F64_ne(v726, float64(0)) != 0 {
		v832 = v682
		goto L171
	} else {
		goto L179
	}
L178:
	;
	goto L176
L179:
	;
	if base.B2i32(v701 < int32(0)) == int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v735 = *(*float64)(unsafe.Add(mBase, uint32(v706+(v711+v703)<<(uint(int32(3))%32))))
	if base.F64_ne(v735, float64(0)) != 0 {
		v832 = v682
		goto L171
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v739 = v711 + int32(1)
	if v739 != v703 {
		v711 = v739
		goto L177
	} else {
		goto L184
	}
L183:
	;
	goto L182
L184:
	;
	goto L178
L185:
	;
	v756 = v700
	goto L187
L186:
	;
	v756 = v703
	goto L187
L187:
	;
	if v756 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v856 = int32(1)
	goto L170
L189:
	;
	goto L190
L190:
	;
	v760 = int32(8)
	v761 = v22 + v760
	v763 = v9 + v760
	v767 = int32(0)
	goto L191
L191:
	;
	v780 = v767 << (uint(int32(3)) % 32)
	v782 = *(*float64)(unsafe.Add(mBase, uint32(v763+v780)))
	v784 = base.B2i32(v698 < int32(0))
	if v698 < int32(0) {
		v792 = v782
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v832 = v824
	goto L171
L193:
	;
	v794 = *(*float64)(unsafe.Add(mBase, uint32(v761+v780)))
	v796 = base.B2i32(v701 < int32(0))
	if v701 < int32(0) {
		v804 = v794
		goto L196
	} else {
		goto L197
	}
L194:
	;
	v789 = *(*float64)(unsafe.Add(mBase, uint32(v763+(v767+v698)<<(uint(int32(3))%32))))
	if base.F64_lt(v782, v789) != 0 {
		v792 = v782
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v792 = v789
	goto L193
L196:
	;
	v805 = int32(0)
	if base.F64_gt(v792, v804) != 0 {
		v832 = v805
		goto L171
	} else {
		goto L199
	}
L197:
	;
	v801 = *(*float64)(unsafe.Add(mBase, uint32(v761+(v767+v701)<<(uint(int32(3))%32))))
	if base.F64_lt(v794, v801) != 0 {
		v804 = v794
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v804 = v801
	goto L196
L199:
	;
	if v698 < int32(0) {
		v813 = v782
		goto L200
	} else {
		goto L201
	}
L200:
	;
	if v701 < int32(0) {
		v821 = v794
		goto L203
	} else {
		goto L204
	}
L201:
	;
	v811 = *(*float64)(unsafe.Add(mBase, uint32(v763+(v767+v698)<<(uint(int32(3))%32))))
	if base.F64_gt(v782, v811) != 0 {
		v813 = v782
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v813 = v811
	goto L200
L203:
	;
	if base.F64_gt(v821, v813) != 0 {
		v832 = v805
		goto L171
	} else {
		goto L206
	}
L204:
	;
	v819 = *(*float64)(unsafe.Add(mBase, uint32(v761+(v767+v701)<<(uint(int32(3))%32))))
	if base.F64_gt(v794, v819) != 0 {
		v821 = v794
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v821 = v819
	goto L203
L206:
	;
	v824 = int32(1)
	v826 = v767 + v824
	if v826 != v756 {
		v767 = v826
		goto L191
	} else {
		goto L207
	}
L207:
	;
	goto L192
L208:
	;
	v862 = int32(1) << (uint(v858) % 32)
	if v862&int32(8384) != 0 {
		goto L5
	} else {
		goto L209
	}
L209:
	;
	if v862&int32(16640) != 0 {
		goto L6
	} else {
		goto L210
	}
L210:
	;
	if v858 != int32(3) {
		v1248 = v2
		goto L4
	} else {
		goto L211
	}
L211:
	;
	goto L6
L212:
	;
	v1248 = v1072
	goto L4
L213:
	;
	v1072 = v1051
	goto L212
L214:
	;
	if v9 == int32(0) {
		v1051 = v875
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v888 = int32(2147483647)
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v893 = base.B2i32(base.Ui32(v887&v888) < base.Ui32(v890&v888))
	if base.Ui32(v887&v888) < base.Ui32(v890&v888) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v894 = v9
	goto L218
L217:
	;
	v894 = v22
	goto L218
L218:
	;
	if base.Ui32(v887&v888) < base.Ui32(v890&v888) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v985 = v976 & int32(2147483647)
	if base.Ui32(v985) <= base.Ui32(v898) {
		goto L243
	} else {
		goto L244
	}
L220:
	;
	v895 = v22
	goto L222
L221:
	;
	v895 = v9
	goto L222
L222:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v895)+4))
	v898 = v896 & int32(2147483647)
	if v898 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v894)+4))
	v976 = v901
	goto L219
L224:
	;
	goto L225
L225:
	;
	v902 = int32(8)
	v903 = v895 + v902
	v905 = v894 + v902
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v894)+4))
	v908 = int32(0)
	goto L226
L226:
	;
	v923 = v908 << (uint(int32(3)) % 32)
	v925 = *(*float64)(unsafe.Add(mBase, uint32(v905+v923)))
	v927 = base.B2i32(v906 < int32(0))
	if v906 < int32(0) {
		v935 = v925
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v976 = v906
	goto L219
L228:
	;
	v937 = *(*float64)(unsafe.Add(mBase, uint32(v923+v903)))
	v939 = base.B2i32(v896 < int32(0))
	if v896 < int32(0) {
		v947 = v937
		goto L231
	} else {
		goto L232
	}
L229:
	;
	v932 = *(*float64)(unsafe.Add(mBase, uint32(v905+(v908+v906)<<(uint(int32(3))%32))))
	if base.F64_lt(v925, v932) != 0 {
		v935 = v925
		goto L228
	} else {
		goto L230
	}
L230:
	;
	v935 = v932
	goto L228
L231:
	;
	v948 = int32(0)
	if base.F64_gt(v935, v947) != 0 {
		v1051 = v948
		goto L213
	} else {
		goto L234
	}
L232:
	;
	v944 = *(*float64)(unsafe.Add(mBase, uint32(v903+(v908+v896)<<(uint(int32(3))%32))))
	if base.F64_gt(v937, v944) != 0 {
		v947 = v937
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v947 = v944
	goto L231
L234:
	;
	if v906 < int32(0) {
		v956 = v925
		goto L235
	} else {
		goto L236
	}
L235:
	;
	if v896 < int32(0) {
		v964 = v937
		goto L238
	} else {
		goto L239
	}
L236:
	;
	v954 = *(*float64)(unsafe.Add(mBase, uint32(v905+(v908+v906)<<(uint(int32(3))%32))))
	if base.F64_gt(v925, v954) != 0 {
		v956 = v925
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v956 = v954
	goto L235
L238:
	;
	if base.F64_lt(v956, v964) != 0 {
		v1051 = v948
		goto L213
	} else {
		goto L241
	}
L239:
	;
	v962 = *(*float64)(unsafe.Add(mBase, uint32(v903+(v908+v896)<<(uint(int32(3))%32))))
	if base.F64_lt(v937, v962) != 0 {
		v964 = v937
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v964 = v962
	goto L238
L241:
	;
	v968 = v908 + int32(1)
	if v968 != v898 {
		v908 = v968
		goto L226
	} else {
		goto L242
	}
L242:
	;
	goto L227
L243:
	;
	v1072 = int32(1)
	goto L212
L244:
	;
	goto L245
L245:
	;
	v989 = v894 + int32(8)
	v993 = v898
	goto L246
L246:
	;
	v1008 = v989 + v993<<(uint(int32(3))%32)
	v1009 = *(*float64)(unsafe.Add(mBase, uint32(v1008)))
	if base.B2i32(v976 < int32(0)) == int32(0) {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	v1051 = int32(0)
	goto L213
L248:
	;
	goto L247
L249:
	;
	if base.F64_lt(v1031, float64(0)) != 0 {
		goto L248
	} else {
		goto L261
	}
L250:
	;
	v1013 = int32(3)
	v1015 = v989 + (v993+v976)<<(uint(v1013)%32)
	v1020 = *(*float64)(unsafe.Add(mBase, uint32(v989+(v993+v985)<<(uint(v1013)%32))))
	if base.F64_lt(v1009, v1020) != 0 {
		goto L253
	} else {
		goto L254
	}
L251:
	;
	goto L252
L252:
	;
	if base.F64_gt(v1009, float64(0)) != 0 {
		goto L248
	} else {
		goto L260
	}
L253:
	;
	v1022 = v1008
	goto L255
L254:
	;
	v1022 = v1015
	goto L255
L255:
	;
	v1023 = *(*float64)(unsafe.Add(mBase, uint32(v1022)))
	if base.F64_gt(v1023, float64(0)) != 0 {
		goto L248
	} else {
		goto L256
	}
L256:
	;
	v1026 = *(*float64)(unsafe.Add(mBase, uint32(v1015)))
	if base.F64_gt(v1009, v1026) != 0 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1028 = v1009
	goto L259
L258:
	;
	v1028 = v1026
	goto L259
L259:
	;
	v1031 = v1028
	goto L249
L260:
	;
	v1031 = v1009
	goto L249
L261:
	;
	v1036 = int32(1)
	v1038 = v993 + v1036
	if v985 != v1038 {
		v993 = v1038
		goto L246
	} else {
		goto L262
	}
L262:
	;
	v1051 = v1036
	goto L213
L263:
	;
	v1248 = v1247
	goto L4
L264:
	;
	v1247 = v1223
	goto L263
L265:
	;
	if v9 == int32(0) {
		v1223 = v1073
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v1090 = int32(2147483647)
	v1091 = v1089 & v1090
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v1094 = v1092 & v1090
	if base.Ui32(v1091) < base.Ui32(v1094) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1097 = v9 + int32(8)
	v1102 = v1091
	goto L270
L268:
	;
	goto L269
L269:
	;
	if base.Ui32(v1091) < base.Ui32(v1094) {
		goto L278
	} else {
		goto L279
	}
L270:
	;
	v1117 = *(*float64)(unsafe.Add(mBase, uint32(v1097+v1102<<(uint(int32(3))%32))))
	if base.F64_ne(v1117, float64(0)) != 0 {
		v1223 = v1073
		goto L264
	} else {
		goto L272
	}
L271:
	;
	goto L269
L272:
	;
	if base.B2i32(v1092 < int32(0)) == int32(0) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1126 = *(*float64)(unsafe.Add(mBase, uint32(v1097+(v1102+v1094)<<(uint(int32(3))%32))))
	if base.F64_ne(v1126, float64(0)) != 0 {
		v1223 = v1073
		goto L264
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1130 = v1102 + int32(1)
	if v1130 != v1094 {
		v1102 = v1130
		goto L270
	} else {
		goto L277
	}
L276:
	;
	goto L275
L277:
	;
	goto L271
L278:
	;
	v1147 = v1091
	goto L280
L279:
	;
	v1147 = v1094
	goto L280
L280:
	;
	if v1147 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1247 = int32(1)
	goto L263
L282:
	;
	goto L283
L283:
	;
	v1151 = int32(8)
	v1152 = v9 + v1151
	v1154 = v22 + v1151
	v1158 = int32(0)
	goto L284
L284:
	;
	v1171 = v1158 << (uint(int32(3)) % 32)
	v1173 = *(*float64)(unsafe.Add(mBase, uint32(v1154+v1171)))
	v1175 = base.B2i32(v1089 < int32(0))
	if v1089 < int32(0) {
		v1183 = v1173
		goto L286
	} else {
		goto L287
	}
L285:
	;
	v1223 = v1215
	goto L264
L286:
	;
	v1185 = *(*float64)(unsafe.Add(mBase, uint32(v1152+v1171)))
	v1187 = base.B2i32(v1092 < int32(0))
	if v1092 < int32(0) {
		v1195 = v1185
		goto L289
	} else {
		goto L290
	}
L287:
	;
	v1180 = *(*float64)(unsafe.Add(mBase, uint32(v1154+(v1158+v1089)<<(uint(int32(3))%32))))
	if base.F64_lt(v1173, v1180) != 0 {
		v1183 = v1173
		goto L286
	} else {
		goto L288
	}
L288:
	;
	v1183 = v1180
	goto L286
L289:
	;
	v1196 = int32(0)
	if base.F64_gt(v1183, v1195) != 0 {
		v1223 = v1196
		goto L264
	} else {
		goto L292
	}
L290:
	;
	v1192 = *(*float64)(unsafe.Add(mBase, uint32(v1152+(v1158+v1092)<<(uint(int32(3))%32))))
	if base.F64_lt(v1185, v1192) != 0 {
		v1195 = v1185
		goto L289
	} else {
		goto L291
	}
L291:
	;
	v1195 = v1192
	goto L289
L292:
	;
	if v1089 < int32(0) {
		v1204 = v1173
		goto L293
	} else {
		goto L294
	}
L293:
	;
	if v1092 < int32(0) {
		v1212 = v1185
		goto L296
	} else {
		goto L297
	}
L294:
	;
	v1202 = *(*float64)(unsafe.Add(mBase, uint32(v1154+(v1158+v1089)<<(uint(int32(3))%32))))
	if base.F64_gt(v1173, v1202) != 0 {
		v1204 = v1173
		goto L293
	} else {
		goto L295
	}
L295:
	;
	v1204 = v1202
	goto L293
L296:
	;
	if base.F64_gt(v1212, v1204) != 0 {
		v1223 = v1196
		goto L264
	} else {
		goto L299
	}
L297:
	;
	v1210 = *(*float64)(unsafe.Add(mBase, uint32(v1152+(v1158+v1092)<<(uint(int32(3))%32))))
	if base.F64_gt(v1185, v1210) != 0 {
		v1212 = v1185
		goto L296
	} else {
		goto L298
	}
L298:
	;
	v1212 = v1210
	goto L296
L299:
	;
	v1215 = int32(1)
	v1217 = v1158 + v1215
	if v1217 != v1147 {
		v1158 = v1217
		goto L284
	} else {
		goto L300
	}
L300:
	;
	goto L285
L301:
	;
	F_pfree(m, v9)
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L1
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	return v1248
L304:
	;
	goto L303
}
