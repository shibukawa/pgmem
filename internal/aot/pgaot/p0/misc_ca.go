package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_calc_hist_selectivity_contains(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) float64 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 float64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 float64
	_ = v82
	var v83 int32
	_ = v83
	var v86 float64
	_ = v86
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v96 float64
	_ = v96
	var v101 float64
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 float64
	_ = v111
	var v112 float64
	_ = v112
	var v116 float64
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 float64
	_ = v123
	var v124 int32
	_ = v124
	var v127 float64
	_ = v127
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v134 float64
	_ = v134
	var v137 float64
	_ = v137
	var v142 float64
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 float64
	_ = v152
	var v154 float64
	_ = v154
	var v158 int32
	_ = v158
	var v163 float64
	_ = v163
	var v174 float64
	_ = v174
	var v176 float64
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v207 int32
	_ = v207
	var v211 float64
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v231 float64
	_ = v231
	var v232 float64
	_ = v232
	var v233 float64
	_ = v233
	var v234 int32
	_ = v234
	var v235 float64
	_ = v235
	var v236 float64
	_ = v236
	var v238 int32
	_ = v238
	var v252 int32
	_ = v252
	var v261 float64
	_ = v261
	var v265 float64
	_ = v265
	var v267 int32
	_ = v267
	var v273 float64
	_ = v273
	var v276 float64
	_ = v276
	var v277 float64
	_ = v277
	var v285 int32
	_ = v285
	var v290 float64
	_ = v290
	var v292 float64
	_ = v292
	var v293 float64
	_ = v293
	var v299 int32
	_ = v299
	var v303 float64
	_ = v303
	var v310 float64
	_ = v310
	var v313 float64
	_ = v313
	var v323 float64
	_ = v323
	var v328 float64
	_ = v328
	var v331 float64
	_ = v331
	var v332 float64
	_ = v332
	var v333 int32
	_ = v333
	var v334 float64
	_ = v334
	var v336 int32
	_ = v336
	var v350 int32
	_ = v350
	var v359 float64
	_ = v359
	var v363 float64
	_ = v363
	var v364 float64
	_ = v364
	var v369 float64
	_ = v369
	var v376 int32
	_ = v376
	var v379 float64
	_ = v379
	var v381 float64
	_ = v381
	var v382 float64
	_ = v382
	var v384 float64
	_ = v384
	var v388 float64
	_ = v388
	var v392 float64
	_ = v392
	var v402 float64
	_ = v402
	var v422 float64
	_ = v422
	var v434 float64
	_ = v434
	var v439 int32
	_ = v439
	var v442 float64
	_ = v442
	var v443 float64
	_ = v443
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v464 float64
	_ = v464
	var v465 int32
	_ = v465
	var v468 float64
	_ = v468
	var v470 int32
	_ = v470
	var v471 int64
	_ = v471
	var v472 int64
	_ = v472
	var v473 int64
	_ = v473
	var v474 int32
	_ = v474
	var v475 float64
	_ = v475
	var v478 float64
	_ = v478
	var v483 float64
	_ = v483
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 float64
	_ = v493
	var v494 float64
	_ = v494
	var v498 int32
	_ = v498
	var v503 float64
	_ = v503
	var v514 float64
	_ = v514
	var v516 float64
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v547 int32
	_ = v547
	var v551 float64
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v570 int32
	_ = v570
	var v571 float64
	_ = v571
	var v572 float64
	_ = v572
	var v573 float64
	_ = v573
	var v574 int32
	_ = v574
	var v575 float64
	_ = v575
	var v576 float64
	_ = v576
	var v578 int32
	_ = v578
	var v592 int32
	_ = v592
	var v601 float64
	_ = v601
	var v605 float64
	_ = v605
	var v607 int32
	_ = v607
	var v613 float64
	_ = v613
	var v616 float64
	_ = v616
	var v617 float64
	_ = v617
	var v625 int32
	_ = v625
	var v630 float64
	_ = v630
	var v632 float64
	_ = v632
	var v633 float64
	_ = v633
	var v639 int32
	_ = v639
	var v643 float64
	_ = v643
	var v650 float64
	_ = v650
	var v653 float64
	_ = v653
	var v663 float64
	_ = v663
	var v668 float64
	_ = v668
	var v671 float64
	_ = v671
	var v672 float64
	_ = v672
	var v673 int32
	_ = v673
	var v674 float64
	_ = v674
	var v676 int32
	_ = v676
	var v690 int32
	_ = v690
	var v699 float64
	_ = v699
	var v703 float64
	_ = v703
	var v704 float64
	_ = v704
	var v709 float64
	_ = v709
	var v716 int32
	_ = v716
	var v719 float64
	_ = v719
	var v721 float64
	_ = v721
	var v722 float64
	_ = v722
	var v724 float64
	_ = v724
	var v728 float64
	_ = v728
	var v732 float64
	_ = v732
	var v742 float64
	_ = v742
	var v762 float64
	_ = v762
	var v772 float64
	_ = v772
	var v783 float64
	_ = v783
	v21 = l4 - int32(1)
	v34 = v21
	v35 = int32(-1)
	goto L1
L1:
	;
	v44 = base.I32_div_s(v34+v35+int32(1), int32(2))
	v48 = F_range_cmp_bounds(m, l0, l3+v44<<(uint(int32(4))%32), l1)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	if v54 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	return float64(0)
L4:
	;
	v53 = base.B2i32(v48 <= int32(0))
	if v48 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v54 = v44
	goto L7
L6:
	;
	v54 = v35
	goto L7
L7:
	;
	if v48 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v57 = v34
	goto L10
L9:
	;
	v57 = v44 - int32(1)
	goto L10
L10:
	;
	if v54 < v57 {
		v34 = v57
		v35 = v54
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L2
L12:
	;
	return float64(0)
L13:
	;
	goto L14
L14:
	;
	v64 = l0 + int32(268)
	v66 = l4 - int32(2)
	if base.Ui32(v54) < base.Ui32(v66) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v68 = v54
	goto L17
L16:
	;
	v68 = v66
	goto L17
L17:
	;
	v71 = l3 + v68<<(uint(int32(4))%32)
	v74 = F_get_position(m, l0, l1, v71, v71+int32(16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v76 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v116 = base.F64_convert_i32_u(v21)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+8)))
	if v117 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L20:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v81 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v103 != int32(1) {
		v112 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L19
	} else {
		goto L35
	}
L23:
	;
	v82 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L25
L24:
	;
	v82 = float64(1)
	goto L25
L25:
	;
	if v81 != 0 {
		v112 = v82
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v83 == int32(0) {
		v112 = v82
		goto L19
	} else {
		goto L27
	}
L27:
	;
	v86 = float64(1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v90 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v91 = F_FunctionCall2Coll(m, v64, v88, v89, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v93 = base.F64_reinterpret_i64(v91)
	if base.F64_lt(v93, float64(0)) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v96 = v86
	goto L31
L30:
	;
	v96 = v93
	goto L31
L31:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v91&int64(9223372036854775807)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v101 = v86
	goto L34
L33:
	;
	v101 = v96
	goto L34
L34:
	;
	v112 = v101
	goto L19
L35:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
	if v108 == v109 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v111 = float64(0)
	goto L38
L37:
	;
	v111 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L38
L38:
	;
	v112 = v111
	goto L19
L39:
	;
	v158 = int32(0)
	v163 = float64(0)
	if base.F64_lt(v154, v163) != 0 {
		v422 = v163
		goto L60
	} else {
		goto L61
	}
L40:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v122 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v144 != int32(1) {
		v154 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L39
	} else {
		goto L55
	}
L43:
	;
	v123 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L45
L44:
	;
	v123 = float64(1)
	goto L45
L45:
	;
	if v122 != 0 {
		v154 = v123
		goto L39
	} else {
		goto L46
	}
L46:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v124 == int32(0) {
		v154 = v123
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v127 = float64(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v130 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
	v132 = F_FunctionCall2Coll(m, v64, v129, v130, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v134 = base.F64_reinterpret_i64(v132)
	if base.F64_lt(v134, float64(0)) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v137 = v127
	goto L51
L50:
	;
	v137 = v134
	goto L51
L51:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v132&int64(9223372036854775807)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v142 = v127
	goto L54
L53:
	;
	v142 = v137
	goto L54
L54:
	;
	v154 = v142
	goto L39
L55:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+10)))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
	if v149 == v150 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v152 = float64(0)
	goto L58
L57:
	;
	v152 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L58
L58:
	;
	v154 = v152
	goto L39
L59:
	;
	v434 = base.F64_add(base.F64_div(base.F64_mul(v74, base.F64_sub(float64(1), v422)), v116), float64(0))
	if v68 != 0 {
		goto L131
	} else {
		goto L132
	}
L60:
	;
	goto L59
L61:
	;
	v174 = float64(1)
	v176 = base.F64_abs(v154)
	if base.F64_eq(v176, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v179 = v158
	goto L64
L63:
	;
	v179 = int32(0)
	goto L64
L64:
	;
	if v179 != 0 {
		v422 = v174
		goto L60
	} else {
		goto L65
	}
L65:
	;
	v182 = l6 - int32(1)
	if v182 < int32(0) {
		v422 = v174
		goto L60
	} else {
		goto L66
	}
L66:
	;
	v186 = v182
	v190 = int32(-1)
	goto L67
L67:
	;
	v207 = base.I32_div_s(v186+v190+int32(1), int32(2))
	v211 = *(*float64)(unsafe.Add(mBase, uint32(l5+v207<<(uint(int32(3))%32))))
	if base.F64_gt(v112, v211) != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	if v182 <= v221 {
		v422 = v174
		goto L60
	} else {
		goto L80
	}
L69:
	;
	if v221 < v219 {
		v186 = v219
		v190 = v221
		goto L67
	} else {
		goto L79
	}
L70:
	;
	v219 = v186
	v221 = v207
	goto L69
L71:
	;
	goto L72
L72:
	;
	v216 = v158 & base.F64_ge(v112, v211)
	if v216 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v217 = v186
	goto L75
L74:
	;
	v217 = v207 - int32(1)
	goto L75
L75:
	;
	if v216 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v218 = v207
	goto L78
L77:
	;
	v218 = v190
	goto L78
L78:
	;
	v219 = v217
	v221 = v218
	goto L69
L79:
	;
	goto L68
L80:
	;
	if v221 < int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v276 = base.F64_convert_i32_u(v182)
	v277 = base.F64_div(base.F64_add(v273, base.F64_convert_i32_u(v267)), v276)
	if base.F64_eq(v112, v154) != 0 {
		v422 = v277
		goto L60
	} else {
		goto L96
	}
L82:
	;
	v267 = int32(0)
	v273 = float64(0)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v230 = l5 + v221<<(uint(int32(3))%32)
	v231 = *(*float64)(unsafe.Add(mBase, uint32(v230)+8))
	v232 = base.F64_abs(v231)
	v233 = math.Float64frombits(uint64(0x7ff0000000000000))
	v234 = base.F64_eq(v232, v233)
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v230)))
	v236 = base.F64_abs(v235)
	v238 = base.F64_eq(v236, v233)
	if v234|v238 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if base.F64_eq(base.F64_abs(v112), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v267 = v221
		v273 = float64(0.5)
		goto L81
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v252 = int32(0)
	if v234|base.B2i32(v238 == v252) == v252 {
		v267 = v221
		v273 = float64(1)
		goto L81
	} else {
		goto L89
	}
L88:
	;
	v267 = v221
	v273 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v231, v112), base.F64_sub(v231, v235)))
	goto L81
