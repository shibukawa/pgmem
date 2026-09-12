package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_calc_hist_selectivity_contains(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) float64 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 float64
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 float64
	_ = v80
	var v81 int32
	_ = v81
	var v84 float64
	_ = v84
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
	var v91 float64
	_ = v91
	var v97 float64
	_ = v97
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 float64
	_ = v110
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 float64
	_ = v134
	var v140 float64
	_ = v140
	var v143 float64
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v163 float64
	_ = v163
	var v174 float64
	_ = v174
	var v179 float64
	_ = v179
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 float64
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 float64
	_ = v237
	var v238 float64
	_ = v238
	var v239 float64
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 float64
	_ = v242
	var v243 float64
	_ = v243
	var v245 int32
	_ = v245
	var v262 float64
	_ = v262
	var v266 float64
	_ = v266
	var v268 int32
	_ = v268
	var v274 float64
	_ = v274
	var v277 float64
	_ = v277
	var v278 float64
	_ = v278
	var v286 int32
	_ = v286
	var v292 float64
	_ = v292
	var v294 float64
	_ = v294
	var v295 float64
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 float64
	_ = v306
	var v313 int32
	_ = v313
	var v316 float64
	_ = v316
	var v319 float64
	_ = v319
	var v329 float64
	_ = v329
	var v334 int32
	_ = v334
	var v335 float64
	_ = v335
	var v338 float64
	_ = v338
	var v339 float64
	_ = v339
	var v340 int32
	_ = v340
	var v341 float64
	_ = v341
	var v343 int32
	_ = v343
	var v360 float64
	_ = v360
	var v364 float64
	_ = v364
	var v365 float64
	_ = v365
	var v370 float64
	_ = v370
	var v377 int32
	_ = v377
	var v381 float64
	_ = v381
	var v383 float64
	_ = v383
	var v384 float64
	_ = v384
	var v386 float64
	_ = v386
	var v390 float64
	_ = v390
	var v394 float64
	_ = v394
	var v404 float64
	_ = v404
	var v425 float64
	_ = v425
	var v437 float64
	_ = v437
	var v447 float64
	_ = v447
	var v449 float64
	_ = v449
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 float64
	_ = v469
	var v470 int32
	_ = v470
	var v473 float64
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 float64
	_ = v480
	var v486 float64
	_ = v486
	var v489 float64
	_ = v489
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 float64
	_ = v499
	var v500 float64
	_ = v500
	var v509 float64
	_ = v509
	var v520 float64
	_ = v520
	var v525 float64
	_ = v525
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 float64
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 float64
	_ = v583
	var v584 float64
	_ = v584
	var v585 float64
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 float64
	_ = v588
	var v589 float64
	_ = v589
	var v591 int32
	_ = v591
	var v608 float64
	_ = v608
	var v612 float64
	_ = v612
	var v614 int32
	_ = v614
	var v620 float64
	_ = v620
	var v623 float64
	_ = v623
	var v624 float64
	_ = v624
	var v632 int32
	_ = v632
	var v638 float64
	_ = v638
	var v640 float64
	_ = v640
	var v641 float64
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 float64
	_ = v652
	var v659 int32
	_ = v659
	var v662 float64
	_ = v662
	var v665 float64
	_ = v665
	var v675 float64
	_ = v675
	var v680 int32
	_ = v680
	var v681 float64
	_ = v681
	var v684 float64
	_ = v684
	var v685 float64
	_ = v685
	var v686 int32
	_ = v686
	var v687 float64
	_ = v687
	var v689 int32
	_ = v689
	var v706 float64
	_ = v706
	var v710 float64
	_ = v710
	var v711 float64
	_ = v711
	var v716 float64
	_ = v716
	var v723 int32
	_ = v723
	var v727 float64
	_ = v727
	var v729 float64
	_ = v729
	var v730 float64
	_ = v730
	var v732 float64
	_ = v732
	var v736 float64
	_ = v736
	var v740 float64
	_ = v740
	var v750 float64
	_ = v750
	var v771 float64
	_ = v771
	var v781 float64
	_ = v781
	var v793 float64
	_ = v793
	v8 = float64(0)
	v21 = l4 - int32(1)
	v34 = int32(-1)
	v35 = v21
	goto L1
L1:
	;
	v44 = base.I32_div_s(v34+v35+int32(1), int32(2))
	v48 = F_range_cmp_bounds(m, l0, l3+v44<<(uint(int32(3))%32), l1)
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
		v793 = v8
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
	v54 = v34
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
	v57 = v35
	goto L10
L9:
	;
	v57 = v44 - int32(1)
	goto L10
L10:
	;
	if v54 < v57 {
		v34 = v54
		v35 = v57
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L2
L12:
	;
	return v793
L13:
	;
	v62 = l0 + int32(268)
	v64 = l4 - int32(2)
	if v54 < v64 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = v54
	goto L16
L15:
	;
	v66 = v64
	goto L16
L16:
	;
	v69 = l3 + v66<<(uint(int32(3))%32)
	v72 = F_get_position(m, l0, l1, v69, v69+int32(8))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v74 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v66 < int32(0) {
		v793 = v8
		goto L12
	} else {
		goto L38
	}
L19:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v79 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v102 != int32(1) {
		v112 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L18
	} else {
		goto L34
	}
L22:
	;
	v80 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L24
L23:
	;
	v80 = float64(1)
	goto L24
L24:
	;
	if v79 != 0 {
		v112 = v80
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v81 == int32(0) {
		v112 = v80
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v84 = float64(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v89 = F_FunctionCall2Coll(m, v62, v86, v87, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v89)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v91)&int64(9223372036854775807)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v97 = v84
	goto L30
L29:
	;
	v97 = v91
	goto L30
L30:
	;
	if base.F64_lt(v91, float64(0)) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v100 = v84
	goto L33
L32:
	;
	v100 = v97
	goto L33
L33:
	;
	v112 = v100
	goto L18
L34:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v107 == v108 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v110 = float64(0)
	goto L37
L36:
	;
	v110 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L37
L37:
	;
	v112 = v110
	goto L18
L38:
	;
	v116 = base.F64_convert_i32_u(v21)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+4)))
	if v117 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v163 = float64(0)
	if base.F64_lt(v154, v163) != 0 {
		v425 = v163
		goto L60
	} else {
		goto L61
	}
L40:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
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
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v145 != int32(1) {
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
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v132 = F_FunctionCall2Coll(m, v62, v129, v130, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v134 = *(*float64)(unsafe.Add(mBase, uint32(v132)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v134)&int64(9223372036854775807)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v140 = v127
	goto L51
L50:
	;
	v140 = v134
	goto L51
L51:
	;
	if base.F64_lt(v134, float64(0)) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v143 = v127
	goto L54
L53:
	;
	v143 = v140
	goto L54
L54:
	;
	v154 = v143
	goto L39
L55:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+6)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v150 == v151 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v153 = float64(0)
	goto L58
L57:
	;
	v153 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L58
L58:
	;
	v154 = v153
	goto L39
L59:
	;
	v437 = base.F64_add(base.F64_div(base.F64_mul(v72, base.F64_sub(float64(1), v425)), v116), float64(0))
	if v66 == int32(0) {
		v793 = v437
		goto L12
	} else {
		goto L135
	}
L60:
	;
	goto L59
L61:
	;
	v174 = float64(1)
	v179 = base.F64_abs(v154)
	if int32(0)&base.F64_eq(v179, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v425 = v174
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v185 = l6 - int32(1)
	if v185 < int32(0) {
		v425 = v174
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v189 = v185
	v193 = int32(-1)
	goto L64
L64:
	;
	v210 = int32(2)
	v211 = base.I32_div_s(v189+v193+int32(1), v210)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l5+v211<<(uint(v210)%32))))
	v216 = *(*float64)(unsafe.Add(mBase, uint32(v215)))
	if base.F64_gt(v112, v216) != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v185 <= v226 {
		v425 = v174
		goto L60
	} else {
		goto L77
	}
L66:
	;
	if v226 < v224 {
		v189 = v224
		v193 = v226
		goto L64
	} else {
		goto L76
	}
L67:
	;
	v224 = v189
	v226 = v211
	goto L66
L68:
	;
	goto L69
L69:
	;
	v221 = int32(0) & base.F64_ge(v112, v216)
	if v221 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v222 = v189
	goto L72
L71:
	;
	v222 = v211 - int32(1)
	goto L72
L72:
	;
	if v221 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v223 = v211
	goto L75
L74:
	;
	v223 = v193
	goto L75
L75:
	;
	v224 = v222
	v226 = v223
	goto L66
L76:
	;
	goto L65
L77:
	;
	if v226 < int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v277 = base.F64_convert_i32_s(v185)
	v278 = base.F64_div(base.F64_add(v274, base.F64_convert_i32_u(v268)), v277)
	if base.F64_eq(v112, v154) != 0 {
		v425 = v278
		goto L60
	} else {
		goto L95
	}
L79:
	;
	v268 = int32(0)
	v274 = float64(0)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v235 = l5 + v226<<(uint(int32(2))%32)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	v237 = *(*float64)(unsafe.Add(mBase, uint32(v236)))
	v238 = base.F64_abs(v237)
	v239 = math.Float64frombits(uint64(0x7ff0000000000000))
	v240 = base.F64_eq(v238, v239)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v242 = *(*float64)(unsafe.Add(mBase, uint32(v241)))
	v243 = base.F64_abs(v242)
	v245 = base.F64_eq(v243, v239)
	if v245 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if v240 != 0 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	if v240 != 0 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	if base.F64_eq(base.F64_abs(v112), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v268 = v226
		v274 = float64(0.5)
		goto L78
	} else {
		goto L85
	}
L85:
	;
	v268 = v226
	v274 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v237, v112), base.F64_sub(v237, v242)))
	goto L78
