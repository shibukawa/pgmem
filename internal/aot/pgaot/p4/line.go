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
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v31 float64
	_ = v31
	var v33 float64
	_ = v33
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v70 float64
	_ = v70
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v93 float64
	_ = v93
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v109 float64
	_ = v109
	var v115 int32
	_ = v115
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v120 float64
	_ = v120
	var v125 float64
	_ = v125
	var v131 float64
	_ = v131
	var v139 float64
	_ = v139
	var v142 float64
	_ = v142
	var v144 float64
	_ = v144
	var v158 float64
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
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
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L61
	}
L2:
	;
	F_float_zero_divide_error(m)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L5
	} else {
		goto L60
	}
L3:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L59
	}
L4:
	;
	v164 = F_Float8GetDatum(m, v158)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L5
	} else {
		goto L58
	}
L5:
	;
	return int32(0)
L6:
	;
	if v13 != 0 {
		v158 = float64(0)
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v17 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	v18 = base.F64_abs(v17)
	v19 = base.I64_reinterpret_f64(v18)
	if base.F64_le(v18, float64(1e-06)) != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v75 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
	v76 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
	v77 = base.F64_mul(v70, v76)
	v78 = base.F64_abs(v77)
	if base.F64_ne(v78, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L24
	} else {
		goto L25
	}
L9:
	;
	v44 = float64(1)
	v45 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
	v46 = base.F64_abs(v45)
	if base.F64_le(v46, float64(1e-06)) != 0 {
		v70 = v44
		goto L8
	} else {
		goto L17
	}
L10:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v19) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v24 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	v25 = base.F64_abs(v24)
	if base.F64_le(v25, float64(1e-06)) != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v25)) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v31 = base.F64_div(v17, v24)
	v33 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v31), v33)&base.F64_ne(v18, v33) != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	if base.F64_ne(v31, float64(0)) != 0 {
		v70 = v31
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if base.F64_eq(v25, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v70 = v31
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L1
L17:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v46)) {
		v70 = v44
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
	v53 = base.F64_abs(v52)
	if base.F64_le(v53, float64(1e-06)) != 0 {
		v70 = v44
		goto L8
	} else {
		goto L19
	}
L19:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v53)) {
		v70 = v44
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v59 = base.F64_div(v45, v52)
	v61 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v59), v61)&base.F64_ne(v46, v61) != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if base.F64_ne(v59, float64(0)) != 0 {
		v70 = v59
		goto L8
	} else {
		goto L22
	}
L22:
	;
	if base.F64_ne(v53, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v70 = v59
	goto L8
L24:
	;
	if base.F64_ne(v77, float64(0)) != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	if base.F64_eq(base.F64_abs(v70), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if base.F64_ne(base.F64_abs(v76), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v93 = math.Float64frombits(uint64(0x7ff0000000000000))
	v94 = base.F64_sub(v75, v77)
	v95 = base.F64_abs(v94)
	if base.F64_ne(v95, v93) != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	if base.F64_eq(v70, float64(0)) != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if base.F64_ne(v76, float64(0)) != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	if base.F64_eq(v18, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v139 = v93
		goto L36
	} else {
		goto L37
	}
L33:
	;
	if base.F64_eq(base.F64_abs(v75), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if base.F64_ne(v78, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v142 = base.F64_div(v95, v139)
	v144 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v142), v144)&base.F64_ne(v95, v144) != 0 {
		goto L3
	} else {
		goto L54
	}
L37:
	;
	v105 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
	v106 = base.F64_abs(v105)
	if base.F64_eq(v106, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v139 = v93
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v109 = math.Float64frombits(uint64(0x7ff8000000000000))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v19) {
		v139 = v109
		goto L36
	} else {
		goto L39
	}
L39:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v106)) {
		v139 = v109
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v115 = base.F64_lt(v18, v106)
	if v115 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v116 = v106
	goto L43
L42:
	;
	v116 = v18
	goto L43
L43:
	;
	if v115 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v117 = v18
	goto L46
L45:
	;
	v117 = v106
	goto L46
L46:
	;
	if base.F64_ne(v117, float64(0)) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v120 = base.F64_div(v117, v116)
	v125 = base.F64_mul(v116, base.F64_sqrt(base.F64_add(base.F64_mul(v120, v120), float64(1))))
	if base.F64_eq(base.F64_abs(v125), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L50
	}
L48:
	;
	v131 = v116
	goto L49
L49:
	;
	if base.F64_ne(v131, float64(0)) != 0 {
		v139 = v131
		goto L36
	} else {
		goto L52
	}
L50:
	;
	if base.F64_eq(v125, float64(0)) != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v131 = v125
	goto L49
L52:
	;
	if base.Ui64(base.I64_reinterpret_f64(v95)) <= base.Ui64(int64(9218868437227405312)) {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	v139 = v131
	goto L36
L54:
	;
	if base.F64_ne(v142, float64(0)) != 0 {
		v158 = v142
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if base.F64_eq(v94, float64(0)) != 0 {
		v158 = v142
		goto L4
	} else {
		goto L56
	}
L56:
	;
	if base.F64_ne(base.F64_abs(v139), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v158 = v142
	goto L4
L58:
	;
	return v164
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
