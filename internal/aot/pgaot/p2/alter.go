package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AlterTypeOwnerInternal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v9 = m.G0
	v11 = v9 - int32(240)
	m.G0 = v11
	v15 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = F_SearchSysCacheCopy(m, int32(82), l0, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L52
	}
L4:
	;
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+72)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v11)+124)) = l1
	v40 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+51)) = uint8(v40)
	v42 = v21 + v22
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+18)))
	if v44&int32(2016) != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L49
	}
L8:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v120 = F_heap_modify_tuple(m, v19, v113, v11+int32(112), v11+int32(80), v11+int32(48))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L36
	}
L9:
	;
	v103 = F_pg_detoast_datum(m, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L34
	}
L10:
	;
	v47 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)) = uint8(v47)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+20)))
	if v49&int32(1) == v47 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v98 = F_getmissingattr(m, v43, int32(32), v11+int32(47))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L32
	}
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+516))
	if int32(0) <= v54 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v87 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21)+26)))
	if int32(0) <= v87 {
		goto L28
	} else {
		goto L29
	}
L16:
	;
	v57 = v54 + v42
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+522)))
	if v58 != int32(1) {
		v101 = v57
		goto L9
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v85 = F_nocachegetattr(m, v19, int32(32), v43)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L27
	}
L19:
	;
	v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+520)))
	switch v61&int32(_a_F_AlterTypeOwnerInternal_3) - int32(1) {
	case 0:
		goto L23
	case 1:
		goto L22
	default:
		goto L20
	case 3:
		goto L21
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v101 = v68
	goto L9
L22:
	;
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57))))
	v101 = v67
	goto L9
L23:
	;
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57))))
	v101 = v66
	goto L9
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v61
	F_errmsg_internal(m, int32(_a_F_AlterTypeOwnerInternal_4), v11+int32(32))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_AlterTypeOwnerInternal_5), int32(70), int32(_a_F_AlterTypeOwnerInternal_6))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
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
	v101 = v85
	goto L9
L28:
	;
	v90 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)) = uint8(v90)
	goto L8
L29:
	;
	goto L30
L30:
	;
	v93 = F_nocachegetattr(m, v19, int32(32), v43)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v101 = v93
	goto L9
L32:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)))
	if v100 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	v101 = v98
	goto L9
L34:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v42)+72))
	v106 = F_aclnewowner(m, v103, v105, l1)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v106
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)) = uint8(v109)
	goto L8
L36:
	;
	F_CatalogTupleUpdate(m, v15, v120+int32(4), v120)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	if v126 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_AlterTypeOwnerInternal(m, v126, l1)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+79)))
	if v129 == int32(114) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v132 = F_get_range_multirange(m, l0)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_relation_close(m, v15, int32(3))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	if v132 == int32(0) {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	F_AlterTypeOwnerInternal(m, v132, l1)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	m.G0 = v11 + int32(240)
	return
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(_a_F_AlterTypeOwnerInternal_7), v11)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_AlterTypeOwnerInternal_1), int32(4002), int32(_a_F_AlterTypeOwnerInternal_2))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v165 = F_format_type_be(m, l0)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v165
	F_errmsg(m, int32(_a_F_AlterTypeOwnerInternal_0), v11+int32(16))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_AlterTypeOwnerInternal_1), int32(4042), int32(_a_F_AlterTypeOwnerInternal_2))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
