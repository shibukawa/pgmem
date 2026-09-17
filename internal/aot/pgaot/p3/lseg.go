package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_lseg_distance(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_lseg_closept_lseg(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_Float8GetDatum(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_lseg_in(m *base.Module, l0 int32) int32 {
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
	var v30 int32
	_ = v30
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_palloc(m, int32(32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v24 = F_path_decode(m, v12, int32(1), int32(2), v16, v9+int32(15), int32(0), int32(_a_F_lseg_in_0), v12, v11)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			if v24 == int32(0) {
				v28 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
				v30 = int32(0)
			} else {
				v30 = v16
			}
			m.G0 = v9 + int32(16)
			return v30
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
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v54 float64
	_ = v54
	var v60 float64
	_ = v60
	var v62 float64
	_ = v62
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 float64
	_ = v89
	var v90 int32
	_ = v90
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v94 float64
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v107 float64
	_ = v107
	var v113 float64
	_ = v113
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v117 float64
	_ = v117
	var v120 int64
	_ = v120
	var v127 float64
	_ = v127
	var v134 float64
	_ = v134
	var v140 float64
	_ = v140
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 float64
	_ = v171
	var v172 float64
	_ = v172
	var v175 int64
	_ = v175
	var v176 float64
	_ = v176
	var v179 int64
	_ = v179
	var v189 float64
	_ = v189
	var v195 float64
	_ = v195
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v204 float64
	_ = v204
	var v207 int64
	_ = v207
	var v214 float64
	_ = v214
	var v226 float64
	_ = v226
	var v245 int32
	_ = v245
	var v252 float64
	_ = v252
	var v255 int64
	_ = v255
	var v256 float64
	_ = v256
	var v259 int64
	_ = v259
	var v282 int32
	_ = v282
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v296 int32
	_ = v296
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v19 = l1 + int32(16)
	v20 = F_point_sl(m, l1, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L1:
	;
	m.G0 = v16 + int32(48)
	return v296
L2:
	;
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v282)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v285
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v287
	v296 = v104
	goto L1
L3:
	;
	v195 = *(*float64)(unsafe.Add(mBase, uint32(v19)))
	if base.Ui64(base.I64_reinterpret_f64(v195)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L55
	} else {
		goto L56
	}
L4:
	;
	v176 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	v179 = base.I64_reinterpret_f64(v176) & int64(9223372036854775807)
	if base.Ui64(v175) <= base.Ui64(int64(9218868437227405312)) {
		goto L46
	} else {
		goto L47
	}
L5:
	;
	if base.B2i32(v164 == int32(0))|base.F64_ne(v113, v107) != 0 {
		v189 = v113
		goto L3
	} else {
		goto L45
	}
L6:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L44
	}
L7:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L9
	} else {
		goto L43
	}
L8:
	;
	v82 = v16 + int32(32)
	v85 = F_line_interpt_line(m, v82, v16+int32(8), l2)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L21
	}
L9:
	;
	return int32(0)
L10:
	;
	if base.F64_eq(base.F64_abs(v20), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(-4616189618054758400)
	v31 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v31
	goto L8
L12:
	;
	goto L13
L13:
	;
	if base.F64_eq(v20, float64(0)) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
	v39 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v39
	goto L8
L15:
	;
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = int64(-4616189618054758400)
	*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v20
	v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v45 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v46 = base.F64_mul(v20, v45)
	v47 = base.F64_abs(v46)
	v48 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v47, v48)&base.F64_ne(base.F64_abs(v45), v48) != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v54 = float64(0)
	if base.F64_eq(v46, v54)&base.F64_ne(v45, v54) != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v60 = math.Float64frombits(uint64(0x7ff0000000000000))
	v62 = base.F64_sub(v44, v46)
	if base.B2i32(base.F64_eq(base.F64_abs(v44), v60)|base.F64_ne(base.F64_abs(v62), v60) == int32(0))&base.F64_ne(v47, v60) != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v62
	if base.F64_ne(v62, float64(0)) != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = int64(0)
	goto L8
L21:
	;
	if v85 == int32(0) {
		v296 = v8
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v89 = F_point_dt(m, v82, l1)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v91 = F_point_dt(m, v82, v19)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	v93 = base.F64_add(v89, v91)
	v94 = F_point_dt(m, l1, v19)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	if base.F64_ne(v93, v94)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v93, v94)), float64(1e-06)) == int32(0)) != 0 {
		v296 = v8
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v104 = int32(1)
	if l0 == int32(0) {
		v296 = v104
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v107 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui64(base.I64_reinterpret_f64(v107)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v107, v113)), float64(1e-06)) == int32(0))&base.F64_ne(v113, v107) != 0 {
		v189 = v113
		goto L3
	} else {
		goto L38
	}
