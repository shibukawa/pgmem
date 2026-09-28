package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_circle_poly_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 float64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 float64
	_ = v94
	var v98 float64
	_ = v98
	var v99 int32
	_ = v99
	var v103 float64
	_ = v103
	var v104 int32
	_ = v104
	var v105 float64
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v124 float64
	_ = v124
	var v127 float64
	_ = v127
	var v134 float64
	_ = v134
	var v135 int32
	_ = v135
	var v138 float64
	_ = v138
	var v144 float64
	_ = v144
	var v145 int32
	_ = v145
	var v146 float64
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 float64
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v176 float64
	_ = v176
	var v180 int32
	_ = v180
	var v181 float64
	_ = v181
	var v182 float64
	_ = v182
	var v187 float64
	_ = v187
	var v189 float64
	_ = v189
	var v191 float64
	_ = v191
	var v194 float64
	_ = v194
	var v198 float64
	_ = v198
	var v202 float64
	_ = v202
	var v204 float64
	_ = v204
	var v212 float64
	_ = v212
	var v213 int32
	_ = v213
	var v214 float64
	_ = v214
	var v222 float64
	_ = v222
	var v223 int32
	_ = v223
	var v224 float64
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	var v238 float64
	_ = v238
	var v249 float64
	_ = v249
	var v250 int32
	_ = v250
	var v251 float64
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 float64
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v285 float64
	_ = v285
	var v289 int32
	_ = v289
	var v290 float64
	_ = v290
	var v291 float64
	_ = v291
	var v297 float64
	_ = v297
	var v298 float64
	_ = v298
	var v300 float64
	_ = v300
	var v302 float64
	_ = v302
	var v304 float64
	_ = v304
	var v310 float64
	_ = v310
	var v312 float64
	_ = v312
	var v320 float64
	_ = v320
	var v321 int32
	_ = v321
	var v322 float64
	_ = v322
	var v330 float64
	_ = v330
	var v331 int32
	_ = v331
	var v332 float64
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v343 float64
	_ = v343
	var v345 float64
	_ = v345
	var v346 float64
	_ = v346
	var v355 float64
	_ = v355
	var v356 int32
	_ = v356
	var v357 float64
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v371 float64
	_ = v371
	var v372 float64
	_ = v372
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v388 float64
	_ = v388
	var v389 float64
	_ = v389
	var v390 float64
	_ = v390
	var v391 float64
	_ = v391
	var v395 int32
	_ = v395
	var v396 float64
	_ = v396
	var v401 int32
	_ = v401
	var v405 float64
	_ = v405
	var v411 float64
	_ = v411
	var v412 float64
	_ = v412
	var v413 float64
	_ = v413
	var v418 int32
	_ = v418
	var v422 float64
	_ = v422
	var v428 float64
	_ = v428
	var v429 float64
	_ = v429
	var v430 float64
	_ = v430
	var v432 float64
	_ = v432
	var v438 float64
	_ = v438
	var v439 float64
	_ = v439
	var v441 float64
	_ = v441
	var v447 float64
	_ = v447
	var v449 int32
	_ = v449
	var v460 float64
	_ = v460
	var v461 float64
	_ = v461
	var v462 float64
	_ = v462
	var v463 float64
	_ = v463
	var v473 int32
	_ = v473
	v4 = int32(0)
	v15 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_le(base.F64_abs(v15), float64(1e-06)) != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	return v473
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v20 = F_errsave_start(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if l0 <= int32(1) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	return int32(0)
L7:
	;
	if v20 == int32(0) {
		v473 = v4
		goto L2
	} else {
		goto L8
	}
L8:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_errmsg(m, int32(_a_F_circle_poly_internal_0), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	F_errsave_finish(m, v19, int32(_a_F_circle_poly_internal_1), int32(_a_F_circle_poly_internal_2), int32(_a_F_circle_poly_internal_3))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L1
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v41 = F_errsave_start(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v58 = l0 << (uint(int32(4)) % 32)
	v59 = base.I32_div_s(v58, l0)
	if base.B2i32(v59 == int32(16))&base.B2i32(v58 <= int32(2147483607)) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	if v41 == int32(0) {
		v473 = v4
		goto L2
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	F_errmsg(m, int32(_a_F_circle_poly_internal_4), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	F_errsave_finish(m, v40, int32(_a_F_circle_poly_internal_1), int32(_a_F_circle_poly_internal_5), int32(_a_F_circle_poly_internal_3))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v68 = F_errsave_start(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v85 = v58 + int32(40)
	v86 = F_palloc0(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L28
	}
L23:
	;
	if v68 == int32(0) {
		v473 = v4
		goto L2
	} else {
		goto L24
	}
L24:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(_a_F_circle_poly_internal_6), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	F_errsave_finish(m, v67, int32(_a_F_circle_poly_internal_1), int32(_a_F_circle_poly_internal_7), int32(_a_F_circle_poly_internal_3))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	goto L1
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v85 << (uint(int32(2)) % 32)
	v94 = base.F64_div(float64(6.283185307179586), base.F64_convert_i32_u(l0))
	if base.F64_eq(v94, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v112 = v108
	v115 = v4
	goto L36
L30:
	;
	v98 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L6
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if base.F64_ne(v94, float64(0)) != 0 {
		v105 = v94
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v105 = v98
	goto L29
L34:
	;
	v103 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v105 = v103
	goto L29
L36:
	;
	v124 = math.Float64frombits(uint64(0x7ff0000000000000))
	v127 = base.F64_mul(v105, base.F64_convert_i32_u(v115))
	if base.F64_eq(base.F64_abs(v105), v124)|base.F64_ne(base.F64_abs(v127), v124) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v371 = *(*float64)(unsafe.Add(mBase, uint32(v86)+48))
	v372 = *(*float64)(unsafe.Add(mBase, uint32(v86)+40))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v373 < int32(2) {
		goto L113
	} else {
		goto L114
	}
L38:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v147 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	v134 = F_float_overflow_error_ext(m, v112)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v138 = float64(0)
	if base.B2i32(v115 == int32(0))|(base.F64_eq(v105, v138)|base.F64_ne(v127, v138)) != 0 {
		v146 = v127
		goto L38
	} else {
		goto L43
	}
L42:
	;
	v146 = v134
	goto L38
L43:
	;
	v144 = F_float_underflow_error_ext(m, v112)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	v146 = v144
	goto L38
L45:
	;
	v156 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v160 = m.G0
	v162 = v160 - int32(16)
	m.G0 = v162
	v169 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v146))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v169) <= base.Ui32(int32(1072243195)) {
		goto L52
	} else {
		goto L53
	}
L46:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if v150 != int32(453) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+4)))
	if v153 == int32(0) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L1
L49:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v225 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L50:
	;
	v202 = base.F64_mul(v156, v198)
	v204 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v202), v204)|base.F64_eq(base.F64_abs(v156), v204) == int32(0) {
		goto L61
	} else {
		goto L62
	}
