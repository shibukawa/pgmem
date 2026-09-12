package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_calc_hist_selectivity_contained(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) float64 {
	mBase := m.M
	_ = mBase
	var v14 float64
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 float64
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 float64
	_ = v114
	var v115 int32
	_ = v115
	var v118 float64
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v131 float64
	_ = v131
	var v134 float64
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 float64
	_ = v144
	var v146 float64
	_ = v146
	var v150 float64
	_ = v150
	var v151 int32
	_ = v151
	var v152 float64
	_ = v152
	var v155 float64
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 float64
	_ = v162
	var v163 int32
	_ = v163
	var v166 float64
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 float64
	_ = v173
	var v179 float64
	_ = v179
	var v182 float64
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 float64
	_ = v192
	var v194 float64
	_ = v194
	var v195 float64
	_ = v195
	var v202 float64
	_ = v202
	var v213 float64
	_ = v213
	var v218 float64
	_ = v218
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 float64
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 float64
	_ = v276
	var v277 float64
	_ = v277
	var v278 float64
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 float64
	_ = v281
	var v282 float64
	_ = v282
	var v284 int32
	_ = v284
	var v301 float64
	_ = v301
	var v305 float64
	_ = v305
	var v307 int32
	_ = v307
	var v313 float64
	_ = v313
	var v316 float64
	_ = v316
	var v317 float64
	_ = v317
	var v325 int32
	_ = v325
	var v331 float64
	_ = v331
	var v333 float64
	_ = v333
	var v334 float64
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 float64
	_ = v345
	var v352 int32
	_ = v352
	var v355 float64
	_ = v355
	var v358 float64
	_ = v358
	var v368 float64
	_ = v368
	var v373 int32
	_ = v373
	var v374 float64
	_ = v374
	var v377 float64
	_ = v377
	var v378 float64
	_ = v378
	var v379 int32
	_ = v379
	var v380 float64
	_ = v380
	var v382 int32
	_ = v382
	var v399 float64
	_ = v399
	var v403 float64
	_ = v403
	var v404 float64
	_ = v404
	var v409 float64
	_ = v409
	var v416 int32
	_ = v416
	var v420 float64
	_ = v420
	var v422 float64
	_ = v422
	var v423 float64
	_ = v423
	var v425 float64
	_ = v425
	var v429 float64
	_ = v429
	var v433 float64
	_ = v433
	var v443 float64
	_ = v443
	var v464 float64
	_ = v464
	var v474 float64
	_ = v474
	v14 = float64(0)
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v19)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	v23 = v21 ^ v19
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v23)
	v27 = l4 - v19
	v35 = int32(-1)
	v37 = v27
	goto L1
L1:
	;
	v50 = base.I32_div_s(v35+v37+int32(1), int32(2))
	v54 = F_range_cmp_bounds(m, l0, l3+v50<<(uint(int32(3))%32), l2)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	if v60 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	return float64(0)
L4:
	;
	v59 = base.B2i32(v54 < int32(0))
	if v54 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v60 = v50
	goto L7
L6:
	;
	v60 = v35
	goto L7
L7:
	;
	if v54 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v63 = v37
	goto L10
L9:
	;
	v63 = v50 - int32(1)
	goto L10
L10:
	;
	if v60 < v63 {
		v35 = v60
		v37 = v63
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
	v70 = l0 + int32(268)
	v73 = l4 - int32(2)
	if v60 < v73 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v75 = v60
	goto L17
L16:
	;
	v75 = v73
	goto L17
L17:
	;
	v78 = l3 + v75<<(uint(int32(3))%32)
	v81 = F_get_position(m, l0, l2, v78, v78+int32(8))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v91 = v75
	v97 = v81
	v98 = v14
	v99 = v14
	goto L19
L19:
	;
	v103 = l3 + v91<<(uint(int32(3))%32)
	v104 = F_range_cmp_bounds(m, l0, v103, l1)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L22
	}
L20:
	;
	return v474
L21:
	;
	v202 = float64(0)
	if base.F64_lt(v194, v202) != 0 {
		v464 = v202
		goto L70
	} else {
		goto L71
	}
L22:
	;
	if v104 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v108 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
	if v156 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L26:
	;
	v150 = F_get_position(m, l0, l1, v103, v103+int32(8))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L46
	}
L27:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v113 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v136 != int32(1) {
		v146 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L26
	} else {
		goto L42
	}
L30:
	;
	v114 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L32
L31:
	;
	v114 = float64(1)
	goto L32
