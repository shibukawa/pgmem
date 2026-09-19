package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_role_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13921(m, l0, l1, int32(_a_F_get_role_oid_0), int32(_a_F_get_role_oid_1), int32(_a_F_get_role_oid_2), int32(_a_F_get_role_oid_3), int32(67137668), int32(10))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_show_role(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_show_role[0]))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_show_role[1])))
	if v7 != 0 {
		v8 = v4
	} else {
		v8 = int32(0)
	}
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_show_role[2]))
	if v10 != 0 {
		v12 = v10
	} else {
		v12 = int32(_a_F_show_role_0)
	}
	if v8 != 0 {
		v14 = v12
	} else {
		v14 = int32(_a_F_show_role_0)
	}
	return v14
}
