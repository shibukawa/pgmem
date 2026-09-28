package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_cube_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v23 float64
	_ = v23
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 float64
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 float64
	_ = v115
	var v118 float64
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v168 float64
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 float64
	_ = v179
	var v180 float64
	_ = v180
	var v187 float64
	_ = v187
	var v188 float64
	_ = v188
	var v191 float64
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v221 float64
	_ = v221
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 float64
	_ = v232
	var v233 float64
	_ = v233
	var v259 float64
	_ = v259
	var v266 int64
	_ = v266
	var v267 int64
	_ = v267
	var v268 int64
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v314 float64
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 float64
	_ = v324
	var v325 float64
	_ = v325
	var v332 float64
	_ = v332
	var v333 float64
	_ = v333
	var v336 float64
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v367 float64
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 float64
	_ = v377
	var v378 float64
	_ = v378
	var v405 float64
	_ = v405
	var v409 float64
	_ = v409
	var v411 int32
	_ = v411
	var v412 float64
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 float64
	_ = v455
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v508 float64
	_ = v508
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 float64
	_ = v517
	var v518 float64
	_ = v518
	var v525 float64
	_ = v525
	var v526 float64
	_ = v526
	var v529 float64
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v561 float64
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v570 float64
	_ = v570
	var v571 float64
	_ = v571
	var v599 float64
	_ = v599
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v653 float64
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 float64
	_ = v661
	var v662 float64
	_ = v662
	var v669 float64
	_ = v669
	var v670 float64
	_ = v670
	var v673 float64
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v706 float64
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v714 float64
	_ = v714
	var v715 float64
	_ = v715
	var v744 float64
	_ = v744
	var v750 int32
	_ = v750
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v780 float64
	_ = v780
	var v781 float64
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v799 int32
	_ = v799
	var v805 float64
	_ = v805
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
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 float64
	_ = v816
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v850 int32
	_ = v850
	var v857 float64
	_ = v857
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v868 float64
	_ = v868
	var v869 float64
	_ = v869
	var v876 float64
	_ = v876
	var v877 float64
	_ = v877
	var v880 float64
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v910 float64
	_ = v910
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v921 float64
	_ = v921
	var v922 float64
	_ = v922
	var v948 float64
	_ = v948
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v986 int32
	_ = v986
	var v994 float64
	_ = v994
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1004 float64
	_ = v1004
	var v1005 float64
	_ = v1005
	var v1012 float64
	_ = v1012
	var v1013 float64
	_ = v1013
	var v1016 float64
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1047 float64
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1057 float64
	_ = v1057
	var v1058 float64
	_ = v1058
	var v1085 float64
	_ = v1085
	var v1093 int32
	_ = v1093
	var v1100 int32
	_ = v1100
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1130 float64
	_ = v1130
	var v1131 float64
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1165 int32
	_ = v1165
	v23 = float64(0)
	v28 = int32(1)
	v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v30 = base.I32_wrap_i64(v29)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v36 = (v32 + int32(_a_F_g_cube_picksplit_0)) & int32(_a_F_g_cube_picksplit_1)
	v40 = v36<<(uint(v28)%32) + int32(4)
	v41 = F_palloc(m, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v41
	v46 = F_palloc(m, v40)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v46
	v49 = int32(2)
	if base.Ui32(v49) <= base.Ui32(v36) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v53 = v31 + int32(8)
	v54 = int32(1)
	v64 = v54
	v67 = v54
	v74 = v28
	v75 = v49
	v80 = v23
	goto L7
L5:
	;
	v429 = v46
	v441 = v28
	v442 = v49
	goto L6
L6:
	;
	v450 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v450
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v455 = float64(0)
	v457 = v31 + int32(8)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v457+v441&int32(_a_F_g_cube_picksplit_1)*int32(24))))
	v464 = F_pg_detoast_datum(m, v463)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L50
	}
