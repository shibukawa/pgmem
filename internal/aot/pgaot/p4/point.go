package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_point_add(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v19 float64
	_ = v19
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v54 int32
	_ = v54
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_palloc(m, int32(16))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v17 = base.F64_add(v15, v16)
		v19 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.B2i32(base.F64_ne(base.F64_abs(v17), v19)|base.F64_eq(base.F64_abs(v15), v19) == int32(0))&base.F64_ne(base.F64_abs(v16), v19) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v31 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v33 = base.F64_add(v31, v32)
			v35 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.B2i32(base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35) == int32(0))&base.F64_ne(base.F64_abs(v32), v35) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v33
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
				return v11
			}
		}
	}
}
func F_point_div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		F_point_div_point(m, v7, v5, v4)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_point_sl(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v53 float64
	_ = v53
	var v62 float64
	_ = v62
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v12 = math.Float64frombits(uint64(0x7ff0000000000000))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_eq(v13, v14) != 0 {
		v62 = v12
		return v62
	} else {
		v16 = base.F64_sub(v13, v14)
		v17 = base.F64_abs(v16)
		if base.F64_le(v17, float64(1e-06)) != 0 {
			v62 = v12
			return v62
		} else {
			v20 = float64(0)
			v21 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_eq(v21, v22) != 0 {
				v62 = v20
				return v62
			} else {
				v24 = base.F64_sub(v21, v22)
				v25 = base.F64_abs(v24)
				if base.F64_le(v25, float64(1e-06)) != 0 {
					v62 = v20
					return v62
				} else {
					v28 = math.Float64frombits(uint64(0x7ff0000000000000))
					v29 = base.F64_ne(v25, v28)
					v34 = int32(0)
					v41 = base.F64_ne(v17, v28)
					if base.B2i32(v29|base.F64_eq(base.F64_abs(v21), v28) == v34)&base.F64_ne(base.F64_abs(v22), v28)|base.B2i32(v41|base.F64_eq(base.F64_abs(v13), v28) == v34)&base.F64_ne(base.F64_abs(v14), v28) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v53 = base.F64_div(v24, v16)
						if v29&base.F64_eq(base.F64_abs(v53), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							if base.F64_ne(v53, float64(0)) != 0 {
								v62 = v53
								return v62
							} else {
								if v41 != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v62 = v53
									return v62
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_point_vert(m *base.Module, l0 int32) int32 {
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
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	return base.F64_eq(v5, v7) | base.F64_le(base.F64_abs(base.F64_sub(v5, v7)), float64(1e-06))
}