L86:
	;
	if base.F64_eq(v238, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	if v245 == int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v268 = v226
	v274 = float64(1)
	goto L78
L89:
	;
	v262 = float64(0)
	goto L91
L90:
	;
	v262 = float64(0.5)
	goto L91
L91:
	;
	if base.F64_eq(v243, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v266 = v262
	goto L94
L93:
	;
	v266 = float64(0.5)
	goto L94
L94:
	;
	v268 = v226
	v274 = v266
	goto L78
L95:
	;
	if v185 <= v268 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v390 = float64(0)
	v394 = base.F64_div(base.F64_add(v386, base.F64_convert_i32_u(v377)), v277)
	if base.F64_gt(v381, v390)|base.F64_gt(v394, v390) != 0 {
		goto L128
	} else {
		goto L129
	}
L97:
	;
	v377 = v268
	v381 = v278
	v383 = v112
	v384 = v163
	v386 = v163
	goto L96
L98:
	;
	goto L99
L99:
	;
	v286 = v268
	v292 = v278
	v294 = v163
	v295 = v112
	goto L101
L100:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l5+v286<<(uint(int32(2))%32))))
	v335 = *(*float64)(unsafe.Add(mBase, uint32(v334)))
	if base.F64_eq(v306, v335) != 0 {
		goto L111
	} else {
		goto L112
	}
L101:
	;
	v300 = int32(1)
	v301 = v286 + v300
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l5+v301<<(uint(int32(2))%32))))
	v306 = *(*float64)(unsafe.Add(mBase, uint32(v305)))
	if base.B2i32(base.F64_ge(v154, v306) == int32(0))|int32(1) != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v377 = v185
	v381 = v319
	v383 = v306
	v384 = v329
	v386 = v163
	goto L96
L103:
	;
	v313 = base.F64_lt(v306, v154)
	goto L105
L104:
	;
	v313 = v300
	goto L105
L105:
	;
	if v313 == int32(0) {
		goto L100
	} else {
		goto L106
	}
L106:
	;
	v316 = float64(0)
	v319 = base.F64_div(base.F64_convert_i32_u(v286), v277)
	if base.F64_gt(v292, v316)|base.F64_gt(v319, v316) != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v329 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v292, v319), float64(0.5)), base.F64_sub(v306, v295)), v294)
	goto L109
L108:
	;
	v329 = v294
	goto L109
L109:
	;
	if v185 != v301 {
		v286 = v301
		v292 = v319
		v294 = v329
		v295 = v306
		goto L101
	} else {
		goto L110
	}
L110:
	;
	goto L102
L111:
	;
	v370 = float64(0)
	goto L113
L112:
	;
	v338 = base.F64_abs(v306)
	v339 = math.Float64frombits(uint64(0x7ff0000000000000))
	v340 = base.F64_eq(v338, v339)
	v341 = base.F64_abs(v335)
	v343 = base.F64_eq(v341, v339)
	if v343 != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v377 = v286
	v381 = v292
	v383 = v295
	v384 = v294
	v386 = v370
	goto L96
L114:
	;
	v370 = v365
	goto L113
L115:
	;
	if v340 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	if v340 != 0 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	if base.F64_eq(base.F64_abs(v154), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v365 = float64(0.5)
		goto L114
	} else {
		goto L118
	}
L118:
	;
	v365 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v306, v154), base.F64_sub(v306, v335)))
	goto L114
L119:
	;
	if base.F64_eq(v338, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	if v343 == int32(0) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v365 = float64(1)
	goto L114
L122:
	;
	v360 = float64(0)
	goto L124
L123:
	;
	v360 = float64(0.5)
	goto L124
L124:
	;
	if base.F64_eq(v341, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v364 = v360
	goto L127
L126:
	;
	v364 = float64(0.5)
	goto L127
L127:
	;
	v365 = v364
	goto L114
L128:
	;
	v404 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v381, v394), float64(0.5)), base.F64_sub(v154, v383)), v384)
	goto L130
L129:
	;
	v404 = v384
	goto L130
L130:
	;
	if base.F64_eq(v179, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if base.F64_eq(base.F64_abs(v404), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v425 = float64(0.5)
		goto L60
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v425 = base.F64_div(v404, base.F64_sub(v154, v112))
	goto L60
L134:
	;
	goto L133
L135:
	;
	v447 = v154
	v449 = v437
	v454 = v66
	goto L136
L136:
	;
	v459 = v454 - int32(1)
	v462 = l3 + v459<<(uint(int32(3))%32)
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+4)))
	if v463 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	v793 = v781
	goto L12
L138:
	;
	v509 = float64(0)
	if base.F64_lt(v500, v509) != 0 {
		v771 = v509
		goto L159
	} else {
		goto L160
	}
L139:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v468 != 0 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	goto L141
L141:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v491 != int32(1) {
		v500 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L138
	} else {
		goto L154
	}
L142:
	;
	v469 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L144
L143:
	;
	v469 = float64(1)
	goto L144
L144:
	;
	if v468 != 0 {
		v500 = v469
		goto L138
	} else {
		goto L145
	}
L145:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v470 == int32(0) {
		v500 = v469
		goto L138
	} else {
		goto L146
	}
L146:
	;
	v473 = float64(1)
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v478 = F_FunctionCall2Coll(m, v62, v475, v476, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L3
	} else {
		goto L147
	}
L147:
	;
	v480 = *(*float64)(unsafe.Add(mBase, uint32(v478)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v480)&int64(9223372036854775807)) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v486 = v473
	goto L150
L149:
	;
	v486 = v480
	goto L150
L150:
	;
	if base.F64_lt(v480, float64(0)) != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v489 = v473
	goto L153
L152:
	;
	v489 = v486
	goto L153
L153:
	;
	v500 = v489
	goto L138
L154:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+6)))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v496 == v497 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v499 = float64(0)
	goto L157
L156:
	;
	v499 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L157
L157:
	;
	v500 = v499
	goto L138
L158:
	;
	v781 = base.F64_add(v449, base.F64_div(base.F64_sub(float64(1), v771), v116))
	if base.Ui32(int32(1)) < base.Ui32(v454) {
		v447 = v500
		v449 = v781
		v454 = v459
		goto L136
	} else {
		goto L234
	}
L159:
	;
	goto L158
L160:
	;
	v520 = float64(1)
	v525 = base.F64_abs(v500)
	if int32(0)&base.F64_eq(v525, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v771 = v520
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v531 = l6 - int32(1)
	if v531 < int32(0) {
		v771 = v520
		goto L159
	} else {
		goto L162
	}
L162:
	;
	v535 = v531
	v539 = int32(-1)
	goto L163
L163:
	;
	v556 = int32(2)
	v557 = base.I32_div_s(v535+v539+int32(1), v556)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l5+v557<<(uint(v556)%32))))
	v562 = *(*float64)(unsafe.Add(mBase, uint32(v561)))
	if base.F64_gt(v447, v562) != 0 {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	if v531 <= v572 {
		v771 = v520
		goto L159
	} else {
		goto L176
	}
L165:
	;
	if v572 < v570 {
		v535 = v570
		v539 = v572
		goto L163
	} else {
		goto L175
	}
L166:
	;
	v570 = v535
	v572 = v557
	goto L165
L167:
	;
	goto L168
L168:
	;
	v567 = int32(0) & base.F64_ge(v447, v562)
	if v567 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v568 = v535
	goto L171
L170:
	;
	v568 = v557 - int32(1)
	goto L171
L171:
	;
	if v567 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v569 = v557
	goto L174
L173:
	;
	v569 = v539
	goto L174
L174:
	;
	v570 = v568
	v572 = v569
	goto L165
L175:
	;
	goto L164
L176:
	;
	if v572 < int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v623 = base.F64_convert_i32_s(v531)
	v624 = base.F64_div(base.F64_add(v620, base.F64_convert_i32_u(v614)), v623)
	if base.F64_eq(v447, v500) != 0 {
		v771 = v624
		goto L159
	} else {
		goto L194
	}
L178:
	;
	v614 = int32(0)
	v620 = float64(0)
	goto L177
L179:
	;
	goto L180
L180:
	;
	v581 = l5 + v572<<(uint(int32(2))%32)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	v583 = *(*float64)(unsafe.Add(mBase, uint32(v582)))
	v584 = base.F64_abs(v583)
	v585 = math.Float64frombits(uint64(0x7ff0000000000000))
	v586 = base.F64_eq(v584, v585)
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	v588 = *(*float64)(unsafe.Add(mBase, uint32(v587)))
	v589 = base.F64_abs(v588)
	v591 = base.F64_eq(v589, v585)
	if v591 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	if v586 != 0 {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	if v586 != 0 {
		goto L181
	} else {
		goto L183
	}
L183:
	;
	if base.F64_eq(base.F64_abs(v447), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v614 = v572
		v620 = float64(0.5)
		goto L177
	} else {
		goto L184
	}
L184:
	;
	v614 = v572
	v620 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v583, v447), base.F64_sub(v583, v588)))
	goto L177
