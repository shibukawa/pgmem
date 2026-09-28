package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_encoding_max_length_sql(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.B2i32(base.Ui32(int32(41)) < base.Ui32(v3))|base.B2i32(v3 == int32(7)) == int32(0) {
		v15 = int64(*(*int32)(unsafe.Add(mBase, uint32(v3*int32(28))+uint32(_c_F_pg_encoding_max_length_sql[0]))))
		return v15
	} else {
		v17 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
		return int64(0)
	}
}
func F_pg_encoding_mblen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) != 0 {
		v18 = int32(1)
		return v18
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_c_F_pg_encoding_mblen[0])))
		v14 = m.T0[v13].(func(*base.Module, int32) int32)(m, l1)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = v14
			return v18
		}
	}
}
func F_pg_encoding_to_char_private(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) != 0 {
		v13 = int32(_a_F_pg_encoding_to_char_private_0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_pg_encoding_to_char_private[0])))
		v13 = v12
	}
	return v13
}
