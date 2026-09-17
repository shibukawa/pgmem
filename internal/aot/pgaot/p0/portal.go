package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetPortalByName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_GetPortalByName[0]))
		v8 = int32(0)
		v10 = F_hash_search(m, v7, l0, v8, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v10 != 0 {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
				v16 = v14
			} else {
				v16 = int32(0)
			}
			return v16
		}
	}
}
