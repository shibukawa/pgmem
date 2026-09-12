package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_replorigin_by_name(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_cstring_to_text(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_SearchSysCache1(m, int32(59), v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v15 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18)))
				F_ReleaseCatCache(m, v15)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v40 = v20
					m.G0 = v8 + int32(16)
					return v40 & int32(65535)
				}
			} else {
				if l1 != 0 {
					v40 = int32(0)
					m.G0 = v8 + int32(16)
					return v40 & int32(65535)
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							F_errmsg(m, int32(71888), v8)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(496373), int32(246), int32(378165))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
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
}
