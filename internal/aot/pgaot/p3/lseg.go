package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_lseg_contain_point(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v15 float64
	_ = v15
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v37 float64
	_ = v37
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v77 float64
	_ = v77
	var v90 float64
	_ = v90
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	var v97 float64
	_ = v97
	var v110 float64
	_ = v110
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v115 float64
	_ = v115
	var v117 float64
	_ = v117
	var v130 float64
	_ = v130
	var v131 int32
	_ = v131
	var v132 float64
	_ = v132
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v149 int32
	_ = v149
	var v150 float64
	_ = v150
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v156 float64
	_ = v156
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	var v172 float64
	_ = v172
	var v180 float64
	_ = v180
	var v185 float64
	_ = v185
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v196 float64
	_ = v196
	var v197 float64
	_ = v197
	var v199 float64
	_ = v199
	var v201 float64
	_ = v201
	var v208 float64
	_ = v208
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 float64
	_ = v226
	var v227 float64
	_ = v227
	var v230 int32
	_ = v230
	var v231 float64
	_ = v231
	var v232 int64
	_ = v232
	var v234 int64
	_ = v234
	var v237 float64
	_ = v237
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v253 float64
	_ = v253
	var v261 float64
	_ = v261
	var v266 float64
	_ = v266
	var v267 float64
	_ = v267
	var v268 float64
	_ = v268
	var v277 float64
	_ = v277
	var v278 float64
	_ = v278
	var v280 float64
	_ = v280
	var v282 float64
	_ = v282
	var v289 float64
	_ = v289
	var v295 float64
	_ = v295
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 float64
	_ = v308
	var v309 float64
	_ = v309
	var v312 int32
	_ = v312
	var v313 float64
	_ = v313
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v319 float64
	_ = v319
	var v322 int64
	_ = v322
	var v324 int64
	_ = v324
	var v335 float64
	_ = v335
	var v343 float64
	_ = v343
	var v348 float64
	_ = v348
	var v349 float64
	_ = v349
	var v350 float64
	_ = v350
	var v359 float64
	_ = v359
	var v360 float64
	_ = v360
	var v362 float64
	_ = v362
	var v364 float64
	_ = v364
	var v371 float64
	_ = v371
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v13 = base.F64_sub(v11, v12)
	v15 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v13), v15)|base.F64_eq(base.F64_abs(v11), v15)|base.F64_eq(base.F64_abs(v12), v15) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v32 = v13
	goto L3
L3:
	;
	v33 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v34 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = base.F64_sub(v33, v34)
	v37 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v35), v37)|base.F64_eq(base.F64_abs(v33), v37)|base.F64_eq(base.F64_abs(v34), v37) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v32 = v28
	goto L3
L6:
	;
	v50 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v52 = v35
	goto L8
L8:
	;
	v53 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v54 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v55 = base.F64_sub(v53, v54)
	v57 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v55), v57)|base.F64_eq(base.F64_abs(v53), v57)|base.F64_eq(base.F64_abs(v54), v57) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v52 = v50
	goto L8
L10:
	;
	v70 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	v72 = v55
	goto L12
L12:
	;
	v73 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v74 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v75 = base.F64_sub(v73, v74)
	v77 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v75), v77)|base.F64_eq(base.F64_abs(v73), v77)|base.F64_eq(base.F64_abs(v74), v77) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v72 = v70
	goto L12
L14:
	;
	v90 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	v92 = v75
	goto L16
L16:
	;
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v94 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v95 = base.F64_sub(v93, v94)
	v97 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v95), v97)|base.F64_eq(base.F64_abs(v93), v97)|base.F64_eq(base.F64_abs(v94), v97) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v92 = v90
	goto L16
L18:
	;
	v110 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	v112 = v95
	goto L20
L20:
	;
	v113 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v114 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v115 = base.F64_sub(v113, v114)
	v117 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v115), v117)|base.F64_eq(base.F64_abs(v113), v117)|base.F64_eq(base.F64_abs(v114), v117) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v112 = v110
	goto L20
L22:
	;
	v130 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	v132 = v115
	goto L24
L24:
	;
	v141 = m.G0
	v143 = v141 - int32(32)
	m.G0 = v143
	v145 = base.F64_abs(v32)
	v146 = base.F64_abs(v52)
	v149 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v145)) < base.Ui64(base.I64_reinterpret_f64(v146)))
	if base.Ui64(base.I64_reinterpret_f64(v145)) < base.Ui64(base.I64_reinterpret_f64(v146)) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v132 = v130
	goto L24
