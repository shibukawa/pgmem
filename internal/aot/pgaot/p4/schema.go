package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetSchemaPublicationRelations(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v15 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = int32(3)
	F_ScanKeyInit(m, v11, v19, v19, int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = F_table_beginscan_catalog(m, v15, int32(1), v11)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = F_heap_getnext(m, v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v30 = v28
	v32 = v3
	goto L9
L7:
	;
	v74 = v3
	goto L8
L8:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+188))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	m.T0[v82].(func(*base.Module, int32))(m, v26)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L26
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
	v40 = v38 + v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+119)))
	switch v42 - int32(112) {
	case 0, 2:
		goto L12
	case 1:
		v69 = v32
		goto L11
	default:
		goto L13
	}
L10:
	;
	v74 = v69
	goto L8
L11:
	;
	v70 = F_heap_getnext(m, v26)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L24
	}
L12:
	;
	goto L15
L13:
	;
	if v42 != int32(83) {
		v69 = v32
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if base.B2i32(base.Ui32(v41) < base.Ui32(int32(_a_F_GetSchemaPublicationRelations_0)))|base.B2i32(base.Ui32(v41) < base.Ui32(int32(_a_F_GetSchemaPublicationRelations_1))) != 0 {
		v69 = v32
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+118)))
	if v52 != int32(112) {
		v69 = v32
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v55 = F_get_rel_relkind(m, v41)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v65 = F_GetPubPartitionOptionRelations(m, int32(0), l1, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	v61 = F_lappend_oid(m, v32, v41)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	switch v55&int32(255) - int32(112) {
	case 0:
		goto L18
	default:
		v69 = v32
		goto L11
	case 2:
		goto L19
	}
L21:
	;
	v69 = v61
	goto L11
L22:
	;
	v67 = F_list_concat_unique_oid(m, v32, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v69 = v67
	goto L11
L24:
	;
	if v70 != 0 {
		v30 = v70
		v32 = v69
		goto L9
	} else {
		goto L25
	}
L25:
	;
	goto L10
L26:
	;
	F_relation_close(m, v15, int32(1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	m.G0 = v11 - int32(-64)
	return v74
}
func F_checkSchemaNameRV(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L15
	} else {
		goto L21
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if base.B2i32(v13 == int32(0))|base.B2i32(v13 != v16) != 0 {
		v34 = v13
		v35 = v16
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v37 == int32(116) {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	if v34-v35 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	goto L6
L8:
	;
	v19 = v10
	v20 = v9
	goto L9
L9:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v24 == int32(0) {
		v34 = v24
		v35 = v23
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v34 = v24
	v35 = v23
	goto L7
L11:
	;
	v27 = int32(1)
	if v24 == v23 {
		v19 = v19 + v27
		v20 = v20 + v27
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	goto L5
L14:
	;
	m.G0 = v7 + int32(16)
	return
L15:
	;
	return
L16:
	;
	F_errcode(m, int32(84279428))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v50
	F_errmsg(m, int32(_a_F_checkSchemaNameRV_0), v7)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_parser_errposition(m, v57, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_checkSchemaNameRV_1), int32(_a_F_checkSchemaNameRV_2), int32(_a_F_checkSchemaNameRV_3))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(_a_F_checkSchemaNameRV_4), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_parser_errposition(m, v77, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_checkSchemaNameRV_1), int32(_a_F_checkSchemaNameRV_5), int32(_a_F_checkSchemaNameRV_3))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L15
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_schema_to_xmlschema(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = F_text_to_cstring(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			F_schema_to_xmlschema_internal(m, v2, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
