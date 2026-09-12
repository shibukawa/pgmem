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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
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
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L43
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
	F_ScanKeyInit(m, v7+int32(-52), int32(1), int32(3), int32(184), v17)
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
	v43 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v58 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L10:
	;
	v58 = int32(1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	v48 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v48 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v58 = base.B2i32(v46 != int32(0))
	goto L9
L14:
	;
	if v46 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v53 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v58 = int32(0)
	goto L9
L17:
	;
	v65 = F_systable_beginscan_ordered(m, v20, v40, int32(4212056), int32(1), v7+int32(-52))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v68 = F_systable_getnext_ordered(m, v65, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	if v68 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v70 = v68
	goto L23
L21:
	;
	goto L22
L22:
	;
	F_systable_endscan_ordered(m, v65)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L33
	}
L23:
	;
	v77 = v70 + int32(4)
	if l1 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L22
L25:
	;
	v83 = F_systable_getnext_ordered(m, v65, int32(1))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L31
	}
L26:
	;
	F_heap_abort_speculative(m, v20, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_simple_heap_delete(m, v20, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L30
	}
L29:
	;
	goto L25
L30:
	;
	goto L25
L31:
	;
	if v83 != 0 {
		v70 = v83
		goto L23
	} else {
		goto L32
	}
L32:
	;
	goto L24
L33:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if int32(0) < v93 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v97 = int32(0)
	goto L37
L35:
	;
	goto L36
L36:
	;
	F_pfree(m, v36)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L41
	}
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v36+v97<<(uint(int32(2))%32))))
	F_relation_close(m, v106, int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L39
	}
L38:
	;
	goto L36
L39:
	;
	v111 = v97 + int32(1)
	if v111 != v93 {
		v97 = v111
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	F_sequence_close(m, v20, int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	goto L2
L43:
	;
	F_errmsg_internal(m, int32(93062), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(516845), int32(653), int32(92621))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_toast_flatten_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	v11 = m.G0
	v13 = v11 - int32(9984)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_heap_deform_tuple(m, l0, l1, v13+int32(3328), v13+int32(1664))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = F__emscripten_memset_bulkmem(m, v13, base.I32_extend8_s(int32(0)), v15)
	mBase = m.M
	goto L3
L3:
	;
	v27 = int32(0)
	v28 = base.B2i32(v15 <= v27)
	if v28 == v27 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v36 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v88 = F_heap_form_tuple(m, l1, v26+int32(3328), v26+int32(1664))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L15
	}
L7:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+int32(1664)+v36))))
	if v47 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v72 = v36 + int32(1)
	if v72 != v15 {
		v36 = v72
		goto L7
	} else {
		goto L14
	}
L10:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(24)+v36<<(uint(int32(4))%32)))))
	if v51 != int32(65535) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v58 = v26 + int32(3328) + v36<<(uint(int32(2))%32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v60 != int32(1) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v63 = F_detoast_external_attr(m, v59)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v63
	v67 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36+v26))) = uint8(v67)
	goto L9
L14:
	;
	goto L8
L15:
	;
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)) = uint16(v90)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v102)+12)) = v104
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v102)+16)) = uint16(v106)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+20)))
	v111 = v109 & int32(15)
	*(*uint16)(unsafe.Add(mBase, uint32(v108)+20)) = uint16(v111)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113)+20)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+20)))
	v119 = v114 | v116&int32(65520)
	*(*uint16)(unsafe.Add(mBase, uint32(v113)+20)) = uint16(v119)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+18)))
	v124 = v122 & int32(8191)
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+18)) = uint16(v124)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+18)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128)+18)))
	v132 = v127 | v129&int32(57344)
	*(*uint16)(unsafe.Add(mBase, uint32(v126)+18)) = uint16(v132)
	if v28 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v143 = int32(0)
	goto L19
L17:
	;
	goto L18
L18:
	;
	m.G0 = v26 + int32(9984)
	return v88
