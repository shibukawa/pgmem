package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_relation_is_publishable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_SearchSysCache1(m, int32(57), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		} else {
			v17 = int32(0)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
			v20 = v18 + v19
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+119)))
			switch v21 - int32(112) {
			case 0, 2:
				if base.Ui32(v6) < base.Ui32(int32(12000)) {
					v32 = v17
				} else {
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+118)))
					v32 = base.B2i32(v26 == int32(112)) & base.B2i32(base.Ui32(int32(16383)) < base.Ui32(v6))
				}
			default:
				v32 = v17
			}
			F_ReleaseCatCache(m, v7)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				return v32
			}
		}
	}
}