L26:
	;
	v222 = m.G0
	v224 = v222 - int32(32)
	m.G0 = v224
	v226 = base.F64_abs(v72)
	v227 = base.F64_abs(v92)
	v230 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v226)) < base.Ui64(base.I64_reinterpret_f64(v227)))
	if base.Ui64(base.I64_reinterpret_f64(v226)) < base.Ui64(base.I64_reinterpret_f64(v227)) {
		goto L48
	} else {
		goto L49
	}
L27:
	;
	m.G0 = v143 + int32(32)
	goto L26
L28:
	;
	v150 = v145
	goto L30
L29:
	;
	v150 = v146
	goto L30
L30:
	;
	v151 = base.I64_reinterpret_f64(v150)
	v153 = int64(base.Ui64(v151) >> (uint(int64(52)) % 64))
	if v153 == int64(2047) {
		v208 = v150
		goto L27
	} else {
		goto L31
	}
L31:
	;
	if base.Ui64(base.I64_reinterpret_f64(v145)) < base.Ui64(base.I64_reinterpret_f64(v146)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v156 = v146
	goto L34
L33:
	;
	v156 = v145
	goto L34
L34:
	;
	if v151 == int64(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v208 = v156
	goto L27
L36:
	;
	v159 = base.I64_reinterpret_f64(v156)
	v161 = int64(base.Ui64(v159) >> (uint(int64(52)) % 64))
	if v161 == int64(2047) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	if int32(65) <= base.I32_wrap_i64(v161)-base.I32_wrap_i64(v153) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v208 = base.F64_add(v145, v146)
	goto L27
L39:
	;
	goto L40
L40:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v159) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	F_sq(m, v143+int32(24), v143+int32(16), v185)
	mBase = m.M
	F_sq(m, v143+int32(8), v143, v186)
	mBase = m.M
	v196 = *(*float64)(unsafe.Add(mBase, uint32(v143)))
	v197 = *(*float64)(unsafe.Add(mBase, uint32(v143)+16))
	v199 = *(*float64)(unsafe.Add(mBase, uint32(v143)+8))
	v201 = *(*float64)(unsafe.Add(mBase, uint32(v143)+24))
	v208 = base.F64_mul(v187, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v196, v197), v199), v201)))
	goto L27
L42:
	;
	v172 = float64(1.90109156629516e-211)
	v185 = base.F64_mul(v156, v172)
	v186 = base.F64_mul(v150, v172)
	v187 = float64(5.260135901548374e+210)
	goto L41
L43:
	;
	goto L44
L44:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v151) {
		v185 = v156
		v186 = v150
		v187 = float64(1)
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v180 = float64(5.260135901548374e+210)
	v185 = base.F64_mul(v156, v180)
	v186 = base.F64_mul(v150, v180)
	v187 = float64(1.90109156629516e-211)
	goto L41
L46:
	;
	v295 = base.F64_add(v208, v289)
	v304 = m.G0
	v306 = v304 - int32(32)
	m.G0 = v306
	v308 = base.F64_abs(v112)
	v309 = base.F64_abs(v132)
	v312 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v308)) < base.Ui64(base.I64_reinterpret_f64(v309)))
	if base.Ui64(base.I64_reinterpret_f64(v308)) < base.Ui64(base.I64_reinterpret_f64(v309)) {
		goto L68
	} else {
		goto L69
	}
L47:
	;
	m.G0 = v224 + int32(32)
	goto L46
L48:
	;
	v231 = v226
	goto L50
L49:
	;
	v231 = v227
	goto L50
L50:
	;
	v232 = base.I64_reinterpret_f64(v231)
	v234 = int64(base.Ui64(v232) >> (uint(int64(52)) % 64))
	if v234 == int64(2047) {
		v289 = v231
		goto L47
	} else {
		goto L51
	}
