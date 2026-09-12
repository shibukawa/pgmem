package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cursor_to_xml(m *base.Module, l0 int32) int32 {
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
	var v19 int32
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
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v82 int64
	_ = v82
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
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
	F_initStringInfo(m, v10-int32(-64))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v19 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(_a_F_cursor_to_xml_0)
	F_appendStringInfo(m, v10-int32(-64), int32(_a_F_cursor_to_xml_1), v10+int32(48))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
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
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L10:
	;
	F_appendStringInfoString(m, v10-int32(-64), int32(_a_F_cursor_to_xml_2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v46 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v24
	F_appendStringInfo(m, v10-int32(-64), int32(_a_F_cursor_to_xml_3), v10+int32(32))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_appendStringInfoString(m, v10-int32(-64), int32(_a_F_cursor_to_xml_4))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	F_appendStringInfoChar(m, v10-int32(-64), int32(10))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L9
L18:
	;
	v68 = F_GetPortalByName(m, v17)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v68 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_SPI_cursor_fetch(m, v68, v20)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
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
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L37
	}
L23:
	;
	v73 = *(*int64)(unsafe.Add(mBase, _c_F_cursor_to_xml[0]))
	if v73 != int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v82 = int64(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	v101 = F_SPI_finish(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L31
	}
L27:
	;
	F_SPI_sql_row_to_xmlelement(m, v10-int32(-64), base.B2i32(v19 != int32(0)), v24)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v90 = v82 + int64(1)
	v92 = *(*int64)(unsafe.Add(mBase, _c_F_cursor_to_xml[0]))
	if base.Ui64(v90) < base.Ui64(v92) {
		v82 = v90
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if v19 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_cursor_to_xml_0)
	F_appendStringInfo(m, v10-int32(-64), int32(_a_F_cursor_to_xml_5), v10+int32(16))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
	v116 = F_cstring_to_text_with_len(m, v114, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
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
	return v116
L37:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v17
	F_errmsg(m, int32(_a_F_cursor_to_xml_6), v10)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_cursor_to_xml_7), int32(2937), int32(_a_F_cursor_to_xml_8))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
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
