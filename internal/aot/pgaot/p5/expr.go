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
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4 == int32(0) {
		v7 = int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v7
		v11 = F_palloc_mul(m, int32(40), v7)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v24 = v11
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v24
			v26 = v24
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v27 + int32(1)
			v33 = v26 + v27*int32(40)
			v34 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v33)+32)) = v34
			v36 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v36
			v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = v38
			v40 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v40
			v42 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v33))) = v42
			return
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v13 != v4 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v26 = v15
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v27 + int32(1)
			v33 = v26 + v27*int32(40)
			v34 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v33)+32)) = v34
			v36 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v36
			v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = v38
			v40 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v40
			v42 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v33))) = v42
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v4 << (uint(int32(1)) % 32)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v22 = F_repalloc(m, v19, v4*int32(80))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = v22
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v24
				v26 = v24
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v27 + int32(1)
				v33 = v26 + v27*int32(40)
				v34 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v33)+32)) = v34
				v36 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v36
				v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = v38
				v40 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v40
				v42 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, uint32(v33))) = v42
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
	case 315:
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
func F_expr_is_nonnullable(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	v4 = int32(0)
	v9 = l1
	goto L2
L1:
	;
	return v264 & int32(1)
L2:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v264 = int32(1)
	goto L1
L4:
	;
	goto L3
L5:
	;
	return int32(0)
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v19 != int32(27) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	switch v19 - int32(6) {
	case 0:
		goto L14
	case 1:
		goto L13
	default:
		v264 = v4
		goto L1
	case 12, 29, 46, 47:
		goto L4
	case 26:
		goto L10
	case 32:
		goto L12
	case 33:
		goto L11
	}
L8:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v9 = v255
	goto L2
L10:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v225 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L75
	}
L11:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v200 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L69
	}
L12:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v175 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L63
	}
L13:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+32)))
	v264 = v172 ^ int32(1)
	goto L1
L14:
	;
	if l0 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if v30 != 0 {
		v167 = v4
		goto L16
	} else {
		goto L17
	}
L16:
	;
	m.G0 = v28 + int32(16)
	v264 = v167
	goto L1
L17:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if v31 != 0 {
		v167 = v4
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	if v32 != 0 {
		v167 = v4
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+8)))
	if v33 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v167 = int32(1)
	goto L16
L21:
	;
	goto L22
L22:
	;
	if v33 == int32(0) {
		v167 = v4
		goto L16
	} else {
		goto L23
	}
L23:
	;
	switch l2 {
	case 0:
		goto L24
	case 1:
		goto L27
	case 2:
		goto L26
	default:
		goto L25
	}
L24:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v158 = F_find_base_rel(m, l0, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L5
	} else {
		goto L61
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L58
	}
L26:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v68 != 0 {
		goto L40
	} else {
		goto L41
	}
L27:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v39 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if v55 != 0 {
		v167 = v4
		goto L16
	} else {
		goto L32
	}
L29:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v53 = v39 + v40<<(uint(int32(2))%32)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+52))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v53 = v46 + v47<<(uint(int32(2))%32) - int32(4)
	goto L28
L32:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+20)))
	if v56 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+21)))
	if v59 != int32(112) {
		v167 = v4
		goto L16
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v63 = F_find_relation_notnullatts(m, l0, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+8)))
	v66 = F_bms_is_member(m, v65, v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v167 = v66
	goto L16
L39:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	if v84 != 0 {
		v167 = v4
		goto L16
	} else {
		goto L43
	}
L40:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v82 = v68 + v69<<(uint(int32(2))%32)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+52))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v82 = v75 + v76<<(uint(int32(2))%32) - int32(4)
	goto L39
L43:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+20)))
	if v85 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v131 = F_table_open(m, v129, int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L56
	}
L45:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v89 = m.G0
	v91 = v89 - int32(16)
	m.G0 = v91
	v95 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v88))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	if v95 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+22)))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+v113)+126)))
	F_ReleaseCatCache(m, v95)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L5
	} else {
		goto L53
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v88
	F_errmsg_internal(m, int32(_a_F_expr_is_nonnullable_0), v91)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_expr_is_nonnullable_1), int32(363), int32(_a_F_expr_is_nonnullable_2))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	m.G0 = v91 + int32(16)
	if v115 == int32(0) {
		goto L44
	} else {
		goto L54
	}
L54:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+21)))
	if v123 != int32(112) {
		v167 = v4
		goto L16
	} else {
		goto L55
	}
L55:
	;
	goto L44