L7:
	;
	v85 = v53 + v64*int32(24)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v87 = F_pg_detoast_datum(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v429 = v422
	v441 = v414
	v442 = v413
	goto L6
L9:
	;
	v90 = v64 + int32(1)
	v91 = v90
	v92 = v90
	v102 = v67
	v109 = v74
	v110 = v75
	v115 = v80
	goto L10
L10:
	;
	v118 = float64(0)
	v122 = v53 + v91*int32(24)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v124 = F_pg_detoast_datum(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	if v90 != v36 {
		v64 = v90
		v67 = v415
		v74 = v414
		v75 = v413
		v80 = v412
		goto L7
	} else {
		goto L48
	}
L12:
	;
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v122)))
	v268 = F_DirectFunctionCall2Coll(m, int32(_a_F_g_cube_picksplit_2), int32(0), v266, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L26
	}
L13:
	;
	v126 = F_cube_union_v0(m, v87, v124)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v126 == int32(0) {
		v259 = v118
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v130 <= int32(0) {
		v259 = v118
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v134 = v126 + int32(8)
	if v130 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v226 = int32(3)
	v228 = v134 + v200<<(uint(v226)%32)
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v228+v130<<(uint(v226)%32))))
	v233 = *(*float64)(unsafe.Add(mBase, uint32(v228)))
	v259 = base.F64_mul(v221, base.F64_abs(base.F64_sub(v232, v233)))
	goto L12
L18:
	;
	v200 = int32(0)
	v221 = float64(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v144 = int32(0)
	v147 = v144
	v149 = v144
	v168 = float64(1)
	goto L21
L21:
	;
	v173 = int32(3)
	v175 = v134 + v147<<(uint(v173)%32)
	v177 = v130 << (uint(v173) % 32)
	v179 = *(*float64)(unsafe.Add(mBase, uint32(v175+v177)))
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v175)))
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v175+int32(8)+v177)))
	v188 = *(*float64)(unsafe.Add(mBase, uint32(v175)+8))
	v191 = base.F64_mul(base.F64_mul(v168, base.F64_abs(base.F64_sub(v179, v180))), base.F64_abs(base.F64_sub(v187, v188)))
	v192 = int32(2)
	v193 = v147 + v192
	v195 = v149 + v192
	if v195 != v130&int32(2147483646) {
		v147 = v193
		v149 = v195
		v168 = v191
		goto L21
	} else {
		goto L23
	}