L32:
	;
	if v113 != 0 {
		v146 = v114
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v115 == int32(0) {
		v146 = v114
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v118 = float64(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v123 = F_FunctionCall2Coll(m, v70, v120, v121, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v125)&int64(9223372036854775807)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v131 = v118
	goto L38
L37:
	;
	v131 = v125
	goto L38
L38:
	;
	if base.F64_lt(v125, float64(0)) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v134 = v118
	goto L41
L40:
	;
	v134 = v131
	goto L41
L41:
	;
	v146 = v134
	goto L26
L42:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v141 == v142 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v144 = float64(0)
	goto L45
L44:
	;
	v144 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L45
L45:
	;
	v146 = v144
	goto L26
L46:
	;
	v152 = base.F64_sub(v97, v150)
	if base.F64_lt(v152, float64(0)) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v155 = float64(0)
	goto L49
L48:
	;
	v155 = v152
	goto L49
L49:
	;
	v194 = v146
	v195 = v155
	goto L21
L50:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v161 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v184 != int32(1) {
		v194 = math.Float64frombits(uint64(0x7ff0000000000000))
		v195 = v97
		goto L21
	} else {
		goto L65
	}
L53:
	;
	v162 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L55
L54:
	;
	v162 = float64(1)
	goto L55
L55:
	;
	if v161 != 0 {
		v194 = v162
		v195 = v97
		goto L21
	} else {
		goto L56
	}
L56:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v163 == int32(0) {
		v194 = v162
		v195 = v97
		goto L21
	} else {
		goto L57
	}
L57:
	;
	v166 = float64(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v171 = F_FunctionCall2Coll(m, v70, v168, v169, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v173 = *(*float64)(unsafe.Add(mBase, uint32(v171)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v173)&int64(9223372036854775807)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v179 = v166
	goto L61
L60:
	;
	v179 = v173
	goto L61
L61:
	;
	if base.F64_lt(v173, float64(0)) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v182 = v166
	goto L64
L63:
	;
	v182 = v179
	goto L64
L64:
	;
	v194 = v182
	v195 = v97
	goto L21
L65:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+6)))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v189 == v190 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v192 = float64(0)
	goto L68
L67:
	;
	v192 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L68
L68:
	;
	v194 = v192
	v195 = v97
	goto L21
L69:
	;
	v474 = base.F64_add(v98, base.F64_div(base.F64_mul(v195, v464), base.F64_convert_i32_u(v27)))
	if int32(0) <= v104 {
		goto L145
	} else {
		goto L146
	}
L70:
	;
	goto L69
L71:
	;
	v213 = float64(1)
	v218 = base.F64_abs(v194)
	if int32(1)&base.F64_eq(v218, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v464 = v213
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v224 = l6 - int32(1)
	if v224 < int32(0) {
		v464 = v213
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v228 = v224
	v232 = int32(-1)
	goto L74
L74:
	;
	v249 = int32(2)
	v250 = base.I32_div_s(v228+v232+int32(1), v249)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l5+v250<<(uint(v249)%32))))
	v255 = *(*float64)(unsafe.Add(mBase, uint32(v254)))
	if base.F64_gt(v99, v255) != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	if v224 <= v265 {
		v464 = v213
		goto L70
	} else {
		goto L87
	}
L76:
	;
	if v265 < v263 {
		v228 = v263
		v232 = v265
		goto L74
	} else {
		goto L86
	}
L77:
	;
	v263 = v228
	v265 = v250
	goto L76
L78:
	;
	goto L79
L79:
	;
	v260 = int32(1) & base.F64_ge(v99, v255)
	if v260 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v261 = v228
	goto L82
L81:
	;
	v261 = v250 - int32(1)
	goto L82
L82:
	;
	if v260 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v262 = v250
	goto L85
L84:
	;
	v262 = v232
	goto L85
L85:
	;
	v263 = v261
	v265 = v262
	goto L76
L86:
	;
	goto L75
L87:
	;
	if v265 < int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v316 = base.F64_convert_i32_s(v224)
	v317 = base.F64_div(base.F64_add(v313, base.F64_convert_i32_u(v307)), v316)
	if base.F64_eq(v99, v194) != 0 {
		v464 = v317
		goto L70
	} else {
		goto L105
	}
L89:
	;
	v307 = int32(0)
	v313 = float64(0)
	goto L88
L90:
	;
	goto L91
L91:
	;
	v274 = l5 + v265<<(uint(int32(2))%32)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	v276 = *(*float64)(unsafe.Add(mBase, uint32(v275)))
	v277 = base.F64_abs(v276)
	v278 = math.Float64frombits(uint64(0x7ff0000000000000))
	v279 = base.F64_eq(v277, v278)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	v281 = *(*float64)(unsafe.Add(mBase, uint32(v280)))
	v282 = base.F64_abs(v281)
	v284 = base.F64_eq(v282, v278)
	if v284 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v279 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	if v279 != 0 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	if base.F64_eq(base.F64_abs(v99), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v307 = v265
		v313 = float64(0.5)
		goto L88
	} else {
		goto L95
	}
L95:
	;
	v307 = v265
	v313 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v276, v99), base.F64_sub(v276, v281)))
	goto L88
L96:
	;
	if base.F64_eq(v277, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	if v284 == int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v307 = v265
	v313 = float64(1)
	goto L88
L99:
	;
	v301 = float64(0)
	goto L101
L100:
	;
	v301 = float64(0.5)
	goto L101
L101:
	;
	if base.F64_eq(v282, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v305 = v301
	goto L104
L103:
	;
	v305 = float64(0.5)
	goto L104
L104:
	;
	v307 = v265
	v313 = v305
	goto L88
L105:
	;
	if v224 <= v307 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v429 = float64(0)
	v433 = base.F64_div(base.F64_add(v425, base.F64_convert_i32_u(v416)), v316)
	if base.F64_gt(v420, v429)|base.F64_gt(v433, v429) != 0 {
		goto L138
	} else {
		goto L139
	}
L107:
	;
	v416 = v307
	v420 = v317
	v422 = v99
	v423 = v202
	v425 = v202
	goto L106
L108:
	;
	goto L109
L109:
	;
	v325 = v307
	v331 = v317
	v333 = v202
	v334 = v99
	goto L111
L110:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l5+v325<<(uint(int32(2))%32))))
	v374 = *(*float64)(unsafe.Add(mBase, uint32(v373)))
	if base.F64_eq(v345, v374) != 0 {
		goto L121
	} else {
		goto L122
	}
L111:
	;
	v339 = int32(1)
	v340 = v325 + v339
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l5+v340<<(uint(int32(2))%32))))
	v345 = *(*float64)(unsafe.Add(mBase, uint32(v344)))
	if base.B2i32(base.F64_ge(v194, v345) == int32(0))|int32(0) != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v416 = v224
	v420 = v358
	v422 = v345
	v423 = v368
	v425 = v202
	goto L106
L113:
	;
	v352 = base.F64_lt(v345, v194)
	goto L115
L114:
	;
	v352 = v339
	goto L115
L115:
	;
	if v352 == int32(0) {
		goto L110
	} else {
		goto L116
	}
L116:
	;
	v355 = float64(0)
	v358 = base.F64_div(base.F64_convert_i32_u(v325), v316)
	if base.F64_gt(v331, v355)|base.F64_gt(v358, v355) != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v368 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v331, v358), float64(0.5)), base.F64_sub(v345, v334)), v333)
	goto L119
L118:
	;
	v368 = v333
	goto L119