L51:
	;
	m.G0 = v162 + int32(16)
	goto L50
L52:
	;
	if base.Ui32(v169) < base.Ui32(int32(1044816030)) {
		v198 = float64(1)
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v169) {
		v198 = base.F64_sub(v146, v146)
		goto L51
	} else {
		goto L56
	}
L55:
	;
	v176 = F___cos(m, v146, float64(0))
	mBase = m.M
	v198 = v176
	goto L51
L56:
	;
	v180 = F___rem_pio2(m, v146, v162)
	mBase = m.M
	v181 = *(*float64)(unsafe.Add(mBase, uint32(v162)+8))
	v182 = *(*float64)(unsafe.Add(mBase, uint32(v162)))
	switch v180&int32(3) - int32(1) {
	case 0:
		goto L59
	case 1:
		goto L58
	case 2:
		goto L57
	default:
		goto L60
	}
L57:
	;
	v194 = F___sin(m, v182, v181, int32(1))
	mBase = m.M
	v198 = v194
	goto L51
L58:
	;
	v191 = F___cos(m, v182, v181)
	mBase = m.M
	v198 = base.F64_neg(v191)
	goto L51
L59:
	;
	v189 = F___sin(m, v182, v181, int32(1))
	mBase = m.M
	v198 = base.F64_neg(v189)
	goto L51
