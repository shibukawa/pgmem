package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SearchSysCacheAttName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_SearchSysCacheAttName[0]))
	v5 = F_SearchCatCache2(m, v4, l0, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v10)+91)))
			if v12 != int32(1) {
				return v5
			} else {
				F_ReleaseCatCache(m, v5)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		} else {
			return int32(0)
		}
	}
}
func F_SearchSysCacheExistsAttName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_SearchSysCacheExistsAttName[0]))
	v5 = F_SearchCatCache2(m, v4, l0, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v10)+91)))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v19 = v12 ^ int32(1)
				return v19 & int32(1)
			}
		} else {
			v19 = int32(0)
			return v19 & int32(1)
		}
	}
}