L89:
	;
	if base.F64_eq(v232, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v261 = float64(0)
	goto L92
L91:
	;
	v261 = float64(0.5)
	goto L92
L92:
	;
	if base.F64_eq(v236, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v265 = v261
	goto L95
L94:
	;
	v265 = float64(0.5)
	goto L95
L95:
	;
	v267 = v221
	v273 = v265
	goto L81
L96:
	;
	if v182 <= v267 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v388 = float64(0)
	v392 = base.F64_div(base.F64_add(v384, base.F64_convert_i32_u(v376)), v276)
	if base.F64_gt(v381, v388)|base.F64_gt(v392, v388) != 0 {
		goto L124
	} else {
		goto L125
	}
L98:
	;
	v376 = v267
	v379 = v112
	v381 = v277
	v382 = v163
	v384 = v163
	goto L97
L99:
	;
	goto L100
L100:
	;
	v285 = v267
	v290 = v277
	v292 = v163
	v293 = v112
	goto L102
L101:
	;
	v328 = *(*float64)(unsafe.Add(mBase, uint32(l5+v285<<(uint(int32(3))%32))))
	if base.F64_eq(v303, v328) != 0 {
		goto L109
	} else {
		goto L110
	}
L102:
	;
	v299 = v285 + int32(1)
	v303 = *(*float64)(unsafe.Add(mBase, uint32(l5+v299<<(uint(int32(3))%32))))
	if base.F64_lt(v303, v154)|v158&base.F64_ge(v154, v303) == int32(0) {
		goto L101
	} else {
		goto L104
	}
L103:
	;
	v376 = v182
	v379 = v303
	v381 = v313
	v382 = v323
	v384 = v163
	goto L97
L104:
	;
	v310 = float64(0)
	v313 = base.F64_div(base.F64_convert_i32_u(v285), v276)
	if base.F64_gt(v290, v310)|base.F64_gt(v313, v310) != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v323 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v290, v313), float64(0.5)), base.F64_sub(v303, v293)), v292)
	goto L107
L106:
	;
	v323 = v292
	goto L107
L107:
	;
	if v182 != v299 {
		v285 = v299
		v290 = v313
		v292 = v323
		v293 = v303
		goto L102
	} else {
		goto L108
	}
L108:
	;
	goto L103
L109:
	;
	v369 = float64(0)
	goto L111
L110:
	;
	v331 = base.F64_abs(v303)
	v332 = math.Float64frombits(uint64(0x7ff0000000000000))
	v333 = base.F64_eq(v331, v332)
	v334 = base.F64_abs(v328)
	v336 = base.F64_eq(v334, v332)
	if v333|v336 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	v376 = v285
	v379 = v293
	v381 = v290
	v382 = v292
	v384 = v369
	goto L97
L112:
	;
	v369 = v364
	goto L111
L113:
	;
	if base.F64_eq(base.F64_abs(v154), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v364 = float64(0.5)
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v350 = int32(0)
	if v333|base.B2i32(v336 == v350) == v350 {
		v364 = float64(1)
		goto L112
	} else {
		goto L117
	}
L116:
	;
	v364 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v303, v154), base.F64_sub(v303, v328)))
	goto L112
L117:
	;
	if base.F64_eq(v331, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v359 = float64(0)
	goto L120
L119:
	;
	v359 = float64(0.5)
	goto L120
L120:
	;
	if base.F64_eq(v334, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v363 = v359
	goto L123
L122:
	;
	v363 = float64(0.5)
	goto L123
L123:
	;
	v364 = v363
	goto L112
L124:
	;
	v402 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v381, v392), float64(0.5)), base.F64_sub(v154, v379)), v382)
	goto L126
L125:
	;
	v402 = v382
	goto L126
L126:
	;
	if base.F64_eq(v176, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if base.F64_eq(base.F64_abs(v402), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v422 = float64(0.5)
		goto L60
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v422 = base.F64_div(v402, base.F64_sub(v154, v112))
	goto L60
L130:
	;
	goto L129
L131:
	;
	v439 = v68
	v442 = v154
	v443 = v434
	goto L134
L132:
	;
	v783 = v434
	goto L133
L133:
	;
	return v783
L134:
	;
	v454 = v439 - int32(1)
	v457 = l3 + v454<<(uint(int32(4))%32)
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457)+8)))
	if v458 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v783 = v772
	goto L133
L136:
	;
	v498 = int32(0)
	v503 = float64(0)
	if base.F64_lt(v494, v503) != 0 {
		v762 = v503
		goto L157
	} else {
		goto L158
	}
L137:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v463 != 0 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L139
L139:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v485 != int32(1) {
		v494 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L136
	} else {
		goto L152
	}
L140:
	;
	v464 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L142
L141:
	;
	v464 = float64(1)
	goto L142
L142:
	;
	if v463 != 0 {
		v494 = v464
		goto L136
	} else {
		goto L143
	}
L143:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v465 == int32(0) {
		v494 = v464
		goto L136
	} else {
		goto L144
	}
L144:
	;
	v468 = float64(1)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v471 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v472 = *(*int64)(unsafe.Add(mBase, uint32(v457)))
	v473 = F_FunctionCall2Coll(m, v64, v470, v471, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L3
	} else {
		goto L145
	}
L145:
	;
	v475 = base.F64_reinterpret_i64(v473)
	if base.F64_lt(v475, float64(0)) != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v478 = v468
	goto L148
L147:
	;
	v478 = v475
	goto L148
L148:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v473&int64(9223372036854775807)) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v483 = v468
	goto L151
L150:
	;
	v483 = v478
	goto L151
L151:
	;
	v494 = v483
	goto L136
L152:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457)+10)))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
	if v490 == v491 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v493 = float64(0)
	goto L155
L154:
	;
	v493 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L155
L155:
	;
	v494 = v493
	goto L136
L156:
	;
	v772 = base.F64_add(v443, base.F64_div(base.F64_sub(float64(1), v762), v116))
	if base.Ui32(int32(1)) < base.Ui32(v439) {
		v439 = v454
		v442 = v494
		v443 = v772
		goto L134
	} else {
		goto L228
	}
L157:
	;
	goto L156
L158:
	;
	v514 = float64(1)
	v516 = base.F64_abs(v494)
	if base.F64_eq(v516, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v519 = v498
	goto L161
L160:
	;
	v519 = int32(0)
	goto L161
L161:
	;
	if v519 != 0 {
		v762 = v514
		goto L157
	} else {
		goto L162
	}
L162:
	;
	v522 = l6 - int32(1)
	if v522 < int32(0) {
		v762 = v514
		goto L157
	} else {
		goto L163
	}
L163:
	;
	v526 = v522
	v530 = int32(-1)
	goto L164
L164:
	;
	v547 = base.I32_div_s(v526+v530+int32(1), int32(2))
	v551 = *(*float64)(unsafe.Add(mBase, uint32(l5+v547<<(uint(int32(3))%32))))
	if base.F64_gt(v442, v551) != 0 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	if v522 <= v561 {
		v762 = v514
		goto L157
	} else {
		goto L177
	}
L166:
	;
	if v561 < v559 {
		v526 = v559
		v530 = v561
		goto L164
	} else {
		goto L176
	}
L167:
	;
	v559 = v526
	v561 = v547
	goto L166
L168:
	;
	goto L169
L169:
	;
	v556 = v498 & base.F64_ge(v442, v551)
	if v556 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v557 = v526
	goto L172
L171:
	;
	v557 = v547 - int32(1)
	goto L172
L172:
	;
	if v556 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v558 = v547
	goto L175
L174:
	;
	v558 = v530
	goto L175
L175:
	;
	v559 = v557
	v561 = v558
	goto L166
L176:
	;
	goto L165
L177:
	;
	if v561 < int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v616 = base.F64_convert_i32_u(v522)
	v617 = base.F64_div(base.F64_add(v613, base.F64_convert_i32_u(v607)), v616)
	if base.F64_eq(v442, v494) != 0 {
		v762 = v617
		goto L157
	} else {
		goto L193
	}
L179:
	;
	v607 = int32(0)
	v613 = float64(0)
	goto L178
L180:
	;
	goto L181
L181:
	;
	v570 = l5 + v561<<(uint(int32(3))%32)
	v571 = *(*float64)(unsafe.Add(mBase, uint32(v570)+8))
	v572 = base.F64_abs(v571)
	v573 = math.Float64frombits(uint64(0x7ff0000000000000))
	v574 = base.F64_eq(v572, v573)
	v575 = *(*float64)(unsafe.Add(mBase, uint32(v570)))
	v576 = base.F64_abs(v575)
	v578 = base.F64_eq(v576, v573)
	if v574|v578 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	if base.F64_eq(base.F64_abs(v442), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v607 = v561
		v613 = float64(0.5)
		goto L178
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v592 = int32(0)
	if v574|base.B2i32(v578 == v592) == v592 {
		v607 = v561
		v613 = float64(1)
		goto L178
	} else {
		goto L186
	}
L185:
	;
	v607 = v561
	v613 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v571, v442), base.F64_sub(v571, v575)))
	goto L178
