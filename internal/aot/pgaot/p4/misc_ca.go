package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_calc_hist_selectivity_contained(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) float64 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v94 float64
	_ = v94
	var v97 int32
	_ = v97
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
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v128 float64
	_ = v128
	var v133 float64
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 float64
	_ = v143
	var v145 float64
	_ = v145
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
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v172 int32
	_ = v172
	var v173 float64
	_ = v173
	var v176 float64
	_ = v176
	var v181 float64
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 float64
	_ = v191
	var v193 float64
	_ = v193
	var v194 float64
	_ = v194
	var v196 int32
	_ = v196
	var v201 float64
	_ = v201
	var v212 float64
	_ = v212
	var v214 float64
	_ = v214
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v245 int32
	_ = v245
	var v249 float64
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v269 float64
	_ = v269
	var v270 float64
	_ = v270
	var v271 float64
	_ = v271
	var v272 int32
	_ = v272
	var v273 float64
	_ = v273
	var v274 float64
	_ = v274
	var v276 int32
	_ = v276
	var v290 int32
	_ = v290
	var v299 float64
	_ = v299
	var v303 float64
	_ = v303
	var v305 int32
	_ = v305
	var v311 float64
	_ = v311
	var v314 float64
	_ = v314
	var v315 float64
	_ = v315
	var v323 int32
	_ = v323
	var v328 float64
	_ = v328
	var v330 float64
	_ = v330
	var v331 float64
	_ = v331
	var v337 int32
	_ = v337
	var v341 float64
	_ = v341
	var v348 float64
	_ = v348
	var v351 float64
	_ = v351
	var v361 float64
	_ = v361
	var v366 float64
	_ = v366
	var v369 float64
	_ = v369
	var v370 float64
	_ = v370
	var v371 int32
	_ = v371
	var v372 float64
	_ = v372
	var v374 int32
	_ = v374
	var v388 int32
	_ = v388
	var v397 float64
	_ = v397
	var v401 float64
	_ = v401
	var v402 float64
	_ = v402
	var v407 float64
	_ = v407
	var v414 int32
	_ = v414
	var v417 float64
	_ = v417
	var v419 float64
	_ = v419
	var v420 float64
	_ = v420
	var v422 float64
	_ = v422
	var v426 float64
	_ = v426
	var v430 float64
	_ = v430
	var v440 float64
	_ = v440
	var v460 float64
	_ = v460
	var v470 float64
	_ = v470
	v8 = float64(0)
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)) = uint8(v19)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	v23 = v21 ^ v19
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)) = uint8(v23)
	v27 = l4 - v19
	v40 = v27
	v41 = int32(-1)
	goto L1
L1:
	;
	v50 = base.I32_div_s(v40+v41+int32(1), int32(2))
	v54 = F_range_cmp_bounds(m, l0, l3+v50<<(uint(int32(4))%32), l2)
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
	v60 = v41
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
	v63 = v40
	goto L10
L9:
	;
	v63 = v50 - int32(1)
	goto L10
L10:
	;
	if v60 < v63 {
		v40 = v63
		v41 = v60
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
	if base.Ui32(v60) < base.Ui32(v73) {
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
	v78 = l3 + v75<<(uint(int32(4))%32)
	v81 = F_get_position(m, l0, l2, v78, v78+int32(16))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v91 = v81
	v92 = v8
	v94 = v8
	v97 = v75
	goto L19
L19:
	;
	v103 = l3 + v97<<(uint(int32(4))%32)
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
	return v470
L21:
	;
	v196 = int32(1)
	v201 = float64(0)
	if base.F64_lt(v193, v201) != 0 {
		v460 = v201
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
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
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
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+8)))
	if v156 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L26:
	;
	v150 = F_get_position(m, l0, l1, v103, v103+int32(16))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L46
	}
L27:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
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
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v135 != int32(1) {
		v145 = math.Float64frombits(uint64(0x7ff0000000000000))
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
		v145 = v114
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v115 == int32(0) {
		v145 = v114
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v118 = float64(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v121 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v122 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
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
	v125 = base.F64_reinterpret_i64(v123)
	if base.F64_lt(v125, float64(0)) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v128 = v118
	goto L38
L37:
	;
	v128 = v125
	goto L38
L38:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v123&int64(9223372036854775807)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v133 = v118
	goto L41
L40:
	;
	v133 = v128
	goto L41
L41:
	;
	v145 = v133
	goto L26
L42:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
	if v140 == v141 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v143 = float64(0)
	goto L45
L44:
	;
	v143 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L45
L45:
	;
	v145 = v143
	goto L26
L46:
	;
	v152 = base.F64_sub(v91, v150)
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
	v193 = v145
	v194 = v155
	goto L21
L50:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
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
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v183 != int32(1) {
		v193 = math.Float64frombits(uint64(0x7ff0000000000000))
		v194 = v91
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
		v193 = v162
		v194 = v91
		goto L21
	} else {
		goto L56
	}
L56:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v163 == int32(0) {
		v193 = v162
		v194 = v91
		goto L21
	} else {
		goto L57
	}
L57:
	;
	v166 = float64(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v169 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
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
	v173 = base.F64_reinterpret_i64(v171)
	if base.F64_lt(v173, float64(0)) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v176 = v166
	goto L61
L60:
	;
	v176 = v173
	goto L61
L61:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v171&int64(9223372036854775807)) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v181 = v166
	goto L64
L63:
	;
	v181 = v176
	goto L64
L64:
	;
	v193 = v181
	v194 = v91
	goto L21
L65:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+10)))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
	if v188 == v189 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v191 = float64(0)
	goto L68
L67:
	;
	v191 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L68
