package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cost_agg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64, l6 int32, l7 int32, l8 float64, l9 float64, l10 float64, l11 float64) {
	mBase := m.M
	_ = mBase
	var v13 float64
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 float64
	_ = v24
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v58 float64
	_ = v58
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v75 int32
	_ = v75
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v84 float64
	_ = v84
	var v93 float64
	_ = v93
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v99 float64
	_ = v99
	var v101 int32
	_ = v101
	var v102 float64
	_ = v102
	var v108 float64
	_ = v108
	var v111 float64
	_ = v111
	var v113 float64
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v155 float64
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 float64
	_ = v167
	var v173 float64
	_ = v173
	var v181 int64
	_ = v181
	var v187 int32
	_ = v187
	var v188 float64
	_ = v188
	var v194 float64
	_ = v194
	var v200 float64
	_ = v200
	var v202 float64
	_ = v202
	var v205 float64
	_ = v205
	var v208 float64
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v232 float64
	_ = v232
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 float64
	_ = v244
	var v248 float64
	_ = v248
	var v256 int64
	_ = v256
	var v273 int64
	_ = v273
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 float64
	_ = v286
	var v287 int64
	_ = v287
	var v289 float64
	_ = v289
	var v291 float64
	_ = v291
	var v292 float64
	_ = v292
	var v293 float64
	_ = v293
	var v296 float64
	_ = v296
	var v297 float64
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 float64
	_ = v303
	var v305 float64
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 float64
	_ = v313
	var v314 float64
	_ = v314
	var v315 float64
	_ = v315
	var v317 float64
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v334 float64
	_ = v334
	var v335 float64
	_ = v335
	var v337 float64
	_ = v337
	var v338 float64
	_ = v338
	var v342 float64
	_ = v342
	var v350 float64
	_ = v350
	var v351 int32
	_ = v351
	var v352 float64
	_ = v352
	var v353 float64
	_ = v353
	var v362 int64
	_ = v362
	var v367 int32
	_ = v367
	var v370 float64
	_ = v370
	var v376 int32
	_ = v376
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 float64
	_ = v403
	var v404 float64
	_ = v404
	var v416 float64
	_ = v416
	var v422 float64
	_ = v422
	var v423 float64
	_ = v423
	var v426 float64
	_ = v426
	var v427 int32
	_ = v427
	var v430 float64
	_ = v430
	var v431 int32
	_ = v431
	var v432 float64
	_ = v432
	var v442 float64
	_ = v442
	var v451 float64
	_ = v451
	var v454 float64
	_ = v454
	var v455 float64
	_ = v455
	v13 = float64(0)
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	switch l2 {
	case 0:
		goto L6
	case 1:
		v52 = l7
		goto L4
	default:
		goto L3
	case 3:
		goto L5
	}
L1:
	;
	if l6 == int32(0) {
		v451 = v350
		v454 = v352
		v455 = v353
		goto L96
	} else {
		goto L97
	}
L2:
	;
	v111 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[0]))
	v113 = base.F64_add(base.F64_mul(v111, l5), v108)
	if l2&int32(-2) != int32(2) {
		v350 = l5
		v351 = v101
		v352 = v102
		v353 = v113
		goto L1
	} else {
		goto L23
	}
L3:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_agg[1])))
	if l3 != 0 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	if l3 != 0 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_agg[1])))
	v52 = l7 + (v48 ^ int32(1))
	goto L4
L6:
	;
	if l3 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v42 = base.F64_add(v40, v41)
	v44 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[0]))
	v350 = float64(1)
	v351 = l7
	v352 = v42
	v353 = base.F64_add(v42, v44)
	goto L1
L8:
	;
	v24 = float64(0)
	v40 = base.F64_add(base.F64_add(base.F64_mul(l10, v24), base.F64_add(l9, v24)), v24)
	v41 = v24
	goto L7
L9:
	;
	goto L10
L10:
	;
	v32 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v34 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	v37 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v39 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v40 = base.F64_add(base.F64_add(base.F64_mul(v32, l10), base.F64_add(l9, v34)), v37)
	v41 = v39
	goto L7
L11:
	;
	v53 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v54 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	v55 = v54
	v56 = v53
	goto L13
L12:
	;
	v55 = v13
	v56 = v13
	goto L13
L13:
	;
	v58 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[2]))
	if l3 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v67 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v68 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v69 = v67
	v70 = v68
	goto L16
L15:
	;
	v69 = v13
	v70 = float64(0)
	goto L16
L16:
	;
	v101 = v52
	v102 = l8
	v108 = base.F64_add(base.F64_mul(v70, l5), base.F64_add(base.F64_add(base.F64_mul(base.F64_mul(v58, base.F64_convert_i32_s(l4)), l10), base.F64_add(base.F64_mul(v56, l10), base.F64_add(l9, v55))), v69))
	goto L2
L17:
	;
	v77 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v78 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	v79 = v78
	v80 = v77
	goto L19
L18:
	;
	v79 = v13
	v80 = v13
	goto L19
L19:
	;
	v84 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[2]))
	if l3 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v94 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v95 = v94
	v96 = v93
	goto L22
