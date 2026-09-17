package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExprEvalPushStep(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v4 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
		v10 = F_palloc(m, int32(640))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v23 = v10
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
			v25 = v23
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v26 + int32(1)
			v32 = v25 + v26*int32(40)
			v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v33
			v35 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+24)) = v35
			v37 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v37
			v39 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v39
			v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v32))) = v41
			return
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		if v12 != v4 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v25 = v14
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v26 + int32(1)
			v32 = v25 + v26*int32(40)
			v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v33
			v35 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+24)) = v35
			v37 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v37
			v39 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v39
			v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v32))) = v41
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v4 << (uint(int32(1)) % 32)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v21 = F_repalloc(m, v18, v4*int32(80))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = v21
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
				v25 = v23
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v26 + int32(1)
				v32 = v25 + v26*int32(40)
				v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v33
				v35 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+24)) = v35
				v37 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v37
				v39 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v39
				v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, uint32(v32))) = v41
				return
			}
		}
	}
}
func F_exprType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
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
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	if l0 == v2 {
		v173 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(48)
	return v173
L2:
	;
	v11 = l0
	goto L4
L3:
	;
	v173 = v171
	goto L1
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	switch v17 - int32(6) {
	case 0:
		goto L33
	case 1:
		goto L32
	case 2:
		goto L31
	case 3, 5, 9, 21, 22, 24, 30, 34, 49, 53:
		goto L12
	case 4:
		v171 = int32(23)
		goto L3
	default:
		goto L8
	case 7:
		goto L30
	case 8:
		goto L29
	case 10:
		goto L28
	case 11, 12, 13:
		goto L27
	case 14, 15, 31, 40, 46, 47, 52:
		v173 = int32(16)
		goto L1
	case 16:
		goto L26
	case 17:
		goto L25
	case 18:
		goto L24
	case 19:
		goto L23
	case 20:
		goto L22
	case 23:
		goto L21
	case 25:
		goto L20
	case 26, 28, 29, 32, 33:
		goto L19
	case 35:
		goto L18
	case 38:
		goto L17
	case 39:
		goto L16
	case 41:
		goto L14
	case 42:
		goto L15
	case 50, 51:
		goto L13
	case 54:
		goto L11
	case 55:
		goto L10
	case 313:
		goto L9
	}
L5:
	;
	v171 = int32(0)
	goto L3
L6:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v168 != 0 {
		v11 = v168
		goto L4
	} else {
		goto L71
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L38
	} else {
		goto L68
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L38
	} else {
		goto L65
	}
L9:
	;
	v167 = v11 + int32(4)
	goto L6
L10:
	;
	v167 = v11 + int32(12)
	goto L6
L11:
	;
	v167 = v11 + int32(4)
	goto L6
L12:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v173 = v133
	goto L1
L13:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v173 = v132
	goto L1
L14:
	;
	v167 = v11 + int32(8)
	goto L6
L15:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v173 = v129
	goto L1
L16:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	v173 = v127
	goto L1
L17:
	;
	v167 = v11 + int32(8)
	goto L6
L18:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v117 == int32(6) {
		goto L59
	} else {
		goto L60
	}
L19:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v173 = v113
	goto L1
L20:
	;
	v167 = v11 + int32(4)
	goto L6
L21:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v173 = v110
	goto L1
L22:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v173 = v109
	goto L1
L23:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v173 = v108
	goto L1
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v167 = v107
	goto L6
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v76 - int32(4) {
	case 0, 2:
		goto L50
	case 1:
		v173 = int32(2249)
		goto L1
	default:
		goto L49
	}
L26:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v29 - int32(4) {
	case 0, 2:
		goto L35
	case 1:
		v173 = int32(2249)
		goto L1
	default:
		goto L34
	}
L27:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v173 = v27
	goto L1
L28:
	;
	v167 = v11 + int32(4)
	goto L6
L29:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v173 = v24
	goto L1
L30:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v173 = v23
	goto L1
L31:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v173 = v22
	goto L1
L32:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v173 = v21
	goto L1
L33:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v173 = v20
	goto L1
L34:
	;
	v173 = int32(16)
	goto L1
L35:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if v32 == int32(0) {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v35 != int32(67) {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+76))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v42 = F_exprType(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return int32(0)
L39:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v46 != int32(6) {
		v173 = v42
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v49 = F_get_promoted_array_type(m, v42)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	if v49 != 0 {
		v173 = v49
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v59 = F_exprType(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L38
	} else {
		goto L45
	}
L45:
	;
	v61 = F_format_type_be(m, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L38
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v61
	F_errmsg(m, int32(_a_F_exprType_0), v7+int32(16))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L38
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_exprType_1), int32(119), int32(_a_F_exprType_2))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L38
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v173 = int32(16)
	goto L1
L50:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v76 != int32(6) {
		v173 = v79
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v82 = F_get_promoted_array_type(m, v79)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L38
	} else {
		goto L52
	}
L52:
	;
	if v82 != 0 {
		v173 = v82
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L38
	} else {
		goto L54
	}
L54:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L38
	} else {
		goto L55
	}
L55:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v92 = F_format_type_be(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L38
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v92
	F_errmsg(m, int32(_a_F_exprType_0), v7+int32(32))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L38
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_exprType_1), int32(150), int32(_a_F_exprType_2))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L38
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v120 = int32(25)
	goto L61
L60:
	;
	v120 = int32(142)
	goto L61
L61:
	;
	if v117 == int32(7) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v123 = int32(16)
	goto L64
L63:
	;
	v123 = v120
	goto L64
L64:
	;
	v173 = v123
	goto L1
L65:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v144
	F_errmsg_internal(m, int32(_a_F_exprType_3), v7)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L38
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_exprType_1), int32(288), int32(_a_F_exprType_2))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L38
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errmsg_internal(m, int32(_a_F_exprType_4), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L38
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_exprType_1), int32(108), int32(_a_F_exprType_2))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L38
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	goto L5
}
func F_find_expr_references_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v821 int32
	_ = v821
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1053 int32
	_ = v1053
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1136 int32
	_ = v1136
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1206 int32
	_ = v1206
	var v1214 int32
	_ = v1214
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1250 int32
	_ = v1250
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1340 int32
	_ = v1340
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1384 int32
	_ = v1384
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1421 int32
	_ = v1421
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1459 int32
	_ = v1459
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1479 int32
	_ = v1479
	var v1484 int32
	_ = v1484
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1496 int32
	_ = v1496
	var v1501 int32
	_ = v1501
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	if l0 == v3 {
		v1421 = v3
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L48
	} else {
		goto L344
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L48
	} else {
		goto L340
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L48
	} else {
		goto L337
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L48
	} else {
		goto L334
	}
