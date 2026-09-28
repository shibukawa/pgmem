package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cvt_date_timestamp(m *base.Module, l0 int64) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_cvt_date_timestamp[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v9
	v12 = *(*int64)(unsafe.Add(mBase, _c_F_cvt_date_timestamp[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v12
	v15 = base.I32_wrap_i64(l0)
	if v15 == int32(-2147483648) {
		v48 = int64(-9223372036854775807 - 1)
		m.G0 = v6 + int32(16)
		return v48
	} else {
		if v15 == int32(2147483647) {
			v48 = int64(9223372036854775807)
			m.G0 = v6 + int32(16)
			return v48
		} else {
			if int32(106751983) <= v15 {
				v24 = F_errsave_start(m, v6)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					if v24 == int32(0) {
						v46 = int64(9223372036854775807)
						v48 = v46
						m.G0 = v6 + int32(16)
						return v48
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_cvt_date_timestamp_0), int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, v6, int32(_a_F_cvt_date_timestamp_1), int32(646), int32(_a_F_cvt_date_timestamp_2))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int64(0)
								} else {
									v48 = int64(9223372036854775807)
									m.G0 = v6 + int32(16)
									return v48
								}
							}
						}
					}
				}
			} else {
				v46 = base.I64_extend_i32_s(v15) * int64(86400000000)
				v48 = v46
				m.G0 = v6 + int32(16)
				return v48
			}
		}
	}
}
func F_cvt_float8_float4(m *base.Module, l0 int64) int64 {
	return base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_demote_f64(base.F64_reinterpret_i64(l0))))
}
func F_cvt_int2_int4(m *base.Module, l0 int64) int64 {
	return base.I64_extend16_s(l0)
}
func F_cvt_timestamp_date(m *base.Module, l0 int64) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_cvt_timestamp_date[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v9
	v12 = *(*int64)(unsafe.Add(mBase, _c_F_cvt_timestamp_date[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v12
	v14 = F_timestamp2date_safe(m, l0, v6)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		m.G0 = v6 + int32(16)
		return base.I64_extend_i32_s(v14)
	}
}
func F_cvt_timestamptz_timestamp(m *base.Module, l0 int64) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_cvt_timestamptz_timestamp[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v8
	v11 = *(*int64)(unsafe.Add(mBase, _c_F_cvt_timestamptz_timestamp[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v11
	v13 = F_timestamptz2timestamp_safe(m, l0, v5)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		m.G0 = v5 + int32(16)
		return v13
	}
}