L186:
	;
	if base.F64_eq(v572, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v601 = float64(0)
	goto L189
L188:
	;
	v601 = float64(0.5)
	goto L189
L189:
	;
	if base.F64_eq(v576, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v605 = v601
	goto L192
L191:
	;
	v605 = float64(0.5)
	goto L192
L192:
	;
	v607 = v561
	v613 = v605
	goto L178
L193:
	;
	if v522 <= v607 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v728 = float64(0)
	v732 = base.F64_div(base.F64_add(v724, base.F64_convert_i32_u(v716)), v616)
	if base.F64_gt(v721, v728)|base.F64_gt(v732, v728) != 0 {
		goto L221
	} else {
		goto L222
	}
L195:
	;
	v716 = v607
	v719 = v442
	v721 = v617
	v722 = v503
	v724 = v503
	goto L194
L196:
	;
	goto L197
L197:
	;
	v625 = v607
	v630 = v617
	v632 = v503
	v633 = v442
	goto L199
L198:
	;
	v668 = *(*float64)(unsafe.Add(mBase, uint32(l5+v625<<(uint(int32(3))%32))))
	if base.F64_eq(v643, v668) != 0 {
		goto L206
	} else {
		goto L207
	}
L199:
	;
	v639 = v625 + int32(1)
	v643 = *(*float64)(unsafe.Add(mBase, uint32(l5+v639<<(uint(int32(3))%32))))
	if base.F64_lt(v643, v494)|v498&base.F64_ge(v494, v643) == int32(0) {
		goto L198
	} else {
		goto L201
	}
L200:
	;
	v716 = v522
	v719 = v643
	v721 = v653
	v722 = v663
	v724 = v503
	goto L194
L201:
	;
	v650 = float64(0)
	v653 = base.F64_div(base.F64_convert_i32_u(v625), v616)
	if base.F64_gt(v630, v650)|base.F64_gt(v653, v650) != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v663 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v630, v653), float64(0.5)), base.F64_sub(v643, v633)), v632)
	goto L204
L203:
	;
	v663 = v632
	goto L204
L204:
	;
	if v522 != v639 {
		v625 = v639
		v630 = v653
		v632 = v663
		v633 = v643
		goto L199
	} else {
		goto L205
	}
L205:
	;
	goto L200
L206:
	;
	v709 = float64(0)
	goto L208
L207:
	;
	v671 = base.F64_abs(v643)
	v672 = math.Float64frombits(uint64(0x7ff0000000000000))
	v673 = base.F64_eq(v671, v672)
	v674 = base.F64_abs(v668)
	v676 = base.F64_eq(v674, v672)
	if v673|v676 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L208:
	;
	v716 = v625
	v719 = v633
	v721 = v630
	v722 = v632
	v724 = v709
	goto L194
L209:
	;
	v709 = v704
	goto L208
L210:
	;
	if base.F64_eq(base.F64_abs(v494), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v704 = float64(0.5)
		goto L209
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v690 = int32(0)
	if v673|base.B2i32(v676 == v690) == v690 {
		v704 = float64(1)
		goto L209
	} else {
		goto L214
	}
L213:
	;
	v704 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v643, v494), base.F64_sub(v643, v668)))
	goto L209
L214:
	;
	if base.F64_eq(v671, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v699 = float64(0)
	goto L217
L216:
	;
	v699 = float64(0.5)
	goto L217
L217:
	;
	if base.F64_eq(v674, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v703 = v699
	goto L220
L219:
	;
	v703 = float64(0.5)
	goto L220
L220:
	;
	v704 = v703
	goto L209
L221:
	;
	v742 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v721, v732), float64(0.5)), base.F64_sub(v494, v719)), v722)
	goto L223
L222:
	;
	v742 = v722
	goto L223
L223:
	;
	if base.F64_eq(v516, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	if base.F64_eq(base.F64_abs(v742), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v762 = float64(0.5)
		goto L157
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v762 = base.F64_div(v742, base.F64_sub(v494, v442))
	goto L157
L227:
	;
	goto L226
L228:
	;
	goto L135
}
func F_calc_joinrel_size_estimate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 float64, l6 int32, l7 int32) float64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 float64
	_ = v56
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v336 int32
	_ = v336
	var v352 int32
	_ = v352
	var v365 int32
	_ = v365
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 float64
	_ = v394
	var v395 float64
	_ = v395
	var v398 float64
	_ = v398
	var v399 float64
	_ = v399
	var v401 float64
	_ = v401
	var v403 float64
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v416 int32
	_ = v416
	var v421 float64
	_ = v421
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 float64
	_ = v458
	var v459 int32
	_ = v459
	var v463 float64
	_ = v463
	var v464 float64
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v479 float64
	_ = v479
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 float64
	_ = v502
	var v505 float64
	_ = v505
	var v508 float64
	_ = v508
	var v516 int32
	_ = v516
	var v518 float64
	_ = v518
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v561 int32
	_ = v561
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v662 int32
	_ = v662
	var v677 float64
	_ = v677
	var v678 int32
	_ = v678
	var v680 float64
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v687 float64
	_ = v687
	var v688 int32
	_ = v688
	var v697 float64
	_ = v697
	var v699 float64
	_ = v699
	var v718 float64
	_ = v718
	var v720 float64
	_ = v720
	var v724 float64
	_ = v724
	var v726 float64
	_ = v726
	var v728 float64
	_ = v728
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v756 float64
	_ = v756
	var v760 float64
	_ = v760
	var v769 float64
	_ = v769
	var v773 float64
	_ = v773
	v12 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v33 == v12 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if int32(1)<<(uint(v32)%32)&int32(174) != 0 {
		goto L121
	} else {
		goto L122
	}
L2:
	;
	v516 = l7
	v518 = float64(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v38 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v516 = l7
	v518 = float64(1)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v47 = base.B2i32(v32&int32(-2) != int32(4))
	v50 = l7
	v56 = float64(1)
	v71 = v12
	goto L8
L8:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v71<<(uint(int32(2))%32))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v81 = F_bms_is_member(m, v80, v43)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v502 = float64(0)
	if base.F64_lt(v479, v502) != 0 {
		v508 = v502
		goto L117
	} else {
		goto L118
	}
L10:
	;
	v499 = v71 + int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v499 < v500 {
		v50 = v473
		v56 = v479
		v71 = v499
		goto L8
	} else {
		goto L116
	}
L11:
	;
	if v50 == l7 {
		goto L44
	} else {
		goto L45
	}
L12:
	;
	if v32&int32(-2) != int32(4) {
		goto L24
	} else {
		goto L25
	}
L13:
	;
	return float64(0)
L14:
	;
	if v81 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v86 = F_bms_is_member(m, v85, v42)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v89 = F_bms_is_member(m, v88, v43)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L13
	} else {
		goto L20
	}
L18:
	;
	if v86 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	if v89 == int32(0) {
		v473 = v50
		v479 = v56
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v94 = F_bms_is_member(m, v93, v42)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v96 = int32(0)
	if base.B2i32(v94 == v96)|base.B2i32(v47 == v96) != 0 {
		v473 = v50
		v479 = v56
		goto L10
	} else {
		goto L23
	}
L23:
	;
	v152 = int32(0)
	goto L11
L24:
	;
	v152 = int32(0)
	goto L11
L25:
	;
	goto L26
L26:
	;
	v104 = int32(0)
	if v42 == v104 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v149 != int32(1) {
		v473 = v50
		v479 = v56
		goto L10
	} else {
		goto L43
	}
L28:
	;
	v149 = int32(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v112 = int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v113 <= v112 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v116 = v112
	goto L33
L32:
	;
	v116 = v113
	goto L33
L33:
	;
	v120 = int32(0)
	v122 = v104
	goto L34
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(8)+v120<<(uint(int32(2))%32))))
	if v129 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v149 = v141
	goto L27
L36:
	;
	goto L35
L37:
	;
	v130 = int32(2)
	if v122 != 0 {
		v141 = v130
		goto L36
	} else {
		goto L40
	}
L38:
	;
	v136 = v122
	goto L39
L39:
	;
	v138 = v120 + int32(1)
	if v138 != v116 {
		v120 = v138
		v122 = v136
		goto L34
	} else {
		goto L42
	}
L40:
	;
	v131 = int32(1)
	if base.Ui32(v131) < base.Ui32(base.I32_popcnt(v129)) {
		v141 = v130
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v136 = v131
	goto L39
L42:
	;
	v141 = v136
	goto L36
L43:
	;
	v152 = int32(1)
	goto L11
L44:
	;
	v154 = F_list_copy(m, v50)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L13
	} else {
		goto L47
	}
L45:
	;
	v156 = v50
	goto L46
L46:
	;
	if v156 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v156 = v154
	goto L46