L56:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+52))
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+8)))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+v134<<(uint(int32(3))%32))+27)))
	F_relation_close(m, v131, int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v167 = base.B2i32(v138 == int32(118))
	goto L16
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = l2
	F_errmsg_internal(m, int32(_a_F_expr_is_nonnullable_3), v28)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_expr_is_nonnullable_4), int32(_a_F_expr_is_nonnullable_5), int32(_a_F_expr_is_nonnullable_6))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+8)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)+100))
	v162 = F_bms_is_member(m, v160, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v167 = v162
	goto L16
L63:
	;
	v180 = int32(0)
	goto L64
L64:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	v187 = base.B2i32(v180 < v186)
	if v186 <= v180 {
		v264 = v187
		goto L1
	} else {
		goto L66
	}
L65:
	;
	v264 = v187
	goto L1
L66:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v180<<(uint(int32(2))%32)+v193)))
	v196 = F_expr_is_nonnullable(m, l0, v195, l2)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	if v196 == int32(0) {
		v180 = v180 + int32(1)
		goto L64
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	v205 = int32(0)
	goto L70
L70:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	v212 = base.B2i32(v205 < v211)
	if v211 <= v205 {
		v264 = v212
		goto L1
	} else {
		goto L72
	}
L71:
	;
	v264 = v212
	goto L1
L72:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v205<<(uint(int32(2))%32)+v218)))
	v221 = F_expr_is_nonnullable(m, l0, v220, l2)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	if v221 == int32(0) {
		v205 = v205 + int32(1)
		goto L70
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	v228 = F_expr_is_nonnullable(m, l0, v225, l2)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	if v228 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v232 == int32(0) {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	v237 = int32(0)
	goto L79
L79:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	v244 = base.B2i32(v243 <= v237)
	if v243 <= v237 {
		v264 = v244
		goto L1
	} else {
		goto L81
	}
L80:
	;
	v264 = v244
	goto L1
L81:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v237<<(uint(int32(2))%32)+v249)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	v253 = F_expr_is_nonnullable(m, l0, v252, l2)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	if v253 != 0 {
		v237 = v237 + int32(1)
		goto L79
	} else {
		goto L83
	}
L83:
	;
	goto L80
}
func F_find_expr_references_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v252 int64
	_ = v252
	var v253 int64
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v292 int64
	_ = v292
	var v293 int64
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int64
	_ = v307
	var v308 int64
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int64
	_ = v322
	var v323 int64
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int64
	_ = v337
	var v338 int64
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int64
	_ = v352
	var v353 int64
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int64
	_ = v367
	var v368 int64
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
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
	var v598 int32
	_ = v598
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v858 int32
	_ = v858
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1097 int32
	_ = v1097
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1254 int32
	_ = v1254
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1299 int32
	_ = v1299
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1319 int32
	_ = v1319
	var v1330 int32
	_ = v1330
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1391 int32
	_ = v1391
	var v1400 int32
	_ = v1400
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1436 int32
	_ = v1436
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1474 int32
	_ = v1474
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1513 int32
	_ = v1513
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1533 int32
	_ = v1533
	var v1538 int32
	_ = v1538
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1550 int32
	_ = v1550
	var v1555 int32
	_ = v1555
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	if l0 == v3 {
		v1474 = v3
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L48
	} else {
		goto L350
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L48
	} else {
		goto L346
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L48
	} else {
		goto L343
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L48
	} else {
		goto L340
	}
L5:
	;
	m.G0 = v16 + int32(96)
	return v1474
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v20 - int32(4) {
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
	v1468 = F_expression_tree_walker_impl(m, l0, int32(498), l1)
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L48
	} else {
		goto L339
	}
L8:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), v1449, int32(0), v1451)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L48
	} else {
		goto L338
	}
L9:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1311 == int32(0) {
		goto L316
	} else {
		goto L317
	}
L10:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1174 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L11:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v1172 = F_find_expr_references_walker(m, v1171, l1)
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L48
	} else {
		goto L293
	}
L12:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v847 == int32(0) {
		goto L237
	} else {
		goto L238
	}
L13:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v827 != 0 {
		goto L227
	} else {
		goto L228
	}
L14:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v804 != 0 {
		goto L217
	} else {
		goto L218
	}
L15:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v791, int32(0), v793)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L48
	} else {
		goto L214
	}
L16:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v782 == int32(0) {
		goto L7
	} else {
		goto L212
	}
L17:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1259), v777, int32(0), v779)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L48
	} else {
		goto L211
	}
L18:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v771, int32(0), v773)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L48
	} else {
		goto L210
	}
L19:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v644 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L20:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v639, int32(0), v641)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L48
	} else {
		goto L190
	}
L21:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v633, int32(0), v635)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L48
	} else {
		goto L189
	}
L22:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v627, int32(0), v629)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L48
	} else {
		goto L188
	}
L23:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v610, int32(0), v612)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L48
	} else {
		goto L185
	}
L24:
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
L25:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v576, int32(0), v578)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L48
	} else {
		goto L179
	}
L26:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v512 = F_get_typ_typrelid(m, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L48
	} else {
		goto L165
	}
L27:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v483 = F_exprType(m, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L48
	} else {
		goto L155
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L48
	} else {
		goto L151
	}
L29:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v459 == v460 {
		goto L7
	} else {
		goto L148
	}
L30:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), v454, int32(0), v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L48
	} else {
		goto L147
	}
