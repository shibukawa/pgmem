package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pa_set_stream_apply_worker(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[669])) = l0
	return
}
func F_pa_set_xact_state(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v4 != 0 {
		F_s_lock(m, l0, int32(496843), int32(1317), int32(352622))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
		return
	}
}