L119:
	;
	if v224 != v340 {
		v325 = v340
		v331 = v358
		v333 = v368
		v334 = v345
		goto L111
	} else {
		goto L120
	}
L120:
	;
	goto L112
L121:
	;
	v409 = float64(0)
	goto L123
L122:
	;
	v377 = base.F64_abs(v345)
	v378 = math.Float64frombits(uint64(0x7ff0000000000000))
	v379 = base.F64_eq(v377, v378)
	v380 = base.F64_abs(v374)
	v382 = base.F64_eq(v380, v378)
	if v382 != 0 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	v416 = v325
	v420 = v331
	v422 = v334
	v423 = v333
	v425 = v409
	goto L106
L124:
	;
	v409 = v404
	goto L123
L125:
	;
	if v379 != 0 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	if v379 != 0 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	if base.F64_eq(base.F64_abs(v194), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v404 = float64(0.5)
		goto L124
	} else {
		goto L128
	}
L128:
	;
	v404 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v345, v194), base.F64_sub(v345, v374)))
	goto L124
L129:
	;
	if base.F64_eq(v377, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	if v382 == int32(0) {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v404 = float64(1)
	goto L124
L132:
	;
	v399 = float64(0)
	goto L134
L133:
	;
	v399 = float64(0.5)
	goto L134
L134:
	;
	if base.F64_eq(v380, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v403 = v399
	goto L137
L136:
	;
	v403 = float64(0.5)
	goto L137
L137:
	;
	v404 = v403
	goto L124
L138:
	;
	v443 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v420, v433), float64(0.5)), base.F64_sub(v194, v422)), v423)
	goto L140
L139:
	;
	v443 = v423
	goto L140
L140:
	;
	if base.F64_eq(v218, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	if base.F64_eq(base.F64_abs(v443), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v464 = float64(0.5)
		goto L70
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v464 = base.F64_div(v443, base.F64_sub(v194, v99))
	goto L70
L144:
	;
	goto L143
L145:
	;
	if int32(0) < v91 {
		v91 = v91 - int32(1)
		v97 = float64(1)
		v98 = v474
		v99 = v194
		goto L19
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	goto L20
L148:
	;
	goto L147
}
func F_calc_length_hist_frac(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 int32) float64 {
	mBase := m.M
	_ = mBase
	var v11 float64
	_ = v11
	var v22 float64
	_ = v22
	var v24 int32
	_ = v24
	var v27 float64
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v93 int32
	_ = v93
	var v110 float64
	_ = v110
	var v114 float64
	_ = v114
	var v116 int32
	_ = v116
	var v122 float64
	_ = v122
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v134 int32
	_ = v134
	var v140 float64
	_ = v140
	var v142 float64
	_ = v142
	var v143 float64
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 float64
	_ = v154
	var v161 int32
	_ = v161
	var v164 float64
	_ = v164
	var v167 float64
	_ = v167
	var v177 float64
	_ = v177
	var v182 int32
	_ = v182
	var v183 float64
	_ = v183
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v188 int32
	_ = v188
	var v189 float64
	_ = v189
	var v191 int32
	_ = v191
	var v208 float64
	_ = v208
	var v212 float64
	_ = v212
	var v213 float64
	_ = v213
	var v218 float64
	_ = v218
	var v225 int32
	_ = v225
	var v229 float64
	_ = v229
	var v231 float64
	_ = v231
	var v232 float64
	_ = v232
	var v234 float64
	_ = v234
	var v238 float64
	_ = v238
	var v242 float64
	_ = v242
	var v252 float64
	_ = v252
	var v273 float64
	_ = v273
	v11 = float64(0)
	if base.F64_lt(l3, v11) != 0 {
		v273 = v11
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v273
L2:
	;
	v22 = float64(1)
	v24 = l4 ^ int32(1)
	v27 = base.F64_abs(l3)
	if base.B2i32(v24 == int32(0))&base.F64_eq(v27, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v273 = v22
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = l1 - int32(1)
	if v33 < int32(0) {
		v273 = v22
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = v33
	v41 = int32(-1)
	goto L5
L5:
	;
	v58 = int32(2)
	v59 = base.I32_div_s(v37+v41+int32(1), v58)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0+v59<<(uint(v58)%32))))
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v63)))
	if base.F64_gt(l2, v64) != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v33 <= v74 {
		v273 = v22
		goto L1
	} else {
		goto L18
	}
L7:
	;
	if v74 < v72 {
		v37 = v72
		v41 = v74
		goto L5
	} else {
		goto L17
	}
L8:
	;
	v72 = v37
	v74 = v59
	goto L7
L9:
	;
	goto L10
L10:
	;
	v69 = l4 & base.F64_ge(l2, v64)
	if v69 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v70 = v37
	goto L13
L12:
	;
	v70 = v59 - int32(1)
	goto L13
L13:
	;
	if v69 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = v59
	goto L16
L15:
	;
	v71 = v41
	goto L16
L16:
	;
	v72 = v70
	v74 = v71
	goto L7
L17:
	;
	goto L6
L18:
	;
	if v74 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v125 = base.F64_convert_i32_s(v33)
	v126 = base.F64_div(base.F64_add(v122, base.F64_convert_i32_u(v116)), v125)
	if base.F64_eq(l2, l3) != 0 {
		v273 = v126
		goto L1
	} else {
		goto L36
	}
L20:
	;
	v116 = int32(0)
	v122 = float64(0)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v83 = l0 + v74<<(uint(int32(2))%32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v85 = *(*float64)(unsafe.Add(mBase, uint32(v84)))
	v86 = base.F64_abs(v85)
	v87 = math.Float64frombits(uint64(0x7ff0000000000000))
	v88 = base.F64_eq(v86, v87)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v89)))
	v91 = base.F64_abs(v90)
	v93 = base.F64_eq(v91, v87)
	if v93 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v88 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	if v88 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if base.F64_eq(base.F64_abs(l2), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v116 = v74
		v122 = float64(0.5)
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v116 = v74
	v122 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v85, l2), base.F64_sub(v85, v90)))
	goto L19
