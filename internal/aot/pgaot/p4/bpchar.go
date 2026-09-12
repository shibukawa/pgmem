package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bpchar(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v17 < int32(4) {
		v127 = v13
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v10 + int32(16)
	return v127
L4:
	;
	v23 = v17 - int32(4)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v24 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(36))))
	v56 = int32(1)
	if v24&v56 != 0 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v27 = int32(4)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v29&int32(254) == int32(2) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v42 = int32(1)
	if v24&v42 != 0 {
		v54 = int32(base.Ui32(v24)>>(uint(v42)%32)) - v42
		goto L5
	} else {
		goto L15
	}
L9:
	;
	v38 = v27
	goto L11
L10:
	;
	v38 = base.B2i32(v29 == int32(18)) << (uint(v27) % 32)
	goto L11
L11:
	;
	if v29 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = v27
	goto L14
L13:
	;
	v41 = v38
	goto L14
L14:
	;
	v54 = v41
	goto L5
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L5
L16:
	;
	v60 = v56
	goto L18
L17:
	;
	v60 = int32(4)
	goto L18
L18:
	;
	v61 = v13 + v60
	v62 = F_pg_mbstrlen_with_len(m, v61, v54)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v62 == v23 {
		v127 = v13
		goto L3
	} else {
		goto L20
	}
L20:
	;
	if v23 < v62 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v109 = v102 + int32(4)
	v110 = F_palloc(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L41
	}
L22:
	;
	v66 = F_pg_mbcharcliplen(m, v61, v54, v23)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v101 = v54
	v102 = v54 + v23 - v62
	goto L21
L25:
	;
	if v55 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v101 = v66
	v102 = v66
	goto L21
L27:
	;
	goto L28
L28:
	;
	if v54 <= v66 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v101 = v66
	v102 = v66
	goto L21
L30:
	;
	goto L31
L31:
	;
	v71 = v66
	goto L33
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L37
	}
L33:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v61))))
	if v77 != int32(32) {
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v101 = v66
	v102 = v66
	goto L21
L35:
	;
	v81 = v71 + int32(1)
	if v81 != v54 {
		v71 = v81
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v23
	F_errmsg(m, int32(657071), v10)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(490154), int32(313), int32(227802))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v109 << (uint(int32(2)) % 32)
	v116 = v110 + int32(4)
	if v101 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v102 <= v101 {
		v127 = v110
		goto L3
	} else {
		goto L46
	}
L43:
	;
	v117 = F__emscripten_memcpy_bulkmem(m, v116, v61, v101)
	mBase = m.M
	v118 = v117
	goto L45
L44:
	;
	v118 = v116
	goto L45
L45:
	;
	goto L42
L46:
	;
	v124 = F__emscripten_memset_bulkmem(m, v101+v118, base.I32_extend8_s(int32(32)), v102-v101)
	mBase = m.M
	goto L47
L47:
	;
	v127 = v110
	goto L3
}
func F_bpchar_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v79 = F_palloc0(m, int32(64))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L23
	}
L2:
	;
	v61 = v57
	goto L19
L3:
	;
	if v50 <= int32(0) {
		v74 = v50
		v77 = v52
		goto L1
	} else {
		goto L18
	}
L4:
	;
	return int32(0)
