package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_cube_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 float64
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 float64
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 float64
	_ = v109
	var v111 float64
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v160 float64
	_ = v160
	var v165 int32
	_ = v165
	var v168 float64
	_ = v168
	var v172 float64
	_ = v172
	var v177 int32
	_ = v177
	var v182 float64
	_ = v182
	var v186 float64
	_ = v186
	var v189 float64
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v216 float64
	_ = v216
	var v223 int32
	_ = v223
	var v226 float64
	_ = v226
	var v230 float64
	_ = v230
	var v255 float64
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v307 float64
	_ = v307
	var v311 int32
	_ = v311
	var v314 float64
	_ = v314
	var v318 float64
	_ = v318
	var v323 int32
	_ = v323
	var v328 float64
	_ = v328
	var v332 float64
	_ = v332
	var v335 float64
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v363 float64
	_ = v363
	var v369 int32
	_ = v369
	var v372 float64
	_ = v372
	var v376 float64
	_ = v376
	var v402 float64
	_ = v402
	var v405 float64
	_ = v405
	var v409 int32
	_ = v409
	var v410 float64
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 float64
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v503 float64
	_ = v503
	var v506 int32
	_ = v506
	var v509 float64
	_ = v509
	var v513 float64
	_ = v513
	var v518 int32
	_ = v518
	var v523 float64
	_ = v523
	var v527 float64
	_ = v527
	var v530 float64
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v559 float64
	_ = v559
	var v564 int32
	_ = v564
	var v567 float64
	_ = v567
	var v571 float64
	_ = v571
	var v598 float64
	_ = v598
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v650 float64
	_ = v650
	var v652 int32
	_ = v652
	var v655 float64
	_ = v655
	var v659 float64
	_ = v659
	var v664 int32
	_ = v664
	var v669 float64
	_ = v669
	var v673 float64
	_ = v673
	var v676 float64
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v706 float64
	_ = v706
	var v710 int32
	_ = v710
	var v713 float64
	_ = v713
	var v717 float64
	_ = v717
	var v745 float64
	_ = v745
	var v750 int32
	_ = v750
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v779 float64
	_ = v779
	var v780 float64
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v797 int32
	_ = v797
	var v803 float64
	_ = v803
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 float64
	_ = v814
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v839 int32
	_ = v839
	var v854 float64
	_ = v854
	var v859 int32
	_ = v859
	var v862 float64
	_ = v862
	var v866 float64
	_ = v866
	var v871 int32
	_ = v871
	var v876 float64
	_ = v876
	var v880 float64
	_ = v880
	var v883 float64
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v910 float64
	_ = v910
	var v917 int32
	_ = v917
	var v920 float64
	_ = v920
	var v924 float64
	_ = v924
	var v949 float64
	_ = v949
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v977 int32
	_ = v977
	var v993 float64
	_ = v993
	var v997 int32
	_ = v997
	var v1000 float64
	_ = v1000
	var v1004 float64
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1014 float64
	_ = v1014
	var v1018 float64
	_ = v1018
	var v1021 float64
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1049 float64
	_ = v1049
	var v1055 int32
	_ = v1055
	var v1058 float64
	_ = v1058
	var v1062 float64
	_ = v1062
	var v1088 float64
	_ = v1088
	var v1095 int32
	_ = v1095
	var v1102 int32
	_ = v1102
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1131 float64
	_ = v1131
	var v1132 float64
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1163 int32
	_ = v1163
	v22 = float64(0)
	v26 = int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v33 = (v29 + int32(_a_F_g_cube_picksplit_0)) & int32(_a_F_g_cube_picksplit_1)
	v37 = v33<<(uint(v26)%32) + int32(4)
	v38 = F_palloc(m, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v38
	v43 = F_palloc(m, v37)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v43
	v46 = int32(2)
	if base.Ui32(v46) <= base.Ui32(v33) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v50 = v28 + int32(4)
	v51 = int32(1)
	v60 = v51
	v63 = v51
	v68 = v26
	v69 = v46
	v76 = v22
	goto L7
L5:
	;
	v426 = v43
	v436 = v26
	v437 = v46
	goto L6
L6:
	;
	v446 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v446
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v451 = float64(0)
	v452 = int32(4)
	v453 = v28 + v452
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v453+v436&int32(_a_F_g_cube_picksplit_1)<<(uint(v452)%32))))
	v460 = F_pg_detoast_datum(m, v459)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L50
	}
