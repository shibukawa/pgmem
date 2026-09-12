package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XmlTableDestroyOpaque(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	F_errstart_cold(m, int32(21), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_errcode(m, int32(1088))
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_errmsg(m, int32(379781), int32(0))
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_errdetail(m, int32(604837), int32(0))
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_errfinish(m, int32(520163), int32(5115), int32(360843))
					v21 = m.ExcPending
					if v21 != 0 {
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
func F_xml_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v15 = F_text_to_cstring(m, v8)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_pq_begintypsend(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = F_strlen(m, v15)
				mBase = m.M
				F_pq_sendtext(m, v5, v15, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v15)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v25))) = v26 << (uint(int32(2)) % 32)
						m.G0 = v5 + int32(16)
						return v25
					}
				}
			}
		}
	}
}
