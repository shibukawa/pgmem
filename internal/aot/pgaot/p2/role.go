package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_role_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v13 = F_GetSysCacheOid(m, int32(10), l0, v3, v3, v3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if l1 != 0 {
			m.G0 = v7 + int32(16)
			return v13
		} else {
			if v13 != 0 {
				m.G0 = v7 + int32(16)
				return v13
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg(m, int32(72247), v7)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(497340), int32(5562), int32(434753))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
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
	}
}
func F_show_role(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, _consts[427]))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[428])))
	if v7 != 0 {
		v8 = v4
	} else {
		v8 = int32(0)
	}
	v10 = *(*int32)(unsafe.Add(mBase, _consts[429]))
	if v10 != 0 {
		v12 = v10
	} else {
		v12 = int32(372507)
	}
	if v8 != 0 {
		v14 = v12
	} else {
		v14 = int32(372507)
	}
	return v14
}
