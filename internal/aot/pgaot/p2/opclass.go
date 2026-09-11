package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_opclass_opfamily_and_input_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v5 = F_SearchSysCache1(m, int32(14), l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
			v11 = v9 + v10
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v12
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v14
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v5 != int32(0))
			}
		} else {
			return base.B2i32(v5 != int32(0))
		}
	}
}