L5:
	;
	m.G0 = v15 + int32(80)
	return v1421
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v19 - int32(4) {
	case 0:
		goto L9
	default:
		goto L7
	case 2:
		goto L39
	case 3:
		goto L38
	case 4:
		goto L37
	case 5:
		goto L31
	case 7:
		goto L30
	case 10:
		goto L29
	case 11:
		goto L36
	case 13:
		goto L35
	case 14:
		goto L34
	case 15:
		goto L33
	case 16:
		goto L32
	case 19:
		goto L28
	case 21:
		goto L27
	case 22:
		goto L26
	case 23:
		goto L25
	case 24:
		goto L24
	case 25:
		goto L23
	case 26:
		goto L22
	case 27:
		goto L21
	case 32:
		goto L20
	case 33:
		goto L19
	case 51:
		goto L18
	case 55:
		goto L17
	case 62:
		goto L16
	case 63:
		goto L12
	case 99:
		goto L10
	case 100:
		goto L8
	case 102:
		goto L15
	case 104:
		goto L14
	case 110:
		goto L13
	case 138:
		goto L11
	}
L7:
	;
	v1415 = F_expression_tree_walker_impl(m, l0, int32(463), l1)
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L48
	} else {
		goto L333
	}
L8:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), v1397, int32(0), v1399)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L48
	} else {
		goto L332
	}
L9:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1262 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L10:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1128 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L11:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v1126 = F_find_expr_references_walker(m, v1125, l1)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L48
	} else {
		goto L287
	}
L12:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v810 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L13:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v790 != 0 {
		goto L221
	} else {
		goto L222
	}
L14:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v767 != 0 {
		goto L211
	} else {
		goto L212
	}
L15:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v754, int32(0), v756)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L48
	} else {
		goto L208
	}
L16:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v745 == int32(0) {
		goto L7
	} else {
		goto L206
	}
L17:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1259), v740, int32(0), v742)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L48
	} else {
		goto L205
	}
L18:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v734, int32(0), v736)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L48
	} else {
		goto L204
	}
L19:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v610 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L20:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v605, int32(0), v607)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L48
	} else {
		goto L184
	}
L21:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v599, int32(0), v601)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L48
	} else {
		goto L183
	}
L22:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v593, int32(0), v595)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L48
	} else {
		goto L182
	}
L23:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v576, int32(0), v578)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L48
	} else {
		goto L179
	}
L24:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v559, int32(0), v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L48
	} else {
		goto L176
	}
L25:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v542, int32(0), v544)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L48
	} else {
		goto L173
	}
L26:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v479 = F_get_typ_typrelid(m, v478)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L48
	} else {
		goto L159
	}
L27:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v450 = F_exprType(m, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L48
	} else {
		goto L149
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L48
	} else {
		goto L145
	}
L29:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v426 == v427 {
		goto L7
	} else {
		goto L142
	}
L30:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), v421, int32(0), v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L48
	} else {
		goto L141
	}
L31:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), v415, int32(0), v417)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L48
	} else {
		goto L140
	}
L32:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v409, int32(0), v411)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L48
	} else {
		goto L139
	}
L33:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v403, int32(0), v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L48
	} else {
		goto L138
	}
L34:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v397, int32(0), v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L48
	} else {
		goto L137
	}