L51:
	;
	if base.Ui64(base.I64_reinterpret_f64(v226)) < base.Ui64(base.I64_reinterpret_f64(v227)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v237 = v227
	goto L54
L53:
	;
	v237 = v226
	goto L54
L54:
	;
	if v232 == int64(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v289 = v237
	goto L47
L56:
	;
	v240 = base.I64_reinterpret_f64(v237)
	v242 = int64(base.Ui64(v240) >> (uint(int64(52)) % 64))
	if v242 == int64(2047) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	if int32(65) <= base.I32_wrap_i64(v242)-base.I32_wrap_i64(v234) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v289 = base.F64_add(v226, v227)
	goto L47
L59:
	;
	goto L60
L60:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v240) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	F_sq(m, v224+int32(24), v224+int32(16), v266)
	mBase = m.M
	F_sq(m, v224+int32(8), v224, v267)
	mBase = m.M
	v277 = *(*float64)(unsafe.Add(mBase, uint32(v224)))
	v278 = *(*float64)(unsafe.Add(mBase, uint32(v224)+16))
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v224)+8))
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v224)+24))
	v289 = base.F64_mul(v268, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v277, v278), v280), v282)))
	goto L47
L62:
	;
	v253 = float64(1.90109156629516e-211)
	v266 = base.F64_mul(v237, v253)
	v267 = base.F64_mul(v231, v253)
	v268 = float64(5.260135901548374e+210)
	goto L61
L63:
	;
	goto L64
L64:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v232) {
		v266 = v237
		v267 = v231
		v268 = float64(1)
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v261 = float64(5.260135901548374e+210)
	v266 = base.F64_mul(v237, v261)
	v267 = base.F64_mul(v231, v261)
	v268 = float64(1.90109156629516e-211)
	goto L61
L66:
	;
	return base.F64_eq(v295, v371) | base.F64_le(base.F64_abs(base.F64_sub(v295, v371)), float64(1e-06))
L67:
	;
	m.G0 = v306 + int32(32)
	goto L66
L68:
	;
	v313 = v308
	goto L70
L69:
	;
	v313 = v309
	goto L70
L70:
	;
	v314 = base.I64_reinterpret_f64(v313)
	v316 = int64(base.Ui64(v314) >> (uint(int64(52)) % 64))
	if v316 == int64(2047) {
		v371 = v313
		goto L67
	} else {
		goto L71
	}
L71:
	;
	if base.Ui64(base.I64_reinterpret_f64(v308)) < base.Ui64(base.I64_reinterpret_f64(v309)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v319 = v309
	goto L74
L73:
	;
	v319 = v308
	goto L74
L74:
	;
	if v314 == int64(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v371 = v319
	goto L67
L76:
	;
	v322 = base.I64_reinterpret_f64(v319)
	v324 = int64(base.Ui64(v322) >> (uint(int64(52)) % 64))
	if v324 == int64(2047) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	if int32(65) <= base.I32_wrap_i64(v324)-base.I32_wrap_i64(v316) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v371 = base.F64_add(v308, v309)
	goto L67
L79:
	;
	goto L80
L80:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v322) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	F_sq(m, v306+int32(24), v306+int32(16), v348)
	mBase = m.M
	F_sq(m, v306+int32(8), v306, v349)
	mBase = m.M
	v359 = *(*float64)(unsafe.Add(mBase, uint32(v306)))
	v360 = *(*float64)(unsafe.Add(mBase, uint32(v306)+16))
	v362 = *(*float64)(unsafe.Add(mBase, uint32(v306)+8))
	v364 = *(*float64)(unsafe.Add(mBase, uint32(v306)+24))
	v371 = base.F64_mul(v350, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v359, v360), v362), v364)))
	goto L67
L82:
	;
	v335 = float64(1.90109156629516e-211)
	v348 = base.F64_mul(v319, v335)
	v349 = base.F64_mul(v313, v335)
	v350 = float64(5.260135901548374e+210)
	goto L81
L83:
	;
	goto L84
