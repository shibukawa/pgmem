package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timestamp_cmp_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_date_cmp_timestamp_internal(m, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_s(int32(0) - v5)
	}
}
func F_timestamp_cmp_timestamptz(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamp_cmp_timestamptz[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamp_cmp_timestamptz[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
				} else {
					if v9 == int64(9223372036854775807) {
						v32 = int32(-1)
					} else {
						v32 = int32(1)
					}
					v41 = v32
				}
			} else {
				if v9 == int64(-9223372036854775807-1) {
					v37 = int32(1)
				} else {
					v37 = int32(-1)
				}
				v41 = v37
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_s(v41)
	}
}
func F_timestamp_eq_timestamptz(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamp_eq_timestamptz[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamp_eq_timestamptz[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u((v21 ^ int32(-1) | base.B2i32(base.Ui64(v17+int64(9223372036854775807)) < base.Ui64(int64(-2)))) & base.B2i32(v17 == v9))
	}
}
func F_timestamp_finite(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_extend_i32_u(base.B2i32(base.Ui64(v2+int64(9223372036854775807)) < base.Ui64(int64(-2))))
}
func F_timestamp_ge_timestamptz(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamp_ge_timestamptz[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamp_ge_timestamptz[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v33 = base.B2i32(v9 <= v17)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v33 = base.B2i32(v9 <= v17)
				} else {
					v33 = base.B2i32(v9 != int64(9223372036854775807))
				}
			} else {
				v33 = base.B2i32(v9 == int64(-9223372036854775807-1))
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v33)
	}
}
func F_timestamp_izone(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14382(m, l0, int32(_a_F_timestamp_izone_0), int32(_a_F_timestamp_izone_1), int32(_a_F_timestamp_izone_2), int64(1000000), int32(_a_F_timestamp_izone_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_timestamp_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v58 int64
	_ = v58
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v8-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		v13 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
		v58 = int64(0)
		m.G0 = v6 + int32(48)
		return v58
	} else {
		v16 = int32(0)
		v21 = F_timestamp2tm(m, v8, v16, v6+int32(4), v6, v16, v16)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			if v21 != 0 {
				v25 = int64(0)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v27 = F_errsave_start(m, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					if v27 == int32(0) {
						v58 = v25
						m.G0 = v6 + int32(48)
						return v58
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamp_time_0), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, v26, int32(_a_F_timestamp_time_1), int32(2029), int32(_a_F_timestamp_time_2))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int64(0)
								} else {
									v58 = v25
									m.G0 = v6 + int32(48)
									return v58
								}
							}
						}
					}
				}
			} else {
				v43 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6))))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v47 = int32(60)
				v58 = v43 + base.I64_extend_i32_s(v44+(v45+v46*v47)*v47)*int64(1000000)
				m.G0 = v6 + int32(48)
				return v58
			}
		}
	}
}
