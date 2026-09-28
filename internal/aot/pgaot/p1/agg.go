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
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 float64
	_ = v191
	var v197 float64
	_ = v197
	var v203 float64
	_ = v203
	var v205 float64
	_ = v205
	var v208 float64
	_ = v208
	var v211 float64
	_ = v211
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v241 float64
	_ = v241
	var v246 int64
	_ = v246
	var v253 int32
	_ = v253
	var v266 int32
	_ = v266
	var v268 float64
	_ = v268
	var v269 int64
	_ = v269
	var v271 float64
	_ = v271
	var v273 float64
	_ = v273
	var v274 float64
	_ = v274
	var v275 float64
	_ = v275
	var v278 float64
	_ = v278
	var v279 float64
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 float64
	_ = v285
	var v289 float64
	_ = v289
	var v290 float64
	_ = v290
	var v291 float64
	_ = v291
	var v293 float64
	_ = v293
	var v294 float64
	_ = v294
	var v297 float64
	_ = v297
	var v298 float64
	_ = v298
	var v300 float64
	_ = v300
	var v303 float64
	_ = v303
	var v310 float64
	_ = v310
	var v311 int32
	_ = v311
	var v312 float64
	_ = v312
	var v319 float64
	_ = v319
	var v322 int64
	_ = v322
	var v327 int32
	_ = v327
	var v330 float64
	_ = v330
	var v336 int32
	_ = v336
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 float64
	_ = v363
	var v364 float64
	_ = v364
	var v376 float64
	_ = v376
	var v382 float64
	_ = v382
	var v383 float64
	_ = v383
	var v386 float64
	_ = v386
	var v387 int32
	_ = v387
	var v390 float64
	_ = v390
	var v391 int32
	_ = v391
	var v392 float64
	_ = v392
	var v402 float64
	_ = v402
	var v411 float64
	_ = v411
	var v414 float64
	_ = v414
	var v415 float64
	_ = v415
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
		v411 = v310
		v414 = v312
		v415 = v319
		goto L84
	} else {
		goto L85
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
	v310 = float64(1)
	v311 = l7
	v312 = v41
	v319 = base.F64_add(v41, v43)
	goto L1
L8:
	;
	v123 = v122 + l7
	v126 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[0]))
	v128 = base.F64_add(base.F64_mul(v126, l5), v119)
	if l2&int32(-2) != int32(2) {
		v310 = l5
		v311 = v123
		v312 = v118
		v319 = v128
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
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+348))
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
	v164 = base.F64_convert_i32_u((base.I32_trunc_sat_f64_u(l11)+int32(23))&int32(-8) + v136<<(uint(int32(4))%32) + v160 + int32(12))
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
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v268 = base.F64_div(base.F64_mul(l5, v164), base.F64_convert_i32_u(v266))
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
	v271 = base.F64_div(l5, base.F64_convert_i64_u(v269))
	if base.F64_gt(v268, v271) != 0 {
		goto L75
	} else {
		goto L76
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
	v185 = int32(32)
	v190 = F_get_hash_memory_limit(m)
	mBase = m.M
	v191 = base.F64_convert_i32_u(v190)
	v197 = base.F64_mul(base.F64_add(base.F64_mul(v191, float64(0.25)), float64(-8192)), float64(0.0001220703125))
	v203 = base.F64_add(base.F64_div(base.F64_mul(v164, base.F64_mul(l5, float64(1.5))), v191), float64(1))
	if base.F64_gt(v203, v197) != 0 {
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
	v205 = v197
	goto L53
L52:
	;
	v205 = v203
	goto L53
L53:
	;
	if base.F64_lt(v205, float64(4)) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v208 = float64(4)
	goto L56
L55:
	;
	v208 = v205
	goto L56
L56:
	;
	if base.F64_gt(v208, float64(1024)) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v211 = float64(1024)
	goto L59
L58:
	;
	v211 = v208
	goto L59
L59:
	;
	v212 = base.I32_trunc_sat_f64_s(v211)
	if base.Ui32(int32(2)) <= base.Ui32(v212) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v220 = v185 - base.I32_clz(v212-int32(1))
	goto L62
L61:
	;
	v220 = int32(0)
	goto L62
L62:
	;
	if int32(31) < int32(0)+v220 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v224 = v185
	goto L65
L64:
	;
	v224 = v220
	goto L65
L65:
	;
	if v20 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1) << (uint(v224) % 32)
	goto L68
L67:
	;
	goto L68
L68:
	;
	v231 = int32(_a_F_cost_agg_0)<<(uint(v224)%32) - int32(-8192)
	if base.Ui32(v231<<(uint(int32(2))%32)) < base.Ui32(v175) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v239 = v175 - v231
	goto L71
L70:
	;
	v239 = base.I32_trunc_sat_f64_u(base.F64_mul(v176, float64(0.75)))
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v239
	v241 = base.F64_convert_i32_u(v239)
	if base.F64_gt(v241, v164) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v246 = base.I64_trunc_sat_f64_u(base.F64_div(v241, v164))
	goto L74
L73:
	;
	v246 = int64(1)
	goto L74
L74:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v169))) = v246
	goto L44