L185:
	;
	if base.F64_eq(v584, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	if v591 == int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v614 = v572
	v620 = float64(1)
	goto L177
L188:
	;
	v608 = float64(0)
	goto L190
L189:
	;
	v608 = float64(0.5)
	goto L190
L190:
	;
	if base.F64_eq(v589, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v612 = v608
	goto L193
L192:
	;
	v612 = float64(0.5)
	goto L193
L193:
	;
	v614 = v572
	v620 = v612
	goto L177
L194:
	;
	if v531 <= v614 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v736 = float64(0)
	v740 = base.F64_div(base.F64_add(v732, base.F64_convert_i32_u(v723)), v623)
	if base.F64_gt(v727, v736)|base.F64_gt(v740, v736) != 0 {
		goto L227
	} else {
		goto L228
	}
L196:
	;
	v723 = v614
	v727 = v624
	v729 = v447
	v730 = v509
	v732 = v509
	goto L195
L197:
	;
	goto L198
L198:
	;
	v632 = v614
	v638 = v624
	v640 = v509
	v641 = v447
	goto L200
L199:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l5+v632<<(uint(int32(2))%32))))
	v681 = *(*float64)(unsafe.Add(mBase, uint32(v680)))
	if base.F64_eq(v652, v681) != 0 {
		goto L210
	} else {
		goto L211
	}
L200:
	;
	v646 = int32(1)
	v647 = v632 + v646
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l5+v647<<(uint(int32(2))%32))))
	v652 = *(*float64)(unsafe.Add(mBase, uint32(v651)))
	if base.B2i32(base.F64_ge(v500, v652) == int32(0))|int32(1) != 0 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v723 = v531
	v727 = v665
	v729 = v652
	v730 = v675
	v732 = v509
	goto L195
L202:
	;
	v659 = base.F64_lt(v652, v500)
	goto L204
L203:
	;
	v659 = v646
	goto L204
L204:
	;
	if v659 == int32(0) {
		goto L199
	} else {
		goto L205
	}
L205:
	;
	v662 = float64(0)
	v665 = base.F64_div(base.F64_convert_i32_u(v632), v623)
	if base.F64_gt(v638, v662)|base.F64_gt(v665, v662) != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v675 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v638, v665), float64(0.5)), base.F64_sub(v652, v641)), v640)
	goto L208
L207:
	;
	v675 = v640
	goto L208
L208:
	;
	if v531 != v647 {
		v632 = v647
		v638 = v665
		v640 = v675
		v641 = v652
		goto L200
	} else {
		goto L209
	}
L209:
	;
	goto L201
L210:
	;
	v716 = float64(0)
	goto L212
L211:
	;
	v684 = base.F64_abs(v652)
	v685 = math.Float64frombits(uint64(0x7ff0000000000000))
	v686 = base.F64_eq(v684, v685)
	v687 = base.F64_abs(v681)
	v689 = base.F64_eq(v687, v685)
	if v689 != 0 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	v723 = v632
	v727 = v638
	v729 = v641
	v730 = v640
	v732 = v716
	goto L195
L213:
	;
	v716 = v711
	goto L212
L214:
	;
	if v686 != 0 {
		goto L218
	} else {
		goto L219
	}
L215:
	;
	if v686 != 0 {
		goto L214
	} else {
		goto L216
	}
L216:
	;
	if base.F64_eq(base.F64_abs(v500), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v711 = float64(0.5)
		goto L213
	} else {
		goto L217
	}
L217:
	;
	v711 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v652, v500), base.F64_sub(v652, v681)))
	goto L213
L218:
	;
	if base.F64_eq(v684, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	if v689 == int32(0) {
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v711 = float64(1)
	goto L213
L221:
	;
	v706 = float64(0)
	goto L223
L222:
	;
	v706 = float64(0.5)
	goto L223
L223:
	;
	if base.F64_eq(v687, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v710 = v706
	goto L226
L225:
	;
	v710 = float64(0.5)
	goto L226
L226:
	;
	v711 = v710
	goto L213
L227:
	;
	v750 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v727, v740), float64(0.5)), base.F64_sub(v500, v729)), v730)
	goto L229
L228:
	;
	v750 = v730
	goto L229
L229:
	;
	if base.F64_eq(v525, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	if base.F64_eq(base.F64_abs(v750), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v771 = float64(0.5)
		goto L159
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v771 = base.F64_div(v750, base.F64_sub(v500, v447))
	goto L159
L233:
	;
	goto L232
L234:
	;
	goto L137
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
	var v56 float64
	_ = v56
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
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
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
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
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
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
	var v221 int32
	_ = v221
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
	var v275 int32
	_ = v275
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
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
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
	var v465 float64
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v479 float64
	_ = v479
	var v482 int32
	_ = v482
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
	var v551 int32
	_ = v551
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
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
	var v629 int32
	_ = v629
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
	var v651 float64
	_ = v651
	var v652 int32
	_ = v652
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v681 float64
	_ = v681
	var v682 int32
	_ = v682
	var v684 float64
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v698 float64
	_ = v698
	var v700 float64
	_ = v700
	var v719 float64
	_ = v719
	var v721 float64
	_ = v721
	var v725 float64
	_ = v725
	var v727 float64
	_ = v727
	var v729 float64
	_ = v729
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v755 float64
	_ = v755
	var v756 float64
	_ = v756
	var v764 float64
	_ = v764
	var v768 float64
	_ = v768
	v12 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v33 == v12 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if int32(1)<<(uint(v32)%32)&int32(174) != 0 {
		goto L123
	} else {
		goto L124
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
	v56 = float64(1)
	v59 = l7
	v72 = v12
	goto L8
L8:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v72<<(uint(int32(2))%32))))
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
		goto L118
	} else {
		goto L119
	}
L10:
	;
	v499 = v72 + int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v499 < v500 {
		v56 = v479
		v59 = v482
		v72 = v499
		goto L8
	} else {
		goto L117
	}
L11:
	;
	if l7 == v59 {
		goto L45
	} else {
		goto L46
	}
L12:
	;
	if v32&int32(-2) != int32(4) {
		goto L25
	} else {
		goto L26
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
		v479 = v56
		v482 = v59
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
	if v94 == int32(0) {
		v479 = v56
		v482 = v59
		goto L10
	} else {
		goto L23
	}
L23:
	;
	if v47 == int32(0) {
		v479 = v56
		v482 = v59
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v152 = int32(0)
	goto L11
L25:
	;
	v152 = int32(0)
	goto L11
L26:
	;
	goto L27
L27:
	;
	if v42 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v149 != int32(1) {
		v479 = v56
		v482 = v59
		goto L10
	} else {
		goto L44
	}
L29:
	;
	v149 = int32(0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v111 = int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v112 <= v111 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v115 = v111
	goto L34
L33:
	;
	v115 = v112
	goto L34
L34:
	;
	v118 = int32(0)
	v120 = v118
	v121 = v118
	goto L35
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(8)+v120<<(uint(int32(2))%32))))
	if v129 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v149 = v142
	goto L28
L37:
	;
	goto L36
L38:
	;
	v130 = int32(2)
	if v121 != 0 {
		v142 = v130
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v135 = v121
	goto L40
L40:
	;
	v138 = v120 + int32(1)
	if v138 != v115 {
		v120 = v138
		v121 = v135
		goto L35
	} else {
		goto L43
	}
L41:
	;
	v131 = int32(1)
	if base.Ui32(v131) < base.Ui32(base.I32_popcnt(v129)) {
		v142 = v130
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v135 = v131
	goto L40
L43:
	;
	v142 = v135
	goto L37
L44:
	;
	v152 = int32(1)
	goto L11
L45:
	;
	v154 = F_list_copy(m, v59)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L13
	} else {
		goto L48
	}
L46:
	;
	v156 = v59
	goto L47
L47:
	;
	if v156 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v156 = v154
	goto L47
L49:
	;
	v159 = int32(0)
	v161 = F_list_concat(m, v159, v159)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v166 = v79 + int32(288)
	v167 = int32(0)
	v180 = v156
	v185 = v167
	v186 = v167
	goto L53
L52:
	;
	v479 = v56
	v482 = v161
	goto L10
L53:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v185 < v196 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v367 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L55:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if int32(0) < v198 {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v361 = v180
	v367 = v186
	goto L57
L57:
	;
	goto L54
L58:
	;
	if v332 != 0 {
		v180 = v332
		v185 = v324 + int32(1)
		v186 = v338
		goto L53
	} else {
		goto L87
	}
L59:
	;
	v317 = F_list_delete_nth_cell(m, v180, v185)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L13
	} else {
		goto L85
	}
L60:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201+v185<<(uint(int32(2))%32))))
	v210 = int32(0)
	v221 = v198
	goto L63
L61:
	;
	goto L62
L62:
	;
	v324 = v185
	v332 = v180
	v338 = v186
	goto L58
L63:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v205)+60))
	if v234 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L62
L65:
	;
	v286 = v210 + int32(1)
	if v286 < v284 {
		v210 = v286
		v221 = v284
		goto L63
	} else {
		goto L84
	}
L66:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v166+v210<<(uint(int32(2))%32))))
	if v238 != v234 {
		v284 = v221
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v79+int32(544)+v210<<(uint(int32(2))%32))))
	v244 = int32(0)
	if v243 == v244 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L59
L70:
	;
	if v282 != 0 {
		goto L59
	} else {
		goto L83
	}