L31:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), v448, int32(0), v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L48
	} else {
		goto L146
	}
L32:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v442, int32(0), v444)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L48
	} else {
		goto L145
	}
L33:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v436, int32(0), v438)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L48
	} else {
		goto L144
	}
L34:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v430, int32(0), v432)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L48
	} else {
		goto L143
	}
L35:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v424, int32(0), v426)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L48
	} else {
		goto L142
	}
L36:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), v418, int32(0), v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L48
	} else {
		goto L141
	}
L37:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v401, int32(0), v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L48
	} else {
		goto L138
	}
L38:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v204, int32(0), v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L48
	} else {
		goto L79
	}
L39:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v23 == int32(0) {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if base.Ui32(v27) <= base.Ui32(v26) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v29 <= int32(0) {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v26<<(uint(int32(2))%32))))
	if v36 == int32(0) {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v39 < v29 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v41 == int32(0) {
		v1474 = v3
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+v29<<(uint(int32(2))%32)-int32(4))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	switch v51 {
	case 0:
		goto L47
	default:
		v1474 = v3
		goto L5
	case 3:
		goto L46
	}
L46:
	;
	v59 = int32(0)
	v61 = m.G0
	v63 = v61 - int32(16)
	m.G0 = v63
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v50)+68))
	if v66 == v59 {
		v162 = int32(1)
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1259), v53, v41, v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	return int32(0)
L49:
	;
	v1474 = v3
	goto L5
L50:
	;
	m.G0 = v63 + int32(16)
	v1474 = v3
	goto L5
L51:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+72)))
	if base.B2i32(v163 == int32(1))&base.B2i32(v162 == v41) != 0 {
		goto L50
	} else {
		goto L74
	}
L52:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v70 <= int32(0) {
		v162 = int32(1)
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v73 = int32(0)
	if v73 < v70 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v76 = v70
	goto L56
L55:
	;
	v76 = v73
	goto L56
L56:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v78 = v59
	v83 = v59
	goto L58
L57:
	;
	v162 = v96 + int32(1)
	goto L51
L58:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v77+v83<<(uint(int32(2))%32))))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v96 = v78 + v95
	if base.B2i32(v41 <= v96)&base.B2i32(v78 < v41) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	if v105 != 0 {
		goto L50
	} else {
		goto L64
	}
L60:
	;
	v103 = v83 + int32(1)
	if v76 != v103 {
		v78 = v96
		v83 = v103
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v108 = F_get_expr_result_tupdesc(m, v106, int32(1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L48
	} else {
		goto L65
	}
L65:
	;
	if v108 == int32(0) {
		goto L50
	} else {
		goto L66
	}
L66:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v112 == int32(2249) {
		goto L50
	} else {
		goto L67
	}
L67:
	;
	v115 = F_get_typ_typrelid(m, v112)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L48
	} else {
		goto L68
	}
L68:
	;
	if v115 == int32(0) {
		goto L50
	} else {
		goto L69
	}
L69:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	if v123 <= v122 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+12)) = v123 << (uint(int32(1)) % 32)
	v130 = F_repalloc(m, v121, v123*int32(24))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L48
	} else {
		goto L73
	}
L71:
	;
	v134 = v121
	v135 = v122
	goto L72
L72:
	;
	v138 = v135*int32(12) + v134
	*(*int32)(unsafe.Add(mBase, uint32(v138)+8)) = v41 - v78
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = int32(1259)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = v143 + int32(1)
	goto L50
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v130
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v134 = v130
	v135 = v133
	goto L72
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L48
	} else {
		goto L75
	}
L75:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L48
	} else {
		goto L76
	}
L76:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v41
	F_errmsg(m, int32(_a_F_find_expr_references_walker_0), v63)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L48
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(2529), int32(_a_F_find_expr_references_walker_2))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
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
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v210 = int32(0)
	if base.B2i32(v209 == v210)|base.B2i32(v209 == int32(100)) == v210 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v209, int32(0), v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L48
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v222 != 0 {
		v1474 = v3
		goto L5
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v223 <= int32(3733) {
		goto L92
	} else {
		goto L93
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L48
	} else {
		goto L134
	}
L86:
	;
	v367 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v368 = int64(0)
	v371 = F_SearchSysCacheExists(m, int32(38), v367, v368, v368, v368)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L48
	} else {
		goto L131
	}
L87:
	;
	v352 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v353 = int64(0)
	v356 = F_SearchSysCacheExists(m, int32(74), v352, v353, v353, v353)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L48
	} else {
		goto L128
	}
L88:
	;
	v337 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v338 = int64(0)
	v341 = F_SearchSysCacheExists(m, int32(16), v337, v338, v338, v338)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L48
	} else {
		goto L125
	}
L89:
	;
	v322 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v323 = int64(0)
	v326 = F_SearchSysCacheExists(m, int32(82), v322, v323, v323, v323)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L48
	} else {
		goto L122
	}
