package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_toast_delete_datum(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L42
	}
L2:
	;
	m.G0 = v9 - int32(-64)
	return
L3:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v14 != int32(18) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+10))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+14))
	v20 = F_table_open(m, v18, int32(3))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v27 = F_toast_open_indexes(m, v20, int32(3), v7+int32(-4), v7+int32(-56))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v30 = v7 + int32(-52)
	F_ScanKeyInit(m, v30, int32(1), int32(3), int32(184), v17)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v27<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_toast_delete_datum[0]))
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v59 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L10:
	;
	v59 = int32(1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_toast_delete_datum[1]))
	v47 = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_toast_delete_datum[2]))
	if base.B2i32(v46 == v47)|base.B2i32(v50 == v47) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v59 = base.B2i32(v50 != int32(0))
	goto L9
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v54 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v59 = int32(0)
	goto L9
L16:
	;
	v64 = F_systable_beginscan_ordered(m, v20, v40, int32(_a_F_toast_delete_datum_0), int32(1), v30)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v67 = F_systable_getnext_ordered(m, v64, int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	if v67 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v69 = v67
	goto L22
L20:
	;
	goto L21
L21:
	;
	F_systable_endscan_ordered(m, v64)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L32
	}
L22:
	;
	v76 = v69 + int32(4)
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
	v82 = F_systable_getnext_ordered(m, v64, int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L30
	}
L25:
	;
	F_heap_abort_speculative(m, v20, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	F_simple_heap_delete(m, v20, v76)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
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
	if v82 != 0 {
		v69 = v82
		goto L22
	} else {
		goto L31
	}
L31:
	;
	goto L23
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if int32(0) < v92 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = int32(0)
	goto L36
L34:
	;
	goto L35
L35:
	;
	F_pfree(m, v36)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L40
	}
L36:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v36+v96<<(uint(int32(2))%32))))
	F_relation_close(m, v105, int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	v110 = v96 + int32(1)
	if v110 != v92 {
		v96 = v110
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	F_relation_close(m, v20, int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
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
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_toast_delete_datum_2), int32(653), int32(_a_F_toast_delete_datum_3))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
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
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
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
	v78 = F_heap_form_tuple(m, l1, v11+int32(3328), v11+int32(1664))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
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
	v64 = v30 + int32(1)
	if v64 != v13 {
		v30 = v64
		goto L9
	} else {
		goto L16
	}
L12:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v30<<(uint(int32(4))%32))+24)))
	if v43 != int32(_a_F_toast_flatten_tuple_1) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v50 = v11 + int32(3328) + v30<<(uint(int32(2))%32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v52 != int32(1) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v55 = F_detoast_external_attr(m, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v55
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30+v11))) = uint8(v59)
	goto L11
L16:
	;
	goto L10
L17:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v78)+8)) = uint16(v80)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = v88
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v92)+16)) = uint16(v94)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+20)))
	v101 = v99 & int32(15)
	*(*uint16)(unsafe.Add(mBase, uint32(v98)+20)) = uint16(v101)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+20)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+20)))
	v109 = v104 | v106&int32(_a_F_toast_flatten_tuple_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v103)+20)) = uint16(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111)+18)))
	v114 = v112 & int32(_a_F_toast_flatten_tuple_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v111)+18)) = uint16(v114)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+18)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+18)))
	v122 = v117 | v119&int32(_a_F_toast_flatten_tuple_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+18)) = uint16(v122)
	if v25 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v129 = int32(0)
	goto L21
L19:
	;
	goto L20
L20:
	;
	m.G0 = v11 + int32(_a_F_toast_flatten_tuple_0)
	return v78
L21:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v11))))
	if v136 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L20
L23:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(3328)+v129<<(uint(int32(2))%32))))
	F_pfree(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v148 = v129 + int32(1)
	if v148 != v13 {
		v129 = v148
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
func F_toast_flatten_tuple_to_datum(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
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
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
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
	return int32(0)
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
		v113 = v4
		v115 = v37
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v120 = v14 + int32(3328)
	v122 = v14 + int32(1664)
	v123 = F_heap_compute_data_size(m, l2, v120, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
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
	v96 = int32(0)
	if v88&int32(1) == v96 {
		v113 = v96
		v115 = v37
		goto L6
	} else {
		goto L18
	}
L10:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v42<<(uint(int32(4))%32))+24)))
	if v62 != int32(_a_F_toast_flatten_tuple_to_datum_2) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v88 = int32(1)
	goto L12
L12:
	;
	v92 = v42 + int32(1)
	if v92 != v16 {
		v42 = v92
		v46 = v88
		goto L8
	} else {
		goto L17
	}
L13:
	;
	v88 = v46
	goto L12
L14:
	;
	v67 = int32(2)
	v69 = v14 + int32(3328) + v42<<(uint(v67)%32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if base.B2i32(v71 != int32(1))&base.B2i32(v71&int32(3) != v67) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v79 = F_detoast_attr(m, v70)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v79
	v83 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v42+v14))) = uint8(v83)
	goto L13
L17:
	;
	goto L9
L18:
	;
	v102 = base.I32_div_s(v16+int32(7), int32(8))
	v113 = int32(1)
	v115 = (v102 + int32(30)) & int32(-8)
	goto L6
L19:
	;
	v125 = v123 + v115
	v126 = F_palloc0(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)+15))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+15)) = v128
	v130 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v126))) = v130
	v132 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+22)) = uint8(v115)
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v125 << (uint(int32(2)) % 32)
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+18)))
	v141 = v138&int32(_a_F_toast_flatten_tuple_to_datum_3) | v16
	*(*uint16)(unsafe.Add(mBase, uint32(v126)+18)) = uint16(v141)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = v143
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v145
	v147 = int32(0)
	if v113 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v154 = v126 + int32(23)
	goto L23
L22:
	;
	v154 = v147
	goto L23
L23:
	;
	F_heap_fill_tuple(m, l2, v120, v122, v126+v115, v126+int32(20), v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
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
	v160 = v147
	goto L28
L26:
	;
	goto L27
L27:
	;
	m.G0 = v14 + int32(_a_F_toast_flatten_tuple_to_datum_0)
	return v126
L28:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v14))))
	if v171 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L27
L30:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(3328)+v160<<(uint(int32(2))%32))))
	F_pfree(m, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v183 = v160 + int32(1)
	if v183 != v16 {
		v160 = v183
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
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
	v22 = F_palloc(m, v18<<(uint(int32(2))%32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v52 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v52
	v55 = F_palloc(m, v52)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v22
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v25 <= int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v33 = int32(0)
	goto L9
L9:
	;
	v39 = v33 << (uint(int32(2)) % 32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39+v40)))
	v43 = F_index_open(m, v42, l1)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L11
	}
L10:
	;
	goto L1
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v45+v39))) = v43
	v49 = v33 + int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v49 < v50 {
		v33 = v49
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
	F_errfinish(m, int32(_a_F_toast_open_indexes_1), int32(609), int32(_a_F_toast_open_indexes_2))
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
