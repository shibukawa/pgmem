package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cvt_date_timestamptz(m *base.Module, l0 int64) int64 {
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
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_cvt_date_timestamptz[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v8
	v11 = *(*int64)(unsafe.Add(mBase, _c_F_cvt_date_timestamptz[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v11
	v14 = F_date2timestamptz_safe(m, base.I32_wrap_i64(l0), v5)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		m.G0 = v5 + int32(16)
		return v14
	}
}
func F_cvt_float4_float8(m *base.Module, l0 int64) int64 {
	return base.I64_reinterpret_f64(base.F64_promote_f32(base.F32_reinterpret_i32(base.I32_wrap_i64(l0))))
}
func F_cvt_int4_int2(m *base.Module, l0 int64) int64 {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = int32(-32768)
	v5 = base.I32_wrap_i64(l0)
	if v5 <= v4 {
		v8 = v4
	} else {
		v8 = v5
	}
	if int32(_a_F_cvt_int4_int2_0) <= v8 {
		v11 = int32(_a_F_cvt_int4_int2_0)
	} else {
		v11 = v8
	}
	return base.I64_extend_i32_s(v11)
}
func F_cvt_int4_int8(m *base.Module, l0 int64) int64 {
	return base.I64_extend32_s(l0)
}
func F_cvt_timestamp_timestamptz(m *base.Module, l0 int64) int64 {
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
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_cvt_timestamp_timestamptz[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v8
	v11 = *(*int64)(unsafe.Add(mBase, _c_F_cvt_timestamp_timestamptz[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v11
	v13 = F_timestamp2timestamptz_safe(m, l0, v5)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		m.G0 = v5 + int32(16)
		return v13
	}
}
func F_cvt_timestamptz_date(m *base.Module, l0 int64) int64 {
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
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_cvt_timestamptz_date[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v9
	v12 = *(*int64)(unsafe.Add(mBase, _c_F_cvt_timestamptz_date[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v12
	v14 = F_timestamptz2date_safe(m, l0, v6)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		m.G0 = v6 + int32(16)
		return base.I64_extend_i32_s(v14)
	}
}