L90:
	;
	v307 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v308 = int64(0)
	v311 = F_SearchSysCacheExists(m, int32(57), v307, v308, v308, v308)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L48
	} else {
		goto L119
	}
L91:
	;
	v292 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v293 = int64(0)
	v296 = F_SearchSysCacheExists(m, int32(40), v292, v293, v293, v293)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L48
	} else {
		goto L116
	}
L92:
	;
	switch v223 - int32(2202) {
	case 0:
		goto L95
	case 1, 2:
		goto L91
	case 3:
		goto L90
	case 4:
		goto L89
	default:
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	if v223 <= int32(4088) {
		goto L101
	} else {
		goto L102
	}
L95:
	;
	v231 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v232 = int64(0)
	v235 = F_SearchSysCacheExists(m, int32(47), v231, v232, v232, v232)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L48
	} else {
		goto L98
	}
L96:
	;
	if v223 != int32(24) {
		v1474 = v3
		goto L5
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	if v235 == int32(0) {
		v1474 = v3
		goto L5
	} else {
		goto L99
	}
L99:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), base.I32_wrap_i64(v231), int32(0), v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L48
	} else {
		goto L100
	}
L100:
	;
	v1474 = v3
	goto L5
L101:
	;
	if v223 == int32(3734) {
		goto L87
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	switch v223 - int32(4089) {
	case 0:
		goto L86
	case 1, 2, 3, 4, 5, 6:
		v1474 = v3
		goto L5
	case 7:
		goto L85
	default:
		goto L109
	}
L104:
	;
	if v223 != int32(3769) {
		v1474 = v3
		goto L5
	} else {
		goto L105
	}
L105:
	;
	v252 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v253 = int64(0)
	v256 = F_SearchSysCacheExists(m, int32(76), v252, v253, v253, v253)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L48
	} else {
		goto L106
	}
L106:
	;
	if v256 == int32(0) {
		v1474 = v3
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3600), base.I32_wrap_i64(v252), int32(0), v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L48
	} else {
		goto L108
	}
L108:
	;
	v1474 = v3
	goto L5
L109:
	;
	if v223 == int32(_a_F_find_expr_references_walker_3) {
		goto L88
	} else {
		goto L110
	}
L110:
	;
	if v223 != int32(_a_F_find_expr_references_walker_4) {
		v1474 = v3
		goto L5
	} else {
		goto L111
	}
L111:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L48
	} else {
		goto L112
	}
L112:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L48
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = int32(_a_F_find_expr_references_walker_5)
	F_errmsg(m, int32(_a_F_find_expr_references_walker_6), v16+int32(48))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L48
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(1989), int32(_a_F_find_expr_references_walker_7))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L48
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	if v296 == int32(0) {
		v1474 = v3
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), base.I32_wrap_i64(v292), int32(0), v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L48
	} else {
		goto L118
	}
L118:
	;
	v1474 = v3
	goto L5
L119:
	;
	if v311 == int32(0) {
		v1474 = v3
		goto L5
	} else {
		goto L120
	}
L120:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1259), base.I32_wrap_i64(v307), int32(0), v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L48
	} else {
		goto L121
	}
L121:
	;
	v1474 = v3
	goto L5
L122:
	;
	if v326 == int32(0) {
		v1474 = v3
		goto L5
	} else {
		goto L123
	}
L123:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), base.I32_wrap_i64(v322), int32(0), v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L48
	} else {
		goto L124
	}
L124:
	;
	v1474 = v3
	goto L5
L125:
	;
	if v341 == int32(0) {
		v1474 = v3
		goto L5
	} else {
		goto L126
	}
L126:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), base.I32_wrap_i64(v337), int32(0), v348)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L48
	} else {
		goto L127
	}
L127:
	;
	v1474 = v3
	goto L5
L128:
	;
	if v356 == int32(0) {
		v1474 = v3
		goto L5
	} else {
		goto L129
	}
L129:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3602), base.I32_wrap_i64(v352), int32(0), v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L48
	} else {
		goto L130
	}
L130:
	;
	v1474 = v3
	goto L5
L131:
	;
	if v371 == int32(0) {
		v1474 = v3
		goto L5
	} else {
		goto L132
	}
L132:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2615), base.I32_wrap_i64(v367), int32(0), v378)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L48
	} else {
		goto L133
	}
L133:
	;
	v1474 = v3
	goto L5
L134:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L48
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(_a_F_find_expr_references_walker_8)
	F_errmsg(m, int32(_a_F_find_expr_references_walker_6), v16+int32(32))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L48
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(1978), int32(_a_F_find_expr_references_walker_7))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L48
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.B2i32(v406 == int32(0))|base.B2i32(v406 == int32(100)) != 0 {
		goto L7
	} else {
		goto L139
	}
L139:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v406, int32(0), v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L48
	} else {
		goto L140
	}
L140:
	;
	goto L7
L141:
	;
	goto L7
L142:
	;
	goto L7
L143:
	;
	goto L7
L144:
	;
	goto L7
