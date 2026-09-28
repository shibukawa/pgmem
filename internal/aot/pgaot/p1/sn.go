package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SN_delete_env(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v3 != 0 {
			F_lose_s(m, v3)
			mBase = m.M
			v5 = m.ExcPending
			if v5 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v7 = m.ExcPending
				if v7 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v7 = m.ExcPending
			if v7 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_sn_array_end(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14376(m, l0, int32(93))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
