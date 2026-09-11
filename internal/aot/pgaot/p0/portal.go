package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetPortalByName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v2 = int32(0)
	if l0 == v2 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_GetPortalByName[0]))
		v9 = int32(0)
		v11 = F_hash_search(m, v8, l0, v9, v9)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v11 != 0 {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
				v16 = v15
			} else {
				v16 = v2
			}
			return v16
		}
	}
}