L60:
	;
	v187 = F___cos(m, v182, v181)
	mBase = m.M
	v198 = v187
	goto L51
L61:
	;
	v212 = F_float_overflow_error_ext(m, v147)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v214 = float64(0)
	if base.F64_eq(v198, v214)|base.F64_ne(v202, v214)|base.F64_eq(v156, v214) != 0 {
		v224 = v202
		goto L49
	} else {
		goto L65
	}
L64:
	;
	v224 = v212
	goto L49
L65:
	;
	v222 = F_float_underflow_error_ext(m, v147)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	v224 = v222
	goto L49
L67:
	;
	v235 = math.Float64frombits(uint64(0x7ff0000000000000))
	v237 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v238 = base.F64_sub(v237, v224)
	if base.F64_eq(base.F64_abs(v224), v235)|base.F64_ne(base.F64_abs(v238), v235)|base.F64_eq(base.F64_abs(v237), v235) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	if v228 != int32(453) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+4)))
	if v231 == int32(0) {
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L1
L71:
	;
	v249 = F_float_overflow_error_ext(m, v225)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L6
	} else {
		goto L74
	}
L72:
	;
	v251 = v238
	goto L73
L73:
	;
	v254 = v86 + int32(40) + v115<<(uint(int32(4))%32)
	*(*float64)(unsafe.Add(mBase, uint32(v254))) = v251
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v256 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v251 = v249
	goto L73
L75:
	;
	v265 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v269 = m.G0
	v271 = v269 - int32(16)
	m.G0 = v271
	v278 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v146))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v278) <= base.Ui32(int32(1072243195)) {
		goto L82
	} else {
		goto L83
	}
L76:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if v259 != int32(453) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+4)))
	if v262 == int32(0) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L1
L79:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v333 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L80:
	;
	v310 = base.F64_mul(v265, v304)
	v312 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v310), v312)|base.F64_eq(base.F64_abs(v265), v312) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L81:
	;
	m.G0 = v271 + int32(16)
	goto L80
L82:
	;
	if base.Ui32(v278) < base.Ui32(int32(1045430272)) {
		v304 = v146
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v278) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v285 = F___sin(m, v146, float64(0), int32(0))
	mBase = m.M
	v304 = v285
	goto L81
L86:
	;
	v304 = base.F64_sub(v146, v146)
	goto L81
L87:
	;
	goto L88
L88:
	;
	v289 = F___rem_pio2(m, v146, v271)
	mBase = m.M
	v290 = *(*float64)(unsafe.Add(mBase, uint32(v271)+8))
	v291 = *(*float64)(unsafe.Add(mBase, uint32(v271)))
	switch v289&int32(3) - int32(1) {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	default:
		goto L92
	}
L89:
	;
	v302 = F___cos(m, v291, v290)
	mBase = m.M
	v304 = base.F64_neg(v302)
	goto L81
L90:
	;
	v300 = F___sin(m, v291, v290, int32(1))
	mBase = m.M
	v304 = base.F64_neg(v300)
	goto L81
L91:
	;
	v298 = F___cos(m, v291, v290)
	mBase = m.M
	v304 = v298
	goto L81
L92:
	;
	v297 = F___sin(m, v291, v290, int32(1))
	mBase = m.M
	v304 = v297
	goto L81
L93:
	;
	v320 = F_float_overflow_error_ext(m, v256)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L6
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v322 = float64(0)
	if base.F64_eq(v304, v322)|base.F64_ne(v310, v322)|base.F64_eq(v265, v322) != 0 {
		v332 = v310
		goto L79
	} else {
		goto L97
	}
L96:
	;
	v332 = v320
	goto L79
L97:
	;
	v330 = F_float_underflow_error_ext(m, v256)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	v332 = v330
	goto L79
