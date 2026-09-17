package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_cube_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 float64
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 float64
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 float64
	_ = v112
	var v114 float64
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v165 float64
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 float64
	_ = v174
	var v175 float64
	_ = v175
	var v182 float64
	_ = v182
	var v183 float64
	_ = v183
	var v186 float64
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v217 float64
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 float64
	_ = v226
	var v227 float64
	_ = v227
	var v254 float64
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v305 float64
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 float64
	_ = v315
	var v316 float64
	_ = v316
	var v323 float64
	_ = v323
	var v324 float64
	_ = v324
	var v327 float64
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v357 float64
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 float64
	_ = v367
	var v368 float64
	_ = v368
	var v394 float64
	_ = v394
	var v398 float64
	_ = v398
	var v400 int32
	_ = v400
	var v401 float64
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 float64
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v496 float64
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 float64
	_ = v504
	var v505 float64
	_ = v505
	var v512 float64
	_ = v512
	var v513 float64
	_ = v513
	var v516 float64
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v548 float64
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v556 float64
	_ = v556
	var v557 float64
	_ = v557
	var v585 float64
	_ = v585
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v638 float64
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v645 float64
	_ = v645
	var v646 float64
	_ = v646
	var v653 float64
	_ = v653
	var v654 float64
	_ = v654
	var v657 float64
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v690 float64
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v697 float64
	_ = v697
	var v698 float64
	_ = v698
	var v727 float64
	_ = v727
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v762 float64
	_ = v762
	var v763 float64
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v780 int32
	_ = v780
	var v786 float64
	_ = v786
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 float64
	_ = v797
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v828 int32
	_ = v828
	var v839 float64
	_ = v839
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v848 float64
	_ = v848
	var v849 float64
	_ = v849
	var v856 float64
	_ = v856
	var v857 float64
	_ = v857
	var v860 float64
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v891 float64
	_ = v891
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v900 float64
	_ = v900
	var v901 float64
	_ = v901
	var v928 float64
	_ = v928
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v961 int32
	_ = v961
	var v971 float64
	_ = v971
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v981 float64
	_ = v981
	var v982 float64
	_ = v982
	var v989 float64
	_ = v989
	var v990 float64
	_ = v990
	var v993 float64
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1023 float64
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1033 float64
	_ = v1033
	var v1034 float64
	_ = v1034
	var v1060 float64
	_ = v1060
	var v1068 int32
	_ = v1068
	var v1075 int32
	_ = v1075
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1092 int32
	_ = v1092
	var v1105 float64
	_ = v1105
	var v1106 float64
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1123 int32
	_ = v1123
	var v1138 int32
	_ = v1138
	v23 = float64(0)
	v27 = int32(1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v34 = (v30 + int32(_a_F_g_cube_picksplit_0)) & int32(_a_F_g_cube_picksplit_1)
	v38 = v34<<(uint(v27)%32) + int32(4)
	v39 = F_palloc(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v39
	v44 = F_palloc(m, v38)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v44
	v47 = int32(2)
	if base.Ui32(v47) <= base.Ui32(v34) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v51 = v29 + int32(4)
	v52 = int32(1)
	v59 = v52
	v69 = v27
	v70 = v47
	v71 = v52
	v78 = v23
	goto L7
L5:
	;
	v416 = v44
	v427 = v27
	v428 = v47
	goto L6
L6:
	;
	v438 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v438
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v443 = float64(0)
	v444 = int32(4)
	v445 = v29 + v444
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v445+v427&int32(_a_F_g_cube_picksplit_1)<<(uint(v444)%32))))
	v452 = F_pg_detoast_datum(m, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L50
	}
