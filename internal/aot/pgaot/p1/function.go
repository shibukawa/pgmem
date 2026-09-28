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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	v14 = v12 + v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
	if v15 != int32(2249) {
		v48 = v2
		m.G0 = v10 + int32(16)
		return v48
	} else {
		v20 = F_heap_attisnull(m, l0, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if v20 != 0 {
				v48 = v2
				m.G0 = v10 + int32(16)
				return v48
			} else {
				v26 = F_heap_attisnull(m, l0, int32(22), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					if v26 != 0 {
						v48 = v2
						m.G0 = v10 + int32(16)
						return v48
					} else {
						v30 = F_SysCacheGetAttrNotNull(m, int32(47), l0, int32(21))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v34 = F_SysCacheGetAttrNotNull(m, int32(47), l0, int32(22))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v40 = F_SysCacheGetAttr(m, int32(47), l0, int32(23), v10+int32(15))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+96)))
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
									if v44 != 0 {
										v45 = int64(0)
									} else {
										v45 = v40
									}
									v46 = F_build_function_result_tupdesc_d(m, v42, v30, v34, v45)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										v48 = v46
										m.G0 = v10 + int32(16)
										return v48
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
func F_has_function_privilege_id(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14302(m, l0, int32(_a_F_has_function_privilege_id_0), int32(1255))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_has_function_privilege_id_id(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14303(m, l0, int32(_a_F_has_function_privilege_id_id_0), int32(1255))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_has_function_privilege_name_name(m *base.Module, l0 int32) int64 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14307(m, l0, int32(_a_F_has_function_privilege_name_name_0), int32(1255), int32(_a_F_has_function_privilege_name_name_1), int32(3589), int32(_a_F_has_function_privilege_name_name_2), int32(52461700), int32(1365))
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		return v9
	}
}
