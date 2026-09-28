package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_pa_set_stream_apply_worker(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_pa_set_stream_apply_worker[0])) = l0
	return
}
func F_pa_set_xact_state(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v5 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v5 != 0 {
		F_s_lock(m, l0, int32(_a_F_pa_set_xact_state_0))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
			v10 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v10))
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
		v10 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v10))
		return
	}
}