L71:
	;
	v282 = int32(0)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if v250 <= int32(0) {
		v275 = v244
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v282 = v275
	goto L70
L75:
	;
	v253 = int32(0)
	if v253 < v250 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v256 = v250
	goto L78
L77:
	;
	v256 = v253
	goto L78
L78:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	v259 = int32(0)
	goto L79
L79:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v257+v259<<(uint(int32(2))%32))))
	v268 = base.B2i32(v267 == v205)
	if v267 == v205 {
		v275 = v268
		goto L74
	} else {
		goto L81
	}
L80:
	;
	v275 = v268
	goto L74
L81:
	;
	v270 = v259 + int32(1)
	if v270 != v256 {
		v259 = v270
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v284 = v283
	goto L65
L84:
	;
	goto L64
L85:
	;
	v319 = F_lappend(m, v186, v205)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L13
	} else {
		goto L86
	}
L86:
	;
	v324 = v185 - int32(1)
	v332 = v317
	v338 = v319
	goto L58
L87:
	;
	v361 = v332
	v367 = v338
	goto L57
L88:
	;
	v380 = F_list_concat(m, v361, int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L13
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v79)+284))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v79)+272))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v79)+276))
	if v382 != v383+(v384-v385) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v479 = v56
	v482 = v380
	goto L10
L92:
	;
	v389 = F_list_concat(m, v361, v367)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L13
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v392 = F_find_base_rel(m, l0, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L13
	} else {
		goto L96
	}
L95:
	;
	v479 = v56
	v482 = v389
	goto L10
L96:
	;
	v394 = *(*float64)(unsafe.Add(mBase, uint32(v392)+120))
	v395 = float64(1)
	if base.F64_gt(v394, v395) != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v398 = v394
	goto L99
L98:
	;
	v398 = v395
	goto L99
L99:
	;
	if v152 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v392)+16))
	v401 = v399
	goto L102
L101:
	;
	v401 = float64(1)
	goto L102
L102:
	;
	v403 = base.F64_mul(v56, base.F64_div(v401, v398))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v79)+276))
	if v404 <= int32(0) {
		v479 = v403
		v482 = v361
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v407 <= int32(0) {
		v479 = v403
		v482 = v361
		goto L10
	} else {
		goto L104
	}
L104:
	;
	v416 = int32(0)
	v421 = v403
	goto L105
L105:
	;
	v441 = v416 << (uint(int32(2)) % 32)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v166+v441)))
	if v443 == int32(0) {
		v465 = v421
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v479 = v465
	v482 = v361
	goto L10
L107:
	;
	v468 = v416 + int32(1)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v468 < v469 {
		v416 = v468
		v421 = v465
		goto L105
	} else {
		goto L116
	}
L108:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+40)))
	if v446 != int32(1) {
		v465 = v421
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v441+(v79+int32(416)))))
	v451 = int32(0)
	v453 = F_ec_search_derived_clause_for_ems(m, l0, v443, v450, v451, v451)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L13
	} else {
		goto L110
	}
L110:
	;
	if v453 == int32(0) {
		v465 = v421
		goto L107
	} else {
		goto L111
	}
L111:
	;
	v458 = F_clause_selectivity(m, l0, v453, int32(0), v32, l6)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	if base.F64_gt(v458, float64(0)) != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v463 = base.F64_div(v421, v458)
	goto L115
L114:
	;
	v463 = v421
	goto L115
L115:
	;
	v465 = v463
	goto L107
L116:
	;
	goto L106
L117:
	;
	goto L9
L118:
	;
	v516 = v482
	v518 = v508
	goto L1
L119:
	;
	v505 = float64(1)
	if base.F64_gt(v479, v505) != 0 {
		v508 = v505
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v516 = v482
	v518 = v479
	goto L1
L121:
	;
	switch v32 {
	case 0:
		goto L159
	case 1:
		goto L164
	case 2:
		goto L163
	default:
		goto L160
	case 4:
		goto L162
	case 5:
		goto L161
	}
L122:
	;
	v681 = F_clauselist_selectivity(m, l0, v667, int32(0), v32, l6)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L13
	} else {
		goto L154
	}
L123:
	;
	v540 = int32(0)
	if v516 == v540 {
		v666 = v540
		v667 = v540
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v651 = F_clauselist_selectivity(m, l0, v516, int32(0), v32, l6)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L13
	} else {
		goto L153
	}
L126:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v544 <= int32(0) {
		v666 = v540
		v667 = v540
		goto L122
	} else {
		goto L127
	}
L127:
	;
	v551 = int32(0)
	v561 = v540
	v562 = v540
	goto L128
L128:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v575+v551<<(uint(int32(2))%32))))
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+8)))
	if v580 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	v666 = v643
	v667 = v644
	goto L122
L130:
	;
	v646 = v551 + int32(1)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v646 < v647 {
		v551 = v646
		v561 = v643
		v562 = v644
		goto L128
	} else {
		goto L152
	}
L131:
	;
	v641 = F_lappend(m, v562, v579)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L13
	} else {
		goto L151
	}
L132:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v579)+32))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v585 = int32(0)
	if v583 == v585 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	goto L134
L134:
	;
	v639 = F_lappend(m, v561, v579)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L13
	} else {
		goto L150
	}
L135:
	;
	if v638 != 0 {
		goto L131
	} else {
		goto L149
	}
L136:
	;
	v638 = int32(1)
	goto L135
L137:
	;
	goto L138
L138:
	;
	if v584 == int32(0) {
		v629 = v585
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v638 = v629
	goto L135
L140:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v583)+4))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	if v595 < v594 {
		v629 = v585
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v597 = int32(1)
	if v594 <= v597 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v600 = v597
	goto L144
L143:
	;
	v600 = v594
	goto L144
L144:
	;
	v601 = int32(8)
	v606 = int32(0)
	goto L145
L145:
	;
	v613 = v606 << (uint(int32(2)) % 32)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v583+v601+v613)))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v613+(v584+v601))))
	v620 = v615 & (v617 ^ int32(-1))
	v622 = base.B2i32(v620 == int32(0))
	if v620 != 0 {
		v629 = v622
		goto L139
	} else {
		goto L147
	}
L146:
	;
	v629 = v622
	goto L139
L147:
	;
	v624 = v606 + int32(1)
	if v624 != v600 {
		v606 = v624
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	goto L134
L150:
	;
	v643 = v639
	v644 = v562
	goto L130
L151:
	;
	v643 = v561
	v644 = v641
	goto L130
L152:
	;
	goto L129
L153:
	;
	v698 = v651
	v700 = float64(0)
	goto L121
L154:
	;
	v684 = F_clauselist_selectivity(m, l0, v666, int32(0), v32, l6)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L13
	} else {
		goto L155
	}
L155:
	;
	F_list_free(m, v667)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L13
	} else {
		goto L156
	}
L156:
	;
	F_list_free(m, v666)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L13
	} else {
		goto L157
	}
L157:
	;
	v698 = v681
	v700 = v684
	goto L121
L158:
	;
	v756 = float64(1e+100)
	if base.F64_gt(v755, v756) != 0 {
		v768 = v756
		goto L177
	} else {
		goto L178
	}
L159:
	;
	v755 = base.F64_mul(base.F64_mul(base.F64_mul(l4, l5), v518), v698)
	goto L158
L160:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L13
	} else {
		goto L174
	}
L161:
	;
	v755 = base.F64_mul(v700, base.F64_mul(l4, base.F64_sub(float64(1), base.F64_mul(v518, v698))))
	goto L158
L162:
	;
	v755 = base.F64_mul(base.F64_mul(l4, v518), v698)
	goto L158
L163:
	;
	v725 = base.F64_mul(base.F64_mul(base.F64_mul(l4, l5), v518), v698)
	if base.F64_gt(l4, v725) != 0 {
		goto L168
	} else {
		goto L169
	}
L164:
	;
	v719 = base.F64_mul(base.F64_mul(base.F64_mul(l4, l5), v518), v698)
	if base.F64_gt(l4, v719) != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v721 = l4
	goto L167
L166:
	;
	v721 = v719
	goto L167
L167:
	;
	v755 = base.F64_mul(v700, v721)
	goto L158
L168:
	;
	v727 = l4
	goto L170
L169:
	;
	v727 = v725
	goto L170
L170:
	;
	if base.F64_gt(l5, v727) != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v729 = l5
	goto L173
L172:
	;
	v729 = v727
	goto L173
