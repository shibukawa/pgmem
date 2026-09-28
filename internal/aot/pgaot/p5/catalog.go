package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatalogCacheInitializeCache(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v21 = F_table_open(m, v19, int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v23 = int32(_a_F_CatalogCacheInitializeCache_0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheInitializeCache[0]))
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheInitializeCache[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheInitializeCache[0])) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	v30 = F_CreateTupleDescCopyConstr(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v35 = F_pstrdup(m, v32+int32(4))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+117)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)) = uint8(v39)
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheInitializeCache[0])) = v24
	F_relation_close(m, v21, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if int32(0) < v46 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v62 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v30
	m.G0 = v17 + int32(16)
	return
L9:
	;
	v72 = v62 << (uint(int32(2)) % 32)
	v73 = l0 + int32(48) + v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v74 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72+(l0+int32(16))))) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v72+(l0+int32(32))))) = v173
	v181 = l0 + int32(104) + v62*int32(56)
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheInitializeCache[1]))
	F_fmgr_info_cxt(m, v172, v181+int32(16), v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L44
	}
L12:
	;
	v172 = v168
	v173 = int32(1781)
	v174 = int32(1782)
	goto L11
L13:
	;
	v168 = int32(184)
	goto L12
L14:
	;
	if int32(0) <= v74 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v30+v92<<(uint(int32(3))%32)+v74*int32(100)-int32(4))))
	if v101 <= int32(2201) {
		goto L29
	} else {
		goto L30
	}
L17:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errmsg_internal(m, int32(_a_F_CatalogCacheInitializeCache_1), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_CatalogCacheInitializeCache_2), int32(1210), int32(_a_F_CatalogCacheInitializeCache_3))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
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
	v172 = int32(61)
	v173 = v105
	v174 = v106
	goto L11
L22:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L41
	}
L23:
	;
	if v101 == int32(3769) {
		goto L13
	} else {
		goto L40
	}
L24:
	;
	v172 = int32(679)
	v173 = int32(1779)
	v174 = int32(1780)
	goto L11
L25:
	;
	v172 = int32(67)
	v173 = int32(1777)
	v174 = int32(1778)
	goto L11
L26:
	;
	v168 = int32(65)
	goto L12
L27:
	;
	v172 = int32(63)
	v173 = int32(1775)
	v174 = int32(1776)
	goto L11
L28:
	;
	v172 = int32(62)
	v173 = int32(1773)
	v174 = int32(1774)
	goto L11
L29:
	;
	v105 = int32(1771)
	v106 = int32(1772)
	switch v101 - int32(16) {
	case 0:
		v172 = int32(60)
		v173 = v105
		v174 = v106
		goto L11
	default:
		goto L22
	case 2:
		goto L21
	case 3:
		goto L28
	case 5:
		goto L27
	case 7:
		goto L26
	case 8, 10:
		goto L13
	case 9:
		goto L25
	case 14:
		goto L24
	}
L30:
	;
	goto L31
L31:
	;
	if v101 <= int32(3768) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if base.B2i32(v101 == int32(3734))|base.B2i32(base.Ui32(v101-int32(2202)) < base.Ui32(int32(5))) != 0 {
		goto L13
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v101 <= int32(_a_F_CatalogCacheInitializeCache_4) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L22
L36:
	;
	switch v101 - int32(4089) {
	case 0, 7:
		goto L13
	case 1, 2, 3, 4, 5, 6:
		goto L22
	default:
		goto L23
	}
L37:
	;
	goto L38
L38:
	;
	if base.B2i32(v101 == int32(_a_F_CatalogCacheInitializeCache_5))|base.B2i32(v101 == int32(_a_F_CatalogCacheInitializeCache_6)) != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	goto L22
L40:
	;
	goto L22
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v101
	F_errmsg_internal(m, int32(_a_F_CatalogCacheInitializeCache_7), v17)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_CatalogCacheInitializeCache_2), int32(342), int32(_a_F_CatalogCacheInitializeCache_8))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	*(*int64)(unsafe.Add(mBase, uint32(v181)+8)) = int64(4080218931200)
	v191 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v181)+6)) = uint16(v191)
	*(*uint16)(unsafe.Add(mBase, uint32(v181)+4)) = uint16(v188)
	v195 = v62 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v195 < v196 {
		v62 = v195
		goto L9
	} else {
		goto L45
	}
L45:
	;
	goto L10
}
func F_CatalogCloseIndexes(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	F_ExecCloseIndices(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		F_pfree(m, l0)
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			return
		}
	}
}
func F_IsCatalogRelationOid(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0) < base.Ui32(int32(_a_F_IsCatalogRelationOid_0)))
}
func F_get_catalog_object_by_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_get_catalog_object_by_oid_extended(m, l0, l1, l2, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