L145:
	;
	goto L7
L146:
	;
	goto L7
L147:
	;
	goto L7
L148:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v459 == v462 {
		goto L7
	} else {
		goto L149
	}
L149:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v459, int32(0), v466)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L48
	} else {
		goto L150
	}
L150:
	;
	goto L7
L151:
	;
	F_errmsg_internal(m, int32(_a_F_find_expr_references_walker_9), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L48
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(2083), int32(_a_F_find_expr_references_walker_7))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L48
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.B2i32(v500 == int32(0))|base.B2i32(v500 == int32(100)) != 0 {
		goto L7
	} else {
		goto L163
	}
L155:
	;
	v485 = F_getBaseType(m, v483)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L48
	} else {
		goto L156
	}
L156:
	;
	v487 = F_get_typ_typrelid(m, v485)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L48
	} else {
		goto L157
	}
L157:
	;
	if v487 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v490 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1259), v487, v490, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L48
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v495, int32(0), v497)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L48
	} else {
		goto L162
	}
L161:
	;
	goto L154
L162:
	;
	goto L154
L163:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v500, int32(0), v508)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L48
	} else {
		goto L164
	}
L164:
	;
	goto L7
L165:
	;
	if v512 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v514 == int32(0) {
		goto L7
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v570, int32(0), v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L48
	} else {
		goto L178
	}
L169:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v517 <= int32(0) {
		goto L7
	} else {
		goto L170
	}
L170:
	;
	v522 = v3
	goto L171
L171:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v514)+12))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v533+v522<<(uint(int32(2))%32))))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v538)+8))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v538)+12))
	if v541 <= v540 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	goto L7
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538)+12)) = v541 << (uint(int32(1)) % 32)
	v548 = F_repalloc(m, v539, v541*int32(24))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L48
	} else {
		goto L176
	}
L174:
	;
	v552 = v539
	v553 = v540
	goto L175
L175:
	;
	v556 = v553*int32(12) + v552
	*(*int32)(unsafe.Add(mBase, uint32(v556)+8)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v556)+4)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v556))) = int32(1259)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v538)+8))
	v562 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v538)+8)) = v561 + v562
	v566 = v522 + v562
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v566 < v567 {
		v522 = v566
		goto L171
	} else {
		goto L177
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v548
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v538)+8))
	v552 = v548
	v553 = v551
	goto L175
L177:
	;
	goto L172
L178:
	;
	goto L7
L179:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
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
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.B2i32(v598 == int32(0))|base.B2i32(v598 == int32(100)) != 0 {
		goto L7
	} else {
		goto L183
	}
L183:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v598, int32(0), v606)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L48
	} else {
		goto L184
	}
L184:
	;
	goto L7
L185:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.B2i32(v615 == int32(0))|base.B2i32(v615 == int32(100)) != 0 {
		goto L7
	} else {
		goto L186
	}
L186:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v615, int32(0), v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L48
	} else {
		goto L187
	}
L187:
	;
	goto L7
L188:
	;
	goto L7
L189:
	;
	goto L7
L190:
	;
	goto L7
L191:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v713 == int32(0) {
		goto L7
	} else {
		goto L201
	}
L192:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v644)+4))
	if v647 <= int32(0) {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v652 = v3
	goto L194
L194:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v644)+12))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v663+v652<<(uint(int32(2))%32))))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v668)))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v668)+8))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v668)+12))
	if v671 <= v670 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	goto L191
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v668)+12)) = v671 << (uint(int32(1)) % 32)
	v678 = F_repalloc(m, v669, v671*int32(24))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L48
	} else {
		goto L199
	}
L197:
	;
	v682 = v669
	v683 = v670
	goto L198
L198:
	;
	v686 = v683*int32(12) + v682
	*(*int32)(unsafe.Add(mBase, uint32(v686)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v686)+4)) = v667
	*(*int32)(unsafe.Add(mBase, uint32(v686))) = int32(2617)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v668)+8))
	v693 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v668)+8)) = v692 + v693
	v697 = v652 + v693
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v644)+4))
	if v697 < v698 {
		v652 = v697
		goto L194
	} else {
		goto L200
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v668))) = v678
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v668)+8))
	v682 = v678
	v683 = v681
	goto L198
L200:
	;
	goto L195
L201:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	if v716 <= int32(0) {
		goto L7
	} else {
		goto L202
	}
L202:
	;
	v722 = int32(0)
	goto L203
L203:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v713)+12))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v733+v722<<(uint(int32(2))%32))))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v738)))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v738)+8))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v738)+12))
	if v741 <= v740 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	goto L7
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v738)+12)) = v741 << (uint(int32(1)) % 32)
	v748 = F_repalloc(m, v739, v741*int32(24))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L48
	} else {
		goto L208
	}
L206:
	;
	v752 = v739
	v753 = v740
	goto L207