L99:
	;
	v343 = math.Float64frombits(uint64(0x7ff0000000000000))
	v345 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v346 = base.F64_add(v332, v345)
	if base.F64_eq(base.F64_abs(v332), v343)|base.F64_ne(base.F64_abs(v346), v343)|base.F64_eq(base.F64_abs(v345), v343) != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	if v336 != int32(453) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+4)))
	if v339 == int32(0) {
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L1
L103:
	;
	v357 = v346
	goto L105
L104:
	;
	v355 = F_float_overflow_error_ext(m, v333)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L106
	}
L105:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v254)+8)) = v357
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v359 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v357 = v355
	goto L105
L107:
	;
	v369 = v115 + int32(1)
	if v369 != l0 {
		v112 = v359
		v115 = v369
		goto L36
	} else {
		goto L111
	}
L108:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	if v362 != int32(453) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+4)))
	if v365 == int32(0) {
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L1
L111:
	;
	goto L37
L112:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v86)+32)) = v462
	*(*float64)(unsafe.Add(mBase, uint32(v86)+8)) = v460
	*(*float64)(unsafe.Add(mBase, uint32(v86)+24)) = v463
	*(*float64)(unsafe.Add(mBase, uint32(v86)+16)) = v461
	v473 = v86
	goto L2
L113:
	;
	v460 = v372
	v461 = v371
	v462 = v371
	v463 = v372
	goto L112
L114:
	;
	goto L115
L115:
	;
	v381 = int32(1)
	v388 = v372
	v389 = v371
	v390 = v371
	v391 = v372
	goto L116
L116:
	;
	v395 = v86 + int32(40) + v381<<(uint(int32(4))%32)
	v396 = *(*float64)(unsafe.Add(mBase, uint32(v395)))
	v401 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v396)&int64(9223372036854775807)))
	if v401 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v460 = v438
	v461 = v447
	v462 = v429
	v463 = v412
	goto L112
L118:
	;
	if base.F64_gt(v391, v396) != 0 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v412 = v391
	goto L120
L120:
	;
	v413 = *(*float64)(unsafe.Add(mBase, uint32(v395)+8))
	v418 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v413)&int64(9223372036854775807)))
	if v418 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L121:
	;
	v405 = v396
	goto L123
L122:
	;
	v405 = v391
	goto L123
L123:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v391)&int64(9223372036854775807)) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v411 = v396
	goto L126
L125:
	;
	v411 = v405
	goto L126
L126:
	;
	v412 = v411
	goto L120
L127:
	;
	if base.F64_lt(v413, v390) != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v429 = v390
	goto L129
L129:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v396)&int64(9223372036854775807)) {
		goto L136
	} else {
		goto L137
	}
L130:
	;
	v422 = v413
	goto L132
L131:
	;
	v422 = v390
	goto L132
L132:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v390)&int64(9223372036854775807)) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v428 = v413
	goto L135
L134:
	;
	v428 = v422
	goto L135
L135:
	;
	v429 = v428
	goto L129
L136:
	;
	v430 = v396
	goto L138
L137:
	;
	v430 = v388
	goto L138
L138:
	;
	if base.F64_lt(v388, v396) != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v432 = v396
	goto L141
L140:
	;
	v432 = v430
	goto L141
L141:
	;
	if base.Ui64(base.I64_reinterpret_f64(v388)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v438 = v432
	goto L144
L143:
	;
	v438 = v388
	goto L144
L144:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v413)&int64(9223372036854775807)) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v439 = v413
	goto L147
L146:
	;
	v439 = v389
	goto L147
L147:
	;
	if base.F64_gt(v413, v389) != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v441 = v413
	goto L150
L149:
	;
	v441 = v439
	goto L150
L150:
	;
	if base.Ui64(base.I64_reinterpret_f64(v389)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v447 = v441
	goto L153
L152:
	;
	v447 = v389
	goto L153
L153:
	;
	v449 = v381 + int32(1)
	if v449 != v373 {
		v381 = v449
		v388 = v438
		v389 = v447
		v390 = v429
		v391 = v412
		goto L116
	} else {
		goto L154
	}
L154:
	;
	goto L117
}
func F_circle_radius(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)+16))
	return v3
}