L7:
	;
	v82 = v51 + v59<<(uint(int32(4))%32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = F_pg_detoast_datum(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v416 = v411
	v427 = v403
	v428 = v402
	goto L6
L9:
	;
	v87 = v59 + int32(1)
	v88 = v87
	v89 = v87
	v103 = v69
	v104 = v70
	v105 = v71
	v112 = v78
	goto L10
L10:
	;
	v114 = float64(0)
	v118 = v51 + v88<<(uint(int32(4))%32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v120 = F_pg_detoast_datum(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	if v87 != v34 {
		v59 = v87
		v69 = v403
		v70 = v402
		v71 = v404
		v78 = v401
		goto L7
	} else {
		goto L48
	}
L12:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v261 = F_DirectFunctionCall2Coll(m, int32(_a_F_g_cube_picksplit_2), int32(0), v259, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L26
	}
L13:
	;
	v122 = F_cube_union_v0(m, v84, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v122 == int32(0) {
		v254 = v114
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v126 <= int32(0) {
		v254 = v114
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v130 = v122 + int32(8)
	if v126 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v220 = int32(3)
	v222 = v130 + v195<<(uint(v220)%32)
	v226 = *(*float64)(unsafe.Add(mBase, uint32(v222+v126<<(uint(v220)%32))))
	v227 = *(*float64)(unsafe.Add(mBase, uint32(v222)))
	v254 = base.F64_mul(v217, base.F64_abs(base.F64_sub(v226, v227)))
	goto L12
L18:
	;
	v195 = int32(0)
	v217 = float64(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v140 = int32(0)
	v143 = v140
	v148 = v140
	v165 = float64(1)
	goto L21
L21:
	;
	v168 = int32(3)
	v170 = v130 + v143<<(uint(v168)%32)
	v172 = v126 << (uint(v168) % 32)
	v174 = *(*float64)(unsafe.Add(mBase, uint32(v170+v172)))
	v175 = *(*float64)(unsafe.Add(mBase, uint32(v170)))
	v182 = *(*float64)(unsafe.Add(mBase, uint32(v170+int32(8)+v172)))
	v183 = *(*float64)(unsafe.Add(mBase, uint32(v170)+8))
	v186 = base.F64_mul(base.F64_mul(v165, base.F64_abs(base.F64_sub(v174, v175))), base.F64_abs(base.F64_sub(v182, v183)))
	v187 = int32(2)
	v188 = v143 + v187
	v190 = v148 + v187
	if v190 != v126&int32(2147483646) {
		v143 = v188
		v148 = v190
		v165 = v186
		goto L21
	} else {
		goto L23
	}
L22:
	;
	if v126&int32(1) == int32(0) {
		v254 = v186
		goto L12
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v195 = v188
	v217 = v186
	goto L17
L25:
	;
	v398 = base.F64_sub(v254, v394)
	v400 = v105 | base.F64_gt(v398, v112)
	if v400 != 0 {
		goto L38
	} else {
		goto L39
	}
L26:
	;
	v263 = F_pg_detoast_datum(m, v261)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v263 == int32(0) {
		v394 = v114
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v267 <= int32(0) {
		v394 = v114
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v271 = v263 + int32(8)
	if v267 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v361 = int32(3)
	v363 = v271 + v336<<(uint(v361)%32)
	v367 = *(*float64)(unsafe.Add(mBase, uint32(v363+v267<<(uint(v361)%32))))
	v368 = *(*float64)(unsafe.Add(mBase, uint32(v363)))
	v394 = base.F64_mul(v357, base.F64_abs(base.F64_sub(v367, v368)))
	goto L25
L31:
	;
	v336 = int32(0)
	v357 = float64(1)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v281 = int32(0)
	v284 = v281
	v289 = v281
	v305 = float64(1)
	goto L34
L34:
	;
	v309 = int32(3)
	v311 = v271 + v284<<(uint(v309)%32)
	v313 = v267 << (uint(v309) % 32)
	v315 = *(*float64)(unsafe.Add(mBase, uint32(v311+v313)))
	v316 = *(*float64)(unsafe.Add(mBase, uint32(v311)))
	v323 = *(*float64)(unsafe.Add(mBase, uint32(v311+int32(8)+v313)))
	v324 = *(*float64)(unsafe.Add(mBase, uint32(v311)+8))
	v327 = base.F64_mul(base.F64_mul(v305, base.F64_abs(base.F64_sub(v315, v316))), base.F64_abs(base.F64_sub(v323, v324)))
	v328 = int32(2)
	v329 = v284 + v328
	v331 = v289 + v328
	if v331 != v267&int32(2147483646) {
		v284 = v329
		v289 = v331
		v305 = v327
		goto L34
	} else {
		goto L36
	}
L35:
	;
	if v267&int32(1) == int32(0) {
		v394 = v327
		goto L25
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	v336 = v329
	v357 = v327
	goto L30
L38:
	;
	v401 = v398
	goto L40
L39:
	;
	v401 = v112
	goto L40
L40:
	;
	if v400 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v402 = v89
	goto L43
L42:
	;
	v402 = v104
	goto L43
L43:
	;
	if v400 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v403 = v59
	goto L46
L45:
	;
	v403 = v103
	goto L46
L46:
	;
	v404 = int32(0)
	v406 = v89 + int32(1)
	v408 = v406 & int32(_a_F_g_cube_picksplit_1)
	if base.Ui32(v408) <= base.Ui32(v34) {
		v88 = v408
		v89 = v406
		v103 = v403
		v104 = v402
		v105 = v404
		v112 = v401
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
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v445+v428&int32(_a_F_g_cube_picksplit_1)<<(uint(int32(4))%32))))
	v593 = F_pg_detoast_datum(m, v592)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L62
	}
L50:
	;
	if v452 == int32(0) {
		v585 = v443
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	if v456 <= int32(0) {
		v585 = v443
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v460 = v452 + int32(8)
	if v456 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v550 = int32(3)
	v552 = v460 + v525<<(uint(v550)%32)
	v556 = *(*float64)(unsafe.Add(mBase, uint32(v552+v456<<(uint(v550)%32))))
	v557 = *(*float64)(unsafe.Add(mBase, uint32(v552)))
	v585 = base.F64_mul(v548, base.F64_abs(base.F64_sub(v556, v557)))
	goto L49
L54:
	;
	v525 = int32(0)
	v548 = float64(1)
	goto L53
L55:
	;
	goto L56
L56:
	;
	v470 = int32(0)
	v473 = v470
	v478 = v470
	v496 = float64(1)
	goto L57
L57:
	;
	v498 = int32(3)
	v500 = v460 + v473<<(uint(v498)%32)
	v502 = v456 << (uint(v498) % 32)
	v504 = *(*float64)(unsafe.Add(mBase, uint32(v500+v502)))
	v505 = *(*float64)(unsafe.Add(mBase, uint32(v500)))
	v512 = *(*float64)(unsafe.Add(mBase, uint32(v500+int32(8)+v502)))
	v513 = *(*float64)(unsafe.Add(mBase, uint32(v500)+8))
	v516 = base.F64_mul(base.F64_mul(v496, base.F64_abs(base.F64_sub(v504, v505))), base.F64_abs(base.F64_sub(v512, v513)))
	v517 = int32(2)
	v518 = v473 + v517
	v520 = v478 + v517
	if v520 != v456&int32(2147483646) {
		v473 = v518
		v478 = v520
		v496 = v516
		goto L57
	} else {
		goto L59
	}
L58:
	;
	if v456&int32(1) == int32(0) {
		v585 = v516
		goto L49
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	v525 = v518
	v548 = v516
	goto L53
L61:
	;
	if v30&int32(_a_F_g_cube_picksplit_1) != int32(1) {
		goto L73
	} else {
		goto L74
	}
L62:
	;
	if v593 == int32(0) {
		v727 = v23
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v593)+4))
	if v597 <= int32(0) {
		v727 = v23
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v601 = v593 + int32(8)
	if v597 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v691 = int32(3)
	v693 = v601 + v666<<(uint(v691)%32)
	v697 = *(*float64)(unsafe.Add(mBase, uint32(v693+v597<<(uint(v691)%32))))
	v698 = *(*float64)(unsafe.Add(mBase, uint32(v693)))
	v727 = base.F64_mul(v690, base.F64_abs(base.F64_sub(v697, v698)))
	goto L61
L66:
	;
	v666 = int32(0)
	v690 = float64(1)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v611 = int32(0)
	v614 = v611
	v619 = v611
	v638 = float64(1)
	goto L69
L69:
	;
	v639 = int32(3)
	v641 = v601 + v614<<(uint(v639)%32)
	v643 = v597 << (uint(v639) % 32)
	v645 = *(*float64)(unsafe.Add(mBase, uint32(v641+v643)))
	v646 = *(*float64)(unsafe.Add(mBase, uint32(v641)))
	v653 = *(*float64)(unsafe.Add(mBase, uint32(v641+int32(8)+v643)))
	v654 = *(*float64)(unsafe.Add(mBase, uint32(v641)+8))
	v657 = base.F64_mul(base.F64_mul(v638, base.F64_abs(base.F64_sub(v645, v646))), base.F64_abs(base.F64_sub(v653, v654)))
	v658 = int32(2)
	v659 = v614 + v658
	v661 = v619 + v658
	if v661 != v597&int32(2147483646) {
		v614 = v659
		v619 = v661
		v638 = v657
		goto L69
	} else {
		goto L71
	}
L70:
	;
	if v597&int32(1) == int32(0) {
		v727 = v657
		goto L61
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	v666 = v659
	v690 = v657
	goto L65
L73:
	;
	v732 = int32(1)
	v738 = v452
	v739 = v732
	v742 = v416
	v743 = v593
	v744 = v732
	v749 = v442
	v762 = v585
	v763 = v727
	goto L76
L74:
	;
	v1112 = v452
	v1116 = v416
	v1117 = v593
	v1123 = v442
	goto L75
L75:
	;
	v1138 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1116))) = uint16(v1138)
	*(*uint16)(unsafe.Add(mBase, uint32(v1123))) = uint16(v1138)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v1117
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v1112
	return v28
L76:
	;
	v764 = int32(_a_F_g_cube_picksplit_1)
	v765 = v744 & v764
	if v765 == v427&v764 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v1112 = v1081
	v1116 = v1085
	v1117 = v1086
	v1123 = v1092
	goto L75
L78:
	;
	v1108 = v744 + int32(1)
	v1110 = v1108 & int32(_a_F_g_cube_picksplit_1)
	if base.Ui32(v1110) <= base.Ui32((v30-v732)&int32(_a_F_g_cube_picksplit_1)) {
		v738 = v1081
		v739 = v1110
		v742 = v1085
		v743 = v1086
		v744 = v1108
		v749 = v1092
		v762 = v1105
		v763 = v1106
		goto L76
	} else {
		goto L113
	}
L79:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v749))) = uint16(v427)
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v770 + int32(1)
	v1081 = v738
	v1085 = v742
	v1086 = v743
	v1092 = v749 + int32(2)
	v1105 = v762
	v1106 = v763
	goto L78
L80:
	;
	goto L81
L81:
	;
	if v428&int32(_a_F_g_cube_picksplit_1) == v765 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v742))) = uint16(v428)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v780 + int32(1)
	v1081 = v738
	v1085 = v742 + int32(2)
	v1086 = v743
	v1092 = v749
	v1105 = v762
	v1106 = v763
	goto L78
L83:
	;
	goto L84
L84:
	;
	v786 = float64(0)
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v445+v739<<(uint(int32(4))%32))))
	v791 = F_pg_detoast_datum(m, v790)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v793 = F_cube_union_v0(m, v738, v791)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v795 = F_cube_union_v0(m, v743, v791)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v797 = float64(0)
	if v793 == int32(0) {
		v928 = v797
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v795 == int32(0) {
		v1060 = v786
		goto L99
	} else {
		goto L100
	}
L89:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v793)+4))
	if v800 <= int32(0) {
		v928 = v797
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v804 = v793 + int32(8)
	if v800 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v894 = int32(3)
	v896 = v804 + v869<<(uint(v894)%32)
	v900 = *(*float64)(unsafe.Add(mBase, uint32(v896+v800<<(uint(v894)%32))))
	v901 = *(*float64)(unsafe.Add(mBase, uint32(v896)))
	v928 = base.F64_mul(v891, base.F64_abs(base.F64_sub(v900, v901)))
	goto L88
L92:
	;
	v869 = int32(0)
	v891 = float64(1)
	goto L91
L93:
	;
	goto L94
L94:
	;
	v814 = int32(0)
	v817 = v814
	v828 = v814
	v839 = float64(1)
	goto L95
L95:
	;
	v842 = int32(3)
	v844 = v804 + v817<<(uint(v842)%32)
	v846 = v800 << (uint(v842) % 32)
	v848 = *(*float64)(unsafe.Add(mBase, uint32(v844+v846)))
	v849 = *(*float64)(unsafe.Add(mBase, uint32(v844)))
	v856 = *(*float64)(unsafe.Add(mBase, uint32(v844+int32(8)+v846)))
	v857 = *(*float64)(unsafe.Add(mBase, uint32(v844)+8))
	v860 = base.F64_mul(base.F64_mul(v839, base.F64_abs(base.F64_sub(v848, v849))), base.F64_abs(base.F64_sub(v856, v857)))
	v861 = int32(2)
	v862 = v817 + v861
	v864 = v828 + v861
	if v864 != v800&int32(2147483646) {
		v817 = v862
		v828 = v864
		v839 = v860
		goto L95
	} else {
		goto L97
	}
L96:
	;
	if v800&int32(1) == int32(0) {
		v928 = v860
		goto L88
	} else {
		goto L98
	}
L97:
	;
	goto L96
L98:
	;
	v869 = v862
	v891 = v860
	goto L91
L99:
	;
	if base.F64_lt(base.F64_sub(v928, v762), base.F64_sub(v1060, v763)) != 0 {
		goto L110
	} else {
		goto L111
	}
L100:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	if v933 <= int32(0) {
		v1060 = v786
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v937 = v795 + int32(8)
	if v933 == int32(1) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v1027 = int32(3)
	v1029 = v937 + v1002<<(uint(v1027)%32)
	v1033 = *(*float64)(unsafe.Add(mBase, uint32(v1029+v933<<(uint(v1027)%32))))
	v1034 = *(*float64)(unsafe.Add(mBase, uint32(v1029)))
	v1060 = base.F64_mul(v1023, base.F64_abs(base.F64_sub(v1033, v1034)))
	goto L99
L103:
	;
	v1002 = int32(0)
	v1023 = float64(1)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v947 = int32(0)
	v950 = v947
	v961 = v947
	v971 = float64(1)
	goto L106
L106:
	;
	v975 = int32(3)
	v977 = v937 + v950<<(uint(v975)%32)
	v979 = v933 << (uint(v975) % 32)
	v981 = *(*float64)(unsafe.Add(mBase, uint32(v977+v979)))
	v982 = *(*float64)(unsafe.Add(mBase, uint32(v977)))
	v989 = *(*float64)(unsafe.Add(mBase, uint32(v977+int32(8)+v979)))
	v990 = *(*float64)(unsafe.Add(mBase, uint32(v977)+8))
	v993 = base.F64_mul(base.F64_mul(v971, base.F64_abs(base.F64_sub(v981, v982))), base.F64_abs(base.F64_sub(v989, v990)))
	v994 = int32(2)
	v995 = v950 + v994
	v997 = v961 + v994
	if v997 != v933&int32(2147483646) {
		v950 = v995
		v961 = v997
		v971 = v993
		goto L106
	} else {
		goto L108
	}
L107:
	;
	if v933&int32(1) == int32(0) {
		v1060 = v993
		goto L99
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	v1002 = v995
	v1023 = v993
	goto L102
L110:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v749))) = uint16(v744)
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v1068 + int32(1)
	v1081 = v793
	v1085 = v742
	v1086 = v743
	v1092 = v749 + int32(2)
	v1105 = v928
	v1106 = v763
	goto L78
L111:
	;
	goto L112
L112:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v742))) = uint16(v744)
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v1075 + int32(1)
	v1081 = v738
	v1085 = v742 + int32(2)
	v1086 = v795
	v1092 = v749
	v1105 = v762
	v1106 = v1060
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
	v50 = int32(0)
	goto L5
L5:
	;
	return v50
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(4)+v23<<(uint(int32(4))%32))))
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
	v50 = v34
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
