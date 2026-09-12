package p0

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_XmlTableGetValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	F_errstart_cold(m, int32(21), int32(0))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(380365), int32(0))
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(605506), int32(0))
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(520832), int32(5068), int32(364051))
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
		}
	}
}
