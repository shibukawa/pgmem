package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_line_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 int64
	_ = v19
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v49 float64
	_ = v49
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v64 float64
	_ = v64
	var v66 float64
	_ = v66
	var v75 float64
	_ = v75
	var v80 float64
	_ = v80
	var v82 float64
	_ = v82
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v96 float64
	_ = v96
	var v106 float64
	_ = v106
	var v110 float64
	_ = v110
	var v111 float64
	_ = v111
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v127 int64
	_ = v127
	var v133 int32
	_ = v133
	var v134 float64
	_ = v134
	var v135 float64
	_ = v135
	var v138 float64
	_ = v138
	var v143 float64
	_ = v143
	var v149 float64
	_ = v149
	var v156 float64
	_ = v156
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v166 float64
	_ = v166
	var v177 float64
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_line_interpt_line(m, int32(0), v11, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L5
	} else {
		goto L45
	}
L2:
	;
	F_float_zero_divide_error(m)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L5
	} else {
		goto L44
	}
L3:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L43
	}
L4:
	;
	v182 = F_Float8GetDatum(m, v177)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L5
	} else {
		goto L42
	}
L5:
	;
	return int32(0)
L6:
	;
	if v13 != 0 {
		v177 = float64(0)
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v17 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	v18 = base.F64_abs(v17)
	v19 = base.I64_reinterpret_f64(v18)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v19))|base.F64_le(v18, float64(1e-06)) != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v80 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
	v82 = math.Float64frombits(uint64(0x7ff0000000000000))
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
	v85 = base.F64_mul(v75, v84)
	v86 = base.F64_abs(v85)
	if base.B2i32(base.F64_eq(base.F64_abs(v75), v82)|base.F64_ne(v86, v82) == int32(0))&base.F64_ne(base.F64_abs(v84), v82) != 0 {
		goto L3
	} else {
		goto L19
	}
L9:
	;
	v47 = float64(1)
	v48 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
	v49 = base.F64_abs(v48)
	if base.F64_le(v49, float64(1e-06))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v49))) != 0 {
		v75 = v47
		goto L8
	} else {
		goto L14
	}
L10:
	;
	v25 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	v26 = base.F64_abs(v25)
	if base.F64_le(v26, float64(1e-06))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v26))) != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v33 = base.F64_div(v17, v25)
	v35 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v33), v35)&base.F64_ne(v18, v35) != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	if base.F64_eq(v26, math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v33, float64(0)) != 0 {
		v75 = v33
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L1
L14:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
	v57 = base.F64_abs(v56)
	if base.F64_le(v57, float64(1e-06))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v57))) != 0 {
		v75 = v47
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v64 = base.F64_div(v48, v56)
	v66 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v64), v66)&base.F64_ne(v49, v66) != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	if base.F64_ne(v64, float64(0)) != 0 {
		v75 = v64
		goto L8
	} else {
		goto L17
	}
L17:
	;
	if base.F64_ne(v57, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v75 = v64
	goto L8
L19:
	;
	v96 = float64(0)
	if base.B2i32(base.F64_eq(v75, v96)|base.F64_ne(v85, v96) == int32(0))&base.F64_ne(v84, v96) != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v106 = math.Float64frombits(uint64(0x7ff0000000000000))
	v110 = base.F64_sub(v80, v85)
	v111 = base.F64_abs(v110)
	if base.B2i32(base.F64_eq(base.F64_abs(v80), v106)|base.F64_ne(v111, v106) == int32(0))&base.F64_ne(v86, v106) != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if base.F64_eq(v18, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v156 = v106
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v159 = base.F64_div(v111, v156)
	v161 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v159), v161)&base.F64_ne(v111, v161) != 0 {
		goto L3
	} else {
		goto L39
	}
L23:
	;
	v122 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
	v123 = base.F64_abs(v122)
	if base.F64_eq(v123, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v156 = v106
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v127 = int64(9218868437227405312)
	if base.B2i32(base.Ui64(v127) < base.Ui64(v19))|base.B2i32(base.Ui64(v127) < base.Ui64(base.I64_reinterpret_f64(v123))) != 0 {
		v156 = math.Float64frombits(uint64(0x7ff8000000000000))
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v133 = base.F64_gt(v123, v18)
	if v133 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v134 = v123
	goto L28
L27:
	;
	v134 = v18
	goto L28
L28:
	;
	if v133 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v135 = v18
	goto L31
L30:
	;
	v135 = v123
	goto L31
L31:
	;
	if base.F64_ne(v135, float64(0)) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v138 = base.F64_div(v135, v134)
	v143 = base.F64_mul(v134, base.F64_sqrt(base.F64_add(base.F64_mul(v138, v138), float64(1))))
	if base.F64_eq(base.F64_abs(v143), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	v149 = v134
	goto L34
L34:
	;
	if base.F64_ne(v149, float64(0)) != 0 {
		v156 = v149
		goto L22
	} else {
		goto L37
	}
L35:
	;
	if base.F64_eq(v143, float64(0)) != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v149 = v143
	goto L34
L37:
	;
	if base.Ui64(base.I64_reinterpret_f64(v111)) <= base.Ui64(int64(9218868437227405312)) {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v156 = v149
	goto L22
L39:
	;
	v166 = float64(0)
	if base.F64_eq(v110, v166)|base.F64_ne(v159, v166) != 0 {
		v177 = v159
		goto L4
	} else {
		goto L40
	}
L40:
	;
	if base.F64_ne(base.F64_abs(v156), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v177 = v159
	goto L4
L42:
	;
	return v182
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
	base.Wasm_trap_unreachable()
	for {
	}
}