L21:
	;
	v95 = float64(0)
	v96 = float64(0)
	goto L22
L22:
	;
	v99 = base.F64_add(base.F64_add(base.F64_mul(base.F64_mul(v84, base.F64_convert_i32_s(l4)), l10), base.F64_add(base.F64_mul(v80, l10), base.F64_add(l9, v79))), v95)
	v101 = l7 + (v75 ^ int32(1))
	v102 = v99
	v108 = base.F64_add(base.F64_mul(v96, l5), v99)
	goto L2
L23:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+332))
	if v118 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v120 = v119
	goto L26
L25:
	;
	v120 = int32(0)
	goto L26
L26:
	;
	if base.F64_lt(l11, float64(4.294967296e+09))&base.F64_ge(l11, float64(0)) != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if l3 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v127 = base.I32_trunc_f64_u(l11)
	v129 = v127
	goto L27
L29:
	;
	goto L30
L30:
	;
	v129 = int32(0)
	goto L27
L31:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v131 = v130
	goto L33
L32:
	;
	v131 = int32(0)
	goto L33
L33:
	;
	if v131 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v155 = base.F64_convert_i32_u(v144 + ((v129+int32(23))&int32(-8) + v120<<(uint(int32(3))%32)) + int32(12))
	v158 = v20 + int32(4)
	v160 = v20 + int32(8)
	v166 = F_get_hash_memory_limit(m)
	mBase = m.M
	v167 = base.F64_convert_i32_u(v166)
	if base.F64_ge(v167, base.F64_mul(v155, l5)) != 0 {
		goto L44
	} else {
		goto L45
	}
L35:
	;
	v132 = int32(1)
	if v131&(v131-v132) != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v144 = int32(0)
	goto L37
L37:
	;
	goto L34
L38:
	;
	v140 = v132 << (uint(int32(32)-base.I32_clz(v131)) % 32)
	goto L40
L39:
	;
	v140 = v131
	goto L40
L40:
	;
	v144 = v140 + int32(8)
	goto L37
L41:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v286 = base.F64_div(base.F64_mul(l5, v155), base.F64_convert_i32_u(v284))
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
	v289 = base.F64_div(l5, base.F64_convert_i64_u(v287))
	if base.F64_gt(v286, v289) != 0 {
		goto L80
	} else {
		goto L81
	}
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v160))) = v273
	goto L41
L43:
	;
	v273 = int64(0)
	goto L42
L44:
	;
	if v20 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v187 = F_get_hash_memory_limit(m)
	mBase = m.M
	v188 = base.F64_convert_i32_u(v187)
	v194 = base.F64_mul(base.F64_add(base.F64_mul(v188, float64(0.25)), float64(-8192)), float64(0.0001220703125))
	v200 = base.F64_add(base.F64_div(base.F64_mul(v155, base.F64_mul(l5, float64(1.5))), v188), float64(1))
	if base.F64_gt(v200, v194) != 0 {
		goto L52
	} else {
		goto L53
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(0)
	goto L49
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v166
	v173 = base.F64_div(v167, v155)
	if base.F64_lt(v173, float64(1.8446744073709552e+19))&base.F64_ge(v173, float64(0)) == int32(0) {
		goto L43
	} else {
		goto L50
	}
L50:
	;
	v181 = base.I64_trunc_f64_u(v173)
	*(*int64)(unsafe.Add(mBase, uint32(v160))) = v181
	goto L41
L51:
	;
	v215 = F_my_log2(m, v214)
	mBase = m.M
	if int32(31) < int32(0)+v215 {
		goto L64
	} else {
		goto L65
	}
L52:
	;
	v202 = v194
	goto L54
L53:
	;
	v202 = v200
	goto L54
L54:
	;
	if base.F64_lt(v202, float64(4)) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v205 = float64(4)
	goto L57
L56:
	;
	v205 = v202
	goto L57
L57:
	;
	if base.F64_gt(v205, float64(1024)) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v208 = float64(1024)
	goto L60
L59:
	;
	v208 = v205
	goto L60
L60:
	;
	if base.F64_lt(base.F64_abs(v208), float64(2.147483648e+09)) != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v212 = base.I32_trunc_f64_s(v208)
	v214 = v212
	goto L51
L62:
	;
	goto L63
L63:
	;
	v214 = int32(-2147483648)
	goto L51
L64:
	;
	v219 = int32(32)
	goto L66
L65:
	;
	v219 = v215
	goto L66
L66:
	;
	if v20 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1) << (uint(v219) % 32)
	goto L69
L68:
	;
	goto L69
L69:
	;
	v226 = int32(_a_F_cost_agg_0)<<(uint(v219)%32) - int32(-8192)
	v232 = base.F64_mul(v167, float64(0.75))
	if base.F64_lt(v232, float64(4.294967296e+09))&base.F64_ge(v232, float64(0)) != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if base.Ui32(v226<<(uint(int32(2))%32)) < base.Ui32(v166) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v238 = base.I32_trunc_f64_u(v232)
	v240 = v238
	goto L70
