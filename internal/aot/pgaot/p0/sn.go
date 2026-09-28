package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SN_new_env(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v13 int64
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v3 = F_palloc(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3)+24)) = int32(0)
			v13 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v3)+16)) = v13
			*(*int64)(unsafe.Add(mBase, uint32(v3)+8)) = v13
			*(*int64)(unsafe.Add(mBase, uint32(v3))) = v13
			v19 = F_create_s(m)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v3))) = v19
				if v19 != 0 {
					return v3
				} else {
					F_pfree(m, v3)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				}
			}
		}
	}
}