L7:
	;
	v80 = v50 + v63<<(uint(int32(4))%32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v82 = F_pg_detoast_datum(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v426 = v420
	v436 = v412
	v437 = v411
	goto L6
L9:
	;
	v85 = v63 + int32(1)
	v86 = v85
	v87 = v85
	v93 = v60
	v101 = v68
	v102 = v69
	v109 = v76
	goto L10
L10:
	;
	v111 = float64(0)
	v115 = v50 + v86<<(uint(int32(4))%32)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v117 = F_pg_detoast_datum(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	if v85 != v33 {
		v60 = v413
		v63 = v85
		v68 = v412
		v69 = v411
		v76 = v410
		goto L7
	} else {
		goto L48
	}
L12:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v263 = F_DirectFunctionCall2Coll(m, int32(_a_F_g_cube_picksplit_2), int32(0), v261, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L26
	}
L13:
	;
	v119 = F_cube_union_v0(m, v82, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v119 == int32(0) {
		v255 = v111
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v123 <= int32(0) {
		v255 = v111
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v126 = int32(1)
	v129 = v119 + int32(8)
	if v123 == v126 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v123&v126 == int32(0) {
		v255 = v216
		goto L12
	} else {
		goto L24
	}
L18:
	;
	v195 = int32(0)
	v216 = float64(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v137 = int32(0)
	v139 = v137
	v142 = v137
	v160 = float64(1)
	goto L21
L21:
	;
	v165 = int32(3)
	v168 = *(*float64)(unsafe.Add(mBase, uint32(v129+(v139+v123)<<(uint(v165)%32))))
	v172 = *(*float64)(unsafe.Add(mBase, uint32(v129+v139<<(uint(v165)%32))))
	v177 = v139 | int32(1)
	v182 = *(*float64)(unsafe.Add(mBase, uint32(v129+(v177+v123)<<(uint(v165)%32))))
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v129+v177<<(uint(v165)%32))))
	v189 = base.F64_mul(base.F64_mul(v160, base.F64_abs(base.F64_sub(v168, v172))), base.F64_abs(base.F64_sub(v182, v186)))
	v190 = int32(2)
	v191 = v139 + v190
	v193 = v142 + v190
	if v193 != v123&int32(2147483646) {
		v139 = v191
		v142 = v193
		v160 = v189
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v195 = v191
	v216 = v189
	goto L17
L23:
	;
	goto L22
L24:
	;
	v223 = int32(3)
	v226 = *(*float64)(unsafe.Add(mBase, uint32(v129+(v195+v123)<<(uint(v223)%32))))
	v230 = *(*float64)(unsafe.Add(mBase, uint32(v129+v195<<(uint(v223)%32))))
	v255 = base.F64_mul(v216, base.F64_abs(base.F64_sub(v226, v230)))
	goto L12
L25:
	;
	v405 = base.F64_sub(v255, v402)
	v409 = (base.F64_gt(v405, v109) | v93) & int32(1)
	if v409 != 0 {
		goto L38
	} else {
		goto L39
	}
L26:
	;
	v265 = F_pg_detoast_datum(m, v263)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v265 == int32(0) {
		v402 = v111
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	if v269 <= int32(0) {
		v402 = v111
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v272 = int32(1)
	v275 = v265 + int32(8)
	if v269 == v272 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v269&v272 == int32(0) {
		v402 = v363
		goto L25
	} else {
		goto L37
	}
L31:
	;
	v341 = int32(0)
	v363 = float64(1)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v283 = int32(0)
	v285 = v283
	v288 = v283
	v307 = float64(1)
	goto L34
L34:
	;
	v311 = int32(3)
	v314 = *(*float64)(unsafe.Add(mBase, uint32(v275+(v285+v269)<<(uint(v311)%32))))
	v318 = *(*float64)(unsafe.Add(mBase, uint32(v275+v285<<(uint(v311)%32))))
	v323 = v285 | int32(1)
	v328 = *(*float64)(unsafe.Add(mBase, uint32(v275+(v323+v269)<<(uint(v311)%32))))
	v332 = *(*float64)(unsafe.Add(mBase, uint32(v275+v323<<(uint(v311)%32))))
	v335 = base.F64_mul(base.F64_mul(v307, base.F64_abs(base.F64_sub(v314, v318))), base.F64_abs(base.F64_sub(v328, v332)))
	v336 = int32(2)
	v337 = v285 + v336
	v339 = v288 + v336
	if v339 != v269&int32(2147483646) {
		v285 = v337
		v288 = v339
		v307 = v335
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v341 = v337
	v363 = v335
	goto L30
L36:
	;
	goto L35
L37:
	;
	v369 = int32(3)
	v372 = *(*float64)(unsafe.Add(mBase, uint32(v275+(v341+v269)<<(uint(v369)%32))))
	v376 = *(*float64)(unsafe.Add(mBase, uint32(v275+v341<<(uint(v369)%32))))
	v402 = base.F64_mul(v363, base.F64_abs(base.F64_sub(v372, v376)))
	goto L25
L38:
	;
	v410 = v405
	goto L40
L39:
	;
	v410 = v109
	goto L40
L40:
	;
	if v409 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v411 = v87
	goto L43
L42:
	;
	v411 = v102
	goto L43
L43:
	;
	if v409 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v412 = v63
	goto L46
L45:
	;
	v412 = v101
	goto L46
L46:
	;
	v413 = int32(0)
	v415 = v87 + int32(1)
	v417 = v415 & int32(_a_F_g_cube_picksplit_1)
	if base.Ui32(v417) <= base.Ui32(v33) {
		v86 = v417
		v87 = v415
		v93 = v413
		v101 = v412
		v102 = v411
		v109 = v410
		goto L10
	} else {
		goto L47
	}
L47:
	;
	goto L11
L48:
	;
	goto L8
L49:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v453+v437&int32(_a_F_g_cube_picksplit_1)<<(uint(int32(4))%32))))
	v606 = F_pg_detoast_datum(m, v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L62
	}
L50:
	;
	if v460 == int32(0) {
		v598 = v451
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	if v464 <= int32(0) {
		v598 = v451
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v467 = int32(1)
	v470 = v460 + int32(8)
	if v464 == v467 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v464&v467 == int32(0) {
		v598 = v559
		goto L49
	} else {
		goto L60
	}
L54:
	;
	v536 = int32(0)
	v559 = float64(1)
	goto L53
L55:
	;
	goto L56
L56:
	;
	v478 = int32(0)
	v480 = v478
	v483 = v478
	v503 = float64(1)
	goto L57
L57:
	;
	v506 = int32(3)
	v509 = *(*float64)(unsafe.Add(mBase, uint32(v470+(v480+v464)<<(uint(v506)%32))))
	v513 = *(*float64)(unsafe.Add(mBase, uint32(v470+v480<<(uint(v506)%32))))
	v518 = v480 | int32(1)
	v523 = *(*float64)(unsafe.Add(mBase, uint32(v470+(v518+v464)<<(uint(v506)%32))))
	v527 = *(*float64)(unsafe.Add(mBase, uint32(v470+v518<<(uint(v506)%32))))
	v530 = base.F64_mul(base.F64_mul(v503, base.F64_abs(base.F64_sub(v509, v513))), base.F64_abs(base.F64_sub(v523, v527)))
	v531 = int32(2)
	v532 = v480 + v531
	v534 = v483 + v531
	if v534 != v464&int32(2147483646) {
		v480 = v532
		v483 = v534
		v503 = v530
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v536 = v532
	v559 = v530
	goto L53
L59:
	;
	goto L58
L60:
	;
	v564 = int32(3)
	v567 = *(*float64)(unsafe.Add(mBase, uint32(v470+(v536+v464)<<(uint(v564)%32))))
	v571 = *(*float64)(unsafe.Add(mBase, uint32(v470+v536<<(uint(v564)%32))))
	v598 = base.F64_mul(v559, base.F64_abs(base.F64_sub(v567, v571)))
	goto L49
L61:
	;
	if v29&int32(_a_F_g_cube_picksplit_1) != int32(1) {
		goto L73
	} else {
		goto L74
	}
L62:
	;
	if v606 == int32(0) {
		v745 = v22
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v606)+4))
	if v610 <= int32(0) {
		v745 = v22
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v613 = int32(1)
	v616 = v606 + int32(8)
	if v610 == v613 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v610&v613 == int32(0) {
		v745 = v706
		goto L61
	} else {
		goto L72
	}
L66:
	;
	v682 = int32(0)
	v706 = float64(1)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v624 = int32(0)
	v626 = v624
	v629 = v624
	v650 = float64(1)
	goto L69
L69:
	;
	v652 = int32(3)
	v655 = *(*float64)(unsafe.Add(mBase, uint32(v616+(v626+v610)<<(uint(v652)%32))))
	v659 = *(*float64)(unsafe.Add(mBase, uint32(v616+v626<<(uint(v652)%32))))
	v664 = v626 | int32(1)
	v669 = *(*float64)(unsafe.Add(mBase, uint32(v616+(v664+v610)<<(uint(v652)%32))))
	v673 = *(*float64)(unsafe.Add(mBase, uint32(v616+v664<<(uint(v652)%32))))
	v676 = base.F64_mul(base.F64_mul(v650, base.F64_abs(base.F64_sub(v655, v659))), base.F64_abs(base.F64_sub(v669, v673)))
	v677 = int32(2)
	v678 = v626 + v677
	v680 = v629 + v677
	if v680 != v610&int32(2147483646) {
		v626 = v678
		v629 = v680
		v650 = v676
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v682 = v678
	v706 = v676
	goto L65
L71:
	;
	goto L70
L72:
	;
	v710 = int32(3)
	v713 = *(*float64)(unsafe.Add(mBase, uint32(v616+(v682+v610)<<(uint(v710)%32))))
	v717 = *(*float64)(unsafe.Add(mBase, uint32(v616+v682<<(uint(v710)%32))))
	v745 = base.F64_mul(v706, base.F64_abs(base.F64_sub(v713, v717)))
	goto L61
L73:
	;
	v750 = int32(1)
	v757 = v750
	v759 = v750
	v761 = v426
	v765 = v450
	v769 = v460
	v770 = v606
	v779 = v598
	v780 = v745
	goto L76
L74:
	;
	v1143 = v426
	v1147 = v450
	v1151 = v460
	v1152 = v606
	goto L75
L75:
	;
	v1163 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1143))) = uint16(v1163)
	*(*uint16)(unsafe.Add(mBase, uint32(v1147))) = uint16(v1163)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v1152
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v1151
	return v27
L76:
	;
	v781 = int32(_a_F_g_cube_picksplit_1)
	v782 = v759 & v781
	if v782 == v436&v781 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v1143 = v1113
	v1147 = v1117
	v1151 = v1121
	v1152 = v1122
	goto L75
L78:
	;
	v1134 = v759 + int32(1)
	v1136 = v1134 & int32(_a_F_g_cube_picksplit_1)
	if base.Ui32(v1136) <= base.Ui32((v29-v750)&int32(_a_F_g_cube_picksplit_1)) {
		v757 = v1136
		v759 = v1134
		v761 = v1113
		v765 = v1117
		v769 = v1121
		v770 = v1122
		v779 = v1131
		v780 = v1132
		goto L76
	} else {
		goto L113
	}
L79:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v765))) = uint16(v436)
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v787 + int32(1)
	v1113 = v761
	v1117 = v765 + int32(2)
	v1121 = v769
	v1122 = v770
	v1131 = v779
	v1132 = v780
	goto L78
L80:
	;
	goto L81
L81:
	;
	if v437&int32(_a_F_g_cube_picksplit_1) == v782 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v761))) = uint16(v437)
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v797 + int32(1)
	v1113 = v761 + int32(2)
	v1117 = v765
	v1121 = v769
	v1122 = v770
	v1131 = v779
	v1132 = v780
	goto L78
