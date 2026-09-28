package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cursor_to_xml(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v17 = F_text_to_cstring(m, v13)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v22 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = F_text_to_cstring(m, v22)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v27 = v10 - int32(-64)
	F_initStringInfo(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v19 == int64(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(_a_F_cursor_to_xml_0)
	F_appendStringInfo(m, v27, int32(_a_F_cursor_to_xml_1), v10+int32(48))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L18
	}
L10:
	;
	F_appendStringInfoString(m, v27, int32(_a_F_cursor_to_xml_2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v42 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v24
	F_appendStringInfo(m, v27, int32(_a_F_cursor_to_xml_3), v10+int32(32))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v50 = v10 - int32(-64)
	F_appendStringInfoString(m, v50, int32(_a_F_cursor_to_xml_4))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	F_appendStringInfoChar(m, v50, int32(10))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L9
L18:
	;
	v61 = F_GetPortalByName(m, v17)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v61 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_SPI_cursor_fetch(m, v61, v20)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L37
	}
L23:
	;
	v66 = *(*int64)(unsafe.Add(mBase, _c_F_cursor_to_xml[0]))
	if v66 != int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v77 = int64(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	v94 = F_SPI_finish(m)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L31
	}
L27:
	;
	F_SPI_sql_row_to_xmlelement(m, v10-int32(-64), base.B2i32(v19 != int64(0)), v24)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v83 = v77 + int64(1)
	v85 = *(*int64)(unsafe.Add(mBase, _c_F_cursor_to_xml[0]))
	if base.Ui64(v83) < base.Ui64(v85) {
		v77 = v83
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if v19 == int64(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_cursor_to_xml_0)
	F_appendStringInfo(m, v10-int32(-64), int32(_a_F_cursor_to_xml_5), v10+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
	v109 = F_cstring_to_text_with_len(m, v107, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	m.G0 = v10 + int32(80)
	return base.I64_extend_i32_u(v109)
L37:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v17
	F_errmsg(m, int32(_a_F_cursor_to_xml_6), v10)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_cursor_to_xml_7), int32(2980), int32(_a_F_cursor_to_xml_8))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