L5:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v10 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v13 = int32(1)
	v14 = v6 + v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	if base.Ui32((v16-v13)&int32(255)) < base.Ui32(int32(3)) {
		v57 = int32(4)
		v59 = v6
		v60 = v14
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v27 = int32(1)
	if v10&v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v50 = base.B2i32(v16 == int32(18)) << (uint(int32(4)) % 32)
	v52 = v14
	goto L3
L10:
	;
	v31 = v27
	goto L12
L11:
	;
	v31 = int32(4)
	goto L12
L12:
	;
	v32 = v6 + v31
	if v10&int32(1) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v35 = int32(1)
	v44 = int32(base.Ui32(v10)>>(uint(v35)%32)) - v35
	goto L15
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v44 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L15:
	;
	if v44 < int32(64) {
		v50 = v44
		v52 = v32
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v48 = F_pg_mbcliplen(m, v32, v44, int32(63))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v50 = v48
	v52 = v32
	goto L3
L18:
	;
	v57 = v50
	v59 = v52 - int32(1)
	v60 = v52
	goto L2
L19:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v59))))
	if v66 != int32(32) {
		v74 = v61
		v77 = v60
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v74 = int32(0)
	v77 = v60
	goto L1
L21:
	;
	v69 = int32(1)
	if v69 < v61 {
		v61 = v61 - v69
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v74 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	return v82
L25:
	;
	v81 = F__emscripten_memcpy_bulkmem(m, v79, v77, v74)
	mBase = m.M
	v82 = v81
	goto L27
L26:
	;
	v82 = v79
	goto L27
L27:
	;
	goto L24
}
func F_bpchar_pattern_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = l0 + int32(28)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(1)
	v24 = v7 + v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v27 = v25 & v23
	if v25 == v23 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v157 != v7 {
		goto L59
	} else {
		goto L60
	}
L5:
	;
	if v27 != 0 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v30 = int32(4)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v32&int32(254) == int32(2) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v45 = int32(1)
	if v27 != 0 {
		v55 = int32(base.Ui32(v25)>>(uint(v45)%32)) - v45
		goto L5
	} else {
		goto L15
	}
L9:
	;
	v41 = v30
	goto L11
L10:
	;
	v41 = base.B2i32(v32 == int32(18)) << (uint(v30) % 32)
	goto L11
L11:
	;
	if v32 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = v30
	goto L14
L13:
	;
	v44 = v41
	goto L14
L14:
	;
	v55 = v44
	goto L5
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L5
L16:
	;
	v58 = v24
	goto L18
L17:
	;
	v58 = v7 + int32(4)
	goto L18
L18:
	;
	v64 = v55
	goto L19
L19:
	;
	if v64 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v81 = int32(1)
	v82 = v14 + v81
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v85 = v83 & v81
	if v83 == v81 {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	goto L20
L22:
	;
	v80 = v55 >> (uint(int32(31)) % 32) & v55
	goto L21
L23:
	;
	goto L24
L24:
	;
	v74 = v64 - int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v74))))
	if v76 == int32(32) {
		v64 = v74
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v80 = v64
	goto L21
L26:
	;
	if v85 != 0 {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v88 = int32(4)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v90&int32(254) == int32(2) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v103 = int32(1)
	if v85 != 0 {
		v113 = int32(base.Ui32(v83)>>(uint(v103)%32)) - v103
		goto L26
	} else {
		goto L36
	}
L30:
	;
	v99 = v88
	goto L32
L31:
	;
	v99 = base.B2i32(v90 == int32(18)) << (uint(v88) % 32)
	goto L32
L32:
	;
	if v90 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v102 = v88
	goto L35
L34:
	;
	v102 = v99
	goto L35
L35:
	;
	v113 = v102
	goto L26
L36:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v113 = int32(base.Ui32(v107)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L37:
	;
	v116 = v82
	goto L39
L38:
	;
	v116 = v14 + int32(4)
	goto L39
L39:
	;
	v122 = v113
	goto L40
L40:
	;
	if v122 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v139 = int32(1)
	if v25&v139 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v138 = v113 >> (uint(int32(31)) % 32) & v113
	goto L42
L44:
	;
	goto L45
L45:
	;
	v132 = v122 - int32(1)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v132))))
	if v134 == int32(32) {
		v122 = v132
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v138 = v122
	goto L42
L47:
	;
	goto L4
L48:
	;
	v143 = v139
	goto L50
L49:
	;
	v143 = int32(4)
	goto L50
L50:
	;
	v145 = int32(1)
	if v83&v145 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v149 = v145
	goto L53
L52:
	;
	v149 = int32(4)
	goto L53
L53:
	;
	v151 = base.B2i32(v80 < v138)
	if v80 < v138 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v152 = v80
	goto L56
L55:
	;
	v152 = v138
	goto L56
L56:
	;
	v153 = F_memcmp(m, v7+v143, v14+v149, v152)
	mBase = m.M
	if v153 != 0 {
		v156 = v153
		goto L47
	} else {
		goto L57
	}
L57:
	;
	if v80 < v138 {
		v156 = int32(-1)
		goto L47
	} else {
		goto L58
	}
L58:
	;
	v156 = base.B2i32(v138 < v80)
	goto L47
L59:
	;
	F_pfree(m, v7)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v161 != v14 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	F_pfree(m, v14)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	return int32(base.Ui32(v156) >> (uint(int32(31)) % 32))
L66:
	;
	goto L65
}
func F_bpchar_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	v4 = int32(4476144)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v9
	F_varstr_sortsupport(m, v6, int32(1042), v7)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[28])) = v5
		return int32(0)
	}
}