L84:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v314) {
		v348 = v319
		v349 = v313
		v350 = float64(1)
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v343 = float64(5.260135901548374e+210)
	v348 = base.F64_mul(v319, v343)
	v349 = base.F64_mul(v313, v343)
	v350 = float64(1.90109156629516e-211)
	goto L81
}
func F_lseg_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = F_lseg_closept_lseg(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return base.I64_reinterpret_f64(v5)
	}
}
func F_lseg_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int64
	_ = v32
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = F_palloc(m, int32(32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v24 = F_path_decode(m, v12, int32(1), int32(2), v16, v9+int32(15), int32(0), int32(_a_F_lseg_in_0), v12, v11)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int64(0)
		} else {
			if v24 == int32(0) {
				v28 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
				v32 = int64(0)
			} else {
				v32 = base.I64_extend_i32_u(v16)
			}
			m.G0 = v9 + int32(16)
			return v32
		}
	}
}
func F_lseg_interpt_line(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 float64
	_ = v20
	var v23 int32
	_ = v23
	var v31 float64
	_ = v31
	var v39 float64
	_ = v39
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v65 float64
	_ = v65
	var v66 int32
	_ = v66
	var v67 float64
	_ = v67
	var v69 float64
	_ = v69
	var v71 float64
	_ = v71
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 float64
	_ = v109
	var v115 float64
	_ = v115
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v119 float64
	_ = v119
	var v122 int64
	_ = v122
	var v129 float64
	_ = v129
	var v136 float64
	_ = v136
	var v142 float64
	_ = v142
	var v161 int32
	_ = v161
	var v168 float64
	_ = v168
	var v169 float64
	_ = v169
	var v172 int64
	_ = v172
	var v173 float64
	_ = v173
	var v176 int64
	_ = v176
	var v186 float64
	_ = v186
	var v192 float64
	_ = v192
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v201 float64
	_ = v201
	var v204 int64
	_ = v204
	var v211 float64
	_ = v211
	var v223 float64
	_ = v223
	var v242 int32
	_ = v242
	var v249 float64
	_ = v249
	var v252 int64
	_ = v252
	var v253 float64
	_ = v253
	var v256 int64
	_ = v256
	var v279 int32
	_ = v279
	var v282 int64
	_ = v282
	var v284 int64
	_ = v284
	var v293 int32
	_ = v293
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v19 = l1 + int32(16)
	v20 = F_point_sl(m, l1, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v95 = v16 + int32(32)
	v98 = F_line_interpt_line(m, v95, v16+int32(8), l2)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L23
	}
L2:
	;
	return int32(0)
L3:
	;
	if base.F64_eq(base.F64_abs(v20), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(-4616189618054758400)
	v31 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v31
	goto L1
L5:
	;
	goto L6
L6:
	;
	if base.F64_eq(v20, float64(0)) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
	v39 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v39
	goto L1
L8:
	;
	goto L9
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = int64(-4616189618054758400)
	*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v20
	v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v45 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v46 = base.F64_mul(v20, v45)
	v48 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v46), v48)|base.F64_eq(base.F64_abs(v45), v48) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v69 = math.Float64frombits(uint64(0x7ff0000000000000))
	v71 = base.F64_sub(v44, v67)
	if base.F64_eq(base.F64_abs(v44), v69)|base.F64_ne(base.F64_abs(v71), v69)|base.F64_eq(base.F64_abs(v67), v69) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v57 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v59 = float64(0)
	if base.F64_eq(v45, v59)|base.F64_ne(v46, v59) != 0 {
		v67 = v46
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v67 = v57
	goto L10
L15:
	;
	v65 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v67 = v65
	goto L10
L17:
	;
	v83 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	v85 = v71
	goto L19
L19:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v85
	if base.F64_ne(v85, float64(0)) != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v85 = v83
	goto L19
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = int64(0)
	goto L1
L22:
	;
	m.G0 = v16 + int32(48)
	return v293
L23:
	;
	if v98 == int32(0) {
		v293 = v8
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v102 = F_lseg_contain_point(m, l1, v95)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	if v102 == int32(0) {
		v293 = v8
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v106 = int32(1)
	if l0 == int32(0) {
		v293 = v106
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v109 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui64(base.I64_reinterpret_f64(v109)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v279)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v282
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v279)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v284
	v293 = v106
	goto L22
L29:
	;
	v192 = *(*float64)(unsafe.Add(mBase, uint32(v19)))
	if base.Ui64(base.I64_reinterpret_f64(v192)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L57
	} else {
		goto L58
	}
L30:
	;
	v173 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	v176 = base.I64_reinterpret_f64(v173) & int64(9223372036854775807)
	if base.Ui64(v172) <= base.Ui64(int64(9218868437227405312)) {
		goto L48
	} else {
		goto L49
	}
L31:
	;
	if base.B2i32(v161 == int32(0))|base.F64_ne(v115, v109) != 0 {
		v186 = v115
		goto L29
	} else {
		goto L47
	}
L32:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v109, v115)), float64(1e-06)) == int32(0))&base.F64_ne(v115, v109) != 0 {
		v186 = v115
		goto L29
	} else {
		goto L42
	}
