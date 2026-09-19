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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_DropDatabaseBuffers[0]))
	if v2 < v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = v2
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
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_DropDatabaseBuffers[1]))
	v23 = v20 + v17<<(uint(int32(6))%32)
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
	v96 = v17 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_DropDatabaseBuffers[0]))
	if v96 < v98 {
		v17 = v96
		goto L4
	} else {
		goto L31
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(_a_F_DropDatabaseBuffers_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(_a_F_DropDatabaseBuffers_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_DropDatabaseBuffers_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v36 = int32(_a_F_DropDatabaseBuffers_3)
	v38 = base.AtomicRmwOr32(m, v23, int32(24), v36)
	if v38&v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	goto L11
L9:
	;
	v59 = v38
	goto L10
L10:
	;
	v63 = int32(_a_F_DropDatabaseBuffers_4)
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_DropDatabaseBuffers[2]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v66 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v59 = v52
	goto L10
L13:
	;
	return
L14:
	;
	v50 = int32(_a_F_DropDatabaseBuffers_3)
	v52 = base.AtomicRmwOr32(m, v23, int32(24), v50)
	if v52&v50 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if l0 == v83 {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DropDatabaseBuffers[2])) = v81
	goto L17
L19:
	;
	if int32(999) < v64 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v64 < int32(11) {
		goto L17
	} else {
		goto L26
	}
L22:
	;
	v71 = int32(900)
	if v71 <= v64 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v74 = v71
	goto L25
L24:
	;
	v74 = v64
	goto L25
L25:
	;
	v81 = v74 + int32(100)
	goto L18
L26:
	;
	v81 = v64 - int32(1)
	goto L18
L27:
	;
	F_InvalidateBuffer(m, v23)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L13
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v59 & int32(-4194305)
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
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_database_to_xml[0]))
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
