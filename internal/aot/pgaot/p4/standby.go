package p4

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_standby_identify(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	v2 = int32(0)
	v4 = l0 & int32(240)
	switch v4 - int32(16) {
	case 0:
		return int32(543499)
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		v12 = v2
		return v12
	case 16:
		v12 = int32(544022)
		return v12
	default:
		if v4 != 0 {
			v12 = v2
			return v12
		} else {
			return int32(555008)
		}
	}
}
