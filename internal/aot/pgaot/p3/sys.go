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
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_SearchSysCacheAttName[0]))
	v7 = F_SearchCatCache2(m, v4, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v12)+91)))
			if v14 != int32(1) {
				return v7
			} else {
				F_ReleaseCatCache(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
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
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_SearchSysCacheExistsAttName[0]))
	v7 = F_SearchCatCache2(m, v4, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v12)+91)))
			F_ReleaseCatCache(m, v7)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v21 = v14 ^ int32(1)
				return v21 & int32(1)
			}
		} else {
			v21 = int32(0)
			return v21 & int32(1)
		}
	}
}
