package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_opclass_opfamily_and_input_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v6 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(l0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
			v12 = v10 + v11
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v13
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v15
			F_ReleaseCatCache(m, v6)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v6 != int32(0))
			}
		} else {
			return base.B2i32(v6 != int32(0))
		}
	}
}
