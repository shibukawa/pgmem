package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cash_dist(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13880(m, l0, int32(_a_F_cash_dist_0), int32(_a_F_cash_dist_1), int32(_a_F_cash_dist_2))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_cash_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	return base.B2i32(v3 == v5)
}
func F_cash_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	return base.B2i32(v5 <= v3)
}
func F_cash_mul_int4(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13883(m, l0, int32(_a_F_cash_mul_int4_0), int32(150), int32(_a_F_cash_mul_int4_1), int32(_a_F_cash_mul_int4_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_cash_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pq_getmsgint64(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_Int64GetDatum(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_cash_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	F_pq_begintypsend(m, v6)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		F_enlargeStringInfo(m, v6, int32(8))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v20 = int64(56)
			v22 = int64(65280)
			v24 = int64(40)
			v27 = int64(16711680)
			v29 = int64(24)
			v31 = int64(4278190080)
			v33 = int64(8)
			*(*int64)(unsafe.Add(mBase, uint32(v17+v18))) = v9<<(uint(v20)%64) | v9&v22<<(uint(v24)%64) | (v9&v27<<(uint(v29)%64) | v9&v31<<(uint(v33)%64)) | (int64(base.Ui64(v9)>>(uint(v33)%64))&v31 | int64(base.Ui64(v9)>>(uint(v29)%64))&v27 | (int64(base.Ui64(v9)>>(uint(v24)%64))&v22 | int64(base.Ui64(v9)>>(uint(v20)%64))))
			v57 = v17 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v57
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(v60))) = v57 << (uint(int32(2)) % 32)
			m.G0 = v6 + int32(16)
			return v60
		}
	}
}
