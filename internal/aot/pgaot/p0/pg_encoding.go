package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_encoding_max_length_sql(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v3) <= base.Ui32(int32(41)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v3*int32(28))+uint32(_c_F_pg_encoding_max_length_sql[0])))
		return v10
	} else {
		v12 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
		return int32(0)
	}
}
func F_pg_encoding_mblen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	if base.Ui32(l0) <= base.Ui32(int32(41)) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_c_F_pg_encoding_mblen[0])))
		v8 = m.T0[v7].(func(*base.Module, int32) int32)(m, l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v13 = v8
			return v13
		}
	} else {
		v13 = int32(1)
		return v13
	}
}
func F_pg_encoding_to_char_private(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if base.Ui32(l0) <= base.Ui32(int32(41)) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_pg_encoding_to_char_private[0])))
		v8 = v6
	} else {
		v8 = int32(_a_F_pg_encoding_to_char_private_0)
	}
	return v8
}