L35:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v391, int32(0), v393)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L48
	} else {
		goto L136
	}
L36:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), v385, int32(0), v387)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L48
	} else {
		goto L135
	}
L37:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v368, int32(0), v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L48
	} else {
		goto L132
	}
L38:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v200, int32(0), v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L48
	} else {
		goto L79
	}
L39:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v22 == int32(0) {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if base.Ui32(v26) <= base.Ui32(v25) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v28 <= int32(0) {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v25<<(uint(int32(2))%32))))
	if v35 == int32(0) {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v38 < v28 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v40 == int32(0) {
		v1421 = v3
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+v28<<(uint(int32(2))%32)-int32(4))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	switch v50 {
	case 0:
		goto L47
	default:
		v1421 = v3
		goto L5
	case 3:
		goto L46
	}
L46:
	;
	v58 = int32(0)
	v60 = m.G0
	v62 = v60 - int32(16)
	m.G0 = v62
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v49)+68))
	if v65 == v58 {
		v159 = int32(1)
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1259), v52, v40, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	return int32(0)
L49:
	;
	v1421 = v3
	goto L5
L50:
	;
	m.G0 = v62 + int32(16)
	v1421 = v3
	goto L5
L51:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+72)))
	if base.B2i32(v160 == int32(1))&base.B2i32(v159 == v40) != 0 {
		goto L50
	} else {
		goto L74
	}
L52:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v69 <= int32(0) {
		v159 = int32(1)
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v72 = int32(0)
	if v72 < v69 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v75 = v69
	goto L56
L55:
	;
	v75 = v72
	goto L56
L56:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v77 = v58
	v82 = v58
	goto L58
L57:
	;
	v159 = v94 + int32(1)
	goto L51
L58:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v76+v82<<(uint(int32(2))%32))))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	v94 = v77 + v93
	if base.B2i32(v40 <= v94)&base.B2i32(v77 < v40) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	if v103 != 0 {
		goto L50
	} else {
		goto L64
	}
L60:
	;
	v101 = v82 + int32(1)
	if v75 != v101 {
		v77 = v94
		v82 = v101
		goto L58
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	goto L57
L64:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v106 = F_get_expr_result_tupdesc(m, v104, int32(1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L48
	} else {
		goto L65
	}
L65:
	;
	if v106 == int32(0) {
		goto L50
	} else {
		goto L66
	}
L66:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v110 == int32(2249) {
		goto L50
	} else {
		goto L67
	}
L67:
	;
	v113 = F_get_typ_typrelid(m, v110)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L48
	} else {
		goto L68
	}
L68:
	;
	if v113 == int32(0) {
		goto L50
	} else {
		goto L69
	}
L69:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	if v121 <= v120 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+12)) = v121 << (uint(int32(1)) % 32)
	v128 = F_repalloc(m, v119, v121*int32(24))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L48
	} else {
		goto L73
	}
L71:
	;
	v132 = v119
	v133 = v120
	goto L72
L72:
	;
	v136 = v133*int32(12) + v132
	*(*int32)(unsafe.Add(mBase, uint32(v136)+8)) = v40 - v77
	*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = int32(1259)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+8)) = v141 + int32(1)
	goto L50
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v128
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v132 = v128
	v133 = v131
	goto L72
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L48
	} else {
		goto L75
	}
L75:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L48
	} else {
		goto L76
	}
L76:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v40
	F_errmsg(m, int32(_a_F_find_expr_references_walker_0), v62)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L48
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(2391), int32(_a_F_find_expr_references_walker_2))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L48
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v206 = int32(0)
	if base.B2i32(v205 == v206)|base.B2i32(v205 == int32(100)) == v206 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v205, int32(0), v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L48
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v218 != 0 {
		v1421 = v3
		goto L5
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v219 <= int32(3733) {
		goto L94
	} else {
		goto L95
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L48
	} else {
		goto L128
	}
L86:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v336 = int32(0)
	v339 = F_SearchSysCacheExists(m, int32(38), v335, v336, v336, v336)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L48
	} else {
		goto L125
	}
L87:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v322 = int32(0)
	v325 = F_SearchSysCacheExists(m, int32(74), v321, v322, v322, v322)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L48
	} else {
		goto L122
	}
L88:
	;
	if v219 != int32(_a_F_find_expr_references_walker_3) {
		v1421 = v3
		goto L5
	} else {
		goto L118
	}
L89:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v292 = int32(0)
	v295 = F_SearchSysCacheExists(m, int32(82), v291, v292, v292, v292)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L48
	} else {
		goto L115
	}
L90:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v278 = int32(0)
	v281 = F_SearchSysCacheExists(m, int32(57), v277, v278, v278, v278)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L48
	} else {
		goto L112
	}
L91:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v264 = int32(0)
	v267 = F_SearchSysCacheExists(m, int32(40), v263, v264, v264, v264)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L48
	} else {
		goto L109
	}
L92:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v250 = int32(0)
	v253 = F_SearchSysCacheExists(m, int32(47), v249, v250, v250, v250)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L48
	} else {
		goto L106
	}