L27:
	;
	if base.F64_eq(v86, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v93 == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v116 = v74
	v122 = float64(1)
	goto L19
L30:
	;
	v110 = float64(0)
	goto L32
L31:
	;
	v110 = float64(0.5)
	goto L32
L32:
	;
	if base.F64_eq(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v114 = v110
	goto L35
L34:
	;
	v114 = float64(0.5)
	goto L35
L35:
	;
	v116 = v74
	v122 = v114
	goto L19
L36:
	;
	if v33 <= v116 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v238 = float64(0)
	v242 = base.F64_div(base.F64_add(v234, base.F64_convert_i32_u(v225)), v125)
	if base.F64_gt(v229, v238)|base.F64_gt(v242, v238) != 0 {
		goto L69
	} else {
		goto L70
	}
L38:
	;
	v225 = v116
	v229 = v126
	v231 = l2
	v232 = v11
	v234 = v11
	goto L37
L39:
	;
	goto L40
L40:
	;
	v134 = v116
	v140 = v126
	v142 = v11
	v143 = l2
	goto L42
L41:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0+v134<<(uint(int32(2))%32))))
	v183 = *(*float64)(unsafe.Add(mBase, uint32(v182)))
	if base.F64_eq(v154, v183) != 0 {
		goto L52
	} else {
		goto L53
	}
L42:
	;
	v148 = int32(1)
	v149 = v134 + v148
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0+v149<<(uint(int32(2))%32))))
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v153)))
	if base.B2i32(base.F64_ge(l3, v154) == int32(0))|v24 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v225 = v33
	v229 = v167
	v231 = v154
	v232 = v177
	v234 = v11
	goto L37
L44:
	;
	v161 = base.F64_lt(v154, l3)
	goto L46
L45:
	;
	v161 = v148
	goto L46
L46:
	;
	if v161 == int32(0) {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v164 = float64(0)
	v167 = base.F64_div(base.F64_convert_i32_u(v134), v125)
	if base.F64_gt(v140, v164)|base.F64_gt(v167, v164) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v177 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v140, v167), float64(0.5)), base.F64_sub(v154, v143)), v142)
	goto L50
L49:
	;
	v177 = v142
	goto L50
L50:
	;
	if v33 != v149 {
		v134 = v149
		v140 = v167
		v142 = v177
		v143 = v154
		goto L42
	} else {
		goto L51
	}
L51:
	;
	goto L43
L52:
	;
	v218 = float64(0)
	goto L54
L53:
	;
	v186 = base.F64_abs(v154)
	v187 = math.Float64frombits(uint64(0x7ff0000000000000))
	v188 = base.F64_eq(v186, v187)
	v189 = base.F64_abs(v183)
	v191 = base.F64_eq(v189, v187)
	if v191 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v225 = v134
	v229 = v140
	v231 = v143
	v232 = v142
	v234 = v218
	goto L37
L55:
	;
	v218 = v213
	goto L54
L56:
	;
	if v188 != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	if v188 != 0 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	if base.F64_eq(base.F64_abs(l3), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v213 = float64(0.5)
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v213 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v154, l3), base.F64_sub(v154, v183)))
	goto L55
L60:
	;
	if base.F64_eq(v186, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	if v191 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v213 = float64(1)
	goto L55
L63:
	;
	v208 = float64(0)
	goto L65
L64:
	;
	v208 = float64(0.5)
	goto L65
L65:
	;
	if base.F64_eq(v189, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v212 = v208
	goto L68
L67:
	;
	v212 = float64(0.5)
	goto L68
L68:
	;
	v213 = v212
	goto L55
L69:
	;
	v252 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v229, v242), float64(0.5)), base.F64_sub(l3, v231)), v232)
	goto L71
L70:
	;
	v252 = v232
	goto L71
L71:
	;
	if base.F64_eq(v27, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if base.F64_eq(base.F64_abs(v252), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v273 = float64(0.5)
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v273 = base.F64_div(v252, base.F64_sub(l3, l2))
	goto L1
L75:
	;
	goto L74
}
func F_calc_non_nestloop_required_outer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v6 = v5
	} else {
		v6 = v3
	}
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v7 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		v9 = v8
	} else {
		v9 = v3
	}
	v10 = F_bms_union(m, v6, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_calc_rank(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v29 float32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v96 int32
	_ = v96
	var v109 float32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v305 float32
	_ = v305
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v372 float32
	_ = v372
	var v374 float32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 float32
	_ = v384
	var v385 int32
	_ = v385
	var v386 float32
	_ = v386
	var v388 int32
	_ = v388
	var v392 float32
	_ = v392
	var v393 int32
	_ = v393
	var v426 float32
	_ = v426
	var v428 float32
	_ = v428
	var v429 int32
	_ = v429
	var v440 float32
	_ = v440
	var v442 int32
	_ = v442
	var v475 float32
	_ = v475
	var v480 int32
	_ = v480
	var v516 float32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v560 float32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v739 int32
	_ = v739
	var v757 float32
	_ = v757
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v795 int32
	_ = v795
	var v816 float32
	_ = v816
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v851 int32
	_ = v851
	var v867 float32
	_ = v867
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v887 int32
	_ = v887
	var v911 float32
	_ = v911
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v925 float32
	_ = v925
	var v926 int32
	_ = v926
	var v931 float32
	_ = v931
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v948 float64
	_ = v948
	var v956 float32
	_ = v956
	var v958 float32
	_ = v958
	var v961 float64
	_ = v961
	var v972 float32
	_ = v972
	var v975 int32
	_ = v975
	var v1005 float32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1040 float32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1075 float32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1113 float32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1132 int32
	_ = v1132
	var v1152 float32
	_ = v1152
	var v1159 float32
	_ = v1159
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1270 float64
	_ = v1270
	var v1271 float64
	_ = v1271
	var v1304 float32
	_ = v1304
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1411 float32
	_ = v1411
	var v1419 int32
	_ = v1419
	var v1425 float32
	_ = v1425
	var v1430 int32
	_ = v1430
	var v1437 float64
	_ = v1437
	var v1443 float32
	_ = v1443
	var v1459 int32
	_ = v1459
	var v1479 float32
	_ = v1479
	v5 = int32(0)
	v29 = float32(0)
	v33 = m.G0
	v35 = v33 - int32(16)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v37 == v5 {
		v1459 = v35
		v1479 = v29
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v1459 + int32(16)
	return v1479
L2:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v40 == int32(0) {
		v1459 = v35
		v1479 = v29
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v43 != int32(2) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if base.F32_lt(v1152, float32(0)) != 0 {
		goto L147
	} else {
		goto L148
	}
L5:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v525 = F_palloc0(m, v522<<(uint(int32(2))%32))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L9
	} else {
		goto L73
	}
L6:
	;
	v63 = m.G0
	v65 = v63 - int32(16)
	m.G0 = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = int32(1)
	v73 = F_SortAndUniqItems(m, l2, v65+int32(4))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L13
	}
L7:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	switch v46 - int32(2) {
	case 0, 2:
		goto L8
	default:
		goto L6
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v40
	v52 = F_SortAndUniqItems(m, l2, v35+int32(4))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return float32(0)
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if int32(1) < v56 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	F_pfree(m, v52)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v75 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v516 = float32(0)
	goto L16
L15:
	;
	v96 = v5
	v109 = v29
	goto L17
L16:
	;
	F_pfree(m, v73)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L9
	} else {
		goto L72
	}
L17:
	;
	v113 = int32(2)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v73+v96<<(uint(v113)%32))))
	v117 = int32(8)
	v118 = v65 + v117
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = int32(0)
	v129 = l1 + v117
	v132 = v129 + v125<<(uint(v113)%32)
	if base.Ui32(v132) <= base.Ui32(v129) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v516 = base.F32_div(v475, base.F32_convert_i32_u(v75))
	goto L16