L173:
	;
	v755 = base.F64_mul(v700, v729)
	goto L158
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v32
	F_errmsg_internal(m, int32(_a_F_calc_joinrel_size_estimate_0), v30)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L13
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(_a_F_calc_joinrel_size_estimate_1), int32(_a_F_calc_joinrel_size_estimate_2), int32(_a_F_calc_joinrel_size_estimate_3))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L13
	} else {
		goto L176
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	m.G0 = v30 + int32(16)
	return v768
L178:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v755)&int64(9223372036854775807)) {
		v768 = v756
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v764 = float64(1)
	if base.F64_le(v755, v764) != 0 {
		v768 = v764
		goto L177
	} else {
		goto L180
	}
L180:
	;
	v768 = base.F64_nearest(v755)
	goto L177
}
func F_calc_rank_cd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 float32
	_ = v31
	var v37 float64
	_ = v37
	var v40 float32
	_ = v40
	var v46 float64
	_ = v46
	var v50 float64
	_ = v50
	var v54 float32
	_ = v54
	var v60 float64
	_ = v60
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v70 float32
	_ = v70
	var v76 float64
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var __phi573 int32
	_ = __phi573
	var v575 int32
	_ = v575
	var __phi575 int32
	_ = __phi575
	var v577 int32
	_ = v577
	var __phi577 int32
	_ = __phi577
	var v580 int32
	_ = v580
	var __phi580 int32
	_ = __phi580
	var v582 int32
	_ = v582
	var __phi582 int32
	_ = __phi582
	var v583 int32
	_ = v583
	var __phi583 int32
	_ = __phi583
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 float64
	_ = v673
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v693 float64
	_ = v693
	var v694 float64
	_ = v694
	var v695 float64
	_ = v695
	var v711 int32
	_ = v711
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v815 int32
	_ = v815
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v922 int32
	_ = v922
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v957 int32
	_ = v957
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1069 int32
	_ = v1069
	var v1075 int32
	_ = v1075
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1162 int32
	_ = v1162
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1197 int32
	_ = v1197
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1243 int32
	_ = v1243
	var v1244 float64
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1265 float64
	_ = v1265
	var v1272 int32
	_ = v1272
	var v1278 float64
	_ = v1278
	var v1279 float64
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1302 float64
	_ = v1302
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1328 float64
	_ = v1328
	var v1338 float64
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1378 int32
	_ = v1378
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1465 float64
	_ = v1465
	var v1466 float64
	_ = v1466
	var v1486 float64
	_ = v1486
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1513 int32
	_ = v1513
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1575 float64
	_ = v1575
	var v1594 float64
	_ = v1594
	var v1599 int32
	_ = v1599
	var v1605 float64
	_ = v1605
	var v1610 int32
	_ = v1610
	var v1616 float64
	_ = v1616
	var v1621 float64
	_ = v1621
	var v1627 float64
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1658 float32
	_ = v1658
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1681 int32
	_ = v1681
	v5 = int32(0)
	v25 = m.G0
	v27 = v25 + int32(-64)
	m.G0 = v27
	v31 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	if base.F32_ge(v31, float32(0)) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L20
	} else {
		goto L245
	}
L2:
	;
	if base.F32_gt(v31, float32(1)) != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v37 = float64(0.10000000149011612)
	goto L4
L4:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v27)+16)) = base.F64_div(float64(1), v37)
	v40 = *(*float32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.F32_ge(v40, float32(0)) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v37 = base.F64_promote_f32(v31)
	goto L4
L6:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v27)+24)) = base.F64_div(float64(1), v50)
	v54 = *(*float32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.F32_ge(v54, float32(0)) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v50 = float64(0.20000000298023224)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v46 = base.F64_promote_f32(v40)
	*(*float64)(unsafe.Add(mBase, uint32(v27)+24)) = v46
	if base.F32_gt(v40, float32(1)) != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v50 = v46
	goto L6
L11:
	;
	v65 = float64(1)
	*(*float64)(unsafe.Add(mBase, uint32(v27)+32)) = base.F64_div(v65, v64)
	v70 = *(*float32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.F32_ge(v70, float32(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v64 = float64(0.4000000059604645)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v60 = base.F64_promote_f32(v54)
	*(*float64)(unsafe.Add(mBase, uint32(v27)+32)) = v60
	if base.F32_gt(v54, float32(1)) != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v64 = v60
	goto L11
L16:
	;
	if base.F32_gt(v70, float32(1)) != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v76 = v65
	goto L18
L18:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v27)+40)) = base.F64_div(v65, v76)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = l2
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v83 = F_palloc0(m, v80*int32(_a_F_calc_rank_cd_0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v76 = base.F64_promote_f32(v70)
	goto L18
L20:
	;
	return float32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v83
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v91 = F_palloc(m, v88*int32(48))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v93 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	m.G0 = v27 - int32(-64)
	return v1658
L24:
	;
	F_pg_qsort(m, v498, v496, int32(12), int32(1532))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L20
	} else {
		goto L94
	}
L25:
	;
	v96 = int32(8)
	v99 = l1 + v96
	v104 = l2
	v106 = v88 << (uint(int32(2)) % 32)
	v108 = v5
	v110 = v91
	v114 = v5
	goto L28
L26:
	;
	v522 = v83
	v530 = v91
	goto L27
L27:
	;
	F_pfree(m, v530)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L20
	} else {
		goto L92
	}
L28:
	;
	v128 = l2 + v96 + v114*int32(12)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v129 != int32(1) {
		v494 = v106
		v496 = v108
		v498 = v110
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if int32(0) < v496 {
		goto L24
	} else {
		goto L91
	}
L30:
	;
	v515 = v114 + int32(1)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v515 < v517 {
		v104 = v516
		v106 = v494
		v108 = v496
		v110 = v498
		v114 = v515
		goto L28
	} else {
		goto L90
	}
L31:
	;
	v133 = v25 + int32(-4)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = int32(0)
	v144 = l1 + int32(8)
	v147 = v144 + v140<<(uint(int32(2))%32)
	if base.Ui32(v147) <= base.Ui32(v144) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if v286 == int32(0) {
		v494 = v106
		v496 = v108
		v498 = v110
		goto L30
	} else {
		goto L62
	}
L33:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+2)))
	if v214 != int32(1) {
		goto L49
	} else {
		goto L50
	}
L34:
	;
	v208 = v144
	v209 = v147
	v210 = v147
	goto L33
L35:
	;
	goto L36
L36:
	;
	v155 = v144
	v157 = v147
	goto L37
L37:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v162 = int32(12)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v172 = int32(2)
	v179 = base.I32_div_s((v157-v155)>>(uint(v172)%32), v172)
	v182 = v155 + v179<<(uint(v172)%32)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v191 = int32(0)
	v192 = F_tsCompareString(m, v104+int32(8)+v161*v162+int32(base.Ui32(v165)>>(uint(v162)%32)), v165&int32(4095), v144+v171<<(uint(v172)%32)+int32(base.Ui32(v183)>>(uint(v162)%32)), int32(base.Ui32(v183)>>(uint(int32(1))%32))&int32(2047), v191)
	mBase = m.M
	if v192 == v191 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v208 = v201
	v209 = v182
	v210 = v202
	goto L33
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = int32(1)
	v208 = v155
	v209 = v182
	v210 = v182
	goto L33
L40:
	;
	goto L41
L41:
	;
	v200 = base.B2i32(int32(0) < v192)
	if int32(0) < v192 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v201 = v182 + int32(4)
	goto L44
L43:
	;
	v201 = v155
	goto L44
L44:
	;
	if int32(0) < v192 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v202 = v157
	goto L47
L46:
	;
	v202 = v182
	goto L47
L47:
	;
	if base.Ui32(v201) < base.Ui32(v202) {
		v155 = v201
		v157 = v202
		goto L37
	} else {
		goto L48
	}
L48:
	;
	goto L38
L49:
	;
	v282 = int32(0)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v282 < v283 {
		goto L59
	} else {
		goto L60
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = int32(0)
	if base.Ui32(v208) < base.Ui32(v210) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v220 = v209
	goto L53
L52:
	;
	v220 = v210
	goto L53
L53:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v144+v221<<(uint(int32(2))%32)) <= base.Ui32(v220) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v232 = v220
	v233 = v221
	goto L55
L55:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v239 = int32(12)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v255 = int32(1)
	v260 = F_tsCompareString(m, v104+int32(8)+v238*v239+int32(base.Ui32(v242)>>(uint(v239)%32)), v242&int32(4095), v144+v233<<(uint(int32(2))%32)+int32(base.Ui32(v251)>>(uint(v239)%32)), int32(base.Ui32(v251)>>(uint(v255)%32))&int32(2047), v255)
	mBase = m.M
	if v260 != 0 {
		goto L49
	} else {
		goto L57
	}
L56:
	;
	goto L49
L57:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v261 + int32(1)
	v266 = v232 + int32(4)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v266) < base.Ui32(v144+v267<<(uint(int32(2))%32)) {
		v232 = v266
		v233 = v267
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v286 = v210
	goto L61
L60:
	;
	v286 = v282
	goto L61
L61:
	;
	goto L32
L62:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	if v289 <= int32(0) {
		v494 = v106
		v496 = v108
		v498 = v110
		goto L30
	} else {
		goto L63
	}
L63:
	;
	v292 = v289
	v294 = v286
	v296 = v106
	v298 = v108
	v300 = v110
	goto L64
L64:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	if v316&int32(1) != 0 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v494 = v379
	v496 = v466
	v498 = v383
	goto L30
L66:
	;
	if v335 != 0 {
		goto L77
	} else {
		goto L78
	}
L67:
	;
	v348 = v296
	v352 = v300
	goto L73
L68:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v323 = int32(1)
	v334 = v99 + v319<<(uint(int32(2))%32) + (int32(base.Ui32(v316)>>(uint(v323)%32))&int32(2047)+int32(base.Ui32(v316)>>(uint(int32(12))%32))+v323)&int32(_a_F_calc_rank_cd_1)
	v335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334))))
	v336 = v298 + v335
	if v296 <= v336 {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v339 = v294 + int32(4)
	if (v339-v286)>>(uint(int32(2))%32) < v292 {
		v294 = v339
		goto L64
	} else {
		goto L72
	}
L71:
	;
	v379 = v296
	v383 = v300
	goto L66
L72:
	;
	v494 = v296
	v496 = v298
	v498 = v300
	goto L30
L73:
	;
	v370 = F_repalloc(m, v352, v348*int32(24))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L20
	} else {
		goto L75
	}
