package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_point_div_point(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v22 float64
	_ = v22
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v36 float64
	_ = v36
	var v45 float64
	_ = v45
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v71 float64
	_ = v71
	var v72 int32
	_ = v72
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v78 float64
	_ = v78
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v102 float64
	_ = v102
	var v103 int32
	_ = v103
	var v104 float64
	_ = v104
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v109 float64
	_ = v109
	var v122 float64
	_ = v122
	var v123 int32
	_ = v123
	var v124 float64
	_ = v124
	var v133 float64
	_ = v133
	var v134 int32
	_ = v134
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v139 float64
	_ = v139
	var v151 float64
	_ = v151
	var v152 int32
	_ = v152
	var v153 float64
	_ = v153
	var v165 float64
	_ = v165
	var v166 int32
	_ = v166
	var v168 float64
	_ = v168
	var v170 float64
	_ = v170
	var v178 float64
	_ = v178
	var v179 int32
	_ = v179
	var v183 float64
	_ = v183
	var v190 float64
	_ = v190
	var v191 int32
	_ = v191
	var v192 float64
	_ = v192
	var v193 float64
	_ = v193
	var v194 float64
	_ = v194
	var v195 float64
	_ = v195
	var v197 float64
	_ = v197
	var v210 float64
	_ = v210
	var v211 int32
	_ = v211
	var v212 float64
	_ = v212
	var v221 float64
	_ = v221
	var v222 int32
	_ = v222
	var v223 float64
	_ = v223
	var v224 float64
	_ = v224
	var v225 float64
	_ = v225
	var v226 float64
	_ = v226
	var v228 float64
	_ = v228
	var v241 float64
	_ = v241
	var v242 int32
	_ = v242
	var v243 float64
	_ = v243
	var v252 float64
	_ = v252
	var v253 int32
	_ = v253
	var v254 float64
	_ = v254
	var v256 float64
	_ = v256
	var v258 float64
	_ = v258
	var v270 float64
	_ = v270
	var v271 int32
	_ = v271
	var v272 float64
	_ = v272
	var v284 float64
	_ = v284
	var v285 int32
	_ = v285
	var v287 float64
	_ = v287
	var v289 float64
	_ = v289
	var v297 float64
	_ = v297
	var v298 int32
	_ = v298
	var v302 float64
	_ = v302
	var v309 float64
	_ = v309
	var v310 int32
	_ = v310
	var v311 float64
	_ = v311
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v11 = base.F64_mul(v10, v10)
	v13 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v11), v13)|base.F64_eq(base.F64_abs(v10), v13) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v33 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v34 = base.F64_mul(v33, v33)
	v36 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v34), v36)|base.F64_eq(base.F64_abs(v33), v36) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v22 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v24 = float64(0)
	if base.F64_eq(v10, v24)|base.F64_ne(v11, v24) != 0 {
		v32 = v11
		goto L1
	} else {
		goto L7
	}
L5:
	;
	return
L6:
	;
	v32 = v22
	goto L1
L7:
	;
	v30 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v32 = v30
	goto L1
L9:
	;
	v57 = math.Float64frombits(uint64(0x7ff0000000000000))
	v59 = base.F64_add(v32, v55)
	if base.F64_eq(base.F64_abs(v32), v57)|base.F64_ne(base.F64_abs(v59), v57)|base.F64_eq(base.F64_abs(v55), v57) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v45 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v47 = float64(0)
	if base.F64_eq(v33, v47)|base.F64_ne(v34, v47) != 0 {
		v55 = v34
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v55 = v45
	goto L9
L14:
	;
	v53 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v55 = v53
	goto L9
L16:
	;
	v71 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	v73 = v59
	goto L18
L18:
	;
	v74 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v75 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v76 = base.F64_mul(v74, v75)
	v78 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v76), v78)|base.F64_eq(base.F64_abs(v74), v78)|base.F64_eq(base.F64_abs(v75), v78) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v73 = v71
	goto L18
L20:
	;
	v105 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v106 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v107 = base.F64_mul(v105, v106)
	v109 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v107), v109)|base.F64_eq(base.F64_abs(v105), v109)|base.F64_eq(base.F64_abs(v106), v109) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v91 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v93 = float64(0)
	if base.F64_eq(v74, v93)|base.F64_ne(v76, v93)|base.F64_eq(v75, v93) != 0 {
		v104 = v76
		goto L20
	} else {
		goto L25
	}
L24:
	;
	v104 = v91
	goto L20
L25:
	;
	v102 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v104 = v102
	goto L20