L93:
	;
	if v219 != int32(24) {
		v1421 = v3
		goto L5
	} else {
		goto L105
	}
L94:
	;
	switch v219 - int32(2202) {
	case 0:
		goto L92
	case 1, 2:
		goto L91
	case 3:
		goto L90
	case 4:
		goto L89
	default:
		goto L93
	}
L95:
	;
	goto L96
L96:
	;
	if v219 <= int32(4088) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	if v219 == int32(3734) {
		goto L87
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	switch v219 - int32(4089) {
	case 0:
		goto L86
	case 1, 2, 3, 4, 5, 6:
		v1421 = v3
		goto L5
	case 7:
		goto L85
	default:
		goto L88
	}
L100:
	;
	if v219 != int32(3769) {
		v1421 = v3
		goto L5
	} else {
		goto L101
	}
L101:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v232 = int32(0)
	v235 = F_SearchSysCacheExists(m, int32(76), v231, v232, v232, v232)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L48
	} else {
		goto L102
	}
L102:
	;
	if v235 == int32(0) {
		v1421 = v3
		goto L5
	} else {
		goto L103
	}
L103:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3600), v231, int32(0), v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L48
	} else {
		goto L104
	}
L104:
	;
	v1421 = v3
	goto L5
L105:
	;
	goto L92
L106:
	;
	if v253 == int32(0) {
		v1421 = v3
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), v249, int32(0), v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L48
	} else {
		goto L108
	}
L108:
	;
	v1421 = v3
	goto L5
L109:
	;
	if v267 == int32(0) {
		v1421 = v3
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v263, int32(0), v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L48
	} else {
		goto L111
	}
L111:
	;
	v1421 = v3
	goto L5
L112:
	;
	if v281 == int32(0) {
		v1421 = v3
		goto L5
	} else {
		goto L113
	}
L113:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1259), v277, int32(0), v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L48
	} else {
		goto L114
	}
L114:
	;
	v1421 = v3
	goto L5
L115:
	;
	if v295 == int32(0) {
		v1421 = v3
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v291, int32(0), v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L48
	} else {
		goto L117
	}
L117:
	;
	v1421 = v3
	goto L5
L118:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v308 = int32(0)
	v311 = F_SearchSysCacheExists(m, int32(16), v307, v308, v308, v308)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L48
	} else {
		goto L119
	}
L119:
	;
	if v311 == int32(0) {
		v1421 = v3
		goto L5
	} else {
		goto L120
	}
L120:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v307, int32(0), v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L48
	} else {
		goto L121
	}
L121:
	;
	v1421 = v3
	goto L5
L122:
	;
	if v325 == int32(0) {
		v1421 = v3
		goto L5
	} else {
		goto L123
	}
L123:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3602), v321, int32(0), v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L48
	} else {
		goto L124
	}
L124:
	;
	v1421 = v3
	goto L5
L125:
	;
	if v339 == int32(0) {
		v1421 = v3
		goto L5
	} else {
		goto L126
	}
L126:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2615), v335, int32(0), v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L48
	} else {
		goto L127
	}
L127:
	;
	v1421 = v3
	goto L5
L128:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L48
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(_a_F_find_expr_references_walker_4)
	F_errmsg(m, int32(_a_F_find_expr_references_walker_5), v15+int32(32))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L48
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(1851), int32(_a_F_find_expr_references_walker_6))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L48
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.B2i32(v373 == int32(0))|base.B2i32(v373 == int32(100)) != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v373, int32(0), v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L48
	} else {
		goto L134
	}
L134:
	;
	goto L7
L135:
	;
	goto L7
L136:
	;
	goto L7
L137:
	;
	goto L7
L138:
	;
	goto L7
L139:
	;
	goto L7
L140:
	;
	goto L7
L141:
	;
	goto L7
L142:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v426 == v429 {
		goto L7
	} else {
		goto L143
	}
L143:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v426, int32(0), v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L48
	} else {
		goto L144
	}
L144:
	;
	goto L7
L145:
	;
	F_errmsg_internal(m, int32(_a_F_find_expr_references_walker_7), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L48
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(1945), int32(_a_F_find_expr_references_walker_6))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L48
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.B2i32(v467 == int32(0))|base.B2i32(v467 == int32(100)) != 0 {
		goto L7
	} else {
		goto L157
	}
L149:
	;
	v452 = F_getBaseType(m, v450)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L48
	} else {
		goto L150
	}
L150:
	;
	v454 = F_get_typ_typrelid(m, v452)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L48
	} else {
		goto L151
	}
L151:
	;
	if v454 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v457 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1259), v454, v457, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L48
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v462, int32(0), v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L48
	} else {
		goto L156
	}
L155:
	;
	goto L148
L156:
	;
	goto L148
L157:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v467, int32(0), v475)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L48
	} else {
		goto L158
	}
L158:
	;
	goto L7
L159:
	;
	if v479 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v481 == int32(0) {
		goto L7
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v536, int32(0), v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L48
	} else {
		goto L172
	}
