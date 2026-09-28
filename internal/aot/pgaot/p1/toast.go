package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_toast_delete_datum(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = base.I32_wrap_i64(l0)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L42
	}
L2:
	;
	m.G0 = v10 + int32(80)
	return
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v16 != int32(18) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+10)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+14))
	v22 = F_table_open(m, v20, int32(3))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v29 = F_toast_open_indexes(m, v22, int32(3), v10+int32(76), v10+int32(12))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v32 = v10 + int32(16)
	F_ScanKeyInit(m, v32, int32(1), int32(3), int32(184), v19)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v29<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_toast_delete_datum[0]))
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v61 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L10:
	;
	v61 = int32(1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_toast_delete_datum[1]))
	v49 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_toast_delete_datum[2]))
	if base.B2i32(v48 == v49)|base.B2i32(v52 == v49) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v61 = base.B2i32(v52 != int32(0))
	goto L9
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v56 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v61 = int32(0)
	goto L9
L16:
	;
	v66 = F_systable_beginscan_ordered(m, v22, v42, int32(_a_F_toast_delete_datum_0), int32(1), v32)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v69 = F_systable_getnext_ordered(m, v66, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	if v69 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v73 = v69
	goto L22
L20:
	;
	goto L21
L21:
	;
	F_systable_endscan_ordered(m, v66)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L32
	}
L22:
	;
	v79 = v73 + int32(4)
	if l1 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L21
L24:
	;
	v85 = F_systable_getnext_ordered(m, v66, int32(1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L30
	}
L25:
	;
	F_heap_abort_speculative(m, v22, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	F_simple_heap_delete(m, v22, v79)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L29
	}
L28:
	;
	goto L24
L29:
	;
	goto L24
L30:
	;
	if v85 != 0 {
		v73 = v85
		goto L22
	} else {
		goto L31
	}
L31:
	;
	goto L23
L32:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if int32(0) < v96 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v102 = int32(0)
	goto L36
L34:
	;
	goto L35
L35:
	;
	F_pfree(m, v38)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L40
	}
L36:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v38+v102<<(uint(int32(2))%32))))
	F_relation_close(m, v110, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	v115 = v102 + int32(1)
	if v115 != v96 {
		v102 = v115
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	F_relation_close(m, v22, int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	goto L2
L42:
	;
	F_errmsg_internal(m, int32(_a_F_toast_delete_datum_1), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_toast_delete_datum_2), int32(644), int32(_a_F_toast_delete_datum_3))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_toast_flatten_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	v9 = m.G0
	v11 = v9 - int32(_a_F_toast_flatten_tuple_0)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_heap_deform_tuple(m, l0, l1, v11+int32(3328), v11+int32(1664))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	base.MemoryFill(m, v11, int32(0), v13)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v24 = int32(0)
	v25 = base.B2i32(v13 <= v24)
	if v25 == v24 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v30 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v77 = F_heap_form_tuple(m, l1, v11+int32(3328), v11+int32(1664))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L17
	}
L9:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(1664)+v30))))
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v63 = v30 + int32(1)
	if v63 != v13 {
		v30 = v63
		goto L9
	} else {
		goto L16
	}
L12:
	;
	v41 = v30 << (uint(int32(3)) % 32)
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v41)+30)))
	if v43 != int32(_a_F_toast_flatten_tuple_1) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v48 = v11 + int32(3328) + v41
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v50 != int32(1) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v53 = F_detoast_external_attr(m, v49)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = base.I64_extend_i32_u(v53)
	v58 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30+v11))) = uint8(v58)
	goto L11
L16:
	;
	goto L10
L17:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v77)+8)) = uint16(v79)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = v87
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
	*(*int64)(unsafe.Add(mBase, uint32(v85))) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+16)) = uint16(v93)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+20)))
	v100 = v98 & int32(15)
	*(*uint16)(unsafe.Add(mBase, uint32(v97)+20)) = uint16(v100)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+20)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+20)))
	v108 = v103 | v105&int32(_a_F_toast_flatten_tuple_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v102)+20)) = uint16(v108)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+18)))
	v113 = v111 & int32(_a_F_toast_flatten_tuple_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v110)+18)) = uint16(v113)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+18)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+18)))
	v121 = v116 | v118&int32(_a_F_toast_flatten_tuple_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v115)+18)) = uint16(v121)
	if v25 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v128 = int32(0)
	goto L21
L19:
	;
	goto L20
L20:
	;
	m.G0 = v11 + int32(_a_F_toast_flatten_tuple_0)
	return v77
L21:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128+v11))))
	if v135 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L20
