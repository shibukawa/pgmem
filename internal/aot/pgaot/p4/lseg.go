package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_lseg_ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v109 int32
	_ = v109
	var v110 float64
	_ = v110
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v116 float64
	_ = v116
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v132 float64
	_ = v132
	var v140 float64
	_ = v140
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v147 float64
	_ = v147
	var v156 float64
	_ = v156
	var v157 float64
	_ = v157
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v168 float64
	_ = v168
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 float64
	_ = v188
	var v189 float64
	_ = v189
	var v192 int32
	_ = v192
	var v193 float64
	_ = v193
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v199 float64
	_ = v199
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v215 float64
	_ = v215
	var v223 float64
	_ = v223
	var v228 float64
	_ = v228
	var v229 float64
	_ = v229
	var v230 float64
	_ = v230
	var v239 float64
	_ = v239
	var v240 float64
	_ = v240
	var v242 float64
	_ = v242
	var v244 float64
	_ = v244
	var v251 float64
	_ = v251
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
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
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
	v34 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
	v35 = base.F64_sub(v33, v34)
	v37 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v35), v37)|base.F64_eq(base.F64_abs(v33), v37)|base.F64_eq(base.F64_abs(v34), v37) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int64(0)
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
	v53 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
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
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
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
	v101 = m.G0
	v103 = v101 - int32(32)
	m.G0 = v103
	v105 = base.F64_abs(v32)
	v106 = base.F64_abs(v52)
	v109 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v105)) < base.Ui64(base.I64_reinterpret_f64(v106)))
	if base.Ui64(base.I64_reinterpret_f64(v105)) < base.Ui64(base.I64_reinterpret_f64(v106)) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v92 = v90
	goto L16
L18:
	;
	v184 = m.G0
	v186 = v184 - int32(32)
	m.G0 = v186
	v188 = base.F64_abs(v72)
	v189 = base.F64_abs(v92)
	v192 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v188)) < base.Ui64(base.I64_reinterpret_f64(v189)))
	if base.Ui64(base.I64_reinterpret_f64(v188)) < base.Ui64(base.I64_reinterpret_f64(v189)) {
		goto L40
	} else {
		goto L41
	}
L19:
	;
	m.G0 = v103 + int32(32)
	goto L18
L20:
	;
	v110 = v105
	goto L22
L21:
	;
	v110 = v106
	goto L22
L22:
	;
	v111 = base.I64_reinterpret_f64(v110)
	v113 = int64(base.Ui64(v111) >> (uint(int64(52)) % 64))
	if v113 == int64(2047) {
		v168 = v110
		goto L19
	} else {
		goto L23
	}
L23:
	;
	if base.Ui64(base.I64_reinterpret_f64(v105)) < base.Ui64(base.I64_reinterpret_f64(v106)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v116 = v106
	goto L26
L25:
	;
	v116 = v105
	goto L26
L26:
	;
	if v111 == int64(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v168 = v116
	goto L19
L28:
	;
	v119 = base.I64_reinterpret_f64(v116)
	v121 = int64(base.Ui64(v119) >> (uint(int64(52)) % 64))
	if v121 == int64(2047) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if int32(65) <= base.I32_wrap_i64(v121)-base.I32_wrap_i64(v113) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v168 = base.F64_add(v105, v106)
	goto L19
L31:
	;
	goto L32
L32:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v119) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	F_sq(m, v103+int32(24), v103+int32(16), v145)
	mBase = m.M
	F_sq(m, v103+int32(8), v103, v146)
	mBase = m.M
	v156 = *(*float64)(unsafe.Add(mBase, uint32(v103)))
	v157 = *(*float64)(unsafe.Add(mBase, uint32(v103)+16))
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v103)+8))
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v103)+24))
	v168 = base.F64_mul(v147, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v156, v157), v159), v161)))
	goto L19
L34:
	;
	v132 = float64(1.90109156629516e-211)
	v145 = base.F64_mul(v116, v132)
	v146 = base.F64_mul(v110, v132)
	v147 = float64(5.260135901548374e+210)
	goto L33
L35:
	;
	goto L36