L68:
	;
	v193 = v191
	v194 = v91
	goto L21
L69:
	;
	v470 = base.F64_add(v92, base.F64_div(base.F64_mul(v194, v460), base.F64_convert_i32_u(v27)))
	if int32(0) <= v104 {
		goto L141
	} else {
		goto L142
	}
L70:
	;
	goto L69
L71:
	;
	v212 = float64(1)
	v214 = base.F64_abs(v193)
	goto L72
L72:
	;
	goto L74
L74:
	;
	if base.F64_eq(v214, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v460 = v212
		goto L70
	} else {
		goto L75
	}
L75:
	;
	v220 = l6 - int32(1)
	if v220 < int32(0) {
		v460 = v212
		goto L70
	} else {
		goto L76
	}
L76:
	;
	v224 = v220
	v228 = int32(-1)
	goto L77
L77:
	;
	v245 = base.I32_div_s(v224+v228+int32(1), int32(2))
	v249 = *(*float64)(unsafe.Add(mBase, uint32(l5+v245<<(uint(int32(3))%32))))
	if base.F64_gt(v94, v249) != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	if v220 <= v259 {
		v460 = v212
		goto L70
	} else {
		goto L90
	}
L79:
	;
	if v259 < v257 {
		v224 = v257
		v228 = v259
		goto L77
	} else {
		goto L89
	}
L80:
	;
	v257 = v224
	v259 = v245
	goto L79
L81:
	;
	goto L82
L82:
	;
	v254 = v196 & base.F64_ge(v94, v249)
	if v254 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v255 = v224
	goto L85
L84:
	;
	v255 = v245 - int32(1)
	goto L85
L85:
	;
	if v254 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v256 = v245
	goto L88
L87:
	;
	v256 = v228
	goto L88
L88:
	;
	v257 = v255
	v259 = v256
	goto L79
L89:
	;
	goto L78
L90:
	;
	if v259 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v314 = base.F64_convert_i32_u(v220)
	v315 = base.F64_div(base.F64_add(v311, base.F64_convert_i32_u(v305)), v314)
	if base.F64_eq(v94, v193) != 0 {
		v460 = v315
		goto L70
	} else {
		goto L106
	}
L92:
	;
	v305 = int32(0)
	v311 = float64(0)
	goto L91
L93:
	;
	goto L94
L94:
	;
	v268 = l5 + v259<<(uint(int32(3))%32)
	v269 = *(*float64)(unsafe.Add(mBase, uint32(v268)+8))
	v270 = base.F64_abs(v269)
	v271 = math.Float64frombits(uint64(0x7ff0000000000000))
	v272 = base.F64_eq(v270, v271)
	v273 = *(*float64)(unsafe.Add(mBase, uint32(v268)))
	v274 = base.F64_abs(v273)
	v276 = base.F64_eq(v274, v271)
	if v272|v276 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if base.F64_eq(base.F64_abs(v94), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v305 = v259
		v311 = float64(0.5)
		goto L91
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v290 = int32(0)
	if v272|base.B2i32(v276 == v290) == v290 {
		v305 = v259
		v311 = float64(1)
		goto L91
	} else {
		goto L99
	}
L98:
	;
	v305 = v259
	v311 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v269, v94), base.F64_sub(v269, v273)))
	goto L91
L99:
	;
	if base.F64_eq(v270, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v299 = float64(0)
	goto L102
L101:
	;
	v299 = float64(0.5)
	goto L102
L102:
	;
	if base.F64_eq(v274, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v303 = v299
	goto L105
L104:
	;
	v303 = float64(0.5)
	goto L105
L105:
	;
	v305 = v259
	v311 = v303
	goto L91
L106:
	;
	if v220 <= v305 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v426 = float64(0)
	v430 = base.F64_div(base.F64_add(v422, base.F64_convert_i32_u(v414)), v314)
	if base.F64_gt(v419, v426)|base.F64_gt(v430, v426) != 0 {
		goto L134
	} else {
		goto L135
	}
L108:
	;
	v414 = v305
	v417 = v94
	v419 = v315
	v420 = v201
	v422 = v201
	goto L107
L109:
	;
	goto L110
L110:
	;
	v323 = v305
	v328 = v315
	v330 = v201
	v331 = v94
	goto L112
L111:
	;
	v366 = *(*float64)(unsafe.Add(mBase, uint32(l5+v323<<(uint(int32(3))%32))))
	if base.F64_eq(v341, v366) != 0 {
		goto L119
	} else {
		goto L120
	}
L112:
	;
	v337 = v323 + int32(1)
	v341 = *(*float64)(unsafe.Add(mBase, uint32(l5+v337<<(uint(int32(3))%32))))
	if base.F64_lt(v341, v193)|v196&base.F64_ge(v193, v341) == int32(0) {
		goto L111
	} else {
		goto L114
	}
L113:
	;
	v414 = v220
	v417 = v341
	v419 = v351
	v420 = v361
	v422 = v201
	goto L107
L114:
	;
	v348 = float64(0)
	v351 = base.F64_div(base.F64_convert_i32_u(v323), v314)
	if base.F64_gt(v328, v348)|base.F64_gt(v351, v348) != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v361 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v328, v351), float64(0.5)), base.F64_sub(v341, v331)), v330)
	goto L117
L116:
	;
	v361 = v330
	goto L117
L117:
	;
	if v220 != v337 {
		v323 = v337
		v328 = v351
		v330 = v361
		v331 = v341
		goto L112
	} else {
		goto L118
	}
L118:
	;
	goto L113
L119:
	;
	v407 = float64(0)
	goto L121