L33:
	;
	v115 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
	v117 = int64(9223372036854775807)
	v118 = base.I64_reinterpret_f64(v115) & v117
	v119 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v122 = base.I64_reinterpret_f64(v119) & v117
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v122) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v136 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
	if base.Ui64(base.I64_reinterpret_f64(v136)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		v186 = v136
		goto L29
	} else {
		goto L41
	}
L36:
	;
	v161 = base.B2i32(base.Ui64(v118) < base.Ui64(int64(9218868437227405313)))
	goto L31
L37:
	;
	goto L38
L38:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v118) {
		v186 = v115
		goto L29
	} else {
		goto L39
	}
L39:
	;
	v129 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	if base.Ui64(base.I64_reinterpret_f64(v129)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L32
	} else {
		goto L40
	}
L40:
	;
	v161 = int32(1)
	goto L31
L41:
	;
	v142 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v168 = v136
	v169 = v142
	v172 = base.I64_reinterpret_f64(v142) & int64(9223372036854775807)
	goto L30
L42:
	;
	if base.F64_eq(v119, v129) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v279 = l1
	goto L28
L44:
	;
	goto L45
L45:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v119, v129)), float64(1e-06)) == int32(0) {
		v186 = v115
		goto L29
	} else {
		goto L46
	}
L46:
	;
	v279 = l1
	goto L28
L47:
	;
	v168 = v115
	v169 = v119
	v172 = v122
	goto L30
L48:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v176))|base.F64_ne(v169, v173) != 0 {
		v186 = v168
		goto L29
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if base.Ui64(v176) <= base.Ui64(int64(9218868437227405312)) {
		v186 = v168
		goto L29
	} else {
		goto L52
	}
L51:
	;
	v279 = l1
	goto L28
L52:
	;
	v279 = l1
	goto L28
L53:
	;
	v279 = v16 + int32(32)
	goto L28
L54:
	;
	v253 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	v256 = base.I64_reinterpret_f64(v253) & int64(9223372036854775807)
	if base.Ui64(v252) <= base.Ui64(int64(9218868437227405312)) {
		goto L70
	} else {
		goto L71
	}
L55:
	;
	if base.B2i32(v242 == int32(0))|base.F64_ne(v186, v192) != 0 {
		goto L53
	} else {
		goto L69
	}
L56:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v192, v186)), float64(1e-06)) == int32(0))&base.F64_ne(v186, v192) != 0 {
		goto L53
	} else {
		goto L66
	}
L57:
	;
	v199 = int64(9223372036854775807)
	v200 = base.I64_reinterpret_f64(v186) & v199
	v201 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v204 = base.I64_reinterpret_f64(v201) & v199
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v204) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	if base.Ui64(base.I64_reinterpret_f64(v186)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L53
	} else {
		goto L65
	}
L60:
	;
	v242 = base.B2i32(base.Ui64(v200) < base.Ui64(int64(9218868437227405313)))
	goto L55
L61:
	;
	goto L62
L62:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v200) {
		goto L53
	} else {
		goto L63
	}
L63:
	;
	v211 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	if base.Ui64(base.I64_reinterpret_f64(v211)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L56
	} else {
		goto L64
	}
L64:
	;
	v242 = int32(1)
	goto L55
L65:
	;
	v223 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v249 = v223
	v252 = base.I64_reinterpret_f64(v223) & int64(9223372036854775807)
	goto L54
L66:
	;
	if base.F64_eq(v201, v211) != 0 {
		v279 = v19
		goto L28
	} else {
		goto L67
	}
L67:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v201, v211)), float64(1e-06)) == int32(0) {
		goto L53
	} else {
		goto L68
	}
L68:
	;
	v279 = v19
	goto L28
L69:
	;
	v249 = v201
	v252 = v204
	goto L54
L70:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v256))|base.F64_ne(v253, v249) != 0 {
		goto L53
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v256) {
		v279 = v19
		goto L28
	} else {
		goto L74
	}
L73:
	;
	v279 = v19
	goto L28
L74:
	;
	goto L53
}
func F_lseg_perp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 float64
	_ = v9
	var v12 int32
	_ = v12
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_point_sl(m, v6, v6+int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v15 = F_point_invsl(m, v5, v5+int32(16))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(base.F64_eq(v9, v15) | base.F64_le(base.F64_abs(base.F64_sub(v9, v15)), float64(1e-06)))
		}
	}
}