L83:
	;
	goto L84
L84:
	;
	v803 = float64(0)
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v453+v757<<(uint(int32(4))%32))))
	v808 = F_pg_detoast_datum(m, v807)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v810 = F_cube_union_v0(m, v769, v808)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v812 = F_cube_union_v0(m, v770, v808)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v814 = float64(0)
	if v810 == int32(0) {
		v949 = v814
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v812 == int32(0) {
		v1088 = v803
		goto L99
	} else {
		goto L100
	}
L89:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v810)+4))
	if v817 <= int32(0) {
		v949 = v814
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v820 = int32(1)
	v823 = v810 + int32(8)
	if v817 == v820 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v817&v820 == int32(0) {
		v949 = v910
		goto L88
	} else {
		goto L98
	}
L92:
	;
	v889 = int32(0)
	v910 = float64(1)
	goto L91
L93:
	;
	goto L94
L94:
	;
	v831 = int32(0)
	v833 = v831
	v839 = v831
	v854 = float64(1)
	goto L95
L95:
	;
	v859 = int32(3)
	v862 = *(*float64)(unsafe.Add(mBase, uint32(v823+(v833+v817)<<(uint(v859)%32))))
	v866 = *(*float64)(unsafe.Add(mBase, uint32(v823+v833<<(uint(v859)%32))))
	v871 = v833 | int32(1)
	v876 = *(*float64)(unsafe.Add(mBase, uint32(v823+(v871+v817)<<(uint(v859)%32))))
	v880 = *(*float64)(unsafe.Add(mBase, uint32(v823+v871<<(uint(v859)%32))))
	v883 = base.F64_mul(base.F64_mul(v854, base.F64_abs(base.F64_sub(v862, v866))), base.F64_abs(base.F64_sub(v876, v880)))
	v884 = int32(2)
	v885 = v833 + v884
	v887 = v839 + v884
	if v887 != v817&int32(2147483646) {
		v833 = v885
		v839 = v887
		v854 = v883
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v889 = v885
	v910 = v883
	goto L91
L97:
	;
	goto L96
L98:
	;
	v917 = int32(3)
	v920 = *(*float64)(unsafe.Add(mBase, uint32(v823+(v889+v817)<<(uint(v917)%32))))
	v924 = *(*float64)(unsafe.Add(mBase, uint32(v823+v889<<(uint(v917)%32))))
	v949 = base.F64_mul(v910, base.F64_abs(base.F64_sub(v920, v924)))
	goto L88
L99:
	;
	if base.F64_lt(base.F64_sub(v949, v779), base.F64_sub(v1088, v780)) != 0 {
		goto L110
	} else {
		goto L111
	}
L100:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v812)+4))
	if v955 <= int32(0) {
		v1088 = v803
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v958 = int32(1)
	v961 = v812 + int32(8)
	if v955 == v958 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v955&v958 == int32(0) {
		v1088 = v1049
		goto L99
	} else {
		goto L109
	}
L103:
	;
	v1027 = int32(0)
	v1049 = float64(1)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v969 = int32(0)
	v971 = v969
	v977 = v969
	v993 = float64(1)
	goto L106
L106:
	;
	v997 = int32(3)
	v1000 = *(*float64)(unsafe.Add(mBase, uint32(v961+(v971+v955)<<(uint(v997)%32))))
	v1004 = *(*float64)(unsafe.Add(mBase, uint32(v961+v971<<(uint(v997)%32))))
	v1009 = v971 | int32(1)
	v1014 = *(*float64)(unsafe.Add(mBase, uint32(v961+(v1009+v955)<<(uint(v997)%32))))
	v1018 = *(*float64)(unsafe.Add(mBase, uint32(v961+v1009<<(uint(v997)%32))))
	v1021 = base.F64_mul(base.F64_mul(v993, base.F64_abs(base.F64_sub(v1000, v1004))), base.F64_abs(base.F64_sub(v1014, v1018)))
	v1022 = int32(2)
	v1023 = v971 + v1022
	v1025 = v977 + v1022
	if v1025 != v955&int32(2147483646) {
		v971 = v1023
		v977 = v1025
		v993 = v1021
		goto L106
	} else {
		goto L108
	}
L107:
	;
	v1027 = v1023
	v1049 = v1021
	goto L102
L108:
	;
	goto L107
L109:
	;
	v1055 = int32(3)
	v1058 = *(*float64)(unsafe.Add(mBase, uint32(v961+(v1027+v955)<<(uint(v1055)%32))))
	v1062 = *(*float64)(unsafe.Add(mBase, uint32(v961+v1027<<(uint(v1055)%32))))
	v1088 = base.F64_mul(v1049, base.F64_abs(base.F64_sub(v1058, v1062)))
	goto L99
L110:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v765))) = uint16(v759)
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v1095 + int32(1)
	v1113 = v761
	v1117 = v765 + int32(2)
	v1121 = v810
	v1122 = v770
	v1131 = v949
	v1132 = v780
	goto L78
L111:
	;
	goto L112
L112:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v761))) = uint16(v759)
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v1102 + int32(1)
	v1113 = v761 + int32(2)
	v1117 = v765
	v1121 = v769
	v1122 = v812
	v1131 = v779
	v1132 = v1088
	goto L78
L113:
	;
	goto L77
}
func F_g_cube_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v14 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(base.Ui32(v13) >> (uint(v14) % 32))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v14 <= v18 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = int32(1)
	v26 = v9
	goto L6
L4:
	;
	v50 = int32(0)
	goto L5
L5:
	;
	return v50
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(4)+v24<<(uint(int32(4))%32))))
	v33 = F_pg_detoast_datum(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v50 = v35
	goto L5
L8:
	;
	v35 = F_cube_union_v0(m, v26, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(base.Ui32(v37) >> (uint(int32(2)) % 32))
	v42 = v24 + int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v42 < v43 {
		v24 = v42
		v26 = v35
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
}