L207:
	;
	v756 = v753*int32(12) + v752
	*(*int32)(unsafe.Add(mBase, uint32(v756)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v756)+4)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v756))) = int32(2753)
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v738)+8))
	v763 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v738)+8)) = v762 + v763
	v767 = v722 + v763
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	if v767 < v768 {
		v722 = v767
		goto L203
	} else {
		goto L209
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v738))) = v748
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v738)+8))
	v752 = v748
	v753 = v751
	goto L207
L209:
	;
	goto L204
L210:
	;
	goto L7
L211:
	;
	goto L7
L212:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2606), v782, int32(0), v787)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L48
	} else {
		goto L213
	}
L213:
	;
	goto L7
L214:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v796 == int32(0) {
		v1474 = v3
		goto L5
	} else {
		goto L215
	}
L215:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v796, int32(0), v801)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L48
	} else {
		goto L216
	}
L216:
	;
	v1474 = v3
	goto L5
L217:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), v804, int32(0), v807)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L48
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v810 != 0 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	goto L219
L221:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1255), v810, int32(0), v813)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L48
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.B2i32(v816 == int32(0))|base.B2i32(v816 == int32(100)) != 0 {
		goto L7
	} else {
		goto L225
	}
L224:
	;
	goto L223
L225:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v816, int32(0), v824)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L48
	} else {
		goto L226
	}
L226:
	;
	goto L7
L227:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(1247), v827, int32(0), v830)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L48
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v833 != 0 {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	goto L229
L231:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(3456), v833, int32(0), v836)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L48
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v839 == int32(0) {
		goto L7
	} else {
		goto L235
	}
L234:
	;
	goto L233
L235:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_add_object_address(m, int32(2617), v839, int32(0), v844)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L48
	} else {
		goto L236
	}
L236:
	;
	goto L7
L237:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v987&int32(-2) != int32(2) {
		goto L262
	} else {
		goto L263
	}
L238:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v847)+4))
	if v850 <= int32(0) {
		goto L237
	} else {
		goto L239
	}
L239:
	;
	v858 = v3
	goto L240
L240:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v847)+12))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v866+v858<<(uint(int32(2))%32))))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v870)+12))
	switch v871 {
	case 0:
		goto L244
	default:
		goto L242
	case 2:
		goto L243
	case 7:
		goto L2
	}
L241:
	;
	goto L237
L242:
	;
	v971 = v858 + int32(1)
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v847)+4))
	if v971 < v972 {
		v858 = v971
		goto L240
	} else {
		goto L261
	}
L243:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v903 = F_lcons(m, v901, v902)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L48
	} else {
		goto L249
	}
L244:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v870)+16))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v873)))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v873)+8))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v873)+12))
	if v876 <= v875 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v873)+12)) = v876 << (uint(int32(1)) % 32)
	v883 = F_repalloc(m, v874, v876*int32(24))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L48
	} else {
		goto L248
	}
L246:
	;
	v887 = v874
	v888 = v875
	goto L247
L247:
	;
	v891 = v888*int32(12) + v887
	*(*int32)(unsafe.Add(mBase, uint32(v891)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v891)+4)) = v872
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = int32(1259)
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v873)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v873)+8)) = v897 + int32(1)
	goto L242
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v873))) = v883
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v873)+8))
	v887 = v883
	v888 = v886
	goto L247
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v903
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v870)+48))
	if int32(0) < v906 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v913 = v906
	v914 = int32(0)
	goto L253
L251:
	;
	v953 = v903
	goto L252
L252:
	;
	v954 = F_list_delete_first(m, v953)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L48
	} else {
		goto L260
	}
L253:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v870)+52))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v923)+12))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v924+v914<<(uint(int32(2))%32))))
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v928)))
	if v929 != int32(6) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v953 = v939
	goto L252
L255:
	;
	v932 = F_find_expr_references_walker(m, v928, l1)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L48
	} else {
		goto L258
	}
L256:
	;
	v935 = v913
	goto L257
L257:
	;
	v937 = v914 + int32(1)
	if v937 < v935 {
		v913 = v935
		v914 = v937
		goto L253
	} else {
		goto L259
	}
L258:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v870)+48))
	v935 = v934
	goto L257
L259:
	;
	goto L254
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v954
	goto L242
L261:
	;
	goto L241
L262:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v1088 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L263:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v992 <= int32(0) {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v995 == int32(0) {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v995)+4))
	if v998 < v992 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v995)+12))
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1000+v992<<(uint(int32(2))%32)-int32(4))))
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+12))
	if v1007 != 0 {
		goto L262
	} else {
		goto L267
	}
L267:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v1008 == int32(0) {
		goto L262
	} else {
		goto L268
	}
L268:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+4))
	if v1011 <= int32(0) {
		goto L262
	} else {
		goto L269
	}
L269:
	;
	v1017 = int32(0)
	goto L270
L270:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+12))
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1028+v1017<<(uint(int32(2))%32))))
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1032)+26)))
	if v1033 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	goto L262
L272:
	;
	v1036 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1032)+8)))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+16))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1038)))
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+8))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+12))
	if v1041 <= v1040 {
		goto L275
	} else {
		goto L276
	}
