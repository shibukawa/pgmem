package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LookupCollation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0 == int32(0) {
		v13 = F_get_collation_oid(m, l1, int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v37 = v13
			m.G0 = v8 + int32(32)
			return v37
		}
	} else {
		v18 = v8 + int32(12)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(489)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v18
		v24 = int32(_a_F_LookupCollation_0)
		v25 = *(*int32)(unsafe.Add(mBase, _c_F_LookupCollation[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v25
		*(*int32)(unsafe.Add(mBase, _c_F_LookupCollation[0])) = v8 + int32(20)
		v32 = F_get_collation_oid(m, l1, int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
			*(*int32)(unsafe.Add(mBase, _c_F_LookupCollation[0])) = v35
			v37 = v32
			m.G0 = v8 + int32(32)
			return v37
		}
	}
}
func F_get_collation_isdeterministic(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v5 = int32(16)
	v6 = v4 - v5
	m.G0 = v6
	v9 = F_SearchSysCache1(m, v5, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_collation_isdeterministic_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_collation_isdeterministic_1), int32(1154), int32(_a_F_get_collation_isdeterministic_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v29)+77)))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