L36:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v111) {
		v145 = v116
		v146 = v110
		v147 = float64(1)
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v140 = float64(5.260135901548374e+210)
	v145 = base.F64_mul(v116, v140)
	v146 = base.F64_mul(v110, v140)
	v147 = float64(1.90109156629516e-211)
	goto L33
L38:
	;
	return base.I64_extend_i32_u(base.F64_ge(base.F64_add(v168, float64(1e-06)), v251))
L39:
	;
	m.G0 = v186 + int32(32)
	goto L38
L40:
	;
	v193 = v188
	goto L42
L41:
	;
	v193 = v189
	goto L42
L42:
	;
	v194 = base.I64_reinterpret_f64(v193)
	v196 = int64(base.Ui64(v194) >> (uint(int64(52)) % 64))
	if v196 == int64(2047) {
		v251 = v193
		goto L39
	} else {
		goto L43
	}
L43:
	;
	if base.Ui64(base.I64_reinterpret_f64(v188)) < base.Ui64(base.I64_reinterpret_f64(v189)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v199 = v189
	goto L46
L45:
	;
	v199 = v188
	goto L46
L46:
	;
	if v194 == int64(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v251 = v199
	goto L39
L48:
	;
	v202 = base.I64_reinterpret_f64(v199)
	v204 = int64(base.Ui64(v202) >> (uint(int64(52)) % 64))
	if v204 == int64(2047) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if int32(65) <= base.I32_wrap_i64(v204)-base.I32_wrap_i64(v196) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v251 = base.F64_add(v188, v189)
	goto L39
L51:
	;
	goto L52
L52:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v202) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	F_sq(m, v186+int32(24), v186+int32(16), v228)
	mBase = m.M
	F_sq(m, v186+int32(8), v186, v229)
	mBase = m.M
	v239 = *(*float64)(unsafe.Add(mBase, uint32(v186)))
	v240 = *(*float64)(unsafe.Add(mBase, uint32(v186)+16))
	v242 = *(*float64)(unsafe.Add(mBase, uint32(v186)+8))
	v244 = *(*float64)(unsafe.Add(mBase, uint32(v186)+24))
	v251 = base.F64_mul(v230, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v239, v240), v242), v244)))
	goto L39
L54:
	;
	v215 = float64(1.90109156629516e-211)
	v228 = base.F64_mul(v199, v215)
	v229 = base.F64_mul(v193, v215)
	v230 = float64(5.260135901548374e+210)
	goto L53
L55:
	;
	goto L56
L56:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v194) {
		v228 = v199
		v229 = v193
		v230 = float64(1)
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v223 = float64(5.260135901548374e+210)
	v228 = base.F64_mul(v199, v223)
	v229 = base.F64_mul(v193, v223)
	v230 = float64(1.90109156629516e-211)
	goto L53
}
func F_lseg_lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v109 int32
	_ = v109
	var v110 float64
	_ = v110
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v116 float64
	_ = v116
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v132 float64
	_ = v132
	var v140 float64
	_ = v140
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v147 float64
	_ = v147
	var v156 float64
	_ = v156
	var v157 float64
	_ = v157
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v168 float64
	_ = v168
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 float64
	_ = v188
	var v189 float64
	_ = v189
	var v192 int32
	_ = v192
	var v193 float64
	_ = v193
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v199 float64
	_ = v199
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v215 float64
	_ = v215
	var v223 float64
	_ = v223
	var v228 float64
	_ = v228
	var v229 float64
	_ = v229
	var v230 float64
	_ = v230
	var v239 float64
	_ = v239
	var v240 float64
	_ = v240
	var v242 float64
	_ = v242
	var v244 float64
	_ = v244
	var v251 float64
	_ = v251
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
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
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
	v34 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
	v35 = base.F64_sub(v33, v34)
	v37 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v35), v37)|base.F64_eq(base.F64_abs(v33), v37)|base.F64_eq(base.F64_abs(v34), v37) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int64(0)
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
	v53 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
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
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
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
	v101 = m.G0
	v103 = v101 - int32(32)
	m.G0 = v103
	v105 = base.F64_abs(v32)
	v106 = base.F64_abs(v52)
	v109 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v105)) < base.Ui64(base.I64_reinterpret_f64(v106)))
	if base.Ui64(base.I64_reinterpret_f64(v105)) < base.Ui64(base.I64_reinterpret_f64(v106)) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v92 = v90
	goto L16
L18:
	;
	v184 = m.G0
	v186 = v184 - int32(32)
	m.G0 = v186
	v188 = base.F64_abs(v72)
	v189 = base.F64_abs(v92)
	v192 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v188)) < base.Ui64(base.I64_reinterpret_f64(v189)))
	if base.Ui64(base.I64_reinterpret_f64(v188)) < base.Ui64(base.I64_reinterpret_f64(v189)) {
		goto L40
	} else {
		goto L41
	}
L19:
	;
	m.G0 = v103 + int32(32)
	goto L18
L20:
	;
	v110 = v105
	goto L22
L21:
	;
	v110 = v106
	goto L22
