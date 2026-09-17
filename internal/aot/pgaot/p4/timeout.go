package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LockTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_LockTimeoutHandler[0]))
	v6 = F_kill(m, int32(0)-v3, int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_LockTimeoutHandler[0]))
		v11 = F_kill(m, v9, int32(2))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
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