L19:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v143))))
	if v147 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(3328)+v143<<(uint(int32(2))%32))))
	F_pfree(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v159 = v143 + int32(1)
	if v159 != v15 {
		v143 = v159
		goto L19
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	goto L20
}
func F_toast_flatten_tuple_to_datum(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(10016)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[88]))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[89]))) = v4
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[90]))) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[91]))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[92]))) = l1
	F_heap_deform_tuple(m, v15+int32(9996), l2, v15+int32(3328), v15+int32(1664))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v38 = F__emscripten_memset_bulkmem(m, v15, base.I32_extend8_s(int32(0)), v17)
	mBase = m.M
	goto L3
L3:
	;
	v40 = base.B2i32(v17 <= int32(0))
	if v17 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v142 = F_heap_compute_data_size(m, l2, v38+int32(3328), v38+int32(1664))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L18
	}
L5:
	;
	v132 = int32(0)
	v137 = int32(24)
	goto L4
L6:
	;
	v47 = int32(0)
	v48 = v4
	goto L7
L7:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(1664)+v47))))
	if v60 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v98 = int32(1)
	if v92&v98 == int32(0) {
		goto L5
	} else {
		goto L17
	}
L9:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(24)+v47<<(uint(int32(4))%32)))))
	if v66 != int32(65535) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v92 = int32(1)
	goto L11
L11:
	;
	v96 = v47 + int32(1)
	if v96 != v17 {
		v47 = v96
		v48 = v92
		goto L7
	} else {
		goto L16
	}
L12:
	;
	v92 = v48
	goto L11
L13:
	;
	v71 = int32(2)
	v73 = v38 + int32(3328) + v47<<(uint(v71)%32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if base.B2i32(v75 != int32(1))&base.B2i32(v75&int32(3) != v71) != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v83 = F_detoast_attr(m, v74)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v83
	v87 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v38+v47))) = uint8(v87)
	goto L12
L16:
	;
	goto L8
L17:
	;
	v106 = base.I32_div_s(v17+int32(7), int32(8))
	v132 = v98
	v137 = (v106 + int32(30)) & int32(-8)
	goto L4
L18:
	;
	v144 = v137 + v142
	v145 = F_palloc0(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v147 = *(*int64)(unsafe.Add(mBase, uint32(l0)+15))
	*(*int64)(unsafe.Add(mBase, uint32(v145)+15)) = v147
	v149 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v145))) = v149
	v152 = v145 + int32(8)
	v153 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v152))) = v153
	*(*uint8)(unsafe.Add(mBase, uint32(v145)+22)) = uint8(v137)
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v144 << (uint(int32(2)) % 32)
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+18)))
	v162 = v159&int32(63488) | v17
	*(*uint16)(unsafe.Add(mBase, uint32(v145)+18)) = uint16(v162)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = v166
	if v132 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v178 = v145 + int32(23)
	goto L22
L21:
	;
	v178 = int32(0)
	goto L22
L22:
	;
	F_heap_fill_tuple(m, l2, v38+int32(3328), v38+int32(1664), v145+v137, v145+int32(20), v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v40 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v189 = v4
	goto L27
L25:
	;
	goto L26
L26:
	;
	m.G0 = v38 + int32(10016)
	return v145
L27:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v189))))
	if v196 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L26
L29:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(3328)+v189<<(uint(int32(2))%32))))
	F_pfree(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v208 = v189 + int32(1)
	if v208 != v17 {
		v189 = v208
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	goto L28
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = F_RelationGetIndexList(m, l0)
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
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v20 = v18
	goto L5
L4:
	;
	v20 = int32(0)
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v20
	v24 = F_palloc(m, v20<<(uint(int32(2))%32))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
	if v14 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if int32(0) < v65 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v29 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v30 <= v29 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v37 = v29
	goto L10
L10:
	;
	v43 = v37 << (uint(int32(2)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43+v44)))
	v47 = F_index_open(m, v46, l1)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L7
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v49+v43))) = v47
	v53 = v37 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v53 < v54 {
		v37 = v53
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	F_list_free(m, v14)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L26
	}
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v74 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	F_list_free(m, v14)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L22
	}
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v68+v74<<(uint(int32(2))%32))))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+192))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+18)))
	if v84 != 0 {
		goto L14
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	v86 = v74 + int32(1)
	if v86 != v65 {
		v74 = v86
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
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
	F_errmsg_internal(m, int32(60179), v12)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(516845), int32(609), int32(166868))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
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
	return v74
}