L120:
	;
	v369 = base.F64_abs(v341)
	v370 = math.Float64frombits(uint64(0x7ff0000000000000))
	v371 = base.F64_eq(v369, v370)
	v372 = base.F64_abs(v366)
	v374 = base.F64_eq(v372, v370)
	if v371|v374 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v414 = v323
	v417 = v331
	v419 = v328
	v420 = v330
	v422 = v407
	goto L107
L122:
	;
	v407 = v402
	goto L121
L123:
	;
	if base.F64_eq(base.F64_abs(v193), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v402 = float64(0.5)
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v388 = int32(0)
	if v371|base.B2i32(v374 == v388) == v388 {
		v402 = float64(1)
		goto L122
	} else {
		goto L127
	}
L126:
	;
	v402 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v341, v193), base.F64_sub(v341, v366)))
	goto L122
L127:
	;
	if base.F64_eq(v369, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v397 = float64(0)
	goto L130
L129:
	;
	v397 = float64(0.5)
	goto L130
L130:
	;
	if base.F64_eq(v372, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v401 = v397
	goto L133
L132:
	;
	v401 = float64(0.5)
	goto L133
L133:
	;
	v402 = v401
	goto L122
L134:
	;
	v440 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v419, v430), float64(0.5)), base.F64_sub(v193, v417)), v420)
	goto L136
L135:
	;
	v440 = v420
	goto L136