L19:
	;
	v480 = v96 + int32(1)
	if v480 != v75 {
		v96 = v480
		v109 = v475
		goto L17
	} else {
		goto L71
	}
L20:
	;
	if v271 == int32(0) {
		v475 = v109
		goto L19
	} else {
		goto L50
	}
L21:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+2)))
	if v199 != int32(1) {
		goto L37
	} else {
		goto L38
	}
L22:
	;
	v193 = v129
	v194 = v132
	v195 = v132
	goto L21
L23:
	;
	goto L24
L24:
	;
	v140 = v129
	v142 = v132
	goto L25
L25:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v147 = int32(12)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v157 = int32(2)
	v164 = base.I32_div_s((v142-v140)>>(uint(v157)%32), v157)
	v167 = v140 + v164<<(uint(v157)%32)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v176 = int32(0)
	v177 = F_tsCompareString(m, l2+int32(8)+v146*v147+int32(base.Ui32(v150)>>(uint(v147)%32)), v150&int32(4095), v129+v156<<(uint(v157)%32)+int32(base.Ui32(v168)>>(uint(v147)%32)), int32(base.Ui32(v168)>>(uint(int32(1))%32))&int32(2047), v176)
	mBase = m.M
	if v177 == v176 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v193 = v186
	v194 = v167
	v195 = v187
	goto L21
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = int32(1)
	v193 = v140
	v194 = v167
	v195 = v167
	goto L21
L28:
	;
	goto L29
L29:
	;
	v185 = base.B2i32(int32(0) < v177)
	if int32(0) < v177 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v186 = v167 + int32(4)
	goto L32
L31:
	;
	v186 = v140
	goto L32
L32:
	;
	if int32(0) < v177 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v187 = v142
	goto L35
L34:
	;
	v187 = v167
	goto L35
L35:
	;
	if base.Ui32(v186) < base.Ui32(v187) {
		v140 = v186
		v142 = v187
		goto L25
	} else {
		goto L36
	}
L36:
	;
	goto L26
L37:
	;
	v267 = int32(0)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v267 < v268 {
		goto L47
	} else {
		goto L48
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = int32(0)
	if base.Ui32(v193) < base.Ui32(v195) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v205 = v194
	goto L41
L40:
	;
	v205 = v195
	goto L41
L41:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v129+v206<<(uint(int32(2))%32)) <= base.Ui32(v205) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v217 = v205
	v218 = v206
	goto L43
L43:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v224 = int32(12)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v240 = int32(1)
	v245 = F_tsCompareString(m, l2+int32(8)+v223*v224+int32(base.Ui32(v227)>>(uint(v224)%32)), v227&int32(4095), v129+v218<<(uint(int32(2))%32)+int32(base.Ui32(v236)>>(uint(v224)%32)), int32(base.Ui32(v236)>>(uint(v240)%32))&int32(2047), v240)
	mBase = m.M
	if v245 != 0 {
		goto L37
	} else {
		goto L45
	}
L44:
	;
	goto L37
L45:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v246 + int32(1)
	v251 = v217 + int32(4)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v251) < base.Ui32(v129+v252<<(uint(int32(2))%32)) {
		v217 = v251
		v218 = v252
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v271 = v195
	goto L49
L48:
	;
	v271 = v267
	goto L49
L49:
	;
	goto L20
L50:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if v274 <= int32(0) {
		v475 = v109
		goto L19
	} else {
		goto L51
	}
L51:
	;
	v287 = v271
	v305 = v109
	goto L52
L52:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	if v311&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v475 = v440
	goto L19
L54:
	;
	v440 = base.F32_demote_f64(base.F64_add(base.F64_div(base.F64_promote_f32(base.F32_sub(base.F32_add(v428, v426), base.F32_div(v426, base.F32_convert_i32_s(v429*v429)))), float64(1.64493406685)), base.F64_promote_f32(v305)))
	v442 = v287 + int32(4)
	if (v442-v271)>>(uint(int32(2))%32) < v274 {
		v287 = v442
		v305 = v440
		goto L52
	} else {
		goto L70
	}
L55:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v318 = int32(1)
	v330 = l1 + int32(8) + v314<<(uint(int32(2))%32) + (int32(base.Ui32(v311)>>(uint(v318)%32))&int32(2047)+int32(base.Ui32(v311)>>(uint(int32(12))%32))+v318)&int32(_a_F_calc_rank_0)
	goto L57
L56:
	;
	v330 = v65 + int32(12)
	goto L57
L57:
	;
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v330))))
	if v331 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v426 = float32(-1)
	v428 = float32(0)
	v429 = int32(1)
	goto L54