L163:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	if v484 <= int32(0) {
		goto L7
	} else {
		goto L164
	}
L164:
	;
	v489 = v3
	goto L165
L165:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v481)+12))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v499+v489<<(uint(int32(2))%32))))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v504)+8))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v504)+12))
	if v507 <= v506 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	goto L7
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v504)+12)) = v507 << (uint(int32(1)) % 32)
	v514 = F_repalloc(m, v505, v507*int32(24))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L48
	} else {
		goto L170
	}
L168:
	;
	v518 = v505
	v519 = v506
	goto L169
L169:
	;
	v522 = v519*int32(12) + v518
	*(*int32)(unsafe.Add(mBase, uint32(v522)+8)) = v503
	*(*int32)(unsafe.Add(mBase, uint32(v522)+4)) = v479
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = int32(1259)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v504)+8))
	v528 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v504)+8)) = v527 + v528
	v532 = v489 + v528
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	if v532 < v533 {
		v489 = v532
		goto L165
	} else {
		goto L171
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = v514
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v504)+8))
	v518 = v514
	v519 = v517
	goto L169
L171:
	;
	goto L166
L172:
	;
	goto L7
L173:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.B2i32(v547 == int32(0))|base.B2i32(v547 == int32(100)) != 0 {
		goto L7
	} else {
		goto L174
	}
L174:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v547, int32(0), v555)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L48
	} else {
		goto L175
	}
L175:
	;
	goto L7
L176:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.B2i32(v564 == int32(0))|base.B2i32(v564 == int32(100)) != 0 {
		goto L7
	} else {
		goto L177
	}
L177:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v564, int32(0), v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L48
	} else {
		goto L178
	}
L178:
	;
	goto L7
L179:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.B2i32(v581 == int32(0))|base.B2i32(v581 == int32(100)) != 0 {
		goto L7
	} else {
		goto L180
	}
L180:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v581, int32(0), v589)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L48
	} else {
		goto L181
	}
L181:
	;
	goto L7
L182:
	;
	goto L7
L183:
	;
	goto L7
L184:
	;
	goto L7
L185:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v677 == int32(0) {
		goto L7
	} else {
		goto L195
	}
L186:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v610)+4))
	if v613 <= int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v618 = v3
	goto L188
L188:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v610)+12))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v628+v618<<(uint(int32(2))%32))))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v633)+8))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v633)+12))
	if v636 <= v635 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	goto L185
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v633)+12)) = v636 << (uint(int32(1)) % 32)
	v643 = F_repalloc(m, v634, v636*int32(24))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L48
	} else {
		goto L193
	}
L191:
	;
	v647 = v634
	v648 = v635
	goto L192
L192:
	;
	v651 = v648*int32(12) + v647
	*(*int32)(unsafe.Add(mBase, uint32(v651)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v651)+4)) = v632
	*(*int32)(unsafe.Add(mBase, uint32(v651))) = int32(2617)
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v633)+8))
	v658 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v633)+8)) = v657 + v658
	v662 = v618 + v658
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v610)+4))
	if v662 < v663 {
		v618 = v662
		goto L188
	} else {
		goto L194
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v633))) = v643
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v633)+8))
	v647 = v643
	v648 = v646
	goto L192
L194:
	;
	goto L189
L195:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v677)+4))
	if v680 <= int32(0) {
		goto L7
	} else {
		goto L196
	}
L196:
	;
	v686 = int32(0)
	goto L197
L197:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v677)+12))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v696+v686<<(uint(int32(2))%32))))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v701)+8))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v701)+12))
	if v704 <= v703 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	goto L7
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v701)+12)) = v704 << (uint(int32(1)) % 32)
	v711 = F_repalloc(m, v702, v704*int32(24))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L48
	} else {
		goto L202
	}
L200:
	;
	v715 = v702
	v716 = v703
	goto L201
L201:
	;
	v719 = v716*int32(12) + v715
	*(*int32)(unsafe.Add(mBase, uint32(v719)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v719)+4)) = v700
	*(*int32)(unsafe.Add(mBase, uint32(v719))) = int32(2753)
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v701)+8))
	v726 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v701)+8)) = v725 + v726
	v730 = v686 + v726
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v677)+4))
	if v730 < v731 {
		v686 = v730
		goto L197
	} else {
		goto L203
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v701))) = v711
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v701)+8))
	v715 = v711
	v716 = v714
	goto L201
L203:
	;
	goto L198
L204:
	;
	goto L7
L205:
	;
	goto L7
L206:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2606), v745, int32(0), v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L48
	} else {
		goto L207
	}
L207:
	;
	goto L7
L208:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v759 == int32(0) {
		v1421 = v3
		goto L5
	} else {
		goto L209
	}
L209:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v759, int32(0), v764)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L48
	} else {
		goto L210
	}
L210:
	;
	v1421 = v3
	goto L5
