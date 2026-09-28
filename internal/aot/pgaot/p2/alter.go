package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AlterTypeOwnerInternal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	v10 = m.G0
	v12 = v10 - int32(368)
	m.G0 = v12
	v16 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = F_SearchSysCacheCopy(m, int32(82), base.I64_extend_i32_u(l0), int64(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L56
	}
L4:
	;
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
	v25 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v12)+104)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v12)+136)) = base.I64_extend_i32_u(l1)
	v43 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+51)) = uint8(v43)
	v45 = v24 + v23
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+18)))
	if v47&int32(2016) != 0 {
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
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L53
	}
L8:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	v129 = F_heap_modify_tuple(m, v21, v122, v12+int32(112), v12+int32(80), v12+int32(48))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L40
	}
L9:
	;
	v110 = F_pg_detoast_datum(m, base.I32_wrap_i64(v108))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L38
	}
L10:
	;
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+47)) = uint8(v50)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+20)))
	if v52&int32(1) == v50 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v103 = F_getmissingattr(m, v46, int32(32), v12+int32(47))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L36
	}
L13:
	;
	v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+276)))
	if int32(0) <= v57 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v92 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+26)))
	if int32(0) <= v92 {
		goto L32
	} else {
		goto L33
	}
L16:
	;
	v60 = v57 + v45
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+280)))
	if v61 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v90 = F_nocachegetattr(m, v21, int32(32), v46)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L31
	}
L19:
	;
	v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+278)))
	if base.I32_popcnt(v64) != int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v108 = base.I64_extend_i32_u(v60)
	goto L9
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L28
	}
L23:
	;
	switch base.I32_ctz(v64) {
	case 0:
		goto L27
	case 1:
		goto L26
	case 2:
		goto L25
	case 3:
		goto L24
	default:
		goto L22
	}
L24:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	v108 = v72
	goto L9
L25:
	;
	v71 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60))))
	v108 = v71
	goto L9
L26:
	;
	v70 = int64(*(*int16)(unsafe.Add(mBase, uint32(v60))))
	v108 = v70
	goto L9
L27:
	;
	v69 = int64(*(*int8)(unsafe.Add(mBase, uint32(v60))))
	v108 = v69
	goto L9
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v64
	F_errmsg_internal(m, int32(_a_F_AlterTypeOwnerInternal_0), v12+int32(32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_AlterTypeOwnerInternal_1), int32(123), int32(_a_F_AlterTypeOwnerInternal_2))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v108 = v90
	goto L9
L32:
	;
	v95 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+47)) = uint8(v95)
	goto L8
L33:
	;
	goto L34
L34:
	;
	v98 = F_nocachegetattr(m, v21, int32(32), v46)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v108 = v98
	goto L9
L36:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+47)))
	if v105 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v108 = v103
	goto L9
L38:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v45)+72))
	v113 = F_aclnewowner(m, v110, v112, l1)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+79)) = uint8(v115)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+360)) = base.I64_extend_i32_u(v113)
	goto L8
L40:
	;
	F_CatalogTupleUpdate(m, v16, v129+int32(4), v129)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v45)+96))
	if v135 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_AlterTypeOwnerInternal(m, v135, l1)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+79)))
	if v138 == int32(114) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	v141 = F_get_range_multirange(m, l0)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_relation_close(m, v16, int32(3))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L52
	}
L49:
	;
	if v141 == int32(0) {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	F_AlterTypeOwnerInternal(m, v141, l1)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	m.G0 = v12 + int32(368)
	return
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg_internal(m, int32(_a_F_AlterTypeOwnerInternal_7), v12)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_AlterTypeOwnerInternal_4), int32(4070), int32(_a_F_AlterTypeOwnerInternal_6))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v174 = F_format_type_be(m, l0)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v174
	F_errmsg(m, int32(_a_F_AlterTypeOwnerInternal_3), v12+int32(16))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_AlterTypeOwnerInternal_4), int32(_a_F_AlterTypeOwnerInternal_5), int32(_a_F_AlterTypeOwnerInternal_6))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
