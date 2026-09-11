package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_make_and_qual(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == int32(0) {
		v31 = l1
		m.G0 = v6 + int32(16)
		return v31
	} else {
		if l1 == int32(0) {
			v31 = l0
			m.G0 = v6 + int32(16)
			return v31
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
			v18 = F_list_make2_impl(m, v6+int32(4), v6)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v23 = F_palloc0(m, int32(16))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v18
					*(*int64)(unsafe.Add(mBase, uint32(v23))) = int64(21)
					v31 = v23
					m.G0 = v6 + int32(16)
					return v31
				}
			}
		}
	}
}
