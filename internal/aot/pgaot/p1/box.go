package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_box_ar(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v32 float64
	_ = v32
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = base.F64_sub(v8, v9)
	v11 = base.F64_abs(v10)
	if base.F64_ne(v11, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L19
	} else {
		goto L21
	}
L2:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L19
	} else {
		goto L20
	}
L3:
	;
	v20 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v21 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = base.F64_sub(v20, v21)
	v23 = base.F64_abs(v22)
	if base.F64_ne(v23, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	if base.F64_eq(base.F64_abs(v8), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if base.F64_ne(base.F64_abs(v9), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	v32 = base.F64_mul(v10, v22)
	if base.F64_ne(base.F64_abs(v32), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if base.F64_eq(base.F64_abs(v20), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if base.F64_ne(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	if base.F64_ne(v32, float64(0)) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if base.F64_eq(v11, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if base.F64_ne(v23, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	return v32
L16:
	;
	if base.F64_eq(v10, float64(0)) != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if base.F64_ne(v22, float64(0)) != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	return float64(0)
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_box_center(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(16))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		F_box_cn(m, v5, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v5
		}
	}
}
func F_box_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 float64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_box_cn(m, v6+int32(16), v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		F_box_cn(m, v6, v8)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = F_point_dt(m, v6+int32(16), v6)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = F_Float8GetDatum(m, v20)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(32)
					return v22
				}
			}
		}
	}
}
func F_box_le(m *base.Module, l0 int32) int32 {
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
	var v9 float64
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_box_ar(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_box_ar(m, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return base.F64_le(v5, base.F64_add(v9, float64(1e-06)))
		}
	}
}
func F_box_left(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+16))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	return base.F64_gt(v3, base.F64_add(v5, float64(1e-06)))
}
func F_box_overlap(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v32 int32
	_ = v32
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.F64_le(v5, base.F64_add(v7, float64(1e-06))) == v2 {
		v32 = v2
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
		if base.F64_le(v13, base.F64_add(v14, float64(1e-06))) == int32(0) {
			v32 = v2
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			if base.F64_le(v20, base.F64_add(v21, float64(1e-06))) == int32(0) {
				v32 = v2
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
				v32 = base.F64_le(v27, base.F64_add(v28, float64(1e-06)))
			}
		}
	}
	return v32
}