L136:
	;
	if base.F64_eq(v214, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	if base.F64_eq(base.F64_abs(v440), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v460 = float64(0.5)
		goto L70
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v460 = base.F64_div(v440, base.F64_sub(v193, v94))
	goto L70
L140:
	;
	goto L139
L141:
	;
	if int32(0) < v97 {
		v91 = float64(1)
		v92 = v470
		v94 = v193
		v97 = v97 - int32(1)
		goto L19
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	goto L20
L144:
	;
	goto L143
}
func F_calc_length_hist_frac(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 int32) float64 {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
	var v21 float64
	_ = v21
	var v23 float64
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v54 int32
	_ = v54
	var v58 float64
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v81 int32
	_ = v81
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v85 int32
	_ = v85
	var v99 int32
	_ = v99
	var v108 float64
	_ = v108
	var v112 float64
	_ = v112
	var v114 int32
	_ = v114
	var v120 float64
	_ = v120
	var v123 float64
	_ = v123
	var v124 float64
	_ = v124
	var v132 int32
	_ = v132
	var v137 float64
	_ = v137
	var v139 float64
	_ = v139
	var v140 float64
	_ = v140
	var v146 int32
	_ = v146
	var v150 float64
	_ = v150
	var v157 float64
	_ = v157
	var v160 float64
	_ = v160
	var v170 float64
	_ = v170
	var v175 float64
	_ = v175
	var v178 float64
	_ = v178
	var v179 float64
	_ = v179
	var v180 int32
	_ = v180
	var v181 float64
	_ = v181
	var v183 int32
	_ = v183
	var v197 int32
	_ = v197
	var v206 float64
	_ = v206
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v216 float64
	_ = v216
	var v223 int32
	_ = v223
	var v226 float64
	_ = v226
	var v228 float64
	_ = v228
	var v229 float64
	_ = v229
	var v231 float64
	_ = v231
	var v235 float64
	_ = v235
	var v239 float64
	_ = v239
	var v249 float64
	_ = v249
	var v269 float64
	_ = v269
	v10 = float64(0)
	if base.F64_lt(l3, v10) != 0 {
		v269 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v269
L2:
	;
	v21 = float64(1)
	v23 = base.F64_abs(l3)
	if base.F64_eq(v23, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = l4
	goto L5
L4:
	;
	v26 = int32(0)
	goto L5
L5:
	;
	if v26 != 0 {
		v269 = v21
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v29 = l1 - int32(1)
	if v29 < int32(0) {
		v269 = v21
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v33 = v29
	v37 = int32(-1)
	goto L8
L8:
	;
	v54 = base.I32_div_s(v33+v37+int32(1), int32(2))
	v58 = *(*float64)(unsafe.Add(mBase, uint32(l0+v54<<(uint(int32(3))%32))))
	if base.F64_gt(l2, v58) != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v29 <= v68 {
		v269 = v21
		goto L1
	} else {
		goto L21
	}
L10:
	;
	if v68 < v66 {
		v33 = v66
		v37 = v68
		goto L8
	} else {
		goto L20
	}
L11:
	;
	v66 = v33
	v68 = v54
	goto L10
L12:
	;
	goto L13
L13:
	;
	v63 = l4 & base.F64_ge(l2, v58)
	if v63 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v64 = v33
	goto L16
L15:
	;
	v64 = v54 - int32(1)
	goto L16
L16:
	;
	if v63 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v65 = v54
	goto L19
L18:
	;
	v65 = v37
	goto L19
L19:
	;
	v66 = v64
	v68 = v65
	goto L10
L20:
	;
	goto L9
L21:
	;
	if v68 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v123 = base.F64_convert_i32_u(v29)
	v124 = base.F64_div(base.F64_add(v120, base.F64_convert_i32_u(v114)), v123)
	if base.F64_eq(l2, l3) != 0 {
		v269 = v124
		goto L1
	} else {
		goto L37
	}
L23:
	;
	v114 = int32(0)
	v120 = float64(0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v77 = l0 + v68<<(uint(int32(3))%32)
	v78 = *(*float64)(unsafe.Add(mBase, uint32(v77)+8))
	v79 = base.F64_abs(v78)
	v80 = math.Float64frombits(uint64(0x7ff0000000000000))
	v81 = base.F64_eq(v79, v80)
	v82 = *(*float64)(unsafe.Add(mBase, uint32(v77)))
	v83 = base.F64_abs(v82)
	v85 = base.F64_eq(v83, v80)
	if v81|v85 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if base.F64_eq(base.F64_abs(l2), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v114 = v68
		v120 = float64(0.5)
		goto L22
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v99 = int32(0)
	if v81|base.B2i32(v85 == v99) == v99 {
		v114 = v68
		v120 = float64(1)
		goto L22
	} else {
		goto L30
	}
L29:
	;
	v114 = v68
	v120 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v78, l2), base.F64_sub(v78, v82)))
	goto L22
L30:
	;
	if base.F64_eq(v79, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v108 = float64(0)
	goto L33
L32:
	;
	v108 = float64(0.5)
	goto L33
L33:
	;
	if base.F64_eq(v83, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v112 = v108
	goto L36
L35:
	;
	v112 = float64(0.5)
	goto L36
L36:
	;
	v114 = v68
	v120 = v112
	goto L22
L37:
	;
	if v29 <= v114 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v235 = float64(0)
	v239 = base.F64_div(base.F64_add(v231, base.F64_convert_i32_u(v223)), v123)
	if base.F64_gt(v228, v235)|base.F64_gt(v239, v235) != 0 {
		goto L65
	} else {
		goto L66
	}
L39:
	;
	v223 = v114
	v226 = l2
	v228 = v124
	v229 = v10
	v231 = v10
	goto L38
L40:
	;
	goto L41
L41:
	;
	v132 = v114
	v137 = v124
	v139 = v10
	v140 = l2
	goto L43
L42:
	;
	v175 = *(*float64)(unsafe.Add(mBase, uint32(l0+v132<<(uint(int32(3))%32))))
	if base.F64_eq(v150, v175) != 0 {
		goto L50
	} else {
		goto L51
	}
L43:
	;
	v146 = v132 + int32(1)
	v150 = *(*float64)(unsafe.Add(mBase, uint32(l0+v146<<(uint(int32(3))%32))))
	if base.F64_lt(v150, l3)|l4&base.F64_ge(l3, v150) == int32(0) {
		goto L42
	} else {
		goto L45
	}
L44:
	;
	v223 = v29
	v226 = v150
	v228 = v160
	v229 = v170
	v231 = v10
	goto L38
L45:
	;
	v157 = float64(0)
	v160 = base.F64_div(base.F64_convert_i32_u(v132), v123)
	if base.F64_gt(v137, v157)|base.F64_gt(v160, v157) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v170 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v137, v160), float64(0.5)), base.F64_sub(v150, v140)), v139)
	goto L48
L47:
	;
	v170 = v139
	goto L48
L48:
	;
	if v29 != v146 {
		v132 = v146
		v137 = v160
		v139 = v170
		v140 = v150
		goto L43
	} else {
		goto L49
	}
L49:
	;
	goto L44
L50:
	;
	v216 = float64(0)
	goto L52
L51:
	;
	v178 = base.F64_abs(v150)
	v179 = math.Float64frombits(uint64(0x7ff0000000000000))
	v180 = base.F64_eq(v178, v179)
	v181 = base.F64_abs(v175)
	v183 = base.F64_eq(v181, v179)
	if v180|v183 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v223 = v132
	v226 = v140
	v228 = v137
	v229 = v139
	v231 = v216
	goto L38
L53:
	;
	v216 = v211
	goto L52
L54:
	;
	if base.F64_eq(base.F64_abs(l3), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v211 = float64(0.5)
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v197 = int32(0)
	if v180|base.B2i32(v183 == v197) == v197 {
		v211 = float64(1)
		goto L53
	} else {
		goto L58
	}
L57:
	;
	v211 = base.F64_sub(float64(1), base.F64_div(base.F64_sub(v150, l3), base.F64_sub(v150, v175)))
	goto L53
L58:
	;
	if base.F64_eq(v178, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v206 = float64(0)
	goto L61
L60:
	;
	v206 = float64(0.5)
	goto L61
L61:
	;
	if base.F64_eq(v181, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v210 = v206
	goto L64
L63:
	;
	v210 = float64(0.5)
	goto L64
L64:
	;
	v211 = v210
	goto L53
L65:
	;
	v249 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v228, v239), float64(0.5)), base.F64_sub(l3, v226)), v229)
	goto L67
L66:
	;
	v249 = v229
	goto L67
L67:
	;
	if base.F64_eq(v23, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if base.F64_eq(base.F64_abs(v249), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v269 = float64(0.5)
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v269 = base.F64_div(v249, base.F64_sub(l3, l2))
	goto L1
L71:
	;
	goto L70
}
func F_calc_non_nestloop_required_outer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		v6 = v4
	} else {
		v6 = int32(0)
	}
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v7 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		v10 = v8
	} else {
		v10 = int32(0)
	}
	v11 = F_bms_union(m, v6, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
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
	var v95 int32
	_ = v95
	var v114 float32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v293 int32
	_ = v293
	var v311 float32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 float32
	_ = v341
	var v342 float32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v373 float32
	_ = v373
	var v374 float32
	_ = v374
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 float32
	_ = v386
	var v387 int32
	_ = v387
	var v388 float32
	_ = v388
	var v390 int32
	_ = v390
	var v394 float32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v430 float32
	_ = v430
	var v431 float32
	_ = v431
	var v434 float32
	_ = v434
	var v443 float32
	_ = v443
	var v445 int32
	_ = v445
	var v481 float32
	_ = v481
	var v483 int32
	_ = v483
	var v519 float32
	_ = v519
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v562 float32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v736 int32
	_ = v736
	var v760 float32
	_ = v760
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v802 int32
	_ = v802
	var v819 float32
	_ = v819
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v855 int32
	_ = v855
	var v871 float32
	_ = v871
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v892 int32
	_ = v892
	var v915 float32
	_ = v915
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v929 float32
	_ = v929
	var v930 int32
	_ = v930
	var v935 float32
	_ = v935
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v952 float64
	_ = v952
	var v960 float32
	_ = v960
	var v962 float32
	_ = v962
	var v965 float64
	_ = v965
	var v976 float32
	_ = v976
	var v979 int32
	_ = v979
	var v1009 float32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1044 float32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1079 float32
	_ = v1079
	var v1084 int32
	_ = v1084
	var v1117 float32
	_ = v1117
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1137 int32
	_ = v1137
	var v1156 float32
	_ = v1156
	var v1163 float32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1240 float64
	_ = v1240
	var v1273 float32
	_ = v1273
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1295 int32
	_ = v1295
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1381 float32
	_ = v1381
	var v1389 int32
	_ = v1389
	var v1395 float32
	_ = v1395
	var v1400 int32
	_ = v1400
	var v1407 float64
	_ = v1407
	var v1413 float32
	_ = v1413
	var v1430 int32
	_ = v1430
	var v1449 float32
	_ = v1449
	v5 = int32(0)
	v29 = float32(0)
	v33 = m.G0
	v35 = v33 - int32(16)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v37 == v5 {
		v1430 = v35
		v1449 = v29
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v1430 + int32(16)
	return v1449
L2:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v40 == int32(0) {
		v1430 = v35
		v1449 = v29
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
	if base.F32_lt(v1156, float32(0)) != 0 {
		goto L147
	} else {
		goto L148
	}
L5:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v527 = F_palloc0_mul(m, int32(4), v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
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
	v519 = float32(0)
	goto L16
L15:
	;
	v95 = v5
	v114 = v29
	goto L17
L16:
	;
	F_pfree(m, v73)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L9
	} else {
		goto L72
	}
L17:
	;
	v115 = int32(2)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v73+v95<<(uint(v115)%32))))
	v119 = int32(8)
	v120 = v65 + v119
	v121 = int32(0)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v121
	v131 = l1 + v119
	v134 = v131 + v127<<(uint(v115)%32)
	if v127 <= v121 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v519 = base.F32_div(v481, base.F32_convert_i32_u(v75))
	goto L16
L19:
	;
	v483 = v95 + int32(1)
	if v483 != v75 {
		v95 = v483
		v114 = v481
		goto L17
	} else {
		goto L71
	}
L20:
	;
	if v274 == int32(0) {
		v481 = v114
		goto L19
	} else {
		goto L50
	}
L21:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
	if v202 != int32(1) {
		goto L37
	} else {
		goto L38
	}
L22:
	;
	v196 = v134
	v197 = v131
	v198 = v134
	goto L21
L23:
	;
	goto L24
L24:
	;
	v144 = v131
	v145 = v134
	goto L25
L25:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v150 = int32(12)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v160 = int32(2)
	v167 = base.I32_div_s((v145-v144)>>(uint(v160)%32), v160)
	v170 = v144 + v167<<(uint(v160)%32)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v179 = int32(0)
	v180 = F_tsCompareString(m, l2+int32(8)+v149*v150+int32(base.Ui32(v153)>>(uint(v150)%32)), v153&int32(4095), v131+v159<<(uint(v160)%32)+int32(base.Ui32(v171)>>(uint(v150)%32)), int32(base.Ui32(v171)>>(uint(int32(1))%32))&int32(2047), v179)
	mBase = m.M
	if v180 == v179 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v196 = v170
	v197 = v189
	v198 = v190
	goto L21
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = int32(1)
	v196 = v170
	v197 = v144
	v198 = v170
	goto L21
L28:
	;
	goto L29
L29:
	;
	v188 = base.B2i32(int32(0) < v180)
	if int32(0) < v180 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v189 = v170 + int32(4)
	goto L32
L31:
	;
	v189 = v144
	goto L32
L32:
	;
	if int32(0) < v180 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v190 = v145
	goto L35
L34:
	;
	v190 = v170
	goto L35
L35:
	;
	if base.Ui32(v189) < base.Ui32(v190) {
		v144 = v189
		v145 = v190
		goto L25
	} else {
		goto L36
	}
L36:
	;
	goto L26
L37:
	;
	v270 = int32(0)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v270 < v271 {
		goto L47
	} else {
		goto L48
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = int32(0)
	if base.Ui32(v197) < base.Ui32(v198) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v208 = v196
	goto L41
L40:
	;
	v208 = v198
	goto L41
L41:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v131+v209<<(uint(int32(2))%32)) <= base.Ui32(v208) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v220 = v209
	v221 = v208
	goto L43
L43:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v227 = int32(12)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v243 = int32(1)
	v248 = F_tsCompareString(m, l2+int32(8)+v226*v227+int32(base.Ui32(v230)>>(uint(v227)%32)), v230&int32(4095), v131+v220<<(uint(int32(2))%32)+int32(base.Ui32(v239)>>(uint(v227)%32)), int32(base.Ui32(v239)>>(uint(v243)%32))&int32(2047), v243)
	mBase = m.M
	if v248 != 0 {
		goto L37
	} else {
		goto L45
	}
L44:
	;
	goto L37
L45:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v249 + int32(1)
	v254 = v221 + int32(4)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v254) < base.Ui32(v131+v255<<(uint(int32(2))%32)) {
		v220 = v255
		v221 = v254
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v274 = v198
	goto L49
L48:
	;
	v274 = v270
	goto L49
L49:
	;
	goto L20
L50:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if v277 <= int32(0) {
		v481 = v114
		goto L19
	} else {
		goto L51
	}
L51:
	;
	v293 = v274
	v311 = v114
	goto L52
L52:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	if v314&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v481 = v443
	goto L19
L54:
	;
	v317 = int32(1)
	v327 = (int32(base.Ui32(v314)>>(uint(v317)%32))&int32(2047) + int32(base.Ui32(v314)>>(uint(int32(12))%32)) + v317) & int32(_a_F_calc_rank_0)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v330 = v328 << (uint(int32(2)) % 32)
	v338 = l1 + v330 + v327 + int32(10)
	v339 = v327 + (l1 + int32(8) + v330)
	goto L56
L55:
	;
	v338 = v65 + int32(14)
	v339 = v65 + int32(12)
	goto L56
L56:
	;
	v341 = float32(-1)
	v342 = float32(0)
	v343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v339))))
	if v343 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v443 = base.F32_demote_f64(base.F64_add(base.F64_div(base.F64_promote_f32(base.F32_add(v431, base.F32_sub(v430, base.F32_div(v431, v434)))), float64(1.64493406685)), base.F64_promote_f32(v311)))
	v445 = v293 + int32(4)
	if (v445-v274)>>(uint(int32(2))%32) < v277 {
		v293 = v445
		v311 = v443
		goto L52
	} else {
		goto L70
	}