L22:
	;
	if v130&int32(1) == int32(0) {
		v259 = v191
		goto L12
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v200 = v193
	v221 = v191
	goto L17
L25:
	;
	v409 = base.F64_sub(v259, v405)
	v411 = v102 | base.F64_gt(v409, v115)
	if v411 != 0 {
		goto L38
	} else {
		goto L39
	}
L26:
	;
	v271 = F_pg_detoast_datum(m, base.I32_wrap_i64(v268))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v271 == int32(0) {
		v405 = v118
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	if v275 <= int32(0) {
		v405 = v118
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v279 = v271 + int32(8)
	if v275 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v371 = int32(3)
	v373 = v279 + v345<<(uint(v371)%32)
	v377 = *(*float64)(unsafe.Add(mBase, uint32(v373+v275<<(uint(v371)%32))))
	v378 = *(*float64)(unsafe.Add(mBase, uint32(v373)))
	v405 = base.F64_mul(v367, base.F64_abs(base.F64_sub(v377, v378)))
	goto L25
L31:
	;
	v345 = int32(0)
	v367 = float64(1)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v289 = int32(0)
	v292 = v289
	v294 = v289
	v314 = float64(1)
	goto L34
L34:
	;
	v318 = int32(3)
	v320 = v279 + v292<<(uint(v318)%32)
	v322 = v275 << (uint(v318) % 32)
	v324 = *(*float64)(unsafe.Add(mBase, uint32(v320+v322)))
	v325 = *(*float64)(unsafe.Add(mBase, uint32(v320)))
	v332 = *(*float64)(unsafe.Add(mBase, uint32(v320+int32(8)+v322)))
	v333 = *(*float64)(unsafe.Add(mBase, uint32(v320)+8))
	v336 = base.F64_mul(base.F64_mul(v314, base.F64_abs(base.F64_sub(v324, v325))), base.F64_abs(base.F64_sub(v332, v333)))
	v337 = int32(2)
	v338 = v292 + v337
	v340 = v294 + v337
	if v340 != v275&int32(2147483646) {
		v292 = v338
		v294 = v340
		v314 = v336
		goto L34
	} else {
		goto L36
	}
L35:
	;
	if v275&int32(1) == int32(0) {
		v405 = v336
		goto L25
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	v345 = v338
	v367 = v336
	goto L30
L38:
	;
	v412 = v409
	goto L40
L39:
	;
	v412 = v115
	goto L40
L40:
	;
	if v411 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v413 = v92
	goto L43
L42:
	;
	v413 = v110
	goto L43
L43:
	;
	if v411 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v414 = v64
	goto L46
L45:
	;
	v414 = v109
	goto L46
L46:
	;
	v415 = int32(0)
	v417 = v92 + int32(1)
	v419 = v417 & int32(_a_F_g_cube_picksplit_1)
	if base.Ui32(v419) <= base.Ui32(v36) {
		v91 = v419
		v92 = v417
		v102 = v415
		v109 = v414
		v110 = v413
		v115 = v412
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
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v457+v442&int32(_a_F_g_cube_picksplit_1)*int32(24))))
	v608 = F_pg_detoast_datum(m, v607)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L62
	}
L50:
	;
	if v464 == int32(0) {
		v599 = v455
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	if v468 <= int32(0) {
		v599 = v455
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v472 = v464 + int32(8)
	if v468 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v564 = int32(3)
	v566 = v472 + v538<<(uint(v564)%32)
	v570 = *(*float64)(unsafe.Add(mBase, uint32(v566+v468<<(uint(v564)%32))))
	v571 = *(*float64)(unsafe.Add(mBase, uint32(v566)))
	v599 = base.F64_mul(v561, base.F64_abs(base.F64_sub(v570, v571)))
	goto L49
L54:
	;
	v538 = int32(0)
	v561 = float64(1)
	goto L53
L55:
	;
	goto L56
L56:
	;
	v482 = int32(0)
	v485 = v482
	v487 = v482
	v508 = float64(1)
	goto L57
L57:
	;
	v511 = int32(3)
	v513 = v472 + v485<<(uint(v511)%32)
	v515 = v468 << (uint(v511) % 32)
	v517 = *(*float64)(unsafe.Add(mBase, uint32(v513+v515)))
	v518 = *(*float64)(unsafe.Add(mBase, uint32(v513)))
	v525 = *(*float64)(unsafe.Add(mBase, uint32(v513+int32(8)+v515)))
	v526 = *(*float64)(unsafe.Add(mBase, uint32(v513)+8))
	v529 = base.F64_mul(base.F64_mul(v508, base.F64_abs(base.F64_sub(v517, v518))), base.F64_abs(base.F64_sub(v525, v526)))
	v530 = int32(2)
	v531 = v485 + v530
	v533 = v487 + v530
	if v533 != v468&int32(2147483646) {
		v485 = v531
		v487 = v533
		v508 = v529
		goto L57
	} else {
		goto L59
	}
L58:
	;
	if v468&int32(1) == int32(0) {
		v599 = v529
		goto L49
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	v538 = v531
	v561 = v529
	goto L53
L61:
	;
	if v32&int32(_a_F_g_cube_picksplit_1) != int32(1) {
		goto L73
	} else {
		goto L74
	}
L62:
	;
	if v608 == int32(0) {
		v744 = v23
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	if v612 <= int32(0) {
		v744 = v23
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v616 = v608 + int32(8)
	if v612 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v708 = int32(3)
	v710 = v616 + v682<<(uint(v708)%32)
	v714 = *(*float64)(unsafe.Add(mBase, uint32(v710+v612<<(uint(v708)%32))))
	v715 = *(*float64)(unsafe.Add(mBase, uint32(v710)))
	v744 = base.F64_mul(v706, base.F64_abs(base.F64_sub(v714, v715)))
	goto L61
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
	v626 = int32(0)
	v629 = v626
	v631 = v626
	v653 = float64(1)
	goto L69
L69:
	;
	v655 = int32(3)
	v657 = v616 + v629<<(uint(v655)%32)
	v659 = v612 << (uint(v655) % 32)
	v661 = *(*float64)(unsafe.Add(mBase, uint32(v657+v659)))
	v662 = *(*float64)(unsafe.Add(mBase, uint32(v657)))
	v669 = *(*float64)(unsafe.Add(mBase, uint32(v657+int32(8)+v659)))
	v670 = *(*float64)(unsafe.Add(mBase, uint32(v657)+8))
	v673 = base.F64_mul(base.F64_mul(v653, base.F64_abs(base.F64_sub(v661, v662))), base.F64_abs(base.F64_sub(v669, v670)))
	v674 = int32(2)
	v675 = v629 + v674
	v677 = v631 + v674
	if v677 != v612&int32(2147483646) {
		v629 = v675
		v631 = v677
		v653 = v673
		goto L69
	} else {
		goto L71
	}
L70:
	;
	if v612&int32(1) == int32(0) {
		v744 = v673
		goto L61
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	v682 = v675
	v706 = v673
	goto L65
L73:
	;
	v750 = int32(1)
	v757 = v750
	v759 = v750
	v760 = v454
	v762 = v429
	v768 = v608
	v772 = v464
	v780 = v599
	v781 = v744
	goto L76
L74:
	;
	v1142 = v454
	v1144 = v429
	v1150 = v608
	v1154 = v464
	goto L75
L75:
	;
	v1165 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1144))) = uint16(v1165)
	*(*uint16)(unsafe.Add(mBase, uint32(v1142))) = uint16(v1165)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = base.I64_extend_i32_u(v1150)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = base.I64_extend_i32_u(v1154)
	return v29 & int64(4294967295)
L76:
	;
	v783 = int32(_a_F_g_cube_picksplit_1)
	v784 = v759 & v783
	if v784 == v441&v783 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v1142 = v1110
	v1144 = v1112
	v1150 = v1118
	v1154 = v1122
	goto L75
L78:
	;
	v1134 = v759 + int32(1)
	v1136 = v1134 & int32(_a_F_g_cube_picksplit_1)
	if base.Ui32(v1136) <= base.Ui32((v32-v750)&int32(_a_F_g_cube_picksplit_1)) {
		v757 = v1136
		v759 = v1134
		v760 = v1110
		v762 = v1112
		v768 = v1118
		v772 = v1122
		v780 = v1130
		v781 = v1131
		goto L76
	} else {
		goto L113
	}
L79:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v760))) = uint16(v441)
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v789 + int32(1)
	v1110 = v760 + int32(2)
	v1112 = v762
	v1118 = v768
	v1122 = v772
	v1130 = v780
	v1131 = v781
	goto L78
L80:
	;
	goto L81
L81:
	;
	if v442&int32(_a_F_g_cube_picksplit_1) == v784 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v762))) = uint16(v442)
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v799 + int32(1)
	v1110 = v760
	v1112 = v762 + int32(2)
	v1118 = v768
	v1122 = v772
	v1130 = v780
	v1131 = v781
	goto L78
