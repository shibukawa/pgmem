package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LockTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v2 = int32(_a_F_LockTimeoutHandler_0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_LockTimeoutHandler[0]))
	v5 = int32(2)
	v6 = F_pgmem_kill(m, int32(0)-v3, v5)
	mBase = m.M
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_LockTimeoutHandler[0]))
	v10 = F_pgmem_kill(m, v8, v5)
	mBase = m.M
	return
}
func F_get_timeout_finish_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0*int32(40))+uint32(_c_F_get_timeout_finish_time[0])))
	return v4
}
func F_get_timeout_indicator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v4 = l0 * int32(40)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+uint32(_c_F_get_timeout_indicator[0]))))
	if v5&int32(1) != 0 {
		v10 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+uint32(_c_F_get_timeout_indicator[0]))) = uint8(v10)
	} else {
	}
	return v5 & int32(1)
}