L48:
	;
	v159 = int32(0)
	v161 = F_list_concat(m, v159, v159)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v166 = v79 + int32(288)
	v167 = int32(0)
	v171 = v156
	v181 = v167
	v184 = v167
	goto L52
L51:
	;
	v473 = v161
	v479 = v56
	goto L10
L52:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	if v181 < v196 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v365 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L54:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if int32(0) < v198 {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	v352 = v171
	v365 = v184
	goto L56
L56:
	;
	goto L53
L57:
	;
	if v323 != 0 {
		v171 = v323
		v181 = v324 + int32(1)
		v184 = v336
		goto L52
	} else {
		goto L86
	}
L58:
	;
	v317 = F_list_delete_nth_cell(m, v171, v181)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L13
	} else {
		goto L84
	}
L59:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201+v181<<(uint(int32(2))%32))))
	v210 = int32(0)
	v220 = v198
	goto L62
L60:
	;
	goto L61
L61:
	;
	v323 = v171
	v324 = v181
	v336 = v184
	goto L57
L62:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v205)+60))
	if v234 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L61
L64:
	;
	v286 = v210 + int32(1)
	if v286 < v284 {
		v210 = v286
		v220 = v284
		goto L62
	} else {
		goto L83
	}
L65:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v166+v210<<(uint(int32(2))%32))))
	if v238 != v234 {
		v284 = v220
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v79+int32(544)+v210<<(uint(int32(2))%32))))
	v244 = int32(0)
	if v243 == v244 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L58
L69:
	;
	if v282 != 0 {
		goto L58
	} else {
		goto L82
	}
L70:
	;
	v282 = int32(0)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if v250 <= int32(0) {
		v276 = v244
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v282 = v276
	goto L69
L74:
	;
	v253 = int32(0)
	if v253 < v250 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v256 = v250
	goto L77
L76:
	;
	v256 = v253
	goto L77
L77:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	v259 = int32(0)
	goto L78
L78:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v257+v259<<(uint(int32(2))%32))))
	v268 = base.B2i32(v267 == v205)
	if v267 == v205 {
		v276 = v268
		goto L73
	} else {
		goto L80
	}
L79:
	;
	v276 = v268
	goto L73
L80:
	;
	v270 = v259 + int32(1)
	if v270 != v256 {
		v259 = v270
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v284 = v283
	goto L64
L83:
	;
	goto L63
L84:
	;
	v319 = F_lappend(m, v184, v205)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	v323 = v317
	v324 = v181 - int32(1)
	v336 = v319
	goto L57
L86:
	;
	v352 = v323
	v365 = v336
	goto L56
L87:
	;
	v380 = F_list_concat(m, v352, int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L13
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v79)+284))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v79)+272))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v79)+276))
	if v382 != v383+(v384-v385) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v473 = v380
	v479 = v56
	goto L10
L91:
	;
	v389 = F_list_concat(m, v352, v365)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L13
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v392 = F_find_base_rel(m, l0, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L13
	} else {
		goto L95
	}
L94:
	;
	v473 = v389
	v479 = v56
	goto L10
L95:
	;
	v394 = *(*float64)(unsafe.Add(mBase, uint32(v392)+128))
	v395 = float64(1)
	if base.F64_gt(v394, v395) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v398 = v394
	goto L98
L97:
	;
	v398 = v395
	goto L98
L98:
	;
	if v152 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v392)+16))
	v401 = v399
	goto L101
L100:
	;
	v401 = float64(1)
	goto L101
L101:
	;
	v403 = base.F64_mul(v56, base.F64_div(v401, v398))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v79)+276))
	if v404 <= int32(0) {
		v473 = v352
		v479 = v403
		goto L10
	} else {
		goto L102
	}
L102:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v407 <= int32(0) {
		v473 = v352
		v479 = v403
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v416 = int32(0)
	v421 = v403
	goto L104
L104:
	;
	v441 = v416 << (uint(int32(2)) % 32)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v166+v441)))
	if v443 == int32(0) {
		v464 = v421
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v473 = v352
	v479 = v464
	goto L10
L106:
	;
	v468 = v416 + int32(1)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v468 < v469 {
		v416 = v468
		v421 = v464
		goto L104
	} else {
		goto L115
	}
L107:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+40)))
	if v446 != int32(1) {
		v464 = v421
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v79+int32(416)+v441)))
	v451 = int32(0)
	v453 = F_ec_search_derived_clause_for_ems(m, l0, v443, v450, v451, v451)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	if v453 == int32(0) {
		v464 = v421
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v458 = F_clause_selectivity(m, l0, v453, int32(0), v32, l6)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L13
	} else {
		goto L111
	}
L111:
	;
	if base.F64_gt(v458, float64(0)) != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v463 = base.F64_div(v421, v458)
	goto L114
L113:
	;
	v463 = v421
	goto L114
L114:
	;
	v464 = v463
	goto L106
L115:
	;
	goto L105
L116:
	;
	goto L9
L117:
	;
	v516 = v473
	v518 = v508
	goto L1
L118:
	;
	v505 = float64(1)
	if base.F64_gt(v479, v505) != 0 {
		v508 = v505
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v516 = v473
	v518 = v479
	goto L1
L120:
	;
	switch v32 {
	case 0:
		goto L158
	case 1:
		goto L163
	case 2:
		goto L162
	default:
		goto L159
	case 4:
		goto L161
	case 5:
		goto L160
	}
L121:
	;
	v540 = int32(0)
	if v516 == v540 {
		v651 = v540
		v662 = v540
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v687 = F_clauselist_selectivity(m, l0, v516, int32(0), v32, l6)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L13
	} else {
		goto L156
	}
L124:
	;
	v677 = F_clauselist_selectivity(m, l0, v662, int32(0), v32, l6)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L13
	} else {
		goto L152
	}
L125:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v544 <= int32(0) {
		v651 = v540
		v662 = v540
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v550 = v540
	v551 = int32(0)
	v561 = v540
	goto L127
L127:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v575+v551<<(uint(int32(2))%32))))
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+8)))
	if v580 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v651 = v643
	v662 = v644
	goto L124
L129:
	;
	v646 = v551 + int32(1)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v646 < v647 {
		v550 = v643
		v551 = v646
		v561 = v644
		goto L127
	} else {
		goto L151
	}
L130:
	;
	v641 = F_lappend(m, v561, v579)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L13
	} else {
		goto L150
	}
L131:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v579)+32))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v585 = int32(0)
	if v583 == v585 {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	goto L133
L133:
	;
	v639 = F_lappend(m, v550, v579)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L13
	} else {
		goto L149
	}
L134:
	;
	if v638 != 0 {
		goto L130
	} else {
		goto L148
	}
L135:
	;
	v638 = int32(1)
	goto L134
L136:
	;
	goto L137
L137:
	;
	if v584 == int32(0) {
		v631 = v585
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v638 = v631
	goto L134
L139:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v583)+4))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	if v595 < v594 {
		v631 = v585
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v597 = int32(1)
	if v594 <= v597 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v600 = v597
	goto L143
L142:
	;
	v600 = v594
	goto L143
L143:
	;
	v601 = int32(8)
	v606 = int32(0)
	goto L144
L144:
	;
	v613 = v606 << (uint(int32(2)) % 32)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v583+v601+v613)))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v584+v601+v613)))
	v620 = v615 & (v617 ^ int32(-1))
	v622 = base.B2i32(v620 == int32(0))
	if v620 != 0 {
		v631 = v622
		goto L138
	} else {
		goto L146
	}
L145:
	;
	v631 = v622
	goto L138
L146:
	;
	v624 = v606 + int32(1)
	if v624 != v600 {
		v606 = v624
		goto L144
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	goto L133
L149:
	;
	v643 = v639
	v644 = v561
	goto L129
L150:
	;
	v643 = v550
	v644 = v641
	goto L129
L151:
	;
	goto L128
L152:
	;
	v680 = F_clauselist_selectivity(m, l0, v651, int32(0), v32, l6)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L13
	} else {
		goto L153
	}
L153:
	;
	F_list_free(m, v662)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L13
	} else {
		goto L154
	}
L154:
	;
	F_list_free(m, v651)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L13
	} else {
		goto L155
	}
L155:
	;
	v697 = v677
	v699 = v680
	goto L120
L156:
	;
	v697 = v687
	v699 = float64(0)
	goto L120
L157:
	;
	m.G0 = v30 + int32(16)
	v760 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v756)&int64(9223372036854775807)))|base.F64_gt(v756, v760) != 0 {
		v773 = v760
		goto L176
	} else {
		goto L177
	}
L158:
	;
	v756 = base.F64_mul(base.F64_mul(base.F64_mul(l4, l5), v518), v697)
	goto L157
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L13
	} else {
		goto L173
	}
L160:
	;
	v756 = base.F64_mul(v699, base.F64_mul(l4, base.F64_sub(float64(1), base.F64_mul(v518, v697))))
	goto L157
L161:
	;
	v756 = base.F64_mul(base.F64_mul(l4, v518), v697)
	goto L157
L162:
	;
	v724 = base.F64_mul(base.F64_mul(base.F64_mul(l4, l5), v518), v697)
	if base.F64_gt(l4, v724) != 0 {
		goto L167
	} else {
		goto L168
	}
L163:
	;
	v718 = base.F64_mul(base.F64_mul(base.F64_mul(l4, l5), v518), v697)
	if base.F64_gt(l4, v718) != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v720 = l4
	goto L166
L165:
	;
	v720 = v718
	goto L166
L166:
	;
	v756 = base.F64_mul(v699, v720)
	goto L157
L167:
	;
	v726 = l4
	goto L169
L168:
	;
	v726 = v724
	goto L169
L169:
	;
	if base.F64_lt(v726, l5) != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v728 = l5
	goto L172
