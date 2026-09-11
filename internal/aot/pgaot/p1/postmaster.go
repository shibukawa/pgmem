package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsPostmasterChildWalSender(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_IsPostmasterChildWalSender[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v3+l0<<(uint(int32(2))%32))+44))
	return base.B2i32(v7 == int32(3))
}
func F_PostmasterChildName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(12))+uint32(_c_F_PostmasterChildName[0])))
	return v6
}
func F_RegisterPostmasterChildActive(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterPostmasterChildActive[0]))
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterPostmasterChildActive[1]))
	v5 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v2+v4<<(uint(v5)%32))+44)) = v5
	F_on_shmem_exit(m, int32(1100), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		return
	}
}
