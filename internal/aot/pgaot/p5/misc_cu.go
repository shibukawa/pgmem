package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cursor_to_xmlschema(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = F_text_to_cstring(m, v11)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = F_text_to_cstring(m, v19)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v26 = F_GetPortalByName(m, v15)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L42
	}
L8:
	;
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+92))
	if v28 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L38
	}
L12:
	;
	v31 = int32(0)
	v34 = F_map_sql_table_to_xmlschema(m, v28, v31, base.B2i32(v17 != v31), v21)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v34&int32(3) == int32(0) {
		v59 = v34
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v94 = v92 + int32(1)
	v95 = F_SPI_palloc(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L15:
	;
	v92 = v84 - v34
	goto L14
L16:
	;
	v63 = v59
	goto L25
L17:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v43 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v92 = int32(0)
	goto L14
L19:
	;
	goto L20
L20:
	;
	v48 = v34
	goto L21
L21:
	;
	v52 = v48 + int32(1)
	if v52&int32(3) == int32(0) {
		v59 = v52
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v84 = v52
	goto L15
L23:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v57 != 0 {
		v48 = v52
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v72 = int32(-2139062144)
	if (int32(16843008)-v69|v69)&v72 == v72 {
		v63 = v63 + int32(4)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v78 = v63
	goto L28
L27:
	;
	goto L26
L28:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v82 != 0 {
		v78 = v78 + int32(1)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v84 = v78
	goto L15
L30:
	;
	goto L29
L31:
	;
	if v94 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v99 = F_SPI_finish(m)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L36
	}
L33:
	;
	v97 = F__emscripten_memcpy_bulkmem(m, v95, v34, v94)
	mBase = m.M
	v98 = v97
	goto L35
L34:
	;
	v98 = v95
	goto L35
L35:
	;
	goto L32
L36:
	;
	v101 = F_cstring_to_text(m, v98)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	m.G0 = v8 + int32(32)
	return v101
L38:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
	F_errmsg(m, int32(_a_F_cursor_to_xmlschema_0), v8)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_cursor_to_xmlschema_1), int32(3107), int32(_a_F_cursor_to_xmlschema_2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_errcode(m, int32(258))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v15
	F_errmsg(m, int32(_a_F_cursor_to_xmlschema_3), v8+int32(16))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_cursor_to_xmlschema_1), int32(3111), int32(_a_F_cursor_to_xmlschema_2))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