L171:
	;
	v728 = v726
	goto L172
L172:
	;
	v756 = base.F64_mul(v699, v728)
	goto L157
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v32
	F_errmsg_internal(m, int32(_a_F_calc_joinrel_size_estimate_0), v30)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L13
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_calc_joinrel_size_estimate_1), int32(_a_F_calc_joinrel_size_estimate_2), int32(_a_F_calc_joinrel_size_estimate_3))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L13
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	return v773
L177:
	;
	v769 = float64(1)
	if base.F64_le(v756, v769) != 0 {
		v773 = v769
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v773 = base.F64_nearest(v756)
	goto L176
}
func F_calc_rank_cd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 float32
	_ = v32
	var v39 float64
	_ = v39
	var v42 float32
	_ = v42
	var v48 float64
	_ = v48
	var v52 float64
	_ = v52
	var v56 float32
	_ = v56
	var v62 float64
	_ = v62
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v71 float32
	_ = v71
	var v78 float64
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v483 int32
	_ = v483
	var v500 int32
	_ = v500
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v547 int32
	_ = v547
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v587 int32
	_ = v587
	var __phi587 int32
	_ = __phi587
	var v589 int32
	_ = v589
	var __phi589 int32
	_ = __phi589
	var v591 int32
	_ = v591
	var __phi591 int32
	_ = __phi591
	var v593 int32
	_ = v593
	var __phi593 int32
	_ = __phi593
	var v594 int32
	_ = v594
	var __phi594 int32
	_ = __phi594
	var v596 int32
	_ = v596
	var __phi596 int32
	_ = __phi596
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
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
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 float64
	_ = v686
	var v697 int32
	_ = v697
	var v704 int32
	_ = v704
	var v708 float64
	_ = v708
	var v709 float64
	_ = v709
	var v710 float64
	_ = v710
	var v715 int32
	_ = v715
	var v725 int32
	_ = v725
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v831 int32
	_ = v831
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v945 int32
	_ = v945
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v982 int32
	_ = v982
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
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
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1078 int32
	_ = v1078
	var v1100 int32
	_ = v1100
	var v1106 int32
	_ = v1106
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1195 int32
	_ = v1195
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1232 int32
	_ = v1232
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1278 float64
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1302 float64
	_ = v1302
	var v1308 int32
	_ = v1308
	var v1314 float64
	_ = v1314
	var v1315 float64
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1341 float64
	_ = v1341
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1366 float64
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1379 float64
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1422 int32
	_ = v1422
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1485 float64
	_ = v1485
	var v1507 float64
	_ = v1507
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1536 int32
	_ = v1536
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1601 float64
	_ = v1601
	var v1612 int32
	_ = v1612
	var v1622 float64
	_ = v1622
	var v1627 int32
	_ = v1627
	var v1633 float64
	_ = v1633
	var v1638 int32
	_ = v1638
	var v1644 float64
	_ = v1644
	var v1649 float64
	_ = v1649
	var v1655 float64
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1688 float32
	_ = v1688
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1706 int32
	_ = v1706
	var v1711 int32
	_ = v1711
	v5 = int32(0)
	v27 = m.G0
	v29 = v27 + int32(-64)
	m.G0 = v29
	v32 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	if base.F32_ge(v32, float32(0)) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L20
	} else {
		goto L244
	}
L2:
	;
	if base.F32_gt(v32, float32(1)) != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v39 = float64(0.10000000149011612)
	goto L4
L4:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+16)) = base.F64_div(float64(1), v39)
	v42 = *(*float32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.F32_ge(v42, float32(0)) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v39 = base.F64_promote_f32(v32)
	goto L4
L6:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+24)) = base.F64_div(float64(1), v52)
	v56 = *(*float32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.F32_ge(v56, float32(0)) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v52 = float64(0.20000000298023224)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v48 = base.F64_promote_f32(v42)
	*(*float64)(unsafe.Add(mBase, uint32(v29)+24)) = v48
	if base.F32_gt(v42, float32(1)) != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v52 = v48
	goto L6
L11:
	;
	v67 = float64(1)
	*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = base.F64_div(v67, v66)
	v71 = *(*float32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.F32_ge(v71, float32(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v66 = float64(0.4000000059604645)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v62 = base.F64_promote_f32(v56)
	*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = v62
	if base.F32_gt(v56, float32(1)) != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v66 = v62
	goto L11
L16:
	;
	if base.F32_gt(v71, float32(1)) != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v78 = float64(1)
	goto L18
L18:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+40)) = base.F64_div(v67, v78)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = l2
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v84 = F_palloc0_mul(m, int32(_a_F_calc_rank_cd_0), v83)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v78 = base.F64_promote_f32(v71)
	goto L18
L20:
	;
	return float32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v84
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v92 = v90 << (uint(int32(2)) % 32)
	v93 = F_palloc_mul(m, int32(12), v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v95 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	m.G0 = v29 - int32(-64)
	return v1688
L24:
	;
	F_pg_qsort(m, v513, v515, int32(12), int32(1730))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L20
	} else {
		goto L94
	}
L25:
	;
	v98 = int32(8)
	v101 = l1 + v98
	v104 = l2
	v106 = v92
	v110 = v93
	v111 = v5
	v112 = v5
	goto L28
L26:
	;
	v539 = v84
	v547 = v93
	goto L27
L27:
	;
	F_pfree(m, v547)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L20
	} else {
		goto L92
	}
L28:
	;
	v130 = l2 + v98 + v111*int32(12)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v131 != int32(1) {
		v509 = v106
		v513 = v110
		v515 = v112
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if int32(0) < v515 {
		goto L24
	} else {
		goto L91
	}
L30:
	;
	v532 = v111 + int32(1)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if v532 < v534 {
		v104 = v533
		v106 = v509
		v110 = v513
		v111 = v532
		v112 = v515
		goto L28
	} else {
		goto L90
	}
L31:
	;
	v135 = v27 + int32(-4)
	v136 = int32(0)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v136
	v146 = l1 + int32(8)
	v149 = v146 + v142<<(uint(int32(2))%32)
	if v142 <= v136 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if v289 == int32(0) {
		v509 = v106
		v513 = v110
		v515 = v112
		goto L30
	} else {
		goto L62
	}
L33:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+2)))
	if v217 != int32(1) {
		goto L49
	} else {
		goto L50
	}
L34:
	;
	v211 = v149
	v212 = v146
	v213 = v149
	goto L33
L35:
	;
	goto L36
L36:
	;
	v159 = v146
	v160 = v149
	goto L37
L37:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v165 = int32(12)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v175 = int32(2)
	v182 = base.I32_div_s((v160-v159)>>(uint(v175)%32), v175)
	v185 = v159 + v182<<(uint(v175)%32)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v194 = int32(0)
	v195 = F_tsCompareString(m, v104+int32(8)+v164*v165+int32(base.Ui32(v168)>>(uint(v165)%32)), v168&int32(4095), v146+v174<<(uint(v175)%32)+int32(base.Ui32(v186)>>(uint(v165)%32)), int32(base.Ui32(v186)>>(uint(int32(1))%32))&int32(2047), v194)
	mBase = m.M
	if v195 == v194 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v211 = v185
	v212 = v204
	v213 = v205
	goto L33
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(1)
	v211 = v185
	v212 = v159
	v213 = v185
	goto L33
L40:
	;
	goto L41
L41:
	;
	v203 = base.B2i32(int32(0) < v195)
	if int32(0) < v195 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v204 = v185 + int32(4)
	goto L44
L43:
	;
	v204 = v159
	goto L44
L44:
	;
	if int32(0) < v195 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v205 = v160
	goto L47
L46:
	;
	v205 = v185
	goto L47
L47:
	;
	if base.Ui32(v204) < base.Ui32(v205) {
		v159 = v204
		v160 = v205
		goto L37
	} else {
		goto L48
	}
L48:
	;
	goto L38
L49:
	;
	v285 = int32(0)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v285 < v286 {
		goto L59
	} else {
		goto L60
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(0)
	if base.Ui32(v212) < base.Ui32(v213) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v223 = v211
	goto L53
L52:
	;
	v223 = v213
	goto L53
L53:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v146+v224<<(uint(int32(2))%32)) <= base.Ui32(v223) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v235 = v224
	v236 = v223
	goto L55
L55:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v242 = int32(12)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	v258 = int32(1)
	v263 = F_tsCompareString(m, v104+int32(8)+v241*v242+int32(base.Ui32(v245)>>(uint(v242)%32)), v245&int32(4095), v146+v235<<(uint(int32(2))%32)+int32(base.Ui32(v254)>>(uint(v242)%32)), int32(base.Ui32(v254)>>(uint(v258)%32))&int32(2047), v258)
	mBase = m.M
	if v263 != 0 {
		goto L49
	} else {
		goto L57
	}
L56:
	;
	goto L49
L57:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v264 + int32(1)
	v269 = v236 + int32(4)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v269) < base.Ui32(v146+v270<<(uint(int32(2))%32)) {
		v235 = v270
		v236 = v269
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v289 = v213
	goto L61
L60:
	;
	v289 = v285
	goto L61
L61:
	;
	goto L32
L62:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	if v292 <= int32(0) {
		v509 = v106
		v513 = v110
		v515 = v112
		goto L30
	} else {
		goto L63
	}
L63:
	;
	v295 = v292
	v297 = v289
	v299 = v106
	v303 = v110
	v305 = v112
	goto L64
L64:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	if v321&int32(1) != 0 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v509 = v386
	v513 = v390
	v515 = v483
	goto L30
L66:
	;
	if v340 != 0 {
		goto L77
	} else {
		goto L78
	}
L67:
	;
	v353 = v299
	v357 = v303
	goto L73
L68:
	;
	v324 = int32(1)
	v334 = (int32(base.Ui32(v321)>>(uint(v324)%32))&int32(2047) + int32(base.Ui32(v321)>>(uint(int32(12))%32)) + v324) & int32(_a_F_calc_rank_cd_1)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v337 = v335 << (uint(int32(2)) % 32)
	v340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334+(v101+v337)))))
	v341 = v305 + v340
	if v299 <= v341 {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v344 = v297 + int32(4)
	if (v344-v289)>>(uint(int32(2))%32) < v295 {
		v297 = v344
		goto L64
	} else {
		goto L72
	}
L71:
	;
	v386 = v299
	v390 = v303
	goto L66
L72:
	;
	v509 = v299
	v513 = v303
	v515 = v305
	goto L30
L73:
	;
	v377 = F_repalloc(m, v357, v353*int32(24))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L20
	} else {
		goto L75
	}