L74:
	;
	v379 = v373
	v383 = v370
	goto L66
L75:
	;
	v373 = v348 << (uint(int32(1)) % 32)
	if v373 <= v336 {
		v348 = v373
		v352 = v370
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v400 = v334 + int32(2)
	v402 = int32(0)
	v408 = v298
	goto L80
L78:
	;
	v460 = v292
	v466 = v298
	goto L79
L79:
	;
	v485 = v294 + int32(4)
	if (v485-v286)>>(uint(int32(2))%32) < v460 {
		v292 = v460
		v294 = v485
		v296 = v379
		v298 = v466
		v300 = v383
		goto L64
	} else {
		goto L89
	}
L80:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	if v426 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	v460 = v459
	v466 = v453
	goto L79
L82:
	;
	v457 = v402 + int32(1)
	if v457 != v335 {
		v402 = v457
		v408 = v453
		goto L80
	} else {
		goto L88
	}
L83:
	;
	v447 = v383 + v408*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v447)+4)) = v294
	*(*uint16)(unsafe.Add(mBase, uint32(v447)+8)) = uint16(v444)
	*(*int32)(unsafe.Add(mBase, uint32(v447))) = v128
	v453 = v408 + int32(1)
	goto L82
L84:
	;
	v432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400+v402<<(uint(int32(1))%32)))))
	v444 = v432
	goto L83
L85:
	;
	goto L86
L86:
	;
	v433 = int32(1)
	v436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400+v402<<(uint(v433)%32)))))
	if int32(base.Ui32(v426)>>(uint(int32(base.Ui32(v436)>>(uint(int32(14))%32)))%32))&v433 == int32(0) {
		v453 = v408
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v444 = v436
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
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v522 = v521
	v530 = v498
	goto L27
L92:
	;
	F_pfree(m, v522)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L20
	} else {
		goto L93
	}
L93:
	;
	v1658 = float32(0)
	goto L23
L94:
	;
	v555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v498)+8)))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v559 = F_palloc(m, v556<<(uint(int32(2))%32))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L20
	} else {
		goto L95
	}
L95:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	*(*int32)(unsafe.Add(mBase, uint32(v559))) = v561
	if v496 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v671 = int32(0)
	v673 = float64(0)
	v687 = v671
	v689 = v671
	v693 = v673
	v694 = float64(0)
	v695 = v673
	goto L108
L97:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v498)+8)) = uint16(v555)
	v566 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v498)+4)) = uint16(v566)
	*(*int32)(unsafe.Add(mBase, uint32(v498))) = v559
	v670 = v566
	goto L96
L98:
	;
	goto L99
L99:
	;
	__phi573 = v498
	__phi575 = v498
	__phi577 = v498 + int32(12)
	__phi580 = int32(1)
	__phi582 = v559
	__phi583 = v555
	v573 = __phi573
	v575 = __phi575
	v577 = __phi577
	v580 = __phi580
	v582 = __phi582
	v583 = __phi583
	goto L100
L100:
	;
	v598 = v575 + int32(20)
	v599 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598))))
	v600 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575)+8)))
	if v599 != v600 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v628)+8)) = uint16(v630)
	*(*uint16)(unsafe.Add(mBase, uint32(v628)+4)) = uint16(v631)
	*(*int32)(unsafe.Add(mBase, uint32(v628))) = v629
	v642 = int32(12)
	v645 = base.I32_div_s(v628-v498+v642, v642)
	v670 = v645
	goto L96
L102:
	;
	v632 = int32(12)
	v633 = v577 + v632
	v636 = base.I32_div_s(v633-v498, v632)
	if v636 < v496 {
		__phi573 = v628
		__phi575 = v577
		__phi577 = v633
		__phi580 = v631
		__phi582 = v629
		__phi583 = v630
		v573 = __phi573
		v575 = __phi575
		v577 = __phi577
		v580 = __phi580
		v582 = __phi582
		v583 = __phi583
		goto L100
	} else {
		goto L107
	}
L103:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v573)+8)) = uint16(v583)
	*(*uint16)(unsafe.Add(mBase, uint32(v573)+4)) = uint16(v580)
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = v582
	v616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598))))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v617)+4))
	v621 = F_palloc(m, v618<<(uint(int32(2))%32))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L20
	} else {
		goto L106
	}
L104:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v575)+16))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	if v602 != v603 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v575)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v582+base.I32_extend16_s(v580)<<(uint(int32(2))%32)))) = v609
	v628 = v573
	v629 = v582
	v630 = v583
	v631 = v580 + int32(1)
	goto L102
L106:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v575)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v621))) = v623
	v628 = v573 + int32(12)
	v629 = v621
	v630 = v616
	v631 = int32(1)
	goto L102
L107:
	;
	goto L101
L108:
	;
	v711 = v687
	goto L110
L109:
	;
	if l3&int32(1) == int32(0) {
		v1486 = v693
		goto L202
	} else {
		goto L203
	}
L110:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L20
	} else {
		goto L112
	}
L111:
	;
	goto L109
L112:
	;
	v725 = int32(0)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v726)+4))
	if v725 < v727 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v732 = v725
	goto L116
L114:
	;
	goto L115
L115:
	;
	v797 = int32(12)
	v798 = v711 * v797
	v800 = base.I32_div_s(v798, v797)
	if v670 <= v800 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v755 = v732 * int32(_a_F_calc_rank_cd_0)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v758 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v755+v756))) = uint8(v758)
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v760+v755)+1)) = uint8(v758)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v764+v755)+4)) = v758
	v769 = v732 + int32(1)
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v770)+4))
	if v769 < v771 {
		v732 = v769
		goto L116
	} else {
		goto L118
	}
L117:
	;
	goto L115
L118:
	;
	goto L117
L119:
	;
	goto L111
L120:
	;
	v804 = v498 + v711*int32(12)
	v805 = v804
	v815 = v798
	goto L121
L121:
	;
	v829 = int32(*(*int16)(unsafe.Add(mBase, uint32(v805)+4)))
	if int32(0) < v829 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L119
L123:
	;
	v835 = int32(0)
	goto L126
L124:
	;
	goto L125
L125:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v964 = F_TS_execute(m, v957+int32(8), v25+int32(-56), int32(0), int32(1533))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L20
	} else {
		goto L145
	}
L126:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v805)))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v857+v835<<(uint(int32(2))%32))))
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861))))
	if v862 != int32(1) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	goto L125
L128:
	;
	v930 = v835 + int32(1)
	v931 = int32(*(*int16)(unsafe.Add(mBase, uint32(v805)+4)))
	if v930 < v931 {
		v835 = v930
		goto L126
	} else {
		goto L144
	}
L129:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v871 = base.I32_div_s(v861-v866-int32(8), int32(12))
	v874 = v865 + v871*int32(_a_F_calc_rank_cd_0)
	v875 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v874))) = uint8(v875)
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874)+1)))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v874)+4))
	if v878 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v874)+4)) = v922
	goto L128
L131:
	;
	if v877&int32(1) != 0 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	v890 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v805)+8)))
	v892 = v874 + int32(8)
	v895 = int32(1)
	v898 = v877 & v895
	if v898 != 0 {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v885 = int32(_a_F_calc_rank_cd_2)
	goto L136
L135:
	;
	v885 = int32(0)
	goto L136
L136:
	;
	v887 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v805)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v874+v885)+8)) = uint16(v887)
	v922 = int32(1)
	goto L130
L137:
	;
	v899 = int32(_a_F_calc_rank_cd_3) - v878
	goto L139
L138:
	;
	v899 = v878 - v895
	goto L139
L139:
	;
	v903 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v892+v899<<(uint(int32(1))%32)))))
	if (v890^v903)&int32(_a_F_calc_rank_cd_4) == int32(0) {
		goto L128
	} else {
		goto L140
	}
L140:
	;
	if v898 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v911 = int32(_a_F_calc_rank_cd_4) - v878
	goto L143
L142:
	;
	v911 = v878
	goto L143
L143:
	;
	v912 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v892+v911<<(uint(v912)%32)))) = uint16(v890)
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v874)+4))
	v922 = v916 + v912
	goto L130
L144:
	;
	goto L127
L145:
	;
	if v964 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v966 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v805)+8)))
	v968 = v966 & int32(_a_F_calc_rank_cd_4)
	if v968 == int32(0) {
		goto L119
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v1344 = int32(12)
	v1345 = v805 + v1344
	v1346 = v1345 - v498
	v1348 = base.I32_div_s(v1346, v1344)
	if v1348 < v670 {
		v805 = v1345
		v815 = v1346
		goto L121
	} else {
		goto L201
	}
L149:
	;
	v971 = int32(0)
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v972)+4))
	if v971 < v973 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v978 = v971
	goto L153
L151:
	;
	goto L152
L152:
	;
	v1043 = v498 + v815
	if base.Ui32(v1043) < base.Ui32(v804) {
		goto L157
	} else {
		goto L158
	}
L153:
	;
	v1001 = v978 * int32(_a_F_calc_rank_cd_0)
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v1004 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1001+v1002))) = uint8(v1004)
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v1008 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1006+v1001)+1)) = uint8(v1008)
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1010+v1001)+4)) = v1004
	v1015 = v978 + v1008
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+4))
	if v1015 < v1017 {
		v978 = v1015
		goto L153
	} else {
		goto L155
	}
