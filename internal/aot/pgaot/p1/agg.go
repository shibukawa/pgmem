package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cost_agg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64, l6 int32, l7 int32, l8 float64, l9 float64, l10 float64, l11 float64) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v30 float64
	_ = v30
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v63 float64
	_ = v63
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 float64
	_ = v87
	var v93 float64
	_ = v93
	var v96 float64
	_ = v96
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v103 float64
	_ = v103
	var v106 float64
	_ = v106
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 float64
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 float64
	_ = v176
	var v189 int32
	_ = v189
	var v190 float64
	_ = v190
	var v196 float64
	_ = v196
	var v202 float64
	_ = v202
	var v204 float64
	_ = v204
	var v207 float64
	_ = v207
	var v210 float64
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v233 float64
	_ = v233
	var v238 int64
	_ = v238
	var v245 int32
	_ = v245
	var v258 int32
	_ = v258
	var v260 float64
	_ = v260
	var v261 int64
	_ = v261
	var v263 float64
	_ = v263
	var v265 float64
	_ = v265
	var v266 float64
	_ = v266
	var v267 float64
	_ = v267
	var v270 float64
	_ = v270
	var v271 float64
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 float64
	_ = v277
	var v281 float64
	_ = v281
	var v282 float64
	_ = v282
	var v283 float64
	_ = v283
	var v285 float64
	_ = v285
	var v286 float64
	_ = v286
	var v289 float64
	_ = v289
	var v290 float64
	_ = v290
	var v292 float64
	_ = v292
	var v295 float64
	_ = v295
	var v302 float64
	_ = v302
	var v303 int32
	_ = v303
	var v304 float64
	_ = v304
	var v311 float64
	_ = v311
	var v314 int64
	_ = v314
	var v319 int32
	_ = v319
	var v322 float64
	_ = v322
	var v328 int32
	_ = v328
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 float64
	_ = v355
	var v356 float64
	_ = v356
	var v368 float64
	_ = v368
	var v374 float64
	_ = v374
	var v375 float64
	_ = v375
	var v378 float64
	_ = v378
	var v379 int32
	_ = v379
	var v382 float64
	_ = v382
	var v383 int32
	_ = v383
	var v384 float64
	_ = v384
	var v394 float64
	_ = v394
	var v403 float64
	_ = v403
	var v406 float64
	_ = v406
	var v407 float64
	_ = v407
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l6 == int32(0) {
		v403 = v302
		v406 = v304
		v407 = v311
		goto L81
	} else {
		goto L82
	}
L2:
	;
	if l3 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if l2&int32(-3) == int32(1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v25 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v27 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	v30 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v32 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v41 = base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v25, l10), base.F64_add(l9, v27)), v30), v32)
	goto L7
L6:
	;
	v34 = float64(0)
	v41 = base.F64_add(base.F64_add(base.F64_mul(l10, v34), base.F64_add(l9, v34)), v34)
	goto L7
L7:
	;
	v43 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[0]))
	v302 = float64(1)
	v303 = l7
	v304 = v41
	v311 = base.F64_add(v41, v43)
	goto L1
L8:
	;
	v123 = v122 + l7
	v126 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[0]))
	v128 = base.F64_add(base.F64_mul(v126, l5), v119)
	if l2&int32(-2) != int32(2) {
		v302 = l5
		v303 = v123
		v304 = v118
		v311 = v128
		goto L1
	} else {
		goto L30
	}
L9:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_agg[1])))
	if l3 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_agg[1])))
	v87 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[2]))
	if l3 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	v51 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v52 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	v56 = v51
	v57 = base.F64_add(l9, v52)
	goto L14
L13:
	;
	v56 = float64(0)
	v57 = base.F64_add(l9, float64(0))
	goto L14
L14:
	;
	v63 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[2]))
	v69 = base.F64_add(base.F64_mul(base.F64_mul(v63, base.F64_convert_i32_s(l4)), l10), base.F64_add(base.F64_mul(v56, l10), v57))
	if l3 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if l2 == int32(3) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v72 = float64(0)
	v78 = v72
	v79 = base.F64_add(v69, v72)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v75 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v76 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v78 = v75
	v79 = base.F64_add(v69, v76)
	goto L15