L211:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), v767, int32(0), v770)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L48
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v773 != 0 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	goto L213
L215:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), v773, int32(0), v776)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L48
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.B2i32(v779 == int32(0))|base.B2i32(v779 == int32(100)) != 0 {
		goto L7
	} else {
		goto L219
	}
L218:
	;
	goto L217
L219:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v779, int32(0), v787)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L48
	} else {
		goto L220
	}
L220:
	;
	goto L7
L221:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v790, int32(0), v793)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L48
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v796 != 0 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	goto L223
L225:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v796, int32(0), v799)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L48
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v802 == int32(0) {
		goto L7
	} else {
		goto L229
	}
L228:
	;
	goto L227
L229:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v802, int32(0), v807)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L48
	} else {
		goto L230
	}
L230:
	;
	goto L7
L231:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v945&int32(-2) != int32(2) {
		goto L256
	} else {
		goto L257
	}
L232:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v810)+4))
	if v813 <= int32(0) {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v821 = v3
	goto L234
L234:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v810)+12))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v828+v821<<(uint(int32(2))%32))))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v832)+12))
	switch v833 {
	case 0:
		goto L238
	default:
		goto L236
	case 2:
		goto L237
	case 7:
		goto L2
	}
L235:
	;
	goto L231
L236:
	;
	v930 = v821 + int32(1)
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v810)+4))
	if v930 < v931 {
		v821 = v930
		goto L234
	} else {
		goto L255
	}
L237:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v865 = F_lcons(m, v863, v864)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L48
	} else {
		goto L243
	}
L238:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v832)+16))
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v835)))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v835)+8))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v835)+12))
	if v838 <= v837 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v835)+12)) = v838 << (uint(int32(1)) % 32)
	v845 = F_repalloc(m, v836, v838*int32(24))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L48
	} else {
		goto L242
	}
L240:
	;
	v849 = v836
	v850 = v837
	goto L241
L241:
	;
	v853 = v850*int32(12) + v849
	*(*int32)(unsafe.Add(mBase, uint32(v853)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v853)+4)) = v834
	*(*int32)(unsafe.Add(mBase, uint32(v853))) = int32(1259)
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v835)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v835)+8)) = v859 + int32(1)
	goto L236
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v835))) = v845
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v835)+8))
	v849 = v845
	v850 = v848
	goto L241
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v865
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v832)+48))
	if int32(0) < v868 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v875 = v868
	v876 = int32(0)
	goto L247
L245:
	;
	v913 = v865
	goto L246
L246:
	;
	v914 = F_list_delete_first(m, v913)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L48
	} else {
		goto L254
	}
L247:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v832)+52))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v884)+12))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v885+v876<<(uint(int32(2))%32))))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v889)))
	if v890 != int32(6) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v913 = v900
	goto L246
L249:
	;
	v893 = F_find_expr_references_walker(m, v889, l1)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L48
	} else {
		goto L252
	}
L250:
	;
	v896 = v875
	goto L251
L251:
	;
	v898 = v876 + int32(1)
	if v898 < v896 {
		v875 = v896
		v876 = v898
		goto L247
	} else {
		goto L253
	}
L252:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v832)+48))
	v896 = v895
	goto L251
L253:
	;
	goto L248
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v914
	goto L236
L255:
	;
	goto L235
L256:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v1044 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L257:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v950 <= int32(0) {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v953 == int32(0) {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v953)+4))
	if v956 < v950 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v953)+12))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v958+v950<<(uint(int32(2))%32)-int32(4))))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v964)+12))
	if v965 != 0 {
		goto L256
	} else {
		goto L261
	}
L261:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v966 == int32(0) {
		goto L256
	} else {
		goto L262
	}
L262:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v966)+4))
	if v969 <= int32(0) {
		goto L256
	} else {
		goto L263
	}
L263:
	;
	v975 = int32(0)
	goto L264
L264:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v966)+12))
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v985+v975<<(uint(int32(2))%32))))
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v989)+26)))
	if v990 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	goto L256
L266:
	;
	v993 = int32(*(*int16)(unsafe.Add(mBase, uint32(v989)+8)))
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v964)+16))
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v995)))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v995)+8))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v995)+12))
	if v998 <= v997 {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	goto L268
L268:
	;
	v1029 = v975 + int32(1)
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v966)+4))
	if v1029 < v1030 {
		v975 = v1029
		goto L264
	} else {
		goto L273
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v995)+12)) = v998 << (uint(int32(1)) % 32)
	v1005 = F_repalloc(m, v996, v998*int32(24))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L48
	} else {
		goto L272
	}
L270:
	;
	v1009 = v996
	v1010 = v997
	goto L271
L271:
	;
	v1013 = v1010*int32(12) + v1009
	*(*int32)(unsafe.Add(mBase, uint32(v1013)+8)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v1013)+4)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v1013))) = int32(1259)
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v995)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v995)+8)) = v1018 + int32(1)
	goto L268
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v995))) = v1005
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v995)+8))
	v1009 = v1005
	v1010 = v1008
	goto L271
L273:
	;
	goto L265
L274:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1114 = F_lcons(m, v1112, v1113)
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L48
	} else {
		goto L284
	}