L22:
	;
	v111 = base.I64_reinterpret_f64(v110)
	v113 = int64(base.Ui64(v111) >> (uint(int64(52)) % 64))
	if v113 == int64(2047) {
		v168 = v110
		goto L19
	} else {
		goto L23
	}
L23:
	;
	if base.Ui64(base.I64_reinterpret_f64(v105)) < base.Ui64(base.I64_reinterpret_f64(v106)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v116 = v106
	goto L26
L25:
	;
	v116 = v105
	goto L26
L26:
	;
	if v111 == int64(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v168 = v116
	goto L19
L28:
	;
	v119 = base.I64_reinterpret_f64(v116)
	v121 = int64(base.Ui64(v119) >> (uint(int64(52)) % 64))
	if v121 == int64(2047) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if int32(65) <= base.I32_wrap_i64(v121)-base.I32_wrap_i64(v113) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v168 = base.F64_add(v105, v106)
	goto L19
L31:
	;
	goto L32
L32:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v119) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	F_sq(m, v103+int32(24), v103+int32(16), v145)
	mBase = m.M
	F_sq(m, v103+int32(8), v103, v146)
	mBase = m.M
	v156 = *(*float64)(unsafe.Add(mBase, uint32(v103)))
	v157 = *(*float64)(unsafe.Add(mBase, uint32(v103)+16))
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v103)+8))
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v103)+24))
	v168 = base.F64_mul(v147, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v156, v157), v159), v161)))
	goto L19
L34:
	;
	v132 = float64(1.90109156629516e-211)
	v145 = base.F64_mul(v116, v132)
	v146 = base.F64_mul(v110, v132)
	v147 = float64(5.260135901548374e+210)
	goto L33
L35:
	;
	goto L36
L36:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v111) {
		v145 = v116
		v146 = v110
		v147 = float64(1)
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v140 = float64(5.260135901548374e+210)
	v145 = base.F64_mul(v116, v140)
	v146 = base.F64_mul(v110, v140)
	v147 = float64(1.90109156629516e-211)
	goto L33
L38:
	;
	return base.I64_extend_i32_u(base.F64_lt(base.F64_add(v168, float64(1e-06)), v251))
L39:
	;
	m.G0 = v186 + int32(32)
	goto L38
L40:
	;
	v193 = v188
	goto L42
L41:
	;
	v193 = v189
	goto L42
L42:
	;
	v194 = base.I64_reinterpret_f64(v193)
	v196 = int64(base.Ui64(v194) >> (uint(int64(52)) % 64))
	if v196 == int64(2047) {
		v251 = v193
		goto L39
	} else {
		goto L43
	}
L43:
	;
	if base.Ui64(base.I64_reinterpret_f64(v188)) < base.Ui64(base.I64_reinterpret_f64(v189)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v199 = v189
	goto L46
L45:
	;
	v199 = v188
	goto L46
L46:
	;
	if v194 == int64(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v251 = v199
	goto L39
L48:
	;
	v202 = base.I64_reinterpret_f64(v199)
	v204 = int64(base.Ui64(v202) >> (uint(int64(52)) % 64))
	if v204 == int64(2047) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if int32(65) <= base.I32_wrap_i64(v204)-base.I32_wrap_i64(v196) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v251 = base.F64_add(v188, v189)
	goto L39
L51:
	;
	goto L52
L52:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v202) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	F_sq(m, v186+int32(24), v186+int32(16), v228)
	mBase = m.M
	F_sq(m, v186+int32(8), v186, v229)
	mBase = m.M
	v239 = *(*float64)(unsafe.Add(mBase, uint32(v186)))
	v240 = *(*float64)(unsafe.Add(mBase, uint32(v186)+16))
	v242 = *(*float64)(unsafe.Add(mBase, uint32(v186)+8))
	v244 = *(*float64)(unsafe.Add(mBase, uint32(v186)+24))
	v251 = base.F64_mul(v230, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v239, v240), v242), v244)))
	goto L39
L54:
	;
	v215 = float64(1.90109156629516e-211)
	v228 = base.F64_mul(v199, v215)
	v229 = base.F64_mul(v193, v215)
	v230 = float64(5.260135901548374e+210)
	goto L53
L55:
	;
	goto L56
L56:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v194) {
		v228 = v199
		v229 = v193
		v230 = float64(1)
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v223 = float64(5.260135901548374e+210)
	v228 = base.F64_mul(v199, v223)
	v229 = base.F64_mul(v193, v223)
	v230 = float64(1.90109156629516e-211)
	goto L53
}
func F_lseg_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_path_encode(m, int32(1), int32(2), v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v5)
	}
}