L58:
	;
	v349 = int32(0)
	v356 = int32(0)
	v373 = v342
	v374 = v341
	goto L61
L59:
	;
	goto L60
L60:
	;
	v430 = v342
	v431 = v341
	v434 = float32(1)
	goto L57
L61:
	;
	v380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v338+v349<<(uint(int32(1))%32)))))
	v381 = int32(12)
	v386 = *(*float32)(unsafe.Add(mBase, uint32(l0+int32(base.Ui32(v380)>>(uint(v381)%32))&v381)))
	v387 = base.F32_lt(v374, v386)
	if v387 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v398 = v395 + int32(1)
	v430 = v394
	v431 = v388
	v434 = base.F32_convert_i32_s(v398 * v398)
	goto L57
L63:
	;
	v388 = v386
	goto L65
L64:
	;
	v388 = v374
	goto L65
L65:
	;
	v390 = v349 + int32(1)
	v394 = base.F32_add(v373, base.F32_div(v386, base.F32_convert_i32_s(v390*v390)))
	if v387 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v395 = v349
	goto L68
L67:
	;
	v395 = v356
	goto L68
L68:
	;
	if v343 != v390 {
		v349 = v390
		v356 = v395
		v373 = v394
		v374 = v388
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
	v1129 = l1
	v1131 = l3
	v1137 = v35
	v1156 = v519
	goto L4
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = int32(1073676289)
	v534 = l0
	v535 = l1
	v536 = l2
	v537 = l3
	v540 = v527
	v542 = v5
	v543 = v35
	v548 = v52
	v552 = v56
	v553 = l1 + int32(8)
	v562 = float32(-1)
	goto L74
L74:
	;
	v566 = int32(2)
	v567 = v542 << (uint(v566) % 32)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v548+v567)))
	v570 = int32(8)
	v571 = v543 + v570
	v572 = int32(0)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = v572
	v582 = v535 + v570
	v585 = v582 + v578<<(uint(v566)%32)
	if v578 <= v572 {
		goto L79
	} else {
		goto L80
	}
