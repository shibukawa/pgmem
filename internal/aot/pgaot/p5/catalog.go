package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatalogCacheInitializeCache(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v63 int32
	_ = v63
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v22 = F_table_open(m, v20, int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = int32(4480304)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v28 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v31 = F_CreateTupleDescCopyConstr(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v36 = F_pstrdup(m, v33+int32(4))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v36
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+117)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)) = uint8(v40)
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v25
	F_sequence_close(m, v22, int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if int32(0) < v47 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v63 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v31
	m.G0 = v18 + int32(16)
	return
L9:
	;
	v77 = v63 << (uint(int32(2)) % 32)
	v78 = l0 + int32(48) + v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77+(l0+int32(16))))) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v77+(l0+int32(32))))) = v170
	v178 = l0 + int32(104) + v63*int32(48)
	v182 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	F_fmgr_info_cxt(m, v168, v178+int32(16), v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L42
	}
L12:
	;
	v168 = v165
	v170 = int32(1584)
	v171 = int32(1585)
	goto L11
L13:
	;
	v165 = int32(184)
	goto L12
L14:
	;
	if int32(0) <= v79 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v31-int32(12)+v97<<(uint(int32(4))%32)+v79*int32(100))))
	if v104 <= int32(2201) {
		goto L28
	} else {
		goto L29
	}
L17:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errmsg_internal(m, int32(167843), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(494158), int32(1178), int32(396123))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
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
	v168 = int32(61)
	v170 = v108
	v171 = v109
	goto L11
L22:
	;
	v168 = int32(679)
	v170 = int32(1582)
	v171 = int32(1583)
	goto L11
L23:
	;
	v168 = int32(67)
	v170 = int32(1580)
	v171 = int32(1581)
	goto L11
L24:
	;
	v165 = int32(65)
	goto L12
L25:
	;
	v168 = int32(63)
	v170 = int32(1578)
	v171 = int32(1579)
	goto L11
L26:
	;
	v168 = int32(62)
	v170 = int32(1576)
	v171 = int32(1577)
	goto L11
L27:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L39
	}
L28:
	;
	v108 = int32(1574)
	v109 = int32(1575)
	switch v104 - int32(16) {
	case 0:
		v168 = int32(60)
		v170 = v108
		v171 = v109
		goto L11
	default:
		goto L27
	case 2:
		goto L21
	case 3:
		goto L26
	case 5:
		goto L25
	case 7:
		goto L24
	case 8, 10:
		goto L13
	case 9:
		goto L23
	case 14:
		goto L22
	}
L29:
	;
	goto L30
L30:
	;
	if v104 <= int32(3768) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if base.Ui32(v104-int32(2202)) < base.Ui32(int32(5)) {
		goto L13
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	switch v104 - int32(4089) {
	case 0, 7:
		goto L13
	case 1, 2, 3, 4, 5, 6:
		goto L27
	default:
		goto L36
	}
L34:
	;
	if v104 == int32(3734) {
		goto L13
	} else {
		goto L35
	}
L35:
	;
	goto L27
L36:
	;
	if v104 == int32(3769) {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	if v104 == int32(4191) {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	goto L27
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v104
	F_errmsg_internal(m, int32(22312), v18)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(494158), int32(330), int32(172443))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
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
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*int64)(unsafe.Add(mBase, uint32(v178)+8)) = int64(4080218931200)
	v188 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v178)+6)) = uint16(v188)
	*(*uint16)(unsafe.Add(mBase, uint32(v178)+4)) = uint16(v185)
	v192 = v63 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v192 < v193 {
		v63 = v192
		goto L9
	} else {
		goto L43
	}
L43:
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
	return base.B2i32(base.Ui32(l0) < base.Ui32(int32(12000)))
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
