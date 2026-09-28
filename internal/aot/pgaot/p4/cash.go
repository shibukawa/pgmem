package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cash_dist(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14242(m, l0, int32(_a_F_cash_dist_0), int32(_a_F_cash_dist_1), int32(_a_F_cash_dist_2))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_cash_eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3))
}
func F_cash_ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(v3 <= v2))
}
func F_cash_mul_int4(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14245(m, l0, int32(_a_F_cash_mul_int4_0), int32(151), int32(_a_F_cash_mul_int4_1), int32(_a_F_cash_mul_int4_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_cash_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pq_getmsgint64(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_cash_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_pq_begintypsend(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		F_enlargeStringInfo(m, v6, int32(8))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v19 = int64(56)
			v21 = int64(65280)
			v23 = int64(40)
			v26 = int64(16711680)
			v28 = int64(24)
			v30 = int64(4278190080)
			v32 = int64(8)
			*(*int64)(unsafe.Add(mBase, uint32(v16+v17))) = v8<<(uint(v19)%64) | v8&v21<<(uint(v23)%64) | (v8&v26<<(uint(v28)%64) | v8&v30<<(uint(v32)%64)) | (int64(base.Ui64(v8)>>(uint(v32)%64))&v30 | int64(base.Ui64(v8)>>(uint(v28)%64))&v26 | (int64(base.Ui64(v8)>>(uint(v23)%64))&v21 | int64(base.Ui64(v8)>>(uint(v19)%64))))
			v56 = v16 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v56
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(v59))) = v56 << (uint(int32(2)) % 32)
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v59)
		}
	}
}