L75:
	;
	v273 = v268
	goto L77
L76:
	;
	v273 = v271
	goto L77
L77:
	;
	v274 = base.F64_ceil(v273)
	v275 = float64(1)
	if base.F64_gt(v274, v275) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v278 = v274
	goto L80
L79:
	;
	v278 = v275
	goto L80
L80:
	;
	v279 = F_log(m, v278)
	mBase = m.M
	v280 = int32(2)
	if v253 <= v280 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v283 = v280
	goto L83
L82:
	;
	v283 = v253
	goto L83
L83:
	;
	v285 = F_log(m, base.F64_convert_i32_u(v283))
	mBase = m.M
	v289 = base.F64_convert_i32_s(base.I32_trunc_sat_f64_s(base.F64_ceil(base.F64_div(v279, v285))))
	v290 = base.F64_mul(base.F64_mul(base.F64_mul(l10, base.F64_convert_i32_u((base.I32_trunc_sat_f64_s(l11)+int32(7))&int32(-8)+int32(24))), float64(0.0001220703125)), v289)
	v291 = base.F64_add(v290, v290)
	v293 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[3]))
	v294 = base.F64_mul(v291, v293)
	v297 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[0]))
	v298 = base.F64_mul(l10, v289)
	v300 = base.F64_mul(v297, base.F64_add(v298, v298))
	v303 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[4]))
	v310 = l5
	v311 = v123
	v312 = base.F64_add(base.F64_add(v294, v118), v300)
	v319 = base.F64_add(v300, base.F64_add(base.F64_mul(v291, v303), base.F64_add(v294, v128)))
	goto L1
L84:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v415
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v414
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v311
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v411
	m.G0 = v20 + int32(32)
	return
L85:
	;
	v322 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v322
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v327 <= int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v383 = base.F64_add(v312, v382)
	v386 = base.F64_add(v319, base.F64_add(base.F64_mul(v310, v376), v382))
	v387 = int32(0)
	v390 = F_clauselist_selectivity(m, l1, l6, v387, v387, v387)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L92
	} else {
		goto L95
	}
L87:
	;
	v330 = float64(0)
	v376 = v330
	v382 = v330
	goto L86
L88:
	;
	goto L89
L89:
	;
	v336 = int32(0)
	goto L90
L90:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v350+v336<<(uint(int32(2))%32))))
	v357 = F_cost_qual_eval_walker(m, v354, v20+int32(8))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v363 = *(*float64)(unsafe.Add(mBase, uint32(v20)+24))
	v364 = *(*float64)(unsafe.Add(mBase, uint32(v20)+16))
	v376 = v363
	v382 = v364
	goto L86
L92:
	;
	return
L93:
	;
	v360 = v336 + int32(1)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v360 < v361 {
		v336 = v360
		goto L90
	} else {
		goto L94
	}
L94:
	;
	goto L91
L95:
	;
	v392 = base.F64_mul(v310, v390)
	if base.F64_gt(v392, float64(1e+100)) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v411 = float64(1e+100)
	v414 = v383
	v415 = v386
	goto L84
L97:
	;
	goto L98
L98:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v392)&int64(9223372036854775807)) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v411 = float64(1e+100)
	v414 = v383
	v415 = v386
	goto L84
L100:
	;
	goto L101
L101:
	;
	v402 = float64(1)
	if base.F64_le(v392, v402) != 0 {
		v411 = v402
		v414 = v383
		v415 = v386
		goto L84
	} else {
		goto L102
	}
L102:
	;
	v411 = base.F64_nearest(v392)
	v414 = v383
	v415 = v386
	goto L84
}