L275:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1044)+4))
	if v1047 <= int32(0) {
		goto L274
	} else {
		goto L276
	}
L276:
	;
	v1053 = int32(0)
	goto L277
L277:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1044)+12))
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1063+v1053<<(uint(int32(2))%32))))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1068)))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1068)+8))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1068)+12))
	if v1071 <= v1070 {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	goto L274
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1068)+12)) = v1071 << (uint(int32(1)) % 32)
	v1078 = F_repalloc(m, v1069, v1071*int32(24))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L48
	} else {
		goto L282
	}
L280:
	;
	v1082 = v1069
	v1083 = v1070
	goto L281
L281:
	;
	v1086 = v1083*int32(12) + v1082
	*(*int32)(unsafe.Add(mBase, uint32(v1086)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1086)+4)) = v1067
	*(*int32)(unsafe.Add(mBase, uint32(v1086))) = int32(2606)
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1068)+8))
	v1093 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1068)+8)) = v1092 + v1093
	v1097 = v1053 + v1093
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1044)+4))
	if v1097 < v1098 {
		v1053 = v1097
		goto L277
	} else {
		goto L283
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1068))) = v1078
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1068)+8))
	v1082 = v1078
	v1083 = v1081
	goto L281
L283:
	;
	goto L278
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1114
	v1119 = F_query_tree_walker_impl(m, l0, int32(463), l1, int32(132))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L48
	} else {
		goto L285
	}
L285:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1122 = F_list_delete_first(m, v1121)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L48
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1122
	v1421 = v1119
	goto L5
L287:
	;
	goto L7
L288:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1195 == int32(0) {
		goto L7
	} else {
		goto L298
	}
L289:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+4))
	if v1131 <= int32(0) {
		goto L288
	} else {
		goto L290
	}
L290:
	;
	v1136 = v3
	goto L291
L291:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+12))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1146+v1136<<(uint(int32(2))%32))))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1151)))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+8))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+12))
	if v1154 <= v1153 {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	goto L288
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1151)+12)) = v1154 << (uint(int32(1)) % 32)
	v1161 = F_repalloc(m, v1152, v1154*int32(24))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L48
	} else {
		goto L296
	}
L294:
	;
	v1165 = v1152
	v1166 = v1153
	goto L295
L295:
	;
	v1169 = v1166*int32(12) + v1165
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+4)) = v1150
	*(*int32)(unsafe.Add(mBase, uint32(v1169))) = int32(1247)
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+8))
	v1176 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1151)+8)) = v1175 + v1176
	v1180 = v1136 + v1176
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+4))
	if v1180 < v1181 {
		v1136 = v1180
		goto L291
	} else {
		goto L297
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1151))) = v1161
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+8))
	v1165 = v1161
	v1166 = v1164
	goto L295
L297:
	;
	goto L292
L298:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1195)+4))
	if v1198 <= int32(0) {
		goto L7
	} else {
		goto L299
	}
L299:
	;
	v1206 = int32(0)
	goto L300
L300:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1195)+12))
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1214+v1206<<(uint(int32(2))%32))))
	v1219 = int32(0)
	if base.B2i32(v1218 == v1219)|base.B2i32(v1218 == int32(100)) == v1219 {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	goto L7
L302:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1226)))
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1226)+8))
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1226)+12))
	if v1229 <= v1228 {
		goto L305
	} else {
		goto L306
	}
L303:
	;
	goto L304
L304:
	;
	v1259 = v1206 + int32(1)
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1195)+4))
	if v1259 < v1260 {
		v1206 = v1259
		goto L300
	} else {
		goto L309
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1226)+12)) = v1229 << (uint(int32(1)) % 32)
	v1236 = F_repalloc(m, v1227, v1229*int32(24))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L48
	} else {
		goto L308
	}
L306:
	;
	v1240 = v1227
	v1241 = v1228
	goto L307
L307:
	;
	v1244 = v1241*int32(12) + v1240
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+4)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v1244))) = int32(3456)
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1226)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1226)+8)) = v1250 + int32(1)
	goto L304
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1226))) = v1236
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1226)+8))
	v1240 = v1236
	v1241 = v1239
	goto L307
L309:
	;
	goto L301
L310:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1329 == int32(0) {
		goto L7
	} else {
		goto L320
	}
L311:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1262)+4))
	if v1265 <= int32(0) {
		goto L310
	} else {
		goto L312
	}
L312:
	;
	v1270 = v3
	goto L313
L313:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1262)+12))
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1280+v1270<<(uint(int32(2))%32))))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1285)))
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1285)+8))
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1285)+12))
	if v1288 <= v1287 {
		goto L315
	} else {
		goto L316
	}
L314:
	;
	goto L310
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1285)+12)) = v1288 << (uint(int32(1)) % 32)
	v1295 = F_repalloc(m, v1286, v1288*int32(24))
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L48
	} else {
		goto L318
	}
L316:
	;
	v1299 = v1286
	v1300 = v1287
	goto L317
