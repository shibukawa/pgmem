package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetPublicationSchemas(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v12 = F_table_open(m, int32(_a_F_GetPublicationSchemas_0), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = v6 + int32(-56)
	F_ScanKeyInit(m, v17, int32(2), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = int32(0)
	v26 = int32(1)
	v29 = F_systable_beginscan(m, v12, int32(_a_F_GetPublicationSchemas_1), v26, v24, v26, v17)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = v24
	goto L5
L5:
	;
	v36 = F_systable_getnext(m, v29)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	F_systable_endscan(m, v29)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38+v39)+8))
	v42 = F_lappend_oid(m, v31, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L6
L11:
	;
	v31 = v42
	goto L5
L12:
	;
	F_relation_close(m, v12, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	m.G0 = v8 - int32(-64)
	return v31
}
func F_PublicationDropTables(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L27
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L10
	} else {
		goto L23
	}
L3:
	;
	m.G0 = v11 + int32(16)
	return
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v20 = int32(0)
	goto L6
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v20<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v33 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	goto L3
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v36 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35)+56)))
	v37 = int64(0)
	v39 = F_GetSysCacheOid(m, int32(53), v36, base.I64_extend_i32_u(l0), v37, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v75 = v20 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v75 < v76 {
		v20 = v75
		goto L6
	} else {
		goto L22
	}
L10:
	;
	return
L11:
	;
	if v39 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if l2 != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v62 != 0 {
		goto L1
	} else {
		goto L20
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v50 + int32(4)
	F_errmsg(m, int32(_a_F_PublicationDropTables_0), v11)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_PublicationDropTables_1), int32(2103), int32(_a_F_PublicationDropTables_2))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	v63 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(_a_F_PublicationDropTables_3)
	F_performDeletion(m, v11+int32(4), int32(1), v63)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	goto L9
L22:
	;
	goto L7
L23:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	F_errmsg(m, int32(_a_F_PublicationDropTables_4), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_PublicationDropTables_1), int32(2090), int32(_a_F_PublicationDropTables_2))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(_a_F_PublicationDropTables_5), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_PublicationDropTables_1), int32(2109), int32(_a_F_PublicationDropTables_2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_publication_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14290(m, l0, l1, int32(_a_F_get_publication_oid_0), int32(3968), int32(_a_F_get_publication_oid_1), int32(_a_F_get_publication_oid_2), int32(67137668), int32(48))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