L29:
	;
	v113 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
	v115 = int64(9223372036854775807)
	v116 = base.I64_reinterpret_f64(v113) & v115
	v117 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v120 = base.I64_reinterpret_f64(v117) & v115
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v120) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v134 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
	if base.Ui64(base.I64_reinterpret_f64(v134)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		v189 = v134
		goto L3
	} else {
		goto L37
	}
L32:
	;
	v164 = base.B2i32(base.Ui64(v116) < base.Ui64(int64(9218868437227405313)))
	goto L5
L33:
	;
	goto L34
L34:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v116) {
		v189 = v113
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v127 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	if base.Ui64(base.I64_reinterpret_f64(v127)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v164 = int32(1)
	goto L5
L37:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v171 = v134
	v172 = v140
	v175 = base.I64_reinterpret_f64(v140) & int64(9223372036854775807)
	goto L4
L38:
	;
	if base.F64_eq(v117, v127) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v282 = l1
	goto L2
L40:
	;
	goto L41
L41:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v117, v127)), float64(1e-06)) == int32(0) {
		v189 = v113
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v282 = l1
	goto L2
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	v171 = v113
	v172 = v117
	v175 = v120
	goto L4
L46:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v179))|base.F64_ne(v172, v176) != 0 {
		v189 = v171
		goto L3
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if base.Ui64(v179) <= base.Ui64(int64(9218868437227405312)) {
		v189 = v171
		goto L3
	} else {
		goto L50
	}
L49:
	;
	v282 = l1
	goto L2
L50:
	;
	v282 = l1
	goto L2
L51:
	;
	v282 = v16 + int32(32)
	goto L2
L52:
	;
	v256 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	v259 = base.I64_reinterpret_f64(v256) & int64(9223372036854775807)
	if base.Ui64(v255) <= base.Ui64(int64(9218868437227405312)) {
		goto L68
	} else {
		goto L69
	}
L53:
	;
	if base.B2i32(v245 == int32(0))|base.F64_ne(v189, v195) != 0 {
		goto L51
	} else {
		goto L67
	}
L54:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v195, v189)), float64(1e-06)) == int32(0))&base.F64_ne(v189, v195) != 0 {
		goto L51
	} else {
		goto L64
	}
L55:
	;
	v202 = int64(9223372036854775807)
	v203 = base.I64_reinterpret_f64(v189) & v202
	v204 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v207 = base.I64_reinterpret_f64(v204) & v202
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v207) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	if base.Ui64(base.I64_reinterpret_f64(v189)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L51
	} else {
		goto L63
	}
L58:
	;
	v245 = base.B2i32(base.Ui64(v203) < base.Ui64(int64(9218868437227405313)))
	goto L53
L59:
	;
	goto L60
L60:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v203) {
		goto L51
	} else {
		goto L61
	}
L61:
	;
	v214 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	if base.Ui64(base.I64_reinterpret_f64(v214)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L54
	} else {
		goto L62
	}
L62:
	;
	v245 = int32(1)
	goto L53
L63:
	;
	v226 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v252 = v226
	v255 = base.I64_reinterpret_f64(v226) & int64(9223372036854775807)
	goto L52
L64:
	;
	if base.F64_eq(v204, v214) != 0 {
		v282 = v19
		goto L2
	} else {
		goto L65
	}
L65:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v204, v214)), float64(1e-06)) == int32(0) {
		goto L51
	} else {
		goto L66
	}
L66:
	;
	v282 = v19
	goto L2
L67:
	;
	v252 = v204
	v255 = v207
	goto L52
L68:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v259))|base.F64_ne(v256, v252) != 0 {
		goto L51
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v259) {
		v282 = v19
		goto L2
	} else {
		goto L72
	}
L71:
	;
	v282 = v19
	goto L2
L72:
	;
	goto L51
}
func F_lseg_perp(m *base.Module, l0 int32) int32 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_point_sl(m, v6, v6+int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v15 = F_point_invsl(m, v5, v5+int32(16))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return base.F64_eq(v9, v15) | base.F64_le(base.F64_abs(base.F64_sub(v9, v15)), float64(1e-06))
		}
	}
}