L74:
	;
	v386 = v380
	v390 = v377
	goto L66
L75:
	;
	v380 = v353 << (uint(int32(1)) % 32)
	if v380 <= v341 {
		v353 = v380
		v357 = v377
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v411 = l1 + v337 + v334 + int32(10)
	v413 = int32(0)
	v423 = v305
	goto L80
L78:
	;
	v473 = v295
	v483 = v305
	goto L79
L79:
	;
	v500 = v297 + int32(4)
	if (v500-v289)>>(uint(int32(2))%32) < v473 {
		v295 = v473
		v297 = v500
		v299 = v386
		v303 = v390
		v305 = v483
		goto L64
	} else {
		goto L89
	}
L80:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	if v439 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v473 = v472
	v483 = v467
	goto L79
L82:
	;
	v470 = v413 + int32(1)
	if v470 != v340 {
		v413 = v470
		v423 = v467
		goto L80
	} else {
		goto L88
	}
L83:
	;
	v460 = v390 + v423*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v460)+4)) = v297
	*(*uint16)(unsafe.Add(mBase, uint32(v460)+8)) = uint16(v457)
	*(*int32)(unsafe.Add(mBase, uint32(v460))) = v130
	v467 = v423 + int32(1)
	goto L82
L84:
	;
	v445 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v411+v413<<(uint(int32(1))%32)))))
	v457 = v445
	goto L83
L85:
	;
	goto L86
L86:
	;
	v446 = int32(1)
	v449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v411+v413<<(uint(v446)%32)))))
	if int32(base.Ui32(v439)>>(uint(int32(base.Ui32(v449)>>(uint(int32(14))%32)))%32))&v446 == int32(0) {
		v467 = v423
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v457 = v449
	goto L83
L88:
	;
	goto L81
L89:
	;
	goto L65
L90:
	;
	goto L29
L91:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v539 = v538
	v547 = v513
	goto L27
L92:
	;
	F_pfree(m, v539)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L20
	} else {
		goto L93
	}
L93:
	;
	v1688 = float32(0)
	goto L23
L94:
	;
	v574 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v513)+8)))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	v577 = F_palloc_mul(m, int32(4), v576)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L20
	} else {
		goto L95
	}
L95:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	*(*int32)(unsafe.Add(mBase, uint32(v577))) = v579
	if v515 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v651)+8)) = uint16(v658)
	*(*uint16)(unsafe.Add(mBase, uint32(v651)+4)) = uint16(v657)
	*(*int32)(unsafe.Add(mBase, uint32(v651))) = v660
	v681 = int32(12)
	v682 = v651 - v513 + v681
	v684 = base.I32_div_s(v682, v681)
	v686 = float64(0)
	v697 = int32(0)
	v704 = v5
	v708 = v686
	v709 = v686
	v710 = float64(0)
	goto L108
L97:
	;
	v651 = v513
	v657 = int32(1)
	v658 = v574
	v660 = v577
	goto L96
L98:
	;
	goto L99
L99:
	;
	__phi587 = v513
	__phi589 = v513
	__phi591 = v513 + int32(12)
	__phi593 = int32(1)
	__phi594 = v574
	__phi596 = v577
	v587 = __phi587
	v589 = __phi589
	v591 = __phi591
	v593 = __phi593
	v594 = __phi594
	v596 = __phi596
	goto L100
L100:
	;
	v613 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+20)))
	v614 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+8)))
	if v613 != v614 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v651 = v641
	v657 = v644
	v658 = v642
	v660 = v643
	goto L96
L102:
	;
	v645 = int32(12)
	v646 = v591 + v645
	v649 = base.I32_div_s(v646-v513, v645)
	if v649 < v515 {
		__phi587 = v641
		__phi589 = v591
		__phi591 = v646
		__phi593 = v644
		__phi594 = v642
		__phi596 = v643
		v587 = __phi587
		v589 = __phi589
		v591 = __phi591
		v593 = __phi593
		v594 = __phi594
		v596 = __phi596
		goto L100
	} else {
		goto L107
	}
L103:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v587)+8)) = uint16(v594)
	*(*uint16)(unsafe.Add(mBase, uint32(v587)+4)) = uint16(v593)
	*(*int32)(unsafe.Add(mBase, uint32(v587))) = v596
	v630 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589)+20)))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	v634 = F_palloc_mul(m, int32(4), v633)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L20
	} else {
		goto L106
	}
L104:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v589)+16))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v589)+4))
	if v616 != v617 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v589)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v596+base.I32_extend16_s(v593)<<(uint(int32(2))%32)))) = v623
	v641 = v587
	v642 = v594
	v643 = v596
	v644 = v593 + int32(1)
	goto L102
L106:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v589)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v634))) = v636
	v641 = v587 + int32(12)
	v642 = v630
	v643 = v634
	v644 = int32(1)
	goto L102
L107:
	;
	goto L101
L108:
	;
	if v684 < v697 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if l3&int32(1) == int32(0) {
		v1507 = v708
		goto L205
	} else {
		goto L206
	}
L110:
	;
	v715 = v697
	goto L112
L111:
	;
	v715 = v684
	goto L112
L112:
	;
	v725 = v697
	goto L113
L113:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L20
	} else {
		goto L115
	}
L114:
	;
	goto L109
L115:
	;
	v744 = int32(0)
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v745)+4))
	if v744 < v746 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v751 = v744
	goto L119
L117:
	;
	goto L118
L118:
	;
	if v725 == v715 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v776 = v751 * int32(_a_F_calc_rank_cd_0)
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v779 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v776+v777))) = uint8(v779)
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v781+v776)+1)) = uint8(v779)
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v785+v776)+4)) = v779
	v790 = v751 + int32(1)
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v791)+4))
	if v790 < v792 {
		v751 = v790
		goto L119
	} else {
		goto L121
	}
L120:
	;
	goto L118
L121:
	;
	goto L120
L122:
	;
	goto L114
L123:
	;
	v822 = v725 * int32(12)
	v823 = v513 + v822
	v824 = v823
	v831 = v822
	goto L124
L124:
	;
	v850 = int32(*(*int16)(unsafe.Add(mBase, uint32(v824)+4)))
	if int32(0) < v850 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L122
L126:
	;
	v856 = int32(0)
	goto L129
L127:
	;
	goto L128
L128:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v989 = F_TS_execute(m, v982+int32(8), v27+int32(-56), int32(0), int32(1731))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L20
	} else {
		goto L148
	}
L129:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v824)))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v880+v856<<(uint(int32(2))%32))))
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v884))))
	if v885 != int32(1) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	goto L128
L131:
	;
	v953 = v856 + int32(1)
	v954 = int32(*(*int16)(unsafe.Add(mBase, uint32(v824)+4)))
	if v953 < v954 {
		v856 = v953
		goto L129
	} else {
		goto L147
	}
L132:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v894 = base.I32_div_s(v884-v889-int32(8), int32(12))
	v897 = v888 + v894*int32(_a_F_calc_rank_cd_0)
	v898 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v897))) = uint8(v898)
	v900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v897)+1)))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v897)+4))
	if v901 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v897)+4)) = v945
	goto L131
L134:
	;
	if v900&int32(1) != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L136
L136:
	;
	v913 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v824)+8)))
	v915 = v897 + int32(8)
	v918 = int32(1)
	v921 = v900 & v918
	if v921 != 0 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	v908 = int32(_a_F_calc_rank_cd_2)
	goto L139
L138:
	;
	v908 = int32(0)
	goto L139
L139:
	;
	v910 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v824)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v897+v908)+8)) = uint16(v910)
	v945 = int32(1)
	goto L133
L140:
	;
	v922 = int32(_a_F_calc_rank_cd_3) - v901
	goto L142
L141:
	;
	v922 = v901 - v918
	goto L142
L142:
	;
	v926 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v915+v922<<(uint(int32(1))%32)))))
	if (v913^v926)&int32(_a_F_calc_rank_cd_4) == int32(0) {
		goto L131
	} else {
		goto L143
	}
L143:
	;
	if v921 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v934 = int32(_a_F_calc_rank_cd_4) - v901
	goto L146
L145:
	;
	v934 = v901
	goto L146
L146:
	;
	v935 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v915+v934<<(uint(v935)%32)))) = uint16(v913)
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v897)+4))
	v945 = v939 + v935
	goto L133
L147:
	;
	goto L130
L148:
	;
	if v989 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v991 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v824)+8)))
	v993 = v991 & int32(_a_F_calc_rank_cd_4)
	if v993 == int32(0) {
		goto L122
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v1389 = v824 + int32(12)
	v1390 = v1389 - v513
	if v1390 < v682 {
		v824 = v1389
		v831 = v1390
		goto L124
	} else {
		goto L204
	}
L152:
	;
	v996 = int32(0)
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v997)+4))
	if v996 < v998 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v1003 = v996
	goto L156
L154:
	;
	goto L155
L155:
	;
	if v831 < v822 {
		goto L160
	} else {
		goto L161
	}
