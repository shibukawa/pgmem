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
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v37 int64
	_ = v37
	var v46 int64
	_ = v46
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int64
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
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
	v23 = v20 + v17*int32(56)
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
	v122 = v17 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_DropDatabaseBuffers[0]))
	if v122 < v124 {
		v17 = v122
		goto L4
	} else {
		goto L37
	}
L7:
	;
	v26 = int64(4194304)
	v28 = base.AtomicRmwOr64(m, v23, int32(24), v26)
	if v28&v26 != int64(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = v28
	goto L11
L9:
	;
	goto L10
L10:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if l0 == v109 {
		goto L33
	} else {
		goto L34
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(_a_F_DropDatabaseBuffers_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(_a_F_DropDatabaseBuffers_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_DropDatabaseBuffers_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	v46 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v46
	if v37&int64(4194304) != v46 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	goto L16
L14:
	;
	goto L15
L15:
	;
	v77 = int32(_a_F_DropDatabaseBuffers_3)
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_DropDatabaseBuffers[2]))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v80 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	return
L19:
	;
	v61 = int64(0)
	v64 = base.AtomicRmwCmpxchg64(m, v23, int32(24), v61, v61)
	if v64&int64(4194304) != v61 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v97 = int64(4194304)
	v99 = base.AtomicRmwOr64(m, v23, int32(24), v97)
	if v99&v97 != int64(0) {
		v37 = v99
		goto L11
	} else {
		goto L32
	}
L22:
	;
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DropDatabaseBuffers[2])) = v95
	goto L22
L24:
	;
	if int32(999) < v78 {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v78 < int32(11) {
		goto L22
	} else {
		goto L31
	}
L27:
	;
	v85 = int32(900)
	if v85 <= v78 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v88 = v85
	goto L30
L29:
	;
	v88 = v78
	goto L30
L30:
	;
	v95 = v88 + int32(100)
	goto L23
L31:
	;
	v95 = v78 - int32(1)
	goto L23
L32:
	;
	goto L12
L33:
	;
	F_InvalidateBuffer(m, v23)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L18
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v115 = base.AtomicRmwSub64(m, v23, int32(24), int64(4194304))
	goto L6
L36:
	;
	goto L6
L37:
	;
	goto L5
}
func F_database_to_xml(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = F_text_to_cstring(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int64(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_database_to_xml[0]))
			v11 = F_get_database_name(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int64(0)
			} else {
				v13 = F_map_sql_identifier_to_xml_name(m)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