L273:
	;
	goto L274
L274:
	;
	v1072 = v1017 + int32(1)
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+4))
	if v1072 < v1073 {
		v1017 = v1072
		goto L270
	} else {
		goto L279
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1038)+12)) = v1041 << (uint(int32(1)) % 32)
	v1048 = F_repalloc(m, v1039, v1041*int32(24))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L48
	} else {
		goto L278
	}
L276:
	;
	v1052 = v1039
	v1053 = v1040
	goto L277
L277:
	;
	v1056 = v1053*int32(12) + v1052
	*(*int32)(unsafe.Add(mBase, uint32(v1056)+8)) = v1036
	*(*int32)(unsafe.Add(mBase, uint32(v1056)+4)) = v1037
	*(*int32)(unsafe.Add(mBase, uint32(v1056))) = int32(1259)
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1038)+8)) = v1061 + int32(1)
	goto L274
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1038))) = v1048
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+8))
	v1052 = v1048
	v1053 = v1051
	goto L277
L279:
	;
	goto L271
L280:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1160 = F_lcons(m, v1158, v1159)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L48
	} else {
		goto L290
	}
L281:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+4))
	if v1091 <= int32(0) {
		goto L280
	} else {
		goto L282
	}
L282:
	;
	v1097 = int32(0)
	goto L283
L283:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+12))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1108+v1097<<(uint(int32(2))%32))))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1113)))
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+8))
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+12))
	if v1116 <= v1115 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	goto L280
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1113)+12)) = v1116 << (uint(int32(1)) % 32)
	v1123 = F_repalloc(m, v1114, v1116*int32(24))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L48
	} else {
		goto L288
	}
L286:
	;
	v1127 = v1114
	v1128 = v1115
	goto L287
L287:
	;
	v1131 = v1128*int32(12) + v1127
	*(*int32)(unsafe.Add(mBase, uint32(v1131)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1131)+4)) = v1112
	*(*int32)(unsafe.Add(mBase, uint32(v1131))) = int32(2606)
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+8))
	v1138 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1113)+8)) = v1137 + v1138
	v1142 = v1097 + v1138
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+4))
	if v1142 < v1143 {
		v1097 = v1142
		goto L283
	} else {
		goto L289
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1113))) = v1123
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+8))
	v1127 = v1123
	v1128 = v1126
	goto L287
L289:
	;
	goto L284
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1160
	v1165 = F_query_tree_walker_impl(m, l0, int32(498), l1, int32(132))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L48
	} else {
		goto L291
	}
L291:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1168 = F_list_delete_first(m, v1167)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L48
	} else {
		goto L292
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1168
	v1474 = v1165
	goto L5
L293:
	;
	goto L7
L294:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1243 == int32(0) {
		goto L7
	} else {
		goto L304
	}
L295:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+4))
	if v1177 <= int32(0) {
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1182 = v3
	goto L297
L297:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+12))
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1193+v1182<<(uint(int32(2))%32))))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1198)))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+8))
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+12))
	if v1201 <= v1200 {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	goto L294
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1198)+12)) = v1201 << (uint(int32(1)) % 32)
	v1208 = F_repalloc(m, v1199, v1201*int32(24))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L48
	} else {
		goto L302
	}
L300:
	;
	v1212 = v1199
	v1213 = v1200
	goto L301
L301:
	;
	v1216 = v1213*int32(12) + v1212
	*(*int32)(unsafe.Add(mBase, uint32(v1216)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1216)+4)) = v1197
	*(*int32)(unsafe.Add(mBase, uint32(v1216))) = int32(1247)
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+8))
	v1223 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1198)+8)) = v1222 + v1223
	v1227 = v1182 + v1223
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+4))
	if v1227 < v1228 {
		v1182 = v1227
		goto L297
	} else {
		goto L303
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1198))) = v1208
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+8))
	v1212 = v1208
	v1213 = v1211
	goto L301
L303:
	;
	goto L298
L304:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1243)+4))
	if v1246 <= int32(0) {
		goto L7
	} else {
		goto L305
	}
L305:
	;
	v1254 = int32(0)
	goto L306
L306:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1243)+12))
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1263+v1254<<(uint(int32(2))%32))))
	v1268 = int32(0)
	if base.B2i32(v1267 == v1268)|base.B2i32(v1267 == int32(100)) == v1268 {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	goto L7
L308:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1275)))
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+8))
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+12))
	if v1278 <= v1277 {
		goto L311
	} else {
		goto L312
	}
L309:
	;
	goto L310
L310:
	;
	v1308 = v1254 + int32(1)
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1243)+4))
	if v1308 < v1309 {
		v1254 = v1308
		goto L306
	} else {
		goto L315
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+12)) = v1278 << (uint(int32(1)) % 32)
	v1285 = F_repalloc(m, v1276, v1278*int32(24))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L48
	} else {
		goto L314
	}