L59:
	;
	goto L60
L60:
	;
	v341 = int32(0)
	v349 = v341
	v355 = v341
	v372 = float32(-1)
	v374 = float32(0)
	goto L61
L61:
	;
	v378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v330+int32(2)+v349<<(uint(int32(1))%32)))))
	v379 = int32(12)
	v384 = *(*float32)(unsafe.Add(mBase, uint32(l0+int32(base.Ui32(v378)>>(uint(v379)%32))&v379)))
	v385 = base.F32_lt(v372, v384)
	if v385 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v426 = v386
	v428 = v392
	v429 = v393 + int32(1)
	goto L54
L63:
	;
	v386 = v384
	goto L65
L64:
	;
	v386 = v372
	goto L65
L65:
	;
	v388 = v349 + int32(1)
	v392 = base.F32_add(v374, base.F32_div(v384, base.F32_convert_i32_s(v388*v388)))
	if v385 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v393 = v349
	goto L68
L67:
	;
	v393 = v355
	goto L68
L68:
	;
	if v331 != v388 {
		v349 = v388
		v355 = v393
		v372 = v386
		v374 = v392
		goto L61
	} else {
		goto L69
	}
L69:
	;
	goto L62
L70:
	;
	goto L53
L71:
	;
	goto L18
L72:
	;
	m.G0 = v65 + int32(16)
	v1125 = l1
	v1127 = l3
	v1132 = v35
	v1152 = v516
	goto L4
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = int32(1073676289)
	v532 = l0
	v533 = l1
	v534 = l2
	v535 = l3
	v540 = v35
	v551 = v56
	v552 = v5
	v553 = v52
	v554 = l1 + int32(8)
	v559 = v525
	v560 = float32(-1)
	goto L74
L74:
	;
	v564 = int32(2)
	v565 = v552 << (uint(v564) % 32)
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v553+v565)))
	v568 = int32(8)
	v569 = v540 + v568
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = int32(0)
	v580 = v533 + v568
	v583 = v580 + v576<<(uint(v564)%32)
	if base.Ui32(v583) <= base.Ui32(v580) {
		goto L79
	} else {
		goto L80
	}
L75:
	;
	F_pfree(m, v559)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L9
	} else {
		goto L145
	}
L76:
	;
	v1118 = v552 + int32(1)
	if v1118 != v551 {
		v552 = v1118
		v560 = v1113
		goto L74
	} else {
		goto L144
	}
L77:
	;
	if v722 == int32(0) {
		v1113 = v560
		goto L76
	} else {
		goto L107
	}
L78:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+2)))
	if v650 != int32(1) {
		goto L94
	} else {
		goto L95
	}
L79:
	;
	v644 = v580
	v645 = v583
	v646 = v583
	goto L78
L80:
	;
	goto L81
L81:
	;
	v591 = v580
	v593 = v583
	goto L82
L82:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	v598 = int32(12)
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	v608 = int32(2)
	v615 = base.I32_div_s((v593-v591)>>(uint(v608)%32), v608)
	v618 = v591 + v615<<(uint(v608)%32)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	v627 = int32(0)
	v628 = F_tsCompareString(m, v534+int32(8)+v597*v598+int32(base.Ui32(v601)>>(uint(v598)%32)), v601&int32(4095), v580+v607<<(uint(v608)%32)+int32(base.Ui32(v619)>>(uint(v598)%32)), int32(base.Ui32(v619)>>(uint(int32(1))%32))&int32(2047), v627)
	mBase = m.M
	if v628 == v627 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v644 = v637
	v645 = v618
	v646 = v638
	goto L78
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = int32(1)
	v644 = v591
	v645 = v618
	v646 = v618
	goto L78
L85:
	;
	goto L86
L86:
	;
	v636 = base.B2i32(int32(0) < v628)
	if int32(0) < v628 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v637 = v618 + int32(4)
	goto L89
L88:
	;
	v637 = v591
	goto L89
L89:
	;
	if int32(0) < v628 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v638 = v593
	goto L92
L91:
	;
	v638 = v618
	goto L92
L92:
	;
	if base.Ui32(v637) < base.Ui32(v638) {
		v591 = v637
		v593 = v638
		goto L82
	} else {
		goto L93
	}
L93:
	;
	goto L83
L94:
	;
	v718 = int32(0)
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	if v718 < v719 {
		goto L104
	} else {
		goto L105
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = int32(0)
	if base.Ui32(v644) < base.Ui32(v646) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v656 = v645
	goto L98
L97:
	;
	v656 = v646
	goto L98
L98:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if base.Ui32(v580+v657<<(uint(int32(2))%32)) <= base.Ui32(v656) {
		goto L94
	} else {
		goto L99
	}
L99:
	;
	v668 = v656
	v669 = v657
	goto L100
L100:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	v675 = int32(12)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v668)))
	v691 = int32(1)
	v696 = F_tsCompareString(m, v534+int32(8)+v674*v675+int32(base.Ui32(v678)>>(uint(v675)%32)), v678&int32(4095), v580+v669<<(uint(int32(2))%32)+int32(base.Ui32(v687)>>(uint(v675)%32)), int32(base.Ui32(v687)>>(uint(v691)%32))&int32(2047), v691)
	mBase = m.M
	if v696 != 0 {
		goto L94
	} else {
		goto L102
	}
L101:
	;
	goto L94
