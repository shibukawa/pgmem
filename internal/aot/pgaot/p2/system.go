package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SystemTypeName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = F_makeString(m, int32(_a_F_SystemTypeName_0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v8
		v13 = F_makeString(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v13
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v13
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v17
			v21 = F_list_make2_impl(m, v5+int32(4), v5)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_makeTypeNameFromNameList(m, v21)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(16)
					return v23
				}
			}
		}
	}
}