L19:
	;
	v83 = v50 ^ int32(1)
	goto L21
L20:
	;
	v83 = int32(0)
	goto L21
L21:
	;
	v118 = l8
	v119 = base.F64_add(base.F64_mul(v78, l5), v79)
	v122 = v83
	goto L8
L22:
	;
	v103 = base.F64_add(base.F64_mul(base.F64_mul(v87, base.F64_convert_i32_s(l4)), l10), base.F64_add(base.F64_mul(v100, l10), v99))
	if l3 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v93 = float64(0)
	v99 = base.F64_add(l9, v93)
	v100 = v93
	goto L22
L24:
	;
	goto L25
L25:
	;
	v96 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	v98 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v99 = base.F64_add(l9, v96)
	v100 = v98
	goto L22
L26:
	;
	v118 = v113
	v119 = base.F64_add(base.F64_mul(v112, l5), v113)
	v122 = v85 ^ int32(1)
	goto L8
L27:
	;
	v106 = float64(0)
	v112 = v106
	v113 = base.F64_add(v103, v106)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v109 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v110 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v112 = v109
	v113 = base.F64_add(v103, v110)
	goto L26
L30:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+332))
	if v133 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v136 = v134
	goto L33
L32:
	;
	v136 = int32(0)
	goto L33
L33:
	;
	if l3 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v140 = v138
	goto L36
L35:
	;
	v140 = int32(0)
	goto L36
L36:
	;
	v148 = int32(1)
	if v140&(v140-v148) != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v164 = base.F64_convert_i32_u((base.I32_trunc_sat_f64_u(l11)+int32(23))&int32(-8) + v136<<(uint(int32(3))%32) + v160 + int32(12))
	v167 = v20 + int32(4)
	v169 = v20 + int32(8)
	v175 = F_get_hash_memory_limit(m)
	mBase = m.M
	v176 = base.F64_convert_i32_u(v175)
	if base.F64_ge(v176, base.F64_mul(v164, l5)) != 0 {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v156 = v148 << (uint(int32(32)-base.I32_clz(v140)) % 32)
	goto L40
L39:
	;
	v156 = v140
	goto L40
L40:
	;
	if v140 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v160 = v156 + int32(8)
	goto L43
L42:
	;
	v160 = int32(0)
	goto L43
L43:
	;
	goto L37
L44:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v260 = base.F64_div(base.F64_mul(l5, v164), base.F64_convert_i32_u(v258))
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
	v263 = base.F64_div(l5, base.F64_convert_i64_u(v261))
	if base.F64_gt(v260, v263) != 0 {
		goto L72
	} else {
		goto L73
	}
L45:
	;
	if v20 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v189 = F_get_hash_memory_limit(m)
	mBase = m.M
	v190 = base.F64_convert_i32_u(v189)
	v196 = base.F64_mul(base.F64_add(base.F64_mul(v190, float64(0.25)), float64(-8192)), float64(0.0001220703125))
	v202 = base.F64_add(base.F64_div(base.F64_mul(v164, base.F64_mul(l5, float64(1.5))), v190), float64(1))
	if base.F64_gt(v202, v196) != 0 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(0)
	goto L50
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v175
	*(*int64)(unsafe.Add(mBase, uint32(v169))) = base.I64_trunc_sat_f64_u(base.F64_div(v176, v164))
	goto L44
L51:
	;
	v204 = v196
	goto L53
L52:
	;
	v204 = v202
	goto L53
L53:
	;
	if base.F64_lt(v204, float64(4)) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v207 = float64(4)
	goto L56
L55:
	;
	v207 = v204
	goto L56
L56:
	;
	if base.F64_gt(v207, float64(1024)) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v210 = float64(1024)
	goto L59
L58:
	;
	v210 = v207
	goto L59
L59:
	;
	v212 = F_my_log2(m, base.I32_trunc_sat_f64_s(v210))
	mBase = m.M
	if int32(31) < int32(0)+v212 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v216 = int32(32)
	goto L62
L61:
	;
	v216 = v212
	goto L62
L62:
	;
	if v20 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1) << (uint(v216) % 32)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v223 = int32(_a_F_cost_agg_0)<<(uint(v216)%32) - int32(-8192)
	if base.Ui32(v223<<(uint(int32(2))%32)) < base.Ui32(v175) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v231 = v175 - v223
	goto L68
