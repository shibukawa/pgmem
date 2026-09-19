package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XmlTableInitOpaque(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	Fn13843(m, l0, l1, int32(_a_F_XmlTableInitOpaque_0), int32(_a_F_XmlTableInitOpaque_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_XmlTableSetDocument(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	Fn13843(m, l0, l1, int32(_a_F_XmlTableSetDocument_0), int32(_a_F_XmlTableSetDocument_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_XmlTableSetNamespace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_F_XmlTableSetNamespace_0), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_errdetail(m, int32(_a_F_XmlTableSetNamespace_1), int32(0))
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_XmlTableSetNamespace_2), int32(_a_F_XmlTableSetNamespace_3), int32(_a_F_XmlTableSetNamespace_4))
					v23 = m.ExcPending
					if v23 != 0 {
						return
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
func F_xml_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_text_to_cstring(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_xml_recv(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13842(m, l0, int32(_a_F_xml_recv_0), int32(431))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
