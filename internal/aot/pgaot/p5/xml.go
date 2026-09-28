package p5

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_xml_is_well_formed(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14411(m, l0, int32(_a_F_xml_is_well_formed_0), int32(_a_F_xml_is_well_formed_1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