L67:
	;
	v231 = base.I32_trunc_sat_f64_u(base.F64_mul(v176, float64(0.75)))
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v231
	v233 = base.F64_convert_i32_u(v231)
	if base.F64_gt(v233, v164) != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v238 = base.I64_trunc_sat_f64_u(base.F64_div(v233, v164))
	goto L71
L70:
	;
	v238 = int64(1)
	goto L71
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v169))) = v238
	goto L44
L72:
	;
	v265 = v260
	goto L74
L73:
	;
	v265 = v263
	goto L74
L74:
	;
	v266 = base.F64_ceil(v265)
	v267 = float64(1)
	if base.F64_gt(v266, v267) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v270 = v266
	goto L77
L76:
	;
	v270 = v267
	goto L77
L77:
	;
	v271 = F_log(m, v270)
	mBase = m.M
	v272 = int32(2)
	if v245 <= v272 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v275 = v272
	goto L80
L79:
	;
	v275 = v245
	goto L80
L80:
	;
	v277 = F_log(m, base.F64_convert_i32_u(v275))
	mBase = m.M
	v281 = base.F64_convert_i32_s(base.I32_trunc_sat_f64_s(base.F64_ceil(base.F64_div(v271, v277))))
	v282 = base.F64_mul(base.F64_mul(base.F64_mul(l10, base.F64_convert_i32_u((base.I32_trunc_sat_f64_s(l11)+int32(7))&int32(-8)+int32(24))), float64(0.0001220703125)), v281)
	v283 = base.F64_add(v282, v282)
	v285 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[3]))
	v286 = base.F64_mul(v283, v285)
	v289 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[0]))
	v290 = base.F64_mul(l10, v281)
	v292 = base.F64_mul(v289, base.F64_add(v290, v290))
	v295 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[4]))
	v302 = l5
	v303 = v123
	v304 = base.F64_add(base.F64_add(v286, v118), v292)
	v311 = base.F64_add(v292, base.F64_add(base.F64_mul(v283, v295), base.F64_add(v286, v128)))
	goto L1
L81:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v407
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v303
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v403
	m.G0 = v20 + int32(32)
	return
L82:
	;
	v314 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v314
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v319 <= int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v375 = base.F64_add(v304, v374)
	v378 = base.F64_add(v311, base.F64_add(base.F64_mul(v302, v368), v374))
	v379 = int32(0)
	v382 = F_clauselist_selectivity(m, l1, l6, v379, v379, v379)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L89
	} else {
		goto L92
	}
L84:
	;
	v322 = float64(0)
	v368 = v322
	v374 = v322
	goto L83
L85:
	;
	goto L86
L86:
	;
	v328 = int32(0)
	goto L87
L87:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v342+v328<<(uint(int32(2))%32))))
	v349 = F_cost_qual_eval_walker(m, v346, v20+int32(8))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v355 = *(*float64)(unsafe.Add(mBase, uint32(v20)+24))
	v356 = *(*float64)(unsafe.Add(mBase, uint32(v20)+16))
	v368 = v355
	v374 = v356
	goto L83
L89:
	;
	return
L90:
	;
	v352 = v328 + int32(1)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v352 < v353 {
		v328 = v352
		goto L87
	} else {
		goto L91
	}
L91:
	;
	goto L88
L92:
	;
	v384 = base.F64_mul(v302, v382)
	if base.F64_gt(v384, float64(1e+100)) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v403 = float64(1e+100)
	v406 = v375
	v407 = v378
	goto L81
L94:
	;
	goto L95
L95:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v384)&int64(9223372036854775807)) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v403 = float64(1e+100)
	v406 = v375
	v407 = v378
	goto L81
L97:
	;
	goto L98
L98:
	;
	v394 = float64(1)
	if base.F64_le(v384, v394) != 0 {
		v403 = v394
		v406 = v375
		v407 = v378
		goto L81
	} else {
		goto L99
	}
L99:
	;
	v403 = base.F64_nearest(v384)
	v406 = v375
	v407 = v378
	goto L81
}