L75:
	;
	F_pfree(m, v540)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L9
	} else {
		goto L145
	}
L76:
	;
	v1122 = v542 + int32(1)
	if v1122 != v552 {
		v542 = v1122
		v562 = v1117
		goto L74
	} else {
		goto L144
	}
L77:
	;
	if v725 == int32(0) {
		v1117 = v562
		goto L76
	} else {
		goto L107
	}
L78:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+2)))
	if v653 != int32(1) {
		goto L94
	} else {
		goto L95
	}
L79:
	;
	v647 = v585
	v648 = v582
	v649 = v585
	goto L78
L80:
	;
	goto L81
L81:
	;
	v595 = v582
	v596 = v585
	goto L82
L82:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	v601 = int32(12)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v569)+8))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	v611 = int32(2)
	v618 = base.I32_div_s((v596-v595)>>(uint(v611)%32), v611)
	v621 = v595 + v618<<(uint(v611)%32)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v621)))
	v630 = int32(0)
	v631 = F_tsCompareString(m, v536+int32(8)+v600*v601+int32(base.Ui32(v604)>>(uint(v601)%32)), v604&int32(4095), v582+v610<<(uint(v611)%32)+int32(base.Ui32(v622)>>(uint(v601)%32)), int32(base.Ui32(v622)>>(uint(int32(1))%32))&int32(2047), v630)
	mBase = m.M
	if v631 == v630 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v647 = v621
	v648 = v640
	v649 = v641
	goto L78
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = int32(1)
	v647 = v621
	v648 = v595
	v649 = v621
	goto L78
L85:
	;
	goto L86
L86:
	;
	v639 = base.B2i32(int32(0) < v631)
	if int32(0) < v631 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v640 = v621 + int32(4)
	goto L89
L88:
	;
	v640 = v595
	goto L89
L89:
	;
	if int32(0) < v631 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v641 = v596
	goto L92
L91:
	;
	v641 = v621
	goto L92
L92:
	;
	if base.Ui32(v640) < base.Ui32(v641) {
		v595 = v640
		v596 = v641
		goto L82
	} else {
		goto L93
	}
L93:
	;
	goto L83