L317:
	;
	v1303 = v1300*int32(12) + v1299
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+4)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v1303))) = int32(1247)
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1285)+8))
	v1310 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1285)+8)) = v1309 + v1310
	v1314 = v1270 + v1310
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1262)+4))
	if v1314 < v1315 {
		v1270 = v1314
		goto L313
	} else {
		goto L319
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1285))) = v1295
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1285)+8))
	v1299 = v1295
	v1300 = v1298
	goto L317
L319:
	;
	goto L314
L320:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1329)+4))
	if v1332 <= int32(0) {
		goto L7
	} else {
		goto L321
	}
L321:
	;
	v1340 = int32(0)
	goto L322
L322:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1329)+12))
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1348+v1340<<(uint(int32(2))%32))))
	v1353 = int32(0)
	if base.B2i32(v1352 == v1353)|base.B2i32(v1352 == int32(100)) == v1353 {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	goto L7
L324:
	;
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v1360)))
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+8))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+12))
	if v1363 <= v1362 {
		goto L327
	} else {
		goto L328
	}
L325:
	;
	goto L326
L326:
	;
	v1393 = v1340 + int32(1)
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1329)+4))
	if v1393 < v1394 {
		v1340 = v1393
		goto L322
	} else {
		goto L331
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1360)+12)) = v1363 << (uint(int32(1)) % 32)
	v1370 = F_repalloc(m, v1361, v1363*int32(24))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L48
	} else {
		goto L330
	}
L328:
	;
	v1374 = v1361
	v1375 = v1362
	goto L329
L329:
	;
	v1378 = v1375*int32(12) + v1374
	*(*int32)(unsafe.Add(mBase, uint32(v1378)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1378)+4)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v1378))) = int32(3456)
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1360)+8)) = v1384 + int32(1)
	goto L326
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1360))) = v1370
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+8))
	v1374 = v1370
	v1375 = v1373
	goto L329
L331:
	;
	goto L323
L332:
	;
	goto L7
L333:
	;
	v1421 = v1415
	goto L5
L334:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v1438
	F_errmsg_internal(m, int32(_a_F_find_expr_references_walker_8), v15)
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L48
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(1711), int32(_a_F_find_expr_references_walker_6))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L48
	} else {
		goto L336
	}
L336:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L337:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v1453
	F_errmsg_internal(m, int32(_a_F_find_expr_references_walker_9), v15+int32(16))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L48
	} else {
		goto L338
	}
L338:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(1714), int32(_a_F_find_expr_references_walker_6))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L48
	} else {
		goto L339
	}
L339:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L340:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L48
	} else {
		goto L341
	}
L341:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v832)+8))
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v1473
	F_errmsg(m, int32(_a_F_find_expr_references_walker_10), v15-int32(-64))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L48
	} else {
		goto L342
	}
L342:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(2206), int32(_a_F_find_expr_references_walker_6))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L48
	} else {
		goto L343
	}
L343:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L344:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v1490
	F_errmsg_internal(m, int32(_a_F_find_expr_references_walker_11), v15+int32(48))
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L48
	} else {
		goto L345
	}
L345:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(2229), int32(_a_F_find_expr_references_walker_6))
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L48
	} else {
		goto L346
	}
L346:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_match_expr_to_partition_keys(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 == int32(27) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = l0
	goto L4
L2:
	;
	v24 = l0
	goto L3
L3:
	;
	v32 = int32(-1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+2)))
	if v34 <= int32(0) {
		v142 = v32
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v21 == int32(27) {
		v12 = v20
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v24 = v20
	goto L3
L6:
	;
	goto L5
L7:
	;
	return v142
L8:
	;
	v42 = int32(0)
	goto L9
L9:
	;
	v46 = v42 << (uint(int32(2)) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46+v47)))
	if v49 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v142 = v32
	goto L7
L11:
	;
	v132 = v42 + int32(1)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133)+2)))
	if v132 < v134 {
		v42 = v132
		goto L9
	} else {
		goto L30
	}
L12:
	;
	v142 = v42
	goto L7
L13:
	;
	if l2 == int32(0) {
		goto L11
	} else {
		goto L22
	}
L14:
	;
	v52 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v53 <= v52 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v60 = v52
	goto L16
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v60<<(uint(int32(2))%32))))
	v69 = F_equal(m, v68, v24)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L13
L18:
	;
	return int32(0)
L19:
	;
	if v69 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v74 = v60 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v74 < v75 {
		v60 = v74
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+268))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87+v46)))
	if v89 == int32(0) {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	v92 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v93 <= v92 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v100 = v92
	goto L25
L25:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v100<<(uint(int32(2))%32))))
	v109 = F_equal(m, v108, v24)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L18
	} else {
		goto L27
	}
L26:
	;
	goto L11
L27:
	;
	if v109 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v112 = v100 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v112 < v113 {
		v100 = v112
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	goto L10
}
func F_transformExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = l2
	v7 = F_transformExprRecurse(m, l0, l1)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v5
		return v7
	}
}
