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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_cstring_to_text(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = F_SearchSysCache1(m, int32(59), base.I64_extend_i32_u(v11))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v16 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v18+v19)))
				F_ReleaseCatCache(m, v16)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v41 = v21
					m.G0 = v8 + int32(16)
					return v41 & int32(_a_F_replorigin_by_name_0)
				}
			} else {
				if l1 != 0 {
					v41 = int32(0)
					m.G0 = v8 + int32(16)
					return v41 & int32(_a_F_replorigin_by_name_0)
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							F_errmsg(m, int32(_a_F_replorigin_by_name_1), v8)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_replorigin_by_name_2), int32(263), int32(_a_F_replorigin_by_name_3))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
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