L72:
	;
	goto L73
L73:
	;
	v240 = int32(0)
	goto L70
L74:
	;
	v241 = v166 - v226
	goto L76
L75:
	;
	v241 = v240
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v241
	v244 = base.F64_convert_i32_u(v241)
	if base.F64_lt(v155, v244) == int32(0) {
		v273 = int64(1)
		goto L42
	} else {
		goto L77
	}
L77:
	;
	v248 = base.F64_div(v244, v155)
	if base.F64_lt(v248, float64(1.8446744073709552e+19))&base.F64_ge(v248, float64(0)) == int32(0) {
		goto L43
	} else {
		goto L78
	}
L78:
	;
	v256 = base.I64_trunc_f64_u(v248)
	*(*int64)(unsafe.Add(mBase, uint32(v160))) = v256
	goto L41
L79:
	;
	v313 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[0]))
	v314 = base.F64_convert_i32_s(v311)
	v315 = base.F64_mul(l10, v314)
	v317 = base.F64_mul(v313, base.F64_add(v315, v315))
	if base.F64_lt(base.F64_abs(l11), float64(2.147483648e+09)) != 0 {
		goto L93
	} else {
		goto L94
	}
L80:
	;
	v291 = v286
	goto L82
L81:
	;
	v291 = v289
	goto L82
L82:
	;
	v292 = base.F64_ceil(v291)
	v293 = float64(1)
	if base.F64_gt(v292, v293) != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v296 = v292
	goto L85
L84:
	;
	v296 = v293
	goto L85
L85:
	;
	v297 = F_log(m, v296)
	mBase = m.M
	v298 = int32(2)
	if v282 <= v298 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v301 = v298
	goto L88
L87:
	;
	v301 = v282
	goto L88
L88:
	;
	v303 = F_log(m, base.F64_convert_i32_u(v301))
	mBase = m.M
	v305 = base.F64_ceil(base.F64_div(v297, v303))
	if base.F64_lt(base.F64_abs(v305), float64(2.147483648e+09)) != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v309 = base.I32_trunc_f64_s(v305)
	v311 = v309
	goto L79
L90:
	;
	goto L91
L91:
	;
	v311 = int32(-2147483648)
	goto L79
L92:
	;
	v334 = base.F64_mul(base.F64_mul(base.F64_mul(l10, base.F64_convert_i32_u((v323+int32(7))&int32(-8)+int32(24))), float64(0.0001220703125)), v314)
	v335 = base.F64_add(v334, v334)
	v337 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[3]))
	v338 = base.F64_mul(v335, v337)
	v342 = *(*float64)(unsafe.Add(mBase, _c_F_cost_agg[4]))
	v350 = l5
	v351 = v101
	v352 = base.F64_add(base.F64_add(v338, v102), v317)
	v353 = base.F64_add(v317, base.F64_add(base.F64_mul(v335, v342), base.F64_add(v338, v113)))
	goto L1
L93:
	;
	v321 = base.I32_trunc_f64_s(l11)
	v323 = v321
	goto L92
L94:
	;
	goto L95
L95:
	;
	v323 = int32(-2147483648)
	goto L92
L96:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v455
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v351
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v451
	m.G0 = v20 + int32(32)
	return
L97:
	;
	v362 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v362
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v362
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = l1
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v367 <= int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v423 = base.F64_add(v352, v422)
	v426 = base.F64_add(v353, base.F64_add(base.F64_mul(v350, v416), v422))
	v427 = int32(0)
	v430 = F_clauselist_selectivity(m, l1, l6, v427, v427, v427)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L104
	} else {
		goto L107
	}
L99:
	;
	v370 = float64(0)
	v416 = v370
	v422 = v370
	goto L98
L100:
	;
	goto L101
L101:
	;
	v376 = int32(0)
	goto L102
L102:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v390+v376<<(uint(int32(2))%32))))
	v397 = F_cost_qual_eval_walker(m, v394, v20+int32(8))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v403 = *(*float64)(unsafe.Add(mBase, uint32(v20)+24))
	v404 = *(*float64)(unsafe.Add(mBase, uint32(v20)+16))
	v416 = v403
	v422 = v404
	goto L98
L104:
	;
	return
L105:
	;
	v400 = v376 + int32(1)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v400 < v401 {
		v376 = v400
		goto L102
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	v432 = base.F64_mul(v350, v430)
	if base.F64_gt(v432, float64(1e+100)) != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v451 = float64(1e+100)
	v454 = v423
	v455 = v426
	goto L96
L109:
	;
	goto L110
L110:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v432)&int64(9223372036854775807)) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v451 = float64(1e+100)
	v454 = v423
	v455 = v426
	goto L96
L112:
	;
	goto L113
L113:
	;
	v442 = float64(1)
	if base.F64_le(v432, v442) != 0 {
		v451 = v442
		v454 = v423
		v455 = v426
		goto L96
	} else {
		goto L114
	}
L114:
	;
	v451 = base.F64_nearest(v432)
	v454 = v423
	v455 = v426
	goto L96
}
