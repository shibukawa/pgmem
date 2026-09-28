package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cash_div_int2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	v4 = F_cash_div_int64(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_cash_div_int4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
	v4 = F_cash_div_int64(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_cash_div_int8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = F_cash_div_int64(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_cash_mul_flt4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 float32
	_ = v3
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = F_cash_mul_float8(m, v2, base.F64_promote_f32(v3))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_cash_numeric(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v45 int64
	_ = v45
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v92 int64
	_ = v92
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_PGLC_localeconv(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+41)))
		v13 = F_int64_to_numeric(m, v7)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = base.I64_extend_i32_u(v13)
			if base.Ui32(int32(10)) < base.Ui32(v12) {
				v19 = int32(2)
			} else {
				v19 = v12
			}
			v20 = base.I32_extend8_s(v19)
			if int32(0) < v20 {
				v23 = int64(1)
				if base.Ui32(int32(8)) <= base.Ui32(v19) {
					v29 = int32(0)
					v30 = v23
					for {
						v36 = v30 * int64(100000000)
						v38 = v29 + int32(8)
						if v38 != v20&int32(120) {
							v29 = v38
							v30 = v36
							continue
						} else {
							break
						}
						break
					}
					if v19&int32(7) == int32(0) {
						v65 = v36
					} else {
						v45 = v36
						v53 = int32(0)
						v54 = v45
						for {
							v60 = v54 * int64(10)
							v62 = v53 + int32(1)
							if v62 != v20&int32(7) {
								v53 = v62
								v54 = v60
								continue
							} else {
								break
							}
							break
						}
						v65 = v60
					}
				} else {
					v45 = v23
					v53 = int32(0)
					v54 = v45
					for {
						v60 = v54 * int64(10)
						v62 = v53 + int32(1)
						if v62 != v20&int32(7) {
							v53 = v62
							v54 = v60
							continue
						} else {
							break
						}
						break
					}
					v65 = v60
				}
				v70 = int32(1390)
				v71 = int32(0)
				v76 = F_int64_to_numeric(m, v65)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int64(0)
				} else {
					v79 = base.I64_extend_i32_u(v19)
					v80 = F_DirectFunctionCall2Coll(m, v70, v71, base.I64_extend_i32_u(v76), v79)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int64(0)
					} else {
						v82 = F_DirectFunctionCall2Coll(m, int32(1391), v71, v15, v80)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int64(0)
						} else {
							v84 = F_DirectFunctionCall2Coll(m, v70, v71, v82, v79)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int64(0)
							} else {
								v92 = v84
								return v92
							}
						}
					}
				}
			} else {
				v92 = v15
				return v92
			}
		}
	}
}
