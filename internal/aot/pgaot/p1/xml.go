package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XmlTableFetchRow(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13862(m, l0, int32(_a_F_XmlTableFetchRow_0), int32(_a_F_XmlTableFetchRow_1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_XmlTableSetRowFilter(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	Fn13863(m, l0, l1, int32(_a_F_XmlTableSetRowFilter_0), int32(_a_F_XmlTableSetRowFilter_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_map_xml_name_to_sql_identifier(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v39 int32
	_ = v39
	var v53 int32
	_ = v53
	var v67 int32
	_ = v67
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	F_initStringInfo(m, v4+int32(-16))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = l0
	goto L3
L3:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v17 != int32(95) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v116 = F_pg_mblen_cstr(m, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L27
	}
L6:
	;
	v110 = F_pg_mblen_cstr(m, v14)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L25
	}
L7:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
	m.G0 = v6 - int32(-64)
	return v103
L8:
	;
	if v17 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v22 != int32(120) {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	goto L6
L12:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)))
	goto L13
L13:
	;
	if base.B2i32(base.Ui32(v25-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v25|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)))
	goto L15
L15:
	;
	if base.B2i32(base.Ui32(v39-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v39|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)))
	goto L17
L17:
	;
	if base.B2i32(base.Ui32(v53-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v53|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)))
	goto L19
L19:
	;
	if base.B2i32(base.Ui32(v67-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v67|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+6)))
	if v81 != int32(95) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v4 + int32(-52)
	v92 = F_sscanf(m, v14+int32(2), int32(_a_F_map_xml_name_to_sql_identifier_0), v6)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v96 = v4 + int32(-48)
	F_pg_unicode_to_server(m, v94, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_appendStringInfoString(m, v4+int32(-16), v96)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v114 = v14 + int32(6)
	goto L5
L25:
	;
	F_appendBinaryStringInfo(m, v4+int32(-16), v14, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v114 = v14
	goto L5
L27:
	;
	v14 = v116 + v114
	goto L3
}
func F_xml_is_well_formed_document(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13862(m, l0, int32(_a_F_xml_is_well_formed_document_0), int32(_a_F_xml_is_well_formed_document_1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
