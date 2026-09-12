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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
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
	v165 = m.ExcPending
	if v165 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11-int32(-64)))) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+72)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v11)+124)) = l1
	v42 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+51)) = uint8(v42)
	v44 = v21 + v22
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+18)))
	if v46&int32(2016) != 0 {
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
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L49
	}
L8:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v124 = F_heap_modify_tuple(m, v19, v117, v11+int32(112), v11+int32(80), v11+int32(48))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L36
	}
L9:
	;
	v107 = F_pg_detoast_datum(m, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L34
	}
L10:
	;
	v49 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)) = uint8(v49)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+20)))
	if v51&int32(1) == v49 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v100 = F_getmissingattr(m, v45, int32(32), v11+int32(47))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L32
	}
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+516))
	if int32(0) <= v56 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21)+26)))
	if int32(0) <= v89 {
		goto L28
	} else {
		goto L29
	}
L16:
	;
	v59 = v56 + v44
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+522)))
	if v60 != int32(1) {
		v105 = v59
		goto L9
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v87 = F_nocachegetattr(m, v19, int32(32), v45)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L19:
	;
	v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+520)))
	switch v63&int32(65535) - int32(1) {
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
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v105 = v70
	goto L9
L22:
	;
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59))))
	v105 = v69
	goto L9
L23:
	;
	v68 = int32(*(*int8)(unsafe.Add(mBase, uint32(v59))))
	v105 = v68
	goto L9
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v63
	F_errmsg_internal(m, int32(475879), v11+int32(32))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(321686), int32(70), int32(67010))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
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
	v105 = v87
	goto L9
L28:
	;
	v92 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)) = uint8(v92)
	goto L8
L29:
	;
	goto L30
L30:
	;
	v95 = F_nocachegetattr(m, v19, int32(32), v45)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v105 = v95
	goto L9
L32:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)))
	if v102&int32(1) != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	v105 = v100
	goto L9
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	v110 = F_aclnewowner(m, v107, v109, l1)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v110
	v113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)) = uint8(v113)
	goto L8
L36:
	;
	F_CatalogTupleUpdate(m, v15, v124+int32(4), v124)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v44)+96))
	if v130 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_AlterTypeOwnerInternal(m, v130, l1)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+79)))
	if v133 == int32(114) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v136 = F_get_range_multirange(m, l0)
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
	F_sequence_close(m, v15, int32(3))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	if v136 == int32(0) {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	F_AlterTypeOwnerInternal(m, v136, l1)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
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
	F_errmsg_internal(m, int32(49916), v11)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(486893), int32(4002), int32(307877))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
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
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v169 = F_format_type_be(m, l0)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v169
	F_errmsg(m, int32(190813), v11+int32(16))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(486893), int32(4042), int32(307877))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
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