L102:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = v697 + int32(1)
	v702 = v668 + int32(4)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if base.Ui32(v702) < base.Ui32(v580+v703<<(uint(int32(2))%32)) {
		v668 = v702
		v669 = v703
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v722 = v646
	goto L106
L105:
	;
	v722 = v718
	goto L106
L106:
	;
	goto L77
L107:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v540)+8))
	if v725 <= int32(0) {
		v1113 = v560
		goto L76
	} else {
		goto L108
	}
L108:
	;
	v739 = v722
	v757 = v560
	goto L109
L109:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v739)))
	if v763&int32(1) != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v1113 = v1075
	goto L76
L111:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	v770 = int32(1)
	v782 = v554 + v766<<(uint(int32(2))%32) + (int32(base.Ui32(v763)>>(uint(v770)%32))&int32(2047)+int32(base.Ui32(v763)>>(uint(int32(12))%32))+v770)&int32(_a_F_calc_rank_0)
	goto L113
L112:
	;
	v782 = v540 + int32(12)
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v565+v559))) = v782
	if v552 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v782))))
	v795 = int32(0)
	v816 = v757
	goto L117
L115:
	;
	v1075 = v757
	goto L116
L116:
	;
	v1080 = v739 + int32(4)
	if (v1080-v722)>>(uint(int32(2))%32) < v725 {
		v739 = v1080
		v757 = v1075
		goto L109
	} else {
		goto L143
	}
L117:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v559+v795<<(uint(int32(2))%32))))
	if v823 == int32(0) {
		v1040 = v816
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v1075 = v1040
	goto L116
L119:
	;
	v1045 = v795 + int32(1)
	if v1045 != v552 {
		v795 = v1045
		v816 = v1040
		goto L117
	} else {
		goto L142
	}
L120:
	;
	if v786 == int32(0) {
		v1040 = v816
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v830 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v823))))
	v832 = v540 + int32(12)
	v851 = int32(0)
	v867 = v816
	goto L122
L122:
	;
	if v830 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v1040 = v1005
	goto L119
L124:
	;
	v874 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v782+int32(2)+v851<<(uint(int32(1))%32)))))
	v876 = v874 & int32(_a_F_calc_rank_1)
	v877 = int32(12)
	v887 = int32(0)
	v911 = v867
	goto L127
L125:
	;
	v1005 = v867
	goto L126
L126:
	;
	v1010 = v851 + int32(1)
	if v1010 != v786 {
		v851 = v1010
		v867 = v1005
		goto L122
	} else {
		goto L141
	}
L127:
	;
	v918 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v823+int32(2)+v887<<(uint(int32(1))%32)))))
	v920 = v918 & int32(_a_F_calc_rank_1)
	v921 = base.B2i32(v876 != v920)
	if base.B2i32(v782 == v832)|base.B2i32(v823 == v832)|v921 == int32(0) {
		v972 = v911
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v1005 = v972
	goto L126
L129:
	;
	v975 = v887 + int32(1)
	if v975 != v830 {
		v887 = v975
		v911 = v972
		goto L127
	} else {
		goto L140
	}
L130:
	;
	v925 = *(*float32)(unsafe.Add(mBase, uint32(v532+int32(base.Ui32(v874)>>(uint(v877)%32))&v877)))
	v926 = int32(12)
	v931 = *(*float32)(unsafe.Add(mBase, uint32(v532+int32(base.Ui32(v918)>>(uint(v926)%32))&v926)))
	v933 = v876 - v920
	v935 = v933 >> (uint(int32(31)) % 32)
	if v876 != v920 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v939 = v933 ^ v935 - v935
	goto L133
L132:
	;
	v939 = int32(_a_F_calc_rank_2)
	goto L133
L133:
	;
	if base.Ui32(v939) <= base.Ui32(int32(100)) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v948 = F_exp(m, base.F64_add(base.F64_div(base.F64_convert_i32_u(v939), float64(1.5)), float64(-2)))
	mBase = m.M
	v956 = base.F32_demote_f64(base.F64_div(float64(1), base.F64_add(base.F64_mul(v948, float64(0.05)), float64(1.005))))
	goto L136
L135:
	;
	v956 = float32(1e-30)
	goto L136
L136:
	;
	v958 = base.F32_sqrt(base.F32_mul(base.F32_mul(v925, v931), v956))
	if base.F32_lt(v911, float32(0)) != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v972 = v958
	goto L129
L138:
	;
	goto L139
L139:
	;
	v961 = float64(1)
	v972 = base.F32_demote_f64(base.F64_sub(v961, base.F64_mul(base.F64_sub(v961, base.F64_promote_f32(v911)), base.F64_sub(v961, base.F64_promote_f32(v958)))))
	goto L129
L140:
	;
	goto L128
L141:
	;
	goto L123
L142:
	;
	goto L118
L143:
	;
	goto L110
L144:
	;
	goto L75
L145:
	;
	F_pfree(m, v553)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L9
	} else {
		goto L146
	}
L146:
	;
	v1125 = v533
	v1127 = v535
	v1132 = v540
	v1152 = v1113
	goto L4
L147:
	;
	v1159 = float32(1e-20)
	goto L149
L148:
	;
	v1159 = v1152
	goto L149
L149:
	;
	if v1127&int32(1) == int32(0) {
		v1304 = v1159
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if v1127&int32(2) == int32(0) {
		v1411 = v1304
		goto L165
	} else {
		goto L166
	}
L151:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+4))
	if v1164 <= int32(0) {
		v1304 = v1159
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v1169 = v1125 + int32(8)
	v1172 = v1169 + v1164<<(uint(int32(2))%32)
	if base.Ui32(v1169) < base.Ui32(v1172) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v1179 = v1169
	v1180 = int32(0)
	goto L156
L154:
	;
	v1270 = float64(1)
	goto L155
L155:
	;
	v1271 = F_log(m, v1270)
	mBase = m.M
	v1304 = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(v1159), base.F64_div(v1271, float64(0.6931471805599453))))
	goto L150
