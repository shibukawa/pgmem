package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_build_function_result_tupdesc_t(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
	v12 = v10 + v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	if v13 != int32(2249) {
		v47 = v2
		m.G0 = v8 + int32(16)
		return v47
	} else {
		v18 = F_heap_attisnull(m, l0, int32(21), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				v47 = v2
				m.G0 = v8 + int32(16)
				return v47
			} else {
				v24 = F_heap_attisnull(m, l0, int32(22), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v24 != 0 {
						v47 = v2
						m.G0 = v8 + int32(16)
						return v47
					} else {
						v28 = F_SysCacheGetAttrNotNull(m, int32(47), l0, int32(21))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v32 = F_SysCacheGetAttrNotNull(m, int32(47), l0, int32(22))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v38 = F_SysCacheGetAttr(m, int32(47), l0, int32(23), v8+int32(15))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+96)))
									v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
									if v42 != 0 {
										v43 = int32(0)
									} else {
										v43 = v38
									}
									v44 = F_build_function_result_tupdesc_d(m, v40, v28, v32, v43)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										v47 = v44
										m.G0 = v8 + int32(16)
										return v47
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_has_function_privilege_id(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13935(m, l0, int32(_a_F_has_function_privilege_id_0), int32(1255))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_has_function_privilege_id_id(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13936(m, l0, int32(_a_F_has_function_privilege_id_id_0), int32(1255))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_has_function_privilege_name_name(m *base.Module, l0 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13940(m, l0, int32(_a_F_has_function_privilege_name_name_0), int32(1255), int32(_a_F_has_function_privilege_name_name_1), int32(3565), int32(_a_F_has_function_privilege_name_name_2), int32(52461700), int32(1237))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
