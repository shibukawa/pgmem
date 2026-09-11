package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_point_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v18 float64
	_ = v18
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 float64
	_ = v22
	var v25 int64
	_ = v25
	var v32 float64
	_ = v32
	var v39 int64
	_ = v39
	var v44 float64
	_ = v44
	var v61 int32
	_ = v61
	var v67 float64
	_ = v67
	var v71 int64
	_ = v71
	var v72 float64
	_ = v72
	var v75 int64
	_ = v75
	var v92 int32
	_ = v92
	v8 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
		v20 = int64(9223372036854775807)
		v21 = base.I64_reinterpret_f64(v18) & v20
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
		v25 = base.I64_reinterpret_f64(v22) & v20
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v25) {
			v61 = base.B2i32(base.Ui64(v21) < base.Ui64(int64(9218868437227405313)))
			if base.F64_ne(v12, v18) != 0 {
				v92 = v8
				return v92
			} else {
				if v61 == int32(0) {
					v92 = v8
					return v92
				} else {
					v67 = v22
					v71 = v25
					v72 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
					v75 = base.I64_reinterpret_f64(v72) & int64(9223372036854775807)
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v71) {
						return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v75))
					} else {
						return base.B2i32(base.Ui64(v75) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v72, v67)
					}
				}
			}
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v21) {
				v92 = v8
				return v92
			} else {
				v32 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
				if base.Ui64(base.I64_reinterpret_f64(v32)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
					if base.F64_ne(v12, v18) != 0 {
						if base.F64_le(base.F64_abs(base.F64_sub(v12, v18)), float64(1e-06)) == int32(0) {
							v92 = v8
						} else {
							v92 = base.F64_eq(v22, v32) | base.F64_le(base.F64_abs(base.F64_sub(v22, v32)), float64(1e-06))
						}
					} else {
						v92 = base.F64_eq(v22, v32) | base.F64_le(base.F64_abs(base.F64_sub(v22, v32)), float64(1e-06))
					}
					return v92
				} else {
					v61 = int32(1)
					if base.F64_ne(v12, v18) != 0 {
						v92 = v8
						return v92
					} else {
						if v61 == int32(0) {
							v92 = v8
							return v92
						} else {
							v67 = v22
							v71 = v25
							v72 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							v75 = base.I64_reinterpret_f64(v72) & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v71) {
								return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v75))
							} else {
								return base.B2i32(base.Ui64(v75) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v72, v67)
							}
						}
					}
				}
			}
		}
	} else {
		v39 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		if base.Ui64(v39&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			v92 = v8
			return v92
		} else {
			v44 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			v67 = v44
			v71 = base.I64_reinterpret_f64(v44) & int64(9223372036854775807)
			v72 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
			v75 = base.I64_reinterpret_f64(v72) & int64(9223372036854775807)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v71) {
				return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v75))
			} else {
				return base.B2i32(base.Ui64(v75) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v72, v67)
			}
		}
	}
}
func F_point_mul(m *base.Module, l0 int32) int32 {
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
		F_point_mul_point(m, v7, v5, v4)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_point_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 int32
	_ = v10
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(16))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_pq_getmsgfloat8(m, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v5))) = v9
			v12 = F_pq_getmsgfloat8(m, v3)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v5)+8)) = v12
				return v5
			}
		}
	}
}