L94:
	;
	v721 = int32(0)
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v571)))
	if v721 < v722 {
		goto L104
	} else {
		goto L105
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = int32(0)
	if base.Ui32(v648) < base.Ui32(v649) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v659 = v647
	goto L98
L97:
	;
	v659 = v649
	goto L98
L98:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	if base.Ui32(v582+v660<<(uint(int32(2))%32)) <= base.Ui32(v659) {
		goto L94
	} else {
		goto L99
	}
L99:
	;
	v671 = v660
	v672 = v659
	goto L100
L100:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	v678 = int32(12)
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v569)+8))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v672)))
	v694 = int32(1)
	v699 = F_tsCompareString(m, v536+int32(8)+v677*v678+int32(base.Ui32(v681)>>(uint(v678)%32)), v681&int32(4095), v582+v671<<(uint(int32(2))%32)+int32(base.Ui32(v690)>>(uint(v678)%32)), int32(base.Ui32(v690)>>(uint(v694)%32))&int32(2047), v694)
	mBase = m.M
	if v699 != 0 {
		goto L94
	} else {
		goto L102
	}
L101:
	;
	goto L94
L102:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v571)))
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = v700 + int32(1)
	v705 = v672 + int32(4)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	if base.Ui32(v705) < base.Ui32(v582+v706<<(uint(int32(2))%32)) {
		v671 = v706
		v672 = v705
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v725 = v649
	goto L106
L105:
	;
	v725 = v721
	goto L106
L106:
	;
	goto L77
L107:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v543)+8))
	if v728 <= int32(0) {
		v1117 = v562
		goto L76
	} else {
		goto L108
	}
L108:
	;
	v736 = v725
	v760 = v562
	goto L109
L109:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v736)))
	if v766&int32(1) != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v1117 = v1079
	goto L76
L111:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	v773 = int32(1)
	v785 = v553 + v769<<(uint(int32(2))%32) + (int32(base.Ui32(v766)>>(uint(v773)%32))&int32(2047)+int32(base.Ui32(v766)>>(uint(int32(12))%32))+v773)&int32(_a_F_calc_rank_0)
	goto L113
L112:
	;
	v785 = v543 + int32(12)
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567+v540))) = v785
	if v542 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v789 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v785))))
	v802 = int32(0)
	v819 = v760
	goto L117
L115:
	;
	v1079 = v760
	goto L116
L116:
	;
	v1084 = v736 + int32(4)
	if (v1084-v725)>>(uint(int32(2))%32) < v728 {
		v736 = v1084
		v760 = v1079
		goto L109
	} else {
		goto L143
	}
L117:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v540+v802<<(uint(int32(2))%32))))
	v827 = int32(0)
	if base.B2i32(v826 == v827)|base.B2i32(v789 == v827) == v827 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v1079 = v1044
	goto L116
L119:
	;
	v836 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v826))))
	v838 = v543 + int32(12)
	v855 = int32(0)
	v871 = v819
	goto L122
L120:
	;
	v1044 = v819
	goto L121
L121:
	;
	v1049 = v802 + int32(1)
	if v1049 != v542 {
		v802 = v1049
		v819 = v1044
		goto L117
	} else {
		goto L142
	}
L122:
	;
	if v836 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v1044 = v1009
	goto L121
L124:
	;
	v878 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v785+int32(2)+v855<<(uint(int32(1))%32)))))
	v880 = v878 & int32(_a_F_calc_rank_1)
	v881 = int32(12)
	v892 = int32(0)
	v915 = v871
	goto L127
L125:
	;
	v1009 = v871
	goto L126
L126:
	;
	v1014 = v855 + int32(1)
	if v1014 != v789 {
		v855 = v1014
		v871 = v1009
		goto L122
	} else {
		goto L141
	}
L127:
	;
	v922 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v826+int32(2)+v892<<(uint(int32(1))%32)))))
	v924 = v922 & int32(_a_F_calc_rank_1)
	v925 = base.B2i32(v880 != v924)
	if base.B2i32(v785 == v838)|base.B2i32(v826 == v838)|v925 == int32(0) {
		v976 = v915
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v1009 = v976
	goto L126
L129:
	;
	v979 = v892 + int32(1)
	if v979 != v836 {
		v892 = v979
		v915 = v976
		goto L127
	} else {
		goto L140
	}
L130:
	;
	v929 = *(*float32)(unsafe.Add(mBase, uint32(v534+int32(base.Ui32(v878)>>(uint(v881)%32))&v881)))
	v930 = int32(12)
	v935 = *(*float32)(unsafe.Add(mBase, uint32(v534+int32(base.Ui32(v922)>>(uint(v930)%32))&v930)))
	v937 = v880 - v924
	v939 = v937 >> (uint(int32(31)) % 32)
	if v880 != v924 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v943 = v937 ^ v939 - v939
	goto L133
L132:
	;
	v943 = int32(_a_F_calc_rank_2)
	goto L133
L133:
	;
	if base.Ui32(v943) <= base.Ui32(int32(100)) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v952 = F_exp(m, base.F64_add(base.F64_div(base.F64_convert_i32_u(v943), float64(1.5)), float64(-2)))
	mBase = m.M
	v960 = base.F32_demote_f64(base.F64_div(float64(1), base.F64_add(base.F64_mul(v952, float64(0.05)), float64(1.005))))
	goto L136
L135:
	;
	v960 = float32(1e-30)
	goto L136
L136:
	;
	v962 = base.F32_sqrt(base.F32_mul(base.F32_mul(v929, v935), v960))
	if base.F32_lt(v915, float32(0)) != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v976 = v962
	goto L129
L138:
	;
	goto L139
L139:
	;
	v965 = float64(1)
	v976 = base.F32_demote_f64(base.F64_sub(v965, base.F64_mul(base.F64_sub(v965, base.F64_promote_f32(v915)), base.F64_sub(v965, base.F64_promote_f32(v962)))))
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
	F_pfree(m, v548)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L9
	} else {
		goto L146
	}