L154:
	;
	goto L152
L155:
	;
	goto L154
L156:
	;
	v1243 = base.I32_div_s(v1049-v498, int32(12))
	v1244 = float64(0)
	if base.Ui32(v1049) <= base.Ui32(v805) {
		goto L189
	} else {
		goto L190
	}
L157:
	;
	v711 = v711 + int32(1)
	goto L110
L158:
	;
	v1049 = v1043
	goto L159
L159:
	;
	v1069 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1049)+4)))
	if int32(0) < v1069 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v1211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1049)+8)))
	v1213 = v1211 & int32(_a_F_calc_rank_cd_4)
	if base.Ui32(v1213) <= base.Ui32(v968) {
		goto L156
	} else {
		goto L188
	}
L161:
	;
	v1075 = int32(0)
	goto L164
L162:
	;
	goto L163
L163:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v1204 = F_TS_execute(m, v1197+int32(8), v25+int32(-56), int32(0), int32(1533))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L20
	} else {
		goto L183
	}
L164:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1049)))
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1097+v1075<<(uint(int32(2))%32))))
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101))))
	if v1102 != int32(1) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	goto L163
L166:
	;
	v1170 = v1075 + int32(1)
	v1171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1049)+4)))
	if v1170 < v1171 {
		v1075 = v1170
		goto L164
	} else {
		goto L182
	}
L167:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v1111 = base.I32_div_s(v1101-v1106-int32(8), int32(12))
	v1114 = v1105 + v1111*int32(_a_F_calc_rank_cd_0)
	v1115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1114))) = uint8(v1115)
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1114)+1)))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+4))
	if v1118 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1114)+4)) = v1162
	goto L166
L169:
	;
	if v1117&int32(1) != 0 {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	v1130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1049)+8)))
	v1132 = v1114 + int32(8)
	v1135 = int32(1)
	v1138 = v1117 & v1135
	if v1138 != 0 {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v1125 = int32(_a_F_calc_rank_cd_2)
	goto L174
L173:
	;
	v1125 = int32(0)
	goto L174
L174:
	;
	v1127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1049)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1114+v1125)+8)) = uint16(v1127)
	v1162 = int32(1)
	goto L168
L175:
	;
	v1139 = int32(_a_F_calc_rank_cd_3) - v1118
	goto L177
L176:
	;
	v1139 = v1118 - v1135
	goto L177
L177:
	;
	v1143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1132+v1139<<(uint(int32(1))%32)))))
	if (v1130^v1143)&int32(_a_F_calc_rank_cd_4) == int32(0) {
		goto L166
	} else {
		goto L178
	}
L178:
	;
	if v1138 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v1151 = int32(_a_F_calc_rank_cd_4) - v1118
	goto L181
L180:
	;
	v1151 = v1118
	goto L181
L181:
	;
	v1152 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1132+v1151<<(uint(v1152)%32)))) = uint16(v1130)
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+4))
	v1162 = v1156 + v1152
	goto L168
L182:
	;
	goto L165
L183:
	;
	if v1204 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1209 = v1049 - int32(12)
	if base.Ui32(v804) <= base.Ui32(v1209) {
		v1049 = v1209
		goto L159
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	goto L160
L187:
	;
	goto L157
L188:
	;
	goto L157
L189:
	;
	v1248 = v1049
	v1265 = v1244
	goto L192
L190:
	;
	v1302 = v1244
	goto L191
L191:
	;
	v1307 = v805 - v1049
	v1309 = base.I32_div_s(v1307, int32(12))
	v1315 = base.I32_div_s(v1307, int32(24))
	v1317 = v968 - (v1309 + v1213)
	if v1317 < int32(0) {
		goto L195
	} else {
		goto L196
	}
L192:
	;
	v1272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1248)+8)))
	v1278 = *(*float64)(unsafe.Add(mBase, uint32(v25+int32(-48)+int32(base.Ui32(v1272)>>(uint(int32(11))%32))&int32(24))))
	v1279 = base.F64_add(v1265, v1278)
	v1281 = v1248 + int32(12)
	if base.Ui32(v1281) <= base.Ui32(v805) {
		v1248 = v1281
		v1265 = v1279
		goto L192
	} else {
		goto L194
	}
L193:
	;
	v1302 = v1279
	goto L191
L194:
	;
	goto L193
L195:
	;
	v1320 = v1315
	goto L197
L196:
	;
	v1320 = v1317
	goto L197
L197:
	;
	v1328 = base.F64_mul(base.F64_convert_i32_u(v1213+v968), float64(0.5))
	if v689 <= int32(0) {
		v1338 = v695
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v1339 = int32(1)
	v687 = v1243 + v1339
	v689 = v689 + v1339
	v693 = base.F64_add(v693, base.F64_div(base.F64_div(base.F64_convert_i32_s(v1309+int32(1)), v1302), base.F64_convert_i32_s(v1320+int32(1))))
	v694 = v1328
	v695 = v1338
	goto L108
L199:
	;
	if base.F64_gt(v1328, v694) == int32(0) {
		v1338 = v695
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1338 = base.F64_add(v695, base.F64_div(float64(1), base.F64_sub(v1328, v694)))
	goto L198
L201:
	;
	goto L122
L202:
	;
	if l3&int32(2) == int32(0) {
		v1575 = v1486
		goto L217
	} else {
		goto L218
	}
L203:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1378 <= int32(0) {
		v1486 = v693
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1384 = v99 + v1378<<(uint(int32(2))%32)
	if base.Ui32(v99) < base.Ui32(v1384) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1389 = v99
	v1391 = int32(0)
	goto L208
L206:
	;
	v1465 = float64(1)
	goto L207
L207:
	;
	v1466 = F_log(m, v1465)
	mBase = m.M
	v1486 = base.F64_div(v693, v1466)
	goto L202
L208:
	;
	v1411 = int32(1)
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v1389)))
	if v1412&v1411 != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v1465 = base.F64_convert_i32_s(v1434 + int32(1))
	goto L207
L210:
	;
	v1415 = int32(1)
	v1428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1384+(int32(base.Ui32(v1412)>>(uint(v1415)%32))&int32(2047)+int32(base.Ui32(v1412)>>(uint(int32(12))%32))+v1415)&int32(_a_F_calc_rank_cd_1)))))
	if base.Ui32(v1428) <= base.Ui32(v1415) {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	v1433 = v1411
	goto L212
L212:
	;
	v1434 = v1433 + v1391
	v1436 = v1389 + int32(4)
	if base.Ui32(v1436) < base.Ui32(v1384) {
		v1389 = v1436
		v1391 = v1434
		goto L208
	} else {
		goto L216
	}
L213:
	;
	v1431 = v1415
	goto L215
L214:
	;
	v1431 = v1428
	goto L215
L215:
	;
	v1433 = v1431
	goto L212
L216:
	;
	goto L209
L217:
	;
	if l3&int32(4) == int32(0) {
		v1594 = v1575
		goto L230
	} else {
		goto L231
	}
L218:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1499 = v99 + v1496<<(uint(int32(2))%32)
	if base.Ui32(v1499) <= base.Ui32(v99) {
		v1575 = v1486
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v1502 = int32(0)
	v1513 = v99
	goto L220
L220:
	;
	v1526 = int32(1)
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1513)))
	if v1527&v1526 != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	if v1549 <= int32(0) {
		v1575 = v1486
		goto L217
	} else {
		goto L229
	}
L222:
	;
	v1530 = int32(1)
	v1543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1499+(int32(base.Ui32(v1527)>>(uint(v1530)%32))&int32(2047)+int32(base.Ui32(v1527)>>(uint(int32(12))%32))+v1530)&int32(_a_F_calc_rank_cd_1)))))
	if base.Ui32(v1543) <= base.Ui32(v1530) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	v1548 = v1526
	goto L224
L224:
	;
	v1549 = v1548 + v1502
	v1551 = v1513 + int32(4)
	if base.Ui32(v1551) < base.Ui32(v1499) {
		v1502 = v1549
		v1513 = v1551
		goto L220
	} else {
		goto L228
	}
L225:
	;
	v1546 = v1530
	goto L227
L226:
	;
	v1546 = v1543
	goto L227
L227:
	;
	v1548 = v1546
	goto L224
L228:
	;
	goto L221
L229:
	;
	v1575 = base.F64_div(v1486, base.F64_convert_i32_u(v1549))
	goto L217
L230:
	;
	if l3&int32(8) == int32(0) {
		v1605 = v1594
		goto L234
	} else {
		goto L235
	}
L231:
	;
	if v689 <= int32(0) {
		v1594 = v1575
		goto L230
	} else {
		goto L232
	}
L232:
	;
	if base.F64_gt(v695, float64(0)) == int32(0) {
		v1594 = v1575
		goto L230
	} else {
		goto L233
	}
L233:
	;
	v1594 = base.F64_div(v1575, base.F64_div(base.F64_convert_i32_u(v689), v695))
	goto L230
L234:
	;
	if l3&int32(16) == int32(0) {
		v1621 = v1605
		goto L237
	} else {
		goto L238
	}
