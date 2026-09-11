package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SubqueryNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+52))
	if v3 != 0 {
		F_ExecReScan(m, v2)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return int32(0)
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v2)+12))
			v9 = m.T0[v8].(func(*base.Module, int32) int32)(m, v2)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				return v9
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v2)+12))
		v9 = m.T0[v8].(func(*base.Module, int32) int32)(m, v2)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