L146:
	;
	v1129 = v535
	v1131 = v537
	v1137 = v543
	v1156 = v1117
	goto L4
L147:
	;
	v1163 = float32(1e-20)
	goto L149
L148:
	;
	v1163 = v1156
	goto L149
L149:
	;
	if v1131&int32(1) == int32(0) {
		v1273 = v1163
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if v1131&int32(2) == int32(0) {
		v1381 = v1273
		goto L162
	} else {
		goto L163
	}
L151:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+4))
	if v1168 <= int32(0) {
		v1273 = v1163
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v1172 = v1129 + int32(8)
	v1175 = v1172 + v1168<<(uint(int32(2))%32)
	v1177 = int32(0)
	v1182 = v1172
	goto L153
L153:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1182)))
	if v1209&int32(1) != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v1240 = F_log(m, base.F64_convert_i32_s(v1232+int32(1)))
	mBase = m.M
	v1273 = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(v1163), base.F64_div(v1240, float64(0.6931471805599453))))
	goto L150
L155:
	;
	v1212 = int32(1)
	v1225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1175+(int32(base.Ui32(v1209)>>(uint(v1212)%32))&int32(2047)+int32(base.Ui32(v1209)>>(uint(int32(12))%32))+v1212)&int32(_a_F_calc_rank_0)))))
	if base.Ui32(v1225) <= base.Ui32(v1212) {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v1231 = int32(1)
	goto L157
L157:
	;
	v1232 = v1231 + v1177
	v1234 = v1182 + int32(4)
	if base.Ui32(v1234) < base.Ui32(v1175) {
		v1177 = v1232
		v1182 = v1234
		goto L153
	} else {
		goto L161
	}
L158:
	;
	v1228 = v1212
	goto L160
L159:
	;
	v1228 = v1225
	goto L160
L160:
	;
	v1231 = v1228
	goto L157
L161:
	;
	goto L154
L162:
	;
	if v1131&int32(8) == int32(0) {
		v1395 = v1381
		goto L175
	} else {
		goto L176
	}
L163:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+4))
	if v1281 <= int32(0) {
		v1381 = v1273
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v1285 = v1129 + int32(8)
	v1288 = v1285 + v1281<<(uint(int32(2))%32)
	v1290 = int32(0)
	v1295 = v1285
	goto L165
L165:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	if v1322&int32(1) != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	if v1345 <= int32(0) {
		v1381 = v1273
		goto L162
	} else {
		goto L174
	}
L167:
	;
	v1325 = int32(1)
	v1338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1288+(int32(base.Ui32(v1322)>>(uint(v1325)%32))&int32(2047)+int32(base.Ui32(v1322)>>(uint(int32(12))%32))+v1325)&int32(_a_F_calc_rank_0)))))
	if base.Ui32(v1338) <= base.Ui32(v1325) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v1344 = int32(1)
	goto L169
L169:
	;
	v1345 = v1344 + v1290
	v1347 = v1295 + int32(4)
	if base.Ui32(v1347) < base.Ui32(v1288) {
		v1290 = v1345
		v1295 = v1347
		goto L165
	} else {
		goto L173
	}
L170:
	;
	v1341 = v1325
	goto L172
L171:
	;
	v1341 = v1338
	goto L172
L172:
	;
	v1344 = v1341
	goto L169
L173:
	;
	goto L166
L174:
	;
	v1381 = base.F32_div(v1273, base.F32_convert_i32_u(v1345))
	goto L162
L175:
	;
	if v1131&int32(16) == int32(0) {
		v1413 = v1395
		goto L178
	} else {
		goto L179
	}
L176:
	;
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+4))
	if v1389 <= int32(0) {
		v1395 = v1381
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v1395 = base.F32_div(v1381, base.F32_convert_i32_u(v1389))
	goto L175
L178:
	;
	if v1131&int32(32) == int32(0) {
		v1430 = v1137
		v1449 = v1413
		goto L1
	} else {
		goto L181
	}
L179:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+4))
	if v1400 <= int32(0) {
		v1413 = v1395
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v1407 = F_log(m, base.F64_convert_i32_s(v1400+int32(1)))
	mBase = m.M
	v1413 = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(v1395), base.F64_div(v1407, float64(0.6931471805599453))))
	goto L178
L181:
	;
	v1430 = v1137
	v1449 = base.F32_div(v1413, base.F32_add(v1413, float32(1)))
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
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v14 == int32(0) {
		v84 = v13
		m.G0 = v11 - int32(-64)
		return v84
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
				v84 = v13
				m.G0 = v11 - int32(-64)
				return v84
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
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															v84 = int32(0)
															m.G0 = v11 - int32(-64)
															return v84
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
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														v84 = int32(0)
														m.G0 = v11 - int32(-64)
														return v84
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
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														v84 = int32(0)
														m.G0 = v11 - int32(-64)
														return v84
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
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													v84 = int32(0)
													m.G0 = v11 - int32(-64)
													return v84
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
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															v84 = int32(0)
															m.G0 = v11 - int32(-64)
															return v84
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
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														v84 = int32(0)
														m.G0 = v11 - int32(-64)
														return v84
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
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														v84 = int32(0)
														m.G0 = v11 - int32(-64)
														return v84
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
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													v84 = int32(0)
													m.G0 = v11 - int32(-64)
													return v84
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
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v84 = int32(0)
							m.G0 = v11 - int32(-64)
							return v84
						}
					}
				}
			}
		}
	}
}
