package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_isQueryUsingTempRelation(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_isQueryUsingTempRelation_walker(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_query_has_required_values(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 <= int32(0) {
		return int32(0)
	} else {
		v11 = F_contains_required_value(m, l0+v3<<(uint(int32(3))%32))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_query_to_xml(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
			v9 = int32(0)
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v15 = F_pg_detoast_datum_packed(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = F_text_to_cstring(m, v15)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = F_query_to_xml_internal(m, v7, v9, v9, base.B2i32(v11 != v9), v17)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
						v23 = F_cstring_to_text_with_len(m, v21, v22)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							return v23
						}
					}
				}
			}
		}
	}
}
func F_query_to_xml_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if l1 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L45
	}
L2:
	;
	v14 = F_makeStringInfo(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v114 = F_map_sql_identifier_to_xml_name(m)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L44
	}
L5:
	;
	return int32(0)
L6:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v23 = F_SPI_execute(m, l0, int32(1), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v23 != int32(5) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if l3 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v29 = m.G0
	v31 = v29 - int32(48)
	m.G0 = v31
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = int32(412760)
	F_appendStringInfo(m, v14, int32(186782), v31+int32(32))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if l2 != 0 {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	F_appendStringInfoString(m, v14, int32(761246))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v43 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = l4
	F_appendStringInfo(m, v14, int32(721315), v31+int32(16))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if l2 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	F_appendStringInfoString(m, v14, int32(786139))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L26
	}
L20:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v52 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = l4
	F_appendStringInfo(m, v14, int32(762781), v31)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_appendStringInfoString(m, v14, int32(762746))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L25
	}
L24:
	;
	goto L19
L25:
	;
	goto L19
L26:
	;
	m.G0 = v31 + int32(48)
	F_appendStringInfoChar(m, v14, int32(10))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	goto L12
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l2
	F_appendStringInfo(m, v14, int32(789045), v10+int32(16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v77 = *(*int64)(unsafe.Add(mBase, _consts[510]))
	if v77 != int64(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	v86 = int64(0)
	goto L35
L33:
	;
	goto L34
L34:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	F_SPI_sql_row_to_xmlelement(m, v14, l3, l4)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L37
	}
L36:
	;
	goto L34
L37:
	;
	v90 = v86 + int64(1)
	v92 = *(*int64)(unsafe.Add(mBase, _consts[510]))
	if base.Ui64(v90) < base.Ui64(v92) {
		v86 = v90
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(412760)
	F_appendStringInfo(m, v14, int32(784695), v10)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v108 = F_SPI_finish(m)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	m.G0 = v10 + int32(32)
	return v14
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(16987), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(520163), int32(3019), int32(325349))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