L156:
	;
	v1028 = v1003 * int32(_a_F_calc_rank_cd_0)
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v1031 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1028+v1029))) = uint8(v1031)
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v1035 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1033+v1028)+1)) = uint8(v1035)
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1037+v1028)+4)) = v1031
	v1042 = v1003 + v1035
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+4))
	if v1042 < v1044 {
		v1003 = v1042
		goto L156
	} else {
		goto L158
	}
L157:
	;
	goto L155
L158:
	;
	goto L157
L159:
	;
	v1278 = float64(0)
	if base.Ui32(v1078) <= base.Ui32(v824) {
		goto L192
	} else {
		goto L193
	}
L160:
	;
	v725 = v725 + int32(1)
	goto L113
L161:
	;
	v1078 = v831 + v513
	goto L162
L162:
	;
	v1100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1078)+4)))
	if int32(0) < v1100 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v1246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1078)+8)))
	v1248 = v1246 & int32(_a_F_calc_rank_cd_4)
	if base.Ui32(v1248) <= base.Ui32(v993) {
		goto L159
	} else {
		goto L191
	}
L164:
	;
	v1106 = int32(0)
	goto L167
L165:
	;
	goto L166
L166:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v1239 = F_TS_execute(m, v1232+int32(8), v27+int32(-56), int32(0), int32(1731))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L20
	} else {
		goto L186
	}
L167:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1078)))
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1130+v1106<<(uint(int32(2))%32))))
	v1135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1134))))
	if v1135 != int32(1) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L166
L169:
	;
	v1203 = v1106 + int32(1)
	v1204 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1078)+4)))
	if v1203 < v1204 {
		v1106 = v1203
		goto L167
	} else {
		goto L185
	}
L170:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v1144 = base.I32_div_s(v1134-v1139-int32(8), int32(12))
	v1147 = v1138 + v1144*int32(_a_F_calc_rank_cd_0)
	v1148 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1147))) = uint8(v1148)
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1147)+1)))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1147)+4))
	if v1151 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1147)+4)) = v1195
	goto L169
L172:
	;
	if v1150&int32(1) != 0 {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	goto L174
L174:
	;
	v1163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1078)+8)))
	v1165 = v1147 + int32(8)
	v1168 = int32(1)
	v1171 = v1150 & v1168
	if v1171 != 0 {
		goto L178
	} else {
		goto L179
	}
L175:
	;
	v1158 = int32(_a_F_calc_rank_cd_2)
	goto L177
L176:
	;
	v1158 = int32(0)
	goto L177
L177:
	;
	v1160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1078)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1147+v1158)+8)) = uint16(v1160)
	v1195 = int32(1)
	goto L171
L178:
	;
	v1172 = int32(_a_F_calc_rank_cd_3) - v1151
	goto L180
L179:
	;
	v1172 = v1151 - v1168
	goto L180
L180:
	;
	v1176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1165+v1172<<(uint(int32(1))%32)))))
	if (v1163^v1176)&int32(_a_F_calc_rank_cd_4) == int32(0) {
		goto L169
	} else {
		goto L181
	}
L181:
	;
	if v1171 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v1184 = int32(_a_F_calc_rank_cd_4) - v1151
	goto L184
L183:
	;
	v1184 = v1151
	goto L184
L184:
	;
	v1185 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1165+v1184<<(uint(v1185)%32)))) = uint16(v1163)
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1147)+4))
	v1195 = v1189 + v1185
	goto L171
L185:
	;
	goto L168
L186:
	;
	if v1239 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v1244 = v1078 - int32(12)
	if base.Ui32(v823) <= base.Ui32(v1244) {
		v1078 = v1244
		goto L162
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	goto L163
L190:
	;
	goto L160
L191:
	;
	goto L160
L192:
	;
	v1282 = v1078
	v1302 = v1278
	goto L195
L193:
	;
	v1341 = v1278
	goto L194
L194:
	;
	v1345 = v824 - v1078
	v1347 = base.I32_div_s(v1345, int32(12))
	v1353 = base.I32_div_s(v1345, int32(24))
	v1355 = v993 - (v1347 + v1248)
	if v1355 < int32(0) {
		goto L198
	} else {
		goto L199
	}
L195:
	;
	v1308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1282)+8)))
	v1314 = *(*float64)(unsafe.Add(mBase, uint32(v27+int32(-48)+int32(base.Ui32(v1308)>>(uint(int32(11))%32))&int32(24))))
	v1315 = base.F64_add(v1302, v1314)
	v1317 = v1282 + int32(12)
	if base.Ui32(v1317) <= base.Ui32(v824) {
		v1282 = v1317
		v1302 = v1315
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v1341 = v1315
	goto L194
L197:
	;
	goto L196
L198:
	;
	v1358 = v1353
	goto L200
L199:
	;
	v1358 = v1355
	goto L200
L200:
	;
	v1366 = base.F64_mul(base.F64_convert_i32_u(v1248+v993), float64(0.5))
	v1368 = int32(0)
	if base.B2i32(base.F64_lt(v710, v1366) == v1368)|base.B2i32(v704 <= v1368) == v1368 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1379 = base.F64_add(v709, base.F64_div(float64(1), base.F64_sub(v1366, v710)))
	goto L203
L202:
	;
	v1379 = v709
	goto L203
L203:
	;
	v1382 = base.I32_div_s(v1078-v513, int32(12))
	v1383 = int32(1)
	v697 = v1382 + v1383
	v704 = v704 + v1383
	v708 = base.F64_add(v708, base.F64_div(base.F64_div(base.F64_convert_i32_s(v1347+int32(1)), v1341), base.F64_convert_i32_s(v1358+int32(1))))
	v709 = v1379
	v710 = v1366
	goto L108
L204:
	;
	goto L125
L205:
	;
	if l3&int32(2) == int32(0) {
		v1601 = v1507
		goto L217
	} else {
		goto L218
	}
L206:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1422 <= int32(0) {
		v1507 = v708
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v1427 = v101 + v1422<<(uint(int32(2))%32)
	v1431 = v101
	v1433 = int32(0)
	goto L208
L208:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1431)))
	if v1455&int32(1) != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v1485 = F_log(m, base.F64_convert_i32_s(v1478+int32(1)))
	mBase = m.M
	v1507 = base.F64_div(v708, v1485)
	goto L205
L210:
	;
	v1458 = int32(1)
	v1471 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1427+(int32(base.Ui32(v1455)>>(uint(v1458)%32))&int32(2047)+int32(base.Ui32(v1455)>>(uint(int32(12))%32))+v1458)&int32(_a_F_calc_rank_cd_1)))))
	if base.Ui32(v1471) <= base.Ui32(v1458) {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	v1477 = int32(1)
	goto L212
L212:
	;
	v1478 = v1477 + v1433
	v1480 = v1431 + int32(4)
	if base.Ui32(v1480) < base.Ui32(v1427) {
		v1431 = v1480
		v1433 = v1478
		goto L208
	} else {
		goto L216
	}
L213:
	;
	v1474 = v1458
	goto L215
L214:
	;
	v1474 = v1471
	goto L215
L215:
	;
	v1477 = v1474
	goto L212
L216:
	;
	goto L209
L217:
	;
	v1612 = int32(0)
	if base.B2i32(base.F64_gt(v709, float64(0)) == v1612)|(base.B2i32(l3&int32(4) == v1612)|base.B2i32(v704 <= v1612)) != 0 {
		goto L230
	} else {
		goto L231
	}
L218:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1517 <= int32(0) {
		v1601 = v1507
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v1522 = v101 + v1517<<(uint(int32(2))%32)
	v1524 = int32(0)
	v1536 = v101
	goto L220
L220:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1536)))
	if v1550&int32(1) != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	if v1573 <= int32(0) {
		v1601 = v1507
		goto L217
	} else {
		goto L229
	}
L222:
	;
	v1553 = int32(1)
	v1566 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1522+(int32(base.Ui32(v1550)>>(uint(v1553)%32))&int32(2047)+int32(base.Ui32(v1550)>>(uint(int32(12))%32))+v1553)&int32(_a_F_calc_rank_cd_1)))))
	if base.Ui32(v1566) <= base.Ui32(v1553) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	v1572 = int32(1)
	goto L224
L224:
	;
	v1573 = v1572 + v1524
	v1575 = v1536 + int32(4)
	if base.Ui32(v1575) < base.Ui32(v1522) {
		v1524 = v1573
		v1536 = v1575
		goto L220
	} else {
		goto L228
	}
L225:
	;
	v1569 = v1553
	goto L227
L226:
	;
	v1569 = v1566
	goto L227
L227:
	;
	v1572 = v1569
	goto L224
L228:
	;
	goto L221
L229:
	;
	v1601 = base.F64_div(v1507, base.F64_convert_i32_u(v1573))
	goto L217
L230:
	;
	v1622 = v1601
	goto L232
L231:
	;
	v1622 = base.F64_div(v1601, base.F64_div(base.F64_convert_i32_u(v704), v709))
	goto L232