L235:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1599 <= int32(0) {
		v1605 = v1594
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v1605 = base.F64_div(v1594, base.F64_convert_i32_u(v1599))
	goto L234
L237:
	;
	if l3&int32(32) != 0 {
		goto L240
	} else {
		goto L241
	}
L238:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1610 <= int32(0) {
		v1621 = v1605
		goto L237
	} else {
		goto L239
	}
L239:
	;
	v1616 = F_log(m, base.F64_convert_i32_s(v1610+int32(1)))
	mBase = m.M
	v1621 = base.F64_div(v1605, base.F64_div(v1616, float64(0.6931471805599453)))
	goto L237
L240:
	;
	v1627 = base.F64_div(v1621, base.F64_add(v1621, float64(1)))
	goto L242
L241:
	;
	v1627 = v1621
	goto L242
L242:
	;
	F_pfree(m, v498)
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L20
	} else {
		goto L243
	}
L243:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	F_pfree(m, v1630)
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L20
	} else {
		goto L244
	}
L244:
	;
	v1658 = base.F32_demote_f64(v1627)
	goto L23
L245:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L20
	} else {
		goto L246
	}
L246:
	;
	F_errmsg(m, int32(_a_F_calc_rank_cd_5), int32(0))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L20
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_calc_rank_cd_6), int32(876), int32(_a_F_calc_rank_cd_7))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L20
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cancel_prior_stmt_triggers(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[0]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[1]))
	v16 = v11 + v13*int32(20)
	v18 = v16 + int32(16)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v19 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
	if v85 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v59 = int32(_a_F_cancel_prior_stmt_triggers_0)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[2]))
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[2])) = v63
	v66 = F_palloc0(m, int32(36))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v22 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v30 = int32(0)
	goto L5
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25+v30<<(uint(int32(2))%32))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v40 != l0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	v48 = v30 + int32(1)
	if v22 != v48 {
		v30 = v48
		goto L5
	} else {
		goto L11
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v42 != l1 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)))
	if v44 != int32(1) {
		v80 = v39
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
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = l0
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v71 = F_lappend(m, v70, v66)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v71
	*(*int32)(unsafe.Add(mBase, _c_F_cancel_prior_stmt_triggers[2])) = v60
	v80 = v66
	goto L1
L15:
	;
	v173 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v173)
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v80)+12)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = v177
	return
L16:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	if v88 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v97 = v94
	v99 = v95
	goto L22
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	v94 = v88
	v95 = v89
	goto L17
L19:
	;
	goto L20
L20:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v91 == v90 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v94 = v91
	v95 = v90
	goto L17
L22:
	;
	if v99 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L15
L24:
	;
	v107 = v99
	goto L26
L25:
	;
	v107 = v97 + int32(16)
	goto L26
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if base.Ui32(v107) < base.Ui32(v108) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v113 = v107
	goto L30
L28:
	;
	goto L29
L29:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v163 != 0 {
		v97 = v163
		v99 = int32(0)
		goto L22
	} else {
		goto L42
	}
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v122 = v113 + v119&int32(134217727)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	if v123 != l0 {
		goto L15
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v125&int32(3) != l2 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	if v125&int32(28) != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v119&int32(1073741823) | int32(-2147483648)
	v138 = v119 & int32(939524096)
	if v138 == int32(134217728) {
		v149 = int32(24)
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v150 = v149 + v113
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if base.Ui32(v150) < base.Ui32(v151) {
		v113 = v150
		goto L30
	} else {
		goto L41
	}
L36:
	;
	if v138 == int32(268435456) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v145 = int32(12)
	goto L39
L38:
	;
	v145 = int32(4)
	goto L39
L39:
	;
	if v138 != int32(805306368) {
		v149 = v145
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v149 = int32(16)
	goto L35
L41:
	;
	goto L31
L42:
	;
	goto L23
}
func F_canonicalize_ec_expression(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
					v35 = v7
					v38 = F_exprCollation(m, l0)
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						if v38 == l2 {
							v52 = l0
							return v52
						} else {
							v41 = F_exprTypmod(m, l0)
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = v35
								v45 = v41
								v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
						v43 = l1
						v45 = int32(-1)
						v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = v50
							return v52
						}
					} else {
						v35 = l1
						v38 = F_exprCollation(m, l0)
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 == l2 {
								v52 = l0
								return v52
							} else {
								v41 = F_exprTypmod(m, l0)
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = v35
									v45 = v41
									v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
							v43 = l1
							v45 = int32(-1)
							v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = v50
								return v52
							}
						} else {
							v35 = l1
							v38 = F_exprCollation(m, l0)
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								if v38 == l2 {
									v52 = l0
									return v52
								} else {
									v41 = F_exprTypmod(m, l0)
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v43 = v35
										v45 = v41
										v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
						v35 = v7
						v38 = F_exprCollation(m, l0)
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 == l2 {
								v52 = l0
								return v52
							} else {
								v41 = F_exprTypmod(m, l0)
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = v35
									v45 = v41
									v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
					v35 = v7
					v38 = F_exprCollation(m, l0)
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						if v38 == l2 {
							v52 = l0
							return v52
						} else {
							v41 = F_exprTypmod(m, l0)
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = v35
								v45 = v41
								v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
							v43 = l1
							v45 = int32(-1)
							v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = v50
								return v52
							}
						} else {
							v35 = l1
							v38 = F_exprCollation(m, l0)
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								if v38 == l2 {
									v52 = l0
									return v52
								} else {
									v41 = F_exprTypmod(m, l0)
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v43 = v35
										v45 = v41
										v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
						v35 = v7
						v38 = F_exprCollation(m, l0)
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 == l2 {
								v52 = l0
								return v52
							} else {
								v41 = F_exprTypmod(m, l0)
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = v35
									v45 = v41
									v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
			if base.Ui32(l1-int32(_a_F_canonicalize_ec_expression_0)) < base.Ui32(int32(4)) {
				v35 = v7
				v38 = F_exprCollation(m, l0)
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					if v38 == l2 {
						v52 = l0
						return v52
					} else {
						v41 = F_exprTypmod(m, l0)
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = v35
							v45 = v41
							v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
				if base.Ui32(l1-int32(_a_F_canonicalize_ec_expression_1)) < base.Ui32(int32(2)) {
					v35 = v7
					v38 = F_exprCollation(m, l0)
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						if v38 == l2 {
							v52 = l0
							return v52
						} else {
							v41 = F_exprTypmod(m, l0)
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = v35
								v45 = v41
								v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
					if l1 == int32(3831) {
						v35 = v7
						v38 = F_exprCollation(m, l0)
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 == l2 {
								v52 = l0
								return v52
							} else {
								v41 = F_exprTypmod(m, l0)
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = v35
									v45 = v41
									v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
							v43 = l1
							v45 = int32(-1)
							v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = v50
								return v52
							}
						} else {
							v35 = l1
							v38 = F_exprCollation(m, l0)
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								if v38 == l2 {
									v52 = l0
									return v52
								} else {
									v41 = F_exprTypmod(m, l0)
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v43 = v35
										v45 = v41
										v50 = F_applyRelabelType(m, l0, v43, v45, l2, int32(2), int32(-1), int32(0))
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32((l0-v15*v14)&v12<<(uint(v21)%32))+uint32(_c_F_casemap[0])))
	v26 = int32(8)
	v27 = int32(base.Ui32(l0) >> (uint(v26) % 32))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_casemap[1]))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30*int32(86)+v15)+uint32(_c_F_casemap[1]))))
	v41 = base.I32_rem_u_s(int32(base.Ui32(v25*v36)>>(uint(int32(11))%32)), int32(6))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_casemap[2]))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32((v41+v44)<<(uint(v21)%32))+uint32(_c_F_casemap[3])))
	v52 = v50 >> (uint(v26) % 32)
	v54 = v50 & v12
	if base.Ui32(v54) <= base.Ui32(int32(1)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v52&(int32(0)-(l1^v54)) + l0
L4:
	;
	goto L5
L5:
	;
	v64 = v52 & int32(255)
	if v64 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v71 = v64
	v72 = int32(base.Ui32(v52) >> (uint(int32(8)) % 32))
	goto L7
L7:
	;
	v78 = int32(1)
	v79 = int32(base.Ui32(v71) >> (uint(v78) % 32))
	v80 = v79 + v72
	v82 = v80 << (uint(v78) % 32)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+uint32(_c_F_casemap[4]))))
	if v85 == v13 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L1
L9:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+uint32(_c_F_casemap[5]))))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87<<(uint(int32(2))%32))+uint32(_c_F_casemap[3])))
	v94 = v92 & int32(255)
	if base.Ui32(v94) <= base.Ui32(int32(1)) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v110 = base.B2i32(base.Ui32(v13) < base.Ui32(v85))
	if base.Ui32(v13) < base.Ui32(v85) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	return (int32(0)-(l1^v94))&(v92>>(uint(int32(8))%32)) + l0
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
	v107 = int32(-1)
	goto L17
L16:
	;
	v107 = int32(1)
	goto L17
L17:
	;
	return v107 + l0
L18:
	;
	v111 = v72
	goto L20
L19:
	;
	v111 = v80
	goto L20
L20:
	;
	if base.Ui32(v13) < base.Ui32(v85) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v113 = v79
	goto L23
L22:
	;
	v113 = v71 - v79
	goto L23
L23:
	;
	if v113 != 0 {
		v71 = v113
		v72 = v111
		goto L7
	} else {
		goto L24
	}
L24:
	;
	goto L8
}
func F_cashlarger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v7 < v5 {
		v9 = v5
	} else {
		v9 = v7
	}
	v10 = F_Int64GetDatum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
