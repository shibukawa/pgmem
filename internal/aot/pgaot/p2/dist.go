package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_dist_cpoint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v14 float64
	_ = v14
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v31 float64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_point_dt(m, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v12 = base.F64_sub(v7, v11)
		v14 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v12), v14)|base.F64_eq(base.F64_abs(v7), v14)|base.F64_eq(base.F64_abs(v11), v14) == int32(0) {
			F_float_overflow_error(m)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v28 = float64(0)
			if base.F64_lt(v12, v28) != 0 {
				v31 = v28
			} else {
				v31 = v12
			}
			v32 = F_Float8GetDatum(m, v31)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				return v32
			}
		}
	}
}
func F_dist_pc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v14 float64
	_ = v14
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v31 float64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_point_dt(m, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v12 = base.F64_sub(v7, v11)
		v14 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v12), v14)|base.F64_eq(base.F64_abs(v7), v14)|base.F64_eq(base.F64_abs(v11), v14) == int32(0) {
			F_float_overflow_error(m)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v28 = float64(0)
			if base.F64_lt(v12, v28) != 0 {
				v31 = v28
			} else {
				v31 = v12
			}
			v32 = F_Float8GetDatum(m, v31)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				return v32
			}
		}
	}
}
func F_dist_polyc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_dist_ppoly_internal(m, v11, v7)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
			v15 = base.F64_sub(v12, v14)
			v17 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v15), v17)|base.F64_eq(base.F64_abs(v12), v17)|base.F64_eq(base.F64_abs(v14), v17) == int32(0) {
				F_float_overflow_error(m)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v31 = float64(0)
				if base.F64_lt(v15, v31) != 0 {
					v34 = v31
				} else {
					v34 = v15
				}
				v35 = F_Float8GetDatum(m, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v35
				}
			}
		}
	}
}
func F_dist_sb(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_box_closept_lseg(m, int32(0), v3, v4)
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
func F_dist_sl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 float64
	_ = v14
	var v15 int32
	_ = v15
	var v19 float64
	_ = v19
	var v20 int32
	_ = v20
	var v22 float64
	_ = v22
	var v25 float64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_lseg_interpt_line(m, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v25 = float64(0)
			v26 = F_Float8GetDatum(m, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				return v26
			}
		} else {
			v14 = F_line_closept_point(m, int32(0), v7, v6)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v19 = F_line_closept_point(m, int32(0), v7, v6+int32(16))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if base.F64_lt(v14, v19) != 0 {
						v22 = v14
					} else {
						v22 = v19
					}
					v25 = v22
					v26 = F_Float8GetDatum(m, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						return v26
					}
				}
			}
		}
	}
}