L312:
	;
	v1289 = v1276
	v1290 = v1277
	goto L313
L313:
	;
	v1293 = v1290*int32(12) + v1289
	*(*int32)(unsafe.Add(mBase, uint32(v1293)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1293)+4)) = v1267
	*(*int32)(unsafe.Add(mBase, uint32(v1293))) = int32(3456)
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+8)) = v1299 + int32(1)
	goto L310
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1275))) = v1285
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+8))
	v1289 = v1285
	v1290 = v1288
	goto L313
L315:
	;
	goto L307
L316:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1380 == int32(0) {
		goto L7
	} else {
		goto L326
	}
L317:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+4))
	if v1314 <= int32(0) {
		goto L316
	} else {
		goto L318
	}
L318:
	;
	v1319 = v3
	goto L319
L319:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+12))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1330+v1319<<(uint(int32(2))%32))))
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1335)))
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+8))
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+12))
	if v1338 <= v1337 {
		goto L321
	} else {
		goto L322
	}
L320:
	;
	goto L316
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1335)+12)) = v1338 << (uint(int32(1)) % 32)
	v1345 = F_repalloc(m, v1336, v1338*int32(24))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L48
	} else {
		goto L324
	}
L322:
	;
	v1349 = v1336
	v1350 = v1337
	goto L323
L323:
	;
	v1353 = v1350*int32(12) + v1349
	*(*int32)(unsafe.Add(mBase, uint32(v1353)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1353)+4)) = v1334
	*(*int32)(unsafe.Add(mBase, uint32(v1353))) = int32(1247)
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+8))
	v1360 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1335)+8)) = v1359 + v1360
	v1364 = v1319 + v1360
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+4))
	if v1364 < v1365 {
		v1319 = v1364
		goto L319
	} else {
		goto L325
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1335))) = v1345
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+8))
	v1349 = v1345
	v1350 = v1348
	goto L323
L325:
	;
	goto L320
L326:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+4))
	if v1383 <= int32(0) {
		goto L7
	} else {
		goto L327
	}
L327:
	;
	v1391 = int32(0)
	goto L328
L328:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+12))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1400+v1391<<(uint(int32(2))%32))))
	v1405 = int32(0)
	if base.B2i32(v1404 == v1405)|base.B2i32(v1404 == int32(100)) == v1405 {
		goto L330
	} else {
		goto L331
	}
L329:
	;
	goto L7
L330:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1412)))
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+8))
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+12))
	if v1415 <= v1414 {
		goto L333
	} else {
		goto L334
	}
L331:
	;
	goto L332
L332:
	;
	v1445 = v1391 + int32(1)
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+4))
	if v1445 < v1446 {
		v1391 = v1445
		goto L328
	} else {
		goto L337
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1412)+12)) = v1415 << (uint(int32(1)) % 32)
	v1422 = F_repalloc(m, v1413, v1415*int32(24))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L48
	} else {
		goto L336
	}
L334:
	;
	v1426 = v1413
	v1427 = v1414
	goto L335
L335:
	;
	v1430 = v1427*int32(12) + v1426
	*(*int32)(unsafe.Add(mBase, uint32(v1430)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1430)+4)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v1430))) = int32(3456)
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1412)+8)) = v1436 + int32(1)
	goto L332
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1412))) = v1422
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+8))
	v1426 = v1422
	v1427 = v1425
	goto L335
L337:
	;
	goto L329
L338:
	;
	goto L7
L339:
	;
	v1474 = v1468
	goto L5
L340:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v1492
	F_errmsg_internal(m, int32(_a_F_find_expr_references_walker_10), v16)
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L48
	} else {
		goto L341
	}
L341:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(1838), int32(_a_F_find_expr_references_walker_7))
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L48
	} else {
		goto L342
	}
L342:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L343:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v1507
	F_errmsg_internal(m, int32(_a_F_find_expr_references_walker_11), v16+int32(16))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L48
	} else {
		goto L344
	}
L344:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(1841), int32(_a_F_find_expr_references_walker_7))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L48
	} else {
		goto L345
	}
L345:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L346:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L48
	} else {
		goto L347
	}
L347:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v870)+8))
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v1527
	F_errmsg(m, int32(_a_F_find_expr_references_walker_12), v16+int32(80))
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L48
	} else {
		goto L348
	}
L348:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(2344), int32(_a_F_find_expr_references_walker_7))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L48
	} else {
		goto L349
	}
L349:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L350:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v1544
	F_errmsg_internal(m, int32(_a_F_find_expr_references_walker_13), v16-int32(-64))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L48
	} else {
		goto L351
	}
L351:
	;
	F_errfinish(m, int32(_a_F_find_expr_references_walker_1), int32(2367), int32(_a_F_find_expr_references_walker_7))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L48
	} else {
		goto L352
	}
L352:
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
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
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
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+288))
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
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
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
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+292))
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l2
	v7 = F_transformExprRecurse(m, l0, l1)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v5
		return v7
	}
}
