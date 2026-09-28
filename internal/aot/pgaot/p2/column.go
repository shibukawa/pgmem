package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_convert_column_name(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_text_to_cstring(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = F_SearchSysCache2(m, int32(6), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(v12))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v17 != 0 {
				v19 = int32(0)
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
				v22 = v20 + v21
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+91)))
				if v23 == v19 {
					v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+74)))
					v27 = v26
				} else {
					v27 = v19
				}
				F_ReleaseCatCache(m, v17)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v53 = v27
					F_pfree(m, v12)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return base.I32_extend16_s(v53)
					}
				}
			} else {
				v31 = F_get_rel_name(m, l0)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v31 == int32(0) {
						v53 = int32(0)
						F_pfree(m, v12)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return base.I32_extend16_s(v53)
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50360452))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v31
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
								F_errmsg(m, int32(_a_F_convert_column_name_0), v8)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_convert_column_name_1), int32(2963), int32(_a_F_convert_column_name_2))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
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
}