L232:
	;
	if l3&int32(8) == int32(0) {
		v1633 = v1622
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if l3&int32(16) == int32(0) {
		v1649 = v1633
		goto L236
	} else {
		goto L237
	}
L234:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1627 <= int32(0) {
		v1633 = v1622
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v1633 = base.F64_div(v1622, base.F64_convert_i32_u(v1627))
	goto L233
L236:
	;
	if l3&int32(32) != 0 {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1638 <= int32(0) {
		v1649 = v1633
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v1644 = F_log(m, base.F64_convert_i32_s(v1638+int32(1)))
	mBase = m.M
	v1649 = base.F64_div(v1633, base.F64_div(v1644, float64(0.6931471805599453)))
	goto L236
L239:
	;
	v1655 = base.F64_div(v1649, base.F64_add(v1649, float64(1)))
	goto L241
L240:
	;
	v1655 = v1649
	goto L241
L241:
	;
	F_pfree(m, v513)
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L20
	} else {
		goto L242
	}
L242:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	F_pfree(m, v1658)
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L20
	} else {
		goto L243
	}
L243:
	;
	v1688 = base.F32_demote_f64(v1655)
	goto L23
L244:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L20
	} else {
		goto L245
	}
L245:
	;
	F_errmsg(m, int32(_a_F_calc_rank_cd_5), int32(0))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L20
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_calc_rank_cd_6), int32(899), int32(_a_F_calc_rank_cd_7))
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L20
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cancel_prior_stmt_triggers(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v167 int32
	_ = v167
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[0]))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[1]))
	v15 = v10 + v12*int32(20)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v16 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+10)))
	if v79 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v54 = int32(_a_F_cancel_prior_stmt_triggers_0)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[2]))
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[2])) = v58
	v61 = F_palloc0(m, int32(36))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v27 = int32(0)
	goto L5
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22+v27<<(uint(int32(2))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 != l0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	v44 = v27 + int32(1)
	if v19 != v44 {
		v27 = v44
		goto L5
	} else {
		goto L11
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v38 != l1 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+8)))
	if v40 != int32(1) {
		v75 = v35
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	goto L6
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = l0
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v66 = F_lappend(m, v65, v61)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v66
	*(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[2])) = v55
	v75 = v61
	goto L1
L15:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+10)) = uint8(v163)
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	*(*int64)(unsafe.Add(mBase, uint32(v75)+12)) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v167
	return
L16:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	if v82 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v91 = v88
	v93 = v89
	goto L22
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	v88 = v82
	v89 = v83
	goto L17
L19:
	;
	goto L20
L20:
	;
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v85 == v84 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v88 = v85
	v89 = v84
	goto L17
L22:
	;
	if v93 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L15
L24:
	;
	v100 = v93
	goto L26
L25:
	;
	v100 = v91 + int32(16)
	goto L26
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if base.Ui32(v100) < base.Ui32(v101) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v106 = v100
	goto L30
L28:
	;
	goto L29
L29:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v154 != 0 {
		v91 = v154
		v93 = int32(0)
		goto L22
	} else {
		goto L41
	}
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v114 = v106 + v111&int32(134217727)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	if v115 != l0 {
		goto L15
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if base.B2i32(v117&int32(3) != l2)|v117&int32(28) != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v111&int32(1073741823) | int32(-2147483648)
	v131 = v111 & int32(939524096)
	if v131 == int32(134217728) {
		v141 = int32(24)
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v142 = v141 + v106
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if base.Ui32(v142) < base.Ui32(v143) {
		v106 = v142
		goto L30
	} else {
		goto L40
	}
L35:
	;
	if v131 != int32(805306368) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v131 == int32(268435456) {
		v141 = int32(12)
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v141 = int32(16)
	goto L34
L39:
	;
	v141 = int32(4)
	goto L34
L40:
	;
	goto L31
L41:
	;
	goto L23
}
func F_canonicalize_ec_expression(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
	v7 = F_exprType(m, l0)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if l1 <= int32(3830) {
			if l1 <= int32(2775) {
				switch l1 - int32(2277) {
				case 0, 6:
					v37 = v7
					v39 = F_exprCollation(m, l0)
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if v39 == l2 {
							v52 = l0
							return v52
						} else {
							v42 = F_exprTypmod(m, l0)
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = v37
								v46 = v42
								v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = v50
									return v52
								}
							}
						}
					}
				case 1, 2, 3, 4, 5:
					if l1 != v7 {
						v44 = l1
						v46 = int32(-1)
						v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = v50
							return v52
						}
					} else {
						v37 = l1
						v39 = F_exprCollation(m, l0)
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 == l2 {
								v52 = l0
								return v52
							} else {
								v42 = F_exprTypmod(m, l0)
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v44 = v37
									v46 = v42
									v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										v52 = v50
										return v52
									}
								}
							}
						}
					}
				default:
					if l1 != int32(2249) {
						if l1 != v7 {
							v44 = l1
							v46 = int32(-1)
							v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = v50
								return v52
							}
						} else {
							v37 = l1
							v39 = F_exprCollation(m, l0)
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								if v39 == l2 {
									v52 = l0
									return v52
								} else {
									v42 = F_exprTypmod(m, l0)
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										v44 = v37
										v46 = v42
										v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											v52 = v50
											return v52
										}
									}
								}
							}
						}
					} else {
						v37 = v7
						v39 = F_exprCollation(m, l0)
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 == l2 {
								v52 = l0
								return v52
							} else {
								v42 = F_exprTypmod(m, l0)
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v44 = v37
									v46 = v42
									v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										v52 = v50
										return v52
									}
								}
							}
						}
					}
				}
			} else {
				if l1 == int32(2776) {
					v37 = v7
					v39 = F_exprCollation(m, l0)
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if v39 == l2 {
							v52 = l0
							return v52
						} else {
							v42 = F_exprTypmod(m, l0)
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = v37
								v46 = v42
								v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = v50
									return v52
								}
							}
						}
					}
				} else {
					if l1 != int32(3500) {
						if l1 != v7 {
							v44 = l1
							v46 = int32(-1)
							v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = v50
								return v52
							}
						} else {
							v37 = l1
							v39 = F_exprCollation(m, l0)
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								if v39 == l2 {
									v52 = l0
									return v52
								} else {
									v42 = F_exprTypmod(m, l0)
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										v44 = v37
										v46 = v42
										v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											v52 = v50
											return v52
										}
									}
								}
							}
						}
					} else {
						v37 = v7
						v39 = F_exprCollation(m, l0)
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 == l2 {
								v52 = l0
								return v52
							} else {
								v42 = F_exprTypmod(m, l0)
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v44 = v37
									v46 = v42
									v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										v52 = v50
										return v52
									}
								}
							}
						}
					}
				}
			}
		} else {
			if base.B2i32(base.Ui32(l1-int32(_a_F_canonicalize_ec_expression_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l1-int32(_a_F_canonicalize_ec_expression_1)) < base.Ui32(int32(2)))|base.B2i32(l1 == int32(3831)) != 0 {
				v37 = v7
				v39 = F_exprCollation(m, l0)
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					if v39 == l2 {
						v52 = l0
						return v52
					} else {
						v42 = F_exprTypmod(m, l0)
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = v37
							v46 = v42
							v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = v50
								return v52
							}
						}
					}
				}
			} else {
				if l1 != v7 {
					v44 = l1
					v46 = int32(-1)
					v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = v50
						return v52
					}
				} else {
					v37 = l1
					v39 = F_exprCollation(m, l0)
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if v39 == l2 {
							v52 = l0
							return v52
						} else {
							v42 = F_exprTypmod(m, l0)
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = v37
								v46 = v42
								v50 = F_applyRelabelType(m, l0, v44, v46, l2, int32(2), int32(-1), int32(0))
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = v50
									return v52
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_canonicalize_qual(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = F_find_duplicate_ors(m, l0, l1)
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_casemap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	if base.Ui32(int32(_a_F_casemap_0)) < base.Ui32(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	v12 = int32(255)
	v13 = l0 & v12
	v14 = int32(3)
	v15 = base.I32_div_u_s(v13, v14)
	v21 = int32(2)
	v23 = *(*int32)(unsafe.Add(mBase, uint32((l0-v15*v14)&v12<<(uint(v21)%32))+uint32(_c_F_casemap[0])))
	v24 = int32(8)
	v25 = int32(base.Ui32(l0) >> (uint(v24) % 32))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_casemap[1]))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v26*int32(86))+uint32(_c_F_casemap[1]))))
	v35 = base.I32_rem_u_s(int32(base.Ui32(v23*v30)>>(uint(int32(11))%32)), int32(6))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_casemap[2]))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35<<(uint(v21)%32)+v38<<(uint(v21)%32))+uint32(_c_F_casemap[3])))
	v44 = v42 >> (uint(v24) % 32)
	v46 = v42 & v12
	if base.Ui32(v46) <= base.Ui32(int32(1)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v44&(int32(0)-(l1^v46)) + l0
L4:
	;
	goto L5
L5:
	;
	v56 = v44 & int32(255)
	if v56 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v63 = int32(base.Ui32(v44) >> (uint(int32(8)) % 32))
	v64 = v56
	goto L7
L7:
	;
	v70 = int32(1)
	v71 = int32(base.Ui32(v64) >> (uint(v70) % 32))
	v72 = v71 + v63
	v74 = v72 << (uint(v70) % 32)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_casemap[4]))))
	if v75 == v13 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L1
L9:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_casemap[5]))))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79<<(uint(int32(2))%32))+uint32(_c_F_casemap[3])))
	v84 = v82 & int32(255)
	if base.Ui32(v84) <= base.Ui32(int32(1)) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v100 = base.B2i32(base.Ui32(v13) < base.Ui32(v75))
	if base.Ui32(v13) < base.Ui32(v75) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	return (int32(0)-(l1^v84))&(v82>>(uint(int32(8))%32)) + l0
L13:
	;
	goto L14
L14:
	;
	if l1 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v97 = int32(-1)
	goto L17
L16:
	;
	v97 = int32(1)
	goto L17
L17:
	;
	return v97 + l0
L18:
	;
	v101 = v63
	goto L20
L19:
	;
	v101 = v72
	goto L20
L20:
	;
	if base.Ui32(v13) < base.Ui32(v75) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v103 = v71
	goto L23
L22:
	;
	v103 = v64 - v71
	goto L23
L23:
	;
	if v103 != 0 {
		v63 = v101
		v64 = v103
		goto L7
	} else {
		goto L24
	}
L24:
	;
	goto L8
}
func F_cashlarger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if v5 < v4 {
		v7 = v4
	} else {
		v7 = v5
	}
	return v7
}
