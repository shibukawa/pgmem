package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetCurrentCommandId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentCommandId[0]))
		if int32(0) <= v3 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(322))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_GetCurrentCommandId_0), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetCurrentCommandId_1), int32(843), int32(_a_F_GetCurrentCommandId_2))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v7 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_GetCurrentCommandId[1])) = uint8(v7)
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentCommandId[2]))
			return v10
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentCommandId[2]))
		return v10
	}
}
func F_current_database(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = F_palloc(m, int32(64))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_current_database[0]))
		v9 = F_get_database_name(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v12 = F_strncpy(m, v3, v9, int32(64))
			mBase = m.M
			v13 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+63)) = uint8(v13)
			return v3
		}
	}
}
func F_current_schema(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v5 = F_fetch_search_path(m, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			v17 = F_get_namespace_name(m, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_list_free(m, v5)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v17 == int32(0) {
						v23 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
						return int32(0)
					} else {
						v29 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v17)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							return v29
						}
					}
				}
			}
		}
	}
}
func F_get_current_ts_config(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_getTSCurrentConfig(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
