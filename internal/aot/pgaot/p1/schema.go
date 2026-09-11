package p1

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_is_schema_publication(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v11 = F_table_open(m, int32(_a_F_is_schema_publication_0), int32(1))
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		F_ScanKeyInit(m, v7, int32(2), int32(3), int32(184), l0)
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = int32(1)
			v24 = F_systable_beginscan(m, v11, int32(_a_F_is_schema_publication_1), v21, int32(0), v21, v7)
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = F_systable_getnext(m, v24)
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_systable_endscan(m, v24)
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_sequence_close(m, v11, int32(1))
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(48)
							return base.B2i32(v26 != int32(0))
						}
					}
				}
			}
		}
	}
}
