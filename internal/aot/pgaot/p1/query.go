package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_query_contains_extern_params_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v8 != int32(67) {
			if v8 != int32(8) {
				v25 = F_expression_tree_walker_impl(m, l0, int32(496), l1)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					return v25
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				return base.B2i32(v13 == int32(0))
			}
		} else {
			v19 = F_query_tree_walker_impl(m, l0, int32(496), l1, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return v19
			}
		}
	}
}
func F_query_to_xmlschema(m *base.Module, l0 int32) int32 {
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
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
	v26 = int32(0)
	v28 = F_SPI_prepare(m, v15, v26, v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L43
	}
L8:
	;
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v30 = F_SPI_cursor_open(m, v28)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
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
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L40
	}
L12:
	;
	if v30 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+92))
	v35 = int32(0)
	v38 = F_map_sql_table_to_xmlschema(m, v34, v35, base.B2i32(v17 != v35), v21)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v38&int32(3) == int32(0) {
		v63 = v38
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v98 = v96 + int32(1)
	v99 = F_SPI_palloc(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L32
	}
L16:
	;
	v96 = v88 - v38
	goto L15
L17:
	;
	v67 = v63
	goto L26
L18:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v47 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v96 = int32(0)
	goto L15
L20:
	;
	goto L21
L21:
	;
	v52 = v38
	goto L22
L22:
	;
	v56 = v52 + int32(1)
	if v56&int32(3) == int32(0) {
		v63 = v56
		goto L17
	} else {
		goto L24
	}
L23:
	;
	v88 = v56
	goto L16
L24:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v61 != 0 {
		v52 = v56
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v76 = int32(-2139062144)
	if (int32(16843008)-v73|v73)&v76 == v76 {
		v67 = v67 + int32(4)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v82 = v67
	goto L29
L28:
	;
	goto L27
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v86 != 0 {
		v82 = v82 + int32(1)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v88 = v82
	goto L16
L31:
	;
	goto L30
L32:
	;
	if v98 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	F_SPI_cursor_close(m, v30)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L37
	}
L34:
	;
	v101 = F__emscripten_memcpy_bulkmem(m, v99, v38, v98)
	mBase = m.M
	v102 = v101
	goto L36
L35:
	;
	v102 = v99
	goto L36
L36:
	;
	goto L33
L37:
	;
	v105 = F_SPI_finish(m)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v107 = F_cstring_to_text(m, v102)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	m.G0 = v8 + int32(32)
	return v107
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
	F_errmsg_internal(m, int32(433358), v8)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(474019), int32(3077), int32(482204))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v15
	F_errmsg_internal(m, int32(433329), v8+int32(16))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(474019), int32(3080), int32(482204))
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
