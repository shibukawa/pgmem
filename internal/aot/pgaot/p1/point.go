package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_point_div_point(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v24 float64
	_ = v24
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v38 float64
	_ = v38
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v99 float64
	_ = v99
	var v101 float64
	_ = v101
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v128 float64
	_ = v128
	var v129 float64
	_ = v129
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v158 float64
	_ = v158
	var v160 float64
	_ = v160
	var v185 int32
	_ = v185
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	v15 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v16 = base.F64_mul(v15, v15)
	v17 = base.F64_abs(v16)
	v18 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v17, v18)&base.F64_ne(base.F64_abs(v15), v18) != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_float_zero_divide_error(m)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L64
	} else {
		goto L67
	}
L2:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L64
	} else {
		goto L66
	}
L3:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L64
	} else {
		goto L65
	}
L4:
	;
	v24 = float64(0)
	if base.F64_eq(v16, v24)&base.F64_ne(v15, v24) != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v29 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v30 = base.F64_mul(v29, v29)
	v31 = base.F64_abs(v30)
	v32 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v31, v32)&base.F64_ne(base.F64_abs(v29), v32) != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v38 = float64(0)
	if base.F64_eq(v30, v38)&base.F64_ne(v29, v38) != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v43 = base.F64_add(v16, v30)
	v44 = base.F64_abs(v43)
	if base.F64_ne(v44, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v51 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v52 = base.F64_mul(v15, v51)
	v53 = base.F64_abs(v52)
	if base.F64_ne(v53, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	if base.F64_eq(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if base.F64_ne(v31, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	if base.F64_ne(v52, float64(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if base.F64_eq(base.F64_abs(v15), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if base.F64_ne(base.F64_abs(v51), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v68 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v69 = base.F64_mul(v29, v68)
	v70 = base.F64_abs(v69)
	if base.F64_ne(v70, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if base.F64_eq(v15, float64(0)) != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if base.F64_ne(v51, float64(0)) != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	if base.F64_ne(v69, float64(0)) != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	if base.F64_eq(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if base.F64_ne(base.F64_abs(v68), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v85 = base.F64_add(v52, v69)
	v86 = base.F64_abs(v85)
	if base.F64_ne(v86, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	if base.F64_eq(v29, float64(0)) != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if base.F64_ne(v68, float64(0)) != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v86)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v43, float64(0)) != 0 {
		goto L1
	} else {
		goto L32
	}
L29:
	;
	if base.F64_eq(v53, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if base.F64_ne(v70, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v99 = base.F64_div(v85, v43)
	v101 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v99), v101)&base.F64_ne(v86, v101) != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	if base.F64_ne(v99, float64(0)) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v112 = base.F64_mul(v15, v68)
	v113 = base.F64_abs(v112)
	if base.F64_ne(v113, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	if base.F64_eq(v44, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	if base.F64_ne(v85, float64(0)) != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	if base.F64_ne(v112, float64(0)) != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	if base.F64_eq(base.F64_abs(v15), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	if base.F64_ne(base.F64_abs(v68), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	goto L38
L42:
	;
	v128 = base.F64_mul(v29, v51)
	v129 = base.F64_abs(v128)
	if base.F64_ne(v129, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	if base.F64_eq(v15, float64(0)) != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	if base.F64_ne(v68, float64(0)) != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	if base.F64_ne(v128, float64(0)) != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	if base.F64_eq(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	if base.F64_ne(base.F64_abs(v51), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	v144 = base.F64_sub(v112, v128)
	v145 = base.F64_abs(v144)
	if base.F64_ne(v145, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	if base.F64_eq(v29, float64(0)) != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	if base.F64_ne(v51, float64(0)) != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v145)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v43, float64(0)) != 0 {
		goto L1
	} else {
		goto L58
	}
L55:
	;
	if base.F64_eq(v129, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	if base.F64_ne(v113, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	v158 = base.F64_div(v144, v43)
	v160 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v158), v160)&base.F64_ne(v145, v160) != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	if base.F64_ne(v158, float64(0)) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v158
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v99
	return
L61:
	;
	if base.F64_eq(v44, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if base.F64_ne(v144, float64(0)) != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	return
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_point_horiz(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
	return base.F64_eq(v5, v7) | base.F64_le(base.F64_abs(base.F64_sub(v5, v7)), float64(1e-06))
}
func F_point_right(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	return base.F64_gt(v3, base.F64_add(v5, float64(1e-06)))
}
func F_point_send(m *base.Module, l0 int32) int32 {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		F_pq_sendfloat8(m, v5, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			F_pq_sendfloat8(m, v5, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v20 << (uint(int32(2)) % 32)
				m.G0 = v5 + int32(16)
				return v19
			}
		}
	}
}
