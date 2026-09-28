package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OidInputFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = v9 + int32(4)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_OidInputFunctionCall[0]))
	F_fmgr_info_cxt_security(m, l0, v12, v14, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = F_InputFunctionCall(m, v12, l1, l2, l3)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			m.G0 = v9 + int32(32)
			return v20
		}
	}
}
func F_get_transform_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v12 = int64(0)
	v14 = F_GetSysCacheOid(m, int32(71), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1), v12, v12)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if l2|v14 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = F_format_type_be(m, l0)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v31 = F_get_language_name(m, l1, int32(0))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v31
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v28
							F_errmsg(m, int32(_a_F_get_transform_oid_0), v7)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_transform_oid_1), int32(2064), int32(_a_F_get_transform_oid_2))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
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
			return v14
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
