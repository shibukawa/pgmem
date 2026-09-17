package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OidInputFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = v8 + int32(4)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_OidInputFunctionCall[0]))
	F_fmgr_info_cxt_security(m, l0, v11, v13, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = F_InputFunctionCall(m, v11, l1, l2, l3)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(32)
			return v19
		}
	}
}
func F_get_transform_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v12 = F_GetSysCacheOid(m, int32(71), l0, l1, v4, v4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if l2|v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = F_format_type_be(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v29 = F_get_language_name(m, l1, int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v29
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v26
							F_errmsg(m, int32(_a_F_get_transform_oid_0), v7)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_transform_oid_1), int32(2049), int32(_a_F_get_transform_oid_2))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			}
		} else {
			m.G0 = v7 + int32(16)
			return v12
		}
	}
}
func F_oid_elem_desc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	F_appendStringInfo(m, l0, int32(_a_F_oid_elem_desc_0), v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
