package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_line_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v42 float64
	_ = v42
	var v43 int32
	_ = v43
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v81 float64
	_ = v81
	var v82 int32
	_ = v82
	var v89 float64
	_ = v89
	var v90 int32
	_ = v90
	var v91 float64
	_ = v91
	var v96 float64
	_ = v96
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v101 float64
	_ = v101
	var v113 float64
	_ = v113
	var v114 int32
	_ = v114
	var v115 float64
	_ = v115
	var v124 float64
	_ = v124
	var v125 int32
	_ = v125
	var v126 float64
	_ = v126
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v129 float64
	_ = v129
	var v133 float64
	_ = v133
	var v140 float64
	_ = v140
	var v141 int32
	_ = v141
	var v143 float64
	_ = v143
	var v144 float64
	_ = v144
	var v148 float64
	_ = v148
	var v149 float64
	_ = v149
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 float64
	_ = v162
	var v163 float64
	_ = v163
	var v166 int32
	_ = v166
	var v167 float64
	_ = v167
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v173 float64
	_ = v173
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v189 float64
	_ = v189
	var v197 float64
	_ = v197
	var v202 float64
	_ = v202
	var v203 float64
	_ = v203
	var v204 float64
	_ = v204
	var v213 float64
	_ = v213
	var v214 float64
	_ = v214
	var v216 float64
	_ = v216
	var v218 float64
	_ = v218
	var v225 float64
	_ = v225
	var v237 float64
	_ = v237
	var v238 int32
	_ = v238
	var v239 float64
	_ = v239
	var v241 float64
	_ = v241
	var v249 float64
	_ = v249
	var v250 int32
	_ = v250
	var v251 float64
	_ = v251
	var v261 float64
	_ = v261
	var v262 int32
	_ = v262
	var v263 float64
	_ = v263
	var v270 int64
	_ = v270
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = F_line_interpt_line(m, int32(0), v9, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v270 = int64(0)
	goto L5
L4:
	;
	v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v17 = base.F64_abs(v16)
	if base.F64_le(v17, float64(1e-06))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v17))) != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return v270
L6:
	;
	v96 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
	v98 = math.Float64frombits(uint64(0x7ff0000000000000))
	v100 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
	v101 = base.F64_mul(v91, v100)
	if base.F64_eq(base.F64_abs(v91), v98)|base.F64_ne(base.F64_abs(v101), v98)|base.F64_eq(base.F64_abs(v100), v98) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L7:
	;
	v54 = float64(1)
	v55 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
	v56 = base.F64_abs(v55)
	if base.F64_le(v56, float64(1e-06))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v56))) != 0 {
		v91 = v54
		goto L6
	} else {
		goto L16
	}
L8:
	;
	v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	v25 = base.F64_abs(v24)
	if base.F64_le(v25, float64(1e-06))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v25))) != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v32 = math.Float64frombits(uint64(0x7ff0000000000000))
	v34 = base.F64_div(v16, v24)
	if base.F64_eq(v17, v32)|base.F64_ne(base.F64_abs(v34), v32) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v42 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if base.F64_eq(v25, math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v34, float64(0)) != 0 {
		v91 = v34
		goto L6
	} else {
		goto L14
	}
L13:
	;
	v91 = v42
	goto L6
L14:
	;
	v50 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v91 = v50
	goto L6
L16:
	;
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
	v64 = base.F64_abs(v63)
	if base.F64_le(v64, float64(1e-06))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v64))) != 0 {
		v91 = v54
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v71 = math.Float64frombits(uint64(0x7ff0000000000000))
	v73 = base.F64_div(v55, v63)
	if base.F64_eq(v56, v71)|base.F64_ne(base.F64_abs(v73), v71) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v81 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if base.F64_eq(v64, math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v73, float64(0)) != 0 {
		v91 = v73
		goto L6
	} else {
		goto L22
	}
L21:
	;
	v91 = v81
	goto L6
L22:
	;
	v89 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v91 = v89
	goto L6
L24:
	;
	v127 = math.Float64frombits(uint64(0x7ff0000000000000))
	v128 = base.F64_sub(v96, v126)
	v129 = base.F64_abs(v128)
	if base.F64_ne(v129, v127) != 0 {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	v113 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v115 = float64(0)
	if base.F64_eq(v91, v115)|base.F64_ne(v101, v115)|base.F64_eq(v100, v115) != 0 {
		v126 = v101
		goto L24
	} else {
		goto L29
	}
L28:
	;
	v126 = v113
	goto L24
L29:
	;
	v124 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v126 = v124
	goto L24
L31:
	;
	v148 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v149 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
	v158 = m.G0
	v160 = v158 - int32(32)
	m.G0 = v160
	v162 = base.F64_abs(v148)
	v163 = base.F64_abs(v149)
	v166 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v162)) < base.Ui64(base.I64_reinterpret_f64(v163)))
	if base.Ui64(base.I64_reinterpret_f64(v162)) < base.Ui64(base.I64_reinterpret_f64(v163)) {
		goto L40
	} else {
		goto L41
	}
