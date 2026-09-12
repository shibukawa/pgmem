package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DropDatabaseBuffers(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	if v2 < v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = v2
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v8 + int32(32)
	return
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v23 = v20 + v18<<(uint(int32(6))%32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v24 != l0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v98 = v18 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	if v98 < v100 {
		v18 = v98
		goto L4
	} else {
		goto L31
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(220922)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(478364)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v37 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v36 | v37
	if v36&v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	goto L11
L9:
	;
	v60 = v36
	goto L10
L10:
	;
	v65 = int32(4074876)
	v66 = *(*int32)(unsafe.Add(mBase, _consts[605]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v68 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v60 = v51
	goto L10
L13:
	;
	return
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v52 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v51 | v52
	if v51&v52 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if l0 == v85 {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[605])) = v83
	goto L17
L19:
	;
	if int32(999) < v66 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v66 < int32(11) {
		goto L17
	} else {
		goto L26
	}
L22:
	;
	v73 = int32(900)
	if v73 <= v66 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v76 = v73
	goto L25
L24:
	;
	v76 = v66
	goto L25
L25:
	;
	v83 = v76 + int32(100)
	goto L18
L26:
	;
	v83 = v66 - int32(1)
	goto L18
L27:
	;
	F_InvalidateBuffer(m, v23)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L13
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v60 & int32(-4194305)
	goto L6
L30:
	;
	goto L6
L31:
	;
	goto L5
}
func F_database_to_xml(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v3 = F_pg_detoast_datum_packed(m, v2)
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
			v10 = *(*int32)(unsafe.Add(mBase, _consts[108]))
			v11 = F_get_database_name(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				v13 = F_map_sql_identifier_to_xml_name(m)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
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
