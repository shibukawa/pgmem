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
			F_errmsg(m, int32(356759), int32(0))
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_errdetail(m, int32(553927), int32(0))
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_errfinish(m, int32(488417), int32(5115), int32(338569))
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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v15 = F_text_to_cstring(m, v8)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v15&int32(3) == int32(0) {
		v42 = v15
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_pq_sendtext(m, v5, v15, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L23
	}
L7:
	;
	v75 = v67 - v15
	goto L6
L8:
	;
	v46 = v42
	goto L17
L9:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v26 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v75 = int32(0)
	goto L6
L11:
	;
	goto L12
L12:
	;
	v31 = v15
	goto L13
L13:
	;
	v35 = v31 + int32(1)
	if v35&int32(3) == int32(0) {
		v42 = v35
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v67 = v35
	goto L7
L15:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v40 != 0 {
		v31 = v35
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v55 = int32(-2139062144)
	if (int32(16843008)-v52|v52)&v55 == v55 {
		v46 = v46 + int32(4)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v61 = v46
	goto L20
L19:
	;
	goto L18
L20:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v65 != 0 {
		v61 = v61 + int32(1)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v67 = v61
	goto L7
L22:
	;
	goto L21
L23:
	;
	F_pfree(m, v15)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v82 << (uint(int32(2)) % 32)
	goto L25
L25:
	;
	m.G0 = v5 + int32(16)
	return v81
}