L32:
	;
	v143 = v129
	v144 = v128
	goto L31
L33:
	;
	goto L34
L34:
	;
	v133 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v96), v133)|base.F64_eq(base.F64_abs(v126), v133) != 0 {
		v143 = v127
		v144 = v128
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v140 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v143 = base.F64_abs(v140)
	v144 = v140
	goto L31
L37:
	;
	v270 = base.I64_reinterpret_f64(v263)
	goto L5
L38:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v143)))|base.F64_ne(v225, float64(0)) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L39:
	;
	m.G0 = v160 + int32(32)
	goto L38
L40:
	;
	v167 = v162
	goto L42
L41:
	;
	v167 = v163
	goto L42
L42:
	;
	v168 = base.I64_reinterpret_f64(v167)
	v170 = int64(base.Ui64(v168) >> (uint(int64(52)) % 64))
	if v170 == int64(2047) {
		v225 = v167
		goto L39
	} else {
		goto L43
	}
L43:
	;
	if base.Ui64(base.I64_reinterpret_f64(v162)) < base.Ui64(base.I64_reinterpret_f64(v163)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v173 = v163
	goto L46
L45:
	;
	v173 = v162
	goto L46
L46:
	;
	if v168 == int64(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v225 = v173
	goto L39
L48:
	;
	v176 = base.I64_reinterpret_f64(v173)
	v178 = int64(base.Ui64(v176) >> (uint(int64(52)) % 64))
	if v178 == int64(2047) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if int32(65) <= base.I32_wrap_i64(v178)-base.I32_wrap_i64(v170) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v225 = base.F64_add(v162, v163)
	goto L39
L51:
	;
	goto L52
L52:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v176) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	F_sq(m, v160+int32(24), v160+int32(16), v202)
	mBase = m.M
	F_sq(m, v160+int32(8), v160, v203)
	mBase = m.M
	v213 = *(*float64)(unsafe.Add(mBase, uint32(v160)))
	v214 = *(*float64)(unsafe.Add(mBase, uint32(v160)+16))
	v216 = *(*float64)(unsafe.Add(mBase, uint32(v160)+8))
	v218 = *(*float64)(unsafe.Add(mBase, uint32(v160)+24))
	v225 = base.F64_mul(v204, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v213, v214), v216), v218)))
	goto L39
L54:
	;
	v189 = float64(1.90109156629516e-211)
	v202 = base.F64_mul(v173, v189)
	v203 = base.F64_mul(v167, v189)
	v204 = float64(5.260135901548374e+210)
	goto L53
L55:
	;
	goto L56
L56:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v168) {
		v202 = v173
		v203 = v167
		v204 = float64(1)
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v197 = float64(5.260135901548374e+210)
	v202 = base.F64_mul(v173, v197)
	v203 = base.F64_mul(v167, v197)
	v204 = float64(1.90109156629516e-211)
	goto L53
L58:
	;
	v237 = F_float_zero_divide_error_ext(m, int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v239 = math.Float64frombits(uint64(0x7ff0000000000000))
	v241 = base.F64_div(v143, v225)
	if base.F64_eq(v143, v239)|base.F64_ne(base.F64_abs(v241), v239) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v263 = v237
	goto L37
L62:
	;
	v249 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v251 = float64(0)
	if base.F64_eq(v144, v251)|base.F64_ne(v241, v251)|base.F64_eq(base.F64_abs(v225), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v263 = v241
		goto L37
	} else {
		goto L66
	}
L65:
	;
	v263 = v249
	goto L37
L66:
	;
	v261 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v263 = v261
	goto L37
}
