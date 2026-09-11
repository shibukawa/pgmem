package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReleaseCatCache(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseCatCache[0]))
	v7 = l0 - int32(8)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v8 - int32(1)
	if v5 != 0 {
		F_ResourceOwnerForget(m, v5, l0, int32(_a_F_ReleaseCatCache_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0-int32(4)))))
			if v17 != int32(1) {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				if v20 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v21 != 0 {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
						if v22 != 0 {
							return
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							F_CatCacheRemoveCTup(m, v23, l0-int32(40))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						F_CatCacheRemoveCTup(m, v23, l0-int32(40))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0-int32(4)))))
		if v17 != int32(1) {
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v21 != 0 {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
					if v22 != 0 {
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						F_CatCacheRemoveCTup(m, v23, l0-int32(40))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					F_CatCacheRemoveCTup(m, v23, l0-int32(40))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_SearchCatCache1(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = int32(0)
	v7 = F_SearchCatCacheInternal(m, l0, int32(1), l1, v4, v4, v4)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