L156:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1179)))
	if v1207&int32(1) != 0 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v1270 = base.F64_convert_i32_s(v1230 + int32(1))
	goto L155
L158:
	;
	v1210 = int32(1)
	v1223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1172+(int32(base.Ui32(v1207)>>(uint(v1210)%32))&int32(2047)+int32(base.Ui32(v1207)>>(uint(int32(12))%32))+v1210)&int32(_a_F_calc_rank_0)))))
	if base.Ui32(v1223) <= base.Ui32(v1210) {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	v1229 = int32(1)
	goto L160
L160:
	;
	v1230 = v1229 + v1180
	v1232 = v1179 + int32(4)
	if base.Ui32(v1232) < base.Ui32(v1172) {
		v1179 = v1232
		v1180 = v1230
		goto L156
	} else {
		goto L164
	}
L161:
	;
	v1226 = v1210
	goto L163
L162:
	;
	v1226 = v1223
	goto L163
L163:
	;
	v1229 = v1226
	goto L160
L164:
	;
	goto L157
L165:
	;
	if v1127&int32(8) == int32(0) {
		v1425 = v1411
		goto L178
	} else {
		goto L179
	}
L166:
	;
	v1313 = v1125 + int32(8)
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+4))
	v1317 = v1313 + v1314<<(uint(int32(2))%32)
	if base.Ui32(v1317) <= base.Ui32(v1313) {
		v1411 = v1304
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v1324 = v1313
	v1325 = int32(0)
	goto L168
L168:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1324)))
	if v1352&int32(1) != 0 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v1375 <= int32(0) {
		v1411 = v1304
		goto L165
	} else {
		goto L177
	}
L170:
	;
	v1355 = int32(1)
	v1368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1317+(int32(base.Ui32(v1352)>>(uint(v1355)%32))&int32(2047)+int32(base.Ui32(v1352)>>(uint(int32(12))%32))+v1355)&int32(_a_F_calc_rank_0)))))
	if base.Ui32(v1368) <= base.Ui32(v1355) {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	v1374 = int32(1)
	goto L172
L172:
	;
	v1375 = v1374 + v1325
	v1377 = v1324 + int32(4)
	if base.Ui32(v1377) < base.Ui32(v1317) {
		v1324 = v1377
		v1325 = v1375
		goto L168
	} else {
		goto L176
	}
L173:
	;
	v1371 = v1355
	goto L175
L174:
	;
	v1371 = v1368
	goto L175
L175:
	;
	v1374 = v1371
	goto L172
L176:
	;
	goto L169
L177:
	;
	v1411 = base.F32_div(v1304, base.F32_convert_i32_u(v1375))
	goto L165
L178:
	;
	if v1127&int32(16) == int32(0) {
		v1443 = v1425
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+4))
	if v1419 <= int32(0) {
		v1425 = v1411
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v1425 = base.F32_div(v1411, base.F32_convert_i32_u(v1419))
	goto L178
L181:
	;
	if v1127&int32(32) == int32(0) {
		v1459 = v1132
		v1479 = v1443
		goto L1
	} else {
		goto L184
	}
L182:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+4))
	if v1430 <= int32(0) {
		v1443 = v1425
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v1437 = F_log(m, base.F64_convert_i32_s(v1430+int32(1)))
	mBase = m.M
	v1443 = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(v1425), base.F64_div(v1437, float64(0.6931471805599453))))
	goto L181
L184:
	;
	v1459 = v1132
	v1479 = base.F32_div(v1443, base.F32_add(v1443, float32(1)))
	goto L1
}
func F_call_real_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v14 == int32(0) {
		v81 = v13
		m.G0 = v11 - int32(-64)
		return v81
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[0])) = int32(50856066)
		v21 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[1])) = v21
		*(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[2])) = v21
		*(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[3])) = v21
		v29 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			if v29 != 0 {
				v81 = v13
				m.G0 = v11 - int32(-64)
				return v81
			} else {
				v34 = F_errstart(m, l4, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					if v34 != 0 {
						v37 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[0]))
						F_errcode(m, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[1]))
							if v41 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v41
								F_errmsg_internal(m, int32(_a_F_call_real_check_hook_0), v9+int32(-16))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[2]))
									if v59 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v59
										F_errdetail_internal(m, int32(_a_F_call_real_check_hook_0), v9+int32(-48))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[3]))
											if v67 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
												F_errhint(m, int32(_a_F_call_real_check_hook_0), v11)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_FlushErrorState(m)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															v81 = int32(0)
															m.G0 = v11 - int32(-64)
															return v81
														}
													}
												}
											} else {
												F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[3]))
										if v67 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
											F_errhint(m, int32(_a_F_call_real_check_hook_0), v11)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										} else {
											F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_FlushErrorState(m)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													v81 = int32(0)
													m.G0 = v11 - int32(-64)
													return v81
												}
											}
										}
									}
								}
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+40)) = v49
								*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v48
								F_errmsg(m, int32(_a_F_call_real_check_hook_4), v9+int32(-32))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[2]))
									if v59 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v59
										F_errdetail_internal(m, int32(_a_F_call_real_check_hook_0), v9+int32(-48))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[3]))
											if v67 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
												F_errhint(m, int32(_a_F_call_real_check_hook_0), v11)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_FlushErrorState(m)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															v81 = int32(0)
															m.G0 = v11 - int32(-64)
															return v81
														}
													}
												}
											} else {
												F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, _c_F_call_real_check_hook[3]))
										if v67 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
											F_errhint(m, int32(_a_F_call_real_check_hook_0), v11)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										} else {
											F_errfinish(m, int32(_a_F_call_real_check_hook_1), int32(_a_F_call_real_check_hook_2), int32(_a_F_call_real_check_hook_3))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_FlushErrorState(m)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													v81 = int32(0)
													m.G0 = v11 - int32(-64)
													return v81
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_FlushErrorState(m)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							v81 = int32(0)
							m.G0 = v11 - int32(-64)
							return v81
						}
					}
				}
			}
		}
	}
}