L27:
	;
	v137 = math.Float64frombits(uint64(0x7ff0000000000000))
	v139 = base.F64_add(v104, v135)
	if base.F64_eq(base.F64_abs(v104), v137)|base.F64_ne(base.F64_abs(v139), v137)|base.F64_eq(base.F64_abs(v135), v137) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v122 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v124 = float64(0)
	if base.F64_eq(v105, v124)|base.F64_ne(v107, v124)|base.F64_eq(v106, v124) != 0 {
		v135 = v107
		goto L27
	} else {
		goto L32
	}
L31:
	;
	v135 = v122
	goto L27
L32:
	;
	v133 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v135 = v133
	goto L27
L34:
	;
	v151 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	v153 = v139
	goto L36
L36:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v153)&int64(9223372036854775807)))|base.F64_ne(v73, float64(0)) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v153 = v151
	goto L36
L38:
	;
	v193 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v194 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v195 = base.F64_mul(v193, v194)
	v197 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v195), v197)|base.F64_eq(base.F64_abs(v193), v197)|base.F64_eq(base.F64_abs(v194), v197) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v165 = F_float_zero_divide_error_ext(m, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v168 = math.Float64frombits(uint64(0x7ff0000000000000))
	v170 = base.F64_div(v153, v73)
	if base.F64_eq(base.F64_abs(v153), v168)|base.F64_ne(base.F64_abs(v170), v168) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v192 = v165
	goto L38
L43:
	;
	v178 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v183 = float64(0)
	if base.F64_eq(base.F64_abs(v73), math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v170, v183)|base.F64_eq(v153, v183) != 0 {
		v192 = v170
		goto L38
	} else {
		goto L47
	}
L46:
	;
	v192 = v178
	goto L38
L47:
	;
	v190 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v192 = v190
	goto L38
L49:
	;
	v224 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v225 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v226 = base.F64_mul(v224, v225)
	v228 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v226), v228)|base.F64_eq(base.F64_abs(v224), v228)|base.F64_eq(base.F64_abs(v225), v228) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	v210 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L5
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v212 = float64(0)
	if base.F64_eq(v193, v212)|base.F64_ne(v195, v212)|base.F64_eq(v194, v212) != 0 {
		v223 = v195
		goto L49
	} else {
		goto L54
	}
L53:
	;
	v223 = v210
	goto L49
L54:
	;
	v221 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v223 = v221
	goto L49
L56:
	;
	v256 = math.Float64frombits(uint64(0x7ff0000000000000))
	v258 = base.F64_sub(v223, v254)
	if base.F64_eq(base.F64_abs(v223), v256)|base.F64_ne(base.F64_abs(v258), v256)|base.F64_eq(base.F64_abs(v254), v256) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	v241 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L5
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v243 = float64(0)
	if base.F64_eq(v224, v243)|base.F64_ne(v226, v243)|base.F64_eq(v225, v243) != 0 {
		v254 = v226
		goto L56
	} else {
		goto L61
	}
L60:
	;
	v254 = v241
	goto L56
L61:
	;
	v252 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v254 = v252
	goto L56
L63:
	;
	v270 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L66
	}
L64:
	;
	v272 = v258
	goto L65
L65:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v272)&int64(9223372036854775807)))|base.F64_ne(v73, float64(0)) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v272 = v270
	goto L65
L67:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v311
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v192
	return
L68:
	;
	v284 = F_float_zero_divide_error_ext(m, int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v287 = math.Float64frombits(uint64(0x7ff0000000000000))
	v289 = base.F64_div(v272, v73)
	if base.F64_eq(base.F64_abs(v272), v287)|base.F64_ne(base.F64_abs(v289), v287) == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v311 = v284
	goto L67
L72:
	;
	v297 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L5
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v302 = float64(0)
	if base.F64_eq(base.F64_abs(v73), math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v289, v302)|base.F64_eq(v272, v302) != 0 {
		v311 = v289
		goto L67
	} else {
		goto L76
	}
L75:
	;
	v311 = v297
	goto L67
L76:
	;
	v309 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	v311 = v309
	goto L67
}
func F_point_horiz(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
	return base.I64_extend_i32_u(base.F64_eq(v5, v7) | base.F64_le(base.F64_abs(base.F64_sub(v5, v7)), float64(1e-06)))
}
func F_point_right(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	return base.I64_extend_i32_u(base.F64_gt(v3, base.F64_add(v5, float64(1e-06))))
}
func F_point_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		F_pq_sendfloat8(m, v5, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			F_pq_sendfloat8(m, v5, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v20 << (uint(int32(2)) % 32)
				m.G0 = v5 + int32(16)
				return base.I64_extend_i32_u(v19)
			}
		}
	}
}