L23:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(3328)+v128<<(uint(int32(3))%32))))
	F_pfree(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v147 = v128 + int32(1)
	if v147 != v13 {
		v128 = v147
		goto L21
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	goto L22
}
func F_toast_flatten_tuple_to_datum(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(_a_F_toast_flatten_tuple_to_datum_0)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_toast_flatten_tuple_to_datum[0]))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_toast_flatten_tuple_to_datum[1]))) = v4
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_toast_flatten_tuple_to_datum[2]))) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_toast_flatten_tuple_to_datum[3]))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_toast_flatten_tuple_to_datum[4]))) = l1
	F_heap_deform_tuple(m, v14+int32(_a_F_toast_flatten_tuple_to_datum_1), l2, v14+int32(3328), v14+int32(1664))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	base.MemoryFill(m, v14, int32(0), v16)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v37 = int32(24)
	v39 = base.B2i32(v16 <= int32(0))
	if v16 <= int32(0) {
		v112 = v4
		v114 = v37
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v119 = v14 + int32(3328)
	v121 = v14 + int32(1664)
	v122 = F_heap_compute_data_size(m, l2, v119, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L19
	}
L7:
	;
	v42 = int32(0)
	v46 = v4
	goto L8
L8:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(1664)+v42))))
	if v56 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v95 = int32(0)
	if v87&int32(1) == v95 {
		v112 = v95
		v114 = v37
		goto L6
	} else {
		goto L18
	}
L10:
	;
	v60 = v42 << (uint(int32(3)) % 32)
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v60)+30)))
	if v62 != int32(_a_F_toast_flatten_tuple_to_datum_2) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v87 = int32(1)
	goto L12
L12:
	;
	v91 = v42 + int32(1)
	if v91 != v16 {
		v42 = v91
		v46 = v87
		goto L8
	} else {
		goto L17
	}
L13:
	;
	v87 = v46
	goto L12
L14:
	;
	v67 = v14 + int32(3328) + v60
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if base.B2i32(v69 != int32(1))&base.B2i32(v69&int32(3) != int32(2)) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v77 = F_detoast_attr(m, v68)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = base.I64_extend_i32_u(v77)
	v82 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v42+v14))) = uint8(v82)
	goto L13
L17:
	;
	goto L9
L18:
	;
	v101 = base.I32_div_s(v16+int32(7), int32(8))
	v112 = int32(1)
	v114 = (v101 + int32(30)) & int32(-8)
	goto L6
L19:
	;
	v124 = v122 + v114
	v125 = F_palloc0(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v127 = *(*int64)(unsafe.Add(mBase, uint32(l0)+15))
	*(*int64)(unsafe.Add(mBase, uint32(v125)+15)) = v127
	v129 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v125))) = v129
	v131 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v125)+8)) = v131
	*(*uint8)(unsafe.Add(mBase, uint32(v125)+22)) = uint8(v114)
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v124 << (uint(int32(2)) % 32)
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+18)))
	v140 = v137&int32(_a_F_toast_flatten_tuple_to_datum_3) | v16
	*(*uint16)(unsafe.Add(mBase, uint32(v125)+18)) = uint16(v140)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+8)) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v144
	v146 = int32(0)
	if v112 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v153 = v125 + int32(23)
	goto L23
L22:
	;
	v153 = v146
	goto L23
L23:
	;
	F_heap_fill_tuple(m, l2, v119, v121, v125+v114, v125+int32(20), v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v39 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v159 = v146
	goto L28
L26:
	;
	goto L27
L27:
	;
	m.G0 = v14 + int32(_a_F_toast_flatten_tuple_to_datum_0)
	return base.I64_extend_i32_u(v125)
L28:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+v14))))
	if v170 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L27
L30:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(3328)+v159<<(uint(int32(3))%32))))
	F_pfree(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v182 = v159 + int32(1)
	if v182 != v16 {
		v159 = v182
		goto L28
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	goto L29
}
func F_toast_open_indexes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
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
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if int32(0) < v67 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	return int32(0)
L3:
	;
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v18
	v21 = F_palloc_mul(m, int32(4), v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v51
	v55 = F_palloc_mul(m, int32(4), v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v21
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v24 <= int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v32 = int32(0)
	goto L9
L9:
	;
	v38 = v32 << (uint(int32(2)) % 32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38+v39)))
	v42 = F_index_open(m, v41, l1)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L11
	}
L10:
	;
	goto L1
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v44+v38))) = v42
	v48 = v32 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v48 < v49 {
		v32 = v48
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v55
	goto L1
L14:
	;
	F_list_free(m, v14)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L26
	}
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v76 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	F_list_free(m, v14)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L22
	}
L18:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v70+v76<<(uint(int32(2))%32))))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+192))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+18)))
	if v86 != 0 {
		goto L14
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	v88 = v76 + int32(1)
	if v88 != v67 {
		v76 = v88
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v105
	F_errmsg_internal(m, int32(_a_F_toast_open_indexes_0), v12)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_toast_open_indexes_1), int32(600), int32(_a_F_toast_open_indexes_2))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	m.G0 = v12 + int32(16)
	return v76
}
