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
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v3*int32(28))+uint32(_consts[864])))
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	if base.Ui32(l0) <= base.Ui32(int32(41)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_consts[953])))
		v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, l1)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = v12
			return v16
		}
	} else {
		v16 = int32(1)
		return v16
	}
}
func F_pg_encoding_to_char_private(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if base.Ui32(l0) <= base.Ui32(int32(41)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_consts[769])))
		v11 = v10
	} else {
		v11 = int32(722455)
	}
	return v11
}
