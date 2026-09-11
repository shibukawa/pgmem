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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v14 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = int32(3)
	F_ScanKeyInit(m, v10, v18, v18, int32(184), l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = F_table_beginscan_catalog(m, v14, int32(1), v10)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = F_heap_getnext(m, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = v26
	v30 = v3
	goto L9
L7:
	;
	v69 = v3
	goto L8
L8:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+188))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	m.T0[v76].(func(*base.Module, int32))(m, v24)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L25
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	v37 = v35 + v36
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+119)))
	switch v38 - int32(112) {
	case 0, 2:
		goto L12
	default:
		v63 = v30
		goto L11
	}
L10:
	;
	v69 = v63
	goto L8
L11:
	;
	v65 = F_heap_getnext(m, v24)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L23
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	goto L13
L13:
	;
	if base.Ui32(v41) < base.Ui32(int32(12000)) {
		v63 = v30
		goto L11
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(v41) < base.Ui32(int32(16384)) {
		v63 = v30
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+118)))
	if v46 != int32(112) {
		v63 = v30
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v49 = F_get_rel_relkind(m, v41)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v59 = F_GetPubPartitionOptionRelations(m, int32(0), l1, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L21
	}
L18:
	;
	v55 = F_lappend_oid(m, v30, v41)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	switch v49&int32(255) - int32(112) {
	case 0:
		goto L17
	default:
		v63 = v30
		goto L11
	case 2:
		goto L18
	}
L20:
	;
	v63 = v55
	goto L11
L21:
	;
	v61 = F_list_concat_unique_oid(m, v30, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v63 = v61
	goto L11
L23:
	;
	if v65 != 0 {
		v28 = v65
		v30 = v63
		goto L9
	} else {
		goto L24
	}
L24:
	;
	goto L10
L25:
	;
	F_sequence_close(m, v14, int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	m.G0 = v10 + int32(48)
	return v69
}
func F__equalCreateSchemaStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v55
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v40 = F_equal(m, v38, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L17
	} else {
		goto L18
	}
L3:
	;
	if v6 == int32(0) {
		v55 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v6 != v7 {
		v55 = v3
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v13 == int32(0) {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v32 == int32(0) {
		goto L2
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v12 != v13 {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v17 = v7
	v18 = v6
	goto L11
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v21
		v33 = v22
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v32 = v21
	v33 = v22
	goto L8
L13:
	;
	v25 = int32(1)
	if v21 == v22 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v55 = v3
	goto L1
L16:
	;
	goto L2
L17:
	;
	return int32(0)
L18:
	;
	if v40 == int32(0) {
		v55 = v3
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v48 = F_equal(m, v46, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	if v48 == int32(0) {
		v55 = v3
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v55 = base.B2i32(v52 == v53)
	goto L1
}
func F_schema_to_xmlschema(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_text_to_cstring(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_schema_to_xmlschema_internal(m, v2, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