L83:
	;
	goto L84
L84:
	;
	v805 = float64(0)
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v457+v757*int32(24))))
	v810 = F_pg_detoast_datum(m, v809)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v812 = F_cube_union_v0(m, v772, v810)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v814 = F_cube_union_v0(m, v768, v810)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v816 = float64(0)
	if v812 == int32(0) {
		v948 = v816
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v814 == int32(0) {
		v1085 = v805
		goto L99
	} else {
		goto L100
	}
L89:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v812)+4))
	if v819 <= int32(0) {
		v948 = v816
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v823 = v812 + int32(8)
	if v819 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v915 = int32(3)
	v917 = v823 + v889<<(uint(v915)%32)
	v921 = *(*float64)(unsafe.Add(mBase, uint32(v917+v819<<(uint(v915)%32))))
	v922 = *(*float64)(unsafe.Add(mBase, uint32(v917)))
	v948 = base.F64_mul(v910, base.F64_abs(base.F64_sub(v921, v922)))
	goto L88
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
	v833 = int32(0)
	v836 = v833
	v850 = v833
	v857 = float64(1)
	goto L95
L95:
	;
	v862 = int32(3)
	v864 = v823 + v836<<(uint(v862)%32)
	v866 = v819 << (uint(v862) % 32)
	v868 = *(*float64)(unsafe.Add(mBase, uint32(v864+v866)))
	v869 = *(*float64)(unsafe.Add(mBase, uint32(v864)))
	v876 = *(*float64)(unsafe.Add(mBase, uint32(v864+int32(8)+v866)))
	v877 = *(*float64)(unsafe.Add(mBase, uint32(v864)+8))
	v880 = base.F64_mul(base.F64_mul(v857, base.F64_abs(base.F64_sub(v868, v869))), base.F64_abs(base.F64_sub(v876, v877)))
	v881 = int32(2)
	v882 = v836 + v881
	v884 = v850 + v881
	if v884 != v819&int32(2147483646) {
		v836 = v882
		v850 = v884
		v857 = v880
		goto L95
	} else {
		goto L97
	}
L96:
	;
	if v819&int32(1) == int32(0) {
		v948 = v880
		goto L88
	} else {
		goto L98
	}
L97:
	;
	goto L96
L98:
	;
	v889 = v882
	v910 = v880
	goto L91
L99:
	;
	if base.F64_lt(base.F64_sub(v948, v780), base.F64_sub(v1085, v781)) != 0 {
		goto L110
	} else {
		goto L111
	}
L100:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v814)+4))
	if v955 <= int32(0) {
		v1085 = v805
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v959 = v814 + int32(8)
	if v955 == int32(1) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v1051 = int32(3)
	v1053 = v959 + v1025<<(uint(v1051)%32)
	v1057 = *(*float64)(unsafe.Add(mBase, uint32(v1053+v955<<(uint(v1051)%32))))
	v1058 = *(*float64)(unsafe.Add(mBase, uint32(v1053)))
	v1085 = base.F64_mul(v1047, base.F64_abs(base.F64_sub(v1057, v1058)))
	goto L99
L103:
	;
	v1025 = int32(0)
	v1047 = float64(1)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v969 = int32(0)
	v972 = v969
	v986 = v969
	v994 = float64(1)
	goto L106
L106:
	;
	v998 = int32(3)
	v1000 = v959 + v972<<(uint(v998)%32)
	v1002 = v955 << (uint(v998) % 32)
	v1004 = *(*float64)(unsafe.Add(mBase, uint32(v1000+v1002)))
	v1005 = *(*float64)(unsafe.Add(mBase, uint32(v1000)))
	v1012 = *(*float64)(unsafe.Add(mBase, uint32(v1000+int32(8)+v1002)))
	v1013 = *(*float64)(unsafe.Add(mBase, uint32(v1000)+8))
	v1016 = base.F64_mul(base.F64_mul(v994, base.F64_abs(base.F64_sub(v1004, v1005))), base.F64_abs(base.F64_sub(v1012, v1013)))
	v1017 = int32(2)
	v1018 = v972 + v1017
	v1020 = v986 + v1017
	if v1020 != v955&int32(2147483646) {
		v972 = v1018
		v986 = v1020
		v994 = v1016
		goto L106
	} else {
		goto L108
	}
L107:
	;
	if v955&int32(1) == int32(0) {
		v1085 = v1016
		goto L99
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	v1025 = v1018
	v1047 = v1016
	goto L102
L110:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v760))) = uint16(v759)
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v1093 + int32(1)
	v1110 = v760 + int32(2)
	v1112 = v762
	v1118 = v768
	v1122 = v812
	v1130 = v948
	v1131 = v781
	goto L78
L111:
	;
	goto L112
L112:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v762))) = uint16(v759)
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v1100 + int32(1)
	v1110 = v760
	v1112 = v762 + int32(2)
	v1118 = v814
	v1122 = v772
	v1130 = v780
	v1131 = v1085
	goto L78
L113:
	;
	goto L77
}
func F_g_cube_union(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v51 int64
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v14 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(base.Ui32(v13) >> (uint(v14) % 32))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v14 <= v17 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = int32(1)
	v25 = v9
	goto L6
L4:
	;
	v51 = int64(0)
	goto L5
L5:
	;
	return v51
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(8)+v23*int32(24))))
	v32 = F_pg_detoast_datum(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v51 = base.I64_extend_i32_u(v34)
	goto L5
L8:
	;
	v34 = F_cube_union_v0(m, v25, v32)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(base.Ui32(v36) >> (uint(int32(2)) % 32))
	v41 = v23 + int32(1)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v41 < v42 {
		v23 = v41
		v25 = v34
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
}
