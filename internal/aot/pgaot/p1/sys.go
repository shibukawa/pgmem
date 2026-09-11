package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SearchSysCache2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_SearchSysCache2[0])))
	v9 = F_SearchCatCache2(m, v8, l1, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_SearchSysCache4(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_SearchSysCache4[0])))
	v12 = F_SearchCatCacheInternal(m, v10, int32(4), l1, l2, l3, l4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
func F_SearchSysCacheExists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_SearchSysCacheExists[0])))
	v11 = F_SearchCatCache(m, v10, l1, l2, l3, l4)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			F_ReleaseCatCache(m, v11)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v11 != int32(0))
			}
		} else {
			return base.B2i32(v11 != int32(0))
		}
	}
}
func F_SearchSysCacheLockedCopy1(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v3 = F_SearchSysCacheLocked1(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			return int32(0)
		} else {
			v11 = F_heap_copytuple(m, v3)
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v3)
				v14 = m.ExcPending
				if v14 != 0 {
					return int32(0)
				} else {
					return v11
				}
			}
		}
	}
}
