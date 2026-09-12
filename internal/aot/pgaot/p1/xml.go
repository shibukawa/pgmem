package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XmlTableFetchRow(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(356992), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(554342), int32(0))
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(488832), int32(4912), int32(31403))
					v23 = m.ExcPending
					if v23 != 0 {
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
}
func F_XmlTableSetRowFilter(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	F_errstart_cold(m, int32(21), int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		F_errcode(m, int32(1088))
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_errmsg(m, int32(356992), int32(0))
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				F_errdetail(m, int32(554342), int32(0))
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					F_errfinish(m, int32(488832), int32(4837), int32(212528))
					v22 = m.ExcPending
					if v22 != 0 {
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
func F_map_xml_name_to_sql_identifier(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	var v57 int32
	_ = v57
	var v71 int32
	_ = v71
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	F_initStringInfo(m, v5+int32(-16))
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
	v15 = l0
	goto L3
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v19 != int32(95) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v123 = F_pg_mblen_cstr(m, v120)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L27
	}
L6:
	;
	v116 = F_pg_mblen_cstr(m, v15)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L25
	}
L7:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
	m.G0 = v7 - int32(-64)
	return v107
L8:
	;
	if v19 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v24 != int32(120) {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	goto L6
L12:
	;
	v28 = v15 + int32(2)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	goto L13
L13:
	;
	if base.B2i32(base.Ui32(v29-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v29|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)))
	goto L15
L15:
	;
	if base.B2i32(base.Ui32(v43-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v43|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
	goto L17
L17:
	;
	if base.B2i32(base.Ui32(v57-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v57|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+5)))
	goto L19
L19:
	;
	if base.B2i32(base.Ui32(v71-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v71|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v86 = v15 + int32(6)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v87 != int32(95) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v5 + int32(-52)
	v94 = F_sscanf(m, v28, int32(507204), v7)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	F_pg_unicode_to_server(m, v96, v5+int32(-48))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_appendStringInfoString(m, v5+int32(-16), v5+int32(-48))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v120 = v86
	goto L5
L25:
	;
	F_appendBinaryStringInfo(m, v5+int32(-16), v15, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v120 = v15
	goto L5
L27:
	;
	v15 = v123 + v120
	goto L3
}
func F_xml_is_well_formed_document(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(356992), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(554342), int32(0))
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(488832), int32(4628), int32(93280))
					v23 = m.ExcPending
					if v23 != 0 {
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
}
