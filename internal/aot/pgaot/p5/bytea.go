package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bytea_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	v3 = int32(_a_F_bytea_sortsupport_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_bytea_sortsupport[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, _c_F_bytea_sortsupport[0])) = v7
	F_varstr_sortsupport(m, v6, int32(17), int32(950))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_bytea_sortsupport[0])) = v4
		return int32(0)
	}
}
